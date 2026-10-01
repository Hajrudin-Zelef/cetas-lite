---
id: collect-261001-general-networking/general-networking/fr-review-kioxia-cm9-r-15-36tb-review-bics8-flash-hits-full-speed-at-low-queue-d-f4c610a5-1
title: "fr-review-kioxia-cm9-r-15-36tb-review-bics8-flash-hits-full-speed-at-low-queue-d-f4c610a5"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["datacenter", "intel", "nvidia"]
source: docs/RAG/collect-261001-general-networking/fr-review-kioxia-cm9-r-15-36tb-review-bics8-flash-hits-full-speed-at-low-queue-d-f4c610a5.md
source_anchor: ""
source_lines: [1, 52]
sha256: b047c454e1203fc43de976b69ab6bd32c6718581fef5cbbb2bec1cbc0c4a3e1b
---

# fr-review-kioxia-cm9-r-15-36tb-review-bics8-flash-hits-full-speed-at-low-queue-d-f4c610a5

La famille de SSD KIOXIA CM9-R est la première à intégrer la mémoire flash BiCS de huitième génération. Le modèle E3.S de 15.36 To testé en laboratoire offre les performances les plus intéressantes de la gamme : 3 400 000 IOPS en lecture aléatoire et 14 800 Mo/s en lecture séquentielle, pour une consommation typique de 25 W. Les SSD d'entreprise de cinquième génération atteignent généralement des vitesses supérieures à 20 000 Mo/s pour afficher les chiffres annoncés. KIOXIA met en avant le fait que la conception CBA (CMOS directement lié à la matrice) de la mémoire BiCS8 offre un débit exceptionnel sans surconsommation énergétique.
Le CM9-R est le modèle de la gamme dédié à la lecture intensive (1 DWPD) ; une gamme mixte CM9-V (3 DWPD) l'accompagne. Les capacités s'échelonnent de 1.92 To à 30.72 To en E3.S, avec une version haut de gamme de 61.44 To réservée au format U.2 2.5 pouces. En dessous de 7.68 To, la gamme utilise la mémoire flash BiCS de 5e génération ; les modèles de 7.68 To, 15.36 To et 30.72 To intègrent la 8e génération. Notre modèle de test de 15.36 To (KCM9XRJE15T3) offre précisément les meilleures performances d'écriture de la gamme : 11 000 Mo/s en écriture séquentielle et 540 000 IOPS en écriture aléatoire, des records inégalés.
Ce disque est compatible PCIe 5.0 x4 avec prise en charge double port x2 pour les topologies haute disponibilité, NVMe 2.0 et NVMe-MI 1.2c, et prend en charge la spécification OCP Datacenter NVMe SSD v2.5 (toutes les exigences ne sont pas couvertes). La protection contre les coupures de courant et la protection des données de bout en bout sont incluses de série, avec des variantes de sécurité SIE, SED et FIPS 140-3 SED disponibles dans la gamme.
Spécifications du KIOXIA CM9-R
| Spécifications | KIOXIA CM9-R 15.36 To (E3.S) | 
|---|---|
| En savoir plus sur la plateforme |  | 
| Modèle | KCM9XRJE15T3 (SIE) KCM9DRJE15T3 (SED) KCM9FRJE15T3 (FIPS SED) | 
| Capacités | 15,360 GB | 
| Facteur de forme | E3.S, 7.5 mm | 
| Interface | PCIe 5.0 (simple x4, double x2), NVMe 2.0, NVMe-MI 1.2c | 
| NON | KIOXIA BiCS FLASH génération 8 3D TLC (CBA) | 
| Performances (port unique x4, jusqu'à) |  | 
| Lecture séquentielle (128 KiB) | 14,800 Mo / s | 
| Écriture séquentielle (128 Kio) | 11,000 Mo / s | 
| Lecture aléatoire (4 KiB) | 3,400K IOPS | 
| Écriture aléatoire (4 KiB) | 540K IOPS | 
| Latence de lecture/écriture (4 KiB QD1, typ.) | 65 µs / 10 µs | 
| Puissance et endurance |  | 
| Alimentation (active / prête) | 25 W typ. / 5 W typ. | 
| Endurance | 1 DWPD | 
| MTTF / Garantie | 2,500,000 5 heures / XNUMX ans | 
| Fonctionnalité |  | 
| Protection | Protection contre la perte de puissance Protection des données de bout en bout Double port pour HA | 
| Conformité | SSD NVMe OCP Datacenter v2.5 (partiel) | 
Contexte des tests et comparables
Nous utilisons un serveur Dell PowerEdge R760 exécutant Ubuntu 22.04.2 LTS comme plateforme de test pour toutes les charges de travail présentées dans ce rapport. Équipé d'un boîtier JBOF Serial Cables Gen5, il offre une large compatibilité avec les SSD U.2, E1.S, E3.S et M.2. La configuration de notre système est détaillée ci-dessous :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
- Nvidia L4
Comparaison des lecteurs
- KIOXIA CD9P-R 7.68 To (1 DWPD)
- Disque dur externe Sandisk DC SN861 7.68 To (1 DWPD)
- Solidigm PS1010 7.68 To (1 DWPD)
- Solidigm PS1030 12.8 To (3 DWPD)
- Micron 7600 MAX 6.4 To (3 DWPD)
- Micron 9550 MAX 12.8 To (3 DWPD)
- Micron 9550 Pro 7.68 To (1 DWPD)
L'association avec le CD9P-R de KIOXIA est primordiale : le CD9P-R a démontré les capacités du BiCS8 en matière de positionnement dans les centres de données, et le CM9-R est la version phare pour entreprises de cette même mémoire flash.
Performances FIO
Écriture séquentielle de 128 K (IODepth 16 / NumJobs 1)
Le CM9-R a débuté son test FIO avec un excellent résultat en écriture séquentielle 128K en régime stable : 8 668,1 Mo/s à 230.4 µs, se classant troisième du groupe derrière les Micron 9550 MAX (10 957,9 Mo/s) et 9550 Pro (10 354,6 Mo/s), et largement en tête des autres modèles, dont les performances se situaient entre 6 370 et 7 127 Mo/s. Face à son homologue CD9P-R (6 912,4 Mo/s), le disque professionnel affiche un avantage de 25 %, un écart significatif étant donné que les deux utilisent la même mémoire flash BiCS8.
Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Le résultat en lecture est le seul de notre jeu de données qui s'écarte sensiblement des performances annoncées du disque. Avec une profondeur d'E/S de 64 sur un seul nœud de calcul, le CM9-R a atteint 9 974,6 Mo/s en 801.7 µs, se classant dernier du groupe et bien en deçà des 14 800 Mo/s spécifiés et des 14 235,9 Mo/s obtenus par son homologue CD9P-R dans la même configuration. Il s'agit d'un test sur une seule tâche ; les graphiques ci-dessous montrent le CM9-R fonctionnant à pleine vitesse lorsque la charge est répartie sur plusieurs tâches. Toutefois, les utilisateurs effectuant des lectures de gros blocs sur un seul flux doivent prendre en compte ce comportement et tester leur propre pipeline.
Écriture aléatoire 64K
Lors du test d'écriture aléatoire de 64 Ko, le CM9-R s'est classé deuxième du groupe, avec un pic à 9 635,3 Mo/s (profondeur d'E/S 8 / nombre de tâches 2, à seulement 103.4 µs), juste derrière le Micron 9550 MAX à 10 878,1 Mo/s. Le test à faible profondeur de file d'attente a été le meilleur du lot : 3 334,8 Mo/s à 18.4 µs (profondeur d'E/S 1 / nombre de tâches 1), devançant le Solidigm PS1030 (2 901,7 Mo/s à 21.1 µs) et tous les autres. La latence maximale est restée maîtrisée, plafonnant à 1 942,5 µs sur l'ensemble du test, tandis que plusieurs concurrents ont dépassé les 2 300 µs et le PS1010 a atteint près de 6 000 µs.
Lecture aléatoire 64K
Avec des débits de lecture de 64 000 Ko, le CM9-R a affiché le meilleur débit d'ouverture en flux unique du groupe, atteignant 1 359,0 Mo/s à 45.6 µs (IODepth 1 / NumJobs 1), surpassant même le CD9P-R (1 334,0 Mo/s), les autres disques testés affichant des débits environ deux fois inférieurs. Autre performance remarquable : sa faible consommation d'énergie requise. Le disque a atteint son pic de 13 402,4 Mo/s (IODepth 4 / NumJobs 8) avec une latence de seulement 148.7 µs, alors que la plupart des autres disques nécessitaient une configuration IODepth 32 pour atteindre les mêmes performances. Le Solidigm PS1030 a dominé le classement des débits en file d'attente profonde avec 14 162,9 Mo/s, mais aucun autre disque du groupe n'offre des performances aussi bonnes que les deux disques KIOXIA pour les lectures de 64 000 Ko à faible consommation d'énergie.
Écriture séquentielle 16K
Note de programmation : notre couverture 16K est désormais séquentielle plutôt qu’aléatoire, une charge de travail que nous avons initialement analysée dans notre test du Micron 9550 MAX et que nous appliquons à l’ensemble du groupe, car les flux séquentiels de taille moyenne dominent les déploiements réels, en particulier les pipelines d’IA qui déplacent des données ordonnées par séquences plutôt que de disperser des données aléatoires.
