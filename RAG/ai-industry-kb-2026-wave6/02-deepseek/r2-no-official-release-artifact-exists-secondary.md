---
id: ai-industry-kb-2026-wave6/02-deepseek/r2-no-official-release-artifact-exists-secondary
title: "R2 — no official release artifact exists [SECONDARY]"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Alibaba", "China", "DeepSeek", "Huawei", "Nvidia", "United States"]
dates: ["2025-05", "2026-04-29", "2026-05", "2026-06", "2026-06-03", "2026-09", "2026-09-22"]
keywords: ["agents", "agi", "ascend", "benchmark", "compute", "consumer", "cost", "deepseek", "inference", "ipo", "memory", "nvidia"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [756, 778]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: a35318e82e955f512da1560e7f755f7eddb07a30d7a218c5ecdb14480088ec67
---

# R2 — no official release artifact exists [SECONDARY]

### R2 — no official release artifact exists [SECONDARY]
- As of 2026-09-22: **no first-party API ID, model card, pricing row, weight release, or release notice** for R2 exists. Every parameter count, benchmark, price, and release date circulating is unverified. [SECONDARY]
- R2 was originally planned for a **May 2026 launch** (one secondary source says May 2025 — the 2025 date conflicts with the R1 January-2025 lineage and is likely an error; treat May 2026 as the better-attested plan). The Information (via NDTV): DeepSeek has **not yet determined the timing** of R2's release; KR-Asia: DeepSeek **denied reports of an August release**. [SECONDARY]
- The delay's causes, per multiple secondary reports: **no successful training run was achieved on Ascend hardware** despite Huawei engineers deployed on-site; **NVIDIA hardware remained necessary for training** while Ascend was reserved for inference; **longer-than-expected data labeling** for R2's updated training dataset; and Liang Wenfeng's **dissatisfaction with R2's performance** — he pushed the team to prioritize performance over political compliance. [SECONDARY]
- The ecosystem gap, as one report put it: **"Without CUDA, even the most powerful chip remains a nicely packaged problem"** — Huawei lacks a mature CUDA equivalent, and gaps in stability, inter-chip connectivity, and software maturity vs Nvidia are the binding constraint, not raw specs like memory bandwidth. [SECONDARY]
- Competitive cost: the delay gave rivals — notably **Alibaba's Qwen3** — time to gain ground, with Qwen3 reported to have incorporated DeepSeek's core training algorithms while improving efficiency. [SECONDARY]

### DeepSeek's second round — frozen after the leak, September 2026 [SECONDARY]
- The June 2026 first round **did close** — reported at **$7 billion** (confirming the June 3 "set to raise" RMB50B story). [SECONDARY]
- Weeks later, DeepSeek was pursuing a **second round**: at least **¥10B (~$1.4B)** at a **$71–74B pre-money valuation** (~¥480B), earmarked for compute-capacity expansion. [SECONDARY]
- **September 2026** — DeepSeek **verbally froze the second round** after notes from a private four-hour investor meeting leaked and went viral on Chinese social media. The leaked remarks, attributed to Liang Wenfeng: DeepSeek **still relies heavily on Nvidia chips** and **China remains behind the US in AI capability**. Liang's frustration at the leak itself was reportedly part of the freeze decision. Bloomberg/Yicai corroborate the suspension; negotiations remain fluid and could resume. [SECONDARY]
- Leaked strategy notes (unverified authenticity, widely circulated): DeepSeek's **inference margins ~85%** (roughly sixfold profit) — the aggressive open-source pricing as a deliberate squeeze on competitors; **coding agents as the year's top priority**, continuous learning as the next milestone; consumer products and enterprise revenue explicitly framed as **byproducts of an AGI roadmap**, not primary goals; a **possible IPO** being prepared. [SECONDARY]

### Ascend lineage — exact wording constraint [SECONDARY]
- Reuters, **2026-04-29**: Huawei said the **Ascend 950 SuperNode** line supports V4 inference and that **Ascend chips were used for part of V4-Flash training**. [SECONDARY]
- DeepSeek itself did **not** disclose whether NVIDIA chips were also used for V4. Therefore write **"Huawei claimed partial V4-Flash training on Ascend"**, never "V4 was trained entirely on Ascend." [SECONDARY]

### DeepSeek's first external raise — "set to raise," not closed [SECONDARY]
- **2026-06-03** reporting: DeepSeek was **set to raise RMB50B (~$7.4B)** at post-money **RMB350–400B (~$52–59B)**, with Liang Wenfeng himself committing **RMB20B**, Tencent considering **RMB10B**, CATL **RMB5B**, and NetEase, JD.com, and a national AI fund in talks — fewer than ten investors total. [SECONDARY]
- Status is **"slated to raise"**, not confirmed closed. The corpus must not write "closed" until later confirmation arrives. If it closes, it is DeepSeek's **maiden external financing** — the lab had been entirely High-Flyer-funded before. [DIRECTIONAL]

---

