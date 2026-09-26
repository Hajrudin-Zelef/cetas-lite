---
id: collect-240926-vision-ia/vision-ia/google-prend-du-retard-face-a-la-concurrence-gemini-3-5-pro-repousse-une-3e-fois-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "California", "EU", "Google", "Hugging Face", "Meta", "Nvidia", "OpenAI", "Perplexity", "Samsung", "Z.ai", "xAI"]
dates: ["2027-01", "2027-07"]
keywords: ["research", "agent", "alignment", "apache", "benchmark", "chatgpt", "claude", "compute", "consumer", "cost", "cybersecurity", "diffusion"]
source: docs/RAG/clean_en/vision-ia/google-prend-du-retard-face-a-la-concurrence-gemini-3-5-pro-repousse-une-3e-fois.md
source_anchor: ""
source_lines: [93, 170]
sha256: b81b8d54562bb37a2d016d6209ec3d0667c7d63985b6c10e473ea8932dd40019
---

# 🧠 **RESEARCH**

Colibrì is an open-source runtime that runs **GLM-5.2**, a MoE model with **744 billion** parameters, on consumer hardware. The trick: each token activates only about **40 billion** parameters. The runtime keeps the dense tensors quantized in int4 (~9.9 GB) in RAM and leaves the **experts** (~370 GB) on an NVMe SSD, loaded on demand. It's **slow, but it works**, and it puts a cutting-edge model within reach of a beefy PC.

**Grok 4.5 beats Claude Opus 4.8 as an orchestrator, for half the price**

On the **WANDR** benchmark (complex multi-step research tasks) within Perplexity Computer, **Grok 4.5** earns the best score (**0.328**) for **$4.76** per run. It beats **Opus 4.8** (0.254 for $9.46) and **GPT-5.6** (0.289 for $2.64). A good reminder that, for multi-agent orchestration, the best performance-to-price ratio doesn't always come from the most prestigious model.

**NVIDIA and Hugging Face open up fine-tuning of video models at scale**

The two players are releasing an **open-source** integration (Apache 2.0) between **NeMo Automodel** and **Diffusers**. You can now train and fine-tune image and video diffusion models (**FLUX.1-dev, Wan 2.1, HunyuanVideo**) without checkpoint conversion, and go from a single GPU to hundreds by changing a single line of configuration. Enough to democratize the training of generative video models.

**Plan an entire weekend by talking to GPT-Live**

This guide shows how to use **GPT-Live**, the new voice model behind ChatGPT Voice, to organize a trip out loud: you interrupt it, redirect its searches in real time, then turn the conversation into an action plan (accommodation, itinerary, order of bookings). A six-step method. **GPT-Live-1** is reserved for paid accounts; a mini version is offered to free users.


# **🗞️MORE NEWS**

**xAI open-sources its Grok Build coding agent**

xAI has published under the **Apache 2.0** license the **Rust** code for Grok Build, its terminal coding agent: agent loop, file tools, shell execution, web search, interface. You can compile it yourself and plug in **local inference** via config.toml. The repository has already passed **1,900 stars**, even though external contributions remain closed.

**Amazon Zoox recalls 105 robotaxis after driving into thick smoke**

A Zoox robotaxi with no passengers drove into a **smoky fire scene** in Las Vegas on June 20 before braking and stopping. The Amazon subsidiary is recalling **105** vehicles to fix the software bug that prevented them from detecting dense smoke. The NHTSA has just required autonomous vehicle makers to resolve interference with emergency responders by the end of July.

**Strike at Hyundai against humanoid robots**

Thousands of workers at the Ulsan plant (South Korea) walked off the job after negotiations failed over the deployment of Boston Dynamics' **Atlas** robots. It's the first automotive work stoppage linked to fear of humanoids: Hyundai wants to deploy **more than 25,000** of them. Each Atlas costs about **$130,000**, recoupable in two years, with an operating cost soon to be lower than the U.S. federal minimum wage.

**Kaiser nurses denounce an AI that grades their empathy**

At Kaiser Permanente, AI software measures call duration and evaluates everything down to nurses' **tone of voice** and empathy. Beyond **15 minutes** of conversation, even for a suicidal patient, punishment looms. The union is negotiating a contract for 25,000 caregivers while California studies protective legislation. (495 points on Hacker News.)

**TikTok tests an AI-generated likeness detection tool**

TikTok is experimenting with an **opt-in** tool that lets creators flag videos exploiting their appearance without authorization. The test is limited to certain American creators, with identity verification (real-time selfie plus ID via **Jumio**). YouTube already offers a similar system to all its adult users.

**Patreon now actively blocks AI bots**

No more politely ignored robots.txt file: Patreon relies on **Cloudflare** (AI Crawl Control) to block crawlers that train AI on creators' work. Weekly access attempts reportedly dropped from **several thousand to zero**. Indexing bots that send traffic back, however, remain allowed.

**Claude Code: anatomy of a failed feature**

Anthropic had slipped into Claude Code (v2.1.198) a **60-second timer** after which the agent continues on its own if the human doesn't respond, without mentioning it in the changelog. The author points to the lack of transparency and the risk of unsupervised actions. Anthropic fixed the behavior a few days after it came to light. (139 points on Hacker News.)

**AI data centers' hunger for memory is driving up smartphone prices**

Samsung, SK Hynix and Micron are redirecting their production toward the **HBM** memory used in AI accelerators, which is far more profitable. The result in India, the world's 2nd market: smartphone shipments fell **10%** in the 2nd quarter, the worst drop in six years. The sub-$150 segment is collapsing by **45%**, and people are keeping their phones longer.

**The EU forces Google to open Android to rival AI assistants**

In the name of the **DMA**, the European Commission is ordering Google to open **11 Android functions** (voice search, bookings) to competing AI assistants by July 2027, and to share anonymized search data as early as January 2027. Google warns that these rules could create privacy and cybersecurity risks.

**The Pentagon: slow AI adoption is riskier than imperfect alignment**

The U.S. Navy signs a strategy for an **"AI-first"** fleet, with large language models running **directly on** warships and an "AI war council" to prioritize mission scenarios. The doctrine's stated message: delaying military AI adoption would be more dangerous than deploying imperfectly aligned systems.

**OpenAI proposes a "scorecard" to measure AI's real ROI**

OpenAI's chief financial officer, **Sarah Friar**, proposes a framework to move past vague marketing: measuring the **useful work** produced, the **cost per successful task**, reliability, and the return on compute invested. The goal: give companies tangible metrics to judge whether AI truly pays off for them, rather than promises.

**Meta could rent its excess AI compute to Anthropic**

Meta is reportedly negotiating with Anthropic to lease capacity from its data centers, a deal that could reach **$10 billion** over two years. Anthropic lacks compute given demand for Claude Code; Meta, which plans up to **$145 billion** in spending this year, is looking to make its infrastructure profitable. Nothing has been signed yet.

**Apple sues OpenAI: what's the real goal?**

Apple is suing OpenAI for alleged trade secret theft, a complaint several experts consider partly exaggerated. The Vergecast wonders whether Apple fears a serious competitor or is taking advantage of a moment of weakness for OpenAI, as it simultaneously launches the public beta of its AI-powered **Siri** in iOS 27.

**Apple overtakes Nvidia and becomes the most valuable company again**

Apple and Nvidia both crossed the **$4.8 trillion** market capitalization mark on Friday, vying for first place. Apple gained about **23%** in 2026 thanks to its AI announcements, while Nvidia is up only 9%, with Wall Street now turning toward memory manufacturers like Micron.

**Neil Rimer (Index Ventures): AI money will eventually flow back out**

The co-founder of Index Ventures (an Anthropic investor) predicts a **redistribution**, voluntary or not, of the wealth generated by AI. He also points to the decline of tech philanthropy: the Giving Pledge went from **113 signatories** in its first year to only **4 in 2024**.

Every week, a new AI tool makes a skill obsolete. Those who know how to use these tools save time, money and clients. Those who don't, watch others do it.
