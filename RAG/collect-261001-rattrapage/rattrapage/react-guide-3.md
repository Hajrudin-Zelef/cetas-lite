---
id: collect-261001-rattrapage/rattrapage/react-guide-3
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [507, 790]
sha256: f5527666b2e06ee7c0eefe86222f8cb9d2fae300b66517599f7dfe88c06fdab2
---

# Guide React — Le manuel complet

```jsx
function Bouton({ variante, children, ...reste }) {
  return <button className={`btn btn-${variante}`} {...reste}>{children}</button>;
}

// Spécialisations par composition, pas par héritage
function BoutonDanger(props) {
  return <Bouton variante="danger" {...props} />;
}
function BoutonPrimaire(props) {
  return <Bouton variante="primaire" {...props} />;
}
```

### Slots multiples (plusieurs `children` nommés)

```jsx
function MiseEnPage({ entete, lateral, children }) {
  return (
    <div className="layout">
      <header>{entete}</header>
      <aside>{lateral}</aside>
      <main>{children}</main>
    </div>
  );
}

<MiseEnPage
  entete={<BarreNavigation />}
  lateral={<MenuSites />}
/>;
```

**Règle :** si vous êtes tenté d'écrire `class X extends Y` pour des composants → utilisez la composition. L'héritage de composants n'existe quasiment jamais en React moderne.

---

## 11. L'état avec useState

Le **state** (état) est la mémoire d'un composant : des données qui, quand elles changent, provoquent un nouveau rendu.

```jsx
import { useState } from 'react';

function Compteur() {
  // [valeur actuelle, fonction pour la modifier] = useState(valeur initiale)
  const [compteur, setCompteur] = useState(0);

  return (
    <div>
      <p>Clics : {compteur}</p>
      <button onClick={() => setCompteur(compteur + 1)}>+1</button>
    </div>
  );
}
```

**Exemple métier — bascule d'un équipement :**

```jsx
function InterrupteurEquipement({ nom }) {
  const [actif, setActif] = useState(true);

  return (
    <div>
      <span>{nom} : {actif ? '🟢 En service' : '🔴 Hors service'}</span>
      <button onClick={() => setActif(!actif)}>
        {actif ? 'Mettre hors service' : 'Remettre en service'}
      </button>
    </div>
  );
}
```

**Points fondamentaux :**
1. `useState` retourne **toujours un tableau de 2 éléments** : `[état, setÉtat]`.
2. Appeler `setÉtat` **ne modifie pas la variable immédiatement** : il planifie un nouveau rendu.
3. Chaque instance de composant a son **propre état isolé** (deux `<Compteur />` ne partagent rien).
4. L'état initial n'est évalué qu'au **premier rendu** (pour un calcul coûteux : `useState(() => calculCouteux())` — initialisation paresseuse).

**Types d'état courants :**

```jsx
const [nom, setNom] = useState('');           // string
const [charge, setCharge] = useState(0);      // number
const [enLigne, setEnLigne] = useState(true); // boolean
const [ups, setUps] = useState(null);         // null en attendant les données
const [mesures, setMesures] = useState([]);   // tableau
const [config, setConfig] = useState({ seuil: 80 }); // objet
```

---

## 12. useState — pièges et bonnes pratiques

### Piège 1 : l'état est immuable — ne jamais le muter

```jsx
const [mesures, setMesures] = useState([72, 75, 71]);

// ❌ INTERDIT : mutation directe → React ne détecte pas le changement
mesures.push(78);
setMesures(mesures);

// ✅ Créer un NOUVEAU tableau
setMesures([...mesures, 78]);

// ❌ INTERDIT : mutation d'objet
config.seuil = 90;
setConfig(config);

// ✅ Nouvel objet
setConfig({ ...config, seuil: 90 });
```

**Tableau récapitulatif des mises à jour immuables :**

| Opération | ❌ Mutant | ✅ Immuable |
|---|---|---|
| Ajouter | `push()` | `[...t, x]` |
| Retirer | `splice()` | `t.filter(...)` |
| Modifier un élément | `t[i] = x` | `t.map((e, i) => ...)` |
| Trier | `sort()` | `[...t].sort()` |
| Changer une clé | `o.k = v` | `{...o, k: v}` |
| Objet imbriqué | `o.a.b = v` | `{...o, a: {...o.a, b: v}}` |

### Piège 2 : le batching — l'état n'est pas mis à jour immédiatement

```jsx
function doubleIncrement() {
  setCompteur(compteur + 1);
  setCompteur(compteur + 1);
  // ❌ compteur n'augmente que de 1 : les deux appels voient l'ANCIENNE valeur
}

// ✅ Forme fonctionnelle : reçoit la valeur à jour
function doubleIncrement() {
  setCompteur((c) => c + 1);
  setCompteur((c) => c + 1); // +2 garanti
}
```

**Règle :** dès que la nouvelle valeur **dépend de l'ancienne**, utiliser la forme `setÉtat((prev) => ...)`.

### Piège 3 : état redondant (dérivable)

```jsx
// ❌ Mauvais : nomComplet est dérivable de prenom + nom
const [prenom, setPrenom] = useState('Jean');
const [nom, setNom] = useState('Dupont');
const [nomComplet, setNomComplet] = useState(''); // risque de désynchro !

// ✅ Calculer pendant le rendu
const nomComplet = `${prenom} ${nom}`;
```

**Ne mettre dans l'état que ce qui ne peut pas être calculé** à partir des props ou d'un autre état.

---

## 13. Rendu conditionnel

Plusieurs techniques pour afficher ou non du contenu :

```jsx
function StatutEquipement({ enLigne, enAlarme }) {
  // 1. Ternaire (if/else)
  const badge = enLigne
    ? <span className="ok">En ligne</span>
    : <span className="ko">Hors ligne</span>;

  return (
    <div>
      {badge}

      {/* 2. && pour "si vrai, afficher" */}
      {enAlarme && <AlerteCritique />}

      {/* 3. Ternaire inline */}
      {enLigne ? <DetailsTechniques /> : <p>Équipement injoignable</p>}
    </div>
  );
}
```

**Attention au piège du `&&` avec `0` :**

```jsx
// ❌ Si nbAlertes vaut 0, React affiche "0" !
{nbAlertes && <p>{nbAlertes} alertes</p>}

// ✅ Comparaison explicite
{nbAlertes > 0 && <p>{nbAlertes} alertes</p>}
```

**Retourner `null` pour ne rien afficher :**

```jsx
function BanniereMaintenance({ maintenanceEnCours }) {
  if (!maintenanceEnCours) return null;
  return <div className="banniere">⚠️ Maintenance planifiée ce soir 22h</div>;
}
```

**Tableau de choix :**

| Besoin | Technique |
|---|---|
| if / else simple | Ternaire `? :` |
| Afficher si vrai | `&&` (avec garde anti-`0`) |
| Logique complexe | `if` avant le `return`, variable intermédiaire |
| Ne rien afficher | `return null` |

---

## 14. Événements en React

Les événements se déclarent en camelCase (`onClick`, `onChange`, `onSubmit`) et reçoivent une **fonction**, pas une string.

```jsx
function PanneauControle() {
  const demarrer = () => console.log('Démarrage…');
  const arreter = (e) => {
    e.preventDefault(); // comme en JS classique
    console.log('Arrêt…');
  };

  return (
    <div>
      {/* Référence de fonction : sans parenthèses */}
      <button onClick={demarrer}>Démarrer</button>

      {/* Arrow inline pour passer un argument */}
      <button onClick={() => arreter('manuel')}>Arrêt manuel</button>

      {/* ❌ NE PAS faire : appelle la fonction pendant le rendu ! */}
      {/* <button onClick={demarrer()}>Démarrer</button> */}
    </div>
  );
}
```

**L'objet événement :** React utilise des événements synthétiques (wrapper cross-browser). Les propriétés utiles : `e.target.value` (inputs), `e.preventDefault()`, `e.stopPropagation()`.

```jsx
function Formulaire() {
  const envoyer = (e) => {
    e.preventDefault(); // empêche le rechargement de la page
    // … logique d'envoi
  };
  return <form onSubmit={envoyer}>…</form>;
}
```

**Événements courants :**

| Événement | Déclencheur |
|---|---|
| `onClick` | Clic |
| `onChange` | Champ modifié (texte, select, checkbox) |
| `onSubmit` | Formulaire soumis |
| `onKeyDown` / `onKeyUp` | Touche clavier |
| `onFocus` / `onBlur` | Focus gagné/perdu |

---

## 15. Formulaires contrôlés

En React, un **formulaire contrôlé** stocke la valeur de chaque champ dans l'état. C'est le pattern standard (vs "non contrôlé" avec `useRef`, réservé à des cas simples).

```jsx
import { useState } from 'react';

function FormulaireTicket() {
  const [titre, setTitre] = useState('');
  const [priorite, setPriorite] = useState('normale');
  const [urgent, setUrgent] = useState(false);

  const envoyer = (e) => {
    e.preventDefault();
    console.log({ titre, priorite, urgent });
    // → POST vers l'API (section 28)
  };

