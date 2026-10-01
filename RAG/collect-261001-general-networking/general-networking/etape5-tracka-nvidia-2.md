---
id: collect-261001-general-networking/general-networking/etape5-tracka-nvidia-2
title: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "CoreWeave", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia", "OpenAI", "Oracle", "TSMC"]
dates: ["2026-06-01"]
keywords: ["gpu", "gpus", "nvidia", "rubin", "agentic", "attention", "aws", "blackwell", "capex", "compute", "cost", "disaggregated"]
source: docs/RAG/collect-261001-general-networking/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [101, 161]
sha256: cab80991031749c0331d7d199bb70337ed19e6a7d0f25d6c012fd5fd715abf50
---

# Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking

### 2.4 Vera Rubin NVL72 rack
- **Configuration:** 72 Rubin GPUs + 36 Vera CPUs, 18 compute trays, 9 NVLink switch trays, ~5,000 copper cables in NVLink spine (over 2 miles) [secondary]
- **Memory:** 20.7 TB HBM4 + 54 TB LPDDR5x; 1.6 PB/s HBM bandwidth [secondary]
- **Performance:** 3.6 EFLOPS NVFP4 inference; 2.5 EFLOPS training [vendor-reported]
- **Scale:** full Vera Rubin POD = 40 racks, 1,152 Rubin GPUs, ~60 EFLOPS NVFP4 [secondary]
- **Production status:** NVIDIA stated Rubin "in full production" in Q1 2026 (per GTC 2026 coverage) — earlier guidance had pointed to mass production in H2 2026; partner availability H2 2026 [independent/secondary — https://videocardz.com/newz/nvidia-vera-rubin-nvl72-detailed-72-gpus-36-cpus-260-tb-s-scale-up-bandwidth]
- **Cloud rollout (H2 2026 confirmed):** CoreWeave (first rack operational June 1, 2026), AWS, Google Cloud, Microsoft Azure, Oracle Cloud, Lambda, Nebius, Nscale [independent/secondary — https://ts2.tech/en/nebius-group-nbis-stock-jumps-after-nvidia-rubin-nvl72-plan-what-investors-watch-next/]
- **Reported customers:** OpenAI (deploying at scale Q3 2026 per Bloomberg reports), Meta, Dell; CoreWeave reported 10× token output vs Grace Blackwell generation [secondary — https://finance.biggo.com/news/202607220220_Nvidia_Vera_Rubin_NVL72_full_production]
- **Europe:** Bull (in partnership with Foxconn) announced European commercial availability, components manufactured in France and the Czech Republic [secondary — https://www.techtimes.com/articles/318651/20260618/nvidia-vera-rubin-nvl72-cloud-rollout-expands-europe-h2-deployments-near.htm]
- ⚠️ One secondary source claims AWS and Google Cloud will brand the Rubin R100 GPU as **"H300"** for catalog continuity — single-source, treat as [unverified]

### 2.5 Rubin CPX (prefill/context GPU)
- **Positioning:** new GPU category purpose-built for the prefill (context) phase of disaggregated LLM inference — million-token software coding, generative video, agentic workloads [independent — https://www.eetimes.com/nvidia-specializes-gpu-for-first-stage-of-transformer-inference/]
- **Silicon:** monolithic single die (departure from NVIDIA's dual-GPU packages) [independent — DCD, EE Times]
- **Compute:** up to 30 PFLOPS NVFP4 [official via press]
- **Memory:** 128 GB **GDDR7** (not HBM) — deliberate cost/throughput trade for context processing [official via press]
- **Media:** 4× NVENC + 4× NVDEC for video workflows [official via press]
- **Claims:** up to 3× faster attention performance vs GB300 NVL72 for long-context processing [vendor-reported]
- **Availability:** end of 2026 / late 2026 [official via press — https://www.datacenterdynamics.com/en/news/nvidia-launches-rubin-cpx-gpu-for-large-scale-inferencing/]
- **Vera Rubin NVL144 CPX rack:** 144 Rubin CPX GPUs + 144 Rubin GPUs + 36 Vera CPUs; 8 EFLOPS NVFP4 AI compute; 100 TB fast memory; 1.7 PB/s memory bandwidth; liquid-cooled MGX system; 7.5× AI performance vs GB300 NVL72 [official via press]
- **Economics claim:** NVIDIA's Ian Buck said $100M capex in CPX racks with Dynamo orchestration could deliver up to $5B in token revenue (30–50× ROI) [vendor-reported — https://www.eetimes.com/nvidia-specializes-gpu-for-first-stage-of-transformer-inference/]
- **Early adopters:** Cursor, Runway, Magic [secondary — https://www.datacenterdynamics.com/en/news/nvidia-launches-rubin-cpx-gpu-for-large-scale-inferencing/]

### 2.6 Rubin Ultra (2027) and Feynman (2028) roadmap
- **Rubin Ultra — 2027:** 600 kW rack-class system previewed at GTC 2026 [secondary — https://letsdatascience.com/blog/jensen-huang-walked-out-with-a-chip-doing-50-petaflops-the-ai-industry-held-its-breath]
  - NVL576 configuration: 576 Rubin-class GPUs [secondary — https://www.tweaktown.com/news/107645/nvidia-rubin-cpx-gpu-to-feature-128gb-gddr7-memory-launches-end-of-2026/index.html]
  - Up to 4 GPUs per module; HBM4E memory; NVIDIA Kyber networking (copper + co-packaged optics) [secondary]
  - ⚠️ One medium source describes "Vera Rubin Ultra" as 144 GPUs in vertically oriented compute trays — **conflicts with NVL576/576-GPU figure; flagged [unverified]**
- **Feynman — 2028 (first public preview at GTC 2026):**
  - Rosa CPU (named after Rosalind Franklin) [secondary]
  - LP40 LPU — next-generation language processing unit, continuing Groq LPU integration [secondary]
  - BlueField-5 DPU, ConnectX-10 (CX10) networking [secondary]
  - NVLink 8, Spectrum-7, co-packaged optics via NVIDIA Kyber [secondary]
  - TSMC A16 1.6nm process [secondary — https://letsdatascience.com/blog/jensen-huang-walked-out-with-a-chip-doing-50-petaflops-the-ai-industry-held-its-breath]

---

## 3. H100 / H200: availability and 2026 price trends

### 3.1 Purchase prices (2026)
- **H100:** $25,000–$40,000 per GPU [secondary — https://intuitionlabs.ai/articles/nvidia-ai-gpu-pricing-guide]
- **H200:** ~$31,000 per GPU; 8-GPU systems at ~$315,000 [secondary]
- A100 40 GB: ~$10–12K; A100 80 GB: ~$15–17K (for context) [secondary]

### 3.2 Cloud rental (2026)
- **H100:** $1.38–$12.29/hr range across providers; on-demand median $2.01–$3.33/hr; spot as low as $1.25/hr; average ~$3.11/hr early 2026 [secondary]
  - SemiAnalysis 1-year contract index (100+ market participants): $1.70/hr (Oct 2025) → **$2.35/hr (Mar 2026)** [independent]
  - Historical: $8–10/hr (Q4 2024 peak) → $5.50–7 (Q1 2025) → $3.50–4.50 (Q2 2025) → $2.85–3.50 (Q3–Q4 2025): 64–75% decline from peak [secondary]
- **H200:** 25–30% premium over H100; on-demand from $3.72/hr; dedicated bare-metal $2.14–$3.59/hr; range $3.50–$10.60/hr [secondary]
  - ⚠️ **H200 contract pricing rose ~40% between Oct 2025 and Mar 2026** even as B200 supply ramped — roughly half of tracked specialist providers reported no Hopper-class capacity coming off contract; demand outstripping supply across both Hopper and early Blackwell simultaneously [secondary — https://tech-insider.org/nvidia-b200-residual-value-158-percent-2026/]
- HBM memory scarcity (not GPU die availability) is now the primary price driver; H100/H200 contract pricing up ~40% in a six-month window while secondary-market older-gen pricing collapsed ~85% [secondary — https://intuitionlabs.ai/pdfs/data-center-gpu-pricing-2026.pdf]

### 3.3 Availability
- H100/A100/H200 cluster at 63–70% confirmed-stock rates across providers (rest provisioning-dependent) [secondary — AIMultiple]
- SemiAnalysis: "hunting for even 8 nodes (64 GPUs) of H100s or H200s is not easy; half the providers we asked were completely sold out" (early 2026); all capacity coming online until Aug–Sep 2026 already booked [independent — SemiAnalysis newsletter]
- H100 remains the cost-performance sweet spot and the default for training under 70B params in 2026; H200 is a drop-in rack replacement (same 700W TDP) tripling KV-cache headroom (141 GB HBM3e at 4.8 TB/s) [secondary — https://www.emma.ms/blog/nvidia-h200-vs-h100-comparison]

---

## 4. DGX systems: Spark, Station, Cloud

