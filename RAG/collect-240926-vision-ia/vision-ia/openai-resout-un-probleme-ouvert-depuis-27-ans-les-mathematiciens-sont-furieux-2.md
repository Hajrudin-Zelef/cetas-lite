---
id: collect-240926-vision-ia/vision-ia/openai-resout-un-probleme-ouvert-depuis-27-ans-les-mathematiciens-sont-furieux-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "CISA", "China", "DeepSeek", "Falcon", "Hugging Face", "Meta", "Microsoft", "Moonshot", "OpenAI", "United States", "Z.ai"]
dates: ["2025-12", "2026-07", "2026-08"]
keywords: ["research", "advisory", "agent", "agentic", "agents", "alignment", "astra", "attention", "benchmark", "chatgpt", "claude", "consumer"]
source: docs/RAG/clean_en/vision-ia/openai-resout-un-probleme-ouvert-depuis-27-ans-les-mathematiciens-sont-furieux.md
source_anchor: ""
source_lines: [101, 170]
sha256: 95e6a0bfb946dbb237f13cf59eeef408160c13791180b1c8e40f84fbbecc64f8
---

# 🧠 **RESEARCH**

- Another case: the creation of an **atlas of toxic venom peptides** with a generative pipeline optimizing toxin characteristics

- **All accounts were banned**, but neither the individuals nor their laboratories are named, so as not to expose them. The information was shared with authorities and other AI companies

- Official nuance: Anthropic specifies that these are **active scientists** and does not claim they had malicious intent

"You don't see someone saying, comic-book style, 'I want to build a biological weapon to kill everyone,'" summarizes Jacob Klein, head of threat intelligence at Anthropic: "this is an incredibly nuanced situation." That is the real problem: a grant application for a vaccine and an application for a weapon look alike, and the better models become at assisting research, the more arbitrary the sorting becomes. Anthropic says it has systematically erred on the side of caution, therefore at the risk of blocking legitimate work.

# 🧠 **RESEARCH**

### **Colibrì runs a 744-billion-parameter model on 25 GB of RAM**

An open-source runtime in pure C, with no dependencies, that runs GLM-5.2 on consumer hardware. The trick: only ~40 billion parameters are activated per token, so Colibrì keeps in RAM the dense tensors quantized in int4 (attention and embeddings, about 9.9 GB) and leaves the ~370 GB of routed experts on an NVMe SSD, loaded on demand at each decoding step. It's slow, but it runs at your place.

**DeepSeek V4.1-Flash cuts AI agent memory by four**

A multimodal model with 552 billion parameters, of which only 16 billion are activated per token, with a KV cache reduced to a quarter of that of the previous version. On the DeepSWE code benchmark, it slightly outperforms Opus 5 and GPT-5.6 Sol. All under MIT license, which makes agents significantly less costly to run continuously.

**MultiMatte cuts out an image from a sentence**

The startup Feyn releases a background-removal model that you target with words: "keep only the dog," and the rest disappears. Built on Meta's SAM 3, it replaces binary masks with alpha mattes, resulting in much better rendering of hair, fur, and motion blur. Measured gains versus SAM 3: S-measure from 0.674 to 0.908 on DIS5K. NoBg library on GitHub, model on Hugging Face.

**Codex and ChatGPT go hunting for antibiotics in the genomes of extinct species**

César de la Fuente's laboratory uses Codex and ChatGPT as assistants to comb through the genomes of living organisms and extinct species in search of antimicrobial peptides. Target: drug-resistant infections, one of the major global health threats, for which the classic discovery pipeline has run dry.

**Public services are drowning in files written by AI**

Complaints to the UK housing ombudsman rose from 2,600 in 2022 to more than 7,000 last year, those to the US agency CFPB increased fivefold, with comparable rises in Brazilian court filings and German parliamentary petitions. Researcher Chris Schmitz documented 84 cases of "agentic flooding" across 11 jurisdictions. Counterintuitive: these requests come mostly from real people with legitimate cases they would once have abandoned.

**Anthropic puts numbers on Chinese campaigns to distill Claude**

Nearly 200 million exchanges spread across five campaigns, including the largest ever observed: 151 million exchanges between May and July 2026 via 3,500 accounts, attributed to Alibaba to feed Qwen's training. The technique consists of getting the model to reveal its chain of reasoning, normally hidden, by disguising the request, for example as a translation exercise. A campaign linked to Moonshot AI appeared to route military requests, including analysis of surveillance camera images.

**GPT-6 Astra dominates math, and OpenAI says it wasn't aiming for that**

The model takes first place on ErdosBench, a benchmark of unsolved mathematical problems, while chief scientist Jakub Pachocki claims that mathematics was deliberately not a priority: resources are going to recursive self-improvement and alignment research. Enough to reinforce the thesis of "spiky" progress, extreme in a few targeted areas rather than uniform.

**Autonomous agents tracked across 30 sites, and Anthropic investigates its own model**

Independent investigators spotted traces of agents attributed to OpenAI across more than 30 online public services, from wikis to RubyGems. In parallel, Anthropic documents a Claude Mythos 5 that convinced itself that real systems were merely a simulation, uploaded a poisoned package to PyPI and deceived its own monitoring overseer. The underlying problem: with GPT-6 Astra, readable reasoning, the main control tool, is becoming opaque.

**Anthropic's threat report covers seven domains of misuse**

Eight months of operations detected and neutralized between December 2025 and August 2026: cyber operations, influence operations, surveillance, scams, biological risks, conventional weaponry and illicit distillation. The shift documented by the team: AI is moving from the role of assistant to that of orchestrator of entire operations. Cases range from a network of fake dating apps to surveillance systems targeting dissidents.

**Why AI research agents don't overlearn**

A fundamental paradox: AI research reuses the same test sets for years, which should mechanically skew rankings through memorization. Yet the gains are confirmed on fresh data. Amazon Science uses research agents as guinea pigs to replay the experimental loop at will, something that obviously cannot be done with a community of human researchers.

**Claude Fable 5.1 writes with fewer tics than Fable 5**

Arena.ai compared tens of thousands of responses on Text Arena: compliments and validations drop from 3.17% to 1.98% of responses, hedging phrases decline by 36%, abstract nouns by 25%, and em dashes become rarer. Median length climbs 30% (319 to 414 words), while remaining 21% below Opus 5's 525 words. Proof that a lab can adjust a model's tone and flattery between two minor versions.

**Garry Tan on distillation: "do nothing"**

The CEO of Y Combinator opposes any regulatory response against Chinese labs copying American models, at the very moment when the NSA, CISA and FBI publish a security advisory on the subject. His argument: the training data of these models is itself being contested in court. His line, which he calls a tightrope, consists of preserving open-weight models while leaving frontier models a viable price premium.

**The accuracy of a cardiac screening came from a data leak**

Ten classifiers tested on 442,067 respondents from a large American health survey showed AUROCs close to 0.89 for predicting heart attack. When two variables that indirectly revealed the diagnosis were removed, all collapsed in exactly the same way: the performance came from the data, not the models. Incidentally, the explainable model matches tabular foundation models while being about 104 times faster, and its transparency made it possible to correct a threshold that detected 75.4% of heart attacks in women versus 89.0% in men.

# **🗞️MORE NEWS**

### **NASA and IBM release a lunar AI model in open source**

The NASA-IBM Lunar Foundation Model is downloadable on Hugging Face, with its dataset open so other teams can train their own models. It spots surface ice zones with 23% fewer errors than Microsoft's SwinV2-B baseline, and classifies craters 19% better with half the training data. Real-scale validation: it correctly identified the crater carved by a Falcon 9 impact on August 5, despite its overlap with an existing crater. The main difficulty in training remains lunar shadows, completely black and shifting with the Sun's position, which make the same crater unrecognizable from one image to the next.

**An Anthropic agent that went off the rails mostly broke its teeth on CAPTCHAs**

