---
id: collect-260926-rattrapage/rattrapage/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752-2
title: "Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["arr", "gpu", "intel"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752.md
source_anchor: ""
source_lines: [69, 122]
sha256: 71a1a9b4f0c29d153311aaec152d1d33e7ea8fde7efa136180ad20d91ef41d6b
---

# Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant

| Spécifications | DÉTAILS |
|---|---|
| Facteur de forme | rack 2U |
| Processeur | Jusqu'à 2 processeurs Intel Xeon série 6700/6400 (cœurs P) ; jusqu'à 86 cœurs par processeur ; TDP jusqu'à 350 W |
| Mémoire | 32 emplacements DIMM (16 par processeur) ; prend en charge les modules TruDDR5 RDIMM jusqu'à 6 400 MHz et les modules MRDIMM jusqu'à 8 000 MHz. |
| Extension de mémoire | Prend en charge les modules de mémoire CXL 2.0 au format E3.S 2T (jusqu'à 12 modules DIMM). |
| Mémoire maximale | Jusqu'à 8 To de mémoire système (utilisant des modules RDIMM 3DS de 256 Go) |
| Baies de lecteur de disque | Avant : Jusqu'à 24 disques NVMe/SAS/SATA 2.5 pouces, 16 disques SAS/SATA 3.5 pouces ou 32 disques NVMe E3.S. Milieu : Jusqu'à 8 disques de 2.5 pouces à échange simple. Arrière : Jusqu'à 8 disques de 2.5 pouces ou 4 disques de 3.5 pouces remplaçables à chaud |
| Contrôleur de stockage | Jusqu'à 36 ports NVMe intégrés (connectivité 1:1) ; prise en charge RAID/HBA |
| Interfaces réseau | Deux emplacements OCP 3.0 SFF avec interface hôte PCIe 5.0 x16 |
| Emplacements d'extension PCI | Jusqu'à 10 emplacements PCIe 5.0 (à l'arrière) ; emplacements PCIe avant en option |
| Prise en charge du GPU | Jusqu'à 10 GPU simple largeur ou 2 GPU double largeur |
| Ports | Avant : Port de diagnostic externe, USB en option et Mini-DP. Arrière : 2 ports USB 3.0, 1 port VGA, 1 port RJ-45 de gestion, port série en option |
| Refroidissement | Jusqu'à 6 ventilateurs remplaçables à chaud (redondance N+1) ; modules de refroidissement liquide Neptune en option |
| Alimentation | Jusqu'à deux alimentations AC/DC redondantes remplaçables à chaud (800 W – 3200 W) |
| Systems Management | Contrôleur XClarity 3 (XCC3) ; XCC3 Premier en option |
| Caractéristiques de sécurité | TPM 2.0, racine de confiance PFR (NIST SP800-193), détection d'intrusion dans le châssis, lunette verrouillable |
| Systèmes d'exploitation | Serveur Microsoft Windows, RHEL, SLES, Serveur Ubuntu |
| Garantie | Garantie de base de 3 ans (configurable) |
| Dimensions | 87 mm (3.4 po) H x 440 mm (17.3 po) L x 800 mm (31.5 po) P |
| Poids | Poids maximal : 38.8 kg (85.5 lb) |

## Conception et fabrication du Lenovo ThinkSystem SR650 V4

Le Lenovo ThinkSystem SR650 V4 est une plateforme 2U standard qui optimise son encombrement, en conciliant densité des composants, ventilation et facilité de maintenance. Son châssis mesure 3.4 cm de hauteur, 17.3 cm de largeur et 31.5 cm de profondeur, pour un poids maximal de 85.5 kg. Ces dimensions le placent au même niveau que les autres serveurs d'entreprise 2U, tout en offrant une profondeur suffisante pour les processeurs hautes performances, les configurations mémoire haute densité et les fonds de panier de stockage flexibles.

Le design industriel ThinkSystem de Lenovo est immédiatement reconnaissable, avec son châssis en acier rigide et son étiquetage clair et fonctionnel. À l'intérieur, l'agencement est épuré et fonctionnel, privilégiant la circulation d'air directe, les zones de composants modulaires et l'accès sans outil pour la maintenance courante. La construction générale inspire confiance et témoigne d'une ingénierie de pointe, reflétant une plateforme conçue pour un fonctionnement continu en environnement de centre de données plutôt que pour l'esthétique.

### Panneau avant

Le panneau avant est hautement configurable. Notre système de test est équipé d'un fond de panier 8 baies 2.5 pouces, mais le châssis peut facilement être configuré avec 16 ou même 24 baies. Les clients peuvent également opter pour des baies 3.5 pouces pour une capacité brute supérieure ou les nouveaux fonds de panier E3.S pour une densité NVMe haute performance. La compatibilité AnyBay permet aux disques SAS, SATA et NVMe de coexister dans le même châssis, offrant une grande flexibilité pour les configurations de stockage mixtes.

Le panneau avant comprend les commandes essentielles : un bouton marche/arrêt, un bouton d'identification et des voyants d'état indiquant le fonctionnement du système et l'activité réseau. Une languette d'information amovible permet d'accéder rapidement à l'étiquette d'accès réseau XCC et au numéro de série. La connectivité locale optionnelle inclut un port Mini DisplayPort et des ports USB pour l'accès via un chariot d'intervention, indispensable pour le dépannage dans les centres de données hors ligne.

### Panneau arrière

L'arrière du châssis est conçu pour une densité d'E/S maximale. Il peut accueillir jusqu'à 10 emplacements PCIe Gen 5, selon la configuration des cartes d'extension. L'intégration de deux emplacements OCP 3.0 libère les cartes d'extension PCIe standard pour d'autres cartes d'extension, telles que des GPU ou des contrôleurs de stockage.

Outre l'extension, le panneau arrière abrite le port réseau de gestion BMC intégré, deux ports USB-A et une sortie VGA pour l'accès à la console locale. Deux alimentations remplaçables à chaud sont montées à l'arrière, préservant ainsi la facilité d'entretien sans impacter la circulation de l'air ni les possibilités d'extension.

### Interne

À l'intérieur, l'agencement est épuré et optimisé pour la circulation de l'air. 32 emplacements DIMM encadrent deux sockets CPU, disposés de manière à assurer un refroidissement optimal de l'avant vers l'arrière grâce à un ensemble de six ventilateurs haute performance remplaçables à chaud, situés à l'avant du châssis. Ces ventilateurs sont entièrement démontables et remplaçables sans outil, et le plateau de ventilateurs est amovible sans outil. Le retrait de ce plateau offre un accès direct et dégagé à la baie de stockage AnyBay, simplifiant ainsi la maintenance, les mises à niveau et le changement de fond de panier.

Les câbles de stockage sont soigneusement acheminés le long des côtés du châssis afin de ne pas obstruer la circulation de l'air vers les processeurs et la mémoire. Pour les configurations utilisant des modules de refroidissement liquide Neptune, l'agencement interne est légèrement modifié afin d'intégrer les circuits de refroidissement, mais l'accessibilité des composants reste optimale. Le module de disque de démarrage M.2 est monté sur le déflecteur d'air ou la cage de disques, ce qui permet un accès facile sans avoir à démonter d'autres composants.

### Contrôleur XClarity 3

Le SR650 V4 intègre le nouveau contrôleur XClarity 3 (XCC3). Ce moteur de gestion offre des performances nettement supérieures aux générations précédentes, avec des temps de démarrage plus rapides et une interface HTML5 plus réactive. Le gestionnaire de ressources de la plateforme fournit des données télémétriques détaillées sur l'alimentation et la température, permettant aux administrateurs d'optimiser l'efficacité du centre de données.

XCC3 prend en charge un large éventail de fonctions de gestion à distance, notamment le contrôle à distance (KVM), le montage de supports virtuels et les mises à jour du firmware. Son interface intuitive offre une vue d'ensemble de l'état du système, des événements en cours et de l'inventaire matériel. Pour l'automatisation, XCC3 est entièrement compatible avec les API REST de Redfish.

## Performances du Lenovo ThinkSystem SR650 V4

