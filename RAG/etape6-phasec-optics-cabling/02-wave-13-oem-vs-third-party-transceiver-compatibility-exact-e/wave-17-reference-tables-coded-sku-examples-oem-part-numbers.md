---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/wave-17-reference-tables-coded-sku-examples-oem-part-numbers
title: "Wave 17 — Reference tables: coded-SKU examples, OEM part numbers, 800G catalog detail, retailer price matrices, programming workflow (research date 2026-09-22)"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["EU", "Nvidia", "United States"]
dates: ["2026-09-22"]
keywords: ["research", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2831, 2888]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 4e5df804d31feb3ff87542aa4b25f43387bff459c62d2da2ccd1e309a93faf9d
---

# Wave 17 — Reference tables: coded-SKU examples, OEM part numbers, 800G catalog detail, retailer price matrices, programming workflow (research date 2026-09-22)

## Wave 17 — Reference tables: coded-SKU examples, OEM part numbers, 800G catalog detail, retailer price matrices, programming workflow (research date 2026-09-22)

**Method:** editorial synthesis wave — facts restated from Waves 1–16 with original provenance; no new web searches. Dense reference tables for RAG retrieval.

### 17.1 Coded-SKU examples — one P/N, many OEM codes (FS.com)

Illustrates how a single third-party P/N is sold under per-OEM SKUs (see Wave 1 §1.1) [official]:

| Third-party P/N | OEM code (SKU) | Speed/type | Price observed |
|---|---|---|---|
| QSFP-SR4-40G | Generic (17931) | 40GBASE-SR4 | US$43.00 |
| QSFP-SR4-40G | HPE (36848) | 40GBASE-SR4 | US$43.00 (category) |
| QSFP-SR4-40G | Juniper (36439) | 40GBASE-SR4 | US$43.00 (category) |
| QSFP-SR4-40G | Gigamon (75298) | 40GBASE-SR4 | US$43.00 (category) |
| QSFP-LR4-40G | Generic (24422) | 40GBASE-LR4 10km | US$299.00 |
| QSFP-LR4-40G | Arista (36202) | 40GBASE-LR4 10km | US$299.00 |
| QSFP-LR4-40G | Dell (36696) | 40GBASE-LR4 10km | US$299.00 |
| QSFP-LR4-40G | Custom (48564) | 40GBASE-LR4 10km | US$299.00 |
| QSFP-LR4-40G | APRESIA (75299) | 40GBASE-LR4 10km | US$299.00 |
| QSFP-100G-LR4 | Generic (104861) | 100GBASE-LR4 10km | US$399.00 |
| QSFP-100G-LR4 | Arista (104847) | 100GBASE-LR4 10km | US$399.00 |
| QSFP-100G-LR4 | HW (104863) | 100GBASE-LR4 10km | US$399.00 / €399.00 excl. VAT (EU) |
| QSFP-100G-LR4 | H3C (104862) | 100GBASE-LR4 10km | US$399.00 |
| QSFP-100G-SR4-S | Cisco (104861) | 100GBASE-SR4 | US$99.00 (113.8K sold) |
| QDD-DR4-400G | HW/Dell (234483 / 183428) | 400G DR4 500m | US$549.00 |
| QDD-DR4-400G-Si | Cisco (128242) | 400G DR4 500m SiPh | US$749.00 |
| QDD-400G-SR8 | Arista | 400G SR8 100m | £168.00 excl. VAT (UK) / US$219.00 (US) |
| SFP-10GSR-85 | Generic (11552) / Arista (36982) | 10GBASE-SR | US$25.00 / €19.00 excl. VAT (EU) |

[official] — https://www.fs.com product/category pages (see Wave 1). Note the same-SKU collision flagged in Wave 1 §1.1 (SKU 104861 returned for both SR4 and LR4 pages) — treat SKU-level joins cautiously.

### 17.2 OEM part-number reference examples (genuine parts cited in compatibility contexts)

| OEM | Part number | Description | Seen in context |
|---|---|---|---|
| Cisco | QSFP-100G-SR4-S | 100GBASE-SR4 QSFP28 | reseller MSRP US$2,440.10 [vendor-reported] |
| Cisco | QDD-400G-DR4-S | 400G QSFP-DD DR4 | reseller "list" US$5,195.85 [vendor-reported] |
| Cisco | GLC-SX-MM / GLC-SX-MMD | 1G SFP SX | Ubiquiti EdgeSwitch tested list [official] |
| Cisco | SFP-H10GB-CU1M | 10G DAC 1m | Ubiquiti EdgeSwitch tested list [official] |
| Cisco | GLC-T | 1G RJ45 SFP | common reference |
| Juniper | JNP-QSFP-100G-SR4 | 100GBASE-SR4 | FS.com Juniper-coded US$99.00 [vendor-reported] |
| Juniper | JNP-QSFP-100G-PSM4 | 100G PSM4 | Juniper HCT example entry [official] |
| Juniper | JNP-SFP-10G-BX40U | 10G BX40-U | Juniper HCT example entry [official] |
| Juniper | OSFP-800G-DR8-P | 800G OSFP DR8 | Axiom equivalent US$2,866 [vendor-reported] |
| NVIDIA/Mellanox | MMA1B00-C100D | 100G QSFP28 SR4 | FS/Approved compatible US$99.00 [vendor-reported/official] |
| NVIDIA/Mellanox | MMA1B00-E100 | 100G QSFP28 (NVIDIA-era) | Approved Networks list [official] |
| NVIDIA/Mellanox | MMA1T00-VS | 200G QSFP56 SR4 | FS.com US$299.00 [official] |
| NVIDIA/Mellanox | MMA4Z00-NS400 | 400G OSFP SR4 | FS.com US$769.00 [official] |
| NVIDIA/Mellanox | MMS1X00-NS400 | 400G QSFP112 DR4 | FS.com US$1,119.00 [official] |
| NVIDIA/Mellanox | MMA1Z00-NS400 | 400G QSFP112 SR4 | FS.com US$1,119.00 [official] |
| NVIDIA/Mellanox | MMS4X00-NS | 800G OSFP DR8 | QSFPTEK compatible US$799 [vendor-reported] |
| Dell | Q56DD-400G-SR4.2-ON-C | 400G QSFP-DD SR4.2 (ProLabs Cisco-style SKU for Dell) | ProLabs NA stock 229 [official] |
| Dell | FTLF series | 10G SFP+ (Finisar-based) | Ubiquiti EdgeSwitch tested list [official] |
| HPE/Aruba | JL274A | 100G QSFP28 SR4 | Approved Networks compatible list [official] |
| HPE | J4858C / J4859B | 1G/10G SFP | Ubiquiti EdgeSwitch tested list [official] |
| MikroTik | S+85DLC03D / S+31DLC10D | SFP+ modules | Ubiquiti EdgeSwitch tested list [official] |
| Arista | (various QDD- codes) | 400G QSFP-DD | FS/ProLabs coded SKUs [vendor-reported/official] |

