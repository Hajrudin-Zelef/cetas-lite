---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/2-3-pm9d3a-gen5-mainstream-successor
title: "2.3 PM9D3a — Gen5 mainstream successor"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2021-12-22"]
keywords: ["datacenter", "latency", "nand", "pricing"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [46, 62]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 7c6efb69a08946f5bc8e2df1573f5260a83bfb7fe40fe9af6b5fd8795650f5d0
---

# 2.3 PM9D3a — Gen5 mainstream successor

- Announced by Samsung 2021-12-22 as its first PCIe 5.0 enterprise SSD with dual-port support [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers).
- Announcement targets: up to 13,000 MB/s seq read, 2,500K random read IOPS; 6,600 MB/s seq write, 250K random write IOPS; ~1.9x/1.7x over PCIe 4.0 predecessors; 608 MB/s-per-watt power efficiency (~30% better than prior gen) [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers).
- Capacities 1.92–15.36 TB; form factors 2.5" (15 mm U.2) and 3" EDSFF E3.S (7.5 mm doubles density vs 15 mm 2.5") [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers).
- In-market specs (Lenovo/SHI listings, 2026): 1.92 TB — 14,000/3,000 MB/s, 2,000K/150K IOPS, 3,504 TBW endurance, 1 DWPD, TLC, TCG Opal 2.0, 19.9 W typical, MTBF 2.5M h; 15.36 TB — 14,000/7,100 MB/s, 2,500K/360K IOPS, 28,032 TBW, 25 W typical [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD) [secondary](https://www.shidirect.com/product/45832454/THINKSYSTEM-2.5IN-U.3-PM1743-15.36TB-READ-INTENSIVE-NVME-PCIE-5.0).
- Drive class: Read Intensive; dual-port for HA failover [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD) [vendor-reported](https://www.businesswire.com/news/home/20211222005474/en/5120105/Samsung-Develops-High-Performance-PCIe-5.0-SSD-for-Enterprise-Servers).
- 2026 reseller pricing (Lenovo-branded): 1.92 TB at $8,968 (MSRP $8,699) — very high per-TB, reflects Lenovo channel markup, not NAND market price; see Section 14 [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD).

### 2.3 PM9D3a — Gen5 mainstream successor

- Interface PCIe Gen5 x4; up to 12,000/6,800 MB/s seq; up to 2,000K/400K random IOPS [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Capacities 960 GB–30.72 TB; form factors 2.5" (U.2), E1.S (9.5 mm), E3.S [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Endurance 1 DWPD, 5 years; enhanced PLP; OCP 2.5 compliant; positioned as Gen5 foundation for AI/hyperscale/cloud [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Reseller spec example (7.68 TB, SHI): 2,000K/300K IOPS, ~9 µs average latency, TLC with advanced ECC + end-to-end protection + SMART, TCG Opal 2.02 [secondary](https://www.shi.com/product/48742687/Samsung-PM9D3a-SSD).
- 2026 reseller pricing (Lenovo-branded): 7.68 TB at $17,362 (MSRP $16,899) ≈ $2,261/TB — channel pricing, not street NAND pricing [secondary](https://www.shi.com/product/48742687/Samsung-PM9D3a-SSD).

---

