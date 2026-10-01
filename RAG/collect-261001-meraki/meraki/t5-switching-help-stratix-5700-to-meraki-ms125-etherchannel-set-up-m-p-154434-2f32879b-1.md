---
id: collect-261001-meraki/meraki/t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b-1
title: "t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b.md
source_anchor: ""
source_lines: [1, 18]
sha256: 711a7b31a6ab96403d7c99eccdb3d79f17d3f1895da5129b7a329aae46f12b7c
---

# t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-11-2022 12:20 PM
I am trying to configure an EtherChannel between a Stratix 5700 (Cisco product) and a Meraki MS125.
Stratix 5700 config: STP is set as MSTP and am using 3 ports. All 3 ports are configured as: Auto Speed, Auto Duplex, Trunk, Native VLAN 3, Allow all VLANs. The I combined them into an EtherChannel using device manager. FA1/14, FA1/15, and FA1/16 Channel# =1, Channel Mode = LACP (Active). This creates a new port called Po1. Po1 configuration: Auto Speed, Auto Duplex, Trunk, Native VLAN 3, Allow all VLANs.
Meraki MS125 config: STP set to RSTP. 3 ports 13,14, and 15 all set as Auto Speed, Auto Duplex, Trunk, Native VLAN 3, Allow VLANS 1,3,11,12 - choose all 3 ports the select Link Aggregation.
When I do this the link comes up with connectivity and then fall out. I get an error message in the syslog of the 5700 saying %PM-4-ERR_DISABLE: Channel-misconfig (STP) error detected on Po1, putting Po1 in err-disable state. Any advise on setting up the Meraki side would be greatly apricated. The only options on the 5700 side is to choose the Channel Mode. If it should be something other than LCAP (Active) let me know, please. Thank you in advance for your help.
Solved! Go to Solution.
- Labels:
- 
						
							
		
