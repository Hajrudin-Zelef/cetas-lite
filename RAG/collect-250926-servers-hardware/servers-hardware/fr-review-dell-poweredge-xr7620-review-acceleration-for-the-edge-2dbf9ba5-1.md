---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5-1
title: "fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["gpu", "intel"]
source: docs/RAG/clean4/fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5.md
source_anchor: ""
source_lines: [1, 44]
sha256: e51fa184b058aad73949239bdfe6925ed093696d4c82d3af1731678388dbeb63
---

# fr-review-dell-poweredge-xr7620-review-acceleration-for-the-edge-2dbf9ba5

StorageReview a testé plusieurs serveurs durcis, en se concentrant sur leurs performances, leur capacité et leur facilité d'utilisation. Nous avons récemment testé la solution RAID Cheetah équipée de SSD Solidigm. Malheureusement, pour des raisons évidentes, nous n'avons pas pu tester la résistance de ces serveurs. Cependant, Dell a placé la barre encore plus haut en fournissant un serveur Edge PowerEdge XR7620 et un conteneur de transport robuste, capables de résister à des conditions extrêmes.
La vie est dure au bord
L'infrastructure robuste en périphérie n'est pas un modèle nouveau. Les serveurs de périphérie existent depuis un certain temps déjà ; ils collectent des données et les renvoient au siège social pour en exploiter tout le potentiel. L'intelligence artificielle a mis l'accent sur un traitement plus rapide des données pour une utilisation quasi immédiate, ce qui implique un traitement à la source. L'IA en périphérie a eu un impact considérable, apportant un traitement plus innovant, plus rapide et plus sécurisé à de nombreuses applications concrètes. Un assistant personnel toujours à portée de main, prenant des décisions en temps réel, vous faisant gagner du temps, de l'argent et parfois même des vies !
Configuration Dell PowerEdge XR7620
Dell a présenté ces serveurs PowerEdge spécialement conçus l'année dernière. Le modèle que nous examinons est le Dell PowerEdge XR7620, un serveur robuste à double socket, optimisé pour la périphérie, à courte profondeur, spécialement conçu et compact, offrant des solutions axées sur l'accélération pour la périphérie. La différence entre le XR7620 et les serveurs Edge classiques est que ce serveur répond à la maturation rapide de l'IA/ML avec la prise en charge des charges de travail les plus exigeantes, notamment l'automatisation industrielle, l'analyse vidéo, l'analyse des points de vente, l'inférence IA et l'agrégation de périphériques Edge.
Notre XR7620 est configuré avec 2x Intel Xeon Gold 6426Y, 128 Go DDR5 4800, 480 Go Dell BOSS RAID1 et 4x SSD Solidigm. Nos tests auront lieu dans le laboratoire de recherche environnementale de Jordan au Canada. Nous avons donc acheté un disque de 60 To pour renvoyer les données au laboratoire Storagereview de Cincinnati.
Spécifications du Dell PowerEdge XR7620
| Fonctionnalité | Spécifications techniques | 
|---|---|
| Processeur | Deux processeurs Intel® Xeon® Scalable de 4e génération avec jusqu'à 32 cœurs par processeur | 
| Mémoire | 16 emplacements DIMM DDR5, prend en charge RDIMM 1 To maximum, vitesses jusqu'à 4800 5 MT/s. Prend en charge uniquement les DIMM DDRXNUMX ECC enregistrés | 
| Contrôleurs de stockage |  | 
| Baies de disques | Baies avant : jusqu'à 4 disques SSD SAS/SATA/NVMe de 2.5 pouces, 61.44 To maximum, jusqu'à 8 disques directs NVMe E3.S, 51.2 To maximum | 
| Alimentations |  | 
| Options de refroidissement | refroidissement par air | 
| Ventilateurs | Six ventilateurs de refroidissement échangeables à froid | 
| Dimensions |  | 
| Poids | Maximum 21.16 kg (46.64 livres) | 
| Facteur de forme | Serveur rack 2U | 
| Gestion intégrée | iDRAC9, iDRAC Direct, iDRAC RESTfulAPI avec Redfish, module de service iDRAC | 
| Biseau | Cadre de sécurité en option avec filtre à poussière (capteur de poussière disponible uniquement pour les systèmes de configuration à accès frontal) | 
| Logiciel OpenManage |  | 
| Mobilité | OuvrirGérer Mobile | 
| Intégrations OpenManage |  | 
| Sécurité |  | 
| Options de processeur graphique | Jusqu'à 5 x 75 W (simple largeur, pleine hauteur/demi-longueur, profil bas) GPU ou jusqu'à 2 x 300 W (double largeur, pleine hauteur/pleine longueur) | 
| Carte réseau intégrée | 2 x 1 GbE LOM | 
| Options réseau | 1 x carte OCP 3.0 (en option) | 
| Ports |  | 
| PCIe | Configuration à 2 processeurs : jusqu'à 5 emplacements PCIe (4 x16 Gen4/5, 1 x16 LP Gen4) | 
| Système d'exploitation et hyperviseurs |  | 
Dell PowerEdge XR7620 – Fonctionnalités Edge critiques
Sécurité
Comme tous les autres serveurs PowerEdge, le XR7620 est conçu dans un souci de sécurité à chaque étape du cycle de vie du serveur. Le processus de sécurité commence par une chaîne d'approvisionnement sécurisée avant la construction du serveur. Il s'étend aux serveurs livrés, avec Secure Lifecycle Management et Silicon Root of Trust, puis sécurise ce qui est créé/stocké par le serveur dans Data Protection.
Cette approche de sécurité zéro confiance suppose des autorisations d'accès privilégiées nécessitant une validation à chaque point d'accès et de mise en œuvre avec des fonctionnalités telles que la gestion des accès aux identités (IAM) et l'authentification multifacteur (MFA). Ceci est essentiel pour les déploiements en périphérie où les serveurs sont généralement installés dans des environnements peu sécurisés. Cette politique de sécurité stricte offre la capacité de détecter et de réagir en cas de falsification ou d'intrusion. La plateforme Root of Trust basée sur silicium de Dell crée un environnement sécurisé qui garantit que le micrologiciel provient d'une source fiable. Toute modification détectée dans les versions ou les configurations du micrologiciel peut forcer PowerEdge à verrouiller la configuration et à lancer une restauration vers le dernier bon environnement connu.
Direction
Tous les serveurs Dell PowerEdge suivent une approche de gestion de système standard à trois niveaux qui inclut le contrôleur d'accès à distance Dell intégré (iDRAC), OpenManage Enterprise (OME) et CloudIQ. Cette approche fournit une solution de gestion de système unifiée, simple, automatisée et sécurisée. Il vous permet de gérer un seul serveur à l'aide de la console iDRAC Baseboard Management Controller (BMC), de gérer des milliers de serveurs simultanément avec OME et d'utiliser des informations intelligentes sur l'infrastructure et des analyses prédictives pour maximiser la productivité du serveur avec CloudIQ.
Refroidissement
Nous avons abordé l'importance du refroidissement sur les performances et l'efficacité globales du serveur, en particulier en ce qui concerne l'IA. Ce n’est peut-être pas sexy, mais offrir des performances thermiques optimales est essentiel lors de la conception de serveurs Edge résilients et robustes. Les serveurs PowerEdge XR ont été conçus avec un flux d'air équilibré et efficace en matière de refroidissement et, si nécessaire, une gestion thermique complète qui optimise le flux d'air, régule la vitesse des ventilateurs et réduit la consommation d'énergie.
Ce niveau de détail permet aux serveurs PowerEdge XR de fonctionner entre -5 °C et 55 °C. Dell travaille sur des solutions susceptibles d'étendre encore davantage cette plage opérationnelle.
Tous les serveurs PowerEdge XR sont conçus avec plusieurs doubles ventilateurs contrarotatifs. Cela équivaut à placer deux ventilateurs dans le même boîtier. Il prend en charge la redondance des ventilateurs N+1, permettant au serveur de continuer à fonctionner avec une seule panne de ventilateur à une température maximale de 40 °C pendant au moins quatre heures.
Évolutivité
La capacité du serveur à héberger jusqu'à huit disques NVMe et la prise en charge de capacités d'accélération denses le positionnent bien pour les avancées futures dans les applications de données de pointe et d'inférence d'IA. Sa conception robuste et son système de refroidissement complet signifient également qu'il continuera à fonctionner efficacement à mesure que les charges de travail et les conditions environnementales évoluent.
Charges de travail ciblées
