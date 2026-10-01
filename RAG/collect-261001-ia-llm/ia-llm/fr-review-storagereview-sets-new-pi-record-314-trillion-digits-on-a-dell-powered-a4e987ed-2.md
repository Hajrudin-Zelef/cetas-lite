---
id: collect-261001-ia-llm/ia-llm/fr-review-storagereview-sets-new-pi-record-314-trillion-digits-on-a-dell-powered-a4e987ed-2
title: "fr-review-storagereview-sets-new-pi-record-314-trillion-digits-on-a-dell-powered-a4e987ed"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Google"]
dates: []
keywords: ["amd"]
source: docs/RAG/collect-261001-ia-llm/fr-review-storagereview-sets-new-pi-record-314-trillion-digits-on-a-dell-powered-a4e987ed.md
source_anchor: ""
source_lines: [23, 45]
sha256: 7ba9994013f99dfc1bc8aff3c780698c1fada050d47b37009c61d5653096b76f
---

# fr-review-storagereview-sets-new-pi-record-314-trillion-digits-on-a-dell-powered-a4e987ed

Il est important de noter que lors du calcul, notre test de 314 To a utilisé des SSD en configuration JBOD sans résilience des données. La consommation d'énergie et les performances globales du système ont motivé ce choix, mais ont également suscité une réflexion sur la conception d'une solution de stockage adaptée à la charge de travail. Chaque charge de travail est différente, et certaines, pouvant être redémarrées avec un impact minimal sur la production, ne nécessitent pas forcément le même niveau de tolérance aux pannes. Dans notre cas, nous avons privilégié la protection des données de sortie grâce à un RAID logiciel traditionnel.
Durée totale d'exécution : 110 jours
Malgré un nombre de chiffres calculés supérieur à celui de toutes les exécutions précédentes, le temps d'exécution réel a été nettement inférieur au record précédent, qui nécessitait environ 225 jours (175 jours de calcul hors interruptions). Cette durée ininterrompue de 110 jours s'explique par un système d'exploitation stable, une charge de calcul en arrière-plan minimale, une topologie NUMA équilibrée et une matrice de calcul temporaire conçue pour le modèle généré par le processeur de calcul à cette échelle.
Points forts techniques
- Total des chiffres calculés: 314,000,000,000,000
- Matériel utiliséServeur Dell PowerEdge R7725 avec 2 processeurs AMD EPYC 9965, 1.5 To de mémoire vive DDR5 et 40 disques durs Micron 6550 Ion de 61.44 To chacun.
- Logiciels et algorithmes: y-cruncher v0.8.6.9545, Chudnovsky
- Usure du SSD par SMART7.3 Po écrits par disque, soit 249.11 Po répartis sur les 34 SSD utilisés pour l'échange.
- Le plus grand point de contrôle logique: 850 538 385 064 992 (774 Tio)
- Utilisation logique du disque de pointe: 126 658 805 195 776 600 (1.43 PiB)
- Octets de disque logique lus: 126 658 805 195 776 600 (132 PiB)
- Octets de disque logique écrits: 126 658 805 195 776 600 (112 PiB)
- Date de début : : Jeu 31 juil 17:16:41 2025
- Date de fin: mar. 18 nov. 05:57:08 2025
- pi: 8 793 223,144 secondes, 101.773 jours
- Temps de calcul total: 9274878.580 seconde
- Temps de mur du début à la fin: 9463226.454 seconde
Réflexions de clôture
Pendant des décennies, les tentatives de calcul extrême de pi ont permis de mettre en valeur les machines considérées comme « puissantes » à l'époque. Les premiers records reposaient sur des ordinateurs de bureau hautes performances et du stockage externe, puis sur des équipements d'entreprise locaux. Plus récemment, la course s'est déplacée vers le cloud, où des exploits comme la tentative de Google d'atteindre 100 000 milliards de décimales ont prouvé qu'il était possible d'établir un record par la force brute, avec suffisamment d'instances et d'entrées/sorties. Puis sont apparus les grands clusters de stockage partagé, sacrifiant la simplicité au profit d'un parallélisme brut et d'une consommation énergétique et de refroidissement colossale.
Notre approche a été radicalement différente. Lors de plusieurs tentatives de test record, nous avons considéré y-cruncher comme une charge de travail HPC sérieuse, et non comme un simple exercice ponctuel. Les solutions à 105T et 202T nous ont permis d'identifier les véritables goulots d'étranglement, de dimensionner et d'optimiser le stockage temporaire, d'alimenter les processeurs sans surcharger la couche d'E/S et de renforcer la robustesse du système afin qu'une tâche de plusieurs mois puisse s'achever. La tentative à 314T est le fruit de cette expérience. Il ne s'agit pas simplement d'un chiffre plus élevé, mais d'une conception plus aboutie.
Les indicateurs confirment ces résultats. Nous avons dépassé les 300 billions de chiffres sur un seul serveur Dell PowerEdge R7725 2U équipé de 40 SSD Micron 6550 Ion et de deux processeurs AMD EPYC à 192 cœurs. Le système est resté opérationnel pendant 110 jours consécutifs, sans aucune interruption. Le débit de stockage a plus que doublé par rapport à notre plateforme 202T, tandis que le serveur a consommé en moyenne 1 600 W et 4 305 kWh au total. Cela représente 13.70 kWh par billion de chiffres, soit une fraction de l'énergie consommée par notre précédent cluster 300T : moins de nœuds, moins de complexité, moins d'énergie, et une productivité accrue.
C’est pourquoi ce record dépasse le simple cadre de la simple performance : si un serveur 2U commercial peut supporter une exécution de calcul intensif de cette ampleur, avec un tel niveau de fiabilité et d’efficacité, les mêmes principes de conception sont directement applicables à la recherche scientifique de production. Les modèles climatiques de longue durée, les simulations physiques, les pipelines de génomique et les tâches d’entraînement d’IA reposent tous sur les mêmes fondamentaux : des E/S équilibrées, une gestion thermique prévisible, un firmware stable et une architecture capable de fonctionner en continu pendant des mois. Cette plateforme a désormais prouvé qu’elle pouvait le faire précisément dans des conditions ne tolérant aucune erreur.
Oui, StorageReview a reconquis le titre de champion du calcul en nombres entiers (π) avec 314 billions de décimales. Plus important encore, nous avons établi la nouvelle norme en matière de calcul numérique à grande échelle sur du matériel réel. Si quelqu'un souhaite battre ce record, nous aimerions qu'il atteigne le niveau maximal : plus de décimales, moins de consommation d'énergie, un temps d'exécution plus court et la même fiabilité sans interruption de service. En attendant, il s'agit de la référence en matière d'efficacité.
