---
id: collect-261001-rattrapage/rattrapage/javascript-guide-6
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [1219, 1488]
sha256: ec2ab0072e7a128b122c28edd6269ba8f115bd5854b7fc4dc46034f049833c1c
---

# JavaScript — Guide ultra-complet

```js
async function retry(fn, { tentatives = 3, delaiMs = 1000, facteur = 2 } = {}) {
  let derniereErreur;
  for (let i = 1; i <= tentatives; i++) {
    try { return await fn(); }
    catch (err) {
      derniereErreur = err;
      if (i < tentatives) {
        const attente = delaiMs * facteur ** (i - 1);
        console.warn(`Tentative ${i}/${tentatives} échouée, nouvel essai dans ${attente}ms`);
        await new Promise(r => setTimeout(r, attente));
      }
    }
  }
  throw derniereErreur;
}
// await retry(() => fetch("https://api.instable.example.com/status"));
```

### Pièges promises

- Oublier `return` dans un `.then` → le chaînage reçoit `undefined`.
- `new Promise` sans `catch` final → `unhandledRejection` (crash possible en Node).
- L'exécuteur est **synchrone** : `new Promise(() => { throw ... })` rejette immédiatement.
- `.then` retourne toujours une nouvelle promise (chaînage infini possible).

---

## 24. async / await en profondeur

`async` = fonction qui retourne **toujours** une promise. `await` = pause jusqu'à résolution (sans bloquer le thread).

```js
async function recupererStatut(url) {
  try {
    const reponse = await fetch(url);          // pause ici
    if (!reponse.ok) throw new HttpError(reponse.status, reponse.statusText);
    const data = await reponse.json();         // pause ici
    return data;
  } catch (err) {
    console.error("échec :", err.message);
    throw err; // re-propagation si l'appelant doit gérer
  }
}
```

### Séquentiel vs parallèle

```js
// SÉQUENTIEL (lent : 3 × 1s = 3s) — seulement si dépendances entre appels
const a = await tache(1);
const b = await tache(2);

// PARALLÈLE (rapide : max(1s) = 1s) — indépendants
const [ra, rb] = await Promise.all([tache(1), tache(2)]);

// Parallèle avec limite de concurrence (pattern sysadmin : ne pas DoS la cible)
async function pool(taches, limite) {
  const resultats = [];
  const enCours = new Set();
  for (const t of taches) {
    const p = t().then(r => { enCours.delete(p); return r; });
    enCours.add(p); resultats.push(p);
    if (enCours.size >= limite) await Promise.race(enCours);
  }
  return Promise.all(resultats);
}
// await pool(urls.map(u => () => fetch(u)), 5); // 5 requêtes max en parallèle
```

### async dans les boucles — LE piège

```js
const urls = ["https://a.example.com", "https://b.example.com"];

// BUG : forEach n'attend pas — les promises partent en "fire and forget"
urls.forEach(async (u) => { await fetch(u); console.log("fini", u); });
console.log("ceci s'affiche AVANT les fetch");

// CORRECT séquentiel : for...of
for (const u of urls) { await fetch(u); }

// CORRECT parallèle : map + Promise.all
await Promise.all(urls.map(async (u) => { await fetch(u); }));
```

### Top-level await (ES2022, modules uniquement)

```js
// En module ESM : await hors fonction autorisé
const config = await fetch("/config.json").then(r => r.json());
```

### Fonction async = toujours une promise

```js
async function f() { return 42; }
f() instanceof Promise; // true
f().then(console.log);  // 42
```

---

## 25. Fetch : GET, POST, erreurs

`fetch` = API native (navigateur + Node 18+). Retourne une promise vers un objet `Response`.

```js
// GET simple
const res = await fetch("https://api.example.com/serveurs");
console.log(res.status); // 200
console.log(res.ok);     // true si 200-299
console.log(res.headers.get("content-type"));
const data = await res.json();  // ou .text(), .blob(), .arrayBuffer(), .formData()
```

### L'erreur n°1 : fetch ne rejette PAS sur 404/500

```js
// fetch ne rejette que sur échec RÉSEAU (DNS, connexion refusée, CORS...)
// Un 404 ou 500 donne une promise FULFILLED avec ok=false !
const r = await fetch("https://api.example.com/inexistant");
if (!r.ok) throw new HttpError(r.status, `HTTP ${r.status} sur ${r.url}`);
```

### POST JSON

```js
const reponse = await fetch("https://api.example.com/equipements", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Authorization": "Bearer " + process.env.API_TOKEN, // jamais en dur !
  },
  body: JSON.stringify({ hostname: "SW9", ip: "10.0.0.9" }),
  signal: AbortSignal.timeout(8000), // timeout natif (Node 18+ / navigateurs récents)
});
if (!reponse.ok) throw new HttpError(reponse.status, await reponse.text());
const cree = await reponse.json();
```

### Annulation explicite

```js
const ctrl = new AbortController();
const timer = setTimeout(() => ctrl.abort(), 5000); // annule après 5s
try {
  const r = await fetch(url, { signal: ctrl.signal });
} catch (e) {
  if (e.name === "AbortError") console.log("requête annulée");
} finally { clearTimeout(timer); }
```

### Client robuste complet (GET + erreurs + timeout + retry)

```js
class HttpError extends Error {
  constructor(status, message) { super(message); this.name = "HttpError"; this.status = status; }
}

async function apiGet(url, { timeoutMs = 8000, retries = 2 } = {}) {
  let err;
  for (let i = 0; i <= retries; i++) {
    try {
      const res = await fetch(url, { signal: AbortSignal.timeout(timeoutMs) });
      if (!res.ok) throw new HttpError(res.status, `${res.status} ${res.statusText}`);
      return await res.json();
    } catch (e) {
      err = e;
      const rejouable = e.name === "AbortError" || e.name === "TypeError" ||
                        (e instanceof HttpError && e.status >= 500);
      if (!rejouable || i === retries) throw e;
      await new Promise(r => setTimeout(r, 1000 * 2 ** i)); // backoff
    }
  }
  throw err;
}
```

### Tableau des méthodes Response

| Méthode | Usage |
|---|---|
| `res.json()` | corps JSON → objet |
| `res.text()` | corps → chaîne |
| `res.blob()` | fichier binaire (navigateur) |
| `res.arrayBuffer()` | binaire brut |
| `res.formData()` | formulaire multipart |

---

## 26. JSON

```js
const obj = { hostname: "R1", ports: [22, 443], actif: true };

// Sérialiser
JSON.stringify(obj);                    // '{"hostname":"R1",...}'
JSON.stringify(obj, null, 2);            // indenté (lisible, logs/fichiers)
JSON.stringify(obj, ["hostname"]);       // whitelist de clés
JSON.stringify(obj, (k, v) => k === "mdp" ? "***" : v); // replacer : filtrer les secrets !

// Désérialiser — TOUJOURS dans try/catch (données externes = hostiles)
try {
  const data = JSON.parse(texteRecu);
} catch (e) {
  console.error("JSON invalide :", e.message);
}
```

### Limites de JSON

- Pas de `undefined`, fonctions, `Symbol` (ignorés ou `null` dans les tableaux).
- `Date` → chaîne ISO (re-parser à la main si besoin).
- `BigInt` → **TypeError** (`JSON.stringify(10n)` échoue !).
- Références circulaires → TypeError.
- `NaN`/`Infinity` → `null`.

```js
// Contourner BigInt
JSON.stringify({ id: 10n }, (k, v) => typeof v === "bigint" ? v.toString() : v);
```

---

## 27. Timers

```js
// Exécution différée
const id1 = setTimeout(() => console.log("une fois"), 1000);
clearTimeout(id1); // annuler

// Répétition
const id2 = setInterval(() => console.log("tick"), 5000);
clearInterval(id2);

// Immédiat (différé au prochain tour de boucle)
setImmediate(() => console.log("tour suivant")); // Node uniquement
queueMicrotask(() => console.log("microtâche")); // avant les timers !

// Timers promisifiés (Node)
const { setTimeout: attendre } = require("timers/promises");
await attendre(1000); // await sleep(1000)
```

### Dérive de setInterval

`setInterval` ne garantit pas l'heure exacte (dérive si le callback est long). Pour des sondes régulières, préférer le `setTimeout` récursif :

```js
function sonde() {
  const debut = Date.now();
  // ... mesure ...
  const duree = Date.now() - debut;
  setTimeout(sonde, Math.max(0, 5000 - duree)); // compense la durée d'exécution
}
sonde();
```

### Node : ne pas bloquer la sortie

```js
const t = setInterval(collecte, 60000);
t.unref(); // le processus peut se terminer même si le timer est actif
```

---

## 28. Modules ESM vs CommonJS

