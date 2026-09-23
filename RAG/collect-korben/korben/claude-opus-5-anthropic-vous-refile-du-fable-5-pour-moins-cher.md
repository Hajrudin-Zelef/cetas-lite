---
id: collect-korben/korben/claude-opus-5-anthropic-vous-refile-du-fable-5-pour-moins-cher
title: "Claude Opus 5 : Anthropic vous refile du Fable 5 pour moins cher"
domain: korben
role: reference
task: article
actors: ["Anthropic", "Google", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["claude", "fable 5", "opus 5", "agent", "agentic", "agi", "benchmark", "benchmarks", "context window", "opus 4", "pricing", "reasoning"]
source: docs/RAG/Collect RAG/01_korben/claude-opus-5-anthropic-vous-refile-du-fable-5-pour-moins-cher.md
source_anchor: ""
source_lines: [1, 57]
sha256: 5de4a5e715f0c4ae78f240f931e205ad4d0ea41ea18f8cb115d5a1995bfd2371
---

# Claude Opus 5 : Anthropic vous refile du Fable 5 pour moins cher

## Metadata

- **Source** : https://korben.info/claude-opus-5-anthropic-vous-refile-du-fable-5-pour-moins-cher.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article covers Anthropic's announcement of Claude Opus 5, a new model positioned just below its flagship Fable 5 while costing roughly half as much to use. Anthropic, the company behind the Claude assistant and a main rival of OpenAI and Google, made the announcement the day before.

The pricing does not change: $5 per million input tokens and $25 per million output tokens, exactly the same as the previous Opus 4.8. A token is the small piece of text the model reads or writes, roughly a syllable, and is the billing unit for all these AIs. The context window, the amount of text the model holds at once, reaches one million tokens, the equivalent of a very large book, and the response can extend up to 128,000 tokens. Reasoning is enabled by default, so the model takes time to reason before answering, and those in a hurry can switch to a mode about 2.5 times faster, billed at double the rate.

The announced progress concerns mainly code, agent tasks (where the AI chains actions on its own to accomplish work), and computer use (where it directly drives a computer by clicking, navigating and filling forms like a human). On benchmarks, the standardized tests used to compare models, Opus 5 goes from 19 to 43% on FrontierBench, from 1.5 to 30% on ARC-AGI-3 (which measures solving never-before-seen problems), and climbs to 71% on OSWorld, the computer-driving test, versus 56% for its predecessor.

On SWE-bench Verified, which has it fix real code bugs, it reaches 96%. However, on the harder version of that test, Fable 5 keeps a very short lead, 80% versus 79%, so parity with the house flagship deserves an asterisk. Another caveat: all these figures come from Anthropic itself, and experience shows there is sometimes a gap between announcement scores and daily use.

On subscriptions, Opus 5 becomes the default model of the Claude Max offering and the most powerful accessible with a Claude Pro subscription. The real question is therefore no longer whether this model is good, but who still has a reason to pay twice as much for Fable 5. The author concludes that Opus 5 becomes the default choice, with Fable 5 confined to the rare uses where the last percentage point is paid without question.

## Key points

- Claude Opus 5 sits just below the flagship Fable 5 while costing about half as much to use.
- Pricing unchanged from Opus 4.8: $5/million input tokens, $25/million output tokens.
- Context window of 1 million tokens; responses up to 128,000 tokens; reasoning enabled by default.
- A ~2.5x faster mode is available but billed at double the rate.
- Main gains in code, agent tasks and computer use.
- Benchmarks: FrontierBench 19→43%, ARC-AGI-3 1.5→30%, OSWorld 56→71%, SWE-bench Verified 96%.
- On the harder SWE-bench version, Fable 5 still leads narrowly, 80% vs 79%.
- Opus 5 becomes default for Claude Max and the strongest model in Claude Pro; all figures are from Anthropic.

## Technical data / figures

| Item | Value |
|---|---|
| Model | Claude Opus 5 (Anthropic) |
| Flagship above it | Fable 5 |
| Input price | $5 per million tokens |
| Output price | $25 per million tokens |
| Same price as | Opus 4.8 |
| Context window | 1,000,000 tokens |
| Max response | 128,000 tokens |
| Reasoning | Enabled by default |
| Fast mode | ~2.5x faster, double price |
| FrontierBench | 19% → 43% |
| ARC-AGI-3 | 1.5% → 30% |
| OSWorld | 56% → 71% |
| SWE-bench Verified | 96% |
| Hard SWE-bench | 79% (Fable 5: 80%) |
| Availability | Default on Claude Max; strongest in Claude Pro |

## Why this source matters for the RAG

This article provides a detailed release overview of a major Anthropic model, with pricing, context limits and benchmark comparisons against a higher-tier model. It is valuable for questions about LLM pricing, capability tiers, agentic coding and computer-use performance.
