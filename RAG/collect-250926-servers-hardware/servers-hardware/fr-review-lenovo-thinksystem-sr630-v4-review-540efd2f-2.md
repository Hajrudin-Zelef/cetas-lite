---
id: collect-250926-servers-hardware/servers-hardware/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-2
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Samsung"]
dates: []
keywords: ["arr", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [46, 89]
sha256: 03a97672652980aacba5129573eeb081679b3731f7738b3e910cc20c7a1d034a
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

| Garantie | Trois ans ou un an (selon le modèle) avec des mises à niveau de service facultatives pour des temps de réponse plus rapides et une couverture étendue. | 
| Dimensions | Largeur : 440 mm (17.3 po), Hauteur : 43 mm (1.7 po), Profondeur : 788 mm (31 po). | 
| Poids | Poids maximal : 20.2 kg (44.5 lb) | 
Conception et fabrication du Lenovo ThinkSystem SR630 V4
Le Lenovo ThinkSystem SR630 V4 conserve le format compact 1U qui est la norme pour de nombreux serveurs rack d'entreprise. Sa conception se concentre sur une combinaison de fonctionnalité, d'accessibilité et de flexibilité. Nous avons apprécié sa disposition simple mais efficace, qui maximise la circulation de l'air, la modularité et la facilité d'utilisation pour les administrateurs informatiques.
Passons aux détails.
Panneau avant
Le panneau avant prend en charge jusqu'à 10 baies de disques durs remplaçables à chaud de 2.5 pouces et offre une flexibilité pour diverses configurations de stockage, notamment les disques SAS, SATA, NVMe ou AnyBay. Cela permet aux entreprises de personnaliser le stockage en fonction de leurs charges de travail, qu'elles privilégient la vitesse, la capacité ou la rentabilité.
Le panneau avant peut également être configuré avec un port vidéo Mini DisplayPort en option, qui peut être utilisé pour une surveillance et des diagnostics locaux rapides sans avoir à accéder à l'arrière du rack. Jusqu'à deux ports USB 3.0 en option sont également disponibles ; l'un est explicitement désigné pour la connexion au contrôleur Lenovo XClarity (XCC). Cette connectivité simplifie les tâches de gestion, telles que le téléchargement de mises à jour du micrologiciel ou l'exécution de diagnostics directement à partir d'un périphérique USB.
Lenovo a intégré un port de diagnostic externe sur notre système, ce qui peut s'avérer très utile pour le personnel informatique sur site qui doit résoudre les problèmes matériels (tels que l'état du système et les pannes). Cela peut accélérer la résolution des problèmes et minimiser les temps d'arrêt. Il dispose également d'une étiquette d'information amovible pour accéder rapidement aux détails essentiels du système, tels que les numéros de série, les configurations et les informations réseau.
Les indicateurs et commandes du panneau de commande avant fournissent les informations habituelles en un coup d'œil sur l'état et l'activité du système, y compris les boutons d'alimentation et de réinitialisation et les indicateurs LED pour l'état et la santé du lecteur.
Panneau arrière
Le panneau arrière comprend des blocs d'alimentation remplaçables à chaud (de 800 W à 2000 3.0 W) situés de chaque côté du système, qui assurent la redondance et peuvent être remplacés sans éteindre le système. Les emplacements Dual OCP 5 prennent en charge PCIe Gen 16 x200 pour une mise en réseau avancée, autorisant des adaptateurs à large bande passante tels que des cartes 3.0 GbE à double port. Le panneau comprend également un port vidéo, deux ports USB XNUMX et un port de gestion dédié XClarity Controller (XCC) pour la gestion locale et à distance du système. Les indicateurs LED offrent des mises à jour visuelles rapides de l'état du système.
Le panneau arrière comprend un mélange d'emplacements PCIe à profil bas et pleine hauteur pour l'extension, permettant aux utilisateurs d'ajouter des GPU, des contrôleurs de stockage ou d'autres adaptateurs. Les options de stockage incluent des disques durs remplaçables à chaud de 2.5 pouces et des disques durs M.2 remplaçables à chaud, offrant des configurations flexibles pour les périphériques de démarrage ou le stockage supplémentaire. Le panneau arrière est également disponible en quatre configurations refroidies par air et deux configurations refroidies par eau.
Interne
Lorsque vous ouvrez le Lenovo ThinkSystem SR630 V4, vous verrez les deux processeurs entourés de leurs emplacements DIMM respectifs. Ce châssis offre 16 DIMM par processeur, soit 32 au total, ce qui permet d'installer jusqu'à 2 To de RAM.
De face, vous remarquerez jusqu'à huit ventilateurs remplaçables à chaud alignés, canalisant le flux d'air sur les composants critiques et gardant tout au frais sous pression. Notre système est équipé de SSD NVMe directement connectés, qui sont câblés directement sur la carte mère.
Ce NVMe à connexion directe offre les performances NVMe les plus élevées disponibles, bien que, pour RAID, les utilisateurs doivent choisir entre des options logicielles ou matérielles comme Graid.
À l'arrière, les emplacements PCIe sont prêts à accueillir des GPU ou d'autres cartes d'extension, tandis que les emplacements OCP ajoutent une couche supplémentaire de polyvalence pour les besoins de réseau spécialisés. Les blocs d'alimentation remplaçables à chaud sont faciles d'accès et de remplacement, ce qui minimise les temps d'arrêt pendant la maintenance.
Contrôleur XClarity 3
Le Lenovo ThinkSystem SR630 V4 est équipé du contrôleur XClarity Controller 3 (XCC3) pour la gestion à distance et du cycle de vie. Il offre des capacités de gestion hors bande via un port LAN dédié, permettant aux utilisateurs de configurer et de déployer de nouveaux matériels, d'interagir avec le système si les réseaux principaux sont en panne, d'effectuer des activités de gestion du micrologiciel et d'effectuer d'innombrables autres tâches.
Les utilisateurs peuvent obtenir un aperçu rapide de tous les principaux composants et avertissements à partir de l'écran d'accueil principal. L'état du système, les événements actifs et l'alimentation sont les domaines clés ici.
Vous pouvez voir comment la plateforme gère les mises à jour en explorant plusieurs domaines, tels que la mise à jour du micrologiciel. Vous importez un package de micrologiciel dans le stockage local à l'intérieur du contrôleur XClarity et cliquez sur Mettre à jour le système pour démarrer le processus de mise à jour de la charge utile du micrologiciel téléchargée.
La télécommande est une autre fonction courante que Lenovo XClarity gère bien via une interface Web HTML5. Cela permet à une large gamme de systèmes clients de gérer la plateforme, y compris les plateformes mobiles. Bien qu'un iPad ne soit pas le mécanisme de support principal, vous devez parfois utiliser ce qui se trouve à proximité.
Performances du Lenovo ThinkSystem SR630 V4
Cette section examine les résultats des tests de performance de y-cruncher, Cinebench, Blackmagic, 7-Zip et Geekbench. Nous avons comparé le Lenovo ThinkSystem SR630 V4 à double processeur avec le Supermicro Hyper 1U SYS-112H-TN à processeur unique récemment testé. Les deux systèmes sont équipés du processeur Intel Xeon 6780E, ce qui nous permet de voir comment le 6780E évolue à partir de configurations à un ou deux processeurs.
En plus du Supermicro, nous avons ajouté un ancien serveur Intel Ice Lake, fourni lors de la sortie des premiers processeurs Ice Lake Xeon 8380. Cela montre comment les modèles E-core se positionnent comme une mise à niveau rentable pour les plates-formes héritées qui privilégient l'efficacité à la puissance de traitement brute. Cette comparaison entre la plate-forme biprocesseur SR630 V4 et Intel Ice Lake et le Supermicro Hyper 1U monoprocesseur illustre la différence de performances.
Voici les configurations pour chaque système.
Configuration du Lenovo ThinkSystem SR630 V4
- CPU: 2 x Intel Xeon 6780E (144 cœurs)
- RAM: 512GB DDR5
- SSD Samsung MZWL6960HFJA-00AW7
- Système opérateur: 2025 serveur
Configuration du Supermicro Hyper 1U SYS-112H-TN
- CPU: Intel Xeon 6780E (144 cœurs)
- RAM: 512GB DDR5
- SSD Disque SSD Micron 7450 NVMe pour centre de données
- Système opérateur: 2022 serveur
Serveur Intel Ice Lake
- CPU: 2 x Intel Xeon 8380 (80 cœurs)
- RAM: 512GB DDR5
- SSD
- Système opérateur: 2025 serveur
Mixeur OptiX
