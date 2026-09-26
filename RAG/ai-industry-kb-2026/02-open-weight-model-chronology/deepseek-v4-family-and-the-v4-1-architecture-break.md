---
id: ai-industry-kb-2026/02-open-weight-model-chronology/deepseek-v4-family-and-the-v4-1-architecture-break
title: "DeepSeek — V4 family and the V4.1 architecture break"
domain: open-weight-model-chronology
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "Meta", "Nvidia", "OpenRouter"]
dates: ["2025-04-28", "2025-04-29", "2026-02", "2026-02-10", "2026-02-16", "2026-04-02", "2026-04-16", "2026-04-21", "2026-05-20", "2026-05-21", "2026-07-19", "2026-08-12", "2026-08-14", "2026-09-02", "2026-09-18"]
keywords: ["deepseek", "agentic", "agents", "apache", "attention", "blackwell", "claude", "fable 5", "fp8", "gpus", "kv cache", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [679, 692]
section: "2. Open-Weight Model Chronology"
sha256: f1ecfa243c49d5017eca20812bd36aba6ffe0243c0d20ccd3b0fc04d48918e21
---

# DeepSeek — V4 family and the V4.1 architecture break

- **2025-04-29 — Qwen3 family launched; Qwen3-30B-A3B is a 2025-04-29 SKU (CORRECTION: not a 2026-02-10 release).** The "February 10, 2026 Qwen3-30B-A3B release" is contradicted: 30.5B total / 3.3B active MoE, 128 experts (8 active), OpenRouter catalog April 28, 2025. The brief attached the wrong model name to the right week — the real mid-February 2026 Alibaba release was Qwen3.5-397B-A17B. No "Qwen3.5-30B-A3B" SKU was found.
- **2026-02-16/17 — Qwen3.5-397B-A17B released (Apache 2.0).** Unveiled in Beijing on Chinese New Year Eve (Feb 16); OpenRouter dates Feb 16, Artificial Analysis Feb 17. 397B total / 17B active, 512 experts top-K (10 routed + 1 shared per token), hybrid 3:1 Gated DeltaNet → Gated Attention stack (60 layers, 15 attention layers with 2 KV heads each; ~7.7 GB KV cache at 262K — exceptionally small), 262K native context (YaRN → 1M), natively multimodal (text + 1344×1344 images + 2h video; no separate VL variant), 201 languages, thinking mode default, tool calling. Hosted counterpart Qwen3.5-Plus: 1M context with built-in tools, ~$0.40/$2.40 per M tokens internationally.
- **2026-04-02 — Qwen3.6-Plus (proprietary API, NOT open weights).** Agentic coding focus (repository-level engineering, autonomous plan/test/iterate loops), native multimodal perception (screenshots/design drafts → code), 1M context, ~¥2/M input and ¥12/M output on Bailian; reported to rival Claude Opus 4.5 on SWE-bench.
- **2026-04-16 — Qwen3.6-35B-A3B (open weights).** 35B total / 3B active MoE, positioned as the efficient agentic-coding model "ordinary developers could self-host"; independent testers reported better results than Claude Opus 4.7 on specific coding tasks. Only 10 full-attention layers hold length-growing KV cache: 10 KiB/token at FP8 (3.2× smaller than the 27B dense sibling; the 30 DeltaNet layers hold a constant recurrent state, length-independent).
- **2026-04-21 — Qwen3.6-Max-Preview.** Preview of the flagship tier; developers could integrate while the flagship tier became paid and lower tiers stayed open-source.
- **2026-05-20 — Qwen3.7-Max unveiled (proprietary, NOT open weights).** Announced at the Alibaba Cloud Summit / Apsara Conference in Hangzhou; API publicly available 2026-05-21–23; 1M context, up to 65,536 output tokens; Artificial Analysis Index 56.6 (global rank 5, #1 in China at launch). Qwen3.7-Plus also proprietary (text/image/video input, ¥0.4/1.6 per 1M tokens).
- **2026-07-19 — Qwen3.8-Max announced ("second only to Fable 5" [VENDOR]).** Announcement precedes weights by ~3.5 weeks.
- **2026-08-12 — Qwen3.8-Max weights released on Hugging Face and ModelScope.** First downloadable Qwen-Max-tier flagship (all previous Max models were API-only): 2.4T total / 95B active per token, 512 experts (10 routed + 1 shared), 92 layers, MTP training, 262K native → 1M extended context. License: **custom restrictive "qwen3.8-max" license** — above 100M MAU or $20M monthly revenue must display the model name; MaaS/AI-assistant businesses above $50M trailing revenue need a separate commercial license; internal use exempt. Critical caveat: downloaded weights are **text-only** — the model card states "Qwen3.8-2.4T-A95B is a text-only model… Multimodal inputs are not supported"; vision belongs to the hosted Max API only. Flexible thinking control via `reasoning_effort` + `preserve_thinking`. [VENDOR] Alibaba self-reported: Terminal-Bench 2.1 86.6, PaperBench 93.0, GPQA Diamond 92.6, HLE 43.6 (Fable 5: 53.3), OSWorld-Verified 86.1 (Fable 5: 85.0), Agents' Last Exam 52.4; independent BenchLM composite #6 of 225 at 79/100. Serving reality: cheapest usable quant ~450GB combined RAM+VRAM; lossless BF16 ~4.9TB; Alibaba reference deployment 72 Blackwell Ultra GPUs.
- **2026-08-14 — Qwen3.8-27B weights (Apache 2.0).** 27.8B dense multimodal (native vision encoder, Multi-Token Prediction heads, 262,144 native context → 1M), ~14–16GB VRAM at 4-bit (RTX 4090 class), AA Index 52 — the independent size-class leader (Qwen3.6-27B: 38; Muse Glimmer 30B: 35); [VENDOR] vendor-reported Terminal-Bench 2.1 73.0.
- **2026-09-02 — Qwen3.8-Max-0902.** Newer checkpoint of the Max line via Alibaba Cloud, 1M context; placed 4th in the September 5 Code Arena WebDev table (preliminary); license unchanged (custom restrictive).
- **2026-09-18 — Qwen3.8-Omni-Flash.** Omni, 1M context — API-only, not open-weight (chronology note to prevent miscategorization).

### DeepSeek — V4 family and the V4.1 architecture break

