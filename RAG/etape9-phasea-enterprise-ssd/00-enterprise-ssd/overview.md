---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/overview
title: "Step 9 — Enterprise SSD Hardware (Phase A)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2025-08", "2026-04", "2026-08-23", "2026-09-18", "2026-09-22"]
keywords: ["consumer", "datacenter", "dram", "hyperscaler", "ipo", "nand", "pricing", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [1, 45]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: d5702406083af953549082c8a972b4cbd7d53123968cf8fe477a6a122cee52cd
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

