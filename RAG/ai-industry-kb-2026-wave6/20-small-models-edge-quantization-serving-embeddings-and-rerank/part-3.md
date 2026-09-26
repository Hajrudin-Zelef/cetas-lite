---
id: ai-industry-kb-2026-wave6/20-small-models-edge-quantization-serving-embeddings-and-rerank/part-3
title: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 3)"
domain: small-models-edge-quantization-serving-embeddings-and-rerank
role: deep-dive
task: quantization
actors: ["Alibaba", "China", "Google", "Hugging Face", "Microsoft", "vLLM"]
dates: ["2024-08", "2025-02-26", "2025-03-13", "2025-04-29", "2025-04-30", "2025-07-08", "2026-09"]
keywords: ["apache", "attention", "benchmarks", "claude", "copilot", "cost", "deepseek", "fine-tuning", "gguf", "gpus", "gqa", "inference"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9692, 9719]
section: "§20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers"
delta_of: ai-industry-kb-2026
sha256: c41c64719836add47d7850ec42bb0ada877dbe99f68e406973d18e76e0531028
---

# §20. Small Models, Edge, Quantization, Serving, Embeddings, and Rerankers (part 3)

- The original Qwen3 dense family launched 2025-04-29 in six sizes: 0.6B, 1.7B, 4B, 8B, 14B, 32B, plus the 30B-A3B MoE — all under Apache 2.0 [SECONDARY]. Source: http://dev.to/best_codes/qwen-3-benchmarks-comparisons-model-specifications-and-more-4hoa
- Qwen3-0.6B, 1.7B, and 4B ship with 32K native context; Qwen3-8B, 14B, and 32B ship with 128K native context per a secondary spec table [SECONDARY]. Source: http://dev.to/best_codes/qwen-3-benchmarks-comparisons-model-specifications-and-more-4hoa
- Qwen3-30B-A3B: 30.5B total / 3.3B active parameters, 48 layers, 128 experts with 8 active per token, native 32,768 context extendable to 131,072 via YaRN [SECONDARY]. Sources: https://huggingface.co/unsloth/Qwen3-30B-A3B-GGUF and https://huggingface.co/Qwen/Qwen3-30B-A3B-MLX-4bit
- Qwen3-Coder-30B-A3B is a distinct variant with 262,144 native context and a non-thinking mode — those specs must not be transferred to base Qwen3-30B-A3B [SECONDARY]. Source: https://huggingface.co/unsloth/Qwen3-Coder-30B-A3B-Instruct-1M-GGUF
- Gemma 3 launched 2025-03-13 in 1B, 4B, 12B, and 27B sizes; the 270M variant was added later [SECONDARY]. Source: https://themunicheye.com/google-launches-gemma-3-open-source-ai-model-13085
- Gemma 3 context split: 1B and 270M ship 32K; 4B, 12B, and 27B ship 128K [SECONDARY]. Source: https://huggingface.co/google/gemma-3-270m-it
- Gemma 3 training-token figures from the model card: 270M on 6T, 1B on 2T, 4B on 4T, 12B on 12T, 27B on 14T tokens; knowledge cutoff August 2024 [SECONDARY]. Source: https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_3.md
- Gemma 3 architectural details: QK-norm, 5:1 interleaved local/global attention, multimodal image input — the small-line template Google later reused [SECONDARY]. Source: https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_3.md
- Phi-4 (14B) launched late 2024, trained on 9.8 trillion tokens over three weeks on synthetic data plus curated web content, under MIT license [SECONDARY]. Sources: https://techcommunity.microsoft.com/blog/educatordeveloperblog/phi-4-small-language-models-that-pack-a-punch/4464167 and https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- Phi-4-mini (3.8B) and Phi-4-multimodal (5.6B) were announced February 26, 2025: mini carries a 200K-token vocabulary, grouped-query attention, and 128K context; multimodal handles text, vision, and speech via a mixture-of-LoRAs architecture with 128K context [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- Phi-4-reasoning (14B) launched April 30, 2025: supervised fine-tuning on chain-of-thought traces, 32K context; Phi-4-reasoning-plus (14B) adds outcome-based reinforcement learning, ~1.5x longer reasoning traces, and effective handling up to 64K [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- Phi-4-reasoning-vision-15B shipped in 2026 as the newest Phi family member, adding vision to the reasoning line [SECONDARY]. Source: https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- Phi Silica is the Copilot+ PC on-device variant of the Phi family, optimized for NPU inference on Windows [SECONDARY]. Source: https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- Phi-4-mini was optimized for MediaTek NPU hardware, running at over 800 tokens per second on the Dimensity 9400 chip [SECONDARY]. Source: https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- Rho-alpha (ρα), announced 2026, is Microsoft's first robotics model built from the Phi series: natural-language commands to control signals, bimanual manipulation focus, tactile sensing in the perception stack [SECONDARY]. Source: https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- No Phi-5 evidence: as of September 2026 the Phi line tops out at Phi-4-reasoning-vision-15B and Rho-alpha; no credible source documents a Phi-5 release [DIRECTIONAL — absence of evidence across the searched corpus].
- Phi-3/Phi-4 config quirk: both report model_type "phi3" — Microsoft kept the family identifier across major versions; Phi-4-mini additionally uses partial_rotary_factor 0.75 (trailing 25% of each head not rotated) [COMMUNITY]. Source: https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- A production Phi-4-reasoning deployment note (September 2026) shows it served on a whole L40S via vLLM v0.20.2 with --reasoning-parser=deepseek_r1 (--enable-reasoning was removed in v0.20.2), 32K context, always-on chain-of-thought with no Qwen-style enable_thinking=false [COMMUNITY]. Source: https://github.com/ualberta-rcg/aleph/blob/HEAD/models/phi-4-reasoning/CLAUDE.md
- SmolLM3-3B released July 8, 2025 by Hugging Face under Apache 2.0: 3.08B parameters, decoder-only transformer with GQA and NoPE at 3:1 ratio [SECONDARY]. Sources: https://atomic.chat/models/smollm3-3b and https://github.com/huggingface/blog/blob/HEAD/smollm3.md
- SmolLM3 pretraining: 11.2T tokens in three stages — Stage 1 (0→8T): 85% web (FineWeb-Edu, DCLM, FineWeb2), 12% code, 3% math; Stage 2 (8→10T): math up to 10%, code up to 15%; Stage 3 (10→11.1T): math 13%, code 24% [SECONDARY]. Source: https://www.xugj520.cn/en/archives/smollm3-compact-multilingual-ai.html
- SmolLM3 context: 64K trained, up to 128K via YaRN rope scaling; NoPE (rotary removed every 4th layer) plus intra-document masking (32% faster long-sequence convergence) [SECONDARY]. Sources: https://atomic.chat/models/smollm3-3b and https://www.xugj520.cn/en/archives/smollm3-compact-multilingual-ai.html
- SmolLM3 dual reasoning: extended thinking on by default, /no_think flag for fast direct replies; post-training was midtraining on 140B reasoning tokens then SFT and Anchored Preference Optimization [SECONDARY]. Sources: https://atomic.chat/models/smollm3-3b and https://github.com/huggingface/blog/blob/HEAD/smollm3.md
- SmolLM3 languages: six native (English, French, Spanish, German, Italian, Portuguese); Arabic, Chinese, Russian seen with fewer tokens — treat as secondary [SECONDARY]. Source: https://atomic.chat/models/smollm3-3b
- SmolLM3 benchmarks by mode (think vs no-think): AIME 2025 36.7% vs 9.3%; LiveCodeBench 30.0% vs 15.2%; GPQA Diamond 41.7% vs 35.7% — thinking roughly doubles to quadruples scores [SECONDARY]. Source: https://www.xugj520.cn/en/archives/smollm3-compact-multilingual-ai.html
- SmolLM3 competitive position: outperforms Llama 3.2 3B and Qwen2.5 3B, stays competitive with 4B alternatives (Qwen3, Gemma3); fully open — weights plus complete training recipe including data mixtures [SECONDARY]. Source: https://github.com/huggingface/smollm/
- SmolLM3 minimum hardware: 2GB RAM via UD-IQ2_M GGUF (1.15GB file); recommended sampling temperature 0.6, top_p 0.95; local runtimes: llama.cpp, ONNX, MLX, MLC, ExecuTorch [SECONDARY]. Source: https://atomic.chat/models/smollm3-3b
- SmolLM3 training cost: 384x H100 GPUs for 24 days, 2.36M-token global batch, 4,096 sequence length, LR 2e-4 with Warmup-Stable-Decay [SECONDARY]. Source: https://www.xugj520.cn/en/archives/smollm3-compact-multilingual-ai.html
- SmolLM3 tool calling: two formats — JSON blobs inside XML tags, or Python-style function calls in a code snippet [SECONDARY]. Source: https://atomic.chat/models/smollm3-3b
