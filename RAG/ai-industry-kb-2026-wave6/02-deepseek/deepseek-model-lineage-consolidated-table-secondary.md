---
id: ai-industry-kb-2026-wave6/02-deepseek/deepseek-model-lineage-consolidated-table-secondary
title: "DeepSeek model lineage — consolidated table [SECONDARY]"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Baseten", "China", "DeepSeek", "Huawei", "Meta", "Nvidia", "OpenAI", "SGLang", "United States", "Z.ai", "vLLM"]
dates: ["2023-11-29", "2024-05-06", "2024-09-05", "2024-12-26", "2025-01-20", "2025-05", "2025-08-21", "2025-12-01", "2026-04-24", "2026-04-29", "2026-05", "2026-06", "2026-06-03", "2026-07-24", "2026-07-25", "2026-07-28", "2026-09", "2026-09-11", "2026-09-14", "2026-09-16", "2026-09-22"]
keywords: ["deepseek", "sol", "agentic", "agents", "agi", "ascend", "awq", "benchmark", "blackwell", "compute", "consumer", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [708, 778]
section: "§2. DeepSeek"
sha256: 4163d1d07905c6c15d7e9a3b3f077d16631c9f1ba2faa84780ba81a6ae3d879c
---

# DeepSeek model lineage — consolidated table [SECONDARY]

### DeepSeek model lineage — consolidated table [SECONDARY]
- One community inference guide consolidates the full lineage (all MIT open-weight MoE unless noted):
  - **DeepSeek-LLM 7B/67B** (2023-11-29): 7B/67B dense, 4K, MIT.
  - **DeepSeek-Coder 1.3B–33B** (2023-11): 16K, MIT/DeepSeek License.
  - **DeepSeek-V2/Lite** (2024-05-06): 236B (16B Lite), 21B (2.4B) active, 128K (32K Lite).
  - **DeepSeek-V2.5** (2024-09-05): 236B, 21B active, 128K.
  - **DeepSeek-R1/R1-Zero** (2025-01-20): 671B, 37B active, 128K, MIT.
  - **DeepSeek-V3** (2024-12-26): 671B, 37B, 64K (128K via V3.1).
  - **DeepSeek-V3.1/Terminus** (2025-08-21): 671B, 37B, 128K, hybrid reasoning.
  - **DeepSeek-V3.2/Speciale** (2025-12-01): 671B, 37B, 128K, MoE + DSA.
  - **DeepSeek-V4-Flash/Base** (2026-04-24): 284B, 13B, 1M.
  - **DeepSeek-V4-Pro/Base** (2026-04-24): 1.6T, 49B, 1M.
- Cadence note: from ~1 major family/year (V2→V3) to multiple substantial releases/year by 2025–2026, "with a clear shift toward long-context efficiency and agentic/reasoning-first capabilities." [COMMUNITY]

### V4.1-Flash vs Union Alpha — the September 2026 contrast [SECONDARY]
- One comparison piece (OrcaRouter) frames V4.1-Flash against **Union Alpha** (launched 2026-09-16, six days after V4.1-Flash) as opposite objects: V4.1-Flash has leaderboard entries, vendor benchmark tables, and a model card; Union Alpha (at the time) had none of the three. [SECONDARY]
- Spec contrast: V4.1-Flash 1M context / 384K max output vs Union Alpha 262K / 131K; V4.1-Flash text+image input (native vision) vs Union Alpha's narrower surface. [SECONDARY]
- The piece's verdict pattern — open weights + published card + reproducible serving recipes vs closed announcement — is the corpus's template for grading September 2026 releases by evidence quality, not headline claims. [DIRECTIONAL]
### V4.1-Flash — audited revision and technical report [SECONDARY]
- One community audit (alkinun/speck) pins the release: repository `deepseek-ai/DeepSeek-V4.1-Flash`, audited revision **`df42c109f1defefcbfcedbe7d905718a12266e40`**, technical report `DeepSeek_V41_Tech_Report.pdf` (1,809,802 bytes, SHA-256 `ba68e2e4…`). The Hub API counts **~484.6B stored tensor elements** because routed-expert FP4 values are packed — not the logical model size. The artifact occupies **~510 GB**. [COMMUNITY]
- CSA2 roles, per the report: **Full** (produce main KV and indexer K, compute new Top-K selection); **Reindex** (reuse main KV and indexer K, fresh depth-specific selection); **Reuse** (reuse representation and latest selection). [SECONDARY]
- **Engram internals** (per secondary technical write-ups): two modules, half the 196B each; looks up patterns spanning **2, 3, and 4 tokens**; **8 hash heads**; ~16M entries per head; **FP8 tables**; lookup address predictable from input so retrieval from host memory can start while the transformer computes. [SECONDARY]

### V4.1-Flash — hardware reality [SECONDARY]
- The checkpoint is **510 GB across 48 shards** (476 GiB); V4 Flash was 166.9 GB on two H200s — the replacement is **3× the size**, and day-one recipes from vLLM, SGLang, and NVIDIA start at **four Blackwell-class GPUs or eight H200s**. "The model's clever memory tricks are real, but they're on the cache side. The weights are the problem." [SECONDARY]
- GPU memory floor: **~614 GB with headroom** (vLLM recipe). Precision as shipped: **FP8 dense, FP4 routed experts, FP8 Engram** — "the easy compression is already spent," so no "just quantize it" path; GPTQ/AWQ re-quantization is neither necessary nor recommended. [SECONDARY]
- Engram is **~183–189 GiB** of the total and **not optional** — every token touches it; SGLang can move it to host RAM, the one route to fewer GPUs. [SECONDARY]
- Context: **1,048,576 tokens**; max output 384K; DSpark built in (5 draft tokens/step in recipes); text+image input (day-one recipes serve text only). [SECONDARY]
- **Baseten** added V4.1-Flash to its Model APIs on **2026-09-11** — the day after launch — describing it as DeepSeek's **third open-weight flash release of 2026** and "the only model of its scale" using CED; support for Baseten's Loops training product was announced as coming soon. [SECONDARY]
- **First non-experimental native image input**: V4.1-Flash is DeepSeek's first non-experimental model with native image understanding — previously limited to the experimental **V4-Flash-Vision-Exp**. Card figures: **Chartography 78.9 vs 64.3**, **ZeroBench 49 vs 35** (new vs experimental). [SECONDARY]
- Launch-day cross-lab comparisons (vendor card): **GPQA Diamond 90.9**, **Codeforces 3471**, **HLE with tools 63.9**; **CyberGym 88.1 vs 84.5 for GPT-5.6 Sol and GLM 5.3** (per one secondary write-up). [SECONDARY]
- The size dispute, reconciled: one community analysis (r/LocalLLaMA, cited by OrcaRouter) reads the checkpoint as **748B total, ~510 GB in FP8** — the reconciliation offered is that DeepSeek's "552B" counts only the compute backbone while the download also carries a large sparsely-accessed memory table looked up rather than computed through. Treat 552B as vendor-reported, 748B as an unrefuted community reading. [COMMUNITY]

### V4 Pro retirement — 2026-09-14 [SECONDARY]
- DeepSeek retired the ~4×-more-expensive **V4 Pro on 2026-09-14**, telling API users: "V4.1 Flash has comprehensively surpassed V4 Pro across all key metrics, including performance, cost, speed, and task completion time." Legacy **V4 Flash and V4 Pro requests are being routed to V4.1 Flash at Flash pricing** — the corpus must treat `deepseek-chat`/`deepseek-reasoner` traffic as landing on V4.1-Flash after this date. [SECONDARY]
- Migration mechanics: from **04:00 UTC on 2026-09-14**, the `deepseek-v4-pro` endpoint stopped serving V4 Pro and started serving V4.1-Flash at Flash prices — DeepSeek said this holds until a future V4.1 Pro arrives. [SECONDARY]
- The migration table (one secondary source): V4 Pro input ~$0.435 / output ~$0.87 → V4.1-Flash off-peak $0.15/$0.60; context 1M both; **HLE (no tools) 42.7 → 36.8** — a regression; V4.1-Flash is weaker on unassisted HLE even as it wins on tool-using HLE (63.9). [SECONDARY]
- Off-peak output price cut vs V4-Pro: **~70%** ($1.98 → $0.60 per million). [SECONDARY]
- The operational critique (one digest): the silent migration means any prompt tuned against Pro now runs against a different model — "your request semantics, latency profile, and output distribution changed without a config change on your side." The source does not report measured throughput or quality deltas against Pro; treat benchmark claims as vendor statements until independent evals land. [SECONDARY]
- The framing from one commentary: "The 'budget' model got good enough that the 'premium' model no longer earns its markup" — V4.1-Flash (284B→552B, nearly double the previous Flash) is "not small"; the flash label now means serving economics, not model size. [SECONDARY]

### V4.1-Flash pricing — peak/off-peak refined [SECONDARY]
- Off-peak: **$0.15 input / $0.60 output / $0.003 cached input** per million. Peak (Mon–Fri, **01:00–04:00 and 06:00–10:00 UTC**): **$0.30 / $1.20 / $0.006** — exactly doubled. **Weekends are always off-peak.** [SECONDARY]

### Alias retirement — cutoff, then compatibility routing [SECONDARY]
- Official cutoff: `deepseek-chat` and `deepseek-reasoner` aliases retired **2026-07-24 15:59 UTC**. [SECONDARY]
- Third-party observations: **HTTP 400 on 2026-07-25** (aliases dead), then **HTTP 200 compatibility routing on 2026-07-28** — but `/models` still omitted the aliases. Treat post-07-28 behavior as **unsupported compatibility**, not a reversal of retirement: the retirement stands, the routing is best-effort. [SECONDARY]

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

