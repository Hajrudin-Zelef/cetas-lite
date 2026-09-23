---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/overview
title: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Nvidia"]
dates: ["2026-05", "2026-07-13", "2026-07-17", "2026-09-22"]
keywords: ["research", "amd", "asic", "benchmark", "blackwell", "distribution", "dram", "ethernet", "fp8", "gpu", "gpus", "latency"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1823, 1895]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 063ccf092f0f6c949f6b99e017ecd579b0fa392fce0d12718cd5b974389e076a
---

# Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail

> **Scope.** Single-pass addition on top of Rounds 1–4. Every item below was verified absent or less-detailed in the current file (re-grepped 2026-09-22: `C9550`, `Z9964`, `BlueField-5`, `LIC-MS-*`). Earlier sections were not modified. All facts carry provenance tags; open items are logged in the R5 verification log at the end of this section.

## R5-A. Cisco C9550 Series Fixed Core Smart Switches — full datasheet specifications (NEW detail)

Round 4 (R4-C, lines 1119–1122) logged the C9550 as orderable since May 2026 with 4×400G uplinks and 6.4 Tbps, from a secondary Cisco-Live recap. The **official Cisco data sheet and release notes** (crawled 49–82 days before 2026-09-22) now supply complete specifications:

- **Three models, two ASIC tiers** [official — Cisco C9550 data sheet]:
  - **C9550-96L4D** ("E100 model"): 1× Cisco **Silicon One E100** ASIC; up to **6.4 Tbps** system switching, **3.9 Bpps** forwarding; AMD x86 3.3 GHz 8-core CPU; 32 GB DRAM; up to **960 GB SSD**; 2×10G AppGig ports; app-hosting allocation up to 8 GB DRAM / 4 vCPUs. Scale: 128,000 MAC, **1,000,000 IPv4 routes**, 500,000 IPv6 routes, 128K ARP/NDP entries, 32K multicast routes, 16K IGMP/MLD snooping entries, 5K QoS ACL, 21K/21K security ACL (ingress/egress), 32K/32K NetFlow entries, 64 MB packet buffer, 4094 VLANs, 4000 SVIs, jumbo 9216.
  - **C9550-48L4CD and C9550-24L4CD** ("E104 models"): 1× Cisco **Silicon One E104** ASIC; up to **3.2 Tbps**, **2.6 Bpps**; AMD x86 3.0 GHz 4-core CPU; 16 GB DRAM; up to 960 GB SSD; app-hosting up to 4 GB DRAM / 2 vCPUs. Scale: 64,000 MAC, **512,000 IPv4 routes**, 256,000 IPv6 routes, 8K IGMP/MLD snooping, 16K multicast routes (same ACL/NetFlow/VLAN numbers as E100 model).
- **Port layouts** [official — IOS XE 26.2.x release notes]:
  - C9550-24L4CD: 24× 50/25/10/1G downlinks + 4× 100/40G or 2× 400G fixed uplinks; 2 PSU slots; 5 fan trays.
  - C9550-48L4CD: 48× QSFP-DD 50/25/10/1G downlinks + 4× 100/40G or 2× 400G fixed uplinks; 2 PSU slots; 5 fan trays.
  - C9550-96L4D: 96× QSFP-DD 50/25/10/1G downlinks + 4× QSFP-DD 40G/100G/400G; 2 PSU slots; 3 fan trays; 2RU; default PSU C9K-PWR-1100WAC (1100W AC).
- **Stacking and storage:** "next-generation Cisco Smart Stacking with StackWise Virtual (SVL) front-side stacking on all ports"; 960 GB SATA SSD local storage for container-based application hosting [official — data sheet].
- **Positioning:** designed primarily for the core and distribution layers of medium-sized to large enterprise campus networks; also fabric border or aggregation for medium/large fabric networks; serves as a consolidation point aggregating multiple 50/25/10/1G SFP access switches into a 400/100/40G QSFP backbone [official — data sheet use-case text].
- **Software:** introductory release **Cisco IOS XE 26.2.1ea** (all three models) [official — release notes].
- **Physical:** 1U models 1.73×17.5×18.15 in, 20.13–20.79 lb (9.15–9.45 kg) with 2 PSUs; 2U model 3.47×17.5×17.85 in, 33.55 lb (15.25 kg). Sound power ≤87.2 LWAd (24L), ≤85.7 (48L), ≤89.3 (96L) [official — hardware installation guide, updated 2026-07-17].
- **Management/licensing:** a dedicated Cisco licensing doc exists — "Cisco Networking Subscription for Cisco C9000 Series Smart Switches" (cns-licensing-c9000-smart-switches.pdf) — and Cisco publishes a "Switching Licensing Feature Matrix"; no list/street prices for C9550 hardware or its subscription tiers located [gap — R4 gap lines 1155/1340 still open].

## R5-B. Dell PowerSwitch Z9964/Z9864 — official radix and two-tier engineering detail (NEW detail)

Base §1.1 (lines 32–42) records the Z9964F-ON/Z9964FL-ON November-2025 launch (Tomahawk-6, 102.4 Tbps, 64×1.6 TbE, 100k+ accelerators, air/DLC). Dell's official AI-networking blog adds the engineering rationale [official — Dell blog, "Power Your AI Future"]:

- **Tomahawk-5 reference (Z9864F-ON):** radix 64 at 800GE, 128 at 400GE, 256 at 200GE → two-tier fabrics supporting 2,048 endpoints @800GE, 8,192 @400GE, 32,768 @200GE.
- **Tomahawk-6 (Z9964F-ON, "Ultra Ethernet compliant"):** radix 64 at **1.6TE**, 128 at 800GE, 256 at 400GE, **512 at 200GE** → two-tier fabrics supporting **2,048 endpoints @1.6TE**, 8,192 @800GE, 32,768 @400GE, **131,072 (128K) endpoints @200GE**.
- **Claimed two-tier benefit:** replacing three-tier networks eliminates one switch layer → **66.7% fewer physical switches/racks**, **40% savings on optics**, lower power/cooling/carbon footprint; the Z9964FL-ON variant uses direct liquid cooling (DLC) for dense AI/ML deployments [vendor-reported — Dell blog].
- **Independent datapoint:** Signal65 benchmarked a Dell PowerEdge XE9680 H200 cluster (64 NVIDIA H200 GPUs, 8 nodes) with Broadcom BCM57608 Thor2 NICs and Dell PowerSwitch Z9864F switches: ~1,979 TOPS (INT8) / 1,979 TFLOPS (FP8) per GPU and **97.3% network efficiency under peak AI workloads**, attributed to the lossless Ethernet fabric with hardware-accelerated collectives and congestion-resilient flow control [independent — Signal65 research].

## R5-C. Meraki MS licensing — official SKU matrix and tier definitions (NEW detail)

Earlier rounds captured MS street prices and the MS150 Enterprise/Advanced tier naming. Cisco's official "Networking Subscription Data Sheet" (©2026) provides the current tiered structure [official — Cisco subscription data sheet]:

- **MS tiers by model size:** LIC-MS-100, LIC-MS-200, LIC-MS-300, LIC-MS-400; each in **S/M/L** sizes with **-A (Advantage)** or **-E (Essentials)** suffixes (e.g. LIC-MS-100-L-E, LIC-MS-300-M-A).
- **Tier split:** **Essentials** = core L2/L3 functionality; **Advantage** = Adaptive Policy, advanced analytics, SD-Branch capabilities. Terms available: 1, 3, 5, 7, and 10 years; tiers cannot be mixed within one organization [secondary — Stratus Information Systems comparison guide].
- **Legacy naming:** the MS150 still uses the older **Enterprise / Advanced** tier names with the same 1/3/5/7/10-year terms, also non-mixable [secondary — reseller listings].
- **Firmware floor:** Cisco switching subscription minimum = IOS-XE 17.18.1 (wireless 17.15.2, SD-WAN 20.15.3); no Meraki firmware requirement for cloud-managed tiers [official — Cisco subscription data sheet].

## R5-D. NVIDIA networking — Spectrum-X platform composition and roadmap naming note (NEW detail)

- **Spectrum-X platform components** (per the Jul-2026 Network World launch article): Spectrum Ethernet switches, Spectrum-X SuperNICs, ConnectX NICs, BlueField DPUs, **LinkX cabling and transceivers**, and Spectrum-XGS for multi-DC networking; Spectrum-6 (102.4 Tbps) is positioned as part of the Vera Rubin platform with silicon photonics to improve bandwidth while lowering latency and power [vendor-reported via press briefings — Network World].
- **Naming caution:** NVIDIA's official Rubin table reportedly calls the 51.2 Tb/s Grace-Blackwell-generation switch **Spectrum-4**; unofficial references to "Spectrum5" for that generation appear to be a naming discrepancy (Spectrum-6 at 102.4 Tb/s is confirmed; Spectrum-7 is future roadmap; ConnectX-10 is future roadmap; BlueField-5 slated for 2028 with Feynman) [unverified — third-party research memo on GitHub; do not treat as official naming].

### R5 verification log (2026-09-22)

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

