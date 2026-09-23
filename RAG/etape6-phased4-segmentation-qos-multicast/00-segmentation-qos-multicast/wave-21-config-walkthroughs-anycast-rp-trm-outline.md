---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-21-config-walkthroughs-anycast-rp-trm-outline
title: "Wave 21 — Config walkthroughs: anycast-RP + TRM outline"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Huawei"]
dates: []
keywords: ["asic", "ethernet"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [659, 708]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 060a33213dfd581e60c2dde002bf1c3954aa521c864694168377846c6d02eae4
---

# Wave 21 — Config walkthroughs: anycast-RP + TRM outline

## Wave 21 — Config walkthroughs: anycast-RP + TRM outline

### 21.1 Anycast-RP setup (structural, per RFC 3446 + vendor notes)
```
! On each RP router (RP1, RP2):
! 1. Shared anycast loopback — SAME address on all RPs
interface loopback10
  ip address 192.0.2.1/32        ! anycast RP address (example)
  ip pim sparse-mode
! 2. Unique loopback for MSDP peering — MUST differ per RP
interface loopback11
  ip address 192.0.2.11/32       ! RP1 unique; RP2 uses .12
  ip pim sparse-mode
! 3. Group->RP mapping uses the ANYCAST address on ALL routers
ip pim rp-address 192.0.2.1
! 4. MSDP mesh between RPs using UNIQUE addresses
ip msdp peer 192.0.2.12 connect-source loopback11
ip msdp mesh-group anycast-rp 192.0.2.12
! 5. Logical RP address for SA messages (so peers don't discard
!    SAs whose RP address equals their own) — per Huawei note
```
`[official]` — RFC 3446 mechanics; Huawei anycast-RP config guide (mesh-group + logical-RP-address notes); Juniper example (static mapping recommended). Template — exact CLI nouns vary by OS `[secondary]`.

### 21.2 TRM bring-up outline (Nexus 9000, per Cisco whitepaper)
1. Underlay PIM-SM + RP (default VRF); verify `show ip pim neighbor` on all VTEP peerings `[official]` — Cisco TRM whitepaper (Wave 8).
2. Per-VNI underlay multicast groups (`mcast-group` under NVE) OR fabric-wide IR — one mode only `[official]` — same.
3. Tenant VRF: `ip pim rp-address <overlay-RP>` reachable inside the tenant VRF; enable PIM-SSM range if used `[official]` — same.
4. L3VNI + TRM enabled on the tenant VRF; verify EVPN Type-6 (SMET) routes appear when receivers join `[official]/[secondary]` — same whitepaper; operator CLI.
5. Source validation: check TTL at the source (dropped at FHR/LHR if too low) and confirm 224.0.0.0/24 is bridged not routed `[official]` — same.

### 21.3 IGMP snooping + querier checklist
- Snooping enabled per VLAN/VNI; querier enabled on exactly one deterministic device per broadcast domain (commonly the anycast-gateway SVI owner or a designated leaf) `[secondary]` — operator practice.
- In EVPN: confirm IGMP state propagates via Type-7/8 (multi-homed) and selective forwarding via Type-6 — otherwise remote VTEPs flood `[secondary]` — EVPN multicast literature.
- Storm-control on host ports bounds the damage when snooping is misconfigured `[secondary]` — operator practice (Wave 9).

### Wave 21 verification
- Sources: 5. Config blocks are structural templates; TRM steps follow Cisco's whitepaper order.

---

## Final verification (all 21 waves)

- Verification: line count ≥ 750, tail intact, backtick parity even, headers well-formed — checked by script after the final append.
- Scope-to-wave map (final): Waves 1–9 core coverage · 10 EVPN control plane · 11 VRF deep-dive · 12 QoS deep-dive · 13 congestion-control algorithms + UEC · 14 PIM/mVPN deep-dive · 15 services deep-dive · 16 segmentation practice · 17 multicast practice · 18 QoS practice · 19 EVPN config walkthrough · 20 RoCE QoS config walkthrough · 21 anycast-RP/TRM walkthrough.
- Open items (consolidated): VNI allocation practice · ACL TCAM table · TrustSec/SGT deep dive · per-ASIC queue/buffer table · PFC-watchdog defaults · IGMP/EVPN Type-6–8 matrix · PTP profile matrix · DNS-anycast prevalence · UEC deployment timing · per-ASIC PFC calculators — all `[unverified]`.
- Conflicts: C1 anycast-RP/MSDP platform variance · C2 Ethernet-vs-IB parity (vendor perspective) · C3 self-reported lab numbers · C4 QoS defaults non-comparable · C5 RoCE DSCP values are conventions, not standards.
- Provenance: factual claims tagged; editorial meta-lines untagged by design; no identifiers guessed; URLs verbatim from search results.

*End of Phase D4 — 21 waves, append-only, complete.*

