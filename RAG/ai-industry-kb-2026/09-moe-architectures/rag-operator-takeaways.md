---
id: ai-industry-kb-2026/09-moe-architectures/rag-operator-takeaways
title: "RAG-operator takeaways"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "CISA", "DeepSeek", "MiniMax", "Moonshot", "Nvidia", "Z.ai", "vLLM"]
dates: ["2025-08-05", "2025-12-16", "2026-04-16", "2026-04-24", "2026-05", "2026-05-22", "2026-06-01", "2026-07-24", "2026-07-27", "2026-08-03", "2026-08-16", "2026-08-22", "2026-08-25", "2026-08-28", "2026-09-09"]
keywords: ["agent", "agentic", "agents", "apache", "attention", "awq", "blackwell", "compute", "cost", "deepseek", "fp4", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [5175, 5229]
section: "9. MoE Architectures"
sha256: 2d16174e3c03b1f49c0b9b8b4df3cf16e919c0ac5c406732fcb8fc652c962843
---

# RAG-operator takeaways

- **Mixtral 8x7B (2023):** 46.7B/12.9B, Apache 2.0 — the original open MoE economics proof: dense-70B-class capability at ~6× inference speed; the AWQ 4-bit vLLM case (90 GB → 26 GB VRAM, 2.1× throughput, −75% TTFT) is the canonical local-serving datum.
- **Nemotron 3.5 Lightning (2026):** 30B/3B MoE + Mamba-2 hybrid with speculative decoding — 235–494 t/s at $0.22/M output (AA, Aug 2026) vs Gemma 4 31B dense at 37–222 t/s and $0.40/M; capability 24 vs 30. The MoE+SSM hybrid is the 2026 efficiency frontier at ~30B total.
- **Llama-4-Maverick-class ultra-sparse MoE (<1% activation density):** EP is 7–12% SLOWER (ROCm measurements) — all-to-all overhead exceeds sparse-compute savings. Sparsity has a floor below which parallelism loses; the 2026 serving rule is ≥3% density for EP to win.
- **MoUE (arXiv 2603.04971, Mar 2026):** reusing a layer-agnostic universal expert pool across layers ("virtual width" from depth) — up to +1.3% over matched MoEs, +4.2% when converting existing MoE checkpoints. Research-scale only; the open question is whether it holds at 2T+ params and 1M context.
- **FLAME (arXiv 2605.09355, May 2026):** adaptive MoE for continual multimodal learning with shared + task-specific experts — the 2026 direction for specialization without full retraining.
- **Phase-aware routing (2026):** routing at environment-step granularity for RL agents (temporal consistency across multi-step plans) — promising for agentic workloads, unproven in production serving stacks.
- **Confidence-gated dynamic expert count (early 2026 work):** easy tokens → fewer experts, suggesting 20–40% compute savings — the one untapped compute axis beyond the active/total ratio itself.

### RAG-operator takeaways

- **$/token economics favor MoE for retrieval-augmented pipelines:** GLM-5.3-Flash at $0.075/M input (promo, through 2026-09-09) to $0.15/M; DeepSeek-V4-Flash cache-hit pricing at fractions of a cent — long retrieved contexts are cheap per token.
- **1M context reduces chunking pressure:** all 2026 flagships accept 1M tokens natively, so "stuff the whole corpus section" beats fine-grained retrieval for many workloads — but KV residency caps concurrent long sessions, making prefix caching and session batching the operative optimizations.
- **Active params predict per-token latency, total params predict fit:** size the GPU fleet on total params (quantized), the latency SLO on active params.
- **Tool-use MoE models are the RAG agent substrate:** Qwen3.6-35B-A3B's 37.0 MCPMark (2× Gemma 4-31B) and DeepSeek-V4's 67% agentic pass rate show sparse models now lead on function calling — the RAG agent loop (retrieve → call tools → synthesize) runs best on exactly these models.
- **Prefix/cache-aware routing:** for RAG serving, the KV cache is the bottleneck, not the experts — invest in prefix caching, HybridKV-style compression, and attention-efficient models rather than denser retrieval.

### Serving-relevant spec snapshot (dated, economics columns)

Total / active per token / context / KV-attention design / license / release:

- gpt-oss-120b: 117B / 5.1B / 128K / standard attention / Apache 2.0 / 2025-08-05
- gpt-oss-20b: 21B / 3.6B / 128K / standard attention / Apache 2.0 / 2025-08-05
- MiMo-V2-Flash: 309B / 15B / 256K / 5:1 SWA(128)/global + sink bias / MIT / 2025-12-16
- Qwen3.6-35B-A3B: 35B / 3B / 262K native (1M w/ YaRN) / 3:1 Gated DeltaNet:gated attention / Apache 2.0 / 2026-04-16
- Gemma 4 26B-A4B: 25.2B / 3.8B / 256K / SWA(1024)+global hybrid, shared KV / Apache 2.0 / ~2026-04
- DeepSeek-V4-Pro: 1.6T / 49B / 1M (384K max output) / NSA lineage / MIT (disputed — see §27 list) / 2026-04-24
- DeepSeek-V4-Flash: 284B / 13B / 1M / NSA + DSA + token-wise compression / MIT (disputed) / 2026-04-24
- MiniMax M3: 428B / 23B / 1M / MSA block-sparse + GQA(64Q/4KV) / MiniMax Community License / 2026-06-01
- Kimi K3: 2.8T / 104B (disputed ~50B — see §27 list) / 1M / 69 KDA + 24 Gated MLA / open-weight / 2026-07-27
- Qwen3.8-Max: 2.4T / 95B / 1M / Gated DeltaNet + gated attention (64Q/4KV) / non-Apache open license / 2026-08-03
- GLM-5.3-Flash: 320B / 18B / 1M (config; 300K cited in one eval — see §27 list) / hybrid sparse+linear + mHC / MIT / 2026-08-25/26
- GLM-5.3 (flagship): 753B / undisclosed / 1M / MoE base from 5.2 / custom glm-5.3 license / 2026-08-28
- MiMo-V2.6-Flash: 309B / 15B / 1M / 5:1 SWA(128)/global / MIT / ~2026-09
- MiMo-V2.6-Pro: 1.02T / 42B / 1M / same family / MIT / ~2026-09
- Nemotron 3.5 Lightning: 30B / 3B / — / MoE + Mamba-2 hybrid / — / 2026 (AA efficiency reference)

### Pricing mechanics notes (dated)

- **2026-05-22:** DeepSeek makes V4-Pro's 75% discount permanent ($0.435/$0.003625/$0.87 input/cached/output).
- **2026-07-24:** legacy `deepseek-chat`/`deepseek-reasoner` aliases retired, routing to V4-Flash during the grace period — model-name continuity matters for cost tracking.
- **2026-08-16:** V4 peak/off-peak tiers introduced (Pro $1.32/$0.66 input; Flash $0.44/$0.22 input) — time-of-day becomes a pricing dimension for MoE serving.
- **2026-09-09:** GLM-5.3-Flash launch promo ends ($0.075/$0.015/$0.25 → $0.15/$0.03/$0.50 input/cached/output).
- **2026-08-22 adoption signal:** Qwen3.8-Max HF downloads 17,386 (bf16) / 1,141 likes; FP8 repo 21,400 downloads — the serving-motivated variant leads.
- **The price ladder mirrors active compute almost monotonically:** Kimi K3 (104B active) and Qwen3.8-Max (95B active) at $3.00/$2.00 per 1M input; 3B-active models (Qwen3.6, Nemotron Lightning) at the bottom. Capability per dollar is dominated by post-training, not architecture.

### RAG session-economics patterns

- **KV residency caps concurrent long-context sessions.** All experts occupy VRAM, so the KV pool left over is smaller than in a same-size dense model — for RAG this means session-level batching and prefix caching are the operative optimizations, not model shrinking.
- **Retrieved-context prefixes are the cost lever.** DeepSeek-V4-Pro cache-hit pricing ($0.003625/M) vs miss ($0.435/M) is a ~120× spread; GLM-5.3-Flash cached $0.03/M vs $0.15/M. Stable system-prompt + retrieval-template prefixes make MoE's cheap long context actually cheap.
- **1M context changes chunking economics:** "stuff the whole corpus section" beats fine-grained retrieval for many workloads — but only where the model's post-training actually uses long context (the attention-side innovations, not MoE sparsity, decide).
- **Tool-use-first model selection for RAG agents:** Qwen3.6-35B-A3B's 37.0 MCPMark (2× Gemma 4-31B) and DeepSeek-V4's 67% agentic pass rate put sparse models ahead on function calling — the retrieve → call tools → synthesize loop runs best on exactly these models.
- **Self-host vs API breakpoint (2026):** self-hosted ballpark $0.25–0.30/MTok on Blackwell (Qwen3-MoE 235B-A22B FP4 on 8× B200; DeepSeek V3.1 FP4+MTP on NVL72) vs API Flash-tier $0.075–$0.15/M input — at Flash-tier prices, API undercuts self-host for most RAG workloads unless residency or data-sovereignty demands otherwise; at flagship tier ($1.40–$3.00/M), self-host breaks even at scale.

### Unresolved discrepancies and unverified claims (do not cite without checking)

