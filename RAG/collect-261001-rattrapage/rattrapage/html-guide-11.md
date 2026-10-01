---
id: collect-261001-rattrapage/rattrapage/html-guide-11
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [2343, 2633]
sha256: 84e6fadd1e03a9a7679a339bfc97064f00c161bbf3fb07e9b626f5eedfda971f
---

# Guide HTML & CSS — Le manuel complet

| Notation | Sens |
|----------|------|
| `1fr` | 1 fraction de l'espace **restant** |
| `repeat(3, 1fr)` | 3 colonnes égales |
| `repeat(auto-fit, minmax(250px, 1fr))` | **Responsive auto** : autant de colonnes de 250 px min que possible |
| `minmax(200px, 1fr)` | Entre 200 px et l'espace dispo |

```css
/* Galerie responsive sans media query ! */
.galerie {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}
```

### Placer les items

```css
.item-important {
  grid-column: 1 / 3;   /* de la ligne 1 à la ligne 3 → occupe 2 colonnes */
  grid-row: 1 / 2;
  /* ou */
  grid-column: span 2;  /* occupe 2 colonnes */
}
```

> 🎯 `repeat(auto-fit, minmax(250px, 1fr))` : LA recette d'une grille responsive sans breakpoint. Retiens-la.

---

## 62. Grid : areas et layout de page

### `grid-template-areas` : dessine ta page en ASCII

```css
.page {
  display: grid;
  grid-template-columns: 220px 1fr;
  grid-template-rows: auto 1fr auto;
  grid-template-areas:
    "header header"
    "sidebar main"
    "footer footer";
  min-height: 100vh;
}
.page > header  { grid-area: header; }
.page > nav     { grid-area: sidebar; }
.page > main    { grid-area: main; }
.page > footer  { grid-area: footer; }
```

```html
<div class="page">
  <header>…</header>
  <nav>…</nav>
  <main>…</main>
  <footer>…</footer>
</div>
```

### Version mobile : on redessine

```css
@media (max-width: 768px) {
  .page {
    grid-template-columns: 1fr;
    grid-template-areas:
      "header"
      "main"
      "sidebar"
      "footer";
  }
}
```

### Flexbox ou Grid ? La règle

| Besoin | Choix |
|--------|-------|
| Aligner des éléments sur **un axe** (nav, ligne de boutons) | **Flexbox** |
| Disposer en **lignes ET colonnes** (dashboard, page) | **Grid** |
| Ils se **combinent** : grid pour la page, flexbox dans les cartes | ✅ |

> 💡 Un dashboard typique : `grid` pour la coquille (header/sidebar/main), `flexbox` à l'intérieur des cartes de KPI.

---

## 63. Positionnement CSS

```css
.boite { position: static; }    /* défaut : dans le flux normal */
```

| Valeur | Comportement |
|--------|--------------|
| `relative` | Décalé par rapport à **sa position normale** (`top/left/…`) ; reste dans le flux |
| `absolute` | Sorti du flux, positionné par rapport au **parent positionné** le plus proche |
| `fixed` | Sorti du flux, positionné par rapport au **viewport** (reste visible au scroll) |
| `sticky` | Hybride : normal, puis « collé » quand on scrolle jusqu'à lui |

```css
/* Badge sur une carte */
.carte { position: relative; }
.badge {
  position: absolute;
  top: 0.5rem; right: 0.5rem;
}

/* En-tête qui reste visible au scroll */
.entete { position: sticky; top: 0; background: #fff; z-index: 10; }

/* Bouton d'aide flottant */
.aide { position: fixed; bottom: 1rem; right: 1rem; }
```

> ⚠️ `absolute` se réfère au parent **positionné** le plus proche (pas `static`). Oublier `position: relative` sur le parent = l'élément part se caler sur la page entière. Classique !

---

## 64. Z-index et contextes d'empilement

`z-index` contrôle **qui passe devant qui** — mais seulement entre éléments **positionnés** (ou flex/grid items).

```css
.modale { position: fixed; z-index: 100; }
.menu-deroulant { position: absolute; z-index: 50; }
```

**Échelle conseillée** (évite les `z-index: 9999`) :

| Couche | z-index |
|--------|---------|
| Contenu normal | auto / 0 |
| Header sticky | 10 |
| Menu déroulant | 50 |
| Modale + overlay | 100 |
| Tooltip | 200 |
| Skip link | 300 |

### Le piège : les contextes d'empilement

Un `z-index` ne s'applique qu'**à l'intérieur** du contexte de son parent. Un enfant avec `z-index: 9999` dans un parent à `z-index: 1` **restera sous** un voisin à `z-index: 2`.

Créent un contexte : `position` + `z-index`, `opacity < 1`, `transform`, `filter`, `isolation: isolate`.

```css
/* Isole proprement un composant */
.carte { isolation: isolate; }
```

> 💡 Si ton menu passe **sous** un élément malgré un gros z-index, c'est un problème de contexte, pas de valeur. Remonte au parent.

---

## 65. Responsive : viewport et media queries

### La meta viewport (rappel §35) — sans elle, pas de responsive

```html
<meta name="viewport" content="width=device-width, initial-scale=1">
```

Sans elle, le mobile affiche la page « zoomée arrière » (viewport virtuel ~980 px).

### Media queries : adapter selon l'écran

```css
/* Base : mobile (mobile-first, §66) */
.carte { padding: 1rem; }

/* À partir de 768 px : tablette */
@media (min-width: 768px) {
  .carte { padding: 2rem; }
}

/* À partir de 1024 px : desktop */
@media (min-width: 1024px) {
  .grille { grid-template-columns: repeat(3, 1fr); }
}
```

| Syntaxe | Sens |
|---------|------|
| `(min-width: 768px)` | À partir de 768 px (mobile-first ✅) |
| `(max-width: 767px)` | Jusqu'à 767 px (desktop-first) |
| `(orientation: landscape)` | Paysage |
| `(prefers-color-scheme: dark)` | Dark mode OS |
| `(prefers-reduced-motion: reduce)` | Animations réduites |

**Opérateurs :**

```css
@media (min-width: 768px) and (max-width: 1023px) { /* tablette uniquement */ }
@media (min-width: 768px), (orientation: landscape) { /* OU logique */ }
@media not (prefers-reduced-motion: reduce) { /* sauf… */ }
```

---

## 66. Mobile-first et breakpoints

### Mobile-first : la méthode

1. Écris le CSS pour **mobile** d'abord (1 colonne, navigation simplifiée).
2. Ajoute des `min-width` pour **enrichir** sur grand écran.

```css
/* 1. Mobile : une colonne */
.dashboard { display: grid; gap: 1rem; }

/* 2. Tablette : deux colonnes */
@media (min-width: 768px) {
  .dashboard { grid-template-columns: repeat(2, 1fr); }
}

/* 3. Desktop : sidebar + contenu */
@media (min-width: 1024px) {
  .dashboard { grid-template-columns: 240px repeat(3, 1fr); }
}
```

**Pourquoi mobile-first ?**

- Le CSS mobile est le plus simple → base saine.
- `min-width` s'empile naturellement (pas de conflits d'annulation).
- Performances : le mobile charge le CSS de base, pas les surcharges desktop.

### Breakpoints : pas de valeurs magiques

| Usage courant | Breakpoint |
|---------------|------------|
| Mobile | < 768 px (base) |
| Tablette | ≥ 768 px |
| Desktop | ≥ 1024 px |
| Grand écran | ≥ 1440 px |

> 🎯 Mieux : des breakpoints **au contenu** (« quand ma carte devient trop étroite ») plutôt qu'aux appareils. Et `auto-fit`/`minmax` (§61) évitent souvent toute media query.

---

## 67. Images et médias responsives (CSS)

```css
/* L'image ne dépasse JAMAIS de son conteneur */
img, video, svg {
  max-width: 100%;
  height: auto;   /* garde les proportions (avec width/height HTML, §15) */
}

/* Image de fond qui couvre son cadre */
.hero {
  background-image: url("../img/local.jpg");
  background-size: cover;       /* couvre tout, rogne si besoin */
  background-position: center; /* centre le point d'intérêt */
  min-height: 40vh;
}
/* alternatives background-size : contain (tout visible), auto */
```

### `<picture>` vs CSS : qui fait quoi ?

| Besoin | Solution |
|--------|----------|
| Même image, taille adaptée | `srcset`/`sizes` (§17) + `max-width: 100%` |
| Image **différente** selon l'écran | `<picture>` (§17) |
| Image décorative de fond | CSS `background-image` + media queries |

```css
/* Fond différent sur mobile (plus léger) */
.hero { background-image: url("../img/hero-mobile.jpg"); }
@media (min-width: 768px) {
  .hero { background-image: url("../img/hero-desktop.jpg"); }
}
```

> ⚠️ `background-image` n'a **pas de `alt`** : réservé au décoratif. Toute image informative = `<img>` avec `alt`.

---

## 68. Variables CSS (custom properties)

```css
:root {
  --couleur-primaire: #2563eb;
  --couleur-danger: #dc2626;
  --fond-carte: #ffffff;
  --rayon: 0.5rem;
  --espacement: 1rem;
}

