---
id: etape5-trackc-huawei-intel/00-huawei-intel/3-intel
title: "3. INTEL"
domain: step-5-track-c-huawei-chinese-hardware-intel-cpus-feb-1-sep-
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Apple", "DeepSeek", "Falcon", "Google", "Huawei", "Intel", "Microsoft", "Nvidia", "Qualcomm", "Samsung", "TSMC", "United States"]
dates: ["2025-03", "2025-08", "2026-01-05", "2026-05", "2026-06-01", "2026-06-26", "2026-09-22"]
keywords: ["intel", "18a", "3nm", "accelerator", "agentic", "agents", "agi", "ai pc", "amd", "ascend", "clearwater forest", "consumer"]
source: docs/RAG/etape5_trackC_huawei_intel.md
source_anchor: ""
source_lines: [104, 214]
section: "Step 5 — Track C: Huawei + Chinese Hardware + Intel CPUs (Feb 1 → Sep 22, 2026)"
sha256: 44def0f1286ee2754fd114dcc19943a4dccc6827642e9893e572d799ea8def4f
---

# 3. INTEL

## 3. INTEL

### 3.1 Corporate / foundry 2026

- CEO **Lip-Bu Tan** (ex-Cadence chief) took the helm **March 2025**; aggressively restructured with Intel Foundry separation and 18A focus **[secondary — financialcontent]**.
- **US government stake:** $8.9B in CHIPS Act grants converted to a **9.9% equity stake (August 2025)** — the administration as stakeholder, not just grantor **[secondary — theeconomicrule citing Manufacturing Dive; financialcontent]**.
- **Intel 18A:** first leading-edge node with RibbonFET (GAA) + PowerVia (backside power); **Microsoft committed** as a foundry customer; **NVIDIA tested 18A via multi-project wafer runs** without committing volume (TrendForce) — one secondary source (financialcontent, Jan 2026) claims NVIDIA signed for 18A/14A, which is **conflicting/single-sourced — [unverified]**.
- **Stock:** ~**$128.76 on June 26, 2026** after a ~500% run under Tan (Fortune); analyst consensus Hold; bull targets ~$150 **[secondary — theeconomicrule]**.
- **Trump–Tan meeting Jan 8, 2026:** "national champion" framing; Advanced Manufacturing Investment Credit expected to rise **25% → 35%** for Ohio/Arizona mega-fabs **[secondary — financialcontent]** — treat policy specifics as **[secondary]**.

### 3.2 Xeon 6 family (Granite Rapids / Sierra Forest)

- **Granite Rapids (Xeon 6900P):** launched **Sep 24, 2024**; up to **128 Redwood Cove P-cores**; 504 MB L3; up to **500 W TDP**; 12 channels DDR5-6400; **MRDIMM up to 8800 MT/s** (market debut of the standard); 136 PCIe 5.0 lanes; CXL 2.0; AMX with MXFP4 support **[secondary — Wikipedia/techpowerup/dlcompare]**.
- **Sierra Forest (Xeon 6700E):** launched **Jun 4, 2024**; up to **288 Crestmont E-cores**; Intel 3; 8/12 DDR5 channels; 88 PCIe 5.0; CXL 2.0 **[secondary — Wikipedia]**.
- In-window 2026 news for the base Xeon 6 lineup is thin — the action moved to the successor (Clearwater Forest).

### 3.3 Clearwater Forest — Xeon 6+, launched June 1, 2026

- **Launched June 1, 2026 at Computex**; **first data-center processor on Intel 18A** **[secondary — Wikipedia]**.
- Up to **288 Darkmont E-cores** per socket; 2-socket configs; **12 channels DDR5-8000**; **96 PCIe 5.0 lanes**; **64 CXL 2.0 lanes**; up to 6 UPI links; 576 MB max L3; built-in accelerators (QAT, DSA, DLB, IAA); SGX/TDX **[secondary — Wikipedia]**.
- Succeeds Sierra Forest; shares I/O tiles with Granite Rapids; successor line points to **Diamond Rapids (2027)** — no 2026 Diamond Rapids/Panther Cove news located **[gap]**.

### 3.4 Gaudi 3 / Falcon Shores / Jaguar Shores — AI accelerator strategy

- **Gaudi 3** (launched fall 2024, Habana Labs): 64 TPCs + 8 MMEs; **128 GB HBM2e**; 24×200Gb Ethernet scale-out; SynapseAI stack. **Commercial miss:** Intel admitted (Nov 2024) it would miss its **$500M Gaudi 3 sales target** on software issues; beyond IBM, few major providers committed **[independent — TechCrunch]**.
- **Falcon Shores CANCELLED as a product (Jan 30, 2025):** interim co-CEO Michelle Johnston Holthaus announced Falcon Shores — the TSMC-3nm, ~1500W programmable GPU meant to succeed Gaudi 3 in late 2025 — would be an **internal test chip only**, not brought to market **[independent — CRN; TechCrunch; TweakTown]**.
- **Jaguar Shores:** the successor is now framed as a **rack-scale system-level solution**, not a standalone chip — Intel's lesson from Gaudi being "it's not enough to just deliver the silicon" **[independent — CRN]**.
- **2026 status:** no verified Jaguar Shores silicon/launch news in-window **[gap]**.
- **Crescent Island:** next-gen **inference-focused datacenter GPU (Xe3P)**, reported with 160 GB LPDDR5X, air-cooled, **sampling H2 2026** — sourced only from a Red Hat ecosystem GitHub repo; **treat as [unverified]**.

### 3.5 Panther Lake — Core Ultra Series 3 (client AI PC, Jan 2026)

- **Launched January 5, 2026 at CES**; **first consumer chip on Intel 18A**, designed and manufactured entirely in the US; RibbonFET + PowerVia **[independent — KitGuru; Micro Center]**.
- Up to **16 cores** (4 Cougar Cove P + 8 Darkmont E + 4 LP E-cores); **Xe3 "Celestial" iGPU** up to 12 Xe cores (Arc B390), matching a laptop RTX 4050 per Intel; XeSS 3 with multi-frame generation **[independent — KitGuru]**.
- **NPU 5: 50 TOPS** (Copilot+ compliant); **180 platform TOPS** total (CPU+GPU+NPU); 4.3× faster LLM inference than AMD XDNA2 (Intel claim — **[vendor-reported]**) **[independent — KitGuru; secondary — GitHub research note]**.
- Memory: **LPDDR5X-9600**, DDR5-7200; 25W base / 80W turbo TDP; Thunderbolt 4, Wi-Fi 7, PCIe 5.0 (12/20 lanes) **[secondary]**.
- Intel claims +50–60% multithreaded perf gen-on-gen at 25W; "battery life king" Core Ultra X9 388H with up to 27h Netflix streaming **[vendor-reported]**.
- **200+ system designs** planned (ASUS Zenbook, MSI Prestige/Titan/Stealth, Samsung Galaxy Book6, Lenovo Yoga Slim 7i) **[secondary]**.
- CES demo: **70B-parameter LLM running locally** with 32K context on a thin-and-light reference design via unified memory (up to 128 GB allocatable to AI tasks) **[vendor-reported — financialcontent]**.
- Intel positioning: "Agentic AI PCs" running autonomous agents on-device **[secondary]**.

---

## 4. CPU LANDSCAPE 2026 — EPYC vs XEON, MEMORY

### 4.1 Market share (Mercury Research)

- **Q1 2026:** AMD hit **32.6% of the x86 CPU market** (units, all segments) — all-time high, up from 27.1% a year earlier; Intel 67.4% (down from 72.9%) **[secondary — TweakTown citing Mercury Research]**. AMD **server unit share 33.2%** (up from 27.2%); **EPYC 46.2% of server CPU revenue** — record, up from 41.3% in Q4 2025 — while Intel held 53.8% of x86 server revenue and 54.9% of server units; **Arm servers at 17.7%** **[secondary — TweakTown; theoutpost; walaw/UBS]**.
- **Q2 2026:** AMD server unit share **34.5%** (from 27.3% a year earlier); in a head-to-head **EPYC vs Xeon SP** view, AMD approaches **46.4%** **[secondary — webpronews citing Mercury]**.
- ⚠️ Segment-definition conflicts exist between outlets (e.g., Q2 "overall x86 share 30.7%" vs Q1's 32.6%) — figures above are quoted as reported per outlet; treat cross-outlet comparisons cautiously **[secondary]**.
- **Economics:** server CPU shipments +10–20% YoY driven "almost entirely by AI data center demand"; Intel server shipments flat; **AMD Q1 data-center revenue $5.8B of $10.3B total (55.9%), +57% YoY** — surpassing Intel's DCAI segment ($5.1B) in the same quarter, a structural inflection **[secondary — TweakTown; walaw]**.
- Lisa Su's framing: moving toward a **1:1 CPU:GPU ratio** in AI infrastructure as agentic workloads grow **[secondary — TweakTown]**.
- **Arm:** server shipments nearly doubled; 17.7% share with growth expected after **Arm AGI** launch — details thin **[secondary]**.
- **EPYC generation:** no new EPYC generation launch (9006/Venice) verified in-window — **gap**; the installed story remains EPYC 9005 (Turin).

### 4.2 Memory: DDR5, MRDIMM, CXL, and the DRAM crunch

- **MRDIMM** debuted with Granite Rapids — up to **8800 MT/s** on Xeon 6900P, targeting memory-bandwidth-bound AI workloads **[secondary — dlcompare]**.
- **Clearwater Forest** pushes to **DDR5-8000** across 12 channels **[secondary — Wikipedia]**.
- **CXL 2.0** is standard on Xeon 6 and Clearwater Forest (64 lanes on the latter) for memory expansion/pooling; no verified CXL 3.0 deployment news in-window — **gap**.
- **The 2026 DRAM crunch:** chipmakers funneled production toward **HBM for AI accelerators**, leaving consumer-grade DRAM scarce and expensive; Mercury Research: "higher PC prices and limited GPU supplies are having a significant impact on end demand for desktop PCs" **[secondary — webpronews citing Mercury/The Register]**.
- Server DRAM pricing specifics for 2026 were not verified in this pass — **gap**; re-verify against DRAMeXchange/TrendForce contract-price series.

### 4.3 Other 2026 CPU/server-chip launches relevant to AI

- **Qualcomm Snapdragon X2 Plus** (client): up to 80 platform TOPS — cited as Panther Lake's competitor at CES 2026 **[secondary — Micro Center]**.
- **Apple M5-class silicon** referenced in local-inference contexts (Step 4 track B); no server CPU play.
- No verified **NVIDIA Grace CPU** refresh, **Amazon Graviton 5**, or **Google Axion** next-gen launch news was collected in this pass — **gaps**; worth a dedicated sweep if server-CPU completeness is required.

---

## 5. MASTER TIMELINE (Feb 1 → Sep 22, 2026)

- **Jan 2, 2026** — Biren lists on HKEX (6082); +119% intraday, closes +76%; raises HK$5.58B **[secondary]**.
- **Jan 5, 2026** — Intel launches Panther Lake / Core Ultra Series 3 at CES (Intel 18A) **[independent]**.
- **Jan 8, 2026** — Intel Iluvatar CoreX lists on HKEX **[secondary]**.
- **Jan 8, 2026** — Trump–Tan "national champion" meeting; Intel +10.5% **[secondary]**.
- **Jan 22, 2026** — Enflame STAR IPO accepted (first A-share IPO of 2026) **[secondary]**.
- **Feb–Mar 2026** — Cambricon Siyuan 690 in testing; H2 2026 mass-production target **[independent]**.
- **Mar 2026** — Moore Threads wins RMB 660M KUAE 10K-GPU cluster order **[independent]**.
- **Apr 28, 2026** — Moore Threads reports Q1 2026: first quarterly profit (RMB 29.36M attributable) **[independent]**.
- **May 2026 (expected)** — Huawei Ascend 910D first samples (Bloomberg reporting) **[independent — via relay]**.
- **Jun 1, 2026** — Intel Clearwater Forest (Xeon 6+) launches at Computex (Intel 18A, 288 E-cores) **[secondary]**.
- **Jun 2026** — Moore Threads and MetaX begin Hong Kong "A+H" listing processes **[secondary]**.
- **Jun 26, 2026** — Intel stock ~$128.76 (Fortune) **[secondary]**.
- **H2 2026 (planned)** — Intel Crescent Island inference GPU sampling **[unverified]**.
- **Q2–Q3 2026 (rumored)** — Ascend 910D mass production **[unverified]**.

---

## 6. OPEN VERIFICATION ITEMS

1. Ascend 920 production/shipment status in 2026 — no confirmation found **[unverified]**.
2. Ascend 910D specs (5nm? 4-die? FP8?) and mass-production timeline — Bloomberg only confirms H100-targeting samples **[secondary]**.
3. Huawei's "double output in 2026" — Bloomberg private-source reporting **[secondary]**.
4. Cambricon's 500K-unit 2026 plan — company publicly called reports "inaccurate" **[secondary]**.
5. SMIC N+2 yield figures (20% Cambricon 590/690; 40% Huawei 910C) — single-source relays **[secondary]**.
6. Moore Threads' 1,000 TFLOPS/S5000, KUAE 2.0 efficiency (60%/40%), DeepSeek R1 inference numbers — vendor-reported **[vendor-reported]**.
7. MetaX Q1 financials (2.3B yuan loss, 8.9B inventory) — IPO-filing-era reporting **[secondary]**.
8. Intel Crescent Island existence/specs — single GitHub-ecosystem source **[unverified]**.
9. NVIDIA as Intel 18A/14A volume customer — conflicting/secondary **[unverified]**.
10. Clearwater Forest specs — Wikipedia infobox (secondary); cross-check with Intel ARK **[secondary]**.
11. Panther Lake 70B-local-LLM demo, 27h battery, 4.3× vs XDNA2 — Intel claims **[vendor-reported]**.
12. Mercury Research share figures — consistent across relays but primary report not fetched **[independent-via-secondary]**.
13. Kunpeng 930/935, EPYC 9006/Venice, CXL 3.0 deployments, server DRAM pricing — **gaps, not located**.

## 7. COLLECTION METADATA

- **Research date:** September 22, 2026.
- **Priority sources used:** ServeTheHome-class coverage via TechPowerUp/TweakTown/Tom's Hardware; SemiAnalysis (CloudMatrix); Mercury Research (share); TrendForce/DIGITIMES relays; TrendForce news; Bloomberg relays; official-vendor claims via press.
- **Conventions kept:** every fact carries a provenance tag; uncertain items numbered in §6; no identifiers guessed; prices are dated snapshots.
