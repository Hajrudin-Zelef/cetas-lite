---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-c6615-server-review-e5a753e4-3
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "benchmark", "inference", "intel", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [121, 157]
sha256: 08adf649fa96920d986fa61eb1c5c25df54e25318450ca5e6233bb3a9a83b7b0
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

| CPU 8K | FPS 121 | FPS 121 | FPS 118 | FPS 119 | FPS 119.75 | 
Test de vitesse du disque Blackmagic
Vient ensuite le test de vitesse du disque Blackmagic. Ce test exécute un exemple de fichier de 5 Go pour les vitesses de lecture et d'écriture. Comme il est monothread, il n'affichera pas les vitesses les plus élevées du disque, mais il donne quand même une bonne perspective.
Les C6615 ont une carte BOSS à l'intérieur, utilisant deux disques M.2 en RAID1, de sorte que les performances sont légèrement dégradées pour la fiabilité. Pour les vitesses d'écriture, nous avons constaté une moyenne de 991.6 Mo/s et pour les vitesses de lecture, une moyenne de 2,801 XNUMX Mo/s.
| Test de vitesse du disque Blackmagic | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Écrire | 999.8 Mo / s | 977.4 Mo / s | 991.4 Mo / s | 997.7 Mo / s | 991.6 Mo / s | 
| Lire | 2,807.4 Mo / s | 2,790.1 Mo / s | 2,828.0 Mo / s | 2,780.4 Mo / s | 2,801.5 Mo / s | 
Y-Cruncher
y-cruncher est un programme multi-thread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des billions de chiffres. Depuis son lancement en 2009, il est devenu une application populaire d'analyse comparative et de test de résistance pour les overclockeurs et les passionnés de matériel.
Pour nos vitesses moyennes, nous avons vu 9.5 secondes pour 1 milliard, 24.20 secondes pour 2.5 milliards et 50.73 secondes pour 5 milliards. Sur les calculs des chiffres les plus significatifs, nous avons vu 105.73 secondes pour 10 milliards, 288.85 secondes pour 25 milliards et 633.5 secondes pour 50 milliards.
| Y Cruncher (temps de calcul total, en secondes) | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| 1 milliard | 9.587 | 9.459 | 9.350 | 9.633 | 9.507 | 
| 2.5 milliard | 24.490 | 24.225 | 23.334 | 24.740 | 24.197 | 
| 5 milliard | 51.427 | 50.990 | 49.303 | 51.214 | 50.734 | 
| 10 milliard | 107.084 | 107.646 | 103.772 | 107.443 | 105.736 | 
| 25 milliard | 291.918 | 290.944 | 280.632 | 291.902 | 288.849 | 
| 50 milliard | 641.709 | 640.289 | 619.100 | 640.917 | 635.504 | 
Benchmark de vision par ordinateur UL Procyon AI
UL Procyon AI Inference est conçu pour évaluer les performances d'une station de travail dans des applications professionnelles. Il est important de noter que ce test n'exploite pas les capacités de plusieurs processeurs. Plus précisément, cet outil évalue la capacité de la station de travail à gérer les tâches et les flux de travail basés sur l'IA, fournissant une analyse détaillée de son efficacité et de sa rapidité de traitement des algorithmes et applications d'IA complexes.
Pour ce test, nous utilisons Procyon V2.7.0. Dans ce test, des temps inférieurs sont meilleurs. Sur l’ensemble des nœuds, les moyennes étaient de 3.91 ms sur MobileNet V3, de 8.4.0 ms pour Resnet50 et de 29.47 ms. Sur le reste des scores, nous avons vu 30.96 ms sur DeepLab V3, 44.68 ms sur YOLO V3 et 2008.65 ms sur Real-ESRGAN. Pour le score global, les nœuds étaient en moyenne de 133.5.
| Vision par ordinateur UL Procyon (Temps d'inférence moyen) | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Mobile Net V3 | 3.87 ms | 3.94 ms | 3.84 ms | 4.00 ms | 3.91 ms | 
| ResNet50 | 8.47 ms | 8.45 ms | 8.23 ms | 8.46 ms | 8.40 ms | 
| Création V4 | 29.76 ms | 29.55 ms | 28.74 ms | 29.84 ms | 29.47 ms | 
| Deep Lab V3 | 30.39 ms | 30.21 ms | 33.18 ms | 30.07 ms | 30.96 ms | 
| YOLO V3 | 44.71 ms | 44.58 ms | 44.79 ms | 44.63 ms | 44.68 ms | 
| Réel-ESRGAN | 2003.18 ms | 1971.97 ms | 2018.26 ms | 2041.18 ms | 2008.65 ms | 
| Note globale | 134 | 134 | 133 | 133 | 133.5 | 
Conclusion
Les nœuds Dell PowerEdge C6615 offrent un seul processeur AMD EPYC avec jusqu'à 64 cœurs et six emplacements DDR5 prenant en charge des DIMM de 96 Go. Le châssis C6600 qui héberge ces nœuds propose quelques configurations de stockage. Notre système d’évaluation dispose du fond de panier SSD 8x E3.S Gen5. Dans la conception C6600, chaque nœud a accès à deux de ces SSD ; le châssis fournit simplement l'alimentation et un accès câblé direct aux disques. Pour la gestion, chaque C6615 propose iDRAC ; le châssis n'a pas de gestion dédiée.
Nous avons évalué de manière indépendante les capacités de chaque nœud C6615 lors de nos tests de performances et avons calculé la moyenne des scores des quatre nœuds pour identifier les anomalies de performances. Les données de performances mettent en évidence que les nœuds fonctionnent de manière cohérente, sans valeurs aberrantes ni performances inégales. Cette prévisibilité est essentielle pour les fournisseurs de services et les clients hyperscale qui peuvent bénéficier de systèmes denses comme celui-ci.
Nous avons trouvé le système bien conçu pour le cas d'utilisation prévu ; notre seul reproche concerne la prise en charge relativement limitée des SSD Gen5 : seulement deux disques par nœud. Dell suggérerait probablement que les clients à forte densité de calcul n'ont pas besoin d'autant de stockage local et que le refroidissement de davantage de disques Gen5 constitue un défi technique sérieux, et ils ont probablement raison, nous préférons simplement plus de disques que moins à presque chaque occasion. Une autre remarque qui mérite d'être mentionnée, nous examinons le C6615 ici, mais comme indiqué en haut de cette revue, Dell propose des types de nœuds supplémentaires pour cette plate-forme, le C6620 basé sur Intel est disponible dans une version refroidie par liquide, que certains peut trouver convaincant.
Les nœuds de calcul Dell PowerEdge C6615 offrent aux fournisseurs de services une étonnante combinaison de performances par rack U. Nous avons déjà vu de nombreuses configurations 2U4N, mais cette conception permet une plus grande largeur, et donc une plus grande flexibilité d'extension, dans chaque serveur que de nombreux systèmes concurrents. Associez l'excellente conception à des logiciels de gestion comme iDRAC et OpenManage Enterprise et nous sommes de grands fans du résultat final.
Page produit Dell PowerEdge C6615
