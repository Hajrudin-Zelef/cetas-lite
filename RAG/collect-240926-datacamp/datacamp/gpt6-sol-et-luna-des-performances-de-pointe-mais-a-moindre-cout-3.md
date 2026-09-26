---
id: collect-240926-datacamp/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout-3
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["luna", "sol", "agent", "agents", "alignment", "astra", "benchmarks", "chatgpt", "claude", "consumer", "cost", "gpt-6"]
source: docs/RAG/clean_en/datacamp/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [247, 285]
sha256: c87332dd1f509111b03e2f2b465dfdf628212172bd22a886030137b0029f5dba
---

# gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout

If you are migrating from GPT‑5.6, OpenAI's GPT‑6 model recommendations, written for Astra, advise removing `temperature` and `top_p`, starting at `low` if you were using `none` or `minimal`, otherwise keeping your current effort level, and replacing `prompt_cache_retention` with `prompt_cache_options.ttl` set to 30 minutes.

## To conclude

Sol and Luna take the training recipe from the flagship Astra model to make the lower tier significantly more affordable and faster. OpenAI's communication explicitly cites Anthropic's models and claims gains on cost-adjusted benchmarks. The comparison is relevant, because Opus 5.5 also focuses on efficiency.

My take: if you run agents or coding workloads on GPT‑5.6 Sol or Luna, switch now. Same effort scale, half the price, and in our test, the only thing GPT‑6 Sol did differently from its predecessor was refuse to give a confident answer to a question the algorithm could not solve. If you are on Claude, the cost-adjusted figures argue for running your own comparison rather than taking OpenAI at its word.

If you want to build on these models, we recommend the OpenAI Fundamentals skill track, which covers the API end to end in 15 hours.

I am a writer and editor in the field of data science. I am particularly interested in linear algebra, statistics, R, and so on. I also play a lot of chess!

**Editor-in-Chief, Data Science at DataCamp |** **I am passionate about forecasting and development using APIs.**

## FAQ on GPT‑6 Sol and Luna

### What are GPT‑6 Sol and Luna, and what is their relationship to GPT‑6 Astra?

These are two new variants in the GPT‑6 family, positioned below Astra (OpenAI's top-tier model) in terms of price and capability. They use training methods similar to Astra, but are optimized to reduce costs and improve latency, in order to meet everyday needs rather than the most demanding projects for which Astra is reserved.

### How are they cheaper than the previous generation?

Both models cost 50% less than their GPT‑5.6 promotional rates. Concretely: Sol goes from $4/$20 to $2/$10 per million input/output tokens, and Luna from $0.20/$1.20 to $0.10/$0.50 per million tokens.

### How do they compare with competitors like Claude?

OpenAI claims strong cost-efficiency — for example, on AutomationBench, GPT‑6 Sol at high effort outperforms Claude Opus 5 for a fraction of its cost per task. Similar comparisons are presented on coding benchmarks (FrontierCode, DeepSWE) and computer use (OSWorld), generally highlighting better or comparable scores at a significantly lower cost. Note that these are results reported by OpenAI, and that competitor figures come from public sources rather than tests conducted by OpenAI.

### What improvements have been made to caching and alignment?

The caching improvements aim for higher hit rates by default (with up to a 90% discount on cached input tokens), as well as new tools such as a cache dashboard and reasoning/tooling effort settings that preserve the cache. On alignment, both models show improved honesty indicators (e.g., fewer misleading claims about coding work) compared with their GPT‑5.6 predecessors.

### When and where can I access these models?

They are available now in ChatGPT Work and Codex for Plus, Pro, Business, Enterprise, and Edu subscribers, with Luna also accessible to Free/Go users in the desktop app. Via the API, they are accessible as `gpt-6-sol` and `gpt-6-luna`. They are not yet available in the standard consumer ChatGPT offering, and the rollout is happening gradually throughout the day.

### Should I use GPT‑6 Sol or GPT‑6 Luna?

Use Sol for agent workflows, coding on real repositories, and highly factual professional work; it is the tier on which OpenAI's results on AutomationBench, DeepSWE, and factuality rest. Use Luna for high-volume, cost-sensitive tasks (classification, extraction, routing), where its $0.10 input and $0.50 output per million tokens matter more than peak capability. Luna is also the only GPT‑6 model available to Free and Go users in the ChatGPT desktop app.
