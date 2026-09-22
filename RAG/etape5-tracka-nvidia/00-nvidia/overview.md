---
id: etape5-tracka-nvidia/00-nvidia/overview
title: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
domain: step-5-track-a-nvidia-gpus-b200-b300-rubin-dgx-systems-netwo
role: deep-dive
task: hardware
actors: ["Nvidia", "TSMC"]
dates: ["2024-03", "2026-01", "2026-02-01", "2026-03", "2026-04", "2026-05", "2026-09-22"]
keywords: ["gpu", "gpus", "nvidia", "rubin", "benchmarks", "blackwell", "compute", "fp4", "fp8", "inference", "liquid cooling", "memory"]
source: docs/RAG/etape5_trackA_nvidia.md
source_anchor: ""
source_lines: [1, 78]
section: "Step 5 — Track A: NVIDIA GPUs (B200/B300/Rubin), DGX Systems & Networking"
sha256: c8ebd882c4e11c3d661996015eb1383c470e37446158443d04a87b3aea5b8be1
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

