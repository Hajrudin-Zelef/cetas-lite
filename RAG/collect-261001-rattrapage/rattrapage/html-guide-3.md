---
id: collect-261001-rattrapage/rattrapage/html-guide-3
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26", "2026-10-15", "2026-11-02"]
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [464, 697]
sha256: 25313e9566a4435267fabcf26ad946352d10582bc48f89c148356b49825f1abc
---

# Guide HTML & CSS — Le manuel complet

> ⚠️ Un `<li>` doit **toujours** être enfant direct de `<ul>`, `<ol>` ou `<menu>`. Jamais de `<li>` orphelin, jamais de texte brut directement dans `<ul>`.

---

## 14. Liens hypertextes

```html
<!-- Lien externe -->
<a href="https://www.apc.com">Site APC</a>

<!-- Lien interne -->
<a href="docs/procedure.html">Procédure de consignation</a>

<!-- Ancre vers une section de la même page -->
<a href="#maintenance">Aller à la maintenance</a>
…
<h2 id="maintenance">Maintenance</h2>

<!-- Email / téléphone -->
<a href="mailto:support@entreprise.fr">Nous écrire</a>
<a href="tel:+33123456789">01 23 45 67 89</a>

<!-- Téléchargement forcé -->
<a href="docs/rapport.pdf" download> Télécharger le rapport (PDF)</a>
```

### Attributs importants

| Attribut | Rôle |
|----------|------|
| `href` | Destination (URL, ancre, `mailto:`, `tel:`) |
| `target="_blank"` | Ouvre dans un nouvel onglet |
| `rel="noopener"` | **Obligatoire** avec `target="_blank"` (sécurité : empêche la page cible de piloter la tienne) |
| `download` | Propose le téléchargement au lieu de l'affichage |
| `title` | Info-bulle (à utiliser avec parcimonie) |

### Bonnes pratiques (accessibilité + SEO)

- [ ] **Jamais** « cliquez ici » : l'intitulé doit décrire la destination → `Télécharger la procédure de consignation (PDF, 2 Mo)` ✅
- [ ] `target="_blank"` + `rel="noopener"` systématiques pour les liens externes.
- [ ] Les liens doivent être **visibles** : soulignés ou nettement distincts du texte (pas que par la couleur — §47).

```html
<!-- ❌ -->
<p>Pour la doc, <a href="doc.pdf">cliquez ici</a>.</p>
<!-- ✅ -->
<p>Consultez la <a href="doc.pdf">documentation de l'onduleur (PDF)</a>.</p>
```

---

## 15. Images : `<img>`

```html
<img src="img/onduleur-3s.jpg" alt="Onduleur Easy UPS 3S 30 kVA en baie" width="800" height="600">
```

| Attribut | Rôle |
|----------|------|
| `src` | Chemin de l'image (**obligatoire**) |
| `alt` | **Texte alternatif** : lu par les lecteurs d'écran, affiché si l'image ne charge pas (**obligatoire**) |
| `width` / `height` | Dimensions intrinsèques → **évite les sauts de mise en page** (CLS) |
| `loading="lazy"` | Chargement différé hors écran (perf, §76) |
| `decoding="async"` | Décodage non bloquant |

### Écrire un bon `alt` (arbre de décision, §42 pour le détail)

- Image **informative** → décris la fonction : `alt="Schéma unifilaire du TGBT"`.
- Image **décorative** → `alt=""` (vide : les lecteurs d'écran l'ignorent).
- Image **lien/bouton** → décris l'action : `alt="Télécharger le rapport"`.
- **Jamais** « image de… » (c'est implicite), jamais de pavé de 3 phrases.

### Formats : lequel choisir ?

| Format | Usage | Particularité |
|--------|-------|---------------|
| **SVG** | Logos, icônes, schémas | Vectoriel, léger, net à toutes tailles |
| **WebP** | Photos, images générales | ~30 % plus léger que JPEG |
| **JPEG** | Photos (si WebP indisponible) | Avec pertes, pas de transparence |
| **PNG** | Captures, transparence | Lourd pour les photos |
| **AVIF** | Photos nouvelle génération | Excellent taux de compression, support croissant |
| **GIF** | Petites animations simples | Limité à 256 couleurs |

> 🎯 Règle : SVG pour tout ce qui est vectoriel, WebP/AVIF pour les photos, `width`/`height` + `loading="lazy"` systématiques.

---

## 16. Chemins de fichiers : relatifs vs absolus

C'est la source n°1 des « images qui ne s'affichent pas ». Comprends une fois, c'est réglé à vie.

```
site/
├── index.html
├── docs/
│   └── procedure.html
├── css/
│   └── style.css
└── img/
    └── logo.svg
```

| Depuis… | Pour atteindre `img/logo.svg` | Type |
|---------|-------------------------------|------|
| `index.html` | `img/logo.svg` | Relatif |
| `docs/procedure.html` | `../img/logo.svg` | Relatif (`..` = dossier parent) |
| N'importe où | `/img/logo.svg` | Absolu **depuis la racine du site** |
| N'importe où | `https://intranet.entreprise.fr/img/logo.svg` | Absolu complet |

**Règles :**

- `./` = dossier courant (optionnel : `img/x.png` ≡ `./img/x.png`).
- `../` = remonte d'un niveau. `../../` = deux niveaux.
- Le chemin absolu `/img/logo.svg` part de la **racine du site** — pratique, mais casse tout si le site est servi depuis un sous-dossier.
- **En CSS**, les chemins dans `url()` sont relatifs **au fichier CSS**, pas au HTML !

```css
/* css/style.css → pour viser img/fond.png : */
.hero { background-image: url("../img/fond.png"); }
```

> ⚠️ Erreur classique : `C:\Users\…\img\logo.png` ou `file:///…` dans `src`. Ça marche sur ton PC, nulle part ailleurs. Toujours des chemins relatifs au projet.

---

## 17. Figure, picture, srcset : images avancées

### `<figure>` + `<figcaption>` : image légendée

```html
<figure>
  <img src="img/courbe-charge.png" alt="Courbe de charge de l'onduleur sur 24 h">
  <figcaption>Fig. 3 — Courbe de charge mesurée le 26/09/2026.</figcaption>
</figure>
```

### `<picture>` : art direction (images différentes selon l'écran)

```html
<picture>
  <source media="(max-width: 600px)" srcset="img/schema-mobile.svg">
  <source media="(max-width: 1200px)" srcset="img/schema-tablette.svg">
  <img src="img/schema-desktop.svg" alt="Schéma unifilaire du local technique">
</picture>
```

### `srcset` : même image, plusieurs résolutions

```html
<img src="img/photo-800.jpg"
     srcset="img/photo-400.jpg 400w, img/photo-800.jpg 800w, img/photo-1600.jpg 1600w"
     sizes="(max-width: 600px) 400px, (max-width: 1200px) 800px, 1600px"
     alt="Baie de brassage du local technique">
```

Le navigateur choisit la meilleure résolution selon l'écran et la densité de pixels. **Indispensable** pour des pages rapides sur mobile.

---

## 18. Tableaux : les bases

```html
<table>
  <caption>Parc onduleurs — site principal</caption>
  <thead>
    <tr>
      <th scope="col">Équipement</th>
      <th scope="col">Puissance</th>
      <th scope="col">Local</th>
      <th scope="col">Prochaine maintenance</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row">UPS-01</th>
      <td>30 kVA</td>
      <td>Local TGBT</td>
      <td>15/10/2026</td>
    </tr>
    <tr>
      <th scope="row">UPS-02</th>
      <td>60 kVA</td>
      <td>Datacenter</td>
      <td>02/11/2026</td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td colspan="4">2 équipements suivis — mise à jour : 26/09/2026</td>
    </tr>
  </tfoot>
</table>
```

| Élément | Rôle |
|---------|------|
| `<caption>` | Titre du tableau (lu en premier par les lecteurs d'écran) |
| `<thead>` / `<tbody>` / `<tfoot>` | En-tête / corps / pied — structure + permet le défilement du corps en CSS |
| `<th>` | Cellule d'en-tête (**gras + centrée** par défaut) |
| `scope="col"` / `scope="row"` | Indique si l'en-tête décrit une colonne ou une ligne (**accessibilité**) |
| `colspan` / `rowspan` | Fusionne des cellules (colonnes / lignes) |

> ⚠️ **Jamais** de tableau pour la mise en page (c'était les années 2000). Tableau = données tabulaires uniquement.

---

## 19. Tableaux accessibles avancés

### En-têtes complexes : `id` + `headers`

Quand un tableau a des en-têtes sur deux niveaux, `scope` ne suffit plus : on lie chaque cellule à ses en-têtes.

```html
<table>
  <caption>Mesures électriques par départ</caption>
  <thead>
    <tr>
      <th id="depart" scope="col">Départ</th>
      <th id="u" scope="col">Tension (V)</th>
      <th id="i" scope="col">Courant (A)</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th id="d1" scope="row">TGBT-1</th>
      <td headers="d1 u">231</td>
      <td headers="d1 i">142</td>
    </tr>
  </tbody>
</table>
```

### Checklist tableau accessible

