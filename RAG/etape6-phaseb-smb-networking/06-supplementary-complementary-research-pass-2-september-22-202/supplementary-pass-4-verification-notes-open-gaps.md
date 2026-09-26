---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/supplementary-pass-4-verification-notes-open-gaps
title: "Supplementary pass #4 — verification notes / open gaps"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["EU", "United States"]
dates: ["2026-05-15", "2026-09", "2026-09-22"]
keywords: ["ethernet", "license", "optics", "pricing", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1034, 1089]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 822ef61caee671d6a2aafcd4ebc0284c5beab7f5dc0cddc08d98be28c41d420e
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

