---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8-2
title: "fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "dram", "gpu", "intel", "valuation"]
source: docs/RAG/clean4/fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8.md
source_anchor: ""
source_lines: [58, 79]
sha256: a3dfd3eb0e52446061d98fb47ad940f5faece8b73f34b922a2ce4aefabdeb0aa
---

# fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8

| Processeur (multicœur) (points) | 3,215 | 2,014 | 
| Processeur (monocœur) (points) | 67 | 71 | 
| Rapport PM | 48.31x | 28.29x | 
Évaluation du processeur Geekbench
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Toutefois, il serait intéressant d'analyser les performances monocœur et multicœur, ainsi que les résultats du benchmark OpenCL. Un score élevé indique de meilleures performances. Précisons que nous n'avons examiné que les résultats du processeur, ce serveur ne disposant pas de carte graphique.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6 | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cœurs, 2.9 GHz) | HPE ProLiant DL320 (processeur Intel Xeon-G 4 de 6430e génération, 32 cœurs, 2.1 GHz) | 
| CPU Benchmark - Monocœur | 3,215 | 2,014 | 
| Référence CPU - Multi-Core | 67 | 71 | 
| Référence GPU – OpenCL | 48.31x | 28.29x | 
croque-y
y-cruncher est un programme multi-thread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des billions de chiffres. Depuis son lancement en 2009, il est devenu une application populaire d'analyse comparative et de test de résistance pour les overclockeurs et les passionnés de matériel.
| croque-y (Temps de calcul total) | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cœurs, 2.9 GHz) | HPE ProLiant DL320 (processeur Intel Xeon-G 4 de 6430e génération, 32 cœurs, 2.1 GHz) | 
| 1 milliard de chiffres (secondes) | 10.012 | 21.452 | 
| 2.5 milliard de chiffres (secondes) | 26.844 | 50.418 | 
| 5 milliards de chiffres (secondes) | 56.400 | N/D | 
| 10 milliard de chiffres (secondes) | 120.233 | 131.135 | 
| 25 milliards de chiffres (secondes) | 340.047 | N/D | 
Conclusion
Même s’il n’est pas déjà très clair, ce système est impressionnant. Les configurations système haut de gamme époustouflent avec des spécifications telles que 16 To de DRAM, la prise en charge de six petits GPU ou deux doubles largeurs et plus qu'assez d'espace pour une capacité de stockage importante. La boucle liquide est la clé qui permet de combiner puissance et densité.
Bien que les systèmes à quatre processeurs ne soient pas la norme dans la plupart des centres de données, ils continuent d'offrir une proposition de valeur tout à fait unique pour les charges de travail qui peuvent tirer parti de l'empreinte massive de la DRAM et de la puissance du processeur. Tenez compte du fait que ce boîtier peut également prendre quelques GPU, et vous auriez du mal à trouver un serveur capable de faire plus, dans un espace de rack de 2U.
Le HP ProLiant DL560 Gen11 met au premier plan la nécessité pour les entreprises de s'adapter aux systèmes de refroidissement modernes pour les serveurs puissants. Dans ce cas, nous avons affaire à un simple système en boucle fermée pour refroidir les quatre processeurs, ce qui répond très bien au besoin. Mais à mesure que les TDP progressent dans les processeurs et progressent dans les GPU, les centres de données devront planifier une boucle plus grande ou des systèmes alternatifs à base de liquide, il n'y a aucun moyen de l'éviter. Pour l’instant, le DL560 Gen11 constitue un grand pas en avant pour les plates-formes de calcul denses et une considération intéressante pour les charges de travail compatibles.
