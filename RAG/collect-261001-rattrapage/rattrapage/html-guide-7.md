---
id: collect-261001-rattrapage/rattrapage/html-guide-7
title: "Guide HTML & CSS — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: ["2026-09-26"]
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/html_guide.md
source_anchor: ""
source_lines: [1388, 1605]
sha256: 0bdccb699b58c6dbc6256e963fe913c3088f40ef51623b63c640627e201455d4
---

# Guide HTML & CSS — Le manuel complet

- [ ] **Un `<title>` unique** par page, ~50–60 caractères, mot-clé en premier : `Tickets onduleurs — Service Systèmes & Énergies` ✅ / `Accueil` ❌
- [ ] **Une `meta description` unique** par page, ~150–160 caractères, incitative.
- [ ] **Un seul `<h1>`**, hiérarchie sans saut (§11).
- [ ] URLs **lisibles** : `/docs/maintenance-ups` ✅ / `/page.php?id=42` ❌
- [ ] Images avec `alt` descriptif + nom de fichier explicite (`schema-tgbt.svg` ✅).
- [ ] **Temps de chargement** < 3 s (poids images, §76).
- [ ] Page **responsive** (Google indexe en mobile-first).
- [ ] Pas de contenu important uniquement en image (le texte d'une image n'est pas lu… sauf via `alt`).

```html
<title>Maintenance des onduleurs — Guide du service énergies</title>
<meta name="description" content="Procédures de maintenance préventive des onduleurs :
  checklists, périodicités, valeurs de référence. Service Systèmes & Énergies.">
```

> 🎯 Pour des pages **internes**, l'essentiel SEO = `noindex` (§35) + structure propre. Pour de la doc publique, applique toute la checklist.

---

## 38. Open Graph et Twitter Cards : beaux partages

Quand on colle ton URL dans Teams, Slack ou LinkedIn, ces metas génèrent la **carte d'aperçu** (titre + image + description).

```html
<!-- Open Graph (Facebook, LinkedIn, Teams, Slack…) -->
<meta property="og:title" content="Guide de maintenance des onduleurs">
<meta property="og:description" content="Checklists, périodicités et valeurs de référence.">
<meta property="og:type" content="article">
<meta property="og:url" content="https://docs.entreprise.fr/guides/maintenance-ups">
<meta property="og:image" content="https://docs.entreprise.fr/img/og-ups.jpg">
<meta property="og:image:width" content="1200">
<meta property="og:image:height" content="630">
<meta property="og:locale" content="fr_FR">

<!-- Twitter Cards -->
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:title" content="Guide de maintenance des onduleurs">
<meta name="twitter:description" content="Checklists, périodicités et valeurs de référence.">
<meta name="twitter:image" content="https://docs.entreprise.fr/img/og-ups.jpg">
```

**Image OG idéale :** 1200×630 px, < 300 Ko, texte lisible en miniature, URLs **absolues** (https://…).

> 💡 Teste le rendu : validateur de cartes Twitter, LinkedIn Post Inspector, ou simplement colle l'URL dans Teams.

---

## 39. Données structurées : JSON-LD

Le **JSON-LD** décrit le contenu aux moteurs dans un vocabulaire standard (schema.org) — placé dans un `<script type="application/ld+json">`.

```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "TechArticle",
  "headline": "Maintenance préventive des onduleurs",
  "description": "Checklists et périodicités de maintenance des UPS.",
  "inLanguage": "fr",
  "author": {
    "@type": "Organization",
    "name": "Service Systèmes & Énergies"
  },
  "datePublished": "2026-09-26"
}
</script>
```

**Autres types utiles :**

| Type schema.org | Usage |
|-----------------|-------|
| `TechArticle` / `Article` | Documentation, guides |
| `FAQPage` | Pages de questions fréquentes (affiche en résultats enrichis) |
| `HowTo` | Procédures pas à pas |
| `BreadcrumbList` | Fil d'Ariane |
| `Organization` | Infos de l'entreprise |

**Exemple FAQ (résultats enrichis Google) :**

```html
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "FAQPage",
  "mainEntity": [{
    "@type": "Question",
    "name": "À quelle fréquence tester les batteries d'un onduleur ?",
    "acceptedAnswer": {
      "@type": "Answer",
      "text": "Test d'impédance annuel et contrôle visuel trimestriel."
    }
  }]
}
</script>
```

> 🎯 Valide ton JSON-LD : Google Rich Results Test. Une virgule oubliée = tout le bloc ignoré.

---

## 40. Sitemap, robots.txt, canonical, hreflang

### robots.txt (à la racine du site)

```text
User-agent: *
Disallow: /admin/
Disallow: /api/
Allow: /

Sitemap: https://docs.entreprise.fr/sitemap.xml
```

### Balise canonical : éviter le contenu dupliqué

```html
<!-- Cette page existe aussi via ?tri=date : on désigne l'URL de référence -->
<link rel="canonical" href="https://docs.entreprise.fr/guides/maintenance-ups">
```

### hreflang : versions multilingues

```html
<link rel="alternate" hreflang="fr" href="https://docs.entreprise.fr/fr/guide">
<link rel="alternate" hreflang="en" href="https://docs.entreprise.fr/en/guide">
<link rel="alternate" hreflang="x-default" href="https://docs.entreprise.fr/fr/guide">
```

### Sitemap XML (extrait)

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://docs.entreprise.fr/guides/maintenance-ups</loc>
    <lastmod>2026-09-26</lastmod>
    <changefreq>monthly</changefreq>
  </url>
</urlset>
```

| Fichier/élément | Rôle |
|-----------------|------|
| `robots.txt` | Dit aux robots quoi explorer |
| `sitemap.xml` | Liste les pages à indexer |
| `canonical` | URL de référence en cas de doublons |
| `hreflang` | Associe les versions linguistiques |

---

## 41. Accessibilité : pourquoi et principes

**L'accessibilité web** = rendre tes pages utilisables par **tous** : aveugles (lecteurs d'écran), malvoyants, daltoniens, personnes à mobilité réduite (clavier seul), etc. En France, le **RGAA** (référentiel public) impose des obligations aux services publics — et c'est une bonne pratique partout.

### Les 4 principes (WCAG)

| Principe | Signification | Exemple |
|----------|---------------|---------|
| **Perceptible** | L'info doit être percevable | `alt` sur images, contrastes suffisants |
| **Utilisable** | Navigable au clavier | Focus visible, pas de piège clavier |
| **Compréhensible** | Lisible et prévisible | Labels explicites, erreurs claires |
| **Robuste** | Interprétable par les techs d'assistance | HTML valide, ARIA correct |

### Ce que ça change concrètement pour toi

- Tes dashboards seront utilisables par un collègue malvoyant (zoom 200 % sans casse).
- Tes formulaires de tickets seront remplissables **sans souris**.
- Bonus : l'accessibilité améliore le **SEO** et la **robustesse** générale.

> 💡 Réflexe : 80 % de l'accessibilité = **HTML sémantique correct** (titres, labels, alt, landmarks). ARIA ne vient qu'en complément (§44–45).

---

## 42. Texte alternatif : l'arbre de décision

Le `alt` est souvent mal écrit. Voici la méthode :

```
L'image est-elle décorative (bordure, illustration d'ambiance) ?
├── OUI → alt=""
└── NON → L'image porte-t-elle de l'information ?
    ├── NON (redondante avec le texte adjacent) → alt=""
    └── OUI → Décris la FONCTION, pas l'apparence
```

**Exemples concrets :**

```html
<!-- ❌ -->
<img src="graph.png" alt="image d'un graphique">
<!-- ✅ : ce que le graphique APPREND -->
<img src="graph.png" alt="Charge de l'onduleur : pic à 78 % à 14 h, stable sinon">

<!-- Image décorative -->
<img src="img/motif-fond.svg" alt="">

<!-- Image-lien : décris l'ACTION -->
<a href="rapport.pdf">
  <img src="img/pdf-icon.svg" alt="Télécharger le rapport d'intervention (PDF)">
</a>

<!-- Schéma complexe : alt court + description longue -->
<img src="img/unifilaire.svg" alt="Schéma unifilaire du TGBT"
     aria-describedby="desc-unifilaire">
<p id="desc-unifilaire">Le schéma montre : arrivée EDF 400 V… [description détaillée]</p>
```

> ⚠️ `alt` trop long (> 125 caractères) : coupe en `alt` court + `aria-describedby`. Les vieux lecteurs d'écran tronquent.

---

## 43. Landmarks et skip links

### Landmarks : les repères de navigation

Les balises sémantiques (§32) créent des **landmarks** que les lecteurs d'écran listent :

