---
id: labs-hyperscalers-2026/00-labs-hyperscalers/4-10-fair-meta-ai-research-milestones-2026
title: "4.10 FAIR / Meta AI research milestones 2026"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["AMD", "Alibaba", "Broadcom", "China", "Google", "Hugging Face", "Meta", "Microsoft", "Nvidia", "TSMC", "United States"]
dates: ["2025-04", "2025-07", "2026-03", "2026-04-08", "2026-04-29", "2026-07-09", "2026-07-28", "2026-09", "2026-09-18"]
keywords: ["research", "accelerator", "agent", "amd", "benchmark", "blackwell", "capex", "compute", "cost", "distillation", "embeddings", "energy"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [742, 777]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 43e341f9ff573a355d5a28cb8c16e21984b6979ad1a183c683db8330fae56ee8
---

# 4.10 FAIR / Meta AI research milestones 2026

- Zuckerberg published a **~6,500-word essay, "The Future is for Everyone,"** on Meta's website on Aug 10, 2026, simultaneous with the Muse Glimmer release [secondary — Reuters Breakingviews; The Guardian].
- Core arguments: everyone should have a personalized "superintelligence" agent; concentrated AI control (by companies, governments, or AI itself) is the biggest risk; **open weights "will be the best way to protect safety and security over time"**; US should loosen restrictions on AI training data to keep US open-weight labs competitive with China; build energy capacity faster [secondary].
- Practical commitments: independent safety review board to approve safety rules before future model releases; **$1B community fund** for communities near Meta data centers; 2026 infra spend $130–145B; offer to share some internal AI training data with officials [secondary].
- Context: shift from his July 2025 stance that Meta likely wouldn't open-source all superintelligence models; follows the Jan 2026 Llama 4 benchmark controversy (Zuckerberg acknowledged the LM Arena test-version vs. shipped-version mismatch — specifics [unverified]) [secondary/unverified].
- The letter's "exceptionally capable personal agent for everyone" and "fully private mode where even Meta cannot see data" do not exist in any shipped Meta product — quote the letter as direction, not capability [secondary].

### 4.10 FAIR / Meta AI research milestones 2026

- **"Beyond Language Modeling" (March 2026)** — large-scale study on building native multimodal foundation models from scratch (no language-pretrained backbone), using the Transfusion framework; findings: RAE encoder built on SigLIP 2 beats VAE encoders/raw pixels; mixed text+vision training is synergistic; MoE architectures naturally specialize experts by modality [secondary].
- **JASCO** — joint audio-and-symbolic conditioning model for temporally controlled text-to-music generation; paper + sample page released; inference code slated for AudioCraft repo under MIT [official — Meta AI blog].
- **NeuralSet (April 29, 2026)** — Python package unifying neural recordings (fMRI, M/EEG, spikes) with deep learning: single PyTorch-ready DataLoader, HuggingFace embeddings support [secondary].
- **AI Research Preference Models / RPMs (Sep 6, 2026)** — FAIR + Oxford + UCL paper: RPMs rank unexecuted ML-experiment candidates before spending GPU hours; scaffold and AIRS-Bench open source; backbone Qwen3.6-27B [secondary].
- **SAM 3** — next-gen "Promptable Concept Segmentation" (2× claimed gain over prior systems); research paper submitted to ICLR 2026, models NOT yet publicly released [unverified].
- **Byte-model fact-check (Sep 2026)** — a viral X thread misrepresented Meta's Dec-2024 Byte Latent Transformer as new 2026 work; the actual new paper (UW + Meta FAIR: "Breaking the Token Ceiling") shows byte models keep improving with compute and can exceed token models' ceiling under distillation with ~1/6th the data [secondary].

### 4.11 Meta infrastructure: capex, chips, data centers

- **Capex guidance 2026:** initially $115–135B; raised (April) to **$125–145B** including finance-lease payments; Q2 reporting raised the floor again to **$130–145B**. ~70% of capex is now AI infrastructure. Q2 2026 revenue **$60.8B (+28% YoY)**; profit missed expectations; shares fell on margin concerns [secondary].
- **Compute targets (Reuters, July 9, 2026, internal comms):** ~7 GW of compute capacity by end of 2026 (1 GW deployed H1 2026 + 2.5 GW planned H2), doubling to **14 GW in 2027** [secondary].
- **MTIA "Iris":** next MTIA accelerator reportedly begins production **September 2026**; co-developed with **Broadcom**, manufactured by **TSMC**; plan to refresh chips roughly every 6 months through 2027 [secondary — Reuters via DataFLOQ].
- **AMD deal (Feb 2026):** up to **$60B over 5 years** for AMD AI chips, per Reuters reporting (also described as "parallel 6 gigawatts of AMD GPUs") [secondary — primary Reuters URL not directly fetched].
- **Nvidia:** remains Meta's primary GPU supplier; training compute reportedly on 500,000+ Blackwell B200-class GPUs [unverified]; Meta also rents large Google TPU clusters [secondary/unverified].
- **BlackRock JV (July 28, 2026):** 1-GW AI data-center campus in **El Paso, Texas**, ~$14B total cost; BlackRock funds hold 80% equity, Meta 20%; capacity online from 2028. Separate 1-GW facility in **Alberta** announced alongside [secondary].
- **Tulsa:** $1B AI-optimized data center announced (1,000+ construction jobs, ~100 permanent); water-efficient cooling, clean-energy matching [secondary].
- **Corning (Jan 2026):** multi-year **$6B** agreement for fiber-optic cables for data centers [secondary].
- **"Tents":** Meta experimenting with lightweight rapidly-deployable data-center structures — permanent facilities take 12–24 months vs. weeks for tents; signal is speed-to-capacity [secondary].
- Power strategy: direct gas-turbine deals, long-term nuclear PPAs, on-site battery storage; behind-the-meter generation toward GW-scale campuses [secondary].
- IDC (via secondary): global AI infrastructure spend projected **$497B in 2026 (+56% YoY)** [secondary].

### 4.12 Llama status — detailed (from Track B)

- **Last verified Llama-family release:** Llama 4 Scout / Maverick, April 2025. No Llama 4.x point release verified in Feb–Sep 2026.
- **Llama 5: UNVERIFIED — conflicting claims. Do not treat as a shipped product fact.** Several low-quality aggregator pages claim Llama 5 launched **April 8, 2026** at a "Meta AI Connect summit" with "600B+ parameters" and "5M-token context," citing a CNBC article whose actual title references the Muse Spark line, not a Llama launch [unverified]. A separate aggregator narrative ("Llama 5 Avocado," Aug–Sep 2026) says Meta's next open model project finished basic training in early 2026 but launch was pushed back [unverified]. A structured Meta-Llama tracking page (verification stamp 2026-09-18) states explicitly: *"No Llama 4.1, 4.5 or 5 exists on any first-party source"* [secondary].
- **AI pricing tracker (Aug 10, 2026 checkpoint):** Meta announced it will *resume* open-source model releases "soon" but published no model name, checkpoint, license, hardware target, or endpoint; hosted Llama API pricing remains third-party inference providers only — Meta publishes weights, not a first-party token price [secondary] https://www.aipricing.guru/meta-llama-pricing/.
- **Meta AI app / Reality Labs:** Meta AI app and meta.ai added "Thinking" mode powered by Muse Spark 1.1 (July 9, 2026) [vendor-reported/secondary]. No verified 2026 Reality Labs AI-specific launches found in fetched sources. **Meta Connect 2026 scheduled Sept 23–24, 2026** — after this report's cutoff; re-verify post-Connect [secondary].

