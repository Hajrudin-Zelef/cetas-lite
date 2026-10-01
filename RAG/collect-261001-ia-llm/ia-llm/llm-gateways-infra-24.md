---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-24
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Moonshot", "Nvidia", "SGLang", "Unsloth", "vLLM"]
dates: []
keywords: ["gpu", "deepseek", "fine-tuning", "fp8", "gguf", "kimi", "lora", "moe", "nvidia", "open source", "qlora", "quantization"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2794, 2944]
sha256: b5b750e161c4ea4c6adf2deb5551ceaac45f0803f6683f90320d58b1d742cdb6
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

Ordres de grandeur (7B–8B, QLoRA r=16, séquences 2048) : ~8–14 Go de VRAM, quelques heures sur une 4090 pour 1–3 époques sur quelques milliers d'exemples. C'est du « fine-tuning du dimanche », pas du pré-entraînement.

### 162.4. Exemple minimal (SFT, correct et testable)

```python
# pip install unsloth  (Linux/WSL2 + GPU NVIDIA)
from unsloth import FastLanguageModel
from trl import SFTTrainer, SFTConfig
from datasets import Dataset

# 1. Jeu de données jouet au format instruction (remplace par tes vraies données)
data = Dataset.from_list([
    {"instruction": "Quelle est la tension de floating d'un bloc VRLA 12 V ?",
     "output": "13,5 à 13,6 V par bloc à 20 °C, avec compensation -3 mV/°C/cellule."},
    {"instruction": "Que signifie le code C6000 sur un copieur Kyocera ?",
     "output": "Rupture de chauffe (fuser) : vérifier thermistance, thermostat et lampe."},
])

def to_chatml(ex):
    return {"text":
        f"<|im_start|>user\n{ex['instruction']}<|im_end|>\n"
        f"<|im_start|>assistant\n{ex['output']}<|im_end|>"}
data = data.map(to_chatml)

# 2. Chargement du modèle en 4 bits
model, tokenizer = FastLanguageModel.from_pretrained(
    model_name="unsloth/Qwen3-8B",   # version pré-patchée par Unsloth ; adapte à ton modèle
    max_seq_length=2048,
    dtype=None,            # auto : float16 (T4/V100) ou bfloat16 (Ampere+)
    load_in_4bit=True,     # QLoRA : ~5 Go pour un 8B
)

# 3. Greffe des adaptateurs LoRA
model = FastLanguageModel.get_peft_model(
    model,
    r=16,                  # rank : 8/16/32 ; 16 = bon défaut
    lora_alpha=16,
    target_modules=["q_proj", "k_proj", "v_proj", "o_proj",
                    "gate_proj", "up_proj", "down_proj"],
    lora_dropout=0,        # 0 = optimisé dans Unsloth
    bias="none",
    use_gradient_checkpointing=True,  # "unsloth" : économise la VRAM au prix d'un peu de vitesse
    random_state=3407,
)

# 4. Entraînement
trainer = SFTTrainer(
    model=model,
    train_dataset=data,
    processing_class=tokenizer,
    args=SFTConfig(
        output_dir="./output",
        per_device_train_batch_size=2,
        gradient_accumulation_steps=4,
        num_train_epochs=3,
        learning_rate=2e-4,
        packing=True,      # 2 à 5x plus rapide sur séquences de longueurs mixtes
    ),
)
trainer.train()

# 5. Export GGUF pour Ollama / LM Studio
model.save_pretrained_gguf("mon-modele-metier", tokenizer, quantization_method="q4_k_m")
```

### 162.5. Limites honnêtes

- **Le fine-tuning ne remplace pas le RAG** : il apprend un style et un domaine, il n'apprend pas *tes* documents. Pour des faits qui changent (prix, versions firmware), le RAG reste supérieur — le modèle fine-tuné « sait » ce qu'il a vu à l'entraînement et hallucine le reste avec assurance.
- **Qualité des données > quantité** : 500 paires propres valent mieux que 10 000 paires bruitées. Prévois du temps de curation (c'est 80 % du travail).
- **Unsloth OSS = DDP uniquement** : une réplique complète par GPU, pas de FSDP/ZeRO — le modèle doit tenir sur **une seule GPU**. Pour du multi-GPU sérieux, c'est DeepSpeed/FSDP manuel ou un framework pro.
- **Évalue avant/après** : repasse tes 30 questions d'éval (section 127) sur le modèle de base puis fine-tuné. Un fine-tuning raté dégrade le modèle (catastrophic forgetting sur les capacités générales).

### 162.6. GRPO : le fine-tuning par récompense (exemple minimal)

Le SFT apprend des exemples parfaits. Le GRPO apprend d'une **fonction de récompense** — idéal quand tu sais dire « c'est bien / c'est mal » mais que tu n'as pas 2000 exemples parfaits. Exemple : récompenser les réponses qui citent une référence documentaire.

```python
from unsloth import FastLanguageModel
from trl import GRPOConfig, GRPOTrainer
import re

model, tokenizer = FastLanguageModel.from_pretrained(
    model_name="unsloth/Qwen3-8B",
    max_seq_length=2048,
    load_in_4bit=True,
)
model = FastLanguageModel.get_peft_model(model, r=16)

def reward_citation(completions, **kwargs):
    """+2 si la réponse cite une référence type [DOC-xxx] ou §x.y, 0 sinon."""
    scores = []
    for c in completions:
        text = c[0]["content"] if isinstance(c, list) else str(c)
        has_ref = bool(re.search(r"\[DOC-[\w-]+\]|§\s?\d+\.\d+|U\d{3,4}|C\d{4}", text))
        scores.append(2.0 if has_ref else 0.0)
    return scores

trainer = GRPOTrainer(
    model=model,
    processing_class=tokenizer,
    args=GRPOConfig(
        output_dir="./grpo-out",
        num_generations=4,     # 4 réponses candidates par prompt
        beta=0.04,             # garde-fou : ne pas trop s'éloigner du modèle de base
        learning_rate=5e-6,    # LR très faible : le RL est instable par nature
        max_prompt_length=1024,
        max_completion_length=512,
    ),
    train_dataset=dataset_prompts,   # juste des prompts, PAS de réponses attendues
    reward_funcs=[reward_citation],
)
trainer.train()
```

Notes : le GRPO consomme plus de VRAM que le SFT (plusieurs générations en parallèle) — prévois large ou réduis `num_generations` à 2 sur 24 Go. Et la fonction de récompense, c'est 80 % du résultat : une récompense mal conçue produit un modèle qui « game » la métrique (ex. : il cite des références inventées pour toucher le bonus — d'où le regex strict sur des formats vérifiables).

### 162.7. VRAM : combien pour quel modèle (QLoRA r=16, seq 2048 — ordres de grandeur)

| Modèle | Poids 4 bits | Entraînement (SFT) | Recommandation GPU |
|---|---|---|---|
| 1B–3B | ~2 Go | ~6–8 Go | RTX 3060 12 Go, T4 Colab |
| 7B–8B | ~5 Go | ~10–14 Go | RTX 3090/4090 24 Go |
| 13B–14B | ~8 Go | ~18–22 Go | 4090 24 Go (limite), A100 40 Go confort |
| 30B–32B | ~18 Go | ~40–48 Go | A100 80 Go, 2×4090 (non — DDP = 1 réplique/GPU, donc 1× GPU ≥ 48 Go) |
| 70B | ~40 Go | ~80 Go+ | H100 80 Go (limite) |

Rappel : Unsloth OSS = **une réplique complète par GPU** (DDP, pas de sharding FSDP/ZeRO) — le modèle + l'entraînement doivent tenir sur **UNE** carte. Si ça ne passe pas : réduis `max_seq_length`, `r`, le batch, ou passe au FP8 (RTX 40+/H100 : -60 % VRAM, 1,4x plus vite selon Unsloth).

### 162.8. Workflow complet : du dataset au modèle servi

```
1. Curation : 500–2000 paires Q/R métier au format ChatML/ShareGPT/Alpaca
   (tickets résolus, procédures validées — PAS du scraping brut)
2. Split : 90 % train / 10 % éval (tes 30 questions d'éval, section 127)
3. SFT Unsloth : 1–3 époques, LR 2e-4, r=16, packing=True
4. Mesure : éval avant/après — si pas de gain net, jette l'adaptateur
5. (Optionnel) GRPO : affine le comportement avec une fonction de récompense
6. Export : GGUF q4_k_m → Ollama / LM Studio (test humain)
7. Prod : LoRA mergé en 16 bits → vLLM/SGLang avec --enable-lora (hot-swap d'adaptateurs)
```

Le hot-swap LoRA de vLLM/SGLang mérite d'être connu : un seul modèle de base servi, plusieurs adaptateurs métier commutables par requête — ton « modèle onduleurs » et ton « modèle copieurs » sur la même GPU, sans recharger les poids.

---

## 163. SGLang : le serveur d'inférence qui monte (état sept 2026)

**En une phrase :** le challenger open source de vLLM pour servir des LLM en production — son arme : **RadixAttention** (cache de préfixes partagés entre requêtes) et un support day-one agressif des gros MoE (DeepSeek, Qwen, Kimi), avec des gains massifs sur le trafic à préfixes répétés… et presque aucun sur le trafic à prompts uniques.

### 163.1. Présentation

