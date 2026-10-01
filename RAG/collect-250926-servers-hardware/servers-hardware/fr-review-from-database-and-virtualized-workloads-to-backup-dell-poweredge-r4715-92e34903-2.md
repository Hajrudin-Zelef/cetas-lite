---
id: collect-250926-servers-hardware/servers-hardware/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903-2
title: "fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["acquisition", "amd", "benchmark", "exploit", "valuation"]
source: docs/RAG/clean4/fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903.md
source_anchor: ""
source_lines: [31, 67]
sha256: 67adcc90c2a58f0b5b44612ca8b32063dabd5b8aa526754a761118658a5a91e5
---

# fr-review-from-database-and-virtualized-workloads-to-backup-dell-poweredge-r4715-92e34903

| Alimentations | Platine 800W, 1100W Titane 800W, 1100W FTR pris en charge |  | 
| Refroidissement et ventilateurs |  |  | 
| Options de refroidissement | refroidissement par air |  | 
| Ventilateurs | Jusqu'à quatre ensembles (module à double ventilateur) de ventilateurs remplaçables à chaud | Jusqu'à six ventilateurs remplaçables à chaud | 
| Dimensions |  |  | 
| Hauteur | 42.8 mm (1.68 pouces) | 86.8 mm (3.41 pouces) | 
| Largeur | 482.0 mm (18.97 pouces) |  | 
| Profondeur (avec lunette) | 816.921 mm (32.16 pouces) | 802.4 mm (31.59 pouces) | 
| Profondeur (sans lunette) | 815.141 mm (32.09 pouces) | 801.51 mm (31.55 pouces) | 
| Biseau | Lunette métallique en option |  | 
Quatre options de processeur AMD EPYC série 9005
Dell propose quatre références de processeurs spécifiques pour ses deux serveurs. Ce choix est délibéré et couvre l'ensemble des besoins des PME sans chevauchement ni complexité inutile. Chaque processeur repose sur la microarchitecture Zen 5 d'AMD et partage les mêmes caractéristiques de mémoire et PCIe au niveau de la plateforme.
| Processeur | Coloré | TDP par défaut | Plage de cTDP | Horloge de base | Max Boost | L3 Cache | 
|---|---|---|---|---|---|---|
| EPYC 9335 | 32 | 210W | 200-240W | 3.0 GHz | 4.4 GHz | 128 MB | 
| EPYC 9255 | 24 | 200W | 200-240W | 3.2 GHz | 4.3 GHz | 128 MB | 
| EPYC 9135 | 16 | 200W | 200-240W | 3.65 GHz | 4.3 GHz | 64 MB | 
| EPYC 9015 | 8 | 125W | 120-155W | 3.6 GHz | 4.1 GHz | 64 MB | 
Le modèle 9335 à 32 cœurs est le plus performant de notre gamme et offre la plus grande flexibilité pour les charges de travail exigeantes en calcul. Le modèle 9255 à 24 cœurs est celui qui se rapproche le plus du meilleur rapport qualité-prix observé lors de nos tests, notamment pour les bases de données, où les gains marginaux liés au nombre de cœurs deviennent négligeables au-delà de 24 cœurs. Les modèles 9135 à 16 cœurs et 9015 à 8 cœurs présentent le meilleur rapport qualité-prix. Étant donné que la plupart des logiciels utilisés par les PME sont facturés par cœur, notamment Windows Server, de nombreuses bases de données relationnelles et certaines plateformes d'hyperviseur et de sauvegarde, le nombre de cœurs choisi à l'achat représente un coût récurrent pour toute la durée de vie du déploiement.
Choisir un processeur 8 ou 16 cœurs adapté à la charge de travail, plutôt que de surdimensionner les cœurs inactifs, permet de réduire les coûts d'acquisition et les dépenses liées aux licences. Le modèle 9135 à 16 cœurs est particulièrement pertinent, car il correspond au nombre minimal de cœurs requis par les licences Windows Server. Il constitue ainsi un point de départ idéal pour les environnements Windows qui souhaitent dimensionner correctement leur processeur en fonction des contraintes de la licence, sans compromettre les performances. Le modèle 9015 à 8 cœurs est l'option la plus économique et la plus économe en énergie, idéale pour les charges de travail axées sur le stockage ou les rôles où le processeur n'est pas un facteur limitant. Chaque processeur fonctionne à la même vitesse mémoire et prend en charge le même nombre de lignes PCIe Gen5 ; les choix de configuration des composants supérieurs ne sont donc pas limités par le choix du modèle.
Test de performance
| Configurations de test | Dell PowerEdge R4715 | Dell PowerEdge R5715 | 
|---|---|---|
| Processeurs testés | AMD EPYC 9335, 9255, 9135, 9015 | AMD EPYC 9015 | 
| Mémoire | 384GB DDR5 | 384GB DDR5 | 
| Stockage de démarrage | BOSS RAID1 | BOSS RAID1 | 
| Configuration du rangement avant | 8 SSD Samsung PM9D3a RI U.2 Gen5 NVMe (1.92 To) RAID 10 x 6 | 12 disques durs de 20 To en RAID 6 | 
Performances des bases de données : HammerDB MariaDB TPC-C
La charge de travail principale de cette évaluation est HammerDB exécutant TPC-C face à MariaDB 12.3.1. TPC-C est un benchmark OLTP reconnu qui produit des résultats mesurables et comparables quelle que soit la configuration du processeur et du stockage, et représente le type de charge de travail de base de données transactionnelle au cœur de la plupart des applications pour PME. Nous avons testé deux profils distincts : un profil intensif en ressources processeur, qui sollicite fortement le traitement transactionnel, et un profil intensif en E/S, qui impose une charge plus importante au sous-système de stockage. Les deux profils ont été exécutés sur les quatre options de processeur de la configuration flash R4715 afin d'obtenir une courbe de montée en charge du processeur claire, puis sur la configuration disque dur R5715 afin d'observer les variations en fonction du type de stockage.
Dell PowerEdge R4715
Les résultats de HammerDB montrent une nette amélioration des performances avec l'augmentation du nombre de cœurs sur la plateforme R4715. Avec le processeur EPYC 9015 à 8 cœurs, le système a atteint 480 818 NOPM en mode CPU intensif et 296 105 NOPM en mode E/S intensif avant de se stabiliser malgré l'augmentation du nombre d'utilisateurs virtuels. Le passage au processeur EPYC 9135 à 16 cœurs a permis un gain de débit considérable, portant les performances à 737 445 NOPM en mode CPU intensif et à 493 093 NOPM en mode E/S intensif, tout en permettant au système de supporter un nombre d'utilisateurs virtuels plus élevé avant saturation.
Le passage au processeur EPYC 9255 à 24 cœurs a permis à la plateforme de dépasser le million de NOPM en mode CPU intensif, atteignant un pic de 1 017 429 NOPM, tandis qu'en mode E/S intensif, elle a culminé à 740 574 NOPM. À ce stade, les cœurs supplémentaires continuaient de se traduire directement par un débit transactionnel exploitable, tandis que le sous-système de stockage NVMe suivait la charge croissante de la base de données.
En haut de gamme, le processeur EPYC 9335 à 32 cœurs a affiché les meilleurs résultats sur les deux profils, atteignant 1 133 714 NOPM pour la charge de travail intensive en CPU et 910 321 NOPM pour celle intensive en E/S. La courbe de performance est restée relativement stable même avec un nombre élevé d'utilisateurs virtuels, ce qui indique que la configuration flash R4715 a exploité efficacement les configurations CPU plus puissantes sans que les goulots d'étranglement du stockage ne limitent prématurément les performances.
Dell PowerEdge R5715
Nous avons ensuite testé le Dell PowerEdge R5715, configuré avec 12 disques durs de 20 To en RAID 6, associé au processeur AMD EPYC 9015 à 8 cœurs. Dans le profil sollicitant fortement le processeur, la plateforme a atteint un pic de 484 715 NOPM avec 16 utilisateurs virtuels, le débit augmentant de manière constante avec l'ajout d'utilisateurs supplémentaires avant de se stabiliser près de la saturation.
Le profil d'activité E/S intensive a atteint un pic de 308 012 NOPM avec 24 utilisateurs virtuels, témoignant des excellentes performances transactionnelles de la baie de disques durs haute capacité en conditions de concurrence modérée. À mesure que la charge de travail augmentait, la courbe de montée en charge s'est aplatie, le sous-système de disques durs approchant ses limites de performance pratiques sous une activité de base de données concurrente soutenue.
Stockage partagé Windows Server
