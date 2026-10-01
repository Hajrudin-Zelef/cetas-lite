---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-1
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["attention", "compute", "datacenter", "distribution", "ethernet", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [1, 32]
sha256: b282c3e5b514521f86ae8abedbb35b5f0ee6287cfc2c8841566a314ff5937a1c
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

Les serveurs Dell PowerEdge série R7x0 sont depuis longtemps un pilier des centres de données, réputés pour leur qualité de fabrication exceptionnelle, leur conception soignée, leurs performances, leur densité et leur fiabilité, le tout dans un format 2U polyvalent. Ces serveurs ont constamment évolué pour répondre aux exigences changeantes. Aujourd'hui, avec l'introduction du Dell PowerEdge R770, la série franchit une étape décisive.
Le R770 inaugure la nouvelle famille de processeurs Intel Xeon 6, avec les processeurs à cœurs Xeon 6500 et 6700 (architectures P et E). Il marque la première adoption complète par Dell de la norme OCP Data Center Modular Hardware System (DC MHS) dans sa gamme de serveurs grand public. Ensemble, ces deux évolutions promettent une avancée significative en termes de performances et de conception.
Répondre aux exigences des centres de données modernes
Le lancement du R770 intervient alors que les centres de données sont confrontés à une pression croissante. Les charges de travail sont de plus en plus diversifiées et exigeantes. La croissance incessante des données renforce le besoin d'analyses et de bases de données robustes. De l'entraînement de modèles complexes au déploiement d'inférences en temps réel, l'intelligence artificielle n'est plus une application de niche, mais un moteur métier essentiel nécessitant une puissance de calcul importante et une accélération spécialisée.
Parallèlement, l'efficacité énergétique et l'optimisation du coût total de possession font l'objet d'une attention particulière. De plus, l'industrie se tourne de plus en plus vers les normes ouvertes pour favoriser l'innovation, améliorer l'interopérabilité et potentiellement réduire la dépendance vis-à-vis d'un fournisseur. Le R770, avec ses nouvelles options de processeur et l'adoption de l'OCP DC MHS, est conçu pour relever ces défis.
Processeurs Intel Xeon 6 P-Core
Le processeur R770 utilise les processeurs Intel Xeon série 6, notamment les séries 6700 et 6500, intégrant les cœurs Performance et Efficiency basés sur la plateforme Socket E2 (LGA4710-2). Dans cette analyse, nous nous concentrons spécifiquement sur les références de la série P.
Intel construit ces processeurs selon une conception en tuiles, combinant des tuiles d'E/S avec une ou deux tuiles de calcul. Cela permet une évolutivité au sein de la série, avec des configurations allant jusqu'à 86 cœurs P (XCC) avec deux tuiles de calcul, et jusqu'à 48 cœurs P (HCC) ou 16 cœurs P (LCC) avec une seule tuile de calcul.
Comparés aux processeurs Sapphire et Emerald Rapids de génération précédente, ces processeurs se distinguent par la disponibilité universelle d'accélérateurs intégrés sur tous les processeurs Xeon 6. Parmi ceux-ci figurent la technologie Intel QuickAssist pour le chiffrement et la compression, l'accélérateur de streaming de données Intel pour le transfert de données, l'accélérateur d'analyse en mémoire Intel pour l'accélération des bases de données et de l'analyse, et l'équilibreur de charge dynamique Intel pour l'efficacité du traitement réseau.
La mémoire et la bande passante d'E/S bénéficient également d'améliorations substantielles. Les processeurs Xeon 6700/6500 P-core prennent en charge la mémoire DDR8 à 5 canaux. Ils ouvrent également la voie aux modules MRDIMM (Multiplexed Rank DIMM), qui offrent des vitesses allant jusqu'à 8,800 5.0 MT/s. Côté E/S, ces processeurs prennent en charge les normes PCIe 2.0 et CXL 88. En configuration double socket, la plateforme peut offrir jusqu'à 176 voies PCIe par socket (soit XNUMX voies au total).
Malgré la différenciation entre les processeurs P-core et E-core, la famille Xeon 6 conserve une cohérence dans les jeux d'instructions, le BIOS, les pilotes, la prise en charge des systèmes d'exploitation et des applications, ainsi que les fonctionnalités RAS, simplifiant ainsi l'intégration et la gestion entre différents types de déploiement. Les processeurs P-core sont destinés aux charges de travail où les performances par cœur, l'accélération de l'IA, une bande passante mémoire élevée et des E/S importantes sont primordiales ; pensez aux bases de données exigeantes, aux simulations HPC, à l'analyse avancée et à un large éventail d'applications d'IA.
Spécifications Dell PowerEdge R770
| Spécifications | Dell PowerEdge R770 | 
| Processeur | Deux processeurs Intel Xeon 6 avec jusqu'à 144 cœurs E ou 86 cœurs P par processeur | 
| Mémoire | 32 emplacements DIMM DDR5, prend en charge RDIMM 8 To max, vitesses jusqu'à 6400 5 MT/s, prend en charge uniquement les DIMM DDRXNUMX ECC enregistrés | 
| Contrôleurs de stockage | Démarrage interne : Sous-système de stockage optimisé pour le démarrage (BOSS-N1 DC-MHS) : HWRAID 1, 2 x SSD M.2 NVMe ou carte intercalaire M.2 (DC-MHS) : 2 x SSD M.2 NVMe ou USB, Contrôleurs internes : PERC H965i avant, PERC H975i avant, PERC H365i avant | 
| Baies avant et arrière |  | 
| Blocs d'alimentation remplaçables à chaud |  | 
| Options de refroidissement | Refroidissement par air et refroidissement liquide direct (DLC est une solution de rack et nécessite des collecteurs de rack et une unité de distribution de refroidissement (CDU) pour fonctionner) | 
| Ventilateurs | Ventilateurs Silver hautes performances (HPR SLVR)/Ventilateurs Gold hautes performances (HPR GOLD), jusqu'à 6 ventilateurs remplaçables à chaud | 
| Dimensions et poids | Hauteur – 86.8 mm (3.42 pouces), largeur – 482 mm (18.97 pouces), poids – 28.53 kg (62.89 livres), profondeur (pour la configuration E/S arrière) – 802.40 mm (31.59 pouces) avec cadre, 801.51 mm (31.56 pouces) sans cadre, profondeur (pour la configuration E/S avant) – 814.52 mm (32.07 pouces) sans cadre | 
| Facteur de forme | Serveur rack 2U | 
| Gestion intégrée | iDRAC, iDRAC Direct, API RESTful iDRAC avec Redfish, CLI RACADM, module de service iDRAC (iSM), point de terminaison NativeEdge, orchestrateur NativeEdge | 
| Biseau | Lunette de sécurité en option | 
| Sécurité | Micrologiciel signé cryptographiquement, chiffrement des données au repos (SED avec gestion des clés locale ou externe), démarrage sécurisé, vérification des composants sécurisés (contrôle de l'intégrité du matériel), racine de confiance en silicium, verrouillage du système, verrouillage du système (nécessite iDRAC10 Enterprise ou Datacenter), détection d'intrusion dans le châssis, TPM 2.0 FIPS, certifié CC-TCG | 
| Options réseau |  | 
| Options GPU | Jusqu'à 6 x 75 W FHHL ou jusqu'à 2 x 350 W DWFL | 
| Ports | Ports avant : 1 port USB 2.0 Type C, 1 port USB 2.0 Type A (en option), 1 port Mini-DisplayPort (en option), 1 port série DB9 (avec configuration E/S avant), 1 port Ethernet dédié à la gestion iDRAC ; Ports arrière : 1 port Ethernet dédié à la gestion iDRAC, 1 port VGA, 2 ports USB 3.1 Type A ; Ports internes : 1 port USB 3.1 Type A | 
| PCIe |  | 
| Systèmes d'exploitation et hyperviseurs | Serveur Ubuntu LTS canonique, serveur Microsoft Windows avec Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware avec vSphere | 
Dell PowerEdge R770 : la modularité avec OCP DC MHS
Le Dell PowerEdge R770 présente des avancées notables et une flexibilité dans sa conception physique et son architecture de composants, en adoptant la norme Data Center Modular Hardware System (OCP DC MHS) du projet Open Compute.
