import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";
// jsdom : résolution standard (node_modules du projet) puis repli /tmp.
const require = createRequire(import.meta.url);
let JSDOM;
try {
  ({ JSDOM } = require("jsdom"));
} catch {
  ({ JSDOM } = await import("/tmp/node_modules/jsdom/lib/api.js"));
}

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.confirm = () => true;

import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT = join(__dirname, "..");

const enginesPayload = {
  engines: [
    { id: "llamacpp", label: "llama.cpp", url: "", has_key: false },
    { id: "ollama", label: "Ollama", url: "http://192.168.1.10:11434", has_key: true },
    { id: "lmstudio", label: "LM Studio", url: "", has_key: false },
  ],
};

const calls = [];
globalThis.fetch = async (path, opts = {}) => {
  const method = opts.method || "GET";
  calls.push({ path, method, body: opts.body ? JSON.parse(opts.body) : undefined });
  const bodyFor = (p) => {
    if (p === "/api/local/engines") return enginesPayload;
    if (p === "/api/local/engines/ollama/test") return { ok: true, latency_ms: 12, models: 3 };
    if (p === "/api/local/engines/lmstudio/test") return { ok: false, error: "url non configuree" };
    if (p.startsWith("/api/local/engines/")) return { ok: true, id: "x", url: "", has_key: true };
    return {};
  };
  return { ok: true, status: 200, text: async () => JSON.stringify(bodyFor(path)) };
};

document.body.innerHTML = '<div id="samgen-body"></div>';

const mod = await import(`${ROOT}/samgen.js`);
await mod.initSamGenPanel();

test("init : 3 cartes moteurs, URL et état clé", () => {
  const cards = document.querySelectorAll("#samgen-body .samgen-card");
  assert.equal(cards.length, 3);
  assert.equal(document.getElementById("samgen-url-ollama").value, "http://192.168.1.10:11434");
  // Clé configurée : ligne masquée + lien Supprimer, pas de fuite.
  const masked = document.querySelector('.samgen-card[data-engine="ollama"] .apikey-masked-key');
  assert.ok(masked, "ligne masquée présente pour ollama");
  assert.ok(document.body.innerHTML.indexOf("gw-secret") === -1);
  assert.ok(!document.querySelector('.samgen-card[data-engine="llamacpp"] .apikey-masked-key'));
});

test("Enregistrer : PUT /api/local/engines/{id} avec url+clé", async () => {
  document.getElementById("samgen-url-lmstudio").value = "http://127.0.0.1:1234";
  document.getElementById("samgen-key-lmstudio").value = "k-test";
  document.getElementById("samgen-save-lmstudio").dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 50));
  const put = calls.find((c) => c.path === "/api/local/engines/lmstudio" && c.method === "PUT");
  assert.ok(put, "PUT émis vers le namespace SamGen");
  assert.equal(put.body.url, "http://127.0.0.1:1234");
  assert.equal(put.body.key, "k-test");
});

test("Tester : POST test, statut OK affiché", async () => {
  document.getElementById("samgen-test-ollama").dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 50));
  const t = calls.find((c) => c.path === "/api/local/engines/ollama/test" && c.method === "POST");
  assert.ok(t, "POST test émis");
  const st = document.getElementById("samgen-status-ollama");
  assert.ok(st.classList.contains("success"));
  assert.ok(st.textContent.includes("3 modèle"));
});

test("Tester sans URL : statut d'erreur affiché", async () => {
  document.getElementById("samgen-test-lmstudio").dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 50));
  const st = document.getElementById("samgen-status-lmstudio");
  assert.ok(st.classList.contains("error"));
});

test("Supprimer la clé : DELETE puis rechargement", async () => {
  const del = document.querySelector('.samgen-card[data-engine="ollama"] .apikey-delete-link');
  assert.ok(del, "lien Supprimer présent");
  del.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 80));
  const d = calls.find((c) => c.path === "/api/local/engines/ollama" && c.method === "DELETE");
  assert.ok(d, "DELETE émis vers le namespace SamGen");
});
