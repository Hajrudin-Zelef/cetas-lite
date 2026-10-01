---
id: collect-261001-ia-llm/ia-llm/fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64-3
title: "fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64.md
source_anchor: ""
source_lines: [73, 92]
sha256: 0173933d3aae17a1bc9bffbad3a75e88ac877eef6448d5e7b0de1a105e38e0e0
---

# fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64

Dans Paramètres > Sauvegardes , l'UNAS permet de sauvegarder les disques des utilisateurs sur des destinations distantes, notamment des serveurs tels que d'autres UNAS Pro distants, des serveurs CIFS/SMB ou Google Drive. Les utilisateurs peuvent appliquer des règles, comme le remplacement de dossiers par les données sources ou la mise à jour de dossiers pour refléter les modifications apportées à des fichiers supprimés. Les intervalles de sauvegarde peuvent être programmés (quotidiens, hebdomadaires ou mensuels), offrant ainsi aux utilisateurs une grande flexibilité quant à la fréquence de sauvegarde de leurs données. Ceci garantit la sécurité et la mise à jour des données, quel que soit leur emplacement. Notez que les prochaines versions d'UniFi Drive 3.0 intégreront les services S3, Backblaze et Wasabi.
Performances de l'UNAS
Pour évaluer les performances du système UNAS, nous avons utilisé CrystalDiskMark et l'avons testé sous trois configurations RAID : RAID 5, RAID 6 et RAID 10. Le système UNAS était partagé avec un poste de travail de notre environnement de laboratoire via le réseau via son port 10 GbE, où chaque configuration a été soumise à des tests pour mesurer ses capacités d'E/S séquentielles et aléatoires. Le test a été réalisé en cinq passes et avec une taille de test de 64 Gio, avec une charge de travail en lecture/écriture de 70 % en lecture et 30 % en écriture.
- RAID10 Il offre les performances globales les plus équilibrées, alliant des vitesses séquentielles élevées et un débit constant pour les charges de travail mixtes. Avec un débit de lecture et d'écriture élevé et des IOPS respectables, il est idéal pour les environnements exigeant vitesse et redondance, comme les charges de travail virtualisées ou les partages de fichiers actifs.
- RAID6 Tolérance aux pannes prioritaire avec protection à double parité, tout en offrant des performances d'écriture compétitives. Il affiche la latence la plus faible et les IOPS d'écriture aléatoire les plus élevées du groupe, ce qui en fait un choix judicieux pour les systèmes critiques ou à forte charge d'écriture, où la disponibilité et la protection des données sont essentielles.
- RAID5 Il a dominé les vitesses d'écriture séquentielle et les IOPS en lecture aléatoire, affichant d'excellentes performances en lecture. Cependant, il présentait la latence la plus élevée lors des écritures aléatoires en raison de la surcharge de parité, ce qui le rend plus adapté au stockage à faible capacité, où les performances en lecture sont prioritaires sur la réactivité en écriture.
| Métrique | UNAS RAID10 | UNAS RAID6 | UNAS RAID5 | 
| Lecture séquentielle (Q8T1) | 799 Mo / s | 535 Mo / s | 704 Mo / s | 
| Écriture séquentielle (Q8T1) | 675 Mo / s | 648 Mo / s | 691 Mo / s | 
| Lecture aléatoire 4K (Q32T1) | 36.76 Mo / s | 35.10 Mo / s | 38.41 Mo / s | 
| Écriture aléatoire 4K (Q32T1) | 4.18 Mo / s | 4.99 Mo / s | 4.00 Mo / s | 
| Lecture aléatoire des IOPS | 8,974 | 8,570 | 9,377 | 
| Écrire au hasard IOPS | 1,020 | 1,219 | 977 | 
| Latence d'écriture aléatoire (µs) | 30,747 μs | 23,717 μs | 32,150 μs | 
| Charge de travail mixte (70R/30W) | 672.51 Mo / s | 424.00 Mo / s | 674.66 Mo / s | 
Conclusion
En résumé, l'Ubiquiti UNAS Pro, proposé à 499 $, est une option abordable par rapport à de nombreuses autres solutions NAS grand public de même taille ou plus petites sur le marché. Ce prix, combiné à ses fonctionnalités robustes et à sa capacité de 7 baies de disques 2.5/3.5 pouces, en fait un choix attractif pour les particuliers et les professionnels à la recherche d'une solution de stockage de taille moyenne et économique.
L'UNAS Pro prend en charge différentes configurations RAID, notamment RAID 5, 6 et 10. Il propose désormais des groupes RAID flexibles, permettant aux utilisateurs de combiner des tailles de disques et de créer plusieurs pools de stockage indépendants adaptés à leurs besoins. L'UNAS Pro offre également un accès à distance et un partage de fichiers simplifié, pour un accès pratique à vos fichiers où que vous soyez. L'application UniFi Drive offre une interface intuitive pour gérer votre stockage, et l'appareil prend en charge les connectivités 1 GbE et 10 GbE pour des transferts de données flexibles et à haut débit.
Cependant, l'écosystème limité d'applications tierces constitue un inconvénient majeur. Contrairement à d'autres systèmes NAS qui offrent un large éventail d'applications tierces, l'UNAS Pro est actuellement dépourvu de nombreuses applications NAS classiques. Cela peut limiter ses fonctionnalités pour les utilisateurs qui dépendent d'options logicielles spécifiques pour leurs flux de travail. Cela dit, la solide expérience d'Ubiquiti en matière d'expansion et d'amélioration de son offre de produits au fil du temps suggère que l'écosystème d'applications de l'UNAS Pro pourrait se développer, ce qui pourrait permettre de remédier à cette limitation lors de futures mises à jour.
Page produit (lien affilié)
