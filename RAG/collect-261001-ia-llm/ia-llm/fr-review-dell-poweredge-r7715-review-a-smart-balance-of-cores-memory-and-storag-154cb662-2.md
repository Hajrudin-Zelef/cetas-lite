---
id: collect-261001-ia-llm/ia-llm/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662-2
title: "fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662"
domain: ia-llm
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "arr", "ethernet", "gpu"]
source: docs/RAG/collect-261001-ia-llm/fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662.md
source_anchor: ""
source_lines: [22, 81]
sha256: 2719b52b33c79fbf14b86ce2a85fcf24a634ceb0ad6495c3f88ff2f4fef39a8b
---

# fr-review-dell-poweredge-r7715-review-a-smart-balance-of-cores-memory-and-storag-154cb662

Le système peut être configuré pour des environnements plus exigeants avec 4 x 16 emplacements Gen5 pleine hauteur, avec une seconde carte réseau OCP incluse. Cette configuration est idéale pour les cas d'utilisation nécessitant plusieurs cartes d'extension, comme des cartes réseau doubles associées à des GPU ou d'autres accélérateurs. Imaginez que vous repoussez les limites de l'extension PCIe. Dans ce cas, le R7715 prend en charge une configuration haute densité avec 6 x 16 emplacements Gen5 pleine hauteur, optimisant ainsi la bande passante disponible et le nombre d'emplacements, ce qui le rend idéal pour les charges de travail gourmandes en E/S ou multi-accélérateurs.
Pour les utilisateurs nécessitant davantage de flexibilité en termes de largeur de voie et de type de carte, une option de carte hybride est disponible : 2 x 16 emplacements Gen5 pleine hauteur et 6 x 8 emplacements Gen5 pleine hauteur. Cette configuration est pratique pour combiner des périphériques à large bande passante avec d'autres pouvant fonctionner efficacement avec un nombre de voies PCIe plus faible. Le R7715 prend également en charge une configuration avec 2 x 16 emplacements Gen5 bas profil et 3 x 16 emplacements Gen5 double largeur pleine hauteur, conçus pour accueillir des cartes accélératrices plus volumineuses et gourmandes en énergie, avec un espacement d'air optimisé.
Enfin, Dell propose une version de cette configuration axée sur le GPU pour les environnements à refroidissement liquide, avec 2 x 16 emplacements Gen5 extra-plats et 3 x 16 emplacements Gen5 double largeur et pleine hauteur, validés pour les configurations à refroidissement liquide direct. Cela garantit la compatibilité avec les exigences de conception physique et thermique des châssis et composants à refroidissement liquide, faisant du R7715 une plateforme flexible pour les déploiements hautes performances.
Configurations de mémoire et d'alimentation du Dell PowerEdge R7715
Le PowerEdge R7715 prend en charge jusqu'à 24 emplacements DIMM DDR5, offrant ainsi une impressionnante capacité de 6 To de mémoire DDR5 ECC enregistrée. Cette configuration mémoire peut gérer jusqu'à 5200 XNUMX MT/s et est conçue pour être évolutive, offrant une bande passante élevée et une faible latence. Elle est donc idéale pour des cas d'utilisation tels que la virtualisation, la gestion de bases de données ou l'analyse de données intensive.
Le système offre également diverses options d'alimentation. Il prend en charge des alimentations redondantes remplaçables à chaud de 800 W à 3200 3200 W, offrant ainsi une flexibilité adaptée à différents besoins et configurations. Vous disposez d'unités haute performance de classe Titanium (2400 1800 W, 1500 1100 W, 800 1100 W, 800 277 W, 3200 1500 W et 1400 W), ainsi que d'options de classe Platinum (48 60 W et XNUMX W). Des configurations spéciales sont également disponibles, telles que XNUMX VCA et CCHT (XNUMX XNUMX W et XNUMX XNUMX W), ainsi qu'une alimentation CC (XNUMX XNUMX W -XNUMX--XNUMX VCC).
Dell PowerEdge R7715 et iDRAC 10
iDRAC, ou Integrated Dell Remote Access Controller, est l'outil de gestion à distance intégré de Dell qui simplifie la surveillance, la mise à jour et le dépannage des serveurs PowerEdge, tels que le R7715, sans nécessiter de présence physique. Avec la dernière version iDRAC10, Dell a apporté plusieurs améliorations pour renforcer la sécurité et la convivialité. Elle intègre désormais un processeur de sécurité dédié avec une racine de confiance intégrée, des algorithmes de chiffrement améliorés et une attestation au niveau du périphérique, rendant la gestion des serveurs plus sécurisée que jamais.
L'interface utilisateur a également été repensée pour une expérience plus homogène sur toutes les consoles Dell Technologies, avec une navigation simplifiée qui rend la gestion de votre serveur encore plus intuitive. De plus, iDRAC10 permet de créer des rôles utilisateurs personnalisés et introduit une structure de licences simplifiée, spécifiquement pour le PowerEdge 17e génération. La récupération de l'alimentation secteur est désormais gérée directement par iDRAC, au lieu d'être contrôlée par le BIOS, offrant ainsi aux administrateurs un contrôle plus centralisé, une fonctionnalité appréciable.
Spécifications Dell PowerEdge R7715
| Fonctionnalité | PowerEdge R7715 | 
| Processeur | Un processeur AMD EPYC 5 Series de 9005e génération avec jusqu'à 160 cœurs pour le processeur Zen5 | 
| Chipset | Chipset AMD | 
| Accélérateurs | Jusqu'à trois GPU double largeur de 400 W ou six GPU simple largeur de 75 W | 
| Mémoire |  | 
| Vitesse des modules DIMM | Jusqu'à 5200 MT/S | 
| Type de mémoire | RDIMM | 
| Emplacements pour module de mémoire | 24 emplacements DIMM DDR5 | 
|  | Prend en charge uniquement les modules DIMM ECC DDR5 enregistrés. | 
| Stockage |  | 
| Baies avant |  | 
| Baies arrière | N/D | 
| Contrôleurs de stockage |  | 
| Contrôleurs internes | PERC H365i, H965i, H975i | 
| Contrôleurs externes | HBA465e, H965e | 
| RAID logiciel | N/D | 
| Démarrage interne |  | 
| Source d'alimentation |  | 
| Options de refroidissement |  | 
| Ventilateurs | Jusqu'à six ventilateurs hot-plug Gold/Très hautes performances | 
| Ports |  | 
| Options réseau | Port Ethernet BMC dédié 1 Gb | 
|  | 2 cartes OCP NIC 3.0 (en option) | 
| Ports avant | 1 x USB 2.0 Type-A (LCP KVM en option) | 
|  | 1 x USB 2.0 Type-C (HÔTE/BMC Direct) | 
|  | 1 x Mini-DisplayPort (LCP KVM en option) | 
| Ports arrière | Port Ethernet BMC dédié 1 Gb | 
|  | 2 x USB 3.1 | 
|  | 1 x VGA | 
| Ports internes | 1 port USB 3.0 (en option) | 
| Slots |  | 
| PCIe | Jusqu'à huit emplacements PCIe Gen5 | 
| Facteur de forme | Serveur rack 2U | 
| Hauteur | 86.8 mm (3.41 pouces) | 
| Largeur | 482.0 mm (18.97 pouces) | 
| Profondeur | 802.4 mm (31.59 pouces) avec poignée d'alimentation | 
| Poids | Maximum 28.68 kg (63.22 lb) | 
| Biseau | Lunette métallique en option | 
| La gestion du système |  | 
| Gestion embarquée |  | 
| Console OpenManage |  | 
| Mobilité | N/D | 
| Outils | IPMI | 
| La Gestion du changement |  | 
| Intégrations OpenManage |  | 
| Sécurité |  | 
| Systèmes d'exploitation et hyperviseurs |  | 
Conception et construction du Dell PowerEdge R7715
Le PowerEdge R7715 présente un design épuré conçu pour un refroidissement efficace et des performances constantes. Son installation modulaire et sans outil simplifie la maintenance et les mises à niveau, vous permettant ainsi de maintenir un fonctionnement optimal avec un minimum de temps d'arrêt. Il est également équipé d'un cadre verrouillable en option, qui offre une sécurité accrue et confère à la face avant un aspect épuré et professionnel. Facile à retirer, il permet un accès rapide pour la maintenance ou les mises à niveau, pour une utilisation simple et pratique.
Une fois le panneau avant retiré, vous pourrez accéder à l'intégralité du panneau avant. Sur le côté droit, le panneau de commande comprend le bouton d'alimentation, un port USB (pour brancher des périphériques externes tels que des clés USB ou des périphériques pour la maintenance ou l'accès direct), un port micro iDRAC Direct et le voyant d'état iDRAC Direct.
