// Menu + : cascade multi-niveaux.
// Famille › → [Défaut (= routage par effort de l'alias, valeur "auto"), Mode › → …]
//   (Défaut affiché seulement si l'alias a plus d'un mode ; Nano n'a que Free)
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
const settingsPuts = [];
let mockPrefs = {};
globalThis.fetch = async (url, init = {}) => {
  const u = String(url);
  if (u.includes("/api/settings") && (init.method || "GET") === "PUT") {
    try { settingsPuts.push(JSON.parse(init.body)); } catch {}
    Object.assign(mockPrefs, JSON.parse(init.body));
    return { ok: true, status: 200, text: async () => JSON.stringify({ ok: true }) };
  }
  if (u.includes("/api/settings")) {
    return { ok: true, status: 200, text: async () => JSON.stringify(clone(mockPrefs)) };
  }
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
    <select id="effort-select">
      <option value="default">Défaut</option>
      <option value="low">Faible</option>
      <option value="medium">Moyen</option>
      <option value="high">Max</option>
    </select>
    <div id="plus-effort-pills">
      <button type="button" class="plus-menu-pill" data-effort="default">Défaut</button>
      <button type="button" class="plus-menu-pill" data-effort="low">Faible</button>
      <button type="button" class="plus-menu-pill" data-effort="medium">Moyen</button>
      <button type="button" class="plus-menu-pill" data-effort="high">Max</button>
    </div>
    <div id="plus-model-list"></div>
    <input id="prompt-input" />
    <span id="input-hint"></span>`;
  localStorage.clear();
  putBodies.length = 0;
  settingsPuts.length = 0;
  mockPrefs = {};
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

  test("Nano › : Free seul (Défaut retiré : mode unique, doublon exact)", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    assert.deepEqual(rowLabels(menu), ["Free"]);
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

  test("N4 › : Défaut + Flash + Standard ; Flash › : Auto + Models", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-n4");
    assert.deepEqual(rowLabels(menu), ["Défaut", "Flash", "Standard"]);
    const defaut = submenuRow(menu, "Défaut");
    assert.equal(defaut.querySelector(".plus-model-option-rule").textContent, "Routage par effort");
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

  test("clic Défaut d'un alias : routage par effort, sans PUT d'aliases", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-n4");
    submenuRow(menu, "Défaut").click();
    await new Promise((r) => setTimeout(r, 30));
    assert.equal(putBodies.length, 0, "pas de PUT /api/aliases pour le Défaut d'alias");
    assert.equal(document.getElementById("mode-select").value, "auto", "valeur interne inchangée");
    const opt = document.querySelector('#mode-select option[value="auto"]');
    assert.equal(opt.textContent, "Défaut", "libellé du sélecteur : Défaut");
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
    assert.deepEqual(rowLabels(menus[0]), ["Défaut", "Flash", "Standard"]);
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

  test("pills d'effort : clic → select + persistance thinking_effort", async () => {
    await setupDOM();
    const pill = document.querySelector('#plus-effort-pills .plus-menu-pill[data-effort="high"]');
    pill.click();
    await new Promise((r) => setTimeout(r, 30));
    assert.equal(document.getElementById("effort-select").value, "high");
    assert.ok(pill.classList.contains("active"), "pill cliqué actif");
    assert.equal(
      document.querySelectorAll('#plus-effort-pills .plus-menu-pill.active').length,
      1,
      "un seul pill actif"
    );
    assert.ok(settingsPuts.length >= 1, "PUT /api/settings émis");
    assert.equal(settingsPuts[settingsPuts.length - 1].thinking_effort, "high");
  });

  test("pills d'effort : changement du select → pill actif synchronisé", async () => {
    await setupDOM();
    const sel = document.getElementById("effort-select");
    sel.value = "low";
    sel.dispatchEvent(new Event("change", { bubbles: true }));
    await new Promise((r) => setTimeout(r, 30));
    const pill = document.querySelector('#plus-effort-pills .plus-menu-pill[data-effort="low"]');
    assert.ok(pill.classList.contains("active"), "pill low actif après change");
    assert.equal(settingsPuts[settingsPuts.length - 1].thinking_effort, "low");
  });

  test("pills d'effort : restauration depuis les prefs au chargement", async () => {
    mockPrefs = { thinking_effort: "medium" };
    // Recharge le DOM avec des prefs pré-remplies (sans réinitialiser mockPrefs).
    document.body.innerHTML = `
      <select id="family-select"></select>
      <select id="mode-select"></select>
      <select id="effort-select">
        <option value="default">Défaut</option>
        <option value="low">Faible</option>
        <option value="medium">Moyen</option>
        <option value="high">Max</option>
      </select>
      <div id="plus-effort-pills">
        <button type="button" class="plus-menu-pill" data-effort="default">Défaut</button>
        <button type="button" class="plus-menu-pill" data-effort="low">Faible</button>
        <button type="button" class="plus-menu-pill" data-effort="medium">Moyen</button>
        <button type="button" class="plus-menu-pill" data-effort="high">Max</button>
      </div>
      <div id="plus-model-list"></div>
      <input id="prompt-input" />
      <span id="input-hint"></span>`;
    await mod.initModels();
    assert.equal(document.getElementById("effort-select").value, "medium");
    assert.ok(
      document.querySelector('#plus-effort-pills .plus-menu-pill[data-effort="medium"]').classList.contains("active"),
      "pill medium actif depuis les prefs"
    );
  });
});
