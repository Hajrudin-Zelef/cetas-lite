---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/r5-sources-verbatim-urls
title: "R5 sources (verbatim URLs)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: reference
actors: ["Nvidia", "xAI"]
dates: ["2024-10-28", "2025-08", "2025-10", "2026-07-13", "2026-09", "2026-09-22"]
keywords: ["benchmark", "compute", "cost", "cpo", "ethernet", "gpu", "gpus", "license", "nvidia", "pricing", "research", "rubin"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1869, 1932]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 9fa454f61c4508554b3dba36224db3b5ac9f5b5bc2024209840a3c3737f6b005
---

# R5 sources (verbatim URLs)

| # | Check | Result |
|---|-------|--------|
| R5.1 | C9550 full official specs | Found (cisco.com data sheet + release notes + HW install guide) — R5-A |
| R5.2 | Z9964 radix/two-tier engineering detail | Found (official Dell blog) — R5-B |
| R5.3 | Signal65 independent Z9864F/H200 benchmark | Found — R5-B |
| R5.4 | Meraki MS subscription SKU matrix + tier definitions | Found (official Cisco subscription data sheet; Stratus guide) — R5-C |
| R5.5 | Spectrum-X platform composition / naming caution | Found — R5-D |
| R5.6 | C9550 list/street pricing | **Still not located** — gap remains open (R4 lines 1155/1340) |

### R5 sources (verbatim URLs)

- https://www.cisco.com/c/en/us/products/collateral/networking/switches/c9550-series-smart-switches-ds.html
- https://www.cisco.com/c/en/us/td/docs/switches/lan/c9000/release-notes/c9550-series-smart-switches-release-notes-262x.html
- https://www.cisco.com/c/en/us/td/docs/switches/lan/ciscoc9550/hardware-install/cisco-c9550-series-smart-switches-hig.pdf
- https://www.cisco.com/c/en/us/td/docs/switches/lan/c9000/licensing/cns-licensing-c9000-smart-switches.pdf
- https://www.cisco.com/c/en/us/products/collateral/switches/catalyst-9000-switches/c9550-series-smart-switches-og.pdf
- https://www.dell.com/en-us/blog/power-your-ai-future-how-dell-powerswitch-unlocks-next-generation-ai-network-performance-and-scale/
- https://signal65.com/research/dell-poweredge-xe9680-h200-cluster-with-dell-400gbe-networking/
- https://www.cisco.com/c/en/us/products/collateral/networking/software/networking-subscription-ds.pdf
- https://www.stratusinfosystems.com/cisco-meraki-ms-series-comparison-guide/
- https://www.networkworld.com/article/4200086/nvidia-unveils-spectrum-x-networking-platform-designed-to-connect-millions-of-gpus.html
- https://github.com/hczhu/stock-research/blob/HEAD/memos/2026-07-13-nvidia-hardware-lineup-ai-factory-ecosystem.md

**Round-5 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags. One gap carried forward (C9550 pricing). Earlier sections were not modified.

---

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

