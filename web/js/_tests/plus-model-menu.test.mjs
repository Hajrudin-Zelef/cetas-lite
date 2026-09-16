import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { createRequire } from "node:module";
import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";

const ROOT = join(dirname(fileURLToPath(import.meta.url)), "..", "..");

// jsdom : résolution standard puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
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

const FAMILIES = [
  {
    id: "samagent-nano",
    label: "SamAgent Nano",
    modes: [
      {
        mode: "free", label: "Free",
        pool: [{ provider: "openrouter", model: "openrouter/free", label: "Free Models Router" }],
      },
    ],
  },
  {
    id: "samagent-n4",
    label: "SamAgent N4",
    modes: [
      {
        mode: "flash", label: "Flash",
        pool: [
          { provider: "deepseek", model: "deepseek-flash", label: "DeepSeek V4.1 Flash" },
          { provider: "openrouter", model: "mimo-v2.5-flash", label: "MiMo V2.5 Flash" },
        ],
      },
      {
        mode: "standard", label: "Standard",
        pool: [{ provider: "deepseek", model: "deepseek-v4", label: "DeepSeek V4" }],
      },
    ],
  },
  {
    id: "code",
    label: "Code",
    modes: [
      {
        mode: "flash", label: "Flash", agent: true,
        pool: [{ provider: "deepseek", model: "deepseek-flash", label: "DeepSeek V4.1 Flash" }],
      },
    ],
  },
  {
    id: "samgen",
    label: "SamGen",
    local: true,
    modes: [
      { mode: "nano", label: "Nano (llama.cpp)", engine: "llamacpp", local: true, pool: [] },
      { mode: "n4", label: "N4 (Ollama)", engine: "ollama", local: true, pool: [] },
    ],
  },
];

// api() lit resp.text() : le mock le fournit.
globalThis.fetch = async (url) => {
  const u = String(url);
  const body = u.includes("/api/aliases") ? { families: FAMILIES } : {};
  return { ok: true, status: 200, text: async () => JSON.stringify(body) };
};

document.body.innerHTML = `
  <select id="family-select"></select>
  <select id="mode-select"></select>
  <div id="plus-model-list"></div>
`;

const { initModels, modeOptionsFor, plusMenuFamilies } = await import("../model-select.js");
await initModels();

const groupLabels = () =>
  [...document.querySelectorAll("#plus-model-list .plus-model-group-label")].map((e) => e.textContent);
const groupByLabel = (label) =>
  [...document.querySelectorAll("#plus-model-list .plus-model-group")].find(
    (g) => g.querySelector(".plus-model-group-label").textContent === label
  );
const optionRows = (group) =>
  [...group.querySelectorAll(".plus-model-option")].map((b) => ({
    btn: b,
    mode: b.dataset.mode,
    name: b.querySelector(".plus-model-option-name").textContent,
    rule: b.querySelector(".plus-model-option-rule")
      ? b.querySelector(".plus-model-option-rule").textContent
      : null,
  }));

await test("menu + : 3 groupes (Nano, N4, SamGen), pas de Code", () => {
  assert.deepEqual(groupLabels(), ["Nano", "N4", "SamGen"]);
  assert.ok(!groupLabels().includes("Code"), "le volet Code est supprimé du menu +");
});

await test("menu + : Nano = Auto (fallback) + Modèle — [sélection]", () => {
  const rows = optionRows(groupByLabel("Nano"));
  assert.equal(rows.length, 2);
  assert.equal(rows[0].name, "Auto (fallback)");
  assert.equal(rows[0].mode, "auto");
  assert.equal(rows[0].rule, null);
  assert.equal(rows[1].name, "Modèle");
  assert.equal(rows[1].mode, "free");
  assert.equal(rows[1].rule, "Free Models Router");
  // Au chargement sans préférences : famille Nano, mode Auto.
  assert.ok(rows[0].btn.classList.contains("active"), "Auto est actif par défaut");
});

await test("menu + : N4 = Auto + un rang par mode avec sa sélection", () => {
  const rows = optionRows(groupByLabel("N4"));
  assert.equal(rows.length, 3);
  assert.equal(rows[0].name, "Auto (fallback)");
  assert.equal(rows[0].mode, "auto");
  assert.equal(rows[1].name, "Flash");
  assert.equal(rows[1].rule, "Fallback · 2 modèles");
  assert.equal(rows[2].name, "Standard");
  assert.equal(rows[2].rule, "DeepSeek V4");
  const text = groupByLabel("N4").textContent;
  assert.ok(!text.includes("offres gratuites"), "jamais de règle technique");
});

await test("menu + : SamGen = routes locales, sans Auto", () => {
  const rows = optionRows(groupByLabel("SamGen"));
  assert.equal(rows.length, 2);
  assert.ok(rows.every((r) => r.mode !== "auto"), "pas d'Auto pour SamGen");
  assert.equal(rows[0].name, "Nano (llama.cpp)");
  assert.equal(rows[0].rule, "llama.cpp");
  assert.equal(rows[1].name, "N4 (Ollama)");
  assert.equal(rows[1].rule, "Ollama");
});

await test("menu + : clic sur Auto sélectionne famille + mode auto", () => {
  const rows = optionRows(groupByLabel("N4"));
  rows[0].btn.click();
  const familySel = document.getElementById("family-select");
  const modeSel = document.getElementById("mode-select");
  assert.equal(familySel.value, "samagent-n4");
  assert.equal(modeSel.value, "auto");
  // L'option Auto existe dans le select des modes (familles cloud).
  assert.ok([...modeSel.options].some((o) => o.value === "auto"), "option Auto dans mode-select");
});

await test("mode-select : pas d'option Auto pour SamGen (local)", () => {
  const familySel = document.getElementById("family-select");
  familySel.value = "samgen";
  familySel.dispatchEvent(new Event("change", { bubbles: true }));
  const modeSel = document.getElementById("mode-select");
  assert.ok(![...modeSel.options].some((o) => o.value === "auto"), "pas d'Auto en local");
  assert.deepEqual(
    [...modeSel.options].map((o) => o.value),
    ["nano", "n4"]
  );
});

await test("modeOptionsFor : Auto en tête pour le cloud, rien pour le local", () => {
  const cloud = modeOptionsFor(FAMILIES[1]);
  assert.equal(cloud[0].mode, "auto");
  assert.equal(cloud[0].label, "Auto (fallback)");
  assert.deepEqual(cloud.slice(1).map((m) => m.mode), ["flash", "standard"]);
  const local = modeOptionsFor(FAMILIES[3]);
  assert.deepEqual(local.map((m) => m.mode), ["nano", "n4"]);
});

await test("plusMenuFamilies : ordre Nano/N4/N8/SamGen, sans Code", () => {
  assert.deepEqual(plusMenuFamilies().map((f) => f.id), ["samagent-nano", "samagent-n4", "samgen"]);
});

await test("menu + : le CSS définit toutes les classes de la liste", () => {
  const css = fs.readFileSync(`${ROOT}/css/features/mx-menus.css`, "utf8");
  for (const cls of [
    "plus-model-group",
    "plus-model-group-label",
    "plus-model-option",
    "plus-model-option-main",
    "plus-model-option-name",
    "plus-model-option-rule",
  ]) {
    assert.ok(css.includes("." + cls), `classe .${cls} stylée`);
  }
});

await test("avatar : le dropdown de la sidebar s'ouvre vers le haut", () => {
  const css = fs.readFileSync(`${ROOT}/css/features/mx-menus.css`, "utf8");
  const m = css.match(/#user-menu-dropdown\.user-menu-dropdown\s*\{([^}]*)\}/);
  assert.ok(m, "règle #user-menu-dropdown présente dans mx-menus.css");
  assert.ok(/position:\s*absolute/.test(m[1]), "position absolute (ancré au wrapper)");
  assert.ok(/bottom:\s*calc\(100%/.test(m[1]), "bottom: calc(100% + …) → ouverture vers le haut");
});
