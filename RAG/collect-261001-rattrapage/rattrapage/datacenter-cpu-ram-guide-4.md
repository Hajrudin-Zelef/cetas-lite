---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-4
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Oracle"]
dates: ["2026-09-27"]
keywords: ["amd", "attention", "datacenter", "gpu", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [457, 609]
sha256: 3db355b4a0ba53a06a9d2258fc92f4cedc21f3d5f03ec668724cca57d59f935e
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Critère | AMD EPYC Turin (Zen 5c) | Intel Xeon 6 (E-core / P-core) |
|---|---|---|
| Densité max | 192 c./384 threads (SMT) | 288 c./288 threads (6+, sans HT) |
| RAM/cœur | 12 canaux / 192 cœurs | 12 canaux / 288 cœurs (6+) |
| Licences VMware (par cœur) | SMT = threads comptés ? vérifier l'éditeur | pas de SMT : 1 cœur = 1 cœur |
| Prix CPU amiral | 14 813 $ (9965) | 14 995 $ (6990E+) |
| Maturité plateforme | SP5 éprouvé depuis 2022 | LGA 7529 éprouvé depuis 2024 |

Point d'attention **licences** : certains éditeurs comptent les threads
(avec SMT) et d'autres les cœurs physiques. Un EPYC 192 cœurs/384 threads
peut coûter 2× plus cher en licences qu'un Xeon E-core 288 cœurs sans HT
selon la métrique de l'éditeur. **Le choix CPU ne se fait jamais sans la
calculette licences** (section 25).

## 21. Comparatif EPYC vs Xeon : bases de données

| Critère | AMD EPYC Turin | Intel Xeon 6900P |
|---|---|---|
| OLTP (latence) | 9175F/9575F : 5,0 GHz, gros L3 | 6960P/6972P : haute fréquence + AMX |
| Analytique (débit) | 9755 : 12 canaux DDR5-6400 | 6980P + MRDIMM-8800 : +24 % mesuré (HPCG) |
| Prix amiral | 12 984 $ (9755) | 17 800 $ (6980P) |
| Bande passante/socket | ~460 Go/s (théorique 12×6400) | ~460 Go/s RDIMM, ~630 Go/s MRDIMM |

Ordres de grandeur théoriques : 12 canaux × 6400 MT/s × 8 octets =
614 Go/s crête par socket (formule : canaux × débit × 8). En pratique
comptez 70-80 % de cette valeur en débit soutenu (STREAM).

Verdict : **OLTP → fréquence et cache (les deux se valent, comparez au
prix) ; analytique/HPC → avantage Intel avec MRDIMM aujourd'hui, avantage
AMD avec Venice + MRDIMM-12800 demain (2027).** Et Oracle Database se
licencie au cœur × 0,5 (facteur x86) : le nombre de cœurs **physiques**
est une donnée financière directe (section 25).

## 22. Comparatif EPYC vs Xeon : HPC et inférence IA sur CPU

- **HPC (AVX-512)** : Turin a enfin un chemin 512 bits complet ; Granite
  Rapids a l'avantage historique du vectoriel Intel + MRDIMM. Les deux
  sont crédibles ; le workload tranche (CFD, FEA, chimie quantique…).
  Chiffres constructeur : AMD revendique +60 % par cœur vs Xeon 5e gén en
  FEA/CFD ; Intel revendique 2,1× en HPC vs sa propre génération précédente
  et jusqu'à 5,5× en inférence IA vs EPYC (chiffres fournisseur — à
  relativiser, section 2).
- **Inférence IA sur CPU** : **AMX est l'atout maître d'Intel.** Sur les
  modèles quantifiés (INT8/BF16), un Xeon 6900P avec AMX surclasse
  largement un EPYC à cœurs équivalents. AMD répond par AVX-512/VNNI mais
  l'écart mesuré reste en faveur d'Intel sur l'inférence CPU pure.
  Si l'inférence est votre workload principal, **privilégiez soit des GPU,
  soit du Xeon P-core avec AMX** — un EPYC dense n'est pas l'outil adapté.
- **Bande passante mémoire** : le nerf de la guerre pour l'inférence de
  grands modèles sur CPU. MRDIMM-8800 (Intel, aujourd'hui) puis
  MRDIMM-12800 (Venice, 2027) changent la donne plus que les cœurs eux-mêmes.

---

## 23. 1 socket, 2 sockets, 4 sockets : panorama

| Config | Sockets | Cœurs max (Turin) | Canaux mémoire | Usage type |
|---|---|---|---|---|
| 1S (mono-socket) | 1 | 192 c. / 384 th. | 12 | virtualisation PME, DB petite/moyenne, edge |
| 2S (bi-socket) | 2 | 384 c. / 768 th. | 24 | standard datacenter, virtualisation dense |
| 4S (quad) | 4 | 768 c. / 1536 th. | 48 | DB monolithiques, in-memory |
| 8S | 8 | 688 P-cores (Intel 6700P) | 64 | niche extrême |

Réalité du marché : **> 90 % des serveurs vendus sont 1S ou 2S.** Le 4S
est une niche (bases de données monolithiques type SAP HANA, Oracle) et le
8S une curiosité. Ce guide traite le 4S pour que vous sachiez quand (ne pas)
le choisir, pas pour vous le vendre.

## 24. NUMA expliqué simplement

NUMA (Non-Uniform Memory Access) : dans un serveur 2S, chaque CPU possède
**sa** mémoire locale (ses 12 canaux). Accéder à la mémoire de l'autre
socket coûte un détour par l'interconnexion inter-socket (Infinity Fabric
chez AMD, UPI chez Intel).

Schéma ASCII (serveur 2S) :

```
        +-----------------+                    +-----------------+
        |   SOCKET 0      |                    |   SOCKET 1      |
        |  192 coeurs     |<---- UPI / IF ---->|  192 coeurs     |
        |  (noeud NUMA 0) |   ~200-250 ns      |  (noeud NUMA 1) |
        +--------+--------+                    +--------+--------+
                 | 12 canaux                              | 12 canaux
                 | ~80-90 ns                              | ~80-90 ns
          +------+-------+                          +-----+--------+
          |  RAM locale  |                          |  RAM locale  |
          |  (rapide)    |                          |  (rapide)    |
          +--------------+                          +--------------+
```

Règles pratiques :
- **Latence locale ~80-90 ns, distante ~200-250 ns** (ordres de grandeur,
  vérifiés via analyses publiques au 27/09/2026).
- L'OS voit 2 nœuds NUMA ; un bon hyperviseur place chaque VM sur **un
  seul nœud** (vNUMA). Une VM à cheval sur 2 nœuds perd 10 à 30 % de
  performance mémoire.
- En 1S, pas de NUMA inter-socket : plus simple, plus prévisible. C'est
  une des raisons du succès du mono-socket dense moderne.

## 25. Coûts cachés : les licences par socket et par cœur

C'est le poste qui fait le plus mal quand on l'oublie. Principes
(licences = matière mouvante : **toujours vérifier auprès de l'éditeur**) :

| Éditeur / produit | Métrique usuelle (indicatif, à vérifier) |
|---|---|
| VMware vSphere | par cœur, minimum 16 cœurs/CPU |
| Microsoft SQL Server | par cœur, minimum 4 cœurs/CPU |
| Oracle Database | par cœur × facteur 0,5 (x86) |
| Veeam, divers | par socket ou par VM |

Exemples d'impact :
- Un serveur 2× 64 cœurs = 128 cœurs à licencier. Le même en 2× 32 cœurs
  = 64 cœurs. **Diviser les cœurs par 2 peut diviser la facture logicielle
  par 2**, pour une perte de performance souvent < 30 % sur des workloads
  non parallélisés à l'extrême.
- Les références « F » (peu de cœurs, haute fréquence) sont des **armes
  anti-licences** : un 9175F (16 cœurs à 5,0 GHz) fait tourner une base
  Oracle plus vite qu'un 64 cœurs dense, pour 4× moins de licences.
- En 4S, les licences explosent (4 sockets × N cœurs) : le TCO logiciel
  dépasse vite le TCO matériel. D'où la règle : **on ne choisit le 4S que
  si le workload l'exige techniquement** (section 27).

## 26. Coût total 1S vs 2S : exemple chiffré (indicatif)

Hypothèse : virtualisation, besoin de ~96 cœurs et 1,5 To de RAM.

Option A — 1 serveur 2S : 2× EPYC 9455 (48 c., 300 W, 4 819 $ pièce).
Option B — 2 serveurs 1S : 2× EPYC 9655P (96 c., 400 W, 10 811 $ pièce).

| Poste (indicatif) | Option A (1× 2S) | Option B (2× 1S) |
|---|---|---|
| CPU | 2 × 4 819 = 9 638 $ | 2 × 10 811 = 21 622 $ |
| Cœurs totaux | 96 | 192 |
| Châssis/licences par socket | 1 serveur à gérer | 2 serveurs (2× ports, 2× contrats) |
| NUMA | oui (2 nœuds) | non (1 nœud/serveur) |
| Tolérance de panne | aucune (1 châssis) | 1 nœud survivant |

Lecture : l'option B coûte plus cher en CPU mais donne 2× plus de cœurs
et une redondance. **Il n'y a pas de bonne réponse universelle** : en
cluster Proxmox/VMware, 2× 1S est souvent préférable (HA native) ; pour
une base monolithique, 1× 2S avec CPU « F » gagne sur les licences.

## 27. Quand choisir 4 sockets (et quand l'éviter)

Choisissez 4S si **toutes** ces conditions sont réunies :
1. Le workload ne se distribue pas (base monolithique, SAP HANA, grosse
   instance Oracle/SQL Server > 2 To de RAM active).
2. Vous avez besoin de plus de 3 To de RAM **dans une seule image OS**
   (au-delà : scale-out ou in-memory distribué).
3. Le budget licences a été validé **avant** le budget matériel.

