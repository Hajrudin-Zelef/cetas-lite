---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/part-9
title: "Supplementary / Complementary Research Pass #2 — September 22, 2026 (part 9)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2025-04", "2026-02-10", "2026-04-24", "2026-06", "2026-07", "2026-07-01"]
keywords: ["cost", "distribution", "latency", "memory", "optics", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1177, 1201]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 7d05cbaade7587c9ee4fb6a0ba596f08e3379080d920e0d220f0e90a117c826f
---

# Supplementary / Complementary Research Pass #2 — September 22, 2026 (part 9)

- **Memory surcharge**: Ubiquiti began applying a separate-line-item **"Memory Surcharge" of up to 5.8%** on select products effective **April 24, 2026**, attributed to global memory/storage market volatility; Ubiquiti says it absorbs part of the cost. The affected-SKU list was **revised July 1, 2026** (per reseller Streakwave). **[secondary]** https://getuniqcli.com/news/ubiquiti-memory-surcharge-2026 · **[vendor-reported]** https://blog.streakwave.com/ubiquiti-memory-surcharge-notice-april-2026 · https://blog.streakwave.com/hubfs/Ubiquiti/2026/Notice-Ubiquiti-Memory-Surcharge-July-2026-with-Tariffs.pdf
- Scope per NASCompares explainer: cloud gateways, routers, Wi-Fi APs, switches, Protect cameras/NVRs, Talk devices, selected NAS/enterprise hardware. **[secondary]** https://nascompares.com/news/unifi-memory-surcharge-explained-ubiquiti-ram-storage-costs/
- Observed instances at UI Store: Enterprise Campus Switch Core 32×100G $5,393 **incl. $394** memory surcharge **[independent]** https://www.notebookcheck.net/Ubiquiti-introduces-Enterprise-Campus-Switch-Core-with-3-2-Tbit-s-throughput.1389733.0.html; EF-Core $3,499 → **$3,775 surcharge incl.** ($276) **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core
- **Tariff surcharge**: consumers report ~7.5% tariff surcharge on many products after April 2025 IEEPA "Liberation Day" tariffs; the US Supreme Court held those tariffs unlawful and importers are being refunded — a class action, **Higgins v. Ubiquiti Inc., No. 1:26-cv-00659 (D. Del.)**, filed by Chimicles Schwartz Kriner & Donaldson-Smith, seeks recovery of tariff surcharges for UI Store and distributor purchasers, alleging Ubiquiti has not refunded them. **[secondary — law-firm announcement; status of case unverified]** https://chimicles.com/ubiquiti-tariff-surcharge-refund-class-action/
  - Flag: this may overlap with the already-covered lawsuits section — dedupe if consolidating.
- Retailer stock snapshot (Sep 2026): EF-Core **sold out** at Baltic Networks **[vendor-reported]** https://www.balticnetworks.com/en-ca/collections/ubiquiti-networks?limit=288; UISP-FIBER-OLT-XGS **low stock** at MiRO SA (JHB: 2, CPT: 1) and **unavailable** at JW.com.au **[vendor-reported]** https://miro.co.za/12-fibre-optics-ubiquiti-ufiber-gpon-olt/7105-ubiquiti-uisp-xgs-olt-8-port-uisp-fiber-olt-xgs-810084691175.html · https://www.jw.com.au/product/ubiquiti-uisp-fiber-olt-xgs-2048-cc-25g-sfp28; Adorama lists USW-Enterprise-24-PoE as special order, estimated ship **10/02/2026** **[vendor-reported]** https://www.adorama.com/ubiquiti-unifi-enterprise-24-poe-managed-switch-400w-2x-sfp/p/ubuswenter; wisp.net.au shows mixed "in stock / last items / backorder" across the UniFi range **[vendor-reported]** https://wisp.net.au/s/15/ubiquiti-network-switches/

#### A.7. UniFi Pro AV line — 2026 snapshot (brief)

- **EAV-Fiber** (Enterprise Audio/Video Fiber switch): L3, 1× GbE RJ45 + 20× 10G SFP+ + 2× 100G QSFP28, 400 Gbps non-blocking / 800 Gbps switching / 595 Mpps, timing engine for PTP / SMPTE ST-2110 / SDVoE / AES67 with OCXO clock and GPS grandmaster input, managed in UniFi Network alongside EAV-Bridge endpoints. Price **$4,259.00**; sold out at Baltic Networks (crawled ~164 days before Sep 22). **[vendor-reported]** https://arcip.com.au/shop/eav-fiber-ubiquiti-unifi-eav-fiber-enterprise-audio-video-1u-aggregation-switch-with-20-x-10g-sfp-and-2-x-100g-qsfp28-800-gbps-switching-ptp-smpte-st-2110-sdvoe-aes67-timing-engine-ocxo-clock-gps-grandmaster-input-9724 · https://www.balticnetworks.com/en-ca/collections/ubiquiti-networks?limit=288
- **EAV-Bridge**: AV endpoint supporting one-to-many AV distribution (details unverified beyond product family mention). **[unverified]** https://www.youtube.com/watch?v=wBk8fXy78Y0
- 2026 shows: ISC West 2026 booth showed "black switches designed for high-end audio and SFP connectivity" under Pro AV **[independent]** https://www.youtube.com/watch?v=m61LhBlqTIQ; UWC London 2026 gave a closer look at the Pro AV lineup incl. bridge **[independent]** https://www.youtube.com/watch?v=wBk8fXy78Y0; ASBIS distributor webinar/page on "UniFi Enterprise Audio Video Switching" (deterministic timing, predictable latency) **[vendor-reported]** https://news.asbis.com/news/?s=ubiquiti

#### A.8. Open uncertainties (Pass A)

1. Whether any 400G UniFi hardware or roadmap exists — not found anywhere; absence of evidence only. **[unverified]**
2. The "UOLT-4"/"UONU-XG" SKUs from the brief — not found under those names; all PON gear found is UISP-branded. **[unverified]**
3. WaveFiber ONU exact launch date — stocked at resellers in Sep 2026; no launch announcement found. **[unverified]**
4. STH/LTT-specific 2026 endorsements — not surfaced in this pass. **[unverified]**
5. EF-Core general availability date — "available now" ~mid-June 2026 (Dong Knows), distributor announcement ~mid-July 2026; no single official launch-date stamp found. **[unverified exact date]**
6. Higgins v. Ubiquiti class-action status (filed vs. certified vs. dismissed) — not verified. **[unverified]**


---

