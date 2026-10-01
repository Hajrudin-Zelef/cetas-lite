---
id: collect-261001-general-networking/general-networking/fr-review-dell-perc13-70dd67e6-7
title: "fr-review-dell-perc13-70dd67e6"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-perc13-70dd67e6.md
source_anchor: ""
source_lines: [168, 185]
sha256: 2453c9eb5cfccdba217de3767c4748681aa88ab4ff8a013ea9a723a2e73bf8d7
---

# fr-review-dell-perc13-70dd67e6

| Écriture séquentielle – Activité légère | 7.51 | 0.125 GB / s | 4.98 | 0.125 GB / s | 
| Écriture séquentielle – Activité intense | 13.09 | 12 GB / s | 15.29 | 62.5 GB / s | 
En passant à Priority Host, qui protège intentionnellement les E/S des applications au détriment de la vitesse de reconstruction, les performances sont similaires en lecture et plus nuancées en écriture. En lecture, le contrôleur PERC13 termine les reconstructions beaucoup plus rapidement que l'ancien PERC12, réduisant le temps de charge faible de 11.23 à 6.70 min/Tio et le temps de charge élevée de 38.44 à 19.75 min/Tio, tout en gérant davantage de trafic hôte (46.2 Go/s contre 24.1 Go/s en charge élevée). En écriture, Priority Host privilégie les performances de production : le PERC13 est plus rapide à la charge la plus faible (7.80 contre 5.67 min/Tio), mais en écriture la plus élevée, sa reconstruction atteint 32.81 min/Tio, contre 12 min/Tio pour le PERC25.40. Le temps de reconstruction s'allonge légèrement, mais le PERC13 offre à l'hôte une bande passante d'écriture bien plus élevée (62.4 Go/s contre 12.5 Go/s).
| Vitesse de reconstruction – (Hôte prioritaire) (RAID5) |  |  |  |  | 
|---|---|---|---|---|
| Scénario | Double PERC 12 (2 × RAID5) |  | Double PERC 13 (2 × RAID5) |  | 
|---|---|---|---|---|
|  | Min/TiB | Bande passante totale | Min/TiB | Bande passante totale | 
| Lecture séquentielle – Activité légère | 11.23 | 0.125 GB / s | 6.70 | 0.125 GB / s | 
| Lecture séquentielle – Activité intense | 38.44 | 24.1 GB / s | 19.75 | 46 GB / s | 
| Écriture séquentielle – Activité légère | 7.80 | 0.125 GB / s | 5.67 | 0.125 GB / s | 
| Écriture séquentielle – Activité intense | 25.40 | 12.5 GB / s | 32.81 | 62.4 GB / s | 
Du point de vue du déploiement, le choix est simple. Lorsqu'il est nécessaire de minimiser la fenêtre de vulnérabilité et que vous pouvez tolérer une certaine dépriorisation des E/S, la priorité de reconstruction du PERC13 réduit les temps de reconstruction, notamment pour les scénarios à lecture intensive. Lorsque le maintien de la réactivité des applications est indispensable, Priority Host s'en charge parfaitement ; le PERC13 excelle toujours pour les reconstructions en lecture, tandis que les périodes d'écriture intensives peuvent justifier une planification ou une légère limitation si le temps de reconstruction absolu est un problème.
Conclusion
Le Dell PERC H975i positionne le RAID matériel comme une solution incontournable pour les datacenters d'entreprise centrés sur NVMe. Si les implémentations JBOD et RAID logiciel ont gagné en popularité dans les environnements scale-out, ces approches engendrent une complexité opérationnelle, une surcharge CPU et des temps de récupération prolongés en cas de panne de disque. Le H975i offre une accélération matérielle sur mesure avec des moteurs de parité dédiés, des opérations de reconstruction accélérées et des capacités de gestion intégrées à la pile d'infrastructure Dell.
Pour les charges de travail d'IA et d'apprentissage automatique qui exigent des caractéristiques de débit cohérentes, une variation de latence minimale et une fiabilité de disponibilité maximale, les architectures RAID gérées par le matériel offrent à la fois des performances de calcul et une résilience opérationnelle sans consommer de ressources de traitement hôte critiques.
Les tests de performance valident les améliorations architecturales : le H975i offre une bande passante de lecture séquentielle supérieure de 88 % et une bande passante d'écriture séquentielle supérieure de 318 % à celle de la génération PERC12. Des pics de débit de 103 Go/s et 25.2 millions d'IOPS démontrent les capacités du contrôleur pour les charges de travail gourmandes en données. De plus, les temps de reconstruction sont passés de plus de 80 minutes par téraoctet à seulement 10 minutes par téraoctet, tout en maintenant des performances proches de la production lors des opérations de récupération.
Les interfaces PCIe Gen975 x5 et la conception intégrée en façade du H16i prennent en charge les déploiements GPU denses sans contention de stockage, permettant une évolutivité prévisible des performances dans les configurations multi-accélérateurs. Avec de nombreux serveurs PowerEdge proposant les contrôleurs RAID H965i et H975i, il est évident que les entreprises exploitant des charges de travail émergentes devraient opter pour la nouvelle offre. Si vous déployez une infrastructure d'IA à grande échelle, le H975i offre la base de stockage à large bande passante et à faible latence nécessaire pour optimiser l'utilisation des ressources de calcul.
