---
id: collect-261001-rattrapage/rattrapage/typescript-guide-9
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [1946, 2220]
sha256: 36151153db8a6e3c7150a4f9549648d013ae94441c013d2e1710baa6b27f4f3a
---

# TypeScript — Guide ultra-complet

```typescript
try {
  await operationRisquee();
} catch (e) {
  // e: unknown → narrowing obligatoire
  if (e instanceof Error) {
    console.error(`Échec: ${e.message}`);
  } else {
    console.error(`Échec non-Error: ${String(e)}`);
  }
}
```

**Ne jamais laisser une promesse sans catch dans un CLI** :

```typescript
// main.ts — point d'entrée blindé
async function main(): Promise<void> {
  // ... logique
}

main().catch((e: unknown) => {
  console.error(e instanceof Error ? e.message : e);
  process.exit(1);
});
```

**Erreurs à ne pas avaler** : un `catch` vide ou un `catch { return null; }` masque les pannes. Logguez toujours, même brièvement. En cron, un script qui échoue **silencieusement** est pire qu'un script qui plante bruyamment.

---

## 55. Le pattern Result : gérer les erreurs sans exceptions

Alternative aux exceptions pour les erreurs **prévisibles** (validation, HTTP 404) : retourner un objet `Result` — le succès ou l'échec est explicite dans le type.

```typescript
type Result<T, E = Error> =
  | { ok: true; valeur: T }
  | { ok: false; erreur: E };

function diviser(a: number, b: number): Result<number, string> {
  if (b === 0) return { ok: false, erreur: "Division par zéro" };
  return { ok: true, valeur: a / b };
}

const r = diviser(10, 0);
if (r.ok) {
  console.log(r.valeur); // number
} else {
  console.error(r.erreur); // string — impossible d'oublier le cas d'échec
}
```

**Helpers** :

```typescript
function ok<T>(valeur: T): Result<T, never> { return { ok: true, valeur }; }
function ko<E>(erreur: E): Result<never, E> { return { ok: false, erreur }; }

// Convertir une promesse en Result (plus de try/catch au call-site)
async function tenter<T>(p: Promise<T>): Promise<Result<T>> {
  try { return ok(await p); }
  catch (e) { return ko(e instanceof Error ? e : new Error(String(e))); }
}

const res2 = await tenter(fetch("http://sonde/api"));
if (!res2.ok) { console.error("Sonde injoignable:", res2.erreur.message); process.exit(1); }
```

> Exceptions vs Result : exceptions pour l'**exceptionnel** (bugs, infra), Result pour l'**attendu** (validation, 404, timeout métier). Les deux cohabitent très bien.

---

## 56. Décorateurs : métaprogrammation déclarative

Les décorateurs (TS 5.0+, standard ECMAScript) ajoutent un comportement à une classe/méthode **déclarativement**. Cas d'usage : logging, mesure de temps, retry, validation.

```typescript
// tsconfig : "experimentalDecorators" n'est PLUS requis pour les décorateurs
// standard TS 5.0+ (il reste requis pour l'ancienne syntaxe expérimentale).

function Mesurer(
  _cible: unknown,
  nom: string,
  descripteur: PropertyDescriptor,
): PropertyDescriptor {
  const original = descripteur.value as (...args: unknown[]) => unknown;
  descripteur.value = function (this: unknown, ...args: unknown[]): unknown {
    const debut = Date.now();
    try {
      return original.apply(this, args);
    } finally {
      console.log(`${nom}: ${Date.now() - debut}ms`);
    }
  };
  return descripteur;
}

class Collecteur3 {
  @Mesurer
  collecter(hote: string): string {
    return `donnees de ${hote}`;
  }
}
```

**Décorateur de classe** (ex : enregistrer automatiquement un plugin) :

```typescript
const REGISTRE: Array<new () => unknown> = [];

function Plugin(nom: string) {
  return function (classe: new () => unknown): void {
    REGISTRE.push(classe);
    console.log(`Plugin enregistré: ${nom}`);
  };
}

@Plugin("snmp")
class CollecteurSnmpPlugin {}
```

> **Signal de version** : syntaxe standard = TS 5.0+. L'ancienne syntaxe (`experimentalDecorators: true`) reste supportée mais diverge du standard — ne mélangez pas les deux.

---

## 57. ts-node, tsx et l'exécution directe

Compiler puis exécuter (`tsc` + `node`) est lourd en développement. Trois exécuteurs directs :

| Outil | Moteur | Vérification des types | Usage |
|---|---|---|---|
| `ts-node` | tsc | Oui (lent, `--transpileOnly` pour aller vite) | `npx ts-node src/index.ts` |
| `tsx` | esbuild | **Non** (transpile seule) | `npx tsx src/index.ts` |
| `node --experimental-strip-types` | natif | Non | Node 22.6+ |

**Recommandation** : `tsx` en dev (instantané), `tsc --noEmit` en parallèle ou en CI pour la vérification :

```jsonc
// package.json
{
  "scripts": {
    "dev": "tsx watch src/index.ts",  // relance à chaque sauvegarde
    "build": "tsc",
    "start": "node dist/index.js",
    "check": "tsc --noEmit",
    "test": "vitest run"
  }
}
```

> ⚠️ `tsx` **ne vérifie pas les types** : une erreur de type passe silencieusement à l'exécution. Le `npm run check` en CI est donc obligatoire, pas optionnel.

---

## 58. ESLint + Prettier : setup complet

Prettier **formate**, ESLint **analyse**. Les deux ensemble = code homogène sans débats.

```bash
npm i -D eslint @eslint/js typescript-eslint prettier eslint-config-prettier
```

```javascript
// eslint.config.js (format "flat", ESLint 9+)
import js from "@eslint/js";
import tseslint from "typescript-eslint";
import prettier from "eslint-config-prettier";

export default tseslint.config(
  js.configs.recommended,
  ...tseslint.configs.recommendedTypeChecked, // règles avec vérification de types
  prettier, // désactive les règles de style en conflit avec Prettier
  {
    languageOptions: {
      parserOptions: {
        projectService: true,
        tsconfigRootDir: import.meta.dirname,
      },
    },
    rules: {
      "@typescript-eslint/no-explicit-any": "error",   // bannit any
      "@typescript-eslint/no-unused-vars": "error",
      "no-console": "off", // on est en CLI, console.* est légitime
    },
  },
);
```

```jsonc
// .prettierrc
{ "semi": true, "singleQuote": false, "printWidth": 100, "trailingComma": "all" }
```

```jsonc
// package.json
{ "scripts": { "lint": "eslint src tests", "format": "prettier --write src tests" } }
```

**Règles à activer en priorité** : `no-explicit-any`, `no-floating-promises` (promesse non awaitée = bug en puissance), `no-misused-promises`.

---

## 59. Tests unitaires avec Vitest

Vitest = runner moderne, compatible API Jest, support TS natif, rapide.

```bash
npm i -D vitest
```

```typescript
// src/seuils.ts
export function niveauAlerte(charge: number): "ok" | "warn" | "critique" {
  if (charge >= 90) return "critique";
  if (charge >= 75) return "warn";
  return "ok";
}

// tests/seuils.test.ts
import { describe, it, expect } from "vitest";
import { niveauAlerte } from "../src/seuils.js";

describe("niveauAlerte", () => {
  it("retourne ok sous 75", () => {
    expect(niveauAlerte(42)).toBe("ok");
  });

  it("gère les bornes exactes", () => {
    expect(niveauAlerte(75)).toBe("warn");
    expect(niveauAlerte(90)).toBe("critique");
  });

  it.each([
    [0, "ok"],
    [74.9, "ok"],
    [75, "warn"],
    [89.9, "warn"],
    [90, "critique"],
    [100, "critique"],
  ])("charge %i → %s", (charge, attendu) => {
    expect(niveauAlerte(charge)).toBe(attendu);
  });
});
```

```bash
npx vitest run          # une fois (CI)
npx vitest              # mode watch (dev)
npx vitest run --coverage # couverture (nécessite @vitest/coverage-v8)
```

**Matchers essentiels** : `toBe` (égalité stricte), `toEqual` (égalité profonde), `toThrow`, `toContain`, `resolves`/`rejects` pour les promesses.

---

## 60. Mocks, fixtures et tests async

```typescript
// src/sonde.ts
export async function chargeSonde(fetcher: (url: string) => Promise<{ charge: number }>, id: string) {
  const data = await fetcher(`http://sondes/${id}`);
  return data.charge;
}

// tests/sonde.test.ts — injection de dépendance : pas de mock framework requis
import { it, expect, vi } from "vitest";
import { chargeSonde } from "../src/sonde.js";

it("remonte la charge", async () => {
  const fauxFetcher = vi.fn(async (_url: string) => ({ charge: 66 }));
  await expect(chargeSonde(fauxFetcher, "s1")).resolves.toBe(66);
  expect(fauxFetcher).toHaveBeenCalledWith("http://sondes/s1");
});

