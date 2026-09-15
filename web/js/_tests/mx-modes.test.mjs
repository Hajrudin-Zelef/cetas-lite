import test from "node:test";
import assert from "node:assert/strict";
// ESM : NODE_PATH ne s'applique pas aux imports, chemin absolu requis.
import { JSDOM } from "/tmp/node_modules/jsdom/lib/api.js";

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.cancelAnimationFrame = dom.window.cancelAnimationFrame.bind(dom.window);
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;
globalThis.EventSource = class {
  constructor() {}
  close() {}
};
globalThis.fetch = async (url) => {
  const u = String(url);
  const json = (body) => ({ ok: true, status: 200, json: async () => body });
  if (u.includes("/api/aliases")) {
    return json({ families: [{ id: "f1", label: "F1", modes: [{ mode: "code", label: "Code", agent: true }] }] });
  }
  if (u.includes("/api/me")) return json({ username: "sam" });
  if (u.includes("/api/metrics")) return json({ cpu: 1, mem: 2, disk: 3 });
  return json({});
};

const { initAgents, renderMxTodos, clearMxTodos } = await import("../agents.js");
const tb = document.createElement("button");
tb.id = "agents-btn";
document.body.appendChild(tb);
localStorage.removeItem("mx.perm");
initAgents();
await new Promise((r) => setTimeout(r, 50));

const $ = (s) => document.querySelector(s);
const tabOnInput = () => {
  const input = $("#mx-input");
  input.dispatchEvent(new dom.window.KeyboardEvent("keydown", { key: "Tab", bubbles: true, cancelable: true }));
};

test("mode : BUILD par défaut (barre bleue, badge, permission)", () => {
  const badge = $("#mx-mode-badge");
  const composer = $("#mx-composer");
  assert.ok(badge, "badge présent");
  assert.equal(badge.textContent, "BUILD");
  assert.ok(badge.classList.contains("build"), "badge bleu");
  assert.ok(composer.classList.contains("mode-build"), "composer mode-build");
  assert.ok(!composer.classList.contains("mode-plan"), "pas mode-plan");
  assert.equal($("#mx-label-perm").textContent, "Espace Write");
});

test("Tab : bascule vers PLAN (barre jaune, lecture seule)", () => {
  tabOnInput();
  const badge = $("#mx-mode-badge");
  const composer = $("#mx-composer");
  assert.equal(badge.textContent, "PLAN");
  assert.ok(badge.classList.contains("plan"), "badge jaune");
  assert.ok(composer.classList.contains("mode-plan"), "composer mode-plan");
  assert.ok(!composer.classList.contains("mode-build"), "plus mode-build");
  assert.equal($("#mx-label-perm").textContent, "Read only", "permission synchronisée");
  assert.equal(localStorage.getItem("mx.perm"), "read", "persisté");
});

test("Tab : retour en BUILD, restaure la permission d'écriture", () => {
  tabOnInput();
  assert.equal($("#mx-mode-badge").textContent, "BUILD");
  assert.ok($("#mx-composer").classList.contains("mode-build"));
  assert.equal($("#mx-label-perm").textContent, "Espace Write", "permission restaurée");
  assert.equal(localStorage.getItem("mx.perm"), "write");
});

test("clic sur le badge : même bascule que Tab", () => {
  $("#mx-mode-badge").click();
  assert.equal($("#mx-mode-badge").textContent, "PLAN");
  $("#mx-mode-badge").click();
  assert.equal($("#mx-mode-badge").textContent, "BUILD");
});

test("Tab depuis 'Ask permission' : restaure 'Ask permission' en Build", () => {
  // Choisir Ask permission via le menu.
  const items = [...document.querySelectorAll('#mx-menu-perm .cdrop-item[data-perm]')];
  const ask = items.find((it) => it.dataset.perm === "ask");
  assert.ok(ask, "item ask présent");
  ask.click();
  assert.equal($("#mx-label-perm").textContent, "Ask permission");
  tabOnInput();
  assert.equal($("#mx-mode-badge").textContent, "PLAN");
  tabOnInput();
  assert.equal($("#mx-mode-badge").textContent, "BUILD");
  assert.equal($("#mx-label-perm").textContent, "Ask permission", "ask restauré, pas write");
});

test("todos : panneau masqué par défaut", () => {
  const panel = $("#mx-todos");
  assert.ok(panel, "panneau présent");
  assert.equal(panel.style.display, "none");
});

test("todos : rendu avec statuts, progression et échappement HTML", () => {
  renderMxTodos([
    { content: "explorer le code", status: "completed" },
    { content: "implémenter <b>le mode</b>", status: "in_progress" },
    { content: "tester", status: "pending" },
  ]);
  const panel = $("#mx-todos");
  assert.notEqual(panel.style.display, "none", "visible");
  assert.match(panel.innerHTML, /1\/3/, "progression 1/3");
  const items = panel.querySelectorAll(".mx-todo-item");
  assert.equal(items.length, 3);
  assert.ok(items[0].classList.contains("todo-completed"));
  assert.ok(items[1].classList.contains("todo-in_progress"));
  assert.ok(items[2].classList.contains("todo-pending"));
  assert.ok(items[0].querySelector(".mx-todo-mark").textContent.includes("✓"));
  assert.ok(!panel.innerHTML.includes("<b>le mode</b>"), "HTML échappé");
});

test("todos : liste vide et clear masquent le panneau", () => {
  renderMxTodos([]);
  assert.equal($("#mx-todos").style.display, "none");
  renderMxTodos([{ content: "x", status: "pending" }]);
  assert.notEqual($("#mx-todos").style.display, "none");
  clearMxTodos();
  assert.equal($("#mx-todos").style.display, "none");
  assert.equal($("#mx-todos").innerHTML, "");
});
