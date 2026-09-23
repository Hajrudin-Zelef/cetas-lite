// Logique pure des questions suggérées (itération 4) : pas de DOM,
// pas de jsdom requis.
import { test, describe } from "node:test";
import assert from "node:assert/strict";
import { isGreeting, drawSuggestions, keyOf, readSeen, markSeen } from "../suggest.js";

function memStorage() {
  const m = new Map();
  return {
    getItem: (k) => (m.has(k) ? m.get(k) : null),
    setItem: (k, v) => m.set(k, String(v)),
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

describe("isGreeting", () => {
  test("salutations reconnues", () => {
    for (const s of ["salut", "Salut !", "  bonjour ", "BONSOIR.", "hello", "Hi?", "coucou…", "yo !"])
      assert.equal(isGreeting(s), true, s);
  });
  test("non-salutations rejetées", () => {
    for (const s of ["", "salut, ça va ?", "bonjour à tous", "salu", "salutation", "ok", "merci", "salut salut"])
      assert.equal(isGreeting(s), false, JSON.stringify(s));
  });
});

describe("drawSuggestions", () => {
  test("3 chips par défaut", () => {
    const d = drawSuggestions(pool);
    assert.equal(d.length, 3);
  });
  test("épinglés en priorité", () => {
    for (let i = 0; i < 20; i++) {
      const d = drawSuggestions(pool);
      assert.ok(d.some((s) => s.pinned), "au moins un épinglé attendu");
    }
  });
  test("mixité des corpus", () => {
    for (let i = 0; i < 30; i++) {
      const d = drawSuggestions(pool);
      const corpora = new Set(d.map((s) => s.corpus));
      assert.equal(corpora.size, d.length, "corpus dupliqué dans le tirage");
    }
  });
  test("exclusion anti-répétition", () => {
    const d = drawSuggestions(pool, { count: 6, exclude: [keyOf(pool[0]), keyOf(pool[1])] });
    assert.ok(!d.some((s) => s.question === "Q Kimi" || s.question === "Q vLLM"));
  });
  test("pool vide ou count 0", () => {
    assert.deepEqual(drawSuggestions([], {}), []);
    assert.deepEqual(drawSuggestions(pool, { count: 0 }), []);
  });
  test("pool homogène : on complète quand même", () => {
    const same = [
      { question: "A", corpus: "c" },
      { question: "B", corpus: "c" },
      { question: "C", corpus: "c" },
    ];
    assert.equal(drawSuggestions(same, { count: 3 }).length, 3);
  });
});

describe("anti-répétition (seen)", () => {
  test("markSeen puis readSeen", () => {
    const st = memStorage();
    assert.deepEqual(readSeen(st), []);
    markSeen([pool[0], pool[1]], st);
    const seen = readSeen(st);
    assert.equal(seen.length, 2);
    assert.ok(seen.includes(keyOf(pool[0])));
    // Pas de doublon au second marquage.
    markSeen([pool[0]], st);
    assert.equal(readSeen(st).length, 2);
  });
  test("plafond de 15 entrées", () => {
    const st = memStorage();
    const many = Array.from({ length: 25 }, (_, i) => ({ question: "Q" + i, corpus: "c" }));
    markSeen(many, st);
    assert.ok(readSeen(st).length <= 15);
  });
  test("JSON corrompu => []", () => {
    const st = memStorage();
    st.setItem("cetas-lite-suggest-seen", "pas du json{");
    assert.deepEqual(readSeen(st), []);
  });
});
