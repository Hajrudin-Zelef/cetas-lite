---
id: collect-240926-mindstudio/mindstudio/ox-alpha-free-on-opencode-pricing-limits-and-how-long-it-lasts
title: "ox-alpha-free-on-opencode-pricing-limits-and-how-long-it-lasts"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Moonshot", "Xiaomi", "Z.ai", "xAI"]
dates: []
keywords: ["pricing", "agent", "agentic", "agents", "benchmark", "benchmarks", "context window", "cost", "deepseek", "fine-tuning", "glm", "grok"]
source: docs/RAG/clean_en/mindstudio/ox-alpha-free-on-opencode-pricing-limits-and-how-long-it-lasts.md
source_anchor: ""
source_lines: [1, 64]
sha256: bec385ec5c812331947b4400ca763a67af0d49954f009f93a9b52e60d3d952a0
---

# ox-alpha-free-on-opencode-pricing-limits-and-how-long-it-lasts

<!-- source: https://www.mindstudio.ai/blog/ox-alpha-free-access-opencode -->

## What is Ox Alpha?

Ox Alpha is an anonymous “stealth” model that appeared on OpenCode with no confirmed provider attached. OpenCode lists it as free to use, with a 1 million token context window, multimodal support, zero data retention, and a stated daily capacity of 100 trillion tokens. Nobody has officially claimed it, but early testing suggests it performs at a frontier level, and community investigation points toward it being connected to Zhipu AI’s GLM model family.

## TL;DR

- **Free access window** : Ox Alpha is available at no cost on OpenCode, with the free period reportedly ending around August 27th.
- **Context and capacity** : The model ships with a 1 million token context window and OpenCode says it has capacity for 100 trillion tokens per day, effectively removing usage friction while the promotion lasts.
- **Benchmark performance** : On one independent test suite (Kingbench), Ox Alpha scored 87.5%, placing second only to GLM 5.3 and ahead of models like Opus 4.8 and Qwen 3.8 Max.
- **Agentic coding strength** : On a 10-task subset of Deep SWE, Ox Alpha scored around 80%, beating GLM 5.3, Grok 4.6, and GPT 5.6 Sol on the same tasks by a wide margin.
- **Likely origin** : Fingerprinting analysis of video token counts, tokenizer behavior, and writing style points toward this being a next-generation GLM model from Zhipu AI, though this remains unconfirmed.
- **Pattern of stealth releases** : This is reportedly the fifth stealth model release in a series, and every prior one was eventually claimed by a Chinese AI lab.

## How long is Ox Alpha free on OpenCode?

OpenCode has stated the free access period runs for about a week from launch, with the cutoff landing around August 27th. After that, the model is expected to either get formally revealed and priced, or disappear if it was purely a testing exercise. This is a common pattern for stealth model drops: labs use free access windows to gather real-world usage data and community feedback before committing to a public launch and pricing structure. If you want to test Ox Alpha’s coding or reasoning capabilities firsthand, the free window is the only guaranteed opportunity, since nothing has been confirmed about pricing or availability after the reveal.

## What kind of usage limits does it actually have?

The headline number here is the 100 trillion token daily capacity figure that OpenCode has publicized. In practice, that number is less about a hard cap you’re likely to bump into and more a signal that the free tier is not artificially throttled the way many trial periods are. Combined with the 1 million token context window, this means users can run large codebases, long documents, or extended agent sessions through the model without hitting the kind of aggressive rate limiting typically seen on free tiers. Zero data retention is also part of the offering, which matters for teams testing the model on proprietary code or sensitive data during the trial.

## How does Ox Alpha perform on benchmarks?

Independent testing gives a clearer picture than the marketing language around “free” and “unlimited.” On Kingbench, a benchmark suite covering coding, math, and generation tasks, Ox Alpha scored 70 out of 80 (87.5%). That places it second on the leaderboard behind GLM 5.3 (91.25%), and ahead of Kimi K2 variants, Qwen 3.8 Max (81.25%), and Opus 4.8 (80%). It posted perfect scores on tasks like a Three.js contact lens rendering case, an SVG generation task, a hard math permutation problem, and a Gemma fine-tuning task, the last of which is notoriously difficult for most models to handle cleanly.

Separately, developer Ben Davis ran Ox Alpha through a 10-task subset of Deep SWE, a benchmark focused on software engineering tasks. Ox Alpha scored 80%, described as a near-miss on the tasks it didn’t fully pass, meaning the real score may be slightly higher. For comparison, GLM 5.3 scored 62%, Grok 4.6 scored 62%, and GPT 5.6 Sol scored 52% on the same subset. Ox Alpha also solved a specific “Marriott task” in a single attempt where GLM 5.3, GPT 5.6 Sol, and Grok 4.6 all failed on four separate attempts. It’s worth noting this is a small 10-task sample, so the margin could narrow with broader testing, but the gap observed is large enough to be notable.

## Who actually made Ox Alpha?

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

No lab has officially claimed the model, but circumstantial evidence points strongly toward Zhipu AI’s GLM line. The investigation by Ben Davis, cited at around 90% confidence, rests on several technical fingerprints. Video encoding tests showed Ox Alpha’s token counts matching GLM 5V Turbo exactly across multiple scenarios, including frame sampling rate, duration scaling (around 147 tokens per second), and per-frame resolution scaling. Other model providers tested showed different signatures entirely.

Tokenizer counts matched GLM 5.3 exactly across 25 different prompts, which generally requires the same underlying vocabulary. The writing style, including emoji-heavy responses, is consistent with patterns seen in GLM and Qwen model families. Ox Alpha also rejects audio input, matching GLM 5V’s behavior and ruling out models like MiMo that do support audio. This process of elimination also excluded DeepSeek, Qwen, Xiaomi, and Western labs based on the same fingerprinting approach. None of this constitutes official confirmation, and it’s possible some of these signals are misleading or intentionally obscured, but the pattern is consistent with previous stealth releases in this series, all of which were eventually claimed by Chinese AI labs.

## Is Ox Alpha worth trying while it’s free?

For anyone building with AI models, yes, mostly because there’s no real downside during the free window. You get a 1 million token context, strong benchmark results across both one-shot generation and agentic coding tasks, and no data retention concerns, all without a subscription commitment. The interesting wrinkle is that Ox Alpha scored slightly below GLM 5.3 on general benchmarks like Kingbench but notably above it on agentic coding tasks like Deep SWE. If Ox Alpha is indeed a next-generation GLM checkpoint, this split suggests a model tuned more heavily toward agentic and coding workflows rather than general-purpose one-shot generation. That’s a meaningfully different use case than a straight upgrade, and worth considering if your primary interest is coding assistance rather than open-ended tasks.

## Frequently Asked Questions

### What is Ox Alpha’s context window size?

Ox Alpha has a 1 million token context window, according to OpenCode’s listing, along with multimodal support and zero data retention during the free access period.

### When does free access to Ox Alpha end?

The free period is expected to end around August 27th, based on OpenCode’s stated one-week trial window from launch.

### Is Ox Alpha confirmed to be a GLM model?

No. It remains unconfirmed. Independent fingerprinting analysis of video encoding, tokenizer behavior, and response style points toward a next-generation GLM model from Zhipu AI with roughly 90% confidence, but no official lab has claimed the model.

### How does Ox Alpha compare to other frontier models?

On one benchmark suite, it scored second overall behind GLM 5.3, ahead of models including Opus 4.8 and Qwen 3.8 Max. On a coding-focused benchmark subset, it outperformed GLM 5.3, Grok 4.6, and GPT 5.6 Sol by a wide margin.

### What happens to Ox Alpha after the free period ends?

That depends on whether the model gets an official reveal and pricing structure, which has happened with previous stealth releases in this same series. Nothing has been confirmed yet about post-trial pricing or availability.
