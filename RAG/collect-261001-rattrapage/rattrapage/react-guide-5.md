---
id: collect-261001-rattrapage/rattrapage/react-guide-5
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [1066, 1298]
sha256: 1504ec4d37a38a4b53219c2b3c0e93850726685763bb2713e8a8ca801588cbcf
---

# Guide React — Le manuel complet

  const demarrer = () => {
    timerRef.current = setInterval(() => console.log('tick'), 1000);
  };
  const arreter = () => clearInterval(timerRef.current);
  // …
}
```

**useRef vs useState :**

|  | `useState` | `useRef` |
|---|---|---|
| Modif → re-rendu | ✅ Oui | ❌ Non |
| Lecture pendant le rendu | ✅ | ⚠️ Éviter (valeur "en retard" possible) |
| Cas typique | Données affichées | Timers, éléments DOM, valeur précédente |

**Piège :** ne pas lire/écrire `ref.current` **pendant** le rendu (sauf initialisation) — le faire dans des effets ou des gestionnaires d'événements.

---

## 20. useMemo — mémoriser un calcul

`useMemo` mémorise le **résultat** d'un calcul coûteux et ne le recalcule que si ses dépendances changent.

```jsx
import { useMemo } from 'react';

function Statistiques({ mesures }) {
  // Sans useMemo : ce tri+calcul tourne à CHAQUE rendu (même si mesures inchangé)
  const stats = useMemo(() => {
    console.log('Calcul des stats…');
    const triees = [...mesures].sort((a, b) => a - b);
    return {
      min: triees[0],
      max: triees[triees.length - 1],
      moyenne: triees.reduce((s, v) => s + v, 0) / triees.length,
    };
  }, [mesures]);

  return <p>Min {stats.min} — Max {stats.max} — Moy {stats.moyenne.toFixed(1)}</p>;
}
```

**Quand l'utiliser :**
- Calcul réellement coûteux (tri/filtrage de milliers de lignes, agrégations).
- Valeur passée à un composant mémorisé (`React.memo`) ou à un `useEffect` en dépendance (référence stable).

**Quand NE PAS l'utiliser :** pour un calcul trivial — `useMemo` a lui-même un coût. **Ne pas optimiser prématurément** : d'abord mesurer (React DevTools Profiler, section 34).

---

## 21. useCallback — mémoriser une fonction

`useCallback` mémorise une **fonction** : même référence entre les rendus tant que les dépendances ne changent pas.

```jsx
import { useState, useCallback } from 'react';

function Liste({ items }) {
  const [selection, setSelection] = useState(new Set());

  // Sans useCallback : nouvelle fonction à chaque rendu → LigneItem (memo) re-rend inutilement
  const basculer = useCallback((id) => {
    setSelection((prev) => {
      const next = new Set(prev);
      next.has(id) ? next.delete(id) : next.add(id);
      return next;
    });
  }, []); // setSelection est stable → pas de dépendance

  return items.map((item) => (
    <LigneItem key={item.id} item={item} onBasculer={basculer} />
  ));
}

const LigneItem = React.memo(function LigneItem({ item, onBasculer }) {
  return <li onClick={() => onBasculer(item.id)}>{item.nom}</li>;
});
```

**`useCallback(fn, deps)` ≈ `useMemo(() => fn, deps)`** — c'est du sucre syntaxique.

**Règle pratique :** `useCallback` n'est utile que si la fonction est passée à un composant **mémorisé** ou utilisée comme **dépendance d'un effet**. Sinon, une fonction inline suffit.

---

## 22. React.memo — éviter les re-rendus inutiles

Par défaut, quand un parent se re-rend, **tous ses enfants se re-rendent aussi** (même si leurs props n'ont pas changé). `React.memo` mémorise le rendu d'un composant : il ne se re-rend que si ses props changent (comparaison superficielle).

```jsx
import { memo } from 'react';

// Ne se re-rend que si `ups` (référence) change
const CarteOnduleur = memo(function CarteOnduleur({ ups }) {
  console.log('Rendu', ups.nom);
  return <div>{ups.nom} — {ups.charge} %</div>;
});
```

**Comparaison personnalisée** (rarement nécessaire) :

```jsx
const Ligne = memo(function Ligne({ data }) { /* ... */ },
  (prev, next) => prev.data.id === next.data.id // true = skip le re-rendu
);
```

**Limites importantes :**
- Comparaison **superficielle** : si le parent recrée un objet à chaque rendu (`data={{...}}`), `memo` ne sert à rien → combiner avec `useMemo`/`useCallback`.
- Ne pas enrober **tous** les composants : la comparaison a un coût. Cibler les composants coûteux ou très souvent re-rendus (lignes de tableau, listes longues).
- `memo` ne protège pas contre un changement de `children` recréé à chaque rendu.

**Ordre d'optimisation :** 1) structure correcte (état au bon niveau), 2) `memo` ciblé, 3) `useMemo`/`useCallback` en support.

---

## 23. useContext et la Context API

Le **contexte** permet de partager des données à toute une sous-arborescence **sans prop drilling** (passer des props à travers 5 niveaux de composants qui n'en ont pas besoin).

Cas typiques : thème, utilisateur connecté, langue, configuration.

```jsx
import { createContext, useContext, useState } from 'react';

// 1. Créer le contexte (avec une valeur par défaut)
const ThemeContext = createContext('clair');

// 2. Fournir la valeur en haut de l'arbre
function App() {
  const [theme, setTheme] = useState('clair');
  return (
    <ThemeContext.Provider value={theme}>
      <Entete />
      <button onClick={() => setTheme(theme === 'clair' ? 'sombre' : 'clair')}>
        Basculer le thème
      </button>
    </ThemeContext.Provider>
  );
}

// 3. Consommer où on veut, sans props intermédiaires
function Entete() {
  const theme = useContext(ThemeContext);
  return <header className={`theme-${theme}`}>Supervision</header>;
}
```

### Pattern "contexte + hook personnalisé" (recommandé)

```jsx
// contexts/AuthContext.jsx
import { createContext, useContext, useState } from 'react';

const AuthContext = createContext(null);

export function AuthProvider({ children }) {
  const [utilisateur, setUtilisateur] = useState(null);
  const connecter = (u) => setUtilisateur(u);
  const deconnecter = () => setUtilisateur(null);

  return (
    <AuthContext.Provider value={{ utilisateur, connecter, deconnecter }}>
      {children}
    </AuthContext.Provider>
  );
}

// Hook qui masque useContext + vérifie le provider
export function useAuth() {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error('useAuth doit être utilisé dans <AuthProvider>');
  return ctx;
}

// Utilisation
function Profil() {
  const { utilisateur, deconnecter } = useAuth();
  // …
}
```

**Avertissements :**
- Tout changement de `value` re-rend **tous** les consommateurs → mémoriser la valeur (`useMemo`) si elle contient des objets recréés à chaque rendu.
- Ne pas mettre dans un contexte des données qui changent très souvent (ex. position souris) — préférer un état local ou une librairie dédiée.
- Le contexte ne remplace pas un vrai store (Zustand, Redux) pour un état global complexe.

---

## 24. useReducer — l'état complexe

`useReducer` est une alternative à `useState` quand l'état a **plusieurs sous-valeurs liées** ou que les transitions sont complexes. Inspiré de Redux : un **reducer** pur `(état, action) => nouvelÉtat`.

```jsx
import { useReducer } from 'react';

const initial = { charge: 72, enLigne: true, alertes: [] };

function reducer(etat, action) {
  switch (action.type) {
    case 'MESURE':
      return { ...etat, charge: action.valeur };
    case 'PANNE':
      return { ...etat, enLigne: false, alertes: [...etat.alertes, action.alerte] };
    case 'REPARER':
      return { ...etat, enLigne: true };
    case 'RESET':
      return initial;
    default:
      throw new Error(`Action inconnue : ${action.type}`);
  }
}

function SupervisionUps() {
  const [etat, dispatch] = useReducer(reducer, initial);

  return (
    <div>
      <p>Charge : {etat.charge} % — {etat.enLigne ? 'En ligne' : 'En panne'}</p>
      <button onClick={() => dispatch({ type: 'MESURE', valeur: 75 })}>
        Simuler mesure
      </button>
      <button onClick={() => dispatch({ type: 'PANNE', alerte: 'Surchauffe' })}>
        Simuler panne
      </button>
    </div>
  );
}
```

**useState vs useReducer :**

