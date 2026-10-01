---
id: collect-261001-rattrapage/rattrapage/typescript-guide-8
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [1693, 1945]
sha256: 6caf3f54e197366d4db4a5d51e088fd8d2b2413db194a30e090c3552ba8eb287
---

# TypeScript — Guide ultra-complet

```typescript
// src/legacy.d.ts — typer un vieux module JS sans types
declare module "vieux-client-snmp" {
  export function get(oid: string): Promise<string>;
  export const VERSION: string;
}
```

**@types/node** : le package qui type toute l'API Node (`fs`, `process`, `path`…). Sans lui, `import fs from "node:fs"` est `any` implicite → erreur en `strict`.

```bash
npm i -D @types/node   # types Node
npm i -D @types/express # types d'express (si pas de types natifs)
```

> De plus en plus de packages embarquent leurs propres types (`"types"` dans leur package.json) : vérifiez avant d'installer un `@types/*`.

---

## 49. Node.js : fs, path, os — les essentiels typés

```typescript
import { readFile, writeFile, mkdir, readdir } from "node:fs/promises";
import path from "node:path";
import os from "node:os";

// Chemins multi-plateformes
const dossier = path.join(os.homedir(), "inventaires");
const fichier = path.resolve(dossier, "srv-01.json");
path.extname(fichier); // ".json"
path.basename(fichier, ".json"); // "srv-01"

// Infos système
os.hostname();        // string
os.cpus().length;     // nombre de cœurs
os.freemem();         // octets libres
os.uptime();          // secondes
```

**Lire un dossier et filtrer** :

```typescript
const entrees = await readdir(dossier, { withFileTypes: true });
const jsons = entrees
  .filter((e) => e.isFile() && e.name.endsWith(".json"))
  .map((e) => path.join(dossier, e.name));
```

> `node:` prefix (ex : `node:fs/promises`) : explicite, évite les collisions avec des packages npm du même nom. Recommandé depuis Node 16+.

---

## 50. Lire et écrire des fichiers et du JSON

```typescript
import { readFile, writeFile } from "node:fs/promises";

interface InventaireServeur { nom: string; ip: string; }

// Lecture JSON typée (le chaînon manquant : JSON.parse → unknown, section 22)
async function lireInventaire(chemin: string): Promise<InventaireServeur[]> {
  const brut: unknown = JSON.parse(await readFile(chemin, "utf8"));
  if (!Array.isArray(brut)) throw new Error(`Format invalide: ${chemin}`);
  // Validation minimale champ par champ (ou Zod, section 68)
  return brut.map((e) => {
    if (typeof e !== "object" || e === null || !("nom" in e) || !("ip" in e)) {
      throw new Error(`Entrée invalide dans ${chemin}`);
    }
    const { nom, ip } = e as { nom: unknown; ip: unknown };
    if (typeof nom !== "string" || typeof ip !== "string") throw new Error("Types invalides");
    return { nom, ip };
  });
}

// Écriture
async function ecrireInventaire(chemin: string, data: InventaireServeur[]): Promise<void> {
  await writeFile(chemin, JSON.stringify(data, null, 2) + "\n", "utf8");
}

// Import JSON direct (resolveJsonModule: true) — pour la config statique
import config from "./config.json" with { type: "json" };
// Note : la syntaxe "with" (import attributes) = Node 20.10+ / TS 5.3+
```

**Streaming pour les gros fichiers** (logs multi-Go : ne jamais tout charger en RAM) :

```typescript
import { createReadStream } from "node:fs";
import { createInterface } from "node:readline";

async function* lignes(chemin: string): AsyncGenerator<string> {
  const rl = createInterface({ input: createReadStream(chemin), crlfDelay: Infinity });
  for await (const ligne of rl) yield ligne;
}

let erreurs = 0;
for await (const ligne of lignes("/var/log/app.log")) {
  if (ligne.includes("ERROR")) erreurs++;
}
console.log(`${erreurs} erreurs`);
```

---

## 51. CLI : process.argv, stdin/stdout, codes de sortie

```typescript
// args.ts — parsing minimal sans dépendance
const args = process.argv.slice(2); // ["--hote", "srv-01", "--port", "22"]

function lireOption(nom: string, defaut?: string): string | undefined {
  const i = args.indexOf(nom);
  if (i === -1) return defaut;
  return args[i + 1];
}

const hote = lireOption("--hote") ?? "localhost";
const port = Number(lireOption("--port", "22"));
```

**Conventions CLI à respecter** :

| Élément | Convention |
|---|---|
| Sortie normale | `console.log` → stdout |
| Erreurs / diagnostics | `console.error` → stderr |
| Succès | `process.exit(0)` (implicite) |
| Échec | `process.exit(1)` (ou 2 pour usage invalide) |
| `--help` / `-h` | Toujours implémenté |

```typescript
function aide(): never {
  console.log(`Usage: inventaire --hote <nom> [--port <n>]`);
  process.exit(2);
}
if (args.includes("--help") || args.includes("-h")) aide();

// Lecture stdin (pipe)
async function lireStdin(): Promise<string> {
  const morceaux: Buffer[] = [];
  for await (const chunk of process.stdin) morceaux.push(chunk as Buffer);
  return Buffer.concat(morceaux).toString("utf8");
}
// Usage : cat hosts.txt | mon-outil
```

> Pour un vrai CLI (sous-commandes, validation, help auto), utilisez **commander** ou **yargs** — typés nativement.

---

## 52. Réseau : fetch typé

Node 18+ a `fetch` global. Le typage de la **réponse** reste à votre charge (`response.json()` retourne `Promise<any>`) :

```typescript
interface SondageApi { charge: number; temperature: number; }

async function getSondage(baseUrl: string, id: string): Promise<SondageApi> {
  const res = await fetch(`${baseUrl}/api/sondes/${id}`, {
    headers: { "Accept": "application/json" },
  });
  if (!res.ok) {
    throw new Error(`HTTP ${res.status} ${res.statusText} sur ${id}`);
  }
  const data: unknown = await res.json();
  // Validation (ici manuelle, voir Zod section 68)
  if (typeof data !== "object" || data === null) throw new Error("Réponse invalide");
  return data as SondageApi;
}
```

**POST JSON typé** :

```typescript
async function postAlerte(url: string, alerte: { niveau: string; message: string }): Promise<void> {
  const res = await fetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(alerte),
  });
  if (!res.ok) throw new Error(`POST ${url}: HTTP ${res.status}`);
}
```

> `fetch` ne rejette **que** sur erreur réseau (DNS, connexion refusée). Un HTTP 500 **ne rejette pas** : toujours tester `res.ok`.

---

## 53. HTTP avancé : timeouts, retry, erreurs propres

```typescript
// Timeout via AbortController (le standard fetch)
async function fetchAvecTimeout(url: string, ms: number): Promise<Response> {
  const ctrl = new AbortController();
  const timer = setTimeout(() => ctrl.abort(new Error(`Timeout ${ms}ms`)), ms);
  try {
    return await fetch(url, { signal: ctrl.signal });
  } finally {
    clearTimeout(timer);
  }
}

// Retry avec backoff exponentiel — typé générique, réutilisable
export async function retry<T>(
  fn: () => Promise<T>,
  tentatives = 3,
  delaiMs = 500,
): Promise<T> {
  let derniere: unknown;
  for (let i = 0; i < tentatives; i++) {
    try {
      return await fn();
    } catch (e) {
      derniere = e;
      if (i < tentatives - 1) await new Promise((r) => setTimeout(r, delaiMs * 2 ** i));
    }
  }
  throw derniere;
}

// Usage : 3 tentatives, 500ms / 1s / 2s
const res = await retry(() => fetchAvecTimeout("http://sonde/api", 3000), 3);
```

**Hiérarchie d'erreurs métier** (mieux que des `Error` génériques) :

```typescript
class HttpError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = "HttpError";
  }
}
class TimeoutError extends Error {
  constructor(public ms: number) { super(`Timeout après ${ms}ms`); this.name = "TimeoutError"; }
}

// Discrimination au catch (section 54)
try {
  await getSondage("http://x", "s1");
} catch (e) {
  if (e instanceof TimeoutError) { /* replanifier */ }
  else if (e instanceof HttpError && e.status === 404) { /* sonde inconnue */ }
  else throw e;
}
```

---

## 54. Gestion d'erreurs : try/catch moderne

Depuis TS 4.4, en `strict` (`useUnknownInCatchVariables`), la variable de `catch` est **`unknown`** — pas `any`. C'est une feature, pas un bug :

