---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-17
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Hugging Face", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2025-12-15", "2026-07-21", "2026-09", "2026-09-27"]
keywords: ["agent", "agentic", "attention", "aws", "bedrock", "benchmarks", "fine-tuning", "fp8", "gguf", "gpu", "int4", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1407, 1498]
sha256: 769dd06402bac3ff3ec422fb34d8a22277f6937a1261baf1931d79f3a1ce4a38
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Sortie | **21 juillet 2026** (checkpoint mis à jour en août 2026 — re-télécharger les poids) |
| Statut | **disponible** — modèle le plus capable / le plus récent de Poolside au 27 sept. 2026 (aucune sortie ultérieure trouvée) |
| Architecture | **MoE : 117,6B totaux / 8,5B actifs par token** (annoncé 118B / 8B), 256 experts + 1 expert partagé, 48 couches (36 SWA fenêtre 512 avec per-head gating + 12 globales, ratio 3:1) ; optimiseur **Muon** ; pré-entraînement commencé le **22 mai 2026** sur **4 096 GPU NVIDIA H200** ; du début du training au lancement : <9 semaines ; **premier modèle Poolside où le RL tourne en précision FP8** |
| Contexte | **1 048 576 tokens** (1M, réductible à 262 144 côté serving gratuit OpenRouter) |
| Features | interleaved thinking (raisonnement entre tool calls), thinking activable/désactivable par requête ; trajectoires d'évaluation publiées (trajectories.poolside.ai) ; trajectoire Erdős Problem #397 publiée ; agent `pool` (v1.0.14, 21/07/2026) |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | **ouverts — OpenMDW-1.1** (licence Linux Foundation) ; formats : BF16, FP8, INT4, NVFP4 (~71 Go), conversions GGUF + MLX officielles, DFlash drafts ; tourne sur un seul NVIDIA DGX Spark |
| KV cache / optimisations | KV cache **FP8** ; SWA per-head gating pour réduire la mémoire KV |
| Benchmarks (vendor, harness `pool`) | Terminal-Bench 2.1 : 70,2 % ; SWE-Bench Multilingual : 78,5 % ; SWE-Bench Pro (public) : 59,4 % ; DeepSWE v1.1 : 40,4 % ; Toolathlon Verified : 49,7 % ; SWE Atlas (Codebase QnA) : 46,2 % |
| Prix API | ~$0,09 / $0,18 (appelant + OpenRouter, 10 % off) ; free tier OpenRouter sans date de retrait connue |
| Déploiement | Hugging Face, Ollama, llama.cpp (BF16, Q4_K_M), API Poolside, OpenRouter, Vercel AI Gateway |

**« Malibu » (famille de modèles Poolside ?) — NON TROUVÉ ❌** : aucune trace d'une famille de modèles « Malibu » chez Poolside au 27 sept. 2026. Une note de recherche (août 2026) indique explicitement : « no Malibu family found for CLI/API use ». Requêtes tentées : « Poolside AI code model releases 2026 » ; « Poolside new model August September 2026 Laguna S 3.0 Malibu ».

---

# PARTIE XI — NVIDIA Nemotron + NousResearch Hermes

## 125. NVIDIA Nemotron 3 : le contexte

NVIDIA n'est plus seulement le vendeur de pelles : la famille **Nemotron 3** (annoncée le 15/12/2025 ; Nano immédiat, Super/Ultra teasés H1 2026) est une lignée **hybride Mamba-Transformer MoE** open-weight, pensée pour l'agentic reasoning efficient. Particularités : **Mamba-2** dans la boucle (KV cache réduit), **NVFP4** natif, **MTP** (multi-token prediction) pour le speculative decoding, recettes d'entraînement complètes publiées (NeMo). Attention : les licences varient par repo (voir § 133).

## 126. Nemotron 3 Nano 30B-A3B — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | `NVIDIA-Nemotron-3-Nano-30B-A3B` (variantes : `-Base-BF16`, `-BF16`, `-FP8`) |
| Sortie | **15 décembre 2025** (annonce famille Nemotron 3 ; Nano shippe immédiatement) |
| Statut | disponible |
| Architecture | **hybrid MoE Mamba-Transformer** : **30B total / 3,2–3,6B actifs** par forward (3,2B sur les fiches HF, 3,6B dans une recette de training — divergence légère) ; 23 couches Mamba-2 + couches MoE + 6 couches self-attention ; LatentMoE / MTP selon les fiches |
| Contexte | **1M natif** (262K sur HF selon une table — la fiche HF limite à 262K ; NVIDIA annonce 1M ; Token Factory listé à 262K en sept. 2026) |
| Features | agentic reasoning, tool calling (OpenAI-style, 0 argument malformé sur 521 tool calls dans un test tiers), budget de reasoning configurable, JSON-schema output |
| Fine-tuning | **oui** — weights + datasets + recettes d'entraînement publiés (NeMo), checkpoint Base disponible |
| Poids | ouverts — **NVIDIA Open Model License** (version Nemotron) ; BF16 / FP8 / GGUF |
| KV cache / optimisations | 25T tokens de pré-entraînement ; ~3,3× le throughput de modèles de taille similaire |
| Benchmarks clés | pas de chiffres détaillés trouvés dans les sources (vendu comme « agentic reasoning »). Nano Omni (dérivé multimodal) : max output 20 480 tokens en thinking, 1 024 en instruct |
| Prix API | Token Factory (sept. 2026) : **$0,06 / 1M in, $0,24 / 1M out** |
| Déploiement | vLLM, SGLang, TensorRT-LLM, Ollama, llama.cpp, NIM, OpenRouter |

## 127. Nemotron 3 Super 120B-A12B — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | `nvidia/nemotron-3-super-120b-a12b` (HF : `nvidia/Nemotron-3-Super-49B-v1` pour une variante — à ne pas confondre, voir § 133) |
| Sortie | **11 mars 2026** |
| Statut | disponible |
| Architecture | **hybrid Mamba Latent-MoE Transformer** : **120B total / 12B actifs** (recettes NeMo : 120,6B / 12,7B). **NVFP4, LatentMoE, MTP**. Pipeline RL multi-stage : 3× RLVR + 2× SWE-RL + RLHF sur 21 environnements de récompense, GRPO asynchrone à l'échelle 1K GPU |
| Contexte | **1M natif** (256K–262K sur HF) |
| Features | frontier reasoning, coding, agentic multi-agent enterprise, tool calling, structured output via `guided_json` (xgrammar/outlines) recommandé par NVIDIA |
| Fine-tuning | **oui** — recettes complètes publiées (NeMo), training data ouverte |
| Poids | ouverts — **NVIDIA Nemotron Open Model License** ; variantes BF16 / FP8 / NVFP4 / GGUF |
| KV cache / optimisations | NVFP4 natif ; MTP pour speculative decoding |
| Benchmarks clés | pas de chiffres vérifiés dans les sources trouvées (tech report NVIDIA publié — chiffres **non extraits, volontairement non inventés**). AWS Bedrock le liste comme `nvidia.nemotron-super-3-120b` |
| Prix API | Token Factory (sept. 2026) : **$0,30 / 1M in, $0,90 / 1M out** |
| Déploiement | vLLM (0.17.1 cité), SGLang, TensorRT-LLM (1.3.0rc5), Ollama, llama.cpp, NIM, AWS Bedrock, OpenRouter |

## 128. Nemotron-Cascade-2 30B-A3B — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Nemotron-Cascade-2-30B-A3B |
| Sortie | **16–19 mars 2026** (sources divergent : 16 ou 19 mars — jour exact incertain) |
| Statut | disponible |
| Architecture | **30B total / 3B actifs**, MoE hybride (base Nemotron-3 Nano), post-entraîné avec **Cascade RL** (thinking + instruct unifiés) |
| Contexte | **262K** |
| Features | reasoning (thinking) + instruct dans un seul modèle |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | ouverts — **NVIDIA Open Model License** |
| Benchmarks clés | **médaille d'or IMO 2025 et IOI 2025** (revendiqué par NVIDIA — research.nvidia.com/labs/nemotron/nemotron-cascade-2) |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | **non vérifié en détail** — probablement vLLM/SGLang via la famille |

## 129. Nemotron 3 Nano Omni 30B-A3B — fiche ✅ (multimodal)

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Nemotron 3 Nano Omni 30B-A3B |
| Sortie | **28–29 avril 2026** |
| Statut | disponible |
| Architecture | **30B total / 3B actifs**, hybrid Mamba-Transformer MoE **multimodal** |
| Contexte | **256K** |
| Features | **multimodal natif : texte, image, vidéo, audio** — perception multimodale agentique, sub-agent. Max output : 20 480 tokens en thinking, 1 024 en instruct |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | ouverts — licence **NVIDIA Open Model Agreement** (**diffère** de l'Open Model License des autres — à vérifier par repo) |
| Benchmarks clés | **non vérifié au 27/09/2026** |
| Prix API | **non vérifié au 27/09/2026** |
| Déploiement | vLLM, SGLang, NIM |

## 130. Nemotron 3 Ultra 550B-A55B — fiche ✅

