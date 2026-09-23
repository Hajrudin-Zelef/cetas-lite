---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/17-workload-decision-guide
title: "17. Workload decision guide"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["cost", "datacenter", "dram", "ipo", "latency", "memory", "nand"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [422, 476]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 15a7c74d6512b70d0b6b5b139c24bc0dadf221c554198bdbeb649142ae369d3f
---

# 17. Workload decision guide

## 17. Workload decision guide

- Read-intensive (1 DWPD): CDN, media streaming, boot images, OS/application read caches, AI model serving/checkpoints, web content, analytics reads — Samsung PM9A3/PM9D3a, Kioxia CM7-R/CD8P-R, Micron 9550 PRO, WD SN861 1-DWPD SKUs, Solidigm D7-PS1010 [independent guidance].
- Mixed-use (3 DWPD): OLTP databases, virtualization hosts (VMware/Hyper-V), Ceph/Rook OSD journals + data on the same pool, container image registries with heavy layer churn, VDI — Micron 9550 MAX / 7450 MAX, Kioxia CM7-V/CD8P-V, Solidigm D7-PS1030, WD SN861 3-DWPD SKUs [independent guidance].
- Ultra-write (25–60 DWPD pSLC): write caches, Ceph WAL/DB devices, ZFS SLOG, RocksDB WAL, deduplication fingerprint stores — Phison X200Z (60 DWPD), Solidigm D7-P5810 SLC line [secondary](https://www.thessdreview.com/our-reviews/enterprise/phison-pascari-x200z-gen5-800gb-1-6tb-enterprise-ssd-review-slc-gold-commands-a-lightning-fast-60-dwpd-data-center-ssd/).
- Read-heavy capacity (QLC ≤1 DWPD): object storage, backup/archival, AI data lakes, video surveillance retention, CDN origin — Solidigm D5-P5336 61.44/122.88 TB, Kioxia LC9 245.8 TB / Sandisk UltraQLC SN670 256 TB (announced) [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Mixed workloads on 1-DWPD TLC: acceptable when the average write rate stays under capacity/5yr/day; a 15.36 TB 1-DWPD drive tolerates ~8.4 TB/day average — measure with `nvme smart-log` data-units-written before deciding a drive is "read-only class" [independent guidance].
- ZFS note: special vdevs and metadata benefit from low-latency TLC (Micron 9550 60/15 µs, WD SN861 65/8 µs); L2ARC is read-only so QLC is acceptable; SLOG needs PLP + high DWPD (pSLC ideal) [independent guidance].
- Ceph note: BlueStore DB/WAL on 3-DWPD or pSLC devices, bulk OSDs on 1-DWPD TLC; bluestore allocator benefits from FDP/ZNS-capable drives (Kioxia CM7, WD SN861) for WAF reduction [independent guidance].
- VMware vSAN ESA / Azure Stack HCI: prefer dual-port (Kioxia CM7, Phison X200 dual-port, Samsung PM1743) for HA path redundancy; single-port CD8P/SN861 U.2 suit single-host or JBOD topologies [independent guidance].
- Do not size endurance from the workload's peak day — size from the 95th-percentile daily write over the warranty window, plus headroom for compaction/GC write amplification (WAF 1.5–3x on random-write workloads) [independent guidance].

---

## 18. Form factors and interface math

- U.2/U.3 2.5" × 15 mm: the dominant datacenter form factor; U.3 (SFF-TA-1001) is tri-mode (NVMe/SAS/SATA) backplane-compatible; 25 W envelope typical [secondary](https://www.tweaktown.com/news/95016/kioxia-ssds-achieve-pcie-5-0-compliance-up-to-30tb-capacity-and-14-000-mb/index.html).
- EDSFF E3.S 1T (7.5 mm): ~2x density vs 15 mm 2.5" at the same rack height; native Gen5 x4 (and x8 in E3.L); adopted by Samsung PM1743/PM9D3a, Kioxia CM7, Micron 9550, Solidigm D7-PS1010, Phison X200 [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers).
- EDSFF E1.S (9.5/15/25 mm): 1U "ruler" optimized; Marvell's 16-channel SC5 in 20×20 mm enabled the first 16CH E1.S SSDs; Samsung PM9A3/PM9D3a, Micron 7450/9550, WD SN861 ship E1.S variants [secondary](https://www.guru3d.com/story/marvell-announces-bravera-sc5-ssd-controller-family-supports-pcie-gen-5-(17gbs)/?print).
- M.2 2280/22110: boot and edge roles (Micron 7450, Samsung PM9A3, Phison B-Series); thermally constrained — Gen5 M.2 enterprise drives are rare [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- PCIe bandwidth: Gen4 x4 ≈ 8 GB/s theoretical (~6.9 GB/s achievable); Gen5 x4 ≈ 16 GB/s theoretical (~14–14.8 GB/s achieved by X200/9550/PS1010) — current flagships saturate Gen5 x4 sequential reads [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- Dual-port (2×2): two independent PCIe paths to the same namespaces for active-active HA failover — Kioxia CM7, Samsung PM1743, Phison X200; single-port is cheaper and sufficient for non-HA topologies [vendor-reported](https://www.phisonenterprise.com/wp-content/uploads/2025/01/PascariProductBrochure_X200_010925.pdf).
- NVMe 2.0 feature set (ZNS, FDP, KV command set, NVMe-MI 1.2c) is the 2025–2026 baseline on Gen5 drives; Gen4 drives sit on NVMe 1.4/1.4b [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- OCP Datacenter NVMe SSD spec v2.0 (+ v2.5 telemetry): required by hyperscale buyers; supported by Samsung PM9D3a, Micron 9550, WD SN861, Kioxia CD8P — check "not all requirements" footnotes on Kioxia [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).

---

## 19. DRAM cache vs HMB and enterprise firmware

- Enterprise SSDs carry onboard DRAM (typically 1 GB per TB, e.g. 8 GB DDR4 on Phison X200E 6.4 TB) for the FTL mapping tables, buffering, and power-loss-safe metadata — onboard DRAM is a reliability/performance requirement, not an option [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- HMB (Host Memory Buffer): borrows host DRAM for FTL tables over PCIe — common in DRAM-less client SSDs (cost/power play); generally avoided in enterprise designs because it couples drive behavior to host memory pressure and complicates PLP guarantees [independent guidance].
- The exception: some E1.S boot/datacenter drives use HMB to hit cost targets — check the datasheet "DRAM" line rather than assuming [independent guidance].
- Enterprise firmware differentiators vs client: deterministic QoS (99.999th-percentile latency SLAs), full-path PLP, TCG security stacks, SR-IOV/CMB, NVMe-MI manageability, telemetry (OCP 2.5), FDP/ZNS support, signed firmware + secure boot, and dual-port pathing [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- Marvell's "elastic SLA enforcer" is a hardware QoS engine example — allocates IOPS/latency budgets per namespace in the controller rather than firmware alone [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- Thermal design: Gen5 drives are 18–29 W active (Solidigm PS1010 7.68 TB ~22–29 W max; Micron 9550 ≤18 W read; WD SN861 E1.S 12 W avg/21 W max) — all require directed airflow; the liquid-cooled D7-PS1010 E1 variant exists precisely because air-cooling dense Gen5 E1.S is marginal in AI racks [secondary](https://www.tomshardware.com/pc-components/ssds/solidigm-touts-industrys-first-liquid-cooled-enterprise-ssd-d7-ps1010-is-an-e-1-pcie-5-0-drive-with-a-wrap-around-cold-plate).

---

## 20. Endurance worked examples (5-year warranty)

- Samsung PM9A3 7.68 TB, 1 DWPD: 7.68 × 1 × 365 × 5 = 14,016 TBW ≈ 14.0 PBW (published 14.02 PBW) [secondary](https://www.superstorage.pl/pdf_datasheet.php?products_id=5921&osCsid=639b917cb1adac10beac4c137376fb04).
- Samsung PM9A3 1.92 TB, 1 DWPD: 1.92 × 365 × 5 = 3,504 TBW ≈ 3.50 PBW (published) [secondary](https://www.superstorage.pl/pdf_datasheet.php?products_id=5921&osCsid=639b917cb1adac10beac4c137376fb04).
- Samsung PM1743 15.36 TB, 1 DWPD: 15.36 × 365 × 5 = 28,032 TBW (published 28,032 TBW — exact match) [secondary](https://www.shidirect.com/product/45832454/THINKSYSTEM-2.5IN-U.3-PM1743-15.36TB-READ-INTENSIVE-NVME-PCIE-5.0).
- Samsung PM1743 1.92 TB, 1 DWPD: 1.92 × 365 × 5 = 3,504 TBW (published) [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD).
- Solidigm D7-PS1010 7.68 TB, 1 DWPD: 14,016 TBW (published; matches formula) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- Micron 9550 PRO 30.72 TB, 1 DWPD: 30.72 × 365 × 5 = 56,064 TBW random (published 56,064 — exact) [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Micron 9550 MAX 25.6 TB, 3 DWPD: 25.6 × 3 × 365 × 5 = 140,160 TBW random (published 140,160 — exact) [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Micron 7450 PRO, 1 DWPD: up to 28,000 TBW (15.36 TB: formula gives 28,032 — vendor rounds to 28,000) [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Micron 7450 MAX, 3 DWPD: up to 70,000 TBW (12.8 TB: formula gives 70,080 — vendor rounds) [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Phison X200E 6.4 TB, 3 DWPD: 6.4 × 3 × 365 × 5 = 35,040 TBW (published) [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- Phison X200Z 1.6 TB, 60 DWPD: 1.6 × 60 × 365 × 5 = 175,200 TBW — pSLC endurance at ~27x the X200E's per-TB rate [secondary](https://www.thessdreview.com/our-reviews/enterprise/phison-pascari-x200z-gen5-800gb-1-6tb-enterprise-ssd-review-slc-gold-commands-a-lightning-fast-60-dwpd-data-center-ssd/).
- Rule of thumb: remaining life ≈ (1 − percentage_used/100) × rated TBW; at 10 TB/day host writes with WAF 2.0 on NAND, a 14,016-TBW drive lasts ~1,918 days ≈ 5.25 years — right at warranty, which is by design [independent guidance].

---

