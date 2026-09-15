import test from "node:test";
import assert from "node:assert/strict";
// jsdom : dependance de test declaree dans package.json (npm install).
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

const nextFrame = () =>
  new Promise((resolve) => {
    dom.window.requestAnimationFrame(() => resolve());
  });
const sleep = (ms) => new Promise((r) => setTimeout(r, ms));

// Géométrie simulée : jsdom ne fait pas de layout.
function makeView() {
  const parent = document.createElement("div");
  const log = document.createElement("div");
  parent.appendChild(log);
  document.body.appendChild(parent);
  const geom = { sh: 2000, ch: 600 };
  let top = 1400; // = floor -> épinglé au départ
  Object.defineProperty(log, "scrollHeight", { get: () => geom.sh, configurable: true });
  Object.defineProperty(log, "clientHeight", { get: () => geom.ch, configurable: true });
  Object.defineProperty(log, "scrollTop", {
    get: () => top,
    set: (v) => { top = v; },
    configurable: true,
  });
  const view = new ThreadView({
    log,
    streamURL: () => "/x",
    sendURL: "/x",
    stopURL: "/x",
    getPayload: (t) => ({ message: t }),
    actions: false,
  });
  return { view, log, parent, geom, getTop: () => top, setTop: (v) => { top = v; } };
}
const floorOf = (geom) => geom.sh - geom.ch;

test("état initial : épinglé, requestFollow écrit en bas sur une frame", async () => {
  const { view, log, geom, getTop } = makeView();
  assert.equal(view.pinned, true);
  view.requestFollow();
  await nextFrame();
  await nextFrame();
  assert.equal(getTop(), geom.sh);
  assert.equal(view.observedTop, geom.sh);
  document.body.innerHTML = "";
});

test("geste lecteur vers le haut : désépingle après échantillonnage, bouton visible", async () => {
  const { view, log, geom, setTop } = makeView();
  setTop(floorOf(geom) - 400); // le lecteur remonte
  log.dispatchEvent(new dom.window.Event("scroll"));
  await sleep(650); // SCROLL_SAMPLE_MS = 500
  assert.equal(view.pinned, false);
  assert.equal(view.toBottomBtn.hidden, false);
  document.body.innerHTML = "";
});

test("requestFollow non épinglé : ne scrolle pas, signale le nouveau contenu", async () => {
  const { view, log, geom, getTop, setTop } = makeView();
  setTop(floorOf(geom) - 400);
  log.dispatchEvent(new dom.window.Event("scroll"));
  await sleep(650);
  const before = getTop();
  view.requestFollow();
  await nextFrame();
  await nextFrame();
  assert.equal(getTop(), before); // pas de yank
  assert.equal(view.newWhileUnpinned, true);
  assert.ok(view.toBottomBtn.classList.contains("has-new"));
  document.body.innerHTML = "";
});

test("livraison programmatique différée : ne désépingle pas", async () => {
  const { view, log, geom, setTop } = makeView();
  // writeBottom simulé : observedTop synchronisé, le navigateur délivre après
  setTop(geom.sh);
  view.observedTop = geom.sh;
  assert.equal(view.pinned, true);
  log.dispatchEvent(new dom.window.Event("scroll")); // branche pinned -> sample immédiat
  assert.equal(view.pinned, true);
  assert.equal(view.toBottomBtn.hidden, true);
  document.body.innerHTML = "";
});

test("seuil d'épinglage : 24px comme la référence", async () => {
  const { view, log, geom, setTop } = makeView();
  const floor = floorOf(geom);
  setTop(floor - 10);
  log.dispatchEvent(new dom.window.Event("scroll"));
  await sleep(650);
  assert.equal(view.pinned, true); // dans les 24px -> épinglé
  setTop(floor - 30);
  log.dispatchEvent(new dom.window.Event("scroll"));
  await sleep(650);
  assert.equal(view.pinned, false); // au-delà -> lecture libre
  document.body.innerHTML = "";
});

test("retour en bas : le bouton ré-épingle et masque le bouton", async () => {
  const { view, log, geom, getTop, setTop } = makeView();
  setTop(floorOf(geom) - 400);
  log.dispatchEvent(new dom.window.Event("scroll"));
  await sleep(650);
  assert.equal(view.pinned, false);
  view.toBottomBtn.click();
  assert.equal(getTop(), geom.sh);
  assert.equal(view.pinned, true);
  assert.equal(view.toBottomBtn.hidden, true);
  assert.equal(view.newWhileUnpinned, false);
  document.body.innerHTML = "";
});

test("élagage : borne le DOM, protège approbations et tool boxes en cours", () => {
  const { view, log } = makeView();
  for (let i = 0; i < 520; i++) {
    const d = document.createElement("div");
    d.className = "msg-system";
    d.textContent = "m" + i;
    log.appendChild(d);
  }
  // Nœuds protégés placés dans la zone élagable (en tête).
  const approval = document.createElement("div");
  approval.className = "msg-approval";
  log.insertBefore(approval, log.children[10]);
  const toolbox = document.createElement("div");
  toolbox.className = "tool-box streaming";
  log.insertBefore(toolbox, log.children[20]);
  view.maybePrune();
  // 500 nœuds de contenu max + 1 sentinelle .history-cut
  assert.ok(log.children.length <= 501, "borne respectée : " + log.children.length);
  assert.ok(log.querySelector(":scope > .history-cut"), "sentinelle présente");
  assert.ok(log.contains(approval), "approbation en attente conservée");
  assert.ok(log.contains(toolbox), "tool box en streaming conservée");
  assert.ok(view.prunedCount > 0);
  document.body.innerHTML = "";
});

test("reset : vide l'état de scroll et l'élagage", () => {
  const { view, log, geom, setTop } = makeView();
  for (let i = 0; i < 520; i++) log.appendChild(document.createElement("div"));
  view.maybePrune();
  assert.ok(view.prunedCount > 0);
  setTop(100);
  view.pinned = false;
  view.reset();
  assert.equal(view.pinned, true);
  assert.equal(view.observedTop, 0);
  assert.equal(view.prunedCount, 0);
  assert.equal(log.querySelector(":scope > .history-cut"), null);
  assert.equal(view.newWhileUnpinned, false);
  document.body.innerHTML = "";
});
