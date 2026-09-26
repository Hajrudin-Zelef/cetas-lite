---
id: etape6-phasec-optics-cabling/00-front-matter/6-8-reliability-compatibility
title: "6.8 Reliability / compatibility"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "EU", "Nvidia", "United States"]
dates: ["2025-04", "2025-12", "2026-04", "2026-09", "2026-09-22"]
keywords: ["cost", "dsp", "gpu", "hyperscaler", "latency", "lpo", "nvidia", "optics", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1154, 1193]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: edd16c79347f40093d745b49d95ea6244e36114106fa7af5f79a7e24f987adf8
---

# 6.8 Reliability / compatibility

**Others:**
- Optcore Cisco QDD-400-CU3M-compat 400G QSFP-DD passive DAC 3 m, 26AWG — US$209.00 excl. VAT — optcore.net QDD-400G-DAC-P3M (page updated ~700 days ago; still live ~April 2026 crawl).
- ROBOfiber/DataInterfaces (US, crawled ~September 2026): 100G QSFP28 passive DAC 1 m $40.00; 100G→4×SFP28 breakout 3 m $99.00; 400G QSFP-DD AOC 3 m $705.00; 400G→4×QSFP56 passive DAC 3 m $330.00; 100G→4×SFP28 AOC breakout 3 m $201.00 — datainterfaces.com.

**Volume-price context note:** 800G DR8 optical transceivers at volume (1000+ units) list $1,000–1,400, LPO $700–900 [secondary] (saastisfy.fr price list, April 2025) — context only; transceivers are not DAC/AOC/AEC and are included solely to frame the "AEC ≈ half of optics" claim.

### 6.8 Reliability / compatibility

**EEPROM coding.** Every DAC/AOC/AEC carries an EEPROM (SFF-8636 for QSFP28-class; CMIS for QSFP-DD/OSFP-class) readable over I2C, programmed with vendor name, part number, length, and compliance codes [vendor-reported] (Dawnray, FS.com). Third-party vendors (FS.com, QSFPTEK, Optcore) operate in-house coding labs that program cables to mimic OEM identities (Cisco, Arista, Juniper, NVIDIA/Mellanox, Dell EMC…); FS.com: "Our in-house coding facility programs all of our parts to standard OEM specs for compatibility" [vendor-reported] (FS datasheet 20240428120743z6yjda.pdf). **Vendor lock-in is real:** "Some switches enforce vendor coding (lock-in)" — always confirm vendor-approved/coded cables for critical deployments [secondary] (network-switch.com). AECs additionally run a microcontroller implementing CMIS (Microchip META-DX2C SDK implements CMIS 5.2; Credo CLOS AEC uses CMIS) — AECs present a much richer management interface (telemetry, diagnostics, firmware) than passive DACs [official] (Microchip; Credo CLOS brief).

**Failure modes (by type).** Passive DAC: mechanical — conductor fatigue from tight bends (respect min. bend radius: 33–72 mm depending on AWG/length), connector mating-cycle wear, crosstalk/insertion-loss degradation. Thick 26AWG bundles also create airflow/thermal issues in dense racks (a deployment-level "failure" via overheating neighbors) [secondary]. AOC: VCSEL/laser wear-out and fiber-connector contamination/damage; no re-termination — one bad end scraps the whole assembly [secondary]. AEC: active-silicon failure modes (retimer chip, PMIC, thermal) — but far fewer components than an optical link. Credo claims up to 100M hours MTBF and "100 times better reliability" than optical solutions for its 800G AECs [vendor-reported] — **flag: vendor marketing claim, not independently verified.** Link-flap sensitivity in AI backends: Credo's ZeroFlap family specifically targets "zero soft link flaps" for lossless RDMA AI fabrics [official].

**BER considerations at 100G PAM4 lanes.** IEEE 802.3 requires pre-FEC BER ≤ 2.4×10⁻⁴ for 100G-PAM4-class clauses (e.g., 100GBASE-CR2/KP4), with KP4 RS(544,514) Reed-Solomon FEC mandatory on 400G and most 100G-PAM4 electrical/optical interfaces, delivering post-FEC BER <10⁻¹²–10⁻¹⁵ [official-via-secondary] (Tektronix PAM4 primer; EDN 400G FEC coverage). Raw (pre-FEC) BER on PAM4 links "can easily reach 10⁻⁶ to 10⁻⁴" vs the traditional 10⁻¹² NRZ target; RS-FEC provides ~7–8 dB coding gain [secondary] (fibermall.com FEC explainer, December 2025). Cable-vendor BER claims: 400G passive DAC "BER better than 1E-15" (Vchung) [vendor-reported]; 800G AEC "BER (Post-FEC) <10⁻¹⁵" (QSFPTEK) [vendor-reported]; Credo retimer/DSP cables are specified against KP4 FEC thresholds (optional inner Hamming (128,120) FEC in some DSPs) [official] (Credo Bluebird brief). Practical implication: at 100G/lane PAM4, DAC/AEC link budgets assume FEC is ON; a passive DAC that meets pre-FEC 2.4e-4 at its rated length/loss will deliver effectively error-free post-FEC operation. Switch ASICs (Broadcom Tomahawk/Jericho, NVIDIA Spectrum, Cisco Silicon One) terminate FEC in hardware; disabling FEC breaks auto-negotiation on standards-compliant ports [secondary]. AEC retimers regenerate the signal and can terminate/monitor FEC, which is why AECs hold BER margin at 5–7 m where passive copper's pre-FEC BER would exceed the KP4 correction threshold (mechanism per vendor DSP descriptions; specific threshold margins per cable are vendor-proprietary — [unverified] at per-SKU level).

### 6.9 Flags & conflicts summary (Wave 6)
1. **Credo "Dove" chip** — could not verify; do not use. **Seagull** verified only as optical DSP, not AEC retimer.
2. **100G QSFP28 passive max length** — 3 m standard vs 5 m vendor-stretch (Optcore); unresolved.
3. **ACC chip placement** — Rx-end only (Fibermall) vs both ends (VEEX); unresolved, likely vendor-dependent.
4. **800G "DAC hits the wall" (Credo 2021)** vs 2026 reality of widely sold 800G passive DACs (FS.com, QSFPTEK) — vendor positioning vs market fact.
5. **AEC retimer latency** — no cable-level ns figure sourced; ~100 ns/hop is industry lore [unverified]; only verified figure is Credo Bluebird *optical* DSP <40 ns/direction.
6. **All prices** are single-unit web list prices (crawled Jul–Sept 2026), region/currency/SKU-dependent, and not comparable to hyperscaler volume pricing; several comparisons above are explicitly directional.
7. **Credo reliability/power/cost claims** (100M-hr MTBF, 100× reliability, half-power vs optics, $1,000/GPU saving) are vendor-reported marketing figures without independent verification.
8. **10G/25G passive DAC lengths and TE/Molex AEC SKUs** were not directly sourced — omitted or flagged rather than invented.

---

## Wave 7 — FS.com Optics Catalog & Pricing (Second Pass: SKUs, DAC/AOC Gaps, Coding, Warranty)

*Research date: 2026-09-22. Single-writer wave. This wave complements Wave 1 (FS.com catalog) and Wave 6 (DAC/AOC/AEC). Price caveat: all fs.com prices are public US$ list prices unless noted; pages were crawled 53–343 days ago, so prices may have drifted since capture; EU prices include VAT at country rates; no volume/discount pricing captured. Tags: [official] (fs.com pages), [vendor-reported], [independent], [secondary], [unverified].*

### 7.1 Transceiver SKUs/prices by speed

**10G SFP+.** Flagship FS volume line: "10G SFP+ SFP-10G-SR 300m 850nm Duplex LC/UPC Module, Cisco" US$25.00 (1.4M sold, 2.5K reviews); "SFP-10G-LR 10km 1310nm Duplex LC/UPC, Cisco" US$34.00 (1.1M sold) [official]. Also "10G SFP+ 30m RJ-45 Module, Cisco" US$86.00; custom 10GBASE-BX SFP+ 1270nm-TX/1330nm-RX 40km Simplex LC US$69.00 [official]. 1G reference points on same category page: SFP-1G-SX US$8.80, SFP-1G-LH/LX US$10.00, GLC-T (RJ-45) US$35.00 [official].

**25G SFP28.** "SFP-25G-SR-S 100m 850nm Duplex LC/UPC, Cisco" US$49.00 (241.6K sold); "SFP-25G-LR-S 10km 1310nm Duplex LC/UPC, Cisco" US$74.00 (66K sold) [official]. Dual-rate 10/25G: SFP-10/25G-LR-S US$86.00 (P/N SFP-25GMLR-31, SKU 155189); 10/25G MSR 100m US$49.00; SFP-10/25G-CSR-S 400m US$74.00 [official]. Industrial variants carry ~+40% premium: 25G SR industrial US$70.00, 25G LRL industrial US$74.00, 25G MLR industrial 10km US$99.00 [official].

**40G QSFP+.** "QSFP-40G-SR4-S 150m 850nm MPO-12/UPC, Cisco" US$43.00 (140.5K sold, 820 reviews) — best seller; "QSFP-40GE-LR4 10km, Cisco" US$299.00 (39.3K sold); "QSFP-4X10G-LR-S 10km PLR4, Cisco" US$359.00; "QSFP-40G-SR-BD 300m SWDM4, Cisco" US$369.00; "QSFP-40G-PLRL4 2km, Cisco" US$309.00; "QSFP-40G-ER4 40km, Cisco" US$1,109.00 [official].

**50G.** "50G SFP56 100m 850nm Duplex LC/UPC Module, Generic" €303.45 incl. VAT / €255.00 VAT excl. [official, FS Europe]; "50G QSFP28 80km 1310nm Duplex LC/UPC, Cisco" (50G ZR2) €3,716.37 incl. / €3,123.00 VAT excl. [official, FS Europe].

