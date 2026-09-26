---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/8-4-subsystem-local-memory-computational-programs-2-1-newcom
title: "8.4 Subsystem Local Memory & Computational Programs (2.1 newcomers)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [188, 204]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: a8953c1f3449345e31fc17ed58ced1d83185804d16159c45ae7fe8f169eeecfa
---

# 8.4 Subsystem Local Memory & Computational Programs (2.1 newcomers)

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

