---
id: collect-261001-huawei/huawei/fr-review-the-micron-6550-ion-ssd-gen5-performance-energy-efficiency-and-high-ca-d75c9796-2
title: "fr-review-the-micron-6550-ion-ssd-gen5-performance-energy-efficiency-and-high-ca-d75c9796"
domain: huawei
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["gpu", "nand", "nvidia"]
source: docs/RAG/collect-261001-huawei/fr-review-the-micron-6550-ion-ssd-gen5-performance-energy-efficiency-and-high-ca-d75c9796.md
source_anchor: ""
source_lines: [23, 54]
sha256: 29554b158cfabe2052b16423d5b616700e6346686d10ec1e814a431aae047c32
---

# fr-review-the-micron-6550-ion-ssd-gen5-performance-energy-efficiency-and-high-ca-d75c9796

| Facteurs de forme | E3.S-1T (7.5 mm), U.2 (15 mm), E1.L (9.5 mm) | 
| Capacités | 30.72 To – 61.44 To (mêmes spécifications de performances pour tous les SKU) | 
| 128 128 lectures séquentielles, QDXNUMX | 12,000 Mo / s | 
| 128 128 séquences d'écriture, QDXNUMX | 5,000 Mo / s | 
| Lectures aléatoires 4K, QD512 | 1,600,000 IOPS | 
| Écritures aléatoires 4K, QD128 | 70,000 IOPS | 
| Endurance (pour une garantie de 5 ans) |  | 
| Tuning Moteur |  | 
| Fiabilité (UBER = taux d'erreur binaire non corrigible) |  | 
| Fonctionnalité |  | 
| Garantie | 5 ans | 
Validation des performances synthétiques du Micron 6550 ION
Avec le Micron 6550 ION dans le laboratoire StorageReview, nous avons mis la version U.2 du disque à l'épreuve dans une plateforme de test Dell PowerEdge R760. Micron a inclus les chiffres qu'ils ont mesurés dans leur laboratoire, que nous avons pu valider lors de nos tests, aux côtés du SSD Solidigm P5336 de 61.44 To basé sur QLC, le seul SSD largement commercialisé disponible à ce niveau de capacité.
En plus de la validation des performances que nous avons effectuée sur le Micron 6550 ION, nous avons testé les états d'alimentation pour limiter la consommation d'énergie du Micron 6550 ION. Alors que tous les périphériques NVMe peuvent fonctionner sans restriction (consommant jusqu'à 25 W d'énergie), les SSD NVMe prennent en charge la réduction de l'utilisation maximale pour économiser sur la consommation d'énergie. La réduction de la consommation d'énergie peut avoir un impact significatif sur les performances de certains disques, mais le Micron 6550 ION est optimisé pour une consommation d'énergie plus faible, offrant des performances de centre de données dans une enveloppe de puissance plus petite.
Avant de commencer nos tests de performances, nous avons mis le Micron 6550 ION en mode PS1 et laissé le Solidigm P5335 en mode PS0 ou sans restriction. Le Micron 6550 ION prend en charge huit états d'alimentation (PS0 à PS7), qui vont de 25 W maximum à seulement 10 W dans la limite inférieure. PS1 définit la consommation d'énergie maximale de ce SSD à 20 W. En comparaison, le Solidigm P5336 prend en charge trois états d'alimentation (PS0 à PS2), qui vont de 25 W à 10 W par incréments de 5 à 10 W.
Chaque SSD a été soumis à une charge de travail de préconditionnement séquentielle de 128 K pour cette comparaison de performances, où les disques ont été remplis deux fois. Une fois la précondition terminée, les disques ont été testés pour des performances de lecture et d'écriture séquentielles de 128 K dans l'état préconditionné. La charge de travail a ensuite été commutée sur une précondition aléatoire de 16 K pour deux remplissages de disques, et les 16 K restants ont été testés en lecture et en écriture.
| Taille de bloc | Opération | Consommation électrique moyenne en microns à l'état d'alimentation PS1 | Fil/file d'attente de charge de travail FIO | Performances attendues du Micron 6550 ION | Performances mesurées du Micron 6550 ION | Performances mesurées du Solidigm P5336 | 
|---|---|---|---|---|---|---|
| 16KB | Lecture aléatoire | 17.9 W | 8T/256Q | > 700 XNUMX IOP > 11.2 Go/s | 11.6GB / s 707 XNUMX IOP | 7.4GB / s 454 XNUMX IOP | 
| 16KB | Écriture aléatoire | 17.5 W | 2T/1Q | > 70 XNUMX IOP > 1.1 Go/s | 1.4GB / s 85.4 XNUMX IOP | 850MB / s 51.9 XNUMX IOP | 
| 128KB | Lecture séquentielle | 17.5 W | 1T/64Q | > 12 Go/s | 12.7GB / s | 7.5GB / s | 
| 128KB | Écriture séquentielle | 17.5 W | 1T/16Q | > 5 Go/s | 8.2GB / s | 3.2GB / s | 
Le Micron 6550 ION de notre laboratoire a maintenu l'alignement avec les performances citées du Micron 6550 ION et, dans certains cas, a dépassé la fiche technique de Micron.
Dans notre charge de travail de lecture aléatoire de 16 6550 octets, le Micron 11.6 ION a maintenu un débit de 7.4 Go/s, contre 5536 Go/s pour le Solidigm P16. En régime permanent dans la charge de travail d'écriture aléatoire de 6550 1.4 octets, le Micron 850 ION pouvait offrir 5336 Go/s, contre 128 Mo/s pour le P6550. En passant à la charge de travail de lecture séquentielle de 12.7 7.5 octets, le Micron 5336 ION a mesuré 128 Go/s, contre 6550 Go/s pour le Solidigm P8.2. En passant à l'écriture séquentielle de 3.2 5336 octets, le Micron XNUMX ION a mesuré XNUMX Go/s contre XNUMX Go/s pour le Solidigm PXNUMX.
En le comparant au Solidigm P5336, nous avons noté de nombreux domaines dans lesquels la nouvelle interface Gen5 du Micron 6550 ION excellait par rapport à l'interface Gen4 du P5336, ce qui était attendu. La différence entre les disques Gen4 et Gen5 concerne la quantité de bande passante que chaque disque peut transmettre sur le réseau. L'interface PCIe Gen4 avec quatre voies de connectivité est capable d'environ 7 Go/s, tandis que la nouvelle interface PCI Gen5 double ce débit avec 14 Go/s. L'utilisation par Micron de la NAND TLC lui confère un autre avantage de conception par rapport aux concurrents basés sur QLC.
Performances de charge de travail du Micron 6550 ION AI
Les charges de travail liées à l'IA mettent les administrateurs à rude épreuve en exigeant des serveurs hautes performances et une efficacité énergétique optimale. Micron a comparé le 6550 ION au Solidigm P5336 afin de démontrer ses capacités à l'aide de quatre charges de travail distinctes : NVIDIA® Magnum IO™, l'entraînement et la capture Unet3D, et la sauvegarde de modèles d'IA. Les tests GPUDirect Storage (GDS) permettent un transfert direct de données entre le stockage et les GPU en contournant la mémoire du processeur, ce qui réduit la surcharge et la latence liées aux transferts de données.
Conclusion
Le passage aux SSD d'entreprise Gen5 a donné aux concepteurs de serveurs plus de flexibilité que jamais en matière de stockage. Le format E3.S a ouvert la voie à une densité de disques extrêmement élevée. Avec l'E3.S, les serveurs 1U standard peuvent contenir 20 SSD E3.S, contre 8 à 10 disques U.2. Les serveurs 2U bénéficient d'avantages similaires, prenant en charge jusqu'à 32 (ou plus) SSD E3.S contre 24 dans un système U.2. Il est intéressant de noter que certains fournisseurs de serveurs, comme Dell Technologies, ont opté pour l'E3.S comme seule option de stockage Gen5, ce qui rend la croissance de l'écosystème SSD E3.S cruciale.
Lors de nos premiers tests de performances, nous avons validé les déclarations de Micron sur le 6550 ION. Un élément crucial du lancement du 6550 ION est son profil puissance/performance. Nous avons limité le Micron 6550 ION à PS1 (ou 20 W maximum) et l'avons soumis à des charges de travail aléatoires de 16 128 et séquentielles de 5336 128 contre le Solidigm P12.7 pour voir comment cela se passait. En lecture séquentielle de 6550 7.5, nous avons mesuré 5336 Go/s sur le 16 ION contre 6550 Go/s sur le Soldigim P84.5. La mémoire NAND TLC du lecteur Micron présente également des avantages significatifs en termes de performances d'écriture, que nous avons constatés dans notre charge de travail d'écriture aléatoire de 51.9 5336. Le Micron XNUMX ION a mesuré XNUMX XNUMX IOP contre seulement XNUMX XNUMX IOP sur le PXNUMX.
Lors de ces tests, nous avons comparé le Micron 6550 ION au seul Solidigm P61.44 de 5336 To actuellement commercialisé. Ce n'est pas une comparaison équitable, mais c'est aussi le but. L'équipe d'ingénieurs de Micron a trouvé comment intégrer une quantité incroyable de NAND dans un boîtier de disque de 7.5 mm tout en mettant l'accent sur l'efficacité énergétique, les performances et la rentabilité. Ce premier aperçu des performances nous donne envie d'enregistrer plus d'heures sur ces disques pour mieux comprendre leurs performances dans diverses applications.
Ce rapport est sponsorisé par Micron. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits concernés.
