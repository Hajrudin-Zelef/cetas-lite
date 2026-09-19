import test from "node:test";
import assert from "node:assert/strict";

// --- Faux DOM minimal (même base que opencode-style.test.mjs) ---
function makeEl(tag) {
  const kids = [];
  const e = {
    tagName: String(tag || "").toUpperCase(),
    children: kids,
    style: {},
    dataset: {},
    className: "",
    textContent: "",
    _innerHTML: "",
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
  // innerHTML = "" vide les enfants, comme dans un vrai navigateur.
  Object.defineProperty(e, "innerHTML", {
    get() { return e._innerHTML; },
    set(v) {
      e._innerHTML = String(v);
      if (v === "") { kids.length = 0; }
    },
    configurable: true,
  });
  return e;
}
const store = {};
let lastEvent = null;
globalThis.document = {
  createElement: (t) => makeEl(t),
  createTextNode: (t) => ({ textContent: String(t), nodeType: 3, appendData(s) { this.textContent += s; } }),
  querySelector: () => null,
};
globalThis.window = globalThis;
globalThis.addEventListener = () => {};
globalThis.removeEventListener = () => {};
globalThis.dispatchEvent = (e) => { lastEvent = e; return true; };
globalThis.localStorage = {
  getItem: (k) => (k in store ? store[k] : null),
  setItem: (k, v) => { store[k] = String(v); },
  removeItem: (k) => { delete store[k]; },
};
globalThis.requestAnimationFrame = (fn) => 0;
globalThis.cancelAnimationFrame = () => {};

const styleMod = await import("../agentic-style.js");
const { ThreadView } = await import("../thread-view.js");

function resetStore() {
  for (const k of Object.keys(store)) delete store[k];
  lastEvent = null;
}

function makeView(agentic) {
  const log = makeEl();
  const badge = makeEl();
  const v = new ThreadView({
    log,
    statsBadge: badge,
    streamURL: () => "",
    sendURL: "",
    stopURL: "",
    agentic: agentic === true,
    reasonHooks: { append() {}, finish() {}, reset() {} },
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

function textOf(n) {
  let s = n.textContent || "";
  for (const k of n.children || []) s += textOf(k);
  return s;
}

// ---------- Module agentic-style.js ----------
test("défaut : harness", () => {
  resetStore();
  assert.equal(styleMod.getAgenticStyle(), "harness");
});

test("setAgenticStyle : persiste + émet l'événement", () => {
  resetStore();
  const v = styleMod.setAgenticStyle("opencode");
  assert.equal(v, "opencode");
  assert.equal(store["mx.agentic.style"], "opencode");
  assert.equal(styleMod.getAgenticStyle(), "opencode");
  assert.ok(lastEvent, "événement émis");
  assert.equal(lastEvent.type, styleMod.AGENTIC_STYLE_EVENT);
  assert.equal(lastEvent.detail.style, "opencode");
});

test("setAgenticStyle : valeur invalide -> harness", () => {
  resetStore();
  assert.equal(styleMod.setAgenticStyle("nimporte"), "harness");
  assert.equal(styleMod.getAgenticStyle(), "harness");
});

// ---------- Vue Agents + Harness : rendu 18/09 ----------
test("agentic+harness : ligne '✨ Tool call · Nom · détail'", () => {
  resetStore();
  styleMod.setAgenticStyle("harness");
  const { v, log } = makeView(true);
  v.handleEvent({ seq: 1, tool: { name: "Read", phase: "start", args: { file_path: "a.txt" } } });
  const det = log.children.find((c) => (c.className || "").includes("harness-tool"));
  assert.ok(det, "details.harness-tool créé");
  const txt = textOf(det);
  assert.ok(txt.includes("Tool call"), "libellé 'Tool call' restauré");
  assert.ok(txt.includes("Read"), "nom de l'outil");
  assert.ok(txt.includes("a.txt"), "hint");
});

test("agentic+harness : pas de ligne thought, pas de thinking", () => {
  resetStore();
  styleMod.setAgenticStyle("harness");
  const { v, log } = makeView(true);
  v.handleEvent({ seq: 1, user: "lis" });
  v.lastEventTs = Date.now() - 1200;
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "start", args: { file_path: "a.txt" } } });
  assert.equal(findByClass(log, "thought-line").length, 0, "pas de thought");
  v.handleEvent({ seq: 3, reasoning_content: "je réfléchis" });
  assert.equal(findByClass(log, "thinking-indicator").length, 0, "pas de thinking");
});

test("agentic+harness : résultat complet, diff replié", () => {
  resetStore();
  styleMod.setAgenticStyle("harness");
  const { v, log } = makeView(true);
  const big = Array.from({ length: 25 }, (_, i) => "ligne " + (i + 1)).join("\n");
  v.handleEvent({ seq: 1, tool: { name: "Read", phase: "start", args: { file_path: "g.txt" } } });
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "end", args: { file_path: "g.txt" }, result: big } });
  assert.equal(findByClass(log, "tool-result-wrap").length, 0, "pas de repli 10 lignes");
  assert.equal(findByClass(log, "tool-expand-btn").length, 0, "pas de bouton expand");
  const pres = findByClass(log, "tool-result");
  assert.equal(pres.length, 1, "un seul <pre> complet");
  assert.ok(textOf(pres[0]).includes("ligne 25"), "résultat intégral");
  v.handleEvent({ seq: 3, tool: { name: "Edit", phase: "start", args: { file_path: "f.txt" } } });
  v.handleEvent({ seq: 4, tool: { name: "Edit", phase: "end", args: { file_path: "f.txt" }, diff: [{ kind: "+", text: "x" }] } });
  const det = log.children.filter((c) => (c.className || "").includes("harness-tool"))[1];
  assert.ok(det && det.open !== true, "details replié par défaut");
});

// ---------- Vue Agents + OpenCode : TUI fidèle ----------
test("agentic+opencode : en-tête 'Nom: action…' puis 'Nom: paramètres'", () => {
  resetStore();
  styleMod.setAgenticStyle("opencode");
  const { v, log } = makeView(true);
  v.handleEvent({ seq: 1, tool: { name: "Bash", phase: "start", args: { command: "ls -la" } } });
  const box = findByClass(log, "oc-tool")[0];
  assert.ok(box, "bloc oc-tool créé");
  const head = findByClass(log, "oc-tool-head")[0];
  assert.ok(head, "en-tête présent");
  assert.ok(textOf(head).includes("Bash: "), "nom + deux-points");
  assert.ok(textOf(head).includes("Building command..."), "action TUI pendant l'exécution");
  v.handleEvent({ seq: 2, tool: { name: "Bash", phase: "end", args: { command: "ls -la" }, result: "total 0" } });
  assert.ok(textOf(head).includes("ls -la"), "paramètres après exécution");
  assert.ok(!textOf(head).includes("Building command..."), "action remplacée");
});

test("agentic+opencode : thought + thinking affichés", () => {
  resetStore();
  styleMod.setAgenticStyle("opencode");
  const { v, log } = makeView(true);
  v.handleEvent({ seq: 1, user: "go" });
  v.lastEventTs = Date.now() - 1200;
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "start", args: { file_path: "a.txt" } } });
  assert.equal(findByClass(log, "thought-line").length, 1, "ligne thought");
  v.handleEvent({ seq: 3, reasoning_content: "je réfléchis" });
  assert.equal(findByClass(log, "thinking-indicator").length, 1, "thinking affiché");
});

test("agentic+opencode : résultat borné à 10 lignes + erreur en rouge", () => {
  resetStore();
  styleMod.setAgenticStyle("opencode");
  const { v, log } = makeView(true);
  const big = Array.from({ length: 25 }, (_, i) => "ligne " + (i + 1)).join("\n");
  v.handleEvent({ seq: 1, tool: { name: "Read", phase: "start", args: { file_path: "g.txt" } } });
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "end", args: { file_path: "g.txt" }, result: big } });
  assert.equal(findByClass(log, "tool-result-wrap").length, 1, "repli 10 lignes");
  v.handleEvent({ seq: 3, tool: { name: "Bash", phase: "start", args: { command: "x" } } });
  v.handleEvent({ seq: 4, tool: { name: "Bash", phase: "end", args: { command: "x" }, result: "[erreur] boom" } });
  const err = findByClass(log, "oc-tool-error");
  assert.equal(err.length, 1, "bloc erreur");
  assert.ok(err[0].textContent.includes("boom"), "contenu d'erreur");
});

test("agentic+opencode : diff visible + pied ' modèle (durée)' + 'Error:'", () => {
  resetStore();
  styleMod.setAgenticStyle("opencode");
  const { v, log } = makeView(true);
  v.handleEvent({ seq: 1, tool: { name: "Edit", phase: "start", args: { file_path: "f.txt" } } });
  v.handleEvent({ seq: 2, tool: { name: "Edit", phase: "end", args: { file_path: "f.txt" }, diff: [{ kind: "+", text: "n" }] } });
  assert.equal(findByClass(log, "tool-diff").length, 1, "diff rendu");
  assert.equal(findByClass(log, "harness-tool").length, 0, "pas de details harness");
  const box = makeEl("div");
  v.turnStats = {};
  v.turnRoute = { label: "MiMo-V2.5" };
  v.turnElapsedMs = 1700;
  v.addTurnStats(box);
  const stats = findByClass(box, "oc-turn-stats");
  assert.equal(stats.length, 1, "pied TUI");
  assert.ok(stats[0].textContent.includes("MiMo-V2.5"), "modèle");
  assert.match(stats[0].textContent, /1\.7s/, "durée");
  v.addError("raté");
  const errs = findByClass(log, "message-error");
  assert.ok(errs.length >= 1, "bulle erreur");
  assert.ok(textOf(errs[errs.length - 1]).startsWith("Error: "), "préfixe TUI");
});

// ---------- Panneau Agentic ----------
test("loadAgenticPanel : seg Harness/OpenCode, clic = choix immédiat", async () => {
  resetStore();
  const els = {};
  const bodyEl = makeEl("div");
  bodyEl.id = "agentic-body";
  const saveEl = makeEl("span");
  saveEl.id = "agentic-save-state";
  els["agentic-body"] = bodyEl;
  els["agentic-save-state"] = saveEl;
  globalThis.document.getElementById = (id) => els[id] || null;

  const { loadAgenticPanel } = await import("../agentic-panel.js");
  loadAgenticPanel();
  assert.equal(bodyEl.dataset.loaded, "1", "marqué chargé");
  const findBtns = () => {
    const out = [];
    const walk = (n) => {
      if ((n.className || "").split(" ").includes("ms-seg-btn")) out.push(n);
      for (const k of n.children || []) walk(k);
    };
    walk(bodyEl);
    return out;
  };
  const btns = findBtns();
  assert.equal(btns.length, 2, "deux choix");
  assert.deepEqual(
    btns.map((b) => b.dataset.style).sort(),
    ["harness", "opencode"]
  );
  // Clic OpenCode : sauvegarde immédiate + événement.
  btns.find((b) => b.dataset.style === "opencode").click();
  assert.equal(store["mx.agentic.style"], "opencode", "choix persisté");
  assert.equal(saveEl.textContent, "Enregistré ✓", "état de sauvegarde");
  assert.ok(lastEvent && lastEvent.detail.style === "opencode", "événement émis");
  // Second appel : pas de re-rendu (dataset.loaded).
  const n = bodyEl.children.length;
  loadAgenticPanel();
  assert.equal(bodyEl.children.length, n, "pas de doublon");
});

// ---------- Chat général : comportement actuel inchangé ----------
test("non-agentic : rendu actuel préservé (pas de 'Tool call', thought, statut)", () => {
  resetStore();
  styleMod.setAgenticStyle("opencode"); // même avec OpenCode choisi…
  const { v, log } = makeView(false); // …le chat général ne change pas
  v.handleEvent({ seq: 1, user: "go" });
  v.lastEventTs = Date.now() - 1200;
  v.handleEvent({ seq: 2, tool: { name: "Read", phase: "start", args: { file_path: "a.txt" } } });
  const det = log.children.find((c) => (c.className || "").includes("harness-tool"));
  assert.ok(det, "details créé");
  assert.ok(!textOf(det).includes("Tool call"), "toujours sans 'Tool call'");
  assert.equal(findByClass(log, "thought-line").length, 1, "thought conservé");
  assert.equal(findByClass(log, "oc-tool").length, 0, "pas de bloc opencode");
  v.handleEvent({ seq: 3, tool: { name: "Bash", phase: "start", args: { command: "x" } } });
  assert.equal(findByClass(log, "agent-status").length, 1, "statut 'Writing command' conservé");
});
