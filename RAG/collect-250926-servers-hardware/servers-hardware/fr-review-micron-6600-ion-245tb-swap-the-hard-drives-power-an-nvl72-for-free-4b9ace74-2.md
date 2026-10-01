---
id: collect-250926-servers-hardware/servers-hardware/fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74-2
title: "fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["blackwell", "gpu"]
source: docs/RAG/clean4/fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74.md
source_anchor: ""
source_lines: [37, 65]
sha256: 5f823ffa1df63c611bef9235b537ae455da167200299047d8fd58aaa8646f9f5
---

# fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74

Nous avons testé les performances et la consommation du SSD Micron 6600 ION 245 To sur un serveur Dell R5715 , aux côtés d'une configuration de huit disques durs de 30 To en RAID 5. Cette plateforme a été choisie pour sa compatibilité avec les disques durs 3.5 pouces et son contrôleur RAID PERC12, permettant ainsi d'optimiser les performances des disques.
Étant donné que la surcharge de la plateforme R5715 est identique dans les deux configurations, la différence de consommation par charge de travail entre le Micron 6600 ION et huit disques durs en RAID 5 représente la différence de consommation attribuable au stockage. Nous avons utilisé notre module d'analyse de consommation électrique Quarch pour mesurer la consommation du serveur. Les mesures de consommation en veille incluent tous les composants du serveur au repos ; l'écart reflète la différence de consommation des disques, de l'électronique en veille et de la réactivité des ventilateurs entre les deux configurations. Les différences de consommation en charge de travail suivent la même logique : la consommation du serveur, hors stockage, ne varie que marginalement d'une charge de travail à l'autre. Les économies réalisées en charge de travail proviennent donc du sous-système de stockage, ainsi que d'une légère augmentation de la charge du processeur dans les zones où le SSD génère davantage d'E/S.
Configurations testées
- RAID 5 : Huit disques durs Seagate Exos M de 30 To sont installés via le fond de panier standard du R5715, connectés au contrôleur RAID intégré et configurés en RAID 5, offrant une capacité utilisable de 210 To. Le fond de panier, le contrôleur et les huit disques étaient entièrement installés et fonctionnels lors des mesures.
- SSD NVMe : Le fond de panier du disque dur a été entièrement retiré du châssis. Une carte d'extension E3.L dédiée a été installée à sa place, et le SSD Micron 6600 ION 245 To a été inséré dans cette carte et connecté via PCIe Gen5 x4 NVMe, présentant ainsi la totalité des 245 To comme un seul espace de noms. Aucun disque dur en rotation, aucune activité de contrôleur RAID ni aucun composant électronique du fond de panier n'étaient présents lors des mesures du SSD.
Sur le papier, le SSD Micron 6600 ION de 245 To affiche une consommation en veille inférieure à 5 W et une consommation maximale inférieure à 30 W. À titre de comparaison, les disques durs Seagate Exos M de 30 To consomment 6.9 W en veille et 9.5 W en fonctionnement.
Charges de travail FIO
- Séquentiel 128K : Les opérations de lecture et d'écriture ont été testées indépendamment. Nous nous sommes concentrés sur les transferts séquentiels monothread pour les SSD et les HDD, reflétant les schémas d'analyse et d'ingestion qui prédominent dans les niveaux de lac de données à haute capacité.
Chaque charge de travail s'est exécutée pendant 3 minutes, les données de consommation et de performance étant moyennées sur cette durée.
Comparaison de la consommation électrique en fonction de la charge de travail
Au repos, l'écart est immédiat et structurel :
- Tirage au sort inactif : 115.9 W pour la configuration 6600 ION contre 173.5 W pour huit disques durs Exos M sur leur fond de panier avec le contrôleur RAID.
- L'écart : 57.6 watts. La configuration disque dur consomme 50 % d'énergie en plus au repos, ou, en d'autres termes, la configuration flash réduit la consommation d'énergie en veille d'un tiers.
- Normalisé à la capacité : 0.47 W/TB en veille contre 0.72 W/TB pour la baie de disques durs à capacité brute égale, soit une réduction de 35 % de la consommation d'énergie par téraoctet.
Ce chiffre, en sommeil, est plus important qu'il n'y paraît. Le stockage n'est pas toujours saturé, notamment dans les environnements de stockage d'objets à grande échelle, d'archivage et de lacs de données d'IA où les données sont ingérées par vagues et lues de manière intermittente.
Sous charge, la séparation se maintient dans les deux sens de circulation :
- Lecture séquentielle : 175.0 W pour la configuration 6600 ION contre 224.7 W pour la baie de disques durs.
- Ecriture séquentielle : 170.2W contre 234.0W.
- Pas de chevauchement : La configuration flash sous sa charge la plus élevée mesurée a consommé 49.7 W de moins que la baie de disques durs sous sa charge la plus faible.
- Le chiffre principal : La consommation en écriture séquentielle de 170.2 W est inférieure à celle de la baie de disques durs en veille complète, qui est de 173.5 W.
Un quart de pétaoctet d'écriture flash à pleine vitesse consomme moins d'énergie système que huit disques durs inactifs.
La puissance consommée est instantanée ; l’énergie est la consommation multipliée par le temps. Cette distinction est cruciale ici, car la configuration flash ne se contente pas de consommer moins, elle achève les opérations plus rapidement. Notre test du Micron 6600 ION de 245 To a mesuré 12 729,8 Mo/s en lecture séquentielle 128K, tandis que notre test du Seagate Exos M de 30 To a mesuré 292 Mo/s sur un seul disque. En supposant une mise à l’échelle linéaire parfaite du RAID 5 sur ses huit disques (cas idéal qu’aucun RAID 5 n’atteint en pratique), le débit côté disque est plafonné à environ 2 336 Mo/s. Si l’on compare ces valeurs à la puissance consommée mesurée lors des lectures séquentielles, l’écart d’efficacité dépasse largement l’écart de consommation.
- Débit par watt : 72.7 Mo/s par watt pour la configuration 6600 ION contre 10.4 Mo/s par watt pour la baie de disques durs.
- Énergie nécessaire pour lire un téraoctet : environ 3.8 wattheures pour la mémoire flash contre environ 26.7 wattheures pour la mémoire disque.
- Dans les deux cas, environ 7 fois, et c'est un plancher, puisqu'il attribue aux disques durs une capacité d'adaptation qu'ils ne fournissent pas.
Les chiffres de débit proviennent de nos tests produits respectifs plutôt que des tests de consommation électrique mentionnés ci-dessus ; le ratio est donc calculé par déduction plutôt que mesuré de bout en bout, et il décrit les lectures séquentielles, le modèle pour lequel ce disque a été conçu.
Micron annonce un avantage bien plus important d'après ses propres tests, avec une efficacité énergétique jusqu'à 84 fois supérieure pour le prétraitement IA. Micron isole la consommation au niveau du disque, ce qui élimine la consommation partagée d'environ 116 W présente dans nos deux configurations. La comparaison porte sur seize disques durs au lieu de huit, et l'entreprise utilise des modèles d'accès au pipeline IA qui sollicitent davantage les disques durs rotatifs que les lectures séquentielles. Notre gain de 7x représente le minimum au niveau du système ; le gain maximal de Micron peut être bien supérieur selon le flux de travail appliqué.
Les mathématiques de Watts et des racks
Le cadre pertinent pour évaluer la consommation énergétique du stockage n'est pas la facture d'électricité ; en effet, les budgets énergétiques des centres de données sont limités, l'accès au réseau est de plus en plus restreint et la capacité de refroidissement constitue une contrainte majeure pour les déploiements possibles. Chaque watt consommé par un système de stockage est un watt qui ne peut être alloué à un GPU. Dans les environnements où les déploiements Blackwell sont limités par la puissance disponible plutôt que par des contraintes d'approvisionnement ou de disponibilité logicielle, l'efficacité du stockage est un choix stratégique en termes de capacité de calcul.
