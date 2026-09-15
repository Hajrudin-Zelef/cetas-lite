import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";
// jsdom : résolution standard (node_modules du projet) puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

const dom = new JSDOM(
  `<!doctype html><html><body>
  <div id="features-rows"></div>
  <select id="search-mode"><option value="race">Rapide</option><option value="priority">Priorité</option></select>
  <select id="cfg-websearch-mode"><option value="auto">Auto</option><option value="natif">Natif</option><option value="outils">Outils</option><option value="off">Off</option></select>
  <div id="search-providers-tabs"></div>
  <div id="search-provider-content"></div>
  <select id="cfg-theme"><option value="clair">Clair</option><option value="sombre">Sombre</option><option value="hard_dark">Hard Dark</option></select>
  <div id="palette-swatches"></div>
  <div id="config-savebar" style="display:none"><span id="config-savebar-hint">Modifications non enregistrées</span></div>
</body></html>`,
  { pretendToBeVisual: true, url: "http://localhost/" }
);
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.localStorage = dom.window.localStorage;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;

// --- Mocks ---
const putBodies = [];
let searchMock = {
  mode: "race",
  providers: [
    { id: "brave", label: "Brave", enabled: true, configured: false, key_url: "https://brave.com/key", key_label: "Clé API Brave", keyless: false },
    { id: "tavily", label: "Tavily", enabled: false, configured: false, key_url: "https://tavily.com", key_label: "Clé API Tavily", keyless: false },
    { id: "duckduckgo", label: "DuckDuckGo", enabled: true, configured: false, keyless: true },
  ],
};
globalThis.fetch = async (url, opts = {}) => {
  const u = String(url);
  const method = opts.method || "GET";
  const json = (body) => ({ ok: true, status: 200, text: async () => JSON.stringify(body) });
  if (u === "/api/settings" && method === "GET")
    return json({ theme: "clair", palette: "bleu", tts: "system", transcription: "system", prompt_enhance: "none", summarizer: "none", title_gen: "conversation", error_analysis: "none", websearch_mode: "auto" });
  if (u === "/api/settings" && method === "PUT") {
    putBodies.push(JSON.parse(opts.body));
    return json({ ok: true });
  }
  if (u === "/api/providers")
    return json({ providers: [{ id: "openrouter", label: "OpenRouter", configured: true, local: false }, { id: "ollama", label: "Ollama", configured: false, local: true }] });
  if (u === "/api/search/settings" && method === "GET") return json(searchMock);
  if (u === "/api/search/settings" && method === "PUT") {
    const body = JSON.parse(opts.body);
    if (body.mode) searchMock.mode = body.mode;
    for (const [id, p] of Object.entries(body.providers || {})) {
      const cur = searchMock.providers.find((x) => x.id === id);
      if (!cur) continue;
      if (p.enabled !== undefined) cur.enabled = p.enabled;
      if (p.key !== undefined && p.key !== null) cur.configured = p.key !== "";
    }
    return json(searchMock);
  }
  return json({});
};

const panels = await import("../config-panels.js");
const { _test } = panels;
const savebar = () => document.getElementById("config-savebar");

await test("Fonctionnalités : six lignes rendues avec les bonnes options", async () => {
  await _test.loadFeaturesPanel();
  const rows = document.querySelectorAll("#features-rows .feat-row");
  assert.equal(rows.length, 6, "six lignes attendues");
  const names = [...rows].map((r) => r.querySelector(".feat-name").textContent);
  assert.deepEqual(names, ["Synthèse vocale", "Transcription", "Amélioration de prompts/rôles", "Résumé IA", "Génération du titre", "Analyse des erreurs"]);
  const ttsSel = document.querySelector('select[data-feature="tts"]');
  assert.ok(ttsSel, "select tts présent");
  assert.deepEqual([...ttsSel.options].map((o) => o.value), ["system", "none"]);
  assert.equal(ttsSel.value, "system");
  const tgSel = document.querySelector('select[data-feature="title_gen"]');
  assert.ok([...tgSel.options].some((o) => o.value === "conversation"), "option Modèle de la conversation");
  assert.equal(tgSel.value, "conversation");
  const peSel = document.querySelector('select[data-feature="prompt_enhance"]');
  assert.ok([...peSel.options].some((o) => o.value === "openrouter"), "providers proposés");
});

await test("Fonctionnalités : changement stagé (pas de PUT), Enregistrer persiste", async () => {
  putBodies.length = 0;
  const ttsSel = document.querySelector('select[data-feature="tts"]');
  ttsSel.value = "none";
  ttsSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 10));
  assert.equal(putBodies.length, 0, "aucun PUT immédiat");
  assert.ok(_test.isConfigDirty(), "onglet marqué sale");
  assert.notEqual(savebar().style.display, "none", "barre de sauvegarde visible");
  await _test.saveConfigPanel("models");
  assert.equal(putBodies.length, 1);
  assert.deepEqual(putBodies[0], { tts: "none" });
  assert.ok(!_test.isConfigDirty(), "plus rien à sauvegarder");
  assert.match(document.getElementById("config-savebar-hint").textContent, /Enregistré/, "confirmation affichée");
});

await test("Recherche Web : onglets moteurs + DuckDuckGo sans clé", async () => {
  await _test.loadSearchPanel();
  const tabs = [...document.querySelectorAll("#search-providers-tabs .provider-tab")].map((t) => t.textContent);
  assert.deepEqual(tabs, ["Brave", "Tavily", "DuckDuckGo"]);
  assert.equal(document.getElementById("search-mode").value, "race");
  const ddgSec = document.querySelector('#search-provider-content .provider-section[data-provider="duckduckgo"]');
  assert.ok(ddgSec, "section DuckDuckGo présente");
  assert.equal(ddgSec.querySelector(".apikey-input"), null, "aucun champ clé pour DuckDuckGo");
  assert.match(ddgSec.textContent, /Aucune clé requise/);
  const braveSec = document.querySelector('#search-provider-content .provider-section[data-provider="brave"]');
  const input = braveSec.querySelector('input[type="password"]');
  assert.ok(input, "champ clé Brave présent");
  const eye = braveSec.querySelector(".search-eye-btn");
  assert.ok(eye, "bouton œil présent");
  eye.click();
  assert.equal(input.type, "text", "œil affiche la clé");
  eye.click();
  assert.equal(input.type, "password", "œil masque à nouveau");
  const link = braveSec.querySelector(".apikey-get-link");
  assert.ok(link && link.href.includes("brave.com"), "lien d'obtention présent");
  assert.equal(braveSec.querySelector(".search-key-actions button.models-save-btn"), null, "plus de bouton par fournisseur");
});

await test("Recherche Web : toggle stagé puis persisté à l'enregistrement", async () => {
  const tavSec = document.querySelector('#search-provider-content .provider-section[data-provider="tavily"]');
  const cb = tavSec.querySelector('input[type="checkbox"]');
  assert.equal(cb.checked, false);
  cb.checked = true;
  cb.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 10));
  assert.equal(searchMock.providers.find((p) => p.id === "tavily").enabled, false, "pas de PUT immédiat");
  assert.match(tavSec.querySelector(".search-key-staged").textContent, /Enregistrer/);
  await _test.saveConfigPanel("search");
  assert.equal(searchMock.providers.find((p) => p.id === "tavily").enabled, true, "persisté après Enregistrer");
});

await test("Recherche Web : clé API stagée puis persistée", async () => {
  const braveSec = document.querySelector('#search-provider-content .provider-section[data-provider="brave"]');
  const input = braveSec.querySelector('input[type="password"]');
  input.value = "brave-secret-123";
  input.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 10));
  assert.equal(searchMock.providers.find((p) => p.id === "brave").configured, false, "clé pas encore envoyée");
  await _test.saveConfigPanel("search");
  assert.equal(searchMock.providers.find((p) => p.id === "brave").configured, true, "clé persistée");
});

await test("Recherche Web : mode moteur + comportement stagés", async () => {
  const modeSel = document.getElementById("search-mode");
  modeSel.value = "priority";
  modeSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  const wsmSel = document.getElementById("cfg-websearch-mode");
  wsmSel.value = "off";
  wsmSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 10));
  assert.equal(searchMock.mode, "race", "mode pas encore persisté");
  putBodies.length = 0;
  await _test.saveConfigPanel("search");
  assert.equal(searchMock.mode, "priority", "mode moteur persisté");
  assert.deepEqual(putBodies[0], { websearch_mode: "off" }, "comportement persisté via /api/settings");
});

await test("Sauvegarde générale : saveAllConfigPanels persiste tous les onglets sales", async () => {
  putBodies.length = 0;
  const trSel = document.querySelector('select[data-feature="transcription"]');
  trSel.value = "none";
  trSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  const braveSec = document.querySelector('#search-provider-content .provider-section[data-provider="brave"]');
  const cb = braveSec.querySelector('input[type="checkbox"]');
  cb.checked = false;
  cb.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 10));
  assert.ok(_test.isConfigDirty());
  await _test.saveAllConfigPanels();
  assert.ok(putBodies.some((b) => b.transcription === "none"), "fonctionnalité persistée");
  assert.equal(searchMock.providers.find((p) => p.id === "brave").enabled, false, "recherche persistée");
  assert.ok(!_test.isConfigDirty());
});

await test("Apparence : sept swatches, clic = aperçu sans PUT, Enregistrer persiste", async () => {
  await _test.loadAppearancePanel();
  const swatches = document.querySelectorAll("#palette-swatches .palette-swatch");
  assert.equal(swatches.length, 7, "sept palettes");
  const violet = [...swatches].find((s) => s.dataset.palette === "violet");
  assert.ok(violet.style.getPropertyValue("--sw"), "couleur du swatch définie");
  putBodies.length = 0;
  violet.click();
  await new Promise((r) => setTimeout(r, 10));
  assert.equal(document.documentElement.dataset.palette, "violet", "aperçu immédiat");
  assert.ok(violet.classList.contains("active"));
  assert.equal(putBodies.length, 0, "aucun PUT avant Enregistrer");
  await _test.saveConfigPanel("appearance");
  assert.deepEqual(putBodies[0], { palette: "violet" });
});

await test("Apparence : thème stagé avec aperçu, revertUnsavedConfig replie", async () => {
  const themeSel = document.getElementById("cfg-theme");
  themeSel.value = "sombre";
  themeSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 10));
  assert.ok(document.body.classList.contains("dark"), "aperçu sombre appliqué");
  putBodies.length = 0;
  const swatches = document.querySelectorAll("#palette-swatches .palette-swatch");
  const rouge = [...swatches].find((s) => s.dataset.palette === "rouge");
  rouge.click();
  await new Promise((r) => setTimeout(r, 10));
  assert.equal(document.documentElement.dataset.palette, "rouge");
  _test.revertUnsavedConfig();
  assert.equal(document.documentElement.dataset.palette, "violet", "palette repliée sur la sauvegardée");
  assert.ok(!document.body.classList.contains("dark"), "thème replié");
  assert.equal(putBodies.length, 0, "rien persisté");
  assert.ok(!_test.isConfigDirty());
  // Enregistrer le thème sombre pour de bon, puis vérifier la persistance.
  themeSel.value = "sombre";
  themeSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await _test.saveConfigPanel("appearance");
  assert.deepEqual(putBodies[0], { theme: "sombre" });
});

await test("helpers : searchStatusText", () => {
  assert.match(_test.searchStatusText({ enabled: false }), /Désactivé/);
  assert.match(_test.searchStatusText({ enabled: true, keyless: true }), /sans clé/);
  assert.match(_test.searchStatusText({ enabled: true, configured: true }), /configurée/);
  assert.match(_test.searchStatusText({ enabled: true, configured: false }), /ignoré/);
});
