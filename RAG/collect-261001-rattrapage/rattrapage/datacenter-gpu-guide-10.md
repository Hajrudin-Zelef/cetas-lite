---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-10
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Nvidia", "vLLM"]
dates: []
keywords: ["gpu", "agents", "amd", "awq", "benchmarks", "compute", "deepseek", "embedding", "embeddings", "fine-tuning", "fp4", "fp8"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1277, 1437]
sha256: 5e6d0f26b0d4bd9591b8b02a9f08a4d8f38c66d225998b81c0dee605550a0317
---

# Les GPU datacenter / IA — présent et futur vérifié

## 47. Retour terrain : 8× L40S en inférence RAG

- **Config** : serveur 4U 8× L40S, 2× EPYC, 1 To RAM.
- **Workload** : embeddings (BGE/E5) + reranking + LLM 8B/32B en support.
- **Débit** : ~2 000–4 000 requêtes embedding/s (batch 512), ~800 tok/s sur 8B.
- **Coût** : ~110 k$ serveur, ~6 300 €/an d'élec, PUE 1,4.
- **Pourquoi ça marche** : l'embedding est compute-léger et batchable — le L40S
  à 350 W est imbattable en perf/watt sur ce créneau.

## 48. Retour terrain : 8× H100 en training 70B

- **Run type** : fine-tuning complet Llama 70B, 100 Md tokens, FP8.
- **Durée** : ~10–14 jours sur 8× H100 (MFU ~40 %).
- **Coût compute** : 14 j × 24 h × 8 × 3,33 $/h ≈ **8 900 $** en location.
- **Checkpoints** : 140 Go × 2 (poids + optimiseur partiel) toutes les 2 h →
  ~3,4 To de checkpoints par run. Prévoir le stockage NVMe en conséquence.
- **Leçon** : 90 % des échecs viennent du **stockage** (checkpoint corrompu) ou du
  **réseau** (NCCL timeout), pas du GPU.

## 49. Retour terrain : MI300X en inférence (le pari AMD)

- **Constat 2026** : vLLM + ROCm 6.x + AITER sur MI300X = 90–95 % des perfs H100
  par dollar sur l'inférence 70B, avec 2,4× la mémoire.
- **Prérequis** : image Docker `rocm/vllm` à jour, gfx942, kernels AITER activés
  (`VLLM_ROCM_USE_AITER=1`).
- **Point dur restant** : les modèles « exotiques » (nouveaux MoE, MLA custom)
  mettent 2–8 semaines à être optimisés sur ROCm après leur sortie.
- **Verdict** : viable en prod si l'équipe maîtrise Linux et accepte un léger
  décalage de support des nouveaux modèles.

## 50. Benchmarks MLPerf : ce qu'ils valent vraiment

- **MLPerf Training v3.0** : H100 recordman 2023 (0,82 min sur 3D U-Net à 432 GPU).
- **MLPerf Inference v3.1** : H100 jusqu'à 4,5× l'A100.
- **MLPerf Training (B200)** : jusqu'à **2,2× le H100** (soumissions 2025-2026).
- **MLPerf Inference v6.0** (avril 2026) : MI355X à 92–104 % du B300 sur Llama 2 70B.
- **Lecture critique** : MLPerf mesure le **pic optimisé par le vendeur** sur un
  workload figé. En prod, diviser par 1,5 à 2 (framework générique, batch variable,
  overhead réseau). Utile pour comparer des générations, pas pour dimensionner.

## 51. FP8 vs FP4 : quand la basse précision casse (ou pas)

| Précision | Perte qualité typique (LLM) | Cas OK | Cas à éviter |
|---|---|---|---|
| BF16 | Référence | Training, tout | — |
| FP8 (E4M3) | < 1 % perplexité | Inférence 7B–405B, training | Modèles < 3B sensibles |
| FP8 (E5M2) | Légèrement plus | Gradients (training) | Inférence pure |
| INT8 (dynamique) | 1–2 % | Inférence, CPU-friendly | MoE avec experts rares |
| FP4 (NVFP4/MXFP4) | 2–5 % | Chatbots, RAG, agents | Code, maths, raisonnement long |
| INT4 (GPTQ/AWQ) | 2–4 % | 70B sur petit GPU | < 13B (dégradation visible) |

**Règle** : valider **chaque** modèle en FP4/FP8 sur **vos** benchmarks métier
(MMLU ne suffit pas — tester le français, le code, vos prompts réels).
Le FP4 est un multiplicateur de débit, pas un lunch gratuit.

## 52. MoE : dimensionner autrement

- **Params totaux vs actifs** : DeepSeek R1 = 671B totaux / 37B actifs.
  La VRAM suit le total, le compute suit l'actif.
- **Conséquence** : un MoE 671B Q4 (335 Go) tourne sur 4× H200 là où un dense
  671B exigerait 8× B200.
- **Piège** : l'**expert parallelism** exige un inter-GPU rapide (NVLink) —
  les experts sont répartis entre GPU et chaque token route vers 2–8 experts.
  En PCIe sans NVLink, la latence explose.
- **Batch** : les MoE adorent les gros batchs (amortissent le routage).

## 53. Parallélismes : TP, PP, DP, EP — qui fait quoi

| Parallélisme | Découpe | Communication | Quand |
|---|---|---|---|
| Tensor (TP) | Chaque couche sur N GPU | All-reduce à chaque couche | Modèle > 1 GPU, inférence |
| Pipeline (PP) | Couches réparties en étages | P2P entre étages | Gros modèles, training |
| Data (DP) | Même modèle, données différentes | All-reduce gradients | Training multi-GPU/nœuds |
| Expert (EP) | Experts MoE répartis | All-to-all | MoE uniquement |
| Sequence (SP) | Contexte découpé | All-gather | Contexte > 128K |

**Règle d'or** : TP intra-nœud (NVLink), DP inter-nœuds (InfiniBand), PP quand le
modèle dépasse le nœud, EP pour les MoE. Un mauvais placement = -50 % de MFU.

## 54. MFU : le vrai rendement du training

- **MFU** (Model FLOPs Utilization) : fraction du pic théorique réellement atteinte.
- **Repères** : 30–45 % = bon (training LLM standard), 50–60 % = excellent
  (kernels custom, parallélisme soigné), < 25 % = problème (I/O, réseau, config).
- **Exemple** : 8× H100 = 8 × 1 979 TFLOPS FP8 = 15,8 PFLOPS pic.
  À 40 % MFU : 6,3 PFLOPS utiles.
- **Leviers** : gros batchs, activation checkpointing équilibré, overlap
  communication/compute (NCCL async), FlashAttention-3, FP8 stable.

## 55. Temps de training : formule pratique

```
temps = (6 × N_param × N_tokens) / (N_GPU × FLOPS_pic × MFU)
```

Exemple : 70B params, 1 000 Md tokens, 64× H100, MFU 40 % :
6 × 70e9 × 1e12 / (64 × 1 979e12 × 0,40) ≈ 8,3 M secondes ≈ **96 jours**.
D'où les clusters de 1 000+ GPU pour les frontier models — et l'intérêt du B200
(2,2× → ~44 jours) ou de Rubin (4× → ~24 jours, claims NVIDIA).

## 56. Stockage pour l'IA : le parent pauvre des BOM

| Besoin | Débit requis | Solution type |
|---|---|---|
| Checkpoints training 8× H100 | 5–10 Go/s écriture | 4× NVMe U.2 RAID0 local + NAS |
| Dataset streaming (web-scale) | 2–5 Go/s lecture | NVMe local + cache, ou WEKA/VAST |
| Inférence (poids) | Lecture unique au chargement | NVMe simple 4 To |
| Logs/traces | Faible | SATA/NAS |

**Piège classique** : un NAS 1 GbE pour alimenter 8× H100 en training → les GPU
attendent les données (MFU < 15 %). Minimum : **100 GbE** vers le stockage ou
staging NVMe local avec pré-chargement.
## 57. Dimensionnement électrique détaillé : méthode

### 57.1. Étapes

1. **Inventaire** : lister chaque équipement avec sa puissance max (plaque ou TDP × 1,1).
2. **Foisonnement** : 1,0 pour les GPU (ils picent ensemble en all-reduce), 0,8 pour
   le réseau/stockage.
3. **Marge** : +20 % pour l'évolutivité et les appels de courant.
4. **Conversion** : P(kW) → I(A) = P / (√3 × 400 × 0,9) en triphasé.

### 57.2. Exemple : salle 4 nœuds 8× H200

| Poste | Calcul | Puissance |
|---|---|---|
| 4 nœuds × 8 kW AC | 4 × 8,0 | 32,0 kW |
| 2 switches IB 800G | 2 × 1,5 | 3,0 kW |
| Stockage NAS NVMe | 1 | 2,0 kW |
| Refroidissement (PUE 1,4 → 0,4 × IT) | 0,4 × 37 | 14,8 kW |
| Éclairage/services | — | 2,0 kW |
| **Total** | | **53,8 kW** |
| +20 % marge | × 1,2 | **64,6 kW** → arrivée **80 kVA** |

Soit ~117 A par phase en 400 V triphasé. Un TGBT 250 A s'impose.

## 58. Onduleurs pour salles GPU : dimensionnement

### 58.1. Règles

- **Puissance** : 1,25× la charge IT (jamais plus de 80 % de charge continue).
- **Autonomie** : 10 min à pleine charge (le temps du démarrage groupe).
- **Topologie** : double conversion (VFI) obligatoire — les PSU GPU sont sensibles
  aux micro-coupures.
- **Batteries** : VRLA (5 ans) ou Li-ion (10–15 ans, 2× le prix, 3× la durée de vie).

### 58.2. Tableau par config

| Config | Charge IT | Onduleur conseillé | Batteries (10 min) |
|---|---|---|---|
| 1 nœud 8× L40S (5 kW) | 5 kW | 10 kVA N+1 | ~15 kWh utiles |
| 1 nœud 8× H100 (8 kW) | 8 kW | 15 kVA N+1 | ~20 kWh utiles |
| 1 nœud 8× B200 (11 kW) | 11 kW | 20 kVA N+1 | ~28 kWh utiles |
| 4 nœuds 8× H200 (37 kW) | 37 kW | 60 kVA N+1 | ~95 kWh utiles |
| Rack GB200 NVL72 (~120 kW) | 120 kW | 200 kVA N+1 | ~300 kWh utiles |

**Note** : les batteries d'un 200 kVA / 10 min pèsent ~3–5 tonnes (VRLA) —
prévoir la portance du local. (Détail complet : voir le guide onduleurs de Zelef.)

## 59. Groupes électrogènes et démarrage

