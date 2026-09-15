import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { createRequire } from "node:module";

// ---------------------------------------------------------------------------
// Inter auto-hébergée + highlight.js câblé.
// ---------------------------------------------------------------------------

const repoRoot = new URL("../../../", import.meta.url);
const read = (rel) => fs.readFileSync(new URL(rel, repoRoot), "utf8");

// 1. Police Inter : fichier présent, taille plausible.
test("fonts : inter-latin.woff2 présent et non vide", () => {
  const url = new URL("web/fonts/inter/inter-latin.woff2", repoRoot);
  const st = fs.statSync(url);
  assert.ok(st.size > 10_000, "woff2 non trivial");
  assert.ok(st.size < 500_000, "woff2 variable raisonnable (< 500 Ko)");
});

// 2. mx-fonts.css déclare le @font-face Inter -> woff2 local.
test("css : mx-fonts.css déclare Inter en local", () => {
  const css = read("web/css/features/mx-fonts.css");
  assert.match(css, /@font-face/);
  assert.match(css, /font-family:\s*"Inter"/);
  assert.match(css, /url\("\/fonts\/inter\/inter-latin\.woff2"\)/);
  assert.doesNotMatch(css, /fonts\.googleapis|fonts\.gstatic/, "aucune dépendance externe");
});

// 3. index.html charge mx-fonts.css AVANT marex-agents.css (qui demande Inter).
test("html : mx-fonts.css chargé avant marex-agents.css", () => {
  const html = read("web/index.html");
  const iFonts = html.indexOf("mx-fonts.css");
  const iAgents = html.indexOf("marex-agents.css");
  assert.ok(iFonts > -1, "mx-fonts.css lié");
  assert.ok(iAgents > -1, "marex-agents.css lié");
  assert.ok(iFonts < iAgents, "police déclarée avant usage");
});

// 4. go:embed inclut fonts (sinon 404 sur /fonts/...).
test("go : embed.go embarque le dossier fonts", () => {
  const go = read("web/embed.go");
  const m = go.match(/\/\/go:embed\s+([^\n]+)/);
  assert.ok(m, "directive go:embed présente");
  assert.ok(m[1].split(/\s+/).includes("fonts"), "fonts dans go:embed");
});

// 5. highlight.esm.min.js est un module ESM valide qui expose highlightElement.
test("highlight : le bundle ESM s'importe et expose highlightElement", async () => {
  const mod = await import("../vendor/highlight.esm.min.js");
  const hljs = mod.default || mod.HighlightJS;
  assert.ok(hljs, "export hljs présent");
  assert.equal(typeof hljs.highlightElement, "function");
});

// 6. markdown.js câble highlight (import paresseux + garde anti-doublon).
test("markdown : câblage highlight présent", () => {
  const src = read("web/js/markdown.js");
  assert.match(src, /highlight\.esm\.min\.js/);
  assert.match(src, /highlightElement/);
  assert.match(src, /data-hl/, "garde anti re-coloration");
});

// 7. Rendu d'un bloc de code : pas d'exception même si hljs indisponible
//    (en node, l'import absolu /js/... échoue -> dégradation gracieuse).
test("markdown : renderInto avec code ne jette pas sans hljs", async () => {
  const require = createRequire(import.meta.url);
  let JSDOM;
  try {
    ({ JSDOM } = require("jsdom"));
  } catch {
    ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
  }
  const dom = new JSDOM("<!doctype html><html><body></body></html>", { url: "http://localhost/" });
  globalThis.window = dom.window;
  globalThis.document = dom.window.document;
  const { renderInto } = await import("../markdown.js");
  const el = document.createElement("div");
  // Sans marked global, renderMarkdown fait un <p> d'échappement : on pose
  // directement un bloc pre/code pour exercer highlightNewCode.
  el.innerHTML = '<pre><code class="language-js">const a = 1;</code></pre>';
  const { highlightNewCode } = await import("../markdown.js");
  highlightNewCode(el);
  await new Promise((r) => setTimeout(r, 50));
  assert.equal(el.querySelector("code").textContent, "const a = 1;", "code intact sans hljs");
});
