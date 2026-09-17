// Déconnexion : le cache des sessions (métadonnées) est purgé — aucune
// trace d'un compte précédent ne reste sur la machine.
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

const apiMod = await import("../api.js");
const cache = await import("../session-cache.js");

test("clearSession() purge le cache des sessions", () => {
  dom.window.localStorage.clear();
  dom.window.localStorage.setItem("cetas-lite-user", JSON.stringify({ username: "sam" }));
  dom.window.localStorage.setItem("cetas-lite-token", "tok");
  cache.setCachedSessions([{ id: "a", title: "A", createdAt: 1, updatedAt: 1, messages: 1 }]);
  assert.ok(cache.getCachedSessions());

  apiMod.clearSession();

  assert.equal(dom.window.localStorage.getItem("cetas-lite-token"), null);
  assert.equal(dom.window.localStorage.getItem("cetas-lite-user"), null);
  assert.equal(cache.getCachedSessions(), null);
  assert.equal(dom.window.localStorage.getItem("cetas.sessions.v1:sam"), null);
});
