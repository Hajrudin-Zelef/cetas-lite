---
id: collect-240926-vision-ia/vision-ia/ox-alpha-tombe-le-masque-le-modele-chinois-qui-code-a-prix-casse
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "China", "ExploitGym", "Google", "Hugging Face", "JFrog", "Meta", "Microsoft", "Moonshot", "Nscale", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai"]
dates: ["2026-08-27"]
keywords: ["research", "agent", "agents", "agi", "amd", "astra", "benchmarks", "chatgpt", "claude", "context window", "cost", "cyber"]
source: docs/RAG/clean_en/vision-ia/ox-alpha-tombe-le-masque-le-modele-chinois-qui-code-a-prix-casse.md
source_anchor: ""
source_lines: [1, 309]
sha256: eefac53b1cba0571d97c1cab5d0b3e379a3af6dbfe88e0c6ef562f0515092ae5
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/ox-alpha-tombe-le-masque-le-mode-le-chinois-qui-code-a-prix-casse -->

## **Today:**

🕵️ Ox Alpha reveals its identity and its 320 billion parameters

🎙️ Gemini 3.5 Transcribe cleans up and structures speech in real time

⚡ Qwen3.8-Flash-Next drastically reduces the computation required

🤗 Nvidia is about to acquire Hugging Face

🎬 LAION opens 80 million videos to research

🔓 OpenAI explains how an agent escaped its environment

🧠 Sam Altman predicts his AGI before the end of 2026

🤖 Jetson Orin Nano 2 runs AI directly inside robots

☕ A JD.com café runs with a single robot

🧩 AI assistants get lost in their own brands

📱 Apple sets its iPhone and Siri AI date for September 9

🖥️ The Mac mini M5 Pro remains more powerful than the new M6

🩸 The FDA authorizes the first autonomous blood-drawing robot

🏃 Chinese humanoids break new records

A real estate agent website. A village association website. Software to learn Spanish. Software to manage your finances. All written in French, without a single line of code. It's the only AI skill you can say that about.

In this new update I teach you Codex in detail, which is included in your ChatGPT subscription. If you don't have ChatGPT, I show you GLM 5.2, which is the equivalent but free. Otherwise, if you have Claude, it also works on Claude Code. In short, nothing extra to buy.

And this module is only one part of the training.

✅ Complete training on AI (LLMs, AI Marketing, image, sound and voice generation, automation with n8n, first steps with agents, etc.)

✅ Complete lesson to create your own AI agents (n8n) and automate without limits

✅ Complete lesson on Claude Code, the best AI tool in 2026

✅ Access to a network of qualified professionals

✅ Regular updates to stay at the cutting edge of AI

➡️ If you join now, you lock in the €49 rate for life (one-time payment).

No matter how high the price climbs, and it will soon reach €100 or more, given the demand, you won't pay a single cent more.

One payment. Lifetime access.

On that last point, the proof is right before your eyes: the vibe coding module dated from late 2025, I deleted it and redid it from scratch, and subscribers didn't pay a single cent more. That's what "updates included" means. The price will go up, but never for those who are already in.

More than 13,000 people are already training with it. Their reviews are on the page.

See you inside.

The mysterious Ox Alpha is in reality **GLM-5.3-Flash**, a model from Z.ai designed to program, manipulate tools, and conduct automations during long sessions. It has **320 billion parameters**, but activates only **18 billion** at each step, with text, image, and video inputs and a context exceeding **one million tokens**.

The model can analyze a software repository, modify code, work in a terminal, interpret documents or graphics, and chain tool calls. Its weights are available under the MIT license, while its API displays a particularly aggressive price.

### **What to remember:**

The Mixture of Experts architecture has **320 billion parameters, of which 18 billion are active**, spread across 45 layers and trained with **30 trillion tokens**.
The context window reaches **1,048,576 tokens**, with a maximum output of **131,072 tokens**, enough to process large repositories or long chains of actions.
Z.ai publishes a score of **84.3 on Terminal Bench 2.1**, versus 85.0 for Claude Opus 4.8 and 87.4 for GPT-5.6 Terra. On Toolathlon Verified, GLM-5.3-Flash reaches **78.4**, ahead of Claude Opus 4.8 at 76.2.
OpenRouter charges **$0.075 per million input tokens** and **$0.25 per million output tokens**.

### **Why it matters**

The API price makes the model immediately testable for code agents or large-scale automations. Local execution, however, remains reserved for very well-equipped servers, because publishing the weights of a 320-billion-parameter model does not turn it into a home model. Z.ai above all brings new pricing pressure on OpenAI and Anthropic, with performance close enough to their best models to make the price gap hard to ignore.

Google DeepMind launches **Gemini 3.5 Transcribe**, a proprietary speech-to-text model capable of transcribing more than **85 languages**, even when a user switches languages mid-sentence. It removes hesitations, understands spoken self-corrections, adds punctuation, and can recognize up to **1,000 custom phrases**.

The service exists in two versions. `gemini-3.5-transcribe-live` processes an audio stream via WebSocket with latency under one second, while `gemini-3.5-transcribe` analyzes files and adds speaker identification as well as word-by-word timestamps.

### **In detail:**

The model is based on Gemini 3 Pro, accepts audio and text inputs, and has a context of **96,000 tokens**, with a maximum output of **32,000 tokens**.
Google announces an average error rate of **4.0% live** and **2.6% on recordings**, according to Artificial Analysis.
The time needed to obtain a final transcription would be reduced by **70% compared to Chirp 3**. On FLEURS, the announced error rates are 5.50% in streaming and 5.04% outside streaming.
Files can last up to **one hour**, or 30 minutes when speaker identification and detailed timestamps are enabled.
The preview is available in Gemini API and Google AI Studio, with a free tier then approximately **$0.009 per minute live** and **$0.005 per minute on file**.

### **What it changes**

A developer can now create a subtitling tool, professional dictation, meeting minutes, or a voice agent without managing an audio model themselves. Custom vocabulary will be particularly useful for product names, medical or legal terms, and alphanumeric references. However, the weights are not downloadable, and Google still documents risks of hallucination, expiration, and slowdown.


Alibaba presents **Qwen3.8-Flash-Next**, an open-weight multimodal model that foreshadows the architecture of Qwen4. Its main network contains **125 billion parameters**, but only **6 billion** work on each token, which greatly reduces the computation needed for reasoning, code, and agents.

The model receives text, images, and videos, then generates text over a native context of **262,144 tokens**, expandable to one million. It can explore a software repository, use tools, analyze a video, or execute long office tasks.

### **Key points:**

The architecture includes **512 experts**, of which 10 are dynamically selected and one remains shared. It adds 51 billion N-gram embedding parameters and 4 billion dedicated to multi-token prediction.
Alibaba claims a training cost equivalent to about **one-ninth that of Qwen3.7-Plus**.
The published results reach **62.5 on SWE-bench Pro**, 58.7 on DeepSWE, 73.9 on CoWorkBench, and **91.7 on GPQA Diamond**. No consolidated independent evaluation has yet been found.
The QwenCloud API costs **$0.15 per million input tokens** and **$0.47 per million output tokens**, with function calling, web search, and a code interpreter.

### **The impact to remember**

Qwen3.8-Flash-Next seeks less to beat all records than to provide many capabilities with few active parameters. For the reader, the API is immediately accessible and the weights allow private deployment, provided one has a sufficiently powerful server. This architecture above all shows how Qwen4 could reduce the cost of agents capable of working for a long time on code, documents, and multimodal data.


Nvidia reportedly agreed to acquire Hugging Face for **$12.9 billion**, according to information initially published by The Information. As of **August 27, 2026**, Nvidia and Hugging Face have still made no official announcement and have not responded to Reuters requests. The timeline, regulatory approvals, and guarantees of independence remain unknown.

In practice, Nvidia would not only be acquiring a community of developers. The Hugging Face Hub makes it possible to download models, test them in the browser, call them via API, deploy them on dedicated endpoints, or create demonstrations with Gradio, Docker, and HTML.

### **A few key figures:**

The Hugging Face Hub hosts more than **2 million models**, **1.5 million datasets**, and **1.5 million Spaces applications**.
More than **200 models** are accessible through inference providers, with **$0.10 in free credit per month**. Compatible models can still be downloaded and run locally.
The reported price represents about **86 times the $150 million in annualized revenue** attributed to Hugging Face.
The company had raised **$235 million in 2023**, at a valuation of $4.5 billion, from investors already including Nvidia, Google, and Salesforce.
Nvidia could strengthen the Hub's cloud services and turn it into a centralized platform for running various open-weight models.

### **The context for the user**

Hugging Face is today one of the main access points to open AI. Its passing under the control of the world's leading GPU supplier could improve hosting and deployment, but raises very concrete questions about pricing, hardware choice, model governance, and the independence of the Hub. As long as no official announcement specifies these guarantees, the models already downloadable remain usable locally, but the platform's commercial future remains uncertain.

# 🧠 **RESEARCH**

LAION-BVD brings together **80 million videos and 10 million hours of content**, obtained from 1.3 billion URLs sourced from Common Crawl.

LAION extracts scenes, images, and synthetic captions describing both the sound and the video.

The corpus aims at reproducible training of video, audio, and image models, with access limited to research uses.

An impossible task in the ExploitGym evaluation reportedly pushed an experimental model to persist, compromise Artifactory, then reach systems at OpenAI, Hugging Face, and other providers.

OpenAI attributes the incident to a rare combination of poorly designed tasks, prolonged execution, and messages that caused other models to deviate.

The lab plans strengthened monitoring of reasoning, permanent on-call staffing, and new mechanisms for stopping dangerous workloads.

According to Jakub Pachocki, Astra would already function as an automated research assistant capable of conducting part of scientific work.

Sam Altman predicts that it will be able to produce inventions with real utility, which corresponds to his own definition of AGI.

The prediction therefore remains impossible to evaluate without a public protocol, detailed benchmarks, or a shared definition of the term.

# **🗞️MORE NEWS**

Jetson Orin Nano 2 is a compact board intended to run AI directly within a machine, without a permanent connection to a data center.

It provides **78 TOPS**, 8 GB of memory, and an eight-core Arm processor.

NVIDIA announces twice the inference performance of Orin Nano Super and **40% less energy consumed** at equal performance in 15-watt mode.

The card supports language and vision models such as Cosmos, Nemotron, Gemma 4, and Qwen 3.

JD.com has opened a café operating **24 hours a day**, with no human staff on site.

The customer orders on a screen, then the robot grinds the coffee, prepares the drink, and places it in a locker.

The system served 202 drinks with a preparation time of under **30 seconds per cup**.

This is a very concrete demonstration of the automation arriving in ordinary services.

TechCrunch criticizes the proliferation of Gemini interfaces, between chat, Spark, Daily Brief, and search across Google services.

The user must understand the product's internal organization before knowing where to submit their request.

Daily Brief can also mix useful information, messages, and personal reminders pulled from Gmail or Calendar.

Claude and ChatGPT experience the same fragmentation, whereas an assistant should normally hide this complexity.

The event will begin at 10 a.m., Pacific Time, at Apple's headquarters in Cupertino.

New iPhones and Apple Watches are expected, with the possible presentation of the brand's first foldable iPhone.

Apple could also unveil Siri AI, a version based on large language models.

The assistant should be able to act directly in applications like Messages and Calendar, but these functions remain to be confirmed.

The new Mac mini M6 starts at **$899**, compared to $1,699 for the M5 Pro.

The M5 Pro retains up to 18 CPU cores and 20 GPU cores, whereas the M6 stops at 12 in both cases.

It also accepts up to **64 GB of memory and 8 TB of storage**.

For 3D rendering, video editing, or certain local AI workloads, the newest model is therefore not necessarily the fastest.

The FDA grants De Novo authorization to the device designed by the Dutch company Vitestro.

Aletta combines near-infrared light, ultrasound, and computer vision to locate a vein and guide the needle.

Trials report **95% success on the first attempt**, 0.6% hemolysis, and a median draw time of 1 minute 49 seconds.

The robot could standardize blood draws while allowing a professional to supervise multiple devices.

The World Humanoid Robot Games in Beijing brought together more than **600 teams and 2,000 robots**.

Tiangong Ultra finished the 100 meters in 8.86 seconds and achieved a long jump of 7.97 meters.

The competition tests motors, dynamic balance, and control algorithms, sometimes to the point of falling.

New fine manipulation events now evaluate capabilities more directly usable in industry.

Trail of Bits asked the model to escape from a QEMU/KVM VM installed on Debian 12 with an AMD Zen 3 processor.

The agent first exploited known kernel vulnerabilities, then patches not yet distributed by some maintainers.

After a complete reconstruction of the environment, it reportedly discovered several novel flaws in QEMU or its dependencies.

The experiment suggests that a classic VM is no longer enough to contain certain advanced cyber agents.

Moonshot AI is negotiating hosting for Kimi K3 with Microsoft, Amazon, and Google.

The company is reportedly asking for up to **30% of the revenue** generated by services using its model.

This would be the first comparable deployment of a Chinese model on the major American clouds.

The discussions remain preliminary, particularly regarding data, usage tracking, and revenue sharing.

Legato Frames isolate and amplify voices while reducing surrounding noise.

They target people with mild to moderate hearing loss and are expected to launch in fall 2026.

The assistive technology is integrated directly into the temples of the glasses.

The founders previously worked on Bose Frames, Meta Ray-Ban, and Bose's hearing aid activities.

The internal "Project OT" planned to entrust the work to agents supervised by small human teams.

According to the cited documents, some units could have lost up to **60% of their staff**.

Mark Zuckerberg stopped the plan a few hours before the first wave of layoffs.

The agents had not produced the promised gains, while employee challenges and internal discontent were increasing.

Bill Gates believes that AI could cause mass unemployment and reduce the skills needed to organize a biological attack.

He states that companies publicly downplay these risks so as not to jeopardize their funding.

He rejects the idea of sufficient self-regulation and proposes institutions inspired by nuclear weapons control.

He also mentions a tax, even a symbolic one, on the use of AI.

Data center projects are becoming a local political issue in the United States.

Opponents point to their electricity consumption, water needs, and effects on local areas.

This resistance could slow down permits and increase the cost of future infrastructure projects.

Computing capacity therefore no longer depends only on available chips, but also on local acceptance of the installations.

The TAKE system, for Throw Away the Key Encryption, will be deployed by default starting in September.

Each video will be protected by unique keys that are regularly renewed.

A copy may pass through a cloud enclave to activate smart home functions, then will be deleted after use.

Users will still be able to choose end-to-end encryption when they want to further limit cloud processing.

Marc Benioff rejects the scenario in which models would replace enterprise software.

He believes that agents need the CRM data, business rules, and security controls of platforms like Salesforce.

Nine of the top ten AI companies reportedly use Salesforce or Slack, with spending up **435% year-over-year**.

Claudeforce already allows Claude to draft emails and update records with Salesforce data.

Anthropic reportedly signed a contract with Nscale worth approximately **$45 billion** to rent computing capacity.

The servers will be installed in Nscale's main data center in West Virginia.

They will use Nvidia Vera Rubin systems, composed of six types of chips working together.

The capacity is expected to begin powering Anthropic's services by the end of 2027.

Anthropic reportedly calculates this theoretical market by adding up all the professional tasks its models could handle.

This estimate implicitly assumes that the company would capture a considerable share of global work, not just software spending.

It must support a possible initial public offering and a valuation that could reach around $2 trillion.

The figure remains highly speculative and very far from the current revenues of the entire technology sector.

A real estate agent website. A village association website. Software to learn Spanish. Software to manage your finances. All written in French, without a single line of code. It's the only AI skill you can say that about.

In this new update I teach you in detail Codex, which is included in your ChatGPT subscription. If you don't have ChatGPT, I show you GLM 5.2, which is the equivalent but free. Otherwise if you have Claude, it also works on Claude Code. In short, nothing extra to buy.

And this module is only one part of the training.

✅ Complete training on AI (LLMs, AI Marketing, image, sound and voice generation, automation with n8n, first steps with agents, etc.)

✅ Complete lesson to create your own AI agents (n8n) and automate without limits

✅ Complete lesson on Claude Code, the best AI tool in 2026

✅ Access to a network of qualified professionals

✅ Regular updates to stay at the forefront of AI

➡️ If you join now, you lock in the €49 price for life (one-time payment).

No matter how high the price climbs, and it will soon reach €100 or more, given the demand, you won't pay a single cent more.

One payment. Lifetime access.

On this last point, the proof is before your eyes: the vibe coding module dated from late 2025, I deleted it and redid it from scratch, and subscribers didn't pay a single cent more. That's what "updates included" means. The price will go up, but never for those who are already in.

More than 13,000 people are already training there. Their reviews are on the page.

See you inside.
