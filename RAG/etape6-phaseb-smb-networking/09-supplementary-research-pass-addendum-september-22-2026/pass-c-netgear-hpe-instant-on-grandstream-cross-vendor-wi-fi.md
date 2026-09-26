---
id: etape6-phaseb-smb-networking/09-supplementary-research-pass-addendum-september-22-2026/pass-c-netgear-hpe-instant-on-grandstream-cross-vendor-wi-fi
title: "Pass C — Netgear, HPE Instant On, Grandstream, cross-vendor Wi-Fi 7 table, 25G/40G pricing"
domain: supplementary-research-pass-addendum-september-22-2026
role: deep-dive
task: pricing
actors: ["CISA", "United States"]
dates: ["2024-03", "2024-08", "2026-01-27", "2026-02", "2026-03", "2026-05", "2026-06-30", "2026-09", "2026-09-22"]
keywords: ["pricing", "ethernet", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1638, 1691]
section: "Supplementary Research Pass — Addendum (September 22, 2026)"
sha256: a305a1ee5645aa1d3d966e62e3d5609c8b9eb5f76043a4cf065e3a17c9af4a57
---

# Pass C — Netgear, HPE Instant On, Grandstream, cross-vendor Wi-Fi 7 table, 25G/40G pricing

1. XS1935-12F port mix: retailer listings were inconsistent; per U4's official-datasheet correction the -12F is **10× SFP+ + 2× multi-gig RJ45** — do NOT repeat the "all-fiber" label from the superseded draft [corrected].
2. USW-Pro-XG-Aggregation surcharge differs by channel ($197 ui.com vs $147 DoubleRadius) [unverified].
3. ER-X-SFP delisted at one retailer only — not a formal EOL announcement; treat as availability note.
4. CISA BOD 26-02 is Feb 2026 — do not misdate.
5. Ax Wireless ITC: pending status inferred from absence of a found ruling, flagged [unverified].
6. TRENDnet URL typo: trendnet.com product path "TEG-3824WS-v1" vs catalog "TEG-3284WS" [unverified].
7. Still not researched (open for a future pass): Zyxel XS1935-10 / -12HP street prices; D-Link DAP Wi-Fi 7 street prices; Netgear M4350 2026 refresh pricing; Grandstream GWN7670LR street price; D-Link DAP-E3620/E9560 (covered by another pass — do not duplicate).


---

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

