---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-7
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Broadcom", "Meta", "Nvidia", "OpenAI", "xAI"]
dates: ["2026-09-27"]
keywords: ["asic", "aws", "cpo", "ethernet", "gpu", "llama", "nvidia", "throughput", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [870, 1012]
sha256: f38b2dcbce3cf07ee939f9ad98028df5fc8a6487b5e87ee6e07443ad682c33b8
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| Poste | Qté | Prix unit. (⚠️) | Total (⚠️) |
|---|---|---|---|
| Switch 64×800G (leaf/spine) | 12 | 150-300 k€ | 1,8-3,6 M€ |
| NIC 400G | 1024 | 1,5-3 k€ | 1,5-3,1 M€ |
| DAC/AOC 400G intra-rack | 1024 | 150-600 € | 150-600 k€ |
| Optiques 400G DR4 inter-rangées | 512 | 0,6-1,5 k€ | 300-770 k€ |
| **Total réseau** | | | **~3,8-8,1 M€** |

Soit **~30-63 k€ par nœud GPU** de réseau. À comparer aux ~250-400 k€ d'un
nœud 8×H100 (⚠️) : le réseau = **10-20 %** du nœud. Le sous-dimensionner
pour « économiser » 5 % du budget total, c'est brider 100 % du calcul.

## 42. Pièges spine-leaf — l'essentiel (détail §107)

- **P12** — Backend IA et front-end sur le même fabric : le training noie tout.
- **P13** — Zéro port libre : la croissance impose un nouveau pod.
- **P14** — MTU VXLAN oublié (+50 o.) : fragmentation silencieuse.
- **P15** — ECMP sans hash seed : polarisation, un spine saturé.

# PARTIE D — SWITCHES POUR L'IA

## 43. Ce qui distingue un switch « IA » d'un switch normal

Un switch IA doit absorber des **incast** : 1000 GPU envoient leur
gradient à 1 GPU en même temps (many-to-one). Un switch classique
droppe ; un switch IA :

| Exigence | Mécanisme | Switch classique |
|---|---|---|
| Ne pas perdre sous incast | Deep buffers + ECN + PFC | Buffers shallow, drop |
| Répartir les éléphants | Packet spraying / adaptive routing | ECMP par flux |
| Congestion de bout en bout | DCQCN / CC UEC, télémétrie | Rien (best effort) |
| Latence prévisible | Cut-through, ~600 ns-1 µs | Store-and-forward |
| Non-bloquant réel | 1:1, pas d'oversubscription interne | Souvent 3:1 interne |

**En 2026, « switch IA » =** Broadcom Jericho3-AI / Tomahawk 5-6,
NVIDIA Spectrum-4, Cisco Silicon One G200 (⚠️ détails non vérifiés au
27/09/2026). Le silicium merchant (Broadcom) équipe la majorité des
switchs OEM (Arista, Dell, HPE, Edgecore) — **le différenciateur migre
vers le logiciel** (NOS, QoS, télémétrie), pas l'ASIC.

## 44. Buffers : deep vs shallow — le choix structurant

| ASIC | Buffer | Philosophie | Usage |
|---|---|---|---|
| Broadcom Tomahawk 5 | ~Shallow (dizaines de Mo) | Faible latence, ECN agressif | Leaf IA, cloud |
| Broadcom Jericho3-AI | **Deep** (Go) | Absorbe l'incast | Spine IA ✅ |
| NVIDIA Spectrum-4 | 160 MB partagé ✅ | Cut-through, CC fine | Leaf/spine Spectrum-X |
| Broadcom Tomahawk 6 | (⚠️ à vérifier) | + Cognitive Routing 2.0 | 102.4T, 2025+ |

**Règle** : leaf = shallow + ECN (on signale vite), spine = deep
(on absorbe). Inverser = drops au spine ou latence inutile au leaf.
Le buffer se dimensionne en **BDP** (bandwidth-delay product) : à 400G
avec 10 µs de RTT, BDP = 500 Ko par flux — multiplier par le nombre de
flux concurrents donne l'ordre de grandeur du buffer utile.

## 45. ECN + PFC : le lossless Ethernet (et ses dangers)

- **PFC** (Priority Flow Control, 802.1Qbb) : pause par classe de trafic —
  crée un réseau « sans perte » pour RoCEv2.
- **ECN** : marquage des paquets en congestion → la NIC réduit son débit
  (DCQCN). C'est le mécanisme **préféré** : il évite les pauses.
- **Danger PFC** : les **deadlocks** (pause qui se propage en boucle) et les
  **pauses storms**. Un PFC mal configuré = pannes en cascade sur tout le
  fabric. D'où la règle : **PFC seulement là où c'est indispensable**
  (backend RoCE), ECN partout, et jamais de PFC sur le front-end.

Seuils usuels (⚠️ à ajuster par fabric) :
- Marquage ECN : file > 30-50 % du buffer.
- Pause PFC : file > 70-80 % (dernier recours).

## 46. RoCEv2 en production : la checklist qui fait marcher

1. **MTU 9000** partout (NIC, leaf, spine) — pas de fragmentation.
2. **DSCP 3** (ou 5) réservé au RoCE, mappé en file lossless.
3. **PFC** sur cette seule file, **ECN** activé (seuils §45).
4. **DCQCN** : paramètres par défaut du constructeur, puis tuning avec
   la télémétrie (jamais « au doigt mouillé » sur un cluster > 100 nœuds).
5. **Pas de QoS qui réordonne** : le réordonnancement tue la perf RDMA.
6. **Compteurs** : `pfc_pause`, `ecn_marked`, `roce_cnp` en supervision —
   un fabric RoCE sain a des CNP **faibles et stables**, pas zéro (zéro =
   seuils trop hauts ou supervision cassée).

## 47. InfiniBand vs Ethernet — état du marché vérifié au 27/09/2026

| Acteur | Choix backend IA | Source / statut |
|---|---|---|
| NVIDIA | **Les deux** : Quantum-X800 IB (115,2T, variante CPO Q3450-LD) et Spectrum-X Ethernet (51,2T) | Datasheet NVIDIA ✅ |
| Meta | **Les deux en test** : 2 clusters 24 576 H100 — l'un Arista RoCE, l'autre Quantum-2 IB ; **> 90 % d'utilisation réseau** sur les deux pour Llama 3 | ✅ vérifié 27/09/2026 |
| xAI (Colossus) | Spectrum-X Ethernet (contrôle de congestion NVIDIA propriétaire) : **95 % de throughput** vs 60 % en Ethernet standard | ✅ vérifié 27/09/2026 |
| OpenAI | Spectrum-X **multi-plane** | ✅ vérifié 27/09/2026 |
| Hyperscalers (AWS/GCP/Azure) | Ethernet (custom ou merchant) | Notoriété publique ⚠️ |

**Lecture** : à performance égale au sommet, le choix n'est plus
technique mais **opérationnel** :
- **InfiniBand** : lossless natif (crédits), NCCL/MPI par défaut, 20 ans
  d'outillage — idéal si l'équipe réseau est petite. Verrou : écosystème
  NVIDIA quasi-exclusif, prix.
- **Ethernet (RoCEv2/UEC)** : compétences Ethernet transférables, multi-
  fournisseurs, convergence front/back-end — mais **exige un vrai tuning**
  (ECN/DCQCN, §46). UEC 1.0 vise à standardiser ce tuning.

## 48. Tableau comparatif InfiniBand / RoCEv2 / Ultra Ethernet

| Critère | InfiniBand (NDR/XDR) | RoCEv2 | Ultra Ethernet (UEC 1.0) |
|---|---|---|---|
| Débit/port max (2026) | 800G (XDR) | 800G | 800G (1,6T en route) |
| Lossless | Natif (crédits) | À construire (PFC+ECN) | **Non requis** (conçu lossy) |
| Ordre des paquets | Garanti | Requis | **Non requis** (placement désordonné) |
| Congestion | Natif | DCQCN à régler | CC émetteur/récepteur programmable |
| Multipath | Natif | ECMP/spraying proprio | **Natif niveau paquet** |
| Retransmission | Natif | Go-Back-N (coûteux) | **Sélective** |
| Interopérabilité | Faible (1 vendeur) | Moyenne | **Objectif : totale** |
| Maturité | 20 ans | 10 ans | Spec 06/2025, produits 2025-2026 |
| Coût | Élevé | Moyen | Promesse : coût Ethernet |

**Décision 2026** : nouveau cluster NVIDIA → IB si l'équipe est petite et le
budget suit ; Ethernet si on veut du multi-vendeur ou si on opère déjà du
RoCE. **Ne pas choisir en fonction du prix catalogue** : le coût d'un tuning
RoCE raté (semaines de GPU idle) dépasse l'écart IB/Ethernet.

## 49. NVIDIA Spectrum-X — la pile Ethernet fermée

| Élément | Référence | Spec (✅ vérifié 27/09/2026) |
|---|---|---|
| ASIC | Spectrum-4 | 51,2 Tb/s, ~600 ns cut-through |
| Switch | SN5600 | 64×800G OSFP, 2U, 33,3 Bpps, 160 Mo buffer |
| Switch compact | SN5610 | 64×800G OSFP, 2U |
| NIC | ConnectX-8 SuperNIC | 800G, PCIe Gen6 |
| NIC/DPU | BlueField-3 SuperNIC | 2×400G ou 1×800G, réordonnancement HW |
| Topologie | Rail-optimized + multi-plane | LB HW inter-plans |

**Conso SN5600 (✅)** : 940 W typique (câbles cuivre passifs) ; SN5610 :
0,9 kW typique, **2,08 kW avec 64 modules optiques** — l'optique double la
facture énergétique du switch (voir §90). Ports OSFP jusqu'à 18 W/module.

**Philosophie** : boucle fermée (switch + NIC + CC + telemetry du même
vendeur) = performance garantie, **lock-in assumé**. Le packet spraying +
contrôle de congestion adaptatif + télémétrie par flux expliquent les
95 % de xAI. Alternative « ouverte » : UEC (§54).

## 50. Broadcom — le silicium derrière la moitié du marché

