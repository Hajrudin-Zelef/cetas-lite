---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-1-bgp-as-the-data-center-underlay-design-patterns-addre
title: "Wave 1 — BGP as the data-center underlay: design patterns, addressing, and tuning"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["cost", "ethernet", "nvidia"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [19, 74]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: e9776e5fe34f980c66a070eef34008eeaabff192ec1c8364ed66f97d294ad9a0
---

# Wave 1 — BGP as the data-center underlay: design patterns, addressing, and tuning

## Wave 1 — BGP as the data-center underlay: design patterns, addressing, and tuning

### 1.1 Why eBGP dominates the modern DC underlay

- The dominant contemporary design for spine/leaf fabrics is **eBGP between every tier**, with a unique ASN per device (or per tier), replacing IGP-based underlays in most greenfield Clos fabrics [independent]. The motivation: BGP carries policy natively, scales without flooding domains, and provides loop prevention via AS_PATH without additional mechanisms [independent].
- iBGP remains common in two places: (a) as the **EVPN overlay control plane** (iBGP with route reflectors, typically on spines or dedicated RRs), and (b) in legacy/enterprise DC designs where OSPF/IS-IS underlays persist with iBGP or eBGP only at borders [independent].
- The classic modern pattern: **eBGP underlay for loopback/VTEP reachability + iBGP EVPN overlay for tenant routes** — documented across NVIDIA Cumulus, Arista, Cisco Nexus and Juniper designs [official][vendor-reported].
- Trade-off to record: eBGP-per-link multiplies BGP sessions (every fabric link is a session) but each session is simple and failure domains stay small; iBGP-with-RR reduces session count but reintroduces RR placement and failure-domain design [independent].

### 1.2 ASN planning: private ASNs and common schemes

- Private ASN ranges available for DC use: **64512–65534** (2-byte, RFC 6996) and **4200000000–4294967294** (4-byte, RFC 6996) [official]. 4-byte private ASNs are now widely supported across DC NOSes; older platforms may still be 2-byte-only — check per platform [official][secondary].
- Common schemes observed in the field [independent]:
  - **One ASN per device:** every spine and leaf gets a unique private ASN (e.g., spines 65100–65109, leaves 65200+). Maximizes ECMP visibility and troubleshooting clarity; consumes ASNs fastest.
  - **One ASN per tier:** all spines share one ASN, each leaf (or leaf pair) gets its own. Reduces ASN consumption; requires `as-path multipath-relax` (or vendor equivalent) on spines so ECMP across same-ASN paths is allowed.
  - **Single ASN fabric with confederations or iBGP:** legacy pattern; still seen in some enterprise DCs.
- `as-path multipath-relax` (FRR/Cumulus) / `bgp bestpath as-path multipath-relax` (NX-OS) / equivalent EOS knob: permits ECMP across paths whose AS_PATHs differ only in ASN sequence length/content when the ASNs are the same length — required for per-tier ASN designs [official].
- **Gap:** no authoritative survey of which ASN scheme hyperscalers use in production; public references are lab and enterprise designs.

### 1.3 Unnumbered BGP (RFC 5549): interface-peered eBGP with IPv6 link-local next-hops

- **RFC 5549** defines advertising IPv4 NLRI with an IPv6 next-hop. Combined with IPv6 link-local addressing and interface-based peering, it yields **BGP unnumbered**: eBGP sessions with no IPv4 addresses configured on fabric links at all [official].
- Mechanics [independent][official]:
  - IPv6 neighbor discovery router advertisements run on the point-to-point fabric interface; each side learns the neighbor's link-local address automatically.
  - The BGP neighbor is configured **by interface, not by IP address** (peer-group + interface neighbor); the session forms to the auto-discovered link-local.
  - IPv4 prefixes (loopbacks, VTEP addresses) are advertised with the IPv6 link-local as next-hop; the data plane resolves the next-hop via NDP and forwards the IPv4 packet.
  - Result: **no IPv4 address planning, no IPAM for fabric links** — plug a link between two switches and the session comes up.
- Vendor support status (as documented) [official][independent]:
  - **NVIDIA Cumulus Linux:** supports unnumbered BGP; `extended-nexthop` capability required for IPv4 NLRI over IPv6 global peerings (flapping the capability flaps sessions) [official — docs.nvidia.com, Cumulus Linux 5.15]. Older 3.7.x behavior: extended next-hop encoding sent only for link-local peerings in ≤3.7.1, extended to global unicast in ≥3.7.2 [official].
  - **Juniper:** calls it "BGP Auto Discovered Neighbors"; supported on Junos for eBGP underlay [independent].
  - **Dell SmartFabric OS10 (10.5.6):** supports BGP unnumbered, declares RFC 5549 compliance; documents a `link-local-only-nexthop` knob for interop with FRR-based switches that send link-local in both next-hop fields; notes sessions to FRR switches can drop without it [official — Dell SmartFabric OS10 User Guide 10.5.6].
  - **Arista EOS, Cisco NX-OS/IOS-XR, FRR:** support per community documentation; "some older platforms don't support RFC 5549" — verify per line card [independent][unverified for specific old line cards].
- Operational caveats [independent]:
  - Troubleshooting changes: `show` outputs reference interfaces and link-locals instead of IPv4 peer addresses; monitoring templates must adapt.
  - `redistribute connected` for loopbacks yields origin `incomplete` in some lab configs; `network` statements are the cleaner production pattern [independent].
- Sources: https://github.com/silas2323/fabric-chaos-sentry, https://github.com/alukacs03/clauntainerlab/blob/HEAD/labs/28-bgp-unnumbered/README.md, https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-515/Layer-3/Border-Gateway-Protocol-BGP/Optional-BGP-Configuration/, https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/smartfabric-os-user-guide-10-5-6/bgp-unnumbered?guid=guid-f9be4925-931f-430d-87e2-8dba17218952&lang=en-us

### 1.4 ECMP in the underlay

- BGP ECMP is the load-sharing mechanism of the Clos fabric: every leaf learns every remote loopback via all spines and installs all equal-cost paths [independent].
- Cumulus Linux 5.15 documents that **Type-5 EVPN routes continue to use all available ECMP paths in the underlay regardless of ASN** [official].
- Maximum-paths defaults vary by NOS (commonly 8–64); verify `maximum-paths` per platform and per address family when scaling beyond 8-way ECMP [official].
- ECMP hashing: 5-tuple (or more fields with resilient hashing) in hardware; BGP itself does not hash — forwarding ASICs do. Uneven elephant flows (e.g., RoCE/AI traffic) can still polarize despite ECMP — relevant for AI fabrics but out of scope for this file's HA focus [independent].

### 1.5 BGP timers and tuning for fast DC convergence

- Default BGP timers (keepalive 60s / hold 180s) are far too slow for DC fabrics; production DC designs tune aggressively or rely on BFD [independent].
- Common practice: **BFD for sub-second failure detection** plus relaxed BGP timers, rather than sub-second BGP keepalives (which risk session flaps under control-plane load) [independent].
- Documented BFD timer floors [official]:
  - **Cisco IOS-XR (NCS 5500/560):** single-hop BFD minimum 4 ms, multiplier minimum 3; BGP multihop minimum 50 ms; up to 6 unique timer profiles [official — Cisco NCS 5500/560 BFD guides, IOS XR 7.6.x–7.10.x].
  - **Cisco Nexus 3000 (NX-OS):** BFD v1, IPv4; supports single-hop eBGP and iBGP-with-update-source; **no BFD authentication, no per-link BFD on port-channels, no stateless restart/ISSU support for BFD**; LACP required on port-channels used by BFD [official — Nexus 3000 NX-OS Unicast Routing Guide].
- Practical DC BFD values commonly seen: 300 ms × 3 or 500 ms × 3 for fabric links (conservative), down to 100 ms × 3 on modern platforms; aggressive 3.3 ms values exist but are rarely used at DC scale due to CPU/scale trade-offs [independent][unverified as a universal recommendation].
- **Gap:** no published cross-vendor study of BFD scale (sessions per switch) at aggressive timers in production DC fabrics; vendor maximums are datasheet values.

### 1.6 Graceful restart / NSF for BGP in the DC

