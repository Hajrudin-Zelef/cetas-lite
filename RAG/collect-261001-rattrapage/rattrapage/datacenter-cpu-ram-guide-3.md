---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-3
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-06-01", "2026-09-27"]
keywords: ["18a", "accelerator", "amd", "clearwater forest", "datacenter", "inference", "intel", "memory", "packaging"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [290, 456]
sha256: da2a30c75704b3ded9f92071f032021775f6f9a21f144a597b89fd140dada822
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Caractéristique | Xeon 6900P |
|---|---|
| Cœurs | jusqu'à 128 P-cores / 256 threads (HT) |
| Socket | LGA 7529 (« Avenue City ») |
| Sockets/serveur | 1 ou 2 |
| TDP | jusqu'à 500 W (réf. 550 W type 6966P-C) |
| Canaux mémoire | 12 × DDR5-6400 (RDIMM) |
| MRDIMM | jusqu'à 8800 MT/s (1re génération) |
| PCIe | 96 lignes Gen5 (+ CXL 2.0 sur 64 lignes) |
| Interconnexion | 6 liens UPI 2.0 à 24 GT/s |
| Réf. amirale | 6980P : 128 cœurs, 500 W, 504 Mo de cache, 17 800 $ |
| Autres réf. | 6979P (120 c.), 6972P (96 c.), 6960P/6962P (72 c.), 6952P (96 c., 400 W), 6944P (72 c., 350 W) |

Points distinctifs vs AMD :
- **MRDIMM natif** : premier au monde à supporter les barrettes à rangs
  multiplexés (voir section 17) — +37,5 % de bande passante vs DDR5-6400.
- **Accélérateurs intégrés** : DSA, IAA, QAT, DLB (section 18).
- **AMX** : accélération IA matricielle par cœur, atout décisif pour
  l'inférence CPU (section 22).

## 13. Sierra Forest : Xeon 6700E / 6900E (E-core)

La réponse d'Intel au scale-out dense (vérifié le 27/09/2026) :

| Caractéristique | Xeon 6700E | Xeon 6900E |
|---|---|---|
| Cœurs | jusqu'à 144 E-cores | jusqu'à 288 E-cores |
| Socket | LGA 4710 | LGA 7529 |
| Sockets/serveur | 1 ou 2 | 1 ou 2 |
| TDP max | 330 W | 500 W |
| Canaux mémoire | 8 × DDR5 | 12 × DDR5 |
| Hyper-threading | non | non |
| Cible | cloud, edge dense | hyperscale, telco |

Note importante : les E-cores Crestmont de Sierra Forest sont des cœurs
« petits » mais modernes (bien plus rapides que les Atom d'il y a dix ans).
Ils excellent en débit/rack sur workloads parallèles, mais perdent face aux
P-cores et aux Zen 5 sur tout workload sensible à la performance mono-thread
ou aux instructions vectorielles larges.

## 14. Xeon 6700/6500 P-core : le milieu de gamme 8 canaux

Pour les serveurs généralistes qui n'ont pas besoin de 12 canaux (lancé
début 2025, vérifié le 27/09/2026) :

| Caractéristique | Xeon 6700P / 6500P |
|---|---|
| Cœurs | jusqu'à 86 P-cores |
| Socket | LGA 4710 |
| Sockets/serveur | 1, 2, **4 et 8** (le 6900P est limité à 2S) |
| TDP max | 350 W |
| Canaux mémoire | 8 × DDR5-6400 / MRDIMM 8000 |
| PCIe | 88 lignes Gen5 (136 en mono-socket) |
| UPI | 4 liens |

C'est la gamme « scale-up » d'Intel : un serveur 8 sockets 6700P peut aligner
688 P-cores dans une seule image OS — le pendant du monde x86 face aux
gros systèmes RISC/Itanium d'autrefois. En pratique, 4S/8S restent des niches
(bases de données monolithiques, voir section 27).

## 15. Sockets Intel : LGA-7529 vs LGA-4710

| Caractéristique | LGA 7529 (« Birch Stream ») | LGA 4710 (« Beechnut City ») |
|---|---|---|
| Gammes | Xeon 6900P / 6900E / 6+ | Xeon 6700P / 6700E / 6500 |
| Contacts | 7529 | 4710 |
| Canaux mémoire | 12 | 8 |
| TDP max | 500 W | 350 W (330 W en E) |
| Sockets max | 2 | 8 (P-core) / 2 (E-core) |
| PCIe max | 96 (Gen5) | 88 (Gen5) |

Conséquence d'achat : **les deux sockets sont incompatibles entre eux.**
Choisir 6700 ou 6900, c'est choisir une carte mère, un châssis et une
trajectoire d'upgrade pour 5 ans. Le 6900 coûte plus cher en plateforme
(cartes mères 12 canaux, VRM 500 W) mais offre 50 % de bande passante mémoire
en plus — décisif pour HPC et bases de données (section 30).

## 16. Xeon 6+ « Clearwater Forest » (vérifié le 27/09/2026)

Lancé au Computex le 01/06/2026, c'est le premier CPU datacenter Intel gravé
en **Intel 18A** (transistors RibbonFET + alimentation par l'arrière
PowerVia), avec packaging Foveros Direct 3D :

| Caractéristique | Xeon 6+ (Clearwater Forest) |
|---|---|
| Cœurs | jusqu'à **288 E-cores Darkmont** (6990E+) |
| Gamme | 4 références, 144 à 288 cœurs (6960E+ → 6990E+) |
| Cache L3 | 576 Mo (amiral) |
| Mémoire | 12 canaux **DDR5-8000** natif |
| PCIe | 96 lignes Gen5, 64 lignes CXL 2.0 |
| TDP | 300 à 450 W (6990E+ : versions 450 W et 330 W) |
| Socket | LGA 7529 (compatible plateformes 69xx) |
| Prix publics | 6990E+ : 14 995 $ ; 6980E+ (264 c.) : 9 850 $ |
| Cible | cloud-natif, telco, IA agentique |

Chiffres constructeur (à relativiser, voir section 2) : +30 % de performance
par thread vs EPYC 9965 (192 cœurs), jusqu'à 1,3× perf/watt à 40 % de charge,
consolidation jusqu'à 9:1 vs Xeon de 2e génération. Les tests indépendants
(Phoronix, ServeTheHome) sont attendus pour trancher — **ne basez pas un
appel d'offres uniquement sur ces chiffres.**

## 17. MRDIMM : le pont entre DDR5 et DDR6 (vérifié le 27/09/2026)

MRDIMM (Multiplexed Rank DIMM) : une barrette DDR5 dont un multiplexeur
embarqué combine **deux rangs simultanément**, doublant le débit effectif
sans changer le slot physique (standard JEDEC, promu par Intel) :

| Génération MRDIMM | Débit | Disponibilité (vérifié) |
|---|---|---|
| Gen1 | 8800 MT/s | en production (Xeon 6900P) |
| Gen2 | 12 800 MT/s | systèmes attendus T1 2027 |
| Gen3 | 17 600 MT/s | horizon ~2030 |

Mesures indépendantes (Phoronix, serveur 1× Xeon 6980P, 12 barrettes) :
- HPCG : **+24 %** avec MRDIMM-8800 vs DDR5-6400.
- Surcoût : environ **+100 $ par barrette 64 Go** vs DDR5-6400 ECC.
- Surconsommation : en charge mémoire intensive, le serveur est passé de
  **800 W à 1237 W moyens** ; sur charges peu mémoire, l'écart tombe à
  ~15 W. Les barrettes MRDIMM ne chauffent que ~5 °C de plus.
- AMD suit : EPYC Venice (SP7) supportera les MRDIMM Gen2 à 12 800 MT/s
  (voir section 93).

Verdict : MRDIMM = bande passante de classe DDR6 **sans changer de
plateforme**, au prix d'un surcoût mémoire et d'un surcroît de puissance
à budgéter (partie H). Pertinent pour HPC/IA/DB analytique, inutile pour
de la virtualisation généraliste.

## 18. Accélérateurs intégrés Intel : DSA, IAA, QAT, DLB

Les Xeon 6 embarquent des moteurs fixes (pas des cœurs) qui déchargent le CPU.
Ils n'existent pas chez AMD — avantage différenciant à évaluer selon le
workload :

| Accélérateur | Rôle | Workload gagnant |
|---|---|---|
| DSA (Data Streaming Accelerator) | copies/mémoire, CRC, mouvements de données | stockage, réseau, virtualisation |
| IAA (In-Memory Analytics Accelerator) | compression/décompression, analytic | bases de données, data lakes |
| QAT (QuickAssist Technology) | chiffrement, compression asymétrique | TLS/IPsec massif, VPN, backup |
| DLB (Dynamic Load Balancing) | répartition de files réseau | NFV, telco, pare-feu |

Conditions : il faut un OS et des applications qui les exploitent (pilotes,
librairies). Un QAT inutilisé vaut zéro. **Avant de choisir Intel « pour les
accélérateurs », vérifiez que votre stack logicielle les supporte** (DPDK
pour le réseau, etc.). Sinon, comparez au prix/perf brut.

## 19. Choisir son Xeon par workload : table de décision

| Workload | Famille | Références types (indicatif) | Critère n°1 |
|---|---|---|---|
| Virtualisation généraliste | 6700P | 6731P/6740P (classes) | prix/cœur, 8 canaux suffisent |
| Virtualisation très dense | 6900E / 6+ | 6980E+, 6960E+ | cœurs/rack, perf/watt |
| Base de données OLTP | 6900P | 6960P, 6972P (haute fréquence) | fréquence, AMX |
| Base de données analytique | 6900P + MRDIMM | 6980P | bande passante mémoire |
| HPC / simulation | 6900P + MRDIMM | 6980P | AVX-512, BP mémoire |
| Inference IA sur CPU | 6900P | 6980P, 6972P | AMX (atout majeur) |
| Telco / NFV | 6700E / 6+ | 6780E, 6990E+ | débit/watt, DLB/QAT |
| Edge / PME | 6500P | entrée de gamme | prix, TDP contenu |

Règle d'or Intel : **P-core = performance par thread et accélérateurs ;
E-core = densité et efficacité.** Ne mélangez jamais les deux dans un même
comparatif prix sans préciser lequel vous chiffrez.

## 20. Comparatif EPYC vs Xeon : virtualisation

La virtualisation (VMware, Proxmox, Hyper-V) est le workload n°1 des
datacenters d'entreprise. Critères : cœurs/€, RAM/cœur, licences, I/O.

