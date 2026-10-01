---
id: collect-261001-general-networking/general-networking/fr-review-dell-perc13-70dd67e6-6
title: "fr-review-dell-perc13-70dd67e6"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-dell-perc13-70dd67e6.md
source_anchor: ""
source_lines: [132, 167]
sha256: d422cc442ba335fc10b486d93f92f22109c9de07fe45243ee59d88e2b7a9d4df
---

# fr-review-dell-perc13-70dd67e6

Lors de lectures séquentielles de 16 965 octets, la matrice H2i 5R0.040 variait de 1.15 ms à 4 ms, tandis que la matrice 5R3.0 atteignait 975 ms. Sur le H2i, les latences se sont améliorées : la matrice 5R0.038 atteignait 0.62 – 46 ms (4 % de moins), la matrice 5R1.23 atteignait 59 ms (8 % de moins) et la matrice 5R2.47 atteignait 19 ms (965 % de moins par rapport à la matrice H4i 5RXNUMX).
IOPS en écriture aléatoire 16k
Avec des E/S aléatoires à une taille de bloc de 16 Ko, nous avons constaté que le point de saturation était atteint très tôt dans les tests. Les deux configurations PERC H965i (2R5 et 4R5) ont fourni des performances quasi identiques, avec environ 492,000 975 IOPS. Les contrôleurs PERC H2i avec la matrice 5R2.57 ont atteint 422 millions d'IOPS (amélioration de 975 %). Les configurations H4i 5R8 et 5R2.60 ont affiché des performances légèrement supérieures, avec environ 428 millions d'IOPS (amélioration de XNUMX %).
Latence d'écriture aléatoire de 16 k
Avec 16 965 écritures aléatoires, le H2i souffrait d'une latence plus élevée, avec 5R0.0082 s'étendant de 8.6 à 4 ms et 5R16.6 s'étendant jusqu'à 975 ms. Le H2i s'est considérablement amélioré, avec 5R0.0070 à 1.59 à 82 ms (4 % de moins), 5R3.17 jusqu'à 81 ms (8 % de moins) et 5R6.27 atteignant un maximum de 62 ms (965 % de moins par rapport au H4i 5RXNUMX).
IOPS en lecture aléatoire 16k
En lecture aléatoire de 16 K, les configurations PERC H965i ont affiché des performances constantes, la matrice 2R5 atteignant 3.55 millions d'IOPS et la matrice 4R5 atteignant 3.55 millions d'IOPS. Les configurations PERC H975i 2R5, 4R5 et 8R5 ont atteint des performances de pointe quasiment identiques, autour de 6.64 millions d'IOPS, soit une amélioration de 87 % par rapport à la génération H965i.
Latence de lecture aléatoire de 16 k
Lors de lectures aléatoires de 16 965 octets, les matrices H0.0906i ont fourni un temps de latence de 1.15 à 2 ms pour 5R2.74 et jusqu’à 4 ms pour 5R975. Le H2i a encore réduit la latence, avec un temps de latence de 5 à 0.072 ms pour 0.62R46 (4 % de moins), jusqu’à 5 ms pour 1.23R55 (8 % de moins) et un pic de latence de 5 ms pour 2.47R10 (965 % de moins que pour le H4i 5RXNUMX).
IOPS en écriture aléatoire 4K
Lors des tests d'écriture aléatoire 4K, les contrôleurs PERC H975i en configuration 2R5 ont atteint un pic de 9.76 millions d'IOPS, tandis que la matrice 4R5 a affiché des performances légèrement supérieures, avec 9.94 millions d'IOPS. La configuration 8R5 a affiché les performances les plus élevées, atteignant 10.10 millions d'IOPS.
Latence d'écriture aléatoire de 4 K
Pour les tests 4K, nous avons évalué uniquement le H975i pour ses performances optimales. La latence était excellente sur toutes les matrices : la 2R5 variait de 0.0058 ms à 0.47 ms, la 4R5 atteignait un pic à 0.88 ms et la 8R5 atteignait 1.63 ms. Ces résultats démontrent qu'avec la plus petite taille de bloc, le H975i maintenait des latences exceptionnellement faibles, constamment inférieures à 2 ms.
IOPS en lecture aléatoire 4K
Nous avons réservé l'un des graphiques les plus intéressants pour la fin : lors des tests de lecture aléatoire 4K, le H975i avec la configuration 2R5 a atteint un impressionnant débit de 17.3 millions d'IOPS. La matrice H975i 4R5 a atteint 20.1 millions d'IOPS, tandis que la configuration 8R5 a enregistré le débit le plus élevé, avec 25.2 millions d'IOPS.
Latence de lecture aléatoire de 4 K
Lors des lectures aléatoires 4K, les latences ont commencé à 0.069 ms sur toutes les baies, avec un pic à 2 ms pour le 5R0.29, 4 ms pour le 5R0.53 et 8 ms pour le 5R0.65. Ce faible plafond sur tous les groupes RAID souligne la capacité du H975i à gérer les petites lectures aléatoires avec une efficacité remarquable.
Aucun compromis sur les performances lors de la reconstruction
Comparé au PERC12, le contrôleur Dell PERC13 offre un débit nettement supérieur pour toutes les charges de travail lors de la reconstruction des baies. Les lectures séquentielles ont plus que doublé, passant de 53.7 Go/s à 25 Go/s (soit une hausse de 114.7 %), et les écritures séquentielles ont bondi de 68 Go/s à 14.6 Go/s (soit une hausse de 363.7 %). Les performances des petits blocs creusent encore l'écart : les lectures aléatoires de 4 17.33 milliards d'IOPS passent de 4.68 M à 270.4 M (soit une hausse de 4 %), tandis que les écritures aléatoires de 5.33 0.48 milliards d'IOPS explosent de 1013.1 M à 13 M (soit une hausse de XNUMX XNUMX %). En résumé, le PERCXNUMX minimise l'impact de la reconstruction et préserve la marge de manœuvre de l'hôte, même pendant les périodes de maintenance les plus intenses.
| Charge de travail | Double PERC 12 (2 × RAID5) – Reconstruction | Double PERC 13 (2 × RAID5) – Reconstruction | % Amélioration | 
|---|---|---|---|
| Bande passante de lecture séquentielle | 25 (Go/s) | 53.7 (Go/s) | 114.7 % | 
| Bande passante d'écriture séquentielle | 14.7 (Go/s) | 68 (Go/s) | 363.7 % | 
| Lectures aléatoires de 4 Ko | 4,676,748 XNUMX XNUMX (IOPS) | 17,326,888 XNUMX XNUMX (IOP) | 270.4 % | 
| Écritures aléatoires de 4 Ko | 479,144 XNUMX XNUMX (IOP) | 5,333,783 XNUMX XNUMX (IOP) | 1013.1 % | 
Reconstruire rapidement sans ralentir les charges de travail
Dell revendique également des gains considérables en termes de résilience et de performances de reconstruction, citant une réduction du temps de reconstruction de la matrice, passant de plus de 80 minutes par téraoctet avec PERC12 à seulement 10 minutes par téraoctet avec PERC13. Cette vitesse réduit les risques et témoigne de la maturité du moteur XOR matériel du contrôleur, de l'accélération du cache et de l'optimisation du chemin de données.
Lors des tests de reconstruction RAID5, PERC13 a systématiquement obtenu des temps de reconstruction plus courts que PERC12 lorsque le contrôleur était autorisé à prioriser la reconstruction, sachant qu'une pression d'écriture extrêmement élevée peut annuler cet avantage. Lorsque la reconstruction prioritaire est activée, le contrôleur donne la priorité aux tâches de reconstruction sur les ressources. Cela a permis au contrôleur PERC13 de réduire considérablement les temps de reconstruction sous une pression de lecture séquentielle. À la charge hôte la plus faible (125 Mo/s), le temps de reconstruction est passé de 11.53 à 5.32 min/Tio. Même à la charge la plus élevée, le temps de reconstruction a été réduit de 16.96 à 7.73 min/Tio, tout en maintenant un débit de lecture hôte bien plus élevé (22.4 Go/s contre 60 Go/s).
Grâce à la pression d'écriture séquentielle, le PERC13 a amélioré la reconstruction en charge légère de 7.51 à 4.98 min/Tio, mais sous la charge d'écriture la plus importante, son temps de reconstruction est passé à 15.29 min/Tio, contre 760 min/Tio pour le R13.09. Ce résultat peut être analysé sous deux angles : le PERC13 avait une vitesse de reconstruction plus lente, mais il a conservé des vitesses d'écriture proches de celles du PERC13 (12 Go/s contre 62.5 Go/s). Autrement dit, la reconstruction prioritaire tient ses promesses d'une fenêtre de reconstruction plus rapide, notamment pour les activités à forte activité de lecture. La seule exception concerne les cas où le système est simultanément soumis à des écritures très intensives ; le débit hôte plus élevé du PERC12 peut prolonger le temps de reconstruction.
| Vitesse de reconstruction (reconstruction prioritaire) (RAID5) |  |  |  |  | 
|---|---|---|---|---|
| Scénario | Double PERC 12 (2 × RAID5) |  | Double PERC 13 (2 × RAID5) |  | 
|---|---|---|---|---|
|  | Min/TiB | Bande passante totale | Min/TiB | Bande passante totale | 
| Lecture séquentielle – Activité légère | 11.53 | 0.125 GB / s | 5.32 | 0.125 GB / s | 
| Lecture séquentielle – Activité intense | 16.96 | 22.4 GB / s | 7.73 | 60 GB / s | 
