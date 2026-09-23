---
id: collect-mindstudio/mindstudio/meta-muse-spark-1-3-open-weight
title: "Meta Muse Spark 1.3 open weight - MindStudio"
domain: mindstudio
role: reference
task: article
actors: ["AWS", "Anthropic", "Meta", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["muse", "muse spark", "agentic", "aws", "benchmark", "benchmarks", "context window", "cost", "gpt-5.6", "multimodal", "open-weight", "opus 5"]
source: docs/RAG/Collect RAG/02_mindstudio/meta-muse-spark-1-3-open-weight.md
source_anchor: ""
source_lines: [1, 52]
sha256: 398a372b3adc505cebfca9678ab9fc070bc43b8e105e6c6015e6fb22f00b0c71
---

# Meta Muse Spark 1.3 open weight - MindStudio

## Metadata

- **Source** : https://www.mindstudio.ai/blog/meta-muse-spark-1-3-open-weight
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article reports a hands-on test of Meta's **Muse Spark 1.3**, a multimodal reasoning model designed for long-horizon agentic work, coding, and computer use, featuring a 1 million token context window.

**Pricing and availability.** Meta prices Muse Spark 1.3 at $1.25 per million input tokens and $4.25 per million output tokens, undercutting most frontier competitors. Meta founder Mark Zuckerberg confirmed the model will be released as open weight soon, which would put a frontier-class model within reach of anyone willing to self-host.

**Benchmark picture.** Muse Spark tops the field on professional tool use and agentic computer use, edging out GPT-5.6 and Opus 5. On end-to-end business workflows it runs neck and neck with Opus 5. Coding is the closest race: Muse Spark, GPT-5.6 and Opus 5 cluster tightly on long-horizon agentic coding tasks. The standout is long-context retrieval: on the million-token retrieval test Muse Spark scores in the high 90s while GPT-5.6 and Opus 5 fall into the 60s and 70s (Opus reportedly posting no score at all). Opus 5 still reportedly leads on some knowledge work tasks.

**Real-world agentic deployment.** In the flagship test the model was given a live AWS account with no existing resources and one instruction: build a self-contained animated website and deploy it globally via S3 + CloudFront. Working in parallel (drafting the HTML animation while provisioning infrastructure), it created the S3 bucket, linked it as the CloudFront origin, and returned a live working URL — an animated rotisserie chicken simulation with flickering flames, rotating spit, fire cam overlay, and cook-time/temperature counters — from a single prompt with zero manual intervention.

**Vision / situational reasoning.** On a WhatsApp screenshot thread (an employee mentions the boss's wife in a work message), the model correctly attributed messages by bubble position, recognized that the wife only saw part of the conversation, understood the joke, and correctly concluded the employee was still on for the evening event.

**Scientific and multilingual reasoning.** On a combined chemistry+math problem (buffer pH after adding strong acid, with two acid dissociation stages), the model correctly identified which reaction dominates, justified ignoring the second equilibrium stage based on the ~5.55 pKa gap, and verified that fully protonated species formation was negligible. On a multilingual test covering 79 languages, it produced correct native-script answers with culturally specific picks (soju for Korean, rakija for Serbo-Croatian, kumis for Kazakh, kava for Hawaiian, lassi for Punjabi, etc.). Total cost for all tests combined including the cloud deployment: roughly $2.49.

**Caveat.** Heavy server load caused repeated throttling during testing, suggesting demand is currently outpacing capacity.

## Key points

- Muse Spark 1.3: multimodal reasoning model, 1M token context, $1.25 in / $4.25 out per million tokens.
- Confirmed by Zuckerberg to become open weight soon.
- Wins on professional tool use and agentic computer use; long-context retrieval in the high 90s vs 60s-70s for GPT-5.6 and Opus 5.
- Completed a fully autonomous AWS S3 + CloudFront deployment from a single prompt.
- Strong social/situational vision reasoning and rigorous multi-step scientific reasoning.
- All tests combined cost roughly $2.49; throttling noted under load.

## Technical data / figures

| Metric | Muse Spark 1.3 | GPT-5.6 | Opus 5 |
|---|---|---|---|
| Context window | 1M tokens | — | — |
| Input price (per M tokens) | $1.25 | — | — |
| Output price (per M tokens) | $4.25 | — | — |
| Long-context retrieval (1M) | high 90s | 60s-70s | 60s-70s / no score |
| Agentic computer use | top | edged out | edged out |
| Professional tool use | top | edged out | edged out |
| Coding (long-horizon) | tight cluster | tight cluster | tight cluster |
| Knowledge work | strong | — | reportedly leads |

## Why this source matters for the RAG

Provides current, first-hand performance data on a frontier-class open-weight model about to be released, including concrete pricing, benchmarks, and real-world agentic capabilities — useful for reducing hallucinations about model capabilities and for up-to-date competitive comparisons against GPT-5.6 and Opus 5.
