---
id: collect-261001-cisco/cisco/t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by-b0cf47c4-2
title: "t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by--b0cf47c4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by--b0cf47c4.md
source_anchor: ""
source_lines: [15, 195]
sha256: 1e77deaeb1c0d585caaa7ae4431df3082e6f0c83e16a75ebd931fcde344542e6
---

# t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by--b0cf47c4

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-17-2026 12:20 PM - edited 05-17-2026 12:27 PM
Hello community,
I am experiencing an intermittent connectivity issue in an environment with Cisco ASA firewalls in failover mode and Nexus 5K switches configured with vPC. I would appreciate any insights or confirmation of possible root cause scenarios.
Environment Overview
Firewalls: Cisco ASA in Active/Standby failover
ASA version: 9.12(4)
Core switches: 2x Cisco Nexus 5K (vPC configured)
Production gateway: HSRP virtual IP → 10.10.17.5
Connectivity: ASA connected to Nexus via Port-Channel (trunk links)
Approximately every ~20 minutes:
- The ASA loses connectivity to the production network
- It cannot reach the HSRP gateway (10.10.17.5)
- No configuration changes or failover events occur during the issue
- No physical link failures or interface flapping are observed
From the ASA perspective, the gateway becomes unreachable until intervention is performed.
Temporary Workaround
Connectivity is immediately restored after executing:
clear arp
After clearing the ARP table:
- Communication to 10.10.17.5 is restored
- Normal operation resumes
- The issue reoccurs again after ~20 minutes
Observations
- No routing changes detected during the event
- No interface down/up events observed
- HSRP gateway remains active and reachable from other devices
- Issue is isolated to ASA-to-production network communication
- Behavior strongly suggests ARP-related instability
Current Hypotheses
1. ARP inconsistency toward HSRP VIP
The ASA appears to retain or fail to refresh a stale ARP entry for 10.10.17.5, and clear arp forces relearning.
2. Asymmetric routing / vPC forwarding inconsistency
Traffic may be entering and returning through different Nexus switches, potentially causing inconsistent MAC/ARP learning behavior.
3. vPC or L2 forwarding inconsistency on Nexus
Possible issues include:
- VLAN inconsistency across vPC peers
- MAC address learning issues
- Inconsistent forwarding paths between Nexus switches
4. HSRP virtual MAC / ARP aging interaction
The ASA may not properly refresh the ARP entry associated with the HSRP virtual MAC, leading to intermittent loss of connectivity.
- Has anyone observed similar behavior where an ASA loses connectivity to an HSRP VIP and recovers only after clear arp?
- Can vPC inconsistencies in Nexus switches lead to ARP instability or stale MAC/ARP entries on downstream firewalls?
- Are there any known issues or bugs in ASA 9.12 related to:
  - ARP aging
  - HSRP MAC resolution
  - Intermittent ARP refresh failures
- What specific checks would you recommend on the Nexus side (vPC consistency, STP state, MAC learning behavior)?
Many thanks in advance.
Regards
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-17-2026 10:59 PM
   - @Scott12                 - What is the full model name of the ASA ?
                                      - Setup a common syslog server for the ASA and nexus; (and all cisco equipment;)
                                         that will give the opportunity to collect logs at a central space for further analysis
  M.
-- ' Listen to the wind, it talks
Listen to the silence, it speaks
Listen to your heart, it knows
Ganado Mucho (1809 to 1893 ) Navajo Indian
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 01:41 AM
Did you check arp cache after problem occurs and search mac address in switch forwarding database "show mac address-table" ? Seems like there is a duplicate IP-address in the network.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 01:58 AM
Hello @Scott12 
Sounds duplicate address or ARP flooding issue.
You can note the mac address of the next hop while everything is running, when the problem occurs follow up the stale mac address in ASA from L2 perspective to identify from what port is responding with ARP reply to 10.10.17.5.
if you could not identify the issue create a static arp entry on ASA to map 10.10.17.5 to NH mac address as a temporary option to avoid clearing arp each time, until you involve TAC or share the logs in the community.
Regards!
Amine ZAKARIA
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 03:46 AM
Thanks for the feedback.
Here are additional details from my troubleshooting:
- ASA model: Cisco ASA 5525-X
- Failover is working correctly and stateful sync looks healthy
- No failover or interface flaps occur during the issue
I performed further checks on the Nexus (vPC environment), and I found something interesting:
- The VLAN used by the ASA (production VLAN) is NOT consistently allowed on both vPC peers
- Example:
  - On one Neus, VLAN is present on the Port-Channel toward the ASA
  - On the other Nexus, the same VLAN is missing from the trunk
Additionally:
- MAC address learning is inconsistent between both Nexus switches
- The ASA-facing Port-Channel (Po52) shows different MAC visibility depending on the Nexus
- During tests, I also observed very high output errors on one physical member of the Port-Channel
This suggests a possible Layer 2 forwarding inconsistency, potentially causing:
- Incorrect or asymmetric MAC learning
- ARP replies not reaching the ASA consistently
- Stale ARP behavior (symptom, not root cause)
Regarding the ARP behavior:
- When the issue occurs, clearing ARP on the ASA immediately restores connectivity
- This reinforces that the ASA is likely keeping a valid but unusable MAC entry
- However, the root cause seems to be inconsistent L2 forwarding in the Nexus/vPC domain
Have you seen similar behavior where vPC VLAN inconsistency or MAC learning issues cause intermittent ARP reachability problems toward an HSRP VIP?
Thanks again!.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 03:57 AM
- @Scott12
- The VLAN used by the ASA (production VLAN) is NOT consistently allowed on both vPC peers
- Example:  On one Nexus, VLAN is present on the Port-Channel toward the ASA
 On the other Nexus, the same VLAN is missing from the trunk
Correct these items
M.
-- ' Listen to the wind, it talks
Listen to the silence, it speaks
Listen to your heart, it knows
Ganado Mucho (1809 to 1893 ) Navajo Indian
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 04:20 AM
Hello,
The vlans config in PO should be symmetric. Btw did you verify if any HSRP failovers ?
Regards!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 05:53 AM - edited 05-18-2026 06:03 AM
Hello there!
That actually aligns with something I found during troubleshooting.
I identified a VLAN inconsistency on the Port-Channel toward the ASA:
VLAN 119 (and potentially others) was allowed on the Port-Channel on one Nexus (TC), but missing on the peer Nexus (GC).
As a result, the ASA-facing Port-Channel (Po52) was not symmetric across vPC peers.
This has now been corrected to ensure VLAN consistency on both sides.
However, the intermittent issue (loss of connectivity to HSRP VIP every ~20 minutes) is still occurring even after fixing the VLAN mismatch.
Additional observations:
No HSRP failovers are occurring (state remains stable).
Failover on ASA is healthy and stateful sync is working.
Issue is resolved immediately after clear arp on the ASA.
MAC address-table inspection shows inconsistent learning paths for some VLANs across vPC peers.
Given this, I suspect there might still be:
MAC/ARP instability or duplication in the L2 domain
Or asymmetric forwarding behavior across the vPC
- Mark as New
- Bookmark
