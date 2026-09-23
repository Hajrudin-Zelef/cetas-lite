---
id: etape6-phasec-optics-cabling/00-front-matter/1-9-compatibility-coding-warranty-shipping-2026-pricing-note
title: "1.9 Compatibility coding, warranty, shipping, 2026 pricing notes"
domain: front-matter
role: reference
task: pricing
actors: ["Alibaba", "EU", "Nvidia", "United States"]
dates: ["2017-07", "2018-03", "2019-07", "2021-05", "2024-06", "2026-09-22"]
keywords: ["pricing", "asic", "datacenter", "ethernet", "gpu", "liquid cooling", "nvidia", "optics", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [174, 237]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 9e32bf8b813d8cae197a63a08e447ddb8871a5f6dc294d78ed3ad4de5f6c8f38
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

- QSFP-DD backward compatible with QSFP+, QSFP28, QSFP56, QSFP112 (DD cage accepts 4-lane modules) [vendor-reported]. Note: inserting QSFP112 into QSFP-DD cage requires host ASIC support for 4×100G PAM4 [vendor-reported].
- **QSFP-DD MSA hardware spec revisions** [secondary — DKN Consulting PDF + fibermall]: Rev 5.0 July 2019; Rev 5.1 Aug 2020; **QSFP-DD800 (8×100G) released 24 May 2021** (qsfptek [secondary]); **Rev 7.0 announced at ECOC 2023** defines QSFP-DD1600 (8×200G PAM4, 1.6T, backward compat with DD800/QSFP56/QSFP28, CMIS 5.x, typical power 20–28 W, riding heatsink); **Rev 7.1 released June 2024** (refines module/connector definitions) [secondary — fibermall.com/blog/1-6t-qsfp-roadmap-qsfp-dd1600-guide.htm].
- Power notes (see also §2.6): early QSFP-DD thermal class 7–12 W [secondary]; QSFP-DD800 modules rated up to 18 W (vendor-reported FS 800G guide); QSFP-DD1600 typical 20–28 W [secondary].

#### OSFP family
- **OSFP** — initially designed 2016 for 8 × 50G = 400G [secondary — qsfptek]; dimensions 22.58 × 107.8 × 13.0 mm vs QSFP-DD 18.35 × 89.4 × 8.5 mm [vendor-reported]; 36 ports per 1U = 14.4 Tbps at 400G [vendor-reported]. Not backward compatible with QSFP+/QSFP28 (larger cage) [vendor-reported].
- **OSFP 4.0** released 2021 — 800G via 8 × 100G [secondary — qsfptek].
- **OSFP MSA Revision 5.1** defines implementations for 50G-, 100G-, and 200G-class signaling per electrical lane → supports 1.6T-class modules (8 × 200G) [secondary — roboticsandautomationnews.com, 22 Sept 2026]. Source: https://roboticsandautomationnews.com/2026/09/22/osfp-modules-the-complete-guide-to-400g-800g-and-1-6t-optical-transceivers-for-ai-and-hyperscale-data-centers/104982/
- Power class: early OSFP 12–15 W [secondary]; 800G OSFP modules rated up to 18 W in vendor tables (see §2.6) [vendor-reported]. 36-port 1U panel delivers 28.8 Tbps at 800G/port [secondary].
- Three thermal tops: finned-top (open top, air-cooled), closed finned-top (air + liquid), flat top (cold-plate liquid cooling) [vendor-reported — FS]. Flat-top OSFP used with NVIDIA DGX H100 Cedar7 GPU links [vendor-reported].
- **OSFP-XD** (eXtra Dense) — MSA-defined 1.6T variant: **16 electrical lanes** (16×100G = 1.6T; future path 16×200G = 3.2T), same cage/panel cutout/thermal as OSFP, up to ~40W design target, up to 32 ports per 1RU / 64 per 2U, supports 100GBASE-CR1 DAC [secondary — Medium/aicplight, 16 Sept 2026; Cosolvic blog]. Sources: https://medium.com/@aicplight888/1-6t-optical-transceiver-form-factor-comparison-osfp1600-vs-osfp-xd-7bc0dcc84670 ; https://cosolvic.com/blog/800g-data-center-connectors-osfp-vs-qsfp-dd-comparison/
- Use-case pattern: OSFP strong in AI/HPC and InfiniBand/NDR (e.g., 800G OSFP DR8 for IB NDR fabrics) in addition to Ethernet; QSFP-DD dominant in Ethernet datacenter [vendor-reported].

