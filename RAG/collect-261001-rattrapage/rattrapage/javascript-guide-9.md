---
id: collect-261001-rattrapage/rattrapage/javascript-guide-9
title: "JavaScript — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-rattrapage/javascript_guide.md
source_anchor: ""
source_lines: [1979, 2228]
sha256: 7ee5bdd398d70dc0e268b84163c1846accb2003f053ee1da159d46583a810619
---

# JavaScript — Guide ultra-complet

// Empêcher le comportement par défaut / la propagation
form.addEventListener("submit", (e) => {
  e.preventDefault();   // ne pas recharger la page
  // ... validation + fetch ...
});
```

### Propagation : capture → cible → bubbling

```js
// Par défaut les handlers s'exécutent au BUBBLING (remontée vers document)
document.getElementById("externe").addEventListener("click", () => console.log("externe"));
document.getElementById("interne").addEventListener("click", (e) => {
  console.log("interne");
  e.stopPropagation(); // stoppe la remontée
});
```

### Délégation d'événements (pattern clé)

Un seul listener sur le parent gère tous les enfants, **même ajoutés dynamiquement** :

```js
document.getElementById("liste-serveurs").addEventListener("click", (e) => {
  const bouton = e.target.closest("button[data-action]");
  if (!bouton) return; // clic hors bouton : ignorer
  const ligne = bouton.closest("tr");
  const ip = ligne.dataset.ip;
  if (bouton.dataset.action === "ping") pinger(ip);
  if (bouton.dataset.action === "supprimer") ligne.remove();
});
```

> Sans délégation : N listeners (fuites mémoire, rien sur les éléments futurs). Avec : 1 seul.

### Événements utiles

| Événement | Déclencheur |
|---|---|
| `click`, `dblclick` | souris |
| `input`, `change` | champ modifié (`input` = chaque frappe) |
| `submit` | formulaire |
| `keydown` / `keyup` | clavier (`e.key`) |
| `DOMContentLoaded` | DOM prêt (avant images) |
| `load` | page complète |

---

## 37. Formulaires

```js
const form = document.getElementById("form-equipement");

form.addEventListener("submit", async (e) => {
  e.preventDefault();

  // FormData : récupère tout le formulaire d'un coup
  const fd = new FormData(form);
  const equipement = Object.fromEntries(fd.entries());
  // { hostname: "SW9", ip: "10.0.0.9", site: "Paris" }

  // Validation native HTML5 : required, type="email", pattern, min/max...
  if (!form.checkValidity()) { form.reportValidity(); return; }

  const res = await fetch("/api/equipements", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(equipement),
  });
  if (!res.ok) { /* afficher l'erreur */ return; }
  form.reset(); // vider
});
```

### Lire les champs individuellement

```js
const ip = document.getElementById("ip").value.trim();
const actif = document.getElementById("actif").checked; // checkbox
const site = document.querySelector('input[name="site"]:checked')?.value; // radio
```

### Validation custom

```js
const champIp = document.getElementById("ip");
champIp.addEventListener("input", () => {
  const ok = /^(?:\d{1,3}\.){3}\d{1,3}$/.test(champIp.value);
  champIp.setCustomValidity(ok ? "" : "IPv4 invalide (ex : 10.0.0.1)");
});
```

---

## 38. Stockage navigateur : localStorage, cookies

```js
// localStorage : ~5 Mo, persistant, SYNCHRONE (petites données uniquement)
localStorage.setItem("theme", "sombre");
localStorage.getItem("theme");            // "sombre" (null si absent)
localStorage.removeItem("theme");

// Stocker un objet → JSON
const prefs = { theme: "sombre", refresh: 30 };
localStorage.setItem("prefs", JSON.stringify(prefs));
const chargees = JSON.parse(localStorage.getItem("prefs") ?? "{}");

// sessionStorage : idem mais vidé à la fermeture de l'onglet
```

### Tableau des stockages

|  | localStorage | sessionStorage | Cookies |
|---|---|---|---|
| Capacité | ~5 Mo | ~5 Mo | ~4 Ko |
| Durée | persistant | onglet | selon `Expires`/`Max-Age` |
| Envoyé au serveur | non | non | **oui** (chaque requête) |
| Usage | préférences UI, cache | état temporaire | sessions (HttpOnly !) |

### Cookies côté JS

```js
document.cookie = "theme=sombre; Max-Age=31536000; Path=/; SameSite=Lax";
// Lire : parser document.cookie ("a=1; b=2")
const cookies = Object.fromEntries(document.cookie.split("; ").map(c => c.split("=")));
```

> Cookies de session/auth : `HttpOnly` + `Secure` + `SameSite` → **inaccessibles en JS** (protection XSS). Les poser côté serveur.

---

## 39. Web APIs utiles

```js
// Presse-papiers
await navigator.clipboard.writeText("10.0.0.1");
const texte = await navigator.clipboard.readText();

// Géolocalisation (avec permission utilisateur)
navigator.geolocation.getCurrentPosition(
  (pos) => console.log(pos.coords.latitude, pos.coords.longitude),
  (err) => console.error(err.message)
);

// Notifications (avec permission)
if (await Notification.requestPermission() === "granted") {
  new Notification("Supervision", { body: "SW1 ne répond plus" });
}

// Fetch + AbortController : voir section 25
// WebSocket : temps réel
const ws = new WebSocket("wss://supervision.example.com/live");
ws.onmessage = (e) => afficher(JSON.parse(e.data));
ws.send(JSON.stringify({ action: "abonner", cible: "SW1" }));

// IntersectionObserver : lazy-load / infinite scroll
const obs = new IntersectionObserver((entrees) => {
  for (const e of entrees) if (e.isIntersecting) chargerPlus();
});
obs.observe(document.getElementById("sentinelle"));

// URL et URLSearchParams : parser/construire des URLs proprement
const u = new URL("https://api.example.com/equipements?site=paris&page=2");
u.searchParams.get("site");      // "paris"
u.searchParams.set("page", "3");
u.toString();
```

---

## 40. RegExp

```js
// Littéral /.../flags  ou  new RegExp("...", "flags")
const reIp = /^(?:\d{1,3}\.){3}\d{1,3}$/;
reIp.test("10.0.0.1");   // true (test rapide booléen)

"err: timeout sur eth0".match(/timeout sur (\w+)/); // ["timeout sur eth0", "eth0"] (groupe)
// matchAll (ES2020) : tous les matchs avec groupes
for (const m of "a1 b2".matchAll(/([a-z])(\d)/g)) console.log(m[1], m[2]);
```

### Flags

| Flag | Effet |
|---|---|
| `g` | global (tous les matchs) |
| `i` | insensible à la casse |
| `m` | `^`/`$` par ligne |
| `s` | `.` matche aussi `\n` (dotAll, ES2018) |
| `u` | unicode (ES2015) |
| `y` | sticky (ES2015) |

### Syntaxes essentielles

| Motif | Signifie |
|---|---|
| `\d` `\w` `\s` | chiffre, mot `[A-Za-z0-9_]`, espace |
| `\D` `\W` `\S` | inverses |
| `^` `$` | début / fin (de chaîne ou ligne avec `m`) |
| `*` `+` `?` `{2,4}` | 0+, 1+, 0/1, répétitions |
| `*?` `+?` | versions **non-gourmandes** (lazy) |
| `(...)` `(?:...)` | groupe capturant / non-capturant |
| `(?<nom>...)` | groupe nommé (ES2018) : `m.groups.nom` |
| `(?<=...)` `(?=...)` | lookbehind / lookahead (ES2018) |
| `[abc]` `[^abc]` | classe / négation |
| `a\|b` | alternative |
| `\.` `\*` `\\` | caractères spéciaux échappés |

### Recettes sysadmin

```js
const IPV4 = /^(?:(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)\.){3}(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)$/;
const MAC = /^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$/;
// Extraire les IP d'un texte
const ips = texte.match(/(?:\d{1,3}\.){3}\d{1,3}/g) ?? [];
// Parser une ligne de log Apache (combined)
const RE_APACHE = /^(\S+) \S+ \S+ \[([^\]]+)\] "(\S+) (\S+) \S+" (\d{3}) (\d+|-) "[^"]*" "([^"]*)"/;
```

### Piège : regex avec `g` et `.test()` en boucle

```js
const re = /a/g;
re.test("a"); // true  — lastIndex avance !
re.test("a"); // false — repart de lastIndex=1
// Solution : ne pas réutiliser une regex /g/ avec test(), ou reset re.lastIndex = 0
```

---

## 41. Map, Set, WeakMap, WeakSet

```js
// Map : clés de TOUT type (objets !), ordre d'insertion, .size
const inventaire = new Map();
inventaire.set("SW1", { ip: "10.0.0.2", ports: 48 });
inventaire.set("SW1", { ip: "10.0.0.3", ports: 24 }); // écrase (clé unique)
inventaire.get("SW1");       // { ip: "10.0.0.3", ports: 24 }
inventaire.has("R1");        // false
inventaire.delete("SW1");
inventaire.size;             // 0
for (const [nom, fiche] of inventaire) { /* ... */ }

// Set : valeurs UNIQUES
const ipsVues = new Set();
ipsVues.add("10.0.0.1"); ipsVues.add("10.0.0.1"); // doublon ignoré
ipsVues.has("10.0.0.1");     // true
[...new Set([1, 2, 2, 3])];   // [1, 2, 3] — dédupliquer un tableau !

