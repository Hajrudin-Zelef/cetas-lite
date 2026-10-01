---
id: collect-261001-rattrapage/rattrapage/react-guide-2
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["containment"]
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [223, 506]
sha256: 7253c8a4ba1c16afb8a7cca84f448129ca29c9b2360651ffe62e5ebd7dc69d9a
---

# Guide React — Le manuel complet

**Points clés :**
- `StrictMode` : en dev, il double l'invocation de certains rendus pour révéler les effets de bord. **Ne pas le retirer** : un bug visible uniquement sans StrictMode est un bug réel.
- `createRoot` : l'API React 18+ (l'ancien `ReactDOM.render` est déprécié).
- Le fichier `index.html` contient `<div id="root"></div>` et le script `/src/main.jsx`.

**Convention de nommage :** fichiers de composants en `PascalCase` (`CarteOnduleur.jsx`), un composant par fichier, extension `.jsx` (ou `.tsx`).

---

## 5. JSX — la syntaxe au cœur de React

JSX est une extension de syntaxe qui permet d'écrire du HTML dans du JavaScript :

```jsx
const titre = <h1>État des onduleurs</h1>;
```

Ce n'est ni une string ni du HTML : c'est du **sucre syntaxique** pour `React.createElement('h1', null, 'État des onduleurs')`. Le navigateur ne comprend pas JSX — Vite le transpile (via esbuild) en JavaScript pur au build.

```jsx
// Ce que vous écrivez
const carte = <div className="carte">UPS-01</div>;

// Ce que Vite produit (simplifié)
const carte = React.createElement('div', { className: 'carte' }, 'UPS-01');
```

**Conséquences pratiques :**
1. On peut mélanger logique JS et markup dans le même fichier — c'est voulu, c'est la **colocalisation**.
2. `class` devient `className` (mot réservé JS), `for` devient `htmlFor`.
3. Tout attribut HTML standard existe en camelCase : `tabIndex`, `onClick`, `readOnly`, `maxLength`.

**Pense-bête JSX :**
- [ ] `class` → `className`
- [ ] Attributs en camelCase
- [ ] Les balises doivent être fermées : `<img />`, `<br />`, `<input />`
- [ ] Commentaires : `{/* commentaire */}`, pas `<!-- -->`

---

## 6. JSX — expressions, attributs, enfants

Les accolades `{}` permettent d'injecter **n'importe quelle expression JavaScript** dans le JSX :

```jsx
const ups = { nom: 'UPS-01', puissance: 40, enLigne: true };

function CarteUps() {
  return (
    <div className="carte">
      <h2>{ups.nom}</h2>
      <p>Puissance : {ups.puissance} kVA</p>
      <p>Double : {ups.puissance * 2} kVA</p>
      <p>Statut : {ups.enLigne ? 'En ligne' : 'Hors ligne'}</p>
      <p>Prochaine maintenance : {prochaineDate().toLocaleDateString('fr-FR')}</p>
    </div>
  );
}
```

**Ce qui est autorisé dans `{}` :** variables, appels de fonctions, ternaires, calculs, accès à des tableaux/objets.

**Ce qui est INTERDIT :** les instructions (`if`, `for`, `while`, `return` au milieu). Pour de la logique complexe, la calculer **avant** le `return` :

```jsx
function Badge({ niveau }) {
  // Logique AVANT le return
  const couleur = niveau > 80 ? 'rouge' : niveau > 50 ? 'orange' : 'vert';
  const libelle = `Charge ${niveau} %`;

  return <span style={{ color: couleur }}>{libelle}</span>;
}
```

**Styles inline :** un objet JS, clés en camelCase :

```jsx
<div style={{ backgroundColor: '#111', padding: '16px', fontSize: 14 }}>
```

> Note : `fontSize: 14` → React ajoute automatiquement `px` pour les propriétés numériques qui l'acceptent.

---

## 7. JSX — les 3 règles de base

### Règle 1 : un seul élément parent

Un composant doit retourner **un seul** élément racine. Sinon, erreur de syntaxe.

```jsx
// ❌ ERREUR : deux racines
return (
  <h1>Titre</h1>
  <p>Texte</p>
);

// ✅ OK : envelopper dans une div…
return (
  <div>
    <h1>Titre</h1>
    <p>Texte</p>
  </div>
);

// ✅ MIEUX : Fragment (n'ajoute pas de div inutile dans le DOM)
return (
  <>
    <h1>Titre</h1>
    <p>Texte</p>
  </>
);
```

`<>...</>` est le raccourci de `<React.Fragment>...</React.Fragment>`.

### Règle 2 : fermer toutes les balises

```jsx
<img src="ups.png" alt="Onduleur" />   // ✅
<br />                                 // ✅
<input type="text" />                  // ✅
```

### Règle 3 : camelCase pour les attributs

`class` → `className`, `onclick` → `onClick`, `tabindex` → `tabIndex`.

---

## 8. Composants fonctionnels

Un composant est une **fonction JavaScript qui retourne du JSX**. Son nom commence par une majuscule (sinon React le traite comme une balise HTML).

```jsx
// Composant simple
function Titre() {
  return <h1>Supervision énergétique</h1>;
}

// Utilisation (comme une balise HTML personnalisée)
function App() {
  return (
    <div>
      <Titre />
      <Titre />
    </div>
  );
}
```

**Anatomie complète d'un composant :**

```jsx
import { useState } from 'react';

export default function CarteOnduleur({ nom, puissance }) {
  const [allume, setAllume] = useState(true);

  return (
    <article className="carte">
      <h2>{nom}</h2>
      <p>{puissance} kVA — {allume ? 'ON' : 'OFF'}</p>
      <button onClick={() => setAllume(!allume)}>Basculer</button>
    </article>
  );
}
```

**Règles d'un composant :**
- Nom en `PascalCase`.
- Retourne du JSX (ou `null` pour ne rien afficher).
- Ne modifie jamais ses props (elles sont en lecture seule).
- Peut contenir de l'état (hooks), des effets, des fonctions internes.

**Composition :** les composants s'imbriquent comme des briques :

```jsx
function Dashboard() {
  return (
    <main>
      <Entete titre="Supervision" />
      <Grille>
        <CarteOnduleur nom="UPS-01" puissance={40} />
        <CarteOnduleur nom="UPS-02" puissance={60} />
      </Grille>
      <PiedDePage />
    </main>
  );
}
```

> **Héritage (classes) :** l'ancien modèle à base de `class MonComposant extends React.Component` existe encore dans du code ancien. En React 18+, on écrit **100 % fonctionnel**. Les classes ne sont détaillées ici que pour lire du code legacy.

---

## 9. Props en détail

Les **props** (propriétés) sont les arguments d'un composant : le mécanisme pour faire passer des données du parent vers l'enfant. Elles sont **en lecture seule** dans l'enfant.

```jsx
// Parent : passe les props comme des attributs HTML
<CarteOnduleur nom="UPS-01" puissance={40} critique={true} />

// Enfant : les reçoit dans un objet en premier paramètre
function CarteOnduleur(props) {
  return <h2>{props.nom} — {props.puissance} kVA</h2>;
}

// Version idiomatique : déstructuration directe
function CarteOnduleur({ nom, puissance, critique = false }) {
  return (
    <h2 className={critique ? 'critique' : ''}>
      {nom} — {puissance} kVA
    </h2>
  );
}
```

**Types de props courants :**

| Type | Exemple | Usage |
|---|---|---|
| String | `nom="UPS-01"` | Texte simple |
| Nombre | `puissance={40}` | Accolades obligatoires |
| Booléen | `critique` (= `critique={true}`) | Flag |
| Objet / tableau | `mesures={[1, 2, 3]}` | Données |
| Fonction | `onAlerte={gererAlerte}` | Callback enfant → parent |
| Élément | `icone={<IconeUps />}` | Composition |

**Props spéciales :**
- `children` : le contenu entre les balises ouvrante/fermante.
- `key` : réservée à React pour les listes (section 16), **non accessible** dans le composant.

```jsx
function Cadre({ titre, children }) {
  return (
    <section className="cadre">
      <h2>{titre}</h2>
      <div className="contenu">{children}</div>
    </section>
  );
}

// Utilisation
<Cadre titre="Batteries">
  <p>Tension : 432 V</p>
  <p>Autonomie : 12 min</p>
</Cadre>;
```

**Valeurs par défaut** (déstructuration, syntaxe moderne) :

```jsx
function Badge({ texte, couleur = 'gris', taille = 'm' }) { /* ... */ }
```

> L'ancienne syntaxe `Composant.defaultProps = {...}` est dépréciée — utiliser les valeurs par défaut de la déstructuration.

---

## 10. Composition vs héritage

React **n'utilise pas l'héritage** entre composants : on compose. C'est un choix de design assumé.

### Containment (conteneur générique)

```jsx
function Panneau({ titre, children }) {
  return (
    <div className="panneau">
      <header>{titre}</header>
      {children}
    </div>
  );
}

<Panneau titre="Alertes">
  <ListeAlertes />
</Panneau>;
```

### Spécialisation (variante d'un composant générique)

