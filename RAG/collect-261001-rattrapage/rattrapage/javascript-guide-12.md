---
id: collect-261001-rattrapage/rattrapage/javascript-guide-12
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [2726, 2972]
sha256: 0b2b397476089678cd3a8d620e3e795dcac58b23d1bdaba5aca1ca27231b6e09
---

# JavaScript — Guide ultra-complet

### 2. `NaN !== NaN`
```js
// MAL : if (x === NaN)
// BIEN :
if (Number.isNaN(x)) { /* ... */ }
```

### 3. Oublier le comparateur de `sort`
```js
[10, 2, 30].sort(); // [10, 2, 30] — tri alphabétique !
[10, 2, 30].sort((a, b) => a - b); // [2, 10, 30]
```

### 4. `var` dans une boucle + closure
```js
for (var i = 0; i < 3; i++) setTimeout(() => console.log(i), 0); // 3,3,3
for (let i = 0; i < 3; i++) setTimeout(() => console.log(i), 0); // 0,1,2
```

### 5. Muter un tableau/objet partagé
```js
const defaut = { port: 22 };
function connecter(opts = defaut) { opts.port = 2222; } // mute le défaut pour tous !
connecter(); connecter(); // 2e appel : port déjà à 2222
// BIEN : function connecter(opts = {}) { const cfg = { port: 22, ...opts }; }
```

### 6. `async` dans `forEach`
```js
// forEach n'attend pas les callbacks async → voir section 24
// BIEN : for...of ou Promise.all(urls.map(...))
```

### 7. Oublier que `fetch` ne rejette pas sur HTTP 4xx/5xx
```js
const r = await fetch(url);
if (!r.ok) throw new Error(`HTTP ${r.status}`); // TOUJOURS vérifier !
```

### 8. `this` perdu dans un callback
```js
class A { constructor(){ this.x = 1; } go(){ setTimeout(function(){ console.log(this.x); }, 0); } }
// undefined — BIEN : setTimeout(() => console.log(this.x), 0)
```

### 9. Comparer des objets avec `===`
```js
{ a: 1 } === { a: 1 }; // false — références différentes !
// BIEN : comparer champ à champ ou assert.deepStrictEqual
```

### 10. Hoisting / TDZ
```js
console.log(x); // ReferenceError (TDZ)
let x = 5;
```

### 11. `parseInt` sans base
```js
parseInt("08");      // 8 aujourd'hui, mais...
parseInt("08", 10);  // 8 — TOUJOURS préciser la base
```

### 12. Flottants : `0.1 + 0.2 !== 0.3`
```js
// BIEN : travailler en centimes (entiers) ou comparer avec EPSILON
Math.abs((0.1 + 0.2) - 0.3) < Number.EPSILON; // true
```

### 13. `Array` vide truthy
```js
if ([]) console.log("vrai"); // s'affiche ! — tester .length
if (resultats.length === 0) { /* ... */ }
```

### 14. Oublier `return` dans `.then` / fonction fléchée à bloc
```js
const f = () => { 42; }; // undefined !
const g = () => 42;      // 42
```

### 15. Bloquer l'event loop avec du sync lourd
```js
// const data = fs.readFileSync("2go.log"); // bloque TOUT le serveur
// BIEN : await readFile / streams (sections 31, 34)
```

---

## 52. Cas pratique 1 — parser de logs Apache/nginx

Objectif : lire un access log, compter les codes HTTP, top 10 des IP, détecter les 5xx. Streaming (gros fichiers OK).

```js
// parse-access-log.js — usage : node parse-access-log.js /var/log/nginx/access.log
import { createReadStream } from "fs";
import readline from "readline";

const RE = /^(\S+) \S+ \S+ \[([^\]]+)\] "(\S+) (\S+) \S+" (\d{3}) (\d+|-) "[^"]*" "([^"]*)"/;

async function analyser(chemin) {
  const parCode = new Map();   // "200" -> n
  const parIp = new Map();     // ip -> n
  const erreurs5xx = [];
  let total = 0, ignorees = 0;

  const rl = readline.createInterface({
    input: createReadStream(chemin),
    crlfDelay: Infinity,
  });

  for await (const ligne of rl) {
    const m = ligne.match(RE);
    if (!m) { ignorees++; continue; }
    const [, ip, , methode, url, code] = m;
    total++;
    parCode.set(code, (parCode.get(code) ?? 0) + 1);
    parIp.set(ip, (parIp.get(ip) ?? 0) + 1);
    if (code.startsWith("5")) erreurs5xx.push({ ip, methode, url, code });
  }

  const top = (map, n) => [...map.entries()].sort((a, b) => b[1] - a[1]).slice(0, n);

  console.log(`Lignes : ${total} (ignorées : ${ignorees})`);
  console.log("\nPar code HTTP :");
  for (const [code, n] of [...parCode.entries()].sort())
    console.log(`  ${code} : ${n} (${(n / total * 100).toFixed(1)}%)`);
  console.log("\nTop 10 IP :");
  for (const [ip, n] of top(parIp, 10)) console.log(`  ${ip} : ${n}`);
  console.log(`\nErreurs 5xx : ${erreurs5xx.length}`);
  for (const e of erreurs5xx.slice(0, 20))
    console.log(`  ${e.code} ${e.methode} ${e.url} ← ${e.ip}`);
}

const fichier = process.argv[2];
if (!fichier) { console.error("Usage : node parse-access-log.js <fichier>"); process.exit(1); }
await analyser(fichier);
```

---

## 53. Cas pratique 2 — client API REST avec retry

Objectif : interroger une API d'inventaire avec timeout, retry backoff, et gestion d'erreurs propre.

```js
// api-client.js
class HttpError extends Error {
  constructor(status, message) { super(message); this.name = "HttpError"; this.status = status; }
}

const sleep = (ms) => new Promise(r => setTimeout(r, ms));

export async function apiRequest(url, {
  method = "GET", body = null, token = process.env.API_TOKEN,
  timeoutMs = 8000, retries = 3,
} = {}) {
  let lastError;
  for (let attempt = 1; attempt <= retries + 1; attempt++) {
    try {
      const res = await fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
          ...(token ? { Authorization: `Bearer ${token}` } : {}),
        },
        ...(body ? { body: JSON.stringify(body) } : {}),
        signal: AbortSignal.timeout(timeoutMs),
      });
      if (!res.ok) throw new HttpError(res.status, `${method} ${url} → HTTP ${res.status}`);
      const text = await res.text();
      return text ? JSON.parse(text) : null;
    } catch (err) {
      lastError = err;
      const retryable = err.name === "AbortError" || err.name === "TypeError" ||
        (err instanceof HttpError && (err.status >= 500 || err.status === 429));
      if (!retryable || attempt > retries) throw err;
      const wait = 1000 * 2 ** (attempt - 1);
      console.warn(`[${attempt}/${retries + 1}] ${err.message} — nouvel essai dans ${wait}ms`);
      await sleep(wait);
    }
  }
  throw lastError;
}

// Utilisation
const equipements = await apiRequest("https://inventaire.example.com/api/equipements?site=paris");
const cree = await apiRequest("https://inventaire.example.com/api/equipements",
  { method: "POST", body: { hostname: "SW9", ip: "10.0.0.9" } });
```

---

## 54. Cas pratique 3 — mini-serveur HTTP maison

Objectif : API interne de supervision sans dépendance (santé, liste d'équipements, réception d'alertes).

```js
// mini-api.js — node mini-api.js (PORT=3000)
import http from "http";

const equipements = new Map([
  ["SW1", { ip: "10.0.0.2", statut: "ok" }],
  ["R1",  { ip: "10.0.0.1", statut: "ok" }],
]);
const alertes = [];

function json(res, code, data) {
  res.writeHead(code, { "Content-Type": "application/json; charset=utf-8" });
  res.end(JSON.stringify(data));
}

async function lireBody(req) {
  let corps = "";
  for await (const chunk of req) {
    corps += chunk;
    if (corps.length > 1e6) throw new Error("body trop volumineux"); // anti-DoS basique
  }
  return corps ? JSON.parse(corps) : {};
}

const server = http.createServer(async (req, res) => {
  try {
    const url = new URL(req.url, `http://${req.headers.host}`);
    const route = `${req.method} ${url.pathname}`;

    if (route === "GET /sante") return json(res, 200, { statut: "ok", uptime: Math.round(process.uptime()) });
    if (route === "GET /equipements") return json(res, 200, [...equipements.entries()].map(([nom, f]) => ({ nom, ...f })));
    if (route === "GET /alertes") return json(res, 200, alertes.slice(-50));

    if (route === "POST /alertes") {
      const a = await lireBody(req);
      if (!a.equipement || !a.message) return json(res, 400, { erreur: "equipement et message requis" });
      const alerte = { ...a, recuLe: new Date().toISOString() };
      alertes.push(alerte);
      console.log(`[ALERTE] ${a.equipement} : ${a.message}`);
      return json(res, 201, alerte);
    }

    const m = url.pathname.match(/^\/equipements\/([\w-]+)$/);
    if (m && req.method === "PATCH") {
      const fiche = equipements.get(m[1]);
      if (!fiche) return json(res, 404, { erreur: "inconnu" });
      Object.assign(fiche, await lireBody(req));
      return json(res, 200, { nom: m[1], ...fiche });
    }

