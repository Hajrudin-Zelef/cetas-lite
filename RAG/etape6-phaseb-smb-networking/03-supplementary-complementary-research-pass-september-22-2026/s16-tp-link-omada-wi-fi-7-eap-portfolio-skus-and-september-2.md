---
id: etape6-phaseb-smb-networking/03-supplementary-complementary-research-pass-september-22-2026/s16-tp-link-omada-wi-fi-7-eap-portfolio-skus-and-september-2
title: "S16. TP-Link Omada Wi-Fi 7 EAP portfolio — SKUs and September 2026 street prices (was: EAPs absent)"
domain: supplementary-complementary-research-pass-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "EU", "Qualcomm", "United States"]
dates: ["2026-03-20", "2026-09"]
keywords: ["cost", "dram", "ethernet", "foldable", "omni", "pricing", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [534, 582]
section: "Supplementary / Complementary Research Pass — September 22, 2026"
sha256: 54cd7be7a1e2350f95fd41b7fd92aa092c86c1aa3caa06d0f75121a7e9cf2680
---

# S16. TP-Link Omada Wi-Fi 7 EAP portfolio — SKUs and September 2026 street prices (was: EAPs absent)

## S16. TP-Link Omada Wi-Fi 7 EAP portfolio — SKUs and September 2026 street prices (was: EAPs absent)

Full Omada-managed Wi-Fi 7 AP stack, current 2026, from UK (linitx.com, Sep 2026), Canada (PC-Canada, Sep 2026), US (directdial.com), and Australia (ple.com.au) retailer snapshots [secondary]:

| SKU | Rating / form | Uplink | Sep 2026 snapshot |
|---|---|---|---|
| EAP723 | BE3600 dual-band, 1×2.5G, 160 MHz, MLO, 802.3at PoE | 2.5GbE | CAD $149.99 / £89.99 inc VAT |
| EAP725-Wall | BE3600 wall-plate: 1×2.5G PoE input + 1×2.5G pass-through + 2× GbE downlink, BT 5.2 | 2.5GbE | CAD $189.99 / £129.98 inc VAT |
| EAP725-Outdoor | BE3600 indoor/outdoor, directional + detachable omni antennas, IP66, 802.3at | 2.5GbE | £174.98 inc VAT |
| EAP770 | BE11000-class (10.4 Gbps) tri-band, 10G Ethernet, BT 5.2, 21.8 W | 10G | CAD $279.99 |
| EAP772 | BE9300 tri-band (5760 Mbps 6 GHz + 2880 Mbps 5 GHz + 574 Mbps 2.4 GHz), 1×2.5G, BT 5.2 | 2.5GbE | £154.99 inc VAT |
| EAP772-Outdoor | BE9300 tri-band, IP68, pole/wall mount, BT 5.2 | 2.5GbE | $369.99 CAD |
| EAP773 | BE9300 tri-band, 1×10G PoE+ port, 380+ concurrent clients, 1500 sq ft coverage, BT 5.2 | 10G | $190 USD / £191.99 inc VAT / AU $429 |
| EAP787 | BE12000 tri-band ceiling-mount | — | listed, price not captured — **flag** |

- **[independent]** TechRadar review of EAP772-Outdoor: US model BE11000 ($245.90 on Amazon.com), UK/EU model BE9300 (£261.50 on Amazon.co.uk) due to bandwidth restrictions; requires ~28 W PoE. Same review compares directly to UniFi U7 Pro Outdoor ($279 US / £195 UK) — cross-vendor outdoor Wi-Fi 7 pricing data point
- **Flag:** launch dates for individual EAP SKUs not verified — present as the current 2026 lineup, not as 2026 launches
- Sources: https://linitx.com/category/omada-wifi-7/1392 ; https://www.pc-canada.com/item/tp-link-omada-eap773-tri-band-wi-fi-7-ieee-802-11-a-b-g-n-ac-ax-be-10-40-gbit-s-wireless-access-point/eap773 ; https://www.directdial.com/us/item/tp-link-omada-eap773-tri-band-ieee-802-11-a-b-g-n-ac-ax-be-10-40-gbit-s-wireless-access-point/eap773 ; https://www.techradar.com/pro/phone-communications/i-reviewed-the-tp-link-omada-eap772-outdoor-and-it-seamlessly-provided-comprehensive-wi-fi-7-coverage-inside-and-out

## S17. Flagship / adjacent Wi-Fi 7 APs — Zyxel WBE660S and Grandstream GWN7670LR (was: S5 covered budget NWA + WBE510D/630S; Grandstream had no Wi-Fi 7)

- **Zyxel WBE660S** (BE22000 flagship, above the S5 lineup) [vendor-reported/official via Zyxel and reviews]: triple-radio (4×4 2.4 GHz + 4×4 5 GHz + 4×4 6 GHz), 12 spatial streams, up to 22 Gbps, Qualcomm Networking Pro 1220 quad-core (2.2 GHz), 2 GB DRAM, 1× 10GbE PoE++ uplink + 1× GbE, smart antenna with 4G/5G interference filter, NebulaFlex Pro (cloud / on-prem controller / standalone), USB-C power option, ~41 W draw, 1000+ clients / 8 SSIDs [secondary: TechRadar review, Newegg listing]
- WBE660S street price: **$540.61** (Newegg, Sep 2026) [secondary]; launch pricing context $599.99 [secondary]. Adjacent NWA130BE (BE11000, 1×2.5G uplink) launched $179.99 [secondary]
- **Flag:** WBE660S launch dates to 2023 (not a 2026 launch) — presented as the current flagship, not new hardware
- **Grandstream GWN7670LR** [official via GWN Series brochure]: outdoor long-range Wi-Fi 7 AP — 3.6 Gbps aggregate throughput, dual-band 2×2:2 MU-MIMO, 1× 2.5GbE RJ45 + 1× 2.5G SFP interface, PoE/PoE+, 256 clients, 350 m coverage (up to 1.5 km PtP bridge mode), IP66
- **Flag:** the source brochure is dated 2025 — verify whether GA occurred in 2026 before calling it a 2026 launch
- Sources: https://www.zyxel.com/global/en/products/wireless/be22000-12-stream-wifi-7-triple-radio-nebulaflex-pro-access-point-wbe660s/overview ; https://www.techradar.com/pro/zyxel-wbe660s-review ; https://dongknows.com/zyxel-nwa130be-be11000-wi-fi-7-access-point-review/ ; https://content.grandstream.com/hubfs/GWN_Series_Brochure_2025.pdf

## S18. New SKUs / products not previously captured: Instant On 1960 multi-gig SKU, UniFi Travel Router Long-Range, homelab 25G options

- **HPE Networking Instant On 1960 S0F35A** (new SKU vs S7's JL809A) [secondary via retailer specs]: 12-port multi-gigabit — 8× 1G PoE+ + 4× 1/2.5G PoE++ + 2× 10GBase-T + 2× 10G SFP+, 480 W PoE budget, 116 Gbps switching capacity, cloud/local management, stacking
- Price snapshots Sep 22, 2026 [secondary]: NZD **$2,876.66** (PB Tech, listed same day) ; UK £604.66 ex VAT, awaiting stock
- **[independent]** March 20, 2026 comms-express guide confirms "HPE Networking Instant On" branding, zero-cost cloud management, and the 1430/1830/1930/1960 segmentation (no confirmed new 2026 Instant On hardware)
- **UniFi Travel Router Long-Range** [secondary]: launched ~Sep 2026 (reports from ~16 days before Sep 22, 2026) at **$99** — foldable external antenna array; sold out within hours of launch, no restock timeline published by Ubiquiti; original UTR retails $79 and also sells out within minutes per restock (Mar 7, 2026 restock sold out same day); eBay resale $200–300
- UTR specs [secondary]: Wi-Fi 5 (866 Mbps), 2× GbE, USB-C (adapter not included), WireGuard + OpenVPN client, standalone mode (no UniFi account required), 89 g
- **Flag:** one Long-Range report (travelandtourworld.com) reads AI-generated — treat specs as [secondary], price $99 and sold-out status as consistent across two sources
- **Homelab / prosumer 25G switching options (2026 community observations)** [secondary/independent — not rankings]:
  - MikroTik **CRS518-16XS-2XQ-RM**: 16× 25G SFP28 + 2× 100G QSFP28, RouterOS v7, dual PSUs, compact rack-mount [secondary, 2026 roundup]
  - MikroTik **CRS510-8XS-2XQ-IN**: 8× 25G SFP28 + 2× 100G QSFP28, dual hot-swap PSUs — newer budget 100G entry [secondary]
  - FS.com **S5860-20SQ**: used as quiet 25G option (supports QSFP+ breakout and SFP28), no RoCE support [independent, STH forums]
  - Used **Brocade/Ruckus ICX6610 / ICX7250**: ~$300–400 used (24-port PoE examples), 8× 10G SFP+ (7250) / 40G QSFP (6610), L3 FastIron, fanless-ish ICX7150-C12P for living spaces [independent, STH forums + HN]
  - **Flag:** no verified 2026 "best homelab switch" ranking; these are observed community mentions only
- Sources: https://www.pbtech.com/pacific/product/SWHAUB19605/Aruba-Instant-On-1960-S0F35A-12-Port-Multi-gigabit ; https://networkwarehouse.co.uk/collections/aruba-instant-on-1960-series-switches ; https://www.comms-express.com/blog/hpe-aruba-meet-the-aruba-instant-on-switches/ ; https://www.nomadlawyer.org/ubiquiti-unifi-travel-router-long-range-features-sold-out-2026 ; https://resellcalendar.com/news/news/ubiquiti-unifi-travel-router-utr-reseller/ ; https://nerdtechy.com/best-25gbe-switches-with-sfp28-ports ; https://forums.servethehome.com/index.php?threads/quiet-25gbit-switch.32489/ ; https://news.ycombinator.com/item?id=33386230

---


---

