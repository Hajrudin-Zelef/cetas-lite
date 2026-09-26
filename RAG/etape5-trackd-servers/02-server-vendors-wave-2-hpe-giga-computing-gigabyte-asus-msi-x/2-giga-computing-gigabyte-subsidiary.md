---
id: etape5-trackd-servers/02-server-vendors-wave-2-hpe-giga-computing-gigabyte-asus-msi-x/2-giga-computing-gigabyte-subsidiary
title: "2. Giga Computing (GIGABYTE subsidiary)"
domain: server-vendors-wave-2-hpe-giga-computing-gigabyte-asus-msi-x
role: deep-dive
task: reference
actors: ["AMD", "DeepSeek", "Intel", "Nvidia"]
dates: ["2026-03", "2026-06-16", "2026-07", "2026-11", "2027-03"]
keywords: ["agentic", "amd", "benchmarks", "blackwell", "compute", "deepseek", "disclosure", "ethernet", "gpu", "gpus", "hyperscaler", "inference"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [543, 596]
section: "Server Vendors — Wave 2: HPE + Giga Computing (Gigabyte) + ASUS + MSI + xFusion"
sha256: f7b1e4f173876e6fb23e57c6b712baa80e82a4d5dd19c565ed4aa04fcdb10688
---

# 2. Giga Computing (GIGABYTE subsidiary)

## 2. Giga Computing (GIGABYTE subsidiary)

### 2.1 2026 AI server lineup & launches (dates/specs)

**GB200 NVL4 — XN24-VC0-LA61 (announced Jan 26, 2026)**
- New AI/HPC server on the **NVIDIA GB200 NVL4** platform; debut at **SCA/HPC Asia 2026, Osaka**; heterogeneous CPU+GPU architecture with liquid cooling [official — GIGABYTE press release, https://www.gigabyte.com/press/news/2360].
- **Customer/design win**: XN24 selected for the **RIKEN Center for Computational Science (R-CCS) next-generation HPC-Quantum hybrid platform** — integrated into **FugakuNEXT** (successor to Japan's flagship Fugaku supercomputer) using the **NVIDIA CUDA-Q** platform for hybrid quantum–GPU supercomputing research [official — same release].

**HGX B300 / B200 flagships (OCP EMEA Summit 2026, Barcelona)**
- Two flagship servers on **NVIDIA HGX B300 (liquid-cooled)** and **NVIDIA HGX B200 (air-cooled)**, OCP-standard platforms for AI data centers; plus air- and liquid-cooled compute nodes for rack-level integration [secondary — https://finance.biggo.com/news/dh3B3Z0B-PfaobXfvRh6]. ⚠️ Exact model numbers for the B300 flagships were not stated in the sources consulted.

**G893 series — 8-GPU HGX flagships** (note: launched ~Dec 2024; in-market through 2026)
- **G893-SD1-AAX5** (Intel 5th Gen Xeon) and **G893-ZD1-AAX5** (AMD EPYC): **8× NVIDIA HGX B200**, NVLink with **1800 GB/s GPU-to-GPU bandwidth** [secondary — https://www.pressebox.de/pressemitteilung/giga-computing-technology-co-ltd/GIGABYTE-Unveils-Next-Generation-AI-Flagship-Servers-G893-SD1-AAX5-G893-ZD1-AAX5/boxid/1230117].

**G4L3 / G894 series — HGX B200 thermal-flex portfolio (late 2025)**
- **G4L3-SD1-LAX5** (Intel Xeon) / **G4L3-ZD1-LAX5** (AMD EPYC): 4U direct-liquid-cooled, dedicated CPU & GPU chambers, HGX B200 [secondary — https://www.techpowerup.com/338575/giga-computing-unveils-liquid-and-air-cooled-gigabyte-ai-servers-accelerated-by-nvidia-hgx-b200-platform].
- **G894-AD1-AAX5** (Intel Xeon 6900) / **G894-SD1-AAX5** (Intel Xeon 6700 & 6500): air-cooled HGX B200 [secondary — same source].

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

