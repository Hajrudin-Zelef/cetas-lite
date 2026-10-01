---
id: collect-261001-general-networking/general-networking/fr-review-eaton-g3-universal-pdu-review-bfd817e8-1
title: "fr-review-eaton-g3-universal-pdu-review-bfd817e8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-general-networking/fr-review-eaton-g3-universal-pdu-review-bfd817e8.md
source_anchor: ""
source_lines: [1, 35]
sha256: a912e1039d34b0e4b4e37a021db70a6c944a2da04744895179fc59e1fdf3f2a8
---

# fr-review-eaton-g3-universal-pdu-review-bfd817e8

Le PDU universel Eaton G3 (UPDU) offre aux administrateurs informatiques la possibilité d'alimenter presque tous les équipements dans un format compact. Le PDU est une solution unique pour répondre aux besoins les plus divers de votre centre de données, pouvant accueillir une alimentation monophasée de 5 kVA à une alimentation triphasée de 23 kVA dans des environnements 3 V et 120 V avec 230 prises dans un châssis 42U. Sur ces 42 points de vente, 42 sont des C21 et les 14 autres sont des C21.
L'ajout de prises C39 est une véritable innovation pour le secteur, car elles permettent de brancher des fiches C14 et C20 sur une seule prise ! Cela peut paraître un changement mineur par rapport aux modèles précédents, mais grâce à sa conception compacte et aux prises C39, ce nouveau produit permet une meilleure organisation des centres de données avec moins d'unités de distribution d'énergie (PDU). Il remplace avantageusement les deux PDU existants qui offraient auparavant la même sélection de ports. Cette innovation place le nouveau produit d'Eaton en position de force face à la concurrence.
Entrée PDU universelle Eaton G3
Le G3 UPDU dispose d'une entrée universelle située à l'extrémité du PDU et répond à un large éventail de besoins en alimentation d'applications. De plus, il contient six banques de prises à code couleur et les disjoncteurs correspondants pour simplifier l'équilibrage de charge. Il existe un total de 12 options de câbles d'alimentation pour l'entrée universelle proposées à la vente d'une longueur de 10 pieds.
Options de câble d'entrée PDU universel Eaton G3 :
| Type de fiche | Numéro de catalogue | Disjoncteur de panneau | Courant déclassé | Tension | Région | Phase de puissance | Puissance maximale par PDU | 
| L6-30P | CBL350-10 | 30A | 24A | 208V | AMER | Single | 5.0kVA | 
| L21-30P | CBL351-10 | 30A | 24A | 208/120 V ÉTOILE | AMER | 3 | 8.6kVA | 
| 332P6W | CBL354-10 | 30A 32A | 24A 32A | 240V 230V | AMER EMEA / APAC | Single Single | 5.8kVA 7.4kVA | 
| 516P6W | CBL355-10 | 20A 16A | 16A 16A | 240/415 V ÉTOILE 230/400 V ÉTOILE | AMER EMEA / APAC | 3 3 | 11.5kVA 11.0kVA | 
| 460P9W | CBL356-10 | 60A | 48A | Delta 208 V | AMER | 3 | 17.3kVA | 
| Couettes | CBL357-10 | 60A | 48A | Delta 208 V | AMER | 3 | 17.3kVA | 
| 532P6W | CBL358-10 | 30A 32A | 24A 32A | 240/415 V ÉTOILE 230/400 V ÉTOILE | AMER EMEA / APAC | 3 3 | 17.3kVA 22.1kVA | 
| 560P6W | CBL360-10 | 60A | 32A | 240 / 415V | AMER | 3 | 23.0kVA | 
| 360P6W | CBL362-10 | 60A | 48A | 240V | AMER | Single | 11.5kVA | 
| CS8365 | CBL364-10 | 50A | 40A | Delta 208 V | AMER | 3 | 14.4kVA | 
| L15-30P | CBL365-10 | 30A | 24A | Delta 208 V | AMER | 3 | 8.6kVA | 
| 560P9W | CBL366-10 | 60A | 48A | 208/120 V ÉTOILE | AMER | 3 | 17.3kVA | 
Construire et concevoir
En regardant la partie centrale de l'onduleur, vous remarquerez le module de gestion et de contrôle du réseau. Cela comprend une carte réseau évolutive intégrée, des commandes à bouton-poussoir et un écran avec des lectures au niveau de la prise et de la banque. Cette baie comprend également des ports série/Ethernet et un port USB pour les mises à niveau du micrologiciel.
L'écran peut pivoter dans le menu pour plus de lisibilité, quelle que soit l'orientation dans laquelle vous choisissez de le monter. Vous pouvez également contrôler l'état d'alimentation de prises individuelles, voire de banques entières à la fois. Pour un aspect de sécurité physique, Eaton vous permet de mettre une épingle sur cet écran pour empêcher quiconque de modifier la configuration au sein de l'interface d'affichage.
L'ensemble de ce module réseau peut être retiré et remplacé sans interférer avec la puissance de sortie de la PDU. Cela permet un remplacement facile si vous rencontrez des problèmes avec le module réseau, permettant également une mise à niveau si un module plus récent est publié.
Facteur de forme du PDU universel Eaton G3
Le bloc d'alimentation testé dans cet article est le modèle EMAGU23X-3 . Nous avons rencontré un léger problème d'encombrement, car il est environ 6 cm plus haut que les blocs EMA107-10 et EMI104-10 de notre laboratoire. Ce problème ne se posait que sur les anciennes baies et armoires de la série S d'Eaton et a depuis été résolu grâce à l'espace plus important prévu pour les blocs d'alimentation dans les baies de la série R. Eaton a fourni un support de compatibilité pour installer ce bloc d'alimentation dans les baies de la série S, mais ce support peut être difficile à trouver.
L'utilisation de ce support de compatibilité place l'unité de distribution d'alimentation (PDU) plus au centre du rack, gênant ainsi l'accès aux ports, aux rails et autres éléments. Autre problème d'encombrement : l'extrémité volumineuse des câbles universels. La prise, de la taille d'une canette de Red Bull, voire plus grande, obstrue également l'accès à la partie supérieure du rack. Bien que ces problèmes aient été résolus depuis sur les racks de la série R , il était important de signaler ceux rencontrés dans notre environnement.
Interconnectivité
Comme les PDU gérés précédents, jusqu'à 8 de ces PDU peuvent être connectés en série sous une seule adresse IP à l'aide des répartiteurs Ethernet inclus, permettant une gestion plus facile et moins d'espace réseau.
Gestion Web universelle des PDU Eaton G3
L'UPDU G3 peut être gérée en ligne via un navigateur web ou par intégration SNMP dans un système DCIM . Elle peut également être gérée via la suite Brightlayer Data Centers d'Eaton . Des systèmes de gestion tiers sont également compatibles. Les commandes de gestion peuvent être envoyées instantanément à plusieurs unités grâce à la configuration en masse.
La gestion à distance des PDU réduit les visites sur site en envoyant des alertes automatisées sur les problèmes d'alimentation et en redémarrant à distance une PDU ou des prises individuelles. C'est parfait que vous ayez un seul centre de données ou même un projet global sur plusieurs sites.
Tout d'abord, lorsque vous vous connectez directement à l'interface de la PDU, vous obtenez immédiatement un écran d'accueil qui affiche l'état des ports, l'ampérage actuel et le pourcentage de charge de chaque phase utilisée, le facteur de crête, la tension d'entrée et la puissance. résumés pour l’unité.
Deuxièmement, nous avons mentionné plus tôt les banques de six prises du PDU, qui peuvent également toutes être surveillées individuellement dans l'interface. Chaque banque affiche sa charge en ampères, son pourcentage de capacité, sa consommation active en watts et sa consommation totale en KWH pour une certaine période de temps.
Troisièmement, pour une vue encore plus granulaire, les prises individuelles elles-mêmes peuvent être surveillées pour l'état de l'alimentation (qui est également affiché par une LED rouge ou verte à côté de la prise respective), la consommation d'énergie active et la consommation totale.
Enfin, pour une organisation et une surveillance facile, chaque prise, ou un groupe de 8 maximum, peut être relié comme un appareil. Les appareils peuvent être répartis sur des PDU connectées en série si vous avez quelque chose comme une configuration gauche-droite. Il s'agit d'une aide précieuse pour les serveurs ou autres unités disposant de plusieurs alimentations, vous permettant de visualiser la consommation ou de basculer l'état ou les horaires d'alimentation en même temps.
Surveillance environnementale
