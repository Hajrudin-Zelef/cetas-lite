---
id: collect-240926-vision-ia/vision-ia/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "ByteDance", "California", "China", "CoreWeave", "DeepSeek", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Samsung", "SpaceX", "TSMC", "United States", "Z.ai", "xAI"]
dates: []
keywords: ["research", "agent", "agentic", "agents", "amd", "attention", "awq", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/clean_en/vision-ia/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours.md
source_anchor: ""
source_lines: [1, 176]
sha256: b07779e340fbaee6f0432fcdca4265704c0988b79a3962a6507fe5cfc3be757a
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

Orchard is built on Orchard Env, a service of reusable environments that allows agents to be trained and evaluated directly in the harnesses where they will actually run: Codex, OpenClaw, ZeroClaw. The same infrastructure serves software development, web browsing and personal assistance. The Orchard-SWE model reaches **69.7% on SWE-bench Verified** (73% with value-model reranking) with only **3 billion active parameters**, ten times fewer than comparable proprietary systems. Data and training methods are published.

**Chinese military researchers are distilling American models**

A Reuters investigation covering more than **80 Chinese articles and patents** shows that researchers linked to the military are using the outputs of OpenAI and Anthropic models to train smaller systems running locally. Intended applications: surveillance, cyber operations, code analysis and tactical uses. The technique, distillation, transfers targeted capabilities, but Reuters notes that it reproduces neither the full power of a frontier model nor the compute capacity needed to train one.

**Managing an agent's memory as a lifecycle, not as storage**

This paper proposes "Agentic Context Management," a five-stage framework for deciding what an agent should **keep, retrieve, share, prepare and compress** in its context. The starting observation is concrete: agents in production fail less from a lack of intelligence than from saturation, when old messages, tool descriptions and outdated outputs clog the window. Storing more is not enough; you need rules for relevance, access, timing and compression.


# **🗞️MORE NEWS**

**Hugging Face CEO: "China clearly dominates open models"**

Clément Delangue said Monday on CNBC that China was winning the AI race, and that it could catch up to the American technological frontier as early as late 2026 or 2027. He pointed to a Chinese ecosystem of open collaboration versus American labs that "build in silos." He also revisited the hacking of Hugging Face by OpenAI agents that escaped a training environment last month, while assuring that collaboration between the two companies remains healthy. Microsoft, Palantir and Nvidia have also signed a letter calling for open-weight models not to be restricted.

**A centaur robot with a goat's head, with chainsaws instead of hands**

California startup Satyress has built Threehalves, a machine over 1m80 that combines a humanoid torso and a four-legged base braked by friction. Its arms end in a quick connector powered at 12v, 18v or 48v, capable of holding various tools, including chainsaws. The target: fire fronts, toxic industrial zones, landslides and debris searches. This is not an autonomous robot; an operator pilots it with a joystick.

**Pepper-spray drones deployed in American schools**

Florida, Georgia, and Colorado are funding pilot programs, up to **$557,000**, to equip facilities with drones designed to intercept an active shooter in about fifteen seconds. Capable of reaching **60 mph** and flying through a window, they are piloted remotely from Austin by Mithril Defense operators, and are not autonomous. School safety experts are raising alarms about the risk of misidentification and the fact that the device diverts attention from the real issues.

**AI has conquered code, now it's attacking the drive-thru**

After its catastrophic and viral debut in 2023, voice AI has become more reliable and is being deployed massively across American fast food. Taco Bell has equipped **more than 890 drive-thru lanes**, over 10% of its restaurants in the United States. Dairy Queen is deploying across **25 states** with the goal of covering its entire network, and White Castle is running its assistant "Julia" at 12% of its locations, systematically installed in new openings. McDonald's had kicked off the movement back in 2019 by acquiring Apprente.

**Anthropic removed 80% of Claude Code's system prompt, without losing anything**

An Anthropic engineer explains that the company removed **more than 80%** of Claude Code's system instructions without any degradation in code evaluations. The reason for the original length: early models needed rigid guardrails to prevent file deletions or unprofessional comments. The team knew these rules would sometimes be counterproductive, but the trade-off was worth it. With models capable of judging a situation on their own, that no longer holds.

**OpenAI unveils GPT-Live, continuous voice without turn-taking**

OpenAI has published the behind-the-scenes of GPT-Live, a continuous voice interaction system developed in six months. It relies on a so-called "turnless" speech model and a low-latency architecture, meant to eliminate the interruptions characteristic of classic voice assistants, where you have to wait for the end of a sentence to get a response. The stated goal is a conversation that feels like a real conversation.

**Fish Audio clones a voice in seconds, live**

The startup has publicly launched its flagship model S2.1 Pro, whose founders demonstrated it by cloning their own voices in real time in front of their audience. The announcement comes with a **$52 million** seed round, closing out its first year of existence. Worth noting alongside the new European labeling requirements that came into force this week.

**Gemini Spark browses Chrome with your connected accounts**

Google's agentic assistant is now integrated into the browser and can "use your connected accounts and saved passwords to handle web chores." The examples cited by Google: searching for flight options and starting a booking, or scheduling an appointment to visit an apartment. Google announces defenses against prompt injection and assures that Spark never finalizes a payment without validation. Rollout first in the United States, while Google AI Pro is now available in more than 160 countries.

**Google pulls image generation from Google Earth after 24 hours**

Launched on July 30, the Nano Banana 2 feature allowed users to generate images anchored in Google Earth's satellite, aerial, and 3D views: historical reconstructions, educational infographics, real estate projects, visualization of urban planning. By July 31, Google was already backtracking after seeing screenshots circulating that violated its rules. The images were watermarked and did not appear in the main experience. "People place a special trust in Google Earth for a reliable view of the world," the company explains.

**RAMageddon: AI is driving up the price of your next computer**

Semiconductor factories are redirecting their capacity toward everything related to AI, causing a global memory shortage that is rippling through PCs, consoles, and smartphones. Samsung has already sold its entire production through the end of 2026. Tim Cook is considering diversifying suppliers beyond TSMC, AMD anticipates **20% less gaming revenue** in the second half, and Microsoft expects a decline in PC sales linked to rising memory prices.

**MacBook Air out of stock, deliveries pushed back to late August**

A direct consequence of the shortage: the price of the MacBook Air has gone from **$1,099 to $1,299**, and an order placed today won't arrive before the end of the month, according to Mark Gurman's sources at Bloomberg. Unusually, Apple is actively steering buyers toward the 14-inch MacBook Pro on its website and in its stores, while awaiting an M6 refresh of the Air lineup.

**Apple has finally fixed Siri, and no one cares anymore**

With the public beta of iOS 27, the new Siri does what Apple promised: it understands personal context (photos, emails, calendar, contacts), draws on general knowledge, and retrieves information stored on the iPhone. It gains a dedicated app, a text mode, visual recognition through the camera, and back-and-forth conversations. The problem isn't quality—it's decent—but timing: during Apple's years of delay, the rest of the industry moved on to multi-step autonomous agents.

**AWS partners with Superblocks, and the signal is bigger than it appears**

The vibe coding startup integrates natively with AWS customers' private clouds: generated applications create Amazon Aurora databases internally and go through Amazon Bedrock, rather than sending data to external services like Supabase. AWS will market the tool through its Marketplace. The underlying trend is clear: hyperscalers are pushing companies to separate the model from the rest of the agentic infrastructure, to avoid direct dependency on labs like Anthropic or OpenAI.

**Boris Cherny: stop optimizing token costs**

The head of Claude Code at Anthropic takes a stance counter to most technical teams: "I use Fable for everything." His reasoning boils down to one line: there may be 50% savings to be scraped off the token bill, but **1,000%, 10,000% or even 100,000%** of potential gain on the return. His recommendation: take the most expensive model and ask how to get more out of it, rather than looking for where to cut back.

**The EU activates its oversight powers over Anthropic, OpenAI and Google**

Alongside transparency obligations, the European Commission can now require an evaluation of a model **before its public release** in the EU, restrict its access to the market, and impose fines of up to 3% of annual revenue. These powers apply to any company offering a general-purpose model in Europe, regardless of nationality. Brussels is already negotiating with OpenAI and Anthropic following cyberattacks involving their models, and has obtained access to Anthropic's Mythos model. The diplomatic context is tense: Donald Trump threatened the EU with "substantial" tariffs after the billion-dollar fine imposed on Google in July.

**OpenAI's luxury influencer trip turns into a PR disaster**

OpenAI took a handful of creators to an upscale retreat in New York State, dubbed "Summer Club": farm-to-table dinners, beekeeping, wellness workshops and classes on using ChatGPT Work. The videos were polished, the comments much less so: internet users criticized the creators for turning a blind eye to the environmental impact of data centers in exchange for a free stay. One of them eventually deleted his video.

**OpenAI responds publicly to Apple's complaint**

In a post bluntly titled "Apple is getting this wrong," OpenAI calls Apple's complaint baseless, corrects claims concerning its employees and publishes message exchanges to back up its version of the facts. The move comes amid already tense relations between the two companies in the AI arena.

**Palantir's CEO calls the AI industry "Marxist"**

In his letter to shareholders, Alex Karp, a doctor of social theory, writes that the builders of large language models are seeking, knowingly or not, to "capture the means of production" of their own enterprise customers. His thesis: generic AIs make it possible to harvest customers' know-how to then build competing products. Palantir positions itself as an agnostic alternative that leaves data control with companies. The quarter gives him volume: **$1.9 billion** in revenue, up **93%** year over year, and $1.1 billion in profit.

**June exits stealth with $20 million to deploy AI in the enterprise**

Founded by former executives of Bonobo AI, acquired by Salesforce, the startup raises in pre-seed from Time Ventures (Marc Benioff), Michael Dell, Aaron Levie and George Kurtz. Its platform scans a company's existing systems (Salesforce, ServiceNow, Workday) to spot duplicates and bottlenecks, then generates an agent deployment roadmap. "AI, paradoxically, increases demand for professional services," summarizes co-founder Efrat Rapoport.

**A former OpenAI researcher's fund loses 67% in one month**

Leopold Aschenbrenner, formerly of OpenAI's superalignment team and author of the essay "Situational Awareness," manages a hedge fund of the same name. In July, his 4x leveraged bets on AI infrastructure (SK Hynix, CoreWeave, Nebius, Micron, Bloom Energy) cost him **67% of its value**. The fund nonetheless remains at **+80% for the year**, after a first half of 2026 at +439%.

**Amazon crosses $3 trillion in market capitalization**

The stock closed up 4% on Monday, its best day since May, buoyed by solid quarterly results. AWS generated **$42.2 billion** in the second quarter, above expectations. The company is raising its 2026 investment budget from 200 to **$220 billion**, partly because of soaring memory prices, and judges that even at this level capacity remains insufficient against demand.

**Jim Cramer: Wall Street has finally accepted AI spending**

According to the CNBC host, it was Andy Jassy's detailed explanations of the expected return on investment that unlocked the situation, more than the amount itself. Microsoft escapes criticism because it already monetizes through Azure and Copilot. Meta remains the most attacked, for failing to explain how its spending will turn into revenue.

**SpaceX publishes its first results after a drop of more than 50%**

The publicly traded company's first quarterly report is expected Tuesday after the close, while the stock has lost more than half its value since its post-IPO peak, amounting to **more than $500 billion** in market capitalization wiped out since its first listing on June 12. The $1.4 trillion valuation rests largely on Elon Musk's promises: data centers in orbit, Mars, and an AI offensive against Google, OpenAI and Anthropic.
