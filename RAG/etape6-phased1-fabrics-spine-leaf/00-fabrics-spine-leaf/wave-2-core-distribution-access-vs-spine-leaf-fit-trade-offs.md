---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-2-core-distribution-access-vs-spine-leaf-fit-trade-offs
title: "Wave 2 — Core/distribution/access vs spine-leaf: fit, trade-offs, migration"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Meta", "Microsoft", "Nvidia", "xAI"]
dates: ["2020-04"]
keywords: ["distribution", "acquisition", "alignment", "benchmarks", "cost", "dci", "ethernet", "gpu", "latency", "nvidia", "pricing", "research"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [69, 157]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 27d5bcea8303fc3432e61f2f73897e2a8085cc68f50aaf3d827f79d08754f1ed
---

# Wave 2 — Core/distribution/access vs spine-leaf: fit, trade-offs, migration

## Wave 2 — Core/distribution/access vs spine-leaf: fit, trade-offs, migration

### 2.1 Where classic three-tier still fits

The three-tier model (access → distribution/aggregation → core) remains rational where [secondary]:
- North-south traffic dominates (client-to-server, internet edge) rather than east-west (server-to-server); classic campus and small enterprise DCs qualify.
- Chassis-based cores already exist and are amortized; forklift replacement is hard to justify below ~500–1,000 server ports.
- Oversubscription is naturally high (20:1+ at access→distribution is normal in campus) and acceptable.

Spine-leaf wins where east-west dominates (virtualization, big data, AI, microservices), latency must be predictable, and scale-out growth is expected [secondary].

### 2.2 Migration paths operators actually use

Common migration patterns [secondary]:
- **Greenfield pod alongside legacy:** build a leaf-spine pod for new workloads, interconnect to the legacy core via border leafs; migrate VLANs/VRFs progressively. Lowest risk; standard enterprise approach.
- **Distribution-to-leaf conversion:** in smaller fabrics, existing distribution switches become spines and access switches become leaves, with an underlay (eBGP) overlaid — feasible when port counts and NOS support it (e.g., Aruba CX, SONiC) [secondary].
- **ACI-style forklift:** Cisco ACI migrations are typically rip-and-replace per pod with APIC-managed fabrics; Cisco publishes migration guides from NX-OS mode to ACI mode on the same Nexus 9000 hardware (certain models support both) [official].
- **Brownfield EVPN-VXLAN overlay:** keep existing L2/L3 devices, add VXLAN VTEPs at the new leaf layer; the overlay abstracts the underlay so legacy and new fabric interoperate during transition [secondary].

**[gap]** — vendor-published cost comparisons of migration vs greenfield at specific port counts were not found; TCO claims in vendor collateral are [vendor-reported] and workload-specific.

---

## Wave 3 — Vendor reference designs

### 3.1 Cisco: Nexus (NX-OS) and ACI

Cisco's data-center portfolio centers on Nexus 9000 series switches with two operating modes [official]:
- **NX-OS mode:** traditional CLI, supports VXLAN/EVPN fabrics with BGP underlays; the basis of Cisco's "VXLAN EVPN multi-site" reference designs [official].
- **ACI mode:** Application Centric Infrastructure — spine-leaf fabric managed by APIC controllers with a policy model (tenants, EPGs, contracts) rather than per-box CLI; uses VXLAN with a Cisco-proprietary control plane (OpFlex/COOP) [official].

Reference scale figures published by Cisco for ACI (verify against current release notes; figures move by APIC/NX-OS version): historically up to ~200 leaf switches per fabric in large designs [vendor-reported]. Cisco's validated design guides (CVDs) document 2-tier and 3-tier (with super-spine) ACI topologies [official].

Current hardware context (2026): Nexus 9000 lines include 400G (e.g., 9332D-H2R class) and 800G-capable platforms; Cisco's Silicon One-based 800G portfolio (e.g., Cisco 8122 series noted in Phase C research) targets AI fabrics [vendor-reported]. Cisco acquired Splunk (2024) and its ThousandEyes/W ThousandEyes integration feeds ACI/NDO observability story [secondary].

**Positioning note:** ACI is the most opinionated fabric (controller-mandatory, policy model); NX-OS EVPN is the open-protocols alternative on the same hardware. Choice between them is the single biggest Cisco DC design decision [independent].

### 3.2 Arista: EOS and CloudVision

Arista's architecture is EOS (single binary image across all platforms) with CloudVision as the management/automation plane [official]. Reference designs are published as "Validated Designs" (AVD — Arista Validated Designs), an Ansible-based framework that generates complete leaf-spine configurations from a data model [official].

Key fabric elements [official][vendor-reported]:
- eBGP underlay with unnumbered interfaces (RFC 5549) as the documented default in AVD designs.
- EVPN-VXLAN overlay for multi-tenancy; Arista documents L2 and L3 EVPN services extensively.
- CloudVision provides config management, telemetry (LANZ — Latency Analyzer), and change control; CloudVision as-a-service is the SaaS option [official].
- Hardware: 7800R4 series with 800G ports (576×800G density cited in Phase C research), 7060/7280R3/7388 lines for leaf/spine roles [vendor-reported].

Arista's design guidance emphasizes large flat L3 fabrics for cloud and AI; its "AI centers" reference architectures pair Spectrum-X-class Ethernet (NVIDIA) or Arista 800G fabrics with rail-optimized topologies for GPU backends [vendor-reported]. **[gap]** — independent head-to-head fabric benchmarks Arista vs alternatives at 800G were not found this wave.

### 3.3 NVIDIA: Spectrum/Spectrum-X, Cumulus Linux, SONiC

NVIDIA's networking stack after the Mellanox acquisition (completed April 2020, ~$6.9B) [secondary]:
- **Spectrum** Ethernet switches (Spectrum-4 at 51.2T class) for general data-center fabrics [vendor-reported].
- **Spectrum-X** — Ethernet platform co-designed for AI: Spectrum-X switches + BlueField DPUs + DOCA software, with adaptive routing and congestion control (telemetry-driven) tuned for collective-communication (NCCL) traffic patterns; positioned against InfiniBand for AI training fabrics [vendor-reported].
- **Cumulus Linux** — NVIDIA's Linux-based NOS (acquired 2020), supporting VXLAN/EVPN fabrics with FRRouting; positioned for open-networking shops [official].
- **SONiC** support on Spectrum platforms for hyperscale-aligned operators [official].

Reference designs: NVIDIA publishes "Spectrum-X" validated designs for AI pods (rail-optimized backend + frontend Ethernet fabric), and Cumulus reference topologies for enterprise leaf-spine with BGP unnumbered underlays [official][vendor-reported]. Quantum-X800 InfiniBand (noted in Phase A research) remains the InfiniBand counterpart for the largest training clusters [vendor-reported].

**Design note:** the NVIDIA-recommended AI data center is two fabrics: frontend (Ethernet, storage/management) and backend (Spectrum-X or InfiniBand, GPU collectives) — a pattern Meta, xAI and others have publicly mirrored in broad terms [secondary].

### 3.4 Dell: Enterprise SONiC and SmartFabric OS10

Dell offers two NOS tracks on its PowerSwitch (Z/N/S series) hardware [official]:
- **Enterprise SONiC** — Dell-hardened distribution of community SONiC with Dell support; supports BGP-EVPN/VXLAN fabrics; Dell publishes validated "fabric design" guides for 2-tier leaf-spine with Enterprise SONiC [official].
- **SmartFabric OS10** — Dell's own NOS with SmartFabric Services (SFS): automated fabric provisioning (plug-and-play leaf discovery, automated underlay) aimed at midmarket/private cloud; supports full automation via SFS orchestrator [official].

Positioning: OS10/SFS for automated "hands-off" fabrics; Enterprise SONiC for shops wanting community-SONiC alignment with vendor support [secondary]. Dell's 800G-class PowerSwitch models (e.g., Z9664F-RT class noted in Phase A) serve spine/super-spine roles [vendor-reported]. **[gap]** — current Dell list pricing for Enterprise SONiC licensing tiers was not collected this wave.

### 3.5 HPE Aruba: CX switching for data center

Aruba CX (CX 10000, 9300, 8325/8360 series) with AOS-CX supports data-center leaf-spine with EVPN-VXLAN overlays and VSX (Virtual Switching Extension) for multi-chassis LAG at the leaf/aggregation layer [official]. Aruba's data-center design guides document 2-tier fabrics with VSX pairs as "distributed" leaves and EVPN for multi-tenancy [official].

HPE's acquisition of Juniper (completed mid-2025, ~$14B) [secondary] reshapes this portfolio: Juniper Apstra + QFX and Aruba CX/Fabric Composer now sit under one roof; integration roadmaps were still being published through 2026 [secondary]. **[unverified]** — long-term NOS consolidation outcomes (AOS-CX vs Junos vs Apstra-managed) remain officially uncommitted as of this writing.

### 3.6 Juniper: Apstra and QFX

Juniper's data-center story is intent-based: **Apstra** (acquired 2021) manages QFX/PTX fabrics from declared intent, with built-in validation, telemetry, and support for SONiC as well as Junos devices [official]. Reference blueprints cover 2-stage and 5-stage Clos (collapsed, 3-stage, 5-stage templates are Apstra's standard vocabulary) with EVPN-VXLAN overlays [official].

Hardware: QFX5120/5130/5220/5230/5240/5210 lines span leaf to spine; PTX for super-spine/DCI roles [official]. Juniper publishes "data center fabric" validated designs with eBGP underlays and documented oversubscription options per blueprint [official].

### 3.7 Whitebox / community SONiC designs

Community SONiC (Software for Open Networking in the Cloud, originally Microsoft) runs on whitebox/brite-box switches (Accton/EdgeCore, Celestica, etc.) and on vendor platforms (Dell, NVIDIA, Arista via AVD-adjacent tooling) [secondary]. Reference fabric designs are community-maintained: the classic Microsoft Azure SONiC deployment (tens of thousands of switches, multi-tier Clos) is the existence proof at hyperscale [secondary].

Design tooling: SONiC fabrics are typically provisioned with Ansible/Nornir + Jinja2 templates or vendor wrappers; Batfish/SuzieQ-style validation (see Phase E) is common in the community [secondary]. Trade-off vs vendor NOS: maximum openness and cost leverage against self-supported integration burden [independent].

---

