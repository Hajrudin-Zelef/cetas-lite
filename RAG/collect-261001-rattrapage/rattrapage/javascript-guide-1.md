---
id: collect-261001-rattrapage/rattrapage/javascript-guide-1
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "attention"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [1, 196]
sha256: 26cf276ef273494f95c6acc3b6fac9feb98f4def9ff5fb5dedb89434f0634352
---

# JavaScript — Guide ultra-complet

> **Pour Zelef** — chef de service systèmes & énergies, sysadmin.
> Objectif : maîtriser JavaScript pour les scripts système, les outils web internes et Node.js.
> Ton : direct, dense, tutoriel + référence. Chaque section contient du code exécutable, des tableaux et des pense-bête.
> Version ciblée : **ES2024** (syntaxes récentes signalées quand elles dépendent de la version).

---

## Sommaire

1. [Installation & tooling](#1-installation--tooling)
2. [Syntaxe de base](#2-syntaxe-de-base)
3. [var / let / const](#3-var--let--const)
4. [Types primitifs](#4-types-primitifs)
5. [Opérateurs](#5-opérateurs)
6. [Structures de contrôle](#6-structures-de-contrôle)
7. [Fonctions — les fondamentaux](#7-fonctions--les-fondamentaux)
8. [Closures](#8-closures)
9. [Paramètres avancés des fonctions](#9-paramètres-avancés-des-fonctions)
10. [Tableaux](#10-tableaux)
11. [Objets](#11-objets)
12. [Destructuration, spread, rest](#12-destructuration-spread-rest)
13. [Chaînes de caractères](#13-chaînes-de-caractères)
14. [Nombres et Math](#14-nombres-et-math)
15. [Dates](#15-dates)
16. [this — les 4 règles](#16-this--les-4-règles)
17. [Prototypes](#17-prototypes)
18. [Classes](#18-classes)
19. [Héritage et polymorphisme](#19-héritage-et-polymorphisme)
20. [Gestion d'erreurs](#20-gestion-derreurs)
21. [Event loop : call stack, micro/macro-tâches](#21-event-loop--call-stack-micromacro-tâches)
22. [Callbacks et callback hell](#22-callbacks-et-callback-hell)
23. [Promises en profondeur](#23-promises-en-profondeur)
24. [async / await en profondeur](#24-async--await-en-profondeur)
25. [Fetch : GET, POST, erreurs](#25-fetch--get-post-erreurs)
26. [JSON](#26-json)
27. [Timers](#27-timers)
28. [Modules ESM vs CommonJS](#28-modules-esm-vs-commonjs)
29. [npm et package.json](#29-npm-et-packagejson)
30. [Node.js — process, env, CLI](#30-nodejs--process-env-cli)
31. [Node.js — fs et path](#31-nodejs--fs-et-path)
32. [Node.js — http](#32-nodejs--http)
33. [Node.js — child_process](#33-nodejs--child_process)
34. [Node.js — events et streams](#34-nodejs--events-et-streams)
35. [Manipulation du DOM](#35-manipulation-du-dom)
36. [Événements et délégation](#36-événements-et-délégation)
37. [Formulaires](#37-formulaires)
38. [Stockage navigateur : localStorage, cookies](#38-stockage-navigateur--localstorage-cookies)
39. [Web APIs utiles](#39-web-apis-utiles)
40. [RegExp](#40-regexp)
41. [Map, Set, WeakMap, WeakSet](#41-map-set-weakmap-weakset)
42. [Itérateurs et générateurs](#42-itérateurs-et-générateurs)
43. [Symbols et BigInt](#43-symbols-et-bigint)
44. [Proxy et Reflect](#44-proxy-et-reflect)
45. [Programmation fonctionnelle](#45-programmation-fonctionnelle)
46. [Tests unitaires](#46-tests-unitaires)
47. [Debug](#47-debug)
48. [Bonnes pratiques & style](#48-bonnes-pratiques--style)
49. [Performance](#49-performance)
50. [Sécurité](#50-sécurité)
51. [Erreurs classiques des débutants (15)](#51-erreurs-classiques-des-débutants-15)
52. [Cas pratique 1 — parser de logs Apache/nginx](#52-cas-pratique-1--parser-de-logs-apachenginx)
53. [Cas pratique 2 — client API REST avec retry](#53-cas-pratique-2--client-api-rest-avec-retry)
54. [Cas pratique 3 — mini-serveur HTTP maison](#54-cas-pratique-3--mini-serveur-http-maison)
55. [Cas pratique 4 — script d'inventaire réseau](#55-cas-pratique-4--script-dinventaire-réseau)
56. [Cas pratique 5 — dashboard web interne](#56-cas-pratique-5--dashboard-web-interne)
57. [Cas pratique 6 — automatisation de fichiers](#57-cas-pratique-6--automatisation-de-fichiers)
58. [Cas pratique 7 — planificateur avec file de tâches](#58-cas-pratique-7--planificateur-avec-file-de-tâches)
59. [Pense-bête de poche](#59-pense-bête-de-poche)
60. [Glossaire](#60-glossaire)
61. [Quiz — 10 questions + réponses](#61-quiz--10-questions--réponses)
62. [Pour aller plus loin](#62-pour-aller-plus-loin)
63. [Checklist de mise en production](#63-checklist-de-mise-en-production)
64. [Annexe — tableau de compatibilité ES](#64-annexe--tableau-de-compatibilité-es)
65. [Annexe — recettes express sysadmin](#65-annexe--recettes-express-sysadmin)

---

## 1. Installation & tooling

### 1.1 Installer Node.js (LTS recommandée)

| Méthode | Commande / action |
|---|---|
| Windows | Télécharger le MSI LTS sur nodejs.org, ou `winget install OpenJS.NodeJS.LTS` |
| Debian/Ubuntu | `curl -fsSL https://deb.nodesource.com/setup_22.x \| sudo -E bash - && sudo apt install -y nodejs` |
| Gestionnaire de versions (recommandé) | `nvm` (Linux/macOS) ou `fnm` (multi-OS, rapide) |

```bash
# Avec fnm (recommandé, rapide, multi-shell)
curl -fsSL https://fnm.vercel.app/install | bash
fnm install 22
fnm default 22

# Vérifications
node --version   # v22.x.x
npm --version    # 10.x.x
```

### 1.2 L'environnement de travail

- **Éditeur** : VS Code + extensions ESLint, Prettier, "Error Lens".
- **REPL** : `node` en terminal pour tester vite ; `.help`, `.exit`, `_` = dernier résultat.
- **Navigateur** : console DevTools (F12) pour le DOM et le web.
- **Deno / Bun** : alternatives à Node (Bun = très rapide, compatible npm).

### 1.3 Premier script

```js
// hello.js
console.log("Bonjour depuis Node.js !");
console.log(`Version Node : ${process.version}`);
```

```bash
node hello.js
```

### 1.4 Checklist tooling

- [ ] Node LTS installé + vérifié (`node -v`)
- [ ] VS Code + ESLint + Prettier
- [ ] `npm init -y` compris (section 29)
- [ ] Savoir lancer `node script.js` et `node --watch script.js` (rechargement auto, Node 18.11+)
- [ ] `"type": "module"` vs CommonJS compris (section 28)

---

## 2. Syntaxe de base

```js
// Commentaire sur une ligne
/* Commentaire
   sur plusieurs lignes */

// Point-virgule : optionnel (ASI) mais recommandé en scripts sérieux
let a = 1;
let b = 2

// Template literals (backticks) : interpolation
const nom = "Zelef";
console.log(`Bonjour ${nom}, il est ${new Date().getHours()}h`);

// console : l'outil n°1
console.log("info");        // standard
console.error("erreur");    // stderr
console.warn("attention");  // avertissement
console.table([{a:1},{a:2}]); // tableau lisible
console.time("t"); /* ... */ console.timeEnd("t"); // chrono
```

### Mots-clés réservés (à ne pas utiliser comme identifiants)

`break case catch class const continue debugger default delete do else export extends finally for function if import in instanceof new return super switch this throw try typeof var void while with yield` (+ `let static enum implements package protected interface private public`, `await` en module).

### Conventions de nommage

| Usage | Convention | Exemple |
|---|---|---|
| Variables, fonctions | camelCase | `nbUtilisateurs`, `calculerTotal()` |
| Constantes globales | SNAKE_UPPER | `MAX_RETRY` |
| Classes | PascalCase | `ServeurHttp` |
| Fichiers | kebab-case | `parse-logs.js` |

---

## 3. var / let / const

C'est **le** point à comprendre en premier : la portée (scope) et la remontée (hoisting).

```js
// var : portée FONCTION, hoisting, réassignable, redéclarable
function demoVar() {
  console.log(x); // undefined (hoisting : déclaration remontée, pas la valeur)
  var x = 5;
  var x = 10; // OK avec var (à éviter !)
}

// let : portée BLOC, pas d'accès avant déclaration (TDZ), réassignable
function demoLet() {
  // console.log(y); // ReferenceError : zone morte temporelle (TDZ)
  let y = 5;
  y = 6; // OK
  // let y = 7; // SyntaxError : déjà déclaré
}

// const : portée BLOC, non réassignable, MAIS objet mutable !
const PI = 3.14159;
// PI = 3; // TypeError
const cfg = { host: "localhost" };
cfg.host = "10.0.0.1"; // OK : on mute le contenu, pas la référence
```

### Tableau comparatif

