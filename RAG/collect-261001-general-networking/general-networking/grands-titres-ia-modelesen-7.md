---
id: collect-261001-general-networking/general-networking/grands-titres-ia-modelesen-7
title: "VOLET 1 — Vague 1 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-06", "2026-06-01", "2026-07-16", "2026-09-02", "2026-09-10", "2026-09-21"]
keywords: ["agent", "agentic", "astra", "bedrock", "claude", "cost", "deepseek", "exploit", "fable 5", "gemini", "glm", "grok"]
source: docs/RAG/collect-261001-general-networking/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [489, 510]
sha256: 8bc7cf217857932d8d7d9c75de9509d5fc689707eda01cf23f209b9efa849b12
---

# VOLET 1 — Vague 1 (EN)

1. **Chinese labs own the open-weight coding/agent leaderboard.** Sept 2026 AA open rankings: MiMo-V2.6-Pro (46) > GLM-5.3 (45) > Qwen3.8-Max-0902 (45) > Kimi K3 (44) — all Chinese. SWE-bench Verified 78–80.6% cluster: DeepSeek V4 Pro 80.6%, MiniMax M3 80.5%, Kimi K2.6 80.2%, GLM-5.1 ~78%, MiMo-V2.5 Pro 78.0%.
2. **The efficiency frontier moved to Flash-tier models:** DeepSeek V4.1 Flash (Sept 10, 2026): 552B MoE (8B/16B active), 1M ctx, MIT, $0.14/$0.28 — near-frontier coding at the cheapest frontier-class price. MiMo-V2.6-Flash and GLM-5.3-FlashX play the same game.
3. **Closed models still lead on the hardest evals:** Claude Opus 5.x / Fable 5.1, GPT-5.x–6 (Sol/Astra), Gemini 3.x, Grok 4.x lead on Terminal-Bench 4.0 (Opus 5: 49.0 vs MiMo 34.9), ProgramBench competitive programming (37.0 vs 26.5), exploit-generation (ExploitBench 78.5 vs 47.9), and top AA Index absolute scores (Opus 5 max 63.1 vs best open ~58–60 on v4.1.1 scale).
4. **Cost asymmetry is the story:** open-weight Chinese models deliver ~90–95% of closed-frontier agentic coding at 1/10 to 1/60 the price (MiMo-Pro $0.13/task vs dollars-per-task closed equivalents; M3 at 1/30th of Fable 5-class).
5. **Transparency as strategy:** Xiaomi's live-streamed RL training (public dashboard, published costs, open-sourced RL tooling/environments) contrasts with US labs' closed training — reproducibility as competitive positioning. NVIDIA's Nemotron 3 Ultra (June 2026, 550B/55B, 71.9% Verified, free on OpenRouter) is the first Western open model partially breaking the Chinese monopoly in this tier.
6. **Caveats for RAG:** (a) AA Intelligence Index versions (v4.1.1 vs v4.3/v4.3.2) are NOT comparable across versions — always record the version; (b) vendor-reported SWE/DeepSWE figures are provisional until independently reproduced; (c) "open-weight" ≠ "open-source" — K3's MaaS revenue clause, Qwen3.8-Max's unshipped weights, and MiniMax's license terms all restrict the label.

---

## 6. Quick-reference table (Sept 2026 snapshots)

| Model | Release | Params (total/active) | Context | AA Index (v4.3) | SWE-bench Verified | License | API $/M (in/out) |
|---|---|---|---|---|---|---|---|
| MiMo-V2.6-Pro (Xiaomi) | 2026-09-21 | 1.02T / 42B | 1M, omni | 46 | n/r (DeepSWE 71.9) | MIT | 0.435 / 0.87 |
| MiMo-V2.6-Flash | 2026-09-21 | 309B / 15B | 1M, omni | — | n/r (DeepSWE 67.9) | MIT | 0.14 / 0.28 |
| Kimi K3 (Moonshot) | 2026-07-16 (wt late Jul) | 2.8T / ~50B | 1M | 44 | — (K2.6: 80.2%) | Open-weight* | via Bedrock/API |
| Qwen3.8-Max-0902 (Alibaba) | 2026-09-02 (GA Aug 3) | 2.4T / 95B | 1M | 45 | — | Proprietary (wt promised) | 2.00 / 6.00 |
| MiniMax M3 | 2026-06-01 | 428B / 23B | 1M, native multi | — | 80.5% | Open-weight‡ | 0.60 / 2.40 |
| DeepSeek V4.1 Flash | 2026-09-10 | 552B / 8–16B | 1M | 39 | 79.0% | MIT | 0.14 / 0.28 |
| GLM-5.3 (Z.ai) | 2026 (mid) | — | 1M (5.2+) | 45 | ~78% (5.1) | MIT | — / 3.20 |

\* K3: "open weight" license with >$20M/yr MaaS separate-agreement clause. ‡ M3: confirm commercial terms.
