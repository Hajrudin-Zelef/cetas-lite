---
id: collect-261001-general-networking/general-networking/t5-switching-how-to-configure-quot-management-vlan-quot-td-p-1642567-1df1d168
title: "t5-switching-how-to-configure-quot-management-vlan-quot-td-p-1642567-1df1d168"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-how-to-configure-quot-management-vlan-quot-td-p-1642567-1df1d168.md
source_anchor: ""
source_lines: [1, 147]
sha256: 34ba476534ddb58d486e9d34dfb057b6938cb7e8453b8613bff4abcc2b8b168d
---

# t5-switching-how-to-configure-quot-management-vlan-quot-td-p-1642567-1df1d168

How to configure " Management VLAN "
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2011 12:01 PM - edited 03-06-2019 03:59 PM
Dear All,
I have configured VLAN 2 as a Native Vlan. Now I want to configure VLAN 3 as a Management Vlan.
I have read from one article that we can configure Management Vlan by a command "management-vlan vlan-id". I tried this command on my cisco 2950 & 2960 Switches, but was not able to find.
I have read on the forums as well that how to configure the Management Vlan, I found that by assigning only the ip address to VLAN 3, we can configure Management Vlan.
I am too much confused about it that how to configure the Management Vlan.
Please guide me the exact way to configure the Management Vlan.
- Labels:
- 
						
							
		
			LAN Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2011 02:42 PM
Management vlan is nothing special its just the vlan you determine that you want to use on your switches . Just trunk the vlan you want to use down to your access swsitches and assign a ip address and the default gateway in that vlan . I have never seen that command that you mentioned and older switches like a 2950 would not have it .
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-10-2011 06:09 AM
Hi,
A sample would simply be like this:
Switch(config)#int vl3
Switch(config-if)#ip add 192.168.1.254 255.255.255.0
It would be helpful if you assign your switches on the same subnet for consistency and for troubleshooting ease.
Sent from Cisco Technical Support iPhone App
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-10-2011 02:43 PM
Hi,
I did the configuration according to the Mr glen.grant reply....... below is the configuration.
Switch(config)#vlan 3
Switch(config)#int fa 0/3
Switch(config-if)#switchport access vlan3
Switch(config)# int fa 0/3
Switch(config-if)# switchport mode trunk
Switch(config)# int vlan 3
Switch(config-vlan)#ip add 192.168.1.5 255.255.255.192
Switch(config)# ip default-gateway 192.168.1.1
I think, I did the right configuration ? But you can guide me wherever you feel, I did anything wrong or missing.
Thanks............. for all of you for your help..................
Message was edited by: NAVEED ANWER
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2021 10:18 AM
LOL!!!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2021 11:19 AM
This is an old thread but perhaps worth a little bit more discussion. To some extent the answer to the question of how to configure the management vlan depends on whether the switch is operating as layer 2 or as layer 3. If the switch is layer 2 then the only thing it needs an IP address for is management processing. So for the vlan that you want to be the management vlan you configure an SVI and assign an IP address and mask. And configure ip default-gateway so that the switch can reach remote subnets. And you should be sure that any trunk configured on the switch carries that management vlan.
If the switch is operating as layer 3 then there are probably multiple SVIs with IP addresses and any of them could be used for management purposes.
To be clear when we talk about a management vlan this is to enable the network administrators to access and to manage the switch, to telnet or SSH to the switch, perhaps HTTP/HTTPS if the switch has a GUI enabled, to send syslog messages to a syslog server, to respond to SNMP and to send traps if they are configured.
I believe that @glen.grant was spot on with his suggestion.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2023 11:19 AM
Nice explanation right there. Mine is short, I just want to confirm whether a layer 2 switch can have more than one management VLAN, by creating SVIs for maybe VLAN 10 and VLAN 20. My belief is that those in VLAN 10 should be able to access VLAN 10 SVI likewise those in VLAN 20 access VLAN 20 SVI and manage the switch.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-13-2023 06:51 PM
You have adopted a very old thread to append your question. But your question deserves a response. You asked "I just want to confirm whether a layer 2 switch can have more than one management VLAN, by creating SVIs for maybe VLAN 10 and VLAN 20". The answer is that it depends on the type of switch (and perhaps on the version of code that the switch is running). I have worked with some Catalyst switches that allow configuration of SVI for 2 interfaces. But I have also worked with many Catalyst switches that allow only ONE active SVI. On those switches if you configure a second SVI then the first SVI is put into the DOWN state.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-31-2017 02:39 AM - edited 12-31-2017 02:41 AM
This is an old thread but I had the same problem but found the solution. First my switch was running IOS 12.4. My management vlan that I wanted to always be up was vlan 2. I do NOT have any ports in vlan 2 but I needed to be able to reach this SVI at all times. I made sure my trunk port was set to allow vlan 2 as well as all the other vlans I needed.
Switch# Conf t
Switch(config)# int vlan 2
Switch(config-if)# ip add ###.###.###.### 255.255.255.0
Switch(config-if)# exit 
Switch(config)# vlan 2
Switch(config-vlan)# state active
Switch(config-vlan)# end
Switch# wr mem
That set the VLAN to always be active and brought the SVI with the IP address I assigned to an up/up state.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-20-2022 12:41 PM
This explanation was very helpful.
Thanks
