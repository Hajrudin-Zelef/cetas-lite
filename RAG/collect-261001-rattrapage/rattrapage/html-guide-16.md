---
id: collect-261001-rattrapage/rattrapage/html-guide-16
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [3581, 3748]
sha256: 440220de712259cb2d2f35c4aaee79e1b506ddd160ad70cb6b33cbdd6b544871
---

# Guide HTML & CSS — Le manuel complet

### Balises à connaître par cœur
`header nav main article section aside footer` · `h1–h6 p ul ol li dl dt dd` ·
`a img figure figcaption` · `table caption thead tbody tfoot tr th td` ·
`form label input select textarea button fieldset legend` · `video audio iframe` ·
`details summary dialog` · `strong em mark code kbd time`

### Formulaire express
```html
<form action="…" method="post">
  <label for="x">Nom</label>
  <input type="text" id="x" name="x" required>
  <button type="submit">OK</button>
</form>
```

### Sélecteurs CSS express
```css
.class #id element [attr="v"]   /* base */
parent enfant  parent > enfant   /* descendant / direct */
a + b  a ~ b                     /* adjacent / suivants */
a:hover a:focus-visible          /* états */
li:first-child li:nth-child(2n)  /* position */
::before ::after                 /* pseudo-éléments */
```

### Flexbox express
```css
.c { display:flex; justify-content:center; align-items:center; gap:1rem; }
/* justify = axe principal · align = axe secondaire */
```

### Grid express
```css
.g { display:grid; grid-template-columns:repeat(auto-fit,minmax(250px,1fr)); gap:1rem; }
```

### Centrage express
```css
.centrer { display:flex; justify-content:center; align-items:center; }
```

### Media query express
```css
@media (min-width:768px){ … }
```

### Variables express
```css
:root{ --p:#2563eb; }  .x{ color:var(--p); }
```

### Commandes utiles
| Action | Commande / touche |
|--------|-------------------|
| Valider HTML | https://validator.w3.org |
| Inspecter | `F12` ou `Ctrl+Maj+I` |
| Émulateur mobile | `Ctrl+Maj+M` (dans DevTools) |
| Audit | Onglet Lighthouse |
| Squelette Emmet | `!` + `Tab` (VS Code) |
| Formater (Prettier) | `Maj+Alt+F` (VS Code) |

---

## 86. Glossaire

| Terme | Définition |
|-------|------------|
| **Accessibilité (a11y)** | Pratiques rendant le web utilisable par tous (WCAG, RGAA) |
| **Ancre** | Lien vers un `id` de la page (`href="#section"`) |
| **ARIA** | Attributs enrichissant l'accessibilité (`role`, `aria-*`) |
| **Attribut** | Paramètre d'une balise (`href`, `src`, `class`) |
| **Balise** | Marqueur HTML (`<p>`, `<div>`) — ouvrante/fermante ou orpheline |
| **BEM** | Convention de nommage CSS : Bloc__Élément--Modificateur |
| **Box model** | Modèle de boîte : contenu + padding + border + margin |
| **Breakpoint** | Largeur d'écran déclenchant une media query |
| **Cascade** | Règles de priorité entre déclarations CSS |
| **Charset** | Jeu de caractères (`utf-8`) |
| **CLS** | Cumulative Layout Shift : sauts de mise en page au chargement |
| **DOM** | Arbre d'objets construit par le navigateur depuis le HTML |
| **Doctype** | Déclaration `<!DOCTYPE html>` (évite le quirks mode) |
| **Emmet** | Abréviations générant du HTML/CSS (intégré à VS Code) |
| **Entité HTML** | Code pour caractère spécial (`&lt;`, `&amp;`) |
| **Favicon** | Icône d'onglet du site |
| **Flexbox** | Mise en page CSS sur un axe |
| **Grid** | Mise en page CSS sur deux axes |
| **Héritage** | Transmission de propriétés CSS parent → enfant |
| **JSON-LD** | Données structurées pour les moteurs (schema.org) |
| **Landmark** | Repère sémantique (header, nav, main…) pour lecteurs d'écran |
| **Lazy loading** | Chargement différé (`loading="lazy"`) |
| **Lighthouse** | Outil d'audit (perf, accessibilité, SEO) intégré à Chrome |
| **Media query** | Règle CSS conditionnelle (`@media (min-width: …)`) |
| **Mobile-first** | Concevoir mobile d'abord, enrichir ensuite |
| **Open Graph** | Metas pour les aperçus de partage (réseaux sociaux, Teams…) |
| **Quirks mode** | Mode de compatibilité buggé sans doctype |
| **RGAA** | Référentiel d'accessibilité français |
| **Responsive** | Mise en page s'adaptant à toutes les tailles d'écran |
| **Sémantique** | Balises décrivant le sens (`article`, `nav`…) vs `<div>` générique |
| **SEO** | Optimisation pour les moteurs de recherche |
| **Skip link** | Lien « aller au contenu » pour la navigation clavier |
| **Spécificité** | Score déterminant quelle règle CSS s'applique |
| **Srcset** | Attribut fournissant plusieurs résolutions d'image |
| **Viewport** | Zone visible de la page (meta indispensable au responsive) |
| **W3C** | Organisme de standardisation du web (+ son validateur) |
| **WCAG** | Règles internationales d'accessibilité web |
| **WHATWG** | Groupe maintenant le standard HTML vivant |
| **z-index** | Ordre d'empilement des éléments positionnés |

---

## 87. Quiz : 10 questions + réponses

**Q1.** Où doit se trouver `<!DOCTYPE html>` et pourquoi ?
> **R.** En toute première ligne du document. Sans lui, le navigateur passe en quirks mode (box model faussé, comportements incohérents).

**Q2.** Quelle est la différence entre `<strong>` et `<b>` ?
> **R.** `<strong>` indique une importance forte (les lecteurs d'écran insistent dessus) ; `<b>` n'est que visuel. Même logique pour `<em>` vs `<i>`.

**Q3.** Pourquoi ce bouton pose problème dans un formulaire : `<button>Fermer</button>` ?
> **R.** Sans `type`, un `<button>` dans un formulaire vaut `type="submit"` : cliquer « Fermer » enverrait le formulaire. Il faut `type="button"`.

**Q4.** Un champ avec seulement `placeholder="Nom"` : quel est le problème ?
> **R.** Le placeholder disparaît à la saisie et n'est pas un label fiable pour les lecteurs d'écran. Il faut un vrai `<label>` ; le placeholder n'est qu'un indice complémentaire.

**Q5.** Calcule la spécificité : `nav ul li a` vs `.menu a`. Laquelle gagne ?
> **R.** `nav ul li a` = (0,0,0,4) ; `.menu a` = (0,0,1,1). Une classe bat 4 éléments : **`.menu a` gagne**.

**Q6.** Que change `box-sizing: border-box` ?
> **R.** `width`/`height` incluent padding et border (au lieu du seul contenu). Fini les largeurs qui dépassent : `width: 300px` = 300 px affichés.

**Q7.** Flexbox ou Grid pour une barre de navigation ? Pour un dashboard complet ?
> **R.** Flexbox pour la nav (un seul axe), Grid pour le dashboard (lignes + colonnes). Les deux se combinent.

**Q8.** Cite 3 attributs indispensables d'un `<img>` bien écrit.
> **R.** `src`, `alt` (texte alternatif, éventuellement vide si décoratif), `width` + `height` (évitent les sauts de mise en page). Bonus : `loading="lazy"`.

**Q9.** À quoi sert `rel="noopener"` ?
> **R.** Avec `target="_blank"`, il empêche la page cible d'accéder à `window.opener` (la page d'origine) — faille de sécurité dite « tabnabbing ».

**Q10.** Ton menu passe sous un bandeau malgré `z-index: 9999`. Diagnostic ?
> **R.** Problème de **contexte d'empilement** (§64), pas de valeur : le parent du menu crée un contexte (ex. `opacity`, `transform`, `position`+`z-index`) qui l'enferme sous le bandeau. Solution : restructurer/remonter le menu dans le DOM ou isoler avec `isolation: isolate`.

---

## 88. Pour aller plus loin

### Approfondir HTML/CSS
- **MDN Web Docs** (developer.mozilla.org) — LA référence, en français, avec compatibilités navigateurs.
- **web.dev** (Google) — guides perf, accessibilité, responsive.
- **W3C Validator** — valider en continu.
- **Can I Use** (caniuse.com) — vérifier le support d'une propriété CSS.

### Accessibilité
- **WCAG 2.2** (w3.org) — les règles officielles.
- **RGAA** (accessibilite.numerique.gouv.fr) — le référentiel français.
- **NVDA** (gratuit) — tester au lecteur d'écran sous Windows.
- **axe DevTools** (extension) — audits automatisés.

### S'entraîner
- **Flexbox Froggy** / **Grid Garden** — apprendre flexbox/grid en jouant.
- **CSS Diner** — les sélecteurs en s'amusant.
- Refais les 4 cas pratiques (§80–83) **sans regarder**, puis compare.

### Étapes suivantes (JavaScript)
1. Manipuler le DOM (`querySelector`, `addEventListener`).
2. Valider les formulaires côté client + messages custom (§27).
3. `fetch()` : envoyer le formulaire de ticket vers ton API.
4. Stocker le thème dark/light en `localStorage`.

