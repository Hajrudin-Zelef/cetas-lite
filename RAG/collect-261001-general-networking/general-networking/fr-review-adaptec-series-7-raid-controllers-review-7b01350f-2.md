---
id: collect-261001-general-networking/general-networking/fr-review-adaptec-series-7-raid-controllers-review-7b01350f-2
title: "fr-review-adaptec-series-7-raid-controllers-review-7b01350f"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-adaptec-series-7-raid-controllers-review-7b01350f.md
source_anchor: ""
source_lines: [37, 51]
sha256: 141154c5ec75d75e527ce2c812cf85cf9bd60c14704835915b7a7a2c17b74d98
---

# fr-review-adaptec-series-7-raid-controllers-review-7b01350f

Les performances du modeste 11 To de stockage de masse ont offert d'excellents résultats, avec des vitesses de lecture et d'écriture de transfert séquentiel mesurant respectivement 858 Mo/s et 458 Mo/s. Avec maxCache 3.0 d'Adaptec, nous avons constaté des augmentations significatives de la lecture et de l'écriture en utilisant deux SSD Corsair comme cache. Transferts aléatoires 4k cadencés à 70,200 20,500 IOPS en lecture et 6.5 3.0 IOPS en écriture. Quant à notre baie SSD beaucoup plus grande, nous avons mesuré plus de 4 Go/s en lecture et 239,000 Go/s en écriture. En se déplaçant vers des transferts aléatoires 4k de pointe, il a montré 932 4 IOPS en écriture 267,000k (soit environ XNUMX Mo/s en accès en écriture aléatoire pur). Les vitesses de lecture aléatoires XNUMXk étaient un peu rapides avec XNUMX XNUMX IOPS.
Le serveur Advatronix Cirrus 1200 est une alternative performante aux serveurs rack traditionnels. Il est équipé d'un processeur Intel Xeon E3, soit le E3-1220LV2, soit le E3-1265LV2, plus puissant (cette version était celle que nous avons testée). Son format compact, de forme cubique bleue, lui confère une taille unique tout en offrant le même accès interne qu'un ordinateur de bureau, facilitant ainsi la maintenance. Ce serveur Advatronix de classe entreprise est doté d'une mémoire vive de 32 Go ECC DDR3 et peut accueillir jusqu'à 12 disques durs 3.5 pouces, répartis dans 12 baies accessibles en façade et remplaçables à chaud. Le Cirrus 1200 intègre un adaptateur RAID Adaptec 71605Q, avec 12 disques durs Hitachi Ultrastar 7K4000 de 4 To et 5 disques Micron P400m de 400 Go.
Les performances de l'Advatronix Cirrus 1200 étaient assez impressionnantes, coûtant beaucoup moins cher que les autres appareils de stockage d'entreprise tout en les surpassant sans sacrifier l'intégrité des données ou la capacité de stockage. Nous avons pu atteindre plus de 100,000 4 IOPS en lecture aléatoire 10K sur iSCSI et saturer 1.36GbE avec des transferts séquentiels en lecture et en écriture cadencés respectivement à 1.42 Go/s et XNUMX Go/s.
Les baies Aberdeen AberNAS N21L sont entièrement conçues avec la technologie SSD et privilégient un débit élevé, offrant un nombre important d'IOPS. Leur capacité de stockage maximale est de 15.3 To SSD dans un format 2U. La personnalisation de l'AberNAS N21L, ainsi que son système d'exploitation Linux, permettent d'unifier les fonctionnalités NAS, iSCSI/IP-SAN et la prise en charge des cibles Fibre Channel au sein d'un seul appareil. Les services informatiques peuvent ainsi choisir entre NAS, iSCSI et Fibre Channel.
Équipé à l'intérieur de l'AberNAS N21L est une série Adaptec 16 ports 7, configurée avec 16 SSD Intel de 800 Go en RAID5 comme livré. La configuration offrait de bonnes performances de lecture, atteignant plus de 100 4 IOPS dans notre charge de travail 2.8K aléatoire et dépassant 128 Go/s en lecture dans notre charge de travail séquentielle de XNUMXK.
Conclusion
La carte Adaptec Série 7 est à la hauteur des attentes et établit une nouvelle norme pour les cartes d'adaptation RAID. Nous l'avons testée en profondeur dans de nombreux scénarios et l'avons intégrée à des solutions de production. Offrant des performances et une efficacité impressionnantes, les cartes Série 7 ont démontré une excellente tenue de route avec des groupes de disques durs traditionnels, des groupes de disques durs accélérés par SSD tirant parti de maxCache 3.0, ainsi que des configurations 100 % SSD pour des E/S et une bande passante maximales. Grâce aux nombreuses configurations disponibles pour la carte Adaptec Série 7, les clients peuvent choisir la configuration de ports adaptée à leurs besoins et l'intégrer à des solutions sans expandeur grâce au nombre accru de ports pris en charge. L'ASR-72405 gère efficacement un grand nombre de SSD et de disques durs, prenant en charge jusqu'à 24 connexions individuelles. Les entreprises qui intègrent ces cartes à leurs serveurs et stations de travail bénéficient ainsi d'une architecture simplifiée et de performances accrues.
Avantages
- Prend en charge jusqu'à 24 appareils sans utiliser d'expandeurs
- Excellente implémentation du logiciel de mise en cache
- Une solution rentable
Inconvénients
- Dans les environnements de système d'exploitation marginaux, la prise en charge des pilotes n'est pas aussi solide
Conclusion
La famille de cartes RAID Adaptec Series 7 offre aux constructeurs de systèmes et aux administrateurs de stockage une flexibilité incroyable avec jusqu'à 24 ports sans avoir besoin d'expandeurs. Pour ceux qui ont besoin d'un peu plus de performances, les cartes 7Q avec maxChache 3.0 offrent de gros gains avec une modeste infusion de flash SSD.
