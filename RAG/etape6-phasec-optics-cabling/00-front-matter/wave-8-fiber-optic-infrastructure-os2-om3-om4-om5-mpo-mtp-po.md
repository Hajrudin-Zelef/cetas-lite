---
id: etape6-phasec-optics-cabling/00-front-matter/wave-8-fiber-optic-infrastructure-os2-om3-om4-om5-mpo-mtp-po
title: "Wave 8 — Fiber-Optic Infrastructure: OS2/OM3/OM4/OM5, MPO/MTP, Polarity, Trunks, Cassettes, Testing"
domain: front-matter
role: reference
task: reference
actors: ["China", "United States"]
dates: ["2026-06", "2026-09-22"]
keywords: ["datacenter", "optics", "pricing", "research", "tpu", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1254, 1303]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: e75e89bc069e8b49de099d0dfddfebd0d7f1d3210106a1fdbd274df6a6cb5ce6
---

# Wave 8 — Fiber-Optic Infrastructure: OS2/OM3/OM4/OM5, MPO/MTP, Polarity, Trunks, Cassettes, Testing

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

- MPO (Multi-fiber Push-On) is the IEC 61754-7/TIA-604-5 (FOCIS-5) standard connector; MTP® is US Conec's enhanced, fully intermateable implementation with floating ferrule, removable housing, and round (vs angular) guide pins [official: IEC 61754-7, TIA-604-5; vendor-reported: US Conec].
- Fiber counts: **MPO-8** (1 row of 8), **MPO-12** (1×12), **MPO-16** (1×16, new MT-16 ferrule, FOCIS-18/TIA-604-18), **MPO-24** (2×12) [official: TIA-604-5/604-18; secondary: network-switch.com]. MTP-12 supports 40G/100G SR4 (8 of 12 fibers active, 4 dark); MTP-16 is the native interface for 400G/800G SR8/DR8 (8 Tx + 8 Rx) [secondary: network-switch.com; vendor-reported: US Conec catalog]. MTP-24 is used for trunk backbones and breakouts (2×MTP-12) [secondary: network-switch.com].
- Female/male (unpinned/pinned): transceivers and modules are typically unpinned (female); trunks are typically pinned (male); cassettes unpinned — the system balances so mated pairs always have one male and one female [secondary: network-switch.com]. Convention: **devices expect female; trunk cables male; patch cords at the electronics female** [independent/secondary consensus].
- APC vs UPC: single-mode MPO is normally APC (8° angled polish, return loss ≥60 dB); multimode is UPC (flat, return loss ≥20–50 dB). APC is mandatory for reflection-sensitive SM links [vendor-reported: FS.com (all SM MPO=APC), network-switch.com; official: return-loss requirements per TIA-568].
- Insertion-loss grades: **Standard ≤0.5 dB (typical 0.25–0.35 dB), Elite ≤0.35 dB** (US Conec MTP Elite max 0.35 dB) [vendor-reported: US Conec, FOCC press material]. FS.com lists its trunk/MPO products at "0.35 dB Max" with US Conec connectors [vendor-reported: FS.com].
- **MTP PRO field feature**: field-configurable polarity and gender (pins) without removing the connector housing — key retracts on one side while a hidden key extends on the other via a polarity-change port; pins can be installed/removed repeatedly with a pin-keeper design exceeding industry pin-retention forces (IEC 61754-7 requires ≥19.6 N). MTP PRO X is the APC-optimized variant (no field polarity change, field gender change only) [vendor-reported: US Conec product catalog]. FS.com sells MTP PRO trunk cables (OS2/OM3/OM4, 8–12 fibers, OFNP/LSZH) marketed for pin/polarity change in the field [vendor-reported: fs.com].

### 8.3 Polarity methods (TIA-568)

- TIA-568.3-D defines three connectivity/polarity methods using MPO-12 as the worked example; TIA-568.7 was cited by one source for MPO polarity — a vendor-blog claim, unverified against the standard text itself [official: TIA-568-C.0/568.3-D Methods A/B/C; unverified: "TIA-568.7" reference].
  - **Method A**: Type A straight-through cable, key-up to key-down (Position 1→1). The physical flip provides the crossover; duplex A-to-B cords + one A-to-A cord in the link [official: TIA-568-C.0; secondary: cablinginstall.com, sumitomoelectriclightwave.com].
  - **Method B**: Type B cable, key-up to key-up, with internal fiber reversal (1→12, 2→11…); A-to-B duplex cords [official: TIA-568-C.0].
  - **Method C**: Type C cable, key-up to key-down, pairwise flips (1→2, 2→1…); used with A-to-B duplex cords [official: TIA-568-C.0]. Least common in datacenters.
- **Mapping to 400G/800G optics**: 8-fiber systems (Base-8/MTP-8) map 1:1 to SR8/DR8 transceivers (8 Tx + 8 Rx on MTP-16, split to 2× MTP-8) [vendor-reported: Corning EDGE8, US Conec]. 12-fiber MTP-12 serves 40G/100G SR4 and 400G-DR4 (MPO-12 APC) [vendor-reported: FS.com, network-switch.com]. 16-fiber MTP-16 is native for 400G/800G SR8/DR8 [vendor-reported: US Conec; secondary: network-switch.com]. 24-fiber MTP-24 trunks break to 2×MTP-12 or 3×MTP-8 [vendor-reported: FS.com conversion-cable listings].
  - Corning's "EDGE8" Base-8 approach (8-fiber MTP) achieves 100% fiber utilization for 40/100/400G parallel optics with 1:1 port mapping and avoids conversion modules (claimed up to 50% less link attenuation vs Base-12 conversions) [vendor-reported: Corning].
- Note: TIA-568-C.0 did not standardize polarity for multi-row (24-fiber) connectors; vendor configurations (Type A/B/C equivalents) are non-standardized options [secondary: cablinginstall.com].

