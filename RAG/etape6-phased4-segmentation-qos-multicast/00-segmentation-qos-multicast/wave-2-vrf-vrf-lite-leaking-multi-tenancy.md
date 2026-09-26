---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy
title: "Wave 2 — VRF: VRF-lite, leaking, multi-tenancy"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Huawei"]
dates: ["2026-09-22"]
keywords: ["agent", "agents", "containment", "cost", "datacenter", "ethernet", "license"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [46, 91]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 20243254db02ead8f60fda935881128f67f4ba1d0090a057bd7383a8d86bb0d9
---

# Wave 2 — VRF: VRF-lite, leaking, multi-tenancy

## Wave 2 — VRF: VRF-lite, leaking, multi-tenancy

### 2.1 Definitions
- VRF (Virtual Routing and Forwarding) = isolated L3 routing table instance per tenant/function; VRF-lite = VRF without MPLS in the data plane (per-interface VRF assignment on a shared box) `[secondary]` — netlab docs; Cisco Community. https://github.com/ipspace/netlab/blob/HEAD/docs/module/vrf.md
- In EVPN-VXLAN fabrics the modern equivalent is per-tenant VRF + L3 VNI with symmetric IRB; strict RT separation is what prevents leakage `[secondary]` — NetPilot multi-tenant EVPN example.

### 2.2 Route leaking mechanics and limits
- Inter-VRF route leaking imports/exports prefixes between VRFs, typically via BGP route-targets or static/route-map filters; commonly used to reach shared services (DNS, DHCP, AD, internet egress) `[official]` — Cisco NX-OS L3 virtualization guide (Nexus 3548, 10.5x): leaking supported between any two non-default VRFs and between default VRF and any other VRF; route leaking TO the default VRF is not allowed (it is the global VRF); max 1000 leaked prefixes by default (configurable 0–1000); requires Enterprise license + BGP `[official]`. https://www.Cisco.com/c/en/us/td/docs/dcn/nx-os/nexus3548/105x/configuration/unicast/b-cisco-nexus-3548-switch-nx-os-unicast-routing-configuration-guide-105x/m-configuring-layer3-virtualization.pdf
- Classic IOS-style config uses per-VRF `route-target import/export` with route-maps for Global→VRF leaking `[secondary]` — Cisco Community example (2013-era, still structurally representative).
- In Cisco ACI, the equivalent of VRF leaking is configuring a subnet as "shared"; alternatives are using the common tenant's VRF or hanging shared services off outside routers `[official]` — Cisco Press ACI multi-tenancy chapter. https://www.ciscopress.com/articles/article.asp?p=2928191&seqNum=4
- Arista EOS documents EVPN multicast VRF leaking (EOS 4.36.2F, Sept 2026): egress PE sends SMET to receiver VRF; with PIM EVPN Gateway the PEG acts as RP for leaked groups; MSDP sync between source/receiver VRFs may be required; DF/non-DF VTEP config must match for multi-homed Ethernet segments `[official]`. https://www.arista.com/en/um-eos/eos-evpn-multicast-vrf-leaking?searchword=eos%20evpn

### 2.3 Platform support snapshot (2026-09)
- netlab VRF module docs: VRF config supported on Arista EOS, Aruba AOS-CX, Cisco IOS/IOS XE/IOS XR/NX-OS, Cumulus 4.x/5.x, Dell OS10, FRR, Junos, MikroTik RouterOS 6/7, Nokia SR Linux/SR OS, VyOS; route leaking supported on all except Nokia SR OS (supported in OS but not yet in netlab) — with caveats for Cumulus 5.x NVUE, Junos cRPD, FRR, SR Linux `[secondary]` — ipspace netlab (updated 6 days before 2026-09-22). https://github.com/ipspace/netlab/blob/HEAD/docs/module/vrf.md
- Routing protocols in VRF (netlab matrix): eBGP widely supported; OSPFv2 mostly; OSPFv3 absent on NX-OS/Cumulus/Dell OS10; IS-IS on Arista/Junos/IOS XR/Nokia/SR Linux/FRR; IBGP within a VRF does not work — PE and CE must use different ASNs `[secondary]` — same source. Treat as lab-tool matrix, not vendor datasheet — cross-check before design use.

### 2.4 Multi-tenancy patterns
- Pattern A — strict isolation: unique RT per tenant VRF, no leaking; verified hosts in different tenants cannot reach each other `[secondary]` — NetPilot.
- Pattern B — shared services: dedicated shared-services VRF (DNS/DHCP/SIEM/internet), selective leaking from tenant VRFs `[secondary]` — NetPilot; Cisco ACI "common tenant" equivalent `[official]`.
- Pattern C — internet egress: single shared Internet VRF on border leaf only, reached via RT leaking (documented live with DSVNI/asymmetric IRB evidence on Nexus 9000v, NX-OS 10.5.3, CML lab) `[secondary]` — enizaksoy VXLAN EVPN multi-AS lab (updated 2026-07). https://github.com/enizaksoy/cisco-vxlan-evpn-multi-as-lab
- Gap: independent measurements of route-leaking scale limits (prefix counts, convergence) on current DC platforms were not found — open item `[unverified]`.

### Wave 2 verification
- Sources: 6. One notable per-platform limit captured verbatim (Cisco 1000-prefix cap). Conflict note: Huawei/Cisco/Juniper anycast-RP docs differ on MSDP necessity — covered in Wave 7, not a VRF conflict. No internal contradictions.

---

## Wave 3 — Microsegmentation & policy

### 3.1 Enforcement taxonomy
- Five enforcement categories (structure of the market, 2026): network vendors (Cisco ACI/TrustSec-SGT, Arista, Juniper, Nokia) enforce in the fabric; firewall vendors (Palo Alto, Fortinet, Check Point) enforce at zones and increasingly at the workload; hypervisor/overlay (VMware NSX archetype) enforces at the vSwitch; host-agent specialists (Illumio, Akamai Guardicore) are platform-independent with best east-west visibility at the cost of per-host agents; cloud-native controls (security groups, NACLs, K8s NetworkPolicy) are already present but often left at defaults `[secondary]` — ronutz/arsenal segmentation guide (updated 2026-09). https://github.com/ronutz/arsenal/blob/HEAD/src/content/learn/en/network-segmentation-from-vlans-to-microsegmentation.mdx
- The buying question per that analysis: where you can enforce consistently across everything you own — a segmentation program with a large enforcement gap relocates risk rather than removing it `[secondary]` — same source.

### 3.2 Measurable benefits
- Blast-radius reduction (compromised host reaches only what policy allows — the same argument as ZTNA at a different layer), ransomware containment (most encryption campaigns depend on broad east-west reachability), compliance scoping (shrinking audit scope often funds the project), and denied-flow telemetry as high-quality detection signal `[secondary]` — same source.

### 3.3 VMware NSX Distributed Firewall (representative overlay model)
- DFW is embedded in the hypervisor kernel on each host; security rules apply at each VM's vNIC with near-maximum performance; logical switches isolated by default; policies move with the workload on vMotion including session state `[vendor-reported]` — VMware NSX micro-segmentation docs; Medium/TechTarget summaries (2025). https://www.vmware.com/docs/vmware-nsx-microsegmentation · https://www.techtarget.com/it-infrastructure/feature/NSX-distributed-firewall-ensures-security-across-VMs-containers
- Eliminates hair-pinning east-west traffic to a central firewall; supports L3/L4 and L7 App-ID firewalling plus distributed IDS/IPS in the hypervisor `[vendor-reported]` — TechTarget.
- ACI-vs-NSX design tension (NetCraftsmen, older but structurally current): when both are deployed, enforce policy where the traffic actually flows — VM-to-VM traffic inside NSX never exits the overlay, so ACI/physical firewalls cannot enforce on it; ACI then automates the fabric and enforces for physical endpoints `[secondary]`. https://netcraftsmen.com/designing-the-datacenter-aci-nsx-or-both/

### 3.4 ACLs and policy-based routing (fabric-level)
- Fabric ACLs (port/VLAN/routed ACLs on ToR/leaf) remain the coarse segmentation tool in EVPN-VXLAN fabrics: VACLs filter intra-VLAN, PACLs at the port, RACLs at SVIs; scale is TCAM-bound and varies by platform/generation — record per-platform, never assume `[secondary]` — standard design references; exact TCAM numbers are platform-specific and were not enumerated here (gap).
- Policy-based routing: overrides destination-based forwarding using ACL/class-map matching (source, app, DSCP) to steer traffic to next-hops or VRFs; DC use cases include steering backup/replication flows, forcing inspection via firewall VRF, and tenant-aware egress `[secondary]` — widely documented pattern; vendor CLIs differ (Cisco route-map PBR, Arista policy-based routing, Junos filter-based forwarding).
- Gap: no independently verified current comparison of ACL scale (entries) across 2026 DC switch ASICs was collected — open item `[unverified]`.

