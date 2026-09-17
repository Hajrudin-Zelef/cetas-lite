// Panneau "Requêtes" façon DeepSeek : sommaire des requêtes de la
// conversation en cours, clic = scroll vers le message.
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
globalThis.requestAnimationFrame =
  dom.window.requestAnimationFrame?.bind(dom.window) || ((fn) => setTimeout(fn, 16));
globalThis.MutationObserver = dom.window.MutationObserver;

// scrollIntoView : jsdom ne l'implémente pas, on l'enregistre.
const scrolledTo = [];
dom.window.Element.prototype.scrollIntoView = function (opts) {
  scrolledTo.push({ el: this, opts });
};

const { initRequestsPanel } = await import("../requests-panel.js");

function setupDom() {
  scrolledTo.length = 0;
  document.body.innerHTML = `
    <main class="main">
      <div id="chat-container">
        <div class="message-wrapper message-wrapper-user">
          <div class="message message-user"><div class="message-text">première requête
deuxième ligne</div></div>
        </div>
        <div class="message-wrapper"><div class="message message-assistant"><div class="message-text">réponse</div></div></div>
        <div class="message-wrapper message-wrapper-user">
          <div class="message message-user"><div class="message-text">seconde requête</div></div>
        </div>
      </div>
      <div class="input-area"></div>
    </main>`;
}

function addUserMessage(text) {
  const log = document.getElementById("chat-container");
  const wrapper = document.createElement("div");
  wrapper.className = "message-wrapper message-wrapper-user";
  const bubble = document.createElement("div");
  bubble.className = "message message-user";
  const t = document.createElement("div");
  t.className = "message-text";
  t.textContent = text;
  bubble.appendChild(t);
  wrapper.appendChild(bubble);
  log.appendChild(wrapper);
  return bubble;
}

const flush = () => new Promise((r) => setTimeout(r, 30));

test("panneau : sommaire des requêtes de la conversation en cours", async () => {
  setupDom();
  const ctl = initRequestsPanel();
  assert.ok(ctl, "contrôleur retourné");
  const panel = document.querySelector("main.main > .requests-panel");
  assert.ok(panel, "panneau dans main.main");
  assert.equal(panel.hidden, false);
  const items = panel.querySelectorAll(".request-item");
  assert.equal(items.length, 2, "2 requêtes utilisateur, pas les réponses");
  assert.equal(items[0].textContent, "première requête", "première ligne uniquement");
  assert.equal(items[0].title, "première requête");
  assert.equal(items[1].textContent, "seconde requête");
});

test("panneau : clic scrolle vers le message", async () => {
  setupDom();
  initRequestsPanel();
  const items = document.querySelectorAll(".request-item");
  items[1].click();
  assert.equal(scrolledTo.length, 1);
  assert.ok(scrolledTo[0].el.classList.contains("message-user"));
  assert.equal(
    scrolledTo[0].el.querySelector(".message-text").textContent,
    "seconde requête",
    "c'est le bon message qui est ciblé"
  );
  assert.equal(scrolledTo[0].opts.behavior, "smooth");
});

test("panneau : requête la plus proche du haut surlignée", async () => {
  setupDom();
  initRequestsPanel();
  const items = document.querySelectorAll(".request-item");
  // jsdom : tous les rects à 0 → la première gagne (déterministe).
  assert.ok(items[0].classList.contains("active"));
  assert.ok(!items[1].classList.contains("active"));
});

test("panneau : se reconstruit à l'arrivée d'un message", async () => {
  setupDom();
  initRequestsPanel();
  addUserMessage("troisième requête");
  await flush(); // MutationObserver
  const items = document.querySelectorAll(".request-item");
  assert.equal(items.length, 3);
  assert.equal(items[2].textContent, "troisième requête");
});

test("panneau : masqué sans requête", async () => {
  scrolledTo.length = 0;
  document.body.innerHTML = `
    <main class="main"><div id="chat-container"></div><div class="input-area"></div></main>`;
  initRequestsPanel();
  const panel = document.querySelector("main.main > .requests-panel");
  assert.ok(panel);
  assert.equal(panel.hidden, true);
});

test("panneau : pas de doublon à l'init répétée", async () => {
  setupDom();
  initRequestsPanel();
  initRequestsPanel();
  assert.equal(document.querySelectorAll("main.main > .requests-panel").length, 1);
});

test("panneau : reset du chat reconstruit le sommaire", async () => {
  setupDom();
  initRequestsPanel();
  document.getElementById("chat-container").innerHTML = "";
  window.dispatchEvent(new CustomEvent("cetas:chat-reset"));
  await flush();
  const panel = document.querySelector("main.main > .requests-panel");
  assert.equal(panel.hidden, true, "plus de requêtes après reset");
  assert.equal(panel.querySelectorAll(".request-item").length, 0);
});

test("css : flèche pastille blanche en mode clair + centrée au-dessus du composer", () => {
  const css = read("web/css/cetas-lite.css");
  assert.match(css, /\.thread-to-bottom\s*\{[^}]*background:\s*#ffffff/);
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
