---
id: etape6-phasec-optics-cabling/00-front-matter/8-8-price-data-points-summary-single-unit-list-2026-09-22
title: "8.8 Price data points summary (single-unit list, 2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["Nvidia", "United States"]
dates: ["2026-09-22"]
keywords: ["inference", "nvidia", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1356, 1406]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: a6b1bca6d25b5308d1e6a42b649b027888bd0f9488f37b23874ba48418bdb872
---

# 8.8 Price data points summary (single-unit list, 2026-09-22)

### 8.8 Price data points summary (single-unit list, 2026-09-22)

| Item | SKU/P-N | Price | Source |
|---|---|---|---|
| FS OS2 G.657.A2 LC duplex 3 m | 142870 / SMLCDX | US$7.70 | [vendor-reported: fs.com] |
| FS OS2 G.657.A2 LC duplex 2 m | (blog table) | US$5.30 | [vendor-reported: FS.com blog] |
| FS OM4 LC simplex 1 m | 17508 / OM4XXSX | US$3.90 | [vendor-reported: fs.com] |
| FS OM4 LC armored duplex 3 m | 41028 / AM-OM4LCDX | US$16.00 | [vendor-reported: fs.com] |
| FS OS2 MTP-12 trunk 8F 1 m | 30976 / XXMTPSMF | US$105.08 | [vendor-reported: fs.com] |
| FS OM4 MTP-12 trunk 8F 1 m | 30962 / XXMTPOM4 | US$85.08 | [vendor-reported: fs.com] |
| FS OS2 MTP-24 hybrid trunk 24F 1 m | 31066 / XXMTPXXSMF | US$213.08 | [vendor-reported: fs.com] |
| FS OS2 MTP-8 → LC hybrid 8F 2 m | 209711 | US$105.25 | [vendor-reported: fs.com] |
| FS MTP-16 APC → 2×MTP-8 OM4 1 m | (conversions page) | US$150.00 | [vendor-reported: fs.com] |
| FS MTP-16 → 2×MTP-8 OS2 1 m | (conversions page) | US$249.00 | [vendor-reported: fs.com] |
| FS MTP-24 → 2×MTP-12 OM4 1 m | (conversions page) | US$166.00 | [vendor-reported: fs.com] |
| FS MTP-24 → 2×MTP-12 OS2 1 m | (conversions page) | US$253.00 | [vendor-reported: fs.com] |
| FS MTP-12 → 2×MTP-4 APC OS2 1 m (NVIDIA IB) | (conversions page) | US$150.00 | [vendor-reported: fs.com] |
| FS OS2 MTP-12 → 6×LC cassette | 57016 / FHD-1MTP6LCDOS2A | US$110.00 | [vendor-reported: fs.com] |
| FS OM4 MTP-12 → 6×LC cassette AF | 57038 | US$119.00 | [vendor-reported: fs.com] |
| FS FHD 1U enclosure (4 cassettes) | 70419 | US$74.00 | [vendor-reported: fs.com] |
| FS FHD 4U sliding enclosure | 73206 / FHD-4UFCE | US$419.00 | [vendor-reported: fs.com] |
| FS FHU 1U 96F OM4 breakout panel | 43514 | US$1,089.00 | [vendor-reported: fs.com] |
| OptoSpan 48F OS2 MTP→LC breakout 10 m | MBLC-FS248NXR10 | US$832.50 | [independent: sanspot.com] |
| qsfptek 24F MTP→12×LC OS2 6 m | 102446 | US$168.78 | [independent: qsfptek.com] |
| Scan UK FS 10 m MTP-12 Elite OM4 trunk | LN115934 | £137.99 | [independent: scan.co.uk] |
| Panduit OM4+ LC uniboot duplex 3 m | FS2RPU1U1NNM003 | US$82.00 (MOQ 10) | [independent reseller: hisco.com] |
| Panduit OM4 LC duplex 3 m LSZH | FZ2ELLNLNSNM003 | US$73.78/EA | [independent reseller: lumen.ca] |
| US Conec MTP PRO field tool | — | US$150–166 | [independent: cablesplususa.com] |
| US Conec MTP-16 adapter (SC duplex footprint) | — | from US$4.49–6.47 | [independent: cablesplususa.com] |

**Caveats:** (1) All prices single-unit list; volume/hyperscale pricing will be substantially lower and is not publicly listed. (2) OM3 patch-cord and duplex OM4 SKU-level pricing was not captured — flagged gap. (3) The 100GBASE-SR4 reach conflict (70 m vs 100 m on OM3) is resolved in favor of IEEE 802.3ba (70 m OM3 / 100 m OM4). (4) The 800G DR8 100-m reach claim and the MTP-12 DR8 connector claim from secondary sources are flagged as suspect vs IEEE 802.3df (500 m, MPO-16). (5) "TIA-568.7" as the MPO-polarity standard appears only in a vendor blog; canonical references are TIA-568-C.0 / 568.3-D — tagged [unverified] where that label appears.

---

## Wave 9 — Copper Structured Cabling: Categories, OOB Roles, Pricing, PoE, Testing, Trends

*Research date: 2026-09-22. Single-writer wave. Claim tags: [official] = standards/vendor spec sheets; [vendor-reported] = vendor marketing; [independent] = trade press/analysts; [secondary] = third-party explainers; [unverified] = single/weak source or inference. Price data points are single-unit list prices (Sept 2026), before volume discounts. Tags kept per-claim as researched.*

### 9.1 Twisted-pair categories (bandwidth, rate, distance, shielding)

Category ↔ ISO class mapping: Cat5e↔Class D, Cat6↔Class E, Cat6A↔Class EA, Cat7↔Class F, Cat7A↔Class FA, Cat8↔Class I/II [official — TIA-568 / ISO/IEC 11801 series, confirmed in Fluke DSX datasheet test-standard list].

- **Cat5e** — 100 MHz, 1 Gbps (1000BASE-T) to 100 m [secondary].
- **Cat6** — 250 MHz, 1 Gbps to 100 m; 10GBASE-T to 55 m "in ideal conditions" [secondary]. Lower bound of 37 m cited for noisy/dense-bundle environments [secondary]. CONFLICT: sources quote 37–55 m range; the variance reflects alien-crosstalk conditions, not a single standards number. TIA-568.2-D does not certify Cat6 for 10G by default; short-reach 10GBASE-T on Cat6 requires alien-crosstalk qualification [secondary].
- **Cat6A** — 500 MHz, 10GBASE-T full 100 m channel (90 m permanent link + 10 m patch-cord budget; superseded TIA-568-B capped any single patch cord at 5 m) [official rule via TIA-568 explainer, secondary]. Typical construction 23 AWG solid bare copper; U/UTP and F/UTP/S/FTP variants [vendor-reported — CTi/Casa datasheet 23 AWG, 7.9 mm OD].
- **Cat7** — 600 MHz, ISO/IEC Class F only; NOT recognized by TIA/EIA [secondary — cablinginstall Q&A]. Fully shielded pairs (S/FTP). 10 Gbps to 100 m. Standardization catch: designed for non-RJ45 connectors (GG45/TERA); "Cat7" cables sold with RJ45 ends are effectively connector-limited to Cat6A-class performance [secondary]. Consequence: Cat7 never gained broad commercial-LAN adoption, particularly in North America [secondary].
- **Cat7A** — 1000 MHz, ISO Class FA [secondary]. 10 Gbps to 100 m; higher headroom. Some vendor materials historically positioned it for longer-term 40G readiness, but no IEEE 40GBASE-T application was ever standardized on it — 25/40GBASE-T went to Cat8 [unverified — no source found confirming any ratified >10G BASE-T on Cat7A].
- **Cat8.1 / Class I** — 2000 MHz, 25GBASE-T and 40GBASE-T to 30 m channel (24 m permanent link, 2-connector channel) [official numbers via cablinginstall CommScope Q&A, secondary]. RJ45 interface; backward compatible — 10G and 1G run the full 100 m on Cat8 channel [secondary]. Shielded construction mandatory; most likely 22 AWG S/FTP [secondary]. TIA-568.2-D recognizes only the RJ45 Class I solution; ISO/IEC 11801-1 recognizes both Class I (Cat8.1) and Class II (Cat8.2, non-RJ45 connectors like GG45/TERA) [secondary — cablinginstall Q&A].

**Shielding notation** (ISO/IEC 11801: overall screen / per-pair screen): U/UTP (unshielded), F/UTP (foil overall), U/FTP (foil per pair), F/FTP, S/FTP (braid overall + foil per pair), SF/UTP [secondary]. Shielded designs increasingly common in Cat6A for alien-crosstalk margin [secondary].

