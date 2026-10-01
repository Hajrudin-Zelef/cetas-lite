---
id: collect-250926-servers-hardware/servers-hardware/fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356
title: "fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd"]
source: docs/RAG/clean4/fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356.md
source_anchor: ""
source_lines: [1, 50]
sha256: 9a95e53f7df4389cc2e041500d487a3a3adeaeebb38e7409002a252e488848df
---

# fr-review-amd-epyc-9754s-review-a-cpu-with-a-very-particular-set-of-skills-6f36d356

L'année dernière, AMD a élargi sa gamme de processeurs pour serveurs avec l'EPYC de 4e génération. Alors que l'EPYC 128 à 256 cœurs et 9754 threads occupe la première place, juste en dessous de la matrice SKU se trouve l'AMD EPYC 9754S. La différence entre les deux puces est simple mais dramatique. Le 9754S a le multithreading simultané (SMT) désactivé. Cela signifie que le 9754S offre les mêmes 128 cœurs que le 9754, mais avec SMT désactivé, seulement 128 threads, contre 256. Ce changement entraîne une belle réduction pour les clients désactivant déjà SMT.
| Modèle | Coloré | Nombre maximum de fils | TDP par défaut | Fréq. (GHz) | Boost Fréq. (GHz) | Cache L3 (Mo) | 
|---|---|---|---|---|---|---|
| 9754 | 128 | 256 | 360W | 2.25 | 3.10 | 256 | 
| 9754S | 128 | 128 | 360W | 2.25 | 3.10 | 256 | 
| 9734 | 112 | 224 | 320W | 2.2 | 3.0 | 256 | 
Qu'est-ce qu'AMD SMT et pourquoi le 9754S existe-t-il ?
Avec SMT, un seul cœur de processeur EPYC peut traiter deux threads simultanément, ce qui peut conduire à une utilisation plus efficace des ressources du processeur. Lorsqu'un thread attend que les données soient chargées depuis la mémoire ou est inactif, l'autre thread peut exécuter des instructions. Cela signifie que le cœur passe moins de temps inactif, ce qui améliore potentiellement les performances. Cela est particulièrement vrai dans des cas d'utilisation tels que la virtualisation et le rendu.
La désactivation de SMT peut permettre aux fabricants de commercialiser ces puces en tant que produits de niveau inférieur, garantissant qu'elles répondent toujours à des critères de performances et de stabilité spécifiques. Les processeurs avec SMT désactivé peuvent être influencés par les processus de regroupement, les stratégies de segmentation du marché et le désir de répondre à des besoins spécifiques de performances ou d'efficacité, démontrant l'approche nuancée adoptée par les fabricants en matière de planification et de positionnement des produits.
Cela dit, toutes les charges de travail ne bénéficient pas du SMT et il arrive souvent que le SMT soit désactivé dans le BIOS sur un serveur AMD. Bien que cela puisse être un ajustement efficace, cela soulève un autre point important. La puce 9754S avec SMT désactivé est un peu moins chère que la 9754. Dans les deux cas, les applications monothread, les charges de travail informatiques et tous les cas d'utilisation dans lesquels la latence du processeur est d'une importance cruciale peuvent bénéficier de la désactivation de SMT.
Performances AMD EPYC 9754S et EPYC 9754
Nous souhaitons effectuer deux de nos tests réguliers, y-cruncher et Cinebench 2024, et voir quelles différences de performances nous obtenons avec et sans SMT. Nous avons comparé le 9754S et le 9754 tout en exécutant le 9754 avec SMT activé et désactivé pour voir quels avantages le 9754S présente sans SMT.
Plateforme de test et spécifications :
- TYAN Transport HX TN85-B8261
- 512GB DDR5
- Windows Server 2022
Cinebench 2024
Le premier est Cinebench 2024, avec SMT activé sur notre modèle non S. Ici, nous pouvons voir que nous sommes dans des différences de variation d’une séquence à l’autre.
| Processeur Cinebench 2024 | 2x EPYC9754S | 2x EPYC9754 | 
|---|---|---|
| Processeur multicœur | 2,682 | 2,587 | 
| Processeur monocœur | 68 | 69 | 
| Rapport PM | 39.19x | 37.64x | 
y-cruncher a été spécifiquement sélectionné en raison de l'architecture du programme, positionné comme un test complet du système. En effectuant un calcul Pi aussi important que possible dans la mémoire système, nous avons cherché à prouver notre intuition de longue date, selon laquelle SMT peut avoir un impact négatif sur les charges de travail liées au processeur et à la mémoire. Jetons d’abord un coup d’œil aux résultats avant de plonger dans ce que tout cela signifie.
y-cruncher 0.8.3
| y-cruncher 0.8.3 Temps de calcul total en secondes (plus bas c'est mieux) | 2x EPYC9754S | 2x EPYC 9754 (SMT désactivé) | 2x EPYC 9754 (SMT activé) | 9754 SMT désactivé Augmentation des performances | 
|---|---|---|---|---|
| 1 milliard | 13.481 | 13.546 | 14.139 | 4.65 % | 
| 2.5 milliard | 23.818 | 24.144 | 28.111 | 15.27 % | 
| 5 milliard | 40.760 | 40.797 | 49.271 | 17.27 % | 
| 10 milliard | 77.409 | 77.959 | 95.420 | 18.88 % | 
| 25 milliard | 203.303 | 202.124 | 233.629 | 12.98 % | 
| 50 milliard | 475.557 | 476.949 | 520.349 | 8.61 % | 
| 100 milliard | 1,248.458 | 1,251.36 | 1,242.419 | -0.49% | 
y-cruncher 0.8.4
| y-cruncher 0.8.4 Temps de calcul total en secondes (plus bas c'est mieux) | 2x EPYC9754S | 2x EPYC 9754 (SMT désactivé) | 2x EPYC 9754 (SMT activé) | 9754 SMT désactivé Augmentation des performances | 
|---|---|---|---|---|
| 1 milliard | 13.480 | 13.56 | 14.573 | 7.50 % | 
| 2.5 milliard | 23.680 | 23.501 | 28.649 | 17.34 % | 
| 5 milliard | 40.819 | 40.547 | 50.082 | 18.50 % | 
| 10 milliard | 78.523 | 77.466 | 93.842 | 16.32 % | 
| 25 milliard | 206.399 | 206.078 | 236.070 | 12.57 % | 
| 50 milliard | 483.797 | 482.79 | 521.867 | 7.29 % | 
| 100 milliard | 1,269.484 | 1,266.83 | 1,253.446 | -1.28% | 
L'analyse des résultats
En plongeant dans les subtilités d'AMD SMT, il existe un dialogue passionnant au sein de la communauté technologique sur ses implications sur les performances du système. À la base, SMT semble être un choix simple pour ceux qui recherchent des performances améliorées. La théorie est la suivante : si l'activation de SMT peut conduire à une mise à l'échelle idéale, alors pourquoi ne pas l'adopter comme un choix architectural bénéfique ?
La relation entre l’efficacité SMT et l’architecture de base n’est pas noire ou blanche. Une mise à l’échelle médiocre du SMT n’indique pas nécessairement un défaut dans sa mise en œuvre. En fait, cela pourrait faire allusion à une conception de base robuste qui ne laisse guère de place au SMT pour faire une différence notable. Ce paradoxe met en lumière un point crucial du secteur : les fabricants de processeurs ne peuvent pas revendiquer un avantage unique avec le SMT ou des technologies similaires. Ils reconnaissent que même si le SMT peut générer des performances supplémentaires dans certains cas d'utilisation, il n'est pas sans inconvénients dans d'autres scénarios.
À travers le prisme du calcul haute performance et des tâches de calcul intensif, les limites du SMT deviennent plus évidentes. Même si l’idée de doubler le nombre de threads par cœur peut sembler prometteuse, la réalité n’est pas comparable à un doublement du nombre de cœurs. Dans des cas extrêmes, cela peut entraîner une baisse des performances lorsque les threads se disputent les ressources du cache. Néanmoins, pour la majorité des applications multithread, en particulier celles dépourvues de concurrence en matière de cache, SMT améliore les performances, en particulier dans les tâches qui peuvent exploiter pleinement son potentiel.
Réflexions de clôture
AMD SMT est incroyablement utile pour une grande variété de charges de travail courantes dans l'entreprise. Mais toutes les charges de travail n’ont pas besoin ou ne bénéficient pas du SMT. Grâce à nos tests, nous avons montré comment AMD est capable de tirer parti des variations de fabrication pour proposer un produit solide doté d'une proposition de valeur unique. Les organisations qui conçoivent des plates-formes pour des types spécifiques de charges de travail nécessitant un cœur pur sans SMT peuvent économiser un peu d'argent en achetant l'AMD EPYC 9754S, dont le SMT est définitivement désactivé en usine.
