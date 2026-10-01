---
id: collect-261001-rattrapage/rattrapage/typescript-guide-3
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [428, 703]
sha256: 2e625029c50cd17d577ad79d9092982e77b4c6db10647777c412f3189273a2a3
---

# TypeScript — Guide ultra-complet

```typescript
const niveau = "info"; // type: "info" (littéral conservé)
let niveau2 = "info";  // type: string (élargi)
```

**`as const`** (assertion const, section 23) fige un objet entier en littéraux — très utile pour les tables de correspondance :

```typescript
const SEUILS = {
  critique: 90,
  warning: 75,
} as const;
// { readonly critique: 90; readonly warning: 75 }

type Seuil = keyof typeof SEUILS; // "critique" | "warning"
```

---

## 13. Enums : valeurs nommées (et leurs alternatives)

```typescript
enum StatutEquipement {
  EnLigne,      // 0
  HorsLigne,    // 1
  Maintenance,  // 2
}

enum Protocole {
  Snmp = "snmp",
  Ssh = "ssh",
  Http = "http",
}

const p: Protocole = Protocole.Snmp; // "snmp"
```

**Tableau comparatif : enum vs union de littéraux :**

| Critère | `enum` | Union de littéraux (`"a" \| "b"`) |
|---|---|---|
| Valeur au runtime | Objet réel généré | Disparaît (zéro coût) |
| Itération (`Object.values`) | Possible | Impossible |
| Reverse mapping (numérique) | Oui (`Statut[0] === "EnLigne"`) | Non |
| Interop JSON/API | Verbeux | Naturel (ce sont des strings) |
| Tree-shaking | Mauvais | Parfait |

**Recommandation moderne** : pour des API/JSON, préférez les unions de littéraux + `as const`. Gardez `enum` quand vous avez besoin d'itérer ou d'un mapping bidirectionnel.

```typescript
// Alternative recommandée : objet const
const PROTOCOLE = {
  Snmp: "snmp",
  Ssh: "ssh",
} as const;
type Protocole = (typeof PROTOCOLE)[keyof typeof PROTOCOLE]; // "snmp" | "ssh"
// Itérable via Object.values(PROTOCOLE), zéro surcoût, JSON-friendly.
```

> **Signal de version** : les `const enum` (zéro coût runtime) sont interdits avec `isolatedModules: true` (donc avec tsx/esbuild). Évitez-les.

---

## 14. Interfaces : les bases

Une interface décrit la **forme** d'un objet. C'est le type le plus courant dans le code TypeScript.

```typescript
interface Serveur {
  nom: string;
  ip: string;
  roles: string[];
  enLigne: boolean;
}

const srv: Serveur = {
  nom: "srv-01",
  ip: "10.0.0.1",
  roles: ["hyperviseur", "dns"],
  enLigne: true,
};
```

**Propriétés optionnelles** (`?`) et **readonly** :

```typescript
interface ConfigCollecte {
  readonly sondeId: string;  // assignable une seule fois
  intervalleSec?: number;     // optionnel
  destinations: string[];
}

const cfg: ConfigCollecte = { sondeId: "sonde-01", destinations: ["influx"] };
cfg.sondeId = "x"; // Erreur : readonly
```

**Interfaces de fonctions et d'index** :

```typescript
// Signature de fonction
interface PingFn {
  (hote: string, timeoutMs?: number): Promise<number>;
}

// Index signature : dictionnaire à clés dynamiques
interface Inventaire {
  [nom: string]: Serveur; // toute clé string -> Serveur
}
const inv: Inventaire = { "srv-01": srv };
```

> **Excès de propriétés** : un objet littéral avec une propriété inconnue est rejeté (frappe), mais pas un objet passé via variable. C'est une protection contre les fautes de frappe, pas une étanchéité.

```typescript
const s2: Serveur = { nom: "x", ip: "y", roles: [], enLigne: true, cpu: 4 };
// Erreur : 'cpu' n'existe pas dans Serveur (object literal excess check)
```

---

## 15. Type aliases : `type`

`type` crée un alias pour **n'importe quel** type : union, tuple, fonction, primitif, objet.

```typescript
type Ip = string;
type Handler = (event: string) => void;
type Matrice = number[][];
type Reponse<T> = { data: T; erreur?: string }; // générique (section 33)
```

**Objet via `type`** (équivalent à l'interface pour la forme simple) :

```typescript
type ServeurT = {
  nom: string;
  ip: string;
};
```

Différences techniques interface/type (détaillées section 16) : l'interface supporte la **déclaration fusionnée** et `extends`/`implements` ; le `type` supporte les unions, tuples et types calculés.

---

## 16. Interfaces vs Types : quand utiliser quoi

**Règle pratique (recommandation officielle TS) :**

| Situation | Privilégier |
|---|---|
| Forme d'un objet / d'une classe (contrat) | `interface` |
| Union, intersection, tuple, fonction, primitif | `type` |
| Type calculé (mapped, conditional) | `type` (obligatoire) |
| API publique d'une bibliothèque | `interface` (extensible par l'utilisateur) |
| Extension d'un type existant | `interface extends` |

```typescript
// ✅ interface : contrat extensible
interface Equipement { nom: string; }
interface Onduleur extends Equipement { kva: number; }

// ✅ type : composition impossible en interface
type Charge = number | "inconnue";
type Paire = [string, number];

// ✅ type : utilitaires (section 36-38)
type Partiel = Partial<Equipement>;
```

**Declaration merging** (interfaces uniquement) : deux déclarations du même nom **fusionnent**. Utile pour étendre des types tiers, dangereux si accidentel :

```typescript
interface Fenetre { titre: string; }
interface Fenetre { largeur: number; }
// Fenetre = { titre: string; largeur: number }
```

> En pratique d'équipe : **interfaces pour les modèles métier** (`Serveur`, `Onduleur`, `Alerte`), **types pour tout le reste**.

---

## 17. Optionnel, readonly, index signatures : le trio du modelage

```typescript
interface Sondage {
  readonly id: string;          // écriture unique
  hote: string;
  port?: number;                // peut être absent
  tags: readonly string[];      // tableau non modifiable
  meta: Record<string, string>; // dictionnaire (section 37)
  [clef: string]: unknown;       // attrape-tout (attention : affaiblit le type)
}
```

**Tableau de décision :**

| Besoin | Syntaxe |
|---|---|
| Champ facultatif | `port?: number` |
| Champ non réassignable | `readonly id: string` |
| Tableau non modifiable | `readonly string[]` |
| Dictionnaire homogène | `Record<string, T>` ou `{ [k: string]: T }` |
| Rendre tout optionnel | `Partial<T>` |
| Rendre tout requis | `Required<T>` |

**`?.` et `!` à l'usage** :

```typescript
declare const s: Sondage;
const p = s.port ?? 443;        // ✅ valeur par défaut propre
const p2 = s.port!.toFixed();   // ⚠️ non-null assertion : "je sais qu'il est défini"
```

> Le `!` (non-null assertion) désactive la vérification : à n'utiliser que quand vous avez une **garantie externe** (ex : validé juste avant). En `strict`, préférez le narrowing.

---

## 18. Fonctions : signatures et paramètres

```typescript
// Annotation complète
function ping(hote: string, timeoutMs: number): Promise<number> {
  // ...
  return Promise.resolve(12);
}

// Inférence du retour (possible, mais annotez les fonctions publiques)
function addition(a: number, b: number) {
  return a + b; // inféré: number
}
```

**Pourquoi annoter le retour des fonctions exportées** :
1. Le compilateur vérifie que tous les chemins retournent le bon type.
2. Le `.d.ts` généré est explicite.
3. Évite les fuites de types internes dans l'API publique.

**Paramètres typés et `this`** :

```typescript
// this typé explicitement (utile avec noImplicitThis)
function afficher(this: Serveur): void {
  console.log(this.nom);
}
```

---

## 19. Paramètres optionnels, par défaut et rest

```typescript
function connecter(
  hote: string,
  port: number = 22,        // défaut → type number, optionnel à l'appel
  options?: { timeout?: number }, // vraiment optionnel
  ...tags: string[]          // rest : 0..n strings
): void {
  console.log(hote, port, options?.timeout, tags);
}

connecter("srv-01");                    // OK
connecter("srv-01", 2222);              // OK
connecter("srv-01", 22, { timeout: 5 }); // OK
connecter("srv-01", 22, undefined, "prod", "critique"); // OK
```

**Ordre obligatoire** : requis → défaut/optionnel → rest. Un paramètre requis ne peut pas suivre un optionnel :

```typescript
function f(a?: string, b: string) {} // Erreur : b requis après a optionnel
```

---

## 20. Surcharge de fonctions (overloads)

