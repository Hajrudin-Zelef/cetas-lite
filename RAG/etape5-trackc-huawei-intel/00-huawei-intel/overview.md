---
id: etape5-trackc-huawei-intel/00-huawei-intel/overview
title: "Step 5 — Track C: Huawei + Chinese Hardware + Intel CPUs (Feb 1 → Sep 22, 2026)"
domain: step-5-track-c-huawei-chinese-hardware-intel-cpus-feb-1-sep-
role: deep-dive
task: hardware
actors: ["ByteDance", "China", "DeepSeek", "Huawei", "Intel", "Nvidia", "TSMC", "United States"]
dates: ["2025-04", "2025-11", "2026-02-01", "2026-04", "2026-09-22"]
keywords: ["intel", "accelerator", "ascend", "blackwell", "chiplet", "clearwater forest", "compute", "consumer", "cost", "datacenter", "deepseek", "export controls"]
source: docs/RAG/etape5_trackC_huawei_intel.md
source_anchor: ""
source_lines: [1, 56]
section: "Step 5 — Track C: Huawei + Chinese Hardware + Intel CPUs (Feb 1 → Sep 22, 2026)"
sha256: f207105b7074ea50c22921c38e7ff4423a91c89099ddc162677cd661fd1afbda
---

# Step 5 — Track C: Huawei + Chinese Hardware + Intel CPUs (Feb 1 → Sep 22, 2026)
## Research report (English) — collected September 22, 2026

**Project:** RAG data-collection, Step 5 (Serveurs & hardware), Track C
**Coverage window:** February 1, 2026 → September 22, 2026
**Topics:** Huawei Ascend (910C/910D/920), CloudMatrix 384, Chinese AI chips (Cambricon, Moore Threads, Biren, MetaX), Intel Xeon 6/Clearwater Forest, Gaudi 3/Jaguar Shores, Panther Lake, CPU market share, memory

### Provenance legend
- **[official]** — vendor's own blog, docs, filings, earnings statements.
- **[vendor-reported]** — figure claimed by the vendor without independent audit.
- **[independent]** — reputable third-party press or independent measurement (SemiAnalysis, Mercury Research, TrendForce, Bloomberg, Reuters, Tom's Hardware, ServeTheHome, EE Times, DIGITIMES, SiliconANGLE).
- **[secondary]** — lower-tier press, blogs, analyst summaries; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

---

## 1. HUAWEI ASCEND

### 1.1 Ascend 910C — current flagship in production

- **Architecture:** dual-chiplet accelerator — two Ascend 910B dies packaged together — fabricated on **SMIC's N+2 7nm-class process** (no EUV) **[independent — Tom's Hardware]**. DaVinci NPU architecture, stacked HBM2E memory (8 stacks per package per TechPowerUp; 128 GB total per the LinkedIn/TechInsights summary) **[independent/secondary]**.
- **Performance (vendor-reported via press):** up to **780 TFLOPS dense BF16** per chip; package power ~**350 W**; chip-to-chip interconnect bandwidth 400 GB/s **[independent — Tom's Hardware; secondary — TechPowerUp]**.
- **Real-world standing:** DeepSeek research put Ascend 910C at roughly **60% of NVIDIA H100 inference performance** — despite only ~25% of the FLOPs — because inference is memory-bandwidth-bound **[secondary — SDxCentral, citing DeepSeek research]**. Yields reported near **40%** on SMIC 7nm in early 2025, at which point the line was described as profitable — "a crisis at TSMC, a milestone in Shenzhen" **[secondary — LinkedIn analysis]**.
- **2026 shipments:** Bloomberg (Apr 2025, context) expected Huawei to ship **over 800,000 Ascend 910B and 910C chips in 2025**, including to state telecom carriers and ByteDance **[independent — TweakTown citing Bloomberg]**. For 2026, Bloomberg (Dec 2025) reports **Huawei is preparing to double output** of its most advanced AI chips **[independent — via i3investor/Bloomberg relay]**.
- **Export-control context:** NVIDIA's H20 was effectively banned from China in April 2025 (NVIDIA took a ~$5.5B charge); Jensen Huang said in November 2025 that NVIDIA is "effectively blocked" from China; the Trump administration was reportedly considering allowing H200 sales, with no guarantee Beijing would permit adoption **[independent — SDxCentral; secondary — Bloomberg relay]**.

### 1.2 Ascend 910D — H100-class successor

- **Bloomberg (April 2026 reporting, relayed by TweakTown):** Huawei's Ascend 910D targets performance **matching or exceeding NVIDIA's Hopper H100** (2022 generation); first batch of samples expected in **late May**, development "still at an early stage" **[independent — Bloomberg via TweakTown]**.
- An unreliable community source (Medium) claims 910D uses a 5nm-class process, 4-die packaging, FP8 support, with volume production in Q2–Q3 2026 — **treat as [unverified]**.
- No credible first-party 910D announcement or shipment confirmation found as of Sep 22, 2026 **[unverified]**.

### 1.3 Ascend 920 — next-generation (general-purpose direction)

- **Unveiled April 2025** at a Huawei partner conference, days after NVIDIA's H20 export-license news **[independent — SDxCentral, citing DigiTimes Asia]**.
- Reported specs (DigiTimes/TechPowerUp): **SMIC 6nm** process; training-focused **Ascend 920C** delivering **>900 TFLOPS BF16** per card (~30–40% over 910C); **HBM3, 4,000 GB/s** memory bandwidth (up from 910C's 3,200 GB/s); PCIe 5.0 support; refined tensor engines for Transformer/MoE; **30–40% training-efficiency improvement** projected over 910C **[secondary — TechPowerUp; techhounder]**.
- Mass production was expected in **H2 2025** per DigiTimes — **no verified 2026 production/shipment confirmation located**; treat 920 availability as **[unverified]**.
- Huawei is reportedly repositioning Ascend 920 as a **general-purpose GPU** (not AI-only), a strategic pivot toward competing with NVIDIA across broader compute workloads **[secondary — techhounder]**.
- ⚠️ Some outlets conflate "Ascend 920" with "Ascend 920C" naming; keep the distinction flagged.

### 1.4 CloudMatrix 384 — rack-scale answer to GB200 NVL72

- **What it is:** 384 Ascend 910C chips across **16 racks (12 compute + 4 networking)**, all-to-all topology, showcased at WAIC Shanghai and the April 2025 partner conference **[independent — SDxCentral]**.
- **SemiAnalysis analysis (Apr 2025):** ~**300 PFLOPS BF16** total (vs 180 for GB200 NVL72); **49.2 TB aggregate HBM** (vs 13.8 TB on NVL72, ~3.6×); **2.1× scale-up bandwidth** and **5.3× scale-out bandwidth** vs NVL72; **6,912 pluggable optical transceivers (LPO @ 800G)** forming an all-optical all-to-all fabric with >5.5 Pbps aggregate internal bandwidth **[independent — SemiAnalysis, via TweakTown/SDxCentral/winbuzzer]**.
- **The cost:** ~**559 kW** peak power draw — roughly **4×** the GB200 NVL72 — plus 5× the chip count and 16× the floor space. SemiAnalysis's one-line logic: "having five times as many Ascends more than offsets each one being a third the performance of a Blackwell" **[independent — SemiAnalysis]**.
- **Economics:** electricity in parts of China fell to ~**$56/MWh in early 2025** (from ~$91/MWh in 2022), making power-hungry systems viable **[secondary — winbuzzer]**.
- **Software:** Huawei's **CANN** programming environment and **MindSpore** framework, with a translation layer ingesting PyTorch/TensorFlow graphs; Huawei plans to open-source more of the toolchain **[independent — Tom's Hardware]**.
- SemiAnalysis also flagged supply-chain reality: while Ascend *can* be fabbed at SMIC, the chip remains "global" — HBM from Korea, wafer production links to TSMC, tens of billions in fab equipment from the US/Netherlands/Japan — describing it as "aggressive skirting of export controls" **[independent — SemiAnalysis]**.

### 1.5 Huawei 2026 announcements (other)

- No verified new Ascend consumer/datacenter product launch at Huawei's 2026 flagship events was located in this research pass beyond the 920/910D roadmap items above — flag as a **coverage gap**; the Ascend story in-window is dominated by 910C volume ramp and CloudMatrix deployments.
- **Kunpeng server CPUs:** no in-window Kunpeng 930/935 launch news verified in this pass — **gap**; re-verify against Huawei's official server pages.

---

