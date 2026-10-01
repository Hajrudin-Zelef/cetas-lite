---
id: collect-261001-general-networking/general-networking/fr-review-adaptec-series-7-raid-controllers-review-7b01350f-1
title: "fr-review-adaptec-series-7-raid-controllers-review-7b01350f"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["dram", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-adaptec-series-7-raid-controllers-review-7b01350f.md
source_anchor: ""
source_lines: [1, 36]
sha256: 739a239ea8b066127e0d49b36f2063606bd97a3de32c6c0e3595aa877bd16bff
---

# fr-review-adaptec-series-7-raid-controllers-review-7b01350f

La famille Series 7 est la gamme polyvalente de contrôleurs RAID 6 Gb/s d'Adaptec, amenant Adaptec sur le marché PCI Express 3.0. La force motrice de la série 7 est qu'aucun extenseur n'est nécessaire pour connecter jusqu'à 24 appareils, ce qui permet une conception moins coûteuse et moins complexe. Il y a aussi clairement un avantage de densité avec un déploiement basé sur Adaptec, car la plupart des solutions concurrentes utilisent des conceptions à 8 ports qui nécessitent des extensions en fonction du nombre de ports. La nouvelle famille Series 7 est la première génération de contrôleurs d'Adaptec qui bénéficie des avantages de la fusion avec PMC en termes de capacités d'ingénierie et de production.
La série 7 est basée sur le RAID-on-Chip PM24 à 8015 ports de PMC. Cette technologie combine une interface x8 PCIe Gen3 avec des ports SAS 6 Gb/s, créant une nouvelle génération d'adaptateurs RAID hautes performances capables de rivaliser avec n'importe quel autre ROC sur le marché. Les adaptateurs RAID Adaptec série 7 sont disponibles avec 8, 16 ou 24 ports SAS/SATA natifs, ce qui offre aux constructeurs de systèmes et à ceux qui préfèrent rouler eux-mêmes une énorme flexibilité. Pour les systèmes centrés sur le stockage, la libération de ports PCIe peut faire une différence substantielle lorsqu'il s'agit d'un stockage PCIe hautes performances et de plusieurs cartes d'interconnexion. Dans certains cas, les cartes Adaptec seront la seule option viable pour résoudre les déploiements de stockage créatifs.
Adaptec propose également la gamme Series 7Q, que nous exploitons dans un certain nombre de projets. La plus grande différence entre les deux versions est que les contrôleurs 7Q disposent de maxCache 3.0, offrant une mise en cache SSD en lecture et en écriture pour augmenter les IOPS et réduire la latence. Avec la possibilité de désigner les SSD comme grands caches, ce type de configuration augmente considérablement la quantité d'espace total de mise en cache disponible par rapport aux configurations de cartes RAID plus standard. De plus, maxCache 3.0 inclut également la fonction d'utilisation optimisée du disque (ODU), atteignant jusqu'à 450,000 XNUMX IOPS selon la société. À mesure que les capacités des SSD continuent d'augmenter, l'utilisation de l'intégralité du disque comme pool de cache devient de moins en moins attrayante ; cependant, avec ODU, les SSD peuvent être partitionnés à la fois en un pool de cache et en un périphérique logique, offrant à l'utilisateur final plus d'options de configuration avec leur stockage flash. Contrairement au pool de cache, la partition logique est exposée au système d'exploitation et peut être utilisée pour y installer un système d'exploitation ainsi que pour stocker d'autres données nécessitant un accès rapide et à faible latence.
Les contrôleurs de la famille Adaptec Series 7 sont livrés avec une garantie de 3 ans.
Toutes les spécifications Adaptec série 7 :
- Modèles
  - 78165 : 24 ports (8 internes et 16 externes)
  - 72405 : 24 ports internes
  - 71685 : 24 ports (16 internes/8 externes)
  - 71605 : 16 ports internes 7805 : 8 ports internes
  - 71605E : 16 ports internes
  - 71605Q : Quatre mini SAS HD internes (SFF-8643)
  - 7805Q : Deux mini SAS HD internes (SFF-8643)
- Débit de 6 Gb/s sur chaque port
- PMC PM8015 ROC
- Performances maximales pour jusqu'à 24 SSD (aucun module d'extension nécessaire)
- Module Flash Adaptec en option (AFM-700) pour une protection du cache sans maintenance
- Nouveaux connecteurs mini SAS HD et câbles SAS HD — pour s'adapter et fonctionner dans des configurations de serveur denses
- Configuration automatique des modes Simple Volume et HBA - "réglez-le et oubliez-le"
- Prise en charge de jusqu'à 256 périphériques SAS et SATA
- Prise en charge de SSD, HDD, chargeur automatique, RBOD, bande
- Certifié Intel EPSD THOL
- Dimensions physiques : 2.535"H x 6.6"L (64 mm x 167 mm) 7805, 71605, 71605E ; 4.198"H x 6.6"L (107mm x 167mm) 72405, 71685
- Température de fonctionnement : 0°C à 55°C* (avec débit d'air de 200 LFM ; sans flash), 0°C à 50°C* (avec débit d'air de 200 LFM ; avec flash)
- Produits PCIe 3.0 désormais conformes à la norme PCI-SIG : ASR-72405 et ASR-71605
- Garantie: ans 3
Construire et concevoir
Les cartes RAID de la série 7 sont disponibles dans des formats demi-longueur et pleine hauteur, mesurant 4.198″ x 6.6″, fournissant 24 ports SAS/SATA natifs.
Tout de suite, vous remarquerez quelque chose de différent sur la carte Adaptec. Au lieu de voir les connecteurs SFF-8087 (mini-SAS) habituels qui sont un incontournable parmi les cartes RAID, vous verrez des connecteurs SFF-8643 (mini-SAS HD, iPASS+ HD). Ces nouveaux connecteurs et câblage haute densité permettent aux cartes Adaptec série 7 d'intégrer 16 ports dans un facteur de forme PCIe demi-longueur à profil bas. C'est une énorme amélioration pour nous chez StorageReview, car la plupart des slots d'extension PCIe dans les serveurs sont LP MD2. Adaptec a décidé assez tôt d'utiliser les câbles SAS HD afin de maximiser le potentiel du PM8015. En raison du profil plus fin du câblage et de sa nature flexible dans les espaces restreints, le routage, la configuration et la gestion des câbles sont plus faciles et bien plus idéaux dans des situations comme la nôtre.
Les quatre connecteurs SFF-8643 sont orientés vers le haut près du support ainsi que deux orientés vers l'arrière. L'adaptateur se connecte via un emplacement PCIe Gen3 x8 ; cependant, il est également rétrocompatible avec PCIe Gen2. L'Adaptec Series 7 prend en charge 1024 Mo de DRAM DDR3-1333 pour la mise en cache, à l'exception du 71605E, qui est proposé en 256 Mo.
Les adaptateurs de la série 7 fonctionnent de 0°C à 50°C sans le module ZMCP, tout en fonctionnant entre 0°C et 50°C avec celui-ci installé. Pour des performances optimales, la carte nécessite un débit d'air de 200 LFM ; si le seuil de chaleur dépasse la limite de température, l'adaptateur étranglera. En ce qui concerne l'alimentation, les adaptateurs de la série 7 consommeront environ 12 à 18 W, selon le nombre de ports.
Au centre des contrôleurs RAID de la série 7 se trouve le processeur PMC8015. Malgré sa taille relativement petite, cette puce a la capacité de se vanter d'une vitesse séquentielle impressionnante de 6,600 450,000 Mo/s et de 8 3 IOPS en vitesse de lecture aléatoire, surpassant facilement la connexion xXNUMX PCIe GenXNUMX.
Performances
Les cartes Adaptec série 7 ont été utilisées dans les appareils de stockage suivants que nous avons précédemment examinés :
Le boîtier Obsidian Series 900D Super Tower est actuellement le modèle le plus haut de gamme de Corsair. Il est équipé de dix emplacements d'extension, jusqu'à quinze baies internes pour disques durs 3.5", prend en charge deux alimentations en bas et offre une large gamme d'autres fonctionnalités pour diverses applications. Ce système est destiné aux passionnés exigeants ; qu'il s'agisse d'un PC de jeu puissant ou d'un serveur de fichiers ultra-performant, le 900D offre une grande flexibilité de configuration.
Inclus dans la série 900D, nous avons installé une carte Adaptec série 16Q à 7 ports (71605Q) utilisée comme stockage de masse accéléré maxCache 3.0. Attachés à la série 7Q, six disques durs Hitachi Ultrastar A7K4000 de 4 To en RAID10 nous donnent un peu plus de 11 To de capacité ainsi que deux SSD Corsair Neutron GTX de 200 Go pour le cache. L'Adaptec 72405 pour l'espace de travail est également équipé, auquel sont connectés vingt-quatre SSD Corsair Neutron GTX 200 Go en RAID0.
