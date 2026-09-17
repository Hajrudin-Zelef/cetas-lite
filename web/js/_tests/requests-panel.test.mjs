// Panneau "Requêtes" façon DeepSeek + flèche retour-en-bas.
// (Ref : capture DeepSeek fournie par l'utilisateur.)
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

// Mock fetch avec text() (api() lit resp.text()).
const calls = [];
globalThis.fetch = async (url, opts = {}) => {
  calls.push({ url: String(url), method: opts.method || "GET", body: opts.body });
  if (String(url).startsWith("/api/conversations/restore")) {
    return { ok: true, status: 200, text: async () => "{}" };
  }
  if (String(url).startsWith("/api/conversations")) {
    return {
      ok: true,
      status: 200,
      text: async () =>
        JSON.stringify({
          archives: [
            { id: "c1", title: "comment appelle ton une fonction qui…" },
            { id: "c2", title: "ok corrigeons 2 points, je veux…" },
          ],
        }),
    };
  }
  if (String(url).startsWith("/api/chat/state")) {
    return { ok: true, status: 200, text: async () => JSON.stringify({ id: "c2", turns: 0 }) };
  }
  return { ok: false, status: 404, text: async () => JSON.stringify({ error: "nope" }) };
};

const { initRequestsPanel } = await import("../requests-panel.js");

function setupDom() {
  document.body.innerHTML = `
    <main class="main">
      <div id="chat-container"></div>
      <div class="input-area"></div>
    </main>`;
}

const flush = () => new Promise((r) => setTimeout(r, 20));

test("panneau : construit avec les conversations, active surlignée", async () => {
  setupDom();
  calls.length = 0;
  const ctl = initRequestsPanel();
  assert.ok(ctl, "contrôleur retourné");
  await flush();
  const panel = document.querySelector("main.main > .requests-panel");
  assert.ok(panel, "panneau dans main.main");
  assert.equal(panel.hidden, false, "visible quand il y a des archives");
  const items = panel.querySelectorAll(".request-item");
  assert.equal(items.length, 2);
  assert.equal(items[0].textContent, "comment appelle ton une fonction qui…");
  assert.equal(items[0].title, "comment appelle ton une fonction qui…");
  assert.equal(items[0].dataset.id, "c1");
  assert.ok(!items[0].classList.contains("active"), "c1 non active");
  assert.ok(items[1].classList.contains("active"), "c2 active (id courant)");
});

test("panneau : clic restaure la conversation et notifie", async () => {
  setupDom();
  calls.length = 0;
  let resetFired = false;
  window.addEventListener("cetas:chat-reset", () => (resetFired = true), { once: true });
  initRequestsPanel();
  await flush();
  const items = document.querySelectorAll(".request-item");
  items[0].click();
  await flush();
  const restore = calls.find((c) => c.url.startsWith("/api/conversations/restore"));
  assert.ok(restore, "POST restore appelé");
  assert.equal(restore.method, "POST");
  assert.match(restore.body, /"c1"/);
  assert.ok(resetFired, "événement cetas:chat-reset émis");
});

test("panneau : masqué quand aucune archive", async () => {
  setupDom();
  const origFetch = globalThis.fetch;
  globalThis.fetch = async (url, opts = {}) => {
    if (String(url).startsWith("/api/conversations") && !String(url).includes("restore")) {
      return { ok: true, status: 200, text: async () => JSON.stringify({ archives: [] }) };
    }
    return origFetch(url, opts);
  };
  try {
    initRequestsPanel();
    await flush();
    const panel = document.querySelector("main.main > .requests-panel");
    assert.ok(panel, "panneau créé");
    assert.equal(panel.hidden, true, "masqué sans archives");
  } finally {
    globalThis.fetch = origFetch;
  }
});

test("panneau : pas de doublon à l'init répétée", async () => {
  setupDom();
  initRequestsPanel();
  initRequestsPanel();
  await flush();
  assert.equal(document.querySelectorAll("main.main > .requests-panel").length, 1);
});

test("css : flèche pastille blanche en mode clair + centrée au-dessus du composer", () => {
  const css = read("web/css/cetas-lite.css");
  // Fond blanc en mode clair (fini le tout-noir).
  assert.match(css, /\.thread-to-bottom\s*\{[^}]*background:\s*#ffffff/);
  // Ancrage au-dessus du composer, centré (vue principale).
  assert.match(css, /\.input-area > \.thread-to-bottom\s*\{[^}]*left:\s*50%/);
  assert.match(css, /\.input-area > \.thread-to-bottom\s*\{[^}]*bottom:\s*calc\(100% \+ 10px\)/);
});

test("css : panneau requêtes scrollable, responsive", () => {
  const css = read("web/css/cetas-lite.css");
  assert.match(css, /\.requests-panel\s*\{[^}]*overflow-y:\s*auto/);
  assert.match(css, /\.requests-panel\s*\{[^}]*max-height:\s*46vh/);
  assert.match(css, /\.request-item\s*\{[^}]*text-overflow:\s*ellipsis/);
  assert.match(css, /@media\s*\(max-width:\s*1100px\)/);
});

test("chat : initRequestsPanel câblé + flèche ancrée au composer", () => {
  const src = read("web/js/chat.js");
  assert.match(src, /import\s*\{\s*initRequestsPanel\s*\}\s*from\s*"\.\/requests-panel\.js"/);
  assert.match(src, /initRequestsPanel\(\)/);
  assert.match(src, /inputArea\.appendChild\(view\.toBottomBtn\)/);
});
