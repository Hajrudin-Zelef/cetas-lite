---
id: collect-261001-general-networking/general-networking/fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376-3
title: "fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376.md
source_anchor: ""
source_lines: [48, 101]
sha256: bad2c85447283d7f40cb2daf1e6d7974544f9988181c07f0b9ab1e5fcfe0188e
---

# fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376

- Associer des ID de machine virtuelle à des politiques de QoS de stockage spécifiques
- Mappage d'ID de machine virtuelle spécifiques à des ressources de stockage dédiées
- Utiliser les ID de VM pour l'équilibrage de charge
- Utiliser les ID de VM en conjonction avec des outils de surveillance
Ainsi, même si l'ID de VM n'a pas d'impact direct sur l'effet du mélangeur d'E/S, offrant une visibilité et un contrôle sur les machines virtuelles individuelles, les administrateurs peuvent personnaliser l'approvisionnement du stockage et hiérarchiser les ressources en fonction des ID de VM, ce qui entraîne des performances améliorées, une réduction des conflits et meilleure gestion globale de l’effet du mixeur d’E/S.
Marvell continue d'innover
Les cartes HBA Fibre Channel de Marvell offrent des performances et des fonctionnalités optimales pour le protocole Fibre Channel (FCP) et NVMe sur Fibre Channel ( FC-NVMe ). Conçues avec des chemins isolés pour chaque port, elles garantissent des performances à débit maximal par port et une fiabilité exceptionnelle. Ces adaptateurs offrent des millions d'IOPS, une latence de l'ordre de la microseconde et un débit maximal jusqu'à 64 Gbit/s par canal. Les technologies Marvell StorFusion™ et VM-ID simplifient le déploiement et l'intégration dans les réseaux SAN Fibre Channel.
Marvell StorFusion
La technologie Marvell StorFusion inclut des fonctionnalités avancées activées lorsqu'elles sont déployées avec les commutateurs Brocade et Cisco pris en charge. En combinant ces solutions, les administrateurs SAN peuvent profiter de fonctionnalités améliorées qui améliorent la disponibilité, accélèrent le déploiement et augmentent les performances du réseau.
À partir de la carte HBA QLE2690 et améliorée par les séries QLE2770 et QLE2870, les adaptateurs Marvell ont pris en charge plusieurs fonctionnalités de virtualisation basées sur des normes qui optimisent le déploiement, le dépannage et les performances des applications des serveurs virtuels.
La technologie Marvell VM-ID s'intègre facilement aux commutateurs Brocade et Cisco, permettant aux clients de surveiller et de gérer la qualité de service dans leurs réseaux de stockage Fibre Channel ; par exemple, l'équilibrage de charge des clusters de machines virtuelles avec le stockage pour garantir une utilisation efficace des ressources de stockage.
À partir de VMware ESXi 6.x, la prise en charge du balisage des demandes et des réponses d'E/S avec l'ID de VM de la machine virtuelle correspondante offre une visibilité complète au niveau de la VM.
De plus, la technologie Marvell StorFusion Universal SAN Congestion Mitigation ( USCM), basée sur la norme industrielle Fabric Performance Impact Notifications (FPIN), permet aux HBA et aux commutateurs du SAN d'identifier et d'atténuer les problèmes de congestion potentiels au sein du réseau. La prise en charge de la virtualisation N_Port ID (NPIV) permet à un seul port d'adaptateur FC de fournir plusieurs ports virtuels pour une évolutivité accrue du réseau. La technologie QoS basée sur le contrôle spécifique à la classe (CS_CTL) standard par port NPIV permet des contrôles et des garanties de bande passante multiniveaux par machine virtuelle. Par conséquent, les charges de travail critiques peuvent se voir attribuer une priorité plus élevée que le trafic de stockage moins sensible au facteur temps pour des performances optimales.
Balisage des trames d'E/S pour les machines virtuelles Fibre Channel
Sous les couvertures, la technologie FC VM-ID implique que le FC HBA balise les trames d'E/S avec des balises VM et que le commutateur FC lise ces balises et enregistre des statistiques par VM. Cela apporte plusieurs avantages, notamment une visibilité améliorée, une allocation des ressources et un dépannage. Le marquage des trames d'E/S avec des balises VM offre une visibilité améliorée du trafic généré par chaque VM individuelle. Cela permet une surveillance, une analyse et une gestion efficaces du trafic de stockage. Grâce au balisage des VM, les administrateurs obtiennent une vue claire de l'allocation des ressources, mettent en œuvre des politiques de QoS spécifiques à une VM individuelle, simplifient le dépannage et améliorent la sécurité et le contrôle d'accès.
La technologie Fibre Channel VM-ID nécessite la prise en charge de VM-ID dans le HBA FC, le commutateur FC et la baie de stockage. La majorité des HBA et commutateurs FC modernes prennent en charge VM-ID. Cependant, VM-ID n'est pris en charge que sur les baies de stockage de NetApp et PureStorage, ce qui pose un défi à l'adoption et au déploiement généralisés de cette technologie. L'innovation récente sur les commutateurs Brocade supprime la limitation VM-ID grâce à la technologie Tagless VM-ID ou VM-ID+.
Marvell VM-ID au Flash Memory Summit 2023
Lors du Flash Memory Summit (FMS), une vitrine internationale annuelle sur la mémoire et le stockage, nous avons visité le stand Marvell pour avoir un aperçu de ce sur quoi ils travaillaient. Parallèlement aux démonstrations de contrôleurs SSD, d'accélérateurs NVMe et de chipsets CXL, il y avait une démonstration en direct de l'ID VM Fibre Channel. La démonstration de VM-ID a attiré un grand nombre de clients et de partenaires. VM-ID pour Fibre Channel prend définitivement son essor.
VM-ID sans balise ou VM-ID+
Comme mentionné ci-dessus, VM-ID+ supprime la dépendance à la matrice de stockage pour prendre en charge le balisage VM-ID. VM-ID+ est configuré sur les ports de la structure SAN auxquels la matrice de stockage est connectée. Lorsque VM-ID+ est activé, la balise VM-ID des trames envoyées depuis l'hyperviseur vers la baie de stockage est supprimée par le commutateur Brocade Gen 7 au niveau du port de sortie connecté à la baie de stockage. Les trames que la baie de stockage envoie à l'hyperviseur ont la balise VM-ID ajoutée par la structure. Les commutateurs Fabric maintiennent le mappage et la collecte des données de télémétrie des VM.
Suivi des machines virtuelles avec VM-ID à partir de la ligne de commande
Les commandes du commutateur Brocade FC qui affichent les machines virtuelles actuelles et leurs statistiques fonctionnant au sein de la structure :
A. Exécutez la commande « appserver –show -all ». Le résultat est affiché ci-dessous ; les informations indiquent le nombre total de machines virtuelles (VM) en fonctionnement, identifiées par leur ID. Dans ce cas, trois VM sont présentes sur chaque hôte ESX. La dernière ligne affiche le nombre total de VM.
sw0-G720:FID128: > appserver –show-all
--------------------
Affichage des résultats pour Tissu
--------------------
N_ID de port : 010300
ID d'entité (ASCII) : 52 b3 0f fc 5a 05 47 a6-18 eb aa b4 b4 8f 9a 5f
ID d'entité (hexadécimal) : 0x35322062332030662066632035612030352034372061362d3138206562206161206234206234203866203961203566
ID de l'application : 0x00000010 (16)
Nom de l'entité :
Identifiant de l'hôte :
Données symboliques :
-------
N_ID de port : 010300
ID d'entité (ASCII) : 52 2c c3 8f c8 3f f5 75-a5 6c db bd 89 3a 95 13
ID d'entité (hexadécimal) : 0x35322032632063332038662063382033662066352037352d6135203663206462206264203839203361203935203133
ID de l'application : 0x00000012 (18)
Nom de l'entité :
Identifiant de l'hôte :
Données symboliques :
-------
N_ID de port : 010300
ID d'entité (ASCII) : 52 b1 ac 8d 2a aa 93 c4-5e 51 98 24 84 63 e0 c2
ID d'entité (hexadécimal) : 0x35322062312061632038642032612061612039332063342d3565203531203938203234203834203633206530206332
ID de l'application : 0x00000018 (24)
Nom de l'entité :
Identifiant de l'hôte :
Données symboliques :
-------
N_ID de port : 010800
ID d'entité (ASCII) : 52 bb 51 48 8a 5c 98 33-7a 74 c6 d5 27 05 58 49
ID d'entité (hexadécimal) : 0x35322062622035312034382038612035632039382033332d3761203734206336206435203237203035203538203439
