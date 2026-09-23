---
id: collect-presse-fr/presse-fr/meta-muse-spark-1-3
title: "Meta lance Muse Spark 1.3, axé sur le code et le contexte long"
domain: presse-fr
role: reference
task: article
actors: ["Anthropic", "Google", "Meta", "OpenAI"]
dates: ["2026-08", "2026-09", "2026-09-23"]
keywords: ["muse", "muse spark", "agent", "agentic", "benchmark", "benchmarks", "claude", "gemini", "gpt-5.6", "open weights", "open-weight", "opus 5"]
source: docs/RAG/Collect RAG/05_presse_fr/meta-muse-spark-1-3.md
source_anchor: ""
source_lines: [1, 50]
sha256: 9e87052c424ac940ac7ea44fc38d3cb92609bf2c14c0e5eedf133c4b00e0e555
---

# Meta lance Muse Spark 1.3, axé sur le code et le contexte long

## Metadata

- **Source** : https://www.blogdumoderateur.com/meta-muse-spark-1-3/
- **Site** : BDM (Blog du Modérateur)
- **Type** : Article de presse
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Published 3 September 2026 by Matthieu Eugène, this BDM article covers Meta's release of Muse Spark 1.3, the fourth version in five months, from Meta Superintelligence Labs. It arrives one month after Muse Spark 1.2, which shipped alongside Meta's programming agent Muse Code. Muse Spark 1.3 is available in Muse Code and via the Meta Model API. Meta continues its rapid cadence on homegrown models, much like Google with Gemini.

The update is positioned around **sustained performance over time rather than raw power**. Meta says the model can run several workstreams in parallel within a single conversation thread, build its own context from messy or contradictory sources, and correct gaps in its plan as it goes. The company also emphasizes more cooperative behavior: the model asks questions when a request is ambiguous, solicits the user when stuck, and asks for confirmation before high-consequence actions. The most notable argument concerns **self-awareness of limits**: Meta says it trained the model to better understand "what it can and cannot do, what it knows and does not know," rather than inventing a result when blocked. On code, Meta claims a less verbose style and fewer useless loop turns, with roughly **20% fewer tool calls and 25% fewer tokens** according to its own engineers' comparisons.

Like every release, Meta publishes its own scorecard against OpenAI's GPT-5.6 Sol and Anthropic's Claude Opus 5. Muse Spark 1.3 leads its two rivals on code and long-context management. However, it takes **no first place on the six agentic evaluations** selected, even though those tasks are one of the announcement's selling points. The table requires a major caveat: the scores attributed to Muse Spark 1.3 come from its **"max reasoning" mode, which is not available at launch** — Meta is delaying it pending additional safety tests. Muse Spark 1.2 was measured in a lower reasoning mode. Part of the displayed gap between the two versions therefore stems from configuration, not from the model generation — further proof of the limits of vendor-published benchmarks.

Meta also announces larger models and the publication of the weights of a Muse Spark version, without a timeline. This reiterates a commitment made in August, at the time of Mark Zuckerberg's superintelligence manifesto, which came with the open-weight model Muse Glimmer. The article closes with related Meta coverage (Meta One subscription, Muse agent, etc.).

## Key points

- Muse Spark 1.3 is Meta's fourth Muse Spark version in five months, one month after 1.2.
- Available in Muse Code and via the Meta Model API.
- Focus: long-context management, parallel workstreams, self-awareness of limits, cooperative behavior.
- Claims ~20% fewer tool calls and 25% fewer tokens on code.
- Leads GPT-5.6 Sol and Claude Opus 5 on code and long context in Meta's own scorecard.
- Takes no first place on the six agentic evaluations, despite agentic tasks being a key selling point.
- Published 1.3 scores use a "max reasoning" mode not yet publicly available (delayed for safety tests), partly invalidating the 1.2-vs-1.3 gap.
- Meta reaffirms a promise to publish weights of a Muse Spark version (no timeline), after Muse Glimmer in August.

## Technical data / figures

| Item | Detail |
|---|---|
| Model | Muse Spark 1.3 |
| Developer | Meta Superintelligence Labs |
| Release date | 3 September 2026 |
| Previous version | Muse Spark 1.2 (one month earlier) |
| Availability | Muse Code, Meta Model API |
| Tool calls | ~20% fewer (Meta internal comparison) |
| Tokens | ~25% fewer (Meta internal comparison) |
| Benchmarks | Leads on code and long context; no first place on 6 agentic evals |
| Benchmark caveat | 1.3 scores from unreleased "max reasoning" mode |
| Open weights | Promised for a Muse Spark version (no date); Muse Glimmer released August 2026 |

## Why this source matters for the RAG

This article documents Meta's Muse Spark lineage and the release cadence of its homegrown models, while flagging the methodological weakness of vendor-published benchmarks (unreleased max-reasoning mode). It is useful for questions about Meta's AI strategy, open-weight commitments, and benchmark reliability.
