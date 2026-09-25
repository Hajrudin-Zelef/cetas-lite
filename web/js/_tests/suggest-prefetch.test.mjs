// Lot 5 : le callback onDrawn de fillChips expose les questions tirées
// (pré-génération). DOM/fetch/localStorage stubés, sans jsdom.
import { test, describe, beforeEach } from "node:test";
import assert from "node:assert/strict";

function memStorage() {
  const m = new Map();
  return {
    getItem: (k) => (m.has(k) ? m.get(k) : null),
    setItem: (k, v) => m.set(k, String(v)),
    removeItem: (k) => m.delete(k),
  };
}

const pool = [
  { question: "Q Kimi", corpus: "wave6", pinned: true },
  { question: "Q vLLM", corpus: "vllm", pinned: true },
  { question: "Q CUDA", corpus: "cuda" },
  { question: "Q ROCm", corpus: "amd" },
  { question: "Q BGP", corpus: "net" },
  { question: "Q Proxmox", corpus: "proxmox" },
];

function fakeButton() {
  return {
    type: "", className: "", textContent: "", title: "",
    _listeners: {},
    addEventListener(t, fn) { (this._listeners[t] = this._listeners[t] || []).push(fn); },
    click() { for (const fn of this._listeners.click || []) fn(); },
  };
}

function fakeBox() {
  return {
    innerHTML: "",
    classList: { add() {} },
    children: [],
    appendChild(c) { this.children.push(c); return c; },
  };
}

function moreButton(box) {
  return box.children.find((c) => c.className.includes("suggest-chip-more"));
}

let suggest;
beforeEach(async () => {
  globalThis.localStorage = memStorage();
  globalThis.fetch = async () => ({
    ok: true, status: 200,
    text: async () => JSON.stringify({ suggestions: pool }),
  });
  globalThis.document = { createElement: () => fakeButton() };
  suggest = await import("../suggest.js");
  suggest.clearPoolCache();
});

describe("fillChips onDrawn (lot 5)", () => {
  test("onDrawn reçoit les 3 questions du tirage initial", async () => {
    const box = fakeBox();
    const drawn = [];
    await suggest.fillChips(box, () => {}, (picked) => drawn.push(picked));
    assert.equal(drawn.length, 1, "un seul appel après le tirage initial");
    assert.equal(drawn[0].length, 3);
    for (const s of drawn[0]) {
      assert.ok(s.question && s.corpus, "question + corpus exposés");
    }
  });

  test("« Autre… » renotifie avec le nouveau tirage", async () => {
    const box = fakeBox();
    const drawn = [];
    await suggest.fillChips(box, () => {}, (picked) => drawn.push(picked));
    moreButton(box).click();
    assert.equal(drawn.length, 2, "renotifié après « Autre… »");
    assert.equal(drawn[1].length, 3);
  });

  test("sans onDrawn : aucun crash, comportement inchangé", async () => {
    const box = fakeBox();
    let clicked = null;
    await suggest.fillChips(box, (s) => { clicked = s; });
    assert.equal(box.children.length, 4, "3 chips + Autre…");
    box.children[0].click();
    assert.ok(clicked && clicked.question, "onPick toujours appelé");
  });

  test("onDrawn qui lève ne casse pas l'affichage", async () => {
    const box = fakeBox();
    await suggest.fillChips(box, () => {}, () => { throw new Error("boom"); });
    assert.equal(box.children.length, 4);
  });
});
