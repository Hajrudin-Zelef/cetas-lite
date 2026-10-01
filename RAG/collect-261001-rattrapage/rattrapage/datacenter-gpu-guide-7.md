---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-7
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "amd", "blackwell", "compute", "distribution", "ethernet", "kv cache", "nvidia", "nvlink", "rubin", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [844, 1005]
sha256: 9d211b188fa4ea227327b6282eda6cc2fcff3d50c524a107ef8099649de96ef4
---

# Les GPU datacenter / IA — présent et futur vérifié

Une carte OEM Dell dans un serveur HPE (ou l'inverse) : le vBIOS peut refuser de
booter, brider la ventilation (courbes OEM), ou désactiver MIG. Les cartes
« blanches » sans firmware constructeur sont un classique des lots gris. Exiger le
firmware du constructeur du **serveur**, pas de la carte.

## 31. Pièges terrain (2/3)

### Piège n°7 : contrefaçons et cartes « rebadgées »

Le marché gris regorge de RTX 3090 rebadgées en 4090 (vBIOS modifié) et de Tesla
P40 maquillées en cartes récentes. Vérifications : `nvidia-smi -q` (Device ID, UUID),
poids et dimensions vs spec officielle, prix « trop beau » (< 60 % du marché = fuyez).
Acheter via distribution agréée (PNY, OEM) pour toute prod.

### Piège n°8 : garantie gaming vs datacenter

Faire tourner des 4090/5090 en 24/7 dans un datacenter : **garantie constructeur
non applicable** (usage non conforme aux conditions grand public chez la plupart
des AIB). En cas de RMA, le fabricant peut refuser. Pour de la prod : cartes pro
(L40S, RTX PRO, H100+) avec garantie 3 ans on-site constructeur.

### Piège n°9 : le NVLink bridge oublié (H200 NVL)

Les H200 NVL se vendent par paire avec un **bridge NVLink** dédié. Commander 4 GPU
sans les bridges = 4 GPU isolés en PCIe. Le bridge est une référence séparée sur
le devis — la vérifier ligne par ligne.

### Piège n°10 : refroidissement passif sans flux d'air

L40S passive, RTX PRO 6000 Server Edition passive : elles **n'ont pas de ventilateurs**.
Elles comptent sur le flux d'air frontal du châssis (ex. 8× ventilateurs 80 mm à
plein régime + allée froide ≤ 27 °C). Les monter dans un boîtier tour ou un rack
sans confinement = surchauffe en 10 minutes et throttling à 1 500 MHz.

### Piège n°11 : MIG mal dimensionné

La RTX PRO 6000 accepte 4 instances MIG (4× 24 Go). Mais une instance MIG 24 Go ne
fait pas tourner un 70B Q4 + KV cache confortable : le découpage MIG est **rigide**
(tailles fixes). Tester le profil exact (`nvidia-smi mig -lgip`) avant de promettre
4× clients par carte.

### Piège n°12 : ROCm — la version exacte compte

Sur AMD, PyTorch 2.4 + ROCm 6.2 + gfx942 (MI300X) : OK. PyTorch compilé pour gfx90a
(MI200) sur MI300X : **ne fonctionne pas** — le target gfx doit matcher. Toujours
vérifier `rocminfo` et la matrice de compatibilité ROCm avant de déployer. Les
images Docker `rocm/pytorch` officielles règlent 90 % des cas.

## 32. Pièges terrain (3/3)

### Piège n°13 : CUDA compute capability et PyTorch

H100/H200 = sm_90, B200 = sm_100/sm_100a, RTX 4090 = sm_89, RTX 5090 = sm_120.
Un binaire PyTorch compilé sans le bon SM tourne en **JIT PTX** (lent au premier
lancement) ou refuse de tourner. En conteneur : prendre les images NGC taguées
pour l'architecture exacte (`nvcr.io/nvidia/pytorch:xx.xx-py3`).

### Piège n°14 : NUMA et affinité CPU-GPU

Sur serveur bi-CPU, un GPU accroché au CPU 0 qui lit la RAM du CPU 1 traverse
l'UPI/Infinity Fabric : -30 % de bande passante host→device. Fixer l'affinité
(`numactl`, `CUDA_VISIBLE_DEVICES` + placement des dataloaders). `nvidia-smi topo -m`
montre la matrice exacte : **la lire avant d'optimiser**.

### Piège n°15 : l'alimentation mono-phase pour 4 GPU

4× RTX 5090 = 2,3 kW GPU + système ≈ 3 kW au mur → **13 A sur 230 V mono**.
Ça passe sur une prise 16 A dédiée, mais pas sur une multiprise de bureau avec
3 autres charges. En rack : PDU metered par phase, équilibrage des phases,
et jamais plus de 80 % de charge continue par disjoncteur (règle NEC/CEI).

### Piège n°16 : la poussière, tueuse de datacenter GPU

Un filtre à air saturé sur un nœud 8× H100 : +15 °C en entrée GPU en quelques
semaines → throttling → -20 % de perf + RMA prématurés. Programme de maintenance :
contrôle mensuel des filtres, thermographie semestrielle des points chauds
(connecteurs 16-pin, VRM).

### Piège n°17 : acheter du H100 en 2026 sans négocier

Le H100 est en fin de cycle (B200/B300 disponibles, Rubin en production). Les
stocks OEM se négocient : viser **-15 à -25 %** sur les prix catalogue 2024.
À l'inverse, ne pas acheter du H100 « pour le futur » : pour un déploiement neuf
fin 2026, le B200 ou le H200 ont un meilleur TCO perf/watt.

### Piège n°18 : oublier le réseau dans le BOM

Un nœud 8× B200 sans InfiniBand 800G : le training multi-nœuds s'effondre
(le scale-out devient le goulot). Compter **8× NIC 800G + 2 switches IB par nœud**
soit ~30 000–50 000 $ de réseau par nœud — 10 % du BOM, 100 % du scaling.

## 33. Où acheter : circuits d'achat

### 33.1. Cartes à l'unité (gaming / pro PCIe)

| Canal | Produits | Délai | Note |
|---|---|---|---|
| Retailers (Amazon, Newegg, LDLC) | RTX 4090/5090 | Immédiat | Prix volatils, garantie grand public |
| Distribution pro (PNY, Tech Data, Ingram) | L40S, RTX PRO 6000 | 1–4 sem. | Devis, garantie 3 ans |
| Marché secondaire (eBay, Jawa) | 4090/5090, L40S recond. | Immédiat | Vérifier piège n°7 |

### 33.2. Serveurs intégrés (H100/H200/B200/MI300X…)

| OEM | Gammes GPU | Positionnement |
|---|---|---|
| Dell | H100/H200/B200, MI300X | PowerEdge XE9680/XE9712, support ProSupport |
| HPE | H100/H200/B200, MI300X | ProLiant Compute XD, Cray pour HPC |
| Supermicro | Tout (le plus large) | SYS-821GE (HGX), prix agressifs, délais courts |
| Lenovo | H100/H200/B200 | ThinkSystem SR675 V3 |
| Gigabyte | H100/H200, MI300X | G593 (HGX 8U), bon rapport qualité/prix |
| NVIDIA DGX | H100/H200/B200 | Clé en main + Base Command, premium ~10 % |

**Règle d'achat** : < 4 GPU → cartes PCIe à l'unité ou serveurs 4U généralistes.
≥ 8 GPU → plateforme HGX/UBB intégrée, **jamais** d'assemblage artisanal (support,
thermique et NVSwitch ne s'improvisent pas).

## 34. BOM type : serveur 4× H200 NVL PCIe (inférence 70B–405B)

| Composant | Référence type | Qté | Prix indicatif (USD) |
|---|---|---|---|
| GPU | NVIDIA H200 NVL 141 Go + NVLink bridges | 4 | 140 000–165 000 |
| CPU | 2× AMD EPYC 9454 (48c) | 2 | 8 000 |
| Carte mère | Supermicro H13DSG-OM | 1 | 1 500 |
| RAM | 1 To DDR5-4800 ECC RDIMM | 1 lot | 4 800 |
| Stockage | 8 To NVMe U.2 RAID | 1 lot | 1 000 |
| Réseau | 2× ConnectX-7 400G | 2 | 3 000 |
| Alimentation | 4× 3 000 W redondantes | 1 lot | 1 500 |
| Châssis | 4U/5U rack, flux d'air GPU | 1 | 500–1 000 |
| **Total** | | | **~160 000–185 000** |

Prix vérifiés le 27/09/2026 via BOM publique (kinvert/knowledge-base). Le GPU
représente ~85 % du total — normal.

## 35. BOM type : serveur 8× RTX PRO 6000 Blackwell (inférence économique)

| Composant | Référence type | Qté | Prix indicatif (USD) |
|---|---|---|---|
| GPU | RTX PRO 6000 Blackwell Server Edition 96 Go | 8 | 88 000–104 000 |
| CPU | 2× AMD EPYC 9354 (32c) | 2 | 5 000 |
| Carte mère | H13 8× PCIe 5.0 x16 | 1 | 2 000 |
| RAM | 768 Go DDR5 ECC | 1 lot | 3 500 |
| Stockage | 4 To NVMe | 1 lot | 500 |
| Réseau | 2× ConnectX-7 200G | 2 | 2 000 |
| Alimentation | 4× 3 000 W (4,8 kW GPU + système) | 1 lot | 1 500 |
| Châssis | 4U 8× DW PCIe 5.0 | 1 | 1 500 |
| **Total** | | | **~105 000–120 000** |

768 Go VRAM cumulés pour ~110 k$ : imbattable en $/Go, mais **sans NVLink** —
réservé à l'inférence data-parallel ou au sharding pipeline.

## 36. BOM type : plateforme AMD 8× MI350X (alternative B200)

| Composant | Référence type | Qté | Prix indicatif |
|---|---|---|---|
| GPU | 8× MI350X OAM sur UBB AMD | 1 plateforme | À vérifier (pas de prix public) |
| CPU | 2× EPYC Turin 64c | 2 | ~8 000 $ |
| RAM | 1,5 To DDR5 ECC | 1 lot | ~7 000 $ |
| Réseau | 8× Pensando/800G Ethernet (UEC) | 8 | ~15 000 $ (est.) |
| Alimentation | 6× 3 000 W | 1 lot | ~2 000 $ |
| Refroidissement | Air haute densité ou DLC | 1 | +10–30 % vs air |
| **Total (hors GPU)** | | | **~35 000–45 000 $** |

