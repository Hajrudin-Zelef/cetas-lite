---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-5
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "attention", "dram", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [610, 773]
sha256: 234d06ebbeb1f31012a303a8a1603bb4e53ce257565fae561f5231692b158021
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Évitez 4S si :
- Le workload se parallélise (alors 2× 2S coûtent moins cher en licences
  et en pannes).
- C'est « pour la marge de croissance » : un châssis 4S coûte 2 à 3× plus
  cher à performance égale, et la croissance se fait rarement socket par
  socket.
- Vous n'avez pas d'équipe capable de tuner NUMA sur 4 nœuds (les pertes
  NUMA croissent avec le nombre de sockets).

## 28. Scale-out contre scale-up : la décision structurante

- **Scale-up** (gros serveur) : moins de nœuds à administrer, licences par
  socket contenues, NUMA à gérer, panne = gros impact.
- **Scale-out** (beaucoup de petits serveurs) : tolérance de panne native,
  upgrade par vagues, plus de ports réseau/énergie, plus d'administration
  (ou d'automatisation).

La tendance 2024-2026 est au **scale-out avec des nœuds 1S/2S denses**
(EPYC 5c, Xeon 6+), parce que le coût du cœur dense a chuté et que
l'automatisation (Ansible, Terraform) a réduit le coût d'administration.
Le scale-up ne se justifie plus que pour les workloads monolithiques
(section 27) et quelques niches HPC.

## 29. Interconnexions inter-sockets : UPI et Infinity Fabric

| Technologie | Constructeur | Débit (vérifié) | Usage |
|---|---|---|---|
| UPI 2.0 | Intel Xeon 6 | 24 GT/s par lien | 4 liens (6700), 6 liens (6900) |
| Infinity Fabric | AMD EPYC | génération 5 (Turin) | 1 à 2 liens selon SKU |

En pratique vous n'avez rien à régler : la topologie est fixe par la carte
mère. Ce qu'il faut savoir :
- Tous les liens UPI/IF consomment de l'énergie **même à vide** (quelques
  dizaines de watts par socket en 2S) : un serveur 2S consomme plus qu'un
  1S à charge égale, à cause de l'interconnexion et du 2e jeu de DIMM/VRM.
- En 2S, la bande passante inter-socket (~200-250 ns de latence, quelques
  dizaines de Go/s) est **10× plus lente** que la mémoire locale : d'où
  l'importance du placement NUMA (section 24).

## 30. La règle d'or : bande passante mémoire par cœur

Formule : **BP/cœur = (canaux × débit × 8) / nombre de cœurs.**

Exemples (théoriques, débit crête) :

| Config | Calcul | BP/cœur |
|---|---|---|
| 1S EPYC 9965 (192 c., 12×6400) | 614 / 192 | **3,2 Go/s/cœur** |
| 1S EPYC 9755 (128 c., 12×6400) | 614 / 128 | **4,8 Go/s/cœur** |
| 1S EPYC 9175F (16 c., 12×6400) | 614 / 16 | **38 Go/s/cœur** |
| 2S Xeon 6980P (256 c., 24×6400) | 1228 / 256 | **4,8 Go/s/cœur** |
| 1S Xeon 6990E+ (288 c., 12×8000) | 768 / 288 | **2,7 Go/s/cœur** |
| 1S Venice 256 c. (16×8000) | 1600 / 256 | **6,25 Go/s/cœur** |

Lecture : **plus le CPU est dense, plus chaque cœur est « affamé ».**
Un 9965 à 3,2 Go/s/cœur est parfait pour du web/conteneurs (peu gourmands
en mémoire), mais sous-alimenté pour du HPC. C'est pour cela que les gammes
HPC restent à 64-128 cœurs avec toute la bande passante pour eux — et que
MRDIMM existe (section 17).

## 31. Exemple de config : virtualisation 2S (indicatif)

Besoin : 150 VM généralistes, 2 To de RAM, haute dispo sur 3 nœuds.

| Composant | Choix (indicatif) | Justification |
|---|---|---|
| CPU | 2× EPYC 9655 (96 c. Zen 5) | 192 cœurs, bon compromis fréquence/densité |
| RAM | 24× 64 Go DDR5-5600 RDIMM = 1,5 To | 12 canaux/socket remplis, symétrie |
| Stockage | 8× NVMe 3,84 To (RAID logiciel/CEP H) | I/O par VM |
| Réseau | 2× 25 GbE | vMotion + production |
| Alimentation | 2× 1600 W redondantes | voir section 84 |
| Conso estimée | ~900-1100 W en charge | voir section 85 |

Variante Intel : 2× Xeon 6780E/6980E+ si la métrique licence favorise les
cœurs sans HT, ou 2× 6700P si le budget est serré (8 canaux suffisent en
virtualisation généraliste).

## 32. Exemple de config : base de données 1S (indicatif)

Besoin : PostgreSQL 800 Go actifs, OLTP, licences à minimiser.

| Composant | Choix (indicatif) | Justification |
|---|---|---|
| CPU | 1× EPYC 9375F (32 c., 4,8 GHz) | fréquence max, 32 cœurs = licences contenues |
| RAM | 12× 64 Go DDR5-6000/6400 = 768 Go | dataset + 30 % overhead (section 50) |
| Stockage | 4× NVMe 7,68 To, faible latence | WAL + données séparés |
| Réseau | 2× 25 GbE | réplication |
| Pourquoi 1S | pas de NUMA, licences / socket minimes | version P si dispo (-10 %) |

Le « F » coûte plus cher au CPU (5 306 $) mais économise des dizaines de
milliers d'euros de licences sur 3 ans. **C'est l'exemple type où le CPU
le plus cher au catalogue est le moins cher au TCO.**

## 33. Exemple de config : HPC 2S + MRDIMM (indicatif)

Besoin : simulation CFD, nœud de cluster, bande passante maximale.

| Composant | Choix (indicatif) | Justification |
|---|---|---|
| CPU | 2× Xeon 6980P (128 c., 500 W) | AVX-512, 12 canaux, MRDIMM natif |
| RAM | 24× 64 Go MRDIMM-8800 = 1,5 To | +24 % mesuré sur HPCG (section 17) |
| Refroidissement | vérifié 2× 500 W + MRDIMM (+437 W mesurés) | watercooling ou air très haut de gamme |
| Réseau | 1× 200 GbE (InfiniBand ou RoCE) | MPI inter-nœuds |
| Alimentation | 2× 2000 W+ redondantes | pic > 2 kW par serveur (mesuré 1237 W pour 1 CPU + 12 DIMM) |
| Conso estimée | 1800-2400 W en charge | à budgéter en kW/rack (section 78) |

Attention : un nœud HPC 2S moderne consomme **autant qu'un radiateur
électrique**. Dix nœuds = 20 kW = une petite salle à refroidir. Le
dimensionnement électrique et frigorifique (partie H) n'est pas optionnel.

---

## 34. Pourquoi la RAM serveur n'est pas de la RAM PC

Quatre différences fondamentales avec la DDR5 « desktop » :

1. **ECC obligatoire** (section 41) : la RAM ECC corrige les erreurs 1 bit
   et détecte les 2 bits. Sur des To de RAM en 24/7, les erreurs sont une
   certitude statistique, pas un risque théorique.
2. **Registered (buffered)** : un registre entre le contrôleur et les puces
   DRAM permet d'adresser beaucoup plus de barrettes par canal (section 35).
3. **Densité et validation** : barrettes 64/96/128/256 Go, validées par
   l'OEM sur la carte mère exacte (QVL). Une barrette « compatible sur le
   papier » qui n'est pas sur la QVL = risque d'instabilité ou de
   sous-fréquence.
4. **Télémétrie** : température par barrette (TSOD), compteurs d'erreurs
   ECC remontés au BMC, PPR (section 55). La RAM serveur se **supervise**.

Conséquence d'achat : **n'achetez jamais de RAM « desktop » pour un
serveur**, même si le format physique ressemble. Et prévoyez la RAM dans le
devis initial : elle représente 20 à 40 % du prix d'un serveur configuré.

## 35. RDIMM : le standard (Registered DIMM)

Le RDIMM intercale un **registre d'adresse/commande** entre le contrôleur
mémoire et les puces DRAM. Effet : le contrôleur ne « voit » qu'une charge
électrique réduite, ce qui permet 2 barrettes par canal (2 DPC) en restant
stable.

| Caractéristique | RDIMM DDR5 |
|---|---|
| Capacités courantes | 16, 32, 48, 64, 96, 128 Go (256 Go en 3DS, voir 37) |
| Fréquences serveur réelles | 4800, 5200, 5600, 6000, 6400 MT/s (vérifié) |
| Rangs | 1R (single) ou 2R (dual) |
| DPC | 1 ou 2 (2 DPC = fréquence réduite, section 47) |
| Usage | 80 %+ des serveurs : le choix par défaut |

**Règle d'achat : si vous hésitez entre RDIMM et LRDIMM, prenez RDIMM.**
Le LRDIMM ne se justifie qu'au-delà des capacités RDIMM standard (section
36) et coûte plus cher avec une latence légèrement supérieure.

## 36. LRDIMM : quand la densité l'exige (Load-Reduced DIMM)

Le LRDIMM ajoute un **buffer de données** (en plus du registre) qui isole
totalement les puces DRAM du bus : le contrôleur ne voit qu'une seule
charge, quelle que soit la densité.

| Critère | RDIMM | LRDIMM |
|---|---|---|
| Capacité max/barrette | 128 Go (256 Go en 3DS) | 256 Go (standard), plus en 3DS |
| Latence | référence | +1 à 3 ns (buffer) |
| Prix | référence | +20 à 50 % (indicatif) |
| Cas d'usage | standard | > 1,5-2 To par socket, in-memory |

