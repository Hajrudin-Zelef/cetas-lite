// Tests phase 3 refonte (2026-09-20) — diffs compacts, raisonnement mobile,
// ligne de statut Codex. Exécutables sans jsdom (faux DOM minimal).
import assert from "node:assert/strict";
import { test } from "node:test";

// --- Faux DOM ---
function makeEl(tag) {
  const kids = [];
  const el = {
    tag, children: kids, parentNode: null,
    className: "", textContent: "", hidden: false, disabled: false,
    dataset: {}, style: {}, open: false,
    _attrs: {}, _listeners: {},
    appendChild(c) { c.parentNode = el; kids.push(c); return c; },
    insertBefore(c, ref) {
      c.parentNode = el;
      const i = ref ? kids.indexOf(ref) : -1;
      if (i < 0) kids.push(c); else kids.splice(i, 0, c);
      return c;
    },
    remove() {
      if (el.parentNode) el.parentNode.children.splice(el.parentNode.children.indexOf(el), 1);
    },
    closest(sel) {
      let n = el;
      while (n) {
        if (sel.startsWith(".") && String(n.className || "").split(" ").includes(sel.slice(1))) return n;
        n = n.parentNode;
      }
      return null;
    },
    setAttribute(k, v) { el._attrs[k] = String(v); },
    getAttribute(k) { return k in el._attrs ? el._attrs[k] : null; },
    addEventListener(t, fn) { (el._listeners[t] = el._listeners[t] || []).push(fn); },
    querySelector(sel) { return el.querySelectorAll(sel)[0] || null; },
    querySelectorAll(sel) {
      const out = [];
      const walk = (n) => {
        for (const c of n.children) {
          if (sel.startsWith(".") && String(c.className || "").split(" ").includes(sel.slice(1))) out.push(c);
          else if (sel === c.tag) out.push(c);
          walk(c);
        }
      };
      walk(el);
      return out;
    },
  };
  return el;
}

let agenticStyle = "harness";
let mobileViewport = false;
const document = {
  createElement(tag) { return makeEl(tag); },
  createTextNode(t) { const e = makeEl("#text"); e.textContent = t; return e; },
  body: makeEl("body"),
};
globalThis.document = document;
globalThis.window = globalThis;
globalThis.matchMedia = (q) => ({ matches: q === "(max-width: 1024px)" ? mobileViewport : false, media: q });
globalThis.addEventListener = () => {};
globalThis.removeEventListener = () => {};
globalThis.dispatchEvent = () => true;
globalThis.requestAnimationFrame = (fn) => setTimeout(fn, 0);
globalThis.cancelAnimationFrame = (id) => clearTimeout(id);
globalThis.localStorage = {
  getItem: (k) => (k === "mx.agentic.style" ? agenticStyle : null),
  setItem() {}, removeItem() {},
};

function textOf(root, cls) {
  const n = root.querySelector("." + cls);
  return n ? n.textContent : null;
}
function fullText(n) {
  let t = n.textContent || "";
  for (const c of n.children) t += fullText(c);
  return t;
}

const { ThreadView, renderDiff, diffVerb } = await import("../thread-view.js");

function newView(agentic = true) {
  const log = makeEl("div");
  const v = new ThreadView({
    log, streamURL: () => "", sendURL: "", stopURL: "",
    approveURL: "/api/agents/x/approve", agentic,
  });
  return { v, log };
}

function sampleDiff(n) {
  // n lignes : alternance + / - / contexte.
  const lines = [];
  for (let i = 0; i < n; i++) {
    const k = i % 3 === 0 ? "-" : i % 3 === 1 ? "+" : " ";
    lines.push({ kind: k, text: "ligne " + i });
  }
  return lines;
}

// --- diffVerb ---
test("diffVerb : verbe par outil", () => {
  assert.equal(diffVerb("Write"), "Écrit");
  assert.equal(diffVerb("Edit"), "Modifié");
  assert.equal(diffVerb("Sed"), "Modifié");
  assert.equal(diffVerb("Bash"), "Diff");
});

// --- renderDiff, vue Agents ---
test("renderDiff agents : verbe + chemin + +n -m colorés", () => {
  const lines = [
    { kind: "-", text: "ancien" },
    { kind: "+", text: "nouveau 1" },
    { kind: "+", text: "nouveau 2" },
    { kind: " ", text: "contexte" },
  ];
  const d = renderDiff(lines, "src/app.js", "Modifié");
  assert.equal(textOf(d, "diff-verb"), "Modifié");
  assert.equal(textOf(d, "diff-file"), "src/app.js");
  assert.equal(textOf(d, "diff-add-n"), "+2");
  assert.equal(textOf(d, "diff-del-n"), "-1");
  assert.equal(d.querySelector(".diff-verb"), d.querySelector(".diff-head").children[0]);
});

// --- renderDiff, chat général : en-tête historique inchangé ---
test("renderDiff général : pas de verbe, en-tête historique", () => {
  const lines = [{ kind: "+", text: "x" }];
  const d = renderDiff(lines, "f.txt", null);
  assert.equal(d.querySelector(".diff-verb"), null);
  assert.equal(textOf(d, "diff-stats"), "+1 / -0");
  assert.equal(d.querySelectorAll(".diff-row").length, 1);
});

// --- Aperçu borné ---
test("renderDiff : aperçu borné à 24 lignes, compteurs sur tout le diff", () => {
  const lines = sampleDiff(40); // 14 "-" (i%3==0), 13 "+" (i%3==1), 13 " "
  const d = renderDiff(lines, "gros.txt", "Modifié");
  const rows = d.querySelectorAll(".diff-row");
  // 24 lignes affichées + 1 ligne "… N lignes de plus".
  assert.equal(rows.length, 25);
  assert.equal(d.querySelectorAll(".diff-more").length, 1);
  assert.match(textOf(d, "diff-more"), /16 lignes de plus/);
  // Compteurs calculés sur les 40 lignes, pas sur l'aperçu.
  assert.equal(textOf(d, "diff-add-n"), "+13");
  assert.equal(textOf(d, "diff-del-n"), "-14");
});

test("renderDiff : pas de ligne '…' quand le diff tient dans l'aperçu", () => {
  const d = renderDiff(sampleDiff(10), "petit.txt", "Écrit");
  assert.equal(d.querySelector(".diff-more"), null);
  assert.equal(d.querySelectorAll(".diff-row").length, 10);
});

// --- Raisonnement : replié sur mobile (agents), déplié sinon ---
test("raisonnement : replié par défaut sur mobile (vue agents)", () => {
  mobileViewport = true;
  const { v } = newView(true);
  const det = v.ensureReasoning();
  assert.equal(det.open, false);
  mobileViewport = false;
});

test("raisonnement : déplié sur desktop (vue agents)", () => {
  mobileViewport = false;
  const { v } = newView(true);
  assert.equal(v.ensureReasoning().open, true);
});

test("raisonnement : chat général inchangé même sur mobile", () => {
  mobileViewport = true;
  const { v } = newView(false);
  assert.equal(v.ensureReasoning().open, true);
  mobileViewport = false;
});

// --- Ligne de statut : variante Codex ---
test("addTurnStats : variante codex discrète 'modèle · durée'", () => {
  agenticStyle = "codex";
  const { v } = newView(true);
  v.turnRoute = { label: "deepseek-chat" };
  v.turnElapsedMs = 4200;
  v.turnStats = { prompt_tokens: 1000, completion_tokens: 250 };
  const box = makeEl("div");
  v.addTurnStats(box);
  const cx = box.querySelector(".cx-turn-stats");
  assert.ok(cx, "ligne cx-turn-stats présente");
  assert.equal(cx.textContent, "deepseek-chat · 4.2s");
  assert.equal(box.querySelector(".turn-stats"), null);
  assert.equal(box.querySelector(".oc-turn-stats"), null);
  assert.match(cx.title, /Entrée/);
  agenticStyle = "harness";
});

test("addTurnStats : harness et opencode inchangés", () => {
  agenticStyle = "harness";
  const { v } = newView(true);
  v.turnRoute = { label: "m" };
  v.turnElapsedMs = 1000;
  v.turnStats = { prompt_tokens: 0, completion_tokens: 10 };
  const box = makeEl("div");
  v.addTurnStats(box);
  assert.equal(box.querySelector(".cx-turn-stats"), null);
  assert.ok(box.querySelector(".turn-stats"));

  agenticStyle = "opencode";
  const { v: v2 } = newView(true);
  v2.turnRoute = { label: "m" };
  v2.turnElapsedMs = 1000;
  const box2 = makeEl("div");
  v2.addTurnStats(box2);
  assert.equal(box2.querySelector(".cx-turn-stats"), null);
  assert.ok(box2.querySelector(".oc-turn-stats"));
  agenticStyle = "harness";
});
