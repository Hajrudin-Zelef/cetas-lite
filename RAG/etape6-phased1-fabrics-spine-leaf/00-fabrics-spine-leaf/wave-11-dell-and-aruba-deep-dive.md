---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-11-dell-and-aruba-deep-dive
title: "Wave 11 — Dell and Aruba deep-dive"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["compute", "lean", "omni"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [319, 370]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 5b70dcec87d33a8c1d2ae7687416c67129ffe4dba8501d89260ffc06bfb4ecd3
---

# Wave 11 — Dell and Aruba deep-dive

## Wave 11 — Dell and Aruba deep-dive

### 11.1 Dell SmartFabric Services (SFS) personalities

Per the Dell SFS User Guide (Release 10.5.5) [official — https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/sfs-solution_ug-10-5-5/SFS-personalities?guid=guid-474f43b5-2c48-460e-8a41-5ffc25d7893a&lang=en-us]:
- **L2 single-rack personality is deprecated**; Dell recommends all new deployments use the **L3 leaf-spine fabric personality**.
- L3 personality: up to 20 switches in a leaf-spine design, expandable up to 8 racks; available on SmartFabric OS10.5.0.5+.
- Existing L2-personality deployments upgrading to 10.5.0.5 stay on L2; no automatic migration to L3 personality.
- SFS automates fabric bring-up (plug-and-play leaf discovery, automated underlay); managed via CLI/API/UI or OMNI.

### 11.2 Dell OS10 BGP unnumbered (RFC 5549)

OS10 documents BGP unnumbered with ENHE/RFC 5549 compliance [official — https://www.dell.com/support/manuals/en-us/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-1/BGP-over-unnumbered-interfaces?guid=guid-fe393409-2b87-4ad8-9324-d73f94707030&lang=en-us]: unnumbered peers form neighborship over IPv6 link-local via router advertisements; restrictions — no IPv6 VRRP on unnumbered interfaces, no route reflectors, confederations, or dampening with unnumbered peers; a `link-local-only-nexthop` knob exists for interop with FRR-based implementations.

### 11.3 Aruba validated DCN architectures

Aruba publishes Ansible-automated DCN reference architectures on its developer portal [official — https://developer.arubanetworks.com/aoscx/docs/architecture-iii-dedicated-data-center-layer-3-spineleaf-topology-ebgp-evpn-multi-as-vxlan-with-vsx]:
- **Architecture III**: dedicated DC L3 spine/leaf, eBGP EVPN multi-AS, VXLAN with VSX; spines in a VSX pair must support EVPN (8325/8400 required); leaves in VSX pairs must support EVPN/VXLAN (8325); playbook `deploy_ebgp_evpn_vxlan.yml` with Jinja2 templates and Excel/YAML inventory.
- **Aruba Fabric Composer**: out-of-band orchestration for leaf-spine provisioning on CX 8360/8325 underlay switches; integrations with vSphere, Nutanix, iLO Amplifier [secondary — https://packetpushers.net/blog/aruba-fabric-composer-the-data-center-simplified/].
- Hardware data points: CX 9300 — 25.6 Tbps, 32×400G in 1RU, VSX redundancy [official via datasheet mirror — https://andovercg.com/datasheets/hpe-aruba-CX-9300-Series.pdf]; CX 8360-32Y4C v2 (JL700C): 32×25G SFP28 + 4×100G QSFP28, 2.4 Tbps, AOS-CX with VSX [secondary].
- NOS gotcha for automation authors: AOS-CX ports are **L3 by default** (inverse of NX-OS); `no routing` opts into L2; `active-gateway` (not VRRP) is the anycast gateway surface [secondary — https://github.com/netcanon/netcanon/blob/HEAD/docs/vendors/aruba_aoscx.md].

---

## Wave 12 — Juniper deep-dive: 5-stage EVPN-VXLAN JVD with Apstra

### 12.1 Validated design under test

Juniper's 5-stage EVPN-VXLAN data-center JVD (JVD-DCFABRIC-5STAGE-01-01, doc dated 29-Jan-25) validates a 5-stage Clos, ERB (edge-routed bridging) architecture with Apstra [official — https://www.juniper.net/documentation/us/en/software/jvd/jvd-dcfabric-5-stage/jvd-dcfabric-5-stage.pdf and https://github.com/juniper/jvd/blob/HEAD/data_center/adc/5stage_evpn_vxlan/documentation/design-guide.md]:

- Software under test: **Apstra 5.0.0-64**, **Junos OS 23.4R2** (S3/S4 service releases).
- Validated functionality: 5-stage Clos + ERB; single- and multi-homed servers (ESI LAG with LACP); ECMP; **eBGP for both underlay and overlay**; EVPN Type 2 + Type 5; **BFD on underlay and overlay eBGP**; symmetric IRB with anycast IP on L3 leaves; IPv4+IPv6; OISM multicast (BDNE); ECN/PFC QoS profiles for RoCEv2; DHCP, duplicate-MAC detection; inter-VRF via external router.
- "Lean" spine/super-spine concept: spines and super-spines do IP forwarding + route relay only.
- 5-stage adopted for large datastore/compute designs needing RoCEv2 + multicast at scale [official].

### 12.2 Platform roles in the JVD

| Role | Validated platforms |
|---|---|
| Server leaf | QFX5120-48Y-8C, QFX5120-48YM, QFX5130-32CD (EVO), QFX5110-48S |
| Border leaf | QFX5700 (EVO), QFX5130-48C (EVO), PTX10001-36MR |
| Spine | QFX5220-32CD (EVO), QFX5210-64C, QFX5120-32C |
| Super-spine | QFX5230-64CD (EVO) |

Juniper notes same-chipset variants are covered by one tested model (exceptions: QFX5130-48C vs 32CD) [official]. A collapsed-fabric JVD also exists (spine+leaf+border roles on one tier, edge DCs) [official — https://github.com/juniper/jvd/blob/HEAD/data_center/adc/collapsed_dc_fabric_with_access/documentation/design-guide.md].

### 12.3 Apstra design workflow

Apstra's model: **rack types** (logical device links, redundancy like MLAG/ESI, links-per-spine) → **templates** (3-stage, 5-stage, or collapsed; IP fabric or BGP EVPN) → **blueprints** (instantiated with IP pools, ASN pools, interface maps) → virtual networks/VRFs/connectivity templates [official]. Apstra 5.x manages Junos, Junos Evolved, vQFX, plus Cisco NX-OS, Arista EOS, and Enterprise SONiC devices [secondary]. Cabling maps support anti-affinity policies.

---

