---
id: etape6-phased3-bgp-ha/00-bgp-ha/11-2-bgp-timer-defaults-reference-control-plane
title: "11.2 BGP timer defaults reference (control-plane)"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["asic"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [637, 685]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: 6cb4b4d0efda4afd8fa36284afef955e2259e58ed50f0777535942cf00f8f8bc
---

# 11.2 BGP timer defaults reference (control-plane)

### 11.2 BGP timer defaults reference (control-plane)

| NOS | Default keepalive / hold | Notes |
|---|---|---|
| FRR / Cumulus | 60 s / 180 s | `timers bgp 3 9` style tuning per neighbor |
| Arista EOS | 60 s / 180 s | per-peer-group `timers` |
| Cisco NX-OS | 60 s / 180 s | `timers bgp <ka> <hold>` |
| Junos | 30 s / 90 s | `hold-time` (keepalive = hold/3) |
| Dell OS10 | 60 s / 180 s | per-neighbor timers |

- DC practice: leave BGP timers near defaults and get speed from **BFD** (§1.5), not from sub-second keepalives [independent].
- All values [official] per vendor docs; verify against the running release train.

### 11.3 BFD capability matrix (fabric-relevant)

| Platform | Min single-hop timer | BFD on port-channel | Auth | Notes |
|---|---|---|---|---|
| Cisco IOS-XR (NCS 5500/560) | 4 ms ×3 | BoB 4 ms ×3 | yes | 6 timer profiles max [official] |
| Cisco NX-OS (N3K) | platform-dependent | LACP required | **no** | no stateless restart/ISSU for BFD [official] |
| Arista EOS | 50 ms typical floor* | yes | yes | S-BFD hold-down in 4.34+ [official] |
| Junos (QFX) | 300 ms ×3 typical* | yes | yes | ICCP liveness 2 s ×4 documented [official] |
| Cumulus/FRR | 300 ms ×3 typical* | yes | yes | ptm-bfd integration [official] |
| Dell OS10 | platform-dependent | yes | yes | [official] |

`*` Typical deployed values; absolute floors vary by platform/ASIC — verify per datasheet. [independent][unverified for exact floors]

### 11.4 Failure walkthroughs (what happens, step by step)

**Scenario A — one spine lost in a 2-spine Clos (eBGP underlay, BFD 300 ms ×3):**
1. Leaf BFD sessions to the dead spine time out (~0.9 s worst case; link-down is faster if the failure is local).
2. Each leaf withdraws/expires the eBGP session; ECMP groups shrink from 2-way to 1-way — reconvergence is local, no BGP best-path recomputation needed for surviving paths.
3. EVPN overlay sessions to the spine (if spines are RRs) drop; leaves reconverge EVPN via the surviving RR. With dual RRs, no control-plane outage.
4. Traffic impact: sub-second with BFD; ECMP rehash may briefly reorder flows [independent].

**Scenario B — vPC peer-link failure (Cisco):**
1. Keepalive (1 s) confirms the peer is alive → secondary disables all vPC member ports (hold-timeout 3 s guards against transients).
2. Downstream access switches see member links drop and reconverge LACP to the primary's members.
3. Orphan devices on the secondary are blackholed until the peer link recovers — design orphan ports off vPC pairs [official].

**Scenario C — MLAG peer-link failure with dual-primary detection (Arista):**
1. Heartbeat over the separate path (VRF mgmt) keeps flowing → both peers know the other is alive.
2. Without dual-primary detection both would go primary and forward independently (split-brain); with it, behavior is deterministic per `dual-primary` configuration.
3. After recovery, `dual-primary recovery delay` staggers port bring-up (e.g., MLAG 60 s) to let control planes settle [official].

**Scenario D — ISSU on one Nexus 9500 (dual sup):**
1. New image loads on standby sup; NSF/SSO keep data plane forwarding.
2. Switchover: BGP GR helpers preserve routes; control plane reconverges in seconds (EISSU: ~3 s control-plane downtime claimed).
3. vPC peer sees no flap if keepalives stay within timers [vendor-reported].

