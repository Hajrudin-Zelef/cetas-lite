---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-18
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Meta", "TSMC"]
dates: ["2025-10-08", "2026-09-27"]
keywords: ["datacenter", "cpo", "dci", "gpu", "hyperscaler", "incident", "intel", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [2506, 2674]
sha256: 741e9af7a5f2ddc36aace27b486470d66d15a8bb1135c1817f97540b836cc297
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| Capacité | **102,4 Tb/s** (2× TH5) |
| Ports (pluggable) | 64×1,6T / 128×800G / 256×400G / 512×200G |
| Version CPO | **Davisson** : 16 moteurs optiques 6,4T co-packagés |
| Conso CPO | **3,5 W/port à 800G** (−36 % vs TH5 CPO, **−70 % vs pluggable**) |
| Lasers | **Remplaçables par la face avant** (l'objection hyperscaler est levée) |
| Échelle | 131 072 XPU en 2 tiers |
| Routing | Cognitive Routing 2.0 |
| Statut | Annoncé (juin 2025) ; Davisson **early access** (08/10/2025) |
| Partenaires | Celestica, Accton, Micas, Nexthop AI, TSMC (COUPE) |
| Usage type | Fabrics IA géants 2026-2028, hyperscalers d'abord |
| Prix indicatif | ❌ non public (⚠️) |
| Point fort | Le CPO devient crédible : −70 % de conso optique, lasers remplaçables |
| Point faible | Early access : pas de volume enterprise avant 2026-2027 |
| Verdict | **La trajectoire à suivre** — ne pas verrouiller du 800G pluggable à 7 ans |

## 132. Fiche — Arista 7060X6 / 7800R4 (Etherlink AI)

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| 7060X6 (leaf) | TH5, 64×800G OSFP, 51,2T, **2× PSU 2400 W**, 2U |
| 7060X5 | TH4, 32×800G, 2× 1500 W |
| 7800R4 (spine) | Jericho3-AI, jusqu'à **460 Tb/s**, 576×800G ou 1152×400G par châssis |
| Gain 7800R4 | **−65 % de conso** vs 7800R3 (Jericho2) |
| NOS | **EOS** (le plus mature), CloudVision, ZTP |
| Features IA | Etherlink : monitoring par flux, tuning RoCE, AI jobs visibility |
| Référence client | Meta (cluster RoCE 24 576 H100, ✅) |
| Prix indicatif | ❌ non public — 64×800G : 150-300 k€ (⚠️) |
| Point fort | EOS + support enterprise, choix « sans surprise » |
| Point faible | Prix premium vs whitebox |
| Verdict | **Le choix enterprise du backend IA** si l'équipe connaît Arista |

## 133. Fiche — Intel E810 (la NIC 100G enterprise)

| Rubrique | Détail (✅ vérifié 27/09/2026 sauf note) |
|---|---|
| Débit | 2×100G (10/25/50/100G auto) |
| PCIe | **Gen4 x16** |
| Connecteur | 2× QSFP28 |
| Virtualisation | **256 VF SR-IOV**, VMDq, 768 VSI |
| Offloads | VXLAN/GENEVE/NVGRE, ADQ, DPDK |
| RDMA | iWARP + RoCEv2 |
| Timing | IEEE 1588 PTP v1/v2, per-packet timestamp |
| DCB | PFC (Qbb), ETS (Qaz) — lossless ready |
| Conso | **15 W** (2 ports) à **27,1 W** (avec AOC) ✅ |
| Driver | `ice` inbox (kernel 5.x+) — le plus simple à maintenir |
| Prix indicatif | 300-800 € (⚠️) |
| Usage type | Serveurs généralistes, VMs, stockage, firewall |
| Point fort | Driver inbox, SR-IOV mature, sobre |
| Point faible | Pas de 200/400G (❌ non trouvée au 27/09/2026) |
| Verdict | **La NIC 100G par défaut** quand on ne fait pas d'IA |

# PARTIE M — DIMENSIONNEMENTS PAS-À-PAS

## 134. Cas complet — cluster 2048 GPU, 3 tiers, 800G

**Hypothèses** : 256 nœuds × 8 GPU (B200), 1 NIC 800G/GPU, 1:1, rail-optimized.

**Étape 1 — Endpoints** : 2048 NIC 800G.

**Étape 2 — Pods** : pod = 256 GPU (32 nœuds). 2048/256 = **8 pods**.

**Étape 3 — Par pod** (rail-optimized, 8 plans) :
- Par plan : 256/8 = 32 endpoints → 1 leaf 64×800G (32 down + 32 up).
- Leafs/pod : 8 plans × 1 = **8 leafs**.
- Spines/pod : 8 plans × 2 = **16 spines** 64×800G (32 ports utilisés).
- Vérification 1:1 : 32×800G up / 32×800G down par leaf ✅.

**Étape 4 — Super-spines** : chaque pod expose 16 spines × 32 uplinks =
512 liens 800G vers le haut. 8 pods = 4096 liens.
Super-spines 64×800G : 4096/64 = **64 super-spines**.
Oversubscription pod→super-spine : on met 16 uplinks par spine de pod
(au lieu de 32) → ratio 2:1 inter-pods (acceptable : le trafic inter-pods
est < 20 % en training, ⚠️ hypothèse à valider par workload).

**Étape 5 — Comptes** :

| Niveau | Switchs | Ports 800G |
|---|---|---|
| Leafs | 8 pods × 8 = 64 | 64 × 64 = 4096 |
| Spines pod | 8 × 16 = 128 | 128 × 64 = 8192 |
| Super-spines | 64 | 64 × 64 = 4096 |
| **Total** | **256 switchs** | **16 384 ports** |

**Étape 6 — Énergie** (⚠️) : 256 × 1,5 kW = 384 kW (switchs) + 2048 ×
60 W = 123 kW (NIC) + ~8000 optiques × 15 W = 120 kW → **~627 kW IT**,
soit **~878 kW** au compteur (PUE 1,4).

**Étape 7 — Budget** (⚠️) : 256 × 200 k€ = 51 M€ (switchs) + 2048 ×
4,5 k€ = 9,2 M€ (NIC) + optiques/câbles ~25 M€ → **~85 M€** de réseau
pour ~2048 GPU (~500 M€ de GPU, ⚠️) : **~15 %** du cluster.

## 135. Cas complet — datacenter cloud 2000 serveurs généralistes

**Hypothèses** : 2000 serveurs × 2×25G, 3:1, EVPN/VXLAN, 2 sites (actif/actif).

**Étape 1 — Leafs** : 48×25G par leaf → 2000/48 ≈ 42 leafs (prendre 44).
Uplinks : 1200 Gb/s down → 400 Gb/s up (3:1) = 4×100G par leaf.
Uplinks totaux : 44 × 4 = 176×100G.

**Étape 2 — Spines** : switch 32×100G → 176/32 ≈ 6 spines (prendre 6).

**Étape 3 — Adressage** :
- Loopbacks : 50 × /32 dans 10.0.0.0/24.
- Liens : 176 × /31 = 352 adresses → **10.0.1.0/23** (512 adresses).
- ASN : spines 65101-65106, leafs 65001-65044.

**Étape 4 — Overlay** : VNI par tenant (16M possibles), anycast gateway,
route-reflectors sur 2 spines.

**Étape 5 — Énergie** (⚠️) : 44 × 350 W + 6 × 500 W = 18,4 kW (switchs) +
4000 NIC × 12 W = 48 kW → **~66 kW IT**, ~93 kW au compteur.

**Étape 6 — Budget** (⚠️) : 44 × 22 k€ + 6 × 37 k€ = ~1,2 M€ (switchs) +
4000 × 500 € = 2 M€ (NIC) + DAC ~0,3 M€ → **~3,5 M€**, soit **1 750 €/
serveur** de réseau. C'est le coût « banalisé » du 25/100G : deux ordres
de grandeur sous le backend IA.

## 136. Cas complet — DCI 80 km chiffré entre 2 sites

**Hypothèses** : 2 datacenters à 80 km, besoin 800 Gb/s utiles, chiffré.

**Étape 1 — Technologie** : 80 km = limite du ZR. **2× 400G-ZR** (cohérent
enfichable, OpenZR+) ou 1× 800G-ZR (⚠️ dispo à vérifier au 27/09/2026 —
prendre 2×400G-ZR, mature).

**Étape 2 — Fibre** : 1 paire OS2 louée (ou propre). Budget : 80 km ×
0,25 dB/km (⚠️ G.652 à 1550 nm avec connecteurs) ≈ 20 dB + marges →
vérifier que le budget du ZR (typiquement 22-26 dB, ⚠️) passe. Sinon :
ampli EDFA (baie optique).

**Étape 3 — Chiffrement** : MACsec 400G en HW sur les ports DCI (si
supporté) ou IPsec sur routeurs. **Ne pas transporter de clair** sur une
fibre louée.

**Étape 4 — Routage** : eBGP entre sites (ASN différents), BFD 300 ms,
pas d'extension du fabric (voir P43).

**Étape 5 — Budget** (⚠️) : 4 optiques 400G-ZR × 8-15 k€ = 32-60 k€ +
2 ports routeur 400G + fibre (location : ~1-3 k€/km/an → 80-240 k€/an !).
**La fibre représente 70 %+ du TCO** : négocier l'IRU (droit d'usage
15 ans) plutôt que la location annuelle si le besoin est durable.

## 137. Cas complet — lab de test d'un fabric (avant production)

**Objectif** : valider FW, CC, ECMP sur 8 nœuds avant de déployer 1000.

| Poste | Détail |
|---|---|
| Nœuds | 8 serveurs (même réf que la prod, ou 4 + simulateurs) |
| NIC | Les mêmes qu'en prod (même FW !) |
| Switchs | 2 leafs + 2 spines (même modèle, même NOS, même version) |
| Câblage | Même média qu'en prod (DAC/AOC/optiques) |
| Outils | nccl-tests, ib_write_bw, sondes latence, capture ECN/PFC |

**Protocole** (4 semaines type) :
1. **S1** : câblage + FW + config de référence (la « golden config »).
2. **S2** : tests §107 (bande passante, latence, ECMP, perte d'un spine).
3. **S3** : chaos — kill -9 d'un process BGP, flap d'un lien, panne PSU.
   Tout ce qui n'est pas testé **arrivera** en prod.
4. **S4** : gel de la config + documentation (c'est elle qui part en prod).

**Budget** (⚠️) : ~150-300 k€ pour un lab 800G 8 nœuds. À comparer au coût
d'une semaine de cluster 1000 GPU idle : **le lab est 100× rentabilisé**
dès le premier incident évité.
**Règle** : le lab reste **branché et à jour** après le déploiement — c'est
là qu'on teste chaque changement avant la prod (FW, config, NOS).

