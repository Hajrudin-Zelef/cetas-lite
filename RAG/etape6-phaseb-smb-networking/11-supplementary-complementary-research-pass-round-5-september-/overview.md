---
id: etape6-phaseb-smb-networking/11-supplementary-complementary-research-pass-round-5-september-/overview
title: "Supplementary / Complementary Research Pass — Round 5 — September 22, 2026"
domain: supplementary-complementary-research-pass-round-5-september-
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2024-12", "2024-12-02", "2026-03", "2026-07-31", "2026-09-22"]
keywords: ["research", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1944, 2015]
section: "Supplementary / Complementary Research Pass — Round 5 — September 22, 2026"
sha256: b539651bc6ef9509b42f038f89478cbf33cee455ca6e5fcca4ba507f0ff24143
---

# Supplementary / Complementary Research Pass — Round 5 — September 22, 2026

**Scope of this pass:** fills gaps NOT covered by the base report, the S1–S13 wave, the S14–S18 third wave, Round 2 (R1–R10), Round 3 (T1–T10), or the U1–U13 fourth wave. New in this round: (1) Ubiquiti E7/E7 Campus/E7 Audience product confirmation — **corrects the "unconfirmed" status** in the header table (line 104) and complements the patent-complaint naming; (2) NETGEAR M4350 MAP prices — **closes the "M4350-16M4V / M4350-16C street prices not captured yet" gap**; (3) EnGenius ECS8830F/ECS8854F datasheet detail — complements the Jul 16, 2026 launch announcement; (4) TP-Link EAP772 UK price data point — complements R1. No existing section was altered.

## V1. Ubiquiti E7 family — commercial-product confirmation (corrects the "unconfirmed" header-table row)

**Correction to the base report:** the header table lists "E7 / E7 Campus / E7 Audience" as unconfirmed enterprise Wi-Fi 7 family named only in the Velocity patent complaint. Additional retail-channel evidence observed September 22, 2026 shows the E7 family as **shipping commercial products** [independent]:

- **E7 (standard, Wi-Fi 7 enterprise AP)** — unveiled **December 2, 2024** and shipping since December 2024 / 2025 across reseller channels; 10-stream Wi-Fi 7, **10GbE** uplink plus redundant GbE port, 1,000+ client capacity [independent via resellers + product coverage].
- **E7 Campus** — shipping and broadly stocked at resellers in 2026; enterprise Wi-Fi 7, 10-stream, 10GbE, 1,000+ clients, **IP67** weather rating (outdoor/indoor), 465 m coverage claim per retailer listing [independent].
- **E7 Audience** — commercialized very-high-density variant; 12-stream, up to ~1,500 clients per secondary coverage [independent/secondary — retailer spec copy, not vendor-confirmed].
- Sources observed September 22, 2026:
  - https://www.galaxus.at/en/s1/product/ubiquiti-e7-8600-mbits-access-points-53700430
  - https://dongknows.com/uniquiti-unifi-e7-enterprise-access-point-is-here/
  - https://9to5mac.com/2025/01/04/ubiquiti-e7/
  - https://global.microless.com/product/ubiquiti-unifi-e7-campus-wifi-7-enterprise-access-point-10-stream-10gbe-465m-coverage-1000-clients-ip67-ipx6-outdoor-indoor-ap-white-e7-campus/
- **Retention note:** the E7 models' named presence in the D. Del. patent complaint (filed Jul 30, 2026) is consistent with a family that has been commercial since late 2024 — the complaint identifies shipping products, not future SKUs [secondary].
- **Remaining flags:** vendor list/MSRP pricing for the E7 family not captured in this pass [unverified]; E7 Audience 1,500-client claim is retailer copy [unverified].

## V2. NETGEAR M4350 — MAP prices Sep 2026 (closes the "street prices not captured" gap)

**Closes open gap #3** noted earlier ("NETGEAR M4350-16M4V / M4350-16C street prices — not captured yet"). US dealer Broadfield (news.broadfield.com, crawled ~28 days before Sep 22, 2026) publishes MAP (Minimum Advertised Price) for the full M4350 line, all listed **IN STOCK** [vendor-reported via reseller]:

| Model | Ports | MAP (USD) | Status |
|---|---|---|---|
| M4350-16M4V (MSM4320) | 16× 2.5G PoE++ (incl. 8× Neutrik etherCON) + 4× 25G SFP28 | **$6,999.99** | in stock |
| M4350-16C (CSM4316) | 16× 100G QSFP28 (AV-over-IP aggregation/core) | **$9,999.99** | in stock |
| M4350-24G4XF (GSM4328) | 24× 1G PoE+ + 4× 10G SFP+ | $2,579.99 | in stock |
| M4350-8M2V (MSM4310) | 8× 2.5G PoE++ + 2× 25G SFP28 | $3,019.99 | in stock |
| M4350-24M4X4V (MSM4332) | 24× 2.5G PoE++ + 4× 10G/multi-gig PoE++ + 4× 25G SFP28 | $4,479.99 | in stock |
| M4350-44M4X4V (MSM4352) | 44× 2.5G PoE++ + 4× 10G/multi-gig PoE++ + 4× 25G SFP28 | $5,599.99 | in stock |

- The M4350-16M4V/-16C entries confirm the two ISE-2026-announced models are commercially orderable (announced available March 2026; this price list post-dates that window) — contradicting the earlier claim in the file's historical table that "no confirmed NEW NETGEAR switch model in 2026" [vendor-reported via reseller, in-stock listing].
- Supplementary detail on the M4350-16M4V SKU naming: reseller channel also lists it as **MSM4320-100NES** (EAN 0606449175073); a Swiss retailer listed it at CHF 7,179.00 with delivery date July 31, 2026 — regional corroboration that the model shipped mid-2026 [independent].
- Sources observed September 22, 2026:
  - https://news.broadfield.com/netgear-m4350-in-stock-at-broadfield/
  - https://www.broadcastbruce.com/product/netgear-av-line-m4350-msm4320/
  - https://alltronstamm.shop.concertopro.ch/catalog/it/export/pdf/urlnew?url=https%3A%2F%2Falltronstamm.shop.concertopro.ch%2Fcatalog%2Fit%2Fproducts%2F4180506-msm4320-100nes-netgear-poe-switch-m4350-16m4v-20-port-sfp%3Fprint%3Dpdf&hash=fce8efd8fea71851c6011f047d22d35555d3f353a383202f33e1e1f1aee040b9

## V3. EnGenius ECS8830F / ECS8854F — datasheet detail (complements the Jul 16, 2026 launch)

- **ECS8830F** [vendor-reported via official product page/datasheet]: 24× 10G SFP+ + 6× 100G QSFP28; **1,680 Gbps switching capacity**; **up to 600 Mpps** forwarding; cloud-managed/stackable L3; enterprise core with **VXLAN/EVPN, MLAG, VSF**; hot-swappable PSUs; OSPF, BGP, VRRP; L2/L3 VPN support; hybrid cloud/standalone/on-prem management.
- **ECS8854F** [vendor-reported via announcement coverage]: 48× 10G SFP+ + 6× 100G QSFP28; **2.16 Tbps** switching capacity; up to 600 Mpps [vendor-reported].
- **ECS6824F** (SMB core/branch/MSP tier): announced in the same wave; port-mix and performance specs not verified in this pass — **gap retained**.
- Street/MSRP pricing for the ECS8830F/ECS8854F/ECS6824F: **not found in retail channels as of September 22, 2026** — gap retained [unverified].
- Sources observed September 22, 2026:
  - https://www.engeniustech.com/ecs8830f-layer-3-stackable-switch.html
  - https://www.engeniustech.com/wp-content/uploads/2026/07/EnGenius-ECS8830F-Datasheet.pdf
  - https://www.engeniustech.com/wp-content/uploads/2026/07/EnGenius-ECS8830F-ECS8854F-Command-Guide-v1.0.pdf
  - http://cerebral-overload.com/2026/07/engenius-strengthens-cloud-managed-layer-3-switching-portfolio-with-new-core-and-aggregation-switches/

## V4. TP-Link EAP772 — UK price data point (complements R1)

- **EAP772 (BE9300 tri-band)**: NetXL (UK) listing observed September 22, 2026 — **£165.49 inc VAT** (£137.91 ex VAT), **63 in stock** [independent]. Additional data point alongside the PB Tech AU price in R1.
- **EAP772-Outdoor**: BE9300/BE11000 and 2.5G/10G uplink claims conflict across retailer listings — **flag retained from R1** [unverified — do not merge].
- Source observed September 22, 2026: https://www.netxl.com/wifi-access-points/tp-link-omada-eap772-wifi-7-access-point/

## V5. Open gaps after Round 5

1. Ubiquiti E7-family official vendor MSRP/list pricing — retailer confirmation exists; Ubiquiti's own product pages not verified in this pass.
2. EnGenius ECS8830F/ECS8854F/ECS6824F street pricing — not found; model-launched but not retail-listed as of September 22, 2026.
3. Zyxel XS1935-10 and XS1935-12HP prices — officially unannounced per Sep 21, 2026 press coverage.
4. TP-Link US regulatory outcome — developing; still [unverified] beyond secondary sources.
5. EAP783 "BE19000 vs BE22000" naming conflict — unresolved.
6. Currency conversions not applied; all prices are raw regional snapshots.

**Round-5 verification log:** read-only web research (browser_search), September 22, 2026. No live-browser visits; nothing sent externally. The E7 correction is additive: the base-report header row and the patent-complaint references were NOT edited. V2 closes the earlier "M4350-16M4V/-16C street prices not captured" gap. V3 is additive to the EnGenius Jul-16-2026 launch (already covered) — only datasheet-level specs and official links are new. V4 is additive to R1.

---

---

