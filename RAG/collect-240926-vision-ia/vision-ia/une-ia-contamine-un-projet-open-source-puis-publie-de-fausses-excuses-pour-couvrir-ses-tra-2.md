---
id: collect-240926-vision-ia/vision-ia/une-ia-contamine-un-projet-open-source-puis-publie-de-fausses-excuses-pour-couvrir-ses-tra-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "China", "CoreWeave", "Google", "Groq", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "Perplexity", "Stripe", "xAI"]
dates: []
keywords: ["research", "accelerator", "agent", "agents", "benchmark", "chatgpt", "claude", "compute", "cybersecurity", "gemini", "gpus", "grok"]
source: docs/RAG/clean_en/vision-ia/une-ia-contamine-un-projet-open-source-puis-publie-de-fausses-excuses-pour-couvrir-ses-traces.md
source_anchor: ""
source_lines: [107, 176]
sha256: db7271e141bf01de8f7bab24f9f8915be8bafc6bd9462a7d4507a4f59bba7354
---

# 🧠 **RESEARCH**

Hugging Face is the place from which almost all the open models you download to run at home originate: the model and dataset hub, Spaces, the Transformers library. According to Business Insider, the company has been approached for a sale and is talking to banks to evaluate offers. **Nothing is signed**, and the potential buyer has not been identified.

The key points:

- Valuation mentioned: **$13 billion or more**, compared with **$4.5 billion** at the last raise in 2023, led by Salesforce Ventures with Alphabet, GV, and IBM Ventures.

- Earlier this year, Hugging Face **turned down $500 million from Nvidia** (a $7 billion valuation), to avoid a dominant investor weighing on its decisions.

- Clem Delangue, the CEO, describes a company **"close to profitability"**, which has "recently started to touch the money raised three years ago," and says it aims for long-term sustainability rather than maximizing raises.

- The market is pushing in that direction: Stripe has just acquired **OpenRouter for $7 billion**. AI infrastructure is being bought up.

- Echoing our first topic of the day: the platform recently was the **target of an OpenAI system** that escaped its sandbox during a cybersecurity evaluation, going so far as to reach its servers.

For you, nothing changes today, since no agreement exists. But Hugging Face is the neutral pipe through which the open weights of Qwen, Mistral, Llama, and tens of thousands of other models pass: the day that pipe belongs to a player with its own models to sell, the question of neutrality takes on a different meaning. Delangue insists on his responsibility toward a community "that entrusts us with its data and its models," which looks less like a hurried seller than someone politely listening to offers.

# 🧠 **RESEARCH**

**Why children learn to speak with 100,000 times fewer words than an AI**

A child hears between 100 and 300 million words before turning 20, and masters their language perfectly. Llama 3.1 needed 15 trillion tokens for a comparable result. Michael C. Frank (Stanford) and Ethan Gotlieb Wilcox (Georgetown) dig into this "data efficiency gap," which is becoming strategic as the reservoir of text available on the internet could run out as early as the 2030s.

**Chatbots steer pregnant women toward anti-abortion sites without flagging it**

AlgorithmWatch analyzed 270 responses from ChatGPT, Gemini, Grok, and Claude on unplanned pregnancies. The anti-abortion organization Profemina appears in 17% of responses, without its ideological orientation being mentioned. In Germany, the chatbots also refer to Caritas for the mandatory prior consultation, even though the association does not issue the required legal certificate.

**Cerebras doubles its accelerator's performance without changing the chip**

The CS-4 runs on the same 5 nm WSE-3 chip as the CS-3, but clocks higher thanks to more electrical power and better cooling: double the performance. Three wafers per rack instead of two, up to 4,400 tokens per second per user, or 30 times a setup based on Nvidia GPUs according to the company. Memory remains at 44 GB per wafer. The hardware already equips OpenAI's Codex Spark.

**Spline V2: coding agents now edit 3D scenes live**

The 3D editor in the browser has been rebuilt on WebGPU, with a built-in agent that inspects the scene, uses the editor's tools, and takes screenshots to verify its own work. It modifies objects, materials, lights, cameras, particles, variables, and states, and writes HTML and JavaScript for the interface and game logic. An MCP server opens the editor to external agents, including Claude Code.

**Coding AI will prevent the formation of expertise**

The skilled orchestrator paradox: the best results with coding agents come from already seasoned developers, but prolonged use of these tools removes precisely the friction that forges that expertise. The author fears a generation of juniors in "confidence without understanding," capable of delivering code that works without mastering its mechanisms. The topic is generating a lot of reaction in the tech community.

**How to encourage smarter use of AI in the classroom**

Cheshire Academy chose to train its teachers in general techniques rather than imposing a specific tool. The result: a mix of general-purpose chatbots (ChatGPT, Perplexity) and specialized tools like MagicSchool. Despite recommendations from OpenAI or UNESCO, many teachers are still moving forward without clear guidance.

# **🗞️MORE NEWS**

**Gradio turns your AI pipelines into visual graphs**

Hugging Face integrates `gr.Workflow` directly into Gradio: you describe your pipeline as a graph of typed nodes, each executable and inspectable separately. No more print-debugging to find the step that produced the weird value. The graph automatically becomes a drag-and-drop interface, a REST API with one endpoint per output, and an application deployable in a single command on Spaces. The provided examples range from prompt-based image editing with Qwen-Image-Edit to a media studio combining FLUX, speech synthesis, and an LLM.

**Qwen 3.6 becomes much simpler to run locally on Mac**

JetBrains has simplified local execution of Alibaba's open model on macOS. In practice, less tinkering to install and run Qwen 3.6 on your machine, and one less cloud dependency for data that must not leave. It's the kind of integration that does more for local AI than one more benchmark.

**OpenAI wants an agent for every profession, but it remains to be seen who will give it the keys**

ChatGPT Work, launched last month at $20 per month, adapts Codex for non-developers and gives the agent full access to emails, Slack, Notion, and Figma. Andrew Ambrosino, chief engineer of the desktop app, entrusted it with his inbox, his Slack, and his phone to test the thing. The stakes are twofold: autonomous agents consume far more tokens, and OpenAI must convince accountants, lawyers, or doctors against vertical specialists like Harvey in law or Clay in sales.

**XPENG raises more than 900 million for its IRON humanoid**

The Chinese manufacturer's robotics division is valued at $6.3 billion, with IDG Capital as lead and Tencent as well as Alibaba as strategic investors. XPENG keeps control of the subsidiary. The money will go to the "physical" foundation model, in-house Turing AI chips, mass production, and commercial expansion outside China. The company presents the operation as the largest private round ever closed in Chinese physical AI.

**Unitree's founder gains 13 billion dollars in a single session**

Unitree's stock jumped 460% for its IPO in Shanghai, up to 629% during the session, valuing the humanoid maker at around $50 billion. It is the first pure player in the sector listed in mainland China, and retail investors requested more than 8,000 times the shares offered. The most delicious part remains the position of Wang Xingxing himself, suddenly worth nearly $16 billion: according to him, the real "ChatGPT moment" for humanoids is still 2 or 10 years away. Investors, however, are clearly not waiting.

**General Intuition trains its future robots with hours of gameplay**

The New York startup is negotiating a raise at a $6 billion pre-money valuation with Valor Equity Partners, Point72 Ventures, and Seven Seven Six, only weeks after a $320 million round at $2.3 billion. Its foundation model teaches agents to move through space and time, based on hundreds of millions of hours of video game clips from Medal, the sharing platform founded by its CEO Pim de Witte. The focus is now on robotic embodiments, with a compute partnership at CoreWeave.

**Nvidia puts Groq racks into production for the end of the year**

