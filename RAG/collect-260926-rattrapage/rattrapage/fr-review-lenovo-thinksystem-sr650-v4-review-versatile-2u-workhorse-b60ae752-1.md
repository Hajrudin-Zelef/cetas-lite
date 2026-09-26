---
id: collect-260926-rattrapage/rattrapage/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752-1
title: "Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["compute", "intel"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse-b60ae752.md
source_anchor: ""
source_lines: [1, 68]
sha256: b15e2fb48bc5c9a0f504ee0ee4effcb5c7c5af49f8932f1d31533bea4ee5cb8a
---

# Test du Lenovo ThinkSystem SR650 V4 : Un serveur 2U polyvalent et performant

*Source : https://www.storagereview.com/fr/review/lenovo-thinksystem-sr650-v4-review-versatile-2u-workhorse*

Le Lenovo ThinkSystem SR650 V4 est un serveur rack 2U à deux sockets, flexible et puissant, conçu pour répondre aux besoins de secteurs tels que les services cloud, les télécommunications et le calcul haute performance (HPC). Que ce soit pour optimiser les charges de travail à forte croissance horizontale ou pour pérenniser votre centre de données grâce à une puissance de calcul haute densité, le SR650 V4 offre des améliorations significatives par rapport à son prédécesseur, le SR650 V3.

## Différences entre Lenovo SR650 V4 et SR650 V3

### Processeurs

Les performances du processeur constituent l'une des améliorations majeures du SR650 V4. Alors que le SR650 V3 s'appuyait sur des processeurs Intel Xeon Scalable de 4e génération, le SR650 V4 inaugure la plateforme Intel Xeon 6 (anciennement nommée « Granite Rapids »). Notre test porte sur l'architecture à cœurs de performance (P-core), conçue spécifiquement pour les charges de travail exigeantes en calcul. Le V4 prend en charge un ou deux processeurs, avec jusqu'à 86 cœurs P par socket (soit jusqu'à 172 threads), offrant des fréquences plus élevées (jusqu'à 4 GHz) et une enveloppe thermique (TDP) jusqu'à 350 W.

Cette transition permet aux entreprises de gérer des applications plus exigeantes par serveur, réduisant ainsi leur encombrement physique. L'architecture offre une base solide pour la virtualisation et les opérations de bases de données lourdes. De plus, l'architecture V4 prend en charge une bande passante PCIe massive, avec jusqu'à 88 lignes PCIe 5.0 par processeur, essentielle pour alimenter les adaptateurs réseau haut débit modernes et les baies de stockage NVMe sans goulots d'étranglement.

| Modèle de CPU | Noyaux / Threads | Base Freq | Max Turbo | L3 Cache | TDP |
|---|---|---|---|---|---|
| 6787P | 86/172 | 2.0 GHz | 3.8 GHz | 336 MB | 350 W |
| 6781P | 80/160 | 2.0 GHz | 3.8 GHz | 336 MB | 350 W |
| 6767P | 64/128 | 2.4 GHz | 3.9 GHz | 336 MB | 330 W |
| 6761P | 64/128 | 2.5 GHz | 3.9 GHz | 320 MB | 330 W |
| 6760P | 64/128 | 2.2 GHz | 3.8 GHz | 320 MB | 330 W |
| 6747P | 48/96 | 2.7 GHz | 3.9 GHz | 288 MB | 330 W |
| 6745P | 32/64 | 3.1 GHz | 4.3 GHz | 336 MB | 300 W |
| 6741P | 48/96 | 2.5 GHz | 3.8 GHz | 288 MB | 300 W |
| 6740P | 48/96 | 2.1 GHz | 3.8 GHz | 288 MB | 270 W |
| 6737P | 32/64 | 2.9 GHz | 4.0 GHz | 144 MB | 270 W |
| 6736P | 36/72 | 2.0 GHz | 4.1 GHz | 144 MB | 205 W |
| 6732P | 32/64 | 3.8 GHz | 4.3 GHz | 144 MB | 350 W |
| 6731P | 32/64 | 2.5 GHz | 4.1 GHz | 144 MB | 245 W |
| 6730P | 32/64 | 2.5 GHz | 3.8 GHz | 288 MB | 250 W |
| 6724P | 16/32 | 3.6 GHz | 4.3 GHz | 72 MB | 210 W |
| 6714P | 8/16 | 4.0 GHz | 4.3 GHz | 48 MB | 165 W |
| 6530P | 32/64 | 2.3 GHz | 4.1 GHz | 144 MB | 225 W |
| 6527P | 24/48 | 3.0 GHz | 4.2 GHz | 144 MB | 255 W |
| 6521P | 24/48 | 2.6 GHz | 4.1 GHz | 144 MB | 225 W |
| 6520P | 24/48 | 2.4 GHz | 4.0 GHz | 144 MB | 210 W |
| 6517P | 16/32 | 3.2 GHz | 4.2 GHz | 72 MB | 190 W |
| 6515P | 16/32 | 2.3 GHz | 3.8 GHz | 72 MB | 150 W |
| 6511P | 16/32 | 2.3 GHz | 4.2 GHz | 72 MB | 150 W |
| 6507P | 8/16 | 3.5 GHz | 4.3 GHz | 48 MB | 150 W |
| 6505P | 12/24 | 2.2 GHz | 4.1 GHz | 48 MB | 150 W |

### Mémoire

En matière de mémoire, le SR650 V4 améliore considérablement les performances du V3, optimisant ainsi le fonctionnement du sous-système. Il prend en charge des vitesses de mémoire DDR5 jusqu'à 6 400 MHz (1 DIMM par canal), un gain important par rapport à la limite de 4 800 MHz de la génération précédente. Le système dispose de 32 emplacements DIMM (16 par processeur) répartis sur huit canaux mémoire par processeur. Il est compatible avec les modules RDIMM standard, les modules RDIMM 3DS et les nouveaux modules MRDIMM (Multiplexed Rank DIMM), qui peuvent fonctionner à des vitesses allant jusqu'à 8 000 MHz pour les applications gourmandes en bande passante. Avec des modules RDIMM 3DS de 256 Go, le serveur prend en charge une capacité de mémoire système totale impressionnante de 8 To.

De plus, la version 4 introduit la prise en charge de l'extension de mémoire CXL 2.0 (Compute Express Link). Les administrateurs peuvent installer jusqu'à 12 modules de mémoire CXL dans les baies de disques E3.S (plus précisément au format E3.S 2T). Ceci permet au système d'étendre la capacité et la bande passante de la mémoire au-delà des limites traditionnelles du processeur, réduisant ainsi la latence de calcul pour les charges de travail de nouvelle génération et diminuant le coût total de possession (TCO) en optimisant l'utilisation de la mémoire inutilisée.

### Stockage

Les capacités de stockage du Lenovo SR650 V4 témoignent d'une évolution vers une densité NVMe haute performance et des formats flexibles. Tout en prenant en charge les disques 3.5 pouces et 2.5 pouces classiques, le V4 introduit une compatibilité massive avec le format NVMe E3.S. Le système peut accueillir jusqu'à 32 disques E3.S de 1 To ou 12 disques E3.S de 2 To, offrant une densité supérieure et une gestion thermique améliorée par rapport aux SSD U.2 traditionnels.

Point essentiel, la V4 prend en charge la connectivité NVMe directe sans sursouscription (1:1), garantissant ainsi que le stockage haute vitesse ne soit pas limité par le partage des lignes PCIe. L'intégration d'options de disque de démarrage M.2 remplaçables à chaud améliore encore la facilité de maintenance.

### Networking

Pour la connectivité réseau, la SR650 V4 dispose de deux emplacements OCP 3.0 dédiés, compatibles PCIe Gen 5 x16. Cette mise à niveau majeure permet une connectivité réseau redondante et haut débit (par exemple, deux ports 200 GbE ou un port 400 GbE) sans monopoliser les emplacements d'extension PCIe standard. Le passage à la norme PCIe 5.0 double la bande passante théorique (32 GT/s contre 16 GT/s), permettant ainsi à la V4 de gérer le trafic réseau le plus exigeant pour le cloud et l'IA.

### Puissance et refroidissement

Le SR650 V4 propose des options d'alimentation améliorées, de 800 W à 3 200 W, disponibles en versions Titane et Platine. Le système prend également en charge les options -48 V CC et HVAC/HVDC pour répondre aux exigences spécifiques des centres de données.

Pour gérer la chaleur dégagée par les processeurs et la mémoire à forte consommation, Lenovo propose les solutions de refroidissement liquide Neptune. Le module « Compute Complex Neptune Core Module » utilise un système de refroidissement liquide en circuit ouvert pour évacuer la chaleur des processeurs, de la mémoire et des régulateurs de tension. Il est capable de capter plus de 80 % de la chaleur des serveurs et de réduire considérablement les coûts de refroidissement des centres de données.

Globalement, le SR650 V4 représente un véritable bond en avant en termes de performances, de densité et d'efficacité thermique pour les environnements d'entreprise modernes.

## Spécifications du Lenovo Think System SR650 V4

