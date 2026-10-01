---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf-2
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "ethernet", "gpu"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [56, 84]
sha256: cc71f362433cea4ab7218aa0361b690a1aca12d2b3340b97042f001aa0cff6d0
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

Le R4715 est disponible avec trois configurations de baies avant : 4 baies SAS 3.5 pouces, 8 baies SAS/SATA 2.5 pouces ou 8 baies NVMe Gen4 U.2. Les deux dernières configurations partagent le même format 8 baies 2.5 pouces, le choix du fond de panier déterminant la compatibilité avec les protocoles de disques. Le RAID interne est géré par le contrôleur PERC H365i ou le contrôleur PERC H965i, plus performant. Le démarrage du système d'exploitation utilise un module BOSS-N1 DC-MHS dédié, qui isole complètement le volume de démarrage du pool de stockage de données et évite d'avoir à allouer de l'espace au système d'exploitation au sein d'une baie active.
L'unité testée était livrée avec un seul SSD SATA de 480 Go et deux disques NVMe U.2 de 1.92 To.
Processeur et mémoire
La carte mère R4715 prend en charge un processeur AMD EPYC série 9005 de 5e génération, jusqu'à 32 cœurs. La mémoire est gérée par 24 emplacements DIMM DDR5 compatibles uniquement avec les modules RDIMM (ni UDIMM ni LRDIMM). La capacité maximale est de 1.5 To avec des modules DIMM de 64 Go par emplacement, pour des vitesses allant jusqu'à 5 200 MT/s.
Refroidissement
Le R4715 est exclusivement refroidi par air. Le processeur est équipé d'un dissipateur thermique à cinq caloducs et d'un imposant empilement d'ailettes qui s'étend dans l'espace habituellement vide à côté du socket, augmentant ainsi la surface d'échange thermique. La ventilation est assurée par huit modules de ventilateurs haute performance remplaçables à chaud, disposés en groupes, qui brassent l'air de l'avant vers l'arrière du châssis. Aucun système de refroidissement liquide n'est disponible sur cette plateforme.
Tuning Moteur
Le R4715 prend en charge les alimentations redondantes remplaçables à chaud, disponibles en deux puissances : 800 W et 1 100 W, chacune proposée avec les certifications 80 PLUS Platinum ou Titanium. La certification FTR (Flex Titanium Rating) est également prise en charge sur l’ensemble de la gamme d’alimentations.
Extension et mise en réseau PCIe
La carte mère R4715 offre jusqu'à trois emplacements PCIe Gen5 avec connecteurs x16. L'emplacement 1 prend en charge les cartes pleine hauteur ou profil bas, l'emplacement 2 les cartes profil bas ou OCP 3.0, et l'emplacement 4 les cartes pleine hauteur ou profil bas. Deux emplacements pour cartes réseau OCP 3.0 (emplacements 2 et 5, Gen5 x16) couvrent les options d'adaptateurs 1 GbE, 10 GbE et 25 GbE. Pour les besoins en bande passante plus élevés, les cartes réseau PCIe AIC prennent en charge jusqu'à 100 GbE et 400 GbE, avec NDR VPI (400 GbE). La gestion hors bande s'effectue via un port Ethernet BMC 1 Gb dédié, isolant ainsi le trafic de gestion du plan de données. Cette plateforme ne prend pas en charge les GPU.
Panneau arrière
La connectique arrière comprend deux ports USB 3.1 Type-A, un port VGA et un port Ethernet 1 Gb dédié (BMC) pour la gestion iDRAC. Un port USB 3.1 Type-A supplémentaire est disponible en interne.
Gestion iDRAC10
La gestion à distance du R4715 utilise iDRAC10, la même plateforme que Dell propose en standard sur toute sa gamme PowerEdge de 17e génération, y compris les PowerEdge R770 et R7725 présentés précédemment. L'interface étant identique sur l'ensemble de la gamme, les administrateurs connaissant déjà iDRAC sur d'autres serveurs PowerEdge s'y retrouveront immédiatement.
Le tableau de bord iDRAC10 offre une vue d'ensemble complète de l'état de santé de tous les sous-systèmes clés : état du système, processeur, mémoire, refroidissement, stockage, tensions, alimentations, batteries et détection d'intrusion. L'unité de test affiche tous les sous-systèmes comme étant opérationnels au moment du test. Les informations système et les détails de la version du firmware sont affichés directement sur le tableau de bord, ainsi que le statut de la licence, qui est confirmé comme étant de type Entreprise sur l'unité de test. Le panneau « Résumé des tâches » suit les tâches en attente, en cours et terminées. L'unité de test affiche les tâches terminées d'un cycle de provisionnement initial, y compris un petit nombre d'erreurs et un échec, ce qui est typique d'un nouveau déploiement.
En explorant la section « Environnements système », vous accédez aux détails du refroidissement, notamment l'état de chaque ventilateur, les vitesses PWM, les paramètres du profil thermique et les relevés de température d'entrée, le tout en temps réel. Ceci est particulièrement utile pour vérifier le flux d'air dans les configurations de racks denses ou pour diagnostiquer les problèmes thermiques sans accès physique au serveur.
La visibilité de la consommation électrique suit le même principe. La section « Informations sur l'alimentation » détaille l'état du bloc d'alimentation, la consommation de courant et le taux d'utilisation, ainsi qu'un graphique historique glissant. Les administrateurs peuvent ainsi visualiser rapidement la consommation moyenne et maximale au fil du temps, ce qui est utile pour la planification de la capacité et la détection des pics de consommation liés à la charge de travail, sans avoir besoin d'un outil de surveillance supplémentaire.
Ensemble, ces points de vue font d'iDRAC10 une solution de gestion hors bande performante qui couvre l'intégralité du cycle de vie opérationnel du R4715, du déploiement initial à la surveillance quotidienne, le tout accessible à distance via un navigateur ou l'API RESTful Redfish.
Performances Dell PowerEdge R4715
Pour évaluer les performances du Dell PowerEdge R4715, nous l'avons comparé à son équivalent 2U, le Dell PowerEdge R5715 . Les deux plateformes partagent la même configuration mémoire et la même architecture PowerEdge, ce qui rend leur comparaison pertinente. La principale différence entre les deux unités testées réside dans le processeur : le R4715 est équipé d'un processeur AMD EPYC 9335 à 32 cœurs, tandis que le R5715 intègre un processeur AMD EPYC 9015 à 8 cœurs.
Il est important de noter que les deux serveurs prennent en charge la même gamme de processeurs EPYC série 9005 et peuvent être configurés avec l'une ou l'autre puce selon les besoins de la charge de travail. La différence du nombre de cœurs entre ces deux unités se reflétera dans les résultats, mais ces derniers indiquent les performances réelles de chaque plateforme, telles qu'elles sont livrées, et non une comparaison des performances maximales entre les deux.
Pour évaluer la capacité des processeurs sur l'ensemble des systèmes, nous avons utilisé un ensemble ciblé de tests de calcul. y-cruncher a mesuré le débit arithmétique brut et les performances en virgule flottante multithread. Blender a fourni une charge de travail de rendu réaliste, évolutive en fonction du nombre de cœurs disponibles et de la bande passante mémoire. La suite de tests Phoronix a complété cet ensemble de tests en ajoutant une plus grande variété de charges de travail gourmandes en ressources processeur, offrant ainsi une vision plus complète des performances de calcul soutenues sur les deux plateformes.
Spécifications du système de test
- Plate-forme: Dell PowerEdge R4715
- CPU: AMD EPYC 9335 monocœur
- Mémoire: 384GB DDR5
- Stockage: Boss RAID1
croque-y
y-cruncher est un programme multithread et évolutif capable de calculer Pi et d'autres constantes mathématiques jusqu'à des milliers de milliards de chiffres. Depuis son lancement en 2009, il est devenu une application de benchmarking et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel informatique.
