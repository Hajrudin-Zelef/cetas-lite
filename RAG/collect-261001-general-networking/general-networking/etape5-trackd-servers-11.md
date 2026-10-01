---
id: collect-261001-general-networking/general-networking/etape5-trackd-servers-11
title: "Step 5 — Server Vendors + AI Server Market + Data-Center Networking"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom", "CoreWeave", "Meta", "Nvidia", "xAI"]
dates: ["2025-09", "2026-07", "2026-07-23"]
keywords: ["asic", "cpo", "ethernet", "gpu", "gpus", "hyperscaler", "inference", "latency", "llama", "nvidia", "nvlink", "revenue"]
source: docs/RAG/collect-261001-general-networking/etape5_trackD_servers.md
source_anchor: ""
source_lines: [720, 777]
sha256: 1ee296e839dd51036500ffa86cbdf7519a6a14f5adc4568151ca55a2aa2c7a68
---

# Step 5 — Server Vendors + AI Server Market + Data-Center Networking

- **Etherlink AI portfolio:** Arista's AI-optimized Ethernet portfolio; customer base expanded to **over 100 customers** (Q2 2026), up from just 4–5 in 2024. [vendor-reported] (cryptobriefing summarizing Arista Q2 2026 earnings)
- During Q2 2026 Arista introduced its **7506c/7 Series AI fabric portfolio** and **1.6 Tbps platforms with liquid-cooled options** (also reported as the **7060XE7 Series 1.6 Tbps platforms**) for the largest AI training/inference clusters. [vendor-reported] [secondary]
- Arista's official Q2 2026 release notes: "Introduced 1.6 Tbps AI fabric platforms, including liquid-cooled options optimized for scale-up, scale-out, and scale-across networks." [official]
- The 7060X6 (Tomahawk 5-based 800G) and 7800R4/7700R4 (modular 800G systems) lines are Arista's shipping 800G portfolio for AI backends; Meta's RoCE-over-Ethernet 24,576-GPU H100 cluster for Llama 3 used Arista 7800 systems side-by-side with a Quantum-2 InfiniBand cluster, both reaching 90%+ network utilization. [secondary]
- **Uncertainty flag:** I did not independently confirm per-SKU shipment volumes for the 7060X6/7800R4/7700R4; the task's "shipments" granularity is not disclosed at SKU level in the sources found. Do not fabricate SKU shipment numbers.

### 1.4 Cisco Nexus 9000 800G / Silicon One

- Cisco introduced **Silicon One G300** — its own 102.4 Tbps networking silicon — debuting in **liquid-cooled N9000 and 8000 series switches**, aimed at competing with Broadcom and NVIDIA switching chips. [secondary] (ainvest)
- Simultaneously, Cisco ships **N9100 series switches explicitly powered by NVIDIA Spectrum-X Ethernet silicon**, running Cisco NX-OS on NVIDIA hardware. [secondary] (ainvest)
- Cisco FY2026 hyperscaler AI design wins: three new wins in Q4 alone — one **Silicon One P200 scale-across** deployment, one **G200 scale-out** project, and one optical line-system deployment; management flagged line-of-sight to more wins over the next six months across G300, G200, P200 and A100 platforms. [vendor-reported] (infotechlead)
- TrendForce platform table (2025 vintage, cited 2026): Cisco Silicon One G200 series at 51.2 Tbps (2023) alongside Marvell Teralynx 10; CPO prototypes. [secondary]

### 1.5 NVIDIA Spectrum-X switches (SN5000-class / SN5600 / Spectrum-X800)

- **Spectrum-X800 Ethernet switch** (SN5600-class): 800 Gb/s per port, 51.2 Tb/s switching capacity (per temperature2.com / TrendForce table). Based on the **Spectrum-4 switch ASIC**; SN5600 = 64-port 800GbE. [secondary]
- **Spectrum-6** (announced July 2026): 102.4 Tbps Ethernet switch system, 2× previous-generation capacity, integrated into the Rubin-architecture stack (Vera CPU, Rubin GPU, NVLink 6, ConnectX-9 SuperNICs, BlueField-4 DPUs, Spectrum-6). [vendor-reported] (cxotoday citing NVIDIA)
- TrendForce roadmap table: Spectrum-X800 (800G/port, 51.2 Tbps) → Spectrum-X1600 (1.6T/port, 102.4 Tbps, paired with Rubin) → Spectrum-X3200 (3.2T/port, 204.8 Tbps). ConnectX-8 (800G, PCIe 6.0) → ConnectX-9 (1.6T, PCIe 7.0) → ConnectX-10 (3.2T, PCIe 8.0). [secondary]

### 1.6 Celestica's position

- Celestica regained #1 in Ethernet AI backend networks in Q1 2026 per Dell'Oro (above). Celestica is NVIDIA's primary Ethernet-switch manufacturing partner for Spectrum-X systems; its ranking reflects white-box/OEM volume into AI fabrics. [independent] [secondary]

---

## 2. NVIDIA Spectrum-X / Spectrum-XGS Ethernet Networking Platform

### 2.1 Platform composition (per NVIDIA)

- Spectrum-X = Spectrum Ethernet switches + Spectrum-X/ConnectX SuperNICs + BlueField DPUs + LinkX cabling and transceivers (+ Spectrum-XGS for multi-DC). [official]
- Performance claim: 1.6× greater bandwidth density than off-the-shelf Ethernet for multi-tenant hyperscale AI factories, "including the world's largest AI supercomputer." [official] (investor.nvidia.com Spectrum-XGS release)
- NVIDIA reports Spectrum-XGS "nearly doubles the performance of the NVIDIA Collective Communications Library (NCCL)," accelerating multi-GPU/multi-node communication across geographically distributed clusters. [official]
- Delivery mechanism for XGS: primarily **software and firmware updates to existing Spectrum-X switches and ConnectX SuperNICs**, not new silicon. [secondary] (techpowerup)

### 2.2 Spectrum-XGS (scale-across), available now

- Announced by NVIDIA (investor.nvidia.com; circa September 2025 based on the release cadence and coverage; the "available now" release was covered at Hot Chips). Adds a **"scale-across"** axis beyond scale-up/scale-out, linking data centers across cities, nations and continents into "giga-scale AI super-factories." [official]
- Technologies: auto-adjusted distance congestion control, precision latency management, end-to-end telemetry; algorithms dynamically adapt the network to inter-facility distance. [official]
- **CoreWeave** named as one of the first adopters, connecting its data centers into a single supercomputer; CTO quote from Peter Salanki in the release. [official]
- Spectrum-XGS is fully integrated into the Spectrum-X platform. [official]
- Arista has publicly identified **scale-across** as a separate networking opportunity (its 1.6T platforms target scale-up/scale-out/scale-across), mirroring NVIDIA's framing. [vendor-reported] [secondary]
- **Uncertainty flag:** the exact announcement date of the Spectrum-XGS release could not be verified from the investor.nvidia.com page text fetched; the techpowerup URL (340218) and Hot Chips mention place it in the September 2025 timeframe. Treat the date as approximate.

### 2.3 Spectrum-6 (Rubin era, July 2026)

- Announced ~July 23, 2026 (cxotoday). 102.4 Tbps Ethernet switch system, 2× previous-gen capacity; designed to connect hundreds of thousands of GPUs with improved power efficiency; part of the Rubin AI infrastructure stack alongside Vera CPU, Rubin GPU, NVLink 6 switches, ConnectX-9 SuperNICs, BlueField-4 DPUs. [vendor-reported] (cxotoday citing NVIDIA)
- NVIDIA quarterly networking revenue context (recent quarter as of July 2026): **$11B**, "massive over 250% growth over the previous year"; CEO Jensen Huang declared "We're now the largest networking company in the world." [vendor-reported] (cxotoday citing NVIDIA)
- Per Dell'Oro figures cited by secondary press (2026): NVIDIA ~$2.1B quarterly data-center Ethernet **switch** revenue, 21.5% share — the largest Ethernet-switch vendor, up from <4% two years ago. (Note: this is switch revenue only, not total NVIDIA networking revenue.) [secondary]

### 2.4 Spectrum-X adoption and performance claims

- **xAI Colossus (flagship win):** 100,000 Hopper GPUs (Memphis) built in 122 days using Spectrum-X Ethernet (SN5600 64-port 800GbE switches, BlueField-3 SuperNICs, 400GbE per GPU) for the RDMA network — instead of InfiniBand. NVIDIA claims: 95% data throughput maintained, zero application latency degradation/packet loss from flow collisions across all three fabric tiers; standard Ethernet would deliver only 60% throughput with thousands of flow collisions. xAI doubling to 200,000 Hopper GPUs; Colossus 2 adds ~30,000 GB200s. [official] (nvidianews.nvidia.com)
- Spectrum-X closes 80–90% of the InfiniBand gap on NCCL all-reduce — within ~5% at 8 nodes, widening at 64+ nodes (cited by rdp.in from Spheron, 2026). [secondary]
- Inference/multi-tenant positioning: Spectrum-X Ethernet ~5–10% behind IB on effective all-reduce at 8 nodes, ~1.7 µs vs ~0.9 µs latency; recommended for inference, multi-tenant and mixed clouds. [secondary]
- DriveNets claimed (Aug/Sep 2026) its AI Fabric Ethernet solution achieved 6% better NCCL performance than NVIDIA InfiniBand (Quantum-X) on average and 15% better than Spectrum-X on a WhiteFiber Iceland deployment — vendor claim using SemiAnalysis test data; treat as [vendor-reported] and un-audited.

### 2.5 Spectrum-4 ASIC context

