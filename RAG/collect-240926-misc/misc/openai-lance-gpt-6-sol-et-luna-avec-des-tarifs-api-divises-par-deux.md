---
id: collect-240926-misc/misc/openai-lance-gpt-6-sol-et-luna-avec-des-tarifs-api-divises-par-deux
title: "openai-lance-gpt-6-sol-et-luna-avec-des-tarifs-api-divises-par-deux"
domain: blogdumoderateur
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["gpt-6", "luna", "sol", "agents", "astra", "chatgpt", "claude", "cost", "fable 5", "gpt-5.6", "opus 5", "pricing"]
source: docs/RAG/clean_en/misc/openai-lance-gpt-6-sol-et-luna-avec-des-tarifs-api-divises-par-deux.md
source_anchor: ""
source_lines: [1, 32]
sha256: e52d1399b16c14e7e7e0a6de18d9249eefb4af7cbc8c44e00debf3c280593058
---

# openai-lance-gpt-6-sol-et-luna-avec-des-tarifs-api-divises-par-deux

<!-- source: https://www.blogdumoderateur.com/openai-lance-gpt-6-sol-luna-tarifs-api-divises-deux/ -->

In early September, OpenAI unveiled GPT-6 Astra, its most powerful model to date. On this September 22, the startup completed its lineup with GPT-6 Sol and GPT-6 Luna. Trained with methods similar to those used for Astra, these two models are sold at significantly lower prices. Sol is designed for complex tasks such as software development, while Luna targets high-volume repetitive tasks, such as document summarization, information extraction, or quick answers.

The announcement came barely 90 minutes after Anthropic's unveiling of Claude Opus 5.5.

## Performance close to high-end, at a much lower cost

According to tests published by OpenAI, GPT-6 Sol rivals Anthropic's high-end models on several fronts.

On AutomationBench, which evaluates the ability to chain business tasks across different applications, it outperforms Claude Opus 5 and Claude Fable 5.1, with a cost per task 11.1 times lower than that of Opus 5. In development, it approaches Fable 5's best score on DeepSWE, a test based on complex tasks in real projects, for roughly 80% less cost. Luna, for its part, reaches a level comparable to that of Opus 5 at medium effort on the same test, for a cost per task 93% lower. OpenAI also announces roughly half as many factual errors with Sol as with GPT-5.6 Sol.

Autonomous computer use is also improving. On OSWorld 2.0, GPT-6 Sol achieves a score similar to that of Claude Opus 5 at medium effort, for roughly 80% less cost. On the developer side, the company also improves GPT-6's prompt caching. Agents reuse more context from one request to the next, respond faster, and benefit from a 90% discount on input tokens read from cache.

Note: the various results do not include Claude Opus 5.5, released the same day.

## Significantly reduced pricing

GPT-6 Sol and Luna are available in ChatGPT Work and Codex for Plus, Pro, Business, Enterprise, and Edu subscribers. Free and Go accounts have access to Luna in the desktop app. Both models are also available in the API, but not in the ChatGPT chat.

In the API, OpenAI halves its prices compared to GPT-5.6's promotional rates. The new rates, per million tokens:

- **GPT-6 Sol:** $2 input and $10 output, compared to $4 and $20 previously.
- **GPT-6 Luna:** $0.10 input and $0.50 output, compared to $0.20 and $1.20 previously.

## Clearer and shorter responses, in the vein of Astra

Sol and Luna inherit the response style introduced with Astra. OpenAI promises responses *"clearer, with less jargon, fewer unusual turns of phrase, fewer superfluous details and, overall, slightly shorter, without losing substance."*

The change will be noticeable, according to the firm, especially in technical exchanges and around code.

GPT-6 Sol does not jump to hasty conclusions, spends less time repeating details that might be obvious to the interlocutor, uses less vague language, is more transparent about what it has verified and what it has not verified, and does not unnecessarily share implementation details, OpenAI explains.
