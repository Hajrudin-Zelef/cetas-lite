---
id: etape6-phasec-optics-cabling/00-front-matter/6-3-aoc-active-optical-cable
title: "6.3 AOC (Active Optical Cable)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "United States"]
dates: ["2023-11"]
keywords: ["cost", "dsp", "latency", "packaging", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1039, 1070]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: fcc69e26701acd1c3eb51784449f6c0750fb7959285c6b2e8053ab29ad5986a3
---

# 6.3 AOC (Active Optical Cable)

**Extended reach vs passive.** FS.com: 400G ACC reach is "relatively longer by 2–3 meters compared to 400G DAC" (i.e., up to ~5–6 m total at 400G) [vendor-reported] (FS.com 400G DAC/AOC/ACC/AEC guide). FS.com QSFP-DD datasheet order table lists active copper (AC) cables: QDD-400G-AC01/02/03/05/07 = 1/2/3/5/7 m; QSFP112 active 1–5 m; while passive (PC) tops out at 3 m [vendor-reported] (FS datasheet 20240403150013b786tc.pdf). Network-switch.com comparison: active DAC 7–15 m, power <1W (speed-dependent; likely lower-speed generations) [secondary].

**Power.** 400G QSFP-DD active copper (FS P/N QDD-400G-AC03, 3 m, Arista-compatible, SKU 177377): "Built-in Macom Chip, Max. power consumption 2.5W" [vendor-reported] (fs.com/sg product 177377). Generic comparison tables quote ACC <1W at lower speeds; at 400G with retimer-class chips power is higher — FS's own blog notes retimer-based active chips carry "higher prices and power consumption" than redrivers [vendor-reported] (FS.com 400G guide).

**Vendors (chips + cable assemblers).** Chip side (sourced): MACOM (redriver/linear chips inside FS.com 400G active DAC) [vendor-reported]. Credo and Astera Labs are retimer/DSP vendors whose chips are used in AEC assemblies (see §6.4); no source names them as ACC redriver suppliers — do not conflate [unverified for ACC]. Cable/assembler side (sourced): FS.com (QDD-400G-ACxx series), Optcore, Amphenol, Molex, TE Connectivity (general DAC/AEC assembler market; no per-SKU evidence pulled for TE/Molex ACC — flagged [unverified] at SKU level). QSFPTEK sells "active breakout DAC" 800G products (e.g., 800G OSFP→2×400G QSFP112 active 3 m, ≤1.5W/≤0.6W, US$1,136) [vendor-reported] (fs.com 800g-dac-aoc category).

**Price premium over passive.** "The cost of 400G ACC is higher than that of 400G passive DAC due to the presence of active chips internally" [vendor-reported] (FS.com 400G guide). Non-comparable datapoints (flag): FS.com Singapore 400G QSFP-DD active 3 m (QDD-400G-AC03) = SGD 683.43 GST incl. [vendor-reported]; Optcore Cisco-compatible 400G QSFP-DD passive 3 m (26AWG) = US$209 excl. VAT [vendor-reported] (optcore.net QDD-400G-DAC-P3M). Different vendors, currencies, and OEM-coding — directional only: active costs a multiple of passive.

### 6.3 AOC (Active Optical Cable)

**How it works.** Two permanently attached transceiver ends (no field re-termination possible) joined by multimode fiber; electrical→optical conversion at one end, optical→electrical at the other [secondary] (kb.veexinc.com: "permanently connected xSFP transceiver modules at both ends"). Typical engine: 850 nm VCSEL transmitter + PIN photodiode receiver, internal CDR on Tx and Rx channels, digital diagnostic monitoring (DDM) via SFF-8636/CMIS [vendor-reported] (mellanoxnetwork.com 25G SFP28 AOC listing). From the host's perspective the electrical interface is the same pluggable cage (QSFP28/QSFP-DD/OSFP); the host sees a module with CDR.

**Lengths.** Offered from 1–3 m up to 100 m: 25G SFP28 AOC up to 100 m on OM4 [vendor-reported]; 200G QSFP56 SR4 AOC: 100 m OM4 / 70 m OM3 [vendor-reported] (mellanoxnetwork.com 200G QSFP56 SR4 AOC listing); 100G QSFP28 AOC commonly sold 3–100 m (FS.com 3 m listing; UnitekFiber 5 m listing) [vendor-reported]; general comparison tables: AOC up to 100 m (OM4), ~70 m (OM3) [secondary] (isp-home.com; optcore.net article090).

**Fiber inside.** Multimode OM3/OM4 fiber (850 nm VCSEL) [vendor-reported/secondary].

**Power (sourced per-SKU).** 25G SFP28 AOC: <1.0 W [vendor-reported]. 100G QSFP28 AOC 3 m: <2.2 W (FS.com, US$204) [vendor-reported] (fs.com 100g-qsfp28-dac category); <2.5 W per end (UnitekFiber 5 m) [vendor-reported]. 200G QSFP56 SR4 AOC: <4 W typical [vendor-reported]. 400G QSFP-DD AOC 3 m (Cisco compat, P/N QDD-400G-AO03, SKU 146374, Broadcom chip): <8 W per end, US$1,011 [vendor-reported] (fs.com product 146374). 400G QSFP-DD→4×100G AOC breakout 3 m (P/N QDD-400G-4QAO03, SKU 150593, Inphi chip): <3.5 W per 100G end / 12 W per 400G end, AUD 2,504.70 GST incl. [vendor-reported] (fs.com/au product 150593). 800G OSFP AOC 1 m: ≤14 W (FS.com Mexico listing, MXN$48,135) [vendor-reported] (fs.com/mx 800g-1.6t-osfp-qsfp-dd category). Generic 400G AOC "~10W" (FS.com guide) [vendor-reported]; comparison tables: AOC 1–2W+ (lower-speed gens) [secondary].

**Latency.** Very low; slightly above passive DAC due to E/O conversion and CDR, well below any routed hop. Universally ranked: DAC < AOC ≈ AEC < transceiver+patch in vendor comparison tables [secondary].

**Use cases.** Inter-rack, row-to-row (EoR/MoR), zone-to-zone links where passive DAC cannot reach (typically 5–100 m); EMI-sensitive environments (fiber immune to EMI); high-density racks where thin fiber improves airflow vs thick DAC bundles [secondary] (isp-home.com AOC page).

**Price vs DAC.** "The price of 400G AOC is typically higher than that of 400G DAC (assuming similar levels or packaging)" — fiber material + lasers at both ends [vendor-reported] (FS.com 400G guide). Sourced ratio: 100G QSFP28: passive DAC 3 m US$39–40 vs AOC 3 m US$204 (FS.com/QSFPTEK/ROBOfiber) — AOC ≈5× passive DAC at 100G [vendor-reported]; 400G: passive DAC 3 m US$209 (Optcore) vs AOC 3 m US$705 (ROBOfiber)/$1,011 (FS.com Cisco-compat) — AOC ≈3–5× passive [vendor-reported] (non-identical vendors — approximate).

**Fixed-length limitation.** AOCs are factory-terminated, pre-tested assemblies; length cannot be changed in the field and a failed end means replacing the whole cable ("permanently connected" module ends) [secondary].

### 6.4 AEC (Active Electrical Cable)

**How it works.** Copper twinax like a DAC, but with **retimer chips with CDR at both ends** of the cable (plus a small MCU for CMIS/I2C management). The retimer fully recovers clock and data, equalizes (DSP-based: FFE/DFE), and re-transmits a clean signal — resetting both the loss and jitter budgets at each end [vendor-reported] (FS.com AEC blog: "retimers… recondition the data signal… reducing noise and amplifying the signal"). [official] Microchip press release for META-DX2C: "high-performance, long-reach 112G SerDes that can support up to 40 dB reach" (globenewswire.com, November 2023). This enables thin-gauge copper (32–34AWG at 800G) at lengths passive DAC cannot achieve.

**How AEC differs from ACC.** ACC = linear redriver (analog CTLE), no CDR; amplifies signal *and* noise, jitter accumulates; modest reach gain (~+2–3 m at 400G), lower power/cost [vendor-reported/secondary]. AEC = DSP retimer with CDR at both ends; recovers clock, re-times, re-drives; blocks jitter propagation; longer reach (up to 7 m at 400G/800G) at higher power and cost [vendor-reported] (FS.com guide; Fibermall; FS.com AEC blog). FS.com: "AEC with re-timers provides a superior solution for maintaining data integrity… compared to ACC which uses linear amplifiers" [vendor-reported].

