---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb-2
title: "fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "gpu", "intel", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb.md
source_anchor: ""
source_lines: [35, 88]
sha256: 37ad14b95e8696754e04e2967f184223b3fcc151c8e1746417c6bbe253834061
---

# fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb

La façade abrite les huit baies de disque 2.5″ remplaçables à chaud et, dans notre modèle d'examen, l'espace où quatre baies de disque SATA/SAS en option peuvent être ajoutées pour une extension ultérieure. Le panneau avant comprend également des commandes et indicateurs essentiels, tels que des boutons d'alimentation et UID, des LED d'état pour surveiller l'état du système et un port USB 3.0 pour un accès rapide aux périphériques externes.
L'arrière du serveur comprend deux emplacements PCIe 5.0 x16 pour des cartes pleine hauteur et pleine longueur (l'un est double largeur) et l'emplacement OCP, permettant des extensions de réseau ou de stockage à haut débit. Les alimentations redondantes de niveau Titanium du serveur garantissent fiabilité et fonctionnement ininterrompu avec des options allant jusqu'à 1600 XNUMX W. Une connectivité supplémentaire pour la gestion USB et OoB se trouve également à l'arrière.
À l'intérieur du SYS-112H-TN, l'agencement est organisé pour optimiser le flux d'air et maintenir une gestion thermique efficace. Le serveur dispose de huit ventilateurs internes pour assurer un refroidissement constant de tous les composants critiques. Le processeur unique Intel Xeon 6 Scalable est situé au centre, flanqué de 16 emplacements DIMM prenant en charge la mémoire DDR5. La carte mère comprend également deux emplacements NVMe M.2.
Performances du Supermicro Hyper 1U SYS-112H-TN
Notre unité d'examen Supermicro Hyper 1U SYS-112H-TN est configurée avec les éléments suivants :
- Processeur : Intel Xeon 6780E 144 cœurs
- RAM: 512GB DDR5
- SSD : SSD de centre de données Micron 7450 NVMe
- Windows 10 64 bits (10.0.20348)
Dans nos tests, nous comparons les performances du Supermicro Hyper 1U SYS-112H-TN avec les résultats de notre examen initial du serveur de lancement d'Intel équipé de la même famille de processeurs Xeon 6. Cette comparaison est utile car elle examine les performances de ces processeurs sur différentes architectures de serveur, ce qui nous permet de fournir des résultats de référence plus précieux. Il peut nous indiquer l'impact de la conception et de la configuration du serveur sur les performances et l'efficacité globales des applications réelles.
Mélangeur OptiX 4.0 / 4.1
Blender est une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Le Supermicro Hyper 1U 112H-TN chargé avec un seul Xeon 6780E affiche des performances solides mais est, comme on s'y attendait, en deçà du serveur Intel (lancement) doté de deux processeurs Xeon. Par exemple, dans la scène « Monster », la configuration Supermicro atteint 781.42 échantillons par minute, tandis que la configuration double Xeon 6780E atteint 1,410.463 3 échantillons par minute. Cela montre que même si la configuration à un seul processeur est compétente pour le rendu 3D, la configuration à deux processeurs double presque les performances, ce qui la rend mieux adaptée aux charges de travail XNUMXD plus exigeantes. Le Supermicro surpasserait probablement le serveur Intel de lancement s'il était équipé de deux processeurs.
| Processeur Blender 4.0 | Supermicro Hyper 1U 112H-TN (1x Xeon 6780E, 512 Go DDR5) | 2x Xeon 6780E (256 Go DDR5) | 2x Xeon 6766E (256 Go DDR5) | 
| Monster | 781.42 | 1410.463 | 1297.715 | 
| Brocanteur | 514.658 | 862.418 | 777.716 | 
| Salle de classe | 370.52 | 696.543 | 628.960 | 
| Processeur Blender 4.1 |  |  |  | 
| Monster | 764.112 | N/D | N/D | 
| Brocanteur | 511.872 | N/D | N/D | 
| Salle de classe | 363.672 | N/D | N/D | 
Geekbench 6
Geekbench 6 est un outil d'évaluation des performances multiplateforme qui mesure les performances globales d'un système. Le navigateur Geekbench permet de comparer n'importe quel système à cet outil.
Le score CPU Single-Core de 1,154 15,167 reflète des performances monothread décentes, suffisantes pour les tâches qui dépendent de la vitesse de chaque cœur. Le score CPU Multi-Core de 6780 6780 met en évidence la capacité du Xeon XNUMXE à gérer efficacement les tâches multithread. Cela rend le système bien adapté aux charges de travail exploitant plusieurs cœurs, telles que la virtualisation, la gestion de bases de données et les environnements multi-utilisateurs. Bien que ces scores soient solides, ils suggèrent que le Xeon XNUMXE dans cette configuration est davantage orienté vers les tâches de traitement parallèle plutôt que d'exceller dans les applications monocœur sensibles à la latence.
| Geekbench 6 (Plus c'est haut, mieux c'est) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | 
| Processeur monocœur | 1,154 | 
| Processeur multicœur | 15,167 | 
Cinebench R23
L'outil de référence Cinebench R23 évalue les performances du processeur d'un système en restituant une scène 3D complexe à l'aide du moteur Cinema 4D. Il mesure les performances monocœur et multicœur, offrant une vue complète des capacités du processeur dans la gestion des tâches de rendu 3D.
La configuration Supermicro Xeon a obtenu 92,516 6780 points dans le test multicœur, soit plus que la plate-forme double Xeon 67,984E de lancement d'origine (XNUMX XNUMX pts).
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | 2x Xeon 6780E (256 Go DDR5) | 2x Xeon 6766E (256 Go DDR5) | 
| Processeur multicœur | 92,516 | 67,984 pts | 64,326 pts | 
| Processeur monocœur | 888 pts | 873 pts | 793 pts | 
| Rapport PM | 104.20 x | 77.91 x | 81.10 x | 
Cinebench 2024
Cinebench 2024 étend les capacités de référence de R23 en ajoutant une évaluation des performances du GPU. Il continue de tester les performances du processeur mais inclut également des tests qui mesurent la capacité du GPU à gérer les tâches de rendu.
Le Supermicro Hyper 112H-TN obtient 2,941 6780 pts en multicœur, surpassant la configuration double Xeon 2,687E, qui a obtenu XNUMX XNUMX pts. Il s'agit d'un autre exemple où la configuration d'un processeur unique montre une optimisation robuste dans le serveur Supermicro, offrant des performances supérieures dans des applications multithread spécifiques.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | 2x Xéon 6780E (256 Go DDR5) | 2x Xeon 6766E (256 Go DDR5) | 
| Processeur multicœur | 2,565 pts | 2,687 pts | 2,347 pts | 
| Processeur monocœur | 53 pts | 43 pts | 45 pts | 
| Rapport PM | 48.38 x | 62.85 x | 52.65 x | 
Y-Cruncher
Y-cruncher est une application populaire d'analyse comparative et de tests de résistance lancée en 2009. Ce test est multithread et évolutif, calculant Pi et d'autres constantes jusqu'à des milliards de chiffres. Plus vite, c'est mieux dans ce test.
Dans ce test, la configuration de Supermicro était plus lente, en particulier dans les calculs plus importants, par rapport au serveur de lancement Intel doté de deux processeurs. Par exemple, le calcul de Pi à 50 milliards de chiffres prend 680.090 secondes sur la configuration Supermicro, alors que la configuration double Xeon 6780E le complète en 565.913 secondes.
| Y-Cruncher (0.8.3.9) (plus bas est mieux) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Xéon 6780E (256 Go DDR5) | Xéon 6766E (256 Go DDR5) | 
| 1 milliard | 9.754 secondes | 6.927 secondes | 7.254 secondes | 
| 2.5 milliard | 25.215 secondes | 17.898 secondes | 19.507 secondes | 
| 5 milliard | 55.242 secondes | 38.454 secondes | 41.116 secondes | 
| 10 milliard | 118.657 secondes | 81.146 secondes | 87.403 secondes | 
| 25 milliard | 315.085 secondes | 217.530 secondes | 238.813 secondes | 
| 50 milliard | 680.090 secondes | 565.913 secondes | 502.245 secondes | 
Voici les résultats de la version 0.8.5.9 :
| Y-Cruncher (0.8.5.9) (plus bas est mieux) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | 
| 1 milliard | 8.757 secondes | 
