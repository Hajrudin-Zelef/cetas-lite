---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868-2
title: "fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Samsung"]
dates: []
keywords: ["intel", "benchmark", "ethernet", "gpu", "open source"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868.md
source_anchor: ""
source_lines: [43, 83]
sha256: 4467e35c0463242d0d7a30ed2375b38266aee7c2d430bc9290ebcfaadcf51c76
---

# fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868

Sur le panneau avant de l'unité, vous trouverez plusieurs indicateurs LED, notamment l'activité du disque dur, l'activité LAN1, l'état de l'alimentation et les informations système, qui fournissent des mises à jour d'état critiques en un coup d'œil. Le panneau comporte également un bouton d'alimentation et UID pour l'identification du système. La configuration par défaut comprend huit baies de lecteur 2.5 pouces remplaçables à chaud à l'avant qui prennent en charge les disques NVMe, SAS ou SATA, ce qui la rend idéale pour les solutions de stockage flexibles et à grande vitesse. De plus, une configuration optionnelle étend l'unité à 24 baies de lecteur SAS ou SATA 2.5 pouces remplaçables à chaud à l'avant. Supermicro n'a pas inclus d'options de connectivité USB ou d'affichage sur le panneau avant, conservant ainsi une conception simplifiée pour le stockage et la gestion.
Panneau arrière
À l'arrière de l'appareil, vous trouverez un port VGA pour la sortie vidéo, un port LAN BMC dédié pour la gestion hors bande et deux ports USB 3.2 Gen 1 pour la connectivité des périphériques. Deux blocs d'alimentation redondants remplaçables à chaud garantissent un fonctionnement ininterrompu.
Pour l'extension, il offre plusieurs emplacements PCIe 5.0, dont un emplacement PCIe 5.0 x16 pleine hauteur, 10.5 pouces de long et double largeur ; un emplacement PCIe 5.0 x8 pleine hauteur, pleine longueur ; un emplacement PCIe 5.0 x16 demi-hauteur, pleine longueur et double largeur ; et un emplacement PCIe 5.0 x16 AIOM compatible OCP 3.0. Ces emplacements prennent en charge de nombreux modules complémentaires, tels que des cartes réseau, des GPU et d'autres composants hautes performances.
Interne
La carte mère Super X14SBH-AP est dotée d'un seul socket LGA 7529 et de deux emplacements M.2 NVMe (M-key 2280/22110/25110) pour les supports de démarrage. De plus, la carte mère est équipée de plusieurs connecteurs, offrant la flexibilité d'intégrer des GPU optionnels pour améliorer les capacités du système pour les charges de travail exigeantes telles que l'IA, l'apprentissage automatique et la virtualisation.
Dans l'image ci-dessus, nous pouvons également voir deux cartes riser. La plus petite est destinée aux emplacements 3-4, tandis que la plus grande donne accès aux emplacements 5-8. Selon la configuration, ces cartes riser offrent une flexibilité supplémentaire pour étendre les capacités GPU et PCIe du serveur. Certains emplacements supplémentaires peuvent nécessiter des câbles spécifiques pour être pleinement utilisés, ce qui permet aux utilisateurs d'adapter la configuration du serveur à leurs besoins de performances.
La carte mère présente une conception refroidie par air capable de gérer un TDP impressionnant de 500 W, présentant une gestion thermique bien conçue pour les processeurs hautes performances.
Gestion intelligente de Supermicro
Ce châssis est équipé de l'interface BMC de gestion intelligente de Supermicro , pilotée par un contrôleur Aspeed AST2600 doté de 4 Go de mémoire vive DDR4. Il comprend un port Ethernet 1G pour la gestion à distance, permettant la surveillance du système, les mises à jour du firmware et la fonctionnalité KVM, même lorsque le système est hors tension. Cette configuration garantit une administration serveur efficace et fiable, idéale pour les environnements d'entreprise.
Le tableau de bord fournit un aperçu de toutes les informations essentielles du serveur, notamment des détails sur l'hôte, le système, la consommation d'énergie et l'accès à la console distante. Dans le panneau latéral, vous trouverez des commandes d'alimentation pour gérer l'unité, permettant des réglages et une surveillance faciles de l'état de l'alimentation du serveur.
La section Informations sur les composants fournit une vue détaillée de la configuration actuelle du serveur et de son état de santé, y compris le processeur, la mémoire, les blocs d'alimentation (PSU), la consommation d'énergie, les cartes réseau complémentaires, les capteurs, les systèmes de refroidissement et les GPU. Cela permet aux clients de surveiller les performances globales du système et de détecter tout problème potentiel.
En plus des informations sur les composants, la section Refroidissement vous permet d'ajuster le profil du ventilateur du serveur, offrant un contrôle sur les vitesses des ventilateurs pour optimiser les performances de refroidissement en fonction des besoins du système.
Dans la section Gestion du micrologiciel sous le volet Maintenance, deux fonctionnalités clés sont disponibles : la section Mise à jour, dans laquelle vous pouvez choisir le type de mise à jour (BMC ou BIOS) et les options de sauvegarde des anciennes images ou de préservation des configurations, et le volet Inventaire, qui affiche les versions actuelles du micrologiciel du serveur. Cela facilite le suivi et la vérification des versions utilisées.
Contrôle à distance : section avec plug-in HTML5 ou JAVAPour la gestion à distance, le BMC utilise un plug-in HTML5 ou Java, avec une option permettant de configurer le mode souris pour des performances optimales en fonction du système d'exploitation installé. C'est également ici que vous pouvez réinitialiser l'IKVM si nécessaire.
Serveur Hyper SuperServer SYS-212HA-TN de Supermicro Performances
Il s'agit de notre premier serveur Intel Granite Rapids. Nous l'avons donc comparé aux plates-formes Sierra Forest précédemment testées en laboratoire pour illustrer les différences de performances du processeur. Les plates-formes comparables proposent le processeur Sierra Forest Xeon 6780E avec 144 cœurs E, qui cible les charges de travail plus optimisées en termes de coûts. Il peut s'agir de domaines dans lesquels les clients cherchent à remplacer des plates-formes obsolètes sans augmentation significative des performances. En revanche, le processeur Granite Rapids Xeon 6980P avec 128 cœurs P cible les charges de travail plus performantes.
Serveur Hyper SuperServer SYS-212HA-TN de Supermicro
- CPU: 1x Intel Xeon 6980P (128 cœurs)
- RAM: 384GB DDR5
- SSD Disque SSD Micron 7450 NVMe pour centre de données
- Système opérateur: 2025 serveur
Configuration du Lenovo ThinkSystem SR630 V4
- CPU: 2 x Intel Xeon 6780E (144 cœurs)
- RAM: 512GB DDR5
- SSD Samsung MZWL6960HFJA-00AW7
- Système opérateur: 2025 serveur
Configuration du Supermicro Hyper 1U SYS-112H-TN
- CPU: 1x Intel Xeon 6780E (144 cœurs)
- RAM: 512GB DDR5
- SSD Disque SSD Micron 7450 NVMe pour centre de données
- Système opérateur: 2022 serveur
Mixeur OptiX
Tout d'abord, nous allons passer au test Blender, une application de modélisation 3D open source. Ce test a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Le Lenovo ThinkSystem SR630 V4 a excellé avec Blender 4.2, en fournissant 1,432 212 échantillons par minute dans la scène Monster. Le Supermicro SYS-1,134HA-TN a atteint 112 4.0 et le Supermicro SYS-781H-TNRT (exécutant Blender 915) a atteint 212. Lenovo a rapporté 759 échantillons dans la scène Junkshop, suivi du Supermicro SYS-112HA-TN avec 515 et du SYS-657H-TNRT avec 212. Dans la scène Classroom, Lenovo a mené avec 540 échantillons, le SYS-112HA-TN en a fourni 371 et le SYS-XNUMXH-TNRT a suivi avec XNUMX.
| Analyse comparative du CPU Blender | Supermicro Hyper 1U 112H-TN (1x Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2x Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| Monstre (Blender 4.0/4.2) | 781.42 | 1432.09 | 1134.77 | 
| Boutique de bric-à-brac (Blender 4.0/4.2) | 514.658 | 914.75 | 758.65 | 
| Salle de classe (Blender 4.0/4.2) | 370.52 | 656.68 | 540.47 | 
croque-y
