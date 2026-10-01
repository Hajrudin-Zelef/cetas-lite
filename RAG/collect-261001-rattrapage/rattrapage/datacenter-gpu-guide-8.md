---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-8
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Anthropic", "Broadcom", "CoreWeave", "Google", "Groq", "Hugging Face", "Intel", "Meta", "Nvidia", "OpenAI", "Oracle", "SGLang", "TSMC", "TensorRT-LLM", "vLLM"]
dates: ["2026-09-27"]
keywords: ["gpu", "agents", "amd", "blackwell", "compute", "crescent island", "ethernet", "fp4", "fp8", "hbm4", "helios", "inference"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1006, 1144]
sha256: 1f20505402e124eb095e64057adc8462c6c689bca577f01e193adf2944d3e3c8
---

# Les GPU datacenter / IA — présent et futur vérifié

Ordre de grandeur attendu plateforme complète : **300 000–450 000 $** (estimation,
à vérifier — AMD ne publie pas de prix, les OEM chiffrent sur devis).

## 37. Pile logicielle : CUDA, ROCm, frameworks

### 37.1. NVIDIA : CUDA

| Composant | Rôle | Version min. (H100+) |
|---|---|---|
| Driver NVIDIA | Pilote kernel | 550.x+ |
| CUDA Toolkit | Compilateur/runtime | 12.4+ (FP8 : 12+) |
| cuDNN / cuBLAS | Libs deep learning | 9.x |
| NCCL | Collectives multi-GPU | 2.20+ |
| TensorRT / TRT-LLM | Inférence optimisée | 10.x |
| Triton Inference Server | Serving | 24.x+ |
| Images NGC | Conteneurs prêts | `nvcr.io/nvidia/pytorch` |

### 37.2. AMD : ROCm

| Composant | Rôle | Note 2026 |
|---|---|---|
| ROCm | Stack complète (équiv. CUDA) | 6.x–7.x, open-source |
| RCCL | Collectives (équiv. NCCL) | Bon en 8× homogène |
| AITER / ATOM | Kernels optimisés vLLM/SGLang | Le « secret » des perfs MI300X |
| MIOpen / hipBLAS | Libs maths | Équiv. cuDNN/cuBLAS |
| Images Docker | `rocm/pytorch`, `rocm/vllm` | Vérifier le tag gfx942/gfx950 |

### 37.3. Frameworks d'inférence (les deux mondes)

- **vLLM** : PagedAttention, le standard de l'inférence LLM. Supporte CUDA et ROCm.
- **SGLang** : RadixAttention (prefix caching agressif), très fort en agents.
- **TensorRT-LLM** : le plus rapide sur NVIDIA, mais verrouillé à l'écosystème NVIDIA.
- **TGI** (Hugging Face) : simple, bon pour démarrer.
- **Ollama / LM Studio** : poste local, RTX 4090/5090 — pas pour la prod.

### 37.4. Règle de décision logicielle

```
Si équipe déjà CUDA + besoin immédiat → NVIDIA (risque zéro).
Si inférence vLLM standard + budget serré → AMD MI300X/MI355X (tester d'abord !).
Si training custom kernels → NVIDIA (l'écosystème Triton/CUDA est 3 ans devant).
```

## 38. Rent vs buy : seuil de rentabilité

| GPU | Achat (carte/serveur) | Location on-demand | Seuil (h/mois) |
|---|---|---|---|
| RTX 4090 | 2 400 $ | 0,34 $/h | ~590 h/mois (rentable d'acheter si > 80 % d'utilisation) |
| L40S | 10 000 $ | 0,96 $/h | ~870 h/mois |
| H100 (serveur/8) | 36 000 $ | 3,33 $/h | ~900 h/mois |
| H200 (serveur/8) | 46 000 $ | 3,70 $/h | ~1 035 h/mois — jamais atteint → **louer** sauf usage > 85 % |
| B200 (serveur/8) | 56 000 $ | 5,50 $/h | ~850 h/mois |

Formule : `seuil = prix_achat / (prix_location × 0,7)` (0,7 = décote revente/obsolescence).
**Règle** : acheter si utilisation > 60–70 % sur 24 mois ; louer sinon. Le B200 acheté
à 40 % d'utilisation coûte **2× plus cher** que la location.
## 39. À venir — roadmap NVIDIA vérifiée au 27/09/2026

### 39.1. Vera Rubin (en production H2 2026)

- **Statut** : plateforme **en pleine production** (annonce GTC 2026, confirmée mai 2026
  à GTC Taipei ; livraisons partenaires H2 2026, ramp-up Q4 2026 → début 2027).
- **Composition** : 7 puces — CPU **Vera** (88 cœurs Olympus custom), GPU **Rubin**
  (50 PFLOPS NVFP4), switch **NVLink 6** (3,6 To/s par GPU), **ConnectX-9** SuperNIC,
  **BlueField-4** DPU, switch **Spectrum-6** Ethernet, **Groq 3 LPU** intégré.
- **NVL72** : 72 GPU Rubin + 36 CPU Vera par rack, 20,7 To de HBM4, 260 To/s de
  scale-up, inlet eau 45 °C, **refroidissement liquide sec sans ventilateurs**,
  compatible sidecar 800VDC. Mise sous tension : 47 min après livraison (chiffre NVIDIA).
- **Claims NVIDIA** : 4× le training et 10× l'inférence par watt vs Blackwell ;
  coût par token ÷10 ; MoE entraîné avec 4× moins de GPU.
- **Clients** : OpenAI (déploiement Q3 2026, 10× tokens vs GB200 selon CoreWeave),
  CoreWeave, Google Cloud, Azure, Meta, Dell, Supermicro (DLC-2).
- **Racks** : 5 types (NVL72 compute, rack 256 CPU Vera, rack Groq 3 LPX 256 LPU pour
  latence ultra-faible, rack stockage BlueField-4 STX, rack réseau Spectrum-6 SPX).

### 39.2. B300 / GB300 (Blackwell Ultra — en production)

- **B300** : 288 Go HBM3e, 8 To/s, FP4 dense 15 PFLOPS, TDP **1 400 W**, NVLink 5.
- **DGX B300** : 8 GPU, ~2,3 To, 11,2 kW, **liquide fortement recommandé**.
- **GB300 NVL72** : rack 72 GPU, production de masse mi-2026 (~89 Md$ de revenus
  attendus selon la presse).
- **Prix location** : ~9,08 $/h on-demand (juillet 2026, Spheron).

### 39.3. Après Rubin : Feynman (officialisé par NVIDIA)

- NVIDIA a officialisé le nom de la génération post-Rubin : **Feynman**
  (annoncé dans la roadmap GTC 2024-2025 : Blackwell → Rubin → Feynman).
- **Aucune spec officialisée au 27/09/2026** : pas de chiffres de mémoire, TDP ou
  calendrier précis au-delà de « après Rubin ».

## 40. À venir — roadmap AMD vérifiée au 27/09/2026

### 40.1. MI450 / MI455X « Altair » + rack Helios (H2 2026)

- **Statut** : **lancé à AMD Advancing AI 2026** (août 2026), en pleine production,
  livraisons fin Q3 2026 → ramp-up 2027. Demande « au-dessus des attentes » (Lisa Su).
- **Specs** : CDNA 5, TSMC 2 nm, 320 milliards de transistors, **432 Go HBM4**,
  **19,6 To/s** par GPU, 40 PFLOPS FP4 dense / 20 PFLOPS FP8.
- **Déclinaisons** : **MI455X** (flagship), **MI450** (volume hyperscale),
  **MI430X** (HPC, FP64 optimisé).
- **Helios** : rack 72× MI455X + EPYC Venice + NIC Pensando Vulcano, **31 To HBM4**,
  1,4 Po/s agrégés, 2,9 exaflops FP4, UALink 260 To/s scale-up, Ethernet UEC 43 To/s
  scale-out. Design Schneider Electric : **246 kW par rack**, 100 % liquide.
- **Clients** : Anthropic (jusqu'à 2 GW + 5 Md$ d'investissement), OpenAI (6 GW,
  Helios en prod depuis 3 mois pour des workloads GPT), Meta, Oracle.
- **Claim AMD** : jusqu'à **30 % de tokens/$ de plus que Rubin NVL72** (chiffre AMD).

### 40.2. EPYC Venice (6e gén, H2 2026)

- CPU serveur Zen 6, optimisé pour l'IA (inférence co-processing), intégré à Helios.
- Morgan Stanley : 6,75 M d'unités prévues en 2027.

### 40.3. MI500 (2027)

- Annoncé à l'Analyst Day AMD : génération post-CDNA 5, **2027**.
- Aucune spec officialisée au 27/09/2026.

## 41. Rumeurs — explicitement marquées RUMEUR

> Tout ce qui suit n'est **pas** officialisé. Ne pas dimensionner un achat dessus.

- **RUMEUR** : NVIDIA « Blackwell-Next » repéré dans un patch du kernel Linux
  (juin 2026) — nom de code inconnu, rien d'officialisé.
- **RUMEUR** : déclinaison **B200A / B300A** allégée pour la Chine (dans la lignée
  des H20/H800) — aucune annonce NVIDIA au 27/09/2026.
- **RUMEUR** : prix et TDP du **Rubin GPU seul** (hors rack) — non publiés par NVIDIA.
- **RUMEUR** : Intel **Crescent Island** (Xe3P, 160 Go LPDDR5X, 350 W, sampling H2 2026)
  — specs publiées par Intel mais produit non disponible ; les chiffres de perf sont
  des estimations tierces (chipsandcheese), pas des mesures.
- **RUMEUR** : Google **TPU v9 « Triggerfish »** avec MediaTek — développement
  rapporté par la presse, pas de fiche produit.
- **RUMEUR** : OpenAI + Broadcom « Jalapeño » (puce interne, déploiement fin 2026) —
  annonce presse, pas de specs publiques.
- **Rien d'officialisé** au 27/09/2026 concernant : RTX 6090, L50S/L60S, H300,
  MI400 (le nom « MI400 » circule dans la presse mais AMD communique sur « MI450 »),
  B400.

## 42. Glossaire (35 termes)

