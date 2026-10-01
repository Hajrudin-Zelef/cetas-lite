---
id: collect-261001-rattrapage/rattrapage/html-guide-14
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [3131, 3346]
sha256: cdfdd53961b62732d6f14af753b2b504a15e3ded0b4847af31a93e4cfe10b7d4
---

# Guide HTML & CSS — Le manuel complet

      <!-- Tableau accessible (§18-19) dans conteneur scrollable -->
      <section aria-labelledby="t-equipements">
        <h2 id="t-equipements">Détail par équipement</h2>
        <div class="table-scroll">
          <table>
            <caption>État temps réel du parc</caption>
            <thead>
              <tr>
                <th scope="col">Équipement</th>
                <th scope="col">Type</th>
                <th scope="col">Charge</th>
                <th scope="col">Statut</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th scope="row">UPS-01</th>
                <td>Onduleur 30 kVA</td>
                <td>42 %</td>
                <!-- Statut : couleur + TEXTE (§47) -->
                <td><span class="statut statut--ok"><span aria-hidden="true">●</span> En ligne</span></td>
              </tr>
              <tr>
                <th scope="row">UPS-02</th>
                <td>Onduleur 60 kVA</td>
                <td>—</td>
                <td><span class="statut statut--ko"><span aria-hidden="true">●</span> Défaut batterie</span></td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </main>

    <footer class="pied">
      <p>Service Systèmes & Énergies — usage interne</p>
    </footer>
  </div>
</body>
</html>
```

```css
/* css/dashboard.css — extraits clés */
*, *::before, *::after { box-sizing: border-box; }
* { margin: 0; }

:root {
  --fond: #f8fafc; --carte: #fff; --texte: #1e293b;
  --ok: #16a34a; --warn: #d97706; --ko: #dc2626;
  --rayon: .75rem;
}

/* Layout : grid areas (§62) */
.page {
  display: grid;
  grid-template-rows: auto 1fr auto;
  min-height: 100vh;
}
.topbar {
  display: flex; align-items: center; gap: 1.5rem;  /* flexbox §59 */
  padding: 1rem 1.5rem; background: #0f172a; color: #fff;
}
.topbar ul { display: flex; gap: 1rem; list-style: none; padding: 0; }
.topbar a { color: #e2e8f0; }
.topbar a[aria-current="page"] { color: #fff; font-weight: 700; }
.horloge { margin-left: auto; color: #94a3b8; font-size: .875rem; }

.contenu { padding: 1.5rem; max-width: 1200px; width: 100%; margin: 0 auto; }
.contenu h1 { margin-bottom: 1rem; }

/* KPIs : grid auto-responsive (§61) */
.kpis {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 1rem; list-style: none; padding: 0; margin-bottom: 2rem;
}
.kpi {
  background: var(--carte); border-radius: var(--rayon);
  padding: 1rem; border-left: 5px solid var(--ok);
}
.kpi--warn { border-color: var(--warn); }
.kpi--ko   { border-color: var(--ko); }
.kpi__valeur { font-size: 2rem; font-weight: 800; }
.kpi__label  { color: #64748b; font-size: .875rem; }

/* Statuts : pastille + texte (jamais couleur seule) */
.statut { font-weight: 600; white-space: nowrap; }
.statut--ok   { color: var(--ok); }
.statut--warn { color: var(--warn); }
.statut--ko   { color: var(--ko); }

/* Tableau */
.table-scroll { overflow-x: auto; }
table { border-collapse: collapse; width: 100%; min-width: 640px; background: var(--carte); }
caption { text-align: left; font-weight: 700; padding: .5rem 0; }
th, td { padding: .75rem; border-bottom: 1px solid #e2e8f0; text-align: left; }
tbody tr:nth-child(even) { background: #f8fafc; }  /* zébré §69 */
tbody tr:hover { background: #e0f2fe; }

.pied { text-align: center; padding: 1rem; color: #64748b; font-size: .875rem; }

/* Focus visible partout (§46) */
:focus-visible { outline: 3px solid #2563eb; outline-offset: 2px; }

/* Skip link (§43) */
.skip-link { position: absolute; left: -9999px; background: #fff; padding: .5rem 1rem; z-index: 300; }
.skip-link:focus { left: 1rem; top: 1rem; }
```

> 💡 Ce dashboard coche : sémantique, skip link, `aria-current`, zone live prête, statuts texte+couleur, tableau scrollable, grid responsive sans breakpoint, focus visible.

---

## 81. Cas pratique 2 : formulaire de ticket (commenté)

```html
<!DOCTYPE html>
<html lang="fr">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="robots" content="noindex, nofollow">
  <title>Nouveau ticket — Service Systèmes & Énergies</title>
  <link rel="stylesheet" href="css/ticket.css">
</head>
<body>
  <main class="form-page">
    <h1>Créer un ticket d'intervention</h1>
    <p>* Champs obligatoires</p>

    <!-- novalidate retiré : on garde la validation native (§26) -->
    <form action="/api/tickets" method="post" enctype="multipart/form-data">

      <fieldset>
        <legend>Demandeur</legend>
        <div class="champ">
          <label for="nom">Nom *</label>
          <input type="text" id="nom" name="nom" required autocomplete="name"
                 minlength="2" maxlength="60">
        </div>
        <div class="champ">
          <label for="email">E-mail professionnel *</label>
          <input type="email" id="email" name="email" required autocomplete="email"
                 aria-describedby="aide-email">
          <p class="aide" id="aide-email">La confirmation sera envoyée à cette adresse.</p>
        </div>
        <div class="champ">
          <label for="tel">Téléphone (10 chiffres)</label>
          <input type="tel" id="tel" name="tel" inputmode="numeric"
                 pattern="[0-9]{10}" title="10 chiffres sans espaces">
        </div>
      </fieldset>

      <fieldset>
        <legend>Équipement concerné</legend>
        <div class="champ">
          <label for="famille">Famille *</label>
          <select id="famille" name="famille" required>
            <option value="">— Choisir —</option>
            <optgroup label="Énergie">
              <option value="ups">Onduleur</option>
              <option value="batterie">Batteries</option>
              <option value="ge">Groupe électrogène</option>
              <option value="tgbt">TGBT / armoire</option>
            </optgroup>
            <optgroup label="Climatisation">
              <option value="clim">Climatiseur</option>
            </optgroup>
          </select>
        </div>
        <div class="champ">
          <label for="ref">Référence (format ABC-12) *</label>
          <input type="text" id="ref" name="ref" required
                 pattern="[A-Z]{3}-[0-9]{2}" title="Ex : UPS-02">
        </div>
        <div class="champ">
          <label for="date">Date souhaitée *</label>
          <!-- min = demain : pas d'intervention dans le passé -->
          <input type="date" id="date" name="date" required min="2026-09-27">
        </div>
      </fieldset>

      <fieldset>
        <legend>Criticité</legend>
        <!-- radios : même name = choix unique (§21) -->
        <label class="choix"><input type="radio" name="crit" value="basse"> Basse</label>
        <label class="choix"><input type="radio" name="crit" value="normale" checked> Normale</label>
        <label class="choix"><input type="radio" name="crit" value="haute"> Haute</label>
        <label class="choix"><input type="radio" name="crit" value="critique"> Critique</label>
      </fieldset>

      <fieldset>
        <legend>Description</legend>
        <div class="champ">
          <label for="desc">Décrivez le problème *</label>
          <textarea id="desc" name="desc" rows="5" required minlength="20"
                    placeholder="Ex : alarme batterie depuis hier 14 h, tension floating à 12,1 V…"></textarea>
        </div>
        <div class="champ">
          <label for="pj">Photo du défaut (JPG/PNG, optionnel)</label>
          <input type="file" id="pj" name="pj" accept="image/jpeg,image/png">
        </div>
      </fieldset>

      <div class="actions">
        <!-- type="button" : ne soumet PAS (§23) -->
        <button type="button" id="btn-brouillon">Enregistrer en brouillon</button>
        <button type="submit" class="primaire">Envoyer le ticket</button>
      </div>
    </form>
  </main>
</body>
</html>
```

