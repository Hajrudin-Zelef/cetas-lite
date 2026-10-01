---
id: collect-261001-rattrapage/rattrapage/javascript-guide-10
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [2229, 2483]
sha256: 93af45cdc072ccad1a7e6b8de2974acfcbfac81be276cf3653c7fc8d044c36d1
---

# JavaScript — Guide ultra-complet

// Opérations ensemblistes (ES2025)
const a = new Set([1, 2, 3]), b = new Set([2, 3, 4]);
a.union(b);            // Set {1,2,3,4}
a.intersection(b);     // Set {2,3}
a.difference(b);       // Set {1}

// WeakMap / WeakSet : clés OBJETS, références faibles (pas d'empêchement GC)
const cache = new WeakMap();
cache.set(objetEquipement, donneesCalculees); // libéré si objetEquipement est GC
// Pas de .size, pas d'itération — usage : caches, métadonnées privées
```

### Map vs Object

|  | `Map` | `Object` |
|---|---|---|
| Clés | tout type | string/symbol (coercition !) |
| Ordre | insertion garanti | quasi (entiers d'abord triés) |
| Taille | `.size` | `Object.keys(o).length` |
| Itération | directe | via `Object.entries` |
| Perf ajout/suppr fréquents | meilleure | — |
| JSON | non natif | natif |

> Clés dynamiques non-string, ajouts/suppressions fréquents → `Map`. Config/objet métier sérialisable → `Object`.

---

## 42. Itérateurs et générateurs

```js
// Tout objet avec [Symbol.iterator] est itérable (for...of, spread, ...)
const iterable = {
  *[Symbol.iterator]() { yield 1; yield 2; yield 3; }
};
[...iterable]; // [1, 2, 3]

// Générateur : fonction* — exécution PARESSEUSE, état interne
function* scanSousReseau(base) {
  for (let i = 1; i < 255; i++) {
    yield `${base}.${i}`; // pause ici, reprend au prochain next()
  }
}
const gen = scanSousReseau("10.0.0");
gen.next(); // { value: "10.0.0.1", done: false }

// Cas sysadmin : générer des IP sans allouer un tableau de 65000 entrées
function* plageIp(cidrBase, debut, fin) {
  for (let i = debut; i <= fin; i++) yield `${cidrBase}.${i}`;
}
for (const ip of plageIp("192.168.1", 1, 254)) {
  // await ping(ip); // traitement à la volée, mémoire constante
  if (ip === "192.168.1.10") break;
}

// Générateur async (for await...of)
async function* lireLignes(chemin) {
  const rl = readline.createInterface({ input: createReadStream(chemin) });
  for await (const ligne of rl) yield ligne;
}
```

---

## 43. Symbols et BigInt

```js
// Symbol : identifiant UNIQUE — clés "cachées", pas de collision
const CLE_PRIVEE = Symbol("cle");
const obj = { [CLE_PRIVEE]: "secret interne", public: "visible" };
Object.keys(obj);              // ["public"] — le symbol n'apparaît pas
obj[CLE_PRIVEE];               // "secret interne"

// Symboles connus (well-known) : customiser le comportement natif
class Plage {
  constructor(debut, fin) { this.debut = debut; this.fin = fin; }
  *[Symbol.iterator]() { for (let i = this.debut; i <= this.fin; i++) yield i; }
  get [Symbol.toStringTag]() { return "Plage"; }
}
[...new Plage(1, 3)]; // [1, 2, 3]

// BigInt : entiers au-delà de 2^53-1 (ES2020)
const compteur = 9007199254740993n;
typeof compteur;              // "bigint"
compteur + 1n;                // OK
// compteur + 1;              // TypeError : conversion explicite requise
Number(compteur);             // 9007199254740992 (perte de précision !)
// Usage : compteurs, timestamps ns, crypto, identifiants 64 bits
```

---

## 44. Proxy et Reflect

`Proxy` = intercepter les opérations sur un objet (validation, logs, valeurs par défaut).

```js
const config = new Proxy({ port: 3000 }, {
  get(cible, prop) {
    if (!(prop in cible)) console.warn(`Config inconnue : ${String(prop)}`);
    return cible[prop];
  },
  set(cible, prop, valeur) {
    if (prop === "port" && (valeur < 1 || valeur > 65535))
      throw new RangeError("port invalide");
    cible[prop] = valeur;
    return true;
  }
});
config.port = 99999; // RangeError
console.log(config.hote); // warn + undefined

// Reflect : mêmes opérations en version fonctionnelle
Reflect.get(config, "port");
Reflect.has(config, "port");
```

> Usage réel : validation de config, ORM, observabilité. À manier avec parcimonie (coût perf, debug plus dur).

---

## 45. Programmation fonctionnelle

```js
// Fonctions pures : même entrée → même sortie, pas d'effet de bord
const ajouterTva = (prix, taux = 0.2) => prix * (1 + taux); // pure ✓
// let total = 0; const ajouter = (p) => total += p; // impure ✗ (état externe)

// Immuabilité : créer plutôt que muter
const srv1 = { nom: "SW1", ports: 48 };
const srv2 = { ...srv1, ports: 24 }; // srv1 intact

// Composition : enchaîner des petites fonctions
const pipe = (...fns) => (x) => fns.reduce((v, f) => f(v), x);
const normaliser = pipe(
  s => s.trim(),
  s => s.toLowerCase(),
  s => s.replace(/\s+/g, "-")
);
normaliser("  SW 1 Core  "); // "sw-1-core"

// Curryfication : spécialiser progressivement
const requete = (base) => (chemin) => (params) =>
  fetch(`${base}${chemin}?${new URLSearchParams(params)}`);
const apiInterne = requete("https://api.interne.example.com");
const getEquipements = apiInterne("/equipements");
await getEquipements({ site: "paris" });
```

### Méthodes fonctionnelles de tableau (référence)

| Méthode | Retourne | Usage |
|---|---|---|
| `map` | tableau transformé | extraire un champ |
| `filter` | sous-ensemble | garder les erreurs |
| `reduce` | une valeur | somme, groupement, agrégation |
| `find` / `findIndex` | élément / index | premier match |
| `some` / `every` | booléen | au moins un / tous |
| `flatMap` | aplati | 1 → N |
| `sort` (avec comparateur) | trié (mutant !) | trier |

---

## 46. Tests unitaires

### Avec le runner natif Node (Node 18+, `--test`)

```js
// maths.js
export const somme = (a, b) => a + b;
export const diviser = (a, b) => {
  if (b === 0) throw new RangeError("division par zéro");
  return a / b;
};

// maths.test.js
import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { somme, diviser } from "./maths.js";

describe("maths", () => {
  it("somme additionne", () => {
    assert.equal(somme(2, 3), 5);
  });
  it("diviser lève une erreur sur zéro", () => {
    assert.throws(() => diviser(1, 0), RangeError);
  });
  it("async supporté", async () => {
    const data = await Promise.resolve(42);
    assert.deepEqual(data, 42);
  });
});
```

```bash
node --test                 # lance tous les *.test.js
node --test --test-reporter=spec
```

```json
// package.json
{ "scripts": { "test": "node --test" } }
```

### Assertions essentielles (`node:assert/strict`)

| Assertion | Vérifie |
|---|---|
| `assert.equal(a, b)` | `==` (éviter) |
| `assert.strictEqual(a, b)` | `===` |
| `assert.deepStrictEqual(a, b)` | structure profonde |
| `assert.throws(fn, Erreur)` | lève bien l'erreur |
| `assert.rejects(promise, Erreur)` | promise rejetée |
| `assert.ok(valeur)` | truthy |
| `assert.match(str, regex)` | regex |

### Bonnes pratiques de test

- Nommage : `fichier.test.js` à côté du fichier.
- Tester les **comportements** (contrats), pas l'implémentation.
- Un test = un comportement ; nom explicite ("rejette les ports hors plage").
- Mocker les I/O (fs, réseau, temps) — un test unitaire ne touche ni disque ni réseau.
- Viser les cas limites : vide, `null`, zéro, bornes, erreurs.
- Alternatives : **Vitest** (rapide, moderne), **Jest** (écosystème historique).

---

## 47. Debug

### Console avancée

```js
console.log("valeur :", x);
console.log("objet : %o", objet);           // format objet
console.table(tableauObjets);               // tableau lisible
console.dir(objet, { depth: null });        // profondeur illimitée
console.trace("où suis-je ?");              // stack trace
console.count("appel");                     // compteur par label
console.time("requete"); /* ... */ console.timeEnd("requete");
console.group("section"); /* ... */ console.groupEnd();
```

### Debugger (le vrai outil)

```js
function calculer(x) {
  debugger; // point d'arrêt : pause ici si DevTools / --inspect attaché
  return x * 2;
}
```

```bash
node --inspect script.js        # DevTools Chrome : chrome://inspect
node --inspect-brk script.js    # pause dès le démarrage
```

