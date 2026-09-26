---
id: etape6-phasec-optics-cabling/00-front-matter/7-8-gaps-and-caveats-wave-7
title: "7.8 Gaps and caveats (Wave 7)"
domain: front-matter
role: reference
task: reference
actors: ["China", "EU", "United States"]
dates: ["2026-06", "2026-09-22"]
keywords: ["datacenter", "optics", "pricing", "research", "tpu", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1244, 1286]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: bcd1fcc01434a7c7a6005853c3c08b7a65928e99aa399a1f08a2dd1a901b0926
---

# 7.8 Gaps and caveats (Wave 7)

- **Availability**: "7 Local Warehouses Enabling 90% Same-Day Shipping" [official]; "Full inventories are provided in FS local warehouses in the US, AU, Europe, and Asia covering global markets and supporting same-day shipping" [official]; US HQ 380 Centerpoint Blvd, New Castle, Delaware 19720 [official]. "8 Local Presences Supporting 450K+ Enterprise Clients in Over 200 Countries" [official]. "FREE SHIPPING on Orders Over US$79" (US store banner) [official].
- **Support**: "900+ Technical Support and Service Professionals", "0.5-Day Average Solution Response Time", "5x24 Phone Support" [official]; "600+ R&D Experts" [official].
- **Warranty**: FS materials are inconsistent across products — FS 40G optics portfolio PDF states "Lifetime Warranty" [official]; an FS DAC product page lists "5-Year Warranty, 30-Day Returns, 30-Day Exchange" [official]; a third-party 400G AOC listing sourced from FS specs states 5 years [secondary]; eBay listing of FS OS2 patch cord states "Manufacturer Warranty: Lifetime" [secondary]. Net: FS advertises lifetime warranty on optical modules/patch cabling lines, with some DAC/AOC SKUs showing 5-year terms — verify per-SKU at purchase [official/secondary].

### 7.8 Gaps and caveats (Wave 7)
- [unverified gaps]: 400G straight-through DAC pricing; 200G AOC and 50G/56G DAC/AOC SKU-level pricing; 400G LR4/coherent pricing; 1.6T transceiver list price; FS cassette/enclosure pricing; QSFP-DD 800G AOC beyond the single OSFP AOC data point.
- All fs.com prices are public list prices (often with OEM-compatible coding variants); no volume/discount pricing captured. Prices scraped from category pages crawled 53–343 days ago may have drifted; EU prices include VAT at country rates.

---

## Wave 8 — Fiber-Optic Infrastructure: OS2/OM3/OM4/OM5, MPO/MTP, Polarity, Trunks, Cassettes, Testing

*Research date: 2026-09-22. Single-writer wave. Price caveat: all prices are single-unit list prices as scraped from vendor storefronts (USD unless noted); not comparable to volume pricing; web-crawl dates vary. No prices or SKUs invented. Tags: [official] (standards bodies/vendor spec sheets), [vendor-reported], [independent], [secondary], [unverified].*

### 8.1 Fiber types — OS2 single-mode, OM3/OM4/OM5 multimode

**OS2 single-mode (G.652.D, G.657.A1/A2).** OS2 = 9/125 µm single-mode fiber. ITU-T G.652.D is the baseline standard; G.657.A1 and G.657.A2 are bend-insensitive variants fully compliant with G.652.D, allowing 10 mm (A1) and 7.5 mm (A2) minimum bend radii vs. 30 mm for standard G.652.D [vendor-reported: FS.com blog G.652D-vs-G.657A1-vs-G.657A2; official-standard: ITU-T G.657]. G.657.A1/A2 are compatible with standard single-mode fiber and are the default fiber type in datacenter patch cords because of rack-space/tight-routing benefits [vendor-reported: FS.com; plugsters.com]. Attenuation: 1310 nm ≤0.35–0.36 dB/km; 1550 nm ≤0.21–0.25 dB/km depending on manufacturer (FS lists G.657.A1 at 0.36/0.22 dB/km, G.657.A2 at 0.40/0.25 dB/km for its cables) [vendor-reported: FS.com]. Zero-dispersion wavelength 1300–1324 nm, PMD ≤0.2 ps/√km [vendor-reported: fiberfuture]. OS2 supports 10 km at 1310 nm (LR) and up to 40 km at 1550 nm (ER-class transceivers) [vendor-reported: plugsters.com]. Jacket color: yellow [official: TIA-598 color code].

**OM3/OM4/OM5 multimode (50/125 µm, laser-optimized).**
| Type | OFL bandwidth @850nm | Effective modal bandwidth (EMB) @850nm | 953nm EMB | Attenuation | Jacket |
|---|---|---|---|---|---|
| OM3 | 1500 MHz·km | 2000 MHz·km | n/a | ~3.5 dB/km | Aqua |
| OM4 | 3500 MHz·km | 4700 MHz·km | n/a | ~3.5 dB/km | Aqua/Violet |
| OM5 | 3500 MHz·km | 4700 MHz·km | 2470 MHz·km | ~3.0 dB/km | Lime green |
[secondary: ydfiberoptic.com, cloudswit.ch; vendor-reported: FS.com OM5 FAQ, qsfptek store]. OM5 is identical to OM4 at 850 nm but adds the 850–953 nm window for SWDM (Shortwave Wavelength Division Multiplexing) — 4 wavelengths per fiber — and is backward compatible (same 50 µm core) [vendor-reported: FS.com]. OM5 typically costs ~50% more than OM4 [vendor-reported: FS.com OM5 FAQ].

**Distances.** 10GBASE-SR: OM3 300 m / OM4 550 m. 40GBASE-SR4: OM3 100 m / OM4 150 m. 100GBASE-SR4: OM3 70 m / OM4 100 m (IEEE 802.3ba) [secondary; official-standard: IEEE 802.3ba reach values].
- **Conflict flagged:** some vendor content lists "100GBASE-SR4 @ 100 m on OM3" or conflates 40G-SR4 (100/150 m) with 100G-SR4 (70/100 m). The correct IEEE values per IEEE 802.3ba are 100GBASE-SR4: 70 m OM3 / 100 m OM4; 40GBASE-SR4: 100 m OM3 / 150 m OM4 [official: IEEE 802.3ba; conflicting secondary sources flagged — do not use the 70-m figure for 40G or the 100-m figure for OM3-100G interchangeably].

**OS2 vs multimode for datacenters.** OS2 is the choice for 400G-DR4/FR4 and 800G-DR8/FR8 (parallel and WDM single-mode). OM3/OM4 remain for short-reach SR4/SR8 links. Because silicon-photonics single-mode optics prices have fallen with hyperscale volumes, single-mode is displacing OM4 for new high-speed builds [vendor-reported: FS.com OM5 FAQ; secondary: Medium/aicplight June 2026].

**Patch-cord prices (LC-LC duplex).**
- FS.com, 3 m OS2 G.657.A2 LC UPC duplex, OFNR riser, P/N SMLCDX / SKU 142870: US$7.70 [vendor-reported: fs.com/products/142870.html].
- FS.com Europe, 7 m OS2 G.657.A1 LC UPC duplex, OFNR: €7.50 (€6.30 excl. VAT) [vendor-reported: fs.com/eu-en/products/40201.html].
- FS.com, 2 m OS2 G.657.A1 LC UPC duplex: US$4.70; G.657.A2: US$5.30 (simplex A1 US$2.80 / A2 US$2.90) [vendor-reported: FS.com blog table].
- FS.com, OM4 LC-LC simplex 1 m, OFNR, P/N OM4XXSX / SKU 17508: US$3.90 (1 m; lengths 1–30 m) [vendor-reported: fs.com/products/17508.html]. Multimode duplex pricing for OM4 was not captured at SKU level in this wave (simplex data point only).
- FS.com, 3 m OM4 LC-UPC armored duplex, P/N AM-OM4LCDX / SKU 41028: US$16.00 [vendor-reported: fs.com/products/41028.html].
- FS.com industrial armored OS2 LC-LC duplex TPU: 10 m US$47.00, 30 m US$102.00 (SKU 106589/106591) [vendor-reported: fs.com].
- Contrast: Panduit OM4+ Signature Core LC Uniboot duplex 3 m plenum (FS2RPU1U1NNM003): US$82.00 each list (US$97.72 before discount), MOQ 10 [independent reseller: hisco.com]; Panduit Opti-Core OM4 duplex 3 m LSZH (FZ2ELLNLNSNM003): US$73.78/EA [independent reseller: lumen.ca]. Branded-plant pricing is ~10–20× FS.com list for comparable MM patch cords.
- China-factory spot pricing (OEM): LC-LC OS2 duplex 3.0 mm LSZH at US$0.01–1.99/pc, MOQ 100 pcs [vendor-reported: ecer.com] — volume-tier, not comparable to single-unit retail.

### 8.2 MPO/MTP connectors

