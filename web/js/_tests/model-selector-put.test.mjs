import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

// Régression : PUT /api/aliases répondait 400 "corps JSON invalide" car
// persist() pré-encodait le corps en JSON alors que api() l'encode déjà
// (double encodage : le serveur recevait une chaîne, pas un objet).

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
  { provider: "openrouter", model: "openrouter/free", label: "Free Models Router", input_per_1m: 0, output_per_1m: 0 },
  { provider: "opencode", model: "big-pickle-zen", label: "Big Pickle", input_per_1m: 0, output_per_1m: 0 },
];
const FAMILIES = [
  { id: "samagent-nano", label: "SamAgent Nano", modes: [{ mode: "free", label: "Free", pool: POOL }] },
];
const clone = (o) => JSON.parse(JSON.stringify(o));

let putBody = null;
globalThis.fetch = async (url, init = {}) => {
  const u = String(url);
  if (u.includes("/api/aliases") && (init.method || "GET") === "PUT") {
    putBody = init.body; // tel que api() l'a transmis à fetch
    return { ok: true, status: 200, text: async () => JSON.stringify({ ok: true, families: FAMILIES }) };
  }
  const body = u.includes("/api/aliases")
    ? { families: clone(FAMILIES), defaults: clone(FAMILIES) }
    : {};
  return { ok: true, status: 200, text: async () => JSON.stringify(body) };
};

document.body.innerHTML = `<div id="selector-body"></div>`;

const { loadSelectorPanel } = await import("../model-selector.js");

await test("PUT /api/aliases : le corps est un objet JSON, pas une chaîne doublement encodée", async () => {
  await loadSelectorPanel();
  // Bascule "1 modèle" sur le mode Free (2 modèles -> 1).
  const btnSingle = document.querySelector(".ms-seg-btn");
  assert.ok(btnSingle, "bouton 1 modèle présent");
  btnSingle.click();
  // persist() est débouncé (450 ms).
  await new Promise((r) => setTimeout(r, 800));
  assert.ok(putBody !== null, "le PUT a été émis");
  assert.equal(typeof putBody, "string", "fetch reçoit une chaîne (encodée par api())");
  const decoded = JSON.parse(putBody);
  assert.equal(typeof decoded, "object", "le serveur reçoit un objet, pas une chaîne");
  assert.deepEqual(decoded, {
    "samagent-nano": { free: [{ provider: "openrouter", model: "openrouter/free" }] },
  });
});

await test("onglet Fallback : re-bascule depuis « 1 modèle » (tout l'ensemble)", async () => {
  // État : le test précédent a basculé le mode en « 1 modèle » (1 radio).
  // Avant le correctif, cliquer « Fallback » ne faisait rien : switchKind
  // conservait le seul modèle coché, le pool restait à 1, kind = "single".
  const btnFallback = [...document.querySelectorAll(".ms-seg-btn")].find(
    (b) => b.textContent === "Fallback"
  );
  assert.ok(btnFallback, "bouton Fallback présent");
  btnFallback.click();
  await new Promise((r) => setTimeout(r, 800)); // persist() débouncé (450 ms)
  const inputs = [...document.querySelectorAll("#selector-body input")];
  assert.ok(inputs.length >= 2, "liste des modèles affichée");
  assert.equal(inputs[0].type, "checkbox", "retour en mode Fallback (cases à cocher)");
  const decoded = JSON.parse(putBody);
  assert.equal(
    decoded["samagent-nano"].free.length,
    2,
    "l'override couvre tout l'ensemble, pas le seul modèle coché"
  );
});
