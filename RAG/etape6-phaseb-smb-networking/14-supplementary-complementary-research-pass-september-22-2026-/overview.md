---
id: etape6-phaseb-smb-networking/14-supplementary-complementary-research-pass-september-22-2026-/overview
title: "Supplementary / Complementary Research Pass — September 22, 2026 (fifth wave; V1–V9)"
domain: supplementary-complementary-research-pass-september-22-2026-
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2026-09-22"]
keywords: ["research", "ethernet", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2204, 2250]
section: "Supplementary / Complementary Research Pass — September 22, 2026 (fifth wave; V1–V9)"
sha256: db4e1edbe17aea26a20f969ff435df534036004e824979d7e51641389e7a13c2
---

# Supplementary / Complementary Research Pass — September 22, 2026 (fifth wave; V1–V9)

This pass targets the remaining open gaps logged in U13 plus a few genuinely new 2026 data points (MikroTik 400G, CRS304-4XG-IN, UniFi U7 Pro Max price snapshots). No existing section was modified. All facts carry provenance tags; regional prices are raw snapshots with no currency normalization.

## V1. Zyxel NWA210BE — specs and pricing RESOLVED (was: U13 gap #3 — model name only)

- **Specs (official via resellers/Icecat):** BE12300 dual-radio Wi-Fi 7; 2×2:2 at 2.4 GHz + 4×4:4 at 5 OR 6 GHz (BandFlex design — configurable 5/6 GHz); up to 12.3 Gbps aggregate; 2× 2.5GbE ports; 802.3at PoE+ (21.5 W max); internal antennas; USB-C; ceiling/wall mount; 250×160×47 mm, 815 g; max 8 SSIDs; standalone or Zyxel Nebula cloud management [vendor-reported via resellers].
- **Review verdict [independent]:** IT Pro tested real-world close-range copy ~190 MB/s and distance ~143 MB/s; Auto-MLO noted as a smart feature, BandFlex easy to use; limited-lifetime warranty; UK: £177 ex VAT at Broadbandbuyer. NWA130BE (full tri-band, half 6 GHz speed) mentioned at £133 as cheaper alternative [independent — review ~late 2025].
- **Prices observed September 22, 2026:**
  - **US$169.00** at microcom.us [vendor-reported via reseller].
  - **$169.99** (list $229.99) at Newegg — out of stock at observation [vendor-reported via reseller].
  - **€339.20 inc VAT**, 14 in stock, at maltazon.com (Malta) [vendor-reported via reseller].
  - **6,149 CZK** at alza.cz (Czech) [vendor-reported via reseller].
  - **€283.70** (€228.79 ex VAT) at galador.eu [vendor-reported via reseller].
  - **R6,606** at itworkup.co.za (South Africa) [vendor-reported via reseller].
- Sources:
  - https://maltazon.com/access-point-pro/zyxel-nwa210be-11530-mbit-s-white-power-over-ethernet-poe
  - https://www.alza.cz/EN/zyxel-nwa210be-d12880389.htm
  - https://www.itpro.com/hardware/routers/zyxel-nwa210be-review-a-dual-personality-wi-fi-7-access-point-for-budget-conscious-businesses
  - https://www.newegg.com/p/0ED-005N-00042?Item=0ED-001T-000S6&
  - https://www.microcom.us/nwa210be.html
- **Label:** not established as a 2026 launch — treated as current SMB-lineup pricing/spec data, launch timing [unverified].

## V2. Zyxel XS1935-12F — street pricing found (was: U13 gap #2 — regional pricing/availability unannounced)

- **Specs:** XS1935-12F (XS1935-12F-ZZ0101F, EAN 4718937637805): 12-port managed multi-gigabit/10G L2 switch, 2× 10GBase-T RJ-45 + remaining SFP+ (fiber variant); 240 Gbps switching; 178 Mpps; 32K MAC; 9K jumbo; rack-mount; AC adapter powered; **not stackable**; no PoE [vendor-reported via resellers].
- **Prices observed September 22, 2026:**
  - **R9,714** at itworkup.co.za (South Africa) [vendor-reported via reseller].
  - **€448.99** (later snapshot €432.99) at shopro.lt (Lithuania) [vendor-reported via reseller].
- Context: XS1935 sibling XGS1935-28HP (24× GbE PoE+ + 4× 10G SFP+, 375 W PoE) at £410.39 VAT-inc at Newegg UK; XGS1935 series street prices from £161.91 ex VAT at comms-express.com (UK, in stock next-day); XGS1935-52HP at £528 ex VAT per IT Pro review [vendor-reported via resellers / independent review].
- Sources:
  - https://itworkup.co.za/index.php?route=product/pdf&product_id=610220
  - https://www.shopro.lt/index.php?route=product/pdf&product_id=540230
  - https://www.comms-express.com/categories/zyxel-xgs1935-series-smart-managed-switches/
  - https://www.itpro.com/hardware/routers/zyxel-xgs1935-52hp-review-a-port-dense-gigabit-poe-switch-thats-priced-right-for-smbs

## V3. TP-Link EAP783 — BE19000 vs BE22000 naming conflict RESOLVED (was: U13 gap #5)

- **Resolution:** The official product is **BE19000**, not BE22000. TP-Link's official page (tp-link.com/au) specifies: BE19000 tri-band Wi-Fi 7 = 11,520 Mbps (6 GHz) + 5,760 Mbps (5 GHz) + 1,376 Mbps (2.4 GHz); 2× 10GbE ports; 320 MHz channels on 6 GHz; MLO, Multi-RUs, 4K-QAM, 4×4 MU-MIMO; 802.3bt PoE++ or 12 V DC; 39 W max; 280×280×46.5 mm, 1.385 kg; 12 internal antennas; 200 m² coverage claim; 640–760+ concurrent clients (marketing variance); Omada SDN (OC200/OC300/software/cloud controller) [official].
- **Where the confusion came from:** at least two New Zealand reseller listings (themall.aucklandairport.co.nz, techtonic.nz) quote the 5 GHz radio as **8,640 Mbps**, which would sum to ~21.5K and is how a "BE22000" label appears to have leaked into third-party catalogs. That 8,640 Mbps figure contradicts TP-Link's own spec (5,760 Mbps on 5 GHz) and should be treated as a **reseller spec-sheet error** [secondary — conflict resolved in favor of official].
- TechRadar hands-on (PC Pro issue 360, ~Jan 2026): "TP-Link's first Wi-Fi 7 business AP"; two 10GbE ports (first on PoE++, second for aggregation); 8 SSIDs per radio; standalone mode with quick-start wizard; slightly heavier than Zyxel WBE660S [independent].
- Sources:
  - https://www.tp-link.com/au/business-networking/omada-wifi-wifi7/eap783/
  - https://www.techradar.com/pro/tp-link-omada-eap783-review
  - https://themall.aucklandairport.co.nz/en/intl-duty-free/product/pb-tech-m_NAPTPL7831/tp-link-omada-eap783-%28be19000%29-tri-band-12-stream-wi-fi-7-access-point-with-10gbe-x2-%28poe%2B%2B-39w%29.html
  - https://www.singular.com.cy/tp-link-omada-eap783-v1-wi-fi-7-2456ghz-cloud-bt-2-ports-wallceiling-mount.html

