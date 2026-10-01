---
id: collect-261001-general-networking/general-networking/fr-review-micron-7600-max-review-9d96536b-4
title: "fr-review-micron-7600-max-review-9d96536b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-micron-7600-max-review-9d96536b.md
source_anchor: ""
source_lines: [107, 130]
sha256: 2ac0a28a09f2cf9c712e1a6f03c3f2f5381e2d27a625012ae90f9dc7d42c3b04
---

# fr-review-micron-7600-max-review-9d96536b

Lors du test de lecture séquentielle 16 K, le Micron 7600 MAX (6.4 To) a affiché une excellente constance, avec un débit initial de 1.03 Go/s, un pic à 11.0 Go/s et une moyenne de 6.08 Go/s sur l'ensemble du cycle. Sa forte évolutivité en milieu de gamme lui a permis de devancer légèrement le 9550 MAX en termes d'équilibre global et de performances soutenues.
Le Micron 9550 MAX (12.8 To) suivait de près, avec un débit initial de 1.02 Go/s, un pic à 12.5 Go/s et une moyenne de 5.59 Go/s. Bien qu'il ait atteint un débit absolu supérieur, sa courbe de performance présentait des fluctuations plus importantes selon la profondeur des files d'attente que les résultats plus stables du 7600 MAX.
Sur le graphique général, le Kingston DC3000ME a dominé les profondeurs de file d'attente les plus élevées, dépassant brièvement 12.8 Go/s, tandis que le Pascari X200P et le Solidigm PS1010 ont tous deux atteint les 12 Go/s. Le SanDisk DC SN861 a suivi de près, se maintenant juste en dessous de 10 Go/s en haut du classement.
Latence de lecture séquentielle de 16 K
Lors du test de latence de lecture séquentielle de 16 K, le Micron 7600 MAX (6.4 To) a démontré un contrôle de latence légèrement plus strict, commençant à 0.014 ms, culminant à 0.71 ms et atteignant une moyenne de 0.13 ms sur l'ensemble du cycle. Cela lui a conféré un léger avantage en termes de réactivité de lecture, maintenant une latence fluide et constante tout au long de la charge de travail.
Le Micron 9550 MAX (12.8 To) suivait de près avec des résultats allant de 0.015 ms en bas de gamme à 0.78 ms en haut de gamme, pour une moyenne globale de 0.15 ms. Bien que légèrement supérieures, ses performances restaient parmi les meilleures du secteur, affichant une excellente constance lors d'opérations de lecture séquentielles soutenues.
Sur le graphique général, les disques Kingston DC3000ME et Pascari X200P ont affiché des performances similaires en milieu de gamme, avec une moyenne comprise entre 0.1 et 0.2 ms, avec des pics légèrement supérieurs à 0.8 ms. Le Solidigm PS1010 a affiché une variabilité légèrement supérieure, plafonnant à près de 0.75 ms, tandis que le SanDisk DC SN861 a suivi de près le Kingston, mais a affiché des fluctuations plus importantes à mesure que la profondeur des files d'attente augmentait.
Écriture aléatoire 16K
Lors du test de lecture aléatoire de 16 000 octets, le Micron 7600 MAX (6.4 To) a affiché des performances constantes tout au long du cycle, allant de 17 000 IOPS en bas de la plage à environ 350 000 IOPS en moyenne, avec un pic à près de 720 000 IOPS aux profondeurs de file d'attente les plus élevées. Sa stabilité en a fait l'un des disques les plus prévisibles, assurant une mise à l'échelle fluide tout au long du cycle, même s'il n'a pas atteint le sommet du classement.
Le Micron 9550 MAX (12.8 To) a atteint un débit global supérieur, allant de 18 000 IOPS en bas de gamme à un pic légèrement supérieur à 900 000 IOPS, avec une moyenne d'environ 420 000 IOPS sur l'ensemble du cycle. Il a devancé les deux Micron en termes de performances brutes, mais a affiché une variation d'évolutivité légèrement supérieure à celle du 7600 MAX.
Sur le graphique général, le Pascari X200P et le Solidigm PS1010 ont tous deux enregistré de solides performances. Le Pascari a presque égalé le 9550 MAX en haut de la fourchette, plafonnant juste en dessous de 900 000 IOPS, tandis que le Solidigm s'est maintenu entre 820 000 et 850 000 IOPS. Le Kingston DC3000ME a initialement dominé, mais a plafonné à environ 620 000 IOPS, tandis que le SanDisk DC SN861 a suivi, atteignant un maximum légèrement supérieur à 500 000 IOPS.
Latence d'écriture aléatoire de 16 K
Lors du test de latence d'écriture aléatoire de 16 K, le Micron 7600 MAX (6.4 To) a conservé un profil compétitif et stable, allant de 0.016 ms en bas de la plage à 1.26 ms en pointe, avec une latence moyenne de 0.21 ms sur l'ensemble du cycle. Sa réactivité est restée excellente sous pression, affichant un contrôle constant sur différentes profondeurs de file d'attente, même s'il n'a pas atteint la même efficacité que son homologue de plus grande capacité.
Le Micron 9550 MAX (12.8 To) a démontré une discipline de latence globalement supérieure, restant entre 0.015 ms et 0.77 ms, avec une moyenne de 0.13 ms, consolidant sa position de disque le plus efficace sous une charge d'écriture soutenue.
D'après le graphique général, les disques Kingston DC3000ME et Pascari X200P se situent dans la moyenne, avec des temps de latence généralement compris entre 0.2 et 1.5 ms. Le SanDisk DC SN861 présente des pics de latence plus fréquents dans les files d'attente les plus longues, dépassant 1.8 ms, tandis que le Solidigm PS1010 présente les plus grandes difficultés à maintenir la stabilité, dépassant 3 ms aux points les plus critiques.
Lecture aléatoire 16K
Lors du test de lecture aléatoire de 16 000 octets, le Micron 7600 MAX (6.4 To) a affiché une excellente constance tout au long de l'exécution, commençant à 17 100 IOPS et atteignant 720 000 IOPS, avec une moyenne de 362 000 IOPS sur l'ensemble du cycle. Sa courbe de performance est restée régulière et prévisible, témoignant d'un excellent contrôle aux profondeurs de file d'attente basses et moyennes, même s'il n'a pas atteint le sommet du classement.
Le Micron 9550 MAX (12.8 To) a atteint un débit maximal plus élevé, allant de 16.7 000 IOPS en bas de gamme à 904 000 IOPS en haut de gamme, avec une moyenne de 433 000 IOPS. Il a dominé en termes de capacité d'évolutivité brute, bien que le 7600 MAX ait offert une meilleure cohérence sur toute la fenêtre de test.
Parmi ses concurrents, le Pascari X200P a quasiment égalé le 9550 MAX, atteignant un pic similaire de 900 000 IOPS. Le Solidigm PS1010 suivait de près avec une plage de 820 000 à 850 000 IOPS, tandis que le Kingston DC3000ME plafonnait autour de 620 000 IOPS. Le SanDisk DC SN861 complétait le peloton, terminant juste au-dessus de 500 000 IOPS et affichant une évolutivité limitée aux profondeurs de file d'attente plus élevées.
Latence de lecture aléatoire de 16 K
Lors du test de latence de lecture aléatoire de 16 K, le Micron 7600 MAX (6.4 To) a maintenu un profil de latence réactif et stable, commençant à 0.065 ms, culminant à 0.71 ms et atteignant une moyenne de 0.14 ms sur l'ensemble de l'exécution. Ses performances sont restées fluides et prévisibles tout au long du test, offrant une excellente réactivité même avec une augmentation de la profondeur des files d'attente.
Le Micron 9550 MAX (12.8 To) le talonnait de près, avec une latence comprise entre 0.073 ms en bas de gamme et 0.57 ms en haut de gamme, avec une latence moyenne de 0.12 ms. Si le disque plus grand présentait une courbe légèrement plus serrée aux profondeurs de file d'attente plus élevées, les deux modèles Micron ont fait preuve d'une cohérence et d'un contrôle de premier ordre.
Parmi les autres modèles, les disques Pascari X200P et Kingston DC3000ME se sont montrés compétitifs en milieu de gamme, avec des temps de latence compris entre 0.1 et 0.3 ms pendant la majeure partie du test, avant de grimper jusqu'à 0.8 ms en charge maximale. Les disques SanDisk DC SN861 et Solidigm PS1010 ont affiché une plus grande variabilité, Solidigm affichant notamment des pics de latence entre 0.6 et 0.65 ms, ce qui les place derrière les leaders Micron et Pascari en termes d'efficacité et de régularité.
Écriture aléatoire 4K
