---
id: etape6-phasec-optics-cabling/00-front-matter/wave-4-fiber-optic-infrastructure-research-date-2026-09-22
title: "Wave 4 — Fiber-optic infrastructure (research date 2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["research", "cost", "dci", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [490, 563]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 7d6b0df4a2284e9ade4602bb7d2d2dc3e1e7da7ebfd1b8a140282b98d2b7d1c7
---

# Wave 4 — Fiber-optic infrastructure (research date 2026-09-22)

## Wave 4 — Fiber-optic infrastructure (research date 2026-09-22)

**Method:** browser_search / browser_open (read-only web research; no files written).

### 4.1 Fiber types

#### 4.1.1 OS2 (G.652.D) — single-mode

**General characteristics [official]:**
- ITU-T G.652.D = standard single-mode fiber (SMF) with low water peak; "OS2" is the ISO/IEC 11801 and EN 50173-1 cabling category for 9/125 µm indoor/outdoor SMF meeting G.652C/D [official — Prysmian spec sheet, https://australia.prysmian.com/sites/australia.prysmian.com/files/media/documents/Tight%20Buffer%20Fibre.pdf]
- ITU-T G.652.D wavelength range: 1310–1625 nm; attenuation ≤0.4 dB/km across the range, ≤0.3 dB/km at 1530–1565 nm [official — ITU G.652; via FS.com comparison http://fs.com/blog/single-mode-fiber-comparison-g652-vs-g655-3213.html]
- Maximum attenuation (cabled fiber): 1310 nm ≤0.35–0.40 dB/km (tight-buffered ≤0.40, loose-tube ≤0.36), 1550 nm ≤0.24–0.40 dB/km (tight ≤0.40, loose ≤0.24), 1383 nm (water peak) ≤0.40 dB/km incl. H₂ ageing [official — OPTOKON datasheet https://nwww.optokon.com/_data/product-files/996/Optical%20Fibres%20Datasheet.pdf; Legrand F01076EN-05 https://assets.legrand.com/pim/NP-FT-GT/F01076EN-05.pdf]
- Mode field diameter 9.2 ± 0.4 µm (1310 nm), 10.4 ± 0.5 µm (1550 nm); cable cutoff λcc ≤ 1260 nm; PMD ≤ 0.2 ps/√km; zero dispersion 1300–1324 nm [official — OPTOKON]
- Typical Panduit OS1/OS2 spec: 1310 nm ≤0.34 dB/km, water peak 1383 nm ≤0.31 dB/km, 1550 nm ≤0.22 dB/km; cladding 125 ± 0.7 µm; NA 0.14 typical [vendor-reported — Panduit TR4 https://panduit.com/content/dam/panduit/en/products/media/0/60/460/7460/106587460.pdf]
- Applications: data-center backbones, campus/inter-building, DCI and WAN; supports CWDM across O/E/S/C/L bands 1260–1625 nm because the water peak is suppressed [secondary — FS.com; beyondtech.us https://beyondtech.us/blogs/beyond-blog/differences-os1-and-os2-fiber]
- Jacket color: yellow [secondary — industry standard color code]

#### 4.1.2 Bend-insensitive OS2: G.657.A1 / G.657.A2

- ITU-T G.657 specifies bending-loss-insensitive SMF; category A is fully/partially compatible with G.652.D; category B need not be [official — ITU-T G.657 (11/2009) https://www.itu.int/rec/dologin_pub.asp?lang=e&id=T-REC-G.657-200911-S!!PDF-E&type=items]
- **G.657.A1:** minimum bend radius 10 mm; macrobending loss ≤0.25 dB (1550 nm, 15 mm radius, 10 turns), ≤0.75 dB (10 mm, 1 turn) [official — ITU G.657 Table 7-1]
- **G.657.A2:** minimum bend radius 7.5 mm; macrobending loss ≤0.03 dB @1550 nm (15 mm, 10 turns), ≤0.1 dB (10 mm, 1 turn), ≤0.5 dB (7.5 mm, 1 turn) [official — ITU G.657 Table 7-1]
- A1 is fully compatible with G.652.D; A2 partially compatible (slightly smaller mode-field diameter, small splice-loss penalty vs G.652.D) [secondary — heyoptics.net https://www.heyoptics.net/blogs/wiki/bend-insensitive-singlemode-fibers-g657; patent CN115542455A https://patents.google.com/patent/CN115542455A/en]
- G.657.A attenuation: ≤0.35 dB/km @1310, ≤0.21 dB/km @1550, ≤0.23 dB/km @1625 [official — Dark Fibre Africa spec https://dfafrica.co.za/documents/procurement/rfp/cto00011121/Annexures/Annexure%20B1%20%E2%80%93%20ENG-SPE-001%20%E2%80%93%20Fibre%20Optic%20Cable%20Specification%20Rev2.2%2020201126.pdf]
- Virtually all data-center patch cords/trunks sold by FS.com, etc. are G.657.A1 (10 mm min. bend) [vendor-reported — FS.com product pages]
- **Bend-insensitive multimode (BIMMF):** available in OM1–OM5; ~10× less loss in tight bends; loops as small as 7.5 mm radius with <0.2 dB bending loss @850 nm and <0.5 dB @1300 nm [vendor-reported — FS.com https://www.fs.com/blog/still-worried-about-bend-radius-come-and-see-the-bendinsensitive-fiber-optic-cables-5112.html]

#### 4.1.3 OM3 / OM4 / OM5 multimode

**Bandwidth (min. specifications):**

| Parameter | OM3 | OM4 | OM5 |
|---|---|---|---|
| Core/cladding | 50/125 µm | 50/125 µm | 50/125 µm |
| OFL bandwidth @850 nm (MHz·km) | 1500 | 3500 | 3500 |
| Effective modal bandwidth (EMB) @850 nm, laser (MHz·km) | 2000 | 4700 | 4700 |
| OFL @1300 nm (MHz·km) | 500 | 500 | 500 |
| EMB @953 nm (MHz·km) | — | — | 2470 |
| OFL @953 nm (MHz·km) | — | — | 1820 |
| Jacket color | Aqua | Aqua / Erika Violet | Lime green |
| Attenuation @850 nm | ≤3.5 dB/km | ≤3.5 dB/km | ≤3.0 dB/km |

[vendor-reported / secondary — FS.com OM5 FAQ https://www.fs.com/blog/om5-fiber-faqs-must-know-for-highspeed-transmission-3505.html; ydfiberoptic.com (updated ~Aug 2026) https://ydfiberoptic.com/what-is-the-difference-between-om3-om4-and-om5/; cablesplususa.com (crawled 2026-09-22) https://store.cablesplususa.com/cabling-insider-blog/multimode-fiber-cabling-types-from-om1-to-om5/]

**Reaches (vendor-reported — FS.com):**

| Transceiver | OM3 | OM4 | OM5 |
|---|---|---|---|
| 10GBASE-SR | 300 m | 550 m | 550 m |
| 40GBASE-SR4 | 100 m | 150 m | 150 m |
| 40G-SWDM4 | 240 m | 350 m | 440 m |
| 100GBASE-SR4 | 70 m | 100 m | 100 m |
| 100G-SWDM4 | 75 m | 100 m | 150 m |
| 400GBASE-SR4.2 | — | — | 150 m |

Note: some vendors list 100 m for OM3 @100G-SR4 and 150 m for OM4/OM5 [secondary — qsfptek.com http://www.qsfptek.com/qt-news/om1-om2-om3-om4-om5-comparision]; **conflict flagged**: 100GBASE-SR4 reach on OM3 is reported as 70 m (FS.com) or 100 m (cablesplususa/qsfptek). IEEE 802.3bm specifies 100 m max over OM3 — treat FS's 70 m figure as conservative/vendor-specific [flagged].

- **OM5 (WBMMF):** specified for 850–950 nm wideband operation; enables SWDM (4 wavelengths: 850/880/910/940 nm, or 850–953 nm range) — 100G over 1 duplex pair (100G-SWDM4: 150 m on OM5 vs 75 m OM3) [vendor-reported — FS.com]
- OM5 costs ~50% more than OM4 (cabling) [vendor-reported — FS.com]
- **Cost pressure:** falling single-mode transceiver costs (silicon photonics, hyperscale volumes) are shifting data centers from MM to SM; FS.com cites 100GBase-PSM4 at $750 supporting 500 m over SM [vendor-reported — FS.com]

#### 4.1.4 OS1a vs OS2 [secondary — beyondtech.us; qsfptek]

- OS1: G.652A/B compliant, **tight-buffered** (indoor), wavelength 1310/1550 nm, cable attenuation up to 1.0 dB/km; limited reach (~2 km at 10G per source — **flagged**: this is a loose, vendor-simplified figure; 10GBASE-LR runs 10 km on OS1/OS1a per standards)
- OS2: G.652C/D compliant (low-water-peak), **loose-tube** (outdoor/underground; widely used indoors in data centers too), attenuation ≤0.4 dB/km (1310–1550 nm), reach up to 10 km @ 10G and hundreds of km with DWDM
- OS1a = indoor-rated variant in ISO/IEC 11801, tight-buffered, commonly rated to ~10 km @ 10GBASE-LX4; OS2 = premises/outside-plant SMF for all longer reaches [secondary — industry convention; **unverified** as a single citable source in this pass]
- Both OS1/OS1a and OS2 use yellow jackets (color does not distinguish them) [secondary]

#### 4.1.5 Standards references
- **ISO/IEC 11801** (incl. 11801-1:2017, which added OM5): defines OS1a/OS2, OM3/OM4/OM5 categories [official]
- **EN 50173-1**: OS2, OS1a categories [official]
- **TIA-568.3-D** (Optical Fiber Cabling and Components) + **TIA-492 series** (fiber specs, e.g., TIA-492-CAAB for OS2, TIA-492-AAAD for OM4) [official]
- **ITU-T G.652 / G.657**, **IEC 60793-2-50** (fiber product specs), **IEC 61754-7 / TIA-604-5 (FOCIS 5)** for MPO interface, **TIA-604-18 (FOCIS 18)** for MTP-16 [official]

