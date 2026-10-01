---
id: collect-261001-rattrapage/rattrapage/typescript-guide-12
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [2620, 2840]
sha256: 1b0896f7eea97c968f0514502c9cd76fe1062e8459fd251d316fd055a116feb5
---

# TypeScript — Guide ultra-complet

**1. tsconfig de build (lib)** :
```jsonc
{
  "compilerOptions": {
    "target": "ES2022", "module": "NodeNext", "moduleResolution": "NodeNext",
    "outDir": "dist", "rootDir": "src",
    "declaration": true,          // ← génère les .d.ts
    "declarationMap": true,       // ← map vers les sources
    "sourceMap": true,
    "strict": true, "noEmitOnError": true,
    "verbatimModuleSyntax": true  // ← force import type / import valeur
  },
  "include": ["src"]
}
```

**2. package.json** — les champs `types`/`exports` :
```jsonc
{
  "name": "@mon-org/outils-reseau",
  "version": "1.0.0",
  "type": "module",
  "main": "./dist/index.js",
  "types": "./dist/index.d.ts",
  "exports": {
    ".": { "types": "./dist/index.d.ts", "default": "./dist/index.js" }
  },
  "files": ["dist"],
  "scripts": {
    "build": "tsc",
    "prepublishOnly": "npm run build && npm run check && npm run test"
  }
}
```

**3. Checklist publication :**
- [ ] `declaration: true` → les consommateurs ont l'autocomplétion sans `@types`.
- [ ] `verbatimModuleSyntax` → pas d'imports de types fantômes au runtime.
- [ ] README avec exemples typés + matrice de compatibilité Node.
- [ ] `npm publish --dry-run` pour inspecter le contenu avant envoi.
- [ ] Versionnage semver strict : un changement de type public = **breaking** (majeur) ou au minimum mineur documenté.
- [ ] `attw --pack` (package `@arethetypeswrong/cli`) pour vérifier que les types se résolvent correctement côté consommateur.

> Alternative sans publication : un **monorepo** avec `paths`/`references` TypeScript (`tsc -b`) pour partager le code typé en interne sans npm.

---

## 70. Cas pratique 1 : CLI d'inventaire réseau

**Besoin** : lire un CSV d'équipements, pinger chacun (parallélisme borné), sortir un JSON + un résumé. Tout ce qu'on a vu, assemblé.

```typescript
// src/inventaire.ts
import { readFile, writeFile } from "node:fs/promises";
import { z } from "zod";
import { paralleleBorne } from "./concurrence.js"; // section 66
import { pingSys } from "./reseau.js";             // section 67 (execFile)

// --- 1. Modèle validé à la frontière (Zod) ---
const EquipementSchema = z.object({
  nom: z.string().min(1),
  ip: z.string().ip({ version: "v4" }),
});
type Equipement = z.infer<typeof EquipementSchema>;

// --- 2. Parsing CSV : string -> unknown -> validé ---
function parseCsv(contenu: string): Equipement[] {
  const lignes = contenu.split("\n").map((l) => l.trim()).filter(Boolean);
  return lignes.map((ligne, i) => {
    const [nom = "", ip = ""] = ligne.split(";");
    const res = EquipementSchema.safeParse({ nom: nom.trim(), ip: ip.trim() });
    if (!res.success) {
      throw new Error(`Ligne ${i + 1} invalide: ${res.error.issues[0]?.message}`);
    }
    return res.data;
  });
}

// --- 3. Résultat typé en union discriminée ---
type ResultatSonde =
  | { nom: string; ip: string; ok: true; latenceMs: number }
  | { nom: string; ip: string; ok: false; erreur: string };

async function sonder(e: Equipement): Promise<ResultatSonde> {
  try {
    const debut = Date.now();
    await pingSys(e.ip);
    return { nom: e.nom, ip: e.ip, ok: true, latenceMs: Date.now() - debut };
  } catch (err) {
    return {
      nom: e.nom, ip: e.ip, ok: false,
      erreur: err instanceof Error ? err.message : String(err),
    };
  }
}

// --- 4. Orchestration ---
export async function inventaire(entreeCsv: string, sortieJson: string): Promise<void> {
  const equipements = parseCsv(await readFile(entreeCsv, "utf8"));
  console.error(`Sondage de ${equipements.length} équipements…`);

  const resultats = await paralleleBorne(equipements, 20, sonder);

  await writeFile(sortieJson, JSON.stringify(resultats, null, 2) + "\n");

  // Résumé avec narrowing sur l'union discriminée
  const ok = resultats.filter((r) => r.ok);
  const ko = resultats.filter((r) => !r.ok);
  console.log(`OK: ${ok.length} | KO: ${ko.length}`);
  for (const r of ko) {
    if (!r.ok) console.log(`  ✗ ${r.nom} (${r.ip}): ${r.erreur}`);
  }
  if (ko.length > 0) process.exitCode = 1; // échec bruyant pour la supervision
}

// src/index.ts
import { inventaire } from "./inventaire.js";

const [csv = "equipements.csv", json = "resultats.json"] = process.argv.slice(2);
inventaire(csv, json).catch((e: unknown) => {
  console.error(e instanceof Error ? e.message : e);
  process.exit(1);
});
```

**Ce que ce cas pratique démontre** : validation Zod à la frontière, union discriminée + narrowing, parallélisme borné, `execFile` sécurisé, exit codes propres, séparation parsing / métier / orchestration.

---

## 71. Cas pratique 2 : client API de supervision avec cache

**Besoin** : interroger une API de supervision (métriques d'onduleurs), mettre en cache avec TTL, gérer timeout + retry. Réutilisable dans un dashboard.

```typescript
// src/client-supervision.ts
import { z } from "zod";
import { retry } from "./retry.js";           // section 53
import { avecTimeout } from "./timeout.js";   // section 44

// --- Schémas : le contrat de l'API, validé au runtime ---
const MesureSchema = z.object({
  horodatage: z.string().datetime(),
  chargePct: z.number().min(0).max(100),
  tensionV: z.number().positive(),
  autonomieMin: z.number().nonnegative(),
});
const ReponseSchema = z.object({ mesures: z.array(MesureSchema) });

type Mesure = z.infer<typeof MesureSchema>;

interface OptionsClient {
  baseUrl: string;
  token: string;
  timeoutMs?: number;
  ttlCacheMs?: number;
}

export class ClientSupervision {
  private cache = new Map<string, { expire: number; data: Mesure[] }>();
  private timeoutMs: number;
  private ttl: number;

  constructor(private opts: OptionsClient) {
    this.timeoutMs = opts.timeoutMs ?? 5000;
    this.ttl = opts.ttlCacheMs ?? 60_000;
  }

  private async get<T>(chemin: string, schema: z.ZodType<T>): Promise<T> {
    const url = `${this.opts.baseUrl}${chemin}`;
    const res = await retry(
      () =>
        avecTimeout(
          fetch(url, { headers: { Authorization: `Bearer ${this.opts.token}` } }),
          this.timeoutMs,
        ),
      3,
    );
    if (res.status === 401) throw new Error("Token invalide (401)");
    if (!res.ok) throw new Error(`HTTP ${res.status} sur ${chemin}`);
    return schema.parse(await res.json()); // throw ZodError si contrat rompu
  }

  async mesures(onduleurId: string): Promise<Mesure[]> {
    const enCache = this.cache.get(onduleurId);
    if (enCache && enCache.expire > Date.now()) return enCache.data;

    const data = await this.get(
      `/api/onduleurs/${encodeURIComponent(onduleurId)}/mesures`,
      ReponseSchema,
    );
    this.cache.set(onduleurId, { expire: Date.now() + this.ttl, data: data.mesures });
    return data.mesures;
  }

  async derniereCharge(onduleurId: string): Promise<number | null> {
    const mesures = await this.mesures(onduleurId);
    return mesures.at(-1)?.chargePct ?? null; // .at(-1) : ES2022
  }

  invalider(onduleurId: string): void {
    this.cache.delete(onduleurId);
  }
}
```

**Points à noter** :
- `encodeURIComponent` sur les segments d'URL (injection évitée).
- Le token ne transite que dans le header `Authorization`, jamais dans l'URL (les URLs finissent dans les logs).
- Cache en `Map` avec TTL : simple, suffisant pour un outil interne (pas besoin de Redis à cette échelle).
- `z.ZodType<T>` rend `get` générique sur n'importe quel schéma.

---

## 72. Cas pratique 3 : parseur de logs avec alertes

**Besoin** : analyser un fichier de log en streaming (gros volume), détecter des motifs, émettre des alertes typées, sortir un rapport.

```typescript
// src/parse-logs.ts
import { lignes } from "./streaming.js"; // section 50 (AsyncGenerator)

