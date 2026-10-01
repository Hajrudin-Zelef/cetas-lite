---
id: collect-261001-general-networking/general-networking/fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce-1
title: "fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["blackwell", "distribution", "ethernet", "gpu", "lpddr5x", "nvfp4", "nvidia", "nvlink"]
source: docs/RAG/collect-261001-general-networking/fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce.md
source_anchor: ""
source_lines: [1, 45]
sha256: 054cfe523c3a53d918d9934f15da5bc2efe62a039c72e15d684376004a219c99
---

# fr-review-asus-expertcenter-pro-et900n-g3-review-7a4e59ce

L'ASUS ExpertCenter Pro ET900N G3 est la deuxième station GB300 DGX que nous avons testée et la première que nous avons pu manipuler en laboratoire ; nos tests de la MSI XpertStation WS300 ont été effectués à distance. Elle utilise la même puce NVIDIA que celle testée dans la MSI XpertStation WS300 : un processeur graphique Grace Blackwell Ultra Desktop Superchip avec 72 cœurs Grace, un GPU B300, 252 Go de mémoire HBM3e, 496 Go de mémoire LPDDR5X et une carte réseau ConnectX-8 SuperNIC avec deux ports 400 GbE. ASUS a intégré cette plateforme dans un boîtier tour compact conçu pour l'IA de bureau, en y ajoutant quelques détails absents de la WS300 : deux poignées de transport sur le dessus, un ventilateur dédié aux baies optiques ConnectX-8 et une alimentation de 1 600 W certifiée 80 PLUS Titanium.
ASUS nous a envoyé le PC portable ET900N G3 pour une semaine de tests de performances, de prises de photos et de vérifications physiques, avec une RTX PRO 2000 installée. Le test du WS300 aborde la plateforme elle-même, la carte mère B300, Grace, NVLink-C2C, MIG et l'utilité de la station DGX. Ce test-ci porte sur les caractéristiques techniques du processeur graphique ASUS et compare ses performances à celles du premier PC portable GB300 que nous avons testé.
Spécifications de l'ASUS ExpertCenter Pro ET900N G3
| Spécifications | ASUS ExpertCenter Pro ET900N G3 | 
|---|---|
| En savoir plus sur la plateforme |  | 
| Superpuce | NVIDIA GB300 Grace Blackwell Superpuce de bureau ultra-performante | 
| Processeur | NVIDIA Grace, 72 cœurs Arm Neoverse V2 | 
| GPU | NVIDIA Blackwell Ultra (B300), jusqu'à 20 PFLOPS NVFP4 avec sparsité | 
| Interconnexion CPU-GPU | NVLink-C2C, 900 Go/s bidirectionnel | 
| Mémoire |  | 
| Mémoire cohérente | 748GB | 
| Mémoire GPU | 252 Go HBM3e, 7.1 To/s | 
| Mémoire CPU | 496 Go LPDDR5X, 396 Go/s, 4 ports SOCAMM | 
| Stockage |  | 
| Disques de démarrage | 2 emplacements M.2 2280 PCIe 5.0 x4 (clé M), équipés de 2 disques NVMe de 2 To en RAID 1. | 
| Entraînements d'expansion | 2 emplacements M.2 2280 (clé M), PCIe 6.0 x4 (norme ASUS), carte mère ouverte. | 
| Networking |  | 
| SuperNIC | NVIDIA ConnectX-8, 2 ports QSFP112 400 GbE | 
| Ethernet | 1x Marvell 10GbE 1x Realtek 1GbE (BMC) | 
| Expansion |  | 
| Emplacements PCIe | 1x PCIe 5.0 x16 2 ports PCIe 5.0 x16 (8 signaux x8) | 
| Carte graphique additionnelle | Jusqu'à une carte NVIDIA RTX PRO Blackwell ; le modèle testé était équipé d'une RTX PRO 2000 Blackwell. | 
| I / O |  | 
| Avant | 2 ports USB Type-C 10 Gbit/s 2 ports USB 10 Gbit/s de type A 1x USB 2.0 Prises casque et microphone | 
| Arrière | 4 ports USB 10 Gbit/s de type A 1x Micro-USB COM (console série BMC) 1x Mini DisplayPort (BMC) 3 prises audio | 
| Gestion et sécurité |  | 
| BMC | ASPEED AST2600 avec firmware AMI MegaRAC, IPMI et Redfish | 
| Sécurité | TPM 2.0 intégré | 
| Puissance et refroidissement |  | 
| Alimentation | 1 alimentation ATX de 1 600 W, certification 80 PLUS Titanium | 
| Refroidissement | Refroidissement liquide AIO en circuit fermé (selon ASUS) avec plaques froides sur le GB300, le SOCAMM et le ConnectX-8 ; radiateurs avant et supérieur, ventilateur arrière du châssis, ventilateur dédié au-dessus des cages QSFP112 | 
| Température de fonctionnement | 10C à 35C | 
| Logiciel et facteur de forme |  | 
| Système d'exploitation | Ubuntu avec les outils de développement IA de NVIDIA ; ASUS annonce un support pour le développement IA sous Windows. | 
| Dimensions | 584 x 232 x 565 mm (23.0 x 9.1 x 22.3 pouces), selon ASUS | 
| Poids | 27 kg net, 32 kg brut | 
Concevoir et construire
ASUS présente l'ET900N G3 comme une station de travail professionnelle, avec un châssis argent brossé, une façade en mesh foncé, un capot et un pied chromés, et le logo ASUS en bas de la façade. Ses dimensions sont conformes à la plateforme : 584 x 232 x 565 mm selon ASUS, soit légèrement plus haut et un peu moins large que le WS300 (529 mm de hauteur, 570 mm de profondeur et 246 mm de largeur). Avec ses 27 kg, ce n'est pas un ordinateur que l'on déplace à la légère, et c'est là qu'intervient la première idée propre à ASUS : deux poignées métalliques sur le bord supérieur, une à l'avant et une à l'arrière, permettent à deux personnes de soulever la tour et de la poser sur un bureau ou un chariot sans avoir à saisir les panneaux. Elles peuvent paraître un peu incongrues sur un ordinateur de bureau, jusqu'à ce qu'on ait déjà eu à transporter une tour GB300 dans un laboratoire.
Le panneau avant comprend le bouton d'alimentation, deux ports USB Type-A 10 Gbit/s, deux ports USB Type-C 10 Gbit/s, un port USB 2.0 et des prises casque et microphone séparées, le tout disposé en bande verticale en haut à droite de la grille. Le panneau latéral est une large plaque ventilée, avec une longue colonne perforée à l'avant pour l'entrée d'air du radiateur et une grille carrée plus petite à l'arrière, alignée avec les ventilateurs du système. L'absence de fenêtre et d'éclairage correspond aux attentes de la clientèle visée.
À l'arrière, la configuration sépare les entrées/sorties de la station de travail du matériel serveur situé en dessous. Le bloc supérieur comprend quatre ports USB Type-A 10 Gbit/s, les prises RJ45 10 Gbit/s et 1 Gbit/s, les trois prises audio, le port console micro-USB du BMC et le Mini DisplayPort utilisé par le BMC pour la configuration. Les deux cages QSFP112 pour le ConnectX-8 se trouvent en haut de la plaque d'E/S ; un ventilateur d'extraction est placé à côté ; les caches des trois emplacements d'extension sont disposés au centre ; et l'alimentation avec son connecteur C19 est en bas. Comme sur le WS300, le Mini DisplayPort est une sortie du BMC pour la configuration initiale et le dépannage ; pour utiliser cette machine comme un bureau, il faut une carte graphique RTX PRO.
À l'intérieur du ET900N G3
Une fois le panneau latéral retiré, l'ET900N G3 apparaît très proche du WS300, ce qui est logique puisqu'ils partagent la même carte mère NVIDIA. Le Superchip GB300 est placé sous une plaque froide en cuivre, à gauche du châssis. Les modules SOCAMM sont installés sous leurs propres plaques de cuivre à côté, et le ConnectX-8 sous un troisième bloc près des connecteurs d'E/S arrière. Des tubes tressés relient ces blocs à un collecteur de distribution monté sur la traverse du châssis, puis alimentent les radiateurs. Les trois emplacements PCIe pleine longueur se trouvent sous le Superchip, l'alimentation occupe le coin inférieur avant et un cache métallique dissimule le chemin des câbles en bas.
L'agencement interne est identique à celui du WS300 : un radiateur monté verticalement derrière la grille avant et un second en haut, chacun équipé d'une rangée de ventilateurs, plus un ventilateur de châssis à l'arrière. Les conduites de liquide de refroidissement sont gainées et fixées à la traverse par des attaches auto-agrippantes. Le collecteur regroupe les conduites provenant des quatre plaques froides, de sorte que seules deux conduites atteignent chaque radiateur.
Le fan d'optique
