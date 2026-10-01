---
id: collect-261001-general-networking/general-networking/fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8-2
title: "fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "gpu", "llama"]
source: docs/RAG/collect-261001-general-networking/fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8.md
source_anchor: ""
source_lines: [64, 88]
sha256: 31d3e81be36315caead50ae8b099fb35e6ecf7b5fe58b8d225e0c41fd34c980e
---

# fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8

Le graphique ci-dessous illustre le comportement des disques lors du processus à travers 18 points de contrôle. Lors de l'entraînement de modèles d'apprentissage automatique, les points de contrôle sont essentiels pour sauvegarder régulièrement l'état du modèle et éviter toute perte de données en cas d'interruption ou de coupure de courant. Cette exigence de stockage requiert des performances robustes, notamment sous des charges de travail soutenues ou intensives. Nous avons utilisé la version 2.0 du benchmark DLIO, publiée le 13 août 2024.
Afin de garantir que nos tests de performance reflètent des scénarios réels, nous les avons basés sur l'architecture du modèle LLAMA 3.1 405B. Nous avons implémenté la sauvegarde des données à l'aide de `torch.save()` pour enregistrer les paramètres du modèle, l'état de l'optimiseur et l'état des couches. Notre configuration simulait un système à huit GPU, mettant en œuvre une stratégie de parallélisme hybride avec un parallélisme tensoriel à quatre voies et un traitement parallèle par pipeline à deux voies, répartis sur les huit GPU. Cette configuration a généré une taille de point de contrôle de 1 636 Go, reflétant les exigences de l'entraînement des grands modèles de langage modernes.
L'analyse des temps de passage révèle que le KIOXIA CD9P-R a débuté à 464.7 secondes lors du premier passage, avant de passer à 575.6 secondes au deuxième et de se stabiliser à 572.2 secondes au troisième. Ce comportement est très proche de celui de la majorité des autres SSD du groupe de comparaison, dont les temps de passage se sont situés entre 553 et 590 secondes environ. Le Pascari X200P fait figure d'exception, avec un temps nettement supérieur de 674.5 secondes. Globalement, le CD9P-R a démontré une évolution prévisible de ses performances lors d'opérations de points de contrôle répétées et est resté compétitif face aux SSD d'entreprise Gen5 les plus utilisés lors de ce test.
Lors du test DLIO Checkpoint Benchmark jusqu'au point de contrôle 12, le KIOXIA CD9P-R 7.68 To s'est avéré être l'un des disques les plus constants du groupe de comparaison. Après un démarrage à 471.4 secondes au premier point de contrôle, son temps de fonctionnement s'est stabilisé dans une plage relativement étroite, entre 560 et 580 secondes environ, pour atteindre 569.7 secondes au point de contrôle 12. Le disque KIOXIA a affiché des performances très proches de celles du Solidigm PS1010, du Micron 7600 MAX et du Kingston DC3000ME pendant la majeure partie de la charge de travail.
Le Pascari X200P s'est clairement démarqué, enregistrant une forte hausse de ses performances après le point de contrôle 4 et se maintenant largement au-dessus du lot, atteignant près de 690 secondes au point de contrôle 12. Le Micron 9550 MAX a affiché les temps de contrôle les plus courts et les plus constants durant la seconde moitié du test, descendant jusqu'à 531.3 secondes avant de terminer à 569.1 secondes. Bien que le CD9P-R n'ait pas été le disque le plus rapide à chaque point de contrôle, il a évité les fortes variations observées chez plusieurs concurrents et a offert des performances stables tout au long du test.
Benchmark de performance FIO
Pour mesurer les performances de stockage de chaque SSD selon les indicateurs courants du secteur, nous utilisons FIO. Chaque disque est soumis au même processus de test, qui comprend une étape de préconditionnement avec deux remplissages complets du disque avec une charge de travail d'écriture séquentielle, suivie d'une mesure des performances en régime permanent. À chaque modification du type de charge de travail mesuré, nous effectuons un nouveau remplissage de préconditionnement avec cette nouvelle taille de transfert.
Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Séquentiel 128K
- 64K Aléatoire
- 16K Aléatoire
- 4K Aléatoire
Écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Lors du test d'écriture séquentielle à 128 Ko en régime permanent, avec une profondeur d'E/S réduite à 16, le classement général est resté globalement inchangé par rapport à la période de préconditionnement. Le Micron 9550 Max (12.8 To) a conservé la tête avec 10 957,9 Mo/s, suivi de près par le Micron 9550 Pro (7.68 To) à 10 354,6 Mo/s. Le Kingston DC3000ME (7.68 To) s'est maintenu en troisième position à 8 477,4 Mo/s, et le Pascari X200P (7.68 To) a suivi de près avec 8 369,7 Mo/s.
Le KIOXIA CD9P-R (7.68 To) a atteint 6 912,4 Mo/s, se classant en dernière position. Le Solidigm PS1010 (7.68 To) à 7 126,5 Mo/s et le SanDisk DC SN861 (7.68 To) à 7 116,5 Mo/s sont tous deux en retrait par rapport aux disques de milieu de gamme, mais devancent tout de même le CD9P-R et le Micron 7600 Max (6.4 To) à 6 960,6 Mo/s. Ce résultat du KIOXIA est cohérent et prévisible pour un disque NVMe optimisé en lecture.
Latence d'écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Lors d'un test d'écriture en régime permanent avec une profondeur d'E/S de 16, la latence a considérablement diminué sur tous les disques par rapport aux conditions de préconditionnement. Le Micron 9550 Max (12.8 To) a de nouveau affiché la latence moyenne la plus faible (182.2 µs), devançant largement le Micron 9550 Pro (7.68 To) à 192.9 µs. Ces deux disques ont bénéficié de leur débit d'écriture supérieur, ce qui leur a permis de gérer les E/S plus efficacement.
Le disque dur KIOXIA CD9P-R (7.68 To) a enregistré une latence d'écriture de 289.0 µs, la plus élevée du groupe pour cette profondeur de file d'attente. Les disques Solidigm PS1010 (7.68 To) à 280.3 µs et SanDisk DC SN861 (7.68 To) à 280.7 µs se situaient juste en dessous du CD9P-R, tandis que le Micron 7600 Max (6.4 To) a atteint 287.1 µs. Les Kingston DC3000ME (7.68 To) et Pascari X200P (7.68 To) se situaient en milieu de classement avec des latences respectives de 235.6 µs et 238.6 µs.
Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Le test de lecture séquentielle à 128 Ko a complètement inversé le classement en écriture, et le KIOXIA CD9P-R (7.68 To) s'est révélé être l'un des plus performants du groupe. Le CD9P-R a atteint 14 235,9 Mo/s, égalant ainsi le Pascari X200P (7.68 To) à 14 242,1 Mo/s, en tête du classement. Le Solidigm PS1010 (7.68 To) à 14 163,3 Mo/s, le Micron 9550 Pro (7.68 To) à 14 050,1 Mo/s et le Micron 9550 Max (12.8 To) à 14 047,5 Mo/s se sont tous regroupés dans une bande de 200 Mo/s en tête.
Le Kingston DC3000ME (7.68 To) se classait en retrait par rapport aux leaders avec 13 513,8 Mo/s, tandis que le SanDisk DC SN861 (7.68 To) atteignait 12 631,2 Mo/s. Le Micron 7600 Max (6.4 To), avec 11 240,5 Mo/s, était le seul disque à descendre sous la barre des 12 Go/s.
Latence de lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Les résultats de latence en lecture séquentielle 128K reflètent fidèlement les performances en bande passante. Le Pascari X200P (7.68 To) arrive en tête avec 561.4 µs, tandis que le KIOXIA CD9P-R (7.68 To) affiche un résultat quasi identique de 561.7 µs. Le Solidigm PS1010 (7.68 To) à 564.5 µs, le Micron 9550 Pro (7.68 To) à 569.0 µs et le Micron 9550 Max (12.8 To) à 569.1 µs se situent tous dans une fourchette de 8 µs autour du disque le plus performant, confirmant ainsi que les performances de cette gamme de disques sont limitées par la bande passante de l'interface Gen5 plutôt que par une latence interne.
Le Kingston DC3000ME (7.68 To) a suivi à 591.6 µs et le SanDisk DC SN861 (7.68 To) à 633.0 µs, tandis que le Micron 7600 Max (6.4 To) à 711.4 µs a affiché une latence 26 % plus élevée que les meilleurs, ce qui correspond à son débit de lecture séquentiel inférieur.
Écriture aléatoire 64K
