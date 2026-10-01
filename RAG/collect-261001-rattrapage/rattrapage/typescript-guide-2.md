---
id: collect-261001-rattrapage/rattrapage/typescript-guide-2
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [187, 427]
sha256: 7b3bb42cbcea699310c9cd9baeb1f1396d89a114cb832f1bc203346e5aa72b33
---

# TypeScript — Guide ultra-complet

| Sous-option | Ce qu'elle interdit | Exemple d'erreur évitée |
|---|---|---|
| `strictNullChecks` | `null`/`undefined` assignés n'importe où | Appeler `.length` sur une valeur peut-être `null` |
| `noImplicitAny` | Paramètre au type `any` implicite | `function f(x) {}` → `x` est `any` caché |
| `strictFunctionTypes` | Contravariance des paramètres de fonction | Assigner un callback incompatible |
| `strictBindCallApply` | `bind`/`call`/`apply` non vérifiés | Mauvais arguments via `.call()` |
| `strictPropertyInitialization` | Propriété de classe non initialisée | `this.host` utilisé avant affectation |
| `noImplicitThis` | `this: any` implicite | `this` dans une fonction détachée |
| `alwaysStrict` | Émet `"use strict"` dans le JS | Oubli du mode strict JS |

**Deux options strictes hors du pack `strict` (à activer à la main) :**

```typescript
// "noUncheckedIndexedAccess": true
const tab: string[] = ["a", "b"];
const x = tab[10]; // type: string | undefined (au lieu de string)
// → vous force à gérer le cas undefined. Précieux quand on parse des données.

// "exactOptionalPropertyTypes": true
interface Config { timeout?: number; }
const c: Config = { timeout: undefined }; // ERREUR : undefined n'est pas assignable
// → distingue "propriété absente" de "propriété = undefined".
```

**Règle d'or** : commencez tout nouveau projet avec `"strict": true`. Sur un projet existant, migrez progressivement (`// @ts-nocheck` en haut des fichiers non migrés, puis fichier par fichier).

---

## 6. Types primitifs : le socle

```typescript
const nom: string = "onduleur-01";
const puissance: number = 40;          // pas de int/float : tout est number
const enLigne: boolean = true;
const gros: bigint = 9007199254740993n; // grands entiers (target ES2020+)
const id: symbol = Symbol("id");        // identifiants uniques
const rien: null = null;
const indefini: undefined = undefined;
```

**Inférence** : TypeScript devine le type quand c'est évident. N'annotez que quand ça apporte de l'information :

```typescript
let compteur = 0;            // inféré: number — pas besoin d'annoter
let etat: "ok" | "ko" = "ok"; // annotation utile : union de littéraux
```

**`number` et ses pièges** (le type ne protège pas de tout) :

```typescript
console.log(0.1 + 0.2);      // 0.30000000000000004 — arithmétique flottante
Number.isSafeInteger(2 ** 53); // false → passer à bigint si besoin
```

> Pense-bête : `string`, `number`, `boolean` s'écrivent **en minuscules**. `String`, `Number`, `Boolean` (majuscules) sont les objets wrappers — ne les utilisez jamais comme types.

---

## 7. string, number, boolean : les subtilités utiles

**Template literals typés** (TS 4.1+) : des strings dont la forme est vérifiée.

```typescript
type Fqdn = `${string}.${string}`;      // "srv-01.lan" OK
type IpV4 = `${number}.${number}.${number}.${number}`;

function ping(hote: Fqdn): void { /* ... */ }
ping("srv-01.lan");   // OK
ping("srv-01");       // Erreur : ne matche pas `${string}.${string}`
```

**Conversions explicites** (TypeScript ne convertit jamais implicitement) :

```typescript
const portStr: string = "8080";
const port: number = Number(portStr);        // "8080" -> 8080
const ok: boolean = portStr === "8080";     // comparaison stricte

// Depuis une variable d'environnement (toujours string | undefined !)
const portEnv: number = Number(process.env.PORT ?? "3000");
```

**Opérateur `??` vs `||`** — capital en scripting :

```typescript
const timeout = config.timeout ?? 5000;  // défaut seulement si null/undefined
const timeout2 = config.timeout || 5000; // défaut aussi si 0 ! dangereux
// Avec ||, un timeout configuré à 0 serait écrasé. Préférez ??.
```

---

## 8. Arrays typés

```typescript
const hosts: string[] = ["srv-01", "srv-02"];
const ports: Array<number> = [22, 80, 443]; // syntaxe équivalente

hosts.push("srv-03");   // OK
hosts.push(42);         // Erreur : number non assignable à string
```

**Tableaux en lecture seule** (évite les mutations accidentelles) :

```typescript
const LISTE: readonly string[] = ["a", "b"];
LISTE.push("c"); // Erreur : push n'existe pas sur readonly
```

**Itérations typées** — l'inférence fait le travail :

```typescript
const onduleurs = [
  { nom: "ups-01", charge: 42 },
  { nom: "ups-02", charge: 67 },
];
// onduleurs: { nom: string; charge: number }[]

for (const o of onduleurs) {
  console.log(o.nom.toUpperCase()); // o.nom est string : autocomplétion OK
}

const charges = onduleurs.map((o) => o.charge * 2); // number[]
const critiques = onduleurs.filter((o) => o.charge > 80); // même type d'objet[]
```

> Avec `noUncheckedIndexedAccess` (section 5), `onduleurs[0]` vaut `{...} | undefined` : pensez au test ou au `!` assumé.

---

## 9. Tuples : des tableaux à forme fixe

Un tuple = un tableau dont **chaque position a son type**. Idéal pour les paires clé/valeur, les coordonnées, les retours multiples.

```typescript
type EntreeInventaire = [nom: string, ip: string, enLigne: boolean];
const srv: EntreeInventaire = ["srv-01", "10.0.0.1", true];
// Les labels (nom:, ip:) sont purement documentaires.

srv[0] = "srv-02"; // OK : string
srv[1] = 123;      // Erreur : number non assignable à string

// Destructuration typée
const [nom, ip, enLigne] = srv; // string, string, boolean

// Tuple à longueur variable (reste typé)
type LigneCsv = [string, ...number[]];
const ligne: LigneCsv = ["ups-01", 42, 67, 12]; // OK
```

**Cas d'usage sysadmin** : retour de fonction multi-valeurs sans créer d'interface.

```typescript
function parseLigneInventaire(ligne: string): [string, string] | null {
  const [nom = "", ip = ""] = ligne.split(";");
  if (!nom || !ip) return null;
  return [nom, ip];
}
```

---

## 10. Union types : « soit l'un, soit l'autre »

```typescript
type Statut = "en-ligne" | "hors-ligne" | "maintenance";
let s: Statut = "en-ligne";
s = "inconnu"; // Erreur

type Id = string | number;
function afficher(id: Id): void {
  console.log(`ID: ${id}`); // OK : les deux ont toString via template
}
```

**Unions d'objets** — la base des machines à états :

```typescript
type ResultatPing =
  | { ok: true; latenceMs: number }
  | { ok: false; erreur: string };

function traiter(r: ResultatPing): void {
  if (r.ok) {
    console.log(r.latenceMs); // OK : narrowing automatique (section 24)
  } else {
    console.log(r.erreur);    // OK
  }
}
```

**Règle** : sur une union, vous ne pouvez utiliser que les propriétés **communes** à tous les membres sans narrowing préalable.

```typescript
declare const r: ResultatPing;
r.latenceMs; // Erreur : n'existe pas sur le membre { ok: false }
```

---

## 11. Intersection types : « à la fois l'un et l'autre »

```typescript
type Identifiable = { id: string };
type Horodate = { creeLe: Date };

type Equipement = Identifiable & Horodate;
// { id: string; creeLe: Date }

const ups: Equipement = { id: "ups-01", creeLe: new Date() };
```

**Cas d'usage** : composer des comportements (mixins), enrichir un type tiers sans le modifier :

```typescript
import type { Request } from "express";

type RequeteAuthentifiee = Request & { utilisateur: { login: string; roles: string[] } };
// On ajoute l'utilisateur injecté par un middleware d'auth.
```

> **Union vs intersection** : `A | B` = "l'un OU l'autre" (au moins un), `A & B` = "l'un ET l'autre" (les deux). L'intersection de types incompatibles donne `never` : `string & number` est inhabitable.

---

## 12. Types littéraux : des valeurs comme types

```typescript
type NiveauLog = "debug" | "info" | "warn" | "error";
type CodeSortie = 0 | 1 | 2;

function log(niveau: NiveauLog, msg: string): void {
  if (niveau === "error") console.error(msg);
  else console.log(`[${niveau}] ${msg}`);
}
log("info", "démarrage");  // OK
log("fatal", "x");         // Erreur
```

**`const` + littéraux** : une variable `const` garde son littéral, `let` s'élargit.

