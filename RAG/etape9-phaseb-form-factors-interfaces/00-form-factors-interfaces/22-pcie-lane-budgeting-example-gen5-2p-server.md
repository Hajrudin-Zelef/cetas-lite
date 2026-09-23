---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/22-pcie-lane-budgeting-example-gen5-2p-server
title: "22. PCIe lane budgeting example (Gen5 2P server)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["AMD", "Broadcom", "Samsung"]
dates: []
keywords: ["compute", "cost", "datacenter", "inference", "memory", "nvidia", "serdes"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [612, 668]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: a242da7871a364bd92f4ebe0351a5bce421a3dfe53c0187fbc61f8d821322f66
---

# 22. PCIe lane budgeting example (Gen5 2P server)

## 22. PCIe lane budgeting example (Gen5 2P server)

A dual-socket Xeon 6 / EPYC 9005 server offers on the order of 160+ PCIe Gen5 lanes per system (socket-dependent; verify exact SKU):

- 24× E3.S Gen5 x4 NVMe = 96 lanes direct-attached (no switch).
- 32× E3.S via Gen5 x2 (Dell R570 32-drive variant) = 64 lanes — the x2 tradeoff documented on the R570 chassis [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations).
- PCIe switch (e.g. Broadcom PEX89048, 48 ports): 1× x16 host → 16× x4 NVMe on the HighPoint Rocket 1624A pattern, at the cost of oversubscription when all drives run flat-out [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).
- Rule of thumb: direct-attach lanes give deterministic bandwidth; switched/oversubscribed lanes trade peak bandwidth for density — acceptable when average, not peak, I/O dominates [independent].
- Bifurcation must be enabled in BIOS/UEFI for passive x16→4×x4 adapters; switch adapters do not need it [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).

## 23. Retimers, redrivers, and cable reach

- **Redriver**: analog equalizer/amplifier; extends copper reach at Gen4/5 but adds jitter; passive in-band, no protocol awareness [independent].
- **Retimer**: full clock-data recovery, re-transmits a clean signal; required for longer Gen5 (and all Gen6/7) channels; the PCIe 7.0 Optical Aware Retimer ECN extends retimers to fiber [secondary](https://convergedigest.com/pci-sig-finalizes-pcie-7-0-specification-at-128-0-gt-s/).
- Gen5 MCIO copper is rated to ~1.0 m; beyond that, or through multiple connectors, budget a retimer card or retimed cable [vendor-reported](https://www.newegg.ca/p/238-00D8-00027)[independent].
- Motherboard MCIO/U.2 traces already account for channel loss; adding third-party breakout cables on top can exceed the loss budget — prefer vendor-qualified cables [independent].
- Clock buffers (e.g. 100 MHz on Icy Dock's Gen5 adapter) are part of the same signal-integrity story for enterprise timing [vendor-reported](https://datasheet.itscope.com/2.1/t/Hc4xNMicnBB3UhUhPI7m5PtNPzD0l5IEpv3OCf9-3DLS6QFmBJ1X4enr1g6nayNaxVVAqsib773xL6nayNaxVVAqsib773xL6rAddIgK0ysFp4G5BB-pHMRZdELYWY76YkdU7i2THKwwJYzCHH9wzfwyQ9eeDukqo7AP5WTJh8fSFx7-2rfftuzR-I0fP8Q_HPlvGNMEFdCTcgRsloU).

## 24. CXL note (storage-adjacent)

- CXL rides PCIe 5.0/6.0 PHYs and appears on Dell 17G (CXL 2.0 on R670/R770/R7725/R6725) as memory expansion, not storage [secondary](https://www.wwt.com/article/unlocking-the-future-of-the-data-center-the-evolution-of-dell-poweredge-servers).
- Relevance to this file: CXL and NVMe compete for the same SerDes/lane budget and retimer infrastructure; CXL-attached memory pools and NVMe-oF storage pools are the two faces of PCIe fabric disaggregation in 2026 [independent].
- CXL.mem vs NVMe block semantics are distinct — CXL is byte-addressable load/store, NVMe is block/FDP/ZNS/KV; they complement rather than replace [independent].
- Deep CXL coverage (CXL 2.0/3.x, memory pooling, HDM) belongs to the memory track of this phase, not this file. [scope note]

## 25. Storage software stack checklist (hardware-adjacent)

- **OS/NVMe driver**: Linux nvme + nvme-cli (MI admin commands, ISH bit per NVMe-MI 2.1) [secondary](https://github.com/linux-nvme/nvme-cli/commit/5be98d6689ed478cdea17562b008189583824323).
- **Multipath**: native nvme-multipath + ANA; dual-port NVMe for path redundancy [independent].
- **Software RAID**: HPE VROC, Dell S160, Linux mdadm — for boot mirrors and simple NVMe sets [secondary](https://farnamco.net/wp-content/uploads/2026/01/HPE-ProLiant-Compute-Family-Guide.pdf).
- **Hardware RAID**: Broadcom 96xx (RAID 0/1/5/6/10/50/60, CacheVault); PERC 12 H365i/H965i on Dell 17G [vendor-reported](https://5.imimg.com/data5/SELLER/Doc/2025/11/559109987/XE/CD/PL/228806307/broadcom-24gb-sas-nvme-adapter.pdf).
- **SDS**: Ceph, WEKA (validated on Supermicro Petascale), Lightbits NVMe/TCP [official][vendor-reported](https://www.storagenewsletter.com/2026/01/16/lightbits-labs-delivers-record-growth-as-organizations-standardize-on-nvme-over-tcp-for-ai-ready-infrastructure/).
- **Userspace I/O**: SPDK with KV command-set support and NVMe-oF target/initiator — the AI-storage fast path [secondary](https://github.com/spdk/spdk/blob/HEAD/CHANGELOG.md).
- **Zoned storage tooling**: zonefs, f2fs, and LSM engines (RocksDB) for ZNS; D2PF (Samsung/WD) for ecosystem convergence [secondary](https://www.eetimes.com/samsung-western-digital-unite-around-zoned-storage).
- **Telemetry**: NVMe-MI + OCP Datacenter NVMe SSD telemetry; BMC/I2C sideband via UBM [vendor-reported](https://www.microsemi.com/product-directory/upcoming-technology/5558-universal-backplane-management-ubm).

## 26. Capacity points and endurance classes (2026)

### 26.1 Where the TB ladder stands

| Class | Examples (2026) | Form factors | Notes |
|---|---|---|---|
| 7.68 TB | Mixed-use NVMe workhorses | U.2/U.3/E3.S | Common enterprise sweet spot |
| 15.36 TB | KIOXIA CM9-R (read-intensive) | 2.5" / E3.S | Gen5 flagship reviewed [secondary](https://www.tweaktown.com/reviews/11605/kioxia-cm9-r-15-36tb-e3-s-enterprise-ssd-performance-king-at-1-dwpd/index.html) |
| 30.72 TB | High-capacity TLC | U.2/U.3/E3.S | Mainstream high-cap |
| 61.44 TB | Supermicro Petascale node drives | E3.S | 16× = 983 TB raw per 1U node [official](https://www.supermicro.com/en/pressreleases/supermicro-introduces-new-petascale-all-flash-storage-server-using-nvidia-grace-cpu) |
| 122.88 TB | E1.L/E3.L rulers | E1.L / E3.L | Dell XE8640: 8× E3.S = 122.88 TB max [secondary](https://expresscomputersystems.com/collections/all-computer-components/products/dell-poweredge-xe8640-8-e3-s-edsff-rack-server-chassis) |
| 245.76 TB | Next-gen ruler class | E1.L | 2025–2026 announcements [secondary] |

### 26.2 Endurance classes

- **Read-intensive (RI)**: ~1 DWPD — boot, CDN, AI inference read caches; CM9-R is the cited 1-DWPD-class example [secondary](https://www.tweaktown.com/reviews/11605/kioxia-cm9-r-15-36tb-e3-s-enterprise-ssd-performance-king-at-1-dwpd/index.html).
- **Mixed-use (MU)**: ~3 DWPD — virtualization, databases [independent].
- **Write-intensive (WI)**: higher DWPD, smaller capacities — journaling, write caches [independent].
- ZNS changes the math: Samsung PM1731a claims "up to four times longer" endurance by eliminating GC-driven write amplification [vendor-reported](https://www.enterprisestorageforum.com/news/samsung-zns-ssd/).
- Endurance and over-provisioning interact: ZNS "eliminating the need for overprovisioning" gives back the full advertised capacity [vendor-reported](https://www.enterprisestorageforum.com/news/samsung-zns-ssd/).

