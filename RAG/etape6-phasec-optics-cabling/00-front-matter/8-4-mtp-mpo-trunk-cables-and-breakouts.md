---
id: etape6-phasec-optics-cabling/00-front-matter/8-4-mtp-mpo-trunk-cables-and-breakouts
title: "8.4 MTP/MPO trunk cables and breakouts"
domain: front-matter
role: reference
task: reference
actors: ["EU", "Nvidia", "United States"]
dates: ["2025-12", "2026-09-22"]
keywords: ["cost", "gpu", "nvidia", "optics", "pricing"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1304, 1349]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 99aa03673b5cc73117924653dbe3ffe3320b0db8232004793c432953ece3e379
---

# 8.4 MTP/MPO trunk cables and breakouts

### 8.4 MTP/MPO trunk cables and breakouts

FS.com (US store, OFNP plenum default, US Conec MTP, 0.35 dB, G.657.A1 for SM):
- OS2 MTP-12 trunk, SKU 30976, P/N XXMTPSMF, **8-fiber 1 m: US$105.08**; available 8–144 fibers, 1–30 m, female APC [vendor-reported: fs.com/products/30976.html].
- OM4 MTP-12 trunk, SKU 30962, P/N XXMTPOM4, **8-fiber 1 m: US$85.08**; available 8–288 fibers, 1–30 m, female UPC, Type B [vendor-reported: fs.com/products/30962.html].
- OS2 MTP-24 hybrid trunk, SKU 31066, P/N XXMTPXXSMF, **24-fiber 1 m: US$213.08** [vendor-reported: fs.com/products/31066.html].
- OS2 MTP-8 hybrid trunk (MTP-8 to LC uniboot), SKU 209711, **8-fiber 2 m: US$105.25** [vendor-reported: fs.com/products/209711.html].
- OS2 MTP-12 hybrid (MTP-12 to 6× LC duplex uniboot), EU store, SKU 209731, **30 m: €177.31 (€149.00 excl. VAT)** [vendor-reported: fs.com/eu-en].
- MTP-12 → MTP-12 OM4 jumper 1 m, Type B: US$59.00; 5 m Type A: US$116.00 [vendor-reported: FS.com MTP/MPO cables category page].
- MTP-8 APC → 4×LC UPC harness OS2 1 m Type B: US$78.00; OM4: US$59.00; customized MTP-12 harness OS2 8F: US$76.00 [vendor-reported: FS.com category page].

Breakout/conversion cables (FS.com, 1 m unless noted):
- MTP-16 APC → 2×MTP-8 UPC OM4 1 m (crossover): US$150.00; OS2: US$249.00 [vendor-reported: fs.com/c/mtp-mpo-conversions-3089].
- MTP-24 → 2×MTP-12 UPC OM4 1 m Type A: US$166.00; OS2: US$253.00 [vendor-reported: FS.com].
- MTP-12 → 2×MTP-4 APC OM4 1 m Type B (NVIDIA IB variant): US$122.00; OS2: US$150.00; OM3: US$119.00 [vendor-reported: FS.com].
- MTP-16 APC → 2×MTP-8 UPC OM4 1 m: 340 units sold; OS2 variant for 400G DR8 applications [vendor-reported: FS.com product 221891].

Other vendors:
- OptoSpan 48F OS2 MTP(M)-to-LC breakout, 10 m (MBLC-FS248NXR10): US$832.50 [independent: sanspot.com].
- qsfptek 24F MTP(F)-to-12×LC duplex OS2 6 m, Type B, Elite: US$168.78 [independent: qsfptek.com].
- FS EU, 8F OS2 MPO harness, 1 m LSZH, Type A: €42.84 (€36.00 excl. VAT) [vendor-reported: fs.com/eu-en/products/226643.html].
- Scan UK, FS 10 m MTP-12 Elite OM4 trunk (female/female, Type B, OFNP): £137.99 [independent reseller: scan.co.uk].

### 8.5 Cassettes/modules and panels

FS.com (FHD series, 1U enclosure holds 4 cassettes = up to 48 LC fibers; FHD 4U enclosure holds 12 cassettes = up to 432 LC / 3,456 MTP-24 fibers):
- MTP-12 → 6×LC duplex cassette, OS2, Type A, P/N FHD-1MTP6LCDOS2A / SKU 57016: US$110.00, 0.35 dB max [vendor-reported: fs.com/products/57016.html].
- MTP-12 → 6×LC duplex cassette, OM4, Type AF (polarity flipped), SKU 57038: US$119.00 [vendor-reported: fs.com/products/57038.html].
- Enclosures (unloaded): 1U modular FHD (4 cassettes, 144F LC): US$74.00 (SKU 70419 US store; AUD 99.00 on AU store); 1U fixed: US$149.00; 1U sliding: US$289.00; 4U sliding (12 cassettes): US$419.00 (P/N FHD-4UFCE, SKU 73206) [vendor-reported: fs.com category/products pages].
- FHU 1U breakout panel, 96F OM4 (12×MTP-8 male → 24×LC quad): US$1,089.00 (SKU 43514) [vendor-reported: fs.com].
- Accessories: MPO-8/12/24 SC-footprint opposed-key adapter ≤0.35 dB: US$14.00; one-push MTP/MPO-8/12/24 pen cleaner (600+ uses): US$90.00; US Conec IBC Brand Cleaner MPO II (525+ uses): US$99.00 [vendor-reported: FS.com category page].
- OptoSpan 2U enclosure (12 cassettes, 288 LC / 2304 MTP fibers): US$294.25 (EEUX-X00X00-2XT) [independent: sanspot.com].
- Corning EDGE8 reference: 144F OM4 trunk MTP-PC pinned→pinned, Type B; 16F MTP-16 APC non-pinned → 2×8F MTP PC y-harness, Type B (from Corning's Dell 10G–800G cabling guide, LAN-2495-AEN) — official manufacturer SKU examples, pricing not published [official: Corning guide PDF].

### 8.6 Fiber for AI clusters

- Dominant optics (2026): **400G-DR4** (MPO-12 APC, 4×100G PAM4 lanes, OS2, 500 m reach per IEEE 802.3bs) and **800G-DR8** (MPO-16 APC, 8×100G PAM4 lanes, OS2, 500 m reach per IEEE 802.3df-2024). FR4 variants use duplex LC with CWDM4 wavelengths for 2 km; 800G 2×FR4 breaks to two 400G FR4 ports [secondary: roboticsandautomationnews.com 2026-09-22; official: IEEE 802.3bs/802.3df reach values].
  - **Conflict flagged:** one Medium source (December 2025) states "800G DR8 supports up to 100 meters over SMF." This contradicts IEEE 802.3df's 500 m DR8 reach and other sources (fibermall, roboticsandautomationnews); treat the 100-m claim as suspect/outlier [unverified/conflicting].
  - **Conflict flagged:** one source (fibermall.com) lists the 800G DR8 connector as "MTP-12 APC." Native 800G DR8 is MPO-16 APC (8 Tx + 8 Rx); MTP-12 applies to 400G DR4. Treat MTP-12-for-DR8 as an error [secondary conflicting vs. vendor-reported US Conec MTP-16].
- Multimode in AI: 800G SR8 (MPO-16 UPC, VCSEL 850 nm, OM4 up to 50–100 m class) remains the lowest-cost option for intra-rack/adjacent-rack GPU-to-leaf links; 400G SR4 (MPO-12) for legacy 40G→100G-era [secondary: aicplight, fibermall].
- Fiber-counts-per-rack guidance: no single published "N fibers per GPU rack" standard was found in this wave. Vendor guidance converges on: (a) design trunk counts in multiples of 8/16 to match QSFP-DD/OSFP parallel interfaces 1:1 (Base-8/16), (b) provision dual-homing/redundancy per leaf-spine pair, (c) standard 24–144F trunks for intra-row and high-count (96–144F) for inter-row/pod, and (d) MTP PRO field-configurability to avoid stranding polarity/gender inventory [vendor-reported: Corning EDGE8 papers, US Conec, FS.com]. Any specific per-rack number in a downstream document should be tagged [unverified] unless sourced.
- 800G loss budgets are tightening: vendor commentary cites ~3.0 dB channel budgets at 400G shrinking to <2.0 dB (sometimes ~1.5 dB) at 800G with PAM4 lanes, making ultra-low-loss trunk/breakout assemblies near-mandatory for GPU racks [secondary: Medium/beta.holightoptic December 2025 — vendor marketing, treat as directional not normative].
- Breakout complexity at 800G: 2×400G breakout (2×FR4/2×DR4 inside one OSFP) changes MPO polarity mapping requirements; NVIDIA InfiniBand variants use MTP-12 → 2×MTP-4 APC (FS.com stocks this SKU explicitly "For NVIDIA IB") [vendor-reported: FS.com].

### 8.7 Installation and testing

