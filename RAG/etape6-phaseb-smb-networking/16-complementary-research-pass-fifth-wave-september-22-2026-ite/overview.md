---
id: etape6-phaseb-smb-networking/16-complementary-research-pass-fifth-wave-september-22-2026-ite/overview
title: "Complementary Research Pass — Fifth Wave (September 22, 2026), Items S25–S33"
domain: complementary-research-pass-fifth-wave-september-22-2026-ite
role: deep-dive
task: reference
actors: ["EU", "United States"]
dates: ["2026-05", "2026-09", "2026-09-22"]
keywords: ["research", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2384, 2450]
section: "Complementary Research Pass — Fifth Wave (September 22, 2026), Items S25–S33"
sha256: f23774a8cb5fb105af7366438f1a8ae4462ae409de5989033e14f7376307a493
---

# Complementary Research Pass — Fifth Wave (September 22, 2026), Items S25–S33

**Purpose:** Resolve retained gaps and add genuinely new 2026 developments not covered in earlier passes.
**Method:** Public web research. Tags: `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`.
**Prior wave:** S1–S24, R1–R13, T1–T7 and earlier append sections; this wave adds S25–S33.

---

## S25. UniFi U7 Pro XGS / U7 Pro XG / E7 family — launch CONFIRMED (resolves retained gap)

Earlier passes (S12, T1, S20) left the U7 Pro XGS launch status unconfirmed. This is now resolved.

### U7 Pro XG / U7 Pro XGS launch — May 2026 `[independent]`
- Dong Knows Tech reports Ubiquiti officially announced availability of its next-gen flagship UniFi APs, the **U7 Pro XG ($199)** and **U7 Pro XGS ($299)** — described as fan-less alternatives to the earlier fan-cooled U7 Pro / U7 Pro Max, for users who "can't afford the fan-less UniFi E7". In-depth review of the U7 Pro XGS published separately. [independent]
  - https://dongknows.com/ubiquiti-u7-pro-xg-vs-u7-pro-xgs-review/
- Reseller availability confirms the dates: Microless lists the single-pack **U7-Pro-XGS** as "date first available: 09 May, 2026" and the 3-pack (U7-Pro-XGS-3 / -B-3) as first available 13 June, 2026, with stock in hand as of September 2026. [secondary]
  - https://sc.microless.com/product/ubiquiti-u7-pro-xgs-access-point-wi-fi-7-with-6ghz-support-8-data-streams-powered-by-poe-hotspot-2-0-zero-wait-dfs-internal-antenna-type-1x-10gbe-poe-rj45-port-interface-white-u7-pro-xgs/

### U7 Pro XGS specifications (reseller + reviewer data) `[secondary]`/`[independent]`
- Ceiling-mounted 8-stream Wi-Fi 7 AP with **dedicated spectral scanning radio** and 10/5/2.5/1 GbE support.
- 6 GHz support, 8 data streams, PoE++ powered, Hotspot 2.0, Zero-Wait DFS, internal antennas.
- Coverage up to 160 m²; more than 500 connected devices.
- Requires **UniFi Network 8.2.93 or later** for management.
- 1× 10 GbE PoE+ RJ45 port.
- Regional pricing snapshot: **SCR 4,742.16** (Seychelles rupees, reseller list) — regional, not comparable to US MSRP.

### E7 / E7 Campus confirmed shipping `[secondary]`
- The same reseller's UniFi Wi-Fi 7 model selector lists both **E7 (SCR 7,730.06)** and **E7 Campus (SCR 12,574.08)** as available SKUs — the first direct commercial confirmation found of the E7 Campus variant being sold. [secondary]
- Dong Knows confirms the **E7 is the fan-less flagship**, positioned above the U7 Pro XG/XGS. [independent]

### Gap status
- **RESOLVED:** U7 Pro XGS/XG launched and shipping since May 2026. E7 and E7 Campus confirmed as shipping SKUs.
- **Retained:** Official Ubiquiti store list prices for the XG/XGS and E7 Campus (only reviewer-reported $199/$299 and regional reseller prices found); E7 Campus hardware specs (indoor/outdoor ruggedization details).

---

## S26. UniFi Cloud Gateway Max (UCG-Max) — 2026 street pricing and availability

UniFi's most recommended SMB router/gateway of 2026 (S2, S20, R8) is widely stocked. Prices below are as listed by resellers; note all are regional and NOT directly comparable across currencies/VAT regimes.

### Key specifications (multiple resellers) `[secondary]`
- Model/SKU: **UCG-Max**. Compact 2.5G Cloud Gateway.
- CPU: quad-core ARM Cortex-A53 at 1.5 GHz; **4 GB RAM** (per rlan.bg).
- **1.5 Gbps routing with IDS/IPS**; multi-WAN load balancing (1× WAN + 4× LAN, all 1/2.5 GbE; one port configurable LAN/WAN).
- Manages **30+ UniFi devices and 300+ clients**; runs full UniFi application suite (Network, Protect, Access, Talk); UniFi Site Manager ready.
- **512 GB NVMe SSD included** for NVR storage (replaceable with higher-capacity drive).
- 0.96" LCM status display; USB-C powered (5V/5A adapter included); 16 W max power draw.
- VPN: WireGuard, OpenVPN, L2TP, IPSec, site-to-site.

### Price snapshots, September 2026 `[secondary]` (regional, non-comparable)
- **UK:** £199.50 ex VAT, in stock (4gon.co.uk) — https://www.4gon.co.uk/ubiquiti-unifi-cloud-gateway-max-ucgmax-p-10815.html
- **UK:** £222.83 ex VAT / £267.40 inc VAT, 96 in stock (networkwarehouse.co.uk) — https://networkwarehouse.co.uk/products/ubiquiti-ucg-max
- **UK:** £230.52 ex VAT / £276.62 inc VAT, next-day (comms-express.com) — https://www.comms-express.com/products/ubiquiti-ucg-max-uk-unifi-cloud-gateway-max-512gb/
- **Australia:** A$449 (RRP A$569), in stock (shoppingexpress.com.au) — https://www.shoppingexpress.com.au/buy/ubiquiti-ucg-max-2.5gbps-multi-wan-unifi-cloud-gateway-max-router/UCG-Max
- **Australia:** A$484 in stock (devicedeal.com.au); A$553.91 in stock (4cabling.com.au)
- **New Zealand:** NZ$607.44 ex GST / NZ$698.56 inc GST (pbtech.co.nz, multiple branches stocked; page created 22-09-2026) — https://www.pbtech.co.nz/product/NETUBI2209/Ubiquiti-UniFi-UCG-Max-Cloud-Gateway-Max-25GbE-WAN?qr=popular_related_products
- **Czechia:** 7,459 Kč / 6,164 Kč ex VAT, in stock (tsbohemia.cz)
- **Bulgaria:** €325.91 inc VAT (20%), backorder with planned courier delivery 11-11-2026 (rlan.bg)
- **US:** $477.70, out of stock, 3–5 day lead time (icttech.net)

### Notable
- PBTech New Zealand's fresh listing (created 22 Sept 2026) and 96-unit UK stock confirm healthy supply in ANZ/EU at this date; only the US reseller surveyed was out of stock.
- Australian prices cluster A$449–$554 vs NZ NZ$607–698 — Pacific pricing variance of ~15–20% before tax, a useful reminder for RAG procurement notes.
- UCG-Max appears repeatedly in homelab recommendations (S20, R8/R10) as the default UniFi router for 2.5G homes.

---

