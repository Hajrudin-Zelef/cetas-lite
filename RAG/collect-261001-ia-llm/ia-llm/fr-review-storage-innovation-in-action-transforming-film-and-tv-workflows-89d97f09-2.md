---
id: collect-261001-ia-llm/ia-llm/fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09-2
title: "fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09"
domain: ia-llm
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["exploit", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09.md
source_anchor: ""
source_lines: [16, 44]
sha256: 696e947d7c8303981e21dc4eeeaafe229b534cf3d1c0cca2ec92f356229b2104
---

# fr-review-storage-innovation-in-action-transforming-film-and-tv-workflows-89d97f09

Nous avons exploité deux stations de travail HP Z4 Rack G5 pour agir comme chargeurs dans notre environnement de test. Chaque système disposait de processeurs Intel Xeon W16-5X à 2465 cœurs et de 128 Go de RAM DDR5. Ils proposent également des GPU NVIDIA RTX 5000 Ada et des cartes réseau NVIDIA ConnectX-6 Dx 100GbE pour en faire des centrales de montage vidéo. Les systèmes ont été configurés avec Windows Server 2022. Il convient de noter que HP ne prend pas directement en charge Windows Server et les cartes réseau 100GbE, mais aucun n'a posé de problème lors de ces tests.
La structure réseau haut débit les connecte au CheetahRAID Raptor via des liaisons 100 GbE, offrant à chaque client une liaison de 10 Go/s vers le stockage partagé. Avec 12 SSD Solidigm P61.44 QLC de 5336 To dans un pool de stockage Graid, nous pourrions partager un énorme volume de 613 To avec chaque hôte.
Pour mesurer les avantages du logiciel Fusion de Tuxera, nous avons comparé la bande passante et la latence de lecture et d'écriture séquentielles de 1 Mo de chaque client interagissant simultanément avec l'hôte de stockage CheetahRAID. Nous avons utilisé quatre tâches FIO sur chaque client, chacune mesurant les performances d'un fichier de 50 Go. Au total, cela équivaut à une empreinte de 400 Go. Notre référence a été mesurée à l'aide de Samba version 4.15.13 par rapport à Fusion sur TCP et à nouveau avec Fusion tirant parti du déchargement RDMA.
La bande passante de base de lecture Samba sur chaque station de travail HP Z4 Rack G5 mesurait 4 Go/s et 3.8 Go/s, pour un total de 7.8 Go/s. Dans une charge de travail d'écriture séquentielle, nous avons mesuré 2.4 Go/s et 2.6 Go/s pour un total de 5 Go/s. La latence de lecture moyenne mesurait 17.23 ms, avec une latence d'écriture moyenne de 26.6 ms.
Le passage au logiciel Fusion de Tuxera a eu un impact considérable, sans aucun changement nécessaire du côté client. En regardant le protocole TCP, nous avons mesuré 11.1 Go/s et 11.1 Go/s pour chaque client, ce qui nous donne 22.2 Go/s de bande passante en lecture. C'était la limite de la connexion 100 GbE à chaque hôte. Pour la bande passante en écriture, nous avons mesuré 5.7 Go/s et 5.8 Go/s, ce qui nous donne un total de 11.5 Go/s. La latence de lecture était en moyenne de 6 ms tandis que la latence d'écriture était de 11.65 ms.
En plus de TCP, Fusion File Share de Tuxera prend également en charge RDMA. Nous avons mesuré le protocole Fusion RDMA, qui nous a donné une bande passante en lecture mesurant 11.6 Go/s et 11.6 Go/s pour chaque hôte, soit 23.2 Go/s. La bande passante en écriture était de 5 Go/s et 5.2 Go/s, soit un total de 10.2 Go/s. La latence de lecture dans cette configuration était de 5.8 ms, tandis que la latence d'écriture était de 13.2 ms.
La comparaison de Samba à Fusion a montré d'énormes gains pour les clients Windows. La bande passante de lecture a été multipliée par près de 3, avec une latence de seulement 33 % de celle proposée par Samba. La bande passante en écriture a également été multipliée par 2.3, avec une latence de seulement 44 % de celle mesurée avec Samba.
| Passerelle | Métrique | Client1 | Client2 | Total | 
|---|---|---|---|---|
| Samba | Lire la bande passante | 4GB / s | 3.8GB / s | 7.8GB / s | 
|  | Lire la latence | 16.77ms | 17.68ms | 17.23ms | 
|  | Bande passante d'écriture | 2.4GB / s | 2.6GB / s | 5GB / s | 
|  | Latence d'écriture | 27.5ms | 25.7ms | 26.6ms | 
| Fusion TCP | Lire la bande passante | 11.1GB / s | 11.1GB / s | 22.2GB / s | 
|  | Lire la latence | 6ms | 6ms | 6ms | 
|  | Bande passante d'écriture | 5.7GB / s | 5.8GB / s | 11.5GB / s | 
|  | Latence d'écriture | 11.8ms | 11.5ms | 11.65ms | 
| FusionRDMA | Lire la bande passante | 11.6GB / s | 11.6GB / s | 23.2GB / s | 
|  | Lire la latence | 5.8ms | 5.8ms | 5.8ms | 
|  | Bande passante d'écriture | 5GB / s | 5.2GB / s | 10.2GB / s | 
|  | Latence d'écriture | 13.4ms | 13ms | 13.2ms | 
D’énormes SSD suivent le rythme du S&E
Les disques SSD massifs comme le P61.44 de 5336 To de Solidigm apportent des avantages substantiels aux secteurs des médias et du divertissement, où la vitesse et la fiabilité sont cruciales pour gérer des fichiers volumineux tels que des vidéos haute définition, des graphiques complexes et des pistes audio étendues. Contrairement aux disques durs traditionnels (Hard Disk Drives), les SSD offrent des temps d'accès aux données plus rapides, des vitesses de lecture/écriture supérieures et un plus grand nombre d'opérations d'E/S par seconde. Cet avantage en termes de performances permet des flux de travail d'édition, de rendu et de traitement plus efficaces, réduisant considérablement le temps nécessaire au chargement et à l'utilisation de fichiers multimédias volumineux. L'absence de pièces mobiles dans les disques SSD améliore leur fiabilité et leur durabilité, les rendant moins sensibles aux pannes mécaniques et à la perte de données (des problèmes critiques lorsqu'il s'agit de contenu multimédia de valeur).
L'adoption de disques SSD denses dans les environnements de production multimédia et de post-production rationalise les flux de travail, permettant l'édition, l'étalonnage des couleurs et le traitement des effets en temps réel sans compromettre la qualité ou l'efficacité. Ces disques peuvent gérer de manière transparente plusieurs flux de vidéo 4K, éliminant ainsi le besoin de fichiers proxy ou d'espaces réservés basse résolution. De plus, le format compact des disques SSD, combiné à leurs capacités de stockage élevées, simplifie la gestion des données en permettant de stocker des projets entiers sur un seul disque ou sur une matrice minimale, facilitant ainsi l'accès et la gestion de gros volumes de données.
Au-delà des avantages en termes de performances et de capacité, les disques SSD massifs contribuent à un environnement de production plus propice grâce à leur fonctionnement silencieux, leur faible production de chaleur et leur efficacité énergétique. Ces fonctionnalités sont particulièrement utiles dans les suites de montage densément remplies ou mobiles, où la réduction du bruit et le refroidissement sont des préoccupations constantes. L'efficacité énergétique des disques SSD réduit non seulement les coûts opérationnels, mais permet également de créer un espace de travail plus frais et plus silencieux, améliorant ainsi la productivité globale et le confort des professionnels des médias. Essentiellement, les disques SSD massifs transforment le paysage des médias et du divertissement, permettant des processus de production plus rapides, plus fiables et plus efficaces, capables de suivre le rythme de la demande croissante de contenu numérique de haute qualité.
Bridge Digital est un grand fan
Lors de nos propres tests avec la plate-forme CheetahRAID, nous avons constaté des avantages impressionnants pour les charges de travail de S&E. Mais nous voulions un autre avis de l’industrie, pour voir si nos conclusions étaient cohérentes avec celles de ceux qui sont à fond dans la gestion des données sur le plateau. Nous avons contacté notre ami Richie Murray, fondateur et président de Bridge Digital.
Bridge Digital est une entreprise possédant une expertise dans les flux de travail vidéo numériques et dans les technologies permettant de mieux les faire fonctionner.
Ils aident les créateurs et les propriétaires de contenu numérique à créer une infrastructure pour créer, gérer, distribuer et monétiser efficacement leurs ressources vidéo. Les solutions de Bridge Digital couvrent l'ensemble du flux de travail des médias numériques, de l'ingestion à la livraison finale.
