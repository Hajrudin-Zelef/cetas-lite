---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-3
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "accelerator", "amd", "blackwell", "compute", "embeddings", "fine-tuning", "fp4", "fp8", "gpus", "hbm"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [279, 418]
sha256: 90151f1cc69f28683d228f1acb893ead30b149fc0a72399581d978617d8d83f5
---

# Les GPU datacenter / IA — présent et futur vérifié

- **Référence exacte** : AMD Instinct MI300A (APU), CDNA 3 + Zen 4 sur le même package
  (vérifiée le 27/09/2026).
- **Architecture** : APU : 228 CU GPU + cœurs CPU Zen 4, mémoire **unifiée**.
- **Mémoire** : **128 Go HBM3 unifiés** CPU+GPU, **5,3 To/s**. Zéro copie hôte/device.
- **TDP** : **760 W**.
- **Précisions IA** : FP8 FNUZ comme MI300X (réduit par le nombre de CU).
- **Cas d'usage** : **HPC** (supercalculateurs : El Capitan au LLNL), charges mémoire-unifiée.
  Moins pertinent pour l'inférence LLM pure que le MI300X.
- **Prix** : ~10 000–15 000 $ estimé (pas de prix public unitaire, nœuds 4 APU).
- **Où l'acheter** : intégrateurs HPC, HPE Cray. Pas de carte PCIe standard.

---
## 13. Tableau comparatif : mémoire

| GPU | Capacité | Type | Bande passante | ECC | Go / 1 000 $ (indicatif) |
|---|---|---|---|---|---|
| RTX 4090 | 24 Go | GDDR6X | 1 008 Go/s | Non | ~10 (2 400 $) |
| RTX 5090 | 32 Go | GDDR7 | 1 792 Go/s | Non | ~12 (2 600 $) |
| L40S | 48 Go | GDDR6 | 864 Go/s | Oui | ~4,8 (10 000 $) |
| RTX PRO 6000 Blackwell | 96 Go | GDDR7 | 1 792 Go/s | Oui | ~8 (12 000 $) |
| H100 SXM | 80 Go | HBM3 | 3 350 Go/s | Oui | ~2,7 (30 000 $) |
| H100 NVL | 188 Go | HBM3 | ~7,8 To/s | Oui | ~3,1 (60 000 $ est.) |
| H200 | 141 Go | HBM3e | 4 800 Go/s | Oui | ~4 (35 000 $) |
| B200 | 192 Go | HBM3e | 8 000 Go/s | Oui | ~4,8 (40 000 $) |
| MI300X | 192 Go | HBM3 | 5 300 Go/s | Oui | n.c. (plateforme) |
| MI325X | 256 Go | HBM3e | 6 000 Go/s | Oui | n.c. (plateforme) |
| MI350X / MI355X | 288 Go | HBM3e | 8 000 Go/s | Oui | n.c. (plateforme) |
| MI300A | 128 Go | HBM3 unifiée | 5 300 Go/s | Oui | n.c. (HPC) |

Lecture : pour l'inférence LLM, la bande passante mémoire gouverne le débit (tokens/s).
Pour le training, c'est le produit bande passante × compute × interconnexion qui compte.
Le ratio Go/1 000 $ montre que les cartes gaming écrasent tout… sans ECC ni support datacenter.

## 14. Tableau comparatif : TDP / performance

| GPU | TDP | FP8 dense | FP4 dense | Perf/Watt FP8 (TFLOPS/W) |
|---|---|---|---|---|
| RTX 4090 | 450 W | ~330 TFLOPS (est.) | — | ~0,73 |
| RTX 5090 | 575 W | 419 TFLOPS | 3 352 TOPS (sparse) | ~0,73 |
| L40S | 350 W | 733 TFLOPS | — | ~2,09 |
| RTX PRO 6000 | 600 W | ~800 TFLOPS | ~2 000 (dense, est.) | ~1,33 |
| H100 | 700 W | 1 979 TFLOPS | — | ~2,83 |
| H200 | 700 W | 1 979 TFLOPS | — | ~2,83 |
| B200 | 1 000 W | 4 500 TFLOPS | 9 000 TFLOPS | ~4,50 |
| MI300X | 750 W | 2 615 TFLOPS | — | ~3,49 |
| MI325X | 1 000 W | 2 615 TFLOPS | — | ~2,62 |
| MI350X | 1 000 W | ~4 600 TFLOPS (est.) | 9 200 TFLOPS | ~4,60 |
| MI355X | 1 400 W | ~5 000 TFLOPS (est.) | 10 100 TFLOPS | ~3,57 |

Note : les chiffres FP8/FP4 AMD CDNA 4 sont les pics constructeur (MXFP4/MXFP6/MXFP8).
Le B200 reste le champion perf/watt FP8 ; le MI350X le talonne en air à 1 000 W.

## 15. Tableau comparatif : prix / performance

Prix indicatifs vérifiés le 27/09/2026 (achat carte nue quand possible, sinon estimation
intégrée /8 pour les serveurs). Coût par TFLOPS FP8 :

| GPU | Prix indicatif (USD) | FP8 dense | $ / PFLOPS FP8 | $ / Go VRAM |
|---|---|---|---|---|
| RTX 4090 | ~2 400 (rue) | ~0,33 PFLOPS | ~7 300 | ~100 |
| RTX 5090 | ~2 600 (rue) | 0,42 PFLOPS | ~6 200 | ~81 |
| L40S | ~10 000 | 0,73 PFLOPS | ~13 700 | ~208 |
| RTX PRO 6000 | ~12 000 | ~0,80 PFLOPS | ~15 000 | ~125 |
| H100 | ~30 000 | 1,98 PFLOPS | ~15 200 | ~375 |
| H200 | ~35 000 | 1,98 PFLOPS | ~17 700 | ~248 |
| B200 | ~40 000 | 4,50 PFLOPS | ~8 900 | ~208 |
| MI300X | n.c. | 2,61 PFLOPS | n.c. | n.c. |
| MI355X | n.c. | ~5,0 PFLOPS | n.c. | n.c. |

Lecture brutale : le B200 est le meilleur $/PFLOPS du datacenter ; la 5090 le meilleur
$/PFLOPS tout court (mais sans ECC, sans support, sans NVLink). Le H200 se paie cher
pour sa mémoire, pas son compute. Ces ratios ignorent l'énergie, le refroidissement
et le logiciel — voir sections 26-29 et 37.

## 16. Tableau comparatif par cas d'usage

| Cas d'usage | 1er choix | 2e choix | À éviter | Pourquoi |
|---|---|---|---|---|
| Inférence 7B–13B locale | RTX 4090/5090 | L40S | H100 | Le H100 est surdimensionné et 10× plus cher |
| Inférence 70B FP16 1 GPU | H200 (141 Go) | MI300X (192 Go) | 2× RTX PRO 6000 | 140 Go tiennent sur 1 GPU, pas de sharding |
| Inférence 70B quantifié 1 GPU | RTX PRO 6000 (96 Go) | 2× 5090 | L40S | Q4 70B ≈ 40 Go + KV cache |
| Training 7B–70B | H100 8× | B200 8× | Cartes gaming | NVLink + ECC + NCCL |
| Training > 100B | B200 8× / GB200 | MI355X 8× | Tout PCIe | Bande passante inter-GPU critique |
| Fine-tuning LoRA/QLoRA | RTX 5090 / 4090 | RTX PRO 6000 | H100 | Coût, simplicité |
| RAG / embeddings | L40S | RTX PRO 6000 | B200 | Débit/$ imbattable du L40S |
| HPC classique (FP64) | MI300A / H100 | MI430X (à venir) | RTX PRO 6000 | FP64 quasi absent des RTX |
| Lab / dev | RTX 5090 | RTX 4090 | — | Prix, dispo immédiate |
| Inférence $/token à l'échelle | B200 | MI355X | H100 | FP4 + 8 To/s |

## 17. Formats physiques : SXM vs PCIe vs OAM/UBB

### 17.1. SXM (NVIDIA)

- **Principe** : module mezzanine soudé sur une **baseboard** (HGX), pas de connecteur PCIe
  pour les données GPU↔GPU (le PCIe ne sert qu'au host). Refroidissement par plaque froide
  ou dissipateur passif traversé par le flux d'air du châssis.
- **Générations** : SXM2 (V100), SXM3 (A100), SXM4 (A100 80 Go), **SXM5 (H100/H200)**.
  Le B200 utilise un format SXM dérivé Blackwell (même philosophie, TDP 1 000 W).
- **Avantages** : TDP élevé possible (700–1 000 W), NVLink natif entre GPU via NVSwitch,
  densité maximale (8 GPU en 8U).
- **Inconvénients** : **non interchangeable** — un H100 SXM ne va pas dans un slot PCIe ;
  achat = serveur complet ; maintenance = remplacement de baseboard.

### 17.2. PCIe (tout le monde)

- **Principe** : carte standard FHFL (full-height, full-length) ou double-slot, slot
  PCIe x16. Le GPU parle au CPU via PCIe, aux autres GPU via NVLink bridge (H200 NVL)
  ou pas du tout (L40S, RTX PRO 6000, 4090/5090).
- **TDP plafond pratique** : ~350–600 W en air (L40S 350 W, RTX PRO 6000 600 W) ;
  au-delà, le format ne suit plus thermiquement en 19".
- **Avantages** : interopérabilité, achat à l'unité, serveurs généralistes, revente.
- **Limites** : PCIe 5.0 x16 = 128 Go/s théoriques — **64× moins** que la HBM d'un B200.
  Tout échange GPU↔GPU ou GPU↔CPU intensif passe par ce goulot.

### 17.3. OAM / UBB (AMD, standard OCP)

- **OAM** (OCP Accelerator Module) : l'équivalent AMD du SXM, standard ouvert OCP.
  MI300X/MI325X/MI350X/MI355X sont des modules OAM 2.0.
- **UBB** (Universal Baseboard) : la baseboard 8× OAM — l'équivalent du HGX NVIDIA,
  mais **standard ouvert** : un UBB peut théoriquement accueillir des modules de
  fournisseurs différents (en pratique : un seul fournisseur par board).
- **Avantage stratégique** : standard ouvert = concurrence entre OEM (Dell, HPE,
  Supermicro, Gigabyte) sur le même design de référence AMD.
- **MI350P** : exception — carte **PCIe** 144 Go HBM3e, 350 W, pour les datacenters
  classiques sans UBB.

### 17.4. Tableau récapitulatif des formats

| Format | GPUs concernés | TDP max typique | Inter-GPU | Achat |
|---|---|---|---|---|
| SXM5 | H100, H200 | 700 W | NVLink 4 + NVSwitch | Serveur 8× uniquement |
| SXM Blackwell | B200 | 1 000 W | NVLink 5 + NVSwitch | Serveur 8× / rack NVL72 |
| PCIe FHFL | L40S, RTX PRO 6000, 4090/5090, H100/H200 PCIe, MI350P | 350–600 W | Bridge ou rien | À l'unité |
| OAM 2.0 | MI300X, MI325X, MI350X, MI355X | 750–1 400 W | Infinity Fabric + switch | Plateforme 8× |
| NVL (double) | H100 NVL, H200 NVL | 2× 700 W | NVLink direct carte-à-carte | Serveur 4× NVL |

## 18. NVLink et NVSwitch

### 18.1. Générations NVLink (vérifié le 27/09/2026)

