---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb-3
title: "fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["gpu", "inference", "intel", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb.md
source_anchor: ""
source_lines: [89, 135]
sha256: 77cf2aebacf1e82bf5ce47f92290b2025d8e6b149f21761ffb4a0d69669287e9
---

# fr-review-supermicro-hyper-superserver-sys-112h-tn-review-efficiency-and-power-w-799f43eb

| 2.5 milliard | 24.928 secondes | 
| 5 milliard | 53.489 secondes | 
| 10 milliard | 113.727 secondes | 
| 25 milliard | 308.218 secondes | 
| 50 milliard | 674.299 secondes | 
Test de vitesse du disque Blackmagic
Lors du test de vitesse de disque de Blackmagic, le SSD Micron 7450 NVMe a produit des vitesses de lecture de 3,627.4 2,849.5 Mo/an et de XNUMX XNUMX Mo/s en écriture.
Test de vitesse Blackmagic RAW
Le Blackmagic RAW Speed Test est un outil d'analyse comparative des performances conçu pour mesurer les capacités d'un système à gérer la lecture et l'édition vidéo à l'aide du codec Blackmagic RAW. Il évalue la capacité d'un système à décoder et à lire des fichiers vidéo haute résolution, en fournissant des fréquences d'images pour le traitement basé sur le CPU et le GPU.
Le Supermicro Hyper 112H-TN a obtenu 116 FPS avec un processeur 8K. Nous n'avons pas effectué la partie GPU de ce test.
7-Zip
Le test de mémoire intégré du populaire utilitaire 7-Zip mesure les performances du processeur et de la mémoire d'un système pendant les tâches de compression et de décompression, fournissant ainsi une indication de la capacité du système à gérer des opérations gourmandes en données.
Ici, le Supermicro Hyper 112h-tn atteint une note totale de 205.449 GIPS, nettement inférieure aux 271.217 GIPS obtenus par la configuration double Xeon 6780E.
| Compression à 7 zips | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Xeon 6780E (256 Go DDR5) | Xeon 6766E (256 Go DDR5) | 
| Compression |  |  |  | 
| Utilisation actuelle du processeur | 5287 % | 5,891 % | 4,768 % | 
| Courant nominal/utilisation | 4.647 GIPS | 4.985 GIPS | 4.614 GIPS | 
| Courant | 245.699 GIPS | 293.689 GIPS | 220.001 GIPS | 
| Utilisation résultante du processeur | 5296 % | 5,603 % | 5,103 % | 
| Évaluation/utilisation résultante | 4.642 GIPS | 4.954 GIPS | 4.638 GIPS | 
| Note résultante | 245.823 GIPS | 277.670 GIPS | 236.910 GIPS | 
| Décompression |  |  |  | 
| Utilisation actuelle du processeur | 6236 % | 5,962 % | 5,798 % | 
| Courant nominal/utilisation | 4.261 GIPS | 4.550 GIPS | 4.152 GIPS | 
| Courant | 265.709 GIPS | 271.266 GIPS | 240.693 GIPS | 
| Utilisation résultante du processeur | 6236 % | 5,832 % | 6,029 % | 
| Évaluation/utilisation résultante | 4.341 GIPS | 4.540 GIPS | 4.161 GIPS | 
| Note résultante | 269.373 GIPS | 264.764 GIPS | 250.853 GIPS | 
| Note totale |  |  |  | 
| Utilisation totale du processeur | 5751 % | 5,717 % | 5,566 % | 
| Note totale/utilisation | 4.491 GIPS | 4.747 GIPS | 4.399 GIPS | 
| Note totale | 257.598 GIPS | 271.217 GIPS | 243.882 GIPS | 
Inférence UL Procyon AI
UL Procyon AI Inference est conçu pour évaluer les performances d'une station de travail dans des applications professionnelles. Il est important de noter que ce test n'exploite pas les capacités de plusieurs processeurs. Plus précisément, cet outil évalue la capacité de la station de travail à gérer les tâches et les flux de travail basés sur l'IA, fournissant une analyse détaillée de son efficacité et de sa rapidité de traitement des algorithmes et applications d'IA complexes.
Les résultats du Supermicro Hyper 112h-tn (Xeon 6780E, 512 Go DDR5) donnent un aperçu de la capacité du système à gérer diverses tâches pilotées par l'IA. D'après le tableau ci-dessous, il est clair que le serveur peut traiter efficacement des modèles plus simples comme MobileNet V3 et ResNet 50, avec des temps d'inférence de 5.23 ms et 6.60 ms, respectivement. Cependant, à mesure que la complexité des modèles augmente, comme avec Real-ESRGAN, le temps d'inférence augmente considérablement pour atteindre 966.48 ms. Le score global de 168 indique que le Xeon 6780E est performant mais peut être mieux adapté aux charges de travail d'IA modérées et bénéficierait d'une configuration à double processeur.
| Temps d'inférence moyens UL Procyon (le plus faible est le mieux) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | 
| Mobile Net V3 | 5.23 ms | 
| ResNet 50 | 6.60 ms | 
| Création V4 | 23.12 ms | 
| Deep Lab V3 | 24.61 ms | 
| YOLO V3 | 35.66 ms | 
| Réel-ESRGAN | 966.48 ms | 
| Note globale | 168 | 
Conclusion
Alimenté par la série Intel Xeon 6, le Supermicro Hyper SuperServer SYS-112H-TN offre une combinaison convaincante de performances et d'efficacité, ce qui en fait un concurrent sérieux pour diverses applications d'entreprise grand public. Son processeur 6780E équipé lui permet d'exceller dans les charges de travail multithread, offrant des performances robustes tout en conservant l'efficacité énergétique. Les gains d’efficacité sont particulièrement impressionnants par rapport aux systèmes plus anciens qu’ils sont susceptibles de remplacer.
Au-delà des capacités des processeurs Xeon 6, le Supermicro Hyper 1U SYS-112H-TN est équipé d'une gamme de fonctionnalités qui améliorent sa polyvalence et ses performances. Il dispose de deux emplacements PCIe 5.0 x16 et d'un emplacement PCIe 5.0 AIOM supplémentaire, permettant une connectivité haut débit pour les GPU, les interfaces réseau et autres cartes d'extension.
Xeon 6 offre jusqu'à 144 cœurs d'efficacité dans le Supermicro Hyper 1U SYS-112H-TN, ce qui est excellent pour les charges de travail qui n'ont pas besoin du silicium le plus puissant. Il devrait constituer un élément de base flexible pour les intégrateurs de systèmes et les entreprises qui souhaitent un serveur abordable mais moderne, pour une grande variété de tâches.
