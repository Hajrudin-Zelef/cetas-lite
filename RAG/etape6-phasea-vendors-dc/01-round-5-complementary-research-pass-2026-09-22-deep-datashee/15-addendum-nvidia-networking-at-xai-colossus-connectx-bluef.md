---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/15-addendum-nvidia-networking-at-xai-colossus-connectx-bluef
title: "15. Addendum — NVIDIA networking at xAI Colossus, ConnectX/BlueField roadmap, Meraki price points (September 22, 2026)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: actor-profile
actors: ["Broadcom", "Nvidia", "xAI"]
dates: ["2024-10-28", "2025-08", "2025-10", "2026-03", "2026-09", "2026-09-22"]
keywords: ["nvidia", "compute", "cost", "cpo", "distribution", "ethernet", "gpu", "gpus", "license", "pricing", "research", "rubin"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1896, 1945]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: e3f590206516b6eab67576abf426daec7b0f5b5524d5a8df121e0abfa1c8fcf0
---

# 15. Addendum — NVIDIA networking at xAI Colossus, ConnectX/BlueField roadmap, Meraki price points (September 22, 2026)

## 15. Addendum — NVIDIA networking at xAI Colossus, ConnectX/BlueField roadmap, Meraki price points (September 22, 2026)

### 15.1 xAI Colossus runs on NVIDIA Spectrum-X Ethernet (not just Dell/Supermicro servers)

- Per NVIDIA's October 28, 2024 press release [official — relayed via GlobeNewswire], the "world's largest AI supercomputer built by xAI" runs on NVIDIA Ethernet networking: **100,000 H100/Hopper GPUs linked by Spectrum-X Ethernet**, with BlueField-3 DPUs/SuperNICs. NVIDIA claimed ~95% sustained throughput and zero packet loss — NVIDIA's own claims, not independently measured [vendor-reported].
  - Sources: https://globenewswire.com/news-release/2024/10/28/2970195/0/en/NVIDIA-Ethernet-Networking-Accelerates-World-s-Largest-AI-Supercomputer-Built-by-xAI.html ; https://www.techradar.com/pro/xais-colossus-supercomputer-cluster-uses-100-000-nvidia-hopper-gpus-and-it-was-all-made-possible-using-nvidias-spectrum-x-ethernet-networking-platform
- Per SDxCentral, Colossus uses **Spectrum-X SN5600 switches** and each GPU is associated with a **400 GbE SuperNIC link**; xAI planned to double compute capacity to 200,000 GPUs, and the cluster had not yet exhibited RDMA/roce congestion problems at 100K scale [secondary].
  - Source: https://www.sdxcentral.com/news/xai-to-double-colossus-compute-capacity-reveals-cluster-uses-nvidia-spectrum-x-ethernet/
- Reference design partner Supermicro confirmed its role in "building Colossus" (GPU/server compute layer); the network layer is NVIDIA Spectrum-X + BlueField-3 [secondary — VentureBeat]: https://venturebeat.com/ai/building-colossus-supermicros-groundbreaking-ai-supercomputer-built-for-elon-musks-xai
- Note: this complements §2/§7 Dell/Supermicro server coverage — compute supply (Dell ~50K, Supermicro ~50K phase 1) is separate from the NVIDIA Spectrum-X fabric.

### 15.2 ConnectX/BlueField cadence 2025–2028 (roadmap)

- **ConnectX-8**: 800 Gb/s per GPU; part of Vera Rubin-generation Spectrum-X fabric (Spectrum-X switches + ConnectX-8 SuperNICs), launched with Spectrum-XGS (August 2025, geo-distributed AI factories) [official — NVIDIA investor release]: https://investor.nvidia.com/news/press-release-details/2025/NVIDIA-Introduces-Spectrum-XGS-Ethernet-to-Connect-Distributed-Data-Centers-Into-Giga-Scale-AI-Super-Factories/default.aspx
- **ConnectX-9**: 1.6 Tb/s per GPU, RDMA, PCIe Gen 6, hardware crypto at line speed; introduced at NVIDIA GTC DC (Washington, D.C., October 2025); part of the Vera Rubin platform [secondary — Network World, SDxCentral].
  - Sources: https://www.networkworld.com/article/4080459/nvidia-looks-to-power-ai-factory-networks.html ; https://www.sdxcentral.com/news/nvidia-reveals-next-gen-dpu-to-help-offload-gigascale-ai-infrastructure/
- **ConnectX-10**: future roadmap item (post-CX9, alongside Spectrum-7 204.8T CPO in preliminary roadmap tables); **no official specs announced as of September 2026** [unverified].
- **BlueField-4**: 800 Gb/s throughput (2x BlueField-3), 6x more compute than BlueField-3, 64-core Grace (Arm Neoverse) CPU + ConnectX-9 SuperNIC; "designed to power the operating system of AI factories" [vendor-reported — NVIDIA briefing via Network World, October 2025].
- **BlueField-5**: slated ~2028 per NVIDIA's annual cadence, part of the Feynman architecture; roadmap only [unverified].

### 15.3 Meraki street pricing (reseller spot checks, September 2026 — [unverified] channel prices)

- **MS390-48UX2-HW** (48x mGig up to 5G UPoE, 645W PoE budget, 640 Gbps, 480 Gbps stacking, 476.19 Mpps, 32K MAC table):
  - MSRP **$14,310.31**; authorized channel ~$5,983.14 (hummingbirdnetworks, new, out of stock) [unverified].
  - Gray/secondary market: $789.99 new gray-market (NetworkTigers), $1,386.81 refurbished (GoToDirect), €624 (eonetix, in stock) — not comparable to official channel [unverified].
  - UAE listing: AED 10,449.60 (itechdevices) [unverified].
  - Sources: https://www.networktigers.com/products/ms390-48ux2-hw-cisco-switch-new ; https://www.hummingbirdnetworks.com/meraki-ms390-48ux2-hw-48-port-5gbe-upoe-1x-mod-slot-unclaimed-switch ; https://www.gotodirect.com/ms390-48ux2-hw-cisco-network-switch ; https://www.eonetix.com/shop/MS390-48UX2-HW-Cisco-Meraki-48-port-5GbE-UPoE-8023bt-Switch-Chassis-only ; https://www.itechdevices.ae/cisco-ms390-48ux2-hw-meraki-ms390-48ux2-48-ports-upoe-ge-switch.html
- **MS130-48-HW** (48x 1GbE, 4x 1G SFP, L2 access):
  - MSRP **$4,276.17**; CDW street $2,667.99 (backordered), hummingbirdnetworks $2,437.42 (182-day Cisco lead time), wiretap AU $2,559.81 ex GST, Bulgaria €2,876.76 incl. VAT [unverified].
  - 1-year enterprise license LIC-MS130-48-1Y: MSRP $298, street ~$156–$263 [unverified].
  - Sources: https://www.cdw.com/product/cisco-meraki-ms130-48-switch-48-ports-managed-rack-mountable/8550294 ; https://www.hummingbirdnetworks.com/meraki-ms130-48-cloud-managed-48ge-switch-ms130-48-hw ; https://wiretap.com.au/product/meraki-ms130-48/ ; https://www.directdial.com/us/item/meraki-enterprise-license-and-support/lic-ms130-48-1y
- **MS130-48P-HW** (48x 1GbE PoE+, 740W): MSRP $5,776.38; street ~$2,519–$2,503 volume [unverified]. Finland: €5,676 incl. 25.5% VAT [unverified].
  - Sources: https://www.tech-america.com/item/cisco-meraki-ms130-48p-hw-ethernet-switch/ms130-48p-hw ; https://www.multitronic.fi/en/products/4014012/meraki-ms130-48p-cloud-managed-48ge-740w
- Meraki list-vs-street discount depth (~40–60% off MSRP) is consistent with Cisco Meraki channel practice; license recurring cost (LIC-MS130-48-1Y ~$156 street) adds ~6% of hardware per year.

## Supplementary / Complementary Research Pass — round 4 (2026-09-22)

**Scope note:** the file already contains the base report (§§1–8) plus three supplementary waves (§§A–AG). This pass covers only genuinely uncovered ground: SONiC 4.7 status check, Cumulus EOL-table refresh + 5.19.0 evidence, Meraki MS210 pricing detail, ConnectX-8 street pricing/SKU detail, Aruba CX 6000 NRF 2026 models, Gartner 2026 MQ additional detail, Z9864F-ON spec depth. Nothing existing was modified.

### R4-A. Dell Enterprise SONiC — 4.6.0 remains the latest; no 4.7.0 found

- Dell support KB 000228560 (Minimum/Recommended/Latest code versions, crawl ~Jul 2026) shows **Enterprise SONiC 4.6.0 as the latest release** on both Broadcom (30-June-2026) and NVIDIA Spectrum (SN-4.6.0, 11-June-2026); **recommended remains 4.5.3 (05-June-2026)**. No 4.7.x train was located as of 2026-09-22 — 4.6.0 is the current top-of-train [official].
  Source: https://www.dell.com/support/kbdoc/en-ph/000228560/minimum-recommended-and-latest-code-versions-for-networking-products
- The same KB documents Dell Enterprise SONiC 4.6.0 availability for **EdgeCore white-box models** — AS4630-54PE, AS5835-54T, AS7326-56X, AS7712-32X, AS7726-32X, AS7816-64X, AS9716-32D (latest 4.6.0 released 30-June-2026, recommended 4.5.3; note: EdgeCore products are supported by EdgeCore, Dell supports only the Enterprise SONiC distribution) [official].
- **SN5601** appears in the KB as a Spectrum platform carrying SN-4.6.0 (SN2201, SN4700, SN5600, SN5601) — a model name not previously itemized in this file [official].
- Dell InfoHub white paper **H04658** "Dell Technologies AI Fabrics Overview: NVIDIA Spectrum with Dell SONiC" (March 2026) and PoC technical brief **HO4617** "AI Fabrics NVIDIA Spectrum with Dell SONiC" (March 2026): describe the GPU-cluster back-end fabric, compute multi-tenant front-end fabric, and out-of-band management network using Spectrum switches with Dell SONiC [official].
  Sources: http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_60dba377-fc13-4251-9dd6-c81409b4095d.pdf ; http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_215fe847-2754-45d2-9f66-b54fe3bfbd6a.pdf
- Dell blog (2026): "Open Ethernet for AI: NVIDIA Spectrum-X with Dell SONiC" — RoCEv2 integrated, adaptive routing for congestion avoidance, ECN-marking + PFC congestion management integrated with Spectrum-X telemetry, native in-band telemetry with Grafana integrations, compliant with the NVIDIA Cloud Partner Reference Architecture [official].
  Source: https://www.dell.com/en-us/blog/open-ethernet-for-ai-nvidia-spectrum-x-with-dell-sonic/

