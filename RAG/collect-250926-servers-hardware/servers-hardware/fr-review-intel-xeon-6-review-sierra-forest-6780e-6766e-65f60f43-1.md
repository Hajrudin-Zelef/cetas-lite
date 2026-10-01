---
id: collect-250926-servers-hardware/servers-hardware/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43-1
title: "fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "compute", "dram"]
source: docs/RAG/clean4/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43.md
source_anchor: ""
source_lines: [1, 24]
sha256: 89c1fef5407c862c3017da2556e27a39896d6008ded54b8fe62457b33f7c84ee
---

# fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43

La famille de processeurs de serveur Intel Xeon 6 a été lancée il y a quelques semaines, remplaçant l'ancienne marque « Scalable ». Il existe deux gammes de puces : Sierra Forest, qui propose des cœurs E, et Granite Rapids, qui propose des cœurs P. Dans cette revue, nous avons échantillonné quelques premiers ensembles de processeurs Sierra Forest, y compris le 6780E haut de gamme, qui commercialise 144 cœurs.
Les processeurs Xeon 6 sont conçus pour prendre en charge pratiquement toutes les charges de travail, depuis les cas d'utilisation évidents de l'IA jusqu'aux déploiements moins exigeants en périphérie. Pour aider les clients à comprendre les deux couloirs de nage, Intel indique que Sierra Forest est « optimisé pour les performances par watt dans les charges de travail de calcul haute densité et évolutives ». Le processeur P-core Granite Rapids est « optimisé pour les performances par cœur dans les charges de travail gourmandes en calcul ». Cependant, tous les processeurs Xeon 6 utilisent une plate-forme et une pile de micrologiciels communes.
Pour donner un peu plus de contexte sur les performances des processeurs, Intel positionne Sierra Forest comme un excellent remplacement pour les systèmes ayant un cycle de rafraîchissement de 5 ans. Dans ce cas, les chiffres de performance par watt devraient être nettement plus favorables.
Le déploiement du Xeon 6 se déroulera par étapes, avec le lancement de la famille 6700E aujourd'hui. Mais il y a bien plus à venir. Les processeurs Intel Xeon 6900P devraient sortir cette année au troisième trimestre. Nous nous attendons à voir un déluge de puces, notamment les 3E, 6900P, 6700P, Xeon 6500 SoC et 6P, au cours du premier trimestre 6300.
Innovations architecturales
Le Xeon 6 présente deux conceptions de microarchitecture, une plate-forme de cœurs optimisés en termes de performances (cœurs P) et une plate-forme de cœurs optimisés en efficacité (cœurs E). Cette approche hybride permet aux centres de données de concevoir un rack dans lequel les charges de travail à forte intensité de calcul, telles que l'IA et le HPC, bénéficient de performances maximales tandis que les tâches orientées débit, comme les microservices et la mise en réseau, atteignent un nouveau niveau d'efficacité énergétique.
Augmentation du nombre de cœurs
La principale caractéristique des processeurs Xeon 6 Sierra Forest est l'augmentation significative du nombre de cœurs. En intégrant davantage de cœurs par processeur, Intel permet aux centres de données de gérer simultanément un plus large éventail de charges de travail. Cette augmentation de la densité du cœur est particulièrement bénéfique pour les applications qui nécessitent des niveaux élevés de traitement parallèle, garantissant que les ressources sont utilisées de manière optimale pour maximiser le débit et minimiser la latence.
Bande passante mémoire améliorée
Depuis quelques générations, on a l’impression que nous sommes coincés avec des vitesses de DRAM de serveur plus lentes alors que le marché grand public a basculé jusqu’à bien au-dessus de 7000 5 MT/s. L'intégration de la mémoire DDR2.0 avec Ultra Path Interconnect (UPI) 6 dans les processeurs Xeon XNUMX Sierra Forest améliore considérablement la bande passante mémoire. Cela permet un accès plus rapide aux données et réduit les goulots d'étranglement, garantissant ainsi le bon fonctionnement des applications hautes performances. L'amélioration de la bande passante mémoire est essentielle pour les applications qui gèrent de grands ensembles de données et nécessitent une récupération et un traitement rapides des données, comme la formation à l'IA et l'analyse du Big Data.
Capacités d'E/S avancées d'Intel Xeon 6
La prise en charge de PCIe 5.0 et Compute Express Link (CXL) 2.0 offre aux processeurs Xeon 6 Sierra Forest des capacités d'E/S avancées. PCIe 5.0 offre le double de la bande passante de son prédécesseur, permettant une communication plus rapide entre le processeur et les périphériques. CXL 2.0 améliore encore la connectivité en fournissant une interconnexion à large bande passante et à faible latence pour les processeurs, les accélérateurs, la mémoire et le stockage. Cela garantit que les centres de données peuvent répondre aux exigences des applications modernes à haut débit et s'intégrer de manière transparente aux technologies futures.
Architecture modulaire multi-matrices
Les processeurs Xeon 6 Sierra Forest utilisent une architecture multi-puces modulaire activée par la technologie Embedded Multi-die Interconnect Bridge (EMIB). Cette conception permet de combiner plusieurs puces dans un seul boîtier, offrant une bande passante élevée et une faible latence tout en maintenant une consommation d'énergie efficace. L'approche modulaire offre flexibilité et évolutivité, permettant aux centres de données de personnaliser leur infrastructure pour répondre aux exigences spécifiques de la charge de travail. Cette architecture prend également en charge une meilleure gestion thermique et une meilleure efficacité énergétique, qui sont essentielles au maintien des performances et de la fiabilité dans les environnements haute densité.
Architecture de matrice Intel Xeon 6
L'architecture des puces des processeurs Xeon 6 Sierra Forest constitue un aperçu intéressant de l'approche innovante d'Intel visant à maximiser les performances et l'efficacité. Contrairement à un grand interposeur en silicium traditionnel, EMIB implémente une petite puce de pont avec plusieurs couches de routage. En tirant parti de la technologie EMIB, Intel peut interconnecter plusieurs puces dans un seul boîtier, réduisant ainsi la taille globale du boîtier et améliorant l'intégrité du signal. Cette configuration permet des taux de transfert de données plus rapides entre les matrices, ce qui est essentiel pour maintenir des performances élevées dans une large gamme d'applications.
Importance pour les centres de données
Performances et efficacité
Par rapport aux générations précédentes, les processeurs Xeon 6 offrent des performances par watt jusqu'à 2.7 fois supérieures, ce qui les rend idéaux pour l'inférence d'IA, le transcodage multimédia et les tâches de calcul générales. Cet équilibre garantit que les centres de données peuvent gérer davantage de charges de travail avec une consommation d'énergie inférieure, ce qui se traduit directement par une réduction des coûts opérationnels et une durabilité améliorée.
Scalabilité et flexibilité
L'architecture modulaire permet aux centres de données de personnaliser leur infrastructure en fonction des demandes de charge de travail spécifiques. La possibilité de déployer une combinaison de cœurs P et de cœurs E sur plusieurs plates-formes offre une approche personnalisée pour gérer divers besoins informatiques, améliorant à la fois les performances et l'efficacité. Cela présente un concept intéressant dans lequel un châssis peut être sélectionné puis déployé, certains disposant de processeurs E-core disponibles pour exécuter efficacement les services de base et gérer les opérations à des moments moins exigeants, tandis que d'autres sont équipés de processeurs P-core pour gérer la demande de pointe et charges de travail exigeantes.
Sécurité et fiabilité améliorées
Construits avec des fonctionnalités de sécurité améliorées au niveau matériel, les processeurs Xeon 6 offrent une protection robuste pour l'intégrité des données et la fiabilité du système. Ceci est essentiel pour maintenir la confiance et la conformité dans les environnements sensibles aux données.
Pérennité grâce aux technologies avancées
