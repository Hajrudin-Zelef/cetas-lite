---
id: collect-261001-general-networking/general-networking/fr-review-dell-perc13-70dd67e6-2
title: "fr-review-dell-perc13-70dd67e6"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "dram", "gpu", "nvidia"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-perc13-70dd67e6.md
source_anchor: ""
source_lines: [44, 66]
sha256: 32dc3ae8f39e7024447bdb2317b8455e42cc2ca91280eaaaabefa6747a4d3533
---

# fr-review-dell-perc13-70dd67e6

Le H975i implémente une architecture de sécurité complète, allant de l'attestation matérielle au niveau du silicium au chiffrement intégral des données, mis en place avec les disques SED. La racine de confiance matérielle (HRC) établit une chaîne immuable de vérification cryptographique, de la ROM de démarrage interne à chaque composant du micrologiciel, garantissant ainsi que seuls les micrologiciels certifiés Dell authentifiés peuvent s'exécuter sur le contrôleur. Cette sécurité matérielle s'étend à l'implémentation du protocole SPDM (Security Protocol and Data Model), où chaque contrôleur contient un certificat d'identité de périphérique unique permettant à l'iDRAC d'effectuer une vérification d'authentification en temps réel. Le contrôleur étend la protection cryptographique au-delà des scénarios traditionnels de données au repos, en incluant la mémoire cache. Il conserve les clés de chiffrement dans des zones mémoire sécurisées, inaccessibles aux micrologiciels non autorisés. Ainsi, les données sensibles restent protégées, qu'elles soient stockées sur les disques ou traitées activement dans le cache.
La protection de l'alimentation du H975i constitue une autre évolution significative par rapport aux systèmes traditionnels alimentés par batterie grâce à l'intégration d'un supercondensateur. Ce dernier assure une alimentation instantanée en cas de coupure de courant imprévue, garantissant un vidage complet et chiffré du cache vers un stockage non volatile, où les données restent protégées indéfiniment. De plus, contrairement aux systèmes alimentés par batterie qui nécessitent 4 à 8 heures pour les cycles d'apprentissage, le supercondensateur du H975i effectue son cycle d'apprentissage transparent en 5 à 10 minutes sans dégradation des performances lors de l'étalonnage. Cette conception élimine les frais de maintenance et les problèmes de dégradation inhérents aux solutions sur batterie, tout en offrant une fiabilité supérieure pour la protection des données critiques.
Surveillance et gestion intégrées
Le contrôleur RAID PERC13 de Dell, comme de nombreuses solutions RAID de Dell, peut être géré et surveillé de nombreuses manières, notamment lors du démarrage de la plate-forme via la configuration du système dans le BIOS, via l'interface graphique Web iDRAC, l'utilitaire PERC12 et même l'interface utilisateur et la CLI Dell OpenManage.
Gestion du contrôleur iDRAC
Dans l'interface de gestion iDRAC, l'onglet « Contrôleurs » offre un aperçu du matériel de stockage du serveur. À côté de la carte BOSS, vous trouverez les deux contrôleurs PERC H975i, ainsi que des informations sur les versions du firmware, la mémoire cache et l'état de la batterie. Ce résumé vous permet de vérifier rapidement l'état de préparation et la configuration des contrôleurs sans avoir à accéder au BIOS ni à utiliser d'outils CLI.
L'onglet « Disques virtuels » d'iDRAC affiche les baies de stockage créées, y compris leur niveau RAID, leur taille et leur politique de mise en cache. Ce système répertorie deux groupes RAID-10, tous basés sur des SSD. Depuis cette vue, les administrateurs peuvent vérifier la connexion des volumes, créer de nouveaux disques virtuels ou utiliser le menu Actions pour ajuster ou supprimer les configurations existantes.
Utilitaire de configuration du contrôleur RAID
L'image ci-dessus montre un exemple d'accès à l'utilitaire de configuration avant PERC H975i sur la plateforme PowerEdge R7715. Cette interface vous permet de gérer tous les paramètres clés du contrôleur RAID, notamment la gestion de la configuration, la gestion des contrôleurs et des périphériques. Cet utilitaire simplifie la configuration des disques virtuels et la surveillance des composants matériels directement pendant le démarrage de la plateforme.
Après avoir sélectionné le niveau RAID, nous passons au choix des disques physiques pour la matrice. Dans cet exemple, tous les SSD NVMe disponibles sont répertoriés et marqués comme compatibles RAID. Nous sélectionnons plusieurs disques Dell DC NVMe de 3.2 Tio dans le pool de capacité non configuré. Des filtres tels que le type de support, l'interface et la taille du secteur logique permettent d'affiner la sélection. Une fois les disques souhaités sélectionnés, nous pouvons cliquer sur « OK » pour finaliser la sélection et poursuivre la création du disque virtuel.
Avant de finaliser la création du disque virtuel, le système affiche un avertissement confirmant que toutes les données des disques physiques sélectionnés seront définitivement supprimées. Pour continuer, cochez la case « Confirmer » et sélectionnez « Oui » pour autoriser l'opération. Cette mesure de sécurité permet d'éviter toute perte accidentelle de données lors de la création du RAID.
Une fois le disque virtuel créé, il apparaît dans le menu « Gestion des disques virtuels ». Dans cet exemple, notre nouveau disque virtuel RAID 5 affiche une capacité de 43.656 Tio et le statut « Prêt ». En quelques étapes simples, le stockage est configuré et prêt à l'emploi.
Bien que l'utilitaire de configuration du BIOS PERC et l'interface iDRAC offrent des options intuitives pour la gestion locale et à distance, Dell propose également un outil en ligne de commande puissant : PERC CLI (perccli2). Cet utilitaire est compatible avec Windows, Linux et VMware, ce qui le rend idéal pour la création de scripts, l'automatisation ou la gestion des contrôleurs PERC dans des environnements sans interface graphique. Dell met également à disposition une documentation détaillée sur l'installation et l'utilisation des commandes de PERC CLI sur son site de support.
Tests de performances du Dell PERC13
Avant de nous lancer dans les tests de performances, nous avons préparé notre environnement avec la plateforme Dell PowerEdge R7715 configurée avec deux contrôleurs frontaux PERC H975i. Ceux-ci étaient associés à trente-deux disques Dell NVMe de 3.2 To, chacun pouvant atteindre 12,000 5,500 Mo/s en lecture séquentielle et 128 13 Mo/s en écriture séquentielle avec des blocs de XNUMX Kio. Cette base hautes performances nous permet de repousser les limites du débit du contrôleur PERCXNUMX et d'évaluer le comportement RAID à grande échelle.
- Plate-forme: Dell PowerEdge R7715
- CPU: Processeur AMD EPYC 9655P 96 cœurs
- Percussion : 768 Go (12 x 64 Go) DDR5-5200 ECC
- Contrôleur de raid : 2 x PERC13 H975i
- Stockage: 32 disques durs Dell CD3.2P NVMe de 8 To
- Accélérateurs PCIe : 2 GPU NVIDIA H100
Stockage direct GPU NVIDIA Magnum IO : l'IA rencontre le stockage
Les pipelines d'IA modernes sont souvent liés aux E/S, et non aux calculs. Les lots de données, les intégrations et les points de contrôle doivent être transférés du stockage vers la mémoire GPU suffisamment rapidement pour occuper les accélérateurs. Le GDS Magnum IO de NVIDIA (via cuFile) court-circuite le chemin traditionnel « SSD → DRAM CPU → GPU » et permet aux données de passer directement du NVMe à la mémoire GPU en DMA. Cela supprime la surcharge du tampon de rebond du processeur, réduit la latence et rend le débit plus prévisible sous charge, ce qui se traduit par une meilleure utilisation du GPU, des temps d'époque plus courts et des cycles de sauvegarde/chargement des points de contrôle plus rapides.
