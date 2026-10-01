---
id: collect-261001-general-networking/general-networking/fr-review-qnap-thunderbolt-nas-das-review-creative-professional-take-3f127875-1
title: "fr-review-qnap-thunderbolt-nas-das-review-creative-professional-take-3f127875"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-qnap-thunderbolt-nas-das-review-creative-professional-take-3f127875.md
source_anchor: ""
source_lines: [1, 29]
sha256: cfa03137d89c97387294bd3453d774d5ecf664b4832820ee825e9a5ed8a1451c
---

# fr-review-qnap-thunderbolt-nas-das-review-creative-professional-take-3f127875

Le QNAP TVS-872XT est un NAS Thunderbolt à 8 baies conçu principalement pour les applications vidéo et multimédia. Il fait partie de la gamme de NAS x72XT de QNAP qui offre à la fois deux ports Thunderbolt 3 et 10 GbE intégrés pour une collaboration multiplateforme entre les utilisateurs macOS et Windows. Aujourd'hui, nous allons examiner comment j'utilise cet appareil dans mon flux de travail à la fois sur et hors des plateaux de tournage et commerciaux.
J'utilise quatre disques durs Seagate IronWolf Pro de 16 To en RAID 5, ainsi que quatre SSD Seagate IronWolf 125 Go en RAID 5 avec un surprovisionnement de 23 %. Sur le papier, c'est un NAS très complet, et vu son prix, c'est bien le moins. Passons rapidement en revue les caractéristiques techniques de ce NAS, puis parlons un peu de mon utilisation.
Spécifications du QNAP TVS-872XT Thunderbolt NAS/DAS
| Processeur | Intel Core i5-8400T, 6 cœurs 1.7 GHz – 3.3 GHz | 
| GPU | Intel UHD Graphics 630 | 
| Transcodage accéléré par le matériel | Oui | 
| Mémoire système | 16 Go DDR4 (2x8 Go), extensible à 32 Go | 
| Baies de disques | 8 × 3.5 pouces SATAIII 6 Gb/s (remplaçable à chaud) | 
| Assistance Drive | Disque dur SATA 3.5 pouces Disque dur SATA 2.5 pouces SSD SATA 2.5 pouces | 
| Fente M.2 | 2 emplacements M.2 2280 PCIe Gen3 | 
| Ethernet | 1 port RJ10 45GBASE-T (10GbE, 5GbE, 2.5GbE, 1GbE) 2 ports Ethernet Gigabit RJ45 | 
| Assistance Thunderbolt | 2 ports Thunderbolt 3 via la carte d'extension incluse | 
| Emplacements PCIe | Emplacement 1 : PCIe Gen 3 x16 (processeur, occupé) Emplacement 2 : PCIe Gen 3 x4 (PCH) | 
| USB | 1 ports USB 3.2 Gen 1 2x USB 3.2 Gen 2 (Type A) 2x USB 3.2 Gen 2 (Type-C) | 
| Dimensions (LxHxP) | 7.41 x 12.96 x 11.01 pouces | 
| Consommation d'énergie | 65.03W (typique) 41.47 W (mode veille du disque dur) | 
Le jour, je suis technicien en imagerie numérique (DIT) et membre de l'IATSE Local 600, et je travaille principalement à New York. Un DIT, brièvement, est en charge du pipeline de couleurs sur les plateaux de cinéma, de télévision et commerciaux, travaillant avec le directeur de la photographie pour gérer l'image sortant de l'appareil photo afin d'obtenir l'apparence et la sensation qu'ils désirent. C'est beaucoup plus compliqué que cela, mais je simplifie à l'extrême par souci de brièveté.
En plus des responsabilités d'image et de couleur, je suis également en charge de la gestion de toutes les données sur le plateau, y compris le téléchargement des images de la caméra, le transcodage des négatifs de l'appareil photo numérique dans un format pour la post-production et la création de «dailies», qui sont des versions plus petites des images que le réalisateur et le directeur de la photographie doivent revoir à la fin de la journée. Comme vous pouvez l'imaginer, tout cela nécessite beaucoup de données et c'est là qu'intervient un appareil comme le QNAP TVS-872XT.
Le débit brut est le nom du jeu, avec des tournages multi-caméras capables de produire des téraoctets de données par jour, avec des heures de séquences devant être transcodées. Le plus gros goulot d'étranglement dans la plupart de ces cas est la vitesse d'entraînement. Alors que la production me fournira quelque chose comme un RAID externe à 2 baies pour la livraison en post-production (les G-RAID G-Tech sont une option populaire), ces disques s'étouffent rapidement lors du déchargement de plusieurs caméras et du transcodage du lot de médias précédent au en même temps.
Malheureusement pour moi, la production n'a généralement pas le budget pour m'acheter une matrice RAID XNUMX % flash haut de gamme pour chaque travail, donc décharger ce travail est essentiel. J'aime aussi avoir une copie du tournage sur quelque chose que je ramène à la maison comme sauvegarde supplémentaire. Voyons comment j'ai configuré mon QNAP et comment cette configuration m'aide à éviter les goulots d'étranglement susmentionnés.
QNAP TVS-872XT Configuration et performances du NAS Thunderbolt
J'utilise quatre disques durs Seagate IronWolf Pro de 16 To configurés en RAID 5, ce qui me donne environ 43 To utilisables. Cette configuration me permet de tolérer la panne d'un disque, ce qui est suffisant pour mon usage, car mes données sont toujours sauvegardées sur au moins deux autres supports. Le RAID 5 améliore considérablement ma vitesse de lecture, sans trop sacrifier la vitesse d'écriture, et ces disques de 16 To sont déjà assez rapides pour des disques durs mécaniques.
La capacité me laisse également amplement d'espace pour stocker des séquences de travaux pendant des mois, ainsi que pour enregistrer des séquences en direct pour une lecture par moi-même et le directeur de la photographie. Cela peut sembler exagéré, et franchement ça l'est un peu, mais je ne peux même pas compter le nombre de fois où j'ai entendu une production un mois après le tournage me demander si j'aurais une sauvegarde des images ; répondre à cela par un oui est toujours agréable.
J'utilise également quatre SSD Seagate IronWolf 125 de 1 To en RAID 5 avec un surprovisionnement de 23 %, pour une capacité utilisable totale de 2 To. Bien que le RAID 5 soit généralement déconseillé pour les SSD, ceux-ci ne fonctionnent pas 24 h/24 et 7 j/7. Seagate annonce une endurance en écriture de 1.4 pétaoctets pour ces disques, sans tenir compte du surprovisionnement. À titre de comparaison, cela signifierait que je devrais écrire 1 To par jour, sans interruption, pendant plus de deux ans et demi.
Comme avec les disques durs, je peux perdre un seul disque et continuer à travailler, et étant des SSD, il y a des pénalités de vitesse minimales. Les vitesses ne sont pas spectaculaires (par rapport à mes disques système NVMe internes), car ce sont des SSD SATA, mais nous y reviendrons plus tard. Les SSD sont utilisés pour transcoder et déplacer des séquences vers d'autres disques. Cela me permet de décharger des médias sur d'autres lecteurs et de transcoder en même temps sans ralentir l'un ou l'autre des processus, une sorte de fac-similé de hiérarchisation du stockage.
Je n'utilise pas de disques NVMe dans mon NAS car le type de travail que je fais ne bénéficie pas vraiment d'un cache, et malheureusement, les disques NVMe sont limités à seulement 2 voies de PCIe, ce qui serait plus lent que ma baie SSD . RAID0 est une option, mais je ne veux pas risquer cela. Plus d'informations à ce sujet plus tard également.
Ci-dessus, un test de vitesse de disque Blackmagic Design pour la matrice de disques durs via Thunderbolt 3. Il a été testé sur un Mac mini 2018 avec le processeur Intel i6-7B à 8700 cœurs et 32 Go de mémoire DDR4. Je vois généralement des vitesses d'environ 600 Mo/s en écriture et 1,550 5 Mo/s en lecture, et cela concerne les performances attendues pour quatre disques durs en RAID2.0. C'est tout à fait adéquat car le support le plus courant que je décharge est les cartes CFast 550, qui atteignent de toute façon environ XNUMX Mo/s.
Ci-dessus, le test Blackmagic pour la baie SSD et a été testé sur le même Mac mini 2018 que les disques durs. Les performances sont un peu décevantes et bien que les IOPS aléatoires soient bien meilleures que les disques rotatifs, mon travail consiste en des lectures et des écritures séquentielles à 99 %, comme c'est la nature des gros fichiers vidéo. Ce n'est cependant pas vraiment la faute du NAS, car ces disques sont goulottés par SATA. Même avec un surapprovisionnement, pour lequel l'outil intégré de QNAP est excellent, environ 1,300 1,750 Mo/s en écriture et XNUMX XNUMX Mo/s en lecture sont les meilleurs que vous obtiendrez avec les SSD SATA. Ces chiffres ne sont pas tout à fait reflétés dans la capture d'écran ci-dessus, mais les performances augmentent parfois.
Le bon
