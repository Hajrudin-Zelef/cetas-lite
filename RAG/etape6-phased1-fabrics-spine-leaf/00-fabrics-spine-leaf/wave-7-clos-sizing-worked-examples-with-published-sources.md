---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-7-clos-sizing-worked-examples-with-published-sources
title: "Wave 7 — Clos sizing worked examples with published sources"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Broadcom"]
dates: ["2025-06", "2026-03"]
keywords: ["agents", "cost", "governance", "memory", "throughput"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [204, 267]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 1181efd1fd6425d4dd3730e71be82dc409ab43f8ea8d769387eb191368b0b8ac
---

# Wave 7 — Clos sizing worked examples with published sources

## Wave 7 — Clos sizing worked examples with published sources

### 7.1 Juniper's published oversubscription options (32×100G leaf)

Juniper's white paper "Design Considerations for Spine-and-Leaf IP Fabrics" gives concrete per-leaf splits for a 32×100G leaf [official — https://Www.juniper.net/content/dam/www/assets/white-papers/us/en/design-considerations-for-spine-and-leaf-ip-fabrics.pdf]:

| Fabric-facing | Workload-facing | Oversubscription | Notes |
|---|---|---|---|
| 16×100G | 16×100G | 1:1 | Non-blocking |
| 12×100G | 20×100G | 1:1.67 | 20/12 |
| 8×100G | 24×100G | 1:3 | Cost-optimized |
| 8×100G | 24×40G (24 ports at 40G workloads) | 1:1.2 | Mixed speed |
| 8×100G | 24×25G workloads | 1:1 | 600G workload vs 800G fabric |

Key Juniper design rules from the same paper [official]: it is not possible to go below 1:1 (workloads cannot consume more than fabric capacity); not every leaf needs the same oversubscription as long as every leaf has the same number of spine-facing ports; oversubscription is "usually considered safe" when application profiles are known to load well below port speeds or traffic is cyclical and non-overlapping; hot spots (database, backup servers) are the classic failure mode, causing drops, delay and jitter that are "generally difficult to diagnose and resolve."

### 7.2 Two-tier scale ceilings from FS.com's design guide

FS.com's leaf-spine design guide gives worked scale figures [secondary — https://www.fs.com/sg/blog/what-is-spineleaf-architecture-and-how-to-design-it-2927.html]:
- Rule: maximum leaf count = spine port density; spine count governed by required inter-leaf throughput, ECMP path count, and port density.
- L2 vs L3: L2 designs give VLAN/MAC mobility anywhere; L3 designs give fastest convergence and fan-out ECMP "supporting up to 32 or more active spine switches."
- 3:1 or less is "the most appropriate oversubscription ratio for modern network architectures."

FS.com's 25G/100G design article adds three-tier numbers [secondary — https://www.fs.com/blog/fs-25g-portfolio-for-data-center-25g100g-leafspine-network-13082.html]:
- A 1.5:1 leaf-uplink design can support up to ~13,000 servers (288 leaves × 48 servers); beyond ~20,000 servers a multi-tier architecture becomes necessary.
- Three-tier (super-spine) with 1.5:1 at leaf and 3:1 at spine, uplinks upgradeable to 400G+: 64 pods × 2,304 servers = 140,000+ servers, with fault-domain isolation between pods.
- Hyperscale "fabric plane" Clos variant: each server pod fully meshed to multiple spine "planes" of parallel non-blocking switches, ECMP across planes.

### 7.3 Community sizing rules (cross-checked)

A community data-center fabric skill document corroborates the arithmetic [secondary — https://github.com/chrishuffman5/agents/blob/HEAD/skills/networking/dc-fabric/SKILL.md]: max leaves = spine ports; oversubscription = total leaf downlink BW ÷ total leaf uplink BW; target 3:1 or better general, 1:1 storage/HPC; worked examples — 48×10G down + 6×100G up = 480:600 = 0.8:1 (non-blocking); 48×25G down + 6×100G up = 1200:600 = 2:1.

MTU rule for VXLAN fabrics from the same source [secondary]: VXLAN adds 50 bytes overhead; all fabric links must support jumbo frames — minimum fabric MTU 9214 bytes; ACI auto-configures fabric MTU via APIC; NSX/open EVPN require manual underlay MTU setting.

### 7.4 Update to Wave 1.3 (Tomahawk 6 — now confirmed shipping)

The [gap] in Wave 1.3 is now partially closed: Broadcom announced Tomahawk 6 shipments in June 2025 [official — https://www.globenewswire.com/news-release/2025/06/03/3092820/19933/en/Broadcom-Ships-Tomahawk-6-World-s-First-102-4-Tbps-Switch.html], and third-party reporting indicates production-volume shipments began March 2026 [secondary — treat date as indicative]. Full system-level detail (Edgecore, DriveNets) is in Wave 13.

---

## Wave 8 — Cisco deep-dive: verified scalability numbers and the 2026 ACI-vs-EVPN debate

### 8.1 NX-OS mode verified scalability (VXLAN/EVPN)

Cisco's Nexus 9000 NX-OS Verified Scalability Guides publish per-release tested limits [official]:
- Release 10.4(3)F and later: **1,000 VTEP scale** supported, with a minimum 32 GB memory requirement on the device [official — https://www.Cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/104x/configuration/scalability/cisco-nexus-9000-series-nx-os-verified-scalability-guide-1046.pdf].
- Older reference (Release 10.1(2), iBGP-centric VXLAN/EVPN topology): 128 VTEPs, 2,000 L2 VNIs, 500 L3 VNIs/VRFs, 64,000 overlay MAC addresses, 60,000 IPv4 host routes, 16,000 IPv6 host routes verified on Nexus 9200/9300/9500 and X9700-EX/FX line cards [official — https://WWW.Cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/101x/configuration/scalability/cisco-nexus-9000-series-nx-os-verified-scalability-guide-1012.pdf]. **Do not compare these directly** with the 10.4 figures — different releases, different test topologies.
- CloudSec scale: 128 hardware security associations = M×N×L (peers × uplinks × border gateways); not supported on 9348GC-FX3/FX3PH, 9332D-H2R, 93108TC-FX3 [official].

### 8.2 ACI verified scalability (APIC 5.2(7)/5.2(8)/5.3.x)

From the Cisco APIC Verified Scalability Guides [official — https://www.cisco.com/c/en/us/td/docs/dcn/aci/apic/5x/verified-scalability/cisco-aci-verified-scalability-guide-527.html and https://www.cisco.com/c/en/us/td/docs/dcn/aci/apic/5x/verified-scalability/cisco-aci-verified-scalability-guide-528.html]:
- Multi-pod / remote leaf: **50 remote-leaf pairs (100 RLs total)** per fabric.
- VRFs: 1,200; external EPGs: 2,000 total (100 per VRF).
- EVPN sessions (SR-MPLS context): 4 per leaf, 100 per fabric; ECMP paths 16 per leaf.
- vPC: 64 total ports, 32 per leaf pair.
- First-hop security: 2,000 endpoints / 1,000 bridge domains per leaf.

### 8.3 The 2026 industry debate: NX-OS EVPN gaining on ACI

A widely discussed 2026 practitioner piece argues more data-center teams are choosing NX-OS VXLAN EVPN over ACI for fresh platform decisions [secondary — https://dev.to/firstpasslab/why-more-data-center-teams-are-choosing-nx-os-vxlan-evpn-over-cisco-aci-in-2026-529j]. Its comparison dimensions: day-1 provisioning (ACI strong greenfield), day-2 troubleshooting (EVPN easier with standard tools), policy model (ACI powerful but proprietary), automation (EVPN fits broader IaC/multivendor workflows), hiring (larger transferable EVPN skill base), fabric portability (EVPN standards-oriented). It concedes ACI still wins for teams already operationally invested, with mature segmentation and centralized governance. **Assessment:** directionally consistent with the broader 2024–2026 industry shift toward open EVPN, but a single practitioner source — treat as [secondary], not consensus.

---

