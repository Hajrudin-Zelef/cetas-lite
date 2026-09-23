---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/supplementary-pass-4-verification-notes-open-gaps
title: "Supplementary pass #4 — verification notes / open gaps"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["EU", "Qualcomm", "United States"]
dates: ["2026-05", "2026-05-15", "2026-09", "2026-09-22"]
keywords: ["dram", "ethernet", "license", "memory", "optics", "pricing", "research", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1034, 1126]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: c265e27944da46ff6b591f3442928062bdf7d1f57f424c653ea6d86d2d3a2c35
---

# Supplementary pass #4 — verification notes / open gaps

## Supplementary pass #4 — verification notes / open gaps

1. Omada PRO SDN Controller licensing terms (per-device vs free) — not published [unverified].
2. Zyxel XS1935-12HP / XS1935-10 street prices — not found [unverified].
3. NETGEAR M4350-16M4V / M4350-16C street prices — not captured yet [unverified].
4. Grandstream GWN7672 (tri-band flagship) street price — not captured [unverified].
5. D-Link DQS-5000-56ZS public price — quote-only [unverified].
6. 25G SFP28 transceiver/DAC street-price survey for the SMB/prosumer tier — deferred to Phase C (optics/fiber/cabling scope).


## Supplementary / Complementary Research Pass — September 22, 2026 (fourth wave; S19–S24)

**Scope:** adds only items absent from the base sections and the S1–S18 waves above. Zyxel XGS2220 street prices (was: context-only mention), NETGEAR M4250/M4500 + extra M4350 SKUs (was: M4350 only), MikroTik 100G/25G street pricing detail (was: community mentions), EnGenius ECW526/ECW520 + ECS5512F (was: ECW515/536S only), UniFi UDR7 GA confirmation + regional pricing (was: name only), Grandstream GWN7700M unmanaged (was: GWN7700MP only), 2026 supply/tariff notes. Omada Pro pricing, U7 Pro XGS/E7 confirmation, Aruba Instant On Wi-Fi 7, and Netgear 2026 financials remain open gaps and are **not** repeated. All claims carry source-class tags per the file convention. No existing section was altered.

Provenance tags: [official] = vendor price list/store/datasheet; [vendor-reported] = vendor claim via trade press; [independent] = third-party test/retailer listing; [secondary] = press/blog/forum; [unverified] = single weak source or inferred.

## S19. Zyxel XGS2220 series — 2026 street prices + GS1915 price data points (was: old-series context only)

XGS2220 (launched 2023) remains the current Zyxel L3 access line in 2026; no 2026 hardware refresh found [secondary — flag]. September 2026 street snapshots [independent]:
- **XGS2220-30HP** (24× GbE PoE++ 400W + 2× 10G multi-gig + 4× 10G SFP+, 168 Gbps, NebulaFlex Pro): €873.83 incl VAT (maltazon.com, Malta — **1 unit, very low stock**) [independent]
- **XGS2220-30HP** (same): regular $1,119.99 → **$799.99 sale** (lttpartners.com, US) [independent]
- **XGS2220-30** (non-PoE): ₹239,859 incl taxes (blumaple.com, India, low stock) [independent]
- **XGS2220-30F-DC** (24× SFP, DC power): regular $1,599.99 → **$849.99 sale** (lttpartners.com, US) [independent]
- **XGS2220-54HP** (48-port, 40× PoE+ + 10× PoE++, 600W): ₹264,969 incl taxes (blumaple.com) [independent]
- **GS1915-24EP** (24× GbE PoE, L2 smart): €300.36 (goalad.com, IT) — first GS1915 price data point observed [independent]
- **GS1915-8** (8-port): €82.59 (goalad.com, IT) [independent]
- **Zyxel CX4800-56F** (MultiGig 48-port, seen in 2026 EU listing): €3,934.27 (galaxus.de) — **flag: specs not verified** [independent]
- Sources: https://maltazon.com/switch-poe-gbe/13588-zyxel-xgs2220-30hp-managed-l3-gigabit-ethernet-10-100-1000-power-over-ethernet-poe-black.html ; https://www.lttpartners.com/products/zyxel-xgs2220-30hp-24-port-gigabit-l3-managed-and-nebula-pro-cloud-managed-switch-26-poe-ports-at-400w-with-4-sfp-10g-uplinks-and-2-10g-ethernet-ports ; https://goalad.com/annunci-italia/zyxel-xgs2220-30-eu0101f-switch-managed-layer-lite

## S20. NETGEAR AV/IT lines — M4250 full price table + M4500 context + extra M4350 SKUs (was: M4350 only)

M4250 = the 1G-focused AV-over-IP line (complements S6's M4350 IT mid-enterprise line); M4500 = 100G flagship above both [vendor-reported via resellers]. September 2026 price snapshots [independent]:
- **M4250-12M2XF** (MSM4214X-100NAS, 12× multi-gig + 2× 10G SFP+): $1,079.99 (videoguys.com, 5 in stock) [independent]
- **M4250-10G2XF-PoE++** (GSM4212UX-100NAS): MSRP **$1,939.29** / MAP $1,479.99 (broadfield.com, 0 in stock) [independent]
- **M4250-26G4F-PoE+** (GSM4230P-100NAS, 24× GbE + 4× SFP+): $2,045.34 (globalmagnetique.com) [independent]
- **M4250-26G4F-PoE++** (GSM4230UP-100NAS, Ultra90 PoE++): $3,633.01 (globalmagnetique.com, sold out) / €2,987.13 incl VAT (bekafun.com, BE — temporarily out of stock) [independent]
- **M4250-26G4F-PoE++** (same SKU): $3,599.69 sale from $3,929.64 (lttpartners.com, 40 in stock) [independent]
- **M4250-40G8XF-PoE+** (GSM4248PX-100NAS, 40× GbE + 8× SFP+): $4,934.56 (globalmagnetique.com) [independent]
- AVB licensing: AVB profile requires a license **sold separately** across M4250 [vendor-reported via reseller specs]
- **M4500** (100G AV/IT flagship): M4500-32XF8C referenced as $6,649.31-class SKU in reseller catalogs — **flag: not verified Sep 2026** [unverified]
- Extra **M4350 SKUs** seen in 2026 EU listings [independent]: M4350-16C (16× 100G?) €8,584.15 / M4350-16V4C €9,108.65 / M4350-24X8F8V (24× 10G + 8× 25G + 8× 100G, 40 ports) / M4350-40F4C (44-port) — specs thin, treat port counts as listing claims [unverified]
- **NETGEAR 24-port 10G/Multi-Gig Easy Smart switch** (26-port, MSB-style line): €334 (galaxus.de) — SMB easy-smart positioning below M4350 [independent]
- Sources: https://videoguys.com/products/netgear-m4250-12m2xf-av-switch ; https://www.broadfield.com/M4250-10G2XF-PoE ; https://globalmagnetique.com/en/products/netgear-m4250-40g8xf-poe-av-line-managed-switch-1 ; https://www.bekafun.com/en/a/32023619/netgear-network-solutions-ng42504-netgear-av-netwerkswitch-m4250-26g4f-poe

## S21. MikroTik 100G/25G street pricing — CRS520 / CRS504 / CRS304 (was: community mentions only)

Detailed September 2026 pricing for the prosumer/homelab 100G line [independent]:
- **CRS520-4XS-16XQ-RM** (16× 100G QSFP28 + 4× 25G SFP28 + 2× 10G RJ45, Marvell 98CX8410, quad-core 2 GHz ARM, 4 GB DDR4, dual hot-swap PSU, RouterOS v7 L5): MSRP **$2,195.00** → street **$1,955.20** (shop.linktechs.net, US, 8 in stock); AU $3,520 (wisp.net.au, last items); NZ $4,095 +GST (gowifi.co.nz, 4 in stock); Indonesia Rp 34,700,000 (mikrotik.id, listed 2026-05-15) [independent]
- **CRS504-4XQ-IN** (4× 100G QSFP28, compact, ~25 W): $698.50 (networkdevicesinc.com) / $710–$799 (balticnetworks.com, AU qty pricing $1,159–$1,304 at 10–20 units); ZA R17,082 incl VAT (geewiz.co.za, in stock at external supplier); CA out of stock at pc-canada.com (MSRP $1,985.32 CAD shown) [independent]
- **Flag — listing confusion:** geewiz.co.za's "CRS504-4XQ-IN" page describes 16× 25G ports, which matches the CRS510-8XS-2XQ-IN (see S18), not the 4× 100G CRS504 — treat that listing's spec block as unreliable [unverified]
- **CRS304-4XG-IN** (compact 4× 10G RJ45, no modules needed): $294.03 (wisp.net.au, in stock) — cheapest 10G entry in the line [independent]
- **US Section 301 tariff surcharge, Sep 2026:** balticnetworks.com adds **+$73.13 tariff surcharge at checkout** on the CRS504 — visible US tariff pass-through on imported network hardware [independent]
- Sources: https://shop.linktechs.net/mikrotik-crs520-4xs-16xq-rm-100-gbe-enterprise-cloud-router-switch ; https://www.balticnetworks.com/products/mikrotik-crs504-4x-qsfp28-100gbps-cloud-router-switch ; https://networkdevicesinc.com/products/crs504-4xq-in

## S22. EnGenius 2026 detail — ECW526 street price, ECW520, ECS5512F (was: ECW515/ECW536S only)

- **ECW526** (Cloud7 2x2x2 tri-band Wi-Fi 7, 10G PoE+ uplink, 9.4 Gbps aggregate, Qualcomm platform, 320 MHz, MLO, 4096-QAM): official store **$369.00** (store.engeniustech.com, Sep 2026); street **under $300** per Dong Knows review (~May 2026, updated ~125 days before Sep 22, 2026) [official/independent]
- Dong Knows assessment [independent]: "virtually the Wi-Fi 7 version" of EnGenius's Wi-Fi 6E AP, cloud-managed only (no local operation), performance "subdued" vs the 10G uplink — but recommended as a quick Wi-Fi 7 upgrade path
- **ECW520** (Cloud7 2x3x3 Lite tri-band Wi-Fi 7, 2.5 GbE uplink, 802.3at): listed on the 2026 EnGenius switch/AP category staging page — budget Wi-Fi 7 tier below ECW526 [vendor-reported]
- **ECW516L** (2x2:2/3x3:3 asymmetric Wi-Fi 7): also on the 2026 category page [vendor-reported]
- **ECS5512F** (Cloud Managed 12-Port 10-Gigabit Half-Rack Aggregate Fiber Switch): 12× 10G, half-rack form factor — new SKU in the 2026 switch lineup page, complements ECS2530FP/ECS2552FP aggregation options [vendor-reported]
- Current 2026 EnGenius switch segmentation per official staging page [vendor-reported]: Layer 3 models / 48-port L2 / 24-port L2 / half-rack aggregate (ECS5512F)
- Sources: https://store.engeniustech.com/products/ecw526-cloud-managed-2x2x2-indoor-tri-band-wifi-7-access-point ; https://dongknows.com/engenius-ecw526-wi-fi-7-access-point-review/ ; https://www.engeniustech.com/cloud-managed-network-switch.html ; https://static.engeniuscdn.com/wp-content/uploads/2026/02/12114759/DS_EnGenius-Cloud-APs_v6.0.pdf

## S23. UniFi UDR7 — GA confirmed, September 2026 regional prices/availability (was: name only)

UDR7 is shipping broadly as of Sep 22, 2026 [independent]:
- BG (rlan.bg): **€320.23 incl VAT**, 90+ in stock, delivery ~Sep 30, 2026
- EU distributor (aerial.net): **€268.49 ex VAT**, 1 pc in stock + 110 pcs ETA Sep 30
- UK (senetic.co.uk): **£247.68 ex VAT** / £297.22 incl VAT, 8 pcs, delivery ~Oct 1, 2026
- ZA (scoop.co.za): **R6,775.00 incl VAT** (60 units across branches) / (geewiz.co.za): R6,242 incl VAT, in stock at external supplier
- IL (senetic.co.il): ₪1,172.66 incl VAT, 2 pcs
- ZA (miro.co.za): **sold out, new stock ETA Nov 5, 2026** — R6,941.40 incl VAT retail
- Specs confirmed across retailers [vendor-reported via retailers]: quad-core, 3 GB RAM, 1× 10G SFP+ WAN + 1× 2.5G RJ45 WAN + 3× 2.5G LAN (1× 802.3af PoE-out), tri-band Wi-Fi 7 (10.7 Gbps aggregate: 5.7 + 4.3 + 0.688 Gbps), 64 GB microSD included, 0.96" LCD, manages 30+ UniFi devices / 300+ clients, ~2.3 Gbps IDS/IPS throughput, Protect/Access/Talk pre-installed
- **Flag:** US store.ui.com list price not captured in this pass — do not assume
- Sources: https://rlan.bg/en/29104-ubiquiti-udr7-unifi-dream-router-7 ; https://www.senetic.co.uk/product/UDR7 ; https://miro.co.za/01-wi-fi-ubiquiti-unifi-cloud-gateways/8256-ubiquiti-unifi-cloud-gateway-dream-router-7-udr7-810084699980.html ; https://scoop.co.za/ubiquiti-unifi-dream-router-7-tri-band-cloud-gateway-udr7.html

## S24. Grandstream GWN7700M unmanaged + 2026 supply/availability notes (was: GWN7700MP only; no supply section)

- **GWN7700M** (non-PoE sibling of the GWN7700MP in Part 2 §5): unmanaged 5× 2.5G RJ45 + 1× 10G SFP+, fanless metal desktop/wall; **€62.92 ex VAT** (€76.21 incl VAT, CZ discomp, in stock, ships within 48h, Sep 2026) [independent]
- **No new Grandstream switch model announced in 2026 found** — GWN7700M/MP, GWN7800 Pro, GWN7810/7816, GWN7821P/22P, GWN7830/31/32 remain the current line [unverified — absence of evidence]; GWN78xx firmware train documented (e.g. 1.0.13.18, Apr 2025) with GDMS cloud-management compatibility fixes [secondary]
- **Supply/availability snapshots, Sep 22, 2026** [independent]:
  - US Section 301 tariff surcharge now itemized at checkout by US retailers (e.g. +$73.13 on MikroTik CRS504 at balticnetworks.com) — tariff pass-through is becoming explicit in networking hardware pricing
  - Zyxel XGS2220-30HP: very low stock at EU reseller (1 unit, Malta)
  - Ubiquiti UDR7: mixed — broad EU/UK/IL stock, sold out in ZA (restock Nov 5); UTR Long-Range (S18) still selling out within hours
  - QNAP QSW-M7308R-4X (4× 100G + 8× 25G, half-width, MSRP $1,927): **backordered** at directdial.com (CA, $1,898 CAD) and pc-canada.com (MSRP $1,985.32); but **>10 pcs in stock** at galaxus.de (€1,615.02 incl VAT, −10%); SGD $2,215.55 (aceperipherals.com); ₹214,500 incl GST (tanotis.com, India) — availability is channel-dependent
  - Aruba Instant On 1960 S0F35A: NZD $2,876.66 (PB Tech, listed Sep 22, 2026); UK £604.66 ex VAT, awaiting stock
- **Flag:** no 2026 DRAM/memory-price-driven switch price hikes beyond Ubiquiti's surcharge (S9) were independently verified in this pass
- Sources: https://www.discomp.cz/grandstream-gwn7700m-unmanaged-network-switch-5x-2-5gb-portu-1-sfp-_d127130.html ; https://www.balticnetworks.com/products/mikrotik-crs504-4x-qsfp28-100gbps-cloud-router-switch ; https://www.directdial.com/ca/item/qnap-half-rackmount-switch-qsw-m7308r-4x-us-layer-2-12-port-managed-switch-4-p/qsw-m7308r-4x-us ; https://www.galaxus.de/en/s1/product/qnap-qsw-m7308r-4x-managed-switch-4-port-100gbe-8-port-25gbe-half-rackmount-design-12-ports-network--39444919


---

