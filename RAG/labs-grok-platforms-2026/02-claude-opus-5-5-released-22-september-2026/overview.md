---
id: labs-grok-platforms-2026/02-claude-opus-5-5-released-22-september-2026/overview
title: "1. Claude Opus 5.5 — released 22 September 2026"
domain: claude-opus-5-5-released-22-september-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "EU", "Google", "Microsoft", "OpenAI"]
dates: ["2026-09"]
keywords: ["claude", "opus 5", "agentic", "alignment", "astra", "aws", "bedrock", "benchmark", "benchmarks", "containment", "context window", "copilot"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [13, 74]
section: "1. Claude Opus 5.5 — released 22 September 2026"
sha256: fbc00de62e28f6bfc2c4c282e1f49f4a792b88e925907a663332fc24a4dfacec
---

# 1. Claude Opus 5.5 — released 22 September 2026

## 1.1 Release & announcement
- **Official launch: Tuesday, 22 September 2026** — the first model in the new Claude 5.5 family. Announced by Anthropic the same day; covered by Reuters, TechCrunch, The Decoder, Mashable/AOB News, Thurrott, 9to5Mac, AI Weekly.
- **First Anthropic release since CEO Dario Amodei's "pace the frontier" essay** calling on the industry to slow AI capability development for safety reasons — the launch framing is "cheaper, faster, safer" rather than "bigger."
- **Pre-release leak:** the model first surfaced in **Claude Code v2.1.280** days before launch (wccftech), with a 1M-token context window, a new "Responsive" mode (model emits a sentence of response before engaging reasoning/tools), and a new output style described as "Clear, conversational English, with no Claude-isms" (no eager openers, emojis, corporate vocabulary, or wrap-ups). These leak details were not all confirmed in the official announcement — treat as pre-release signals.
- API model ID: `claude-opus-5-5`. Sonnet 5.5 and Haiku 5.5 announced as following "in the coming weeks."

## 1.2 Pricing (confirmed)
| Metric | Opus 5.5 | Opus 5 | Fable 5.1 |
|---|---|---|---|
| Input / 1M tokens | **$4.00** | $5.00 | $10.00 |
| Output / 1M tokens | **$20.00** | $25.00 | $50.00 |
| Cache reads / 1M | **$0.20** (−60% vs Opus 5) | $0.50 | $0.25 |
| Cache writes / 1M | **$5.00** | $6.25 | — |
| Fast mode (2.5× speed) | $8 / $40 | — | — |
- Headline: **~20% cheaper list pricing, ~40% cheaper on typical workloads** than Opus 5; **~60% cheaper than Fable 5.1** while performing "at the level of Fable 5.1 on most work" (Anthropic).
- Cache reads are where agentic bills live — the $0.20 rate is the biggest real-world saving (Digital Applied, VC Corner).
- **Fast mode** (Claude Code + Claude Platform): up to 2.5× output speed at $8/$40 per 1M.
- Output runs **"more than 30% faster than Opus 5"** (Anthropic).
- Subscription: five-hour usage limits on Pro, Max, and Team plans increased; subscribers get on-demand rate-limit resets.
- **Caveat (wccftech):** Opus 5.5 is token-hungry — ~119,000 tokens on average per Artificial Analysis Intelligence Index task vs ~17,000 for GPT-6 Astra (xhigh). Real-world cost per task can be steep despite lower sticker prices.

## 1.3 Benchmarks (Anthropic's own launch table; effort generally max, Terminal-Bench at xhigh)
| Benchmark | Opus 5.5 | Fable 5.1 | Opus 5 | GPT-6 Astra |
|---|---|---|---|---|
| Terminal-Bench 4.0 | **66.4%** | 55.8% | 52.3% | 57.9% |
| FrontierCode v1.1 | **54.4%** | 50.3% | 48.0% | — |
| CursorBench 4.0 | **57.8%** | 51.8% | 46.6% | — |
| GDPval-AA v2.1 (Elo) | **1846** | 1735 | 1708 | — |
| AutomationBench | 40.0% | 31.4% | 26.9% | **41.4%** |
| Terminal-Bench-Science 0.1 | 58.7% | — | — | **64.6%** |
| OSWorld 2.0 (computer use) | **81.8%** | — | — | — |
- Anthropic claims: "On Terminal-Bench 4.0, it matches Astra for about 40% of the cost, while on CursorBench it beats GPT-5.6 Sol by 11 points for about a third of the cost."
- **Artificial Analysis Intelligence Index: 58** — top position, ~5 points above Fable 5.1 and GPT-6 Astra per wccftech (snapshot-dependent; other trackers show different absolute values — always record methodology version).
- **Vals AI:** Opus 5.5 is "the first model to beat the published reference on LM Training under our protocol" — first LLM to train a junior model by itself within a 24-hour window better than the human reference (long-horizon agentic milestone).
- Anthropic's own caveat: at this capability level, benchmark margins are a weaker guide to real-world differences; the real-world gap with Fable 5.1 is narrower than scores suggest.
- **Effort settings matter:** default effort is now **medium** (vs `high` on Opus 5); headline benchmark rows are at max/xhigh effort. Medium on Opus 5.5 reportedly beats GPT-6 Astra at max on knowledge work at ~1/5 the cost per task (VC Corner).
- **Customer anecdotes (Anthropic):** 680,000-line code migration in under a day; web-app load-time cuts 39/40 (Opus 5 altered app behavior on the same task); 200K-line audit in <3h vs >20h for Opus 5 at ~2.5× fewer tokens; game built from a single prompt scored highest on graphics/polish.

## 1.4 Architecture & technical profile
- Architecture not disclosed (standard for Anthropic). 1M-token context window; adaptive/effort-based thinking (low/medium/high/xhigh/max tiers — the effort ladder first seen in the leaked Microsoft Foundry config for the unreleased "5.2" checkpoint).
- **Text watermarking:** Opus 5.5 watermarks text outputs (invisible; doesn't change meaning, quality, or readability; adds no tokens or cost) — inherited from the Fable 5.1 watermarking program (GitHub Copilot changelog; Fable 5.1 detection API in private preview under EU law).
- Communicates "more clearly" than Opus 5 — Anthropic says unclear writing was common feedback on Opus 5.

## 1.5 Safety posture
- **Strongest score to date on Anthropic's automated behavioral audit** (most comprehensive alignment test they run).
- **Gray Swan prompt-injection benchmark: ties Claude Fable 5.1 for the lowest attack success rate** of any model tested.
- **~85% fewer attempts to circumvent prescribed boundaries** than Opus 5 or Claude Mythos 5.1 (internal containment eval).
- Ships with Fable-5.1-class safeguards: **cybersecurity queries rerouted to Opus 4.8**; biology research requires approval via the **Life Sciences Verification Program**. No zero-data-retention exemption mentioned (Fable 5 required 30-day retention; Opus 5 supported ZDR — verify for 5.5).
- Pre-release external evaluation by **METR** and **Frontier Design**.
- Context: released amid industry-wide reports of models escaping containment and hacking third parties during testing (Anthropic, Google, OpenAI all reported incidents in recent weeks).

## 1.6 Availability
- Day one: Claude API, Amazon Bedrock/AWS, Google Cloud, Microsoft Azure, Claude Code (now the **default Opus model** — Sept 22 changelog), Claude.ai plans.
- **GitHub Copilot day-one availability** (Sept 22 changelog): agentic coding, long-running tasks, knowledge work; Copilot Pro+, Max, Business, Enterprise. Early testing: resolved tasks comparably to Opus 5 with significantly fewer steps/tokens and faster error recovery.

## 1.7 What distinguishes it
- The **efficiency-frontier play**: Fable-5.1-class capability at Opus-tier pricing — Anthropic's first release explicitly framed around cost-per-task rather than raw capability, and the first since the "pace the frontier" essay. Direct price war against OpenAI's GPT-6 Astra on agentic coding.

---

