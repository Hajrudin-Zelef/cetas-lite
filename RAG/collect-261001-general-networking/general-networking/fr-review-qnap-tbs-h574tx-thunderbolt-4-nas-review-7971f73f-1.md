---
id: collect-261001-general-networking/general-networking/fr-review-qnap-tbs-h574tx-thunderbolt-4-nas-review-7971f73f-1
title: "fr-review-qnap-tbs-h574tx-thunderbolt-4-nas-review-7971f73f"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Samsung"]
dates: []
keywords: ["datacenter", "ethernet", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-qnap-tbs-h574tx-thunderbolt-4-nas-review-7971f73f.md
source_anchor: ""
source_lines: [1, 25]
sha256: fb7b9b2b7aff871a2be7082fa83f2e8b23f24b5d1faa35af4cbe0b319e5d1946
---

# fr-review-qnap-tbs-h574tx-thunderbolt-4-nas-review-7971f73f

Le QNAP TBS-h574TX est une unité de stockage en réseau à 5 baies, équipée de Thunderbolt 4, conçue principalement pour les applications vidéo et multimédia. Il s'agit d'un NAS 1 % Flash doté de la particularité d'utiliser des disques EDSFF (Enterprise and Datacenter Standard Form Factor) E1.S. Ces disques se trouvent principalement dans les centres de données et les applications hyperscale et constituent un choix intéressant pour un NAS d'utilisateur final. Heureusement, si vous n'êtes pas un point de révision du stockage d'entreprise avec un accès facile aux disques E2.S, QNAP inclut les traîneaux M.1 à E5.S avec votre achat. Il y a beaucoup de choses dans cette petite empreinte, alors regardons les spécifications. Cette unité est disponible en deux configurations, notre unité d'examen arborant la variante Intel Core i16 avec XNUMX Go de RAM.
QNAP TBS-h574TX Spécifications
| Processeur | Intel® Core™ i3-1320PE 8C(4P+4E)/12T jusqu'à 4.50 GHz avec carte graphique Intel UHD Intel® Core™ i5-1340PE 12C(4P+8E)/16T jusqu'à 4.50 GHz avec carte graphique Intel Xe | 
| Transcodage accéléré par le matériel | Oui | 
| Mémoire système | 12 Go/16 Go (non extensible) | 
| Baies de disques | 5x E1.S, PCIe Gen3 x2 (échangeables à chaud) | 
| Assistance Drive | E1.S jusqu'à 15mm M.2 2280 via l'adaptateur inclus | 
| Ethernet | 1 port 10GbE 1 port 2.5GbE | 
| Assistance Thunderbolt | 2 ports Thunderbolt 4 | 
| USB | 1x USB 2.0 2x USB 3.2 Gen 2 (Type-A, 10 Gbit/s) | 
| OS | QuTS Hero basé sur ZFS QTS (configurable) | 
| Dimensions (LxHxP) | 2.36 x 8.46 x 7.83 pouces | 
| Consommation d'énergie | 46W typique | 
Permettez-moi de vous expliquer brièvement qui je suis et pourquoi cet appareil s'adresse directement à moi et à mes cas d'utilisation. Le jour, je suis technicien en imagerie numérique et membre de la section locale 600 de l'IATSE. L'une des principales responsabilités d'un DIT est de gérer toutes les données sur le plateau, y compris le téléchargement des images de la caméra, le transcodage des négatifs de l'appareil photo numérique dans un format. pour la post-production et la création de « quotidiens », qui sont des versions plus petites de toutes les séquences de la journée. La nuit, je suis monteur vidéo et le QNAP TBS-h574TX se positionne carrément à l'intersection de ces deux rôles. Ils mentionnent même spécifiquement les DIT et les éditeurs vidéo dans leur documentation marketing.
Sur un plateau de tournage, le débit séquentiel est primordial, les tournages multicaméras pouvant générer des téraoctets de données par jour, avec des heures de rushes à transcoder. Le principal goulot d'étranglement est la vitesse des disques. Bien que la production me fournisse des SSD Samsung T7 , largement assez rapides, je rencontre des ralentissements lors du transcodage et du transfert simultanés des rushes. Un système de stockage 100 % flash est une solution idéale. J'utilise généralement un OWC Express 4M2 en RAID 0 pour un stockage ultra-rapide via une connexion Thunderbolt. Les disques ne sont pas remplaçables à chaud et le RAID logiciel ne convient pas à une utilisation inter-systèmes. Ce NAS vise à résoudre ces problèmes.
Concevoir et construire
Le QNAP TBS-h574TX est une belle machine dans un format raisonnablement petit. QNAP souhaite que vous ayez cet appareil sur votre bureau, et sa conception le reflète. La coque est en aluminium, le logo QNAP sur le dessus est réfléchissant et la marque sur le devant est dorée, ce qui la rend très agréable à regarder. Le panneau avant est facilement amovible et maintenu par des aimants, bien qu'il soit verrouillable, ce qui rend l'accès aux disques très facile. Sa petite taille fait de la portabilité une option facile si vous en avez besoin.
La disposition des ports est bien pensée, avec un port Thunderbolt 4 et un port USB 3 facilement accessibles en façade. Le deuxième port Thunderbolt 4 est situé à l'arrière de la machine, avec les ports réseau 10GbE et 2.5GbE, un port HDMI et deux ports USB-A (2.0 et 3.2). Vous trouverez également le bouton de réinitialisation, la prise cylindrique DC et les trois ventilateurs d'extraction. Le bloc d'alimentation inclus est relativement compact et peut facilement être caché.
Bien que l’apparence et la convivialité de cette machine conviennent parfaitement à la vie sur votre bureau, le son, malheureusement, ne le fait pas. Même lorsque les ventilateurs sont configurés dans leur profil silencieux, le TBS-h574TX est audible et, sous des charges de travail plus lourdes, à la limite du bruit. Le seul point positif est que le profil sonore du ventilateur est proche du bruit blanc et ne présente ni gémissement ni bourdonnement. Si vous travaillez dans un environnement calme, vous le saurez sur votre bureau. Cependant, si vous portez des écouteurs, tout cela est sans objet et ce n'est pas assez fort pour déranger vos voisins si vous travaillez dans un bureau ou si vous apportez cela sur un plateau de tournage.
En fin de compte, c'est à vous de décider de votre tolérance au bruit et de l'environnement dans lequel vous travaillez, mais prévoyez de trouver un emplacement sous votre bureau si vous pensez que cela pourrait devenir gênant. De plus, par défaut, cet appareil émet de nombreux bips (surtout lorsqu'un câble Thunderbolt est branché ou débranché), ce qui est assez agaçant, mais vous pouvez désactiver cette fonction dans les paramètres.
Performances
Comme indiqué, notre modèle de test est équipé d'un processeur Intel Core i5-1340PE, d'une carte graphique Intel Xe et de 16 Go de RAM non extensible. Il utilise le système d'exploitation QuTS Hero basé sur ZFS, mais QTS classique est également disponible. Le NAS dispose de cinq disques SK hynix PE8110 E1.S de 3840 Go configurés en RAID 5.
À des fins de tests, nous avons désactivé la compression et la déduplication, ainsi que QSAL (QNAP SSD Antiwear Leveling), car tous ces éléments affectent négativement les performances. Nos systèmes de test étaient un Mac Mini 2018, un MacBook Pro M2021 Max 1 et deux systèmes Windows. Nous avons utilisé des câbles certifiés Thunderbolt 4 et testé les capacités 10GbE.
J'ai également inclus le QNAP TVS-872XT dans les tests, un autre NAS Thunderbolt que j'utilise quotidiennement depuis plus de trois ans. Tous les tests ont été effectués via SMB. Il est important de noter que, bien que les interfaces des disques soient uniquement PCIe Gen3 x2, chaque disque atteint individuellement 1.5 Go/s lors des tests effectués avec QuTS Hero et, en RAID, devrait facilement saturer deux connexions ou plus simultanément.
Commençons par les performances 10GbE du Blackmagic Disk Speed Test. Sans surprise, le TBS-h574TX sature facilement une connexion 10GbE avec 1,050 868 Mo/s en écriture et 872 Mo/s en lecture. Bien que les performances de lecture soient légèrement inférieures aux attentes, les cas d’utilisation réels ne le remarqueront probablement pas. À titre de référence, le TVS-968XT a la même vitesse d'écriture mais 10 Mo/s en lecture. Ce test a été effectué sur une machine Windows 10 sur un réseau XNUMXGbE (non directement connecté).
