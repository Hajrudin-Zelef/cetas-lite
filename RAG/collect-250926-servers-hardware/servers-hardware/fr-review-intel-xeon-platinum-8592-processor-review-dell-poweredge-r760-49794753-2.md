---
id: collect-250926-servers-hardware/servers-hardware/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753-2
title: "fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["intel", "amd", "valuation"]
source: docs/RAG/clean4/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753.md
source_anchor: ""
source_lines: [79, 122]
sha256: 3a3c42b03d550684c06b9b4761ab9ecce28e36eebd6d91c6b0543ec90ae05928
---

# fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753

|  | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Compression |  |  |  | 
| Utilisation actuelle du processeur | 5,609 % | 5,732 % | 5,482 % | 
| Courant nominal/utilisation | 4.912 GIPS | 3.912 GIPS | 4.628 GIPS | 
| Courant | 275,503 GIPS | 224.209 GIPS | 253.724 GIPS | 
| Utilisation résultante du processeur | 5,605 % | 5,669 % | 5,475 % | 
| Évaluation/utilisation résultante | 4.883 GIPS | 3.923 GIPS | 4.628 GIPS | 
| Note résultante | 273.716 GIPS | 222.407 GIPS | 253.382 GIPS | 
| Décompression |  |  |  | 
| Utilisation actuelle du processeur | 6,243 % | 5,852 % | 6,219 % | 
| Courant nominal/utilisation | 3.635 GIPS | 3.423 GIPS | 3.745 GIPS | 
| Courant | 226.917 GIPS | 200.350 GIPS | 231.916 GIPS | 
| Utilisation résultante du processeur | 6,232 % | 5,894 % | 6,129 % | 
| Évaluation/utilisation résultante | 3.654 GIPS | 3.385 GIPS | 3.871 GIPS | 
| Note résultante | 227.744 GIPS | 199.363 GIPS | 237.259 GIPS | 
| Note totale |  |  |  | 
| Utilisation totale du processeur | 5,919 % | 5,781 % | 5,802 % | 
| Note totale/utilisation | 4.269 GIPS | 3.654 GIPS | 4.249 GIPS | 
| Note totale | 250.730 GIPS | 210.363 GIPS | 245.320 GIPS | 
Inférence UL Procyon AI
La suite de tests d' inférence IA Procyon d'UL évalue les performances de différents moteurs d'inférence IA utilisant des réseaux neuronaux de pointe. Ces tests ont été effectués uniquement sur le processeur. Chaque valeur représente un temps d'inférence moyen (plus le temps est court, meilleures sont les performances), et la dernière ligne indique un score global (plus le score est élevé, meilleures sont les performances). Les différences entre ces processeurs sont moins marquées dans ce test : le Xeon 8480+ se distingue même dans certains tests, mais au final, le Xeon 8592+ obtient tout de même un score global supérieur.
|  | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Mobile Net V3 | 3.17 | 3.71 | 2.34 | 
| ResNet 50 | 5.23 | 4.37 | 5.76 | 
| Création V4 | 18.66 | 22.59 | 21.70 | 
| Deep Lab V3 | 23.99 | 28.25 | 23.00 | 
| YOLO V3 | 35.47 | 40.28 | 30.81 | 
| RÉEL-ESRGAN | 1021.49 | 1277.08 | 1535.27 | 
| Note globale | 196 | 161 | 191 | 
Processeur Z
Bien qu’il ne s’agisse pas exactement d’une référence, nous avons pensé qu’il valait la peine de mentionner les différences de vitesse de mémoire, le seul changement étant le processeur installé.
Voici d’abord une photo de l’onglet mémoire des Sapphire Rapids Xeon 6430 qui étaient initialement installés dans notre R760.
Voici ensuite une photo de l'onglet mémoire de l'Emerald Rapids Xeon 8592+.
Il existe une nette différence dans la synchronisation de la mémoire sur le Xeon 8592+, ce qui la rend un peu inférieure à celle de notre Xeon 6430, mais probablement pas suffisamment pour avoir un impact sur les performances dans la plupart des applications du monde réel.
Test de mémoire
En analysant les informations mémoire fournies par CPU-Z, nous pouvons exécuter des microbenchmarks depuis Clamchowder . Cet outil est conçu pour évaluer les paramètres liés au processeur et à la mémoire, notamment la taille des fichiers ROB et de registres, la latence de cohérence des verrous et du cache, ainsi que les performances du cache et de la mémoire. Nous avons choisi ce test pour analyser les gains de performance entre les différentes générations de processeurs et pour différentes tailles d'accès.
Pour ce test, nous avons utilisé 64 threads et exécuté un profil complet de 64 Ko à 3,145,728 5 XNUMX Ko. En nous concentrant sur les performances une fois que nous quittons le cache du processeur et accédons à la mémoire système, nous pouvons constater une augmentation significative du débit avec le Xeon de XNUMXe génération dans tous les domaines lors du test d'écriture.
En ce qui concerne la bande passante de lecture, l’histoire n’est pas aussi dramatique, mais remarquable. La différence de vitesse d'horloge entre le processeur de 4e génération et le processeur de 5e génération que nous avons testés devient un peu plus claire ici dans ce test. Malgré le léger inconvénient de la 5e génération, nous pouvons toujours constater que la 5e génération Xeon Scalable peut offrir une baisse de l'augmentation des performances.
Conclusion
Comme prévu, le processeur Intel Xeon Platinum 8592+ a affiché des performances nettement supérieures aux Xeon 6430 dans la plupart des tests, ce qui est logique puisqu'il ne s'agit pas de processeurs concurrents directs. En intégrant le Xeon 8480+, l'écart de performances se réduit considérablement, ce dernier étant un prédécesseur direct du Xeon 8592+. Un avantage majeur des processeurs Xeon de 5e génération réside dans leur compatibilité avec le même socket que la 4e génération, permettant ainsi la mise à niveau de plateformes plus anciennes sur site, comme nous l'avons fait pour certains de nos serveurs.
Nous avons constaté des améliorations significatives des performances sur toutes nos charges de travail entre les deux processeurs Intel Xeon de 4e génération et le Xeon 5+ de 8592e génération. Les plus grandes améliorations ont été visibles dans nos charges de travail informatiques, telles que y-cruncher. Au cours de nos courses à 25 milliards de chiffres, nous avons mesuré des temps sur le Gold 6430 et le Platinum 8480+ de 229 et 187 secondes respectivement, tombant à seulement 157 secondes sur le Platinum 8592+. Les performances ont également été considérablement améliorées dans Cinebench R23 et 2024, passant du Platinum 4+ de 8480e génération au 5+ de 8592e génération. Dans la R23, nous avons vu les performances passer de 79 110 à 2024 4,699 km, et en 6,001, les vitesses sont passées de XNUMX XNUMX à XNUMX XNUMX.
L'Intel Xeon Platinum 8592+ marque un grand pas en avant pour l'entreprise et même s'il ne peut pas aller de cœur à cœur avec les puces robustes d'AMD, le 8592+ possède des atouts uniques qui en font un choix convaincant. Intel a considérablement optimisé ces processeurs pour les charges de travail IA et HPC, offrant des performances robustes à tous les niveaux.
