---
id: collect-huggingface/huggingface/unsloth-kimi-k2-7-code-gguf
title: "Kimi K2.7 Code GGUF - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-09-23"]
keywords: ["gguf", "kimi", "agent", "agentic", "attention", "benchmark", "benchmarks", "claude", "int4", "license", "llama", "llama.cpp"]
source: docs/RAG/Collect RAG/03_huggingface/unsloth-Kimi-K2.7-Code-GGUF.md
source_anchor: ""
source_lines: [1, 48]
sha256: 2386644ab569dafcdb2514b3a66d9b18cf3a3ca2fed8a9fdca60999ae01809c9
---

# Kimi K2.7 Code GGUF - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/unsloth/Kimi-K2.7-Code-GGUF
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

unsloth/Kimi-K2.7-Code-GGUF is Unsloth's GGUF conversion (Unsloth Dynamic 2.0 quantization) of Moonshot AI's Kimi K2.7 Code, a coding-focused agentic model built upon Kimi K2.6. The embedded model card describes Kimi K2.7 Code as a 1T-total/32B-active MoE: 61 layers (1 dense), MLA attention (hidden 7168, 64 heads), MoE hidden 2048 per expert, 384 experts with 8 selected per token plus 1 shared expert, SwiGLU, 160K vocabulary, 256K context, and a MoonViT vision encoder (400M). It strengthens end-to-end completion of complex software engineering workflows while reducing thinking-token usage by ~30% compared to Kimi K2.6, and it forces thinking and `preserve_thinking` modes (instant mode unsupported). Native INT4 quantization follows the K2-Thinking method. License: Modified MIT.

Benchmark improvements over K2.6 (vs GPT-5.5 and Claude Opus 4.8): Kimi Code Bench v2 62.0 (vs 50.9 for K2.6), Program Bench 53.6 (48.3), MLS Bench Lite 35.1 (26.7), Kimi Claw 24/7 Bench 46.9 (42.9), MCP Atlas 76.0 (69.4), MCP Mark Verified 81.1 (72.8). Footnotes detail the in-house Kimi Code Bench V2 (10+ languages, production incidents), Program Bench (recreate behavior from compiled binary against 248,000 fuzz tests), MLS-Bench-Lite (30 tasks, 5-hour exploration), Kimi Claw 24/7 Bench (17 scenarios, 610 points, OpenClaw harness, avg of 3 runs), MCP-Atlas and MCPMark-Verified protocols.

The Unsloth-specific guidance: full-precision lossless is Q8 (UD-Q8_K_XL, 595 GB), only ~10 GB bigger than Q4 (UD-Q4_K_XL, 584 GB). The hardware table lists Dynamic 2.0 quants: 1-bit UD-IQ1_M 304 GB; 2-bit UD-IQ2_XXS 318 GB, UD-IQ2_M 318 GB, UD-Q2_K_XL 339 GB; 3-bit UD-IQ3_S 419 GB, UD-Q3_K_M 464 GB, UD-Q3_K_XL 464 GB; 4-bit UD-IQ4_XS 495 GB, UD-Q4_K_XL 584 GB; 8-bit UD-Q8_K_XL 595 GB. Architecture tag is `deepseek2` (like K2.5/K2.6). Usage: llama.cpp (`llama serve -hf unsloth/Kimi-K2.7-Code-GGUF:UD-Q4_K_XL`), Ollama, LM Studio, Jan, vLLM, SGLang, Pi, OpenClaw, Hermes; recommended temperature 1.0, top_p 0.95; video input experimental (official API only). Downloads ~362K/month.

## Key points

- Unsloth Dynamic 2.0 GGUF of Kimi K2.7 Code (1T-A32B, 256K context).
- Coding-focused: +30% token efficiency (fewer thinking tokens) vs K2.6.
- Forces thinking + preserve_thinking; no instant mode.
- Lossless Q8 (UD-Q8_K_XL 595 GB) only ~10 GB bigger than Q4 (584 GB).
- Gains: Kimi Code Bench v2 62.0, MCP Mark Verified 81.1, MCP Atlas 76.0 vs K2.6.
- Modified MIT license; architecture tag deepseek2.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | unsloth |
| Model name | Kimi-K2.7-Code-GGUF |
| Base model | moonshotai/Kimi-K2.7-Code |
| Architecture | MoE, 61 layers, MLA, 384 experts (8/token) |
| Total / active params | 1T / 32B |
| Context length | 256K |
| Vision encoder | MoonViT (400M) |
| Quantization | Native INT4 (source); Dynamic 2.0 GGUFs |
| License | modified-mit |
| Quant sizes | UD-IQ1_M 304 GB; UD-Q2_K_XL 339 GB; UD-Q3_K_XL 464 GB; UD-Q4_K_XL 584 GB; UD-Q8_K_XL 595 GB |
| Key benchmarks | Kimi Code Bench v2 62.0, Program Bench 53.6, MLS Bench Lite 35.1, Kimi Claw 24/7 46.9, MCP Atlas 76.0, MCP Mark Verified 81.1 |
| Downloads/month | ~362,293 |

## Why this source matters for the RAG

This card documents a specialized open-weight coding agentic model (Kimi K2.7 Code) with its Unsloth Dynamic 2.0 GGUF sizes and lossless-Q8 guidance, plus benchmark data against GPT-5.5 and Claude Opus 4.8. It provides citable facts on coding-agent specialization, token-efficiency gains, and local quantization of 1T-class models.
