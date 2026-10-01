---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-4
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "Nvidia", "OpenAI", "United States"]
dates: ["2026-02-12", "2026-04", "2026-07-24", "2026-07-26", "2026-08-10", "2026-09-10", "2026-09-22"]
keywords: ["agentic", "agents", "agi", "aws", "benchmarks", "claude", "compute", "cost", "cybersecurity", "exploit", "fable 5", "foundry"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [113, 146]
sha256: 877010217851d21df37d2349dbd926bc8e521a222c7c4ec0af0be0fbe6ec07c9
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

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

### 1.2 Funding, valuation, revenue (Feb → Sep 2026)

#### Series G — February 12, 2026: $30B at $380B valuation
- Announced Feb 12, 2026; co-led by D. E. Shaw Ventures, ICONIQ, and MGX; included portion of previously announced Microsoft and Nvidia investments. [independent] https://www.thehindu.com/sci-tech/anthropic-clinches-380-billion-valuation-after-30-billion-funding-round/article70627131.ece
- At announcement: run-rate revenue $14B; Claude Code alone $2.5B+ run rate (more than doubled since start of 2026); Claude Code business subscriptions quadrupled since start of year; enterprise = >50% of Claude Code revenue. [vendor-reported via press] https://www.thehindu.com/sci-tech/anthropic-clinches-380-billion-valuation-after-30-billion-funding-round/article70627131.ece
- Company had raised $13B Series F at $183B in early September (2025, per article context). [independent] https://www.thehindu.com/sci-tech/anthropic-clinches-380-billion-valuation-after-30-billion-funding-round/article70627131.ece

#### Amazon commitment — April 2026: up to $25B (total up to $33B)
- Amazon announced additional commitment up to $25B ($5B immediate) on top of earlier $8B; total Amazon commitment up to $33B — largest investor by committed capital. [secondary] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/openai_anthropic_funding_history.md
- Circular-deals reporting (Sep 2026): Amazon has $13B invested in Anthropic with up to $20B more; Anthropic has $100B+ AWS commitment over 10 years — ⚠️ [secondary][unverified], only partially corroborated. [secondary] https://www.aistockwire.com/

#### Google commitment — April 2026: up to $40B
- Google committed up to $40B ($10B immediate), dwarfing prior ~$2–3B; alongside Alphabet/Google as existing investor. [secondary] https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/openai_anthropic_funding_history.md ; https://siliconvalleyinvestclub.com/companies/anthropic/tear-sheet.pdf
- Valuation reported as $350B (Bloomberg) vs $380B (some outlets); contingent-$30B triggers undisclosed; no official joint announcement. See §3.6. [independent/secondary]

