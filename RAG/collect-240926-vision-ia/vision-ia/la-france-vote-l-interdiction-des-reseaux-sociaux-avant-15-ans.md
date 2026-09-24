---
id: collect-240926-vision-ia/vision-ia/la-france-vote-l-interdiction-des-reseaux-sociaux-avant-15-ans
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "China", "EU", "Google", "Intel", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Samsung", "SpaceX", "United States", "Xiaomi", "xAI"]
dates: ["2026-03", "2026-06"]
keywords: ["research", "agent", "agentic", "agents", "amd", "apache", "benchmark", "benchmarks", "blackwell", "chatgpt", "claude", "copilot"]
source: docs/RAG/clean_en/vision-ia/la-france-vote-l-interdiction-des-reseaux-sociaux-avant-15-ans.md
source_anchor: ""
source_lines: [1, 110]
sha256: 81cb1c7f2b99dcf0dc0bd4c618590cdcb69a905b3f94c090ff3e706ffc96a495
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

**Former Intel CEO wants to revive Moore's law through light**, Pat Gelsinger, now a partner at the deep tech fund Playground Capital, is betting on a new generation of lithography startups. He sits on the board of **xLight**, which is developing light-based etching and has received investment from the US government, with the ambition of surpassing current physical limits (13.5 nm at ASML) to restart the transistor race.

**Data centers will consume four times more electricity by 2035**, According to BloombergNEF, the electricity demand of data centers will **quadruple** by 2035 to reach **one fifth** of the electricity produced in the United States, driven by the explosion of AI computing. Enough to strain already fragile grids (PJM suspended new connections for four years) and drive electricity prices up by **76%** in a year.


# **🗞️MORE NEWS**

**Washington clashes over Chinese open-weight models**, The release of **Kimi K3** by Moonshot AI, the largest open-weight model ever published, has reopened a debate in Washington. Dean Ball (OpenAI) predicts an informal regulatory risk (suspicion of backdoors) rather than a ban, while David Sacks accuses the big closed labs of wanting to crush open source. A battle that will weigh on corporate purchasing well beyond the United States. Picked up by 5 sources.

**Meta's open models power the Genesis Mission**, Meta's open-source models (**Segment Anything, DINO**) are being deployed in the first scientific projects of the Genesis Mission, an initiative by the US Department of Energy carried out at Lawrence Berkeley National Laboratory, to accelerate research via computer vision. The topic is climbing on Hacker News (95 points).

**An OpenAI executive calls open models "decelerationist"**, Dean Ball, Head of Strategic Futures at OpenAI, caused a stir (**nearly 11 million views**) by commenting on Kimi K3, a Chinese multimodal model with **2.8 trillion parameters**. He acknowledges its strength in agentic coding and long context, while arguing that open models would slow AI progress rather than accelerate it.

**The US Army exhausts its AI token quotas**, US Army employees burned through their annual allocation of **100 million tokens** in a few months via the Ask Sage platform (Gemini, Llama, ChatGPT), forcing a return of usage limits one month after an announcement of unlimited access. The Pentagon is said to have consumed around **20 billion tokens per day** during a 38-day operation.

**Substack launches an AI-written article detector**, The platform is deploying a tool powered by **Pangram** that estimates the probability that a text was generated or assisted by AI. Usable on articles, notes, replies, and comments of more than 100 words via the three-dot menu, first on the web and iOS, with Android to follow.

**Claude Cowork learns skills by filming your screen**, Anthropic is adding a "Record a skill" option to Claude Cowork: you record your screen while commenting on your actions out loud, and the tool turns it all into a **reusable skill** automatically relaunched the next time. Reserved for Pro, Max, and Team subscribers. OpenAI offers a comparable feature in Codex.

**OpenAI launches a ChatGPT program for small businesses**, The new **ChatGPT for Small Business** program aims to help entrepreneurs build AI skills, automate their daily tasks and grow their business, building on the **ChatGPT Work** business offering. A way to make generative AI more accessible to small organizations.

**Meta tests StoryKit, an app for AI-generated children's tales**, The app generates personalized stories from a photo (character, setting, life lesson, music) without parents having to write a single word. Spotted in the App Store, it is confirmed by Meta as a pilot limited to a few countries, with no social features and restricted to **18 and over**.

**Samsung launches Health Assistant, a chatbot for your health data**, Integrated into Samsung Health and in beta in the United States, this chatbot cross-references sleep, activity, nutrition and vital signs to offer recommendations validated by doctors, without making a diagnosis. It relies on the Samsung ecosystem (watch, ring, connected home), on the eve of Galaxy Unpacked.

**Gritt emerges from stealth with $32M for solar panel-installing robots**, Founded by two roboticists from Carnegie Mellon, Gritt operates rented Kawasaki robotic arms to position panels with **sub-millimeter** precision. A team of 8 people installs **3,000 to 4,000 per day** versus 800 without a robot, and the startup is already under contract for **2.8 gigawatts** of solar.

**Anduril and Archer unveil "Thunder," an autonomous attack drone**, Designed as a wingman for crewed helicopters, this drone takes off vertically then tilts its rotors to fly like a fixed-wing aircraft, without a runway. The pilot assigns it objectives rather than remotely controlling it, a response to the growing threat of cheap drones in close combat.

**Bristol Myers Squibb buys an Nvidia AI supercomputer**, The pharmaceutical company becomes the first life sciences company to acquire a **DGX SuperPOD** based on the Vera Rubin architecture (8 NVL72 systems). Goal: train its own models and accelerate predictions on compounds and proteins, opening access to all its researchers, without the current queues.

**On TikTok Shop, AI videos compete with human creators**, Brands are increasingly replacing small creators with AI-generated influencers and product visuals, favored by the GMV Max advertising algorithm. Reactions diverge: SharkNinja bans AI for fear of bad publicity, others use it to test concepts before handing them to real creators.

**Nvidia unveils its Vera CPU and takes on AMD and Intel**, The GPU giant publishes the full specifications of **Vera**, its data center processor already delivered to OpenAI, Anthropic and SpaceX. The rise of autonomous AI agents gives the CPU a strategic role again, tasked with feeding the GPUs, in a server market Nvidia estimates at **$200 billion**.

**Nvidia wants to supply all the chips in AI data centers**, Ahead of AMD's product event, Nvidia unveils new benchmarks for its **Vera Rubin NVL72** system (36 CPUs for 72 GPUs), which it claims is **10 times more efficient** in tokens per watt than Grace Blackwell. The strategy: sell complete racks, GPU and CPU, rather than standalone chips.

**Stanford students walk out of Sundar Pichai's speech**, More than **100 students** left the graduation ceremony to protest Google's contracts with the American and Israeli governments and ICE. The episode, picked up by CNN and the BBC, illustrates a growing defiance among Generation Z: a Gallup poll shows **22% enthusiastic** about AI, versus 42% anxious.

**Super Micro jumps 15% amid record orders and SpaceX**, The server maker announces **more than $60 billion** in new orders in the last quarter, with margins raised to 15-17%. Its CEO mentions a one-gigawatt AI data center co-built with **SpaceX and xAI**, driven by demand for Nvidia GPU servers.

**Goldman Sachs launches a private markets platform for wealthy clients**, The bank consolidates its alternative activities and two teams dedicated to direct stakes in high-growth private companies, to allow its clients to invest early in future unicorns. The boom in AI investments (data centers, infrastructure) intensifies demand.

**OpenAI and Anthropic beef up their lobbying in Washington**, The two labs spent **$3.17 million** on federal lobbying in the second quarter (+23%), a record. Anthropic (**$1.97 million**) even surpasses Nvidia, on cybersecurity, copyright, cloud and public procurement files, while legacy tech and defense stagnate or decline.
