---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/7-nvme-specification-family-2-0-2-4-august-2026
title: "7. NVMe specification family: 2.0 → 2.4 (August 2026)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2025-08-05", "2026-08", "2026-08-04", "2026-09-22"]
keywords: ["datacenter", "memory", "nand", "research"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [141, 204]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 1c13dbe26243e092b74aad295ecb1b16ffe95d24184035b1756253f5104b2e5b
---

# 7. NVMe specification family: 2.0 → 2.4 (August 2026)

## 7. NVMe specification family: 2.0 → 2.4 (August 2026)

### 7.1 The spec set, not a single spec

- NVMe grew "from a single PCIe SSD specification" into "nearly a dozen specifications, including multiple command sets" covering all major transports [secondary](https://www.guru3d.com/story/nvm-express-announces-release-of-nvme-21-specifications/).
- The family: Base specification + I/O Command Set specs (NVM, ZNS, Key Value, Subsystem Local Memory, Computational Programs) + Transport specs (PCIe, RDMA, TCP, Fibre Channel) + NVMe Management Interface (NVMe-MI) + NVMe Boot [official](https://nvmexpress.org/SPECIFICATIONS/).
- The FC transport spec and the original NVMe-oF spec are carried as historical references; current transports are PCIe/RDMA/TCP/FC individual specs [official](https://nvmexpress.org/SPECIFICATIONS/).

### 7.2 Revision timeline

| Release | Date | Base rev | Headline features |
|---|---|---|---|
| NVMe 2.0 family | 2021-06 | 2.0 | Spec refactor (Base + transports + command sets), ZNS/KV as command sets, NVMe-MI 1.1 |
| NVMe 2.1 | 2024-08 | 2.1 | FDP (host-directed placement), PCIe controller live migration, computational offload, NVMe Boot spec (new), Subsystem Local Memory (new), Computational Programs (new), NVMe-oF zoning |
| NVMe 2.3 | 2025-08-05 | 2.3 | Rapid Path Failure Recovery, Power Limit Config, Configurable Device Personality, sustainability; NVMe-MI 2.1; Transports PCIe 1.3 / RDMA 1.2 / TCP 1.2 |
| NVMe 2.4 | 2026-08-04 | 2.4 | Post-Quantum Cryptography, PCIe Exported NVM Subsystem Migration, Voltage Monitoring; NVMe-MI 2.2; Transports PCIe 1.4 / RDMA 1.3 / TCP 1.3 |

- NVMe 2.1 details: "three new specifications" (NVMe Boot, Subsystem Local Memory command set, Computational Programs command set) and "eight updated"; capabilities include live migration of PCIe NVMe controllers between subsystems, host-directed data placement "backwards compatible with previous NVMe specifications", offloading host processing to devices, network boot for NVMe-oF, and NVMe-oF zoning [secondary](https://www.guru3d.com/story/nvm-express-announces-release-of-nvme-21-specifications/)[secondary](https://www.techpowerup.com/325310/nvm-express-releases-nvme-2-1-specifications).
- NVMe 2.3 (Aug 5, 2025): all 11 specs; Base 2.3, NVM CS 1.2, ZNS CS 1.4, KV CS 1.3, SLM CS 1.2, Computational Programs CS 1.2; PCIe Transport 1.3, RDMA 1.2, TCP 1.2; NVMe-MI 2.1; NVMe Boot 1.3; features: Rapid Path Failure Recovery (alternate-channel communication on controller loss), Power Limit Config (host control of device max power), Configurable Device Personality, sustainability enhancements [secondary](https://markets.financialcontent.com/sweetwaterreporter/article/bizwire-2025-8-5-nvm-express-publishes-set-of-nvme-specifications-enabling-new-capabilities-for-ai-cloud-enterprise-and-client-storage).
- NVMe 2.4 (Aug 4, 2026 — six weeks before cutoff): all 11 specs; Base 2.4, NVM CS 1.3, ZNS CS 1.5, KV CS 1.4, SLM CS 1.3, Computational Programs CS 1.3; PCIe Transport 1.4, RDMA 1.3, TCP 1.3; NVMe-MI 2.2; NVMe Boot 1.4; headline: Post-Quantum Cryptography, PCIe Exported NVM Subsystem Migration, Voltage Monitoring — "foundational enablers to remain secure in the post-quantum era" per NVM Express president Amber Huffman [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- "75 new authorized technical proposals" were already in progress at the 2.1 launch — the pipeline behind 2.3/2.4 [secondary](https://www.techpowerup.com/325310/nvm-express-releases-nvme-2-1-specifications).

### 7.3 FDP (Flexible Data Placement) — TP4146

- FDP is the NVMe 2.1 host-directed data placement mechanism: "a new method for placing logical blocks into the non-volatile storage in an effort to reduce Write Amplification Factor (WAF) by the SSD" [vendor-reported](https://download.semiconductor.samsung.com/resources/white-paper/FDP_Whitepaper_102423_Final.pdf).
- Spec mechanics (optional feature): new fields in the NVM Command Set Identify Namespace data structure; Namespace Management extended to create FDP-enabled namespaces; directives added to Write Zeroes and Write Uncorrectable; Reclaim Unit Handle Status Descriptor in I/O Management Receive; Media Reallocation Event Type in the FDP Events log page [official](https://nvmexpress.org/wp-content/uploads/NVM-Express-Revision-2.1-Changes-08_07.pdf).
- FDP vs the older mechanisms: **Streams** give independent write paths (host picks a stream ID per write directive); **ZNS** splits namespaces into sequential-write zones tied to physical NAND (host placement via LBA); **FDP** is designed to be "backwards compatible with previous NVMe specifications" and simpler to integrate [vendor-reported](https://download.semiconductor.samsung.com/resources/white-paper/FDP_Whitepaper_102423_Final.pdf).
- Host software stack support for FDP was still being built out at whitepaper time (Oct 2023); Linux ecosystem enablement is the gating factor for datacenter use [vendor-reported](https://download.semiconductor.samsung.com/resources/white-paper/FDP_Whitepaper_102423_Final.pdf) [unverified as of 2026-09-22 — no 2026-dated FDP drive announcements found in this research].

## 8. NVMe I/O command sets in 2026

### 8.1 NVM Command Set (the default)

- The NVM command set is what "NVMe SSD" means by default: read/write/compare/write-zeroes/dataset-management on logical block namespaces [independent].
- Current rev in the 2.4 family: NVM Command Set 1.3 [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- NVMe 2.1 NVM-CS changes beyond FDP: Key Per I/O (TP4055 — 16-bit key tag field carved out of Dword 13, CETYPE/CEV on I/O commands), Performance Characteristics Reporting (TP4077), dispersed namespaces clarification (TP4034a) [official](https://nvmexpress.org/wp-content/uploads/NVM-Express-Revision-2.1-Changes-09_05_24.pdf).

### 8.2 ZNS (Zoned Namespaces)

- ZNS divides each namespace into equally-sized zones, each mapped to physical NAND; "A zone is required to be written sequentially" and the host performs placement via the LBA in the write command [vendor-reported](https://download.semiconductor.samsung.com/resources/white-paper/FDP_Whitepaper_102423_Final.pdf).
- Samsung PM1731a (enterprise ZNS SSD): WAF "close to one, a major improvement over typical server SSD values between three and four", "last up to four times longer than conventional NVMe SSDs", and full usable capacity by "eliminating the need for overprovisioning" [vendor-reported](https://www.enterprisestorageforum.com/news/samsung-zns-ssd/).
- Samsung + Western Digital signed an MOU (2022) to unify zoned-storage standards — ZNS SSDs plus SMR HDDs — under "data placement, processing, and fabrics (D2PF)", explicitly to fix fragmentation from vendors' differing implementations [secondary](https://www.eetimes.com/samsung-western-digital-unite-around-zoned-storage).
- ZNS software adoption "makes it hard": the MOU's purpose was standardization so "end-users [have] confidence that these emerging technologies will have support from multiple device vendors and a vertically integrated hardware and software ecosystem" [secondary](https://www.eetimes.com/samsung-western-digital-unite-around-zoned-storage).
- ZNS in 2026: current ZNS Command Set rev is 1.5 (NVMe 2.4 family); real-world deployment remains niche vs FDP's simpler integration story — ZNS suits log-structured workloads (LSM key-value, journaling) where sequential zones match the I/O pattern [official][independent].

### 8.3 Key Value (KV) Command Set

- KV replaces block LBA addressing with Store/Retrieve/Delete/Exist/List on key-value pairs; the KV Command Set 1.0a defines Store/Retrieve/Delete/Exist/List as mandatory (M) for I/O controllers supporting it [official](https://nvmexpress.org/wp-content/uploads/NVMe-Key-Value-Command-Set-Specification-1.0a-2021.07.26-Ratified.pdf).
- Current rev: KV Command Set 1.4 in the NVMe 2.4 family [official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- Ecosystem signal: SPDK "added support for the NVMe Key Value (KV) command set" with a new public header exposing spdk_nvme_kv_store/retrieve/delete/exist/list and KV-aware identify data updated to KV Command Set Specification 1.3 [secondary](https://github.com/spdk/spdk/blob/HEAD/CHANGELOG.md).
- Target workloads: KV stores, metadata, AI feature stores — eliminating the block translation layer [independent].

### 8.4 Subsystem Local Memory & Computational Programs (2.1 newcomers)

- Subsystem Local Memory (SLM) Command Set and Computational Programs Command Set were both introduced as new specifications in NVMe 2.1; current revs are 1.3 each in the 2.4 family [secondary](https://www.techpowerup.com/325310/nvm-express-releases-nvme-2-1-specifications)[official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- They formalize "host processing offloading" — "support for offloading some host processing to NVMe storage devices" — the standards track behind computational storage [secondary](https://www.guru3d.com/story/nvm-express-announces-release-of-nvme-21-specifications/).
- Industry context: "computational storage trims data movement by processing information where it sits" and is listed alongside ZNS and NVMe-oF as one of the three technologies "central to how servers are evolving in 2026" [secondary](https://www.techtimes.com/articles/317663/20260603/how-new-storage-tech-like-zns-ssds-nvme-computational-storage-are-changing-servers.htm).

### 8.5 NVMe Boot & NVMe-MI

- NVMe Boot specification (new in 2.1, now rev 1.4) "introduces a boot mechanism for NVMe over Fabrics (NVMe-oF)" [secondary](https://www.guru3d.com/story/nvm-express-announces-release-of-nvme-21-specifications/)[official](https://www.businesswire.com/news/home/20260804315628/en/NVM-Express-Publishes-Set-of-NVMe-Specifications-Enhancing-Security-Manageability-and-Sustainability-for-AI-Cloud-Enterprise-and-Client-Storage).
- NVMe-MI (Management Interface) is at rev 2.2 in the 2.4 family; shipping drives already implement it in-band — KIOXIA CM9 is "NVMe-MI 1.2c" compliant [official][vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology).
- NVMe-MI carries enclosure/thermal/power management, including the Linux nvme-cli MI admin-command paths (e.g. ISH bit handling per NVMe-MI 2.1) [secondary](https://github.com/linux-nvme/nvme-cli/commit/5be98d6689ed478cdea17562b008189583824323).

