---
id: collect-261001-cisco/cisco/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a-1
title: "fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "ethernet", "intel"]
source: docs/RAG/collect-261001-cisco/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a.md
source_anchor: ""
source_lines: [1, 47]
sha256: 8eb6f8b071bb8a4262cbfc2475f2e1f7b7ad9f9469fdc13f9560a3823129a880
---

# fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a

La famille EPYC 4005 d'AMD intègre Zen 5 à l'AM5 avec un objectif opérationnel clair. Maintenir la cohérence de la plateforme pour les constructeurs et les hébergeurs de systèmes, puis optimiser les éléments essentiels à l'utilisation quotidienne. Le résultat est une approche de l'AM5 axée sur le serveur, qui privilégie des déploiements prévisibles, un firmware stable et une matrice de système d'exploitation conçue pour l'évolutivité.
Dans la gamme AMD, l'EPYC 4005 se situe entre Ryzen sur AM5 et les plateformes EPYC haut de gamme sur SP5 et SP6. Ryzen peut ancrer des rôles de serveur léger, mais il n'offre pas la profondeur de validation, le cycle de vie ni l'ensemble de fonctionnalités attendus par les opérateurs de centres de données. Les gros sockets offrent davantage de cœurs, de canaux mémoire et de voies, mais un coût bien plus élevé. L'EPYC 4005 comble cet écart. Vous conservez le côté économique et la simplicité de l'AM5 tout en bénéficiant d'un cadre d'entreprise compatible avec les licences et les enveloppes de puissance courantes.
La gamme de produits reflète cette intention. La couverture s'étend dans les deux sens, permettant aux acheteurs de standardiser avec un budget énergétique modeste ou d'opter pour un niveau de cache supérieur important pour les services sensibles à la latence, le tout sans changer de châssis ni de forfait d'E/S.
AMD EPYC 4004 vs 4005 : les améliorations pratiques
L'EPYC 4005 est une évolution générationnelle simple par rapport à la 4004, qui conserve l'héritage AM5 tout en améliorant la plateforme. Avec la 4005, vous bénéficiez de cœurs Zen 5, d'une augmentation des performances de la RAM jusqu'à la DDR5-5600 et de la même capacité PCIe Gen28 jusqu'à 5 voies. AMD ajoute également l'AVX-512 avec un chemin de données complet de 512 bits, ce qui offre à cette catégorie de serveurs monosocket une marge de manœuvre accrue pour les tâches vectorielles intensives sans changer de carte mère ni de châssis.
Au sommet de la pile, la comparaison est simple. L'EPYC 4565P du 4005 est un processeur 16 cœurs avec 64 Mo de mémoire L3, un TDP de 170 W et une fréquence d'horloge allant jusqu'à 5.7 GHz. Son homologue du 4004, l'EPYC 4564P, est également un processeur 16 cœurs avec 64 Mo de mémoire L3 et un TDP de 170 W, avec la même augmentation maximale de 5.7 GHz. Cela signifie que les gains générationnels proviennent des mises à jour du cœur et de la plateforme plutôt que d'un saut significatif des spécifications.
Si vous souhaitez travailler avec du cache, les deux familles proposent une option V-Cache 3D. Cependant, le 4005 place le processeur EPYC 4585PX au sommet avec 170 W et 128 Mo de mémoire L3, tandis que le 4004 plafonne avec l'EPYC 4584PX à 120 W et 128 Mo de mémoire. Les fréquences de base du 4585PX sont également légèrement supérieures.
L'ajout pratique le plus significatif du 4005 est une référence 16 cœurs et 65 W. L'EPYC 4545P permet aux serveurs hébergés bare metal et PME de standardiser 16 cœurs pour les licences Windows tout en respectant des enveloppes de puissance et de refroidissement plus strictes. Cette option n'existait pas dans le 4004, où 16 cœurs signifiaient 170 W sur la puce standard ou un composant V-Cache de 120 W avec des fréquences de base plus basses.
Comme indiqué précédemment, la prise en charge de la mémoire a été légèrement améliorée. Le 4004 a validé deux canaux de DDR5 jusqu'à 5200 4005 MT/s. Le 5600 porte cette vitesse à 192 128 MT/s et augmente la capacité maximale prise en charge au niveau de la plateforme de XNUMX Go à XNUMX Go, ce qui facilite les charges virtualisées mixtes et les petites instances de bases de données.
Comparaison des processeurs EPYC 4004 et 4005
| Série AMD EPYC 4004 Modèles | Arch. | Noyaux / Threads | Cache L3 (Mo) | TDP (W) | Horloge de base (GHz) | Boost horloge (GHz) | Prix (1KU, USD) | 
|---|---|---|---|---|---|---|---|
| EPYC 4124P | Zen 4 | 4/8 | 16 | 65W | 3.8 | 5.1 | $149 | 
| EPYC 4244P | Zen 4 | 6/12 | 32 | 65W | 3.8 | 5.1 | $229 | 
| EPYC 4344P | Zen 4 | 8/16 | 32 | 65W | 3.8 | 5.3 | $329 | 
| EPYC 4364P | Zen 4 | 8/16 | 32 | 105W | 4.5 | 5.4 | $399 | 
| EPYC 4464P | Zen 4 | 12/24 | 64 | 65W | 3.7 | 5.4 | $429 | 
| EPYC4484PX | Zen 4 | 12/24 | 128 | 120W | 4.4 | 5.6 | $599 | 
| EPYC 4564P | Zen 4 | 16/32 | 64 | 170W | 4.5 | 5.7 | $699 | 
| EPYC4584PX | Zen 4 | 16/32 | 128 | 120W | 4.2 | 5.7 | $699 | 
| Modèles de la série AMD EPYC 4005 |  |  |  |  |  |  |  | 
| EPYC 4245P | Zen 5 | 6/12 | 32 | 65W | 3.9 | 5.4 | $239 | 
| EPYC 4345P | Zen 5 | 8/16 | 32 | 65W | 3.8 | 5.5 | $329 | 
| EPYC 4465P | Zen 5 | 12/24 | 64 | 65W | 3.4 | 5.4 | $399 | 
| EPYC 4545P | Zen 5 | 16/32 | 64 | 65W | 3.0 | 5.4 | $549 | 
| EPYC 4565P | Zen 5 | 16/32 | 64 | 170W | 4.3 | 5.7 | $589 | 
| EPYC4585PX | Zen 5 | 16/32 | 128 | 170W | 4.3 | 5.7 | $699 | 
Plateforme de test : MSI S1102-02 avec refroidissement liquide
Pour ce test, nous avons utilisé la carte mère MSI S1102-02, un barebone 1U AM5 parfaitement adapté au marché cible des processeurs EPYC 4005. Elle prend en charge les processeurs AM5 jusqu'à 170 W et notre modèle est livré avec un système de refroidissement liquide tout-en-un en option, une caractéristique rare dans cette catégorie et idéale pour les processeurs 4005 haut de gamme. Ce système devrait notamment permettre de réduire la vitesse et le niveau sonore des ventilateurs. Côté refroidissement, le boîtier intègre sept ventilateurs système de 40 mm, une alimentation Platinum 1+1 de 600 W et une gestion complète via un contrôleur BMC AST2600 compatible IPMI et Redfish.
La mise en réseau s'effectue via deux ports RJ10 45 GbE sur le processeur Intel X710. Le stockage est simple avec quatre baies SATA remplaçables à chaud en façade, deux baies SATA internes 2.5 pouces et deux emplacements NVMe M.2 2280/22110. L'extension comprend un emplacement Gen5 x16 pour cartes FHHL et un chemin Gen4 x4 partagé pouvant être acheminé vers un emplacement PCIe secondaire ou vers l'un des sockets M.2. La mémoire est composée de quatre emplacements DDR5 UDIMM avec prise en charge ECC, validés jusqu'à 5600 1 MT/s à 48 DPC et 192 Go par DIMM pour une capacité maximale de 26 Go. La profondeur physique est de XNUMX pouces, ce qui offre une grande flexibilité de déploiement dans les racks courts.
Spécifications complètes du MSI S1102-02
| Spécifications | DÉTAILS | 
| Facteur de forme | 1U | 
| Dimensions | 438.5 mm (17.26″) L x 43.5 mm (1.71″) H x 660 mm (26″) P | 
| Processeur | Processeur AMD Ryzen™ série 7000/9000 et EPYC™ série 4004/4005, jusqu'à TDP 170 W | 
| Douille | (1) Prise AMD AM5 | 
| Chipset | AMD B650 | 
| Mémoire | (4) emplacements DIMM DDR5, 2DPC, UDIMM ECC/non-ECC – Max. Fréquence 5600MT/s(1DPC) et 3600MT/s(2DPC) – Max. Capacité par DIMM : 48 Go | 
| Baies de disques | (4) Les baies de disque 3.5"/2.5" remplaçables à chaud prennent en charge SATA 3.0 | 
| Stockage interne | (2) ports M.2 2280/22110 PCIe4.0 x4 (2) Les baies de disque internes 2.5″ prennent en charge SATA 2.0 | 
| Slots d'extension | (1) L'emplacement PCIe 4.0 x16 du processeur prend en charge la carte PCIe FHHL | 
| Networking | (2) ports Ethernet 10GBase-T (Intel® X710AT2) *JLAN1,2 prend en charge NCSI | 
| RAID | N/D | 
| Front I / O | (4) Baies de lecteur 3.5" remplaçables à chaud (2) Ports USB3.2 Gen1 Type-A (1) Bouton d'alimentation du système (1) Bouton UID (1) LED UID (5) LED d'état : Alimentation/Défaut/HDD/(2)LAN | 
| I / O arrière | (2) ports Ethernet 10GBase-T (1) Port de gestion de serveur dédié 1000Base-T (4) Ports USB3.2 Gen1 Type-A (1) port VGA D-Sub (1) Port COM DB9 (1) Bouton LED UID | 
| TPM | (1) En-tête TPM avec interface SPI | 
| Sécurité | TPM 2.0 | 
