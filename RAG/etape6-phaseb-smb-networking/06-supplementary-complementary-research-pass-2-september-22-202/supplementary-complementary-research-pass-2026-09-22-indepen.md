---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/supplementary-complementary-research-pass-2026-09-22-indepen
title: "Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["EU", "United States"]
dates: ["2025-04", "2026-02-10", "2026-04-24", "2026-06", "2026-07", "2026-07-01", "2026-09-22"]
keywords: ["research", "advisory", "agents", "claude", "cost", "distribution", "ethernet", "latency", "license", "memory", "optics", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1127, 1201]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 00b162f56fef8d09c9abf1db65cdcdc8323fde83d4f40c15f41de967950234c6
---

# Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)

## Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)

**Scope note:** This section contains material from three independent research passes run on 2026-09-22 against gaps in the original Parts 1–2. It is distinct from the sibling S1–S18 and R1–R6 supplementary waves added separately to this file; some topics overlap across waves (flagged where material), but this pass's research is independent and adds new facts, prices, and firmware details. Parts 1–2 above are untouched. Every fact carries a provenance tag: [official] / [vendor-reported] / [independent] / [secondary] / [unverified]. Identifiers, dates, specs, prices and verbatim source URLs are given; uncertainties are flagged explicitly.

---

### Pass A — Ubiquiti deep gaps: 400G, fiber/PON, EF-Core launch, software window Sep 10–22, homelab consensus, supply/surcharges, Pro AV

Research conducted September 22, 2026; all facts dated on/before that cutoff. **[official]** = ui.com/Ubiquiti-owned; **[vendor-reported]** = reseller/distributor specs; **[independent]** = press/reviewer/institutional; **[secondary]** = aggregator/analysis; **[unverified]** = could not be confirmed.

#### A.1. UniFi 400G (QSFP-DD) status

- No UniFi switch with 400G (QSFP-DD) ports was announced, launched, or road-mapped on/before Sep 22, 2026, in any source searched. **[unverified — absence of evidence; not proof of absence]**
- The fastest UniFi ports found in-market remain 100G QSFP28: Enterprise Campus Switch Core (32× QSFP28, $5,393 incl. $394 memory surcharge), EAV-Fiber (2× 100G QSFP28), EF-Core gateway (4× 100G QSFP28). **[independent]** https://www.notebookcheck.net/Ubiquiti-introduces-Enterprise-Campus-Switch-Core-with-3-2-Tbit-s-throughput.1389733.0.html · **[vendor-reported]** https://arcip.com.au/shop/eav-fiber-ubiquiti-unifi-eav-fiber-enterprise-audio-video-1u-aggregation-switch-with-20-x-10g-sfp-and-2-x-100g-qsfp28-800-gbps-switching-ptp-smpte-st-2110-sdvoe-aes67-timing-engine-ocxo-clock-gps-grandmaster-input-9724
- Searched: "Ubiquiti UniFi 400G QSFP-DD switch announced 2026", "store.ui.com 400G QSFP-DD switch UniFi", community/Reddit threads — no hits on a UniFi 400G product or roadmap statement. **[unverified]**

#### A.2. Fiber / PON 2026: UISP line (no UniFi-branded OLT/ONU found)

- **No UniFi-branded OLT or ONU (no "UOLT-4" or "UONU-XG" SKUs) was found** as of Sep 22, 2026 — all Ubiquiti PON products found are UISP/UFiber-branded, managed by the UISP app (not the UniFi Network app). The example SKUs from the brief could not be verified; treat them as unconfirmed. **[unverified]**
- **UISP-FIBER-OLT-XGS** — 8× XGS-PON ports, 4× 25G SFP28 uplinks, 2048-client capacity (256:1 split ratio), 1U rackmount, 2× hot-swappable redundant AC/DC PSUs, max 70 W, managed via UISP app 1.5.7+. MSRP **$2,499**. **[vendor-reported]** https://www.publicsector.shidirect.com/product/47460833/Ubiquiti-UISP-Fiber-OLT-XGS · https://hw.ddfu.org/en/printout/printproduct?StiId=293801 · https://www.jw.com.au/product/ubiquiti-uisp-fiber-olt-xgs-2048-cc-25g-sfp28
- **UISP-FIBER-XGS** (XGS-PON ONU/ONT) — 1× SC/APC XGS-PON (10 Gbps symmetric) + 1× 10GbE RJ45 (100/1G/2.5G/5G/10G), bridge and router modes, USB-C power, managed via UISP app 1.4.9+; works **only** with Ubiquiti OLTs. Street pricing: €120.65 ex-VAT (PL) **[vendor-reported]** https://en.cdr.pl/p9664,ubiquiti-uisp-fiber-xgs-uisp-fiber-xgs-pon-ont-terminal-1x-10ge.html · £100.79 inc VAT, 18 in stock (UK) **[vendor-reported]** https://linitx.com/category/ubiquiti-gpon/1267 · €150.90 (PL) / R3,347.65 inc VAT (ZA) **[vendor-reported]** https://interprojekt.pl/en/p/ubiquiti-uisp-fiber-xgs-eu.html · https://miro.co.za/12-fibre-optics-ubiquiti-ufiber-gpon-cpe/7103-ubiquiti-uisp-fiber-xgs-gpon-onu-ont-810084691236.html
- **UISP-FIBER-XG** (XG-PON ONU/ONT) — 10 Gbps down / 2.5 Gbps up, 1× 2.5GbE RJ45 + 24V PoE-in. £93.59 inc VAT, 24 in stock (UK). **[vendor-reported]** https://linitx.com/category/ubiquiti-gpon/1267
- **Wave-Fiber-ONU** (new "WaveFiber" GPON CPE line) — 1× SC/APC GPON WAN (ITU-G.984) + 1× 1/2.5 GbE RJ45, 24V passive PoE-in (5W max), 76.5×76.5×26.4 mm / 78 g, desktop/wall mount, OMCI-managed via UISP (remote provisioning, firmware upgrades, GEM port encryption, NAT/firewall/VLAN). £47.21 (UK, in stock, "Despatched today"); 5-pack **Wave-Fiber-ONU-5** £230.09, 38 in stock. **[vendor-reported]** https://linitx.com/product/ubiquiti-wavefiber-gpon-cpe-optical-network-unit-wave-fiber-onu/18297 · https://www.firstshop.co.za/collections/wireless-adapters/products/ubiquiti-uisp-wavefiber-onu-sc-apc-2-5gbps-ethernet-gpon-optical-network-media-converter-wave-fiber-onu-372469 · https://www.getic.com/product/ubiquiti-wavefiber-onu
  - Flag: the WaveFiber ONU is widely stocked at resellers in Sep 2026 and looks like a 2026 addition, but no official launch announcement or date was found — exact launch date **[unverified]**.
- Accessories spotted at UK retail: **UACC-UF-OM-XGS** (XGS/XG optical transceiver, 10 Gbps, 20 km) £131.99 — awaiting restock; **UACC-UF-WDM-XGS** (XGS/XG-PON + GPON coexistence WDM filter) £238.22, 9 in stock. **[vendor-reported]** https://linitx.com/category/ubiquiti-gpon/1267
- Legacy GPON hardware still sold in 2026: **UF-OLT** (8× GPON, 2× SFP+, 1024 ONUs, MSRP $1,499 at 2017 launch) and **UF-OLT-4** (~$1,190–$1,291), UF-Nano / UF-LOCO ONUs, managed via UISP/UNMS. **[vendor-reported]** https://globenewswire.com/news-release/2017/07/31/1065075/0/en/Ubiquiti-Networks-Launches-Plug-and-Play-UFiber-GPON-Platform.html · https://www.howardcomputers.com/accessories/detail.cfm?id=S20654584

#### A.3. Enterprise Firewall Core (EF-Core, "EFG Core" 4×100G) — launched, in market

- The "EFG Core" previewed at UWC is officially the **Enterprise Firewall Core (EF-Core)** and is **launched and orderable on the UI Store**: **$3,499.00** ($3,775.00 surcharge incl.); UK store **£3,005.00** (£3,606.00 VAT incl.). **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core · https://uk.store.ui.com/uk/en/products/ef-core
- Specs (official + vendor-reported): 24 hyperscale-class Arm Neoverse N2 cores @ 2.5 GHz (ARM v9), 32 GB RAM, 128 GB SSD; ports = 4× 100G QSFP28 + 4× 25G SFP28 + 8× 10GbE RJ45 + 2× 1GbE RJ45 + 1 management + 1 console (default WAN = 1× 100G, 1× 25G, 1× 10GbE); 2× hot-swappable 550W AC/DC PSUs; 1U rackmount; managed through UniFi Network; **79 Gbps IDS/IPS**, 61 Gbps full SSL inspection, 22,500+ clients, 10M concurrent sessions, 120K new sessions/sec, 5,000+ concurrent IPsec/WireGuard tunnels, 38 Gbps aggregate IPsec; shadow-mode/VRRP HA; license-free SD-WAN. **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core · **[vendor-reported]** https://news.asbis.com/news/suppliers/ubiquiti-introduces-enterprise-firewall/ · https://aceperipherals.com/products/ubiquiti-ef-core-enterprise-firewall-core-100gbps-cloud-gateway-with-high-availability
- Timeline: reviewed as "available now" ~mid-June 2026 **[independent]** https://dongknows.com/ubiquiti-enterprise-firewall-core-ef-core-review/; ASBIS (distributor) announced availability ~70 days before Sep 22, 2026 (i.e., ~mid-July 2026) **[vendor-reported]** https://news.asbis.com/news/suppliers/ubiquiti-introduces-enterprise-firewall/; shown on the show floor at UWC London 2026 **[independent]** https://www.youtube.com/watch?v=wBk8fXy78Y0
- Paid add-ons at UI Store checkout: CyberSecure Enterprise (Proofpoint threat intelligence) $499.00/unit billed annually; 5-year coverage/replacement $699.00/unit; 90 days of professional phone support included in US/CA/EU/UK. **[official]** https://store.ui.com/us/en/category/cloud-gateways-enterprise-scale/products/ef-core?utm_source=web&utm_medium=whatsnew&utm_campaign=1carousel&utm_content=ef-core
- Availability Sep 2026: **sold out** at Baltic Networks ($4,969.00). **[vendor-reported]** https://www.balticnetworks.com/en-ca/collections/ubiquiti-networks?limit=288

#### A.4. UniFi software, Sep 10 → Sep 22, 2026 (window check)

- **No Network Application release found after 10.6.106 (Sep 10, 2026).** Third-party package trackers (Home Assistant add-ons, FreeBSD net-mgmt/unifi10 port) still list **10.6.106** as current as of ~4–12 days before Sep 22, 2026. **[secondary]** https://github.com/zglate/addon-unifi/blob/HEAD/README.md · https://postgoo.com/posts/6aae49d13acbd737ce601ad4?type=article · https://github.com/hassio-addons/repository-edge/blob/HEAD/unifi/CHANGELOG.md
- **No 10.7 or 11.0** Network release notes or early-access builds found in the Sep 10–22 window. **[unverified — absence of evidence]**
- No UniFi OS release found dated Sep 10–22, 2026. Most recent security baselines cited by third parties in this period: UniFi OS **5.1.12** for most supported devices, **5.1.11** for UDM-Beast, **5.1.10** for UNAS family, UniFi OS Server **5.0.8** (per Vanuatu CERT Advisory 152, citing Ubiquiti Security Bulletin 064). **[independent]** https://cert.gov.vu/images/publications/2026/Advisory_152_Ubiquiti%20UniFi%20OS%20Improper%20Input%20Validation%20Vulnerability.pdf
- 10.6 feature highlights circulating in the period (for context): Topology Spotlight, Time Machine for radios + ports, Channel AI nightly automation, port locking, expanded SafeOps, HA readiness score; Port Locking + Multicast Suppressor noted as "coming soon to Early Access" (requires UAP 8.8+ / switch firmware 7.6+). **[secondary]** https://www.youtube.com/shorts/HZd4vj0d-1Q · https://www.youtube.com/watch?v=mMUARWRLx_U

#### A.5. Homelab consensus picks & trends, 2026

- Switch shortlist (ComputingForGeeks roundup, checked live June 2026, updated ~20h before Sep 22): best overall **MikroTik CRS310-8G+2S+IN** (8× 2.5GbE + 2× 10G SFP+, $210–230); best value/easy **TP-Link Omada ES210X-M2** (8× 2.5G + 2× 10G SFP+, fanless, ~$100); PoE pick **TP-Link SG3210XHP-M2** (8× 2.5G PoE+, 240W budget, ~$385); fanless managed **TP-Link SG3210X-M2** (~$230); "best if you already run UniFi" **UniFi Flex 2.5G PoE** (~$260); pure-fiber aggregation **MikroTik CRS305-1G-4S+IN** (4× 10G SFP+, fanless, $130–150). **[independent]** https://computingforgeeks.com/best-25gbe-10gbe-managed-switch-homelab/
- Trend (2026 GitHub homelab docs): **UCG-Fiber gateway + MikroTik switching** is a recurring combo (e.g., UCG-Fiber + MikroTik CRS310, VLAN-segmented, Talos Kubernetes cluster) **[secondary]** https://github.com/reptambe/digital-garden/blob/HEAD/content/Portfolio/Homelab/4%20-%202026%20Homelab/1%20-%20Architecture.md; Minisforum MS-01 / N5 Pro Proxmox nodes common; Tailscale tailnets for remote access; `home.arpa` local DNS naming; MikroTik CRS812/CRS310 and UniFi UDM-Pro Max appear in larger builds **[secondary]** https://github.com/lushanoperera/homelab/blob/HEAD/hardware-purchases/CLAUDE.md · https://github.com/tyrion70/multica-agents/blob/HEAD/skills/homelab/SKILL.md
- ServeTheHome / LTT-specific 2026 consensus: no dedicated STH or LTT roundup surfaced in this search pass; treat the ComputingForGeeks list + r/homelab-adjacent GitHub builds as the proxy evidence. **[unverified for STH/LTT specifically]**

#### A.6. Supply / availability / surcharges, 2026

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

