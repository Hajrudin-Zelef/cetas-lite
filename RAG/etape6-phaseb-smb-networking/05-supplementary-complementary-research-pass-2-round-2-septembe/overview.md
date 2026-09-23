---
id: etape6-phaseb-smb-networking/05-supplementary-complementary-research-pass-2-round-2-septembe/overview
title: "Supplementary / Complementary Research Pass #2 (Round 2) — September 22, 2026"
domain: supplementary-complementary-research-pass-2-round-2-septembe
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2026-08-04", "2026-09-22"]
keywords: ["research", "cost", "license", "memory", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [702, 750]
section: "Supplementary / Complementary Research Pass #2 (Round 2) — September 22, 2026"
sha256: 732f62891fda31ffb06727b17baa025c0331a7842070f888ccbb266c3ce977b5
---

# Supplementary / Complementary Research Pass #2 (Round 2) — September 22, 2026

**Scope:** second complementary pass on Phase B. Resolves previously open gaps (Omada Cloud Standard pricing, DXS-3130 pricing, Grandstream Wi-Fi 7 status) and adds new 2026 detail: NETGEAR Wi-Fi 7 APs, UniFi Pro XG official US pricing + memory surcharge, UNAS storage lineup, UniFi Access G6 Pro Entry, Zyxel XS1935 retail pricing, EnGenius 2026 launches (Cloud-Lite + L3 core), homelab favorites 2026, FCC router-ban market shock. No existing section was altered.

Provenance tags: [official] = vendor price list/store/datasheet; [vendor-reported] = vendor claim via PR/trade press; [independent] = retailer listing / third-party review; [secondary] = press/blog/forum; [unverified] = single weak source or inferred.

---

## R1. TP-Link Omada Cloud Standard licensing — pricing RESOLVED (was: gap #1)

The main report's open gap "Omada Cloud Standard per-device license price not found" is now resolved via distributor/retailer listings of the **LIC-OCC license SKUs** [independent]:

- **LIC-OCC-1YR** (1 device, 1 year): **$13.99 USD** — CDW (CDW #8343487), listing crawled Sep 2026 [independent]. Also listed at $9.99 (Wamatek, out of stock) [independent].
- **LIC-OCC-3YR** (1 device, 3 years): **$66.40 AUD** — Ezy P.C. Sales (out of stock) [independent]; regular ~$92.95 AUD (Refurbly) [independent].
- **LIC-OCC-5YR** (1 device, 5 years): **$63 CAD** — DirectDial Canada (in stock, UPC 840030712012) [independent]; **$95.49 AUD** incl. GST — ITSpot (sold out) [independent].
- TP-Link's official controller comparison page confirms the tier structure: Essentials = Free, Standard = "Device License Fee", Hardware controller = hardware cost + free cloud, Software controller = free [official].
  Source: https://www.tp-link.com/dk/business-networking/omada-sdn-controller/omada-cloud-based-controller/
- Sources: https://www.cdw.com/product/tp-link-omada-cloud-based-controller-license-1-device/8343487 ; https://www.directdial.com/ca/item/tp-link-omada-cloud-based-controller/lic-occ-5yr ; https://ezypcsales.com.au/product/tp-link-omada-cloud-based-controller-3-year-license-one-device-cbc/
- **Context:** at ~$14/device/year (US), Omada Cloud Standard is meaningfully cheaper than Cisco Meraki or Zyxel Nebula Plus/Pro per-device licensing, preserving TP-Link's "licensing-light" competitive angle [independent assessment].
- **Open gap retained:** Omada Pro controller hardware/software SKU + pricing — still not found publicly [unverified].

## R2. D-Link DXS-3130 — street pricing (was: launch announced, no prices)

Multi-distributor pricing, observed Sep 2026 [independent]:

| SKU | Price (Sep 2026) | MSRP/List | Source |
|---|---|---|---|
| DXS-3130-28P (24× 10G multigig PoE++ + 4× 25G SFP28) | **$3,433.09 USD** (CompSource, updated 2026-08-04) | Regular $4,058.99 | compsource.com |
| DXS-3130-28P | $2,916.04 (Multioculus) / $3,887.00 (Hypertec, 4 avail.) | MSRP $4,058.99 | multioculus.com / store.hypertecsp.com |
| DXS-3130-28P | $9,116.79 AUD (eyo.com.au, supplier stock TBA) | — | eyo.com.au |
| DXS-3130-28 (non-PoE) | **$3,078.00** (SHI) | MSRP $3,769.99 | shi.com |

- Non-PoE variant confirms the DXS-3130-28 model number; PoE variant DXS-3130-28P UPC 790069800641, specs: 680 Gbps fabric, 4 MB packet buffer, 32K MAC, 12KB jumbo, ERPS, stackable (9 units / 200 Gbps per launch materials) [secondary/vendor-reported].
- Sources: https://www.shi.com/product/50538104/D-Link-DXS-3130-28-Switch ; https://www.multioculus.com/Products/overview/M025167496 ; https://www.eyo.com.au/659751_d-link-28-port-stackable-10-gigabit-poe-layer-3-switch-with-24-multi-gigabit-10g-poe-ports-and-4-25g-sfp28-ports-poe-budget-up-to-1440w-dxs-3130-28p.html
- **Note:** AU pricing ($9,116 AUD) is roughly 2.6× the US street price — regional premium + import chain, not a global list price [independent assessment; flag].

## R3. Grandstream Wi-Fi 7 — lineup CONFIRMED (corrects main §5 flag)

The main report's flag "no confirmed Grandstream Wi-Fi 7 AP launch found in 2026" is **superseded**: Grandstream has a full Wi-Fi 7 GWN7670-series portfolio, listed on grandstream.com [official] and stocked by distributors [independent]:

- **GWN7670** — indoor dual-band Wi-Fi 7, 2×2:2 MIMO, 3.6 Gbps aggregate wireless / 5 Gbps wired, 2× 2.5G PoE+, 256 clients, 175 m coverage; **$119.00** (VoIP Supply) [independent/official].
- **GWN7670LR** — long-range dual-band Wi-Fi 7 AP/bridge; **$179.00** [independent].
- **GWN7672** — tri-band Wi-Fi 7 (BE11000), 11 Gbps wireless / 10 Gbps wired, 384 clients, embedded controller (50 APs), GDMS/GWN Manager; **$169.00** [independent/official].
- **GWN7672L** — tri-band Wi-Fi 7 lite; **$169.00** [independent].
- **GWN7672WM** — wall-mount tri-band Wi-Fi 7 (BE11000), 11 Gbps wireless / 5 Gbps wired, 384 clients, 2,500 PPSK accounts [official].
- **GWN7674** — pro tri-band Wi-Fi 7 (BE21000), 4×4:4 MU-MIMO on 5/6 GHz, 21 Gbps wireless / 12.5 Gbps wired, 768 clients, independent scanning radio, 1× 10G + 1× 2.5G PoE++ [vendor-reported via reseller specs].
- All: controller-less embedded UI + free GDMS Networking cloud + GWN Manager on-prem; secure boot, per-device certificates [official].
- Sources: https://www.grandstream.com/products/networking-solutions/indoor-wifi-access-points/product/gwn7672 ; https://www.grandstream.com/networking-solutions/indoor-wifi-access-points/product/gwn7672wm ; https://www.ipphone-warehouse.com/grandstream-gwn7674-indoor-tri-band-wi-fi-7-access-point/ ; https://www.voipsupply.com/manufacturer/grandstream/wifi-access-points/gwn-series-routers/grandstream-gwn7001-multi-wan-gigabit-vpn-router

