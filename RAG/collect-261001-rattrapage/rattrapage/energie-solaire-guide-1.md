---
id: collect-261001-rattrapage/rattrapage/energie-solaire-guide-1
title: "Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sol", "arr"]
source: docs/RAG/collect-261001-rattrapage/energie_solaire_guide.md
source_anchor: ""
source_lines: [1, 161]
sha256: 5fcc1fb0c685ecd08acc26497fb834cb3c5c2b796f419bc13f20b62bfce8a804
---

# Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)

> De la physique du photon au dimensionnement d'une centrale : panneaux,
> onduleurs, batteries, installation, sécurité, maintenance, économie.
> Exemples chiffrés calibrés pour l'**Afrique de l'Ouest**
> (ensoleillement ~4,5–5,5 kWh/m²/jour, réseau instable, saison des pluies).
>
> Rédigé à partir des connaissances générales du domaine (2025-2026) :
> vérifiez les prix et réglementations locales, ils évoluent vite.

---

## 1. Bases physiques — comprendre avant d'acheter

### 1.1 L'effet photovoltaïque
- Un photon frappe une cellule en silicium → il arrache un électron →
  le champ électrique interne de la jonction P-N dirige les électrons →
  **courant continu (DC)**.
- Une cellule produit ~0,5–0,6 V. 60–72 cellules en série = panneau
  « 12/24 V » (en réalité 30–45 V à vide, Voc).
- La puissance d'un panneau se mesure en **watts-crête (Wc)** : puissance
  sous conditions standard STC (1000 W/m², 25 °C cellule, AM1.5).

### 1.2 L'irradiation : la vraie unité qui compte
- **kWh/m²/jour** (ou « heures de soleil équivalent plein ») : l'énergie
  reçue par jour. C'est LE chiffre du dimensionnement.
- Ordres de grandeur :
  - Abidjan / Afrique de l'Ouest : **4,5 à 5,5** (saison sèche haute,
    saison des pluies basse).
  - Sahel : 5,5 à 6,5. Europe du Nord : 2,5 à 3,5.
- Un panneau de 550 Wc à Abidjan produit environ :
  `550 Wc × 5 h × 0,75 (rendement système) ≈ 2 kWh/jour`.
- La **température** fait perdre du rendement : −0,35 %/°C au-dessus de
  25 °C (cellule). Sous 65 °C de cellule (courant en Afrique), perte
  ~14 %. D'où l'importance de la ventilation sous les panneaux.

### 1.3 Orientation et inclinaison
- Hémisphère nord : plein **sud**. Hémisphère nord proche de
  l'équateur (Côte d'Ivoire, ~5°N) : inclinaison faible, **5–15°**.
- Règle simple : inclinaison ≈ latitude (Abidjan : ~10°).
- 10–15° minimum quand même : en dessous, la poussière et l'eau
  stagnent (l'harmattan et la poussière rouge encrassent vite).
- **L'ombrage est l'ennemi n°1** : 10 % d'ombre sur un panneau peut
  faire perdre 50 % de sa production (les cellules sont en série).
  Étudiez l'ombre à 9h, 12h et 16h, en saison sèche ET des pluies
  (le soleil est plus bas en décembre).

---

## 2. Les panneaux — choisir le bon

### 2.1 Technologies

| Technologie | Rendement | Durée de vie | Notes |
|---|---|---|---|
| Mono PERC | 20–22 % | 25–30 ans | Standard actuel, bon rapport qualité/prix |
| Mono TOPCon (N-type) | 22–24 % | 25–30 ans | Meilleur en chaleur, devient le standard 2024+ |
| HJT | 23–25 % | 25–30 ans | Excellent coefficient température, plus cher |
| Bifaciaux | +5–15 % de gain | 25–30 ans | Face arrière capte le réfléchi (au sol, surélevés) |
| Poly (P-type) | 16–18 % | 20–25 ans | En fin de vie commerciale, à éviter en neuf |
| Couches minces (CdTe) | 17–19 % | 25 ans | Bon en chaleur/ombrage partiel, grandes surfaces |

- **En 2025-2026 : achetez du mono TOPCon 550–580 Wc** (ou HJT si le
  budget suit). C'est le meilleur compromis chaleur/rendement/prix.
- Bifacial = rentable **au sol** (surélevé, sol clair) ; inutile sur
  toit collé (la face arrière ne voit rien).

### 2.2 Lire une fiche technique
- **Pmax** : puissance crête (Wc). **Voc** : tension à vide.
  **Isc** : courant de court-circuit. **Vmpp/Impp** : point de
  puissance max.
- **Coefficient de température** : le plus bas (en valeur absolue) est
  le mieux pour l'Afrique (−0,30 %/°C bat −0,40 %/°C).
- **Garanties** : 2 niveaux — produit (12–15 ans, défauts) et
  **puissance** (25–30 ans, ≥ 87 % à 25 ans). Exigez les deux par écrit.
- **Certifications** : IEC 61215, IEC 61730, CE. Méfiez-vous des panneaux
  sans marquage (contrefaçons fréquentes sur les marchés).

### 2.3 Tester un panneau avant achat (lot)
- Mesurez **Voc** et **Isc** au soleil avec un multimètre : comparez à
  la fiche (±5 % acceptable).
- Inspectez : microfissures, délamination, jaunissement, boîtier de
  jonction bien scellé, diodes bypass présentes.
- Flash-test en usine : demandez le rapport (puissance réelle ≥ Pmax
  annoncée, tolérance positive 0/+5 W de préférence).

---

## 3. Architecture d'une installation

```
 [Panneaux] --DC--> [Régulateur MPPT] --DC--> [Batteries]
     |                                              |
     +------------DC-------------------> [Onduleur] --AC 230V--> [Consommateurs]
                                                |
                                         [Réseau / Groupe] (hybride)
```

### 3.1 Les 4 architectures

| Architecture | Principe | Idéal pour |
|---|---|---|
| **Off-grid (autonome)** | 100 % solaire + batteries | Site isolé, village, relais télécom |
| **Hybride** | Solaire + batteries + réseau/groupe en secours | Maison/bureau avec délestages |
| **Raccordé réseau (grid-tie)** | Solaire injecté, pas de batteries | Réduire la facture (réseau fiable) |
| **Pompage solaire** | Panneaux → variateur → pompe, sans batteries | Forage, irrigation |

- **En Afrique de l'Ouest, l'hybride est roi** : le réseau existe mais
  coupe ; le solaire couvre le jour, les batteries les coupures, le
  réseau/groupe en dernier recours.
- Le grid-tie pur sans batteries est inutile là où le réseau coupe
  (l'onduleur grid-tie **s'arrête** quand le réseau tombe — norme
  anti-îlotage).

### 3.2 Onduleurs

| Type | Usage | Notes |
|---|---|---|
| **String** | 1–2 trackers MPPT, installation simple | Le standard résidentiel (3–10 kW) |
| **Micro-onduleurs** | 1 par panneau (AC dès le toit) | Cher, mais top contre l'ombrage partiel |
| **Hybride** | Solaire + batteries + réseau/groupe | **Le choix Afrique** (Deye, Victron, Growatt…) |
| **Central** | > 100 kW | Centrales au sol |

- **Dimensionnement** : puissance onduleur ≈ 80–100 % de la puissance
  crête panneaux (on surdimensionne légèrement les panneaux, jamais
  l'inverse — un onduleur trop gros tourne à faible rendement).
- **Surge** : un onduleur 5 kW doit encaisser le démarrage d'un
  congélateur/clim (2–3× le nominal quelques secondes). Vérifiez la
  puissance de surcharge (ex. : 10 kW pendant 5 s).
- **Forme d'onde** : **sinusoïdale pure** obligatoire (les « modifiés »
  tuent les moteurs et l'électronique).
- **Rendement** : ≥ 95 % (européen), comptez les pertes dans le
  dimensionnement.

### 3.3 Régulateurs de charge

- **MPPT** (Maximum Power Point Tracking) : +20–30 % de production vs
  PWM, **obligatoire** au-delà de quelques centaines de Wc.
- **Dimensionnement** : courant max = P_panneaux / V_batterie.
  Ex. : 3000 Wc en 48 V → 62,5 A → régulateur **80 A**.
- Tension PV max du régulateur > Voc totale de la string **par temps
  froid** (le Voc augmente quand il fait froid : +0,3 %/°C sous 25 °C).
- Plusieurs strings d'orientations différentes = plusieurs MPPT
  (ou micro-onduleurs).

---

## 4. Batteries — le poste le plus cher et le plus critique

### 4.1 Comparatif des technologies

|  | Plomb (AGM/Gel) | Lithium LiFePO4 |
|---|---|---|
| Profondeur de décharge | 50 % max | 80–90 % |
| Cycles | 500–1 200 | 4 000–6 000+ |
| Durée de vie | 3–5 ans (chaleur) | 10–15 ans |
| Rendement charge/décharge | 80–85 % | 95 %+ |
| Entretien | Nul (AGM) mais chaleur = mort | BMS intégré, zéro entretien |
| Prix (par kWh utile) | Moins cher à l'achat, **plus cher sur 10 ans** | Plus cher à l'achat, rentable |
| Sécurité | Dégage de l'hydrogène (ventiler) | Très stable (LiFePO4, pas NMC) |

