---
id: collect-261001-cisco/cisco/fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2-1
title: "fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2"
domain: cisco
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["arr", "attention", "ethernet", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-cisco/fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2.md
source_anchor: ""
source_lines: [1, 39]
sha256: 2b6e6921336358a3cde4632b24d53936aca0fc4e6271608c12426259b9f77921
---

# fr-review-dell-poweredge-xe7740-inside-the-architecture-of-enterprise-ai-inferen-a8c6e9c2

Le marché des infrastructures d'IA n'évolue pas dans une direction unique ; il se divise en deux univers distincts. D'un côté, les clusters d'entraînement de pointe, conçus pour développer des modèles fondamentaux à grande échelle, étroitement liés à des architectures propriétaires et à un nombre restreint d'accélérateurs. De l'autre, la réalité en pleine expansion de l'inférence en entreprise, où les organisations déploient des modèles pour servir les utilisateurs, traiter des données en temps réel et générer une valeur commerciale mesurable. Le Dell PowerEdge XE7740 est conçu spécifiquement pour ce second univers.
Points clés à retenir
- Le PowerEdge XE7740 est conçu pour l'inférence en entreprise. avec un système de refroidissement à double zone, une topologie PCIe Gen5 structurée et une architecture réseau évolutive adaptée aux charges de travail réelles de production.
- L'équilibre du système est intentionnel, Combinant la densité de 6 cœurs Xeon, une bande passante mémoire élevée et le PCIe Gen 5 E3.S NVMe pour prendre en charge le déchargement et l'orchestration du cache KV.
- La flexibilité du silicium est fondamentale, prise en charge d'une large gamme d'accélérateurs PCIe Gen5 sans nécessiter de refonte de l'infrastructure.
- La plateforme évolue facilement au fil du temps. de l'installation partielle de GPU dans un seul châssis à l'inférence distribuée sur plusieurs racks utilisant huit emplacements réseau dédiés Gen5 x16 à l'arrière.
Au cœur du XE7740 se trouve un engagement fort en faveur de la diversité des puces. Plutôt que de baser la plateforme sur une seule feuille de route d'accélérateurs, Dell a conçu un système qui s'adapte à la disponibilité, au coût et au niveau de préparation de l'organisation. Le XE7740 prend en charge une gamme d'accélérateurs PCIe Gen5, notamment les GPU NVIDIA RTX PRO 6000, H100/200, L40S, L4 et A16 pour les organisations privilégiant une large compatibilité avec leur écosystème, ainsi que l'Intel Gaudi 3 pour les équipes recherchant une solution d'inférence plus économique et immédiatement disponible. Les accélérateurs Gaudi 3 sont disponibles dès aujourd'hui, permettant aux organisations de passer de la planification au déploiement sans les délais d'approvisionnement qui caractérisent souvent les stratégies d'accélération.
L'inférence devenant la charge de travail dominante en IA, la disponibilité et le coût sont des facteurs essentiels. La plupart des entreprises n'entraînent pas de modèles à l'échelle de la pointe de la technologie. Elles exécutent des pipelines d'inférence, gèrent des modèles de langage de taille moyenne, alimentent des flux de travail de génération augmentée par la recherche et déploient la vision par ordinateur en production. Dans ce contexte, Gaudi 3 se positionne comme l'un des accélérateurs d'inférence modernes les plus abordables du marché, offrant une architecture contemporaine avec une mémoire à large bande passante et une évolutivité horizontale via Ethernet, sans le coût des GPU d'entraînement haut de gamme. Au sein du XE7740, Gaudi 3 vise moins à remplacer les solutions existantes qu'à permettre des déploiements d'inférence durables.
La plateforme entourant les accélérateurs a été conçue avec le même soin. Le XE7740 repose sur des processeurs Intel Xeon 6 et, dans les systèmes dédiés à l'inférence, le processeur demeure un composant essentiel. Le nombre élevé de cœurs et la bande passante mémoire accrue offrent la marge de manœuvre nécessaire aux planificateurs, à la tokenisation, au prétraitement et aux tâches d'orchestration qui se situent directement sur le chemin critique de l'inférence. Le stockage NVMe E3.S en façade prend en charge le stockage local des données et le déchargement du cache KV, réduisant ainsi la charge des accélérateurs et améliorant l'efficacité globale du système. Cette conception équilibrée témoigne de la conviction que les performances d'inférence dépendent de l'ensemble du système, et non des seuls accélérateurs.
Le XE7740 est conçu pour une évolutivité fluide. Les entreprises peuvent démarrer avec une configuration modeste, par exemple deux ou quatre accélérateurs, et en tirer immédiatement profit sans saturer le châssis. À mesure que les besoins augmentent, la même plateforme peut évoluer verticalement ou se transformer en inférence distribuée. Huit emplacements PCIe Gen5 x16 à l'arrière offrent une bande passante dédiée pour la mise en réseau haut débit, permettant au XE7740 de servir de base à la création de clusters d'inférence à extension horizontale. La prise en charge optionnelle des DPU renforce encore cette flexibilité en déchargeant les tâches de réseau et de communication à mesure que les déploiements évoluent.
Caractéristiques principales du Dell PowerEdge XE7740
| Spécifications | PowerEdge XE7740 | 
|---|---|
| Caractéristiques du PowerEdge XE7740 |  | 
| Processeur | Deux processeurs Intel® Xeon® série 6, avec jusqu'à 86 cœurs par processeur | 
| Slots |  | 
| Accélérateurs PCIe | 8 ports PCIe Gen 5 x16 DW-FHFL jusqu'à 600 W, ou 16 ports PCIe Gen 5 x16 SW-FHFL jusqu'à 75 W | 
| Cartes réseau PCIe |  | 
| Facteur de forme |  | 
| Facteur de forme | Serveur rack 4U | 
| Mémoire |  | 
| Vitesse des modules DIMM, capacité maximale | Jusqu'à 6 400 MT/s, 4 To max. | 
| Emplacements pour module de mémoire | 32 emplacements DIMM DDR5 Prend uniquement en charge les barrettes ECC DDR5 RDIMM enregistrées. | 
| Stockage |  | 
| Baies avant | Jusqu'à 8 x EDSFF E3.S Gen5 NVMe (SSD) max 122.88 To | 
| Contrôleurs de stockage |  | 
| Démarrage interne | Sous-système de stockage optimisé pour le démarrage (BOSS-N1 DC-MHS) : HWRAID 1, 2 x M.2 SSD NVMe | 
| Source d'alimentation |  | 
| Source d'alimentation | 3200 W Titanium 200-240 V CA ou 240 V CC, redondant remplaçable à chaud Alimentation multi-capacité pour 3200 W :  Alimentation multi-capacité pour 2400 W :  ATTENTION : Le système nécessite au moins une alimentation dans la zone CPU et une dans la zone GPU pour assurer l’alimentation du BMC et l’alimentation de secours. Si aucune alimentation n’est installée dans la zone GPU, le système restera en veille. Pour une redondance totale, installez N+N alimentations dans chaque zone : 1+1 dans la zone CPU et 3+3 dans la zone GPU. Le retrait de toutes les alimentations de la zone CPU alors que le système est sous tension provoquera un arrêt immédiat et peut entraîner une perte de données. | 
| Options de refroidissement |  | 
| Options de refroidissement | Refroidissement par air | 
| Ventilateurs | Jusqu'à quatre modules de ventilateurs haute performance (HPR) de qualité platine (module à double ventilateur) installés dans le plateau intermédiaire Jusqu'à douze ventilateurs haute performance (HPR) de qualité platine installés à l'avant du système Tous sont des ventilateurs remplaçables à chaud | 
| Ports |  | 
| Options réseau | 1 interface E/S compatible PCIe Gen 5 OCP 3.0 (prise en charge par 8 lignes PCIe) | 
| Ports avant | 1 port USB 2.0 Type-A (en option) 1 port Mini-DisplayPort (en option) 1 port USB 2.0 Type-C double mode (port hôte/iDRAC direct) | 
| Ports arrière | 1 port Ethernet direct dédié iDRAC/BMC 2 ports USB 3.1 Type A 1 x VGA | 
| Ports internes | 1 x USB 3.1 Type-A | 
Conception et fabrication du XE7740
Architecture à double zone : séparation du CPU et du GPU
