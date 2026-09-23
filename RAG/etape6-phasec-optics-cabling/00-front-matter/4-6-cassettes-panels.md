---
id: etape6-phasec-optics-cabling/00-front-matter/4-6-cassettes-panels
title: "4.6 Cassettes & panels"
domain: front-matter
role: reference
task: reference
actors: ["Nvidia", "United States"]
dates: []
keywords: ["cost", "distribution", "gpu", "gpus", "nvidia", "optics", "pricing"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [680, 737]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 3688cdd4929e45795303f8b9dd020901e3ebb1b20e842ea7e468dd28f92c65b6
---

# 4.6 Cassettes & panels

### 4.6 Cassettes & panels

#### 4.6.1 FS.com FHD series (US store, crawled 2026)

| Product | P/N / SKU | Price |
|---|---|---|
| FHD MTP-12→6×LC duplex OS2 cassette, Type A, 12F, 0.35 dB | FHD-1MTP6LCDOS2A / 57016 | US$110.00 |
| FHD MTP-12→6×LC duplex OM5 cassette, Type A, 12F, 0.35 dB | FHD-1MTP6LCDOM5A / 174246 | US$139.00 |
| FHD 3×MTP-12→18×LC shuttered, OS2, 36F, Type A | FHD-3MTP18LCDOS2A / 105333 | US$234.00 ex VAT (UK £234) |
| FHD 24F (3×MTP-8→LC quad) OM4 universal polarity | FHD-3MTP6LCQOM4U / 68399 | US$259.00 |
| FHU 1U MTP→LC breakout patch panel 96F OM4 universal | FHU-96FLC12MOM4 / 43514 | US$1,089.00 |
| FHD rack-mount modular enclosure 1U, up to 144F LC (unloaded) | — | US$74.00 |
| FHD rack-mount fixed enclosure 1U, up to 144F | — | US$149.00 |
| FHD sliding-drawer enclosure 1U | — | US$289.00 |

Sources: https://www.FS.COM/products/57016.html, https://www.fs.com/products/174246.html, https://www.fs.com/uk/products/105333.html, https://www.fs.com/products/68399.html, https://www.FS.COM/products/43514.html, https://www.fs.com/c/rack-mount-enclosures-45

- 4 FHD cassettes per 1U → up to 144F (36F cassettes) or 288F with MTP-24 panels [vendor-reported — FS.com]

#### 4.6.2 Corning EDGE / EDGE-16
- EDGE HD housings: 1U = 144F (EDGE-01U-SP), 2U = 288F, 4U = 576F; EDGE trunks 12–144F [official — Corning LAN-1405 https://www.corning.com/catalog/coc/documents/line-card/LAN-1405-A4-BEN.pdf]
- Distributor list prices (amerisponse.com, crawled 2026-09): EDGE 12F module LC-duplex→MTP OS2 (ECM-UM12-05-93Q) $506.17; OM4 version (ECM-UM12-05-93Q? listed as ECMUM120593Q) $507.94; 1U housing EDGE-01U $566.38; 24F MTP panel $86.41 [vendor-reported — https://www.amerisponse.com/Corning/Corning-Pricelist-P4.html]
- **EDGE-16** (16-fiber modules for 400G-SR8/DR8): exists in Corning portfolio; **no current SKU/pricing captured in this pass — flagged as a gap** (Corning EDGE-16 modules use 16F MT ferrules per FOCIS-18)
- CommScope (NVIDIA NVL72 ordering guide) uses MPO-12/MPO-16 trunks with MPO-16 APC pinned backbones, modules like PPL-DM-16AU-16LC-SM-BEU (16F MPO→16 LC) and LC-uniboot patch cords; part numbers: MMS4X50-NM (400G DR4 OS2 500 m), URXQVQVF8 (MPO-12/8 cable) etc. [vendor-reported — https://webresources.commscope.com/download/assets/CO-120443-EN_NVIDIA+NVL72+DC+Ordering+Guide.pdf/e6cd640a10b611f1944a3a47d01923dd]

#### 4.6.3 Panduit HD Flex
- Enclosures: FLEX1U06 (1U, 12 cassettes, **144F**), FLEX2U06 (288F), FLEX4U06 (576F); convertible 4/6/12-port cassettes [official — Panduit FBSP92 https://objects.eanixter.com/PD442248.PDF]
- With MPO adapter panels: up to 864F/1U (12F MPO FAPs) or 1728F/1U (24F MPO FAPs) [official — Panduit]
- Cassette example: FHC39N-08H-10A — HD Flex 4-to-1 breakout cassette, OS2, standard IL, Method A [official — Panduit solution guide https://www.panduit.com/content/dam/panduit/en/corporate-non-product/documents/cpcb223-sa-eng-multi-tenant-data-center-solution-guide/multi-tenant-data-center-solution-guide-web-cpcb223-sa-eng.pdf]
- No current retail pricing found for Panduit HD Flex — **flagged gap** (Panduit sells through distribution; pricing is quote-based)

### 4.7 Fiber for AI clusters

- **Why OS2 dominates AI scale-out fabrics:** (1) reach — 800G-DR8 OS2 runs 500 m (2 km DR8-2) vs 100 m MM limit [official — CommScope/802.3df table]; (2) single-mode transceiver cost collapse (silicon photonics, hyperscale volumes); (3) leaf-spine fabrics span buildings/rows beyond MM reach; (4) future-proofing for 1.6T (200G/lane) where MM budgets are tightest [secondary — synthesis of CommScope ordering guide + holightoptic 400G/800G article]
- **MPO-16 adoption for 800G:** 800G-SR8/DR8 = 8 Tx + 8 Rx over one MPO-16/APC; dual-MPO-12 (2×DR4/2×SR4) is the OSFP twin-port alternative [vendor-reported — CommScope; medium/@aicplight888]
- **NVIDIA DGX B300 reference design (NADDOD, crawled 2026):** rail-optimized fat tree; 8 scalable units, 576 nodes / 4,608 GPUs; Quantum-X800 Q3400 IB switches; DGX↔Leaf and Leaf↔Spine use 1.6T 2×DR4 OSFP224 XDR transceivers with **MPO-12 APC OS2 jumpers** (~9,216 MPO cables per 8-SU build) [secondary — https://www.naddod.com/ai-insights/nvidia-dgx-b300-cluster-cabling-guide-infiniband-and-roce-network-deployment]
- **CommScope Propel for NVIDIA NVL72:** MPO-16 APC pinned trunks for serial-to-serial 800G/400G; breakout modules split 800G→2×400G or 4×100G [vendor-reported — CommScope guide + tde 800G breakout module https://www.cablinginstall.com/connectivity/product/55362193/800g-breakout-module-with-front-facing-mpo-port]
- **Loss budgets at 800G:** 400G links tolerate ~3.0 dB; 800G DR8 links often <2.0 dB, sometimes ~1.5 dB — drives low-loss/Elite MPO and fewer mated pairs [secondary — medium/@beta.holightoptic; treat as vendor/editorial guidance, not a standard]
- **Bend-insensitive fiber in dense racks:** G.657.A1 (10 mm radius) is the de-facto patch-cord fiber; congestion in 40 kW+ GPU cabinets makes tight-bend tolerance and thin uniboot/VSFF jumpers (SN/MDC/CS) essential for airflow and MACs [secondary — gcabling.com https://www.gcabling.com/ai-server-network-infrastructure-fiber-rack-guide/; FS.com]

### 4.8 Testing & standards

#### 4.8.1 Link loss budgets
- Budget = cable loss (length × dB/km) + connector-pair loss + splice loss [secondary — FOA; telecron.net poster]
- TIA-568 default allowances: **0.75 dB per mated connector pair, 0.30 dB per splice** [secondary — cablinginstall.com https://www.cablinginstall.com/testing/article/16467189/fiber-optic-testing-keeping-it-simple-in-the-field]
- MPO reality: Elite/low-loss MPO (0.35 dB max) is required to stay within 800G's <2 dB channel budgets across multiple mated pairs [secondary synthesis]
- Channel examples: 800G-DR8 total channel insertion loss per IEEE 802.3df is 4.0 dB max (500 m); vendor guidance warns practical AI-cluster budgets are far tighter [official — IEEE 802.3df via CommScope table; vendor note]

#### 4.8.2 Tier 1 vs Tier 2 certification
- **Tier 1 (mandatory per TIA-568-C.0):** OLTS (light source + power meter) measures end-to-end insertion loss, length, and polarity. One-jumper reference method is TIA-preferred; ISO prefers the three-jumper method [secondary — cablinginstall.com; trend-networks.com]
- **Tier 2 (supplemental):** OTDR characterizes each component (connectors, splices, bends) with distance-to-event; requires Tier 1 first; OTDR-derived loss is calculated, not directly measured — OLTS remains the more accurate IL measurement [secondary — trend-networks.com https://www.trend-networks.com/wp-content/uploads/2019/05/Tier-1-vs-Tier-2-Fibre-Optic-Certification-Part-1.pdf]
- Test wavelengths: MM @ 850 nm (opt. 1300 nm); SM @ 1310 nm (opt. 1550 nm) [secondary — FOA]

#### 4.8.3 Cleaning & inspection
- **IEC 61300-3-35** defines pass/fail criteria for connector end-face visual inspection (scratches/digs/defects by zone); inspection scopes (e.g., Fluke FI-1000, EXFO) give automated pass/fail to this standard [secondary — cablinginstall.com https://www.cablinginstall.com/testing/article/16465842/tier-1-and-tier-2-testing-troubleshooting-and-documentation]
- MPO/MTP inspection: 12-fiber ferrules must be inspected all-fibers; contamination is the #1 cause of high IL in MPO systems — "inspect before you connect" discipline is mandatory, especially for APC single-mode parallel optics [vendor-reported — holightoptic.com]
- MPO polarity field check: VFL into position 1 → position 1 lights = Type A; position 2 = Type C; position 12 = Type B [secondary — medium/@Yanissss]

