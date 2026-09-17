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
      assert.equal(log.querySelector(".reason-btn"), null, "pas de bouton avant la réponse");

      view.handleEvent({ route: { label: "MiMo V2.5", provider: "opencode", model: "mimo" } });
      view.handleEvent({ reasoning_content: "je réfléchis…" });
      assert.ok(document.getElementById("reason-panel").classList.contains("open"), "panneau ouvert");

      view.handleEvent({ content: "bonjour" });
      const bubble = log.querySelector(".message-assistant");
      assert.ok(bubble, "bulle assistant créée");
      const btn = bubble.querySelector(".reason-btn");
      assert.ok(btn, "bouton Raisonnement dans la bulle assistant");
      assert.equal(bubble.firstChild, btn, "bouton en tête de la bulle");
      assert.equal(btn.textContent, "▸ Raisonnement");
      assert.equal(btn.hidden, false, "visible dès le premier delta");

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
      const awrapper = log.querySelector(".message-wrapper-assistant");
      assert.ok(awrapper._reasonSnap, "instantané stocké sur la réponse");
      assert.ok(awrapper._reasonSnap.text.includes("je réfléchis"));
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

  test("auto : le tour scellé est traduit même si un nouveau tour a commencé", async () => {
    setupTranslated();
    panelMod._setAutoTranslateDelayForTests(40);
    localStorage.setItem("cetas.deepthink.auto", "1");
    try {
      panelMod.finishReasoning(); // scelle le tour 1, programme son auto-traduction
      // Un nouveau tour commence avant la fin du délai : son bloc se crée
      // EN DESSOUS, le texte du tour 1 ne change plus -> son auto part.
      panelMod.beginReasonTurn({ label: "MiMo", provider: "opencode", model: "x" });
      panelMod.appendReasoningPanel("Contenu du tour suivant.");
      await wait(200);
      assert.equal(translateCalls.length, 1, "le tour 1 scellé est traduit");
      assert.equal(translateCalls[0], "The user says hello.", "texte du tour 1 envoyé");
      const turns = document.querySelectorAll(".reason-turn");
      assert.equal(turns.length, 2, "deux blocs dans l'historique");
      assert.ok(turns[0].querySelector(".reason-translation"), "traduction dans le bloc du tour 1");
      assert.equal(turns[1].querySelector(".reason-translation"), null, "pas de traduction pour le tour 2 en cours");
      assert.ok(
        turns[1].querySelector(".reason-text").textContent.includes("Contenu du tour suivant."),
        "texte du tour 2 intact"
      );
    } finally {
      teardownTranslated();
    }
  });

  test("auto : annulée si le bloc disparaît avant le délai (nouvelle conversation)", async () => {
    setupTranslated();
    panelMod._setAutoTranslateDelayForTests(40);
    localStorage.setItem("cetas.deepthink.auto", "1");
    try {
      panelMod.finishReasoning();
      panelMod.resetReasonPanel(); // nouvelle conversation : le bloc n'existe plus
      await wait(200);
      assert.equal(translateCalls.length, 0, "traduction annulée (bloc supprimé)");
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

describe("panneau raisonnement : historique des tours", () => {
  function setupHistory() {
    setupPanelDom();
    panelMod.resetReasonPanel();
    localStorage.removeItem("cetas.deepthink.auto");
    return document.getElementById("log");
  }

  function teardownHistory() {
    panelMod.resetReasonPanel();
    localStorage.removeItem("cetas.deepthink.auto");
    document.body.innerHTML = "";
  }

  const wait = (ms) => new Promise((r) => setTimeout(r, ms));

  test("nouveau tour : le précédent est conservé, le suivant s'ajoute en dessous", () => {
    setupHistory();
    try {
      panelMod.beginReasonTurn({ label: "M1", provider: "p", model: "m1" });
      panelMod.appendReasoningPanel("raisonnement du tour 1");
      panelMod.finishReasoning();
      // Nouveau message utilisateur : scelle le tour 1, ne vide pas.
      panelMod.sealReasonTurn();
      panelMod.beginReasonTurn({ label: "M2", provider: "p", model: "m2" });
      panelMod.appendReasoningPanel("raisonnement du tour 2");
      const turns = document.querySelectorAll(".reason-turn");
      assert.equal(turns.length, 2, "deux blocs dans l'historique");
      assert.ok(turns[0].querySelector(".reason-text").textContent.includes("tour 1"), "tour 1 conservé");
      assert.ok(turns[1].querySelector(".reason-text").textContent.includes("tour 2"), "tour 2 en dessous");
      assert.ok(turns[0].hasAttribute("data-sealed"), "tour 1 scellé");
      assert.ok(!turns[1].hasAttribute("data-sealed"), "tour 2 en cours");
      assert.equal(turns[0].querySelector(".ri-model").textContent, "M1", "infos du tour 1 intactes");
      assert.equal(turns[1].querySelector(".ri-model").textContent, "M2", "infos du tour 2");
    } finally {
      teardownHistory();
    }
  });

  test("chaque tour a son bouton Traduire (traduction ciblée)", async () => {
    const calls = [];
    globalThis.fetch = async (url, opts) => {
      if (String(url).includes("/api/deepthink/translate")) {
        const body = JSON.parse(opts.body);
        calls.push(body.text);
        return { ok: true, text: async () => JSON.stringify({ translation: "TRADUIT", lang: "fr" }) };
      }
      throw new Error("fetch inattendu: " + url);
    };
    setupHistory();
    try {
      panelMod.beginReasonTurn({ label: "M1" });
      panelMod.appendReasoningPanel("texte un");
      panelMod.finishReasoning();
      panelMod.sealReasonTurn();
      panelMod.beginReasonTurn({ label: "M2" });
      panelMod.appendReasoningPanel("texte deux");
      const btns = document.querySelectorAll(".reason-turn .reason-translate-btn");
      assert.equal(btns.length, 2, "un bouton par tour");
      btns[1].click(); // traduit le tour 2 uniquement
      await wait(50);
      assert.equal(calls.length, 1, "un seul appel");
      assert.equal(calls[0], "texte deux", "le texte du tour 2 est envoyé");
      const turns = document.querySelectorAll(".reason-turn");
      assert.equal(turns[0].querySelector(".reason-translation"), null, "tour 1 non traduit");
      assert.ok(turns[1].querySelector(".reason-translation"), "traduction dans le bloc du tour 2");
      assert.ok(turns[0].querySelector(".reason-text").textContent.includes("texte un"), "original tour 1 intact");
      assert.ok(turns[1].querySelector(".reason-text").textContent.includes("texte deux"), "original tour 2 intact");
    } finally {
      teardownHistory();
    }
  });

  test("bouton Raisonnement : défile vers le tour dans l'historique sans vider", () => {
    setupHistory();
    try {
      panelMod.beginReasonTurn({ label: "M1" });
      panelMod.appendReasoningPanel("raisonnement un");
      const snap = panelMod.finalizeReasonTurn(1000);
      assert.ok(snap.turnId, "instantané avec turnId");
      panelMod.finishReasoning();
      panelMod.sealReasonTurn();
      panelMod.beginReasonTurn({ label: "M2" });
      panelMod.appendReasoningPanel("raisonnement deux");
      panelMod.finishReasoning();
      assert.equal(document.querySelectorAll(".reason-turn").length, 2);
      panelMod.restoreReasonSnapshot(snap);
      assert.equal(document.querySelectorAll(".reason-turn").length, 2, "historique conservé");
      const target = document.querySelector('.reason-turn[data-turn-id="' + snap.turnId + '"]');
      assert.ok(target, "bloc du tour retrouvé");
      assert.ok(target.classList.contains("reason-flash"), "surlignage du tour ciblé");
      assert.ok(target.querySelector(".reason-text").textContent.includes("raisonnement un"));
      assert.ok(document.getElementById("reason-panel").classList.contains("open"), "panneau rouvert");
    } finally {
      teardownHistory();
    }
  });

  test("dropCurrentReasonTurn : retire le tour en cours, garde l'historique", () => {
    setupHistory();
    try {
      panelMod.beginReasonTurn({ label: "M1" });
      panelMod.appendReasoningPanel("tour un");
      panelMod.finishReasoning();
      panelMod.beginReasonTurn({ label: "M2" });
      panelMod.appendReasoningPanel("tour deux en cours");
      panelMod.dropCurrentReasonTurn();
      const turns = document.querySelectorAll(".reason-turn");
      assert.equal(turns.length, 1, "seul le tour en cours est retiré");
      assert.ok(turns[0].querySelector(".reason-text").textContent.includes("tour un"), "tour 1 conservé");
    } finally {
      teardownHistory();
    }
  });

  test("nouveau message via thread-view : le tour précédent n'est pas remplacé", () => {
    const log = setupHistory();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "q1" });
      view.handleEvent({ reasoning_content: "premier raisonnement" });
      view.handleEvent({ turn_done: {} });
      view.handleEvent({ user: "q2" }); // scelle le tour 1
      view.handleEvent({ reasoning_content: "deuxième raisonnement" });
      const turns = document.querySelectorAll(".reason-turn");
      assert.equal(turns.length, 2, "deux blocs, pas de remplacement");
      assert.ok(turns[0].querySelector(".reason-text").textContent.includes("premier raisonnement"));
      assert.ok(turns[1].querySelector(".reason-text").textContent.includes("deuxième raisonnement"));
    } finally {
      view.reset();
      teardownHistory();
    }
  });

  test("nouvelle conversation : l'historique est vidé", () => {
    const log = setupHistory();
    const view = makeView(log);
    try {
      view.handleEvent({ user: "q1" });
      view.handleEvent({ reasoning_content: "raisonnement un" });
      view.handleEvent({ turn_done: {} });
      view.handleEvent({ user: "q2" });
      view.handleEvent({ reasoning_content: "raisonnement deux" });
      view.handleEvent({ turn_done: {} });
      assert.equal(document.querySelectorAll(".reason-turn").length, 2, "deux tours dans l'historique");
      view.reset(); // nouvelle conversation
      assert.equal(document.querySelectorAll(".reason-turn").length, 0, "historique vidé");
      assert.ok(document.querySelector(".reason-panel-empty"), "état vide affiché");
    } finally {
      view.reset();
      teardownHistory();
    }
  });
});

await test("exports : tout ce que thread-view.js importe de reasoning-panel.js existe", async () => {
  // Non-régression : un import statique manquant casse tout le chat dans le
  // navigateur (« does not provide an export named ... »). Couvre sealReasonTurn
  // et dropCurrentReasonTurn ajoutés pour l'historique des tours.
  const fs = await import("node:fs");
  const src = fs.readFileSync(new URL("../thread-view.js", import.meta.url), "utf8");
  const m = src.match(/import\s*\{([^}]+)\}\s*from\s*"\.\/reasoning-panel\.js"/);
  assert.ok(m, "import de reasoning-panel.js trouvé dans thread-view.js");
  const names = m[1].split(",").map((s) => s.trim()).filter(Boolean);
  assert.ok(names.length > 0, "au moins un nom importé");
  for (const name of names) {
    assert.equal(typeof panelMod[name], "function", `reasoning-panel.js doit exporter ${name}`);
  }
});
