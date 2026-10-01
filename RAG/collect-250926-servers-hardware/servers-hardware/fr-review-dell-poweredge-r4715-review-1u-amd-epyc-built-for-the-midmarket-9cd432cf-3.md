---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf-3
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "gpu", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [85, 125]
sha256: fec266f78a167fba1ba2cafef73ea4af0e9493e2ca5c29d9df7652dd76d16277
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

Le R4715 a systématiquement surpassé le R5715 quelle que soit la taille de la charge de travail, réalisant les calculs environ 2.8 à 2.9 fois plus rapidement sur toute la plage de 1 à 50 milliards de chiffres. À 1 milliard de chiffres, le R4715 a terminé le calcul en 5.305 secondes, contre 14.537 secondes pour le R5715, et cet écart est resté constant malgré l'augmentation de la charge de travail. À 50 milliards de chiffres, le R4715 a atteint 445.440 secondes, tandis que le R5715 a nécessité 1 273,734 secondes. Ce résultat reflète directement la différence du nombre de cœurs : l'EPYC 9335 possède 32 cœurs contre 8 cœurs pour l'EPYC 9015 du R5715.
| y-cruncher (une durée plus courte est préférable) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cœurs \| 384 Gio de RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cœurs \| 384 Gio de RAM) | 
|---|---|---|
| 25 millions | 0.11 secondes | 0.25 secondes | 
| 50 millions | 0.23 secondes | 0.51 secondes | 
| 100 millions | 0.46 secondes | 1.08 secondes | 
| 250 millions | 1.22 secondes | 3.00 secondes | 
| 500 millions | 2.49 secondes | 6.60 secondes | 
| 1 milliard | 5.30 secondes | 14.53 secondes | 
| 2.5 milliard | 14.58 secondes | 41.32 secondes | 
| 5 milliard | 32.38 secondes | 92.99 secondes | 
| 10 milliard | 71.54 secondes | 202.87 secondes | 
| 25 milliard | 203.40 secondes | 576.87 secondes | 
| 50 milliard | 445.44 secondes | 1,273.73 secondes | 
Mixeur 4.5
Blender 4.5 est un logiciel de modélisation 3D open source. Ce test de performance a été réalisé à l'aide de l'utilitaire en ligne de commande Blender Benchmark. Le score est basé sur le nombre d'échantillons par minute ; plus la valeur est élevée, meilleures sont les performances.
Le R4715 a affiché un débit de rendu environ 3.8 à 4 fois supérieur à celui du R5715 sur les trois scènes, un écart légèrement plus important que sur le test Y-Ccruncher. Ceci illustre la forte augmentation des performances du moteur de rendu CPU de Blender en fonction du nombre de cœurs lors de la parallélisation des calculs de lancer de rayons. Sur la scène « Monster », le R4715 a atteint 523.29 échantillons par minute, contre 135.21 pour le R5715. Dans « Junkshop », le score était de 355.43 contre 88.61, et dans « Classroom », de 264.70 contre 68.48.
| Test de performance du processeur avec Blender 4.5 (un nombre d'échantillons par minute plus élevé est préférable) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cœurs \| 384 Gio de RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cœurs \| 384 Gio de RAM) | 
|---|---|---|
| Monster | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
| Brocanteur | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
| Salle de classe | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
Points de repère Phoronix
La suite de tests Phoronix est une plateforme d'évaluation des performances automatisée et open source qui prend en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org. Elle gère l'ensemble du processus, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue. Nous comparerons les performances des processeurs R4715 et R5715 avec les tests Stream, 7-Zip, la compilation du noyau Linux, Apache et OpenSSL.
Le débit du serveur web Apache figurait parmi les résultats les plus proches de la suite, le R4715 atteignant 177 839,86 requêtes par seconde contre 123 710,75 pour le R5715. Apache peut maintenir un débit raisonnable avec moins de cœurs lorsque la bande passante mémoire est suffisante, ce qui explique un écart plus faible ici qu'avec des charges de travail plus fortement parallélisées.
Le débit de transfert OpenSSL a affiché un écart plus important, le R4715 atteignant 533 318 299 283 octets/s contre 148 168 050 733 octets/s pour le R5715. Le débit cryptographique augmente fortement avec le nombre de threads, et cet écart le reflète directement.
Le test de compilation du noyau Linux a révélé l'un des écarts les plus marqués de la suite. Le R4715 a terminé en 379.53 secondes contre 1 244,86 secondes pour le R5715 ; la compilation du noyau étant l'une des mesures les plus directes du nombre de threads qu'un système peut exécuter simultanément.
La compression 7-Zip a atteint 260 124 MIPS sur le R4715 contre 98 555 MIPS sur le R5715, suivant de manière cohérente avec le reste de la suite.
Le débit de mémoire de flux a été mesuré à 370 228,9 Mo/s sur le R4715, contre 230 123,6 Mo/s sur le R5715.
| Points de repère Phoronix | Dell PowerEdge R4715 (AMD EPYC 9335 32 cœurs \| 384 Gio de RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cœurs \| 384 Gio de RAM) | 
|---|---|---|
| Requêtes Apache par seconde | 177,839.86 | 123,710.75 | 
| Débit de transfert OpenSSL (octets/s) | 533,318,299,283 | 148,168,050,733 | 
| Temps de compilation du noyau (secondes) (plus le temps est court, mieux c'est) | 379.531 | 1,244.86 | 
| 7-ZIP MIPS | 260,124 | 98,555 | 
| Débit du flux (Mo/s) | 370,228.9 | 230,123.6 | 
Conclusion
Le Dell PowerEdge R4715 est une plateforme 1U performante qui justifie pleinement une architecture mono-processeur pour les charges de travail courantes des PME. Les entreprises utilisant la virtualisation, des bases de données à grande échelle et des déploiements en périphérie de réseau, pour lesquelles l'efficacité des licences et la simplicité d'utilisation sont primordiales, trouveront le R4715 parfaitement adapté. Son format 1U, ses trois emplacements PCIe Gen5, ses 24 emplacements DDR5 RDIMM et ses options de stockage flexibles de 2.5 et 3.5 pouces confèrent à cette plateforme une grande polyvalence sans les coûts supplémentaires d'un châssis bi-processeur.
Les résultats de performance démontrent les capacités de la plateforme. Testé avec le processeur EPYC 9335, le R4715 a systématiquement surpassé le R5715 sur tous les benchmarks, l'avantage du nombre de cœurs étant particulièrement visible dans les charges de travail fortement parallélisées telles que la compilation du noyau, OpenSSL et Blender. Les utilisateurs n'ayant pas besoin de 32 cœurs peuvent opter pour l'EPYC 9255 (24 cœurs), l'EPYC 9135 (16 cœurs) ou l'EPYC 9015 (8 cœurs), en choisissant un processeur adapté à leurs besoins et à leur budget.
Il est important de préciser ce que le R4715 n'est pas. Il ne prend pas en charge les GPU ni les DPU, un choix délibéré pour limiter les coûts et l'encombrement. Pour les charges de travail nécessitant des accélérateurs, les modèles R6715 et R7715, basés sur AMD, sont les plus adaptés. Une connectivité réseau plus rapide est disponible via PCIe AIC, avec des options 100 GbE et 400 GbE.
Le R4715 excelle constamment dans les domaines essentiels pour ses cas d'utilisation cibles. Il offre une densité de calcul élevée dans un châssis 1U, une efficacité énergétique optimale grâce à ses alimentations Platinum et Titanium de 800 W et 1 100 W, des configurations de stockage NVMe et SAS/SATA flexibles, ainsi qu'une expérience de gestion iDRAC10 Enterprise éprouvée, parfaitement intégrée à la gamme PowerEdge de 17e génération. Pour les PME et les entreprises de taille intermédiaire recherchant une plateforme de calcul mono-socket adaptée à leurs besoins, sans surcoût en stockage ni en capacité d'extension, le R4715 est un excellent choix et un complément naturel au R5715 dans la gamme actuelle de serveurs Dell basés sur AMD.
