---
id: collect-261001-general-networking/general-networking/fr-review-micron-7600-max-review-9d96536b-2
title: "fr-review-micron-7600-max-review-9d96536b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "gpu", "llama"]
source: docs/RAG/collect-261001-general-networking/fr-review-micron-7600-max-review-9d96536b.md
source_anchor: ""
source_lines: [51, 77]
sha256: 052e466a993d4215bb0754c1910f6fa4dd58bf9ca12de84b9db2acd2fa6dcc5b
---

# fr-review-micron-7600-max-review-9d96536b

Pour évaluer les performances réelles des SSD dans les environnements d'entraînement d'IA, nous avons utilisé l'outil de référence DLIO (Data and Learning Input/Output). Développé par l'Argonne National Laboratory, DLIO est spécialement conçu pour tester les schémas d'E/S dans les charges de travail d'apprentissage profond. Il fournit des informations sur la façon dont les systèmes de stockage gèrent les défis tels que les points de contrôle, l'ingestion de données et l'entraînement des modèles. Le graphique ci-dessous illustre la gestion de ces processus par les deux disques SSD sur 36 points de contrôle. Lors de l'entraînement des modèles d'apprentissage automatique, les points de contrôle sont essentiels pour sauvegarder périodiquement l'état du modèle et éviter ainsi toute perte de progression en cas d'interruption ou de panne de courant. Cette demande de stockage exige des performances robustes, notamment sous des charges de travail soutenues ou intensives. Nous avons utilisé la version 2.0 du benchmark DLIO du 13 août 2024.
Afin de garantir que nos analyses comparatives reflètent des scénarios réels, nous avons basé nos tests sur l'architecture du modèle LLAMA 3.1 405B. Nous avons implémenté des points de contrôle à l'aide de torch.save() pour capturer les paramètres du modèle, les états de l'optimiseur et les états des couches. Notre configuration simulait un système à huit GPU, mettant en œuvre une stratégie de parallélisme hybride avec un parallélisme tensoriel à quatre voies et un traitement parallèle par pipeline à deux voies réparti sur les huit GPU. Cette configuration a généré des points de contrôle de 1 636 Go, ce qui reflète les exigences d'entraînement des modèles modernes de langages volumineux.
Lors de ce test, le Micron 9550 MAX 12.8 To s'est imposé comme le leader incontesté. Sur l'ensemble des 18 points de contrôle, il a conservé les temps d'exécution moyens les plus courts, compris entre 457 et 575 s. Le disque a offert une stabilité exceptionnelle avec une variation minimale entre les points de contrôle, témoignant d'une conception de micrologiciel bien équilibrée et optimisée pour les charges de travail mixtes en lecture/écriture.
Juste derrière, le Micron 7600 MAX 6.4 To a enregistré des temps compris entre 459 et 586 s. Bien que sa moyenne soit restée compétitive, le disque a présenté une brève fluctuation de performances entre les points de contrôle 4 et 7 avant de se stabiliser vers la fin du test. Malgré cela, il est resté solidement dans le peloton de tête, affichant une excellente efficacité pour les charges de travail soutenues en IA et HPC.
Le Micron 9550 7.68 To a obtenu des résultats juste derrière les deux modèles phares, avec des performances allant de 458 à 582. Il a maintenu une évolutivité constante et est resté compétitif face aux disques MAX haut de gamme, renforçant ainsi la puissance de la plateforme Micron 9550 sous-jacente.
Parmi les autres SSD d'entreprise testés, les Solidigm PS1010, SanDisk SN861 et Kingston DC3000ME se situaient en milieu de gamme, franchissant la plupart des points de contrôle entre 450 et 610 secondes. Le Pascari X200P a affiché les performances les moins constantes, atteignant plus de 690 secondes en milieu d'exécution avant de se stabiliser vers la fin.
Lors de ce test de réussite, le Solidigm PS1010 7.68 To a dominé le groupe avec les temps d'exécution moyens les plus rapides, allant de 458 à 564 secondes sur les trois passes. Le disque a fait preuve d'une excellente constance, maintenant une faible variabilité entre les exécutions et démontrant une grande efficacité sous des charges de travail d'E/S mixtes.
Le SanDisk SN861 7.68 To suivait de près, affichant des résultats presque identiques avec des moyennes comprises entre 461 s et 553 s, confirmant sa capacité à fournir des performances de point de contrôle fiables avec une dégradation minimale.
Le Micron 9550 7.68 To a suivi, terminant entre 461 et 559 secondes lors des mêmes passes. Ses performances sont restées très compétitives, se plaçant juste derrière les leaders, tout en conservant une évolutivité stable et un débit élevé à toutes les itérations.
Les disques Micron 9550 MAX 12.8 To et Micron 7600 MAX 6.4 To complètent le top 5, avec des moyennes légèrement supérieures, respectivement de 462 à 555 s et de 464 à 567 s. Leurs performances sont restées constantes dans le temps, mais ils restent derrière le disque Micron de plus petite capacité et les deux disques leaders de Solidigm et SanDisk.
Parmi les autres périphériques testés, les Kingston DC3000ME et Pascari X200P ont affiché les temps d'exécution les plus longs, avec des moyennes respectives de 580 s et 660 s. Ces résultats mettent en évidence un écart de performance plus important en conditions de sauvegarde continue, notamment pour les charges de travail nécessitant des écritures fréquentes sur le stockage persistant.
Benchmark de performance FIO
Pour mesurer les performances de stockage de chaque SSD selon les indicateurs courants du secteur, nous utilisons FIO. Chaque disque est soumis au même processus de test, qui comprend une étape de préconditionnement impliquant deux remplissages complets du disque avec une charge de travail d'écriture séquentielle, suivie d'une mesure des performances en régime permanent. À chaque modification du type de charge de travail mesuré, nous effectuons un nouveau remplissage de préconditionnement avec cette nouvelle taille de transfert.
Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Séquentiel 128K
- 64K Aléatoire
- 16K Aléatoire
- 16k séquentiel
- 4K Aléatoire
Écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Lors du test d'écriture séquentielle à 128 Ko, les résultats étaient quasiment identiques à ceux observés lors du préconditionnement. Le Micron 9550 Max (12.8 To) a une fois de plus largement dominé le classement, avec un débit de 10 957,9 Mo/s, conservant ainsi la première place. Le Kingston DC3000ME (7.68 To) suivait en deuxième position avec 8 477,4 Mo/s, suivi de près par le Pascari X200P (7.68 To) avec 8 369,7 Mo/s.
Plus loin derrière se trouvent le Solidigm PS1010 (7 126,5 Mo/s) et le SanDisk DC SN861 (7 116,5 Mo/s), tandis que le Micron 7600 Max (6.4 To) se situe en bas du classement avec 6 960,6 Mo/s.
Latence d'écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Concernant la latence, le test d'écriture séquentielle de 128 Ko a été exécuté avec une profondeur d'entrée-sortie de 16 et une seule tâche, contre une profondeur de file d'attente plus importante de 256 utilisée lors du préconditionnement. Comme prévu, la latence a diminué significativement sur tous les disques. Le Micron 9550 Max (12.8 To) a de nouveau dominé le marché avec la latence la plus faible, soit 0.18 ms, démontrant ainsi sa capacité à maintenir un débit maximal avec un délai minimal.
Le Kingston DC3000ME (7.68 To) suivait de près avec 0.24 ms, suivi de près par le Pascari X200P (7.68 To). Parallèlement, le Solidigm PS1010 (0.28 ms) et le SanDisk DC SN861 (0.28 ms) affichaient des résultats similaires, tandis que le Micron 7600 Max (6.4 To) arrivait en queue de peloton avec 0.29 ms.
Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
En ce qui concerne les lectures, le test de lecture séquentielle 128 Ko a donné des résultats bien plus proches de ceux des disques concurrents. Le Pascari X200P (7.68 To) a pris la tête avec 14 242,1 Mo/s, juste devant le Solidigm PS1010 (7.68 To) avec 14 163,3 Mo/s, et le Micron 9550 Max (12.8 To) juste derrière avec 14 047,5 Mo/s. Ces trois disques se situent dans une marge étroite, affichant des différences minimes en termes de débit de lecture séquentielle soutenue.
