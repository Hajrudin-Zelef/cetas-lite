---
id: etape6-phasea-vendors-dc/00-front-matter/a-dell-z9964f-on-z9864f-on-software-support-closes-supplemen
title: "A. Dell Z9964F-ON / Z9864F-ON software support — closes Supplementary Pass #2 §W item 7"
domain: front-matter
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: ["2025-12-09", "2026-07", "2026-09-22", "2027-12", "2028-11"]
keywords: ["asic", "benchmark", "cpo", "distribution", "dsp", "ethernet", "full-duplex", "gpu", "intel", "latency", "nvidia", "optics"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1360, 1406]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 5bd1bd7cfeef53a9765f29dda72f42713e9541f65fc56f168e8f8216e40cf39d
---

# A. Dell Z9964F-ON / Z9864F-ON software support — closes Supplementary Pass #2 §W item 7

### A. Dell Z9964F-ON / Z9864F-ON software support — closes Supplementary Pass #2 §W item 7

- Dell's **SmartFabric OS10 Hardware Compatibility List (KB 000192674, current 2026)** supports the **400G column** on **Z9264F, Z9332F, Z9432F, Z9664F** (10.6.1.x, EoM November 2028; 10.6.0.x, EoM December 2027). **Neither Z9864F-ON (800G, Tomahawk-5) nor Z9964F-ON (1.6T, Tomahawk-6) appears in the OS10 compatibility list** [official — Dell KB 000192674, 2026].
- Dell's Secure Connect Gateway support matrix for Enterprise SONiC lists Z9264F/Z9332F/Z9432F/Z9664F (telemetry compatibility), with Z9964F-ON absent from that matrix as well — that matrix only tracks telemetry-collection compatibility, not NOS support itself [official — Dell Secure Connect Gateway 5.x support matrix].
- Finding: as of 2026-09-22, the **800G/1.6T Tomahawk-5/6 platforms are Enterprise SONiC-only** in Dell's documented matrix; OS10 coverage tops out at 400G (Z9664F-ON). Dell's published Enterprise SONiC spec sheet covers the Z-series (Z9332F/Z9264F/Z9432F/Z9664F/Z9864F) but the Z9964F-ON's exact NOS support document was not located — flagged [unverified].

### B. Dell Z-series street-pricing evidence (secondary market)

- **Z9664F-ON (64× 400G)**: new units remain **quote-only** from Dell (confirmed again on Dell Canada's data-center switch page, Sep 2026 crawl — "Shop Now" leads to quote/configurator) [official]. Secondary market: refurbished/tested Z9664F-ON listed at **$8,495** (expresscomputersystems.com; 1 in stock, RAF airflow, OS10) [secondary].
- Note: the same reseller listing describes the Z9664F-ON as "51.2 Tbps" — this is **full-duplex** capacity (64× 400G × 2); Dell's official 25.6 Tbps is the simplex switching figure. The discrepancy is labeling, not a spec conflict [secondary — flagged].
- **Z9864F-ON / Z9964F-ON**: no public list or street prices located; quote-only stands (base §1.7 / §7.1 unchanged) [gap].

### C. Enterprise SONiC bundles + ecosystem (new detail)

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

