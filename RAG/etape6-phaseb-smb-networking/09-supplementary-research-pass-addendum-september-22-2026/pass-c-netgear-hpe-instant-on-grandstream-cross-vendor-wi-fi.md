---
id: etape6-phaseb-smb-networking/09-supplementary-research-pass-addendum-september-22-2026/pass-c-netgear-hpe-instant-on-grandstream-cross-vendor-wi-fi
title: "Pass C — Netgear, HPE Instant On, Grandstream, cross-vendor Wi-Fi 7 table, 25G/40G pricing"
domain: supplementary-research-pass-addendum-september-22-2026
role: deep-dive
task: pricing
actors: ["United States"]
dates: ["2023-12-06", "2024-03", "2024-05-01", "2024-08", "2025-01-31", "2026-01", "2026-01-27", "2026-02", "2026-02-17", "2026-03", "2026-05", "2026-06-30", "2026-07", "2026-07-02", "2026-08-07", "2026-09", "2026-09-22", "2026-14-07", "2026-19-05", "2026-20-03", "2026-30-10"]
keywords: ["pricing", "ethernet", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1649, 1737]
section: "Supplementary Research Pass — Addendum (September 22, 2026)"
sha256: 8cac9390c4ccc9ab6fb77353e1132011d392a2fdcf722614743c8db7e379dee4
---

# Pass C — Netgear, HPE Instant On, Grandstream, cross-vendor Wi-Fi 7 table, 25G/40G pricing

### Pass C — Netgear, HPE Instant On, Grandstream, cross-vendor Wi-Fi 7 table, 25G/40G pricing

Evidence caveat: "September 2026" prices are retailer page snapshots whose crawl/index dates fall within roughly September 2026, or dated listings explicitly noted. Where a listing's crawl date is older, the crawl recency is stated. Provenance tags: [official] (vendor page/release notes/datasheet), [vendor-reported] (vendor PR statement), [independent] (hands-on review/lab), [secondary] (press, distributor, or retailer page), [unverified] (could not be confirmed or negative finding). Cutoff: September 22, 2026.

#### C.1. NETGEAR

**C.1.1. 2026 AV switch launches (ISE 2026).**
- NETGEAR announced two new M4350 AV switches at ISE 2026 in a release dated January 27, 2026: the M4350-16M4V and the M4350-16C, with March 2026 availability [official, via syndicated release] [secondary]: https://www.varindia.com/news/netgear-introduces-next-gen-m4350-switches-with-rugged-design-and-offline-av-ready-network-configuration-at-ise-2026
- The M4350-16M4V uses industry-standard Neutrik locking connectors (etherCON for network, opticalCON for fiber, powerCON for power), delivers 16 × 2.5G PoE++ ports (up to 1,130 W total) and 4 × 25G SFP28 uplinks, and has four modular interface-card slots for uplink customization [official] [secondary].
- The M4350-16C provides 16 ports of 100G connectivity aimed at aggregation/core layers in large AV-over-IP deployments [official] [secondary].
- The two models brought the M4350 portfolio to 18 models total (1G to 100G), with TAA-compliant versions available for government projects [official] [secondary].
- Note on SKU naming: retailer listings observed in 2026 also show an "M4350-16V4V/16V4C (VSM4320C)" variant described as 16 × SFP28 25G + 4 × QSFP28 100G, which is a different port configuration from the newly announced 16 × 100G "M4350-16C"; the two SKUs should not be conflated [secondary] (Videoguys: https://videoguys.com/search?q=netgear+m4350; Thomas-Krenn: https://www.thomas-krenn.com/en/products/infrastructure/ethernet-switches/netgear-fully-managed-m4350).
- No evidence of a 2026 M4250 refresh, M4500 update, MS510TXUP successor, or new XS/SX managed-switch launch was located; this is a negative finding, not a claim that none exist [unverified].

**C.1.2. Management: Engage Controller 2.4 and Insight 10.0.**
- NETGEAR Engage Controller 2.4 (available February 2026) added offline provisioning, virtual switches/Wi-Fi 7 APs, pre-site testing, configuration export, and reusable templates [official] [secondary].
- Insight 10.0 launched on June 30, 2026 [vendor-reported] (dated release mirror: https://stockhouse.com/news/press-releases/2026/06/30/netgear-introduces-the-next-generation-of-insight-advancing-the-future-of-ai). [secondary]
- The official Insight 10.0 data sheet describes a redesigned UI/UX, a streamlined single-tier subscription, map/topology reporting, AI-defined networking in beta, Wi-Fi heatmaps, deeper device insights, improved client tracking, hierarchical site structure, onboarding, and role-based access control [official] (https://assets.netgear.com/asset/9e3910f8-5cf2-4920-bb90-a21c4ee40121/NG_ENT_Data-Sheet_Insight10.pdf).
- Independent coverage summarizes Insight 10.0 as AI-driven network management aimed at SMEs and MSPs [independent] (https://www.networkworld.com/article/4191261/netgear-brings-ai-driven-network-management-to-smes-and-msps.html).

**C.1.3. September 2026 street-price snapshots (NETGEAR).**

| Product | Observed price | Retailer / date evidence |
|---|---|---|
| M4350-16V4C (VSM4320C, 16×25G SFP28 + 4×100G QSFP28) | US $10,799.99 TAA / $8,999.99 non-TAA (sale) [secondary] | Videoguys, page crawled ~Sep 8, 2026 (https://videoguys.com/collections/netgear) |
| M4350-16V4C | €7,390.00 [secondary] | Thomas-Krenn, page crawled ~Sep 18, 2026 (https://www.thomas-krenn.com/en/products/infrastructure/ethernet-switches/netgear-fully-managed-m4350) |
| M4350-16V4C | £9,690.00, in stock [secondary] | DJKit UK, page crawled ~Jul 2026 (https://www.djkit.com/products/netgear-av-m4350-16v4c-managed-switch-16xsfp28-25g-and-4xqsfp28-100g-vsm4320c-100nes) |
| M4350-16M4V (MSM4320, new 2026 model) | AUD $14,799.00 (was $19,799) [secondary] | Broadcast Bruce AU, page updated ~Feb 2026 (https://www.broadcastbruce.com/product/netgear-av-line-m4350-msm4320/) |
| M4350-16M4V | CHF 7,179.00 [secondary] | Alltron Stamm (CH), page crawled ~Jun 2026 |
| MS510TXUP (8-port multigig/10G PoE++, 295 W, 2×SFP+) | US $691.33 in stock [secondary] | Barcodes Inc, page crawled ~Apr 2026 (https://www.barcodesinc.com/netgear/part-ms510txup-100nas.htm) |
| MS510TXUP | US $725.99 new, 4 available [secondary] | Compsource, page crawled ~May 2026 (https://www.compsource.com/buy/MS510TXUP100NAS/Netgear-306/MS510TXUP-Ethernet-Switch-MS510TXUP100NAS/) |
| MS510TXUP | £570.19 ex VAT (£684.23 inc), in stock next day [secondary] | Comms Express UK, page crawled ~Jun 2026 (https://www.comms-express.com/products/netgear-ms510txup-100eus-8-port-multi-gigabit-10g-ethernet-poe-smart-managed-switch/) |
| WBE710 (BE9400 tri-band Wi-Fi 7, 9.4 Gbps, 2.5G PoE+) | US $339.99, sold by Newegg [secondary] | Newegg, current-era listing (https://www.newegg.com/netgear-wbe710/p/N82E16833222347?Item=N82E16833222347) |
| WBE710 | listed at Mwave AU among Wi-Fi 7 APs (BE9400, 2.5GbE, PoE) [secondary] | Mwave, page crawled ~May 2026 (https://www.mwave.com.au/wireless-networking/access-points-bridges-repeaters/netgear) |

**C.1.4. Wi-Fi 7 AP family correction ("WAX655").**
- No NETGEAR product named "WAX655" was found in any search result; the example in the research brief appears to be erroneous [unverified].
- NETGEAR's business Wi-Fi 7 AP family uses **WBE** identifiers: WBE750 (tri-band, launched March 2024 at US $699.99) and WBE710 (tri-band, launched August 2024 at US $349.99) [official] (launch releases: https://www.nasdaq.com/press-release/netgear-unveils-the-ultimate-tri-band-wifi-7-access-point-wbe750-for-heavily and https://www.nasdaq.com/press-release/netgear-expands-its-wifi-7-solution-portfolio-businesses-new-wifi-7-access-point-2024).
- WBE700 is a dual-band Wi-Fi 7 (BE5000) AP with 5 Gbps aggregate, 2×2 on 2.4/5 GHz, one 2.5GbE PoE+ port, MLO/4K-QAM, 150 idle/75 active clients, Insight plus standalone local GUI, and a one-year Insight subscription included [official] (https://www.downloads.netgear.com/files/GDC/WBE700/WBE700_DS.pdf); its exact launch date was not verified, so it is **not** claimed as a 2026 launch [unverified].
- WAX630E is a Wi-Fi 6E (AXE7800) AP, not Wi-Fi 7 [secondary].

#### C.2. HPE Networking Instant On

**C.2.1. Hardware status 2026 — no Wi-Fi 7, no new switch family found.**
- No credible Instant On Wi-Fi 7 AP was located as of September 22, 2026; the AP lineup still tops out at Wi-Fi 6/6E (AP21/AP22/AP22D/AP25/AP27/AP32) [unverified].
- The flagship AP32 (US SKU S1T22A) is Wi-Fi 6E (2.4/5/6 GHz, max 2.4 Gbps on 6 GHz, 1.2 Gbps on 5 GHz, 574 Mbps on 2.4 GHz, 1×2.5GbE, PoE, 2-year warranty) [secondary] (NEBuMAX, page crawled ~Sep 19, 2026: https://nebumax.com/Products/overview/M025106084).
- The brand changed from "Aruba Instant On" to "HPE Networking Instant On" on May 1, 2024, not in 2026 [official] (https://community.hpe.com/t5/networking/evolving-excellence-aruba-instant-on-is-now-hpe-networking/ba-p/7213752?nobounce).
- The 2026 lineup guide still lists the 1430/1830/1930/1960 switch families and states cloud management carries no subscription fee [secondary] (https://www.comms-express.com/blog/hpe-aruba-meet-the-aruba-instant-on-switches/).
- HPE's January 2026 retail announcements (NRF 2026) covered the enterprise CX 6000 switch, not Instant On [secondary] (https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html).
- New enterprise Wi-Fi 7 APs (slimmed-down 740/720 models) reported in late 2025 belong to the Aruba enterprise line, not Instant On, and must not be conflated [secondary] (https://www.sdxcentral.com/news/hpes-switch-update-taunts-ciscos-tardiness-vmware-value/).

**C.2.2. Software: Instant On 3.4.0.**
- The official Instant On 3.4.0 user guide (mobile-app version, PDF observed in 2026) documents: per-protocol ALG toggles (SIP, H.323, RTSP, FTP, TFTP), bulk PoE power-cycle for wired clients, IP-reservation enhancements, firewall policies (Allow/Block/Restrict) plus the AI policy assistant extended to switch-only sites, software-update retry, service-outage email notifications, and post-onboarding device recovery via the local WebUI without factory reset [official] (https://arubanetworking.hpe.com/techdocs/instanton/3.4.0/HPE-Networking-Instant-On-3.4.0-User-Guide-Mobile-App-Version.pdf). [unverified] The exact 3.4.0 release date was not verified from this document.

**C.2.3. September 2026 price snapshots.**

| Product | Observed price | Retailer / date evidence |
|---|---|---|
| Instant On 1960 12-port multigig (S0F35A: 8×PoE+, 4×2.5G PoE++, 2×SFP+, 2×10G RJ45, 480 W max) | NZD 2,876.66 [secondary] | PB Tech Pacific, snapshot Sep 22, 2026 (https://www.pbtech.com/pacific/product/SWHAUB19605/Aruba-Instant-On-1960-S0F35A-12-Port-Multi-gigabit) |
| Instant On AP32 (US, S1T22A) | US $187.32 new (MSRP $397.00); promo $122.38 through 10/30/2026 [secondary] | NEBuMAX, page crawled ~Sep 19, 2026 (https://nebumax.com/Products/overview/M025106084) |
| Office bundle: 1930 8-port PoE (JL681A-ACC, 124 W) + 2×AP22D (Wi-Fi 6) | £569.96 inc VAT (£474.97 ex), out of stock [secondary] | NetXL UK, page crawled ~Sep 22, 2026 (https://www.netxl.com/poe-network-switches/hpe-networking-office-bundle/) |

**C.2.4. Positioning vs UniFi/Omada.**
- Instant On's documented differentiators are subscription-free cloud management and SMB-targeted simplicity (app-based onboarding, 1430→1960 switch range) [secondary].
- Its AP portfolio lags a generation at the top end: no Wi-Fi 7 Instant On AP exists, while TP-Link, Zyxel, Grandstream, and Cambium have all shipped Wi-Fi 7 APs (see §C.4); UniFi/Omada positioning details are in the main document and are not repeated here. [unverified] Whether HPE plans an Instant On Wi-Fi 7 AP was not determinable from public sources.

#### C.3. Grandstream

**C.3.1. Wi-Fi 7 confirmation and 2026 additions.**
- Grandstream's first Wi-Fi 7 AP launch was the GWN7670 (dual-band BE3600, 3.6 Gbps wireless, 5 Gbps aggregate wired, 256 clients), announced January 31, 2025 — **not** 2026 [official] (https://blog.grandstream.com/press-releases/archive/2025/01).
- **Genuinely new 2026 Wi-Fi 7 models located:**
  - **GWN7672L** — official firmware release notes dated 07/14/2026 state "This is the initial release for GWN7672L" (firmware 1.0.27.5) [official] (https://firmware.grandstream.com/Release_Note_GWN7672L_1.0.27.5.pdf). A New Zealand master distributor lists GWN7672L at $272.00 +GST with 50+ in stock [secondary] (https://www.gowifi.co.nz/manufacturer/grandstreamnetworks.html?Itemid=878) — [unverified] note the distributor page crawl (~Feb 2026) predates the July 2026 firmware initial-release date, an unresolved discrepancy.
  - **GWN7672WM** — wall-mounted tri-band Wi-Fi 7 (BE11000, 11 Gbps wireless, 5 Gbps wired, 384 clients, 175 m, MLO/4K-QAM/MRU/preamble puncturing) [official] (https://www.grandstream.com/networking-solutions/indoor-wifi-access-points/product/gwn7672wm); its quick-installation guide is dated July 2, 2026 [secondary] (https://manuals.plus/category/grandstream).
  - **GWN7670WM and GWN7670LR** — distributor reporting says these were added to the Wi-Fi 7 lineup after the GWN7670/GWN7672 [secondary] (https://info.teledynamics.com/blog/grandstream-expands-its-wi-fi-7-ap-lineup-with-the-gwn7670wm-and-7670lr); a retailer lists GWN7670LR (dual-band Wi-Fi 7, long-range, 256 clients, 350 m, 1×2.5G RJ45 + 1×2.5G SFP) with a first-available date of February 17, 2026 [secondary] (https://ru.microless.com/en/product/grandstream-gwn7670lr-dual-band-wi-fi-7-access-point-256-client-devices-350-meters-range-1x-2-5g-ethernet-wan-lan-rj-45-poe-input-1x-2-5g-sfp-interface-white-gwn7670lr/).
  - **GWN7670E** — dual-band Wi-Fi 7 AP present in the 2026 firmware tables (1.0.27.5) and with an installation guide dated August 7, 2026 [official/secondary]; NZ distributor price $171.00 +GST [secondary]; exact launch date not verified [unverified].
  - **GWN7674** — appears in the 2026 official firmware table (1.0.27.6) [official]; launch date and specs not verified [unverified].

**C.3.2. Management: GDMS Networking (renamed from GWN.Cloud).**
- "GWN.Cloud was renamed to 'GDMS Networking'" per the official GDMS release notes [official] (https://firmware.grandstream.com/Release_Note_GDMS.pdf).
- **Dated 2026 update:** GDMS Networking firmware version 1.1.37.6, dated 5/19/2026, adds wireless-bridge configuration and status management, AI floor-plan wall recognition and AI AP placement (beta), RF-planning report export, SSID multi-VLAN, scheduled SSID-password rotation, 802.1X identity display in the client list, client batch configuration, client offline alarms and bandwidth alerts, voucher templates with QR codes, unified static-IP binding, PPSK quantity limits per device model, URL access-log forwarding, clone-network support for WAN/VPN and device groups, a User Administrator role, GS Home account management, wireless-router mesh management, WireGuard auto-peer allowed-IP/DNS settings, and select-model router QoS/WAN-access-control/FXS features [official] (https://firmware.grandstream.com/Release_Note_GWN_Cloud.pdf, crawled ~Jul–Aug 2026).
- Current management architecture per product sources: embedded controller manages up to 50 local APs, GDMS Networking supports unlimited APs, and GWN Manager (on-premises) supports up to 3,000 APs [secondary] (Best4Systems GWN7672 listing: https://www.best4systems.co.uk/grandstream-gwn7672-tri-band-wi-fi-7-access-point.html).

**C.3.3. Switches — no new 2026 model found.**
- Negative finding: no genuinely new 2026 switch SKU beyond the already-documented GWN7816/7816P, GWN7810 series, GWN7821P/GWN7822P, and GWN7830-32 was located; search results surfaced only older SKUs (GWN7813/P, GWN7806P, GWN7832, GCC6011), none with 2026 launch evidence [unverified].
- The GWN7816/7816P launched December 6, 2023 [official] (https://blog.grandstream.com/press-releases/grandstream-releases-48-port-layer-3-network-switches).
- GWN7062M (Wi-Fi 6 router, not a switch and not Wi-Fi 7) had its first firmware release on 03/20/2026 [official] (https://firmware.grandstream.com/Release_Note_GWN7062M_1.0.1.98.pdf).

**C.3.4. September 2026 street-price snapshots (Grandstream Wi-Fi 7).**

