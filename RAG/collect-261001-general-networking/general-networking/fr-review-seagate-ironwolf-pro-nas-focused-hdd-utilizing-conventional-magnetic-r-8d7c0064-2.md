---
id: collect-261001-general-networking/general-networking/fr-review-seagate-ironwolf-pro-nas-focused-hdd-utilizing-conventional-magnetic-r-8d7c0064-2
title: "fr-review-seagate-ironwolf-pro-nas-focused-hdd-utilizing-conventional-magnetic-r-8d7c0064"
domain: general-networking
role: reference
task: reference
actors: ["DeepSeek", "Meta"]
dates: []
keywords: ["benchmark", "benchmarks", "deepseek", "llama"]
source: docs/RAG/collect-261001-general-networking/fr-review-seagate-ironwolf-pro-nas-focused-hdd-utilizing-conventional-magnetic-r-8d7c0064.md
source_anchor: ""
source_lines: [61, 110]
sha256: 36fb6d7d0b1c17bf2609a2a7707638c8c19f3533a7ed902056e682f22082f980
---

# fr-review-seagate-ironwolf-pro-nas-focused-hdd-utilizing-conventional-magnetic-r-8d7c0064

| Seagate Iron Wolf Pro 30 To | 287 Mo/s (29.23 ms) | 267 Mo/s (31.39 ms) | 205 155.74 IOPS (XNUMX ms) | 301 105.95 IOPS (XNUMX ms) | 
| Seagate x24 24 To | 285 Mo/s (29.42 ms) | 285 Mo/s (29.42 ms) | 210 152.03 IOPS (XNUMX ms) | 749 42.70 IOPS (XNUMX ms) | 
| WD Or 24 To | 283 Mo/s (29.66 ms) | 286 Mo/s (29.36 ms) | 214 148.98 IOPS (XNUMX ms) | 651 49.11 IOPS (XNUMX ms) | 
| WD Red Pro 22 To | 271 Mo/s (31.00 ms) | 276 Mo/s (30.37 ms) | 214 149.17 IOPS (XNUMX ms) | 421 75.92 IOPS (XNUMX ms) | 
| WD Ultrastar DC HC590 | 268 Mo/s (31.28 ms) | 280 Mo/s (30.00 ms) | 198 161.18 IOPS (XNUMX ms) | 663 48.23 IOPS (XNUMX ms) | 
Temps de chargement moyen du LLM
Le test de temps de chargement moyen des LLM a évalué les temps de chargement de trois LLM différents : DeepSeek R1 7B, Meta Llama 3.2 11B et DeepSeek R1 32B. Chaque modèle a été testé 10 fois et le temps de chargement moyen a été calculé. Ce test mesure la capacité du lecteur à charger rapidement des modèles de langage volumineux (LLM) en mémoire. Les temps de chargement des LLM sont essentiels pour les tâches liées à l'IA, notamment pour l'inférence en temps réel et le traitement de grands ensembles de données. Un chargement plus rapide permet au modèle de traiter rapidement les données, améliorant ainsi la réactivité de l'IA et réduisant les temps d'attente.
Lors du test de temps de chargement moyen LLM, l'IronWolf 30 To s'est classé dans la moyenne des trois modèles testés. Il a chargé DeepSeek R1 7 To en 48.42 secondes et Meta Llama 3.2 11 To en 69.91 secondes, tous deux légèrement plus lents que l'Exos M et le WD Gold, mais devant des disques comme le x24 24 To et le WD Red Pro. Le modèle DeepSeek R1 32 To, plus grand, a terminé le chargement en 72.38 secondes, derrière le WD Gold et le HC590, mais proche de l'Exos M.
| Temps de chargement moyen du LLM (plus c'est bas, mieux c'est) | DeepSeek R1 7 milliard | Meta Llama 3.2 11B Vision | DeepSeek R1 32 milliard | 
| Seagate Exos 30 To | 46.4424s | 68.7064s | 72.7249s | 
| WD Or 24 To | 46.7133s | 68.8183s | 68.9720s | 
| WD Ultrastar DC HC590 26 To | 47.9877s | 71.0063s | 69.7892s | 
| Seagate Iron Wolf Pro 30 To | 48.4175s | 69.9071s | 72.3803s | 
| Seagate x24 24 To | 48.6615s | 71.4855s | 73.8097s | 
| WD Red Pro 22 To | 49.0575s | 71.4783s | 71.1382s | 
Stockage 3DMark
Le benchmark de stockage 3DMark teste les performances de jeu de votre SSD en mesurant des tâches telles que le chargement, la sauvegarde de la progression, l'installation de fichiers et l'enregistrement des parties. Il évalue la capacité de votre stockage à gérer les activités de jeu réelles et prend en charge les dernières technologies de stockage pour des analyses de performances précises.
Lors du benchmark de stockage 3DMark, l'IronWolf 30 To a obtenu un score de 231, se classant troisième au classement général et se rapprochant des deux meilleurs disques. Le x24 24 To a obtenu un score de 234, suivi de près par l'Exos M 30 To avec 223, les trois modèles Seagate se situant à quelques points près. Ce résultat souligne la constance des performances de l'IronWolf Pro, qui se rapproche de celles de ses homologues Seagate lors de simulations de charges de travail réelles.
| 3DMark Benchmark de stockage (plus c'est mieux) | Note globale | 
| Seagate x24 24 To | 234 | 
| Seagate Exos 30 To | 223 | 
| Seagate Iron Wolf Pro 30 To | 231 | 
| WD Ultrastar DC HC590 26 To | 168 | 
| WD Red Pro 22 To | 156 | 
| WD Or 24 To | 150 | 
Test de vitesse du disque BlackMagic
Le test de vitesse BlackMagic Disk évalue les vitesses de lecture et d'écriture d'un disque et évalue ses performances, notamment pour le montage vidéo. Il permet aux utilisateurs de s'assurer que leur stockage est suffisamment rapide pour les contenus haute résolution, comme les vidéos 4K ou 8K.
Lors du test de vitesse du disque BlackMagic, l'IronWolf 30 To a obtenu d'excellents résultats avec des vitesses de lecture de 267.6 Mo/s et d'écriture de 272.7 Mo/s. Bien qu'il soit légèrement inférieur à l'Exos M 30 To en lecture, il a conservé la deuxième vitesse d'écriture la plus élevée parmi tous les disques testés.
| Vitesse du disque BlackMagic (Mo/s, plus c'est élevé, mieux c'est) | Lire Mo/s | Écrire Mo/s | 
| Seagate Exos 30 To | 274.6 | 275.2 | 
| WD Or 24 To | 272.8 | 213.0 | 
| Seagate x24 24 To | 271.0 | 164.4 | 
| Seagate Iron Wolf Pro 30 To | 267.6 | 272.7 | 
| WD Ultrastar DC HC590 26 To | 267.0 | 264.5 | 
| WD Red Pro 22 To | 260.9 | 258.3 | 
Stockage PCMark 10
Les benchmarks de stockage PCMark 10 évaluent les performances de stockage en conditions réelles à l'aide de traces applicatives. Ils testent le système et les disques de données, en mesurant la bande passante, les temps d'accès et la cohérence sous charge. Ces benchmarks offrent des informations pratiques allant au-delà des tests synthétiques, permettant aux utilisateurs de comparer efficacement les solutions de stockage modernes.
Dans le benchmark PCMark 10 Data Drive, l'IronWolf 30 To a obtenu un score de 771, se classant deuxième au classement général, juste devant l'Exos M 30 To à 769. Il n'est devancé que par le WD Ultrastar DC HC590, qui est en tête avec 853. L'IronWolf Pro a surpassé le x24 24 To, le WD Gold et le WD Red Pro par une marge significative, démontrant de solides performances applicatives réelles et une réactivité constante sous charge.
| Lecteur de données PCMark 10 (plus c'est élevé, mieux c'est) | Note globale | 
| WD Ultrastar DC HC590 26 To | 853 | 
| Seagate Iron Wolf Pro 30 To | 771 | 
| Seagate Exos 30 To | 769 | 
| Seagate x24 24 To | 671 | 
| WD Or 24 To | 397 | 
| WD Red Pro 22 To | 380 | 
Conclusion
Le Seagate IronWolf Pro 30 To offre des performances solides et complètes, adaptées aux environnements NAS commerciaux et d'entreprise. Lancé en même temps que l'Exos M 30 To, il a constamment résisté aux tests, affichant un débit, une fiabilité et une réactivité comparables en conditions réelles. En charges de travail séquentielles, l'IronWolf Pro a légèrement distancé l'Exos M, mais a conservé une excellente régularité dans les opérations aléatoires, notamment en termes de latence de lecture et de réactivité.
Lors de tests de performance axés sur les applications, comme PCMark 10 et 3DMark Storage, l'IronWolf Pro a obtenu des résultats proches de ceux de l'Exos M et du x24 24 To, les trois disques terminant dans une courte marge. Ses vitesses d'écriture élevées au BlackMagic Disk Speed Test et ses temps de chargement LLM respectables démontrent sa polyvalence face à des charges de travail mixtes.
Conçu pour un fonctionnement continu, l'IronWolf Pro supporte une charge de travail de 550 To par an et un MTBF de 2.5 millions d'heures. Il intègre des fonctionnalités NAS comme le micrologiciel AgileArray et la résistance aux vibrations, ainsi qu'IronWolf Health Management pour une surveillance proactive. Les utilisateurs bénéficient également de trois ans de services de récupération de données Rescue inclus, offrant une garantie supplémentaire aux entreprises.
Pour les utilisateurs qui construisent des systèmes NAS multi-baies, des environnements de stockage collaboratifs ou des baies RAID haute capacité, l'IronWolf Pro 30 To offre un mélange fiable de capacité, de performances et d'endurance, correspondant à la valeur et à l'évolutivité attendues de la gamme NAS d'entreprise de Seagate.
