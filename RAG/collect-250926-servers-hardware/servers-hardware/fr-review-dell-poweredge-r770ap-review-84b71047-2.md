---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770ap-review-84b71047-2
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["datacenter", "ethernet", "intel"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [48, 74]
sha256: 8510c39d2d7c8730b2651e9a9873dc34db0f1f9fa1588afba189e9514874d8d7
---

# fr-review-dell-poweredge-r770ap-review-84b71047

| Caractéristiques de sécurité | Firmware signé cryptographiquement, chiffrement des données au repos (SED avec gestion de clés locale ou externe), démarrage sécurisé, vérification des composants sécurisés (contrôle d'intégrité matérielle), effacement sécurisé, racine de confiance matérielle, verrouillage du système (nécessite iDRAC10 Enterprise ou Datacenter), TPM 2.0 certifié FIPS/CC-TCG, détection d'intrusion dans le châssis | 
| Systèmes d'exploitation et hyperviseurs |  | 
| Systèmes d'exploitation/hyperviseurs pris en charge | Canonical Ubuntu Server LTS, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware vSAN / VMware ESXi*, Microsoft Windows, Microsoft Windows Server, Microsoft Windows Server Datacenter | 
Concevoir et construire
Le Dell PowerEdge 770AP est un serveur rack 2U de la 17e génération de serveurs Dell PowerEdge, reprenant le même design que le R770 que nous avons testé. Il mesure 3.42 cm de hauteur, 19.0 cm de largeur et 31.59 cm de profondeur. La façade avant est optionnelle. Le panneau avant intègre un accès direct à iDRAC, un port USB 2.0 Type-C, un bouton d'alimentation et un bouton d'identification du système.
Stockage
Le serveur 770AP prend en charge trois configurations de stockage. Il est livré avec jusqu'à 16 SSD NVMe Gen 5 x4 de 2.5 pouces, pour une capacité maximale de 245.76 To. Il est également possible d'opter pour jusqu'à 16 SSD NVMe Gen 5 x2 de 2.5 pouces, plafonnés à 245.76 To, ou encore jusqu'à 32 SSD NVMe EDSFF E3.S Gen 5, extensibles jusqu'à 491.52 To. Dans les configurations à 16 baies, Dell répartit les disques en deux groupes de huit, situés à gauche et à droite du serveur, la partie centrale servant d'entrée d'air.
En regardant plus en détail à l'intérieur du châssis, on constate que le 770AP présente un câblage NVMe direct et épuré. Les câbles relient directement le fond de panier du stockage au bord avant de la carte mère, ce qui raccourcit le trajet du signal et optimise l'organisation interne.
E/S arrière et réseau
Deux alimentations redondantes de 2 400 W sont fixées à l’arrière du 770AP, à chaque extrémité. Le module BOSS-N1 gère le démarrage et intègre deux disques de 480 Go pour le système d’exploitation.
Pour l'extension, le serveur offre jusqu'à cinq emplacements PCIe Gen 5 répartis sur les emplacements 2, 3, 5, 7 et 9, tous équipés de connecteurs x16 en configuration pleine hauteur. La connectivité réseau OCP 3.0 est assurée par deux cartes maximum : l'emplacement 4 prend en charge les interfaces Gen 5 x8 ou x16, et l'emplacement 10 offre une connexion x16 Gen 5 dédiée. Notre machine était livrée avec une carte OCP 200 GbE et plusieurs cartes 100 GbE, garantissant une bande passante réseau largement suffisante.
Les E/S arrière standard comprennent un port Ethernet BMC dédié, deux ports USB 3.1 Type-A et un port VGA.
Un examen plus attentif du module BOSS-N1 révèle la présence de deux disques de démarrage de 480 Go côte à côte, tous deux remplaçables à chaud et faciles d'accès et de remplacement en cas de besoin.
Une fois le capot supérieur et les carénages d'aération retirés, l'intérieur du R770AP se révèle propre et bien organisé. Six ventilateurs remplaçables à chaud brassent l'air à travers les larges dissipateurs thermiques, refroidissant ainsi les processeurs Xeon série 6900, dont la configuration double processeur et mémoire est disposée symétriquement. On remarque également les languettes bleues présentes sur le châssis, qui servent de guides pour le démontage, le retrait des câbles et l'accès aux composants.
Processeur
Une fois le processeur retiré, la taille imposante de la puce Intel Xeon série 6900 saute aux yeux. La carte mère R770AP utilise le socket LGA 7529, et notre modèle de test était équipé de deux processeurs Intel Xeon 6978P. Chaque puce affiche un TDP de 500 W et 120 cœurs, portant le nombre total de cœurs à 240 pour les deux sockets.
Refroidissement et mémoire
Pour gérer la dissipation thermique de 1 000 W du processeur par simple refroidissement à air, Dell a conçu un système de refroidissement spécifique. Les dissipateurs avant et arrière utilisent des ailettes horizontales et des caloducs pour une dissipation thermique efficace. La partie centrale, quant à elle, est dotée d'un empilement d'ailettes verticales qui augmente le temps de contact et la surface d'échange thermique, permettant ainsi aux ventilateurs d'évacuer plus efficacement la chaleur avant sa sortie du châssis. Au total, 24 emplacements DIMM sont intégrés au système de refroidissement : chaque processeur est flanqué de 12 emplacements, soit six de chaque côté.
Tuning Moteur
Le R770AP prend en charge quatre options d'alimentation, toutes certifiées 80 Plus Titanium et redondantes à chaud : 1 500 W, 1 800 W, 2 400 W et 3 200 W. Avec une consommation pouvant atteindre 1 000 W pour les seuls processeurs, la configuration de base de 1 500 W offre une marge de manœuvre très réduite une fois les disques et les cartes d'extension pris en compte. Notre modèle était équipé d'une alimentation de 2 400 W, affichant un rendement de 96 %, ce qui représente le minimum pratique pour une configuration de stockage complète.
Gestion iDRAC 10
La gestion à distance du R770AP est assurée par iDRAC10, la même plateforme que Dell propose en standard sur l'ensemble de sa gamme PowerEdge de 17e génération, notamment les PowerEdge R770 et R7725 que nous avons testés précédemment. L'interface étant identique pour toute la gamme, les administrateurs connaissant déjà iDRAC sur d'autres plateformes PowerEdge s'y retrouveront facilement.
Le tableau de bord iDRAC10 offre un aperçu complet et instantané de l'état de santé de chaque sous-système principal : état du système, processeur, mémoire, refroidissement, stockage, tensions, alimentations, batteries et détection d'intrusion. L'unité de test indique que tous les sous-systèmes étaient opérationnels au moment du test. Les informations système et les détails de la version du firmware sont affichés directement sur le tableau de bord, ainsi que l'état de la licence, qui, sur l'unité de test, est confirmé comme étant de type Entreprise. Le panneau « Résumé des tâches » suit les tâches en attente, en cours et terminées. L'unité de test affiche les tâches terminées d'un cycle de provisionnement initial, dont quelques-unes avec des erreurs et une ayant échoué, ce qui est typique d'un nouveau déploiement.
En explorant la section « Environnements système », vous accédez aux détails du refroidissement, notamment l'état de chaque ventilateur, les vitesses PWM, les paramètres du profil thermique et les relevés de température d'entrée, le tout en temps réel. Cette fonctionnalité est particulièrement utile pour vérifier le flux d'air dans les configurations de racks denses ou pour diagnostiquer les problèmes thermiques sans avoir à accéder physiquement au serveur.
La visibilité de la consommation électrique suit le même principe. La section « Informations sur l'alimentation » détaille l'état du bloc d'alimentation, la consommation de courant et le taux d'utilisation, ainsi qu'un graphique historique glissant. Les administrateurs peuvent ainsi visualiser rapidement la consommation moyenne et maximale au fil du temps, ce qui est précieux pour la planification de la capacité et l'identification des pics de consommation liés à la charge de travail, sans avoir besoin d'un outil de surveillance supplémentaire.
Ensemble, ces points de vue font d'iDRAC10 une solution de gestion hors bande performante qui couvre l'intégralité du cycle de vie opérationnel du R770AP, du déploiement initial à la surveillance quotidienne, le tout accessible à distance via un navigateur ou l'API RESTful Redfish.
Performances du Dell PowerEdge R770AP
