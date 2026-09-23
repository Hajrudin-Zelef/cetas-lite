---
id: etape6-phaseb-smb-networking/01-part-1-ubiquiti-unifi/3-unifi-os-network-application-software-2026
title: "3. UniFi OS / Network Application software 2026"
domain: part-1-ubiquiti-unifi
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2026-05-18", "2026-09-22"]
keywords: ["advisory", "latency", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [70, 131]
section: "Part 1 — Ubiquiti UniFi"
sha256: 98e490e59556020027327ce23ec33b3d1ee643433d553285406701fb0c7a115b
---

# 3. UniFi OS / Network Application software 2026

## 3. UniFi OS / Network Application software 2026

### 3.1 UniFi Network Application 10.x — major 2026 versions [secondary/official via community release notes]
- **10.1.89** (Mar 18, 2026) — security-fix release; also 10.2.97 (RC track). Patch for CVE-2026-22557/22558 (see §7). Full notes: https://community.ui.com/releases/UniFi-Network-Application-10-4-57/92694b29-fd78-4d52-906a-3211136610e2 (10.4.57 link; 10.1.x notes on community.ui.com).
- **10.3.58** (Apr 28, 2026) — maintenance.
- **10.4.57** (May 18, 2026) — **new Routing experience**: centralized Routing view managing Static, OSPF, **BGP**, Connected, and Policy-Based routes + unified FIB table (destinations, next hops, interfaces, metrics, sources); expanded PBR; dashboard WiFi Airtime + Multicast graphs; WAN/ISP side panel with SLA, latency, loss, utilization; improved 5G/cellular visibility (requires U5G firmware 7.4.1+) [official via community release notes, mirrored in GitHub changelog].
- **10.5.67** (Aug 2026) — current in security bulletins [secondary].
- **10.6.106** (Sep 10, 2026) — latest tracked; Home Assistant fork packaged same day [secondary].
- **Zone-based firewall** is the default policy system (since Network 8.x); legacy rules coexist with zone policies in 9.x/10.x — a new zone defaults to *block* vs Internal both directions [secondary/community].
- Feature inventory 2026: VLAN management, policy-based routing, WireGuard VPN server/peers (incl. Terraform `unifi_wireguard_peer` support), Teleport VPN, Site Magic (multi-site VPN mesh), WLAN PPSK (private pre-shared keys), IPv6 (static/RA/PD/DHCPv6), OSPF + BGP in routing view, 802.1X/MAB RADIUS VLAN assignment [secondary/community].
- Sources: https://github.com/zglate/addon-unifi/blob/HEAD/unifi/CHANGELOG.md ; https://github.com/ubiquiti-community/terraform-provider-unifi/blob/HEAD/CHANGELOG.md

### 3.2 UniFi OS 5.x
- Current track in 2026: **UniFi OS 5.1.x** — fixes shipped in 5.1.19 (Jul 2026 CVE batch) and 5.1.37 (Aug 2026 bulletin) [secondary]. UniFi OS Server 5.1.37 named in Aug 2026 advisory.
- UniFi Express (UX) firmware **4.0.13** (bundles Network 9.0.114/9.0.118) [secondary].
- Note: CVE advisory lists affected families incl. UDM/UDM-Pro/UDM-SE/UDM-Pro-Max/UDM-Beast/EFG/UDW/UDR/UDR7/UDR-5G/Express 7/Cloud Keys/UNVR/ENVR/UCG/EF-Core/UNAS — confirming the 2026 hardware matrix [secondary].

---

## 4. Wi-Fi 7 APs (U7) — lineup & pricing, Sep 2026

U7 is the current Wi-Fi 7 generation (tri-band with 6 GHz, MLO) [secondary]. Prices below = US list/typical; most corroborated by Aug–Sep 2026 sources:
| Model | Price (USD) | Notes |
|---|---|---|
| U7 Lite | $99 | 4-stream, 2.5GbE, budget Wi-Fi 7 (launched Mar 2025 per 9to5Mac) |
| U7 In-Wall (U7-IW) | $149 | 4-stream, 2.5GbE, PoE output (launched Mar 2025) |
| U7 Pro | $189 | tri-band 6 GHz, 6-stream "sweet spot" |
| U7 Pro XG | $199 | tri-band + **10GbE uplink** |
| U7 Pro Wall | $199 | in-wall, 2.5GbE passthrough |
| U7 Pro Outdoor | ~$299 | IP67 |
| U7 Pro Max | $279 | tri-band, 8-stream, dedicated scanning radio |
| U7 Mesh / U7 Outdoor / U7 Long-Range / U7 Extender | ~$179–199 | U6 Mesh ($179) still sold |
| **U7 Pro XGS** | unconfirmed | Named in Jul 2026 patent complaint + UAE retailer (AED 1,275) — specs/price TBD [unverified] |
| **U7 Pro XG Wall** | unconfirmed | Named in patent complaint [unverified] |
| E7 / E7 Campus / E7 Audience | unconfirmed | Enterprise Wi-Fi 7 family named in complaint; E7 Campus listed AED 3,333 (5-pack?) in UAE [unverified] |

- Sources: https://9to5mac.com/2025/03/03/ubiquiti-new-wi-fi-7-gear-starts-at-99-and-can-supercharge-your-iphone-16/ ; https://ifeeltech.com/blog/ubiquiti-unifi-u6-series-review (Aug 2026 buying guide) ; https://uae.microless.com/product/ubiquiti-unifi-u7-pro-xg-wifi-7-tri-band-access-point-with-10-gbe-uplink-2x2-mu-mimo-per-band-5-8-gbps-6-ghz-speed-300-clients-poe-5-pack-white-u7-pro-xg-5/ ; https://www.comx-computers.co.za/UB-UAP-U7-IW-Ubiquiti-UniFi-WiFi-7-In-Wall-Access-Point-Buy-p-323240.php (U7-IW R3,365, price valid 2026-09-22)
- U6 line (U6 Lite/Pro/Mesh/In-Wall/LR/Enterprise) is now "fleet continuity" — mature firmware, some models delisted from ui.com, U6+ at $129 [secondary].

---

## 5. Street prices, dated Sep 2026 (key SKUs)

| SKU | Price (USD) | Dated | Source |
|---|---|---|---|
| UDM-Beast | $1,499 | May–Sep 2026 | vendor-reported |
| EFG / UXG-Enterprise | $1,999 | Sep 2026 | retailer (ispsupplies/centralcomputer) |
| UDM-Pro-Max | $599 | list (not re-verified) | official list |
| UDM-SE | $499 | list (not re-verified) | official list |
| UDM-Pro | $379 | list (not re-verified) | official list |
| USW-Pro-XG-8-PoE | $499 | Sep 2026 | STH review |
| USW-Pro-XG-10-PoE | $699 | Sep 2026 | STH review |
| USW-Pro-XG-Aggregation (32×SFP28) | ~$2,999? | uncertain — not confirmed; ZAR 59,975 retail | [uncertainty] |
| Enterprise Campus Switch Core 32×100G | $5,393 | Sep 2026 | Notebookcheck (incl. $394 surcharge) |
| Enterprise Campus Switch Core 25G | $4,314 | Sep 2026 | Notebookcheck |
| U7 Pro / Pro XG / Pro Max | $189 / $199 / $279 | Aug–Sep 2026 | secondary |
| U7 Lite / In-Wall | $99 / $149 | Sep 2026 | secondary |

Non-US datapoints: Scoop (South Africa) Sep 2026 — USW-Pro-XG-8-PoE ZAR 12,295; 10-PoE ZAR 16,995; Pro-XG-24 ZAR 26,995; 24-PoE ZAR 43,695; Aggregation ZAR 59,975. ComX (ZA): U7-IW R3,365 (price valid 2026-09-22). Microless (UAE, Sep 2026): U7-Pro-XG 5-pack AED 4,544.98.

---

