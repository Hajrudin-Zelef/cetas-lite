---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-10-oem-integration-and-form-factors
title: "Wave 10 — OEM integration and form factors"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Broadcom", "Intel", "Meta", "Microsoft", "Nvidia"]
dates: ["2023-07", "2023-07-19", "2025-06-11", "2025-11", "2026-06", "2026-09-15", "2026-09-22"]
keywords: ["amd", "ethernet", "gpu", "gpus", "intel", "latency", "lean", "nvidia", "training"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [361, 426]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 0894dfe833d45e7fdd939c77d3f571ce9f4882c40cf705884166645af513abfe
---

# Wave 10 — OEM integration and form factors

## Wave 10 — OEM integration and form factors

### 10.1 Official NVIDIA BF-3 firmware SKU table (extract)

The official BlueField-3 firmware compatibility page lists [official] [source: https://docs.nvidia.com/networking/display/bluefield3firmwarev32432402lts/firmware+compatible+products ]:

| SKU | PSID | Description |
|---|---|---|
| 900-9D3B6-00CN-AB0 | MT_0000000883 | BF-3 B3240 P-Series dual-slot FHHL DPU; 400GbE/NDR IB; dual QSFP112; PCIe Gen5 x16 (x16 extension option); 16 Arm cores; 32 GB DDR; BMC; crypto enabled |
| 900-9D3B6-00CV-AAH | MT_0000000884 | BF-3 B3220 P-Series FHHL DPU; 200GbE/NDR200 IB; dual QSFP112; PCIe Gen5 x16; 16 Arm cores; 32 GB DDR; BMC; crypto enabled |
| 900-9D3B6-00SN-AB0 | MT_0000000964 | B3240 P-Series; same as above; crypto disabled |
| 900-9D3B6-00SV-AA0 | MT_0000000965 | B3220 P-Series; same as above; crypto disabled |
| 900-9D3B4-00CC-EA0 | MT_0000000966 | BF-3 B3210L E-Series FHHL SuperNIC; 100GbE/HDR100 IB; dual QSFP112; PCIe Gen5 x16; 8 Arm cores; 16 GB DDR; BMC; crypto enabled |
| 900-9D3B4-00SC-EA0 | MT_0000000967 | B3210L E-Series; same; crypto disabled |
| 900-9D3B4-00EN-EA0 | MT_0000001010 | BF-3 B3140L E-Series FHHL SuperNIC; 400GbE/NDR IB; single QSFP112; PCIe Gen5 x16; 8 Arm cores; 16 GB DDR; BMC; crypto enabled |
| 900-9D3B4-00PN-EA0 | MT_0000001011 | B3140L E-Series; same; crypto disabled |

- Note: E-series = "SuperNIC" branding (lean Arm complex, 8 cores/16 GB); P-series = full DPU (16 cores/32 GB). Both carry an integrated BMC (separate from host BMC) [official].
- Earlier DOCA table variants also listed `900-9D3B6-00CV-AA0` (B3220 crypto-disabled) and `900-9D3B4-00CC-EA0` described as B3210L E-series 100G (same as above) [official].

### 10.2 Lenovo ThinkSystem BF-3 adapters

Lenovo Press product guide (lp1809) documents [official/vendor-reported] [source: https://lenovopress.lenovo.com/lp1809-thinksystem-nvidia-bluefield-3-qsfp112-adapters ]:

| Lenovo part | Feature | Description | NVIDIA model | NVIDIA P/N |
|---|---|---|---|---|
| 4XC7A93809 | C0Q4 | ThinkSystem NVIDIA BlueField-3 B3140H VPI QSFP112 1P 400G PCIe Gen5 x16 Adapter | B3140H SuperNIC | 900-9D3D4-00EN-HA0 |
| 4XC7A87752 | BVBG | BlueField-3 B3220 VPI QSFP112 2P 200G PCIe Gen5 x16 Adapter (gold-plated power connector) | B3220 DPU | 900-9D3B6-00CV-AA0 |
| 4XC7B08942 | CCVX | B3220 2P 200G PCIe Gen5 x16 with tin-plating connector | B3220 DPU | — |
| 4XC7A96568 | C4GD | BlueField-3 B3240 2P 400G PCIe Gen5 x16 crypto-enabled (generic FW, tin-plated, dual-slot) | B3240 DPU | 900-9D3B6-00CN-AB0 |

- Auxiliary power cable kits required for B3220/B3240: 4X97B02426 (V3 2U), 4X97A91527 (SR675 V3), 4X97B04861 (SR650/a V4) — DPUs above PCIe-slot 75 W envelope draw auxiliary power [official].
- Dell-OEM B3220 (PC-Canada, Sept 2026 snapshot): ~$6,982–$7,229 vs MSRP $11,579.78 [secondary].

### 10.3 HPE and power

- HPE sells a **ProLiant DL385 Gen11 DPU Power Cable Kit for 3x NVIDIA BlueField-3** (P71949-B21), showing multi-DPU-per-server configurations in AMD EPYC Genoa/Bergamo platforms [official] [source: https://www.hpe.com/psnow/generateDDS/HPE%20ProLiant%20DL385%20Gen11%20DPU%20Power%20Cable%20Kit%20for%203x%20NVIDIA%20BlueField-3%20data%20sheet-PSN1014914516DKEN.pdf?oid=1014914516&cc=DK&lc=EN&softroll=false&print=&section=&prelaunchSection=&softrollSection=&deepLink=&utm_source=&utm_medium=&utm_campaign=&utm_content=&utm_term= ].

### 10.4 Software validation snapshot (AI-factory reference HCL)

- A third-party AI-factory infra-controller HCL (updated ~Sept 2026) pins: BlueField-2 → DOCA 3.2.0; BlueField-3 → DOCA/HBN 3.2.2 — "tested software pins, not necessarily latest" [secondary] [source: https://github.com/dsx-ai-factory/infra-controller/blob/HEAD/docs/hcl.md ].
- Compsource: BF-3 B3210E 900-9D3B6-00CC-EA0 new at $4,040.60, out of stock (listing last updated 2026-09-15) [secondary] [source: https://www.compsource.com/buy/9009D3B600CCEA0/Mellanox-4134/Nvidia-Bluefield3-B3210e-ESeries-Fhhl-Dpu-9009D3B600CCEA0-9009D3B600CCEA0/ ].

## Wave 11 — Ultra Ethernet 1.0 deep-dive (transport context)

### 11.1 Specification facts

- UEC formed **July 2023** (announced July 19, 2023) by AMD, Arista, Broadcom, Cisco, Eviden, HPE, Intel, Meta, Microsoft under the Linux Foundation's Joint Development Foundation; NVIDIA joined ~Aug 2024; >100 member companies / 1,500 participants as of 2026 [secondary].
- **UEC Specification 1.0 released June 11, 2025** — 560+ page document defining a vertically integrated networking stack from physical layer to application API [secondary] [source: https://ultraethernet.org/ultra-ethernet-consortium-uec-launches-specification-1-0-transforming-ethernet-for-ai-and-hpc-at-scale/ — via https://github.com/rishijain905/localmodelresearch/blob/HEAD/topics/training-infrastructure/high-speed-fabrics/Ultra-Ethernet.md ].
- Errata/maintenance releases: 1.0.1 (Sep 5, 2025 — RCCC source algorithm correction); 1.0.2 (Jan 2026 — CMS congestion-control corrections); **1.0.3 (Jul 16, 2026, current)** — 200 Gb/s-per-lane PHY, LLR/CBFC race fixes, MP_Range<128, CtlOS protection [secondary] [source: https://github.com/abuabdurahman82/llm-systems-wiki/blob/HEAD/AI-Factory-Networking/30-ultra-ethernet-consortium.md ].

### 11.2 Architecture blocks

- **UET (Ultra Ethernet Transport)**: packet format/delivery model replacing RoCEv2's BTH layer.
- **Retransmit**: SACK-based selective retransmission — reliability **without lossless fabric** (key difference vs RoCEv2); unordered delivery is first-class.
- **Congestion control**: ACK-based, end-to-end (sender→receiver), not PFC pause-based; spec modes: NSCC (network-signal: ECN + latency, source congestion window), RCCC (receiver-credit CC: destination credits control all sources), TFC (transport flow control to avoid destination buffer overflow; not CC by itself) [secondary] [source: https://a.ifif.workers.dev:443/https/ultraethernet.org/wp-content/uploads/sites/20/2025/06/UE-Specification-6.11.25.pdf ].
- **Multipath**: native packet spraying across paths; receiver reassembles out-of-order [secondary].
- Link layer: LLR (link retry), CBFC (credit-based flow control, optional); profiles AI-Base / AI-Full / HPC; in-network (INC) collectives [secondary].
- Common misconception corrections: spec terms are **NSCC + RCCC** (not "CCv1/CCv2/TACC"); there was **no November 2025 release** (1.0.1 only); NVIDIA ships **no UET NIC** (pushes Spectrum-X) [secondary] [source: https://github.com/abuabdurahman82/llm-systems-wiki/blob/HEAD/AI-Factory-Networking/30-ultra-ethernet-consortium.md ].

### 11.3 Ecosystem readiness

- Broadcom shipped first UE-compatible Tomahawk switch (~250 ns latency); Arista announced Etherlink portfolio support [secondary] [source: https://medium.com/teradata-labs/the-ultra-ethernet-endgame-cec7242a243b ].
- VIAVI launched (June 2026) the first validation solution for Ultra Ethernet Transport, emulating AI workload traffic without GPUs; validates ordered/unordered delivery, packet trimming, CC, dynamic multipathing [secondary] [source: https://cryptobriefing.com/viavi-ultra-ethernet-validation-solution/ ].
- Claimed positioning: ~30% performance loss of RoCEv2 vs InfiniBand at 10,000-GPU scale motivating UEC; UEC claims InfiniBand-class tail latency on open Ethernet — **vendor ecosystem claims, not independently benchmarked at scale as of 2026-09-22** [unverified].

