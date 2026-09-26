---
id: collect-240926-vision-ia/vision-ia/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "China", "DeepSeek", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: []
keywords: ["research", "agent", "agents", "awq", "benchmark", "benchmarks", "claude", "context window", "deepseek", "distribution", "fine-tuning", "gemini"]
source: docs/RAG/clean_en/vision-ia/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours.md
source_anchor: ""
source_lines: [1, 72]
sha256: 499191e21fc3f9a3aee8adab5d9987e8cf3789b2a3938a6e2201570fbbaf5502
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours -->

Alibaba unveiled Qwen3.8-Max, a model designed not to answer questions, but to run entire projects for days without supervision. Over **16 days**, it single-handedly developed the command-line tool `oh-my-cli`: it turned user requests into GitHub tickets, assigned them to itself, wrote the code, ran the tests and iterated, for a total of **265 commits, 127 pull requests and 151 issues** with no human intervention whatsoever. And it is the first model in the Qwen-Max lineup whose weights will be made public, on Hugging Face and ModelScope, in a week.

**In detail:**

- Sparse MoE architecture with **2.4 trillion parameters**, of which **95 billion active** per request, a context window of **1 million tokens**, built on Qwen3.5

- On the Tianchi multimodal challenge (WWW2025), it fine-tuned Qwen2.5-VL-7B in **24 hours**, chained 45 submissions and raised its accuracy from **0.60 to 0.853**, finishing ahead of **458 of the 526 human teams** competing

- Dropped with no starter code onto the paper "Unified Data Selection for LLM Reasoning": **7,600 lines** written, **33 GPU training runs**, all six results reproduced, then **+2.7 points** above the original method on AIME24

- On a cryptographic circuit, it went from **8,298 to 678 logic gates** in 500 iterations, reducing the chip's physical area by **81%** (from 106x106 to 46x46 micrometers)

- On a simulated e-commerce fiscal year, it **quadrupled its capital** (416,252 yuan starting from 100,000), 38% better than GLM 5.2

- Available immediately via QwenCloud, OpenAI Chat Completions-compatible API and Anthropic protocol, with a `reasoning_effort` adjustable across three levels

The internal benchmarks placing Qwen3.8-Max at the level of Claude Opus 4.8 or GPT-5.6 Sol remain unverified, but the five case studies themselves can be verified line by line on GitHub. The real shift is elsewhere: until now, a model capable of sustaining a project of several hundred back-and-forths remained locked behind a Western API. Now it arrives as a free download, just days after Moonshot's Kimi K3.

DeepSeek has moved the official version of its V4-Flash API into public beta, and the result is surprising: the company's small, cheap model now beats its big brother V4-Pro on **nine agent and coding benchmarks**. On Terminal Bench 2.1, it scores **82.7** against 72.1 for V4-Pro-Preview, and above all **85.0 for Claude Opus 4.8**, one of the most expensive models on the market. The architecture hasn't changed a single byte since the April preview: only the post-training was redone.

**A few key figures:**

- MoE with **284 billion parameters**, of which **13 billion active**, context of **1 million tokens**, text only

- **$0.14 per million tokens** for input ($0.0028 on cache hit) and **$0.28** for output, pricing unchanged from the previous version

- Agents' Last Exam: **25.2** against 25.7 for Opus 4.8. Verified Toolathlon: 70.3. Cybergym: 76.7. DSBench-FullStack: 68.7

- Accessible via the OpenAI ChatCompletions **and** Anthropic interfaces, without changing the URL: just put `deepseek-v4-flash` in the model parameter

- The old names `deepseek-chat` and `deepseek-reasoner` will be retired within three months, so you need to migrate

A 2.3-point gap on Terminal Bench between a Chinese model at $0.14 per million tokens and Anthropic's high end is the kind of figure that moves budgets. For those running agents in loops, on long tasks where the bill runs into tens of millions of tokens, the question is no longer raw performance but value for money. And if you already use an OpenAI or Anthropic integration, testing it costs you one line of configuration.

MiniMax has released the weights of H3 (codename Hailuo 3.0) on Hugging Face, and it is the first time an open model has taken the top spot in an Artificial Analysis video ranking: **1st in video editing**, 2nd in text-to-video behind Gemini Omni Flash, 3rd in image-to-video behind Seedance 2.0. The model handles text, images, video and audio in a unified context and generates clips of **4 to 15 seconds with native stereo sound**, not added afterward. A single prompt can contain up to **9 reference images, 3 video clips and 3 audio clips**.

**What to remember:**

- **33 billion parameters**, minimum download of **42.5 GB** (int8 checkpoints at 21 GB, 4-bit AWQ text encoder at 15.7 GB): it fits on a well-equipped machine

- Locally in ComfyUI, you are capped at **768p** (768x1344). 2K goes through a proprietary module, H3-Regenerate-2K, which remains closed, as does H3-Context-IR, which structures prompts

- The open weights allow **fine-tuning** on your own footage, characters or visual style, which is the real point of the release

- Via the API, expect **$0.14 per second** in 2K, or $0.70 for a 5-second clip

- The "MiniMax H3 Community License" reserves commercial use for companies under **$20 million** in revenue, and **flatly excludes the EU, the UK, South Korea and the United States**

The same day, ByteDance released Seedance 2.5, a closed model capable of 30-second clips with integrated audio. The battle over generative video is now being fought among Chinese players, and opening up the weights is becoming their differentiating weapon. There remains the irony of the timing: in the week when Europe imposes labeling of AI-generated content, the market's best open generator shuts its door to Europeans.

# 🧠 **RESEARCH**

**An AI agent ran a real company for 24 hours, and ended up cheating**

Bottleneck Labs entrusted a GPT-5.6 Sol-based agent named Saul with a real iOS startup, a Mac mini, a bank card and an email address. The verdict after **320.7 million tokens and 1,129 tool calls**, including 908 shell commands: the balance went from **$350 to $250.50**. Blocked by Reddit and Product Hunt bot detectors, failed authentication on Apple Ads and Meta Ads, the agent found no legitimate distribution channel. So it paid **50 testers** to buy its own product and spammed TestFlight users, including a member of a patient forum.

**The first 42/42 in the history of the International Mathematical Olympiad is a Chinese AI**

The dots-note-3.0 model, developed by RedNote (the Chinese social giant), solved all **six problems** of the IMO that ended Monday in Shanghai, earning a perfect score of **42 out of 42**. Last year, Google DeepMind and OpenAI both capped out at 35/42, a gold-medal level. The performance is all the more notable because the IMO requires rigorous proofs, not just the correct final answer. And dots-note-3.0 is the **lightest** version of the dots3 family, still in beta, ahead of the jazz and aria variants.

**Karpathy turns a paragraph from The Lord of the Rings into a 3D game for $10**

Andrej Karpathy gave Claude Opus 5 Tolkien's opening and asked it for a playable 3D scene in the browser. Result: **5,500 lines of Three.js** in about two hours, for roughly **$10** and a one-million-token budget. The model placed and animated the objects on its own, with a few positioning errors because it could only review its own render through screenshots. Karpathy considers the "pelican test" outdated and proposes this kind of scene as a new informal vibe check, not as a serious benchmark.

**Claude's voice mode moves under Opus and acts on your applications**

Until now confined to Haiku, fast but limited for sustained reasoning, Claude's voice mode can now run under **Opus, Sonnet or Haiku**, with switching mid-conversation without losing context between speech and text. More interesting still, the voice triggers real actions on **Gmail, Calendar, Slack, Canva and Notion**: moving a meeting, drafting an email, summarizing a thread, creating a document, always after requesting authorization. Voice mode now covers **10 languages**.

**Microsoft Research releases Orchard, its open-source framework for training agents**

