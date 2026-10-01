---
id: collect-261001-rattrapage/rattrapage/html-guide-13
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [2907, 3130]
sha256: 330ac39a3790c3daa59a1e500ba536b542d208e57b1c760d6d858bf8c7ef07b5
---

# Guide HTML & CSS — Le manuel complet

| Propriété | Effet |
|-----------|-------|
| `break-before: page` | Saut de page avant |
| `break-inside: avoid` | Pas de coupure à l'intérieur |
| `break-after: avoid` | Pas de coupure juste après (titre orphelin) |

**Checklist print :**

- [ ] Fond blanc, texte noir (économie d'encre).
- [ ] Navigation / boutons / pubs masqués.
- [ ] Tableaux : répète l'en-tête (`thead { display: table-header-group; }`).
- [ ] Teste avec `Ctrl+P` → aperçu avant impression.

---

## 75. Reset CSS et méthodologie BEM

### Reset moderne minimal

Les navigateurs ont des styles par défaut incohérents. Ce reset les harmonise :

```css
/* 1. Box model sain */
*, *::before, *::after { box-sizing: border-box; }

/* 2. Marges par défaut supprimées (on les remet où besoin) */
* { margin: 0; }

/* 3. Médias responsives */
img, video, svg { max-width: 100%; height: auto; display: block; }

/* 4. Formulaires : héritent la typo */
input, button, textarea, select { font: inherit; }

/* 5. Body : base saine */
body {
  line-height: 1.6;
  font-family: system-ui, sans-serif;
  min-height: 100vh;
}
```

### BEM : nommer ses classes sans chaos

**BEM** = Bloc __ Élément -- Modificateur. Parfait pour des dashboards maintenus à plusieurs.

```html
<article class="carte carte--alerte">
  <h2 class="carte__titre">UPS-02 — Défaut batterie</h2>
  <p class="carte__meta">Datacenter · 26/09/2026</p>
  <button class="carte__action carte__action--primaire" type="button">
    Voir le ticket
  </button>
</article>
```

```css
.carte { /* bloc : la carte */ }
.carte__titre { /* élément : le titre DANS la carte */ }
.carte--alerte { /* modificateur : variante alerte */ }
.carte__action--primaire { /* bouton primaire de la carte */ }
```

**Avantages :** pas de conflits de noms, spécificité basse (que des classes), on comprend le HTML sans ouvrir le CSS.

---

## 76. Performance web : l'essentiel

| Levier | Action | Gain |
|--------|--------|------|
| Images | WebP/AVIF, `srcset`, `loading="lazy"`, dimensions | Souvent **-70 %** du poids |
| CSS/JS | 1 fichier chacun, minifiés, `defer` pour JS | Moins de requêtes |
| Polices | `system-ui` ou `font-display: swap` | Pas de texte invisible |
| Iframes/vidéos | `loading="lazy"`, `preload="metadata"` | Charge au besoin |
| Cache | En-têtes serveur (avec ton admin sys 😉) | Visites suivantes instantanées |

```html
<!-- Ordre de chargement optimal -->
<link rel="stylesheet" href="css/style.css">
<script src="js/main.js" defer></script>  <!-- defer : exécuté après le HTML -->
```

**Mesurer :**

- DevTools → **Network** : poids total, requêtes, temps.
- **Lighthouse** : score + recommandations concrètes.
- Objectifs raisonnables (outil interne) : < 1 Mo par page, < 2 s de chargement.

```css
/* Évite le CLS (sauts de mise en page) : réserve la place des images */
img { aspect-ratio: attr(width) / attr(height); }
/* ou directement : */
.vignette { aspect-ratio: 16 / 9; object-fit: cover; width: 100%; }
```

> 🎯 Le plus rentable : **optimiser les images**. Sur la plupart des pages, elles font 60–80 % du poids.

---

## 77. Bonnes pratiques HTML

- [ ] `<!DOCTYPE html>` + `<html lang="fr">` + `<meta charset>` + viewport (§7).
- [ ] **Un seul** `<h1>`, hiérarchie sans saut (§11).
- [ ] Balises **sémantiques** au lieu de `<div>` partout (§32–33).
- [ ] Chaque `<img>` a un `alt` pertinent (§42) + `width`/`height` (§15).
- [ ] Chaque champ de formulaire a un `<label>` (§20).
- [ ] Liens explicites (jamais « cliquez ici ») (§14).
- [ ] `rel="noopener"` avec `target="_blank"` (§14).
- [ ] Attributs entre guillemets, minuscules, balises fermées dans l'ordre (§9).
- [ ] Pas de style inline `style=""` (sauf dynamique JS) — tout en CSS (§50).
- [ ] Pas de balise obsolète (§34), pas de tableau de mise en page (§18).
- [ ] `id` uniques dans la page.
- [ ] Validation W3C à zéro erreur avant mise en prod (§6).
- [ ] Commentaires pour structurer, **jamais** d'infos sensibles dedans (§10).

---

## 78. Bonnes pratiques CSS

- [ ] `box-sizing: border-box` global (§56).
- [ ] Variables CSS pour la palette et les espacements (§68).
- [ ] Classes plutôt qu'ID ; **zéro `!important`** (§53).
- [ ] Nommage cohérent (BEM ou équivalent) (§75).
- [ ] `rem` pour les textes, `%`/`fr` pour les largeurs fluides (§55).
- [ ] Mobile-first : base mobile + `min-width` (§66).
- [ ] `:focus-visible` toujours visible (§46) ; jamais de `:hover` sans équivalent clavier (§69).
- [ ] `prefers-reduced-motion` dès qu'il y a animation (§49).
- [ ] Dark mode via variables (§73) si pertinent.
- [ ] Print CSS pour les documents (§74).
- [ ] Un seul fichier CSS par site (ou découpage logique importé), minifié en prod.
- [ ] Contrastes AA minimum (§47).

---

## 79. Erreurs classiques des débutants (12)

### 1. Oublier le doctype
→ Quirks mode, box model cassé (§8). **Toujours** `<!DOCTYPE html>` en ligne 1.

### 2. `<button>` sans `type` dans un formulaire
→ Il soumet le formulaire par défaut (§23). Toujours `type="button|submit|reset"`.

### 3. Placeholder au lieu de label
→ Inaccessible, disparaît à la saisie (§20). Label visible **+** placeholder en complément.

### 4. « Cliquez ici » comme texte de lien
→ Incompréhensible hors contexte (§14). Décris la destination.

### 5. Images sans dimensions
→ Sauts de mise en page (CLS) (§15, §76). Toujours `width` + `height`.

### 6. Chemins absolus Windows dans `src`
→ `C:\Users\…` ne marche que sur ton PC (§16). Chemins relatifs au projet.

### 7. Sauter des niveaux de titres (`h1` → `h3`)
→ Casse la navigation des lecteurs d'écran (§11). Ordre strict.

### 8. `display: none` vs `visibility: hidden` confondus
→ Le second garde l'élément focusable = focus fantôme (§58).

### 9. `outline: none` sans remplacement
→ Plus aucun repère clavier = site inutilisable sans souris (§46). Toujours un `:focus-visible` visible.

### 10. ID en double
→ `document.getElementById` ne trouve que le premier, ancres cassées, HTML invalide. Un ID = **un seul** élément.

### 11. `z-index: 9999` qui « ne marche pas »
→ Problème de **contexte d'empilement**, pas de valeur (§64). Remonte au parent.

### 12. Oublier `enctype` pour l'upload
→ Sans `enctype="multipart/form-data"`, le fichier n'est jamais transmis (§28).

---

## 80. Cas pratique 1 : page dashboard (commentée)

```html
<!DOCTYPE html>
<html lang="fr">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="robots" content="noindex, nofollow"><!-- outil interne -->
  <title>Supervision énergie — Dashboard</title>
  <link rel="stylesheet" href="css/dashboard.css">
</head>
<body>
  <!-- Skip link : accès direct au contenu (§43) -->
  <a class="skip-link" href="#contenu">Aller au contenu</a>

  <div class="page">
    <header class="topbar">
      <p class="logo">⚡ Supervision Énergie</p>
      <nav aria-label="Navigation principale">
        <ul>
          <li><a href="/" aria-current="page">Dashboard</a></li>
          <li><a href="/tickets">Tickets</a></li>
          <li><a href="/docs">Docs</a></li>
        </ul>
      </nav>
      <!-- Zone live : annonces des changements d'état (§44) -->
      <p class="horloge" aria-live="off">Maj : <time datetime="2026-09-26T23:30">23:30</time></p>
    </header>

    <main id="contenu" class="contenu">
      <h1>État des équipements</h1>

      <!-- Cartes KPI : <ul> car c'est une LISTE de données -->
      <ul class="kpis">
        <li class="kpi kpi--ok">
          <p class="kpi__valeur">42</p>
          <p class="kpi__label">Équipements en ligne</p>
        </li>
        <li class="kpi kpi--warn">
          <p class="kpi__valeur">3</p>
          <p class="kpi__label">Avertissements</p>
        </li>
        <li class="kpi kpi--ko">
          <p class="kpi__valeur">1</p>
          <p class="kpi__label">Défaut critique</p>
        </li>
      </ul>

