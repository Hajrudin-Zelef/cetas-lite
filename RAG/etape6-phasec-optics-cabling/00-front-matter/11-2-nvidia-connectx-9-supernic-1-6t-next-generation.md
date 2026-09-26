---
id: etape6-phasec-optics-cabling/00-front-matter/11-2-nvidia-connectx-9-supernic-1-6t-next-generation
title: "11.2 NVIDIA ConnectX-9 SuperNIC (1.6T, next generation)"
domain: front-matter
role: reference
task: actor-profile
actors: ["Nvidia", "United States"]
dates: ["2025-05-20", "2025-06", "2025-06-10", "2025-10", "2026-09-22"]
keywords: ["nvidia", "compute", "ethernet", "gpu", "pricing", "rubin", "vera rubin"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1646, 1672]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 62654fa07dad7561aecaa9d6581e18c5e516c4c125abcaeaf8caaa651db7726b
---

# 11.2 NVIDIA ConnectX-9 SuperNIC (1.6T, next generation)

- NVIDIA ConnectX-8 is an 800 Gb/s network adapter available in InfiniBand and Ethernet/VPI variants [official].
- ConnectX-8 uses a PCIe Gen6 x16 host interface [official].
- Firmware release notes for ConnectX-8 v40.45.1200 were dated May 20, 2025, indicating pre-GA hardware in the field by mid-2025 [official]. A June 2025 firmware revision (40.45.1202) followed, documentation page updated June 10, 2025 [official].
- Form factors in NVIDIA documentation: CEM PCIe x16, OCP 3.0, and Socket Direct auxiliary-card configurations [official].
- Connector variants: single-port OSFP224 and dual-port 112G-class connector variants [official].
- Documented variant C8180: 800 Gb/s XDR InfiniBand default mode, also configurable to 2×400GbE, single OSFP, PCIe 6.0 x16, multi-host and Socket Direct capable; part number 900-9X85E-00EX-MC0 [official].
- Documented variant C8240: dual QSFP112, aggregate 800G as 2×400G [secondary].
- ConnectX-8 supports in-network computing (SHARP), MPI_Alltoall acceleration, QoS, and congestion-control functions per NVIDIA firmware release notes [official].
- ConnectX-8 supports RoCEv2, InfiniBand (XDR/NIC generations), and GPUDirect RDMA across compute fabrics [official].
- Official firmware release notes: docs.nvidia.com ConnectX-8 SuperNIC firmware release notes v40.45.1200 [official]. Official OCP user manual: docs.nvidia.com ConnectX-8 SuperNIC for OCP 3.0 user manual [official].
- As of September 22, 2026, ConnectX-8 cards are listed as in-stock items by third-party distributors, consistent with commercial availability through channels; NVIDIA has not published a single global GA date in the documents reviewed [secondary].
- FS.com listed, as of September 22, 2026: ConnectX-8 VPI one-port 800G OSFP PCIe 6.0 x16 at US$2,519; ConnectX-8 VPI dual-port 400G QSFP112 at US$2,519 [secondary] (fs.com NVIDIA Ethernet NICs category).
- NADDOD listed, as of September 22, 2026: C8240 dual QSFP112 at US$1,779 and C8180 single OSFP XDR 800G at US$1,779 [secondary].
- A private resale listing observed September 22, 2026 asked US$2,000 each for new single-port 800G cards; anecdotal, not representative pricing [secondary] (forums.servethehome.com).
- No official NVIDIA list price for ConnectX-8 was found in public documentation; all prices above are channel/retail [unverified].

### 11.2 NVIDIA ConnectX-9 SuperNIC (1.6T, next generation)

- NVIDIA introduced the ConnectX-9 SuperNIC at GTC DC in Washington, D.C. in October 2025, as part of the Vera Rubin platform [independent].
- ConnectX-9 is positioned as a 1.6 Tb/s-per-GPU next-generation network interface card with advanced RDMA capabilities [vendor-reported].
- ConnectX-9 supports PCIe Gen6 [vendor-reported].
- NVIDIA claims ConnectX-9 enables four-times-larger clusters than BlueField-3-era hardware; treat as vendor projection [vendor-reported].
- ConnectX-9 incorporates the BlueField-4 architecture path: BlueField-4 combines Arm-based CPUs with the ConnectX-9 SuperNIC [vendor-reported].
- Launch coverage: networkworld.com (NVIDIA looks to power AI factory networks) [independent]; sdxcentral.com (NVIDIA next-gen DPU) [independent].
- A four-way ConnectX-9 IO card designed for the NVIDIA Vera Rubin NVL144 system was described in secondary coverage; treat integration details as vendor-reported [secondary].
- ConnectX-9 ships in the Rubin timeframe (2026), which is also NVIDIA's stated 1.6T adoption window; no standalone retail price was found [unverified].

