---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/9-nvme-over-fabrics-transports-in-2026
title: "9. NVMe over Fabrics — transports in 2026"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: []
dates: ["2025-06-11", "2026-01"]
keywords: ["accelerator", "compute", "consumer", "cost", "datacenter", "disaggregated", "dram", "ethernet", "gpus", "inference", "latency", "memory"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [205, 276]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: bea41c258dc5f329386f22517292b415d36a02c6f590ba59e5ecaf39c80d0a55
---

# 9. NVMe over Fabrics — transports in 2026

## 9. NVMe over Fabrics — transports in 2026

### 9.1 The transport menu

| Transport | Spec rev (2.4 family) | Media | Latency profile | 2026 position |
|---|---|---|---|---|
| PCIe (local) | PCIe Transport 1.4 | PCIe lanes | Lowest | Default for DAS |
| RDMA (RoCE v2 / iWARP / IB) | RDMA Transport 1.3 | Ethernet / InfiniBand | Very low | Perf-sensitive fabrics |
| TCP | TCP Transport 1.3 | Any Ethernet | Low | Mainstream disaggregation |
| Fibre Channel | FC Transport (2.1-era) | FC fabric | Very low | Legacy SAN migration |

- "NVMe technology adoption continues to grow and has succeeded in unifying client, cloud, AI and enterprise storage around a common architecture" across all major transports [secondary](https://www.techpowerup.com/325310/nvm-express-releases-nvme-2-1-specifications).

### 9.2 NVMe/TCP — the mainstream disaggregation play

- NVMe/TCP runs on standard Ethernet: "NVMe-TCP works on any modern Ethernet network, without requiring specialized hardware or significant network redesign", making it "a logical upgrade for organisations that use iSCSI SAN storage" [secondary](https://www.computerweekly.com/feature/NVMe-over-Fabrics-How-NVMe-oF-revolutionises-shared-storage)[secondary](https://www.blueally.com/why-nvme-of-is-the-future-of-datacenter-connectivity/).
- The RDMA alternative is fragmented: "Choosing to run NVMe over RDMA requires committing to either RDMA over Converged Ethernet (RoCE) or its predecessor, Internet Wide-area RDMA Protocol (iWARP), as very few devices will handle both. RoCE requires converged Ethernet… that means I need to get the network team involved" — Howard Marks, DeepStorage [secondary](https://www.eetimes.com/nvme-over-tcp-will-take-time-to-eclipse-rdma/).
- Hyperscale driver: direct-attached flash is stranded at "30% to 40% utilization of the SSDs"; disaggregating via NVMe/TCP lets operators "right-size your storage footprint without sacrificing performance" [secondary](https://www.eetimes.com/nvme-over-tcp-will-take-time-to-eclipse-rdma/)[secondary](https://www.blueally.com/why-nvme-of-is-the-future-of-datacenter-connectivity/).
- Market proof: Lightbits Labs (self-described "inventor of the NVMe over TCP storage protocol") reported January 2026 "a 3X Y/Y increase in software purchases" as "enterprises standardize on software-defined, NVMe over TCP-based storage", with customers claiming "up to 5x greater hardware efficiency than legacy SDS, such as Ceph Storage" [vendor-reported](https://www.storagenewsletter.com/2026/01/16/lightbits-labs-delivers-record-growth-as-organizations-standardize-on-nvme-over-tcp-for-ai-ready-infrastructure/).
- AI angle: "NVMe-oF makes it possible to build high-throughput, low-latency storage backends that keep GPUs constantly busy" — disaggregated NVMe pools behind AI training/inference clusters [secondary](https://www.blueally.com/why-nvme-of-is-the-future-of-datacenter-connectivity/).
- FAQ reality check: "NVMe-oF can run over standard Ethernet using TCP, but low-latency RDMA-capable NICs and higher-speed links usually deliver the best results" [secondary](https://www.techtimes.com/articles/317663/20260603/how-new-storage-tech-like-zns-ssds-nvme-computational-storage-are-changing-servers.htm).

### 9.3 RDMA and Fibre Channel notes

- RoCE v2 needs a lossless-ish fabric (PFC + ECN end-to-end) — operational cost is the tax vs TCP [independent].
- FC-NVMe: Gen6 FC HBAs (16/32 Gb) can carry NVMe alongside SCSI; the FCIA pushes single adapters covering disk, SSD and NVMe [secondary](https://www.computerweekly.com/feature/NVMe-over-Fabrics-How-NVMe-oF-revolutionises-shared-storage).
- NVMe-oF zoning (new in 2.1) brings fabric-level access control to shared NVMe pools [secondary](https://www.guru3d.com/story/nvm-express-announces-release-of-nvme-21-specifications/).

## 10. Advanced NVMe features (hardware-relevant)

- **Dual-port NVMe**: KIOXIA CM9 supports dual-port in both 2.5" and E3.S — two PCIe paths to one drive for active/active HA without SAS [vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology).
- **Multipath / ANA**: NVMe native multipathing (Linux nvme-multipath) plus Asymmetric Namespace Access log pages let hosts survive path loss; NVMe 2.3's Rapid Path Failure Recovery adds "communication with the NVM subsystem through alternative channels" at spec level [independent][secondary](https://markets.financialcontent.com/sweetwaterreporter/article/bizwire-2025-8-5-nvm-express-publishes-set-of-nvme-specifications-enabling-new-capabilities-for-ai-cloud-enterprise-and-client-storage).
- **Controller Memory Buffer (CMB)**: on-controller memory host-mappable for admin queues/small I/O — reduces host DRAM traffic; a standard NVMe feature since 1.2, still relevant for DPU-adjacent designs [independent].
- **SR-IOV for NVMe**: single physical NVMe controller exposing virtual functions to VMs — the hardware path to vGPU-style storage sharing; complements virtio-blk/vhost in virtualized stacks [independent].
- **Live migration of PCIe NVMe controllers** between NVM subsystems (NVMe 2.1, TP4159): LBA Migration Queue + Controller Data Queue + Migration Send command track changed blocks during migration [official](https://nvmexpress.org/wp-content/uploads/NVM-Express-Revision-2.1-Changes-08_07.pdf).
- **Power Limit Config** (2.3): host-controlled device max power — "particularly important for older systems with limited power capabilities" [secondary](https://markets.financialcontent.com/sweetwaterreporter/article/bizwire-2025-8-5-nvm-express-publishes-set-of-nvme-specifications-enabling-new-capabilities-for-ai-cloud-enterprise-and-client-storage).
- **Voltage Monitoring** (2.4) and **Post-Quantum Cryptography** (2.4): telemetry and crypto-agility hardening for the next decade [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- **OCP Datacenter NVMe SSD spec** (CM9 is "OCP Datacenter NVMe SSD 2.5 specification-compliant"): hyperscale requirements on top of NVMe (SMBus/VPD, telemetry, firmware, form-factor behaviors) — the compliance badge that matters for OCP/DC-MHS platforms [vendor-reported](https://www.silicon.co.uk/press-release/kioxia-announces-first-enterprise-nvme-ssd-built-with-8th-generation-bics-flash-tlc-based-flash-memory-technology).

## 11. PCIe generations: the lane underneath everything

### 11.1 Generation table (PCI-SIG)

| Generation | Spec year | Signaling | Raw rate/lane | x16 bidir bandwidth | 2026 status |
|---|---|---|---|---|---|
| PCIe 3.0 | 2010 | NRZ 128b/130b | 8.0 GT/s | ~64 GB/s | Legacy |
| PCIe 4.0 | 2017 | NRZ 128b/130b | 16 GT/s | ~128 GB/s | Mainstream installed |
| PCIe 5.0 | 2019 | NRZ 128b/130b | 32 GT/s | ~256 GB/s | Current server standard |
| PCIe 6.0 | 2022 | PAM4 + FEC, Flit | 64 GT/s | ~256 GB/s bidir (x16) | Early silicon, not widespread |
| PCIe 7.0 | 2025-06-11 | PAM4, Flit | 128 GT/s | 512 GB/s bidir (x16) | Spec ratified; products ~2028 |
| PCIe 8.0 | 2028 (planned) | — | 256 GT/s | ~1 TB/s (x16) | Pathfinding |

- PCIe 6.0 was "introduced in 2022 but has not yet seen widespread implementation in consumer products"; 7.0 products are "unlikely… before 2028" — spec ratification runs ~3 years ahead of volume [secondary](https://themunicheye.com/pci-express-7-0-finalized-data-transfer-speeds-23001).
- PCIe 7.0 official facts: 128.0 GT/s raw, "up to 512.0 GB/s bi-directional bandwidth via a sixteen lane (x16) configuration", PAM4 signaling, Flit-based encoding, improved power efficiency, full backward compatibility [official](https://pcisig.com/faq?field_category_value%5B%5D=pci_express_7.0&keys=).
- PAM4 arrived with 6.0 ("allows each transfer to convey two bits of data using four voltage levels, necessitating more sophisticated controller designs"); 7.0 keeps PAM4 + Flit [secondary](https://themunicheye.com/pci-express-7-0-finalized-data-transfer-speeds-23001)[secondary](https://www.eetimes.com/pcie-7-0-keeps-pace-with-ai-demands/).
- PCIe 7.0 was released to members on June 11, 2025; PCI-SIG president Al Yanes: "It takes three years to develop a specification. It takes three years for innovation. It takes three years for feedback on the previous technology" [secondary](https://convergedigest.com/pci-sig-finalizes-pcie-7-0-specification-at-128-0-gt-s/)[secondary](https://www.eetimes.com/pcie-7-0-keeps-pace-with-ai-demands/).
- **Optical Aware Retimer ECN**: alongside 7.0, PCI-SIG updated PCIe 6.4 and 7.0 "to enable standardized PCIe operation over optical fiber using retimer-based solutions" — "a pivotal step toward mainstream adoption of optical links in data center and AI accelerator topologies" [secondary](https://convergedigest.com/pci-sig-finalizes-pcie-7-0-specification-at-128-0-gt-s/).
- PCIe underpins CXL ("PCIe serves as the foundation for Compute Express Link, enabling the connection of additional memory modules") — storage and memory expansion ride the same SerDes [secondary](https://themunicheye.com/pci-express-7-0-finalized-data-transfer-speeds-23001).

### 11.2 What Gen5 means for storage in 2026

- PCIe 5.0 x4 ≈ 16 GB/s raw per drive link — the reason flagship Gen5 SSDs (KIOXIA CM9: 14.8 GB/s seq read) finally saturate what Gen4 x4 (8 GB/s) could not [vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology)[independent].
- All 2026 server platforms in this research are PCIe Gen5 native: Dell 17G, HPE Gen12, Supermicro Petascale, MSI DC-MHS [secondary].
- Gen5 signal integrity forces retimers/redrivers on longer traces and cables: Icy Dock's PCIe 5.0 x8 MCIO adapter integrates a "PCIe Redriver" plus 100 MHz clock buffer "for stable enterprise timing" [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).
- Note: the datasheet URL above is a long vendor CDN link — cited verbatim as returned; treat link rot risk as [unverified] for long-term retrieval.

### 11.3 Bifurcation, switches, retimers

- PCIe bifurcation (splitting x16 into x4/x4/x4/x4 etc.) is how motherboards feed multiple M.2/U.2 NVMe drives from one slot; bifurcation is a **motherboard** capability — adapter cards like the Icy Dock MB409A5 "itself does not perform bifurcation" [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).
- PCIe switch adapters multiply NVMe density: HighPoint Rocket 1624A — PCIe Gen5 x16 host, dual MCIO 8i ports, "supports up to 16 NVMe SSDs" per adapter, configurable downstream 1x16 down to 16x1, hot-plug/hot-swap, direct P2P between GPUs/NICs/SSDs bypassing the CPU [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw)[vendor-reported](https://electronicsbuzz.in/highpoint-unveils-comprehensive-pcie-gen-5-mcio-expansion-ecosystem/).
- Native OS support (Linux/Windows, x86 and ARM) with no proprietary driver overhead is the selling point of switch-based NVMe expansion [vendor-reported](https://electronicsbuzz.in/highpoint-unveils-comprehensive-pcie-gen-5-mcio-expansion-ecosystem/).

