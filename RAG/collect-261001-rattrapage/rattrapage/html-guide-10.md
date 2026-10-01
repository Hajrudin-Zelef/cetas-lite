---
id: collect-261001-rattrapage/rattrapage/html-guide-10
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [2092, 2342]
sha256: 3b5e765337bb6939518661085d96b5593a0baf6f8be3cd94ee48225f7d6d097f
---

# Guide HTML & CSS — Le manuel complet

**HSL, ton ami :** pour une palette, fixe la teinte et joue sur saturation/luminosité.

```css
:root {
  --teinte: 217; /* bleu */
  --primaire: hsl(var(--teinte) 91% 60%);
  --primaire-fonce: hsl(var(--teinte) 91% 45%);
  --primaire-clair: hsl(var(--teinte) 91% 92%);
}
```

---

## 56. Box model : le concept central

Chaque élément = **4 couches** emboîtées (de l'intérieur vers l'extérieur) :

```
┌──────────── margin ────────────┐
│ ┌────────── border ──────────┐ │
│ │ ┌─────── padding ────────┐ │ │
│ │ │      CONTENU           │ │ │
│ │ │  (width × height)      │ │ │
│ │ └────────────────────────┘ │ │
│ └────────────────────────────┘ │
└────────────────────────────────┘
```

```css
.boite {
  width: 300px;            /* contenu */
  padding: 20px;           /* intérieur : fond coloré inclus */
  border: 5px solid #333;   /* bordure */
  margin: 30px;            /* extérieur : transparent, espace avec voisins */
}
```

### `box-sizing` : le réglage qui change tout

```css
/* ✅ À mettre en tête de CHAQUE projet */
*, *::before, *::after { box-sizing: border-box; }
```

| `box-sizing` | `width: 300px` + `padding: 20px` + `border: 5px` = largeur totale |
|--------------|-------------------------------------------------------------------|
| `content-box` (défaut historique) | 300 + 40 + 10 = **350 px** 😱 |
| `border-box` | **300 px** tout compris ✅ |

> 💡 Avec `border-box`, `width` = largeur **totale** visible. Les calculs de mise en page deviennent intuitifs. C'est le défaut de tous les frameworks modernes.

### Fusion des marges (margin collapsing)

Deux marges verticales adjacentes **fusionnent** (la plus grande gagne) :

```css
h2 { margin-bottom: 30px; }
p  { margin-top: 20px; }
/* Espace réel entre les deux : 30px (pas 50px) */
```

Pas un bug : c'est la spec. On le contourne avec flexbox/grid (qui ne fusionnent pas) ou un padding.

---

## 57. Typographie CSS

```css
body {
  font-family: system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
  font-size: 1rem;          /* 16 px par défaut */
  line-height: 1.6;        /* interlignage : 1.5–1.7 = confortable */
  color: #1e293b;
}

h1, h2, h3 {
  line-height: 1.2;        /* titres plus resserrés */
  font-weight: 700;
  text-wrap: balance;      /* équilibre les lignes des titres (moderne) */
}

p { max-width: 65ch; }     /* ~65 caractères/ligne = lisibilité optimale */

.code {
  font-family: ui-monospace, "Cascadia Code", Consolas, monospace;
}
```

| Propriété | Effet |
|-----------|-------|
| `font-weight` | 100–900 ou `bold`/`normal` |
| `font-style` | `italic`, `normal` |
| `text-align` | `left`, `center`, `right`, `justify` |
| `text-transform` | `uppercase`, `lowercase`, `capitalize` |
| `letter-spacing` | Chasse (espacement des lettres) |
| `text-decoration` | `underline`, `line-through`, `none` |
| `white-space: nowrap` | Empêche le retour à la ligne |
| `text-overflow: ellipsis` | `…` quand ça dépasse (avec `overflow: hidden`) |

```css
/* Titre qui tronque proprement */
.titre-carte {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
```

> ⚠️ `text-align: justify` sur de petites largeurs crée des « rivières » de blancs. Sur mobile, préfère `left`.

---

## 58. Display : block, inline, et les autres

```css
.block       { display: block; }        /* toute la largeur, retour ligne */
.inline      { display: inline; }       /* dans le flux, pas de width/height */
.inline-bloc { display: inline-block; } /* inline + width/height réglables */
.cacher      { display: none; }         /* retiré du flux ET des lecteurs d'écran */
.flex        { display: flex; }         /* §59 */
.grille      { display: grid; }         /* §61 */
```

| Valeur | Particularité |
|--------|---------------|
| `display: none` | L'élément **n'existe plus** (mise en page + accessibilité) |
| `visibility: hidden` | Invisible mais **occupe toujours sa place** (et reste focusable ⚠️) |
| `opacity: 0` | Transparent mais **cliquable/focusable** (piège !) |

**Cacher proprement selon le besoin :**

```css
/* Caché pour tout le monde */
.cache { display: none; }

/* Caché visuellement mais lu par les lecteurs d'écran (labels discrets…) */
.sr-only {
  position: absolute; width: 1px; height: 1px;
  overflow: hidden; clip-path: inset(50%);
}
```

> ⚠️ `visibility: hidden` + `tabindex` = focus invisible au clavier. Pour masquer temporairement un élément interactif, préfère `display: none` ou `hidden`.

---

## 59. Flexbox : le conteneur

**Flexbox** = mise en page **1D** (une ligne OU une colonne). Parfait pour : barres de navigation, cartes, formulaires, centrages.

```css
.conteneur {
  display: flex;
  flex-direction: row;      /* row | row-reverse | column | column-reverse */
  justify-content: center;  /* alignement sur l'AXE PRINCIPAL */
  align-items: center;      /* alignement sur l'AXE SECONDAIRE */
  gap: 1rem;                /* espace entre enfants (moderne, remplace les margins) */
  flex-wrap: wrap;          /* retour à la ligne si ça déborde */
}
```

### `justify-content` (axe principal)

| Valeur | Effet |
|--------|-------|
| `flex-start` | Au début (défaut) |
| `center` | Centré |
| `flex-end` | À la fin |
| `space-between` | Répartis, collés aux bords |
| `space-around` | Répartis, marges égales autour |
| `space-evenly` | Espaces parfaitement égaux |

### `align-items` (axe secondaire)

`stretch` (défaut : étire) | `flex-start` | `center` | `flex-end` | `baseline`

**Le centrage parfait (fini les calculs) :**

```css
.centrer {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
}
```

> 💡 `gap` marche aussi en grid. Oublie les `margin-right` sur tous-les-enfants-sauf-le-dernier : `gap` fait ça proprement.

---

## 60. Flexbox : les items et cas concrets

```css
.item {
  flex-grow: 1;    /* grandit pour remplir l'espace (proportion) */
  flex-shrink: 1;  /* rétrécit si manque de place */
  flex-basis: 200px; /* taille de départ */
  /* raccourci : */
  flex: 1;         /* = flex: 1 1 0% → se partagent l'espace équitablement */
  align-self: flex-start; /* aligne CET item seul (surcharge align-items) */
  order: 2;        /* change l'ordre VISUEL (pas l'ordre DOM/lecteur d'écran ⚠️) */
}
```

### Cas concret 1 : barre de navigation

```html
<nav class="navbar">
  <a class="logo" href="/">⚡ Énergies</a>
  <ul class="liens">
    <li><a href="/tickets">Tickets</a></li>
    <li><a href="/docs">Docs</a></li>
  </ul>
  <button type="button">+ Nouveau ticket</button>
</nav>
```

```css
.navbar { display: flex; align-items: center; gap: 1rem; padding: 1rem; }
.navbar .liens { display: flex; gap: 1rem; list-style: none; margin: 0; padding: 0; }
.navbar button { margin-left: auto; } /* pousse le bouton à droite */
```

### Cas concret 2 : carte équipement

```css
.carte { display: flex; gap: 1rem; align-items: flex-start; }
.carte img { flex-shrink: 0; }        /* l'image ne s'écrase pas */
.carte .infos { flex: 1; }            /* le texte prend le reste */
```

> ⚠️ `order` ne change que l'affichage : un lecteur d'écran et le `Tab` suivent toujours l'ordre DOM. Ne t'en sers pas pour réordonner du contenu porteur de sens.

---

## 61. CSS Grid : les bases

**Grid** = mise en page **2D** (lignes ET colonnes). Parfait pour : dashboards, galeries, pages complètes.

```css
.grille {
  display: grid;
  grid-template-columns: 200px 1fr 1fr;  /* 3 colonnes : fixe + 2 flexibles */
  grid-template-rows: auto 1fr auto;     /* 3 lignes */
  gap: 1rem;
}
```

### Unités grid

