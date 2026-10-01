---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-4
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmarks", "open source"]
source: docs/RAG/clean4/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [140, 167]
sha256: aec664e18ea518952411c3bbc8c8f5ecfb24b1648b15004dd03a443afe8087ca
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

| 2.5 milliard | 15.245 s | 13.811 secondes | 
| 5 milliard | 24.961 s | 22.107 secondes | 
| 10 milliard | 44.350 | 40.111 secondes | 
| 25 milliard | 114.107 | 98.445 secondes | 
| 50 milliard | 246.688 | 211.567 secondes | 
| 100 milliard | 572.800 | 481.207 secondes | 
Points de repère Phoronix
Phoronix Test Suite est une plateforme d'analyse comparative automatisée et open source qui prend en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org. Elle gère l'ensemble du processus, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue.
Bande passante de la mémoire de flux
Lors du test de bande passante mémoire Stream, le Supermicro AS-2126HS-TN atteint un débit soutenu de 807 766 Mo/s. Le Dell PowerEdge R7725, quant à lui, culmine à 883 312 Mo/s, ce qui indique une bande passante mémoire maximale légèrement supérieure. Toutefois, les deux systèmes se situent dans la plage de performances attendue pour les plateformes biprocesseurs EPYC 9005.
Compression à 7 zips
Pour la compression 7-Zip, le système Supermicro atteint 1 262 832 MIPS, démontrant ainsi d'excellentes performances en calcul entier et une mise à l'échelle multithread efficace. Le système Dell affiche 1 326 967 MIPS, ce qui représente un léger avantage pour cette charge de travail, tout en restant dans la même catégorie de performances globales.
Compilation du noyau
Lors de la compilation du noyau (allmod), le Supermicro AS-2126HS-TN exécute la tâche en 117.97 secondes, devançant ainsi le Dell PowerEdge R7725 qui réalise la même opération en 139.36 secondes. Ce résultat souligne l'efficacité de la plateforme Supermicro dans les scénarios de compilation parallèle courants dans les environnements de développement et d'intégration continue.
Serveur Web Apache
Le test de performance Apache montre que le système Supermicro traite 90 623,69 requêtes par seconde, tandis que le système Dell atteint 96 782,75 requêtes par seconde. Les deux systèmes affichent un débit élevé pour les serveurs web, la configuration Dell présentant toutefois un pic de requêtes légèrement supérieur.
Vérification OpenSSL
Lors des tests OpenSSL, le Supermicro AS-2126HS-TN atteint 3.55 To/s, ce qui témoigne d'un débit cryptographique important, adapté aux charges de travail exigeantes en matière de chiffrement. Le Dell PowerEdge R7725 va encore plus loin avec 4.41 To/s, indiquant des performances cryptographiques de pointe supérieures. Les deux plateformes offrent des performances largement supérieures aux besoins habituels des entreprises.
| Points de repère Phoronix | Serveur Supermicro A+ AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (double processeur AMD EPYC 9965 192C) | 
|---|---|---|
| Discussions | 807,766.0 Mo / s | 883,312.0 Mo / s | 
| 7-ZIP | 1,262,832 MIPS | 1,326,967 XNUMX MIP/s | 
| Compilation du noyau (allmod) | 117.97 secondes | 139.356 secondes | 
| Apache (requêtes par seconde) | 90,623.69 XNUMX R/s | 96,782.75 XNUMX R/s | 
| OpenSSL | 3,545,769,484,910 XNUMX XNUMX XNUMX Vérifications | 4,409,642,672,307 XNUMX XNUMX XNUMX XNUMX Vérifications | 
Conclusion
Le serveur Supermicro A+ AS-2126HS-TN offre une plateforme Turin biprocesseur équilibrée, privilégiant la densité de calcul, l'évolutivité énergétique et l'extension PCIe flexible dans un format 2U. Compatible avec deux processeurs AMD EPYC série 9005, les configurations à TDP élevé et les configurations d'E/S flexibles, il est idéal pour les environnements d'entreprise, de cloud et de calcul haute performance (HPC) où la performance parallèle soutenue et la configurabilité sont essentielles.
Dans notre suite de tests d'entreprise, le AS-2126HS-TN a offert des performances constantes et prévisibles pour les charges de travail gourmandes en ressources CPU. Bien qu'il ait légèrement esquivé certains indicateurs de débit absolu, la plateforme Supermicro est restée compétitive, démontrant une excellente évolutivité et une grande efficacité pour le rendu, le calcul et les charges de travail mixtes. Ces résultats confirment que le AS-2126HS-TN constitue une base performante et flexible pour les déploiements à grand nombre de cœurs, notamment lorsque la polyvalence de la plateforme et la marge de puissance sont aussi importantes que les performances maximales des benchmarks.
