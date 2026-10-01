---
id: collect-250926-servers-hardware/servers-hardware/fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655-2
title: "fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Broadcom"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "exploit"]
source: docs/RAG/clean4/fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655.md
source_anchor: ""
source_lines: [17, 45]
sha256: a2a8e31a4ddf261ba8081c6d41fcd2372bbe40f60a683fb9bfc3ecdefc9cdc7f
---

# fr-review-300-gb-s-in-2u-the-dell-poweredge-r7725xd-resets-expectations-for-stor-1e3ea655

Présentation du Dell PowerEdge R7725xd iDRAC 10
Cette génération du R7725xd, comme de nombreuses autres plateformes de 17e génération que nous avons testées, intègre la nouvelle plateforme iDRAC 10 de Dell, qui centralise la gestion à distance, la surveillance de l'état du système et le contrôle hors bande. Le tableau de bord offre un aperçu immédiat de l'état général du système, de l'état du stockage et de l'activité récente. Sur notre unité de test, le rapport d'état du système et du stockage est au vert, confirmant le fonctionnement normal du serveur. Les informations système essentielles, telles que le modèle, le nom d'hôte, la version du BIOS, le niveau du firmware iDRAC, l'adresse IP et les informations de licence, sont affichées à droite de l'interface.
Le tableau de bord comprend également un panneau récapitulatif des tâches affichant les opérations terminées, en attente et en cours. En dessous, une liste des journaux récents recense les événements d'intrusion dans le châssis et les messages relatifs à l'alimentation, offrant ainsi un aperçu rapide des changements d'état du matériel sans nécessiter de navigation dans les menus. Le panneau de la console virtuelle est accessible dans le coin inférieur droit pour un contrôle KVM à distance complet.
La section stockage d'iDRAC 10 offre une vue d'ensemble complète de tous les disques physiques installés sur le R7725xd. Le panneau récapitulatif affiche le nombre total de disques connectés, accompagné d'un graphique circulaire illustrant leur état. Dans cette configuration, 24 SSD NVMe sont actifs et prêts à l'emploi, et deux périphériques de démarrage supplémentaires sont présents dans le système, distincts du banc NVMe principal en façade.
À droite, le panneau « Résumé des disques » détaille les disques physiques et les disques virtuels associés. Le R7725xd utilisant une architecture NVMe directe sans contrôleur RAID traditionnel, tous les disques sont identifiés comme non-RAID et adressables individuellement, conformément à la conception du système pour les grands pools NVMe et les plateformes SDS.
Sous le résumé d'état, la section « Événements de stockage récemment enregistrés » répertorie les journaux d'insertion de chaque SSD PCIe, classés par baie et emplacement. Cet enregistrement confirme la bonne détection des disques dans toutes les baies et permet d'identifier tout problème d'insertion, de câblage ou d'échange à chaud. Pour les déploiements de grande envergure, ces journaux sont utiles pour suivre le provisionnement des disques ou vérifier que la capacité a été correctement allouée.
La dernière capture d'écran présente la vue détaillée des périphériques NVMe dans iDRAC10. Chaque disque NVMe installé dans le système est listé avec son état, sa capacité et son emplacement. La sélection d'un disque permet d'afficher le détail complet de ses caractéristiques.
Dans cet exemple, le panneau d'informations du disque affiche la chaîne de modèle complète, le protocole du périphérique, le format et les paramètres PCIe négociés. Les périphériques NVMe fonctionnent à une vitesse de liaison de 32 GT/s avec une connexion x4 négociée, confirmant ainsi que les disques exploitent toute la bande passante du fond de panier PCIe Gen5 du système. La section d'informations indique également le pourcentage d'endurance, l'état des disques disponibles et le type de protocole, permettant ainsi aux administrateurs de surveiller l'état et la durée de vie prévue des disques.
Ce rapport détaillé sur les disques est précieux dans les configurations NVMe haute densité où la largeur de la liaison, la vitesse négociée et l'état du support influencent directement le comportement de la charge de travail et les performances de stockage.
Globalement, l'interface iDRAC 10 offre une vue claire et axée sur le matériel de l'architecture de stockage NVMe du R7725xd, permettant une validation facile de l'état de la liaison, de l'état du disque et de l'intégrité du système en un coup d'œil.
Performances du Dell PowerEdge R7725xd
Avant les tests, notre système a été configuré avec une configuration à la fois équilibrée et performante. Il est équipé de deux processeurs AMD EPYC 9575F, chacun doté de 64 cœurs haute fréquence, et de 24 modules DIMM DDR5 de 32 Go fonctionnant à 6 400 MT/s. Pour le stockage, le châssis est entièrement équipé de 24 SSD NVMe Micron 9550 PRO U.2 de 15.36 To, chacun connecté via une interface PCIe Gen5 x4 dédiée. Ceci offre une capacité brute totale de 368.64 To, et les disques Micron 9550 PRO atteignent des vitesses de lecture séquentielles allant jusqu'à 14 000 Mo/s et des vitesses d'écriture séquentielles allant jusqu'à 10 000 Mo/s. La mise en réseau est assurée par quatre adaptateurs Broadcom BCM57608 qui fournissent un total de huit ports 200 Gb, ainsi que par une carte réseau OCP BCM57412 offrant deux ports 10 gigabits supplémentaires.
Spécifications du système de test
- CPU: 2 processeurs AMD EPYC 9575F à 64 cœurs haute fréquence
- Mémoire: 24 x 32 Go DDR5 à 6400 MT/s
- Stockage: 24 disques Micron 9550 PRO U.2 de 15.36 To (connectés chacun sur 4 lignes PCIe Gen5) ; compatible avec des disques jusqu’à 128 To actuellement, et des capacités supérieures à venir.
- Réseau: 4 cartes réseau Broadcom BCM57608 2x200G, 1 carte réseau OCP BCM57412 2x10Gb
- Commutateur: Dell PowerSwitch Z9664
Benchmark de performance FIO
Pour mesurer les performances de stockage du PowerEdge R7725xd, nous avons utilisé des indicateurs standard du secteur et l'outil FIO. Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Aléatoire 4K – 1M
- Séquentiel 4K – 1M
FIO – Local – Bande passante
Lors des tests d'accès local aux 24 disques NVMe PCIe Gen5 du Dell PowerEdge R7725xd, le système affiche les performances attendues d'une plateforme où chaque disque est connecté aux processeurs via une liaison PCIe Gen5 x4 complète. Sans couche réseau, il s'agit du débit interne pur de l'architecture de stockage Gen5 de Dell et de la bande passante PCIe de la plateforme AMD EPYC exploitée sans restriction.
Les lectures séquentielles débutent à 184 Go/s avec des blocs de 4 Ko et augmentent rapidement avec la taille des blocs. De 512 Ko à 1 Mo, le serveur maintient un débit constant de 312 à 314 Go/s, ce qui témoigne de l'excellente capacité du système à agréger les 24 × 4 voies Gen5 pour obtenir une bande passante de lecture soutenue, sans aucun goulot d'étranglement au niveau du contrôleur.
Les écritures séquentielles suivent une courbe différente, mais restent dans la plage attendue. Commençant à 149 Go/s, les résultats augmentent jusqu'à une centaine de Go/s et atteignent 182 Go/s à 1 million d'écritures. Ceci correspond au comportement en écriture des SSD Micron 9550 PRO et à la surcharge inhérente aux écritures NVMe hautement parallèles sur un si grand nombre de périphériques indépendants.
Les performances en lecture aléatoire constituent un autre point fort. Le système atteint des vitesses proches de 300 Go/s avec les plus petites tailles de blocs, diminue légèrement en moyenne, puis remonte entre 200 et 300 Go/s avec les plus grandes tailles de blocs. À 1 Mo, les lectures aléatoires atteignent un maximum de 318 Go/s, démontrant ainsi la capacité de la plateforme à répartir uniformément les opérations mixtes sur l'ensemble des 24 disques.
Les écritures aléatoires sont effectuées à un débit inférieur, ce qui est typique des tâches de métadonnées dispersées et d'allocation d'écriture sur un large ensemble NVMe. Les résultats se maintiennent entre 140 et 160 Go/s pendant la majeure partie du test et diminuent légèrement pour atteindre un peu moins de 100 Go/s à 1 Mbit/s.
FIO – Local – IOPS
