---
id: collect-261001-automatisation-infra/automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead-2
title: "etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead.md
source_anchor: ""
source_lines: [102, 199]
sha256: f213262656c96ab1ec4065c9eaf5f292818fdb6a1e5a064a9610ac0bb37100d3
---

# etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead

- Build a NSSA on one router: ASBR injects an external static; ABR translates Type-7 → Type-5. Verify with show ospfv3 ipv4 database nssa-external vsshow ospfv3 ipv4 database external .
- FRR cross-check: same triangle on FRR Docker containers. Confirm identical LSA behavior — write a one-paragraph diff in your notes on any CLI differences.
Prove it:
- show ospfv3 ipv4 neighbor andshow ospfv3 ipv6 neighbor both show Full.
- Both AFs show correct routes in show ip route ospf andshow ipv6 route ospf .
GNS3vault labs (gns3vault-archive/OSPF/) — these cover the full week:
| Lab | What to do | 
|---|---|
| ospf-stub-area | Stub area — LSA type 5 filtered. | 
| ospf-totally-stub | Totally stubby — type 3 + 5 filtered. | 
| ospf-nssa-not-so-stubby-area | NSSA — Type-7 to Type-5 conversion. | 
| ospf-totally-nssa | Totally NSSA — all filtering combined. | 
| ospf-virtual-link | Virtual link — connect discontiguous area-0. | 
| ospf-lsa-type-3-summarization | ABR summarization. | 
| ospf-lsa-type-5-summarization | ASBR summarization. | 
Network types extension (hand-written): GNS3vault Frame-Relay network-type labs are obsolete. Build this instead on vIOS-L3:
Topology: 1 hub + 3 spokes using sub-interfaces to simulate NBMA. Apply each network type in sequence and note the DR/BDR behavior:
| Network type | DR elected? | Hello/dead | When to use | 
|---|---|---|---|
| broadcast | yes | 10/40 | Ethernet | 
| non-broadcast | yes | 30/120 | legacy NBMA | 
| point-to-multipoint | no | 30/120 | hub-spoke, no DR needed | 
| point-to-point | no | 10/40 | any p2p link | 
Then redesign the virtual-link topology from ospf-virtual-link to remove the virtual link entirely — add a physical or GRE link between the disconnected ABR and area-0. Write in your notes why virtual links are an anti-pattern in production.
GNS3vault coverage: minimal. The archive has no usable IS-IS labs for modern IOS. Use FRR containers instead — IS-IS is identical conceptually across vendors.
Topology: 4 FRR Docker containers + 2 vIOS-L3. Two areas: 49.0001 (L1 only) and 49.0002 (backbone L2). Two L1/L2 routers bridge the areas.
Tasks:
- Configure NETs on all routers (49.0001.0000.0000.0001.00 etc.). Enablerouter isis on FRR (net 49.0001.0000.0000.0001.00 ),router isis on IOS.
- Set router types: two routers is-type level-1 , twois-type level-2-only , twois-type level-1-2 .
- Use wide metrics everywhere: metric-style wide on IOS,metric-style wide on FRR.
- SHA-256 hello authentication: authentication key-chain IS-IS-AUTH on IOS;isis authentication mode md5 on FRR (FRR supports MD5 only — note the gap).
- Leak selected L2 → L1 routes with a route-map on the L1/L2 routers.
- Insert a deliberate metric tie; resolve with manual metric tuning.
Prove it:
- show isis neighbors shows expected L1/L2 adjacencies.
- show isis database detail shows TLV-extended IS reachability (wide metrics active).
- An L1-only router's show ip route isis shows only intra-area routes + a default toward the L1/L2 router.
Notes to write: compare IS-IS area model vs OSPF area model — specifically why IS-IS has no Area-0 constraint and why that matters for large flat backbones.
GNS3vault labs (gns3vault-archive/BGP/) — this is where the archive is richest. Do these in order:
| Lab | What to do | 
|---|---|
| bgp-basic | Baseline — session, network advertisement. | 
| ibgp-internal-bgp | iBGP full mesh, next-hop behavior. | 
| bgp-next-hop-self | next-hop-self — why it's needed on iBGP. | 
| bgp-route-reflectors | RR + client config. Verify cluster-list and originator-id. | 
| bgp-attribute-local-preference | LOCAL_PREF — AS-wide scope. | 
| bgp-attribute-as-path | AS-path prepending — influence inbound. | 
| bgp-attribute-med | MED — inter-AS metric. | 
| bgp-attribute-weight | Weight — local-router only. | 
| bgp-communities | Standard communities. | 
| bgp-communities-no-export | no-export well-known community. | 
| bgp-as-path-access-list | AS-path regex — block private AS leakage. | 
| bgp-peer-group | Peer groups — operational efficiency. | 
IPv6 AF extension (hand-written): GNS3vault BGP labs are IPv4 only. After the labs above, extend your topology:
Add address-family ipv6 unicast to every peer. Use link-local next-hops where applicable (neighbor X update-source GigE0/1). Confirm with show bgp ipv6 unicast summary.
GNS3vault labs — use these as a foundation, then build the gauntlet yourself:
| Lab | Purpose | 
|---|---|
| bgp-ebgp-multihop | Covers the multihop failure mode. | 
| bgp-update-source | Covers the update-source mismatch failure mode. | 
| bgp-md5-authentication | Covers the MD5 mismatch failure mode. | 
Gauntlet (hand-written extension): take the Lab 8 topology and break it six ways, one at a time. For each bug, predict the BGP state machine state first, then diagnose:
- TCP/179 blocked by an ACL on a transit router → session stuck in Active .
- eBGP multihop missing for loopback-sourced peering → Idle .
- update-source Loopback0 set on one side only →Active .
- MD5 password mismatch → Active (TCP connects, Open rejected).
- iBGP RR client missing next-hop-self → sessionEstablished but routes black-holed.
- Outbound route-map silently denying a prefix → session up, prefix missing from peer's table.
For each: write the minimum diagnostic command sequence in your notes. These become your BGP troubleshooting runbook.
GNS3vault labs (gns3vault-archive/MPLS/) — the archive has solid MPLS coverage:
| Lab | What to do | 
|---|---|
| mpls-ldp | LDP basics — label distribution, LFIB, FIB. Verify with show mpls ldp neighbor andshow mpls forwarding-table . | 
| vrf-lite | VRF review if needed — RD/RT mechanics. | 
| basic-mpls-vpn | Core L3VPN lab — PE-CE with static routes, MP-BGP VPNv4, RT import/export. | 
| mpls-vpn-pe-ce-using-ospf | PE-CE using OSPF — important for the capstone. | 
Extensions (hand-written): after the GNS3vault labs:
- Add a second VRF (VPN_BLUE ) with eBGP as PE-CE protocol instead of OSPF. This gives you both PE-CE options in one topology — important for the Week 16 capstone.
- Prove route isolation: overlap 10.0.0.0/24 in both VRFs. Verifyshow bgp vpnv4 unicast all shows distinct RDs.
- Stretch — 6VPE: add address-family vpnv6 unicast to both PEs. Configure IPv6 CEs. Verifyshow bgp vpnv6 unicast all .
- Juniper cross-check (optional): swap one PE for vMX. CLI differs; RSVP-TE vs LDP differs; L3VPN concept is identical. Write a one-page CLI comparison in your notes.
Prove it:
- show bgp vpnv4 unicast all summary — both PE peers Established.
- show mpls forwarding-table on a P router shows label stacks.
- traceroute vrf VPN_RED 10.0.0.X from a CE shows MPLS labels in output.
GNS3vault coverage: the archive has GRE and IPsec labs but no DMVPN. Use the GNS3vault IPsec labs as a warm-up, then build DMVPN from scratch.
GNS3vault warm-up (gns3vault-archive/Tunneling/):
| Lab | What to do | 
|---|---|
| gre-tunnel-basic | GRE recap — mGRE is an extension of this. | 
| gre-over-ipsec | GRE+IPsec mechanics — the template for DMVPN phase 1. | 
| site-to-site-ipsec-vpn | Plain IPsec — verify IKEv1 SA establishment before upgrading to IKEv2. | 
DMVPN hand-written lab:
Topology: 1 hub CSR1000v + 3 spoke CSR1000v. Add 1 MikroTik CHR as a 4th spoke for Phase 1 only (MikroTik DMVPN interop is IKEv1/Phase-1 only; document the limitation).
Tasks:
- Phase 1 — hub-and-spoke only. Hub: tunnel mode gre multipoint , NHRP NHS config. Spokes:tunnel destination = hub, NHRP NHS = hub. EIGRP across tunnel withno ip split-horizon eigrp on hub.
- Phase 2 — enable NHRP redirect on hub, NHRP shortcut on spokes. Verify spoke-to-spoke tunnel formation with show dmvpn detail .
- Phase 3 — NHRP shortcut + redirect on hub. Spoke-to-spoke triggered by hub redirect. Confirm traffic bypasses hub: capture on hub's transport interface during a spoke-to-spoke ping — hub should see only the NHRP redirect, not the data traffic.
