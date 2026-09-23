---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/3-5-aliasing-and-backup-paths
title: "3.5 Aliasing and backup paths"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [413, 478]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: fda0512cae1396ef65db4a875eae9b18fd46b8a99fcfccdeccd606ab1663f271
---

# 3.5 Aliasing and backup paths

### 3.5 Aliasing and backup paths

- **Aliasing**: Type-1 per-ES routes let a remote PE send traffic for a MAC to
  *any* PE on the ES (not just the PE that advertised the Type-2). Enables ECMP
  across multihomed PEs — the load-balancing half of all-active [official].
- **Backup path / fast convergence**: on PE failure or ES down, the affected PE
  withdraws its Type-1 per-ES routes ("mass withdrawal") instead of withdrawing
  thousands of Type-2 MAC routes — remote PEs reconverge in one step. This is
  the sub-second convergence story vs. MAC-by-MAC withdrawal [official]
  [vendor-reported].
- **Preferred DF / don't-preempt** knobs exist on several platforms to avoid
  BUM flaps during planned maintenance [vendor-reported][unverified].

### 3.6 Gaps — Wave 3

- HRW / RFC 8584 algorithm support per shipping platform: not collected —
  gap; default modulo behavior is [official], everything else [unverified].
- Exact duplicate-MAC-detection thresholds and freeze behaviors are
  vendor-specific and were not tabulated this wave.

---
---

## Wave 4 — DCI: EVPN multi-site, OTV comparison

### 4.1 Why not stretch raw EVPN-VXLAN across the WAN

Stretching one flat EVPN-VXLAN fabric across sites is an anti-pattern
[vendor-reported][secondary]:

- Full-mesh VTEP reachability across the WAN; every site's MAC table sees every
  other site's MACs — control-plane and data-plane flooding scale collapse.
- BUM flooded everywhere; a broadcast storm in one DC hits all DCs.
- No failure-domain isolation: a control-plane issue propagates globally.

The fix is hierarchy: each site runs its own EVPN-VXLAN fabric; sites are
interconnected through border gateways (BGWs) with selective re-origination
[vendor-reported].

### 4.2 EVPN multi-site architecture (Cisco model; concepts are generic)

Per Cisco's VXLAN EVPN Multi-Site design white paper [vendor-reported]:

- **Border gateway (BGW)**: separates the site-internal fabric from the
  site-external DCI transport. Site-internal VTEPs are **masked behind the
  BGWs** — only BGW underlay IPs are visible in the inter-site transport.
- BGWs **re-originate** EVPN routes: site-local Type-2/Type-5 routes stay
  local; cross-site routes are re-advertised by the BGW with its own IP as
  next-hop. Remote sites therefore see one next-hop (the BGW), not hundreds of
  VTEPs.
- **Anycast BGW**: a pair of BGWs per site share a virtual IP; site-internal
  VTEPs see one logical DCI next-hop [vendor-reported].
- **DCI tracking** (`evpn multisite dci-tracking` on site-external interfaces):
  mandatory on Cisco NX-OS to enable the multi-site virtual IP; at least one
  tracked interface must be up [vendor-reported].
- BUM between sites uses **ingress replication exclusively** (simplifies the
  site-external underlay — no PIM needed across the WAN) [vendor-reported].
- I-E-I model: IGP + iBGP EVPN inside the site; **eBGP** between sites over the
  DCI transport [vendor-reported][secondary].
- Site-internal nodes interoperating with a BGW must support: VXLAN with PIM-ASM
  or ingress replication (Type-3), Type-2/Type-5, RR capable of Type-4, VXLAN OAM
  [vendor-reported].

Arista's equivalent is branded **EVPN multi-domain / DCI with BGWs**; the
re-origination concept is the same, config syntax differs [vendor-reported].

