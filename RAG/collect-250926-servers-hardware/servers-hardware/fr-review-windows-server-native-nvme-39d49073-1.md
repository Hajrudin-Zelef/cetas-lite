---
id: collect-250926-servers-hardware/servers-hardware/fr-review-windows-server-native-nvme-39d49073-1
title: "fr-review-windows-server-native-nvme-39d49073"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd", "benchmarks", "valuation"]
source: docs/RAG/clean4/fr-review-windows-server-native-nvme-39d49073.md
source_anchor: ""
source_lines: [1, 34]
sha256: dd2297b3565309a2f9f36c88e6cc87a37e70095a333ee84599e97b3df5029252
---

# fr-review-windows-server-native-nvme-39d49073

Le 15 décembre 2025, Microsoft a annoncé que Windows Server 2025 intégrerait enfin nativement la norme NVMe à son architecture de stockage. Or, le stockage NVMe est un format répandu sur les serveurs, les stations de travail professionnelles et les PC grand public depuis de nombreuses années, sa compatibilité avec les systèmes d'exploitation étant assurée depuis Windows Server 2012 R2 et Windows 8.1. De ce fait, une annonce concernant la prise en charge « native » de NVMe pourrait paraître anodine, voire anecdotique. Pourtant, il y a plus à découvrir.
Que signifie exactement « NVMe natif » ?
Dans les versions précédentes des piles de stockage Windows (éditions grand public et serveur), les commandes de lecture et d'écriture de données, indépendamment des protocoles matériels sous-jacents, étaient systématiquement traduites en commandes SCSI. La norme SCSI (Small Computer System Interface) remonte au début des années 1980 et a été conçue pour connecter des périphériques et des disques de stockage aux ordinateurs (Storage Networking Industry Association, s.d.). Elle constitue la base de plusieurs protocoles de stockage modernes utilisés pour diverses charges de travail, notamment les protocoles réseau tels que iSCSI (Internet Small Computer System Interface) et FCP (Fibre Channel Protocol), ainsi que les interfaces de stockage local telles que SAS (Serial Attached SCSI) et UASP (USB Attached SCSI).
L'ancienne manière
En convertissant différents protocoles en commandes SCSI, Microsoft a unifié les commandes de stockage aux niveaux supérieurs du système d'exploitation. Toutefois, cette approche a sacrifié une grande partie des gains d'évolutivité et de performance des architectures de stockage modernes. L'ancien cheminement des opérations d'E/S était le suivant :
- Les opérations de lecture et d'écriture ont lieu dans la couche de stockage supérieure, au niveau du système de fichiers.
- Les commandes sont transmises au pilote Disk.sys.
- Disk.sys traduit les commandes de stockage génériques en commandes SCSI.
- Storport reçoit les commandes SCSI et les envoie au pilote Miniport approprié (par exemple, StorAHCI.sys pour les disques SATA).
- Le pilote Miniport correspondant communique directement avec le périphérique de stockage, en le traduisant à nouveau au format de commande de stockage approprié.
D'autres conventions SCSI, telles que les LUN (Logical Unit Numbers) utilisés pour identifier les partitions de données sur un périphérique de stockage, ont été reprises dans la pile de stockage Windows, même si des concepts plus récents comme les espaces de noms NVMe existent depuis un certain temps (Hands, Worley et Lakhveer Kaur, nd).
La nouvelle norme
La nouvelle architecture de stockage de Microsoft pour Windows Server 2025 active de nouvelles fonctionnalités dans Storport et remplace Disk.sys par NVMeDisk.sys, offrant ainsi un cadre évolutif, prêt pour l'avenir et performant.
- Les opérations de lecture et d'écriture ont lieu dans la couche de stockage supérieure, au niveau du système de fichiers.
- Les commandes sont transmises directement de NVMeDisk.sys au nouveau code StorMQ au sein de Storport.
- StorMQ génère les commandes NVMe (ou autre type de stockage) appropriées pour chaque opération de lecture et d'écriture et les envoie directement au matériel lui-même.
(Image tirée de la présentation de Scott Lee à la conférence des développeurs SNIA du 16 septembre 2025)
Cette nouvelle norme pour les opérations de disques sur Windows Server 2025 élimine une couche de traduction et s'intègre pleinement aux files d'attente de commandes de stockage sur les périphériques NVMe, RAID et HBA. La rationalisation du système de stockage Windows offre également des avantages supplémentaires, tels qu'une réduction de l'utilisation des ressources du processeur grâce à la suppression des traductions inutiles des commandes de stockage et une meilleure utilisation des processeurs logiques. La nouvelle architecture adopte d'autres spécifications NVMe, telles que les espaces de noms NVMe et la prise en charge du Plug-and-Play. Elle permet la création de pilotes Miniport de stockage spécifiques au fournisseur ou au type de périphérique et leur intégration à Windows pour une compatibilité et des performances accrues avec les nouvelles générations de périphériques de stockage (Lee, SNIA SDC 2025 – Storage Multi-Queue on Windows, 2025).
Prêt à tester ?
Lors de sa présentation à la conférence SNIA Developer le 16 septembre 2025, Scott Lee a révélé que Microsoft travaillait déjà en étroite collaboration avec les fournisseurs pour développer de nouveaux pilotes pour des périphériques tels que les cartes RAID et les HBA. Cela laisse présager que les améliorations apportées à StorMQ seront bientôt disponibles, voire déjà activées sur de nombreux périphériques de stockage. Bien que la fonctionnalité ait été annoncée comme disponible pour tous en décembre dernier, la nouvelle pile de stockage n'est activée que sur option, nécessitant l'ajout d'une clé de registre. Vous trouverez la procédure d'activation dans l'article de Microsoft consacré à la prise en charge native de NVMe.
Avertissement : Toute modification incorrecte du registre peut entraîner de graves problèmes. Il est donc impératif de tester cette opération au préalable sur un serveur non critique. Plusieurs utilisateurs ayant activé cette fonctionnalité ont signalé des problèmes avec les disques NVMe dont la déduplication est activée. Bien qu'un correctif officiel de Microsoft soit en préparation, vous procédez à vos risques et périls.
Test du NVMe natif sur Windows Server 2025
Notre plateforme de test pour l'évaluation du NVMe natif sur Windows Server 2025 (version 26100.32370) comprenait un serveur à deux sockets SP5 équipé de deux processeurs AMD EPYC 9754 à 128 cœurs. À ces processeurs multicœurs s'ajoutait une mémoire DDR5 de 768 Go fonctionnant à 4 800 MT/s.
Remarque : Selon Yash Shekar de Microsoft, une amélioration intermédiaire sans lien avec le NVMe natif a déjà été publiée pour Windows Server 2025, ce qui pourrait avoir apporté un gain supplémentaire à la pile de stockage non native, réduisant ainsi l’écart potentiel entre les résultats.
Pour évaluer le potentiel de la nouvelle architecture de stockage, nous avons utilisé quinze SSD NVMe Solidigm P5316 de 30.72 To avec PCIe 4.0 en configuration JBOD. Il est important de noter que le Solidigm P5316 possède une unité d'indirection de 64 kilo-octets, ce qui signifie que les performances d'écriture pour les petites tailles (comme les tests 4K) sont souvent inférieures aux attentes. Compte tenu de cette unité d'indirection plus importante, nous avons exécuté des benchmarks FIO avec des tests de lecture et d'écriture pour des blocs aléatoires de 4K, des blocs aléatoires et séquentiels de 64K, et des blocs séquentiels de 128K afin de comparer la vitesse globale pour différentes tailles de blocs. Nous avons également surveillé l'utilisation du processeur pendant les tests afin d'évaluer les affirmations de Microsoft concernant une efficacité accrue.
Points forts
- Augmentation considérable de la bande passante et des IOPS en lecture aléatoire 4K et 64K
- Latence de lecture aléatoire réduite pour les formats 4K et 64K
- Diminution significative de l'utilisation du processeur pour les lectures et écritures séquentielles, quelle que soit la taille des blocs.
| Métrique | 4K aléatoires |  | 64K aléatoires |  | Séquentiel 64K |  | Séquentiel 128K |  | 
|---|---|---|---|---|---|---|---|---|
|  | Non-autochtone | Originaire | Non-autochtone | Originaire | Non-autochtone | Originaire | Non-autochtone | Originaire | 
| Lire |  |  |  |  |  |  |  |  | 
| Bande passante (Gio/s) | 6.1 | 10.058 | 74.291 | 91.165 | 35.596 | 35.623 | 86.791 | 92.562 | 
