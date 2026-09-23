---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-4-2026-09-22-5
title: "Supplementary / Complementary Research Pass — round 4 (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Broadcom", "China", "EU", "Meta", "Nvidia", "United States"]
dates: ["2019-03-11", "2019-12", "2020-04-16", "2020-04-27", "2025-12", "2025-12-02", "2026-02-26", "2026-03-02", "2026-06-09", "2026-07-10", "2026-08-31", "2026-09-02", "2026-09-22"]
keywords: ["research", "acquisition", "amd", "asic", "compute", "ethernet", "fp4", "gpu", "gpus", "hbm4", "helios", "latency"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1672, 1721]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 7cc5db44570e698ed8a8326bfabe60ff928fa97dc4159d617cf46c8787c8879e
---

# Supplementary / Complementary Research Pass — round 4 (2026-09-22)

## Supplementary / Complementary Research Pass — round 4 (2026-09-22)

**Scope note:** Base §§1–8, Supplementary Passes #1 (§§A–M), #2 (§§N–X), round 2 (§§N–S), and round 3 (R3-A…H, AA…II, plus Pass #3) already exist. This round-4 pass adds only material not present in any of them. All earlier sections were left untouched. Provenance legend from the base applies. All facts as of 2026-09-22.

### R4-A. HPE + AMD Helios — the Juniper scale-up switch (new 2026 networking detail)

Round-3 §R3-C mentioned a "Helios platform expected to bring further acceleration" with no detail. The networking architecture is now documented [secondary — Tom's Hardware 2025-12, The Register 2025-12-02, TechRadar, ai-daily.news; HPE statements]:

- HPE is the **first major OEM to commit to AMD's Helios rack-scale AI architecture**, offering Helios AI Racks globally starting **2026** (announced ahead of HPE Discover Barcelona, December 2025) [vendor-reported — HPE].
- Helios rack: **72 AMD Instinct MI455X GPUs + EPYC "Venice" CPUs**, double-wide liquid-cooled chassis (Meta Open Rack Wide), up to **2.9 exaFLOPS FP4 per rack**, 31 TB HBM4, 260 TB/s aggregated scale-up bandwidth, 18 compute trays × (4 MI455X + 1 Venice) [vendor-reported — AMD/HPE].
- **Scale-up fabric: UALink-over-Ethernet (UALoE)** — an Ethernet-based scale-up fabric developed by AMD + Broadcom, explicitly positioned against NVIDIA NVLink; scale-out via Ultra Ethernet Consortium-aligned hardware; Pensando components for networking offload; ROCm software stack [vendor-reported — AMD; secondary — Tom's Hardware].
- **Purpose-built Juniper Networks scale-up switch** based on **Broadcom Tomahawk 6 silicon (102.4 Tbps)** — the GPU-interconnect backbone of the rack, running the UALink protocol over Ethernet instead of proprietary UALink hardware [vendor-reported — HPE; secondary — The Register].
- HPE branding: "HPE Scaling Network Solution" — Antonio Neri (HPE CEO): "our purpose-built HPE Scaling Network Solution … providing our cloud service provider customers with faster deployments, greater flexibility and lower risk in how they scale AI computing across their businesses" [vendor-reported — HPE via vmvirtualmachine.com].
- Scale-out NICs: **AMD Pensando Vulcano NICs** handle inter-rack scale-out networking [secondary — mlq.ai].
- Timeline: engineering samples and low-volume production targeted **2H 2026**, mass production ramp expected by Q2 2027; AMD denies reported delays, stating Helios "on target for 2H 2026" [secondary — mlq.ai].
- Significance for Phase A: this is the first named Juniper-heritage AI-fabric switch inside the HPE Networking portfolio (alongside the CX 10040/Pensando DC line), and the first public HPE application of Tomahawk-6 outside the Aruba CX 9300 — directly relevant to the HPE (Aruba+Juniper) competitive framing in round-3 §U.

### R4-B. NVIDIA DOCA SDK 2026 — data-plane release line (new, not the DPF orchestration train)

Pass §DD covered the **DOCA Platform Framework (DPF)** orchestration train (v25.x/v26.x). The **DOCA SDK data-plane line** runs separate 3.x releases [official — NVIDIA DOCA docs, docs.nvidia.com/doca/sdk]:

- Version cadence in 2026: **DOCA 3.3.0** (NGC container dated 2026-02-26), **3.4.0** (2026-06-09; firmware xx.49.1014, MFT 4.36.0-147, BlueField ATF/EDK2 4.15.0, SHARP 2.50, HPC-X 2.50), **3.5.0** (NGC dated 2026-08-31; release-notes page updated 2026-09-02) [official — NVIDIA docs, NGC catalog].
- **DOCA VERBs — new Multi-Rail API**: supports NIC resiliency for VR [2:1] and load-balancing for non-NVIDIA comm libraries [official — DOCA SDK changes page, Sep 2026].
- DOCA-Host/driver updates: **FRMR (Fast Registration Memory Region) pools replace the MR cache** — user-space control of pool sizes and aging via iproute2 RDMA tool; Adaptive-ReTransmission Q-counters exposed to all function types, not just physical functions [official — DOCA changes page].
- **BlueField-3 firmware: PCIe in-band SPDM attestation** added (host reaches the PSC over in-band PCIe channels) [official — DOCA changes page].
- SNAP virtio-fs service 1.5.0 targets **DOCA 3.2.0**; known issue: FUSE_STATX on modern kernels can hang the virtiofs driver; hot-unplug of PCIe functions with in-flight IO can time out [official — NVIDIA SNAP release notes, updated 2026-03-02].
- Warning for readers: do not conflate the DOCA SDK 3.x data-plane line with DPF v26.x orchestration (see §DD) — they version independently [editorial note].

### R4-C. Quantum-X800 XDR — switch spec depth (new detail)

Pass §K covered orderability; the datasheet (via Dell OEM mirror) gives the engineering detail [official — NVIDIA Quantum-X800 datasheet via delltechnologies.com]:

- **Q3400-RA** (4U): **144× non-blocking 800 Gb/s XDR ports** (72 twin-port OSFP224 cages), **115.2 Tb/s** aggregate throughput, Quantum-3 ASIC, sub-100 ns port-to-port latency; ~2,900 W typical (up to ~7,000 W with active cables); 200–240 V AC [official — datasheet; secondary — gpusmith.com verified 2026-07-10].
- **Q3200-RA** (2U): architecturally **two fully independent 36-port switches** (2× 28.8 Tb/s) in one chassis — cannot be pooled into a single 72-port fabric; 18 OSFP224 cages per internal switch; designed for dual-plane AI networks or isolated compute/storage fabrics [official — datasheet; vendor-reported via reseller FAQ].
- **Q3401-RD**: DC-power variant (48–54 V DC), same 144× 800G / 115.2 Tb/s envelope [official — datasheet].
- Common: i3-8100H management CPU, **NVOS software**, IRoT-based hardware security (CPU/CPLD/switch IC), dedicated **400 Gb/s OSFP in-band management port (UFM)**, 1-year service [official].
- Fabric scale: NVIDIA reference design scales a single two-tier fat tree of Q3400-RA to **10,368 NICs**; leaf or spine roles in two-tier fat trees; an InfiniBand Subnet Manager is required, UFM recommended for production [vendor-reported — gpusmith.com; secondary — aicplight.com FAQ].
- Note: ports A17/A18/B17/B18 on the Q3200-RA do not support split mode or LACC (Linear Active Copper Cable) — cable-planning constraint [vendor-reported — reseller FAQ].

### R4-D. Mellanox heritage — acquisition and milestone timeline (new, consolidated)

No earlier pass documented the heritage arc itself; it is the causal background for §§2.x:

- **Announced March 11, 2019**: NVIDIA to acquire Mellanox for **$125/share in cash, $6.9B total** — NVIDIA's largest acquisition at the time, immediately accretive to non-GAAP gross margin/EPS/free cash flow expected [official — NVIDIA press release, nvidianews.nvidia.com].
- **Regulatory path**: US waiting period expired; EU approval December 2019; Mexico approved; China SAMR conditional approval April 16, 2020 (no unreasonable bundling in China). **Closed April 27, 2020** [official — NVIDIA].
- Pre-close Mellanox performance (2020): revenue +22% to $1.33B, net income +48% to $394M [secondary — datacentrereview.com].
- Post-acquisition integration milestones: **Cumulus Networks** acquired 2020 (~$100M est.) to complete the stack (NIC + switch + OS); Mellanox team credited with InfiniBand/Spectrum/RDMA; BlueField DPUs, ConnectX line, Quantum/Quantum-2/Quantum-X switches, UFM [secondary — GitHub NVIDIA deep-dive, 2026].
- 2026 state: the Mellanox-built unit is the core of NVIDIA's networking business — **$11B in the most recent quarter cited in round-2 §O (+263% YoY)**; **#1 in DC Ethernet switching by revenue** (IDC Q1/Q2 2026, §§2.1/G); CEO: "We're now the largest networking company in the world" [vendor-reported/secondary].
- Note (from §HH.5): the $6.9B figure and April 27, 2020 close are well-established but re-verified here against primary NVIDIA sources [official].

