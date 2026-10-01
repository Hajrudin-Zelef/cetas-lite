---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf-2
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia", "Samsung"]
dates: []
keywords: ["amd", "benchmark", "gpu", "nvidia", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [43, 121]
sha256: 0c4374e1b3bd56f4da83bf7164083c4a17bd0d87d6cf47737ce718140d262494
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

| Fonctionnalité |  | 
| Surveillance de la santé PC |  | 
| FAN |  | 
| Température |  | 
| Châssis |  | 
| Facteur de forme | 1U | 
| Modèle | CSE-116BTS-R000WNP | 
| Dimensions et poids |  | 
| Hauteur | 1.7 mm (43 po) | 
| Largeur | 17.2 mm (437 po) | 
| Profondeur | 23.5 mm (597 po) | 
| Forfait | 7.75 "(H) x 23.5" (L) x 31.5 "(P) | 
| Poids |  | 
| Couleur disponible | Noir | 
| Panneau avant |  | 
| DEL |  | 
| Boutons |  | 
| Slots d'extension |  | 
| Configuration PCI-Express (PCIe) |  | 
| M.2 | 2 emplacements M.2 PCIe 3.0 x4 NVMe (clé M) | 
| Baies de lecteur/stockage |  | 
| Configuration des baies de lecteur |  | 
| M.2 | 2 emplacements M.2 PCIe 3.0 x4 NVMe (clé M 2280/22110) Afficher les options M.2 | 
| Refroidissement du système |  | 
| Ventilateurs | 6 ventilateurs robustes contrarotatifs de 40 x 40 x 56 mm. | 
| Source d'alimentation |  | 
|  |  | 
Performances du Supermicro 1115SV-WTNRT
Notre unité d'examen Supermicro Storage A+ 1115SV-WTNRT est configurée avec les éléments suivants :
- Processeur : AMD EPYC 8534P 64 cœurs
- RAM: 194GB DDR5
- GPU : NVIDIA A2 (15,360 6 Mo GDDRXNUMX)
- SSD : SSD pour centre de données Ultrastar SN655 NVMe (15.36 To)
- 2x Mellanox ConnectX-6 DX 100G
- 10 64-bit Windows
Nous avons récemment testé la solution RAID Graid SupremeRAID SR-1001 avec le SSD Supermicro AS-1115SV-WTNRT. Pour découvrir leurs performances respectives, consultez notre test complet.
Lors de cette analyse des performances, nous le comparerons cependant au Supermicro Storage A+ ASG-1115S-NE316R. Il est important de noter qu’il ne s’agira pas d’une évaluation directe. Nous visons plutôt à fournir une perspective sur l’échelle des deux processeurs AMD.
L'unité d'examen A+ ASG-1115S-NE316R est configurée avec les éléments suivants :
- Processeur : EPYC 84 à 9634 cœurs
- RAM : 384 Go DDR5 (12 x DIMM de 32 Go)
- SSD : 1 To Samsung PM9A3
- 2x Mellanox ConnectX-6 DX 100G
- Windows Server 2022
Mélangeur OptiX 4.0
Le premier est Blender OptiX, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire CLI Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Le système Supermicro a présenté des performances variables dans les tâches de rendu 3D sur différentes scènes lorsqu'il a été testé avec et sans accélération GPU. Dans tous les scénarios (Monster, Junkshop et Classroom), la configuration avec processeur uniquement a fourni des échantillons par minute plus élevés que la configuration assistée par GPU.
Plus précisément, la scène Monster a atteint 492.88 échantillons/min sans GPU, surpassant le score de 427.25 échantillons/min avec GPU. De même, dans la scène Junkshop, le test CPU uniquement a atteint 349.35 échantillons/min, dépassant largement les 266.65 échantillons/min du GPU. La scène Classroom a également enregistré de meilleures performances avec le processeur seul, avec un score de 255.98 échantillons/min contre 238.33 échantillons/min pour le GPU.
| Mixeur 4.0 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5, GPU) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5, processeur) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| Monster | 427.25 | 492.88 | 673.21 | 
| Brocanteur | 266.65 | 349.35 | 475.17 | 
| Salle de classe | 238.33 | 255.98 | 342.06 | 
Mélangeur OptiX 4.1
Blender OptiX 4.1 apporte de nouvelles fonctionnalités, telles que le débruitage accéléré par GPU, la rationalisation du processus de rendu et la réduction du temps nécessaire aux tâches de débruitage. Malgré ces progrès, les améliorations globales des performances dans les scores de référence par rapport à la version 4.0 sont minimes, indiquant seulement de légères améliorations en termes d'efficacité (comme vous le remarquerez dans les résultats ci-dessous).
Les résultats montrent que la configuration CPU uniquement a systématiquement surpassé la configuration améliorée GPU dans tous les tests, comme dans la version 4.0. Plus précisément, le mode CPU uniquement a atteint 493.167 échantillons/min dans Monster, 356.36 échantillons/min dans Junkshop et 250.87 échantillons/min dans Classroom, par rapport aux scores assistés par GPU de 428.90, 269.39 et 238.58 échantillons/min respectivement.
| Mixeur 4.0 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5, GPU) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5, processeur) | 
| Monster | 428.90 | 493.167 | 
| Brocanteur | 269.39 | 356.36 | 
| Salle de classe | 238.58 | 250.87 | 
Test de vitesse Blackmagic RAW
Nous avons effectué le test de vitesse RAW de Blackmagic pour étendre la lecture vidéo. Il s’agit plutôt d’un test hybride incluant les performances du CPU et du GPU pour le décodage RAW réel. Ici, nous avons uniquement testé le processeur, qui a pu atteindre 117 FPS.
| Test de vitesse Blackmagic RAW (Plus c'est haut, mieux c'est) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| CPU 8K | FPS 117 | FPS 131 | 
Test de vitesse du disque Blackmagic
Le Blackmagic Disk Speed Test est un autre test pour lequel nous disposons uniquement de résultats pour le Supermicro. Ce test exécute un exemple de fichier de 5 Go pour les vitesses de lecture et d'écriture. Ce test a montré des vitesses de lecture de 3.57 Go/s et de près de 2.54 Go/s en lecture sur les SSD Ultrastar SN655 NVMe Data Center installés par Supermicro.
| Test de vitesse du disque Blackmagic (Plus c'est haut, mieux c'est) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| Écrire | 2,536MB / s | 1,415.1 Mo / s | 
| Lire | 3,568MB / s | 3,031.0 Mo / s | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs. Voici les résultats pour toutes les puces EPYC.
Ici, le Supermicro 1115SV-WTNRT a obtenu un score de 63,332 1,093 points au test multicœur, démontrant la robuste capacité du processeur à gérer plusieurs threads simultanément. D'autre part, le score au test monocœur était de 57.92 XNUMX points, ce qui reflète son efficacité dans les tâches nécessitant des performances monothread. Le ratio MP, qui compare les performances multicœurs aux performances monocœur, s'élève à XNUMX, indiquant une architecture bien équilibrée pour le traitement parallèle et l'efficacité de chaque cœur.
| Cinebench R23 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| Processeur (multicœur) (points) | 63,332 | 81,148 | 
| Processeur (monocœur) (points) | 1,093 | 1,309 | 
| Rapport PM | 57.92 | 61.99x | 
Cinebench 2024
Cinebench 2024 de Maxon est une référence de rendu CPU et GPU qui utilise tous les cœurs et threads du processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Nous n'avons pas ces chiffres car la configuration ASG-1115S-NE316R ne dispose pas de GPU. Des scores plus élevés sont meilleurs.
Le 1115SV-WTNRT a obtenu 3,928 69 points au test multicœur, soulignant ses solides performances en termes de puissance de traitement pour le multitâche et les tâches de calcul exigeantes. Les performances monocœur ont été mesurées à 3,634 points, tandis que le benchmark GPU a atteint un score de 57.27 XNUMX points, démontrant sa maîtrise du traitement graphique. Le ratio MP a été calculé à XNUMX, démontrant un équilibre cohérent entre les capacités de traitement multithread et monothread.
| Cinebench 2024 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| Processeur (multicœur) (Points) | 3,928 | 4,913 | 
