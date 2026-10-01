---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770ap-review-84b71047-4
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["apache", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [123, 135]
sha256: 3956e392c529e11f4e56859c292440b54b4cfcc9bd584b300eac32a4832605d3
---

# fr-review-dell-poweredge-r770ap-review-84b71047

La méthodologie de Metrum AI repose sur un outil personnalisé appelé jitter-c, qui mesure la gigue de latence de réveil par cœur, c'est-à-dire la régularité avec laquelle un thread programmé pour s'exécuter à un instant précis démarre effectivement. Cette métrique isole la variabilité de la planification du processeur des facteurs réseau, mémoire et applicatifs, offrant ainsi un point de comparaison fiable entre les générations de processeurs. En comparant un R770AP équipé de deux processeurs Xeon 6980P (256 cœurs au total) à un R760 de génération précédente doté de deux processeurs Xeon Platinum 8592+ (128 cœurs au total), l'étude a révélé que l'architecture Granite Rapids-AP réduisait la gigue de réveil p99 à environ 1 microseconde, soit environ la moitié de celle de l'ancienne plateforme, tout en doublant la densité de cœurs. Ces profils de gigue ont ensuite été intégrés à un moteur de simulation de backtesting afin de modéliser l'impact financier. Les résultats sont résumés ci-dessous.
| Résultats du backtest Metrum AI HFT | Dell PowerEdge R760 (2x Xeon 8592+, 128 cœurs) | Dell PowerEdge R770AP (2x Xeon 6980P, 256 cœurs) | 
|---|---|---|
| p99 Réveil instable | ~2 µs | ~1 µs | 
| Retour à la moyenne : Total des transactions | 5,175 | 6,229 (+ 20.4%) | 
| Retour à la moyenne : Transactions/sec | 819 | 991 (+ 21.1%) | 
| Tenue de marché : Total des transactions | 21,765 | 32,491 (+ 49.3%) | 
| Tenue de marché : Transactions/sec | 2,067 | 3,072 (+ 48.6%) | 
Comme l'a souligné Seamus Jones de Dell dans son commentaire sur l'étude, la valeur ajoutée ne réside pas dans la rapidité, mais dans la prévisibilité de cette rapidité. En effet, en trading, un système rapide mais incohérent est source de risque. À l'inverse, un système déterministe constitue un atout stratégique.
Conclusion
Le Dell PowerEdge R770AP occupe une place bien définie au sein de la gamme PowerEdge de 17e génération. Il ne remplace pas le R770 et Dell ne le présente pas comme tel. Le R770 demeure la plateforme Intel 2U polyvalente et hautement configurable qu'il a toujours été, avec prise en charge des GPU, stockage mixte SAS/SATA/NVMe, options de processeur E-core et P-core, et jusqu'à 8 To de mémoire répartis sur 32 emplacements DIMM. Pour les entreprises exécutant des solutions de virtualisation générales, des applications d'entreprise mixtes ou des charges de travail tirant parti de cette flexibilité de configuration, le R770 reste le choix idéal.
Le R770AP est conçu pour les charges de travail pour lesquelles le R770 n'a jamais été optimisé. En adoptant la plateforme Granite Rapids-AP, avec son architecture mémoire à 12 canaux, jusqu'à 128 cœurs de traitement par socket et 504 Mo de cache L3, Dell a créé un système 2U qui privilégie la densité de calcul, la bande passante mémoire et la déterminisme d'exécution à la polyvalence. Nos tests de performance reflètent cette priorité : la bande passante de STREAM a quasiment doublé, le rendu Blender s'est amélioré de 29 à 36 % et la mise à l'échelle du processeur graphique s'est étendue de manière constante à mesure que les ensembles de travail dépassaient la capacité du cache. La régression d'Apache est un point important à noter, car elle démontre que la topologie NUMA du R770AP nécessite une prise en compte de la charge de travail pour exploiter pleinement ses performances, et que toutes les applications ne bénéficieront pas de ce changement de plateforme sans optimisation.
Les tests d'IA Metrum publiés par Dell en parallèle de cette plateforme mettent en lumière le déterminisme sous-jacent. Réduire de moitié la gigue de planification p99 tout en doublant la densité des cœurs représente une amélioration architecturale significative pour les entreprises exécutant des opérations de trading haute fréquence, des moteurs de risque en temps réel, des analyses en mémoire à grande échelle et des simulations massivement parallèles. Pour ces charges de travail, le R770AP est une plateforme performante et parfaitement adaptée. Pour toutes les autres applications, les R770 et R7725 restent les options les plus pertinentes au sein de la gamme PowerEdge.
