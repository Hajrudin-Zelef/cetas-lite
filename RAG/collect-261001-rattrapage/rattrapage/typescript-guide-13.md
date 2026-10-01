---
id: collect-261001-rattrapage/rattrapage/typescript-guide-13
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [2841, 3055]
sha256: 9b6ebc377e62db53484e010b4bd0570ffc3435eac662fce44202c52d2063f943
---

# TypeScript — Guide ultra-complet

// --- Événements : union discriminée (section 26) ---
type EvenementLog =
  | { type: "erreur"; heure: string; service: string; message: string }
  | { type: "timeout"; heure: string; service: string; dureeMs: number }
  | { type: "auth"; heure: string; utilisateur: string; succes: boolean };

// Format attendu : "2026-09-26T10:00:00Z [svc-ups] ERROR: surcharge détectée"
const REGEX = /^(\S+) \[([^\]]+)\] (ERROR|TIMEOUT \d+ms|AUTH (OK|KO) \S+): (.*)$/;

function parseLigne(ligne: string): EvenementLog | null {
  const m = REGEX.exec(ligne);
  if (!m) return null;
  const [, heure = "", service = "", code = "", message = ""] = m;
  if (code === "ERROR") return { type: "erreur", heure, service, message };
  if (code.startsWith("TIMEOUT")) {
    return { type: "timeout", heure, service, dureeMs: Number(code.split(" ")[1]?.replace("ms", "")) || 0 };
  }
  const ok = code.includes("OK");
  return { type: "auth", heure, utilisateur: message.trim(), succes: ok };
}

export interface Rapport {
  totalLignes: number;
  erreurs: number;
  timeouts: number;
  echecsAuth: number;
  topServicesEnErreur: Array<{ service: string; nb: number }>;
}

export async function analyser(chemin: string): Promise<Rapport> {
  let totalLignes = 0, erreurs = 0, timeouts = 0, echecsAuth = 0;
  const parService = new Map<string, number>();

  for await (const ligne of lignes(chemin)) {
    totalLignes++;
    const evt = parseLigne(ligne);
    if (!evt) continue;
    switch (evt.type) {
      case "erreur":
        erreurs++;
        parService.set(evt.service, (parService.get(evt.service) ?? 0) + 1);
        break;
      case "timeout":
        timeouts++;
        break;
      case "auth":
        if (!evt.succes) echecsAuth++;
        break;
      default: {
        const impossible: never = evt; // exhaustivité (section 26)
        throw new Error(`Non géré: ${JSON.stringify(impossible)}`);
      }
    }
  }

  const topServicesEnErreur = [...parService.entries()]
    .map(([service, nb]) => ({ service, nb }))
    .sort((a, b) => b.nb - a.nb)
    .slice(0, 5);

  return { totalLignes, erreurs, timeouts, echecsAuth, topServicesEnErreur };
}
```

**Pourquoi ce design tient la route en prod** : streaming (mémoire constante même sur 10 Go de logs), regex ancrée (`^...$`, pas de catastrophic backtracking ici car pas de quantificateurs imbriqués), exhaustivité garantie par `never`, rapport sérialisable en JSON tel quel.

---

## 73. Cas pratique 4 : agrégateur de métriques pour dashboard

**Besoin** : agréger des mesures de plusieurs onduleurs (min/max/moyenne par tranche de 5 min) pour alimenter un dashboard. Pur, testable, sans I/O.

```typescript
// src/agregation.ts — 100% pur : testable sans mock
export interface MesureBrute {
  onduleur: string;
  horodatage: string; // ISO
  chargePct: number;
}

export interface Tranche {
  debut: string; // ISO, arrondi à 5 min
  nbMesures: number;
  min: number;
  max: number;
  moyenne: number;
}

export function arrondir5min(iso: string): string {
  const d = new Date(iso);
  d.setSeconds(0, 0);
  d.setMinutes(d.getMinutes() - (d.getMinutes() % 5));
  return d.toISOString();
}

export function agreger(mesures: MesureBrute[]): Tranche[] {
  const groupes = new Map<string, number[]>();
  for (const m of mesures) {
    const cle = arrondir5min(m.horodatage);
    const g = groupes.get(cle);
    if (g) g.push(m.chargePct);
    else groupes.set(cle, [m.chargePct]);
  }
  return [...groupes.entries()]
    .map(([debut, vals]) => ({
      debut,
      nbMesures: vals.length,
      min: Math.min(...vals),
      max: Math.max(...vals),
      moyenne: Math.round((vals.reduce((s, v) => s + v, 0) / vals.length) * 10) / 10,
    }))
    .sort((a, b) => a.debut.localeCompare(b.debut));
}

// tests/agregation.test.ts
import { it, expect } from "vitest";
import { agreger } from "../src/agregation.js";

it("agrège par tranche de 5 minutes", () => {
  const tranches = agreger([
    { onduleur: "ups-01", horodatage: "2026-09-26T10:01:00Z", chargePct: 40 },
    { onduleur: "ups-01", horodatage: "2026-09-26T10:03:00Z", chargePct: 60 },
    { onduleur: "ups-01", horodatage: "2026-09-26T10:07:00Z", chargePct: 80 },
  ]);
  expect(tranches).toHaveLength(2);
  expect(tranches[0]).toMatchObject({ min: 40, max: 60, moyenne: 50, nbMesures: 2 });
  expect(tranches[1]).toMatchObject({ min: 80, max: 80, moyenne: 80, nbMesures: 1 });
});
```

**Leçon** : séparer le **calcul pur** (testé en millisecondes) de la **collecte I/O** (ClientSupervision, section 71). Le dashboard (React, section 62) consomme les `Tranche[]` via une API qui expose ce résultat.

---

## 74. Migrer un projet JavaScript vers TypeScript

**Stratégie incrémentale** (ne jamais tout réécrire d'un coup) :

**Étape 1 — cohabitation** : autorisez le JS dans le projet TS.
```jsonc
// tsconfig.json
{ "compilerOptions": { "allowJs": true, "checkJs": false, "strict": true } }
```
Renommez `index.js` → `index.ts` progressivement, en commençant par les feuilles (utilitaires sans dépendances).

**Étape 2 — JSDoc d'abord** : sur les fichiers encore en `.js`, ajoutez des annotations JSDoc — `tsc` les comprend avec `checkJs: true` :
```javascript
// @ts-check
/**
 * @param {string} hote
 * @returns {Promise<number>}
 */
async function ping(hote) { /* ... */ }
```

**Étape 3 — renommage par vagues** : `.js` → `.ts`, corrigez les erreurs fichier par fichier. Mettez `// @ts-nocheck` en haut des fichiers pas encore migrés pour garder un build vert.

**Étape 4 — durcissement** : une fois tout en `.ts`, retirez `allowJs`, activez `noUncheckedIndexedAccess`, puis les règles ESLint strictes.

**Checklist migration :**
- [ ] Un seul fichier migré à la fois, tests verts entre chaque vague
- [ ] Les frontières (I/O, API) typées en premier — c'est là que les bugs se cachent
- [ ] `any` temporaires marqués `// TODO(ts-migration): typer` pour ne pas les oublier
- [ ] Pas de changement fonctionnel pendant la migration (même comportement, juste des types)

---

## 75. Pense-bête de poche

### Commandes
```bash
npx tsc --noEmit          # vérifier les types (CI)
npx tsc                   # compiler vers dist/
npx tsx src/index.ts      # exécuter direct (dev, sans vérif)
npx tsx watch src/index.ts # dev avec relance auto
npx vitest run            # tests une fois
npx eslint src tests      # lint
```

### Snippets à copier
```typescript
// Main CLI blindé
async function main(): Promise<void> { /* ... */ }
main().catch((e: unknown) => {
  console.error(e instanceof Error ? e.message : e);
  process.exit(1);
});

// Env validée
import { z } from "zod";
const ENV = z.object({
  PORT: z.coerce.number().default(3000),
  API_URL: z.string().url(),
}).parse(process.env);

// fetch avec timeout + retry (sections 53-54)
const res = await retry(() => fetchAvecTimeout(url, 3000), 3);
if (!res.ok) throw new Error(`HTTP ${res.status}`);

// Dictionnaire exhaustif
const SEUILS: Record<"warn" | "critique", number> = { warn: 75, critique: 90 };

// Union discriminée + exhaustivité
switch (e.type) {
  case "a": /* ... */ break;
  default: { const x: never = e; throw new Error("non géré"); }
}

// Narrowing unknown (données externes)
if (typeof v === "object" && v !== null && "port" in v) { /* ... */ }

// Parallélisme borné
await paralleleBorne(items, 20, (x) => traiter(x));
```

