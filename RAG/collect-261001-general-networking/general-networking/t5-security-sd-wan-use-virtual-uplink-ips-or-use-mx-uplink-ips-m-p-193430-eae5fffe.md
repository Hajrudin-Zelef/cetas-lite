---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-use-virtual-uplink-ips-or-use-mx-uplink-ips-m-p-193430-eae5fffe
title: "t5-security-sd-wan-use-virtual-uplink-ips-or-use-mx-uplink-ips-m-p-193430-eae5fffe"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-use-virtual-uplink-ips-or-use-mx-uplink-ips-m-p-193430-eae5fffe.md
source_anchor: ""
source_lines: [1, 136]
sha256: 4c5e82d9d92c852916f986b8aa299946e519cbdf97c639ad1865ae4b89546895
---

# t5-security-sd-wan-use-virtual-uplink-ips-or-use-mx-uplink-ips-m-p-193430-eae5fffe

Use virtual Uplink IPs or Use MX uplink IPs
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-02-2023 04:47 AM
Hi,
We have two MXs of the same model, and we want to design an HA for our environment. I am considering having WAN 1 (from ISP 1) on both the main and spare MX, and WAN 2 (from ISP 2) on both the main and spare MX.
ISP1-> (MXA_Wan1, MXB_Wan1)
ISP2-> (MXA_Wan2, MXB_Wan2)
The MX also has a virtual uplink setting.
Which option is the best? Does anyone have any ideas?"
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
05-02-2023 06:26 AM
In either way, you'll need two IP addresses from ISP A, and to addresses from ISP B - one for each WAN port.
The Virtual IP allows you to have an extra, third IP address that is shared between the two MX's WAN1 port. When using Warm Spare (which is based off VRRP), you'd have the third IP address which point towards the Active-Primary MX. Incase of a failure on the Primary MX, the Secondary MX will become Active, and take over the third IP address.
If you don't use the Virtual IP, in the event of a failure on the Primary MX, for VPN connections, you'll have to manually reconfigure endpoints to use the Spare MX WAN IP. In terms of sessions, your clients may also experience short outages, as all their TCP traffic will be reset, and connections have to be re-established as your Public IP address would have changed.
However, in order to obtain a third IP address for Warm Spare, it would require your ISP to atleast provide a /29 handoff, which in some cases may be a bit more difficult and more expensive.
For more details on Meraki Warm Spare I'd refer you to the documentation page here; https://documentation.meraki.com/MX/Deployment_Guides/MX_Warm_Spare_-_High_Availability_Pair
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-02-2023 11:24 PM
Thanks for your response. The idea behind having one IP address from each ISP is to have the same IP address when we experience a failure in the primary MX. Some services rely on the IP address, and we don't want to lose clients' access to those services due to a failure in the primary MX.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-02-2023 11:43 PM
I'm rather certain that having the same IP address on WAN1 of both MXs won't work.
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-03-2023 01:55 AM
I asked Meraki support, and this is their response:
Thank you for contacting Cisco Meraki support.
Yes, having 2 WAN uplinks on Primary and Spare MX is supported. Please refer to the following KB for recommended topologies:
https://documentation.meraki.com/MX/Deployment_Guides/MX_Warm_Spare_-_High_Availability_Pair#Recommended_Topologies
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-03-2023 11:11 PM
Yes, but I'm still rather certain, that it won't work if you reuse the same IP address from MX_A WAN1 on MX_B WAN1. This is what the Virtual IP is for. But you'll still need infividual IP addresses on each WAN interface. WAN1 and WAN2 can be different ISPs, but still need individual IP addresses.
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-19-2024 05:14 AM
Hi Everyone .
Just trying to understand how Failover works on WAN side . let say /29 Public subnet i have .
Internet Link is via UNTRUST-Switch ( layer 2 switch) and 2 MX boxes (WAN1 of MX-A and WAN1 of MX-B ) , ISP -Gateway connected to the UNTRUST Layer-2 switch . VIP ip used (common IP used ) the idea to not have the NAT public IP change so that TCP-Sessions will NOTget terminated . NAT is happening on the MX-A. Classic Outbound flows to INTERNET /North bound . EVENT = MX-A link to UNTRUST switch fails (LINK Failure) , or Huge packet loss (PL) /Packet drops on Link from MX-A connected to UNTREUST switch , how the failover is happen . I read in some articles VRRP hello mu;icast/hear beat packets will be sent viaLAN interface only , NOT via WAN interface . wondering how the other MX-B will take over the VIP /Floating ip /shared public . is there any HA-cluster link which will be used. MX-A will inform MX-B that WAN1 interface of MX-A is DOWN. secondly on the brownout scenario , MX-A 's WAN 1 interface having packetv loss , any mechanism , after 60 seconds , MX-B will become Active for the VIP -IP ( Public ip ) shared between the 2 MX-Boxes .
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-19-2024 05:16 AM
Just comparing with other SDWAN -products . SILVER PEAK , Juniper 128T SSR which i am currently working on . appreciate your valuable response .
Thanks
SUBU
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-02-2023 11:29 PM
Also this is the Meraki recommendation design for best practice HA.
https://documentation.meraki.com/MX/Deployment_Guides/MX_Warm_Spare_-_High_Availability_Pair
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-04-2023 02:59 AM
We have the same situation. Two different WAN connections on an HA pair of MX's. @Rasmus Hoffmann Birkelund is correct that you can't have the same IP address on separate WAN interfaces - this is what the virtual IP is for.
