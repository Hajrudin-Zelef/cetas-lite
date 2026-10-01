---
id: collect-261001-rattrapage/rattrapage/javascript-guide-11
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "valuation"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [2484, 2725]
sha256: aa6b739a0688cb7ef150c70cb62f6e2ac29c923b91ce6d56a36f7032cf8c4891
---

# JavaScript — Guide ultra-complet

Dans VS Code : F5 avec une config `launch.json` → breakpoints visuels, variables, pile d'appels, évaluation d'expressions.

### Stratégie de debug

1. **Reproduire** avec le cas minimal.
2. **Lire l'erreur** : `message` + 3 premières lignes de `stack`.
3. **Isoler** : `console.log`/`debugger` en dichotomie (moitié du code).
4. **Vérifier les hypothèses** : `typeof`, valeurs réelles vs supposées.
5. Pièges fréquents : async non attendu, `this` perdu, mutation surprise, cache.

### Erreurs Node courantes : décodage

| Message | Cause probable |
|---|---|
| `Cannot read properties of undefined (reading 'x')` | objet `undefined` (chaînage `?.` ou garde) |
| `x is not a function` | mauvaise import, écrasement, `this` perdu |
| `Unexpected token` | syntaxe invalide / JSON malformé |
| `Cannot find module './x'` | chemin/extension faux (ESM exige l'extension) |
| `require is not defined` | fichier ESM (`"type": "module"`) |
| `await is only valid in async` | `await` hors fonction async / hors module |
| `Maximum call stack size exceeded` | récursion infinie |
| `EADDRINUSE` | port déjà occupé |
| `ENOENT` | fichier/chemin inexistant |

---

## 48. Bonnes pratiques & style

### Les 12 règles d'or

1. `const` par défaut, `let` si réassigné, jamais `var`.
2. `===` toujours, jamais `==` (sauf `x == null` idiome, à connaître).
3. Fonctions courtes, un seul niveau d'abstraction, un seul "travail".
4. Noms explicites : `recupererEquipementsEnPanne()` > `getData2()`.
5. Retour anticipé (early return) plutôt qu'imbrications.
6. Ne jamais muter les arguments d'entrée.
7. Gérer les erreurs là où on peut agir ; logger avec contexte.
8. Async/await plutôt que `.then` imbriqués ; `Promise.all` pour le parallèle.
9. Valider les entrées externes (API, CLI, fichiers) — ne jamais faire confiance.
10. Pas de secrets en dur : variables d'environnement.
11. Commenter le **pourquoi**, pas le **quoi**.
12. Linter (ESLint) + formateur (Prettier) dès le premier jour.

### Early return

```js
// MAL : pyramide
function traiter(req) {
  if (req) {
    if (req.body) {
      if (req.body.ip) { /* ... */ }
    }
  }
}
// BIEN : gardes
function traiter(req) {
  if (!req?.body?.ip) return { erreur: "IP manquante" };
  /* ... logique principale au premier niveau ... */
}
```

### ESLint + Prettier (setup minimal)

```bash
npm init -y && npm install -D eslint prettier
npx eslint --init   # ou config manuelle
```

```js
// eslint.config.js (flat config, ESLint 9+)
export default [
  { rules: {
    "no-var": "error",
    "eqeqeq": ["error", "always"],
    "no-unused-vars": "warn",
    "prefer-const": "error",
  }}
];
```

### Style : points qui comptent

- Point-virgule : choisir (avec/sans) et laisser Prettier trancher.
- Guillemets simples en JS, backticks pour l'interpolation.
- `camelCase` variables/fonctions, `PascalCase` classes, `SNAKE_UPPER` constantes.
- Fichiers : un module = une responsabilité ; `kebab-case.js`.

---

## 49. Performance

### Mesurer d'abord

```js
console.time("boucle");
// ... code ...
console.timeEnd("boucle"); // "boucle: 12.345ms"

// Précis : performance.now()
const t0 = performance.now();
// ...
console.log(`${(performance.now() - t0).toFixed(2)} ms`);

// Benchmark sérieux : répéter + médiane (bruit de mesure)
function bench(fn, iterations = 10000) {
  const temps = [];
  for (let i = 0; i < iterations; i++) {
    const t0 = performance.now(); fn(); temps.push(performance.now() - t0);
  }
  temps.sort((a, b) => a - b);
  return temps[Math.floor(temps.length / 2)]; // médiane
}
```

### Règles de performance

| Règle | Pourquoi |
|---|---|
| Éviter le travail dans les boucles chaudes | hoister les invariants, pré-calculer |
| `for` > `forEach`/`map` en boucle ultra-chaude | moins d'overhead (micro-opt, à mesurer) |
| `Map`/`Set` pour recherches fréquentes | O(1) vs O(n) de `Array.includes`/`find` |
| Concaténer avec `push` + `join`, pas `+=` en boucle géante | chaînes immuables = copies |
| `JSON.parse/stringify` : coûteux sur gros volumes | streamer si possible |
| Regex pré-compilées hors boucle | `new RegExp` à chaque itération = lent |
| Éviter les fuites : listeners, timers, closures sur gros objets | mémoire qui ne redescend jamais |
| `structuredClone` > `JSON.parse(JSON.stringify())` | plus rapide et plus fidèle |

### Mémoire Node

```bash
node --max-old-space-size=4096 app.js  # augmenter le heap si besoin
node --inspect app.js                   # heap snapshots dans DevTools
```

```js
// Surveiller
const { heapUsed } = process.memoryUsage();
console.log(`heap : ${(heapUsed / 1024 / 1024).toFixed(1)} Mo`);
```

### Anti-patterns perf

```js
// MAL : N+1 requêtes séquentielles
for (const id of ids) { await fetch(`/api/eq/${id}`); } // N allers-retours
// BIEN : parallèle borné (voir pool section 24) ou endpoint batch

// MAL : JSON.parse/stringify pour cloner en boucle chaude
// BIEN : structuredClone, ou mieux : ne pas cloner (immuabilité ciblée)

// MAL : regex recompilée
for (const l of lignes) { if (/error/i.test(l)) n++; } // en fait /.../ littéral = pré-compilé, OK
for (const l of lignes) { if (new RegExp(motif).test(l)) n++; } // MAL : compilation à chaque tour
```

---

## 50. Sécurité

### Checklist sécurité (à appliquer systématiquement)

- [ ] Aucun secret en dur (code, git) → variables d'environnement / vault
- [ ] `execFile`/`spawn` avec args séparés, **jamais** d'interpolation dans `exec`
- [ ] Validation de toutes les entrées externes (schéma, plages, types)
- [ ] `textContent` plutôt que `innerHTML` avec des données externes
- [ ] Cookies auth : `HttpOnly; Secure; SameSite=Lax` (posés serveur)
- [ ] Dépendances : `npm audit` régulier, lockfile commité
- [ ] Erreurs : messages génériques côté client, détails côté logs
- [ ] `npm ci` en prod, pas `npm install`
- [ ] Principe du moindre privilège (utilisateur dédié, pas root)
- [ ] Timeouts sur toutes les I/O réseau

### Injection de commande (rappel section 33)

```js
// INTERDIT avec entrée utilisateur
exec(`ping -c 1 ${hote}`);
// AUTORISÉ
execFile("ping", ["-c", "1", hote]);
```

### XSS (navigateur)

```js
// DANGEREUX si `nom` vient d'un utilisateur/API
el.innerHTML = `<b>${nom}</b>`;
// SÛR
el.textContent = nom;
// ou échapper manuellement si HTML nécessaire :
const echapper = (s) => s.replace(/[&<>"']/g, c =>
  ({ "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[c]));
```

### Prototype pollution

```js
// DANGEREUX : fusion récursive naïve d'objets externes
// { "__proto__": { "isAdmin": true } } peut polluer Object.prototype !
// Défenses :
const cible = Object.create(null); // objet sans prototype
// ou : valider les clés ("__proto__", "constructor", "prototype" interdites)
// ou : structuredClone / lib de merge sûre
```

### Secrets et tokens

```js
// Générer un token sûr (jamais Math.random() pour la sécurité !)
import { randomBytes } from "crypto";
const token = randomBytes(32).toString("hex"); // 64 caractères hex

// Comparer des secrets en temps constant (anti timing-attack)
import { timingSafeEqual } from "crypto";
const ok = timingSafeEqual(Buffer.from(a), Buffer.from(b)); // mêmes longueurs requises
```

### Validation d'entrée (exemple)

```js
function validerEquipement(e) {
  const erreurs = [];
  if (typeof e.hostname !== "string" || !/^[A-Za-z0-9-]{1,32}$/.test(e.hostname))
    erreurs.push("hostname invalide");
  if (!/^(?:\d{1,3}\.){3}\d{1,3}$/.test(e.ip ?? ""))
    erreurs.push("ip invalide");
  if (e.port !== undefined && (!Number.isInteger(e.port) || e.port < 1 || e.port > 65535))
    erreurs.push("port invalide");
  if (erreurs.length) throw new Error("Validation : " + erreurs.join("; "));
  return e;
}
```

---

## 51. Erreurs classiques des débutants (15)

### 1. `==` au lieu de `===`
```js
if (code == "200") // true aussi pour 200 (number) — ambigu
if (code === 200)  // strict, prévisible
```

