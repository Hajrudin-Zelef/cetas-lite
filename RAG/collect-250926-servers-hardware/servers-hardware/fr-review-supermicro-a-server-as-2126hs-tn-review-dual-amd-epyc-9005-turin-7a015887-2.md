---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-2
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "distribution", "gpu", "nvidia"]
source: docs/RAG/clean4/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [78, 104]
sha256: 5752258103b6d10a7991be00792c9b69337fa304a80b43f29096979e79393228
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

Conception et montage du Supermicro AS-2126HS-TN
Le serveur Supermicro A+ AS-2126HS-TN est une plateforme 2U double socket compacte, conçue pour les configurations à très grand nombre de cœurs, processeurs à TDP élevé et accélérateurs multiples. Son architecture interne optimise la circulation de l'air de l'avant vers l'arrière, assurant un équilibre entre efficacité thermique, évolutivité de la consommation et facilité de maintenance pour les sous-systèmes de calcul, de mémoire, de stockage et d'extension.
Partie avant : Rangement et prise d'air
La face avant du châssis est conçue pour un stockage haute densité et une circulation d'air optimale. Selon la configuration, le système prend en charge jusqu'à 8 ou 24 baies pour disques NVMe, offrant une connectivité PCIe directe pour les charges de travail de stockage hautes performances. Ces baies sont positionnées en amont du flux d'air principal, garantissant ainsi une distribution homogène de l'air frais entrant avant qu'il n'atteigne les zones de calcul et d'extension.
Cette configuration permet au stockage de fonctionner indépendamment des températures du processeur et de l'accélérateur, maintenant ainsi des performances constantes sous une charge d'E/S soutenue tout en minimisant les perturbations du flux d'air à l'intérieur du châssis.
Architecture de refroidissement du châssis central
Juste derrière le compartiment de stockage avant se trouve un ensemble de six ventilateurs internes haute performance remplaçables à chaud. Ces ventilateurs créent un canal de refroidissement direct, générant une pression statique suffisante pour alimenter deux processeurs AMD EPYC d'une enveloppe thermique allant jusqu'à 500 W et plusieurs cartes graphiques PCIe double largeur.
La vitesse des ventilateurs est contrôlée dynamiquement via des connecteurs PWM et surveillée par des tachymètres intégrés, tandis que la gestion thermique s'effectue au niveau système. Cette conception permet à la plateforme d'adapter son refroidissement en temps réel en fonction des températures du processeur, de la mémoire et du châssis, garantissant ainsi un fonctionnement stable pour une large gamme de charges de travail.
Cœur de calcul : processeurs et mémoire
Au cœur du système se trouve la carte mère Super H14DSH, compatible avec deux processeurs AMD EPYC séries 9005 ou 9004. Prenant en charge jusqu'à 384 cœurs et 768 threads, l'AS-2126HS-TN est parfaitement adapté aux charges de travail hautement parallèles telles que la virtualisation, le calcul haute performance (HPC) et l'inférence en intelligence artificielle.
La plateforme prend en charge les processeurs dont le TDP atteint 500 W, avec un refroidissement par air possible au-delà de 400 W sous certaines configurations thermiques. Chaque processeur est associé à un ensemble complet de canaux mémoire alimentant 24 emplacements DIMM DDR5, prenant en charge jusqu'à 6 To de mémoire ECC DDR5 RDIMM en configuration 1DPC. La vitesse de la mémoire évolue avec la génération du processeur, atteignant jusqu'à 6 400 MT/s avec les processeurs EPYC 9005 et 4 800 MT/s avec les processeurs EPYC 9004.
Cette disposition symétrique de la mémoire minimise la latence et assure un flux d'air équilibré sur toutes les barrettes DIMM, même dans des configurations entièrement occupées.
Stockage de démarrage et connectivité embarquée
Pour le système d'exploitation et le stockage de gestion, le système comprend deux emplacements M.2 PCIe 3.0 NVMe intégrés, isolés du fond de panier NVMe principal. Cette séparation permet aux administrateurs de dédier les baies avant au stockage d'applications ou de données tout en préservant l'indépendance thermique et logique du support de démarrage.
La gestion du réseau s'effectue via un emplacement PCIe 5.0 x16 AIOM (OCP 3.0), permettant des configurations réseau flexibles sans monopoliser les emplacements d'extension PCIe standard. Cette approche préserve l'espace PCIe disponible pour les accélérateurs ou les contrôleurs de stockage supplémentaires.
Extension et prise en charge GPU
La partie arrière du châssis est dédiée à l'extension PCIe et à la prise en charge des accélérateurs. La carte mère AS-2126HS-TN prend en charge jusqu'à trois GPU PCIe double largeur, ce qui la rend idéale pour les environnements de calcul intensifs. Parmi les accélérateurs compatibles figurent les GPU NVIDIA PCIe tels que les H100 NVL, RTX 6000 Ada Generation et L4, ainsi que les accélérateurs AMD Instinct MI210.
Tous les emplacements d'extension sont PCIe Gen 5 et agencés de manière à garantir une ventilation optimale et une stabilité mécanique parfaite pour les cartes pleine hauteur et pleine longueur. La communication entre les GPU s'effectue via PCIe, ce qui assure la flexibilité et la compatibilité de la plateforme avec tous les fournisseurs, tout en prenant en charge un large éventail de configurations d'accélérateurs.
E/S arrière et gestion
Le panneau d'E/S arrière centralise l'accès à la gestion locale et hors bande. Il comprend un port LAN IPMI 1 GbE dédié, deux ports USB 3.0 et une sortie VGA pour un accès direct à la console. La gestion s'effectue via IPMI 2.0, avec prise en charge KVM sur LAN et des supports virtuels intégrée à la suite logicielle de gestion Supermicro, qui inclut SuperCloud Composer, Supermicro Server Manager et IPMIView.
Distribution de l'énergie et flexibilité électrique
L'alimentation est assurée par des blocs d'alimentation redondants à haut rendement (titane) disponibles en configurations de 1 200 W, 1 300 W, 1 600 W, 2 000 W et 2 600 W. Cette large gamme d'alimentations permet à la plateforme d'évoluer des configurations à forte densité de processeurs aux systèmes entièrement équipés de GPU, sans surdimensionnement ni sous-dimensionnement de la capacité d'alimentation.
Ce système prend en charge les tensions d'entrée de 120 V CA, 240 V CA et 48-60 V CC, ce qui le rend adapté aux baies d'entreprise traditionnelles, aux centres de données haute densité et aux environnements de télécommunications. La gestion de l'alimentation ACPI/APM intégrée permet une reprise contrôlée après une coupure de courant, un comportement configurable à la mise sous tension et une fonction de priorité sur le bouton d'alimentation. Grâce à sa conception à alimentation redondante, la plateforme offre à la fois une efficacité énergétique optimale et une grande fiabilité opérationnelle sous charge soutenue.
Sécurité, surveillance et fiabilité
Le contrôleur AS-2126HS-TN intègre un cadre de sécurité et de surveillance complet, comprenant TPM 2.0, une racine de confiance matérielle, le démarrage sécurisé, un firmware signé cryptographiquement et la restauration automatique du firmware. Les protections BMC en cours d'exécution et les fonctions de verrouillage du système contribuent à garantir l'intégrité de la plateforme tout au long de son cycle de vie.
La surveillance de l'état du système couvre le processeur, la mémoire, les lignes d'alimentation, la vitesse des ventilateurs et la température du châssis, avec une télémétrie en temps réel intégrée aux outils de gestion de Supermicro. Ceci permet une surveillance proactive et une réponse rapide aux anomalies thermiques ou d'alimentation dans les environnements de production.
Gestion Supermicro (BMC)
