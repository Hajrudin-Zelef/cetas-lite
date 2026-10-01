---
id: collect-261001-huawei/huawei/fr-review-vergeio-high-performance-vmware-alternative-040958db-3
title: "fr-review-vergeio-high-performance-vmware-alternative-040958db"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "valuation"]
source: docs/RAG/collect-261001-huawei/fr-review-vergeio-high-performance-vmware-alternative-040958db.md
source_anchor: ""
source_lines: [34, 57]
sha256: cbac06f20032482edecdb90b75d2da6e4fa58664a3e3f9940dd42a7a71743c56
---

# fr-review-vergeio-high-performance-vmware-alternative-040958db

Pour mesurer les performances de la plateforme VergeIO, nous avons déployé 16 machines virtuelles (4 par nœud) pour mesurer les performances globales de ce cluster HCI. Ces machines virtuelles ont été utilisées pour orchestrer une charge de travail Vdbench exécutée de manière uniforme sur le cluster, toutes faisant rapport à une seule machine virtuelle. Ces sessions Vdbench ont également été configurées pour tester des données incompressibles afin de voir comment le cluster se comportait dans le pire des cas, car elles prennent en charge la réduction des données. En ce qui concerne l'empreinte des données, chaque machine virtuelle disposait d'un disque de données de 500 Go, soit 8 To au total sur l'ensemble du cluster.
Nous nous sommes concentrés sur les performances des quatre coins ainsi que sur les performances d'une base de données synthétique en utilisant les charges de travail suivantes :
- Lecture et écriture séquentielles de 2 Mo
- Lecture et écriture aléatoire 4K
- Charge de travail SQL
En ce qui concerne la bande passante séquentielle maximale, nous avons mesuré 4.7 Go/s en lecture sur notre niveau TLC et 4.2 Go/s sur le niveau QLC. En passant à la bande passante en écriture, le niveau TLC a mesuré 6.9 Go/s tandis que le niveau QLC a mesuré 5 Go/s.
| Charge de travail Vdbench | VergeIO Tier 1 Solidigm TLC | VergeIO Tier 2 Solidigm QLC | 
|---|---|---|
| 2 Mo de lecture séquentielle | 4.7 Go/s (27 ms) | 4.2 Go/s (30 ms) | 
| 2 Mo d'écriture séquentielle | 6.9 Go/s (17.6 ms) | 5.0 Go/s (21.5 ms) | 
| Lecture aléatoire 4K | 215 Mo/s (2.6 ms) | 243 Mo/s (8.2 ms) | 
| Écriture aléatoire 4K | 263 Mo/s (0.96 ms) | 200 Mo (0.85 ms) | 
| SQL | 533 Mo/s (0.89 ms) | 525 Mo/s (0.97 ms) | 
VergeIO VSAN affiche des performances impressionnantes sur les charges de travail des niveaux de stockage Solidigm TLC et QLC. Les opérations séquentielles affichent un excellent débit, le niveau TLC atteignant 6.9 Go/s pour les écritures et 4.7 Go/s pour les lectures. Les performances d'E/S aléatoires sont respectables, les deux niveaux atteignant plus de 200 Mo/s pour les opérations 4K. La plateforme excelle notamment dans les charges de travail de démarrage SQL et VDI, en maintenant des latences inférieures à la milliseconde et un débit élevé.
Ces résultats indiquent que les SSD Solidigm n'ont eu aucune difficulté à suivre le rythme des niveaux de stockage intégrés à notre plateforme, les contraintes de réseau et de plateforme étant les principaux facteurs limitatifs plutôt que les disques eux-mêmes. L'infrastructure ultra-convergée de VergeIO peut prendre en charge efficacement une large gamme d'applications d'entreprise, des transferts de fichiers volumineux aux opérations de base de données et aux environnements de bureau virtuel, le niveau TLC offrant généralement des performances supérieures pour les scénarios à forte intensité d'écriture.
Performances VDI
Le VDI étant une charge de travail courante déployée sur les plateformes VergeIO, nous avons voulu tester un bootstorm extrême pour voir comment le cluster se comporte avec 1000 2 machines virtuelles tournant simultanément. Chaque machine virtuelle avait 2 processeurs, 10 Go de RAM et un disque de 22.04 Go. La machine virtuelle avait une installation standard d'Ubuntu XNUMX (pas minimale) pour représenter une image réelle. Une fois complètement démarré, un script est appelé via systemd qui utilise curl pour envoyer son mac et son horodatage via HTTP à un collecteur distant.
Nos tests ont révélé que les disques TLC et QLC offraient des performances finales très similaires. En examinant les données de stockage back-end, les disques SSD TLC avaient un avantage en termes d'IOPS totales, même si pour cette plateforme, les processeurs sont devenus le goulot d'étranglement avant le stockage. Les 1000 71 machines virtuelles ont pu démarrer en XNUMX secondes environ. Ces résultats mettent en évidence les avantages de combiner différents disques SSD dans la plateforme VergeIO. Les clients peuvent facilement utiliser le stockage QLC pour les tâches VDI, ce qui, dans ce cas, permet une densité et une rentabilité fantastiques.
Nous avons également testé la haute disponibilité. Nous avons mesuré 138 secondes pour que la VM devienne disponible après une perte totale d'un nœud, ce qui correspond à peu près à ce que VergeIO affirme.
Conclusion
VergeIO se distingue comme une excellente alternative à VMWare. Son wiki complet et son processus de migration convivial en font une option intéressante pour les organisations qui envisagent un changement. Le modèle de licence par nœud simplifié, qui est en moyenne 50 % moins cher que VMWare, facilite l'évaluation et l'adoption par les décideurs.
La gestion convergée en fait également l'une des plateformes les plus simples à gérer. Dans le même temps, des fonctionnalités telles que le catalogue et l'intégration native de Wireguard améliorent l'ensemble, les performances robustes de la plateforme et l'attention méticuleuse portée à la reprise après sinistre, ce qui est vraiment impressionnant.
La nature raffinée de la plateforme VergeIO la place au-dessus des alternatives comme Proxmox. Nous pouvons facilement recommander VergeIO à toute personne à la recherche d'une solution d'hyperviseur. Elle offre un environnement robuste et riche en fonctionnalités adapté aux organisations de toutes tailles. VergeIO combine avec succès simplicité d'utilisation, performances élevées et capacités de niveau entreprise, ce qui en fait un choix de premier ordre dans le paysage de la virtualisation.
Ce rapport est sponsorisé par VergeIO. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits considérés.
