---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r5715-review-bc780bb6-2
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "ethernet", "gpu"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [56, 79]
sha256: 7e03c0464cfb5d774afef42d000adfb3db84914a33f83afd9737390fe8d9eb99
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

L'unité testée est équipée d'une baie avant SAS/SATA 12 x 3.5 pouces, dont quatre baies occupées par des disques durs SATA 6 Gb/s 7 200 tr/min de 20 To et huit baies libres pour de futures extensions. Une configuration alternative avec fond de panier SAS/SATA 16 x 2.5 pouces est également disponible, selon les besoins. La gestion RAID est assurée par le contrôleur interne PERC H365i ou le PERC H965i, plus performant. Le démarrage est géré séparément par un module BOSS-N1 DC-MHS dédié à l'arrière, isolant ainsi le système d'exploitation des données. Cette conception astucieuse évite l'erreur fréquente consistant à exécuter le système d'exploitation et le stockage des données sur la même baie.
Processeur et refroidissement
La R5715 est une plateforme mono-socket basée sur le processeur AMD EPYC série 9005, prenant en charge jusqu'à 32 cœurs. Son imposant dissipateur thermique est un modèle tour à ailettes intégrant des caloducs en cuivre, fixé au socket SP5 par six vis imperdables. Le refroidissement est entièrement par air ; jusqu'à six ventilateurs remplaçables à chaud assurent la circulation de l'air de l'avant vers l'arrière du châssis. Le refroidissement liquide n'est pas compatible avec cette plateforme.
Mémoire
La carte mère R5715 dispose de 24 emplacements DIMM DDR5 répartis en deux groupes de part et d'autre du socket du processeur. Cette plateforme est exclusivement compatible avec les modules RDIMM ; les modules UDIMM et LRDIMM ne sont pas pris en charge. La capacité maximale atteint 1.5 To avec 64 Go de RDIMM par emplacement, fonctionnant à une fréquence maximale de 5 200 MT/s. Le modèle testé est livré avec plusieurs emplacements occupés, tirant parti de l'architecture mémoire multicanal d'EPYC pour offrir une bande passante agrégée élevée sur l'ensemble du sous-système mémoire.
Extension et mise en réseau PCIe
Le R5715 offre jusqu'à quatre emplacements PCIe Gen5 x16 pleine hauteur (emplacements 2, 3, 7 et 9), répartis sur cinq emplacements de riser (Risers 1 à 5) visibles à l'intérieur du châssis. Deux emplacements supplémentaires pour cartes réseau OCP NIC 3.0 (emplacements 4 et 10, Gen5 x16) prennent en charge les adaptateurs réseau OCP 1 GbE, 10 GbE ou 25 GbE. Pour une connectivité haut débit, les cartes réseau PCIe AIC prennent en charge jusqu'à 100 GbE et 400 GbE, avec NDR VPI (400 GbE). Un port Ethernet BMC 1 Gb dédié est intégré au panneau arrière pour la gestion iDRAC hors bande. Le R5715 ne propose pas d'options GPU ; il s'agit d'une plateforme de stockage et de calcul, et non d'un châssis d'accélération.
Gestion iDRAC10
La gestion à distance du R5715 est assurée par iDRAC10, la même plateforme que Dell propose en standard sur l'ensemble de sa gamme PowerEdge de 17e génération, notamment les PowerEdge R770 et R7725 que nous avons testés précédemment. L'interface étant identique pour toute la gamme, les administrateurs familiarisés avec iDRAC sur d'autres plateformes PowerEdge s'y retrouveront immédiatement.
Le tableau de bord iDRAC10 offre un aperçu complet et instantané de l'état de santé de chaque sous-système principal : état du système, processeur, mémoire, refroidissement, stockage, tensions, alimentations, batteries et détection d'intrusion. L'unité de test indique que tous les sous-systèmes étaient opérationnels au moment du test. Les informations système et les détails de la version du firmware sont affichés directement sur le tableau de bord, ainsi que l'état de la licence, qui, sur l'unité de test, est confirmé comme étant de type Entreprise. Le panneau « Résumé des tâches » suit les tâches en attente, en cours et terminées. L'unité de test affiche les tâches terminées d'un cycle de provisionnement initial, dont quelques-unes avec des erreurs et une ayant échoué, ce qui est typique d'un nouveau déploiement.
En explorant la section « Environnements système », vous accédez aux détails du refroidissement, notamment l'état de chaque ventilateur, les vitesses PWM, les paramètres du profil thermique et les relevés de température d'entrée, le tout en temps réel. Cette fonctionnalité est particulièrement utile pour vérifier le flux d'air dans les configurations de racks denses ou pour diagnostiquer les problèmes thermiques sans avoir à accéder physiquement au serveur.
La visibilité de la consommation électrique suit le même principe. La section « Informations sur l'alimentation » détaille l'état du bloc d'alimentation, la consommation de courant et le taux d'utilisation, ainsi qu'un graphique historique glissant. Les administrateurs peuvent ainsi visualiser rapidement la consommation moyenne et maximale au fil du temps, ce qui est précieux pour la planification de la capacité et l'identification des pics de consommation liés à la charge de travail, sans avoir besoin d'un outil de surveillance supplémentaire.
Ensemble, ces points de vue font d'iDRAC10 une solution de gestion hors bande performante qui couvre l'intégralité du cycle de vie opérationnel du R5715, du déploiement initial à la surveillance quotidienne, le tout accessible à distance via un navigateur ou l'API RESTful Redfish.
Performances Dell PowerEdge R5715
Pour évaluer les performances du Dell PowerEdge R5715, nous l'avons comparé à son homologue 1U, le Dell PowerEdge R4715 . Les deux plateformes partagent la même configuration mémoire et la même architecture PowerEdge, ce qui en fait un point de comparaison évident. La principale différence entre les deux unités testées réside dans le processeur. Le R4715 était équipé d'un processeur AMD EPYC 9335 à 32 cœurs, tandis que le R5715 disposait d'un processeur AMD EPYC 9015 à 8 cœurs.
Il est important de noter que les deux plateformes prennent en charge la même gamme de processeurs EPYC série 9005 et peuvent être configurées avec l'une ou l'autre puce selon les besoins. La différence du nombre de cœurs entre ces deux unités se reflétera dans les résultats, mais ces derniers correspondent aux performances réelles de chaque plateforme, et non à une comparaison de performances maximales.
Afin de solliciter les processeurs des deux systèmes, nous avons utilisé un ensemble ciblé de tests de calcul. y-cruncher a permis d'évaluer le débit arithmétique brut et les performances en virgule flottante multithread. Blender a fourni une charge de travail de rendu réaliste, évolutive en fonction du nombre de cœurs disponibles et de la bande passante mémoire. La suite de tests Phoronix a complété cet ensemble avec une collection plus large de charges de travail gourmandes en ressources processeur, offrant ainsi une vision plus complète des performances de calcul soutenues sur les deux plateformes.
Spécifications du système de test
- Plate-forme: Dell PowerEdge R5715
- CPU: AMD EPYC 9015 unique
- Mémoire: 384GB DDR5
- Stockage: Boss RAID1
croque-y
y-cruncher est un programme multithread et évolutif capable de calculer Pi et d'autres constantes mathématiques jusqu'à des milliers de milliards de chiffres. Depuis son lancement en 2009, il est devenu une application de benchmarking et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel informatique.
