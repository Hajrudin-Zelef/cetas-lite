---
id: collect-261001-rattrapage/rattrapage/javascript-guide-7
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [1489, 1729]
sha256: 61234b0bee2cde6901c9104a1e7f9ac1952a549e5c31e23a6b5c90922b30a8d9
---

# JavaScript — Guide ultra-complet

Deux systèmes coexistent. Node historiquement = CommonJS ; le standard moderne = ESM.

### CommonJS (CJS) — `require` / `module.exports`

```js
// maths.cjs
function somme(a, b) { return a + b; }
const PI = 3.14159;
module.exports = { somme, PI };
// ou : exports.somme = somme;

// app.cjs
const { somme, PI } = require("./maths.cjs");
const path = require("path"); // modules natifs Node
```

Caractéristiques : **synchrone**, `require` appelable partout (même conditionnel), `__dirname`/`__filename` disponibles, extension `.cjs` (ou `.js` sans `"type": "module"`).

### ESM (ES2015+) — `import` / `export`

```js
// maths.mjs
export function somme(a, b) { return a + b; }
export const PI = 3.14159;
export default class Calculatrice { /* ... */ }

// app.mjs
import { somme, PI } from "./maths.mjs";       // import nommé (extension OBLIGATOIRE en Node)
import Calculatrice from "./maths.mjs";        // import défaut
import * as maths from "./maths.mjs";          // namespace
const { somme: s } = await import("./maths.mjs"); // import dynamique (conditionnel !)
```

Caractéristiques : **asynchrone**, imports statiques en tête (tree-shaking), strict mode implicite, pas de `__dirname` (utiliser `import.meta.url`), top-level await.

### Tableau comparatif

|  | CommonJS | ESM |
|---|---|---|
| Syntaxe | `require` / `module.exports` | `import` / `export` |
| Chargement | synchrone | asynchrone |
| Strict mode | non (sauf directive) | toujours |
| Tree-shaking | non | oui |
| `__dirname` | oui | non (`import.meta.url` + `fileURLToPath`) |
| Top-level await | non | oui |
| Recommandé pour | legacy, scripts rapides | **tout nouveau code** |

### Choisir dans Node

```json
// package.json
{ "type": "module" }   // .js = ESM (recommandé pour les nouveaux projets)
```

```js
// Équivalent __dirname en ESM
import { fileURLToPath } from "url";
import { dirname, join } from "path";
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
```

### Interopérabilité

- ESM peut importer du CJS : `import pkg from "vieux-package-cjs"`.
- CJS peut importer de l'ESM **uniquement** via `await import()` dynamique (Node 22 : `require(esm)` encore expérimental/limité).
- Règle pratique : publier/écrire en ESM ; si une dépendance n'existe qu'en CJS, l'importer en défaut.

---

## 29. npm et package.json

```bash
npm init -y                 # crée package.json avec défauts
npm install express         # dépendance de prod (+ package-lock.json)
npm install -D nodemon      # dépendance de dev
npm install -g pm2          # installation globale (CLI système)
npm uninstall lodash
npm update                  # met à jour selon les plages semver
npm outdated                # quoi mettre à jour
npm audit                   # vulnérabilités
npm audit fix
npx cowsay salut            # exécute un paquet SANS l'installer
```

### package.json — l'essentiel

```json
{
  "name": "outils-reseau",
  "version": "1.0.0",
  "type": "module",
  "description": "Scripts d'exploitation réseau",
  "main": "index.js",
  "bin": { "scan-reseau": "./bin/scan.js" },
  "scripts": {
    "start": "node src/index.js",
    "dev": "node --watch src/index.js",
    "test": "node --test",
    "lint": "eslint .",
    "inventaire": "node scripts/inventaire.js --sous-reseau 10.0.0.0/24"
  },
  "engines": { "node": ">=20" },
  "dependencies": { "commander": "^12.0.0" },
  "devDependencies": { "eslint": "^9.0.0" }
}
```

### Semver : `^` vs `~`

| Plage | Signification | `1.2.3` accepte |
|---|---|---|
| `^1.2.3` | compatible mineur | `1.x.x` (pas `2.0.0`) |
| `~1.2.3` | compatible patch | `1.2.x` (pas `1.3.0`) |
| `1.2.3` | exact | `1.2.3` uniquement |

### Scripts npm (le "Makefile" du JS)

```bash
npm run inventaire        # lance le script
npm start / npm test      # raccourcis (run implicite)
```

Astuce : chaîner avec `&&`, passer des args avec `--` : `npm run scan -- --rapide`.

### Bonnes pratiques npm

- Commiter `package.json` **et** `package-lock.json` (reproductibilité).
- Ne jamais commiter `node_modules/` (`.gitignore`).
- `npm ci` en CI/prod (installe exactement le lockfile, plus rapide et sûr que `install`).
- Fixer `engines.node` pour éviter les surprises de version.

---

## 30. Node.js — process, env, CLI

```js
// process : infos et contrôle du processus
process.version;          // "v22.x.x"
process.platform;         // "linux" | "win32" | "darwin"
process.arch;             // "x64", "arm64"
process.pid;              // PID
process.cwd();            // répertoire de travail
process.uptime();         // secondes depuis le démarrage
process.memoryUsage();    // { rss, heapTotal, heapUsed, ... }
process.exit(0);          // 0 = OK, 1+ = erreur (vide les stdout avant en sync)

// Variables d'environnement — JAMAIS de secrets en dur !
const token = process.env.API_TOKEN;
if (!token) { console.error("API_TOKEN manquant"); process.exit(1); }

// Arguments CLI : process.argv
// node scan.js --sous-reseau 10.0.0.0/24 --rapide
// argv = ["/usr/bin/node", "/chemin/scan.js", "--sous-reseau", "10.0.0.0/24", "--rapide"]
const args = process.argv.slice(2);

// stdin / stdout / stderr (streams)
process.stdout.write("progression...\r");
process.stdin.on("data", (d) => console.log("reçu :", d.toString().trim()));
```

### Parser d'arguments minimal (sans dépendance)

```js
function parseArgs(argv = process.argv.slice(2)) {
  const opts = { _: [] };
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a.startsWith("--")) {
      const [k, v] = a.slice(2).split("=");
      opts[k] = v ?? argv[++i] ?? true;
    } else if (a.startsWith("-") && a.length === 2) {
      opts[a[1]] = argv[++i] ?? true;
    } else opts._.push(a);
  }
  return opts;
}
// node scan.js --sous-reseau 10.0.0.0/24 -v scan
// → { _: ["scan"], "sous-reseau": "10.0.0.0/24", v: true }
```

> Pour un vrai CLI : paquet `commander` (robuste, help auto).

### Signaux et erreurs globales

```js
// Arrêt propre sur Ctrl+C
process.on("SIGINT", () => { console.log("\nArrêt..."); server.close(() => process.exit(0)); });
process.on("SIGTERM", () => process.exit(0)); // systemd/docker

// Filet de sécurité (logguer, PAS masquer les bugs)
process.on("uncaughtException", (err) => { console.error("FATAL:", err); process.exit(1); });
process.on("unhandledRejection", (raison) => { console.error("Promise non gérée:", raison); process.exit(1); });
```

### Variables d'environnement en pratique

```bash
# .env (jamais commité !) — chargé via paquet `dotenv` ou --env-file (Node 20.6+)
API_TOKEN=xxx
PORT=3000
```

```bash
node --env-file=.env app.js   # Node 20.6+ : natif, sans dépendance
```

---

## 31. Node.js — fs et path

```js
import { readFile, writeFile, appendFile, mkdir, readdir, stat, rm } from "fs/promises";
import { join, resolve, dirname, basename, extname } from "path";
import { existsSync, createReadStream } from "fs";

// Lire / écrire (promises — à privilégier)
const contenu = await readFile("/etc/hosts", "utf8");
await writeFile("/tmp/rapport.txt", contenu);
await appendFile("/var/log/app.log", `[${new Date().toISOString()}] démarrage\n`);
await mkdir("/tmp/a/b", { recursive: true });

// Lister et filtrer
const fichiers = await readdir("/var/log", { withFileTypes: true });
const logs = fichiers.filter(f => f.isFile() && f.name.endsWith(".log")).map(f => f.name);

// Infos fichier
const infos = await stat("/var/log/syslog");
infos.size; infos.mtime; infos.isFile(); infos.isDirectory();

// path : construire des chemins portables (jamais de concaténation à la main !)
join("/var", "log", "app.log");       // "/var/log/app.log"
resolve("relatif", "fichier.txt");    // absolu depuis cwd
basename("/var/log/app.log");         // "app.log"
extname("archive.tar.gz");            // ".gz"
dirname("/var/log/app.log");          // "/var/log"

// Supprimer (moderne)
await rm("/tmp/vieux", { recursive: true, force: true });
```

