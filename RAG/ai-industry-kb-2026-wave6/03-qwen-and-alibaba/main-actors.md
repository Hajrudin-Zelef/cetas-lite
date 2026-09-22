---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/main-actors
title: "Main actors"
domain: qwen-and-alibaba
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "California", "China", "DeepSeek", "Google", "Meta", "Mistral", "OpenAI", "SGLang", "United States"]
dates: ["2026-02-03", "2026-02-15", "2026-02-16", "2026-02-24", "2026-03-30", "2026-03-31", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-22", "2026-05-18", "2026-05-20", "2026-06-01", "2026-07", "2026-07-15", "2026-07-21", "2026-08-03", "2026-08-05", "2026-08-14", "2026-08-20", "2026-08-26", "2026-09-02", "2026-09-17"]
keywords: ["agent", "agentic", "agents", "apache", "attribution", "benchmark", "benchmarks", "capex", "consumer", "cost", "deepseek", "disclosure"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1360, 1446]
section: "§3. Qwen and Alibaba"
sha256: ce21fdef274484f6a5b4f901f9b4b00c21db822918c8d83f9578dd14ca723181
---

# Main actors

## Main actors

- **Alibaba (Qwen team)** — vendor of the Qwen3.5–3.8 lines; Apache-2.0 open mid-tiers plus custom-licensed Max flagships; stated Qwen 4 in training as of the window. [VENDOR]
- **Community trackers** — r/LocalLLaMA/alextech for the 0902 checkpoint score; SGLang ecosystem for `qwen4_exp` day-0 support. [COMMUNITY]

## Timeline and context

- **2026-02-03/04** — Qwen3-Coder-Next (80B/~3B, Apache 2.0, Gated DeltaNet, 262K, 800K+ coding tasks). [VENDOR]
- **2026-02-16** — Qwen3.5-397B-A17B. [SECONDARY]
- **2026-02-24/25** — Medium open Qwen3.5 models + hosted Qwen3.5-Flash. [SECONDARY]
- **2026-03-30** — Qwen3.5-Omni. [SECONDARY]
- **2026-04-02** — Qwen3.6-Plus (closed). [SECONDARY]
- **2026-04-16** — Qwen3.6-35B-A3B (open). [SECONDARY]
- **2026-04-20** — Qwen3.6-Max-Preview (closed). [SECONDARY]
- **2026-04-22** — Qwen3.6-27B (open). [SECONDARY]
- **2026-05-20** — Qwen3.7-Max (closed). [SECONDARY]
- **2026-06-01** — Qwen3.7-Plus (closed). [SECONDARY]
- **2026-07** — Qwen3.7-Flash snapshot/release. [SECONDARY]
- **2026-08-26** — Qwen3.8-Flash-Next (125B/6B). [SECONDARY]
- **2026-08 (SGLang day-0)** — `qwen4_exp` experimental checkpoint (not a Qwen 4 release). [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902 snapshot. [SECONDARY]
- **2026-09-17** — 0902 checkpoint reported at AA v4.3 = 45 (China lead). [SECONDARY]


### New verified timeline entries — expansion

- **2026-02-15/16** — Qwen3.5-397B-A17B released (Apache 2.0): 397B/17B, 512 experts, 262K context (1M extensible), 201 languages, early-fusion multimodal. [SECONDARY]
- **2026-03-30** — Qwen3.5-Omni released: text+image+audio+video, 113 languages. [SECONDARY]
- **2026-03-31** — Qwen3.6-Plus Preview: agentic coding, 1M context. [SECONDARY]
- **2026-04-20** — Qwen3.6-Max-Preview released: **first closed-weights Qwen flagship**. [SECONDARY]
- **2026-05-18** — Qwen3.7-Max-Preview / Qwen3.7-Plus-Preview hit Arena (Alibaba #6 lab text, #5 vision). [SECONDARY]
- **2026-05-20** — Qwen3.7-Max GA at Apsara Conference: 1M context, 35-hour agent-task claims. [SECONDARY]
- **July 2026** — Qwen3.8-Max previewed at World AI Conference; preview opens 07-19. [SECONDARY]
- **2026-07-21** — Qwen Image 3.0 / 3.0 Pro released (no benchmarks, no weights, no model card). [SECONDARY]
- **2026-08-03** — Qwen3.8-Max GA: 2.4T/95B, $2/$6, 131K output. [SECONDARY]
- **2026-08-14 (approx)** — Qwen3.8-2.4T-A95B open weights released under the Qwen3.8-Max License (custom, not Apache 2.0); Qwen3.8-27B Apache 2.0. [SECONDARY]
- **2026-08-20** — Alibaba June-quarter earnings: AI Cloud +45% to RMB48.44B; AI products RMB12.38B (12th consecutive triple-digit quarter); capex +75% to RMB67.68B. [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902 snapshot: coding/agent post-training, AA Index 45, cost per task $5.41. [SECONDARY]
- **2026-09** — Alibaba's new T-Head AI chip unveiled; 5–10T parameter model plans disclosed; cloud unit reports rationing GPU access. [SECONDARY]
- **2026-07-15** — Qwen app + Doubao shut down user-created AI agents (China humanlike-AI rules); Qwen offered no migration path. [SECONDARY]
- **Early July 2026** — Alibaba cut Qwen3.7-Max ~80% / Qwen3.7-Plus ~60% for international Qoder users off-peak (US developer play). [SECONDARY]
- **2026-08-05** — Qwen Image 3.0 opened to all Qwen AI platform users (was invite-only since 07-21). [SECONDARY]
- **2026-06** — US DoD designated Alibaba a military-linked firm; Alibaba sued in California federal court. [SECONDARY]
- **2026-08-03** — DeepSeek V4-Flash launched same day as Qwen3.8-Max GA (dual Chinese flagship day). [SECONDARY]
- **2026-08-14 (approx)** — Qwen3.8-2.4T-A95B open weights (custom license); Qwen3.8-27B (Apache 2.0). [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902 snapshot (AA Index 45, cost/task $5.41). [SECONDARY]

---

## Implications

1. Alibaba runs a two-speed release policy: open Apache-2.0 mid-tiers (3.5/3.6/3.8 ≤27B-class) plus closed/custom-licensed Max flagships — procurement must read the LICENSE file per checkpoint, not per brand. [DIRECTIONAL]
2. The 0902 snapshot is checkpoint news, not generation news; the "Qwen3.8-Max-0902" name invites misreading as a new model. [DIRECTIONAL]
3. The unfulfilled "Plus will be open source" plan is a reminder to treat vendor open-sourcing promises as intent, not inventory. [DIRECTIONAL]
4. `qwen4_exp` must never appear in a Qwen 4 release timeline; conflating an experimental checkpoint with the announced-in-training generation repeats the V4-Lite class of error. [DIRECTIONAL]
5. Qwen's download scale (>3B cumulative, 4.9× Google on HF in 7 months) is the strongest distribution fact for any open-weight family in the corpus — a vendor press claim, but Reuters-cited and directionally consistent with HF's 41%-China figure. [VENDOR]
6. The DeepSWE vendor/secondary split (56.6 vs 69.3) is the canonical "same model, different harness" case in this corpus — any benchmark claim without a named harness is incomplete. [DIRECTIONAL]
7. Qwen3.7-Plus at $0.276/M input (Sept 2026) shows the price war's Chinese front: closed-flagship capability at open-weight prices, repriced quarterly. [SECONDARY]


### New verified implications — expansion

1. The Qwen licensing arc (custom terms at 72B → all-Apache-2.0 at Qwen3 → custom gate at Qwen3.8-Max) is a pricing-strategy timeline: permissiveness was the distribution strategy when Qwen needed adoption; the gate arrives when the weights are strong enough to sell against. The corpus should treat license permissiveness as **cyclical**, not a one-way trend. [DIRECTIONAL]
2. The $50M-MAAS/AI-Work-Assistant trigger is aimed precisely at inference marketplaces and agent startups — the businesses that would commoditize Alibaba's API. Open weights that you cannot resell above a revenue line create a **middle licensing category** the corpus should name explicitly. [DIRECTIONAL]
3. Qwen3.8-Max's open-weight tie with the paid API (AA 57.7 vs 58.1) while the paid tier withholds vision, non-thinking mode, 1M default context, and built-in tools shows Alibaba is differentiating on **service surface**, not cognition — the weights are the commodity, the API surface is the product. [DIRECTIONAL]
4. The Qwen Image 3.0 launch (no benchmarks, no weights, no model card) is the transparency reversal the corpus should log as a data point: the same lab that published same-day technical reports for 1.0 and 2.0 shipped 3.0 on demos alone. Procurement teams cannot rely on precedent for evaluation rigor. [DIRECTIONAL]
5. The 4,500-token prompt window on Qwen Image 3.0 reframes image models as **document-production tools** (infographics, papers, UIs) rather than aesthetic generators — a category shift the corpus's image-model coverage should acknowledge. [DIRECTIONAL]
6. Alibaba's June quarter (AI Cloud +45%, net income −75%) is the cleanest example of the 2026 capex thesis in the corpus: triple-digit AI product growth for 12 quarters running, paid for by collapsing group margins. The growth and the margin damage are the same trade, reported in the same release. [DIRECTIONAL]
7. Eddie Wu's server-rationing disclosure (full-stack customers prioritized, 2–3 year supply bottleneck, GPU pooling cutting usage 82%) is a capacity-constraint signal from a **supply** side — the counterpoint to the corpus's demand-side capex tallies. [DIRECTIONAL]
8. The AA-Omniscience refusal pattern on Qwen3.7-Max (accuracy down, hallucinations down more) suggests the lab is tuning for **calibration over recall** — a meaningful product choice for enterprise deployment, and the corpus should note it as behavior, not regression. [DIRECTIONAL]
9. Qwen3.6-35B-A3B (35B/3B, SWE-Verified 73.4%) and the 4× RTX PRO 6000 Qwen3.5 serving recipe (259 tok/s) together mark the **prosumer-frontier boundary**: 2026 frontier-class open models now run on 4 high-end GPUs, compressing the gap between "open weights" and "self-hostable." [DIRECTIONAL]
10. The "Plus open / Max closed" two-tier pattern (3.7, 3.8) plus the killed Qwen Code free tier is Alibaba's 2026 monetization arc — the corpus should not describe Qwen as "open source" without the tier qualifier. [DIRECTIONAL]
11. Benchmark-label discipline: every 3.5/3.6/3.7/3.8 score in this section comes from vendor-published tables except the Artificial Analysis and Arena readings — the independent column is narrow (AA Index, Arena Elo, Text Arena rank) and the corpus must keep that asymmetry visible. [DIRECTIONAL]
12. The multimodal-vs-Qwen3.7-Plus (not Max) comparison and the peaked RL scaling curve are the two methodological caveats to carry forward whenever Qwen3.8-Max's vendor table is cited. [DIRECTIONAL]
13. The 2026 policy shift (Qwen3.6-Max-Preview, Qwen3.7-Max API-only) is Alibaba's first sustained turn toward the **OpenAI/Anthropic "frontier-closed, smaller-open" strategy** — the most prolific open-weight family in the world now keeps its best weights closed. The corpus should treat Qwen's openness as **tiered**, not binary. [DIRECTIONAL]
14. The 35-hour Zhenwu M890 demo (10× kernel improvement, >1,000 tool calls) is the most extreme long-horizon agent claim in the 2026 dataset — and it was run on **Alibaba's own silicon the model had never seen**. Vertical integration (chip + model + cloud) is the point, not the demo. [DIRECTIONAL]
15. The July 15 agent shutdown (no migration path, permanent deletion) vs the May 20 "AI factory" launch is the **regulatory ceiling on the consumer agent business** in China — Alibaba can build the factory but cannot let users build companion agents on it. The corpus should track this as a structural constraint, not a product decision. [DIRECTIONAL]
16. The Qoder off-peak price cuts (~80% Max, ~60% Plus) are the clearest evidence of the **Chinese token price war** reaching US developers — Alibaba is buying American developer mindshare with Beijing-time off-peak discounts. [DIRECTIONAL]
17. Qwen3.8-Max's **$2/$6 pricing undercuts flagship US tiers by 4×** (Code Arena parity at $6 vs $25 output) — the price war is now at the frontier, not just the mid-tier. [DIRECTIONAL]
18. The August 3 dual-launch (Qwen3.8-Max GA + DeepSeek V4-Flash) marks the moment **Chinese labs ship flagships on the same day** — the domestic race is now the pacing item, not the US race. [DIRECTIONAL]
19. Qwen3.5-Omni's Thinker-Talker + 113-language ASR + open weights makes it **the only voice-native open omni model** in the 2026 dataset — DeepSeek, Mistral, and Meta have no equivalent. For enterprise multimodal RAG, the embedding (Qwen3-Embedding-8B #1 MTEB) + omni (Qwen3.5-Omni) + reranker (Qwen3-Reranker-8B) stack is a **single-vendor open pipeline**. [DIRECTIONAL]
20. The five-releases-in-two-months cadence (Feb–Mar 2026) is the **operational signature** of the Qwen team — the corpus should treat Qwen's velocity, not any single model, as the competitive fact. [DIRECTIONAL]
21. Qwen3.5-397B-A17B at **$0.60/$3.60** (third-party) vs Qwen3.8-Max at **$2/$6** (official) is the **price ladder** Alibaba built in six months — the 397B open model is the loss-leader, the 2.4T Max is the margin. The corpus should track this as intentional tiering, not cost-plus pricing. [DIRECTIONAL]
22. The **Qwen3.8-Max License** (100M MAU / $20M monthly revenue attribution; $50M MaaS/AI-Assistant separate license; internal-use carve-out; no output-training restriction; no royalty) is the **template for 2026 Chinese frontier licensing** — permissive enough to distribute, gated enough to monetize. [DIRECTIONAL]

---

