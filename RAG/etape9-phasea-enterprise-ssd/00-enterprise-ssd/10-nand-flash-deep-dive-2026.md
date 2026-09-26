---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/10-nand-flash-deep-dive-2026
title: "10. NAND flash deep dive (2026)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: ["2024-04"]
keywords: ["nand", "cost", "datacenter", "latency", "pricing", "research"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [228, 269]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 7e0775bcb3b6db9e5188fee5ae59be7011f46af2a728b0784034ead58fd7b488
---

# 10. NAND flash deep dive (2026)

## 10. NAND flash deep dive (2026)

### 10.1 Bits per cell

- SLC (1) / MLC (2) / TLC (3) / QLC (4) / PLC (5): voltage-state count doubles per added bit (SLC 2 states → QLC 16 states/15 thresholds → PLC 32 states/31 thresholds) [secondary](https://www.blocksandfiles.com/flash/2026/01/15/sk-hynix-developing-split-cell-5-bit-flash/4090388).
- TLC is the enterprise performance standard; QLC is the capacity/cost play for read-heavy tiers; PLC is not commercial (read reliability too low, endurance too short) — would add ~25% die capacity vs QLC [secondary](https://www.trendforce.com/news/2026/01/16/news-sk-hynix-unveils-5-bit-nand-that-splits-cells-delivers-20x-faster-reads/).
- Indicative P/E cycles: TLC ~3,000; QLC ~1,000; PLC <500 (estimated) — actual endurance depends on controller, firmware, overprovisioning [secondary](https://oretonstorage.com/blog/nand-flash-tlc-qlc-plc-2026).
- Samsung V9 QLC: delayed to H1 2026 amid performance/design issues (ZDNet via TechPowerUp); Samsung's top QLC products were still on older V7 [secondary](https://www.techpowerup.com/341061/samsung-delays-v9-qlc-nand-production-to-h1-2026-amid-performance-issues).
- SK hynix MSC (multi-site cell) split-cell 5-bit: research demo at IEDM Dec 2025 only — not a 2026 product [secondary](https://www.techradar.com/pro/samsungs-biggest-ssd-rival-has-a-clever-plan-to-make-ssd-bigger-but-not-cheaper-split-cell-5-bit-flash-sounds-like-raid-0-on-hard-drives-but-for-ssd).

### 10.2 Layer counts and density (2026 state of the art)

- Samsung V9: ~286–290 layers, mass production since April 2024, 1 Tb TLC; V10 (10th gen): 400+ active layers, 1 Tb TLC, 5.6 GT/s interface, 28 Gb/mm², hybrid bonding + cell-on-periphery (CoP) — ISSCC 2025 demo, production date unannounced [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding).
- SK hynix Gen9: 321 layers (4D PUC), 1 Tb TLC ~20 Gb/mm², mass production; 321-layer QLC 2 Tb die in mass production H1 2026 [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding) [vendor-reported](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Kioxia/SanDisk BiCS8: 218 layers, CBA (CMOS directly Bonded to Array), 2 Tb QLC; BiCS10 (10th gen): 332 layers, CBA + OPS, Toggle DDR6.0 4.8 GB/s, >37 Gb/mm² (QLC) / >29 Gb/mm² (TLC) — FMS 2026 announcement, ~38% denser than Samsung's 400+ layer TLC [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).
- Micron G9: 276 layers, 1 Tb TLC, 21.0 Gb/mm², up to 3,600 MT/s [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding).
- YMTC Xtacking 3.0: 232 layers, 1 Tb QLC, 19.8 Gb/mm² [secondary](https://www.tomshardware.com/pc-components/ssds/samsung-unveils-10th-gen-v-nand-400-layers-5-6-gt-s-and-hybrid-bonding).
- CBA/hybrid bonding separates array and CMOS wafers to avoid thermal degradation of logic under tall stacks — becoming the industry standard beyond ~300 layers [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).
- Density now matters more than raw layer count for hyperscale cost-per-bit; Samsung leads on interface speed (5.6 GT/s) but trails Kioxia/SanDisk on areal density [secondary](https://www.techtimes.com/articles/323160/20260805/kioxia-squeezes-38-more-data-per-chip-samsung-68-fewer-layers.htm).

---

## 11. Endurance, reliability, and data protection

- DWPD classes: 1 DWPD = read-intensive (CDN, boot, read caches); 3 DWPD = mixed-use (OLTP, virtualization, SDS); pSLC 25–60 DWPD = write-cache tier [independent guidance].
- TBW = capacity × DWPD × 365 × warranty years (approximately; vendors publish exact TBW which may use conservative rounding) — e.g. Solidigm 7.68 TB @ 1 DWPD × 5 yr = 14,016 TBW; Micron 9550 MAX 25.6 TB @ 3 DWPD = 140,160 TBW (RND) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146) [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- UBER standard: 1 sector per 10^17 bits read (Samsung, Micron, Kioxia, WD); Solidigm D7-PS1010 rates 1E-18 (10x better) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf) [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/).
- MTBF/MTTF: 2.0M h (Samsung PM9A3, Micron 7450/9550 @ 0–55 °C) to 2.5M h (Kioxia CM7, Solidigm PS1010, WD SN861, Micron 9550 @ 0–50 °C) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf) [secondary](https://www.Scan.co.uk/products/kioxia-64tb-cm7-v-u3-sie-pcie-gen5-1x4-2x2-u3-15mm-mix-use-3dwpd-enterprise-ssd).
- Write amplification factor (WAF): ratio of NAND writes to host writes; lowered by overprovisioning, TRIM, larger sequential writes, ZNS/FDP — enterprise drives ship 7–28%+ OP (Phison X200E 6.4 TB: 28%; Solidigm PS1010: 14.5%) [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html) [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- Power-loss protection (PLP): capacitor-backed flush of in-flight data — standard on all drives in this file (Samsung, Kioxia, Micron, Solidigm, WD, Phison) [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- End-to-end data path protection (DIF/DIX-style), metadata protection, SECDED ECC, LDPC (Marvell NANDEdge 5th gen) — standard enterprise feature set [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- Flash die failure protection / RAID-across-dies: Kioxia CM7 maintains full reliability on die failure; all vendors implement die-level striping [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html).
- Security: TCG Opal 2.0/2.02, TCG Enterprise, TCG Ruby SSCs, SIE (sanitize crypto erase), ISE, AES-XTS 256, FIPS 140-3 L2 certifiable options, SPDM 1.2 device authentication (Micron 9550, SK hynix PEB110) [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf) [vendor-reported](https://www.storagenewsletter.com/2024/10/09/sk-hynix-unveils-peb110-e1-s-gen5-pcie-up-to-8tb-ssd-for-data-centers/).
- Namespaces: 512 (Micron 9550), 132 (Micron 7450), 128 (Phison X200) — multi-namespace support matters for ZNS/SR-IOV/CMB use [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- ZNS (Zoned Namespaces), FDP (Flexible Data Placement), multistream writes, CMB, SR-IOV, SGL: NVMe 2.0 feature set supported across Kioxia CM7, WD SN861, Micron 9550 — reduces WAF and tail latency at the host level [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html).

---

## 12. Prices, availability, and the homelab angle (2026)

### 12.1 New enterprise NVMe pricing (dated snapshots)

