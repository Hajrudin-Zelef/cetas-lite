---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45-2
title: "fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "gpu", "intel", "valuation"]
source: docs/RAG/clean4/fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45.md
source_anchor: ""
source_lines: [59, 114]
sha256: ab6c71ada25d332aed65c74e4ecb0a2d6b2d5674a2edafefaf22671b480b40d6
---

# fr-review-hpe-proliant-ml350-gen11-tower-server-review-dfc3ad45

Cinebench 2024 de Maxon est une référence de rendu CPU et GPU qui utilise tous les cœurs et threads du processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Comme cette configuration n'a pas de GPU, nous n'avons pas ces chiffres. Des scores plus élevés sont meilleurs.
| Cinebench 2024 | HPE ML350 Gen11 (Double Intel Xeon(r) Platinum 8480+, 112 cœurs, 2 GHz) | 
| Processeur (multicœur) (points) | 4,699 | 
| Processeur (monocœur) (points) | 76 | 
| Rapport PM | 61.44x | 
Évaluation du processeur Geekbench
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Toutefois, il serait intéressant d'analyser les performances monocœur et multicœur, ainsi que les résultats du benchmark OpenCL. Un score élevé indique de meilleures performances. Précisons que nous n'avons examiné que les résultats du processeur, ce serveur ne disposant pas de carte graphique.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6 | HPE ML350 Gen11 (Double Intel Xeon(r) Platinum 8480+, 112 cœurs, 2 GHz) | 
| CPU Benchmark - Monocœur | 1,939 | 
| Référence CPU - Multi-Core | 15,218 | 
| Référence GPU – OpenCL | N/D | 
croque-y
y-cruncher est un programme multi-thread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des billions de chiffres. Depuis son lancement en 2009, il est devenu une application populaire d'analyse comparative et de test de résistance pour les overclockeurs et les passionnés de matériel.
| croque-y (Temps de calcul total) | HPE ML350 Gen11 (Double Intel Xeon(r) Platinum 8480+, 112 cœurs, 2 GHz) | 
| 1 milliard de chiffres (secondes) | 5.136 | 
| 2.5 milliard de chiffres (secondes) | 29.889 | 
| 10 milliard de chiffres (secondes) | 65.194 | 
| 25 milliards de chiffres (secondes) | 186.841 | 
| 50 milliards de chiffres (secondes) | 413.722 | 
Compression à 7 zips
L'utilitaire populaire 7-Zip dispose d'un test de mémoire intégré qui démontre très bien les performances du processeur. Dans ce test, nous l'exécutons avec une taille de dictionnaire de 128 Mo lorsque cela est possible.
|  | HPE ML350 Gen11 (Double Intel Xeon(r) Platinum 8480+, 112 cœurs, 2 GHz) | 
|---|---|
| Compression |  | 
| Utilisation actuelle du processeur | 5,482 % | 
| Courant nominal/utilisation | 4.628 GIPS | 
| Courant | 253.724 GIPS | 
| Utilisation résultante du processeur | 5,475 % | 
| Évaluation/utilisation résultante | 4.628 GIPS | 
| Note résultante | 253.382 GIPS | 
| Décompression |  | 
| Utilisation actuelle du processeur | 6,219 % | 
| Courant nominal/utilisation | 3.745 GIPS | 
| Courant | 231.916 GIPS | 
| Utilisation résultante du processeur | 6,129 % | 
| Évaluation/utilisation résultante | 3.871 GIPS | 
| Note résultante | 237.259 GIPS | 
| Note totale |  | 
| Utilisation totale du processeur | 5,802 % | 
| Note totale/utilisation | 4.249 GIPS | 
| Note totale | 245.320 GIPS | 
Inférence UL Procyon AI
La suite de tests d' inférence IA Procyon de UL évalue les performances de différents moteurs d'inférence IA utilisant des réseaux neuronaux de pointe. Ces tests ont été exécutés uniquement sur le processeur. Chaque valeur représente un temps d'inférence moyen (plus la valeur est basse, meilleures sont les performances), et la dernière ligne indique un score global (plus la valeur est élevée, meilleures sont les performances).
|  | HPE ML350 Gen11 (Double Intel Xeon(r) Platinum 8480+, 112 cœurs, 2 GHz) | 
|---|---|
| Mobile Net V3 | 2.34 | 
| ResNet 50 | 5.76 | 
| Création V4 | 21.70 | 
| Deep Lab V3 | 23.00 | 
| YOLO V3 | 30.81 | 
| RÉEL-ESRGAN | 1535.27 | 
| Note globale | 191 | 
Conclusion
Dans l'ensemble, le HPE ProLiant ML350 Gen11 arrive sur le marché avec une multitude d'options de personnalisation qui séduiront toute entreprise cherchant à mettre à niveau les capacités de son serveur sans sacrifier l'espace physique qui serait généralement requis pour un système monté en rack. Lorsque nous avons commencé l'examen, ce système prenait en charge Intel Xeon Scalable de 4e génération, mais au cours du processus, ils ont également récupéré les puces de 5e génération, apportant encore plus de puissance à cette plate-forme.
De plus, ce serveur peut accueillir plus de deux douzaines de SSD NVMe hot-plug pour une capacité allant jusqu'à 368 To, 32 emplacements DIMM capables d'exécuter jusqu'à 8 To de RAM DDR5, une tonne d'extension PCIe et un système de ventilateur impressionnant à conserver. tout est cool et fonctionne efficacement. Les serveurs tour sont souvent négligés, car ils sont conçus « uniquement pour les PME ». Le HPE ProLiant ML350 Gen11 est un autre excellent rappel que les tours peuvent être bien plus encore ; ils font partie des systèmes les plus diversifiés et personnalisables disponibles pour répondre à l’ensemble de cas d’utilisation de plus en plus diversifiés d’aujourd’hui.
