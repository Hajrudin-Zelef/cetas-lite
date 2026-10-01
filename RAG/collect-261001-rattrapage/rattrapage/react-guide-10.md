---
id: collect-261001-rattrapage/rattrapage/react-guide-10
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition"]
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [2283, 2496]
sha256: 28f70261d5cc9d02860d3ac5893ed23c2858e9d8abe8e918458830a486c9f152
---

# Guide React — Le manuel complet

  return (
    <div ref={parentRef} style={{ height: 400, overflow: 'auto' }}>
      <div style={{ height: virtualizer.getTotalSize(), position: 'relative' }}>
        {virtualizer.getVirtualItems().map((ligne) => {
          const alerte = alertes[ligne.index];
          return (
            <div
              key={alerte.id}
              style={{
                position: 'absolute', top: 0, left: 0, width: '100%',
                height: ligne.size, transform: `translateY(${ligne.start}px)`,
              }}
            >
              {alerte.date} — {alerte.message}
            </div>
          );
        })}
      </div>
    </div>
  );
}
```

**Quand virtualiser :** à partir de quelques centaines de lignes complexes, ou milliers de lignes simples. En dessous, un `.map()` classique suffit.

**Alternative souvent meilleure :** la **pagination côté serveur** (`?page=2&limit=50`) — moins de données transférées ET moins de DOM.

---

## 44. Intro Next.js — SSR et SSG

**Next.js** est le framework React de Vercel : routage basé sur les fichiers, rendu côté serveur, build optimisé. À connaître même si on reste sur Vite.

### Les 3 modes de rendu

| Mode | Où le HTML est généré | Cas d'usage |
|---|---|---|
| **CSR** (Client-Side Rendering) | Dans le navigateur (Vite/SPA classique) | Dashboards internes, apps derrière login |
| **SSR** (Server-Side Rendering) | Sur le serveur à chaque requête | Données fraîches + SEO |
| **SSG** (Static Site Generation) | Au build, une fois | Documentation, pages publiques, SEO max |

**Pour un dashboard interne :** le CSR (Vite) suffit largement — pas besoin de SEO, l'utilisateur est authentifié. Next.js se justifie pour un portail public, de la doc indexée, ou des performances premier affichage critiques.

### Concepts clés Next.js (App Router, v13+)

```
app/
├── page.jsx            # route "/"
├── tickets/
│   └── page.jsx        # route "/tickets"
├── layout.jsx          # layout partagé
└── onduleurs/[id]/
    └── page.jsx        # route dynamique "/onduleurs/ups-01"
```

- **Server Components** (par défaut) : s'exécutent sur le serveur, peuvent appeler la BDD directement, n'envoient pas leur JS au client.
- **Client Components** (`'use client'` en haut du fichier) : pour l'interactivité (`useState`, `onClick`) — c'est le React qu'on connaît.

```jsx
// app/onduleurs/page.jsx — Server Component : fetch direct, zéro JS client
export default async function Onduleurs() {
  const res = await fetch('https://api.interne/onduleurs');
  const data = await res.json();
  return (
    <ul>{data.map((u) => <li key={u.id}>{u.nom}</li>)}</ul>
  );
}
```

**Verdict pragmatique :** maîtriser React + Vite d'abord. Next.js quand un besoin SSR/SSG réel apparaît.

---

## 45. Erreurs classiques des débutants (15)

1. **Oublier les accolades pour une valeur non-string** : `puissance="40"` passe la *string* `"40"`, pas le nombre. → `puissance={40}`.
2. **`class` au lieu de `className`** : warning + style non appliqué.
3. **Appeler la fonction dans le gestionnaire** : `onClick={maFonction()}` l'exécute au rendu. → `onClick={maFonction}` ou `onClick={() => maFonction(arg)}`.
4. **Mutations directes de l'état** : `tableau.push(x)` puis `setTableau(tableau)` → rien ne se met à jour. → créer un nouveau tableau/objet.
5. **Lire l'état juste après `setState`** : la valeur n'est pas encore à jour (batching). → utiliser la valeur calculée localement ou un effet.
6. **`useEffect` sans tableau de dépendances** : tourne à chaque rendu → boucles infinies avec `setState` dedans.
7. **Oublier le cleanup** : `setInterval` / WebSocket qui survivent à la navigation → fuites mémoire.
8. **L'index comme `key`** dans une liste modifiable → états des lignes mélangés.
9. **Props drilling extrême** : passer des props sur 6 niveaux → penser Context ou recomposition.
10. **Tout mettre dans un seul composant géant** : 400 lignes de JSX → découper en sous-composants.
11. **`<a href>` pour la navigation interne** : recharge la page, tue l'état React. → `<Link>`.
12. **Champ contrôlé sans `onChange`** : `value={x}` seul = champ figé + warning.
13. **Afficher un objet brut** : `<p>{monObjet}</p>` → crash *"Objects are not valid as a React child"*. → `{monObjet.nom}` ou `JSON.stringify`.
14. **Le piège du `&&` avec `0`** : `{compte && <X/>}` affiche `"0"`. → `{compte > 0 && <X/>}`.
15. **Fetch dans le corps du composant** (hors `useEffect`) : requête relancée à chaque rendu → boucle. → `useEffect` + dépendances.

---

## 46. Les erreurs React les plus fréquentes — corrections

### "Too many re-renders"

```jsx
// ❌ setCompteur exécuté PENDANT le rendu → boucle infinie
<button onClick={setCompteur(compteur + 1)}>+1</button>

// ✅ Fonction passée, exécutée AU CLIC
<button onClick={() => setCompteur(compteur + 1)}>+1</button>
```

Autre cause : `setState` appelé directement dans le corps du composant (hors gestionnaire/effet).

### "Cannot read properties of undefined (reading 'nom')"

```jsx
// ❌ ups est null pendant le chargement
<h2>{ups.nom}</h2>

// ✅ Garde ou optional chaining
<h2>{ups?.nom ?? 'Chargement…'}</h2>
if (!ups) return <p>Chargement…</p>;
```

### "Each child in a list should have a unique key"

→ Ajouter `key={item.id}` stable (section 16). Si vraiment aucun id : `key={item.nom + index}` en dernier recours, jamais l'index seul sur liste mutable.

### "Objects are not valid as a React child"

```jsx
// ❌
<p>{erreur}</p>              // erreur est un objet Error !
// ✅
<p>{erreur.message}</p>
```

### Warning `exhaustive-deps`

```jsx
useEffect(() => {
  console.log(filtre); // 'filtre' utilisé mais absent des dépendances
}, []); // ⚠️ valeur obsolète capturée

// ✅
}, [filtre]);
```

### "X is not a function" sur une prop callback

La prop n'a pas été passée par le parent → valeur par défaut ou garde :

```jsx
function Enfant({ onValider = () => {} }) { /* ... */ }
// ou
onValider?.();
```

---

## 47. Cas pratique 1 — dashboard de supervision

Un mini-dashboard : cartes d'onduleurs avec polling, badge critique, temps réel simulé. Fichiers : `services/api.js`, `hooks/usePolling.js` (section 25), composants ci-dessous.

```jsx
// components/Dashboard.jsx
import { usePolling } from '../hooks/usePolling.js';
import CarteUps from './CarteUps.jsx';
import ResumeParc from './ResumeParc.jsx';

export default function Dashboard() {
  const { data: parc, erreur, chargement } = usePolling('/api/parc', 10000);

  if (chargement) return <p>⏳ Connexion à la supervision…</p>;
  if (erreur) return (
    <div className="erreur">
      <p>❌ Supervision injoignable : {erreur}</p>
      <p>Vérifiez le réseau et le service d'acquisition.</p>
    </div>
  );

  return (
    <main className="dashboard">
      <h1>Supervision énergétique</h1>
      <ResumeParc onduleurs={parc.onduleurs} />
      <div className="grille">
        {parc.onduleurs.map((ups) => (
          <CarteUps key={ups.id} ups={ups} />
        ))}
      </div>
      <p className="maj">Dernière mise à jour : {new Date(parc.horodatage).toLocaleTimeString('fr-FR')}</p>
    </main>
  );
}
```

```jsx
// components/ResumeParc.jsx — agrégats calculés pendant le rendu
import { useMemo } from 'react';

export default function ResumeParc({ onduleurs }) {
  const resume = useMemo(() => {
    const total = onduleurs.length;
    const enAlarme = onduleurs.filter((u) => u.charge > 80 || !u.enLigne).length;
    const chargeMoy = total
      ? onduleurs.reduce((s, u) => s + u.charge, 0) / total
      : 0;
    return { total, enAlarme, chargeMoy };
  }, [onduleurs]);

  return (
    <section className="resume">
      <div>🔌 {resume.total} onduleurs</div>
      <div>⚠️ {resume.enAlarme} en alarme</div>
      <div>📊 Charge moyenne : {resume.chargeMoy.toFixed(1)} %</div>
    </section>
  );
}
```

