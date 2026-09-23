---
id: etape6-phasea-vendors-dc/00-front-matter/g-meraki-ms355-ms450-ms150-delta-addendum-to-round-3b-r3b-a-
title: "G. Meraki MS355 / MS450 / MS150 — delta addendum to round-3b §R3b-A (new listings only)"
domain: front-matter
role: reference
task: reference
actors: ["EU", "Nvidia"]
dates: ["2026-07-10", "2026-09-22"]
keywords: ["benchmark", "cpo", "ethernet", "inference", "license", "nvidia", "pricing", "research"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1407, 1483]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 665e45cb194961d7fc151ae48fc8bc72699be02a9f396660dbc644c1465cd3c3
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

- https://www.dell.com/support/kbdoc/en-ca/000192674/smartfabric-os10-hardware-compatibility-list
- https://www.dell.com/support/manuals/en-us/secure-connect-gateway-ve/scg_5.x_ve_sm/powerswitch-switches-with-enterprise-sonic-operating-system?guid=guid-e9fa4676-f185-4f30-920b-9934ad2f545f&lang=en-us
- https://www.dell.com/support/kbdoc/en-us/000220710/dell-emc-networking-sonic-os-platform-and-package-support
- https://www.businesswire.com/news/home/20251209613448/en/Aviz-Networks-Unveils-Turnkey-Enterprise-Grade-SONiC-As-Open-Network-Adoption-Accelerates
- https://expresscomputersystems.com/collections/all/products/dell-emc-powerswitch-z9664f-on-64-port-400gbe-qsfp56-dd-switch
- https://www.delltechnologies.com/asset/no-no/products/networking/technical-support/dell-powerswitch-z9664f-on-spec-sheet.pdf
- https://lenovopress.lenovo.com/lp2451.pdf
- https://content.etilize.com/Additional-pdf1/1076533848.pdf
- https://resource.fs.com/mall/doc/20221205184410kkyxaf.pdf
- https://www.lightoptics.co.uk/blogs/news/nvidias-new-ethernet-networking-platform-for-hpc-ai-and-data-center
- https://docs.nvidia.com/networking/display/ufmenterpriseumv6231/gnmi-telemetry-plugin
- https://docs.nvidia.com/networking/display/ufmenterpriseumv6175/release+notes
- https://developer.nvidia.com/blog/benchmarking-nvidia-spectrum-x-for-ai-network-performance-now-available-from-supermicro/
- https://cryptobriefing.com/nvidia-spectrum-x-ethernet-ai-networking/
- http://gpusmith.com/articles/en/pdfs/infiniband-vs-spectrum-x-vs-roce-vs-ethernet-ai-clusters.pdf
- https://github.com/abuabdurahman82/llm-systems-wiki/blob/HEAD/AI-Factory-Networking/26-arista-etherlink.md
- http://gpusmith.com/datasheets/nvidia-connectx-7-datasheet.pdf
- https://www.fastswan.org/Nvidia-ConnectX-7-Benchmark/
- https://www.galaxus.de/en/s1/product/nvidia-connectx-7-vpi-200gbehdr-ib-dual-port-network-cards-58022449
- https://4tekgear.com/meraki-ms450-12-hw-layer-3-switch-hardware-only.html
- https://www.networkgenetics.net/cisco-meraki-ms355-48x2-hw-24x-10gb-rj-45-24x-rj-45-2x-qsfp-unclaimed-switch/
- https://www.networktigers.com/products/ms355-48x2-hw-cisco-switch
- https://networkequipment.net/products/cisco-meraki-ms355-48x2-hw-ref?_pos=20&_fid=2515010fc&_ss=c
- https://www.hummingbirdnetworks.com/cisco-meraki-lic-ms355-48x2-7yr/
- https://www.hummingbirdnetworks.com/cisco-meraki-lic-ms355-48x2-3yr/
- https://github.com/ciscodevnet/meraki-portal-what-s-new-changelog/blob/HEAD/documents/2026.md
- https://telecom4good.org/product/cisco-meraki-switch-ms130-8-compact/
- https://www.dell.com/support/manuals/en-us/poweredge-mx840c/pemx7000_ism_pub/Next-Generation-Modular-overview?guid=guid-00a96ee7-973e-4fc4-b710-89361324e0ba&lang=en-us
- https://infohub.delltechnologies.com/en-us/l/dell-poweredge-mx-networking-deployment-guide-1/scenario-6-connect-mx9116n-fse-to-fibre-channel-storage-fc-direct-attach-6/

**Supplementary Pass #3 collection metadata:** read-only web research (browser_search + one page fetch, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §L. Existing sections §§1–8, §§A–M, §§N–X (Pass #2) and §§N–S (round 2) were not modified.

