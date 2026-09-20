import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { createRequire } from "node:module";

// jsdom : résolution standard puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

// Harness DOM AVANT d'importer agents.js (markdown.js lit window au chargement).
const dom = new JSDOM("<!doctype html><html><head></head><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.cancelAnimationFrame = dom.window.cancelAnimationFrame.bind(dom.window);
// jsdom ne fournit pas matchMedia : agents.js détecte l'overlay mobile via
// window.matchMedia("(max-width: 1024px)") — stub par défaut (matches: false).
dom.window.matchMedia = () => ({
  matches: false,
  addEventListener() {},
  removeEventListener() {},
});
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;
globalThis.EventSource = class {
  constructor() {}
  close() {}
};
globalThis.fetch = async () => ({ ok: true, status: 200, json: async () => ({}) });

const { positionMenu, initAgents } = await import("../agents.js");

// ---------------------------------------------------------------------------
// 1. positionMenu : mathématiques pures du placement.
// ---------------------------------------------------------------------------
const VIEW = { left: 0, top: 0, width: 1920, height: 1080 };

test("positionMenu : ouvre vers le haut par défaut", () => {
  const btn = { left: 900, top: 1000, right: 960, bottom: 1030 };
  const p = positionMenu(VIEW, btn, 260, 320);
  assert.equal(p.top, 1000 - 320 - 8);
  assert.equal(p.left, 900);
});

test("positionMenu : aligné à droite avec right:true (menu modèle)", () => {
  const btn = { left: 1840, top: 1000, right: 1900, bottom: 1030 };
  const p = positionMenu(VIEW, btn, 260, 320, { right: true });
  assert.equal(p.left, 1900 - 260);
});

test("positionMenu : clampé à gauche", () => {
  const btn = { left: 2, top: 1000, right: 40, bottom: 1030 };
  const p = positionMenu(VIEW, btn, 260, 320);
  assert.equal(p.left, 8);
});

test("positionMenu : repli vers le bas si pas la place en haut", () => {
  const btn = { left: 900, top: 100, right: 960, bottom: 130 };
  const p = positionMenu(VIEW, btn, 260, 320);
  assert.equal(p.top, 130 + 8);
});

test("positionMenu : down:true ouvre vers le bas (sélecteur projet)", () => {
  const btn = { left: 100, top: 60, right: 300, bottom: 90 };
  const p = positionMenu(VIEW, btn, 260, 320, { down: true });
  assert.equal(p.top, 90 + 8);
  assert.equal(p.left, 100);
});

test("positionMenu : menu trop grand clampé dans la vue", () => {
  const btn = { left: 900, top: 1000, right: 960, bottom: 1030 };
  const p = positionMenu(VIEW, btn, 260, 2000);
  assert.equal(p.top, 8);
});

// ---------------------------------------------------------------------------
// 2. câblage : le clic sur + déplace le menu dans le calque et le positionne.
// ---------------------------------------------------------------------------
const sideBtn = document.createElement("button");
sideBtn.className = "dev-module-btn";
sideBtn.dataset.module = "agents";
document.body.appendChild(sideBtn);

initAgents();
window.dispatchEvent(new dom.window.CustomEvent("cetas:open-agents"));
await new Promise((r) => setTimeout(r, 50));

const $ = (s) => document.querySelector(s);

test("le calque de menus existe dans la vue, hors du composer", () => {
  const layer = $("#mx-menu-layer");
  assert.ok(layer, "#mx-menu-layer présent");
  assert.ok(layer.closest("#marex-view"), "dans #marex-view");
  assert.ok(!layer.closest(".mx-composer"), "hors du .mx-composer (plus rogné par son overflow:hidden)");
});

test("clic sur + : menu portalé dans le calque et positionné", () => {
  const btn = $("#mx-plus");
  const menu = $("#mx-menu-plus");
  assert.ok(btn && menu, "bouton et menu + présents");
  btn.click();
  assert.ok(menu.classList.contains("open"), "menu ouvert");
  assert.equal(menu.parentElement.id, "mx-menu-layer", "menu déplacé dans le calque (plus rogné par le composer)");
  assert.ok(menu.style.top !== "" && menu.style.left !== "", "positionné en JS (top/left inline)");
  assert.equal(menu.style.bottom, "auto", "bottom neutralisé");
});

test("second clic : le menu se referme", () => {
  const btn = $("#mx-plus");
  const menu = $("#mx-menu-plus");
  btn.click(); // referme (était ouvert par le test précédent)
  assert.ok(!menu.classList.contains("open"), "menu fermé");
});

test("le sélecteur projet ouvre son menu vers le bas", () => {
  const btn = $("#mx-btn-project");
  const menu = $("#mx-menu-project");
  btn.click();
  assert.ok(menu.classList.contains("open"), "menu projet ouvert");
  assert.equal(menu.parentElement.id, "mx-menu-layer", "menu projet portalé lui aussi");
  btn.click();
});

test("fermeture de la vue (arrête les timers)", () => {
  window.dispatchEvent(new dom.window.CustomEvent("cetas:close-agents"));
  assert.ok(!$("#marex-view").classList.contains("open"), "vue fermée");
});

// ---------------------------------------------------------------------------
// 3. statique : CSS du calque + lien dans index.html.
// ---------------------------------------------------------------------------
const cssMenus = fs.readFileSync(new URL("../../css/features/mx-menus.css", import.meta.url), "utf8");
const indexHtml = fs.readFileSync(new URL("../../index.html", import.meta.url), "utf8");

test("css : le calque laisse passer les clics sauf sur les menus", () => {
  assert.match(cssMenus, /\.mx-menu-layer/);
  assert.match(cssMenus, /pointer-events:\s*none/);
  assert.match(cssMenus, /\.mx-menu-layer \.cdrop-menu[\s\S]*?pointer-events:\s*auto/);
});

test("index.html charge mx-menus.css", () => {
  assert.ok(indexHtml.includes("/css/features/mx-menus.css"), "mx-menus.css lié");
});

// ---------------------------------------------------------------------------
// 3. aria-hidden : le calque de menus ne doit jamais exposer un focus caché.
// ---------------------------------------------------------------------------
test("aria-hidden : masqué au repos, exposé à l'ouverture, re-masqué à la fermeture", () => {
  const layer = $("#mx-menu-layer");
  const btn = $("#mx-btn-perm");
  assert.equal(layer.getAttribute("aria-hidden"), "true", "masqué au repos");
  btn.click(); // ouvre
  assert.equal(layer.getAttribute("aria-hidden"), "false", "exposé quand le menu est ouvert");
  btn.click(); // referme
  assert.equal(layer.getAttribute("aria-hidden"), "true", "re-masqué après fermeture");
});

test("aria-hidden : le focus est rendu au déclencheur quand on ferme au clavier", () => {
  const layer = $("#mx-menu-layer");
  const btn = $("#mx-btn-perm");
  const menu = $("#mx-menu-perm");
  btn.click(); // ouvre
  const item = menu.querySelector(".cdrop-item");
  assert.ok(item, "item de menu focusable présent");
  item.focus();
  assert.equal(document.activeElement, item, "focus sur l'item du menu");
  document.dispatchEvent(new dom.window.KeyboardEvent("keydown", { key: "Escape", bubbles: true }));
  assert.ok(!menu.classList.contains("open"), "menu fermé");
  assert.equal(layer.getAttribute("aria-hidden"), "true", "calque re-masqué");
  assert.equal(document.activeElement, btn, "focus rendu au bouton déclencheur");
});

test("aria-hidden : le focus est rendu au déclencheur après sélection d'un item", () => {
  const layer = $("#mx-menu-layer");
  const btn = $("#mx-btn-perm");
  const menu = $("#mx-menu-perm");
  btn.click(); // ouvre
  const item = menu.querySelector(".cdrop-item");
  item.focus();
  item.click(); // sélectionne -> closeAllDrops() dans le handler
  assert.ok(!menu.classList.contains("open"), "menu fermé");
  assert.equal(layer.getAttribute("aria-hidden"), "true", "calque re-masqué");
  assert.equal(document.activeElement, btn, "focus rendu au bouton déclencheur");
});
