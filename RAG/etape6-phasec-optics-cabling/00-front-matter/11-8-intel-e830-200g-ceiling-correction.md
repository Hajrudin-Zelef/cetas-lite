---
id: etape6-phasec-optics-cabling/00-front-matter/11-8-intel-e830-200g-ceiling-correction
title: "11.8 Intel (E830 = 200G ceiling — correction)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "Intel", "Microsoft", "Nvidia", "United States"]
dates: ["2023-01-09", "2025-06-11", "2025-07", "2025-09-05", "2026-01-28", "2026-06", "2026-07-16", "2026-09", "2026-09-22"]
keywords: ["intel", "acquisition", "cost", "ethernet", "full-duplex", "latency", "lpo", "nvidia", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1727, 1773]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 796fde6913133d390873abf8cf629267cab66f1a3d6949096c7b3fd7bd721dfa
---

# 11.8 Intel (E830 = 200G ceiling — correction)

### 11.8 Intel (E830 = 200G ceiling — correction)

- **Intel E830 is not an 800G controller; the family ceiling found is up to 200GbE** [official].
- Intel's E830 feature support matrix: initial public release July 2025, revisions through June 2026 [official].
- Reported configurations: 1×200GbE, 2×100/50/25/10GbE, 8×25/10GbE [official].
- Launch caveat: initial launch hardware included lower-rate models; 200GbE configurations followed [independent].
- A secondary HPE-channel article cites approximately US$7,000 for a four-port 25GbE E830 option; this is an integrated server option price, not comparable with bare 800G NIC retail prices [secondary].

### 11.9 Marvell and Fungible

- Searches on September 22, 2026 did **not** uncover a reliable public Marvell Octeon or FastLinQ 800G NIC/DPU product; search results mainly covered Marvell's 800G optical modules, 1.6T coherent pluggables, DSPs, and electro-optics — not host NICs [independent]. Do not present Marvell as an 800G NIC/DPU vendor without a verified product source [independent].
- Microsoft announced the acquisition of Fungible on January 9, 2023; Fungible's team joined Microsoft's data-center infrastructure engineering organization [vendor-reported]. A reported US$190 million purchase price was not officially disclosed — label as secondary [secondary]. By September 2026, Fungible should not be presented as an independent merchant 800G NIC/DPU supplier [independent]. No Fungible-branded 800G product was found [independent].

### 11.10 Host bus and form factor comparison

- PCIe 6.0: 64 GT/s per lane, up to 256 GB/s for x16 in the bidirectional convention PCI-SIG uses [official].
- PCIe 5.0 provides half the aggregate of 6.0 [official].
- PCIe 6.0 introduces PAM4 signaling, low-latency FEC, and flit-based encoding [official].
- Practical one-direction estimates from secondary engineering tables: Gen5 x16 ≈ 504 Gb/s effective; Gen6 x16 ≈ 1 Tb/s effective [secondary].
- An 800 Gb/s full-duplex NIC needs ~100 GB/s per direction; Gen5 x16 (~64 GB/s per direction) cannot sustain it, so **Gen6 x16 is required** (or multi-link/Socket Direct architectures) [independent].
- Intel's summary: Gen6 offers twice the data transfer rate of Gen5 and quadruple that of Gen4 [vendor-reported].
- 800G NICs therefore pair with PCIe Gen6 x16 host interfaces (ConnectX-8, Thor Ultra, reported Vulcano), while 400G AI NICs commonly use Gen5 x16 (Thor 2, Pollara) [independent].
- Single-port 800G NICs commonly use **OSFP/OSFP112** (8 electrical lanes × 100–112G PAM4) [independent].
- Dual-port 2×400G cards commonly use **QSFP112** (4 lanes × 112G per port) [independent].
- QSFP112 (4-lane, 400G) and QSFP-DD (double-density 8-lane, 800G) are not the same connector class [independent].
- OSFP offers larger thermal and power headroom and is prominent in 800G AI switches and NICs [independent].
- CEM PCIe x16 and OCP 3.0 are the common 800G AI NIC card forms in the products reviewed [official].

### 11.11 Cabling media and connectors

- Practical hierarchy: passive DAC (shortest, lowest power/cost, best in-rack where channel loss permits), AEC (retimed copper for longer rack/adjacent-rack reach, extra power/electronics), AOC (fixed active optical cable, lighter and longer but less field-serviceable), pluggable optics (longest reach, modular replacement, higher power/cost) [independent].
- Broadcom claims BCM57608 can drive passive copper to 5 m [vendor-reported].
- Broadcom's OFC 2026 portfolio includes 200G/lane AECs to 6 m [vendor-reported].
- NVIDIA Spectrum SN5600 supports passive DAC, AEC, and optical transceivers through OSFP ports [vendor-reported].
- Thor Ultra claims support for DAC, AEC, optical transceivers, and linear-drive (LPO) optics [vendor-reported].

### 11.12 Protocols: Ethernet vs InfiniBand vs UEC

- RoCEv2 runs RDMA over Ethernet/IP and typically relies on ECN/PFC plus vendor congestion systems for losslessness [independent].
- InfiniBand uses a native credit-based flow control and mature collectives (e.g., SHARP in-network aggregation) with a long HPC pedigree [independent].
- The Ultra Ethernet Consortium released **Specification 1.0 on June 11, 2025**: 560+ pages covering NICs, switches, optics, cables, transport, congestion control, security, and interoperability [vendor-reported].
- UEC 1.0 coverage: hpcwire.com [independent]; datacenterdynamics.com [independent].
- A secondary GitHub synthesis (needs primary verification) lists UEC 1.0.1 (September 5, 2025), 1.0.2 (January 28, 2026), and 1.0.3 (July 16, 2026, adding a 200 Gb/s-per-lane PHY and link/congestion fixes) [secondary].
- UET (Ultra Ethernet Transport) is a new transport design — packet spraying/multipath, out-of-order handling, selective retry, modern congestion mechanisms — not merely rebranded RoCEv2 [independent].
- Distinguish four stages: existing RoCEv2 deployments, UEC spec availability, sampling UEC-capable silicon (Thor Ultra), and broad production deployment, which gathered sources do not demonstrate [independent].
- Do not make unsupported categorical claims such as "InfiniBand is always faster" or "Ethernet has matched InfiniBand"; results depend on workload, topology, scale, and implementation [independent].

