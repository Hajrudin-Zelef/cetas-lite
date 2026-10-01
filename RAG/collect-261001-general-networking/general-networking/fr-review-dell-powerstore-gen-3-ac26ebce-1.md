---
id: collect-261001-general-networking/general-networking/fr-review-dell-powerstore-gen-3-ac26ebce-1
title: "fr-review-dell-powerstore-gen-3-ac26ebce"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["dram", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-powerstore-gen-3-ac26ebce.md
source_anchor: ""
source_lines: [1, 30]
sha256: 0823dd3ff67eec3419355c8f96aceb3e852c3eea080333cd79d98cade2e69959
---

# fr-review-dell-powerstore-gen-3-ac26ebce

Les mises à jour de stockage se présentent généralement sous deux formes. Il y a la mise à niveau discrète, où un fournisseur intègre un nouveau processeur, annonce un gain de performance de quelques points de pourcentage et livre le même châssis avec une étiquette différente. Et puis il y a la refonte générationnelle, où le châssis, les disques, l'interconnexion, l'architecture du cache et le plan de gestion sont tous remplacés simultanément. Dell PowerStore Gen 3, initialement commercialisé sous le nom de PowerStore Elite, appartient à la seconde catégorie, et de loin. Chaque sous-système majeur de la plateforme a été modifié, et la plupart d'entre eux l'ont été de manière à redéfinir en profondeur ce à quoi une baie unifiée devrait ressembler en 2026.
Nous suivons PowerStore depuis ses débuts. À notre avis, il s'agit de la mise à jour la plus importante depuis son lancement initial en 2020, et sans doute de la refonte de plateforme de stockage la plus ambitieuse proposée par un grand constructeur depuis des années. Dell n'a pas simplement amélioré la Gen 2. L'entreprise a entièrement reconstruit la plateforme, du châssis jusqu'au boîtier, en misant sur des formats et des choix architecturaux que la plupart des acteurs du secteur n'ont pas encore adoptés, et a conçu le tout pour une durée de vie de dix ans avec de multiples mises à niveau de contrôleurs. Pour découvrir ces nouveaux modèles, Dell nous a invités à Hopkinton, dans le Massachusetts.
La densité de stockage est un argument de vente majeur du nouveau Dell PowerStore, et Dell ne déçoit pas sur ce point. La plateforme prend en charge jusqu'à 40 disques NVMe E3.S dans un châssis 3U, avec une compatibilité prévue pour les disques E3.L, tout en conservant des emplacements adressables par l'utilisateur pour les données, au lieu de réserver des emplacements pour les SSD de cache. Dell a également modernisé en profondeur la plateforme matérielle sous-jacente, en adoptant des processeurs Intel de nouvelle génération, de la mémoire DDR5, une connectivité PCIe Gen5 de bout en bout et des modules OCP 3.0 qui remplacent l'ancienne conception de carte SLIC spécifique à Dell.
La connectivité entre les contrôleurs atteint aujourd'hui 200 GbE RDMA, avec la possibilité d'atteindre des débits encore plus élevés grâce à une future mise à niveau des cartes d'E/S. Côté logiciel, PowerStoreOS 5.0 introduit une nouvelle intelligence autonome pour la gestion des chemins de données et des métadonnées structurées en journaux afin d'optimiser les performances et l'endurance des mémoires flash QLC haute capacité. La télémétrie au niveau des E/S prépare le terrain pour la détection intégrée des ransomwares, et le partage dynamique des ressources entre les services de blocs et de fichiers est également intégré. Tous les équipements PowerStore sont désormais unifiés d'emblée, offrant une prise en charge améliorée de l'extension horizontale des fichiers et des blocs, ainsi qu'une mobilité des données non perturbatrice entre les clusters.
Dell a également ajouté la déduplication non alignée et des déchargements de compression améliorés, un facteur clé de la décision de l'entreprise d'augmenter sa garantie de réduction des données de 5:1 dans les générations précédentes à 6:1 avec la nouvelle plateforme.
Les modifications matérielles et les mises à jour logicielles ne suffisent pas à expliquer l'ensemble des changements. Les choix architecturaux sous-jacents de Dell sont essentiels, car ils déterminent si la plateforme restera performante au cours de la prochaine décennie ou si elle paraîtra obsolète dans quelques années. Le passage aux disques E3.S/L, le châssis plus grand, l'adoption de la mémoire persistante définie par logiciel, l'interconnexion sans fil du fond de panier et le découplage de l'architecture inter-nœuds par rapport à la génération de processeurs sont autant de décisions tournées vers l'avenir. Ensemble, elles font du châssis de 3e génération une base solide pour la stratégie de mise à niveau multigénérationnelle proposée par Dell avec l'extension du cycle de vie.
La gamme PowerStore s'enrichit de trois nouveaux modèles. Le PowerStore 1500 est une plateforme mono-processeur dotée de 24 baies de disques dès son lancement et d'une interface inter-nœuds RDMA 100 GbE. Les modèles 5500 et 9500, au format 3U, sont bi-processeurs et offrent 40 baies de disques ainsi qu'une connectivité inter-nœuds RDMA 200 GbE. Le 9500 propose deux fois plus de mémoire et un nombre de cœurs supérieur au 5500. Une future mise à niveau via le remplacement du contrôleur permettra au 1500 d'évoluer jusqu'à 40 disques et une connectivité RDMA 200 GbE. Ces trois modèles fonctionnent sous PowerStoreOS 5.0, partagent la même architecture d'E/S OCP 3.0 et prennent en charge les supports TLC et QLC, sans perte de performance lors du passage à la technologie QLC. Il n'est donc plus nécessaire de choisir entre différents modèles pour répondre à divers besoins ou exigences de performance.
Les plateformes de deuxième génération restent disponibles et les clients PowerStore existants bénéficient d'une solution d'avenir claire grâce au clustering intelligent. Les appliances de première, deuxième et troisième génération peuvent coexister au sein d'un même cluster, assurant une mobilité des charges de travail sans interruption de service. Cette solution est idéale pour une base installée qui ne souhaite pas changer de plateforme tous les deux ou trois ans.
Spécifications du Dell PowerStore Gen3
Les trois modèles de troisième génération partagent un châssis 3U à double nœud, la même architecture d'E/S OCP 3.0 et un système d'exploitation PowerStoreOS unifié prenant en charge nativement le stockage par blocs ou par fichiers. Ils diffèrent par le nombre de sockets du processeur, la capacité de la mémoire DRAM, le nombre de disques et la bande passante inter-nœuds.
| Spécifications | Power Store 1500 | Power Store 5500 | Power Store 9500 | 
|---|---|---|---|
| Marché |  |  |  | 
| placement | Milieu | milieu/haut de gamme | Modèle phare haut de gamme | 
| Châssis | 3U, double nœud |  |  | 
| Calcul et mémoire |  |  |  | 
| Plateforme CPU | Intel mono-socket | Intel double socket | Intel double socket | 
| Processeur par appareil | 2 cœurs 24 GHz | 4 cœurs 24 GHz | 4 cœurs 32 GHz | 
| Mémoire par appareil | 512 Go (16 × 32 Go) | 1,024 Go (32 × 32 Go) | 2,048 Go (64 × 32 Go) | 
| Stockage |  |  |  | 
| Lecteurs par appareil de base | Jusqu'à 24 EDSFF | Jusqu'à 40 EDSFF | Jusqu'à 40 EDSFF | 
| Disques par extension (après RTS) | 44EDSFF |  |  | 
| Assistance routière | TLC : 3.84 / 7.68 / 15.36 To · QLC : 30.72 To |  |  | 
| Configuration minimale du disque dur | TLC 6 × 3.84 To · QLC 7 × 30.72 To | TLC 6 × 3.84 To · QLC 11 × 30.72 To | TLC 6 × 3.84 To · QLC 11 × 30.72 To | 
| Capacité brute maximale (base) | ~737 To (24 × 30.72 To) | ~1.2 PB (40 × 30.72 TB) | ~1.2 PB (40 × 30.72 TB) | 
| E/S et réseau |  |  |  | 
| cartes de ligne OCP 3.0 par nœud | 3 (+1 réservé) | 5 (+1 réservé) | 5 (+1 réservé) | 
| Interconnexion entre nœuds | RDMA 100 GbE | RDMA 200 GbE | RDMA 200 GbE | 
Dell PowerStore Gen 2 vs. Gen 3
Chaque sous-système matériel majeur des appliances PowerStore de la série Gen 3 a bénéficié d'une évolution d'au moins une génération. Les changements les plus importants pour la planification de la capacité sont le châssis 3U, les baies de disques E3.S et le passage de 2 × 10 GbE à jusqu'à 200 GbE sur l'interconnexion des nœuds.
