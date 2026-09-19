import test from "node:test";
import assert from "node:assert/strict";

// --- Faux DOM minimal pour ThreadView (thought, thinking, expand, diff) ---
function makeEl(tag) {
  const kids = [];
  const e = {
    tagName: String(tag || "").toUpperCase(),
    children: kids,
    style: {},
    dataset: {},
    className: "",
    textContent: "",
    innerHTML: "",
    hidden: false,
    open: false,
    classList: {
      add() {}, remove() {}, toggle() {},
      contains() { return false; },
    },
    appendChild(c) { kids.push(c); c.parentNode = e; return c; },
    append(...a) { for (const c of a) c.parentNode = e; kids.push(...a); return a[0]; },
    prepend(...a) { for (const c of a) c.parentNode = e; kids.unshift(...a); },
    insertBefore(c, ref) {
      c.parentNode = e;
      const i = ref ? kids.indexOf(ref) : -1;
      if (i >= 0) kids.splice(i, 0, c); else kids.push(c);
      return c;
    },
    remove() { e._removed = true; },
    before() {},
    querySelector() { return null; },
    querySelectorAll() { return []; },
    closest(sel) {
      let n = e.parentNode;
      const want = String(sel || "").toUpperCase();
      while (n) {
        if (n.tagName === want) return n;
        n = n.parentNode;
      }
      return null;
    },
    setAttribute() {},
    getAttribute() { return null; },
    addEventListener(t, fn) { e._listeners = e._listeners || {}; e._listeners[t] = fn; },
    removeEventListener() {},
    scrollIntoView() {},
    focus() {},
    click() { if (e._listeners && e._listeners.click) e._listeners.click(); },
  };
  return e;
}
globalThis.document = {
  createElement: (t) => makeEl(t),
  createTextNode: (t) => ({ textContent: String(t), nodeType: 3, appendData(s) { this.textContent += s; } }),
  querySelector: () => null,
};
globalThis.window = globalThis;
globalThis.addEventListener = () => {};
globalThis.removeEventListener = () => {};
globalThis.dispatchEvent = () => true;
globalThis.localStorage = { getItem: () => null, setItem: () => {}, removeItem: () => {} };
globalThis.requestAnimationFrame = (fn) => 0;
globalThis.cancelAnimationFrame = () => {};

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
    // Comme la vue Agents : le raisonnement part au panneau latéral.
    reasonHooks: {
      append() {},
      finish() {},
      reset() {},
    },
  });
  return { v, badge, log };
}

function findByClass(log, cls) {
  const out = [];
  const walk = (n) => {
    if ((n.className || "").split(" ").includes(cls)) out.push(n);
    for (const k of n.children || []) walk(k);
  };
  walk(log);
  return out;
}

test("+ Thought: ligne insérée avant l'outil quand le délai >= 300ms", () => {
  const { v, log } = makeView();
  v.handleEvent({ seq: 1, user: "lis le fichier" });
  // Simule 1,2 s de réflexion avant l'appel d'outil.
  v.lastEventTs = Date.now() - 1200;
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "start", args: { file_path: "a.txt" } } });
  const thoughts = findByClass(log, "thought-line");
  assert.equal(thoughts.length, 1, "une ligne thought");
  const txt = thoughts[0].children.map((c) => c.textContent).join("");
  assert.ok(txt.includes("+ Thought:"), "préfixe");
  assert.match(txt, /1\.2s/, "durée formatée");
});

test("+ Thought: pas de ligne si enchaînement immédiat (< 300ms)", () => {
  const { v, log } = makeView();
  v.handleEvent({ seq: 1, user: "go" });
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "start", args: {} } });
  assert.equal(findByClass(log, "thought-line").length, 0, "pas de ligne thought");
});

test("thinking + spinner : affiché au raisonnement, masqué au contenu", () => {
  const { v, log } = makeView();
  v.handleEvent({ seq: 1, user: "réfléchis" });
  v.handleEvent({ seq: 2, reasoning_content: "je réfléchis…" });
  const ind = findByClass(log, "thinking-indicator");
  assert.equal(ind.length, 1, "indicateur thinking affiché");
  assert.ok(findByClass(log, "thinking-spinner").length >= 1, "spinner présent");
  v.handleEvent({ seq: 3, content: "voici" });
  assert.equal(findByClass(log, "thinking-indicator").filter((e) => !e._removed).length, 0, "indicateur masqué");
});

test("lecture de fichier : 10 lignes + bouton Expand/Collapse", () => {
  const { v, log } = makeView();
  const big = Array.from({ length: 25 }, (_, i) => "ligne " + (i + 1)).join("\n");
  v.handleEvent({ seq: 1, tool: { name: "Read", phase: "start", args: { file_path: "gros.txt" } } });
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "end", args: { file_path: "gros.txt" }, result: big } });
  const wraps = findByClass(log, "tool-result-wrap");
  assert.equal(wraps.length, 1, "conteneur repliable");
  const pres = findByClass(log, "tool-result");
  assert.equal(pres.length, 2, "deux <pre> (court + complet)");
  assert.ok(!pres[0].hidden && pres[1].hidden, "court visible, complet masqué");
  assert.ok(pres[0].textContent.split("\n").length <= 10, "10 lignes max");
  const btn = findByClass(log, "tool-expand-btn")[0];
  assert.ok(btn, "bouton présent");
  assert.ok(btn.textContent.includes("25 lignes"), "compte de lignes");
  btn.click();
  assert.ok(pres[0].hidden && !pres[1].hidden, "expand : complet visible");
  assert.ok(btn.textContent.includes("Collapse"), "libellé Collapse");
  btn.click();
  assert.ok(!pres[0].hidden && pres[1].hidden, "collapse : retour");
});

test("petite lecture : pas de bouton Expand", () => {
  const { v, log } = makeView();
  v.handleEvent({ seq: 1, tool: { name: "Cat", phase: "start", args: { file_path: "petit.txt" } } });
  v.handleEvent({ seq: 2, tool: { name: "Cat", phase: "end", args: { file_path: "petit.txt" }, result: "a\nb\nc" } });
  assert.equal(findByClass(log, "tool-expand-btn").length, 0, "pas de bouton");
});

test("diff d'édition : déplié par défaut, lignes vert/rouge", () => {
  const { v, log } = makeView();
  const diff = [
    { kind: " ", text: "contexte" },
    { kind: "-", text: "ancien" },
    { kind: "+", text: "nouveau" },
  ];
  v.handleEvent({ seq: 1, tool: { name: "Edit", phase: "start", args: { file_path: "f.txt" } } });
  v.handleEvent({ seq: 2, tool: { name: "Edit", phase: "end", args: { file_path: "f.txt" }, diff } });
  const diffs = findByClass(log, "tool-diff");
  assert.equal(diffs.length, 1, "diff rendu");
  assert.equal(findByClass(log, "diff-add").length, 1, "ligne verte");
  assert.equal(findByClass(log, "diff-del").length, 1, "ligne rouge");
  // Le details parent doit être ouvert (modifications non masquées).
  const det = log.children.find((c) => (c.className || "").includes("harness-tool"));
  assert.ok(det && det.open === true, "details déplié");
});
