---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/15-3-supermicro
title: "15.3 Supermicro"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "compute", "datacenter", "gpu", "intel", "memory", "nvidia", "research"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [373, 402]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 1f0009b2208e2c9466aaf742d29078708ee8807e3dd85e85b33de50141ae0582
---

# 15.3 Supermicro

- Gen12 = Intel Xeon 6 (and 5th-gen AMD EPYC in DL3x5 lines) with PCIe Gen5, iLO 7, DDR5 up to 8 TB [secondary](https://www.storagereview.com/news/hpe-proliant-gen12-servers-unleashing-intel-xeon-6-performance-and-direct-liquid-cooling?_rt=MnwxfGhwZSBwcm9saWFudHwxNzQ5MTM4MDc3&_rt_nonce=1d16b5525f).
- DL360 Gen12 (1U 2P): "up to 20 EDSFF E3.S NVMe drives"; hybrid front cage supports SFF + E3.S + OS boot device + two front OCP NICs as options [secondary](https://buy.hpe.com/my/en/compute/hpe-proliant-compute-gen12-servers/c/c001030).
- DL380 Gen12 (2U 2P): "up to 36 EDSFF E3.S drives" [secondary](https://buy.hpe.com/my/en/compute/hpe-proliant-compute-gen12-servers/c/c001030).
- DL320/DL340 Gen12 (single-socket): "up to 20x EDSFF, 8x SFF, or 4x/10x/12x LFF" / "up to 20x EDSFF, 8x SFF, or 12x LFF" [secondary](https://www.storagereview.com/news/hpe-proliant-gen12-servers-unleashing-intel-xeon-6-performance-and-direct-liquid-cooling?_rt=MnwxfGhwZSBwcm9saWFudHwxNzQ5MTM4MDc3&_rt_nonce=1d16b5525f).
- DL325 Gen12 (AMD): tri-mode backplane "supports SATA, SAS4 plus NVMe SSD devices"; hybrid solution "up to five drive boxes, with each one supporting two SFF drives or four EDSFF E3.S NVMe SSDs"; backplane direct-attached to motherboard M-XIO connectors; HPE Smart Array Gen11 (no Gen12 models at review time); NS204i-u V2 rear boot with 2× 480 GB M.2 NVMe [secondary](https://www.itpro.com/infrastructure/servers-and-storage/hpe-proliant-compute-dl325-gen12-review-a-deceptively-small-and-powerful-1p-rack-server-with-a-huge-core-count).
- ML350 Gen12 (tower): "Up to 12 LFF SAS/SATA drives or 24 SFF SAS/SATA/NVMe drives or 12 EDSFF NVMe SSD drives"; HPE VROC for NVMe software RAID [secondary](https://farnamco.net/wp-content/uploads/2026/01/HPE-ProLiant-Compute-Family-Guide.pdf).
- Boot: NS204i-u V2 = "hot-pluggable, high-availability, RAID 1 protected M.2 NVMe boot option" [secondary](https://farnamco.net/wp-content/uploads/2026/01/HPE-ProLiant-Compute-Family-Guide.pdf).

### 15.3 Supermicro

- ARS-121L-NE316R (1U, NVIDIA Grace CPU Superchip): 16× hot-swap EDSFF PCIe Gen5 E3.S NVMe; "983 TB raw capacity using 61.44 TB SSDs"; 40 systems per rack = 39.3 PB raw; WEKA Data Platform validated; dual BlueField-3 / ConnectX-8 SuperNIC support; GPUDirect Storage [official](https://www.supermicro.com/en/pressreleases/supermicro-introduces-new-petascale-all-flash-storage-server-using-nvidia-grace-cpu).
- The Petascale line is Supermicro's all-flash NVMe brand for software-defined storage (WEKA, Ceph-class workloads) [official](https://www.supermicro.com/en/pressreleases/supermicro-introduces-new-petascale-all-flash-storage-server-using-nvidia-grace-cpu).
- [unverified] 2026 EPYC 9005 / Xeon 6 general-purpose Supermicro lines with EDSFF counts — not covered in this research; Supermicro's configurator data was not fetched.

### 15.4 OCP / whitebox / DC-MHS

- MSI DC-MHS (AMD EPYC 9005): CD281-S4051-X2 2OU 2-node, 21" ORv3 48 Vdc, "12 E3.S PCIe 5.0 NVMe bays per node", DC-SCM2 modules, EVAC cooling to 500 W TDP/node [vendor-reported](https://www.msi.com/news/detail/MSI-Showcases-DC-MHS-and-MGX-Server-Platforms--for-Cloud-Scale-and-AI-Infrastructure-at-OCP-APAC-2025-146777).
- DC-MHS's HPM/DC-SCM split means storage backplanes and drive cages can persist across CPU generations — relevant for E3.S investments with 5+ year horizons [secondary](https://medium.com/codex/from-traditional-servers-to-modular-servers-understanding-ocp-dc-mhs-hpm-and-dc-scm-2a55471033b8).
- KIOXIA CM9's OCP Datacenter NVMe SSD 2.5 compliance is the drive-side badge for these platforms [vendor-reported](https://www.silicon.co.uk/press-release/kioxia-announces-first-enterprise-nvme-ssd-built-with-8th-generation-bics-flash-tlc-based-flash-memory-technology).

### 15.5 Boot-device landscape (2026)

| Solution | Vendor | Form | Notes |
|---|---|---|---|
| BOSS-N1 | Dell | 2× M.2 NVMe carrier | 17G boot standard |
| NS204i-u V2 | HPE | Rear, 2× 480 GB M.2 NVMe hot-plug | RAID 1, high-availability boot |
| SATADOM / USB | Generic | DOM | Legacy/homelab |
| Network boot | NVMe Boot spec | NVMe-oF | New in NVMe 2.1 spec family |
| E1.S boot | Dell XE | EDSFF | GPU nodes boot from E1.S |

