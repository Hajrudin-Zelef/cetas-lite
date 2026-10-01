---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-1
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sol"]
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [1, 177]
sha256: 9fd23407e76ae53982b3f4c7db44de57d009cc7e96b813b2fe37d5adc916c236
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

> Tout sur les copieurs/multifonctions Canon imageRUNNER / imageRUNNER
> ADVANCE : fonctionnement, consommables, maintenance préventive,
> dictionnaire des codes erreur, bourrages, qualité d'image, mode
> service, procédures de remplacement, réseau, firmware, dépannage.
>
> **Orienté technicien de maintenance** : c'est le cœur du métier qui
> est détaillé ici. Rédigé à partir des connaissances du domaine
> (manuels de service Canon, 2025-2026). Les codes et procédures
> varient selon les séries : **vérifiez toujours le manuel de service
> exact de votre modèle** avant une intervention lourde.

---

## 1. Les gammes Canon — s'y retrouver

### 1.1 imageRUNNER ADVANCE (la gamme pro actuelle)
- **DX C3800 / C5800 series** (A3 couleur, 2023+) : remplace les C5500.
  Vitesses 30–60 ppm. La plus vendue en entreprise.
- **DX 6800 / 4800 series** (A3 N&B) : remplace les 6500/4500.
- **DX C2700 series** (A3 couleur compacte).
- **C3300 series** (A4 couleur, remplace C3500).
- Anciennes : **iR-ADV C5200/C5500**, **iR-ADV 4200/4500/6500**,
  **iR 2520/2530/2545** (très répandues en Afrique, robustes).

### 1.2 Autres gammes
- **imagePRESS** : production (C10000, V1000) — hors sujet ici sauf
  mention.
- **imageFORMULA** : scanners (pas des copieurs).
- **MAXIFY / PIXMA** : jet d'encre, pas concernées par ce guide.
- **imageRUNNER 1600/2600 series** (A4 N&B récents, remplace 1435).

### 1.3 Lire une référence
- **A3/A4** : format max. **C** = couleur, pas de C = N&B.
- **Chiffres** : vitesse approximative (C5860 ≈ 60 ppm).
- **DX** = génération actuelle (2021+), avec sécurité renforcée
  (McAfee embarqué, TPM).

---

## 2. Principe de fonctionnement — l'électrophotographie

Comprendre le cycle = diagnostiquer vite. 7 étapes :

```
1. CHARGE      → le tambour (OPC) est chargé uniformément (-)
2. EXPOSITION  → le laser « dessine » l'image (zones déchargées)
3. DÉVELOPPEMENT → le toner (-) est attiré sur les zones exposées
4. TRANSFERT   → le toner passe du tambour au papier (+)
5. SÉPARATION  → le papier se décolle du tambour
6. FUSION      → rouleaux chauffants (180-200 °C) fixent le toner
7. NETTOYAGE   → la raclette enlève le toner résiduel du tambour
```

- **Couleur** : 4 tambours (CMYK) + **courroie de transfert (ITB)**
  qui assemble les 4 couches avant transfert sur papier.
- Chaque étape = un composant = une famille de pannes :
  - Fond gris → charge ou développeur.
  - Image pâle → laser, développeur, transfert.
  - Lignes verticales → tambour ou raclette rayés.
  - Taches répétitives → mesurez l'intervalle (voir §7).
  - Papier qui s'enroule → fusion (température/pression).

---

## 3. Consommables et pièces d'usure

### 3.1 Durées de vie typiques (iR-ADV A3, à adapter)

| Pièce | Durée indicative | Signes de fin de vie |
|---|---|---|
| Toner | 15 000–40 000 pages (selon cartouche) | Message « toner bas », pâleur |
| Tambour (drum unit) | 100 000–250 000 pages | Lignes, fond gris, points répétitifs |
| Développeur (developer) | 200 000–500 000 pages | Densité instable, fond |
| Unité de fusion (fixing) | 200 000–500 000 pages | Bourrages sortie, toner qui s'efface |
| Courroie transfert (ITB) | 200 000–500 000 pages | Couleurs décalées, bandes |
| Rouleaux d'entraînement | 100 000–200 000 pages | Multi-prises, bourrages pickup |
| Raclette de nettoyage | avec le tambour | Traînées |

- Ces chiffres sont des **ordres de grandeur** : papier épais, humidité
  et poussière (contexte Afrique) divisent par 1,5 à 2.
- **Compteurs** : COPIER > COUNTER > PARTS (mode service) donne la
  vie restante estimée par pièce.

### 3.2 Toner : les règles d'or
- **Toner d'origine Canon uniquement** en maintenance pro : le
  compatible use le développeur et le tambour (granulométrie
  différente), et annule les arguments de garantie.
- Ne secouez jamais violemment une cartouche (poussière de toner =
  toxique à inhaler, portez un **masque**).
- Stockage : à l'horizontale, à l'ombre, < 35 °C. Le toner cuit
  s'agglutine.
- Après remplacement : la machine fait un cycle d'agitation
  (ne pas interrompre).

### 3.3 Le toner usagé (waste toner)
- Bac récupérateur à vider/remplacer quand plein (**E013** si débordé).
- Ne jamais réutiliser le toner usagé (chargé, contaminé).
- Le bac plein non vidé = toner qui remonte dans le développeur =
  fond gris généralisé.

---

## 4. Maintenance préventive — plannings et checklists

### 4.1 Fréquences recommandées

| Fréquence | Actions |
|---|---|
| **Chaque visite** | Compteurs, codes erreur historiques, test copie/scan, état consommables |
| **Tous les 50 000 pages** | Nettoyage optiques, vitre d'exposition, rouleaux, bac toner usagé |
| **Tous les 100 000 pages** | Rouleaux d'entraînement (si usure), filtres, ventilateurs |
| **Tous les 200 000 pages** | Tambour, développeur, ITB, unité de fusion (selon compteurs PARTS) |
| **Annuel** | Firmware, sauvegarde carnet d'adresses, test disques, dépoussiérage complet |

### 4.2 Checklist de visite (à cocher)

```
□ Relever compteurs (total, couleur, N&B) → COUNTER
□ Historique erreurs → mode service > DISPLAY > ERR / JAM
□ Test : 1 copie vitre + 1 copie ADF + 1 scan + 1 impression réseau
□ Qualité : mire de test (fond, lignes, densité, couleurs)
□ Consommables : % restants (toner, bac usagé)
□ Compteurs de pièces : PARTS (drum, fixing, ITB…)
□ Nettoyage : vitre expo, bande blanche ADF, optiques, corona
□ Rouleaux pickup : visuel (lisse = à changer)
□ Ventilateurs : bruit/poussière
□ Firmware : version vs dernière connue
□ Carnet d'adresses : sauvegarde (si modifs)
□ Client : formation express si bourrages récurrents (souvent = mauvais papier)
```

### 4.3 Le papier — 50 % des appels
- **80 g/m²** standard ; papier humide (saison des pluies !) =
  bourrages + ondulation. Stockez le papier **dans son emballage**,
  à l'horizonté, jamais au sol.
- Papier recyclé bas de gamme = poussière = encrassement optiques et
  rouleaux. Imposez une **marque de papier validée** au client.
- Ventilez la rame avant de charger (évite les multi-prises).
- Guides du bac **ajustés** au format (guides lâches = papier de travers
  = bourrages 0108/0208).

### 4.4 Environnement
- Température 15–30 °C, humidité 20–70 %. En Afrique : **clim ou
  local ventilé** — la chaleur cuit le toner et vieillit les
  électroniques.
- Poussière (harmattan) : dépoussiérage à l'air comprimé **sec**
  (jamais d'air humide de compresseur sans filtre), masque obligatoire.
- Onduleur sur l'alimentation : les micro-coupures corrompent le
  disque dur (**E602**) et tuent les cartes.

---

## 5. Dictionnaire des codes erreur (E-codes)

> Les codes sont au format `EXXX-YYYY` (code + détail). Le détail
> précise le sous-ensemble. Procédure standard : noter le code **complet**,
> éteindre/rallumer, si retour → mode service > COPIER > FUNCTION >
> CLEAR > ERR, si persistant → intervention.
>
> ⚠️ Les codes **E000/E001/E002/E003** (fusion) et **E602** (disque)
> sont les plus fréquents en maintenance. Ils sont détaillés en premier.

### 5.1 Bloc FUSION (les plus critiques)

**E000-0001 — Montée anormale en température (fixing)**
- Cause : thermistance sale/défectueuse, carte DC, surchauffe réelle.
- Action : laisser refroidir 15 min, CLEAR > ERR. Si retour :
  vérifier thermistances (résistance), connecteurs, puis remplacer
  l'unité de fusion. **Ne jamais shunter la sécurité thermique.**

**E001-0001 / 0002 — Température excessive**
- Idem E000, seuil dépassé. Vérifier le thermostat (coupure de
  sécurité) : s'il a sauté, la cause est réelle (thermistance ou
  triac de la carte d'alimentation).

