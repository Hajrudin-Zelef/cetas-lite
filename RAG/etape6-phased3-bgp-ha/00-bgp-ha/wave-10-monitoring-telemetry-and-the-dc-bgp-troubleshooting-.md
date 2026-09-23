---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-10-monitoring-telemetry-and-the-dc-bgp-troubleshooting-
title: "Wave 10 — Monitoring, telemetry, and the DC BGP troubleshooting checklist"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci", "optics", "research"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [512, 562]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: bb62fabc293ae23f1ccf284b1e65a2591f2d4e7e790bbc9fcb238d757f3e7706
---

# Wave 10 — Monitoring, telemetry, and the DC BGP troubleshooting checklist

## Wave 10 — Monitoring, telemetry, and the DC BGP troubleshooting checklist

### 10.1 BMP and streaming telemetry for BGP

- **BMP (BGP Monitoring Protocol, RFC 7854):** routers stream Adj-RIB-In/Out, peer up/down events, and statistics to a monitoring station — the standard way to get real-time BGP visibility without polling. Supported on IOS-XR, Junos, EOS (recent trains), and FRR [official][secondary].
- **Model-driven telemetry** (gNMI/OpenConfig) complements BMP with interface counters, BFD session state, and FIB programming stats — the combination answers "did the route leave the box, and did the FIB install it?" [independent].
- **What to alert on in a DC fabric** [independent]:
  - BGP session flaps on fabric links (correlate with BFD session state).
  - Unexpected prefix count changes per peer (leak/hijack inside the fabric).
  - ECMP group width changes (path loss without session loss).
  - MC-LAG/MLAG peer-link and keepalive state transitions.
  - FHRP/anycast gateway inconsistencies (duplicate gateway MAC from two leaves).
- **Gap:** no public reference for BMP scale (sessions × prefixes) on DC fixed-form-factor switches; BMP is documented mostly for SP route-monitoring use.

### 10.2 Key verification commands per vendor (fabric BGP health)

- **Arista EOS:** `show ip bgp summary`, `show ip bgp <prefix>`, `show mlag`, `show bfd peers`, `show ip route` ECMP flags [official].
- **Cisco NX-OS:** `show bgp summary`, `show bgp ipv4 unicast <prefix>`, `show vpc`, `show bfd neighbors`, `show ip route` [official].
- **Cumulus/FRR:** `vtysh -c 'show bgp summary'`, `show bgp ipv4 unicast`, `clagctl`, `show bfd peers` [official].
- **Dell OS10:** `show ip bgp summary`, `show vlt brief`, `show bfd neighbors` [official].
- **Aruba CX:** `show bgp ipv4 unicast summary`, `show vsx brief`, `show bfd` [official].
- **Juniper:** `show bgp summary`, `show route <prefix>`, `show iccp`, `show bfd session` [official].

### 10.3 The DC BGP troubleshooting checklist (field-proven ordering)

1. **Physical first:** link state, optics, FEC mismatches — most "BGP down" is Layer 1 [independent].
2. **BFD state:** if BFD is down but the link is up, suspect CoPP policers or CPU starvation before timers [independent].
3. **BGP session state:** `Idle` → TCP/ACL issue; `Active` → no TCP response (check GTSM/MD5 mismatch, TTL); `OpenConfirm` looping → capability mismatch (e.g., extended-nexthop) [independent].
4. **Routes missing:** check inbound prefix-lists/route-maps, `send-community` on the far end, RR reflection (next-hop-self?), RT import on EVPN [independent].
5. **ECMP asymmetry:** compare `maximum-paths` and multipath-relax knobs across tiers; check for polarization [independent].
6. **MC-LAG split-brain:** verify keepalive path independent of peer link; check dual-primary detection config on both peers [independent].
7. **After maintenance:** confirm GR completed (no stale paths), BFD timers restored, GSHUT communities withdrawn [independent].

### 10.4 Final coverage audit (all waves)

| Scope item | Section(s) |
|---|---|
| eBGP vs iBGP, ASN planning, unnumbered BGP | §1.1–1.3, §6.2–6.3 |
| ECMP, timers, GR, BFD | §1.4–1.7, §9.1–9.2 |
| Prefix lists, route maps, communities, RPKI | §2.1–2.4, §6.4 |
| vPC / MLAG / CLAG / VLT / VSX / MC-LAG | §3.1–3.7 |
| EVPN multihoming vs MC-LAG, STP, L3 over MC-LAG | §7.1–7.4 |
| FHRP, anycast gateway, ISSU/NSF | §4.1–4.3 |
| Failure detection, convergence, PIC | §4.4–4.5, §8.1, §9.1 |
| Security (GTSM/MD5/TCP-AO), dynamic peering | §6.1–6.2 |
| DCI borders, monitoring, troubleshooting | §9.4, §10.1–10.3 |

*End of Wave 10. Phase D3 research complete — append-only, all sections intact.*

---

