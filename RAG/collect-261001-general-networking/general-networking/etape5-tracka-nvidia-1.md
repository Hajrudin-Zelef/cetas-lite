---
id: collect-261001-general-networking/general-networking/etape5-tracka-nvidia-1
title: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia", "Samsung", "TSMC"]
dates: ["2024-03", "2026-01", "2026-02-01", "2026-03", "2026-04", "2026-05", "2026-09-22"]
keywords: ["gpu", "gpus", "nvidia", "rubin", "3nm", "800v dc", "benchmarks", "blackwell", "compute", "cost", "fp4", "fp8"]
source: docs/RAG/collect-261001-general-networking/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [1, 100]
sha256: 7a937c343ab3aad51362db20b6cb66060d550a8a4cd819530fb17f7e20c67a91
---

# Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking
## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 5 (Servers & hardware), Track A
**Coverage window:** February 1, 2026 → September 22, 2026
**Status:** Research snapshot. Prices are dated; re-verify before use. Street prices and cloud rates move weekly.

### Provenance legend
- **[official]** — NVIDIA's own blog, docs, GTC keynote, datasheet, or regulatory filing.
- **[vendor-reported]** — figure claimed by NVIDIA (benchmarks, throughput, ROI) without independent audit.
- **[independent]** — reputable third-party press (Bloomberg, Reuters, CNBC, DCD, EE Times, SemiAnalysis, The Register, Tom's Hardware, ServeTheHome, VideoCardz) or independent measurement (LMSYS, Artificial Analysis).
- **[secondary]** — lower-tier press, blogs, analyst summaries, GitHub research memos; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

---

## 1. Blackwell: B200 / B300 GPU specifications

### 1.1 B200 (Blackwell)
- **Architecture:** Blackwell, TSMC 4NP process, dual-die package (~208B transistors) [secondary]
- **Streaming multiprocessors:** 148 SM (TechPowerUp lists B200 SXM 192 GB as 148 SM / 18,944 CUDA cores) [secondary]
- **Tensor cores:** 5th generation [secondary]
- **Compute:**
  - FP32: ~75 TFLOPS
  - FP16/BF16 tensor: ~2,250 TFLOPS (dense) [secondary — community spec tables]
  - FP8 tensor: ~4,500 TFLOPS dense (~9 PFLOPS sparse) [secondary]
  - FP4 (NVFP4): ~9 PFLOPS dense, ~18–20 PFLOPS sparse [secondary — figures vary across sources: 9/18 (sparse 2x) vs 9/20 per Spheron; treat as vendor-derived]
  - FP64: 37 TFLOPS [secondary]
- **Memory:** 192 GB HBM3e (one source table lists 180 GB "usable"); 8 TB/s bandwidth [secondary]
- **Interconnect:** NVLink 5, 1.8 TB/s per GPU; PCIe 5.0 [secondary]
- **TDP:** ~1,000 W class (configurable to 1,200 W per some tables) [secondary]
- **Launch status:** Announced GTC March 2024; volume shipping Q4 2024; widely deployed through 2025–2026 [secondary]
- **Pricing (2026 snapshot):**
  - Single GPU buy: ~$30,000–$50,000 [secondary — https://intuitionlabs.ai/articles/nvidia-ai-gpu-pricing-guide]
  - Cloud rental: $2.12–$6.04/hr (neo-cloud), on-demand median ~$5.50/hr (May 2026, Spheron); reserved 36-month pricing dropped to $2.25/GPU/hr (April 2026); spot from $2.12/hr [secondary — https://gpuaas.com/blog/h100-vs-h200-vs-b200-which-gpu-to-rent-2026]
  - **Residual value hit 158% in mid-2026 — 58% above launch price** [secondary — https://tech-insider.org/nvidia-b200-residual-value-158-percent-2026/]
  - B200 rental index surged 24% in March 2026 before pulling back (Silicon Data) [secondary]

### 1.2 B300 / Blackwell Ultra
- **Architecture:** Blackwell Ultra, TSMC 4NP, 208B-transistor dual-die [secondary]
- **Streaming multiprocessors:** 160 SM [secondary — per NVIDIA Technical Blog]
- **Compute:**
  - FP4 (NVFP4): ~15 PFLOPS dense (~28–30 PFLOPS sparse est.) [secondary]
  - FP8: ~4.5–13.5 PFLOPS depending on source (figures conflict — flagged)
  - FP16/BF16: ~2.2 PFLOPS [secondary]
  - ⚠️ **FP64 collapses to ~1.25 TFLOPS** (vs 37 TFLOPS on B200) — B300 trades nearly all double-precision performance for FP4 headroom [secondary — https://ifactoryapp.com/sap-integration/on-prem-ai/nvidia-dgx-b200-b300-blackwell-systems]
- **Memory:** 288 GB HBM3e; 8 TB/s bandwidth [secondary]
- **Interconnect:** NVLink 5, 1.8 TB/s per GPU; PCIe 6.0 [secondary]
- **TDP:** 1,200–1,400 W class (DGX B300 system draws 14.5 kW across 8 GPUs) [secondary]
- **Launch status:** Partners shipping from 2H 2025; DGX B300 shipping January 2026 [secondary]
- **Pricing (2026 snapshot):**
  - Single GPU: ~$53,000 [secondary]
  - DGX B300 (8-GPU system): $300,000–$350,000 [secondary]
  - Cloud rental: $3.27–$18.00/hr across 17 tracked providers [secondary — https://gpusmith.com/articles/en/h200-vs-b200-vs-b300-vs-mi325x]

### 1.3 DGX B200 vs DGX B300 systems
| Specification | DGX B200 | DGX B300 |
|---|---|---|
| GPUs | 8× B200 | 8× B300 |
| HBM3e per GPU | 192 GB | 288 GB (+50%) |
| Total system memory | 1.5 TB | 2.3 TB |
| FP4 dense per GPU | 9 PFLOPS | 15 PFLOPS (+67%) |
| FP64 per GPU | 37 TFLOPS | 1.25 TFLOPS |
| TDP per GPU | 1,000 W | 1,400 W |
| System peak power | ~10 kW | ~14 kW |
| NVLink | NVLink 5 | NVLink 5 |
| Cooling | Air or liquid (recommended) | Direct liquid cooling mandatory |
| System price (2026) | ~$280K–$320K | $300K–$350K |
| Lead time (Apr 2026) | 8–16 weeks | 12–20 weeks |
[secondary — https://ifactoryapp.com/sap-integration/on-prem-ai/nvidia-dgx-b200-b300-blackwell-systems]

### 1.4 GB200 / GB300 NVL72 rack-scale systems
- **GB300 NVL72:** 72 Blackwell Ultra GPUs + 36 Grace CPUs per rack; ~1.4 EFLOPS FP4 inference; ~140 kW rack power [secondary]
- Allocation at the top of the Blackwell stack remained constrained through 2026 [secondary]
- Rubin CPX marketing positions Vera Rubin NVL144 CPX at 7.5× the AI performance of GB300 NVL72 [vendor-reported]

---

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

