---
id: ai-industry-kb-2026-wave6/07-minimax/new-verified-implications-expansion-continued
title: "New verified implications — expansion (continued)"
domain: minimax
role: deep-dive
task: actor-profile
actors: ["Anthropic", "China", "DeepSeek", "EU", "Meituan", "MiniMax", "Moonshot", "OpenAI", "United States"]
dates: ["2025-10-28"]
keywords: ["agent", "attention", "benchmark", "benchmarks", "claude", "cost", "deepseek", "distillation", "distribution", "gqa", "inference", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3479, 3541]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 0a66c93aabbdc3d23f9f13ed1ae8b52aa6063df5579a17389432ef2778332488
---

# New verified implications — expansion (continued)

- **The Kilo audit is the most concrete third-party cost-per-outcome data point in §7**: 18–48× cost gaps for equivalent-or-better outcomes, with the failure analysis (invalid JSON, import-time DB setup, async-in-sync-transaction) showing exactly what the price premium buys [DIRECTIONAL].
- **"KV outer gather Q"** (read-each-block-once, contiguous access) is the concrete engineering claim behind MSA's 15.6× number — a more checkable story than the headline speedup [DIRECTIONAL].
- **Sell-the-news on launch day** (−12.38%) despite strong benchmark claims shows the market pricing model releases as commodities — the HK$5,000 peak of January vs June's reaction tells a sentiment story [DIRECTIONAL].
- **Lightning attention → sparse attention** is a genuine architectural reversal: MiniMax abandoned its previous efficiency bet and returned to sparse attention with a better kernel — efficiency layers are still up for grabs [DIRECTIONAL].
- **Token Plan's unified pool** (text+image+voice+music in one subscription) is a bundling move none of the other three labs in this wave have made [DIRECTIONAL].


### New verified implications — expansion (continued)

- **Self-evolving training as proof point**: 100+ rounds of autonomous scaffold optimization on OpenClaw with a measured 30% gain is the most concrete "AI improving AI" result in this wave — and it shipped as an open-weights model, not a lab demo [DIRECTIONAL].
- **M2.5 → M2.7 → M3 in under four months** (Feb → Apr → Jun) is the fastest flagship cadence in the Chinese open-weights cohort — compare Moonshot's K2.5 → K2.6 → K2.7 arc in §5 [DIRECTIONAL].
- **The M2.7 license conflict is unresolved and material**: "officially open source" (MarkTechPost) vs "non-commercial, separate agreement required" (bofai mirror) — enterprises must read the checkpoint file; this is the same discipline §5's K3 conflict requires [DIRECTIONAL].
- **10B-active competitive with frontier**: M2.7's 78% SWE-Bench Verified at 10B active vs M3's 80.5% at 23B active shows the efficiency curve steepening — active-parameter count is becoming a weak proxy for capability [DIRECTIONAL].
- **Dual-protocol API + broad editor ecosystem** is a distribution strategy aimed at zero-friction switching from OpenAI/Anthropic stacks — the Coding Plan's 10–20× price gap does the rest [DIRECTIONAL].


### New verified implications — expansion

- **MSA is a stack play, not just a model**: open-sourcing the sparse-attention kernel (github.com/MiniMax-AI/MSA) alongside the weights mirrors Moonshot's MoonEP/FlashKDA move in §5 — the 2026 open-weights game is won on serving economics, and both labs are productizing the serving layer [DIRECTIONAL].
- **License friction scales with capability**: H3's $20M revenue cap plus reported US/EU/UK/SK local-deployment exclusion makes it the most capable open-weight video model most Western readers cannot legally self-host — the RunAIHome verdict ("the license says the weights aren't yours to run at home") is the honest summary of the trade [DIRECTIONAL] (runaihome.com).
- **Video API-first distribution**: H3 across EvoLink/Pixo/Aura AI within a week shows MiniMax monetizing through third-party creative platforms while the weights' legal reach is restricted — hosted revenue where self-hosting is fenced off [DIRECTIONAL].
- **Training-token opacity as pattern**: the 100T-vs-undisclosed dispute and the 109B/3T research-vs-production confusion are a case study in why this knowledge base must keep research-testbed and production figures in separate columns with setup labels [DIRECTIONAL].
- **Benchmark version hygiene is load-bearing**: the codingfleet Terminal-Bench 2.1-vs-2.0 warning is the concrete exhibit for the no-mixing rule stated in the task — one version digit changes comparability [DIRECTIONAL].
- **M3's custom SwiGLU/RMSNorm/GQA-index details matter for porting**: inference frameworks ingesting M3 must handle swigluoai (α 1.702, limit 7.0), Gemma-style norms, and the custom XML tool-call namespace format — "OpenAI-compatible" is not plug-and-play here [DIRECTIONAL] (sglang cookbook).
- **Anthropic allegations touch all three non-Meituan labs** (§5 Moonshot, §7 MiniMax, plus DeepSeek): the distillation allegation is now a sector-level overhang on Chinese open-weight labs, and every capability claim from these vendors should be read with that context noted — without treating the allegation as fact [DIRECTIONAL].

## Sources and URLs
- https://github.com/MiniMax-AI/MiniMax-M2
- https://www.caixinglobal.com/2025-10-28/minimax-unveils-m2-model-to-compete-on-speed-and-cost-102376624.html
- https://venturebeat.com/technology/minimaxs-new-open-m2-5-and-m2-5-lightning-near-state-of-the-art-while
- https://www.unite.ai/minimax-open-sources-m2-7-a-self-evolving-agent-model/
- https://www.marktechpost.com/2026/04/12/minimax-just-open-sourced-minimax-m2-7-a-self-evolving-agent-model-that-scores-56-22-on-swe-pro-and-57-0-on-terminal-bench-2/
- https://theagenttimes.com/articles/minimax-m2-7-open-source-claim-contradicts-all-available-evi-cbe1eef5
- https://venturebeat.com/technology/minimax-teases-upcoming-m3-model-with-new-sparse-attention-mechanism-and-15-6x-response-speed-boost
- https://ai-beat.github.io/news/2026/06/minimax-m3-sparse-attention/
- https://rits.shanghai.nyu.edu/ai/minimax-m3-frontier-coding-1m-context-and-sparse-attention/
- https://datanorth.ai/news/minimax-releases-minimax-h3
- https://www.explainx.ai/blog/minimax-h3-open-video-model-hailuo-july-2026/


### New sources — expansion (continued)

- https://justbeingresourceful.com/2026/08/26/minimax-h3-the-open-weight-ai-video-model-with-a-licensing-catch-2026/


### New sources — expansion (continued)

- https://pasqualepillitteri.it/en/news/3934/minimax-m3-chinese-open-weights-coding-model


### New sources — expansion (continued)

- https://github.com/gameworkerkim/vibe-investing/blob/HEAD/TechDoc/MiniMax%20Coding%20Guide/minimax-coding-guide.en.md
- https://github.com/bofai/docs/blob/HEAD/docs/llmservice/models/minimax-m2.7.md
- https://news.saerio.com/new-minimax-m2-7-proprietary-ai-model-is-self-evolving-and-can-perform-30-50-of-reinforcement-learning-research-workflow/
- https://the-decoder.com/minimax-m2-5-promises-intelligence-too-cheap-to-meter-as-chinese-labs-squeeze-western-ai-pricing/
- https://venturebeat.com/technology/new-minimax-m2-7-proprietary-ai-model-is-self-evolving-and-can-perform-30-50
- https://particula.tech/blog/minimax-m2-7-vs-claude-opus-coding-benchmarks
- https://www.marktechpost.com/2026/04/12/minimax-just-open-sourced-minimax-m2-7-a-self-evolving-agent-model-that-scores-56-22-on-swe-pro-and-57-0-on-terminal-bench-2/


### New sources — expansion

