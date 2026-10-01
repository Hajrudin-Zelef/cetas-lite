---
id: collect-261001-rattrapage/rattrapage/javascript-guide-4
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["1970-01-01", "2026-09-26"]
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [717, 967]
sha256: 17430380b6b9c0573b9e89bab49b97dba1d4ab3c5a04e1fecbfd1bce3189bdcb
---

# JavaScript — Guide ultra-complet

// Remplacement
"aaa".replace("a", "b");    // "baa" (première occurrence)
"aaa".replaceAll("a", "b"); // "bbb" (ES2021)
"ERR 12".replace(/(\d+)/, "[$1]"); // "ERR [12]" (regex + groupe)

// Découpage / assemblage
"a,b,c".split(",");         // ["a","b","c"]
["a","b"].join("-");        // "a-b"

// Template literals : multiligne + interpolation + taggés
const lignes = `
  ligne1
  ligne2 ${2 + 2}
`;
```

### Cas sysadmin : normaliser des sorties

```js
function parseLigneDf(ligne) {
  // "/dev/sda1       100G   60G   40G  60% /"
  const [fs, taille, used, dispo, pct, mount] = ligne.trim().split(/\s+/);
  return { fs, taille, used, dispo, usage: parseInt(pct), mount };
}
```

---

## 14. Nombres et Math

```js
// Littéraux
const a = 1_000_000;   // séparateur (ES2021) : lisibilité
const b = 0xFF;        // 255 hexadécimal
const c = 0b1010;      // 10 binaire
const d = 0o755;       // 493 octal

// Arrondis
Math.floor(4.9);  // 4    Math.ceil(4.1);   // 5
Math.round(4.5);  // 5    Math.trunc(-4.9); // -4 (ES2015, vers zéro)
(4.567).toFixed(2);      // "4.57" (chaîne !)
Number("4.57").toPrecision(3);

// Aléatoire
Math.floor(Math.random() * 100); // 0..99
// crypto.getRandomValues() pour du vrai aléatoire (tokens, voir section 50)

// Min/Max, puissance, racine
Math.max(3, 9, 1);  Math.min(...[3, 9, 1]);
Math.pow(2, 10);  2 ** 10;
Math.sqrt(144);   Math.cbrt(27);

// Vérifications
Number.isInteger(4);       // true
Number.isNaN("abc" / 1);   // true (ne coerce pas, contrairement au global isNaN)
Number.isFinite(1 / 0);    // false
Number.EPSILON;            // pour comparer des flottants
Math.abs(0.1 + 0.2 - 0.3) < Number.EPSILON; // true
```

---

## 15. Dates

```js
const maintenant = new Date();
maintenant.toISOString();      // "2026-09-26T23:34:27.000Z" (UTC, idéal logs)
maintenant.toLocaleString("fr-FR"); // "26/09/2026 23:34:27" (affichage)
maintenant.getTime();          // timestamp ms depuis 1970-01-01

// Construction : ATTENTION mois 0-indexé !
new Date(2026, 8, 26);         // 26 septembre 2026 (mois 8 = septembre !)
new Date("2026-09-26T10:00:00Z"); // ISO : recommandé
Date.now();                    // timestamp ms (le plus rapide)

// Calculs : travailler en ms, jamais en "jours" naïfs (DST !)
const UN_JOUR_MS = 24 * 60 * 60 * 1000;
const demain = new Date(Date.now() + UN_JOUR_MS);

// Formater proprement : Intl (ES2015+)
new Intl.DateTimeFormat("fr-FR", {
  dateStyle: "full", timeStyle: "short", timeZone: "Europe/Paris"
}).format(maintenant); // "samedi 26 septembre 2026 à 23:34"

// Durée entre deux dates
const debut = Date.now();
// ... traitement ...
const dureeMs = Date.now() - debut;

// Temporal (API moderne, en cours de standardisation — signalée, pas encore partout)
// const zdt = Temporal.Now.zonedDateTimeISO(); // futur remplaçant de Date
```

### Pièges Date

| Piège | Explication |
|---|---|
| Mois 0-indexé | `new Date(2026, 8, 1)` = septembre, pas août |
| `getDay()` vs `getDate()` | jour semaine (0=dimanche) vs jour du mois |
| Fuseaux | `new Date("2026-09-26")` = minuit **UTC** ; `new Date(2026,8,26)` = minuit **local** |
| Mutable | les setters (`setHours`...) modifient l'objet |

---

## 16. this — les 4 règles

`this` dépend du **site d'appel**, pas du site de définition. 4 règles, par ordre de priorité :

```js
// RÈGLE 1 — new : `this` = nouvel objet
function Serveur(host) { this.host = host; }
const s = new Serveur("10.0.0.1"); // this → s

// RÈGLE 2 — appel méthode : `this` = objet avant le point
const obj = {
  nom: "R1",
  qui() { return this.nom; }
};
obj.qui(); // "R1" — this → obj

// RÈGLE 3 — appel explicite : call / apply / bind
function presenter(role) { return `${this.nom} (${role})`; }
presenter.call({ nom: "Zelef" }, "admin");   // "Zelef (admin)"
presenter.apply({ nom: "Zelef" }, ["admin"]); // idem, args en tableau
const f = presenter.bind({ nom: "Zelef" }); // this figé
f("admin");

// RÈGLE 4 — défaut : undefined (strict) ou globalThis (sloppy)
function seule() { return this; }
seule(); // undefined en module/strict ; globalThis en sloppy
```

### Fonctions fléchées : PAS de `this` propre

```js
const equipe = {
  nom: "Systèmes",
  membres: ["a", "b"],
  lister() {
    // `this` fléchée = celui de lister() → equipe. Parfait pour les callbacks !
    this.membres.forEach(m => console.log(`${this.nom} : ${m}`));
  },
  // MAUVAIS : méthode fléchée → this = portée englobante, pas l'objet
  // mauvais: () => this.nom
};
```

### Cas concret : perdre `this` dans un callback

```js
class Sonde {
  constructor() { this.valeur = 0; }
  demarrer() {
    // setInterval(function() { this.valeur++; }, 1000); // BUG : this = timeout/global
    setInterval(() => { this.valeur++; }, 1000); // OK : fléchée hérite du this
    // ou : setInterval(function() { this.valeur++; }.bind(this), 1000);
  }
}
```

> **Pense-bête** : dans une méthode, callback → fléchée. Pour fixer un `this` → `bind`.

---

## 17. Prototypes

Avant les classes (ES2015), JS utilisait les prototypes. Les classes ne sont que du sucre syntaxique par-dessus — comprendre le prototype explique 80 % des bizarreries.

```js
function Animal(nom) { this.nom = nom; }
// Méthode sur le PROTOTYPE : partagée par toutes les instances (1 seule copie en mémoire)
Animal.prototype.parler = function() { return `${this.nom} fait un bruit`; };

const rex = new Animal("Rex");
rex.parler(); // "Rex fait un bruit"

// Chaîne de prototype : rex → Animal.prototype → Object.prototype → null
Object.getPrototypeOf(rex) === Animal.prototype; // true
rex.hasOwnProperty("nom");        // true (propre)
rex.hasOwnProperty("parler");     // false (hérité du prototype)

// __proto__ : accesseur historique (préférer Object.getPrototypeOf / Object.create)
const chien = Object.create(Animal.prototype);
Animal.call(chien, "Médor"); // simule le constructeur
```

### Pourquoi c'est important

- `Array.prototype.map`, `String.prototype.trim`… : toutes les méthodes natives vivent sur des prototypes.
- Ajouter une méthode au prototype = disponible sur toutes les instances existantes.
- **Ne jamais** modifier `Object.prototype` (pollution globale) ni les prototypes natifs en lib partagée.

---

## 18. Classes

```js
class Equipement {
  // Champ public (ES2022)
  statut = "inconnu";
  // Champ privé (ES2022) : inaccessible de l'extérieur
  #serie = "SN-" + Math.random().toString(36).slice(2, 8).toUpperCase();

  constructor(hostname, ip) {
    this.hostname = hostname;
    this.ip = ip;
  }

  // Méthode d'instance (sur le prototype)
  ping() { return `ping ${this.ip}...`; }

  // Getter / setter
  get fiche() { return `${this.hostname} [${this.ip}]`; }
  set fiche(v) { [this.hostname, this.ip] = v.split("|"); }

  // Méthode statique (sur la classe, pas les instances)
  static type() { return "équipement réseau"; }

  // Méthode privée
  #audit() { return this.#serie; }
  serieMasquee() { return this.#audit().slice(0, 5) + "****"; }
}

const sw1 = new Equipement("SW1", "10.0.0.2");
sw1.ping();            // "ping 10.0.0.2..."
sw1.fiche;             // "SW1 [10.0.0.2]" (getter, sans parenthèses)
Equipement.type();     // "équipement réseau"
// sw1.#serie;         // SyntaxError : privé
```

### Tableau : champs et méthodes

| Syntaxe | Accès | Exemple |
|---|---|---|
| `champ = v` | public, par instance | `statut = "ok"` |
| `#champ` | privé (classe uniquement) | `#serie` |
| `static champ` | sur la classe | `static VERSION = 2` |
| `methode()` | prototype → instances | `ping()` |
| `static methode()` | classe uniquement | `type()` |
| `get` / `set` | propriété calculée | `get fiche()` |

---

## 19. Héritage et polymorphisme

```js
class Equipement {  // voir section 18
  constructor(hostname, ip) { this.hostname = hostname; this.ip = ip; }
  decrire() { return `${this.hostname} (${this.ip})`; }
}

