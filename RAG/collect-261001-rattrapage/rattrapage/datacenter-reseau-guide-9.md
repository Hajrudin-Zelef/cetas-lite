---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-9
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia", "Oracle"]
dates: ["2026-09-27"]
keywords: ["agent", "arr", "cpo", "gpu", "kv cache", "nvidia", "rubin", "vera rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1169, 1325]
sha256: 0f113bbe38636a3e38f2c65827f7685d6559062d2184ac97c7b3f203d4801c50
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

- Module 800G-DR8 : **14-18 W** dans un volume de ~10 cm³ → densité de
  puissance comparable à un CPU. Sans airflow, la jonction dépasse 85 °C
  en minutes.
- ✅ SN5600 : ports qualifiés jusqu'à 18 W (vérifié 27/09/2026) — mais
  **tous les switchs ne le sont pas** : vérifier la spec « max power per
  port » avant d'acheter des optiques 18 W.
- **Finned-top vs flat-top** (§16) : à ailettes sur switch (air), plat sur
  NIC (souvent water-cooled ou faible airflow). Inverser = +15-20 °C.
- **Seuils** : alerte à 70 °C boîtier, critique à 80 °C (⚠️ usuels) —
  exposer les températures DOM en supervision, pas seulement les puissances.

## 61. Alimentation : dimensionner les PSU

| Switch | PSU | Conso typ. / max (✅/⚠️) |
|---|---|---|
| SN5600 64×800G | 2× (1+1) | 940 W typ. cuivre / ~2 kW optiques (⚠️ extrapolé SN5610 : 2,08 kW à 64 optiques ✅) |
| SN5610 64×800G | 4× (2+2) | 0,9 kW typ. / 2,08 kW à 64 optiques ✅ |
| Arista 7060X6 64×800G | 2× 2400 W | Dimensionné pour le pire cas ✅ |
| Arista 7060X5 32×800G | 2× 1500 W | ✅ |

**Règles** :
- **Toujours 1+1** minimum sur leaf/spine (un PSU en panne ne doit pas
  couper un plan entier).
- **Deux arrivées électriques** (A/B) par rack réseau, sur deux onduleurs
  ou deux jeux de PDU — le réseau est le dernier à tomber (il porte la
  supervision de l'arrêt d'urgence, voir guide onduleurs).
- **Appel de courant** : les ventilos + optiques au boot = pic ~120 % du
  nominal pendant 30 s (⚠️) — ne pas calibrer les disjoncteurs au nominal.

## 62. Exemple : pod 256 GPU — plan électrique et thermique

| Poste | Calcul | Résultat |
|---|---|---|
| 8 leafs + 16 spines 800G | 24 × 1,5 kW (⚠️ mixte) | 36 kW |
| 2048 NIC 800G (8/GPU × 256) | 2048 × 60 W (⚠️) | 123 kW |
| Optiques leaf/spine | ~1500 × 15 W (⚠️) | 22 kW |
| **Total réseau pod** | | **~181 kW** |
| Dissipation à évacuer (PUE 1,5) | 181 × 1,5 | **~272 kW** |
| En clim (~3,5 kW par tonneau frigo) | 272 / 3,5 | **~78 tonnes** frigorifiques |

**Lecture Zelef** : un pod réseau 800G = la clim d'un petit immeuble.
Prévoir les groupes froids **avec** le lot réseau, pas après — et noter que
le CPO (Tomahawk 6-Davisson : 3,5 W/port vs ~15 W pluggable, ✅) divise par
~4 la chaleur des optiques : **le CPO est d'abord un sujet énergie**.

## 63. Pièges 800G — l'essentiel (détail §107)

- **P20** — Airflow inversé : hotspot de 20 kW par rangée.
- **P21** — DAC 800G de 3 m « qui devrait passer » : ne passe pas.
- **P22** — Optiques 18 W sur switch qualifié 12 W : throttling thermique.
- **P23** — Disjoncteurs calibrés au nominal : déclenchement au boot.

# PARTIE F — DPU (DATA PROCESSING UNIT)

## 64. DPU : définition et place dans le serveur

Un DPU = **une NIC + des cœurs ARM + des accélérateurs** (crypto, RegEx,
décompression, RDMA) sur la même carte, avec son propre OS. Il décharge le
CPU hôte de trois familles de tâches :

```
  ┌──────────── Serveur ─────────────────┐
  │  CPU hôte : applicatif, VMs          │
  │       ▲                              │
  │       │ PCIe                         │
  │  ┌────┴──────────────────────┐       │
  │  │ DPU                       │       │
  │  │  ARM + NIC + accélérateurs│──→ réseau (400/800G)
  │  │  - vSwitch / OVS          │       │
  │  │  - NVMe-oF target         │       │
  │  │  - Firewall / TLS         │       │
  │  │  - Télémétrie             │       │
  │  └───────────────────────────┘       │
  └──────────────────────────────────────┘
```

**Idée** : à 400G, le traitement réseau « coûte » 8-16 cœurs CPU (⚠️).
Les mettre sur le DPU = les rendre au workload (ou ne pas les acheter).

## 65. BlueField-2 / BlueField-3 — l'existant

| Modèle | Réseau | CPU | PCIe | Mémoire | Statut |
|---|---|---|---|---|---|
| BlueField-2 | 2×100G ou 1×200G | 8× ARM A72 | Gen4 | 16-32 Go DDR4 | Legacy, encore déployé |
| BlueField-3 | **400G** (4/2/1 ports) | 16× ARM A78 | Gen5 | 64 Go DDR5 | **Standard actuel** ✅ |
| BlueField-3 SuperNIC | 2×400G ou 1×800G | 16× A78 + ConnectX-7 | Gen5 | 64 Go | Backend IA Spectrum-X ✅ |

**BlueField-3 en bref** : DOCA (SDK), offloads OVS (des millions de flux),
NVMe-oF, chiffrement, isolation multi-tenant (chaque VM ne voit que son
DPU). Conso carte : ~75 W (⚠️ ordre de grandeur) — à comparer aux 8-16
cœurs CPU économisés (~100-200 W, ⚠️).

## 66. BlueField-4 — la génération actuelle (vérifiée 27/09/2026)

✅ **Lancé à GTC Washington (octobre 2025)**, BlueField-4 = le DPU des
« AI factories » :

| Spec | Valeur (✅ vérifié 27/09/2026) |
|---|---|
| Débit | **800 Gb/s** (2× BF-3) |
| CPU | **NVIDIA Grace : 64 cœurs ARM Neoverse V2** |
| Réseau | **ConnectX-9** |
| Mémoire | 128 Go |
| PCIe | **Gen6** |
| Logiciel | DOCA microservices, zero-trust, service chaining natif |
| Dispo | **Early availability 2026**, plateforme Vera Rubin |

**Cas d'usage mis en avant** : stockage « AI-native » (KV cache pour
l'inférence agentique), sécurité runtime IA, multi-tenant networking.
Systèmes partenaires attendus **2ᵉ semestre 2026**.
**Lecture** : le DPU devient un **mini-serveur ARM** (64 cœurs !) — la
frontière NIC/serveur s'efface. En achat 2026 : BF-3 pour le 400G
immédiat, BF-4 à évaluer pour les plateformes Rubin/800G.

## 67. Offloads DPU : réseau, stockage, sécurité

| Famille | Offload | Effet |
|---|---|---|
| **Réseau** | OVS/vSwitch complet, NAT, LB | 0 cœur CPU pour le data-plane |
| **Réseau** | Télémétrie par flux, ECN/CC | Visibilité sans agent hôte |
| **Stockage** | NVMe-oF target/initiator, déduplication | Nœuds stockage sans CPU |
| **Stockage** | Compression/décompression, erasure coding | Moins de CPU par Go/s |
| **Sécurité** | Firewall stateful, TLS/IPsec, RegEx IDS | Micro-segmentation à 400G |
| **Sécurité** | Secure boot, attestation, isolation | Zero-trust hardware |

**Règle** : un offload ne vaut que s'il est **dans le chemin de données
réel**. Offloader le firewall quand le goulot est le stockage = 0 gain.
Mesurer d'abord (`perf`, compteurs NIC), offloader ensuite.

## 68. DOCA : le SDK qui fait (ou défait) le DPU

- **DOCA** = l'API NVIDIA pour programmer BlueField (data-plane, flow,
  crypto, telemetry). **DOCA 2.0** avec ConnectX-8 (✅).
- **Avantage** : écrire une fois, déployer sur BF-3/BF-4/ConnectX.
- **Risque** : lock-in logiciel — une équipe qui développe 50 k lignes de
  DOCA ne changera plus de DPU. **Évaluer le coût de sortie** avant
  d'investir (comme tout SDK propriétaire).
- **Alternative ouverte** : OPI (Open Programmable Infrastructure, Linux
  Foundation) — promet une API DPU multi-vendeurs ; maturité ⚠️ à vérifier
  au 27/09/2026 (non évaluée ici).

## 69. Quand mettre un DPU — matrice de décision

| Situation | DPU ? | Pourquoi |
|---|---|---|
| Cloud privé / multi-tenant (OpenStack, K8s) | **Oui** | Isolation, OVS offload, facturation |
| Nœuds stockage NVMe-oF | **Oui** | CPU rendu au stockage |
| Backend IA pur (NCCL) | **SuperNIC plutôt que DPU** | Le réordonnancement HW compte plus que l'ARM |
| Serveurs généralistes < 100G | **Non** | Surcoût injustifié |
| Edge / 5G (UPF, firewall) | **Oui** | PPS + accélérateurs |
| Pare-feu / sécurité inline 400G | **Oui** | TLS/RegEx en HW |

**Calcul de rentabilité** : DPU ~2 000-5 000 € (⚠️) vs 8-16 cœurs CPU
(~1 500-4 000 € en serveurs, ⚠️) + conso. Le DPU gagne quand le CPU est
cher (licences au cœur : VMware, Oracle…) ou quand l'isolation est
réglementaire.

