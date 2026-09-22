---
id: etape5-trackd-servers/03-data-center-networking-for-ai-2026/2-nvidia-spectrum-x-spectrum-xgs-ethernet-networking-platfor
title: "2. NVIDIA Spectrum-X / Spectrum-XGS Ethernet Networking Platform"
domain: data-center-networking-for-ai-2026
role: deep-dive
task: actor-profile
actors: ["Broadcom", "CoreWeave", "Meta", "Nvidia", "xAI"]
dates: ["2025-09", "2026-07", "2026-07-23"]
keywords: ["ethernet", "nvidia", "asic", "blackwell", "cpo", "energy", "gpu", "gpus", "inference", "latency", "llama", "nvlink"]
source: docs/RAG/etape5_trackD_servers.md
source_anchor: ""
source_lines: [745, 806]
section: "Data-Center Networking for AI (2026)"
sha256: 433fc9c1ad4c5f17f6beafa3a2d17c4659a18d0b2a14f97759208e8a1ad0f2a7
---

# 2. NVIDIA Spectrum-X / Spectrum-XGS Ethernet Networking Platform

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

- Spectrum-X800 SN5600 switches are based on the Spectrum-4 switch ASIC (800G ports, 51.2T). [secondary] (tomshardware citing NVIDIA)

---

## 3. InfiniBand vs Ethernet in 2026 AI fabrics

### 3.1 2026 state of play

- Both NVIDIA Quantum-X800 (InfiniBand) and Spectrum-X800 (Ethernet) now ship at 800 Gb/s per port; Quantum-X800 at 115.2 Tb/s total switching capacity vs Spectrum-X800 51.2 Tb/s per the ascentoptics comparison (note: different switch configurations; treat as [secondary], vendor/product-page sourced). [secondary]
- 8-node reference numbers (secondary, 2026): IB NDR 400G (Quantum-2) ~350 GB/s effective all-reduce at sub-1 µs latency with SHARP in-switch reductions; well-tuned 400GbE RoCEv2 ~270–290 GB/s at ~2.4 µs; Spectrum-X within 5–10% of IB. [secondary] (rdp.in, temperature2.com)
- Decision guidance in 2026 (secondary consensus): IB for large training where all-reduce latency compounds every step; Ethernet (Spectrum-X/RoCE) for inference, multi-tenant clouds, and open-ecosystem deployments. [secondary]
- **Meta example:** two 24,576-H100 clusters run side by side — one RoCE-over-Ethernet (Arista 7800), one NVIDIA Quantum-2 InfiniBand — both reaching 90%+ network utilization training Llama 3. [secondary] (temperature2.com)
- Dell'Oro Q1 2026: Ethernet ~2/3 of AI-backend DC switch sales; IB sales tripled on the Blackwell Ultra 800G ramp (brownfield-upgrade caveat noted). [independent]

### 3.2 NVIDIA InfiniBand updates: Quantum-X800 / Quantum-3 / Quantum-X

- **Quantum-X800 (XDR InfiniBand, 800G/port):** shipping with NVIDIA's Blackwell Ultra platform ramp (Dell'Oro Q1 2026: IB sales "more than tripled ... supported by the ramp of 800 Gbps switches shipping with NVIDIA's Blackwell Ultra platform"). [independent]
- Quantum-X800 features SHARP v4 in-network computing for collective operations (9× improvement claim in the ascentoptics comparison table — [secondary], treat improvement factor as vendor-reported via that source). [secondary]
- NVIDIA newsroom framed the IB vs Ethernet choice for xAI: Colossus chose Spectrum-X Ethernet over InfiniBand — a landmark design decision for the world's largest AI supercomputer at the time. [official]
- **Quantum-X silicon photonics switches:** NVIDIA's announced Quantum-X and Spectrum-X silicon photonics networking switches "enable AI factories to connect millions of GPUs across sites while reducing energy consumption and operational costs" — referenced in the Spectrum-XGS release. [official]
- TrendForce roadmap: Quantum-2 (400G/port, 51.2 T) → Quantum-X800 (800G, 115.2 T) → Quantum-X1600 (1.6T, 230.4 T, paired with Rubin) → Quantum-X3200 (3.2T, paired with Feynman). [secondary]
- **Uncertainty flags:** (a) No Quantum-3 shipping announcement was found in the sources retrieved; the platform table lists Quantum-2 → Quantum-X800 → Quantum-X1600, so "Quantum-3" may be a skipped/renamed generation — do not assert its existence. (b) Quantum-X800 CPO/silicon-photonics variant shipping dates were described as "early 2026" roadmap (per the github silicon-photonics research note) but no dated production shipment was confirmed. (c) The `<100 ns` latency figure in the ascentoptics table vs "sub-1 µs" elsewhere is inconsistent — do not quote either as settled fact.

### 3.3 Broadcom vs NVIDIA "scale-out tech war"

- TrendForce (2025, still cited 2026): frames the AI fabric contest as Broadcom (Tomahawk/Thor/UEC) vs NVIDIA (Spectrum-X/Quantum-X) camps; Marvell Teralynx 10 and Cisco Silicon One G200 (both 51.2T, 2023) also compete; Cisco has CPO prototypes. [secondary]

---

