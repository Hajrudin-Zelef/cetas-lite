---
id: collect-250926-servers-hardware/servers-hardware/fr-review-linux-vs-windows-server-2025-native-nvme-926200e2-2
title: "fr-review-linux-vs-windows-server-2025-native-nvme-926200e2"
domain: servers-hardware
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["benchmarks", "open source", "research"]
source: docs/RAG/clean4/fr-review-linux-vs-windows-server-2025-native-nvme-926200e2.md
source_anchor: ""
source_lines: [50, 59]
sha256: 5063730d0e60a60212a766f1761444b94d508ddbce90f0f3b067f5b8da76f0fd
---

# fr-review-linux-vs-windows-server-2025-native-nvme-926200e2

Il est intéressant de noter que le noyau Linux 6.8 a remporté la victoire lors des tests de bande passante d'écriture séquentielle pour des tailles de blocs de 64 Ko et 128 Ko. Bien que la différence ne soit pas énorme, les piles logicielles open source ont battu le NVMe natif de Windows Server d'environ 2 Gio/s dans les deux cas.
Les résultats de latence ont généralement suivi ceux des tests de débit, comme l'illustre la différence entre les moyennes de lecture aléatoire. Malheureusement pour Tux, libaio et io_uring ont présenté une latence plus élevée, avec une différence maximale de 0.17 ms entre le NVMe natif de Windows Server (0.207 ms) et libaio (0.377 ms) pour des lectures aléatoires de 64 Ko.
La révélation la plus surprenante de nos tests est sans doute l'écart considérable d'utilisation du processeur entre Windows Server 2025 et Ubuntu Server 24.04.4 LTS. Dans trois des quatre tests de lecture aléatoire et séquentielle, le SSD NVMe natif de Windows Server a affiché la plus faible utilisation du processeur. Le résultat le plus remarquable a été observé lors du test de lecture séquentielle de 128 Ko, où Windows a consommé 27.34 % de ressources en moins que Linux.
L'utilisation du processeur avec libaio et io_uring s'est avérée légèrement meilleure lors des tests d'écriture aléatoire et séquentielle, mais cela n'a pas suffi à empêcher le NVMe natif sur Windows Server de l'emporter dans trois des benchmarks. Une exception notable concerne l'utilisation du processeur par libaio lors du test d'écriture aléatoire 4K, qui a atteint 45.76 % du processeur système, tandis que les autres solutions de stockage se situaient autour de 20 %.
Gagnant Gagnant, Dîner CPU
Nos résultats montrent que Windows Server et Ubuntu Server affichent des performances très proches lors de tests comparatifs, tant aléatoires que séquentiels, pour différentes tailles de blocs. En termes de bande passante, Windows Server 2025 avec NVMe natif a généralement surpassé Linux dans la plupart des tests de lecture, tandis que Linux a obtenu des résultats légèrement meilleurs en écriture. Nos mesures de latence confirment cette tendance, mais le point fort réside dans l'efficacité du processeur de Windows Server 2025 avec NVMe natif.
Microsoft a manifestement déployé des efforts considérables pour faire de sa nouvelle solution de stockage la meilleure possible. Bien qu'elle ne surpasse pas systématiquement libaio et io_uring, elle offre une alternative intéressante. Ces résultats, bien que non définitifs pour tous les cas d'utilisation et toutes les configurations serveur, peuvent aider les administrateurs système à choisir entre un serveur Windows et un serveur Linux lorsque les performances de stockage priment sur la compatibilité avec le système d'exploitation.
N'hésitez pas à nous faire part de vos impressions sur ces résultats en commentant sur nos réseaux sociaux ou sur le serveur Discord de SR ! Vous attendiez-vous à ce que Windows Server obtienne d'aussi bons résultats lors de nos tests, ou espériez-vous un résultat plus favorable pour Linux ? Souhaiteriez-vous voir d'autres distributions ou noyaux Linux pour serveurs testés ? Vos retours nous intéressent toujours, et les tests demandés par nos lecteurs, comme celui-ci, deviennent souvent nos articles préférés.
Références
Didona, D., Pfefferle, J., Ioannou, N., Metzler, B. et Trivedi, A. (13 juin 2022). Comprendre les API de stockage modernes : une étude systématique de libaio, SPDK et io_uring. SYSTOR '22, 120-121. Consulté le 3 avril 2026 sur https://atlarge-research.com/pdfs/2022-systor-apis.pdf
