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

// ---------- helpers purs ----------

import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const ROOT = join(__dirname, "..");
const mod = await import(`${ROOT}/apimodeles.js`);

test("helpers : PROVIDERS_META, escHtml, makerLabel, priceStr", () => {
  const { _test } = mod;
  assert.equal(_test.PROVIDERS_META.length, 7);
  assert.deepEqual(_test.PROVIDERS_META.map((m) => m.id),
    ["openrouter", "deepseek", "opencode", "opencode-go", "llamacpp", "ollama", "lmstudio"]);
  assert.ok(_test.PROVIDERS_META.every((m) => m.icon && m.label));
  assert.equal(_test.escHtml('<b>"x"</b>'), "&lt;b&gt;&quot;x&quot;&lt;/b&gt;");
  assert.equal(_test.makerLabel("openai/gpt-4o"), "OpenAI");
  assert.equal(_test.makerLabel("qwen/qwen3"), "Qwen");
  assert.equal(_test.priceStr(0, 0), "Gratuit");
  assert.ok(_test.priceStr(1.25, 10).includes("$1.25"));
  assert.ok(_test.priceStr(1.25, 10).includes("/M"));
});

// ---------- init : onglets et sections ----------

const providersPayload = {
  providers: [
    { id: "openrouter", configured: true },
    { id: "deepseek", configured: false },
    { id: "opencode", configured: false },
    { id: "opencode-go", configured: false },
    { id: "llamacpp", configured: false, url: "" },
    { id: "ollama", configured: true, url: "http://192.168.1.10:11434" },
    { id: "lmstudio", configured: false, url: "" },
  ],
};
const orPayload = {
  models: [
    { id: "openai/gpt-4o", name: "GPT-4o", prompt_per_1m: 2.5, completion_per_1m: 10, context_length: 128000, description: "Multimodal." },
    { id: "anthropic/claude-3", name: "Claude 3", prompt_per_1m: 0, completion_per_1m: 0, context_length: 200000, description: "Free tier." },
    { id: "qwen/qwen3", name: "Qwen 3", prompt_per_1m: 0.1, completion_per_1m: 0.2, context_length: 32768, description: "" },
  ],
};
const staticPayload = {
  providers: [
    { id: "openrouter", models: [] },
    { id: "deepseek", models: [{ id: "deepseek-chat", label: "DeepSeek Chat", input_per_1m: 0.27, output_per_1m: 1.1 }] },
    { id: "opencode", models: [{ id: "opencode-zen-chat", label: "OpenCode Zen Chat", input_per_1m: 0, output_per_1m: 0 }] },
    { id: "opencode-go", models: [] },
  ],
};

const calls = [];
globalThis.fetch = async (path, opts = {}) => {
  const method = opts.method || "GET";
  calls.push({ path, method, body: opts.body ? JSON.parse(opts.body) : undefined });
  const bodyFor = (p) => {
    if (p === "/api/providers") return providersPayload;
    if (p === "/api/catalog/selection") return { disabled: {} };
    if (p === "/api/catalog") return staticPayload;
    if (p.startsWith("/api/openrouter/models")) return orPayload;
    return {};
  };
  return { ok: true, status: 200, text: async () => JSON.stringify(bodyFor(path)) };
};

document.body.innerHTML =
  '<div id="providers-tabs" class="providers-tabs"></div>' +
  '<div id="provider-content" class="provider-content"></div>' +
  '<button id="apimodeles-save-btn">Sauvegarder tout</button>';

await mod.initApiModelesPanel();

test("init : 7 onglets, OpenRouter actif par défaut, sections rendues", () => {
  const tabs = document.querySelectorAll("#providers-tabs .provider-tab");
  assert.equal(tabs.length, 7);
  assert.ok(tabs[0].classList.contains("active"));
  assert.equal(tabs[0].dataset.provider, "openrouter");
  assert.ok(tabs[0].querySelector("img.provider-tab-icon"));
  assert.ok(tabs[0].textContent.includes("OpenRouter"));
  const secs = document.querySelectorAll("#provider-content .provider-section");
  assert.equal(secs.length, 7);
  assert.ok(document.querySelector('.provider-section[data-provider="openrouter"]').classList.contains("active"));
});

test("OpenRouter : toggle Textes/Images, lignes avec prix, recherche", () => {
  const list = document.getElementById("catalog-list-openrouter");
  assert.ok(list, "catalog list OR présent");
  const rows = list.querySelectorAll(".catalog-row");
  assert.equal(rows.length, 3);
  assert.ok(rows[0].textContent.includes("GPT-4o"));
  assert.ok(rows[0].textContent.includes("$2.50"));
  assert.ok(rows[1].textContent.includes("Gratuit"));
  assert.equal(list.querySelectorAll(".catalog-row-info").length, 3); // contexte/descriptions
  // Toggle Images visible uniquement pour OpenRouter.
  const toggle = document.querySelector('.provider-section[data-provider="openrouter"] .catalog-type-btns');
  assert.ok(toggle);
  assert.ok(!document.querySelector('.provider-section[data-provider="deepseek"] .catalog-type-btns'));
  // Recherche filtre.
  const search = document.getElementById("catalog-search-openrouter");
  search.value = "qwen";
  search.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  assert.equal(document.getElementById("catalog-list-openrouter").querySelectorAll(".catalog-row").length, 1);
  search.value = "";
  search.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  // Catégories dérivées des makers.
  const cat = document.getElementById("catalog-cat-openrouter");
  assert.ok(cat.querySelectorAll("option").length >= 4); // Toutes + OpenAI/Anthropic/Qwen
  cat.value = "Anthropic";
  cat.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  assert.equal(document.getElementById("catalog-list-openrouter").querySelectorAll(".catalog-row").length, 1);
});

test("changement d'onglet : DeepSeek, catalogue statique", async () => {
  const tab = document.querySelector('.provider-tab[data-provider="deepseek"]');
  tab.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 60)); // chargement async du catalogue
  assert.ok(tab.classList.contains("active"));
  assert.ok(document.querySelector('.provider-section[data-provider="deepseek"]').classList.contains("active"));
  assert.ok(!document.querySelector('.provider-section[data-provider="openrouter"]').classList.contains("active"));
  const list = document.getElementById("catalog-list-deepseek");
  const rows = list.querySelectorAll(".catalog-row");
  assert.equal(rows.length, 1);
  assert.ok(rows[0].textContent.includes("DeepSeek Chat"));
});

test("œil : afficher/masquer la clé", () => {
  const input = document.getElementById("apikey-deepseek");
  const eye = document.querySelector('.provider-section[data-provider="deepseek"] .apikey-eye-btn');
  assert.equal(input.type, "password");
  eye.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  assert.equal(input.type, "text");
  assert.ok(eye.classList.contains("shown"));
  eye.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  assert.equal(input.type, "password");
});

test("Valider : PUT clé puis ligne masquée affichée", async () => {
  const input = document.getElementById("apikey-deepseek");
  input.value = "sk-test-123";
  const btn = document.querySelector('.provider-section[data-provider="deepseek"] .apikey-validate-btn');
  btn.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 50));
  const put = calls.find((c) => c.path === "/api/providers/deepseek" && c.method === "PUT");
  assert.ok(put, "PUT /api/providers/deepseek émis");
  assert.equal(put.body.key, "sk-test-123");
  const row = document.getElementById("apikey-masked-deepseek");
  assert.notEqual(row.style.display, "none");
});

test("décocher un modèle : PUT /api/catalog/selection (debounce 800ms)", async () => {
  // Réinitialise les filtres posés par les tests précédents.
  const cat0 = document.getElementById("catalog-cat-openrouter");
  cat0.value = "all";
  cat0.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  const search0 = document.getElementById("catalog-search-openrouter");
  search0.value = "";
  search0.dispatchEvent(new dom.window.Event("input", { bubbles: true }));
  const cb = document.querySelector('#catalog-list-openrouter .catalog-cb[data-id="openai/gpt-4o"]');
  assert.ok(cb.checked);
  cb.checked = false;
  cb.dispatchEvent(new dom.window.Event("change", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 1100));
  const put = calls.filter((c) => c.path === "/api/catalog/selection" && c.method === "PUT").pop();
  assert.ok(put, "PUT selection émis");
  assert.deepEqual(put.body.disabled.openrouter, ["openai/gpt-4o"]);
});

test("local : Ollama affiche l'URL serveur, Mettre à jour émet PUT", async () => {
  const input = document.getElementById("apikey-ollama");
  assert.equal(input.value, "http://192.168.1.10:11434");
  input.value = "http://192.168.1.20:11434";
  const btn = document.getElementById("apikey-update-ollama");
  btn.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 50));
  const put = calls.find((c) => c.path === "/api/providers/ollama" && c.method === "PUT");
  assert.ok(put, "PUT /api/providers/ollama émis");
  assert.equal(put.body.url, "http://192.168.1.20:11434");
  const status = document.getElementById("apikey-status-ollama");
  assert.ok(status.classList.contains("success"));
});

test("Sauvegarder tout : flush clés + sélection", async () => {
  const input = document.getElementById("apikey-opencode");
  input.value = "oc-key";
  const btn = document.getElementById("apimodeles-save-btn");
  btn.dispatchEvent(new dom.window.Event("click", { bubbles: true }));
  await new Promise((r) => setTimeout(r, 80));
  const put = calls.find((c) => c.path === "/api/providers/opencode" && c.method === "PUT");
  assert.ok(put);
  assert.equal(put.body.key, "oc-key");
  const sel = calls.filter((c) => c.path === "/api/catalog/selection" && c.method === "PUT").pop();
  assert.ok(sel);
});
