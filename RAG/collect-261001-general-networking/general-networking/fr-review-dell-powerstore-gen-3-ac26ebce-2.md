---
id: collect-261001-general-networking/general-networking/fr-review-dell-powerstore-gen-3-ac26ebce-2
title: "fr-review-dell-powerstore-gen-3-ac26ebce"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-powerstore-gen-3-ac26ebce.md
source_anchor: ""
source_lines: [31, 58]
sha256: bd6b7145dd41abd5ccfebb297cfd41014828d27b70c041c6371f6ebeb447123e
---

# fr-review-dell-powerstore-gen-3-ac26ebce

Le châssis 3U de base des modèles 5500 et 9500 dispose de 40 emplacements pour disques. Le modèle 1500 est livré avec 24 baies occupées et 16 baies supplémentaires accessibles via une future mise à niveau Data-In-Place. Cela représente 13.3 disques par unité de rack (RU), contre 10.5 pour la génération 2, soit environ 40 % de disques supplémentaires par RU dans l'appliance de base et une densité accrue de 83 % avec l'extension à 44 disques après la sortie de la version RTS. Les disques utilisés sont des SSD NVMe E3.S 1 To standard de différents fournisseurs, sans support propriétaire, ce qui réduit les risques liés aux contraintes d'approvisionnement et à une dépendance à un fournisseur unique. Les capacités vont de 3.84 To, 7.68 To et 15.36 To en TLC, plus 30.72 To en QLC, et le même ensemble d'options est pris en charge sur les trois modèles (1500, 5500 et 9500) sans compromis sur les performances entre les types de supports.
Le choix entre TLC et QLC repose sur le rapport prix/Go et la densité plutôt que sur le niveau de performance. Les configurations minimales varient légèrement selon le modèle : la TLC est disponible à partir de 6 × 3.84 To sur toute la gamme, tandis que la QLC est disponible à partir de 7 × 30.72 To sur le modèle 1500 et de 11 × 30.72 To sur les modèles 5500 et 9500. Tous les disques sont certifiés SED ou FIPS, et la plateforme prend en charge les configurations DRE +1 et +2. La dissipation thermique est également améliorée : Dell annonce des besoins en refroidissement environ 50 % inférieurs à ceux d’un déploiement équivalent de 2.5 pouces, grâce notamment à la géométrie du flux d’air EDSFF. Les 40 baies sont désormais accessibles à l’utilisateur pour les données, le cache étant géré par une mémoire persistante définie par logiciel (SDM) et non plus par des disques NVRAM occupant les baies avant, comme sur la génération 2.
|  | PowerStore Gen 2 | PowerStore Elite (Gen 3) | 
|---|---|---|
| Châssis et plateforme |  |  | 
| Châssis de base | 2U, 2 nœuds | 3U, 2 nœuds | 
| Bus E/S | PCIe génération 3 | PCIe génération 5 | 
| Mémoire | DDR4 | DDR5 | 
| Stockage |  |  | 
| Facteur de forme du lecteur | Jusqu'à 25 SSD NVMe U.2 de 2.5 pouces (double port) | Jusqu'à 40× E3.S/L 1T NVMe (double port) | 
| étagère extensible | Jusqu'à 24 SSD NVMe U.2 de 2.5 pouces | Jusqu'à 44× E3.S/L NVMe (post-RTS) | 
| Stratégie de cache | Disques NVRAM U.2 | Cache vers mémoire flash locale (SDPM) | 
| E/S et réseau |  |  | 
| Emplacements d'E/S | 3× SLIC / OCP 2.0 (PCIe Gen 3 x16) | Jusqu'à 5× OCP 3.0 (PCIe Gen 4/5 x16) | 
| Connexion inter-nœuds | 2× 10 GbE, RDMA | Jusqu'à 200 GbE, RDMA (prêt pour 400 GbE) | 
| Gestion et énergie |  |  | 
| Gestion hors bande (BMC) | EMC GEM | iDRAC (version stockage) | 
| Tuning Moteur | Blocs d'alimentation EMC / alimentations personnalisées | PowerEdge BBU / Alimentations | 
Au sein de la plateforme matérielle PowerStore Gen 3
En comparant la face avant du nouveau serveur 3U PowerStore Gen 3 à celle du modèle Gen 2, on constate que le design modernisé reprend l'esthétique des serveurs PowerEdge de 17e génération que nous avons déjà testés. Il est proposé dans la nouvelle teinte grise de Dell, avec le cadre en nid d'abeille assorti, devenu la signature visuelle de la gamme actuelle de matériel d'entreprise de la marque.
Calcul et mémoire
Côté calcul, le système est alimenté par des processeurs Intel, avec un TDP par processeur allant de 165 W à 270 W et jusqu'à 32 cœurs par processeur sur le modèle 9500. Dell annonce jusqu'à 50 % de cœurs supplémentaires par nœud. La mémoire passe à la DDR5 pour l'ensemble de la gamme : le modèle 9500 est livré avec 2 To (64 modules DIMM de 32 Go), le 5500 avec 1 To et le 1500 avec 512 Go. Le passage à la DDR5 améliore considérablement le débit mémoire, offrant une bande passante deux fois supérieure à celle des générations précédentes de DDR4.
L'architecture interne passe à PCIe Gen 5, offrant une bande passante par voie quatre fois supérieure à celle de l'architecture Gen 3 utilisée dans PowerStore Gen 2. La configuration des sockets est le principal élément de différenciation au sein de la gamme : le modèle 1500 est mono-socket, tandis que les modèles 5500 et 9500 sont bi-sockets. Tous utilisent le même châssis et la même image logicielle PowerStore OS, le nombre de cœurs et la mémoire étant les paramètres permettant d'adapter la configuration.
Refroidissement du châssis
Le système de refroidissement du 9500, illustré ci-dessous, est d'une conception impressionnante. Les deux refroidisseurs de processeur d'un même contrôleur sont reliés par des caloducs et forment un large empilement d'ailettes commun, une solution judicieuse compte tenu de la quantité de chaleur qu'un contrôleur biprocesseur doit dissiper. Le châssis 3U, choix de conception antérieur, joue un rôle crucial dans ce refroidissement. Chaque contrôleur mesure désormais 1.5U de hauteur, offrant ainsi à l'air un parcours plus long qu'avec un plateau 1U. Ceci garantit à chaque modèle une marge de refroidissement confortable tout au long de la durée de vie de la plateforme. Le 1500 utilise un refroidisseur de processeur d'apparence plus classique, mais conserve tous les avantages de flux d'air des modèles 5500 et 9500. Quel que soit le modèle, le refroidissement de la gamme est assuré par six ventilateurs par contrôleur, soit un total de 12 ventilateurs qui maintiennent les disques et le reste de l'unité dans leur plage thermique optimale.
Mise en réseau et stockage
Comparativement aux générations précédentes, les nouveaux modèles adoptent une architecture d'E/S modulaire et standardisée, basée sur des emplacements OCP 3.0. Les modèles 5500 et 9500 offrent jusqu'à 5 emplacements par nœud (dont un réservé à l'extension), tandis que le modèle 1500 en propose 3 (dont un réservé). Les modules sont remplaçables à chaud sans outil et remplacent le support SLIC propriétaire utilisé en deuxième génération. Autre point important, l'interface d'E/S passe à PCIe Gen 5 sur les nouveaux modèles, augmentant considérablement la bande passante par voie disponible pour chaque carte et repoussant les limites du réseau. Au lancement, les options de cartes incluent 4 ports FC 32/64 Gb/s, 4 ports Ethernet 1/10 Gb/s, 4 ports Ethernet 10/25 Gb/s et 2 ports Ethernet 100 Gb/s. L'Ethernet 200/400 Gb/s et le FC 128 Gb/s sont prévus pour des sorties ultérieures. Dell renforce également la sécurité du réseau : toutes les cartes Fibre Channel prendront en charge le protocole EDIF (Encrypted Data-in-Flight) grâce à une prochaine mise à jour logicielle non intrusive. Il en résulte jusqu’à 40 ports réseau par appareil, soit le double de la génération précédente et une densité de ports supérieure d’environ 11 % à celle du contrôleur de deuxième génération.
La disposition des contrôleurs a également été repensée. Sur la génération 2, le contrôleur supérieur était inversé par rapport à celui de l'inférieur, ce qui compliquait la maintenance et augmentait le risque de retirer le mauvais contrôleur ou de manipuler le mauvais composant lors d'un remplacement à chaud. Sur la génération 3, les deux contrôleurs sont orientés de la même manière et peuvent être retirés indifféremment à l'aide des deux poignées et leviers situés sur les côtés du châssis, en les faisant glisser comme un plateau 1.5U. Ce changement, bien que mineur en apparence, est significatif pour la facilité de maintenance, notamment dans les baies où l'accès par le haut est limité ou lorsque le technicien travaille sous pression.
