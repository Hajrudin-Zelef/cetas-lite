---
id: collect-261001-rattrapage/rattrapage/html-guide-8
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [1606, 1835]
sha256: 6d0a79a9d4defcb764c0d04531e1c9d7610aa905a18921c097ff5581eeda7c4c
---

# Guide HTML & CSS — Le manuel complet

```html
<header>…</header>          <!-- banner -->
<nav aria-label="Principale">…</nav>  <!-- navigation -->
<main>…</main>              <!-- main -->
<aside aria-label="Compléments">…</aside>  <!-- complementary -->
<footer>…</footer>          <!-- contentinfo -->
```

- Si **deux** `<nav>`, différencie-les avec `aria-label` (« Principale » / « Fil d'Ariane »).
- `<main>` : **un seul** par page, jamais imbriqué dans header/footer.

### Skip link : aller au contenu

Le premier élément interactif de la page : un lien « Aller au contenu » **visible au focus clavier**.

```html
<body>
  <a class="skip-link" href="#contenu">Aller au contenu principal</a>
  <header>…</header>
  <main id="contenu" tabindex="-1">
    …
  </main>
</body>
```

```css
.skip-link {
  position: absolute; left: -9999px; top: 0;
  background: #0f172a; color: #fff; padding: .75rem 1rem; z-index: 100;
}
.skip-link:focus { left: 0; }  /* apparaît quand on tabule dessus */
```

> 🎯 Teste au clavier : `Tab` depuis le haut de la page → le skip link apparaît → `Entrée` → tu arrives au contenu sans traverser tout le menu.

---

## 44. ARIA : rôles

**ARIA** (Accessible Rich Internet Applications) ajoute des informations aux technologies d'assistance. **Règle n°1 : pas d'ARIA si le HTML natif suffit.**

```html
<!-- ❌ ARIA inutile -->
<div role="button" tabindex="0">Envoyer</div>
<!-- ✅ HTML natif (clavier + lecteur d'écran gratuits) -->
<button type="submit">Envoyer</button>
```

### Quand ARIA est légitime

| Situation | Solution ARIA |
|-----------|---------------|
| Menu hamburger custom | `aria-expanded`, `aria-controls` |
| Onglets custom | `role="tablist"`, `role="tab"`, `aria-selected` |
| Message d'erreur dynamique | `role="alert"` |
| Barre de progression custom | `role="progressbar"` + `aria-valuenow` |
| Zone qui se met à jour (dashboard) | `aria-live="polite"` |

```html
<!-- Bouton qui déplie un panneau -->
<button aria-expanded="false" aria-controls="panneau-filtres" id="btn-filtres">
  Filtres
</button>
<div id="panneau-filtres" hidden>
  …
</div>
```

```html
<!-- Zone live : le lecteur d'écran annonce les mises à jour -->
<div aria-live="polite" id="statut-ups">
  UPS-01 : en ligne — charge 42 %
</div>
```

> ⚠️ `aria-live="assertive"` interrompt l'utilisateur : réserve-le aux **alertes critiques** (défaut onduleur !), pas aux mises à jour de routine.

---

## 45. ARIA : états et propriétés essentiels

| Attribut | Rôle | Exemple |
|----------|------|---------|
| `aria-label` | Nom accessible quand il n'y a pas de texte visible | `<button aria-label="Fermer">×</button>` |
| `aria-labelledby` | Référence l'élément qui sert de label | `aria-labelledby="titre-dialog"` |
| `aria-describedby` | Description complémentaire | Champ + aide |
| `aria-hidden="true"` | Cache aux lecteurs d'écran (icône décorative) | `<svg aria-hidden="true">…` |
| `aria-expanded` | État déplié/replié | Accordéon, menu |
| `aria-current="page"` | Page active dans la navigation | |
| `aria-invalid="true"` | Champ en erreur | + message lié |
| `aria-disabled="true"` | Désactivé mais focusable (vs `disabled`) | |

```html
<!-- Navigation : page courante -->
<nav aria-label="Principale">
  <ul>
    <li><a href="/" aria-current="page">Accueil</a></li>
    <li><a href="/tickets">Tickets</a></li>
  </ul>
</nav>

<!-- Champ en erreur lié à son message -->
<label for="mail2">E-mail</label>
<input type="email" id="mail2" aria-invalid="true" aria-describedby="err-mail2">
<p id="err-mail2" role="alert">Format d'e-mail invalide.</p>

<!-- Icône décorative dans un bouton : cachée aux lecteurs d'écran -->
<button type="submit">
  <svg aria-hidden="true" focusable="false">…</svg>
  Envoyer
</button>
```

> ⚠️ `aria-hidden="true"` sur un élément **focusable** = catastrophe (focus invisible). Ne le mets que sur du non-interactif.

---

## 46. Navigation clavier et focus

**Test de base :** débranche ta souris. Tu dois pouvoir tout faire au `Tab` / `Maj+Tab` / `Entrée` / `Espace` / flèches.

### Ordre de tabulation

- Suit l'**ordre du DOM** : structure ton HTML dans l'ordre logique (pas de positionnement CSS qui contredit l'ordre visuel).
- `tabindex="0"` : rend focusable (dans l'ordre naturel).
- `tabindex="-1"` : focusable **en JS** uniquement (utile pour déplacer le focus vers un message d'erreur).
- `tabindex > 0` : **interdit** — casse l'ordre naturel.

### Focus visible (obligatoire)

```css
/* Style de focus personnalisé ET visible */
:focus-visible {
  outline: 3px solid #2563eb;
  outline-offset: 2px;
}

/* ❌ JAMAIS ça sans remplacement */
/* :focus { outline: none; } */
```

### Piège clavier

Si tu ouvres une **modale**, le focus doit rester dedans tant qu'elle est ouverte (focus trap) et revenir au bouton d'origine à la fermeture. Sans JS robuste, préfère `<dialog>` natif :

```html
<dialog id="dlg">
  <h2>Confirmer l'arrêt</h2>
  <p>Couper l'onduleur UPS-02 ?</p>
  <form method="dialog">
    <button value="non">Annuler</button>
    <button value="oui">Confirmer</button>
  </form>
</dialog>
<button onclick="document.getElementById('dlg').showModal()">Arrêter l'UPS</button>
```

> 🎯 `<dialog>` gère nativement : focus trap, `Échap` pour fermer, fond assombri (`::backdrop`). Zéro JS à écrire.

---

## 47. Contrastes et couleurs

### Ratios WCAG

| Niveau | Texte normal | Grand texte (≥ 24 px ou ≥ 18,66 px gras) |
|--------|--------------|------------------------------------------|
| **AA** (minimum) | 4,5:1 | 3:1 |
| **AAA** (renforcé) | 7:1 | 4,5:1 |

**Couples qui passent / qui cassent (fond blanc) :**

| Couleur | Ratio | Verdict |
|---------|-------|---------|
| `#000000` (noir) | 21:1 | ✅ AAA |
| `#374151` (gris foncé) | 10,9:1 | ✅ AAA |
| `#6b7280` (gris moyen) | 4,8:1 | ✅ AA |
| `#9ca3af` (gris clair) | 2,8:1 | ❌ même pas AA |
| `#dc2626` (rouge) | 4,8:1 | ✅ AA |
| `#16a34a` (vert) | 3,7:1 | ❌ pour texte normal |

### Règles

- [ ] Vérifie avec un outil (WebAIM Contrast Checker, extension axe).
- [ ] **Ne jamais coder l'info par la seule couleur** : un statut « défaut » rouge doit aussi avoir un texte/icône/forme (« ⚠ Défaut »).
- [ ] Les **placeholder** ont souvent un contraste insuffisant : ne pas y mettre d'info critique.
- [ ] Teste en **niveaux de gris** : si l'info disparaît, elle dépendait trop de la couleur.

```html
<!-- ❌ Couleur seule -->
<span style="color: red">●</span>
<!-- ✅ Couleur + texte + forme -->
<span class="statut statut-defaut"><span aria-hidden="true">●</span> Défaut</span>
```

---

## 48. Formulaires accessibles : la checklist

- [ ] Chaque champ a un `<label>` visible (pas de placeholder seul).
- [ ] Champs obligatoires **indiqués dans le label** : `Nom *` + mention `* Champs obligatoires` en haut.
- [ ] `fieldset` + `legend` pour les groupes (§25).
- [ ] Messages d'erreur : liés au champ (`aria-describedby` + `role="alert"`), **explicites** (« La date doit être après aujourd'hui » pas « Erreur »).
- [ ] Le focus va au **premier champ en erreur** après soumission.
- [ ] `autocomplete` renseigné (nom, email, tel…) : aide tout le monde, crucial sur mobile.

```html
<form action="/api/tickets" method="post">
  <p>* Champs obligatoires</p>

  <div>
    <label for="nom3">Nom *</label>
    <input type="text" id="nom3" name="nom" required autocomplete="name">
  </div>

  <div>
    <label for="mail3">E-mail professionnel *</label>
    <input type="email" id="mail3" name="email" required autocomplete="email"
           aria-describedby="aide-mail3">
    <p id="aide-mail3">Nous enverrons la confirmation à cette adresse.</p>
  </div>

  <button type="submit">Créer le ticket</button>
</form>
```

**Valeurs `autocomplete` courantes :** `name`, `given-name`, `family-name`, `email`, `tel`, `organization`, `street-address`, `postal-code`, `country`, `current-password`, `new-password`.

---

