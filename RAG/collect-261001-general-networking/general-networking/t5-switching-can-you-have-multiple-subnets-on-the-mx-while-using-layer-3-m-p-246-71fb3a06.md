---
id: collect-261001-general-networking/general-networking/t5-switching-can-you-have-multiple-subnets-on-the-mx-while-using-layer-3-m-p-246-71fb3a06
title: "t5-switching-can-you-have-multiple-subnets-on-the-mx-while-using-layer-3-m-p-246-71fb3a06"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-can-you-have-multiple-subnets-on-the-mx-while-using-layer-3-m-p-246-71fb3a06.md
source_anchor: ""
source_lines: [1, 94]
sha256: 3c4091941a7f341afe037c5f73355b38d05f70c4b098469ed1ff4bb3f0a90147
---

# t5-switching-can-you-have-multiple-subnets-on-the-mx-while-using-layer-3-m-p-246-71fb3a06

Can you have multiple Subnets on the MX while using Layer 3 Switching?
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2024 05:06 AM
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2024 10:08 AM
You can do this, we normally keep the MGMT VLAN and isolated VLANs like the one you are describing on the MX for exactly the security reasons you describe.
You may lose some performance on inter vlan traffic as there will be an extra hop to get between VLAN 20 and VLANs 10/30 but it may not be noticeable. That's really going to be dependent on your overall load.
The routing will be fine provided your default route on the L3 is the MX. If it's not you would just have to add  static route on the L3 pointing VLAN 20 to the MX.
One thing to note is you will have to move/recreate DHCP for VLAN 20 if the Merakis are handling it.  If you have a lot of reservations it can be easier to dump them with the API and reupload that way. 
Otherwise I can't think of anything else to note.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2024 10:15 AM
Thank you for your response!
I followed these steps exactly today and for some reason I couldn't ping any device that picked up an IP address from the MX on VLAN20.
L3 Switch had default route to the MX transit VLAN IP. 
But no dice.
The switches that were downstream of the L3 switch, on the management vlan which is also configured on the MX. Could ping the VLAN 20 gateway and work perfectly fine. But not a client on VLAN20... I double checked to make sure no ACLs, No Firewall  rules were causing an issue and nothing. 
I must be missing something.
Do I need a static route on the MX somewhere to point traffic back to these devices? 
I can't configure a static route to a subnet thats residing on the MX of course so I'm not really to sure what I am missing.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2024 12:12 PM
You do not need a static route on the MX for VLAN 20 as it exists locally.  If you can ping the VLAN 20 gateway remotely my suspicion would be either dhcp is handing out the wrong gateway ip to the devices OR there is a firewall thing you are missing.
It's all Meraki so I'd suggest a quick call to support. They may be able to track it down faster.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2024 12:13 PM
Did you add vlan 20 on the trunk between mx and core switch?
Did you remove the static route for vlan20 from the mx to the coreswitch?
Did you remove the vlan 20 svi from the switch?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-05-2024 01:43 AM
Did you add vlan 20 on the trunk between mx and core switch?
-The trunk link between the Core switch (Catalyst-M) C9300X-12Y is tagged on Management Vlan 5 with allowed VLANs 1-1000
Did you remove the static route for vlan20 from the mx to the coreswitch?
Yes, this was a pre-requisite otherwise I was not able to add the Subnet onto the MX
Did you remove the vlan 20 svi from the switch?
Yes, I removed this.
Core switch has one route on it
0.0.0.0/0 next hop of the MXs transit IP address.
