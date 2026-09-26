---
id: collect-240926-vision-ia/vision-ia/les-eleves-qui-deleguent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-ecole-2
title: "🧠 **RESEARCH**"
domain: vision-ia
role: reference
task: reference
actors: ["Anthropic", "Apple", "China", "Google", "Huawei", "OpenAI", "Samsung", "United States"]
dates: []
keywords: ["research", "alignment", "apache", "benchmark", "compute", "cost", "embeddings", "energy", "foldable", "gpu", "gqa", "inference"]
source: docs/RAG/clean_en/vision-ia/les-eleves-qui-deleguent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-ecole.md
source_anchor: ""
source_lines: [89, 141]
sha256: 174acfde4b85b2fcb68c6f86ce3ac55fe84fd4d9a2ab76625bba61a7f9cbae7b
---

# 🧠 **RESEARCH**

The iPhone Duo unveiled at the "Surprise and Shine" keynote is Apple's first foldable phone, with a **7.6-inch internal display** and a **5.4-inch external display** usable when folded. But the most interesting part is invisible: the hinge is designed and manufactured using AI algorithms, unit by unit, on the production line.

According to Johny Srouji, Apple's hardware chief, **AI algorithms pair each individual hinge with the housing best suited to it** , to guarantee perfect alignment despite manufacturing tolerances.
A **confocal laser scans the topology of each unit** , then a printer deposits**up to 25 micro-layers of a custom photopolymer** to eliminate residual waviness. Each device therefore receives a correction unique to it.
The problem being addressed is well known: a foldable undergoes **twice the mechanical wear** of a conventional smartphone, and that's where the competition has come unstuck.
Added to this are a **nano-texture anti-reflective** finish and a multilayer lamination strategy.
Pricing and schedule: **€2,339** for 256 GB, €2,589 for 512 GB, €3,089 for 1 TB and €3,839 for 2 TB, preorders on**October 16 at 2 p.m.** , release on**October 23** (detailed pricing).

Apple is arriving years after Samsung and Huawei in this market, and is betting everything on durability as a differentiator. The notable point for anyone following AI is not the phone, it's the shifting of the needle: AI is leaving software to enter precision micro-manufacturing, with customized correction for each unit coming off the line. Few manufacturers communicate about this use. The usual unknown remains: Apple publishes no figure for guaranteed folding cycles, nor any durability comparison with competing foldables. We'll have to wait for the first aged devices to decide.


# 🧠 **RESEARCH**

### **A developer single-handedly trains a 3.8-billion-parameter model for $998**

Hugo Vergnes trained from scratch a **3.8-billion-parameter** LLM on **65 billion tokens**, in **43 hours of rented GPU**, for a total cost of **$998**. Result: **0.384 on the CORE benchmark**, above Karpathy's nanochat at a comparable budget. Two practical takeaways: rented B200s offer better cost-per-work than H100s for this type of run, and the author built little-lm, a framework where each experiment fits in a few lines of YAML. The architecture is Llama-type, with RMSNorm, RoPE, GQA, QK-norm and ResFormer-style value embeddings. The zone between the educational toy and the research lab is more accessible than one might think.

**IBM opens a time series forecasting model under commercial license**

Granite Time Series PatchTST-FM-r2 is a **385-million-parameter** foundation model that produces forecasts with no prior training on your data. It handles a context of **8,192 points**, outputs probabilistic forecasts via **99 quantiles** and can fill in missing values. As of September 8, it is **2nd in the GIFT-Eval ranking** among reproducible zero-shot models, and **1st among those under a permissive license**. Weights, architecture and inference code are published under a dual Apache 2.0 and OpenMDW 1.0 license, so usable in production: demand forecasting, price forecasting, energy consumption or traffic.

**Anthropic publishes an economic model that ranks its own CEO's predictions in the extreme scenario**

Anthropic has modeled three trajectories for the US economy through 2030. In the most violent scenario, output **doubles every 4.5 years** and knowledge-worker unemployment reaches **17.9%**. Yet these are precisely the figures Dario Amodei publicly put forward last May. His own lab has therefore just filed its leader's prediction under minority hypotheses, which says a great deal about the gap between the sector's public messaging and its internal models.

**Do readers prefer stories written by an AI?**

Cambridge University Press publicly poses the question, which is becoming serious as models produce readable fiction. At this stage, only the title of the post is accessible: neither the methodology nor the results of the study can be consulted, and it is better to wait for the full publication before drawing a conclusion. The subject to watch remains that of perception, since the same text is not judged the same way depending on whether or not one believes a machine wrote it.

**Batteries set a new installation record in the United States**

**20.2 GWh** of storage installed in the second quarter of 2026, the equivalent of the daily consumption of **700,000 households**, with seven projects of more than one gigawatt-hour brought online. The country is on a trajectory of **71 GWh for the year**, up 20%. Data centers account for about **three quarters** of new batteries in the commercial segment, while residential falls **16%** after the end of a tax credit. Almost all cells still come from China.

# **🗞️MORE NEWS**

### **The Apple Watch now listens continuously, and transcribes**

Three features are arriving on the watch: **Audio Intelligence**, which detects sirens, alarms, doorbells and baby cries **on-device, without an iPhone nearby**, **Live Rewind**, which transcribes the **last 15 seconds** of a conversation on a double press of the Digital Crown and saves the text in the new Siri app, and Siri Recap. Sound detection clearly falls under accessibility. Retroactive transcription, much less so: it raises the question of your interlocutor's consent, who does not know they have just been transcribed. Coming from the company that built its marketing on privacy, the shift is notable, and it looks like an anticipated response to the passive listening devices from OpenAI, Friend or Bee.

**A protein specialist dismantles the AI-designed supervirus scenario**

Economist Noah Smith published a piece imagining a teenager capable, in 2029, of getting an AI to design a virus killing 90% of humanity. A researcher who specifically designs AI-assisted proteins and peptides responds to him point by point. His central argument is empirical: in 2026, doctoral students equipped with the best tools spend **months** designing simple peptide binders, and most of their attempts fail. Then comes the physical obstacle, because designing is not manufacturing, and "you don't order a virus on temu.com." The author does not deny any biological risk linked to AI, he disputes the scale. He notes in passing that Smith's piece is paywalled and closed to comments, and therefore sheltered from contradiction.

**Apple wants to prove your photos are not generated images**

With **Apple Reference Image**, the iPhone 18 Pro captures **signed sensor data** at the moment of shooting, which Private Cloud Compute turns into an unalterable image viewable in the Photos app. It is a digital negative: any modified version can be compared to the original to detect retouching. Apple will open **APIs to developers** to integrate this verification elsewhere, and will adopt the **SynthID** standard to flag images created or modified by AI. The company explicitly targets photojournalists, but the obvious use is broader: insurance, expertise, litigation, anything that relies on a photo that must be believable.

**Apple's Health app calculates your "health age"**

A complete overhaul, powered by Apple Intelligence. An **Insights** tab gathers personalized advice and assessments, with a daily **readiness** score and a **Health Age** calculation that compares your metrics to your real age, based on VO2 max, sleep and blood biomarkers. A **Longevity** tab evaluates sleep, activity and heart health, integrates your medical test results and comes with a partnership with Quest for a panel of **50 biomarkers at $119**. The recommendations are concrete, such as adding intervals to your morning run. Rollout later this year, first in American English in the United States, with the Apple Watch Series 12 and Ultra 4.

**Google DeepMind reconstructs a 70-year-old memory in a short film**

