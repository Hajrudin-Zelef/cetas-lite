---
id: collect-261001-rattrapage/rattrapage/javascript-guide-2
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [197, 455]
sha256: f6fb07f3559165e03fc49ecba949bb6df4643c620bbe2ba898442777c2d73668
---

# JavaScript — Guide ultra-complet

|  | `var` | `let` | `const` |
|---|---|---|---|
| Portée | fonction | bloc `{}` | bloc `{}` |
| Hoisting | oui (→ `undefined`) | TDZ (erreur si accès avant) | TDZ |
| Redéclaration | autorisée | interdite | interdite |
| Réassignation | autorisée | autorisée | interdite |
| Usage recommandé | **jamais** (code legacy) | variable qui change | **par défaut** |

### Le piège classique avec `var` dans les boucles

```js
// BUG classique
for (var i = 0; i < 3; i++) {
  setTimeout(() => console.log("var:", i), 10); // 3, 3, 3 !
}
// CORRIGÉ
for (let i = 0; i < 3; i++) {
  setTimeout(() => console.log("let:", i), 10); // 0, 1, 2
}
```

> **Pense-bête** : `const` par défaut, `let` si ça change, `var` jamais.

---

## 4. Types primitifs

JavaScript a 7 types primitifs + `object`. `typeof` renvoie le type (avec le bug historique `typeof null === "object"`).

```js
typeof "hello"      // "string"
typeof 42           // "number"
typeof 42n          // "bigint"
typeof true         // "boolean"
typeof undefined    // "undefined"
typeof Symbol("s")  // "symbol"
typeof null         // "object"  <-- bug historique, à connaître
typeof {}           // "object"
typeof []           // "object"  (utiliser Array.isArray)
typeof function(){} // "function"
```

### Détails par type

```js
// string : immuable, guillemets simples/doubles/backticks équivalents
const s = "Cisco";
s[0];            // "C"
s.length;        // 5
// s[0] = "c";   // silencieux en sloppy, TypeError en strict : immuable !

// number : TOUS les nombres sont des flottants double précision (IEEE 754)
0.1 + 0.2;       // 0.30000000000000004  <-- piège !
Number.MAX_SAFE_INTEGER; // 9007199254740991 (2^53 - 1)

// bigint (ES2020) : entiers arbitraires, suffixe n
const big = 9007199254740993n;
big + 1n;        // 9007199254740994n
// big + 1;      // TypeError : ne pas mélanger bigint et number

// boolean : truthy / falsy
// Falsy : false, 0, -0, 0n, "", null, undefined, NaN  (7 valeurs, à connaître par cœur)
// Tout le reste est truthy, y compris [] et {} !
Boolean([]);     // true  <-- piège classique
Boolean("0");    // true  (chaîne non vide)

// undefined vs null
let x;           // undefined : "pas encore de valeur" (moteur)
let y = null;    // null : "volontairement vide" (programmeur)

// symbol (ES2015) : identifiant unique, voir section 43
const id = Symbol("id");
```

### Conversions

```js
// Explicites (recommandées)
Number("42");    // 42
Number("");      // 0
Number("abc");   // NaN
String(42);      // "42"
Boolean(0);      // false
parseInt("42px", 10); // 42  (toujours préciser la base !)
parseFloat("3.14");   // 3.14

// Implicites (pièges)
"5" + 1;         // "51"  (concaténation !)
"5" - 1;         // 4     (- force la conversion numérique)
"5" * "2";       // 10
true + 1;        // 2
null + 1;        // 1
undefined + 1;   // NaN
```

---

## 5. Opérateurs

### 5.1 Comparaison : `==` vs `===`

```js
// == : égalité LÂCHE (avec coercition de type) — À ÉVITER
0 == "0";        // true
0 == [];         // true  (!!)
"" == false;     // true
null == undefined; // true (seul cas "utile")

// === : égalité STRICTE (type + valeur) — TOUJOURS utiliser
0 === "0";       // false
NaN === NaN;     // false  <-- NaN n'est jamais égal à lui-même !
Object.is(NaN, NaN); // true (Object.is : comparaison "SameValue")
Object.is(0, -0);    // false (=== dirait true)
```

> **Règle d'or** : toujours `===` et `!==`. Pour tester NaN : `Number.isNaN(x)`.

### 5.2 Logiques et courts-circuits

```js
// && et || retournent un OPÉRANDE, pas forcément un booléen
const port = process.env.PORT || 3000;   // valeur par défaut
const user = input && input.trim();      // garde si input truthy

// ?? : nullish coalescing (ES2020) — ne réagit qu'à null/undefined
const p1 = 0 || 3000;    // 3000 (0 est falsy — bug potentiel !)
const p2 = 0 ?? 3000;    // 0 (correct)

// ?. : optional chaining (ES2020) — navigation sécurisée
const ip = srv?.reseau?.ip;  // undefined si un maillon est null/undefined
srv?.demarrer?.();           // n'appelle que si la méthode existe

// ! et !!
const connecte = !!token;    // force en booléen
```

### 5.3 Arithmétiques et autres

```js
let n = 10;
n++; n--;            // incrément / décrément
2 ** 10;             // 1024 (exponentiation, ES2016)
17 % 5;              // 2 (modulo)
10 / 0;              // Infinity
-10 / 0;             // -Infinity
0 / 0;               // NaN

// Affectation composée
n += 5; n -= 2; n *= 3; n **= 2; n %= 4;

// Opérateurs logiques d'affectation (ES2021)
let a2; a2 ??= 5;    // a2 = 5 (si null/undefined)
let b2 = 0; b2 ||= 9; // b2 = 9 (si falsy)
let c2 = 1; c2 &&= 9; // c2 = 9 (si truthy)

// typeof, instanceof, in, delete, void
"host" in {host: "x"};  // true
const o = {a: 1}; delete o.a;
void 0;              // undefined (idiome historique)
```

---

## 6. Structures de contrôle

```js
// if / else if / else
const charge = 82;
if (charge > 90) console.log("CRITIQUE");
else if (charge > 70) console.log("WARNING");
else console.log("OK");

// Ternaire (pour des valeurs, pas pour des effets de bord complexes)
const etat = charge > 90 ? "CRITIQUE" : charge > 70 ? "WARNING" : "OK";

// switch (comparaison stricte ===)
const os = "linux";
switch (os) {
  case "linux": console.log("apt/dnf"); break;
  case "windows": console.log("winget"); break;
  default: console.log("inconnu");
}

// Boucles
for (let i = 0; i < 3; i++) { /* ... */ }

for (const el of [10, 20, 30]) { console.log(el); }      // valeurs (itérables)
for (const k in {a: 1, b: 2}) { console.log(k); }         // clés (objets) — éviter sur tableaux !

let j = 0;
while (j < 3) { j++; }
do { j--; } while (j > 0);   // exécuté au moins une fois

// break / continue avec labels (rare mais utile pour boucles imbriquées)
boucleExterne:
for (const hote of hotes) {
  for (const port of ports) {
    if (port === 22) continue boucleExterne;
  }
}
```

### Choisir la bonne boucle

| Besoin | Construction |
|---|---|
| Compteur classique | `for (let i = 0; ...)` |
| Parcourir un tableau (valeurs) | `for...of` ou `.forEach/.map` |
| Parcourir les clés d'un objet | `for...in` ou `Object.keys()` |
| Condition d'arrêt inconnue | `while` |
| Au moins une exécution | `do...while` |
| Transformer un tableau | `.map/.filter/.reduce` (section 10) |

---

## 7. Fonctions — les fondamentaux

```js
// 1. Déclaration (hoisted : appelable avant sa définition)
function addition(a, b) { return a + b; }

// 2. Expression (non hoisted)
const soustraction = function(a, b) { return a - b; };

// 3. Fléchée (ES2015) — concise, pas de `this` propre, pas de `arguments`
const multiplication = (a, b) => a * b;          // return implicite
const carre = x => x * x;                        // 1 param : parenthèses optionnelles
const ping = () => console.log("pong");
const fabrique = (hote) => ({ hote, port: 22 }); // objet : parenthèses obligatoires

// Fonctions = objets de première classe
function executer(fn, ...args) { return fn(...args); }
executer(addition, 2, 3); // 5
```

### return implicite vs bloc

```js
const f1 = x => x * 2;        // retourne x*2
const f2 = x => { x * 2; };   // retourne undefined ! (bloc sans return)
const f3 = x => { return x * 2; }; // explicite
```

### Hoisting des fonctions

```js
saluer(); // OK
function saluer() { console.log("salut"); }

// aurevoir(); // TypeError : pas une fonction (undefined au moment de l'appel)
const aurevoir = function() { console.log("bye"); };
```

---

## 8. Closures

Une **closure** = une fonction qui "capture" les variables de la portée où elle a été **définie** (pas appelée). C'est le mécanisme derrière les callbacks, les fabriques et l'encapsulation.

