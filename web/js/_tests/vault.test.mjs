import test from "node:test";
import assert from "node:assert/strict";
import fs from "node:fs";
import { createRequire } from "node:module";
// jsdom : résolution standard (node_modules du projet) puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

const dom = new JSDOM(
  `<!doctype html><html><body><div id="vault-body"></div></body></html>`,
  { pretendToBeVisual: true, url: "http://localhost/" }
);
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.localStorage = dom.window.localStorage;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;

const vault = await import("../vault.js");

// Mock fetch programmable : { [path]: { status, body } | Error }
let routes = {};
let lastReq = null;
globalThis.fetch = async (path, opts = {}) => {
  lastReq = { path, opts };
  const r = routes[path];
  if (r instanceof Error) throw r;
  const status = r?.status ?? 404;
  const body = r?.body;
  return {
    status,
    ok: status >= 200 && status < 300,
    text: async () => (body === undefined ? "" : JSON.stringify(body)),
  };
};
const setRoutes = (r) => {
  routes = r;
};

await test("escapeHtml neutralise le HTML", () => {
  assert.equal(vault.escapeHtml('<b>"x"&'), "&lt;b&gt;&quot;x&quot;&amp;");
});

await test("formatValue : string telle quelle, objet en JSON", () => {
  assert.equal(vault.formatValue("abc"), "abc");
  assert.equal(vault.formatValue({ a: 1 }), '{"a":1}');
  assert.equal(vault.formatValue(42), "42");
});

await test("vaultFetch envoie le token et lève le message serveur", async () => {
  dom.window.localStorage.setItem("cetas-lite-token", "tok123");
  setRoutes({ "/api/vault/status": { status: 200, body: { exists: true, unlocked: false } } });
  const st = await vault.vaultFetch("/api/vault/status");
  assert.equal(st.unlocked, false);
  assert.equal(lastReq.opts.headers["Authorization"], "Bearer tok123");

  setRoutes({ "/api/vault/unlock": { status: 401, body: { error: "mot de passe incorrect" } } });
  await assert.rejects(
    vault.vaultFetch("/api/vault/unlock", { method: "POST", body: { password: "x" } }),
    /mot de passe incorrect/
  );
  // Pas de clearSession : le token reste (un 401 coffre ≠ session expirée).
  assert.equal(dom.window.localStorage.getItem("cetas-lite-token"), "tok123");
});

await test("vaultFetch : erreur réseau", async () => {
  setRoutes({ "/api/vault/status": new Error("boom") });
  await assert.rejects(vault.vaultFetch("/api/vault/status"), /Connexion au serveur impossible/);
});

await test("loadVaultPanel : sans coffre → écran de création", async () => {
  setRoutes({ "/api/vault/status": { status: 200, body: { exists: false, unlocked: false } } });
  await vault.loadVaultPanel();
  const body = document.getElementById("vault-body");
  assert.ok(body.querySelector("#vault-create-btn"), "bouton créer présent");
  assert.ok(body.querySelector("#vault-new-pw"), "champ mot de passe présent");
});

await test("loadVaultPanel : coffre verrouillé → écran de déverrouillage", async () => {
  setRoutes({ "/api/vault/status": { status: 200, body: { exists: true, unlocked: false } } });
  await vault.loadVaultPanel();
  const body = document.getElementById("vault-body");
  assert.ok(body.querySelector("#vault-unlock-btn"), "bouton déverrouiller présent");
});

await test("loadVaultPanel : déverrouillé → liste des secrets", async () => {
  setRoutes({
    "/api/vault/status": { status: 200, body: { exists: true, unlocked: true } },
    "/api/vault/entries": { status: 200, body: { keys: ["github", "openrouter"] } },
  });
  await vault.loadVaultPanel();
  const body = document.getElementById("vault-body");
  const rows = body.querySelectorAll(".vault-row");
  assert.equal(rows.length, 2);
  assert.equal(rows[0].dataset.key, "github");
  assert.ok(body.querySelector("#vault-lock-btn"), "bouton verrouiller présent");
  assert.ok(body.querySelector("#vault-add-btn"), "bouton ajouter présent");
});

await test("modals.js importe bien loadVaultPanel depuis vault.js", async () => {
  const src = fs.readFileSync(new URL("../modals.js", import.meta.url), "utf8");
  const m = src.match(/import\s*\{([^}]+)\}\s*from\s*"\.\/vault\.js"/);
  assert.ok(m, "import de vault.js trouvé dans modals.js");
  for (const name of m[1].split(",").map((s) => s.trim()).filter(Boolean)) {
    assert.equal(typeof vault[name], "function", `vault.js doit exporter ${name}`);
  }
  assert.ok(src.includes('data-tab="vault"') === false, "le bouton d'onglet est dans index.html, pas modals.js");
});

await test("index.html : onglet + panneau coffre présents", () => {
  const html = fs.readFileSync(new URL("../../index.html", import.meta.url), "utf8");
  assert.ok(html.includes('data-tab="vault"'), "bouton onglet coffre");
  assert.ok(html.includes('id="panel-vault"'), "panneau coffre");
  assert.ok(html.includes('id="vault-body"'), "conteneur vault-body");
  assert.ok(html.includes("mx-vault.css"), "CSS du coffre chargé");
});
