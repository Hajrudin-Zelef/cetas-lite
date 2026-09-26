---
id: collect-240926-vision-ia/vision-ia/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors-1
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Apple", "China", "DeepSeek", "Google", "Huawei", "Meta", "Microsoft", "OpenAI", "Perplexity", "SpaceX", "United States", "Xiaomi", "xAI"]
dates: ["2024-06", "2024-06-10", "2025-03-29", "2025-05", "2026-12-21"]
keywords: ["research", "accelerator", "agent", "agents", "astra", "benchmark", "chatgpt", "claude", "compute", "cost", "cyber", "deepseek"]
source: docs/RAG/clean_en/vision-ia/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors.md
source_anchor: ""
source_lines: [1, 90]
sha256: 16b236244b3b4fc05e796084d556bd82efd180565fe61358dd6721d76ea24b7d
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

