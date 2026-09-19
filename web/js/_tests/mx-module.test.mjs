import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { createRequire } from "node:module";
// jsdom : résolution standard (node_modules du projet) puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

// ---------------------------------------------------------------------------
// 1. Structure du HTML livré : métriques au-dessus du module Agent.
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

test("index.html référence le CSS du module", () => {
  assert.ok(indexHtml.includes("/css/features/mx-module.css"), "mx-module.css lié");
});

// ---------------------------------------------------------------------------
// 2. Comportement : le bouton module Agents ouvre la page complète
//    Marexcode en plein écran (overlay), pas un panneau intégré.
// ---------------------------------------------------------------------------
const cssModule = fs.readFileSync(new URL("../../css/features/mx-module.css", import.meta.url), "utf8");
// Extrait minimal des règles générées (positionnement plein écran d'origine).
const cssGenerated = "#marex-view{position:fixed;inset:0;z-index:400;display:none;}#marex-view.open{display:block;}";

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

const { initAgents } = await import("../agents.js");
initAgents();
await new Promise((r) => setTimeout(r, 50));

const $ = (s) => document.querySelector(s);
const toggle = () => window.dispatchEvent(new dom.window.CustomEvent("cetas:toggle-agents"));

test("toggle : ouvre la page Agents en plein écran", () => {
  toggle();
  const view = $("#marex-view");
  assert.ok(view, "#marex-view créée");
  assert.equal(view.parentElement, document.body, "vue rattachée au body (page complète)");
  assert.ok(view.classList.contains("open"), "vue ouverte");
  assert.ok(sideBtn.classList.contains("active"), "bouton module actif");
  assert.equal(getComputedStyle(view).position, "fixed", "plein écran (position:fixed)");
  assert.equal(document.body.style.overflow, "hidden", "scroll du fond bloqué");
});

test("toggle : referme la page et revient au chat", () => {
  toggle();
  const view = $("#marex-view");
  assert.ok(!view.classList.contains("open"), "vue fermée");
  assert.ok(!sideBtn.classList.contains("active"), "bouton module inactif");
  assert.equal(document.body.style.overflow, "", "scroll restauré");
});

test("cetas:open-agents ouvre, cetas:close-agents referme (retour chat)", () => {
  window.dispatchEvent(new dom.window.CustomEvent("cetas:open-agents"));
  assert.ok($("#marex-view").classList.contains("open"), "ouvert via open-agents");
  assert.equal(getComputedStyle($("#marex-view")).position, "fixed", "toujours plein écran");
  window.dispatchEvent(new dom.window.CustomEvent("cetas:close-agents"));
  assert.ok(!$("#marex-view").classList.contains("open"), "fermé via close-agents");
});

// ---------------------------------------------------------------------------
// 3. Sidebar : sections Projets / Discussions restructurées (sans doublons).
// ---------------------------------------------------------------------------
const { Projects } = await import("../projects.js");

test("sidebar projets : structure aplatie, sans doublons", () => {
  const view = $("#marex-view");
  assert.equal(view.querySelector("#mx-toggle-active"), null, "plus de toggle « Projet actif »");
  assert.equal(view.querySelector("#mx-panel-active"), null, "plus de panneau « Projet actif »");
  const tProj = view.querySelector("#mx-toggle-projects");
  assert.ok(tProj, "toggle Projets présent");
  assert.match(tProj.textContent, /Projets/);
  assert.ok(tProj.classList.contains("sb-group-toggle"), "l'intitulé de groupe est le toggle");
  const tDisc = view.querySelector("#mx-toggle-disc");
  assert.ok(tDisc, "toggle Discussions présent");
  assert.match(tDisc.textContent, /Discussions/);
  assert.ok(!view.innerHTML.includes("Toutes les discussions"), "plus d'intitulé « Toutes les discussions »");
  assert.ok(!view.innerHTML.includes("Mes projets"), "plus d'intitulé « Mes projets »");
  assert.ok(!view.innerHTML.includes("Aucun projet actif"), "plus de doublon « Aucun projet actif »");
});

test("sidebar projets : état vide (espace partagé, pas de projet)", () => {
  const view = $("#marex-view");
  assert.equal(view.querySelector("#mx-active-name").textContent, "Espace partagé");
  assert.ok(view.querySelector("#mx-active-badge").hidden, "badge Actif masqué sans projet");
  assert.ok(view.querySelector("#mx-active-sub").hidden, "détail masqué sans projet");
  assert.equal(view.querySelector("#mx-active-tree").style.display, "none", "arbre masqué sans projet");
  assert.notEqual(view.querySelector("#mx-projects-empty").style.display, "none", "« Aucun projet importé » visible");
  assert.equal(view.querySelector("#mx-projects-label").style.display, "none", "« Tous les projets » masqué quand vide");
  assert.ok(view.querySelector("#mx-new-project"), "bouton + Nouveau projet présent");
});

test("sidebar projets : projet actif affiché avec badge et détail", () => {
  Projects.list = [
    { id: "p1", name: "MonProjet", mode: "local" },
    { id: "p2", name: "ServeurProd", mode: "sftp", user: "deploy", host: "srv.ex", remote_path: "/srv/api" },
  ];
  Projects.activeId = "p2";
  Projects.onChange();
  const view = $("#marex-view");
  assert.equal(view.querySelector("#mx-active-name").textContent, "ServeurProd");
  assert.ok(!view.querySelector("#mx-active-badge").hidden, "badge Actif visible");
  assert.equal(view.querySelector("#mx-active-sub").textContent, "deploy@srv.ex:/srv/api");
  assert.notEqual(view.querySelector("#mx-active-tree").style.display, "none", "arbre visible avec projet");
  assert.equal(view.querySelector("#mx-projects-empty").style.display, "none", "état vide masqué");
  assert.notEqual(view.querySelector("#mx-projects-label").style.display, "none", "« Tous les projets » visible");
  const rows = [...view.querySelectorAll("#mx-projects-list .mx-project-row")];
  assert.equal(rows.length, 2, "les deux projets listés");
  assert.ok(rows[1].classList.contains("active"), "projet actif surligné dans la liste");
  // Nettoyage : retour à l'état vide.
  Projects.list = [];
  Projects.activeId = "";
  Projects.onChange();
});

test("sidebar : les toggles replient les sections", () => {
  const tProj = $("#mx-toggle-projects");
  const panel = $("#mx-panel-projects");
  tProj.click();
  assert.ok(panel.classList.contains("hidden"), "section Projets repliée");
  assert.ok(tProj.querySelector("svg").classList.contains("collapsed"), "chevron pivoté");
  tProj.click();
  assert.ok(!panel.classList.contains("hidden"), "section Projets dépliée");
  const tDisc = $("#mx-toggle-disc");
  const dPanel = $("#mx-panel-disc");
  tDisc.click();
  assert.ok(dPanel.classList.contains("hidden"), "section Discussions repliée");
  tDisc.click();
  assert.ok(!dPanel.classList.contains("hidden"), "section Discussions dépliée");
});
