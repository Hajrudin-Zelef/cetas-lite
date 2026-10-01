---
id: collect-261001-rattrapage/rattrapage/react-guide-9
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [2075, 2282]
sha256: 2d38167280a72162a14e2c05d2b675b92bfaa6633167ac8d98e0aafa85f0fd6a
---

# Guide React — Le manuel complet

| Type | Usage |
|---|---|
| `React.ReactNode` | `children` quelconque |
| `React.FC<Props>` | Type de composant (optionnel, controversé) |
| `React.ChangeEvent<HTMLInputElement>` | `onChange` d'un input |
| `React.FormEvent` | `onSubmit` |

**Conseil :** typer les **frontières** (props, retours d'API, fonctions exportées) ; laisser l'inférence faire le reste. Un `any` occasionnel et assumé vaut mieux qu'un typage mensonger.

---

## 39. Sécurité — XSS et injections

### React protège par défaut contre le XSS

Tout ce qui passe dans `{}` est **échappé** automatiquement : impossible d'injecter du HTML par une variable.

```jsx
const nomSaisi = '<img src=x onerror=alert(1)>';
<p>{nomSaisi}</p>
// ✅ Affiche le texte littéralement — le script NE s'exécute PAS
```

### Le danger : `dangerouslySetInnerHTML`

```jsx
// ☠️ À ÉVITER sauf contenu 100 % maîtrisé
<div dangerouslySetInnerHTML={{ __html: htmlBrut }} />
```

**Si on doit afficher du HTML riche** (ex. description d'équipement avec mise en forme) : le **sanitizer** avec une librairie comme **DOMPurify** :

```bash
npm install dompurify
```

```jsx
import DOMPurify from 'dompurify';

<div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(htmlBrut) }} />
```

### Checklist sécurité front

- [ ] Jamais de `dangerouslySetInnerHTML` sur du contenu utilisateur (ou + DOMPurify)
- [ ] Pas de secret dans le code front ni dans les variables `VITE_*` (visibles par tous)
- [ ] Valider aussi côté **serveur** : le front n'est jamais une barrière de sécurité
- [ ] Liens externes : `target="_blank"` **toujours** avec `rel="noopener noreferrer"` (évite le tabnabbing)
- [ ] `npm audit` régulier ; dépendances à jour (faille = souvent une lib tierce)
- [ ] Échapper les URL construites : `encodeURIComponent(id)` dans les `fetch`

```jsx
// ✅ Lien externe sécurisé
<a href="https://doc.interne" target="_blank" rel="noopener noreferrer">
  Documentation
</a>
```

---

## 40. Bonnes pratiques & style de code

### Organisation

- **Un composant = un fichier**, nommé comme le composant (`CarteOnduleur.jsx`).
- **Colocalisation** : le style, les tests et le composant vivent ensemble (`CarteOnduleur.jsx`, `.css`, `.test.jsx`).
- Extraire un composant dès qu'un bloc de JSX est **réutilisé 2 fois** ou dépasse ~50 lignes.

### Nommage

| Élément | Convention | Exemple |
|---|---|---|
| Composant | PascalCase | `CarteOnduleur` |
| Hook personnalisé | camelCase, préfixe `use` | `usePolling` |
| Fonction / variable | camelCase | `chargerMesures` |
| Constante globale | SNAKE_CASE | `SEUIL_CRITIQUE` |
| Gestionnaire d'événement | `gerer…` / `on…` (props) | `gererClic`, prop `onAlerte` |

### Règles d'or

1. **Ne jamais muter** l'état ou les props (section 12).
2. **État au plus bas niveau possible** : remonter ("lifter") uniquement si plusieurs composants en ont besoin.
3. **Dériver plutôt que dupliquer** : pas d'état calculable depuis un autre.
4. **Pas de logique métier dans le JSX** : extraire des fonctions / utilitaires.
5. **Keys stables** sur les listes (section 16).
6. **Nettoyer les effets** (section 17).
7. **Prévenir plutôt que guérir** : ESLint + TypeScript + tests sur les chemins critiques.
8. **Accessibilité de base** : `<button>` pour les actions (pas de `<div onClick>`), `label` pour chaque input, `alt` sur les images.

### Exemple de composant "propre"

```jsx
import { usePolling } from '@/hooks/usePolling';
import { formatPct } from '@/utils/format';
import './CarteUps.css';

const SEUIL_CRITIQUE = 80;

export default function CarteUps({ id }) {
  const { data: ups, erreur, chargement } = usePolling(`/api/ups/${id}`, 5000);

  if (chargement) return <div className="carte">⏳…</div>;
  if (erreur) return <div className="carte erreur">❌ {erreur}</div>;

  const critique = ups.charge > SEUIL_CRITIQUE;

  return (
    <article className={`carte ${critique ? 'critique' : ''}`}>
      <h2>{ups.nom}</h2>
      <p>Charge : {formatPct(ups.charge)}</p>
      {critique && <p className="alerte">⚠️ Charge critique</p>}
    </article>
  );
}
```

---

## 41. Performance — comprendre le rendu

### Le cycle de rendu

1. **Déclencheur** : un `setState` (ou nouvelles props du parent).
2. **Render** : React ré-exécute le composant et construit un arbre virtuel (Virtual DOM).
3. **Réconciliation** : React compare (diff) avec l'arbre précédent.
4. **Commit** : seules les différences sont appliquées au vrai DOM.

Le Virtual DOM n'est pas "magique" : un composant qui rend 10 000 lignes à chaque frappe clavier **restera lent**. Les optimisations :

| Levier | Quand |
|---|---|
| État au bon niveau | Toujours (évite les re-rendus en cascade) |
| `React.memo` | Composant coûteux re-rendu avec mêmes props |
| `useMemo` / `useCallback` | Calcul coûteux / props de composants mémorisés |
| `lazy` + `Suspense` | Code splitting par page |
| Virtualisation | Listes de centaines/milliers de lignes |
| Pagination côté serveur | Listes très longues (mieux que tout) |

### Mesurer avant d'optimiser

1. Onglet **Profiler** de React DevTools : enregistrer une interaction, repérer les composants lents / trop souvent rendus.
2. *"Highlight updates"* : visualiser les re-rendus en cascade.
3. `console.time` autour d'un calcul suspect.

**Règle :** 95 % des "lenteurs" React viennent d'un **état placé trop haut** (ex. le champ de recherche dans le même composant que le tableau de 5 000 lignes → chaque frappe re-rend tout). Descendre l'état ou séparer les composants règle le problème sans `memo`.

---

## 42. Performance — lazy, Suspense, code splitting

`React.lazy` charge un composant **à la demande** (chunk JS séparé) au lieu de l'inclure dans le bundle initial. `Suspense` affiche un fallback pendant le chargement.

```jsx
import { lazy, Suspense } from 'react';
import { Routes, Route } from 'react-router-dom';

// Ces pages ne sont téléchargées que quand on les visite
const Dashboard = lazy(() => import('./pages/Dashboard.jsx'));
const Rapports = lazy(() => import('./pages/Rapports.jsx')); // gros graphiques

function App() {
  return (
    <Suspense fallback={<p>⏳ Chargement de la page…</p>}>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/rapports" element={<Rapports />} />
      </Routes>
    </Suspense>
  );
}
```

**Effet :** le bundle initial passe par ex. de 800 Ko à 200 Ko → première page affichée bien plus vite, crucial sur un intranet avec des postes modestes.

**Bonnes pratiques :**
- Découper **par route** en priorité (le plus rentable).
- Découper les **gros composants rarement utilisés** (éditeur riche, export PDF, graphiques lourds).
- Le `fallback` doit être léger (spinner CSS, pas un composant lourd).
- Combiner avec une Error Boundary : un chunk qui échoue à charger (déploiement entre-temps) ne doit pas écran-blanc l'app.

> **Note :** `React.lazy` ne fonctionne qu'avec les **exports default**. Avec Vite, on peut aussi utiliser l'import dynamique manuel pour précharger : `import('./pages/Rapports.jsx')` au survol d'un lien.

---

## 43. Performance — virtualisation des longues listes

Afficher 10 000 lignes d'historique d'alarmes dans le DOM = lent (mémoire + layout). La **virtualisation** ne rend que les lignes **visibles** (+ une marge).

Librairie standard : **TanStack Virtual** (ex `react-virtual`) ou `react-window`.

```bash
npm install @tanstack/react-virtual
```

```jsx
import { useRef } from 'react';
import { useVirtualizer } from '@tanstack/react-virtual';

function JournalAlertes({ alertes }) { // alertes = tableau de 10 000 entrées
  const parentRef = useRef(null);

  const virtualizer = useVirtualizer({
    count: alertes.length,
    getScrollElement: () => parentRef.current,
    estimateSize: () => 40, // hauteur estimée d'une ligne (px)
    overscan: 10,           // lignes de marge hors viewport
  });

