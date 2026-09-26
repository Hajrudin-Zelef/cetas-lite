---
id: etape6-phasec-optics-cabling/00-front-matter/wave-6-dac-vs-aoc-vs-aec-direct-attach-cable-technologies
title: "Wave 6 — DAC vs AOC vs AEC: Direct-Attach Cable Technologies"
domain: front-matter
role: reference
task: reference
actors: ["Nvidia", "United States"]
dates: ["2021-10", "2026-07", "2026-08", "2026-09", "2026-09-22"]
keywords: ["capex", "cost", "dsp", "gpu", "hyperscaler", "latency", "memory", "nvidia", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [994, 1038]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 8c3579b2b85131af655bec9d30e6d895db63d18b5be751c02ee1fc5359bae7b1
---

# Wave 6 — DAC vs AOC vs AEC: Direct-Attach Cable Technologies

## Wave 6 — DAC vs AOC vs AEC: Direct-Attach Cable Technologies

*Research date: 2026-09-22. Single-writer wave. Tags: [official], [vendor-reported], [independent], [secondary], [unverified].*
*Price caveat (applies throughout Wave 6): third-party retail prices (FS.com, QSFPTEK, Cables on Demand, Optcore, ROBOfiber/DataInterfaces) are single-unit web list prices captured by search crawls between ~July and mid-September 2026, usually excluding VAT/shipping, and are NOT volume/hyperscaler contract prices. Prices across vendors, currencies, regions, and OEM-compatibility SKUs are not directly comparable — flagged where relevant. Every SKU/P-N below was quoted verbatim from its source; none were invented.*

### 6.1 Passive DAC (Direct Attach Copper, twinax)

**How it works.** A passive DAC is a length of high-speed differential coaxial (twinax) cable terminated directly into QSFP/OSFP/SFP housings at both ends: an end-to-end electrical path with NO active components — no lasers, no DSP, no CDR, no amplification [secondary] (fibermall.com DAC/AOC/AEC blog; network-switch.com DAC guide; kb.veexinc.com cable knowledge base). The only electronics in the housings are a small EEPROM accessed over I2C/TWI (SFF-8636, now the CMIS memory map) carrying vendor ID, part number, cable length, and compliance codes so the host switch/NIC can identify and sometimes authenticate the cable [vendor-reported] (dawnraytech.com edge-core DAC page; FS.com AEC blog). Consequence: near-zero power draw (only EEPROM/ID leakage), minimal latency (no electro-optical conversion or digital processing — copper propagation delay only), and lowest cost per link. FS.com: "lowest-cost, lowest-latency, and near-zero-power connections for high-speed links" [vendor-reported] (fs.com product 284167).

An independent test (AICPlight, Medium, ~August 2026) of a 1 m 200G QSFP56 passive DAC (P/N Q200-Q200-CU1) between two NVIDIA DGX Spark systems reported "nanosecond-level latency" with zero raw physical errors per lane pre-FEC using standard RS-FEC on ConnectX-7 NICs [secondary].

**Typical max passive reach by speed (per vendor listings; see conflicts below):**
| Speed / form factor | Lane signaling | Max passive length (vendor listings) | Typical AWG |
|---|---|---|---|
| 10G SFP+ | 10G NRZ | up to ~7 m (commonly listed; some vendors 10 m) [unverified — not directly sourced in this research] | 24–30 |
| 25G SFP28 | 25G NRZ | ≤3–5 m [secondary; conflicting — network-switch.com says ≤3 m for 25G/100G+; Cisco-style listings commonly 1/2/3/5 m; unverified for 25G specifically] | 24–30 |
| 40G QSFP+ | 4×10G NRZ | up to 7 m [secondary] | 24–30 |
| 100G QSFP28 | 4×25G NRZ | 3–5 m: Optcore lists 1/2/3/5 m 100G QSFP28 DAC; QSFPTEK 3 m; most enterprise SKUs stop at 3 m [vendor-reported] | 30AWG typical at 3 m (QSFPTEK QT-Q28-PC3, 33 mm bend) |
| 200G QSFP56 | 4×50G PAM4 | ≤3 m [vendor-reported] (NADDOD blog: 200G QSFP56 DAC 0.5–3 m, <0.1 W) | 28–30 |
| 400G QSFP-DD (8×50G PAM4) | 8×50G PAM4 | 0.5–3 m: FS.com QDD-400G-PCxxx series 0.5–3 m (26–30AWG); FS QSFP-DD passive datasheet: 1 m=30AWG, 2 m/3 m=28AWG [vendor-reported] (fs.com product 82457; FS datasheet 20230602113002ttej4d.pdf); FS 400G OSFP flat-top passive only 0.5–2 m; QSFP112 (4×100G) passive only 0.5–1.5 m [vendor-reported] (FS.com 400G cable table) | 26–30AWG |
| 800G OSFP/QSFP-DD (8×100G PAM4) | 8×100G PAM4 | 0.5–3 m (vendor-limited): FS.com lists 800G OSFP passive 0.5 m (30AWG) and 1 m (30AWG); breakout 2 m (26–28AWG) [vendor-reported] (fs.com 800g-dac-aoc category, crawled ~July 2026); Vchung/ONIS NVIDIA MCP7Y00-Nxxx-compatible 800G→2×400G breakout 1–3 m (26AWG at 2.5 m, 30AWG at 1–1.5 m) [vendor-reported] | 26–30AWG |

**Conflicts/flags:**
- 100G QSFP28 passive max is listed as 3 m by most enterprise vendors but up to 5 m by Optcore; treat "3 m" as the de-facto standard and 5 m as a vendor-stretch claim [unverified-resolution].
- Credo (October 2021): "Credo sees 800G as the point where passive DACs hit the wall – they are far too thick and rigid for many customer applications and impose a high cost and engineering burden on switch manufacturers" (Don Barnetson, VP Product) [vendor-reported]. However, 800G passive DACs are widely sold in 2026 (see pricing), so this reads as vendor positioning, not a technical impossibility.

**AWG gauges.** Rule of thumb: lower AWG = thicker conductors = lower loss = longer reach but stiffer cable and worse airflow; higher AWG = thinner, more flexible, shorter reach [secondary]. Sourced pairings: 400G QSFP-DD 3 m → 28AWG; 400G 1 m → 30AWG (FS) [vendor-reported]; 800G breakout 2.5 m → 26AWG; 800G 1–1.5 m → 30AWG (Vchung) [vendor-reported]; 100G QSFP28 3 m → 30AWG (QSFPTEK) [vendor-reported]; Credo 800G AEC uses 32AWG at up to 2.5 m (vs 26AWG passive DAC) — illustrating the thickness penalty of passive reach [vendor-reported] (cablinginstall.com, Credo AEC coverage).

**Power.** ≤0.1 W per cable assembly for 400G and 800G passive DACs (FS.com spec tables) [vendor-reported]; "extremely low power consumption near 0W" (Dawnray) [vendor-reported]; some comparison tables quote "<0.3W" [secondary] (isp-home.com AOC product page).

**Latency/cost position.** Latency: lowest of all cable types — no conversion, no DSP (nanosecond-scale) [secondary]. Cost: lowest capex per link; stated by all vendor sources [vendor-reported/secondary].

**Use cases.** ToR-switch-to-server links inside a rack, adjacent-rack links, spine→super-spine intra-rack links, InfiniBand storage fabrics (HDD/flash subsystems) [vendor-reported] (FS.com MCP7Y00-N02A page). AI/GPU: 200G QSFP56 DACs validated on NVIDIA DGX Spark (ConnectX-7 200G) [vendor-reported] (NADDOD, 2026); dual-DGX Spark peer-to-peer 200G DAC test [secondary] (AICPlight).

**Breakout variants (sourced SKUs/prices, FS.com ~July 2026 unless noted):**
- 100G QSFP28→4×25G SFP28 passive breakout 3 m, 30AWG, US$136 (5.7K sold) [vendor-reported] (fs.com 100g-qsfp28-dac category).
- 400G QSFP-DD→2×200G QSFP56 passive 0.5–3 m (QDD-400G-2QPC005…2QPC03; 30/28AWG) [vendor-reported] (FS datasheet 20240428120743z6yjda.pdf); 400G QSFP-DD→4×100G passive up to 3 m (Q-4Q28PC005…PC03) [vendor-reported] (FS datasheet 400g-qsfp-dd-to-4x100g passive DAC breakout).
- ROBOfiber 400G QSFP-DD→4×QSFP56 passive 3 m, US$330 [vendor-reported] (datainterfaces.com).
- 800G OSFP→2×400G OSFP passive 2 m, 26AWG, US$299; 800G OSFP→2×400G QSFP112 passive 2 m, 28AWG, US$299; 800G OSFP→4×200G OSFP 2 m, 28AWG, US$461; 800G OSFP→4×200G QSFP112 2 m, 28AWG, US$324 [vendor-reported] (fs.com 800g-dac-aoc category).
- Straight 800G passive: 800G QSFP-DD 0.5 m US$199; 800G OSFP 1 m US$186; 800G OSFP 0.5 m US$161 [vendor-reported] (same page).

### 6.2 Active copper DAC / ACC (Active Copper Cable)

**Where the chips sit / what they do.** ACC adds active signal-conditioning electronics to the copper assembly. Dominant architecture: a linear redriver chip (continuous-time linear equalizer, CTLE — an analog amplifier/equalizer) that boosts and equalizes the signal [secondary]. **Conflict:** Fibermall states the linear redriver sits "at the Rx end of the cable" [secondary]; VEEX states ACC adds "signal equalization and amplification electronics at both ends" [secondary]; FS.com's AEC article says redrivers "amplify the signal" but "do not recondition it… they amplify the noise as well," contrasting redriver (analog) with retimer (clock restore + resample) [vendor-reported] (FS.com AEC blog). Treat exact placement as vendor-dependent. Because a redriver has no clock-data recovery, jitter and noise accumulate and are amplified — ACC reach gain over passive is modest [secondary].

