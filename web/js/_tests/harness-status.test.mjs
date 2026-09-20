import test from "node:test";
import assert from "node:assert/strict";

// --- Faux DOM minimal pour ThreadView (chemins user/tool/stats/turn_done) ---
function makeEl() {
  const kids = [];
  const e = {
    children: kids,
    style: {},
    dataset: {},
    className: "",
    textContent: "",
    innerHTML: "",
    hidden: false,
    classList: {
      add() {}, remove() {}, toggle() {},
      contains() { return false; },
    },
    appendChild(c) { kids.push(c); return c; },
    append(...a) { kids.push(...a); return c0(a); },
    prepend(...a) { kids.unshift(...a); },
    remove() {},
    before() {},
    querySelector() { return null; },
    querySelectorAll() { return []; },
    closest() { return null; },
    setAttribute() {},
    getAttribute() { return null; },
    addEventListener() {},
    removeEventListener() {},
    scrollIntoView() {},
    focus() {},
  };
  function c0(a) { return a[0]; }
  return e;
}
globalThis.document = {
  createElement: () => makeEl(),
  createTextNode: (t) => ({ textContent: String(t), nodeType: 3 }),
  querySelector: () => null,
};
globalThis.window = globalThis;
globalThis.addEventListener = () => {};
globalThis.removeEventListener = () => {};
globalThis.dispatchEvent = () => true;
globalThis.localStorage = { getItem: () => null, setItem: () => {}, removeItem: () => {} };
globalThis.requestAnimationFrame = (fn) => 0;
globalThis.cancelAnimationFrame = () => {};
// setInterval/clearInterval existent en node ; le timer du wait-label tourne.

const { ThreadView } = await import("../thread-view.js");

function makeView() {
  const log = makeEl();
  const badge = makeEl();
  const v = new ThreadView({
    log,
    statsBadge: badge,
    streamURL: () => "",
    sendURL: "",
    stopURL: "",
    // Vue Agents (phase 1 refonte) : les lignes d'outils sobres (glyphes)
    // et le 5+5 ne s'appliquent qu'ici ; le chat général garde son rendu.
    agentic: true,
  });
  return { v, badge, log };
}

test("badge façon Harness : tours, outils, tokens cumulés", () => {
  const { v, badge } = makeView();
  // Tour 1 : 2 outils, 2 réponses provider (5k puis 14k = cumul tour 14k).
  v.handleEvent({ seq: 1, user: "salut" });
  v.handleEvent({ seq: 2, tool: { name: "Tree", phase: "start", args: {} } });
  v.handleEvent({ seq: 3, tool: { name: "Tree", phase: "end", args: {}, result: "ok" } });
  v.handleEvent({ seq: 4, tool: { name: "Read", phase: "start", args: { file_path: "a.md" } } });
  v.handleEvent({ seq: 5, stats: { prompt_tokens: 5000, completion_tokens: 100 } });
  v.handleEvent({ seq: 6, stats: { prompt_tokens: 14000, completion_tokens: 168 } });
  v.handleEvent({ seq: 7, turn_done: { elapsed_ms: 3200 } });
  assert.equal(v.convTurns, 1);
  assert.equal(v.convTools, 2);
  assert.equal(v.convIn, 14000);
  assert.equal(v.convOut, 168);
  assert.match(badge.textContent, /^1 tour · 2 outils · ↑14k ↓168$/);

  // Tour 2 : 1 outil, tokens du tour 2 s'ajoutent au cumul.
  v.handleEvent({ seq: 8, user: "et ça ?" });
  v.handleEvent({ seq: 9, tool: { name: "Bash", phase: "start", args: { command: "ls" } } });
  v.handleEvent({ seq: 10, stats: { prompt_tokens: 14500, completion_tokens: 200 } });
  v.handleEvent({ seq: 11, turn_done: { elapsed_ms: 1500 } });
  assert.equal(v.convTurns, 2);
  assert.equal(v.convTools, 3);
  assert.equal(v.convIn, 28500);
  assert.equal(v.convOut, 368);
  assert.match(badge.textContent, /^2 tours · 3 outils · ↑28,5k ↓368$/);
});

test("ligne d'outil : summary glyphe sobre · nom · hint (sans 'Tool call', sans ✨)", () => {
  const { v, log } = makeView();
  v.handleEvent({ seq: 1, tool: { name: "Read", phase: "start", args: { file_path: "NEVA PVE/README.md" } } });
  const det = log.children.find((c) => (c.className || "").includes("harness-tool"));
  assert.ok(det, "details.harness-tool créé");
  assert.ok((det.className || "").includes("running"), "classe running pendant l'appel");
  const sum = det.children[0];
  const txt = sum.children.map((c) => c.textContent).join("");
  assert.ok(!txt.includes("✨"), "plus d'étincelle (phase 1 refonte)");
  assert.ok(txt.includes("→"), "glyphe sobre par outil");
  assert.ok(!txt.includes("Tool call"), "plus de libellé 'Tool call'");
  assert.ok(txt.includes("Read"), "nom de l'outil");
  assert.ok(txt.includes("NEVA PVE/README.md"), "hint");
  // Fin de l'appel : la classe running disparaît.
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "end", args: { file_path: "NEVA PVE/README.md" }, result: "x" } });
});
