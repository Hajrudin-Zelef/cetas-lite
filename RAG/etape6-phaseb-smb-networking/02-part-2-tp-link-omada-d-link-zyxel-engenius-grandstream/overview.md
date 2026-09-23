---
id: etape6-phaseb-smb-networking/02-part-2-tp-link-omada-d-link-zyxel-engenius-grandstream/overview
title: "Part 2 — TP-Link / Omada, D-Link, Zyxel, Engenius, Grandstream"
domain: part-2-tp-link-omada-d-link-zyxel-engenius-grandstream
role: deep-dive
task: reference
actors: ["EU", "United States"]
dates: ["2026-03", "2026-04-29", "2026-05", "2026-07-17", "2026-07-30", "2026-09", "2026-09-04", "2026-09-22"]
keywords: ["acquisition", "cost", "license", "pricing", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [188, 255]
section: "Part 2 — TP-Link / Omada, D-Link, Zyxel, Engenius, Grandstream"
sha256: e15ef1a9fc3d94b3eb3f929f436aeeb17b52750bc9b5ac5cb9e35f7123bc2387
---

# Part 2 — TP-Link / Omada, D-Link, Zyxel, Engenius, Grandstream

# SMB/Prosumer Networking — TP-Link/Omada + D-Link + Zyxel + EnGenius + Grandstream
Research date: September 22, 2026. All prices "street" = retailer listings observed September 2026; MSRP = vendor price lists. Provenance tags: [official], [vendor-reported], [independent], [secondary], [unverified].

---

## 1. TP-Link / Omada

### 1.1 Omada SDN Controller — 2026 versions & features

- **Current stable release: Omada Network Application v6.3.0.45** (Windows/Linux), published 2026-09-04 [official]. Previous stable v6.2.14.11 (2026-07-17), v6.2.10.17 (2026-04-29); Omada Discovery Utility 5.4.0 (2026-07-30) [official].
  Source: https://support.omadanetworks.com/en/product/omada-software-controller/?resourceType=download
- **v6.3 highlights** (community release notes, ~2 weeks before Sep 22, 2026) [vendor-reported]: Intelligent Anomaly Detection (previously Omada Pro exclusive) now in standard Omada; wireless AP Grouping (one SSID across multiple AP groups); Wireless IDS/IPS (rogue-AP detection, auto-disconnect/lock); Rapid Device Replacement for APs and switches (enter new device key to adopt); OUI-Based VLANs now on EAPs (was Omada Pro exclusive) and on "Agile" switches. Available for OC200/OC220/OC300/OC400 hardware controllers and software controller [vendor-reported].
  Source: https://community.tp-link.com/en/business/forum/topic/871456
- **v6.1.x pre-release** (closed Jan 21, 2026) [vendor-reported]: site-level health dashboards (24h site-health trend, radar score), device health timelines on AP/switch/gateway detail pages, health reports (Site/Wi-Fi/Switch/AP health trends), Client Health toggle (MSP/global view), automatic RF optimization via spectrum scanning + RRM (Adaptive Mode), RadSec (RADIUS over TLS) for SSID/MAC/802.1X on APs and switches (requires firmware; not yet for Built-in RADIUS/Portal/SSL-VPN), EAP 725-Outdoor support.
  Source: https://community.tp-link.com/en/business/forum/topic/850962
- **v6.2.10** (Apr 2026) [vendor-reported]: Wi-Fi Calling support, AP load balancing, SSID band steering, WLAN optimization schedule, AFC acquisition status display, CCI/AP-density charts, SD-WAN access-control logic update.
  Source: https://static.tp-link.com/upload/software/2026/202604/20260429/software%20controller%20v6.2.10%20release%20note%20win&linux.pdf
- **v6.0 pre-release** (Nov 2025, feeds 2026 devices): added gateway support for ER701-5G-Outdoor, ER706W-4G(EU)2.20 / ER706WP-4G, ER706W(EU/US)1.30, ER7206(UN)2.30, ER7406(UN)1.20, ER605(UN)2.30, ER707-M2(UN)1.30, ER7412-M2(UN)1.30, ER8411(UN)1.20, DR3220v-4G/DR3650v(-4G) [vendor-reported].
  Source: https://community.tp-link.com/en/business/forum/topic/840402

### 1.2 Omada tiers: free vs paid

- **Omada Software Controller: free, no license fees**, manages up to 1,500 devices, hybrid cloud access [official].
  Source: https://tp-link.com/ca/business-networking/management-platform/omada-software-controller/v4/
- **Hardware controllers** (OC200/OC300/OC400): one-time hardware cost, "no license fee, no monthly fee" per retailer listings [secondary]. OC200 street ≈ $50–65 USD/EU equivalent (2026 listings); precise Sep-2026 street price not confirmed — **flag: unverified**.
- **Cloud-based Omada controller tiers** (Essentials free / Standard): comparison table by independent reviewer Dong Knows (updated ~May 2026) [independent]: Cloud Essentials = free, essential features, no MSP mode, no PPSK, no ACL/IPS/IDS/DPI; Cloud Standard = "Device License Free" per table (wording ambiguous — TP-Link phrasing; likely per-device license, price not disclosed), adds PPSK, full VPN (L2TP/PPTP/SSL/OpenVPN/IPsec/WireGuard), MSP mode, ACL/IPS/IDS/DPI, advanced portal auth (RADIUS/LDAP/external portal). Exact per-device license pricing for Standard **not found — flagged as unknown** [unverified].
  Source: https://dongknows.com/tp-link-omada-class-diy-mesh-review/
- **Omada Pro**: separate enterprise line (S5500 switches, Pro controllers); features trickling down into standard Omada (Anomaly Detection, OUI VLANs in v6.3). Pro pricing not publicly listed — **unknown** [unverified].
- **Agile Switches**: new 2026 portfolio family referenced in v6.3 notes (firmware must be Controller-6.3-compatible for OUI VLANs) — details beyond this not found; **flag: thin sourcing** [unverified].

### 1.3 Switches — 2026 portfolio (Omada + JetStream + SX)

From TP-Link's 2026 ISC West flyer (March 2026) [official]: current access lines include SG2005P-PD (PoE++ powered), SG2008P, SG2428P, SG3428XMPP / SG3452XMPP (8× PoE++ up to 90W/port, 4× 10G SFP+, 500W/750W budgets), **SG2210XMP-M2** (8× 2.5G PoE+, 2× 10G SFP+, 160W), **SG3218XP-M2 (V2)** (16× 2.5G with 4× PoE++ up to 90W + 12× PoE+, 4× 10G SFP+, 240W), **SG3428XPP-M2 (V2)** (24× 2.5G PoE++, 8× 10G SFP+, 770W), **SX3832MPP** (24× 10G RJ45 PoE++ up to 90W/port, 8× 10G SFP+, 770W, 640 Gbps, 240 Mpps), **SX3032F** (32× 10G SFP+ aggregation) [official].
Source: https://static.tp-link.com/upload/case-study/2026/202603/20260312/2026_ISC%20WEST_Flyer_Omada%20Switch%20Portfolio.pdf

Notable models:
- **SX3206HPP** (6-port 10G L2+ managed PoE++, 4× 10G RJ45 PoE++ + 2× 10G SFP+, 200W budget): NZD 1,129.57 ex GST at PB Tech, listing created 2026-09-22 [independent].
  Source: https://www.pbtech.co.nz/product/SWHTPL103206/TP-Link-Omada-TL-SX3206HPP-6-Port-10GE-L2-Managed
- **SX3832MPP** (see specs above): retailer-confirmed Omada SDN integration, ZTP, per-port 90W [secondary].
- **SG3218XP-M2** (16× 2.5G, 8× PoE+, 2× 10G SFP+, 240W budget, 120 Gbps, L2+): street Sep 2026 — €449 incl VAT (Greece), €542.20 incl VAT (Belgium), AUD 715.39 (eyo.com.au), AUD 774.60 ex GST special / 1,045.60 reg (dataworld.com.au), AED 2,303 Dubai [independent].
  Sources: https://www.xpatit.gr/en/wired-networks/switch/managed/4661-tp-link-sg3218xp-m2-omada-16-port-2-5g-and-2-port-10ge-sfp-l2-managed-switch-with-8-port-poe ; https://www.eyo.com.au/648414_tp-link-omada-16-port-2-5g-and-2-port-10ge-sfp-sg3218xp-m2.html
- **TL-SX3008F** (8× 10G SFP+, L2+, 160 Gbps, fanless): street Sep 2026 — $194.97 USD (AlwaysInTouch), AUD 429 (Computer Alliance), €290.92 incl VAT (Netherlands), ₱14,900 PH, ₹22,500 India, AED 994 Dubai [independent].
  Sources: https://www.alwaysintouch.com/tp-link-tl-sx3008f-8-port-10g-sfp-enterprise-level-switch-l2-smart-managed-omada-sdn-integrated-ipv6-static-routing-l2-l3-l4-qos-igmp-lag-5-year-manufacturer-warranty/ ; https://www.computeralliance.com.au/8-port-tp-link-tl-sx3008f-jetstream-10ge-sfpplus-l2plus-managed-switch
- **TL-SG3210XHP-M2** (8× 2.5G PoE+, 2× 10G SFP+): AUD 774.60 ex GST special / AUD 1,045.60 regular (dataworld.com.au, 4cabling.com.au, in stock Sep 2026) [independent].
  Source: https://dataworld.com.au/product/tp-link-sg3218xp-m2-omada-16-port-2-5g-and-2-port-10ge-sfp-l2-managed-switch-with-8-port-poe/
- **Omada Pro S5500-24MPP4XF** (enterprise line: 8× 2.5G PoE++ + 16× 2.5G PoE+, 4× 10G SFP+, 500W budget, 200 Gbps): listed by US/Canada/Australia resellers; no current street price captured — **flag: price unconfirmed** [unverified].
  Sources: https://www.ziestech.com/products/tp-link-omada-pro-24-port-poe-2-5g-l2-managed-switch-with-4-sfp-slots-s5500-24mpp4xf

### 1.4 Gateways / routers — 2026

- **ER7412-M2** (2× 2.5G RJ45 WAN/LAN, 8× GbE RJ45, 2× GbE SFP, USB 3.0, up to 11 WAN, 1M concurrent sessions): MSRP CAD 329.99 (TP-Link Canada price list Feb 2026) [official]; street Sep 2026 — $234 USD (Tech-America), $251.75 (multilink.us), AUD 315.80 via BIG W [independent].
- **ER707-M2** (2× 2.5G RJ45, 4× GbE RJ45, 1× GbE SFP, USB 2.0, 6 WAN max, 500k sessions, WireGuard/OpenVPN/SSL/IPsec): MSRP CAD 199.99 [official]; street — $144.99 (directmacro.com), $251.75 (multilink.us), ₱11,990 PH, R2,600 ZA [independent].
- **ER8411** (10G: 1× 10G SFP+ WAN + 1× 10G SFP+ WAN/LAN, 8× GbE RJ45, 1.5M sessions, dual PSU): MSRP CAD 599.99 [official].
- **ER7406** (gigabit rack/desktop VPN router): MSRP CAD 179.99 [official].
- **ER706W-4G / ER706WP-4G** (Wi-Fi 6 AX3000 + Cat6 LTE, dual nano-SIM): 2026 hardware revisions supported in controller v6.0; ER706W street ≈ $130.90 (multilink.us, 2026 crawl) [independent].
- **ER7212PC** (3-in-1 router + PoE+ switch + Omada controller): $251.75 (multilink.us) [independent].
  Sources: https://micro-informa.ca/files/Omada_MSRP_2026.pdf ; https://www.tech-america.com/item/tp-link-omada-multi-gigabit-vpn-router/er7412-m2 ; https://multilink.us/tp-link-er707-m2-omada-multi-gigabit-vpn-router/

### 1.5 Market position

- TP-Link remains the prosumer/SMB price leader with a free controller (no per-device licensing) — a key differentiator vs. Zyxel Nebula paid packs and Ubiquiti-adjacent subscription models [independent assessment].
- Enterprise push: Omada Pro line (S5500 switches, Pro controllers) and 2026 feature convergence (Anomaly Detection, OUI VLANs moving from Pro to standard Omada) show TP-Link climbing into mid-enterprise, while keeping a free SDN controller to undercut Cisco/Meraki-style licensing [independent assessment].
- **Flag:** no verifiable 2026 market-share figures for TP-Link in managed SMB switching found; do not cite any.

---

