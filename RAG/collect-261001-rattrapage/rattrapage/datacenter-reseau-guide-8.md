---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-8
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Meta", "Microsoft", "Nvidia", "xAI"]
dates: ["2025-10-08", "2025-10-14", "2026-09-27"]
keywords: ["amd", "asic", "blackwell", "cpo", "distribution", "dsp", "ethernet", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1013, 1168]
sha256: 086ef594288eb8197d4a31eee4a26ea95e170795ec4c217a484c2f0ece944832
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| ASIC | Capacité | Positionnement | Statut |
|---|---|---|---|
| Tomahawk 4 | 25,6T | Génération 400G | Base de grands fabrics (Meta) |
| Tomahawk 5 | 51,2T | 64×800G / 128×400G | **Shipping** ✅ |
| Tomahawk 6 | **102,4T** | 64×1,6T / 128×800G / 512×200G ; Cognitive Routing 2.0 | **Annoncé** ✅ |
| Tomahawk 6-Davisson | 102,4T **CPO** | Optiques co-packagées, 3,5 W/port à 800G | Early access ✅ (08/10/2025) |
| Tomahawk Ultra | UEC | AllReduce **dans le switch** (in-network collectives) | ✅ |
| Jericho3 / 3-AI | Deep-buffer | Spine IA (Arista 7800R4) | Shipping ✅ |

**Taurus** (✅ OFC 2026, mars 2026) : premier DSP optique **400G/lane** —
permet des modules 1,6T et prépare le 3,2T / les switchs 204,8T.
**Roadmap affichée par Broadcom : l'ère 200T.**

**Lecture** : Broadcom vend le silicium, pas (peu) de switchs : Arista,
Dell, HPE, Edgecore, Celestica, H3C construisent dessus. **Comparer les
implémentations logicielles** (EOS, SONiC durci, etc.), pas l'ASIC —
il est souvent identique.

## 51. Arista Etherlink — l'Ethernet IA « enterprise »

| Produit | Base | Spec (✅ vérifié 27/09/2026) |
|---|---|---|
| 7060X6 (leaf) | Tomahawk 5 | 64×800G OSFP, 51,2T, 2× PSU 2400 W |
| 7060X5 | Tomahawk 4 | 32×800G, 2× PSU 1500 W |
| 7800R4 (spine modulaire) | Jericho3-AI | Jusqu'à **460 Tb/s**, 576×800G ou 1152×400G par châssis ; −65 % de conso vs 7800R3 |

**Positionnement** : EOS (le NOS le plus mature du marché enterprise),
CloudVision, AI Etherlink (features IA : monitoring par flux, tuning RoCE).
C'est le choix « sans surprise » pour une équipe qui connaît déjà Arista —
et celui de Meta pour son cluster RoCE 24k H100.

## 52. Cisco, Juniper, HPE, whitebox — le reste du paysage

| Acteur | Offre IA | Note |
|---|---|---|
| Cisco Silicon One G200 | 51,2T, 800G | ⚠️ Détails non vérifiés au 27/09/2026 — vérifier fiche Cisco |
| Juniper (HPE) | PTX / QFX sur Broadcom | Solide SP, moins « IA-first » |
| HPE Aruba | CX sur merchant silicon | Enterprise, pas le backend IA |
| Edgecore / Celestica | Whitebox TH5/TH6 + SONiC | Le moins cher au port (⚠️), support à évaluer |
| Nexthop AI | TH6-Davisson CPO + SONiC durci | ✅ cité par Broadcom (08/10/2025) |

**SONiC** (open-source, Microsoft) : le NOS des whitebox. Mature en
2026 pour le cloud, **à valider** pour les features IA fines (CC, telemetry)
— prendre une distribution durcie (Nexthop, Aviz) plutôt que du SONiC brut
en production IA.

## 53. Dimensionner un fabric IA — méthode en 6 étapes

1. **Compter les endpoints** : N GPU × NIC/GPU (1:1 standard).
2. **Choisir le ratio** : 1:1 backend, 3:1 front-end.
3. **Choisir le radix** : 64×800G → 1024 endpoints en 2 tiers (§33).
4. **Séparer les plans** : rail-optimized, 1 plan/indice NIC (§39).
5. **Vérifier le BDP** : buffers spine ≥ BDP × flux concurrents (§44).
6. **Chiffrer l'énergie** : switch + optiques + NIC (§89-92) — prévoir
   **2× la conso switch** pour la clim (PUE ~1,5-2 sur la partie réseau).

## 54. Ultra Ethernet (UEC) — le standard qui change la donne

- **11 juin 2025** : UEC Specification **1.0** publiée (transport, contrôle
  de congestion, RDMA sur Ethernet, > 100 membres : AMD, Arista, Broadcom,
  Cisco, HPE, Intel, Meta, Microsoft…) ✅.
- **Mi-2025** : spec **1.0.1** ✅. **14/10/2025** : premier NIC conforme,
  Broadcom **Thor Ultra** 800G ✅.
- **Innovations** : UET (Ultra Ethernet Transport) — livraison **non
  ordonnée**, retransmission **sélective**, multipath natif, profils
  applicatifs (collectives IA), pas de lossless requis.
- **Juin 2026** : VIAVI sort la première solution de validation UET ✅
  (TestCenter) — signe de maturité industrielle.
- **Position NVIDIA** : pas membre pilote d'UEC ; Spectrum-X = alternative
  propriétaire. Le marché se divise en **converged** (NVIDIA, boucle fermée)
  vs **discrete** (Broadcom+OEM+UEC, ouvert).

**Impact achat** : exiger « UEC-ready » sur les appels d'offres 2026+ pour
éviter le lock-in, même si on déploie Spectrum-X aujourd'hui — c'est une
assurance interopérabilité.

## 55. CAS CHIFFRÉ — fabric backend pour 1024 GPU (Blackwell)

Hypothèses : 128 nœuds × 8 GPU, 1× NIC 800G par GPU (1:1), rail-optimized
8 plans, 2 tiers.

| Poste | Calcul | Résultat |
|---|---|---|
| NIC 800G | 1024 × ConnectX-8 | 1024 NIC |
| Par plan | 1024 / 8 = 128 endpoints/plan | 2 leafs 64×800G par plan |
| Leafs | 8 plans × 2 = **16 leafs** (64×800G) | 1:1 par plan |
| Spines | 8 plans × 2 spines 64×800G = **16 spines** | — |
| Switchs totaux | 32 × 64×800G | — |
| Conso switchs (⚠️) | 32 × ~1,5 kW (avec optiques) | **~48 kW** |
| Conso NIC 800G (⚠️) | 1024 × ~60 W | **~61 kW** |
| Optiques (⚠️) | ~3000 × ~15 W | **~45 kW** |
| **Total réseau IT** | | **~154 kW** |
| Avec PUE 1,4 (⚠️) | 154 × 1,4 | **~216 kW** tirés au compteur |

**Référence** : 1024 GPU B200 ≈ 1,2-1,5 MW IT (⚠️) → le réseau = **~10-13 %**
de la puissance du cluster. À 0,15 €/kWh (⚠️), le réseau coûte
~280 k€/an d'électricité. Le « petit » poste réseau n'est pas petit.

## 56. Pièges switches IA — l'essentiel (détail §107)

- **P16** — PFC global au lieu de PFC par file : pauses storms.
- **P17** — Leaf shallow-buffer en spine : drops sous incast.
- **P18** — RoCE sans ECN/DCQCN : 60 % du débit (cas xAI : 60 → 95 % avec CC).
- **P19** — Comparer des switchs sur l'ASIC seul : le logiciel fait la différence.

# PARTIE E — CHÂSSIS 800GbE POUR CLUSTERS GPU IA

## 57. Ce qu'implique le 800GbE : radix, ports, câblage

Le 800G change trois choses fondamentalement :

1. **Radix** : un switch 64×800G = 51,2T dans 2U. Le même débit tenait en
   8 switchs 64×400G il y a 2 ans. **Moins de switchs, mais chacun est un
   point de défaillance plus gros** — la redondance N+1 devient critique.
2. **Densité énergétique** : 64 ports × 18 W (optiques) = **1,15 kW rien que
   d'optiques** + ~0,9 kW de switch = 2 kW dans 2U. C'est une densité de
   rack GPU, pas de rack réseau classique.
3. **Câblage** : le DAC 800G ne dépasse pas ~2 m (⚠️) — **tout ce qui n'est
   pas intra-rack passe en AOC ou optique**. Le budget câblage explose si
   la topologie n'est pas compacte (leafs au plus près des nœuds).

## 58. Radix et dimensionnement Clos au 800G

Formule (rappel §33) : switchs n ports → **n²/4 endpoints** non-bloquants
en 2 tiers, n/2 spines.

| Switch | Endpoints 2 tiers | Spines | Cas d'usage |
|---|---|---|---|
| 32×800G (25,6T) | 256 | 16 | Pod 256 GPU |
| 64×800G (51,2T) | 1024 | 32 | Cluster 1k GPU |
| 64×800G en 3 tiers | 16k+ | pods de 1024 | 10k+ GPU |

**Exemple pod 256 GPU** : 8 leafs 32×800G (16 down + 16 up) + 16 spines
32×800G → 256 endpoints 1:1. Seulement **24 switchs** pour 256 GPU :
la densité 800G simplifie la topologie — mais chaque spine porte 1/16ᵉ du
trafic : **prévoir la perte d'un spine sans dégradation** (surprovisionner
de ~10 % ou accepter 94 % en dégradé).

## 59. Câblage 800G : la matrice de choix

| Distance | Solution | Conso/extr. (⚠️) | Note |
|---|---|---|---|
| < 2 m | DAC 800G passif | ~0,5 W | Intra-rack uniquement |
| 2-6 m | AEC 800G (cuivre actif) | ~3-4 W | Racks adjacents |
| 3-30 m | AOC 800G | ~2-3 W | Inter-racks proches |
| < 100 m | 800G-SR8 (OM4, MPO-16) | ~12-14 W | Inter-rangées |
| < 500 m | 800G-DR8 (OS2, MPO-16) | ~14-16 W | Inter-salles |
| 2 km | 800G-2×FR4 (OS2, CS) | ~16-18 W | Campus |

**Règle thermique** : à 64 ports optiques, le switch dissipe ~2 kW —
**vérifier le sens d'airflow** (avant-arrière vs arrière-avant) par rapport
aux allées chaudes/froides. Un switch 800G monté à l'envers = hotspot de
20 kW par rangée (P20, §107).

## 60. Thermique OSFP : prévoir la chaleur, pas la subir

