---
id: collect-261001-general-networking/general-networking/fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a-2
title: "fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "gpu", "llama"]
source: docs/RAG/collect-261001-general-networking/fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a.md
source_anchor: ""
source_lines: [44, 74]
sha256: b400afb5fcdba4b4b043cf2befe330b67a79416b7627156d3f96d19214cbd649
---

# fr-review-solidigm-d7-ps1030-review-3-dwpd-gen5-that-earned-its-keep-in-the-kv-c-14d6508a

Afin de garantir que nos tests de performance reflètent des scénarios réels, nous les avons basés sur l'architecture du modèle LLAMA 3.1 405B. Nous avons implémenté la sauvegarde des données à l'aide de `torch.save()` pour enregistrer les paramètres du modèle, l'état de l'optimiseur et l'état des couches. Notre configuration simulait un système à huit GPU, mettant en œuvre une stratégie de parallélisme hybride avec un parallélisme tensoriel à quatre voies et un traitement parallèle par pipeline à deux voies, répartis sur les huit GPU. Cette configuration a généré une taille de point de contrôle de 1 636 Go, reflétant les exigences de l'entraînement des grands modèles de langage modernes.
| par chaîne | Passe 1 Moyenne (secondes) | Passe 2 Moyenne (secondes) | Passe 3 Moyenne (secondes) | 
|---|---|---|---|
| Disque dur SanDisk DC SN861 7.68 To | 461.3 | 558.6 | 553.3 | 
| Micron 9550 MAX 12.8 To | 462.8 | 558.9 | 555.3 | 
| Micron 9550 Pro 7.68 To | 461.4 | 577.9 | 559.7 | 
| Solidigm PS1010 7.68 To | 458.8 | 561.1 | 564.6 | 
| Micron 7600 MAX 6.4 To | 464.2 | 581.5 | 567.3 | 
| KIOXIA CD9P-R 7.68 To | 464.7 | 575.6 | 570.6 | 
| KIOXIA CM9-R 15.36 To | 462.8 | 571.9 | 580.9 | 
| Solidigm PS1030 12.8 To | 462.3 | 578.0 | 599.2 | 
Au vu des temps de passage moyens, le Solidigm PS1030 s'est classé parmi les meilleurs dès le premier passage, avec un temps de 462.3 secondes. L'ensemble des modèles comparés se situait alors dans une fourchette de sept secondes, entre 459 et 465 secondes environ. L'écart s'est creusé plus tard. Au deuxième passage, le PS1030 a atteint 578.0 secondes, se plaçant dans la partie supérieure d'une série allant de 558.6 secondes pour le SanDisk SN861 à 581.5 secondes pour le Micron 7600 MAX. Au troisième passage, il a grimpé à 599.2 secondes, la meilleure moyenne du groupe, les autres modèles se situant entre 553.3 secondes pour le SN861 et 580.9 secondes pour le KIOXIA CM9-R.
Le journal des points de contrôle du PS1030 illustre cette dérive. Le PS1030 a démarré à 465.3 secondes et s'est maintenu aux alentours de 461 secondes jusqu'au point de contrôle 5 avant d'accélérer. Une fois la transition amorcée, le temps de fonctionnement a oscillé entre 553 et 593 secondes environ, le point de contrôle 12 s'est clôturé à 614.7 secondes, et les derniers points de contrôle ont atteint 629.0 secondes. Ce résultat confirme une faiblesse connue du disque plutôt qu'une nouvelle : la mise en place de points de contrôle DLIO correspond précisément au type de pic d'écriture séquentielle important que le test FIO mono-processeur de 128 Ko a mis en évidence. L'écart est réel mais limité, d'environ 5 % par rapport au KIOXIA CD9P-R sur les moyennes du troisième passage, et les performances du PS1030 ont évolué de manière prévisible au fil des passages, sans fluctuations.
Benchmark de performance FIO
Pour mesurer les performances de stockage de chaque SSD selon les indicateurs courants du secteur, nous utilisons FIO. Chaque disque est soumis au même processus de test, qui comprend une étape de préconditionnement avec deux remplissages complets du disque avec une charge de travail d'écriture séquentielle, suivie d'une mesure des performances en régime permanent. À chaque modification du type de charge de travail mesuré, nous effectuons un nouveau remplissage de préconditionnement avec cette nouvelle taille de transfert.
Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Séquentiel 128K
- 64K Aléatoire
- Séquentiel 16K
- 4K Aléatoire
Écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Le test d'écriture séquentielle à 128 Ko en régime permanent est le seul point faible du PS1030 par rapport aux autres SSD. Avec une vitesse de 6 370,3 Mo/s et une latence de 313.7 µs, il affiche la bande passante la plus faible et la latence la plus élevée du groupe, devançant même le KIOXIA CD9P-R, pourtant très gourmand en lecture, à 6 912,4 Mo/s. Le Micron 9550 MAX arrive en tête avec 10 957,9 Mo/s, suivi du 9550 Pro à 10 354,6 Mo/s, tandis que le KIOXIA CM9-R se classe troisième avec 8 668,1 Mo/s. Le PS1010 (7 126,5 Mo/s), frère du PS1030, et le SanDisk DC SN861 (7 116,5 Mo/s) occupaient le milieu, avec le Micron 7600 MAX à 6 960,6 Mo/s.
Il est important de noter que cette charge de travail est monotâche et que l'architecture d'écriture du PS1030 est conçue pour répartir la charge, et non pour optimiser un flux de travail en particulier. Les sections d'écriture aléatoire ci-dessous montrent que le même disque déplace beaucoup plus de données une fois le parallélisme introduit, ce qui correspond au modèle d'accès généré par les charges de travail cibles.
Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
En lecture, la donne a changé. Le PS1030 a atteint 14 156,4 Mo/s à 564.8 µs, égalant quasiment son homologue PS1010 (14 163,3 Mo/s) et se classant à seulement 0.6 % du CD9P-R, leader du groupe (14 235,9 Mo/s). Les Micron 9550 Pro (14 050,1 Mo/s) et 9550 MAX (14 047,5 Mo/s) complètent le groupe de cinq disques, saturant l'interface Gen5 dans une bande de 200 Mo/s. Le SN861 suit à 12 631,2 Mo/s, le 7600 MAX à 11 240,5 Mo/s, et le CM9-R, pourtant si performant en écriture, ferme la marche avec 9 974,6 Mo/s dans cette configuration à tâche unique. Pour un disque vendu en fonction de son budget d'écriture, ne rien sacrifier en lecture sur gros blocs est le gain discret de ce tableau.
Écriture aléatoire 64K
C'est lors d'un balayage d'écriture aléatoire de 64 Ko que le PS1030 révèle tout son potentiel. Le disque a atteint un pic de 7 224,0 Mo/s, se classant quatrième de sa catégorie, derrière le Micron 9550 MAX (10 878,1 Mo/s), le KIOXIA CM9-R (9 635,3 Mo/s) et le Micron 9550 Pro (9 069,4 Mo/s), et devant le Micron 7600 MAX (6 960,5 Mo/s) et les autres disques testés. La particularité du PS1030 réside dans le contexte où ce pic a été atteint : avec une profondeur d'E/S (IODepth) de 2 et un nombre de tâches (NumJobs) de 2, la latence n'est que de 34.3 µs, alors que la plupart des autres disques nécessitaient des files d'attente importantes pour atteindre leurs performances optimales. Le disque sature presque instantanément, puis se stabilise pour le reste du balayage, ce qui correspond exactement au profil recherché pour une couche d'écriture en régime permanent, fonctionnant à une concurrence modérée 24 h/24. L'écart avec son homologue se traduit également par une meilleure endurance en termes de performances : le PS1010 a atteint un pic de 5 873,9 Mo/s, soit 23 % de moins que le PS1030.
Concernant la latence, le PS1030 a affiché 21.1 µs à une profondeur d'E/S de 1 et un nombre de tâches de 1, juste derrière le CM9-R et ses 18.4 µs. Son pic de latence le plus élevé sur l'ensemble du test a été de 2 831 µs, soit moins de la moitié du pic de 5 987 µs enregistré par le PS1010. Le 9550 MAX est resté le plus stable en cas de forte concurrence, avec une latence maximale de 1 714 µs.
Lecture aléatoire 64K
Le PS1030 a réalisé le meilleur pic de lecture aléatoire 64K du groupe à 14 162,9 Mo/s (IODepth 32 / NumJobs 8), devançant légèrement le Micron 9550 Pro (14 049,9 Mo/s), le 9550 MAX (14 049,7 Mo/s) et le PS1010 (14 013,6 Mo/s), tandis que le CM9-R à 13 402,4 Mo/s et le CD9P-R à 12 036,0 Mo/s étaient plus loin derrière. En matière de faible profondeur de file d'attente, les disques KIOXIA excellent, comme lors de notre test du CD9P-R : le CM9-R affichait 1 359,0 Mo/s et le CD9P-R 1 334,0 Mo/s à une profondeur d'E/S et un nombre de tâches de 1, soit une latence d'environ 45 µs. Le PS1030, quant à lui, démarrait à 768.4 Mo/s et 81.0 µs, se situant en milieu de classement. Le PS1030 se distingue par sa capacité d'adaptation, et non par sa réactivité en flux unique.
Écriture séquentielle 16K
