---
id: collect-250926-servers-hardware/servers-hardware/fr-review-linux-vs-windows-server-2025-native-nvme-926200e2-1
title: "fr-review-linux-vs-windows-server-2025-native-nvme-926200e2"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd", "distribution", "exploit"]
source: docs/RAG/clean4/fr-review-linux-vs-windows-server-2025-native-nvme-926200e2.md
source_anchor: ""
source_lines: [1, 49]
sha256: 4f7857634fe430ead5eabbe5448a048d43b6e102e2f755d523ffc648f4cdc28e
---

# fr-review-linux-vs-windows-server-2025-native-nvme-926200e2

Suite à la publication de notre article sur la fonctionnalité NVMe native optionnelle de Microsoft sur Windows Server 2025 , nous avons reçu de nombreuses demandes pour une comparaison directe des performances de stockage entre Windows Server 2025 avec NVMe natif et un système d'exploitation serveur basé sur Linux. Un utilisateur de Reddit particulièrement enthousiaste nous a même offert de la bière pour que nous réalisions ce test ! Faute d'autres options raisonnables, nous avons décidé d'effectuer les mêmes tests sous Linux.
Il y a très longtemps, dans une version d'OS très, très lointaine
Le noyau Linux prend en charge NVMe depuis la version 3.3 , sortie en mars 2012. De même, le protocole est pris en charge sur Windows Server (de manière non native, via une traduction SCSI) depuis la version 2012 R2, soit environ octobre 2013. Plus d'une décennie plus tard, les utilisateurs débattent encore pour savoir si Windows ou Linux est le meilleur système d'exploitation pour le stockage. Nous avons donc pensé alimenter encore davantage le débat avec des résultats de tests comparatifs entre les deux.
Disposant de résultats de tests pour Windows Server 2025 avec des architectures de stockage natives et non natives, il nous a semblé pertinent d'évaluer deux architectures de stockage sous Linux. Pour nos tests FIO, nous avons utilisé libaio et io_uring, deux des API les plus populaires pour les transactions de stockage. Si io_uring est nettement plus récente et offre de nombreuses améliorations pour les E/S asynchrones, libaio reste largement utilisée pour sa flexibilité et sa simplicité d'utilisation (Didona, Pfefferle, Ioannou, Metzler et Trivedi, 2022). Une description architecturale complète de ces deux architectures dépasse le cadre de cet article, mais nous présentons néanmoins des résultats permettant une comparaison directe.
Tests NVMe sur Ubuntu Server 24.04.4 LTS
Pour cette comparaison, notre plateforme matérielle est le même serveur que celui utilisé dans notre article sur Windows Server 2025 et la prise en charge native du NVMe . Afin de garantir un débit maximal et des résultats constants, il est équipé de deux processeurs AMD EPYC 9754 à 128 cœurs, de 768 Go de mémoire DDR5 à 4 800 MT/s et de quinze SSD NVMe Solidigm P5316 de 30.72 To avec PCIe 4.0 en configuration JBOD.
Comme indiqué dans notre précédent article, le Solidigm P5316 possède une unité d'indirection de 64 kilo-octets, ce qui signifie que les performances d'écriture pour les petites tailles (comme les tests de 4 Ko) sont souvent inférieures aux attentes. Nous avons donc exécuté différents tests avec des tailles de blocs de 4 Ko, 64 Ko et 128 Ko afin d'obtenir un large éventail de résultats pour les opérations de lecture et d'écriture.
Nous avons choisi Ubuntu Server 24.04.4 LTS comme exemple de distribution Linux en raison de sa popularité et de son support à long terme. Il utilise par défaut le noyau Linux 6.8, qui n'est ni le plus récent ni le plus avancé, mais qui représente probablement une part importante des installations dans le monde.
Points forts
- Windows Server 2025 avec NVMe natif remporte trois des quatre tests de performances en lecture.
- Utilisation du processeur plus faible observée avec Windows Server lors de la plupart des tests
- Ubuntu Server 24.04.4 LTS remporte trois des quatre tests de performances en écriture.
| Métrique | 4K aléatoires |  |  |  | 64K aléatoires |  |  |  | 
|---|---|---|---|---|---|---|---|---|
|  | Windows non natif | Windows natif | Linux libaio | Linux io_uring | Windows non natif | Windows natif | Linux libaio | Linux io_uring | 
| Lecture aléatoire |  |  |  |  |  |  |  |  | 
| Bande passante (Gio/s) | 6.1 | 10.058 | 9.198 | 9.504 | 74.291 | 91.165 | 77.517 | 77.7 | 
| IOPS | 1,598,959 | 2,636,516 | 2,411,000 | 2,491,000 | 1,217,176 | 1,493,637 | 1,270,000 | 1,273,000 | 
| Latence moyenne (ms) | 0.169 | 0.104 | 0.198 | 0.192 | 0.239 | 0.207 | 0.377 | 0.376 | 
| Utilisation totale du processeur (%) | 72.67 | 74.22 | 99.77 | 99.76 | 68.44 | 65.11 | 83.16 | 84.72 | 
| Métrique | Séquentiel 64K |  |  |  | Séquentiel 128K |  |  |  | 
|---|---|---|---|---|---|---|---|---|
|  | Windows non natif | Windows natif | Linux libaio | Linux io_uring | Windows non natif | Windows natif | Linux libaio | Linux io_uring | 
| Lecture séquentielle |  |  |  |  |  |  |  |  | 
| Bande passante (Gio/s) | 35.596 | 35.623 | 31.867 | 31.433 | 86.791 | 92.562 | 97.05 | 97 | 
| IOPS | 583,192 | 583,638 | 522,000 | 515,000 | 710,978 | 758,252 | 795,000 | 795,000 | 
| Latence moyenne (ms) | 0.809 | 0.812 | 0.919 | 0.932 | 0.613 | 0.608 | 0.603 | 0.604 | 
| Utilisation totale du processeur (%) | 44.89 | 37.11 | 53.94 | 41.74 | 61.56 | 49.56 | 75.14 | 76.90 | 
| Métrique | 4K aléatoires |  |  |  | 64K aléatoires |  |  |  | 
|---|---|---|---|---|---|---|---|---|
|  | Windows non natif | Windows natif | Linux libaio | Linux io_uring | Windows non natif | Windows natif | Linux libaio | Linux io_uring | 
| Écriture aléatoire |  |  |  |  |  |  |  |  | 
| Bande passante (Gio/s) | 1.803 | 1.756 | 1.876 | 1.815 | 7.654 | 7.655 | 7.652 | 7.651 | 
| IOPS | 472,725 | 460,383 | 492,000 | 476,000 | 125,391 | 125,406 | 125,000 | 125,000 | 
| Latence moyenne (ms) | 0.992 | 1.028 | 0.974 | 1.007 | 3.814 | 3.816 | 3.827 | 3.828 | 
| Utilisation totale du processeur (%) | 26.00 | 20.67 | 45.76 | 22.80 | 12.22 | 9.33 | 20.07 | 10.90 | 
| Métrique | Séquentiel 64K |  |  |  | Séquentiel 128K |  |  |  | 
|---|---|---|---|---|---|---|---|---|
|  | Windows non natif | Windows natif | Linux libaio | Linux io_uring | Windows non natif | Windows natif | Linux libaio | Linux io_uring | 
| Écriture séquentielle |  |  |  |  |  |  |  |  | 
| Bande passante (Gio/s) | 44.67 | 50.087 | 52.283 | 52.25 | 50.477 | 50.079 | 52 | 52.083 | 
| IOPS | 731,859 | 820,603 | 856,000 | 856,000 | 413,495 | 410,232 | 426,000 | 427,000 | 
| Latence moyenne (ms) | 0.399 | 0.558 | 0.560 | 0.560 | 1.022 | 1.149 | 1.126 | 1.125 | 
| Utilisation totale du processeur (%) | 70.44 | 57.78 | 61.88 | 62.75 | 58.44 | 47.33 | 61.49 | 44.27 | 
À noter: Nos résultats d'IOPS sous Linux sont arrondis au millier près en raison des différences de rapport FIO entre Windows Server 2025 et Ubuntu Server 24.04.4 LTS. Les résultats relatifs à la bande passante, à la latence et à l'utilisation du processeur sont arrondis de manière uniforme sur les deux plateformes.
Les chiffres ne mentent pas
On constate immédiatement qu'Ubuntu ne surpasse pas Windows dans tous les domaines. Si libaio et io_uring ont affiché d'excellents débits lors de nos tests de bande passante en lecture aléatoire, leurs performances restent inférieures à celles de la pile NVMe native de Microsoft. Le noyau Windows NT a devancé le noyau Linux d'environ 17 % lors de nos tests de lecture aléatoire 64K, avec un débit de 91.165 Gio/s sur NVMe natif, contre 77.7 Gio/s pour io_uring.
Cependant, tout espoir n'est pas perdu pour la terreur technologique de Torvalds. Ubuntu Server a devancé de peu Windows Server dans l'un de nos tests de performance en lecture : le test séquentiel de 128 Ko. Dans ce test, libaio de Linux a obtenu les meilleurs résultats avec 97.05 Gio/s, contre 92.562 Gio/s pour la technologie NVMe native de Windows, soit une différence d'environ 5 %. Cela indique que Linux pourrait avoir un léger avantage pour la gestion de blocs de taille supérieure à la taille des unités d'indirection des disques.
La bande passante en écriture aléatoire était constante sous Linux et Windows, notamment lors des tests de performances à 64 Ko. L'écart entre les meilleurs et les pires résultats de ces tests n'était que de 0.05 %, ce qui indique que toutes les configurations de stockage ont exploité pleinement le potentiel des disques.
