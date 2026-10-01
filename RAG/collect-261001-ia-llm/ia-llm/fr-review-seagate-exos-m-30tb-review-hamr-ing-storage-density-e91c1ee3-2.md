---
id: collect-261001-ia-llm/ia-llm/fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3-2
title: "fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Meta"]
dates: []
keywords: ["benchmark", "deepseek", "llama"]
source: docs/RAG/collect-261001-ia-llm/fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3.md
source_anchor: ""
source_lines: [59, 104]
sha256: 73354985161169fe4b1e0882903d5e47733c9cc8be40fb5447a3255589e004d8
---

# fr-review-seagate-exos-m-30tb-review-hamr-ing-storage-density-e91c1ee3

Le test FIO est un outil d'analyse comparative flexible et puissant permettant de mesurer les performances des périphériques de stockage, notamment les SSD et les disques durs. Il évalue des indicateurs tels que la bande passante, les IOPS (opérations d'entrée/sortie par seconde) et la latence sous différentes charges de travail, comme les opérations de lecture/écriture séquentielles et aléatoires. Ce test permet d'évaluer les performances maximales des systèmes de stockage, ce qui le rend utile pour comparer différents périphériques ou configurations. Nous avons mesuré les performances maximales en rafale pour ce test, en limitant la charge de travail à 10 Go sur les deux disques durs.
Charges de travail séquentielles (bloc de 128 Ko, 1 thread, profondeur de file d'attente de 64)
Lors des tests de lecture et d'écriture séquentielles, le Seagate Exos M 30 To a enregistré des débits de 292 Mo/s en lecture et de 289 Mo/s en écriture, ce qui le place en tête du classement pour les deux opérations. Ces résultats reflètent la capacité du disque à maximiser le débit soutenu, un facteur essentiel pour les centres de données gérant des transferts séquentiels volumineux, tels que la sauvegarde et l'archivage. Le WD Gold 24 To et le Seagate x24 24 To suivent de près, oscillant tous deux entre 283 et 285 Mo/s.
Les disques durs plus anciens, comme le WD Red Pro 22 To et l'Ultrastar HC590, ont affiché des performances nettement inférieures pour cette charge de travail, tombant sous les 275 Mo/s. Malgré une capacité nettement supérieure, l'Exos M 30 To a conservé sa position de leader en matière de transferts séquentiels sans latence supplémentaire, signe d'une architecture interne optimisée pour les tâches gourmandes en débit.
Charges de travail aléatoires de 4 K (16 threads, 32 files d'attente)
En E/S aléatoires de petits blocs, les performances révèlent souvent la capacité d'un disque à gérer des charges de travail gourmandes en métadonnées ou des applications à forte concurrence. L'Exos M 30 To a atteint 205 IOPS en lecture aléatoire (155.58 ms) et 341 IOPS en écriture aléatoire (93.79 ms).
Les performances en lecture et en écriture par seconde (IOPS) étaient compétitives sur tous les disques, les WD Gold et WD Red Pro se classant légèrement en tête avec 214 IOPS. En revanche, les performances en écriture de l'Exos M étaient plus significatives, surpassant celles de l'IronWolf Pro (301 IOPS) et nettement celles de l'Ultrastar HC590 (663 IOPS) en termes de latence, sans toutefois égaler celles du x24 24 To, qui a atteint 749 IOPS avec une latence bien inférieure de 42.70 ms.
| Test FIO (un débit MB/s/IOPS plus élevé est meilleur) | Lecture séquentielle 128K (1T/64Q) | Écriture séquentielle 128 Ko (1T/64Q) | Lecture 4K aléatoire (16T/32Q) | Écriture 4K aléatoire (16T/32Q) | 
| Seagate Exos 30 To | 292 Mo/s (28.72 ms) | 289 Mo/s (29.04 ms) | 205 155.58 IOPS (XNUMX ms) | 341 93.79 IOPS (XNUMX ms) | 
| Seagate Iron Wolf Pro 30 To | 287 Mo/s (29.23 ms) | 267 Mo/s (31.39 ms) | 205 155.74 IOPS (XNUMX ms) | 301 105.95 IOPS (XNUMX ms) | 
| Seagate x24 24 To | 285 Mo/s (29.42 ms) | 285 Mo/s (29.42 ms) | 210 152.03 IOPS (XNUMX ms) | 749 42.70 IOPS (XNUMX ms) | 
| WD Or 24 To | 283 Mo/s (29.66 ms) | 286 Mo/s (29.36 ms) | 214 148.98 IOPS (XNUMX ms) | 651 49.11 IOPS (XNUMX ms) | 
| WD Red Pro 22 To | 271 Mo/s (31.00 ms) | 276 Mo/s (30.37 ms) | 214 149.17 IOPS (XNUMX ms) | 421 75.92 IOPS (XNUMX ms) | 
| WD Ultrastar DC HC590 26 To | 268 Mo/s (31.28 ms) | 280 Mo/s (30.00 ms) | 198 161.18 IOPS (XNUMX ms) | 663 48.23 IOPS (XNUMX ms) | 
Temps de chargement moyen du LLM
Le test de temps de chargement moyen des LLM a évalué les temps de chargement de trois LLM différents : DeepSeek R1 7B, Meta Llama 3.2 11B et DeepSeek R1 32B. Chaque modèle a été testé 10 fois et le temps de chargement moyen a été calculé. Ce test mesure la capacité du lecteur à charger rapidement des modèles de langage volumineux (LLM) en mémoire. Les temps de chargement des LLM sont essentiels pour les tâches liées à l'IA, notamment pour l'inférence en temps réel et le traitement de grands ensembles de données. Un chargement plus rapide permet au modèle de traiter rapidement les données, améliorant ainsi la réactivité de l'IA et réduisant les temps d'attente.
Le Seagate Exos M 30 To a enregistré les temps de chargement les plus rapides sur deux des trois modèles AI testés, surpassant tous les autres disques sur DeepSeek R1 7B et Meta Llama 3.2 11B. Il a devancé le WD Gold 24 To par une marge faible mais constante, démontrant ainsi sa capacité à maintenir un débit élevé sur des charges de plusieurs gigaoctets.
Bien que l'Exos M ne soit pas en tête du classement des modèles 32B, il reste proche des meilleurs lecteurs comme le WD Gold et l'Ultrastar HC590. Cette constance entre les différentes tailles de modèles démontre que l'Exos M maintient une faible latence et d'excellentes performances de lecture, même dans des conditions exigeantes et à volume élevé.
| Temps de chargement moyen du LLM (plus c'est bas, mieux c'est) | DeepSeek R1 7 milliard | Meta Llama 3.2 11B Vision | DeepSeek R1 32 milliard | 
| Seagate Exos 30 To | 46.4424s | 68.7064s | 72.7249s | 
| WD Or 24 To | 46.7133s | 68.8183s | 68.9720s | 
| WD Ultrastar DC HC590 26 To | 47.9877s | 71.0063s | 69.7892s | 
| Seagate Iron Wolf Pro 30 To | 48.4175s | 69.9071s | 72.3803s | 
| Seagate x24 24 To | 48.6615s | 71.4855s | 73.8097s | 
| WD Red Pro 22 To | 49.0575s | 71.4783s | 71.1382s | 
Stockage 3DMark
Le benchmark de stockage 3DMark teste les performances de jeu de votre SSD en mesurant des tâches telles que le chargement, la sauvegarde de la progression, l'installation de fichiers et l'enregistrement des parties. Il évalue la capacité de votre stockage à gérer les activités de jeu réelles et prend en charge les dernières technologies de stockage pour des analyses de performances précises.
Dans le benchmark de stockage 3DMark, le Seagate Exos M 30 To a obtenu un score de 223, se classant deuxième au classement général, juste un point derrière le Seagate x24 24 To avec un score de 234 et devant l'IronWolf Pro 30 To avec un score de 231. De plus, il s'est placé devant le WD Ultrastar DC HC590 (168) et le WD Red Pro 22 To (156).
| 3DMark Benchmark de stockage (plus c'est mieux) | Note globale | 
| Seagate x24 24 To | 234 | 
| Seagate Exos 30 To | 223 | 
| Seagate Iron Wolf Pro 30 To | 231 | 
| WD Ultrastar DC HC590 26 To | 168 | 
| WD Red Pro 22 To | 156 | 
| WD Or 24 To | 150 | 
Test de vitesse du disque BlackMagic
Le test de vitesse BlackMagic Disk évalue les vitesses de lecture et d'écriture d'un disque et évalue ses performances, notamment pour le montage vidéo. Il permet aux utilisateurs de s'assurer que leur stockage est suffisamment rapide pour les contenus haute résolution, comme les vidéos 4K ou 8K.
Lors du test de vitesse de disque, le Seagate Exos M 30 To a enregistré 274.6 Mo/s en lecture et 275.2 Mo/s en écriture, se classant en tête dans les deux catégories. Il a surpassé tous les autres disques, y compris les WD Gold 24 To et IronWolf Pro 30 To, qui étaient à la traîne en écriture.
| Vitesse du disque BlackMagic (Mo/s, plus c'est élevé, mieux c'est) | Lire Mo/s | Écrire Mo/s | 
| Seagate Exos 30 To | 274.6 | 275.2 | 
| WD Or 24 To | 272.8 | 213.0 | 
| Seagate x24 24 To | 271.0 | 164.4 | 
| Seagate Iron Wolf Pro 30 To | 267.6 | 272.7 | 
| WD Ultrastar DC HC590 26 To | 267.0 | 264.5 | 
| WD Red Pro 22 To | 260.9 | 258.3 | 
Stockage PCMark 10
