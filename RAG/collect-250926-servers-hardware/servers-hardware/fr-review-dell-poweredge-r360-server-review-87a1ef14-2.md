---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r360-server-review-87a1ef14-2
title: "fr-review-dell-poweredge-r360-server-review-87a1ef14"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["benchmark", "gpu", "nvidia", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r360-server-review-87a1ef14.md
source_anchor: ""
source_lines: [65, 111]
sha256: c73de2a3d842ba68fc4a7770d1ee14274a2116f3ca0ed9ef953f7a41d2a04c7e
---

# fr-review-dell-poweredge-r360-server-review-87a1ef14

| Courant nominal/utilisation | 4.371 GIPS | 5.470 GIPS | 7.340 GIPS | 
| Courant | 69.791 GIPS | 87.013 GIPS | 28.925 GIPS | 
| Utilisation résultante du processeur | 1,575 % | 1,569 % | 397 % | 
| Évaluation/utilisation résultante | 4.401 GIPS | 5.451 GIPS | 7.317 GIPS | 
| Note résultante | 69.321 GIPS | 85.517 GIPS | 29.062 GIPS | 
| Note totale |  |  |  | 
| Utilisation totale du processeur | 1,468 % | 1,456 % | 364 % | 
| Note totale/utilisation | 4.718 GIPS | 5.661 GIPS | 8.101 GIPS | 
| Note totale | 68.912 GIPS | 82.163 GIPS | 29.246 GIPS | 
Test de vitesse du disque Blackmagic
Nous exécutons le populaire Blackmagic Disk Speed Test sur le SSD de démarrage BOSS RAID1 du système.
Vision par ordinateur UL Procyon AI
Le test Procyon d'UL évalue les performances d'une station de travail pour les applications professionnelles. Idéalement, ce test s'exécute sur un GPU, mais nous l'effectuons une fois sur le CPU pour nos tests de serveurs. Le PowerEdge R360 a obtenu des résultats conformes aux attentes, surpassant peut-être légèrement le PowerEdge R260 grâce à sa plus grande capacité de RAM.
| Temps d'inférence moyens UL Procyon (ms, plus bas est mieux) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| Mobile Net V3 | 1.02 | 1.10 | 1.46 | 
| ResNet 50 | 13.85 | 14.90 | 17.70 | 
| Création V4 | 39.35 | 44.42 | 52.38 | 
| Deep Lab V3 | 40.29 | 46.80 | 58.16 | 
| YOLO V3 | 106.94 | 117.50 | 141.53 | 
| Réel-ESRGAN | 4,280.9 | 4,617.6 | 5,740.2 | 
| Note globale | 107 | 97 | 79 | 
croque-y
y-cruncher est un programme multithread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des milliards de chiffres. Depuis son lancement en 2009, elle est devenue une application d'analyse comparative et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel. Ces serveurs ne sont pas idéaux pour les applications gourmandes en CPU, mais le PowerEdge R360 s'en sort assez respectablement pour son processeur à huit cœurs.
| y-cruncher (Temps de calcul total en secondes ; plus bas est mieux) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| 1 milliard de chiffres | 40.723 | 35.118 | 68.036 | 
| 2.5 milliards | 116.892 | 100.2 | 192.715 | 
| 5 milliards | 259.398 | 220.128 | 426.003 | 
| 10 milliards | 561.962 | N/D | N/D | 
| 25 milliards | 1,561 | N/D | N/D | 
Geekbench 6
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Le navigateur Geekbench permet de comparer n'importe quel système à ce test. Les scores monocœur étaient proches entre tous les appareils, mais le PowerEdge R260 s'est distingué en multicœur, comme prévu.
| Geekbench 6 (Plus c'est mieux) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| Processeur monocœur | 2,471 | 2,747 | 2,314 | 
| Processeur multicœur | 12,934 | 14,384 | 7,380 | 
Cinebench R23
Ce benchmark utilise tous les cœurs et threads du processeur pour générer un score global. Le résultat ici était similaire à ce que nous avons vu dans Geekbench 6.
| Cinebench R23 (Plus haut, c'est mieux) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| Multi-Core | 12,743 | 16,056 | 6,525 | 
| Single-Core | 1,772 | 2,000 | 1,669 | 
Cinebench 2024
Nous avons également commencé à exécuter le dernier test Cinebench. Les résultats ici sont similaires à ceux de Geekbench 6.
| Cinebench R23 (Plus haut, c'est mieux) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| Multi-Core | 724 | 898 | 380 | 
| Single-Core | 103 | 117 | 98 | 
Conclusion
Le PowerEdge R360 de Dell constitue une valeur prometteuse pour les PME et les applications Near Edge grâce à ses composants abordables et ses configurations polyvalentes. L'amélioration la plus significative par rapport au PowerEdge R350 est sa capacité à héberger un seul GPU NVIDIA A2, ce qui élargit son attrait pour les applications d'IA. Il améliore également son prédécesseur avec un processeur Xeon plus rapide, une bande passante mémoire supérieure et une configuration de disque de stockage BOSS N-1 plus avancée pour le système d'exploitation.
Comme toujours, Dell inclut son logiciel de gestion iDRAC populaire et facile à utiliser. Une version courte de ce serveur, le PowerEdge R260, est également disponible, tout comme une version tour (PowerEdge T360) qui prend en charge davantage d'extension. Le PowerEdge R360 offre une valeur louable pour les PME à la recherche d'un serveur léger et axé sur la valeur.
