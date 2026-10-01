---
id: collect-261001-rattrapage/rattrapage/javascript-guide-13
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [2973, 3215]
sha256: 42eb820f16cd66f5ca96da7fcb12181d28ef5c814b00d320613277dad877789d
---

# JavaScript — Guide ultra-complet

    return json(res, 404, { erreur: "route inconnue" });
  } catch (err) {
    console.error(err);
    return json(res, 500, { erreur: "erreur interne" }); // message générique côté client
  }
});

const PORT = Number(process.env.PORT) || 3000;
server.listen(PORT, () => console.log(`API sur http://localhost:${PORT}`));
process.on("SIGTERM", () => server.close(() => process.exit(0)));
process.on("SIGINT", () => server.close(() => process.exit(0)));
```

---

## 55. Cas pratique 4 — script d'inventaire réseau

Objectif : pinger une plage d'IP en parallèle borné, sortir un JSON d'inventaire.

```js
// inventaire.js — node inventaire.js 10.0.0 1 254
import { execFile } from "child_process";
import { promisify } from "util";
import { writeFile } from "fs/promises";
const execFileAsync = promisify(execFile);

const sleep = (ms) => new Promise(r => setTimeout(r, ms));

async function ping(ip, timeoutSec = 1) {
  try {
    await execFileAsync("ping", ["-c", "1", "-W", String(timeoutSec), ip], { timeout: 5000 });
    return true;
  } catch { return false; }
}

// Parallélisme borné : ne pas inonder le réseau
async function pool(taches, limite) {
  const resultats = [];
  const enCours = new Set();
  for (const t of taches) {
    const p = t().then(r => { enCours.delete(p); return r; });
    enCours.add(p); resultats.push(p);
    if (enCours.size >= limite) await Promise.race(enCours);
  }
  return Promise.all(resultats);
}

function* plageIp(base, debut, fin) {
  for (let i = debut; i <= fin; i++) yield `${base}.${i}`;
}

const [base = "10.0.0", debut = "1", fin = "254"] = process.argv.slice(2);
console.log(`Scan de ${base}.${debut} → ${base}.${fin} (50 parallèles max)...`);
const t0 = Date.now();

const joignables = await pool(
  [...plageIp(base, Number(debut), Number(fin))].map(ip => async () => {
    const ok = await ping(ip);
    if (ok) console.log(`  ✓ ${ip}`);
    return ok ? ip : null;
  }),
  50
).then(r => r.filter(Boolean));

const rapport = {
  genereLe: new Date().toISOString(),
  plage: `${base}.${debut}-${fin}`,
  dureeSec: Math.round((Date.now() - t0) / 1000),
  joignables,
  total: joignables.length,
};
await writeFile("inventaire.json", JSON.stringify(rapport, null, 2));
console.log(`\n${rapport.total} hôtes joignables en ${rapport.dureeSec}s → inventaire.json`);
```

---

## 56. Cas pratique 5 — dashboard web interne

Objectif : page HTML + JS qui interroge la mini-API (section 54) et affiche les équipements avec auto-refresh.

```html
<!DOCTYPE html>
<html lang="fr">
<head>
<meta charset="utf-8">
<title>Supervision interne</title>
<style>
  body { font-family: system-ui, sans-serif; margin: 2rem; }
  table { border-collapse: collapse; width: 100%; }
  th, td { border: 1px solid #ccc; padding: .5rem; text-align: left; }
  .ok { color: green; font-weight: bold; } .ko { color: red; font-weight: bold; }
  #erreur { color: red; }
</style>
</head>
<body>
<h1>Supervision</h1>
<p id="erreur" hidden></p>
<table>
  <thead><tr><th>Nom</th><th>IP</th><th>Statut</th></tr></thead>
  <tbody id="corps"></tbody>
</table>
<p>Dernière mise à jour : <span id="maj">—</span></p>

<script>
const API = "http://localhost:3000";
const corps = document.getElementById("corps");
const maj = document.getElementById("maj");
const erreur = document.getElementById("erreur");

async function rafraichir() {
  try {
    const res = await fetch(`${API}/equipements`, { signal: AbortSignal.timeout(5000) });
    if (!res.ok) throw new Error(`HTTP ${res.status}`);
    const eqs = await res.json();
    erreur.hidden = true;
    const frag = document.createDocumentFragment(); // un seul reflow
    for (const e of eqs) {
      const tr = document.createElement("tr");
      // textContent : pas de XSS même si l'API renvoie du louche
      const tdNom = document.createElement("td"); tdNom.textContent = e.nom;
      const tdIp = document.createElement("td"); tdIp.textContent = e.ip;
      const tdStatut = document.createElement("td");
      tdStatut.textContent = e.statut;
      tdStatut.className = e.statut === "ok" ? "ok" : "ko";
      tr.append(tdNom, tdIp, tdStatut);
      frag.appendChild(tr);
    }
    corps.replaceChildren(frag);
    maj.textContent = new Date().toLocaleString("fr-FR");
  } catch (e) {
    erreur.hidden = false;
    erreur.textContent = `Échec de rafraîchissement : ${e.message}`;
  }
}

rafraichir();
const timer = setInterval(rafraichir, 30000);
timer.unref?.(); // navigateur : unref n'existe pas, le ?. évite l'erreur
</script>
</body>
</html>
```

---

## 57. Cas pratique 6 — automatisation de fichiers

Objectif : renommer/organiser des exports (ex : `export_2026-09-26.csv` → archivage par mois), avec log.

```js
// organiser-exports.js — node organiser-exports.js /chemin/exports
import { readdir, mkdir, rename, stat } from "fs/promises";
import { join, basename, extname } from "path";

const racine = process.argv[2] ?? "./exports";
const journal = [];

for (const f of await readdir(racine)) {
  const chemin = join(racine, f);
  if (!(await stat(chemin)).isFile()) continue;
  // export_SW1_2026-09-26.csv → dossier 2026-09/
  const m = f.match(/^export_(.+?)_(\d{4}-\d{2})-\d{2}(\.\w+)$/);
  if (!m) { journal.push({ fichier: f, action: "ignoré (nom non reconnu)" }); continue; }
  const [, nom, mois, ext] = m;
  const dest = join(racine, mois, `${nom}${ext}`);
  await mkdir(join(racine, mois), { recursive: true });
  await rename(chemin, dest);
  journal.push({ fichier: f, action: `archivé → ${mois}/${basename(dest)}` });
}

console.table(journal);
await import("fs/promises").then(fs =>
  fs.appendFile(join(racine, "journal.txt"),
    `[${new Date().toISOString()}] ${journal.length} fichiers traités\n`));
```

---

## 58. Cas pratique 7 — planificateur avec file de tâches

Objectif : file FIFO persistée en JSON, exécution séquentielle, reprise après crash.

```js
// file-taches.js — node file-taches.js ajouter "sauvegarde SW1" / node file-taches.js traiter
import { readFile, writeFile } from "fs/promises";

const FICHIER = "./file.json";

async function charger() {
  try { return JSON.parse(await readFile(FICHIER, "utf8")); }
  catch { return { enAttente: [], terminees: [] }; }
}
const sauver = (etat) => writeFile(FICHIER, JSON.stringify(etat, null, 2));

const [, , commande, ...reste] = process.argv;

if (commande === "ajouter") {
  const etat = await charger();
  etat.enAttente.push({ id: Date.now(), tache: reste.join(" "), ajouteeLe: new Date().toISOString() });
  await sauver(etat);
  console.log("Tâche ajoutée.");
} else if (commande === "lister") {
  const etat = await charger();
  console.table(etat.enAttente);
} else if (commande === "traiter") {
  const etat = await charger();
  while (etat.enAttente.length > 0) {
    const tache = etat.enAttente.shift(); // FIFO
    console.log(`▶ ${tache.tache}`);
    try {
      await new Promise(r => setTimeout(r, 500)); // ← remplacer par le vrai traitement
      etat.terminees.push({ ...tache, termineeLe: new Date().toISOString(), statut: "ok" });
    } catch (e) {
      etat.terminees.push({ ...tache, statut: "echec", erreur: e.message });
    }
    await sauver(etat); // persisté après CHAQUE tâche → reprise après crash
  }
  console.log("File vide.");
} else {
  console.log("Usage : node file-taches.js <ajouter|lister|traiter> [tâche]");
}
```

---

## 59. Pense-bête de poche

### Syntaxe express

```js
const c = 1; let v = 2;                       // déclarations
const f = (a, b = 1) => a + b;                // fléchée + défaut
const { ip, port = 22 } = srv;                // destructuration
const copie = { ...srv, ip: "x" };            // spread objet
const [p, ...reste] = liste;                  // rest tableau
`ssh ${user}@${ip}`;                          // template literal
obj?.champ?.methode?.();                      // optional chaining
const x = val ?? defaut;                      // nullish coalescing
```

### Async express

