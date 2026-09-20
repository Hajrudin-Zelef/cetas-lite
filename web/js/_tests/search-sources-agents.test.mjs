// Tests "pas de panneau Sources en vue Agents" (2026-09-20) — la recherche
// web n'affiche que le résumé du modèle : ni panneau Sources, ni sortie
// brute dans le fil Agents. Le chat général garde son rendu historique.
// Exécutables sans jsdom (faux DOM minimal).
import assert from "node:assert/strict";
import { test } from "node:test";

// --- Faux DOM ---
function makeEl(tag) {
  const kids = [];
  const el = {
    tag, children: kids, parentNode: null,
    className: "", textContent: "", innerHTML: "", hidden: false, disabled: false,
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
        else if (sel === n.tag) return n;
        n = n.parentNode;
      }
      return null;
    },
    setAttribute(k, v) { el._attrs[k] = String(v); },
    getAttribute(k) { return k in el._attrs ? el._attrs[k] : null; },
    hasAttribute(k) { return k in el._attrs; },
    removeAttribute(k) { delete el._attrs[k]; },
    addEventListener(t, fn) { (el._listeners[t] = el._listeners[t] || []).push(fn); },
    click() { for (const fn of el._listeners.click || []) fn({ preventDefault() {} }); },
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
  Object.defineProperty(el, "classList", {
    get() {
      const sync = (set) => { el.className = [...set].join(" "); };
      const get = () => new Set(String(el.className || "").split(" ").filter(Boolean));
      return {
        add(...cs) { const s = get(); for (const c of cs) s.add(c); sync(s); },
        remove(...cs) { const s = get(); for (const c of cs) s.delete(c); sync(s); },
        toggle(c, force) {
          const s = get();
          const on = force === undefined ? !s.has(c) : !!force;
          if (on) s.add(c); else s.delete(c);
          sync(s);
          return on;
        },
        contains(c) { return get().has(c); },
      };
    },
  });
  return el;
}

let agenticStyle = "harness";
const document = {
  createElement(tag) { return makeEl(tag); },
  createTextNode(t) { const e = makeEl("#text"); e.textContent = t; return e; },
  body: makeEl("body"),
};
globalThis.document = document;
globalThis.window = globalThis;
globalThis.matchMedia = () => ({ matches: false, media: "" });
globalThis.addEventListener = () => {};
globalThis.removeEventListener = () => {};
globalThis.dispatchEvent = () => true;
globalThis.requestAnimationFrame = (fn) => setTimeout(fn, 0);
globalThis.cancelAnimationFrame = (id) => clearTimeout(id);
globalThis.localStorage = {
  getItem: (k) => (k === "mx.agentic.style" ? agenticStyle : null),
  setItem() {}, removeItem() {},
};

const { ThreadView } = await import("../thread-view.js");

function newView(agentic = true, style = "harness") {
  agenticStyle = style;
  const log = makeEl("div");
  const v = new ThreadView({
    log, streamURL: () => "", sendURL: "", stopURL: "",
    approveURL: "/api/agents/x/approve", agentic,
  });
  return { v, log };
}

const SOURCES = [
  { title: "Titre A", url: "https://a.example/x" },
  { title: "Titre B", url: "https://b.example/y" },
];

function runSearch(v, name, args, extraEnd = {}) {
  v.addTool({ name, args, phase: "start" });
  v.addTool({
    name, args, phase: "end",
    result: "Resultats exa (2):\n[1] Titre A\nhttps://a.example/x\ndesc",
    sources: SOURCES,
    ...extraEnd,
  });
}

function fullText(n) {
  let t = n.textContent || "";
  for (const c of n.children) t += fullText(c);
  return t;
}

// --- Vue Agents, style Harness ---
test("agents/harness : web_search ne rend ni Sources ni sortie brute", () => {
  const { v, log } = newView(true, "harness");
  const args = { query: "prix du cuivre" };
  v.addTool({ name: "web_search", args, phase: "start" });
  assert.ok(log.querySelector(".search-status"), "indicateur attendu pendant la recherche");
  v.addTool({
    name: "web_search", args, phase: "end",
    result: "Resultats exa (2):\n[1] Titre A\nhttps://a.example/x\ndesc",
    sources: SOURCES,
  });
  assert.equal(log.querySelector(".search-status"), null, "indicateur retiré");
  assert.equal(log.querySelector(".citations-block"), null, "pas de panneau Sources");
  assert.equal(log.querySelector(".tool-result"), null, "pas de sortie brute");
  assert.ok(log.querySelector(".msg-tool"), "la ligne d'outil reste visible");
});

test("agents/harness : web_fetch sans sources ne rend pas la sortie brute", () => {
  const { v, log } = newView(true, "harness");
  const args = { url: "https://a.example/x" };
  v.addTool({ name: "web_fetch", args, phase: "start" });
  v.addTool({ name: "web_fetch", args, phase: "end", result: "[info] fetch: echec" });
  assert.equal(log.querySelector(".citations-block"), null);
  assert.equal(log.querySelector(".tool-result"), null);
});

// --- Vue Agents, style Codex ---
test("agents/codex : web_search ne rend ni Sources ni sortie brute", () => {
  const { v, log } = newView(true, "codex");
  const args = { query: "prix du cuivre" };
  runSearch(v, "web_search", args);
  assert.equal(log.querySelector(".citations-block"), null, "pas de panneau Sources");
  assert.equal(log.querySelector(".tool-result"), null, "pas de sortie brute");
  const head = log.querySelector(".cx-tool-head");
  assert.ok(head, "l'en-tête d'outil reste visible");
  assert.match(fullText(head), /Ran/);
});

// --- Vue Agents, style OpenCode ---
test("agents/opencode : web_search ne rend ni Sources ni sortie brute", () => {
  const { v, log } = newView(true, "opencode");
  const args = { query: "prix du cuivre" };
  runSearch(v, "web_search", args);
  assert.equal(log.querySelector(".citations-block"), null, "pas de panneau Sources");
  assert.equal(log.querySelector(".tool-result"), null, "pas de sortie brute");
  assert.ok(log.querySelector(".oc-tool"), "le bloc d'outil reste visible");
});

// --- Recherche native (provider), vue Agents ---
test("agents : recherche native sans panneau Sources", () => {
  const { v, log } = newView(true, "harness");
  v.handleEvent({ search: { phase: "start", native: true } });
  assert.ok(log.querySelector(".search-status"), "indicateur natif attendu");
  v.handleEvent({ search: { phase: "done", native: true, sources: SOURCES } });
  assert.equal(log.querySelector(".search-status"), null, "indicateur retiré");
  assert.equal(log.querySelector(".citations-block"), null, "pas de panneau Sources natif");
});

// --- Chat général : comportement historique inchangé ---
test("général : web_search rend toujours le panneau Sources", () => {
  const { v, log } = newView(false, "harness");
  const args = { query: "prix du cuivre" };
  runSearch(v, "web_search", args);
  const block = log.querySelector(".citations-block");
  assert.ok(block, "panneau Sources conservé dans le chat général");
  assert.equal(block.querySelectorAll(".citation-card").length, 2);
});

test("général : recherche native rend toujours le panneau Sources", () => {
  const { v, log } = newView(false, "harness");
  v.handleEvent({ search: { phase: "start", native: true } });
  v.handleEvent({ search: { phase: "done", native: true, sources: SOURCES } });
  assert.ok(log.querySelector(".citations-block"), "panneau Sources natif conservé");
});

// --- Recherche en échec : le signal visuel d'erreur est conservé ---
test("agents/harness : recherche en échec garde la bordure rouge", () => {
  const { v, log } = newView(true, "harness");
  const args = { query: "prix du cuivre" };
  v.addTool({ name: "web_search", args, phase: "start" });
  v.addTool({ name: "web_search", args, phase: "end", result: "[erreur] recherche web: tous les backends ont échoué" });
  const box = log.querySelector(".msg-tool");
  assert.ok(box, "boîte d'outil attendue");
  assert.ok(box.classList.contains("is-error"), "marquage d'erreur attendu");
  assert.equal(log.querySelector(".citations-block"), null);
  assert.equal(log.querySelector(".tool-result"), null);
});

test("agents/harness : recherche réussie sans marquage d'erreur", () => {
  const { v, log } = newView(true, "harness");
  const args = { query: "prix du cuivre" };
  runSearch(v, "web_search", args);
  const box = log.querySelector(".msg-tool");
  assert.ok(box, "boîte d'outil attendue");
  assert.ok(!box.classList.contains("is-error"), "pas de marquage d'erreur");
});

test("agents/opencode : recherche en échec garde la bordure rouge", () => {
  const { v, log } = newView(true, "opencode");
  const args = { query: "prix du cuivre" };
  v.addTool({ name: "web_search", args, phase: "start" });
  v.addTool({ name: "web_search", args, phase: "end", result: "[erreur] recherche web: limite atteinte" });
  const box = log.querySelector(".oc-tool");
  assert.ok(box, "bloc d'outil attendu");
  assert.ok(box.classList.contains("is-error"), "marquage d'erreur attendu");
  assert.equal(log.querySelector(".citations-block"), null);
  assert.equal(log.querySelector(".tool-result"), null);
});

test("agents/codex : recherche en échec garde la pastille rouge", () => {
  const { v, log } = newView(true, "codex");
  const args = { query: "prix du cuivre" };
  v.addTool({ name: "web_search", args, phase: "start" });
  v.addTool({ name: "web_search", args, phase: "end", result: "[erreur] recherche web: timeout" });
  const box = log.querySelector(".cx-tool");
  assert.ok(box, "bloc d'outil attendu");
  assert.ok(box.classList.contains("is-error"), "marquage d'erreur attendu");
  assert.equal(log.querySelector(".citations-block"), null);
  assert.equal(log.querySelector(".tool-result"), null);
});
