---
id: collect-240926-mindstudio/mindstudio/glm-5-3-flash-runs-on-chinese-chips-without-nvidia-1
title: "glm-5-3-flash-runs-on-chinese-chips-without-nvidia"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "China", "Google", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["glm", "nvidia", "agent", "agentic", "agents", "benchmark", "benchmarks", "claude", "cost", "cost per token", "gemini", "gpt-5.6"]
source: docs/RAG/clean_en/mindstudio/glm-5-3-flash-runs-on-chinese-chips-without-nvidia.md
source_anchor: ""
source_lines: [1, 51]
sha256: c6fff1f2a521475131cb92557a4698577c40e86daa0347cbb039c790a7e7354d
---

# glm-5-3-flash-runs-on-chinese-chips-without-nvidia

<!-- source: https://www.mindstudio.ai/blog/china-ai-chips-glm-inference -->

## What did ZAI actually demonstrate with GLM 5.3 Flash?

ZAI, a Chinese AI lab, released GLM 5.3 Flash, an open-weights model built with a mixture-of-experts design (320 billion total parameters, 18 billion active per token). According to reporting from semi analysis cited by the source video, ZAI served more than 100 trillion tokens per day of this model running entirely on Chinese-made AI chips, with no Nvidia hardware in the stack. That’s the headline: a near-frontier-capable model, at real production volume, running on domestic silicon instead of the GPUs that have dominated AI infrastructure for a decade.

This matters because inference at scale has quietly become the bottleneck that determines who can actually deploy AI cheaply and widely. Training a model once is one problem. Serving it to millions of users, every day, at a cost per token low enough to run a business on, is a different and arguably harder problem. ZAI’s claim, if accurate, is that Chinese chip stacks can now handle that harder problem for a model that’s competitive with mainstream frontier systems on coding and agentic tasks.

## TL;DR

- **ZAI’s GLM 5.3 Flash** is a 320 billion parameter mixture-of-experts model with 18 billion active parameters, released as open weights, and it reportedly ran on Chinese chips serving over 100 trillion tokens per day.
- **The cost story is the real headline** : GLM 5.3 Flash reportedly costs around 9 cents per completed task on the artificial analysis intelligence index, compared to roughly $3.14 for Claude Opus 4.5, while landing only a few points behind it on the intelligence index.
- **Hardware and model co-design appears to be the key enabler** , with ZAI describing a serving stack, interconnect, and cluster all optimized together specifically for Chinese chips rather than adapted from an Nvidia-first pipeline.
- **Open weights change who gets to test this** , since anyone can run GLM 5.3 Flash through providers like OpenRouter or ZAI’s own API, and multiple inference providers will compete on price rather than one company gatekeeping access.
- **It’s not beating the biggest models outright** , since GLM 5.3 Flash sits a few points under Claude Opus 4.5 on intelligence benchmarks and used more output tokens per task than some efficient rivals like GPT-5.6 Luna, meaning the token efficiency story is more mixed than the price story.
- **The chip independence angle is bigger than one model** , because it suggests a full-stack alternative to Nvidia (chips, interconnect, serving software) is now viable for frontier-adjacent inference workloads, not just small or toy models.

## Why does running inference on Chinese chips matter?

For years, the assumption in AI infrastructure has been simple: if you want to train or serve a serious model, you need Nvidia GPUs, because the CUDA software ecosystem, the interconnect technology, and the raw chip performance had no real substitute at scale. Export restrictions on advanced chips to China were built on that assumption, betting that cutting off access to Nvidia’s best hardware would meaningfully slow down Chinese AI development.

GLM 5.3 Flash’s reported serving numbers push back on that bet. ZAI describes running the model on “a large-scale cluster of Chinese AI chips supported by high bandwidth interconnect and a serving stack optimized for the underlying hardware.” The framing that matters here is co-design: the chips, the interconnect, and the software serving layer were built together, specifically for this hardware, rather than porting an Nvidia-oriented stack onto substitute silicon. ZAI’s own claim is that this setup achieves “hardware efficiency and per token cost comparable to mainstream Nvidia GPUs.”

If that holds up under independent scrutiny, it means the constraint isn’t just “can China make a chip that’s fast.” It’s “can China make an entire full-stack alternative that’s economically competitive.” Serving over 100 trillion tokens a day is not a lab demo. That’s production-scale traffic, which is a much stronger signal than a benchmark chip test in isolation.

## How does GLM 5.3 Flash perform against frontier models?

On benchmark comparisons against similarly efficient models, GLM 5.3 Flash scored 84.3 on Terminal Bench, ahead of the earlier full-size GLM 5.2, and landed close to GPT-5.6 Tera and Gemini 3.7 Flash. On Deep Suite, considered a more accurate reflection of real-world usage quality, it scored 63.4, a sizable jump from GLM 5.2 and competitive with GPT-5.6 Tera. On GDPval, an OpenAI benchmark for real-world knowledge work, it reportedly ranked first by a large margin among the models compared.

Against the very largest frontier models, the gap is real but narrow. On the artificial analysis intelligence index, GLM 5.3 Flash scored 57 versus Claude Opus 4.5’s 62. That’s a meaningful difference, but the context matters: Opus 4.5 is estimated in the multi-trillion parameter range, while GLM 5.3 Flash runs at 320 billion total parameters with only 18 billion active. Closing most of that gap at a fraction of the size is the notable part, not the fact that a gap still exists.

## Is GLM 5.3 Flash actually cheaper to run?

Cost per completed task is where the numbers get striking. Claude Opus 4.5 reportedly costs around $3.14 per task on the intelligence index measure. GPT-5.6 Soul comes in around 95 cents. Kimi K3 sits at 84 cents. GLM 5.3 Flash reportedly costs about 9 cents per task, roughly 2 to 3 percent of Opus 4.5’s cost while landing only a handful of points behind it on raw intelligence score.

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

That said, cheapest isn’t the same as most efficient. GPT-5.6 Luna Max reportedly costs about half of GLM 5.3 Flash’s price while scoring only a few points lower on the intelligence index, making it arguably the better cost-to-quality tradeoff of the two, even though both are inexpensive in absolute terms. Part of the reason: GLM 5.3 Flash reportedly uses around 47,000 output tokens on average to complete tasks on the intelligence benchmark, more than double what Luna Max uses for a comparable result. So GLM 5.3 Flash is cheap partly because its per-token price is low, not because it’s stingy with tokens. A more token-efficient future version could close that gap further, and open weights mean outside developers and rival labs can try to do exactly that.

## What does open-weights availability mean in practice?

Because ZAI released GLM 5.3 Flash as open weights, it isn’t locked to one API or one company’s pricing. It’s already available through ZAI’s own API and through OpenRouter, and it can be run through any tool that supports an OpenAI-compatible endpoint. That includes coding tools and agent frameworks that developers already use for other models. Multiple inference providers can host it and compete on price, which historically pushes serving costs down faster than a single vendor controlling access would.

Independent testing so far has been mixed by task type. In coding and agentic benchmarks, GLM 5.3 Flash performs close to top frontier models relative to its size. In more open-ended generation tasks, like producing polished visual demos or fully designed websites, comparisons against GPT-5.6 Soul in informal testing showed a split result: Soul produced more polished 3D scene renders, while GLM 5.3 Flash produced a cleaner, better laid-out website in at least one side-by-side test, notably without internet access that Soul had available during that comparison.

## Frequently Asked Questions

### What is GLM 5.3 Flash?

