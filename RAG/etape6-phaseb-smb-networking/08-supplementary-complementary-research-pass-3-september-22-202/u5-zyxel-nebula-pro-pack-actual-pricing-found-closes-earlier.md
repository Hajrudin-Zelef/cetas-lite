---
id: etape6-phaseb-smb-networking/08-supplementary-complementary-research-pass-3-september-22-202/u5-zyxel-nebula-pro-pack-actual-pricing-found-closes-earlier
title: "U5. Zyxel Nebula Pro Pack — actual pricing found (closes earlier gap)"
domain: supplementary-complementary-research-pass-3-september-22-202
role: deep-dive
task: pricing
actors: ["EU", "United States"]
dates: ["2026-09-22"]
keywords: ["pricing", "ethernet", "license", "licenses", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1457, 1508]
section: "Supplementary / Complementary Research Pass #3 — September 22, 2026"
sha256: 610f1f846b093fac5184fcc78e7a34c5fce6400d2054fd8c09c03884dbc76a24
---

# U5. Zyxel Nebula Pro Pack — actual pricing found (closes earlier gap)

## U5. Zyxel Nebula Pro Pack — actual pricing found (closes earlier gap)

Per-device Pro Pack licenses, SKU pattern `LIC-NPRO-ZZ*n*Y00F` [vendor-reported via resellers]:

| Term | Price observed (Sep 22, 2026) | Source region |
|---|---|---|
| 1-year | €30.53 ex VAT / €36.94 incl. VAT (out of stock at listing); CAD 46.99 | vb.net (EU); ca.insight.com (Canada) |
| 2-year | £49.92 ex VAT / £59.90 incl. VAT; $64.83 (CDWG); €57.98 ex VAT | senetic.co.uk; CDWG (US); vb.net (EU) |
| 4-year | €98.90 ex VAT / €119.67 incl. VAT; ₪373.44 ex VAT | vb.net (EU); senetic.co.il (Israel) |
| 7-year | ₪599.93 ex VAT | senetic.co.il (Israel) |

- Note: Nebula Pro Pack is the license tier above the free Nebula base tier; Pro adds the management features surveyed earlier in this file.
- Sources observed September 22, 2026:
  - https://shop.vb.net/en/ZYXLICNPROZZ1Y00F/software-licenses-upgrades/zyxel-nebula-professional-pack-license-per-device-1-year
  - https://www.senetic.co.uk/product/LIC-NPRO-ZZ2Y00F
  - https://shop.vb.net/en/ZYXLICNPROZZ4Y00F/software-licenses-upgrades/zyxel-nebula-professional-pack-license-per-device-4-year
  - https://ca.insight.com/en_CA/shop/product/LICNCCPRO1YR/zyxel/LICNCCPRO1YR/Zyxel-Nebula-Professional-Pack-subscription-license-1-year-1-device/
  - https://www.cdwg.com/product/zyxel-nebula-professional-pack-subscription-license-2-years-1-device/7492128

## U6. Grandstream GWN7816 / GWN7816P (new — above the GWN7800 line)

Enterprise L3 managed aggregation switches (newer than the GWN7800 covered in the base report):

- **GWN7816**: 48× GbE + 6× 10G SFP+, 216 Gbps switching capacity, 108 Gbps non-blocking, 160.704 Mpps, 16 MB buffer [vendor-reported via reseller].
- **GWN7816P**: 740W total PoE budget; 60W max on ports 1–8, 30W on ports 9–48; optional second hot-swap PSU [vendor-reported via reseller].
- Dynamic routing: RIP/RIPng, OSPF/OSPFv3; **BGP not confirmed** — flag for EVPN/BGP research (Step 6 Phase D) [unverified].
- **Stacking conflict:** one reseller lists up to 4-unit stacking; another says stacking is "pending" — unresolved [unverified].
- Prices observed September 22, 2026: **€929 ex VAT / €1,114.80 incl. VAT** (GWN7816P); **AUD 2,313.30 RRP** (GWN7816P); **ZAR 9,213.80 incl. VAT** (non-PoE GWN7816) [vendor-reported via resellers].
- Sources:
  - https://www.snappernet.co.nz/product/57429/grandstream-gwn7816p-54-port-layer3-managed-poe-ethernet-switch-48-x-gige-6-x-sfp
  - https://rts-bg.com/en/network-switches/1790-grandstream-gwn7816p-48-gigabit-poe-ports-l3-managed-switch.html
  - https://miro.co.za/07-networking-switches---managed-layer-3/7677-grandstream-gwn7816-enterprise-l3-48-port-managed-gbe-6x-sfp-stackable-switch-6947273704850.html

## U7. Aruba Instant On 1930 — street-price snapshots (JL683B)

- Specs: 24× GbE PoE+ + 4× 10G SFP+, 195W PoE budget, 128 Gbps switching [official via reseller specs].
- Prices observed September 22, 2026: **£267.26 ex VAT / £320.71 incl. VAT** (UK); **US $422.74–423.84** (Microless); **$514.99** (B&H); $494.99 before discount (link-us-online) [vendor-reported via resellers].
- No recurring management subscription for Instant On cloud [official].
- Sources:
  - https://online.qual.co.uk/hpe-aruba-networking-aruba-instant-on-1930-24g-class4-poe-4sfp-sfp-195w-managed-l2-gigabit-ethernet-10-100-1000-power-over-ethernet-poe-1u-white-hpe-jl683b-acc
  - https://global.microless.com/product/aruba-instant-on-1930-24-port-poe-compliant-managed-network-switch-with-sfp-128-gb-s-switching-capacity-195w-power-budget-95-23-mpps-forwarding-rate-800-mhz-arm-cortex-a9-cpu-white-jl683b/offers/
  - https://www.bhphotovideo.com/c/replacement_for/1561168-REG/aruba_jl683a_instant_on_1930_24_port.html
  - https://link-us-online.com/product/aruba-instant-on-1930-24g-switch-jl683b/

## U8. NETGEAR WBE750 Wi-Fi 7 (ecosystem data point)

- Tri-band Wi-Fi 7 AP, up to 18.4 Gbps aggregate, 10G/multi-gig PoE++ uplink, up to 600 clients; one-year Insight subscription included [vendor-reported via reseller].
- **Label:** current product, not established as a 2026 launch [unverified on launch timing].
- Sources observed September 22, 2026:
  - https://www.techpowerup.com/320362/netgear-unveils-the-ultimate-tri-band-wifi-7-access-point-with-the-wbe750
  - https://www.cdw.com/product/netgear-wbe750-tri-band-wi-fi-7-ieee-802-11-a-b-g-n-ac-ax-be-18.40-gbit-s-w/7865277

