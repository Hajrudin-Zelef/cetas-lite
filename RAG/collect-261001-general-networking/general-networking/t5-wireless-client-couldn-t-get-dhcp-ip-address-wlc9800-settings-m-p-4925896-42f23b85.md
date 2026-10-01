---
id: collect-261001-general-networking/general-networking/t5-wireless-client-couldn-t-get-dhcp-ip-address-wlc9800-settings-m-p-4925896-42f23b85
title: "t5-wireless-client-couldn-t-get-dhcp-ip-address-wlc9800-settings-m-p-4925896-42f23b85"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/t5-wireless-client-couldn-t-get-dhcp-ip-address-wlc9800-settings-m-p-4925896-42f23b85.md
source_anchor: ""
source_lines: [1, 119]
sha256: 01b21b929f2df5a944b3674afe5b8a2e9f344a0918624d0d58f980b8d34dce46
---

# t5-wireless-client-couldn-t-get-dhcp-ip-address-wlc9800-settings-m-p-4925896-42f23b85

Client couldn't get DHCP IP address - WLC9800 settings
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 07:43 AM
Hello Professionals,
Could you please see and figure out my incorrect configurations?
I'm testing with WLC9800, but client device can't get IP address. 
DHCP - Cisco 9300L switch - SVI (ip dhcp pool for each VLAN)
 * VLAN14: 10.50.14.0/24 (SVI - 10.50.14.1)
Client> When I connect AP, it says connected (authenticated - macfiltering)
WLC9800>> 
I have configued correct VLAN - Profile and DHCP server IP.
However, when I connect this SSID, I can't get IP address from DHCP Server.
What did I miss..?
I appreciate your comments.
- Labels:
- 
						
							
		
			Wireless LAN Controller
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 08:30 AM
 - Have a checkup of the WLC9800 configuration with  the CLI command show tech wireless ; feed the output into :
                                https://cway.cisco.com/wireless-config-analyzer/
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
09-19-2023 08:43 AM
where is your vlan 14 SVI configured - check the SVI configuration and add IP helper address as example below :
interface Vlan14
description **SSID**
ip address x.x.x.x 255.255.255.0
ip helper-address a.a.a.a
ip helper-address b.b.b.b
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 09:12 AM
I have ip helper-address config.
interface Vlan14
description Corporate-WIRELESS
ip dhcp relay information trusted
ip address 10.50.14.254 255.255.255.0
ip helper-address 10.50.14.1
However, I got this errors from Cisco Switch (DHCP):
%DHCP_SNOOPING-5-DHCP_SNOOPING_NONZERO_GIADDR: DHCP_SNOOPING drop message with non-zero giaddr or option82 value on untrusted port, message type: DHCPDISCOVER, MAC sa:
Maybe that's why I can't get the IP address? Do you think if I configure dhcp snooping trusted port then it works?
Thank you so much.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 09:23 AM
After I configured the port to be trusted, it works.
Thank you everyone!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 09:25 AM - edited 09-19-2023 09:28 AM
ip dhcp relay information trusted - this should fix the issue for that error.
what DHCP Server (Microsoft - what DHCP Option 82 ?)
you can check #show ip dhcp snooping statistics detail
some good document :
https://community.cisco.com/t5/switching/dhcp-snooping/td-p/1622877
you can use below options :
ip dhcp snooping information option allow-untrusted
no ip dhcp snooping verify no-relay-agent-address
no ip dhcp snooping information option
Suggest to run the Debug :
debug ip dhcp server packet
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
