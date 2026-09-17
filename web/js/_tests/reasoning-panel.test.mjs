// Panneau Raisonnement : bloc infos requete, auto-masquage, bouton par
// requete, loader rond de streaming. (Refs : CETAS complet / Marexcode.)
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
const translateCalls = [];
let translateShouldFail = false;
globalThis.fetch = async (url, opts = {}) => {
  if (String(url).startsWith("/api/model-info")) {
    return {
      ok: true,
      status: 200,
      text: async () =>
        JSON.stringify({ context_window: 131072, input_per_1m: 0.15, output_per_1m: 0.6 }),
    };
  }
  if (String(url) === "/api/deepthink/translate" && (opts.method || "GET") === "POST") {
    const body = JSON.parse(opts.body || "{}");
    translateCalls.push(body.text);
    if (translateShouldFail) {
      return { ok: false, status: 502, text: async () => JSON.stringify({ error: "boom" }) };
    }
    return {
      ok: true,
      status: 200,
      text: async () => JSON.stringify({ translation: "TRADUIT: " + body.text }),
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
      assert.equal(document.querySelector(".reason-sec-label").textContent, "REASONING");
      const tbtn = document.querySelector(".reason-translate-btn");
      assert.ok(tbtn, "bouton de traduction présent dans l'en-tête REASONING");
      assert.equal(tbtn.textContent, "🌐 Traduire");
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

  test("finalize + restore : instantané réutilisable par le bouton de la réponse", () => {
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

describe("thread-view : bouton par réponse + loader rond", () => {
  test("cycle complet : bouton caché → visible, auto-hide, loader, snapshot", async () => {
    const log = setupPanelDom();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "salut" });
      const btnBefore = log.querySelector(".reason-btn");
      assert.ok(btnBefore, "bouton existe (caché) avant la réponse");
      assert.equal(btnBefore.hidden, true, "bouton caché avant la réponse");

      view.handleEvent({ route: { label: "MiMo V2.5", provider: "opencode", model: "mimo" } });
      view.handleEvent({ reasoning_content: "je réfléchis…" });
      assert.ok(document.getElementById("reason-panel").classList.contains("open"), "panneau ouvert");

      view.handleEvent({ content: "bonjour" });
      const bubble = log.querySelector(".message-assistant");
      assert.ok(bubble, "bulle assistant créée");
      // Le bouton Raisonnement est dans le wrapper user (pas l'assistant).
      const wrapper = log.querySelector(".message-wrapper-user");
      assert.ok(wrapper, "wrapper user");
      const btn = wrapper.querySelector(".reason-btn");
      assert.ok(btn, "bouton Raisonnement dans le wrapper user");
      assert.equal(btn.hidden, false, "visible dès le premier delta");
      assert.equal(btn.textContent, "Raisonnement");

      await nextFrame();
      assert.ok(!document.getElementById("reason-panel").classList.contains("open"), "panneau auto-masqué");
      const loader = log.querySelector(".marex-loader");
      assert.ok(loader, "loader rond présent pendant le stream");
      assert.ok(loader.querySelector(".loader__inner"), "point central présent");
      assert.equal(loader.querySelectorAll(".loader__dot").length, 4, "4 satellites");
      const liveStats = log.querySelector(".gen-stats-live");
      assert.ok(liveStats, "stats live présentes pendant le stream");
      assert.match(liveStats.textContent, /^\d+s/, "format 'Ns'");

      view.handleEvent({ stats: { prompt_tokens: 443, completion_tokens: 95 } });
      view.handleEvent({ turn_done: { elapsed_ms: 5000 } });
      assert.equal(view.assistantBody, null, "état assistant réinitialisé");
      assert.equal(log.querySelector(".marex-loader"), null, "loader retiré en fin de tour");
      assert.equal(log.querySelector(".gen-stats-live"), null, "stats live retirées en fin de tour");
      const uwrapper = log.querySelector(".message-wrapper-user");
      assert.ok(uwrapper._reasonSnap, "instantané stocké sur la réponse user");
      assert.ok(uwrapper._reasonSnap.text.includes("je réfléchis"));
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
      assert.ok(btn, "bouton créé dans la bulle assistant");
      assert.equal(btn.hidden, true, "reste caché sans raisonnement");
      assert.equal(document.querySelector(".reason-info"), null, "pas de bloc infos");
      const awrapper = log.querySelector(".message-wrapper-assistant");
      assert.equal(awrapper._reasonSnap, undefined, "pas d'instantané");
    } finally {
      view.reset();
      panelMod.resetReasonPanel();
      document.body.innerHTML = "";
    }
  });

  test("le loader survit au re-rendu markdown (hors du corps)", async () => {
    const log = setupPanelDom();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "test" });
      view.handleEvent({ content: "premier" });
      await nextFrame();
      await nextFrame();
      assert.ok(log.querySelector(".marex-loader"), "loader présent après rendu");
      // Le moteur markdown ne touche qu'au corps : on vide le corps comme
      // le ferait un re-rendu complet, le loader (en tête de bulle) survit.
      const body = log.querySelector(".message-assistant .message-text");
      assert.ok(body, "corps présent");
      body.innerHTML = "";
      assert.ok(log.querySelector(".marex-loader"), "loader intact après vidage du corps");
      // Les stats live suivent le texte : format "Ns · X tok".
      const live = log.querySelector(".gen-stats-live");
      assert.ok(live, "stats live présentes");
      assert.match(live.textContent, /^\d+s/, "format 'Ns'");
      view.handleEvent({ turn_done: {} });
      assert.equal(log.querySelector(".marex-loader"), null, "loader retiré");
      assert.equal(log.querySelector(".gen-stats-live"), null, "stats live retirées");
    } finally {
      view.reset();
      panelMod.resetReasonPanel();
      document.body.innerHTML = "";
    }
  });

  test("showWait affiche le loader rond (aucun timer braille)", async () => {
    const log = setupPanelDom();
    const view = makeView(log);
    try {
      view.showWait();
      const loader = log.querySelector(".stream-waiting .marex-loader");
      assert.ok(loader, "loader rond dans la zone d'attente");
      assert.equal(view.waitTimer, undefined, "plus de timer d'attente");
      view.hideWait();
      assert.equal(log.querySelector(".stream-waiting"), null, "zone d'attente retirée");
    } finally {
      view.reset();
      panelMod.resetReasonPanel();
      document.body.innerHTML = "";
    }
  });
});


describe("panneau raisonnement : traduction DeepThink", () => {
  let translateCalls;
  let translateShouldFail;

  // Mock fetch : compte les appels de traduction, jamais de remplacement.
  function mockFetchTranslate() {
    translateCalls = [];
    translateShouldFail = false;
    globalThis.fetch = async (url, opts) => {
      if (String(url).includes("/api/deepthink/translate")) {
        const body = JSON.parse(opts.body);
        translateCalls.push(body.text);
        if (translateShouldFail) {
          return {
            ok: false,
            status: 500,
            statusText: "boom",
            text: async () => "boom",
          };
        }
        return {
          ok: true,
          text: async () => JSON.stringify({ translation: "TRADUIT: " + body.text, lang: "fr" }),
        };
      }
      throw new Error("fetch inattendu: " + url);
    };
  }

  function setupTranslated() {
    mockFetchTranslate();
    localStorage.removeItem("cetas.deepthink.auto");
    panelMod._setAutoTranslateDelayForTests(5000);
    setupPanelDom();
    panelMod.resetReasonPanel();
    panelMod.beginReasonTurn({ label: "MiMo", provider: "opencode", model: "xiaomi/mimo-v2.5" });
    panelMod.appendReasoningPanel("The user says hello.");
  }

  function teardownTranslated() {
    panelMod.resetReasonPanel();
    localStorage.removeItem("cetas.deepthink.auto");
    panelMod._setAutoTranslateDelayForTests(5000);
    document.body.innerHTML = "";
  }

  function clickTranslateBtn() {
    const btn = document.querySelector(".reason-translate-btn");
    assert.ok(btn, "bouton Traduire présent");
    btn.click();
  }

  const wait = (ms) => new Promise((r) => setTimeout(r, ms));

  test("clic : traduction affichée SOUS l'original, jamais de remplacement, auto activé", async () => {
    setupTranslated();
    try {
      clickTranslateBtn();
      await wait(50);
      assert.equal(translateCalls.length, 1, "un appel /api/deepthink/translate");
      assert.equal(translateCalls[0], "The user says hello.", "texte original envoyé");
      // L'original est intact…
      const t = document.querySelector(".reason-text");
      assert.equal(t.textContent, "The user says hello.", "original non remplacé");
      // …et la traduction est affichée en dessous.
      const box = document.querySelector(".reason-translation");
      assert.ok(box, "bloc de traduction présent");
      assert.match(box.textContent, /TRADUCTION · FR/, "en-tête avec la langue");
      assert.match(box.textContent, /TRADUIT: The user says hello\./, "traduction affichée");
      // Le bouton est passé en mode auto.
      const btn = document.querySelector(".reason-translate-btn");
      assert.equal(btn.textContent, "🌐 Auto ✓", "bouton en mode auto");
      assert.equal(localStorage.getItem("cetas.deepthink.auto"), "1", "auto persisté");
    } finally {
      teardownTranslated();
    }
  });

  test("second clic : désactive l'auto, la traduction affichée reste", async () => {
    setupTranslated();
    try {
      clickTranslateBtn();
      await wait(50);
      assert.ok(document.querySelector(".reason-translation"), "traduction affichée");
      clickTranslateBtn();
      await wait(20);
      const btn = document.querySelector(".reason-translate-btn");
      assert.equal(btn.textContent, "🌐 Traduire", "bouton revenu à l'état initial");
      assert.equal(localStorage.getItem("cetas.deepthink.auto"), null, "auto désactivé");
      assert.equal(translateCalls.length, 1, "aucun nouvel appel de traduction");
      assert.ok(document.querySelector(".reason-translation"), "traduction conservée à l'écran");
    } finally {
      teardownTranslated();
    }
  });

  test("auto : 5s après la fin du raisonnement, traduction automatique sous l'original", async () => {
    setupTranslated();
    panelMod._setAutoTranslateDelayForTests(40);
    localStorage.setItem("cetas.deepthink.auto", "1");
    try {
      panelMod.finishReasoning();
      assert.equal(translateCalls.length, 0, "pas de traduction immédiate");
      await wait(200);
      assert.equal(translateCalls.length, 1, "traduction auto déclenchée après le délai");
      const t = document.querySelector(".reason-text");
      assert.equal(t.textContent, "The user says hello.", "original intact");
      const box = document.querySelector(".reason-translation");
      assert.ok(box, "traduction auto affichée");
      assert.match(box.textContent, /TRADUIT: The user says hello\./);
    } finally {
      teardownTranslated();
    }
  });

  test("auto : si le texte a changé entre-temps, la traduction est annulée", async () => {
    setupTranslated();
    panelMod._setAutoTranslateDelayForTests(40);
    localStorage.setItem("cetas.deepthink.auto", "1");
    try {
      panelMod.finishReasoning();
      // Un nouveau contenu arrive avant la fin du délai (tour suivant).
      panelMod.appendReasoningPanel("Contenu du tour suivant.");
      await wait(200);
      assert.equal(translateCalls.length, 0, "traduction annulée (texte différent)");
      assert.equal(document.querySelector(".reason-translation"), null, "aucun bloc affiché");
    } finally {
      teardownTranslated();
    }
  });

  test("auto : pas de double traduction si déjà traduite (clic manuel)", async () => {
    setupTranslated();
    panelMod._setAutoTranslateDelayForTests(40);
    try {
      clickTranslateBtn(); // traduit maintenant + active l'auto
      await wait(50);
      assert.equal(translateCalls.length, 1, "traduction manuelle faite");
      panelMod.finishReasoning(); // fin du même raisonnement
      await wait(200);
      assert.equal(translateCalls.length, 1, "pas de second appel (déjà traduite)");
    } finally {
      teardownTranslated();
    }
  });

  test("échec de traduction : bloc d'erreur sous l'original, original intact", async () => {
    setupTranslated();
    try {
      translateShouldFail = true;
      clickTranslateBtn();
      await wait(50);
      const t = document.querySelector(".reason-text");
      assert.equal(t.textContent, "The user says hello.", "original intact après échec");
      const box = document.querySelector(".reason-translation-error");
      assert.ok(box, "bloc d'erreur affiché");
      assert.match(box.textContent, /⚠/, "marqueur d'erreur présent");
    } finally {
      teardownTranslated();
    }
  });

  test("le bouton reflète le mode auto persisté à la reconstruction du bloc", async () => {
    setupTranslated();
    try {
      localStorage.setItem("cetas.deepthink.auto", "1");
      panelMod.resetReasonPanel();
      panelMod.beginReasonTurn({ label: "MiMo", provider: "opencode", model: "xiaomi/mimo-v2.5" });
      panelMod.appendReasoningPanel("Bonjour.");
      const btn = document.querySelector(".reason-translate-btn");
      assert.equal(btn.textContent, "🌐 Auto ✓", "mode auto restauré après rebuild");
    } finally {
      teardownTranslated();
    }
  });

  test("finalizeReasonTurn : l'instantané garde toujours le texte original", async () => {
    setupTranslated();
    try {
      clickTranslateBtn();
      await wait(50);
      const snap = panelMod.finalizeReasonTurn(1200);
      assert.equal(snap.text, "The user says hello.", "instantané = original");
    } finally {
      teardownTranslated();
    }
  });
});
