---
id: collect-240926-korben/korben/glm-5-2-le-premier-moda-le-ia-open-source-que-je-garde-korben
title: "GLM 5.2 - The first open source AI model I'm keeping"
domain: korben
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "Moonshot", "Z.ai"]
dates: []
keywords: ["glm", "open source", "benchmark", "claude", "context window", "deepseek", "fable 5", "kimi", "leaderboard", "license", "llama", "mit license"]
source: docs/RAG/clean_en/korben/glm-5-2-le-premier-moda-le-ia-open-source-que-je-garde-korben.md
source_anchor: ""
source_lines: [1, 56]
sha256: 94cfa990c053804f3df94caed2fcc3507765bea5e1c90c95842303dc680311fa
---

# GLM 5.2 - The first open source AI model I'm keeping

<!-- source: https://korben.info/glm-5-2-retour-experience.html -->

# GLM 5.2 - The first open source AI model I'm keeping

## What to remember AI-generated summary

1. GLM 5.2, the open weights model from Z.ai (744 billion parameters in MoE, 1 million tokens of context), works without bugs for the author's coding uses, unlike the other open source models tested.
2. GLM 5.2 integrates directly into Claude Code via Z.ai's Anthropic-compatible API, allowing the use of the same skills and scripts as with Claude.
3. On the Arena.ai leaderboard dedicated to front-end code, GLM 5.2 comes in second place, the first open weights model at this level (all the others in the top are proprietary).

Friends, I need to tell you about GLM 5.2 . I'm using it right now through Z.ai, and it's the first time an open weights model has given me satisfaction with what I ask it to do. And God knows I've tested these fucking models!

GLM 5.2 is the latest from Z.ai, the Chinese lab formerly known as Zhipu AI. It came out this month (in June), and it's a big baby with its 744 billion parameters in Mixture-of-Experts (MoE), of which roughly 40 billion activate for each token, as well as a context window that goes up to 1 million tokens via the glm-5.2[1m] variant. All published, as always, under MIT license, with the weights downloadable on HuggingFace.

Anyway, I didn't really believe in it, but I still took the small Z.ai subscription and launched my usual tools and coded a few new features on my software. And surprise, it does very very well for my uses (I do mean for my uses!). I had no bugs, no endless discussion beating around the bush, nor a conversation ending that drifts into Chinese characters like Qwen often did to me.

Then, the cool thing is that I plugged it directly into Claude Code. If you're interested, I made myself a little specific launcher. It's a gift:

```
#!/usr/bin/env bash
export ANTHROPIC_BASE_URL="https://api.z.ai/api/anthropic"
export ANTHROPIC_AUTH_TOKEN=YOUR_API_KEY
export ANTHROPIC_DEFAULT_SONNET_MODEL="glm-5.2[1m]"
export ANTHROPIC_DEFAULT_OPUS_MODEL="glm-5.2[1m]"
export CLAUDE_CODE_AUTO_COMPACT_WINDOW="1000000"
claude "$@"
```
You save it under the name of your choice, for example "glm". Then you do a:

```
chmod +x glm
```
And then you launch it like this:

```
./glm
```
The idea is that since Z.ai's API is Anthropic-compatible, you just point Claude Code to their endpoint, slip in your key, and it talks to GLM 5.2 like it would talk to Claude. My skills, my scripts, everything works the same, it's fire!

I only regret one thing, which is not being able to run it locally at home. Because the beast is TOO big. Even trimmed down and quantized in 2-bit for home use , it eats up around 240 GB of RAM. At home, I don't have the hardware, and you probably don't either. So for now, the API is the only realistic and affordable entry point.

Whether it's Qwen, Llama, Kimi, DeepSeek, no matter what I've tested locally, for my somewhat elaborate uses, every time I'm super disappointed. So this one, for what I ask of it, holds up very well.

Now, I'm not going to sell it to you as a Claude Killer either but I still found a benchmark that confirms my feeling. On the Arena.ai leaderboard dedicated to front-end code, GLM 5.2 sits in second place, just behind Fable 5. And since everything ahead of it is proprietary, that makes it the first open weights model at that level of the ranking.

So it's not the best AI in the world, mind you, but it's the first open source one that gives me a result that suits me. And you all know how annoying and demanding I am with this kind of tool. In any case, it's the first time I've told myself that open source AI could really enter my daily flow, and not just remain a toy for ranking stuff or making slop on SEO blogs. Now, between us, I'm mainly waiting for Fable 5, or its equivalent, to come back and set things on fire!!

If you're tempted to try it, there's therefore the GLM Coding Plan from Z.ai, which starts at 18 dollars per month and is especially tailored for code. It plugs into Claude Code, Cline and about twenty tools of the same ilk. A little tip along the way, this link to the GLM Plan is indeed an affiliate link, but it gives you 10% off if you use it, and it gives me a little something too, so everyone wins.

There you go, if you've been coding with something else until now, it's worth taking a look out of curiosity.

Source: Z.ai

## Comments

starfix!in Surfshark doesn't make you invMorganein Discord guesses your agests3rv1in The Ray-Ban Display arrive eponponin Openpilot - The NHTSA passesfabiin Claude Code makes you choose
