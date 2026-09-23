---
id: etape6-phaseb-smb-networking/16-complementary-research-pass-fifth-wave-september-22-2026-ite/s27-netgear-insight-cloud-subscription-2026-pricing-state
title: "S27. NETGEAR Insight cloud subscription — 2026 pricing state"
domain: complementary-research-pass-fifth-wave-september-22-2026-ite
role: deep-dive
task: pricing
actors: ["EU", "Microsoft", "Qualcomm", "United States"]
dates: ["2026-03", "2026-07", "2026-09", "2026-09-22"]
keywords: ["pricing", "cost", "distribution", "license", "licenses", "scout"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2451, 2508]
section: "Complementary Research Pass — Fifth Wave (September 22, 2026), Items S25–S33"
sha256: d4f310f23066be1defe9b480d1f04a2724f04eda7dda5f4f7664733454b88e55
---

# S27. NETGEAR Insight cloud subscription — 2026 pricing state

## S27. NETGEAR Insight cloud subscription — 2026 pricing state

The Insight subscription cost question (important for TCO comparisons in the earlier NETGEAR section S6/S21/R4) had not been captured. Two layers:

### Published list pricing (from NETGEAR flyer, aged) `[official, aged 2021]`
- **Insight Basic:** free for the first 2 Insight devices on the network; $4.99/device/year from the 3rd device.
- **Insight Premium:** $9.99/device/year or $0.99/device/month; applies from the first device; adds Cloud Portal, Smart WiFi management, fast roaming, PoE scheduling.
- **Insight Pro (MSPs/resellers):** all Premium features + multi-tenancy & multi-roles; distribution-only, "contact us" pricing.
  - Source: https://www.netgear.com/images/pdf/Insight-4-Page-Flyer.pdf (NETGEAR 4-page flyer, crawled ~2021; current page text still mirrored at snappernet.co.nz mirror of the same flyer)

### 2026 real-world pricing and changes `[independent]`/`[secondary]`
- MightyGadget blog review (crawled ~March 2026, [independent]) reports:
  - **All NETGEAR Insight devices now come with 1-year free Insight Premium** — then ~£8.95/device/year (UK) on renewal.
  - Credit packs barely discount: 5 devices/1 year ≈ £9.05/device/year; 10 devices/1 year ≈ £8.77/device/year; 10 devices/5 years ≈ £8.39/device/year.
  - **The old Insight Basic free tier "can no longer be signed up for"** — the author could find no sign-up path; NETGEAR support pages still reference it. **[unverified / possible contradiction with official flyer]** — treat the Basic tier's availability as uncertain for 2026 deployments.
  - **Insight Pro ~$22.00/device/year** per the same review.
  - Source: https://mightygadget.com/netgear-insight-wax610-indoor-wax610y-outdoor-wi-fi-6-review/
- Reseller license SKUs (2026 listings, [secondary]):
  - **Insight Pro 1-year, 1 device (NPRSNG1P-10000S):** £13.99 inc VAT / £11.66 ex VAT (netxl.com, out of stock) — https://www.netxl.com/networking-software-licenses/netgear-insight-pro-1-year-subscription-nprsng1p-10000s/
  - **Insight Pro 5-year, 1 device (NPR1SNG5-10000S):** $73.99 (was $101.12, −27%), out of stock (wamatek.com, 2026) — https://wamatek.com/products/netgear-insight-pro-subscription-license-1-managed-device-5-year
  - **Insight Pro 3-year, 1 device (NPR1SNG3-10000S):** €129.03 ex VAT, free shipping 24–72h (redcorp.com, crawled Aug 2026) — https://www.redcorp.com/en/product/communication-networking-warranty/netgear/insight-pro-subscription-1-device-3-year-npr1sng3-10000s/20325458
- Insight-managed switches carry a 5-year Limited Hardware Warranty with Next Business Day replacement (official flyer).

### Gap status
- **CAVEAT:** The strongest primary source (netgear.com flyer) dates to ~2021. 2026 blog evidence suggests the free Basic tier may be discontinued, but NETGEAR's official Insight pages were not re-checked in this wave — mark tier availability as **[unverified]** and flag for the next controller/cloud pass.
- **Retained:** Current US list price for Insight Premium renewal beyond the first free year; whether the 1-year-free-Premium bundle applies to all M4350/M4500 SKUs.

---

## S28. MikroTik 2026 hardware wave — MWC 2026 announcements and RouterOS switches

Earlier passes (S18, R4/R9, T1, S24) covered MikroTik only lightly and had no 2026 launches captured. The MWC Barcelona March 2026 showcase was a major event.

### MWC 2026 announcements (March 2026) `[secondary]` — YouTube review + MikroTik community forum
Sources:
- https://www.youtube.com/watch?v=KATjJn3RYUk (video: "MikroTik did it AGAIN — WiFi 7 + GPON, Scout and more news from MWC 2026")
- https://forum.mikrotik.com/t/new-mikrotik-hardware-showed-in-mwc-2026/268913 (MikroTik official forum thread, 85+ posts from March 2–7, 2026)

| Product | What it is | Notes |
|---|---|---|
| **hEX Pro / hEX Pro PoE** | SMB router with M.2 slot for Docker/containers, 10G | Forum users expect it to beat the RB5009 (A73 @ 2.2 GHz); M.2 storage for containers / Back-to-Home shared folders; pictured serial console (may be AI-generated render — **[unverified]**); RouterOS license level 4 noted as a possible ISP limitation; appears to use Qualcomm IPQ-9570 SoC ([independent] analysis, unconfirmed — switch/L2 hardware-offload support questioned) |
| **mAP ax** | Pocket travel Wi-Fi 6 router | With PoE; forum "buy when available" interest |
| **nRAY gen2** | 60 GHz wireless link up to 10 km with 5 GHz backup | Point-to-point/long-haul |
| **MikroTik Scout** | Mesh communication device working without internet | Targeted at agriculture / rural areas |
| **A42GO-HbeP** | **Wi-Fi 7 + GPON ONT in a single device** | Described as the flagship ISP product; forum users plead for SC/APC connector, not UPC |
| **CRS816** | New-generation Cloud Router Switch (98DX3550 switch chip) | Hardware-offloaded MACsec (IEEE 802.1AE GCM-AES-128/256 and XPN), flow-aware VXLAN/VXLAN-GPE/Geneve/IP-GRE/EVPN/SRv6/MPLS-SR encapsulations, hardware-level stacking solution; 50G QSFP28 ports (behave as single/combined stacking ports, likely SFP56-compatible — **40G/100G transceiver support doubted** [secondary]) |
| **hAP be³ media** | Wi-Fi 7 router, five 2.5 GbE ports, container support, USB, microSD storage, Bluetooth, Thread | **Suggested price $179**; YouTube review (68 days before Sept 2026, i.e., ~July 2026): "may be one of the most feature-packed routers MikroTik has ever produced"; pre-orders via Baltic Networks (video: https://www.youtube.com/watch?v=Xk9S4PPTOQ4) |

### RouterOS/SwitchOS switch price snapshots, South Africa, September 2026 `[secondary]`
- **CRS310-8G+2S+IN** (8× 2.5GbE + 2× SFP+, RouterOS v7, license 5): **R4,090**, expected in stock, prices valid 2026-09-22 (comx-computers.co.za) — https://www.comx-computers.co.za/MT-RBCRS310-8G+2S-IN-MikroTik-Cloud-Router-Switch-8x-Buy-p-290412.php
- **CSS610-8G-2S+IN** (8× GbE + 2× SFP+, SwitchOS): **R2,190**, expected in stock (comx.co.za) — https://www.comx.co.za/MT-RBCSS610-8G-2S-IN-MikroTik-CSS610-8G-2S-IN-Cloud-Buy-p-247806.php

### Gap status
- **RESOLVED:** 2026 MikroTik launch slate documented.
- **Retained:** Actual GA/ship dates and final street prices for hEX Pro, CRS816, A42GO-HbeP, mAP ax, nRAY gen2, Scout (all announced, pricing largely TBD); whether hEX Pro's renders/console port are accurate; US/EU reseller availability.

---

