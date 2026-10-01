---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-22
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "Cerebras", "Falcon", "Google", "Groq", "Huawei", "Hugging Face", "Intel", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "agent", "agents", "amd", "ascend", "awq", "aws", "blackwell", "compute", "crescent island", "distillation"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [3271, 3459]
sha256: 328662f918c66f0313ba77b4a2ea9bd2b51511d0917c57fac6532ab8483f46f0
---

# Les GPU datacenter / IA — présent et futur vérifié

ZeRO-3 fait tenir un 70B full-FT sur 8× L40S (48 Go) — au prix d'une
communication ×1,5. **Le choix standard en 2026 : FSDP (PyTorch natif).**

## 180. RLHF / DPO : le surcoût

- **DPO** (Direct Preference Optimization) : ~1,3× la VRAM du SFT — accessible.
- **PPO** (RLHF classique) : 4 modèles en mémoire (policy, ref, reward, value) =
  **~4× la VRAM du SFT**. Un 70B en PPO = ~300 Go → 4× H100 minimum.
- **Tendance 2026** : DPO/ORPO remplace PPO dans 80 % des cas (moins cher,
  résultats équivalents sur l'alignement simple).

## 181. Distillation : le raccourci malin

- **Principe** : un gros modèle (teacher) génère des données pour entraîner un
  petit (student). Le student 8B distillé d'un 70B/405B bat souvent le 8B de base
  de 5–15 % sur la tâche cible.
- **Coût** : génération teacher (1× B200 quelques jours) + training student
  (8× H100 quelques jours) — **10× moins cher** qu'un pre-training 70B.
- **Limite** : le student ne dépasse jamais vraiment le teacher sur le général.

## 182. Moteurs d'inférence : comparatif chiffré (70B FP8, 1× H200)

| Moteur | tok/s (batch 1) | tok/s (batch 64) | Latence p99 | Note |
|---|---|---|---|---|
| vLLM | ~60 | ~2 200 | ~40 ms | Le standard, PagedAttention |
| SGLang | ~62 | ~2 600 | ~35 ms | RadixAttention, top agents |
| TensorRT-LLM | ~70 | ~3 000 | ~30 ms | Le plus rapide, build complexe |
| TGI | ~55 | ~1 800 | ~45 ms | Simple, Hugging Face |

**Choix** : vLLM par défaut (écosystème), SGLang si agents/prefix caching massif,
TensorRT-LLM si chaque tok/s compte et l'équipe est senior.

## 183. Quantifications : comparatif qualité/vitesse (70B)

| Méthode | VRAM | tok/s relatif | Perte qualité | Outil |
|---|---|---|---|---|
| BF16 | 161 Go | 1,0× | Référence | — |
| FP8 (E4M3) | 80 Go | 1,6× | < 1 % | vLLM, TRT-LLM |
| AWQ (INT4) | 40 Go | 1,3× | 1–2 % | AutoAWQ |
| GPTQ (INT4) | 40 Go | 1,3× | 1–3 % | AutoGPTQ |
| GGUF Q4_K_M | 40 Go | 0,9× (CPU+GPU) | 2–4 % | llama.cpp |
| FP4 (NVFP4) | 20 Go | 2,2× | 2–5 % | TRT-LLM (B200+) |

## 184. Batching : le multiplicateur oublié

| Batch | tok/s agrégés (70B FP8, H200) | Latence/token |
|---|---|---|
| 1 | ~60 | ~17 ms |
| 8 | ~450 | ~18 ms |
| 32 | ~1 500 | ~21 ms |
| 64 | ~2 200 | ~29 ms |
| 128 | ~2 600 | ~49 ms |

**Leçon** : à batch 64, un H200 produit **36×** plus de tokens qu'à batch 1.
Le batching continu (vLLM) est le levier n°1 du $/token — bien avant le choix du GPU.

## 185. Latence : TTFT vs TPOT

- **TTFT** (Time To First Token) : dépend du **compute** (prefill) — gros GPU,
  batch petit, contexte court.
- **TPOT** (Time Per Output Token) : dépend de la **bande passante** — HBM rapide.
- **SLA type chatbot** : TTFT < 500 ms, TPOT < 50 ms (20 tok/s).
- **SLA agent** : TTFT < 1 s accepté (les outils prennent de toute façon des secondes).

## 186. Concurrence : combien d'utilisateurs par GPU

| GPU | 70B FP8, 20 tok/s/user | Users simultanés |
|---|---|---|
| RTX PRO 6000 | ~1 100 tok/s (batch 64) | ~50 |
| H200 | ~2 200 tok/s | ~100 |
| B200 | ~3 500 tok/s (FP4) | ~175 |
| 8× B200 | ~28 000 tok/s | ~1 400 |

Hypothèse : 50 % d'utilisation moyenne (pics). Diviser par 2 pour la HA.

## 187. Modèles de serving : 3 architectures

```
A. Monolithe : 1 GPU = 1 modèle (simple, gaspillage si petit modèle)
B. Multiplexé : vLLM multi-LoRA (1 base 70B + 50 adapteurs = 50 clients)
C. Routé : petit modèle (8B) d'abord, escalade vers 70B si besoin (cascade)
```

L'option C divise le coût par 3–5 sur les workloads mixtes (80 % des requêtes
sont simples) — c'est l'architecture des assistants de production en 2026.
## 188. Conversions d'unités (pense-bête)

| Conversion | Formule |
|---|---|
| W → kWh/an (24/7) | W × 8,76 |
| kW → A (tri 400 V, cos φ 0,9) | kW × 1,6 |
| kW → A (mono 230 V, cos φ 0,9) | kW × 4,8 |
| Go/s → To/s | ÷ 1 000 (marketing) ou ÷ 1 024 (réel) |
| TFLOPS → tok/s (théorique) | TFLOPS × 10^12 / (2 × params × tokens) — borne haute |
| $/h → $/an (24/7) | $/h × 8 760 |
| PUE → surcoût élec | (PUE − 1) × P_IT × 8 760 × prix_kWh |

## 189. Ordres de grandeur à connaître par cœur

- PCIe 5.0 x16 : **128 Go/s** | NVLink 4 : **900 Go/s** | NVLink 5 : **1,8 To/s**
- HBM H100 : **3,35 To/s** | HBM B200 : **8 To/s** | HBM4 MI455X : **19,6 To/s**
- 70B BF16 : **140 Go** | 70B FP8 : **70 Go** | 70B Q4 : **35 Go**
- KV cache 70B @ 8K : **~10 Go** | @ 128K : **~160 Go**
- Nœud 8× H100 : **8 kW** | 8× B200 : **10,6 kW** | Rack NVL72 : **~120 kW**
- PUE air : **1,4–1,6** | PUE DLC : **~1,1**
- Break-even achat/location : **~15 mois** à 24/7, **60 %** d'utilisation seuil

## 190. Acronymes du guide (index)

AITER, AWQ, BF16, BMC, BOM, CDU, DCGM, DLC, DPO, ECC, FNUZ, FP4/8/16, FSDP,
GDDR, GQA, GPTQ, HBM, HGX, IB, KV cache, MIG, MLA, MLPerf, MoE, NCCL, NVL,
NVLink, OAM, PPO, PUE, QLoRA, RCCL, RDMA, RLHF, SFT, SXM, TDP, TPOT, TTFT,
UBB, UEC, vLLM, Xid, ZeRO.

## 191. Ce que ce guide ne couvre pas (pistes suivantes)

- **TPU Google / Trainium AWS** : les alternatives cloud-only (pas d'achat).
- **ASICs chinois** (Huawei Ascend, Biren) : marché restreint, à vérifier.
- **GPU Intel datacenter** (Gaudi 3, Falcon Shores) : Crescent Island en
  sampling H2 2026 — à réévaluer en 2027.
- **Cerebras / Groq LPUs** : l'inférence ultra-low-latency (SRAM géante) —
  un monde à part, pertinent pour les agents temps réel.
- **Quantique** : hors sujet en 2026 pour l'IA appliquée.

## 192. Historique des révisions

| Date | Version | Changements |
|---|---|---|
| 27/09/2026 | 1.0 | Création — 176 sections, données vérifiées le jour même |

*Fin du guide. 192 sections. Données vérifiées le 27/09/2026.*
## 193. Fiches terrain par GPU (synthèse décideur)

### 193.1. RTX 4090 — le labo

Forces : prix, dispo immédiate, 24 Go suffisants pour 8B–13B et QLoRA 70B.
Faiblesses : pas d'ECC, garantie gaming, fin de vie (prix en hausse).
Verdict : acheter d'occasion à < 2 000 $ pour un lab, jamais pour la prod.

### 193.2. RTX 5090 — le labo musclé

Forces : 32 Go GDDR7, 1,79 To/s, FP4 — le meilleur $/TOPS du marché.
Faiblesses : 575 W, prix rue +30 à +145 % vs MSRP, pas d'ECC.
Verdict : le choix lab 2026 si on la trouve à < 2 800 $.

### 193.3. L40S — le cheval de labour de l'inférence

Forces : 48 Go ECC, 350 W, MIG, prix contenu, dispo.
Faiblesses : pas de NVLink, FP8 correct sans plus.
Verdict : le meilleur GPU « pro raisonnable » pour RAG/embeddings.

### 193.4. RTX PRO 6000 Blackwell — le 70B à 12 k$

Forces : 96 Go ECC, 1,79 To/s, MIG 4×, achat à l'unité.
Faiblesses : pas de NVLink, GDDR7 chère en 2026 (+55 %).
Verdict : le sweet spot inférence 70B sur site en 2026.

### 193.5. H100 — le standard en fin de cycle

Forces : écosystème, MLPerf partout, prix en baisse.
Faiblesses : 80 Go justes pour 70B, fin de cycle (Rubin en prod).
Verdict : négocier -20 % et l'utiliser 3 ans, ou passer au B200/H200.

### 193.6. H200 — le 70B sans sharding

Forces : 141 Go, même tooling que H100, dispo.
Faiblesses : compute identique au H100 (pas de FP4), cher au Go.
Verdict : pour l'inférence 70B–123B, pas pour le training (prendre B200).

### 193.7. B200 — le champion actuel

Forces : 192 Go, 8 To/s, FP4, 2,2× le H100 en training.
Faiblesses : 1 000 W, DLC quasi-obligatoire, prix, délais.
Verdict : le choix par défaut pour tout projet neuf ambitieux fin 2026.

### 193.8. MI300X — l'alternative crédible

Forces : 192 Go, 5,3 To/s, prix agressifs (1,99 $/h cloud).
Faiblesses : ROCm (POC obligatoire), pas de FP4.
Verdict : viable en inférence si l'équipe est senior Linux.

### 193.9. MI325X — le max VRAM Hopper-like

Forces : 256 Go en 1 GPU — le record 2024-2025.
Faiblesses : 1 000 W, éclipsé par CDNA 4.
Verdict : niche (contexte très long sur 1 GPU), sinon prendre MI355X.

### 193.10. MI350X / MI355X — le B200 d'AMD

