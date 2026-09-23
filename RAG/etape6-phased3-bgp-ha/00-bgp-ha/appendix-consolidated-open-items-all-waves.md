---
id: etape6-phased3-bgp-ha/00-bgp-ha/appendix-consolidated-open-items-all-waves
title: "Appendix — Consolidated open items (all waves)"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["benchmarks", "hyperscaler"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [563, 636]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: 0bf3efbbb4d38e2a7952aa893f5c7f9b18d40bbf186632297055a3f15b04df2c
---

# Appendix — Consolidated open items (all waves)

## Appendix — Consolidated open items (all waves)

1. Production ASN-scheme survey (hyperscaler/enterprise).
2. BFD scale at aggressive timers in production fabrics.
3. Quantified GR benefit in Clos fabrics (lab evidence mixed).
4. RPKI-at-DC-border adoption survey.
5. ISSU support matrix for Cumulus/Dell/Juniper MC-LAG.
6. Independent cross-vendor convergence benchmarks; standardized methodology.
7. Sub-second MC-LAG failover claims are vendor-reported and scenario-specific.
8. 2026 TCP-AO support per DC NOS; ADD-PATH support matrix.
9. L3-over-MC-LAG convergence benchmarks.
10. BMP scale reference on DC fixed-form-factor switches.
11. ECMP polarization measurements in production Clos.

## Wave 11 — Reference configs, timer tables, and failure walkthroughs

### 11.1 Annotated eBGP underlay config patterns

**Arista EOS — eBGP numbered underlay with peer-group (lab pattern):**
```
router bgp 65101
   router-id 10.0.0.11
   no bgp default ipv4-unicast
   maximum-paths 8
   neighbor FABRIC peer group
   neighbor FABRIC remote-as 65100
   neighbor FABRIC send-community
   neighbor FABRIC maximum-routes 12000
   neighbor 10.1.0.0/31 peer group FABRIC   ! one statement per fabric link
   redistribute connected route-map RM-LOOPS
!
ip prefix-list PL-LOOPS seq 10 permit 10.0.0.0/24 eq 32
route-map RM-LOOPS permit 10
   match ip address prefix-list PL-LOOPS
```
[independent — distilled from Arista ATD EVPN lab patterns]

**FRR/Cumulus — BGP unnumbered (interface peering):**
```
router bgp 65101
 bgp router-id 10.0.0.11
 bgp bestpath as-path multipath-relax
 neighbor fabric peer-group
 neighbor fabric remote-as external
 neighbor fabric capability extended-nexthop
 neighbor swp1 interface peer-group fabric
 neighbor swp2 interface peer-group fabric
```
[official — Cumulus Linux BGP docs]

**Cisco NX-OS — eBGP underlay with BFD:**
```
feature bgp
feature bfd
router bgp 65101
  router-id 10.0.0.11
  address-family ipv4 unicast
    maximum-paths 8
  neighbor 10.1.0.0 remote-as 65100
    bfd
    address-family ipv4 unicast
```
[official — NX-OS unicast routing guides]

**Junos — eBGP underlay with BFD and multipath:**
```
set protocols bgp group FABRIC type external
set protocols bgp group FABRIC peer-as 65100
set protocols bgp group FABRIC neighbor 10.1.0.0
set protocols bgp group FABRIC multipath
set protocols bfd ...
```
[official — Junos routing protocols docs; BFD timers per platform]

