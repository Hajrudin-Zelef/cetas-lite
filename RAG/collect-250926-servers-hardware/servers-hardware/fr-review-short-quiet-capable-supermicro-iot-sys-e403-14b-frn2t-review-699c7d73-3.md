---
id: collect-250926-servers-hardware/servers-hardware/fr-review-short-quiet-capable-supermicro-iot-sys-e403-14b-frn2t-review-699c7d73-3
title: "fr-review-short-quiet-capable-supermicro-iot-sys-e403-14b-frn2t-review-699c7d73"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["apache", "benchmark", "gpu", "intel", "nvidia", "open source"]
source: docs/RAG/clean4/fr-review-short-quiet-capable-supermicro-iot-sys-e403-14b-frn2t-review-699c7d73.md
source_anchor: ""
source_lines: [50, 89]
sha256: 3809378ad69e459f09a4c1ca60c3ebd41cadfdef5d05f4ca4469c5cb8d01e74e
---

# fr-review-short-quiet-capable-supermicro-iot-sys-e403-14b-frn2t-review-699c7d73

y-cruncher est un programme multithread et évolutif capable de calculer Pi et d'autres constantes mathématiques jusqu'à des milliers de milliards de chiffres. Depuis son lancement en 2009, il est devenu une application de benchmarking et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel informatique.
Lors des tests Y-Cruncher, le Supermicro SYS-E403-14B-FRN2T, équipé d'un processeur Intel Xeon 6521P (24 cœurs), a affiché des performances constantes d'une exécution à l'autre. Le calcul à 1 milliard de chiffres a été effectué en 10.2 secondes, passant à 27.6 secondes à 2.5 milliards de chiffres, 63.1 secondes à 5 milliards de chiffres et 134.4 secondes à 10 milliards de chiffres. À 25 milliards de chiffres, le système a terminé en 391.8 secondes. Ces résultats soulignent la capacité d'un serveur edge compact à gérer des charges de travail de calcul exigeantes et de longue durée.
| y-cruncher Temps de calcul total (plus c'est bas, mieux c'est) | SuperServer SYS-E403-14B-FRN2T (Intel Xeon 6521P 24C) | 
| 1 milliard | 10.199 secondes | 
| 2.5 milliard | 27.59 secondes | 
| 5 milliard | 63.12 secondes | 
| 10 milliard | 134.44 secondes | 
| 25 milliard | 391.84 secondes | 
Mixeur 4.0
Blender 4.0 est une application de modélisation 3D open source. Ce benchmark a été réalisé à l'aide de l'utilitaire Blender Benchmark CLI. Le score est mesuré en échantillons par minute, les valeurs les plus élevées indiquant de meilleures performances.
Lors des tests de rendu du processeur Blender 4.0, le SYS-E403-14B-FRN2T a atteint 375.5 échantillons par minute dans la scène Monstre, 246.5 échantillons par minute dans la scène Dépôt et 179.3 échantillons par minute dans la scène Salle de classe. Ces chiffres confirment que le processeur Xeon 24 cœurs du système est parfaitement adapté aux tâches de rendu, même sans accélération GPU.
| Échantillons CPU par minute de Blender 4.0 (plus c'est élevé, mieux c'est) | SuperServer SYS-E403-14B-FRN2T (Intel Xeon 6521P 24C) | 
| Monster | 375.53 | 
| Brocanteur | 246.52 | 
| Salle de classe | 179.32 | 
Équipé d'un GPU NVIDIA L4, les performances ont considérablement augmenté. Avec l'accélération GPU activée, le SYS-E403-14B-FRN2T a atteint 1 975,2 échantillons par minute dans Monster, 1 027,2 échantillons par minute dans Junkshop et 1 036,1 échantillons par minute dans Classroom. Cela démontre la capacité du système à passer des charges de travail basées sur le processeur au rendu accéléré par le GPU, ce qui en fait un choix polyvalent pour les applications gourmandes en ressources de calcul.
| Échantillons GPU Blender 4.0 par minute (plus c'est élevé, mieux c'est) | SuperServer SYS-E403-14B-FRN2T (NVIDIA L4) | 
| Monster | 1,975.18 | 
| Brocanteur | 1,027.16 | 
| Salle de classe | 1,036.09 | 
Points de repère Phoronix
Phoronix Test Suite est une plateforme d'analyse comparative automatisée open source prenant en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org. Elle gère toutes les étapes, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue. Nous examinerons les performances des tests Stream, 7-Zip, de la compilation du noyau Linux, d'Apache et d'OpenSSL.
Grâce à la suite de tests Phoronix, le SYS-E403-14B-FRN2T a fourni d'excellents résultats sur plusieurs charges de travail.
- Bande passante de la mémoire de flux : 305 960 Mo/s
- Compression 7-Zip : 235 421 MIPS
- Compilation du noyau (allmod) : 540 secondes
- Serveur Web Apache : 289 885 requêtes par seconde
- Vérification OpenSSL : 408 milliards de vérifications par seconde
Ces résultats montrent que le système compact est plus que capable de gérer des applications gourmandes en mémoire, des charges de travail de développement, des performances de serveur Web et des opérations cryptographiques, tout en conservant les avantages de son petit format.
| Points de repère Phoronix | SuperServer SYS-E403-14B-FRN2T (Intel Xeon 6521P 24C) | 
| Bande passante de la mémoire de flux | 305,960.3 Mo / s | 
| Compression à 7 zips | 235,421 MIPS | 
| Compilation du noyau (allmod) | 540.260 secondes | 
| Apache (requêtes par seconde) | 289,885.15 XNUMX R/s | 
| Vérification OpenSSL | 408,423,815,760 XNUMX XNUMX XNUMX Vérifications | 
Conclusion
Le Supermicro SYS-E403-14B-FRN2T est un système compact de 16 cm de profondeur, conçu spécifiquement pour les déploiements edge embarqués où l'espace et la facilité d'entretien sont aussi importants que la puissance brute. Malgré sa taille, il conserve les fonctionnalités généralement associées aux serveurs rack plus grands, notamment les blocs d'alimentation à double redondance, la prise en charge des GPU pleine hauteur et des options de processeur flexibles jusqu'à 300 W de TDP. Lors de nos tests, la plateforme a géré avec fiabilité les tâches de calcul exigeantes, obtenant d'excellents résultats sur les charges de travail y-cruncher, Blender et Phoronix. Elle a même pris le dessus sur les serveurs web Apache, où son efficacité et sa réactivité lui ont conféré un avantage sur des systèmes beaucoup plus imposants. L'ajout d'une carte graphique NVIDIA L4 a démontré la rapidité d'adaptation du boîtier aux cas d'utilisation d'inférence ou de rendu IA, renforçant ainsi son rôle de plateforme edge polyvalente.
Physiquement, la conception privilégie un fonctionnement silencieux, une facilité d'entretien en façade et une flexibilité de montage, avec prise en charge des installations murales ou en racks peu profonds. Conçu pour des températures de fonctionnement allant jusqu'à 45 °C, il est parfaitement adapté aux déploiements dans des environnements moins contrôlés où les équipements traditionnels montés en rack peuvent ne pas convenir. Des compromis sont à noter, comme la simplicité de la configuration U.2 par rapport aux options E3.S plus denses, mais cela reflète l'équilibre entre évolutivité et légèreté du système.
Globalement, le SYS-E403-14B-FRN2T offre une combinaison unique de compacité, de flexibilité et de performances pratiques. Il n'est pas destiné à remplacer les serveurs rack haute densité. Néanmoins, pour les entreprises qui ont besoin d'une puissance de calcul et d'une accélération GPU fiables en périphérie, avec un service simple et un encombrement réduit, ce système offre une solution robuste et bien pensée.
Page produit du SuperServer IoT Supermicro SYS-E403-14B-FRN2T
