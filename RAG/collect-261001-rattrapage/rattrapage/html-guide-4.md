---
id: collect-261001-rattrapage/rattrapage/html-guide-4
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-01-01", "2026-12-31"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [698, 929]
sha256: 58f31604d08a455c57ca10ed363d8f46556b32e927a644181eac2067906c3863
---

# Guide HTML & CSS — Le manuel complet

- [ ] `<caption>` présent (ou `aria-label` sur le `<table>`).
- [ ] En-têtes en `<th>` avec `scope` (ou `id`/`headers` si complexe).
- [ ] Pas de cellules vides utilisées pour l'espacement.
- [ ] `colspan`/`rowspan` vérifiés (une fusion cassée décale tout).
- [ ] Sur mobile : tableau scrollable horizontalement (`overflow-x: auto` sur un conteneur) plutôt que compressé à l'illisible.

```css
/* Conteneur scrollable pour tableaux sur mobile */
.table-scroll { overflow-x: auto; }
.table-scroll table { min-width: 640px; }
```

---

## 20. Formulaires : structure de base

```html
<form action="/api/tickets" method="post">
  <div>
    <label for="nom">Nom du demandeur</label>
    <input type="text" id="nom" name="nom" required>
  </div>
  <div>
    <label for="email">E-mail</label>
    <input type="email" id="email" name="email" required>
  </div>
  <button type="submit">Envoyer</button>
</form>
```

| Attribut de `<form>` | Rôle |
|----------------------|------|
| `action` | URL qui reçoit les données |
| `method="get"` | Données dans l'URL (recherche, filtres — **jamais** de données sensibles) |
| `method="post"` | Données dans le corps de la requête (formulaires, création) |
| `enctype="multipart/form-data"` | **Obligatoire** pour l'upload de fichiers (§28) |
| `novalidate` | Désactive la validation native (si tu valides en JS) |
| `autocomplete="on"` | Active la complétion du navigateur |

**Règle d'or (accessibilité) :** chaque champ a un `<label>` avec `for` = `id` du champ. Cliquer le label donne le focus au champ — et les lecteurs d'écran annoncent l'intitulé.

```html
<!-- ✅ -->
<label for="tel">Téléphone</label>
<input type="tel" id="tel" name="tel">

<!-- ❌ (placeholder seul = inaccessible, il disparaît à la saisie) -->
<input type="tel" placeholder="Téléphone">
```

> 💡 Le `placeholder` est un **indice**, pas un label. Utilise les deux : label visible + placeholder d'exemple.

---

## 21. Tous les types d'input — partie 1 : texte et choix

```html
<!-- Texte simple -->
<label for="t1">Référence équipement</label>
<input type="text" id="t1" name="ref" maxlength="20" placeholder="Ex : UPS-01">

<!-- Mot de passe (masqué) -->
<label for="mdp">Mot de passe</label>
<input type="password" id="mdp" name="mdp" autocomplete="current-password">

<!-- E-mail (validation native du format) -->
<label for="mail">E-mail</label>
<input type="email" id="mail" name="mail" required>

<!-- Téléphone (clavier adapté sur mobile) -->
<label for="tel">Téléphone</label>
<input type="tel" id="tel" name="tel" pattern="[0-9]{10}">

<!-- URL -->
<label for="site">Site du fournisseur</label>
<input type="url" id="site" name="site" placeholder="https://…">

<!-- Recherche -->
<label for="q">Rechercher</label>
<input type="search" id="q" name="q">

<!-- Zone de texte multiligne -->
<label for="desc">Description du problème</label>
<textarea id="desc" name="desc" rows="5" cols="40"></textarea>
```

| Type | Clavier mobile / comportement |
|------|-------------------------------|
| `text` | Standard |
| `password` | Masqué, propose le gestionnaire de mots de passe |
| `email` | Clavier avec `@` ; validation du format |
| `tel` | Pavé numérique |
| `url` | Clavier avec `/` et `.com` |
| `search` | Champ avec croix d'effacement |

### Cases à cocher et boutons radio

```html
<!-- Cases à cocher : choix multiples -->
<fieldset>
  <legend>Équipements concernés</legend>
  <label><input type="checkbox" name="eq" value="ups" checked> Onduleur</label>
  <label><input type="checkbox" name="eq" value="ge"> Groupe électrogène</label>
  <label><input type="checkbox" name="eq" value="tgbt"> TGBT</label>
</fieldset>

<!-- Radios : choix unique (même name = même groupe) -->
<fieldset>
  <legend>Criticité</legend>
  <label><input type="radio" name="crit" value="basse"> Basse</label>
  <label><input type="radio" name="crit" value="haute" checked> Haute</label>
  <label><input type="radio" name="crit" value="critique"> Critique</label>
</fieldset>
```

> 💡 Astuce : enrober l'`<input>` dans le `<label>` évite d'avoir à gérer `for`/`id` — la liaison est implicite.

---

## 22. Tous les types d'input — partie 2 : nombres, dates, divers

```html
<!-- Nombre avec bornes et pas -->
<label for="kva">Puissance (kVA)</label>
<input type="number" id="kva" name="kva" min="10" max="500" step="10" value="60">

<!-- Curseur -->
<label for="charge">Charge estimée : <output id="o-charge">50</output> %</label>
<input type="range" id="charge" name="charge" min="0" max="100" value="50"
       oninput="document.getElementById('o-charge').value = this.value">

<!-- Date / heure -->
<label for="d">Date d'intervention</label>
<input type="date" id="d" name="date" min="2026-01-01" max="2026-12-31">
<label for="h">Heure</label>
<input type="time" id="h" name="heure">
<label for="dt">Date et heure</label>
<input type="datetime-local" id="dt" name="rdv">
<label for="m">Mois</label>
<input type="month" id="m" name="mois">
<label for="s">Semaine</label>
<input type="week" id="s" name="semaine">

<!-- Couleur -->
<label for="c">Couleur du statut</label>
<input type="color" id="c" name="couleur" value="#16a34a">

<!-- Fichier -->
<label for="pj">Photo du défaut</label>
<input type="file" id="pj" name="pj" accept="image/*" multiple>

<!-- Caché (données techniques, pas pour l'utilisateur) -->
<input type="hidden" name="site_id" value="12">
```

| Type | Rendu |
|------|-------|
| `number` | Champ + flèches d'incrément |
| `range` | Curseur |
| `date`, `time`, `datetime-local`, `month`, `week` | Sélecteurs natifs (calendrier…) |
| `color` | Sélecteur de couleur natif |
| `file` | Bouton « Parcourir… » |
| `hidden` | Invisible, envoyé avec le formulaire |

> ⚠️ `type="number"` : pour un **téléphone** ou un **code postal**, préfère `type="text"` + `inputmode="numeric"` + `pattern` — `number` ajoute des flèches parasites et gère mal les zéros initiaux.

```html
<label for="cp">Code postal</label>
<input type="text" id="cp" name="cp" inputmode="numeric" pattern="[0-9]{5}">
```

---

## 23. Boutons et soumission

```html
<!-- Soumet le formulaire -->
<button type="submit">Créer le ticket</button>

<!-- Réinitialise le formulaire -->
<button type="reset">Effacer</button>

<!-- Bouton générique (piloté en JS) -->
<button type="button" id="btn-calcul">Calculer l'autonomie</button>

<!-- Version image (évite : préfère <button> + CSS) -->
<input type="image" src="img/go.png" alt="Envoyer">

<!-- Désactivé -->
<button type="submit" disabled>Envoyer (brouillon incomplet)</button>
```

| `type` | Effet |
|--------|-------|
| `submit` | Envoie le formulaire (défaut si `type` omis **dans** un formulaire — piège classique !) |
| `reset` | Réinitialise les champs |
| `button` | Ne fait rien seul — pour JavaScript |

> ⚠️ **Piège n°1** : un `<button>` sans `type` dans un formulaire = `type="submit"` implicite. Ton bouton « Annuler » enverra le formulaire si tu oublies `type="button"`.

```html
<!-- ❌ Ce bouton "Fermer" soumet le formulaire ! -->
<button>Fermer</button>
<!-- ✅ -->
<button type="button">Fermer</button>
```

---

## 24. Select, datalist, textarea, output

### Liste déroulante

```html
<label for="famille">Famille d'équipement</label>
<select id="famille" name="famille" required>
  <option value="">— Choisir —</option>
  <optgroup label="Énergie">
    <option value="ups">Onduleur</option>
    <option value="ge" selected>Groupe électrogène</option>
  </optgroup>
  <optgroup label="Climatisation">
    <option value="clim">Climatiseur</option>
  </optgroup>
</select>
```

- `multiple` + `size="4"` : liste à choix multiples.
- `<optgroup>` : regroupe les options sous un libellé.

### Datalist : suggestions + saisie libre

