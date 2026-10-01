---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-4
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "CoreWeave", "Intel", "Nvidia", "xAI"]
dates: ["2024-05", "2025-07", "2025-11", "2026-05-28", "2026-05-29", "2026-06"]
keywords: ["agentic", "amd", "backlog", "blackwell", "dram", "ethernet", "fp4", "fp8", "gpu", "gpus", "inference", "intel"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [261, 366]
sha256: 68c772160dd570b06420cc93459f19d91728d0ca22194bd9e27dbd40e654734a
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

- **PowerEdge XE9680** — Dell's "fastest ramping solution ever" (predecessor to the
  2026 XE generation); 8-GPU platform supporting NVIDIA H100/H200/B200-class GPUs
  **[official]**.
- **PowerEdge XE9780 and XE9785** (air-cooled) + **XE9780L and XE9785L** (liquid-cooled),
  announced as the next-generation XE line with NVIDIA: successors to the XE9680,
  supporting **8-way NVIDIA HGX B300** and delivering **up to 4× faster LLM training**
  vs XE9680; configurations support **up to 192 NVIDIA Blackwell Ultra GPUs** with
  direct-to-chip liquid cooling, customizable to **up to 256 Blackwell Ultra GPUs per
  Dell IR7000 rack** **[official (Dell announcement via TechPowerUp techpowerup.com/336994;
  smetechguru.co.za)]**. The announcement (Dell Technologies World 2025 cycle) also
  introduced: **PowerEdge XE9712 with NVIDIA GB300 NVL72** (rack-scale; **50× more AI
  reasoning inference output, 5× throughput improvement**; new Dell PowerCool
  technology) **[official]**; **PowerEdge XE7745 with NVIDIA RTX PRO 6000 Blackwell
  Server Edition** (available July 2025, 4U, up to 8 GPUs, for physical/agentic AI —
  robotics, digital twins, multimodal) **[official]**; planned support for the **NVIDIA
  Vera CPU** and a new Dell PowerEdge XE server for **Vera Rubin** (Dell Integrated
  Rack Scalable Systems) **[official]**; networking: Dell PowerSwitch SN5600/SN2201
  (Spectrum-X Ethernet) and NVIDIA Quantum-X800 InfiniBand (800 Gb/s)
  **[official]**.
- **PowerEdge XE9785 / XE9785L with AMD Instinct MI355X** (announced at **SC25, St.
  Louis — November 2025**, per HotHardware): dual AMD EPYC CPUs + **8× AMD Instinct
  MI355X per node**; liquid-cooled variant is **3U per node** → **128 MI355X GPUs per
  standard 42U rack**; each MI355X has **288 GB HBM3e**, FP6/FP4 support, ~5 FP8
  PFLOPS per GPU (~40 PFLOPS per air-cooled 10U node); AMD Pensando Pollara 400 NICs;
  Dell PowerSwitch AI fabric (new **PowerSwitch Z9964F-ON / Z9964FL-ON**, liquid-cooled
  networking variant); also **PowerEdge R770AP** with Intel Xeon 6 6900-series
  (up to 128 P-cores per socket, dual socket, 12-channel memory), configurable with
  Instinct accelerators **[independent (HotHardware, Nov 2025)]**. ⚠️ Date note: the
  HotHardware article is from SC25 (Nov 2025), not 2026 — it is included because it
  defines Dell's 2026 AMD platform lineup.
- **PowerEdge XE9680L** (liquid-cooled 4U): supports **8× NVIDIA Blackwell GPUs**;
  "highest possible rack-scale density for NVIDIA GPUs in an industry standard x86
  rack," **33% more GPU density per node**, 20% more PCIe Gen5 slots, double
  North/South network expansion capacity; direct liquid cooling; factory-integrated
  rack-scale deployment **[official (Dell AI Factory update via
  digitalisationworld.com, updated ~June 2026)]**. ⚠️ The XE9680L was first announced
  at Dell Tech World 2024 (May 2024); the June-2026 item may be a re-announcement or
  availability update — flagged **[unverified]** which 2026 event it belongs to.
- **Turnkey rack-scale variants** (announced alongside XE9680L): air-cooled design
  supporting **64 GPUs in a single rack**, or liquid-cooled format with **72 NVIDIA
  Blackwell GPUs in a single rack** **[official]**.
- Dell AI Data Platform / Dell NativeEdge / Generative AI Solutions for Digital
  Assistants updates were announced alongside (non-server software) **[official]**.

### 2.2 xAI / Colossus-related deals

- **Feb 14–15, 2025 — Bloomberg/Reuters report** **[independent]**: Dell was "nearing a
  deal **worth more than $5 billion**" to provide xAI with AI-optimized servers
  containing **NVIDIA GB200** semiconductors, for delivery "this year" (2025); details
  still being discussed and subject to change; Dell and xAI did not comment
  (reuters.com via wdsm710.com; d2649.cms.socastsrm.com).
- **Margins**: a Bloomberg follow-up reported the $5B xAI deal would yield
  **mid-single-digit gross margins**, with **NVIDIA providing the networking** (a
  category Dell also sells); interviewed workers at the server makers described
  Nvidia's demands as overbearing **[independent (Bloomberg via visive.ai)]**.
- **Prior Colossus build**: the Memphis **Colossus** supercomputer was built using both
  **Dell and Supermicro** servers, ~100,000 NVIDIA GPUs; Musk planned expansion to
  200,000 and then stated a 1M-GPU target **[secondary (SDxCentral; en.tmtpost.com)]**.
  Dell COO Jeff Clarke told Bloomberg (Dec 2025) the company had **shipped tens of
  thousands of GPUs to xAI's supercomputer** **[secondary (en.tmtpost.com)]**.
- **Whether the $5B deal was finalized and how much has shipped in 2026: NOT publicly
  confirmed** in collected sources — flagged **[unverified]**. No Dell press release or
  2026 earnings statement names xAI as a completed $5B contract. Parent should verify
  before asserting the deal closed.

### 2.3 AI server revenue / backlog figures from earnings (ISG AI orders)

All figures **[vendor-reported]** (Dell official Q-results press releases):

- **Q4 FY2026** (reported Feb 26, 2026, quarter ended Jan 30, 2026): revenue **$33.4B,
  +39% YoY** (record); non-GAAP diluted EPS **$3.89, +45%**; **AI orders $34.1B** (record);
  **AI server revenue $9.0–9.5B, +342% YoY** (record; sources cite $9.0B and $9.5B —
  minor discrepancy between the official release and secondary reporting, flagged);
  **AI backlog $43.0B exiting Q4** (record); ISG revenue $19.6B +73%, operating income
  $2.9B +41%. **Full FY2026**: revenue **$113.5B, +19%** (record); non-GAAP EPS
  **$10.30, +27%**; **AI orders >$64B ($64.1B)**; **AI shipments $25.2B**; ISG
  **$60.8B, +40%**, operating income **$7.1B, +27%**; cash flow from operations
  **$11.2B**; capital returned to shareholders $7.5B; dividend raised 20%, buyback
  authorization +$10B. Initial FY27 guidance: full-year revenue **$140B** midpoint,
  **AI-optimized server revenue $50B** (>100% growth).
- **Q1 FY2027** (reported May 28, 2026): revenue **$43.8B, +88% YoY** (record); diluted
  EPS $5.24 +282%, non-GAAP EPS **$4.86, +214%**; **AI orders $24.4B**; **AI server
  revenue $16.1B, +757% YoY** (record); **AI backlog $51.3B** (record, per Zacks);
  ISG **$29.0B, +181%**; operating income $3.1B +206%. Raised FY27 revenue guidance to
  **$167B midpoint** (~+50% YoY), **AI server revenue to $60B**. Clarke: AI customer
  count topped **5,000** (+50% in six months) **[secondary (Zacks, May 29, 2026)]**;
  cited DRAM/NAND/CPU supply bottlenecks, 3–5-year customer supply deals
  **[secondary]**.
- **Q2 FY2027** (reported Sep 1, 2026, quarter ended Jul 31, 2026): revenue
  **$47.0B, +58% YoY** (record); GAAP EPS $6.34 +273%, non-GAAP EPS **$7.04, +203%**;
  **AI orders $60.9B** (record, "the most in our history" — Clarke); **AI server
  revenue $16.4B, +100% YoY** (record); **AI backlog $95B** (record, ~8× YoY per
  techx.pk); **H1 FY2027 AI revenue $32.5B** (vs $10.1B prior-year period — more than
  tripled); **trailing-12-month AI orders $131.7B** **[secondary]**; AI customers
  **>6,500** (+3,300 in the last three quarters); ISG **$31.8B, +89%**, operating
  income **$4.8B, +225%**; Traditional Servers & Networking **$10.5B, +122%**; Storage
  **$4.9B, +26%**; CSG **$15.0B, +20%**. Raised FY27 revenue guidance to **$192B**,
  non-GAAP EPS to **$25.50**, **AI server revenue to $74B** **[official (Dell press
  release via morningstar.com, bizwire-20260901)]**. Supporting stock news: $5B
  investment-grade bond offering drew ~$23B of orders; Dell added to the **S&P 100**
  **[secondary (tickeron.com)]**.
- Analyst price-target wave post-Q2: $635 JPMorgan, $600 Citi/Mizuho, $640 RBC
  Outperform initiation **[secondary]** — market commentary, not vendor guidance.

### 2.4 Partnerships (NVIDIA, CoreWeave, etc.)

