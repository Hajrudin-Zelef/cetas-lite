import test from "node:test";
import assert from "node:assert/strict";

import { takeCount, createStreamRenderer } from "../stream-render.js";

test("takeCount renvoie 0 sur buffer vide", () => {
  assert.equal(takeCount(0, 16), 0);
  assert.equal(takeCount(-5, 16), 0);
});

test("takeCount ne depasse jamais la longueur disponible", () => {
  assert.equal(takeCount(3, 1000), 3);
  assert.equal(takeCount(10, 1000), 10);
});

test("takeCount prend base quand dt est grand", () => {
  assert.equal(takeCount(100, 1000, 30), 30);
});

test("takeCount prend catchup quand dt est petit et le buffer long", () => {
  assert.equal(takeCount(100, 16, 30), Math.ceil(100 / 8));
  assert.equal(takeCount(800, 16, 30), Math.ceil(800 / 8));
});

test("takeCount avance d'au moins 1 par frame", () => {
  assert.ok(takeCount(20, 16, 30) >= 1);
});

test("takeCount reste borne par la longueur", () => {
  for (const len of [1, 5, 20, 137, 1000]) {
    const got = takeCount(len, 16, 30);
    assert.ok(got >= 1 && got <= len, `len=${len} got=${got}`);
  }
});

test("renderer replace est immediat et replace-aware", () => {
  const seen = [];
  const r = createStreamRenderer({ onRender: (t) => seen.push(t) });
  r.replace("bonjour");
  assert.equal(r.text(), "bonjour");
  assert.deepEqual(seen, ["bonjour"]);
  r.replace("salut");
  assert.equal(r.text(), "salut");
  assert.deepEqual(seen, ["bonjour", "salut"]);
});

test("renderer replace declenche onFirst une seule fois", () => {
  let first = 0;
  const r = createStreamRenderer({ onRender: () => {}, onFirst: () => { first++; } });
  r.replace("a");
  r.replace("ab");
  assert.equal(first, 1);
  r.add("c");
  assert.equal(first, 1);
});

test("renderer reset vide le buffer", () => {
  const r = createStreamRenderer({ onRender: () => {} });
  r.replace("texte");
  r.reset();
  assert.equal(r.text(), "");
  assert.equal(r.isActive(), false);
});

test("renderer add accumule dans le buffer", () => {
  const r = createStreamRenderer({ onRender: () => {} });
  r.add("a");
  r.add("b");
  assert.equal(r.text(), "ab");
  r.reset();
});

test("renderer reducedMotion rend immediatement sans rAF", () => {
  const seen = [];
  const r = createStreamRenderer({ reducedMotion: true, onRender: (t) => seen.push(t) });
  r.add("abc");
  r.add("def");
  assert.equal(r.text(), "abcdef");
  assert.deepEqual(seen, ["abc", "abcdef"]);
  assert.equal(r.isActive(), false);
});
