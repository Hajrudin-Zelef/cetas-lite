---
id: open-local-models-2026/01-step-1-open-local-ai-models-chinese-track-research-fiche/part-b-deepseek
title: "PART B — DEEPSEEK"
domain: step-1-open-local-ai-models-chinese-track-research-fiche
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "DeepSeek", "Huawei", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Z.ai", "vLLM"]
dates: ["2025-05", "2025-07-11", "2026-04-20", "2026-04-24", "2026-05-20", "2026-06-12", "2026-06-13", "2026-06-16", "2026-07-16", "2026-07-24", "2026-08-14", "2026-08-26", "2026-08-28", "2026-09-10", "2026-09-18"]
keywords: ["deepseek", "agent", "agentic", "agents", "ascend", "attention", "attribution", "benchmarks", "claude", "compute", "cost", "cyber"]
source: docs/RAG/Modèles IA open  locauxEN.md
source_anchor: ""
source_lines: [64, 136]
section: "STEP 1 — Open / Local AI Models (Chinese track) — Research Fiche"
sha256: 5cb52b3bfa6afe38a65e928d63e26e1b8223fb1a244fb317f1e29080638a3ddf
---

# PART B — DEEPSEEK

## PART B — DEEPSEEK

### B0. R-series (reasoning lineage)
- **DeepSeek-R1** — Jan 20, 2025 (updated R1-0528, May 2025). MoE 671B/37B active. 64K → 164K context. AIME 2024 79.8% at ~1/18 the per-token price of o3. **MIT**; HF `deepseek-ai/DeepSeek-R1`. First-party API retired July 24, 2026 (weights remain).
- **DeepSeek-R2 — NOT RELEASED (as of Sept 2026).** All specs online (685B–1.2T, $0.14/M) are unconfirmed leaks. Documented delays: CEO Liang Wenfeng dissatisfied with performance; export controls; failed training run on Huawei Ascend 910C (stability/interconnect/software); pivot back to Nvidia for training, Ascend for inference. No launch timeline.

### B1. DeepSeek V4 family — V4 Preview April 24, 2026
- **DeepSeek V4-Pro:** MoE **1.6T total / 49B active**. **Hybrid CSA + HCA attention + mHC hyper-connections** (replaces V3's MLA); **Muon optimizer**; **mixed FP4+FP8 precision**; 32T+ training tokens. 1M context, 384K max output. At 1M tokens: 27% of V3.2's FLOPs, 10% of KV cache. Text-only. **MIT**; HF `deepseek-ai/DeepSeek-V4-Pro-0813`. LiveCodeBench 93.5, Codeforces 3206, GPQA Diamond 90.1. AA Index 36. Off-peak pricing **$0.66/M in / $1.98/M out**; cache hit $0.022. Self-host: 8× H200/B200 (SGLang NVFP4) or 4× GB300 (vLLM).
- **DeepSeek V4-Flash:** MoE **284B total / 13B active**; same CSA/HCA architecture; 1M context. **MIT**; HF `deepseek-ai/DeepSeek-V4-Flash-0731`. Current off-peak **$0.22/$0.66**. AA Index 35.
- **Engram note:** "Engram conditional memory" was rumored pre-V4 but absent from V4; it shipped in V4.1 Flash.

### B2. DeepSeek V4.1 Flash ⭐ — released September 10, 2026
- **Architecture:** First model on **Causal Encoder-Decoder (CED)** — 40 layers (20 encoder + 20 decoder). **552B backbone, 8B active on input / 16B on output**; plus **196B Engram conditional-memory params** (~763B total). CSA-2 + FP4 KV cache → 890 bytes/token (~¼ of V4 Flash). Trained on **45T tokens**. Native vision. 1M context, 384K max output. Observed 214 tok/s, TTFT ~1.1s.
- **License:** **MIT**. Weights: `deepseek-ai/DeepSeek-V4.1-Flash` on HF (FP8 + technical report). Checkpoint ≈ 510 GB.
- **Pricing (two tiers):** Off-peak — cache-hit input **$0.003/M**, cache-miss **$0.15/M**, output **$0.60/M**. Peak (Mon–Fri 01:00–04:00 and 06:00–10:00 UTC) = double. Weekends always off-peak.
- **Product-line event:** V4-Pro API traffic was to auto-route to V4.1 Flash on Sept 14; **reversed** — V4-Pro continues, no date for V4.1-Pro.
- **Benchmarks (vendor):** Terminal-Bench 2.1 **90.6** (vs Opus 5.0 89.1), DeepSWE v1.1 **74.2** (vs Opus 5.0 74.0, V4-Pro 62.7), CyberGym 88.1, AutomationBench 54.8, GPQA Diamond 90.9. Beats V4-Pro on 12/14 rows at ~⅓ backbone, ~⅙ active params.
- **Independent (Artificial Analysis):** Index **≈39–40**; AutomationBench-AA 69%, cost/task **$0.27** (Pareto frontier). Flagged as "very verbose" (250M output tokens vs 140M median).
- **Distinguishing:** Cheapest intelligent open-weight model per task; first shipped Engram memory; 890 B/token KV cache makes 1M context economically viable.

---

## PART C — KIMI (Moonshot AI, Beijing, est. 2023)

**License trajectory:** K2 (MIT) → K2.5/K2.6/K2.7 Code (modified MIT, attribution trigger ~100M MAU/$20M-mo) → K3 (fully custom license, $20M/12-mo MaaS separate-agreement trigger).

### Kimi K2 — July 11, 2025
MoE **1T total / 32B active**; MLA; 128K → 256K context. **MIT**. Agentic-task optimized. OpenRouter ~$0.57/$2.30. First 1T open-weight agentic model.

### Kimi K2.5 — January 26–27, 2026
Same 1T/32B/384-expert MoE (8+1 shared active, 61 layers); vision-native (MoonViT, ~15T visual+text tokens); native INT4. 256K context. Modified MIT. Benchmarks: HLE **50.2** (vs GPT-5.2 45.5), BrowseComp 74.9, SWE-bench Verified 76.8 (scaffold; independent mini-SWE-agent 70.80%), AIME 2025 96.1. Pricing $0.60/$3.00, cached $0.10/M. **Agent Swarm** primitive (up to 100 parallel sub-agents). Deprecated May 20, 2026.

### Kimi K2.6 — April 20, 2026 (GA Apr 21)
1T/32B/384 experts, 61 layers, 160K vocab, MLA+SwiGLU, MoonViT 400M vision; native INT4; multimodal. 256K context. Modified MIT (`moonshotai/Kimi-K2.6`). Benchmarks: SWE-bench Pro **58.6** (tied GPT-5.5), Verified **80.2**, Terminal-Bench 2.0 66.7, AIME 2026 96.4, HLE-Full 54.0, AA Index **54 — #1 open-weight at launch**. Pricing $0.95/$4.00, cache-hit $0.16/M. Agent Swarm scaled to **300 sub-agents, 12–13h autonomous sessions**.

### Kimi K2.7 Code — June 12, 2026 (HighSpeed June 15)
Coding/agents specialization of K2.6; thinking mode mandatory (temp hard-pinned 1.0). Modified MIT. Benchmarks (vendor-only): Kimi Code Bench v2 62.0 (+21.8% vs K2.6), MCP Mark Verified 81.1 (vs Opus 4.8 76.4). **~30% fewer thinking tokens** than K2.6; HighSpeed mode up to 6× throughput. Pricing $0.95/$4.00. Self-host footprint ~630–640GB at INT4.

### Kimi K3 ⭐ — API July 16, 2026; open weights July 26–27, 2026
- **Architecture:** **2.8T total**, sparse MoE — **896 experts, 16 active/token**; 61 layers. **Kimi Delta Attention (KDA)** (up to 6.3× faster decoding at 1M ctx); **Attention Residuals** (~25% training efficiency at <2% extra params). Native multimodal. Weights shipped as **MXFP4 weights / MXFP8 activations** — download footprint **~1.56 TB**, largest open-weight checkpoint ever. Variants: K3 Max, K3 Swarm Max; always-on reasoning.
- **Context:** 1,048,576 tokens.
- **License:** **NOT MIT.** Custom **Kimi K3 License**: (1) **MaaS clause — operators/affiliates with >$20M aggregate revenue over any 12 consecutive months must sign a separate commercial agreement**; (2) attribution ("Kimi K3") above 100M MAU or $20M/month revenue.
- **Benchmarks:** AA Index v4.1 **57.1** — 4th of all models (behind Fable 5 59.9, GPT-5.6 Sol 58.9); v4.3 rebase: **44**. Vendor: Terminal-Bench 2.1 88.3; topped Claude Fable 5 on Frontier Code Arena (human-judged); launch stunts: autonomous 48h chip design, MiniTriton GPU compiler built from scratch.
- **Pricing:** **$3/M in, $15/M out**, cached ~$0.30/M — ~3× K2.6 ("the end of super-cheap Chinese AI"); AA-measured ~$0.94/task; 21% fewer output tokens than K2.6.
- **HF:** `moonshotai/Kimi-K3` + ModelScope; vLLM/SGLang/TokenSpeed recipes.
- **Market impact:** launch triggered AI/semiconductor stock sell-off ("DeepSeek moment" parallels).

---

## PART D — GLM (Zhipu AI / Z.ai, Tsinghua spin-out 2019)

### GLM-5.2 — Coding Plan June 13, 2026; API + weights June 16, 2026
Sparse MoE **753B total / ~40B active**; 384 experts; Dynamic Sparse Attention; trained on **28.5T tokens**. **1M context** (from 200K), 128K–131K max output; text-only. **MIT** (`zai-org/GLM-5.2`). Benchmarks: SWE-bench Pro **62.1** (vs GPT-5.5 58.6), Terminal-Bench 2.1 **81.0** (first open-weight >80), AIME 2026 99.2, GPQA Diamond 91.2, AA Index **51 — #1 open-weight at launch**. Pricing $1.40/$4.40, cached $0.26/M (~⅙ of GPT-5.5). Released one day after US Commerce forced Anthropic to disable Fable 5/Mythos 5.

### GLM-5.3 — API August 14, 2026; open weights August 28, 2026
- **Architecture:** 753B/~40B active MoE (78 layers); **same base as GLM-5.2 — gains from post-training only**; text-only; 1M context.
- **Staged release:** weights **withheld ~2 weeks for "safety evaluation and hardening"** (emergent offensive-cyber capability, CyberGym 84.5); shipped on HF (`zai-org/GLM-5.3`, `-BF16`) **Aug 28, 2026**, exactly the promised 14 days; safety work reportedly found 2,436 vulns across 269 OSS projects. 756GB native FP8; fits 8×H200, not 8×80GB H100.
- **License:** **NOT MIT.** Custom **GLM-5.3 License**: MIT-style permissions **plus security-review condition for MaaS operators with >$10B trailing-12-month revenue** (~500× looser than Kimi K3's $20M trigger).
- **Benchmarks:** DeepSWE **66.9** (vs 5.2's 46.2), CyberGym 84.5, AutomationBench 48.8 (vs Opus 4.8 41.0), AA v4.3 **45**. AA measures ~$2.15/M blended.
- **Distinguishing:** first explicit **staged open-weight release (pause → harden → ship)**; cyber-capability framing drove both the delay and the license change.

### GLM-5.3-Flash — August 26, 2026 (stealth-tested as "Ox Alpha")
- **Architecture:** **320B total / 18B active** MoE (45 layers); hybrid **KDA linear + NoPE sparse MLA** (~3× less attention compute, 4.4× smaller KV via IndexPool); **natively multimodal** (first in GLM-5 family); **1M context**; native FP8 (~306 GiB).
- **License:** **MIT**. HF: `zai-org/GLM-5.3-Flash` + `-BF16` (321B params).
- **Benchmarks:** DeepSWE v1.1 **63.4** (vs Opus 4.8 58.0), AutomationBench 48.8 (vs Opus 4.8 41.0), Terminal-Bench 2.1 **84.3** (vs Opus 4.8 85.0 — near-parity), AA Index v4.1.1 **57** at ~$0.045/task. Vision is the weak flank.
- **Pricing:** **$0.15/M in, $0.50/M out, $0.03/M cached**; launch half-price promo through Sept 9 (~1/33 in, ~1/50 out of Opus 4.8's $5/$25).
- **The 100,000-chip story:** stealth-tested as **"Ox Alpha"** on OpenRouter/OpenCode — became **most-used model on both, >62T tokens in 6 days (31% of OpenRouter weekly traffic)**. All inference on **>100,000 domestic Chinese accelerators** (Cambricon, Huawei, Moore Threads) with custom SGLang-based disaggregated engine; **Infra Agent powered by GLM-5.3 did most optimization** — first boot on domestic silicon to all production traffic in **~13 days, 3.2× throughput vs baseline**. Serving via SGLang, vLLM, TokenSpeed, KTransformers.

### GLM-5.3-FlashX — September 18, 2026
**Identical weights to GLM-5.3-Flash** (320B/18B, 1M) — a pure **inference-configuration** variant for the domestic-chip fleet. Output **200 tok/s** (~3.2–5× faster than Flash depending on metric); license/pricing presumed same as Flash.

---

