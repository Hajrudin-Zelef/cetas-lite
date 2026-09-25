// Lot lissage + phases (thread-view) : tampon rendu à cadence constante,
// flush intégral en fin de tour, libellé d'attente alimenté par l'événement
// SSE "route". Rien n'est simulé : le texte integral part au flush.
import test from "node:test";
import assert from "node:assert/strict";
import { JSDOM } from "jsdom";

const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.cancelAnimationFrame = dom.window.cancelAnimationFrame.bind(dom.window);
globalThis.getComputedStyle = dom.window.getComputedStyle.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;

const { ThreadView } = await import("../thread-view.js");

const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

function makeView() {
  const log = document.createElement("div");
  document.body.appendChild(log);
  const view = new ThreadView({
    log,
    streamURL: () => "/x",
    sendURL: "/x",
    stopURL: "/x",
    getPayload: (t) => ({ message: t }),
    actions: false,
  });
  return view;
}

test("lissage : le texte integrale est affiche apres turn_done (rien perdu)", async () => {
  const view = makeView();
  view.handleEvent({ user: "salut" });
  // Rafale en plusieurs deltas (comme un SSE saccade).
  view.handleEvent({ content: "bon" });
  view.handleEvent({ content: "jour le " });
  view.handleEvent({ content: "monde" });
  // Tout de suite : le DOM peut être en retard (tampon), assistantText non.
  assert.equal(view.assistantText, "bonjour le monde");
  // flush (finalizeAssistant, appele par finishTurn avant le reset du fil)
  view.finalizeAssistant();
  assert.ok(view.assistantBody.textContent.includes("bonjour le monde"));
  document.body.innerHTML = "";
});

test("lissage : le tampon rend progressivement (pas de rendu instantane obligatoire)", async () => {
  const view = makeView();
  view.handleEvent({ user: "salut" });
  view.handleEvent({ content: "x".repeat(600) });
  // Un rAF apres : le rendu a demarre mais n'est pas force a tout montrer
  // (catch-up len/8) ; l'important est que assistantText soit complet.
  await sleep(50);
  assert.equal(view.assistantText.length, 600);
  view.finalizeAssistant();
  assert.ok(view.assistantBody.textContent.includes("xxxxx"));
  document.body.innerHTML = "";
});

test("phase d'attente : route met a jour le libelle avant le premier token", () => {
  const view = makeView();
  view.handleEvent({ user: "salut" });
  assert.ok(view.waitEl, "loader d'attente affiche");
  assert.match(view.waitLabel.textContent, /En cours/);
  view.handleEvent({ route: { provider: "llamacpp", label: "Qwen3.5", model: "m", local: true } });
  assert.match(view.waitLabel.textContent, /Qwen3.5 · rédige/);
  // Premier token : le loader disparait.
  view.handleEvent({ content: "voila" });
  assert.equal(view.waitEl, null);
  view.handleEvent({ turn_done: { elapsed_ms: 3 } });
  document.body.innerHTML = "";
});

test("replace (replay) synchronise le tampon : les deltas suivants ne perdent rien", async () => {
  const view = makeView();
  view.handleEvent({ user: "q" });
  view.handleEvent({ content: "debut ", replace: true });
  view.handleEvent({ content: "suite" });
  view.finalizeAssistant();
  assert.equal(view.assistantText, "debut suite");
  assert.ok(view.assistantBody.textContent.includes("debut suite"));
  document.body.innerHTML = "";
});
