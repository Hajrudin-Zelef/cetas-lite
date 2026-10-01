---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8-1
title: "fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["dram", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8.md
source_anchor: ""
source_lines: [1, 57]
sha256: 6207784a419d6131011f478900efad1ca9d66260d18232167589b379a748bdf4
---

# fr-review-hpe-proliant-dl560-gen11-review-closed-loop-liquid-cooling-664217b8

Si vous avez déjà envisagé un serveur 2U standard à double socket et souhaité deux fois plus de processeurs et une empreinte RAM massive, le HPE ProLiant Dl560 Gen11 est là pour vous. Ce serveur prend en charge quatre processeurs évolutifs Intel Xeon de 4e génération, jusqu'à 16 To de DRAM, ainsi qu'une grande variété de stockage et d'extension de baie arrière. Qu'est-ce qui est cool d'autre ? Pour les configurations à quatre processeurs avec un TDP de 270 W ou plus (par puce), HPE fournit un refroidissement liquide en boucle fermée pour maintenir ces cœurs en fonctionnement à des performances optimales.
Oui. Vous avez bien lu. Ce système est livré avec un système de refroidissement liquide en boucle fermée pour la configuration quad-CPU ! À quel point cela est cool? (jeu de mots volontaire)
Mais sérieusement, HPE a brisé le moule avec cette configuration et a permis de regrouper toute la puissance et la flexibilité dans un châssis 2U standard. Ce qui le rend si unique, cependant, c'est la configuration empilée des cartes processeurs qui permet aux administrateurs système d'accéder relativement facilement à l'une des quatre puces et la possibilité d'intégrer les systèmes de refroidissement liquide dans un seul package complet sans avoir à se soucier de une vidange ou une fuite sur tout l'équipement.
Matériel HPE ProLiant DL560 Gen11
Une fois le capot retiré, vous pouvez voir les deux processeurs supérieurs et leurs plaques de refroidissement, ainsi que la gamme d'emplacements RAM disponibles, qui dans notre modèle de test est presque vide ; cependant, pour garantir une bonne circulation de l’air, il est livré avec des bouches d’espacement.
Le système en boucle fermée passe du radiateur au premier processeur en passant par le deuxième processeur, puis remonte au radiateur pour refroidir le liquide maintenant chauffé. En dessous se trouve une image miroir de ce que vous voyez en haut.
Une chose à prendre en compte est que le système de refroidissement liquide est couvert par une garantie de 5 ans, après quoi le système doit être remplacé (dans son intégralité) aux frais du client. Aucune partie du système de refroidissement ne peut être réparée ou remplacée par le client et doit être effectuée par du personnel certifié HPE pendant la période de garantie de 5 ans.
En regardant de plus près les plaques de refroidissement, et comme la plupart des refroidisseurs de processeur pour systèmes hautes performances, l'accès au processeur en dessous est simple, comme toujours.
Un regard rapproché et personnel sur la configuration de stockage qui nous a été fournie. Comme la plupart des systèmes HPE sortant de leurs usines, le DL560 Gen11 est très polyvalent et peut s'adapter à plusieurs contrôleurs et voies de stockage différents.
En fonction de la configuration du stockage, le fond de panier aura différentes combinaisons d'options. Notre système de test n'a qu'une seule « case » remplie, mais HPE propose plusieurs configurations de stockage, notamment des fonds de panier SSD E3.S.
L'arrière de cette machine dispose de suffisamment d'espace pour des composants supplémentaires tels que des GPU (jusqu'à 6 !) ou des disques de démarrage NVMe si c'est votre truc. La mise en réseau et la connectivité sont facilement accessibles et configurées à l'arrière.
De plus, même si notre modèle de test n'en avait pas besoin, vous pouvez voir où il y a de la place pour quatre modules d'alimentation pour alimenter cette bête.
Gestion HPE ProLiant DL560 Gen11
Grâce à iLO, vous pouvez consulter des informations générales telles que les configurations du processeur, de la mémoire, du réseau et du stockage, ainsi que l'état de chacun.
Dans l'onglet Alimentation et thermique, nous pouvons voir la lecture de la puissance et l'état de l'alimentation sur la configuration du serveur. Des informations détaillées sur la puissance sont disponibles dans l’onglet du compteur de puissance. Cependant, cela nécessitait une licence iLO payante, que ce système n'incluait pas.
À côté de l'alimentation se trouvent les ventilateurs et les modules de refroidissement. C'est ici que vous pouvez surveiller le système de refroidissement liquide. Malheureusement, il n'affiche que l'état général, si les pompes fonctionnent et à quelle vitesse.
iLO contient un onglet d'informations sur la température sous alimentation et thermique. Basées sur chaque capteur de température sur un graphique 3D, ces informations vous montrent où se trouvent vos points chauds dans le système. Les données de température sont essentielles à surveiller pour garantir la stabilité et la durée de vie de votre matériel.
Spécifications du HPE ProLiant DL560 Gen11
La gamme potentielle de spécifications du DL560 Gen11 comprend :
| Processeur | Processeurs Intel® Xeon® Scalable de 4e génération prenant en charge jusqu'à 60 cœurs | 
| Mémoire |  | 
| Contrôleurs de stockage |  | 
| Baies de disques | 8, 16 ou 24 SFF SAS/SATA/NVMe | 
| Alimentations |  | 
| Ventilateurs | Solution de refroidissement liquide, Ventilateurs redondants enfichables à chaud, Kit de ventilateurs performants ou Kit de ventilateurs hautes performances | 
| Dimensions | 3.4 cm x 17.05 cm x 31.75cm | 
| Facteur de forme | Châssis rackable 2U | 
| Gestion intégrée |  | 
| Utilitaires du serveur |  | 
| Sécurité |  | 
| Port réseau de gestion à distance HPE iLO | 1 Go dédié, arrière | 
| Options réseau | Aucun standard. Un choix de carte réseau OCP ou de carte réseau stand-up est requis. Les modèles BTO seront présélectionnés avec une carte réseau principale. | 
| Options GPU | Jusqu'à 6 petits GPU ou 2 grands | 
| USB | Jusqu'à 7 au total :  | 
| Ports supplémentaires |  | 
| PCIe |  | 
| Système d'exploitation et hyperviseurs |  | 
Performances
Configuration HPE ProLiant DL560 Gen11
- 4 processeurs Intel Xeon Platinum 8444H (16 cœurs, 2.9 GHz)
- 512 Go DDR5 (8 x 64 Go DDR5-4800)
- Windows Server 2022
Nous le comparons au DL320 , récemment testé, à titre de comparaison.
Test de vitesse Blackmagic RAW
Nous exécutons toujours le test de vitesse RAW de Blackmagic, qui teste la lecture vidéo. C'est un bon indicateur. Il s’agit plutôt d’un test hybride incluant les performances du CPU et du GPU pour le décodage RAW réel. Bien qu'il n'y ait pas d'autres machines HPE pour montrer une comparaison directe ici, pour vous donner une idée de la situation de celle-ci, les autres serveurs ne produisent généralement qu'à 50-60 ips.
| Test de vitesse Blackmagic RAW (Plus c'est haut, mieux c'est) | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cœurs, 2.9 GHz) | HPE ProLiant DL320 (processeur Intel Xeon-G 4 de 6430e génération, 32 cœurs, 2.1 GHz) | 
| CPU 8K | 86 | 99 | 
| CUDA 8K | N/D | N/D | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs.
| Cinebench R23 | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cœurs, 2.9 GHz) | HPE ProLiant DL320 (processeur Intel Xeon-G 4 de 6430e génération, 32 cœurs, 2.1 GHz) | 
| Processeur (multicœur) (points) | 83,478 | 38,707 | 
| Processeur (monocœur) (points) | 1,261 | 1,245 | 
| Rapport PM | 66.21x | 31.10x | 
Cinebench 2024
Cinebench 2024 de Maxon est une référence de rendu CPU et GPU qui utilise tous les cœurs et threads du processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Comme cette configuration n’a pas de GPU, nous n’avons pas ces chiffres. Des scores plus élevés sont meilleurs.
| Cinebench 2024 | HPE ProLiant DL560 Gen11 (Quad Intel Xeon Platinum 8444H, 64 cœurs, 2.9 GHz) | HPE ProLiant DL320 (processeur Intel Xeon-G 4 de 6430e génération, 32 cœurs, 2.1 GHz) | 
