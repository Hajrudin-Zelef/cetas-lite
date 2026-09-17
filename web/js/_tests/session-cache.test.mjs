// Cache localStorage des sessions : scopé par utilisateur, write-through,
// purge. Jamais de contenu de messages.
import { test } from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  try {
    ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
  } catch {
    console.log("jsdom indisponible, tests ignorés");
    process.exit(0);
  }
}

const dom = new JSDOM("<!doctype html><html><body></body></html>", { url: "http://localhost/" });
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.localStorage = dom.window.localStorage;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;

const cache = await import("../session-cache.js");

function loginAs(name) {
  dom.window.localStorage.setItem("cetas-lite-user", JSON.stringify({ username: name }));
}
function logout() {
  dom.window.localStorage.clear();
}

test("set/get : roundtrip de la liste", () => {
  logout();
  loginAs("sam");
  const list = [
    { id: "a", title: "Première", createdAt: 1, updatedAt: 3, messages: 2 },
    { id: "b", title: "Seconde", createdAt: 2, updatedAt: 2, messages: 0 },
  ];
  cache.setCachedSessions(list);
  assert.deepEqual(cache.getCachedSessions(), list);
});

test("scopé par utilisateur : aucun mélange", () => {
  logout();
  loginAs("sam");
  cache.setCachedSessions([{ id: "a", title: "Sam", createdAt: 1, updatedAt: 1, messages: 1 }]);
  loginAs("bob");
  assert.equal(cache.getCachedSessions(), null);
  cache.setCachedSessions([{ id: "z", title: "Bob", createdAt: 1, updatedAt: 1, messages: 1 }]);
  loginAs("sam");
  assert.equal(cache.getCachedSessions()[0].id, "a");
  loginAs("bob");
  assert.equal(cache.getCachedSessions()[0].id, "z");
});

test("upsert : insertion en tête + mise à jour", () => {
  logout();
  loginAs("sam");
  cache.setCachedSessions([{ id: "a", title: "A", createdAt: 1, updatedAt: 1, messages: 0 }]);
  cache.upsertCachedSession({ id: "b", title: "(sans titre)", createdAt: 2, updatedAt: 2, messages: 0 });
  let got = cache.getCachedSessions();
  assert.equal(got[0].id, "b");
  assert.equal(got.length, 2);
  cache.upsertCachedSession({ id: "b", title: "Renommée", messages: 2 });
  got = cache.getCachedSessions();
  assert.equal(got.length, 2);
  assert.equal(got[0].title, "Renommée");
  assert.equal(got[0].messages, 2);
});

test("remove : la session supprimée ne réapparaît pas", () => {
  logout();
  loginAs("sam");
  cache.setCachedSessions([
    { id: "a", title: "A", createdAt: 1, updatedAt: 1, messages: 1 },
    { id: "b", title: "B", createdAt: 2, updatedAt: 2, messages: 1 },
  ]);
  cache.removeCachedSession("b");
  const got = cache.getCachedSessions();
  assert.equal(got.length, 1);
  assert.equal(got[0].id, "a");
});

test("clearSessionCache : purge tous les utilisateurs", () => {
  logout();
  loginAs("sam");
  cache.setCachedSessions([{ id: "a", title: "A", createdAt: 1, updatedAt: 1, messages: 1 }]);
  loginAs("bob");
  cache.setCachedSessions([{ id: "z", title: "Z", createdAt: 1, updatedAt: 1, messages: 1 }]);
  cache.clearSessionCache();
  loginAs("sam");
  assert.equal(cache.getCachedSessions(), null);
  loginAs("bob");
  assert.equal(cache.getCachedSessions(), null);
});

test("ne stocke que des métadonnées (pas de contenu de messages)", () => {
  logout();
  loginAs("sam");
  cache.setCachedSessions([{ id: "a", title: "T", createdAt: 1, updatedAt: 1, messages: 4 }]);
  const raw = dom.window.localStorage.getItem("cetas.sessions.v1:sam");
  assert.ok(raw);
  assert.ok(!raw.includes("content"));
  const parsed = JSON.parse(raw);
  assert.deepEqual(Object.keys(parsed.sessions[0]).sort(), ["createdAt", "id", "messages", "title", "updatedAt"]);
});

test("JSON corrompu -> null (pas de crash)", () => {
  logout();
  loginAs("sam");
  dom.window.localStorage.setItem("cetas.sessions.v1:sam", "{corrompu");
  assert.equal(cache.getCachedSessions(), null);
});
