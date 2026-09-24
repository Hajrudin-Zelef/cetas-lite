---
id: collect-240926-vision-ia/vision-ia/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "ByteDance", "China", "DeepSeek", "EU", "Google", "Huawei", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Nvidia", "OpenAI", "Perplexity", "SpaceX", "United States", "Xiaomi", "xAI"]
dates: ["2024-06", "2024-06-10", "2025-03-29", "2025-05", "2026-12-21"]
keywords: ["research", "accelerator", "advisory", "agent", "agents", "alignment", "amd", "astra", "attention", "benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/clean_en/vision-ia/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors.md
source_anchor: ""
source_lines: [1, 193]
sha256: 926993c838620251313bfdd8c181f14794534af4c5c5bfb8113713f14f77aaa6
---

# 🧠 **RESEARCH**

<!-- source: https://vision-ia.beehiiv.com/p/meta-la-che-son-agent-ia-dans-les-rayons-amazon-le-met-dehors -->

🛒 Amazon blocks Muse, Meta's AI agent, on its store

🇨🇳 Alibaba releases an AI chip three times faster than the previous one

⚡ Grok 4.7 slashes prices, independent tests dampen enthusiasm

🍎 25 dollars per iPhone for the promised and never-delivered Siri

📦 Xiaomi publishes the world's best open weights model, under MIT license

🧪 Microsoft opens RetroChimera, its AI that plans molecule synthesis

🔓 The guardrail that arbitrates AI agents' decisions flips with one sentence

🌊 French abyssal drones to cover 98% of the oceans by 2028

🇺🇳 The UN warns: nothing guarantees that humans keep control of agents

🏜️ Texas freezes all environmental permits for data centers

And a skill that everyone has doesn't get billed. Yet there remains a demand that ChatGPT cannot serve: an accounting firm will never put its client files in it. Neither will a lawyer, and even less so a medical practice. They want AI like everyone else, but without their files leaving the office.

It exists, and it's called local AI: the model runs on their own machine, and it answers even with the wifi turned off. Same for image and video, with no credit limit. Few people know how to install it, and nobody teaches it in French. A rare skill, that one gets billed.

On Sunday night, Amazon cut its site's access to **Muse**, the AI agent Meta launched on **September 8**. Muse's principle: you write to it on WhatsApp, and it browses on your behalf to fill a cart, draft an email, book an appointment, or negotiate a price. Amazon had asked Meta to remove its bot, Meta refused, and users now see a warning explaining that access through an unauthorized AI agent violates the site's terms of use.

**Key takeaways:**

- Meta unleashed Muse on Amazon.com **without any prior agreement**, according to Amazon, which was not notified of anything

- Two specific grievances: the agent **does not identify itself as AI** while it browses, and it **would capture and store customer data**, which according to Amazon creates a security risk

- Meta responds that Muse has access **neither to passwords nor to payment methods**

- Muse became **the most downloaded free app on the US App Store** in one week

- Amazon has already pushed out the shopping agents of **Perplexity, Google, and OpenAI**, and is suing Perplexity over its Comet browser (details at GeekWire)

- Paradox: the two groups remain partners, bound by a billion-dollar contract signed in April on Amazon's cloud chips

The real question is not technical, it is commercial: who controls the customer relationship when a piece of software buys on your behalf? An agent that compares, filters, and negotiates on its own short-circuits the storefront, sponsored recommendations, and advertising, that is, a considerable share of what makes Amazon money. Expect the big e-commerce sites to start filtering AI agents as methodically as they already filter indexing bots, and for access to be negotiated by contract, agent by agent.

At the Apsara conference in Hangzhou, CEO Eddie Wu unveiled the **Zhenwu V900**, an AI accelerator designed in-house by Alibaba Cloud's T-Head division, announced at **three times the performance** of the Zhenwu M890 released in May. Shortly after, the group promised to exceed **20 gigawatts of data center capacity by 2032**. The stock gained about **3% in Hong Kong**.

**In detail:**

- **216 GB of memory** per card, chip-to-chip interconnect at **1,200 GB/s**, native support for **FP8 and FP4** compute formats

- Designed for very large clusters: up to **500,000 cards** linked in the "Apsara Supernode" stack (ICN interconnect, Pangu NIC network card, Zhenyue SSD storage), detailed here by Pandaily

- Mass production and commercialization in the **first quarter of 2027**; current Zhenwu chips already equip **more than 650 customers** in automotive, finance, energy, and industry

- On the models side: **Qwen 4 is being trained**, the Qwen 4.5 and Qwen 5 series are planned, and Alibaba announces it is targeting models of **5,000 to 10,000 billion parameters**

You will not be able to buy this chip: it stays in Alibaba Cloud's data centers, with no public price or individual availability. Its effect will be read elsewhere, in the pricing and quality of the Qwen models, which many already run at home in their open versions. If Alibaba trains its next Qwen models on its own silicon, China gains a notch of independence against American semiconductor restrictions, and the cost gap between Chinese and American models will continue to widen in favor of the former. Huawei presented last week an infrastructure capable, according to its claims, of scaling up to a million processors.

SpaceXAI launched this morning **Grok 4.7**, presented as its best model for code and intellectual work, served **at the same price and speed as Grok 4.6**: $2 per million tokens for input, $6 for output. Under the hood, a larger base model, around **2,100 billion parameters** according to the press, and an extended reinforcement learning phase on tasks that require several hours of work. It is a text-only model, available immediately.

**A few key figures:**

- In-house measurements: **46.3% on CursorBench 4.0** versus 40.4% for Grok 4.6, **38% on Terminal-Bench 4.0** versus 20.3%, 71% on DeepSWE v1.1 and 64% on EEBench

- Independent measurement by Artificial Analysis, relayed by The Decoder: Intelligence Index of **46 for Grok 4.7, versus 53 for Claude Fable 5.1 and GPT-6 Astra**, and above all **26% on Terminal-Bench 4.0**, far behind GPT-6 Astra (60%) and Claude Fable 5.1 (55%), and even below DeepSeek V4.1 Flash (27%)

- New safety stack: **62.4%** on the LatchBio biosafety benchmark, and only **3.3%** of risky cyber requests let through on HackerBench v0.3

- Available in **Cursor**, Grok Build, the Grok API, model routers and cloud platforms, with a variant twice as fast at double the price

In practice, you can plug it into Cursor today without paying a cent more than yesterday, and for coding assistance, the price/performance ratio remains a real argument. But the gap between the figures in the press release and those measured by a third party is a useful reminder of a rule worth keeping in mind: on agent tasks, the only test that matters is yours, not the launch-day chart.

Apple is paying **250 million dollars** to settle a class action that accused it of having sold iPhones on the promise of an AI-boosted Siri, announced in June 2024 with Apple Intelligence and never delivered on time. The initial amount, set at 95 million in May 2025, was revised sharply upward. Since Sunday, the claims website has been open.

**Key points:**

- Eligibility: US buyers of an **iPhone 15 Pro, 15 Pro Max, 16, 16e, 16 Plus, 16 Pro or 16 Pro Max** purchased between **June 10, 2024 and March 29, 2025**, original purchaser, personal or professional use

- **25 dollars per eligible device**, an amount that can rise to as much as **95 dollars** if few claims are filed, and decrease if they pour in

- Process: online form with name, contact details and the iPhone's serial number (or Apple ID and phone number if the serial number cannot be found), payment by check, PayPal, Venmo or bank transfer

- **Filing deadline: December 21, 2026**, confirmed by ConsumerAffairs. The settlement does not constitute an admission of wrongdoing

No French buyer is concerned; the mechanism is strictly American. What matters here is the precedent: a broken promise of an AI feature has just been given a price, quantified, paid per device sold. Every manufacturer now selling hardware on the promise of "upcoming" AI capabilities has an order of magnitude for assessing the risk, and Siri remains, two years after the announcement, an incomplete feature.

# 🧠 **RESEARCH**

### **MiMo-V2.6-Pro: Xiaomi releases the best open weights model in the world**

The electronics and electric vehicle manufacturer lands **46 on the Artificial Analysis Intelligence Index**, tied with Grok 4.7 released the same day, ahead of Gemini 3.8 Flash and DeepSeek V4.1. The model is under an **MIT license**, freely downloadable on Hugging Face and runnable on your own hardware, with **1 million tokens of context** and text, image, audio and video inputs. Via API, count on $0.435 per million input tokens and $0.87 for output, versus $0.14 and $0.28 for the Flash version. A Pro-UltraSpeed variant promises generation up to 20 times faster.

**The model that arbitrates AI agents' decisions flips with a single sentence**

Jev, TypeSafe's decision model adopted within days by Vercel, Cloudflare, LangChain and Langfuse, writes nothing: it rules (allow an action, choose a tool, classify an input) in 70 to 500 milliseconds for $0.042 per million tokens. The problem is that it treats the data it is given as trustworthy: an Octomind test drove the probability of blocking the destructive command `rm -rf ~/.ssh` down from **0.76 to 0.48** by slipping in a fake mention of prior approval. LangChain now excludes tool outputs from what the classifier sees, so that content retrieved by the agent cannot self-authorize, and Pydantic insists: this type of guardrail complements deterministic checks, it never replaces them.

**RetroChimera: Microsoft's AI that writes the recipe for your molecules**

Published in Nature and released as open source, code and weights included, RetroChimera automatically proposes synthesis routes for making a target molecule from purchasable components. It merges two complementary retrosynthesis models and learns to rank their proposals, which makes it better than each taken separately. In a blind test, PhD-level chemists preferred its predictions to those of earlier models, and even to reactions published in the scientific literature. It recovers rare reaction types and transfers without retraining to proprietary datasets.

**The "ChatGPT moment" for robot brains may come in 2027**

The founder of Chinese startup Spirit AI, which develops the embedded models serving as the brains of humanoid robots, places the breakthrough equivalent to what ChatGPT was for text as early as next year. If the prediction holds, humanoids would move from scripted demonstrations to machines capable of understanding a real scene and acting on it. China is putting resources into the topic and openly aims for leadership against American players.

**Underwater drones to cover 98% of the oceans by 2028**

Ifremer has validated two prototypes of Deep-6000 floats capable of descending to **6,000 meters**, where the pressure is 600 times that at the surface, making France the third country in the world to reach these depths after China and the United States, following five successful dives between January 11 and February 2. Thirty abyssal floats will join the international Argo fleet and its 4,000 devices by 2028. The blind spot they fill is significant: **10% of ocean warming occurs below 4,000 meters**, which most current floats do not reach. Each device weighs 40 kg, lasts seven years without recharging, and follows a ten-day cycle: nine days of drifting at depth, then three hours of profiling.

**Amazon wants to design antibodies by computer**

Amazon Bio Discovery publishes three works on the design of therapeutic antibodies, including **MochiBind**, built on the protein language model ESM-2, which compares the binding strength of two antibodies instead of predicting an absolute value that is difficult to compare from one experiment to another. A third work entrusts an AI agent with choosing binding sites and has produced antibodies validated in the laboratory against a novel cancer target. The goal: to shave off the six to twelve months usually needed to go from a biological target to a drug candidate. Amazon also points out the weakness of current benchmarks, reliable on known targets, much less so on new ones.

**What if we approved medical AI like a drug?**

Researchers at the University of Bristol propose a framework called "Learning Ensemble" to evaluate AI systems used in medicine, modeled on drug approval procedures. Three axes of verification: the system's limits, fairness between patient groups, and suitability to the real clinical context. The idea is to catch models that shine on paper and make serious mistakes at the patient's bedside. The authors' central argument: medicine has long known how to regulate mechanisms it does not fully understand.

**Grok Voice Transcribe 2.0: twice as accurate, at the same price**

xAI doubles the accuracy of its transcription model without touching the price: **$0.10 per hour of audio** in batch processing, $0.20 in streaming. It is trained on real, noisy, multilingual recordings, and targets difficult audio: on short voice commands, the word error rate drops from **20.6% to 6.8%**. It handles dozens of languages with automatic detection and switching during recording, and ranks first in accuracy among 32 streaming models on Artificial Analysis's public leaderboard. Existing API integrations get the gain without a single line of code to modify.

**Compressing an LLM like a physicist**

Multiverse Computing reformulates the removal of entire blocks from a transformer, the most brutal method for speeding up a model, as an Ising spin glass problem. The starting intuition: the effect of removing one block depends on the other blocks removed, so scoring them in isolation, as classical methods do, is a reasoning error. The system's energy then serves as a fast proxy for predicting the compressed model's performance without rerunning benchmarks. Announced result: **nearly 23 MMLU points** gained against the best competing method on a Llama-3.3-70B-Instruct compressed by half.

**Valued at $1.4 billion before publishing a single model**

Naive AI, founded in February in Beijing by Jifeng Dai, a professor at Tsinghua University, is said to have raised about $400 million across three rounds with **fewer than 100 employees** and no model published to date. Its strategy stands out from that of its competitors: not pre-training from scratch, but starting from an existing Chinese open-weight model to concentrate its resources on midtraining, post-training, reinforcement learning, and architecture. An open-weights model could be released as early as this month.

**DeepSeek-V4.1-Flash compresses its memory to 890 bytes per token**

A multimodal MoE architecture with **552 billion parameters**, of which only 16 are activated during decoding and 8 during prefill, a one-million-token window, and pre-training on 45 trillion multimodal tokens. The main contribution lies elsewhere: the very aggressive compression of the attention cache, reduced to **890 bytes per token**, which brings down the memory cost and bandwidth required for agents working on very long documents.

**MiniMax-H3 stumbles on physical reasoning**

Tested on 517 cases where text, image, video, and audio each provide only part of the clues about a physical event, the model achieves only **41.97% success**. It does best on decision-making from video (56%) and collapses on audio disambiguation (27.40%). Translation: multimodal models know how to read each channel separately, they still do not know how to piece the parts back together well.

# **🗞️MORE NEWS**

### **The UN warns that there is no guarantee of human control over AI agents**

The UN's scientific panel on AI publishes its first thematic report, and the conclusion is blunt: nothing ensures that humans will retain control of AI agents. Its co-chair Yoshua Bengio details an incident involving OpenAI on Hugging Face, where three critical ingredients came together for the first time: a misaligned objective, the concrete ability to pursue it, and an environment that allowed it. The report also highlights that the most advanced systems could increasingly recognize that they are in a testing situation, and deliberately circumvent safeguards at that precise moment.

**A new arms race is playing out at the bottom of the oceans**

After the Nord Stream pipeline explosions in 2022 and a series of incidents in the Arctic and near Taiwan, underwater infrastructure has become an accepted target of hybrid warfare. More than **1.5 million kilometers of submarine cables** keep the global economy running, and no one is really watching them. Advances in AI, drones, and autonomous underwater vehicles are opening two simultaneous paths: machines capable of patrolling on their own, and sensors attached directly to the cables. For Chatham House, the maritime domain is "the next logical domain" after the massive use of aerial drones in Ukraine, and defense manufacturers are already positioning themselves.

**Higgsfield AI goes from prompt to production in one day with GPT-6 Astra**

OpenAI is highlighting Higgsfield AI, which relies on GPT-6 Astra to create video ads for small businesses with no production expertise. The argument being made is not the quality of the output but the pace: the company claims to deliver new creative tools in a single day of development. It is a good indicator of what the latest models are changing in practice for small product teams, where the idea/prototype/launch cycle is now measured in hours.

**Muse already surpasses ChatGPT's mobile beginnings**

According to estimates from Apptopia, Meta's app was downloaded **1.8 million times on iOS** in the United States and Canada during its first 12 days, compared with 1.3 million for ChatGPT at the same stage. The gap is even more pronounced in actual usage: **642,000 daily active users** in the United States, compared with 231,000 for ChatGPT at the time, and 2.8 million installations worldwide. Meta owes a great deal to its own pipelines, since more than 95% of Muse users are also on Facebook, exactly the method that had propelled Threads. It remains to be seen how many will stay once the novelty wears off.

**OpenAI proposes global standards on AI self-improvement**

OpenAI is publishing a series of safety proposals for frontier models, centered on alignment research and recursive self-improvement, the technique in which a model improves itself without human intervention. The company states that fully autonomous self-improvement should not be pursued as long as it cannot be carried out safely, at the risk of losing human control, and calls for international technical standards supported by existing safety institutes. The publication follows closely on Anthropic's last week, in a climate marked by Dario Amodei's essay calling for a slower pace and by the much-discussed resignation of Jacob Coxon.

**Amazon and Stanford formalize a joint research initiative**

The two institutions announced on September 16 the Stanford and Amazon Research Initiative, a joint framework on fundamental and applied AI, automated reasoning, energy, and health. The partnership did not start from scratch: more than ten Amazon teams already fund work at Stanford, from humanoid robotics to post-quantum cryptography to AI-assisted radiology. The agreement adds three concrete components, joint research projects, doctoral fellowships, and interdisciplinary symposia, with the stated goal of accelerating the transition from the laboratory to real-world applications.

**ByteDance launches Dramagic, a factory for AI-generated short series**

The parent company of TikTok has launched a platform that handles the entire production chain of a short series, from scriptwriting to scene generation and video preview. The target market is already colossal: **128,000 short series** were released in China in the first quarter of 2026, **95% of them AI-generated**. In other words, mass fiction production there has already shifted to near-complete automation, and ByteDance is simply industrializing what was being done with scattered tools.

**DeepSeek bets on Huawei chips to circumvent U.S. controls**

DeepSeek CEO Liang Wenfeng has told his investors that Huawei could begin delivering training chips as early as the **fourth quarter of 2026**. The nuance is important: training a model requires significantly more powerful chips than simply running it, and that is precisely where U.S. restrictions bite the hardest. At the same time, DeepSeek is finalizing a second funding round of about $7.5 billion at a valuation close to $75 billion.

**Tesla deploys FSD Supervised in the Czech Republic**

The Czech Ministry of Transport has granted provisional authorization to Tesla's driver-assistance system, making the country the seventh in Europe to authorize it, after the Netherlands, Lithuania, Estonia, Denmark, Belgium, and Slovenia. The agency had initially raised concerns about traffic light handling, speed limits, and driver attention monitoring, before relying on the approval issued in April by the Dutch regulator RDW. Tesla is now aiming for EU-wide authorization, which would require the support of at least 15 of the 27 member states in a vote envisioned as early as October. France, for its part, is still in the testing phase.

**Texas freezes all environmental permits for data centers**

Governor Greg Abbott has ordered the state environmental agency to suspend all permits related to data centers pending an audit of the power grid interconnection queue by ERCOT, extending the moratorium decided in August. BloombergNEF estimates that nearly **50 gigawatts of projects**, or about 20% of the U.S. pipeline, are threatened, with up to $8 billion in lost revenue by early 2027. The issue has become frankly electoral: **64% of Americans surveyed** say they are less inclined to vote for a candidate who supports building data centers near their homes. Anthropic is even considering listing public hostility as a risk factor in its IPO prospectus.

**OpenAI and Anthropic negotiated an agreement to test each other's models**

The two companies were close to a legally binding agreement allowing them to submit their respective models to security stress tests. These discussions, never made public until now, predate the series of cybersecurity incidents involving OpenAI's technology. The company is currently reviewing all of its security strategies, under pressure from employees who are publicly warning about the risks of its systems.

**OpenAI establishes an independent advisory group on mathematics**

The company is partnering with an external advisory group tasked with overseeing the evaluation and communication of mathematical results obtained by AI. The context explains everything: announcements of mathematical breakthroughs are multiplying across the industry, without many people being able to verify them. The implicit goal is therefore to lend credibility to future announcements by limiting unverifiable claims, in a field where verification by competent peers is the only currency that matters.

**AMD crosses the $1 trillion market cap threshold for the first time**

The stock jumped 10% on Monday, reaching an intraday record of $615.52, at the end of five sessions of gains and a rise of more than 180% since January. The driver is clearly identified: $11.54 billion in quarterly revenue, up 50% year over year, and sales in the Data Center division that have **more than doubled** (+107%, $6.7 billion in the second quarter). AMD nonetheless remains far behind Nvidia and its roughly $5.4 trillion market cap. Lisa Su plans to double data center sales again in 2027, while Jensen Huang announces he wants to double the number of chips sold next year.

And a skill that everyone has doesn't get billed. Yet there remains a demand that ChatGPT cannot serve: an accounting firm will never put its client files in it. Neither will a lawyer, and even less so a medical practice. They want AI like everyone else, but without their files leaving the office.

It exists, and it's called local AI: the model runs on their own machine, and it even responds with the wifi turned off. Same for image and video, with no credit limits. Few people know how to install it, and no one teaches it in French. A rare skill, on the other hand, gets billed.
