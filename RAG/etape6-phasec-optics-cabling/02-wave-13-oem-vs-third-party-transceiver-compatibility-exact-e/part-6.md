---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/part-6
title: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior (part 6)"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Intel", "Nvidia"]
dates: ["2020-12"]
keywords: ["cost", "intel", "nvidia", "optics", "pricing", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2504, 2515]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 7e824cbe1e96661de5b05340a9de0331cf060e97397c83c9067fea9f4865cdfd
---

# Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior (part 6)

- **Identity:** "Approved Networks", self-described as "a proud brand of Legrand in the Data, Power, and Control Division"; "Approved Optics" is not an official name [official] (https://approvednetworks.com/about-us/).
- **History/profile:** formed from three third-party brands — Approved Networks, Champion ONE, U.S. Critical — acquired by Legrand in **December 2020** and consolidated under one brand; claims "30 years" of cost-effective optical solutions and 10,000+ customers in 40+ countries [official] (https://approvednetworks.com/about-us/). Portfolio: transceivers, DACs, AOCs, passive solutions; 100G/200G/400G/800G including QSFP-DD SR4/SR4.2/SR8/DR4/DR4+/FR4/LR4/ZR/ZR+ and OSFP [official] (https://approvednetworks.com/transceivers/).
- **Founding-year conflict (flagged):** official page gives no founding year; secondary sources conflict — Champion ONE founded 1992, HQ Cleveland, Ohio [secondary] (https://leadiq.com/c/approved-networks-a-brand-of-legrand-formerly-champion-one/5a1d959f2300005200844eac) vs. Approved Networks, Inc. founded 2009, Independence, Ohio [secondary] (https://Rocketreach.co/approved-networks-a-brand-of-legrand-profile_b5fe787cf42d35d6). Treat founding year as **unresolved**; official history emphasizes the 2020 Legrand consolidation.
- **Coding method (universal reprogramming):** **OnePort programmer** — "code/recode transceivers, DACs, and AOCs for over 90 OEM systems, including: Arista, Cisco, Juniper Networks, Avaya, Hewlett Packard, Nokia, Intel Adtran, and Ciena"; data rates 1G–100G; batch-coding and favorite lists by OEM/part number; multi-code DACs/AOCs [official] (https://approvednetworks.com/blog/the-new-oneport-is-here/). Products also sold pre-coded per OEM ("OP-" OnePort programmable SKUs) [official] (https://approvednetworks.com/products/10gbase-lr-sfp-1310nm-40km-ddm-transceiver.html).
- **Compatibility claim:** "100% guaranteed compatible in over 90 OEM platforms" [official] (https://approvednetworks.com/products/100gbase-sr4-qsfp28-850nm-100m-ddm-transceiver.html?sku=SM100G-SR-A). Confirmed on the official 100G SR4 page: Cisco, Arista, Dell, HP/HPE (incl. JL274A), Juniper, **Mellanox (MMA1B00-C100D)**, **NVIDIA-era Mellanox (MMA1B00-E100)**, plus MSA OnePort Programmable variants [official] (same URL). **Ubiquiti and MikroTik compatibility: not found** on official Approved Networks lists.
- **Warranty (exact official wording):** "Industry Leading Lifetime Warranty on all transceiver products" [official] (same URL); also "Limited Lifetime Warranty" [official] (https://approvednetworks.com/transceivers/). **Exact exclusions: not found** in public official sources.
- **Standards compliance:** example QSFP28 SR4 datasheet cites IEEE Std 802.3bm (100GBASE-SR4/40GBASE-SR4), QSFP28 MSA, SFF-8636 management interface [official] (https://approvednetworks.com/content/data sheets/transceivers/QSFP28/QSFP28-SR4-I.pdf). Broader claim: "100% tested to exact MSA & OEM specifications" [official] (product pages).
- **QA/testing claims:** "99.98% Reliability Rate"; every unit "environmentally tested in its specific port/platform" (distance, wavelength, traffic density, light verification); partners with "the world's largest Tier 1 optical component contract manufacturers"; Serial Number Tracking System [official] (https://approvednetworks.com/transceivers/ and product pages). About page: "60,000 Units in Stock", ">98% On Time Delivery Rate" [official] (https://approvednetworks.com/about-us/).
- **Prices:** not public — official shop states "For pricing, call your Approved Networks sales partner, or Request a Quote" [official] (https://approvednetworks.com/optical-transceivers/). **Quote-only model.** Warehouse locations: not found.

#### 14.2.2 10Gtek

