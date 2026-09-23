---
id: etape6-phaseb-smb-networking/13-supplementary-complementary-research-pass-september-22-2026-/overview
title: "Supplementary / Complementary Research Pass — September 22, 2026 (fourth wave; U1–U7)"
domain: supplementary-complementary-research-pass-september-22-2026-
role: deep-dive
task: reference
actors: ["China", "EU", "Qualcomm", "United States"]
dates: ["2026-03", "2026-03-05", "2026-09-22"]
keywords: ["research", "asic", "consumer", "ethernet", "license", "pricing", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2113, 2203]
section: "Supplementary / Complementary Research Pass — September 22, 2026 (fourth wave; U1–U7)"
sha256: f53e54afee9a5dd1c73c99e9107af41e38b4b42f27855480857a5a8ce14dd377
---

# Supplementary / Complementary Research Pass — September 22, 2026 (fourth wave; U1–U7)

**Scope:** fourth complementary wave on Phase B (SMB/prosumer). Fills: NETGEAR 2026 hardware (ISE 2026), Zyxel GS1915 + XS1935-12F street prices, Grandstream GWN7670 GA/price, EnGenius ECW510 pricing detail, TP-Link Omada Pro controller (explicit finding), Aruba Instant On 2026 status re-check. No existing section was modified.

Provenance tags: [official] = vendor price list/store/datasheet; [vendor-reported] = vendor claim via press release/trade press; [independent] = third-party/retailer listing; [secondary] = press/blog; [unverified] = single weak source or inferred.

## U1. NETGEAR — NEW 2026 hardware: M4350-16M4V + M4350-16C at ISE 2026 (corrects "no confirmed NEW NETGEAR switch model in 2026" flag)

NETGEAR (NASDAQ: NTGR) announced two new IP switches in its M4350 series plus Engage Controller v2.4 at **ISE 2026 (Jan 27, 2026)** [vendor-reported via BusinessWire syndication]:

- **M4350-16M4V** — ruggedized model: industry-standard **Neutrik locking connectors** throughout (etherCON for network, opticalCON for fiber, powerCON for power) to prevent connections coming loose during load-in/mid-show [vendor-reported]. 16× 2.5G PoE++ ports (up to **1,130W** total PoE budget), 4× 25G SFP28 uplinks, plus **four modular interface card slots** for configurable uplink connectivity [vendor-reported].
  Source: https://www.ncnonline.net/netgear-introduces-next-gen-m4350-switches-with-rugged-design-and-offline-av-ready-network-configuration-at-ise-2026/
- **M4350-16C** — **16 ports of 100G**, designed for aggregation and core layers where multiple high-resolution video streams converge; addresses throughput challenges in large AV-over-IP deployments [vendor-reported].
  Source: https://etedge-insights.com/trending/ise-2026-netgear-introduces-next-gen-m4350-switches-for-pro-av-networks/
- Portfolio context: the two additions bring the M4350 family to **18 models total**, covering 1G to 100G with high-power PoE++, redundant modular power supplies, and SMPTE ST 2110 timing support; **TAA-compliant versions** available for government projects; new models available **March 2026** [vendor-reported].
  Source: http://business.minstercommunitypost.com/minstercommunitypost/article/bizwire-2026-1-27-netgear-brings-new-ruggedized-switching-and-anywhere-anytime-network-configuration-to-ise-2026
- **Engage Controller v2.4** (Feb 2026): introduces **offline provisioning** — integrators create virtual switches and Wi-Fi 7 access points, design entire networks and test configurations without physical hardware; export and deploy on-site; configurations saved as reusable templates [vendor-reported].
  Source: https://www.ncnonline.net/netgear-introduces-next-gen-m4350-switches-with-rugged-design-and-offline-av-ready-network-configuration-at-ise-2026/
- Pricing for the two new models: not captured in sources found — **flag: prices unconfirmed** [unverified]. (S6 pricing for five earlier M4350 models stands: $838.99–$8,689.64 street [independent].)

## U2. Zyxel GS1915 — street prices surface (resolves T7 gap #3 for GS1915); XS1935-12F shipping

The 2026 "Just Connect" re-announcement covers the existing GS1915 8/24-port family. Current street snapshots [independent]:

| Model | Street (Sep 2026 snapshot) | Retailer |
|---|---|---|
| GS1915-8EP (8× GbE PoE+, 60W) | **€99.99** (was €179.99) | bestdigit.it (IT) |
| GS1915-24E (24× GbE) | **€125.71** | eestipoisid.eu (EE, listed 25-Jul-2026) |
| GS1915-8EP | **₱9,200** (~$160) | pcx.com.ph (PH) |
| GS1915-8 (non-PoE) | **€93.01** | shopinbit.com (EU) |
| GS1915-24EP (24× GbE PoE+) | **£176.26 ex VAT / £211.51 inc VAT** | screenmoove.com (UK) |
| GS1915-24E | $94.17 / $169.98 (out of stock; dual-price listing, ambiguous) | panacompu.com (PA) — **flag: ambiguous listing** |

Sources: https://bestdigit.it/en-es/products/zyxel-gs1915-8ep-switch-di-rete-gestito-l2-8-porte-gigabit-poe-60w-nero ; https://eng.eestipoisid.eu/computers-and-equipment/network-equipment/other-network-products/zyxel-switch-gs-1915-gs1915-24e-eu0101f/170634/2/p?catalog_print=1 ; https://pcx.com.ph/collections/pc-peripeherals-accessories/products/zyxel-gs1915-8ep-8-port-gbe-smart-managed-switch ; https://shopinbit.com/Switch-ZyXEL-GS1915-8-EU0101/23-S55165326 ; https://screenmoove.com/products/zyxel-l2-gigabit-ethernet-12-port-poe-managed-switch-gs1915-24ep-gb0101f

- **XS1935-12F** (all-fiber, 12× 10G SFP+): shipping in EU — **€448.99** (shopro.lt, Lithuania, EAN 4718937637805) [independent].
  Source: https://www.shopro.lt/index.php?route=product/pdf&product_id=540230
- **XS1935-12HP and XS1935-10 street prices: still unannounced.** The Tech Revolutionist (updated Sep 2026): "Zyxel has not announced regional pricing or a detailed market-by-market availability schedule" — check local channels [secondary].
  Source: https://thetechrevolutionist.com/2026/09/zyxel-xs1935-10gbe-poe-switches-smb-networks.html
- Related: Zyxel 2026 product portfolio PDF lists WBE660S (BE22000 Wi-Fi 7, 10G PoE), WBE530 (BE11000), NWA130BE, XMG2230-28HP, XS3800-28 aggregation switch, USG FLEX 700H as headline SMB products [official catalog].
  Source: https://www.zyxel.com/sites/zyxel/files/solution/Product_Portfolio_2026.pdf

## U3. Grandstream GWN7670 — GA confirmed + EU street price (corrects T7 gap #1)

- **GWN7670 is Grandstream's first Wi-Fi 7 AP (dual-band 2×2:2, up to 3.6 Gbps aggregate wireless, 5 Gbps aggregate wired)**; announced **~Feb 2025** per press-coverage dates ("593 days ago" relative to Sep 22, 2026) — i.e. the GWN767x family predates 2026, correcting the earlier "brochure dated 2025" flag to a confirmed GA lineage [vendor-reported/secondary].
  Source: https://telconews.asia/story/grandstream-unveils-gwn7670-its-first-wi-fi-7-access-point
- Features: MLO, 4K-QAM, multi-RU, preamble puncturing, up to 256 clients, 175 m coverage, BLE 5.3, PoE+ self-power adaptation, embedded controller (up to 50 local GWN APs), GDMS Networking / GWN Manager managed [vendor-reported].
  Source: https://rlan.bg/en/29127-grandstream-gwn7670-wi-fi-7-access-point
- **EU street price: €118.92 incl. 20% VAT (≈ €99.10 ex VAT)** — in stock (<10 units) at rlan.bg (Bulgaria), planned delivery **30-Sep-2026**, 36-month warranty [independent].
  Source: https://rlan.bg/en/29127-grandstream-gwn7670-wi-fi-7-access-point
- GWN7670LR (outdoor) GA date still unverified [unverified]. GWN7303 status unknown [unverified]. Tri-band Wi-Fi 7 models were promised "in the coming months" (2025 launch material) — current 2026 status not confirmed [unverified].
  Source: https://www.grandstreamindia.co.in/blog/grandstream-wifi7-access-point/

## U4. EnGenius ECW510 — $129 MSRP detail (complements §4)

- **ECW510** — affordable Wi-Fi 7 AP for SMB, MSRP **$129** [vendor-reported via PRNewswire]: Qualcomm Networking Pro 1220 platform, dual-band 2×2 Wi-Fi 7 up to **5 Gbps aggregate**, 2.5G Ethernet, PoE+, 18.7W power draw, 1,000 sq ft coverage, up to 400 clients, WPA3, license-free cloud management, 5-year warranty; shipping through authorized resellers [vendor-reported].
  Source: https://pr.gulfmainmagazine.com/article/EnGenius-Brings-Wi-Fi-7-to-Small-Businesses-with-Affordable-ECW510-Access-Point/68b04dcef5f91c4266c32f44
- ECW536S (Jan 15, 2026: 4×4×4 Wi-Fi 7, 18.8 Gbps aggregate, 10 GbE PoE++, AirGuard 24/7 WIDS/WIPS, Qualcomm Networking Pro 1220) and ECW515 wall-plate (Feb 12, 2026: 3.6 Gbps, 2.5 GbE PoE-in, integrated 4-port GbE switch with PoE-out) from main §4 stand confirmed [vendor-reported].
  Source: https://www.morningstar.com/news/pr-newswire/20260115ph63498/engenius-unveils-cloud-managed-wi-fi-7-enterprise-ap-with-247-airguard-security ; https://www.advfn.com/stock-market/stock-news/97823066/engenius-ecw515-brings-wi-fi-7-performance-to-in-r

## U5. TP-Link Omada Pro controller — explicit finding: no Omada-Pro-branded controller SKU

- **No Omada-Pro-branded hardware controller SKU exists publicly.** TP-Link's Omada Pro line (S5500 switches) is managed by the standard Omada controller stack — Omada Pro *features* (Anomaly Detection, OUI VLANs, now trickled down to standard Omada v6.3) run on OC200/OC220/OC300/OC400 hardware and the software/cloud controllers [independent assessment from catalog gap — flag: no TP-Link page confirms a "Pro controller" product].
- Current hardware-controller street snapshots (Sep 2026) [independent]:
  - OC200: **€130.01** incl. tax (tdtprofesional.com, ES). Source: https://www.tdtprofesional.com/en/controller-ap-for-omada-oc200-of-tp-link.html
  - OC300: **$160.00** (Best Buy, US). Source: https://www.bestbuy.com/product/tp-link-oc300-omada-hardware-controller-sdn-integrated-2-gigabit-port-1-usb-3-0-port-manage-up-to-500-devices-black/J36QCHCWK6
  - OC300: **$187.49** (supercircuits.com). Source: https://www.supercircuits.com/tp-link-omada-hardware-controller-oc300
  - A 4-port Omada hardware controller with 10G capability: **$582.30** (Klarna/Best Buy listing; SKU lineage ambiguous — could be OC400-class; do not label as OC400 without confirmation) [independent]. Source: https://www.klarna.com/us/shopping/pl/cl167/3393542695/Switches/TP-Link-Omada-Hardware-Controller-4-Ports/
  - Omada Cloud-Based Controller: **₱499.00** listed (pinklehub, PH; sold out — tier/pricing terms unclear) [independent]. Source: https://pinklehub.com/products/tp-link-omada-cloud-based-controller-i-omada-cloud-based-controller
- **Gap retained:** dedicated Omada Pro controller licensing/pricing not found — appears not to exist as a separate product [unverified].

## U6. Aruba Instant On — 2026 status re-check (explicit finding: no new hardware; CX 6000 is adjacent enterprise)

- **No new Aruba Instant On hardware launch found in 2026** [explicit finding — searched Sep 22, 2026]. The 1930 line (JL680A–JL686A, incl. quieter JL683B/JL684B/JL686B revisions per 2.5.0.42 release notes) remains the current Instant On switching portfolio; latest 1930 firmware tracked: **2.9.1.17** (7th release of major version 2.0) [official].
  Sources: https://www.arubanetworks.com/techdocs/InstantOn_1930_Switch/rn_1930_2.9.1.17.pdf ; https://www.buyrouterswitch.com/jl683a-price.html
- Adjacent enterprise news: at **NRF 2026 (Jan 2026)** HPE announced a new **Aruba CX 6000** L2 switch (8-port, up to 104 Gbps non-blocking, 77.3 Mpps, HPE Aruba ASIC architecture, programmable CX OS; 24/48-port models with 1GbE SFP uplinks and Class 4 PoE up to 740W) aimed at retailers/branch — **CX, not Instant On**; noted here to avoid confusion [secondary].
  Source: https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html
- Instant On 1960 12-port multi-gig SKU S0F35A from S18 stands; no "1930R successor" surfaced [unverified].

## U7. Fourth-wave verification log + remaining gaps

1. NETGEAR M4350-16M4V / M4350-16C street prices: not found — pending.
2. Zyxel XS1935-12HP / XS1935-10 street prices: still unannounced per Zyxel press coverage (Sep 2026).
3. Grandstream GWN7670LR GA date; tri-band Wi-Fi 7 models' 2026 status; GWN7303 status — unconfirmed.
4. TP-Link Omada Pro: no Pro-branded controller product found — treat as "not a separate product" [unverified].
5. D-Link March 5, 2026 "Wi-Fi 7 Practical" announcement — consumer/home portfolio, out of SMB scope; not researched.
6. MikroTik China-supply exposure to 2026 US tariff/FCC action — unknown.
7. Currency conversions not applied; all prices raw regional snapshots.

---

