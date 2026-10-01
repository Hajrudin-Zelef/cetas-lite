---
id: collect-261001-rattrapage/rattrapage/html-guide-6
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Google", "Meta"]
dates: ["2026-09-20"]
keywords: ["sandbox", "transcription"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [1156, 1387]
sha256: 9b6752b7557a73bdf5f50832779236501531c67b1a3b4e06b053295714fe0ec0
---

# Guide HTML & CSS — Le manuel complet

- Fournis une **transcription textuelle** sous le lecteur (obligatoire pour un contenu informatif).
- Ne déclenche **jamais** de son automatique : c'est désorientant pour les lecteurs d'écran (et interdit par les navigateurs sans interaction préalable).

```html
<figure>
  <audio controls src="audio/bruit-ventilateur.mp3"></audio>
  <figcaption>Enregistrement du ventilateur défectueux (12 s).
    <a href="docs/transcription-bruit.txt">Transcription</a></figcaption>
</figure>
```

---

## 31. Iframe : embarquer du contenu

```html
<!-- Carte OpenStreetMap embarquée -->
<iframe src="https://www.openstreetmap.org/export/embed.html?bbox=…"
        width="600" height="400" title="Carte du site principal"
        loading="lazy"></iframe>

<!-- Vidéo YouTube (version respectueuse : youtube-nocookie) -->
<iframe width="560" height="315"
        src="https://www.youtube-nocookie.com/embed/VIDEO_ID"
        title="Tutoriel : remplacement d'un module UPS"
        allow="accelerometer; autoplay; encrypted-media; picture-in-picture"
        allowfullscreen loading="lazy"></iframe>
```

| Attribut | Rôle |
|----------|------|
| `title` | **Obligatoire** : décrit le contenu embarqué aux lecteurs d'écran |
| `loading="lazy"` | Charge quand visible (perf) |
| `sandbox` | Restreint ce que l'iframe peut faire (scripts, formulaires…) — sécurité |
| `allowfullscreen` | Autorise le plein écran |

```html
<!-- Iframe durcie : aucun script, aucun formulaire -->
<iframe src="https://exemple.com/widget" sandbox title="Widget météo" loading="lazy"></iframe>
```

> ⚠️ Chaque iframe = un **document complet** à charger (lourd). N'en abuse pas sur un dashboard. Et `sandbox` sans `allow-scripts` bloque le JS embarqué — parfait pour du contenu tiers non fiable.

---

## 32. Balises sémantiques : structurer la page

```html
<body>
  <header>
    <img src="img/logo.svg" alt="Logo Entreprise">
    <nav aria-label="Navigation principale">
      <ul>
        <li><a href="/">Accueil</a></li>
        <li><a href="/tickets">Tickets</a></li>
        <li><a href="/docs">Documentation</a></li>
      </ul>
    </ul>
    </nav>
  </header>

  <main>
    <article>
      <h1>…</h1>
      …
    </article>
    <aside>
      <h2>Liens utiles</h2>
      …
    </aside>
  </main>

  <footer>
    <p>© 2026 — Service Systèmes & Énergies</p>
  </footer>
</body>
```

| Balise | Sens |
|--------|------|
| `<header>` | En-tête (de page ou de section) |
| `<nav>` | Bloc de navigation |
| `<main>` | Contenu principal — **un seul** par page |
| `<article>` | Contenu autonome (article, ticket, fiche) |
| `<section>` | Regroupement thématique (avec un titre !) |
| `<aside>` | Contenu complémentaire (liens, encadrés) |
| `<footer>` | Pied (de page ou de section) |

**Pourquoi c'est crucial :** les lecteurs d'écran listent ces « landmarks » pour naviguer (`D` dans NVDA). Un `<div>` ne dit rien ; un `<nav>` dit « navigation ».

---

## 33. Balises sémantiques : le contenu fin

```html
<article>
  <header>
    <h1>Remplacement des batteries UPS-02</h1>
    <p>Publié le <time datetime="2026-09-20">20 septembre 2026</time>
       par <address><a href="mailto:tech@entreprise.fr">Karim T.</a></address></p>
  </header>

  <section>
    <h2>Diagnostic</h2>
    <p>…</p>
    <blockquote>
      <p>« La tension de floating est tombée à 12,1 V par bloc. »</p>
      <cite>— Rapport de mesure n°42</cite>
    </blockquote>
  </section>

  <section>
    <h2>Procédure</h2>
    <pre><code># Vérification avant intervention
ups-test --bank A --full</code></pre>
  </section>

  <footer>
    <p>Tags : <a href="/tags/ups">onduleur</a>, <a href="/tags/batterie">batterie</a></p>
  </footer>
</article>
```

| Balise | Usage |
|--------|-------|
| `<address>` | Coordonnées de l'auteur (pas une adresse postale quelconque !) |
| `<blockquote>` + `<cite>` | Citation longue + source |
| `<q>` | Citation courte inline |
| `<pre>` | Texte préformaté (conserve espaces/retours) — idéal pour le code |
| `<hr>` | Séparation thématique (pas juste un « trait décoratif ») |
| `<details>` + `<summary>` | Accordéon natif, sans JS ! |

```html
<details>
  <summary>Voir la procédure de consignation</summary>
  <ol>
    <li>Couper le départ…</li>
    <li>Vérifier l'absence de tension…</li>
  </ol>
</details>
```

---

## 34. Balises obsolètes : à bannir

Ces balises existent encore (compatibilité) mais **ne doivent plus être utilisées**. Le validateur W3C les signale.

| ❌ Obsolète | ✅ Remplacement |
|-------------|-----------------|
| `<font>`, `<basefont>` | CSS `font-family`, `color` |
| `<center>` | CSS `text-align: center` / flexbox |
| `<b>` (présentation pure) | `<strong>` (sens) ou CSS `font-weight` |
| `<i>` (présentation pure) | `<em>` (sens) ou CSS `font-style` |
| `<u>` | CSS `text-decoration` (et évite : ressemble à un lien) |
| `<strike>` | `<del>` ou CSS `text-decoration: line-through` |
| `<table>` de mise en page | Flexbox / Grid (§59–62) |
| `<frame>`, `<frameset>` | `<iframe>` (ou rien) |
| `<marquee>`, `<blink>` | CSS animations (et pitié, ne les imite pas) |
| `bgcolor`, `align`, `border` (attributs) | CSS |

**Test express :** si tu dois mettre en forme, c'est CSS. Si la balise ne décrit que l'apparence, elle est suspecte.

---

## 35. Head : les meta essentielles

```html
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Tickets — Service Systèmes & Énergies</title>
  <meta name="description" content="Suivi des tickets d'intervention : onduleurs, groupes, TGBT.">
  <meta name="robots" content="noindex, nofollow">  <!-- page interne : pas d'indexation -->
  <meta name="theme-color" content="#0f172a">       <!-- couleur barre mobile -->
  <meta name="author" content="Service Systèmes & Énergies">
  <link rel="stylesheet" href="/css/style.css">
  <link rel="icon" href="/img/favicon.ico" sizes="any">
  <link rel="icon" href="/img/favicon.svg" type="image/svg+xml">
  <link rel="apple-touch-icon" href="/img/apple-touch-icon.png">
</head>
```

| Meta | Rôle |
|------|------|
| `charset="utf-8"` | **Première** meta du head — encodage |
| `viewport` | **Indispensable** au responsive (§65) |
| `description` | Résumé affiché dans Google (~155 caractères) |
| `robots` | `noindex` = page interne non indexée |
| `theme-color` | Teinte de l'interface mobile |

**Ordre conseillé :** charset → viewport → title → description → CSS → favicons → scripts `defer`.

---

## 36. Favicon : le détail qui fait pro

Le favicon = la petite icône d'onglet. Un site interne sans favicon fait « chantier ».

```html
<link rel="icon" href="/img/favicon.ico" sizes="any"><!-- fallback universel -->
<link rel="icon" href="/img/favicon.svg" type="image/svg+xml"><!-- moderne -->
<link rel="apple-touch-icon" href="/img/apple-touch-icon.png"><!-- iOS -->
```

**Recette minimale :**

1. Crée un SVG simple (logo ou initiale) → `favicon.svg`.
2. Exporte un PNG 180×180 → `apple-touch-icon.png`.
3. (Optionnel) Génère le `.ico` multi-tailles pour vieux navigateurs.

**Favicon SVG inline élégant :**

```html
<link rel="icon" href="data:image/svg+xml,
  <svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 100 100'>
    <rect width='100' height='100' rx='20' fill='%230f172a'/>
    <text x='50' y='68' font-size='60' text-anchor='middle' fill='white'
          font-family='sans-serif'>⚡</text>
  </svg>">
```

> 💡 Teste l'affichage en mode sombre : un favicon foncé sur onglet sombre = invisible. Prévois un fond contrasté ou `prefers-color-scheme` dans le SVG.

---

## 37. SEO on-page : les fondamentaux

Le SEO (Search Engine Optimization) commence dans le HTML — même pour un intranet, ces réflexes structurent bien les pages.

**Checklist SEO on-page :**

