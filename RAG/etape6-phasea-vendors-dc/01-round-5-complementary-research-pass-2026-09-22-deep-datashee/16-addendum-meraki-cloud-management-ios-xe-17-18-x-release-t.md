---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/16-addendum-meraki-cloud-management-ios-xe-17-18-x-release-t
title: "16. Addendum — Meraki Cloud Management IOS XE 17.18.x release timeline, Aruba CX 10040 street pricing, Broadcom Trident5-X12 silicon context, Dell SFM for SONiC official spec (September 22, 2026)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: pricing
actors: ["Broadcom", "United States"]
dates: ["2023-11-30", "2025-12", "2026-07", "2026-09-22"]
keywords: ["pricing", "distribution", "ethernet", "gpu", "inference", "inference engine", "latency", "license", "research", "serdes", "throughput"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [2071, 2139]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 8a4e7a617dc658cb4f487c9e2847c02ba398b8191004f3a110b85db56ae582fc
---

# 16. Addendum — Meraki Cloud Management IOS XE 17.18.x release timeline, Aruba CX 10040 street pricing, Broadcom Trident5-X12 silicon context, Dell SFM for SONiC official spec (September 22, 2026)

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

- CX 10040 bundle **S4R54A** (32× 100G QSFP28 + 6× 400G QSFP-DD, front-to-back airflow, rack-mountable): MSRP **$112,995.00**, street **$85,804.00** (SHI, listed as updated <1 hour before retrieval) [secondary — SHI].
  - Source: https://www.shi.com/product/50519753/HPE-Aruba-Networking-CX-10040
- Same series in UK channel [secondary — Back to the Office]: from **£85,025.99** (spring sale, 0 stock) for the S4R54A bundle; TAA-compliant back-to-front variant from **£82,907.99** (January sale, 0 stock) — snapshot prices only, not directly comparable to US MSRP.
  - Sources: https://www.backtotheoffice.co.uk/products/hpe-aruba-networking-cx-10040-32p-qsfp28-100g-6p-qsfp-dd-400g-front-to-back-4xfan-2xpsu-ac-bundle ; https://www.backtotheoffice.co.uk/products/hpe-aruba-networking-cx-10040-32p-qsfp28-100g-6p-qsfp-dd-400g-back-to-front-4xfan-2xpsu-ac-taa-bdl
- License SKUs from the official CX 10040 QuickSpecs (reseller-mirrored PDF) [official — via reseller PDF]: CX 10040 Advanced 1/3/5-year and perpetual E-STU (S6J68AAE–S6J71AAE); CX 10040 Premium 1/3/5-year and perpetual (S6J72AAE–S6J75AAE); IPSec Encryption/NAT for CX 10000 E-LTU (R9H26AAE); Fabric Composer Device Management Service Tier 4 switch 1/3/5-year subscriptions (R7G99AAE/R7H00AAE/R7H01AAE).
  - Source: https://b2b.atlantis.cz/attachments/S4R55A/HPE%20Networking%20CX%2010040%20Switch%20QuickSpecs.pdf
- HPE Store official support pricing [official]: **HPE Aruba Networking Foundational Care 1Y 4HR Onsite for CX 10040 (SKU H08HLE)** — $2,987.78 buy-now price.
  - Source: https://buy.hpe.com/us/en/services-support/technology-services/hardware-software-combo-support-service/hpe-aruba-networking-foundational-care-1y-4hr-onsite-technical-support-cx-10040-32c-service/p/h08hle
- Note: US list pricing is quote-gated on the HPE Store (buy.hpe.com shows "Indicative price / request a custom quote") [official]; the SHI MSRP is the first public list reference located for the CX 10040 in this research.

### 16.3 Broadcom Trident5-X12 (StrataXGS BCM78800) — new merchant-silicon context for leaf/ToR class

- The **Trident 5-X12 (BCM78800)** is Broadcom's software-programmable top-of-rack/leaf merchant chip: **16.0 Tbps** bandwidth (2× Trident 4-X9), **5 nm process**, **25% less power per 400G port** vs Trident 4-X9, with **800G port support** via 100G PAM4 SerDes — positioning it as the ToR partner to Tomahawk 5-based spine/fabric tiers [vendor-reported — Broadcom press release 2023-11-30 via Nasdaq].
  - Source: https://www.nasdaq.com/press-release/broadcom-introduces-industrys-first-switch-with-on-chip-neural-network-2023-11-30
- Industry-first claim [vendor-reported]: first Ethernet switch chip with an **on-chip neural-network inference engine — NetGNT** (Networking General-purpose Neural-network Traffic-analyzer), running in parallel with the packet pipeline to detect traffic patterns (e.g., AI incast congestion) and invoke congestion-control at line rate without throughput/latency impact.
- Port configs: designed to enable a **1RU ToR with 48× 200G downlink + 8× 800G uplink** (e.g., 24× 400G + 8× 800G mixes); 160× 100G-PAM4 SerDes with up to 4 m reach; supports Enterprise SONiC and SAI for data-center integration; field-upgradable via NPL (Network Programming Language), compatible with the Trident 4 family [vendor-reported — EDN, allaboutcircuits, The Register coverage].
  - Sources: https://www.edn.com/ethernet-switch-has-on-chip-neural-network/ ; https://www.allaboutcircuits.com/news/in-an-industry-first-broadcom-puts-neural-network-onto-a-switch/ ; https://www.theregister.com/2023/11/30/broadcom_trident_npu/?td=keepreading ; https://convergedigest.com/broadcom-s-new-trident-switching-silicon-doubles-capacity-adds-neural-engine/
- Competitive context [independent]: Trident5-X12 (16T, ~2023 launch, now shipping to qualified customers per Broadcom 2023 PR) sits one silicon generation behind the Tomahawk-5/6 spine class covered elsewhere in this file — it fills the merchant-silicon explanation for the leaf/ToR tiers of the Dell and Aruba 25/100/400G platforms; 800G support also matters for MS390-era 100G→400G-era migration math.

### 16.4 Dell SmartFabric Manager for SONiC — official solution-brief and spec-sheet detail (extends §1.4 and §P; narrows the GA-vs-roadmap gap)

- Per the official Dell "SmartFabric Manager for SONiC" solution brief [official]: SFM provides (1) **validated blueprints and automatic fabric discovery** for design/deploy; (2) **standard APIs** for automation and unified lifecycle management; (3) fabric resiliency via advanced analytics/monitoring; supports **traditional L3, BGP EVPN VXLAN, standalone, and rail-optimized** fabric designs.
  - Source: http://delltechnologies.com/asset/zh-hk/products/networking/briefs-summaries/smartfabric-manager-for-sonic-brief.pdf
- Per the official SFM AI Blueprints Guide (Dell Info Hub, crawled ~19 days before retrieval) [official]: SFM landing page exposes fabric errors, fabric health, and SFM overall health; default access credentials published as **sfmadmin / Dellsfm@123** (deployment documentation).
  - Source: https://infohub.delltechnologies.com/document_parser/crosslinks/chapter/e3a640fd4p/
- Per the SFM for SONiC spec sheet [official — Dell spec sheet via reseller mirror]: SFM supports up to **192 switches** per fabric; fabric types: **3-tier CLOS** for access/storage or front-end fabrics, **BGP EVPN with VXLAN, L3 with BGP, L2 with MCLAG (leaf only)**; multi-fabric management across sites (requires OOB connectivity); lifecycle management includes uniform NOS version maintenance, auto-rollback on failure, switch replacement with auto-config restore, one-click fabric snapshot backup/restore; switch discovery for Dell PowerSwitch S & Z series and N3248TE running SONiC 4.3+.
  - Source: https://netmateit.com/wp-content/uploads/2025/10/dell-smartfabric-manager-for-sonic-spec-sheet.pdf
- AI fabric specialization per the same spec sheet [official]: **built-in blueprints for AI fabrics** covering three fabrics (GPU backend/scale-out, frontend/access-storage, management); **auto-configures RoCEv2 and PFC watchdog, enables DLB (dynamic load balancing) by default on the GPU fabric**; rail-optimized topology support; RBAC with custom read/write profiles.
- Note on the §Q gap: these documents are live Dell documentation and a shipping spec sheet, which supports treating SFM for SONiC as a shipped product rather than pure roadmap — though no explicit "GA" press date was located [unverified].

### 16.5 Verification log and open items

| Check | Result |
|---|---|
| 16.1 IOS XE 17.18.2 timeline / EVPN guide / BGP scale table | Official Meraki docs (4 URLs); one community secondary cross-check. No live browser needed. |
| 16.2 CX 10040 pricing/SKUs | SHI secondary street price; UK reseller snapshots; official QuickSpecs via reseller mirror; official HPE Store support price. |
| 16.3 Trident5-X12 | Vendor-reported (Broadcom PR) + 4 independent/secondary technical press. Not independently benchmarked here. |
| 16.4 SFM for SONiC spec | Official Dell brief, Info Hub guide, spec sheet via reseller mirror. |
| Duplication guard | `grep` before writing: Trident5=0, 17.18.2=1 (partial), 10040=13 (no pricing), SFM=7 (no feature spec). No overlap with §P/§X/R4 content. |

**Round-16 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally; no identifiers invented. Each fact carries its provenance tag in text. Earlier sections were not modified. Phase B and Phase C were not touched — parent to sequence next work per Anicet's instruction.


---

