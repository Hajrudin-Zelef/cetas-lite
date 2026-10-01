---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-1
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "agent", "arr", "gpu", "nvidia"]
source: docs/RAG/clean4/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [1, 77]
sha256: aa0474956cc9873f2e4e1521f9d3fd5416c358ef61191ed7b73959d5d4d16a84
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

Nous avons récemment reçu le serveur Supermicro AS-2126HS-TN, un serveur A+ biprocesseur basé sur l'architecture Turin, pour test. Cette plateforme 2U est conçue pour prendre en charge des charges de travail extrêmement gourmandes en ressources et offre une extension PCIe flexible. Basé sur deux processeurs AMD EPYC séries 9005 (Turin) et 9004 (Genoa), ce système cible un large éventail de cas d'utilisation en entreprise et en centre de données, notamment la virtualisation, le stockage défini par logiciel, l'inférence IA et l'apprentissage automatique, le cloud computing, la consolidation de serveurs d'entreprise et le calcul haute performance (HPC).
La carte AS-2126HS-TN prend en charge plusieurs configurations d'emplacements PCIe, permettant ainsi de privilégier la densité d'accélérateurs ou la flexibilité des E/S. Selon la configuration, la plateforme peut accueillir jusqu'à quatre emplacements PCIe 5.0 x16, jusqu'à huit emplacements PCIe 5.0 x8, ou des configurations mixtes pour intégrer des GPU, des contrôleurs réseau haut débit et des contrôleurs de stockage. Cette approche modulaire permet au système de s'adapter aux charges de travail exigeantes en calcul, en E/S ou axées sur les accélérateurs, sans nécessiter de modifications du châssis.
Pour évaluer la plateforme, nous avons soumis l'AS-2126HS-TN à notre environnement de tests d'entreprise, en nous concentrant sur des charges de travail gourmandes en ressources CPU reflétant des scénarios de déploiement réels. Nos tests mettent l'accent sur les performances multithread soutenues, le comportement de la mémoire et l'efficacité globale de la plateforme, offrant ainsi un aperçu des performances de l'architecture Turin de Supermicro dans des conditions exigeantes d'entreprise et de calcul haute performance (HPC).
Spécifications du Supermicro AS-2126HS-TN
Le tableau ci-dessous présente les spécifications matérielles du Supermicro AS-2126HS-TN, offrant un aperçu de sa conception de plateforme, de ses capacités de calcul, de ses options d'extension et de ses caractéristiques d'alimentation et de refroidissement.
| Spécifications | Serveur Supermicro A+ AS -2126HS-TN | 
|---|---|
| Présentation du système |  | 
| Produit REF | Serveur A+ AS -2126HS-TN | 
| Carte mère | Super H14DSH | 
| Facteur de forme | Montage en rack 2U | 
| Modèle de châssis | CSE-HS201-R000NFP | 
| Processeurs |  | 
| Support de CPU | Processeur double ; processeurs AMD EPYC™ série 9005/9004 | 
| Nombre maximal de cœurs | Jusqu'à 384 °C / 768 °C | 
| TDP maximal du processeur | Prend en charge les processeurs jusqu'à 500 W de TDP* | 
| Note sur le refroidissement du processeur | Les processeurs refroidis par air avec un TDP supérieur à 400 W ne sont pris en charge que dans des conditions spécifiques. | 
| GPU et accélération |  | 
| Nombre maximal de GPU | Jusqu'à 3 GPU double largeur | 
| GPU pris en charge (PCIe) | NVIDIA : H100 NVL, RTX 6000 génération Ada, L4 AMD : Instinct™ MI210 | 
| Interconnexion GPU-GPU | PCIe | 
| Mémoire |  | 
| Emplacements DIMM | 24 emplacements DIMM | 
| Mémoire maximale (1DPC, EPYC 9005) | Jusqu'à 6 To de mémoire RDIMM ECC DDR5 à 4 800 MT/s | 
| Mémoire maximale (1DPC, EPYC 9004) | Jusqu'à 6 To de mémoire RDIMM ECC DDR5 à 4 800 MT/s | 
| Tension de mémoire | 1.1V | 
| Dispositifs embarqués et réseau |  | 
| Chipset | Système sur puce | 
| Connectivité réseau | Via AIOM (options AOC disponibles) | 
| BMC / IPMI | IPMI 2.0 avec prise en charge de Virtual Media over LAN et KVM-over-LAN | 
| Entrée / Sortie |  | 
| LAN (BMC) | 1 port LAN BMC dédié RJ45 1GbE | 
| USB | 2 ports USB 3.0 (arrière) | 
| Vidéo | 1×VGA | 
| TPM | 1 TPM embarqué / port 80 | 
| BIOS |  | 
| Type de BIOS | AMI Flash EEPROM 64 Mo SPI | 
| Fonctionnalités du BIOS | Plug and Play (PnP) ; UEFI 2.8 ; prise en charge des claviers USB ; ACPI 6.5 ; SMBIOS 3.7 ou version ultérieure | 
| Direction |  | 
| Logiciels / Outils | SuperCloud Composer® ; Supermicro Server Manager (SSM) ; Super Diagnostics Offline (SDO) ; KVM avec LAN dédié ; IPMI 2.0 ; Service d’agent léger Supermicro (TAS) ; Assistant d'automatisation SuperServer (SAA) ; IPMIView | 
| Configurations d'alimentation | Gestion de l'alimentation ACPI/APM ; contrôle du mode de mise sous tension pour la récupération après une coupure de courant alternatif ; mécanisme de contournement du bouton d'alimentation | 
| Sécurité |  | 
| Sécurité matérielle | TPM 2.0 ; Racine de confiance en silicium (RoT) – Conforme à la norme NIST 800-193 | 
| Caractéristiques de sécurité | Micrologiciel signé cryptographiquement ; Démarrage sécurisé ; Mises à jour sécurisées du micrologiciel ; Récupération automatique du micrologiciel ; Sécurité de la chaîne d'approvisionnement (attestation à distance) ; protections BMC en temps réel ; verrouillage du système | 
| Surveillance de la santé PC |  | 
| Surveillance du ventilateur | Surveillance du tachymètre ; état du régulateur de vitesse ; connecteurs de ventilateur PWM | 
| Surveillance de la température | Surveillance de l'environnement du processeur et du châssis ; contrôle thermique des connecteurs de ventilateurs | 
| Tension / Capteurs | Température du système ; Température de la mémoire ; Température du processeur ; Veille 3.3 V ; Veille +5 V ; +5 V ; +3.3 V ; +12 V ; Protection contre la surchauffe du processeur | 
| Panneau avant |  | 
| Voyants | Activité du disque dur ; activité du réseau local ; état de l’alimentation ; informations système | 
| Boutons | Marche/Arrêt ; UID | 
| Extension et interconnexion |  | 
| Configuration des emplacements PCIe | Option A*: 4 emplacements PCIe 5.0 x16 FHFL double largeur ; 1 emplacement PCIe 5.0 x16 AIOM (compatible OCP 3.0) Option B*: 8 emplacements PCIe 5.0 x8 (dans x16) FHFL ; 1 module PCIe 5.0 x16 AIOM (compatible OCP 3.0) *Nécessite des pièces supplémentaires ; voir la liste des pièces optionnelles. Reportez-vous aux schémas du système pour plus de détails. | 
| Assistance CXL | Jusqu'à 4 appareils CXL 2.0 x16 | 
| Stockage |  | 
| Baies de disques (par défaut) | 8 baies au total ; 8 baies avant remplaçables à chaud pour disques NVMe*/SATA* de 2.5 pouces | 
| Baies de disques (Option A) | 24 baies au total ; 24 baies avant remplaçables à chaud pour disques NVMe*/SATA* de 2.5 pouces | 
| M.2 | 2 × M.2 PCIe 3.0 x4 NVMe (clé M 2280/22110) | 
| Remarque sur le stockage | La prise en charge NVMe/SATA peut nécessiter un contrôleur de stockage et/ou des câbles supplémentaires. | 
| Refroidissement |  | 
| Ventilateurs | Jusqu'à 6 ventilateurs contrarotatifs de 60×60×56 mm | 
| Suaire d'air | 2 enveloppes d'air | 
| Tuning Moteur |  | 
| Options d'alimentation | 2× 1200W / 1300W / 1600W / 2000W / 2600W (varie selon la configuration) | 
| Alimentation incluse/mentionnée | 2 × 2000 W redondants (1+1) Niveau Titane (96 %) | 
| Dimensions du bloc d'alimentation (L×H×P) | 73.5 × 40 x 265 mm | 
| Entrée (variable selon l'alimentation) | 1000 W : 100–127 Vca / 50–60 Hz 1800 W : 200–220 Vca / 50–60 Hz 2700 W : 200–240 Vca 1980 W : 220–230 Vca / 50–60 Hz 2000 W : 220–240 Vca / 50–60 Hz (homologué UL uniquement) 2000 W : 230–240 Vca / 50–60 Hz 2000 W : 230–240 Vcc / 50–60 Hz (CQC uniquement) | 
| Rail +12V (varie selon l'entrée) | Max 83 A (100–127 Vca) / Max 150 A (200–220 Vca) / Max 225 A (200–240 Vca) Max 165 A (220–230 Vca) / Max 166 A (230–240 Vca) | 
| SB 12V | Max. 3.5 A / Min. 0 A | 
| Type de sortie | Fond de panier (doigt d'or) | 
| Dimensions et poids |  | 
| Hauteur | 3.5 mm (88.9 po) | 
| Largeur | 17.2 mm (437 po) | 
| Profondeur | 31.74 mm (806.2 po) | 
| Emballage (H×L×P) | 9.96 "× 26.46" × 43.31 " | 
| Poids | Poids brut : 34 kg (75 lb) ; Poids net : 20.5 kg (45 lb) | 
| Couleur disponible | un Prix d'argent | 
