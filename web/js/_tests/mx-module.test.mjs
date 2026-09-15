import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { JSDOM } from "jsdom";

// ---------------------------------------------------------------------------
// 1. Structure du HTML livré : métriques au-dessus du module Agent,
//    conteneur #module-agents dans <main>.
// ---------------------------------------------------------------------------
const indexHtml = fs.readFileSync(new URL("../../index.html", import.meta.url), "utf8");

test("sidebar : MÉTRIQUES placées juste au-dessus du module Agent", () => {
  const iMetrics = indexHtml.indexOf('id="sb-metrics"');
  const iModules = indexHtml.indexOf('class="dev-modules"');
  const iAgentBtn = indexHtml.indexOf('data-module="agents"');
  assert.ok(iMetrics > -1, "#sb-metrics présent");
  assert.ok(iModules > -1 && iAgentBtn > -1, "modules présents");
  assert.ok(iMetrics < iModules, "métriques avant les boutons modules");
  assert.ok(iMetrics < iAgentBtn, "métriques avant le bouton Agent");
});

test("main : conteneur #module-agents présent et masqué par défaut", () => {
  const iMain = indexHtml.indexOf("<main");
  const iMainEnd = indexHtml.indexOf("</main>");
  const iMod = indexHtml.indexOf('id="module-agents"');
  assert.ok(iMain > -1 && iMainEnd > iMain, "<main> présent");
  assert.ok(iMod > iMain && iMod < iMainEnd, "#module-agents dans <main>");
  assert.ok(/id="module-agents"[^>]*hidden/.test(indexHtml), "#module-agents masqué par défaut");
});

test("index.html référence le CSS du module", () => {
  assert.ok(indexHtml.includes("/css/features/mx-module.css"), "mx-module.css lié");
});

// ---------------------------------------------------------------------------
// 2. Comportement : la vue agents vit dans le module (plus d'overlay).
// ---------------------------------------------------------------------------
const cssModule = fs.readFileSync(new URL("../../css/features/mx-module.css", import.meta.url), "utf8");
// Extrait minimal des règles générées (positionnement overlay d'origine).
const cssGenerated = "#marex-view{position:fixed;inset:0;z-index:400;display:none;}#marex-view.open{display:block;}";

const dom = new JSDOM("<!doctype html><html><head></head><body></body></html>", {
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

const style = document.createElement("style");
style.textContent = cssGenerated + "\n" + cssModule;
document.head.appendChild(style);

const sideBtn = document.createElement("button");
sideBtn.className = "dev-module-btn";
sideBtn.dataset.module = "agents";
document.body.appendChild(sideBtn);

const main = document.createElement("main");
main.className = "main";
const chatContainer = document.createElement("div");
chatContainer.id = "chat-container";
main.appendChild(chatContainer);
const host = document.createElement("div");
host.id = "module-agents";
host.className = "module-view";
host.hidden = true;
main.appendChild(host);
document.body.appendChild(main);

const { initAgents } = await import("../agents.js");
initAgents();
await new Promise((r) => setTimeout(r, 50));

const $ = (s) => document.querySelector(s);
const toggle = () => window.dispatchEvent(new dom.window.CustomEvent("cetas:toggle-agents"));

test("toggle : ouvre le module Agent (vue intégrée, chat masqué)", () => {
  toggle();
  const view = $("#marex-view");
  assert.ok(view, "#marex-view créée");
  assert.equal(view.parentElement.id, "module-agents", "vue hébergée par #module-agents");
  assert.ok(view.classList.contains("open"), "vue ouverte");
  assert.equal(host.hidden, false, "conteneur module visible");
  assert.ok(main.classList.contains("module-agents-active"), "main marqué module-agents-active");
  assert.ok(sideBtn.classList.contains("active"), "bouton module actif");
  assert.equal(getComputedStyle(view).position, "relative", "plus de position:fixed (fini l'overlay)");
});

test("toggle : referme le module et revient au chat", () => {
  toggle();
  const view = $("#marex-view");
  assert.ok(!view.classList.contains("open"), "vue fermée");
  assert.equal(host.hidden, true, "conteneur module masqué");
  assert.ok(!main.classList.contains("module-agents-active"), "main sans marqueur");
  assert.ok(!sideBtn.classList.contains("active"), "bouton module inactif");
});

test("cetas:open-agents ouvre, cetas:close-agents referme (retour chat)", () => {
  window.dispatchEvent(new dom.window.CustomEvent("cetas:open-agents"));
  assert.ok($("#marex-view").classList.contains("open"), "ouvert via open-agents");
  assert.ok(main.classList.contains("module-agents-active"), "main marqué");
  window.dispatchEvent(new dom.window.CustomEvent("cetas:close-agents"));
  assert.ok(!$("#marex-view").classList.contains("open"), "fermé via close-agents");
  assert.ok(!main.classList.contains("module-agents-active"), "main dé-marqué");
});
