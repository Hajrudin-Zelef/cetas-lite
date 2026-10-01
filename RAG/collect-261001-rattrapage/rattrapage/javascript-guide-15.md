---
id: collect-261001-rattrapage/rattrapage/javascript-guide-15
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [3378, 3476]
sha256: 3daaa33b21622ee87bb4d27f4528aa3c1706dd0a9c8d2f79cfd75dbe6cccaf00
---

# JavaScript — Guide ultra-complet

### Exploitation
- [ ] Timeouts sur toutes les I/O réseau et `child_process`
- [ ] Arrêt propre : `SIGTERM`/`SIGINT` → fermer serveur/connexions
- [ ] Supervision : process manager (`pm2`, systemd) avec restart auto
- [ ] Logs avec timestamps ISO, niveau (info/warn/error), contexte (qui/quoi)
- [ ] Santé : endpoint `/sante` + alerte si down
- [ ] Utilisateur système dédié, jamais root ; principe du moindre privilège
- [ ] Sauvegarde des données persistées (fichiers JSON, etc.)

---

## 64. Annexe — tableau de compatibilité ES

| Version | Année | Apports majeurs | Node min |
|---|---|---|---|
| ES2015 (ES6) | 2015 | `let`/`const`, fléchées, classes, modules, promises, template literals, destructuration, spread | 6+ |
| ES2016 | 2016 | `**`, `Array.includes` | 7+ |
| ES2017 | 2017 | `async/await`, `Object.entries/values`, `padStart/padEnd` | 8+ |
| ES2018 | 2018 | rest/spread objets, groupes nommés regex, lookbehind | 10+ |
| ES2019 | 2019 | `flat`/`flatMap`, `Object.fromEntries`, `trimStart/End`, `?.` catch sans param | 12+ |
| ES2020 | 2020 | `?.`, `??`, `BigInt`, `Promise.allSettled`, `matchAll`, `globalThis` | 14+ |
| ES2021 | 2021 | `replaceAll`, `??=`/`&&=`/`\|\|=`, séparateurs `_`, `Promise.any` | 16+ |
| ES2022 | 2022 | top-level await, `.at()`, champs privés `#`, `Object.hasOwn`, `Error.cause` | 18+ |
| ES2023 | 2023 | `toSorted`/`toReversed`/`toSpliced`, `findLast`, shebang | 20+ |
| ES2024 | 2024 | `Object.groupBy`, `Promise.withResolvers` | 22+ |
| ES2025 | 2025 | `Set` union/intersection/difference, `RegExp.escape` | 22+ |

> Node 22 couvre confortablement ES2024. Pour les navigateurs : vérifier sur caniuse.com si parc ancien.

---

## 65. Annexe — recettes express sysadmin

```js
// 1. Lire un fichier ligne par ligne (stream)
import { createReadStream } from "fs";
import readline from "readline";
const rl = readline.createInterface({ input: createReadStream("/var/log/syslog"), crlfDelay: Infinity });
for await (const ligne of rl) { if (/error/i.test(ligne)) console.log(ligne); }

// 2. Ping rapide d'un hôte (sécurisé)
import { execFile } from "child_process";
import { promisify } from "util";
const ok = await promisify(execFile)("ping", ["-c", "1", "-W", "1", "10.0.0.1"]).then(() => true).catch(() => false);

// 3. Requête API JSON avec timeout (Node 18+)
const data = await fetch("https://api.example.com/status", { signal: AbortSignal.timeout(5000) })
  .then(r => { if (!r.ok) throw new Error(`HTTP ${r.status}`); return r.json(); });

// 4. Horodater un log
const ts = () => new Date().toISOString();
console.log(`[${ts()}] sauvegarde terminée`);

// 5. Charger une config JSON avec valeurs par défaut
import { readFile } from "fs/promises";
const cfg = { port: 3000, hote: "localhost", ...JSON.parse(await readFile("config.json", "utf8").catch(() => "{}")) };

// 6. Lister les fichiers récents d'un dossier
import { readdir, stat } from "fs/promises";
const fichiers = await readdir("/var/log");
const recents = [];
for (const f of fichiers) {
  const s = await stat(`/var/log/${f}`).catch(() => null);
  if (s?.isFile() && Date.now() - s.mtimeMs < 24 * 3600e3) recents.push(f);
}

// 7. Sleep / attendre
const sleep = (ms) => new Promise(r => setTimeout(r, ms));

// 8. Exécuter N tâches en parallèle borné (voir pool() section 24/55)

// 9. Générer un mot de passe/token sûr
import { randomBytes } from "crypto";
const secret = randomBytes(24).toString("base64url"); // Node 15.7+

// 10. Valider une IPv4
const estIpv4 = (s) => /^(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.){3}(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)$/.test(s);

// 11. Top N des occurrences (comptage)
const top = (arr, n = 10) => [...arr.reduce((m, x) => m.set(x, (m.get(x) ?? 0) + 1), new Map())]
  .sort((a, b) => b[1] - a[1]).slice(0, n);

// 12. Échapper pour le shell — NE PAS FAIRE soi-même : utiliser execFile (section 33)

// 13. Dédupliquer et trier des lignes
const uniques = [...new Set((await readFile("hosts.txt", "utf8")).split("\n").map(l => l.trim()).filter(Boolean))].sort();

// 14. Timeout global sur un script
const TIMEOUT_MS = 60000;
setTimeout(() => { console.error("timeout global"); process.exit(2); }, TIMEOUT_MS).unref();

// 15. Arguments CLI simples
const args = Object.fromEntries(process.argv.slice(2).map(a => a.replace(/^--/, "").split("=")));
// node script.js --hote=10.0.0.1 --port=22 → { hote: "10.0.0.1", port: "22" }
```

---

*Fin du guide — bon code, et que tes scripts ne plantent jamais un vendredi soir.*
