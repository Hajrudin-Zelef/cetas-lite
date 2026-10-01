---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770ap-review-84b71047-3
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["apache", "benchmark", "intel", "open source"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [75, 122]
sha256: 76fc390dfba1ac1e037c0352dc77765d485577df96e8b5fbb9533b60e23575a5
---

# fr-review-dell-poweredge-r770ap-review-84b71047

Pour évaluer le R770AP, nous l'avons comparé directement au R770. Le R770AP est équipé de deux processeurs Intel Xeon 6978P, chacun doté de 120 cœurs, pour un total de 240 cœurs et 3 To de mémoire DDR5. Le R770, quant à lui, intègre deux processeurs Intel Xeon 6787P, pour un total de 172 cœurs et 2 To de mémoire DDR5.
Afin de solliciter les processeurs des deux systèmes, nous avons utilisé un ensemble ciblé de tests de calcul. y-cruncher a permis d'évaluer le débit arithmétique brut et les performances en virgule flottante multithread. Blender a fourni une charge de travail de rendu réaliste, évolutive en fonction du nombre de cœurs disponibles et de la bande passante mémoire. La suite de tests Phoronix a complété cet ensemble avec une collection plus large de charges de travail gourmandes en ressources processeur, offrant ainsi une vision plus complète des performances de calcul soutenues sur les deux plateformes.
Spécifications du système de test
- Plate-forme: Dell PowerEdge R770AP
- CPU: Double Intel Xeon 6978P, 120 cœurs
- Mémoire: 3 To DDR5
- Stockage: Boss RAID1
croque-y
y-cruncher est une application populaire de test de performance et de résistance des systèmes, lancée en 2009. Ce test multithread et évolutif calcule Pi et d'autres constantes jusqu'à des billions de décimales. Plus le test est rapide, mieux c'est. Ce logiciel s'est avéré excellent pour tester les plateformes à grand nombre de cœurs et démontrer les avantages de calcul entre les plateformes mono-processeur et bicœur.
Dans le benchmark y-cruncher, le R770AP a systématiquement surpassé le R770, quelle que soit la taille des données testées. Lors du test avec un milliard de décimales, le R770AP a terminé en 2.692 secondes, contre 2.753 secondes pour le R770. À 10 milliards de décimales, le R770AP a réalisé un temps de 30.399 secondes, contre 34.873 secondes pour le R770. À 50 milliards de décimales, le R770AP a affiché un temps de 192.128 secondes, contre 221.255 secondes pour le R770. L'écart s'est creusé pour la charge de travail la plus importante : le test avec 100 milliards de décimales a été réalisé en 430.208 secondes par le R770AP, contre 491.737 secondes par le R770, soit une différence d'environ 61 secondes et un gain de performance d'environ 12.5 % pour le R770AP.
| Y-cruncher (une durée plus courte est meilleure) | Dell PowerEdge R770 (2x Intel Xeon 6787P \| 2 To de RAM) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 To de RAM) | 
|---|---|---|
| 1 milliard | 2.753 secondes | 2.692 secondes | 
| 2.5 milliard | 7.365 secondes | 6.747 secondes | 
| 5 milliard | 16.223 secondes | 14.235 secondes | 
| 10 milliard | 34.873 secondes | 30.399 secondes | 
| 25 milliard | 99.324 secondes | 86.298 secondes | 
| 50 milliard | 221.255 secondes | 192.128 secondes | 
| 100 milliard | 491.737 secondes | 430.208 secondes | 
Mixeur
Une application de modélisation 3D open source. Ce benchmark a été réalisé avec l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Dans le benchmark Blender 4.3, le R770AP a surpassé le R770 sur les trois scènes. Sur la scène « Monster », le R770AP a atteint 2 200,116 échantillons par minute, contre 1 706,002 pour le R770. Sur la scène « Junkshop », le R770AP a réalisé 1 565,643 échantillons par minute, contre 1 169,370 pour le R770. Enfin, sur la scène « Classroom », le R770AP a obtenu 1 076,122 échantillons par minute, contre 791.475 pour le R770, soit un gain de performance d'environ 36 % sur cette charge de travail.
| Test de performance du processeur avec Blender 4.3 (un nombre d'échantillons par minute plus élevé est préférable) | Dell PowerEdge R770 (2x Intel Xeon 6787P \| 2 To de RAM) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 To de RAM) | 
|---|---|---|
| Monster | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
| Brocanteur | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
| Salle de classe | 1 076,122 échantillons/min | 1 076,122 échantillons/min | 
Points de repère Phoronix
Phoronix Test Suite est une plateforme d'analyse comparative automatisée et open source prenant en charge plus de 450 profils de test et plus de 100 suites de tests via OpenBenchmarking.org. Elle gère l'ensemble du processus, de l'installation des dépendances à l'exécution des tests et à la collecte des résultats, ce qui la rend idéale pour les comparaisons de performances, la validation matérielle et l'intégration continue. Nous comparerons ici les performances des R770AP et R770 avec les tests Stream, 7-Zip, la compilation du noyau Linux, Apache et OpenSSL.
Discussions
Lors du test de bande passante mémoire Stream, le R770AP a réalisé une nette amélioration par rapport au R770, atteignant 869 965,3 Mo/s contre 472 135,6 Mo/s. Cela représente quasiment le double de la bande passante mémoire du système de référence, ce qui témoigne de la configuration mémoire plus importante et plus rapide du R770AP.
7-Zip
Dans le test de compression 7-Zip, le R770AP a obtenu un score de 806 375 MIPS, contre 628 206 MIPS pour le R770, une nette amélioration due au nombre de cœurs plus élevé des processeurs 6978P.
Compilation du noyau
Dans le test de compilation du noyau Linux, où un temps plus court est préférable, le R770AP a terminé la compilation allmod en 176.391 secondes contre 188.793 secondes sur le R770, réduisant ainsi le temps de compilation d'environ 12 secondes.
Apache
Le test Apache a été le seul domaine où le R770 a légèrement surpassé le R770AP, avec un score de 60 258,5 requêtes par seconde contre 48 729,63 pour le R770AP. Ce résultat est important, car les charges de travail des serveurs web n'évoluent pas toujours de manière linéaire avec le nombre de cœurs et peuvent être influencées par la latence mémoire et les caractéristiques d'E/S.
OpenSSL
Dans le test de vérification OpenSSL, le R770AP a obtenu un score de 2 515 270 390 853 vérifications/s contre 2 216 883 554 350 vérifications/s sur le R770, un gain significatif en débit cryptographique qui met en évidence l’efficacité de calcul du 6978P à grande échelle.
| Points de repère Phoronix | Dell PowerEdge R770 (2x Intel Xeon 6787P 86C) | Dell PowerEdge R770AP (2x Intel Xeon 6978P \| 3 To de RAM) | 
|---|---|---|
| Discussions | 472,135.6 Mo / s | 869,965.3 Mo / s | 
| 7-ZIP | 628,206 XNUMX MIP/s | 806,375 XNUMX MIP/s | 
| Compilation du noyau (allmod) (plus bas est mieux) | 188.793 secondes | 176.391 secondes | 
| Apache (requêtes par seconde) | 60,258.5 XNUMX R/s | 48,729.63 XNUMX R/s | 
| OpenSSL | 2,216,883,554,350 XNUMX XNUMX XNUMX Vérifications | 2,515,270,390,853 XNUMX XNUMX XNUMX Vérifications | 
Dell PowerEdge R770AP : Performances déterministes et trading haute fréquence
Bien que notre suite de tests standard se concentre sur le débit de calcul, la bande passante mémoire et la montée en charge générale, les priorités de conception du R770AP s'étendent à un domaine que nous ne testons généralement pas : la déterminisme d'exécution à la microseconde. Afin d'illustrer les capacités de cette plateforme pour son public cible le plus exigeant, Dell a publié une note technique en partenariat avec Metrum AI, évaluant le R770AP spécifiquement pour les charges de travail de trading haute fréquence. Nous n'avons pas réalisé ces tests, ni audité les résultats de manière indépendante. Néanmoins, nous en incluons un résumé ici, car il démontre de la manière la plus directe en quoi ce serveur est un produit distinct du R770.
