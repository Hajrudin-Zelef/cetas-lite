---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/7-3-anycast-rp-rfc-3446
title: "7.3 Anycast-RP (RFC 3446)"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [192, 206]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 4fed08e037c63154f17bac73a7d01c88345ffa4bf572a49986ae1c8b70bca4d2
---

# 7.3 Anycast-RP (RFC 3446)

### 7.3 Anycast-RP (RFC 3446)
- Mechanism: all RPs in the set share one anycast IP (typically a loopback), advertise the same group→RP mapping; sources/receivers join/register toward the topologically closest RP via unicast routing; RPs MSDP-peer with each other (using unique addresses — never the anycast address as router-ID) and exchange Source-Active (SA) messages so every RP knows all sources `[official]` — RFC 3446; Huawei anycast-RP docs; Juniper example. https://info.support.huawei.com/hedex/api/pages/EDOC1100149308/AEJ0713J/18/resources/admin/sec_admin_multicast_0167.html
- Benefits: RP load sharing, fast failover (unicast routing reconverges to the next-closest RP), optimal RP-proximity paths `[official]` — Huawei; RFC 3446.
- Platform nuance (conflict recorded): Cisco Community (2026): on IOS-XE, IPv4 anycast-RP requires MSDP (or MP-BGP IPv4-multicast AF); Nexus supports an MSDP-less PIM anycast-RP variant; Junos supports anycast-RP with or without MSDP (PIM-only for IPv4 and IPv6) `[secondary]/[official]` — Cisco Community thread; Juniper docs. Do NOT assume MSDP-less anycast-RP works on all platforms — verify per OS. https://community.cisco.com/t5/routing-and-sd-wan/multicast-rp-placement-msdp-vs-anycast/td-p/5551629
- Huawei implementation notes: with >3 RPs, full-mesh MSDP with all peers in one mesh-group; must configure a logical RP address for SA messages (peers discard SAs whose RP address equals their own) `[official]` — Huawei docs.

### 7.4 MSDP
- Multicast Source Discovery Protocol: TCP-based SA exchange between RPs in different PIM-SM domains (or anycast-RP sets); includes RPF checks on SA messages to prevent loops `[official]` — RFC 3446; Huawei.
- Mesh-groups suppress SA flooding between fully-meshed peers — required config hygiene at scale `[official]` — Huawei.

### Wave 7 verification
- Sources: 7 (RFC 3446, Huawei ×2, Juniper, Cisco Community, Huawei config guide). One platform conflict documented explicitly (MSDP-less anycast-RP support).

---

