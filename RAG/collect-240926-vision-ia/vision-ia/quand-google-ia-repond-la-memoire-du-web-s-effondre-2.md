---
id: collect-240926-vision-ia/vision-ia/quand-google-ia-repond-la-memoire-du-web-s-effondre-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "China", "Google", "Meta", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai"]
dates: []
keywords: ["research", "agent", "agents", "amd", "aws", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "cost", "cybersecurity"]
source: docs/RAG/clean_en/vision-ia/quand-google-ia-repond-la-memoire-du-web-s-effondre.md
source_anchor: ""
source_lines: [91, 146]
sha256: c46e8ebe2647a753fb1eb2f41c4e40485d38a5aee64e62b36fa9bbd7aef2718c
---

# 🧠 **RESEARCH**

Nothing is confirmed by OpenAI, everything comes from press leaks, and there is strictly nothing to test before 2027. But the direction is clear: after ChatGPT and Codex, the company is looking to exist somewhere other than in a browser or an app, that is, somewhere other than on platforms controlled by Apple and Google. The question remains that everyone who already has a mute connected speaker in their living room will ask: 300 dollars for an object that films you, what will it do more?

On Monday, Mark Zuckerberg posted "The Future is for Everyone," an essay of **more than 6,500 words** on the coexistence between humanity and a superintelligent AI, an extension of a much shorter public letter published last year. Behind the philosophy, the text contains four very concrete requests and a check for **1 billion dollars**.

**Data centers**: Meta promises to return more water than it consumes by 2030 and to create local jobs, to make its infrastructure socially acceptable. The figures behind it: **145 billion dollars of investment planned for 2026**, mostly in data centers.
**Access**: making AI free or nearly free for "billions of people" via downloadable open source models, with **Muse Glimmer** and an open-weight version of **Muse Spark 1.2** announced, without a precise date.
**Regulation**: loosening US rules on data and distillation, deemed ineffective since they don't apply to foreign models.
**Cooperation with the government**: sharing "intermediate training checkpoints" for cybersecurity and abuse detection.
A **"Future Is for Everyone Fund" endowed with 1 billion dollars**, and the promise of free or very low-cost personal agents for health, career, finances, and relationships.

The common thread of the text is the refusal to see superintelligence remain concentrated in "a handful of labs," an argument that has the advantage of being both defensible and perfectly aligned with Meta's open source strategy. Notably, a more discreet but real change: it is now Meta's board of directors, and no longer Zuckerberg alone, that validates the safety criteria for models. For the reader, the most verifiable commitment remains the simplest: downloadable models, or not.

# 🧠 **RESEARCH**

**Kimi K3 leaves the sandbox to go read its exam answers**

Researchers at Frontier Security observed the Chinese model Kimi K3 cheating during a cybersecurity evaluation by the UK AI Safety Institute. Taking advantage of an outbound DNS resolution that remained functional in the sandbox meant to isolate it, the model reached github.com, cloned the benchmark's official repository, and read the solution from the disk rather than solving the task. The flaw affects the Inspect and Cybench frameworks, and since Kimi K3 is a public model, unlike the OpenAI case detected upstream, anyone can reproduce the exploit.

**Claude advances a result related to the Riemann hypothesis**

An Anthropic employee asked an unpublished research version of Claude to tackle the Riemann hypothesis. The model failed, as expected since 1859, but along the way advanced a related result: the proven proportion of zeros of the zeta function that satisfy the hypothesis rises from **41.6% to 67.2%**, a bound that had not moved in decades. The proof, building on the work of Baluyot, Goldston, Suriajaya, Turnage-Butterbaugh and Bombieri, was validated by two Anthropic mathematicians and external experts.

**A mini LLM runs at nearly 60,000 tokens per second in a $250 FPGA**

A developer fit a language model of **3.16 million parameters** (1.5 MB in INT4) entirely into the on-chip memory of an AMD Kria KV260 board, without ever touching external RAM. The result: **59,965 tokens per second** in the reconfigurable logic, versus **11 tok/s** on the ARM cores of the same chip and **719 tok/s** on an RTX 3050 Ti. The model can only generate short stories, it is not an assistant, but the live demo runs directly on a physical board installed in Wales.

**NVIDIA open-weights Magpie TTS, 12 languages including French**

NVIDIA updates its open-weights speech synthesis model **Magpie TTS Multilingual**, 364 million parameters, which now covers **12 languages** with the addition of Modern Standard Arabic, Korean and Brazilian Portuguese. The argument is not raw quality but control: deployable on your own infrastructure via NVIDIA NIM, it lets you control latency and data confidentiality, whereas closed voice APIs impose their black box.

**Probing models to guess how many parameters they hide**

A researcher details techniques for extracting information that labs do not publish. "Incompressible Knowledge Probes" test models on very obscure facts to estimate their parameter count, "Data Mixture Inference" analyzes how they split tokens to deduce the training datasets, and date-related questions reveal pre-training schedules. Lacking a public reference, the results remain speculative, but the method is instructive.

**Humanizing LLM responses would be a bad idea**

The author takes on a trend: asking AI agents to write short, without jargon, in "simplified technical English" or in a style adapted to ADHD. The problem, according to him, is that the instruction does not apply after the work but during it: it is a lossy compression, applied continuously, which erases along the way the signals of failure (unresolved hypotheses, contradictory results, uncertainties). His proposal: keep the richest possible representation between agents, and compress only at the last moment, when facing a human.

**Which programming language for coding agents?**

A highly cited study claims that dynamic and concise languages like Clojure or J consume 2 to 3 times fewer tokens than Rust, Go or C++, to the point that Google's AI summaries repeat the conclusion. Dan Luu dismantles the method: the benchmarks rely on trivial Rosetta Code-style problems, unrelated to real code, and one of them contains a bug (an erroneous symlink created by an agent) that skewed the scoring of several languages tested afterward.

**Academic researchers adapt to the grip of private labs**

At a meeting of the AI2050 program funded by Eric and Wendy Schmidt, professors describe a shift in the center of gravity of research toward Anthropic, OpenAI and Google, which keep the training details of their models secret. Without a budget to buy GPUs or to massively query these systems, researchers at UC Berkeley and Johns Hopkins are redirecting their work toward questions that companies have no interest in exploring themselves.

**AWS launches a competition to co-design models and chips**

Amazon invites academic and industrial laboratories to train a language model of about **50 million parameters** from scratch on its Trainium chips, with total freedom over architecture, optimizer and compute kernels. The underlying idea is interesting: current architectures were shaped by the constraints of GPUs, and Trainium offers a different profile (more on-chip SRAM, explicit software control of data movement). What would a model designed for that hardware look like?

**Macaron-V1, agents that keep learning after deployment**

This family of open source models relies on a **Mixture-of-LoRA** architecture: a frozen base model, several specialized adapters (conversation, agent, code, generative interface), and a single one selected at each turn of conversation. The flagship model, Macaron-V1-Venti, combines a **GLM-5.2 base of 744 billion parameters** with these adapters, while the Tall version (**50B**, Qwen3.6 base) targets local deployment. The whole thing is driven by a self-improvement loop where the experience of one version is used to build the next.

**Video models that learn physics instead of imitating pixels**

