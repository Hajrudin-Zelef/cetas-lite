---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-51
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Baidu", "Google", "Hugging Face", "Meta", "Microsoft", "Nvidia", "OpenAI", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2026-09-22", "2026-09-27"]
keywords: ["agents", "awq", "benchmarks", "claude", "dpo", "fable 5", "fine-tuning", "gguf", "gptq", "gpu", "leaderboard", "llama"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3957, 4063]
sha256: fed8f2f237cc7744c7553f87ce9c0360e828e0d593f7d3c20a8851d967a76bdd
---

# IA — Le grand dossier

### 3.4. La gamme de modèles (synthèse — détails dans les volumes IA)

> Les volumes 1-2 du dossier (« ia_modeles_*.md ») documentent modèle par modèle. Ici, la carte
> d'ensemble au 27/09/2026.

| Génération | Modèles | Positionnement | Ordre de prix (entrée/sortie, $/M tokens — « à vérifier » sur la grille du jour) |
|---|---|---|---|
| Claude 4.5 (janv. 2026) | Opus 4.5, Sonnet 4.5, Haiku 4.5 | Génération « agents + code + computer use » ; Sonnet 4.5 décrit comme le plus aligné | Opus ~5/25 ; Sonnet ~3/15 ; Haiku ~1/5 |
| Claude 4.6 (fév. 2026) | Opus 4.6, Sonnet 4.6 | Itération capacité ; contexte 1M | Opus 5/25 ; Sonnet 3/15 |
| Claude 5 (2026) | Opus 5, Sonnet 5, **Fable 5 / 5.1** | Fable = tier « max capability » (~2× le prix d'Opus : 10/50) ; Sonnet 5 passé en tarif standard 2/10 (hausse annulée sept. 2026) | Fable 10/50 ; Opus 5/25 ; Sonnet 2/10 |
| Claude Opus 5.5 (22/09/2026) | Opus 5.5 | Niveau Fable 5.1 à −40 % | « à vérifier » |
| Mythos (avr. 2026) | — | Non public (~40 orgs) | — |

**Règle d'usage** (valable pour tous les labs) : dans un tier donné, une montée de version est un
changement de capacité **à coût unitaire constant** — aucune raison de rester sur l'ancien modèle
du même tier pour le prix. Le routage par complexité (Haiku → Sonnet → Opus/Fable) est le levier
FinOps n°1 : Sonnet 5 à 0,4× le prix d'Opus change l'équation du « tout-Opus par défaut ».

---

## 4. Les stacks techniques (« starck » = stacks)

> Pour un sysadmin, c'est la section à plus forte valeur : comprendre *ce qui tourne où* dans
> une chaîne IA, c'est pouvoir l'exploiter, la superviser et la dépanner. Chaque stack est
> présentée en 3 temps : schéma textuel, composants réels (noms, versions indicatives 2026),
> exemple de code ou de config.

### 4.1. Stack entraînement (training)

**À quoi ça sert** : fabriquer un modèle (pre-training) ou l'adapter (fine-tuning). Tu n'entraîneras
probablement jamais un LLM from scratch (budget : dizaines de millions de $ en calcul), mais le
fine-tuning léger (LoRA) et l'évaluation sont à ta portée — et comprendre la stack aide à
dimensionner l'inférence.

```
┌─────────────┐   ┌──────────────┐   ┌───────────────┐   ┌──────────────┐
│  Données    │──▶│  Tokenizer   │──▶│  Framework    │──▶│  Cluster GPU │
│ (corpus,    │   │  (BPE,       │   │  PyTorch +    │   │  NCCL,       │
│  qualité!)  │   │  tiktoken)   │   │  DeepSpeed /  │   │  InfiniBand, │
└─────────────┘   └──────────────┘   │  FSDP /       │   │  checkpoint  │
                                     │  Megatron-LM  │   │  réguliers   │
                                     └───────────────┘   └──────────────┘
        ┌──────────────┐   ┌──────────────┐   ┌──────────────┐
        │  Éval        │──▶│  Alignement  │──▶│  Quantif.   │
        │  (benchmarks)│   │  RLHF/DPO/   │   │  (GGUF,      │
        │              │   │  RLAIF       │   │  AWQ, GPTQ)  │
        └──────────────┘   └──────────────┘   └──────────────┘
```

| Couche | Outils réels (2026) | Notes sysadmin |
|---|---|---|
| Framework | **PyTorch** (dominant recherche/industrie), JAX (Google/DeepMind), PaddlePaddle (Baidu) | PyTorch 2.x : `torch.compile`, distributed natif |
| Parallélisme | **DeepSpeed** (Microsoft), **FSDP** (natif PyTorch), Megatron-LM (NVIDIA) | ZeRO : partitionne états d'optimiseur/gradients/params sur N GPU |
| Communication | **NCCL** (NVIDIA), InfiniBand / RoCE, NVLink/NVSwitch intra-nœud | Le réseau *est* le goulot : 3,2 Tb/s par GPU sur les dernières générations (« à vérifier ») |
| Orchestration | Slurm (HPC classique), Kubernetes + KubeFlow / Volcano | Checkpointing régulier obligatoire : un run de plusieurs semaines *va* perdre des nœuds |
| Données | Datasets (Hugging Face), traitement : deduplication (MinHash), filtrage qualité, décontamination des benchmarks | « Garbage in, garbage out » — 80 % du travail, 20 % du glamour |
| Fine-tuning léger | **LoRA / QLoRA** (adapteurs de rang faible), **Unsloth**, **Axolotl**, **LLaMA-Factory** | QLoRA : fine-tune un 70B sur *un seul* GPU 48 Go — c'est ça qui rend le fine-tuning accessible |
| Éval | **lm-evaluation-harness** (EleutherAI), Open LLM Leaderboard, benchmarks métier maison | Toujours évaluer *avant/après* sur tes propres données, pas seulement sur les benchmarks publics |

**Exemple réel — fine-tuning QLoRA minimal (Hugging Face `trl` + `peft`) :**

```python
from datasets import load_dataset
from trl import SFTTrainer, SFTConfig
from peft import LoraConfig
from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

model_id = "Qwen/Qwen3-8B"  # exemple : remplace par ton modèle
tok = AutoTokenizer.from_pretrained(model_id)
model = AutoModelForCausalLM.from_pretrained(
    model_id, torch_dtype=torch.bfloat16,
    load_in_4bit=True, device_map="auto",  # QLoRA : 4-bit NF4
)
peft_cfg = LoraConfig(r=16, lora_alpha=32, lora_dropout=0.05,
                      target_modules=["q_proj","k_proj","v_proj","o_proj"],
                      task_type="CAUSAL_LM")
args = SFTConfig(output_dir="qwen3-8b-lora",
                 per_device_train_batch_size=2,
                 gradient_accumulation_steps=8,
                 num_train_epochs=3, learning_rate=2e-4,
                 max_seq_length=2048, packing=True)
trainer = SFTTrainer(model=model, args=args, train_dataset=load_dataset("json", data_files="train.jsonl")["train"],
                     peft_config=peft_cfg, processing_class=tok)
trainer.train()
trainer.model.save_pretrained("qwen3-8b-lora-adapter")  # seul l'adaptateur (~100 Mo) est sauvé
```

**Ordres de grandeur à retenir** : entraîner un 8B « from scratch » ≈ milliers de GPU-heures ;
fine-tuner en LoRA ≈ heures sur 1 GPU ; inférer un 8B quantifié ≈ CPU correct ou petit GPU.
D'où la règle : **on n'entraîne presque jamais, on adapte souvent, on sert toujours.**

### 4.2. Stack inférence (servir un modèle)

**À quoi ça sert** : transformer des poids en API rapide, concurrente, économe. C'est *ta* stack
si tu auto-héberges (Ollama sur un serveur, vLLM en prod).

```
 Requête ──▶ ┌──────────────┐ ┌──────────────┐ ┌──────────────┐ ┌──────────────┐
             │  API / LB    │ │  Moteur      │ │  Optimisations│ │  GPU / CPU   │
             │  (OpenAI-    │▶│  d'inférence │▶│  KV-cache,   │▶│  (CUDA,      │
             │  compatible) │ │  vLLM/TGI/   │ │  continuous  │ │  ROCm,       │
             │              │ │  llama.cpp   │ │  batching,   │ │  Metal…)     │
             └──────────────┘ │  TensorRT-LLM│ │  quantif.    │ └──────────────┘
                             └──────────────┘ └──────────────┘
```

