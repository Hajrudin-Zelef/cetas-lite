---
id: collect-261001-general-networking/general-networking/fr-review-graid-supremeraid-sr-1010-review-9cf28dd3-1
title: "fr-review-graid-supremeraid-sr-1010-review-9cf28dd3"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["dram", "gpu", "intel", "nand", "nvidia", "revenue"]
source: docs/RAG/collect-261001-general-networking/fr-review-graid-supremeraid-sr-1010-review-9cf28dd3.md
source_anchor: ""
source_lines: [1, 53]
sha256: 2982560258293147dd57bd5baff0ae2fd8b2926cfbf5a33c6225fde0ffdbf313
---

# fr-review-graid-supremeraid-sr-1010-review-9cf28dd3

L'année dernière, GRAID ™ a fait sensation en lançant une « carte RAID » NVMe. En réalité, il ne s'agissait pas d'une véritable carte RAID, mais d'un GPU NVIDIA associé au logiciel GRAID. Les résultats étaient stupéfiants : leur carte de première génération atteignait près de 9 millions d'IOPS en lecture aléatoire 4K avec une latence infime, et ce, avec seulement 8 SSD.
Maintenant, la société est de retour avec le GRAID SupremeRAID SR-1010 qui fait monter les enchères avec un GPU Gen4 et un logiciel mis à jour qui peut fournir jusqu'à 19 millions d'IOPS en lecture aléatoire 4K et 110 Go/s en lecture séquentielle en gros blocs.
Apprêt GRAID SupremeRAID
Nous avons testé le SR-1000 il y a quelques mois seulement , nous n'allons donc pas nous étendre sur la mission de GRAID. Pour résumer, GRAID a identifié un problème dans la gestion des disques NVMe. Les utilisateurs avaient recours soit à des solutions RAID logicielles pour regrouper leurs SSD, soit à des cartes RAID NVMe, ce qui présente leurs propres difficultés.
Le RAID logiciel est excellent dans la mesure où il peut être très peu coûteux et facile à mettre en œuvre. Mais le RAID logiciel a des pénalités, en particulier, il y a des coups sur les ressources du processeur et du système. Le RAID logiciel n'a pas de composant de gestion matérielle, les tâches doivent donc être déchargées sur le CPU et la DRAM du système.
Les cartes RAID physiques, qui sont la norme depuis des décennies, n'ont pas été en mesure de suivre le rythme de l'innovation en matière de mémoire flash. Un serveur avec un slot PCIe Gen4 ne peut espérer atteindre que 16 Go/s sur un slot x8 ou 32 Go/s sur un slot x16. Le fait est que la plupart des cartes RAID sont x8, donc un système hôte avec seulement 24 SSD aurait besoin de plusieurs cartes RAID pour atteindre un point de pleine performance pour le flash.
Spécifications du GRAID SupremeRAID SR-1010
Cela nous amène à la dernière itération du RAID compatible GPU, le GRAID SupremeRAID SR-1010. Le SR-1010 prend en charge jusqu'à 32 SSD Gen4 dans un seul boîtier. Le matériel, les logiciels et l'interface mis à jour permettent au SR-1010 d'afficher de beaux gains progressifs par rapport à la première carte, notamment en ce qui concerne les écritures. La carte SR-1010 elle-même est une NVIDIA A2000, conçue pour prendre en charge RAID0, 1, 5, 6 et 10 avec le logiciel GRAID.
La nouvelle carte Gen4 offre une puissance de calcul et une bande passante PCIe accrues. La bande passante PCIe contribue aux performances d'écriture puisque les données d'écriture transitent par la carte GPU. C'est là que les performances d'écriture séquentielle et d'écriture aléatoire doublent en raison de l'augmentation de la bande passante de 10 Go/s à 20 Go/s (PCIe Gen3 vs Gen4).
Vous trouverez ci-dessous un aperçu plus complet des spécifications, et la fiche détaillée de GRAID est disponible ici.
- Niveaux RAID pris en charge - RAID 0, 1, 5, 6, 10
- Nombre maximum de disques physiques - 32
- Groupes d'entraînement maximum - 4
- Nombre maximal de lecteurs virtuels par groupe de lecteurs : 8
- Taille maximale du groupe de lecteurs –
- Défini par la taille du disque physique
- Système d'exploitation
  - Linux : Alma Linux 8.5
  - Rocheux Linux 8.5
  - CentOS 7.9, 8.4, 8.5
  - openSUSE saut 15.2, 15.3
  - RHEL7.9, 8.4, 8.5
  - SLES15 SP2, SP3
  - Ubuntu 20.04
  - Windows Serveur 2019 x86-64
  - Windows Serveur 2022 x86-64
- Interface hôte - x16 PCIe Gen 4.0
- Consommation d'énergie maximale - 70W
- Facteur de forme - 2.713″ H x 6.6″ L, emplacement unique
- Poids du produit - 306 g
Performances GRAID SupremeRAID SR-1010
Pour évaluer les performances du SupremeRAID SR-1010, nous avons utilisé notre serveur Dell PowerEdge R750 sous Ubuntu 20.04 équipé de huit baies NVMe Gen4 en façade. Nous avons utilisé huit SSD Memblaze PBlaze6 6926 de 12.8 To, ce qui nous a permis de disposer d'une importante capacité de mémoire NAND haute performance pour une configuration RAID 5.
Étant donné qu'il s'agit de SSD différents de ceux de notre dernier examen, nous avons effectué une série complète de nouveaux tests comparant les configurations RAID5 utilisant un RAID logiciel par rapport aux SR-1000 et SR-1010. Pour le RAID logiciel, nous avons utilisé mdadm avec une taille de bloc de 64 Ko. Des tests ont été effectués en exécutant FIO sur le volume RAID.
Spécifications Dell PowerEdge R750 :
- 2 processeurs évolutifs Intel Xeon de 8380e génération Intel 3
- 32 x 32 Go DDR4 3200 MHz
- 8 x 12.8 To Memblaze PBlaze6 6926
| Performances RAID 5 FIO |  |  |  |  | 
| Test | Logiciel RAID5 | SR-1000 (Gen3) | SR-1010 (Gen4) |  | 
| 1 Mo d'écriture séquentielle (16T/32Q) | 1.3GB / s | 11.1GB / s | 17.7GB / s |  | 
| 1 Mo de lecture séquentielle (16T/32Q) | 56.2GB / s | 49.4GB / s | 49.4GB / s |  | 
| 8K aléatoire 70/30 (32T/64Q) | 160.2k IOPS | 1.51 M IOPS | 1.95 M IOPS |  | 
| Écriture aléatoire 4K (32T/64Q) | 73.9k IOPS | 838k IOPS | 1.56 M IOPS |  | 
| Lecture aléatoire 4K (32T/64Q) | 2.24 M IOPS | 10.3 M IOPS | 11.0 M IOPS |  | 
En commençant par la bande passante en lecture séquentielle, la configuration RAID logicielle avait un léger avantage, mesurant 56.2 Go/s contre 49.4 Go/s sur les SR-1000 et SR-1010.
Passer à l'écriture séquentielle bien que cette piste soit revenue à GRAID avec une énorme différence. Nous avons mesuré 1.3 Go/s avec le RAID logiciel contre 11.1 Go/s avec le Gen3 SR-1000 et 17.7 Go/s avec le Gen4 SR-1010.
Dans notre charge de travail aléatoire 8K avec un mélange 70/30 R/W, nous avons vu 160.2k IOPS à partir du RAID logiciel contre 1.51M IOPS avec le SR-1000 et 1.95M IOPS sur le SR-1010.
En passant à la charge de travail aléatoire 4K, nous avons continué à voir des gains énormes, avec l'écriture aléatoire 4K. La configuration RAID logicielle ne mesurait que 73.9 838 IOPS, contre 1000 1.56 IOPS sur le SR-1010 et XNUMX million d'IOPS bien améliorés sur le SR-XNUMX.
L'écriture aléatoire 4K des deux versions GRAID avait des performances similaires avec 10.3 millions d'IOPS sur le SR-1000 et 11 millions d'IOPS sur le SR-1010, mais étaient bien supérieures au RAID logiciel qui mesurait 2.24 millions d'IOPS.
Conclusion
En testant la GRAID SupremeRAID SR-1010, nous avons obtenu exactement ce que nous attendions. Si la SR-1000 nous avait agréablement surpris l'an dernier, nos attentes étaient encore plus élevées pour la SR-1010. Afin de lui donner toutes les chances de réussir, nous avons associé la carte GRAID à un serveur haut de gamme et à huit de nos meilleurs SSD d'entreprise, les disques Memblaze 6926.
Il convient de noter que nos résultats de performances, bien que très bons, ne correspondront pas à la spécification GRAID simplement en raison du nombre de disques testés. Bien que nous aimerions revoir la carte GRAID avec une configuration 32 SSD entièrement chargée, nous n'avons tout simplement pas autant de SSD correspondants dans le laboratoire.
En termes de gains de performances les plus importants par rapport à la carte de la génération précédente, le GRAID SupremeRAID SR-1010 a présenté les améliorations mesurées les plus importantes dans les charges de travail d'écriture. Nous avons vu les vitesses d'écriture séquentielle passer de 11.1 Go/s avec le SR-1000 à 17.7 Go/s avec le SR-1010. Les performances d'écriture aléatoire se sont également considérablement améliorées, passant de 838 1000 IOPS avec le SR-1.56 à 1010 M IOPS avec le SR-XNUMX. Dans l'ensemble, nous avons été très impressionnés par ces chiffres, qui étaient exponentiellement supérieurs aux performances RAID logicielles.
