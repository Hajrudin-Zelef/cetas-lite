---
id: labs-hyperscalers-2026/00-labs-hyperscalers/7-8-trainium-inferentia-2026-detailed
title: "7.8 Trainium / Inferentia (2026) — detailed"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "Nvidia", "OpenAI", "TSMC", "United States"]
dates: ["2025-03", "2025-12", "2026-02-04", "2026-03", "2026-03-19", "2026-07", "2026-07-27", "2026-07-30", "2026-07-31"]
keywords: ["trainium", "3nm", "accelerator", "agent", "agentic", "agents", "agi", "aws", "bedrock", "benchmark", "blackwell", "capex"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1037, 1082]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 6abefc1eff764a0dcc89254e1cdecb0ce4b5b8cf7f8cc2d4654bd38832f16542
---

# 7.8 Trainium / Inferentia (2026) — detailed

### 7.8 Trainium / Inferentia (2026) — detailed

- **Trainium3** — announced at re:Invent 2025; **generally available in EC2 Trn3 UltraServers since December 2025**; first 3nm AWS accelerator (TSMC, CoWoS-L); **2.517 MXFP8 PFLOPS/chip**, 144 GB HBM3e, 4.9 TB/s memory bandwidth, ~750W TDP (vs. Nvidia H200 700W, B200 1,000W); supports MXFP8/MXFP4 and dense + expert-parallel topologies [official specs via press].
- **Trn3 UltraServer:** 144 chips, up to **362 FP8 PFLOPs**, 20.7 TB pooled HBM3e, ~706 TB/s aggregate bandwidth (NeuronLink); up to **4.4×** performance, 3.9× memory bandwidth, ~4× perf/watt vs. Trainium2 UltraServer [vendor-reported]. UltraClusters 3.0 for scale-out to **1M chips** [vendor-reported].
- **Pricing:** AWS quoted internal customer pricing ≈ **$1.80/chip-hour** for Trainium3 (early 2026) vs. on-demand H200 ≈ **$4.80/chip-hour** [secondary]. AWS claims up to **50% lower training/inference cost** vs. alternatives and 30–40% better price-performance than Trainium2 [vendor-reported]. **No independent per-token benchmark exists as of July 2026**; SemiAnalysis InferenceMAX listed Trainium2/3 as "being added." Neuron 2.30 dropped PyTorch/XLA training — migration cost for custom CUDA kernels is real [independent analysis].
- **Adoption:** **Anthropic trains Claude on Trainium2** through **Project Rainier**: cluster of **500,000+ chips** [vendor-reported]. **Uber** picked Trainium3 (reported as "50% cheaper than Nvidia") [secondary, Sept 17, 2026] https://tech-insider.org/uber-aws-trainium3-amazon-ai-chip-deal-2026/. Also named: Karakuri, Metagenomi, NetoAI, Ricoh, Splash Music, Decart [vendor-reported]. Most Amazon Bedrock inference now runs on Trainium (125,000+ Bedrock customers) [vendor-reported].
- **Coexistence with Nvidia:** Reuters (March 2026): Nvidia to sell **1M GPUs to AWS by end of 2027** as part of a wider cloud deal; AWS also launched SageMaker HyperPod **P6-B200** (NVIDIA B200) instances (July 27, 2026) [secondary].
- **Rack-scale comparison:** Tom's Hardware (Dec 2025) noted the Trn3 UltraServer's **0.36 ExaFLOPS of FP8** performance matches Nvidia's **NVL72 GB300** rack-scale system [secondary]. Per-chip, however, Trainium3 (2.517 PFLOPS FP8) remains well below a single Blackwell B200 (~10 PFLOPS FP8 equiv., 192 GB HBM3e, 8 TB/s) [secondary]. Fabric-throughput comparisons are modeled estimates cited by secondary sources; treat as approximate.
- **Trn3 Gen2 UltraServer:** AWS reaffirmed July 2026 (144 accelerators, NeuronLink interconnect) [secondary].
- **OpenAI commitment:** as part of the Feb 2026 partnership, OpenAI committed to **2 GW of Trainium capacity** for training workloads [secondary].
- Inferentia: no 2026 Inferentia news surfaced.

### 7.9 Amazon AI labs, Alexa+, consumer AI — detailed

**AI labs:**
- **Amazon AGI SF Lab** (opened Dec 2024 — context): SF-based lab seeded by hires from Adept, focused on AI agents acting in digital and physical worlds; led by **David Luan** (Adept co-founder, VP of Autonomy) with Pieter Abbeel (Amazon Scholar, robotics) closely involved [official].
- **Personnel change (2026):** GeekWire (July 2026) reports David Luan **has since left Amazon** [secondary]. Nova Act co-developer Valliappa Chinnapaiyan Talluri also departed to found Primitive Labs [secondary].
- **Lab126 "Physical AI":** new agentic-AI group within Lab126 developing an agentic AI framework for robotics — robots that "hear, understand and act on natural language commands"; targets warehouse robots (unload trailers, retrieve parts) [secondary — SiliconANGLE].

**Alexa+ (consumer AI):**
- **US general availability: February 4, 2026** (after year-long Early Access from March 2025 preview); free for Prime members; **$19.99/month** standalone [secondary].
- Stack: Anthropic's Claude (complex reasoning/queries) + Amazon Nova (routing/simpler tasks) + Annapurna Labs **Trainium 2**; powered by **Amazon Bedrock** with 70+ LLMs, automatic model selection [vendor-reported][secondary].
- **UK rollout:** invite-only early access from **March 19, 2026**; hundreds of thousands invited; partners include OpenTable, JustEat; after early access: Prime benefit or **£19.99/month** [secondary].
- **India launch: September 16–17, 2026** (early access): Hindi + Hinglish support (English, Hindi, Hinglish; more Indian languages to follow); developed by teams in Bangalore, Hyderabad, Pune, Chennai; integrates Swiggy, Zomato District, TripAdvisor, MakeMyTrip, EazyDiner, Amazon Now; after early access: free for Prime, **₹2,000/month** standalone for non-Prime [secondary] https://www.thehindu.com/sci-tech/technology/amazon-launches-gen-ai-powered-alexa-in-india/article71471832.ece.
- **New Alexa+-optimized devices (Sept 2026):** Echo Show 8, Echo Show 11, Echo Dot Max, Echo Studio — enhanced processing, edge-compute capabilities [secondary].
- Scale: **600M+** Alexa-enabled devices worldwide; 97% compatible with Alexa+; 2–3× more conversations and 3× more purchases reported for users [vendor-reported][secondary].

### 7.10 Amazon investments, partnerships, personnel, data centers (2026) — detailed

**OpenAI partnership + $50B investment:**
- **Announced Feb 27, 2026; completed July 31, 2026** (SEC filing): Amazon invested **$50B** in OpenAI — $15B in Q1 (initial), $13.7B in Q2, $21.3B after June 30. Initial $15B was for preferred stock; $35B tranche conditional on milestones (The Information: tied to IPO or AGI conditions; Amazon did not specify why it released the remainder) [secondary] http://pymnts.com/news/artificial-intelligence/2026/amazon-completes-50-billion-dollar-investment-openai/.
- Terms: **AWS becomes exclusive third-party cloud provider for OpenAI Frontier** (OpenAI's enterprise agent platform); joint development of a **Stateful Runtime Environment** on AWS, integrated with Bedrock AgentCore (expected to launch "within the next few months" of Feb 2026); OpenAI committed to **2 GW of Trainium capacity**; expanded prior AWS infrastructure agreements **up to $100B over 8 years** [secondary].
- Valuation inconsistency: round reported as $110B at $730B pre-money (FT) vs. $840B pre-money (aiforautomation); OpenAI later closed $122B at $852B post-money [secondary]. Flag when reusing.

**Anthropic relationship:**
- Project Rainier (Indiana): Anthropic trains Claude on 500,000+ Trainium2 chips; site described as "fully operational" in 2026 [vendor-reported][secondary].
- Circular-deals reporting (aistockwire, Sept 2026): Amazon has **$13B invested** in Anthropic with up to **$20B more**; Anthropic has **$100B+ AWS commitment over 10 years** [secondary][unverified] — only partially corroborated; note alongside Microsoft's Nov 2025 $5B/up-to-$30B-Azure deal.
- Parallel track: Amazon committed up to $33B to Anthropic days before Google's Apr 24 $40B announcement [secondary].

**Capex and data centers:**
- **2026 capex:** announced **~$200B** by Jassy on **Feb 5, 2026** (Q4 2025 earnings); **raised to ~$220B on July 30, 2026** (Q2 earnings), citing rising memory-chip prices and unmet AWS demand [secondary].
- **AWS AI revenue run rate: $15B in Q1 2026** (~10% of AWS's $142B annualized revenue), per Jassy's shareholder letter [vendor-reported][secondary].
- **Capacity:** Amazon added **3.8–3.9 GW** of data-center capacity in the past year (outpacing competitors, per Jassy); aims to **double capacity by 2027** [vendor-reported][secondary].
- **$50B federal AI data-center investment** (announced Nov 24, 2025 — context; construction begins 2026): up to $50B for AI/supercomputing across AWS Top Secret, Secret, and GovCloud regions; **+1.3 GW** of compute; 11,000+ US government agencies served [secondary via Reuters].
- **India:** Jassy committed **$48B to India 2026–2030** (custom AI accelerators, AWS infra in Mumbai and Hyderabad) [secondary]; separate reporting cites an **$8.3B** AI/cloud infrastructure investment in India [secondary] — figures may overlap; do not sum.

