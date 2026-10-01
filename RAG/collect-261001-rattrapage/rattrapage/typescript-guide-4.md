---
id: collect-261001-rattrapage/rattrapage/typescript-guide-4
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [704, 921]
sha256: 9cdbe656efc9abc000febbb4f989178dfc2853dbcad6b55f493293d3e34c8557
---

# TypeScript — Guide ultra-complet

Quand une fonction accepte **des formes d'appel différentes**, les overloads décrivent chaque forme :

```typescript
// Signatures publiques (ce que l'appelant voit)
function requete(url: string): Promise<string>;
function requete(url: string, format: "json"): Promise<unknown>;
function requete(url: string, format: "text"): Promise<string>;
// Implémentation (invisible de l'extérieur)
function requete(url: string, format: "json" | "text" = "text"): Promise<unknown> {
  // ...
  return fetch(url).then((r) => (format === "json" ? r.json() : r.text()));
}

const t = await requete("http://x");          // Promise<string>
const j = await requete("http://x", "json");  // Promise<unknown>
```

**Règles** :
- L'implémentation doit être compatible avec **toutes** les signatures.
- L'implémentation n'est pas appelable directement avec sa propre signature large.
- Alternative moderne : une seule signature avec union + narrowing interne (souvent plus simple).

---

## 21. Fonctions fléchées et callbacks typés

```typescript
// Callback typé inline
const hosts = ["srv-01", "srv-02"];
hosts.forEach((h, index) => {
  console.log(index, h.toUpperCase()); // h: string, index: number (inférés)
});

// Type de callback réutilisable
type AlerteHandler = (alerte: { niveau: string; message: string }) => void;

function surveiller(handler: AlerteHandler): void {
  handler({ niveau: "critique", message: "UPS en surcharge" });
}
```

**`void` vs `undefined` en retour** : un callback typé `( ) => void` accepte une fonction qui retourne quelque chose (la valeur est ignorée). C'est voulu : `forEach` n'utilise pas le retour.

```typescript
const f: () => void = () => 42; // OK : 42 est ignoré
```

**Async comme callback** : attention, `forEach` n'attend pas les promesses (section 43).

---

## 22. unknown, any, never, void : le quatuor à maîtriser

| Type | Signification | Assignable à | Opérations permises |
|---|---|---|---|
| `any` | « je renonce au typage » | tout (contamine) | tout, sans vérification |
| `unknown` | « je ne sais pas encore » (sûr) | rien sans narrowing | narrowing obligatoire |
| `never` | « impossible » | tout | aucune valeur possible |
| `void` | « pas de retour utile » | `undefined` seul | rien |

```typescript
// any : la trappe (à bannir via eslint @typescript-eslint/no-explicit-any)
declare const sale: any;
sale.nimporteQuoi(); // aucune erreur... jusqu'au runtime

// unknown : le any sûr — typique des données externes
async function lireJson(chemin: string): Promise<unknown> {
  const txt = await import("node:fs/promises").then((fs) => fs.readFile(chemin, "utf8"));
  return JSON.parse(txt); // JSON.parse retourne any → on le remonte en unknown
}
const data: unknown = await lireJson("conf.json");
if (typeof data === "object" && data !== null && "port" in data) {
  console.log((data as { port: number }).port); // narrowing puis usage
}

// never : fonction qui ne revient jamais
function panique(msg: string): never {
  throw new Error(msg);
}
// never : exhaustiveness checking (section 26)

// void : fonction sans retour significatif
function notifier(msg: string): void {
  console.log(msg);
}
```

**Règle d'équipe** : `unknown` pour les entrées externes (JSON, argv, API), `never` pour l'impossible, `any` interdit sauf interop legacy documentée.

---

## 23. Assertions de type et `as const`

L'assertion dit au compilateur : « fais-moi confiance, c'est ce type ». Elle ne convertit rien au runtime.

```typescript
// Depuis unknown / any
const data: unknown = JSON.parse('{"port": 8080}');
const port = (data as { port: number }).port;

// Syntaxe alternative (évitez-la en .tsx : conflit avec le JSX)
const port2 = (<{ port: number }>data).port;
```

**`as const`** : fige en littéraux readonly (déjà vu section 12) :

```typescript
const SEUILS = { critique: 90, warning: 75 } as const;
const NIVEAUX = ["debug", "info", "error"] as const;
type Niveau = (typeof NIVEAUX)[number]; // "debug" | "info" | "error"
```

**`satisfies`** (TS 4.9+) : vérifie la conformité **sans élargir** le type — le meilleur des deux mondes :

```typescript
const config = {
  hote: "srv-01",
  port: 22,
  proto: "ssh",
} satisfies { hote: string; port: number; proto: string };
// config.proto garde le type littéral "ssh" (pas string) ET est vérifié.
```

> **Avertissement** : une assertion fausse = un mensonge au compilateur. `JSON.parse` + `as MonType` sans validation runtime (Zod, section 68) est la source n°1 de bugs « impossibles » en TypeScript.

---

## 24. Narrowing : réduire une union avec typeof, instanceof, in

Le **narrowing** = le compilateur suit vos tests et affine le type dans chaque branche.

```typescript
function formater(valeur: string | number): string {
  if (typeof valeur === "string") {
    return valeur.toUpperCase(); // ici: string
  }
  return valeur.toFixed(2);      // ici: number
}
```

**Les 4 gardes natives :**

| Garde | Usage | Exemple |
|---|---|---|
| `typeof x === "..."` | primitifs | `"string"`, `"number"`, `"boolean"`, `"undefined"`, `"bigint"`, `"function"`, `"symbol"` |
| `x instanceof C` | classes | `err instanceof Error`, `d instanceof Date` |
| `"clef" in obj` | propriété présente | `"latenceMs" in resultat` |
| `Array.isArray(x)` | tableaux | `Array.isArray(donnees)` |

```typescript
function traiter(entree: string[] | string): void {
  if (Array.isArray(entree)) {
    entree.forEach((e) => console.log(e)); // string[]
  } else {
    console.log(entree.split(","));        // string
  }
}

// Garde "in" sur union d'objets
type R = { ok: true; latence: number } | { ok: false; erreur: string };
function log2(r: R): void {
  if ("latence" in r) console.log(r.latence); // { ok: true, ... }
  else console.log(r.erreur);
}
```

**Vérité/faux (truthiness)** : `if (x)` narrow `"" | 0 | null | undefined | NaN` hors du type. Pratique mais attention : `0` et `""` valides sont exclus — préférez `x != null` ou `x !== undefined` quand 0/"" sont légitimes.

---

## 25. Type guards personnalisés (type predicates)

Quand la logique de test est réutilisable, extrayez-la dans un **prédicat** `x is T` :

```typescript
interface Onduleur { type: "onduleur"; kva: number; }
interface Serveur2 { type: "serveur"; cpu: number; }
type Equipement2 = Onduleur | Serveur2;

// Prédicat : le retour "e is Onduleur" informe le compilateur
function estOnduleur(e: Equipement2): e is Onduleur {
  return e.type === "onduleur";
}

function puissance(e: Equipement2): number {
  if (estOnduleur(e)) {
    return e.kva * 0.8; // e: Onduleur ici
  }
  return 0;
}
```

**Assertion functions** (TS 3.7+) : des guards qui lancent au lieu de retourner false :

```typescript
function assertEstDefini<T>(v: T | null | undefined): asserts v is T {
  if (v === null || v === undefined) throw new Error("Valeur manquante");
}

declare const p: number | undefined;
assertEstDefini(p);
p.toFixed(2); // OK : p est number après l'assertion
```

> Les prédicats sont la façon propre de valider des données d'API avant usage — alternative légère à Zod pour des cas simples.

---

## 26. Discriminated unions & exhaustiveness checking

Une **union discriminée** = union d'objets partageant un champ discriminant littéral (souvent `type` ou `kind`). C'est le pattern n°1 pour les machines à états, les événements, les résultats d'opérations.

```typescript
type Evenement =
  | { type: "ping"; hote: string }
  | { type: "alerte"; niveau: "warn" | "critique"; message: string }
  | { type: "fin"; code: number };

