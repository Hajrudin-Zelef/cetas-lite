// Module DeepThink Global : rendu de l'onglet paramètres et persistance.
import { test, describe } from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

// jsdom : résolution standard puis repli /tmp.
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

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.CustomEvent = dom.window.CustomEvent;
globalThis.Event = dom.window.Event;
globalThis.localStorage = dom.window.localStorage;

const putCalls = [];
// Mock fetch avec text() (api() lit resp.text()).
globalThis.fetch = async (url, opts = {}) => {
  const method = opts.method || "GET";
  if (String(url) === "/api/deepthink" && method === "GET") {
    return {
      ok: true,
      status: 200,
      text: async () =>
        JSON.stringify({
          settings: { lang: "fr", provider: "openrouter", model: "openrouter/free" },
          langs: ["es", "fr", "it", "de", "en"],
          models: [
            { provider: "openrouter", model: "openrouter/free", label: "openrouter/free (OpenRouter)" },
            { provider: "deepseek", model: "deepseek-chat", label: "DeepSeek V3.2 (DeepSeek)" },
          ],
        }),
    };
  }
  if (String(url) === "/api/deepthink" && method === "PUT") {
    const body = JSON.parse(opts.body || "{}");
    putCalls.push(body);
    return { ok: true, status: 200, text: async () => JSON.stringify({ ok: true, settings: body }) };
  }
  return { ok: false, status: 404, text: async () => JSON.stringify({ error: "nope" }) };
};

const mod = await import("../deepthink.js");

const wait = (ms) => new Promise((r) => setTimeout(r, ms));

function setupDom() {
  document.body.innerHTML = `
    <div class="apikeys-panel" id="panel-deepthink">
      <h3>DeepThink Global <span id="deepthink-save-state" class="ms-save-state"></span></h3>
      <div class="apikeys-panel-body"><div id="deepthink-body"></div></div>
    </div>`;
}

describe("DeepThink Global : panneau paramètres", () => {
  test("rendu : 5 langues + 2 modèles, FR actif par défaut", async () => {
    setupDom();
    await mod.loadDeepThinkPanel();
    const body = document.getElementById("deepthink-body");
    const segs = body.querySelectorAll(".ms-seg");
    assert.equal(segs.length, 2, "2 groupes (langue + modèle)");
    const langBtns = segs[0].querySelectorAll(".ms-seg-btn");
    assert.equal(langBtns.length, 5, "5 langues");
    assert.deepEqual(
      [...langBtns].map((b) => b.textContent),
      ["Espagnol", "Français", "Italien", "Allemand", "Anglais"]
    );
    const activeLang = segs[0].querySelector(".ms-seg-btn.active");
    assert.equal(activeLang.textContent, "Français", "FR actif par défaut");
    const modelBtns = segs[1].querySelectorAll(".ms-seg-btn");
    assert.equal(modelBtns.length, 2, "2 modèles traducteurs");
    assert.ok(modelBtns[0].textContent.includes("openrouter/free"));
    assert.ok(modelBtns[1].textContent.includes("DeepSeek V3.2"));
    const activeModel = segs[1].querySelector(".ms-seg-btn.active");
    assert.ok(activeModel.textContent.includes("openrouter/free"), "openrouter/free actif par défaut");
  });

  test("clic langue -> PUT /api/deepthink avec la nouvelle langue", async () => {
    setupDom();
    putCalls.length = 0;
    await mod.loadDeepThinkPanel();
    const esBtn = [...document.querySelectorAll(".ms-seg-btn")].find((b) => b.textContent === "Espagnol");
    assert.ok(esBtn, "bouton Espagnol trouvé");
    esBtn.click();
    assert.ok(esBtn.classList.contains("active"), "bouton activé immédiatement");
    await wait(500); // debounce 350ms
    assert.equal(putCalls.length, 1, "un PUT émis");
    assert.equal(putCalls[0].lang, "es", "langue persistée = es");
    assert.equal(document.getElementById("deepthink-save-state").textContent, "Enregistré ✓");
  });

  test("clic modèle -> PUT avec provider/model DeepSeek", async () => {
    setupDom();
    putCalls.length = 0;
    await mod.loadDeepThinkPanel();
    const dsBtn = [...document.querySelectorAll(".ms-seg-btn")].find((b) =>
      b.textContent.includes("DeepSeek V3.2")
    );
    assert.ok(dsBtn, "bouton DeepSeek V3.2 trouvé");
    dsBtn.click();
    await wait(500);
    assert.equal(putCalls.length, 1, "un PUT émis");
    assert.equal(putCalls[0].provider, "deepseek");
    assert.equal(putCalls[0].model, "deepseek-chat");
  });

  test("exports : loadDeepThinkPanel est une fonction (import modals.js)", async () => {
    assert.equal(typeof mod.loadDeepThinkPanel, "function");
  });
});
