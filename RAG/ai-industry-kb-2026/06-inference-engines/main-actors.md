---
id: ai-industry-kb-2026/06-inference-engines/main-actors
title: "Main actors"
domain: inference-engines
role: deep-dive
task: architecture
actors: ["Alibaba", "DeepSeek", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Stripe", "Z.ai", "vLLM"]
dates: ["2026-09"]
keywords: ["accelerator", "agentic", "agents", "apache", "benchmark", "benchmarks", "blackwell", "capex", "compute", "cost", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2429, 2458]
section: "6. Inference Engines"
sha256: bdd870e93adb3ef7a8855866bd44ed44e0fa0af09af4efcbbe5b639ac5600cbc
---

# Main actors

- **Gartner (Aug 2026):** AI-optimized IaaS **$42B in 2026 (+96% YoY)**; inference **$23.3B** vs training **$19B**; 55% inference share in 2026 → **59% in 2027**.
- **Gartner (Oct 2025 vintage):** $37.5B AI-optimized IaaS forecast for 2026, 55% inference — upward revision within ten months.
- **Production-team economics (Aug 2026 analysis):** inference **60–80% of AI GPU spend** for production teams; roughly **two-thirds of all AI accelerator spending** in 2026.
- **Hyperscaler capex:** **$757.7B in 2026 (+60% YoY)** — third straight year of 60%+ growth.
- **GPT-4-level inference unit cost:** ~$20/M tokens (late 2022) → ~$0.40/M tokens (early 2026) [COMMUNITY].
- **Stripe (reported 2026):** **73% cost reduction**, 50M daily API calls, one-third the GPU fleet via vLLM.
- **Self-hosted vLLM cost:** **$0.50–1.00 per million tokens** [COMMUNITY guides, directional].
- **Decision math:** with GPT-4-level inference at ~$0.40/M tokens and agentic workloads multiplying tokens 5–50x per interaction, a 29% throughput delta (SGLang on prefix-heavy load, sibling 06b) or a fleet reduction (Stripe's 73%) translates into seven-figure annual differences for high-volume deployments — which is why 2026 procurement treats the inference engine as infrastructure, not tooling [COMMUNITY synthesis].
- **Disaggregation dividend:** **4–10x cost-per-token reduction** vs aggregated serving [COMMUNITY practitioner analysis]; Dynamo-claimed 7x throughput-per-GPU on DeepSeek R1/GB200 [VENDOR].
- **Red Hat AI Inference Server:** **2–4x token production** with pre-optimized models [VENDOR].
- **llm-d v0.5 benchmarks:** up to **57x faster TTFT, 2x throughput** vs round-robin under high prefix reuse (8 pods/16×H100) [VENDOR]; operators' own measurements: ~25% over defaults, 2–3x tokens/s/GPU with prefix-cache-hit routing, 3–5x cost-per-token on chat-shaped workloads [COMMUNITY].
- **vLLM release cadence:** roughly one minor version per month in 2026 (v0.15.1 → v0.29.0 in eight months). v0.20.0: 752 commits/320 contributors; v0.25.0: 558 commits/232 contributors; v0.26.0: 411 commits/212 contributors; v0.28.0: 584 commits/270 contributors (76 new).
- **GitHub scale:** vLLM ~50K–75K stars (sources disagree; present as range), ~2,400 contributors, 100+ model architectures, Apache 2.0; issue response 12 hours–3 days [COMMUNITY].
- **Kimi-K3 (v0.28.0 release notes) [VENDOR]:** combined all-gathers 1.5–3x kernel-level speedup; adaptive speculative token budget ~60% better DSpark TTFT; shared-expert sharding ~17 GiB/GPU memory savings.
- **vLLM-Omni paper (arXiv:2602.02204):** up to **91.4% lower JCT** vs baselines [UNVERIFIED as independently reproduced].
- **vLLM-Omni diffusion wall-clock (H200, Wan2.2-I2V) [COMMUNITY]:** 133.94 s (v0.16.0 retro) → 93.67 s (v0.18.0) → 79.19 s (v0.20.0).
- **New defaults (v0.28.0):** `max_num_batched_tokens` 8192 → 16384; Blackwell CUDA-graph capture default raised to 1024.
- **Independent TRT-LLM counter-benchmark (Aug 2026, Qwen3-Coder-30B-A3B, workstation Blackwell) [COMMUNITY]:** vLLM led at concurrency 1 (12.3 vs 10.8), 8 (56.3 vs 44.4), 64 (129.0 vs 115.1); TRT-LLM led only at 32 (120.5 vs 119.1, +1%). Peak: vLLM 129.0 vs TRT-LLM 120.5. Published 10–25% TRT-LLM claims did not hold on workstation Blackwell (likely because sm_120 lacks FlashAttention-4); vLLM kept climbing to 64 concurrent while TRT-LLM saturated at 32.
- **Head-to-head counter-evidence on the vLLM side (AIMultiple/GPT-OSS-120B tests) [COMMUNITY]:** vLLM wins **fastest TTFT** across concurrency levels in GPT-OSS-120B tests and leads at **100+ concurrent requests** — the workload-dependent correction to any "SGLang always faster" reading (SGLang's +29% holds on prefix-heavy load, tested with identical FlashInfer kernels, i.e. the difference is orchestration overhead, not kernel performance).
- **AI platforms market:** $181.3B (2026) → $496.9B (2030), 28.7% CAGR [COMMUNITY].
- **The 1M-context serving proof:** GLM-5.3-Flash NVFP4 at 1M-token context on DGX Spark is the longest-context vLLM deployment documented this cycle [COMMUNITY]; the SM121 kernel patch it required shows community patching still carries part of the day-0 burden.
- **Enterprise adoption (PyTorch Conference 2026/Futurum) [COMMUNITY]:** 51% of enterprises pursue balanced in-house/vendor AI; 63.9% deploy on provider-managed infrastructure.
- **FlagOS chip coverage:** vLLM's multi-hardware effort tracked against a **20+ chip test base** — the concrete mechanism behind the "any accelerator" claim [COMMUNITY].
- **Agentic token multiplication:** one user request becomes dozens of model calls; agents generate 5x–50x more tokens per interaction than single-turn chat [COMMUNITY] — the workload shape that made prefix-cache efficiency and structured tool-call generation the 2026 differentiators rather than raw prefill FLOPS.
- **Context-length production costs:** 128K–1M token contexts moved from flagship demos to production RAG and code agents; every long-context request is a prefill bill paid at inference time — a structural driver of the spend flip alongside test-time compute.
- **The MLA compounding pattern (2026):** MLA compresses the KV cache, FP8/FP4 quantizes it, disaggregation moves it, and RadixAttention shares it — four independent multipliers on cost-per-token [COMMUNITY synthesis].
- **New entrants (2026, adjacent):** TileRT (per-user decode-speed runtime pairing with vLLM prefill); Tiny-vLLM (community C++/CUDA lightweight engine, Hacker News traction); Neutree 1.2 / Flex Engine (Arcfra, September 2026 — enterprise platform unifying vLLM and SGLang under one gateway, proprietary Flex Engine for non-LLM models, automatic KV-cache/GPU-memory calculation, project-based API keys).

## Main actors

