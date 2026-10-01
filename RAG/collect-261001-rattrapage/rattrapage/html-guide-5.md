---
id: collect-261001-rattrapage/rattrapage/html-guide-5
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [930, 1155]
sha256: 18b88898d2f71d89b4174612fe76b65bdad13ba1acdbffed776764574177f7fb
---

# Guide HTML & CSS — Le manuel complet

```html
<label for="marque">Marque</label>
<input list="marques" id="marque" name="marque">
<datalist id="marques">
  <option value="Schneider Electric"></option>
  <option value="Eaton"></option>
  <option value="Vertiv"></option>
  <option value="Riello"></option>
</datalist>
```

### Output : afficher un résultat calculé

```html
<form oninput="total.value = (kva.valueAsNumber || 0) * (nb.valueAsNumber || 0)">
  <label>Puissance unitaire (kVA) <input type="number" name="kva" value="30"></label>
  <label>Nombre <input type="number" name="nb" value="2"></label>
  <p>Total : <output name="total" for="kva nb">60</output> kVA</p>
</form>
```

---

## 25. Fieldset, legend : regrouper les champs

```html
<form action="/api/tickets" method="post">
  <fieldset>
    <legend>Demandeur</legend>
    <label for="nom">Nom</label>
    <input type="text" id="nom" name="nom" required>
    <label for="svc">Service</label>
    <input type="text" id="svc" name="service">
  </fieldset>

  <fieldset>
    <legend>Équipement</legend>
    <label for="eq">Référence</label>
    <input type="text" id="eq" name="equipement" required>
  </fieldset>

  <fieldset disabled>
    <legend>Réservé au support</legend>
    <label for="tech">Technicien assigné</label>
    <input type="text" id="tech" name="tech">
  </fieldset>
</form>
```

- `<legend>` = titre du groupe, annoncé par les lecteurs d'écran avant chaque champ du groupe.
- `disabled` sur le `<fieldset>` désactive **tous** ses champs d'un coup.
- En CSS, on peut retirer la bordure par défaut : `fieldset { border: 0; padding: 0; }` (en gardant `legend` visible !).

> 🎯 Un long formulaire sans fieldset = un mur de champs. Groupe par thème : demandeur / équipement / description / pièces jointes.

---

## 26. Validation native HTML5

Le navigateur valide **gratuitement** avant l'envoi — sans JavaScript.

```html
<form action="/api/tickets" method="post">
  <label for="ref">Référence (obligatoire, 3–10 caractères)</label>
  <input type="text" id="ref" name="ref" required minlength="3" maxlength="10">

  <label for="mail">E-mail pro</label>
  <input type="email" id="mail" name="mail" required>

  <label for="tel">Téléphone (10 chiffres)</label>
  <input type="tel" id="tel" name="tel" pattern="[0-9]{10}"
         title="10 chiffres sans espaces, ex : 0612345678">

  <label for="date">Date (dans le futur)</label>
  <input type="date" id="date" name="date" min="2026-09-27" required>

  <button type="submit">Envoyer</button>
</form>
```

| Attribut | Effet |
|----------|-------|
| `required` | Champ obligatoire |
| `minlength` / `maxlength` | Longueur du texte |
| `min` / `max` | Valeur min/max (number, date…) |
| `pattern` | Expression régulière (§27) |
| `title` | Message d'aide affiché avec l'erreur de `pattern` |
| `type="email"` / `type="url"` | Validation du format |

**Pseudo-classes CSS pour le feedback visuel :**

```css
input:required { border-left: 3px solid #f59e0b; }  /* obligatoire */
input:valid { border-color: #16a34a; }               /* valide */
input:invalid { border-color: #dc2626; }             /* invalide */
```

> ⚠️ La validation HTML5 est un **confort**, pas une sécurité : toujours re-valider côté serveur. `novalidate` sur le `<form>` la désactive (utile si tu valides en JS).

---

## 27. Validation : pattern et messages personnalisés

### `pattern` : expressions régulières utiles

```html
<!-- Référence type "UPS-01" : 3 lettres, tiret, 2 chiffres -->
<input type="text" pattern="[A-Z]{3}-[0-9]{2}" title="Format : ABC-12">

<!-- Téléphone français -->
<input type="tel" pattern="(0|\+33)[1-9][0-9]{8}" title="Ex : 0612345678">

<!-- Code postal -->
<input type="text" pattern="[0-9]{5}" title="5 chiffres">

<!-- Mot de passe : 8+ caractères, 1 majuscule, 1 chiffre -->
<input type="password" pattern="(?=.*[A-Z])(?=.*[0-9]).{8,}"
       title="8 caractères min, dont 1 majuscule et 1 chiffre">
```

| Motif regex | Sens |
|-------------|------|
| `[0-9]{5}` | Exactement 5 chiffres |
| `[A-Z]{3}` | Exactement 3 majuscules |
| `.{8,}` | 8 caractères ou plus |
| `(?=.*[A-Z])` | Contient au moins une majuscule (lookahead) |
| `(0\|\+33)` | `0` ou `+33` |

### Messages d'erreur personnalisés (JS minimal)

```html
<form id="f-ticket">
  <label for="ref2">Référence</label>
  <input type="text" id="ref2" name="ref" required pattern="[A-Z]{3}-[0-9]{2}">
  <button type="submit">Envoyer</button>
</form>
<script>
  const ref = document.getElementById('ref2');
  ref.addEventListener('invalid', () => {
    ref.setCustomValidity('Format attendu : 3 lettres, un tiret, 2 chiffres (ex : UPS-01).');
  });
  ref.addEventListener('input', () => ref.setCustomValidity(''));
</script>
```

> 💡 `setCustomValidity('')` à chaque saisie : sinon le message d'erreur reste bloqué même après correction.

---

## 28. Upload de fichiers

```html
<form action="/api/upload" method="post" enctype="multipart/form-data">
  <label for="photo">Photo du défaut (JPG/PNG, max 5 Mo)</label>
  <input type="file" id="photo" name="photo" accept="image/jpeg,image/png" required>

  <label for="docs">Documents (plusieurs fichiers)</label>
  <input type="file" id="docs" name="docs" multiple accept=".pdf,.docx">

  <button type="submit">Envoyer</button>
</form>
```

| Point | Détail |
|-------|--------|
| `enctype="multipart/form-data"` | **Obligatoire** sur le `<form>`, sinon le fichier n'est pas transmis |
| `method="post"` | Obligatoire aussi (GET ne transporte pas de fichiers) |
| `accept` | Filtre la boîte de dialogue (`image/*`, `.pdf`, `audio/*`…) — indicatif, à contrôler côté serveur |
| `multiple` | Autorise plusieurs fichiers |
| `capture="environment"` | Sur mobile : ouvre directement l'appareil photo (dos) |

```html
<!-- Prise de photo directe sur mobile -->
<input type="file" accept="image/*" capture="environment">
```

> ⚠️ `accept` n'est qu'un **filtre d'interface** : un utilisateur malveillant peut envoyer n'importe quoi. Contrôle le type MIME réel, la taille et le contenu côté serveur. Toujours.

---

## 29. Vidéo : `<video>`

```html
<video controls width="640" height="360" preload="metadata"
       poster="img/apercu-maintenance.jpg">
  <source src="video/maintenance-ups.webm" type="video/webm">
  <source src="video/maintenance-ups.mp4" type="video/mp4">
  <track kind="subtitles" src="video/sous-titres-fr.vtt" srclang="fr" label="Français" default>
  <p>Votre navigateur ne supporte pas la vidéo.
     <a href="video/maintenance-ups.mp4">Télécharger la vidéo (MP4)</a></p>
</video>
```

| Attribut | Rôle |
|----------|------|
| `controls` | Affiche les contrôles natifs (lecture, volume, plein écran) |
| `poster` | Image d'aperçu avant lecture |
| `preload="metadata"` | Ne charge que les métadonnées (durée…) — **perf** |
| `autoplay` | Démarre seul — **à éviter** sauf `muted` (les navigateurs bloquent l'autoplay sonore) |
| `loop` / `muted` | Boucle / muet |
| `<track>` | Sous-titres WebVTT — **accessibilité** |

**Formats :** fournis **WebM + MP4** (H.264) pour couvrir tous les navigateurs. Le navigateur lit le premier `<source>` qu'il supporte.

> 🎯 Vidéo de fond décorative ? `autoplay muted loop playsinline` + **jamais** de son + alternative statique si `prefers-reduced-motion` (§49).

---

## 30. Audio : `<audio>`

```html
<audio controls preload="none">
  <source src="audio/alarme-ups.ogg" type="audio/ogg">
  <source src="audio/alarme-ups.mp3" type="audio/mpeg">
  <p><a href="audio/alarme-ups.mp3">Télécharger l'enregistrement (MP3)</a></p>
</audio>
```

| Attribut | Rôle |
|----------|------|
| `controls` | Barre de lecture native |
| `preload="none"` | Ne charge rien avant le clic (économe) |
| `preload="metadata"` / `"auto"` | Métadonnées / tout |

**Accessibilité audio :**

