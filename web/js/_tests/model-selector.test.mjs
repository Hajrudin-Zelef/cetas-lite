import test from "node:test";
import assert from "node:assert/strict";
import { fileURLToPath } from "node:url";
import { dirname, join } from "node:path";

const ROOT = join(dirname(fileURLToPath(import.meta.url)), "..", "..");

const {
  poolKey,
  sameKeySet,
  modeSubtitle,
  familyShortLabel,
  engineName,
  buildStaged,
  fmtPrice,
} = await import(ROOT + "/js/model-selector.js");

test("poolKey : provider/model", () => {
  assert.equal(poolKey({ provider: "deepseek", model: "deepseek-chat" }), "deepseek/deepseek-chat");
});

test("sameKeySet : ordre indifférent, doublons détectés", () => {
  const a = [{ provider: "p", model: "m1" }, { provider: "p", model: "m2" }];
  const b = [{ provider: "p", model: "m2" }, { provider: "p", model: "m1" }];
  assert.ok(sameKeySet(a, b));
  assert.ok(!sameKeySet(a, [{ provider: "p", model: "m1" }]));
  assert.ok(!sameKeySet(a, [{ provider: "p", model: "m1" }, { provider: "p", model: "m3" }]));
});

test("modeSubtitle : 1 modèle -> label, N -> Fallback · N, jamais de règle technique", () => {
  const fam = { id: "samagent-n8" };
  assert.equal(
    modeSubtitle(fam, { pool: [{ provider: "deepseek", model: "deepseek-chat", label: "DeepSeek V3.2" }] }),
    "DeepSeek V3.2"
  );
  assert.equal(
    modeSubtitle(fam, { pool: [{ provider: "a", model: "x" }, { provider: "b", model: "y" }] }),
    "Fallback · 2 modèles"
  );
  assert.equal(modeSubtitle(fam, { pool: [], rule: "deepseek v3.2" }), "");
  // SamGen : route par défaut, sélection locale sinon.
  const sg = { id: "samgen", local: true };
  assert.equal(modeSubtitle(sg, { pool: [], engine: "ollama" }), "Ollama");
  assert.equal(
    modeSubtitle(sg, { pool: [{ provider: "ollama", model: "llama3.1:8b", label: "llama3.1:8b" }] }),
    "llama3.1:8b"
  );
  assert.equal(
    modeSubtitle(sg, { pool: [{ provider: "ollama", model: "a" }, { provider: "ollama", model: "b" }], engine: "ollama" }),
    "Fallback · 2 modèles"
  );
});

test("familyShortLabel : alias courts", () => {
  assert.equal(familyShortLabel({ id: "samagent-nano" }), "Nano");
  assert.equal(familyShortLabel({ id: "samagent-n4" }), "N4");
  assert.equal(familyShortLabel({ id: "samagent-n8" }), "N8");
  assert.equal(familyShortLabel({ id: "code" }), "Code");
  assert.equal(familyShortLabel({ id: "samgen" }), "SamGen");
  assert.equal(engineName("llamacpp"), "llama.cpp");
  assert.equal(engineName("lmstudio"), "LM Studio");
});

test("buildStaged : reconstruit les overrides depuis effectif vs défaut", () => {
  const defaults = [
    { id: "samagent-n4", modes: [{ mode: "flash", pool: [{ provider: "openrouter", model: "a" }, { provider: "openrouter", model: "b" }] }] },
    { id: "samgen", local: true, modes: [{ mode: "n4", pool: [] }] },
  ];
  const families = [
    // identique au défaut -> pas d'override
    { id: "samagent-n4", modes: [{ mode: "flash", pool: [{ provider: "openrouter", model: "b" }, { provider: "openrouter", model: "a" }] }] },
    // sélection locale -> override
    { id: "samgen", local: true, modes: [{ mode: "n4", pool: [{ provider: "ollama", model: "qwen2.5:14b", label: "qwen2.5:14b" }] }] },
  ];
  const staged = buildStaged(families, defaults);
  assert.deepEqual(staged, {
    samgen: { n4: [{ provider: "ollama", model: "qwen2.5:14b" }] },
  });
});

test("buildStaged : pool réduit -> override conservé", () => {
  const defaults = [
    { id: "samagent-n8", modes: [{ mode: "standard", pool: [{ provider: "deepseek", model: "deepseek-flash" }, { provider: "opencode-go", model: "hy3-go" }] }] },
  ];
  const families = [
    { id: "samagent-n8", modes: [{ mode: "standard", pool: [{ provider: "deepseek", model: "deepseek-flash", label: "DeepSeek V4.1 Flash" }] }] },
  ];
  const staged = buildStaged(families, defaults);
  assert.deepEqual(staged, {
    "samagent-n8": { standard: [{ provider: "deepseek", model: "deepseek-flash" }] },
  });
});

test("fmtPrice : gratuit et prix", () => {
  assert.equal(fmtPrice({ input_per_1m: 0, output_per_1m: 0 }), "gratuit");
  assert.ok(fmtPrice({ input_per_1m: 0.14, output_per_1m: 0.28 }).includes("/1M"));
});

test("modals.js importe bien loadSelectorPanel depuis model-selector.js", async () => {
  const fs = await import("node:fs");
  const src = fs.readFileSync(new URL("../modals.js", import.meta.url), "utf8");
  const m = src.match(/import\s*\{([^}]+)\}\s*from\s*"\.\/model-selector\.js"/);
  assert.ok(m, "import de model-selector.js trouvé dans modals.js");
  assert.ok(m[1].split(",").map((s) => s.trim()).includes("loadSelectorPanel"));
  assert.ok(m[1].split(",").map((s) => s.trim()).includes("loadAgentSelectorPanel"));
  assert.ok(src.includes('tab.dataset.tab === "selector"'), "onglet selector câblé");
  assert.ok(src.includes('tab.dataset.tab === "selector-agent"'), "onglet selector-agent câblé");
});

test("index.html : onglet + panneau sélecteur présents", async () => {
  const fs = await import("node:fs");
  const html = fs.readFileSync(new URL("../../index.html", import.meta.url), "utf8");
  assert.ok(html.includes('data-tab="selector"'), "bouton onglet sélecteur");
  assert.ok(html.includes('id="panel-selector"'), "panneau sélecteur");
  assert.ok(html.includes('id="selector-body"'), "conteneur selector-body");
  assert.ok(html.includes("mx-model-selector.css"), "CSS du sélecteur chargé");
});

test("index.html : onglet + panneau sélecteur agent présents", async () => {
  const fs = await import("node:fs");
  const html = fs.readFileSync(new URL("../../index.html", import.meta.url), "utf8");
  assert.ok(html.includes('data-tab="selector-agent"'), "bouton onglet sélecteur agent");
  assert.ok(html.includes('id="panel-selector-agent"'), "panneau sélecteur agent");
  assert.ok(html.includes('id="selector-agent-body"'), "conteneur selector-agent-body");
  assert.ok(html.includes('id="selector-agent-save-state"'), "indicateur de sauvegarde agent");
});

test("model-selector.js exporte loadAgentSelectorPanel", async () => {
  const mod = await import("../model-selector.js");
  assert.equal(typeof mod.loadAgentSelectorPanel, "function");
  assert.equal(typeof mod.loadSelectorPanel, "function");
});
