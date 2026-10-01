---
id: collect-261001-rattrapage/rattrapage/react-guide-6
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [1299, 1552]
sha256: 8dd6ca4cca335945892e82658ed4f8d5a6a19d2aa339f0c91a0a020a9af25b26
---

# Guide React — Le manuel complet

|  | `useState` | `useReducer` |
|---|---|---|
| État simple (1-2 valeurs) | ✅ Idéal | Overkill |
| Plusieurs valeurs liées | Devient verbeux | ✅ Idéal |
| Transitions nommées/testables | ❌ | ✅ Le reducer est une fonction pure |
| Logique réutilisable | ❌ | ✅ Reducer exportable |

**Bonnes pratiques :** actions avec `type` en SNAKE_CASE, payloads explicites, reducer **pur** (pas de mutation, pas d'appels API dedans).

---

## 25. Hooks personnalisés

Un **hook personnalisé** est une fonction qui encapsule de la logique à base de hooks pour la **réutiliser** entre composants. Convention : nom préfixé par `use`.

```jsx
// hooks/useCompteur.js
import { useState, useCallback } from 'react';

export function useCompteur(initial = 0) {
  const [valeur, setValeur] = useState(initial);
  const incrementer = useCallback(() => setValeur((v) => v + 1), []);
  const decrementer = useCallback(() => setValeur((v) => v - 1), []);
  const reset = useCallback(() => setValeur(initial), [initial]);
  return { valeur, incrementer, decrementer, reset };
}

// Utilisation : chaque composant a son état indépendant
function PanneauA() {
  const { valeur, incrementer } = useCompteur(0);
  return <button onClick={incrementer}>A : {valeur}</button>;
}
```

**Exemple métier — polling d'une API :**

```jsx
// hooks/usePolling.js
import { useState, useEffect, useRef } from 'react';

export function usePolling(url, intervalleMs = 5000) {
  const [data, setData] = useState(null);
  const [erreur, setErreur] = useState(null);
  const [chargement, setChargement] = useState(true);

  useEffect(() => {
    let ignore = false;
    let timer;

    async function tick() {
      try {
        const res = await fetch(url);
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        const json = await res.json();
        if (!ignore) { setData(json); setErreur(null); }
      } catch (e) {
        if (!ignore) setErreur(e.message);
      } finally {
        if (!ignore) setChargement(false);
      }
    }

    tick();
    timer = setInterval(tick, intervalleMs);
    return () => { ignore = true; clearInterval(timer); };
  }, [url, intervalleMs]);

  return { data, erreur, chargement };
}

// Utilisation : 3 lignes pour un polling complet
function WidgetUps() {
  const { data, erreur, chargement } = usePolling('/api/ups/ups-01', 5000);
  if (chargement) return <p>Chargement…</p>;
  if (erreur) return <p>Erreur : {erreur}</p>;
  return <p>Charge : {data.charge} %</p>;
}
```

**Checklist hook personnalisé :**
- [ ] Nom en `useQuelqueChose`
- [ ] Retourne ce que le composant consomme (objet ou tableau)
- [ ] Chaque appel = état indépendant (pas de partage magique)
- [ ] Testable : logique pure extraite si possible

---

## 26. Les règles des hooks (linter)

Trois règles **non négociables**, vérifiées par le plugin ESLint `react-hooks` (inclus dans les templates Vite) :

### Règle 1 : appeler les hooks uniquement au niveau racine

```jsx
// ❌ INTERDIT : hook dans une condition / boucle / fonction imbriquée
function Mauvais({ actif }) {
  if (actif) {
    const [x, setX] = useState(0); // ordre des hooks instable !
  }
}

// ✅ Toujours au top-level, dans le même ordre à chaque rendu
function Bon({ actif }) {
  const [x, setX] = useState(0);
  useEffect(() => {
    if (actif) { /* ... */ } // la condition est DANS le hook
  }, [actif]);
}
```

**Pourquoi ?** React identifie les hooks par leur **ordre d'appel**. Un hook conditionnel décale tout l'ordre → états mélangés.

### Règle 2 : appeler les hooks uniquement dans des composants React ou des hooks personnalisés

Pas dans des fonctions utilitaires classiques, pas dans des classes.

### Règle 3 : respecter les dépendances (`exhaustive-deps`)

Le warning ESLint *"React Hook useEffect has a missing dependency"* signale une vraie source de bugs (valeur obsolète capturée). Le corriger, pas le masquer avec `// eslint-disable-next-line`.

**Pour aller plus loin :** le compilateur React (React 19, "React Compiler") mémorise automatiquement — mais les règles des hooks restent valables.

---

## 27. Data fetching — patterns loading / error

Le trio d'états standard pour tout appel réseau : **`data` / `erreur` / `chargement`**.

```jsx
import { useState, useEffect } from 'react';

function ListeEquipements() {
  const [data, setData] = useState(null);
  const [chargement, setChargement] = useState(true);
  const [erreur, setErreur] = useState(null);

  useEffect(() => {
    const controleur = new AbortController(); // pour annuler la requête

    async function charger() {
      try {
        setChargement(true);
        const res = await fetch('/api/equipements', { signal: controleur.signal });
        if (!res.ok) throw new Error(`Erreur HTTP ${res.status}`);
        setData(await res.json());
      } catch (e) {
        if (e.name !== 'AbortError') setErreur(e.message);
      } finally {
        setChargement(false);
      }
    }
    charger();

    return () => controleur.abort(); // annule si démontage / changement
  }, []);

  // Rendu selon l'état — TOUJOURS dans cet ordre
  if (chargement) return <p>⏳ Chargement des équipements…</p>;
  if (erreur) return <p className="erreur">❌ {erreur}</p>;
  if (!data || data.length === 0) return <p>Aucun équipement.</p>;

  return (
    <ul>
      {data.map((e) => <li key={e.id}>{e.nom}</li>)}
    </ul>
  );
}
```

**Les 4 états d'une requête — checklist :**
- [ ] `chargement` : skeleton ou spinner (pas de page blanche)
- [ ] `erreur` : message compréhensible + bouton "Réessayer"
- [ ] `vide` : "Aucun résultat" plutôt qu'une liste vide silencieuse
- [ ] `succès` : les données

**`AbortController` :** annule la requête si le composant se démonte ou si les dépendances changent → évite les `setState` sur composant démonté et les réponses "en retard" qui écrasent les fraîches (race condition).

---

## 28. Data fetching — exemple complet

Un composant de détail qui recharge quand l'`id` change, avec annulation et retry :

```jsx
import { useState, useEffect, useCallback } from 'react';

function DetailOnduleur({ id }) {
  const [ups, setUps] = useState(null);
  const [chargement, setChargement] = useState(true);
  const [erreur, setErreur] = useState(null);

  const charger = useCallback(async (signal) => {
    setChargement(true);
    setErreur(null);
    try {
      const res = await fetch(`/api/onduleurs/${id}`, { signal });
      if (!res.ok) throw new Error(`HTTP ${res.status} — vérifiez l'ID`);
      setUps(await res.json());
    } catch (e) {
      if (e.name !== 'AbortError') setErreur(e.message);
    } finally {
      setChargement(false);
    }
  }, [id]);

  useEffect(() => {
    const controleur = new AbortController();
    charger(controleur.signal);
    return () => controleur.abort();
  }, [charger]);

  if (chargement) return <p>⏳ Chargement de {id}…</p>;
  if (erreur) return (
    <div className="erreur">
      <p>❌ {erreur}</p>
      <button onClick={() => charger()}>Réessayer</button>
    </div>
  );
  if (!ups) return null;

  return (
    <article>
      <h2>{ups.nom}</h2>
      <dl>
        <dt>Puissance</dt><dd>{ups.puissance} kVA</dd>
        <dt>Charge</dt><dd>{ups.charge} %</dd>
        <dt>Batterie</dt><dd>{ups.batterieV} V</dd>
      </dl>
    </article>
  );
}
```

**Envoi de données (POST) :**

```jsx
async function creerTicket(ticket) {
  const res = await fetch('/api/tickets', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(ticket),
  });
  if (!res.ok) {
    const detail = await res.json().catch(() => ({}));
    throw new Error(detail.message || `HTTP ${res.status}`);
  }
  return res.json();
}
```

**Alternatives à `fetch` natif :** pour du cache, retry, déduplication → **TanStack Query** (ex React Query) ou **SWR**. Pour un dashboard avec polling, TanStack Query change la vie (voir section 55).

---

