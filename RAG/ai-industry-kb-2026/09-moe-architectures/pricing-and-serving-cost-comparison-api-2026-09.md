---
id: ai-industry-kb-2026/09-moe-architectures/pricing-and-serving-cost-comparison-api-2026-09
title: "Pricing and serving-cost comparison (API, 2026-09)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "MiniMax", "Moonshot", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-04", "2026-05-22", "2026-07", "2026-08-16", "2026-08-19", "2026-08-22", "2026-09-07", "2026-09-09", "2026-09-22"]
keywords: ["cost", "pricing", "agentic", "apache", "attention", "awq", "benchmarks", "claude", "compute", "deepseek", "fp8", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5080, 5120]
section: "9. MoE Architectures"
sha256: a4551bf5cebcf7db8477ac0de02451b25c9c1081e62a0a16cf985409362d80b5
---

# Pricing and serving-cost comparison (API, 2026-09)

- **SWE-bench Verified:** MiniMax M3 80.5% (Artificial Analysis, 2026-09-07); Qwen3.6-35B-A3B 73.4% (2026-04); MiMo-V2-Flash 73.4% (2025-12, #1 open at release); Gemma 4 26B-A4B 17.4% (community-measured); Qwen3.8-Max 67.7% SWE-bench Pro / 73.5% FrontierSWE (vendor, 2026-08).
- **LiveCodeBench:** DeepSeek-V4-Pro-Max 93.5% (2026-04, ahead of Claude Opus 4.6 and Gemini 3.1 Pro per vendor); V4-Flash Think Max 91.6%; Qwen3.6 v6 80.4% (2026-04); Codeforces 3,052 (V4-Flash).
- **Agentic coding:** DeepSeek-V4-Pro 67% pass rate (vs Sonnet 4.5's 47%, trailing Opus 4.6 Thinking's 80%); MiniMax M3 "matches Claude Sonnet 4.6" on real-world agentic benchmarks (Morph, vendor-adjacent — UNVERIFIED independently).
- **Math/reasoning:** AIME 2025 94.1% (MiMo-V2-Flash, 2025-12); GPQA Diamond 93.5% (Kimi K3, third-party table), 92.6% (Qwen3.8-Max, vendor), 84.1% (Qwen3.6).
- **Multimodal:** MMMU 81.7%, MathVista 86.4% (Qwen3.6); WebDev Arena Elo 1679 (Kimi K3, #1 open); AA Intelligence Index v4.1 composite 57.1 (Kimi K3, #4 globally); PaperBench 93.0 (Qwen3.8-Max).
- **Tool use:** MCPMark 37.0 (Qwen3.6) vs 18.1 (Gemma 4-31B) — the MoE active-compute advantage shows most in agentic/tool benchmarks.
- **Caveat (applies to all vendor numbers):** vendor-reported benchmarks use vendor-chosen harnesses; treat cross-vendor comparisons as directional. The most trustworthy cross-model numbers are third-party (Artificial Analysis, community reproductions) and are marked as such above.

### Pricing and serving-cost comparison (API, 2026-09)

Per 1M tokens, input / cached input / output (where published):

- DeepSeek-V4-Pro: $0.435 / $0.003625 / $0.87 (75% discount permanent since 2026-05-22; peak/off-peak tiers from 2026-08-16: $1.32/$0.66 input).
- DeepSeek-V4-Flash: peak/off-peak $0.44/$0.22 input.
- GLM-5.3-Flash: $0.15 / $0.03 / $0.50 (launch promo through 2026-09-09: $0.075 / $0.015 / $0.25).
- GLM-5.3 (flagship): $1.40 / $4.40 output.
- MiniMax M3: $0.30 / — / — (official; among cheapest above 80% SWE-bench Verified).
- MiMo-V2-Flash: $0.10 / $0.30 output (2025 pricing; V2.6 pricing not yet published as of 2026-09-22).
- Qwen3.8-Max: $2.00 / $0.25 cache reads / $6.00.
- Kimi K3: $3.00 / $0.30 / $15.00 — the most expensive Chinese-lab model to date, priced at Claude Sonnet tier.
- Reference efficiency points: Nemotron 3.5 Lightning (30B/3B MoE+Mamba-2) $0.22/M output; Gemma 4 31B dense $0.40/M output.
- The price ladder mirrors active compute almost monotonically — Kimi K3 (104B active) and Qwen3.8-Max (95B active) sit at the top; 3B-active models (Qwen3.6, Nemotron Lightning) at the bottom — while capability per dollar is dominated by post-training, not architecture. Chinese-lab open MoE APIs undercut Western equivalents by 5–50× at matched capability tiers (the exception proving the rule is Kimi K3 at Claude Sonnet tier for the largest open model ever released). The cleanest 2026 illustration of the MoE efficiency trade is the GLM sibling pair: flagship 753B-class at AA index 60 vs Flash 320B/18B at 57 — 3 points of measured intelligence for ~9× the API cost.

### License landscape (2026)

- **Apache 2.0:** Qwen3.6-35B-A3B, Gemma 4 family (notable shift from earlier Gemma-specific terms), gpt-oss-120b/20b.
- **MIT:** Kimi K3, GLM-5.3-Flash, MiMo-V2-Flash, MiMo-V2.6-Flash/Pro, DeepSeek-V4 (per later sources; early April 2026 coverage cited Apache 2.0 — UNRESOLVED, do not cite without checking).
- **Custom / community:** MiniMax M3 (MiniMax Community License), GLM-5.3 flagship (glm-5.3 license), Qwen3.8-Max (non-Apache open license).
- **Trend:** the most permissive standard licenses (Apache 2.0, MIT) now cover the majority of open-weight MoE releases; custom licenses cluster at the largest/capability-leading models (M3, GLM-5.3, Qwen3.8-Max).

### Quantization formats used across the 2026 MoE fleet

- **FP8 (e4m3):** default serving format for 300B+ models — GLM-5.3-Flash ships FP8 by default (~306–331 GiB); MiMo-V2.6-Flash weights 172.9 GB in FP8; MiniMax M3 MXFP8 variant ~440 GB; Qwen3.8-Max's FP8 repo is the most-downloaded variant (21,400 downloads vs 17,386 for bf16 as of 2026-08-22 — FP8 leading, indicating serving-motivated fetchers).
- **NVFP4 (group-16):** MiniMax M3 (~245 GB safetensors → ~61 GB/GPU at TP4); vLLM v0.15.0 added NVFP4-CUTLASS kernels; Unsloth Studio exports to NVFP4 since July 2026 (v0.1.481-beta); Axolotl added NVFP4 (4-bit) MoE LoRA in July 2026 via ScatterMoE (W4A16) and SonicMoE (W4A4), including lossless adapter merge back into a plain NVFP4 checkpoint. Third-party nuance (unsloth-cli docs, Sept 2026): NVFP4 exports were live-tested for *serving* via vLLM on Thor (JetPack R38.2.2), not for training — Unsloth/NVFP4 claims to date are export-and-serve claims, not NVFP4-training claims.
- **MXFP4/MXFP8:** native in Kimi K3 (MXFP4 weights / MXFP8 activations; per-expert MXFP4 with E8M0 scale per 32 weights, 33.0M params per expert, 2.72T routed params total) and gpt-oss (MXFP4 native, 60.8 GiB / 12.8 GiB checkpoints).
- **AWQ / Marlin 4-bit:** local-serving standard; documented Mixtral 8x7B vLLM case: 90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT.
- **GGUF (Q4_K_M and below):** CPU/offload tier — Qwen3.6 20.9 GB Q4 on MacBook Pro; Gemma 4 26B-A4B ~11 GB routed experts on NVMe with ~24 MB/forward hot-loading. Unsloth's Dynamic v3.0 remains the current UD quant generation (launched 2026-08-19); GLM-5.2's Dynamic GGUF was presented by a secondary source as 84% size reduction (~239 GB at 2-bit, ~217 GB at 1-bit) with ~82%/~76% accuracy retention and ~98% at 4-bit, runnable on 256 GB unified-memory machines — [VENDOR-ORIGINATED, UNVERIFIED].
- **Invariance principle (verified):** quantization scales total and active params equally — the total/active ratio (e.g., 5.40× at fp32 and int4 in the OLMoE profile) does not change.

### Attention innovations glossary (what makes 1M context affordable)

