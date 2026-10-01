---
id: collect-261001-rattrapage/rattrapage/energie-solaire-guide-2
title: "Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sol"]
source: docs/RAG/collect-261001-rattrapage/energie_solaire_guide.md
source_anchor: ""
source_lines: [162, 349]
sha256: 2f1804770b7b8fdf66e7263aa63dd1b65f35b1da03c0c35b1b9578b7cfa50c93
---

# Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)

- **Verdict 2025-2026 : LiFePO4 partout** sauf tout petit budget
  immédiat. Le plomb en Afrique (35–40 °C) meurt en 2–3 ans :
  chaque +10 °C au-dessus de 25 °C **divise par 2** sa durée de vie.
- **BMS** (Battery Management System) : obligatoire en lithium —
  il protège contre surcharge, décharge profonde, court-circuit,
  déséquilibre des cellules. Ne jamais acheter de lithium sans BMS.

### 4.2 Dimensionner le parc batteries

**Formule** :
```
Énergie à stocker (Wh) = Consommation nocturne (Wh) × jours d'autonomie
Capacité utile (Wh) = Énergie / profondeur de décharge max
Capacité nominale (Wh) = Capacité utile / rendement
```

- **Jours d'autonomie** : 1–2 jours en hybride (le réseau/groupe
  prend le relais), 2–3 jours en off-grid (saison des pluies).
- **Tension du parc** : 12 V (< 1 kW), 24 V (1–3 kW), **48 V (> 3 kW,
  le standard)**. Plus la tension est haute, plus les courants sont
  faibles (câbles moins gros, moins de pertes).
- **Exemple** : besoin nocturne 5 kWh, 2 jours d'autonomie, LiFePO4
  (DoD 80 %) :
  `5 × 2 / 0,8 = 12,5 kWh utiles` → parc **48 V 260 Ah** (≈ 12,5 kWh).

### 4.3 Règles d'or
- Jamais de mélange : même marque, même modèle, même âge, même lot
  si possible.
- Local **ventilé**, à l'ombre, 15–30 °C idéal. Jamais en plein soleil.
- Câbles de même longueur entre batteries en parallèle (équilibrage
  des courants).
- Coupe-batterie + fusible **au plus près** des bornes (le
  court-circuit d'un parc lithium = soudure instantanée).

---

## 5. Dimensionnement complet — méthode pas à pas

### Étape 1 : bilan de consommation
Listez chaque appareil : puissance (W) × heures/jour = Wh/jour.

| Appareil | Puissance | h/j | Wh/j |
|---|---|---|---|
| 10 LED | 90 W | 6 | 540 |
| Frigo | 150 W | 8 (cyclé) | 1 200 |
| Congélateur | 200 W | 8 (cyclé) | 1 600 |
| Ventilateurs ×3 | 180 W | 8 | 1 440 |
| TV + box | 120 W | 5 | 600 |
| Chargeurs/PC | 200 W | 6 | 1 200 |
| **Total** | | | **≈ 6 580 Wh/j** |

### Étape 2 : puissance crête panneaux
```
P(Wc) = Conso quotidienne (Wh) / (ensoleillement × rendement système)
P = 6 580 / (5 × 0,75) ≈ 1 755 Wc → 4 panneaux de 550 Wc = 2 200 Wc
```
(rendement système 0,70–0,80 : onduleur, câbles, température, poussière).

### Étape 3 : batteries (hybride, 1,5 jour d'autonomie, nuit = 60 % conso)
```
Nuit ≈ 4 kWh × 1,5 jour / 0,8 (DoD) ≈ 7,5 kWh → 48 V 160 Ah LiFePO4
```

### Étape 4 : onduleur et régulateur
- Pic simultané estimé : ~2,5 kW → onduleur hybride **5 kW** (marge).
- Régulateur : 2 200 Wc / 48 V ≈ 46 A → **MPPT 60 A**.

### Étape 5 : synthèse exemple « villa »
- 4 × 550 Wc TOPCon, onduleur hybride 5 kW 48 V, 1 × batterie
  LiFePO4 48 V 160 Ah (7,7 kWh), MPPT 60 A intégré ou externe.
- Production ≈ 8 kWh/jour en saison sèche, ~5 en saison des pluies.

---

## 6. Trois exemples chiffrés

### 6.1 Studio / petit appartement (secours délestages)
- **Besoin** : éclairage, ventilateur, TV, chargeurs, frigo ≈ 2,5 kWh/j.
- **Kit** : 2 × 550 Wc, onduleur hybride 3 kW 24 V, batterie
  LiFePO4 24 V 100 Ah (2,4 kWh).
- **Ordre de prix** : 1 500–2 500 € / 1–1,7 M FCFA (matériel).

### 6.2 Villa familiale (exemple §5)
- **Besoin** : ~6,5 kWh/j.
- **Kit** : 4 × 550 Wc, hybride 5 kW 48 V, LiFePO4 48 V 160 Ah.
- **Ordre de prix** : 4 000–6 000 € / 2,6–4 M FCFA.

### 6.3 Bureau / site télécom 10 kW
- **Besoin** : ~40 kWh/j (clim, serveurs, onduleurs).
- **Kit** : 24 × 550 Wc (13,2 kWc), 2 × hybrides 8 kW en parallèle,
  4 × LiFePO4 48 V 200 Ah (38 kWh), structure au sol.
- **Ordre de prix** : 18 000–28 000 € / 12–18 M FCFA.
- Envisager un **groupe électrogène** en secours + un contrat de
  maintenance (site critique).

> Prix indicatifs 2025-2026, matériel seul, hors main-d'œuvre et
> transport. Les prix ont fortement baissé depuis 2022 (panneaux
> −40 %, lithium −50 %).

---

## 7. Installation — pas à pas

### 7.1 Étude préalable
1. **Relevé d'ombres** (9h/12h/16h, deux saisons).
2. **Toiture** : orientation, pente, état (un toit à refaire dans
   5 ans = refaire l'installation ; renforcez si tôle fine).
3. **Structure** : point d'ancrage, cheminement des câbles DC,
   local technique (batteries/onduleur) ventilé et sécurisé.
4. **Compteur et protections AC** existants.

### 7.2 Structure de montage
- **Toiture tôle** : rails aluminium + crochets/vis avec
  **étanchéité** (joints EPDM) à chaque perçage. Surélévation
  10–15 cm (ventilation).
- **Toiture terrasse / au sol** : châssis acier galvanisé ou alu,
  **lesté ou ancré** (prise au vent : un panneau = une voile ;
  dimensionnez contre les tornades de la saison des pluies).
- **Inclinaison** : 10–15° en Afrique de l'Ouest, plein sud
  (ou est-ouest pour lisser la production sur la journée).
- **Mise à la terre** : toutes les structures métalliques reliées
  à la terre (câble 16 mm² vert/jaune).

### 7.3 Câblage DC (panneaux → régulateur/onduleur)
- Câble **solaire** (double isolation, résistant UV) : 4 mm² jusqu'à
  ~20 A, 6 mm² au-delà. **Jamais de câble domestique en DC.**
- **Chute de tension < 3 %** en DC (calculez : longueur × 2
  aller-retour × courant).
- Connecteurs **MC4** : sertis proprement, jamais de dominos à l'air
  libre. Un mauvais MC4 = point chaud = **incendie**.
- **Coffret DC** : sectionneur DC, **parafoudre DC type 2**,
  fusibles par string si ≥ 3 strings en parallèle.
- Séparez physiquement DC et AC dans les cheminements.

### 7.4 Câblage AC et protections
- Disjoncteur AC dédié en sortie d'onduleur, **interrupteur
  différentiel 30 mA** sur les circuits desservis.
- **Inverseur de source** (manuel ou automatique) entre
  réseau / solaire / groupe — **jamais de couplage sauvage**.
- Parafoudre **AC type 2** au tableau.
- Terre : une seule prise de terre, toutes les masses reliées
  (panneaux, onduleur, coffrets).

### 7.5 Mise en service
1. Vérifier **toutes** les polarités au multimètre avant branchement.
2. Brancher **batteries d'abord**, puis panneaux (régulateur), puis AC.
3. Configurer : type de batterie (LiFePO4 : tensions bulk/float selon
   fabricant, typiquement 56,8 V / 54,4 V en 48 V), seuils de
   bascule réseau, limites de charge.
4. Tester : coupure réseau simulée → bascule < 20 ms (hybride),
   production à midi ≈ P_théorique × 0,7–0,8.
5. **Documenter** : schéma, réglages, mots de passe → remis au client.

---

## 8. Sécurité — à ne jamais négliger

- **Le DC ne pardonne pas** : une string de 10 panneaux = 400 V DC,
  **l'arc DC ne s'éteint pas tout seul** (pas de passage par zéro
  comme en AC). Sectionneurs DC **adaptés** (pas de disjoncteur AC
  en DC !).
- **Arc-fault** : les onduleurs modernes détectent les arcs
  (AFCI) — activez la fonction.
- **Incendie** : 90 % des sinistres = mauvais connecteurs MC4 ou
  câbles sous-dimensionnés. Serrage au couple, thermographie annuelle
  sur les grosses installations.
- **Batteries** : local ventilé, extincteur **poudre ABC ou CO2**
  (jamais d'eau sur un feu électrique/lithium), pas de stockage de
  carburant à côté.
- **Travail en hauteur** : harnais sur toit, jamais seul.
- **Foudre** : paratonnerre si site exposé + parafoudres DC et AC ;
  en zone orageuse (toute l'Afrique de l'Ouest en saison des pluies),
  c'est indispensable, pas optionnel.

---

## 9. Maintenance et monitoring

### 9.1 Routine
| Fréquence | Action |
|---|---|
| Mensuel | Nettoyage panneaux (eau claire, tôt le matin ; l'harmattan encrasse vite) |
| Mensuel | Vérifier production vs attendu (appli onduleur) |
| Trimestriel | Serrage visuel connecteurs, état câbles (rongeurs !) |
| Semestriel | Vérifier fixations (vents violents), corrosion |
| Annuel | Thermographie des coffrets, test parafoudres, resserrage au couple |
| Annuel | Test capacité batteries (décharge contrôlée) |

