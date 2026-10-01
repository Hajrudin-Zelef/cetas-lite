---
id: collect-261001-general-networking/general-networking/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a-2
title: "fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a.md
source_anchor: ""
source_lines: [20, 44]
sha256: 87f7a913dc59ec28a77b678afc65a2259fff372f7878ef28ee850829899c063c
---

# fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a

La première phase de l'effort de SAN autonome consiste à développer un élément d'architecture appelé Fabric Performance Impact Notifications (FPIN). Les notifications de matrice créent un mécanisme pour notifier les appareils du réseau des événements se produisant dans la matrice qui aideront à prendre des décisions en matière de résilience. Deux types de conditions d'erreur peuvent se produire. Les erreurs logiques peuvent souvent être récupérées par une nouvelle tentative ou une réinitialisation logique et sont relativement peu perturbatrices pour le système. D'autre part, les erreurs physiques nécessitent souvent le même type d'intervention pour terminer la réparation. Les erreurs intermittentes sont plus difficiles à résoudre et peuvent prendre du temps.
Avec les notifications de matrice, la matrice (ou le périphérique final) détecte le problème physique intermittent, surveille pour voir si l'erreur est persistante et, si c'est le cas, génère un message qui est envoyé aux périphériques affectés par l'événement. Grâce à ces informations, la solution multivoie connaît l'emplacement et la nature de l'erreur physique et peut la « contourner ». L'administrateur informatique n'a pas besoin de s'impliquer dans l'identification, l'isolement ou le lancement de commandes de récupération pour effacer l'erreur.
Tous ces mécanismes sont contrôlés par les terminaux via l'échange de capacités et l'enregistrement des opérations. Les commutateurs Fibre Channel de la structure ont la visibilité d'autres composants de la structure qui leur permettent de collecter des informations sur le réseau de stockage, les périphériques connectés et l'infrastructure globale. Ces commutateurs échangent ces informations à travers la structure, créant une véritable vision du SAN autonome pour l'auto-apprentissage, l'auto-optimisation et l'auto-rétablissement.
Les notifications de matrice fournissent finalement des informations aux appareils de la matrice pour éliminer le gaspillage d'énergie lors du dépannage, de l'analyse et de la résolution des problèmes qui sont résolus automatiquement par les appareils qui corrigent les problèmes qui ont un impact sur les performances et les pannes.
VMware a jeté les bases pour fournir un SAN intelligent et autonome avec le processus de réception, d'affichage et d'activation des alertes pour les événements de matrice FC à l'aide de Fabric Performance Impact Notification (FPIN), une technologie standard de l'industrie.
Le FPIN est une trame de notification transmise par un port de matrice pour notifier à un équipement terminal une condition pour un autre port de sa zone. Les conditions comprennent ce qui suit :
- Problèmes d'intégrité des liens qui dégradent les performances
- Notifications de cadre perdu
- Problèmes d'encombrement
Grâce à un mécanisme de notification proactif, les problèmes de port peuvent être résolus rapidement et des actions de récupération peuvent être définies pour atténuer les temps d'arrêt.
Figure 1 : vSphere 8.0 avec Marvell QLogic FC s'enregistre et reçoit des notifications de structure indiquant une sursouscription dans le SAN
Figure 2 : vSphere 8.0 avec Marvell QLogic FC s'enregistre et reçoit des notifications de structure indiquant une détérioration de l'intégrité des liens.
Les adaptateurs HBA Marvell QLogic Enhanced 16GFC, Enhanced 32GFC et 64GFC sont entièrement intégrés à vSphere 8.0 et prennent en charge la technologie de notification de fabric qui sert de bloc de construction pour les SAN autonomes.
Productivité avec vVols
VMware a mis l'accent sur vVols ces dernières versions de vSphere. Avec vSphere 8.0, le stockage de base a intégré la prise en charge de vVols pour NVMe-oF, la prise en charge de FC-NVMe étant initialement limitée. VMware continuera toutefois de valider et de prendre en charge les autres protocoles compatibles avec vSphere NVMe-oF. La nouvelle spécification vVols, le framework VASA/VC, est disponible ici.
Avec l'industrie et de nombreux fournisseurs de baies ajoutant la prise en charge de NVMe-oF pour des performances améliorées et une latence plus faible, VMware voulait s'assurer que les vVols restent à jour avec les technologies de stockage récentes.
Outre les performances améliorées, la configuration des vVols NVMe-oF est simplifiée. Une fois VASA enregistré, la configuration sous-jacente est effectuée en arrière-plan ; il ne reste plus qu’à créer la banque de données. VASA gère toutes les connexions aux points de terminaison de protocole virtuels (vPE). Les clients peuvent désormais gérer les baies de stockage NVMe-oF dans une banque de données vVols via la gestion basée sur les stratégies de stockage dans vCenter. vSphere 8 prend également en charge des espaces de noms et des chemins supplémentaires et améliore les performances de vMotion.
Suivi des machines virtuelles avec la technologie VM-ID
La virtualisation des serveurs a été le catalyseur d'un partage de liens accru, comme en témoigne Fibre Channel. Avec le nombre croissant de machines virtuelles (VM) dans le centre de données, les liens partagés transportent les données associées aux cœurs de processeur, à la mémoire et aux autres ressources système, en utilisant la bande passante maximale disponible. Les données envoyées depuis n'importe quelle machine virtuelle et d'autres systèmes physiques se mélangent. Ces données voyagent le long du même chemin que le trafic du réseau de stockage (SAN), donc tout semble identique et ne peut pas être considéré comme des flux de données individuels.
L'utilisation du VM-ID de Marvell QLogic (une solution de bout en bout utilisant le balisage de trames pour associer les différentes VM et leurs flux d'E/S sur le SAN) permet de déchiffrer chaque VM sur un lien partagé. QLogic a activé cette capacité sur ses derniers adaptateurs de bus hôte (HBA) 64GFC, Enhanced 32GFC et Enhanced 16GFC. Cette technologie dispose d'un moniteur de services d'application intégré, qui recueille l'ID unique au monde de VMware ESX. Il peut ensuite interpréter les différents identifiants de chaque machine virtuelle pour effectuer une surveillance intelligente.
VM-ID apporte une vision approfondie des E/S depuis la machine virtuelle d'origine jusqu'à la structure, ce qui donne aux gestionnaires de SAN la possibilité de contrôler et de diriger les services de niveau application vers chaque charge de travail virtuelle au sein d'un HBA QLogic Fibre Channel.
Figure 3 : Le moteur d'analyse du commutateur Brocade peut désormais afficher des statistiques par machine virtuelle en comptant les trames Fibre Channel marquées avec un ID de machine virtuelle individuel par le HBA Marvell Fibre Channel.
Augmentation des performances avec 64GFC
Les progrès de Fibre Channel se poursuivent depuis le lancement du protocole en 1988. Les premiers produits FC SAN, 1 Go FC, ont commencé à être commercialisés en 1997, et l'évolution se poursuit aujourd'hui, avec des produits 128 Go à l'horizon.
Tous les trois à quatre ans, la vitesse du FC double. Outre les avancées en matière de performances accrues, de nouveaux services tels que Fabric Services, StorFusion™ avec Universal SAN Congestion Mitigation, NPIV (virtualisation) et des services cloud ont été inclus. Les sociétés de mise en réseau et les équipementiers participent au développement de ces normes et continuent de travailler ensemble pour fournir des produits de réseau de stockage fiables et évolutifs.
