---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594-2
title: "fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["compute", "gpu", "intel", "nvidia"]
source: docs/RAG/clean4/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594.md
source_anchor: ""
source_lines: [39, 78]
sha256: 35a668c3fe7a4163c51356295bfd7d3440545a90dd8e2915b55cf1fda3ba2b20
---

# fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594

- Intel Xeon-Bronze de 4e génération
  - Processeur Intel Xeon-Bronze 3408U 1.8 GHz à 8 cœurs 125 W
HPE ProLiant DL320 Gen11 – Variantes de GPU/stockage
Il existe plusieurs configurations physiques différentes du DL320 Gen11, offrant aux clients la flexibilité d'ajuster leurs options de stockage et d'extension en fonction de leur charge de travail.
Notre modèle de révision standard prend en charge 8 +2 disques 2.5″ avec quelques options de fond de panier différentes. HPE propose également le serveur avec une configuration de 4 disques durs de 3.5 pouces pour les cas d'utilisation qui nécessitent plus de capacité que de performances. Ces deux offres sont assez standards dans cette catégorie. Mais HPE devient alors un peu sauvage.
Ils proposent une version 12x HDD du DL320 Gen11, quelque chose que l'on ne voit généralement pas sur un serveur d'entrée de gamme. La cage de disque 12x LFF est disposée selon une disposition de 3 × 4 rangées. Cette cage de disque importante offre une augmentation significative de la capacité de stockage adaptée aux applications gourmandes en données. Cette version du serveur est un peu plus longue pour s'adapter à la profondeur nécessaire pour loger les disques durs.
HPE propose également des configurations de serveur GPU avec le DL320 Gen11. HPE prend en charge jusqu'à quatre GPU simples ou deux doubles largeurs dans la cage avant avec le serveur GPU CTO. Cela signifie que les clients peuvent configurer le serveur avec quatre GPU NVIDIA L4 ou deux GPU NVIDIA L40, par exemple, pour disposer d'une machine d'inférence très rentable pour les cas d'utilisation Edge.
De plus, une option de configuration pour la variante GPU du DL320 Gen11 permet d'installer huit SSD E3.S, remplaçant le châssis de stockage 4x U.3.
Gestion HPE ProLiant DL320 Gen11 – iLO 6
Le DL320 Gen11 dispose d'une suite de fonctionnalités de gestion avancées grâce à l'intégration d'iLO 6, la dernière version de HPE de sa technologie Lights-Out. iLO 6 apporte une série d'avantages (tels que la configuration de serveur à distance, la surveillance de l'état et le contrôle de l'alimentation et de la température) qui améliorent la gérabilité, la sécurité et l'efficacité du DL320 et simplifient les environnements informatiques complexes. Bien qu'iLO ne soit pas nouveau et soit couramment présent sur la plupart des serveurs HPE, nous le notons ici car il constitue une valeur ajoutée considérable pour ceux qui ne sont peut-être pas habitués à une gestion de serveur aussi approfondie dans cette fourchette de prix particulière.
Les fonctionnalités de sécurité avancées d'iLO sont conçues pour protéger contre les menaces et garantir l'intégrité des données tout en offrant une interface simplifiée et des capacités d'automatisation améliorées pour faciliter une expérience utilisateur plus intuitive. En fin de compte, cela réduira le temps et les efforts requis pour diverses tâches de gestion et vous facilitera la vie en tant qu'administrateur.
Pour améliorer la sécurité de votre DL320, iLO 6 intègre la racine de confiance en silicium exclusive de HPE, garantissant que seul le micrologiciel signé par HPE démarre pour établir une base sécurisée dès le départ. Cette sécurité est encore renforcée par HPE Secure Start, qui authentifie le micrologiciel via cette racine de confiance en silicium pour garantir que le micrologiciel démarré est sûr et sécurisé. HPE iLO 6 impose également l'utilisation de cryptographie et d'algorithmes conformes aux normes CNSA (Commercial National Security Algorithm) et introduit la prise en charge de l'authentification à 2 facteurs via des cartes PIV/CAC.
La vérification du micrologiciel d'exécution d'iLO 6 vérifie également régulièrement les micrologiciels essentiels pour détecter les intrusions après le démarrage. S'il détecte une violation, il permet une restauration rapide du micrologiciel à l'usine ou aux derniers paramètres de sécurité connus pour aider à minimiser les dommages potentiels causés par l'intrus.
HPE ProLiant DL320 Gen11 Spécifications complètes
| Spécifications | DÉTAILS | 
| Type de processeur | Intel | 
| Famille de processeurs | Processeurs évolutifs Intel Xeon de 4e génération | 
| Modèles de processeur |  | 
| Nombre de processeur | 1 | 
| Processeur Core disponible | 8 à 32 cœurs, selon le processeur | 
| Cache du processeur | 26.25 – 60 Mo L3, selon le processeur | 
| La vitesse du processeur | Jusqu'à 3.7 GHz, selon le processeur | 
| Type d'alimentation | Alimentations HPE Flex Slot (les options varient : 500 W, 800 W, 1000 1600 W, 1800 2000 W, XNUMX XNUMX-XNUMX XNUMX W) | 
| Slots d'extension | Maximum, 2 PCIe Gen5 et 1 OCP 3.0 PCIe Gen5 | 
| Mémoire maximale | 2.0 To par socket, un seul socket, lorsqu'il est équipé de 128 Go de mémoire DDR5. | 
| Slots mémoire | 16 emplacements DIMM par socket | 
| Type de mémoire | Mémoire intelligente HPE DDR5 | 
| Fonctions de protection de la mémoire | Mémoire HPE Fast Fault Tolerant Advanced ECC, rechange en ligne, mémoire miroir | 
| Type de lecteur optique | Lecteur optique DVD-RW SATA HPE 9.5 mm en option, lecteur DVD-RW USB mobile HPE | 
| Caractéristiques du ventilateur du système | Kit de ventilateur standard ou haute performance, selon le modèle | 
| contrôleur réseau | Large gamme d'options (adaptateur vertical PCIe et OCP3.0) | 
| contrôleur de stockage | Contrôleur SATA intégré (RAID logiciel AHCI ou Intel SATA), en option – Contrôleur HPE Smart Array Gen11 | 
| Capacité DIMM | 16 Go à 256 Go | 
| Gestion de l'infrastructure | HPE GreenLake pour Compute Ops Management, HPE iLO Standard, HPE OneView Standard, en option – HPE iLO Advanced, HPE OneView Advanced | 
| Garantie | 3/3/3 : Trois ans de couverture pour les pièces, la main d’œuvre et l’assistance sur site. Une couverture supplémentaire est disponible. | 
| Lecteur pris en charge | Jusqu'à 8+2 disques durs SAS/SATA SFF ou SSD SATA/SAS/NVMe U.2 ou U.3, selon le modèle | 
HPE ProLiant DL320 Gen11 contre DL360 Gen11
Avec plusieurs serveurs Intel 1U dans le portefeuille, la comparaison des configurations disponibles des serveurs HPE ProLiant DL320 Gen11 et DL360 Gen11 révèle des différences distinctes adaptées aux différents cas d'utilisation et besoins de performances. La différence la plus évidente réside dans le choix du processeur. Le DL320 Gen11 propose des processeurs à 32 cœurs, et le DL360 Gen 11 propose deux processeurs et s'adapte au meilleur processeur d'Intel.
Le DL320 Gen11 offre des configurations de mémoire de base inférieures : 2 To par socket (un seul socket) lorsque chaque 16 DIMMS est équipé de 128 Go de clés mémoire DDR5. Cela met en évidence son adaptabilité aux scénarios à la fois économiques et plus exigeants. Le DL360 Gen11 double ce total de mémoire à 4 To par socket lorsqu’il est équipé de clés RAM DDR256 de 5 Go.
Les configurations PCIe du DL320 Gen11, généralement un emplacement Gen5 et un emplacement OCP 3.0 Gen5, s'adressent aux entreprises ayant besoin d'une extensibilité modérée, peut-être pour des cartes réseau supplémentaires ou un stockage externe raisonnable. Le DL360 Gen11 va plus loin en proposant davantage de slots d'extension. Cela indique une conception adaptée aux environnements où les E/S à haut débit et une connectivité externe étendue sont cruciales, à l'instar de l'informatique en cluster ou des réseaux de stockage à haut débit.
