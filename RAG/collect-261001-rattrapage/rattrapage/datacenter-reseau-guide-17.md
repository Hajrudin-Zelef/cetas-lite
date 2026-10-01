---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-17
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Intel", "Meta", "Nvidia", "xAI"]
dates: ["2025-06-11", "2025-10-08", "2025-10-14", "2026-09-27"]
keywords: ["asic", "blackwell", "chiplet", "cpo", "dsp", "ethernet", "gpu", "intel", "kv cache", "lpo", "nvidia", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [2366, 2505]
sha256: 8ef6d3c58f7b822b4429da81ddba4b7f0881c91a95b292b421c3d1d5ef0216bc
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Fibre**
- [ ] Type : OM4 intra-baie, OS2 en rocade — couleurs conformes
- [ ] Longueurs : DAC ≤ portée (800G : 2 m), sinon AOC/optique
- [ ] MPO : méthode de polarité et genre documentés
- [ ] Budget optique : marge ≥ 3 dB calculée par lien type
- [ ] Kit nettoyage + microscope : 1 par salle (non négociable)

**Projet**
- [ ] 2 fabrics (IA + front-end) ou justification écrite du fabric unique
- [ ] Tests de réception §107 dans le contrat (critères chiffrés)
- [ ] SLA 4 h 24/7 + spares sur site pour le backend IA
- [ ] FW/driver figés par campagne dans le contrat de maintenance
- [ ] PUE / conso : switch + optiques + NIC chiffrés (§89-93)
- [ ] Clause de sortie : pas d'engagement 800G pluggable à 7 ans (CPO)

## 125. Sources vérifiées le 27/09/2026

| Fait | Source |
|---|---|
| ConnectX-8 : 800G, PCIe Gen6, OSFP224, Socket Direct, COMPUTEX 2025, FW 40.46.x | Docs firmware NVIDIA (docs.nvidia.com), Dell |
| Spectrum-X : Spectrum-4 51,2T, SN5600/SN5610 64×800G shipping | Specs NVIDIA, analyses d'écosystème |
| SN5600 : 940 W typique ; SN5610 : 0,9 kW / 2,08 kW à 64 optiques, ports 18 W | Datasheet NVIDIA SN5600 |
| BlueField-4 : 800G, Grace 64 cœurs, ConnectX-9, 128 Go, PCIe Gen6, GTC Washington 10/2025, dispo 2026 | NVIDIA, Canonical/Ubuntu |
| Broadcom Thor 2 : 400G, 5 nm, PCIe Gen5 x16, RoCE + apports UEC | Broadcom, FS, SDxCentral, STH |
| Broadcom Thor Ultra : 800G, premier NIC **UEC 1.0**, 14/10/2025 | Network World, EE Times |
| UEC 1.0 : **11/06/2025** ; 1.0.1 mi-2025 ; > 100 membres | UEC, HPCwire, DCD |
| Tomahawk 6-Davisson CPO : 102,4T, 64×1,6T, 3,5 W/port, 131 072 XPU, **08/10/2025** | Broadcom (GlobeNewswire), Next Platform |
| Broadcom Taurus : DSP 400G/lane, modules 1,6T/3,2T, OFC 03/2026, ère 200T | Broadcom OFC 2026 |
| Quantum-X800 CPO : 115,2T, 9 W/port vs 30 W | Analyses presse |
| ConnectX-9 : 1,6 Tb/s scale-out, Rubin H2 2026 | Presse spécialisée |
| Meta : 2 clusters 24 576 H100 (RoCE Arista vs IB Quantum-2), > 90 % util. | Presse spécialisée |
| xAI : Spectrum-X, 95 % vs 60 % Ethernet standard | Presse spécialisée |
| Arista 7060X6 : 64×800G, 2× 2400 W ; 7800R4 : 460T, Jericho3-AI | Arista, Next Platform |
| Intel E810 : PCIe Gen4 x16, 2×100G, 256 VF, 15-27 W | Fiches distributeurs |
| Intel : **pas de NIC 400G** au 27/09/2026 | ServeTheHome |

**Non trouvées / non vérifiées au 27/09/2026** : NIC Intel 400G ❌ · détails
Cisco Silicon One G200 ⚠️ · roadmap Intel IPU ⚠️ · prix publics précis des
switchs/NIC 800G ⚠️ (fourchettes marché uniquement) · conso exacte
ConnectX-8 / BlueField-4 ❌ (non publiée).

---

*Fin du guide — 125 sections. Vérification : `wc -l` — objectif ≥ 4000 lignes.*

# PARTIE L — FICHES PRODUITS DÉTAILLÉES (vérifiées 27/09/2026)

## 126. Fiche — NVIDIA ConnectX-8 SuperNIC

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| Débit | 800 Gb/s |
| PCIe | **Gen6** x16 (jusqu'à 48 lanes : usage switch PCIe interne) |
| Connecteurs | OSFP224 1×800G ou double 112 (2×400G) |
| Formats | OCP 3.0, CEM PCIe x16 |
| Features | In-Network Computing, MPI_Alltoall, QoS + contrôle de congestion, chiffrement TLS/IPsec inline sans perte |
| Socket Direct | Oui (carte auxiliaire 16 lanes) |
| Logiciel | DOCA 2.0, firmware 40.46.x (09/2025) |
| Lancement | COMPUTEX 2025 |
| Usage type | Serveurs RTX PRO, backend Blackwell, ratio 2:1 GPU:NIC (400G/GPU) |
| Prix indicatif | ❌ non public — fourchette marché 800G : 3 000-6 000 € (⚠️) |
| Alternatives | Broadcom Thor Ultra (UEC), Intel ❌ (pas de 800G au 27/09/2026) |
| Point fort | Commutateur PCIe intégré : remplace plusieurs switchs PCIe discrets |
| Point faible | Écosystème fermé NVIDIA, prix |
| Verdict | **La référence 800G** si on est déjà en Spectrum-X ; sinon évaluer Thor Ultra |

## 127. Fiche — Broadcom Thor 2 (BCM957608)

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| Débit | 400 Gb/s |
| PCIe | Gen5 x16 |
| Gravure | 5 nm (argument efficacité) |
| SerDes | 8 lanes : 100G PAM4 / 50G / 25G NRZ |
| RDMA | RoCEv2 + améliorations issues d'UEC : placement désordonné, ACK sélectif, multipath paquet, CC sans config |
| Cuivre | DAC passif jusqu'à ~5 m vers Tomahawk 5 |
| Optiques | Ultra-low-power linear (LPO) |
| Formats | Carte, chiplet, IP (3 modèles de consommation) |
| Usage type | Backend IA 400G sur fabric Broadcom/Arista, alternative ouverte à CX-7 |
| Prix indicatif | ❌ non public — fourchette 400G : 1 500-3 000 € (⚠️) |
| Point fort | Ouvert, basse conso, DAC longue portée |
| Point faible | Écosystème logiciel moins riche que DOCA |
| Verdict | **Le choix « ouvert » du 400G** — à privilégier hors écosystème NVIDIA |

## 128. Fiche — Broadcom Thor Ultra

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| Débit | **800G sur un seul flux** (pas 2×400G) |
| Conformité | **Premier NIC conforme UEC 1.0** (annoncé 14/10/2025) |
| RDMA | Modernisé : retransmission sélective, CC programmable émetteur/récepteur |
| Positionnement | IA scale-out uniquement (pas enterprise généraliste) |
| Interopérabilité | Fonctionne avec Tomahawk 5 existants et Tomahawk 6 futurs |
| Usage type | Backend IA 800G multi-vendeurs, pari UEC |
| Prix indicatif | ❌ non public (⚠️) |
| Point fort | UEC 1.0 natif, anti-lock-in |
| Point faible | Jeune (fin 2025) : valider le support et les drivers sur votre OS |
| Verdict | **Le pari ouvert du 800G** — exiger une preuve de concept avant volume |

## 129. Fiche — NVIDIA BlueField-4 DPU

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| Débit | **800 Gb/s** (2× BF-3) |
| CPU | **NVIDIA Grace : 64 cœurs ARM Neoverse V2** |
| Réseau | **ConnectX-9** |
| Mémoire | 128 Go |
| PCIe | **Gen6** |
| Logiciel | DOCA microservices, zero-trust, service chaining natif |
| Lancement | GTC Washington, octobre 2025 |
| Dispo | Early availability 2026, plateforme Vera Rubin |
| Usage type | AI factories : stockage AI-native (KV cache), sécurité runtime, multi-tenant |
| Systèmes | Partenaires attendus H2 2026 |
| Prix indicatif | ❌ non public (⚠️) |
| Point fort | 6× la puissance de calcul de BF-3, 800G |
| Point faible | Dispo 2026, écosystème DOCA = lock-in |
| Verdict | **À évaluer pour les plateformes Rubin** ; BF-3 reste le choix 400G immédiat |

## 130. Fiche — NVIDIA Spectrum-X SN5600 / SN5610

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| ASIC | Spectrum-4, 51,2 Tb/s, ~600 ns cut-through |
| Ports | **64× 800G OSFP** (ou 128×400G / 256×200G / 256×100G) |
| Débit forwarding | 33,3 Bpps |
| Buffer | 160 Mo partagé |
| Format | 2U, 23,5 kg |
| Conso | **940 W typique** (cuivre) ; SN5610 : 0,9 kW typ. / **2,08 kW à 64 optiques** |
| Ports optiques | Qualifiés jusqu'à **18 W** par OSFP |
| Breakout | 2×400G, 4×200G sans contrainte ; 8×100G sur port impair consomme le port pair adjacent |
| PSU | SN5600 : 2 (1+1) ; SN5610 : 4 (2+2) |
| NOS | Cumulus Linux / ONIE |
| Prix indicatif | ❌ non public — 64×800G : 150-300 k€ (⚠️) |
| Usage type | Leaf/spine Spectrum-X, backend Blackwell |
| Point fort | Boucle fermée avec SuperNIC : 95 % de throughput (cas xAI) |
| Point faible | Lock-in NVIDIA, OSFP uniquement |
| Verdict | **Le switch du backend NVIDIA** — hors écosystème NVIDIA, préférer TH5 |

## 131. Fiche — Broadcom Tomahawk 6 / 6-Davisson (CPO)

