---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-8-detection-extras-fhrp-tracking-control-plane-protecti
title: "Wave 8 — Detection extras, FHRP tracking, control-plane protection, and GR procedures"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "dci", "optics", "preemption"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [427, 481]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: 269ba9be274a70f6280324b6a0cb998a507f93247b029758c7948f08f0e7b3fd
---

# Wave 8 — Detection extras, FHRP tracking, control-plane protection, and GR procedures

## Wave 8 — Detection extras, FHRP tracking, control-plane protection, and GR procedures

### 8.1 Beyond BFD: UDLD, LACP timers, debounce

- **UDLD** (unidirectional link detection): detects one-way fiber/cabling faults that link-up/link-down misses; aggressive mode can error-disable the port. Relevant on long fiber runs between DC rows or to DCI; rarely used inside the rack [official][independent].
- **LACP timers:** slow (30 s periodic, 90 s timeout) vs fast (1 s periodic, 3 s timeout) per IEEE 802.1AX. Fast LACP is standard on DC server-facing and MC-LAG member ports; both ends must agree or the actor falls back per implementation [official].
- **Link debounce:** physical interfaces often delay signaling link-down by tens of ms to ride out transients; tune down on fabric links where BFD already guards against flaps, but keep some debounce on optics prone to flapping [independent].
- **Carrier-delay / link-flap dampening:** NX-OS and EOS expose carrier-delay timers; use with BFD so a flapping link doesn't flap the whole BGP session [official][independent].

### 8.2 FHRP tracking and deterministic failover

- **Object tracking** (Cisco): track upstream/core reachability (interface line-protocol, IP SLA, route reachability) and tie HSRP/GLBP priority decrement or vPC suspension to it — e.g., the vPC object-tracking design in §3.1 that suspends vPC legs when uplinks die, preventing blackholing [official — Cisco NX-OS 10.1(x)].
- **VRRP tracking:** priority decrement on tracked-object failure with preemption restores deterministic mastership; combine with BFD-for-VRRP where supported for sub-second master-down [official][independent].
- Design rule: the FHRP/anycast layer must fail over **slower** than the underlay reconverges, or gateways flap during fabric reconvergence — coordinate FHRP timers with BFD/BGP convergence budgets [independent].

### 8.3 Control-plane protection: CoPP and BGP/BFD stability

- **CoPP** (control-plane policing): rate-limits traffic punted to the CPU (BGP, BFD, ARP, LACP, STP) so data-plane events can't starve protocol keepalives. Default CoPP policies exist on NX-OS/EOS/Junos; DC fabrics with aggressive BFD timers must verify the BFD class policer admits the configured session scale [official][independent].
- Operational lesson: most "BGP session flapped for no reason" incidents trace to CPU starvation (broadcast storm, ARP flood) rather than protocol bugs — CoPP + storm-control are the fix, not longer timers [independent].
- **BGP update throttling:** advertisement-interval and MRAI defaults (30 s eBGP / 5 s iBGP per RFC 4271) are often lowered or disabled in DC fabrics for faster propagation; FRR/EOS/NX-OS all expose per-neighbor knobs [official].

### 8.4 BGP graceful restart procedures (RFC 4724) in detail

- **Restarting speaker** sets the Graceful Restart capability with a **restart timer**; on restart it re-establishes sessions and asks helpers to preserve forwarding state marked **stale** until the **stale-path (selection-deferral) timer** expires or updates arrive [official — RFC 4724].
- **Helper mode:** peers retain routes and keep forwarding; if the restarting speaker doesn't return within the restart timer, helpers flush stale state [official].
- In DC practice: enable GR helper on all fabric speakers; enable restart capability on devices undergoing ISSU/supervisor switchover. On pure eBGP underlays many operators rely on ECMP reconvergence instead — document the choice per fabric [independent].
- **Interaction with BFD:** BFD will kill the session before GR completes unless BFD is also made GR-aware or timers are relaxed during maintenance windows — a classic maintenance-window footgun [independent].

### 8.5 Redundancy beyond protocols: power, supervisors, and process restart

- **Hardware:** dual PSUs on separate feeds, dual supervisors on modular chassis (Nexus 9500/7000), redundant fabric modules — protocol HA is moot if the box dies from a single PSU fault [independent].
- **Process-level HA:** NX-OS modular processes can restart independently (BGP process restart with GR/NSF); EOS multi-agent model isolates agents similarly [official][secondary].
- **Maintenance discipline:** drain traffic before maintenance (BGP graceful shutdown via GSHUT, RFC 8326 — advertise routes with low local-pref / high MED or use `graceful-shutdown` community) rather than relying on fast reconvergence alone [official — RFC 8326][independent].

### 8.6 Wave-8 verification notes and open items

- **Verified:** UDLD/LACP/debounce mechanics; object-tracking designs; CoPP role; GR procedures; GSHUT (RFC 8326).
- **Open:** per-platform CoPP defaults vs aggressive-BFD scale validation data.
- *End of Wave 8.*

---

## Wave 9 — BGP PIC, best-path determinism, RR scaling, and DCI border design

### 9.1 BGP PIC (Prefix Independent Convergence)

- **PIC Core:** handles failures in the transport (IGP/next-hop resolution) without per-prefix BGP recomputation. BGP prefixes share a hierarchical FIB next-hop pointer; when the IGP reroutes, all prefixes switch at once. Requirements: redundant IGP paths (ECMP/LFA), CEF, fast detection (BFD). Typical result: sub-second, often <200 ms, convergence independent of prefix count [official — Cisco IOS/IOS-XR BGP PIC guides][independent].
- **PIC Edge:** pre-installs a **backup BGP next-hop** per prefix in RIB/FIB/CEF; on edge link/node failure the backup takes over immediately — "subsecond convergence ... regardless of the number of prefixes" per Cisco IOS-XR docs [official]. Needs `bgp additional-paths install/select`, `advertise additional-paths`, and BFD (`fall-over bfd`) on edge peers [independent].
- **Caveat for DC fabrics:** PIC was designed for SP cores with IGP transport. In eBGP-underlay Clos fabrics, ECMP already provides the equivalent of PIC Core (local rehash on link failure), and PIC Edge's backup-next-hop concept maps to having multiple ECMP paths pre-installed — which BGP multipath already does. PIC matters most at **DCI borders** (iBGP/MPLS edges) rather than inside the fabric [independent].
- **RR interaction:** a route reflector advertises only its best path, so clients don't learn backup next-hops — PIC Edge behind an RR needs **ADD-PATH** or diverse-path (multiple RRs) designs [independent — orhanergun.net analysis].
- **Platform notes:** Cisco documents PIC Edge for IPv4/IPv6/VPNv4/VPNv6/6PE/VPLS/VPWS/**EVPN** over labeled unicast; PIC over BVI not supported; one primary + one backup path for labeled loopback peering [official — Cisco NCS 560 / 8000 BGP guides].
- Sources: https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/iproute_bgp/configuration/15-mt/irg-15-mt-book/irg-bgp-mp-pic.html, https://orhanergun.net/bgp-pic-prefix-independent-convergence-fundemantals

### 9.2 Best-path determinism and ECMP stability

