---
id: etape6-phaseb-smb-networking/07-supplementary-complementary-research-pass-round-3-september-/pass-b-tp-link-omada-pro-agile-d-link-nuclias-zyxel-usg-flex
title: "Pass B — TP-Link Omada Pro/Agile, D-Link Nuclias, Zyxel USG Flex / XGS2220 / GS2200 / Wi-Fi 7 APs"
domain: supplementary-complementary-research-pass-round-3-september-
role: deep-dive
task: reference
actors: ["AWS", "EU", "Google", "Intel", "Malaysia", "Microsoft", "Qualcomm", "United States"]
dates: ["2023-03", "2023-10", "2025-04", "2025-07", "2025-10", "2025-12", "2026-01", "2026-01-07", "2026-01-15", "2026-01-21", "2026-03", "2026-03-05", "2026-03-17", "2026-04-09", "2026-04-14", "2026-05", "2026-06-18", "2026-06-26", "2026-07-02", "2026-07-03", "2026-07-29", "2026-08", "2026-09", "2026-09-08", "2026-09-13", "2026-09-16", "2026-09-20", "2026-09-22"]
keywords: ["advisory", "cost", "ethernet", "license", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1316, 1402]
section: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026"
sha256: 984f8716ea0d7ab2b979b429909ac69bc1bf2786d06666a3802350615ea202dc
---

# Pass B — TP-Link Omada Pro/Agile, D-Link Nuclias, Zyxel USG Flex / XGS2220 / GS2200 / Wi-Fi 7 APs

### Pass B — TP-Link Omada Pro/Agile, D-Link Nuclias, Zyxel USG Flex / XGS2220 / GS2200 / Wi-Fi 7 APs

Provenance tags: [official] = vendor's own site/docs; [vendor-reported] = vendor PR statement; [independent] = retailer/review; [secondary] = press/distributor/retailer page; [unverified] = single weak or contradictory source. Cutoff: September 22, 2026.

#### B.1. TP-Link Omada Pro 2026

**S5500 street prices/availability.** US retailer CompSource lists Omada Pro S5500 models (all "Date Inserted: Thursday January 15, 2026"; listings updated June–September 2026) [independent]:
- S5500-8XF (8× 10GE SFP+ L2+): $1,010.23 (listing updated June 26, 2026). Source: https://www.compsource.com/buy/S55008XF/Tp-Link-3623/TPLink-Omada-Pro-8Port-SFP-L2-Managed-Switch--Manageable--10-Gigabit-Ethernet--10GBaseX--2--S55008XF/
- S5500-48GP4XF (48× GbE PoE+, 4× SFP+, 500 W PoE): listed at $3,469.77 in one listing (updated June 18, 2026) and $3,849.99 in a second listing (updated July 3, 2026) — price varies by listing; treat both as snapshots, not confirmed street MSRP. Source: https://www.compsource.com/buy/S550048GP4XF/Tp-Link-3623/TPLink-Omada-Pro-48Port-PoE-Gigabit-L2-Managed-Switch-with-4-SFP-Slots--48-Ports--Manageable--S550048GP4XF/
- S5500-8MHP2XF (8× 2.5G PoE+, 2× SFP+, 240 W PoE): $1,869.99 (listing updated July 3, 2026). Source: https://www.compsource.com/buy/S55008MHP2XF/Tp-Link-3623/TPLink-Omada-Pro-8Port-PoE-25G-L2-Managed-Switch-with-2-SFP-Slots--8-Ports--Manageable--Gig-S55008MHP2XF/
- S4500-8GP2F (Omada Pro 8-port PoE+ GbE smart switch, 61 W PoE — smaller sibling line): $552.73 (listing updated September 16, 2026). Source: https://www.compsource.com/buy/S45008GP2F/TP-LINK-3623/TPLink-Omada-Pro-8Port-PoE-Gigabit-Smart-Switch-with-2-SFP-Slots--8-Ports--Manageable--Gigabit-S45008GP2F/
- A Bulgarian reseller (virtualnazona.com) lists S5500-48GP4XF, S5500-8MHP2XF and S5500-8XF but shows "No stock" [independent]: https://virtualnazona.com/en/48-portov-gigabiten-poe-l2-upravlyaem-komutator-tp-link-omada-pro-s5500-48gp4xf-s-4-sfp-porta.html
- **Availability caveat:** models are orderable in the US via B2B retail; EU/EMEA retail availability appears spotty ("no stock" at one EU reseller). No official MSRP was found for any S5500 model — flagged unknown.

**Omada Pro controller 2026.** TP-Link knowledge-base docs from August 2026 reference "Omada Pro Controller v1.9.20.7" (Aug 12, 2026) [official]: https://support.omadanetworks.com/uk-ua/document/13280/. S7500-series firmware v1.20 release notes (Dec 2025, file dated 20251219) recommend "Omada Pro Controller V1.10.20" and note that v1.20 rebrands the S7500 line's VI "to Omada Central" (rename of the series) [official]: https://static.tp-link.com/upload/firmware/2025/202512/20251219/S7500-26XF6Y(UN)_v1.20_1.20.0%20Build%20251107%20Release%20Note.pdf. Separately, an Omada knowledge-base article dated 08-28-2026 confirms that **starting from Omada Controller V6.2, standard Omada controllers can discover, adopt and mesh Omada Pro devices** in the same site (limited features; more Pro features expected in V6.3) [official]: https://support.omadanetworks.com/us/document/120285/.

**2026 Omada Pro announcement (adjacent).** At TP-Link Malaysia's APAC B2B Summit on 15 September 2026, TP-Link unveiled an **Omada WiFi 8 line** for business in five flavours (Max, Pro X, Pro, UR, Basic); the Pro model is expected in H1 2027, rest later 2027 — future product, noted here only as announced context [secondary]: https://techminit.com/tp-link-malaysia-announces-wifi-8-lineup-featuring-archer-8-ultra-and-deco-8-ultra/. Omada Network v6.3 (announced early September 2026) adds Wireless IDS/IPS, AP grouping, rapid device replacement, and **OUI-based VLANs now compatible with Agile Switches** (requires compatible Agile firmware) [vendor-reported via community]: https://community.tp-link.com/en/business/forum/topic/871456.

**"Agile Switches" details.** "Agile" = TP-Link's Easy Managed switch class (not a separate Pro line) [official]: https://www.omadanetworks.com/us/business-networking/omada-switch-agile/. The 2026 ISC West flyer (March 2026) lists the Agile series lineup [official]:
- ES205GP (5× GbE, 4× PoE+, 65 W budget), ES210GMP (10-port, 8× PoE+, 123 W), ES220GP (20-port, 16× PoE+, 150 W), ES228GP (28-port, 24× PoE+, 250 W), ES206XPP-M2 (6-port 2.5G with 4× PoE++, 120 W + 1× 10G SFP+). Feature pitch: cost-effective SDN compatibility, up-to-820 ft long-range PoE on gigabit models, IntelliRecover remote reboot. Source: https://static.tp-link.com/document/pdf/en/isc-west-2026/2026_ISC_WEST_Omada_Business_Switches.pdf
- The full Agile table adds: ES210XPP-M2 (8× 2.5G PoE++, 200 W), ES210X-M2, ES206X-M2 (2.5G non-PoE), ES228GMP (384 W), ES224G, ES220GMP, ES216G, ES210GP (63 W), ES208GP/ES208G, ES206GP. All desktop/wall-mount except 28-port models (rackmount); nearly all fanless. Source: https://www.omadanetworks.com/us/business-networking/omada-switch-agile/
- Feature limits (Easy Managed class): **no 802.1X control, no 802.1p priority, no LLDP-MED, no DHCP L2 relay, no trust mode** on port profiles [independent, from controller UI text]: https://github.com/daily-nerd/terraform-provider-omada/blob/HEAD/docs/SWITCH_CLASS_MATRIX.md. Agile switches require explicit per-site activation in the controller (community Q&A, Jan 2026) [secondary]: https://community.tp-link.com/en/business/threads/topic/853584.
- **Street prices for Agile models: not found** — flagged unknown.

#### B.2. D-Link Nuclias 2026

**2026 platform releases.** The next-gen Nuclias hardware controllers (DNH-1000 up to 500 devices, DNH-3000 up to 1,500 devices, software-based DNC-5000) were announced **October 2025** (PR Newswire) — pre-cutoff context, included only for continuity [vendor-reported]: https://beta.manilatimes.net/2025/10/16/tmt-newswire/pr-newswire/d-link-unveils-nuclias-network-controllers/2202000. **No 2026-versioned Nuclias Connect software release was found** — flagged unknown (latest version number not identified in 2026 sources).

**2026 Nuclias-compatible hardware.**
- **DAP-E9560 BE9500 Wi-Fi 7 ceiling-mount AP** and **DAP-X3060W AX3000 Wi-Fi 6 wall-plate AP** announced January 7, 2026 [vendor-reported]: DAP-E9560 = tri-band Wi-Fi 7, 320 MHz channels, 4096-QAM, 1× 10GbE PoE + 1× 2.5GbE LAN, WPA3 Enterprise, Nuclias-managed; DAP-X3060W = dual-band AX3000 in-room wall-plate, PoE, multiple GbE ports for IP phones/TVs. Source: https://facilityexecutive.com/d-link-unveils-next-generation-enterprise-wireless-lineup. Note: the PR wire copy carries a December 2025 wire date while trade press dates it Jan 7, 2026 — treat as late-Dec 2025/early-Jan 2026 announcement [secondary]: https://aapnews.aap.com.au/aapreleases/cision20251218AE49456.
- **Prices for DAP-E9560 / DAP-X3060W: not found** — flagged unknown.
- D-Link Japan's "Product Guide 2026-27" (May 2026) lists as planned releases (リリース予定): **DAP-E9560** (Wi-Fi 7 tri-band 2×2×2, 10G PoE 802.3bt), **DAP-E3620** (Wi-Fi 7 2×2), **DAP-E3620OU** (Wi-Fi 7 2×2 outdoor) — i.e., two additional Wi-Fi 7 APs still in "planned" status as of mid-2026 [vendor-reported]: https://files.actibookone.com/contents/4200/737165-20260519113745/original.pdf
- Nuclias-adjacent (surveillance, July 29, 2026): D-Link expanded the Nuclias surveillance portfolio for SMBs with BCB-P01/BCD-P01 Full HD IP cameras and BNR-004/BNR-008 NVRs [vendor-reported]: https://uktechnews.co.uk/2026/07/29/d-link-expands-nuclias-surveillance-portfolio-for-small-businesses/.

**D-View 8 status.** Product status "Live" (Revision A1) on D-Link UK as of 2026 [official]: https://www.dlink.com/uk/en/products/dv-800-d-view-8-network-management-software. Latest release identified: **v2.0.5.109, dated 2025/10/31** (server/probe/MongoDB, per D-Link Australia release notes) [official]: http://files.dlink.com.au/DV-800S/REV_A/Software/DV-800_Software_Release_Notes_v2.0.5.109(WW).pdf. **No 2026 D-View 8 release was found** — latest known remains the Oct 2025 build. 2026 security advisories: **CVE-2026-23754** (CVSS 8.8, improper access control → full admin takeover) and **CVE-2026-23755** (CVSS 7.3, installer DLL preloading), disclosed January 21, 2026, affecting versions "2.0.1.107 and below", patched in 2.0.1.108+ [secondary]: https://stack.watch/product/dlink/d-view-8/ and https://app.opencve.io/cve/?vendor=dlink&product=d-view_8. **Inconsistency flagged:** the advisory cites "2.0.1.107 and below" as affected even though 2.0.5.109 shipped Oct 2025 — version-branch relationship unclear; verify against D-Link's official advisory before acting.

#### B.3. Zyxel USG Flex / security 2026

**Lineup & street prices.** USG FLEX H series (50H/100H/200H/500H/700H) street prices published in the April 2025 launch release [vendor-reported]: USG FLEX 200H $399.99 ($549.99 bundled with 1-year Gold Security License Pack); 200HP $499.99 ($649.99 bundled); 500H $799.99 ($1,099.99 bundled); 700H $1,299.99 ($1,699.99 bundled); 100H $299.99 ($399.99 bundled); 100HP $399.99 ($499.99 bundled). Source: https://businesswire.com/news/home/20250415755419/en/Zyxel-Networks-Upgrades-Leading-Firewall-Family-to-Provide-SMBs-with-Unified-Cloud-and-On-Premises-Security. A January 2026 press item shows slightly higher street prices (200H $449.99 / bundled $599.99; 500H $839.99 / $1,199.99; 700H $1,299.99 / $1,699.99) [vendor-reported]: https://www.einpresswire.com/article_pdf/838918940/zyxel-networks-firmware-enables-zero-touch-nebula-deployment-for-usg-flex-h-series-firewalls. Implied 1-year Gold Security License Pack value ≈ $150 (200H) to $400 (700H) by arithmetic on both snapshots — derived, not officially priced. **Current Sep 2026 retail street prices: not independently verified** — flagged.

**uOS firmware releases 2026** (USG FLEX H series) [official release notes, download.zyxel.com]:
- **uOS V1.37, released 2026/01/15**: Application-Aware Policy Routing, SSL VPN/Captive Portal auth via Microsoft Entra ID / Google OIDC, IPsec IKEv2 AES-GCM + DH groups 31–32, anomaly detection & prevention, anti-malware SHA-256 allow/block lists, mDNS proxy (AirPlay/AirDrop/Chromecast cross-subnet), interface ingress/egress rate limiting. Source: https://download.zyxel.com/USG_FLEX_50H/firmware/USG%20FLEX%2050H_1.37(ACLO.0)C0_2.pdf
- **uOS V1.38 Patch 0, released April 09, 2026** (consolidates V1.37/A___.1, V1.37/A___.0, V1.36, V1.35.x, V1.32 feature sets). Source: https://download.zyxel.com/USG_FLEX_700H/firmware/USG%20FLEX%20700H_1.38(ABZI.0)C0_2.pdf
- uOS 1.35 (early Jan 2026, per press release ~Jan 1, 2026) enabled zero-touch Nebula deployment/preconfiguration for H-series, DNS Safe Search and expanded auth options [vendor-reported]: https://www.einpresswire.com/article_pdf/838918940/zyxel-networks-firmware-enables-zero-touch-nebula-deployment-for-usg-flex-h-series-firewalls
- The V1.38 AP-controller supported-AP list adds **WBE665S** and **IAP500BE** alongside WBE660S/WBE630S/WBE530/WBE510D/WAX series [official, via release-note mirror]: https://download.zyxel.com/USG_FLEX_700H/firmware/USG%20FLEX%20700H_1.38(ABZI.0)C0_2.pdf
- Performance (official datasheet, 09/04/26 rev): 500H — SPI 10 Gbps, VPN 2 Gbps, IPS 4.5 Gbps, UTM 3 Gbps, 1M sessions, 300 IPSec tunnels; 700H — SPI 15 Gbps, VPN 3 Gbps, IPS 7 Gbps, UTM 4 Gbps, 2M sessions, 1,000 IPSec tunnels [official]: https://www.studerus.ch/fr/support/download/212358.

#### B.4. Zyxel XGS2220 / GS2200 series 2026

**XGS2220 firmware 2026.** Zyxel Download Library shows **firmware V5.00 patch 1 released March 05, 2026** across XGS2220-30/30F/30HP/54/54HP/54FP (e.g. 5.00(ABXQ.1)C0 for -54HP), plus refreshed datasheets April 14, 2026 and CLI guides July 2, 2026 [official]: https://www.zyxel.com/us/en-us/support/download?model=xgs2220-30F. The V5.00 base (July 2025) added: enlarged L3 forwarding table (2K IPv4 + 2K IPv6), IPv4 policy-based routing, SSH public-key auth, forced password change on first login, Telnet disabled by default, HTTP→HTTPS redirect, AV-optimized (Dante/AES67) mode, unlockable local GUI in cloud mode [official]: https://community.zyxel.com/en/discussion/30171/nebulaflex-switch-xgs2220-series-v5-00-patch-0-firmware-release. **New XGS2220 SKUs in 2026: none found** (six-model family from March 2023 unchanged: -30/-30F/-30HP/-54/-54HP/-54FP) — flagged as no-change.

**XGS2220 street-price snapshots (Sep 2026)** [independent]:
- XGS2220-30HP (24-port GbE L3 PoE+, 6× 10G uplink, 400 W): **$1,140.00** at SHI (page "updated 6 hours ago" at crawl → Sep 2026): https://www.shi.com/product/45906824/Zyxel-XGS2220-Series-XGS2220-30HP
- XGS2220-54HP (48-port, 600 W): **$1,744.00** at SHI (updated 14 hours ago): https://www.shi.com/product/45906827/Zyxel-XGS2220-Series-XGS2220-54HP
- UK: XGS2220-30HP at **£488.76 ex VAT** (£586.51 inc) at Comms Express, but listed as discontinued there: https://www.comms-express.com/products/zyxel-xgs2220-30hp-24-port-gbe-l3-managed-poe-switch-with-6-10g-uplink-400-w/
- For reference, 2023 launch street prices: -30 $849.99, -30F $999.99, -30HP $1,249.99, -54 $1,249.99, -54HP $1,649.99 [vendor-reported]: https://www.zyxel.com/us/en-us/newsroom/press-releases/zyxel-networks-launches-layer-3-access-switch-family

**GS2200 / GS2220 status.**
- **GS2200 series is End-of-Life** — per Zyxel Community (accepted solutions, 2024–2025 threads): GS2200-8HP latest firmware 4.00(AAAW.4)C0; GS2200-48 latest V3.80(BPR.3); Zyxel staff recommend the GS2220 series as successor. **No 2026 GS2200 firmware exists** [official-via-community]: http://community.zyxel.com/en/discussion/30499/gs2200-8hp-firmware-and-reset-files and http://community.zyxel.com/en/discussion/comment/67161
- GS2220 street prices (UK, e-catalog, Sep 2026 crawl): GS2220-10 from £162.68, -10HP £233.99, -28 £302.06, -28HP £327.49–£392.89, -50 £438.49, -50HP £644.99 [independent]: https://e-catalog.co.uk/ZYXEL-GS2220-50HP.htm and https://www.pricerunner.com/pl/167-3202694026/Switches/Zyxel-GBPGS2220-28HP-EU0101F-Managed-Compare-Prices
- 2023-era US launch street prices for reference: GS2220-10 $179.99, -28 $389.99, -50 $629.99, -10HP $279.99, -28HP $649.99 [vendor-reported]: https://www.businesswire.com/news/home/20200708005109/en/Zyxel-Hybrid-Switches-Feature-Industry-First-Tri-Mode-Management-Capability
- **GS2200/GS2220 2026 firmware updates: none found** — no new firmware releases surfaced for either series in 2026 (active development is on XGS2220 V5.00 and newer families).

#### B.5. Zyxel Wi-Fi 7 APs 2026

**WBE660S.** Not a 2026 launch (launched October 2023) [vendor-reported]: https://aijourn.com/zyxel-networks-launches-22gbps-wifi-7-access-point/. Specs: BE22000 triple-radio 4×4×4 (2.4/5/6 GHz), 320 MHz, Qualcomm Networking Pro 1220, 1× 10GbE (1/2.5/5/10G) PoE + 1× 1GbE, up to 8 SSIDs / 1,000+ clients, wall-mountable, NebulaFlex Pro, no power supply included [independent review]: https://www.techradar.com/pro/zyxel-wbe660s-review. Street prices Sep 2026 [independent]: SHI US $553–$606 (MSRP listed $499.99–$674.99 across SHI storefronts — inconsistent; flagged); SHI Germany €620.00; SHI France €658.00 (2 in stock); Poland 2,775.47 zł; TechRadar notes an Amazon special at $499 against a $799 price tag. Sources: https://www.shi.com/product/47189848/Zyxel-NebulaFlex-Pro-WBE660S, https://www.de.shi.com/product/47253730/Zyxel-NebulaFlex-Pro-WBE660S, https://stylem.pl/upload/files/product/access-point-zyxel-wbe660s-eu0101f-344547.pdf.

**WBE530.** BE11000 tri-radio 2×2 (2.4/5/6 GHz: 688 Mb/s + 4,324 Mb/s + 5,764 Mb/s), Qualcomm quad-core, 2× 2.5GbE, PoE+ (24 W typical), NebulaFlex Pro, bundled 1-year NCC Pro Pack, classroom-targeted [independent]: https://zyxelguard.com/WBE530.asp. Street prices Sep 2026 [independent]: $239.99 (ZyxelGuard, list $289.99); SHI US $282.00 (MSRP $239.99); UK £283.97 (Laptops Direct) / £303.30 (Inception); EU €299.00 (axitech.be) / €311.32 ex VAT (esus-it); Czechia CZK 9,139 (Alza). Sources: https://www.shi.com/product/48606975/Zyxel-NebulaFlex-Pro-WBE530, https://www.laptopsdirect.co.uk/zyxel-wbe530-tri-band-wifi-7-access-point-10.7gbps-be-wbe530-eu0101f/version.asp, https://www.alza.cz/EN/zyxel-wbe530-d12766390.htm.

**NWA series — 2026 launches.**
- **March 17, 2026: four Wi-Fi 7 APs declared OpenWiFi-ready for ISPs** (Anaheim, GlobeNewswire) [vendor-reported]: NWA50BE BE5100 $79.99 street; **NWA55BE** new BE5100 **outdoor** $149.99; NWA130BE tri-band $139.99 street (listed as "BE1100" in the release — likely a typo for BE11000, flagged); NWA210BE BandFlex BE12300 $169.99. Source: https://business.kanerepublican.com/kanerepublican/article/gnwcq-2026-3-17-zyxel-networks-expands-openwifi-ready-portfolio-with-four-wi-fi-7-access-points-for-isps
- **May 2026: six affordable Wi-Fi 7 APs for SMEs** — NWA90BE PRO, NWA50BE PRO, NWA90BE, NWA50BE, NWA55BE, NWA30BE (MSRPs $79.99–$99.99), Smart Mesh MLO with dual-radio mesh backhaul, SSID VLAN tagging, portfolio now 13 Wi-Fi 7 models [secondary]: https://itbrief.com.au/story/zyxel-networks-unveils-six-new-wifi-7-access-points-for-smes; background: https://www.pocnetwork.net/technology-news/zyxel-launches-new-affordable-wifi-7-nebulaflex-access-points/
- **September 8, 2026: IAP500BE** — Zyxel's **first industrial-grade Wi-Fi 7 AP**, aimed at manufacturing/warehousing/logistics MSP opportunities [secondary]: https://manageditmag.co.uk/zyxel-networks-targets-industry-4-0-dx-with-first-industrial-grade-wifi-7-access-point/
- **September 13, 2026: outdoor trio for MSPs** — **WBE665S** (weather-resistant Wi-Fi 7 outdoor), NWA55AX PRO and a PTP point-to-point device (up to 5 km), all Nebula-managed [secondary]: https://technologyreseller.uk/zyxel-networks-launches-trio-of-aps-to-help-msps-target-growing-outdoor-market/
- **Caution:** a Technology Reseller page dated September 20, 2026 announces "four new" APs (NWA110BE, NWA210BE, WBE510D, WBE630S) but its footnote says "available in Q1 2025" — the copy appears recycled from an older announcement; treat this item as [unverified] rather than a genuine Sep 2026 launch: https://technologyreseller.uk/zyxel-networks-extends-industry-leading-wifi-7-portfolio-to-give-customers-more-choice-flexibility-and-better-user-experiences/
- **Street prices for IAP500BE, WBE665S, NWA55AX PRO, NWA110BE, NWA55BE: not found** — flagged unknown.

#### B.6. Open uncertainties (Pass B)

1. Official MSRP for TP-Link S5500 line and EU availability; Agile street prices.
2. Nuclias Connect 2026 software version; DAP-E9560/X3060W pricing.
3. Sep 2026 retail street prices for USG FLEX 200H/500H/700H.
4. GS2200/GS2220 2026 firmware — none found, EOL confirmed for GS2200.
5. Prices for Zyxel's Sep 2026 outdoor/industrial Wi-Fi 7 APs.

---

