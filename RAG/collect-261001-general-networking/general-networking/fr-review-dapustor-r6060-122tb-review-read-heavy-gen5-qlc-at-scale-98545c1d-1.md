---
id: collect-261001-general-networking/general-networking/fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d-1
title: "fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "benchmarks", "intel", "nand"]
source: docs/RAG/collect-261001-general-networking/fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d.md
source_anchor: ""
source_lines: [1, 55]
sha256: 7f32f97870280f2e3402c1015a8d221838f3cda7e4539a5816a07f4e12bbc909
---

# fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d

Le DapuStor R6060 est un SSD QLC PCIe Gen5 pour entreprises, basé sur le contrôleur DP800 et la mémoire NAND 3D QLC. Cette gamme se décline aux formats U.2, E3.L et E1.L, avec des capacités respectives de 15.36 To, 30.72 To, 61.44 To et 122.88 To, ainsi qu'une version haut de gamme de 245 To. L'interface prend en charge le fonctionnement double port PCIe 5.0 x4 ou 2×2 et NVMe 2.0. Notre modèle de test est la variante E3.L 2T de 122.88 To, que DapuStor destine aux infrastructures d'IA denses, aux déploiements cloud et aux pools de stockage à grande échelle où la densité de mémoire flash est primordiale. Le profil de performance correspond à ce rôle, associant une bande passante de lecture séquentielle élevée à un débit d'écriture beaucoup plus faible, ce qui rend le R6060 mieux adapté aux ensembles de données à forte intensité de lecture et aux niveaux de stockage axés sur la capacité qu'aux charges de travail transactionnelles par petits blocs.
DapuStor annonce un débit de lecture séquentielle maximal de 14 Go/s pour l'ensemble de la gamme R6060, et un débit d'écriture séquentielle de 4 Go/s. Les performances en lecture aléatoire atteignent 3 millions d'IOPS pour les modèles de 15.36 To et 30.72 To, tandis que les modèles de 61.44 To et 122.88 To sont évalués à 2.8 millions d'IOPS. Les performances en écriture aléatoire sont nettement inférieures : 40 000 IOPS à 16 Ko pour les capacités les plus faibles et 20 000 IOPS à 32 Ko pour le modèle de 122.88 To. Le R6060 se positionne ainsi comme un outil principalement dédié à la lecture, notamment pour les grands ensembles de données et les niveaux de stockage haute densité où la capacité utilisable prime sur la vitesse d'écriture par petits blocs.
Le R6060 intègre les fonctionnalités professionnelles attendues d'un disque de ce type, notamment la technologie NVMe 2.0 FDP (Flexible Data Placement), particulièrement pertinente pour la mémoire QLC. FDP offre à l'hôte un meilleur contrôle sur l'emplacement d'écriture des données, ce qui contribue à réduire l'amplification d'écriture et à optimiser l'utilisation de la mémoire flash. DapuStor annonce un taux d'écriture par jour (DWPD) de 0.6 DWPD pour ce disque et indique sa conformité OCP 2.5, NVMe-MI 1.2, la protection des données de bout en bout, le démarrage sécurisé, la vérification du firmware, la prise en charge de l'effacement sécurisé des données, la télémétrie, la surveillance de la latence et la prise en charge du double port pour les systèmes nécessitant une redondance des chemins d'accès.
DapuStor offre une garantie de cinq ans sur le R6060 et indique une consommation maximale de 25 W et une consommation en veille de 5 W. La latence aléatoire est de 80/25 µs en lecture/écriture, la latence séquentielle est de 7/8 µs et le MTBF est de 2.5 millions d'heures.
Spécifications du DapuStor R6060 122 To
| Métrique/Champ | 15.36TB | 30.72TB | 61.44TB | 122.88TB | 
|---|---|---|---|---|
| Généralités |  |  |  |  | 
| PCN | R6060 |  |  |  | 
| Capacité (To) | 15.36 | 30.72 | 61.44 | 122.88 | 
| Facteur de forme | U.2/E3.L 2T/E1.L |  |  |  | 
| Interface | PCIe 5.0×4 / 2×2, NVMe 2.0 |  |  |  | 
| Type de flash | Mémoire flash NAND QLC 3D Enterprise |  |  |  | 
| Performances |  |  |  |  | 
| Bande passante de lecture à 128 Ko (Mo/s) | 14000 | 14000 | 14000 | 14000 | 
| Bande passante d'écriture à 128 Ko (Mo/s) | 4000 | 4000 | 4000 | 4000 | 
| Lecture aléatoire @4KB KIOPS | 3000 | 3000 | 2800 | 2800 | 
| KIOPS d'écriture aléatoire | 40@16KB | 40@16KB | 40@16KB | 20@32KB | 
| Latence aléatoire R/W (µs) | 80/25 |  |  |  | 
| Latence séquentielle R/W (µs) | 7/8 |  |  |  | 
| Tuning Moteur |  |  |  |  | 
| Puissance maximale (W) | 25 |  |  |  | 
| Puissance au ralenti (W) | 5 |  |  |  | 
| Fiabilité |  |  |  |  | 
| Endurance | 0.6 DWPD |  |  |  | 
| MTBF | 2.5 millions d'heures |  |  |  | 
| UBER | 1 secteur par 10^18 bits lus |  |  |  | 
| Garantie | 5 ans |  |  |  | 
Performances du DapuStor R6060
Plateforme de test de conduite
Nous utilisons un serveur Dell PowerEdge R760 exécutant Ubuntu 22.04.2 LTS comme plateforme de test pour toutes les charges de travail présentées dans ce rapport. Équipé d'un boîtier JBOF Serial Cables Gen5 , il offre une large compatibilité avec les SSD U.2, E1.S, E3.S et M.2. La configuration de notre système est détaillée ci-dessous :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
Comparaison des lecteurs
- Solidigm P5336 122.88 To (Gen4 | 2.5″ | U.2)
- Solidigm P5336 61.44 To (Gen4 | 2.5″ | U.2)
- Micron 6550 ION 61.44 To (Gen5 | E3.S)
- DapuStor J5060 61.44 To (Gen4 | 2.4″ | U.2)
- DapuStor R6060 122.88 To (Gen5 | E3.L)
Lors de l'analyse des résultats des tests de performance, il est important de tenir compte du positionnement de ces disques. Bien qu'ils ne soient pas tous en concurrence directe dans les scénarios de déploiement, leurs capacités et leurs cibles marketing se recoupent suffisamment pour fournir un contexte utile quant à la place du DapuStor R6060 122.88 To sur le marché actuel des SSD d'entreprise haute capacité.
Ce groupe de comparaison met en lumière différentes approches pour l'extension des capacités de mémoire flash d'entreprise. La Micron 6550 ION de 61.44 To, basée sur la technologie TLC, privilégie des performances brutes Gen5 supérieures, tandis que les Solidigm P5336 de 122.88 To et 61.44 To misent principalement sur l'optimisation de la densité et du rapport coût-efficacité grâce à la technologie NAND QLC. La DapuStor J5060 de 61.44 To constitue un autre point de référence pour la technologie Gen4 U.2, tandis que la R6060, avec ses 122.88 To, s'aventure sur le segment plus récent de la Gen5 E3.L.
L'inclusion de ces disques offre une vision plus large de la façon dont le R6060 se compare aux conceptions axées sur la performance et sur la densité, alors que les fournisseurs continuent de faire évoluer les plateformes de stockage d'entreprise haute capacité.
Benchmark de performance FIO
Pour mesurer les performances de stockage de chaque SSD selon les indicateurs courants du secteur, nous utilisons FIO. Chaque disque est soumis au même processus de test, qui comprend une étape de préconditionnement avec deux remplissages complets du disque avec une charge de travail d'écriture séquentielle, suivie d'une mesure des performances en régime permanent. À chaque modification du type de charge de travail mesuré, nous effectuons un nouveau remplissage de préconditionnement avec cette nouvelle taille de transfert.
Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Séquentiel 128K
- 64K Aléatoire
- 16K Aléatoire
- 4K Aléatoire
Écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Le disque DapuStor R6060 de 122.88 To a atteint une vitesse de 3 920,6 Mo/s lors du test d'écriture séquentielle à 128 Ko, se classant ainsi deuxième de ce groupe. Le Micron 6550 ION de 61.44 To le devance largement avec 10 456,4 Mo/s, mais le R6060 reste nettement en tête devant le Solidigm P5336 de 122.88 To (3 152,5 Mo/s), le DapuStor J5060 de 61.44 To (2 883,1 Mo/s) et le Solidigm P5336 de 61.44 To (2 503,5 Mo/s).
Latence d'écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Le DapuStor R6060 de 122.88 To a affiché une latence d'écriture séquentielle de 509.7 µs (128 Ko), se classant une nouvelle fois deuxième. Le Micron 6550 ION de 61.44 To a dominé le classement avec 191.0 µs, tandis que les autres disques testés étaient à la traîne du R6060, notamment le Solidigm P5336 de 122.88 To (634.0 µs), le DapuStor J5060 de 61.44 To (693.3 µs) et le Solidigm P5336 de 61.44 To (798.4 µs). Malgré une avance confortable de Micron, le R6060 a tout de même réalisé la meilleure performance parmi les disques restants.
