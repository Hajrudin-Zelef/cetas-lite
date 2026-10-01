---
id: collect-261001-rattrapage/rattrapage/javascript-guide-3
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [456, 716]
sha256: 9e8c2ee80c04a6060537fbaf0f6f4cdc3559fb48a8f8fb4bdb7019a3319dae41
---

# JavaScript — Guide ultra-complet

```js
function creerCompteur() {
  let n = 0; // variable "privée", capturée
  return {
    incrementer() { n++; return n; },
    valeur() { return n; }
  };
}
const c1 = creerCompteur();
c1.incrementer(); // 1
c1.incrementer(); // 2
const c2 = creerCompteur();
c2.valeur();      // 0 — chaque appel a SA propre closure
```

### Cas d'usage n°1 : fabrique paramétrée

```js
function creerLogger(prefixe) {
  return (message) => console.log(`[${new Date().toISOString()}] [${prefixe}] ${message}`);
}
const logReseau = creerLogger("RESEAU");
const logSys = creerLogger("SYSTEME");
logReseau("lien up");   // [2026-...] [RESEAU] lien up
```

### Cas d'usage n°2 : encapsulation (données privées avant les champs #)

```js
function creerSession(utilisateur) {
  let token = Math.random().toString(36).slice(2); // privé
  return {
    utilisateur,
    getToken: () => token.slice(0, 4) + "****",  // accès contrôlé
    regenerer() { token = Math.random().toString(36).slice(2); }
  };
}
```

### Cas d'usage n°3 : mémoïsation

```js
function memoiser(fn) {
  const cache = new Map(); // capturé par la closure
  return (...args) => {
    const cle = JSON.stringify(args);
    if (cache.has(cle)) return cache.get(cle);
    const res = fn(...args);
    cache.set(cle, res);
    return res;
  };
}
const dnsLent = memoiser((host) => { /* résolution coûteuse */ return host; });
```

> **Piège mémoire** : une closure retient TOUTE sa portée. Ne pas capturer de gros objets inutilement dans des callbacks longue durée (timers, listeners).

---

## 9. Paramètres avancés des fonctions

```js
// Valeurs par défaut (ES2015) — évaluées à chaque appel
function connecter({ host = "localhost", port = 22, timeout = 5000 } = {}) {
  console.log(`ssh ${host} -p ${port} (timeout ${timeout}ms)`);
}
connecter();                    // ssh localhost -p 22 (timeout 5000ms)
connecter({ host: "10.0.0.5" });

// Rest : regroupe les arguments restants en tableau
function somme(...nombres) { return nombres.reduce((a, b) => a + b, 0); }
somme(1, 2, 3, 4); // 10

// arguments (fonctions classiques uniquement) : objet array-like
function legacy() { return Array.from(arguments).join("-"); }

// Paramètres nommés via destructuration (idiome sysadmin : options lisibles)
async function requete(url, { methode = "GET", headers = {}, timeoutMs = 10000, retries = 3 } = {}) {
  /* ... */
}
```

### Bonnes pratiques

- Max ~3 paramètres positionnels ; au-delà → objet d'options.
- Valeurs par défaut pour tout ce qui a une valeur "normale".
- Ne jamais muter les paramètres d'entrée (effet de bord surprise).

---

## 10. Tableaux

```js
const srv = ["web1", "web2", "db1"];

// Accès / bases
srv[0]; srv.length; srv.at(-1); // "db1" (.at() ES2022 : index négatif)
```

### Méthodes mutantes vs non-mutantes

| Mutantes (modifient le tableau) | Non-mutantes (retournent un nouveau) |
|---|---|
| `push`, `pop`, `shift`, `unshift` | `map`, `filter`, `slice`, `concat` |
| `splice`, `sort`, `reverse`, `fill` | `toSorted`, `toReversed`, `toSpliced` (ES2023) |

```js
const nums = [3, 1, 4, 1, 5];

// Transformation (chaînables — le cœur du JS moderne)
nums.filter(n => n > 2)   // [3, 4, 5]
    .map(n => n * 10)     // [30, 40, 50]
    .sort((a, b) => a - b);

nums.reduce((acc, n) => acc + n, 0); // 14 (somme)
nums.find(n => n > 3);    // 4 (premier match)
nums.findIndex(n => n > 3);// 2
nums.some(n => n > 4);    // true (au moins un)
nums.every(n => n > 0);   // true (tous)
nums.includes(4);         // true (avec SameValueZero : trouve NaN !)
nums.indexOf(1);          // 1 (premier index)

// Aplatir / grouper
[[1, 2], [3]].flat();              // [1, 2, 3]
[1, 2, 3].flatMap(x => [x, x * 2]); // [1,2,2,4,3,6]
Object.groupBy(srv, s => s.startsWith("web") ? "web" : "db"); // ES2024

// Tri : TOUJOURS avec comparateur pour les nombres !
[10, 2, 30].sort();            // [10, 2, 30] — tri LEXICOGRAPHIQUE, bug classique
[10, 2, 30].sort((a, b) => a - b); // [2, 10, 30]

// Copie
const copie = [...nums];          // spread (superficielle)
const copie2 = structuredClone(obj); // profonde (navigateur + Node 17+)
```

### Cas sysadmin : pipeline de traitement

```js
const lignes = [
  "10.0.0.1 - - [26/Sep/2026:10:00:01] \"GET /api/status 200\"",
  "10.0.0.2 - - [26/Sep/2026:10:00:02] \"GET /api/data 500\"",
];
const erreurs = lignes
  .map(l => l.match(/"(GET|POST) (\S+) (\d{3})"/))
  .filter(Boolean)
  .filter(([, , , code]) => code.startsWith("5"))
  .map(([, methode, url]) => ({ methode, url }));
console.log(erreurs); // [{ methode: 'GET', url: '/api/data' }]
```

---

## 11. Objets

```js
const routeur = {
  hostname: "R1",
  ip: "10.0.0.1",
  "nom avec espace": "core",   // clé entre guillemets si spéciale
  interfaces: ["ge-0/0/0"],
  redemarrer() { console.log(`${this.hostname} reboot...`); }, // méthode
  ["vlan" + 10]: "users",      // clé calculée (ES2015)
};

// Accès
routeur.hostname;      // notation point (préférée)
routeur["ip"];         // notation crochet (clé dynamique)
const cle = "hostname";
routeur[cle];          // "R1"

// Ajout / suppression
routeur.os = "JunOS";  // ajout
delete routeur.os;     // suppression

// Parcourir
Object.keys(routeur);    // ["hostname", "ip", ...]
Object.values(routeur);  // valeurs
Object.entries(routeur); // [[clé, valeur], ...]
for (const [k, v] of Object.entries(routeur)) { /* ... */ }

// Fusion / copie superficielle
const clone = { ...routeur, ip: "10.0.0.2" };
const fusion = Object.assign({}, routeur, { site: "Paris" });

// Descripteurs : verrouiller un objet
const CONSTANTES = Object.freeze({ SEUIL_CRITIQUE: 90, SEUIL_WARN: 70 });
// CONSTANTES.SEUIL_CRITIQUE = 95; // silencieux / TypeError en strict
```

### Getters / setters

```js
const sonde = {
  _temp: 21,
  get temp() { return `${this._temp}°C`; },
  set temp(v) { if (v < -50 || v > 80) throw new RangeError("plage invalide"); this._temp = v; }
};
sonde.temp = 25;
console.log(sonde.temp); // "25°C"
```

### Vérifier l'existence d'une propriété

```js
"ip" in routeur;                    // true (chaîne de prototype incluse)
routeur.hasOwnProperty("ip");       // true (propre uniquement)
Object.hasOwn(routeur, "ip");       // true (ES2022, recommandé)
routeur.ip !== undefined;           // approximatif (échoue si valeur = undefined)
```

---

## 12. Destructuration, spread, rest

```js
// Tableaux
const [premier, deuxieme, ...reste] = ["a", "b", "c", "d"];
// premier="a", deuxieme="b", reste=["c","d"]
const [x = 1, y = 2] = [];  // valeurs par défaut : x=1, y=2
let p = 1, q = 2; [p, q] = [q, p]; // swap sans variable temporaire

// Objets (+ renommage + défauts + imbriqué)
const { hostname: nom, ip = "0.0.0.0", reseau: { masque } = {} } = routeur;

// Paramètres de fonction
function afficher({ hostname, ip }) { console.log(hostname, ip); }

// Spread : étaler
const t1 = [1, 2], t2 = [3, 4];
const t3 = [...t1, ...t2];          // [1,2,3,4]
const o2 = { ...routeur };          // copie superficielle
Math.max(...[4, 9, 2]);             // 9 (spread en appel)

// Rest : regrouper (inverse du spread)
const { hostname: h, ...autres } = routeur; // autres = tout sauf hostname
```

> **Attention** : spread/destructuration = copie **superficielle** (les objets imbriqués restent partagés).

---

## 13. Chaînes de caractères

```js
const log = "  ERR  connexion timeout sur eth0  ";

// Nettoyage / inspection
log.trim();                 // "ERR  connexion timeout sur eth0"
log.toLowerCase();          // minuscules
log.includes("timeout");    // true
log.startsWith("  ERR");    // true
log.endsWith("eth0  ");     // true
log.indexOf("timeout");     // position ou -1

// Extraction
log.slice(2, 5);            // "ERR" (indices négatifs OK)
log.substring(2, 5);        // "ERR" (pas de négatifs)
"eth0".padStart(8, "0");    // "0000eth0"
"42".padEnd(5, ".");        // "42..."

