---
id: collect-261001-rattrapage/rattrapage/javascript-guide-8
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [1730, 1978]
sha256: 72b6c9593849742b1cbfc4e7235f84749c037791a3c52e3fffcfcb0a8e26b604
---

# JavaScript — Guide ultra-complet

### Gros fichiers : streams (ne jamais tout charger en RAM)

```js
import { createReadStream } from "fs";
import readline from "readline";

// Compter les lignes d'un log de 2 Go sans exploser la RAM
async function compterLignes(chemin) {
  let n = 0;
  const rl = readline.createInterface({ input: createReadStream(chemin), crlfDelay: Infinity });
  for await (const ligne of rl) { n++; } // for await...of sur un stream (section 42)
  return n;
}
```

### Sync vs Async

|  | Usage |
|---|---|
| `fs/promises` (`await readFile`) | **par défaut** — ne bloque pas l'event loop |
| `fs.readFileSync` | init au démarrage, scripts jetables uniquement |
| streams | gros fichiers, temps réel |

> **Jamais** de `*Sync` dans un serveur ou une boucle chaude.

---

## 32. Node.js — http

```js
import http from "http";

const server = http.createServer(async (req, res) => {
  const url = new URL(req.url, `http://${req.headers.host}`);
  console.log(`${req.method} ${url.pathname}`);

  if (url.pathname === "/sante" && req.method === "GET") {
    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify({ statut: "ok", uptime: process.uptime() }));
    return;
  }

  if (url.pathname === "/echo" && req.method === "POST") {
    let corps = "";
    for await (const chunk of req) corps += chunk; // lire le body en stream
    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify({ recu: JSON.parse(corps) }));
    return;
  }

  res.writeHead(404, { "Content-Type": "application/json" });
  res.end(JSON.stringify({ erreur: "introuvable" }));
});

const PORT = process.env.PORT || 3000;
server.listen(PORT, () => console.log(`Écoute sur http://localhost:${PORT}`));
```

### Client HTTP natif (alternative à fetch)

```js
// fetch suffit dans 95 % des cas (Node 18+). http.request pour le contrôle fin :
import https from "https";
function getJson(url) {
  return new Promise((resolve, reject) => {
    https.get(url, { timeout: 5000 }, (res) => {
      let data = "";
      res.on("data", c => data += c);
      res.on("end", () => res.statusCode === 200 ? resolve(JSON.parse(data)) : reject(new Error(`HTTP ${res.statusCode}`)));
    }).on("error", reject).on("timeout", function() { this.destroy(new Error("timeout")); });
  });
}
```

> En pratique : **fetch** pour les clients, **Express/Fastify** pour les vraies APIs (routage, middlewares).

---

## 33. Node.js — child_process

Exécuter des commandes système depuis Node : le pont sysadmin.

```js
import { exec, execFile, spawn } from "child_process";
import { promisify } from "util";
const execAsync = promisify(exec);

// 1. exec : commande shell complète (simple, mais RISQUE D'INJECTION)
const { stdout } = await execAsync("uptime");
console.log(stdout.trim());

// 2. execFile : binaire + args SANS shell — SÉCURISÉ, à préférer
const execFileAsync = promisify(execFile);
const { stdout: ip } = await execFileAsync("ip", ["-brief", "addr"]);

// 3. spawn : streaming temps réel (ping, tail -f, builds)
const ping = spawn("ping", ["-c", "4", "8.8.8.8"]);
ping.stdout.on("data", d => process.stdout.write(`[ping] ${d}`));
ping.stderr.on("data", d => process.stderr.write(`[ping:err] ${d}`));
const code = await new Promise(res => ping.on("close", res));
console.log("code de sortie :", code);
```

### Tableau : quelle fonction choisir

| Fonction | Shell ? | Streaming | Usage |
|---|---|---|---|
| `exec` | oui | non (buffer) | commande simple, sortie petite |
| `execFile` | non | non (buffer) | **sécurisé**, args dynamiques |
| `spawn` | non (opt `shell:true`) | oui | sortie longue / temps réel |
| `fork` | non | IPC | sous-processus Node (workers) |

### SÉCURITÉ : injection de commande

```js
// DANGEREUX si `hote` vient de l'utilisateur :
await execAsync(`ping -c 1 ${hote}`); // hote = "8.8.8.8; rm -rf /" → catastrophe

// SÛR : execFile sans shell, args séparés
await execFileAsync("ping", ["-c", "1", hote]); // ";" traité comme simple argument
```

### Timeout et kill

```js
const proc = spawn("sleep", ["60"], { timeout: 5000 }); // tué après 5s
proc.on("error", e => console.error(e.message));
// ou manuel : proc.kill("SIGTERM");
```

---

## 34. Node.js — events et streams

### EventEmitter : le cœur de Node

```js
import { EventEmitter } from "events";

class Superviseur extends EventEmitter {}
const sup = new Superviseur();

sup.on("panne", (equipement) => console.log(`ALERTE: ${equipement} en panne`));
sup.once("demarrage", () => console.log("premier démarrage")); // une seule fois
sup.emit("panne", "SW1");
sup.emit("panne", "SW1"); // les .on persistent, les .once non

sup.off("panne", handler);      // désabonner (référence de fonction requise)
sup.removeAllListeners("panne");
```

### Streams : traiter les données par morceaux

Quatre types : **Readable** (lecture), **Writable** (écriture), **Duplex**, **Transform**.

```js
import { createReadStream, createWriteStream } from "fs";
import { Transform, pipeline } from "stream/promises";
import { createGzip } from "zlib";

// Compresser un gros log SANS le charger en RAM
const filtreErreurs = new Transform({
  transform(chunk, enc, cb) {
    const lignes = chunk.toString().split("\n").filter(l => /error|crit/i.test(l));
    cb(null, lignes.join("\n") + "\n");
  }
});

await pipeline(
  createReadStream("/var/log/syslog"), // Readable
  filtreErreurs,                        // Transform
  createGzip(),                         // Transform
  createWriteStream("/tmp/erreurs.log.gz") // Writable
);
console.log("pipeline terminé");
// pipeline() gère les erreurs et ferme les streams — TOUJOURS préférer à .pipe()
```

---

## 35. Manipulation du DOM

Le DOM = représentation objet de la page HTML. `document` = point d'entrée.

```js
// Sélection
const el = document.getElementById("statut");      // par id (unique)
const btns = document.querySelectorAll(".btn");    // tous les .btn (NodeList)
const premier = document.querySelector("#menu a"); // premier match (sélecteur CSS)

// Contenu
el.textContent = "En ligne";        // texte brut (SÛR contre XSS)
el.innerHTML = "<b>En ligne</b>";   // HTML interprété (DANGEREUX si données externes !)
el.innerHTML = "";                  // vider

// Attributs / classes / styles
el.setAttribute("data-ip", "10.0.0.1");
el.getAttribute("data-ip");
el.classList.add("ok"); el.classList.remove("ko"); el.classList.toggle("actif");
el.style.color = "green"; // style inline (préférer les classes)

// Créer / insérer
const li = document.createElement("li");
li.textContent = "SW1 — 10.0.0.2";
document.getElementById("liste").appendChild(li);
// Moderne : append(), prepend(), before(), after(), remove()
li.remove();

// Parcourir
el.parentElement; el.children; el.closest(".carte"); // ancêtre le plus proche
```

### NodeList vs Array

```js
const items = document.querySelectorAll("li"); // NodeList (itérable, .forEach OK)
const tableau = [...items];                    // spread → vrai tableau (map/filter...)
```

### Performance DOM

- Regrouper les insertions via `DocumentFragment` (un seul reflow).
- Éviter `innerHTML +=` en boucle (re-parse tout à chaque fois).
- Mettre en cache les sélections hors des boucles.

```js
const frag = document.createDocumentFragment();
for (const srv of serveurs) {
  const li = document.createElement("li");
  li.textContent = `${srv.nom} — ${srv.ip}`;
  frag.appendChild(li);
}
document.getElementById("liste").appendChild(frag); // un seul reflow
```

---

## 36. Événements et délégation

```js
const btn = document.getElementById("btn-scan");

// addEventListener : la méthode moderne (plusieurs handlers, options)
btn.addEventListener("click", (e) => {
  console.log("cliqué !", e.target, e.clientX, e.clientY);
}, { once: true }); // once : exécuté une seule fois

btn.removeEventListener("click", handler); // nécessite la référence

