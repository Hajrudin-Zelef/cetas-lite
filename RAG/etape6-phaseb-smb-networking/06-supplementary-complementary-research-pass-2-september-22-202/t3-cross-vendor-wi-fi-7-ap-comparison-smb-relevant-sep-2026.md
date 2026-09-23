---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/t3-cross-vendor-wi-fi-7-ap-comparison-smb-relevant-sep-2026
title: "T3. Cross-vendor Wi-Fi 7 AP comparison (SMB-relevant, Sep 2026)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["ethernet", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [889, 939]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 999fe169baefd28bd058d405026fa92d469d9da5eeb35691d500f08d66ce35b7
---

# T3. Cross-vendor Wi-Fi 7 AP comparison (SMB-relevant, Sep 2026)

## T3. Cross-vendor Wi-Fi 7 AP comparison (SMB-relevant, Sep 2026)

Wi-Fi 7 APs are shipping across every SMB vendor now; summary table for orientation [provenance per row]:

| Vendor | Model | Band config / uplink | Street price (Sep 2026) | Source type |
|---|---|---|---|---|
| Ubiquiti UniFi | U7 Lite → U7 Pro Max | up to tri-band + 10GbE (Pro XG) | $99–$279 | [secondary] |
| TP-Link Omada | EAP770 (BE11000) | tri-band, 1× 2.5G RJ45 | CAD $309.99 MSRP | [official] |
| Zyxel | NWA50BE / NWA90BE / NWA50BE PRO / NWA90BE PRO | dual-radio MLO / tri-band | $79.99–$99.99 | [secondary] |
| Zyxel | WBE510D / WBE630S | high-density dual-radio MLO | $199.99 / $299.99 | [secondary] |
| EnGenius | ECW515 (wall-plate), ECW510 | dual-band, 2.5GbE PoE-in | per vendor launch | [vendor-reported] |
| EnGenius | ECW536S | AirGuard WIDS/WIPS | per vendor launch | [vendor-reported] |
| D-Link | DAP-E9560 / DAP-E3620 / DAP-E3620OU | tri-band / 2×2 / outdoor | Japan guide only, prices unconfirmed | [secondary] |
| NETGEAR | WBE750 (Insight-managed, business) | tri-band, 18.4 Gbps, 10G PoE++, 600 clients | per launch PR | [vendor-reported] |
| Grandstream | GWN7670 / GWN7672 / GWN7672L / GWN7670WM / GWN7670LR | dual-band BE3600 / tri-band + 10G | £125.10 (7672L); R2,685 (7670) | [independent] |
| Cisco Meraki | CW9172I | enterprise Wi-Fi 7 | pricing not captured | [secondary] |
| Aruba (enterprise) | 730 Series (ultra tri-band) | Apr 2024 launch, Aruba Central-managed | enterprise, not Instant On | [secondary] |
| Ruckus | R770 | enterprise, AFC-ready, RUCKUS AI | 2023 launch, current | [secondary] |
| Aruba Instant On | — | **no Wi-Fi 7 on the Instant On line as of Sep 2026**; AP32 is Wi-Fi 6E ($199) | current lineup AP22/AP25/AP27/AP32 | [secondary — flag] |

- Enterprise Wi-Fi 7 context [secondary]: HPE Aruba's ultra-tri-band launch (2024, still current 2026) claims +30% capacity vs rivals, IoT radios (Bluetooth/Zigbee), managed by Aruba Central; Meraki's CW9172I is the cloud-managed enterprise Wi-Fi 7 AP; Ruckus R770 launched 2023 with AFC readiness and RUCKUS AI assurance.
- Sources: https://www.technologydecisions.com.au/content/networking/news/hpe-launches-wi-fi-7-access-points-531321353 ; https://www.thefastmode.com/technology-solutions/35386-netgear-launches-powerful-wifi-7-access-point-for-high-traffic-businesses ; https://www.stratusinfosystems.com/news/cisco-meraki-cw9172i-with-wi-fi-7-should-i-upgrade/ ; https://varindia.com/news/commscope-launches-enterpriseclass-wifi-7-access-point-from-ruckus-networks

## T4. Zyxel XGS/XMG switch street prices 2026 (was: absent or MSRP-only)

Current retailer snapshots for Zyxel multigig/10G lines [independent, all Newegg.com unless noted]:

- **XGS2220-54FP** (48× GbE, 40× PoE+ / 10× PoE++, 960W, 4× SFP+ 10G + 2× 10G Ethernet, L3, Nebula Pro) — **$1,980** street [independent].
- **XGS2220-54HP** (48× GbE, 40× PoE+ / 10× PoE++, 600W, 4× SFP+ 10G + 2× 10G Ethernet) — **$1,539** street [independent]; UK street **£582.87 ex VAT** (comms-express, Sep 2026) — wide regional variance, flag.
  Source: https://www.comms-express.com/products/zyxel-xgs2220-54-48-port-gbe-l3-managed-switch-with-6-10g-uplink/
- **XGS1935-52HP** (48× GbE PoE+, 375W, 4× 10G SFP+, Lite-L3) — **$693** street (Newegg); UK IT Pro review found **£528 ex VAT** (~£100 less than predecessor) [independent].
  Source: https://www.itpro.com/hardware/routers/zyxel-xgs1935-52hp-review-a-port-dense-gigabit-poe-switch-thats-priced-right-for-smbs
- **XGS1935-28HP** (24× GbE PoE+, 4× 10G SFP+, Lite-L3) — **$599.99** (comps); 2,075 AED (UAE) [independent].
- **XMG1915-18EP** (16× 2.5G PoE (8× PoE+ / 8× PoE++, 180W), 2× 10G SFP+, fanless, Nebula) — **$400** street [independent].
- **XGS1210-12** (12-port web-managed multigig: 2× 2.5G + 2× 10G SFP+) — **$249** street [independent].
- **XS1930-10** (10× multigig, 2× 10G Ethernet + 2× 10G SFP+, Nebula optional, rackmount) — **$870** street [independent].
- **GS1920-24v2** (24× GbE + 4× SFP, Nebula cloud-managed) — **$200** street [independent].
- Launch MSRP reference (2025 Business Wire) [vendor-reported]: XGS1935-28 $274.99; XGS1935-52 $599.99; XGS1935-28HP $429.99; XGS1935-52HP $749.99 — street prices above now undercut MSRP on the HP models.
  Source: https://www.businesswire.com/news/home/20250513491345/en/Zyxel-Networks-Smart-Managed-Switches-Deliver-Affordable-Eco-Friendly-10GbE-Migration-to-Small--and-Medium-Sized-Businesses
- **Flag:** GS1915 street prices still unannounced (per S13); XS1935-12HP price still unannounced [unverified].

## T5. Aruba Instant On — 2026 status refresh

- **No new Instant On hardware launch found in 2026** — the SMB line remains: AP22 (Wi-Fi 6, $150), AP25 (Wi-Fi 6 4×4, $199–239), AP27 (outdoor Wi-Fi 6, $239), AP32 (Wi-Fi 6E, $199); switches Instant On 1830 (entry) / 1930 (mid, 4× SFP+ 10G) / 1960 (stackable up to 4, 600W PoE) [secondary].
  Source: https://community.cognetic.com/index.php?topic=85.0
- **No Wi-Fi 7 Instant On AP as of Sep 22, 2026** — HPE's Wi-Fi 7 APs (730 Series, ultra tri-band, launched Apr 2024) are Aruba Central-managed enterprise only [secondary].
  Source: https://convergedigest.com/hpe-aruba-debuts-wi-fi-7-access-points/
- Post-Juniper context (from Phase A file): Mist↔Central convergence targets enterprise, not Instant On [unverified assessment].
- IT Pro review data point [independent]: AP25 delivered 151 MB/s close-range and 131 MB/s at 10 m on 160 MHz channels (vs 80/71 MB/s for AP22), paired with a Zyxel XS1930-12HP 10G PoE++ switch in the test lab.
  Source: https://www.itpro.com/hardware/routers/369878/aruba-instant-on-ap25-review-a-simple-wireless-network-management-choice

