---
id: collect-240926-vision-ia/vision-ia/la-premiere-puce-d-openai-bat-nvidia-avec-deux-fois-moins-d-electricite-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Moonshot", "Nvidia", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["research", "agent", "benchmarks", "chatgpt", "claude", "compute", "cost", "deepseek", "embedding", "gemini", "gpu", "grok"]
source: docs/RAG/clean_en/vision-ia/la-premiere-puce-d-openai-bat-nvidia-avec-deux-fois-moins-d-electricite.md
source_anchor: ""
source_lines: [111, 177]
sha256: d5b7908a688995fd3bb42bbeea5587793adb763d5778b7338ee3b5b2d187d681
---

# 🧠 **RESEARCH**

SpaceX has filed with the FCC a dossier named Starmind: a constellation of data centers in low orbit (500 to 2,000 km) that could reach **one million satellites**. Each unit carries **72 Nvidia chips**, the equivalent of a complete server rack, and the first unit, AI1, is targeting a launch in the last quarter of 2027.

**A few key figures:**

- AI1 runs on the **Vera Rubin NVL72 "Space-1"** platform, Rubin GPU plus Vera CPU. Nvidia claims up to **25 times** the AI compute of an H100.

- Sun-synchronous orbit, facing the Sun **98% of the time**. The solar panels produce around **210 kW**, for **175 kW** of average compute power per satellite.

- Heat escapes into the vacuum through liquid radiators of around **1,700 square feet** (nearly 160 m²), the most fragile point of the dossier.

- The results come back down via **laser links** relayed by Starlink. Stated goal: free xAI from the electrical grid, land and terrestrial cooling constraints, and train Grok 5.

- At this stage, it's a filed dossier, not an authorization: no FCC decision, no public budget. Eventually, the project is to rely on Terafab, the chip factory at **119 billion dollars** set up with Tesla.

The irony is neat: the chip that Jalapeño outclasses on paper goes off to conquer orbit. Musk calls the NVL72 the "best architecture" available, but the specialists interviewed by Engadget point out that embedding the most recent and most power-hungry architecture is precisely the hardest way to start: thermal dissipation, cost, radiation, and zero possibility of going to replace a faulty card. The 2027 timeline is to be taken for what it is, a Musk timeline.

# 🧠 **RESEARCH**

**An mRNA adjuvant that melts tumors in mice**

Anti-cancer vaccines often trigger an immune response that is too weak, and boosting it with cytokines causes severe side effects. The team of Daniel Anderson (MIT), with Harvard and the University of Houston, designed an mRNA adjuvant encoding two genes that activate immune cells through other pathways. In mice, it slows or eliminates bladder, colon, lung and melanoma cancers, alone or combined with a targeted vaccine or checkpoint inhibitors. Combined with flu and covid vaccines, it multiplies the T-cell response by 10 to 15.

**Four weeks of chatbot, and you spot fake news less well**

The MIT Media Lab had pairs of headlines and images evaluated over a month. Assisted by a chatbot, participants were initially 21% more accurate at distinguishing true from false. By the fourth week, deprived of AI, they were 15% worse than before the experiment, while a quarter of them declared themselves on the contrary sharper. AIs that ask questions instead of giving the answer clearly limit this effect, at the cost of speed.

**A model compressed to 4 bits that beats its original version**

Multiverse Computing presents Quantization-Aware Healing, a repair method for models that have undergone structural compression (removal of layers, heads, neurons) in addition to quantization, a case that conventional techniques handle poorly. Result on a GPT-OSS 120B pruned to 60B then quantized to 4 bits: the final model outperforms the bfloat16 version it came from on 7 out of 9 benchmarks, including reasoning, math, and code. In other words, a smaller model, cheaper to run, and yet better.

**An MIT AI invents weather disasters that never happened**

Current risk models can only replay what is already in the archives. The Extreme Event Aware method (η-learning), by Kai Chang and Professor Themis Sapsis, generates maps of extreme events that are statistically possible but never observed in a given region, with an estimate of duration, intensity, and affected area. Published in Nature Communications on August 20, it targets insurers, urban planners, and power grid managers who must anticipate a "hundred-year Katrina."

**A sonar to pilot a robot in murky water**

When an underwater vehicle lands or digs, it stirs up a cloud of sediment and its cameras go blind. Amy Phung and Richard Camilli, of the Woods Hole Oceanographic Institution, first have the environment mapped by sonar, whose resolution is low but which works as well in clear water as in murky water, then bring the craft close enough for the cameras to take over. An image-matching algorithm estimates the depth of each pixel to keep up with real time. Intended applications: exploration, underwater maintenance, and demining.

**Faraday reproduces scientific papers with 27 billion parameters**

Inherent, a London laboratory founded by former Google DeepMind employees, claims that its agent Faraday outperforms Claude Opus 4.8 and GPT-5.5 on a specific task: independently rediscovering the results of a published study, without knowing the answer in advance, the equivalent of a standard doctoral exercise. All on Qwen 3.6, a model with 27 billion parameters, far smaller than its competitors. Inherent's bet is on reinforcement learning, to create a researcher's "instinct" rather than follow rules.

**Detecting an AI's lies: EleutherAI's assessment**

EleutherAI publishes its retrospective on Aletheia's Quest, a month-long competition organized by Cadenza Labs and the National Deep Inference Fabric. Finishing 2nd out of 19, the team tested hundreds of detector variants on more than 30 datasets and three families of open-weights models ranging from 27B to 120B. Main surprise: black-box methods, which observe only the outputs, often rival white-box methods that read internal activations, the latter proving highly context-dependent. The datasets are released.

**Ukraine opens its five million combat images to the British**

The United Kingdom becomes the first country to gain access to Avengers Labs, the Ukrainian platform that gathers around five million annotated combat images, raw material for training military AIs. Three British startups are already running pilot projects. The agreement marks a quiet shift: real wartime data is becoming a diplomatic bargaining chip for developing autonomous weaponry.

**Students prefer Gemini for writing their essays**

Out of 6,851 blind votes collected by StudyArena, Gemini wins for writing with 39.6% of preferences, ahead of Claude (31.8%) and ChatGPT (29.2%). Two takeaways more useful than the ranking: long answers win significantly more often (37% longer on average), and increasing the reasoning setting does not improve writing quality. The authors' advice: use AI as a proofreader, to hunt down generic passages and missing details, and write the final text yourself.


# **🗞️MORE NEWS**

**Chinese hackers double their pace thanks to DeepSeek**

Groups linked to the Chinese state carry out more than twice as many attacks since they integrated open source models into their tools, reveals Taiwanese company TeamT5 in a report relayed by Bloomberg. Their model of choice is DeepSeek, for two very prosaic reasons: it costs almost nothing and its guardrails are particularly lax. Researchers have, however, never yet encountered Kimi K3, which is more expensive, in an attack. The downside of open source that people rarely think about when installing a model at home.

**ChatGPT's task scheduling arrives on free accounts**

Until now reserved for Plus, Pro, Business, and Enterprise subscriptions, ChatGPT's scheduling menu is opening up to everyone. Concretely, you schedule a prompt that triggers itself at the time or pace of your choice (morning news roundup, weekly follow-up, Friday recap), and you can now share a scheduled task with other users. Paid accounts additionally gain event-based triggers: a prompt that fires when something changes in Gmail, Slack, or GitHub.

**OpenAI reinstates the 5-hour limit on Codex for Plus subscribers**

