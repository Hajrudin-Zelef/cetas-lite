---
id: collect-261001-ia-llm/ia-llm/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662-5
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: ia-llm
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "gpu", "valuation"]
source: docs/RAG/collect-261001-ia-llm/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [147, 176]
sha256: f26b59af142e28829e6707bdcc3a9c06dac00c67cb84bb6bb01dac10c1cc656d
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

MariaDB 11.4.4 a affiché les meilleures performances transactionnelles. Elle a surpassé les anciennes versions de MariaDB, telles que 10.6.22 et la version 10.11.12 optimisée pour les besoins spécifiques (MDEV-21923). MySQL 8.4.4 a également affiché d'excellentes performances, talonnant de près MariaDB 11.4.4. PostgreSQL 17.2 a obtenu des résultats compétitifs, mais est resté légèrement en retrait par rapport à MariaDB et à la nouvelle version de MySQL. MySQL 5.7.44 était la base de données la plus faible parmi les bases de données testées.
Référence de compression à 7 zips
Le test de mémoire intégré à l'utilitaire 7-Zip mesure les performances du processeur et de la mémoire d'un système pendant les tâches de compression et de décompression, indiquant dans quelle mesure le système peut gérer des opérations gourmandes en données. Nous exécutons ce test avec une taille de dictionnaire de 128 Mo lorsque cela est possible.
Bien qu'il possède moins de cœurs que certains autres processeurs de la série EPYC 9005, il a atteint un score total de 378.469 GIPS. Bien que ce score soit sans aucun doute honorable, il est nettement inférieur à celui de l'EPYC 9755 (443.029 GIPS) et de l'EPYC 9575F (394.900 GIPS). Il est intéressant de noter que le 9965 (266.740 GIPS) est à la traîne dans ce benchmark, ce qui suggère qu'un nombre de cœurs plus élevé ne se traduit pas toujours par de meilleures performances de compression.
Lors des tâches de décompression, le 9665P maintient son niveau actuel avec 395.502 GIPS, mais reste à la traîne derrière le 9755 (487.263 GIPS) et le 9575F (425.580 GIPS). Son architecture monosocket et ses fréquences d'horloge élevées lui permettent de rester compétitif, mais les modèles à plus haut débit ont l'avantage en termes de débit brut.
| Référence de compression à 7 zips (Plus c'est haut, mieux c'est) | AMD EPYC 9655P (96c) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c, SMT désactivé) |  | 
| Compression |  |  |  |  |  | 
| Utilisation actuelle du processeur | 5881 % | 4302 % | 5233 % | 4406 % |  | 
| Note actuelle/utilisation | 6.112 GIPS | 5.830 GIPS | 7.597 GIPS | 7.975 GIPS |  | 
| Courant | 359.471 GIPS | 250.827 GIPS | 397.536 GIPS | 351.358 GIPS |  | 
| Utilisation résultante du processeur | 5875 % | 4041 % | 5306 % | 4555 % |  | 
| Évaluation/utilisation résultante | 6.140 GIPS | 5.804 GIPS | 7.720 GIPS | 8.070 GIPS |  | 
| Note résultante | 360.695 GIPS | 234.317 GIPS | 409.652 GIPS | 367.358 GIPS |  | 
| Décompression |  |  |  |  |  | 
| Utilisation actuelle du processeur | 6168 % | 4322 % | 6041 % | 5017 % |  | 
| Note actuelle/utilisation | 6.412 GIPS | 7.078 GIPS | 8.065 GIPS | 8.483 GIPS |  | 
| Courant | 395.502 GIPS | 305.909 GIPS | 487.263 GIPS | 425.580 GIPS |  | 
| Utilisation résultante du processeur | 6159 % | 4556 % | 5921 % | 4940 % |  | 
| Évaluation/utilisation résultante | 6.434 GIPS | 6.577 GIPS | 8.045 GIPS | 8.569 GIPS |  | 
| Note résultante | 396.243 GIPS | 299.163 GIPS | 476.405 GIPS | 422.441 GIPS |  | 
| Note totale |  |  |  |  |  | 
| Utilisation totale du processeur | 6017 % | 4298 % | 5613 % | 4747 % |  | 
| Note totale/utilisation | 6.287 GIPS | 6.190 GIPS | 7.883 GIPS | 8.319 GIPS |  | 
| Note totale | 378.469 GIPS | 266.740 GIPS | 443.029 GIPS | 394.900 GIPS |  | 
Conclusion
Le Dell PowerEdge R7715 offre un équilibre impressionnant entre performances, évolutivité et efficacité pour les charges de travail des entreprises modernes. Compatible avec les processeurs AMD EPYC série 9005, offrant jusqu'à 160 cœurs et 24 emplacements DIMM DDR5 pour jusqu'à 6 To de mémoire, le R7715 est parfaitement équipé pour gérer les applications gourmandes en données dans les environnements de virtualisation, d'analyse et de stockage défini par logiciel.
Grâce à des fréquences d'horloge élevées et à une architecture performante, le R7715 excelle en monocœur et offre des performances multicœurs compétitives grâce à sa conception monosocket. Bien qu'il ne soit pas à la hauteur des systèmes bi-socket en calcul parallèle brut, il s'en rapproche étonnamment, offrant une alternative plus économe en énergie et plus économique pour de nombreuses charges de travail réelles.
Cette configuration n'est pas idéale pour tous les cas d'utilisation, notamment lorsque la densité maximale des cœurs ou l'accélération GPU sont essentielles. Cependant, Dell devrait déployer la prise en charge GPU pour le R7715 plus tard cette année. De plus, les entreprises ayant besoin de disques de plus grande capacité pour des charges de travail spécifiques pourraient envisager des configurations de stockage plus étendues.
En définitive, le R7715 est une plateforme idéale pour les environnements informatiques qui privilégient un débit élevé, une mémoire rapide et la flexibilité des E/S Gen5, sans la complexité ni le coût des déploiements à double socket. Le R7715 s'impose comme une option judicieuse pour les entreprises cherchant à optimiser leur efficacité sans sacrifier les capacités.
Page de configuration de produit
