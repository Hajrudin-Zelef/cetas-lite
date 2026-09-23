---
id: collect-presse-fr/presse-fr/openai-gpt-6-sol-luna-tarifs
title: "OpenAI lance GPT-6 Sol et Luna, avec des tarifs API divisés par deux"
domain: presse-fr
role: reference
task: article
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09", "2026-09-23"]
keywords: ["gpt-6", "luna", "sol", "agents", "astra", "benchmarks", "chatgpt", "claude", "cost", "fable 5", "gpt-5.6", "opus 5"]
source: docs/RAG/Collect RAG/05_presse_fr/openai-gpt-6-sol-luna-tarifs.md
source_anchor: ""
source_lines: [1, 54]
sha256: 91161d61b9726a1887f5509d49d04bb78e2212a54227fe255c2a40e8463e24bd
---

# OpenAI lance GPT-6 Sol et Luna, avec des tarifs API divisés par deux

## Metadata

- **Source** : https://www.blogdumoderateur.com/openai-lance-gpt-6-sol-luna-tarifs-api-divises-deux/
- **Site** : BDM (Blog du Modérateur)
- **Type** : Article de presse
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Published 23 September 2026 by José Billon, this BDM article covers OpenAI's expansion of the GPT-6 family with GPT-6 Sol and GPT-6 Luna, announced on 22 September, three weeks after GPT-6 Astra. The two models were trained with methods close to Astra's but are sold much cheaper. Sol targets complex tasks such as software development, while Luna targets high-volume repetitive tasks like document summarization, information extraction, and quick answers. The announcement landed barely 90 minutes after Anthropic's Claude Opus 5.5.

**Performance:** According to OpenAI's published tests, GPT-6 Sol rivals Anthropic's high-end models on several fronts. On AutomationBench, which evaluates chaining business tasks across applications, it beats Claude Opus 5 and Claude Fable 5.1 with a cost per task **11.1× lower than Opus 5**. In development, it approaches Fable 5's best score on DeepSWE (a test based on complex tasks in real projects) for about **80% less cost**. Luna reaches a level comparable to Opus 5 at medium effort on the same test, for a cost per task **93% lower**. OpenAI also claims roughly **twice fewer factual errors** with Sol than with GPT-5.6 Sol.

Autonomous computer use also improves: on OSWorld 2.0, GPT-6 Sol scores similarly to Claude Opus 5 at medium effort for about 80% less cost. For developers, OpenAI improves GPT-6 prompt caching: agents reuse more context between requests, respond faster, and benefit from a **90% discount on cached input tokens read**. Important caveat: these results do **not include Claude Opus 5.5**, released the same day.

**Pricing and availability:** GPT-6 Sol and Luna are available in ChatGPT Work and Codex for Plus, Pro, Business, Enterprise and Edu subscribers. Free and Go accounts get Luna in the desktop app. Both models are also available in the API, but **not in the ChatGPT chat**. In the API, OpenAI halves prices versus GPT-5.6's promotional rates:

- **GPT-6 Sol**: $2 input / $10 output per million tokens (previously $4 / $20).
- **GPT-6 Luna**: $0.10 input / $0.50 output per million tokens (previously $0.20 / $1.20).

**Response style:** Sol and Luna inherit the response style introduced with Astra. OpenAI promises "clearer" answers with less jargon, fewer unusual phrasings, fewer superfluous details, and overall slightly shorter responses without losing substance. The change will be most noticeable in technical and code-related exchanges. GPT-6 Sol reportedly avoids jumping to conclusions, spends less time repeating details obvious to the interlocutor, uses less vague language, is more transparent about what it has and has not verified, and avoids unnecessarily sharing implementation details.

## Key points

- OpenAI launched GPT-6 Sol and Luna on 22 September 2026, ~90 minutes after Claude Opus 5.5, three weeks after Astra.
- Sol targets complex tasks (software development); Luna targets high-volume repetitive tasks.
- Sol beats Claude Opus 5 and Fable 5.1 on AutomationBench at 11.1× lower cost per task vs Opus 5.
- On DeepSWE, Sol approaches Fable 5's best score for ~80% less cost; Luna is ~93% cheaper per task than Opus 5 medium.
- ~2× fewer factual errors with Sol vs GPT-5.6 Sol (OpenAI internal evaluation).
- Prompt caching improved with a 90% discount on cached input tokens.
- API prices halved: Sol $2/$10, Luna $0.10/$0.50 per million tokens.
- Available in ChatGPT Work and Codex (Plus/Pro/Business/Enterprise/Edu); Free/Go get Luna in the desktop app; not in the ChatGPT chat.
- Benchmarks exclude the same-day Claude Opus 5.5.

## Technical data / figures

| Model | Input ($/M tokens) | Output ($/M tokens) | Previous (GPT-5.6 promo) |
|---|---|---|---|
| GPT-6 Sol | 2 | 10 | 4 / 20 |
| GPT-6 Luna | 0.10 | 0.50 | 0.20 / 1.20 |

- **Cost per task vs Claude Opus 5**: Sol on AutomationBench 11.1× lower; Sol on DeepSWE ~80% less; Luna on DeepSWE 93% lower.
- **Factual errors**: ~2× fewer for Sol vs GPT-5.6 Sol.
- **Prompt caching**: 90% discount on cached input tokens.
- **OSWorld 2.0**: Sol ≈ Claude Opus 5 (medium effort) at ~80% less cost.
- **Availability**: ChatGPT Work + Codex (Plus, Pro, Business, Enterprise, Edu); Luna for Free/Go desktop; API available, chat not.

## Why this source matters for the RAG

This article provides precise pricing and availability details for GPT-6 Sol and Luna, plus OpenAI's cost-per-task comparisons against Anthropic models and the prompt-caching improvement. It is valuable for questions about API economics, model routing between Sol and Luna, and ChatGPT plan access.
