---
id: etape6-phaseb-smb-networking/07-supplementary-complementary-research-pass-round-3-september-/t5-netgear-m4500-25g-100g-pro-av-line-pricing-sep-2026-compl
title: "T5. NETGEAR M4500 — 25G/100G Pro AV line pricing Sep 2026 (complements S6 M4350)"
domain: supplementary-complementary-research-pass-round-3-september-
role: deep-dive
task: pricing
actors: ["Intel", "United States"]
dates: ["2026-02-16"]
keywords: ["pricing", "asic", "cost", "cyber", "license", "optics", "settlement"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1256, 1315]
section: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026"
sha256: 491965abe0bbe078af4071056c157deeb0e784182e1fb3069dc198ad48162432
---

# T5. NETGEAR M4500 — 25G/100G Pro AV line pricing Sep 2026 (complements S6 M4350)

## T5. NETGEAR M4500 — 25G/100G Pro AV line pricing Sep 2026 (complements S6 M4350)

- **M4500-32C** (32× 100G QSFP28, L3, front-to-back airflow): **MSRP $27,048.27 / SHI street $20,509.99** [independent, shi.com]; CZ street 474,022 Kč ex VAT [independent, suntech.cz]
- **M4500-48XF8C** (48× 25G SFP28 + 8× 100G QSFP28): street **$15,869.99** (proavwarehouse, crawled Sep 21, 2026, in stock) [independent]; UK **£17,196.99 ex VAT** (misco.co.uk, crawled Sep 21, 2026, 2 in stock) [independent]; India refurb ₹2,324,458 (xfurbish.com, Sep 5 2026) [independent]
- Positioning: SDVoE-ready Pro AV switching with IGMP Plus (no L3 PIM routing needed), Engage Controller, ProSAFE lifetime warranty — pairs with M4350 for AV-over-IP [vendor-reported via resellers]
- **Flag:** no NEW NETGEAR switch model in 2026 found; M4500/M4350/M4250 remain the current lines (consistent with S6 flag)
- Sources: https://www.shi.com/product/38680664/NETGEAR-M4500-32C-Switch ; https://proavwarehouse.com/infrastructure-and-cabling/netgear/netgear-xsm4556-100nas-m4500-48xf8c-managed-switch-with-48x10g-25g-sfp28/ ; https://www.misco.co.uk/product/Networking/Switch/NETGEAR/NETGEAR-M4500-48XF8C---Switch---L3---Managed---48-?prodid=5362747 ; https://www.xfurbish.com/netgear-m4500-48xf8c-48-ports-l3-managed-switch-74in80y0h9-refurbished?tag=Layer+3

## T6. HPE Aruba CX 6000 — NEW CX-series switch for SMB/retail (NRF 2026) (was: absent)

- **[vendor-reported via NetworkWorld]** HPE announced the **CX 6000** at NRF 2026 (Jan 2026): an 8-port L2 switch for retailers/branches — up to **104 Gbps non-blocking**, **77.3 Mpps** forwarding, HPE Aruba Networking ASIC architecture, programmable **CX Operating System**
- Family: five fixed 1U models, **24 and 48 access ports of 1GbE** with 4 built-in 1GbE SFP uplinks; PoE models 370W (24-port) / 740W (48-port), 802.3at Class 4 (30W/port); non-PoE models available
- Target: point-of-sale, digital signage, IoT, cameras, staff systems in small retail — "the smallest retailers" per HPE product marketing (Gayle Levin)
- **Significance:** CX OS extends downward into SMB/retail switching, below Instant On — watch for channel overlap with Instant On 1830/1930 [independent assessment]
- **Flag:** no street prices found yet; treat as announced, availability unconfirmed [unverified]
- Sources: https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html

## T7. Aruba Instant On — divestiture status (complements S7; was: convergence-only context)

- **DOJ settlement context** [secondary — single source, vendor blog (networkdevicesinc.com), treat cautiously]: the Jun 28, 2025 HPE–Juniper DOJ settlement required HPE to **divest the global Instant On campus/branch WLAN business within 180 days** to a DOJ-approved buyer, plus auction a license to Juniper's Mist AIOps source code
- Per the same source: **Extreme Networks and Tech Mahindra both declined to bid**; **Fortinet offered $5–15M**; a separate investor group offered **$1** to spin it into a standalone entity; a federal judge ruled **Jan 2026** that HPE may continue integrating Juniper while 13 state AGs' challenge to the settlement is reviewed
- **[independent assessment]** Buyer guidance implication: Instant On faces branding/ownership uncertainty; Aruba CX + Aruba Central is the safer enterprise-growth path per the source's assessment — this explains the thin 2026 Instant On hardware news (S7/S18/R3 flags)
- **Flag strongly:** single secondary source with potential commercial bias; cross-verify before citing as fact
- Sources: https://networkdevicesinc.com/community/blog/aruba-vs-juniper-2026-buyers-guide ; https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html (NRF 2026, confirms Rami Rahim runs HPE Networking, Marvis↔Juniper analytics integration)

## T8. TP-Link Omada hardware controller pricing — Sep 2026 snapshots (closes the §1.2/S1/S13 gap)

- **No separate "Omada Pro" controller SKU found** — the standard Omada hardware controllers (OC400/OC300/OC200) manage Pro-line devices; retailer capacity specs apply across the platform [independent/secondary — flag: absence of evidence, not proof]
- **OC400** (top model: 1,000 APs / 200 switches / 100 routers, 2× 10G SFP+, 4× GbE, dual fixed redundant PSUs, 8 GB DDR4) [vendor-reported]:
  - UK **from £572.02** (5 offers compared Sep 1, 2026, yorkshire.com) [independent]
  - **€477.27 ex VAT** (CPC Farnell Ireland, 1 in stock) [independent]
  - ZA **R17,218** (itworkup.co.za) ; Kenya **KSh 98,719** (microless, 3 left) ; UAE **AED 2,487 incl VAT** (microless, 3 left) ; AU **$1,552.13** (Kogan) [independent]
- **OC300** (500 devices / 15,000 clients): **$129.99** (MSRP $161.12, multilink.us, on sale) ; **$160.00** (Best Buy US, free shipping) ; **£309.99 inc VAT** (£258.33 ex, 30 in stock, netxl UK) ; ₱9,659 (pcworx.ph, out of stock) ; ₱22,037 (shopee.ph) [independent]
- **OC200** (100 devices, PoE/micro-USB powered): **₹5,500**/piece (indiamart Mumbai) ; SG **$183.90** (shopee.sg, sold out) [independent]
- All controllers: no license fee, free cloud access — the consistent Omada cost story vs Meraki/paid models (see §6 controller table) [official]
- Sources: https://www.yorkshire.com/marketplace/electronics-222/networking-342/4141c45c-cf62-4b7f-9077-e0fbf8a2b57a ; https://cpcireland.farnell.com/tp-link/oc400/omada-hardware-controller/dp/CS37095 ; https://shop.multilink.us/tp-link-oc200-omada-hardware-controller/ ; https://www.netxl.com/wifi-access-points/tp-link-oc300-omada-cloud-controller/ ; https://www.bestbuy.com/product/tp-link-oc300-omada-hardware-controller-sdn-integrated-2-gigabit-port-1-usb-3-0-port-manage-up-to-500-devices-black/J36QCHCWK6

## T9. EnGenius AirGuard + ECW515 detail (complements §4)

- **ECW536S with AirGuard** (announced Jan 15, 2026, PRNewswire) [vendor-reported]: AirGuard = WIDS/WIPS built into the AP — 24/7 threat detection (evil twins, rogue APs, MITM, RF jammers, flood attempts) via **dedicated scanning radios** without impacting Wi-Fi performance; professional RF spectrum analysis, **zero-wait DFS** for non-disruptive radar-avoidance channel shifts, BLE device scanning; targeted at finance/healthcare/enterprise
- **ECW515** (announced Feb 12, 2026, PRNewswire) [vendor-reported]: Wi-Fi 7 dual-band 2×2:2 wall-plate — up to **3.6 Gbps** aggregate, in-room coverage ~1,000 sq ft, 2.5GbE PoE-in (802.3at), integrated **4-port Gigabit switch with PoE output**, SmartCasting (mobile-to-TV casting), SSID-on-LAN (extends wireless policies to wired), carrier-class Wi-Fi calling, VLAN per-user segmentation; targets MDU/student housing/senior living/hotels
- Sources: https://www.morningstar.com/news/pr-newswire/20260115ph63498/engenius-unveils-cloud-managed-wi-fi-7-enterprise-ap-with-247-airguard-security ; https://www.advfn.com/stock-market/stock-news/97823066/engenius-ecw515-brings-wi-fi-7-performance-to-in-r

## T10. Round-3 verification log

1. D-Link DAP-E3620 prices: too fresh (Aug 19, 2026 announcement) — no retailer listings by Sep 22; re-check later [open].
2. Zyxel USG FLEX 50H launch date (2026-02-16) comes from Bonanza marketplace metadata, not a Zyxel press release — verify before citing as launch fact [unverified].
3. Grandstream GWN7813 £172.70 / GWN7816 £359.40 (voipon UK) — UK prices; regional variance vs US/CZ/AE/IN/AU captured but not normalized [noted].
4. Instant On divestiture claims (T7) rest on a single vendor-blog source — **cross-verify independently before use** [unverified].
5. Aruba CX 6000 (T6) — announced at NRF 2026 (Jan); no street price or confirmed GA date found [unverified].
6. "No separate Omada Pro controller SKU" (T8) is an absence-of-evidence finding from retail/TP-Link listings — do not state as a confirmed TP-Link policy [unverified].
7. NETGEAR M4500-32C SHI data — page last crawled 61 days ago; MSRP $27,048.27 vs street $20,509.99 should be re-verified live [flag].
8. ECW536S/ECW515 announcements are Jan/Feb 2026 PRNewswire releases re-hosted on aggregator sites; dates taken from the original PR datelines [vendor-reported].
9. Zyxel USG FLEX H still has no 2026 *platform refresh* — 50H/50HP is a range extension, not a generation change [assessment].
10. Phase C scope (FS.com, 25G/40G enterprise optics, Spine/Leaf hardware) remains intentionally excluded here.

**Round-3 sources (verbatim):** https://aapnews.aap.com.au/aapreleases/cision20260819AE28553 ; https://aapnews.aap.com.au/aapreleases/cision20251218AE49456 ; https://www.bonanza.com/listings/Zyxel-USGFLEX50H-Cyber-Security-Firewall-2-Gbps-Up-to-25-Users-Hardware-Onl/1801209230 ; https://www.manilarepublic.com/zyxel-usg-flex-50hp-firewall-wins-taiwan-excellence-award/ ; https://info.zyxel.com/hubfs/USG%20FLEX%20H%20Comparison_Partners_Rev2026.pdf ; https://www.einpresswire.com/article_pdf/838918940/zyxel-networks-firmware-enables-zero-touch-nebula-deployment-for-usg-flex-h-series-firewalls ; https://www.voipon.co.uk/grandstream-gwn7813-managed-network-switch-p-10157.html ; https://www.discomp.cz/grandstream-gwn7816-layer-3-managed-network-switch-48-portu-6x-sfp-_d127143.html ; https://gear-up.me/grandstream-gwn7813p-layer-3-managed-network-poe-switch-24-ports-4-sfp.html ; https://www.shi.com/product/38680664/NETGEAR-M4500-32C-Switch ; https://proavwarehouse.com/infrastructure-and-cabling/netgear/netgear-xsm4556-100nas-m4500-48xf8c-managed-switch-with-48x10g-25g-sfp28/ ; https://www.misco.co.uk/product/Networking/Switch/NETGEAR/NETGEAR-M4500-48XF8C---Switch---L3---Managed---48-?prodid=5362747 ; https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html ; https://networkdevicesinc.com/community/blog/aruba-vs-juniper-2026-buyers-guide ; https://www.yorkshire.com/marketplace/electronics-222/networking-342/4141c45c-cf62-4b7f-9077-e0fbf8a2b57a ; https://cpcireland.farnell.com/tp-link/oc400/omada-hardware-controller/dp/CS37095 ; https://www.netxl.com/wifi-access-points/tp-link-oc300-omada-cloud-controller/ ; https://www.morningstar.com/news/pr-newswire/20260115ph63498/engenius-unveils-cloud-managed-wi-fi-7-enterprise-ap-with-247-airguard-security ; https://www.advfn.com/stock-market/stock-news/97823066/engenius-ecw515-brings-wi-fi-7-performance-to-in-r

---

