---
id: collect-250926-servers-hardware/servers-hardware/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-1
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["compute", "gpu", "intel", "valuation"]
source: docs/RAG/clean4/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [1, 45]
sha256: dcb1be82b2621db787f98337319f8efdb013580fb6ccf6ccf8272db6ff94f737
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

Le Lenovo ThinkSystem SR630 V4 est un serveur rack 2U à 1 sockets flexible et puissant, conçu pour répondre aux besoins des secteurs tels que les services cloud et les télécommunications. Qu'il s'agisse d'optimiser les charges de travail évolutives ou de pérenniser votre centre de données, le SR630 V4 offre des mises à niveau significatives par rapport à son prédécesseur, le SR630 V3. Dans cette évaluation, nous avons examiné les nouveautés et la manière dont Lenovo a peaufiné ce serveur d'entreprise pour répondre aux défis des environnements informatiques modernes.
Différences entre Lenovo SR630 V4 et SR630 V3
Processeurs
L'une des améliorations majeures du SR630 V4 réside dans ses capacités de traitement. Alors que le SR630 V3 s'appuyait sur des processeurs Intel Xeon Scalable de 4e et 5e génération, dotés de jusqu'à 64 cœurs et de la technologie Hyper-Threading, le SR630 V4 intègre les processeurs Intel Xeon série 6700, avec jusqu'à 144 cœurs à haute efficacité énergétique (cœurs E). Cette évolution double le nombre de cœurs tout en privilégiant l'efficacité, malgré l'absence de la technologie Hyper-Threading. De plus, la version V4 prévoit la prise en charge des cœurs Intel Xeon P, ce qui pourrait considérablement améliorer les performances pour certaines charges de travail. La densité de cœurs plus élevée de la version V4 permet aux entreprises de consolider davantage d'applications sur le même nombre de serveurs, réduisant ainsi les coûts d'exploitation et les besoins en serveurs physiques.
Voici un aperçu de tous les processeurs de la série 630 pris en charge par le SR4 V6700 :
| Modèle PU | Noyaux / Threads | Vitesse du cœur (Base/To max) | L3 Cache | Mémoire Chan | Vitesse de mémoire maximale | Liens et vitesse UPI 2.0 | Lignes PCIe | TDP | 
| 6710E | 64/64 | 2.4 / 3.2 GHz | 94 MB | 8 | 5600 MHz | 4 / 16 GT/s | 88 | 205W | 
| 6731E | 96/96 | 2.2 / 3.1 GHz | 96 MB | 8 | 5600 MHz | Aucun‡ | 88 | 250W | 
| 6740E | 96/96 | 2.4 / 3.2 GHz | 96 MB | 8 | 6400 MHz | 4 / 20 GT/s | 88 | 250W | 
| 6746E | 112/112 | 2.2 / 2.7 GHz | 96 MB | 8 | 5600 MHz | 4 / 16 GT/s | 88 | 250W | 
| 6756E | 128/128 | 1.8 / 2.6 GHz | 96 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 225W | 
| 6766E | 144/144 | 1.9 / 2.7 GHz | 108 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 250W | 
| 6780E | 144/144 | 2.2 / 3 GHz | 108 MB | 8 | 6400 MHz | 4 / 24 GT/s | 88 | 330W | 
Mémoire
En ce qui concerne la mémoire, le SR630 V4 améliore la mémoire DDR3 du V5, fonctionnant jusqu'à 5600 MHz en prenant en charge des vitesses de mémoire DDR5 allant jusqu'à 6400 MHz pour les cœurs E. Les deux systèmes disposent de 32 DIMM (16 par processeur) et de deux DIMM par canal sur huit canaux par CPU. Néanmoins, le V4 introduit une protection pour l'avenir avec un support prévu pour les technologies de mémoire avancées telles que Compute Express Link (CXL) et les MCRDIMM pour les cœurs P.
Alors que le SR630 V3 pouvait prendre en charge jusqu'à 8 To de mémoire, le V4 se concentre sur une capacité plus ciblée de 2 To pour les cœurs E, optimisant ainsi les coûts et les performances pour des charges de travail spécifiques.
Stockage
Les capacités de stockage du Lenovo SR630 V4 démontrent une évolution vers des disques NVMe hautes performances tout en réduisant la dépendance aux options SAS/SATA traditionnelles. Le SR630 V3 offre une flexibilité avec des baies de disque SAS/SATA de 3.5 pouces et jusqu'à 16 ports NVMe intégrés. Cependant, le SR630 V4 prend en charge jusqu'à 12 disques NVMe dans les configurations avant et arrière, ajoutant des options futures pour les formats de disque E3.S, qui promettent des capacités et une densité supérieures.
La V4 élimine la prise en charge des disques 3.5 pouces mais introduit des options d'échange à chaud M.2 pour le démarrage du système d'exploitation, améliorant ainsi les performances et la facilité d'entretien. En se concentrant sur NVMe et en éliminant le besoin d'adaptateurs supplémentaires, la V4 maximise la bande passante d'E/S et l'efficacité du système.
Networking
Pour la mise en réseau, le SR630 V4 s'appuie sur le slot OCP unique du V3 en proposant deux slots OCP 3.0, tous deux prenant en charge PCIe Gen 5 x16. Cette mise à niveau double la flexibilité du réseau, permettant une connectivité et un débit améliorés avec des adaptateurs réseau 200 GbE à deux ports ou d'autres solutions réseau avancées.
La bande passante PCIe accrue (32 GT/s en Gen 5 contre 16 GT/s en Gen 4) signifie que le V4 peut mieux gérer les charges de travail exigeantes des centres de données et du cloud, ce qui en fait une option plus polyvalente pour les besoins de réseau modernes.
Tuning Moteur
Enfin, le SR630 V4 offre des options d'alimentation améliorées, passant des options Platinum/Titanium AC 630 W à 3 750 W du SR1800 V800 à une gamme plus large de 2000 W à 9 4 W, y compris des modèles conformes à la norme ErP Lot 48 pour une efficacité énergétique optimale. Le V1300 conserve également la prise en charge de l'alimentation -4 V CC compatible avec les opérateurs de télécommunications tout en introduisant de nouvelles options HVDC XNUMX XNUMX W pour les exigences régionales spécifiques. Ces améliorations rendent le VXNUMX plus adaptable à divers environnements électriques, garantissant qu'il répond aux exigences énergétiques de configurations plus complexes et plus performantes.
Dans l’ensemble, le SR630 V4 représente une avancée solide sur le papier pour mieux répondre au besoin de plus de performances, de flexibilité et d’efficacité dans les environnements informatiques modernes.
Spécifications du Lenovo Think System SR630 V4
| Spécifications du Lenovo Think System SR630 V4 |  | 
| Facteur de forme | rack 1U | 
| Processeur | Un ou deux processeurs Intel Xeon série 6700E (jusqu'à 144 cœurs, 2.4 GHz et TDP 330 W). Prise en charge des processeurs Intel Xeon série 6700P prévue pour le 1er trimestre 2025. | 
| Mémoire | 32 emplacements DIMM (16 par processeur), prend en charge les RDIMM TruDDR5 jusqu'à 6400 1 MHz (5200DPC) ou 2 6700 MHz (1DPC). La mémoire CXL est prévue pour la série Intel Xeon 2025P au premier trimestre XNUMX. | 
| Mémoire maximale | Jusqu'à 2 To en utilisant 32 x 64 Go RDIMM | 
| Baies de lecteur de disque |  | 
| Stockage interne maximal | 184.3 To avec 12 disques SSD NVMe 15.36 pouces 2.5 To | 
| contrôleur de stockage | Jusqu'à 16 ports NVMe intégrés avec prise en charge RAID (Intel VROC). Prise en charge prévue des adaptateurs RAID et non RAID SAS/SATA 12 Go. | 
| Interfaces réseau | Deux emplacements OCP 3.0 SFF avec interface hôte PCIe 5.0 (x8 ou x16), prenant en charge jusqu'à 100 adaptateurs réseau GbE. | 
| Emplacements d'extension PCI |  | 
| Prise en charge du GPU | Prise en charge prévue jusqu'à 3 GPU à largeur unique | 
| Ports | Avant:  Arrière:  | 
| Refroidissement | Jusqu'à 8 ventilateurs remplaçables à chaud (redondance N+1), avec un ventilateur supplémentaire intégré à chaque bloc d'alimentation. | 
| Alimentation | Jusqu'à deux blocs d'alimentation CA redondants remplaçables à chaud (800 W, 1300 2000 W, 80 XNUMX W). Certifications XNUMX PLUS Platinum et Titanium. | 
| Vidéo | Carte graphique intégrée avec deux ports vidéo (VGA arrière et Mini DisplayPort en option), prenant en charge des résolutions jusqu'à 1920 × 1200 à 60 Hz. | 
| Pièces remplaçables à chaud | Lecteurs, alimentations et ventilateurs | 
| Systems Management |  | 
| Caractéristiques de sécurité | Interrupteur d'intrusion dans le châssis, mots de passe de mise sous tension et d'administrateur, TPM 2.0 et cadre de sécurité avant verrouillable en option. | 
| Systèmes d'exploitation supportés | Serveur Microsoft Windows, Red Hat Enterprise Linux, serveur SUSE Linux Enterprise, serveur Ubuntu. | 
