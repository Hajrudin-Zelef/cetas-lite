---
id: etape6-phasea-vendors-dc/00-front-matter/g-meraki-ms355-ms450-ms150-delta-addendum-to-round-3b-r3b-a-
title: "G. Meraki MS355 / MS450 / MS150 — delta addendum to round-3b §R3b-A (new listings only)"
domain: front-matter
role: reference
task: reference
actors: ["EU"]
dates: ["2026-07-10"]
keywords: ["benchmark", "cpo", "ethernet", "inference", "license", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1407, 1451]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: f063572e23d093b8b37d5ad81973e1c15fb1999c33e7d5b638e90d3065c8a03f
---

# G. Meraki MS355 / MS450 / MS150 — delta addendum to round-3b §R3b-A (new listings only)

### G. Meraki MS355 / MS450 / MS150 — delta addendum to round-3b §R3b-A (new listings only)

Round-3b §R3b-A already captured MS450-12-HW (MSRP $20,439.21 / street $17,938.91, datadirectglobal $20,571), MS355-48X2-HW at barcodesinc $12,866.95, and the MS355 license 5YR/1YR terms — those are not repeated here. Net-new reseller snapshots from this pass [secondary]:

- **4tekgear MS450-12-HW listing anomaly**: the same listing shows Special **$17,938.91**, Regular **$29,463.61**, and a listed MSRP of **$20,439.21** — special price below stated MSRP while "regular" is above it; inconsistent list anchors on one page, flagged as reseller-display inconsistency, not a pricing signal [secondary — 4tekgear.com].
- **MS355-48X2-HW additional secondary listings** (R3b-A had only barcodesinc): reconditioned **$10,209.99** (networkequipment.net); refurbished **$2,999** (networkgenetics, "unclaimed"); seller-refurbished **$1,032.99** (networktigers) — refurb variance spans nearly 10×, so treat any single listing as non-representative [secondary].
- **New license terms** (R3b-A covered 5YR/1YR; these terms add): **LIC-MS355-48X2-7YR** street **$3,329.70** (list $6,954.26); **LIC-MS355-48X2-3YR** street **$1,427.02** (list $2,980.40) [secondary — hummingbirdnetworks.com]. Per-switch Enterprise license; 7-year term ≈ 2.33× the 3-year street price.
- **MS130 family detail** (hardware detail not in R4-G's status finding): **10 models**, up to **740W PoE**, multigigabit ports, SFP+ uplinks, **802.3bt up to 30W/port**, 2.5GbE RJ45 on select SKUs; compact branch/campus positioning [secondary — telecom4good.org]. Consistent with R4-G §R4-G's finding of no new 2026 Meraki hardware family — this is depth on an existing family.
- Street-vs-list observation (computed from R3b-A + this pass): Meraki street pricing lands at **~45–55% of list** on typical reseller channels [independent — analyst inference].

### H. Meraki platform evolution 2026 (new)

- **Meraki Dashboard API**: **v1.73.0** (Aug 2026, 924 ops) → **v1.74.0** (Sep 2026, 894 total ops, 6 new + 38 enhanced) [secondary — ciscodevnet/meraki-portal-what-s-new-changelog, 2026.md].
- Switch-relevant API additions 2026: **fastPoe / perpetualPoe** port properties; assurance alert types **copp_drops, fan_tray_failure, line_card_insufficient_power_shutdown**; **campusGateway clusters** (tunneling batchUpdate); staged firmware events now model **switchCatalyst** products [secondary — same changelog].
- Community evidence: Meraki API ecosystem active in 2026 (meraki-dashboard-exporter v2.1.0, Sep 12, 2026; panoramicdata/meraki.api tracking v1.70.0) [secondary].
- Confirms the platform direction from base §4.2: cloud-managed Catalyst convergence is now API-visible (firmware-staged events for Catalyst switches), not just dashboard UI [secondary].

### I. ConnectX-7 / BlueField-3 / ConnectX-8 2026 status and EU street prices (new)

- **ConnectX-7**: status **SHIPPING**, verified 2026-07-10 (independent SKU database) — 400 Gb/s VPI (NDR IB or 400GbE), PCIe Gen5 x16, 330–370 Mmsg/s RDMA, SHARP, ASAP², inline IPsec/TLS/MACsec [independent — gpusmith.com datasheet mirror].
- Firmware: ConnectX-7 at **28.48.1000** (released Feb 2026) seen in an independent benchmark rig [secondary — fastswan.org, 2026].
- EU street snapshots (galaxus.de, Sep 2026 crawl) [secondary]: ConnectX-7 200G VPI dual-port **€2,448**; **ConnectX-8 C8180 HHHL 800G SuperNIC €2,529**; **BlueField-3 B3210E E-Series DPU €5,020**; BlueField-3 B3220L SuperNIC €4,647; ConnectX-6 Dx 100G still listed at **€257**; ConnectX-6 Lx at €689.
- Mellanox-heritage context: ConnectX-6 Dx/Lx remain orderable at the low end while ConnectX-7/8 cover the 200–800G AI/HPC tier; the pricing ladder shows ~10× from CX-6 Dx (€257) to BlueField-3 (€5,020) [independent — analyst inference].

### J. Aruba CX 8350 — finding: no such product

- Targeted searches for "HPE Aruba CX 8350" surfaced only unrelated models (Aruba 2930M, Instant On) — **no CX 8350 product line exists**; the 83xx tier is covered by the **CX 8325** (covered in Round-2 §O) and **CX 8360 v2** (covered in Pass #1 §I / Round-2 §O) [finding — negative result].

### K. Dell PowerEdge MX switching modules — 2026 status (new)

- **MX9116n Fabric Switching Engine** and **MX5108n Ethernet Switch** remain the current MX7000 I/O options (alongside MX7116n Fabric Expander, MXG610s FC, MX5000s SAS) per the 2026 MX7000 installation manual [official — Dell support manual, 2026 crawl].
- MX9116n: 16× 25 GE server-facing + 12× QSFP28-DD (breakout to 8× 10/25GE or 2× 40/100GE) + 2× QSFP28 uplinks + 2× unified Ethernet/FC ports (8/16/32GFC); FC Direct Attach and NPIV Proxy Gateway modes documented in the 2026 deployment guide [official — Dell InfoHub, 2026]. **No end-of-sale announcement located** [unverified — continuation assumed].

### L. Supplementary Pass #3 verification log (new open items)

1. Z9964F-ON exact NOS support document (Enterprise SONiC version) — not located [gap].
2. Cumulus Linux 5.18 feature deltas (docs exist in draft; feature list not extracted) [gap].
3. UFM Enterprise 6.26.1 feature deltas vs 6.23/6.22 [gap].
4. UFM licensing model (per-switch vs per-port) [gap].
5. MS450-12-HW reseller price inconsistency ($17.9K special vs $20.4K MSRP on same listing) [secondary — flagged].
6. Spectrum-X1600 / CPO performance claims — vendor-reported, no independent measurement located [unverified].
7. Spectrum-X vs Etherlink head-to-head independent benchmark — not located; only vendor blogs and community analysis [gap].

### M. Supplementary Pass #3 sources (verbatim URLs)

