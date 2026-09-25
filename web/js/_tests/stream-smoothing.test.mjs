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

function makeView(extra = {}) {
  const log = document.createElement("div");
  document.body.appendChild(log);
  const view = new ThreadView({
    log,
    streamURL: () => "/x",
    sendURL: "/x",
    stopURL: "/x",
    getPayload: (t) => ({ message: t }),
    actions: false,
    ...extra,
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

test("libelle rotatif : variante + compteur, le nom du modele n'y est plus", () => {
  const view = makeView();
  view.handleEvent({ user: "salut" });
  assert.ok(view.waitEl, "loader d'attente affiche");
  assert.match(view.waitTextEl.textContent, /… \d+s$/);
  view.handleEvent({ route: { provider: "llamacpp", label: "Qwen3.5", model: "m", local: true } });
  assert.doesNotMatch(view.waitTextEl.textContent, /Qwen3\.5/);
  // Premier token : le loader disparait.
  view.handleEvent({ content: "voila" });
  assert.equal(view.waitEl, null);
  view.handleEvent({ turn_done: { elapsed_ms: 3 } });
  document.body.innerHTML = "";
});

test("rotation : le message change apres ~1,8 s, sans repetition immediate", async () => {
  const view = makeView({ waitRotateMs: 40 });
  view.handleEvent({ user: "salut" });
  const first = view.waitMsg;
  const seen = new Set([first]);
  for (let i = 0; i < 5; i++) {
    await sleep(60);
    seen.add(view.waitMsg);
  }
  assert.ok(seen.size >= 3, "les messages doivent tourner (" + seen.size + ")");
  view.handleEvent({ content: "x" });
  assert.equal(view.waitRotateTimer, 0, "rotation arretee au premier token");
  document.body.innerHTML = "";
});

test("rotation active même en reduced-motion (seul l'anneau est calmé)", async () => {
  const view = makeView({ reducedMotion: true, waitRotateMs: 20 });
  view.handleEvent({ user: "salut" });
  assert.ok(view.waitRotateTimer, "rotation du texte active");
  const msg = view.waitMsg;
  await sleep(80);
  assert.notEqual(view.waitMsg, msg, "le message doit tourner");
  view.handleEvent({ content: "x" });
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

test("envoi : l'indicateur part avec le POST, pas avec le flux", async () => {
  const view = makeView();
  let release;
  const gate = new Promise((res) => { release = res; });
  const realFetch = globalThis.fetch;
  globalThis.fetch = () => gate.then(() => ({
    ok: true, status: 200, text: async () => "{}",
  }));
  const done = view.sendText("bonjour");
  // Immédiat (avant toute réponse réseau) : spinner + phase Réflexion.
  assert.ok(view.waitEl, "spinner affiche des l'envoi");
  // Message rotatif (une des variantes) + compteur reel.
  assert.match(view.waitTextEl.textContent, /… \d+s$/);
  release();
  await done;
  // POST réussi, flux pas encore ouvert : l'indicateur reste.
  assert.ok(view.waitEl, "indicateur maintenu en attendant le flux");
  // Le nom du modele n'apparait plus dans l'indicateur (rotatif seul).
  const before = view.waitTextEl.textContent;
  view.handleEvent({ route: { provider: "llamacpp", label: "Qwen3.5", model: "m", local: true } });
  assert.equal(view.waitTextEl.textContent, before);
  assert.doesNotMatch(view.waitTextEl.textContent, /Qwen3\.5/);
  view.handleEvent({ content: "voila" });
  assert.equal(view.waitEl, null);
  view.handleEvent({ turn_done: { elapsed_ms: 1 } });
  globalThis.fetch = realFetch;
  document.body.innerHTML = "";
});

test("envoi : erreur réseau => indicateur retiré, jamais bloqué", async () => {
  const view = makeView();
  const realFetch = globalThis.fetch;
  globalThis.fetch = async () => { throw new TypeError("network down"); };
  const ok = await view.sendText("salut");
  assert.equal(ok, false);
  assert.equal(view.waitEl, null, "spinner retiré après échec");
  assert.ok(view.lastSendError);
  globalThis.fetch = realFetch;
  document.body.innerHTML = "";
});

test("timeout client : POST qui pend => envoi coupe, erreur visible", async () => {
  const view = makeView({ sendTimeoutMs: 30 });
  const realFetch = globalThis.fetch;
  globalThis.fetch = (url, opts) =>
    new Promise((_, reject) => {
      if (opts && opts.signal)
        opts.signal.addEventListener("abort", () => {
          const e = new Error("aborted");
          e.name = "AbortError";
          reject(e);
        });
    });
  const done = view.sendText("salut");
  await sleep(100); // tient la boucle : seul le timer d'abort est unrefernce
  const ok = await done;
  assert.equal(ok, false);
  assert.equal(view.waitEl, null, "spinner retire");
  assert.match(view.lastSendError, /ne répond pas/);
  globalThis.fetch = realFetch;
  document.body.innerHTML = "";
});

test("watchdog : POST ok mais flux muet => erreur visible et envoi débloqué", async () => {
  const view = makeView({ turnWatchdogMs: 30 });
  const realFetch = globalThis.fetch;
  globalThis.fetch = async () => ({ ok: true, status: 200, text: async () => "{}" });
  const done = view.sendText("salut");
  await sleep(20);
  const ok = await done;
  assert.equal(ok, true);
  assert.ok(view.generating);
  await sleep(80);
  assert.equal(view.generating, false, "etat de blocage libere");
  assert.equal(view.waitEl, null);
  assert.match(view.lastSendError, /pas démarré/);
  globalThis.fetch = realFetch;
  document.body.innerHTML = "";
});

test("watchdog : un evenement SSE le desarme (pas de fausse alerte)", async () => {
  const view = makeView({ turnWatchdogMs: 40 });
  const realFetch = globalThis.fetch;
  globalThis.fetch = async () => ({ ok: true, status: 200, text: async () => "{}" });
  const done = view.sendText("salut");
  await sleep(20);
  await done;
  view.handleEvent({ seq: 1, user: "salut" });
  await sleep(80);
  assert.ok(view.generating, "le tour demarre ne doit pas declencher l'erreur");
  globalThis.fetch = realFetch;
  document.body.innerHTML = "";
});
