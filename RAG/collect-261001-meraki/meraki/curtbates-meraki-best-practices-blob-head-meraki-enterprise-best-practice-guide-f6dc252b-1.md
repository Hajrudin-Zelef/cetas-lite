---
id: collect-261001-meraki/meraki/curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide-f6dc252b-1
title: "curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "throughput", "voice"]
source: docs/RAG/collect-261001-meraki/curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b.md
source_anchor: ""
source_lines: [1, 59]
sha256: 3befaeb37ceef327b97018f702b763ee1ffdece0f8081c36ae6013f4b85a7427
---

# curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b

Scope: Wired (MS), Wireless (MR), SD-WAN (MX), and Security best practices for a cloud-managed enterprise network. Source basis: Cisco Meraki Best Practice Design documentation (Dashboard Administration → Architectures and Best Practices). Audience: Network, security, and wireless engineers and dashboard administrators.
This guide consolidates Meraki's published design recommendations into a single reference. Every deployment is unique — validate sizing, addressing, and policy against your own requirements, and finalize designs with your Meraki SE or partner before production cutover.
Before configuring any individual product, get these organization-wide decisions right. They are the hardest things to change later.
Organization and network structure. Auto VPN, templates, and most reporting are scoped to a single dashboard organization. Devices in separate organizations generally represent separate SD-WAN domains and cannot form Auto VPN to one another. Plan your org boundary around the administrative and SD-WAN domain you actually want, not around geography alone.
Addressing and segmentation strategy (decide once, apply everywhere).
- Use multiple VLANs/subnets segmented by traffic role (employee, voice, guest, IoT, management). Segmentation is your first and cheapest security control and it reduces broadcast scope.
- Always isolate guest and untrusted traffic onto their own VLANs with no path to business VLANs by default.
- Standardize on /24 subnets for general user access (large enough for most sites, leaves room to grow and re-subnet). Use /23 where a single user segment genuinely needs more hosts.
- Never overlap subnets across the routed domain — overlaps cause inconsistent routing/forwarding.
- Reserve a dedicated management VLAN that is separate from any user data VLAN.
Firmware. Keep MX/MS/MR on the firmware trains your Meraki SE recommends, and standardize the same build across members of a stack or HA pair before provisioning them together.
| Setting | Recommendation | 
|---|---|
| RSTP | Enabled (default). Leave on; disable only after deliberate analysis. | 
| Root bridge | Force the root deterministically — set priority to "0 / likely root" on a stable core/aggregation switch (MS410/MS425 class). Choose a switch that rarely changes config or sees link flaps. | 
| STP diameter | Keep under 7 hops end to end. | 
| BPDU Guard | Enable on all access ports (user/server edge) to block rogue switches. | 
| Loop Guard | Enable on trunk links between switches. | 
| Root Guard | Enable on ports facing switches outside your administrative control. | 
| Catalyst/Nexus interop | Allow VLAN 1 on the Catalyst↔MS trunk (required for RSTP), and make the Catalyst the root when mixing PVST. | 
Trunks.
- Prune unnecessary VLANs from each trunk's allowed list to limit flooding scope.
- Ensure the native VLAN and allowed-VLAN list match on both ends — a native VLAN mismatch bridges traffic between VLANs.
- Use port tags (e.g., trunk ,wireless ,uplink ) for operational clarity.
Link aggregation. Only LACP is supported — confirm the far end runs LACP. Configure the aggregate in the dashboard before physically cabling the partner device.
Other L2 settings.
- MTU: leave at the default (9578) unless an intermediate device lacks jumbo-frame support; jumbo frames help server-to-server/app performance and avoid fragmentation.
- UDLD: enable on fiber trunks in Alert-Only mode.
- Link negotiation: auto-negotiate for Meraki-to-Meraki links; use forced mode only for a device that cannot auto-negotiate.
- Use L3 interfaces for data VLANs only. Do not put an L3 interface on the management VLAN — keep management traffic separate.
- In a switch stack, make sure the management IP subnet does not overlap any configured L3 interface subnet (overlap causes loss when pinging/SNMP-polling stack members). Exception: MS390 is not subject to this.
- L3 table changes on most MS models (MS210/225/250/350/355/410/425/450) flush and rebuild hardware tables, causing momentary disruption — make L3 changes only in a maintenance window.
ACLs.
- Summarize addresses wherever possible to conserve entries.
- Hard limit: 128 ACEs per network. Profile your real traffic first and always permit traffic to the Meraki dashboard (see Help → Firewall Info).
OSPF (Switching → Configure → OSPF Routing).
- Use broadcast mode for hellos; leave hello/dead timers at 10s/40s unless you've tested more aggressive values.
- All areas must attach directly to Area 0 (no virtual links).
- Set an explicit Router ID for manageability and enable MD5 authentication.
- To avoid OSPF forming needless adjacencies on every VLAN of a trunk, define a dedicated transit VLAN (typically core↔aggregation) with OSPF Passive = No, and set Passive = Yes on all other advertised subnets. Keep the management VLAN allowed on those trunks.
DHCP. Configure allowed DHCP servers to block rogue servers. Prefer stacking over warm spare for L3 DHCP, since warm-spare DHCP load balancing can hand clients to a member with no remaining leases.
- Switch stacking is the preferred HA/redundancy method for L3 switches — better redundancy and faster failover than warm spare.
  - Provision stacking carefully: add switches to dashboard without configuring the stack, bring them online, push the same latest firmware to all, then power-off and cable in a ring topology before provisioning the stack.
  - Use two uplink ports (one on the top, one on the bottom switch) and configure cross-stack link aggregation for uplink redundancy.
- Warm spare (VRRP) is available for L3 gateways where stacking isn't used. The pair shares a virtual MAC (00-00-5E-00-01-xx ) and virtual IP. Provide an alternate/direct path between primary and spare for VRRP exchange. L3-interface changes can briefly trigger VRRP transitions — make them in a change window.
- Classify traffic (commonly by VLAN — user/voice/network-control — and optionally refined by protocol/ports).
- MS acts on DSCP only (it does not honor 802.1p CoS in the dot1q header, and CoS markings are not preserved by default). If an endpoint can't mark DSCP itself, set it with a QoS rule.
- There are 6 weighted CoS queues (0–5) with weights 1/2/4/8/16/32. Bandwidth share under congestion = queue weight ÷ sum of configured queue weights. Plan capacity ahead — QoS only engages during congestion.
- Keep a dashboard network to ≤ 400 switches and < 8000 switch ports for reliable topology/port-page loading.
- For multi-gig links (2.5/5 Gbps), use Cat-6a for reliable operation; Cat-5e can work but is vulnerable to alien crosstalk/noise over longer runs and bundles.
- Where possible, keep source and receivers in the same VLAN and use IGMP snooping.
- Disable IGMP snooping if you have no L2 multicast need (it's CPU-bound).
- For multicast routing: place the PIM-SM rendezvous point close to the source (core/aggregation), configure an RP for every group, verify all multicast-routed switches can ping the RP, and ACL off noise like SSDP (239.255.255.250 ). Use the 239.0.0.0/8 range for internal apps. Always run an IGMP querier per multicast VLAN when there's no PIM-enabled device.
A coverage-only survey under-builds dense environments. Treat any area expecting >30 clients per AP as high density and size by capacity:
- Aggregate application throughput = per-user app throughput × concurrent users (compute per area: lobby, classroom, auditorium, etc.).
- Per-device usable throughput ≈ ~½ of the advertised data rate, then reduce a further ~30% for real-world overhead; size to the 20 MHz rate.
- APs by throughput = aggregate throughput ÷ per-device throughput.
- APs by client count = concurrent 5 GHz clients ÷ 25 (target ~25 clients/radio, ~50/AP). Assume roughly a 30/70 split of 2.4 GHz vs 5 GHz clients.
- Required APs = the larger of the two counts.
