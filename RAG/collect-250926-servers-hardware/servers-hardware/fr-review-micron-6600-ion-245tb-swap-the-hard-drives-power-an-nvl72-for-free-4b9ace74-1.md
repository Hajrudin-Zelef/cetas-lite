---
id: collect-250926-servers-hardware/servers-hardware/fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74-1
title: "fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "benchmarks", "gpu", "mai", "nand"]
source: docs/RAG/clean4/fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74.md
source_anchor: ""
source_lines: [1, 36]
sha256: ad0ace7ddf187be15eaffad3b6d99595075e6957ad21f313928b4f8a9adfc377
---

# fr-review-micron-6600-ion-245tb-swap-the-hard-drives-power-an-nvl72-for-free-4b9ace74

Pendant deux décennies, le débat SSD contre HDD s'est soldé de la même manière : la mémoire flash l'emportait sur les performances, le disque sur le prix au téraoctet, et l'ampleur de cet écart de prix a tranché en faveur du stockage de masse. Avec la maturité des technologies de stockage et l'essor de l'IA, cette vision est clairement obsolète. Les plus grands opérateurs de centres de données ne sont plus limités par leur budget d'achat, mais par leurs capacités en alimentation, refroidissement et intégration physique. Lors de notre test du SSD Micron 6600 ION de 245 To en début d'année, les résultats des benchmarks étaient révélateurs, mais le calcul de la consommation en watts et en unités rack était encore plus important. Nos mesures ont montré qu'un seul SSD Micron 6600 ION équivalait à huit disques durs nearline, et que la configuration flash en écriture à pleine vitesse consommait moins d'énergie que la configuration HDD au repos. Étendez cette capacité d'échange à un exaoctet, et l'encombrement passe de 22 racks (densité maximale de disques durs) à 6 racks de mémoire flash, libérant ainsi 16 emplacements pour le calcul. Un disque capable d'intégrer près d'un quart de pétaoctet dans un seul emplacement bouleverse la planification des exaoctets, et dans les infrastructures actuelles, la puissance et le nombre de racks sont réservés des années à l'avance.
Depuis sa commercialisation en mai, Micron a activement promu le SSD 6600 ION de 245 To, et nous disposons des données de test permettant de vérifier ces affirmations. Les performances FIO et la consommation énergétique, ainsi que les résultats de notre test complet du SSD ( GPU Direct Storage et DLIO Checkpointing ), confirment une version précise de son argumentaire : pour le stockage de masse intensif en lecture, ce SSD offre une densité et une efficacité inégalées par les disques durs nearline, avec des compromis prévisibles lorsque la charge de travail correspond à sa conception.
Points clés à retenir
- Un disque dur a remplacé huit : Un seul disque dur Micron 6600 ION de 245 To a remplacé huit disques Seagate Exos M de 30 To en RAID 5 dans le même Dell R5715, avec le fond de panier du disque dur et le contrôleur RAID retirés du châssis.
- L'écriture flash consomme moins d'énergie que la mise en veille du disque : 170.2 W en écriture séquentielle contre 173.5 W pour la configuration du disque dur au repos. Chaque état mesuré, libéré, a consommé entre 49.7 W et 63.8 W, soit une moyenne de 55 W par unité.
- Les watts se convertissent directement en calculs : 44 kW libérés par rack de serveurs flash. Avec 2 182 disques, soit moins de trois racks, l’échange permet de libérer les 120 kW nécessaires au fonctionnement d’un GB200 NVL72 complet.
- L'efficacité, pas seulement le match nul : 72.7 Mo/s par watt contre 10.4 pour la baie de disques durs, ou 3.8 watt-heures pour lire un téraoctet contre 26.7, soit environ 7 fois plus en lectures séquentielles.
- Un exaoctet dans 6 racks au lieu de 22 : 4.2 fois la capacité par rack avec les deux côtés à la densité optimale, ce qui donne 16 emplacements de rack et environ 320 pieds carrés d'espace libre par exaoctet.
L'économie des centres de données est en train d'être réinventée.
L'Agence internationale de l'énergie prévoit que la consommation d'électricité des centres de données plus que doublera d'ici 2030 pour atteindre environ 945 TWh, l'intelligence artificielle étant le principal facteur de croissance. Dans les infrastructures qui absorbent cette augmentation, la puissance, l'espace et le refroidissement sont devenus les contraintes majeures, remplaçant le budget. Dans un bâtiment à capacité énergétique limitée, chaque watt et chaque emplacement de rack alloué au stockage représente une capacité inutilisable pour les GPU générateurs de revenus.
Le SSD Micron 6600 ION de 245 To a été conçu pour répondre à ce besoin. Les disques durs nearline restent plus avantageux en termes de coût d'acquisition par téraoctet, et pour les archives froides et les données rarement consultées, cet avantage demeure décisif. Cependant, les opérateurs qui planifient à l'échelle de l'exaoctet évaluent différemment le coût total de possession et d'exploitation de cette capacité sur toute sa durée de vie, avec une consommation d'énergie et un encombrement fixes. Mesurée de cette façon, la densité du disque se traduit directement en watts économisés, en emplacements rack et en marge de refroidissement, autrement dit, en puissance de calcul que l'infrastructure peut désormais supporter. Travis Vigil, vice-président senior de la gestion des produits ISG chez Dell Technologies, a présenté l'analyse du coût total de possession du point de vue du fournisseur lors du lancement , qualifiant le 6600 ION de « réduction significative du coût total de possession pour les clients qui déploient des environnements d'IA et de centres de données à grande échelle ». Jeff Janukowicz, vice-président de la recherche chez IDC pour les disques SSD et les technologies associées, a décrit cette même évolution dès le lancement : « La croissance rapide des ensembles de données d’IA fait évoluer la rentabilité du stockage, des disques individuels à l’efficacité au niveau des racks. Les opérateurs ont besoin d’une capacité utilisable accrue par rack, tout en respectant des contraintes strictes de consommation d’énergie et de refroidissement. » Les sections suivantes présentent nos propres mesures à l’appui de cette affirmation.
Présentation du Micron 6600 ION 245TB
Le SSD ION 6600 de 245 To est actuellement le SSD commercial le plus performant. Disponible depuis le 5 mai 2026 aux formats E3.L (9.5 mm) et U.2 (15 mm), il est basé sur la mémoire NAND QLC G9 de neuvième génération de Micron, dotée d'une architecture à six plans qui atteint des vitesses d'E/S NAND de 3.6 Go/s, les plus rapides actuellement disponibles pour un SSD de centre de données. Le contrôleur utilise une interface PCIe Gen5 x4 et le disque répond à toutes les exigences de conformité attendues par les équipes d'achat : OCP 2.6, NVMe 2.0d, éligibilité TAA et certification FIPS 140-3 niveau 2 avec prise en charge CNSA 2.0 et SPDM 1.2.
La fiche technique indique clairement l'objectif de conception. Les vitesses de lecture séquentielles atteignent 13 700 Mo/s contre 3 000 Mo/s en écriture, et les performances en lecture aléatoire 1.78 million d'IOPS contre 42 000 IOPS en écriture aléatoire. Le modèle haute capacité utilise une unité d'indirection de 16 Ko au lieu de 4 Ko, ce qui explique son endurance de 1.0 SDWPD pour les écritures séquentielles de 128 Ko et de 0.3 RDWPD pour les écritures aléatoires de 16 Ko. Il s'agit d'un SSD haute capacité optimisé pour la lecture.
| Spécifications | Micron 6600 ION 245 To | 
|---|---|
| En savoir plus sur la plateforme |  | 
| Capacités | 245TB | 
| Facteurs de forme | E3.L (9.5 mm) U.2 (15 mm) | 
| Interface | PCIe Gen5 x4, NVMe 2.0d | 
| NON | Mémoire NAND Micron G9 QLC, six plans, E/S 3.6 Go/s | 
| Performances |  | 
| Lecture séquentielle | 13,700 Mo / s | 
| Écriture séquentielle | 3,000 Mo / s | 
| Lecture aléatoire | 1,780,000 IOPS | 
| Écriture aléatoire (4K/16K) | 42,000 IOPS | 
| Latence (QD1, Lecture/Écriture) | 100 µs / 20 µs | 
| Puissance et endurance |  | 
| Puissance maximale | ≤ 30W | 
| Puissance inactive | ≤ 5W | 
| Endurance | 1.0 SDWPD (128 Ko séquentiel) 0.3 RDWPD (16K aléatoire) | 
| MTTF / UBER | 2.5 millions d'heures à 50 °C <1 secteur pour 1017 bits lus | 
| Fonctionnalité |  | 
| Conformité | OCP2.6 NVMe 2.0d NVMe-MI 1.2d TAA | 
| Sécurité | Certifiable FIPS 140-3 L2 CNSA 2.0 SPDM 1.2 Options Micron SEE et SED | 
Ce que nos tests ont montré
