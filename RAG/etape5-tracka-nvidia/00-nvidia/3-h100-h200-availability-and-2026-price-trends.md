---
id: etape5-tracka-nvidia/00-nvidia/3-h100-h200-availability-and-2026-price-trends
title: "3. H100 / H200: availability and 2026 price trends"
domain: step-5-track-a-nvidia-gpus-b200-b300-rubin-dgx-systems-netwo
role: deep-dive
task: hardware
actors: ["AWS", "Nvidia"]
dates: []
keywords: ["blackwell", "cost", "decode", "fp4", "gpu", "gpus", "hbm", "inference", "llama", "lpddr5x", "memory", "nvidia"]
source: docs/RAG/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [138, 183]
section: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
sha256: 1ae2b6ac5ce9c74776d8d8395d955874a02c769c30bdb6e91d30b21c5d9b58ad
---

# 3. H100 / H200: availability and 2026 price trends

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

### 4.1 DGX Spark (GB10)
- **Specs:** NVIDIA GB10 Grace Blackwell Superchip (20-core Arm: 10× Cortex-X925 + 10× Cortex-A725); Blackwell GPU with 6,144 CUDA cores, 5th-gen Tensor, 4th-gen RT; up to **1 PFLOP (sparse FP4)** AI performance [secondary — https://peterfalkingham.com/2026/09/18/academic-tech-gigabyte-ai-top-nvidia-dgx-spark-review/]
- **Memory:** 128 GB LPDDR5x unified, 273 GB/s; model capacity ~200B params FP4 inference (405B with two units over ConnectX-7) [secondary]
- **Storage:** 4 TB NVMe Gen5, self-encrypting; **Networking:** 10 GbE RJ-45, ConnectX-7 @ 200 Gbps, WiFi 7; **Power:** 240W PSU (GB10 TDP 140W; 40–60W observed at wall under load) [secondary]
- **OS:** NVIDIA DGX OS (Ubuntu 24.04 base) [secondary]
- **Pricing:**
  - Launch MSRP: $3,999 → raised to **$4,699** (Feb 23, 2026) [secondary]
  - Sep 2026 street: NVIDIA marketplace $4,699; Amazon $4,679; retailer markups $4,999.99 (Micro Center via Amazon) to $6,155.99 (Best Buy/Newegg); UK £4,900–£5,485; MSI EdgeXpert variant spiked to $7,342 [secondary — https://www.eteknix.com/rtx-spark-prices-are-already-surging-weeks-before-launch/]
  - Price history: lowest $3,979.19 (Jan 30, 2026), highest $7,346.22 (Sep 15, 2026), average $4,757.26 [secondary — pricehistory.app]
- **Performance (independent):** LMSYS Org — GPT-OSS 20B in Ollama: 2,053 tok/s prefill, 49.7 tok/s decode (~1/4 of RTX Pro 6000 Blackwell WE; slower than single RTX 5090); Llama 3.1 8B @ batch 32: 368 tok/s decode; Sebastian Raschka: single-sample PyTorch inference "roughly on par with the 6× more expensive H100" [independent — https://gpusmith.com/articles/en/pdfs/nvidia-dgx-spark-review-specs-performance.pdf]
- **Issues:** thermal and power-delivery problems on early units — John Carmack publicly reported units capping near 100W of rated 240W and rebooting under sustained load; NVIDIA dev forum moderators confirmed "a known, documented problem across the platform" [independent/secondary]
- **Reception:** Tom's Hardware — "a well-rounded toolkit for local AI... but a pricey platform if you don't intend to use its features to the fullest" [independent]
- **Alternatives (Sep 2026):** Strix Halo 128GB ($2,600–3,600), OEM GB10 boxes (often below FE price), Radeon PRO W7900 ($3,700–3,999), Mac Studio M5 Max (from $2,499), RTX 5090 ($3,700–5,000 street) [secondary — https://memeburn.com/dgx-spark-alternatives/]
- **RTX Spark:** per eTeknix (Sep 21, 2026), "RTX Spark Mini PCs make their market debut" weeks out — prices already surging [secondary — treat launch timing as [unverified]]

### 4.2 DGX Station / DGX Cloud
- No verified 2026 DGX Station (desktop/workstation) refresh details were located in the fetched sources — **gap flagged** [unverified]
- DGX Cloud: NVIDIA's AI supercomputing cloud offering; no standalone 2026 pricing/availability news located in this research pass — **gap flagged** [unverified]
- Context: DGX Cloud was the vehicle for NVIDIA's "AI factory" enterprise push alongside Nemo/Dynamo software at GTC 2026 [secondary]

---

