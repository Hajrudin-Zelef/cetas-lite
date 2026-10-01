---
id: collect-261001-general-networking/general-networking/fr-review-kingston-dc3000me-review-built-for-system-integrators-and-mainstream-e-ebcbb957-4
title: "fr-review-kingston-dc3000me-review-built-for-system-integrators-and-mainstream-e-ebcbb957"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-kingston-dc3000me-review-built-for-system-integrators-and-mainstream-e-ebcbb957.md
source_anchor: ""
source_lines: [89, 115]
sha256: 093f30eced611fd99478c257f26fef516a6a59c7851bbbe3736b89d0edbe9d12
---

# fr-review-kingston-dc3000me-review-built-for-system-integrators-and-mainstream-e-ebcbb957

Latence d'écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Lors du test de latence d'écriture séquentielle 128 Ko, le Kingston DC3000ME a obtenu un excellent résultat avec une latence moyenne de 235.6 µs. Il devance ainsi le SanDisk SN861 et le Solidigm PS1010, qui affichent respectivement des latences de 280.7 µs et 280.3 µs. Bien que moins rapide que le Micron 9550, qui domine avec 192.9 µs, le Kingston DC3000ME reste compétitif et devance légèrement le Pascari X200P avec 238.6 µs.
Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Lors du test de lecture séquentielle de 128 Ko avec une profondeur de file d'attente de 64, le Kingston DC3000ME a atteint 13,513.8 200 Mo/s. Bien que ce résultat le place en quatrième position parmi les disques testés, il a néanmoins fourni un débit élevé (avec des différences minimes en conditions réelles). Il était inférieur d'environ 14,242.1 % au Pascari X5.1P (1010 14,163.3 Mo/s), au Solidigm PS4.6 (9550 14,050.1 Mo/s) de 3.8 % et au Micron 861 (12,631.2 3000 Mo/s) d'environ XNUMX %. Il a également largement surpassé le SanDisk SNXNUMX, qui a atteint une vitesse inférieure de XNUMX XNUMX Mo/s. Néanmoins, les résultats du Kingston DCXNUMXME étaient excellents, avec une baisse minime par rapport aux disques les mieux testés.
Latence de lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
En termes de latence, le Kingston DC3000ME a enregistré une latence moyenne de 591.6 µs, ce qui le place en milieu de groupe. Ce résultat est supérieur de 5.4 % à celui du Micron 9550 (569.0 µs) et inférieur de 5.4 % à celui du Solidigm PS1010 (564.5 µs). Le Pascari X200P arrive légèrement en tête avec 561.4 µs, tandis que le SanDisk SN861 affiche la réponse la plus lente avec 633.0 µs. Au final, le Kingston DC3000ME a maintenu une latence relativement faible dans des conditions de lecture avec une profondeur de file d'attente élevée.
Écriture aléatoire 64K
Lors du test d'écriture aléatoire de 64 Ko, le Kingston DC3000ME a constamment affiché des performances élevées sur différentes profondeurs de file d'attente et combinaisons de threads, atteignant un pic de 6,649 32 Mo/s dans la configuration 8 (profondeur d'E/S)/XNUMX (nombre de tâches). Ce résultat figure parmi les plus élevés, toutes charges de travail et points de test confondus.
Tout au long du graphique, le Kingston DC3000ME a maintenu une bande passante stable entre 4,000 5,000 et 32 4 Mo/s, avec des performances particulièrement élevées dans les configurations de concurrence moyenne à élevée, telles que 5,380/16 (8 5,017 Mo/s) et 1/4 (2 4 Mo/s). Même dans des conditions plus légères, comme 4200/3000 et XNUMX/XNUMX, il s'est maintenu au-dessus de XNUMX XNUMX Mo/s. Comparé aux autres disques, le Kingston DCXNUMXME a généralement dominé ou est resté proche du sommet du groupe sur la plupart des points de test, offrant à la fois un débit de pointe élevé et des performances constantes tout au long du test.
Latence d'écriture aléatoire de 64 K
Lors du test de latence d'écriture aléatoire de 64 K, le Kingston DC3000ME a systématiquement affiché des temps de réponse faibles pour la plupart des profondeurs de file d'attente et des combinaisons de tâches, démontrant ainsi une excellente efficacité d'écriture, même sous forte charge. Par exemple :
- À 4/1, il affichait 49µs
- À 8/1, la latence est restée faible à 102 µs
- À 16/4, il mesurait 1,486 XNUMX µs
- Et à la charge testée la plus élevée, 32/8, elle a atteint 2,402 XNUMX µs
Ces résultats indiquent que le Kingston DC3000ME a évolué de manière prévisible, évitant les pics de latence importants observés sur d'autres disques, en particulier les disques Pascari et Solidigm, qui ont présenté des sauts erratiques au-dessus de 3,000 6,000 à 16 8 µs (notamment à XNUMX/XNUMX).
Lecture aléatoire 64K
Lors du test de lecture aléatoire 64 K, le Kingston DC3000ME a affiché des performances solides et constantes sur l'ensemble de la matrice IOdepth/NumJobs, terminant quatrième à la fin du test (avec une faible avance). La bande passante maximale a atteint 13,515 32 Mo/s à 4/16, avec un débit tout aussi élevé à 4/13,482 (32 8 Mo/s) et 13,512/1 (4 2 Mo/s), démontrant une excellente évolutivité sous de lourdes charges de lecture parallèle. À des charges plus faibles, comme 2/3000 et 2,298/2,234, le Kingston DCXNUMXME a enregistré des débits respectifs de XNUMX XNUMX Mo/s et XNUMX XNUMX Mo/s.
Latence de lecture aléatoire de 64 K
La latence du Kingston DC3000ME 64K est restée relativement faible à tous les points de test. Tous les disques ont affiché des performances similaires, même si le SanDisk SN861 a atteint un pic nettement supérieur à celui de tous les autres disques testés à la fin du test. À partir de 1/2, le Kingston DC3000ME a enregistré 106 µs, suivi de 108 µs à 1/4, 131 µs à 8/1, 133 µs à 4/4 et 177 µs à 8/4. À une concurrence plus élevée, elle est passée à 305 µs à 16/4, 174 µs à 32/1, 301 µs à 32/2, pour atteindre un pic à 1,184 32 µs à 8/3000, ce qui le place dans la même catégorie que le reste du groupe. Dans l'ensemble, le profil de latence du Kingston DCXNUMXME a suivi de près celui des meilleurs disques, avec une gigue minimale ou des pics aberrants (ce qui était le cas de tous les disques testés).
Écriture aléatoire 16K
Lors du test d'écriture aléatoire de 16 Ko, le Kingston DC3000ME a fourni une bande passante élevée sur toute la plage de profondeurs de file d'attente et de nombres de threads, terminant le test à la deuxième place parmi les disques concurrents. Il a atteint 427,592 32 IOPS en configuration 16/338,521. Parmi les autres performances élevées, on compte 32 8 IOPS en configuration 251,428/16, 4 226,606 IOPS en configuration 1/8 et 2 16 IOPS en configuration 1/4, tous démontrant une excellente efficacité du contrôleur sous différentes charges parallèles. Même dans des configurations à charge modérée, comme 218,300/204,867 et 3000/160,000, le disque a atteint respectivement XNUMX XNUMX IOPS et XNUMX XNUMX IOPS. Globalement, le Kingston DCXNUMXME a constamment atteint des IOPS supérieures à XNUMX XNUMX sur l'ensemble de la matrice de test (sauf dans quelques zones), ce qui en fait l'un des disques les plus équilibrés pour cette charge de travail.
Latence d'écriture aléatoire de 16 K
La latence d'écriture du Kingston DC3000ME 16K a été très solide, terminant le test en tête du classement (le lecteur Pascari étant à un cheveu de la distance). Parmi les points forts du Kingston DC3000ME, on peut citer 14 µs à 1/1, 18 µs à 2/1, 19 µs à 1/4 et 29 µs à 1/2. Avec l'augmentation de la charge, Kingston a maintenu un profil de latence élevé : 126 µs à 8/4, 146 µs à 2/16, 254 µs à 16/4 et 575 µs à 16/8. Même avec la configuration la plus lourde, 32/16, la latence est restée sous contrôle à 1,197 XNUMX µs.
Lecture aléatoire de 16 XNUMX
Dans des conditions de lecture aléatoire de 16 3000, le Kingston DC8ME a affiché des performances constantes et solides jusqu'à atteindre 8/800, après quoi il a commencé à accuser un léger retard. Le pic d'IOPS a atteint un peu moins de 648,686 32 (641 4) au 16e trimestre avec quatre tâches, suivi de 623 16 au 3000e trimestre avec XNUMX tâches et de XNUMX XNUMX au XNUMXe trimestre avec quatre tâches. Malheureusement, le Kingston DCXNUMXME a terminé le test en bas du classement, aux côtés du disque SanDisk.
Latence de lecture aléatoire de 16 K
