---
id: collect-261001-general-networking/general-networking/fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc-2
title: "fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc"
domain: general-networking
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "ethernet", "exploit"]
source: docs/RAG/collect-261001-general-networking/fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc.md
source_anchor: ""
source_lines: [19, 45]
sha256: 004c09dece51a438300abc4524f525dbc3cc2198271050dacec38513638a4680
---

# fr-review-maximizing-vsan-8-esa-potential-with-western-digital-ultrastar-dc-sn65-352ebbbc

Les Dell PowerEdge R6625 ont beaucoup à offrir, ce qui les rend idéaux pour notre cas d'utilisation VMware vSAN HCI. Du côté du calcul, il s’agit de serveurs à double socket, prenant en charge suffisamment de puissance de traitement pour gérer n’importe quelle charge de travail. Dans notre configuration, nous avons deux processeurs AMD EPYC 9554 à 64 cœurs avec 128 Go de RAM par nœud. Le R6625 prend également en charge jusqu'à dix SSD NVMe de 2.5 pouces, ce qui leur confère une bonne densité, même si notre version s'est concentrée sur quatre et huit SSD par nœud.
Le PowerEdge R6625 prend en charge de nombreuses options du point de vue de la mise en réseau, avec un emplacement OCP 3.0 et trois emplacements PCIe Gen5. Cette configuration nous permet d'exploiter 100 GbE et au-delà, avec des emplacements supplémentaires disponibles pour d'autres appareils. En parlant de commutation, le Dell PowerSwitch S5232F-ON relie ce cluster. Il s'agit d'un commutateur multicouche 1U avec 32 ports 100 GbE QSFP28 et 2 ports 10GbE SFP+.
Il convient de noter ici que le travail que nous avons effectué dans ce rapport porte sur certains composants qui ne sont pas encore sur VMware vSAN HCI ou qui ne sont pas officiellement pris en charge par Dell à ce jour. Cela dit, nous avons suivi les meilleures pratiques VMware pour vSAN tout au long de ces tests et avons demandé aux ingénieurs VMware vSAN d'examiner les données.
Performances de vSAN 8 mise à jour 2
Dans notre approche pour tester vSAN 8 de VMware avec ESA, nous nous sommes concentrés sur une configuration de quatre et huit SSD par nœud. Avec quatre hôtes, cela représentait 16 ou 32 SSD au total.
Nous avons exploité Ethernet 100 Go dans le cluster, ce qui nous a théoriquement permis de nous concentrer sur le stockage NVMe pour voir où il commence à saturer les ressources système.
Une autre variable mesurée dans ce rapport concerne l'impact du RAID 1 et du RAID 5 sur les performances globales de stockage au sein de notre cluster vSAN. Historiquement, le RAID 1 a été une stratégie de stockage courante pour les déploiements vSAN. La configuration recommandée pour ESA est « Stratégie par défaut vSAN ESA – RAID 5 », qui offre des performances quasi identiques avec une capacité bien supérieure. Le RAID 1 induit une perte de 50 % immédiate, tandis que le RAID 5 limite la surcharge liée à la parité, offrant ainsi une capacité utilisable beaucoup plus importante.
Pour mesurer les performances de notre cluster VMware vSAN exécutant ESXi 8 Update 2, nous avons utilisé HCIbench. HCIbench a été développé pour simplifier les tests des clusters HCI, dans lesquels vous devez déployer des machines virtuelles de travail, générer des disques virtuels et orchestrer les tests sur l'ensemble d'un cluster tout en regroupant tous ces résultats. Lors de nos tests, nous avons utilisé les paramètres suivants pour nos quatre configurations ESA :
Déploiement de la machine virtuelle HCIBench :
- 16 VM
- 8 disques de données de 50 Go par VM
- 16 processeurs virtuels par VM
- 8 Go de RAM par VM
- Durée de test de 3600 XNUMX secondes par charge de travail
- 8 threads par disque pour les charges de travail séquentielles et 16 threads pour les charges de travail aléatoires
| DONNÉES Western Digital Ultrastar DC SN655 vSAN ESA | 1024 XNUMX XNUMX Mo/s en écriture séquentielle | Lecture séquentielle 1024 XNUMX XNUMX Mo/s | IOPS en écriture aléatoire 4K | IOPS en lecture aléatoire 4K | 8K aléatoires 70/30 IOPS | 
|---|---|---|---|---|---|
| vSAN ESA 4-SSD par nœud RAID1 | 10,914 | 17,638 | 378,861 | 563,054 | 318,640 | 
| vSAN ESA 4-SSD par nœud RAID5 | 10,999 | 17,726 | 476,770 | 524,076 | 301,960 | 
| vSAN ESA 8-SSD par nœud RAID1 | 12,702 | 24,128 | 520,632 | 526,504 | 323,674 | 
| vSAN ESA 8-SSD par nœud RAID5 | 14,336 | 21,994 | 504,508 | 523,666 | 292,557 | 
Nous divisons les données de performances en deux sections qui se chevauchent : quatre contre huit SSD par nœud et RAID1 contre RAID5. Dans notre comparaison initiale, nous nous sommes concentrés sur le point auquel les performances commencent à saturer à mesure que davantage de disques sont ajoutés au cluster. Avec quatre hôtes, nous avons examiné 16 SSD dans le cluster, contre 32 lorsque nous avons augmenté chaque nœud à huit SSD. Le deuxième aspect est la différence significative entre la politique RAID1 ESA et la politique RAID5 ESA recommandée. Les avantages de l'utilisation de RAID5 sont assez évidents : les utilisateurs bénéficient d'une énorme quantité de capacité disponible grâce à la politique de protection des données, associée aux denses SSD Western Digital.
Nous avons remarqué quelques tendances apparaître en examinant les données collectées lors de nos tests. Si vous vous concentrez sur la bande passante totale, le cluster VMware vSAN ESA a connu une amélioration significative de ses performances en passant de quatre à huit SSD par nœud. Nous avons mesuré environ 10.9 Go/s en écriture séquentielle à partir des configurations RAID1 et RAID5 avec une configuration à 4 SSD. En lecture séquentielle, nous avons mesuré 17.6-17.7 Go/s en modes RAID1 et RAID5. En passant aux charges de travail aléatoires, nous avons constaté que RAID5 avait un avantage dans la configuration à 4 SSD, mesurant 476 4 IOPS en écriture aléatoire 379K, contre 1 4 IOPS en RAID1. En lecture aléatoire 563K, les deux configurations ont commencé à se rapprocher, le RAID524 mesurant 5 8 IOPS à 70 30 IOPS en RAID1. En 319K 302/5, la différence était beaucoup plus étroite, RAIDXNUMX ayant un bord mesurant XNUMX XNUMX IOPS contre XNUMX XNUMX IOPS pour RAIDXNUMX.
En passant à huit SSD par nœud, nous avons constaté une amélioration des performances, mais cela n’a pas doublé les performances par rapport à quatre SSD. C’est là que les limites supérieures de performances de vSAN et la limite de bande passante de 100 GbE sur nos nœuds ont commencé à apparaître. Oui, des connexions réseau plus rapides sont disponibles, mais cela s’éloigne davantage des déploiements HCI et typiques des PME/PMI.
Concernant la bande passante en écriture, RAID5 avait le dessus, mesurant 14.3 Go/s contre 12.7 Go/s pour RAID1. RAID1 était en tête en termes de bande passante de lecture avec 24.1 Go/s contre RAID5 avec 22 Go/s. Avec des transferts d'écriture 4K aléatoires, RAID1 et RAID5 étaient proches, bien que nous ayons mesuré 521 1 IOPS dans R505 contre 5 4 IOPS dans R527. En examinant les performances de lecture aléatoire 1K, nous nous sommes retrouvés confrontés à une limite de performances du cluster vSAN où quatre et huit configurations SSD étaient très proches les unes des autres. Ici, avec notre configuration de huit SSD, nous avons mesuré 524 5 IOPS en RAID8 et 70 30 IOPS en RAID1. En 324K 293/5, le RAIDXNUMX avait finalement l'avantage, avec XNUMX XNUMX IOPS contre XNUMX XNUMX IOPS en RAIDXNUMX.
Il est important de noter qu’avec VMware vSAN ESA, la configuration RAID1 ou RAID5 n’est pas une décision gravée dans le marbre. Il s'agit d'une politique de stockage appliquée au niveau d'un vDisk de VM, ce qui signifie que les deux peuvent coexister simultanément. Supposons donc que vous disposiez d’une machine virtuelle de base de données dans laquelle chaque bit d’E/S compte ; donnez-lui une politique RAID1, où tout le reste est sur RAID5. De cette façon, vous optimisez au mieux votre utilisation de l’espace.
Réflexions finales
