---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/overview
title: "Step 9 — Enterprise SSD Hardware (Phase A)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2021-12-22", "2025-08", "2026-04", "2026-08-23", "2026-09-18", "2026-09-22"]
keywords: ["consumer", "datacenter", "dram", "hyperscaler", "ipo", "latency", "nand", "pricing", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [1, 62]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: f40c454e21ea38598f68ba333345726637c504dca656eba85940502edcf4cff6
---

# Step 9 — Enterprise SSD Hardware (Phase A)

## 0. Scope, method, provenance

- Scope: flagship datacenter/enterprise SSDs as of 2026-09-22 (specs, NAND, controllers, endurance, prices, homelab angle). Hardware angle only; storage software (Ceph, Rook, ZFS) was covered in Step 7 Phase E.
- Research date: 2026-09-22. All prices dated; enterprise SSD pricing is highly volatile in 2026 — treat every price as a dated snapshot, not a quote.
- Method: read-only web research (vendor datasheets/product pages, reseller listings, independent reviews: Tom's Hardware, ServeTheHome-class press, AnandTech-derived articles, TechPowerUp SSD database, Blocks & Files).
- Provenance tags: [official] = vendor datasheet/page/PR; [vendor-reported] = vendor claim quoted by press; [independent] = independent test/lab review; [secondary] = press/analyst/reseller; [unverified] = could not be corroborated. Conflicts and gaps are registered in Section 14.
- Naming note: "SK Hynix"/"SK hynix" styling varies by source; copied as found. "BiCS FLASH" = Kioxia/SanDisk 3D NAND brand. DWPD = Drive Writes Per Day over the warranty period.

---

## 1. 2026 market and pricing context

- Enterprise SSD prices rose ~6.5x year-over-year while nearline HDD rose ~2.45x over the same window, per VDURA's Flash Volatility Index (Aug 2026 update); VDURA sells hybrid arrays, so read its index accordingly, but TrendForce contract data points the same direction [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- The widely repeated "+472%" figure is an April 2026 VDURA number about its own product economics, already superseded by the same vendor's later data: its August update prices a 30 TB TLC drive at $22,600 [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- Priced 2026-08-23: new enterprise NVMe U.2 runs $300–$1,172/TB; nearline SATA is flat at $25/TB across 12/16/20 TB drives; used enterprise SATA/SAS SSDs at 3.84 TB land at $78–$127/TB — the used channel is priced off a decommissioned-hardware pool, not this quarter's NAND contracts [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- Tiering arithmetic at those prices: 100 TB usable costs about $6,600 as hot NVMe over bulk disk vs ~$37,500 all-new-flash — a 5x gap [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- No price relief expected before H2 2027, per TrendForce; DRAM is forecast to ease while NAND tightens [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- Retail (consumer) divergence, 2026-09-18: MemoryPriceChart's DDR5 index is +397% vs August 2025 backcast while its SSD index is +158%; basket median $159.94/TB for 1–2 TB PCIe 4.0 NVMe drives, mean $168.83/TB [secondary](https://lifestyle.successxl.com/story/867529/ram-and-ssd-prices-diverge-as-ddr5-index-rises-397-versus-158-for-ssds/).
- AI demand is the structural driver: hyperscaler AI storage tiers show read-to-write ratios up to 20:1+, which is why QLC endurance tradeoffs are tolerated there [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- IDC (via Kioxia launch coverage): EDSFF form factor expected to exceed 50% of enterprise SSD unit shipments by 2026 [vendor-reported](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html).

---

## 2. Samsung datacenter SSDs

### 2.1 PM9A3 — Gen4 mainstream workhorse

- Interface: PCIe Gen4 x4, NVMe 1.4, PCI Express Base Spec Rev 4.0 [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- NAND: Samsung V6 (128-layer) V-NAND TLC [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Form factors: U.2 (7 mm), E1.S, M.2 (2280/22110) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Capacities: 960 GB / 1.92 / 3.84 / 7.68 / 15.36 TB (product brief); samsung.com datacenter page lists 480 GB–15.36 TB for U.2/M.2/E1.S — capacity-list conflict, see Section 14 [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf) [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Performance (brief): up to 6,800 MB/s seq read, 4,000 MB/s seq write; 1,000,000 random read IOPS, 180,000 random write IOPS [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- samsung.com headline figures: up to 6,900/4,100 MB/s and 1,100K/200K IOPS (PM9A3) — slightly higher than the product-brief maxima; treat as capacity/best-case variance, not a conflict [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Endurance: 1.0 DWPD for 5 years; UBER 1 sector per 10^17 bits read; MTBF 2 million hours [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Power: read ≤11 W, write ≤13.5 W (U.2/E1.S); ≤8 W both (M.2) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Data path: PLP, end-to-end data protection, AES-XTS 256-bit encryption, TCG Enterprise; 512/4096-byte sectors; SMBus NVMe basic management [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Per-capacity endurance (reseller datasheets, 2026): 1.92 TB = 3.50 PBW; 3.84 TB = 7.01 PBW; 7.68 TB = 14.02 PBW — all at 1 DWPD [secondary](https://www.superstorage.pl/pdf_datasheet.php?products_id=5921&osCsid=639b917cb1adac10beac4c137376fb04).
- Per-capacity performance examples (1.92 TB SED): 6,800/2,700 MB/s, 740K/130K IOPS; (7.68 TB): 6,700/4,000 MB/s, 1,100K/200K IOPS — smaller capacities are write-limited [secondary](https://www.superstorage.pl/pdf_datasheet.php?products_id=5921&osCsid=639b917cb1adac10beac4c137376fb04).
- Positioned for cloud services, transactional databases, CDNs [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Homelab relevance: PM9A3 is one of the most common used enterprise Gen4 drives on the secondary market; check `nvme smart-log` percentage-used and power-on hours before buying [independent guidance].

### 2.2 PM1743 — first PCIe 5.0 enterprise SSD (announced 2021-12)

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

