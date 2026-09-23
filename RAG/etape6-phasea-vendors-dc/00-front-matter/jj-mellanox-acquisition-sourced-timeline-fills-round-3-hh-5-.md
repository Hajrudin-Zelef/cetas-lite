---
id: etape6-phasea-vendors-dc/00-front-matter/jj-mellanox-acquisition-sourced-timeline-fills-round-3-hh-5-
title: "JJ. Mellanox acquisition — sourced timeline (fills round-3 §HH.5 gap)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "EU", "Nvidia", "United States"]
dates: ["2019-03", "2019-03-11", "2020-04-16", "2020-04-17", "2020-04-27", "2020-10-07", "2020-11", "2020-11-04", "2026-23-10"]
keywords: ["acquisition", "cost", "disclosure", "gpu", "latency", "license", "licenses", "nvidia", "pricing", "regulation", "revenue"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1582, 1626]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 11f501935d5ddcc28495115a1cc6a07db53cd19da4df8dfae3ce9c1376cdf4ce
---

# JJ. Mellanox acquisition — sourced timeline (fills round-3 §HH.5 gap)

### JJ. Mellanox acquisition — sourced timeline (fills round-3 §HH.5 gap)

- NVIDIA announced the Mellanox acquisition in **March 2019**: **$125 per share in cash, ~$6.9B total**, financed entirely with cash, NVIDIA's largest acquisition at the time [secondary — wccftech, via NVIDIA press release, 2019-03-11].
- Deal completed on **April 27, 2020** after a 13-month regulatory process; the final approval came from **China's State Administration for Market Regulation (SAMR) on April 16, 2020** (conditional approval; US, EU and Mexico had approved earlier) [secondary — wccftech; Globes, 2020-04-17].
- **Eyal Waldman**, Mellanox founder/CEO, stepped down in **November 2020** after agreeing to stay through deal completion [secondary — Haaretz, 2020-11-04].
- Post-close, Mellanox's BlueField line became the center of NVIDIA's "three-pillar" data-center vision (CPU + GPU + DPU) articulated by Jensen Huang; Mellanox's InfiniBand remained the backbone of NVIDIA's largest AI clusters [secondary — Haaretz, 2020-10-07].
- 2026 retrospectives: NVIDIA networking hit **$4.9B in Q2 FY2026** (+64% QoQ) and roughly **$15B in Q1 FY2027** (~$60B annualized run rate) — one analyst piece computes the $6.9B bet has "returned roughly 8.7 times its cost in annualized revenue alone" and quotes Huang's earlier "homerun deal" characterization [secondary — ainvest, ~Jun 2026]. Israeli press estimates networking will cross **$60B in FY2027** with ~70% of the tech stack originating from Israeli R&D [secondary — SKN, Sep 2026].

### KK. Meraki MS390 hardware + license pricing (new tier vs round-3 §AA's MS450/MS355)

- **MS390-48P-HW** (48-port PoE+ L3 access): **MSRP $13,497**, street **$8,990 CAD** (DirectDial Canada, marked backordered) [secondary — directdial.com/ca].
- **MS390-48UX-HW** (36× 2.5GbE + 12× mGig UPoE): **MSRP $12,945**, special **$9,320.17** [secondary — 4tekgear.com]. Note: the 4tekgear listing shows an "END OF SALE" marking on the -48UX and -48U SKUs — a **single-reseller marking, not an official Cisco EOS notice**; official Meraki EOS list not located [secondary-flagged].
- **MS390-48U-HW** (48-port UPoE): **MSRP $11,050**, special **$7,480.65** [secondary — 4tekgear.com].
- MS390 7-year licenses (new anchors beyond base §4.4):
  - **LIC-MS390-48E-7Y** (Enterprise): MSRP **$6,543** → **$3,100** street (DirectDial; shows a 25% "Instant Rebate" expiring 10/23/2026) [secondary — directdial.com/us].
  - **LIC-MS390-48A-7Y** (Advanced): MSRP **$15,346.53** → **$7,989.00** street (SHI, in stock) [secondary — shi.com].
  - **LIC-MS390-24E-7Y** (Enterprise, 24-port): MSRP **$3,645.72** → **$1,898.00** street (SHI, in stock) [secondary — shi.com].
- Operational note (reseller disclosure, consistent across US Cisco resellers): authorized Cisco partners **cannot stock Meraki hardware in their own warehouses** — orders ship direct from Cisco (target "10 Days or Less"), and license keys are emailed within 1–3 business days [secondary — 4tekgear.com].

### LL. Dell PowerSwitch SN4700 — Spectrum-3 OEM tier (not covered elsewhere in this file)

- **SN4700**: 32× QSFP-DD 400GbE (up to 128× 100GbE / 64× 200GbE via breakout), **12.8 Tb/s switching, 8.4 Bpps**, 8× 50G PAM4 lanes/port, **620 ns latency**, quad-core x86 CPU, 16 GB DDR4, 64 GB SSD, **64 MB packet buffer**; 2 (1+1) hot-swap PSUs, 6 (N+1) hot-swap fans, reversible airflow; 1U (44 × 428 × 568.5 mm), 11.7 kg; 100–264 VAC [official — Dell SN4700 spec sheet, delltechnologies.com].
- Positioning: 400G spine/leaf tier of the Dell OEM Spectrum family below the SN5600 800G line (Spectrum-4, §D) — the NVIDIA-silicon counterpart to the Broadcom-based Z9432/Z9664 class. SN4000/SN3000 (lower tiers) spec detail was **not located** [gap].

### MM. Aruba CX campus/access tier: 6200F / 6300F / 6300M (new below the 8325/8360)

Base §3.3/round-2 §O covered 8325/8360; the access/aggregation tiers underneath were not detailed:

| Feature | CX 6200F (JL724A) | CX 6300F | CX 6300M |
|---|---|---|---|
| Deployment | access | access | access / aggregation |
| Switching capacity | 176 Gbps | 880 Gbps | 880 Gbps (SFP+ models; newer Smart-Rate models up to 1,760 Gbps — see below) |
| Uplinks | 4× 10G | 4× 10G + 4× 25/50G | 28× 10G + combo 25/50G |
| Smart Rate multi-gig | no | no | yes (48× 1/2.5/5G, up to 90W PoE per port) |
| Max PoE | 1,440W (802.3at, 60W max) | 720W | 2,880W (Class 6/8) |
| Stacking (VSF) | 8 members | 10 members | 10 members |
| Hot-swap PSU/fans | no | no | yes |
| EVPN-VXLAN w/ MP-BGP | no | no | yes |

[secondary — Aruba CX 6300 customer presentation via slideshare; HPE CX 6300 datasheet via hpe.com psnow].

- **CX 6200F-24G (JL724A)**: 24× 10/100/1000BASE-T + 4× 1/10G SFP+, 128 Gbps, 95.2 Mpps, VSF 8, fixed 200W PSU, 1.73 × 17.4 × 12.9 in; Dynamic Segmentation, switch-to-switch VXLAN tunnels [secondary — manuals.plus datasheet mirror].
- **Newer 6300M SKUs** (HPE CX 6300 datasheet, 2026): **S0E91A** (6300M 48p SR10 PTP/AVB Class-8 PoE, 4× 100G MACsec, 1,760 Gbps / 1,310 Mpps, 32 MB buffer); **R8S89A** (6300M 24p HPE Smart Rate 1G/2.5G/5G/10G CL6 PoE + 2× 50G + 2× 25G, 780 Gbps / 580 Mpps); **R8S90A** (6300M 48p Smart Rate CL8 PoE + 2× 50G + 2× 25G) [official — HPE CX 6300 datasheet].
- Licensing stance across the CX campus family: **"No software feature licensing or subscription required"**; Limited Lifetime Warranty [secondary — Aruba presentation].

