import test from "node:test";
import assert from "node:assert/strict";
// jsdom : dependance de test declaree dans package.json (npm install).
import { JSDOM } from "jsdom";

// jsdom (meme version que le harnais agents) pour le DOM.
const dom = new JSDOM("<!doctype html><html><body></body></html>", {
  pretendToBeVisual: true,
  url: "http://localhost/",
});
globalThis.window = dom.window;
globalThis.document = dom.window.document;
globalThis.requestAnimationFrame = dom.window.requestAnimationFrame.bind(dom.window);
globalThis.cancelAnimationFrame = dom.window.cancelAnimationFrame.bind(dom.window);
globalThis.localStorage = dom.window.localStorage;

const { splitBlocks, createMarkdownRenderer } = await import("../markdown.js");

const nextFrame = () =>
  new Promise((resolve) => {
    dom.window.requestAnimationFrame(() => resolve());
  });

test("splitBlocks decoupe sur les lignes vides", () => {
  const { blocks, tail } = splitBlocks("aaa\n\nbbb\n\n");
  assert.deepEqual(blocks.map((b) => b.trim()), ["aaa", "bbb"]);
  assert.equal(tail, "");
});

test("splitBlocks ne coupe pas a l'interieur d'un fence", () => {
  const { blocks, tail } = splitBlocks("```js\n\na = 1\n```\n\ntexte");
  assert.equal(blocks.length, 1);
  assert.ok(blocks[0].includes("a = 1"));
  assert.equal(tail, "texte");
});

test("splitBlocks garde le tail live", () => {
  const { blocks, tail } = splitBlocks("aaa\n\nbb");
  assert.deepEqual(blocks.map((b) => b.trim()), ["aaa"]);
  assert.equal(tail, "bb");
});

test("splitBlocks gere le tilde fence", () => {
  const { blocks, tail } = splitBlocks("~~~\n\ncode\n~~~\n\nfin");
  assert.equal(blocks.length, 1);
  assert.equal(tail, "fin");
});

test("update coalesce les deltas sur une frame", async () => {
  const r = createMarkdownRenderer();
  const c = document.createElement("div");
  document.body.appendChild(c);
  r.update(c, "premier");
  r.update(c, "premier deux");
  r.update(c, "premier deux trois");
  // Pas encore de rendu synchrone : tout est bufferise.
  assert.equal(c.children.length, 0);
  await nextFrame();
  assert.ok(c.textContent.includes("premier deux trois"));
  assert.equal(r.text(c), "premier deux trois");
  c.remove();
});

test("update ne re-rend pas les blocs stables", async () => {
  const r = createMarkdownRenderer();
  const c = document.createElement("div");
  document.body.appendChild(c);
  r.update(c, "bloc stable\n\ntail un");
  await nextFrame();
  assert.equal(c.children.length, 2);
  const stableHTML = c.children[0].innerHTML;
  // Nouveau delta qui ne touche que le tail.
  r.update(c, "bloc stable\n\ntail deux !");
  await nextFrame();
  assert.equal(c.children.length, 2);
  assert.equal(c.children[0].innerHTML, stableHTML);
  assert.ok(c.children[1].textContent.includes("tail deux !"));
  c.remove();
});

test("render remplace immediatement et synchrone", () => {
  const r = createMarkdownRenderer();
  const c = document.createElement("div");
  document.body.appendChild(c);
  r.render(c, "abc");
  assert.ok(c.textContent.includes("abc"));
  r.render(c, "xyz");
  assert.ok(!c.textContent.includes("abc"));
  assert.ok(c.textContent.includes("xyz"));
  assert.equal(r.text(c), "xyz");
  c.remove();
});

test("finalize rend synchrone et annule la frame en attente", async () => {
  const r = createMarkdownRenderer();
  const c = document.createElement("div");
  document.body.appendChild(c);
  r.update(c, "brouillon");
  r.finalize(c, "fini");
  assert.ok(c.textContent.includes("fini"));
  await nextFrame();
  await nextFrame();
  assert.ok(c.textContent.includes("fini"));
  assert.ok(!c.textContent.includes("brouillon"));
  c.remove();
});

test("finalize sans texte reprend le buffer", async () => {
  const r = createMarkdownRenderer();
  const c = document.createElement("div");
  document.body.appendChild(c);
  r.update(c, "en cours");
  await nextFrame();
  r.finalize(c);
  assert.ok(c.textContent.includes("en cours"));
  c.remove();
});
