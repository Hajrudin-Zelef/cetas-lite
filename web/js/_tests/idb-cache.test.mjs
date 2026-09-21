import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

// jsdom : résolution standard puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

// Harness DOM AVANT d'importer les modules (projects.js -> dialogs.js).
const dom = new JSDOM("<!doctype html><html><head></head><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.cancelAnimationFrame = dom.window.cancelAnimationFrame.bind(dom.window);
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;
globalThis.EventSource = class {
  constructor() {}
  close() {}
};

// --- faux IndexedDB minimal (requêtes async comme le vrai) -----------------
function fakeReq(result) {
  const r = {};
  r.result = result;
  setTimeout(() => r.onsuccess && r.onsuccess(), 0);
  return r;
}

function makeFakeIdb() {
  const data = new Map();
  const db = {
    _data: data,
    createObjectStore() {},
    transaction() {
      return {
        objectStore() {
          return {
            get: (k) => fakeReq(data.get(k)),
            put: (v, k) => {
              data.set(k, v);
              return fakeReq(undefined);
            },
            delete: (k) => {
              data.delete(k);
              return fakeReq(undefined);
            },
            getAllKeys: () => fakeReq([...data.keys()]),
          };
        },
      };
    },
  };
  return {
    _data: data,
    open() {
      const req = { result: db };
      setTimeout(() => {
        req.onupgradeneeded && req.onupgradeneeded();
        req.onsuccess && req.onsuccess();
      }, 0);
      return req;
    },
  };
}

const idb = makeFakeIdb();
globalThis.indexedDB = idb;

// --- stub fetch --------------------------------------------------------------
let fetchCalls = [];
let blobMode = false;
let bigFile = false;
let treeVersion = 1; // le stub /tree renvoie une version différente à chaque appel

globalThis.fetch = async (url, opts = {}) => {
  const u = String(url);
  const method = opts.method || "GET";
  fetchCalls.push({ url: u, method });
  if (u.includes("/tree")) {
    const v = treeVersion++;
    return {
      ok: true,
      status: 200,
      text: async () => JSON.stringify({ path: "", version: v, children: [{ path: "a.txt", is_dir: false }] }),
    };
  }
  if (u.includes("/file?") && method === "GET") {
    if (blobMode) {
      return { ok: true, status: 200, blob: async () => new Blob(["x".repeat(100)]) };
    }
    const payload = bigFile ? { content: "x".repeat(3 << 20) } : { content: "hello" };
    return { ok: true, status: 200, text: async () => JSON.stringify(payload) };
  }
  if (u.includes("/file?") && method === "DELETE") {
    return { ok: true, status: 200, text: async () => JSON.stringify({ ok: true }) };
  }
  return { ok: true, status: 200, text: async () => "{}" };
};

const tick = (ms = 20) => new Promise((r) => setTimeout(r, ms));

const idbCache = await import("../idb-cache.js");
const IDB_TREE_TTL_MS = idbCache.IDB_TREE_TTL_MS;
const { ProjectsAPI: P } = await import("../projects.js");

test("idb-cache : stale-while-revalidate (cache servi, fond rafraîchi)", async () => {
  fetchCalls = [];
  treeVersion = 1;
  const r1 = await P.tree("p1", "", 1);
  assert.equal(r1.version, 1, "1er appel : version réseau 1");
  const r2 = await P.tree("p1", "", 1);
  assert.equal(r2.version, 1, "2e appel : servi depuis le cache (pas la version 2)");
  await tick(80); // laisse le rafraîchissement discret se produire
  const r3 = await P.tree("p1", "", 1);
  assert.equal(r3.version, 2, "3e appel : fond rafraîchi à la version 2");
});

test("idb-cache : entrée périmée => refetch réseau", async () => {
  // Graine périmée directement dans le faux store.
  idb._data.set("tree:p2::1", { t: Date.now() - IDB_TREE_TTL_MS - 1000, v: { stale: true } });
  fetchCalls = [];
  const r = await P.tree("p2", "", 1);
  assert.equal(fetchCalls.length, 1, "entrée périmée : refetch");
  assert.ok(!r.stale, "donnée fraîche servie");
});

test("idb-cache : removeFile invalide l'arbre et l'aperçu", async () => {
  fetchCalls = [];
  await P.tree("p3", "", 1);
  await P.file("p3", "a.txt");
  assert.ok(idb._data.has("tree:p3::1"));
  assert.ok(idb._data.has("file:p3:a.txt"));
  await P.removeFile("p3", "a.txt");
  assert.ok(!idb._data.has("tree:p3::1"), "clé arbre purgée");
  assert.ok(!idb._data.has("file:p3:a.txt"), "clé aperçu purgée");
  fetchCalls = [];
  await P.tree("p3", "", 1);
  assert.equal(fetchCalls.length, 1, "arbre re-fetché après invalidation");
});

test("idb-cache : invalidateProjectCache purge arbre + aperçus (hook onDone)", async () => {
  const { invalidateProjectCache } = await import("../idb-cache.js");
  fetchCalls = [];
  await P.tree("p8", "", 1);
  await P.file("p8", "a.txt");
  await invalidateProjectCache("p8");
  assert.equal(
    [...idb._data.keys()].filter((k) => k.startsWith("tree:p8:") || k.startsWith("file:p8:")).length,
    0,
    "clés IDB du projet purgées"
  );
  const before = fetchCalls.filter((c) => c.url.includes("/tree")).length;
  await P.tree("p8", "", 1);
  const after = fetchCalls.filter((c) => c.url.includes("/tree")).length;
  assert.equal(after, before + 1, "arbre re-fetché après invalidation");
});
test("idb-cache : aperçu > 2 Mo non mis en cache", async () => {
  bigFile = true;
  try {
    fetchCalls = [];
    await P.file("p4", "gros.bin");
    await P.file("p4", "gros.bin");
    assert.equal(fetchCalls.filter((c) => c.url.includes("/file?")).length, 2);
    assert.ok(!idb._data.has("file:p4:gros.bin"));
  } finally {
    bigFile = false;
  }
});

test("idb-cache : aperçu ≤ 2 Mo mis en cache", async () => {
  fetchCalls = [];
  const r1 = await P.file("p5", "petit.txt");
  const r2 = await P.file("p5", "petit.txt");
  assert.deepEqual(r2, r1);
  assert.equal(fetchCalls.filter((c) => c.url.includes("/file?")).length, 1);
});

test("idb-cache : fileBlob ≤ 2 Mo mis en cache", async () => {
  blobMode = true;
  try {
    fetchCalls = [];
    const b1 = await P.fileBlob("p6", "img.png");
    const b2 = await P.fileBlob("p6", "img.png");
    assert.ok(b1 instanceof Blob && b2 instanceof Blob);
    assert.equal(fetchCalls.filter((c) => c.url.includes("/file?")).length, 1);
  } finally {
    blobMode = false;
  }
});

test("idb-cache : sans IndexedDB, pas de plantage (overlay mémoire + réseau)", async () => {
  delete globalThis.indexedDB;
  try {
    const { ProjectsAPI: P2 } = await import("../idb-cache.js?no-idb=1");
    void P2;
    fetchCalls = [];
    const r1 = await P.tree("p7", "", 1);
    const r2 = await P.tree("p7", "", 1);
    assert.deepEqual(r2, r1, "données correctes sans IDB");
    assert.ok(
      fetchCalls.filter((c) => c.url.includes("/tree")).length <= 2,
      "pas d'appels en boucle sans IDB"
    );
  } finally {
    globalThis.indexedDB = idb;
  }
});
