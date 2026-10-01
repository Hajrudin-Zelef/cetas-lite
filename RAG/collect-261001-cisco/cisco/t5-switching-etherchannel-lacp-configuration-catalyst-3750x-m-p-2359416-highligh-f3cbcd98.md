---
id: collect-261001-cisco/cisco/t5-switching-etherchannel-lacp-configuration-catalyst-3750x-m-p-2359416-highligh-f3cbcd98
title: "t5-switching-etherchannel-lacp-configuration-catalyst-3750x-m-p-2359416-highligh-f3cbcd98"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-etherchannel-lacp-configuration-catalyst-3750x-m-p-2359416-highligh-f3cbcd98.md
source_anchor: ""
source_lines: [1, 115]
sha256: 34e9b68f769ba693bdc620b1dd4204428b427ea035550c54d02ecffaab8026de
---

# t5-switching-etherchannel-lacp-configuration-catalyst-3750x-m-p-2359416-highligh-f3cbcd98

EtherChannel (LACP) Configuration - Catalyst 3750x
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-05-2013 03:17 PM - edited 03-07-2019 04:57 PM
Greetings,
I have a couple questions for the group.
Q1) Should the port configuration options of each individual line of the etherchannel bundle also be reliazed on the logical port-channel interface? In the below example, should "spanning-tree portfast" and "spanning-tree bpduguard enable" be coded on the port-channel interface?
Q2) Are there any best practices, general rules, or nuances you can share regarding the configuration of etherchannels as it pertains to the above question? What about etherchannel trunks?
Thank You
========================
!
interface Port-channel3
description EtherChannel (LACP, 2-port) for Server1
switchport access vlan 999
switchport mode access
!
!
!
interface GigabitEthernet3/0/12
description Server1 NIC1
switchport access vlan 999
switchport mode access
spanning-tree portfast
spanning-tree bpduguard enable
channel-group 3 mode active
!
!
interface GigabitEthernet4/0/12
description Server1 NIC2
switchport access vlan 999
switchport mode access
spanning-tree portfast
spanning-tree bpduguard enable
channel-group 3 mode active
!
========================
.
- Labels:
- 
						
							
		
			Other Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-05-2013 05:39 PM
Hi,
When you configure an EtherChannel, configuration changes applied to the port-channel interface apply to all the physical ports assigned to the port-channel interface, and configuration changes applied to the physical port affect only the port where you apply the configuration.
Below you have a link with the configuration guidelines:
Hope this helps.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2013 11:20 AM
Thank you Leo for the reply and link.
Reading through the documentation you linked, I discovered this:
"NOTE After you configure an EtherChannel, configuration changes applied to the port-channel interface apply to all the physical ports assigned to the port-channel interface, and configuration changes applied to the physical port affect only the port where you apply the configuration."
This statement isnt quite clear to me. Does this mean that line-item configurations to the EtherChannel are propagated down to the individual ports? But line-item configuration changes made to indiviual ports will not propagate up the the EtherChannel?
So with the example I gave above, does Port-channel3 need to have explicit line-item configurations for "spanning-tree portfast" and "spanning-tree bpduguard enable"; and should it look like the following?
interface Port-channel3
description EtherChannel (LACP, 2-port) for Server1
switchport access vlan 999
switchport mode access
spanning-tree portfast
spanning-tree bpduguard enable
.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2013 12:10 PM
Hello
Yes you are correct in your example
L2/L3 etherchannel:
1) Apply the port-channel command to the designated interfaces - This will auto create the port-channel interface
2) apply the necessary interface commands to the port-channel - this will propagate to the physical interfaces.
3) Shutdown and restart the physical interfaces
If you change the physical interfaces after the port channel is created it will NOT propagate to the port channel
 Res
Paul
Sent from Cisco Technical Support iPad App
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
