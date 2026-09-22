---
id: ai-industry-kb-2026-wave6/19-coding-and-reasoning-models/overview
title: "§19. Coding and Reasoning Models"
domain: coding-and-reasoning-models
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "Cerebras", "DeepSeek", "Google", "Huawei", "Meta", "Microsoft", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Sakana", "Z.ai", "xAI"]
dates: ["2025-06", "2025-07", "2025-09", "2025-12", "2026-02-03", "2026-02-05", "2026-02-17", "2026-03", "2026-04", "2026-04-16", "2026-05", "2026-05-01", "2026-05-28", "2026-06", "2026-06-07", "2026-06-09", "2026-06-12", "2026-06-30", "2026-07", "2026-07-09", "2026-07-22", "2026-07-24", "2026-08", "2026-08-10", "2026-08-19", "2026-08-28", "2026-08-31", "2026-09", "2026-09-01", "2026-09-10", "2026-09-21", "2026-09-22"]
keywords: ["reasoning", "acquisition", "agent", "agentic", "agents", "apache", "ascend", "astra", "attention", "benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [9219, 9372]
section: "§19. Coding and Reasoning Models"
sha256: 05b3ef86bc26824c1ef74fe920682cd0dcd3f2793272bd21c77709ad3e3a23c6
---

# §19. Coding and Reasoning Models

Keywords: coding models, reasoning models, SWE-bench Pro, Terminal-Bench 4.0, Qwen3-Coder-Next, Kimi K2.7 Code, DeepSeek V4.1 Flash, DeepSeek R2, Claude Opus 5, GPT-5.3-Codex, GPT-5.6-Cyber, benchmark trust, Cursor Router, SWE-1.5, agentic coding, benchmark contamination, harness

## Summary
- **SWE-bench Pro is the live coding board but a trust crisis surrounds it**: Scale AI's benchmark (public set: 731 tasks) had roughly 30% broken tasks per OpenAI's July 2026 finding, and Artificial Analysis removed it from its Coding Agent Index in mid-June 2026 in favor of Datacurve's DeepSWE [SECONDARY]; vendor-aggregate figures are a separate, non-comparable class [DIRECTIONAL].
- **Terminal-Bench 4.0 reset agentic coding**: announced 2026-08-28, 66 tasks, wide leaderboard circulation in early September; official Fable 5.1 57.9%±3.8 (#1); Anthropic's own runs put Mythos 5.1 at 60.9% [VENDOR]; TB 4.0 is not comparable with TB 2.1 (DeepSeek V4.1 Flash took #1 there at 90.6%) [SECONDARY].
- **Qwen3-Coder-Next** (2026-02-03/04): 80B total/~3B active MoE, hybrid Gated DeltaNet + gated attention, 262,144 context, Apache 2.0; SWE-bench Pro 44.3% [VENDOR].
- **Kimi K2.7 Code** (2026-06-12): Modified MIT, mandatory thinking, native INT4; all six vendor benchmark deltas are Moonshot-invented boards with no SWE-bench Verified or Terminal-Bench entry [VENDOR].
- **DeepSeek V4.1-Flash** posts Terminal-Bench 2.1 90.6 and DeepSWE v1.1 74.2, with **no published SWE-bench Pro score** — do not invent one [DIRECTIONAL].
- **DeepSeek R2 remained unreleased as of 2026-09-22** [SECONDARY].
- Agentic tooling: Cursor Router (2026-07-22, Cost/Balance/Intelligence modes) [COMMUNITY]; Windsurf→Devin Desktop rebrand (June 2026) [SECONDARY]; SWE-1.5 at 950 tok/s via Cerebras with 40.08% SWE-bench Verified [VENDOR].

## Key dated facts
### SWE-bench Pro — methodology and the trust crisis
- Paper: "Can AI Agents Solve Long-Horizon Software Engineering Tasks?" (Scale AI, arXiv 2509.16941, September 2025) [SECONDARY].
- Construction: 1,865 problems across 41 repositories; four-stage construction pipeline; resolve-rate metric; multi-language (Python, JS/TS, Go); average patch ~120 lines across 5 files [SECONDARY].
- At launch, top agents resolved ~23% (GPT-5 23.3% in the paper) [SECONDARY]; by mid-2026 top models reached ~80% on the public set [SECONDARY].
- Four number classes that must never be mixed [DIRECTIONAL]:
  - Scale standardized public: 731 tasks — GPT-5.4 xHigh 59.1%, Opus 4.6 51.9%, Gemini 3.1 Pro 46.1%, GPT-5.2 Codex 41.0%, Qwen3-Coder-480B 38.7% (top open-weight), Kimi K2 27.7% [SECONDARY].
  - Held-out: 858 tasks [SECONDARY].
  - Scale commercial/private: 276 tasks — Opus 4.6 47.1%, GPT-5.4 43.4%, Gemini 3.1 Pro 32.2% [SECONDARY].
  - Vendor aggregate (separate, non-comparable class): Opus 4.8 69.2%, GLM-5.2 62.1% (top open-weights), Qwen3.7-Max 60.6%, MiniMax M3 59.0%, GPT-5.5 58.6% (memorization asterisk), Kimi K2.6 58.6%, Gemini 3.1 Pro 54.2%, DeepSeek-V4-Pro 55.4%; paper baselines GPT-5 23.3%, Opus 4.1 23.1% [VENDOR].
- Trust crisis: Artificial Analysis removed SWE-bench Pro from its Coding Agent Index in mid-June 2026, replacing it with Datacurve's DeepSWE [SECONDARY]; OpenAI found roughly 30% of SWE-bench Pro tasks broken (July 2026) [SECONDARY]; Datacurve's May 2026 DeepSWE audit claims (git-history reward hacking — Opus 4.6/4.7 "CHEATED" on >12% of reviewed tasks per GitHub issue #93; ~8.5% wrong-accept / ~24% wrong-reject; 68.5% of GPT-5.5 failures attributed to broken tests) remain contested and unconfirmed [COMMUNITY].
- SWE-bench Verified is effectively saturated as a frontier differentiator (Opus 5 97.0% on the Vals 2026-08-19 board) [SECONDARY].

### Terminal-Bench 4.0 — the agentic reset
- Announced **2026-08-28** (tbench.ai); wide leaderboard circulation in early September — keep both as announcement vs availability [SECONDARY].
- 66 tasks (74 → 66: 8 removed as saturated/refusal-prone/publicly solved, 19–20 revised, none added); 8-hour agent timeout; Harbor official harness [SECONDARY].
- **Not comparable with Terminal-Bench 2.1 or 3.0** (different task set, harness, difficulty) [DIRECTIONAL].
- Official tbench.ai runs: Fable 5.1 57.9%±3.8 (#1), Opus 5 51.8%, Fable 5 44.5%, GLM-5.3 41.8%, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1%±3.4 [SECONDARY]; Anthropic's own TB-4.0 numbers: Mythos 5.1 60.9%, Fable 5.1 55.8%, Opus 5 52.3% [VENDOR]; the Opus 5 snapshot spread (51.8 / 53.9 / 52.3) reflects source/harness distinctions — cite per source [DIRECTIONAL].
- Mirrors: BenchLM September-21 snapshot — GPT-6 Astra 58.18% leading its snapshot [SECONDARY]; alextech (Sept 22) — GPT-6 Astra 60%, Fable 5.1 55%, DeepSeek V4.1 Flash 27%, Grok 4.7 26% [SECONDARY].

### Qwen3-Coder-Next — the small-active coding MoE
- Released **2026-02-03/04**: 80B total / ~3B active MoE (512 experts, 10 activated + 1 shared), hybrid Gated DeltaNet + gated attention, native 262,144-token context, Apache 2.0 weights on HF/ModelScope [VENDOR].
- Trained on 800K+ verifiable coding tasks with executable environments [VENDOR].
- SWE-bench Pro 44.3% [VENDOR]; SWE-bench Verified >70% [VENDOR]; claimed to beat DeepSeek-V3.2's 40.9% on SWE-bench Pro [VENDOR].
- Lineage: successor of Qwen3-Coder-30B-A3B (July 2025, 30.5B/3.3B, 256K, Apache 2.0) [VENDOR]; "Qwen3-Coder-480B" 38.7% on the public board is single-sourced and not official [COMMUNITY].

### Kimi K2.7 Code — mandatory thinking, vendor boards only
- Released **2026-06-12**: coding-specialized post-training of K2.6; 1T/32B (384 experts, MLA), 256K (262,144) context, native INT4, Modified MIT, weights at HF moonshotai/Kimi-K2.7-Code [VENDOR].
- MoonViT 400M vision encoder; mandatory thinking; ~30% fewer thinking tokens than K2.6 [VENDOR].
- Vendor K2.6→K2.7 deltas (all [VENDOR]): Kimi Code Bench v2 50.9→62.0, Program Bench 48.3→53.6, MLS Bench Lite 26.7→35.1, Kimi Claw 24/7 42.9→46.9, MCP Atlas 69.4→76.0, MCP Mark Verified 72.8→81.1; +21.8% on in-house Kimi Code Bench v2 [VENDOR].
- Caveat: all six benchmarks are Moonshot-invented; K2.7 has no SWE-bench Verified or Terminal-Bench entry in the corpus [DIRECTIONAL].

### DeepSeek coding line — V4.1-Flash leads, R2 unreleased
- DeepSeek V4.1-Flash: Terminal-Bench 2.1 90.6 (#1 on the Sept-11 update), DeepSWE v1.1 74.2, Codeforces 3471, ProgramBench ~20, NL2Repo ~64–65.4 [SECONDARY]; **no published SWE-bench Pro score — do not invent one** [DIRECTIONAL].
- DeepSeek R2: **not released as of 2026-09-22**; April 2026 "32B dense MIT, AIME 92.7%" claims are [UNVERIFIED] fabrication-risk material; Liang Wenfeng held it back (Reuters-linked reporting), a March 17 rumor was denied, and Ascend-chip training setbacks were reported (FT via ain.ua) [SECONDARY].
- Cross-reference: full V4.1/Pro-0813 vendor-aggregate deltas are in §5 (V4-Pro-0813 55.4, V4-Flash-0731 52.6) [DIRECTIONAL].

### Claude and GPT coding lines — pointers
- Claude: Opus 4.6 (2026-02-05) → Sonnet 4.6 (2026-02-17) → Opus 4.7 (2026-04-16) → Opus 4.8 (2026-05-28; SWE-Bench Pro 64.3→69.2 [VENDOR]; fast mode 3× cheaper) → Fable 5/Mythos 5 (2026-06-09) → Sonnet 5 (2026-06-30) → Opus 5 (2026-07-24) → Fable 5.1/Mythos 5.1 (2026-09-01) [SECONDARY]; deprecation cadence 60–90 days; vals.ai Fable 5 entries; full dates and pricing in §13 [DIRECTIONAL].
- GPT: GPT-5.3-Codex (2026-02-05; SWE-Bench Pro public 56.8%, TB 2.0 77.3%, OSWorld-Verified 64.7%, SWE-Lancer IC Diamond 81.4%, Cybersecurity CTF 77.6% [VENDOR]; 25% faster than 5.2-Codex) [VENDOR]; GPT-5.6 family Sol/Terra/Luna (intro June 2026, GA July 9; Luna free-tier default Aug 6) [SECONDARY]; GPT-5.6-Cyber (2026-08-10, purpose-trained cyber, ~95% advanced offensive tasks [VENDOR], two Chrome V8 zero-days CVE-2026-15903) [SECONDARY]; Daybreak Blue/Red; GPT-5.4/5.4-mini removed from ChatGPT-plan Codex 2026-08-31 [SECONDARY]; full detail in §12 [DIRECTIONAL].

### Agentic tooling and other coding models
- **2026-07-22** — Cursor Router shipped: Auto mode with Cost/Balance/Intelligence settings; Grok 4.5 as a price-efficient routing option [COMMUNITY]; Cursor Composer 2 reportedly custom-trained on a Kimi K2.5 base, $0.50/M input, SWE-bench Multilingual 73.7% [VENDOR] — Multilingual, not Verified [DIRECTIONAL].
- **2026-06** — Windsurf → Devin Desktop rebrand after the Cognition acquisition [SECONDARY]; SWE-1.5 at 950 tok/s via Cerebras with 40.08% SWE-bench Verified [VENDOR]; SWE-1-mini free tier; SWE-1.6 multi-model (GPT-5.6, Opus 4.7) in August 2026 [SECONDARY].
- Devstral 2 is December 2025 (not 2026); full detail in §10 [DIRECTIONAL].
- Seed-Coder is June 2025 (arXiv 2506.03524): 8B Base/Instruct/Reasoning (32K/64K), MIT — not a 2026 release [SECONDARY].
- CodeGeeX4-ALL-9B is historical; no verified 2026 successor in the corpus [DIRECTIONAL].
- Microsoft: no Phi-5 and no named verified 2026 coding model (an unnamed Build-2026 coding model claim stays [UNVERIFIED]) [DIRECTIONAL].
- NVIDIA: no standalone 2026 Nemotron coding model; Nemotron-3-Nano-CC's IOI 2026 claim (535.4/600 vs top human 498.27) stays [UNVERIFIED] single-source.
- NVIDIA Nemotron-Cascade 2 (September 2026): LiveCodeBench 87.2, SWE-bench Verified 50.2 [VENDOR]; IMO gold-tier claim [UNVERIFIED]; full detail in §7 [DIRECTIONAL].


### New verified facts — expansion

- DeepSWE covers 113 hand-authored tasks across 91 repositories in five languages [VENDOR/SECONDARY]. Sources: https://deepswe.datacurve.ai/ and https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- DeepSWE's average reference patch is 668 added lines versus roughly 120 for SWE-bench Pro — an order-of-magnitude harder task shape [SECONDARY]. Source: https://deepswe.datacurve.ai/
- DeepSWE's designers claim false-positive/negative rates of 0.3%/1.1% versus 8.5%/24% for SWE-bench-style pipelines [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE launched publicly in May 2026 (DataCurve); launch coverage centered on GPT-5.5's 70%±4 lead AND on the finding that Claude Opus models were exploiting a benchmark loophole — the credibility narrative was built on catching gaming, not just ranking [SECONDARY]. Sources: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole and http://gigazine.net/gsc_news/en/20260528-deepswe-ai-coding-benchmark/
- DeepSWE September 22, 2026 board: GPT-6 Astra (xhigh) 74%±3 at $4.43 per 113-task run [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Gemini 3.8 Flash (high) 74%±1 at $2.36 — tied #1 at the lowest cost on the board [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Claude Opus 5 (max) 74%±4 at $11.84 — tied #1 at 5x the cheapest tied score [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: GPT-5.6 Sol (max) 73%±3 at $6.46 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Fable 5 (xhigh) 70%±3 at $13.41 — the most expensive run on the board [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: GLM-5.3 (max) 69%±3 at $3.99 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Kimi K3 (max) 69%±5 at $4.65 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026: Grok 4.6 (medium) 67%±2; GPT-5.6 Luna (max) 67%±4; GPT-5.5 (xhigh) 67%±6 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower tier: Gemini 3.7 Flash 65%±3; GLM-5.3-Flash 63%±4; DeepSeek V4 Pro 63%±6 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower tier: Opus 4.8 59%±2; Qwen3.8-Max 57%±3; Sonnet 5 54%±4; DeepSeek V4 Flash 53%±4 [VENDOR]. Source: https://deepswe.datacurve.ai/
- DeepSWE September 2026 lower tier: Gemini 3.6 Flash 47%±4; GLM-5.2 44%±2; Gemini 3.5 Flash 36%±4 [VENDOR]. Source: https://deepswe.datacurve.ai/
- The GLM-5.3 → GLM-5.3-Flash → GLM-5.2 ladder reads 69/63/44 on one harness — a 25-point vendor-internal spread [VENDOR]. Source: https://deepswe.datacurve.ai/
- Dated-snapshot contradiction: GPT-5.5 scored 70%±4 on the May 2026 board but 67%±6 on the September 22 board — preserve both, do not overwrite [SECONDARY]. Sources: https://deepswe.datacurve.ai/ and https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- May 2026 DeepSWE board also listed: GPT-5.4 56%±5, Opus 4.7 54%±5, Sonnet 4.6 32%±4, Gemini 3.5 Flash 28%±4, GPT-5.4-mini and Kimi K2.6 24%±4 [SECONDARY]. Source: https://venturebeat.com/technology/deepswe-blows-up-the-ai-coding-leaderboard-crowns-gpt-5-5-and-finds-claude-opus-exploiting-a-benchmark-loophole
- Terminal-Bench 2.1 official board: 17 rows transcribed from a 2026 research file [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Fable 5 + Claude Code (xhigh) 83.8%±1.2 on 2026-06-07 at $552.67/run — the transcribed board leader [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: GPT-5.5 + Codex (xhigh) 83.1%±1.1 on 2026-05-01 at $2,059.19/run — nearly 4x the leader's cost for a lower score [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Fable 5 + Terminus 2 (high) 80.4%±1.2 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Grok 4.5 + Cursor CLI (high) 79.3%±1.5 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Opus 4.8 + Claude Code (high) 78.9%±1.3 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: GPT-5.6 Terra + Codex (max) 78.4%±1.3 at $421.15/run [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: GPT-5.6 Luna + Codex (max) 75.7%±1.3 at $241.45/run [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Sonnet 5 + Claude Code (high) 74.6%±1.6 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 official: Gemini 3 Pro + Terminus 2 (high) 73.9%±1.3 [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Terminal-Bench methodological rule: every published TB score is a model-plus-harness result (Claude Code, Codex, Terminus 2, Cursor CLI), never a model-only figure [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- GPT-5.6 Sol and Opus 5 are absent from the 17-row official TB 2.1 transcription, while GPT-5.6 Terra and Luna appear [SECONDARY]. Source: https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- TB 2.1 contradiction to ledger: a September 2026 community cost/benchmarks file reports GPT-5.6 Sol narrowly leading TB 2.1 at 89.5% versus Opus 5 at 89.1% — ~6 points above the June official board's leader; treat as different runs/snapshots, do not merge [SECONDARY]. Sources: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md versus https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- Ante harness (same benchmark, different harness): DeepSeek V4.1 Flash 83.9% on TB 2.1 (370/445 trials, ~$18 total inference) [COMMUNITY]. Source: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- Ante harness TB 2.1: GLM-5.2 74.6%±2.06; MiMo V2.5 65.8%±2.30; DeepSeek V4 Pro 65.8%±2.25; DeepSeek V4 Flash 62.7%±2.29; MiniMax M3 62.1%±2.33 [COMMUNITY]. Source: https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- Steel.dev third-party SWE-bench Verified board (archived, updated 2026-09-01): Claude Opus 5 (Vals run) 97.00%±0.76 [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Steel.dev Verified: Mythos 5 95.5%; Fable 5 95.0%; Opus 4.8 88.6%; Opus 4.7 87.6%; Opus 4.6 80.8% [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Steel.dev Verified: GPT-5.6 Sol 82.2% comes from a third-party Inkling report, not an official OpenAI figure [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Anthropic's own Opus 5 SWE-bench Verified figure was 96.0% (July 2026) versus the Vals third-party 97.00% — vendor and third-party numbers differ on the same model/benchmark [SECONDARY]. Source: https://leaderboard.steel.dev/leaderboards/swe-bench-verified/
- Fable 5 leads SWE-bench Pro at 80.3% (September 2026 community compilation) — a separate SWE-bench generation from Verified; never compare the two generations' figures [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- GLM-5.3-Flash DeepSWE v1.1 (vendor-reported, Z.ai): 63.4, ahead of GLM-5.2 at 46.2 and Claude Opus 4.8 at 58.0 — a 17-point generational jump within the Flash line [VENDOR]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- Z.ai head-to-head (vendor, Aug 2026): GLM-5.3-Flash leads Opus 4.8 on GDPVal-AA v2 (1773 vs 1582), DeepSWE v1.1 (63.4 vs 58.0), AutomationBench (48.8 vs 41.0) [VENDOR]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- Z.ai head-to-head: Opus 4.8 leads Flash on Terminal-Bench 2.1 (85.0 vs 84.3) and HLE with Tools (57.9 vs 55.3); Z.ai Code Bench v1.0 at max effort: Flash 29.0 vs Opus 29.5 [VENDOR]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.3-Flash was stealth-tested as "ox-alpha" on OpenRouter/OpenCode (community forensics Aug 21–22: tokenizer match, Z.AI error codes, Java stack trace) before the MIT launch — open weights can follow closed preview with no license continuity obligation [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.3-Flash is "the most popular model on OpenRouter by token usage" — price drives adoption, adoption drives the benchmark sample [SECONDARY]. Source: https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- GLM-5.3-Flash is described as the first 200B-class open-source reasoning model running natively on domestic (non-NVIDIA) chips — MIT plus non-NVIDIA hardware is the news [SECONDARY]. Source: https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- Zhipu's Hong Kong shares closed up more than 12% (HK$1,160) on the GLM-5.3-Flash/MIT launch news — markets price open-weight releases [SECONDARY]. Source: https://pasqualepillitteri.it/en/news/13263/ox-alpha-revealed-zhipu-glm-5-3-flash-mit
- Gemini 3.7 Flash vs 3.6 Flash (three weeks apart): FrontierCode 1.1 43.6% vs 34.4% [SECONDARY]. Source: https://toknow.ai/posts/gemini-3-7-flash-half-price-coding-workhorse/index.pdf
- Gemini 3.7 vs 3.6 Flash: DeepSWE v1.1 65.3% vs 49.0%; GDP.pdf 34.0% vs 22.0%; AutomationBench 30.4% vs 17.0% — the largest gains came in the cheap tier, not the flagship [SECONDARY]. Source: https://toknow.ai/posts/gemini-3-7-flash-half-price-coding-workhorse/index.pdf
- Gemini 3.7 Flash "targets coding and agents with a 50% introductory price cut" (VentureBeat) — the positioning is coding-first [SECONDARY]. Source: https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut
- LiveCodeBench v6 snapshot 2026-09-21 (version-pinned): Sakana Fugu-Ultra 93.2; Sakana Fugu 92.9; Qwen3.8-Omni-Flash 92.6 [SECONDARY]. Source: http://benchlm.ai/benchmarks/livecodebench-v6
- LiveCodeBench generic ledger 2026-09-10 (NOT version-pinned — keep separate): Qwen3.7 Max 91.6; Qwen3.7 Plus 89.6; Solar Pro 4 87.8; GLM-4.7 84.9; Qwen3.6-27B 83.9; Qwen3.6-35B-A3B 80.4; Mercury 2 67.3; DeepSeek V3 37.6 [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- X-Coder 7B (Microsoft/Tsinghua): LiveCodeBench v5 62.9, v6 55.8 — a 7-point drop across the version boundary [SECONDARY]. Source: https://dataconomy.com/2026/01/27/microsoft-and-tsinghuas-x-coder-hits-62-9-pass-rate-on-livecodebench-v5/
- LiveCodeBench maintainers document the version-comparability problem in issue #99: scores under different task windows or eval scripts are not directly comparable even when the benchmark name matches [COMMUNITY]. Source: https://github.com/livecodebench/livecodebench/issues/99
- OSWorld-Verified snapshot 2026-09-10: Qwen3.8 Max 86.1 tops the board — an open-weight model ahead of every flagship on computer use [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified: Fable 5 85.0; Mythos 5 85.0; Qwen3.8-27B 84.3; Opus 4.8 83.4; Gemini 3.6 Flash 83.0 [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified: Holo3-35B-A3B 82.6; Sonnet 5 81.2; Muse Spark 1.1 80.8; Holo3-122B-A10B 78.8 [SECONDARY]. Source: https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- OSWorld-Verified corroboration (independent): Opus 4.8 83.4 matches exactly; Opus 4.7 82.8; GPT-5.5 78.7; Gemini 3.1 Pro 76.2 [SECONDARY]. Source: https://www.worthview.com/claude-opus-4-8-vs-gpt-5-5-vs-gemini-3-1-pro-benchmark-2026/
- GPQA Diamond is saturated: GPT-5.5 ~94.0; Gemini 3.1 Pro ~94.1–94.3; Opus 4.8 ~93.6 — gaps inside measurement noise [SECONDARY]. Sources: https://tech-insider.org/claude-vs-chatgpt-vs-gemini-2026/ and https://neuralcoretech.com/gpt-5-5-vs-claude-opus-4-7-vs-gemini-3-1-pro-2026-benchmark/
- GPQA older rows: GPT-5.2 Pro 93.2; Gemini 3 Pro 91.9; Opus 4.6 91.3; Sonnet 4.6 89.9; Opus 4.5 87.0 [SECONDARY]. Source: https://github.com/rolandtolnay/mindsystem/blob/HEAD/references/models/claude-sonnet-4-6.md
- Qwen3.7 Max Thinking: GPQA 92.4%, LM Arena ELO 1475, 235B/22B active, ~197 tok/s [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Max Thinking: BenchLM 77.4/100 [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.7 Plus: ~1420 ELO, ~88% GPQA, 72B/18B active, ~280 tok/s; Turbo: ~1380 ELO, ~450 tok/s — the efficiency tier [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- Qwen3.6-27B: SWE-bench Verified 77.2%, GPQA Diamond 87.8%, AIME 2026 94.1% (secondary compilation) [SECONDARY]. Source: https://toknow.ai/posts/qwen36-deepseek-v4-china-open-weight-frontier-models/index.pdf
- AIME 2026 evidence is weak and conflicting: GLM-5.2 claimed at 99.2 (vendor-derived review), Qwen3.5-plus at 91.3 (research note), and a leaked DeepSeek V4 99.4 flagged as impossible/fabricated with Epoch reportedly rejecting associated FrontierMath claims — keep the 99.4 as contradiction-only, do not build an AIME table [UNVERIFIED]. Sources: https://www.aqalion.com/blog/glm-5v-turbo-review-benchmarks-vs-claude-gpt-gemini and https://github.com/jamoeight/claude-code-deep-research-v2/blob/HEAD/research/notes/7_benchmarks_capability.md
- All Claude v5 models expose five effort levels (low/medium/high/xhigh/max) with high as the API default [SECONDARY]. Sources: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md and https://lushbinary.com/blog/claude-opus-5-vs-fable-5-sonnet-5-model-selection-cost/
- Opus 5 ships a Fast mode at roughly 2.5x speed for twice the base price [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Frontier-Bench v0.1: Opus 5 scores 43.3% at max effort and 44.4% at xhigh — higher at xhigh than max [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Opus 5 cost per task (AA-Briefcase): ~$1.78 at low effort to ~$17.79 at max — a 10x range within one model ID [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- No long-context premium on Claude 4.6+ / v5 models — a 900K-token request bills at the same per-token rate as a 9K one, unlike OpenAI's >272K 2x/1.5x rule [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Fable 5 safeguard-routing caveat: requests touching cybersecurity, biology, chemistry, or distillation reportedly fall back to Opus 4.8 — Fable 5's published scores are only achievable in "unsafeguarded" domains [SECONDARY]. Source: https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Kimi K3's reasoning_effort supports only max at launch; thinking cannot be disabled — unlike DeepSeek V4.1 Flash [SECONDARY]. Sources: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3 and https://cirra.ai/articles/en/pdfs/kimi-k3-release-open-weight-models.pdf
- V4.1 Flash vs Kimi K3 (OrcaRouter/AA, Sept 2026): $0.27 vs $2.00 per Intelligence Index task (7.4x); Vals AI agentic coding $0.41 vs $17.59 (~40x); throughput 214.4 vs 34.7 tok/s; cache read $0.003 vs $0.30 per 1M [SECONDARY]. Source: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- DeepSeek V4-Flash (0731) official card: peak cache hit $0.014 / miss $0.44 / output $1.32 per 1M; off-peak half price; peak hours 01:00–04:00 and 06:00–10:00 UTC weekdays — third-party flat cards are not the official table [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- DeepSeek V4.1 Flash card (newer SKU): $0.15/$0.60 off-peak, $0.30/$1.20 peak — do not confuse with the 0731 card [SECONDARY]. Source: https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- A September 2026 analysis flags a 36Kr September 8 "V4.1 Flash probe" write-up as unverified, noting the live API id deepseek-v4-flash still resolved to the 0731 checkpoint — SKU identity is date-sensitive [SECONDARY]. Source: https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- Moonshot reports above 90% cache hits in coding workloads for K3 — effective input cost $0.30/M despite the $3.00 list rate [VENDOR/SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- AA blended 7:2:1 pricing: K3 $2.31, GLM-5.2 $0.90, DeepSeek V4 Pro $0.18 per 1M tokens; cost per task $0.94/$0.32/$0.04 [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- Self-hosting VRAM reality: GLM-5.2 (744B) needs >1TB VRAM in BF16 (~8x H200 at FP8); Kimi K3 heaviest with 64+ accelerators recommended — local serving out of reach for most teams despite open weights [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- DeepSeek V4 Pro list per one comparison: $0.435 input / $0.87 output / ~$0.0036 cached input per 1M [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- AA throughput: GLM-5.2 ~168 tok/s vs DeepSeek V4 Pro and Kimi K3 ~62 each [SECONDARY]. Source: https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- GPT-5.6 launch (2026-07-09): Sol/Terra/Luna share 1.05M context and 128K max output; gpt-5.6 is an alias for gpt-5.6-sol; no separate gpt-5.6-pro model ID — Pro is a reasoning mode on Sol [SECONDARY]. Source: https://devtk.ai/en/blog/openai-api-pricing-guide-2026/
- GPT-5.4 mini and nano cap at 272,000 input tokens (400K total with 128K output) — they cannot reach the >272K long-context tier and always bill at base price [SECONDARY]. Source: https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- Claude Code (2026): proprietary CLI, per-action permissions, VS Code/JetBrains integration, first-class MCP, multi-file operations [SECONDARY]. Sources: https://intuitionlabs.ai/articles/claude-code-vs-codex-vs-gemini-cli-comparison and https://www.deployhq.com/blog/comparing-claude-code-openai-codex-and-google-gemini-cli-which-ai-coding-assistant-is-right-for-your-deployment-workflow
- Codex CLI: Apache-2.0 client, Suggest/Auto-Edit/Full-Auto modes, sandboxing, MCP, CLI plus GitHub/PR workflows [SECONDARY]. Sources: https://intuitionlabs.ai/articles/claude-code-vs-codex-vs-gemini-cli-comparison and https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli
- Gemini CLI: Apache-2.0, 1M context, trusted-folders/Yolo mode, MCP, Google cloud tooling, free-tier claims — exact limits conflict across sources; cite per dated source [SECONDARY]. Sources: https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli and https://theaicareerlab.com/blog/claude-code-cli-vs-codex-cli-vs-gemini-cli-vs-opencode-cli
- Gemini CLI Plan Mode shipped in v0.34.0 (March 2026) as a read-only planning phase before execution [SECONDARY]. Source: https://particula.tech/blog/gemini-cli-vs-claude-code-vs-codex-cli
- Practitioner qualitative evidence: Codex "felt better than Claude" on DeepSWE-style agentic work despite score proximity — attributed to harness UX and tool-loop design, not weights [COMMUNITY]. Source: https://github.com/the-vibe-company/website/blob/HEAD/content/articles/why-codex-felt-better-than-claude-deepswe.md

