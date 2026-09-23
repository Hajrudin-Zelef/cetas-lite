---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-9-arista-deep-dive-800g-hardware-and-avd-data-model
title: "Wave 9 — Arista deep-dive: 800G hardware and AVD data model"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: hardware
actors: ["Nvidia"]
dates: ["2023-05", "2025-11"]
keywords: ["dci", "ethernet", "full-duplex", "gpu", "lpo", "nvidia", "optics", "research"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [268, 318]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 9c27c5270283fb765d39204c28a68eaa2fbb30f4cf292cd83e7a1f11414efef0
---

# Wave 9 — Arista deep-dive: 800G hardware and AVD data model

## Wave 9 — Arista deep-dive: 800G hardware and AVD data model

### 9.1 R4-series 800G hardware (announced Nov 2024, shipping status 2025–2026)

From Arista's press release and datasheets [official/vendor-reported — https://investors.arista.com/Communications/Press-Releases-and-Events/Press-Release-Detail/2025/Arista-Networks-Unveils-Next-Generation-Data-and-AI-Centers/default.aspx and http://sit.www.arista.com/assets/data/pdf/Datasheets/7800R4-Series-AI-Spine-Datasheet.pdf]:

| Platform | Role | Key specs |
|---|---|---|
| 7800R4 modular (4/8/12/16-slot) | AI/DC spine, super-spine | Up to 576×800G per system; 36-port 800G OSFP line cards (72×400G); 28.8 Tbps forwarding + 10.8 Bpps per LC; 32 GB buffer per LC; cell-based fabric (non-hash), active/active fabric modules; ~460 Tbps system capacity class (7816LR4: 460 Tbps / 920 Tbps full-duplex per third-party summary [secondary]) |
| 7280R4 fixed | Spine / backbone | 32×800G; also 64×100G + 10×800G leaf variant |
| 7020R4 fixed | DC/AI leaf | 10/25G server ports, 100G uplinks, per-port TunnelSec wire-speed encryption |
| HyperPort | Scale-across/DCI | Single 3.2 Tbps Ethernet interface; claimed 44% shorter AI job completion vs 4×800G LAG [vendor-reported] |

Availability per Arista: 7800R4 systems and line cards shipping at announcement; 7280R4 platforms shipping; 7020R4 chassis platforms targeted Q1 2026 [vendor-reported]. November 2025 update adds lossless 800G, LPO optics support, 800G ZR/ZR+ and post-quantum cryptography on R4 lines [official — https://arista.com/en/22897-800g-r4-launch-webinar].

### 9.2 AVD (Arista Validated Designs) data model specifics

AVD is an Ansible collection (arista.avd) generating full EOS configs, docs, and CloudVision deployment from a YAML data model [official — https://github.com/gusmb/ansible-avd and https://github.com/lermilov/ansible-avd/blob/HEAD/ansible_collections/arista/avd/README.md]. Key fabric variables [official — https://github.com/cbcrc/ansible-avd/blob/HEAD/ansible_collections/arista/avd/roles/eos_designs/doc/fabric-variables.md]:

- Underlay protocols: EBGP (default for l3ls-evpn), OSPF, ISIS (incl. SR/LDP variants); overlay: EBGP (default) or IBGP.
- `underlay_rfc5549: true/false` (default false) — point-to-point underlay with RFC 5549 IPv6 unnumbered; requires EBGP underlay.
- `p2p_uplinks_mtu` default 9000; `bgp_maximum_paths` / `bgp_ecmp` default 4.
- Supported design matrix: eBGP/RFC5549(eBGP)/ISIS/OSPF underlays × eBGP/iBGP overlays, 3-stage and 5-stage (multi-stage) topologies + L2 leafs.
- Example bundles render complete configs: single-dc-l3ls (6 leaves + 2 spines; leaf config ~365 lines with MLAG, VXLAN, EVPN, 3 VRFs, 12 VLANs, virtual-router anycast MAC; spine ~138 lines with BGP AS 65100); dual-dc-l3ls (12 leaves + 4 spines, explicit DCI); single-dc-l3ls-ipv6 (IPv6 underlay variant) [secondary — https://github.com/netcanon/netcanon/blob/HEAD/docs/fixture-research-2015/03-arista_eos.md].
- AVD 3.x release notes (current ~Sept 2026) show active EVPN/multicast feature work [official — https://github.com/aristanetworks/avd/blob/HEAD/docs/release-notes/3.x.x.md].

---

## Wave 10 — NVIDIA deep-dive: Spectrum-X launch facts and Cumulus reference design

### 10.1 Spectrum-X launch (May 2023, Computex)

NVIDIA launched Spectrum-X as an "accelerated Ethernet platform for hyperscale generative AI" [official — https://nvidianews.nvidia.com/news/nvidia-launches-accelerated-ethernet-platform-for-hyperscale-generative-ai]:
- Spectrum-4: described at launch as the world's first 51.2 Tb/s Ethernet switch built for AI; 256×200G ports per switch; **16,000 ports in a two-tier leaf-spine topology**.
- End-to-end 400GbE: Spectrum-4 + BlueField-3 DPUs + LinkX optics with "advanced RoCE extensions."
- Software: Cumulus Linux, pure SONiC, NetQ; DOCA framework on BlueField.
- Testbed/blueprint: "Israel-1" supercomputer on Dell PowerEdge XE9680 (HGX H100 8-GPU), BlueField-3 DPUs, Spectrum-4 switches.
- Ecosystem at launch: Dell Technologies, Lenovo, Supermicro.

### 10.2 Cumulus reference design AS scheme

The Cumulus Linux Network Reference Design Guide documents the canonical Clos AS plan [official — https://onix.kiev.ua/pdf/mellanox/Cumulus-Linux-Network-Reference-Design-Guide.pdf]: a unique AS per leaf; a **shared AS per spine group** (all spines serving the same leaves); a shared AS per super-spine group. Peering between tiers; routing policy shares loopbacks and **disallows leaf-as-transit**; load balancing enabled. Spines act as EVPN route forwarders without installing forwarding state (non-VTEP). Cumulus supports EVPN with both eBGP and iBGP peering.

### 10.3 Cumulus BGP unnumbered (NVUE example)

Current Cumulus Linux 5.14 docs show BGP unnumbered config via NVUE: `nv set vrf default router bgp neighbor swp51 remote-as external` — neighbor as interface, no IP on the link [official — https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-514/Layer-3/Border-Gateway-Protocol-BGP/Basic-BGP-Configuration/]. Reference EVPN examples use MLAG leaf pairs with BGP unnumbered underlay [official — https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-44/Network-Virtualization/Ethernet-Virtual-Private-Network-EVPN/Configuration-Examples/].

**Interop caveat:** a community lab verified live that Arista EOS/cEOS does NOT accept `remote-as external` for unnumbered peers — explicit per-peer AS required on EOS [secondary — https://github.com/alukacs03/clauntainerlab/blob/HEAD/labs/28-bgp-unnumbered/README.md]. Dell OS10 documents its own unnumbered restrictions (next wave).

---

