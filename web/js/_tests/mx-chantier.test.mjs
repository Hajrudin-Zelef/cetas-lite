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
// Mock fetch : text() requis (api() lit resp.text()).
let skillsCalls = 0;
globalThis.fetch = async (url, opts = {}) => {
  const u = String(url);
  const json = (body) => ({ ok: true, status: 200, text: async () => JSON.stringify(body) });
  if (u.includes("/api/aliases")) {
    return json({
      families: [
        {
          id: "code",
          label: "Code",
          modes: [
            { mode: "flash", label: "Flash", agent: true },
            { mode: "standard", label: "Standard", agent: true },
            { mode: "elite", label: "Elite", agent: true },
          ],
        },
      ],
    });
  }
  if (u.includes("/api/skills")) {
    skillsCalls++;
    return json({
      skills: [
        { id: "s1", name: "revue", description: "d", instructions: "i", enabled: true },
        { id: "s2", name: "tests", description: "d", instructions: "i", enabled: false },
      ],
    });
  }
  if (u.includes("/api/projects/active")) return json({ active: "" });
  if (u.includes("/api/projects")) {
    return json({
      projects: [
        { id: "p1", name: "local-proj", mode: "local" },
        { id: "p2", name: "mon-serveur", mode: "sftp", user: "u", host: "h.example", remote_path: "/srv" },
      ],
    });
  }
  if (u.includes("/api/me")) return json({ username: "sam" });
  if (u.includes("/api/metrics")) return json({ cpu: 1, mem: 2, disk: 3 });
  if (u.includes("/api/agents")) return json({ agents: [] });
  if (u.includes("/api/capabilities")) return json({ caps: {} });
  return json({});
};

const { initAgents } = await import("../agents.js");
const tb = document.createElement("button");
tb.id = "agents-btn";
document.body.appendChild(tb);
localStorage.clear();
initAgents();
await new Promise((r) => setTimeout(r, 150));

const $ = (s) => document.querySelector(s);
const indexHtml = fs.readFileSync(new URL("../../index.html", import.meta.url), "utf8");

// ---------------------------------------------------------------------------
// 1. Menu + restructuré.
// ---------------------------------------------------------------------------
test("menu + : les 5 sections demandées", () => {
  const menu = $("#mx-menu-plus");
  assert.ok(menu, "menu + présent");
  const t = menu.textContent;
  for (const s of ["Fichiers", "Projet", "Compétences", "Recherche web", "Plugins"]) {
    assert.ok(t.includes(s), "section « " + s + " »");
  }
  assert.ok(t.includes("Ajouter des fichiers"), "action fichiers");
  assert.ok(t.includes("Ajouter un projet"), "action projet");
  assert.ok(t.includes("Ajouter des plugins"), "action plugins");
  assert.ok(t.includes("Worktree git isolé"), "toggle worktree conservé");
});

test("menu + : compteur de compétences actives", () => {
  const t = $("#mx-menu-plus").textContent;
  assert.ok(t.includes("1 active / 2"), "compteur 1 active / 2, obtenu: " + t.slice(0, 200));
});

test("menu + : choisir une profondeur allume le globe", () => {
  const globe = $("#mx-web");
  // Éteindre le globe d'abord.
  if (globe.classList.contains("active")) globe.click();
  assert.ok(!globe.classList.contains("active"), "globe éteint");
  const menu = $("#mx-menu-plus");
  menu.querySelector('.cdrop-item[data-action="webdepth"][data-v="deep"]').click();
  assert.ok(globe.classList.contains("active"), "globe rallumé par le choix deep");
  assert.equal(localStorage.getItem("mx.mx_web"), "1", "mx_web persisté");
});

// ---------------------------------------------------------------------------
// 2. Sélecteur de modèle : sans le préfixe famille.
// ---------------------------------------------------------------------------
test("menu modèle : exactement Flash / Standard / Elite, sans en-tête de famille", async () => {
  $("#agents-btn").click(); // openView() -> loadFamilies() -> buildModelMenu()
  await new Promise((r) => setTimeout(r, 100));
  const menu = $("#mx-menu-model");
  $("#mx-back").click(); // closeView() : coupe les setInterval
  const items = menu.querySelectorAll('.cdrop-item[data-f]');
  assert.equal(items.length, 3, "3 modes, obtenu: " + items.length);
  const labels = [...items].map((it) => it.textContent);
  for (const l of ["Flash", "Standard", "Elite"]) {
    assert.ok(labels.some((t) => t.includes(l)), "mode " + l);
  }
  assert.equal(menu.querySelectorAll(".cdrop-section-label").length, 0, "aucun en-tête de famille");
  // Les identifiants internes famille + mode sont conservés.
  const flash = menu.querySelector('.cdrop-item[data-m="flash"]');
  assert.equal(flash.dataset.f, "code", "famille interne conservée");
});

// ---------------------------------------------------------------------------
// 3. Menu projet : recherche + volet Remote SFTP.
// ---------------------------------------------------------------------------
test("menu projet : champ de recherche présent", () => {
  const menu = $("#mx-menu-project");
  assert.ok(menu.querySelector("#mx-project-filter"), "champ de recherche");
});

test("menu projet : volet Remote avec le serveur SFTP", () => {
  const t = $("#mx-menu-project").textContent;
  assert.ok(t.includes("Remote"), "section Remote");
  assert.ok(t.includes("mon-serveur"), "serveur SFTP listé");
  assert.ok(t.includes("Serveur distant (SFTP)"), "bouton d'ajout SFTP");
});

// ---------------------------------------------------------------------------
// 4. Pièces jointes : chips + input + alerte vision.
// ---------------------------------------------------------------------------
test("composer : zone de pièces jointes et input fichier", () => {
  assert.ok($("#mx-attach-row"), "ligne de chips");
  assert.ok($("#mx-file-input"), "input fichier caché");
  assert.ok($("#mx-vision-warn"), "zone d'alerte vision");
});

// ---------------------------------------------------------------------------
// 5. Configuration : onglets Remote et Compétences.
// ---------------------------------------------------------------------------
test("config : onglets Remote et Compétences + panneaux", () => {
  assert.ok(indexHtml.includes('data-tab="remote"'), "onglet Remote");
  assert.ok(indexHtml.includes('data-tab="competences"'), "onglet Compétences");
  assert.ok(indexHtml.includes('id="panel-remote"'), "panneau Remote");
  assert.ok(indexHtml.includes('id="panel-competences"'), "panneau Compétences");
  assert.ok(indexHtml.includes('id="remote-body"'), "corps Remote");
  assert.ok(indexHtml.includes('id="skills-body"'), "corps Compétences");
});

test("compétences : l'événement cetas:skills-changed recharge le menu +", async () => {
  const before = skillsCalls;
  window.dispatchEvent(new dom.window.CustomEvent("cetas:skills-changed"));
  await new Promise((r) => setTimeout(r, 100));
  assert.ok(skillsCalls > before, "nouvel appel /api/skills après l'événement");
  assert.ok($("#mx-menu-plus").textContent.includes("1 active / 2"), "compteur reconstruit");
});

test("modals.js : chargement paresseux des onglets", () => {
  const src = fs.readFileSync(new URL("../modals.js", import.meta.url), "utf8");
  assert.ok(src.includes("loadRemoteTab"), "loadRemoteTab défini");
  assert.ok(src.includes("loadSkillsTab"), "loadSkillsTab défini");
  assert.ok(src.includes('openNewProjectModal'), "réutilise la modale projet");
  assert.ok(src.includes('/api/skills'), "API skills utilisée");
});
