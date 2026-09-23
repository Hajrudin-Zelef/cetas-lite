---
id: etape6-phasec-optics-cabling/00-front-matter/6-3-aoc-active-optical-cable
title: "6.3 AOC (Active Optical Cable)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "EU", "United States", "xAI"]
dates: ["2021-09", "2022-03", "2022-10", "2023-10", "2023-11", "2024-03", "2024-10", "2024-10-10", "2026-01", "2026-08", "2026-09"]
keywords: ["cost", "disaggregated", "dsp", "ethernet", "gpu", "hyperscaler", "latency", "optics", "packaging", "pricing", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1047, 1098]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 07c359672a1aca0568637f450415cce3db367668e5025543506b588b4fd291f0
---

# 6.3 AOC (Active Optical Cable)

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

**Key chip vendors.**
- **Credo (HiWire AECs)** — vertically integrated: own DSP/SerDes retimer silicon inside HiWire cable assemblies [official] (BusinessWire press releases). Chip-family naming: "HiWire" = Credo's AEC cable brand (purple jacket), verified [official]. "Seagull" = Credo's PAM4 DSP family (Seagull 110 2×50G, Seagull XR8 8×50G) — but sourced as **optical DSPs with integrated VCSEL/EML drivers** for AOC/transceiver use, *not* documented as AEC cable retimers in the sources found [official] (BusinessWire, March 2022; September 2021). "Dove" = [unverified] — no source found for a Credo "Dove" retimer/AEC chip; do not use without verification. Verified Credo retimer families: **Screaming Eagle 112G** (3rd-gen 112G retimer DSP, up to 1.6T capacity, sampling ~2022) [official] (lightwaveonline.com coverage).
- **Astera Labs** — **Taurus Ethernet Smart Cable Modules**: purpose-built modules for cable vendors to build 200/400/800G Ethernet AECs; demonstrated at OCP Global Summit (October 2023) [official] (BusinessWire, October 2023). Note: Astera's **Aries** family = PCIe/CXL retimers and Smart Cable Modules for PCIe AECs (Aries 6 SCM: up to 7 m PCIe 6.x, January 2026 product brief) — PCIe fabric, not Ethernet; do not conflate [official] (asteralabs.com Aries PCIe/CXL SCM product brief).
- **Microchip** — META-DX2C (part #PM6254) 112G retimer for 800G AECs, 40 dB reach, CMIS 5.2 SDK, available since November 2023 [official].
- **Broadcom** — DSP-based AEC models referenced (64-tap FFE/12-tap DFE) in a Fibermall analysis [secondary] (medium.com/@fibermall.com 400G NDR splitter / OSFP 800G analysis) — [unverified] at product level; Broadcom DSPs are documented inside FS.com 400G AOCs, not AECs, in the sources found.
- **Cable assemblers selling AECs:** Amphenol (800G OSFP AEC via Cables on Demand), FS.com, QSFPTEK, Vchung-style white-label vendors; TE Connectivity and Molex are major high-speed cable assemblers generally but no per-SKU AEC evidence was pulled for them — [unverified] at SKU level.

**Max lengths (sourced).**
- **400G AEC:** up to 7 m — Credo LP SPAN Gen 2: 3/5/7 m options at 4.5 W/end (400G PAM4 8×56G); Gen 1: 3/5 m at 8.5 W/end [official] (BusinessWire, September 2021).
- **800G AEC:** Credo LP CLOS: up to 2.5 m (32AWG, 2021) [official]; Credo **ZeroFlap (ZF)** 800G AEC: up to **7 m**, launched October 10, 2024, for AI backend networks; four variants: 800G OSFP↔OSFP, OSFP↔OSFP-RHS, OSFP↔2×OSFP-RHS, OSFP↔2×Q112 (QSFP112) [official] (venturebeat.com, Credo ZeroFlap launch).
- Third-party 800G AEC products (in production/retail as of September 2026): Amphenol 800G OSFP (finned) → 2×400G OSFP (flat) AEC in 3 m (32AWG), 5 m (30AWG), 7 m (28AWG) — in stock at Cables on Demand [vendor-reported] (cablesondemand.com QSFP-DD/OSFP cables page); QSFPTEK generic 800G OSFP→2×400G OSFP AEC 1 m, copper link up to 7 m max, 26–34AWG [vendor-reported] (qsfptek.com product 103796); FS.com EU 800G OSFP→2×400G OSFP AEC 4 m (P/N OSFP-800G-2OFLAE04, SKU 312079) [vendor-reported] (fs.com/eu-en product 312079).

**Power per cable (sourced).** 200G AEC: 2.5 W/end (Credo LP SPAN Gen 2) / 4.5 W/end (Gen 1) [official]. 400G AEC: 4.5 W/end (Gen 2) / 8.5 W/end (Gen 1) [official]. 800G AEC: QSFPTEK generic breakout <9 W max [vendor-reported]; FS.com EU 800G AEC 4 m: 12 W (800G end) / 10 W per 400G end [vendor-reported]; FS.com MX 800G OSFP AEC 1 m: ≤12 W [vendor-reported]; 1.6T OSFP-XD AEC: <20 W/end [official] (Credo, October 2022). Relative claims: Credo LP CLOS AEC = "half the power of optical cabling solutions" [vendor-reported]; ZeroFlap = "power savings of up to 14W per link" vs legacy optics [vendor-reported]; Credo CLOS AEC brief claims 75% less power than optical solutions [vendor-reported] — vendor marketing ratios, treat as [vendor-reported/unverified-independently].

**Latency added by retimers.** **No AEC-cable-specific retimer latency figure was sourced.** Closest verified numbers: Credo Bluebird 1.6T **optical** DSP: "latency below 40ns in each direction" [official-via-press] (stocktitan.net, Credo Bluebird 1.6T optical DSP). Industry rule-of-thumb for a PAM4 DSP retimer hop is on the order of ~100 ns — **[unverified]**; do not quote as fact. FS.com EU lists "ultra-low latency" for its 800G AEC without a number [vendor-reported].

**Use cases.** Rack-to-rack and intra-rack where passive DAC is too short/rigid but AOC is overkill: Distributed Disaggregated Chassis (DDC) CLOS fabrics [official] (Credo CLOS AEC brief); GPU/AI backend lossless RDMA networks — Credo ZeroFlap "zero soft link flaps" for AI backend, endorsed by xAI network engineering for 100,000+ GPU builds [official] (VentureBeat/BusinessWire, October 2024); in-cable **speed-shifting** AECs bridging 112G-lane NICs to legacy 56G-lane 12.8T/25.6T/51.2T ToRs (Credo 400G AI/ML backend family, sampling March 2024, production Q3 2024 — e.g., 400G OSFP-RHS 4×112 ↔ QSFP-DD 8×56; 800G 2×(4×112) ↔ 8×112) [official] (nasdaq.com press release, March 2024).

**Pricing (sourced, single-unit web list, September 2026 unless noted).**
- Amphenol 800G OSFP→2×400G OSFP AEC (OEM P/N NND1JA-0303 etc.), Cables on Demand, in stock ~August 2026: 3 m $1,869.93; 5 m $2,131.02; 7 m $2,503.51 (1–50 qty) [vendor-reported] (cablesondemand.com).
- QSFPTEK generic 800G OSFP→2×400G OSFP AEC 1 m (Product No. 103796): US$1,664.10, stock dated 16 September 2026 [vendor-reported] (qsfptek.com product 103796).
- FS.com EU 800G OSFP→2×400G OSFP AEC 4 m (OSFP-800G-2OFLAE04, SKU 312079): €1,568.00 VAT excl. (€1,865.92 incl.) [vendor-reported] (fs.com/eu-en product 312079).
- FS.com MX 800G OSFP AEC 1 m: MXN$21,151 [vendor-reported]; 800G OSFP AOC 1 m: MXN$48,135 (same page — AEC ≈44% of AOC list price in that storefront; currencies identical, but different cable types/lengths — directional only) [vendor-reported] (fs.com/mx 800g category).
- Credo claim (October 2024): "cost savings of up to $1,000 per GPU" vs legacy optics in AI clusters [vendor-reported] — hyperscaler-context marketing figure, [unverified] independently.
- FS.com US 800G active breakout DAC (not full AEC): 800G OSFP→2×400G QSFP112 active 3 m, US$1,136; 800G OSFP→4×200G QSFP112 active 5 m, US$1,536 [vendor-reported] (fs.com 800g-dac-aoc category).

**800G AEC status as of September 2026.** **In production and broadly available:** Credo ZeroFlap 800G (7 m) shipping since late 2024 with "millions of HiWire AECs deployed at tier-one hyperscalers" [vendor-reported]; Amphenol and white-label (QSFPTEK, FS.com) 800G AECs in retail stock [vendor-reported]; 650 Group (Alan Weckel) quoted 2022: "AECs will quickly replace direct attached copper" as hyperscalers scale [independent-analyst-via-vendor-PR]. AEC is an established, growing category at 800G — not sampling-only.

