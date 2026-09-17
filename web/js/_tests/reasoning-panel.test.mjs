// Panneau Raisonnement : bloc infos requete, auto-masquage, bouton par
// requete, spinner de streaming. (Refs : CETAS complet / Marexcode.)
import { test, describe } from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

// jsdom : résolution standard puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  try {
    ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
  } catch {
    console.log("jsdom indisponible, tests ignorés");
    process.exit(0);
  }
}

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.cancelAnimationFrame = dom.window.cancelAnimationFrame.bind(dom.window);
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;

// Mock fetch avec text() (api() lit resp.text()).
globalThis.fetch = async (url) => {
  if (String(url).startsWith("/api/model-info")) {
    return {
      ok: true,
      status: 200,
      text: async () =>
        JSON.stringify({ context_window: 131072, input_per_1m: 0.15, output_per_1m: 0.6 }),
    };
  }
  return { ok: false, status: 404, text: async () => JSON.stringify({ error: "nope" }) };
};

const panelMod = await import("../reasoning-panel.js");
const { ThreadView } = await import("../thread-view.js");

const nextFrame = () =>
  new Promise((resolve) => {
    dom.window.requestAnimationFrame(() => resolve());
  });
const flush = () => new Promise((r) => setTimeout(r, 20));

function setupPanelDom() {
  document.body.innerHTML = `
    <aside class="reason-panel" id="reason-panel">
      <div class="reason-panel-header">
        <span class="reason-panel-title">
          <span class="reason-spinner" id="reason-spinner" style="display:none"></span>
          <span>Raisonnement</span>
        </span>
        <button class="reason-panel-close" id="reason-panel-close"></button>
      </div>
      <div class="reason-panel-body" id="reason-panel-body"></div>
    </aside>
    <div id="log"></div>`;
  panelMod.initReasonPanel();
  return document.getElementById("log");
}

function makeView(log) {
  return new ThreadView({
    log,
    streamURL: () => "/x",
    sendURL: "/x",
    stopURL: "/x",
    getPayload: (t) => ({ message: t }),
    actions: false,
    reasonPanel: true,
  });
}

describe("panneau raisonnement : bloc infos", () => {
  test("le premier delta construit le bloc infos + REASONING et ouvre le panneau", () => {
    setupPanelDom();
    panelMod.resetReasonPanel();
    panelMod.beginReasonTurn({ label: "MiMo V2.5", provider: "opencode", model: "mimo" });
    panelMod.appendReasoningPanel("je réfléchis");
    try {
      const p = document.getElementById("reason-panel");
      assert.ok(p.classList.contains("open"), "le panneau s'ouvre");
      const info = document.querySelector(".reason-info");
      assert.ok(info, "bloc infos présent");
      const rows = info.querySelectorAll(".ri-row");
      assert.equal(rows.length, 6, "6 lignes d'infos");
      const labels = [...rows].map((r) => r.querySelector(".ri-k").textContent);
      assert.deepEqual(labels, [
        "Modèle : ",
        "Temps de génération : ",
        "Input : ",
        "Output - t/s : ",
        "Contexte : ",
        "Coût : ",
      ]);
      assert.equal(document.querySelector(".ri-model").textContent, "MiMo V2.5");
      assert.equal(document.querySelector(".reason-sec").textContent, "REASONING");
      assert.ok(document.querySelector(".reason-text").textContent.includes("je réfléchis"));
    } finally {
      panelMod.resetReasonPanel();
    }
  });

  test("updateReasonUsage remplit input/output/contexte/coût", async () => {
    setupPanelDom();
    panelMod.resetReasonPanel();
    panelMod.beginReasonTurn({ label: "MiMo V2.5", provider: "opencode", model: "mimo" });
    panelMod.setReasonModel("MiMo V2.5", "opencode", "mimo");
    panelMod.appendReasoningPanel("x");
    panelMod.updateReasonUsage({ prompt_tokens: 443, completion_tokens: 95 });
    await flush(); // fetch model-info mocké
    try {
      assert.ok(document.querySelector(".ri-input").textContent.includes("443"));
      const out = document.querySelector(".ri-output").textContent;
      assert.ok(out.includes("95"), "output : " + out);
      assert.ok(out.includes("/s"), "débit : " + out);
      const ctx = document.querySelector(".ri-ctx").textContent;
      assert.ok(ctx.includes("131K"), "contexte max : " + ctx);
      assert.ok(ctx.includes("utilisé"), "pourcentage : " + ctx);
      const cost = document.querySelector(".ri-cost").textContent;
      assert.ok(cost.startsWith("$"), "coût : " + cost);
    } finally {
      panelMod.resetReasonPanel();
    }
  });

  test("sans raisonnement (Thinking désactivé), les stats ne créent pas le bloc", () => {
    setupPanelDom();
    panelMod.resetReasonPanel();
    panelMod.beginReasonTurn({ label: "M", provider: "p", model: "m" });
    panelMod.updateReasonUsage({ prompt_tokens: 10, completion_tokens: 5 });
    try {
      assert.equal(document.querySelector(".reason-info"), null, "pas de bloc infos");
      assert.equal(document.querySelector(".reason-text"), null, "pas de texte");
    } finally {
      panelMod.resetReasonPanel();
    }
  });

  test("finishReasoning masque automatiquement le panneau", () => {
    setupPanelDom();
    panelMod.resetReasonPanel();
    panelMod.beginReasonTurn({});
    panelMod.appendReasoningPanel("raisonnement...");
    assert.ok(document.getElementById("reason-panel").classList.contains("open"));
    panelMod.finishReasoning();
    try {
      assert.ok(!document.getElementById("reason-panel").classList.contains("open"), "auto-masqué");
      assert.equal(document.getElementById("reason-spinner").style.display, "none", "spinner arrêté");
    } finally {
      panelMod.resetReasonPanel();
    }
  });

  test("finalize + restore : instantané réutilisable par le bouton de la requête", () => {
    setupPanelDom();
    panelMod.resetReasonPanel();
    panelMod.beginReasonTurn({ label: "MiMo V2.5", provider: "opencode", model: "mimo" });
    panelMod.appendReasoningPanel("raisonnement complet");
    panelMod.updateReasonUsage({ prompt_tokens: 443, completion_tokens: 95 });
    panelMod.finishReasoning();
    const snap = panelMod.finalizeReasonTurn(5000);
    assert.ok(snap.text.includes("raisonnement complet"));
    assert.equal(snap.elapsedMs, 5000);
    assert.equal(snap.inputTok, 443);
    // Nouveau tour : le panneau est vide…
    panelMod.resetReasonPanel();
    assert.equal(document.querySelector(".reason-text"), null);
    // …puis restauré via le bouton.
    panelMod.restoreReasonSnapshot(snap);
    try {
      assert.ok(document.getElementById("reason-panel").classList.contains("open"));
      assert.ok(document.querySelector(".reason-text").textContent.includes("raisonnement complet"));
      assert.equal(document.querySelector(".ri-model").textContent, "MiMo V2.5");
      assert.ok(document.querySelector(".ri-input").textContent.includes("443"));
      assert.ok(document.querySelector(".ri-time").textContent.includes("5"), "temps figé : " + document.querySelector(".ri-time").textContent);
    } finally {
      panelMod.resetReasonPanel();
    }
  });
});

describe("thread-view : bouton par requête + spinner de streaming", () => {
  test("cycle complet : bouton caché → visible, auto-hide, spinner, snapshot", async () => {
    const log = setupPanelDom();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "salut" });
      const wrapper = log.querySelector(".message-wrapper-user");
      assert.ok(wrapper, "message utilisateur créé");
      const btn = wrapper.querySelector(".reason-btn");
      assert.ok(btn, "bouton Raisonnement présent");
      assert.equal(btn.hidden, true, "caché avant raisonnement");
      assert.equal(btn.textContent, "Raisonnement");

      view.handleEvent({ route: { label: "MiMo V2.5", provider: "opencode", model: "mimo" } });
      view.handleEvent({ reasoning_content: "je réfléchis…" });
      assert.equal(btn.hidden, false, "visible dès le premier delta");
      assert.ok(document.getElementById("reason-panel").classList.contains("open"), "panneau ouvert");

      view.handleEvent({ content: "bonjour" });
      await nextFrame();
      assert.ok(!document.getElementById("reason-panel").classList.contains("open"), "panneau auto-masqué");
      const spinner = view.assistantBody.querySelector(".stream-spinner");
      assert.ok(spinner, "spinner de streaming présent pendant le stream");

      view.handleEvent({ stats: { prompt_tokens: 443, completion_tokens: 95 } });
      view.handleEvent({ turn_done: { elapsed_ms: 5000 } });
      assert.equal(view.assistantBody, null, "état assistant réinitialisé");
      assert.equal(log.querySelector(".stream-spinner"), null, "spinner retiré en fin de tour");
      assert.ok(wrapper._reasonSnap, "instantané stocké sur la requête");
      assert.ok(wrapper._reasonSnap.text.includes("je réfléchis"));
      assert.equal(btn.hidden, false, "bouton toujours visible");

      // Clic : le panneau se rouvre avec l'instantané du tour.
      panelMod.resetReasonPanel();
      btn.click();
      assert.ok(document.getElementById("reason-panel").classList.contains("open"), "panneau rouvert");
      assert.ok(document.querySelector(".reason-text").textContent.includes("je réfléchis"));
      assert.equal(document.querySelector(".ri-model").textContent, "MiMo V2.5");
    } finally {
      view.reset();
      panelMod.resetReasonPanel();
      document.body.innerHTML = "";
    }
  });

  test("sans thinking : pas de bouton visible, pas de panneau", () => {
    const log = setupPanelDom();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "salut" });
      view.handleEvent({ content: "réponse directe" });
      view.handleEvent({ turn_done: { elapsed_ms: 1000 } });
      const btn = log.querySelector(".reason-btn");
      assert.ok(btn, "bouton créé");
      assert.equal(btn.hidden, true, "reste caché sans raisonnement");
      assert.equal(document.querySelector(".reason-info"), null, "pas de bloc infos");
      const wrapper = log.querySelector(".message-wrapper-user");
      assert.equal(wrapper._reasonSnap, undefined, "pas d'instantané");
    } finally {
      view.reset();
      panelMod.resetReasonPanel();
      document.body.innerHTML = "";
    }
  });

  test("le spinner survit au re-rendu markdown (re-ancrage)", async () => {
    const log = setupPanelDom();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "test" });
      view.handleEvent({ content: "premier" });
      await nextFrame();
      await nextFrame();
      assert.ok(log.querySelector(".stream-spinner"), "spinner présent après rendu");
      // Le moteur markdown retire les enfants excédentaires : on simule en
      // le retirant à la main, le tick suivant le ré-ancre.
      const sp = log.querySelector(".stream-spinner");
      sp.remove();
      await new Promise((r) => setTimeout(r, 120));
      assert.ok(log.querySelector(".stream-spinner"), "spinner ré-ancré par le tick");
      view.handleEvent({ turn_done: {} });
      assert.equal(log.querySelector(".stream-spinner"), null, "spinner retiré");
    } finally {
      view.reset();
      panelMod.resetReasonPanel();
      document.body.innerHTML = "";
    }
  });
});
