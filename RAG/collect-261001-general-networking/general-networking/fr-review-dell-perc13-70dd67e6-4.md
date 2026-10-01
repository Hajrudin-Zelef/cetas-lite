---
id: collect-261001-general-networking/general-networking/fr-review-dell-perc13-70dd67e6-4
title: "fr-review-dell-perc13-70dd67e6"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "benchmark", "mlperf"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-perc13-70dd67e6.md
source_anchor: ""
source_lines: [87, 109]
sha256: 5551f7c14f30af8e1b563672914f3bbadce35cc6337efa30a1207c9a88abd034
---

# fr-review-dell-perc13-70dd67e6

La latence lors des écritures séquentielles est restée remarquablement faible pour des blocs de plus petite taille et un nombre de threads plus faible, restant souvent inférieure à 50 µs sur des blocs de 128 8 jusqu'à 392 threads. À mesure que le nombre de threads augmentait, la latence s'est nettement améliorée. Par exemple, la latence a atteint 512 µs à 32 1 avec 1 threads et a dépassé 64 ms à XNUMX M avec XNUMX threads.
Les effets de saturation sont devenus plus évidents avec les blocs les plus volumineux et les niveaux de concurrence les plus élevés. La latence a atteint 12.4 ms à 5 Mbit/s avec 128 threads et a culminé à 50.3 ms à 10 Mbit/s avec 256 threads.
Le point de fonctionnement le plus efficace pour les charges de travail d'écriture séquentielle s'est produit aux tailles de bloc de 1 M ou 5 M avec 8 à 16 threads, où le débit a atteint 87.9 à 101.2 Gio/s tandis que la latence est restée entre 178 µs et 1.7 ms, offrant des performances solides et soutenues sans déclencher de retards excessifs dans la file d'attente d'écriture.
Performances du stockage MLPerf 2.0
Pour évaluer les performances réelles dans les environnements d'entraînement d'IA, nous avons utilisé la suite de tests MLPerf Storage 2.0. MLPerf Storage est spécialement conçu pour tester les schémas d'E/S dans des charges de travail d'apprentissage profond réelles et simulées. Il fournit des informations sur la manière dont les systèmes de stockage gèrent les défis tels que les points de contrôle et l'entraînement des modèles.
Benchmark de points de contrôle
Lors de l'entraînement de modèles de machine learning, les points de contrôle sont essentiels pour sauvegarder périodiquement l'état du modèle. Cela permet d'éviter les pertes de progression dues à des interruptions, telles que des pannes matérielles, d'effectuer des arrêts anticipés pendant l'entraînement et de permettre aux chercheurs de se diversifier à partir de différents points de contrôle pour les expériences et les ablations.
La comparaison des durées de sauvegarde aux points de contrôle a révélé que le Dell PERC13 surpassait systématiquement le PERC12, toutes configurations confondues. Le PERC13 a atteint des temps de sauvegarde compris entre 7.61 et 10.17 secondes, tandis que le PERC12 a nécessité entre 10.41 et 20.67 secondes pour les mêmes opérations. L'écart de performance était particulièrement marqué avec le modèle à 1 To, où le PERC13 a effectué les sauvegardes en un peu plus de 10 secondes, contre plus de 12 secondes pour le PERC20. Cela représente une réduction d'environ 50 % du temps de sauvegarde pour les modèles les plus volumineux.
L'analyse des résultats de débit de sauvegarde met en évidence l'utilisation supérieure de la bande passante par PERC13, offrant des taux de transfert de données constamment plus élevés. PERC13 atteint un débit compris entre 11.46 et 14.81 Go/s, avec des performances maximales sur le modèle 1T. En revanche, PERC12 atteint un maximum de 9.49 Go/s et descend à 6.98 Go/s pour la configuration la plus importante. Le nouveau contrôleur maintient des performances plus stables sur différentes tailles de modèle, ce qui suggère une meilleure optimisation pour la gestion des écritures séquentielles volumineuses, typiques des opérations de point de contrôle.
Les comparaisons de durée de charge montrent des avantages similaires pour le PERC13, bien que l'écart de performance varie selon la taille du modèle. Pour les modèles plus petits (8B, 70B), le PERC 13 a chargé les points de contrôle environ 35 à 40 % plus rapidement que le PERC12. Cependant, l'amélioration la plus spectaculaire a été observée avec le modèle 1T, où le PERC13 a chargé en 10.58 secondes contre 12 secondes pour le PERC21.22 (soit une réduction de près de 50 %). Ce temps de récupération plus rapide est crucial pour minimiser les temps d'arrêt lors de la reprise de l'entraînement après interruption.
Enfin, en examinant les métriques de débit de charge, PERC13 démontre un net avantage en termes de performances, maintenant constamment un débit supérieur à 18 Go/s sur toutes les configurations et atteignant un pic de 23.73 Go/s sur le modèle 405B. En revanche, PERC12 a affiché des performances inférieures, allant de 6.8 Go/s à 10.68 Go/s.
FIO Benchmark de performance
Bien que de nouvelles approches de test aient été intégrées à cette analyse, afin de mettre en évidence les améliorations, nous avons repris certaines données de notre précédent article sur le contrôleur PERC12 de Dell, illustrant la différence en termes de bande passante maximale et de débit maximal.
Dire que PERC13 apporte des améliorations est un euphémisme. Avec un seul volume RAID5 sur chaque contrôleur, nous avons mesuré une augmentation de 88 % de la bande passante en lecture, de 318 % de la bande passante en écriture, de 31 % des performances en lecture aléatoire 4K et de 466 % des performances en écriture aléatoire 4K. Il ne s'agit pas des performances maximales du contrôleur PERC 13 ; des vitesses supérieures sont possibles avec davantage de disques virtuels. Cependant, ce résultat reflète les performances d'un espace de noms unique lors de l'optimisation de la capacité totale.
| Charge de travail | Double PERC 12 (2 x RAID5) | Double PERC 13 (2 x RAID5) | Augmentation des performances | 
|---|---|---|---|
| 128 XNUMX lectures séquentielles | 56,107 XNUMX (Mo/s) | 105,227 XNUMX (Mo/s) | 88 % | 
| 128 XNUMX écritures séquentielles | 24,351 XNUMX (Mo/s) | 101,723 XNUMX (Mo/s) | 318 % | 
| Lectures aléatoires de 4 Ko | 13,205,656 XNUMX XNUMX (IOP) | 17,342,057 XNUMX XNUMX (IOP) | 31 % | 
| Écritures aléatoires de 4 Ko | 1,725,198 XNUMX XNUMX (IOP) | 9,758,677 XNUMX XNUMX (IOP) | 466 % | 
Nous nous sommes concentrés sur les performances des contrôleurs Dell PERC H975i et PERC H965i, exploitant le RAID 5, qui offre un excellent compromis entre capacité et protection de parité. Nous avons examiné plusieurs configurations de disques virtuels (VD) sur le Dell PERC H975i : 8 VD en RAID 5 (8R5), 4 VD en RAID 5 (4R5) et 2 VD en RAID 5 (2R5). Nous avons également testé deux configurations sur le Dell PERC H965i : 4 VD en RAID 5 (4R5) et 2 VD en RAID 5 (2R5). Les configurations ont été choisies en fonction du nombre de SSD que chaque contrôleur peut gérer. Le tout nouveau contrôleur PERC 13 peut gérer jusqu'à 16 SSD, facilement divisibles en 4 groupes RAID 5 de 4 SSD chacun. L'ancien PERC 12 ne pouvait gérer que 8 SSD, limitant ainsi ses tests à 2 groupes RAID 5 maximum. Cette configuration signifie que dans le cas d'un système 8R5, nous avons quatre RAID 4 à 5 disques sur chaque contrôleur RAID PERC.
Chaque configuration a été soumise à un processus d'analyse comparative identique, commençant par une phase de préconditionnement comprenant deux écritures complètes sur le périphérique avec des charges de travail séquentielles. Une fois l'état stable atteint, nous avons mesuré les performances selon différents modèles d'accès. Avant chaque nouveau test de charge de travail, nous avons réédité un cycle de préconditionnement avec la taille de transfert correspondante afin de garantir la cohérence des résultats.
Bande passante d'écriture séquentielle de 128 K
