import test from "node:test";
import assert from "node:assert/strict";

// --- Faux DOM minimal (même harnais que opencode-style.test.mjs) ---
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
    closest() { return null; },
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
  const v = new ThreadView({
    log,
    streamURL: () => "",
    sendURL: "",
    stopURL: "",
    reasonHooks: { append() {}, finish() {}, reset() {} },
  });
  return { v, log };
}

function findByClass(log, cls) {
  const out = [];
  const walk = (n) => {
    if (!n._removed && (n.className || "").split(" ").includes(cls)) out.push(n);
    for (const k of n.children || []) walk(k);
  };
  walk(log);
  return out;
}

function toolStart(v, seq, name, args) {
  v.handleEvent({ seq, tool: { name, phase: "start", args: args || {} } });
}
function toolEnd(v, seq, name, args) {
  v.handleEvent({ seq, tool: { name, phase: "end", args: args || {}, result: "ok" } });
}

test("Bash → statut « Writing command »", () => {
  const { v, log } = makeView();
  toolStart(v, 1, "Bash", { command: "ls" });
  const st = findByClass(log, "agent-status");
  assert.equal(st.length, 1, "une ligne de statut");
  assert.equal(st[0].textContent, "Writing command");
});

test("RunScript → « Writing command », Edit/Write → « Preparing edit »", () => {
  const { v, log } = makeView();
  toolStart(v, 1, "RunScript", { script: "x" });
  assert.equal(findByClass(log, "agent-status")[0].textContent, "Writing command");
  toolEnd(v, 2, "RunScript", { script: "x" });
  toolStart(v, 3, "Edit", { file_path: "a.txt" });
  assert.equal(findByClass(log, "agent-status")[0].textContent, "Preparing edit");
  toolEnd(v, 4, "Edit", { file_path: "a.txt" });
  toolStart(v, 5, "Write", { file_path: "b.txt" });
  assert.equal(findByClass(log, "agent-status")[0].textContent, "Preparing edit");
});

test("Read → aucun statut", () => {
  const { v, log } = makeView();
  toolStart(v, 1, "Read", { file_path: "a.txt" });
  assert.equal(findByClass(log, "agent-status").length, 0, "pas de statut pour Read");
});

test("fin d'outil → statut retiré", () => {
  const { v, log } = makeView();
  toolStart(v, 1, "Bash", { command: "ls" });
  assert.equal(findByClass(log, "agent-status").length, 1);
  toolEnd(v, 2, "Bash", { command: "ls" });
  assert.equal(findByClass(log, "agent-status").length, 0, "statut retiré");
});

test("outils parallèles : le dernier démarré donne le libellé", () => {
  const { v, log } = makeView();
  toolStart(v, 1, "Bash", { command: "a" });
  toolStart(v, 2, "Edit", { file_path: "f" });
  assert.equal(findByClass(log, "agent-status")[0].textContent, "Preparing edit");
  toolEnd(v, 3, "Edit", { file_path: "f" });
  assert.equal(findByClass(log, "agent-status")[0].textContent, "Writing command", "retour au Bash restant");
  toolEnd(v, 4, "Bash", { command: "a" });
  assert.equal(findByClass(log, "agent-status").length, 0);
});

test("turn_done → statut nettoyé", () => {
  const { v, log } = makeView();
  toolStart(v, 1, "Bash", { command: "ls" });
  assert.equal(findByClass(log, "agent-status").length, 1);
  v.handleEvent({ seq: 2, turn_done: {} });
  assert.equal(findByClass(log, "agent-status").length, 0, "nettoyé en fin de tour");
});
