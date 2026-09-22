---
id: frontier-models-2026/04-deepseek-v4-deepseek-v4-1-flash/overview
title: "2. DeepSeek V4 & DeepSeek V4.1 Flash"
domain: deepseek-v4-deepseek-v4-1-flash
role: deep-dive
task: actor-profile
actors: ["Anthropic", "DeepSeek", "Google", "Hugging Face", "OpenAI", "OpenRouter"]
dates: ["2026-02", "2026-09"]
keywords: ["deepseek", "agent", "agentic", "agents", "astra", "attention", "benchmark", "benchmarks", "claude", "cost", "decode", "distillation"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [197, 275]
section: "2. DeepSeek V4 & DeepSeek V4.1 Flash"
sha256: af82923cb36274a8d340894ad324c9b4add74bca8807554284f16c896bb8faea
---

# 2. DeepSeek V4 & DeepSeek V4.1 Flash

## 2.1 DeepSeek V4 — release
- **Planned/announced for mid-February 2026, ~17 February (Lunar New Year)** — reported by The Information citing people with direct knowledge; DeepSeek declined to comment. Strategy mirrored the R1 launch (Jan 20, 2025, one week before Lunar New Year; R1 had triggered a ~$1T tech selloff).
- **DeepSeek's first flagship built specifically for coding** — successor to V3.2 (Dec 2025).
- Two serving tiers: **V4-Pro** (flagship) and **V4-Flash** (efficient).

## 2.2 DeepSeek V4 — architecture
- **Hybrid MoE architecture** optimized for multi-file refactoring, full-stack scaffolding, and whole-repo ingestion in a single prompt.
- **Engram Conditional Memory** (paper: DeepSeek + Peking University, 12–13 Jan 2026, arXiv:2601.07372) — conditional memory as a complementary sparsity axis to MoE; O(1) lookups for static patterns, freeing GPU for reasoning.
- **mHC (Manifold-Constrained Hyperconnections)** (1 Jan 2026) — training-stability fix enabling larger models on the same hardware.
- **Muon optimizer** for training stability.
- **Hybrid attention:** Compressed Sparse Attention + Heavily Compressed Attention.
- **Pre-training: 32T+ tokens.**
- **V4-Pro:** 1.6T total / 49B active; **V4-Flash:** 284B total / 13B active.
- **Context:** 1M tokens (up from 128K on V3.1).
- **API features:** OpenAI ChatCompletions + Anthropic Messages formats, JSON mode, tool/function calling, up to 128 parallel function calls, chat-prefix completion (beta), FIM completion (beta, non-thinking), dual Thinking/Non-Thinking modes, automatic KV-cache reuse.
- **Integrations:** Claude Code, OpenCode, OpenClaw, CodeBuddy (pre-tuned adapters on Pro).

## 2.3 DeepSeek V4 — benchmarks (vendor-reported)
| Benchmark | V4-Pro | V4-Flash | Reference points |
|---|---|---|---|
| SWE-bench Verified | **80.6%** | ~77% | Claude Opus 4.6: 80.8% — near tie |
| SWE-bench Pro | ~55–58% | ~53% | Claude Opus 4.7: 64.3% (gap real) |
| Terminal-Bench 2.0 | **67.9%** | ~58% | Opus 4.7: 69.4%; GPT-5.5: 82.7% |
| LiveCodeBench | **93.5%** | ~89% | best-in-class reported |
| Codeforces | **3206** | — | GPT-5.4: 3168 |
| GPQA Diamond | ~90–92% | — | Opus 4.7: 94.2% |
| MMLU-Pro | ~90% | — | Opus 4.5: 89.5% |
| AIME 2026 | ~95% | — | Opus 4.5: 95.1% |
| MCPAtlas (agentic) | **73.6** (Pro-Max) | — | tied Opus 4.6; leads open models |
| Toolathlon | 51.8 (Pro-Max) | — | — |
| SkillsBench Avg5 | ~50% | ~45% | — |

Note: Anthropic publishes SWE-bench **Pro** (81.2 for Fable 5.1), OpenAI quotes **DeepSWE** (74.1 for GPT-6 Astra) — three vendors, three benchmarks, no like-for-like comparison. DeepSeek itself acknowledges V4 "falls marginally short of GPT-5.4 and Gemini 3.1 Pro" on world knowledge by ~3–6 months.

## 2.4 DeepSeek V4 — license, price, availability
- **Open weights, MIT license** (both tiers) — Hugging Face / ModelScope.
- **API pricing (per 1M):** V4-Pro $1.74 in / $3.48 out ($0.145 cached in); V4-Flash $0.14 in / $0.28 out ($0.028 cached in). V4-Pro output ~87% cheaper than GPT-5.4 ($15) and ~91% cheaper than Claude Opus 4.6 ($25).
- Typical 10-step agent workflow: <$0.05 with optimized Flash/Pro routing vs $0.50+ with Opus 4.7/GPT-5.5.

---

## 2.5 DeepSeek V4.1 Flash — release (10 September 2026)
- **Released 10 September 2026** — not a leak, not a beta. A two-day API test (model ID `deepseek-v4.1-flash-expires-on-0910`, from Sept 8) ended on schedule; release landed across web app, mobile app, first-party API, with **weights + technical report on Hugging Face** (`deepseek-ai/DeepSeek-V4.1-Flash`).
- API model ID: **`deepseek-flash`** (single name). Legacy IDs `deepseek-v4-flash` / `deepseek-v4-flash-vision-exp` now route to V4.1 Flash at Flash pricing (retired).
- **DeepSeek is retiring the ~4× more expensive V4 Pro on 14 September 2026**, stating "V4.1 Flash has comprehensively surpassed V4 Pro across all key metrics."
- Independent analysis: "not a point release — a generation that declined to name itself."

## 2.6 DeepSeek V4.1 Flash — architecture (detailed)
- **Causal Encoder-Decoder (CED):** 40-layer Transformer split into a **20-layer causal encoder + 20-layer decoder**. The decoder's global KV cache is projected from the encoder's final hidden states instead of rebuilt layer-by-layer.
- **Parameters:** 552B backbone — **8B active per token during prefill, 16B during decode** (the "smallest model in the new architecture family").
- **+196B Engram conditional-memory parameters** (sparse, token-based lookup) → ~763B stored parameters in the HF repo.
- **MoE layout:** each MoE layer has 1 shared expert + 384 routed experts, 6 routed experts activated per token.
- **Attention/cache efficiency:** Compressed Sparse Attention 2 (CSA2) + **FP4 KV caching** (E2M1) + shared attention indices → **~890 bytes/token global KV cache** (~¼ of V4-Flash, ~1/437th of DeepSeek V1).
- **SWA Bounded Replay:** cuts persistent KV-cache footprint to ~⅛ of V4-Flash.
- **Training:** from scratch on **45T mixed text+image tokens**; context extended to 1M at the 34T mark.
- **Native multimodal visual understanding** (text + image in, text out).
- **Context:** 1M tokens; **max output 384K tokens**; concurrency limit 2,500 (vs 500 on V4-Pro — 5× throughput headroom).
- **Reasoning effort dial 1–100** (continuously controllable; thinking on by default with low/high/max presets) — trade accuracy vs cost per request without switching models.
- **Post-training:** SFT + RL + on-policy distillation; agentic gains credited to "large-scale automated synthesis of agent tasks and environments."

## 2.7 DeepSeek V4.1 Flash — benchmarks
| Benchmark | V4.1 Flash | Reference |
|---|---|---|
| Terminal-Bench 2.1 | **90.6** | ahead of Claude Opus 5 (vendor claim) |
| DeepSWE v1.1 | **74.2** (up from 54.4 on prior Flash) | — |
| GPQA Diamond | **90.9** (vs 92.4 prior flagship) | "98% of flagship score, 45% of price" |
| (Flash-0731 preview row, vendor) | TB 2.1: 82.7; Toolathlon-V: 70.3; Agents' Last Exam: 25.2; AutomationBench: 25.1 | Opus 4.8 beats Flash on all nine rows (DeepSeek published both) |

## 2.8 DeepSeek V4.1 Flash — license, price, availability
- **MIT license, weights + repo on Hugging Face and ModelScope** — permissive end of the open-weights spectrum (no revenue threshold, no field-of-use clause); self-host, fine-tune, serve commercially, royalty-free. Inference folder + separate encoding module in repo; community building inference support from day one.
- **Official API pricing (per 1M, off-peak):** $0.15 input / $0.60 output / $0.003 cached input. **Peak** (Mon–Fri 01:00–04:00 & 06:00–10:00 UTC): doubled — $0.30 / $1.20 / $0.006. Weekends always off-peak.
- **Speed (observed):** ~190–427 tok/s community-reported; 215 tok/s p50 in one routing analysis; p50 TTFT 842 ms, p95 2.40 s, ~1.2% error rate.
- **Distribution:** DeepSeek app (rolling out), official API, **OpenRouter** (`deepseek/deepseek-v4.1-flash` — 1,048,576-token context, reasoning default `high`, tools, JSON, streaming; upstreams include DeepInfra, Novita, Venice; 99.92% 3-day availability); partners WorkBuddy, CodeBuddy, OpenCode Go, ClinePass.
- **Positioning:** killed its own premium tier — "a lab deciding the premium tier was no longer worth charging for." Best value in its class on price/perf/agent metrics per multiple independent analyses.

---

