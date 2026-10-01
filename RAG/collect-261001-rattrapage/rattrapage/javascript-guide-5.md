---
id: collect-261001-rattrapage/rattrapage/javascript-guide-5
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [968, 1218]
sha256: 7fa73f2c967960937b9a5ba3f0868bea6fce31973d1a1b050df53c4053def302
---

# JavaScript — Guide ultra-complet

class Switch extends Equipement {
  constructor(hostname, ip, nbPorts) {
    super(hostname, ip);   // OBLIGATOIRE avant d'utiliser this
    this.nbPorts = nbPorts;
  }
  // Surcharge (override) : polymorphisme
  decrire() { return `${super.decrire()} — ${this.nbPorts} ports`; }
}

class Routeur extends Equipement {
  constructor(hostname, ip, as) { super(hostname, ip); this.as = as; }
  decrire() { return `${super.decrire()} — AS${this.as}`; }
}

const parc = [new Switch("SW1", "10.0.0.2", 48), new Routeur("R1", "10.0.0.1", 64512)];
for (const e of parc) console.log(e.decrire()); // polymorphisme : chaque classe répond à sa façon
parc[0] instanceof Switch;      // true
parc[0] instanceof Equipement;  // true (chaîne d'héritage)
```

### Composition > héritage (principe)

```js
// Plutôt que d'hériter, composer : plus souple
const Supervisable = (obj) => ({
  ...obj,
  superviser() { return `supervision de ${this.hostname}`; }
});
const sw = Supervisable(new Switch("SW1", "10.0.0.2", 48));
```

---

## 20. Gestion d'erreurs

```js
// Lancer
function diviser(a, b) {
  if (b === 0) throw new RangeError("division par zéro");
  return a / b;
}

// Capturer : try / catch / finally
let resultat;
try {
  resultat = diviser(10, 0);
} catch (err) {
  console.error(`${err.name}: ${err.message}`); // "RangeError: division par zéro"
  // err.stack : pile d'appels (debug)
  resultat = Infinity;
} finally {
  console.log("nettoyage : toujours exécuté");
}
```

### Types d'erreurs natives

| Classe | Quand |
|---|---|
| `Error` | générique |
| `SyntaxError` | code invalide (compilation) |
| `ReferenceError` | variable inexistante / TDZ |
| `TypeError` | mauvais type (`null.foo()`, réassigner `const`) |
| `RangeError` | valeur hors limites |
| `URIError` | `encodeURI` invalide |
| `AggregateError` (ES2021) | plusieurs erreurs (`Promise.any`) |

### Erreurs métier personnalisées

```js
class HttpError extends Error {
  constructor(status, message) {
    super(message);
    this.name = "HttpError";
    this.status = status;
  }
}
// catch (e) { if (e instanceof HttpError && e.status === 404) {...} }
```

### Bonnes pratiques

- Capturer là où on peut **agir** (retry, fallback, message utilisateur) ; sinon laisser remonter.
- En async : `try/catch` autour de `await`, ou `.catch()` sur la promise.
- Ne jamais `catch {}` vide silencieux — au minimum logger.
- `finally` pour libérer les ressources (fermer fichier, connexion).
- Node : écouter `process.on("uncaughtException")` et `unhandledRejection` (section 30).

---

## 21. Event loop : call stack, micro/macro-tâches

JavaScript est **mono-thread** : une seule pile d'exécution (call stack). L'event loop orchestre l'asynchrone.

```
┌─────────────────────────────┐
│         CALL STACK          │  ← exécution synchrone (LIFO)
└──────────────┬──────────────┘
               │ vide ?
               ▼
┌─────────────────────────────┐
│   MICROTÂCHES (prioritaires)│  ← promises (.then), queueMicrotask, await
└──────────────┬──────────────┘
               │ vides ?
               ▼
┌─────────────────────────────┐
│   MACROTÂCHES               │  ← setTimeout, setInterval, I/O, setImmediate
└─────────────────────────────┘
        ↺ boucle infinie
```

```js
console.log("1. synchrone");

setTimeout(() => console.log("4. macrotâche (setTimeout)"), 0);

Promise.resolve().then(() => console.log("3. microtâche (promise)"));
queueMicrotask(() => console.log("3b. microtâche explicite"));

console.log("2. synchrone");
// Ordre : 1, 2, 3, 3b, 4
// Même avec un délai de 0, setTimeout passe APRÈS toutes les microtâches.
```

### Démonstration : les microtâches affament les macrotâches

```js
// DANGER : boucle de microtâches = event loop bloquée, timers jamais exécutés
function boucleInfinie() {
  Promise.resolve().then(boucleInfinie); // ne jamais faire ça
}
```

### Ce qui bloque / ne bloque pas

| Bloque l'event loop | Ne bloque pas |
|---|---|
| Boucle `for` de 10M d'itérations | `await fetch(...)` |
| `JSON.parse` d'un fichier de 500 Mo | `fs.readFile` (async) |
| `crypto.pbkdf2Sync` | `setTimeout` |
| Regex catastrophique | Workers threads |

> **Règle sysadmin** : en Node, jamais de synchrone lourd dans le chemin d'une requête/API. Déporter vers worker threads ou processus fils.

---

## 22. Callbacks et callback hell

```js
// Callback : fonction passée en argument, appelée "quand c'est prêt"
const fs = require("fs"); // CommonJS, voir section 28
fs.readFile("/etc/hosts", "utf8", (err, data) => {
  if (err) { console.error(err); return; } // convention error-first !
  console.log(data);
});
```

### Le callback hell (pyramide de la mort)

```js
// À NE PAS FAIRE — illisible, erreurs mal propagées
connexion(db, (err, conn) => {
  if (err) return gerer(err);
  requete(conn, "SELECT ...", (err, rows) => {
    if (err) return gerer(err);
    traiter(rows, (err, res) => {
      if (err) return gerer(err);
      envoyer(res, (err) => {
        if (err) return gerer(err);
        console.log("fini");
      });
    });
  });
});
```

### Sorties du callback hell (historique → moderne)

1. **Nommer les fonctions** au lieu d'imbriquer des anonymes.
2. **Promisifier** : `const { promisify } = require("util"); const readFile = promisify(fs.readFile);`
3. **Promises / async-await** (sections 23-24) : la solution moderne.

```js
// Promisification manuelle (à connaître pour les vieilles APIs)
function delai(ms) {
  return new Promise((resolve) => setTimeout(() => resolve(ms), ms));
}
```

---

## 23. Promises en profondeur

Une Promise = valeur **future** avec 3 états : `pending` → `fulfilled` (valeur) ou `rejected` (raison). État **définitif** une fois résolu.

```js
const promesse = new Promise((resolve, reject) => {
  // exécuteur appelé IMMÉDIATEMENT (synchrone !)
  setTimeout(() => {
    Math.random() > 0.5 ? resolve("OK") : reject(new Error("KO"));
  }, 100);
});

promesse
  .then(valeur => { console.log(valeur); return valeur + "!"; }) // chaînage
  .then(v2 => console.log(v2))
  .catch(err => console.error("capturé :", err.message))  // attrape tout le chaînage
  .finally(() => console.log("toujours exécuté"));
```

### Création rapide

```js
Promise.resolve(42);           // déjà fulfilled
Promise.reject(new Error("x"));// déjà rejected
// await 42;                   // await accepte aussi les non-promises
```

### Combinateurs

```js
const p1 = fetch("https://api1.example.com");
const p2 = fetch("https://api2.example.com");

// Tous doivent réussir — rejette dès le premier échec
await Promise.all([p1, p2]);

// Tous se terminent (succès ou échec) — idéal pour bilans
const resultats = await Promise.allSettled([p1, p2]);
// [{status:"fulfilled", value:...}, {status:"rejected", reason:...}]

// Le premier qui se termine (succès ou échec)
await Promise.race([p1, timeout(5000)]);

// Le premier SUCCÈS — rejette seulement si TOUS échouent (AggregateError, ES2021)
await Promise.any([miroir1, miroir2, miroir3]);
```

### Timeout sur une promise (pattern indispensable)

```js
function avecTimeout(promise, ms, message = "timeout") {
  const t = new Promise((_, reject) =>
    setTimeout(() => reject(new Error(`${message} après ${ms}ms`)), ms));
  return Promise.race([promise, t]);
}
await avecTimeout(fetch(url), 5000);
```

### Retry avec backoff exponentiel (pattern sysadmin)

