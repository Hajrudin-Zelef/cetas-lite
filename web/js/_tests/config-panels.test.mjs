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
  <div id="search-providers-tabs"></div>
  <div id="search-provider-content"></div>
  <input id="sessions-filter" type="search">
  <div id="sessions-chat"></div>
  <div id="sessions-agents"></div>
  <div id="palette-swatches"></div>
</body></html>`,
  { pretendToBeVisual: true, url: "http://localhost/" }
);
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.localStorage = dom.window.localStorage;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;
// confirmDialog() replie sur window.confirm hors modale HTML.
let confirmMsg = null;
globalThis.window.confirm = (m) => { confirmMsg = m; return confirmAnswer; };
let confirmAnswer = true;
let chatTurns = 0;
const restoreCalls = [];

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
const sessionsMock = {
  chat: [
    { id: "c1", kind: "chat", title: "Recette carbonara", updated: 1757932800000, messages: 12 },
    { id: "c2", kind: "chat", title: "Debug panique Go", updated: 1757846400000, messages: 4 },
  ],
  agents: [{ id: "a1", kind: "agent", title: "Refactor auth", updated: 1757932800, status: "running", family: "code", mode: "build" }],
};
globalThis.fetch = async (url, opts = {}) => {
  const u = String(url);
  const method = opts.method || "GET";
  const json = (body) => ({ ok: true, status: 200, text: async () => JSON.stringify(body) });
  if (u === "/api/settings" && method === "GET")
    return json({ theme: "clair", palette: "bleu", tts: "system", transcription: "system", prompt_enhance: "none", summarizer: "none", title_gen: "conversation", error_analysis: "none" });
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
  if (u === "/api/sessions") return json(sessionsMock);
  if (u === "/api/chat/state") return json({ turns: chatTurns });
  if (u === "/api/conversations/restore") { restoreCalls.push(JSON.parse(opts.body)); return json({ ok: true }); }
  if (u.startsWith("/api/conversations/") && method === "DELETE") {
    sessionsMock.chat = sessionsMock.chat.filter((s) => !u.endsWith(s.id));
    return json({ ok: true });
  }
  if (u.startsWith("/api/agents/") && method === "DELETE") {
    sessionsMock.agents = sessionsMock.agents.filter((s) => !u.endsWith(s.id));
    return json({ ok: true });
  }
  return json({});
};

const panels = await import("../config-panels.js");
const { _test } = panels;

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

await test("Fonctionnalités : changement persisté via PUT /api/settings", async () => {
  putBodies.length = 0;
  const ttsSel = document.querySelector('select[data-feature="tts"]');
  ttsSel.value = "none";
  ttsSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 30));
  assert.equal(putBodies.length, 1);
  assert.deepEqual(putBodies[0], { tts: "none" });
});

await test("Recherche Web : onglets moteurs + mode + DuckDuckGo sans clé", async () => {
  await _test.loadSearchPanel();
  const tabs = [...document.querySelectorAll("#search-providers-tabs .provider-tab")].map((t) => t.textContent);
  assert.deepEqual(tabs, ["Brave", "Tavily", "DuckDuckGo"]);
  assert.equal(document.getElementById("search-mode").value, "race");
  // Panneau DuckDuckGo : pas de champ clé.
  const ddgSec = document.querySelector('#search-provider-content .provider-section[data-provider="duckduckgo"]');
  assert.ok(ddgSec, "section DuckDuckGo présente");
  assert.equal(ddgSec.querySelector(".apikey-input"), null, "aucun champ clé pour DuckDuckGo");
  assert.match(ddgSec.textContent, /Aucune clé requise/);
  // Panneau Brave : champ clé + œil + lien.
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
});

await test("Recherche Web : toggle d'activation persisté", async () => {
  const tavSec = document.querySelector('#search-provider-content .provider-section[data-provider="tavily"]');
  const cb = tavSec.querySelector('input[type="checkbox"]');
  assert.equal(cb.checked, false);
  cb.checked = true;
  cb.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 30));
  assert.equal(searchMock.providers.find((p) => p.id === "tavily").enabled, true);
});

await test("Recherche Web : changement de mode persisté", async () => {
  const modeSel = document.getElementById("search-mode");
  modeSel.value = "priority";
  modeSel.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 30));
  assert.equal(searchMock.mode, "priority");
});

await test("Sessions : listes CETAS + Agents rendues", async () => {
  await _test.loadSessionsPanel();
  const chatRows = document.querySelectorAll("#sessions-chat .sess-row");
  assert.equal(chatRows.length, 2, "deux conversations");
  assert.match(chatRows[0].querySelector(".sess-title").textContent, /carbonara/);
  assert.match(chatRows[0].querySelector(".sess-meta").textContent, /12 messages/);
  const agentRows = document.querySelectorAll("#sessions-agents .sess-row");
  assert.equal(agentRows.length, 1, "un agent");
  assert.match(agentRows[0].querySelector(".sess-meta").textContent, /En cours/);
});

await test("Sessions : filtre + suppression", async () => {
  const filter = document.getElementById("sessions-filter");
  filter.value = "go";
  filter.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  assert.equal(document.querySelectorAll("#sessions-chat .sess-row").length, 1, "filtre appliqué");
  filter.value = "";
  filter.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  const delBtn = [...document.querySelectorAll("#sessions-chat .sess-btn.danger")].find((b) => b.textContent === "Supprimer");
  delBtn.click();
  await new Promise((r) => setTimeout(r, 30));
  assert.equal(document.querySelectorAll("#sessions-chat .sess-row").length, 1, "conversation supprimée");
});

await test("Apparence : sept swatches, clic applique la palette", async () => {
  await _test.loadAppearancePanel();
  const swatches = document.querySelectorAll("#palette-swatches .palette-swatch");
  assert.equal(swatches.length, 7, "sept palettes");
  const violet = [...swatches].find((s) => s.dataset.palette === "violet");
  putBodies.length = 0;
  violet.click();
  await new Promise((r) => setTimeout(r, 30));
  assert.equal(document.documentElement.dataset.palette, "violet");
  assert.ok(violet.classList.contains("active"));
  assert.deepEqual(putBodies[0], { palette: "violet" });
});

await test("Sessions : restauration demande confirmation si la conversation courante n'est pas vide", async () => {
  // Recharge le panneau (sessionsMock.chat a été réduit par le test précédent).
  sessionsMock.chat = [
    { id: "c1", kind: "chat", title: "Recette carbonara", updated: 1757932800000, messages: 12 },
  ];
  const { refreshSessionsPanel } = _test;
  await refreshSessionsPanel();
  const openBtn = [...document.querySelectorAll("#sessions-chat .sess-btn")].find((b) => b.textContent === "Ouvrir");
  assert.ok(openBtn, "bouton Ouvrir présent");

  // Cas 1 : conversation courante vide -> pas de confirmation, restauration directe.
  chatTurns = 0; confirmAnswer = true; confirmMsg = null; restoreCalls.length = 0;
  openBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  assert.equal(confirmMsg, null, "aucune confirmation si rien à archiver");
  assert.equal(restoreCalls.length, 1, "restauration appelée");

  // Cas 2 : conversation courante non vide -> confirmation, refus = pas de restauration.
  chatTurns = 3; confirmAnswer = false; confirmMsg = null; restoreCalls.length = 0;
  openBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  assert.match(confirmMsg || "", /archivée/, "message d'archivage affiché");
  assert.equal(restoreCalls.length, 0, "restauration annulée");

  // Cas 3 : confirmation acceptée -> restauration.
  confirmAnswer = true; restoreCalls.length = 0;
  openBtn.click();
  await new Promise((r) => setTimeout(r, 50));
  assert.equal(restoreCalls.length, 1, "restauration appelée après confirmation");
  assert.equal(restoreCalls[0].id, "c1");
});

await test("helpers : agentStatusLabel / fmtSessionDate / searchStatusText", () => {
  assert.equal(_test.agentStatusLabel("running"), "En cours");
  assert.equal(_test.agentStatusLabel("done"), "Terminé");
  assert.equal(_test.agentStatusLabel("stopped"), "Arrêté");
  assert.match(_test.fmtSessionDate(1757932800000), /\d/);
  assert.equal(_test.fmtSessionDate(0), "");
  assert.match(_test.searchStatusText({ enabled: false }), /Désactivé/);
  assert.match(_test.searchStatusText({ enabled: true, keyless: true }), /sans clé/);
  assert.match(_test.searchStatusText({ enabled: true, configured: true }), /configurée/);
  assert.match(_test.searchStatusText({ enabled: true, configured: false }), /ignoré/);
});
