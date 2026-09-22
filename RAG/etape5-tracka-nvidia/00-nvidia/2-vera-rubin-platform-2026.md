---
id: etape5-tracka-nvidia/00-nvidia/2-vera-rubin-platform-2026
title: "2. Vera Rubin platform (2026)"
domain: step-5-track-a-nvidia-gpus-b200-b300-rubin-dgx-systems-netwo
role: deep-dive
task: hardware
actors: ["AWS", "CoreWeave", "Google", "Groq", "Lambda", "Meta", "Microsoft", "Nebius", "Nscale", "Nvidia", "OpenAI", "Oracle", "Samsung", "TSMC"]
dates: ["2026-06-01"]
keywords: ["rubin", "vera rubin", "3nm", "800v dc", "agentic", "attention", "aws", "blackwell", "capex", "compute", "cost", "disaggregated"]
source: docs/RAG/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [79, 137]
section: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
sha256: bf9d0ddfbaff3b74c6f38794fbb18f3e325072525051705146c99f185d0a7f02
---

# 2. Vera Rubin platform (2026)

## 2. Vera Rubin platform (2026)

### 2.1 Platform overview
- **Announcement timeline:** CES 2026 (January, Las Vegas) — formal platform announcement; **GTC 2026 (March 16, keynote)** — full production confirmed; partner availability H2 2026 [independent — DCD; secondary]
- **Silicon:** 336-billion-transistor platform on TSMC N3P (3nm) — nearly double Blackwell's density [secondary — https://blockeden.xyz/blog/2026/03/20/nvidia-gtc-2026-vera-rubin-gpu-architecture-depin-compute-bottleneck/; figure from secondary sources, not verified against NVIDIA]
- **Memory:** HBM4 (SK Hynix and Samsung); per-GPU memory bandwidth ~22 TB/s — nearly 3× Blackwell's 8 TB/s [secondary]
- **Performance claims:** 5× inference throughput vs Blackwell at rack level; up to 10× lower inference token cost; 4× fewer GPUs for MoE training vs Blackwell; MoE inference at ~1/7 the token cost of GB200 [vendor-reported]
- **⚠️ Supply:** HBM4 reported completely sold out through 2026 at GTC 2026; GPU lead times 36–52 weeks [secondary]

### 2.2 Rubin GPU (R100)
- 72 Rubin GPUs in NVL72; rack delivers 3.6 EFLOPS NVFP4 inference + 2.5 EFLOPS training → **~50 PFLOPS NVFP4 per GPU** (derived) [secondary]
- NVLink 6: 3.6 TB/s per GPU scale-up (2× Blackwell); 260 TB/s aggregate NVL72 fabric [secondary]
- Vera Rubin NVL72 power: 190–230 kW per rack (vs ~140 kW for GB300 NVL72); requires **800V DC power delivery** rather than the 48V standard [secondary]
- Full liquid cooling (dry cooling), no fans; NVIDIA claims 47 minutes from truck arrival to power-on [secondary — https://finance.biggo.com/news/202607220220_Nvidia_Vera_Rubin_NVL72_full_production]

### 2.3 Vera CPU
- 88 NVIDIA custom "Olympus" cores, 176 threads (Spatial Multi-Threading) [secondary — https://videocardz.com/newz/nvidia-vera-rubin-nvl72-detailed-72-gpus-36-cpus-260-tb-s-scale-up-bandwidth]
- 2 MB L2 per core; 162 MB unified L3 [secondary]
- Memory: up to 1.5 TB LPDDR5X at up to 1.2 TB/s (vs Grace: 480 GB at 512 GB/s) [secondary]
- NVLink-C2C: 1.8 TB/s (vs 900 GB/s on Grace) [secondary]
- PCIe Gen6 / CXL 3.1; confidential compute supported [secondary]

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

