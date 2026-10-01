---
id: collect-261001-rattrapage/rattrapage/react-guide-8
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [1830, 2074]
sha256: 9bbb3d2409ed7a4786320c28aa2a23162f29080d6dd63ee866b2910a34cda1e0
---

# Guide React — Le manuel complet

### Variables d'environnement

- Fichier `.env` à la racine : `VITE_API_URL=https://api.interne.local`
- **Obligatoirement** préfixées par `VITE_` pour être exposées au navigateur.
- Accès : `import.meta.env.VITE_API_URL` (jamais `process.env`, c'est Node).

```bash
# .env.development
VITE_API_URL=http://localhost:3000

# .env.production
VITE_API_URL=https://api.supervision.interne
```

> ⚠️ Tout ce qui est préfixé `VITE_` est **embarqué dans le JS envoyé au navigateur** : jamais de secret/mot de passe dedans.

### Build de production

```bash
npm run build    # → dossier dist/ (HTML + JS + CSS minifiés, hashés)
npm run preview  # teste le build en local
```

Le contenu de `dist/` est **statique** : à servir avec Nginx, Apache, ou copié sur un serveur de fichiers interne. Pas besoin de Node en production pour une SPA Vite.

**Checklist build :**
- [ ] `npm run build` passe sans erreur ni warning bloquant
- [ ] Variables `VITE_*` définies pour la cible
- [ ] Tester avec `npm run preview` avant déploiement

---

## 34. Debug avec React DevTools

**React DevTools** : extension navigateur (Chrome/Firefox/Edge) indispensable. Deux onglets :

### Onglet Components

- Arborescence des composants (noms, pas divs anonymes).
- **Props et hooks** de chaque composant inspectés en direct (valeurs de `useState`, dépendances).
- *"Highlight updates"* : surligne les composants qui se re-rendent → repère les rendus en cascade.
- On peut **modifier une prop ou un état** à la volée pour tester.

### Onglet Profiler

- Enregistre une session → montre **quel composant a rendu, combien de fois, combien de temps**.
- Le workflow d'optimisation : *mesurer d'abord* avec le Profiler, *ensuite* `memo`/`useMemo` ciblés.

**Bonnes pratiques :**
- Nommer les composants (pas de `export default () => …` anonyme : le DevTools affiche "Anonymous").
- En dev, `StrictMode` double certains effets : un double appel API au montage en dev est **normal** (pas en prod).

---

## 35. Debug — stratégies et console

### La méthode en 5 étapes

1. **Reproduire** : le plus petit cas qui déclenche le bug.
2. **Lire l'erreur** : la console donne le composant + la stack. Les erreurs React ont un lien vers la doc.
3. **Isoler** : commenter / extraire le composant suspect.
4. **Vérifier les hypothèses** : `console.log` des props et de l'état aux points clés.
5. **Corriger à la racine** : pas de rustine qui masque le symptôme.

### Console utile

```jsx
console.log('props reçues :', props);
console.table(tableauDeMesures); // affichage tableau
console.time('calcul'); /* … */ console.timeEnd('calcul');
```

### Erreurs typiques et lecture

| Message | Signification |
|---|---|
| *Objects are not valid as a React child* | On essaie d'afficher un objet `{...}` directement → afficher une propriété ou `JSON.stringify` |
| *Cannot read properties of undefined* | Donnée pas encore chargée → garde `if (!data)` ou `?.` |
| *Too many re-renders* | `setState` appelé **pendant** le rendu (ex. `onClick={setX(1)}`) |
| *Each child in a list should have a unique "key"* | `key` manquante dans un `.map()` |

### Debugger VS Code

Avec l'extension "JavaScript Debugger", on peut poser des **breakpoints** directement dans le `.jsx` (source maps de Vite). Plus précis que les `console.log` en série.

---

## 36. Tests avec Vitest + Testing Library

**Vitest** : runner de tests ultra-rapide, natif pour Vite. **Testing Library** (`@testing-library/react`) : teste les composants **comme un utilisateur** (clics, saisie) plutôt que l'implémentation.

```bash
npm install -D vitest @testing-library/react @testing-library/jest-dom jsdom
```

```js
// vite.config.js — ajouter la section test
export default defineConfig({
  plugins: [react()],
  test: {
    environment: 'jsdom',       // simule le DOM du navigateur
    setupFiles: './src/test-setup.js',
    globals: true,
  },
});
```

```js
// src/test-setup.js
import '@testing-library/jest-dom'; // matchers : toBeInTheDocument(), …
```

```json
// package.json — scripts
{
  "scripts": {
    "test": "vitest",          // mode watch
    "test:run": "vitest run"   // une seule passe (CI)
  }
}
```

**Philosophie Testing Library :** on ne teste pas l'état interne ni les détails d'implémentation ; on teste ce que **voit** et **fait** l'utilisateur : texte affiché, bouton cliquable, formulaire remplissable.

**Les requêtes (par ordre de préférence) :**

| Requête | Usage |
|---|---|
| `getByRole('button', { name: 'Envoyer' })` | ✅ Rôle accessible — préférée |
| `getByLabelText('Titre')` | ✅ Champ de formulaire |
| `getByText('En ligne')` | ✅ Texte visible |
| `getByTestId('…')` | ⚠️ Dernier recours (`data-testid`) |

---

## 37. Tests — exemples commentés

```jsx
// components/Badge.test.jsx
import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import Badge from './Badge.jsx';

describe('Badge', () => {
  it('affiche le texte fourni', () => {
    render(<Badge texte="En ligne" />);
    expect(screen.getByText('En ligne')).toBeInTheDocument();
  });

  it('applique la classe critique quand niveau > 80', () => {
    const { container } = render(<Badge texte="Charge" niveau={95} />);
    expect(container.firstChild).toHaveClass('critique');
  });
});
```

```jsx
// components/Compteur.test.jsx — interaction utilisateur
import { render, screen, fireEvent } from '@testing-library/react';
import Compteur from './Compteur.jsx';

it('incrémente au clic', () => {
  render(<Compteur />);
  const bouton = screen.getByRole('button', { name: '+1' });
  fireEvent.click(bouton);
  expect(screen.getByText('Clics : 1')).toBeInTheDocument();
});
```

**Tester un appel API (mock de fetch) :**

```jsx
import { render, screen, waitFor } from '@testing-library/react';
import { vi } from 'vitest';
import ListeEquipements from './ListeEquipements.jsx';

it('affiche les équipements chargés', async () => {
  global.fetch = vi.fn().mockResolvedValue({
    ok: true,
    json: async () => [{ id: '1', nom: 'UPS-01' }],
  });

  render(<ListeEquipements />);

  // waitFor : attend l'update asynchrone
  await waitFor(() => {
    expect(screen.getByText('UPS-01')).toBeInTheDocument();
  });
  vi.restoreAllMocks();
});
```

**Que tester en priorité :** logique métier (calculs, formatage), formulaires (validation, soumission), états loading/error. **Ne pas tester** : le style CSS, les détails d'implémentation, les librairies tierces.

---

## 38. TypeScript + React

TypeScript ajoute le **typage statique** : l'éditeur détecte les erreurs avant l'exécution. Recommandé pour tout projet d'équipe ou durable.

```bash
npm create vite@latest mon-app -- --template react-ts
```

### Typer les props

```tsx
// CarteOnduleur.tsx
interface CarteOnduleurProps {
  nom: string;
  puissance: number;
  critique?: boolean;          // optionnelle
  onAlerte?: (msg: string) => void;
}

export default function CarteOnduleur({ nom, puissance, critique = false, onAlerte }: CarteOnduleurProps) {
  return (
    <div>
      <h2>{nom}</h2>
      <p>{puissance} kVA</p>
      {critique && <button onClick={() => onAlerte?.('Surchauffe')}>Alerter</button>}
    </div>
  );
}
```

### Typer l'état et les événements

```tsx
const [charge, setCharge] = useState<number>(0);
const [ups, setUps] = useState<Onduleur | null>(null); // peut être null avant chargement

interface Onduleur { id: string; nom: string; puissance: number; charge: number; }

const gererChangement = (e: React.ChangeEvent<HTMLInputElement>) => {
  setNom(e.target.value);
};
const envoyer = (e: React.FormEvent) => {
  e.preventDefault();
  // …
};
```

### Types React utiles

