---
id: collect-261001-rattrapage/rattrapage/typescript-guide-14
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [3056, 3159]
sha256: 25c6cc2632607b51caf2d2196e3f2c8aedb4612fcfb532e59fa65e15abe6f084
---

# TypeScript — Guide ultra-complet

### Tableau des opérateurs de types
| Syntaxe | Nom | Exemple |
|---|---|---|
| `A \| B` | Union | `string \| number` |
| `A & B` | Intersection | `{a} & {b}` |
| `T[]` / `Array<T>` | Tableau | `string[]` |
| `[A, B]` | Tuple | `[string, number]` |
| `keyof T` | Clés | `keyof Config` |
| `T[K]` | Accès indexé | `Config["port"]` |
| `typeof v` | Type d'une valeur | `typeof CONFIG` |
| `T extends U ? A : B` | Conditionnel | `Awaited<T>` |
| `` `a${T}` `` | Template literal | `` `sonde:${string}` `` |
| `infer R` | Capture | `ReturnType<F>` |
| `Partial/Pick/Omit/Record` | Utilitaires | `Partial<Config>` |

---

## 76. Glossaire

| Terme | Définition |
|---|---|
| **Annotation** | Type écrit explicitement (`: string`) |
| **Inférence** | Type deviné par le compilateur |
| **Narrowing** | Affinage d'une union via un test (`typeof`, `in`, `instanceof`) |
| **Type guard** | Fonction prédicat `x is T` qui narrow |
| **Assertion** | `as T` : forcer un type (sans contrôle runtime) |
| **Union discriminée** | Union d'objets avec champ discriminant commun |
| **Générique** | Type paramétré (`Array<T>`, `Promise<T>`) |
| **Utility type** | Type utilitaire natif (`Partial`, `Pick`…) |
| **Mapped type** | Type construit en itérant sur des clés (`{ [K in keyof T]: … }`) |
| **Conditional type** | `T extends U ? A : B` |
| **Declaration merging** | Fusion de plusieurs déclarations `interface` du même nom |
| **Ambient** | Déclaration `.d.ts` : types sans implémentation |
| **Branded type** | `string & { __brand }` : distinguer des types structurellement identiques |
| **Exhaustiveness** | Vérifier via `never` que tous les cas d'une union sont traités |
| **Transpilation** | Conversion TS → JS (les types sont effacés) |
| **Erasure** | Les types n'existent pas au runtime |
| **Soundness** | Garantie qu'un typage est correct (TS est *unsound* par endroits : `as`, tableaux covariants — assumé pour la praticité) |
| **Variance** | Sens de compatibilité des sous-types (covariance/contravariance) |
| **Overload** | Plusieurs signatures pour une implémentation |
| **Barrel file** | `index.ts` qui ré-exporte des modules |
| **Triple-slash** | `/// <reference path="..." />` : directive d'inclusion (legacy, préférez les imports) |

---

## 77. Quiz : 10 questions + réponses

**Q1.** Que produit `tsc` à partir d'une `interface` dans le JavaScript généré ?
> **R :** Rien. Les interfaces (comme tous les types) sont effacées à la compilation — zéro trace au runtime.

**Q2.** Quelle est la différence entre `any` et `unknown` ?
> **R :** `any` désactive toute vérification (tout est permis, contamination). `unknown` impose un narrowing avant usage : c'est le « any sûr », à utiliser pour les données externes.

**Q3.** Pourquoi `fetch` peut-il vous piéger avec un HTTP 500 ?
> **R :** `fetch` ne rejette que sur erreur réseau. Un statut 500 résout normalement la promesse : il faut tester `res.ok` (ou `res.status`) explicitement.

**Q4.** Dans `hosts.forEach(async (h) => { await ping(h); })`, les pings sont-ils attendus avant la suite du programme ?
> **R :** Non. `forEach` ignore les promesses retournées. Utilisez `for...of` + `await` (séquentiel) ou `Promise.all(hosts.map(...))` (parallèle).

**Q5.** À quoi sert le contrôle `const x: never = e` dans le `default` d'un `switch` sur une union discriminée ?
> **R :** À garantir l'exhaustivité : si un nouveau variant est ajouté à l'union sans `case` correspondant, `e` n'est plus `never` et la compilation échoue.

**Q6.** Que change `"noUncheckedIndexedAccess": true` ?
> **R :** L'accès `tableau[i]` devient `T | undefined` au lieu de `T`, forçant à gérer l'index hors limites — source fréquente de crashs en prod.

**Q7.** Pourquoi faut-il écrire `from "./utils.js"` (et non `"./utils"`) avec `moduleResolution: NodeNext` ?
> **R :** Node ESM exige l'extension complète à l'exécution. TypeScript aligne la résolution sur ce comportement : l'import doit refléter le fichier JS réellement chargé.

**Q8.** Quelle est la différence entre `??` et `||` pour une valeur par défaut ?
> **R :** `??` ne s'applique que si `null`/`undefined` ; `||` s'applique aussi à `0`, `""`, `false`, `NaN`. Pour un timeout configuré à `0`, `||` écraserait à tort la valeur.

**Q9.** `private` en TypeScript protège-t-il vraiment un champ au runtime ?
> **R :** Non, c'est une vérification compile-time uniquement — le champ reste accessible en JS. Pour une vraie encapsulation runtime, utilisez `#champ`.

**Q10.** Quel est l'intérêt de Zod par rapport à un simple `as MonType` après `JSON.parse` ?
> **R :** `as` est un mensonge potentiel sans contrôle runtime. Zod valide réellement les données **et** infère le type statique (`z.infer`) : un schéma = validation + typage, source unique de vérité.

---

## 78. Pour aller plus loin

**Documentation officielle (en anglais, la référence) :**
- *TypeScript Handbook* : typescriptlang.org/docs/handbook — le lire en entier une fois vaut tous les tutoriels.
- *TSConfig Reference* : chaque option expliquée avec exemples.
- *TypeScript Playground* : typescriptlang.org/play — tester un type en 10 secondes, partager via URL.

**Outils à explorer ensuite :**
- **tRPC** : API typées de bout en bout (le type du serveur devient le type du client — fini les contrats dupliqués).
- **Prisma / Drizzle** : ORM avec types générés depuis le schéma SQL.
- **Hono / Fastify** : frameworks HTTP rapides et bien typés pour vos API internes.
- **Effect** : programmation fonctionnelle avancée (gestion d'erreurs, retry, ressources) — quand `Result` ne suffit plus.
- **tsx + chokidar / pm2** : superviser vos outils TS en production.

**Pratique délibérée (4 exercices) :**
1. Réécrire un script bash/python d'inventaire en CLI TypeScript (section 70) avec tests.
2. Typer un client d'API existant avec Zod (section 71) et mesurer les bugs attrapés par `tsc --noEmit`.
3. Migrer un vieux script JS avec la méthode de la section 74.
4. Implémenter un mini-dashboard React+TS (section 62) qui consomme l'agrégateur de la section 73.

**Communautés :** le Discord TypeScript, r/typescript, et le dépôt DefinitelyTyped pour contribuer des `@types`.

---

*Fin du guide — bon code, et que vos `tsc --noEmit` restent verts.* ✅
