import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";

import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT = join(__dirname, "..", "..");

// La vue Agents suit le thème (clair/sombre/hard_dark) et la palette de
// l'app à 100 % : .light + data-accent synchronisés depuis
// html[data-theme] / html[data-palette] ; les couleurs sombres en dur du
// CSS généré sont corrigées pour fond clair dans la couche d'override.

await test("thème agents : le bloc clair du CSS généré est intact", () => {
  const gen = fs.readFileSync(`${ROOT}/css/features/marex-agents.css`, "utf8");
  assert.ok(gen.includes("#marex-view.light{"), "bloc #marex-view.light présent");
  assert.ok(gen.includes("--bg-app:#f3f4f7;"), "variables claires d'origine intactes");
});

await test("thème agents : mx-module.css complète le thème clair", () => {
  const css = fs.readFileSync(`${ROOT}/css/features/mx-module.css`, "utf8");
  const light = (sel) => {
    const m = css.match(
      new RegExp("#marex-view\\.light " + sel + "\\{([^}]*)\\}")
    );
    assert.ok(m, `règle #marex-view.light ${sel} présente`);
    return m[1];
  };
  // Titres de groupes : bleu soutenu lisible sur fond blanc
  // (le #6EA8FE du sombre serait illisible).
  assert.ok(light("\\.sb-group-label").includes("color:#1a56db"), "sb-group-label bleu soutenu");
  // Contrôles natifs en mode clair.
  assert.ok(css.includes("#marex-view.light{ color-scheme:light; }"), "color-scheme:light");
  // Aucune pastille sombre restante : stop, erreurs, todo, diff, menus, textarea.
  assert.ok(light("\\.stop-btn").includes("#f7dede"), "stop-btn clair");
  assert.ok(
    light("\\.mx-chat-log \\.message-error").includes("#fdecec"),
    "message-error clair"
  );
  assert.ok(css.includes("#marex-view.light .md pre code"), "code pre clair");
  assert.ok(light("\\.cdrop-item\\.selected").includes("#e8ebf0"), "cdrop selected clair");
  assert.ok(
    light("\\.modal-instructions-textarea").includes("#f4f5f7"),
    "textarea instructions claire"
  );
  // TodoWrite jaune pur -> ambre sombre lisible.
  assert.ok(css.includes("#8a6d00 !important"), "chat-todo ambre sombre");
  // Diff +/- : verts/rouges lisibles sur fond clair.
  assert.ok(light("\\.diff-add").includes("#1a7f37"), "diff-add lisible");
  assert.ok(light("\\.diff-del").includes("#d1242f"), "diff-del lisible");
});

await test("thème agents : agents.js synchronise thème + palette", () => {
  const src = fs.readFileSync(`${ROOT}/js/agents.js`, "utf8");
  // Classe .light pilotée par html[data-theme].
  assert.ok(src.includes('classList.toggle("light"'), '.toggle("light") présent');
  assert.ok(src.includes("dataset.theme"), "lecture de html[data-theme]");
  // Palette de l'app -> data-accent de la vue (couleurs respectées).
  for (const p of ["bleu", "violet", "vert", "bleu_ocean", "jaune_or", "rouge"]) {
    assert.ok(src.includes(p + ":"), `palette ${p} mappée`);
  }
  // Répercussion en direct (changement depuis Configuration).
  assert.ok(
    /new MutationObserver\([\s\S]*data-theme/.test(src),
    "MutationObserver sur data-theme/data-palette"
  );
});
