---
id: etape6-phaseb-smb-networking/09-supplementary-research-pass-addendum-september-22-2026/part-3
title: "Supplementary Research Pass — Addendum (September 22, 2026) (part 3)"
domain: supplementary-research-pass-addendum-september-22-2026
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2023-12-06", "2024-05-01", "2025-01-31", "2026-01", "2026-02-17", "2026-07", "2026-07-02", "2026-08-07", "2026-09", "2026-09-22", "2026-14-07", "2026-19-05", "2026-20-03", "2026-30-10"]
keywords: ["ethernet"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1692, 1735]
section: "Supplementary Research Pass — Addendum (September 22, 2026)"
sha256: 994b275184f42c21dd4ff0e0ff362b715374d6416bc86b02f4c05782a8f30c9e
---

# Supplementary Research Pass — Addendum (September 22, 2026) (part 3)

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

