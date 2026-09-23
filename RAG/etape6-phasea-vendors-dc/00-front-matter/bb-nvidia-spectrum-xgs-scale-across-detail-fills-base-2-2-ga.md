---
id: etape6-phasea-vendors-dc/00-front-matter/bb-nvidia-spectrum-xgs-scale-across-detail-fills-base-2-2-ga
title: "BB. NVIDIA Spectrum-XGS — scale-across detail (fills base §2.2 gap)"
domain: front-matter
role: reference
task: actor-profile
actors: ["CoreWeave", "Nvidia", "Oracle"]
dates: ["2025-08-22", "2025-10", "2026-01", "2026-08", "2026-09-01"]
keywords: ["nvidia", "accelerator", "asic", "backlog", "cpo", "dsp", "ethernet", "gpu", "latency", "nvlink", "optics", "revenue"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [854, 905]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 6184fbe66f89b071cb2396e24a4cae91d8123a09f4c4492dbbf95c3843e5126f
---

# BB. NVIDIA Spectrum-XGS — scale-across detail (fills base §2.2 gap)

### BB. NVIDIA Spectrum-XGS — scale-across detail (fills base §2.2 gap)

- **Spectrum-XGS** was announced **August 22, 2025** (Hot Chips, Palo Alto) as a "scale-across" technology extending Spectrum-X Ethernet to interconnect **multiple geographically distributed data centers into unified "giga-scale AI super-factories"** — a third pillar beyond scale-up (within servers) and scale-out (within data centers) [official — NVIDIA press release via syndication].
- Delivery mechanism: primarily **software and firmware updates to existing Spectrum-X switches and ConnectX SuperNICs**, not new silicon — allowing operators to reuse existing hardware investments [vendor-reported — NVIDIA; secondary — TechPowerUp].
- Key algorithms: **auto-adjusted distance congestion control** optimized for long-haul links, **precision latency management** to minimize jitter, and **comprehensive end-to-end telemetry** across sites [vendor-reported].
- Performance claim: **nearly doubles NCCL (NVIDIA Collective Communications Library) throughput** for multi-GPU, multi-node training jobs and large-scale experiments across distributed sites [vendor-reported].
- First customer: **CoreWeave** — co-founder/CTO Peter Salanki: XGS will give its customers "access to giga-scale AI" across CoreWeave's distributed facilities [vendor-reported].
- Jensen Huang: "With NVIDIA Spectrum-XGS Ethernet, we add scale-across to scale-up and scale-out capabilities to link data centers across cities, nations and continents into vast, giga-scale AI super-factories" [vendor-reported].
- 2026 deployment confirmation beyond CoreWeave was not located [gap].

### CC. Mellanox heritage: ConnectX end-of-life landscape (new)

The file frames NVIDIA Networking as "Mellanox heritage"; NVIDIA's official EOL notices define the installed-base sunset timeline:

- **ConnectX-4 EN cards** (MCX413A/MCX414A/MCX513A/MCX514A): last-time-buy **Oct 31, 2021**, last ship **Apr 30, 2022** (reason: low market interest; replacement = ConnectX-5) [official — NVIDIA network.nvidia.com EOL notice LCR-000674].
- **ConnectX-4 OCP variants** (MCX445B/MCX445N): LTB **Aug 31, 2018**, last ship **Feb 28, 2019** [official — LCR-000266].
- **ConnectX-3 VPI/EN**: LTB **Dec 31, 2019**, last ship **Jun 30, 2020** (reason: improved replacement products / low interest in older technology) [official — LCR-000443].
- Software support continues: the mlx5 driver in DPDK 24.11 still supports **ConnectX-4 through ConnectX-7 and BlueField through BlueField-3** — i.e. the full Mellanox-lineage device tree remains driver-current in 2026 [secondary — asterfusion/helium_dpu GitHub mirror of DPDK docs].
- Practical note: 2016–2019 Mellanox hardware (ConnectX-3/4) is now past EOL; 2020–2021 (ConnectX-5/6) sits in maintenance/legacy phase while ConnectX-7/8/9 carry current AI-fabric deployments — consistent with the ConnectX-9 GA covered in base §2.4.

### DD. NVIDIA networking software stack 2026: DOCA Platform Framework + BlueField firmware (new)

Base §2.5 mentioned BlueField-4 and DOCA in passing. The 2026 software train:

- **DOCA Platform Framework (DPF)** versions [official — NVIDIA doca-platform GitHub repo]:
  - **v26.10.0 GA — August 2026** (provisioning/orchestration of BlueField DPUs in Kubernetes).
  - v26.4.0 (2026); v25.10.1 GA — January 2026 (adds NIC-mode → DPU-mode transitioning during provisioning).
  - DPF v25.7.0 (2025): DPU Zero Trust (beta) with VPC service, SNAP-based storage volume management, auto-discovery of DPUs.
- **BlueField-3 DPU NIC firmware v32.43.2408 / v32.43.2026**: mandatory DPA (Data Path Accelerator) outbox non-blocking-mode change and DPA thread-context internal API change — customers programming DPA (e.g. virtio-net/blk/fs, NVMe stacks) must adapt [official — NVIDIA docs].
- **BlueField-4** (GTC DC, October 2025 announcement): **126 billion transistors**, **64-core "Grace" CPU** (Neoverse V2-class), **ConnectX-9 networking**, 800G class; expected in early availability as part of **Vera Rubin platforms in 2026**; PCIe Gen6 class [vendor-reported — NVIDIA; secondary — ServeTheHome].
  - BlueField-4 ecosystem partners named: **CoreWeave, Oracle Cloud Infrastructure, Palo Alto Networks** [vendor-reported — SDxCentral, GTC 2025 briefing].
  - **BlueField-5 hinted for 2028** with the Feynman architecture, per NVIDIA's stated roadmap cadence [vendor-reported — SDxCentral].

### EE. InfiniBand roadmap beyond XDR: X1600 and CPO detail (new)

- **X1600 InfiniBand/Ethernet network processor**: named on NVIDIA's AI-data-center roadmap (presented via SC'24 roadmap reporting) as the networking element of the **Vera Rubin** generation (R100, 2026), alongside NVLink 6 switches (3,600 GB/s) and CX9 SuperNICs (1,600 Gb/s); **Rubin Ultra** (2H 2027) follows with NVLink 576-class scale-up, and **Feynman** (2028) after [secondary — Converge Digest roadmap summary, 2026].
- **CPO product detail** (technical specs, extending base §2.3 / round-1 §K):
  - **Quantum 3450-LD (CPO InfiniBand)**: 144× 800 Gb/s ports, **115 Tb/s** aggregate, four Quantum-X CPO sockets, **liquid-cooled chassis** [secondary — TechPowerUp, NVIDIA CPO briefing].
  - **SN6810 / SN6800 (Spectrum-X CPO Ethernet)**: 128× 800G (102.4 Tb/s) / 512× 800G (**409.6 Tb/s**), 512-port radix; InfiniBand CPO = monolithic ASIC + six CPO modules (36× 800G per module); Spectrum-X CPO = multi-chip (central packet engine + 8 SerDes chiplets); both on **224 Gb/s lanes, 4 lanes/port** [secondary].
  - CPO power claim: **3.5× lower power** than pluggable optics; signal loss 22 dB → 4 dB; optical interconnect power in a 400,000-GPU scenario from ~72 MW to **21.6 MW**; laser source 2 W/port vs 10 W; optical engine 7 W vs 20 W DSP [vendor-reported — NVIDIA via TechPowerUp].
  - Scheduling: Quantum-X (IB) CPO H2 2025, Spectrum-X (Ethernet) CPO **H2 2026** per original announcements — ship confirmations still outstanding (see §L.8) [vendor-reported schedule; gap on confirmation].

### FF. Dell Q2 FY2027 earnings — September 1, 2026 (refreshes base §1.6)

Base §1.6 carried FY2026 figures. The Q2 FY2027 print (reported 2026-09-01) reshapes the Dell networking context [official — Dell 8-K; secondary — press]:

- Record revenue **$47.0B (+58% YoY)**; GAAP diluted EPS $6.34 (+273%); non-GAAP $7.04 (+203%) [official].
- **AI-optimized server revenue: $16.4B (+100% YoY)**; record AI-server orders **$60.9B**; record AI-server backlog **$95B** (~5.8 quarters at the quarter's revenue run-rate) [official].
- **Traditional Servers and Networking: $10.5B (+122% YoY)** — a record quarter; ISG total $31.8B (+89%); ISG operating income $4.8B (+225%, 15.0% margin) [official].
- Full-year FY2027 revenue guidance raised by $25B to **$192B** [official — Dell; secondary — sharesify/pulse2].
- **Networking revenue remains undisclosed** — Dell still does not break out switching revenue or AI-fabric attach rates [gap, unchanged].

