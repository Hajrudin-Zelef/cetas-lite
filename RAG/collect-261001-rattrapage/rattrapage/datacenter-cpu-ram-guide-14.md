---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-14
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["18a", "amd", "arr", "compute", "distribution", "dram", "gpu", "hbm", "intel", "liquid cooling", "lpddr5x", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [2140, 2297]
sha256: 64a3efc57daa95fae8f6dac0ec7cab1c7f4a838e301c08ae20e2cf1fc991741c
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

## 118. Piège n°14 : le câble trop long dans la baie

Des câbles réseau/alimentation de 3 m enroulés derrière un 2U dense
bloquent le flux d'air de sortie : +5 à +10 °C sur les composants arrière,
ventilateurs qui s'emballent, bruit, surconsommation (loi cubique,
section 62). **Câbles courts (0,5-1 m), chemins séparés, rien qui pende**
(section 79).

## 119. Piège n°15 : confondre TDP et consommation

Le TDP (500 W) n'est ni la consommation moyenne (~300-350 W à 30 % de
charge), ni le pic (jusqu'au PPT, parfois +10-15 %). Dimensionner
l'électrique sur le TDP sous-estime les pics ; dimensionner la facture sur
le TDP surestime le coût. **Trois chiffres, trois usages** : pic → protection
électrique, moyenne → facture, TDP → refroidissement (section 9).

## 120. Piège n°16 : acheter la fréquence max pour un workload qui s'en moque

Du DDR5-6400 à prix premium pour de la virtualisation généraliste où le
goulot est le CPU ou le stockage : de l'argent jeté. La fréquence mémoire
ne sert que les workloads **sensibles à la bande passante** (HPC, DB
analytique, section 30). **Achetez la fréquence que votre workload peut
consommer**, pas celle du catalogue.

## 121. Piège n°17 : le 4S « pour la croissance »

Un châssis 4S coûte 2-3× plus cher à performance égale, les licences
explosent (4 sockets), et la croissance se fait rarement socket par socket
(section 27). **La croissance se gère en scale-out** (ajout de nœuds 2S),
pas en scale-up spéculatif. N'achetez du 4S que pour un workload
monolithique avéré et chiffré.

## 122. Piège n°18 : négliger la fin de vie à l'achat

Disques à détruire de façon certifiée, batteries onduleur à recycler,
serveurs à revendre ou à dépolluer (DEEE) : la fin de vie coûte de
l'argent et engage la responsabilité (données). **Négociez la reprise et
la destruction certifiée à l'achat**, quand vous avez du levier, pas 5 ans
plus tard quand le prestataire vous tient.

## 123. Glossaire (A–M)

- **AMX (Advanced Matrix Extensions)** : accélérateur matriciel intégré aux
  Xeon P-core ; décisif pour l'inférence IA sur CPU (sections 18, 22).
- **ASHRAE** : association qui définit les classes d'environnement des
  salles IT (température, humidité). Classe A2 : 18-27 °C.
- **AVX-512** : jeu d'instructions vectorielles 512 bits ; chemin complet
  sur Turin et Xeon 6 (sections 4, 22).
- **Bank / bank group** : subdivisions internes des puces DRAM ; la DDR5
  en a 32 (section 43).
- **Blade** : serveur lame mutualisant alim/ventilateurs/réseau dans un
  châssis (section 77).
- **BMC (Baseboard Management Controller)** : processeur de supervision de
  la carte mère (températures, ventilateurs, alimentations), accessible via
  IPMI/Redfish même serveur éteint.
- **BOM (Bill of Materials)** : liste chiffrée des composants d'une
  configuration (sections 127, 128).
- **Burst length** : nombre de transferts par commande mémoire (16 en DDR5).
- **cTDP** : TDP configurable par firmware (section 9).
- **CXL (Compute Express Link)** : standard d'extension mémoire via PCIe,
  avec cohérence de cache (section 56).
- **DIMM** : barrette mémoire (Dual Inline Memory Module).
- **DLC (Direct Liquid Cooling)** : refroidissement liquide direct sur les
  composants chauds (sections 65, 66).
- **DPC (DIMM Per Channel)** : nombre de barrettes par canal mémoire
  (section 47).
- **ECC** : correction d'erreurs mémoire, obligatoire en serveur
  (sections 41, 42).
- **E-core** : cœur « efficient » Intel, dense, sans hyper-threading
  (section 11).
- **Fan wall** : mur de ventilateurs hot-swap traversant le châssis
  (section 58).
- **Free cooling** : refroidissement utilisant l'air/eau extérieur sans
  groupe froid (sections 62, 92).
- **HBM** : mémoire empilée sur le processeur (GPU), très haut débit.
- **Hyper-threading (HT/SMT)** : 2 threads par cœur (Intel P-core, AMD).
  Les E-cores n'en ont pas.
- **IAA** : accélérateur d'analytics en mémoire des Xeon 6 (section 18).
- **IMC** : contrôleur mémoire intégré au CPU.
- **IPMI / Redfish** : protocoles de gestion à distance du BMC.
- **L3 (cache)** : mémoire cache partagée du CPU ; 384-576 Mo sur les
  amiraux 2026, jusqu'à 1 152 Mo (Venice-X).
- **LRDIMM** : barrette à buffer complet, pour les très grosses capacités
  (section 36).
- **MT/s** : mégatransferts par seconde, l'unité de débit mémoire
  (à ne pas confondre avec MHz).

---

## 124. Glossaire (N–Z)

- **N+1** : redondance avec un élément de secours (ventilateur, alimentation,
  nœud de cluster).
- **NUMA** : architecture où chaque socket a sa mémoire locale plus rapide
  que la mémoire distante (section 24).
- **NUT** : logiciel de supervision onduleur et d'arrêt ordonné des serveurs.
- **Overcommit** : allouer plus de ressources virtuelles que de ressources
  physiques (section 49).
- **P-core** : cœur « performance » Intel, avec HT, haute fréquence
  (section 11).
- **PCIe** : bus d'extension ; Gen5 = 32 GT/s, Gen6 = 64 GT/s (section 98).
- **PDU** : barre de distribution électrique en baie (sections 89, 90).
- **PUE** : ratio énergie totale / énergie IT (sections 81, 82).
- **PPL / PPT** : limites de puissance pilotées par le firmware (section 9).
- **PPR** : réparation de lignes DRAM défectueuses à chaud (section 55).
- **PWM** : pilotage des ventilateurs par modulation de largeur d'impulsion
  (section 60).
- **QAT** : accélérateur chiffrement/compression des Xeon 6 (section 18).
- **QVL** : liste de composants validés par l'OEM (section 115).
- **Rank** : ensemble de puces DRAM adressées simultanément (section 43).
- **RAS** : fiabilité/disponibilité/maintenabilité (ECC, Chipkill, PPR…).
- **RDIMM** : barrette à registre, standard serveur (section 35).
- **RibbonFET / PowerVia** : transistors GAA et alimentation par l'arrière
  du nœud Intel 18A (section 16).
- **Riser** : carte qui déporte les slots PCIe à l'horizontale en 1U
  (section 70).
- **SDDC / Chipkill** : survie à la panne d'une puce DRAM entière
  (section 55).
- **SECDED** : correction 1 bit, détection 2 bits (ECC de base).
- **Shroud** : déflecteur d'air interne du serveur (section 64).
- **SOCAMM2** : format de module mémoire compact remplaçable (LPDDR5X),
  utilisé par EPYC 9006 LP (section 93).
- **SP5 / SP6 / SP7** : sockets AMD EPYC (sections 8, 93).
- **TCO** : coût total de possession sur la durée de vie (sections 101, 102).
- **TDP** : enveloppe thermique de référence du CPU (section 9).
- **Throttling** : baisse de fréquence pour limiter la température
  (section 67).
- **TSV** : vias traversants pour l'empilement 3D des puces (section 37).
- **UPI** : interconnexion inter-sockets Intel, 24 GT/s en version 2.0
  (section 29).
- **VFI / VI / VFD** : topologies d'onduleur (double conversion, interactif,
  off-line).
- **vNUMA** : présentation de la topologie NUMA aux VM par l'hyperviseur
  (section 24).
- **VRM** : régulateurs de tension alimentant le CPU sur la carte mère.
- **3DS** : empilement 3D des puces DRAM pour la densité (section 37).
- **80 PLUS** : certification de rendement des alimentations (section 84).

## 125. Quiz : 10 questions pour valider

**Q1.** Un EPYC 9965 a 192 cœurs et 12 canaux DDR5-6400. Quelle est sa bande
passante mémoire théorique par cœur ? Est-ce adapté au HPC ?

**Q2.** Pourquoi un Xeon 6900P avec AMX est-il préférable à un EPYC dense
pour de l'inférence IA sur CPU ?

**Q3.** Vous avez 768 Go à installer sur un socket 12 canaux : 6× 128 Go ou
12× 64 Go ? Justifiez.

**Q4.** Quelle est la différence entre l'on-die ECC de la DDR5 et le « vrai »
ECC serveur ? Lequel suffit en production ?

**Q5.** Un serveur 2S consomme-t-il la même chose qu'un 1S à charge CPU
égale ? Pourquoi ?

**Q6.** Calculez l'intensité par phase d'une baie de 15 kW IT sur PDU
triphasée 400 V (cos φ = 0,95). Quelle PDU choisir avec la règle des 50 % ?

