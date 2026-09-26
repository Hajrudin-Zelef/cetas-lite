---
id: etape6-phaseb-smb-networking/07-supplementary-complementary-research-pass-round-3-september-/part-5
title: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026 (part 5)"
domain: supplementary-complementary-research-pass-round-3-september-
role: deep-dive
task: reference
actors: ["AWS", "Google", "Microsoft", "Qualcomm", "United States"]
dates: ["2023-03", "2023-10", "2025-04", "2025-07", "2026-01", "2026-01-21", "2026-03-05", "2026-04-09", "2026-04-14", "2026-07-02"]
keywords: ["advisory", "license"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1350, 1382]
section: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026"
sha256: 6d7fb0704df295e08db3dae2c848ab17faebdbe03969a8d9ef803bb86406be15
---

# Supplementary / Complementary Research Pass — Round 3 — September 22, 2026 (part 5)

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

