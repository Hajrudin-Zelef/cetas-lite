---
id: collect-261001-huawei/huawei/fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217-1
title: "fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217"
domain: huawei
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["diffusion", "distribution", "gpu", "intel", "nand", "valuation"]
source: docs/RAG/collect-261001-huawei/fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217.md
source_anchor: ""
source_lines: [1, 60]
sha256: b7f68d5a80b447f67b7a5a6e7b6703855b61eecea15ac0a0279887b676eaca5a
---

# fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217

La gamme Pascari X-Series de Phison est conçue pour répondre aux différents besoins de stockage des entreprises, offrant des solutions personnalisées pour les charges de travail intensives en lecture et en écriture. Le X200P est le modèle haute capacité, prenant en charge jusqu'à 30.72 To avec un DWPD (nombre d'écritures par jour). Il utilise le PCIe Gen5 et la NAND TLC et est disponible aux formats U.3, U.2 et E3.S. Le X200P est conçu pour un large éventail d'usages, notamment la diffusion de contenu à grande échelle, l'inférence IA et l'archivage de données froides.
Phison propose également les modèles haute endurance X200E, optimisés pour les charges de travail à forte intensité d'écriture, prenant en charge jusqu'à trois DWPD et des capacités allant de 1.6 To à 25.6 To, ce qui les rend parfaitement adaptés aux bases de données transactionnelles, aux analyses en temps réel et au traitement des journaux.
Pour ce test, Phison nous a confié le modèle X7.68P U.200 de 2 To. Nous l'avons soumis à notre suite complète de tests d'entreprise afin d'évaluer ses performances sous pression.
Phison Pascari X200P Série Spécifications techniques
| Spécifications Phison Pascari Série X200P | 1.92TB | 3.84TB | 7.68TB | 15.36TB | 30.72TB | 
|---|---|---|---|---|---|
| Facteur de forme | U.2 |  |  |  |  | 
| Interface | PCIe 5.0 x4, 2×2 |  |  |  |  | 
| NVMe | 2.0 |  |  |  |  | 
| NAND flash | 3D TLC |  |  |  |  | 
| Lecture séquentielle (Mo / s) | 14,800 | 14,800 | 14,800 | 14,800 | 14,000 (Est.) | 
| Écriture séquentielle (Mo / s) | 4,300 | 8,600 | 8,700 | 8,350 | 7,500 (Est.) | 
| Lecture aléatoire 4K (IOPS) | 2,400K | 3,000K | 3,000K | 3,000K | 2,300 XNUMX XNUMX (est.) | 
| Écriture aléatoire 4K (IOPS) | 170K | 380K | 500K | 500K | 283 XNUMX XNUMX (est.) | 
| Latence de lecture (μs) | 60 |  |  |  |  | 
| Latence d'écriture (μs) | 10 |  |  |  |  | 
| Puissance – Active (W) |  |  |  |  |  | 
| Puissance – Ralenti (W) | 5 |  |  |  |  | 
| DWPD(7) | 1 |  |  |  |  | 
| UBER | <1 secteur pour 1018 bits lus |  |  |  |  | 
| MTBF (millions d'heures) | 2.5 |  |  |  |  | 
| Garantie limitée (années) | 5 |  |  |  |  | 
| Exploitation temporaire. (° C) | 0 à 70 ans, qui |  |  |  |  | 
| Température hors fonctionnement (°C) | -40 à 85 |  |  |  |  | 
| Dimensions (mm) | 100.10 (L) x 69.85 (L) x 15.00 (H) |  |  |  |  | 
| Poids (g) | 188 | 199 | 201 | 168 |  | 
Conception et construction du Pascari X200P 7.68 To
Notre unité de test est la version U.7.68 2 pouces de 2.5 To du X200P, conçue pour les applications de stockage d'entreprise hautes performances. Elle utilise une interface PCIe 5.0, entièrement conforme à la spécification NVMe 2.0. Le disque est construit autour d'une NAND 3D TLC haute endurance et prend en charge des capacités allant jusqu'à 30.72 To.
Physiquement, le X200P présente un format U.2.5 standard de 2 pouces, mesurant 100.10 mm de longueur, 69.85 mm de largeur et 15.00 mm de hauteur, pour un poids de 201 grammes. L'ensemble est logé dans un élégant boîtier en aluminium noir avec refroidissement passif intégré, conçu pour gérer efficacement les températures sous des charges de travail élevées et soutenues. Le disque prend également en charge les configurations E3.S pour plus de flexibilité dans les environnements de stockage denses.
Côté performances, il offre des débits allant jusqu'à 14,800 8,700 Mo/s en lecture séquentielle, 3 500 Mo/s en écriture séquentielle, ainsi que jusqu'à 25 millions d'IOPS en lecture aléatoire et 5 XNUMX IOPS en écriture aléatoire. Sa consommation en mode actif est inférieure à XNUMX W, avec seulement XNUMX W au repos, ce qui en fait une solution performante pour les opérations à haut débit soutenues.
L'indice d'endurance est de 1 DWPD, avec un MTBF de 2.5 millions d'heures et une garantie limitée de 5 ans. Il est conçu pour les charges de travail d'entreprise fonctionnant 24h/7 et 0j/70, avec une plage de températures de fonctionnement de XNUMX °C à XNUMX °C.
Phison comprend un ensemble complet de fonctionnalités de protection et de gestion des données de niveau entreprise :
- Protection contre les pertes de puissance (PLP)
- Prise en charge de l'ISE (effacement sécurisé instantané) et de TCG Opal 2.0
- Cryptage AES-XTS 256 bits
- Protection du chemin de données de bout en bout
- Protection des métadonnées
- SECDED (Correction d'erreur simple et détection d'erreur double)
- Opérations de désinfection
- NVMe-MI (interface de gestion)
- Compatibilité SMBus
- Prise en charge jusqu'à 128 espaces de noms
Dans l'ensemble, la gamme Pascari X200P combine une qualité de construction robuste de qualité industrielle avec des performances de pointe et une fiabilité de niveau entreprise, ce qui en fait un candidat solide pour les environnements de stockage exigeants tels que le cloud, l'IA/ML et les infrastructures virtualisées.
Test de performance
Plateforme de test de conduite
Nous utilisons un serveur Dell PowerEdge R760 exécutant Ubuntu 22.04.02 LTS comme plateforme de test pour toutes les charges de travail présentées dans cet article. Équipé d'un boîtier JBOF Serial Cables Gen5 , il offre une large compatibilité avec les SSD U.2, E1.S, E3.S et M.2. La configuration de notre système est décrite ci-dessous :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
Comparaison des lecteurs
- Pascari X200P 7.68 To
- Micron 9550 7.68 To
- SanDisk SN861 7.68 To
- Solidigm PS1010 7.68 To
- Kingston DC3000ME 7.68 To
Nous avons comparé le Pascari X200P à un groupe de SSD NVMe PCIe Gen7.68 de 5 To de taille similaire utilisant la mémoire Flash NAND TLC. La comparaison comprenait les Micron 9550, SanDisk SN861, Solidigm PS1010 et Kingston DC3000ME. Ces disques représentent des solutions professionnelles de moyenne capacité conçues pour les environnements hautes performances. Les tests ont été réalisés à l'aide de bancs d'essai réels et synthétiques, notamment CDN, FIO et GDSIO, afin de mesurer les performances en termes de débit soutenu, de latence, de schémas d'E/S mixtes et de charges de travail accélérées par GPU. En standardisant la capacité, l'interface et le type de mémoire NAND, cette évaluation permet de comparer clairement les performances du Pascari X200P à celles de ses concurrents dans des conditions exigeantes.
Performances du CDN
Afin de simuler une charge de travail CDN réaliste à contenu mixte, les SSD ont été soumis à une séquence d'analyse comparative en plusieurs phases conçue pour reproduire les schémas d'E/S des serveurs Edge à contenu important. La procédure de test couvre différentes tailles de blocs, grandes et petites, réparties sur des opérations aléatoires et séquentielles, avec différents niveaux de concurrence.
Avant les principaux tests de performance, chaque SSD a effectué un remplissage complet du périphérique avec une passe d'écriture séquentielle à 100 % utilisant des blocs de 1 Mo. Ce processus utilisait des E/S synchrones et une profondeur de file d'attente de quatre, permettant quatre tâches simultanées. Cette phase garantit que le disque entre dans un état stable représentatif d'une utilisation réelle. Après le remplissage séquentiel, une deuxième phase de saturation d'écriture aléatoire de trois heures a été exécutée selon une distribution pondérée de la taille des blocs (taille de bloc/pourcentage), privilégiant fortement les transferts de 128 Ko (98.51 %), avec des contributions mineures des blocs inférieurs à 128 Ko jusqu'à 8 Ko. Cette étape simule les schémas d'écriture fragmentés et irréguliers souvent observés dans les environnements de cache distribué.
