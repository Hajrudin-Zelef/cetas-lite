import test from "node:test";
import assert from "node:assert/strict";

// --- Faux DOM avec querySelector fonctionnel pour la déduplication ---
function makeEl(tag) {
  const kids = [];
  const attrs = {};
  const e = {
    tagName: String(tag || "").toUpperCase(),
    children: kids,
    style: {},
    dataset: {},
    className: "",
    textContent: "",
    innerHTML: "",
    hidden: false,
    classList: {
      add(c) { if (!e.className.split(" ").includes(c)) e.className += (e.className ? " " : "") + c; },
      remove(c) { e.className = e.className.split(" ").filter((x) => x !== c).join(" "); },
      toggle() {},
      contains(c) { return e.className.split(" ").includes(c); },
    },
    appendChild(c) { kids.push(c); c.parentNode = e; return c; },
    append(...a) { for (const c of a) c.parentNode = e; kids.push(...a); return a[0]; },
    insertBefore(c, ref) {
      c.parentNode = e;
      const i = ref ? kids.indexOf(ref) : -1;
      if (i >= 0) kids.splice(i, 0, c); else kids.push(c);
      return c;
    },
    remove() { e._removed = true; if (e.parentNode) { const i = e.parentNode.children.indexOf(e); if (i >= 0) e.parentNode.children.splice(i, 1); } },
    before() {},
    // Supporte le sélecteur utilisé par la déduplication :
    // ':scope > .message-wrapper-user[data-optimistic="1"]'
    querySelector(sel) {
      const m = sel.match(/:scope > \.([\w-]+)\[([\w-]+)="([^"]+)"\]/);
      if (!m) return null;
      const [, cls, attr, val] = m;
      return kids.find((k) => !k._removed && k.className.split(" ").includes(cls) && k.getAttribute(attr) === val) || null;
    },
    querySelectorAll() { return []; },
    closest() { return null; },
    setAttribute(k, v) { attrs[k] = String(v); },
    getAttribute(k) { return attrs[k] ?? null; },
    removeAttribute(k) { delete attrs[k]; },
    hasAttribute(k) { return k in attrs; },
    addEventListener() {},
    removeEventListener() {},
    scrollIntoView() {},
    get isConnected() { return !e._removed; },
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
globalThis.localStorage = { getItem: () => null, setItem: () => {}, removeItem: () => {} };
globalThis.requestAnimationFrame = (fn) => 0;
globalThis.cancelAnimationFrame = () => {};
Object.defineProperty(globalThis, "crypto", {
  value: { randomUUID: () => "uuid-test-1" },
  configurable: true,
});

const { ThreadView } = await import("../thread-view.js");

const views = [];
test.afterEach(() => {
  // showWait() lance un setInterval : on le nettoie pour laisser
  // l'event loop se vider.
  for (const v of views.splice(0)) v.hideWait();
});

function makeView() {
  const log = makeEl("div");
  const v = new ThreadView({
    log,
    streamURL: () => "",
    sendURL: "",
    stopURL: "",
    optimisticEcho: true,
    reasonHooks: { append() {}, finish() {}, reset() {} },
  });
  views.push(v);
  return { v, log };
}

function userWrappers(log) {
  return log.children.filter(
    (k) => !k._removed && k.className.split(" ").includes("message-wrapper-user")
  );
}

test("écho optimiste : le delta user avec le même client_msg_id ne duplique pas", () => {
  const { v, log } = makeView();
  const w = v.addUserOptimistic("bonjour");
  assert.equal(w.getAttribute("data-optimistic"), "1");
  v.pendingUserId = "uuid-test-1";
  v.handleEvent({ seq: 1, user: "bonjour", client_msg_id: "uuid-test-1" });
  assert.equal(userWrappers(log).length, 1, "un seul message utilisateur");
  assert.equal(v.pendingUserId, null, "pendingUserId réinitialisé");
  assert.equal(w.getAttribute("data-optimistic"), null, "marqueur optimiste retiré");
});

test("client_msg_id différent : le delta user s'affiche normalement", () => {
  const { v, log } = makeView();
  v.addUserOptimistic("bonjour");
  v.pendingUserId = "uuid-test-1";
  v.handleEvent({ seq: 1, user: "bonjour", client_msg_id: "autre-id" });
  assert.equal(userWrappers(log).length, 2, "deux messages (pas de faux positif)");
});

test("sans client_msg_id (ancien serveur) : comportement historique", () => {
  const { v, log } = makeView();
  v.handleEvent({ seq: 1, user: "salut" });
  assert.equal(userWrappers(log).length, 1);
});

test("removeOptimistic retire le nœud et réinitialise", () => {
  const { v, log } = makeView();
  const w = v.addUserOptimistic("oups");
  v.pendingUserId = "x";
  v.removeOptimistic(w);
  assert.equal(userWrappers(log).length, 0);
  assert.equal(v.pendingUserId, null);
});

test("newClientMsgId génère un identifiant non vide", () => {
  const { v } = makeView();
  const id = v.newClientMsgId();
  assert.ok(typeof id === "string" && id.length > 0);
});
