// Tests phase 2 refonte (2026-09-20) — cartes d'approbation premium.
// Exécutables sans jsdom (faux DOM minimal mais fonctionnel).
import assert from "node:assert/strict";
import { test } from "node:test";

// --- Faux DOM (querySelector/querySelectorAll fonctionnels, classes) ---
function makeEl(tag) {
  const kids = [];
  const el = {
    tag, children: kids, parentNode: null,
    className: "", textContent: "", hidden: false, disabled: false,
    dataset: {}, style: {},
    _attrs: {}, _listeners: {},
    appendChild(c) { c.parentNode = el; kids.push(c); return c; },
    insertBefore(c, ref) {
      c.parentNode = el;
      const i = ref ? kids.indexOf(ref) : -1;
      if (i < 0) kids.push(c); else kids.splice(i, 0, c);
      return c;
    },
    remove() {
      const i = kids.indexOf(el);
      if (el.parentNode) el.parentNode.children.splice(el.parentNode.children.indexOf(el), 1);
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
          if (sel.startsWith(".") && c.className.split(" ").includes(sel.slice(1))) out.push(c);
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
      return {
        add: (...cs) => {
          const set = new Set(el.className.split(" ").filter(Boolean));
          for (const c of cs) set.add(c);
          el.className = [...set].join(" ");
        },
        contains: (c) => el.className.split(" ").includes(c),
      };
    },
  });
  return el;
}

const document = {
  createElement(tag) { return makeEl(tag); },
  createTextNode(t) { const e = makeEl("#text"); e.textContent = t; return e; },
  body: makeEl("body"),
};
globalThis.document = document;
globalThis.window = globalThis;
globalThis.addEventListener = () => {};
globalThis.removeEventListener = () => {};
globalThis.dispatchEvent = () => true;
globalThis.requestAnimationFrame = (fn) => setTimeout(fn, 0);
globalThis.cancelAnimationFrame = (id) => clearTimeout(id);
globalThis.localStorage = { getItem: () => null, setItem() {}, removeItem() {} };

function findByClass(root, cls) {
  return root.querySelectorAll("." + cls);
}
function textOf(root, cls) {
  const n = root.querySelector("." + cls);
  return n ? n.textContent : null;
}
function fullText(n) {
  let t = n.textContent || "";
  for (const c of n.children) t += fullText(c);
  return t;
}

const { ThreadView } = await import("../thread-view.js?ap2");

let captured = null;
globalThis.fetch = async (url, opts) => {
  captured = { url, body: JSON.parse(opts.body) };
  return { ok: true, status: 200, text: async () => '{"ok":true}' };
};

function newView() {
  const log = makeEl("div");
  const v = new ThreadView({
    log, streamURL: () => "", sendURL: "", stopURL: "",
    approveURL: "/api/agents/x/approve", agentic: true,
  });
  captured = null;
  return { v, log };
}

function bashReq(id) {
  return { id, phase: "request", kind: "tool", tool: "Bash", args: { command: "npm run build", timeout: 30 } };
}

test("carte premium : bandeau, headline, identité outil, détail structuré", () => {
  const { v, log } = newView();
  v.addApproval(bashReq("a1"));
  const card = findByClass(log, "ap2")[0];
  assert.ok(card, "carte .ap2 créée");
  assert.equal(textOf(card, "ap2-banner-text"), "En attente de décision");
  assert.equal(textOf(card, "ap2-headline"), "Exécuter cette commande ?");
  const tool = findByClass(card, "ap2-tool")[0];
  assert.ok(tool, "ligne d'identité outil");
  assert.ok(fullText(tool).includes("Bash"), "nom de l'outil");
  assert.ok(fullText(tool).includes("$"), "glyphe sobre");
  assert.ok(fullText(tool).includes("npm run build"), "argument clé");
  const detail = findByClass(card, "ap2-detail")[0];
  assert.ok(detail, "détail structuré");
  assert.ok(fullText(detail).includes("npm run build"), "commande visible");
  assert.ok(!fullText(detail).includes('"command"'), "jamais de JSON brut");
  // Actions : Refuser / Toujours ce tour / Approuver.
  const btns = card.querySelectorAll("button").map((b) => b.textContent);
  assert.deepEqual(btns.filter((t) => !["Annuler", "Confirmer le refus"].includes(t)),
    ["Refuser", "Toujours ce tour", "Approuver"]);
});

test("détail Write : Fichier + Contenu borné, sans JSON", () => {
  const { v, log } = newView();
  const big = Array.from({ length: 50 }, (_, i) => "ligne " + i).join("\n");
  v.addApproval({ id: "a2", phase: "request", kind: "tool", tool: "Write",
    args: { file_path: "src/app.js", content: big } });
  const card = findByClass(log, "ap2")[0];
  assert.equal(textOf(card, "ap2-headline"), "Écrire ce fichier ?");
  const detail = findByClass(card, "ap2-detail")[0];
  const t = fullText(detail);
  assert.ok(t.includes("src/app.js"), "chemin visible");
  assert.ok(t.includes("50 lignes au total"), "contenu borné avec compteur");
  assert.ok(!t.includes('"file_path"'), "pas de JSON brut");
});

test("détail Edit : Remplacer / Par", () => {
  const { v, log } = newView();
  v.addApproval({ id: "a3", phase: "request", kind: "tool", tool: "Edit",
    args: { file_path: "a.txt", old: "foo", new: "bar" } });
  const card = findByClass(log, "ap2")[0];
  const t = fullText(findByClass(card, "ap2-detail")[0]);
  assert.ok(t.includes("Remplacer") && t.includes("foo"), "ancien texte");
  assert.ok(t.includes("Par") && t.includes("bar"), "nouveau texte");
});

test("carte plan : headline dédiée, pas de 'Toujours'", () => {
  const { v, log } = newView();
  v.addApproval({ id: "p1", phase: "request", kind: "plan", plan: "1. faire X\n2. faire Y" });
  const card = findByClass(log, "ap2")[0];
  assert.equal(textOf(card, "ap2-headline"), "Valider ce plan ?");
  assert.ok(fullText(card).includes("faire X"), "plan affiché");
  const btns = card.querySelectorAll("button").map((b) => b.textContent);
  assert.ok(!btns.includes("Toujours ce tour"), "pas de 'Toujours' pour un plan");
  assert.ok(btns.includes("Valider le plan"), "bouton plan");
});

test("refus : champ commentaire révélé puis envoyé au serveur", async () => {
  const { v, log } = newView();
  v.addApproval(bashReq("a4"));
  const card = findByClass(log, "ap2")[0];
  const byText = (t) => card.querySelectorAll("button").find((b) => b.textContent === t);
  byText("Refuser").click();
  const denyBox = findByClass(card, "ap2-denybox")[0];
  assert.equal(denyBox.hidden, false, "champ 'que faire différemment ?' révélé");
  assert.equal(textOf(denyBox, "ap2-deny-label"), "Que faire différemment ?");
  assert.ok(byText("Annuler") && !byText("Annuler").hidden, "bouton Annuler visible");
  const input = findByClass(card, "ap2-deny-input")[0];
  input.value = "utilise npm plutôt que yarn";
  byText("Confirmer le refus").click();
  await new Promise((r) => setTimeout(r, 10));
  assert.ok(captured, "POST envoyé");
  assert.equal(captured.url, "/api/agents/x/approve");
  assert.deepEqual(captured.body,
    { id: "a4", approved: false, always: false, comment: "utilise npm plutôt que yarn" });
});

test("résolution : carte figée + trace persistante (refus commenté)", () => {
  const { v, log } = newView();
  v.addApproval(bashReq("a5"));
  const card = findByClass(log, "ap2")[0];
  card._ap2.decision = { approved: false, always: false, comment: "trop risqué" };
  v.addApproval({ id: "a5", phase: "resolved", approved: false });
  assert.equal(card.getAttribute("data-resolved"), "1", "carte figée");
  assert.ok(card.querySelectorAll("button").every((b) => b.disabled), "boutons désactivés");
  assert.equal(textOf(card, "ap2-banner-text"), "Décision enregistrée");
  const traces = findByClass(log, "ap2-trace");
  assert.equal(traces.length, 1, "trace ajoutée au fil");
  const t = fullText(traces[0]);
  assert.ok(t.startsWith("✖ Refusé"), "préfixe refus");
  assert.ok(t.includes("trop risqué"), "commentaire dans la trace");
  assert.ok(t.includes("Bash") && t.includes("npm run build"), "outil + argument clé");
  // La trace suit la carte dans le fil.
  const kids = log.children;
  assert.ok(kids.indexOf(traces[0]) === kids.indexOf(card) + 1, "trace juste après la carte");
});

test("résolution : trace approbation et 'toujours ce tour'", () => {
  const { v, log } = newView();
  v.addApproval({ id: "a6", phase: "request", kind: "tool", tool: "Edit",
    args: { file_path: "src/app.js", old: "a", new: "b" } });
  const card = findByClass(log, "ap2")[0];
  card._ap2.decision = { approved: true, always: true, comment: "" };
  v.addApproval({ id: "a6", phase: "resolved", approved: true });
  const t = fullText(findByClass(log, "ap2-trace")[0]);
  assert.ok(t.startsWith("✔ Toujours approuver (ce tour)"), "trace 'toujours'");
  assert.ok(t.includes("Edit") && t.includes("src/app.js"), "outil + fichier");
});

test("résolution : expiration → trace ⌛", () => {
  const { v, log } = newView();
  v.addApproval(bashReq("a7"));
  const card = findByClass(log, "ap2")[0];
  v.addApproval({ id: "a7", phase: "resolved", approved: false, timeout: true });
  assert.equal(textOf(card, "ap2-banner-text"), "Expirée");
  const t = fullText(findByClass(log, "ap2-trace")[0]);
  assert.ok(t.startsWith("⌛"), "trace expiration");
});

test("chat général : carte historique strictement inchangée", () => {
  const log = makeEl("div");
  const v = new ThreadView({
    log, streamURL: () => "", sendURL: "", stopURL: "",
    approveURL: "/api/chat/approve",
    // pas de flag agentic → chat général
  });
  v.addApproval(bashReq("g1"));
  assert.equal(findByClass(log, "ap2").length, 0, "pas de carte premium");
  const card = findByClass(log, "msg-approval")[0];
  assert.ok(card, "carte historique créée");
  assert.equal(textOf(card, "approval-title"), "Approbation requise");
  assert.equal(textOf(card, "approval-status"), "En attente de ta décision…");
  assert.ok(fullText(findByClass(card, "approval-args")[0]).includes('"command"'),
    "l'ancien JSON brut est conservé (comportement historique)");
  const btns = card.querySelectorAll("button").map((b) => b.textContent);
  assert.deepEqual(btns, ["Approuver", "Refuser", "Toujours approuver (ce tour)"]);
  v.addApproval({ id: "g1", phase: "resolved", approved: true });
  assert.equal(textOf(card, "approval-status"), "Approuvé");
  assert.ok(findByClass(card, "approved").length > 0, "classe approved historique");
  assert.equal(findByClass(log, "ap2-trace").length, 0, "pas de trace premium");
});
