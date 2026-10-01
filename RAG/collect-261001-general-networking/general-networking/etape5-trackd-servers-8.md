---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-8
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Intel", "Nvidia"]
dates: ["2026-03", "2026-06-16", "2026-07", "2026-11", "2027-03"]
keywords: ["agentic", "amd", "benchmarks", "blackwell", "deepseek", "disclosure", "energy", "ethernet", "gpu", "gpus", "hyperscaler", "inference"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [561, 618]
sha256: ba9eb96192c22fc000fa0b2c55a6d833c552f2735a4e61fa92c7218737d1c3ae
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

**AMD platforms — Advancing AI 2026 (July 2026)**
- New server platforms for **6th Gen AMD EPYC 9006 (SP7 and SP8 sockets)**; full portfolio spanning rack, blade, GPU, edge [secondary — https://www.channelpostmea.com/2026/07/29/giga-computing-adds-support-for-6th-gen-amd-epyc-and-debuts-first-9006-server/ and https://sauditechpost.com/2026/07/30/giga-computing-expands-its-server-lineup-with-6th-gen-amd-epyc/].
- GIGABYTE systems around **AMD Instinct MI350P and MI355X accelerators**, **AMD Pensando Pollara 400 AI NIC**, and **UfiSpace 800G switches** — end-to-end AMD AI infrastructure [secondary — sauditechpost].
- New models: **R165-DG2-CS1** (single-socket 1U, closed-loop liquid cooling, SP7 max 256-core config); **B685-D80-LS1** (6U ten-node blade, DLC, high-density HPC) [secondary — channelpostmea].
- Shipment schedule: **9006 SP7 systems — day-one shipments with AMD in November 2026**; **SP8 systems — March 2027** [secondary — channelpostmea; also https://finance.biggo.com/news/0d245922-df28-40e6-8c8e-0f6badfdbb25].
- Giga Computing GM **Daniel (Hou Chih-jen) Hou** quote: platforms re-engineered from board layout, power delivery, and thermal design for agentic-AI-era demands [vendor-reported via secondary outlets above].

**Vera Rubin** — Giga Computing unveiled an **NVIDIA Vera Rubin rack-scale AI platform at ISC 2026** (June); analysts expect the next-gen Rubin platform launch in **H2 2026** [secondary — https://finance.biggo.com/news/ea828e8b-8abc-436e-9e17-10ecc5fe586d].

### 2.2 Key customers
- **RIKEN R-CCS** (FugakuNEXT HPC-Quantum hybrid platform, Japan) — XN24 GB200 NVL4 [official — https://www.gigabyte.com/press/news/2360]. (The task asked for "key customers" — beyond RIKEN, no 2026 customer names were surfaced in this research; hyperscaler/ODM-style volumes flow through without public disclosure — flagged as a gap.)

### 2.3 Market position (2026)
- **Q1 2026**: GIGABYTE AI server revenue **+110% YoY**, representing **71.5% of total sales**; AI servers now **>80% of total server revenue, approaching 90%** [secondary — https://www.ainvest.com/news/gigabyte-record-revenue-9-margin-ai-server-math-praising-2609/].
- Analyst estimates: GIGABYTE (2376.TW) **Q2 2026 revenue ~NT$140.5B (~$4.4B), +34% QoQ / +37% YoY**, EPS NT$9.08; full-year 2026 revenue **NT$500.5B (~$15.8B), +49% YoY**, EPS NT$31.22; AI server business **67% of 2026 revenue, rising to 72% in 2027** (vs 67% in 2025), AI business growth **+86% (2026) / +37% (2027)** [secondary — biggo finance / analyst note summaries, https://finance.biggo.com/news/ea828e8b-8abc-436e-9e17-10ecc5fe586d].
- Margin context: GIGABYTE gross margin **~9.6% (Q2 2026)** — assembler economics vs NVIDIA's ~75%; Q3 margin recovery expected on motherboard/GPU peak season and GB-series vs HGX mix shift; management stated **H2 revenue and earnings will outpace H1** [secondary — ainvest].
- Market backdrop: IDC Q1 2026 — GPU-accelerated servers **$68.9B (+24.8% YoY), 56.2% of the $122.6B market**; branded OEMs gaining share as ODM Direct share compressed (64.1% → 50.2%) [independent — IDC, https://www.idc.com/resource-center/press-releases/1q26-server-tracker/]; Q2 2026 total $166.3B (+52% YoY) [independent — IDC via StorageReview].
- ⚠️ **Price gap**: no public pricing found for G893/XN24 systems (enterprise quotes only).

---

## 3. ASUS

### 3.1 2026 AI server offerings & launches

**XA NB3I-E12 — 8-GPU NVIDIA HGX B300 (Blackwell Ultra)**
- Shipped **Oct/Nov 2025** (announcement ~326 days before Sept 22, 2026): 8× NVIDIA Blackwell Ultra GPUs, dual **Intel Xeon 6** Scalable CPUs, **8× NVIDIA ConnectX-8 InfiniBand SuperNICs**, 5× PCIe expansion slots, 32 DIMMs, 10× NVMe drives, dual 10Gb LAN [official via PR Newswire/Technode — https://technode.global/prnasia/asus-launches-xa-nb3i-e12-ai-server-built-with-nvidia-hgx-b300/].
- **June 16, 2026 — MLPerf Training v6.0**: #1 rankings in Llama 3.1-8B (Server and Offline modes) and GPT-OSS-120B (Interactive); training scores: 6.584 (Llama 2 7B), 84.85 (DeepSeek 671B), 74.75 (Llama 3.1 8B), 2.342 (DLRM-DCN) [official — https://servers.asus.com/news/ASUS-NVIDIA-HGX-B300-AI-Server-Dominated-MLPerf-Training-v60-Benchmarks]. System "now shipping worldwide" [official — same].

**ASUS AI POD with NVIDIA GB300 NVL72 (Blackwell Ultra rack-scale)**
- 72× Blackwell Ultra GPUs + 36× Grace CPUs per rack, **up to 40 TB high-speed memory per rack**, NVIDIA Quantum-X800 InfiniBand + Spectrum-X Ethernet integration, SXM7 and SOCAMM modules, **100% liquid-cooled**, for trillion-parameter LLM training/inference [secondary — https://markets.hanidaily.com/info/t-2503200966.html (carrying ASUS PR), ~March 2026]. NVIDIA VP Kaustubh Sanghani (GPU products) endorsed ASUS/Blackwell Ultra collaboration [vendor-reported via same].

**ASUS AI POD with NVIDIA Vera Rubin NVL72 (Ai4 2026, Aug 2026)**
- 100% liquid-cooled rack, **densities up to 227 kW per rack**; **XA NR1I-E12L**: 8-GPU **NVIDIA HGX Rubin NVL8** server, hybrid-cooled, as the deskside/entry point to the Rubin architecture [secondary — https://www.businesswire.com/news/home/20260805966459/en/ASUS-Brings-Full-Stack-AI-Infrastructure-From-Data-Center-to-Deskside-to-Ai4-2026].
- Shown at **ASUS AI Tech 2026, Seoul** (Sept 2026): enterprise AI server platforms, HGX training systems, rack-scale AI POD on Vera Rubin NVL72 under the "AI factories" theme [official — https://press.asus.com/blog/asus-ai-tech-2026-seoul-ai-factory-infrastructure/].

**ESC NM2N721-E1 — GB200 NVL72 (customer deployment)**
- NVIDIA GB200 NVL72 architecture (36 Grace CPUs + 72 Blackwell GPUs per rack); deployed in Taiwan's **NCHC "Nano4" AI supercomputer** — Taiwan's first fully liquid-cooled GB200 NVL72 system, in operation [secondary — https://www.techpowerup.com/343414/asus-hardware-powers-taiwans-nchc-ai-supercomputer-ranked-29-on-top500].
- The Nano4 system's NVIDIA **HGX H200** partition (with Direct-to-Chip liquid cooling, 81.55 PFLOPS) ranked **#29 on TOP500** [secondary — same]. Also in the build: ESC8000-E12 servers with **NVIDIA MGX H200** + NVLink Bridge [secondary — same].

**Legacy/in-market lines (carried into 2026)**:
- **ESC8000A-E13P** (MGX, up to 8× NVIDIA H200 NVL dual-slot 600 W, BlueField-3) — shown SC24 [secondary — https://www.techedt.com/asus-unveils-next-generation-infrastructure-solutions-at-sc24-with-advanced-cooling].
- **ESC N8-E11V** (NVIDIA HGX H200) and **ESC N8A-E13** (Blackwell platform) — SC24 2024 showcase [secondary — same]. ⚠️ The task references **ESC N8A-E12**; sources surfaced ESC N8A-E13 (Blackwell) and the earlier ESC N8A-E12 (HGX H100, per press.asus.com CloudFest 2024: 7U dual AMD EPYC 9004 + 8× H100) [official — https://press.asus.com/news/press-releases/asus-mgx-powered-server-cloudfest-2024-showcase/]. Confirming a 2026-spec "ESC N8A-E12" B200/B300 variant was not possible — flagged as a gap.

### 3.2 Customer wins
- **Taiwan NCHC Nano4 AI supercomputer** (national HPC center, academia + industry workloads) — ESC NM2N721-E1 (GB200 NVL72) + HGX H200 systems; #29 TOP500 [secondary — TechPowerUp above].
- **Ubilink** (via ASUS partnership): 45.82 PFLOPS green-energy AI data center [secondary — https://www.tweaktown.com/news/101772/asus-unveils-its-new-ai-pod-complete-rack-of-liquid-cooled-nvidia-gb200-nvl72-servers/index.html].
- ⚠️ No 2026 enterprise/CSP customer deal values found publicly — gap.

### 3.3 Revenue/market context
- ASUS server/AI financials for 2026 were not isolated in this research — gap. (ASUS reports server business within its broader segments; AI POD momentum is the stated growth engine.)

---

## 4. MSI

### 4.1 AI server offerings & 2026 launches

