---
id: collect-261001-rattrapage/rattrapage/typescript-guide-11
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [2417, 2619]
sha256: e9976dd59a041eb2a0cbefc4236b6933b9c17f573a0166b577525be8ab3cf582
---

# TypeScript — Guide ultra-complet

### 7. Oublier l'extension `.js` dans les imports (module NodeNext)
```typescript
import { x } from "./utils";   // ❌ TS2307 avec moduleResolution NodeNext
import { x } from "./utils.js"; // ✅
```

### 8. Croire que `private` protège au runtime
`private` est compile-time. Donnée vraiment sensible → `#privé` ou ne pas l'exposer.

### 9. Retourner `undefined` implicitement dans tous les chemins
```typescript
function chercher(nom: string): Serveur { // ❌ TS: pas de retour dans tous les cas
  if (nom) { /* ... */ return srv; }
}
function chercher2(nom: string): Serveur | undefined { // ✅ type honnête
  if (nom) { /* ... */ return srv; }
  return undefined;
}
```

### 10. Négliger `noUncheckedIndexedAccess`
`tableau[i]` peut être `undefined` : en prod, c'est un crash `Cannot read properties of undefined`.

### 11. Mettre la logique métier dans le `catch`
Le `catch` gère l'échec, il ne continue pas le traitement nominal. Structure : try = nominal, catch = dégradé/log, finally = nettoyage.

### 12. Ne pas typer les variables d'environnement
`process.env.PORT` est `string | undefined` : validez au démarrage (section 47 + Zod) au lieu de découvrir le `undefined` en plein run.

---

## 65. Bonnes pratiques & style

**Checklist qualité avant chaque commit :**
- [ ] `npm run check` (tsc --noEmit) passe
- [ ] `npm run lint` passe, zéro `any` explicite
- [ ] `npm run test` passe
- [ ] Prettier appliqué (`npm run format`)
- [ ] Pas de `console.log` de debug oublié (ou logger structuré)

**Conventions de nommage :**

| Élément | Convention | Exemple |
|---|---|---|
| Variables, fonctions | camelCase | `nbTentatives`, `lireInventaire()` |
| Classes, interfaces, types, enums | PascalCase | `Serveur`, `NiveauAlerte` |
| Constantes globales | UPPER_SNAKE | `SEUIL_CRITIQUE` |
| Fichiers | kebab-case | `collecte-snmp.ts` |
| Booléens | préfixe `is/has/peut` | `isEnLigne`, `hasErreur` |

**Principes de design :**
1. **Types honnêtes** : si ça peut être `undefined`, écrivez `| undefined`. Le type est un contrat.
2. **Petites fonctions pures** : parse, formate, calcule → testables sans mock.
3. **I/O aux frontières** : `main()` orchestre ; la logique ne touche ni `fs` ni `fetch` directement (injection, section 60).
4. **Échec bruyant** : jamais de `catch` vide ; exit code ≠ 0 en cas d'échec (cron/supervision).
5. **Un seul `as`** par validation, jamais en chaîne : `as unknown as T` = aveu d'échec du typage.
6. **TSDoc** sur les fonctions publiques (`/** ... */` avec `@param`, `@returns`, `@throws`).

```typescript
/**
 * Calcule le niveau d'alerte à partir d'une charge en %.
 * @param charge - Charge mesurée (0-100)
 * @returns "ok" | "warn" | "critique"
 * @throws {RangeError} Si charge hors de [0, 100]
 */
export function niveauAlerte(charge: number): "ok" | "warn" | "critique" {
  if (charge < 0 || charge > 100) throw new RangeError(`Charge invalide: ${charge}`);
  if (charge >= 90) return "critique";
  if (charge >= 75) return "warn";
  return "ok";
}
```

---

## 66. Performance : compilation et runtime

**Compilation plus rapide :**
- `skipLibCheck: true` (ignore les `.d.ts` tiers) — gain majeur.
- `isolatedModules: true` + `tsx`/esbuild en dev (pas de vérification, transpile seule).
- Évitez les types **récursifs profonds** et les unions géantes (>10k membres) : le checker peut exploser.
- `tsc --noEmit` en CI ; en dev, laissez l'IDE vérifier à la volée.

**Runtime : les types coûtent zéro** (effacés). Les vrais leviers :
- **Éviter les `await` séquentiels** quand l'ordre n'importe pas : `Promise.all` (section 44).
- **Streamer** les gros fichiers au lieu de `readFile` (section 50).
- **Limiter la concurrence** sur les I/O massives (ne pas lancer 10 000 fetch en parallèle) :

```typescript
// Parallélisme borné : simple et robuste
export async function paralleleBorne<T, R>(
  items: T[],
  limite: number,
  fn: (item: T) => Promise<R>,
): Promise<R[]> {
  const resultats: R[] = new Array(items.length);
  let index = 0;
  async function worker(): Promise<void> {
    while (index < items.length) {
      const i = index++;
      resultats[i] = await fn(items[i] as T);
    }
  }
  await Promise.all(Array.from({ length: Math.min(limite, items.length) }, worker));
  return resultats;
}

// 500 hôtes, 20 pings simultanés max
const latences = await paralleleBorne(hosts, 20, (h) => mesurerLatence(h));
```

- `for...of` > `.forEach` + closure en boucle chaude (micro-opt, rarement décisif).
- JSON : `JSON.parse` sur des payloads énormes = synchrone et bloquant ; pour du très gros, streamer (ndjson).

---

## 67. Sécurité : valider à la frontière

TypeScript **ne protège que le code que vous écrivez**. Tout ce qui vient de l'extérieur (argv, env, JSON, API, stdin) est **hostile jusqu'à validation**.

**Checklist sécurité d'un outil interne :**
- [ ] Valider `process.env` au démarrage (présence + format), échouer vite si invalide.
- [ ] Valider tout JSON entrant (fichier, API) avec Zod ou des guards (section 68).
- [ ] Ne jamais interpoler une entrée utilisateur dans une commande shell → utilisez `execFile` (args séparés), jamais `exec` avec concaténation.
- [ ] Timeouts sur tout appel réseau (section 53) — un hang = un cron bloqué.
- [ ] Ne jamais logger de secrets (tokens, mots de passe) ; masquez-les (`***`).
- [ ] `npm audit` régulier ; `npm ci` (pas `npm install`) en CI/prod pour figer le lockfile.
- [ ] Fichiers sensibles (clés) : permissions `0o600`, jamais dans le repo.

```typescript
import { execFile } from "node:child_process";
import { promisify } from "node:util";
const execFileAsync = promisify(execFile);

// ✅ Arguments séparés : pas d'injection shell possible
async function pingSys(hote: string): Promise<string> {
  if (!/^[a-zA-Z0-9.-]+$/.test(hote)) throw new Error(`Hôte invalide: ${hote}`);
  const { stdout } = await execFileAsync("ping", ["-c", "1", "-W", "2", hote]);
  return stdout;
}

// ❌ JAMAIS : `exec(`ping -c 1 ${hote}`)` — injection si hote = "x; rm -rf /"
```

---

## 68. Zod : validation runtime + types inférés

Zod (ou Valibot/ArkType) résout le dilemme « `as` menteur vs validation manuelle verbeuse » : **un schéma = validation runtime + type statique inféré**.

```bash
npm i zod
```

```typescript
import { z } from "zod";

// 1. Schéma (runtime)
const ServeurSchema = z.object({
  nom: z.string().min(1),
  ip: z.string().ip({ version: "v4" }),
  port: z.number().int().min(1).max(65535).default(22),
  tags: z.array(z.string()).default([]),
});

// 2. Type inféré (statique) — source unique de vérité !
type ServeurValide = z.infer<typeof ServeurSchema>;
// { nom: string; ip: string; port: number; tags: string[] }

// 3. Validation à la frontière
function parseInventaire(brut: unknown): ServeurValide[] {
  return z.array(ServeurSchema).parse(brut); // throw ZodError si invalide
}

// Version non-throw (pattern Result, section 55)
const res4 = z.array(ServeurSchema).safeParse(brut);
if (!res4.success) {
  console.error(res4.error.issues.map((i) => `${i.path.join(".")}: ${i.message}`));
  process.exit(1);
}
const serveurs: ServeurValide[] = res4.data; // typé et VALIDÉ
```

**Schémas utiles en sysadmin :**

```typescript
const EnvSchema = z.object({
  PORT: z.coerce.number().int().min(1).max(65535).default(3000),
  SONDE_API_URL: z.string().url(),
  LOG_LEVEL: z.enum(["debug", "info", "warn", "error"]).default("info"),
});
export const ENV = EnvSchema.parse(process.env); // échoue vite au démarrage
// z.coerce.number() convertit "3000" -> 3000 : parfait pour les env vars
```

> **Règle** : Zod (ou équivalent) à **chaque frontière** : argv, env, fichiers JSON, réponses HTTP, webhooks. À l'intérieur, TypeScript suffit.

---

## 69. Publier un package typé sur npm

Pour qu'un package soit agréable à consommer en TS :

