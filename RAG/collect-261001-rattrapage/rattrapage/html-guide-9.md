---
id: collect-261001-rattrapage/rattrapage/html-guide-9
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [1836, 2091]
sha256: 1f1504fe23693dbdb03d793be846e9923c6f2e8d87861830d4123514f87a07f3
---

# Guide HTML & CSS — Le manuel complet

## 49. Préférences utilisateur : reduced motion, dark mode

### `prefers-reduced-motion` : respecter les utilisateurs sensibles

```css
/* Désactive animations/transitions pour ceux qui l'ont demandé */
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
    scroll-behavior: auto !important;
  }
}
```

À mettre **systématiquement** dans tes CSS dès que tu animes quelque chose (§72).

### `prefers-color-scheme` : dark mode natif (§73 pour le détail)

```css
:root { --fond: #ffffff; --texte: #111827; }
@media (prefers-color-scheme: dark) {
  :root { --fond: #111827; --texte: #f9fafb; }
}
body { background: var(--fond); color: var(--texte); }
```

> 💡 Ces media queries lisent les **réglages OS** de l'utilisateur : tu n'as rien à demander, juste à respecter.

---

## 50. CSS : introduction et liaison

**CSS** (Cascading Style Sheets) décrit la présentation. Une règle = **sélecteur** + **déclarations**.

```css
/* sélecteur */ h1 {
  color: #0f172a;        /* déclaration : propriété + valeur */
  font-size: 2rem;
  margin-bottom: 1rem;
}
```

### 3 façons d'appliquer du CSS

```html
<!-- 1. Feuille externe (✅ À PRIVILÉGIER : cache, réutilisable) -->
<link rel="stylesheet" href="css/style.css">

<!-- 2. <style> dans le head (⚠️ page unique, prototypes) -->
<style>
  .alerte { color: #dc2626; }
</style>

<!-- 3. Attribut style (❌ sauf cas dynamiques en JS) -->
<p style="color: red;">Urgent</p>
```

**Ordre de priorité** (la « cascade ») : `style=""` > `<style>`/feuille externe (selon l'ordre) — voir spécificité §53.

### Premier fichier CSS type

```css
/* css/style.css */
:root {
  --bleu: #2563eb;
  --fond: #f8fafc;
}

body {
  font-family: system-ui, sans-serif;
  background: var(--fond);
  color: #1e293b;
  line-height: 1.6;
  margin: 0;
}
```

> 🎯 `system-ui` : utilise la police système (rapide, native, zéro téléchargement). Parfait pour des outils internes.

---

## 51. Sélecteurs CSS : les bases

| Sélecteur | Cible | Exemple |
|-----------|-------|---------|
| `*` | Tout | `* { box-sizing: border-box; }` |
| `p` | Toutes les balises `<p>` | `p { margin: 0 0 1rem; }` |
| `.alerte` | Classe `alerte` | `<p class="alerte">` |
| `#menu` | ID `menu` (unique) | `<nav id="menu">` |
| `[type="email"]` | Attribut | `<input type="email">` |
| `h1, h2` | Groupe | Les deux |

```html
<p class="alerte critique">Défaut onduleur !</p>
<p class="alerte">Avertissement batterie.</p>
```

```css
.alerte { padding: 1rem; border-left: 4px solid #f59e0b; }
.alerte.critique { border-color: #dc2626; background: #fef2f2; } /* les deux classes */
```

**Classes vs ID :**

- **Classe** (`.`) : réutilisable, c'est 99 % de ton CSS.
- **ID** (`#`) : unique par page ; utile pour les ancres et le JS, **évite** en CSS (spécificité trop forte, §53).

> 💡 Convention de nommage : minuscules + tirets (`.carte-equipement`), jamais d'accents ni d'espaces.

---

## 52. Sélecteurs avancés : combinateurs et attributs

### Combinateurs

```css
/* Descendant : tout <a> DANS <nav> (peu importe la profondeur) */
nav a { color: #2563eb; }

/* Enfant direct : <li> enfants directs de <ul class="menu"> */
ul.menu > li { list-style: none; }

/* Frère adjacent : <p> juste APRÈS un <h2> */
h2 + p { margin-top: 0; }

/* Frères suivants : tous les <p> après un <h2> (même parent) */
h2 ~ p { color: #334155; }
```

### Sélecteurs d'attributs

```css
/* Présence */
a[target] { /* liens avec target */ }

/* Valeur exacte */
input[type="email"] { border-color: #2563eb; }

/* Commence par : liens externes http */
a[href^="http"]::after { content: " ↗"; }

/* Finit par : liens PDF */
a[href$=".pdf"]::before { content: "📄 "; }

/* Contient */
img[src*="logo"] { /* images dont le src contient "logo" */ }
```

| Opérateur | Sens |
|-----------|------|
| `[attr]` | L'attribut existe |
| `[attr="v"]` | Égal à `v` |
| `[attr^="v"]` | Commence par `v` |
| `[attr$="v"]` | Finit par `v` |
| `[attr*="v"]` | Contient `v` |

> 🎯 `a[href^="http"]` + `::after "↗"` : signale les liens externes **sans toucher au HTML**. Élégant et maintenable.

---

## 53. Spécificité et cascade : qui gagne ?

Quand plusieurs règles ciblent le même élément, le navigateur calcule un **score** : `(inline, ID, classes, éléments)`.

```css
p { color: black; }              /* (0,0,0,1) */
.alerte { color: orange; }       /* (0,0,1,0) → gagne sur p */
#msg { color: blue; }            /* (0,1,0,0) → gagne sur .alerte */
p.alerte { color: green; }       /* (0,0,1,1) → gagne sur .alerte seul */
```

**Règles de départage :**

1. `!important` gagne sur tout (à éviter — c'est un aveu d'échec).
2. Score le plus élevé.
3. À égalité : la règle **déclarée en dernier** gagne.

```html
<p id="msg" class="alerte">Défaut</p>  <!-- bleu : #msg gagne -->
```

### Bonnes pratiques

- [ ] **Zéro ID** en CSS, **zéro `!important`** → tes scores restent bas et surchargeables.
- [ ] Préfère des classes explicites (`.btn`, `.btn-primaire`) à des sélecteurs longs (`header nav ul li a`).
- [ ] Si une règle « ne s'applique pas », c'est 99 % un problème de **spécificité** : inspecte dans DevTools (les règles barrées = perdantes).

> 💡 DevTools → onglet Elements → panneau Styles : les déclarations **barrées** sont celles battues par la spécificité. Ton meilleur prof de CSS.

---

## 54. Héritage en CSS

Certaines propriétés **s'héritent** (les enfants reprennent la valeur du parent), d'autres non.

| Héritées ✅ | Non héritées ❌ |
|-------------|-----------------|
| `color`, `font-*`, `line-height` | `margin`, `padding`, `border` |
| `text-align`, `visibility` | `background`, `width`, `height` |
| `list-style`, `cursor` | `display`, `position`, `float` |

```css
body {
  font-family: system-ui, sans-serif;  /* hérité par TOUTE la page */
  color: #1e293b;                      /* hérité aussi */
  margin: 0;                           /* NON hérité : à mettre où besoin */
}
```

### Mots-clés de contrôle

```css
.carte {
  color: inherit;   /* force l'héritage depuis le parent */
  margin: initial;  /* revient à la valeur par défaut du navigateur */
  border: unset;    /* inherit si héritable, sinon initial */
}
```

> 🎯 Tire parti de l'héritage : déclare `font-family` et `color` sur `body`, pas sur chaque élément. Moins de code, plus de cohérence.

---

## 55. Unités et couleurs en CSS

### Unités : absolues vs relatives

| Unité | Type | Équivalence / usage |
|-------|------|---------------------|
| `px` | Absolue | Pixel — bordures, détails fins |
| `rem` | Relative | × taille de police racine (`html`, défaut 16 px) — **textes, espacements** |
| `em` | Relative | × taille de police de l'élément parent — espacements proportionnels |
| `%` | Relative | % du parent — largeurs fluides |
| `vw` / `vh` | Relative | % de la largeur/hauteur du viewport — plein écran |
| `ch` | Relative | Largeur du « 0 » — largeurs de texte optimales (`max-width: 65ch`) |

```css
html { font-size: 16px; }   /* 1rem = 16px */
h1 { font-size: 2rem; }      /* = 32px, suit le réglage utilisateur */
.carte { width: 90%; max-width: 60ch; padding: 1.5rem; }
.hero { min-height: 50vh; }  /* moitié de la hauteur d'écran */
```

> 🎯 `rem` pour les tailles de texte : si l'utilisateur a réglé une police plus grande (malvoyance), tout suit. `px` fige — à réserver aux bordures.

### Couleurs : 4 notations

```css
.rouge-nomme  { color: red; }                  /* 148 noms — lisible, limité */
.rouge-hex    { color: #dc2626; }              /* hexadécimal — le standard */
.rouge-rgb    { color: rgb(220 38 38 / 0.8); } /* rgb + alpha (transparence) */
.rouge-hsl    { color: hsl(0 84% 60%); }       /* teinte/saturation/luminosité — intuitif */
```

