---
id: etape6-phased3-bgp-ha/00-bgp-ha/9-3-route-reflector-scaling-and-hierarchy
title: "9.3 Route reflector scaling and hierarchy"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [482, 511]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: ce52b054d643ef35412b094303bbf6653abaf6ba5945e6f5ab79f7a65f8bed38
---

# 9.3 Route reflector scaling and hierarchy

- BGP best-path is a long tie-break chain (weight → local-pref → locally originated → AS_PATH length → origin → MED → eBGP>iBGP → IGP metric → router-ID). In a symmetric Clos fabric the decision must be **identical on every switch**, or ECMP groups diverge and troubleshooting becomes nondeterministic [independent].
- Practical rules [independent]:
  - Keep the underlay policy-free: no local-pref/MED manipulation on fabric eBGP sessions; let AS_PATH length + IGP metric decide.
  - `bgp deterministic-med` (compare MEDs only between same-ASN paths) avoids nondeterministic MED comparisons [official].
  - `bgp bestpath as-path multipath-relax` + consistent `maximum-paths` on every tier keeps ECMP groups uniform.
  - Avoid `bgp bestpath compare-routerid` asymmetry: it's deterministic per box but can produce different ECMP sets across boxes if router-IDs differ in ordering relative to paths — the Arista lab enables it deliberately; understand the effect before copying [independent].
- **ECMP polarization:** identical hash inputs on every tier can synchronize path choices (all switches pick the same member), collapsing ECMP. Resilient/variable hashing seeds per switch mitigate it — a forwarding-plane concern adjacent to BGP design [independent].

### 9.3 Route reflector scaling and hierarchy

- RR scaling dimensions: number of clients, prefixes per client, update churn. Two-tier RR (spine RRs + dedicated RRs for DCI/edge) isolates churn domains [independent].
- **RR redundancy:** always deploy RRs in pairs with identical policy; clients peer to both. Inconsistent policy between RRs of a pair causes persistent asymmetric routing — audit RR configs as a pair [independent].
- **Cluster-ID:** set the same cluster-ID on redundant RRs so clients don't see duplicate reflection loops; unique cluster-IDs per RR pair in multi-pair designs [official].
- **iBGP split-horizon reminder:** an RR must not reflect a route back to the client it learned it from; `no bgp enforce-first-as`-style knobs are unrelated — don't confuse eBGP loop prevention with iBGP reflection rules [independent].

### 9.4 DCI / border-leaf BGP design

- Border leaves peer eBGP to WAN/transit/DCI routers with **full policy**: inbound prefix-lists + AS-path filters + max-prefix + RPKI validation (ROV) on Internet-facing sessions; outbound: only owned aggregates, `remove-private-AS` toward public peers [independent].
- **MED vs local-pref at DCI:** use MED to signal preferred entry to external peers; local-pref to prefer one DCI exit internally. AS-path prepending as a coarse traffic-engineering lever of last resort [independent].
- **Default origination:** border leaves commonly originate default into the fabric (or a summarized aggregate) so tenant VRFs don't carry full Internet tables [independent].
- **BGP communities at the border** tag route source (transit vs peer vs cloud interconnect) for downstream filtering and troubleshooting [independent].

### 9.5 Wave-9 verification notes and open items

- **Verified:** PIC Core/Edge mechanics and platform notes; RR-PIC interaction; DCI border patterns; determinism rules.
- **Open:** DC-NOS PIC support matrix beyond Cisco; ECMP polarization measurements in production Clos.
- *End of Wave 9.*

---

