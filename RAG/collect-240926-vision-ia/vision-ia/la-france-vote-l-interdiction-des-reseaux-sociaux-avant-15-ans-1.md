---
id: collect-240926-vision-ia/vision-ia/la-france-vote-l-interdiction-des-reseaux-sociaux-avant-15-ans-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "EU", "Google", "Microsoft", "Mistral", "Nvidia", "Xiaomi"]
dates: ["2026-03", "2026-06"]
keywords: ["research", "agent", "agents", "apache", "benchmark", "claude", "copilot", "cyber", "foundry", "funding", "funding round", "gemini"]
source: docs/RAG/clean_en/vision-ia/la-france-vote-l-interdiction-des-reseaux-sociaux-avant-15-ans.md
source_anchor: ""
source_lines: [1, 66]
sha256: a46a9fae5cba73fb9a1121a1a0e509676c1e136c37b9e5bd34df9a753496eb3a
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/la-france-vote-l-interdiction-des-r-seaux-sociaux-avant-15-ans -->

The entire tech press is talking about it: Google DeepMind launched on **July 21** not one, but **three new Gemini models**. The throughline announced by the lab: making AI agents **faster and cheaper** to run in production. The flagship model, Gemini 3.6 Flash, consumes **17% fewer tokens** than the previous generation, at equal quality.

**Key takeaways:**

- **Gemini 3.6 Flash**, the "all-purpose" model: better at code (**49% on DeepSWE vs. 37%**) and data analysis (**63.9% on MLE Bench vs. 49.7%**), with knowledge up to date as of **March 2026**. Price: **$1.50 per million input tokens, $7.50 for output** (down from $9 previously).

- **Gemini 3.5 Flash-Lite**, the fastest and most economical: **350 tokens per second**, built for high volumes and document processing, at **$0.30 / $2.50** per million tokens.

- **Gemini 3.5 Flash Cyber**, paired with the CodeMender agent: it detects and fixes security vulnerabilities by orchestrating multiple agents together, but remains reserved for governments and carefully vetted partners (pilot program, not for the general public).

- Everything is available right now on the Gemini API, Google AI Studio, the Gemini app, and Google Search.

**What this changes:**

For a company running AI agents continuously, fewer tokens and a lower price per token translate directly into a lighter bill. A stinging detail: Google is multiplying efficient Flash models while its flagship model, **Gemini 3.5 Pro, remains stuck** due to code issues and still only runs with partners. In parallel, the lab confirms it has started "its most ambitious training run to date," that of **Gemini 4**.

Alibaba's Qwen team unveiled on **July 21** Qwen-Image-3.0, an image generator that accepts prompts up to **4,500 tokens** and renders legible text **from just ten pixels tall**, in **twelve languages** natively. In a single pass, it produces complete infographics, newspaper pages, or multi-line LaTeX equations.

**In detail:**

- Prompts supported up to **4,500 tokens**, native rendering in **12 languages** (including Japanese, Korean, Spanish) and more than **20 character fonts**.

- Generates dense layouts in one go: infographics, nested interface grids, complex mathematical notation with subscripts and superscripts.

- Image editing pushed to the point of restoring traditional ink paintings, and photographic details on portraits (pores, skin texture).

- Designed as a content production tool: e-commerce mockups, web-fiction storyboards, press layouts.

**The caveat:**

Alibaba published **no benchmark, no spec sheet, nor the number of parameters** of the model, and the weights will probably not be open (unlike the first Qwen-Image). Access is currently via API and invitation, with an arrival planned "soon" in Qwen Chat. In other words, a fine display of strength on the rendering side, but more marketing than scientific communication against GPT-Image-2 and Nano Banana Pro. For the user, the usefulness remains nuanced: the output is a pixel image, not an editable file, whereas a researcher or graphic designer will want to regain control over their document.

Microsoft and Mistral are expanding their partnership to deploy AI infrastructure in Europe. Beyond the billions announced, the concrete benefit for organizations: Mistral's **Medium 3.5** and **OCR 4** models are arriving on **Microsoft Foundry**, and can run on your own infrastructure, **including completely disconnected from the Internet**.

**The key points:**

- **Medium 3.5** (text and reasoning) and **OCR 4** (document reading) are now available on Microsoft Foundry, and Medium 3.5 is also available in Copilot Studio.

- Via **Azure Local**, customers can run Mistral's open models in the cloud, on their own servers, or **completely offline**.

- Stated target: **finance, healthcare, industry** and other regulated sectors, with data sovereignty as the central argument (official Microsoft press release).

- On the infrastructure side, Mistral is adding **thousands of Nvidia Vera Rubin GPUs** and has invested **1.2 billion euros** in data centers in Sweden.

**The impact to remember:**

For a European administration, firm, or hospital, the appeal is clear: use cutting-edge AI without sending its data outside the EU, or even without letting it leave its own walls. Behind the scenes, Mistral has multiplied its revenue by **twenty** (more than $400 million in annualized revenue) and is reportedly negotiating a funding round of around **3 billion euros** for a valuation close to 20 billion. Along the way, the French startup has renamed its assistant Le Chat to **Vibe**, repositioned as a productivity tool with a Work Mode and a Code Mode.

# 🧠 **RESEARCH**

**A brain implant restores independence to a paralyzed man**, Three years after a double neural bypass, patient Keith Thomas can eat, drink, and grab fragile objects without breaking them in **90% of cases**. Micro-electrodes read his movement intention with **about 85% accuracy**, and pressure sensors recreate the sensation of touch by connecting his hand to his sensory cortex. Study published in *Nature Medicine* by the Feinstein Institutes.

**China wants to deflect an asteroid with a Mach 26 projectile**, Beijing is preparing a planetary defense test: striking a **30-meter** asteroid with an impactor launched at **more than 9 km/s**, nearly **50% faster** than NASA's Dart mission (6.1 km/s). The goal, led by Li Mingtao's team, is to alter its orbit, "or even break its structure," during a passage planned for **2029 or 2030**.

**Xiaomi-Robotics-1: more data beats bigger models**, Xiaomi trained its robot on **more than 100,000 hours** of movements captured by people equipped with camera-mounted hand grippers, without a real robot. Counterintuitive result: increasing the volume of data improves performance far more than enlarging the model, and the gains have not yet plateaued, even though absolute success rates remain low.

**An AI returns $38.50 per dollar invested to Pakistan's judiciary**, Across **1,559 Pakistani judges**, the JudgeGPT assistant raised the case resolution rate by **6.3%**. The catch: the gain appears only among judges **trained on the tool**, and almost disappears without that training. The researchers estimate a return on investment of up to **$38.50 per dollar** spent, a reminder that the tool is not enough without human support.

**Fractal: an open-source tool that makes AI agents self-organize**, Plasma AI has released under the **Apache 2.0** license (free even for commercial use) Fractal, a command-line tool that turns a Claude Code or Codex session into a **persistent tree of agents** capable of creating others. It runs **entirely locally**, with no remote server, to break down large coding tasks that do not fit within a single agent's context.

**Alibaba's Qwen Audio 3.0 TTS takes the lead in synthetic voices**, Alibaba's new speech synthesis model climbs to **number 1** on Artificial Analysis's Speech Arena. It handles **16 languages** and allows control of voice style in natural language or via tags such as [angry], but remains slower than its rivals Sonic 3.5 and Simba 3.2, at **16 characters per second**.

**Half of the tracks uploaded to Deezer are AI-generated**, Deezer claims that **more than 50%** of the titles submitted each day to its platform are AI-produced, a peak of **90,000 tracks per day** reached in June 2026 (compared with 10% in early 2025). The platform will remove AI tracks never listened to for six months or linked to fraudulent streaming, to protect artists' revenue. Its technology notably identifies tracks originating from Suno and Udio.

