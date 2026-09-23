---
id: collect-huggingface/huggingface/unsloth-kimi-k2-6-gguf
title: "Kimi K2.6 GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Google", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "context window", "gemini", "int4"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Kimi-K2.6-GGUF.md
source_anchor: ""
source_lines: [1, 49]
sha256: 3f949205bfc62cf24f765623b4e91fa7ca3c53b45a6fde91b9bb777c885be62d
---

# Kimi K2.6 GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Kimi-K2.6-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/Kimi-K2.6-GGUF is Unsloth's GGUF conversion (Unsloth Dynamic 2.0 quantization) of Moonshot AI's Kimi K2.6, an open-source native multimodal agentic model. The card embeds the full Kimi K2.6 model card, so the architecture is the same: a 1T-total/32B-active MoE with 61 layers (1 dense), MLA attention (hidden 7168, 64 heads), MoE hidden 2048 per expert, 384 experts with 8 selected per token plus 1 shared expert, SwiGLU activation, 160K vocabulary, 256K context window, and a MoonViT vision encoder (400M params). It advances long-horizon coding, coding-driven design, proactive autonomous execution, and swarm-based task orchestration (up to 300 sub-agents / 4,000 coordinated steps). Native INT4 quantization follows the K2-Thinking method; license is Modified MIT. The hub architecture tag is `deepseek2` (same as K2.5 line). Downloads ~275K/month.

The card includes the K2.6 benchmark suite vs GPT-5.4, Claude Opus 4.6, Gemini 3.1 Pro and K2.5: AIME 2026 96.4, GPQA-Diamond 90.5, HLE-Full 34.7 (54.0 w/ tools), SWE-bench Verified 80.2, SWE-bench Multilingual 76.7, SWE-Bench Pro 58.6, Terminal-Bench 2.0 66.7, LiveCodeBench v6 89.6, BrowseComp 83.2 (86.3 Agent Swarm), DeepSearchQA F1 92.5, MCPMark 55.9, OSWorld-Verified 73.1, MMMU-Pro 79.4, CharXiv 80.4, MathVision 87.4. Usage: Thinking and Instant modes (temp 1.0 / 0.6, top_p 0.95), `preserve_thinking`, interleaved thinking + multi-step tool calls, image/video input (video experimental, official API only), Kimi Code CLI agent framework. Deployment via vLLM / SGLang / KTransformers reusing the K2.5 method; transformers >=4.57.1,<5.0.0.

The Unsloth-specific guidance: full-precision lossless is Q8 (UD-Q8_K_XL, 595 GB), only ~10 GB bigger than Q4 (UD-Q4_K_XL, 584 GB). Hardware table: 2-bit UD-Q2_K_XL 340 GB; 4-bit UD-Q4_K_XL 584 GB; 8-bit UD-Q8_K_XL 595 GB; 16-bit BF16 2.05 TB. llama.cpp serving is via `llama serve -hf unsloth/Kimi-K2.6-GGUF:UD-Q4_K_XL`, with Ollama, LM Studio, Jan, vLLM, SGLang, Pi, OpenClaw, Hermes, and Docker Model Runner options. Base model is moonshotai/Kimi-K2.6 (47 quantizations). A "How to Run Kimi K2.6" guide is linked. Reference paper: Kimi K2.5 (arXiv 2602.02276).

## Key points

- Unsloth Dynamic 2.0 GGUF of Kimi K2.6 (1T-A32B, 256K context, multimodal).
- Lossless Q8 (UD-Q8_K_XL 595 GB) only ~10 GB bigger than Q4 (584 GB).
- Quant sizes: Q2_K_XL 340 GB, Q4_K_XL 584 GB, Q8_K_XL 595 GB, BF16 2.05 TB.
- Thinking + Instant modes; preserve_thinking; agent swarm (300 sub-agents).
- Modified MIT license; architecture tag deepseek2.
- ~275K downloads/month.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Kimi-K2.6-GGUF |
| Base model | moonshotai/Kimi-K2.6 |
| Architecture | MoE, 61 layers, MLA, 384 experts (8/token) |
| Total / active params | 1T / 32B |
| Context length | 256K |
| Vision encoder | MoonViT (400M) |
| Quantization | Native INT4 (source); Dynamic 2.0 GGUFs |
| License | modified-mit |
| Quant sizes | UD-Q2_K_XL 340 GB, UD-Q4_K_XL 584 GB, UD-Q8_K_XL 595 GB, BF16 2.05 TB |
| Key benchmarks | AIME 2026 96.4, GPQA-D 90.5, SWE-bench Verified 80.2, Terminal-Bench 2.0 66.7, BrowseComp 83.2, DeepSearchQA F1 92.5 |
| Paper | Kimi K2.5 (arXiv 2602.02276) |
| Downloads/month | ~274,978 |

## Why this source matters for the RAG

This card is the practical GGUF companion to the Kimi K2.6 release, documenting exact Dynamic 2.0 quant file sizes and lossless-Q8 guidance for a 1T-class multimodal model, plus local-serving commands and usage modes. It complements the base card with essential, citable quantization and local-deployment data.
