---
id: ai-industry-kb-2026-wave6/07-minimax/main-actors
title: "Main actors"
domain: minimax
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "ByteDance", "China", "DeepSeek", "EU", "Hugging Face", "Meituan", "MiniMax", "Moonshot", "OpenAI", "Stability AI", "United States"]
dates: ["2025-10-27", "2025-10-28", "2025-12-30", "2026-01-01", "2026-01-09", "2026-02-12", "2026-02-13", "2026-03-13", "2026-03-18", "2026-04-12", "2026-06-01", "2026-06-03", "2026-06-06", "2026-06-17", "2026-06-18", "2026-07-31", "2026-08-02", "2026-08-03", "2026-08-26", "2026-09-19"]
keywords: ["agent", "agentic", "apache", "attention", "attribution", "benchmark", "benchmarks", "claude", "cost", "deepseek", "distillation", "distribution"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [3401, 3563]
section: "§7. MiniMax"
delta_of: ai-industry-kb-2026
sha256: 8df419ca448374e15d3c81234185da942b33519cf61fdf183a49fadeb27ef032
---

# Main actors

## Main actors
- **MiniMax** — Chinese AI lab; runs the MIT → modified-MIT → non-commercial → geo-excluding Community License tightening sequence across M2→M2.5→M2.7→M3 [SECONDARY].
- **Skyler Miao** — MiniMax R&D lead whose diagram originated the 9.7×/15.6× MSA figures [VENDOR].
- **Hailuo** — MiniMax's video product line/app; H3 is its third-generation branding [SECONDARY].
- **Hugging Face** — hosts the H3 open weights (2026-08-03) [SECONDARY].

## Timeline and context
- **2025-10-27** — MiniMax M2 released under plain MIT [SECONDARY].
- **2026-02-12** — M2.5 released; modified-MIT with "MiniMax M2.5" UI attribution requirement [SECONDARY].
- **2026-03-18** — M2.7 announced (API-only initially); weights ~2026-04-12 under non-commercial terms; MiniMax retreats from "open source" to "open weights" wording [SECONDARY].
- **2026-06-01** — MiniMax M3 released: 428B/~23B MoE, 1M context, MiniMax Sparse Attention, MiniMax Community License [SECONDARY].
- **2026-07-31** — Hailuo 3 API + Hailuo app release [SECONDARY].
- **2026-08-03** — Hailuo 3 weights on Hugging Face [SECONDARY].
- **2026-04-12 (approx)** — M2.7 weights window reported by MarkTechPost; the part-4 chronology centers weights around this date [SECONDARY].
- **2026-06-01** — the agent-times contradiction piece documented the gap between M2.7's "open source" claim and its non-commercial license [SECONDARY].
- The RITS Shanghai-NYU piece framed M3 as a frontier-coding model with 1M context and sparse attention [SECONDARY].

### Corrections applied (wave-6 discipline)
- **M3 date:** the brief's "March 13, 2026" is corrected to **June 1, 2026** — the resolved fact across the corpus [SECONDARY].
- **M3 license:** the one Part-4 "Apache 2.0" line is discarded; the resolved license is the custom MiniMax Community License [SECONDARY].
- **M2 date:** the Oct 23 outlier is discarded in favor of Oct 27 [SECONDARY].
- **M2.5→M3 license arc:** MIT → modified-MIT (attribution) → non-commercial → geo-excluding Community License — cite the full arc, not any single point, when characterizing MiniMax's 2026 posture [SECONDARY/DIRECTIONAL].


### New verified timeline entries — expansion (continued)

- **2026-06-01** — M3 launch; MiniMax stock +5% to HK$907.5 then −12.38% close; weights promised within ~10 days [SECONDARY].
- **2026-06-06/07** — Kilo independent audit: M3 vs Opus 4.8 bug-hunt [SECONDARY].


### New verified timeline entries — expansion (continued)

- **Late 2021** — MiniMax founded in Shanghai (Shanghai Xiyu Jizhi Technology Co., Ltd.) [COMMUNITY].
- **2026-01** — HK listing (0100.HK) [COMMUNITY].
- **2026-02-13** — M2.5 released (MIT, open weights): 230B/10B, SWE-Bench Verified 80.2% [SECONDARY] (the-decoder.com).
- **Early 2026-04** — M2.7 released (self-evolving, OpenClaw scaffold optimization) [SECONDARY] (particula.tech).
- **2026-04-12** — M2.7 officially open-sourced (weights on HF) [SECONDARY] (marktechpost.com).
- **2026-06-01** — M3 released (1M ctx, MSA, native multimodal) [COMMUNITY] (vibe-investing guide).


### New verified timeline entries — expansion

- **2025-12-30** — Reuters: MiniMax targets up to $539M Hong Kong IPO (25.4M shares, HK$151–165; cornerstone Alibaba, ADIA) [SECONDARY].
- **2026-01-01** — Year-end Hong Kong IPO rush coverage (MiniMax among Chinese AI firms listing) [SECONDARY] (bworldonline.com).
- **2026-01-09** — Planned HK debut; ~$6.5B implied valuation [SECONDARY] (reuters.com).
- **2026-06-03** — MLQ.ai quotes MiniMax on ~100T interleaved text+image training tokens for M3 (disputed — see contradiction log) [SECONDARY].
- **2026-06-17** — MarkTechPost piece mixes 109B/3T research-model figures into M3 coverage [SECONDARY] (techtimes.com).
- **2026-06-18** — TechTimes clarifies: 109B/3T = MSA research testbed; production M3 token budget undisclosed as of that date [SECONDARY].
- **2026-07-31** — MiniMax H3 (Hailuo 3.0) announced; teased #MiniMaxH3 the prior day [SECONDARY].
- **2026-08-02** — Third-party H3 model card reviewed; Stability AI H3Max (paid speed-optimized fine-tune) noted [SECONDARY] (ai-tools-mcp card).
- **2026-08-03** — H3 weights on Hugging Face; ComfyUI support merged same day; GGUF/INT4/NVFP4 quants within 24h [SECONDARY] (runaihome.com).
- **2026-08-26** — Independent H3 pricing/positioning analysis ($7.80/min vs $20–22/min rivals) [SECONDARY] (justbeingresourceful.com).
- **2026-09-19** — aifoss.dev licensing survey: revenue-cap/MAU-cap/regional-exclusion clause taxonomy directly applicable to MiniMax terms [SECONDARY].

## Implications
1. MiniMax is the corpus's load-bearing case study for the 2026 license-tightening trend: capability goes up while license freedom goes down, within the same lab (MIT → attribution → non-commercial → geo-exclusion) [DIRECTIONAL].
2. M3's 1M context plus sparse-attention claims target agentic/coding workloads; verify MSA figures against vendor measurement conditions before citing — three different vendor numbers circulate [DIRECTIONAL].
3. The MiniMax Community License's US/EU/UK/South Korea local-deployment exclusion is the most geographically aggressive community license in the corpus — a procurement blocker for Western on-prem use, distinct from Llama's narrower EU-multimodal carve-out [SECONDARY/DIRECTIONAL].
4. H3 shows MiniMax extending open weights to the video line (third Hailuo generation), not just text models [SECONDARY].
5. Cite M3 pricing only with both the standard and promotional figures labeled; the master contradictions log explicitly flags the gap [SECONDARY].
6. The one outlier M2 date (Oct 23) is discarded in favor of Oct 27; cite the corrected date only [SECONDARY].


### New verified implications — expansion (continued)

- **SCMP's framing ("challenges ByteDance")** confirms the H3 competitive set is domestic as much as Western — Kling and Seedance are the price anchors, not just Sora/Veo [DIRECTIONAL].


### New verified implications — expansion (continued)

- **"Open weights" with a territorial license** is a new category in this knowledge base — the weights are downloadable but not legally runnable in the four biggest developer markets; enterprises must read the Applicable Territory clause, not just the "open" headline [DIRECTIONAL].
- **The 2K regeneration module being API-only** means the marketed capability is never actually in the open weights — the classic open-core split applied to a media model [DIRECTIONAL].
- **The lawsuit's May 26 discovery ruling** is the legal context for every MiniMax video license decision; the May timing (2 months before H3) explains the territorial design [DIRECTIONAL].
- **Video is the last closed-model stronghold**: "one of the last corners of generative AI where closed, expensive models still had a comfortable lead" — H3's price/partial-open playbook is the same one DeepSeek ran on text [DIRECTIONAL].


### New verified implications — expansion (continued)

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

- https://huggingface.co/MiniMaxAI/MiniMax-M3
- https://github.com/semianalysisai/inferencex-app/blob/HEAD/packages/app/content/models/minimax-m3.mdx
- https://github.com/bytedance-iaas/sglang/blob/HEAD/docs/cookbook/autoregressive/MiniMax/MiniMax-M3.mdx
- https://www.techtimes.com/articles/318622/20260618/minimax-m3-takes-open-weight-ai-lead-sparse-attention-architecture-now-verified.htm
- https://www.morphllm.com/minimax-m3
- https://aitoolgrade.com/review/minimax-m3.html
- https://codingfleet.com/blog/minimax-m3-vs-gpt-5-5/
- https://pasqualepillitteri.it/en/news/3934/minimax-m3-chinese-open-weights-coding-model
- https://www.reuters.com/world/asia-pacific/chinese-ai-firm-minimax-targets-up-539-million-hong-kong-ipo-2025-12-30/
- https://bworldonline.com/world/2026/01/01/721832/chinese-ai-firm-minimax-others-launch-hong-kong-ipos-in-year-end-rush/
- https://www.computerworld.com/article/4136474/anthropic-alleges-large-scale-distillation-campaigns-targeting-claude-2.html
- https://huggingface.co/blog/ResterChed/minimax-h3-hailuo-3-0
- https://runaihome.com/blog/minimax-h3-open-weights-local-ai-hardware-guide-2026/
- https://evolink.ai/ai-video-generator/hailuo
- https://pixo.video/blog/what-is-minimax-h3
- https://github.com/minecraft9101010/ai-tools-mcp/blob/HEAD/cards/minimax-h3.md
- https://auraai.app/models/minimax-h3
- https://justbeingresourceful.com/2026/08/26/minimax-h3-the-open-weight-ai-video-model-with-a-licensing-catch-2026/
- https://aifoss.dev/blog/open-weight-licences-that-block-commercial-use-2026/

*Contradictions recorded in this file: (1) M3 production training tokens: ~100T interleaved (vendor-quoted via MLQ.ai, June 3) vs "not publicly disclosed as of June 18" (TechTimes) — unresolved; (2) 109B/3T figures = MSA research testbed, repeatedly misattributed to production M3 — keep separate; (3) MSA speedups 9.7×/15.6× vs 9×/15× vs 14.2×/7.6× — different experiment setups; (4) H3 audio-reference-alone rule: one source says audio clips can be used on their own, the HF FAQ says they cannot — unresolved detail; (5) H3 US/EU/UK/SK local-deployment exclusion is single-source [UNVERIFIED] — license file not reviewed this pass; (6) 512K guaranteed-usable context floor is single-source [UNVERIFIED].*

