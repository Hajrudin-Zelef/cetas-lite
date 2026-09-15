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

const { initAgents } = await import("../agents.js");
const tb = document.createElement("button");
tb.id = "agents-btn";
document.body.appendChild(tb);
initAgents();
await new Promise((r) => setTimeout(r, 50));

const $ = (s) => document.querySelector(s);
const view = $("#marex-view");

test("composer agent : globe + thinking + effort presents", () => {
  assert.ok(view, "vue agents montée");
  assert.ok($("#mx-web"), "bouton globe");
  assert.ok($("#mx-think"), "bouton thinking");
  assert.ok($("#mx-btn-effort"), "sélecteur effort");
  assert.ok($("#mx-menu-effort"), "menu effort");
});

test("globe : bascule active et persistance", () => {
  const btn = $("#mx-web");
  localStorage.removeItem("mx.mx_web");
  btn.click();
  assert.ok(btn.classList.contains("active"), "actif après clic");
  assert.equal(localStorage.getItem("mx.mx_web"), "1", "persisté");
  assert.equal(btn.getAttribute("aria-pressed"), "true");
  btn.click();
  assert.ok(!btn.classList.contains("active"), "inactif après 2e clic");
  assert.equal(localStorage.getItem("mx.mx_web"), "0");
});

test("thinking : interrupteur réel, N&B éteint / couleur allumé, persisté", () => {
  const btn = $("#mx-think");
  localStorage.removeItem("mx.mx_think");
  // État initial : actif (défaut), couleur, sans verrou.
  assert.ok(btn.classList.contains("active"), "actif par défaut");
  assert.ok(!btn.classList.contains("locked"), "cliquable, plus verrouillé");
  assert.ok(!btn.classList.contains("off"), "pas en N&B");
  assert.equal(btn.getAttribute("aria-pressed"), "true");
  // Clic : extinction -> noir et blanc.
  btn.click();
  assert.ok(btn.classList.contains("off"), "N&B quand éteint");
  assert.ok(!btn.classList.contains("active"), "plus de couleur");
  assert.equal(btn.getAttribute("aria-pressed"), "false");
  assert.equal(localStorage.getItem("mx.mx_think"), "0", "persisté éteint");
  // Re-clic : rallumage.
  btn.click();
  assert.ok(btn.classList.contains("active"), "couleur quand allumé");
  assert.ok(!btn.classList.contains("off"), "plus de N&B");
  assert.equal(localStorage.getItem("mx.mx_think"), "1", "persisté allumé");
});

test("effort : 4 options, labels FR, persistance", () => {
  const menu = $("#mx-menu-effort");
  const items = menu.querySelectorAll(".cdrop-item[data-effort]");
  assert.equal(items.length, 4, "4 options d'effort");
  const want = { default: "Défaut", low: "Faible", medium: "Moyen", high: "Max" };
  for (const it of items) {
    it.click();
    const v = it.dataset.effort;
    assert.equal(localStorage.getItem("mx.mx_effort"), v, "persisté " + v);
    assert.equal($("#mx-label-effort").textContent, want[v], "label " + v);
  }
});

test("thread-view : événement system -> message système", async () => {
  const { ThreadView } = await import("../thread-view.js");
  const log = document.createElement("div");
  document.body.appendChild(log);
  const view2 = new ThreadView({ log, renderer: { render() {}, update() {}, finish() {}, reset() {} } });
  view2.handleEvent({ system: "MAREX.md chargé" });
  assert.ok(log.querySelector(".msg-system"), "message système inséré");
  assert.match(log.textContent, /MAREX\.md chargé/);
});
