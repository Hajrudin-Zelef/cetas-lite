---
id: collect-240926-vision-ia/vision-ia/un-professeur-piege-un-exercice-32-etudiants-sur-35-ont-utilise-l-ia-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SpaceX", "United States"]
dates: []
keywords: ["agent", "agents", "benchmark", "cost", "cyber", "cybersecurity", "diffusion", "distillation", "gpu", "kimi", "mai", "mistral"]
source: docs/RAG/clean_en/vision-ia/un-professeur-piege-un-exercice-32-etudiants-sur-35-ont-utilise-l-ia.md
source_anchor: ""
source_lines: [72, 128]
sha256: 6e613ddb846ca791a31d385c084ecc79090d39230c2da99e2b2728cadde716b6
---

# 🧠 **RESEARCH**

**NVIDIA Cosmos-H-Dreams: a real-time generative simulator for surgical robotics**

Surgical scenes are a nightmare to simulate: deformable tissues, reflective surfaces, sutures, needles, smoke, occlusions. NVIDIA tackles the problem from the other end with Cosmos-H-Dreams, a model that **generates** the visual consequences of the robot's movements instead of computing them. It receives a starting image and the robot's kinematic stream, then produces the continuation of the scene continuously, in a closed loop, **on a single RTX PRO 6000 GPU**. The model was distilled from Cosmos-H-Surgical-Simulator via the FlashDreams library, and its integration was demonstrated on CMR Surgical's Versius surgical platform.

**A drug candidate in 9 months instead of 4 and a half years**

Insilico Medicine, listed in Hong Kong, claims to reduce to **about one year** (13 months on average, **9 months** for its fastest program) the time needed to identify a drug candidate, compared with about 4 and a half years using conventional methods. Generative AI proposes the biological targets and designs the molecules, which reduces to 60-200 the number of compounds to synthesize and test. Note what this does not cover: clinical trials, manufacturing and regulatory validation remain separate, unaccelerated steps. CEO Alex Zhavoronkov attributes part of the gain to the Chinese ecosystem, with AI R&D conducted in Montreal and Abu Dhabi and biological testing in Shanghai.

**The real bottleneck in AI drug discovery is data**

Since the 1950s, the cost of developing a new drug has roughly doubled every nine years, a phenomenon dubbed **Eroom's law**. Today: 10 to 15 years, between $1 and $2.5 billion, and more than 90% failures. AI is shifting the industry from physical screening of molecules to predictive design before any laboratory handling. But the article insists on the real brake: the quality and authenticity of bench data, and their effective integration into existing R&D systems.

**Dario Amodei: "we have never advocated for a ban on open-weight models"**

Anthropic's boss published a blog post on Monday to head off accusations mounting in the industry. His alternative position comes down to three points: control access to advanced chips, fight industrial distillation (he cites an attack he attributes to Alibaba/Qwen), and impose safety tests on any sufficiently powerful model, open or closed. A clarification that comes after the publication of a collective anti-restrictions letter that Anthropic did not sign.

**Anthropic, the only major AI lab not to have signed the pro open weight letter**

Nvidia, Microsoft, Meta, Palantir, Mistral, Google and finally OpenAI, after some hesitation, signed an open letter on Friday asking Washington not to restrict open weight models. Jensen Huang devoted **his very first post on X** to it: *"Open models strengthen safety and cybersecurity, accelerate innovation and diffusion, and enable sovereignty."* The text implicitly targets the measures the Trump administration is considering against players like Moonshot AI, the publisher of Kimi K3. Anthropic is the only major name absent.

**Microsoft launches MAI-Cyber-1-Flash, but still calls OpenAI for hard cases**

Microsoft unveils a compact model specialized in cybersecurity that reaches **96% on the CyberGym benchmark** once integrated into its MDASH multi-agent system. The appeal is economic as much as technical: intelligent routing sends only the hardest cases to GPT-5.4, which would cut costs by **50%** compared to direct use of frontier models. Dependence on OpenAI for complex reasoning, however, remains complete.

**METR puts a dollar figure on the moment when an AI agent costs more than a human**

The METR lab proposes the **"expenditure horizon"**, a metric that directly compares the dollar cost of an AI agent and that of a human to solve the same problem. First results on the NanoGPT speedrun: unflattering for current agents. METR acknowledges blind spots in the measurement, and notes that the coming generation of models could reshuffle the deck.

**What if the next step wasn't bigger models, but agents that coordinate**

Cisco, through its Outshift division, advocates a two-layer architecture, an *"Internet of Agents"* for connectivity and an *"Internet of Cognition"* for shared meaning, so that specialized agents truly collaborate on a common goal instead of merely exchanging data. The quantified argument stings: a cited study measures a failure rate of **41 to 87%** across seven open source multi-agent systems evaluated. The thesis: move from vertical scaling to horizontal scaling. Note, this is sponsored content.

**Lasers to extract nuclear fuel from old enrichment waste**

Near Paducah, Kentucky, thousands of cylinders store the residue from a closed enrichment plant. Global Laser Enrichment wants to reprocess them with a technique that no longer separates isotopes by mass, as centrifuges do, but selectively targets them using the **vibratory signatures specific to U-235**. Goal: produce fuel at the concentration of natural ore, including for advanced reactors. Nuclear power today accounts for about **9% of global electricity**.

**When should an AI contradict you?**

This study argues that model sycophancy is a subtler phenomenon than simple flattery. Three factors determine whether a model changes its moral judgment: the distance between the opinion expressed and its initial position, the identity of the person expressing disagreement, and whether a group supports it or not. The practical stakes: knowing how to distinguish legitimate reconsideration from submission to social pressure, in conversations where it really matters.


# **🗞️MORE NEWS**

**Alibaba releases open-code-review, a free AI code review tool**

Alibaba open-sources the code review tool it uses internally, on codebases at its scale. The architecture is hybrid: deterministic pipelines for what can be checked mechanically, an LLM agent for the rest, with precise line-by-line comments. The built-in ruleset is fine-tuned to detect NullPointerExceptions, thread-safety issues, XSS vulnerabilities and SQL injections. Compatible with the OpenAI and Anthropic APIs, so it can be plugged into the model of your choice. Already more than 15,000 stars on GitHub.

**Hugging Face has a massive problem with non-consensual nude deepfakes**

The European NGO AI Forensics tested the 9 most popular image-editing "Spaces" on Hugging Face: **7 out of 9** allowed a photo of a clothed woman to be undressed with no obstacle whatsoever. A honeypot-style experiment conducted over one week, with more than 1,000 requests tracked, shows that **73%** of requests were sexual in nature, of which **95% targeted women** and **6.7% apparent minors**. Hugging Face, valued in the billions, did not respond to Wired's questions about its moderation but removed certain pages after being reported. The EU and the United Kingdom are preparing a ban on "nudification" apps by the end of the year.

**A quadcopter drone hoists a man out of floods in China**

Footage broadcast by Chinese state media shows a large quadcopter drone strapping a man into a harness and lifting him above the waters, after Typhoon Maysak passed over Guangxi province. The sequence is spectacular and raises a concrete logistics question: at what point does a drone become a standard means of rescue rather than a demonstration? **More than 130,000 people** have been evacuated in the region, with up to 35 inches of rain (about 89 cm) in a few days.

**Nvidia, Microsoft, SpaceX and IBM found the Open Secure AI Alliance**

