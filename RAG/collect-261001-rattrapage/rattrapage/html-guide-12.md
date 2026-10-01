---
id: collect-261001-rattrapage/rattrapage/html-guide-12
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [2634, 2906]
sha256: c62097e4a27daf9411871f4c2d24f3f9b5e852da4d11879c3ab79cd5db3dd7a4
---

# Guide HTML & CSS — Le manuel complet

.carte {
  background: var(--fond-carte);
  border-radius: var(--rayon);
  padding: var(--espacement);
  border-left: 4px solid var(--couleur-primaire);
}

.carte.danger {
  --couleur-primaire: var(--couleur-danger); /* surcharge LOCALE */
}
```

### Pourquoi c'est puissant

- **Thèmes** : change les variables, toute la page suit (dark mode §73).
- **Valeur de repli** : `var(--inconnue, #333)` → `#333` si non définie.
- **Cascade** : redéfinissable par composant (exemple `.danger` ci-dessus).
- **Pilotable en JS** : `element.style.setProperty('--couleur-primaire', '#16a34a')`.

```css
/* Thème par attribut data- */
:root, [data-theme="clair"] { --fond: #fff; --texte: #111; }
[data-theme="sombre"]       { --fond: #111; --texte: #f5f5f5; }
body { background: var(--fond); color: var(--texte); }
```

```html
<button onclick="document.documentElement.dataset.theme='sombre'">🌙</button>
```

> 💡 Convention : déclare toutes tes variables dans `:root`, documente-les en commentaire. C'est ton « design system » miniature.

---

## 69. Pseudo-classes : les états

```css
/* Liens : L-O-V-E H-A-T-E (ordre mnémotechnique) */
a:link    { color: #2563eb; }   /* non visité */
a:visited { color: #6d28d9; }   /* visité */
a:hover   { text-decoration: underline; }  /* survolé */
a:active  { color: #dc2626; }   /* pendant le clic */
```

### Les plus utiles au quotidien

| Pseudo-classe | Cible |
|---------------|-------|
| `:hover` | Survol souris |
| `:focus` / `:focus-visible` | Focus clavier (§46) |
| `:first-child` / `:last-child` | Premier/dernier enfant |
| `:nth-child(2n)` | 1 ligne sur 2 (tableaux zébrés) |
| `:not(.actif)` | Tout sauf… |
| `:checked` | Checkbox/radio cochée |
| `:disabled` / `:enabled` | Champs (in)actifs |
| `:required` / `:optional` | Champs (non) obligatoires |
| `:valid` / `:invalid` | Validation (§26) |
| `:empty` | Élément sans contenu |

```css
/* Tableau zébré */
tbody tr:nth-child(even) { background: #f8fafc; }
tbody tr:hover { background: #e0f2fe; }

/* Bouton désactivé */
button:disabled { opacity: .5; cursor: not-allowed; }

/* Case cochée : label mis en avant */
input:checked + label { font-weight: bold; }
```

> ⚠️ `:hover` seul = inutilisable au tactile et au clavier. Tout effet `:hover` doit avoir un équivalent `:focus-visible`.

---

## 70. Pseudo-éléments : ::before / ::after

Ils créent des **boîtes virtuelles** sans toucher au HTML — parfait pour la décoration.

```css
/* Icône externe après les liens sortants (§52) */
a[href^="http"]::after {
  content: " ↗";
  font-size: .8em;
}

/* Puce custom devant chaque item */
.liste-puces li::before {
  content: "⚡";
  margin-right: .5rem;
}

/* Grand guillemet décoratif */
blockquote::before {
  content: "“";
  font-size: 3rem;
  color: #cbd5e1;
}
```

| Règle | Détail |
|-------|--------|
| `content` | **Obligatoire** (même `content: ""`) sinon rien ne s'affiche |
| `::before` / `::after` | Syntaxe moderne à 2 deux-points |
| Accessibilité | Le `content` **n'est pas lu** par (la plupart des) lecteurs d'écran → n'y mets que du décoratif |

**Autres pseudo-éléments utiles :**

```css
p::first-line { font-weight: bold; }       /* première ligne */
p::first-letter { font-size: 2em; }        /* lettrine */
::selection { background: #bfdbfe; }      /* texte sélectionné */
input::placeholder { color: #94a3b8; }     /* placeholder */
dialog::backdrop { background: rgb(0 0 0 / .5); } /* fond de <dialog> */
```

---

## 71. Transitions : animer en douceur

```css
.btn {
  background: #2563eb;
  color: #fff;
  padding: .75rem 1.5rem;
  border-radius: .5rem;
  /* ce qui s'anime, durée, courbe, délai */
  transition: background-color .2s ease, transform .2s ease;
}

.btn:hover {
  background: #1d4ed8;
  transform: translateY(-2px);  /* soulève légèrement */
}

.btn:active {
  transform: translateY(0);      /* renfonce au clic */
}
```

| Propriété | Rôle |
|-----------|------|
| `transition-property` | Quoi animer (`background-color`, `transform`, `all`…) |
| `transition-duration` | Durée (`.2s`) |
| `transition-timing-function` | Courbe : `ease`, `linear`, `ease-in-out`, `cubic-bezier(…)` |
| `transition-delay` | Délai avant démarrage |

**Bonnes pratiques :**

- Anime `transform` et `opacity` (performants, GPU) — **évite** `width`, `height`, `top` (recalculent toute la mise en page).
- Durées : 150–300 ms pour les micro-interactions.
- Toujours le bloc `@media (prefers-reduced-motion: reduce)` (§49).

> 💡 `transition: all` = facile mais coûteux (le navigateur surveille tout). Préfère lister les propriétés.

---

## 72. Animations `@keyframes`

Pour les animations **complexes ou en boucle** (spinner de chargement, pulse d'alerte).

```css
/* Spinner de chargement */
.spinner {
  width: 40px; height: 40px;
  border: 4px solid #e2e8f0;
  border-top-color: #2563eb;
  border-radius: 50%;
  animation: rotation 1s linear infinite;
}

@keyframes rotation {
  to { transform: rotate(360deg); }
}

/* Pulse d'alerte */
.alerte-critique {
  animation: pulse 2s ease-in-out infinite;
}
@keyframes pulse {
  0%, 100% { opacity: 1; }
  50%      { opacity: .4; }
}
```

| Propriété | Rôle |
|-----------|------|
| `animation-name` | Nom du `@keyframes` |
| `animation-duration` | Durée d'un cycle |
| `animation-iteration-count` | `infinite` ou nombre |
| `animation-direction` | `alternate` (aller-retour) |
| `animation-fill-mode` | `forwards` : garde l'état final |

```html
<div class="spinner" role="status">
  <span class="sr-only">Chargement des tickets…</span>
</div>
```

> ⚠️ Une animation en boucle + `prefers-reduced-motion` (§49) = **obligatoire**. Et un spinner sans texte accessible = invisible aux lecteurs d'écran : ajoute `role="status"` + texte sr-only.

---

## 73. Dark mode propre avec les variables

```css
/* 1. Palette en variables */
:root {
  --fond: #ffffff;
  --fond-carte: #f8fafc;
  --texte: #1e293b;
  --texte-doux: #64748b;
  --bordure: #e2e8f0;
  --primaire: #2563eb;
  color-scheme: light; /* scrollbars & contrôles natifs en clair */
}

/* 2. Version sombre : on ne change QUE les variables */
@media (prefers-color-scheme: dark) {
  :root {
    --fond: #0f172a;
    --fond-carte: #1e293b;
    --texte: #f1f5f9;
    --texte-doux: #94a3b8;
    --bordure: #334155;
    --primaire: #60a5fa;
    color-scheme: dark;
  }
}

/* 3. Les composants utilisent les variables → les deux thèmes gratuits */
body { background: var(--fond); color: var(--texte); }
.carte {
  background: var(--fond-carte);
  border: 1px solid var(--bordure);
  color: var(--texte);
}
```

**Pièges du dark mode :**

- Les **images** ne s'inversent pas : prévois des visuels qui marchent sur fond sombre (ou `filter: brightness(.9)`).
- `box-shadow` quasi invisible sur fond sombre → utilise des **bordures**.
- Teste les **contrastes** dans les deux thèmes (§47) : un gris qui passe sur blanc peut casser sur noir.

> 💡 Bonus : un toggle manuel (`data-theme`, §68) + `localStorage` pour laisser l'utilisateur choisir, avec le media query comme défaut.

---

## 74. Print CSS : des pages qui s'impriment bien

Tes procédures et rapports seront imprimés. Prévois-le :

```css
@media print {
  /* On masque ce qui ne sert pas sur papier */
  header nav, footer, .actions, .no-print { display: none; }

  /* Lisible sur papier */
  body { font-size: 12pt; color: #000; background: #fff; }
  a { color: #000; text-decoration: none; }
  /* Affiche les URL des liens (utile sur papier !) */
  a[href^="http"]::after { content: " (" attr(href) ")"; }

  /* Évite les coupures moches */
  h1, h2, h3 { break-after: avoid; }
  table, figure, pre { break-inside: avoid; }

  /* Chaque grande section sur sa page */
  .page-break { break-before: page; }
}
```

