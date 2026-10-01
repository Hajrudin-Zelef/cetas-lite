---
id: collect-250926-servers-hardware/servers-hardware/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5-3
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "gpu"]
source: docs/RAG/clean4/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [56, 89]
sha256: a4f87f28c00e3d9f246b15040beee806d873c67e6bbe051c1019b2373e1942e9
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

Ici, les 9965 cœurs de l'EPYC 192 affichent un score multicœur de 11,199 20,217, ce qui permet une puissance de traitement considérable, mais reste en retrait par rapport aux configurations Genoa (17,916 9965) et Bergamo (9755 128) à double socket. Cela suggère que même si le 9965 est un modèle performant, Genoa et Bergamo restent optimisés pour les charges de travail de traitement parallèle les plus extrêmes. Les 1,641 cœurs de l'EPYC 11,800 démontrent des gains d'efficacité, surpassant légèrement le XNUMX dans les scores monocœur (XNUMX XNUMX) et multicœur (XNUMX XNUMX), ce qui indique qu'il est parfaitement adapté aux tâches évolutives et multithread sans la surcharge d'un nombre de cœurs plus important.
Le 9575F (avec SMT désactivé) obtient un score solide de 1,865 64 en mono-cœur, ce qui montre son potentiel pour les applications qui bénéficient de performances mono-thread élevées. En revanche, son score multi-cœur 13,219 cœurs de XNUMX XNUMX, bien que modeste, le rend plus adapté aux charges de travail spécialisées plutôt qu'aux tâches de traitement parallèle les plus lourdes.
| Geekbench 6 | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) | 
| Single Core | 1,453 | 1,641 | 1,865 | 2,048 | 1,723 | 
| Multi-Core | 11,199 | 11,800 | 13,219 | 20,217 | 17,916 | 
Cinebench est un outil d'analyse comparative largement utilisé qui mesure les performances des processeurs et des GPU utilisant Maxon Cinema 4D pour le rendu. Il fournit un score permettant de comparer les performances de différents systèmes et composants. Nous avons exécuté quatre versions populaires de Cinebench afin que vous puissiez comparer les résultats des classements populaires en ligne.
Les résultats de Cinebench R23 montrent que l'EPYC 9755 obtient un score multicœur élevé de 131,846 116,744, surpassant les 102,125 1,294 de Genoa et les XNUMX XNUMX de Bergamo, ce qui le positionne comme un choix idéal pour les tâches de rendu intensif dans des environnements hautement parallèles. Genoa est en tête avec XNUMX XNUMX points en performances monocœur, ce qui lui donne un léger avantage dans les applications qui privilégient les opérations monothread plus rapides.
Dans les tests Cinebench 2024 mis à jour, les 9965 et 9755 continuent d'obtenir de bons résultats avec respectivement 4,845 5,921 et 3 XNUMX points, offrant des scores multithread solides adaptés à des tâches telles que le rendu XNUMXD complexe et les flux de travail de création de contenu.
| Test | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) |  | 
| Cinebench R20 |  |  |  |  |  |  | 
| Processeur | N/D | 43,765 pts | N/D | N/D | N/D |  | 
| Cinebench R23 |  |  |  |  |  |  | 
| Processeur (multicœur) | N/D | 131,846 pts | 111,149 pts | 116,744 pts | 102,125 XNUMX points |  | 
| Processeur (monocœur) | N/D | 1,400 pts | 1,052 pts | 1,294 pts | 1,089 XNUMX points |  | 
| Cinebench 2024 |  |  |  |  |  |  | 
| Processeur (multicœur) | 4,845 pts | 5,921 pts | 4,324 | N/D | N/D |  | 
| Processeur (monocœur) | 77 pts | 84 pts | 103 pts | N/D | N/D |  | 
y-cruncher 0.8.3.9522 est un programme multithread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des milliards de chiffres. Depuis son lancement en 2009, elle est devenue une application d'analyse comparative et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel.
Ici, l'EPYC 9575F (SMT désactivé) atteint les temps de calcul les plus rapides pour les calculs à 1 et 2.5 milliards de chiffres, avec respectivement 4.476 et 10.067 secondes. Cette vitesse surpasse même les modèles 9965 et 9755 à haut cœur, démontrant son efficacité pour les tâches de haute précision qui ne nécessitent pas un nombre de threads extrême.
Les processeurs 9965 et 9755 offrent également de bons temps, surpassant notamment ceux de Gênes et de Bergame à 10 milliards de chiffres, en un peu plus de 41 secondes contre 51 secondes pour Gênes. Cela montre que les processeurs EPYC basés à Turin sont bien optimisés pour les applications à calcul intensif, démontrant des performances évolutives qui permettent un contrôle plus précis de la gestion des threads et des ressources. Le processeur 64F à 9575 cœurs a vraiment brillé ici avec sa vitesse d'horloge incroyablement élevée, montrant une forte avance sur les cœurs des échantillons 9755 ou 9965 dans les gammes de y-cruncher que nous avons examinées.
| y-cruncher Temps de calcul total (Plus bas, c'est mieux) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, 128t) | AMD EPYC 9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) | 
| 1 milliard | 7.346 secondes | 7.747 secondes | 5.408 secondes | 4.476 secondes | 8.882 secondes | 9.184 secondes | 
| 2.5 milliard | 13.661 secondes | 14.113 secondes | 11.376 secondes | 10.067 secondes | N/D | N/D | 
| 5 milliard | 23.211 secondes | 22.820 secondes | 20.177 secondes | 20.030 secondes | N/D | N/D | 
| 10 milliard | 41.750 secondes | 41.512 secondes | 40.767 secondes | 41.518 secondes | 51.071 secondes | 55.683 secondes | 
| 25 milliard | 115.091 secondes | 98.981 secondes | 103.650 secondes | 104.737 secondes | N/D | N/D | 
Ce test de performance Y-Cruncher utilise les formules Bailey-Borwein-Plouffe (BBP) pour calculer des chiffres hexadécimaux massifs de Pi, mesurant le temps de calcul total du processeur, son utilisation et son efficacité multicœur. Les résultats fournissent une analyse détaillée des performances des processeurs AMD Turin dans le cadre de ces calculs intensifs.
L'EPYC 9575F avec SMT désactivé excelle en efficacité mono-thread, réalisant 1 BBP en seulement 0.179 seconde avec une efficacité multi-cœur de 44.73 %. Cela démontre son efficacité dans les scénarios où l'optimisation mono-thread est prioritaire. En comparaison, le 9965, avec ses 192 cœurs, affiche une efficacité multi-cœur inférieure de 4.61 % dans ce test mono-BBP, car son nombre élevé de cœurs introduit une latence dans les scénarios à faible thread.
L'efficacité multicœur devient plus évidente avec des charges de travail plus élevées, comme 10 et 100 BBP. L'EPYC 9575F avec SMT désactivé conserve une solide avance avec une efficacité de 86.59 % à 10 BBP, le terminant en 0.896 seconde. Pour 100 BBP, le 9575F continue d'exceller, atteignant près de 99 % d'efficacité et affichant des performances équilibrées dans des charges de travail hautement parallèles, terminant l'exécution en 8.065 secondes. Dans ce test, le 9965 atteint une efficacité de 85.10 %, tandis que les 9755 et 9575 maintiennent des efficacités multicœurs compétitives de 94 % et 96.83 %, respectivement, démontrant que les processeurs EPYC basés sur Turin peuvent gérer efficacement les charges de travail à faible et à haut thread.
| référence | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | AMD EPYC9575F (64c, SMT désactivé) | 
| 1 BBP |  |  |  |  | 
| 10 BBP |  |  |  |  | 
| 100 BBP |  |  |  |  | 
Le test de mémoire intégré à l'utilitaire 7-Zip mesure les performances du processeur et de la mémoire d'un système pendant les tâches de compression et de décompression, indiquant dans quelle mesure le système peut gérer des opérations gourmandes en données. Nous exécutons ce test avec une taille de dictionnaire de 128 Mo lorsque cela est possible.
