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
const now = Date.now();
// Mock fetch avec text() (api() lit resp.text()).
globalThis.fetch = async (url) => {
  if (String(url).startsWith("/api/conversations")) {
    return {
      ok: true,
      status: 200,
      text: async () =>
        JSON.stringify({
          archives: [
            { id: "a1", title: "récente", updated: now - 3600000, messages: 4 },
            { id: "a2", title: "il y a 3 jours", updated: now - 3 * DAY, messages: 2 },
            { id: "a3", title: "il y a 10 jours", updated: now - 10 * DAY, messages: 7 },
            { id: "a4", title: "il y a 40 jours", updated: now - 40 * DAY, messages: 1 },
            { id: "a5", title: "sans date", updated: 0, messages: 1 },
          ],
        }),
    };
  }
  return { ok: false, status: 404, text: async () => JSON.stringify({ error: "nope" }) };
};

const { initSidebar } = await import("../sidebar.js");

function setupDom() {
  document.body.innerHTML = `
    <button id="new-chat-btn"></button>
    <div class="fav-section" id="fav-section" style="display:none"><div id="fav-list"></div></div>
    <div class="cat-select-row"><div id="cat-select"><span class="cat-select-label"></span></div></div>
    <div id="cat-select-dropdown" style="display:none"></div>
    <input type="text" id="conv-search">
    <div class="conv-list-container"><div id="conv-list" class="conv-list"></div></div>`;
}

const flush = () => new Promise((r) => setTimeout(r, 30));

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
