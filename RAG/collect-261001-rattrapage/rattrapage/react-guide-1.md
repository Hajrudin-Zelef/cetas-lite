---
id: collect-261001-rattrapage/rattrapage/react-guide-1
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [1, 222]
sha256: 065d5f730a2cb829ebec0871198091c504b990408c7b155e7fb92e6b0873a3ce
---

# Guide React — Le manuel complet

> **Public visé :** Zelef, chef de service systèmes & énergies — dashboards internes, outils web de supervision.
> **Ton :** direct, tutoriel + référence. **Version :** React 18+ (notes React 19 signalées quand pertinent).
> **Règle d'or de ce guide :** tout le code est moderne (composants fonctionnels, hooks). Les classes ne sont mentionnées que pour comprendre le code hérité.

---

## Sommaire

1. Pourquoi React pour un SI interne
2. Prérequis : Node.js, npm, éditeur
3. Créer un projet avec Vite
4. Anatomie d'un projet Vite + React
5. JSX — la syntaxe au cœur de React
6. JSX — expressions, attributs, enfants
7. JSX — les 3 règles de base
8. Composants fonctionnels
9. Props en détail
10. Composition vs héritage
11. L'état avec useState
12. useState — pièges et bonnes pratiques
13. Rendu conditionnel
14. Événements en React
15. Formulaires contrôlés
16. Listes et keys
17. useEffect — dépendances et cleanup
18. useEffect — les patterns courants
19. useRef — au-delà du DOM
20. useMemo — mémoriser un calcul
21. useCallback — mémoriser une fonction
22. React.memo — éviter les re-rendus inutiles
23. useContext et la Context API
24. useReducer — l'état complexe
25. Hooks personnalisés
26. Les règles des hooks (linter)
27. Data fetching — patterns loading / error
28. Data fetching — exemple complet
29. Error Boundaries — attraper les plantages
30. Modules — organiser son code
31. React Router — routes de base
32. React Router — params, navigation, layout
33. Vite — configurer et builder
34. Debug avec React DevTools
35. Debug — stratégies et console
36. Tests avec Vitest + Testing Library
37. Tests — exemples commentés
38. TypeScript + React
39. Sécurité — XSS et injections
40. Bonnes pratiques & style de code
41. Performance — comprendre le rendu
42. Performance — lazy, Suspense, code splitting
43. Performance — virtualisation des longues listes
44. Intro Next.js — SSR et SSG
45. Erreurs classiques des débutants (15)
46. Les erreurs React les plus fréquentes — corrections
47. Cas pratique 1 — dashboard de supervision
48. Cas pratique 2 — formulaire de ticket d'intervention
49. Cas pratique 3 — liste d'équipements filtrable
50. Checklist — mettre une app en production
51. Pense-bête de poche
52. Glossaire
53. Quiz — 10 questions + réponses
54. Pour aller plus loin

---

## 1. Pourquoi React pour un SI interne

React est une bibliothèque JavaScript (Meta) pour construire des interfaces utilisateur par **composants réutilisables**. Pour un service systèmes & énergies, les cas d'usage typiques sont :

- **Dashboards de supervision** : onduleurs, températures, états d'équipements, alertes temps réel.
- **Outils internes** : formulaires d'intervention, inventaires, suivi de tickets, GMAO simplifiée.
- **Portails** : pages d'accueil intranet, documentation, tableaux de bord KPI.

Pourquoi React plutôt qu'un site "classique" en HTML/JS vanilla ou PHP ?

| Critère | React | JS vanilla / jQuery |
|---|---|---|
| Mise à jour du DOM | Déclarative : on décrit l'état, React met à jour | Impérative : on manipule le DOM à la main |
| Réutilisabilité | Composants (bouton, carte, tableau) | Copier-coller de code |
| Écosystème | Énorme (charts, tables, formulaires) | Fragmenté |
| Tests | Vitest + Testing Library, standardisés | Artisanal |
| Courbe d'apprentissage | Moyenne (JSX, hooks) | Faible au début, chaotique ensuite |

**Le modèle mental à adopter :** en React, l'UI est une **fonction de l'état**. `UI = f(state)`. Quand l'état change (nouvelle mesure d'un onduleur), React recalcule et met à jour uniquement ce qui a changé. On ne touche (presque) jamais au DOM directement.

**Quand NE PAS utiliser React :** une page statique de documentation, un simple formulaire de contact → un site statique ou un template serveur suffit. React se justifie dès qu'il y a de l'interactivité et de l'état.

---

## 2. Prérequis : Node.js, npm, éditeur

React s'écrit en JavaScript moderne (ES2015+). Avant tout :

### 2.1 Installer Node.js

Node.js permet d'exécuter JavaScript hors navigateur et fournit **npm** (gestionnaire de paquets).

- Télécharger la version **LTS** sur https://nodejs.org (vérifiée : LTS conseillée pour la stabilité).
- Vérifier l'installation :

```bash
node -v   # v20.x.x ou v22.x.x attendu
npm -v    # 10.x.x attendu
```

> Sur un poste Windows d'entreprise, préférer l'installeur MSI officiel. Éviter les versions "Current" en production.

### 2.2 L'éditeur : VS Code

Extensions indispensables :

| Extension | Rôle |
|---|---|
| ES7+ React/Redux/React-Native snippets | Snippets (`rafce` → composant) |
| ESLint | Détecte les erreurs, **dont les règles des hooks** |
| Prettier | Formatage automatique |
| Error Lens | Affiche les erreurs inline |

### 2.3 JavaScript moderne requis

Si ces syntaxes sont floues, les revoir avant de continuer :

```js
// Déstructuration
const { nom, puissance } = onduleur;

// Spread
const copie = { ...onduleur, puissance: 120 };

// Arrow functions
const double = (x) => x * 2;

// Template literals
const msg = `UPS ${nom} : ${puissance} kVA`;

// Modules ES
import { useState } from 'react';
export function Carte() { /* ... */ }

// Optional chaining + nullish coalescing
const ville = site?.adresse?.ville ?? 'Inconnue';
```

**Checklist prérequis :**
- [ ] Node LTS installé, `node -v` OK
- [ ] VS Code + ESLint + Prettier
- [ ] À l'aise avec les arrow functions, la déstructuration, les modules ES

---

## 3. Créer un projet avec Vite

**Vite** est l'outil de build officiellement recommandé pour démarrer un projet React (Create React App est abandonné depuis 2023 — ne plus l'utiliser).

```bash
# Créer le projet (répondre aux questions : React, JavaScript ou TypeScript)
npm create vite@latest supervision-dashboard -- --template react

cd supervision-dashboard
npm install     # installe les dépendances
npm run dev     # lance le serveur de dev → http://localhost:5173
```

Commandes du quotidien :

| Commande | Effet |
|---|---|
| `npm run dev` | Serveur de développement (rechargement instantané) |
| `npm run build` | Build de production dans `dist/` |
| `npm run preview` | Prévisualise le build de production en local |
| `npm install <paquet>` | Ajoute une dépendance |

> **Pourquoi Vite et pas autre chose ?** Démarrage en < 1 s, rechargement à chaud (HMR) quasi instantané même sur gros projet, config simple. Pour du SSR (rendu serveur), voir la section 44 (Next.js).

**Variante TypeScript** (recommandée pour un projet d'équipe) :

```bash
npm create vite@latest supervision-dashboard -- --template react-ts
```

---

## 4. Anatomie d'un projet Vite + React

```
supervision-dashboard/
├── index.html              # point d'entrée HTML (Vite l'utilise directement)
├── package.json            # dépendances + scripts
├── vite.config.js          # config Vite
├── src/
│   ├── main.jsx            # point d'entrée JS : monte <App/> dans le DOM
│   ├── App.jsx             # composant racine
│   ├── App.css
│   ├── index.css           # styles globaux
│   └── assets/             # images, etc.
```

**`src/main.jsx`** — le bootstrap, on n'y touche presque jamais :

```jsx
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.jsx';
import './index.css';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>
);
```

**`src/App.jsx`** — le composant racine, c'est ici que tout commence :

```jsx
export default function App() {
  return <h1>Supervision UPS</h1>;
}
```

