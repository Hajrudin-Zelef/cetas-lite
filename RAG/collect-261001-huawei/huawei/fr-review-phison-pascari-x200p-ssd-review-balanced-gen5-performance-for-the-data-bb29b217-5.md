---
id: collect-261001-huawei/huawei/fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217-5
title: "fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["nand"]
source: docs/RAG/collect-261001-huawei/fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217.md
source_anchor: ""
source_lines: [133, 142]
sha256: a8ce89475d8c058cb8ad1c97d7e47724054b5b8a1ee2d4372a3c9cd316acd3e1
---

# fr-review-phison-pascari-x200p-ssd-review-balanced-gen5-performance-for-the-data-bb29b217

Les résultats de latence de lecture GDSIO du Pascari X200P mettent en évidence une relation claire entre la taille des blocs, le nombre de threads et la latence. Avec une taille de bloc plus petite de 16 Ko et un seul thread, le lecteur atteint une latence moyenne de seulement 0.026 ms. Cependant, avec une taille de bloc de 128 Ko et la même taille de bloc, la latence augmente considérablement pour atteindre 1.076 ms. Avec 128 Ko de blocs, la latence est de 0.050 ms avec un seul thread et atteint 3.056 ms avec 128 threads. Pour des blocs plus importants de 1 Mo, la latence commence à 0.268 ms avec un seul thread et atteint 20.324 ms avec un parallélisme maximal.
Avec une taille de bloc de 16 Ko, le débit X200P démarre à 0.58 Gio/s (latence de 25.17 µs) au premier jour et atteint 1 Gio/s (latence de 1.22 ms) au deuxième jour. Cela représente un gain modeste en bande passante, mais une forte augmentation de la latence, suggérant une saturation précoce à cette petite taille d'E/S.
À 128 Ko, les performances évoluent mieux, commençant à 2.63 Gio/s (45.55 µs) et atteignant 4.94 Gio/s (3.16 ms) à QD128. Cela représente une augmentation significative du débit, mais là encore, la latence évolue fortement, indiquant une surcharge croissante dans des conditions de profondeur de file d'attente élevée.
Avec une taille de bloc de 1 Mo, le disque démarre fort à 4.52 Gio/s (215 µs) et culmine à 5.02 Gio/s (24.9 ms) à QD128. Le gain de débit est minime par rapport à 128 Ko, et la latence à QD128 devient la plus élevée de tous les tests, ce qui indique des gains d'efficacité limités pour les transferts plus importants au-delà de 128 Ko dans les files d'attente profondes.
Les résultats de latence d'écriture GDSIO du Pascari X200P montrent une évolution constante, la latence augmentant avec la taille des blocs et le nombre de threads. Avec une taille de bloc de 16 Ko et un seul thread, le disque affiche une latence moyenne de 0.025 ms, atteignant 1.595 ms à 128 threads. Avec 128 Ko de blocs, la latence passe de 0.046 ms à 3.159 ms dans les mêmes conditions. Avec la taille de bloc maximale de 1 Mo, la latence démarre à 0.215 ms et atteint 24.917 ms à la profondeur de thread maximale. Malgré l'augmentation attendue de la latence, le Pascari X200P domine le groupe avec des tailles de blocs et des nombres de threads plus élevés, maintenant la latence la plus faible sous des charges de travail d'écriture parallèle importantes.
Conclusion
Le SSD Phison Pascari X200P 7.68 To est un disque dur professionnel doté de la technologie NAND TLC et optimisé pour les performances PCIe Gen5, adapté aux charges de travail générales et gourmandes en contenu. Il est conçu pour les environnements où un débit élevé, une forte évolutivité et une flexibilité de déploiement priment sur les réglages spécifiques à l'hyperscale. Avec la prise en charge des formats U.2, U.3 et E3.S, ainsi que des fonctionnalités telles que la protection contre les coupures de courant, le chiffrement AES-XTS 256 bits et la gestion NVMe-MI, le X200P offre une base solide pour votre infrastructure de stockage.
En termes de performances, le X200P excelle dans les scénarios séquentiels et de lecture intensive, se classant régulièrement parmi les meilleurs aux tests 128 K et 64 K et évoluant efficacement sous les charges de travail CDN. Les tests FIO confirment sa performance en lecture séquentielle et affichent des performances compétitives sur les charges de travail de lecture aléatoire. Bien qu'il soit à la traîne des disques durs haut de gamme comme Micron et SanDisk dans les conditions d'écriture intensive et de forte concurrence, son comportement d'écriture prévisible et efficace le rend parfaitement adapté à un large éventail de déploiements d'entreprises de taille moyenne.
Les tests GDSIO mettent en évidence les atouts du disque dans les applications axées sur le débit. Le X200P maintient une excellente latence avec des blocs de petite taille et domine les accès parallèles intensifs avec des transferts de blocs volumineux. Bien que la latence augmente avec des files d'attente plus profondes, le réglage de Phison garantit la stabilité et la réactivité du disque sous une pression soutenue.
Globalement, le Pascari X200P est un SSD d'entreprise complet, offrant d'excellentes performances et des fonctionnalités adaptées aux charges de travail réelles. Il sera intéressant de voir si Phison parviendra à passer d'une entreprise axée sur les contrôleurs à une entreprise proposant une gamme complète de solutions de disques intégrés. Pour l'instant, le X200P semble être un début prometteur dans cette direction.
