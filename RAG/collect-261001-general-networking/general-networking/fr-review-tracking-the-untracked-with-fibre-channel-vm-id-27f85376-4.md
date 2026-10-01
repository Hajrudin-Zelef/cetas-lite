---
id: collect-261001-general-networking/general-networking/fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376-4
title: "fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376.md
source_anchor: ""
source_lines: [102, 152]
sha256: 375033741c5c53f28a0a908a6740ecf7643d38cd9d550d2b5215ab18dae89eba
---

# fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376

ID de l'application : 0x00000010 (16)
Nom de l'entité :
Identifiant de l'hôte :
Données symboliques :
-------
N_ID de port : 010800
ID d'entité (ASCII) : 52 36 64 98 87 5d a5 c6-02 38 0a d7 85 42 3b 4b
ID d'entité (hexadécimal) : 0x35322033362036342039382038372035642061352063362d3032203338203061206437203835203432203362203462
ID de l'application : 0x00000012 (18)
Nom de l'entité :
Identifiant de l'hôte :
Données symboliques :
-------
N_ID de port : 010800
ID d'entité (ASCII) : 52 de 5b 4f a9 9f 98 12-65 4f e7 ca c5 78 c2 3c
ID d'entité (hexadécimal) : 0x35322064652035622034662061392039662039382031322d3635203466206537206361206335203738206332203363
ID de l'application : 0x00000018 (24)
Nom de l'entité :
Identifiant de l'hôte :
Données symboliques :
-------
Le serveur d'applications affiche six entrées.
B. Brocade Analytics Engine a rapporté des statistiques pour les machines virtuelles
Utilisez les commandes ci-dessous pour détailler les métriques d'E/S pour chaque VM
C. Vérification et configuration d'un port de commutateur Brocade à cibler pour le VMID sans étiquette (pour les baies de stockage autres que NetApp et PureStorage) :
sw0-G720:FID128:admin> portcfgappheader -h
Usage:
portCfgAppHeader <[slot/]port> – activer/– désactiver
D. L'exécution de la commande sur un port précédemment configuré affichera ce qui suit :
portcfgappheader 26 – activer
Même configuration pour le port 26
Le mode d'investigation de SANnav fournit un aperçu des performances opérationnelles de chaque machine virtuelle. Il collecte et stocke les statistiques de performances du SAN et les données de télémétrie, puis fournit des graphiques de séries chronologiques clairs et intuitifs qui tracent les principales mesures de trafic. Il comprend des détails sur les violations MAPS pour les ports, les liaisons et les lignes réseau, les tunnels et circuits d'extension et les flux pour aider les utilisateurs à comprendre et à étudier les comportements complexes des modèles de trafic. De plus, il peut collecter des métriques plus fréquemment et en temps quasi réel (à intervalles de 10 secondes) pour les ports sélectionnés.
Une fois connecté, la vue du tableau de bord SANnav affichera la structure/les commutateurs en cours de gestion.
1. Accédez à Inventaire—> Flux —> sélectionnez le filtre défini par l'utilisateur :
« All-VMs-flows » et les flux détaillés seront affichés comme illustré ci-dessous :
2. Cliquez sur l'icône (…) dans le coin supérieur droit pour afficher un menu déroulant, puis cliquez sur « Sélection en masse ».
3. Dans la zone de bannière au-dessus de toutes les machines virtuelles, cochez la case pour sélectionner toutes les machines virtuelles.
4. Une fois les machines virtuelles sélectionnées, cliquez sur le bouton Action en haut à droite, puis sélectionnez « Analyser ».
Dans la fenêtre Mode Investigation :
5. Cochez les cases correspondant à chacune des VM et cliquez sur l'option « Débit de données de lecture » dans le panneau de gauche.
6. Ensuite, cliquez sur la flèche vers le bas à côté de « 30 dernières minutes ».
7. Une autre fenêtre « Sélectionnez la plage de dates » s’affiche. Cliquez sur l’option prédéfinie « Dernière semaine » à gauche, puis sur « Appliquer ».
8. Cliquez ensuite sur la flèche vers le bas à côté de « Intervalle : 5 minutes » et sélectionnez l'option « 6 heures ».
9. Une vue d'informations détaillée sera affichée pour les six machines virtuelles et le trafic associé présent dans la structure. Déplacer le curseur et survoler un index temporel graphique spécifique affichera les performances de débit de données en lecture pour chacun.
10. Cliquez ensuite sur la flèche vers le bas à côté de « Intervalle : 5 minutes » et sélectionnez l'option « 6 heures ».
11. Une vue d'informations détaillée sera affichée pour les six machines virtuelles et le trafic associé présent dans la structure. Déplacer le curseur et survoler un index temporel graphique spécifique affichera les performances de débit de données en lecture pour chacun.
12. Cliquez sur l'option « Temps réel » en haut à droite pour afficher et actualiser les détails toutes les dix secondes.
Les profils de performances et d'E/S des machines virtuelles individuelles activés par VM-ID et affichés dans les écrans SANnav ci-dessus permettent aux administrateurs SAN et de stockage d'obtenir une visibilité sur les modèles de trafic pour chaque VM.
La technologie Marvell QLogic VM-ID et Brocade SANnav sont des leaders de l'innovation en matière de gestion moderne des données. Avec VMware ESXi, les capacités transparentes de déploiement et d'orchestration de VM-ID et les outils complets de gestion du stockage de SANnav, les entreprises peuvent naviguer en toute confiance et facilement dans les complexités des environnements virtualisés.
Ces solutions permettent aux organisations d'exploiter tout le potentiel de leur infrastructure de données, garantissant des performances, une efficacité et une adaptabilité optimales. À mesure que la technologie évolue, Marvell QLogic VM-ID et SANnav restent des partenaires fidèles dans le cheminement vers une gestion rationalisée des données et une excellence opérationnelle améliorée.
Ce rapport est parrainé par Marvell. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
