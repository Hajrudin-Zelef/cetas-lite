---
id: etape6-phaseb-smb-networking/09-supplementary-research-pass-addendum-september-22-2026/overview
title: "Supplementary Research Pass — Addendum (September 22, 2026)"
domain: supplementary-research-pass-addendum-september-22-2026
role: deep-dive
task: reference
actors: ["Apple", "CISA", "California", "EU", "Malaysia", "United States"]
dates: ["2026-09-22"]
keywords: ["research", "benchmark", "license", "memory", "pricing", "voice"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1572, 1637]
section: "Supplementary Research Pass — Addendum (September 22, 2026)"
sha256: 1727e75e9b753c0f5219bec444ae5bb58d467478fee5dd3116d5313fed2f58ff
---

# Supplementary Research Pass — Addendum (September 22, 2026)

**Scope:** addendum written after the R1/R5/R7/S20/U2/U4 waves were added to this file. Keeps only genuinely new material in full: **TRENDnet 2026 catalog (new vendor detail)**, **Zyxel GS1915 street prices (closes the repeatedly-flagged §3.2 gap)**, **EdgeRouter/EdgeMAX 2026 status (new)**, **Ax Wireless ITC status re-check (new)**. Two shorter items are kept as explicit addenda to already-resolved topics: Omada multi-year license data points (addendum to R1) and USW-Pro-XG-Aggregation checkout totals/street snapshots (addendum to R5). The XS1935-12F street-price write-up was **removed** — it duplicated R7/S20 — and its port-mix question is settled by U4's official-datasheet correction (10× SFP+ + 2× multi-gig RJ45, not "all-fiber"). The "Round 3" title used in the superseded draft was also renamed to remove a header collision with two other passes in this file. No other section was altered.

Provenance tags: [official] / [vendor-reported] / [independent] / [secondary] / [unverified] — applied per-fact.

---

## B1. TRENDnet 2026 — full portfolio snapshot (new vendor; was: only R8 budget-pick mentions)

TRENDnet (founded 1990, Torrance, California) published its **2026 EU product catalog** (PDF, dated ~Sep 2026) covering multi-gigabit and 10G switching for SMB [official]:

- **Web Smart multi-gig line (2026 catalog)**:
  - **TEG-3284WS** — 24× 2.5G RJ45 + 4× 10G SFP+, 200 Gbps switching, L2+ (VLAN, LACP, QoS, SNMP, IGMP snooping, voice VLAN, IPv6), smart fan / fanless design, TRENDnet Hive cloud management (no hardware/server required; additional fee may apply), TAA/NDAA compliant (US/CA), Lifetime Warranty [official]
  - **TEG-3182WS** — 16× 2.5G + 2× 10G SFP+ [official]
  - **TEG-3102WS** — 8× 2.5G + 2× 10G SFP+ [official]
  - **TEG-7124WS** — 24× 10G SFP+ (L2) [official]
- **L2 SFP/SFP+ fiber line**: TL2-F7120 (12× 10G SFP+), TL2-F7080 (8× 10G SFP+), TL2-F70284 (24× 1G SFP + 4× 10G SFP+) [official]
- Street pricing [independent]: **TEG-3284WS $582.00** (SHI, Sep 2026, TAA compliant version v1.0R); TL2-F7080 (8× 10G SFP+) ~$229 per 2026 community snapshot [independent]
- Sources: http://downloads.trendnet.com/marketing/catalog/2026_product_catalog_EU_092925_sm.pdf ; https://www.shi.com/product/48974336/TRENDnet-TEG-3284WS ; https://www.trendnet.com/langen/products/2.5g-managed-switch/28-port-multi-gig-web-smart-switch-TEG-3824WS-v1
- Positioning [independent assessment]: web-smart/fanless SMB switches undercutting Omada/UniFi on price, but with a thinner cloud-management story (Hive) and no controller-free ecosystem lock-in.
- Flag: the trendnet.com product URL path references "TEG-3824WS-v1" while the catalog says TEG-3284WS — likely a URL typo [unverified].

## B2. Zyxel GS1915 — street prices (closes the repeated §3.2 gap)

- **GS1915-24EP** (24-port GbE, 8× 802.3at PoE+, 60W… note one listing says 130W budget — variants differ; 48 Gbps fabric, 8K MAC): **€231.99 incl. VAT** (aio.lv, Latvia, 25 pcs in stock, Sep 2026) [independent]; **$239.28** (technologygalaxy.com, US — backordered, ETA unknown) [independent]; **£176.26 ex VAT / £211.51 inc VAT** (screenmoove.com, UK) [independent]; €125.71 (eestipoisid.eu, Estonia, non-PoE 24E variant, Jul 2026) [independent]
- **GS1915-8EP** (8-port GbE, 8× 802.3at PoE+, 60W, 16 Gbps): **$133.05** (technologygalaxy.com, US — backordered) [independent]; **₱9,200** (PC Express, Philippines) [independent]
- The 2026 "Just Connect" GS1915 push is therefore **already in retail channels** in EU/UK/US/PH/Latvia, entry-level pricing confirmed, with US stock tight (backordered at one reseller) [independent assessment]
- Sources: https://aio.lv/en/product--zyxel-gs1915-24ep-eu0101f--5004835/all-questions ; https://www.technologygalaxy.com/Zyxel-GS1915-24EP/p/402523 ; https://www.technologygalaxy.com/Zyxel-GS1915-8EP/p/397462

## B3. EdgeRouter / EdgeMAX — 2026 status (new; not previously researched in this file)

- **ER-X-SFP** (EdgeRouter X SFP, 6-port): **"This item is no longer available"** at Adorama (listing retired) [independent]
- **ER-10X** (EdgeRouter 10X): still listed at EU retailers (cablematic.com, Sep 17, 2026 crawl) [independent]
- **ER-8-XG Infinity** (10G EdgeRouter): still listed at third-party resellers [secondary]
- No formal Ubiquiti EdgeMAX/EdgeRouter EOL bulletin found in 2026 sources searched [unverified — gap, not absence proof]; the line is effectively in **maintenance/still-available-via-channel** status while UniFi gateways get all new development
- Context: CISA **Binding Operational Directive 26-02** (Feb 2026) urges federal agencies to replace end-of-support edge devices (routers/firewalls/switches), noting nation-state actors actively target EOS network edge gear [secondary — SecurityWeek, 2026]
- Sources: https://www.adorama.com/ubierxsfpus.html ; https://cablematic.com/en/products/ubiquiti-er-10x-edge-router-10x-edgemax-router-UI156/ ; https://www.securityweek.com/organizations-urged-to-replace-discontinued-edge-devices/

## B4. Ax Wireless v. Ubiquiti ITC — status re-check (Sep 22, 2026; new)

- **No final determination or exclusion order found** in sources searched as of Sep 22, 2026 — the Section 337 investigation (Feb 2, 2026 complaint) appears still in the investigation/hearing phase [unverified — do not cite a ruling]
- Related context only: the USITC **instituted a separate Section 337 investigation (337-TA-1520, Aug 28, 2026)** against Apple on mobile devices (StayTouch complaint) — showing the Commission's 2026 caseload context; not related to Ax Wireless [official]
- Sources: https://www.usitc.gov/press_room/news_release/2026/er0828_69139.htm ; https://www.tipranks.com/news/the-fly/ubiquiti-discloses-complaint-with-itc-filed-by-ax-wireless-thefly
- Kovalenko v. Ubiquiti and the Velocity suit (§8): no new 2026 developments surfaced in this pass — flagged, not researched further

## B5. Addendum to R1 — Omada Cloud Standard multi-year license data points (new)

R1 resolved the headline pricing; these additional sightings are new:

- **LIC-OCC-3YR** (1 device, 3 years): **AUD $59.00** (themediajoint.com.au — on backorder) [independent]
- **1-year equivalent**: **RM 45.00** (Malaysia, pre-order) [independent]
- Sources: https://www.themediajoint.com.au/store/p/tp-link-omada-cloud-based-controller-3-year-license-fee-for-one-device ; https://wamatek.com/product/tp-link-omada-cloud-based-controller-license-1-device/
- Reminder (R1): US ~$10–14/device/yr (CDW $13.99; Wamatek $9.99 out of stock); CA 5-yr $63 (DirectDial). Use channel price as displayed; no normalized benchmark.

## B6. Addendum to R5 — USW-Pro-XG-Aggregation checkout totals and street snapshots (new)

R5 resolved the $2,499 MSRP; these are new:

- **USW-Pro-XG-Aggregation** (32× 25G SFP28, L3, Etherlighting, 200W max): **$2,499.00 MSRP + $197 memory surcharge = $2,696.00** at checkout (store.ui.com/us, crawled ~Sep 2026) [official]; 5-year coverage add-on **$499/unit** [official]
- Street snapshots [independent]: Adorama **$2,599.00** (ships from manufacturer); Newegg Q&A price reference **$2,646.00**; DoubleRadius **$2,499.00 + $147 surcharge at checkout** (out of stock); BIG W Australia **AUD $5,478**
- **Flag:** memory-surcharge amounts differ by channel ($197 on ui.com vs $147 at DoubleRadius) — take channel price as displayed [independent]
- Sources: https://store.ui.com/us/en/products/usw-pro-xg-aggregation ; https://www.adorama.com/ubiquiti-networks-pro-xg-aggregation-32-port-200w-managed-switch/p/ubuswproxgag ; https://shopdoubleradius.com/products/ubiquiti-unifi-pro-xg-aggregation-usw-pro-xg-aggregation

## B7. Verification log

