---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-3
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "blackwell", "ethernet", "gpu", "intel", "nvidia", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [244, 386]
sha256: 72573a06cb69260600dd7f71ff60a45623adbda68f521ab7d3ec7d3cf30646d9
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| Workload | Débit cible | NIC recommandée | Pourquoi |
|---|---|---|---|
| Serveur généraliste / VMs | 2×25G ou 2×100G | Intel E810 / CX-6 Dx | SR-IOV mature, drivers inbox, prix |
| Stockage NVMe-oF / Ceph | 2×100G | CX-6 Dx / E810 | RDMA, faible latence, multipath |
| Backend IA Hopper (H100) | 400G / GPU | ConnectX-7 ou Thor 2 | RoCEv2 + CC ; Thor 2 si fabric Broadcom |
| Backend IA Blackwell (B200) | 800G / GPU | ConnectX-8 / BF-3 SuperNIC | Spectrum-X de bout en bout |
| Backend IA Rubin (2026+) | 1,6T / GPU | ConnectX-9 (annoncé) | Scale-out 1,6 Tb/s ✅ |
| Management / OOB | 1-25G | LOM / SFP28 éco | Ne pas gaspiller du 100G |
| Firewall/edge | 100G+ DPDK | E810 / CX-6 Dx | PPS, drivers stables |

**Questions à poser au vendeur** (checklist §109) :
1. Firmware minimal requis pour le débit annoncé ? (ex. CX-8 : 40.46.x)
2. Driver inbox kernel ou OFED/DOCA propriétaire ? Qui le maintient ?
3. La carte est-elle codée constructeur (lock optique) ? (P3, §107)
4. Conso max avec optiques (pas la carte seule) ?
5. Support UEC / RoCEv2 / PFC-DCQCN validé sur quel switch ?

## 12. Formats physiques : PCIe CEM, OCP 3.0, Socket Direct

| Format | Usage | Avantage | Limite |
|---|---|---|---|
| PCIe CEM (FH/HL) | Serveurs généralistes | Universel | Encombrement, 1 slot |
| OCP 3.0 (SFF) | Serveurs hyperscale/OEM | Hot-swap, frontal, dense | Châssis compatibles requis |
| Socket Direct | Serveurs IA NVIDIA | Latence NUMA minimale | Carte auxiliaire + câblage |
| Mezzanine (ex. OCP) | Lames | Intégré | Propriétaire |

✅ ConnectX-8 existe en OCP 3.0 et CEM x16 (vérifié 27/09/2026).
**Piège d'achat** : une NIC OCP 3.0 ne rentre pas dans un serveur tour/1U
classique — valider le format **avant** la commande, pas à la réception.

## 13. Firmware & drivers : la moitié du produit

- **NVIDIA** : MLNX_OFED (ou DOCA-OFED) vs driver inbox. Pour RoCEv2 +
  Spectrum-X, prendre le driver validé par le fabric, pas le dernier cri.
  `mlxfwmanager` pour le firmware ; ne jamais mélanger FW 40.45 et 40.46 sur
  un même fabric sans valider.
- **Intel** : drivers `ice` inbox depuis kernel 5.x — le plus simple à
  maintenir. NVM update via `ice` tools.
- **Broadcom** : `bnxt_en` inbox ; firmware via les OEM (Dell/HPE) — prévoir
  le cycle de qualification OEM (souvent 1-2 trimestres de retard).
- **Règle d'or** : figer un triplet (FW NIC, driver, OS) par cluster et
  l'appliquer à tout le fabric. Un nœud avec un FW différent = suspect n°1
  en cas de micro-pertes (P4, §107).

## 14. CAS CHIFFRÉ — 8 nœuds GPU Hopper, backend 400G

Hypothèses : 8 serveurs × 8 GPU H100, 1 NIC 400G par GPU (rail-optimized).

| Poste | Calcul | Résultat |
|---|---|---|
| NIC 400G | 64 × ConnectX-7 / Thor 2 | 64 NIC |
| Slots PCIe Gen5 x16 | 8 par serveur | Vérifier carte mère 8× x16 Gen5 ! |
| Conso NIC seules | 64 × ~35 W (⚠️) | ~2,2 kW |
| Optiques 400G DR4 | 64 × ~9 W (⚠️) | ~0,6 kW |
| Ports leaf 400G | 64 | 1 leaf 64×400G (ex. 32 ports en 2×200G si besoin) |
| Câblage intra-rack | DAC 400G si < 3 m, sinon AOC/DR4 | Voir §22 |

**Budget indicatif (⚠️ fourchettes marché, à vérifier)** : NIC 400G ~1 500-3 000 €,
optique 400G DR4 ~600-1 500 € → **~135-290 k€** pour 64 NIC + optiques seules.
Le réseau backend d'un cluster IA = 10-15 % du coût des GPU : le provisionner
en même temps, pas après.

## 15. Pièges NICs — l'essentiel (détail §107)

- **P1** — NIC 400G dans slot PCIe Gen4 : bridée à ~200G, invisible sans test.
- **P2** — Offloads désactivés par un vieux driver : CPU à 100 % à 100G.
- **P3** — Optiques tierces refusées par firmware verrouillé : tester 1 avant 100.
- **P4** — Firmwares hétérogènes sur un fabric RoCE : micro-pertes aléatoires.
- **P5** — OCP 3.0 commandée pour châssis CEM : ne rentre pas physiquement.

# PARTIE B — FIBRE & TRANSCEIVERS (« module stp fibre »)

## 16. Lexique des formats : du SFP à l'OSFP

« Module STP fibre » = en pratique un **transceiver** (émetteur-récepteur)
enfichable. Les formats, par ordre d'apparition :

| Format | Lanes élec. | Débit max usuel | Taille | Génération |
|---|---|---|---|---|
| SFP / SFP+ | 1 | 1G / 10G | Petit | Legacy |
| SFP28 | 1×25G NRZ | 25G | = SFP+ | 2015+ |
| SFP56 | 1×50G PAM4 | 50G | = SFP+ | 2019+ |
| SFP112 | 1×100G PAM4 | 100G | = SFP+ | 2024+ |
| QSFP+ | 4×10G | 40G | Compact 4 lanes | Legacy |
| QSFP28 | 4×25G NRZ | 100G | = QSFP+ | Standard DC |
| QSFP56 | 4×50G PAM4 | 200G | = QSFP+ | Transition |
| QSFP-DD | 8×25G/50G | 400G | Double densité (DD) | Backend IA |
| QSFP112 | 4×100G PAM4 | 400G | = QSFP-DD | Backend IA |
| OSFP | 8×50G/100G PAM4 | 400G/800G | Plus grand, meilleure thermique | IA / IB |
| QSFP-DD800 | 8×100G PAM4 | 800G | = QSFP-DD | 800G Ethernet |
| OSFP-XD / OSFP1600 | 8×200G PAM4 | 1,6T | En cours | Futur (§99) |

**Règle** : à débit égal, OSFP dissipe mieux que QSFP-DD (plus de surface,
dissipateur intégré). C'est pour ça que NVIDIA (Spectrum-X, Quantum) est
OSFP-natif : 800G = 14-18 W par module, il faut du métal autour.
⚠️ Commander des OSFP « flat-top » pour NIC et « finned-top » (à ailettes)
pour switch : se tromper = surchauffe (voir §107, P6).

## 17. Correspondance débit ⇄ format ⇄ fibre — le tableau de référence

| Débit | Transceiver type | Format | Fibre | Connecteur |
|---|---|---|---|---|
| 10G SR | 10GBASE-SR | SFP+ | OM3/OM4 | LC duplex |
| 25G SR | 25GBASE-SR | SFP28 | OM3/OM4 | LC duplex |
| 100G SR4 | 100GBASE-SR4 | QSFP28 | OM3/OM4 (8 brins) | MPO-12 |
| 100G DR | 100GBASE-DR | QSFP28 | OS2 (2 brins) | LC duplex |
| 100G FR1/LR4 | 100GBASE-FR1/LR4 | QSFP28 | OS2 | LC duplex |
| 200G SR4 | 200GBASE-SR4 | QSFP56 | OM3/OM4 | MPO-12 |
| 400G SR8 | 400GBASE-SR8 | QSFP-DD/OSFP | OM3/OM4 (16 brins) | MPO-16 |
| 400G DR4 | 400GBASE-DR4 | QSFP-DD/OSFP | OS2 (8 brins) | MPO-12 |
| 400G FR4/LR4 | 400GBASE-FR4/LR4 | QSFP-DD/OSFP | OS2 (2 brins, WDM) | LC duplex |
| 800G SR8 | 800GBASE-SR8 | OSFP | OM3/OM4 (16 brins) | MPO-16 |
| 800G DR8 | 800GBASE-DR8 | OSFP | OS2 (16 brins) | MPO-16 |
| 800G 2×FR4 | 800GBASE-2FR4 | OSFP | OS2 (4 brins) | 2× LC duplex (CS) |

**Lecture** : SR = Short Reach (multimode), DR = 500 m (monomode),
FR = 2 km, LR = 10 km. Le chiffre après = nombre de lanes optiques
(SR8 = 8 lanes de 100G sur 16 brins : 8 Tx + 8 Rx).

## 18. Types de fibre : OM3 / OM4 / OM5 / OS2

| Fibre | Type | Cœur | Bandeau modal (EMB) | Usage |
|---|---|---|---|---|
| OM3 | Multimode 50/125 µm | 50 µm | 2000 MHz·km @850 nm | Legacy 10/40/100G |
| OM4 | Multimode 50/125 µm | 50 µm | 4700 MHz·km @850 nm | Standard DC actuel |
| OM5 | Multimode 50/125 µm | 50 µm | 4700 + SWDM (4 λ) | 100G sur 2 brins (BiDi/SWDM4) |
| OS2 | Monomode 9/125 µm | 9 µm | — (une seule voie) | Tout ce qui dépasse 100-150 m |

**Règles de décision :**
- **Intra-datacenter < 100 m** : OM4 (SR). Le moins cher au mètre et au module.
- **> 100 m ou inter-bâtiments** : OS2 systématique. La différence de prix du
  câble est anecdotique face au coût d'un re-câblage.
- **Neuf en 2026** : tirer de l'OS2 partout dans les rocades (même si on
  branche de l'OM4 en bout), et de l'OM4 vers les baies. L'OM5 ne se justifie
  que pour des cas SWDM/BiDi précis — pas un standard de rocade.
- **Couleurs** : OM3/OM4 = gaine **aqua**, OM5 = **vert lime**, OS2 = **jaune**.
  Un jarretière jaune sur un port SR = suspect immédiat (P7, §107).

## 19. Distances maximales — tableau normatif usuel

Valeurs IEEE 802.3 / MSA usuelles (⚠️ vérifier la datasheet du module exact,
les optiques « extended » dépassent parfois la norme) :

