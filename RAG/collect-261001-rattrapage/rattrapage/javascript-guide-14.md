---
id: collect-261001-rattrapage/rattrapage/javascript-guide-14
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [3216, 3377]
sha256: 25f67785197e75375f68dea3da5f96d78b08aa314d993cd2a2f63dc839eef135
---

# JavaScript — Guide ultra-complet

```js
await fetch(url);                             // pause sans bloquer
const [a, b] = await Promise.all([f1(), f2()]); // parallèle
for (const x of xs) await traiter(x);         // séquentiel
await Promise.all(xs.map(x => traiter(x)));   // parallèle (map)
try { await f(); } catch (e) { /* ... */ }    // erreurs async
```

### Les 7 falsy (par cœur)

`false` `0` `-0` `0n` `""` `null` `undefined` `NaN` — tout le reste est truthy (y compris `[]` et `{}`).

### Vérifications express

```js
Array.isArray(x);            // vrai tableau ?
Number.isNaN(x);             // NaN ?
x == null;                   // null ou undefined (idiome toléré)
typeof x === "string";       // type primitif
x instanceof MaClasse;       // instance
Object.hasOwn(o, "cle");     // propriété propre (ES2022)
```

### Node express

```bash
node script.js                 # exécuter
node --watch script.js         # rechargement auto
node --test                    # tests natifs
node --env-file=.env app.js    # variables d'environnement
node --inspect app.js          # debug DevTools
npm init -y && npm i <pkg>     # projet + dépendance
npx <pkg>                      # exécuter sans installer
```

### Commandes qui sauvent

```js
console.table(data);           // inspecter un tableau d'objets
console.dir(obj, {depth:null});// tout voir
JSON.stringify(o, null, 2);    // pretty-print
structuredClone(o);            // copie profonde
[...new Set(t)];               // dédupliquer
t.sort((a,b) => a-b);          // tri numérique
new URL(url).searchParams;     // parser query string
AbortSignal.timeout(5000);     // timeout fetch natif
```

---

## 60. Glossaire

| Terme | Définition |
|---|---|
| ASI | Automatic Semicolon Insertion : le moteur ajoute les `;` manquants (ne pas s'y fier aveuglément) |
| Callback | fonction passée en argument pour être appelée plus tard |
| Closure | fonction + les variables capturées de sa portée de définition |
| Coercition | conversion implicite de type (`"5" - 1` → `4`) |
| Debounce / Throttle | limiter la fréquence d'appels (recherche en direct, resize) |
| Event loop | boucle qui orchestre l'exécution : stack → microtâches → macrotâches |
| Hoisting | remontée des déclarations (`var`, `function`) en haut de la portée |
| IIFE | Immediately Invoked Function Expression : `(function(){...})()` |
| Polyfill | code qui apporte une API moderne aux vieux environnements |
| Promise | valeur future (pending → fulfilled / rejected) |
| Prototype | objet dont hérite une instance (chaîne de recherche des propriétés) |
| Scope | portée : zone où une variable est visible (bloc, fonction, module, globale) |
| Shadowing | variable locale qui masque une variable externe du même nom |
| Strict mode | `"use strict"` : erreurs sur les pratiques dangereuses (implicite en modules/classes) |
| TDZ | Temporal Dead Zone : `let`/`const` inaccessibles avant leur déclaration |
| Thunk | fonction sans argument qui retarde un calcul : `() => couteux()` |
| Tree-shaking | élimination du code inutilisé au build (possible grâce aux imports ESM statiques) |
| Truthy / Falsy | valeurs évaluées à `true`/`false` dans un contexte booléen |

---

## 61. Quiz — 10 questions + réponses

**Q1.** Que vaut `typeof null` et pourquoi ?
> `"object"` — bug historique du moteur, conservé pour compatibilité. Tester `x === null` explicitement.

**Q2.** Quelle est la différence entre `==` et `===` ? Lequel utiliser ?
> `==` compare avec coercition de type (`0 == "0"` → true), `===` compare type + valeur sans conversion. Toujours `===`.

**Q3.** Que se passe-t-il ici et comment corriger ?
> ```js
> for (var i = 0; i < 3; i++) setTimeout(() => console.log(i), 0);
> ```
> Affiche `3, 3, 3` : `var` est partagé par les 3 closures (même variable mutée). Corriger avec `let` (une variable par itération).

**Q4.** Pourquoi `fetch` vers une URL qui répond 500 n'entre-t-il pas dans le `catch` ?
> `fetch` ne rejette que sur échec réseau. Un statut HTTP 500 donne une promise *fulfilled* avec `ok === false`. Il faut tester `res.ok` / `res.status` manuellement.

**Q5.** `urls.forEach(async u => { await fetch(u); })` — quel est le problème ?
> `forEach` n'attend pas les callbacks : les requêtes partent en fire-and-forget et le code suivant s'exécute avant leur fin. Utiliser `for...of` (séquentiel) ou `Promise.all(urls.map(...))` (parallèle).

**Q6.** Citer les 4 règles de `this` par ordre de priorité.
> 1. `new` → le nouvel objet ; 2. appel méthode `obj.m()` → `obj` ; 3. `call`/`apply`/`bind` → l'objet passé ; 4. défaut → `undefined` (strict) ou objet global. Les fléchées n'ont pas de `this` propre (hérité lexicalement).

**Q7.** Dans l'event loop, qui passe en premier : le callback de `setTimeout(..., 0)` ou celui de `Promise.resolve().then(...)` ?
> La microtâche (promise) : l'event loop vide toujours la file des microtâches avant de prendre la prochaine macrotâche (timer).

**Q8.** `const cfg = { port: 22 }; cfg.port = 2222;` — erreur ou pas ? Pourquoi ?
> Pas d'erreur : `const` fige la *référence*, pas le contenu. L'objet reste mutable. Pour figer : `Object.freeze()`.

**Q9.** Comment copier un objet en profondeur, et pourquoi `JSON.parse(JSON.stringify(o))` est-il imparfait ?
> `structuredClone(o)` (moderne). Le hack JSON perd les `Date` (→ chaînes), `undefined`/fonctions/`Symbol`, échoue sur `BigInt` et les références circulaires.

**Q10.** Écrire une fonction `sleep(ms)` retournant une promise, puis l'utiliser pour attendre 2 s.
> ```js
> const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));
> await sleep(2000);
> ```

---

## 62. Pour aller plus loin

### Documentation de référence
- **MDN Web Docs** (developer.mozilla.org) — LA référence JS : chaque méthode avec exemples et compatibilité.
- **node.js.org/docs** — API Node officielle.
- **javascript.info** — tutoriel moderne, excellent sur les subtilités.

### Approfondir par sujet
| Sujet | Ressource |
|---|---|
| Event loop visuelle | Conf "What the heck is the event loop?" (Philip Roberts, YouTube) |
| Async moderne | MDN "Using promises" + "async function" |
| Node en prod | "Node.js Best Practices" (goldbergyoni, GitHub) |
| Regex | regex101.com (tester en direct), regexr.com |
| Sécurité | OWASP Cheat Sheets |
| TypeScript | typescriptlang.org — le successeur typé (recommandé pour projets > quelques fichiers) |

### Prochaines étapes concrètes pour Zelef
1. Convertir un script shell récurrent (sauvegarde, rapport) en script Node CLI avec `commander`.
2. Exposer un endpoint `/sante` sur un outil interne existant (section 54).
3. Ajouter des tests `node --test` sur les fonctions de parsing (section 46).
4. Passer un projet à TypeScript quand la base dépasse ~1000 lignes.
5. Conteneuriser un outil Node avec Docker (multi-stage build, user non-root).

### Frameworks (quand le besoin vient)
- **API** : Express (simple), Fastify (rapide), NestJS (structuré).
- **Front** : pas besoin de framework pour un dashboard interne (section 56 suffit) ; sinon Vue ou Svelte.
- **CLI** : `commander`, `yargs`.
- **Temps réel** : `ws` (WebSocket), Socket.IO.

---

## 63. Checklist de mise en production

### Code
- [ ] ESLint passe sans erreur, Prettier appliqué
- [ ] `node --test` : tous les tests verts
- [ ] Aucun `console.log` de debug restant (logger structuré : `pino`/`winston`)
- [ ] Gestion d'erreurs : `try/catch` aux frontières (I/O, réseau, parsing)
- [ ] `unhandledRejection` / `uncaughtException` gérés (log + exit 1)

### Config & secrets
- [ ] Secrets uniquement via variables d'environnement / vault — rien en dur, rien dans git
- [ ] `.env` dans `.gitignore`, `.env.example` commité (sans valeurs)
- [ ] `package-lock.json` commité, `npm ci` utilisé au déploiement
- [ ] `engines.node` fixé dans package.json

