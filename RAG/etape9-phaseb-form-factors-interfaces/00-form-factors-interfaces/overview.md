---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/overview
title: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Intel", "Microsoft"]
dates: ["2026-04-01", "2026-08", "2026-09-22"]
keywords: ["compute", "datacenter", "hyperscaler", "intel", "latency", "memory", "nand", "nvidia", "research", "throughput"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [1, 36]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 218a18db423215e3a85cb017f793b9df2f9f78cb728ad3c487a6d04668f45775
---

# Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)

*Scope: physical form factors for enterprise/datacenter SSDs (M.2, U.2/U.3, EDSFF E1.S/E1.L/E3.S/E3.L, 2.5" SATA/SAS, AIC), the NVMe specification family (2.0 → 2.4, current August 2026), NVMe-oF transports, PCIe/SAS/SATA interface generations, advanced placement features (ZNS, FDP), backplane management, storage cabling, and 2026 server platform integration.*
*Research date / cutoff: 2026-09-22. Method: read-only web research (search + page fetch). Provenance legend: `[official]` = standards body or vendor primary source; `[vendor-reported]` = vendor press/blog claims; `[independent]` = third-party lab/review measurement; `[secondary]` = press/analyst reporting; `[unverified]` = single-source or otherwise unconfirmed claim. Facts someone will act on (specs, dates, power, prices, SKUs) are tagged; invented SKUs/URLs/prices are never used.*
*Single writer; no other workspace files modified.*

---

## 1. Why form factors still matter in 2026

- Flash outgrew the mechanical envelope: EDSFF was purpose-built for NVMe drives "whether they contain flash or another storage class memory (SCM) such as Intel Optane", because the AIC/CEM slot steals PCIe slots from other devices and has limited hot-plug, the 2.5" HDD-derived form factor "blocks airflow to the hottest components in the server" and "is not natively designed for flash", and M.2 has "low capacity without any hot-plug features and suffer from limited power and thermal scaling for data center use" [secondary](https://eetimes.com/nvme-drives-ready-to-embrace-own-form-factors/).
- All major SSD makers ship EDSFF products, with hyperscaler support from Facebook and Microsoft behind the family [secondary](https://eetimes.com/nvme-drives-ready-to-embrace-own-form-factors/).
- Form factor choice in 2026 is a density/power/servicing tradeoff, not a performance tier per se: "compare the exact drive's capacity, PCIe link, sequential throughput, random IOPS, latency, endurance, workload class, and sustained-performance specifications rather than choosing by EDSFF form factor alone" [secondary](https://www.ariat-tech.com/blog/EDSFF-SSD-Guide-E1.S-vs.E1.L-vs.E3.S-vs.E3.L.html).
- The three 2026 mainstream enterprise buckets: 2.5"/U.3 tri-mode for mixed fleets, EDSFF E3.S for density-optimized NVMe, E1.S for 1U compute-heavy nodes [secondary](https://dev-snia-org.pantheonsite.io/sites/default/files/2025-10/SNIA-SDC25-Constantine-Why-Another-EDSFF-SSD.pdf).
- AI storage nodes in 2026 are built around front-access hot-swap NVMe density (e.g. 16–36+ E3.S bays per 1U/2U server) rather than raw per-drive speed [secondary](https://www.supermicro.com/en/pressreleases/supermicro-introduces-new-petascale-all-flash-storage-server-using-nvidia-grace-cpu).

## 2. M.2 — client-born, datacenter-boot

- M.2 (NGFF) module sizes are named width × length in mm: 2280 (22 × 80 mm) is the mainstream client size; 22110 (22 × 110 mm) is the enterprise/datacenter length with room for power-loss-protection capacitors and more NAND packages [independent].
- M.2 keying: M-key (PCIe x4 NVMe) vs B-key (SATA/PCIe x2); enterprise M.2 NVMe drives use M-key [independent].
- Typical enterprise M.2 NVMe is PCIe x4 with ~8–10 W power envelope; capacity tops out well below 2.5"/EDSFF because of PCB area and thermal limits — the reason it "found their way into the data center because they were small and modular, but they're low capacity without any hot-plug features" [independent][secondary](https://eetimes.com/nvme-drives-ready-to-embrace-own-form-factors/).
- 2026 datacenter role is overwhelmingly **boot/OS**, not data: Dell PowerEdge XE AI servers use "1 x M.2 Boot" / BOSS-N1; HPE Gen12 uses the NS204i-u V2 rear module "fitted with two hot-plug 480GB M.2 NVMe SSDs to provide mirrored redundant storage for running an OS or hypervisor" [secondary](https://secureserve.co.th/uploads/product-datasheet/2026-04-01/69cc9d92b815f.pdf)[secondary](https://www.itpro.com/infrastructure/servers-and-storage/hpe-proliant-compute-dl325-gen12-review-a-deceptively-small-and-powerful-1p-rack-server-with-a-huge-core-count).
- No native hot-swap: M.2 is a board-edge connector on the motherboard or a carrier (BOSS); surprise-removal handling is OS/driver dependent, unlike U.2/U.3/EDSFF sleds [independent].
- Bifurcation note: x4 M.2 NVMe needs motherboard PCIe bifurcation support when split from x8/x16 slots [independent].

## 3. U.2 (SFF-8639) and U.3 (SFF-TA-1001) — the 2.5" NVMe era

### 3.1 U.2 / SFF-8639

- U.2 is the 2.5" (15 mm z-height) NVMe drive connector SFF-8639: PCIe x4 (up to x4 lanes) plus sideband, in the familiar 2.5" hot-swap sled with 25 W typical power envelope [independent].
- U.2 NVMe is electrically distinct from SAS/SATA wiring: early U.2 backplanes needed NVMe-aware cabling, which motivated the U.3 unification [secondary](https://www.storagereview.com/news/evolving-storage-with-sff-ta-1001-u-3-universal-drive-bays?amp).
- 2026 status: U.2 remains the dominant enterprise NVMe bay in installed fleets, but new 2026 platforms bias toward E3.S for density; Dell 17G and HPE Gen12 still offer 2.5" universal/tri-mode bays alongside E3.S options [secondary](https://www.dell.com/en-au/shop/servers-storage-and-networking/poweredge-r570-rack-server/spd/poweredge-r570/promo_r570_1?view=configurations)[secondary](https://buy.hpe.com/my/en/compute/hpe-proliant-compute-gen12-servers/c/c001030).
- KIOXIA CM9 ships in both "2.5-inch and E3.S SSD form factors" with dual-port support, showing the 2.5" bay still gets flagship PCIe 5.0 drives in 2026 [vendor-reported](http://www.techpowerup.com/336800/kioxia-announces-first-enterprise-nvme-ssd-with-8th-gen-bics-flash-technology).

### 3.2 U.3 / SFF-TA-1001 — the universal bay

