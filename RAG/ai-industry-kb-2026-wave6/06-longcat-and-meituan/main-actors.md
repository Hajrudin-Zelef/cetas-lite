---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/main-actors
title: "Main actors"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["China", "DeepSeek", "Huawei", "LongCat", "Meituan", "OpenRouter", "Z.ai"]
dates: ["2025-08-29", "2025-09-05", "2025-09-22", "2025-12-22", "2026-01-14", "2026-02-05", "2026-03-11", "2026-03-12", "2026-04-20", "2026-05-29", "2026-06-12", "2026-06-29", "2026-06-30", "2026-07-01", "2026-07-05", "2026-08-28", "2026-09", "2026-09-02"]
keywords: ["apache", "attention", "benchmark", "consumer", "cost", "deepseek", "glm", "inference", "license", "moe", "omni", "open weights"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2911, 2995]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: 105ee3167665ef3f3f4901873e683cb5ae14198a6a0fa6364efeb6d02e419e9e
---

# Main actors

## Main actors

- **Meituan (LongCat team)** — vendor of the LongCat line; ran the Owl Alpha stealth episode; claims domestic-only training hardware for LongCat-2.0. [VENDOR]
- **Tencent (Hunyuan)** — vendor of Hy4 Preview (Apache 2.0) with the Hy3-Preview/Hy3 license split. [VENDOR]
- **Huawei (Noah's Ark Lab)** — vendor of openPangu 2.0; official license unverified. [SECONDARY]

## Timeline and context

- **2025-09-22/23** — LongCat-Flash-Thinking original release. [VENDOR]
- **~2025-11** — LongCat-Flash-Omni (560B/27B, 128K, MIT). [VENDOR]
- **2026-01-14** — LongCat-Flash-Thinking-2601 effective (Meituan changelog; reporting also cites Jan 16). [VENDOR]
- **2026-05-29** — Six-model LongCat Flash service sunset. [VENDOR]
- **2026-06-12** — Huawei announces openPangu 2.0 at HDC. [SECONDARY]
- **2026-06-29** — Meituan reveals the Owl Alpha stealth episode (~10.1 trillion tokens/month). [SECONDARY]
- **2026-06-30** — LongCat-2.0 revealed; openPangu 2.0 staged release begins on GitCode. [VENDOR]
- **2026-07-01** — Secondary coverage frames Owl Alpha as a stealth model that had been "quietly topping OpenRouter all along" (bitcoinlfg) — reception framing, not vendor claim. [SECONDARY]
- **2026-07-05** — LongCat-2.0 weights + inference code, MIT. [VENDOR]
- **2026-08-28** — Tencent Hy4 Preview released and open-sourced (Apache 2.0). [VENDOR]


### New verified timeline entries — expansion (continued — platform history)

- **2025-08-29** — LongCat-Flash-Chat release [VENDOR].
- **2025-09-05** — LongCat API Platform launch [VENDOR].
- **2025-09-22** — LongCat-Flash-Thinking release [VENDOR].
- **2025-12-22** — Flash-Chat upgraded (256K ctx, 9 languages) [VENDOR].
- **2026-01-14** — LongCat-Flash-Thinking-2601 release [VENDOR].
- **2026-02-05** — LongCat-Flash-Lite release (68.5B/3B, N-gram table) [VENDOR].
- **2026-03-11** — LongCat-Flash-Omni-2603 release [VENDOR].
- **2026-03-12** — Flash-Thinking auto-routed to 2601 [VENDOR].
- **2026-04-20** — LongCat-2.0-Preview (5M tokens/day quota) [VENDOR].
- **2026-05-29** — Six-model sunset (exact list above) [VENDOR].
- **2026-06-30** — LongCat-2.0 release + billing (Token Pack 30-day, Pay-As-You-Go) [VENDOR].
- **2026-09-02 / 09-10** — Enterprise services (verification, invoicing, tiered discounts) [VENDOR].


### New verified timeline entries — expansion (continued)

- **2026-01** — LongCat-Flash-Thinking-2601 technical report (arXiv 2601.16725); model card with full vendor benchmark table [VENDOR/SECONDARY].
- **2026** — LongCat-Flash-Thinking-ZigZag released (sparse-attention variant) [VENDOR].
- **2026** — LongCat-Flash-Thinking GitHub repo documents DORA RL framework and two-phase pipeline [VENDOR].


### New verified timeline entries — expansion

- **2025-09** — LongCat-Flash technical report (arXiv 2509.01322): 560B ScMoE, shortcut-connected MoE, zero-computation experts [SECONDARY] (arxiv.org).
- **2026-05-29** — Six-model platform sunset (existing §6); exact retirement list still unrecovered — open gap [DIRECTIONAL].
- **2026-06-30** — LongCat-2.0 unveiled [SECONDARY] (cryptobriefing.com).
- **2026-07-01** — LongCat-2.0 weights released: 1.6T MoE, 1M context, MIT, meituan-longcat org; already leading OpenRouter pre-release [SECONDARY] (cryptobriefing.com; venturebeat.com).
- Existing §6 anchors retained: LongCat-2.0 release coverage, 10.1T tokens/month platform volume, Owl Alpha.

## Implications

1. Thinking-2601's canonical date is the changelog's Jan 14; the Jan-16 reporting variant is a secondary-citation artifact — vendor changelogs beat coverage dates. [DIRECTIONAL]
2. The Owl-Alpha trillion/billion unit error is a live example of how secondary writeups distort magnitudes; always re-check the source unit on token-volume claims. [DIRECTIONAL]
3. LongCat-2.0's domestic-hardware claim joins DeepSeek V4 and GLM-5.1 as 2026's vendor-reported domestic-silicon set — three independent claims, zero independent verifications, all tagged accordingly. [DIRECTIONAL]
4. June 30, 2026 was a double event day: LongCat-2.0's reveal and openPangu 2.0's staged release — Chinese open weights hit the 500B+ tier from two labs on the same date. [DIRECTIONAL]
5. The Meituan changelog (Jan 14) vs coverage (Jan 16) gap is a two-day secondary-citation lag — prefer primary sources for effective dates across all six sections, not just this one. [DIRECTIONAL]
6. The 5-day reveal-to-weights gap (June 30 → July 5) is the fast end of the corpus's reveal/release spectrum — compare with API-only releases whose weights are promised months out. [DIRECTIONAL]


### New verified implications — expansion (continued)

- **Three price surfaces coexist for the same model**: $2 token packs (retail), $5–10/month aggregator subscriptions (OpenCode), and direct API pay-as-you-go — LongCat is being sold like a consumer good, a subscription, and a utility simultaneously [DIRECTIONAL].


### New verified implications — expansion (continued)

- **LSA-as-evolution-of-DSA** puts LongCat in direct architectural dialogue with DeepSeek — the Chinese labs are iterating on each other's sparse-attention work, not just scaling [DIRECTIONAL].
- **Dynamic activation breaks cost accounting**: per-token pricing for a model whose active parameters vary 33B→56B by token complexity is a pricing fiction — the "cheap tokens" story assumes average-case complexity [DIRECTIONAL].
- **The customized-Muon-at-scale note** (mbrukman) independently corroborates the Muon-family optimizer trend from §5's MuonClip — Chinese labs are converging on Muon variants for trillion-parameter stability [DIRECTIONAL].
- **Token packs "like mobile games"** is the retail packaging of the pay-as-you-go model — the same Token Pack mechanism in the official ChangeLog, now with a $2 entry point [DIRECTIONAL].


### New verified implications — expansion (continued)

- **The sunset list resolves the open question**: May 29 retired Chat, Thinking, Thinking-2601, Omni-2603, Lite, and Chat-2602-Exp — meaning the Flash family was wound down as a *line*, not just selected checkpoints, to make room for 2.0 [DIRECTIONAL].
- **API lifecycles are measured in months**: Omni-2603 (March 11) was sunset by May 29 — developers pinning to dated LongCat checkpoints should expect ~2–3-month support windows [DIRECTIONAL].
- **Quota-for-feedback (5M → 120M tokens/day)** is a clever RLHF flywheel: free inference in exchange for structured model feedback [DIRECTIONAL].
- **Enterprise billing features (invoicing, tiered discounts) landed in September 2026** — the platform is maturing from developer playground to procurement-ready in the same quarter as the 2.0 launch [DIRECTIONAL].
- **In-place vs dated upgrades coexist**: Flash-Chat was upgraded silently under one name, while Thinking got dated checkpoints (2601) and auto-routing — two versioning philosophies in one platform [DIRECTIONAL].


### New verified implications — expansion (continued)

