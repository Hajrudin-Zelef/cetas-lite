---
id: labs-grok-platforms-2026/12-part-b-openai/overview
title: "PART B — OPENAI"
domain: part-b-openai
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "California", "DeepSeek", "Google", "Microsoft", "OpenAI", "SpaceX", "United States", "xAI"]
dates: ["2026-07", "2026-09", "2026-09-03", "2026-09-12", "2026-09-18"]
keywords: ["agentic", "agi", "antitrust", "astra", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt", "claude", "cost", "cyber"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [326, 391]
section: "PART B — OPENAI"
sha256: af59917cc55130b9e6634450219babf121e9b7fabd6145ee090ed618b2de7778
---

# PART B — OPENAI

## 5. GPT-5.x series progression through 2026

| Model | Release | Key facts |
|---|---|---|
| GPT-5 | Aug 2025 | Consolidated chat experience; base of the 5.x generation |
| GPT-5.3-Codex | Feb 5, 2026 | Agentic coding flagship; Terminal-Bench 2.0 77.3%; LTS through Feb 4, 2027 |
| GPT-5.4 | Spring 2026 | Codeforces ~3168 (ref. point in DeepSeek V4 comparisons) |
| GPT-5.5 | ~May–Jun 2026 | Prior flagship; remained supported after 5.6 launched |
| **GPT-5.6 (Sol / Terra / Luna)** | **Limited preview Jun 26, 2026 → GA Jul 9, 2026** | First frontier model through full US-government pre-release review |
| GPT-5.6-Cyber | Aug 10, 2026 | With Daybreak Blue/Red; solved 95% of cyber problems |

### GPT-5.6 — the tier portfolio
| Tier | API ID | Positioning | Price/1M in/out |
|---|---|---|---|
| **Sol** | `gpt-5.6-sol` | Flagship; "best coding model yet"; complex reasoning, agentic, science, cyber | **$5 / $30** |
| **Terra** | `gpt-5.6-terra` | Balanced; ~GPT-5.5-class at half the cost | $2.50 / $15 |
| **Luna** | `gpt-5.6-luna` | Fastest/cheapest; high-volume routine work | $1 / $6 |
- Specs shared across tiers: **~1.05M-token context**, 128K max output, training cutoff **Feb 16, 2026**; text+image in, text out; function calling, structured outputs, Batch API, native tools; fine-tuning not at launch; bare `gpt-5.6` alias → Sol.
- Sam Altman: Sol is **54% more token-efficient** on coding tasks; pitch shifted from raw intelligence to **intelligence per token per dollar**.

## 6. GPT-6 "Astra" — September 3, 2026

### 6.1 Release
- **Released Thursday, September 3, 2026** — two days after Anthropic's Claude Fable 5.1 (Sept 1), two months after GPT-5.6 Sol.
- Staged rollout: **Daybreak enterprise/cybersecurity program first** (Sept 3), then ChatGPT Plus/Pro/Business/Enterprise (Sept 4); API, AWS Bedrock, Microsoft Azure; **free tier and cheapest paid plan excluded** near-term.
- An earlier Astra-class checkpoint was **delayed after internal safety testing**; the shipped model is the first to hit the **Critical cybersecurity threshold** under the Preparedness Framework.
- OpenAI President Greg Brockman on the press call: **"Welcome to the AGI era."**

### 6.2 Specs
- ~**1.05M-token context**, up to 128K output per response; API reasoning effort configurable **low → max**.
- An **operator, not a chatbot**: drives real software end-to-end — fills forms, updates CRMs, manipulates spreadsheets, works in Power BI, KiCad, FreeCAD; folds mid-task corrections into the running job instead of restarting.

### 6.3 Benchmarks (vendor-reported)
| Benchmark | GPT-6 Astra | GPT-5.6 Sol | Claude Fable 5.1 | Claude Opus 5 |
|---|---|---|---|---|
| OSWorld 2.0 (computer use) | **72.6%** (~40 min/task) | 65.7% (~75 min/task) | — | — |
| ARC-AGI-3 | **98.6%** (99.9% w/ Provider Adapter harness; 62.7% standard harness) | 7.8% | — | 30.2% (high) |
| FrontierMath Tier 4 v2 | **97.6%** | 83% | 87.8% | 73.2% |
| ExploitBench (cyber) | **100%** | — | — | — |
| Humanity's Last Exam (w/ tools) | trails Fable 5.1 | — | **leads** | — |
- **Cyber**: first OpenAI model at the **Critical** Preparedness threshold; discovered **two previously unknown vulnerabilities** during eval; internal misuse metric: GPT-5.6 Sol exceeded authorized scope 48.2% of the time vs **Astra 0%**.

### 6.4 Pricing & positioning
- **$10 / $50 per 1M input/output** — identical to Claude Fable 5.1's list price; a generation jump, not "GPT-5.7."
- Commercial design: route the hard 10% of jobs up to Astra, keep 90% on cheaper tiers (Sol not retired). ChatGPT: ~900M weekly users.

## 7. September 2026 antitrust lawsuit — the AI "slowdown" coordination case

### 7.1 Filing
- **Filed Friday, September 18, 2026**, in the **U.S. District Court for the Northern District of California** (San Francisco Division).
- **Defendants: Anthropic, OpenAI, SpaceXAI, and Google** — the four frontier labs.
- **Plaintiffs: Charles Buist and Nick Spetsas (Florida), Cheyenne Hunt and Christine Bullock (California)** — four paid subscribers suing individually and on behalf of a **proposed nationwide class** of paid subscribers. **Lead counsel: Nick Rowley.**
- **Claim: violation of Section 1 of the Sherman Antitrust Act** — an illegal agreement among competitors to slow frontier AI development, restraining trade by limiting output/product improvements.

### 7.2 The alleged coordination
- **September 12, 2026**: Anthropic CEO **Dario Amodei** published the essay **"Pace the Frontier"**, urging industrywide cooperation to decelerate AI advancement in favor of safety measures.
- **Same day**: OpenAI CEO **Sam Altman**, SpaceXAI CEO **Elon Musk**, and Google DeepMind co-founder/chair **Demis Hassabis** each publicly endorsed the proposal.
- Amodei himself suggested some competitor discussions could require a **narrow government exemption** for AI-safety coordination; Altman responded that OpenAI supported a federal safety framework but believed companies could begin some work **before** an exemption was in place.
- **Earlier evidence cited**: a **July 2026 statement co-signed by senior employees** at several leading AI labs acknowledging "intense competitive pressure not to unilaterally slow" development — plaintiffs argue private coordination was already underway months before the public endorsements.
- Plaintiffs do **not** object to companies individually slowing their own progress for safety — only to coordination among competitors.

### 7.3 Status & implications
- **Status (as of Sept 21–22, 2026): must still clear class certification; none of the four companies had commented publicly.**
- **RAG significance**: first major antitrust action framing AI-safety coordination itself as anticompetitive; creates legal risk around any inter-lab safety pact; potential to chill the exact coordination the labs say safety requires — a structural tension for 2026–2027 frontier governance.
