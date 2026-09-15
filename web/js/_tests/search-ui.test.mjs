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

function makeView() {
  const parent = document.createElement("div");
  const log = document.createElement("div");
  parent.appendChild(log);
  document.body.appendChild(parent);
  const view = new ThreadView({
    log,
    streamURL: () => "/x",
    sendURL: "/x",
    stopURL: "/x",
    getPayload: (t) => ({ message: t }),
    actions: false,
  });
  return { view, log };
}
const done = () => { document.body.innerHTML = ""; };

test("web_search : l'indicateur s'affiche au demarrage de l'outil", () => {
  const { view, log } = makeView();
  view.addTool({ name: "web_search", args: { query: "prix du cuivre" }, phase: "start" });
  const box = log.querySelector("details.msg-tool");
  assert.ok(box, "boite d'outil attendue");
  const status = box.querySelector(".search-status");
  assert.ok(status, "indicateur de recherche attendu");
  assert.ok(status.querySelector(".search-spinner"), "spinner attendu");
  assert.match(status.textContent, /prix du cuivre/);
  done();
});

test("web_search : l'indicateur laisse place au panneau Sources", () => {
  const { view, log } = makeView();
  const args = { query: "prix du cuivre" };
  view.addTool({ name: "web_search", args, phase: "start" });
  view.addTool({
    name: "web_search", args, phase: "end",
    result: "Resultats tavily (2):\n[1] T1\nhttps://a.example\n...",
    sources: [
      { title: "Titre A", url: "https://a.example/x" },
      { title: "Titre B", url: "https://www.b.example/y" },
    ],
    search_provider: "tavily",
  });
  const box = log.querySelector("details.msg-tool");
  assert.equal(box.querySelector(".search-status"), null, "indicateur retire");
  const block = box.querySelector(".citations-block");
  assert.ok(block, "panneau Sources attendu");
  assert.equal(block.querySelector(".citations-title").textContent, "Sources");
  const cards = block.querySelectorAll(".citation-card");
  assert.equal(cards.length, 2);
  assert.equal(cards[0].getAttribute("href"), "https://a.example/x");
  assert.equal(cards[0].querySelector(".citation-num").textContent, "[1]");
  assert.equal(cards[0].querySelector(".citation-domain").textContent, "a.example");
  assert.equal(cards[0].querySelector(".citation-card-title").textContent, "Titre A");
  assert.equal(cards[1].querySelector(".citation-domain").textContent, "b.example");
  // Le resultat brut n'est pas duplique quand le panneau existe.
  assert.equal(box.querySelector("pre.tool-result"), null);
  done();
});

test("web_search : sans sources, le resultat texte reste affiche", () => {
  const { view, log } = makeView();
  const args = { query: "zzz" };
  view.addTool({ name: "web_search", args, phase: "start" });
  view.addTool({ name: "web_search", args, phase: "end", result: "[info] recherche web: aucun resultat" });
  const box = log.querySelector("details.msg-tool");
  assert.equal(box.querySelector(".search-status"), null);
  assert.equal(box.querySelector(".citations-block"), null);
  assert.match(box.querySelector("pre.tool-result").textContent, /aucun resultat/);
  done();
});

test("panneau Sources : +N sources deplie le reste", () => {
  const { view, log } = makeView();
  const args = { query: "q" };
  view.addTool({ name: "web_search", args, phase: "start" });
  const sources = Array.from({ length: 6 }, (_, i) => ({ title: "T" + i, url: "https://s" + i + ".example/" }));
  view.addTool({ name: "web_search", args, phase: "end", result: "x", sources });
  const block = log.querySelector(".citations-block");
  assert.equal(block.querySelectorAll("li.citation-hidden").length, 2);
  const more = block.querySelector(".citation-more");
  assert.match(more.textContent, /\+2 sources/);
  more.click();
  assert.equal(block.querySelectorAll("li.citation-hidden").length, 0);
  assert.equal(block.querySelector(".citation-more"), null);
  done();
});

test("recherche native : indicateur puis panneau Sources via ev.search", () => {
  const { view, log } = makeView();
  view.handleEvent({ search: { phase: "start", native: true } });
  const status = log.querySelector(":scope > .search-status");
  assert.ok(status, "indicateur natif attendu");
  assert.ok(status.querySelector(".search-spinner"));
  // Second start : pas de doublon.
  view.handleEvent({ search: { phase: "start", native: true } });
  assert.equal(log.querySelectorAll(":scope > .search-status").length, 1);
  view.handleEvent({
    search: { phase: "done", native: true, sources: [{ title: "N", url: "https://n.example/" }] },
  });
  assert.equal(log.querySelector(":scope > .search-status"), null, "indicateur retire");
  const block = log.querySelector(":scope > .citations-block");
  assert.ok(block, "panneau Sources natif attendu");
  assert.equal(block.querySelector(".citation-card").getAttribute("href"), "https://n.example/");
  done();
});

test("recherche native : done sans sources retire juste l'indicateur", () => {
  const { view, log } = makeView();
  view.handleEvent({ search: { phase: "start", native: true } });
  view.handleEvent({ search: { phase: "done", native: true, sources: [] } });
  assert.equal(log.querySelector(".search-status"), null);
  assert.equal(log.querySelector(".citations-block"), null);
  done();
});
