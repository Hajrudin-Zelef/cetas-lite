---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-18
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agent", "agentic", "agents", "alignment", "attention", "benchmarks", "consumer", "datacenter", "fine-tuning", "fp8", "gguf", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1499, 1562]
sha256: 3ece4f64737340950a8d5709d23e13ec95ffea2a0316ab7f3c54a483d94ee906
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | `nvidia/Nemotron-3-Ultra-550b-a55b` |
| Sortie | **4 juin 2026** |
| Statut | disponible |
| Architecture | **hybrid Mamba-Attention LatentMoE Transformer** : **550B total / 55B actifs**. **MTP, NVFP4** — le plus grand modèle Nemotron 3, pour le datacenter |
| Contexte | **1M** |
| Features | frontier reasoning pour agents long-running : planification multi-turn soutenue, délégation à des sub-agents |
| Fine-tuning | **non vérifié au 27/09/2026** — recette d'entraînement listée dans NeMo (Pretrain → SFT → RLVR → MOPD) |
| Poids | ouverts — **NVIDIA Open Model License** |
| KV cache / optimisations | NVFP4, MTP |
| Benchmarks clés | **non vérifié en chiffres au 27/09/2026** — présenté par NVIDIA comme le tier frontier reasoning au-dessus de Super |
| Prix API | Token Factory (sept. 2026) : **$1,00 / 1M in, $3,00 / 1M out** |
| Déploiement | vLLM, SGLang, TensorRT-LLM, NIM, OpenRouter |

## 131. Nemotron 3.5 Lightning 30B-A3B — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | `nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B` (variantes `-BF16`, `-FP8`, `-NVFP4` ; draft models `-NVFP4-DFlash`, `-NVFP4-DSpark`) |
| Sortie | **11 août 2026** (GA publique, blog NVIDIA). Une note communautaire cite BF16 le 1er août et NVFP4 le 4 août — uploads HF vs annonce, non tranché |
| Statut | disponible |
| Architecture | **MoE hybride Mamba-2 + MoE + attention** : **30B total / 3B actifs**. **MTP** entraîné par continued pre-training sur la base Nano 3. Pré-entraîné sur **>20T tokens** avec une recette NVFP4 (Megatron-LM), puis RL multi-environnements (GRPO asynchrone, NeMo RL / NeMo Gym) |
| Contexte | **jusqu'à 1M** ; NVIDIA documente du single-H100 à 256K |
| Features | exécution agentique « always-on » : tool use, validation, sub-agent. Speculative decoding natif via MTP (1,46–1,96× le throughput en test indépendant Thoughtworks). Reasoning + chat, toggle thinking |
| Fine-tuning | **oui** — base checkpoint publié, recette d'entraînement complète (Pretrain → MTP CPT → SFT → RL → Quantization) |
| Poids | ouverts — **OpenMDW License Agreement v1.1** (permissive, y compris usage commercial/redistribution — rapporté par tech-insider). HF : BF16 (~66 GB), NVFP4 (~22 GB), FP8 ; GGUF communautaires (ggml-org, bartowski, unsloth) |
| KV cache / optimisations | NVFP4 → **~22 GB**, tient sur 1× H100 80GB, 1× A100 80GB, DGX Spark (GB10), RTX 5090 (32 GB) selon config ; MTP speculative decoding ; draft models DFlash (datacenter low-concurrency) et DSpark (DGX Spark) |
| Benchmarks clés | NVIDIA (à vérifier indépendamment) : MMLU Pro **81,94 (BF16) / 81,62 (NVFP4)**, GPQA Diamond **75,44 / 75,57** ; PinchBench agentic : **~86 %** accuracy, 30 % plus rapide que Qwen3.6-35B. Indépendant : OpenRouter mesure GPQA Diamond **63,0–68,8 %** (7–12 pts sous le chiffre NVIDIA) ; Thoughtworks : speculative decoding 1,46–1,96× throughput, coût self-hosted réduit de $0,477 à $0,250 / 1M tokens output |
| Prix API | Token Factory (sept. 2026) : **$0,06 / 1M in, $0,24 / 1M out** |
| Déploiement | vLLM (v0.27.1 avec `--mamba-backend flashinfer`), SGLang, NIM, Ollama/llama.cpp via GGUF communautaires, NVIDIA NIM, OpenRouter |

## 132. Nemotron 3.5 Content Safety — fiche ✅ (guardrail)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Nemotron 3.5 Content Safety |
| Sortie | **4 juin 2026** |
| Statut | disponible |
| Architecture | modèle de **sécurité de contenu** (guardrail) — specs **non vérifiées au 27/09/2026** |
| Contexte | 128K (selon une table tierce) |
| Features | content safety / modération pour les agents |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | ouverts — licence **non vérifiée précisément au 27/09/2026** |
| Benchmarks clés | **non vérifié au 27/09/2026** |
| Prix API | **non vérifié au 27/09/2026** |

*(Nemotron Nano 2 12B — 2 décembre 2025, hybride Transformer-Mamba dense ~12B, 131K contexte, NVIDIA Open Model License — inclus pour mémoire : dernier Nemotron avant la période.)*

## 133. Licences Nemotron : attention, elles varient par repo ⚠️

- Super / Nano / Ultra / Cascade-2 : **NVIDIA Nemotron Open Model License**.
- Nano Omni : **NVIDIA Open Model Agreement** (diffère — à vérifier par repo).
- 3.5 Lightning : **OpenMDW License Agreement v1.1** (permissive, Linux Foundation).
- Ne pas confondre `nvidia/Nemotron-3-Super-49B-v1` (lignée « Llama Nemotron » de mai 2025 : Nano 4B / Super 49B / Ultra 253B, hors période) avec **Nemotron 3 Super 120B-A12B**.

## 134. NousResearch Hermes 4 / Hermes 4.3-36B — fiches ✅

**Hermes 4 (famille)** — sortie **août 2025** (annonce initiale ; paper arXiv:2508.18255) ; dans la période fév–sept 2026 : toujours le modèle chat courant de NousResearch. Variantes : `Hermes-4-70B`, `Hermes-4-405B` ; non-reasoning et reasoning. Architecture : fine-tunes de bases **Llama-3.1-70B et Llama-3.1-405B**, dense transformer, **hybrid reasoning** (mode « deep thinking » explicite `<think>…</think>` / `reasoning_content`), neutral alignment, steerability. Contexte **128K–130K**. Poids ouverts (HF `NousResearch/`, collection `hermes-4-collection`), BF16 + FP8 + GGUF ; fine-tuning : oui (licence permissive). Benchmarks : GPQA Diamond ~83 % pour le 35B (Delphi Digital, marqué estimé). Déploiement : Nous Portal (API OpenAI-compatible), vLLM/SGLang, GGUF, LM Studio. Prix API : **non vérifié au 27/09/2026**.

**Hermes 4.3-36B** — sortie **6 décembre 2025** ; modèle mid courant en 2026. **36B paramètres**, dense (fine-tune), **hybrid reasoning mode**, **entraînement décentralisé sur le réseau Psyche de Nous Research** (revendiqué). Contexte : **jusqu'à 512K** (revendiqué au lancement) vs **128K** dans le catalogue Nous Portal — **divergence non expliquée, non vérifié au 27/09/2026**. Optimisé déploiement local (consumer hardware), user-aligned, minimal content restrictions, steerability. Poids ouverts (BF16 + FP8 + GGUF). Benchmarks : **non vérifiés en chiffres indépendants au 27/09/2026**. Déploiement : Nous Portal, GGUF/llama.cpp, LM Studio, vLLM/SGLang.

## 135. Hermes Agent : le harness auto-évolutif open source ✅

