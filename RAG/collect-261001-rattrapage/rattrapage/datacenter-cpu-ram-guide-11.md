---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-11
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["arr", "datacenter", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1629, 1802]
sha256: 9128e3d37c471bf60d6bf2f553d263752fb3f67ce8437b5243b63f2395664764
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

```
 Facture annuelle = P_max_serveur x taux_charge x 8760 h x PUE x prix_kWh

 où P_max_serveur = estimation section 85 (pas le TDP seul !)
```

Les trois multiplicateurs qu'on oublie :
1. **Taux de charge réel** : un serveur de virtualisation tourne à 20-40 %
   en moyenne, pas à 100 %. À 30 % de charge CPU, la consommation n'est pas
   30 % du max mais **50-60 %** (le socle idle est élevé : ~30-40 % du max
   pour un serveur 2S moderne — ordre de grandeur).
2. **PUE** : multiplie tout (section 81).
3. **Prix du kWh** : partez du prix **tout compris** (acheminement,
   taxes), pas du prix spot.

Exemple : serveur estimé à 1000 W max, charge 30 % → ~550 W réels ;
PUE 1,5 → 825 W soutirés ; 8760 h → 7 226 kWh/an ; à 0,18 €/kWh →
**~1 300 €/an**. Sur 5 ans : 6 500 € — souvent plus que la RAM du serveur.

## 84. Rendement des alimentations : 80 PLUS et au-delà

Les alimentations serveur convertissent l'AC en DC avec des pertes. Le
rendement dépend de la **charge** (courbe en cloche, optimum vers 50 %) :

| Certification (230 V) | 20 % | 50 % | 100 % |
|---|---|---|---|
| 80 PLUS Gold | 88 % | 92 % | 88 % |
| 80 PLUS Platinum | 90 % | 94 % | 91 % |
| 80 PLUS Titanium | 90 % | **96 %** | 94 % |

Règles :
- **Surdimensionner l'alimentation dégrade le rendement** : un 2000 W à
  20 % de charge (400 W) tourne à ~88 % au lieu de 96 % → 8 % de pertes
  en plus, 24/7. Dimensionnez pour viser **40-60 % de charge en régime
  normal**.
- La redondance 1+1 fait tourner chaque bloc à ~50 % : c'est aussi un
  choix de rendement, pas seulement de disponibilité.
- **Pertes = chaleur dans la salle** : 100 W perdus dans les PSU = 100 W
  à climatiser en plus.

## 85. Tableau : consommation typique par configuration (estimations)

Méthode : 2× TDP CPU × 0,75 (charge réaliste haute) + RAM (10 W/DIMM,
25 W/MRDIMM) + 150 W (carte mère, disques, fans) + 10 % pertes PSU.
**Estimations d'ingénierie — à mesurer au wattmètre/PDU pour vos
configurations exactes.**

| Configuration | Calcul indicatif | Estimation |
|---|---|---|
| 1U 1S EPYC 9355 (32 c., 280 W), 12 DIMM | 210 + 120 + 150 | **~500 W** |
| 2U 2S EPYC 9655 (2× 400 W), 24 DIMM | 600 + 240 + 150 | **~1100 W** |
| 2U 2S EPYC 9965 (2× 500 W), 24 DIMM | 750 + 240 + 200 | **~1300 W** |
| 2U 2S Xeon 6980P (2× 500 W), 24 MRDIMM | 750 + 600 + 200 | **~1700 W** |
| 2U 2S Xeon 6700P (2× 350 W), 16 DIMM | 525 + 160 + 150 | **~900 W** |
| 1U 2S Xeon 6780E (2× 330 W), 16 DIMM | 500 + 160 + 150 | **~850 W** |
| 4U 8× GPU (8× 700 W) + 2 CPU | 4200 + 800 + 400 | **~6000 W** |
| Serveur idle (2S moderne) | socle incompressible | **~250-350 W** |

Notez le **socle idle** : un serveur éteint-logiquement mais allumé
consomme 250-350 W. Éteindre (ou mettre en veille profonde) les serveurs
inutilisés est l'économie d'énergie la plus rentable : 300 W × 8760 h ×
1,5 PUE × 0,18 € = **~710 €/an par serveur éteint**.

## 86. Cas chiffré : baie de 10 serveurs 2S (indicatif)

10× serveurs 2U 2S EPYC (estimation 1100 W chacun en charge, 300 W idle) :

| Poste | Calcul | Résultat |
|---|---|---|
| Puissance IT charge | 10 × 1100 W | 11 kW |
| Puissance IT idle (nuit) | 10 × 300 W | 3 kW |
| Moyenne pondérée (30 % charge) | ~5,5 kW | — |
| Avec PUE 1,5 | 5,5 × 1,5 | **8,25 kW soutirés** |
| Énergie annuelle | 8,25 × 8760 | **72 270 kWh/an** |
| Facture (0,18 €/kWh) | — | **~13 000 €/an** |
| Sur 5 ans | — | **~65 000 €** |

C'est le calcul à présenter à la direction avec tout achat de baie :
**le TCO énergie sur 5 ans représente 30 à 60 % du prix d'achat des
serveurs.** Un serveur « moins cher » de 2 000 € qui consomme 200 W de
plus coûte 200 × 8760 × 1,5 × 0,18 / 1000 ≈ **470 €/an** de plus, soit
2 370 € sur 5 ans : la « bonne affaire » est une perte sèche.

## 87. Dimensionnement onduleur : méthode

L'onduleur (UPS) protège contre les coupures et les microcoupures. Méthode
de dimensionnement :

```
1. P_IT_max = somme des P_max serveurs (section 85) x 1,25 (marge)
2. P_ondul_kVA = P_IT_max / facteur_de_puissance (0,9 typique)
3. Autonomie : energie_batteries >= P_IT x duree_cible
4. Verifier : bypass, redondance (N+1), courbe de surcharge
```

Règles :
- **Ne dimensionnez jamais sur la puissance moyenne** : l'onduleur doit
  tenir le **pic** (démarrage simultané des serveurs = appel de courant).
- Comptez large : un onduleur chargé à > 80 % en permanence vieillit mal
  (batteries) et ne laisse aucune marge d'extension.
- **Onduleur triphasé** dès que P_IT > ~10 kW : équilibrage des phases,
  sections de câbles raisonnables.
- Prévoyez l'**arrêt ordonné** (NUT, voir guide onduleurs) : l'autonomie
  batterie sert à éteindre proprement, pas à « tenir » indéfiniment.

## 88. Dimensionnement onduleur : exemple chiffré (indicatif)

Baie de la section 86 : 10 serveurs, 11 kW max IT.

| Étape | Calcul | Résultat |
|---|---|---|
| P_IT_max | 11 kW | — |
| Marge 25 % | 11 × 1,25 | 13,75 kW |
| kVA (FP 0,9) | 13,75 / 0,9 | **~15,3 kVA → onduleur 20 kVA** |
| Autonomie 15 min | 13,75 kW × 0,25 h | 3,4 kWh utiles de batteries |
| Triphasé ? | > 10 kW | **oui** |

Un onduleur 20 kVA triphasé avec 15 minutes d'autonomie pour cette baie :
ordre de grandeur **8 000-15 000 €** (indicatif, selon marque et batteries).
À comparer aux 65 000 € d'énergie sur 5 ans (section 86) : la protection
électrique coûte ~15-20 % du budget énergie. **Ne pas en mettre est une
économie qui se paie au premier orage.**

## 89. PDU : monophasé vs triphasé

| PDU | Usage | Intensité typique |
|---|---|---|
| Monophasée 16 A (230 V) | petites baies, < 3 kW | 16 A → 3,7 kW max |
| Monophasée 32 A | baie moyenne, < 7 kW | 32 A → 7,4 kW max |
| Triphasée 16 A | standard datacenter | 3 × 16 A → ~11 kW |
| Triphasée 32 A | baies denses | 3 × 32 A → ~22 kW |

Règles :
- **2 PDU par baie** (A et B), chaque serveur branché sur les deux
  (alimentations redondantes) : c'est la base de la disponibilité.
- Chaque PDU doit pouvoir supporter **seule** toute la baie (si l'autre
  disjoncte) : ne chargez pas chaque PDU à plus de 50 % en nominal.
- PDU **monitorées** (par prise ou par phase) : la mesure permanente
  remplace les estimations de ce guide par des faits.

## 90. PDU : calcul d'intensité par phase (triphasé)

Formule : **I = P / (√3 × U × cos φ)** avec U = 400 V, cos φ ≈ 0,95.

Exemple : baie de 11 kW IT sur PDU triphasée 400 V :
I = 11 000 / (1,732 × 400 × 0,95) ≈ **16,7 A par phase**.

Vérifications :
- 16,7 A < 16 A ? **Non** → PDU 16 A insuffisante : passer en 32 A ou
  répartir sur 2 baies.
- Avec la règle des 50 % (redondance A/B) : il faut pouvoir tenir
  16,7 A sur **une seule** PDU → PDU 32 A triphasée par baie.
- **Équilibrez les phases** : 3 serveurs par phase, pas 5/1/0. Un
  déséquilibre fait disjoncter le neutre avant les phases.

## 91. Groupes électrogènes et serveurs : repères

Au-delà de l'onduleur (minutes), le groupe électrogène (heures/jours) :
- **Puissance** : ≥ P_IT_max × 1,5 (le groupe n'aime pas tourner à
  100 %, et il alimente aussi le froid — qui peut représenter autant que
  l'IT).
- **Démarrage** : 10-30 s. L'onduleur couvre l'intervalle : vérifiez que
  l'autonomie batterie > temps de démarrage + marge (section 88).
- **Qualité du courant** : les groupes produisent une tension parfois
  « sale » (harmoniques) : l'onduleur **double conversion** (VFI) isole
  les serveurs — un des arguments pour le VFI en datacenter.
- **Tests** : test mensuel en charge (banc de charge ou bascule réelle).
  Un groupe qui n'a pas tourné depuis 2 ans est une décoration.

## 92. Refroidissement de la salle : kW frigorifiques

Bilan thermique : **toute l'énergie IT devient chaleur** (+ les pertes
électriques). Pour 11 kW IT avec PUE 1,5 : 16,5 kW à évacuer.

