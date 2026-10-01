---
id: collect-261001-rattrapage/rattrapage/typescript-guide-10
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [2221, 2416]
sha256: 383079c5bbb09aa3557f2652dae84bd28cf01b1b0e870744ab0531e6b6d8b5e3
---

# TypeScript — Guide ultra-complet

it("propage les erreurs réseau", async () => {
  const enPanne = vi.fn(async (_url: string): Promise<{ charge: number }> => {
    throw new Error("connexion refusée");
  });
  await expect(chargeSonde(enPanne, "s1")).rejects.toThrow("connexion refusée");
});
```

**Bonnes pratiques de test** :
- Testez les **fonctions pures** en priorité (seuils, parsers, formateurs) : rapides, déterministes.
- **Injectez** les dépendances I/O (fetch, fs) en paramètres plutôt que de mocker des modules.
- Un test qui dépend du réseau ou de l'heure est un test **fragile** : isolez l'horloge (`vi.useFakeTimers()`).
- Visez les **cas limites** : bornes, vide, `null`, unicode, timeouts.

---

## 61. Debug : VS Code, source maps, inspector

**launch.json** — debugger le TypeScript source directement (grâce à `sourceMap: true`) :

```jsonc
// .vscode/launch.json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "node",
      "request": "launch",
      "name": "Debug src/index.ts",
      "runtimeExecutable": "npx",
      "runtimeArgs": ["tsx", "--inspect", "src/index.ts"],
      "args": ["--hote", "srv-01"],
      "console": "integratedTerminal"
    }
  ]
}
```

**Techniques de debug par situation :**

| Situation | Outil |
|---|---|
| Script qui plante en dev | Breakpoints VS Code + `tsx` |
| Plantage en prod (dist/) | `node --enable-source-maps dist/index.js` + stack trace mappée |
| Fuite mémoire / perf | `node --inspect` + Chrome DevTools (onglet Memory/Profiler) |
| Promesse qui ne se termine pas | `--trace-warnings`, `process.on("unhandledRejection")` |
| Comportement étrange des types | Survolez le type dans VS Code (`Ctrl+K, Ctrl+I`) |

**Garde-fou global** (à mettre dans `main.ts`) :

```typescript
process.on("unhandledRejection", (raison) => {
  console.error("Promesse non gérée:", raison);
  process.exit(1);
});
```

> En prod, les logs structurés (JSON) valent mieux que `console.log` : envisagez `pino` (rapide, typé).

---

## 62. React + TypeScript : l'essentiel

Même si votre cœur de métier est sysadmin, un **dashboard interne** en React+TS est un cas classique. Les 20% qui couvrent 80% des besoins :

```typescript
// Props typées
interface CarteSondeProps {
  nom: string;
  charge: number;
  onAlerte?: (nom: string) => void; // callback optionnel
}

function CarteSonde({ nom, charge, onAlerte }: CarteSondeProps) {
  return (
    <div onClick={() => onAlerte?.(nom)}>
      {nom} — {charge}%
    </div>
  );
}

// useState typé (inférence suffit souvent)
import { useState, useEffect } from "react";

function ListeSondes() {
  const [sondes, setSondes] = useState<string[]>([]); // annotation requise (init vide)
  const [erreur, setErreur] = useState<string | null>(null);

  useEffect(() => {
    let actif = true;
    fetch("/api/sondes")
      .then((r) => r.json())
      .then((d: unknown) => { if (actif) setSondes(d as string[]); })
      .catch((e: unknown) => { if (actif) setErreur(String(e)); });
    return () => { actif = false; }; // cleanup anti setState après unmount
  }, []);

  if (erreur) return <div>Erreur: {erreur}</div>;
  return <ul>{sondes.map((s) => <li key={s}>{s}</li>)}</ul>;
}
```

**Règles React+TS** :
- `useState<string[]>([])` : annotez quand l'initial est vide/ambigu.
- Événements : `(e: React.ChangeEvent<HTMLInputElement>) => ...`, `(e: React.FormEvent) => ...`.
- `useRef<HTMLInputElement>(null)` pour le DOM.
- **Ne mettez jamais `any`** dans les props : le composant devient un piège pour ses utilisateurs.

---

## 63. Les 15 erreurs du compilateur les plus fréquentes (expliquées + corrigées)

| # | Code | Message typique | Cause | Correction |
|---|---|---|---|---|
| 1 | TS2322 | `Type 'X' is not assignable to type 'Y'` | Affectation incompatible | Aligner les types ou narrower avant |
| 2 | TS2339 | `Property 'x' does not exist on type 'Y'` | Faute de frappe ou union non narrowée | Corriger le nom / ajouter un narrowing |
| 3 | TS2345 | `Argument of type 'X' is not assignable to parameter of type 'Y'` | Mauvais argument | Vérifier l'ordre et le type des args |
| 4 | TS18048 | `'x' is possibly 'undefined'` | `strictNullChecks` + accès non gardé | `if (x)`, `?.`, `??`, assertion |
| 5 | TS2531 | `Object is possibly 'null'` | Idem avec `null` | Narrowing explicite |
| 6 | TS7006 | `Parameter 'x' implicitly has an 'any' type` | `noImplicitAny` | Annoter le paramètre |
| 7 | TS2454 | `Variable 'x' is used before being assigned` | Assignation conditionnelle non prouvée | Initialiser ou restructurer |
| 8 | TS2564 | `Property 'x' has no initializer` | `strictPropertyInitialization` | Initialiser, `?`, ou `!` assumé |
| 9 | TS2571 | `Object is of type 'unknown'` | Usage direct d'un `unknown` | Narrowing (typeof/in/instanceof) |
| 10 | TS2741 | `Property 'x' is missing in type ... but required` | Objet littéral incomplet | Ajouter la propriété ou la rendre optionnelle |
| 11 | TS2769 | `No overload matches this call` | Aucune surcharge ne correspond | Vérifier les combinaisons d'args |
| 12 | TS2349 | `This expression is not callable` | Appeler quelque chose qui n'est pas une fonction (`string \| (() => void)` non narrowé) | Narrower avec `typeof x === "function"` |
| 13 | TS2304 | `Cannot find name 'x'` | Faute de frappe / import manquant / mauvaise portée | Importer ou corriger |
| 14 | TS1192 | `Module has no default export` | `import X from` sans default | Utiliser l'import nommé |
| 15 | TS2307 | `Cannot find module './x'` | Mauvais chemin / extension manquante (NodeNext exige `.js`) | Corriger le chemin + extension |

**Exemples corrigés :**

```typescript
// ❌ TS18048: 'port' is possibly 'undefined'
declare const cfg: { port?: number };
const p1 = cfg.port.toFixed(); // Erreur

// ✅ Corrections possibles
const p2 = cfg.port?.toFixed();          // string | undefined
const p3 = (cfg.port ?? 443).toFixed();  // string — valeur par défaut
if (cfg.port !== undefined) {
  const p4 = cfg.port.toFixed();          // string — narrowing
}

// ❌ TS2339 sur union
declare const r: { ok: true; v: number } | { ok: false; e: string };
console.log(r.v); // Erreur : 'v' n'existe pas sur le 2e membre
// ✅
if (r.ok) console.log(r.v); else console.error(r.e);

// ❌ TS7006
const dbl = [1, 2].map((x) => x * 2); // OK ici (x inféré), mais :
function traiter2(x) { return x; }     // TS7006 en strict
// ✅
function traiter2(x: unknown): unknown { return x; }
```

**Méthode de lecture d'une erreur TS** : 1) le code (TSxxxx) → 2) *ce que* le compilateur attendait → 3) *pourquoi* votre valeur ne convient pas → 4) corrigez la **cause** (le type), pas le symptôme (pas de `as any` pour faire taire).

---

## 64. Erreurs classiques des débutants (12, avec corrections)

### 1. Utiliser `any` comme rustine
```typescript
const data: any = JSON.parse(brut); // ❌ contamine tout le fichier
const data2: unknown = JSON.parse(brut); // ✅ puis narrowing / Zod
```

### 2. Oublier que `fetch` ne rejette pas sur HTTP 500
Toujours tester `res.ok` (section 52).

### 3. `forEach` + `async` sans await
`forEach` n'attend pas : utilisez `for...of` ou `Promise.all` (section 43).

### 4. Comparer avec `==` au lieu de `===`
```typescript
if (port == "22") {} // ❌ true par coercition — en TS on veut du strict
if (port === 22) {}  // ✅
```

### 5. Muter un paramètre par défaut partagé
```typescript
function f(tags: string[] = []): string[] { // ✅ OK : nouveau tableau à chaque appel
  tags.push("x"); return tags;
}
// ⚠️ En revanche, un objet mutable en variable module = état global caché.
```

### 6. Confondre `interface` et valeur au runtime
```typescript
interface S { nom: string; }
const x = new S(); // ❌ Impossible : une interface n'existe pas au runtime
// Les interfaces sont effacées. Pour instancier : class. Pour valider : Zod.
```

