import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";

// ---------------------------------------------------------------------------
// Thème highlight.js assorti au mode clair/sombre + bloc de code intégré.
// ---------------------------------------------------------------------------

const repoRoot = new URL("../../../", import.meta.url);
const read = (rel) => fs.readFileSync(new URL(rel, repoRoot), "utf8");

test("html : deux thèmes hljs liés avec ids", () => {
  const html = read("web/index.html");
  assert.match(html, /highlight-github-dark\.min\.css" id="hljs-theme-dark"/);
  assert.match(html, /highlight-github-light\.min\.css" id="hljs-theme-light"/);
});

test("css : thème clair hljs vendored, local, sans dépendance externe", () => {
  const css = read("web/js/vendor/highlight-github-light.min.css");
  assert.match(css, /\.hljs-string/);
  assert.match(css, /\.hljs-keyword/);
  assert.match(css, /\.hljs-comment/);
  assert.match(css, /\.hljs-title/);
  assert.doesNotMatch(css, /https?:\/\//, "aucune dépendance externe");
  const st = fs.statSync(new URL("web/js/vendor/highlight-github-light.min.css", repoRoot));
  assert.ok(st.size > 500, "thème non trivial");
});

test("css : fond de bloc de code intégré au design (chat principal)", () => {
  const css = read("web/css/cetas-lite.css");
  assert.match(css, /#chat-container \.message-text pre\s*\{[^}]*var\(--bg-sidebar/);
  assert.match(css, /#chat-container \.message-text pre code\s*\{[^}]*background:\s*transparent/);
});

test("theme-init : bascule disabled des thèmes hljs selon le mode", () => {
  // Stub DOM minimal : pas besoin de jsdom pour ce script.
  const src = read("web/js/theme-init.js");
  const cases = [
    ["clair", true, false],
    ["sombre", false, true],
    ["hard_dark", false, true],
  ];
  for (const [theme, darkDisabled, lightDisabled] of cases) {
    const links = {
      "hljs-theme-dark": { disabled: false },
      "hljs-theme-light": { disabled: false },
    };
    const store = { "cetas-lite-theme": theme };
    const document = {
      readyState: "complete",
      documentElement: { dataset: {} },
      body: { className: "" },
      getElementById: (id) => links[id] || null,
      addEventListener: () => {},
    };
    const localStorage = {
      getItem: (k) => (k in store ? store[k] : null),
      setItem: (k, v) => { store[k] = String(v); },
    };
    const window = {};
    new Function("document", "localStorage", "window", src)(document, localStorage, window);
    assert.equal(links["hljs-theme-dark"].disabled, darkDisabled, `${theme} : dark disabled`);
    assert.equal(links["hljs-theme-light"].disabled, lightDisabled, `${theme} : light disabled`);
    // Bascule à chaud via la fonction exposée (chemin applyTheme).
    assert.equal(typeof window.cetasHljsTheme, "function", "cetasHljsTheme exposée");
    window.cetasHljsTheme(theme === "clair" ? "sombre" : "clair");
    assert.equal(links["hljs-theme-dark"].disabled, !darkDisabled, `${theme} : dark rebasculé`);
    assert.equal(links["hljs-theme-light"].disabled, !lightDisabled, `${theme} : light rebasculé`);
  }
});

test("model-select : applyTheme rebascule le thème hljs à chaud", () => {
  const src = read("web/js/model-select.js");
  assert.match(src, /window\.cetasHljsTheme\)\s*window\.cetasHljsTheme\(t\)/);
});
