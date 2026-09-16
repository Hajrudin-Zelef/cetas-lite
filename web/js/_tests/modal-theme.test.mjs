import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";

import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT = join(__dirname, "..", "..");

await test("modale : le CSS généré n'est pas modifié", () => {
  const gen = fs.readFileSync(`${ROOT}/css/features/marex-agents.css`, "utf8");
  // Le bloc d'origine force toujours le sombre en dur (c'est l'override local qui corrige).
  assert.ok(gen.includes("--bg-panel:#0d1622;"), "bloc sombre d'origine intact");
  assert.ok(gen.includes("color-scheme:dark;"), "color-scheme sombre d'origine intact");
});

await test("modale : mx-module.css surcharge vers le thème + la palette", () => {
  const css = fs.readFileSync(`${ROOT}/css/features/mx-module.css`, "utf8");
  // 1. L'accent suit la palette (héritage depuis html[data-palette]).
  assert.ok(
    /\.mx-modal-overlay\s*\{[^}]*--accent:\s*inherit/.test(css),
    "--accent:inherit sur .mx-modal-overlay"
  );
  // 2. Thème clair : variables claires + color-scheme light.
  const m = css.match(/html\[data-theme="clair"\]\s*\.mx-modal-overlay\s*\{([^}]*)\}/);
  assert.ok(m, "règle html[data-theme=clair] .mx-modal-overlay présente");
  const body = m[1];
  for (const decl of [
    "--bg-panel:#ffffff",
    "--text-primary:#1a1d21",
    "color-scheme:light",
  ]) {
    assert.ok(body.includes(decl), `thème clair : ${decl}`);
  }
  // 3. L'override est chargé après le généré (ordre des <link> dans index.html).
  const html = fs.readFileSync(`${ROOT}/index.html`, "utf8");
  assert.ok(
    html.indexOf("marex-agents.css") < html.indexOf("mx-module.css"),
    "mx-module.css chargé après marex-agents.css"
  );
});

await test("modale : aucun style inline sombre dans le HTML du modal", async () => {
  const src = fs.readFileSync(`${ROOT}/js/projects.js`, "utf8");
  const start = src.indexOf("modal nouveau projet");
  const end = src.indexOf("export function", start + 10);
  const chunk = src.slice(start, end);
  assert.ok(!/#0d1622|#1f2024|#25262b/.test(chunk), "pas de couleur sombre en dur dans le JS");
});
