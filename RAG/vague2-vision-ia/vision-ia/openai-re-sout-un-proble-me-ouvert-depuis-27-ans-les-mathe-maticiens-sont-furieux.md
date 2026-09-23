---
id: vague2-vision-ia/vision-ia/openai-re-sout-un-proble-me-ouvert-depuis-27-ans-les-mathe-maticiens-sont-furieux
title: "OpenAI résout un problème ouvert depuis 27 ans, les mathématiciens sont furieux"
domain: vision-ia
role: reference
task: article
actors: ["Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "Google", "Mistral", "OpenAI", "Z.ai"]
dates: ["2026-06-29", "2026-08-01", "2026-09-11", "2026-09-23"]
keywords: ["agent", "agentic", "agents", "astra", "attribution", "benchmark", "chatgpt", "claude", "deepseek", "distillation", "full-duplex", "gemini"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/openai-re-sout-un-proble-me-ouvert-depuis-27-ans-les-mathe-maticiens-sont-furieux.md
source_anchor: ""
source_lines: [1, 51]
sha256: 0751b67a9deb0f47410e5092df2cd4ca889ebee2351b237099d76ee9a726d974
---

# OpenAI résout un problème ouvert depuis 27 ans, les mathématiciens sont furieux

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/openai-re-sout-un-proble-me-ouvert-depuis-27-ans-les-mathe-maticiens-sont-furieux
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This September 11, 2026 issue leads with a controversy over OpenAI's August 1, 2026 announcement that its Astra model solved ten open math problems, including the first construction of a non-sofic group—a question posed by Gromov 27 years ago. The proof directly builds on a 2019 paper by Andreas Thom and Gábor Kun. Thom, a group theorist at TU Dresden, had spent months working on precisely this subject while conversing with ChatGPT. He now asks OpenAI to prove that his private conversations did not feed the model's training. Thom detailed the accusation in a series of Mastodon posts, calling the behavior "dishonest" and the data-origin opacity total. He had disabled training on his data on June 29, 2026, already doubting traceability; this opt-out setting exists for everyone in ChatGPT's data controls. OpenAI researcher Mark Sellke replied "that did not happen," a formula that only covers direct paper access during resolution, not conversation use in the training corpus. On another matter (Navier-Stokes equations), OpenAI admitted it cannot rule out that "de-identified" data improved its models. Thom is the second mathematician in days to raise this accusation, and the company has produced no technical traceability proof.

Next, OpenAI launched ChatGPT for Financial Services on September 10, combining built-in financial datasets (Daloopa, PitchBook, LSEG News, Crunchbase) with GPT-6 Astra. It targets junior analyst work: valuations, LBO models, buyer screening, earnings analysis, pitchbooks, and client notes. Each figure carries a granular citation back to its source. It was designed with Morgan Stanley and Evercore as design partners, with auth/permission integrations at FactSet, S&P Global, Preqin, and Datasite. Availability is on request for eligible financial institutions; pricing is not published (a source mentions ~$10 per million tokens, unconfirmed). Notably absent: no mention of European hosting or DORA compliance.

OpenAI also opened GPT-Live-1, its full-duplex voice model, to developers at $0.05 per minute. Full-duplex interactivity is measured at 80.1% versus 45.4% for GPT-Realtime-2.1; speech-start latency drops to 0.8s from 1.4s; tool-call accuracy rises to 87% from 60%; and on a voice banking support benchmark success rises to 32% from 12.4%. Twelve new voices are included. Yelp already uses it for phone reservations.

Anthropic published "Detecting and Countering Misuse of AI" on September 10, documenting five cases where Claude accounts aided work that could facilitate biological weapons. The most notable: a user asked Claude to write a grant request for gain-of-function research on the chikungunya virus to make it more infectious and immune-evasive. Detection came from Anthropic's internal biological safety classifier. All accounts were banned, but individuals and labs were not named; information was shared with authorities and other AI companies.

Research briefs cover Colibrì running GLM-5.2 (744B parameters) on 25 GB RAM; DeepSeek V4.1-Flash cutting agent memory 4x (552B params, 16B active, MIT license); MultiMatte text-guided background removal; Codex hunting antibiotics in extinct genomes; "agentic flooding" of public services; Anthropic quantifying Chinese distillation campaigns (~200 million exchanges, 151 million attributed to Alibaba); GPT-6 Astra topping ErdosBench; autonomous agents tracked across 30 sites; Anthropic's seven-domain threat report; and why AI research agents don't overfit. Industry briefs include NASA-IBM's open lunar model, Anthropic's Mythos 5 PyPI incident, Gemini on Windows, OpenAI's Data agent, Slack Surfaces, Universal Music–ElevenLabs deal, JD.com's robot plan, Maven Robotics, the iPhone Duo in China, and Mistral's €3B raise.

## Key points

- OpenAI's claim to have solved a 27-year-old group theory problem is contested by Andreas Thom, who says his months of ChatGPT conversations may have trained the model.
- OpenAI provided no technical traceability proof; its researcher's denial only covered direct paper access.
- ChatGPT for Financial Services launched with GPT-6 Astra and built-in datasets, targeting junior analyst work, with no European hosting/DORA mention.
- GPT-Live-1 full-duplex voice opened at $0.05/min with large interactivity and latency gains.
- Anthropic documented five cases of Claude potentially aiding biological-weapons research, including chikungunya gain-of-function.
- DeepSeek V4.1-Flash cuts agent memory 4x under an MIT license.

## Technical data / figures

| Item | Value |
| --- | --- |
| Thom training opt-out date | June 29, 2026 |
| GPT-Live-1 price | $0.05 per minute |
| GPT-Live-1 full-duplex | 80.1% (vs 45.4% GPT-Realtime-2.1) |
| GPT-Live-1 speech latency | 0.8s (vs 1.4s) |
| GPT-Live-1 tool-call accuracy | 87% (vs 60%) |
| Voice banking support success | 32% (vs 12.4%) |
| Colibrì / GLM-5.2 | 744B params on 25 GB RAM (~40B active/token) |
| DeepSeek V4.1-Flash | 552B params, 16B active, KV cache /4, MIT |
| Anthropic Chinese distillation | ~200M exchanges; 151M (Alibaba, May–Jul 2026) |
| ChatGPT for Financial Services data | Daloopa, PitchBook, LSEG News, Crunchbase |
| Financial Services price (unconfirmed) | ~$10 / M tokens |

## Why this source matters for the RAG

It documents a landmark provenance/attribution controversy—whether private expert conversations trained a frontier model—central to data-governance and AI-credit debates. It also captures enterprise AI expansion into regulated finance and full-duplex voice with precise metrics, plus Anthropic's biosecurity misuse findings. These are high-value for safety, legal, and product-trend queries.
