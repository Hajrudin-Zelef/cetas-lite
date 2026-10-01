---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-26
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "CISA", "EU", "Google", "Intel", "Microsoft", "Nvidia", "United States"]
dates: ["2024-08-13", "2026-06-22", "2026-09-14", "2026-09-17", "2026-09-27", "2027-01-01"]
keywords: ["datacenter", "amd", "asic", "aws", "chiplet", "dsp", "gpu", "intel", "luna", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3351, 3484]
sha256: 17cff79f7891ed24bb52e9b4f2077fc9f9e96b69f255a11abc45817dc0152d0e
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Tableaux : FIPS 140 (5), form factors HSM (15), cloud HSM prix (19), YubiHSM 2 (29),
comparatif HSM (30), performances tps (31), certifications (32), custodians (44),
prix HSM (47), TCO 5 ans (48), TPM serveur (52), mécanismes RoT (68), FPGA vs GPU vs
ASIC (80), conso FPGA (86), grille décision FPGA (87), équipements physiques (104),
rotation clés (108), matrice rôles (110), KMS (117), journaux (127), consos énergie (139),
ROI IPsec (176), SmartNIC/DPU/FPGA (177), perf/watt (182), grille notation (193),
registre firmware (200).
Schémas ASCII : placement HSM datacenter (33), PME (34), chaîne de confiance (50),
4 cercles physiques (91), architecture électrique (141), SIEM crypto (189), cycle de
vie des clés (105), rotation DEK (186).

## C. Sources et références (vérifiées le 27/09/2026)

- Thales Luna Network/PCIe HSM — product briefs (cpl.thalesgroup.com, via revendeurs ;
  chiffres : 84 W typ./110 W max, MTBF 171 308 h ; PCIe : 14 W typ./18 W max, MTBF
  997 508 h ; perfs RSA/ECC/AES par modèle ; Luna PQC Functionality Module).
- Utimaco — communiqué septembre 2026 (u.trust GP HSM : 31 conteneurs, Se100→Se40k,
  FIPS 140-2 L3 / L4 en cours, PCI PTS HSM V3, Common Criteria CCN Espagne, Quantum
  Protect) ; datasheet SecurityServer Se Gen2 (PCIe/LAN, n-of-m, SHA-3).
- Marvell + Microsoft + Utimaco — Azure Payment HSM v2, public preview 17/09/2026
  (Western US, Western Europe), Marvell LiquidSecurity + Atalla.
- pqca/wg-readiness-tracking (GitHub, 09/2026) — PQC HSM : Utimaco Ready (ML-KEM,
  ML-DSA, XMSS, LMS, HSS), YubiHSM 2 v2.4+ pilote uniquement.
- Yubico YubiHSM 2 — specs (nano USB-A 12×13×3,1 mm, 1 g, 20/30 mA, RSA/ECC/AES/Ed25519,
  PKCS#11/CNG) ; prix constatés 2026 : ~650 € standard, ~949 € FIPS (revendeurs EU).
- AMD Alveo V80 — product brief (Versal XCV80, 2,6 M LUT, 10 848 DSP, 32 Go HBM2e
  820 Go/s, 4× QSFP56, PCIe Gen4 x16/Gen5 x8, 190 W TDP, SKU A-V80-P64G-PQ-G) ;
  prix MSRP constaté 9 495 $.
- Intel/Altera Agilex 7 — device overview 2025 (PCIe 5.0, CXL, 116 Gb/s, chiplet EMIB) ;
  SmartNIC N6000/N6001-PL ; FPGA AI Suite + OpenVINO (altera-fpga.github.io).
- NIST — FIPS 203/204/205 finalisés 13/08/2024 ; FIPS 206 (HQC) en cours (attendu
  2026-2027, à vérifier) ; IR 8547 (retrait algos classiques 2030-2035).
- NSA CNSA 2.0 v2.1 (12/2024) — jalons 01/2027, 12/2030, 12/2031 ; CNSSP 15.
- EO 14412 (22/06/2026), OMB M-26-15 (5 phases 2026-2035), EO 14144 (TLS 1.3 PQC).
- TLS hybride X25519MLKEM768 — supporté Chrome/Firefox/Cloudflare/AWS/Schannel (2026).
- Affaire HAWK — retrait juillet 2026 après découverte d'une faiblesse (Anthropic).
- DDRop — divulgation 14/09/2026 (KU Leuven, ETH Zurich, Durham, Google) : interposeur
  DDR5 < 200 $ contre l'intégrité mémoire TDX/SGX/SEV-SNP.
- Confidential computing — Intel TDX, AMD SEV-SNP, Arm CCA (Armv9), NVIDIA H100 CC
  (VRAM AES-256-GCM, ~2-5 % overhead), Red Hat OpenShift sandboxed containers 1.12
  (04/2026, Tech Preview GPU confidentiel), Trustee/Confidential Containers.
- AWS CloudHSM — ~1,45-1,50 $/h, FIPS 140-2 L3 ; Azure Dedicated HSM — 4,85 $/h
  (page tarifs Azure).
- Non trouvés au 27/09/2026 (détails produit/tarifs) : Entrust nShield, Futurex,
  Atalla AT1000 (détails), prix publics Thales/Utimaco (sur devis uniquement).

## D. Historique du document

- **v1.0 — 27/09/2026** : création. 207 sections, glossaire 40 termes, quiz 10 Q/R,
  25 pièges terrain. Recherche web préalable le 27/09/2026 ; mentions « vérifié le
  27/09/2026 » / « à vérifier » / « non trouvée au 27/09/2026 ».
- Prochaine revue conseillée : **T1 2027** (finalisation FIPS 206, échéance CNSA 2.0
  du 01/01/2027, évolutions PQC des HSM).

# PARTIE 16 — SAUVEGARDES, RÉSEAU CRYPTO ET PRA ÉTENDU

## 208. Chiffrement des sauvegardes : l'architecture complète

Les sauvegardes contiennent tout — elles méritent le même niveau que la production :

```
[Données prod] ──chiffrées (DEK)──▶ [Sauvegarde chiffrée]
                                         │
[DEK] ──enveloppée (KEK du HSM)──▶ [stockée avec la sauvegarde]
                                         │
[KEK] ──dans le HSM──▶ [backup HSM] ──MofN──▶ [coffre]
```

Règles :

- Le logiciel de sauvegarde (Veeam, etc.) chiffre avec une clé dont la **KEK est dans
  le HSM/KMS** — jamais de mot de passe « dans la doc du logiciel ».
- **Air gap** : une copie offline (bande, disque déconnecté, bucket immuable) contre
  le ransomware — le chiffrement seul ne suffit pas si l'attaquant a les clés.
- **Immuabilité** : verrouillage d'objet (object lock) sur la copie distante : même
  avec les credentials, on ne peut pas effacer avant l'échéance.
- **Test** : restauration complète annuelle, incluant le **déchiffrement** (la clé
  existe-t-elle encore ? — voir section 138).
- **Rétention vs clés** : une sauvegarde de 7 ans exige que la clé vive 7 ans —
  coordonner avec la politique de rotation (section 108).

## 209. Ransomware et clés : le scénario à préparer

Un ransomware moderne fait deux choses : **chiffre** vos données et **vole** vos
sauvegardes si elles sont accessibles. La défense crypto :

1. **Clés hors d'atteinte** : le ransomware qui compromet l'OS n'atteint pas les clés
   du HSM (périmètre inviolable) — mais il peut **utiliser** le HSM via les applis
   compromises : d'où la surveillance des opérations anormales (section 120).
2. **Sauvegardes immuables + offline** : la seule vraie parade (section 208).
3. **Ne pas payer** : la clé de déchiffrement de l'attaquant n'est pas garantie, et
   payer finance l'écosystème — position à faire valider par la direction **avant**
   la crise.
4. **Exercice** : simuler un ransomware un vendredi soir — qui décide quoi, avec
   quelles clés, dans quel ordre (voir section 190).

## 210. Segmentation réseau de la zone crypto (détaillée)

Au-delà du VLAN dédié (section 35), l'architecture cible :

```
[Internet] ─▶ [FW périmétrique] ─▶ [DMZ] ─▶ [FW interne] ─▶ [VLAN prod]
                                                              │
[Postes admin] ─▶ [Bastion] ─▶ [FW mgmt] ─▶ [VLAN management]  │
                                              (BMC, Redfish)  │
[VLAN crypto] ◀── [FW crypto] ◀── autorisé depuis : VLAN prod (clients PKCS#11
   (HSM, Vault,   (règles par IP+port,        déclarés uniquement) + bastion
    PKI, TSA)      TLS mutuel, IDS)            (admin)
```

- **Firewall dédié** (physique ou virtuel) pour la zone crypto — pas de simples ACL
  sur le switch.
- **IDS/IPS** en écoute sur le trafic vers le HSM (détection d'anomalies).
- **Aucune route** entre VLAN crypto et Internet/DMZ — même via le firewall.
- **Bastion** : seul point d'administration, MFA obligatoire, sessions enregistrées.
- **NAC** (802.1X) sur les ports du VLAN crypto : un équipement non autorisé branché
  physiquement n'obtient pas d'IP.

## 211. Bastion d'administration : le poste fortifié

Le bastion est la **porte d'entrée unique** vers la zone crypto et le VLAN management :

- OS durci, **pas d'Internet**, pas de mail, pas de navigation.
- Accès par **MFA** (carte à puce / YubiKey + PIN).
- **Enregistrement des sessions** (audit trail) — qui a fait quoi sur le HSM.
- Comptes nominatifs, pas de compte partagé « admin ».
- Mises à jour prioritaires, antivirus/EDR, chiffrement disque.
- **2 bastions** (ou un accès de secours documenté) : le bastion est un SPOF
  assumé — prévoir la procédure si le bastion tombe.

## 212. Supervision avancée : corréler la crypto et l'infra

Aller au-delà des alertes unitaires (section 144) avec des **règles de corrélation** :

