---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036-1
title: "fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["amd", "benchmark", "gpu", "open source"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036.md
source_anchor: ""
source_lines: [1, 63]
sha256: 763e49ff71a24c76ebea39519a14a6019a503ee90cc1e4cc6e6379e77427e06a
---

# fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036

Supermicro développe des solutions de stockage impressionnantes et innovantes, et l'ASG-1115S-NE316R ne fait pas exception. Il s'agit d'un serveur EPYC série 9004 à socket unique avec jusqu'à 16 SSD E3.S Gen5 remplaçables à chaud dans un châssis 1U.
La plupart des fournisseurs de flash SSD d'entreprise investissent dans les facteurs de forme EDSFF pour les SSD Gen5. Dans l'entreprise, cela signifie généralement que E3.S est le facteur de forme de choix, là où les hyperscalers ont tendance à se tourner vers E1.SU2/U.3 qui a perdu de son élan, car il existe peu d'options Gen5. En tant que telle, cette plate-forme de Supermicro permet une configuration à stockage dense, prenant en charge jusqu'à 16 SSD E3.S, sans compromettre les options d'extension PCIe à l'arrière du châssis.
Entièrement équipé, ce châssis prend toujours en charge 64 voies PCIe pour les E/S arrière sur deux emplacements OCP Gen5 x16 et deux emplacements PCIe Gen5 x16 FH. Supermicro comprend deux emplacements M.2 compatibles avec les SSD de 80 mm et 110 mm pour la capacité de démarrage intégrée. Ce châssis vous offre la plate-forme idéale pour créer un serveur à stockage dense autour du socket AMD Epyc SP5.
Supermicro Storage A+ ASG-1115S-NE316R Conception et construction
L'ASG-1115S-NE316R est un serveur à socket unique avec jusqu'à 24 DIMM DDR5 et 2 M.2 sur la carte avec prise en charge jusqu'à 6 To de mémoire et jusqu'à 96 cœurs. Le refroidisseur de processeur ramifié situé près du centre de la carte est visuellement frappant. Nous avons vu cela récemment sur d'autres serveurs. Cette conception donne plus de surface au refroidisseur de processeur, lui permettant de compenser plus de chaleur dans un si petit boîtier, ce qui est de plus en plus difficile dans les serveurs 1U refroidis par air.
Le reste du système de refroidissement se compose principalement de 8 ensembles de ventilateurs à double empilement qui aspirent l'air autour des baies de disque, le forcent sur le dissipateur thermique et la RAM, puis dans le reste du châssis, l'évacuant par l'arrière.
Côté mémoire, l'ASG-1115S-NE316R dispose de 24 emplacements pour DDR5 ECC. La configuration de cet examen était livrée avec 384 Go de DDR5, composés de 12 clés Samsung DDR32 ECC de 5 Go fonctionnant à 4,800 XNUMX Mbps.
Une autre touche intéressante sur ce serveur est l'utilisation de caddies sans outils pour les disques E3.S, qui permettent de préparer ou de remplacer facilement les disques. C'est formidable de voir des méthodes de rétention de disque sans outils, quelle que soit l'application.
À l'arrière du serveur, nous pouvons voir les quatre emplacements d'extension différents, qui se composent de deux emplacements PCIE 5.0 16x FH et de deux emplacements PCIE 5.0 16x AIOM. Actuellement, dans les emplacements FH PCIe 5.0, il y a une carte réseau Mellanox ConnectX-6 DX 200G dans chaque emplacement. L'autre IO est assez standard, avec un seul RJ45 pour l'IPMI, 2 ports USB 3.0 et un port VGA.
Ce qui est intéressant à propos des E/S arrière, c'est que tous les ports se trouvent sur une carte fille reliée à la carte mère avec un seul connecteur. Cela facilite la conception de la carte mère en éliminant le besoin d'une section dépassant entre les emplacements AIOM, ce qui constituerait un gaspillage de matériau dans le processus de fabrication.
ASG-1115S-NE316R Spécifications
|  | ASG-1115S-NE316R | 
|---|---|
| Processeur | Prise unique SP5 Série Epyc 9004 Jusqu'à 96C/192T Prend en charge les processeurs TDP 200 W-300 W (refroidis par air). | 
| RAM | 24 emplacements DIMM jusqu'à 6 To ECC DDR5 RDIMM/LRDIMM à 4800 XNUMX MT/s | 
| Baies de stockage/disque |  | 
| Slots d'extension |  | 
| Entrée / Sortie |  | 
| Panneau avant |  | 
| Alimentations | 2 alimentations en titane de 1600 XNUMX W  | 
| Carte mère | Super H13SSF | 
| Châssis | 1U Rackmount CSE-126E32-R1K62P | 
| Direction |  | 
| Sécurité | matériel:  Logiciel:  | 
| Environnement d'exploitation |  | 
La fiche technique complète est disponible ici.
Gestion du stockage Supermicro A+ ASG-1115S-NE316R
L'ASG-1115S-NE316R utilise l'interface IPMI pour la gestion hors bande (OOBM), qui est assez uniforme sur les systèmes Supermicro modernes. L'IPMI permet la gestion à distance tout en affichant également l'état des composants individuels.
L'IPMI permet de diagnostiquer rapidement les pannes matérielles et simplifie la configuration et la gestion du matériel. L'utilisation de la console distante supprime le besoin de transporter un moniteur et des périphériques sur le serveur et vous permet de le faire à distance. La console distante dispose également d'une option permettant de monter un support virtuel pour l'installation du système d'exploitation, l'installation de logiciels et les mises à jour.
Performances
Notre unité d'examen Supermicro Storage A+ ASG-1115S-NE316R est configurée avec les éléments suivants :
- Processeur : EPYC 84 à 9634 cœurs
- RAM : 384 Go DDR5 (12 x DIMM de 32 Go)
- SSD : 1 To Samsung PM9A3
- 2x Mellanox ConnectX-6 DX 100G
- Windows Server 2022
Afin d'établir une comparaison, nous ajouterons les données de benchmark des serveurs Genoa-X (96 cœurs) et Bergamo (128 cœurs) issues de notre test des processeurs AMD Bergamo/Genoa-X . Ces données visent à illustrer l'échelle de performances au sein de la gamme de processeurs AMD EPYC ; il s'agit de la première plateforme serveur testée en laboratoire avec l'EPYC 9634.
TYAN SX TS70AB8056
- EPYC 96X à 9684 cœurs
- 512GB DDR5
TYAN SX TS70AB8056
- EPYC 128 à 9754 cœurs
- 512GB DDR5
Mélangeur OptiX 3.6
Le premier est Blender OptiX, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire CLI Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur. Le test de la version 3.6 existe depuis un certain temps et nous avons ici les résultats de toutes les machines de comparaison. Comme prévu, le 84C EPYC du Supermicro a plutôt bien résisté et a fonctionné juste derrière le 96X à 9684 cœurs, qui a fonctionné juste derrière le 128 à 9754 cœurs.
| Processeur Blender 3.6 | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | TYAN SX TS70AB8056 (96C EPYC 9684X, 512 Go DDR5) | TYAN SX TS70AB8056 (128C EPYC 9754, 512 Go DDR5) | 
|---|---|---|---|
| Monster | 702.590 | 879.580 | 1031.495 | 
| Brocanteur | 483.502 | 605.446 | 704.168 | 
| Salle de classe | 341.717 | 421.318 | 506.666 | 
Mélangeur OptiX 4.0
Depuis que nous avons testé les puces Bergamo et Genoa-X il y a quelque temps, nous n'avons pas de résultats pour Blender 4.0, sorti assez récemment. Pour cette raison, nous n'avons que des résultats pour le Supermicro avec le 84C EPYC 9634 dans le test CPU 4.0. Pour cela, EPYC 9634 a marqué 673.21 points sur Monster, 475.17 sur Junkshop et 342.06 sur Classroom.
| Processeur Blender 4.0 | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
|---|---|
| Monster | 673.21 | 
| Brocanteur | 475.17 | 
| Salle de classe | 342.06 | 
Test de vitesse Blackmagic RAW
Nous avons effectué le test de vitesse RAW de Blackmagic, qui teste la lecture vidéo. Il s’agit plutôt d’un test hybride incluant les performances du CPU et du GPU pour le décodage RAW réel. Comme ce serveur n'a pas de GPU, nous n'avons aucun résultat CUDA, mais il atteint 131 FPS avec le CPU.
| Test de vitesse Blackmagic RAW (Plus c'est haut, mieux c'est) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
|---|---|
| CPU 8K | FPS 131 | 
Test de vitesse du disque Blackmagic
