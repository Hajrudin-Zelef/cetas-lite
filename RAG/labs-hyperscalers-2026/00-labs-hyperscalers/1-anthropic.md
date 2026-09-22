---
id: labs-hyperscalers-2026/00-labs-hyperscalers/1-anthropic
title: "§1 — ANTHROPIC"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Microsoft", "OpenAI", "Stripe", "United States"]
dates: ["2025-05", "2026-02", "2026-02-05", "2026-04", "2026-04-16", "2026-05-28", "2026-06-09", "2026-06-30", "2026-07-24", "2026-07-26", "2026-07-31", "2026-08-10", "2026-09-01", "2026-09-10", "2026-09-22"]
keywords: ["accelerator", "agent", "agentic", "agents", "agi", "aws", "bedrock", "benchmark", "benchmarks", "claude", "compute", "context window"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [58, 131]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: caef6414caf4036b5acca67a2a6a653f78d0035884288c50c5fe413ff0a9021b
---

# §1 — ANTHROPIC

## §1 — ANTHROPIC

> **2026 at a glance — Anthropic:** five funding events ($30B G → $65B H → $52B I @ $1.26T, plus a $30B credit facility) made Anthropic the most valuable private company ever; Claude Opus 4.6 / Sonnet 4.6 / Fable 5 shipped alongside the Cowork agentic OS; $25B annualized revenue run rate (80% enterprise); confidential S-1 filed Jun 1 targeting a Nov 2026 NYSE IPO; sovereign compute deals in the US ($200B/10 GW) and UK (£31B/4 GW Stargate); export controls hit Fable/Mythos in Macau (Jul 9).

### 1.1 Model releases (Feb → Sep 2026)

#### Claude Opus 4.6 — February 5, 2026
- Launched with stronger coding, planning, code review, debugging, long-running agent reliability, and 1M-token context window in beta. [secondary] https://github.com/jqueryscript/anthropic-claude-timeline
- Companion benchmark snapshot (from a 2026 tracking table): SWE-bench Verified 71.8%, GPQA Diamond 79.2%, 32,768 max thinking budget. [secondary] https://www.narenvadapalli.com/blog/anthropic-claude-opus-5-architectural-guide/
- Launch timed within minutes of OpenAI's GPT-5.3-Codex launch (Feb 5, 2026). [secondary] https://github.com/xagi-labs/xagi-labs.github.io/blob/HEAD/content/blog/openai-and-anthropic-go-to-war-claude-opus-46-vs-gpt-53-codex.md

#### Claude Sonnet 4.6 — February 2026
- Broad upgrade across coding, computer use, long-context reasoning, agent planning, knowledge work, and design. [secondary] https://github.com/jqueryscript/anthropic-claude-timeline
- Anthropic-reported agentic coding score: 58.1% (vs 63.2% for Sonnet 5). [vendor-reported via press] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/
- Lane note: Sonnet 4 (May 2025) → Sonnet 4.5 (Sep 2025) → Sonnet 4.6 (Feb 2026) → Sonnet 5 (Jun 2026); no Sonnet 4.7/4.8 exists. [secondary] https://hidekazu-konishi.com/entry/anthropic_claude_model_release_timeline.html

#### Claude Opus 4.7 — April 16, 2026
- Positioned as upgrade for advanced software engineering, complex workflows, high-resolution vision, and professional work ("vision synthesis & multi-step professional knowledge work"). [secondary] https://www.scriptbyai.com/anthropic-claude-timeline/
- Snapshot benchmarks: SWE-bench Verified 72.9%, GPQA Diamond 80.4%. [secondary] https://www.narenvadapalli.com/blog/anthropic-claude-opus-5-architectural-guide/
- Anthropic–Google multi-GW TPU deal announced the week of Opus 4.7 (Apr 6, 2026). [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/31-google-anthropic-40b-commitment.md

#### Mythos-class tier announced — April 2026
- Anthropic announced "Mythos" in April 2026, a new tier above Opus. [secondary, via Wikipedia mirror] https://pengen.diewe.workers.dev/thegdsks/openclaw-https-en.wikipedia.org/wiki/Anthropic
- Mythos 5 Preview released ~April 2026 (scores 69% on ExploitBench vs 78% for Mythos 5 at Fable launch). [independent — MLQ News] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/

#### Claude Opus 4.8 — May 28, 2026
- Focus: consistency and autonomy for long-running agent tasks; 65,536 max thinking budget; SWE-bench Verified 73.6%, GPQA Diamond 81.1%. [secondary] https://www.narenvadapalli.com/blog/anthropic-claude-opus-5-architectural-guide/
- Anthropic-reported agentic coding 69.2% (vs Sonnet 5 63.2%). [vendor-reported via press] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/
- API price: $5/M input / $25/M output; offered "Fast mode" (later restricted). [independent] https://www.macrumors.com/2026/07/24/anthropic-opus-5/
- Retired model note: an Aug-2026 API changelog notes "Fast mode has been removed for Claude Opus 4.6 and 4.7; migrate to Opus 4.8 or Opus 5 to keep using it." [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- BenchLM HLE (as of 2026-09-10): Opus 4.8 at 57.9%. [secondary] https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md
- Series H ($65B @ $965B) closed the same day, May 28, 2026. [independent] https://datafloq.com/openai-set-the-ai-valuation-record-in-march-anthropic-broke-it-by-may/

#### Claude Sonnet 5 — June 30, 2026
- Replaced Sonnet 4.6; became the default model for Free and Pro subscription tiers, Claude Code, and API default selection. [independent] https://www.techtimes.com/articles/319409/20260701/claude-sonnet-5-ships-anthropic-default-agentic-performance-closes-opus-gap.htm
- API ids/context: `claude-sonnet-5`, 1M context, 128K max output. [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Pricing: introductory $2/M input / $10/M output through Aug 31, 2026 (then list price $2 in / $10 out per later price tables — no step-up confirmed). [independent] https://memeburn.com/claude-sonnet-5-launched-with-near-opus-performance-at-half-the-price/ ; [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained
- Anthropic-reported benchmarks (launch): agentic coding 63.2% (vs Opus 4.8's 69.2%), SWE-bench Verified 85.2% (vendor-reported), SWE-bench Pro 63.2%, Terminal-Bench 2.1 80.4%, OSWorld Verified 81.2%, GDPval-AA ~1607–1609 Elo, Artificial Analysis Intelligence Index ~53. [vendor-reported] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/ and [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Vals AI independent: Terminal-Bench 2.1 74.53%. [secondary] https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- BenchLM snapshots: HLE 57.4% (rank #7 as of 2026-09-10); BenchLM "BenchAlign" composite ~65, rank #29–33 (July 31, 2026); coding #9 / agentic #15; BrowseComp 84.7%. [secondary] https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md ; https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- Safety: Anthropic claims lower rate of undesirable behaviors vs Sonnet 4.6, "safer to use in agentic contexts" (hallucination reduction claimed). [vendor-reported via press] https://www.neowin.net/news/anthropic-releases-claude-sonnet-5-with-improved-agentic-capabilities/
- API behavioral changes (Aug 2026 changelog): adaptive thinking on by default; manual extended-thinking budgets and non-default sampling parameters now return errors; new tokenizer producing ~30% more tokens for the same text. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md

#### Claude Fable 5 + Claude Mythos 5 — June 9, 2026
- Fable 5: first publicly available "Mythos-class" model, tier above Opus; Mythos 5 is the same underlying weights with cyber safeguards lifted, restricted to Project Glasswing partners (cyberdefense, critical-infrastructure orgs; ~150 organizations in 15+ countries, a US-government collaboration) and limited US government channels. [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/ ; [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md
- Availability channels: Claude API, Amazon Bedrock, Google Cloud, Microsoft Foundry. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/67-anthropic-claude-fable-5-release.md
- Pricing: $10/M input / $50/M output for both; 90% prompt-cache input discount; most expensive GA Anthropic model (~2× Opus output). In Claude Code, Fable selectable via `/model fable` or `best` alias, requires v2.1.170+; 1M context on API; not available under zero-data-retention; Anthropic requires 30-day data retention on all Mythos-class traffic. [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md ; [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Subscription availability: free on Pro/Max/Team/seat Enterprise June 9 → June 22 only; from June 23 requires usage credits; Anthropic said it intended to restore as standard subscription model "as quickly as we can," no date given. [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md
- Anthropic-reported benchmarks (launch): SWE-bench Verified 95.0% (Vals AI independent: 95.0%), SWE-bench Pro 80.3% (leading at launch; contested — produced on Anthropic scaffolding), GPQA Diamond 92.6%, Terminal-Bench 2.1 88.0% (Vals AI independent: 80.52%), τ²-Bench 98.5%, FrontierCode 29.3% (vs 13.4% Opus 4.8, 5.7% GPT-5.5), ExploitBench (Mythos 5) 78%. Artificial Analysis Intelligence Index: #1 across public models, quality score 64.9 at launch. [vendor-reported/secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/67-anthropic-claude-fable-5-release.md ; https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Safety architecture: classifiers route cyber/bio/distillation flagged requests to Opus 4.8 fallback; triggers in <5% of sessions (Anthropic claim); 319-page system card. [vendor-reported via secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/67-anthropic-claude-fable-5-release.md
- **Covert-degradation controversy (June 10–11, 2026):** Simon Willison published "If Claude Fable Stops Helping You, You'll Never Know" documenting that for "frontier LLM development" queries (pretraining pipelines, distributed training infra, accelerator design), the system card described INVISIBLE safeguards — prompt modification, steering vectors, parameter-efficient fine-tuning — with no user notification. On June 11 Anthropic reversed the policy: "we made the wrong tradeoff," committing to make all capability restrictions visible and explained. [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/68-anthropic-fable-covert-capability-degradation.md
- Adoption datapoint: Stripe used Fable 5 to migrate a 50M-line Ruby codebase in one day (task estimated at two months for a team). [independent] https://mlq.ai/news/anthropic-ships-claude-fable-5-to-the-public-keeps-mythos-5-gated-for-cyberdefense/
- Suspension (see §1.5) June 12 → lifted ~June 30. [independent] https://www.rappler.com/technology/united-states-lifts-curbs-anthropic-fable-mythos-ai-models/
- **Fable 5.1 released September 1, 2026:** `claude-fable-5-1`, 1M context, $10/$50 per Mtok with $0.25/Mtok cache reads; became default Fable model in Claude Code v2.1.257. BenchLM HLE 65.0% (rank #1 as of 2026-09-10); ARC-AGI-2 90.0%. [secondary] https://github.com/florianbruniaux/claude-code-ultimate-guide/blob/HEAD/guide/core/claude-code-releases.md ; https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md

#### Claude Opus 5 — July 24, 2026
- Replaces Opus 4.8 at same API price ($5/M input / $25/M output); new default on Claude Max, strongest model on Claude Pro; available on API, AWS, Google Cloud, Microsoft Foundry. Fast mode available at 2.5× speed for 2× base price. [independent] https://www.macrumors.com/2026/07/24/anthropic-opus-5/
- US-only inference option at 1.1× standard price; prompt caching cuts input costs up to 90%; batch 50% savings. [secondary] https://metrotechs.io/news/anthropic-claude-opus-5-release-enterprise-evaluation?v=2513abcab895229909902a0af92abafb
- Anthropic-reported benchmarks: Frontier-Bench v0.1 43.3% max effort (44.4% xhigh; vs Opus 4.8 18.7%, Fable 5 33.7%, GPT-5.6 Sol 37.5% — SOTA), CursorBench 3.2 within 0.5% of Fable 5 peak at half cost, ARC-AGI 3 3× next-best model, Zapier AutomationBench ~1.5× next-best pass rate, OSWorld 2.0 surpasses Fable 5 best at ~1/3 cost, ARC-AGI-2 90.4%, Terminal-Bench 2.1 89.1% max, SWE-bench Verified 96% (vendor) / 97.0% (Vals AI independent), SWE-bench Pro 79.2%, BrowseComp 90.8%, FrontierCode 1.1 Main 53.4%, GDPval-AA ~1861 Elo. Artificial Analysis Intelligence Index ~61, narrowly #1 (Opus 5 holds #1 spot Jul–Sep). [vendor-reported/secondary/independent] https://pulse2.com/anthropic-launches-claude-opus-5/ ; https://www.neowin.net/news/anthropic-launches-claude-opus-5-with-near-fable-intelligence-at-half-the-price/ ; https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- ⚠️ **CONFLICT:** a second tracking source shows Opus 5 SWE-bench Verified at 74.8% "SOTA" in an Opus architecture table — vs 96% (vendor) / 97.0% (Vals AI). Likely different harnesses/effort configs; marked conflicting/uncertain, not reconciled. [secondary] https://www.narenvadapalli.com/blog/anthropic-claude-opus-5-architectural-guide/
- Effort toggle: users choose low/medium/high/max effort; automatic model fallbacks for API. [secondary] https://github.com/abdulhadi446/blog-s/blob/HEAD/blogs/ai-daily-roundup-2026-07-26/blog.md
- BenchLM (as of 2026-09-10): HLE 64.7% (rank #2); BrowseComp 90.8% (rank #4); ARC-AGI-2 90.4% (rank #3). BenchLM composite "BenchAlign" (Jul 31): 82.79, #3 agentic / #4 coding / #1 knowledge. [secondary] https://github.com/leoncuhk/awesome-llm-bench/blob/HEAD/README.md ; https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md

#### Claude Opus 5.5 — September 22, 2026 (today)
- Released Sept 22, 2026; Anthropic says it sets new SOTA in coding and knowledge work, outpaces Fable 5 in many benchmarks; first model released after CEO Dario Amodei embraced "pacing the frontier." Pricing cut: output $20/M (vs $25 predecessor); "other metrics have similar price drops"; faster to run (less compute per token). Sonnet 5.5 and Haiku 5.5 to follow "in the coming weeks." [independent] https://techcrunch.com/2026/09/22/anthropic-releases-opus-5-5-with-lower-prices-and-fable-level-performance/
- Safeguards: Opus 5.5 deemed comparable to Mythos on biology/cybersecurity capabilities; released under same safeguards as Fable (limits on exploit discovery, bioweapon-adjacent work). [independent] https://techcrunch.com/2026/09/22/anthropic-releases-opus-5-5-with-lower-prices-and-fable-level-performance/

#### Haiku line — no new Haiku in Feb–Sep 2026
- Haiku 4.5 (Oct 2025) remains the latest GA Haiku: $1/M in / $5/M out, 200K context, 64K max output, fastest in Anthropic's measurements (~101 output tok/s). No Haiku 4.6/4.7/4.8 or Haiku 5 exist as of July–Sept 2026; Haiku 5.5 promised "in the coming weeks" (Sept 22). [secondary] https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained ; https://hidekazu-konishi.com/entry/anthropic_claude_model_release_timeline.html ; [independent] https://techcrunch.com/2026/09/22/anthropic-releases-opus-5-5-with-lower-prices-and-fable-level-performance/

#### Retirements (API changelog, Aug 2026)
- Claude Opus 4.1, Sonnet 4, and Opus 4 retired — now return errors on all requests. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- API changelog (Aug 10, 2026): thinking changes, Managed Agents, consolidated rate-limit tiers; Opus 4.1/Sonnet 4/Opus 4 retired. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md

