---
id: etape6-phasea-vendors-dc/00-front-matter/d-nvidia-spectrum-family-completion-mellanox-heritage-new-de
title: "D. NVIDIA Spectrum family completion (Mellanox heritage — new detail)"
domain: front-matter
role: reference
task: actor-profile
actors: ["Intel", "Nvidia"]
dates: ["2025-12-09", "2026-07"]
keywords: ["nvidia", "asic", "benchmark", "cpo", "distribution", "dsp", "ethernet", "gpu", "intel", "latency", "optics", "throughput"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1374, 1406]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: ad6b8bb7e229c9e4cb20611af90208f59d668700e29b1f038f05b125804c0edf
---

# D. NVIDIA Spectrum family completion (Mellanox heritage — new detail)

- Dell KB **000220710** ("Dell Networking: SONiC OS platform and package support") documents the Enterprise SONiC bundle matrix: **Cloud Standard, Cloud Premium, Enterprise Standard, Enterprise Premium**, plus **Edge Standard** — the Edge Standard bundle is supported **only on N3248TE-ON, N3248PXE-ON, N3248X-ON, E3248P-ON, E3248PXE-ON**; all other PowerSwitch platforms cannot run the Edge bundle [official — Dell KB 000220710].
- Ecosystem (vendor-reported, Dec 2025): **Aviz Networks "Aviz Certified Community SONiC"** — turnkey enterprise-grade community SONiC with FTAS test reports across 100+ deployment scenarios (IP CLOS, EVPN/VXLAN, MLAG, Spectrum-X validated runbooks), certified images for Accton/Edgecore, Celestica, Cisco, **NVIDIA**, Wistron; claims "**50%+ savings**" from software-first operations [vendor-reported — BusinessWire, 2025-12-09]. Relevant as the third-party SONiC support lane competing with Dell's own distribution.

### D. NVIDIA Spectrum family completion (Mellanox heritage — new detail)

Earlier passes covered Spectrum-4 (SN5600/5610/SN5600D), SN5400 and Spectrum-6. The rest of the shipping family, verified via the Lenovo OEM product guide and NVIDIA pages [official — Lenovo Press lp2451.pdf; NVIDIA spectrum-sn2000 page]:

| Series | Models | ASIC generation | Positioning |
|---|---|---|---|
| SN2201 | 48× 1GbE RJ45 + 4× 100G QSFP28, 448 Gb/s | Spectrum (1st gen) | OOB management / 1G ToR; Intel x86 dual-core, 8 GB ECC DDR4, 20 GB SSD [official — FS.com hardware manual mirror] |
| SN2000 | SN2010 (18× 25G + 4× 100G), SN2100 (16× 100G), SN2410 (48× 25G + 8× 100G), SN2700 (32× 100G, 3.2 Tb/s) | Spectrum-1/2 | 10/25/100G ToR, storage/HCI |
| SN3000 | SN3420 (48× 25G + 12× 100G), SN3510 (48× 50G + 6× 400G), SN3700/SN3700C (32× 200G / 64× 100G) | Spectrum-2 | 100/200G ToR/aggregation |
| SN4000 | **SN4700** (32× 400G QSFP-DD, 1U, 25.6 Tb/s, 8.4 Bpps) [official]; **SN4600** (64× 200G, 2U), SN4600C (64× 100G, 2U); **SN4800 modular chassis** (MSN4800-WS4, 4U, 4 PSUs, line cards MSN4800-C16 16× QSFP28, x86 mgmt cards) | Spectrum-3 | 400G spine/super-spine |

- SN4700 sells with Cumulus Linux, ONYX, or ONIE load options (MSN4700-WS2RC = Cumulus, C2P airflow) [secondary — advancedhpc.com]. All SN4000 series: 64 MB shared packet buffer, single-pass VXLAN, 512B deep DPI, INT-ready [official — SN4000 datasheet via etilize mirror].
- Positioning note: SN2201's dual role (OOB management switch + 1G ToR) explains its presence in both DC management and campus edge designs [independent].

### E. NVIDIA UFM / NetQ / telemetry 2026 (new detail)

- Base timeline cites **UFM Enterprise 6.26.1** (2026 Rel2A, Sep 2026). NVIDIA docs also show active UFM Enterprise manuals for **6.11.3, 6.17.5, 6.19.0, 6.22.2, 6.23.1** — i.e., a rapid 6.x cadence through 2026 [official — docs.nvidia.com]. UFM Enterprise 6.26 feature deltas were not extractable from the search results [gap].
- **UFM gNMI-Telemetry plugin 1.3.8-5**: UFM can stream events in **JSON format through gNMI** with optional device information (`include_dev_details_in_events`), requires UFM ≥ 6.23.1; 1.3.8-3 optimized dynamic XCSET group loading (faster startup, fetches only configured endpoints) [official — NVIDIA UFM Enterprise docs].
- Feature scope (UFM, stable across 6.x): fabric dashboard with congestion detection, real-time health/performance monitoring, fabric segmentation/isolation, QoS, routing optimization, switch auto-provisioning, fabric validation tests, HA, client-certificate auth [official].
- UFM **licensing model** (per-switch/per-port) was not located in 2026 sources [gap].

### F. Independent and vendor AI-fabric benchmark data, 2026 (new)

- **Supermicro + NVIDIA benchmark** (NVIDIA Technical Blog, 2024, still the reference vendor benchmark): Spectrum-X vs traditional Ethernet — **4.6× higher effective RDMA bisection bandwidth, 4.5× lower latency**; NCCL all-to-all/all-reduce gains with near-identical performance in noisy vs non-noisy AI-cloud scenarios (traditional Ethernet varied up to 20% run-to-run) [vendor-reported].
- **NVIDIA claims, 2026**: Spectrum-X delivers up to **1.6× network throughput** vs conventional Ethernet; job-completion metrics **1.3–1.6×**; plane-failover **1.08 s → ~2.68 ms** (>99% reduction); ~**90% bandwidth sustainability** in 8-plane configs with one plane down [vendor-reported — cryptobriefing.com, Sep 2026].
- **Hot Chips, Aug 2026 — "Scale-In"**: NVIDIA framed BlueField-4 + DOCA north-south offload as the "fifth pillar" of AI networking; revealed a multiplane topology scaling to **>512,000 endpoints in two tiers** with a path to million-GPU clusters; silicon-photonics support for ultra-large deployments targeted **H2 2026** [secondary — cryptobriefing.com].
- **Spectrum-X1600 (CPO) claims**: 5× lower power per port vs pluggable DSP optics (72 MW → ~21 MW networking power in a 400K-GPU DC); **~0.5 µs consistent end-to-end latency** by collapsing the electrical path to millimeters [vendor-reported — syndicated GTC coverage; treat as vendor claims, not independent measurements].
- **Independent comparative analysis** (gpusmith.com, Jul 2026): InfiniBand retains a per-hop latency edge "measured in the hundreds of nanoseconds"; Spectrum-X leads on **demonstrated single-fabric GPU count under one commercial platform**; Ultra Ethernet leads on long-term vendor optionality but is "**the least battle-tested of the four** as of July 2026"; buyers should weigh in-house IB vs Ethernet admin expertise and multitenancy (IB's subnet-manager model assumes single-tenant HPC) [independent].
- **Etherlink vs Spectrum-X** (community wiki, Sep 2026): Arista = merchant silicon (TH5/Jericho 3-AI) + 3rd-party NICs + RoCE DCQCN today / UEC NSCC-RCCC later; NVIDIA = in-house Spectrum-4 + BlueField/ConnectX SuperNIC + closed-loop TCC congestion control; Arista = founding UEC member, NVIDIA joined ~2024; Etherlink framed as "a bridge today (RoCEv2 + PFC + DLB) with UEC plumbing being added" [secondary — llm-systems-wiki; community analysis, unverified claims].

