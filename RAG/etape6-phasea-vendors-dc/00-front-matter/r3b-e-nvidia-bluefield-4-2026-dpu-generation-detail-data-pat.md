---
id: etape6-phasea-vendors-dc/00-front-matter/r3b-e-nvidia-bluefield-4-2026-dpu-generation-detail-data-pat
title: "R3b-E. NVIDIA BlueField-4 — 2026 DPU generation detail (data-path networking, Phase A-adjacent)"
domain: front-matter
role: reference
task: actor-profile
actors: ["EU", "Nvidia", "United States"]
dates: ["2025-10", "2026-09", "2026-09-22", "2026-11", "2027-04", "2027-07", "2027-08", "2027-11"]
keywords: ["nvidia", "compute", "consumer", "disclosure", "ethernet", "gpu", "inference", "kv cache", "license", "memory", "optics", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1019, 1097]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: a5f84273b21092d0a5d0d1fed65eb4baf17031bccd02976681190afcea7dc1d1
---

# R3b-E. NVIDIA BlueField-4 — 2026 DPU generation detail (data-path networking, Phase A-adjacent)

### R3b-E. NVIDIA BlueField-4 — 2026 DPU generation detail (data-path networking, Phase A-adjacent)

The base report (§2.5) named BlueField-4 as unveiled at GTC 2026 with few details. New 2026 detail:

- Announced at **GTC DC (October 2025, Washington D.C.)**; **early availability 2026 as part of NVIDIA Vera Rubin platforms**; storage-partner platforms **available second half of 2026** [official-via-press — NVIDIA; secondary — SDxCentral, ServeTheHome].
- Hardware (Hot Chips 2026 disclosure, via ServeTheHome) [secondary]: **64 Arm Neoverse V2 "Grace" CPU cores at 1.7 GHz**, **LPDDR5 bandwidth 275 GB/s**, **ConnectX-9 networking with 800G Ethernet over 200G PAM4 SerDes**, inline encryption, **PCIe Gen6 x16** host link, **126 billion transistors**.
- Performance claims: **6× more compute power than BlueField-3**; enables AI factories **up to 4× larger** than BlueField-3 generation; NVIDIA frames a move from 200 Gb/s-class "cloud DPUs" to **7 Tb/s aggregate bandwidth on the AI DPU**; every Vera Rubin compute tray needs **7 Tb/s aggregate** — four 1.6 Tb/s GPU scale-out links plus **800 Gb/s of scale-in to the DPU** [secondary — SDxCentral, ServeTheHome].
- Software/architecture: **DOCA microservices**, multiservice architecture with native **service function chaining**, **Advanced Secure Trusted Resource Architecture** (zero-trust tenant isolation), software-defined acceleration across AI data storage/networking/security [secondary].
- **KV-cache / inference-context storage play** (CES 2026): NVIDIA Inference Context Memory Storage Platform — Rubin cluster-level KV cache capacity, **up to 5× power efficiency vs traditional storage**, hardware-accelerated KV-cache placement on BlueField-4, DOCA + **NIXL library + Dynamo** integration, RDMA access over **Spectrum-X Ethernet** [vendor-reported — NVIDIA CES 2026 release].
- Storage partners building BlueField-4 platforms: **AIC, Cloudian, DDN, Dell Technologies, HPE, Hitachi Vantara, IBM, Nutanix, Pure Storage, Supermicro, VAST Data, WEKA** [vendor-reported].
- Note: 1.7 GHz clock is deliberately lower-power (vs GB300 Grace cores); only 64 cores active [independent — ServeTheHome analysis].

### R3b-F. Quantum-X800 chassis detail — Q3200-RA vs Q3400-RA design notes

Extends base §2.3 with field-design detail [official — NVIDIA; secondary — gpusmith.com field notes, 2026]:

| Item | Q3400-RA | Q3200-RA |
|---|---|---|
| Form factor | 4U | 2U |
| XDR ports | **144** over 72 OSFP cages | **72** (two fully independent 36-port switches in one chassis) |
| Capacity | **115.2 Tb/s** | **57.6 Tb/s** |
| Role | spine-class, large fabrics | smaller deployments; **bridging legacy HDR/NDR/NDR200/XDR400 into XDR fabric** |

- **Two-level non-blocking fat-tree reference design for the Q3400-RA connects up to 10,368 ConnectX-8 NICs**; GTC 2024 reference: a **576-GPU rail group as a 'virtual modular switch' built from 4 Quantum-X800 switches** in a 3-tier topology [secondary].
- **Not a drop-in upgrade**: cabling, optics, and endpoint NICs (ConnectX-8/9) are a different generation from Quantum-2 gear and do not interoperate at line rate [secondary — field notes].
- Power: **2,900 W typical, up to 7,000 W with active cabling on the Q3400-RA** — chassis power approaches a full GPU compute node; rack power planning must treat the switch as a first-class consumer [secondary].
- **Dedicated management OSFP port** keeps UFM traffic off the data plane (budget a cable/port; not part of the 144-port count) [secondary].
- Fabric software: SHARP v4 hardware in-network computing (**9× in-network computing claim**), adaptive routing, telemetry-based congestion control; 2× bandwidth, 5× data throughput vs prior generation (NVIDIA claims) [official — NVIDIA Quantum-X800 platform page].

### R3b-G. Cumulus Linux 5.18 — GA status and Spectrum-6 support evidence

Extends base §3.3 / Pass #1 §F (EOL table through 5.18):

- **Cumulus Linux 5.18 user guide and 5.18.1 release notes are published on docs.nvidia.com** — 5.18 shipped by September 2026 (the 5.18.1 release-notes page lists fixed issues dated to this train) [official — NVIDIA docs, crawled 2026-09].
- **Spectrum-6 support confirmed in the 5.18 train**: 5.18.1 open/fixed issues explicitly reference Spectrum-6 hardware (e.g., **SN6600-LD** — gNMI shared-buffer cell-size reporting 192 bytes vs correct 256 bytes for Spectrum-6; `bmc-cli collect-all` on Spectrum-6; 5.17 issues also mention "Spectrum-6 switch") [official — NVIDIA 5.18.1 release notes].
- **New 5.18 feature evidence**: **802.1X telemetry server metrics** (new setting; upgrade from pre-5.18 requires a config workaround for a leftover Docker setting) [official].
- Support policy (KB, current): **5.18.z EOL August 2027**; 5.17.z EOL July 2027; 5.16.z April 2027; 5.15.z November 2026 (EOL approaching at research date); LTS: 5.9.z → April 2027, 5.11.z → November 2027 [official — NVIDIA support-policy KB].
- 5.18 upgrade/ops caveats in release notes (NVUE 802.1X telemetry, BGP/BFD FRR reload, maintenance-mode warm-boot behavior, ECMP forwarding corner cases) — relevant for production readiness assessment [official].
- Note: the standalone Cumulus Linux 5.18 **announcement/GA date** was not pinned down in sources searched; the guide exists and the docs GitHub mirror still carries a draft marker for the 5.18 page — treat the "shipped" claim as release-notes-evidenced, not announcement-evidenced [gap].

### R3b-H. Round-3 verification log (new open items)

1. Meraki MS450/MS355 street prices are reseller snapshots (US-centric); EU/APAC price variance not captured [secondary].
2. Essentials vs Advantage per-model price differentials (MS100/200/300/400 tiers) — exact license SKUs per tier not enumerated [gap].
3. Cisco C9350 street pricing / orderability — not located; datasheet only [gap].
4. Cumulus 5.18 exact GA date and release-highlights page — not pinned down [gap].
5. BlueField-4 early-availability ship dates at named Rubin partners — "early availability 2026 / H2 2026" per NVIDIA; partner ship dates not confirmed [unverified].
6. Q3200-RA pricing — not located (all NVIDIA switch pricing quote-only) [gap].
7. C9610 price bands and MS license tier mapping for the 10-slot chassis — not located [gap].

### R3b-I. Round-3 sources (verbatim URLs)

- https://4tekgear.com/meraki-ms450-12-hw-layer-3-switch-hardware-only.html
- https://www.hummingbirdnetworks.com/cisco-meraki-lic-ms450-12-5yr/
- https://www.hummingbirdnetworks.com/cisco-meraki-lic-ms450-12-3yr/
- https://4tekgear.com/meraki-ms355-48x-5-year-hardware-licensing.html
- https://4tekgear.com/meraki-ms355-48x2-1-year-hardware-licensing.html
- https://www.barcodesinc.com/cisco/part-ms355-48x2-hw.htm
- https://www.stratusinfosystems.com/cisco-meraki-ms-series-comparison-guide/
- https://www.cisco.com/c/en/us/products/collateral/networking/software/networking-subscription-ds.pdf
- https://www.crn.com/news/networking/2025/cisco-networking-smart-switches-and-wi-fi-gear-unveiled-at-cisco-live-2025
- https://www.techpowerup.com/337907/cisco-unveils-secure-network-architecture-new-smart-switches-secure-routers-and-wifi7-access-points
- https://www.cisco.com/c/dam/en/us/products/collateral/switches/c9350-series-smart-switches/c9350-series-smart-switches-ds.pdf
- https://www.cisco.com/c/dam/m/cs_cz/training-events/webinars/tech-club-webinars/tech-club-switching-update-december-2025.pdf
- https://documentation.meraki.com/@api/deki/pages/11886/pdf/Cloud%2bConfiguration%253A%2bRelease%2bVersions%2band%2bHighlights.pdf?stylesheet=default
- https://documentation.meraki.com/@api/deki/pages/10591/pdf/How%2bto%2bUpgrade%2bCloud-Managed%2bCatalyst%2bSwitches%2bto%2bIOS%2bXE.pdf?stylesheet=default
- https://www.servethehome.com/nvidia-bluefield-4-processor-at-hot-chips-2026/
- https://www.sdxcentral.com/analysis/nvidias-bluefield-4-a-first-look-at-the-dpu-built-to-run-ai-factories/
- https://www.servethehome.com/nvidia-bluefield-4-with-64-arm-cores-and-800g-networking-announced-for-2026/
- https://investingnews.com/nvidia-bluefield-4-powers-new-class-of-ai-native-storage-infrastructure-for-the-next
- http://gpusmith.com/datasheets/nvidia-quantum-x800-datasheet.pdf
- https://www.nvidia.com/en-gb/networking/products/infiniband/quantum-x800/
- http://docs.nvidia.com/networking-ethernet-software/cumulus-linux/Whats-New/rn/
- https://docs.nvidia.com/networking-ethernet-software/knowledge-base/Support/Support-Offerings/Cumulus-Linux-Release-Versioning-and-Support-Policy/

**Round-3 collection metadata:** read-only web research (browser_search + one page fetch, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §R3b-H above. Existing sections were not modified.

---

