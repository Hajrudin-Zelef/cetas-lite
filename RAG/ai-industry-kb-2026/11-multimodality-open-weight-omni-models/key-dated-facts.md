---
id: ai-industry-kb-2026/11-multimodality-open-weight-omni-models/key-dated-facts
title: "Key dated facts"
domain: multimodality-open-weight-omni-models
role: deep-dive
task: multimodal
actors: ["AMD", "Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "ExploitGym", "Google", "Huawei", "Hugging Face", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "SGLang", "StepFun", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2025-09-23", "2026-01-13", "2026-03-30", "2026-04-02", "2026-06", "2026-06-01", "2026-06-03", "2026-07-31", "2026-08-02", "2026-08-11", "2026-08-20", "2026-08-26", "2026-08-27", "2026-09-10", "2026-09-11", "2026-09-15", "2026-09-17", "2026-09-18", "2026-09-21", "2026-09-22"]
keywords: ["agent", "agentic", "apache", "ascend", "awq", "benchmark", "benchmarks", "claude", "consumer", "decode", "deepseek", "diffusion"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5893, 5982]
section: "11. Multimodality — Open-Weight Omni Models"
sha256: f676eaa2fbb448e4ee21b39e50c8001d49c6526367e076da2e154bf05c702a44
---

# Key dated facts

## Key dated facts

- **2025-09-23/25** — Alibaba releases **Qwen3-Omni-30B-A3B** (Apache 2.0), text/image/audio/video in, streaming speech out, Thinker-Talker MoE (30B/3B); ≥48 GB GPU to serve.
- **2025-11** — **vLLM-Omni** released by the vLLM community (paper arXiv 2602.02204) — first dedicated omni serving framework.
- **2026-01-13** — StepFun releases **Step3-VL-10B** (dense 10.2B VLM, tech report arXiv 2601.09668).
- **2026-01** — vLLM-Omni **v0.14.0**, first stable release.
- **2026-02** — vLLM-Omni **v0.16.0** rebased on upstream vLLM v0.16.0; covers Qwen3-Omni/Qwen3-TTS, MiMo-Audio, diffusion stack, CUDA/ROCm/NPU/XPU.
- **2026-03-30** — Alibaba releases **Qwen3.5-Omni** (Plus/Flash/Light; Light dense open weights): 256K ctx, 113 languages ASR, claimed SOTA on 215 audio/video benchmarks [VENDOR].
- **2026-04-02** — Google DeepMind releases **Gemma 4** under Apache 2.0: E2B/E4B with native audio input (first at that size), 26B-A4B MoE, 31B (LMArena #3 open).
- **2026-06-01** — MiniMax releases **MiniMax M3** (428B/23B, native text/image/video from step 0, 1M ctx, license disputed).
- **2026-06-03** — Google releases **Gemma 4 12B Unified** (Apache 2.0) — claimed encoder-free text/image/audio/video on 16 GB laptops. [UNVERIFIED coverage.]
- **2026-07-31** — MiniMax H3 API launch (model ID `MiniMax-H3`, Hailuo AI app); a model catalog note dated September 7 corrects older "MiniMax Open Weights" labels on H3 entries to territorially-restricted.
- **2026-08-02/03** — MiniMax H3 open weights land on Hugging Face (sources differ by one day); day-0 ComfyUI support; ByteDance Seedance 2.5 (closed, 30s clips) reported the same week by a single secondary source. [UNVERIFIED for Seedance.]
- **2026-08-11** — H3 official repo documents BF16 checkpoints, SGLang/vLLM/Diffusers/ComfyUI reference paths, and a portable `h3-prompt-writing` agent skill.
- **2026-08-20** — Stealth model `ox-alpha` appears free on OpenRouter for one week.
- **2026-08-26** — Zhipu identifies Ox Alpha as **GLM-5.3-Flash** (Bloomberg); trial week served 100 trillion tokens/day on ~100,000 domestic Chinese chips (likely Huawei Ascend).
- **2026-08-27** — Zhipu releases GLM-5.3-Flash weights under **MIT** — first native multimodal GLM (text/image/video; audio inconsistent in coverage).
- **2026-09-10** — **DeepSeek V4.1-Flash GA**: MIT open weights (48 safetensors shards, ~510 GB), MXFP4 KV at 890 bytes/token, built-in reasoning dial, new price card; legacy `deepseek-v4-flash(-vision-exp)` model IDs begin routing to it.
- **2026-09-11** — Cognition's SWE-2 (post-trained from Kimi K3 2.8T) launches; agent-model context for §13.
- **2026-09-15** — fal's **H3 Max** 75% launch discount expires: 5s 768p clip price $0.10→$0.40.
- **2026-09-17** — **HiDream-O1-Video-1.0** announced (native omnimodal video + synchronized audio; company-reported; no independent reproduction as of September 22, 2026). [VENDOR/UNVERIFIED.]
- **2026-09-18** — Alibaba releases **Qwen3.8-Omni-Flash** (API-only): 1M ctx, agentic video perception, OmniVideoBench 63.4→67.8 at −45.7% tokens [VENDOR]; the one omni flagship with no open weights.
- **2026-09-21** — Xiaomi publishes **MiMo-V2.6 Pro (1.02T/42B)** and **Flash (309B/15B)** checkpoints (MIT) with livestreamed RL run ($850K / $2.62M, 60K+ sandboxes).
- **2026-09-22** — Consolidation point: MiMo-V2.6 training environments/RL code promised open-source but not yet shipped; H3 territorial license stands; benchmark independence for MiMo-V2.6 still pending.

## Figures and metrics

### Model scales and footprints

| Model | Total / active params | Context | Weights footprint |
|---|---|---|---|
| Qwen3-Omni-30B-A3B | 30B / 3B | — | ≥48 GB GPU (vLLM-Omni TP2) |
| Qwen3.5-Omni Plus | 30B / 3B | 256K | multi-GPU |
| Qwen3.8-Omni-Flash | — | 1M (991K in / 131K out) | API-only |
| MiniMax M3 | 428B / 23B | 1M | MXFP8 variant; data-center class |
| MiMo-V2.6-Pro | 1.02T / 42B | 1M (128K out) | — (multi-node; API $0.435/$0.87 per 1M) |
| MiMo-V2.6-Flash | 309B / 15B | 1M | 172.9 GB FP8 e4m3, 65 shards (~1 byte/param) |
| Gemma 4 E2B | ~2.3B eff | 128K | QAT 4.3 GB (phone/browser) |
| Gemma 4 E4B | ~4.5B eff | 128K | QAT ~4.3 GB (laptop) |
| Gemma 4 26B-A4B | 26B / ~4B | 256K | single consumer GPU |
| Gemma 4 31B | 30.7B | 256K | workstation GPU |
| Gemma 4 12B Unified | 12B | — | 16 GB laptop |
| GLM-5.3-Flash | 320B / 18B | 1M | vLLM/SGLang/KTransformers |
| Kimi K3 | 2.8T / 104B | 1M | MXFP4 ~1.4 TB resident; 64+ accelerators |
| DeepSeek V4.1-Flash | 552B (+196B Engram; ~748B total) / 8B prefill, 16B decode | 1M | 48 shards ~510 GB; MXFP4 KV 890 bytes/token |
| MiniMax H3 | 33B dense | — | GGUF Q4 14.2 GB; NVFP4 single RTX 5090 |
| Qwen3-VL-8B | 8B | 256K | 4-bit ~6 GB (MLX); AWQ consumer GPU |
| Qwen3-VL-30B-A3B | 30B / 3B | 256K | AWQ-4bit ~19 GB disk; ~10–12 GiB/GPU at TP2 |
| Ministral 3 8B | 8B | 256K | 6.0 GB at 8-bit |

### Benchmark scores (vendor vs independent, source-typed)

| Benchmark | Model | Score | Source type |
|---|---|---|---|
| 36 audio/audio-visual benchmarks | Qwen3-Omni-30B-A3B | overall SOTA 22/36; open SOTA 32/36 | [VENDOR] (technical report, Sept 2025) |
| 215 audio/video benchmarks | Qwen3.5-Omni | claimed SOTA | [VENDOR] (Alibaba, Mar 2026) |
| OmniVideoBench | Qwen3.8-Omni-Flash (agentic perception) | 63.4 → 67.8, −45.7% tokens | [VENDOR] (Sept 18, 2026) |
| SWE-bench Verified | MiniMax M3 | 80.5% | [VENDOR] (June 2026) |
| SWE-bench Pro | MiniMax M3 | 59% | [VENDOR] (June 2026) |
| Terminal-Bench 2.1 | MiniMax M3 | 66% | [VENDOR] (June 2026) |
| LMArena | MiniMax M3 | 1340 | [VENDOR] (June 2026) |
| Intelligence Index v4.3 | MiMo-V2.6-Pro | 46.32 (claimed top open) | [VENDOR] (Sept 21, 2026) |
| Intelligence Index v4.3 | Kimi K3 | 60 (#1 open at publication) | [VENDOR] (mid-2026) |
| DeepSWE v1.1 | MiMo-V2.6-Pro 71.9 / Flash 67.9 vs Claude Opus 5 74.0 | — | [VENDOR] (Sept 21, 2026) |
| DeepSWE v1.1 | MiniMax M3 | 65.68 | [VENDOR] (June 2026) |
| Terminal-Bench 2.1 | MiMo-V2.6-Flash | 87.6 vs Opus 5 89.1 | [VENDOR] (Sept 21, 2026) |
| AutomationBench | MiMo-V2.6-Flash | 52.3 vs Opus 5 50.3 / GPT-5.6 Sol 45.8 | [VENDOR] (Sept 21, 2026) |
| MiMo VisualCoding | MiMo-V2.6-Pro 72.3 / Flash 71.5 vs Opus 5 70.0 | — | [VENDOR] (Sept 21, 2026) |
| CyberGym | MiMo-V2.6-Flash | 95.1 (ExploitGym 6.0, ExploitBench 25.3) | [VENDOR] (Sept 21, 2026) |
| DeepSWE v1.1 | DeepSeek V4.1-Flash | 74.2 (TB4.0 31.2; HLE 36.8 vs Opus 5 51.8/56.3) | [VENDOR] (Sept 10, 2026) |
| TB 2.1 | DeepSeek V4.1-Flash | 90.6 | [VENDOR] |
| GPQA Diamond | DeepSeek V4.1-Flash | 90.9 | [VENDOR] |
| Codeforces | DeepSeek V4.1-Flash | 3471 | [VENDOR] |
| VTB (visual agentic tasks) | GPT-5-think 18.44% vs open 1.16–1.65% | 10–17× gap | Independent (Scale Labs) |
| DocVQA | Qwen3-VL-8B | 96.1% | [VENDOR] |
| OCRBench | Qwen3-VL-8B | ~896 | [VENDOR] |
| ScreenSpot (UI grounding) | Qwen3-VL-8B | 94.4% | [VENDOR] |
| DeepSWE | GLM-5.3-Flash | 63.4 | [VENDOR] (model card, Aug 2026) |
| SWE-bench Pro | GLM-5.3 (744B/40B) | 77.8% | [VENDOR] (Feb 2026) |
| 28 practical tasks, rubric 9.3/10 | GLM-5.3-Flash | 100% pass rate, $0.28 vs $1.43 GPT-5.5 | Independent (Ed Yau, Aug 23, 2026) |
| Arena AI (text) | Gemma 4 31B | #3 open globally | Independent (LMArena) |
| Codeforces ELO | Gemma 4 | 110 → 2150 reported | [VENDOR] |
| LiveBench Coding / Agentic Coding | Kimi K3 | 81.45 / 57.58 | [VENDOR] (mid-2026) |
| Video Editing (With Audio), Artificial Analysis | MiniMax H3 | #1 (first open model to top a ranking) | [COMMUNITY] (leaderboard) |
| Text-to-Video / Image-to-Video, Artificial Analysis | MiniMax H3 | #2 / #3 | [COMMUNITY] |
| Image-to-video Elo | MiniMax H3 | ~1476 (Arena.ai) | [COMMUNITY] |
| Image-to-Video (With Audio), Artificial Analysis | HiDream-O1-Video-1.0 | #4 claimed | [VENDOR] (unreproduced) |

### Training and inference economics

