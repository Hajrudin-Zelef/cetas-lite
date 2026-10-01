---
id: collect-261001-automatisation-infra/automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead-1
title: "etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead.md
source_anchor: ""
source_lines: [1, 101]
sha256: 93f72bb5f5ab5ec0b4215ee83431a44db05369c197cc7a9babf130e19f9c7808
---

# etherne1-obsidian-vault-blob-head-post-20encor-post-encor-labs-eve-ng-md-562ebead

Lab source philosophy: use pre-validated GNS3vault labs (René Molenaar) wherever the topic exists in the archive. Hand-written labs only for topics GNS3vault is too old to cover: DMVPN, IS-IS, BFD, OSPFv3 address-families, named-mode EIGRP, IPv6 FHS, gNMI/telemetry, MikroTik cross-checks, automation, and the capstone.
GNS3vault repo: https://github.com/networklessons/labs/tree/main/gns3vault-archive/
How to use GNS3vault labs:
- Open the lab's .md file — it contains the scenario, numbered task list, and a YouTube walkthrough link.
- Build the topology in EVE-NG matching the .png diagram.
- Apply each router's startup-configs/RN.txt to get the same starting point.
- Work through the task list. Don't open final-configs/ until you've attempted everything.
- Verify with the YouTube walkthrough. Only then diff your config against final-configs/ .
Platform translation: GNS3vault uses IOS 12.4/15.x on 3640/3725/7200. Your CSR1000v 17.3.04a is upward-compatible for all topics below. Rename Serial0/0 → GigabitEthernet1, FastEthernet0/0 → GigabitEthernet1. L2 features use vIOS-L2 instead of the NM-16ESW switch module.
Assumed images in EVE-NG:
- Cisco vIOS-L2, vIOS-L3, CSR1000v 17.3.04a
- MikroTik CHR (free from mikrotik.com — boots on x86 EVE-NG)
- FRR on Alpine/Debian Docker (for IS-IS and vendor-neutral drills)
- Juniper vMX (optional — everything works without it)
Conventions:
- Loopback0 = X.X.X.X/32 where X = router number (R1 → 1.1.1.1/32)
- Customer networks: 10.X.0.0/24 . Backbone p2p links:192.168.0.X/30 slices.
- IPv6: ULA fd00:X::/64 per router.
- Every lab ends with a prove it step and a destroy and rebuild from notes step.
Coverage split: VRF-Lite → GNS3vault. BFD → hand-written (not in GNS3vault).
GNS3vault labs (gns3vault-archive/MPLS/):
| Lab | What to do | 
|---|---|
| vrf-lite | Start here. VRF creation, RD, interface assignment, routing separation. | 
| vrf-routing | Inter-VRF routing via route leaking. Prove overlapping 10.0.0.0/24 is isolated. | 
After both labs work, extend the topology yourself: add a SHARED VRF and leak specific routes into both customer VRFs using RT import. This is the creative step — no task list, just your notes.
MikroTik cross-check (hand-written): repeat the vrf-lite scenario on a CHR (/ip vrf, /routing/ospf/instance, route-leaking via /ip route). Target: same VRF isolation, same overlapping prefix, proven with /ip route print vrf=CUST_A.
GNS3vault has no BFD lab. Build this on two CSR1000v routers:
Tasks:
- OSPF adjacency between R1 and R2 on Gi0/1 (area 0, point-to-point).
- Enable BFD: bfd interval 300 min_rx 300 multiplier 3 on both interfaces.
- Enable BFD under OSPF: bfd all-interfaces .
- Verify: show bfd neighbors detail — confirm asynchronous mode, state Up, echo enabled.
- Simulate failure: shutdown the link. Measure convergence time against OSPF-only dead-interval (40 s default) vs BFD (0.9 s). Write the delta in your notes.
- Repeat for a BGP session: neighbor X fall-over bfd . Tear the session, confirm faster detection.
Prove it:
- show bfd neighbors detail shows300 ms × 3 , stateUp .
- After link failure: show ip ospf neighbor shows neighbor gone within 1 second.
GNS3vault labs (gns3vault-archive/Network Services/ and gns3vault-archive/BGP/):
| Lab | What to do | 
|---|---|
| policy-based-routing | Core PBR lab — route-map match ip address ,set ip next-hop verify-availability . | 
| bgp-filtering-extended-access-list | ACL-based BGP filtering — builds the prefix-list mental model. | 
| bgp-as-path-access-list | AS-path regex filtering — directly exam-relevant. | 
After the GNS3vault labs, do this extension on your own topology (no task list — build from knowledge):
- Six-router partial mesh, 3 ASes, eBGP between them. Apply a route-map outbound that: denies RFC1918 (prefix-list BLOCK_RFC1918 ), sets local-pref 200 on a specific prefix. Verify withshow ip bgp neighbors X advertised-routes .
- PBR: from a stub LAN, route TCP/80 out a backup link, all other traffic via primary. Add track 1 on the primary — PBR falls back when the primary is down.
MikroTik cross-check: same PBR on a CHR via /ip firewall mangle (mark-routing) + /ip route rule. Compare the conceptual flow in your notes — note where Cisco and MikroTik differ in where the policy sits.
GNS3vault does not have a dedicated redistribution lab but the building blocks are there. Use this sequence:
| Lab | Purpose | 
|---|---|
| ospf-single-area | Baseline OSPF — use as left side of the redistribution topology. | 
| eigrp-basic | Baseline EIGRP — use as right side. | 
Then hand-build the redistribution scenario:
Topology: 4 vIOS-L3 in a square. Left edge OSPF area 0, right edge EIGRP AS 100. Two middle routers are the ASBRs doing mutual redistribution.
Tasks:
- Redistribute OSPF into EIGRP and EIGRP into OSPF on both ASBRs — no filtering first. Observe the routing table. Identify the loop or suboptimal path with traceroute .
- Fix with route tags: tag 100 when redistributing OSPF→EIGRP,tag 200 when redistributing EIGRP→OSPF. Addroute-map DENY_RETURNING_TAG on each ASBR to block re-injection.
- Verify correct metric-type (metric-type 1 vs2 ) and seed metrics.
- Optional IS-IS variant: swap EIGRP for IS-IS on FRR containers — proves the technique is vendor-neutral. The tag mechanism is identical.
Prove it:
- show ip route ospf on the EIGRP side shows E2 routes with tag 200.
- No "ghost" EIGRP routes originating via the OSPF path in show ip route eigrp .
- Shut a redistribution link — no microloop during convergence.
Admin-distance comparison table (write in your notes):
| Protocol | Cisco AD | MikroTik distance | Juniper preference | 
|---|---|---|---|
| Connected | 0 | 0 | 0 | 
| Static | 1 | 1 | 5 | 
| eBGP | 20 | 20 | 170 | 
| OSPF intra | 110 | 110 | 10 | 
| IS-IS L1 | 115 | 115 | 15 | 
| EIGRP internal | 90 | — | — | 
| iBGP | 200 | 200 | 170 | 
GNS3vault labs (gns3vault-archive/EIGRP/) — start with these three before the named-mode extension:
| Lab | What to do | 
|---|---|
| eigrp-basic | Classic-mode EIGRP baseline, DUAL verification. | 
| eigrp-authentication | MD5 key-chain — foundation before SHA upgrade. | 
| eigrp-summarization | Summarization, auto-summary disable, discard route. | 
Named-mode extension (hand-written): GNS3vault uses classic-mode EIGRP. Rebuild the eigrp-basic topology in named-mode:
Tasks:
- Convert to named-mode: router eigrp LAB4 /address-family ipv4 unicast autonomous-system 100 /address-family ipv6 unicast autonomous-system 100 . Both AFs on a single adjacency.
- Replace MD5 authentication with SHA-256 key-chain (requires named-mode). Rotate one key mid-lab — confirm zero downtime.
- Configure stub routing on spokes: eigrp stub connected summary . Verify the hub receives only stub-type queries.
- SIA scenario: on a stub spoke, remove a route from the RIB manually with a null0 static. On the active querying router, observeshow ip eigrp topology active . Fix with proper stub configuration.
Prove it:
- show eigrp protocols shows named-mode, AF IPv4 + IPv6.
- show ip eigrp neighbors detail shows SHA-256 auth mode.
- show ip eigrp topology active is clean after the fix.
GNS3vault labs (gns3vault-archive/OSPF/) — complete these first as the OSPFv2 foundation:
| Lab | What to do | 
|---|---|
| ospf-authentication | Authentication mechanics — key-chain concept. | 
| ospf-nssa-not-so-stubby-area | NSSA Type-7/Type-5 conversion — needed in the OSPFv3 lab too. | 
| ospf-virtual-link | Virtual link mechanics — you'll recreate the problem in OSPFv3 too. | 
OSPFv3 address-family lab (hand-written): GNS3vault has no OSPFv3 AF lab.
Topology: 3 CSR1000v in a triangle, dual-stack (10.X.0.0/24 + fd00:X::/64).
Tasks:
- Single OSPFv3 process with two AFs: address-family ipv4 unicast andaddress-family ipv6 unicast . One adjacency carries both.
- SHA-256 key-chain authentication (mandatory for OSPFv3 in modern IOS — no plain-text option).
