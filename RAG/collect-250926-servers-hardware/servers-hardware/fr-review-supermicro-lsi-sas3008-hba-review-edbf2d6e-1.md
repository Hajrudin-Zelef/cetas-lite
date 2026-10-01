---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-lsi-sas3008-hba-review-edbf2d6e-1
title: "fr-review-supermicro-lsi-sas3008-hba-review-edbf2d6e"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/clean4/fr-review-supermicro-lsi-sas3008-hba-review-edbf2d6e.md
source_anchor: ""
source_lines: [1, 43]
sha256: 3dcc8fb3bb8adeec22546db036ad047d0f4ce0082bc89793da9ece58f7f6a01a
---

# fr-review-supermicro-lsi-sas3008-hba-review-edbf2d6e

Les cartes HBA Supermicro LSI SAS3008 (qui partagent le même contrôleur que les LSI 9300-8i ) sont conçues pour offrir des performances maximales. Avec plus d'un million d'IOPS et un débit supérieur à 6 000 Mo/s, elles répondent aux besoins croissants des entreprises exigeant un débit encore plus élevé pour diverses applications, notamment les bases de données transactionnelles, le Web 2.0, l'exploration de données, ainsi que le streaming et le montage vidéo. Le contrôleur LSI SAS 3008 prend en charge 8 lignes PCIe 3.0 et fournit des liaisons SATA/SAS à des débits compris entre 3 et 12 Gb/s. La version IR prend en charge les modes RAID 0, RAID 1 et RAID 10.
Les nouvelles cartes HBA de Supermicro sont conçues pour exploiter pleinement le potentiel des nouveaux SSD SAS 12 Gb/s déjà disponibles sur le marché, tels que le SSD d'entreprise Toshiba PX02SM et le SSD d'entreprise HGST Ultrastar SSD800MM que nous avons récemment testés dans nos laboratoires. Supermicro intègre ces cartes HBA SAS3 dans ses derniers modèles de serveurs SuperStorage, comme le SuperStorage Server 2027R-AR24NV, qui fait partie de notre laboratoire de test d'entreprise. Alors que les performances des SSD SAS3 étaient bridées par l'infrastructure informatique existante équipée de ports SAS 6 Gb/s, les cartes HBA SAS3008 12 Gb/s de Supermicro lèvent ce goulot d'étranglement qui limitait ces SSD et les futurs SSD compatibles SAS 12 Gb/s.
Les cartes HBA SAS 12 Gb/s LSI SAS3008 (IR) et LSI SAS3008 (IT) de Supermicro sont intégrées aux serveurs de stockage Supermicro, tels que le SuperStorage Server 2027R-AR24NV . Elles ne sont pas disponibles séparément.
Spécifications du HBA Supermicro LSI SAS3008 :
- Des modèles:
  - LSI SAS3008 HBA (IR)
  - HBA LSI SAS3008 (informatique)
- Contrôleur d'E/S : LSI SAS3008/Fusion MPT 2.5
- Connectivité de stockage : 8 ports internes
- Taux de transfert de données : 12 Gb/s conforme à la norme SAS 3.0
- Bande passante SAS : Half Duplex (bus large x4) ; 4800 Mo/s
- Configuration du port : 2 ea, port large x4
- Bus hôte : x8 voies, compatible PCI Express 3.0
- Taux de transfert de rafale de données PCI : Half Duplex x8, PCIe 3.0, 8000 Mo/s
- Connecteurs : Mini SAS HD
- Environnement d'exploitation
  - Température de fonctionnement : 10º à 35º C (32º à 95º F)
  - Température hors fonctionnement : -40º à 60º C (-40º à 140º F)
  - Humidité relative de fonctionnement : 20 % à 95 % (sans condensation)
  - Humidité relative hors fonctionnement : 5 % à 95 % (sans condensation)
Concevoir et construire
Pour ceux qui ont déjà possédé un HBA LSI, il y aura certainement une familiarité avec le HBA Supermicro LSI SAS3008 lors de l'examen de la carte. La carte offre un facteur de forme HHHL PCIe similaire à celui du LSI 9300-8i auquel elle ressemble étroitement, ainsi qu'une disposition de port similaire à l'arrière de la carte. Les changements notables sont cependant une couleur de carte de circuit imprimé différente et une position de connecteur de port légèrement différente.
À l'avant et au centre se trouve un gros dissipateur thermique, qui recouvre le contrôleur d'E/S LSI SAS 3008 basé sur Fusion-MPT. De plus, il existe un connecteur de bord de carte PCIe x8 pour son interface PCIe 3.0.
Les deux modèles du SAS3008 sont équipés de huit voies 12 Gb/s et s'interfacent avec une liaison PCI Express 3.0 à huit voies. Pour une fonctionnalité accrue, ils sont compatibles avec les périphériques SATA 3 Gb/s et 6 Gb/s, ainsi qu'avec les périphériques SAS 3, 6 et 12 Gb/s. L'ajout le plus important aux HBA est les deux connecteurs mini-SAS HD à quatre voies, correspondant à ceux que l'on trouve sur les nouveaux fonds de panier Supermicro.
Performances
Pour évaluer les performances de la nouvelle carte HBA LSI SAS3008 de Supermicro, nous avons utilisé notre nouveau serveur SuperStorage 2027R-AR24NV , équipé de trois de ces cartes connectées directement à 24 baies de disques SAS3. À des fins de comparaison, nous avons testé un échantillon de tous nos nouveaux SSD compatibles SAS3, en mesurant les débits d'E/S séquentiels, la bande passante séquentielle et les débits d'E/S aléatoires de pointe.
- Hitachi Ultrastar SSD800MM
- Toshiba PX02SS
- Toshiba PX02SM
- SSD Seagate 1200
Notre premier benchmark évalue les performances des transferts 4k aléatoires composés d'une activité d'écriture à 100 % et de lecture à 100 %. En débit 4k, le HGST SSD800MM est en tête du classement avec 149,078 66,367 IOP en lecture et 02 117 IOP en écriture, le Toshiba PX42.5SS venant en deuxième position avec 02k IOP en lecture et 1200k IOP en écriture, avec le PXXNUMXSM et le Seagate XNUMX SSD en bas du pack.
En doublant la taille de transfert à 8k et en passant à un transfert séquentiel, l'Ultrastar SSD800MM a continué d'afficher d'excellents chiffres avec 107,331 64,953 IOPS en lecture et 858 519 IOPS en écriture. Ces chiffres traduits en bande passante globale se sont avérés être de XNUMX Mo/s en lecture et de XNUMX Mo/s en écriture de transferts séquentiels en petits blocs.
Le benchmark synthétique final a utilisé une taille de transfert beaucoup plus grande de 128k avec des opérations de lecture à 100% et d'écriture à 100%. Sous cette charge de travail, le HGST Ultrastar SSD800MM a continué à briller à travers le Supermicro LSI SAS3008 HBA avec une vitesse de lecture supérieure à 1 Go/s et une écriture de 571 Mo/s. En descendant dans la liste, il était surprenant de voir que peu d'autres disques SAS3 étaient capables de tirer parti de l'interface plus rapide avec le SSD en régime permanent avec des transferts séquentiels multithreads. Le Toshiba PX02SS est arrivé en deuxième position avec des vitesses de lecture légèrement inférieures à 600 Mo/s, tandis que le PX02SM ne mesurait que 533 Mo/s et que le SSD Seagate 1200 ne mesurait que 340 Mo/s.
Conclusion
Les nouveaux HBA 12 Gb/s de Supermicro basés sur le contrôleur SAS3008 de LSI permettent des performances encore plus élevées pour les environnements virtualisés, les centres de données cloud, le streaming vidéo et d'autres scénarios gourmands en bande passante en prenant en charge plus d'un million d'IOPS et plus de 6000 8 Mo/s par carte. Les deux modèles à 3.0 voies (de PCIe 3008) disponibles sont le LSI SAS3008 HBA (IT) et le LSI SAS3.0 HBA (IR), qui incluent tous deux des connecteurs mini-SAS HD pour répondre à la norme SAS 6. Ces HBA laissent derrière eux le goulot d'étranglement du SAS XNUMX Gb/s pour permettre aux entreprises ayant les exigences de performances les plus robustes d'atteindre leurs objectifs informatiques lorsqu'elles sont associées aux bons SSD.
Avec des performances et une polyvalence comme celle-ci, Supermicro aide à faire sa part pour pousser l'adoption des produits SAS 12 Gb/s sur le marché. Alors que les OEM de serveurs grand public ont besoin de plus de temps de cycle pour adopter une nouvelle technologie, Supermicro est à l'avant-garde en supprimant les contraintes qui empêchent la nouvelle génération de SSD SAS d'atteindre leur plein potentiel. Dans le même temps, la nouvelle marge de sécurité illustre également clairement que ce n'est pas parce qu'un disque revendique la compatibilité SAS 12 Gb/s qu'il peut utiliser pleinement la marge de sécurité dont il dispose. Les données incluses ici ne sont qu'un aperçu. Cependant, des tests de performances détaillés pour chacun des SSD SAS3 seront disponibles sous peu, car le Supermicro LSI SAS3008 assume la responsabilité principale des tests de la dernière génération de SSD d'entreprise à haut débit.
Avantages
- Permet aux SSD SAS3 d'atteindre leur potentiel
- Fondation technologique LSI de confiance
- Large prise en charge des pilotes sur de nombreux systèmes d'exploitation
Inconvénients
- Uniquement disponible à l'achat avec un serveur Supermicro
Conclusion
