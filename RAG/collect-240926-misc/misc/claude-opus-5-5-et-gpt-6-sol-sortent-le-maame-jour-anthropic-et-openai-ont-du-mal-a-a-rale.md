---
id: collect-240926-misc/misc/claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-rale
title: "claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-ralentir-a-l-i"
domain: numerama
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "OpenAI", "SpaceX", "xAI"]
dates: ["2023-03-14", "2026-08-02"]
keywords: ["claude", "gpt-6", "sol", "agents", "agi", "astra", "benchmarks", "chatgpt", "cost", "cybersecurity", "fable 5", "gemini"]
source: docs/RAG/clean_en/misc/claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-ralentir-a-l-i.md
source_anchor: ""
source_lines: [1, 80]
sha256: cd0217a1c9aa30f81cf1c808738dec8256d055949105699c39d9195f6ff6bca3
---

# claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-ralentir-a-l-i

<!-- source: https://www.numerama.com/tech/2338131-claude-opus-5-5-et-gpt-6-sol-sortent-le-meme-jour-anthropic-et-openai-ont-du-mal-a-ralentir-lia.html -->

It’s rare to see the two leaders of a sector launch a new product simultaneously. Yet on September 22, Anthropic launched Claude Opus 5.5, which took the top spot in most rankings, and OpenAI, which estimated in early September that it had entered “*the AGI era*” with GPT-6 Astra, launched GPT-6 Sol and GPT-6 Luna. The irony is that both groups called for slowing down the AI race last week and hadn’t launched new models at the same time since March 14, 2023, when LLMs were still in their early days (GPT-4 and the first Claude).

On the agenda: better results, obviously. But above all, lower prices. With GPT-6 Luna, billed at $0.10 per million input tokens, OpenAI is even taking on open-source models and local AI: at that rate, only those seeking confidentiality have any reason to set up a machine locally.

| API Price (per million tokens) | Input | Output |
|---|---|---|
| **GPT-6 Luna** | $0.10 | $0.50 |
| **GPT-6 Sol** | $2 ($4 before) | $10 ($20 before) |
| **Claude Opus 5.5** | $4 ($5 before) | $20 ($25 before) |
| **GPT-6 Astra and Claude Fable 5.1** | $10 | $50 |

## Claude Opus 5.5: Fable 5.1-level performance for 40% less than Opus 5

The first model in the Claude 5.5 family is Opus 5.5. According to Anthropic, it reaches the level of Claude Fable 5.1 on most tasks, despite a much lower price.

On Terminal-Bench 4.0 (command-line programming), it scores 66.4%, compared to 57.9% for GPT-6 Astra and 55.8% for Fable 5.1. On GDPval-AA, which evaluates professional tasks across 44 professions, it reaches 1,846 Elo points, compared to 1,735 for Fable 5.1 and 1,542 for Astra. GPT-6 Astra retains the advantage on two tests, enterprise task automation (41.4% vs. 40%) and scientific research (64.6% vs. 58.7%). But benchmarks no longer mean much at this level.

Opus costs 2.5 times less than Fable but beats it in most tests, which raises the question of the value of Anthropic’s high-end model until the release of the next Fable 5.5 (and thus the constant AI race). The company itself acknowledges that the real gap between the two is smaller than the figures suggest. The price, however, is dropping: $4 per million input tokens and $20 output, compared to $5 and $25 for Opus 5. Pro, Max, and Team subscribers also benefit from increased usage limits over five hours and a free reset of these limits that they can save and trigger whenever they wish, similar to what OpenAI already offers.

Note, however, that in Claude, the effort selector is now set to “Medium” by default, rather than “High.” Part of the announced savings comes from this, while benchmarks are run at maximum effort. It is also no longer possible to disable thinking mode. Finally, Opus 5.5 inherits Fable 5.1’s guardrails: cybersecurity requests are redirected to Opus 4.8, biology requests to Opus 5.

### Claude complies with European regulation: Opus 5.5 watermarks its texts

Another change: Opus 5.5 is the first Opus launched with content marking, as required by the AI Act since August 2, 2026. Anthropic says it applies this marking worldwide, not just in Europe, since it cannot limit it by region. The detection tool is reserved for regulators, media, fact-checkers, and researchers. Unsurprisingly, Anthropic warns that it is not infallible proof.

## GPT-6 Sol and Luna: OpenAI halves its prices

Three weeks after Astra, OpenAI completes its GPT-6 family with GPT-6 Sol, its high/mid-range model, and GPT-6 Luna, its entry-level model, while Astra remains the most powerful. Unlike the GPT-5.6 generation, there is no Terra version for now.

As with Anthropic, prices are falling. The two new models cost half as much as their predecessors in the API: $2 and $10 for Sol, $0.10 and $0.50 for Luna.

OpenAI compares its models to Anthropic’s… but not to the new Opus 5.5, obviously.

On AutomationBench, Sol at xhigh effort reaches 33.2% for $0.27 per task, ahead of Fable 5.1 (31.4%) and Opus 5 (26.9%), which costs 11 times more per task. Opus 5.5 shows 40% on the same test according to Anthropic: it remains ahead, but at a higher price. In programming, Sol scores 68.8% on DeepSWE, 1.1 points behind Fable 5 for about 80% less cost. Luna reaches 66.6%, on par with Opus 5 at medium effort. OpenAI also announces half as many factual errors as with GPT-5.6 Sol, according to an internal evaluation.

OpenAI takes a jab at its rival along the way: the cost of Fable 5.1 is said to be underestimated, as it doesn’t include fallbacks to Opus 5, which occurred on about 40% of the tasks in this test. A mechanism that Opus 5.5 now also uses.

GPT-6 Sol and GPT-6 Luna are available today in ChatGPT Work and Codex for Plus, Pro, Business, Enterprise, and Edu subscribers. Free and Go users have access to Luna in the desktop app.

## Grok 4.7 released the day before, Google absent from subscribers

The day before, xAI, now integrated into SpaceX, had launched Grok 4.7… while Elon Musk also called for slowing down the AI race. Billed at $2 input and $6 output, it remains a notch below Anthropic’s and OpenAI’s models. Elon Musk himself placed it at roughly the level of Claude Opus 5.

All that's left is Google, a fallen giant. The company promised that Gemini 3.5 Pro would arrive in June… but its poor performance is gradually condemning it to never being released. The company now prefers to talk about Gemini 4, which is in training, but no one knows if it will manage to return to the level of Anthropic or OpenAI, which is now betting everything on agents, or even Meta and its Muse, which is convincing more than one person.

+ faster, + practical, + exclusive

Zero advertising, advanced reading features, AI-summarized articles, exclusive content and more.

Discover the many benefits of Numerama+.

You have read **0 articles** on Numerama this month

            **Not everyone can afford to pay** for information.

            That's why we keep our journalism open to everyone.
        

            **But if you can,**

            here are three good reasons to support our work:
        

- 
                1
                Numerama+ helps provide **a free experience for all Numerama readers**.
- 
                2
                You will enjoy **ad-free reading**, many **advanced reading features and exclusive content**.
- 
                3
                **Help Numerama in its mission**: understanding the present to anticipate the future.

**If you believe in a free web** and in quality information accessible to as many people as possible, join Numerama+.

All the tech news at a glance

Add Numerama to your home screen and stay connected to the future!
