// Sidebar : historique groupé par date façon DeepSeek
// (Aujourd'hui / 7 derniers jours / 30 derniers jours / Plus anciens),
// lignes épurées (titre seul).
import { test } from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
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

const repoRoot = new URL("../../../", import.meta.url);
const read = (rel) => fs.readFileSync(new URL(rel, repoRoot), "utf8");

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;

const DAY = 86400000;
// Utiliser un timestamp à midi UTC pour éviter les problèmes de timezone
// (setHours(0,0,0,0) utilise l'heure locale).
const now = new Date().setHours(12, 0, 0, 0);
Date.now = () => now;
// Mock fetch avec text() (api() lit resp.text()).
globalThis.fetch = async (url) => {
  if (String(url).startsWith("/api/sessions/current")) {
    return { ok: true, status: 200, text: async () => JSON.stringify({ id: "a1" }) };
  }
  if (String(url).startsWith("/api/sessions")) {
    return {
      ok: true,
      status: 200,
      text: async () =>
        JSON.stringify({
          sessions: [
            { id: "a1", title: "récente", createdAt: now - 3600000, updatedAt: now - 3600000, messages: 4 },
            { id: "a2", title: "il y a 3 jours", createdAt: now - 3 * DAY, updatedAt: now - 3 * DAY, messages: 2 },
            { id: "a3", title: "il y a 10 jours", createdAt: now - 10 * DAY, updatedAt: now - 10 * DAY, messages: 7 },
            { id: "a4", title: "il y a 40 jours", createdAt: now - 40 * DAY, updatedAt: now - 40 * DAY, messages: 1 },
            { id: "a5", title: "sans date", createdAt: 0, updatedAt: 0, messages: 1 },
          ],
        }),
    };
  }
  return { ok: false, status: 404, text: async () => JSON.stringify({ error: "nope" }) };
};

const { initSidebar } = await import("../sidebar.js");

function setupDom() {
  localStorage.clear();
  document.body.innerHTML = `
    <button id="new-chat-btn"></button>
    <div class="fav-section" id="fav-section" style="display:none"><div id="fav-list"></div></div>
    <div class="cat-select-row"><div id="cat-select"><span class="cat-select-label"></span></div></div>
    <div id="cat-select-dropdown" style="display:none"></div>
    <input type="text" id="conv-search">
    <div class="conv-list-container"><div id="conv-list" class="conv-list"></div></div>`;
}

const flush = () => new Promise((r) => setTimeout(r, 200));

function itemsOf(headerText) {
  const headers = [...document.querySelectorAll(".conv-section-header")];
  const h = headers.find((x) => x.textContent === headerText);
  assert.ok(h, "section « " + headerText + " » présente");
  const items = [];
  let n = h.nextElementSibling;
  while (n && !n.classList.contains("conv-section-header")) {
    if (n.classList.contains("conv-item")) items.push(n);
    n = n.nextElementSibling;
  }
  return items.map((i) => i.querySelector(".conv-item-title").textContent);
}

test("historique : sections par date dans l'ordre", async () => {
  setupDom();
  initSidebar();
  await flush();
  const headers = [...document.querySelectorAll(".conv-section-header")].map((h) => h.textContent);
  assert.deepEqual(headers, ["Aujourd'hui", "7 derniers jours", "30 derniers jours", "Plus anciens"]);
});

test("historique : conversations rangées dans la bonne section", async () => {
  setupDom();
  initSidebar();
  await flush();
  assert.deepEqual(itemsOf("Aujourd'hui"), ["récente"]);
  assert.deepEqual(itemsOf("7 derniers jours"), ["il y a 3 jours"]);
  assert.deepEqual(itemsOf("30 derniers jours"), ["il y a 10 jours"]);
  assert.deepEqual(itemsOf("Plus anciens"), ["il y a 40 jours", "sans date"]);
});

test("historique : lignes épurées façon DeepSeek (titre seul)", async () => {
  setupDom();
  initSidebar();
  await flush();
  assert.equal(document.querySelectorAll(".conv-item-date-line").length, 0, "plus de sous-titre date/messages");
  const titles = [...document.querySelectorAll(".conv-item-title")];
  assert.equal(titles.length, 5);
});

test("css : en-têtes de section stylés", () => {
  const css = read("web/css/cetas-lite.css");
  assert.match(css, /\.conv-section-header\s*\{/);
});

test("session courante surlignée", async () => {
  setupDom();
  initSidebar();
  await flush();
  const active = document.querySelector(".conv-item.active .conv-item-title");
  assert.ok(active, "une session doit porter la classe active");
  assert.equal(active.textContent, "récente");
});

test("ouvrir une session supprimée ailleurs : retirée du cache, pas de fantôme", async () => {
  const origFetch = globalThis.fetch;
  globalThis.fetch = async (url) => {
    if (String(url).includes("/open")) {
      return { ok: false, status: 404, text: async () => JSON.stringify({ error: "session introuvable" }) };
    }
    return origFetch(url);
  };
  try {
    setupDom();
    initSidebar();
    await flush();
    assert.equal(document.querySelectorAll(".conv-item").length, 5);
    const rows = [...document.querySelectorAll(".conv-item")];
    const target = rows.find((r) => r.querySelector(".conv-item-title").textContent === "il y a 3 jours");
    target.querySelector(".conv-item-content").click();
    await flush();
    const titles = [...document.querySelectorAll(".conv-item-title")].map((t) => t.textContent);
    assert.ok(!titles.includes("il y a 3 jours"), "la ligne fantôme est retirée");
  } finally {
    globalThis.fetch = origFetch;
  }
});

test("clic sur une session : ouvre via POST /api/sessions/{id}/open", async () => {
  const calls = [];
  const origFetch = globalThis.fetch;
  globalThis.fetch = async (url, opts) => {
    calls.push(String(url) + " " + ((opts && opts.method) || "GET"));
    return origFetch(url, opts);
  };
  try {
    setupDom();
    initSidebar();
    await flush();
    const rows = [...document.querySelectorAll(".conv-item")];
    const target = rows.find((r) => r.querySelector(".conv-item-title").textContent === "il y a 3 jours");
    assert.ok(target, "ligne a2 présente");
    target.querySelector(".conv-item-content").click();
    await flush();
    assert.ok(
      calls.some((c) => c === "/api/sessions/a2/open POST"),
      "open appelé, got: " + JSON.stringify(calls)
    );
    // render() reconstruit le DOM : on re-requête la ligne.
    const active = document.querySelector(".conv-item.active .conv-item-title");
    assert.ok(active, "une session doit être active après ouverture");
    assert.equal(active.textContent, "il y a 3 jours");
  } finally {
    globalThis.fetch = origFetch;
  }
});
