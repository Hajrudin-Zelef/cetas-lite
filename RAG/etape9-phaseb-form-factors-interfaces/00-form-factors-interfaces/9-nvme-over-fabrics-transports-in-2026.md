---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/9-nvme-over-fabrics-transports-in-2026
title: "9. NVMe over Fabrics — transports in 2026"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: []
dates: ["2025-06-11", "2026-01"]
keywords: ["cost", "datacenter", "disaggregated", "dram", "ethernet", "gpus", "inference", "latency", "memory", "throughput", "training"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [205, 256]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 32564283c2a41403872de80eb0a69763319d1684022d9fc7cf2c74408587b02c
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

