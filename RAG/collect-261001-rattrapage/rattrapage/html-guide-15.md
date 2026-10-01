---
id: collect-261001-rattrapage/rattrapage/html-guide-15
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26", "2026-11-02"]
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [3347, 3580]
sha256: 45a1391eec6397611c70f22b2370678eab46d80ecabcb7717120523dadfa4c71
---

# Guide HTML & CSS — Le manuel complet

```css
/* css/ticket.css — extraits */
.form-page { max-width: 640px; margin: 2rem auto; padding: 0 1rem; }
fieldset {
  border: 1px solid #e2e8f0; border-radius: .75rem;
  padding: 1.25rem; margin-bottom: 1.5rem;
}
legend { font-weight: 700; padding: 0 .5rem; }
.champ { margin-bottom: 1rem; }
.champ label { display: block; font-weight: 600; margin-bottom: .25rem; }
.champ input, .champ select, .champ textarea {
  width: 100%; padding: .6rem .75rem;
  border: 1px solid #cbd5e1; border-radius: .5rem;
}
.aide { font-size: .875rem; color: #64748b; margin-top: .25rem; }
.choix { display: inline-flex; align-items: center; gap: .4rem; margin-right: 1rem; }

/* Feedback validation (§26) */
input:required, select:required, textarea:required { border-left: 3px solid #f59e0b; }
input:invalid:not(:placeholder-shown),
textarea:invalid:not(:placeholder-shown) { border-color: #dc2626; }
input:valid:not(:placeholder-shown) { border-color: #16a34a; }

.actions { display: flex; gap: 1rem; justify-content: flex-end; }
button { padding: .75rem 1.25rem; border-radius: .5rem; border: 1px solid #cbd5e1; cursor: pointer; }
button.primaire { background: #2563eb; color: #fff; border-color: #2563eb; }
button:disabled { opacity: .5; cursor: not-allowed; }
:focus-visible { outline: 3px solid #2563eb; outline-offset: 2px; }
```

> 💡 Points forts : fieldsets légendés, `autocomplete`, `pattern` + `title`, feedback visuel `:valid`/`:invalid`, boutons typés, `enctype` pour l'upload, `min` sur la date.

---

## 82. Cas pratique 3 : page de documentation interne

```html
<article class="doc">
  <header>
    <!-- Fil d'Ariane : nav secondaire labellisée (§44) -->
    <nav aria-label="Fil d'Ariane">
      <ol class="ariane">
        <li><a href="/docs">Documentation</a></li>
        <li><a href="/docs/ups">Onduleurs</a></li>
        <li aria-current="page">Maintenance préventive</li>
      </ol>
    </nav>
    <h1>Maintenance préventive des onduleurs</h1>
    <p class="meta">Mis à jour le <time datetime="2026-09-26">26/09/2026</time>
      · Service Systèmes & Énergies</p>
  </header>

  <!-- Sommaire auto : ancres vers les sections -->
  <nav aria-label="Sommaire">
    <h2>Sommaire</h2>
    <ul>
      <li><a href="#periodicites">Périodicités</a></li>
      <li><a href="#valeurs">Valeurs de référence</a></li>
      <li><a href="#procedure">Procédure</a></li>
    </ul>
  </nav>

  <section id="periodicites">
    <h2>Périodicités</h2>
    <table>
      <caption>Fréquence des contrôles</caption>
      <thead>
        <tr><th scope="col">Contrôle</th><th scope="col">Fréquence</th></tr>
      </thead>
      <tbody>
        <tr><td>Inspection visuelle</td><td>Trimestrielle</td></tr>
        <tr><td>Test d'impédance batteries</td><td>Annuelle</td></tr>
        <tr><td>Thermographie</td><td>Annuelle</td></tr>
      </tbody>
    </table>
  </section>

  <section id="valeurs">
    <h2>Valeurs de référence</h2>
    <dl>
      <dt>Tension de floating / bloc 12 V</dt>
      <dd>13,5 – 13,6 V à 20 °C</dd>
      <dt>Charge maximale conseillée</dt>
      <dd>80 % de la puissance nominale</dd>
    </dl>
  </section>

  <section id="procedure">
    <h2>Procédure</h2>
    <ol>
      <li>Consigner l'armoire (voir <a href="/docs/consignation">procédure de consignation</a>).</li>
      <li>Mesurer la tension du bus DC.</li>
      <li>…
        <!-- Accordéon natif pour le détail (§33) -->
        <details>
          <summary>Détail de la mesure du bus DC</summary>
          <p>Utiliser un multimètre CAT III…</p>
        </details>
      </li>
    </ol>
  </section>

  <footer>
    <p>Document interne — <a href="mailto:energie@entreprise.fr">signaler une erreur</a></p>
  </footer>
</article>
```

> 💡 Structure idéale d'une doc : fil d'Ariane + `aria-current`, sommaire à ancres, sections titrées, `<dl>` pour les valeurs, `<details>` pour les approfondissements, JSON-LD `TechArticle` (§39) dans le head pour la doc publique.

---

## 83. Cas pratique 4 : carte équipement réutilisable

Composant **BEM** (§75) : la même carte sert sur le dashboard, la recherche, les favoris.

```html
<article class="eq-card eq-card--ko">
  <div class="eq-card__entete">
    <h3 class="eq-card__titre">UPS-02</h3>
    <span class="eq-card__statut"><span aria-hidden="true">●</span> Défaut batterie</span>
  </div>
  <dl class="eq-card__specs">
    <div>
      <dt>Puissance</dt><dd>60 kVA</dd>
    </div>
    <div>
      <dt>Local</dt><dd>Datacenter</dd>
    </div>
    <div>
      <dt>Charge</dt><dd>—</dd>
    </div>
    <div>
      <dt>Maintenance</dt><dd><time datetime="2026-11-02">02/11/2026</time></dd>
    </div>
  </dl>
  <div class="eq-card__actions">
    <a class="btn" href="/tickets/nouveau?eq=UPS-02">Créer un ticket</a>
    <a class="btn btn--secondaire" href="/docs/ups-02">Documentation</a>
  </div>
</article>
```

```css
.eq-card {
  --accent: #16a34a; /* variante par défaut */
  background: #fff; border: 1px solid #e2e8f0; border-radius: .75rem;
  border-top: 4px solid var(--accent);
  padding: 1.25rem;
  display: flex; flex-direction: column; gap: 1rem; /* flex vertical §59 */
}
.eq-card--ko   { --accent: #dc2626; }
.eq-card--warn { --accent: #d97706; }

.eq-card__entete { display: flex; justify-content: space-between; align-items: baseline; }
.eq-card__titre { font-size: 1.25rem; }
.eq-card__statut { font-weight: 700; color: var(--accent); white-space: nowrap; }

.eq-card__specs {
  display: grid; grid-template-columns: repeat(2, 1fr); gap: .5rem 1rem;
  margin: 0;
}
.eq-card__specs div { display: contents; } /* dt/dd restent dans la grille */
.eq-card__specs dt { color: #64748b; font-size: .875rem; }
.eq-card__specs dd { margin: 0; font-weight: 600; }

.eq-card__actions { display: flex; gap: .75rem; margin-top: auto; }
.btn {
  display: inline-block; padding: .6rem 1rem; border-radius: .5rem;
  background: #2563eb; color: #fff; text-decoration: none; text-align: center;
  transition: background-color .2s ease;
}
.btn:hover { background: #1d4ed8; }
.btn--secondaire { background: #f1f5f9; color: #1e293b; }
```

> 💡 La variante (`--ko`) ne change qu'une **variable** : le reste suit. C'est la puissance du combo BEM + custom properties.

---

## 84. Checklist de mise en production

Avant de publier une page (même interne), passe cette checklist :

**HTML**
- [ ] Doctype + `lang` + charset + viewport présents.
- [ ] Validateur W3C : zéro erreur.
- [ ] Un seul `<h1>`, hiérarchie sans saut.
- [ ] Toutes les images ont un `alt` (éventuellement vide).
- [ ] Tous les champs ont un `<label>`.
- [ ] Liens explicites, `rel="noopener"` sur les `_blank`.

**Accessibilité**
- [ ] Navigation complète au clavier (test sans souris).
- [ ] Focus visible partout.
- [ ] Contrastes AA vérifiés.
- [ ] Test lecteur d'écran rapide (NVDA gratuit / VoiceOver intégré).
- [ ] `prefers-reduced-motion` géré si animations.

**SEO / metas** (si page publique)
- [ ] `<title>` et `description` uniques et soignés.
- [ ] Open Graph + image 1200×630.
- [ ] JSON-LD validé (Rich Results Test).
- [ ] `canonical` si doublons possibles.

**Performance**
- [ ] Poids total < 1 Mo, images optimisées + `lazy`.
- [ ] Lighthouse ≥ 90 (perf + accessibilité).
- [ ] Test mobile (émulateur + vrai téléphone si possible).

**Contenu**
- [ ] Pas d'info sensible dans le code ou les commentaires.
- [ ] Dates, chiffres, références vérifiés.
- [ ] Page 404 personnalisée si site public.

---

## 85. Pense-bête de poche

### Squelette HTML5
```html
<!DOCTYPE html>
<html lang="fr">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>…</title>
<link rel="stylesheet" href="style.css">
</head>
<body>
</body>
</html>
```

