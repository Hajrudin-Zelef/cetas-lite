---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/24-nand-industry-structure-2026
title: "24. NAND industry structure (2026)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["China", "Samsung"]
dates: ["2024-04", "2025-08", "2026-08-23", "2026-09-18", "2026-09-22"]
keywords: ["nand", "consumer", "cost", "datacenter", "dram", "ipo", "memory", "pricing", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [526, 575]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 8c5409187882c161df8463d07da8d0636d8c81ca845e6f3bc7912f3e98fe8ef5
---

# 24. NAND industry structure (2026)

## 24. NAND industry structure (2026)

- Samsung: V9 in mass production since April 2024 (~286–290L); V10 400+ layers demoed at ISSCC 2025 with hybrid bonding and 5.6 GT/s interface; V9 QLC delayed to H1 2026 [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding) [secondary](https://www.techpowerup.com/341061/samsung-delays-v9-qlc-nand-production-to-h1-2026-amid-performance-issues).
- Kioxia/SanDisk: BiCS8 218L CBA shipping (2 Tb QLC to Pure Storage for DirectFlash Modules); BiCS10 332L CBA announced at FMS 2026 — 4.8 GB/s Toggle DDR6.0, >37 Gb/mm² QLC, ~33% faster than BiCS8; Kioxia and SanDisk operate as separate vendors post-2024 spin-off but share BiCS roadmap [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).
- SK hynix (+Solidigm): 321L Gen9 TLC mass production; 321L QLC 2 Tb die in mass production H1 2026; 4D PUC architecture; Solidigm ETF filed ahead of Sept 4 IPO deadline [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493) [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Micron: G9 276L TLC, 1 Tb die, 21.0 Gb/mm², up to 3,600 MT/s — vertically integrated (NAND + controller + DRAM + firmware) [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding).
- YMTC: Xtacking 3.0 232L, 1 Tb QLC, 19.8 Gb/mm² — China domestic supply, relevant to cost pressure on QLC [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding).
- Structural note: NAND contract prices tightened through 2026 while DRAM is forecast to ease — SSD street prices follow NAND contracts with a 1–2 quarter lag; TrendForce sees no NAND relief before H2 2027 [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).
- CBA/hybrid bonding is now the standard path beyond ~300 layers because tall stacks thermally degrade CMOS logic — all five vendors have adopted some form of array/logic separation [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).

---

## 25. Controller market notes

- Merchant (third-party) controllers: Marvell Bravera SC5 (8/16ch Gen5), Phison X2 (16ch Gen5) — used by drive vendors without in-house controller teams [secondary](https://www.blocksandfiles.com/flash/2021/06/01/marvell-launches-first-pcie-5-ssd-controller/1616031).
- Vertically integrated (NAND + controller + firmware): Samsung, Kioxia, Micron, SK hynix/Solidigm — the four largest NAND producers all ship in-house controllers on flagships [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).
- Phison's dual role (merchant controller vendor + Pascari first-party SSD brand + IMAGIN+ custom design service) makes it both supplier and competitor to its own customers — a noted industry tension [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).
- Channel count matters: 16-channel controllers (Marvell SC5 MV-SS1333, Phison X2, SK hynix Atomos Prime) feed Gen5 x4 bandwidth; 8-channel designs (Marvell MV-SS1331, <8.7 W) target power-optimized E1.S [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- NAND interface speed: Marvell SC5 supports up to 1,600 MT/s; Samsung V10 targets 5.6 GT/s; Kioxia BiCS10 Toggle DDR6.0 4.8 GB/s — interface speed is now a headline spec alongside layer count [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).
- Firmware is the moat: FTL quality, GC efficiency, QoS determinism, and telemetry separate vendors more than raw NAND — which is why hyperscalers co-design via programs like Phison IMAGIN+ [secondary](https://www.tweaktown.com/news/98293/pascari-is-new-enterprise-brand-from-phison-the-biggest-behind-scenes-name-in-storage/index.html).

---

## 26. Retail/consumer price divergence (context)

- MemoryPriceChart (2026-09-18): DDR5 index +397% vs August 2025 backcast; SSD index +158%; DRAM supply remains tight (Nanya stop-ship, allocation) while SSD market is comparatively fluid [secondary](https://lifestyle.successxl.com/story/867529/ram-and-ssd-prices-diverge-as-ddr5-index-rises-397-versus-158-for-ssds/).
- Consumer basket: median $159.94/TB, mean $168.83/TB for 1–2 TB PCIe 4.0 NVMe (15+ retail SKUs) — mainstream consumer SSDs remain far cheaper per TB than new enterprise NVMe ($300–$1,172/TB), reflecting PLP, DRAM, validation, and warranty costs [secondary](https://lifestyle.successxl.com/story/867529/ram-and-ssd-prices-diverge-as-ddr5-index-rises-397-versus-158-for-ssds/).
- Retailer behavior noted: one vendor raised "free shipping" threshold from $50 to $99 as RAM prices surged — a proxy for memory-market stress [secondary](https://lifestyle.successxl.com/story/867529/ram-and-ssd-prices-diverge-as-ddr5-index-rises-397-versus-158-for-ssds/).
- Takeaway for RAG: enterprise and consumer flash markets decoupled in 2026 — do not extrapolate consumer SSD pricing to datacenter SSDs [independent guidance].

---

*End of Step 9 Phase A — Enterprise SSD hardware. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

---

## 27. $/TB and $/TBW economics (dated snapshots, 2026)

| Drive / SKU | Price (snapshot) | $/TB | Rated TBW | Price per TBW | Context |
|---|---|---|---|---|---|
| Samsung PM1743 1.92 TB (Lenovo) | $8,968 | $4,671 | 3,504 | $2.56 | OEM channel markup [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD) |
| Samsung PM9D3a 7.68 TB (Lenovo) | $17,362 | $2,261 | 14,016 | $1.24 | OEM channel [secondary](https://www.shi.com/product/48742687/Samsung-PM9D3a-SSD) |
| Kioxia CD8P-R 30.72 TB SED (Lenovo) | $21,685.99 | $706 | 56,064 | $0.39 | large-capacity discount visible [secondary](https://corgitech.us/products/lenovo-thinksystem-read-intensive-30-72tb-2-5-pcie-5-0-sed-kioxia-cd8p-r-ssd) |
| Micron 9550 PRO 7.68 TB | €1,597 excl. VAT | €208 | 14,016 (RND) | €0.11 | street price, one data point [secondary](https://www.computeruniverse.net/en/p/3314-02V) |
| WD SN861 3.2 TB (3 DWPD) | £572.81 excl. VAT | £179 | 17,520 | £0.033 | 3-DWPD class best $/TBW [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca232psp9x1-ssd-data-centre-3-2-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041686/) |
| Phison X200E 6.4 TB (at review) | $1,309.99 | $205 | 35,040 | $0.037 | pre-2026-surge price [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html) |
| VDURA reference 30 TB TLC | $22,600 | $753 | ~54,750 | $0.41 | vendor-published ref [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/) |
| Used ent. SATA/SAS 3.84 TB | — | $78–$127 | varies | — | decommissioned pool [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/) |
| New ent. NVMe U.2 (market) | — | $300–$1,172 | varies | — | 2026-08-23 band [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/) |

