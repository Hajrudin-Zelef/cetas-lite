---
id: collect-261001-general-networking/general-networking/t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e-3
title: "t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e.md
source_anchor: ""
source_lines: [248, 349]
sha256: e15cebb9afb635adfb1b20c81a8378c27cbb942d5db4dc75d05deabcae58a763
---

# t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e

Hope this helps.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-07-2020 04:11 AM
Hello
@medharrak wrote:
However when I enable "IP Source Guard" nothing works even PCs on the same VLAN on the same Switch can't reach each other.
As far as I know IP source Guard wroks with DHCP snooping to verify the source IP address and it should be enabled on untrusted ports, if the source IP is not on DHCP snooping binding than the packet is dropped.
FYI -  What do the system related logging show for dhcp /IPSG.
sh ip source binding
sh ip dhcp snooping binding
sh log
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-11-2020 01:25 AM
MacAddress IpAddress Lease(sec) Type VLAN Interface
------------------ --------------- ---------- ------------- ---- --------------------
C0:02:18:66:00:00 100.1.1.3 86078 dhcp-snooping 100 GigabitEthernet0/1
C0:04:18:8D:00:00 100.1.1.2 86059 dhcp-snooping 100 GigabitEthernet0/2
C0:05:18:9B:00:00 200.1.1.3 86058 dhcp-snooping 200 GigabitEthernet0/2
C0:03:18:77:00:00 200.1.1.2 86058 dhcp-snooping 200 GigabitEthernet0/1
Total number of bindings: 4
AS-1#show ip source binding
MacAddress IpAddress Lease(sec) Type VLAN Interface
------------------ --------------- ---------- ------------- ---- --------------------
C0:04:18:8D:00:00 100.1.1.2 85738 dhcp-snooping 100 GigabitEthernet1/0
C0:05:18:9B:00:00 200.1.1.3 85737 dhcp-snooping 200 GigabitEthernet1/1
Total number of bindings: 2
AS-2#show ip source binding
MacAddress IpAddress Lease(sec) Type VLAN Interface
------------------ --------------- ---------- ------------- ---- --------------------
C0:02:18:66:00:00 100.1.1.3 85727 dhcp-snooping 100 GigabitEthernet1/0
C0:03:18:77:00:00 200.1.1.2 85707 dhcp-snooping 200 GigabitEthernet1/1
Total number of bindings: 2
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-12-2020 04:55 AM
Hello
Remove Dhcp snooping and IPSG off CS and just have it running on AS-1 & AS-2
On AS-1 apply the following to Gi0/0, Gi0/2
On AS-2 apply the following to Gi0/0, Gi0/1
interface
ip arp inspection trust
ip dhcp snooping trust
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-12-2020 11:03 PM
I did and still the same issue.
I don't understand why I should Enable IP ARP Inspection on AS-1 and AS-2? I'm trying to configure IP DHCP with IPSG not arp inspection , they are not the same , right ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-12-2020 11:50 PM
Hello
Your right, I was thinking of IPSG but for some reaon mentioned DAI...
So for claridication you DONT need trust any interface for DAI ...Apologies
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-12-2020 11:54 PM
What model switches are you using?
The reason I ask is because normally Gi0/[1-48] would be the copper ports used for access, while Gi1/[1-4] would be the SFP ports used for uplink ports on something like a Catalyst 3650. Is it possible that you're mixing these up?
