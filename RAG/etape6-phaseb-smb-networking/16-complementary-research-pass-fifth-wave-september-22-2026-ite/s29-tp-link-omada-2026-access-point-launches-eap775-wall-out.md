---
id: etape6-phaseb-smb-networking/16-complementary-research-pass-fifth-wave-september-22-2026-ite/s29-tp-link-omada-2026-access-point-launches-eap775-wall-out
title: "S29. TP-Link Omada 2026 access-point launches — EAP775-Wall/Outdoor, EAP783, EAP787, EAP725"
domain: complementary-research-pass-fifth-wave-september-22-2026-ite
role: deep-dive
task: reference
actors: ["AWS", "United States"]
dates: ["2025-12", "2026-05"]
keywords: ["ethernet", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2509, 2563]
section: "Complementary Research Pass — Fifth Wave (September 22, 2026), Items S25–S33"
sha256: e7ef5d7c1f8877dc6505b3f36b879b98688836f6669cdfac2baa9bb9f920a376
---

# S29. TP-Link Omada 2026 access-point launches — EAP775-Wall/Outdoor, EAP783, EAP787, EAP725

## S29. TP-Link Omada 2026 access-point launches — EAP775-Wall/Outdoor, EAP783, EAP787, EAP725

### EAP775-Wall (BE9300 / BE11000 wall-plate Wi-Fi 7) `[official]`/`[secondary]`
- Tri-band Wi-Fi 7 wall-plate AP; 6-stream; 5765 Mbps (6 GHz) + 2882 Mbps (5 GHz) + 688 Mbps (2.4 GHz).
- Wired: 1× 2.5G PoE-in uplink + 1× 2.5G + 3× GbE downlinks; **PoE Out (PoE+/dual-PoE) on two Gigabit ports when powered by 802.3bt**.
- 380+ concurrent clients; 100 m² coverage; Bluetooth 5.2; 6× internal antennas; MLO, 4K-QAM, 320 MHz; WPA3; 24 SSIDs; Omada SDN managed.
- Max power 43.35 W (with 25 W PoE out); 172.9 × 91.2 × 41.3 mm.
- **Availability:** Amazon listing first available **17 December 2025** (ASIN B0G88FG4QH) — https://www.amazon.com/dp/B0G88FG4QH
- **Prices:** CAD $309.99 in stock (pc-canada.com, 2026) — https://www.pc-canada.com/item/tp-link-omada-eap773-tri-band-wi-fi-7-ieee-802-11-a-b-g-n-ac-ax-be-10-40-gbit-s-wireless-access-point/eap773 ; UK RRP £225.45+VAT, order on demand (provu.co.uk).
  - https://www.4cabling.com.au/tp-link-eap775-wall-omada-be9300-wall-plate-wi-fi-7-access-point.html

### EAP775-Outdoor (tri-band Wi-Fi 7, rugged) `[official]`
- 6-stream tri-band Wi-Fi 7: 5765 (6 GHz) / 4324 (5 GHz) / 688 Mbps (2.4 GHz).
- **AFC-enabled 6 GHz band** (outdoor legal operation); plug-and-play auto-sensing antennas switching between directional and omnidirectional modes (software per-band control).
- **IP68**, anti-UV coating, **6 kV lightning protection**; 1× 2.5G PoE+ port; multi-axis mounting; mesh + seamless roaming; Omada cloud managed.
- Datasheet dated May 2026 (TP-Link static CDN): https://static.tp-link.com/upload/product-overview/2026/202605/20260522/Datasheet_EAP775-Outdoor%20v1.pdf

### EAP787 (tri-band Wi-Fi 7 ceiling, 10G uplink) `[secondary]`
- Tri-band Wi-Fi 7, ~15 Gbit/s class; **1× 2.5G + 1× 10G Ethernet** uplinks; 24.5 W; ceiling/wall/junction-box mount.
- **CAD $449.99 in stock** (pc-canada.com, 2026); UK RRP £295.00+VAT, order on demand (provu.co.uk).

### EAP783 (Wi-Fi 7 flagship, 21 Gbit/s class) `[secondary]`
- Tri-band Wi-Fi 7 "21.03 Gbit/s" class; 2× network RJ-45 incl. **10 GbE**; 39 W.
- **CAD $699.99 in stock** (pc-canada.com, 2026).

### EAP772 / EAP725-Wall / EAP725-Outdoor (UK RRPs) `[secondary]`
- EAP772 (ceiling Wi-Fi 7): RRP £183.33+VAT, in stock (provu.co.uk).
- EAP725-Wall (Wi-Fi 7 wall plate): RRP £159.17+VAT, in stock.
- EAP725-Outdoor (Wi-Fi 7 outdoor): RRP £175.71+VAT, order on demand.
- Source: https://Www.Provu.Co.Uk/tp-link-omada-range-access-points (crawled Sept 2026)

### Gap status
- **RESOLVED:** The 2026 EAP775 line (wall + outdoor) and upper-tier EAP783/EAP787 pricing documented.
- **Retained:** US list prices (MSRP) for EAP783/EAP787/EAP775-Outdoor; EAP776 existence not confirmed (was a guess in earlier planning, no evidence found — do not invent).

---

## S30. Grandstream GWN7672 (BE11000 Wi-Fi 7) — street pricing confirmed (resolves retained gap)

Earlier passes (S4, S22, S24) carried "GWN7672 street price not captured" as a retained gap. Resolved:

- **UK:** **£145.00 ex VAT / £174.00 inc VAT, in stock** (best4systems.co.uk) — https://www.best4systems.co.uk/grandstream-gwn7672-tri-band-wi-fi-7-access-point.html
- **Canada:** CAD $228.38 regular (dealtargets.ca — sold out at crawl)
- **India:** ₹14,500–20,000/piece across Mumbai/Pune/Bengaluru sellers (indiamart.com, crawled ~Aug 2026) — https://www.indiamart.com/proddetail/grandstream-tri-band-wi-fi-7-access-point-gwn7672-2858838626412.html
- **KSA:** SAR 1,017 (noon.com Saudi, Sept 2026)
- IndiaMART category page (updated 25 Aug 2026) also lists **GWN7670 (dual-band Wi-Fi 7)** at ₹9,999 and **GWN7664E (Wi-Fi 6 4×4)** at ₹10,700 for context.
- Confirmed specs: tri-band Wi-Fi 7 (BE11000), 2×2:2 MU-MIMO (DL/UL OFDMA), MLO/4K-QAM/MRU/preamble-puncturing, up to 175 m range, **384 concurrent clients**, 2× 2.5GbE PoE+, **embedded controller manages up to 50 local GWN APs**; also managed by GDMS Networking (cloud) and GWN Manager (on-prem); BLE 5.3; anti-hacking secure boot.
- At ~£145 ex VAT it is one of the cheapest tri-band Wi-Fi 7 APs documented in this project (vs U7 Pro XGS $299, EAP787 ~$450 CAD class).

### Gap status
- **RESOLVED:** GWN7672 street pricing in UK/Canada/India/KSA.
- **Retained:** US MSRP/street price (not found in US retail channels this wave).

---

