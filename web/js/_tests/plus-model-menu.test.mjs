// Menu + : disposition en cascade — une ligne par famille (Nano / N4 / SamGen,
// pas de Code), survol => sous-menu à droite avec "Auto (fallback)" puis les
// modes. Vérifie la structure DOM, le contenu des sous-menus et la sélection.
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

// api() lit resp.text() : le mock le fournit. Les PUT sont capturés pour
// vérifier la persistance.
const putBodies = [];
globalThis.fetch = async (url, opts) => {
  const u = String(url);
  if (opts && opts.method === "PUT" && opts.body) {
    try { putBodies.push(JSON.parse(opts.body)); } catch {}
  }
  const body = u.includes("/api/aliases") ? { families: FAMILIES } : {};
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
  await mod.initModels();
}

// Sous-menu visible pour une famille (simule le survol).
function hoverFamily(famId) {
  const row = document.querySelector(`.plus-model-family-row[data-family="${famId}"]`);
  assert.ok(row, `ligne famille ${famId} présente`);
  row.dispatchEvent(new Event("mouseenter", { bubbles: false }));
  const menu = document.getElementById("plus-model-submenu");
  assert.ok(menu, "sous-menu créé");
  assert.equal(menu.style.display, "block", "sous-menu visible au survol");
  return { row, menu };
}

const optionRows = (menu) =>
  [...menu.querySelectorAll(".plus-model-option")].map((b) => ({
    btn: b,
    mode: b.dataset.mode,
    name: b.querySelector(".plus-model-option-name").textContent,
    rule: b.querySelector(".plus-model-option-rule")
      ? b.querySelector(".plus-model-option-rule").textContent
      : null,
  }));

describe("menu + : disposition en cascade (familles + sous-menus)", () => {
  test("une ligne par famille, pas de Code", async () => {
    await setupDOM();
    const rows = [...document.querySelectorAll(".plus-model-family-row")];
    assert.deepEqual(
      rows.map((r) => r.dataset.family),
      ["samagent-nano", "samagent-n4", "samgen"],
      "Nano, N4, SamGen uniquement (pas de Code)",
    );
    assert.deepEqual(
      rows.map((r) => r.querySelector(".plus-model-family-name").textContent),
      ["Nano", "N4", "SamGen"],
    );
    for (const r of rows) {
      const chev = r.querySelector(".plus-model-family-chev");
      assert.ok(chev, "chevron présent");
      assert.equal(chev.textContent, "›");
    }
  });

  test("survol Nano : Auto (fallback) + Modèle — [sélection effective]", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    const rows = optionRows(menu);
    assert.equal(rows.length, 2);
    assert.equal(rows[0].mode, "auto");
    assert.equal(rows[0].name, "Auto (fallback)");
    assert.equal(rows[0].rule, null, "pas de règle pour Auto");
    assert.ok(rows[0].btn.classList.contains("active"), "Auto actif par défaut");
    assert.equal(rows[1].mode, "free");
    assert.equal(rows[1].name, "Modèle", "famille à un seul mode : libellé Modèle");
    assert.equal(rows[1].rule, "Free Models Router", "règle = label du modèle effectif");
  });

  test("survol N4 : Auto (fallback) + Flash/Standard", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-n4");
    const rows = optionRows(menu);
    assert.equal(rows.length, 3);
    assert.equal(rows[0].name, "Auto (fallback)");
    assert.equal(rows[0].mode, "auto");
    assert.equal(rows[1].name, "Flash");
    assert.equal(rows[1].rule, "Fallback · 2 modèles");
    assert.equal(rows[2].name, "Standard");
    assert.equal(rows[2].rule, "DeepSeek V4");
    assert.ok(!menu.textContent.includes("offres gratuites"), "jamais de règle technique");
  });

  test("survol SamGen : routes locales, pas d'Auto", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samgen");
    const rows = optionRows(menu);
    assert.equal(rows.length, 2);
    assert.ok(!rows.some((r) => r.mode === "auto"), "pas d'Auto pour SamGen");
    assert.deepEqual(
      rows.map((r) => r.name),
      ["Nano (llama.cpp)", "N4 (Ollama)"],
      "labels SamGen inchangés",
    );
  });

  test("changer de famille remplace le contenu du sous-menu", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-nano");
    assert.equal(optionRows(menu).length, 2);
    hoverFamily("samagent-n4");
    assert.equal(optionRows(menu).length, 3, "contenu remplacé");
    const openRows = document.querySelectorAll(".plus-model-family-row.open");
    assert.equal(openRows.length, 1);
    assert.equal(openRows[0].dataset.family, "samagent-n4");
  });

  test("clic sur une option du sous-menu : sélection appliquée, sous-menu fermé", async () => {
    await setupDOM();
    const { menu } = hoverFamily("samagent-n4");
    const std = optionRows(menu).find((r) => r.mode === "standard").btn;
    std.dispatchEvent(new Event("click", { bubbles: true }));
    assert.equal(document.getElementById("family-select").value, "samagent-n4");
    assert.equal(document.getElementById("mode-select").value, "standard");
    assert.equal(menu.style.display, "none", "sous-menu fermé après sélection");
    const row = document.querySelector('.plus-model-family-row[data-family="samagent-n4"]');
    assert.ok(row.classList.contains("active"), "ligne N4 marquée active");
  });

  test("clic sur Auto (fallback) : mode auto persisté + événement modèle", async () => {
    await setupDOM();
    putBodies.length = 0;
    let changedFired = 0;
    window.addEventListener("cetas:model-changed", () => changedFired++);
    const { menu } = hoverFamily("samagent-nano");
    const auto = optionRows(menu).find((r) => r.mode === "auto").btn;
    auto.dispatchEvent(new Event("click", { bubbles: true }));
    await new Promise((r) => setTimeout(r, 20)); // persistPrefs() est async
    assert.equal(document.getElementById("family-select").value, "samagent-nano");
    assert.equal(document.getElementById("mode-select").value, "auto");
    const last = putBodies[putBodies.length - 1];
    assert.ok(last, "un PUT de préférences a été émis");
    assert.equal(last.family, "samagent-nano");
    assert.equal(last.mode, "auto");
    assert.equal(changedFired, 1, "cetas:model-changed émis (le hint du composer se rafraîchit)");
  });

  test("mouseleave programme la fermeture, mouseenter du sous-menu l'annule", async () => {
    await setupDOM();
    const { row, menu } = hoverFamily("samagent-nano");
    row.dispatchEvent(new Event("mouseleave", { bubbles: false }));
    menu.dispatchEvent(new Event("mouseenter", { bubbles: false }));
    await new Promise((r) => setTimeout(r, 220));
    assert.equal(menu.style.display, "block", "resté ouvert : la souris est sur le sous-menu");
    menu.dispatchEvent(new Event("mouseleave", { bubbles: false }));
    await new Promise((r) => setTimeout(r, 220));
    assert.equal(menu.style.display, "none", "fermé après mouseleave du sous-menu");
  });

  test("plusModelSubmenuContains / closePlusModelSubmenu (intégration chat.js)", async () => {
    await setupDOM();
    hoverFamily("samagent-nano");
    const menu = document.getElementById("plus-model-submenu");
    assert.ok(mod.plusModelSubmenuContains(menu.querySelector(".plus-model-option")));
    assert.ok(!mod.plusModelSubmenuContains(document.body));
    mod.closePlusModelSubmenu();
    assert.equal(menu.style.display, "none");
  });

  test("CSS : classes du menu en cascade définies dans mx-menus.css", () => {
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
