---
id: collect-250926-servers-hardware/servers-hardware/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5-4
title: "fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "gpu", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5.md
source_anchor: ""
source_lines: [90, 134]
sha256: 6501381fe40ce3900af9456c3bc3bd03382a9f06ebe0a54a8aa15409d8f63129
---

# fr-review-amd-epyc-turin-review-192-cores-of-zen-5-c681c8c5

Le test de compression 7-Zip démontre l'efficacité du 9755, atteignant 443.029 GIPS à 5613 9755 % d'utilisation du processeur, le plus élevé de la gamme EPYC. Ce niveau de performance rend le 9575 idéal pour les tâches de compression de données intensives, telles que celles des solutions de sauvegarde ou d'archivage, où une gestion rapide des données est essentielle. Le 394.9F (SMT désactivé) fonctionne également bien, atteignant 9755 GIPS. Ces résultats indiquent que les 9575 et XNUMXF ont de solides capacités de compression et de décompression, avec un débit CPU élevé et une gestion efficace des ressources pour les applications qui exploitent des taux de compression élevés ou traitent de gros volumes de données.
| Référence de compression à 7 zips (Plus c'est haut, mieux c'est) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c, SMT désactivé) |  | 
| Compression |  |  |  |  | 
| Utilisation actuelle du processeur | 4302 % | 5233 % | 4406 % |  | 
| Note actuelle/utilisation | 5.830 GIPS | 7.597 GIPS | 7.975 GIPS |  | 
| Courant | 250.827 GIPS | 397.536 GIPS | 351.358 GIPS |  | 
| Utilisation résultante du processeur | 4041 % | 5306 % | 4555 % |  | 
| Évaluation/utilisation résultante | 5.804 GIPS | 7.720 GIPS | 8.070 GIPS |  | 
| Note résultante | 234.317 GIPS | 409.652 GIPS | 367.358 GIPS |  | 
| Décompression |  |  |  |  | 
| Utilisation actuelle du processeur | 4322 % | 6041 % | 5017 % |  | 
| Note actuelle/utilisation | 7.078 GIPS | 8.065 GIPS | 8.483 GIPS |  | 
| Courant | 305.909 GIPS | 487.263 GIPS | 425.580 GIPS |  | 
| Utilisation résultante du processeur | 4556 % | 5921 % | 4940 % |  | 
| Évaluation/utilisation résultante | 6.577 GIPS | 8.045 GIPS | 8.569 GIPS |  | 
| Note résultante | 299.163 GIPS | 476.405 GIPS | 422.441 GIPS |  | 
| Note totale |  |  |  |  | 
| Utilisation totale du processeur | 4298 % | 5613 % | 4747 % |  | 
| Note totale/utilisation | 6.190 GIPS | 7.883 GIPS | 8.319 GIPS |  | 
| Note totale | 266.740 GIPS | 443.029 GIPS | 394.900 GIPS |  | 
Blender OptiX est une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire CLI Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Ici, le 9755 surpasse légèrement avec un score de 2,606.54 9965 échantillons par minute dans la scène « Monster », suivi de près par le 2,558.43 avec 1,700.65 2,038.71. Les deux surpassent le Genoa et le Bergamo, avec des scores respectifs de 3 9755 et 9965 3, soulignant l'avantage que les processeurs basés sur Turin apportent au rendu XNUMXD. Ces scores mettent en évidence l'adéquation du XNUMX et du XNUMX aux tâches de modélisation XNUMXD haute résolution et aux flux de travail de création de contenu impliquant un rendu de scène complexe.
| Blender 4.0 Échantillons de processeur par minute (plus c'est élevé, mieux c'est) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) | 
| Monster | 2,558.43 | 2,606.54 | 1,196.15 | 1,700.65 | 2,038.71 | 
| Brocanteur | 1,866.65 | 1,843.48 | 802.00 | 1,101.84 | 1,382.58 | 
| Salle de classe | 1,270.17 | 1,251.54 | 637.13 | 869.48 | 1,045.96 | 
La suite de tests d' inférence IA Procyon de UL évalue les performances de différents moteurs d'inférence IA utilisant des réseaux neuronaux de pointe. Ces tests ont été exécutés uniquement sur le processeur. Chaque valeur représente un temps d'inférence moyen ; plus la valeur est basse, meilleures sont les performances. La dernière ligne indique un score global ; plus la valeur est élevée, meilleures sont les performances.
Ici, l'EPYC 9575F avec SMT désactivé offre des temps d'inférence solides sur divers réseaux neuronaux, avec des performances exceptionnelles dans MobileNet V3 (7.02 ms) et ResNet 50 (10.45 ms). Ces scores dépassent à la fois Gênes (3.63 ms pour MobileNet et 6.34 ms pour ResNet) et Bergame (4.16 ms et 8.22 ms, respectivement) pour les charges de travail d'IA, ce qui suggère que la configuration SMT-off du 9575F donne la priorité au contrôle de la latence, ce qui le rend très efficace pour les tâches d'IA en temps réel où des temps d'inférence faibles sont essentiels.
Les EPYC 9965 et 9755, bien que moins rapides dans les tâches d'inférence unique que Genoa, maintiennent de solides performances dans tous les domaines, avec les 9755 ms du 20.07 sur MobileNet et les 20.11 ms sur ResNet 50 indiquant des performances stables et fiables adaptées aux charges de travail modérées.
Alors que Gênes et Bergame offrent une latence systématiquement plus faible, le 9575F, en particulier avec SMT désactivé, offre un excellent contrôle de la latence.
| Moyenne UL Procyon | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC9575F (64c) | Gênes (2p/96c) | Bergame (2p/128c) |  | 
| Temps d'inférence (le plus bas est le mieux) |  |  |  |  |  |  | 
| Mobile Net V3 | 32.86 ms | 20.07 ms | 7.02 ms | 3.63 ms | 4.16 ms |  | 
| ResNet 50 | 38.63 ms | 20.11 ms | 10.45 ms | 6.34 ms | 8.22 ms |  | 
| Création V4 | 116.38 ms | 66.22 ms | 33.69 ms | 25.99 ms | 30.68 ms |  | 
| Deep Lab V3 | 58.16 ms | 33.02 ms | 19.38 ms | 25.33 ms | 30.57 ms |  | 
| YOLO V3 | 84.20 ms | 38.47 ms | 24.54 ms | 34.13 ms | 41.38 ms |  | 
| RÉEL-ESRGAN | 2923.75 ms | 1600.73 ms | 1219.26 ms | 2524.03 ms | 2301.35 ms |  | 
| Note globale (Plus c'est haut, mieux c'est) | 44 | 81 | 148 | N/D | N/D |  | 
L'outil d'analyse comparative des performances du Blackmagic RAW Speed Test mesure les capacités d'un système en matière de gestion de la lecture et du montage vidéo à l'aide du codec Blackmagic RAW. Il évalue la capacité d'un système à décoder et à lire des fichiers vidéo haute résolution, en fournissant des fréquences d'images pour le traitement basé sur le CPU et le GPU.
Dans ce test, l'EPYC 9755 est en tête avec un impressionnant 174 ips en décodage CPU 8K, idéal pour la lecture de vidéos haute résolution dans les environnements de production multimédia. Le 9575F (SMT désactivé) suit de près avec 154 ips, surpassant légèrement le 9965 à 134 ips. Ces résultats suggèrent que les 9755 et 9575F offrent le meilleur équilibre entre cohérence de la fréquence d'images et vitesse de décodage, crucial pour les flux de travail de montage et de production vidéo qui gèrent des fichiers vidéo volumineux et de haute qualité.
| Blackmagic RAW (plus c'est haut, mieux c'est) | AMD EPYC 9965 (192c) | AMD EPYC 9755 (128c) | AMD EPYC 9575F (64c) | 
| CPU 8K | 134 images/s | 174 images/s | 154 images/s | 
La série AMD EPYC Turin 9005 marque une avancée significative dans l'informatique d'entreprise, offrant des performances, une efficacité et une adaptabilité exceptionnelles sur diverses charges de travail. Alimentés par l'architecture Zen 5, ces processeurs sont spécialement conçus pour répondre aux exigences croissantes des environnements IA, cloud et HPC tout en offrant une évolutivité inégalée pour les centres de données cherchant à optimiser les performances et la consommation d'énergie.
La série EPYC 9005 se distingue par sa capacité à répondre aux divers besoins des entreprises modernes avec des configurations allant des modèles d'entrée de gamme à haute fréquence aux puces à très forte densité de cache, en passant par les centrales multicœurs. Qu'il s'agisse d'inférences d'IA en temps réel, de dynamique des fluides computationnelle, d'analyses de données à grande échelle ou de rendu 3D haute résolution, la gamme EPYC offre une réactivité monothread et une efficacité multithread. Des fonctionnalités avancées telles que la prise en charge de la mémoire DDR12 à 5 canaux, les voies PCIe 5.0 et le calcul sécurisé et confidentiel d'AMD font de cette série une mise à niveau des performances et une solution complète pour les centres de données tournés vers l'avenir.
