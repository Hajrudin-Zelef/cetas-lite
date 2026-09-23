---
id: collect-huggingface/huggingface/moonshotai-kimi-vl-a3b-thinking
title: "Kimi-VL-A3B-Thinking - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Hugging Face", "Moonshot", "OpenAI", "vLLM"]
dates: ["2026-09-23"]
keywords: ["kimi", "agent", "benchmarks", "context window", "fine-tuning", "inference", "license", "mit license", "moe", "multimodal", "open-weight", "parameters"]
source: docs/RAG/Collect RAG/03_huggingface/moonshotai-Kimi-VL-A3B-Thinking.md
source_anchor: ""
source_lines: [1, 54]
sha256: 0a1fa328c40eb6a42445b50ea3946ed5300fae6d0fd36744456e7cddec9aea52
---

# Kimi-VL-A3B-Thinking - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/moonshotai/Kimi-VL-A3B-Thinking
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Kimi-VL is an efficient open-source Mixture-of-Experts (MoE) vision-language model (VLM) from Moonshot AI, offering advanced multimodal reasoning, long-context understanding, and strong agent capabilities while activating only 2.8B parameters in its language decoder. Kimi-VL-A3B-Thinking is the long-thinking variant, developed through long chain-of-thought (CoT) supervised fine-tuning (SFT) and reinforcement learning (RL). The architecture combines an MoE language model (16B total / 3B activated), a native-resolution vision encoder (MoonViT), and an MLP projector, with a 128K extended context window. The base is Moonlight-16B-A3B.

As a general-purpose VLM it excels at multi-turn agent interaction (e.g., OSWorld, state-of-the-art comparable to flagship models), college-level image and video comprehension, OCR, mathematical reasoning, and multi-image understanding, competing with GPT-4o-mini, Qwen2.5-VL-7B, and Gemma-3-12B-IT while surpassing GPT-4o in several specialized domains. With the 128K window it achieves 64.5 on LongVideoBench and 35.1 on MMLongBench-Doc; MoonViT native resolution enables 83.2 on InfoVQA and 34.5 on ScreenSpot-Pro. The Thinking variant scores 61.7 on MMMU, 36.8 on MathVision (full), and 71.3 on MathVista (mini), matching 30B/70B frontier open-source VLMs on MathVision.

Recommended sampling: Temperature = 0.8 for thinking models (0.2 for Instruct). Inference via Hugging Face Transformers (python=3.10, torch>=2.1.0, transformers=4.48.2, trust_remote_code=True) and vLLM (via an open MR). A newer version, Kimi-VL-A3B-Thinking-2506, offers improved general visual understanding, reasoning, video, and agent abilities. Licensed under MIT; hub reports 16B params (BF16), 16,461 monthly downloads, 4 community quantizations. Technical report: arXiv 2504.07491.

## Key points

- Efficient MoE VLM: 16B total / 3B activated (2.8B active LLM decoder) parameters, 128K context.
- Long-thinking variant trained with CoT SFT + RL for long-horizon multimodal reasoning.
- Native-resolution MoonViT vision encoder enables ultra-high-resolution input understanding.
- Competitive with GPT-4o-mini, Qwen2.5-VL-7B, Gemma-3-12B; beats GPT-4o on some specialized tasks.
- Thinking benchmarks: MMMU 61.7, MathVision 36.8, MathVista 71.3.
- Agent benchmarks: OSWorld SOTA; LongVideoBench 64.5, InfoVQA 83.2.
- MIT license; use Temperature 0.8 for thinking models.
- Newer 2506 version available with improved video/agent abilities.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | moonshotai |
| Model name | Kimi-VL-A3B-Thinking |
| Architecture | MoE VLM (MoonViT + MLP projector + MoE LLM) |
| Total params | 16B |
| Activated params | 3B (2.8B active decoder) |
| Context length | 128K |
| Base model | moonshotai/Moonlight-16B-A3B |
| Vision encoder | MoonViT (native resolution) |
| Training | CoT SFT + RL (long thinking) |
| License | MIT |
| Weight formats | BF16 |
| Benchmarks | MMMU 61.7, MathVision 36.8, MathVista 71.3, LongVideoBench 64.5, InfoVQA 83.2 |
| Sampling | Temperature 0.8 (thinking) |
| Requirements | python=3.10, torch>=2.1.0, transformers=4.48.2 |
| Paper | arXiv 2504.07491 |
| Downloads/month | 16,461 |

## Why this source matters for the RAG

This card documents a compact open-weight multimodal thinking model, showing how a 3B-active MoE VLM can reach frontier-level reasoning, long-context, and agent performance. It is essential reference for retrieval on efficient VLMs, multimodal RL/CoT training, and native-resolution vision encoders.
