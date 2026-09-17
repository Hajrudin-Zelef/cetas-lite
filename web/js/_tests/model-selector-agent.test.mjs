import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

// Onglet "Sélecteur agent" : ne montre que les familles et modes agent
// (m.agent), avec les mêmes réglages 1 modèle / Fallback.

const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

const dom = new JSDOM("<!doctype html><html><head></head><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;
globalThis.localStorage = dom.window.localStorage;

const POOL = [
  { provider: "openrouter", model: "openrouter/free", label: "Free", input_per_1m: 0, output_per_1m: 0 },
  { provider: "opencode", model: "big-pickle-zen", label: "Big Pickle", input_per_1m: 0, output_per_1m: 0 },
];
const FAMILIES = [
  {
    id: "samagent-n4", label: "SamAgent N4", modes: [
      { mode: "flash", label: "Flash", agent: true, pool: POOL },
      { mode: "standard", label: "Standard", agent: true, pool: POOL },
    ],
  },
  {
    id: "code", label: "Code", modes: [
      { mode: "chat", label: "Chat", agent: false, pool: POOL },
    ],
  },
];
const clone = (o) => JSON.parse(JSON.stringify(o));

globalThis.fetch = async (url, init = {}) => {
  const u = String(url);
  const body = u.includes("/api/aliases")
    ? { families: clone(FAMILIES), defaults: clone(FAMILIES) }
    : {};
  return { ok: true, status: 200, text: async () => JSON.stringify(body) };
};

document.body.innerHTML = `<div id="selector-body"></div><div id="selector-agent-body"></div><span id="selector-save-state"></span><span id="selector-agent-save-state"></span>`;

const { loadSelectorPanel, loadAgentSelectorPanel } = await import("../model-selector.js");

await test("onglet agent : seules les familles/modes agent sont affichés", async () => {
  await loadAgentSelectorPanel();
  const body = document.getElementById("selector-agent-body");
  const cards = [...body.querySelectorAll(".ms-family-card")];
  assert.equal(cards.length, 1, "une seule carte famille (agent)");
  assert.equal(cards[0].dataset.fam, "samagent-n4");
  const modes = [...cards[0].querySelectorAll(".ms-mode-card .ms-mode-title")].map((e) => e.textContent);
  assert.deepEqual(modes, ["Flash", "Standard"], "seuls les modes agent");
  // Les réglages 1 modèle / Fallback sont présents comme dans l'onglet principal.
  assert.ok(cards[0].querySelector(".ms-seg-btn"), "boutons 1 modèle / Fallback présents");
});

await test("onglet principal : toutes les familles sont affichées", async () => {
  await loadSelectorPanel();
  const body = document.getElementById("selector-body");
  const cards = [...body.querySelectorAll(".ms-family-card")];
  assert.equal(cards.length, 2, "les deux familles");
  const fams = cards.map((c) => c.dataset.fam).sort();
  assert.deepEqual(fams, ["code", "samagent-n4"]);
});
