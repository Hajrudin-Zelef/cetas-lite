---
id: collect-261001-huawei/huawei/fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217-3
title: "fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks"]
source: docs/RAG/collect-261001-huawei/fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217.md
source_anchor: ""
source_lines: [79, 111]
sha256: 6ea163fe6cf55af6fd4e88b9943d47dd55e10542c90becb1999731f2c5f89c34
---

# fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217

Cependant, à mi-parcours du benchmark (points de contrôle 5 à 9), les performances du X200P commencent à diverger. Les temps de passage aux points de contrôle augmentent fortement, culminant à 689.68 secondes au point de contrôle 12, le plus élevé du groupe. Sur les trois derniers points de contrôle, le X200P affiche une moyenne de 672 secondes, soit environ 19.3 % de moins que le disque le plus lent suivant (Kingston DC3000ME) et 23 % de moins que la moyenne du groupe.
Si l'on considère les moyennes de passes, le X200P montre une nette tendance à la dégradation des performances au fil du temps. Lors de la première passe, la durée moyenne de la course a été de 1 secondes, ce qui le place légèrement en retrait par rapport au reste du peloton, tout en restant compétitif.
Lors de la passe 2, le X200P s'est nettement démarqué du groupe, affichant un temps de 662.04 secondes, soit 14.5 % de moins que le disque le plus lent suivant (Kingston DC3000ME) et 17.4 % de moins que la moyenne du groupe pour cette passe. Cette tendance s'est poursuivie lors de la passe 3, où il a enregistré 674.48 secondes, conservant ainsi sa position de disque le plus lent. Comparé à la moyenne des quatre autres disques (environ 567 secondes), le X200P a été 18.9 % plus lent à terminer.
Benchmark de performance FIO
Pour mesurer les performances de stockage de chaque SSD selon les indicateurs courants du secteur, nous utilisons FIO. Chaque disque est soumis au même processus de test, qui comprend une étape de préconditionnement avec deux remplissages complets du disque avec une charge de travail d'écriture séquentielle, suivie d'une mesure des performances à l'état stable. À chaque changement de type de charge de travail mesuré, nous effectuons un nouveau remplissage de préconditionnement avec cette nouvelle taille de transfert.
Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Séquentiel 128K
- 64K Aléatoire
- 16K Aléatoire
- 4K Aléatoire
Précondition séquentielle de 128 K (IODepth 256 / NumJobs 1)
Lors du test de préconditionnement d'écriture séquentielle de 128 Ko, le X200P se classe troisième au classement général, avec une bande passante moyenne de 8,371 9550 Mo/s. Bien qu'il conserve d'excellentes performances, le lecteur présente une légère fluctuation récurrente de la bande passante, indiquant une moins bonne régularité par rapport aux courbes plus plates et plus stables du Micron 3000 et du Kingston DCXNUMXME, situés au-dessus.
Latence de précondition séquentielle de 128 K (IODepth 256 / NumJobs 1)
Lors du test de latence préalable à l'écriture séquentielle de 128 Ko, le X200P a enregistré une latence moyenne de 3.822 ms, ce qui le place troisième au classement général. Il se situe derrière le Micron 9550 et le Kingston DC3000ME. À l'instar de sa bande passante, le Pascari présente de légères fluctuations de latence, indiquant une certaine variabilité lors des écritures soutenues, tout en conservant une position solide dans le haut de gamme.
Écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Lors du test d'écriture séquentielle 128 Ko, le X200P a atteint une bande passante moyenne de 8369.7 9550 Mo/s, ce qui le place troisième au classement général. Il se classe derrière le Micron 3000 et le Kingston DC1010ME, mais devant le Solidigm PS861 et le SanDisk SNXNUMX.
Latence d'écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Lors du test d'écriture séquentielle 128K, le X200P a enregistré une latence moyenne de 0.238 ms. Il se classe ainsi quatrième au classement général, juste derrière le Kingston DC3000ME (0.235 ms) et devant le Solidigm PS1010 et le SanDisk SN861. Bien que sa latence soit inférieure à la plupart des autres, il reste inférieur au Micron 9550, le plus performant.
Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Lors du test de lecture séquentielle 128K, le X200P s'est classé premier au classement général avec une bande passante de 14,242.1 1010 Mo/s, devançant de justesse le Solidigm PS9550 et le Micron XNUMX. Il est en tête du peloton en termes de débit de lecture, affichant d'excellentes performances à une profondeur de file d'attente élevée.
Latence de lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Lors du test de latence de lecture séquentielle 128K, le X200P a enregistré une latence moyenne de 561.4 ms, ce qui le place deuxième en termes de latence globale. Il devance de peu le Solidigm PS1010 et surpasse le Micron 9550, le Kingston DC3000ME et le SanDisk SN861.
Écriture aléatoire 64K
Lors du test d'écriture aléatoire 64 Ko, le X200P affiche des performances moyennes, avec quelques fluctuations selon la profondeur des files d'attente et la combinaison de threads. Le disque ne présente pas de performances dans une plage spécifique, mais maintient des performances généralement stables entre 2,500 3,600 Mo/s et 6,625.92 32 Mo/s sur la majeure partie du test. Sa bande passante maximale atteint 8 XNUMX Mo/s avec la combinaison XNUMX/XNUMX IODepth/NumJobs, ce qui le place parmi les meilleurs résultats du test et lui confère une excellente note finale.
Bien qu'il ne soit pas le plus performant, le lecteur Pascari tient bon dans les charges de threads plus lourdes et fonctionne mieux vers les profondeurs de file d'attente plus élevées.
Latence d'écriture aléatoire de 64 K
Lors du test de latence d'écriture aléatoire de 64 K, le X200P affiche généralement une faible latence avec des profondeurs de file d'attente faibles à modérées, avec des valeurs exceptionnelles de 0.023 ms à 1/1 et de 0.041 ms à 2/1. Cependant, avec des combinaisons de threads et de files d'attente plus importantes, telles que 16/8 et 8/8, la latence augmente considérablement, atteignant respectivement 4.045 ms et 3.019 ms.
Lecture aléatoire 64K
Lors du test de lecture aléatoire 64 K, le X200P affiche des performances constantes sur toute la plage de profondeurs de file d'attente et de nombres de threads, suivant de près les meilleurs disques. Bien qu'il ne remporte pas la première place au départ, il domine aux QD 16/8 et 32/8, atteignant une bande passante maximale de 14,232 XNUMX Mo/s, égalant ou surpassant la concurrence aux niveaux de charge les plus élevés. Cela démontre la capacité du disque à évoluer sous un accès parallèle intensif.
Latence de lecture aléatoire de 64 K
Lors du test de latence de lecture aléatoire 64 K, le X200P maintient une latence faible et constante, quelle que soit la profondeur de file d'attente et le nombre de threads, généralement inférieure à 0.2 milliseconde. La latence commence à augmenter sensiblement à QD16/4, atteignant 0.285 ms, puis grimpe encore à QD32/4, atteignant 0.563 ms. L'augmentation la plus significative se produit à QD32/8, où la latence culmine à 1.135 ms.
Écriture aléatoire 16K
Lors du test d'écriture aléatoire de 16 200, le X170P conserve une solide position intermédiaire dans la plupart des combinaisons de files d'attente et de threads, avec des performances comprises entre 190 4 et 4 8 IOPS dans des configurations classiques comme 4/4, 8/221 et 32/8. Les performances commencent à évoluer significativement sous des charges plus importantes, atteignant 413 32 IOPS à 16/3000 et culminant à 428 XNUMX IOPS à XNUMX/XNUMX. À ce stade, il termine juste derrière le Kingston DCXNUMXME (XNUMX XNUMX IOPS), qui devance de peu. Ce bon résultat démontre la capacité du Pascari à évoluer efficacement sous une pression d'écriture maximale, même s'il n'occupe pas la première place.
