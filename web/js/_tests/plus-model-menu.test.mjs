// Menu + : cascade multi-niveaux.
// Famille › → [Auto (= fallback de l'alias), Mode › → …]
// Mode › → [Auto (= fallback du mode), Models › → modèles] — le fallback,
//           c'est le mode Auto ; en dessous, sélection manuelle du modèle.
// Nano : Free › → groupes par fournisseur (« Free de OpenRouter », « Zen Free »).
import { test, describe } from "node:test";
import assert from "node:assert/strict";
import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
import { createRequire } from "node:module";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = join(dirname(fileURLToPath(import.meta.url)), "..", "..");

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

// Harness DOM AVANT d'importer model-select.js.
const dom = new JSDOM("<!doctype html><html><head></head><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;
globalThis.localStorage = dom.window.localStorage;

const M = (provider, model, label) => ({ provider, model, label, input_per_1m: 0, output_per_1m: 0 });

// Pools par défaut (référence pour buildStaged).
const DEFAULTS = [
  { id: "samagent-nano", label: "SamAgent Nano", modes: [
    { mode: "free", label: "Free", pool: [M("openrouter", "openrouter/free", "Free Model Router")] },
  ]},
  { id: "samagent-n4", label: "SamAgent N4", modes: [
    { mode: "flash", label: "Flash", pool: [
      M("deepseek", "deepseek-flash", "DeepSeek V4.1 Flash"),
      M("openrouter", "mimo-v2.5-flash", "MiMo V2.5 Flash"),
    ]},
    { mode: "standard", label: "Standard", pool: [M("deepseek", "deepseek-v4", "DeepSeek V4")] },
  ]},
  { id: "samgen", label: "SamGen", local: true, modes: [
    { mode: "nano", label: "Nano (llama.cpp)", engine: "llamacpp", local: true, pool: [] },
    { mode: "n4", label: "N4 (Ollama)", engine: "ollama", local: true, pool: [] },
  ]},
];

// Pools effectifs : des overrides sont déjà en place (nano/free : 4 modèles,
// n4/flash : 1 modèle, n4/standard : 2 modèles ≠ défaut).
const FAMILIES = [
  { id: "samagent-nano", label: "SamAgent Nano", modes: [
    { mode: "free", label: "Free", pool: [
      M("openrouter", "openrouter/free", "Free Model Router"),
      M("opencode", "big-pickle", "Big Pickle"),
      M("opencode", "ling-3-flash", "Ling 3.0 Flash Fin Free"),
      M("opencode", "mimo-v2.5-free", "MiMo V2.5 Free"),
    ]},
  ]},
  { id: "samagent-n4", label: "SamAgent N4", modes: [
    { mode: "flash", label: "Flash", pool: [M("deepseek", "deepseek-flash", "DeepSeek V4.1 Flash")] },
    { mode: "standard", label: "Standard", pool: [
      M("deepseek", "deepseek-v4", "DeepSeek V4"),
      M("openrouter", "deepseek-v4-free", "DeepSeek V4 Free"),
    ]},
  ]},
  { id: "samgen", label: "SamGen", local: true, modes: [
    { mode: "nano", label: "Nano (llama.cpp)", engine: "llamacpp", local: true, pool: [] },
    { mode: "n4", label: "N4 (Ollama)", engine: "ollama", local: true, pool: [] },
  ]},
];

const clone = (o) => JSON.parse(JSON.stringify(o));

// PUT capturés pour vérifier la fusion des overrides.
const putBodies = [];
globalThis.fetch = async (url, init = {}) => {
  const u = String(url);
  if (u.includes("/api/aliases") && (init.method || "GET") === "PUT") {
    try { putBodies.push(JSON.parse(init.body)); } catch {}
    return { ok: true, status: 200, text: async () => JSON.stringify({ ok: true, families: clone(FAMILIES) }) };
  }
  const body = u.includes("/api/aliases")
    ? { families: clone(FAMILIES), defaults: clone(DEFAULTS) }
    : {};
  return { ok: true, status: 200, text: async () => JSON.stringify(body) };
};

const mod = await import("../model-select.js");

async function setupDOM() {
  document.body.innerHTML = `
    <select id="family-select"></select>
    <select id="mode-select"></select>
    <div id="plus-model-list"></div>
    <input id="prompt-input" />
    <span id="input-hint"></span>`;
  localStorage.clear();
  putBodies.length = 0;
  await mod.initModels();
}

const visibleSubmenus = () =>
  [...document.querySelectorAll(".plus-model-submenu")].filter((el) => el.style.display === "block");

function hover(el) {
  el.dispatchEvent(new Event("mouseenter", { bubbles: false }));
}

function submenuRow(menu, label) {
  const r = [...menu.querySelectorAll(".plus-model-option")].find(
    (b) => b.querySelector(".plus-model-option-name").textContent === label
  );
  assert.ok(r, `option « ${label} » présente`);
  return r;
}

const rowLabels = (menu) =>
  [...menu.querySelectorAll(".plus-model-option")].map((b) => b.querySelector(".plus-model-option-name").textContent);

// Survole une famille et retourne son sous-menu (niveau 0).
function hoverFamily(famId) {
  const row = document.querySelector(`.plus-model-family-row[data-family="${famId}"]`);
  assert.ok(row, `ligne famille ${famId} présente`);
  hover(row);
  const menus = visibleSubmenus();
  assert.equal(menus.length, 1, "un seul sous-menu visible");
  return { row, menu: menus[0] };
}

describe("menu + : cascade multi-niveaux", () => {
  test("lignes familles : Nano, N4, SamGen (pas de Code), chevron ›", async () => {
    await setupDOM();
    const rows = [...document.querySelectorAll(".plus-model-family-row")];
    assert.deepEqual(
      rows.map((r) => r.dataset.family),
      ["samagent-nano", "samagent-n4", "samgen"]
    );
    assert.deepEqual(
      rows.map((r) => r.querySelector(".plus-model-family-name").textContent),
      ["Nano", "N4", "SamGen"]
    );
    for (const r of rows) assert.equal(r.querySelector(".plus-model-family-chev").textContent, "›");
  });

  test("Nano › : Auto + Free (pas « Modèle — Free Model Router »)", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    assert.deepEqual(rowLabels(menu), ["Auto", "Free"]);
    const free = submenuRow(menu, "Free");
    assert.equal(free.querySelector(".plus-model-option-rule").textContent, "Fallback · 4 modèles");
    assert.ok(free.querySelector(".plus-model-family-chev"), "Free ouvre une cascade");
  });

  test("Nano › Free › : Free de OpenRouter + Zen Free, en cascades", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    hover(submenuRow(menu, "Free"));
    const menus = visibleSubmenus();
    assert.equal(menus.length, 2, "deux niveaux ouverts");
    assert.deepEqual(rowLabels(menus[1]), ["Free de OpenRouter", "Zen Free"]);
  });

  test("Nano › Free › Zen Free › : les modèles gratuits de Zen", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    hover(submenuRow(menu, "Free"));
    hover(submenuRow(visibleSubmenus()[1], "Zen Free"));
    const menus = visibleSubmenus();
    assert.equal(menus.length, 3, "trois niveaux ouverts");
    assert.deepEqual(rowLabels(menus[2]), [
      "Big Pickle",
      "Ling 3.0 Flash Fin Free",
      "MiMo V2.5 Free",
    ]);
    const first = submenuRow(menus[2], "Big Pickle");
    assert.equal(first.querySelector(".plus-model-option-rule").textContent, "opencode");
  });

  test("Nano › Free › Free de OpenRouter › : Free Model Router", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    hover(submenuRow(menu, "Free"));
    hover(submenuRow(visibleSubmenus()[1], "Free de OpenRouter"));
    assert.deepEqual(rowLabels(visibleSubmenus()[2]), ["Free Model Router"]);
  });

  test("N4 › : Auto + Flash + Standard ; Flash › : Auto + Models", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-n4");
    assert.deepEqual(rowLabels(menu), ["Auto", "Flash", "Standard"]);
    assert.equal(submenuRow(menu, "Flash").querySelector(".plus-model-option-rule").textContent, "DeepSeek V4.1 Flash");
    assert.equal(submenuRow(menu, "Standard").querySelector(".plus-model-option-rule").textContent, "Fallback · 2 modèles");
    hover(submenuRow(menu, "Flash"));
    assert.deepEqual(rowLabels(visibleSubmenus()[1]), ["Auto", "Models"]);
  });

  test("N4 › Flash › Models › : tous les modèles du mode (défaut + effectifs)", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-n4");
    hover(submenuRow(menu, "Flash"));
    hover(submenuRow(visibleSubmenus()[1], "Models"));
    // Pool par défaut (2) : l'effectif n'en a qu'1, la liste manuelle
    // propose l'union.
    assert.deepEqual(rowLabels(visibleSubmenus()[2]), [
      "DeepSeek V4.1 Flash",
      "MiMo V2.5 Flash",
    ]);
  });

  test("SamGen : routes directes, pas d'Auto", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samgen");
    assert.deepEqual(rowLabels(menu), ["Nano (llama.cpp)", "N4 (Ollama)"]);
  });

  test("clic Auto d'un mode : fallback du mode, overrides des autres modes préservés", async () => {
    await setupDOM();
    let changed = 0;
    window.addEventListener("cetas:model-changed", () => changed++);
    const { menu } = hoverFamily("samagent-n4");
    hover(submenuRow(menu, "Flash"));
    submenuRow(visibleSubmenus()[1], "Auto").click();
    await new Promise((r) => setTimeout(r, 30));
    assert.equal(putBodies.length, 1, "un seul PUT");
    const body = putBodies[0];
    // Le mode passe en fallback : tout le pool par défaut (2 modèles).
    assert.deepEqual(
      body["samagent-n4"].flash.map((m) => m.provider + "/" + m.model).sort(),
      ["deepseek/deepseek-flash", "openrouter/mimo-v2.5-flash"]
    );
    // Les autres overrides sont préservés (le PUT remplace tout).
    assert.equal(body["samagent-n4"].standard.length, 2, "override standard préservé");
    assert.equal(body["samagent-nano"].free.length, 4, "override nano préservé");
    assert.equal(document.getElementById("family-select").value, "samagent-n4");
    assert.equal(document.getElementById("mode-select").value, "flash");
    assert.equal(changed, 1, "cetas:model-changed émis");
    assert.equal(visibleSubmenus().length, 0, "sous-menus fermés");
  });

  test("clic sur un modèle : « 1 modèle » manuel sur ce modèle", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    hover(submenuRow(menu, "Free"));
    hover(submenuRow(visibleSubmenus()[1], "Zen Free"));
    submenuRow(visibleSubmenus()[2], "Big Pickle").click();
    await new Promise((r) => setTimeout(r, 30));
    assert.equal(putBodies.length, 1);
    const body = putBodies[0];
    assert.deepEqual(body["samagent-nano"].free, [
      { provider: "opencode", model: "big-pickle" },
    ]);
    assert.equal(body["samagent-n4"].standard.length, 2, "autres overrides préservés");
    assert.equal(document.getElementById("family-select").value, "samagent-nano");
    assert.equal(document.getElementById("mode-select").value, "free");
  });

  test("clic Auto d'un alias : sélection auto, sans PUT d'aliases", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    submenuRow(menu, "Auto").click();
    await new Promise((r) => setTimeout(r, 30));
    assert.equal(putBodies.length, 0, "pas de PUT /api/aliases pour l'auto d'alias");
    assert.equal(document.getElementById("mode-select").value, "auto");
  });

  test("changer de famille referme les niveaux profonds", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    hover(submenuRow(menu, "Free"));
    hover(submenuRow(visibleSubmenus()[1], "Zen Free"));
    assert.equal(visibleSubmenus().length, 3);
    hoverFamily("samagent-n4");
    const menus = visibleSubmenus();
    assert.equal(menus.length, 1, "seul le niveau 0 reste");
    assert.deepEqual(rowLabels(menus[0]), ["Auto", "Flash", "Standard"]);
  });

  test("mouseleave programme la fermeture, mouseenter du sous-menu l'annule", async () => {
    await setupDOM();
    const { row, menu } = hoverFamily("samagent-nano");
    row.dispatchEvent(new Event("mouseleave", { bubbles: false }));
    menu.dispatchEvent(new Event("mouseenter", { bubbles: false }));
    await new Promise((r) => setTimeout(r, 220));
    assert.equal(visibleSubmenus().length, 1, "resté ouvert");
    menu.dispatchEvent(new Event("mouseleave", { bubbles: false }));
    await new Promise((r) => setTimeout(r, 220));
    assert.equal(visibleSubmenus().length, 0, "fermé");
  });

  test("plusModelSubmenuContains couvre tous les niveaux / closePlusModelSubmenu", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    hover(submenuRow(menu, "Free"));
    const deep = visibleSubmenus()[1].querySelector(".plus-model-option");
    assert.ok(mod.plusModelSubmenuContains(deep), "niveau 1 couvert");
    assert.ok(!mod.plusModelSubmenuContains(document.body));
    mod.closePlusModelSubmenu();
    assert.equal(visibleSubmenus().length, 0, "tous les niveaux fermés");
  });

  test("CSS : classes de la cascade définies dans mx-menus.css", () => {
    const css = readFileSync(join(ROOT, "css", "features", "mx-menus.css"), "utf8");
    for (const cls of [
      ".plus-model-family",
      ".plus-model-family-row",
      ".plus-model-family-chev",
      ".plus-model-submenu",
      ".plus-model-option",
    ]) {
      assert.ok(css.includes(cls), `classe ${cls} stylée`);
    }
  });
});
