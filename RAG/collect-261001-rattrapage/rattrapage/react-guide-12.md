---
id: collect-261001-rattrapage/rattrapage/react-guide-12
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [2730, 2865]
sha256: 32a4d94bc68eed55ca345b60dc8a882862b17412106e78f5115445603779144e
---

# Guide React — Le manuel complet

### Build & tests
- [ ] `npm run build` OK, tester avec `npm run preview`
- [ ] `npm run test:run` : tous les tests passent
- [ ] Taille du bundle vérifiée (découper avec `lazy` si > 500 Ko initial)

### Robustesse
- [ ] Error Boundaries sur les zones à risque (graphiques, widgets)
- [ ] Tous les appels API gèrent loading / error / empty
- [ ] Page 404 (`path="*"`) définie
- [ ] Formulaires : validation + anti double-submit

### Sécurité
- [ ] Pas de `dangerouslySetInnerHTML` non sanitizé
- [ ] `rel="noopener noreferrer"` sur les `target="_blank"`
- [ ] `npm audit` sans vulnérabilité critique

### Déploiement (SPA statique)
- [ ] Servir `dist/` via Nginx/Apache
- [ ] **Fallback SPA** : toute route inconnue → `index.html` (sinon refresh sur `/tickets` = 404 serveur). Ex. Nginx : `try_files $uri /index.html;`
- [ ] En-têtes cache : assets hashés = cache long ; `index.html` = `no-cache`

---

## 51. Pense-bête de poche

### Créer & lancer
```bash
npm create vite@latest mon-app -- --template react
npm install && npm run dev
```

### Composant minimal
```jsx
export default function MonComposant({ titre }) {
  return <h1>{titre}</h1>;
}
```

### Les 8 hooks essentiels
| Hook | Usage en une phrase |
|---|---|
| `useState` | Mémoire du composant → re-rend quand ça change |
| `useEffect` | Effet de bord après rendu (API, timer, abonnement) |
| `useRef` | Réf DOM ou boîte mutable sans re-rendu |
| `useMemo` | Mémorise un **calcul** coûteux |
| `useCallback` | Mémorise une **fonction** (props stables) |
| `useContext` | Lit un contexte (sans prop drilling) |
| `useReducer` | État complexe à transitions nommées |
| `useId` | (bonus) Génère un id unique accessible |

### Les 5 réflexes
1. `UI = f(state)` — décrire l'état, pas manipuler le DOM.
2. Immuable : `[...t, x]`, `{...o, k: v}` — jamais `push` / assignation directe.
3. `useEffect` : toujours un tableau de dépendances + cleanup si abonnement.
4. Listes : `key` stable et unique (jamais l'index sur liste mutable).
5. Dériver pendant le rendu plutôt que dupliquer dans l'état.

### Snippets VS Code (extension ES7+)
- `rafce` → composant fonctionnel avec export
- `useState` → `const [x, setX] = useState()`
- `useEffect` → squelette d'effet

---

## 52. Glossaire

| Terme | Définition |
|---|---|
| **Babel / esbuild** | Transpilers : convertissent JSX/TS en JS compris par le navigateur |
| **Batching** | Regroupement de plusieurs `setState` en un seul re-rendu |
| **Bundler** | Outil qui assemble les modules en fichiers optimisés (Vite utilise Rollup en build) |
| **Cleanup** | Fonction retournée par `useEffect`, exécutée avant ré-exécution et au démontage |
| **Client Component** | (Next.js) Composant interactif exécuté dans le navigateur (`'use client'`) |
| **Code splitting** | Découpage du bundle en morceaux chargés à la demande (`React.lazy`) |
| **Composant contrôlé** | Champ de formulaire dont la valeur est pilotée par l'état React |
| **CSR** | Client-Side Rendering : le HTML est construit dans le navigateur |
| **Dépendance (effet)** | Valeur du tableau de `useEffect` qui redéclenche l'effet quand elle change |
| **Hydratation** | (SSR) React "réhydrate" le HTML serveur en y attachant l'interactivité |
| **HMR** | Hot Module Replacement : rechargement à chaud sans perdre l'état (Vite) |
| **Hook** | Fonction `use*` qui branche un composant aux fonctionnalités React (état, effets…) |
| **Immuabilité** | Ne jamais modifier un objet/tableau existant : en créer un nouveau |
| **JSX** | Extension de syntaxe : écrire du HTML-like dans du JavaScript |
| **Key** | Identifiant stable aidant React à suivre les éléments d'une liste |
| **Lifting state up** | Remonter l'état vers le parent commun quand plusieurs enfants le partagent |
| **Memoïsation** | Mettre en cache un résultat pour éviter de le recalculer (`memo`, `useMemo`) |
| **Prop drilling** | Passer des props à travers des composants intermédiaires qui n'en ont pas besoin |
| **Props** | Données passées du parent à l'enfant (lecture seule) |
| **Réconciliation** | Algorithme de diff entre deux arbres virtuels |
| **Server Component** | (Next.js) Composant exécuté sur le serveur, sans JS envoyé au client |
| **SPA** | Single Page Application : une seule page HTML, navigation côté client |
| **SSG** | Static Site Generation : HTML généré une fois au build |
| **SSR** | Server-Side Rendering : HTML généré sur le serveur à chaque requête |
| **State** | État local d'un composant ; tout changement déclenche un re-rendu |
| **StrictMode** | Mode dev qui révèle les effets de bord (double invocation) |
| **SyntheticEvent** | Wrapper d'événement normalisé par React |
| **Virtual DOM** | Représentation mémoire du DOM utilisée pour calculer les différences |
| **XSS** | Cross-Site Scripting : injection de script via des données non échappées |

---

## 53. Quiz — 10 questions + réponses

**Q1. Que se passe-t-il quand on appelle `setCompteur(compteur + 1)` deux fois de suite ?**
> Le compteur n'augmente que de 1 : les deux appels capturent l'ancienne valeur (batching). Solution : forme fonctionnelle `setCompteur(c => c + 1)`.

**Q2. Pourquoi ce code est-il interdit ? `if (actif) { const [x] = useState(0); }`**
> Les hooks doivent être appelés dans le même ordre à chaque rendu, au niveau racine. Un hook conditionnel casse cet ordre → états mélangés. Mettre la condition *dans* le hook.

**Q3. À quoi sert le tableau de dépendances de `useEffect` ?**
> À déclarer les valeurs dont dépend l'effet : `[]` = une seule fois au montage ; `[a, b]` = au montage + quand `a` ou `b` change ; absent = après chaque rendu (rarement voulu).

**Q4. Pourquoi ne faut-il pas utiliser l'index comme `key` dans une liste modifiable ?**
> React identifie les éléments par leur `key`. Avec l'index, une suppression/réordonnancement fait croire à React que c'est un autre élément qui a changé → états locaux des lignes mélangés, bugs fantômes.

**Q5. Quelle est la différence entre `useMemo` et `useCallback` ?**
> `useMemo` mémorise le *résultat* d'un calcul ; `useCallback` mémorise une *fonction* (référence stable). `useCallback(fn, deps)` équivaut à `useMemo(() => fn, deps)`.

**Q6. Pourquoi ce champ est-il figé ? `<input type="text" value={nom} />`**
> Champ contrôlé sans `onChange` : React impose `value` à chaque rendu, l'utilisateur ne peut rien saisir. Ajouter `onChange={(e) => setNom(e.target.value)}`.

**Q7. Comment éviter qu'une erreur dans un widget ne fasse écran blanc sur tout le dashboard ?**
> Envelopper chaque zone à risque dans une **Error Boundary** (composant classe avec `getDerivedStateFromError` / `componentDidCatch`) qui affiche une UI de secours.

**Q8. `const [v, setV] = useState(0); setV(1); console.log(v);` — qu'affiche le log ?**
> `0` : `setV` ne modifie pas la variable immédiatement, il planifie un nouveau rendu où `v` vaudra 1.

**Q9. Quand utiliser `useReducer` plutôt que `useState` ?**
> Quand l'état comporte plusieurs sous-valeurs liées ou des transitions complexes : le reducer centralise les transitions nommées (`{ type: 'PANNE' }`), testables car fonction pure.

**Q10. Pourquoi `<a href="/tickets">` est-il déconseillé pour la navigation interne ?**
> Il recharge toute la page : on perd l'état React et on retélécharge le bundle. Utiliser `<Link to="/tickets">` de React Router (navigation côté client).

---

## 54. Pour aller plus loin

