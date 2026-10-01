---
id: collect-261001-general-networking/general-networking/fr-review-asus-ascent-qn10-review-550f5c4b-3
title: "fr-review-asus-ascent-qn10-review-550f5c4b"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["benchmark", "gpu", "intel", "valuation"]
source: docs/RAG/collect-261001-general-networking/fr-review-asus-ascent-qn10-review-550f5c4b.md
source_anchor: ""
source_lines: [55, 94]
sha256: 1ceb1dcda2694a15385ed42ecc6521deb301d2ae96ac3ed46afadc835e5c1b39
---

# fr-review-asus-ascent-qn10-review-550f5c4b

Face au Snapdragon X de génération précédente du Neo 50q, le QN10 affiche un score de 3 667 en monocœur contre 2 148 et de 19 884 en multicœur contre 8 565, soit un gain de 2.3 fois en multicœur en une seule génération Arm. La comparaison avec les processeurs x86 est tout aussi déséquilibrée : le Core Ultra 7 165H du NUC 14 Pro atteint 12 477 en multicœur lors de son nouveau test, soit 63 % du score du QN10. Le calcul GPU confirme cette tendance : l’Adreno obtient 39 761 en OpenCL contre 9 634 pour le Neo 50q et devance l’Arc du NUC (35 449).
Cinebench 2026
Cinebench 2026 est la version actuelle de la gamme Cinebench et la seule présentée ici. Ce test évalue les performances du processeur et de la carte graphique à l'aide du moteur de rendu Redshift de Maxon, basé sur le code source de Cinema 4D 2026. Il permet de vérifier la stabilité d'une machine sous forte charge processeur, la capacité de son système de refroidissement à supporter des rendus prolongés et ses performances face à des tâches 3D exigeantes. Les modifications apportées au code et au compilateur ayant accéléré le rendu des scènes, les scores de Cinebench 2026 utilisent une plage ajustée et ne doivent pas être comparés à ceux des versions précédentes. Le benchmark est fourni avec une version native ARM64, utilisée par le QN10. Le test GPU ne prenant pas en charge l'Adreno ni la carte graphique intégrée Intel du NUC, seuls les résultats du processeur sont présentés ici.
| Cinebench 2026 (plus c'est élevé, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Processeur à thread unique | 639 | 444 | N/D | 
| Processeur multithread | 6,476 | 3,684 | N/D | 
| Rapport PM | 10.13x | 8.29x | N/D | 
Avec un score de 639 en monocœur, le QN10 affiche le meilleur score Cinebench 2026 monocœur jamais enregistré sur un système de cette catégorie ; le tableau de référence de Maxon le situe entre l'Apple M4 Max et le processeur Intel Core Ultra 9 285K de bureau. En multicœur, il atteint 6 476 points, soit un ratio de 10.13x, contre 3 684 points (8.29x) pour le NUC 14 Pro, ce qui représente une avance de 76 %. L'écart en monocœur (639 contre 444) reste similaire, ce qui rend ces résultats remarquables pour un boîtier de 0.7 litre.
Profil de processeur 3DMark
Le test 3DMark CPU Profile mesure les performances du processeur à un nombre de cœurs constant, d'un seul cœur jusqu'au maximum disponible, et montre comment les performances évoluent avec l'utilisation de plusieurs cœurs. Plus le score est élevé, meilleures sont les performances. Ce test étant l'un de ceux que nous avions publiés pour le Neo 50q, nous l'avons réexécuté sur le QN10 afin de maintenir la comparaison.
| Profil CPU 3DMark (plus c'est élevé, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Nombre maximum de fils | 9,544 | 7,680 | 3,370 | 
| Fils 16 | 9,118 | 6,899 | 3,360 | 
| Fils 8 | 5,327 | 5,617 | 3,497 | 
| Fils 4 | 3,209 | 3,582 | 2,152 | 
| Fils 2 | 1,702 | 1,914 | 1,422 | 
| 1 discussion | 857 | 1,002 | 712 | 
L'écart générationnel est flagrant à chaque étape : 9 544 cœurs maximum contre 3 370 pour le Neo 50q, et 857 en monocœur contre 712. Le croisement à 8 cœurs est notable : 5 327 contre 3 497, car la plupart des charges de travail bureautiques se situent dans cette plage. La courbe s'aplatit également sensiblement entre 16 cœurs et le maximum sur le QN10, ce qui est normal pour un processeur 18 cœurs sans SMT. Face au NUC 14, le tableau est plus nuancé que ne le suggèrent les résultats Geekbench : le QN10 domine le haut de la courbe (9 544 contre 7 680), mais le 165H prend l'avantage en monocœur (1 002 contre 857) et conserve une légère avance jusqu'à 8 cœurs. Cette charge de travail favorise les fréquences d'horloge en rafale d'une manière que l'architecture Oryon ne prend pas en compte, ce qu'il est bon de garder à l'esprit pour les tâches de bureau peu gourmandes en ressources multithread.
Graphiques 3DMark
La suite graphique 3DMark couvre un large éventail de charges de travail de rendu : Solar Bay et Wild Life sont des tests multiplateformes conçus pour les cartes graphiques intégrées et mobiles, Steel Nomad et Time Spy sont des tests de rastérisation DirectX 12 plus exigeants, Fire Strike est le benchmark DirectX 11, et Port Royal et Speed Way testent le ray tracing DirectX. Un score élevé est synonyme de meilleures performances. Les résultats de Solar Bay, Steel Nomad et Wild Life recoupent nos données publiées pour le Neo 50q ; les autres tests établissent la base de référence du QN10 par rapport au NUC 14 Pro retesté, incluant les premiers tests de ray tracing DXR que nous avons réalisés sur un système Windows-on-Arm.
| Score graphique 3DMark (plus le score est élevé, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Baie solaire | 21,561 | 11,630 | 5,971 | 
| Nomade d'acier | 1,042 | 611 | 235 | 
| Faune | 19,996 | N/D | 11,224 | 
| Spy Time | 4,014 | 3,643 | N/D | 
| Grève de feu | 9,949 | 6,591 | N/D | 
| Port-Royal | 1,520 | 1,429 | N/D | 
| Voie de vitesse | 324 | 335 | N/D | 
Avec un score de 21 561 sur Solar Bay, le QN10 surpasse le Neo 50q (3.6 fois) et le NUC 14 (1.9 fois). Sur Steel Nomad (1 042), il dépasse la génération Arm précédente (4.4 fois) et l'Arc (70 %). Les scores de Time Spy (4 014) et Fire Strike (9 949) placent le QN10 dans la catégorie des GPU discrets d'entrée de gamme, et non en bas du classement des iGPU. Le ray tracing est le seul domaine où l'Arc reprend l'avantage : Port Royal l'emporte de peu sur le QN10 (1 520 contre 1 429), tandis que Speed Way l'emporte sur le NUC (335 contre 324), ce qui donne un résultat quasiment nul. Personne n'achète ces ordinateurs pour le ray tracing, mais la compatibilité avec l'architecture Arm est indéniable. Le test Wild Life n'est pas inclus dans la colonne NUC 14, car il n'a pas été testé lors de sa nouvelle évaluation.
Stockage 3DMark
3DMark Storage mesure les performances d'un SSD lors de tâches liées aux jeux, comme le chargement de jeux, l'installation de logiciels, la sauvegarde de la progression et le déplacement de fichiers. Ce test s'exécute nativement sur les plateformes Arm. Plus le score est élevé, meilleures sont les performances. Le score reflète le disque fourni avec chaque système ; il s'agit donc d'une comparaison des choix de stockage des constructeurs plutôt que des plateformes elles-mêmes.
| Stockage 3DMark (plus c'est élevé, mieux c'est) | ASUS Ascent QN10 | ASUS NUC 14 Pro (Core Ultra 7 165H) | Lenovo ThinkCentre Neo 50q QC | 
|---|---|---|---|
| Score de stockage | 2,393 | 1,887 | 2,871 | 
Avec un score de 2 393, le SSD SanDisk PC SN5100S de 512 Go du QN10 se situe en dessous des 2 871 du Neo 50q, mais devant les 1 887 du NUC 14. Le SN5100S est un disque client de classe Gen4 et s'installe dans l'emplacement Gen4 ; l'emplacement Gen5, étiqueté G5 sur la carte mère, est livré vide et est destiné à accueillir un disque plus rapide, puisque le SSD de 512 Go fourni est fonctionnel.
Compression à 7 zips
Le test de performance intégré de 7-Zip mesure la vitesse à laquelle le processeur peut compresser et décompresser des données à l'aide de plusieurs threads, avec un dictionnaire de 128 Mo et dix passes. La décompression est généralement proportionnelle au nombre de threads, tandis que la compression dépend de la latence mémoire ; les deux mesures donnent donc souvent des résultats différents. Un score GIPS plus élevé est préférable. 7-Zip 26.03 est fourni avec une version native ARM64, utilisée par le QN10.
