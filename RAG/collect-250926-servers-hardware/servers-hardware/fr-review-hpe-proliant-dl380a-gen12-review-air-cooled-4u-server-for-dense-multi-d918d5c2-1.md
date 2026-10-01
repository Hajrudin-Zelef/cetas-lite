---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi-d918d5c2-1
title: "fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["blackwell", "compute", "distribution", "fp4", "gpu", "intel", "nvidia"]
source: docs/RAG/clean4/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2.md
source_anchor: ""
source_lines: [1, 31]
sha256: 3d32ad3eed3915595ef87a8074dd5b1f1a0914921eed2d63043195b1129f407b
---

# fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2

Le serveur HPE ProLiant Compute DL380a Gen12 s'adresse aux équipes d'IA d'entreprise qui recherchent une forte densité de calcul sans modifier la configuration de leurs racks. Ce châssis 4U à refroidissement par air s'intègre facilement, prend en charge jusqu'à huit GPU double largeur et offre une connectivité PCIe Gen5 complète. Il peut être configuré avec deux processeurs Intel Xeon 6, chacun doté de 144 cœurs, 4 To de mémoire DDR5 répartis sur 32 modules DIMM et seize baies NVMe E3.S pour un débit et une capacité élevés. L'objectif est simple : atteindre une capacité d'inférence de niveau production et un réglage fin précis à grande échelle sans recourir à un refroidissement liquide.
Pour les accélérateurs, HPE propose une gamme complète incluant les cartes NVIDIA H200 NVL, H100 NVL, L40S, L20, L4 et la RTX PRO 6000 Blackwell Server Edition, avec des options d'alimentation compatibles avec les composants à forte consommation. Dans ce test, nous nous concentrons sur la RTX PRO 6000 Server, qui offre un excellent compromis pour l'IA d'entreprise. Chaque carte embarque 96 Go de mémoire ECC GDDR7, une interface PCIe Gen5 x16, des cœurs Tensor compatibles FP4 et une enveloppe thermique de 600 W, adaptée aux racks refroidis par air. Notre configuration était équipée de quatre cartes, un point de départ judicieux pour l'inférence à haut débit et l'optimisation ciblée, avec une marge de progression.
HPE complète la plateforme avec les éléments opérationnels essentiels. iLO 7 gère la configuration hors bande, l'état du système et la gestion de l'alimentation, grâce à la Silicon Root of Trust (une enclave sécurisée garantissant l'intégrité du firmware), la prise en charge du chiffrement RSA 4096 bits et un module iLO DC-MHS détachable qui renforce la vérification de la chaîne d'approvisionnement. Le serveur s'intègre également à la plateforme HPE Private Cloud AI pour une gouvernance multi-équipes et des déploiements reproductibles à grande échelle.
HPE ProLiant Compute DL380a Gen12 – Spécifications techniques
| Catégories | Spécifications | 
|---|---|
| Type de processeur | HPE ProLiant Compute DL380a Gen12 | 
| Famille de processeur | Processeurs Intel® Xeon® Scalable de 6e génération | 
| Cœurs de processeur disponibles | De 64 à 144 cœurs, selon le processeur | 
| Nombre de processeurs | 2 | 
| La vitesse du processeur | Jusqu'à 2.4 GHz, selon le processeur | 
| La mémoire maximale | RDIMM 4 To (2 To par processeur) | 
| Logements pour la mémoire | 32 emplacements DIMM | 
| Type de mémoire | Mémoire intelligente HPE DDR5 | 
| Protection de la mémoire | RAS : ECC avancé, mémoire de secours en ligne, mise en miroir, fonctionnalité de canal combiné (verrouillage), mémoire tolérante aux pannes HPE Fast (ADDDC) | 
| Assistance routière | SFF NVMe et EDSFF | 
| Sécurité | Cadre de verrouillage en option, détection d'intrusion et module HPE TPM 2.0 intégré | 
| Gestion de l'infrastructure | HPE iLO Standard avec provisionnement intelligent (intégré), HPE OneView Standard (nécessite un téléchargement) • En option : HPE iLO Advanced et HPE OneView Advanced (licences requises) | 
| Source d'alimentation | Jusqu'à 8 M-CRPS. Redondance simple 1+1 pour la carte mère. Redondance double 2+1 pour les GPU. | 
| Connecteurs d'extension | 6 | 
| Ventilateurs du système | 4 ventilateurs à double rotor et 8 ventilateurs à simple rotor remplaçables à chaud inclus | 
| Facteur de forme | rack 4U | 
| Garantie | 3/3/3 : Garantie du serveur | 
Conception et assemblage du serveur HPE ProLiant DL380a Gen12
Le serveur rack HPE ProLiant Compute DL380a Gen12 est un serveur 4U à deux sockets conçu pour les déploiements évolutifs et hautes performances. Mesurant 6.88 x 17.63 x 31.60 cm, il combine une puissance de calcul CPU et GPU dense avec un refroidissement par air efficace pour un fonctionnement fiable même sous fortes charges de travail.
Pesant entre 82.7 et 137.8 kg selon la configuration, le châssis prend en charge les composants haute capacité, l'alimentation redondante et offre un accès frontal aisé pour la maintenance. Sa conception privilégie la performance, l'évolutivité et une gestion thermique performante, ce qui le rend idéal pour les environnements d'entreprise et de centres de données.
Côté stockage, le HPE ProLiant DL380a Gen12 propose des configurations à 4 ou 8 baies aux formats SFF ou EDSFF. Notre modèle de test était équipé du kit de baie avant HPE DL380a Gen12 NS204i-u, prenant en charge deux périphériques de démarrage NVMe M.2 remplaçables à chaud. Le châssis comprenait également huit baies 2.5 pouces, occupées par deux SSD U.3 HPE d'une capacité totale de 15.36 To. HPE propose plusieurs options de baies avant, offrant une grande flexibilité d'adaptation aux différents besoins de déploiement.
L'unité peut être transportée grâce à ses deux poignées latérales, ce qui implique qu'au moins deux personnes sont nécessaires pour la mise en rack et l'installation en toute sécurité. Elle utilise un kit de rails 2U avec rails télescopiques, permettant une installation aisée et une maintenance facilitée sans démontage complet du rack.
À l'arrière du serveur HPE ProLiant DL380a Gen12, l'agencement optimisé favorise la circulation de l'air, l'extensibilité et la facilité de maintenance. Le système prend en charge jusqu'à huit alimentations MCRPS (1 à 8) grâce à un panneau de ventilation intégré, garantissant un refroidissement optimal même en pleine charge. L'extensibilité est importante, avec plusieurs emplacements PCIe Gen5 x16 (emplacements 1 à 6) compatibles avec les cartes d'extension intégrées et optionnelles, ainsi que les emplacements OCP A et B pour une configuration flexible des adaptateurs réseau.
La connectivité comprend un port réseau iLO dédié, plusieurs ports USB 3.2 Gen 1 et un port VGA pour la gestion locale. Il est important de noter que l'emplacement 1 est disponible uniquement lorsque le câble HPE DL380a Gen12 4EDSFF Direct Cable for NVD (P74716-B21) est installé et ne peut pas être utilisé avec des disques NVMe SFF. Quant à l'emplacement 4, il n'est pas pris en charge dans les configurations comportant 4 ou 8 GPU DW.
L'alimentation du serveur HPE ProLiant DL380a Gen12 est assurée par des kits d'alimentation modulaires M-CRPS Titanium remplaçables à chaud. Les modèles compatibles incluent les versions 1 500 W (P67244-B21), 2 400 W (P67252-B21) et 3 200 W (P67248-B21). Le système prend en charge jusqu'à huit alimentations, offrant une redondance N+1 pour garantir un fonctionnement continu même en cas de défaillance d'un module d'alimentation. Les besoins en énergie et leur distribution peuvent varier en fonction de la configuration des GPU. Notre modèle de test était équipé de cinq alimentations M-CRPS de 2 400 W, fournissant une capacité suffisante pour alimenter les quatre GPU du système (TDP de 600 W) tout en assurant une redondance fiable.
