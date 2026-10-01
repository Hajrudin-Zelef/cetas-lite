---
id: collect-261001-general-networking/general-networking/fr-review-adaptec-series-8-raid-controllers-review-1f1b7292-1
title: "fr-review-adaptec-series-8-raid-controllers-review-1f1b7292"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-general-networking/fr-review-adaptec-series-8-raid-controllers-review-1f1b7292.md
source_anchor: ""
source_lines: [1, 76]
sha256: 131079155125dec93ce73a7de0b87ab325d34919c526bac6284457c8e5ed5e22
---

# fr-review-adaptec-series-8-raid-controllers-review-1f1b7292

La gamme Adaptec Série 8 est composée d'adaptateurs RAID 12 Gb/s dotés d'une architecture PMC complète intégrant les contrôleurs de protocole d'E/S, les contrôleurs RoC, les expandeurs SAS et le logiciel PMC. Les cartes Adaptec Série 8 offrent des performances nettement supérieures à la génération précédente , avec des vitesses considérablement accrues. Elles sont également équipées de la technologie maxCache Plus, qui prend en charge la mise en cache et la hiérarchisation des données, permettant ainsi aux responsables informatiques, aux intégrateurs système et aux éditeurs de logiciels indépendants d'optimiser les performances et le rapport qualité-prix de leurs systèmes de stockage. maxCache Plus offre également aux utilisateurs la possibilité de configurer leurs périphériques de stockage dans des environnements serveur et d'utiliser les supports les plus rapides comme stockage et non plus seulement comme cache. Ainsi, le potentiel des disques 12 Gb/s les plus récents du marché est pleinement exploité.
Nous avions précédemment testé la gamme Adaptec Série 7 (71605Q) et étions satisfaits de ses performances en termes de débit. Cependant, c'est surtout sa grande flexibilité, avec ses 24 ports sans extension, qui nous avait impressionnés, établissant ainsi une nouvelle référence en matière de polyvalence sur le marché. À cette époque, Adaptec était (et est toujours) en concurrence féroce avec LSI, les deux entreprises rivalisant pour proposer soit un grand nombre de ports, soit des performances élevées. La Série 8, quant à elle, offre aux entreprises 16 ports et une bande passante impressionnante de 12 Gbit/s, démontrant ainsi la volonté d'Adaptec de repousser les limites du marché. Cette innovation constante contribue à faire d'Adaptec, l'entreprise l'espère, le leader incontesté du marché des adaptateurs RAID.
La gamme Series 8 d'Adaptec comprend cinq modèles différents :
- 81605ZQ: C'est la carte que nous utilisons dans notre Serveur de fichiers Supermicro, qui dispose de 16 ports internes avec protection de cache SuperCap (incluse) de sauvegarde Flash (intégrée).
- 8885 et 8885Q : 8 ports internes et 8 ports externes ; le 8885Q est équipé de la protection de cache AFM-700 en option.
- 8805: La seule carte de la gamme 8-Series avec 8 ports (internes) au total.
Tous les modèles prennent en charge les niveaux RAID 0, 1, 1E, 5, 6, 10, 50, 50 et 60.
Les 8885 et 8885Q coûtent actuellement respectivement 725.00 $ et 1100.00 $, tandis que les 81605ZQ et 8805 sont disponibles pour 1065.00 $ et 640.00 $. Toutes les versions sont livrées avec une garantie de trois ans. Notre examen se concentrera sur l'unité 81605ZQ.
Spécifications du contrôleur RAID Adaptec série 8
8160ZQ
- Numéro de pièce de commande : 2281600-R (simple)
- Niveaux RAID : 0,1,1E,5,6,10, 50, 60
- Facteur de forme:
  - MD2 – Profil bas
  - 2.535"H x 6.6"L (64mm x 167mm)
- Ports : 16 internes
- Connecteurs : 4 X SFF-8643
- Interface de bus : PCIe Gen8 à 3 voies
- Processeur : PMC PM8063
- Cache: 1024MB
- Protection du cache : Flash backup (intégré) Supercap (inclus)
8885Q
- Numéro de pièce de commande : 2277100-R (simple)
- Niveaux RAID : 0,1,1E,5,6,10, 50, 60
- Facteur de forme:
  - MD2 – Profil bas
  - 2.535"H x 6.6"L (64mm x 167mm)
- Ports : 8 internes/8 externes
- Connecteurs:
  - 2 x SFF-8643
  - 2 x SFF-8644
- Interface de bus : PCIe Gen8 à 3 voies
- Processeur : PMC PM8063
- Cache: 1024MB
- Protection du cache : AFM-700 (inclus)
8885
- Numéro de pièce de commande : 2277000-R (simple)
- Niveaux RAID : 0,1,1E,5,6,10, 50, 60
- Facteur de forme:
  - MD2 – Profil bas
  - 2.535"H x 6.6"L (64mm x 167mm)
- Ports : 8 internes/8 externes
- Connecteurs:
  - 2 x SFF-8643
  - 2 x SFF-8644
- Interface de bus : PCIe Gen8 à 3 voies
- Processeur : PMC PM8063
- Cache: 1024MB
- Protection du cache : AFM-700 (en option)
8805
- Numéro de pièce de commande : 2277500-R (simple)
- Niveaux RAID : 0,1,1E,5,6,10, 50, 60
- Facteur de forme:
  - MD2 – Profil bas
  - 2.535"H x 6.6"L (64mm x 167mm)
- Ports : 8 internes
- Connecteurs : 2 X SFF-8643
- Interface de bus : PCIe Gen8 à 3 voies
- Processeur : PMC PM8063
- Cache: 1024MB
- Protection du cache : AFM-700 (en option)
Spécifications générales:
- Dimensions physiques : 2.535" H x 6.6" L (64 mm x 167 mm)
- Température de fonctionnement : 0°C à 55°C* (avec débit d'air de 200 LFM, sans flash) ; 0°C à 50°C* (avec débit d'air de 200 LFM, avec flash)
- Courant de fonctionnement : 1.0 A à 3.3 V CC ; 1.1 A à 12.0 Vcc 8805, 8885, 8885Q ; 1.5 A à 3.3 Vcc ; 1.0 A à 12.0 Vcc 81605ZQ
- Certifications réglementaires CE, FCC, UL, C-tick, VCCI, KCC, CNS
- Conformité environnementale RoHS
- MTBF en test maintenant
- Garantie: ans 3
Construire et concevoir
Au cœur de chaque nouvelle carte RAID Adaptec de la série 8 se trouve le contrôleur RAID sur puce PM8063, qui s'interface avec le système hôte via une connexion PCIe Gen8 à 3 voies offrant jusqu'à 8 Go/s de bande passante brute combinée avec XOR en ligne pour réduire l'accès à la DDR3 pour optimiser les performances. Le PM8063 prend en charge 16 appareils à des vitesses natives de 12 Gb/s. Pour les besoins de cet examen, nous nous concentrons sur le nouveau 81605ZQ, qui prend en charge 16 périphériques internes, ainsi qu'un module flash intégré (sans avoir besoin d'une carte fille supplémentaire) pour la protection des données en cas de panne de courant. scénario.
La carte 81605ZQ se présente au format HHHL compact et intègre entièrement sa fonction de sauvegarde flash. Contrairement aux autres cartes de la série Adaptec, qui nécessitent des cartes filles additionnelles pour activer cette fonctionnalité, cette conception compacte libère de l'espace au profit de deux ports mini SAS HD supplémentaires à l'arrière. Bien que les autres cartes de la série Adaptec 8 prennent également en charge 16 périphériques, aucune ne propose 16 ports internes. Elles offrent en revanche une combinaison de 8 ports internes et 8 ports externes. Le seul périphérique externe requis pour la 81605ZQ est une batterie de sauvegarde, qui peut être placée quasiment n'importe où à l'intérieur du châssis du serveur. Dans notre cas, pour le serveur Supermicro SuperStorage 2027R-AR24NV, nous avons pu positionner les batteries directement derrière les baies de stockage avant, dans la zone de ventilation la plus optimale du châssis.
Enfin, sont les connecteurs internes SFF-8643 ; l'adaptateur se connecte via un emplacement PCIe Gen3 x8 ; cependant, il est également rétrocompatible avec les systèmes PCIe Gen2. L'adaptateur 81605ZQ n'est pas équipé de connecteurs externes, contrairement à deux autres modèles, les 8885 et 885Q, ce qui leur donne une certaine flexibilité.
Configuration des tests
La carte Adaptec série 8 81605ZQ est utilisée dans notre serveur de stockage Supermicro SuperStorage Server 2027R-AR24NV , fonctionnant sous Windows Server 2012 R2 et faisant office de serveur de fichiers pour le laboratoire de test StorageReview. Le SuperStorage Server AR24NV est une plateforme de stockage haut de gamme conçue spécifiquement pour exploiter les périphériques de stockage SAS3 les plus performants.
Les deux principaux sous-systèmes de l'AR24NV comprennent le châssis SC216A-R920LPB 2U 24 baies et la carte serveur biprocesseur X9DRH-iF-NV. Sa caractéristique la plus importante, cependant, est un tout nouveau fond de panier à connexion directe SAS24 3 Gb/s à 12 baies optimisé pour la dernière génération de SSD, où la série Adaptec 8 entre en jeu.
