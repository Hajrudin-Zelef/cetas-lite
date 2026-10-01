---
id: collect-261001-rattrapage/rattrapage/typescript-guide-6
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [1191, 1446]
sha256: 0f77cc032f3c7b46e23c0b9d66d200c884272a6823c9648f0f559dfcde93c5c1
---

# TypeScript — Guide ultra-complet

```typescript
interface Cache<T = string> { // T vaut string si non précisé
  get(clef: string): T | undefined;
}
const c1: Cache = {} as Cache;        // Cache<string>
const c2: Cache<number> = {} as Cache<number>;
```

**Contrainte `keyof`** — le duo gagnant avec les objets :

```typescript
function lireProp<Objet, Clef extends keyof Objet>(obj: Objet, clef: Clef): Objet[Clef] {
  return obj[clef];
}
const srvX = { nom: "srv-01", port: 22 };
const p = lireProp(srvX, "port"); // number (pas string | number !)
lireProp(srvX, "inexistant");     // Erreur de compilation
```

---

## 35. keyof, typeof et l'accès aux types

Trois opérateurs pour **dériver des types de valeurs existantes** — la base du DRY typé.

```typescript
const CONFIG = {
  hote: "srv-01",
  port: 22,
  tls: true,
} as const;

type ClefConfig = keyof typeof CONFIG;      // "hote" | "port" | "tls"
type Port = (typeof CONFIG)["port"];        // 22 (littéral, grâce à as const)

function getConfig<K extends ClefConfig>(k: K): (typeof CONFIG)[K] {
  return CONFIG[k];
}
getConfig("port");    // 22
getConfig("inconnu"); // Erreur
```

**Tableau récapitulatif :**

| Opérateur | Opère sur | Produit |
|---|---|---|
| `typeof v` | une valeur | son type |
| `keyof T` | un type objet | union de ses clés |
| `T[K]` | un type objet + clé | type de la propriété |
| `typeof import("...")` | un module | son namespace de types |

> Pattern sysadmin typique : définir **une seule** source de vérité (objet `as const` ou interface) et dériver tout le reste (clés, unions, validateurs).

---

## 36. Types utilitaires : Partial, Required, Readonly

TypeScript fournit des **utility types** qui transforment des types existants. Les trois premiers rendent les propriétés optionnelles / requises / readonly **récursivement au premier niveau**.

```typescript
interface ServeurComplet {
  nom: string;
  ip: string;
  port: number;
  tags: string[];
}

// Tout optionnel : parfait pour les PATCH / mises à jour partielles
type ServeurPatch = Partial<ServeurComplet>;
// { nom?: string; ip?: string; port?: number; tags?: string[] }

function majServeur(nom: string, patch: ServeurPatch): void {
  // ... UPDATE ... SET ... avec seulement les champs fournis
}
majServeur("srv-01", { port: 2222 }); // OK

// Tout requis (inverse de Partial)
type ServeurRequis = Required<ServeurComplet>;

// Tout readonly (premier niveau)
type ServeurFige = Readonly<ServeurComplet>;
```

> `Partial` est **shallow** (premier niveau uniquement). Pour du deep-partial, il faut un mapped type récursif (section 39).

---

## 37. Types utilitaires : Pick, Omit, Record

```typescript
interface ServeurComplet {
  nom: string; ip: string; port: number; tags: string[]; secret: string;
}

// Sous-ensemble : Pick (inclusion)
type ServeurPublic = Pick<ServeurComplet, "nom" | "ip" | "port">;
// { nom: string; ip: string; port: number }

// Soustraction : Omit (exclusion) — idéal pour masquer des secrets
type ServeurSansSecret = Omit<ServeurComplet, "secret">;

// Dictionnaire : Record<Clés, Valeur>
type Inventaire2 = Record<string, ServeurPublic>;
const inv2: Inventaire2 = {
  "srv-01": { nom: "srv-01", ip: "10.0.0.1", port: 22 },
};

// Record avec union de clés : exhaustif !
type Niveau = "warn" | "critique";
const seuils: Record<Niveau, number> = { warn: 75, critique: 90 };
// Oublier "critique" → erreur. Le compilateur vous force à tout couvrir.
```

**Cas d'usage par utilitaire :**

| Utilitaire | Quand l'utiliser |
|---|---|
| `Partial<T>` | Body PATCH, options de fonction, formulaires partiels |
| `Required<T>` | Forcer la complétude après validation |
| `Readonly<T>` | Config figée, constantes partagées |
| `Pick<T, K>` | DTO public, projection de champs |
| `Omit<T, K>` | Retirer secrets / champs internes avant export JSON |
| `Record<K, V>` | Dictionnaires, tables de correspondance exhaustives |

---

## 38. Types utilitaires : ReturnType, Parameters, Awaited…

Ces utilitaires **extraient** des informations de fonctions existantes — fini la duplication de signatures.

```typescript
async function fetchSonde(id: string): Promise<{ charge: number }> {
  return { charge: 42 };
}

type SondeData = Awaited<ReturnType<typeof fetchSonde>>;
// { charge: number } — le type "déballe" la promesse

type ArgsFetch = Parameters<typeof fetchSonde>;
// [id: string]

function creer(...args: ArgsFetch): void { /* réutilise la signature */ }

// Autres utiles :
type Ctor = ConstructorParameters<typeof Date>; // params du constructeur
type Inst = InstanceType<typeof Date>;           // type d'instance
```

**Tableau complet des utilitaires natifs :**

| Utilitaire | Effet |
|---|---|
| `Partial<T>` / `Required<T>` / `Readonly<T>` | Modifie l'optionnalité / readonly |
| `Pick<T,K>` / `Omit<T,K>` | Sélection / exclusion de clés |
| `Record<K,V>` | Dictionnaire |
| `Exclude<U,E>` / `Extract<U,E>` | Soustrait / garde des membres d'une union |
| `NonNullable<T>` | Retire `null`/`undefined` |
| `ReturnType<F>` / `Parameters<F>` | Retour / paramètres d'une fonction |
| `Awaited<T>` | Déballe les `Promise` imbriquées |
| `ConstructorParameters` / `InstanceType` | Sur les constructeurs |

```typescript
type Statut2 = "ok" | "ko" | "inconnu";
type SansInconnu = Exclude<Statut2, "inconnu">; // "ok" | "ko"
```

---

## 39. Mapped types : transformer des types par programme

Un **mapped type** applique une transformation à chaque clé d'un type — c'est une boucle `for` au niveau des types.

```typescript
// Rendre optionnel "à la main" (= Partial)
type MonPartial<T> = {
  [K in keyof T]?: T[K];
};

// Rendre nullable
type Nullable<T> = {
  [K in keyof T]: T[K] | null;
};

// Modificateurs : -? retire l'optionnel, -readonly retire le readonly
type FigéEnMutable<T> = {
  -readonly [K in keyof T]: T[K];
};
```

**Exemple concret : tous les champs d'un formulaire avec leurs erreurs** :

```typescript
interface Formulaire { hote: string; port: number; }
type Erreurs<T> = { [K in keyof T]?: string };

const erreurs: Erreurs<Formulaire> = { port: "Le port doit être entre 1 et 65535" };
```

**`as` dans les mapped types** (TS 4.1+) : renommer/filtrer les clés :

```typescript
type Getters<T> = {
  [K in keyof T as `get${Capitalize<string & K>}`]: () => T[K];
};
type G = Getters<{ nom: string }>;
// { getNom: () => string }
```

---

## 40. Conditional types & infer : la logique dans les types

Un **conditional type** = un `if` au niveau des types : `T extends U ? A : B`.

```typescript
type EstString<T> = T extends string ? true : false;
type A = EstString<"hello">; // true
type B = EstString<42>;      // false

// Aplatir un tableau : T[] -> T, sinon T
type Aplatir<T> = T extends (infer E)[] ? E : T;
type C = Aplatir<string[]>; // string
type D = Aplatir<number>;   // number
```

**`infer`** capture un morceau du type dans la branche `true`. C'est ainsi que `ReturnType` et `Awaited` sont implémentés en interne :

```typescript
// Simplifié : voilà comment ReturnType fonctionne vraiment
type MonReturnType<F> = F extends (...args: never[]) => infer R ? R : never;
```

**Cas sysadmin** : extraire le type de données d'un collecteur générique :

```typescript
type DonneesDe<C> = C extends { collecter(): Promise<infer D> } ? D : never;

declare class CollecteurSnmp2 {
  collecter(): Promise<{ charge: number }>;
}
type D2 = DonneesDe<CollecteurSnmp2>; // { charge: number }
```

> Les conditionals sont **distributifs** sur les unions : `Aplatir<string[] | number>` = `string | number`. Pour désactiver la distributivité, enveloppez : `[T] extends [U]`.

---

## 41. Template literal types : des strings calculées

Les template literals existent aussi **au niveau des types** (TS 4.1+) :

```typescript
type EvenementSonde = `sonde:${string}:${"up" | "down"}`;
const e1: EvenementSonde = "sonde:dc1:up";    // OK
const e2: EvenementSonde = "sonde:dc1:panic"; // Erreur

