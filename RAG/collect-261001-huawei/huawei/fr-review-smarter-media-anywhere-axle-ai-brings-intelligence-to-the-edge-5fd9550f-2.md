---
id: collect-261001-huawei/huawei/fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f-2
title: "fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f"
domain: huawei
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["benchmark", "diffusion", "gpu", "nvidia"]
source: docs/RAG/collect-261001-huawei/fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f.md
source_anchor: ""
source_lines: [24, 49]
sha256: 5b51b7a471db29c21a75ba2682d6fe1ad318c97e978a4cd1a05c5022294e32b3
---

# fr-review-smarter-media-anywhere-axle-ai-brings-intelligence-to-the-edge-5fd9550f

Le test FIO est un outil d'analyse comparative flexible et puissant permettant de mesurer les performances des périphériques de stockage, notamment les SSD et les disques durs. Il évalue des paramètres tels que la bande passante, les IOPS (opérations d'entrée/sortie par seconde) et la latence sous diverses charges de travail, telles que les opérations de lecture/écriture séquentielles et aléatoires. Ce test est conçu pour capturer les performances maximales et soumettre le système de stockage à de multiples charges de travail, ce qui le rend particulièrement utile pour comparer différents périphériques ou configurations. Dans ce cas, un test sur toute la surface a été réalisé, testant la capacité totale des disques afin d'obtenir une vue complète de leurs performances soutenues.
Lors du test de lecture séquentielle utilisant des blocs de 128 K, le système a fourni une bande passante de 56.4 Go/s et a maintenu 430,000 1.78 IOPS avec une latence moyenne de 45.4 milliseconde. Les performances d'écriture séquentielle à la même taille de bloc ont atteint 346,000 Go/s, produisant 2.22 4 IOPS et une latence moyenne de 46.6 millisecondes. Pour les opérations de lecture aléatoire utilisant des blocs de 11.4 K, le système a atteint 0.269 Go/s avec 4 millions d'IOPS et une faible latence moyenne de 29.1 milliseconde, soulignant le potentiel de haut débit de la baie de stockage NVMe dans des conditions d'accès intensifs. Les opérations d'écriture aléatoire à 7.1 K ont mesuré 0.432 Go/s, atteignant XNUMX millions d'IOPS avec une latence moyenne de XNUMX milliseconde, confirmant une solide capacité d'écriture soutenue, même en cas d'accès fragmenté.
| Résumé du benchmark HPE DL145 Gen 11 FIO | Bande passante – Go/s | IOPS | Latence moyenne | 
|---|---|---|---|
| Lecture séquentielle (128 Ko) | 56.4 GB / s | 430K | 1.78 ms | 
| Écriture séquentielle (128 Ko) | 45.4 GB / s | 346K | 2.22 ms | 
| Lecture aléatoire (4 Ko) | 46.6 GB / s | 11.4M | 0.269 ms | 
| Écriture aléatoire (4K) | 29.1 GB / s | 7.1M | 0.432 ms | 
Stockage direct du GPU
L'un des tests que nous avons menés sur ce banc d'essai était le test Magnum IO GPU Direct Storage (GDS). GDS est une fonctionnalité développée par NVIDIA qui permet aux GPU de contourner le CPU lors de l'accès aux données stockées sur des disques NVMe ou d'autres périphériques de stockage haute vitesse. Au lieu de faire transiter les données par le CPU et la mémoire système, GDS permet une communication directe entre le GPU et le périphérique de stockage, réduisant ainsi considérablement la latence et améliorant le débit.
Comment fonctionne le stockage direct GPU
Traditionnellement, lorsqu'un GPU traite des données stockées sur un disque NVMe, les données doivent d'abord transiter par le processeur et la mémoire système avant d'atteindre le GPU. Ce processus introduit des goulots d'étranglement, car le processeur devient un intermédiaire, ce qui ajoute de la latence et consomme de précieuses ressources système. Le stockage direct GPU élimine cette inefficacité en permettant au GPU d'accéder directement aux données depuis le périphérique de stockage via le bus PCIe. Ce chemin direct réduit la surcharge associée au déplacement des données, permettant des transferts de données plus rapides et plus efficaces.
Les charges de travail de l’IA, en particulier celles impliquant l’apprentissage profond, sont très gourmandes en données. La formation de grands réseaux neuronaux nécessite le traitement de téraoctets de données, et tout retard dans le transfert de données peut entraîner une sous-utilisation des GPU et des temps de formation plus longs. Le stockage direct GPU relève ce défi en garantissant que les données sont transmises au GPU le plus rapidement possible, en minimisant les temps d’inactivité et en maximisant l’efficacité de calcul.
En outre, GDS est particulièrement utile pour les charges de travail impliquant la diffusion de grands ensembles de données, comme le traitement vidéo, le traitement du langage naturel ou l'inférence en temps réel. En réduisant la dépendance au processeur, GDS accélère le déplacement des données et libère les ressources du processeur pour d'autres tâches, améliorant ainsi encore les performances globales du système.
Lecture séquentielle GDSIO
Lors du test de lecture séquentielle GDSIO du disque Solidigm PS1010 de 7.68 To, les performances ont évolué de manière significative avec la taille des blocs et la profondeur des E/S. Avec la plus petite taille de bloc de 16 Ko, le débit a commencé à seulement 0.2 Gio/s avec une profondeur de file d'attente de 1, puis a progressivement augmenté jusqu'à 1.3 Gio/s avec une profondeur de 128 Ko, ce qui montre une évolutivité limitée à cette granularité. Avec 128 Ko de blocs, les performances ont progressé de manière plus spectaculaire, commençant à 1.1 Gio/s et atteignant 6.5 Gio/s avec la profondeur la plus élevée. Les meilleurs résultats ont été obtenus avec une taille de bloc de 1 Mo, où le débit a atteint initialement 2.4 Gio/s et a culminé à 8.5 Gio/s avec une profondeur de file d'attente de 128 Ko, ce qui indique que le profil de performances optimal du disque est atteint avec des lectures séquentielles volumineuses et des files d'attente plus profondes.
Écriture séquentielle GDSIO
Les performances d'écriture séquentielle du Solidigm PS1010 présentent une évolutivité solide pour des tailles de blocs plus importantes, mais accusent une légère régression pour des profondeurs de file d'attente plus élevées dans les charges de travail de taille moyenne. Avec la plus petite taille de bloc de 16 Ko, les vitesses d'écriture ont débuté à 0.5 Gio/s et ont atteint un pic modeste à 0.9 Gio/s entre les profondeurs de file d'attente de 8 à 64, avant de chuter légèrement à 0.8 Gio/s à la profondeur de 128. Avec 128 Ko de blocs, les performances ont débuté à 2.2 Gio/s, ont atteint un pic de 4.3 Gio/s à la profondeur de 32, puis ont chuté à seulement 1.9 Gio/s à la profondeur de file d'attente la plus élevée, ce qui indique une saturation ou une limitation potentielle des écritures. Les meilleures performances soutenues ont été obtenues avec des tailles de bloc de 1 M, où le débit a évolué nettement de 4.1 Gio/s à la profondeur de 1 à 5.6 Gio/s aux profondeurs de 32 et 64, restant stable jusqu'à la profondeur de 128.
Résumé du GDSIO
Ce tableau fournit une analyse détaillée des mesures de performance GDSIO pour la latence et les IOPS collectées sur le SSD Solidigm D7-PS1010, mesurées pour des tailles de bloc de 16 Ko, 128 Ko et 1 Mo avec une profondeur d'E/S de 128. Avec une profondeur de file d'attente de 128, la latence et les IOPS évoluent de manière prévisible avec la taille du bloc. La lecture d'un bloc de 16 Ko a duré en moyenne 1.549 ms avec 82.3 Ko d'IOPS, tandis que la latence d'écriture était de 2.429 ms avec 52.6 Ko d'IOPS. Avec 128 Ko, la latence de lecture a atteint 2.414 ms (52.9 Ko d'IOPS) et la latence d'écriture 8.050 ms (15.9 Ko d'IOPS). À 1 Mo, la latence de lecture a atteint 14.643 ms avec 8.7 Ko d'IOPS, et la latence d'écriture a atteint 23.030 ms avec 5.6 Ko d'IOPS.
| Graphique GDSIO (tailles moyennes des blocs 16 128, 1 XNUMX et XNUMX M) | HPE DL145 Gen 11 (6 disques SSD Solidigm D7-PS1010 E3s 7.68 To) | 
|---|---|
| (Taille de bloc de 16 Ko, profondeur d'E/S de 128) Lecture moyenne | 1.3 Gio/s (1.549 ms) IOPS : 82.3 K | 
| (Taille de bloc de 16 Ko, profondeur d'E/S de 128) Écriture moyenne | 0.8 Gio/s (2.429 ms) IOPS : 52.6 K | 
| (Taille de bloc de 128 Ko, profondeur d'E/S de 128) Lecture moyenne | 6.5 Gio/s (2.414 ms) IOPS : 52.9 K | 
| (Taille de bloc de 128 Ko, profondeur d'E/S de 128) Écriture moyenne | 1.9 Gio/s (8.050 ms) IOPS : 15.9 K | 
