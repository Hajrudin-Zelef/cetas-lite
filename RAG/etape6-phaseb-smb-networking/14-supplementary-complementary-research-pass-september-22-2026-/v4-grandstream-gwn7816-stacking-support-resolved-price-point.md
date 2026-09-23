---
id: etape6-phaseb-smb-networking/14-supplementary-complementary-research-pass-september-22-2026-/v4-grandstream-gwn7816-stacking-support-resolved-price-point
title: "V4. Grandstream GWN7816 — stacking support RESOLVED + price points (was: U13 gap #6)"
domain: supplementary-complementary-research-pass-september-22-2026-
role: deep-dive
task: reference
actors: ["Falcon", "Intel", "United States"]
dates: ["2026-01", "2026-09", "2026-09-22"]
keywords: ["intel", "license", "licenses", "optics", "pricing", "voice"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2251, 2321]
section: "Supplementary / Complementary Research Pass — September 22, 2026 (fifth wave; V1–V9)"
sha256: 6a21fd2dfe1d5ea7ac1b1d6b4010fd7b55f7b3b287fb105135089ab6e16a37c9
---

# V4. Grandstream GWN7816 — stacking support RESOLVED + price points (was: U13 gap #6)

## V4. Grandstream GWN7816 — stacking support RESOLVED + price points (was: U13 gap #6)

- **Resolution:** Stacking **is supported**. Grandstream's official feature set (via resellers quoting datasheet): "supports stacking for easy management on one interface while creating redundant backup between multiple devices." fgtechstore specifies **up to 4 units** stackable. One reseller (MiRO, South Africa) described it as "pending stacking support" — likely reflecting firmware vintage at listing time rather than hardware limitation; treat MiRO's "pending" note as [unverified]/superseded [secondary — resolved in favor of official datasheet].
- **Specs:** 48× GbE RJ45 + 6× 10G SFP+; 216 Gbps switching; 160.7 Mpps; 4K VLANs (QinQ, protocol/port/MAC-based, voice); static/RIP/RIPng/OSPF/OSPFv3 (BGP pending); 27 LAGs; 4K ACLs; QoS (SP/WRR/WFQ); ARP inspection, IP Source Guard, DHCP snooping, DoS, 802.1X, RADIUS, TACACS+; dual boot; fault detection; GWN.Cloud / GWN Manager / embedded controller / CLI / SNMP; 16 MB packet buffer; 32K MAC; 4 fans; 1U rack; 3-year warranty [vendor-reported via resellers].
- **GWN7816P** (PoE variant): dynamic PoE/PoE+/PoE++ allocation, up to **740 W** total, 60 W max per port [vendor-reported via reseller].
- **Prices observed September 22, 2026:**
  - **R9,213.80 inc VAT** (low stock; JHB 3 / CPT 1) at miro.co.za [vendor-reported via reseller].
  - **AU RRP $1,199.67 inc GST** at alloys.com.au (trade) [vendor-reported via reseller].
  - **₹74,000** at indiamart.com (India) [vendor-reported via reseller].
- Sources:
  - https://fgtechstore.com/product/grandstream-gwn7816/
  - https://miro.co.za/07-networking-switches---managed-layer-3/7677-grandstream-gwn7816-enterprise-l3-48-port-managed-gbe-6x-sfp-stackable-switch-6947273704850.html
  - https://www.alloys.com.au/security-and-automation/communications/ip-telephony-and-unified-comms/48-port-layer-3-managed-switch
  - https://www.ipandgo.com/en/routers-switches-ap-s/1945-grandstream-gwn7816-6947273704850.html

## V5. Aruba Instant On 1930R successor — re-check (was: U13 gap #7)

- **Finding:** No successor model found as of September 22, 2026. The Instant On 1930 line remains current in its A and **B revisions** — Aruba release notes (2.8.1.35, 2.9.1.17) document JL683B / JL684B / JL686B introduced with firmware 2.5.0.42 as acoustic-improved replacements for JL683A/JL684A/JL686A, same features [official].
- 1930 portfolio unchanged: JL680A (8G 2SFP), JL681A (8G PoE 124 W), JL682A (24G 4SFP/SFP+), JL683A/B (24G PoE 195 W), JL684A/B (24G PoE 370 W), JL685A (48G 4SFP/SFP+), JL686A/B (48G PoE 370 W); 128 Gbps switching, 95.23 Mpps on 24-port models; limited lifetime warranty [vendor-reported via resellers / official release notes].
- **Gap retained:** no 1930R / next-generation Instant On switch successor confirmed for 2026 [unverified].
- Sources:
  - https://www.arubanetworks.com/techdocs/InstantOn_1930_Switch/rn-1930-2.8.1.35.pdf
  - https://www.networktigers.com/products/jl682a-hpe-switch-new
  - https://buyrouterswitch.com/jl683a-price.html

## V6. D-Link DQS-5000 — second price attempt + spec correction (was: U13 gap #4)

- **Price attempt:** still no street price. Zoro lists DQS-5000-32Q28/AF (32× 100G QSFP28) but with **no price shown** and sale restricted in CA; AV-iQ lists DQS-5000-32S as "Request Quote" only [vendor-reported via resellers]. **Pricing gap retained** — D-Link 5000 Series appears to be quote-only in retail channels [unverified].
- **Spec correction:** the DQS-5000-56ZS official datasheet specifies **48× 10G/25G SFP28 + 8× 40G/100G QSFP28**; the AU variant DQS-5000-54SQ28 is **48× 25G SFP28 + 6× 100G QSFP28** — not 8× 100G. Earlier pass S8's "48×25G + 8×100G" conflated the two variants; corrected here [official — datasheet]. 2 Tbps switching, 2,380 Mpps, Marvell Falcon silicon, Intel x86 4-core CPU, 8 GB DDR4 ECC, 16 GB eMMC, 512K IPv4 / 256K IPv6 prefixes, ONIE + SONiC-ready, optional D-Link OS activation licenses (DQS-5K-*-DC-LIC) [official].
- Full 5000 Series lineup: DXS-5000-54S (48× 10G SFP+ + 6× 40G QSFP+), DQS-5000-32S (32× 40G QSFP+), DQS-5000-32Q28 (32× 100G QSFP28), DQS-5000-54SQ28 (48× 25G + 6× 100G); front-to-back airflow options [official via ordering guide].
- Sources:
  - https://usweb.dlink.com/en/-/media/product-pages/dqs/dqs-5000-56zs/dqs500056zsa1datasheetv101ww.pdf
  - https://www.dlink.com.au/business-solutions/DQS-5000-54SQ28-54-Port-Data-Centre-Switch-with-48-25GbE-SFP28-and-6-100GbE-QSFP28-Ports
  - https://www.zoro.com/d-link-systems-32-port-100g-qsfp28-open-network-switch-front-to-back-airflow-dqs-5000-32q28af/i/G9650079/

## V7. MikroTik — 2026 launches: CRS804 DDQ 400GbE + CRS304-4XG-IN (new detail)

- **CRS804 DDQ** — MikroTik's second 400GbE switch (after CRS812 DDQ), announced ~January 2026 per ServeTheHome: **4× QSFP56-DD 400G ports** + 2× 10G from management CPU; half-width chassis (two fit side-by-side in 1U); **Marvell 98DX7335** switch chip (same as CRS812); Annapurna Labs **AL52400 quad-core Arm** CPU; 1.6 Tbps total fabric; max 123 W with optics (92 W bare); dual hot-swap fans; redundant PSUs [secondary — STH, Jan 2026]. Significance: sub-$4K-class 400G from MikroTik brings 400GbE into the prosumer/homelab conversation [independent analysis].
- **CRS304-4XG-IN** — 4× **10GBase-T RJ-45** ready out of the box (no SFP modules needed) + 1× GbE management port; fanless heatsink-chassis design; RouterOS v7 L3 hardware offload; **NZD $368.70** at PB Tech (listing created Aug 18, 2026) [vendor-reported via reseller].
- Price snapshot: **CRS310-8G+2S+IN** (8× 2.5G + 2× SFP+) at **R4,090** at comx-computers.co.za, valid 2026-09-22 [vendor-reported via reseller].
- Sources:
  - https://www.servethehome.com/mikrotik-crs804-ddq-announced-4-port-400gbe-switch/
  - https://www.servethehome.com/mikrotik-crs812-ddq-400gbe-switch-launched-crs812-8ds-2dq-2ddq-marvell/
  - https://www.pbtech.com/pacific/product/NETMKT1502/MikroTik-Cloud-Router-Switch-CRS304-4XG-IN-with-4
  - https://www.comx-computers.co.za/MT-RBCRS310-8G-2S-IN-MikroTik-Cloud-Router-Switch-8x-Buy-p-290412.php

## V8. UniFi U7 Pro Max — September 2026 price snapshots (flag: 2024 launch, not 2026)

- **Prices observed September 22, 2026:**
  - **AUD $477.53** at PB Tech (AU), in stock, ships today [vendor-reported via reseller].
  - **£255.99 inc VAT** (£213.33 ex VAT), 128 in stock, bulk tiers down to £249.61 at 10+, at netxl.com (UK) [vendor-reported via reseller].
- Specs (current): BE15000 8-stream Wi-Fi 7, tri-radio 2.4/5/6 GHz, 4×4 5 GHz + 2×2 2.4/6 GHz, 2.5GbE uplink, PoE+, dedicated spectral engine, 160 m² / 500+ clients, real-time spectral analysis via UniFi Network 8.2.93+ [vendor-reported via resellers].
- **Flag:** U7 Pro Max launched ~mid-2024 alongside U7 Pro Wall and U7 Outdoor — included here as 2026 pricing context only, not as a 2026 launch [secondary — asbis/news archives, ~Jun 2024].
- Sources:
  - https://www.pbtech.com/product/NAPUBI1702/Ubiquiti-UniFi-U7-Pro-Max-BE15000-Tri-Band-Wi-Fi-7
  - https://www.netxl.com:443/wifi-access-points/ubiquiti-unifi-u7-pro-max-wifi-7-access-point/
  - https://news.asbis.com/news/suppliers/ubiquiti-revolutionizes-connectivity-with-new-unifi-wi-fi-7-access-points-u7-pro-max-u7-pro-wall-and-u7-outdoor/

## V9. Wave-5 verification log / remaining gaps

**Gaps closed in this wave:** U13 #2 (XS1935 pricing), U13 #3 (NWA210BE specs/pricing), U13 #5 (EAP783 naming conflict), U13 #6 (GWN7816 stacking).
**Gaps retained:**
1. D-Link DQS-5000 street pricing — second attempt failed; quote-only channel [U13 #4 retained].
2. Aruba Instant On 1930R successor — no evidence; B-revision refreshes are the 2026 state [U13 #7 retained].
3. TP-Link US regulatory outcome — developing; beyond this pass's scope [U13 #1 retained].
4. Currency conversions not applied — all prices are raw regional snapshots [U13 #8 retained].
5. Omada Pro controller appliance SKU + license pricing — still not found publicly in this wave (noted across passes).
**New uncertainty introduced:** some NZ reseller listings misquote EAP783 5 GHz radio as 8,640 Mbps — resolved in favor of TP-Link official spec; anyone re-scraping reseller catalogs will hit the conflict again [secondary].

---

