---
id: etape6-phasec-optics-cabling/00-front-matter/1-9-compatibility-coding-warranty-shipping-2026-pricing-note
title: "1.9 Compatibility coding, warranty, shipping, 2026 pricing notes"
domain: front-matter
role: reference
task: pricing
actors: ["Alibaba", "EU", "Nvidia", "United States"]
dates: ["2017-07", "2018-03", "2026-09-22"]
keywords: ["pricing", "datacenter", "ethernet", "nvidia", "optics", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [174, 224]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 5413cf2c65997d1b4b0b6fb6904dc6487db69409150f42aec70d3dd74b8f1af5
---

# 1.9 Compatibility coding, warranty, shipping, 2026 pricing notes

### 1.9 Compatibility coding, warranty, shipping, 2026 pricing notes

- **Compatibility coding:** At purchase, buyer selects "Compatible Brands" / end A & B OEM brand and hardware part number (e.g., Cisco, Arista, Juniper, Dell); FS assigns a distinct SKU per brand coding under a common P/N (e.g., QSFP-LR4-100G: SKUs 24422/36202/36696/48564/75299). FS claims same software codes as the original vendor for major brands, FS's own codes for niche brands [official] (https://www.fs.com/blog/all-about-compatibility-thirdparty-vs-brand-optics-5999.html). FS BOX (V4.0) enables real-time re-coding/config of transceivers, DAC and AOC [official]. HP-compatible modules noted as "Complex Code Makes Higher Price" [official].
- **Warranty:** FS product FAQ states **5-year warranty** on optical modules/transceivers (covers normal-use defects; excludes misuse/abuse/alteration/improper install) [official] (https://www.fs.com/blog/faqs-about-fs-10g-sfp-module-14210.html; https://www.fs.com/sg/blog/faqs-about-fs-1g-sfp-transceivers-7731.html). A **"5-Year Warranty / 30-Day Returns / 30-Day Exchange"** badge appears on most transceiver and DAC/AOC product pages [official]. FS's catalog marketing copy also advertises "lifetime warranty" for transceivers [official] (https://resource.fs.com/mall/file/catalog/optical-module.pdf) — **warranty claim discrepancy: 5-year (product pages/FAQ) vs lifetime (marketing catalog). Treat lifetime claim as [unverified].** Copper patch cords (Cat6A packs) show 1-Year Warranty [official].
- **Shipping/lead times:** "7 Local Warehouses Enabling 90% Same-Day Shipping"; "8 Local Presences Supporting 450K+ Enterprise Clients in Over 200 Countries"; 5x24 phone support [official, sitewide banner]. Free-shipping thresholds: EU "FREE SHIPPING on Orders Over EUR 79 VAT excl."; UK "FREE DELIVERY on Orders over £79 VAT excl."; AU "FREE SHIPPING on Orders Over AUD 79 excl. GST" [official]. US threshold — not captured in these results **[unverified]**.
- **2026 notes/promotions:** No 2026-specific promotions found in the material reviewed; free-product-trial and "Product Custom Request" options appear on select pages (e.g., 100G DAC EU page SKU 65905) [official]. Sold-count data (e.g., 1.4M SFP-10G-SR, 113.8K QSFP-100G-SR4, 14.5K 800G OSFP SR8) suggests core SKUs are stocked items [official].

### 1.10 Regional pricing differences (US vs EU vs UK vs others)

Observed patterns [official, based on indexed page snapshots]:
- FS localizes currency and tax display per region (US $; EU € incl./excl. VAT; UK £ incl./excl. VAT; AU/SG incl. GST). EU/UK ex-VAT prices are generally in the same ballpark as US prices for mainstream optics (e.g., QSFP-SR4-40G: US$43.00 vs €39.00 excl.; SFP-10G-SR: US$25.00 vs €19.00 excl. — **EU slightly cheaper ex-VAT on this SKU**).
- High-end modules: 800G QSFP-DD DR8 Cisco — US$1,429.00 vs €1,458.00 excl. (SKU 304373); 400G DR4 Cisco SiPh — US$749.00 vs €569–599 excl. (EU snapshot variance).
- No evidence of region-exclusive SKUs in the data; pricing deltas are modest but real. **Conflicting snapshots:** the EU QDD-DR4-400G-Si page showed two different prices (€712.81 vs €677.11 incl. VAT) — price changes over time on the live site; current live price should be re-verified.

### 1.11 Explicit gaps / unverified items (Wave 1)

1. **FS.com 800G 2xFR4 own-brand SKU/price:** not found in FS listings reviewed; only third-party equivalents (~US$3.1–3.9K) [secondary]. **[unverified]**
2. **Individual 800G DAC and 800G AOC SKU prices:** categories exist (1.6T/800G DAC/AOC/AEC) but no per-SKU prices captured. **[unverified]**
3. **Individual FS Cat8 patch-cord prices:** product IDs (#72756/#116404) and specs confirmed, but no per-unit prices captured. **[unverified]**
4. **US-listed 400G OSFP DR4 (Ethernet) price:** only InfiniBand NDR variant (US$1,124.00) found; NVIDIA ETH OSFP DR4 price not found. **[unverified]**
5. **400G QSFP-DD active copper (ACC) US price:** only SG (SGD 683.43, SKU 177377) captured. **[unverified]**
6. **Lifetime vs 5-year warranty:** catalog says "lifetime warranty", product pages/FAQ say 5-year — conflict noted, treat lifetime as unverified. **[unverified]**
7. **All prices are from search-index snapshots (crawled 53–361 days ago)**, not a live fetch of fs.com on 2026-09-22 — current prices may differ; the report dates the observation (not the snapshot) as required.

---

## Wave 2 — Optical transceiver landscape & standards (research date 2026-09-22)

### 2.1 Form factors

#### SFP family (all [secondary] unless noted)

- **SFP** — introduced 2001 [secondary]; 1 lane × 1 Gb/s NRZ; LC/RJ45; fiber or twisted pair; MSA-based (not a formal standard) [secondary — qsfptek.com SFP guide]. Source: https://www.qsfptek.com/qt-news/sfp-module-introduction-sfp-meaning-fiber-sfp-and-copper-sfp.html
- **SFP+** — introduced 2009 [secondary — same source]; 1 × 10 Gb/s NRZ (10Gbase); backward compatible with SFP; LC/RJ45 [secondary].
- **SFP28** — introduced 2014 [secondary]; 1 × 25 Gb/s NRZ (28 Gb/s max); targets 25GbE and 32G Fibre Channel; defined under SFF-8402/SFF-8472/SFF-8432 [secondary — lightoptics.co.uk]. Source: https://www.lightoptics.co.uk/blogs/news/a-guide-to-100g-transceiver-qsfp28-sfp56-dd-dsfp-sfp-dd-sfp-dd112-name-and-specs
- **SFP56** — 1 × 50 Gb/s PAM4 (26.5625 GBd); Finisar introduced "industry's first 50G SFP56 transceivers" at OFC in **March 2018** [official — Finisar/GlobeNewswire PR]; electrical standards SFF-8402, SFF-8472, OIF-CEI-04.0 [vendor-reported — FS via ntmmgroup PDF]; use cases: 5G fronthaul, 50G-PAM4 datacenter links. Sources: https://www.globenewswire.com/news-release/2018/03/13/1421253/0/en/Finisar-Introduces-Industry-s-First-400G-QSFP-DD-Active-Optical-Cable-and-Transceivers-for-Switching-and-Routing-Applications-at-OFC-2018.html ; https://www.businesswire.com/news/home/20220307005884/en/ (Source Photonics 50G SFP56, launched Sept 2021 [official]).
- **SFP-DD** ("Small Form-factor Pluggable–Double Density") — SFP-DD MSA formed 12 July 2017, initiated by Alibaba [secondary — fs.com]; 2 electrical lanes (20→40 pins): 2×25G NRZ or 2×56G PAM4 → up to 100/112G aggregate; backward compatible with SFP+/SFP28 in same cage [secondary]. Sources: https://www.fs.com/blog/connectivity-in-the-era-of-big-data-adapting-with-qsfp28-sfpdd-and-dsfp-6591.html ; https://www.lightoptics.co.uk/blogs/news/a-guide-to-100g-transceiver-qsfp28-sfp56-dd-dsfp-sfp-dd-sfp-dd112-name-and-specs
- **SFP-DD112** — 2 × 100G PAM4 → 200G aggregate; SFP112 (single-lane 100G) cages accept SFP+/SFP28 [secondary — same lightoptics source].
- **Use-case pattern** [secondary]: SFP family dominates server/NIC, 5G fronthaul, and enterprise access; SFP-DD targeted high-density 100G ToR server ports.

#### QSFP family (all [vendor-reported] via FS/qsfptek unless noted)
Source: https://www.fs.com/blog/differences-between-qsfpdd-and-qsfp-qsfp28-qsfp56-osfp-cfp8-cobo-1213.html ; https://www.fs.com/blog/how-much-do-you-know-about-qsfp56-1922.html

| Form factor | Cited release year | Lanes | Lane rate/modulation | Aggregate |
|---|---|---|---|---|
| QSFP+ | 2013 per FS; **conflict**: 2012 per QSFPTEK [flagged conflict] | 4 | 10G NRZ | 40G |
| QSFP28 | 2016 | 4 | 25G NRZ | 100G |
| QSFP56 | 2017 | 4 | 50G PAM4 | 200G |
| QSFP112 | 2021 | 4 | 100G PAM4 (400GAUI-4) | 400G |
| QSFP-DD | 2016 (MSA prelim. pub. 19 Sept 2016 [secondary — qsfptek]) | 8 | 25G NRZ / 50G PAM4 / 100G PAM4 | 200G (8×25G) / 400G (8×50G) / 800G (8×100G) |

