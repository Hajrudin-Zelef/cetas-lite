---
id: collect-261001-rattrapage/rattrapage/react-guide-4
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [791, 1065]
sha256: 9b1787d96046e6989f860aeb826cdbcb525c666300a7ec4e3691390076ce287f
---

# Guide React — Le manuel complet

  return (
    <form onSubmit={envoyer}>
      <label>
        Titre :
        <input
          type="text"
          value={titre}
          onChange={(e) => setTitre(e.target.value)}
        />
      </label>

      <label>
        Priorité :
        <select value={priorite} onChange={(e) => setPriorite(e.target.value)}>
          <option value="basse">Basse</option>
          <option value="normale">Normale</option>
          <option value="haute">Haute</option>
        </select>
      </label>

      <label>
        <input
          type="checkbox"
          checked={urgent}
          onChange={(e) => setUrgent(e.target.checked)}
        />
        Urgent
      </label>

      <button type="submit">Créer le ticket</button>
    </form>
  );
}
```

**Points clés :**
- `value` + `onChange` pour texte/select ; `checked` + `onChange` pour checkbox (`e.target.checked`).
- Sans `onChange`, un champ avec `value` est **en lecture seule** (warning React).
- Un seul objet d'état pour un gros formulaire (voir cas pratique 2, section 48).

**Validation simple :**

```jsx
const [erreurs, setErreurs] = useState({});

const valider = () => {
  const e = {};
  if (titre.trim().length < 5) e.titre = 'Titre trop court (min 5 caractères)';
  setErreurs(e);
  return Object.keys(e).length === 0;
};
```

---

## 16. Listes et keys

Pour afficher un tableau, on le transforme avec `.map()` en éléments JSX. Chaque élément **doit** avoir une prop `key` unique et stable.

```jsx
const onduleurs = [
  { id: 'ups-01', nom: 'UPS-01', puissance: 40 },
  { id: 'ups-02', nom: 'UPS-02', puissance: 60 },
];

function ListeOnduleurs() {
  return (
    <ul>
      {onduleurs.map((ups) => (
        <li key={ups.id}>
          {ups.nom} — {ups.puissance} kVA
        </li>
      ))}
    </ul>
  );
}
```

### Pourquoi `key` est critique

React utilise `key` pour **identifier chaque élément entre deux rendus** : réordonner, ajouter, supprimer sans tout re-créer (et sans perdre l'état local des lignes, ex. un champ en cours de saisie).

```jsx
// ❌ JAMAIS l'index comme key si la liste peut changer d'ordre / filtrée
{items.map((item, i) => <Ligne key={i} {...item} />)}

// ✅ Un identifiant stable issu des données
{items.map((item) => <Ligne key={item.id} {...item} />)}
```

**Conséquence d'une mauvaise key :** avec l'index, si on supprime la première ligne, React croit que c'est la *dernière* qui a disparu → états des champs mélangés, bugs "fantômes".

### Filtrage + tri avant le rendu

```jsx
function ListeFiltree({ equipements, recherche }) {
  const visibles = equipements
    .filter((e) => e.nom.toLowerCase().includes(recherche.toLowerCase()))
    .sort((a, b) => a.nom.localeCompare(b.nom));

  if (visibles.length === 0) return <p>Aucun équipement trouvé.</p>;

  return (
    <ul>
      {visibles.map((e) => (
        <li key={e.id}>{e.nom}</li>
      ))}
    </ul>
  );
}
```

> `key` n'est **pas** transmise au composant comme prop. Besoin de l'id dans l'enfant ? Le passer explicitement : `<Ligne key={e.id} id={e.id} />`.

---

## 17. useEffect — dépendances et cleanup

`useEffect` exécute du code **après** le rendu, pour tout ce qui est un **effet de bord** : appels API, timers, abonnements (WebSocket), manipulation du DOM, `document.title`.

```jsx
import { useState, useEffect } from 'react';

function Horloge() {
  const [heure, setHeure] = useState(new Date());

  useEffect(() => {
    // 1. L'EFFET : s'exécute après chaque rendu (selon dépendances)
    const timer = setInterval(() => setHeure(new Date()), 1000);

    // 2. Le CLEANUP : exécuté avant le prochain effet + au démontage
    return () => clearInterval(timer);
  }, []); // 3. Le TABLEAU DE DÉPENDANCES

  return <p>Il est {heure.toLocaleTimeString('fr-FR')}</p>;
}
```

### Le tableau de dépendances — les 3 cas

| Dépendances | Comportement |
|---|---|
| `[]` (vide) | L'effet tourne **une seule fois** (montage) |
| `[a, b]` | L'effet tourne au montage + **quand `a` ou `b` change** |
| *(absent)* | L'effet tourne **après chaque rendu** (rarement voulu) |

```jsx
// ❌ Sans tableau : boucle infinie potentielle (setState → rendu → effet → setState…)
useEffect(() => { setCompteur(compteur + 1); });

// ✅ Avec dépendance : se relance quand l'id change (ex. page détail d'un équipement)
useEffect(() => {
  fetch(`/api/ups/${id}`).then(/* ... */);
}, [id]);
```

### Le cleanup — quand et pourquoi

Le cleanup s'exécute : **avant** chaque ré-exécution de l'effet, et au **démontage** du composant.

```jsx
useEffect(() => {
  const ws = new WebSocket('wss://supervision.local/mesures');
  ws.onmessage = (e) => setMesures(JSON.parse(e.data));
  return () => ws.close(); // ferme l'ancienne connexion avant d'en ouvrir une nouvelle
}, [siteId]);
```

**Sans cleanup :** timers qui continuent après navigation, WebSockets orphelines, `setState` sur composant démonté (warning + fuite mémoire).

> **React 19 :** `useEffect` existe toujours. La nouveauté est le hook `use()` pour consommer des promesses pendant le rendu (Suspense), mais `useEffect` reste le standard pour les effets impératifs.

---

## 18. useEffect — les patterns courants

### Pattern 1 : chargement de données au montage

```jsx
useEffect(() => {
  let ignore = false; // garde anti "setState après démontage"

  async function charger() {
    const res = await fetch('/api/onduleurs');
    const data = await res.json();
    if (!ignore) setOnduleurs(data);
  }
  charger();

  return () => { ignore = true; };
}, []);
```

> Note : on ne peut pas rendre la fonction d'effet `async` directement (elle doit retourner soit rien, soit une fonction de cleanup). On définit une fonction async **à l'intérieur**.

### Pattern 2 : synchroniser le titre du document

```jsx
useEffect(() => {
  document.title = `(${nbAlertes}) Supervision`;
}, [nbAlertes]);
```

### Pattern 3 : s'abonner / se désabonner

```jsx
useEffect(() => {
  const onResize = () => setLargeur(window.innerWidth);
  window.addEventListener('resize', onResize);
  return () => window.removeEventListener('resize', onResize);
}, []);
```

### Pattern 4 : effet conditionné par une valeur

```jsx
// Recharge les mesures quand l'équipement sélectionné change
useEffect(() => {
  if (!equipementId) return;
  chargerMesures(equipementId);
}, [equipementId]);
```

### Anti-patterns useEffect

| ❌ À éviter | ✅ À la place |
|---|---|
| `useEffect` pour calculer une valeur dérivée | Calcul direct pendant le rendu (section 12) |
| Chaîne d'effets qui se déclenchent entre eux | Un seul effet, ou un gestionnaire d'événement |
| `fetch` sans garde `ignore` / AbortController | Annuler proprement (voir section 28) |
| Dépendances manquantes (warning ESLint) | Ajouter toutes les valeurs utilisées, ou `useCallback` |

**Checklist useEffect :**
- [ ] Tableau de dépendances toujours renseigné (même `[]`)
- [ ] Cleanup pour tout timer / abonnement / connexion
- [ ] Pas de fonction `async` directement en callback d'effet
- [ ] Warning `react-hooks/exhaustive-deps` d'ESLint = à corriger, pas à ignorer

---

## 19. useRef — au-delà du DOM

`useRef` retourne un objet `{ current: ... }` **persistant entre les rendus**, dont la modification **ne provoque pas** de nouveau rendu. Deux usages :

### Usage 1 : accéder à un élément du DOM

```jsx
import { useRef } from 'react';

function ChampRecherche() {
  const inputRef = useRef(null);

  const focus = () => inputRef.current?.focus();

  return (
    <div>
      <input ref={inputRef} type="search" placeholder="Rechercher…" />
      <button onClick={focus}>Focus</button>
    </div>
  );
}
```

### Usage 2 : une "boîte" mutable qui survit aux rendus

```jsx
function CompteurRendus() {
  const nbRendus = useRef(0);
  const timerRef = useRef(null);

  useEffect(() => {
    nbRendus.current += 1; // pas de re-rendu → pas de boucle infinie
    console.log('Rendus :', nbRendus.current);
  });

