---
id: collect-261001-ia-llm/ia-llm/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662-4
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: ia-llm
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "gpu", "open source"]
source: docs/RAG/collect-261001-ia-llm/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [101, 146]
sha256: 70d03a5a495bbac2041a41401853a8b5125472af9c73df2fa82b52df71b33c43
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

| Single Core | 2,806 | 1,453 | 1,641 | 1,865 | 2,048 | 1,723 | 
| Multi-Core | 28,645 | 11,199 | 11,800 | 13,219 | 20,217 | 17,916 | 
Test de vitesse Blackmagic RAW
Nous avons effectué le test de vitesse Blackmagic RAW pour évaluer la capacité du PowerEdge R7715 à gérer les tâches de décodage Blackmagic RAW en utilisant uniquement le processeur, sans GPU installé. Ce test mesure les performances à différentes résolutions et niveaux de compression.
Le R7715 a obtenu un score de 171 FPS pour le contenu 8K, ce qui démontre des performances CPU assez solides pour le traitement vidéo haute résolution, même sans accélération GPU.
Maxon Cinébench
Cinebench est un outil d'analyse comparative largement utilisé qui mesure les performances des processeurs et des cartes graphiques (CPU) utilisant Maxon Cinema 4D pour le rendu. Il fournit un score permettant de comparer les performances de différents systèmes et composants. Nous avons testé quatre versions populaires de Cinebench afin que vous puissiez comparer les résultats sur les classements en ligne les plus populaires.
Ici, l'AMD EPYC 9665P a démontré ses capacités de processeur monosocket hautes performances. Avec un score multicœur de 121,254 23 points dans Cinebench R9755, il se démarque des systèmes comparables, devançant de peu le score de 131,846 2 points de l'EPYC 96, mais surpassant les systèmes Genoa (2P/128c) et Bergamo (9665P/1,845c). Cependant, la caractéristique la plus remarquable ici est la performance monocœur, où le XNUMXP a obtenu un score de XNUMX XNUMX points, nettement supérieur à tous les autres processeurs testés.
Dans Cinebench 2024, le 9665P continue d'exceller, obtenant 7,501 9965 points aux tests multicœurs, surpassant ainsi les 9755, 9575 et 109F. Son score monocœur de 9965 points est également impressionnant, largement devant le 77 (9755 points) et le 84 (9665 points). Le XNUMXP offre un équilibre parfait entre vitesse monocœur et puissance multicœur.
| Test | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) |  | 
| Cinebench R23 |  |  |  |  |  |  |  | 
| Processeur (multicœur) | 121,254 pts | N/D | 131,846 pts | 111,149 pts | 116,744 pts | 102,125 XNUMX points |  | 
| Processeur (monocœur) | 1,845 pts | N/D | 1,400 pts | 1,052 pts | 1,294 pts | 1,089 XNUMX points |  | 
| Cinebench 2024 |  |  |  |  |  |  |  | 
| Processeur (multicœur) | 7,501 pts | 4,845 pts | 5,921 pts | 4,324 | N/D | N/D |  | 
| Processeur (monocœur) | 109 pts | 77 pts | 84 pts | 103 pts | N/D | N/D |  | 
croque-y
y-cruncher est un programme multithread et évolutif capable de calculer Pi et d'autres constantes mathématiques jusqu'à des milliers de milliards de chiffres. Depuis son lancement en 2009, il est devenu une application de benchmarking et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel informatique.
Malgré un nombre de cœurs inférieur à celui des autres processeurs testés, le 9665P a affiché de solides performances sur la plupart des calculs. Son temps de 6.836 secondes pour le calcul à 1 milliard de chiffres est honorable, même s'il est inférieur à celui de l'AMD EPYC 9575F (64c), qui a réalisé le meilleur temps avec 4.476 secondes. À mesure que le nombre de chiffres augmente, le 9665P maintient des performances compétitives, excellant particulièrement dans le test à 10 milliards de chiffres, qu'il réalise en 51.851 secondes. Ce résultat est comparable à celui d'autres processeurs à nombre de cœurs plus élevé, tels que le 9965 (41.750 secondes) et le 9755 (41.512 secondes).
Cependant, pour des nombres de chiffres plus élevés (25 et 50 milliards dans ce cas), le 9665P met respectivement 140.339 et 303.842 secondes. Cela signifie probablement que, même si le 9665P gère bien les tâches de calcul lourdes, il commence à perdre du terrain à mesure que la complexité des charges de travail augmente, notamment par rapport aux processeurs dotés d'un nombre de cœurs plus élevé. Néanmoins, le 9665P reste un concurrent sérieux pour la plupart des charges de travail et offre d'excellentes performances pour son nombre de cœurs.
| y-cruncher Temps de calcul total (Plus bas, c'est mieux) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, 128t) | AMD EPYC 9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) | 
| 1 milliard | 6.836 secondes | 7.346 secondes | 7.747 secondes | 5.408 secondes | 4.476 secondes | 8.882 secondes | 9.184 secondes | 
| 2.5 milliard | 13.720 secondes | 13.661 secondes | 14.113 secondes | 11.376 secondes | 10.067 secondes | N/D | N/D | 
| 5 milliard | 25.795 secondes | 23.211 secondes | 22.820 secondes | 20.177 secondes | 20.030 secondes | N/D | N/D | 
| 10 milliard | 51.851 secondes | 41.750 secondes | 41.512 secondes | 40.767 secondes | 41.518 secondes | 51.071 secondes | 55.683 secondes | 
| 25 milliard | 140.339 secondes | 115.091 secondes | 98.981 secondes | 103.650 secondes | 104.737 secondes | N/D | N/D | 
| 50 milliard | 303.842 secondes | N/D | N/D | N/D | N/D | N/D | N/D | 
| 100 milliard | 707.391 secondes | N/D | N/D | N/D | N/D | N/D | N/D | 
Mixeur OptiX
Blender OptiX est une application de modélisation 3D open source. Ce benchmark a été réalisé à l'aide de l'utilitaire Blender Benchmark CLI. Le score est mesuré en échantillons par minute, les valeurs les plus élevées étant les meilleures.
Les résultats du benchmark Blender OptiX indiquent que l'EPYC 9665P offre de bonnes performances compte tenu de son nombre de cœurs, mais se situe derrière les modèles à cœurs plus élevés pour les tâches de rendu intensives. Dans la scène Monster, le 9665P atteint 1,026.50 9965 échantillons par minute, un résultat solide, mais loin des résultats du 2,558.43 (9755 2,606.54) et du 9665 (795.44 511.45). La tendance se poursuit dans les scènes Junkshop et Classroom, où le 2P atteint respectivement 128 et 2 échantillons par minute, bien en dessous des systèmes à double traitement comme Bergamo (96p/XNUMXc) et Genoa (XNUMXp/XNUMXc).
Bien que le 9665P offre des performances correctes pour la plupart des charges de travail, il est clair qu'il existe des processeurs plus adaptés, dotés d'un nombre de cœurs plus élevé. Cependant, le 9665P tient tête à l'EPYC 9575F, notamment dans le secteur des bric-à-brac, où ses performances sont quasiment identiques (795.44 contre 802.00).
| Blender 4.0 Échantillons de processeur par minute (plus c'est élevé, mieux c'est) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) | 
| Monster | 1,026.50 | 2,558.43 | 2,606.54 | 1,196.15 | 1,700.65 | 2,038.71 | 
| Brocanteur | 795.44 | 1,866.65 | 1,843.48 | 802.00 | 1,101.84 | 1,382.58 | 
| Salle de classe | 511.45 | 1,270.17 | 1,251.54 | 637.13 | 869.48 | 1,045.96 | 
Hammer DB TPROC-C
Nous avons également testé les performances des bases de données avec Hammer DB sur le PowerEdge R7715. Ce système a démontré d'excellentes performances OLTP sur toutes les bases de données testées sous la charge de travail HammerDB TPROC-C (basée sur TPC-C, 800 entrepôts).
| Moteur de base de données | Performance des transactions (TPM) | 
|---|---|
| MariaDB 11.4.4 | 3,600,000 | 
| MySQL 8.4.4 | 3,300,000 | 
| PostgreSQL 17.2 | 3,100,000 | 
| MariaDB 10.11.12 (MDEV-21923) | 2,950,000 | 
| MariaDB 10.6.22 | 2,850,000 | 
| MySQL 5.7.44 | 2,700,000 | 
