---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/21-vendor-detail-expansions
title: "21. Vendor detail expansions"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["EU", "Samsung"]
dates: []
keywords: ["cost", "datacenter", "inference", "ipo", "liquid cooling", "memory", "nand", "training"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [477, 525]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: d20930fa269f779d350f53f24be4fd466498fc0212155a6928d86da9a10cbc81
---

# 21. Vendor detail expansions

## 21. Vendor detail expansions

- Samsung PM9A3 uses V6 (128L) while 2024-era Samsung client drives moved to V8/V9 — datacenter qualification cycles lag client NAND by ~1 generation, which is normal [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Kioxia CD8P-R 7.68 TB (KCD8XP U G7T68): 12,000/5,500 MB/s, 2,000K/200K IOPS, 1 DWPD, 2.5" U.3 single-port — the single-port counterpart to CM7-R for non-HA racks [secondary](https://www.digitec.ch/en/s1/product/kioxia-cd8p-r-series-kcd8xpug7t68-ssd-7680-gb-ssd-53141468).
- Kioxia CD8P-V 3.2 TB (KCD8XPUG3T20): 12,000/5,300 MB/s, 2,000K/400K IOPS, 3 DWPD — mixed-use single-port [secondary](https://www.galaxus.ch/en/s1/product/kioxia-x121-cd8p-v-dssd-u2-pcie-sie-3200-gb-25-ssd-49654389).
- Kioxia CM7-R 7.68 TB (KCMYRUG7T68): 14,000/5,300 MB/s, 2,700K/160K IOPS, 14,016 TBW @ 1 DWPD, dual-port x2/x2, 25 W, MTBF 2.5M h [secondary](https://www.directdial.com/us/item/kioxia-7-68tb-cm7-r-series-enterprise-2-5-nvme-ssd-solid-state-drive/kcmyxrug7t68).
- Micron 7450 PRO 3.84 TB example (SHI EU): 6,800/5,600 MB/s, 1,000K/400K IOPS, 1 DWPD, 14,000 TBW, U.3 — formula: 3.84 × 365 × 5 = 7,008 vs published 14,000 TBW; discrepancy noted — reseller TBW figure may be for the 7.68 TB SKU; flagged [secondary](https://www.bigw.com.au/product/micron-7450-pro-7-68tb-gen4-nvme-enterprise-ssd-u-3-6800-5600-mb-s-r-w-1000k-180k-iops-25700tbw-1dwpd-2m-hrs-mtbf-server-data-centre-5yrs/p/9902768775).
- Micron 7450 MAX: up to 12.8 TB at 3 DWPD = 70,080 TBW formula vs 70,000 published (rounding) [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- WD SN861 6.4 TB (3 DWPD class): 13,700/7,500 MB/s, 3,300K/800K IOPS, 2.5" U.2, PLP, TCG Opal 2.01 [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca264psp9x1-ssd-data-centre-6-4-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041687/).
- WD SN861 1.92 TB (1 DWPD class): 13,700/3,600 MB/s, 2,100K/165K IOPS — small-capacity write limits visible [secondary](https://www.singular.com.cy/wd-ultrastar-dc-sn861-wus6ba119psp9x3-ssd-data-centre-192-tb-internal-25-u2-pcie-50-x4-nvme.html).
- WD SN861 3.84 TB: 13,700/5,600 MB/s, 2,900K/430K IOPS [secondary](https://www.singular.com.cy/wd-ultrastar-dc-sn861-wus6ba138psp9x3-ssd-data-centre-384-tb-internal-25-u2-pcie-50-x4-nvme.html?sl=el).
- Solidigm D7-PS1010 1.92 TB: 14,500/4,100 MB/s, 1,300K/90K IOPS, 5/15 W — entry SKU is write-limited, typical of channel-count-bound small capacities [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-1-9-tb.d2148).
- Solidigm D7-PS1010 3.84 TB: 14,500/8,200 MB/s, 2,400K/220K IOPS, 12/20 W [secondary](https://www.techpowerup.com/ssd-specs/solidigm-d7-ps1010-7-5-tb.d2146).
- Samsung PM9D3a vs PM9A3: Gen5 vs Gen4, ~1.75x seq-read uplift (12,000 vs 6,900), 2x random-read IOPS (2,000K vs 1,100K) — the standard generational jump [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Samsung datacenter page (2026) lists five lines: PM9D3a, PM9A3, plus BM1743 (QLC high-density), PM9E1, PM963 — page confirms Samsung's QLC datacenter line exists under BM1743 naming [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Phison Pascari X200 family spans X200E (TLC performance), X200Z (pSLC cache), X100E (Gen4 mainstream) per product brochure [vendor-reported](https://www.phisonenterprise.com/wp-content/uploads/2025/01/PascariProductBrochure_X200_010925.pdf).

---

## 22. Homelab buyer's checklist (expanded)

- Endurance remaining: `nvme smart-log` → `percentage_used` (100 = worn out); compare `data_units_written` × 512,000 bytes to rated TBW; <70% used is the usual comfort zone for TLC [independent guidance].
- Error counters: any non-zero `media_errors` or `num_err_log_entries` growth is a reject; `available_spare` must be well above `available_spare_threshold` [independent guidance].
- Power-on hours vs wear: high POH with low `data_units_written` = lightly used (good); low POH with high writes = hammered (bad) — never use POH alone [independent guidance].
- Critical warnings byte: bit 0 (available spare), bit 1 (temperature), bit 2 (device reliability), bit 3 (read-only), bit 4 (volatile memory backup failed = PLP capacitor issue) — reject on bits 3/4 [independent guidance].
- Unsafe shutdowns: high counts suggest missing/failed PLP protection in prior deployment or dirty power — acceptable if SMART is otherwise clean [independent guidance].
- Firmware: check vendor firmware revision against current; Samsung/Kioxia/Micron OEM-branded drives may need OEM update ISOs (Dell/HPE/Lenovo) — plain-channel drives update via vendor tools [independent guidance].
- Sector format: enterprise drives may ship 520/528-byte or 4096-byte sectors — `nvme format --lbaf` to 512/4096 as needed (destroys data) [independent guidance].
- Namespace/SED state: `sedutil-cli` or `nvme` security commands to verify no TCG lock remains; locked drives are bricks without the PSID (printed on the label — photograph it) [independent guidance].
- Physical: U.2 drives need SFF-8643/SFF-8654 (SlimSAS) cables or U.2 backplanes; M.2→U.2 adapters exist but verify PCIe bifurcation support on the motherboard; 7 mm vs 15 mm z-height must match the bay/caddy [independent guidance].
- Thermal: Gen4 U.2 (8–14 W) runs on passive chassis airflow; Gen5 U.2 (19–29 W) wants direct airflow or heatsink kits — throttling starts ~70–75 °C NAND temp [independent guidance].
- Warranty/returns: used enterprise drives are usually sold as-is; factor one spare drive per RAID set into the budget [independent guidance].
- Price sanity: compare $/TB against the 2026 used band ($78–$127/TB for 3.84 TB SATA/SAS enterprise) and new NVMe band ($300–$1,172/TB) — anything far outside needs an explanation (OEM markup, scarcity, or a scam) [secondary](https://pcserverandparts.com/news/enterprise-ssd-prices-2026-server-storage-buying-guide/).

---

## 23. AI workloads and QLC economics

- AI storage tiers show read-to-write ratios up to 20:1+, which matches QLC's asymmetric profile (fast sequential reads, weak random writes) — the core reason hyperscalers absorb QLC despite ~1,000 P/E cycles [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Solidigm D5-P5336 (122.88 TB): up to 9:1 rack-space reduction and ~90% lower storage power vs hybrid HDD/TLC NAS configs; ~25 W peak per drive; optimized for mid-sized I/O patterns in object storage and AI data pipelines [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Kioxia AI angle: CM9 teased explicitly "aimed at AI workloads" with up to 65% random-write and 95% sequential-write gains over CM7 [vendor-reported](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- SK hynix AIN-D family: high-capacity, low-cost, low-power eSSD intended to replace HDDs in AI datacenters; capacity undisclosed; JEDEC NL-SSD standard implied [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Samsung PM9D3a positioning: "built for AI, hyperscale, and cloud" with OCP 2.5 compliance — hyperscale qualification (OCP) is the gating factor for AI-cluster SSD selection [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Liquid cooling arrives in storage: Solidigm's liquid-cooled D7-PS1010 E1 variant with wrap-around cold plate targets AI datacenters where air-cooling dense Gen5 E1.S is marginal [secondary](https://www.tomshardware.com/pc-components/ssds/solidigm-touts-industrys-first-liquid-cooled-enterprise-ssd-d7-ps1010-is-an-e-1-pcie-5-0-drive-with-a-wrap-around-cold-plate).
- Checkpointing workloads (LLM training): bursty sequential writes favor high sustained-write TLC (Micron 9550 MAX 10 GB/s-class, Phison X200) over QLC [independent guidance].
- Inference/serving (model weights, KV caches): read-heavy sequential/random-read — QLC sweet spot; 256 TB-class QLC drives (Sandisk SN670, Kioxia LC9 announced) target exactly this [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).

---

