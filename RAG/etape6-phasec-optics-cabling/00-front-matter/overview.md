---
id: etape6-phasec-optics-cabling/00-front-matter/overview
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "EU", "Nvidia", "United States"]
dates: ["2026-09-22"]
keywords: ["optics", "cpo", "dsp", "lpo", "npo", "nvidia", "pricing", "research"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1, 66]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: d7b717bfed5081ca84009f34c01332a9d5799b19f251ca573909eb0b2a54b760
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure

**Scope:** FS.com catalog & pricing; optical transceiver landscape (SFP → OSFP-XD, 400G/800G standards); DSP vs LPO vs CPO/NPO; DAC vs AOC vs AEC; fiber infrastructure (OS2/OM3–OM5, MPO, polarity); copper structured cabling (Cat6A/Cat8); third-party coded optics compatibility; 800G ecosystem readiness.
**Research date:** 2026-09-22. All prices, specs and availability verified against sources cited as of this date.
**Method:** read-only web research (browser_search, browser_open). No live-browser visits, no forms, nothing sent externally. No identifiers guessed; no SKUs, URLs or numbers invented.
**Provenance legend:** every factual claim carries one of `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`.
**File discipline:** append-only. Sections are added wave by wave; earlier sections are never rewritten or deleted. Gaps, conflicts, unverified claims and non-comparable figures are flagged explicitly.

---

---

## Wave 1 — FS.com catalog & pricing (research date 2026-09-22)

**Method:** read-only web research via browser_search / browser_open of fs.com product and category pages. Note: FS.com shows live prices on its site; prices below are as captured from indexed page snapshots, so any individual price may have changed since the snapshot — where the index showed two snapshots of the same page with different prices, this is flagged. Fact labels: **[official]** = FS.com's own site/page text; **[secondary]** = third-party retailer carrying FS products; **[independent]** = third-party market commentary; **[unverified]** = not confirmed.

### 1.1 Optical transceivers by form factor/speed

#### 10G SFP+
- **SFP-10GSR-85** (Cisco SFP-10G-SR compatible), SKU 11552 (generic page SKU; Arista-coded variant SKU 36982): 10GBASE-SR, 850nm, 300m OM3/DOM, duplex LC/UPC, ≤1W. **US$25.00** [official] (https://www.fs.com/products/36433.html ; https://www.fs.com/products/36982.html). EU: €19.00 excl. VAT (SKU 11552, SKU EU 21146 page: https://www.fs.com/eu-en/products/21146.html) [official]. UK category list: £20.40 incl. VAT / £17.00 excl. [official] (https://www.fs.com/uk/c/10g-sfp-plus-63). 1.4M sold, 2.5K reviews on US category page [official] (https://www.fs.com/c/1-2-4g-modules-57).
- **SFP-10G-LR** (Cisco SFP-10G-LR compatible): 1310nm, 10km, duplex LC/UPC, ≤1W. **US$34.00** [official]; UK £27.60/£23.00 excl. [official] (same URLs).
- 10GBASE-T SFP+ (Cisco compatible): 30m over Cat6a, ≤2.5W. **US$86.00** [official] (https://www.fs.com/c/1-2-4g-modules-57).
- 1G SFP: SFP-1G-SX 550m **US$8.80**; SFP-1G-LH 10km **US$10.00**; GLC-T RJ45 **US$35.00** [official] (same URL).

#### 25G SFP28
- **SFP-25G-SR-S** (Cisco compatible): 25GBASE-SR, 850nm, 100m, duplex LC/UPC, ≤1W. **US$49.00** [official] (https://www.fs.com/c/1-2-4g-modules-57).
- US category page also lists 50G QSFP28/SFP56 and 100G SFP112 categories [official].

#### 40G QSFP+
- **QSFP-SR4-40G** (40GBASE-SR4, Cisco QSFP-40G-SR4-S compatible), SKU 17931: 850nm, 150m, MPO-12/UPC, ≤1.5W, Macom chip. **US$43.00** [official] (https://www.fs.com/products/17931.html). EU (EU-site SKU 36361): €39.00 excl. VAT [official] (https://www.fs.com/eu-en/products/36361.html). 140.5K sold [official] (https://www.fs.com/c/qsfp-40g-transceivers-1360).
- **QSFP-LR4-40G** (40GBASE-LR4, 10km, 1310nm, duplex LC): **US$299.00** — SKU 24422 (generic), 36202 (Arista), 36696 (Dell), 48564 (custom), 75299 (APRESIA) — all US$299.00 [official] (https://www.fs.com/products/36202.html; https://www.fs.com/products/24422.html). 39.3K sold [official].
- QSFP-40GE-LR4-S 4x10G-LR PLR4 10km MPO-12/APC: **US$359.00** [official].
- QSFP-40G-SR-BD 300m (SWDM4, duplex LC): **US$369.00** [official].
- QSFP-40G-PLRL4 2km MPO-12/APC: **US$309.00** [official].
- QSFP-LR4L-40G (Cisco WSP-Q40GLR4L compatible, 2km, duplex LC), SKU 36172: **US$289.00** [official].
- QSFP-LX4-40G (Dell 407-BBRC, 150m OM3 / 2km OS2), SKU 36491: **US$389.00** [official].
- QSFP-40G-UNIV (LX4/CWDM4 2km): **US$389.00** [official]; QSFP-40G-CSR4 400m MPO: **US$81.00** [official] (https://www.fs.com/c/qsfp-40g-transceivers-1360).

#### 100G QSFP28
- **QSFP-100G-SR4-S** (Cisco QSFP-100G-SR4-S compatible), SKU 104861: 4x25G NRZ, 850nm, 100m, MPO-12/UPC, ≤2.5W, VCSEL. **US$99.00**; 113.8K sold, 98 reviews [official] (https://www.fs.com/products/104861.html).
- **QSFP-LR4-100G** (100GBASE-LR4, 1310nm, 10km, duplex LC/UPC, ≤3.5W, Semtech), SKU 104861 generic: **US$399.00** [official] (https://www.fs.com/products/104861.html — note: search returned same SKU for LR4 page; Arista variant SKU 104847 US$399.00; HW variant SKU 104863; H3C SKU 104862 — all US$399.00) [official]. EU (SKU 104863): €399.00 excl. VAT [official] (https://www.fs.com/eu-en/products/104863.html). 40.8K sold [official] (https://www.fs.com/c/100g-qsfp28-sfp-dd-1159).
- QSFP-100G-CWDM4-S (2km, duplex LC): **US$209.00** [official].
- QSFP-100G-SL4 (850nm 30m MPO-12/UPC): **US$74.00** [official].
- QSFP-100G-SWDM4 (100m duplex LC): **US$419.00** [official].
- QSFP-40/100-SRBD dual-rate SWDM4: **US$549.00** [official].
- QSFP-100G-ER4L-S (40km): **US$1,599.00** [official].
- QSFP-100G-SR4-I industrial (≤2W, Broadcom VCSEL): **US$349.00**; 100G ZR4 80km: **US$6,499.00**; BX40 40km: **US$2,369.00** [official] (https://www.fs.com/c/qsfp28-100g-transceivers-1159).

#### 200G QSFP56 / QSFP-DD
- MMA1T00-VS (NVIDIA ETH compatible) 200G QSFP56 SR4 4x50G PAM4, 850nm, 100m, MPO-12/UPC: **US$299.00**; 4.8K sold [official] (https://www.fs.com/c/200-400-800g-modules-3859).
- Q28DD-200G-2SR4 (Dell compatible) 200G QSFP-DD SR4 8x25G NRZ, 850nm, 100m, MPO-24/UPC: **US$661.00** [official].

#### 400G QSFP-DD / QSFP112 / OSFP
Category page [official] (https://www.fs.com/uk/c/400g-osfp-qsfp112-qsfpdd-3652; prices from UK listing, VAT incl / excl.):
- QDD-400G-DR4-S (Cisco compatible), 1310nm, 500m, MPO-12/APC: **£505.20 / £421.00 excl.** [official] (UK).
- US price for Cisco QDD-DR4-400G-Si (SKU 128242, Broadcom 7nm DSP, ≤10W): **US$749.00** [official] (https://www.fs.com/products/128242%20.html).
- QDD-DR4-400G (HW/Dell compatible), SKU 234483 / 183428, Marvell DSP ≤9W, MPO-12/APC: **US$549.00** [official] (https://www.FS.COM/products/234483.html; https://www.fs.com/products/183428.html).
- EU price for QDD-DR4-400G-Si (SKU 128242): **€712.81 incl. / €599.00 excl.** in one snapshot and **€677.11 incl. / €569.00 excl.** in a second snapshot — **[official] but conflicting across snapshots (time-varying; treat as approximate €569–599 excl.)** (https://www.fs.com/eu-en/products/128242.html).
- QDD-400G-FR4 (Arista compatible), 1310nm, 2km, duplex LC/UPC: **£667.20 / £556.00 excl.** [official].
- QDD-400G-SR4 (Arista), 850nm, 50m, MPO-12/APC: **£667.20 / £556.00 excl.** [official].
- QDD-400G-LR4 (Arista), 1310nm, 10km, duplex LC/UPC: **£808.80 / £674.00 excl.** [official].
- QDD-400G-SR8 (Arista), 850nm, 100m, MPO-16/APC: **£201.60 / £168.00 excl.**; 7K sold [official].
- QDD-400G-SR4.2-BD (Cisco), 850nm, 100m, MPO-12/UPC: **£1,296.00 / £1,080.00 excl.** [official].
- 400G OSFP SR4 MMA4Z00-NS400 (NVIDIA ETH), 850nm, 50m, MPO-12/APC, flat top: **US$769.00** [official]; UK £708.00 [official].
- InfiniBand NDR prices [official] (https://www.fs.com/c/800g-ndr-infiniband-3801): 400G OSFP DR4 (NVIDIA IB), 1310nm, 500m: **US$1,124.00**; 400G OSFP DR4L (100m): **US$869.00**; 400G QSFP112 DR4 MMS1X00-NS400: **US$1,119.00**; 400G QSFP112 SR4 MMA1Z00-NS400: **US$1,119.00**; 400G QSFP112 FR4 2km: **US$1,619.00**.

