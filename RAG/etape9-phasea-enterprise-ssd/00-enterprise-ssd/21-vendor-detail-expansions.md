---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/21-vendor-detail-expansions
title: "21. Vendor detail expansions"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["EU", "Samsung"]
dates: []
keywords: ["datacenter", "nand"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [462, 497]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 87a52010af18eb6a59a4dbf732060237add0a80bda9f5b8e556726e54e728675
---

# 21. Vendor detail expansions

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

