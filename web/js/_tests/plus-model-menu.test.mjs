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
    modes: [{ id: "free", mode: "free", label: "Free", rule: "offres gratuites", agent: false }],
  },
  {
    id: "code",
    label: "Code",
    modes: [{ id: "flash", mode: "flash", label: "Flash", rule: "deepseek flash", agent: true }],
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

const { initModels } = await import("../model-select.js");
await initModels();

await test("menu + : structure premium de la liste des modèles", () => {
  const list = document.getElementById("plus-model-list");
  const groups = list.querySelectorAll(".plus-model-group");
  assert.equal(groups.length, 2);
  assert.equal(groups[0].querySelector(".plus-model-group-label").textContent, "SamAgent Nano");
  const opts = list.querySelectorAll(".plus-model-option");
  assert.equal(opts.length, 2);
  // Nom + règle en sous-titre.
  const main = opts[0].querySelector(".plus-model-option-main");
  assert.ok(main, ".plus-model-option-main présent");
  assert.equal(main.querySelector(".plus-model-option-name").textContent, "Free");
  assert.equal(main.querySelector(".plus-model-option-rule").textContent, "offres gratuites");
  // Pas de badge Agent : le chat général est pur chat, l'agent vit dans le module Agents.
  assert.equal(opts[0].querySelector(".plus-model-option-badge"), null);
  assert.equal(opts[1].querySelector(".plus-model-option-badge"), null);
  // Le mode courant est marqué actif.
  assert.ok(
    opts[0].classList.contains("active") || opts[1].classList.contains("active"),
    "un mode est .active"
  );
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
