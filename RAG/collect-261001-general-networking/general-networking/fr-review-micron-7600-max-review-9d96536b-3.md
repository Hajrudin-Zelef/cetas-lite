---
id: collect-261001-general-networking/general-networking/fr-review-micron-7600-max-review-9d96536b-3
title: "fr-review-micron-7600-max-review-9d96536b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-micron-7600-max-review-9d96536b.md
source_anchor: ""
source_lines: [78, 106]
sha256: 82b9b397c38b9d95559a01fac950bf197051a458c0430384ca66206a581dd70c
---

# fr-review-micron-7600-max-review-9d96536b

Le Kingston DC3000ME (7.68 To) arrive légèrement en retrait du trio de tête avec 13 513,8 Mo/s, tandis que le SanDisk DC SN861 (7.68 To) atteint 12 631,2 Mo/s. En bas de gamme, le Micron 7600 Max (6.4 To) affiche 11 240,5 Mo/s, le seul disque du groupe à descendre sous la barre des 12 Go/s.
Latence de lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
En ce qui concerne la latence, le test de lecture séquentielle 128 Ko (IODepth 64 / NumJobs 1) a mis en évidence la forte concurrence entre les meilleurs. Le Pascari X200P (7.68 To) a dominé avec 0.56 ms, presque égalé par le Solidigm PS1010 (0.56 ms) et le Micron 9550 Max (12.8 To) avec 0.57 ms. Ces trois disques étaient à égalité, reflétant la faible différence de débit observée.
Le Kingston DC3000ME (7.68 To) suivait avec 0.59 ms, tandis que le SanDisk DC SN861 (7.68 To) atteignait 0.63 ms. Le Micron 7600 Max (6.4 To) arrivait dernier avec 0.71 ms, ce qui correspond à sa bande passante séquentielle plus faible.
Écriture aléatoire 64K
Lors du test d'écriture aléatoire de 64 Ko, le Micron 7600 MAX (6.4 To) a obtenu des résultats solides et constants, allant de 2.39 Go/s à 6.8 Go/s, avec un débit moyen de 5.16 Go/s sur l'ensemble du cycle. Il s'est ainsi positionné parmi les disques les plus performants, offrant une excellente stabilité tout au long du test et une évolutivité fiable même avec des files d'attente plus longues.
Le Micron 9550 MAX (12.8 To) est resté le leader incontesté du classement général, avec une plage de performances plus large, de 2.45 Go/s à un pic de 10.6 Go/s et une moyenne de 7.34 Go/s. Il a été le seul à dépasser régulièrement la barre des 10 Go/s, démontrant ainsi les avantages de sa configuration haut de gamme et du réglage optimisé de son micrologiciel.
Parmi les autres disques, les Kingston DC3000ME (7.68 To) et SanDisk DC SN861 (7.68 To) ont réalisé de solides performances dans la plage de 4 à 6 Go/s, restant compétitifs, même s'ils n'ont pas pu atteindre les performances supérieures de Micron. Les Solidigm PS1010 (7.68 To) et Pascari X200P (7.68 To) suivaient, se situant généralement dans la plage de 2 à 4 Go/s, loin derrière les deux disques Micron.
Latence d'écriture aléatoire de 64 K
En termes de latence, le Micron 7600 MAX (6.4 To) a maintenu un contrôle solide sous pression, avec une moyenne de 0.41 ms et un pic à 2.3 ms lors de files d'attente plus longues. Son profil de latence a démontré une réactivité constante tout au long du cycle, ce qui en fait l'un des disques les plus performants en écriture continue.
Le Micron 9550 MAX (12.8 To) est resté la référence en matière de cohérence, avec une moyenne de seulement 0.30 ms avec des pics inférieurs à 1.71 ms, affichant une gestion de la latence supérieure même à charge maximale.
Les Kingston DC3000ME et SanDisk DC SN861 se situaient dans la moyenne, avec des latences généralement comprises entre 0.05 ms et 2.7 ms, offrant un bon équilibre, mais n'atteignant pas la précision de Micron. En revanche, les Pascari X200P et Solidigm PS1010 affichaient la volatilité la plus importante, atteignant respectivement 4.1 ms et 6.0 ms à des profondeurs de file d'attente plus élevées.
Lecture aléatoire 64K
Lors du test de lecture aléatoire de 64 K, le Micron 7600 MAX (6.4 To) a affiché des performances équilibrées, avec un débit initial de 0.61 Go/s, un pic à 11.0 Go/s et une moyenne de 6.94 Go/s sur l'ensemble du cycle. La cohérence de ses lectures et sa mise à l'échelle régulière à des profondeurs de file d'attente plus élevées ont mis en évidence l'efficacité de son architecture et du réglage de son micrologiciel.
Le Micron 9550 MAX (12.8 To) a reproduit fidèlement ce comportement, avec des performances allant de 0.49 Go/s en bas de gamme à 13.7 Go/s, pour une moyenne globale de 6.96 Go/s. Les deux disques Micron se classent ainsi parmi les meilleurs en termes de performances, avec seulement quelques différences marginales entre eux.
Dans l'ensemble, les Solidigm PS1010 et Pascari X200P ont réussi à se démarquer légèrement en termes de débit maximal, atteignant 13-14 Go/s avec des profondeurs de file d'attente plus élevées. Le Kingston DC3000ME suivait de près avec 12 à 13 Go/s, tandis que le SanDisk DC SN861 suivait légèrement en dessous, se stabilisant autour de 12.3 Go/s.
Latence de lecture aléatoire de 64 K
Lors du test de lecture aléatoire de 64 K, le Micron 7600 MAX (6.4 To) a affiché une latence élevée, atteignant en moyenne 0.26 ms, puis 0.10 ms et atteignant un pic à 1.42 ms sous des charges plus importantes. Ses résultats ont montré une excellente constance tout au long du test, conservant une réactivité stable même avec une augmentation de la profondeur des files d'attente.
Le Micron 9550 MAX (12.8 To) a affiché des performances quasiment identiques, avec une moyenne de 0.25 ms, des creux à 0.12 ms et des pics jusqu'à 1.14 ms. Les deux disques Micron ont affiché une latence précise et prévisible, restant étroitement groupés et assurant un fonctionnement fluide tout au long du balayage.
En observant le graphique, les Solidigm PS1010 et Pascari X200P ont affiché des pics de latence légèrement supérieurs, généralement compris entre 0.1 et 1.2 ms. Parallèlement, les Kingston DC3000ME et SanDisk DC SN861 ont suivi dans une fourchette similaire, atteignant un pic juste au-dessus de 1.2 ms. Globalement, les disques Micron sont restés parmi les plus constants et les plus compétitifs du secteur, ne se différenciant que par de subtiles différences des autres disques de premier plan.
Écriture séquentielle 16K
Lors du test d'écriture séquentielle 16 K, le Micron 7600 MAX (6.4 To) a réalisé d'excellentes performances avec un débit compris entre 0.84 Go/s et 6.8 Go/s, pour une moyenne de 5.63 Go/s sur l'ensemble du cycle. Ses résultats ont montré un comportement d'écriture constant, avec une stabilité constante pour les profondeurs de file d'attente moyennes à élevées.
Le Micron 9550 MAX (12.8 To) a dominé sa catégorie, atteignant des débits compris entre 0.85 Go/s et 10.7 Go/s, avec un débit moyen de 7.75 Go/s. Il s'est imposé comme le leader incontesté, étant le seul disque capable de supporter des débits à deux chiffres en gigaoctets par seconde en période de pointe.
D'après le graphique général, les Kingston DC3000ME et Pascari X200P se situent dans la fourchette de 6 à 8 Go/s avec des profondeurs de file d'attente plus élevées, généralement compétitifs mais à la traîne par rapport au 9550 MAX. Le Solidigm PS1010 se situe légèrement en dessous, entre 5 et 6 Go/s, tandis que le SanDisk DC SN861 affiche les résultats les plus faibles, descendant fréquemment sous les 4 Go/s et atteignant des valeurs plancher proches de 1 Go/s.
Latence d'écriture séquentielle de 16 K
Lors du test de latence d'écriture séquentielle de 16 K, le Micron 7600 MAX (6.4 To) a fait preuve d'une excellente réactivité, avec une latence moyenne de 0.18 ms, un minimum de 0.018 ms et un pic de 1.15 ms sous charges plus importantes. Son profil de latence est resté stable tout au long du test, démontrant un contrôle d'écriture fiable quelle que soit la profondeur de la file d'attente.
Le Micron 9550 MAX (12.8 To) a offert la meilleure réactivité globale, avec une moyenne de 0.12 ms, atteignant des minimums de 0.018 ms et culminant à 0.75 ms sous charge, ce qui en fait le modèle le plus constant de cette catégorie.
D'après le graphique général, les disques Kingston DC3000ME et Pascari X200P se situaient en milieu de gamme, avec des temps de latence généralement compris entre 0.05 et 1.2 ms, tandis que le Solidigm PS1010 progressait, dépassant 1.5 ms dans les files d'attente les plus longues. Le SanDisk DC SN861 affichait la latence la plus élevée, dépassant 2.0 ms en conditions de forte charge.
Lecture séquentielle 16K
