---
id: collect-261001-general-networking/general-networking/fr-review-qnap-ts-h1887xu-rp-nas-review-07398158-1
title: "fr-review-qnap-ts-h1887xu-rp-nas-review-07398158"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Intel"]
dates: []
keywords: ["dram", "ethernet", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-qnap-ts-h1887xu-rp-nas-review-07398158.md
source_anchor: ""
source_lines: [1, 50]
sha256: 8f5c12450a1926b780d5bf10b9c4dc59684d75683b473d88be05a237a6cf1e31
---

# fr-review-qnap-ts-h1887xu-rp-nas-review-07398158

QNAP a récemment lancé sa série TS-hx87XU-RP, une gamme de périphériques NAS disponibles en modèles à 9, 18, 22 et 30 baies. La famille de NAS intègre les processeurs de la série Intel Xeon E-2300 avec une connectivité Multi-Gig 10GbE et 2.5GbE et une architecture de stockage hybride HDD/SSD qui équilibre les performances, la capacité et le coût. À l'étude aujourd'hui, c'est le modèle à 18 baies. Le QNAP TS-h1887XU-RP comprend 12 baies HDD 3.5″ à l'avant et 6 baies SATA 2.5″ à l'arrière, dans un châssis 2U.
Présentation du QNAP TS-h1887XU-RP
Ces systèmes peuvent exécuter les systèmes d'exploitation QTS ou QuTS hero. Exécutant le système d'exploitation QuTS hero basé sur ZFS, la série TS-hx87XU-RP donne la priorité à l'intégrité des données et fournit une réduction puissante des données, des instantanés presque illimités et jusqu'à 5 pétaoctets de capacité par dossier partagé. Cela en fait une solution de stockage idéale pour les entreprises, avec des applications dans les serveurs de fichiers, les serveurs de virtualisation, le VDI et la sauvegarde/restauration. Le système d'exploitation QTS offre une amélioration des performances si les utilisateurs n'ont pas besoin de toutes les fonctionnalités de données que ZFS apporte à la table.
Comme mentionné, le TS-h1887XU-RP offre 18 baies au total pour le stockage, avec 12 baies HDD à l'avant et 6 baies SSD à l'arrière, elles sont toutes SATA. QNAP prend en charge le stockage supplémentaire via des cartes d'extension en cas de besoin d'un flash plus rapide via l'adaptateur QM2 PCIe en option. L'option d'extension sur PCIe offre une grande flexibilité dans cette version, QNA comprend trois emplacements disponibles, 2 x Gen4 x4 et 1 x Gen 4 x8.
Pour la connectivité, il y a 2 ports 2.5 GbE à bord ainsi que 2 ports 10GBASE-T (10G/5G/2,5G/1G/100M). En ce qui concerne la DRAM, quatre emplacements sont disponibles et ce système prendra en charge 128 Go au total.
Il existe en réalité deux sous-configurations du NAS 18 baies. Le TS-h1887XU-RP-E2336-32G est équipé d'un processeur Intel Xeon E-2336 (6 cœurs/12 threads) cadencé à 2.9 GHz (jusqu'à 4.8 GHz), de 32 Go de mémoire ECC DDR4 (2 x 16 Go) et d'une alimentation redondante. Le TS-h1887XU-RP-E2334-16G, quant à lui, propose un processeur Intel Xeon E-2334 (4 cœurs/8 threads) cadencé à 3.4 GHz (jusqu'à 4.8 GHz), 16 Go de mémoire ECC DDR4 (1 x 16 Go) et une alimentation redondante. Notre test porte sur cette dernière configuration. Au moment de la publication, son prix public conseillé avoisine les 4 200 $.
Spécifications du QNAP TS-h1887XU-RP
| Processeur | Intel Xeon E-2334 4C 8T 3.4 GHz, jusqu'à 4.8 GHz | 
|---|---|
| Architecture du processeur | 64 bits x86 | 
| Unité à virgule flottante | Oui | 
| Moteur de chiffrement | Oui (AES-NI) | 
| Mémoire système | 16 Go ECC DDR4 (1 x 16 Go) | 
| Mémoire maximale | 128 Go (4 x 32 Go) | 
| slot mémoire | 4 x ECC UDIMM DDR4 | 
| Mémoire flash | 5 Go (protection du système d'exploitation à double démarrage) | 
| Drive Bay | 18 (12 SATA 3.5 pouces + 6 SATA 2.5 pouces) | 
| Fente M.2 | En option via un adaptateur PCIe QM2 | 
| Prise en charge de l'accélération du cache SSD | Oui | 
| Pass-through GPU | Oui | 
| Port Ethernet 2.5 gigabits (2.5 G/1 G/100 M) | 2 (2.5G/1G/100M) | 
| 10 port Ethernet Gigabit | 2 x 10GBASE-T (10G/5G/2,5G/1G/100M) | 
| Emplacement PCIe | 3 (2 x Gen4 x4 +1 x Génération 4 x8) | 
| Facteur de forme | Montage en rack 2U | 
| Bloc D'Alimentation | 2 blocs d'alimentation 550 W | 
| Ventilateur | 4 x 60 mm, 12 V CC | 
Installation et configuration du QNAP TS-h1887XU-RP
Pour ce test, nous utilisons 12 disques durs WD Red Pro de 22 To pour le stockage principal et 6 SSD WD Red de 4 To pour les volumes flash. Chaque groupe de disques a été utilisé pour créer son propre pool de stockage RAID 6, à partir duquel 4 LUN iSCSI et 4 partages de fichiers ont été provisionnés. La compression et la déduplication ne sont pas activées et la taille des blocs est de 32 Ko.
Les disques durs WD Red Pro sont une gamme de disques haute capacité conçus pour les systèmes NAS jusqu'à 24 baies. Optimisés pour les environnements multi-utilisateurs, ils supportent des charges de travail intensives en fonctionnement continu (24h/24 et 7j/7). Idéaux pour stocker, protéger, archiver et partager de grandes quantités de données, ils conviennent à diverses applications gourmandes en données. Dans le cadre de nos tests, ces disques sont parfaitement adaptés au stockage de masse. Ils sont disponibles sur Amazon au prix d'environ 550 $.
Le SSD WD Red SA500 est conçu pour servir de solution de cache dans les systèmes NAS, offrant des performances accrues et un accès plus rapide aux données. Ici, nous l'utilisons cependant comme volume flash. Quoi qu'il en soit, ces SSD sont conçus pour minimiser les goulots d'étranglement et améliorer l'efficacité globale des systèmes NAS. Grâce à leur interface SATA, ils sont également économiques ; leur prix public conseillé est d'environ 380 $.
Performances du QNAP TS-h1887XU-RP
Pour cet examen, nous utilisons 12 x 22 To WD Red Pro en RAID6 et 6 x 4 To WD Red SSD en RAID6.
Notre processus de test NAS d'entreprise préconditionne chaque ensemble de disques dans un état stable avec la même charge de travail avec laquelle l'appareil sera testé sous une lourde charge de 16 threads, avec une file d'attente exceptionnelle de 16 par thread. L'appareil est ensuite testé à des intervalles définis dans plusieurs profils de profondeur de thread/file d'attente pour montrer les performances en cas d'utilisation légère et intensive. Étant donné que les disques durs atteignent très rapidement leur niveau de performance nominal, nous ne représentons graphiquement que les principales sections de chaque test.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend trois profils basés sur des tailles de charge de travail courantes. Ces profils ont été développés pour faciliter la comparaison avec nos références passées, ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4K et 8K 70/30, qui est couramment utilisée pour les périphériques de stockage. Lors des tests de plates-formes NAS, pour chaque pool de stockage que nous comparons, nous provisionnons quatre dossiers partagés pour les fichiers de 25 Go et quatre LUN iSCSI de 25 Go.
4K
- 100 % de lecture ou 100 % d'écriture
- 100% 4K
8K70/30
- 70 % de lecture, 30 % d'écriture
- 100% 8K
128K (séquentiel)
- 100 % de lecture ou 100 % d'écriture
- 100% 128K
Notre premier test de débit mesure les performances aléatoires 4K du QNAP TS-h1887XU-RP. Ici, il a affiché des performances similaires dans les configurations iSCSI et SMB. En termes de performances SSD, il a montré 45,543 17,735 IOPS en lecture et 44,448 17,891 IOPS en écriture (SMB) et 1887 1,878 IOPS en lecture et 1,882 2831 IOPS en écriture (iSCSI). En termes de performances du disque dur, le TS-h2350XU-RP a enregistré XNUMX XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture (SMB) et XNUMX XNUMX en lecture et XNUMX XNUMX IOPS en écriture (iSCSI).
En latence moyenne, le TS-h1887XU-RP enregistre à nouveau des performances similaires en SMB et iSCSI. Pour les performances SSD, il affiche des lectures et écritures de 5.62 ms et 14.43 ms (SMB) et 5.76 ms et 14.31 ms (iSCSI SSD). Pour la configuration du disque dur, le TS-h1887XU-RP avait 136.26 ms en lecture et 135.93 ms en écriture, et 90.40 ms en lecture et 108.89 ms en écriture.
