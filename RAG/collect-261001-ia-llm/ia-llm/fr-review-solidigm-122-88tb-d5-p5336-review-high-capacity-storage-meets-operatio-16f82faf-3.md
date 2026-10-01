---
id: collect-261001-ia-llm/ia-llm/fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf-3
title: "fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "nand"]
source: docs/RAG/collect-261001-ia-llm/fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf.md
source_anchor: ""
source_lines: [63, 99]
sha256: 39aa05f63fc4e0ea3015104e9c59e99359ad9239c111166aae688f05f784f38b
---

# fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf

En écriture monothread, le Solidigm P5336 122.88 To démarre à 1,742 2,572 Mo/s et atteint environ 32 5336 Mo/s à QD61.44. Le Solidigm P461 3,029 To démarre à 6550 Mo/s, mais évolue plus rapidement, culminant à 984 6,288 Mo/s. Le Micron 32 démarre à XNUMX Mo/s et continue d'évoluer de manière constante sur toute la plage de profondeur de file d'attente, atteignant XNUMX XNUMX Mo/s à QDXNUMX. Les modèles Solidigm présentent des caractéristiques d'évolutivité différentes, tandis que le Micron maintient une progression plus linéaire tout au long du test.
Écriture de la charge de travail CDN 2
En passant aux performances d'écriture double thread, la bande passante augmente pour les trois disques. Le Solidigm P5336 61.44 To démarre à 2,771 32 Mo/s et maintient un débit relativement stable jusqu'à QD5336, avec seulement de légères fluctuations. Le Solidigm P122.88 2,468 To fonctionne dans une plage plus étroite, se maintenant entre 2,620 6550 Mo/s et 2,035 6,743 Mo/s sur toutes les profondeurs de file d'attente. Le Micron 32 affiche une évolutivité continue, commençant à XNUMX XNUMX Mo/s et atteignant XNUMX XNUMX Mo/s à QDXNUMX. Les disques Solidigm maintiennent un débit constant, tandis que le Micron affiche un profil d'évolutivité plus large sur la même plage.
Écriture de la charge de travail CDN 4
En conditions de simultanéité maximale, les deux modèles Solidigm P5336 affichent une évolutivité stable, mais limitée. Le P5336 61.44 To démarre à environ 2,935 3,062 Mo/s et atteint un pic à 5336 122.88 Mo/s, tandis que le P2,529 2,562 To démarre à 16 122.88 Mo/s et termine légèrement plus bas à 61.44 6550 Mo/s. Cela se traduit par un débit maximal inférieur d'environ 2,323 % pour le modèle 6,731 To par rapport à la version 32 To. Le Micron XNUMX, quant à lui, évolue régulièrement de XNUMX XNUMX Mo/s à XNUMX XNUMX Mo/s en QDXNUMX.
Performances d'ObjectStorage
Ce test exploite un script FIO approchant une charge de travail ObjectStorage, avec 65 % des requêtes émises à une taille de transfert de 64 Kio pour représenter les opérations courantes sur petits blocs, 15 % à 8 Mio pour les charges de travail de streaming de milieu de gamme, et 15 % supplémentaires à 64 Mio pour solliciter la gestion des gros blocs du disque. Les 5 % restants, à 1 Gio, poussent le débit séquentiel maximal. En entrelaçant ces quatre tailles de blocs dans les proportions spécifiées, il simule une charge de travail mixte qui révèle à la fois l'agilité du contrôleur pour les petites E/S et ses capacités de bande passante brute pour les transferts massifs.
Lecture aléatoire (1 thread, 40QD)
| par chaîne | Bande passante de lecture (Mo/s) | Lire les IOPS | Latence de lecture (ms) | 
|---|---|---|---|
| Micron 6550 61 To | 13,444.10 | 3,165.10 | 12.5011 | 
| Solidigm P5336 61 To | 7,117.38 | 1,673.76 | 23.4513 | 
| Solidigm P5336 122 To | 7,101.97 | 1,674.78 | 23.4385 | 
Lors de ce test de lecture aléatoire monothread et haute profondeur, les Solidigm P5336 122.88 To et P5336 61.44 To affichent des performances quasi identiques. Le modèle 122.88 To atteint 7,101.97 1,674.78 Mo/s et 23.44 61.44 IOPS avec une latence de 7,117.38 ms, tandis que la variante 1,673.76 To affiche 23.45 0.25 Mo/s et 5336 XNUMX IOPS à XNUMX ms. La différence de bande passante entre les deux capacités Solidigm est inférieure à XNUMX %, ce qui témoigne de la constance des performances de la gamme PXNUMX pour les charges de travail en lecture aléatoire.
Le Micron 6550 offre des performances nettement supérieures, atteignant 13,444.10 3,165.10 Mo/s et 12.50 5 IOPS avec une latence réduite de 4 ms. Son avantage dans ce scénario réside dans l'utilisation de la mémoire NAND TLC et d'une interface PCIe GenXNUMX, qui contribuent toutes deux à un débit de lecture aléatoire et une réactivité supérieurs à ceux des disques Solidigm GenXNUMX basés sur QLC.
Lecture séquentielle (1 thread, 40QD)
| par chaîne | Bande passante de lecture (Mo/s) | Lire les IOPS | Latence de lecture (ms) | 
|---|---|---|---|
| Micron 6550 61 To | 13,955.46 | 223.32 | 174.723 | 
| Solidigm P5336 61 To | 7,098.64 | 114.12 | 341.727 | 
| Solidigm P5336 122 To | 7,103.98 | 114.60 | 340.322 | 
Concernant les performances de lecture séquentielle, les Solidigm P5336 122.88 To et P5336 61.44 To affichent des résultats quasiment identiques. Le modèle 122.88 To atteint 7,103.98 114.60 Mo/s avec 340.32 IOPS et une latence de 61.44 ms, tandis que la version 7,098.64 To affiche 114.12 341.73 Mo/s, 0.1 IOPS et 6550 ms. La différence de performances entre les deux est inférieure à 13,955.46 %, ce qui reflète un comportement cohérent des deux capacités lors de charges de travail de lecture séquentielle soutenues. Le Micron 223.32 affiche des performances nettement supérieures, avec 174.72 96 Mo/s et XNUMX IOPS avec une latence de XNUMX ms, offrant un débit supérieur d'environ XNUMX % à celui des deux modèles Solidigm lors de ce test.
Lecture aléatoire (4 thread, 10QD)
| par chaîne | Bande passante de lecture (Mo/s) | Lire les IOPS | Latence de lecture (ms) | 
|---|---|---|---|
| Micron 6550 61 To | 13,301.67 | 3,142.01 | 12.5619 | 
| Solidigm P5336 61 To | 7,131.65 | 1,686.98 | 22.9787 | 
| Solidigm P5336 122 To | 7,131.95 | 1,690.84 | 22.9315 | 
En lecture à quatre threads et avec une profondeur de file d'attente de 10, le Solidigm P5336 122.88 To enregistre 7,131.95 1,690.84 Mo/s, 22.93 5336 IOPS et une latence de 61.44 ms. Le Solidigm P7,131.65 1,686.98 To arrive juste derrière avec 22.98 0.005 Mo/s et 6550 13,301.67 IOPS, avec une latence de 3,142.01 ms. La différence de bande passante entre les deux modèles est inférieure à 12.56 %. Le Micron 86 atteint XNUMX XNUMX Mo/s et XNUMX XNUMX IOPS avec une latence de XNUMX ms, offrant un débit supérieur d'environ XNUMX % à celui des deux disques Solidigm.
Lecture séquentielle (4 thread, 10QD)
| par chaîne | Bande passante de lecture (Mo/s) | Lire les IOPS | Latence de lecture (ms) | 
|---|---|---|---|
| Micron 6550 61 To | 13,524.00 | 218.06 | 171.040 | 
| Solidigm P5336 61 To | 7,130.97 | 115.03 | 315.565 | 
| Solidigm P5336 122 To | 7,130.99 | 114.72 | 316.304 | 
Lors de ce test de lecture séquentielle à quatre threads avec une profondeur de file d'attente de 10, le Solidigm P5336 122.88 To atteint 7,130.99 114.72 Mo/s avec 316.30 IOPS et une latence de 5336 ms. La latence du Solidigm P61.44 7,130.97 To est similaire : 115.03 315.57 Mo/s, 0.01 IOPS et 6550 ms. Les deux modèles affichent des performances séquentielles quasiment identiques sur toutes les capacités, avec une différence inférieure à 13,524.00 %. Le Micron 218.06 atteint 171.04 89 Mo/s et XNUMX IOPS avec une latence de XNUMX ms, offrant un débit environ XNUMX % supérieur à celui des deux disques Solidigm dans les mêmes conditions.
Benchmark de point de contrôle DLIO
