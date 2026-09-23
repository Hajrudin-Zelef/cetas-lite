---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-28-source-index-key-urls-consulted
title: "Wave 28 — Source index (key URLs consulted)"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Broadcom", "Meta", "Nvidia", "xAI"]
dates: []
keywords: ["benchmark", "cpo", "datacenter", "dci", "disaggregated", "distribution", "ethernet", "nvidia", "optics", "training"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [629, 696]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: d5ae0ecd19fec50b44211faeecf489c85e372759c5e113001b4cba66da250df6
---

# Wave 28 — Source index (key URLs consulted)

## Wave 28 — Source index (key URLs consulted)

- Juniper spine-leaf white paper: https://Www.juniper.net/content/dam/www/assets/white-papers/us/en/design-considerations-for-spine-and-leaf-ip-fabrics.pdf
- FS.com spine-leaf guide: https://www.fs.com/sg/blog/what-is-spineleaf-architecture-and-how-to-design-it-2927.html
- FS.com 25G/100G design: https://www.fs.com/blog/fs-25g-portfolio-for-data-center-25g100g-leafspine-network-13082.html
- Cisco NX-OS scalability 10.4(6): https://www.Cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/104x/configuration/scalability/cisco-nexus-9000-series-nx-os-verified-scalability-guide-1046.pdf
- Cisco APIC scalability 5.2(7): https://www.cisco.com/c/en/us/td/docs/dcn/aci/apic/5x/verified-scalability/cisco-aci-verified-scalability-guide-527.html
- Cisco APIC scalability 5.2(8)/5.3.x: https://www.cisco.com/c/en/us/td/docs/dcn/aci/apic/5x/verified-scalability/cisco-aci-verified-scalability-guide-528.html
- ACI vs NX-OS EVPN 2026 debate: https://dev.to/firstpasslab/why-more-data-center-teams-are-choosing-nx-os-vxlan-evpn-over-cisco-aci-in-2026-529j
- Arista 7800R4 datasheet: http://sit.www.arista.com/assets/data/pdf/Datasheets/7800R4-Series-AI-Spine-Datasheet.pdf
- Arista R4 launch PR: https://investors.arista.com/Communications/Press-Releases-and-Events/Press-Release-Detail/2025/Arista-Networks-Unveils-Next-Generation-Data-and-AI-Centers/default.aspx
- Arista Nov 2025 R4 update: https://arista.com/en/22897-800g-r4-launch-webinar
- Arista EVPN Gateway whitepaper: http://sit.www.arista.com/assets/data/pdf/Whitepapers/EVPN-Data-Center-EVPN-Gateway-for-Hierarchical-Multi-Domain-EVPN-and-DCI-WP.pdf
- AVD fabric variables: https://github.com/cbcrc/ansible-avd/blob/HEAD/ansible_collections/arista/avd/roles/eos_designs/doc/fabric-variables.md
- AVD release notes 3.x: https://github.com/aristanetworks/avd/blob/HEAD/docs/release-notes/3.x.x.md
- AVD dual-DC example: https://github.com/aristanetworks/avd/blob/HEAD/ansible_collections/arista/avd/examples/dual-dc-l3ls/README.md
- Arista ATD lab (overlay peering): https://github.com/aristanetworks/atd-public/blob/HEAD/topologies/dual-datacenter/labguides/source/l2l3evpn.rst
- NVIDIA Spectrum-X launch: https://nvidianews.nvidia.com/news/nvidia-launches-accelerated-ethernet-platform-for-hyperscale-generative-ai
- Cumulus reference design guide: https://onix.kiev.ua/pdf/mellanox/Cumulus-Linux-Network-Reference-Design-Guide.pdf
- Cumulus BGP unnumbered: https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-514/Layer-3/Border-Gateway-Protocol-BGP/Basic-BGP-Configuration/
- Dell SFS personalities: https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/sfs-solution_ug-10-5-5/SFS-personalities?guid=guid-474f43b5-2c48-460e-8a41-5ffc25d7893a&lang=en-us
- Dell OS10 BGP unnumbered: https://www.dell.com/support/manuals/en-us/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-1/BGP-over-unnumbered-interfaces?guid=guid-fe393409-2b87-4ad8-9324-d73f94707030&lang=en-us
- Aruba Architecture III: https://developer.arubanetworks.com/aoscx/docs/architecture-iii-dedicated-data-center-layer-3-spineleaf-topology-ebgp-evpn-multi-as-vxlan-with-vsx
- Aruba Fabric Composer: https://packetpushers.net/blog/aruba-fabric-composer-the-data-center-simplified/
- Juniper 5-stage JVD: https://www.juniper.net/documentation/us/en/software/jvd/jvd-dcfabric-5-stage/jvd-dcfabric-5-stage.pdf
- Juniper collapsed fabric JVD: https://github.com/juniper/jvd/blob/HEAD/data_center/adc/collapsed_dc_fabric_with_access/documentation/design-guide.md
- Broadcom Tomahawk 6: https://www.globenewswire.com/news-release/2025/06/03/3092820/19933/en/Broadcom-Ships-Tomahawk-6-World-s-First-102-4-Tbps-Switch.html
- Edgecore 102.4T: https://markets.financialcontent.com/wral/article/bizwire-2026-2-23-edgecore-networks-sets-new-benchmark-for-ai-infrastructure-with-worlds-first-1024t-open-networking-switches
- Davisson CPO: https://www.storagenewsletter.com/2025/10/10/broadcom-shipping-tomahawk-6-davisson-102-4-tb-s-ethernet-switch-with-co-packaged-optics/
- Meta RoCE AI zones: https://engineering.fb.com/2024/08/05/data-center-engineering/roce-network-distributed-ai-training-at-scale/
- Meta DSF: https://engineering.fb.com/2025/10/20/data-center-engineering/disaggregated-scheduled-fabric-scaling-metas-ai-journey/
- NVIDIA xAI Colossus: http://nvidianews.nvidia.com/news/spectrum-x-ethernet-networking-xai-colossus
- Edgecore whitebox TCO: https://www.edge-core.com/download/whitebox-switching-data-center-tco-analysis/?wpdmdl=6365&refresh=6aa0b13dd3e8b1788916029
- Cisco Silicon One G300: https://www.sdxcentral.com/news/next-gen-switches-servers-everything-unveiled-at-cisco-live-emea-2026/
- CVE-2026-20212: https://forkast.news/cisco-nexus-9000-silicon-one-rce-exposes-ai-data-center-fabric-to-root-compromise/
- Arista/Pure NVMe-oF: https://www.purestorage.com/content/dam/pdf/en/white-papers/wp-arista-pure-deploying-nvme-of-enterprise-leaf-spine-architecture.pdf

*End of Phase D1 (28 waves). Append-only; earlier sections untouched. Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

## Wave 29 — Brownfield migration: decision matrix and phase plan

### 29.1 Migration pattern decision matrix

| Starting point | Target pattern | Risk | Downtime | Notes |
|---|---|---|---|---|
| 3-tier L2/STP | Parallel EVPN fabric + migrate | Low | Zero (per-VLAN cutover) | Build new leaf-spine beside old core; move VLANs one by one [secondary] |
| 3-tier L3 | In-place leaf reuse as spines | Medium | Maintenance windows | Old core/distribution become spine tier if port density allows [secondary] |
| Single-vendor fabric | Multivendor pod | Medium | Zero (pod-by-pod) | eBGP/EVPN interop; validate ESI and BUM replication across vendors [secondary] |
| VXLAN flood-and-learn | EVPN control plane | Low-Medium | Rolling | EVPN and flood-and-learn can coexist per-VNI during transition [secondary] |
| ACI fabric | NX-OS EVPN | High | Rebuild | No in-place path; 2026 trend drivers documented (Wave 16) [secondary] |

### 29.2 Phase plan (typical enterprise)

1. **Discover:** cable/asset inventory, traffic matrix (NetFlow/sFlow), oversubscription reality check vs design.
2. **Design:** ASN plan (RFC 7938), IP plan (numbered /31s or unnumbered), multi-homing choice, ECMP paths, BFD timers, anycast gateway.
3. **Build parallel:** rack, cable (Wave 20 deterministic plan), baseline config from AVD/Apstra/NetBox templates.
4. **Validate:** LLDP neighbor matrix, BGP session count = expected, ECMP spread test, BUM replication check, loss/PFC counters clean.
5. **Migrate:** one leaf pair or VLAN at a time; keep rollback = move gateway back.
6. **Decommission:** only after 30 days of clean counters on the new fabric.

[secondary — standard field practice, consistent with vendor migration guides]

*End of Phase D1 (29 waves, 750+ lines). Append-only; earlier sections untouched. Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

