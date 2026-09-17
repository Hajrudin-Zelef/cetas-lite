// Hero façon Gemini (chat vide) : le vrai composer remonte au centre quand
// le chat est vide et redescend en bas dès le premier message.
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
globalThis.MutationObserver = dom.window.MutationObserver;
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;

const heroMod = await import("../chat-hero.js");
const { initChatHero, updateHeroMode, isHeroMode } = heroMod;

const wait = (ms) => new Promise((r) => setTimeout(r, ms));

const EMPTY_HTML = `<div id="empty-chat-placeholder" class="empty-chat-placeholder" data-empty>
  <div class="chat-hero">
    <div class="hero-glow" aria-hidden="true"></div>
    <h1 class="hero-title">Par où commencer ?</h1>
    <div id="hero-composer-slot" class="hero-composer-slot"></div>
  </div>
  <div id="empty-chat-category" class="empty-chat-category" style="display:none"></div>
</div>`;

// Reconstruit un DOM minimal : main > chat-container + model-alert + input-area.
function setupDom() {
  document.body.innerHTML = "";
  document.body.className = "";
  const main = document.createElement("main");
  main.className = "main";
  const log = document.createElement("div");
  log.id = "chat-container";
  log.className = "chat-container";
  log.innerHTML = EMPTY_HTML;
  const alert = document.createElement("div");
  alert.id = "model-alert";
  alert.className = "model-alert";
  alert.style.display = "none";
  const inputArea = document.createElement("div");
  inputArea.className = "input-area";
  const ta = document.createElement("textarea");
  ta.id = "prompt-input";
  ta.setAttribute("placeholder", "Écrivez votre message...");
  inputArea.appendChild(ta);
  // Comme dans le vrai DOM, l'indicateur est DANS .input-area.
  const hint = document.createElement("span");
  hint.id = "input-hint";
  hint.className = "input-hint";
  hint.textContent = "Xiaomi · Flash";
  inputArea.appendChild(hint);
  main.appendChild(log);
  main.appendChild(alert);
  main.appendChild(inputArea);
  document.body.appendChild(main);
  initChatHero();
  return { log, inputArea, alert, hint, ta };
}

describe("chat hero façon Gemini", () => {
  test("chat vide : le composer remonte dans le hero, titre + halo présents", () => {
    const { inputArea, alert, hint, ta } = setupDom();
    try {
      assert.equal(isHeroMode(), true, "mode hero actif");
      const slot = document.getElementById("hero-composer-slot");
      assert.equal(slot.contains(inputArea), true, "input-area déplacée dans le slot");
      assert.equal(slot.contains(alert), true, "model-alert déplacée dans le slot");
      assert.equal(slot.contains(hint), true, "input-hint suit dans le composer");
      assert.deepEqual(
        [slot.children[0], slot.children[1]],
        [alert, inputArea],
        "ordre : alerte puis composer"
      );
      assert.ok(document.body.classList.contains("hero-mode"), "classe hero-mode sur body");
      assert.equal(document.querySelector(".hero-title").textContent, "Par où commencer ?");
      assert.ok(document.querySelector(".hero-glow"), "halo présent");
      assert.equal(ta.getAttribute("placeholder"), "Demander à Cetas");
    } finally {
      document.body.innerHTML = "";
      document.body.className = "";
    }
  });

  test("premier message : le composer et l'alerte redescendent en bas", async () => {
    const { log, inputArea, alert, hint, ta } = setupDom();
    try {
      assert.equal(isHeroMode(), true);
      // clearEmpty() du ThreadView : le placeholder est retiré.
      const ph = document.getElementById("empty-chat-placeholder");
      ph.remove();
      const msg = document.createElement("div");
      msg.className = "message-wrapper message-wrapper-user";
      log.appendChild(msg);
      await wait(20); // MutationObserver = microtask
      assert.equal(isHeroMode(), false, "mode hero désactivé");
      const main = inputArea.parentNode;
      assert.equal(main.tagName, "MAIN", "input-area de retour dans main");
      assert.equal(alert.parentNode, main, "model-alert de retour dans main");
      assert.equal(hint.parentNode, inputArea, "input-hint resté dans le composer");
      assert.deepEqual(
        [alert.nextSibling, inputArea.nextSibling],
        [inputArea, null],
        "ordre d'origine conservé"
      );
      assert.equal(
        document.body.classList.contains("hero-mode"),
        false,
        "classe hero-mode retirée"
      );
      assert.equal(ta.getAttribute("placeholder"), "Écrivez votre message...", "placeholder restauré");
    } finally {
      document.body.innerHTML = "";
      document.body.className = "";
    }
  });

  test("reset avec le composer dans le hero : le nœud survit au innerHTML", async () => {
    const { log, inputArea } = setupDom();
    try {
      assert.equal(isHeroMode(), true, "précondition : hero actif");
      assert.ok(document.getElementById("hero-composer-slot").contains(inputArea));
      // ThreadView.reset() : innerHTML écrase le hero ALORS QUE le composer
      // est dedans — le nœud doit survivre via la référence gardée.
      log.innerHTML = EMPTY_HTML;
      await wait(20);
      assert.equal(isHeroMode(), true, "toujours en hero après reset");
      const slot = document.getElementById("hero-composer-slot");
      assert.ok(slot, "nouveau slot présent");
      assert.equal(slot.contains(inputArea), true, "composer réinséré dans le nouveau slot");
      assert.equal(document.querySelectorAll(".input-area").length, 1, "un seul composer");
    } finally {
      document.body.innerHTML = "";
      document.body.className = "";
    }
  });

  test("cycle complet : hero -> message -> reset -> hero", async () => {
    const { log, inputArea } = setupDom();
    try {
      assert.equal(isHeroMode(), true, "1. hero au chargement");
      document.getElementById("empty-chat-placeholder").remove();
      await wait(20);
      assert.equal(isHeroMode(), false, "2. bas après message");
      assert.equal(inputArea.parentNode.tagName, "MAIN");
      log.innerHTML = EMPTY_HTML; // reset (nouvelle conversation)
      await wait(20);
      assert.equal(isHeroMode(), true, "3. hero après reset");
      assert.ok(document.getElementById("hero-composer-slot").contains(inputArea));
    } finally {
      document.body.innerHTML = "";
      document.body.className = "";
    }
  });
});
