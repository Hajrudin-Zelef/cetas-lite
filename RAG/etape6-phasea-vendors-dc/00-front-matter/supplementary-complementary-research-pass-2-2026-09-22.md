---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-2-2026-09-22
title: "Supplementary / Complementary Research Pass #2 — 2026-09-22"
domain: front-matter
role: reference
task: reference
actors: ["AWS", "CoreWeave", "Google", "Microsoft", "Nebius", "Nvidia", "Oracle", "SpaceX"]
dates: ["2024-12-18", "2026-02", "2026-06", "2026-07", "2026-07-22", "2026-08", "2026-08-18", "2026-09-17", "2026-09-22"]
keywords: ["research", "asic", "aws", "distribution", "ethernet", "gpu", "gpus", "license", "nvidia", "nvlink", "pricing", "rack-scale"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [494, 541]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: a174fc67e41afd33c17054a6e73707833567208900d259b1452e8ec92ba40d14
---

# Supplementary / Complementary Research Pass #2 — 2026-09-22

## Supplementary / Complementary Research Pass #2 — 2026-09-22

**Scope note:** This is a second additive pass. It contains only material gathered after the first supplementary pass was written (same date, independent research). It does not duplicate §§1–8 or the first supplementary pass (§§A–M). Earlier sections were left untouched. Read the base report and Supplementary Pass #1 first; the provenance legend from the base applies here.

### N. Dell SmartFabric OS10 — 2026 release train (new detail)

- **OS10 10.6.0.9** — Release Notes dated **July 2026** (Rev. A00), release type "Minor (MI)". Supported hardware includes S3048-ON, S3248T-ON, S4048-ON/T-ON, S4112F/T-ON, S4128F/T-ON, S4148F/FE/T/U-ON, S4248FBL-ON, S5232F/S5248F/S5296F-ON, S5212F/S5224F-ON, S5448F-ON, Z9100-ON, Z9264F-ON, Z9332F-ON, N3248TE-ON, Z9432F-ON, E3224F-ON, Z9664F-ON [official — Dell release notes mirror, 2026-07].
- **OS10 10.6.1.3** (August 2026): SmartFabric OS10 now "operates with a virtual license applied" — `show license status` displays "Virtual License Applied" for license-related fields [official — Dell Installation/Upgrade/Downgrade Guide revision history, A25 2026-08-18].
- **OS10 10.6.1.0** line: a June 2026 Security Configuration Guide documents X.509v3 certificate handling for 10.5.0.7P3–10.6.1.x; OS 10.5.0.7P3 images display a warning that the cluster manager is using default credentials until a valid certificate is installed [official — Dell Security Configuration Guide, 2026-06].
- **OS10 10.5.6.7** maintenance release for PowerSwitch + PowerEdge MX users (release notes current in Dell support portal, 2026) [official].
- Platform direction notes: base Linux distribution upgraded from **Debian 10 (Buster) to Debian 12 (Bookworm)** starting with **10.6.0.1** (Dec 2024); new CLI command `show image gpg-keys` added; **SmartFabric Director (SFD) deprecation** announced in 10.6.0.1 [official — Dell upgrade guide revision history, A20 2024-12-18].

### O. NVIDIA Spectrum-6 standalone platform launch — July 2026 (new detail)

The base report (§2.2) and Pass #1 cover Spectrum-6 via Dell OEM (SN6000) and the Vera Rubin bundle. A **standalone Spectrum-X/Spectrum-6 platform announcement in July 2026** added [vendor-reported — NVIDIA press release via press syndication, 2026-07-22/23]:

- **Spectrum-6 = 102.4-terabit-per-second Ethernet switch system**, 2× the capacity of the previous generation, "designed to supply the bandwidth, scale and telemetry needed to run an AI factory as a single end-to-end computing system" [vendor-reported].
- Positioned as the **cornerstone of the Spectrum-X Ethernet platform**; one of six components of the Vera Rubin architecture (Vera CPU, Rubin GPU, NVLink 6 switches, ConnectX-9 SuperNICs, BlueField-4 DPUs, Spectrum-6 switches) [vendor-reported].
- First adopters named: **CoreWeave, Microsoft, Nebius, SpaceXAI, Tesla** — "to connect hundreds of thousands of GPUs" [vendor-reported]. CoreWeave's Min Jun (director of product, networking) and Nebius's Laurelle Roseman (VP global partnerships) gave supporting quotes on bandwidth/resilience for frontier-model training [vendor-reported].
- Platform composition confirmed: Spectrum Ethernet switches + Spectrum-X SuperNICs + ConnectX NICs + BlueField DPUs + LinkX cabling/transceivers + **Spectrum-XGS for networking between multiple AI data centers** [vendor-reported].
- Financial context: NVIDIA's most recent quarter at the time showed **networking sales at $11B, +263% YoY**; CEO Jensen Huang: "We're now the largest networking company in the world" [vendor-reported — Q2 FY2027 earnings call, Aug 2026].
- Timeline note: Vera Rubin entered **full production announced at GTC 2026** (seven chips, five rack-scale systems); cloud/hardware partners expected to ship Vera Rubin products H2 2026: AWS, Google Cloud, Microsoft Azure, Oracle, Dell, HPE, Lenovo, Supermicro [secondary — TechRepublic GTC 2026 live blog].

### P. ConnectX-9 SuperNIC GA detail (February 2026)

- **GA firmware v82.48.1000, February 2026** — initial GA release notes [official — NVIDIA ConnectX-9 firmware release notes, 2026-02].
- Specs confirmed: up to **800 Gb/s per port over InfiniBand and Ethernet**; drives up to **1.6 Tb/s of throughput to Rubin GPUs**; fully integrated with Spectrum-X Ethernet and Quantum-X800 networking; advanced programmable IO and intelligent congestion control [official].
- Physical design: **200G PAM4 SerDes**, PCIe Gen 6 support; NVIDIA claims **4× larger clusters vs BlueField-3** SuperNIC; enhanced programmable IO with SDKs for AI workloads [vendor-reported — SDxCentral, GTC 2026 D.C. briefing; ServeTheHome CES 2026 keynote coverage].
- Deployment evidence: CoreWeave's Rubin deployment (504 GPUs reported Sep 2026) uses **two ConnectX-9 SuperNICs per Rubin GPU**, i.e. **1.6 Tbps of scale-out connectivity per GPU**, in a two-tier non-blocking multi-rail/plane design [secondary — neoteo.com, 2026-09-17].

### Q. Dell Enterprise SONiC — pricing status 2026 (no public prices; spec-sheet depth)

Pass #1 (§A) documented the subscription structure (1/3/5-yr, Standard/Premium, speed-banded). This pass confirms the **dollar-price gap remains as of 2026-09-22**: TrustRadius: "Dell Enterprise SONiC Distribution does not currently have any pricing plans listed at this time" [secondary — TrustRadius, 2026].

The 2026 Dell Enterprise SONiC spec sheet adds technical depth (useful for feature comparison, not present elsewhere in this report) [official — Dell Enterprise SONiC spec sheet © 2026]:

- AI/core/edge features (model- and ASIC-dependent), Anycast Gateway, BFD, BGPv4/v6 + unnumbered BGP, EVPN with multihoming, symmetric/asymmetric IRB, LAG/MCLAG/LACP fallback, QinQ, RoCEv2 with DCBx, VRF, VRRPv4/v6, L2+L3 VXLAN [official].
- Certifications: **FIPS 140-2**, USGv6-R1; management: ZTP, gNMI/REST/MF-CLI/OpenConfig, Ansible collection, SmartFabric Manager, Prometheus, Telegraf, sFlow, RSPAN/ERSPAN [official].
- Validated solutions: VxRail, PowerStore, PowerFlex PNA, Microsoft Azure Stack, OpenStack [official].

### R. Meraki MS150 pricing snapshots (new detail)

- **MS150-48MP-4X** (48-port: 32× 1G + 16× 5GBase-T, 4× 10G SFP+, PoE++ 740W): **MSRP $18,694.47, street $8,586.00** (SHI) [secondary — shi.com]. UK reseller: **£4,836.23 excl. VAT (£5,803.48 incl.)**, 5 in stock (2026) [secondary — networkwarehouse.co.uk].
- **MS150-24P-4G** (24× 1G PoE+ 370W, 4× 1G SFP): **£1,386.56 excl. VAT** (UK, 2026) [secondary].
- **MS150-48T-4G** (48× 1G, 4× 1G SFP): **£1,596.04 excl. VAT**; **MS150-24T-4G**: **£1,918.03 excl. VAT** (UK reseller, 2026) [secondary]. (24T price above 48T is a reseller-listing artifact; flagged, not a Cisco price signal.)
- Series context: MS150 = **12 models**, stackable L2 access switches, up to 740W PoE, 2× dedicated stack ports (80G stacking bandwidth, cables sold separately), non-blocking backplane (e.g. 304 Gbps on 48MP-4X), Adaptive Policy (Adv license), 802.1X, perpetual PoE; static routing "coming soon"; managed via Meraki Dashboard; **mandatory per-device cloud license** [official-via-reseller — networkwarehouse.co.uk, Meraki docs].
- License anchor found: **LIC-MS-100-L-E** (MS Series Large Essentials + Support, annual) listed at **$362.99** street (2026, reseller; on backorder) [secondary — pc-canada.com].

