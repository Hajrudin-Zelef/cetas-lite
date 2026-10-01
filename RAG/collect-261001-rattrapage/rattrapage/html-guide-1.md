---
id: collect-261001-rattrapage/rattrapage/html-guide-1
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [1, 213]
sha256: ba54cfaac0469f15483c5de9fedc441c8072fcb5cf1389070c68fc0b59ec71de
---

# Guide HTML & CSS — Le manuel complet

> **Pour Zelef** — chef de service systèmes & énergies : pages internes, dashboards, documentation.
> Ton : direct, tutoriel + référence, dense. Zéro blabla, que de l'utile.
> **Convention de ce guide** : chaque section numérotée `## N.` est autonome. Les blocs de code sont valides HTML5/CSS3 et copiables tels quels.

---

## 1. Bienvenue : comment utiliser ce guide

Ce guide couvre **HTML5** de A à Z + une **introduction solide à CSS** (sélecteurs, box model, flexbox, grid, responsive, variables). Il est pensé pour un usage pro interne :

- **Sections 1–10** : prise en main, tooling, syntaxe de base.
- **Sections 11–35** : HTML en profondeur (texte, liens, images, tableaux, formulaires, médias, sémantique).
- **Sections 36–49** : SEO, accessibilité, ARIA.
- **Sections 50–75** : CSS complet pour mettre en page (flexbox, grid, responsive).
- **Sections 76–88** : bonnes pratiques, erreurs classiques, cas pratiques commentés, pense-bête, glossaire, quiz.

**Méthode de travail conseillée :**

1. Lis les sections 1 à 10 d'une traite (30 min).
2. Fais les cas pratiques (sections 80–83) en tapant le code toi-même.
3. Reviens aux sections de référence (tableaux d'attributs, pense-bête §85) quand tu codes.

> 💡 Tout le code de ce guide est en HTML5 sémantique moderne. Aucune balise obsolète (`<font>`, `<center>`, `<table>` de mise en page…).

---

## 2. C'est quoi, HTML ?

**HTML** (HyperText Markup Language) est le langage de **balisage** qui décrit la **structure** d'une page web. Il ne fait pas la mise en forme (c'est CSS) ni l'interactivité (c'est JavaScript). Retiens le trio :

| Couche      | Rôle                | Langage    |
|-------------|---------------------|------------|
| Structure   | Quoi afficher       | HTML       |
| Présentation| Comment l'afficher  | CSS        |
| Comportement| Ce que ça fait      | JavaScript |

Exemple minimal :

```html
<!DOCTYPE html>
<html lang="fr">
<head>
  <meta charset="utf-8">
  <title>Ma première page</title>
</head>
<body>
  <h1>Bonjour</h1>
  <p>Ceci est un paragraphe.</p>
</body>
</html>
```

Points clés :

- HTML est un langage de **balises** : `<h1>…</h1>` encadre du contenu.
- Les navigateurs **interprètent** le HTML et construisent le **DOM** (Document Object Model), un arbre d'objets manipulable.
- HTML est **tolérant** : un document mal formé s'affiche quand même — mais un code propre évite 90 % des bugs.

---

## 3. Histoire rapide : pourquoi HTML5

| Année | Version | Ce qui change |
|-------|---------|---------------|
| 1991  | HTML (Tim Berners-Lee) | 18 balises, pages scientifiques |
| 1997  | HTML 4.01 | Standard du web « classique » |
| 2000  | XHTML | HTML strict façon XML (échec d'adoption) |
| 2014  | **HTML5** | Sémantique (`<header>`, `<article>`…), `<video>`, `<audio>`, `<canvas>`, formulaires enrichis, API |
| 2021+ | HTML Living Standard | Le HTML n'a plus de numéro : c'est un standard **vivant** maintenu par le WHATWG |

**Ce que HTML5 a apporté d'essentiel :**

- Balises **sémantiques** : le code décrit le *sens*, pas juste l'apparence.
- Médias **natifs** (`<video>`, `<audio>`) : fini Flash.
- Formulaires intelligents : `type="email"`, `required`, `date`, validation native.
- API : géolocalisation, stockage local, drag & drop, etc.

> 🎯 Dans ce guide, « HTML » = HTML5 / HTML Living Standard. Tout le reste est de l'histoire.

---

## 4. Tooling : éditeurs et environnement

### Éditeurs recommandés

| Éditeur | Points forts | Prix |
|---------|--------------|------|
| **VS Code** | Standard de fait, extensions HTML/CSS, Emmet intégré, Live Server | Gratuit |
| Sublime Text | Ultra rapide, léger | Gratuit / licence |
| WebStorm | IDE complet, refactorings puissants | Payant |
| Notepad++ | Léger, Windows | Gratuit |

### Extensions VS Code indispensables

- **Live Server** : recharge la page à chaque sauvegarde (`Alt+L, Alt+O`).
- **Prettier** : formate le code automatiquement à la sauvegarde.
- **Emmet** : déjà intégré — tape `!` + `Tab` pour générer un squelette HTML5.
- **HTMLHint** / **W3C Validator** : signale les erreurs en direct.
- **axe Accessibility Linter** : repère les problèmes d'accessibilité.

### Raccourci Emmet à connaître

Tape puis `Tab` :

| Abréviation | Résultat |
|-------------|----------|
| `!` | Squelette HTML5 complet |
| `ul>li*3` | `<ul>` avec 3 `<li>` |
| `form>input:text+input:submit` | Formulaire avec champ texte + bouton |
| `table>tr*2>td*3` | Tableau 2 lignes × 3 colonnes |

### Organisation conseillée d'un projet

```text
mon-site/
├── index.html
├── css/
│   └── style.css
├── js/
│   └── main.js
├── img/
│   └── logo.svg
└── docs/
    └── procedure-onduleur.html
```

> 💡 Sépare toujours structure (HTML), présentation (CSS) et comportement (JS) en fichiers distincts. C'est la base de tout projet maintenable.

---

## 5. Tooling : DevTools du navigateur

Les **outils de développement** (F12 ou `Ctrl+Maj+I`) sont ton meilleur ami. Onglets essentiels :

| Onglet | Usage |
|--------|-------|
| **Elements** | Inspecte le DOM, modifie HTML/CSS en direct |
| **Console** | Erreurs JS, tests rapides |
| **Network** | Temps de chargement, requêtes, poids des images |
| **Lighthouse** | Audit perf / accessibilité / SEO en 1 clic |
| **Application** | Stockage local, cookies |

**Réflexes à prendre :**

1. `Clic droit → Inspecter` sur n'importe quel élément pour voir son HTML/CSS.
2. Dans **Elements**, double-clique un style CSS pour le modifier en direct (sans toucher au fichier).
3. **Lighthouse** (`Onglet Lighthouse → Analyze page load`) : vise un score ≥ 90 en Performance et Accessibilité pour tes pages internes.
4. L'émulateur mobile (`Ctrl+Maj+M`) : teste ton responsive sans téléphone.

> 🎯 80 % du débogage HTML/CSS se fait dans l'onglet Elements. Apprivoise-le avant tout.

---

## 6. Validation W3C : ton contrôle qualité

Le **validateur du W3C** (https://validator.w3.org) vérifie que ton HTML respecte le standard. C'est gratuit et sans compte.

**3 façons de valider :**

1. **Par URL** : colle l'adresse de ta page (si elle est en ligne).
2. **Par upload** : envoie ton fichier `.html`.
3. **Par saisie directe** : colle ton code.

**Ce que la validation détecte :**

- Balises non fermées ou mal imbriquées (`<p><div></p></div>`)
- Attributs inconnus ou mal orthographiés
- `id` en double dans la page
- `<img>` sans `alt`
- Mauvais placement (`<div>` dans `<head>`, etc.)

**Bonnes pratiques :**

- Vise **zéro erreur** avant chaque mise en production (les avertissements sont tolérables s'ils sont compris).
- Intègre la validation dans ta routine : coder → tester dans le navigateur → valider → corriger.
- Les frameworks génèrent parfois du HTML imparfait : valide le **rendu final** (code source de la page affichée).

> ⚠️ Un HTML valide n'est pas une garantie d'accessibilité ni de bon SEO — mais un HTML invalide est une garantie de problèmes.

---

## 7. Anatomie d'un document HTML

Voici le squelette complet, à connaître par cœur :

```html
<!DOCTYPE html>
<html lang="fr">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Titre de la page — Nom du site</title>
  <meta name="description" content="Description courte pour les moteurs de recherche.">
  <link rel="stylesheet" href="css/style.css">
  <link rel="icon" href="img/favicon.ico">
</head>
<body>
  <header>
    <h1>Titre principal</h1>
  </header>
  <main>
    <p>Contenu…</p>
  </main>
  <footer>
    <p>© 2026 Mon service</p>
  </footer>
  <script src="js/main.js"></script>
</body>
</html>
```

