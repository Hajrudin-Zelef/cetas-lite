---
id: etape6-phased3-bgp-ha/00-bgp-ha/1-7-bfd-integration-with-bgp-session-bring-up-order-and-pitf
title: "1.7 BFD integration with BGP: session bring-up order and pitfalls"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [75, 96]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: 8db163eb2c72e7c1dddf4bb329b1e9367a53607b9ad5cead0fe7830d2ea686f3
---

# 1.7 BFD integration with BGP: session bring-up order and pitfalls

- BGP graceful restart (RFC 4724) lets a restarting speaker ask peers to preserve forwarding state across a control-plane restart; **NSF/NSR** are the platform mechanisms that keep the data plane up during supervisor failover or process restart [official].
- In DC fabrics, graceful restart is commonly enabled on the **EVPN overlay (iBGP)** sessions; on the eBGP underlay its value is debated because ECMP reconvergence around a failed node is often as fast as GR procedures [independent].
- Platform notes: Cisco NX-OS supports BGP graceful restart and NSF/SSO on modular platforms; Nexus 3000's BFD explicitly lacks stateless-restart support (see §1.5) [official]. Arista EOS supports BGP GR; Juniper Junos supports GR and NSR per routing-instance [official][secondary].
- **Unverified:** claims that GR measurably improves user-visible convergence in 3-stage Clos fabrics — lab evidence is mixed; treat vendor GR convergence claims as [vendor-reported].

### 1.7 BFD integration with BGP: session bring-up order and pitfalls

- BFD sessions for BGP are typically **single-hop** on fabric links; multihop BFD is used for iBGP overlay sessions via loopbacks [official][independent].
- Pitfalls documented [official][independent]:
  - BFD must be configured consistently on both peers (IOS-XR notes behavior depends on identical config) [official].
  - Configuring timers below the platform minimum causes "undesirable behavior" — Cisco explicitly warns against 3 ms when the floor is 4 ms [official].
  - On NX-OS, topology changes affecting SVIs can flap BFD sessions; Cisco recommends disabling BFD or raising timers during topology changes [official].
  - BFD and graceful restart interact: aggressive BFD can tear down sessions that GR is trying to preserve — coordinate the two [independent].

### 1.8 Wave-1 verification notes and open items

- **Verified:** RFC 5549 mechanism and multi-vendor support statements; Cisco BFD timer tables; NX-OS BFD limitations; Cumulus extended-nexthop behavior; Dell OS10 `link-local-only-nexthop`.
- **Open:** ASN scheme production survey; BFD scale-at-aggressive-timers data; GR benefit quantification in Clos fabrics.
- *End of Wave 1.*

---

