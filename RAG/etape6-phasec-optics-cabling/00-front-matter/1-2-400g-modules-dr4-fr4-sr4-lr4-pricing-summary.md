---
id: etape6-phasec-optics-cabling/00-front-matter/1-2-400g-modules-dr4-fr4-sr4-lr4-pricing-summary
title: "1.2 400G modules — DR4 / FR4 / SR4 / LR4 (pricing summary)"
domain: front-matter
role: reference
task: pricing
actors: ["Broadcom", "EU", "Nvidia", "United States"]
dates: ["2026-09-22"]
keywords: ["pricing", "dsp", "lpo", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [69, 126]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: d7c5444451a3acc8ecb1f5a2db77434ce712a32da12c7150520f1922186f8e4e
---

# 1.2 400G modules — DR4 / FR4 / SR4 / LR4 (pricing summary)

### 1.2 400G modules — DR4 / FR4 / SR4 / LR4 (pricing summary)

| P/N | Brand coding | Form factor | Reach/type | λ | Connector | Price (USD, 2026-09-22) |
|---|---|---|---|---|---|---|
| QDD-DR4-400G-Si (SKU 128242) | Cisco | QSFP-DD | 500m DR4 | 1310nm | MPO-12 | US$749.00 [official] |
| QDD-DR4-400G (SKU 234483/183428) | HW/Dell | QSFP-DD | 500m DR4 | 1310nm | MPO-12/APC | US$549.00 [official] |
| QDD-400G-DR4-S | Cisco | QSFP-DD | 500m DR4 | 1310nm | MPO-12/APC | US$749.00 (US cat.); £421 excl. VAT UK [official] |
| QDD-400G-FR4 | Arista | QSFP-DD | 2km FR4 | 1310nm | Duplex LC/UPC | £556.00 excl. VAT UK [official] |
| QDD-400G-SR4 | Arista | QSFP-DD | 50m SR4 | 850nm | MPO-12/APC | £556.00 excl. VAT UK [official] |
| QDD-400G-SR8 | Arista | QSFP-DD | 100m SR8 | 850nm | MPO-16/APC | £168.00 excl. VAT UK [official] |
| QDD-400G-LR4 | Arista | QSFP-DD | 10km LR4 | 1310nm | Duplex LC/UPC | £674.00 excl. VAT UK [official] |
| MMA4Z00-NS400 | NVIDIA (ETH) | OSFP | 50m SR4 | 850nm | MPO-12/APC | US$769.00 [official] |

All 400G modules use 4x100G PAM4 (or 8x50G PAM4 for some QSFP-DD) electrical interfaces. 400G-DR4 is marketed for 4x100G breakout to QSFP-DR-100G [official].

**Gap:** No US-listed price found for FS 400G OSFP DR4 (only InfiniBand NDR variant US$1,124.00). **[unverified]**

### 1.3 800G modules — DR8 / FR8 / SR8 / 2xFR4

US category listing [official] (https://www.fs.com/c/800g-qsfp-dd-osfp-4089 ; https://www.fs.com/c/osfp-200-400-800g-4089), 8x100G PAM4:

| Description | Connector | Price (USD) |
|---|---|---|
| 800G OSFP SR8, MMA4Z00-NS, NVIDIA (ETH), 850nm, 50m, Dual MPO-12/APC, IHS/Finned Top | Dual MPO-12/APC | US$879.00 |
| 800G OSFP SR8, NVIDIA (ETH), 850nm, 100m, MPO-16/APC | MPO-16/APC | US$879.00 |
| 800G OSFP SR8, MMA4Z00-NS-FLT, NVIDIA (ETH), 850nm, 50m, RHS/Flat Top | Dual MPO-12/APC | US$1,249.00 |
| 800G OSFP DR8, Arista, 1310nm, 500m, IHS/Closed Finned Top | Dual MPO-12/APC | US$1,319.00 |
| 800G OSFP DR8, MMS4X00-NM16, NVIDIA (ETH), 1310nm, 500m, IHS/Finned Top | MPO-16/APC | US$2,249.00 |
| 800G OSFP DR8L, MMS4X00-NS, NVIDIA (ETH), 1310nm, 100m, IHS/Finned Top | Dual MPO-12/APC | US$1,119.00 |
| 800G OSFP FR8, MMS4X50-NM, NVIDIA (IB NDR), 1310nm, 2km | Dual Duplex LC/UPC | US$1,874.00 |
| 800G LPO OSFP DR8, FS, 1310nm, 500m, IHS/Finned Top (no DSP, lower power) | Dual MPO-12/APC | US$1,749.00 |
| 800G QSFP-DD SR8, Arista, 850nm, 50m | Dual MPO-12/APC (QDD-800G-2VSR) | US$1,119.00 |
| 800G QSFP-DD SR8, Cisco, 850nm, 50m | MPO-16/APC | US$1,749.00 |
| 800G QSFP-DD DR8, Cisco, 1310nm, 500m | Dual MPO-12/APC | US$1,429.00 |
| 800G QSFP-DD DR8, QDD-2X400G-DR4, Juniper, 1310nm, 500m | MPO-16/APC | US$1,819.00 |
| 800G QSFP-DD FR8, QDD-2X400G-FR4, Cisco, 1310nm, 2km | Dual Duplex LC/UPC | US$2,119.00 |
| 800G QSFP-DD XDR8, QDD-800G-2XDR4, Arista, 1310nm, 2km | Dual MPO-12/APC | US$1,999.00 |
| 800G QSFP-DD LR8, QDD-800G-2LR4, Arista, 1310nm, 10km | Dual Duplex LC/UPC | US$2,499.00 |

- **800G 2xFR4 QSFP-DD:** FS product page not directly surfaced; third-party compatible parts (AddOn/Axiom via SHI, MSRP US$5,651.50) retail ~US$3,100–3,950 [secondary] (https://www.shi.com/product/50602106/ARISTA-COMP-QSFP-DD). FS's own 800G QSFP-DD FR8 (QDD-2X400G-FR4) at US$2,119.00 is FS's closest listed equivalent [official]. **FS 2xFR4 own SKU/price — could not confirm [unverified].**
- **QDD-DR8-800G** (Cisco-compatible 800GBASE-DR8, Broadcom 7nm DSP ≤16.5W, breakout to 2x400G/8x100G), SKU 304373: **€1,458.00 excl. VAT** (EU page) [official] (https://www.fs.com/eu-en/products/304373.html).
- **Availability/lead time:** FS claims 90% same-day shipping from 7 local warehouses; 800G category shows sold quantities (e.g., 1K sold on 800G OSFP DR8 Arista, 14.5K sold on OSFP SR8) implying stocked items [official]. No explicit per-SKU lead-time text found; 800G LPO option available. Market context: 800G OSFP/QSFP-DD volume pricing roughly US$1,000–1,400 (DR8), LPO US$700–900 [independent] (invest-bud.com.pl report).

### 1.4 DAC cables

**Passive DAC:**
- 10G SFP+ passive DAC, 1m (3ft), 30AWG: **US$14.00**; 1G SFP passive DAC 0.5m: **US$14.00** [official] (https://www.fs.com/c/10-25g-dac-aoc-2869). 125.5K sold (10G 1m) [official].
- 25G SFP28 passive DAC, 1m (3ft), 30AWG: **US$33.00**; 0.5m: **US$33.00**; Juniper-coded JNP-25G-DAC-1M SKU 65860: **US$33.00**; Arista 5m SKU 65867: **US$71.00** [official] (https://www.fs.com/products/65860.html; https://www.fs.com/products/65867.html).
- 100G QSFP28 passive DAC: 0.5m £30.00 excl. VAT UK; 1m £30.00 excl.; 2m £40.00 excl.; 3m £51.00 excl. [official] (https://www.fs.com/uk/c/100g-qsfp28-dac-2873). EU 2m (SKU 65905, Dell-coded): €43.00 excl. [official] (https://www.fs.com/eu-en/products/65905.html). US: 1m Brocade SKU 50478 **US$40.00**; 2m generic SKU 74661 **US$54.00**; Cisco 2x50G breakout QSFP-100G-2QPC02 SKU 136513 **US$124.00** [official].
- 400G QSFP-DD passive DAC (customized length, 0.5–3m), SKU 82464, P/N QDD-400G-PCxxx: **€89.00 excl. VAT** (EU) [official] (https://www.fs.com/eu-en/products/82464.html). Breakouts: QDD-400G to 2xQSFP56 2m (SKU 207309): €149.00 excl. [official]; QDD-400G to 4xQSFP56 1m (SKU 139216): €129.00 excl. [official]; QDD-400G to 8xSFP56 2.5m (SKU 82463): £271.00 excl. [official]; QSFP-DD to 2xQSFP56 2.5m (QDD-400G-2QPC025): listed via FS India reseller [secondary] (https://shoptheworld.in/p/fs207311xx-2-5m-8ft-dell-compatible-400g-qsfp-dd-to-2-x-200g-qsfp56-passive-direct-attach-copper-breakout-cable).
- 800G DAC: FS category includes 1.6T/800G DAC/AOC/AEC; no individual 800G DAC SKU prices surfaced — **could not confirm [unverified]**.

**Active copper DAC (ACC):**
- 10G SFP+ Active DAC 1m: **US$58.00** [official].
- 25G SFP28 Active DAC 1m: **US$40.00** [official] (https://www.fs.com/c/10-25g-dac-aoc-2869).
- 100G QSFP28 Active DAC 1m: £117.00 excl. VAT (UK) / €125.00 excl. (EU) [official].
- 400G QSFP-DD Active DAC: P/N **QDD-400G-AC03**, SKU 177377 (Arista-coded), 3m, Macom chip ≤2.5W, 70mm bend radius: **SGD 683.43** (SG page) [official] (https://www.fs.com/sg/products/177377.html). US price **not confirmed [unverified]**.

