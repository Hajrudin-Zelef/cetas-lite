---
id: etape6-phaseb-smb-networking/04-supplementary-complementary-research-pass-round-2-september-/overview
title: "Supplementary / Complementary Research Pass — Round 2 — September 22, 2026"
domain: supplementary-complementary-research-pass-round-2-september-
role: deep-dive
task: reference
actors: ["EU"]
dates: ["2026-07", "2026-09-22"]
keywords: ["research", "licenses", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [583, 638]
section: "Supplementary / Complementary Research Pass — Round 2 — September 22, 2026"
sha256: 6f029820bde66cb23ea5478c041c1a97a347de0e4f04539f9af8ed4b829fa886
---

# Supplementary / Complementary Research Pass — Round 2 — September 22, 2026

**Scope of this pass:** fills gaps NOT covered by the base report or by the first supplementary pass (S1–S13). New in this round: TP-Link Omada Wi-Fi 7 APs (EAP772/EAP783), NETGEAR Pro Wi-Fi 7 AP lineup (WBE710/WBE700/WBE750), Aruba Instant On Wi-Fi 7 status re-check, MikroTik 25G/100G switch pricing table 2026, UniFi Talk 2026 software/ATA updates, UniFi Access 2026 hub/reader additions, homelab favorite switches 2026, Zyxel XGS2220/GS2220 status check. No existing section was altered.

Provenance tags: [official] / [vendor-reported] / [independent] / [secondary] / [unverified] — applied per-fact.

---

## R1. TP-Link Omada Wi-Fi 7 APs — EAP772 / EAP772-Outdoor / EAP783 (complements S3 which had EAP770 BE11000)

Omada's 2026 Wi-Fi 7 AP lineup is fully shipping at retailers with dated Sep 2026 price snapshots [independent/secondary]:

- **EAP772 (BE9300) ceiling-mount** — tri-band Wi-Fi 7: 688 Mbps (2.4 GHz) + 2,882 Mbps (5 GHz) + 5,765 Mbps (6 GHz), 6 streams, 1× 2.5G RJ45 (PoE+ 24.05W / DC 20.92W), Bluetooth 5.2, 320 MHz (EHT320), MLO, 4K-QAM, 380+ concurrent clients, 140 m² (1,500 ft²) coverage, 24 SSIDs (8/band), WPA3, 802.11k/v/r, Omada Mesh. Price: **AUD 335.21 ex GST / AUD 368.73 incl GST** at PB Tech AU (listing created Sep 2026; 4 units left, ships Sep 29) [independent]. (Note: PB Tech's own "Features" block mislabels it "BE11000" — datasheet rates are BE9300; **flag**.)
  Source: https://www.pbtech.com/au/product/NAPTPL7732/TP-Link-Omada-EAP772-BE9300-Ceiling-Mount-Tri-Band
- **EAP772-Outdoor** — tri-band omnidirectional Wi-Fi 7, BE9300 (5,765 Mbps 6 GHz + 2,882 5 GHz + 688 2.4 GHz), **IP68**, 1× 2.5GbE PoE+ (21.5W), integrated GNSS (GPS/Galileo/GLONASS/BeiDou) receiver for AFC-compliant outdoor 6 GHz, 300 m² coverage, 380+ clients, wall/pole mount kits. Price not captured — **flag** [vendor-reported/secondary].
  Source: https://www.linkqage.com/product/tp-link-omada-eap772-outdoor-tri-band-omnidirectional-be9300-indoor-outdoor-wi-fi-7-access-point/
- **EAP783 (BE19000)** — tri-band 12-stream flagship: 11,520 Mbps (6 GHz) + 8,640 Mbps (5 GHz) + 1,376 Mbps (2.4 GHz), **2× 10G ports** (20G aggregation), PoE++ (39W), 320 MHz, MLO, mesh, AI roaming. Price: **AUD 815.30 ex GST / AUD 896.83 incl GST** at PB Tech AU (Sep 2026, out of stock; earlier AU listing AUD 842.61) [independent]; NZ: NZD 1,042.61 ex GST / NZD 1,199 incl GST (out of stock) [independent].
  Sources: https://www.pbtech.com/au/product/NAPTPL7831/TP-Link-Omada-EAP783-BE19000-Tri-Band-12-Stream-Wi ; https://www.pbtech.co.nz/product/NAPTPL7831/TP-Link-Omada-EAP783-BE19000-Tri-Band-Wi-Fi-7-Acce?qr=related-view
- Positioning: EAP772 = value tri-band Wi-Fi 7; EAP783 = high-density flagship — all under the free Omada SDN controller (no AP licenses) [independent assessment].

## R2. NETGEAR Pro Wi-Fi 7 APs — WBE710 (BE9400) / WBE700 / WBE750 (complements S6)

NETGEAR's Insight-managed Wi-Fi 7 AP line is current in 2026 — all ship with **1-year Insight subscription included**, no hardware controller required [official/vendor-reported]:

- **WBE710 (BE9400)** — tri-band 9.4 Gbps aggregate, 6 streams, 1× 2.5GbE PoE+, 320 MHz (6 GHz) / 160 MHz (5 GHz), WPA3/WPA2-ENT, 8 SSIDs, dynamic VLAN, Insight Mesh backhaul. Datasheet **July 2026** revision (202-12964-01) [official].
  Sources: https://www.netgear.com/au/business/wifi/access-points/wbe710/ ; https://www.downloads.netgear.com/files/GDC/WBE710/WBE710_UM_EN.pdf
- **WBE700** — higher-tier Insight Wi-Fi 7 AP (datasheet current 2026); 1-yr Insight subscription bundled [official].
  Source: https://www.downloads.netgear.com/files/GDC/WBE700/WBE700_DS.pdf
- **WBE750** — launched as NETGEAR's "first Insight-manageable Wi-Fi 7 AP": 4×4 tri-band, 18.4 Gbps theoretical, 600 clients, **NBASE-T 10GbE/5G/2.5G/1G RJ45 uplink with PoE++**, $700 launch price (1-yr Insight included) [secondary].
  Source: https://comkex.com/tech/netgear-introduces-wbe750-first-insight-manageable-wi-fi-7-access-point-targets-congested-deployments/
- Competitive note (Sep 2026 independent review): reviewer contrasts WBE-style enterprise Wi-Fi 7 against UniFi U7-Pro-XG ($199, 10G wired) and Aruba Instant On AP25 (Wi-Fi 6, simpler management) — NETGEAR sits between on price/features with recurring-subscription model [independent].
  Source: https://medium.com/@mikahwiggins/best-wireless-access-points-for-business-4efe06a18c8f

## R3. Aruba Instant On Wi-Fi 7 — status re-check (complements S7)

- **No Wi-Fi 7 Instant On AP confirmed as of September 22, 2026.** Aruba's Wi-Fi 7 APs remain the enterprise 730 Series (Aruba Central only, launched Apr 2024) — Instant On's current flagship AP is the **Wi-Fi 6 AP25** (Sep 2026 independent business-AP review still compares U7-Pro-XG against the AP25) [independent/secondary — flag retained].
  Source: https://medium.com/@mikahwiggins/best-wireless-access-points-for-business-4efe06a18c8f
- This is now an explicit second confirmation of the S7 flag (no Instant On Wi-Fi 7 in 2026).

## R4. MikroTik 25G/100G switch pricing — Sep 2026 snapshots (new detail)

MikroTik remains the budget leader for 25G+ in SMB/homelab — retailer prices observed Sep 2026 [independent]:

| Model | Ports | Price (Sep 2026) | Source |
|---|---|---|---|
| CRS504-4XQ-OUT | 4× 100G QSFP28, outdoor IP66, 802.3bt | **AUD 1,552.03** | wisp.net.au [independent] |
| CRS510-8XS-2XQ-IN | 8× 25G SFP28 + 2× 100G QSFP28, dual hot-swap PSU | **AUD 1,706.36** (backorder) | wisp.net.au [independent] |
| CRS326-4C+20G+2Q-RM | 20× 2.5G + 4× combo (2.5G/SFP+) + 2× 40G QSFP+ | **€659.91** (EU, ex-tax) / **£791.99** inc VAT (UK) | en.cdr.pl / linitx.com [independent] |
| CRS310-8G+2S+IN | 8× 2.5G + 2× 10G SFP+, Marvell 98DX226S | **AUD 390.00** | wisp.net.au [independent] |
| CRS326-24S+2Q-RM | 24× 10G SFP+ + 2× 40G QSFP+ (suggested $499) | **NZD 1,052.17** ex GST / **CAD 775.00** (usedservers, OOS) | pbtech.co.nz / usedservers.ca [independent] |
| CRS320-8P-8B-4S+RM | 8× 802.3bt PoE 1G + 8× PoE 1G + 4× 10G SFP+, 963W budget | **£419.99** inc VAT | linitx.com [independent] |
| CRS418-8P-8G-2S+5axQ2axQ-RM | 16-port PoE switch w/ integrated Wi-Fi 6 (4x4) | **£383.99** inc VAT (awaiting restock) | linitx.com [independent] |

- Note: MikroTik CRS 25G+ switches run RouterOS v7 / SwitchOS with VLANs, MLAG, ACLs — deep feature set at prices no brand competitor matches [independent/secondary assessment]. No new 2026 MikroTik switch launch surfaced in sources searched — **flag**.
  Source: https://wisp.net.au/s/30/mikrotik-network-switches

