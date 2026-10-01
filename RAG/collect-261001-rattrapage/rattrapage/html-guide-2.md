---
id: collect-261001-rattrapage/rattrapage/html-guide-2
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [214, 463]
sha256: c6d21b6cb552e48bc8a7d924afd3a51d312e8f843102dad2b94a31e878befd91
---

# Guide HTML & CSS — Le manuel complet

| Zone | Rôle | Visible ? |
|------|------|-----------|
| `<!DOCTYPE html>` | Déclare HTML5, évite le quirks mode (§8) | Non |
| `<html lang="fr">` | Racine du document ; `lang` = langue (crucial pour lecteurs d'écran et SEO) | Non |
| `<head>` | Métadonnées : titre, charset, CSS, SEO | Non |
| `<body>` | Tout le contenu affiché | **Oui** |

**Règles d'or :**

- **Un seul** `<h1>` par page (= le sujet de la page).
- Les `<script>` en fin de `<body>` (ou avec `defer`) pour ne pas bloquer l'affichage.
- `lang="fr"` toujours renseigné — un lecteur d'écran prononce différemment selon la langue.

---

## 8. Doctype et quirks mode

La toute première ligne, `<!DOCTYPE html>`, n'est pas une balise : c'est une **instruction** au navigateur.

```html
<!DOCTYPE html>
```

**Sans doctype** (ou avec un doctype incorrect), le navigateur bascule en **quirks mode** (« mode déglingué ») : il imite les bugs d'Internet Explorer 5 pour afficher les très vieilles pages. Conséquences :

- Le **box model** est calculé différemment (§56) — tes largeurs CSS deviennent fausses.
- Des comportements CSS incohérents entre navigateurs.

**Checklist :**

- [ ] `<!DOCTYPE html>` en **première ligne**, rien avant (pas même un commentaire ni une ligne vide — certains vieux navigateurs bronchent).
- [ ] Écrit exactement comme ça : insensible à la casse, mais la forme minuscule est la convention.
- [ ] Ne **jamais** utiliser les vieux doctypes (`HTML 4.01 Transitional…`) : ils déclenchent le mode presque-standard.

> 💡 Si ta mise en page CSS semble « cassée sans raison », vérifie d'abord le doctype. C'est la cause n°1 des mystères de débutant.

---

## 9. Syntaxe : balises, attributs, imbrication

### Balise ouvrante / fermante

```html
<p>Ceci est un paragraphe.</p>
```

- `<p>` = balise ouvrante, `</p>` = balise fermante (note le `/`).
- Certaines balises sont **orphelines** (auto-fermantes) : pas de contenu, pas de balise fermante.

```html
<img src="logo.png" alt="Logo">
<br>
<hr>
<input type="text">
<meta charset="utf-8">
```

> En HTML5, `<br>` suffit (la forme XHTML `<br />` est tolérée mais inutile).

### Attributs

```html
<a href="https://example.com" target="_blank" rel="noopener">Lien</a>
```

| Règle | Exemple |
|-------|---------|
| Toujours entre **guillemets doubles** | `href="page.html"` ✅ / `href=page.html` ⚠️ |
| Noms en **minuscules** | `class`, pas `CLASS` |
| Attributs booléens : présence = vrai | `<input required>` équivaut à `required="required"` |
| Pas d'espaces autour du `=` (convention) | `id="menu"` ✅ |

### Imbrication : la règle du poupée russe

```html
<!-- ✅ Correct : fermeture dans l'ordre inverse -->
<p>Un <strong>mot important</strong> ici.</p>

<!-- ❌ Incorrect : chevauchement -->
<p>Un <strong>mot important</p></strong>
```

### Les 3 familles de balises

| Famille | Exemples | Comportement |
|---------|----------|--------------|
| **Block** | `<p>`, `<div>`, `<h1>`, `<ul>` | Occupe toute la largeur, retour à la ligne avant/après |
| **Inline** | `<a>`, `<strong>`, `<span>`, `<img>` | Dans le flux du texte, pas de retour à la ligne |
| **Inline-block** (via CSS) | — | Comme inline mais avec largeur/hauteur réglables |

---

## 10. Commentaires et entités HTML

### Commentaires

```html
<!-- Ceci est un commentaire : invisible à l'écran, visible dans le code source -->

<!--
  Commentaire
  sur plusieurs lignes —
  pratique pour structurer un long fichier
-->

<!-- TODO: ajouter le tableau des onduleurs quand les specs arrivent -->
```

**Usages pro :** baliser les grandes zones (`<!-- /header -->`), laisser des TODO, désactiver temporairement du code.

> ⚠️ Les commentaires sont **visibles** dans le code source (Ctrl+U). N'y mets jamais d'infos sensibles (mots de passe, remarques sur des collègues…).

### Entités HTML (caractères spéciaux)

Certains caractères ont un sens en HTML : pour les afficher, on utilise des **entités**.

| Caractère | Entité | Usage |
|-----------|--------|-------|
| `<` | `&lt;` | Afficher du code : `&lt;div&gt;` → `<div>` |
| `>` | `&gt;` | Idem |
| `&` | `&amp;` | « R&D » → `R&amp;D` |
| `"` | `&quot;` | Dans un attribut délimité par `"` |
| espace insécable | `&nbsp;` | Empêche un retour à la ligne : `10&nbsp;kVA` |
| `€` | `&euro;` | (ou taper € directement en UTF-8) |
| `©` | `&copy;` | © |

**Règle :** avec `<meta charset="utf-8">`, tu peux taper directement `é`, `€`, `→`. Les entités restent indispensables pour `<`, `>`, `&`.

---

## 11. Titres et hiérarchie

```html
<h1>Rapport d'intervention — Onduleur Q4</h1>
<h2>Contexte</h2>
<h3>Équipement concerné</h3>
<h3>Symptômes observés</h3>
<h2>Actions réalisées</h2>
<h2>Recommandations</h2>
```

**Règles strictes (SEO + accessibilité) :**

1. **Un seul `<h1>`** par page : le sujet principal.
2. **Ne jamais sauter de niveau** : pas de `<h4>` directement sous un `<h2>`.
3. Les titres structurent le **plan** du document — les lecteurs d'écran permettent de naviguer de titre en titre (`H` dans NVDA).
4. Ne choisis **jamais** un niveau de titre pour sa taille visuelle : la taille, c'est CSS (`h2 { font-size: … }`).

**Mauvais réflexe classique :**

```html
<!-- ❌ Titre choisi pour faire "joli petit" -->
<h4>Mon sous-titre</h4>

<!-- ✅ Structure logique, taille gérée en CSS -->
<h2>Mon sous-titre</h2>
```

> 🎯 Test : lis tes titres seuls, dans l'ordre. Si ça forme un sommaire cohérent, ta hiérarchie est bonne.

---

## 12. Paragraphes et texte inline

```html
<p>Un paragraphe regroupe une idée. Les retours à la ligne
dans le code source ne créent <strong>pas</strong> de retour
à la ligne à l'écran : il faut un nouveau <code>&lt;p&gt;</code>.</p>
```

### Balises inline essentielles

| Balise | Sens | Rendu par défaut |
|--------|------|------------------|
| `<strong>` | Importance forte | **gras** |
| `<em>` | Emphase | *italique* |
| `<mark>` | Surlignage / pertinent | <mark>surligné</mark> |
| `<code>` | Morceau de code | `monospace` |
| `<kbd>` | Touche clavier | <kbd>Ctrl</kbd> |
| `<samp>` | Sortie d'un programme | monospace |
| `<var>` | Variable | *italique* |
| `<small>` | Texte secondaire | plus petit |
| `<del>` / `<ins>` | Supprimé / inséré | <del>barré</del> / <ins>souligné</ins> |
| `<sub>` / `<sup>` | Indice / exposant | H<sub>2</sub>O, m<sup>2</sup> |
| `<abbr>` | Abréviation | `title` = signification |
| `<time>` | Date/heure machine | — |

**Exemples :**

```html
<p>Appuyez sur <kbd>Ctrl</kbd> + <kbd>S</kbd> pour enregistrer.</p>
<p>La tension est de <var>U</var> = 230&nbsp;V.</p>
<p>Le <abbr title="Uninterruptible Power Supply">UPS</abbr> est en défaut.</p>
<p>Intervention le <time datetime="2026-09-26">26 septembre 2026</time>.</p>
<p>Ancien prix : <del>1 200 €</del>, nouveau : <ins>990 €</ins>.</p>
```

> 💡 Préfère `<strong>` à `<b>` et `<em>` à `<i>` : les premiers portent du **sens** (les lecteurs d'écran insistent dessus), les seconds ne sont que visuels.

---

## 13. Listes : ordonnées, non ordonnées, définitions

### Liste non ordonnée (`<ul>`)

```html
<ul>
  <li>Vérifier le niveau d'électrolyte</li>
  <li>Contrôler la tension de floating</li>
  <li>Nettoyer les filtres</li>
</ul>
```

### Liste ordonnée (`<ol>`)

```html
<ol>
  <li>Consigner l'armoire</li>
  <li>Mesurer la tension du bus DC</li>
  <li>Remplacer le module</li>
</ol>
```

Attributs utiles de `<ol>` : `start="5"` (commence à 5), `reversed` (décroissant), `type="A"` (`A, B, C…`), `type="i"` (chiffres romains).

### Liste de définitions (`<dl>`)

```html
<dl>
  <dt>UPS</dt>
  <dd>Uninterruptible Power Supply — onduleur.</dd>
  <dt>VFI</dt>
  <dd>Voltage and Frequency Independent — double conversion.</dd>
</dl>
```

### Listes imbriquées

```html
<ul>
  <li>Onduleurs
    <ul>
      <li>Easy UPS 3S</li>
      <li>Easy UPS 3M</li>
    </ul>
  </li>
  <li>Groupes électrogènes</li>
</ul>
```

