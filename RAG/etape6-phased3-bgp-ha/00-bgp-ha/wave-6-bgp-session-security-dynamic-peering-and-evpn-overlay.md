---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-6-bgp-session-security-dynamic-peering-and-evpn-overlay
title: "Wave 6 — BGP session security, dynamic peering, and EVPN overlay control-plane design"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci", "ethernet", "preemption"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [342, 402]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: c936b47d7d111b18eb76a72d164381bdde9a76333e7faec3174644c8803902b5
---

# Wave 6 — BGP session security, dynamic peering, and EVPN overlay control-plane design

## Wave 6 — BGP session security, dynamic peering, and EVPN overlay control-plane design

### 6.1 Securing fabric BGP sessions

- **GTSM / TTL security (RFC 5082):** sender sets TTL=255; receiver discards packets with TTL below 254 (single-hop). Defeats off-path spoofed TCP RST/SYN attacks against eBGP sessions. Configured per neighbor (`ttl-security hops 1`); mutually exclusive with `ebgp-multihop` on most implementations [official — RFC 5082; VyOS/FRR docs]. Netlab exposes it as `bgp.gtsm` [independent].
- **TCP MD5 (RFC 2385):** keyed hash over TCP segments with a shared password; widely supported on fabric BGP sessions (Arista's EVPN lab uses `password` on peer-groups). Susceptible to collision attacks in theory; key rotation is operationally painful [independent][official].
- **TCP-AO (RFC 5925):** obsoletes MD5 with stronger hashes (HMAC-SHA-1-96, AES-128-CMAC) and key rollover via KeyID. Implementation status as presented at APNIC 50 (2020, Juniper): Nokia SR OS (16.0.R15/19.10.R7/20.5.R1), Cisco IOS-XR (stable since 6.6.3/7.0.1), Juniper (20.3R1, tested against Nokia); Arista was "no comment on timelines" at that time [secondary]. **Gap:** 2026 DC-NOS TCP-AO support status not re-verified — treat the 2020 status as dated.
- **Defense in depth:** GTSM + MD5/TCP-AO + prefix-lists + max-prefix + control-plane policing (CoPP) is the recommended stack; no single mechanism suffices [independent].
- Sources: https://conference.apnic.net/50/assets/files/APCS790/The-TCP-Authentication-Option.pdf, https://github.com/ipspace/netlab/blob/HEAD/docs/plugins/bgp.session.md

### 6.2 Dynamic BGP peering and peer groups

- **Dynamic peering** (`bgp listen range` on EOS, `neighbor ... peer-group` + `bgp listen` on NX-OS, `neighbor fabric interface` on FRR/Cumulus): the switch accepts BGP sessions from any neighbor in a configured prefix range, assigning them to a peer-group template. Combined with `peer-filter`/`as-range` checks, this enables **zero-touch leaf onboarding** — a new leaf cabled to spines forms eBGP sessions with no per-neighbor config on the spine [official][independent — Arista ATD lab].
- **Peer groups/templates** reduce config duplication and improve update-generation scaling: members of a peer group share outbound update state [official].
- **Passive peers:** `neighbor ... passive` (accept inbound, never initiate) is used on spines to avoid connection collisions during bring-up [independent — VyOS/FRR docs].
- Trade-off: dynamic peering widens the attack surface — always combine with GTSM, MD5 passwords on the peer-group, and ASN-range filters [independent].

### 6.3 EVPN overlay control plane: iBGP, route reflectors, next-hop handling

- The EVPN address family is virtually always **iBGP** (same ASN across the fabric or a dedicated overlay ASN), with **route reflectors** on spines or dedicated devices; leaves are RR clients [independent].
- **RR placement options:** (a) spines as RRs (common in 2-tier fabrics; spines already see all leaves); (b) dedicated RR pair (larger fabrics; isolates control-plane load); (c) every spine as RR with full mesh between spines [independent].
- **next-hop-self on RRs:** required so leaves resolve VTEP reachability via the RR's loopback only if the design intends it; in most DC designs the RR does **not** change next-hop — the originating leaf's loopback/VTEP IP is preserved end-to-end so VXLAN tunnels terminate on the right VTEP [independent].
- **ADD-PATH (RFC 7911):** lets RRs advertise multiple paths for the same NLRI — useful for fast EVPN reconvergence and for backup-path visibility; support varies by NOS/version [official][secondary].
- **Overlay BFD:** multihop BFD between loopbacks can protect iBGP EVPN sessions; Cisco documents BGP-multihop BFD minimum 50 ms on IOS-XR [official].
- **Route-target discipline:** RTs define tenant/VRF membership; misconfigured RT import/export is a leading cause of "routes in BGP but not in VRF" incidents — validate with `show bgp l2vpn evpn` per VRF [independent].

### 6.4 Communities in the overlay: RTs, ESI, and policy tags

- **Extended communities** carry the EVPN control plane: Route Targets (import/export), Encapsulation (VXLAN), ESI labels (multihoming), Router MAC, and Default Gateway (for anycast-gateway advertisement) [official — RFC 8365].
- **BGP ESI (Ethernet Segment Identifier):** 10-byte identifier shared by all links of a multihomed device; advertised in EVPN Type-1 (per-ESI auto-discovery) and Type-4 (ESR) routes; enables **all-active multihoming** with DF (designated forwarder) election for BUM traffic [official — RFC 7432].
- Operator-defined **large communities** tag tenant, site, or security posture for DCI policy; keep a registry — community sprawl is a real operational hazard [independent].

### 6.5 Wave-6 verification notes and open items

- **Verified:** GTSM/MD5/TCP-AO mechanics and 2020 implementation status; dynamic peering patterns; RR design options; EVPN extended community roles.
- **Open:** 2026 TCP-AO support per DC NOS; ADD-PATH support matrix.
- *End of Wave 6.*

---

## Wave 7 — EVPN multihoming vs MC-LAG, STP interplay, and L3 over MC-LAG

### 7.1 EVPN multihoming (EVPN-MH): the standards-based alternative

- **Mechanism (RFC 7432):** a multihomed CE connects via LACP LAG to two+ PEs sharing an **Ethernet Segment (ES)** identified by a 10-byte **ESI**. PEs discover each other via EVPN Type-4 (ES) routes, elect a **Designated Forwarder (DF)** per ES per service for BUM traffic, and use **aliasing** so remote PEs load-balance known unicast across all PEs of the ES even if the MAC was advertised by only one [official — RFC 7432; Nokia SR OS docs].
- **Redundancy modes:** **all-active** (all PEs forward; per-flow load balancing) and **single-active** (one active, others standby with pre-installed backup path) [official — RFC 7432].
- **DF election algorithms:** default modulo (RFC 7432 §8.5 service carving), Highest Random Weight (RFC 8584), Highest/Lowest Preference (RFC 9785); implementations let operators set preference and non-preemption [official][independent — rustbgpd EVPN docs].
- **Split-horizon:** BUM from non-DF to DF carries the ESI label so the DF doesn't echo it back to the ES [official — Nokia docs].
- **Netris guidance:** recommends EVPN-MH over MC-LAG where hardware supports it — standardized, works with VXLAN overlays, faster convergence via aliasing, scales beyond two devices; MC-LAG "typically scales to two devices" with convergence/scaling limits [vendor-reported — Netris docs].
- **Vendor support notes:** ALE (OmniSwitch) documents all-active multihoming from 8.10R2 [vendor-reported]; Nokia 7750 SR/7450 ESS/7950 XRS implement DF election + split-horizon + aliasing [official].
- Source: https://infocenter.nokia.com/public/7750SR217R1A/topic/com.nokia.L2_Services_and_EVPN_Guide_21.7.R1/evpn_all-active-ai9enrmqel.html, https://www.fs.com/uk/blog/evpn-multihoming-and-mlag-key-differences-16957.html

### 7.2 EVPN-MH vs MC-LAG: decision guidance

- **Choose EVPN-MH when:** the fabric is already EVPN/VXLAN, the CE needs all-active forwarding with per-flow load balance, or more than two switches must back the same ES [independent][vendor-reported].
- **Choose MC-LAG when:** the attachment is pure L2 without an EVPN overlay, the switches lack EVPN-MH support, or operational simplicity of a single logical LAG matters more than standards-based multihoming [independent].
- **They can coexist:** MC-LAG pairs can themselves be VTEPs in an EVPN fabric (e.g., Arista MLAG pair as one logical VTEP with shared VTEP IP) — the MC-LAG handles the local LAG, EVPN handles fabric-wide reachability [independent].
- FS.com's comparison article frames it as protocol-based multihoming (flexible, scalable) vs device-link aggregation (simpler, two-box) [secondary].

### 7.3 STP interaction with MC-LAG

