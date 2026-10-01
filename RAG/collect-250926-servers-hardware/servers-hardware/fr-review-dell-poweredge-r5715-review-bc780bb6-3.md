---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r5715-review-bc780bb6-3
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache", "benchmark", "benchmarks", "open source"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [80, 120]
sha256: f06e4671ccd8791892378c5e7fd34de7503b821f4ff644d4727b093ab3b1792b
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

Le R5715 a affiché des performances prévisibles par rapport au R4715, quelle que soit la taille de la charge de travail. À 1 milliard de décimales, le R5715 a terminé en 14.537 secondes contre 5.305 secondes pour le R4715, et cet écart s'est maintenu. À 50 milliards de décimales, le R5715 a atteint 1 1,273.734 secondes tandis que le R4715 a terminé en 445.440 secondes, soit environ 2.8 à 2.9 fois plus rapide sur toute la plage de 4715 à 50 milliards de décimales. Malgré ses 8 cœurs seulement, l'EPYC 9015 est un processeur serveur dédié, doté d'une bande passante mémoire et d'un cache nettement supérieurs à ceux d'un processeur de bureau classique. Il reste ainsi bien plus performant que la plupart des processeurs grand public sur les mêmes charges de travail.
| Y-cruncher (une durée plus courte est meilleure) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cœurs \| 384 Gio de RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cœurs \| 384 Gio de RAM) | 
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
Blender 4.5 est une application de modélisation 3D open source. Ce benchmark a été réalisé à l'aide de l'utilitaire Blender Benchmark CLI. Le score est mesuré en échantillons par minute, les valeurs les plus élevées indiquant de meilleures performances.
Les résultats de Blender suivent une tendance similaire à celle de Y-Cruncher : le nombre de cœurs supérieur du R4715 se traduit directement par un débit de rendu plus élevé. Sur la scène Monster, le R4715 a atteint 523.29 échantillons par minute contre 135.21 pour le R5715. La scène Junkshop a obtenu 355.43 contre 88.61, et la scène Classroom 264.70 contre 68.48 pour le R5715. Sur l'ensemble des trois scènes, le R4715 a affiché un débit de rendu environ 3.8 à 4 fois supérieur à celui du R5715, un écart légèrement plus important que sur Y-Cruncher. Ceci illustre l'importance du nombre de cœurs pour le rendu CPU de Blender lors de la parallélisation des calculs de lancer de rayons au sein d'une même scène.
| Test de performance du processeur avec Blender 4.5 (un nombre d'échantillons par minute plus élevé est préférable) | Dell PowerEdge R4715 (AMD EPYC 9335 32 cœurs \| 384 Gio de RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cœurs \| 384 Gio de RAM) | 
|---|---|---|
| Monster | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
| Brocanteur | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
| Salle de classe | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
Points de repère Phoronix
Phoronix Test Suite est une plateforme de benchmarking automatisée et open source prenant en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org . Elle gère l'ensemble du processus, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue. Nous comparerons ici les performances des processeurs R5715 et R4715 avec les tests Stream, 7-Zip, la compilation du noyau Linux, Apache et OpenSSL.
En termes de débit pour les serveurs web Apache, le R4715 a atteint 177 839,86 requêtes par seconde, contre 123 710,75 pour le R5715, soit l'un des résultats les plus proches de toute la série. La capacité d'Apache à offrir des performances acceptables même avec un nombre de cœurs inférieur, pourvu que la bande passante mémoire soit suffisante, explique que l'écart soit ici plus faible que pour des charges de travail plus fortement parallélisées.
Le débit de transfert OpenSSL a affiché un écart plus important, le R4715 atteignant 533 318 299 283 octets par seconde contre 148 168 050 733 octets par seconde pour le R5715. Le débit cryptographique est l'une des charges de travail qui évoluent le plus rapidement avec le nombre de threads, et cet écart le reflète clairement.
Le test de compilation du noyau Linux a révélé l'un des écarts les plus marqués de la suite, le R4715 terminant en 379.53 secondes contre 1 244,86 secondes pour le R5715. La compilation du noyau est l'une des mesures les plus directes du nombre de threads qu'un système peut exécuter simultanément.
La compression 7-Zip a atteint 260 124 MIPS sur le R4715 contre 98 555 MIPS sur le R5715, ce qui est cohérent avec les résultats obtenus sur le reste de la suite.
Le débit de mémoire de flux était de 370 228,9 Mo/s sur le R4715, contre 230 123,6 Mo/s sur le R5715.
| Points de repère Phoronix | Dell PowerEdge R4715 (AMD EPYC 9335 32 cœurs \| 384 Gio de RAM) | Dell PowerEdge R5715 (AMD EPYC 9015 8 cœurs \| 384 Gio de RAM) | 
|---|---|---|
| Requêtes Apache par seconde | 177,839.86 | 123,710.75 | 
| Débit de transfert OpenSSL (octets/s) | 533,318,299,283 | 148,168,050,733 | 
| Temps de compilation du noyau (secondes) (plus le temps est court, mieux c'est) | 379.531 | 1,244.86 | 
| 7-ZIP MIPS | 260,124 | 98,555 | 
| Débit du flux (Mo/s) | 370,228.9 | 230,123.6 | 
Conclusion
Le Dell PowerEdge R5715 est une plateforme 2U dédiée au stockage, parfaitement conçue et qui justifie pleinement l'utilisation d'un châssis mono-processeur dans certains contextes de charge de travail. Les entreprises exploitant des services de fichiers, des cibles de sauvegarde, des systèmes de vidéosurveillance ou des bases de données, et qui privilégient la densité de disques et l'extensibilité des E/S à la puissance de calcul brute, trouveront dans le R5715 une solution idéale. Son fond de panier 3.5 pouces à 12 baies, prenant en charge jusqu'à 288 To de capacité brute, associé à quatre emplacements PCIe Gen5 et à la prise en charge de deux cartes réseau 3.0 OCP, offre à la plateforme une marge de progression significative sans nécessiter le passage à un châssis bi-processeur plus onéreux.
Les résultats de performance sont sans équivoque. Testé avec le processeur EPYC 9015, le R5715 affiche des performances nettement inférieures à celles du R4715 (32 cœurs) sur tous les benchmarks, comme prévu. Cependant, cette comparaison est quelque peu hors sujet. Le R5715 n'est pas conçu pour les calculs intensifs, et l'EPYC 9015 n'est pas le processeur que Dell prévoit pour la plupart de ses clients. Configurer le R5715 avec un processeur EPYC 9005 à plus grand nombre de cœurs permet de réduire considérablement cet écart, et l'architecture de la plateforme est parfaitement compatible.
Le R5715 excelle dans les domaines essentiels à ses cas d'utilisation cibles : densité de stockage, flexibilité d'extension, efficacité énergétique et gestion. iDRAC10 Enterprise offre une expérience de gestion hors bande éprouvée et cohérente, directement issue de la gamme PowerEdge de 17e génération, réduisant ainsi les coûts opérationnels pour les équipes ayant déjà investi dans la suite de gestion Dell.
Pour les acheteurs de PME et de marché intermédiaire qui cherchent à consolider leurs charges de travail de stockage sur une plateforme monoprocesseur adaptée sans surdimensionner leurs capacités de calcul, le R5715 est un excellent choix et un complément naturel au R4715 dans la gamme actuelle de Dell basée sur AMD.
