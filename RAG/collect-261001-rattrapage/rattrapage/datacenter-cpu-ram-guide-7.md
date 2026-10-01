---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-7
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Oracle"]
dates: ["2026-09-27"]
keywords: ["inference", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [932, 1113]
sha256: 0255ed9ec1725ea899211f7b1e7278ea7f30b2ba6d202839295efa4f6c6a85dd
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

```
 Socket 0 : [C0][C1][C2][C3][C4][C5][C6][C7][C8][C9][C10][C11]
             64G 64G 64G 64G 64G 64G 64G 64G 64G 64G 64G 64G  = 768 Go
             ^^ tous les canaux remplis, barrettes identiques = pleine BP
```

---

## 46. Canaux et symétrie : la règle n°1 du dimensionnement

La bande passante mémoire d'un serveur = **nombre de canaux remplis ×
débit par canal**. Tout le reste (fréquence, rangs) est secondaire.

Tableau d'impact (socket 12 canaux, DDR5-5600, débit théorique/can. = 44,8 Go/s) :

| Barrettes | Canaux actifs | BP théorique | Perte |
|---|---|---|---|
| 12 × 64 Go | 12/12 | 537 Go/s | 0 % |
| 8 × 64 Go | 8/12 | 358 Go/s | -33 % |
| 6 × 128 Go | 6/12 | 269 Go/s | -50 % |
| 4 × 128 Go | 4/12 | 179 Go/s | -67 % |

**6 × 128 Go = 768 Go avec moitié moins de bande passante que 12 × 64 Go
= 768 Go.** Même capacité, moitié moins de débit, pour un prix souvent
supérieur. C'est l'erreur de dimensionnement la plus chère et la plus
fréquente : on achète de la capacité en sacrifiant la bande passante.

Règle : **dimensionnez d'abord le nombre de barrettes (= canaux), ensuite
la capacité unitaire.**

## 47. 1 DPC vs 2 DPC : l'impact sur la fréquence

DPC = DIMM Per Channel. Ajouter une 2e barrette sur un canal augmente la
charge électrique du bus et force le contrôleur à baisser la fréquence
pour rester stable.

| Configuration | Fréquence typique (indicatif) | Quand l'accepter |
|---|---|---|
| 1 DPC | max JEDEC (6400 sur Turin/6900P) | toujours préférable |
| 2 DPC RDIMM | -1 à -2 crans (ex. 6400 → 5200/5600) | besoin de capacité > 1 DPC |
| 2 DPC LRDIMM | souvent -1 cran | in-memory |
| 2 DPC MRDIMM | chute significative | à éviter si la BP est le critère |

Exemple concret : 12× 64 Go à 6400 (768 Go, pleine vitesse) vs 24× 64 Go
à 5200 (1,5 To, -19 % de débit). Si votre workload est sensible à la bande
passante, **mieux vaut 12× 128 Go à 6400** (1,5 To, pleine vitesse) que
24× 64 Go — même si la barrette 128 Go coûte plus cher à l'unité.

## 48. Règles de calcul : RAM par cœur

Trois ratios de terrain, par workload (RAM installée / cœurs physiques) :

| Workload | Ratio recommandé | Exemple (96 cœurs) |
|---|---|---|
| Virtualisation généraliste | 4 à 8 Go/cœur | 384 à 768 Go |
| Virtualisation dense / VDI | 2 à 4 Go/cœur | 192 à 384 Go |
| Bases de données | 8 à 32 Go/cœur (selon dataset) | dataset + overhead (section 50) |
| HPC | 2 à 4 Go/cœur (sauf gros maillages) | 192 à 384 Go |
| Conteneurs / cloud-natif | 2 à 4 Go/cœur | 192 à 384 Go |
| IA / inference CPU | 8 à 16 Go/cœur (modèles en RAM) | 768 Go à 1,5 To |

Méthode :
1. Partez du workload, pas du CPU : listez les consommateurs (VM, datasets,
   modèles).
2. Appliquez le ratio, arrondissez au multiple de 12 barrettes (ou 8/16
   selon plateforme) supérieur.
3. Vérifiez la règle des canaux (section 46) : le ratio ne doit jamais vous
   faire laisser des canaux vides.

## 49. Règles de calcul : RAM par VM (virtualisation)

Formule de dimensionnement d'un hôte :

```
RAM_hôte = Σ(RAM_allouée_VM × taux_surallocation_mémoire) + overhead_hyperviseur
```

Valeurs de terrain :
- **Overcommit mémoire** : 1,2 à 1,5 en généraliste (la plupart des VM
  n'utilisent pas toute leur RAM) ; 1,0 (pas d'overcommit) pour les VM
  critiques/DB.
- **Overhead hyperviseur** : ~5 % + ~1 Go par dizaine de VM (indicatif).
- **Réserve HA** : en cluster N+1, ne remplissez chaque hôte qu'à
  (N-1)/N de sa capacité (ex. 67 % sur 3 nœux).

Exemple : 60 VM × 8 Go = 480 Go alloués ; overcommit 1,3 → 370 Go réels ;
+ overhead 30 Go = 400 Go ; cluster 3 nœuds → 400 / 0,67 = 600 Go/hôte →
**12 × 64 Go = 768 Go** (12 canaux remplis, marge incluse).

## 50. Règles de calcul : bases de données (dataset + overhead)

Formule :

```
RAM_DB = dataset_actif + overhead_moteur + marge_croissance
```

- **dataset_actif** : la partie chaude des données (pas la taille totale
  sur disque). Mesurez-la (pg_stat, AWR, compteurs OS) au lieu de la deviner.
- **overhead_moteur** : 20 à 40 % du dataset (buffers, caches, connexions,
  tri, WAL). Comptez 30 % par défaut.
- **marge_croissance** : 20 à 30 % pour 2-3 ans, ou la croissance mesurée.

Exemple : PostgreSQL, 500 Go de données chaudes :
500 + 30 % (150) = 650 Go ; + 25 % croissance = ~815 Go →
**12 × 64 Go = 768 Go** (limite) ou **12 × 96 Go = 1 152 Go** (confortable,
recommandé). En 2 DPC si la carte l'exige, en acceptant la baisse de
fréquence (section 47) — ou 12× 96 Go en 1 DPC si la carte a 12 slots.

Cas Oracle/SQL Server : ajoutez la contrainte licences (section 25) —
parfois il vaut mieux **moins de RAM et un CPU « F »** que l'inverse.

## 51. Cas chiffré n°1 : cluster Proxmox 3 nœuds (indicatif)

Besoin : 120 VM généralistes (8 Go allouées en moyenne), HA N+1.

| Poste | Calcul | Choix |
|---|---|---|
| RAM allouée totale | 120 × 8 = 960 Go | — |
| RAM réelle (overcommit 1,3) | 960 / 1,3 ≈ 740 Go | — |
| Par nœud (N+1, 67 %) | 740 / 2 nœuds utiles ≈ 370 Go → 768 Go | 12 × 64 Go DDR5-5600 |
| CPU par nœud | 120 VM / ~1,5 vCPU par cœur ≈ 80 cœurs | 1× EPYC 9655 (96 c.) |
| Conso/nœud estimée | 400 W CPU + 150 W RAM + 150 W divers | ~700 W en charge |

BOM par nœud (indicatif) : 1× EPYC 9655 (~10 800 $ tarif), 12× 64 Go
RDIMM, 4× NVMe 3,84 To, 2× 25 GbE, 2× 1200 W. **La RAM représente ~25-35 %
du total** : ne la rognez jamais en premier.

## 52. Cas chiffré n°2 : serveur PostgreSQL 2 To (indicatif)

Besoin : 1,2 To de données chaudes, OLTP + analytique, croissance 20 %/an.

| Poste | Calcul | Choix |
|---|---|---|
| Dataset + overhead 30 % | 1,2 To × 1,3 = 1,56 To | — |
| + croissance 2 ans | × 1,44 ≈ 2,25 To | **12 × 256 Go LRDIMM = 3 To** |
| CPU | OLTP → fréquence | 1× EPYC 9575F (64 c., 5,0 GHz) |
| Alternative économe | 2× 128 Go ne suffit pas (2,25 To) | LRDIMM obligatoire ici |
| Stockage | WAL séparé, NVMe faible latence | 2× 1,6 To (WAL) + 8× 7,68 To |
| Conso estimée | 400 W CPU + 250 W RAM + 200 W | ~850 W en charge |

Note : 3 To en 12× 256 Go LRDIMM = le cas d'école du LRDIMM (section 36).
En RDIMM il faudrait 24× 128 Go en 2 DPC (fréquence réduite) — le LRDIMM
gagne ici malgré son prix.

## 53. Cas chiffré n°3 : hôte VDI 200 postes (indicatif)

Besoin : 200 bureaux virtuels, 4 Go/poste, pic simultané 70 %.

| Poste | Calcul | Choix |
|---|---|---|
| RAM postes | 200 × 4 Go × 70 % = 560 Go | — |
| + overhead VDI 20 % | ≈ 670 Go | **12 × 64 Go = 768 Go** |
| CPU | 200 postes / ~2,5 postes par cœur | 1× EPYC 9745 (128 c. Zen 5c) |
| Ratio | 768 / 128 = 6 Go/cœur | dans la fourchette VDI |
| Stockage | I/O aléatoires intenses | NVMe, cache écriture dimensionné |
| Réseau | 200 × ~50 Mb/s en pic | 2× 25 GbE minimum |

Le VDI est un workload **cœurs + IOPS**, pas mémoire pure : ne surpayez
pas la fréquence CPU, misez sur la densité (Zen 5c / E-core) et le stockage.

## 54. Consommation électrique des DIMM : ordres de grandeur

| Type de barrette | Au repos (indicatif) | En charge (indicatif) |
|---|---|---|
| RDIMM 32-64 Go DDR5 | 2 à 4 W | 8 à 12 W |
| RDIMM 96-128 Go DDR5 | 3 à 5 W | 10 à 15 W |
| LRDIMM 256 Go | 4 à 6 W | 15 à 20 W |
| MRDIMM-8800 64 Go | 5 à 8 W | 20 à 30 W |

Exemple : 24× RDIMM 64 Go en charge ≈ 24 × 10 = **240 W** — l'équivalent
d'un demi-CPU. Sur un serveur 2S HPC avec 24 MRDIMM, la RAM peut dépasser
**500 W** (mesure Phoronix : +437 W en passant de DDR5-6400 à MRDIMM-8800
sur 12 barrettes — vérifié le 27/09/2026).

Règle : **comptez 10 W par RDIMM et 25 W par MRDIMM** dans vos bilans de
puissance (partie H). C'est une marge prudente, pas une mesure.

## 55. RAS : Chipkill, SDDC, PPR, et la supervision

La RAM serveur ne se contente pas de corriger : elle **anticipe** :

