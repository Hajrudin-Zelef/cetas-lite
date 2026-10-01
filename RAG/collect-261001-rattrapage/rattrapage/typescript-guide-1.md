---
id: collect-261001-rattrapage/rattrapage/typescript-guide-1
title: "TypeScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/typescript_guide.md
source_anchor: ""
source_lines: [1, 186]
sha256: d6b0b89aef8e686fe9a8cc7b027c5aea10c42772283160998429698f97828e82
---

# TypeScript — Guide ultra-complet

> **Public visé** : sysadmin / chef de service systèmes & énergies qui développe des outils internes, des CLI et des dashboards.
> **Ton** : direct, tutoriel + référence. Chaque section contient du code exécutable et des tableaux de synthèse.
> **Version couverte** : TypeScript 5.x (les syntaxes dépendant d'une version sont signalées). Node.js 20+.

---

## 1. Pourquoi TypeScript quand on vient du scripting

En administration système, on écrit beaucoup de scripts : inventaires, collecte SNMP, parsing de logs, glue entre API. JavaScript suffit pour un script jetable de 50 lignes. Dès que l'outil devient critique (un dashboard utilisé par l'équipe, un CLI de provisioning), les types deviennent un filet de sécurité :

- **Les erreurs sont détectées à la compilation**, pas à 3h du matin en astreinte.
- **L'autocomplétion** dans VS Code connaît la forme exacte de vos objets (une réponse d'API REST, une ligne de CSV parsée).
- **Le refactoring est sûr** : renommer un champ `hostname` en `fqdn` dans 40 fichiers devient une opération mécanique.
- **La documentation est vivante** : une interface `Onduleur` décrit le contrat mieux qu'un commentaire.

TypeScript = JavaScript + un système de types statique qui **disparaît à la compilation** (aucun surcoût au runtime).

```typescript
// Sans types : l'erreur n'apparaît qu'à l'exécution
function puissanceTotale(onduleurs) {
  return onduleurs.reduce((s, o) => s + o.puissance, 0);
  // si "puissance" s'appelle "kva" dans les données : NaN silencieux
}

// Avec types : l'erreur est signalée avant même d'exécuter
interface Onduleur { nom: string; kva: number; }

function puissanceTotale(onduleurs: Onduleur[]): number {
  return onduleurs.reduce((s, o) => s + o.kva, 0);
}
```

**Checklist mentale** — adoptez TypeScript quand :
- [ ] le script dépasse ~200 lignes ou est partagé avec l'équipe ;
- [ ] il manipule des données structurées (JSON d'API, CSV, YAML) ;
- [ ] il doit tourner en production / en cron sans surveillance ;
- [ ] plusieurs personnes vont le maintenir.

---

## 2. Installation et tooling

Prérequis : Node.js 20+ (`node -v`). Trois façons d'installer TypeScript :

```bash
# 1. Globale (pratique pour essayer, déconseillé en projet)
npm install -g typescript
tsc --version   # ex : Version 5.6.3

# 2. Par projet (recommandé : version figée par projet)
mkdir mon-outil && cd mon-outil
npm init -y
npm install -D typescript @types/node
npx tsc --version

# 3. Vérifier ce qui est installé
npx tsc --version
```

**Outils compagnons à installer d'emblée :**

| Outil | Rôle | Installation |
|---|---|---|
| `typescript` | Compilateur `tsc` | `npm i -D typescript` |
| `@types/node` | Types de l'API Node (fs, process…) | `npm i -D @types/node` |
| `tsx` | Exécution directe de TS (dev) | `npm i -D tsx` |
| `vitest` | Tests unitaires | `npm i -D vitest` |
| `eslint` + `prettier` | Qualité + formatage | voir section 58 |

**VS Code** (éditeur recommandé) : l'extension TypeScript est intégrée. Réglez la version du workspace pour utiliser celle du projet : `Ctrl+Shift+P` → *TypeScript: Select TypeScript Version* → *Use Workspace Version*.

**Structure de projet conseillée :**

```
mon-outil/
├── src/
│   ├── index.ts        # point d'entrée
│   ├── inventory.ts
│   └── types.ts        # interfaces partagées
├── tests/
│   └── inventory.test.ts
├── dist/               # sortie compilée (généré, à ignorer via .gitignore)
├── package.json
└── tsconfig.json
```

---

## 3. Premier programme : de `hello.ts` au JavaScript

```typescript
// src/hello.ts
const message: string = "Bonjour depuis TypeScript";
console.log(message);
```

Compilation et exécution :

```bash
npx tsc src/hello.ts --outDir dist --target es2022 --module nodenext
node dist/hello.js
```

Points clés :
- `tsc` **transpile** : il enlève les annotations de types et produit du JS pur.
- Par défaut, `tsc fichier.ts` compile **sans vérification stricte**. Activez toujours `--strict` (ou un `tsconfig.json`, section 4).
- Les types n'existent qu'à la compilation : `typeof maVariable` au runtime ne verra jamais une interface.

**Erreurs ≠ blocage** : par défaut, `tsc` émet quand même le JS même s'il y a des erreurs de types. Pour interdire ça en CI : `"noEmitOnError": true` dans le tsconfig.

```bash
# Vérifier sans émettre (idéal en CI / pre-commit)
npx tsc --noEmit
```

---

## 4. tsconfig.json : la référence complète

Le `tsconfig.json` est le centre de contrôle. Voici un modèle **production-ready** pour un outil Node, commenté option par option :

```jsonc
{
  "compilerOptions": {
    /* --- Sortie --- */
    "target": "ES2022",              // version JS générée (ES2022 = Node 18+ OK)
    "module": "NodeNext",            // système de modules (NodeNext = ESM/CJS natif)
    "moduleResolution": "NodeNext",  // comment résoudre les imports
    "outDir": "dist",                // dossier du JS compilé
    "rootDir": "src",                // dossier des sources
    "declaration": true,             // génère les .d.ts (utile si package publié)
    "sourceMap": true,              // .map pour debugger le TS d'origine
    "removeComments": false,

    /* --- Rigueur (voir section 5) --- */
    "strict": true,
    "noUncheckedIndexedAccess": true,
    "exactOptionalPropertyTypes": true,

    /* --- Modules --- */
    "esModuleInterop": true,         // import express from "express"
    "allowSyntheticDefaultImports": true,
    "resolveJsonModule": true,       // import config from "./config.json"
    "isolatedModules": true,         // compatibilité esbuild/tsx/swc

    /* --- Sécurité de compilation --- */
    "noEmitOnError": true,           // ne produit rien si erreur de types
    "forceConsistentCasingInFileNames": true,
    "skipLibCheck": true,            // ignore les erreurs dans les .d.ts tiers

    /* --- Style --- */
    "noUnusedLocals": true,          // erreur si variable locale inutilisée
    "noUnusedParameters": true,      // erreur si paramètre inutilisé
    "noFallthroughCasesInSwitch": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist", "tests"]
}
```

**Tableau des options les plus importantes :**

| Option | Effet | Recommandé |
|---|---|---|
| `target` | Version JS émise (`ES2022`, `ESNext`…) | `ES2022` (Node 20) |
| `module` / `moduleResolution` | Format des modules | `NodeNext`/`NodeNext` |
| `strict` | Active tout le bloc strict | `true` **toujours** |
| `outDir` / `rootDir` | Sépare sources et build | `dist` / `src` |
| `declaration` | Émet les fichiers `.d.ts` | `true` si lib publiée |
| `sourceMap` | Permet de debugger le `.ts` | `true` |
| `noEmitOnError` | Bloque la sortie en cas d'erreur | `true` en CI |
| `resolveJsonModule` | `import` de fichiers JSON | `true` |
| `esModuleInterop` | Compatibilité CommonJS | `true` |
| `skipLibCheck` | Ignore les erreurs des `.d.ts` tiers | `true` (build plus rapide) |
| `isolatedModules` | Chaque fichier compilable seul | `true` (requis par tsx/esbuild) |

> **Piège classique** : `"module": "NodeNext"` exige `"moduleResolution": "NodeNext"`. Si vous mettez l'un sans l'autre, `tsc` proteste.

---

## 5. Le mode strict, option par option

`"strict": true` est un **interrupteur général** qui active 7 vérifications. Les comprendre une par une évite de le désactiver par frustration.

