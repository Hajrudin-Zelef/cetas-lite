---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-4-first-hop-redundancy-anycast-gateway-issu-nsf-and-con
title: "Wave 4 — First-hop redundancy, anycast gateway, ISSU/NSF, and convergence"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "dci", "nvidia", "preemption"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [237, 294]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: dce5db67c781553eb07505397996f681f3cc91b2f7d1d4d4039899f844ed9f5a
---

# Wave 4 — First-hop redundancy, anycast gateway, ISSU/NSF, and convergence

## Wave 4 — First-hop redundancy, anycast gateway, ISSU/NSF, and convergence

### 4.1 FHRP protocols: HSRP, VRRP, GLBP

- **HSRP** (Cisco proprietary, RFC 2281 informational): default hello **3 s**, hold **10 s**; one active router forwards for the virtual IP/MAC (VMAC `0000.0c07.acXX`); preemption configurable. Still the default FHRP on Cisco enterprise/DC L3 boundaries [official][independent].
- **VRRP** (open standard, RFC 5798): default advertisement interval **1 s**; master-down interval ≈ 3× advertisement + skew; virtual MAC `0000.5e00.01XX`; one master forwards per virtual router. Multi-vendor support (Juniper, Arista, Cumulus, Dell) makes it the cross-vendor choice [official].
- **GLBP** (Cisco proprietary): default hello **3 s**, hold **10 s**; one AVG (active virtual gateway) assigns virtual MACs (`0007.b400.XXYY`) so up to 4 AVFs (active virtual forwarders) **load-share** traffic — the only FHRP of the three with native active-active forwarding [official].
- In modern DC fabrics these classic FHRPs are increasingly displaced by **anycast gateway** (§4.2), but they persist at DCI edges, firewall boundaries, and non-EVPN L2 domains [independent].
- **Timers tuning:** sub-second HSRP/VRRP timers are supported on most platforms but aggressive values interact with control-plane load; BFD-for-VRRP exists on some platforms to accelerate master-down detection [official][unverified for universal availability].

### 4.2 Anycast gateway: the DC replacement for FHRP

- **Concept:** every leaf/VTEP presents the **same gateway IP and MAC**; hosts ARP once and the local leaf routes. No master election, no failover event on leaf loss — the gateway is inherently distributed. Host mobility and workload placement become trivial [independent][official].
- **Per-vendor implementations** [official][independent]:
  - **Arista:** `ip address virtual` on the SVI (mutually exclusive with a real `ip address` on the same SVI) + fabric-wide `ip virtual-router mac-address` (must be identical on all leaves). The legacy MLAG-scoped variant is **VARP**; EVPN anycast supersedes VARP in modern fabrics (wider scope: hundreds of leaves vs one MLAG pair; EVPN-discovered vs static) [independent — clauntainerlab EVPN anycast lab].
  - **Cisco:** Distributed Anycast Gateway (DAG) in EVPN VXLAN fabrics — same gateway IP/MAC across all VTEPs; ARP/ND to the gateway IP is intercepted locally and not flooded [official — Cisco EVPN VXLAN IRB guides]. On NCS/IOS-XR, DAG is documented for EVPN-over-MPLS with symmetric IRB [official].
  - **Aruba CX (VSX):** **active-gateway** — same IP/MAC on both VSX peers for VSX LAG-attached hosts [official — AOS-CX VSX whitepaper].
  - **Cumulus:** VRR (virtual router redundancy) anycast mode; netlab docs note vendor naming variance: VARP (Arista), VRR (Cumulus), passive VRRP (Nokia) [independent].
  - **EOS 4.36.2F** added **EVPN VXLAN Single-Gateway Centralized Routing**: for fabrics with a single L3 VTEP/MLAG pair as centralized gateway, eliminates the duplicate BUM traffic caused by having both physical VTEP IP and VARP VTEP IP in the overlay flood set [official — Arista EOS 4.36.2F guide].
- **Netlab convention:** shared anycast MAC default `0200.cafe.00ff` in lab topologies [independent].
- Sources: https://github.com/alukacs03/clauntainerlab/blob/HEAD/labs/32-evpn-anycast-gateway/README.md, https://arista.com/en/um-eos/eos-evpn-vxlan-single-gateway-centralized-routing, https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9500/software/release/17-9/configuration_guide/vxlan/b_179_bgp_evpn_vxlan_9500_cg/configuring_evpn_vxlan_integrated_routing_and_bridging.pdf

### 4.3 ISSU / NSF / SSO: hitless upgrades and supervisor redundancy

- **ISSU** (in-service software upgrade) upgrades system software while the switch keeps forwarding, built on **NSF** (nonstop forwarding: data plane keeps forwarding during control-plane restart) + **SSO** (stateful switchover: standby supervisor takes over with synchronized state) [official — Cisco Nexus 9000/7000 HA guides].
- **Cisco NX-OS specifics** [official][vendor-reported]:
  - Dual-supervisor chassis (Nexus 9500/7000): new image loads on standby sup, switchover occurs, old active then upgrades — "no system downtime" [official].
  - **EISSU** (enhanced ISSU, NX-OS 10.2(2)+): container-based second virtual supervisor; claims **zero data-plane downtime / zero packet loss** and ~3 s control-plane downtime; kernel patches fall back to classic ISSU with longer control-plane outage [vendor-reported — Cisco blogs].
  - **EISSU enabled by default** on Nexus 9300 GX2A/GX2B, and on GX/FX3 with NX-OS 10.3.3+; FX/FX2 need an extra command + reload [vendor-reported].
  - ISSU is **disruptive** (reload required) on chassis with N9K-C95xx-FM-E/Ex, FM-R, or FM-G fabric modules [official — NX-OS 7.x and 10.4(x) HA guides].
  - L2 protocols supported across ISSU/EISSU: STP and vPC [vendor-reported].
- **Other vendors:** Arista EOS supports ISSU with MLAG peer coordination [secondary]; Juniper Junos supports unified ISSU on select platforms [secondary]; Dell OS10 and Cumulus ISSU claims are thinner in public docs — **Gap:** per-vendor ISSU support matrix not fully verified.
- **NSR vs NSF vs GR:** NSR (nonstop routing) keeps protocol adjacencies up across switchover without peer cooperation; NSF/GR needs peer cooperation (RFC 4724 for BGP). DC operators often prefer NSR where available to avoid depending on peer behavior [independent].

### 4.4 Link failure detection: the convergence chain

- End-to-end convergence in a DC fabric is a chain: **physical detection** (link-down, ~ms) → **BFD** (sub-second, §1.5) → **protocol reaction** (BGP withdraw/ECMP rehash) → **FIB reprogram**. The slowest link dominates [independent].
- Transceiver-level signal-detect and port-debounce timers add low-ms delay before the link-down is even signaled; tune `link-debounce` where the platform exposes it [independent].
- BFD over port-channels needs per-member sessions (BoB) to detect single-member failure — Cisco documents BoB minimum 4 ms on IOS-XR [official].
- ECMP reconvergence is local: the detecting switch rehashes flows across surviving paths without waiting for BGP — this is why underlay failures often converge in **tens of ms** even with relaxed BGP timers [independent].

### 4.5 Convergence benchmarks: what is published

- Published numbers are **vendor-reported or lab-scale** and not comparable across vendors or topologies [vendor-reported][independent]:
  - Cisco EISSU claims: zero packet loss on data plane during upgrade [vendor-reported].
  - Juniper documents backup-liveness-detection for MC-LAG peer reboot achieving "sub-second traffic loss" [official — Juniper MC-LAG docs].
  - Community/chaos-engineering labs (e.g., EVPN/RoCEv2 failure-injection on Cumulus) measure detection and time-to-alert rather than standardized convergence [independent].
- **No independent head-to-head benchmark** of link-failure convergence across Cisco/Arista/NVIDIA/Dell/Aruba/Juniper DC fabrics was found. Treat any single-vendor "sub-50 ms convergence" marketing claim as [vendor-reported] and scenario-specific.
- **Gap:** standardized, reproducible convergence benchmark methodology for Clos fabrics.

### 4.6 Wave-4 verification notes and open items

- **Verified:** FHRP protocol defaults; anycast gateway per-vendor implementations; Cisco ISSU/EISSU mechanics and defaults; ISSU disruptiveness on specific fabric modules; convergence chain model.
- **Open:** ISSU support matrix (Cumulus/Dell/Juniper); independent convergence benchmarks.
- *End of Wave 4.*

---

