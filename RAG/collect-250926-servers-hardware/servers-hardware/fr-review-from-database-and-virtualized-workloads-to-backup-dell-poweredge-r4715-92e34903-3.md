---
id: collect-250926-servers-hardware/servers-hardware/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903-3
title: "fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["acquisition", "open source"]
source: docs/RAG/clean4/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903.md
source_anchor: ""
source_lines: [68, 97]
sha256: a00aadf7ff86e6c6c577c2a28e8558471cf60a2d61c9beadd3eb8d78cc0d16b4
---

# fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903

Le second cas d'usage pratique cible un autre type de charge de travail SMB : le stockage partagé sous Windows. Pour les organisations utilisant des partages de fichiers, des applications départementales ou des serveurs Windows à usage général, ces plateformes répondent aux exigences de performance SMB actuelles. Nous avons exécuté FIO sur Windows Server afin de caractériser les performances séquentielles et aléatoires de deux configurations de stockage : une baie RAID 6 sur disque dur (R5715) comme référence, et une baie JBOD SSD à 8 disques (R4715) représentant une configuration de stockage haute performance. La comparaison met en évidence l'écart de performances entre les deux niveaux de stockage de cette catégorie de plateformes.
L'écart entre les deux configurations de stockage apparaît clairement dans les résultats FIO. Alors que la baie RAID 6 de disques durs du R5715 offrait un débit séquentiel respectable pour une plateforme à disques durs haute capacité, le R4715 équipé de SSD affichait des performances nettement supérieures, notamment pour les charges de travail aléatoires, où les environnements PME sont particulièrement sensibles à la latence de stockage.
Les performances séquentielles sur la baie de disques durs ont atteint jusqu'à 3.7 Go/s en écriture et 2.2 Go/s en lecture lors des tests à 4 cœurs, ce qui est largement suffisant pour les tâches classiques de partage de fichiers, de sauvegarde et de stockage de masse. En revanche, la configuration SSD a permis d'atteindre un débit séquentiel de plusieurs dizaines de gigaoctets par seconde, dépassant les 56 Go/s en lecture et les 26 Go/s en écriture, tout en conservant une latence considérablement réduite.
L'écart s'est encore creusé lors des tests d'écriture aléatoire 4K. La baie de disques durs a plafonné à moins de 1 300 IOPS en écriture aléatoire, avec une latence dépassant les 100 ms, tandis que la configuration SSD a atteint plus de 4 millions d'IOPS avec une latence inférieure à la milliseconde. Concrètement, cela se traduit directement par une meilleure réactivité des applications, des performances accrues pour le partage de fichiers multi-utilisateurs, un comportement optimisé du stockage des machines virtuelles et la capacité à gérer des charges de travail SMB simultanées sans que le stockage ne devienne un goulot d'étranglement.
| Charge de travail FIO | Disque dur RAID6 R5715 1T | Disque dur RAID6 R5715 4T | R4715 8x SSD 1T | R4715 8x SSD 4T | 
|---|---|---|---|---|
| Lecture séquentielle (128 Ko) |  |  |  |  | 
| Bande passante | 1,475.89 Mo / s | 2,198.89 Mo / s | 56,861.09 Mo / s | 56,866.76 Mo / s | 
| IOPS | 11,807 | 17,589 | 454,885 | 454,918 | 
| Latence | 2.70ms | 7.28ms | 0.56ms | 2.25ms | 
| Écriture séquentielle (128 Ko) |  |  |  |  | 
| Bande passante | 2,665.31 Mo / s | 3,726.63 Mo / s | 26,739.52 Mo / s | 26,753.48 Mo / s | 
| IOPS | 21,322 | 29,811 | 213,912 | 214,011 | 
| Latence | 1.49ms | 4.39ms | 1.20ms | 4.78ms | 
| Lecture aléatoire (4K) |  |  |  |  | 
| Bande passante | 1.00 Mo / s | 3.60 Mo / s | 7,268.78 Mo / s | 16,143.19 Mo / s | 
| IOPS | 256 | 919 | 1,860,803 | 4,132,645 | 
| Latence | 125.11ms | 139.00ms | 0.13ms | 0.17ms | 
| Écriture aléatoire (4K) |  |  |  |  | 
| Bande passante | 4.88 Mo / s | 4.67 Mo / s | 7,555.53 Mo / s | 16,010.24 Mo / s | 
| IOPS | 1,248 | 1,195 | 1,934,214 | 4,098,613 | 
| Latence | 25.63ms | 106.98ms | 0.07ms | 0.13ms | 
Serveur de sauvegarde Proxmox
Au-delà des simples performances brutes, le R5715 avec stockage HDD est la plateforme idéale pour une charge de travail de sauvegarde virtualisée. Pour le confirmer, nous avons déployé Proxmox Backup Server sur le R5715 configuré avec le processeur EPYC 9015 à 8 cœurs et la même baie de 12 disques durs 3.5 pouces. Proxmox est un exemple représentatif de l'écosystème open source des hyperviseurs et infrastructures qui a connu un essor considérable auprès des PME, et Proxmox Backup Server, en particulier, est parfaitement adapté au profil de stockage de ce serveur.
Nous avons déployé Proxmox Backup Server 4.2.0 et l'avons utilisé pour sauvegarder les machines virtuelles qui alimentent notre environnement de serveur Discord communautaire Proxmox.
Dans notre configuration, les opérations de sauvegarde et de restauration étaient quelque peu limitées par la connexion réseau 1 GbE du système, qui constituait le principal goulot d'étranglement lors des transferts importants. Cependant, la plateforme prend en charge des mises à niveau réseau simples via des cartes d'extension OCP, facilitant ainsi la migration vers une connectivité 10 GbE, voire 25 GbE. Avec un réseau plus rapide, le R5715 serait capable de gérer un débit de sauvegarde et des performances de restauration nettement supérieurs, notamment dans les environnements comportant des ensembles de données de machines virtuelles plus volumineux ou des fenêtres de sauvegarde plus exigeantes.
Conclusion
Les serveurs Dell PowerEdge R4715 et R5715 doivent leur succès à une configuration parfaitement adaptée. Deux châssis aux formats clairement différenciés, quatre options de processeur couvrant l'ensemble des besoins des PME sans chevauchement, et une gamme de solutions de stockage suffisamment large pour répondre à tous les besoins, du stockage de masse économique aux performances 100 % flash. Cette flexibilité de configuration n'est pas purement théorique. Lors de nos tests, la configuration optimale s'est avérée différente selon les charges de travail. Le processeur 9255 à 24 cœurs a offert le meilleur compromis pour les performances des bases de données transactionnelles sur mémoire flash. Le processeur 9015 à 8 cœurs a fourni la puissance de calcul nécessaire au déploiement d'un serveur de sauvegarde Proxmox avec des disques durs de grande capacité. Le R4715 avec stockage flash était le choix idéal pour le stockage partagé Windows, tandis que le R5715 avec disques durs était le choix optimal pour les charges de travail axées sur la capacité, où les pics d'E/S ne constituent pas un facteur limitant.
Pour les PME et leurs partenaires, la valeur de ces plateformes réside dans la possibilité d'adapter l'infrastructure à la charge de travail, plutôt que d'investir dans une capacité surdimensionnée. L'écosystème Dell, incluant la gestion iDRAC10, ProSupport, la suite de sécurité et la prévisibilité de la chaîne d'approvisionnement, renforce cette valeur au niveau opérationnel. Les équipes informatiques agiles et les partenaires bénéficient tous deux d'une plateforme parfaitement maîtrisée.
Le positionnement de Dell concernant ces serveurs, présentés comme une solution pour consolider les infrastructures existantes et réduire les coûts de licences par socket et par cœur, se confirme au vu des données recueillies. Avec quatre options de processeur allant de 8 à 32 cœurs sur une même plateforme, clients et partenaires bénéficient d'une grande flexibilité pour adapter les coûts d'acquisition, de licences et d'exploitation à leurs besoins réels. C'est là toute la valeur ajoutée, et les serveurs R4715 et R5715 la tiennent pleinement.
