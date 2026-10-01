---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-4
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["gpu", "accelerator", "amd", "blackwell", "ethernet", "gpus", "hbm", "hbm3", "hbm4", "helios", "intel", "mi455x"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [419, 537]
sha256: bc5898232fa2c13c68a9442f6bc190c7fd9c53286fd04328b008c070438faf7a
---

# Les GPU datacenter / IA — présent et futur vérifié

| Génération | Débit par GPU | GPUs | Année |
|---|---|---|---|
| NVLink 1 | 160 Go/s | P100 | 2016 |
| NVLink 2 | 300 Go/s | V100 | 2017 |
| NVLink 3 | 600 Go/s | A100 | 2020 |
| NVLink 4 | 900 Go/s | H100, H200 | 2022–2024 |
| NVLink 5 | 1 800 Go/s (1,8 To/s) | B200, B300 | 2024–2026 |
| NVLink 6 | 3 600 Go/s (3,6 To/s) par GPU | Rubin | 2026 (Vera Rubin NVL72) |

Le NVLink 6 de Rubin porte le scale-up du rack NVL72 à **260 To/s** totaux
(données NVIDIA, GTC 2026).

### 18.2. NVSwitch

- **Rôle** : commutateur tout-à-tout entre les 8 GPU d'un nœud. Sans NVSwitch,
  chaque GPU ne parle qu'à ses voisins (topologie en anneau/mesh partiel).
- **Générations** : NVSwitch 1 (V100), 2 (A100), **3 (H100/H200, 900 Go/s)**,
  **4 (B200)**. Le B300/GB300 utilise 6× NVSwitch à 4,8 To/s agrégés.
- **Conséquence pratique** : un serveur 8× H100 **avec** NVSwitch se comporte comme
  un seul GPU de 640 Go pour le training distribué (all-reduce à pleine vitesse).
  Un serveur 8× PCIe **sans** NVSwitch (ex. 8× L40S) fait du training multi-GPU
  3 à 10× moins efficace à taille de modèle égale.

### 18.3. Schéma ASCII : topologie 8× GPU avec NVSwitch

```
        ┌─────────────────────────────────────────────┐
        │              NVSwitch (tout-à-tout)          │
        │   900 Go/s par lien (NVLink 4, H100/H200)   │
        └──┬──────┬──────┬──────┬──────┬──────┬───┬────┘
           │      │      │      │      │      │   │
         GPU0   GPU1   GPU2   GPU3   GPU4   GPU5 GPU6 GPU7
           │      │      │      │      │      │   │
        ┌──┴──────┴──────┴──────┴──────┴──────┴───┴────┐
        │         PCIe 5.0 x16 vers CPU hôte         │
        │   (128 Go/s — le goulot vers le host)      │
        └─────────────────────────────────────────────┘
```

### 18.4. Quand NVLink change tout (et quand il ne sert à rien)

- **Training distribué** : indispensable. All-reduce des gradients à chaque step.
- **Inférence tensor-parallel sur plusieurs GPU** : indispensable au-delà d'1 GPU.
- **Inférence 1 GPU / batch** : inutile — le modèle tient en HBM, rien ne transite.
- **Data-parallel simple** (N requêtes indépendantes sur N GPU) : inutile.

## 19. Infinity Fabric et xGMI côté AMD

- **xGMI** (inter-GPU) : l'équivalent AMD du NVLink pour les liens directs GPU↔GPU.
  Sur MI300X : 7 liens xGMI par GPU vers les 7 autres (topologie tout-à-tout en
  plateforme 8×), ~896 Go/s agrégés par GPU (ordre de grandeur public).
- **Infinity Fabric** (intra-package) : relie les 8 XCD aux 4 IOD dans le MI300X
  via empilement 3D SoIC — ~0,05 pJ/bit, bien meilleur que l'interposeur classique.
- **Inter-plateforme** : AMD pousse **UALink** (Ultra Accelerator Link, standard ouvert)
  pour le scale-up rack : 260 To/s agrégés annoncés sur Helios (MI450).
- **RCCL** : l'équivalent AMD de NCCL (librairie de collectives). Maturité 2025-2026
  en net progrès, mais toujours un cran derrière NCCL sur les topologies exotiques.
- **Point de vigilance** : le duo xGMI+RCCL fonctionne très bien en 8× homogène AMD.
  Le **mixte** (ex. 4× AMD + 4× NVIDIA, ou multi-nœuds hétérogènes) reste un terrain
  miné — à éviter en production.

## 20. Serveurs 4/8 GPU : HGX, DGX et équivalents AMD

### 20.1. Gamme NVIDIA (configs actuelles vérifiées le 27/09/2026)

| Système | GPU | Mémoire totale | TDP GPU totaux | Refroidissement | Prix indicatif |
|---|---|---|---|---|---|
| DGX H100 | 8× H100 SXM | 640 Go | 5,6 kW | Air | ~290 000 $ |
| OEM 8× H100 HGX | 8× H100 SXM | 640 Go | 5,6 kW | Air | 250 000–320 000 $ |
| OEM 4× H100 PCIe | 4× H100 PCIe | 320–384 Go | 1,4–2,8 kW | Air | 150 000–180 000 $ |
| DGX H200 | 8× H200 | ~1,1 To | 5,6 kW | Air | ~350 000 $ (est.) |
| OEM 8× H200 HGX | 8× H200 | ~1,1 To | 5,6 kW | Air | 320 000–420 000 $ |
| DGX B200 | 8× B200 | ~1,5 To | 8 kW | Air/liquide | ~515 000 $ |
| OEM 8× B200 HGX | 8× B200 | ~1,5 To | 8 kW | Liquide recommandé | 400 000–520 000 $ |
| GB200 NVL72 (rack) | 72× B200 + 36 Grace | ~14 To | ~120 kW | Liquide | ~3 000 000 $ |
| DGX B300 | 8× B300 | ~2,3 To | 11,2 kW | Liquide | n.c. |
| 4× RTX PRO 6000 (2U MGX) | 4× RTX PRO 6000 | 384 Go | 2,4 kW | Air | ~80 000–100 000 $ (est.) |

### 20.2. Équivalents AMD (plateformes 8× OAM)

- **Dell PowerEdge XE9680** : 8× MI300X (jusqu'à 1,5 To HBM3), air, ~10 kW.
- **Supermicro AS-8125GS-TNMR2** : 8× MI300X/MI325X, air/liquide.
- **HPE / Lenovo / Gigabyte** : plateformes 8× MI300X/MI350X équivalentes (UBB AMD).
- **Rack Helios** (MI450, H2 2026) : 72× MI455X, 31 To HBM4, design de référence
  Schneider Electric à **246 kW** (vérifié le 27/09/2026).

### 20.3. Schéma ASCII : serveur 8 GPU type HGX

```
 ┌──────────────────────────────────────────────────────────────┐
 │  Châssis 8U — Serveur HGX 8× GPU                             │
 │                                                              │
 │  ┌────────────┐  ┌────────────┐      ┌──────────────────┐     │
 │  │ 2× CPU     │  │ 8× GPU sur │      │ 8× ConnectX-7/8  │     │
 │  │ EPYC/Xeon  │  │ baseboard  │◄────►│ 400G/800G IB ou  │     │
 │  │ 2 To DDR5  │  │ + NVSwitch │ PCIe │ Ethernet         │     │
 │  └────────────┘  └────────────┘      └──────────────────┘     │
 │                                                              │
 │  Alimentation : 6+2 PSU 3000 W  (ex. 8× B200 → 6× 3 kW mini)  │
 │  Refroidissement : 8–12 ventilateurs 80 mm / plaques DLC      │
 └──────────────────────────────────────────────────────────────┘
```

## 21. DLSS / FSR / XeSS — section courte

**Principe commun** : rendre l'image à résolution réduite, puis reconstruire à
résolution native par IA (upscaling) et/ou générer des images intermédiaires
(frame generation). But : plus de FPS à qualité perçue égale. **C'est du gaming.**

| Techno | Acteur | Depuis | Hardware requis | Frame Gen | Statut 27/09/2026 |
|---|---|---|---|---|---|
| DLSS 2 (SR) | NVIDIA | 2020 | RTX 20+ | Non | Modèle CNN/transformer, fermé |
| DLSS 3 (FG) | NVIDIA | 2022 | RTX 40 (Ada) | 1 image / rendue | + Reflex anti-latence |
| DLSS 4 (MFG) | NVIDIA | 2025 | RTX 50 (Blackwell) | **jusqu'à 3 images** | Modèle transformer |
| DLSS 5 | NVIDIA | 2026 | RTX 50 | Oui | Neural rendering 3D, RTX 40 exclu |
| FSR 1–3 | AMD | 2021–2023 | Tous GPU (ouvert) | FSR 3 : oui | Spatial → temporel analytique |
| FSR 4 / Redstone | AMD | 2025 | RDNA 4 (RX 9000) | Oui | **ML-based** (rattrape DLSS) |
| XeSS 1 / 2 / 3 | Intel | 2022–2026 | Arc (XMX) ou DP4a | XeSS 2+ | MFG chez Intel aussi |

