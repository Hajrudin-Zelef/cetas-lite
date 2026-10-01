---
id: collect-261001-general-networking/general-networking/fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8-1
title: "fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "datacenter", "distribution", "intel", "nand", "nvidia"]
source: docs/RAG/collect-261001-general-networking/fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8.md
source_anchor: ""
source_lines: [1, 63]
sha256: d06100bdbd97ae01a3c6aaeefdd014c217114b2eb2052e20f34049f8db3d099a
---

# fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8

Le Kioxia CD9P-R est le modèle dédié aux performances de lecture intensives de la nouvelle génération de SSD NVMe pour centres de données de Kioxia, et le premier disque de la série CD basé sur la mémoire flash BiCS TLC de 8e génération. Cette série associe le contrôleur et le firmware propriétaires de Kioxia aux technologies PCIe 5.0 et NVMe 2.0, avec des performances nominales atteignant 14 800 Mo/s en lecture séquentielle et 2.6 millions d'IOPS en lecture aléatoire, selon la capacité. Les modèles E3.S offrent des capacités allant de 1.92 To à 30.72 To, tandis que la version 2.5 pouces porte la capacité maximale à 61.44 To. Notre modèle de test est le E3.S de 7.68 To en finition SED (KCD9DPJE7T68).
Avant de consulter les graphiques, il convient de comprendre une subtilité concernant la gamme de capacités. Les spécifications publiées par Kioxia mettent en avant les performances optimales des modèles de milieu de gamme, et non celles des modèles haut de gamme : les modèles de 7.68 To et 15.36 To offrent les meilleures vitesses de lecture séquentielle (14 800 Mo/s) et d’écriture aléatoire (450 000 IOPS), tandis que le modèle phare de 30.72 To affiche respectivement 13 500 Mo/s et 270 000 IOPS. Les deux plus petites capacités sont également équipées de la mémoire NAND BiCS de génération 5, et non de la génération 8. En d’autres termes, le disque de 7.68 To que nous avons testé est celui qui révèle tout le potentiel de cette plateforme, et les acheteurs recherchant une densité maximale avec le modèle de 30.72 To devront faire des concessions, ce qui est assez courant.
C'est le passage générationnel au CD8P qui constitue l'argument principal de Kioxia, étayé par les spécifications techniques publiées. Avec une capacité de 7.68 To, le CD8P-R E3.S précédent offrait 200 000 IOPS en écriture aléatoire ; le CD9P-R porte ce chiffre à 450 000, soit une amélioration de 2.25 fois. La vitesse de lecture séquentielle progresse de 23 %, passant de 12 000 Mo/s à 14 800 Mo/s, et la vitesse de lecture aléatoire de 30 %, passant de 2 millions à 2.6 millions d'IOPS. La consommation électrique augmente légèrement, passant de 21 W (valeur typique) à 23 W pour cette capacité, ce qui place le CD9P-R au même niveau que les autres disques de la génération 5 (Gen5) à forte intensité de lecture. L'argument en faveur de l'efficacité réside ici dans les performances réelles du disque compte tenu de cette consommation.
Ce qui définit précisément le public cible de ce disque : le CD9P-R est doté d'un port unique et d'une capacité de 1 DWPD. Il ne propose pas de configuration à double port pour les baies de stockage d'entreprise traditionnelles, et les charges de travail intensives en écriture sont plutôt destinées au CD9P-V, son homologue polyvalent. Ce disque est conçu pour les parcs de serveurs hyperscale et cloud, les niveaux de lecture OLTP, la distribution de contenu et les environnements virtualisés où l'accès est principalement en lecture et la consommation énergétique est fixe. Il répond aux exigences des plateformes standard, notamment la prise en charge OCP Datacenter NVMe SSD v2.5 (toutes les exigences ne sont pas couvertes), la protection contre les coupures de courant, la protection des données de bout en bout, les options de sécurité SIE et SED, un MTTF de 2.5 millions d'heures à 50 °C et une garantie de cinq ans.
Spécifications du KIOXIA CD9P-R
Le tableau ci-dessous présente la série KIOXIA CD9P-R au format E3.S pour différentes capacités, en mettant en évidence les indicateurs de performance, les niveaux d'endurance, la puissance et les spécifications de fiabilité.
| Spécifications de la série KIOXIA CD9P-R (E3.S) |  |  |  |  |  | 
|  | 30.72TB | 15.36TB | 7.68TB | 3.84TB | 1.92TB | 
| Numéros de modèle |  |  |  |  |  | 
| Numéro de modèle SIE | KCD9XPJE30T7 | KCD9XPJE15T3 | KCD9XPJE7T68 | KCD9XPJE3T84 | KCD9XPJE1T92 | 
| Numéro de modèle SED | KCD9DPJE30T7 | KCD9DPJE15T3 | KCD9DPJE7T68 | KCD9DPJE3T84 | KCD9DPJE1T92 | 
| Spécifications de base |  |  |  |  |  | 
| Case Study | Lecture intensive (1 session d'écriture par jour) |  |  |  |  | 
| Facteur de forme | E3.S, épaisseur de 7.5 mm |  |  |  |  | 
| Interface / Protocole | PCIe 5.0 x4, NVMe 2.0 |  |  |  |  | 
| Vitesse d'interface maximale | 128 GT/s (PCIe Gen5 x4) |  |  |  |  | 
| NON | KIOXIA BiCS FLASH 3D TLC (Gen 8 pour 7.68 To à 30.72 To ; Gen 5 pour 1.92 To à 3.84 To) |  |  |  |  | 
| Conformité OCP | Spécification OCP Datacenter NVMe SSD v2.5 (partielle) |  |  |  |  | 
| Sécurité | SIE (Sanitize Instant Erase), SED (TCG Opal & Ruby SSC) |  |  |  |  | 
| Performance (jusqu'à) |  |  |  |  |  | 
| Lecture séquentielle (128 Kio, Mo/s) | 13,500 | 14,800 | 14,800 | 14,500 | 14,500 | 
| Écriture séquentielle (128 Kio, Mo/s) | 7,000 | 7,000 | 7,000 | 7,000 | 3,600 | 
| Lecture aléatoire (4 KiB, K IOPS) | 2,600 | 2,600 | 2,600 | 2,600 | 2,000 | 
| Écriture aléatoire (4 Kio, K IOPS) | 270 | 450 | 450 | 320 | 160 | 
| Exigences d'alimentation |  |  |  |  |  | 
| Tension d'alimentation | 12 V ± 10 %, 3.3 V ± 15 % |  |  |  |  | 
| Puissance (active) | 23W typ. |  |  |  |  | 
| Puissance (Prêt/Inactif) | 5W typ. |  |  |  |  | 
| Fiabilité |  |  |  |  |  | 
| MTTF | 2 500 000 heures à 0–50 °C \| 2 000 000 heures à 0–55 °C |  |  |  |  | 
| UBER | < 1 secteur pour 1017 bits lus |  |  |  |  | 
| DWPD | 1 |  |  |  |  | 
| Garantie | 5 ans |  |  |  |  | 
| Protection des données | Protection contre les coupures de courant (PLP), protection des données de bout en bout |  |  |  |  | 
| Dimensions |  |  |  |  |  | 
| Grosor | 7.5 mm +0.2 / -0.5 mm |  |  |  |  | 
| Largeur | 76 mm ± 0.25 mm |  |  |  |  | 
| longueur du câble | 112.75 mm ± 0.4 mm |  |  |  |  | 
| Poids | 110g maximum |  |  |  |  | 
| Environnemental |  |  |  |  |  | 
| Température (fonctionnement) | 0 ° C à 75 ° C |  |  |  |  | 
| Température (hors fonctionnement) | -40 ° C à 85 ° C |  |  |  |  | 
| Humidité (en fonctionnement) | 5% à 95% RH |  |  |  |  | 
| Vibrations (fonctionnement) | 21.27 m/s² { 2.17 Grms } (5–800 Hz) |  |  |  |  | 
| Choc (fonctionnement) | 9.8 km/s² { 1 000 G } (0.5 ms) |  |  |  |  | 
Performances du KIOXIA CD9P-R
Plateforme de test de conduite
Nous utilisons un serveur Dell PowerEdge R760 exécutant Ubuntu 22.04.2 LTS comme plateforme de test pour toutes les charges de travail présentées dans ce rapport. Équipé d'un boîtier JBOF Serial Cables Gen5 , il offre une large compatibilité avec les SSD U.2, E1.S, E3.S et M.2. La configuration de notre système est détaillée ci-dessous :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
- Nvidia L4
Comparaison des lecteurs
- Pascari X200P 7.68 To
- SanDisk SN861 7.68 To
- Solidigm PS1010 7.68 To
- Kingston DC3000ME 7.68 To
- Micron 7600 Max 6.4 To
- Micron 9550 MAX 12.8 To
- Micron 9550 Pro 7.68 To
Benchmark de point de contrôle DLIO
Pour évaluer les performances réelles des SSD dans les environnements d'entraînement d'IA, nous avons utilisé l'outil de test DLIO (Data and Learning Input/Output). Développé par le Laboratoire national d'Argonne, DLIO est spécifiquement conçu pour tester les modèles d'E/S dans les charges de travail d'apprentissage profond. Il permet de comprendre comment les systèmes de stockage gèrent des problématiques telles que la création de points de contrôle, l'ingestion de données et l'entraînement des modèles.
