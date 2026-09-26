---
id: ai-industry-kb-2026-wave6/17-other-chinese-labs/part-4
title: "§17. Other Chinese Labs (part 4)"
domain: other-chinese-labs
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Baidu", "Falcon", "Huawei", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "TII", "Z.ai", "xAI"]
dates: ["2024-10", "2026-01", "2026-01-01", "2026-05-07", "2026-07", "2026-08"]
keywords: ["acquisition", "agent", "apache", "ascend", "attribution", "benchmark", "benchmarks", "claude", "cost", "distillation", "distribution", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8386, 8414]
section: "§17. Other Chinese Labs"
delta_of: ai-industry-kb-2026
sha256: 14c1fe55c43b34fbcd3ad737f35b101e9b8e367a696bb480f4eea308b0433887
---

# §17. Other Chinese Labs (part 4)

60. Hy4 Preview architecture per independent technical analysis: 78-layer backbone (1 dense FFN layer + 77 MoE layers), 256 routed experts plus 1 shared expert with top-8 activation per token, a native MTP layer (10B total / 0.7B active) for speculative decoding, Gated DSA with IndexCache, and Identity Hyper-Connections [SECONDARY]. (S41, S42)
61. Official Tencent API pricing for Hy4 Preview: $0.834/M input tokens, $2.501/M output tokens, $0.042/M cache-hit tokens; independent OpenRouter measurement reports ~40 tok/s throughput at 3.49 s average latency [SECONDARY]. (S41, S42)
62. Self-hosting weight: BF16 Hy4 weights are ~1.56 TB; AngelSlim published GGUF builds (STQ1_0 at 213.66 GiB, Q4_K_M at 435.20 GiB, UD-IQ1_M at 219.83 GiB), but full residency needs ~214 GiB VRAM and the builds require patching — they do not run on stock llama.cpp; the vendor-reported minimal deployment is an 8-GPU cluster; single secondary coverage [SECONDARY]. (S42)
63. Hy4 Preview's internal blind evaluation (163 professional experts, 203 real-world engineering tasks): Tencent reports Hy4 Preview scoring 2.99/4.00 versus GLM-5.3 at 2.92 and Kimi K3 at 2.94 — a vendor-run parity claim, not an independent benchmark [VENDOR]. (S41 — single vendor source coverage)
64. Hy-lineage detail new to this corpus: Hy3 expanded internationally via the WorkBuddy app and Tencent Cloud TokenHub in August 2026, per Wikipedia's Tencent Hy page [SECONDARY]. (S40 — single secondary coverage)
65. Tencent claims Hy4 Preview is the first Hunyuan model that helped optimize its own training pipeline, improving throughput by 31.8% — vendor claim via a spec aggregator (already cited in the base §17 sources), not independently verified [VENDOR]. (S80 — single vendor source coverage)
66. CONTRADICTION PRESERVED: one technical analysis distinguishes the 770B backbone from ~780B when the native MTP layer is included; sources are inconsistent about which figure is "the" parameter count [SECONDARY]. (S41, S42)
67. openPangu 2.0 distribution detail beyond the base facts: the staged release from June 30 put Flash out first through GitCode Ascend Tribe with weights, inference code, and training/inference operators [SECONDARY]. (S44, S86)
68. openPangu 2.0 specs beyond the base corpus: ~28:1 sparsity ratio (505B/18B), trained on 34 trillion tokens on Ascend NPUs, with a full-stack release — architecture, weights, technical report, inference, pre-training, post-training, and Ascend operators [SECONDARY]. (S44, S45)
69. The official license of openPangu 2.0 is UNVERIFIED — staged sources describe open-sourced components but no confirmed license text has been located; do NOT call it Apache 2.0 or MIT [UNVERIFIED]. (S44 — single unverified coverage)
70. Huawei's own claim: openPangu 2.0 Pro delivers ~2× single-card throughput on Ascend versus other major open models of comparable scale; no independent benchmark (Artificial Analysis/BenchLM) had published results at the time of reporting — Huawei's figure must not be presented as an independent measurement [VENDOR]. (S46, S45)
71. Community NVFP4 CUDA build of openPangu 2.0 Flash reports ~56.9 GB resident memory on a 24 GB RTX 4090 — single community benchmark doc [COMMUNITY]. (S78 — single community coverage)
72. Supply-chain framing dispute PRESERVED: one secondary outlet questions the "trained without Nvidia" narrative, saying the supply chain "tells a different story" — the claim is contested, not settled [SECONDARY]. (S46, S87)
73. Falcon H1R 7B's new detail: "Deep Think with confidence" (DeepConf) prunes low-quality reasoning traces via the model's own confidence scores; TII claims it matches reasoning models 2–7× larger, published as open source on Hugging Face [SECONDARY]. (S47, S48)
74. Falcon-H1 Arabic: hybrid Mamba-Transformer architecture (a departure from previous transformer-only Falcon versions), built Arabic-first rather than translated, in 3B/7B/34B sizes, with context up to 256K tokens [VENDOR]. (S48, S47)
75. Vendor-reported OALL (Open Arabic LLM Leaderboard) results: 3B at 61.87% (~10 points ahead of Phi-4 Mini 4B), 7B at 71.47% (surpassing ~10B-class models including Qatar's Fanar-1-9B and Saudi HUMAIN's ALLaM 7B), 34B at 75.36% (outperforming Qwen2.5 72B and Llama-3.3 70B); plus 3LM (STEM), ArabCulture, and AraDice (dialects) suites — vendor figures, not independent [VENDOR]. (S48, S49)
76. The Poolside deal's reporting trail is new: Newcomer via a Poolside investor letter, confirmed by Bloomberg — $12B pre-money valuation for the $1B investment, offers to 109 of roughly 115 engineers/researchers, the three co-founders staying and Poolside operating independently [SECONDARY]. (S50, S51)
77. The letter's own framing: "not an acquisition and not an acquihire" — Poolside may continue licensing Model Factory to other buyers [SECONDARY]. (S50, S54)
78. The $6B fee is set to be distributed to existing investors by end of 2027; the deal values Poolside at ~4× its ~$3B mark a year earlier; Poolside has raised $1.6B since 2023 (last: $500M Series B, October 2024); co-founded by former GitHub CTO Jason Warner and Eiso Kant [SECONDARY]. (S51, S52)
79. Model Factory: RL from code execution — running millions of coding tasks against real execution results instead of human preference ratings — built by a team of fewer than 70 engineers [SECONDARY]. (S53, S50)
80. Poolside's open Laguna family (July 2026, OpenMDW-1.1 license): XS 2.1 (33B total / 3B active), M.1 (225B/23B), S 2.1 (118B); reported benchmarks: M.1 72.5% SWE-bench Verified, XS 2.1 70.9%, S 2.1 70.2% on Terminal-Bench 2.1 [SECONDARY]. (S51, S53)
81. ERNIE 5.0's new technical detail: ~2.4T-parameter native-multimodal MoE with a unified text/image/audio/video architecture and fewer than 3% of parameters activating per query [SECONDARY]. (S55, S82)
82. LMSYS Arena (January 2026, secondary reporting): Ernie 5.0-0110 text score edged GPT-5.1-High, sitting among Claude Opus 4.5 (1,467) and Grok 4.1 (1,466) — community/secondary reporting of arena votes, not a vendor benchmark; single secondary coverage [COMMUNITY]. (S75)
83. ERNIE 5.1 efficiency claims (new): hyper-sparse MoE that dynamically generates sub-models during a single training run; Baidu claims 94% lower pre-training cost (~6% of industry standard), total parameters compressed to ~1/3, active parameters to ~1/2, and 35% lower single-response latency [SECONDARY]. (S55, S56)
84. Four-stage post-training pipeline for 5.1 addresses Baidu's "seesaw effect" (code/logic/creativity dragging each other down): joint SFT → parallel expert models (code, reasoning, agent) → distillation into a student model → open RL for dialogue and creativity [SECONDARY]. (S55 — single secondary coverage)
85. CONTRADICTION PRESERVED — ERNIE pricing (third source): a French consulting document prices ERNIE 5.1 at $3/M input and $12/M output via the Qianfan API, versus Claude Mythos 5 at $5/$25 and GPT-5.6 Sol at $7/$35 — against the base corpus's $0.85/$3.40 (5.0) and $0.59/$2.65 (5.1). All three figures must be cited; no blended figure [SECONDARY]. (S57, S83)
86. No open weights for ERNIE 5.0 or 5.1 — the-decoder explicitly notes the benchmark and efficiency claims cannot be independently verified; DeepLearning.AI likewise describes 5.0 as proprietary [SECONDARY]. (S55, S82)
87. Kunlun Chip (Baidu's chip arm): confidential HK main-board A1 filing January 1, 2026; STAR Market listing coaching initiated May 7, 2026 (CICC) — a dual A+H path [SECONDARY]. (S56 — single secondary coverage)
88. Ling-3.0-Flash license resolution (new): community sources (best-of-ai catalog, sglang cookbook) attribute an MIT license with Hugging Face availability — superseding the base corpus's license-[UNVERIFIED] tag with community MIT attribution, still not confirmed from vendor legal text . (S59, S60)
