---
id: collect-261001-rattrapage/rattrapage/typescript-guide-7
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [1447, 1692]
sha256: fa4c48d02ac890b1da6c933d1837d6b9d1ae5abf20819e6fd1dfc2579e68df8b
---

# TypeScript — Guide ultra-complet

// Générer des noms de méthodes/event handlers
type HandlerName<E extends string> = `on${Capitalize<E>}`;
type H = HandlerName<"connect">; // "onConnect"
```

**Parsing de strings au niveau des types** avec `infer` :

```typescript
type ExtraireHote<S> = S extends `${string}://${infer Hote}/${string}` ? Hote : never;
type H2 = ExtraireHote<"https://srv-01.lan/api">; // "srv-01.lan"
```

**Cas d'usage** : typer des clés de métriques, des topics MQTT, des routes d'API :

```typescript
type Route = `/api/${"v1" | "v2"}/${"sondes" | "alertes"}`;
const r: Route = "/api/v2/sondes"; // OK
```

> **Signal de version** : template literal types = TS 4.1+. `Capitalize`/`Uncapitalize`/`Uppercase`/`Lowercase` = utilitaires natifs TS 4.8-.

---

## 42. Branded types : des types nominaux dans un monde structurel

TypeScript est **structurel** : deux types de même forme sont interchangeables. Parfois on veut distinguer `UserId` de `string` :

```typescript
// Brand : un champ fantôme qui n'existe qu'au niveau des types
type HostId = string & { readonly __brand: "HostId" };

function creerHostId(s: string): HostId {
  if (!s) throw new Error("id vide");
  return s as HostId; // seule porte d'entrée (validée)
}

function pingHost(id: HostId): void { /* ... */ }

pingHost(creerHostId("srv-01")); // OK
pingHost("srv-01");              // Erreur ! string ≠ HostId
```

**Variante objet** (plus lisible dans les messages d'erreur) :

```typescript
interface KiloWatt { readonly __unite: "kW"; valeur: number; }
```

> À utiliser pour : identifiants (évite de mélanger `userId` et `hostId`), unités (kW vs kVA — critique en énergie !), valeurs validées (email vérifié vs string brute).

---

## 43. async/await typé

```typescript
async function mesurerLatence(hote: string): Promise<number> {
  const debut = Date.now();
  await attendre(50); // simulation
  return Date.now() - debut; // number → Promise<number> automatique
}

function attendre(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

// À l'appel : await déballe
const latence: number = await mesurerLatence("srv-01");
```

**Règles d'or** :
1. Une fonction `async` retourne **toujours** une `Promise` (même `return 42` → `Promise<number>`).
2. `await` n'est autorisé que dans une fonction `async` (ou au top-level d'un module ESM — Node 14.8+).
3. **Ne jamais** passer une fonction `async` à `forEach`/`map` sans gérer les promesses (voir point 5).

```typescript
// ❌ Piège classique : forEach n'attend pas
hosts.forEach(async (h) => { await pingHost2(h); });
console.log("fini"); // s'affiche AVANT la fin des pings !

// ✅ Boucle séquentielle
for (const h of hosts) { await pingHost2(h); }

// ✅ Parallèle contrôlé
await Promise.all(hosts.map((h) => pingHost2(h)));
```

---

## 44. Promesses : typage avancé

```typescript
// Promise.all : tuple typé position par position
const [latence, charge] = await Promise.all([
  mesurerLatence("srv-01"),   // Promise<number>
  Promise.resolve("42%"),     // Promise<string>
]);
// latence: number, charge: string

// Promise.allSettled : n'échoue jamais globalement
const resultats = await Promise.allSettled(hosts.map((h) => pingHost2(h)));
for (const r of resultats) {
  if (r.status === "fulfilled") console.log(r.value);  // narrowing natif
  else console.error(r.reason);
}

// Premier arrivé gagne
const rapide = await Promise.race([
  fetchDonnees("http://a"),
  timeout(2000).then(() => { throw new Error("timeout"); }),
]);
```

**Typer un timeout/race proprement** :

```typescript
function avecTimeout<T>(p: Promise<T>, ms: number, message = "Timeout"): Promise<T> {
  let timer: NodeJS.Timeout;
  const garde = new Promise<never>((_, reject) => {
    timer = setTimeout(() => reject(new Error(`${message} après ${ms}ms`)), ms);
  });
  return Promise.race([p, garde]).finally(() => clearTimeout(timer));
}

const data3 = await avecTimeout(fetch("http://sonde/api").then((r) => r.json()), 3000);
```

> `Promise<never>` = une promesse qui ne se résout jamais positivement : parfaite pour les gardes de timeout.

---

## 45. Modules ESM : import / export

TypeScript utilise les modules ES standard. **Chaque fichier est un module** dès qu'il a un `import`/`export`.

```typescript
// src/types.ts
export interface Alerte { niveau: string; message: string; date: Date; }
export const SEUIL_CRITIQUE = 90;
export default class Moteur { /* ... */ } // un seul default par fichier

// src/alertes.ts
import Moteur, { Alerte, SEUIL_CRITIQUE } from "./types.js";
import type { Alerte as AlerteType } from "./types.js"; // import de type seul
```

**Points critiques avec `module: NodeNext`** :
1. **Extension obligatoire** : `from "./types.js"` (même si le fichier source est `types.ts`). C'est l'exigence Node ESM.
2. `import type` : import **effacé à la compilation** — zéro impact runtime, zéro risque d'import circulaire de valeurs. Activez `verbatimModuleSyntax: true` (TS 5.0+) pour forcer la distinction.
3. Évitez `export default` dans le code partagé : un seul par fichier, refactoring fragile, interop CJS compliquée. Préférez les exports nommés.

```typescript
// ✅ Recommandé
export function ping() {}
export interface Config {}

// Import namespace
import * as utils from "./utils.js";
utils.ping();
```

---

## 46. Résolution de modules, paths et interop

**`paths`** : des alias pour éviter les `../../../` :

```jsonc
// tsconfig.json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"],
      "@types/*": ["src/types/*"]
    }
  }
}
```

```typescript
import { Alerte } from "@/types.js"; // au lieu de ../../../types.js
```

> ⚠️ **Attention** : `paths` n'est qu'une indication pour `tsc`. Au **runtime**, Node ne connaît pas `@/`. Solutions : utiliser un bundler, ou le plugin `tsc-alias`, ou des imports relatifs en prod. Avec `tsx` en dev, ça marche via son resolver.

**Interop CommonJS** (`esModuleInterop: true`) :

```typescript
import express from "express";        // CJS vu comme default
import * as fs from "node:fs";        // namespace
import { readFile } from "node:fs/promises"; // nommé
```

**Ré-exports (barrel files)** :

```typescript
// src/index.ts — point d'entrée public du package
export * from "./types.js";
export * from "./alertes.js";
export { Moteur } from "./moteur.js";
```

> Les barrel files simplifient les imports mais peuvent ralentir la compilation et créer des cycles. Dans un gros projet, importez les modules directement.

---

## 47. Namespaces et declaration merging

Les `namespace` sont l'ancien mécanisme d'organisation interne (pré-modules). **En code moderne : préférez les modules ESM.** Ils restent utiles pour **étendre des types tiers** (declaration merging) :

```typescript
// Étendre les variables d'environnement typées (process.env est string | undefined par défaut)
declare global {
  namespace NodeJS {
    interface ProcessEnv {
      readonly SONDE_API_URL: string;
      readonly PORT?: string;
    }
  }
}
// À placer dans un fichier src/env.d.ts
// → process.env.SONDE_API_URL est maintenant string (pas string | undefined)
```

**Declaration merging utile** : ajouter une méthode à une interface tierce sans la modifier :

```typescript
declare module "express-serve-static-core" {
  interface Request {
    utilisateur?: { login: string };
  }
}
```

> Règle : `namespace` uniquement pour de l'**ambient** (déclarations `.d.ts`, extensions globales). Jamais pour structurer votre code applicatif.

---

## 48. Fichiers .d.ts et @types : les types sans le code

Un fichier `.d.ts` ne contient **que des déclarations de types** — aucun code exécutable. Usages :

1. **Typer du JS existant** sans le réécrire.
2. **Typer les variables globales** (`env.d.ts` ci-dessus).
3. **Publier les types** d'une bibliothèque (`declaration: true` génère les `.d.ts`).

