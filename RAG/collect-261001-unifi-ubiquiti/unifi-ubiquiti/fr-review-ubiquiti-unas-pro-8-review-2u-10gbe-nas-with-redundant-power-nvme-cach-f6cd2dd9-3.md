---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/fr-review-ubiquiti-unas-pro-8-review-2u-10gbe-nas-with-redundant-power-nvme-cach-f6cd2dd9-3
title: "fr-review-ubiquiti-unas-pro-8-review-2u-10gbe-nas-with-redundant-power-nvme-cach-f6cd2dd9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmarks"]
source: docs/RAG/collect-261001-unifi-ubiquiti/fr-review-ubiquiti-unas-pro-8-review-2u-10gbe-nas-with-redundant-power-nvme-cach-f6cd2dd9.md
source_anchor: ""
source_lines: [57, 84]
sha256: 7ce32665cbf37786530ad45af060b897411006674b2cf3664a1c278d1e866560
---

# fr-review-ubiquiti-unas-pro-8-review-2u-10gbe-nas-with-redundant-power-nvme-cach-f6cd2dd9

Cette configuration nous a permis d'exploiter pleinement l'UNAS 8 sans que la saturation du réseau ne devienne un facteur limitant, tout en conservant des modèles d'accès NFS réalistes, représentatifs des déploiements en production avec accès synchrone et asynchrone. Le client et la plateforme de stockage fonctionnant à 10 GbE, les performances reflètent les capacités du sous-système de stockage plutôt que les contraintes de liaison.
Dans cette section, nous nous concentrons sur les benchmarks FIO suivants :
- Lecture aléatoire
- Écriture aléatoire
- Lecture séquentielle
- Écriture séquentielle
Lecture séquentielle de 1 million – Bande passante
Lors du test de lecture séquentielle de 1 million de données, l'UNAS Pro 8 affiche une montée en puissance rapide avec l'augmentation de la concurrence. À faible profondeur de file d'attente et faible nombre de tâches, le débit initial se situe autour de 500-600 Mo/s, mais grimpe rapidement dans la plage intermédiaire, dépassant 1.5 Go/s pour les configurations 8Q/8T à 16Q/16T. À une concurrence plus élevée, le système atteint un plateau stable légèrement supérieur à 2.2 Go/s, où les performances restent constantes jusqu'à la limite supérieure du test. À ce niveau de débit maximal, la bande passante combinée de 20 GbE fournie par les deux interfaces 10 GbE est presque saturée, la marge restante étant principalement due à la surcharge du protocole et du système de fichiers plutôt qu'à des limitations de stockage.
Lecture séquentielle de 1M – Latence
Lors du test de latence de lecture séquentielle de 1 million de requêtes, la latence reste stable tout au long du test, suivant une courbe régulière et prévisible à mesure que la charge augmente. En début de test, la latence oscille entre 2 et 5 ms et demeure inférieure à 10 ms. À des niveaux de charge intermédiaires, la latence augmente progressivement pour atteindre 25 à 75 ms, reflétant la montée en puissance du système vers un débit plus élevé. Enfin, en fin de test, la latence grimpe brutalement, atteignant plusieurs centaines de millisecondes et culminant à environ 1.8 seconde sous une concurrence maximale de 32/64.
Écriture séquentielle de 1 Mo – Bande passante
Lors du test d'écriture séquentielle de 1M, l'UNAS Pro 8 montre une nette séparation entre les quatre configurations testées, soulignant l'impact du comportement de synchronisation et de la mise en cache sur le débit d'écriture soutenu.
La configuration de cache asynchrone offre les meilleures performances globales, avec une montée en puissance rapide et un débit maximal maintenu entre 480 et 515 Mo/s lorsque la concurrence augmente. Cela représente la limite supérieure de la plateforme pour les écritures séquentielles de gros blocs.
En mode asynchrone sans cache, le débit est légèrement inférieur, mais reste élevé, se stabilisant généralement entre 430 et 490 Mo/s une fois le système stabilisé. Bien qu'il ne bénéficie pas des avantages supplémentaires de la mise en mémoire tampon, les performances demeurent constantes et prévisibles, même avec des files d'attente importantes.
La configuration Sync Cache affiche des performances inférieures aux modes asynchrones, avec un débit maximal de 100 à 130 Mo/s. Bien que nettement plus lente, elle conserve une courbe relativement stable malgré l'augmentation de la charge, ce qui reflète le coût des garanties d'écriture synchrone, même avec l'aide du cache.
Enfin, le mode Sync No-Cache offre le débit le plus faible, se situant principalement entre 40 et 60 Mo/s tout au long du test. Ce comportement est normal pour des écritures sur disque dur entièrement synchrones et protégées par parité, sans mise en cache, et représente une configuration privilégiant la durabilité plutôt que les performances.
Écriture séquentielle 1M – Latence
Les résultats de latence d'écriture séquentielle 1M montrent une séparation claire entre le comportement d'écriture asynchrone et synchrone, la latence augmentant de manière prévisible à mesure que la profondeur de la file d'attente et le nombre de tâches augmentent.
Les modes Async Cache et Async No-Cache maintiennent une latence relativement maîtrisée tout au long du test. Pour les faibles profondeurs de file d'attente, la latence se situe entre 50 et 100 ms, puis augmente progressivement pour atteindre environ 500 à 900 ms. En cas de charge de travail maximale, la latence des deux modes asynchrones culmine autour de 1.8 à 2.0 secondes, ce qui indique une saturation stable sans emballement de la file d'attente.
La configuration du cache synchrone présente une latence nettement plus élevée en charge. Dès le début du test, la latence atteint plusieurs centaines de millisecondes, puis grimpe jusqu'à environ 1.5 à 3.0 secondes en milieu de test. À pleine concurrence, les écritures mises en cache synchrones subissent des latences d'environ 4.5 à 5.0 secondes, ce qui reflète le coût du comportement de validation synchrone, même avec l'aide du cache.
Le comportement le plus extrême s'observe en mode Sync No-Cache, où la latence augmente fortement avec la concurrence. Pour des charges de travail moyennes, la latence dépasse les 2 à 3 secondes, et pour les charges les plus élevées, elle grimpe en flèche, atteignant un pic de 6.5 à 8.6 secondes. Cette courbe abrupte est caractéristique des écritures sur disque dur entièrement synchrones et protégées par parité, sans mise en mémoire tampon, où le système privilégie l'intégrité des données à la réactivité.
Lecture aléatoire 4K – IOPS
Le test de lecture aléatoire 4K met en évidence l'impact du comportement de synchronisation sur les performances des petits blocs et sépare clairement les quatre configurations.
La configuration Sync Cache offre de loin les IOPS les plus élevées. Avec une faible profondeur de file d'attente, les performances débutent autour de 5 000 à 6 000 IOPS et augmentent rapidement en milieu de plage, culminant entre 22 000 et 23 500 IOPS. À mesure que la concurrence augmente, les IOPS commencent à diminuer légèrement en fin de plage, chutant à 15 000-18 000 IOPS, ce qui indique une saturation et une contention de la file d'attente plutôt qu'une instabilité.
Avec le cache asynchrone, les performances sont nettement inférieures, mais plus stables. Les IOPS varient généralement de 2 000 à 4 300, les valeurs les plus élevées étant atteintes pour les faibles profondeurs de file d'attente, avant de diminuer progressivement avec l'augmentation de la concurrence. Cette diminution reflète le passage d'une réactivité optimisée par le cache à un comportement dépendant du disque.
Les configurations Async No-Cache et Sync No-Cache se situent toutes deux en bas du graphique, offrant des performances relativement modestes. Ces configurations fonctionnent principalement entre 0.6 et 1 000 IOPS sur la majeure partie du test, avec une montée en puissance minimale pour des profondeurs de file d'attente plus élevées. Cette courbe plate est typique des charges de travail de lecture aléatoire sur disque dur sans mise en cache, où la latence mécanique est prépondérante.
Lecture aléatoire 4K – Latence
Les résultats de latence de lecture aléatoire 4K montrent une progression claire et attendue à mesure que la concurrence augmente, la latence restant bien contrôlée à de faibles profondeurs de file d'attente avant d'augmenter fortement une fois que le système est saturé.
