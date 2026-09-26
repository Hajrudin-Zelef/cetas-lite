---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/16-addendum-meraki-cloud-management-ios-xe-17-18-x-release-t
title: "16. Addendum — Meraki Cloud Management IOS XE 17.18.x release timeline, Aruba CX 10040 street pricing, Broadcom Trident5-X12 silicon context, Dell SFM for SONiC official spec (September 22, 2026)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: pricing
actors: ["Broadcom"]
dates: ["2025-12", "2026-07", "2026-09-22"]
keywords: ["pricing", "distribution", "ethernet", "license", "nvidia"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [2041, 2094]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 8d463dc2197fceb1e8e114746b760198b6882de1a43743495c06143a68d30596
---

# 16. Addendum — Meraki Cloud Management IOS XE 17.18.x release timeline, Aruba CX 10040 street pricing, Broadcom Trident5-X12 silicon context, Dell SFM for SONiC official spec (September 22, 2026)

- https://www.dell.com/support/kbdoc/en-ph/000228560/minimum-recommended-and-latest-code-versions-for-networking-products
- http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_60dba377-fc13-4251-9dd6-c81409b4095d.pdf
- http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_215fe847-2754-45d2-9f66-b54fe3bfbd6a.pdf
- https://www.dell.com/en-us/blog/open-ethernet-for-ai-nvidia-spectrum-x-with-dell-sonic/
- https://docs.nvidia.com/networking-ethernet-software/knowledge-base/Support/Support-Offerings/Cumulus-Linux-Release-Versioning-and-Support-Policy/
- https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-513/Whats-New/rn/
- https://github.com/nvidia/product-security/blob/HEAD/2026/5817/5817.md
- https://blog.ipspace.net/2025/06/cumulus-linux-gone/
- https://www.shi.com/product/34444595/CISCO-MERAKI-MS210-
- https://www.publicsector.shidirect.com/Product/34444595/CISCO-MERAKI-MS210-48FP-1G-L2CLD-MNGD-48X-GIGE-740W-POE-SWITCH
- https://www.cablesandkits.com/networking/switches/meraki-switches/ms210-48fp-hw/pro-27668/
- https://www.shi.com/product/34444596/CISCO-MERAKI-MS210-
- https://www.shi.com/product/34444597/CISCO-MERAKI-MS210-48LP-1G-L2CLD-MNGD-48X-GIGE-370W-POE-SWITCH
- https://networkequipment.net/products/cisco-meraki-ms210-24-hw-new
- https://www.FS.com/c/nvidia-ethernet-nics-4014
- https://www.fs.com/uk/c/nvidia-ethernet-nics-4014
- https://www.naddod.com/collections/nvidia-networking/infiniband-adapters
- https://www.accio.com/plp/connectx-8-c8240
- https://nvdam.widen.net/content/8h0owe2dhm/original/connectx-datasheet-connectx-8-supernic-update-a4-web-zhCN-3523588-R3.pdf?u=rubmrs&use=c8xan&download=true
- https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html
- https://www.networktigers.com/products/r8n85a-hpe-switch-new
- https://www.hpe.com/emea_europe/en/networking/magic-quadrant-wired-wireless.html
- https://www.businesswire.com/news/home/20260520437739/en/Arista-Networks-Positioned-as-a-Leader-in-the-2026-Gartner-Magic-Quadrant-for-Enterprise-Wired-and-Wireless-LAN
- https://www.webdisclosure.com/press-release/huawei-etr-huawei-named-a-leader-in-the-2026-gartner-magic-quadrant-for-enterprise-wired-and-wireless-lan-infrastructure-for-the-fourth-year-in-a-row-p7AObeR7kp0
- https://www.grabnpay.in/products/dell-powerswitch-z9864f-on-64-800gbe-osfp112-high-density-open-networking-switch-with-dual-ac-dc-power-supplies-and-hot-swappable-fans
- https://delltechnologies.com/asset/ko-kr/products/networking/technical-support/dell-powerswitch-z9864f-on-spec-sheet.pdf
- https://bigfastservers.com/products/dell-powerswitch-z9864f-on-with-102-4-tbps-switch-capacity-64x-800-gbe-osfp112-2x-sfp-ports
- https://juaraitsolutions.com/dell-powerswitch-z-series-switches/
---

## 16. Addendum — Meraki Cloud Management IOS XE 17.18.x release timeline, Aruba CX 10040 street pricing, Broadcom Trident5-X12 silicon context, Dell SFM for SONiC official spec (September 22, 2026)

**Scope note:** this pass adds ONLY material not already present. Cross-checks before writing: `Trident5` = 0 occurrences (genuine gap; Tomahawk-5/6 merchant silicon is already covered in §§1.1/1.5, R4-B), `17.18.2` = 1 occurrence (partial mention only, §X line 1014), `SmartFabric Manager` = 7 occurrences (sections 1.4, §P, round-4 notes — but no official feature spec; GA-vs-roadmap status gap was flagged in §Q), `10040` = 13 occurrences (DPU-enabled series described, no street pricing). Earlier sections untouched.

### 16.1 Cisco Meraki Cloud Management with IOS XE 17.18.x — official release timeline and Cloud EVPN fabric technical detail

- Per the official Meraki documentation page "Cloud Management with IOS XE Overview" [official], **IOS XE 17.18.2 is a release candidate** (surfaced under the dashboard's "Other Available Versions" tab); 17.18.1 introduced BGP, VRF, and ISSU support plus full Catalyst 9200/9300/C9500 High Performance family coverage; 17.18.2 adds **cloud EVPN fabric capabilities, VRRP support, and SmartPort Automation (early access)**.
  - Source: https://documentation.meraki.com/@api/deki/pages/10595/pdf/Cloud%2bManagement%2bwith%2bIOS%2bXE%2bOverview.pdf?stylesheet=default
- Per the official "Cloud Managed EVPN Fabric Technical Guide" (documentation.meraki.com, doc updated ~41 days before retrieval) [official]: Cloud-managed EVPN fabric requires **IOS XE 17.18.2 or later**, a **Cloud Switching Advanced license** (for EVPN and Adaptive Policy), Cisco ISE 3.2+ (or Meraki Access Manager) for authentication, OSPFv2 unicast underlay, and an existing Layer 3 underlay topology.
  - Source: https://documentation.meraki.com/Switching/Cloud_Management_with_IOS_XE/Design_and_Configure/Cloud_Managed_EVPN_Fabric_Technical_Guide
- Release-support matrix from the same guide [official]: **IOS XE 17.18.2 shipped December 2025**; **17.18.3 / 26.1.1 in public beta as of July 2026**. Access: MS + Catalyst 9200/9300; distribution leaf: Catalyst 9300X (17.18.2); core spine and network edge/border: Catalyst 9500 High Performance (recommended) or Catalyst 9300X. 17.18.3/26.1.1 adds EVPN Multi-Homing at the leaf layer.
- Per the official "BGP Routing for Cloud Management with IOS XE" doc [official]: BGP route-table scale (IPv4 total / direct / indirect / IPv6 / multicast):
  - **C9300X: 39,000 / 24,000 / 15,000 / 19,500 / 8,000**;
  - **C9300 / MS390: 32,000 / 24,000 / 8,000 / 16,000 / 8,000**;
  - C9300L/LM: 32,000 / 24,000 / 8,000 / 16,000 / 8,000; C9500H: 90,000 / 90,000 / 90,000 / 90,000 / 32,000; C9200CX: 14,000 / 10,000 / 4,000 / 2,000 / 1,000.
  - This confirms the **MS390 as a cloud-managed BGP leaf** at the same hardware scale as the Catalyst 9300 access tier — relevant to Meraki DC campus-leaf positioning.
  - Source: https://documentation.meraki.com/@api/deki/pages/11882/pdf/BGP%2bRouting%2bfor%2bCloud%2bManagement%2bwith%2bIOS%2bXE.pdf?stylesheet=default
- Phased BGP feature support [official]: **Phase 1 (17.18.x)** — IPv4 unicast, default VRF, prefix lists, AS-path, route reflectors; **Phase 2 (future)** — IPv6 unicast, L2VPN EVPN, multicast, multi-VRF, route maps, community strings.
- Independent cross-check [secondary — community-maintained knowledge base, not official Cisco]: a Meraki dashboard support heuristic table maps VLAN/STP/DHCP snooping/LACP/QoS/SVI as native since 17.15.1; BGP/VRF/VRRP/ISSU native since 17.18.1; cloud EVPN fabric and SmartPort Automation native since 17.18.2; OSPF only partial; NetFlow/IPFIX not supported on cloud-managed mode.
  - Source: https://github.com/kebaldwi/meraki-rag/blob/HEAD/rag/knowledge_base/04_dashboard_support_heuristic.md
- Competitive read-through [independent]: Cisco's cloud-EVPN-on-dashboard push (GA path from Dec 2025 RC to 2026 public beta) is the clearest enterprise vendor counterweight to Dell SmartFabric Manager for SONiC intent-fabric automation and Aruba Fabric Composer — all three converge on dashboard/Blueprint-driven EVPN in 2026.

### 16.2 HPE Aruba CX 10040 — street pricing and license SKUs (first price points in this file)

