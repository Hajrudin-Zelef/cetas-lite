---
id: collect-261001-general-networking/general-networking/fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a-1
title: "fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "intel", "nand", "nvidia"]
source: docs/RAG/collect-261001-general-networking/fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a.md
source_anchor: ""
source_lines: [1, 43]
sha256: 07cf58f9deee4d7b3d3eb606eba6f11a4426dc4dc25d6d57e3ba3d1bde126778
---

# fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a

Le Solidigm D7-PS1030 est le modèle d'endurance intermédiaire de la première famille de serveurs PCIe 5.0 de Solidigm. Il partage la même plateforme que le D7-PS1010 et offre une endurance de 3 DWPD ainsi qu'une capacité d'écriture aléatoire maximale de 800 000 IOPS, soit le double des 400 000 IOPS de son homologue d'endurance standard. La famille propose des capacités allant de 1.6 To à 12.8 To en interfaces E3.S et U.2, utilisant de la mémoire NAND TLC 176 couches et des performances nominales allant jusqu'à 14 500 Mo/s en lecture séquentielle et 10 000 Mo/s en écriture séquentielle. Solidigm cible les charges de travail mixtes et fortement axées sur l'écriture : OLTP, journalisation des métadonnées, HPC et pipelines d'IA/ML où la pression en écriture est constante. Notre modèle de test est la version E3.S de 12.8 To.
Ce n'est pas la première fois que nous testons ce disque ; cependant, c'est la première fois qu'un seul disque est évalué. Huit de ces mêmes unités de 12.8 To ont géré la couche flash lors de nos tests de déchargement du cache KV sur le Dell PowerEdge XE7740. La baie a supporté 1.9 Go/s d'écritures KV continues 24 h/24, soit un cycle d'utilisation équivalent à environ 3.2 écritures par jour et par disque en mode miroir, ou environ 1.6 DWPD en RAID 0. C'est précisément la plage de performances pour laquelle le PS1030 est conçu avec sa capacité de 3 DWPD : le déchargement KV est une charge de travail où l'endurance, et non la capacité ou la vitesse de pointe, est la contrainte principale. Les applications à forte intensité de lecture qui font souvent la une des journaux concernant la 5e génération épuiseraient rapidement son budget de 1 DWPD avec une telle charge.
Avec une capacité de 12.8 To, Solidigm annonce 2.75 millions d'IOPS en lecture aléatoire 4K et 800 000 IOPS en écriture aléatoire, le pic de 3.1 millions d'IOPS en lecture aléatoire de la gamme se situant en dessous de ce seuil. La consommation électrique est de 23 W en fonctionnement et de 5 W en veille, conformément aux standards de la génération 5, avec cinq niveaux de consommation configurables de 5 W à 25 W pour les opérateurs disposant d'un budget rack fixe. Solidigm met également en avant deux autres atouts : une efficacité énergétique jusqu'à 70 % supérieure à celle des disques comparables et une constance des IOPS jusqu'à 90 % sur toute la durée de vie du disque. L'endurance est également optimisée sur de courtes périodes : le même support supporte 4.98 DWPD sur trois ans, et le modèle 12.8 To est conçu pour une capacité d'écriture de 70 Po dans les deux sens. Les spécifications de fiabilité sont conformes à la catégorie : MTBF de 2.5 millions d’heures, une garantie de cinq ans et un test UBER Solidigm à 1E-18. Les options de sécurité comprennent le TCG Opal 2.02 SED trim, le matériel certifiable FIPS 140-3 niveau 2, le démarrage sécurisé et la signature du firmware conformes à la norme OCP, ainsi que l’attestation du dispositif et la révocation des clés.
Spécifications du Solidigm D7-PS1030
| Spécifications | Solidigm D7-PS1030 (12.8 To E3.S, gamme familiale indiquée) | 
|---|---|
| En savoir plus sur la plateforme |  | 
| Capacités | 1.6TB 3.2TB 6.4TB 12.8 To (testé) | 
| Facteurs de forme | E3.S 7.5 mm (testé) U.2 15mm | 
| Interface / Protocole | PCIe 5.0x4, NVMe | 
| NON | Mémoire NAND 3D TLC Solidigm à 176 couches | 
| Performance (jusqu'à, selon le fournisseur) |  | 
| Lecture séquentielle (128 Ko) | 10 000 Mo/s (famille) | 
| Écriture séquentielle (128 Ko) | 10 000 Mo/s (famille) | 
| Lecture aléatoire (4K) | 2,750 12.8 IOPS (XNUMX To) Jusqu'à 3 100 000 IOPS (pic de la famille) | 
| Écriture aléatoire (4K) | 800K IOPS | 
| Puissance et endurance |  | 
| Puissance (active / inactive) | 23 W typ. / 5 W typ. Cinq niveaux de puissance configurables, de 5 W à 25 W | 
| Endurance | 3.0 DWPD (sur une base de 5 ans) 4.98 DWPD (sur une base de 3 ans) 70 PBW à 12.8 To | 
| Fiabilité et sécurité |  | 
| MTBF / UBER | de 2,500,000 heures Testé à 1E-18 | 
| Sécurité | TCG Opal 2.02 (variante SED) Certifiable FIPS 140-3 niveau 2 Démarrage sécurisé conforme à la norme OCP et signature du firmware Attestation du dispositif, révocation de la clé | 
| Garantie | 5 ans | 
Performances du Solidigm D7-PS1030
Le contexte des graphiques suivants est le suivant : le PS1030 se positionne dans notre comparatif comme le disque polyvalent par rapport aux disques axés sur la lecture intensive qui ont dominé les tests récents de la 5e génération. Face au KIOXIA CD9P-R testé en juin, le PS1030 affiche une vitesse de lecture séquentielle inférieure (14 500 Mo/s contre 14 800 Mo/s), mais une vitesse d'écriture aléatoire presque deux fois supérieure (800 000 IOPS contre 450 000 IOPS) et un budget d'écriture trois fois plus important (3 DWPD contre 1 DWPD). Son principal concurrent, en termes de performances, est le Micron 7600 MAX , l'autre disque 3 DWPD de ce comparatif. Le Micron 9550 MAX, avec la même capacité de 12.8 To que notre modèle, offre quant à lui des performances mixtes comparables.
Plateforme de test de conduite
Nous utilisons un serveur Dell PowerEdge R760 exécutant Ubuntu 22.04.2 LTS comme plateforme de test pour toutes les charges de travail présentées dans ce rapport. Équipé d'un boîtier JBOF Serial Cables Gen5, il offre une large compatibilité avec les SSD U.2, E1.S, E3.S et M.2. La configuration de notre système est détaillée ci-dessous :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
- Nvidia L4
Comparaison des lecteurs
- KIOXIA CD9P-R 7.68 To (1 DWPD)
- KIOXIA CM9-R 15.36 To (1 DWPD)
- Disque dur SanDisk DC SN861 7.68 To (1 DWPD)
- Solidigm PS1010 7.68 To (1 DWPD)
- Micron 7600 MAX 6.4 To (3 DWPD)
- Micron 9550 MAX 12.8 To (3 DWPD)
- Micron 9550 Pro 7.68 To (1 DWPD)
Benchmark de point de contrôle DLIO
Pour évaluer les performances réelles des SSD dans les environnements d'entraînement d'IA, nous avons utilisé l'outil de test DLIO (Data and Learning Input/Output). Développé par le Laboratoire national d'Argonne, DLIO est spécifiquement conçu pour tester les modèles d'E/S dans les charges de travail d'apprentissage profond. Il permet de comprendre comment les systèmes de stockage gèrent des problématiques telles que la création de points de contrôle, l'ingestion de données et l'entraînement des modèles.
Le tableau ci-dessous présente le temps moyen d'exécution des points de contrôle pour chaque disque, sur trois passages ; plus le temps est court, mieux c'est. Remarque importante : le nombre de points de contrôle par passage est proportionnel à la capacité du disque. Ainsi, les disques de plus grande capacité enregistrent davantage de points de contrôle, et les résultats par point de contrôle ne sont pas directement comparables entre des disques de tailles différentes. C'est pourquoi nous publions des moyennes par passage, qui normalisent les performances de chaque disque et permettent une comparaison directe. Lors de l'entraînement de modèles d'apprentissage automatique, les points de contrôle sont essentiels pour sauvegarder régulièrement l'état du modèle et éviter toute perte de données en cas d'interruption ou de coupure de courant. Cette exigence de stockage requiert des performances robustes, notamment sous des charges de travail soutenues ou intensives. Nous avons utilisé la version 2.0 du benchmark DLIO, publiée le 13 août 2024.
