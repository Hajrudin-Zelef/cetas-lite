---
id: etape6-phasec-optics-cabling/00-front-matter/4-4-mtp-pro-us-conec
title: "4.4 MTP PRO / US Conec"
domain: front-matter
role: reference
task: reference
actors: ["EU", "Nvidia", "United States"]
dates: []
keywords: ["datacenter", "inference", "nvidia", "pricing"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [623, 679]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: eb17f01d28dfe4d59d7c0e61515ac1c58a84a555c88090756b47d8619120a6a5
---

# 4.4 MTP PRO / US Conec

### 4.4 MTP PRO / US Conec

- **MTP® PRO** (released 2017 per FS timeline): field-reversible connector — with a small handheld tool, an installer can change **gender (pins in/out)** and **polarity (key up ↔ key down)** on the connector **without removing the housing** [vendor-reported — Black Box https://www.blackbox.com/en-my/insights/blogs/detail/tps/2022/08/19/the-evolution-of-the-mpo-connector-from-its-original-design-to-what-it-is-today; cablinginstall.com https://www.cablinginstall.com/home/article/16468532/mpo-connectors-and-related-tools]
- Mechanism: one insertion into a "polarity change port" retracts the exposed key and extracts a hidden key on the opposite side; pin-keeper design allows pins to be added/removed repeatedly without handling loose pins [vendor-reported — US Conec product catalog https://www.usconec-cn.com/media/k2pbdmyl/us-conec-product-catalog.pdf; AFL datasheet https://afl-delivery.stylelabs.cloud/api/public/content/5329564cc0af4478854385453f077c7a?v=d0b5e66f]
- MTP PRO pins meet IEC 61754-7 pin-retention force (≥19.6 N) [vendor-reported — AFL]
- **MTP® PRO X**: optimized for APC single-mode applications; field gender change (no polarity change); DirectConec™ push-pull system [vendor-reported — US Conec]
- **MTP® PRO vs standard MPO:** same IEC 61754-7/TIA-604-5 optical interface (interoperable), but PRO adds field configurability; standard MPO gender/polarity is fixed at manufacture [vendor-reported]
- Adoption: used to reduce trunk SKU variety (one trunk serves Type A or B, male or female) and re-purpose expensive pre-terminated trunks in MACs; Panduit's pre-terminated PanMPO trunks offer similar field gender/polarity change [vendor-reported — Black Box; Panduit https://maincdn.getuniqcli.com/products/panduit-corp/flex-rcm2u/docs/2.pdf]
- Hyperscale note: hyperscalers increasingly use **spliced/field-terminated or direct trunk-to-transceiver** designs with minimal cassette layers; MTP PRO's main value is in enterprise/colocation where trunks are reused across migrations [secondary — inference from vendor positioning; **unverified** as a hyperscale-specific adoption figure]
- Field tool price example: MTP® PRO polarity/pin removal tool listed at $150.00 (retail $166.00) at CablesPlusUSA [vendor-reported, date crawled 2026 — https://store.cablesplususa.com/manufacturers/shop-by-brand/us-conec/]

### 4.5 Patch cords & trunks

#### 4.5.1 Duplex LC patch cords (prices, FS.com, US store unless noted)

| Product | P/N / SKU | Price | Source/date |
|---|---|---|---|
| 2 m LC-LC OS2 uniboot BIF | HD-SMULCDX / SKU 68484 | €15.47 (€13.00 ex VAT) | FS.com EU, crawled 2026-09 https://www.fs.com/eu-en/products/68484.html |
| 2 m SN→LC UPC OS2 uniboot OFNR 1.6 mm | HD-SMUSNFULCDX / SKU 155088 | US$32.00 | FS.com US https://www.fs.com/products/155088.html |
| 2 m MDC→LC UPC OS2 uniboot OFNR | HD-SMUMDCLCDX / SKU 139411 | US$41.00 | FS.com US https://www.fs.com/products/139411.html |
| 3 m CS→LC UPC OS2 uniboot OFNR | HD-SMUCSULCDX / SKU 116822 | US$28.00 | FS.com US https://www.fs.com/products/116822.html |
| 1 m MDC→LC UPC OS2 uniboot OFNR | HD-SMUMDCLCDX / SKU 130968 | US$41.00 | FS.com US https://www.fs.com/products/130968.html |
| 3 m Cleerline LC-LC OS2 uniboot OFNR | SKU 149171 | US$31.30 | PacRad https://www.pacrad.com/cleerline-ssf-lc-upc-lc-upc-2-0mm-os2-riser-fiber-uniboot-patch-cord-3m.html |
| 3 m Belden OS2 LC duplex UHD uniboot OFNR | MPN FPSLULU003M | US$60.27 (retail $66.30) | alldataresource.com https://www.alldataresource.com/Belden-FPSLULU003M-Fx-Patch-Cord-Os2-Lc-Duplexuhd-Uniboot-lc-Duplex-Uhduniboot-3m-Ofnr_p_635260.html |

- Uniboot = both fibers in one round cable (2.0 mm), halves patch-field bulk, reversible polarity without tools [vendor-reported — FS.com]
- VSFF connectors (SN, MDC, CS): 1.5×–3× density vs LC duplex, target 200/400/800G front panels (QSFP-DD/OSFP) [vendor-reported — FS.com]
- All OS2 patch cords above use G.657.A1 (10 mm min. bend) [vendor-reported — FS.com]

#### 4.5.2 MPO trunks

| Product | P/N / SKU | Price | Source/date |
|---|---|---|---|
| Custom 24–144F OS2 MTP-12 trunk, 0.35 dB, OFNP, female APC (1 m, 8F config shown) | XXMTPSMF / SKU 30976 | US$105.08 | FS.com US https://www.fs.com/products/30976.html |
| Custom 48–144F OS2 MTP-24 trunk, 0.35 dB, OFNP (1 m, 24F shown) | XXMTPXXSMF / SKU 31066 | US$213.08 | FS.com US https://www.fs.com/products/31066.html |
| Custom MPO-12 harness 8–144F OS2 0.35 dB (1 m) | XXMPOXXSMF / SKU 226643 | €42.84 (€36.00 ex VAT) | FS.com EU https://www.fs.com/eu-en/products/226643.html |
| MPO-12(F)→MPO-12(F) OS2 OFNP Type B (Takfly) | — | US$80.70 | fiberopticsupply.com https://fiberopticsupply.com/mpo-12-female-to-mpo-12-female-os2-sm-trunk-cable-ofnp-type-b/ |

- Trunk ranges: 8–144F common; 288F available (FS offers up to 288F MPO-12 assemblies) [vendor-reported — FS.com]

#### 4.5.3 Breakout / harness cables

| Product | Price | Source/date |
|---|---|---|
| MTP-12 → 2×MTP-4 APC 8F OM4 Type B 1 m ("For NVIDIA IB") | US$122.00 | FS.com https://www.fs.com/c/mtp-mpo-conversions-3089 |
| MTP-12 → 2×MTP-4 APC 8F OS2 Type B 1 m ("For NVIDIA IB") | US$150.00 | FS.com (same page) |
| MTP-16 → 2×MTP-8 APC 16F OS2 crossover 1 m | US$249.00 | FS.com (same page) |
| MTP-24 → 3×MTP-8 APC 24F OS2 Type B 1 m | US$307.00 | FS.com (same page) |
| MTP-24 → 3×MTP-8 UPC 24F OM4 Type B 1 m | US$207.00 | FS.com (same page) |
| MTP-16 → 2×MTP-8 UPC 16F OM4 crossover 1 m | US$150.00 | FS.com (same page) |
| OS2 MPO→4×LC duplex harness 1 m (Patchsave, UK) | £41.12 ex VAT | patchsavesolutions.com https://patchsavesolutions.com/os2-mpo-assemblies-and-fan-outs/9097-os2-mpo-to-4x-lc-duplex-breakout-harness-fan-out-assembly.html |

- Standard configs: MPO-12 → 6×LC duplex; MPO-24 → 12×LC duplex; MPO-16 → 2×MPO-8; MPO-24 → 3×MPO-8; MPO-12 → 2×MTP-4 (8-fiber twin-port for 800G→4×200G NVIDIA IB, e.g., Quantum-2, ConnectX-7) [vendor-reported — FS.com; LinkedIn FS post https://www.linkedin.com/posts/fscomglobal_infiniband-datacenter-cabling-activity-7361071589825695744-Xq2w]
- **Armored fiber:** stainless-steel-armored patch cords/breakouts are offered by vendors (e.g., ecer.com lists armored MPO breakout assemblies) but **no current retail pricing was found in this pass — flagged as a gap**; treat any figure as unverified.

**Price comparability warning:** FS.com prices vary by region (US/EU/SG/AU), length (1 m shown above), jacket (OFNP > LSZH > OFNR) and IL grade; Corning/Panduit list prices run 3–5× higher than FS street prices (see §4.6) — non-comparable without normalizing length/config/discount.

