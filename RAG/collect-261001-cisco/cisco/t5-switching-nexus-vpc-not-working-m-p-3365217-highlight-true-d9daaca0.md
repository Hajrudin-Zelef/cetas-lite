---
id: collect-261001-cisco/cisco/t5-switching-nexus-vpc-not-working-m-p-3365217-highlight-true-d9daaca0
title: "t5-switching-nexus-vpc-not-working-m-p-3365217-highlight-true-d9daaca0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "throughput"]
source: docs/RAG/collect-261001-cisco/t5-switching-nexus-vpc-not-working-m-p-3365217-highlight-true-d9daaca0.md
source_anchor: ""
source_lines: [1, 155]
sha256: 3cc7ec0447012cc084ded221b051a88e7b1935a661cebf76845490e7fae9b36c
---

# t5-switching-nexus-vpc-not-working-m-p-3365217-highlight-true-d9daaca0

Nexus VPC not working
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 09:00 AM - edited 03-08-2019 02:37 PM
Have 2x 3548s, configured VPC and it is not working.
Switches ping each other - Sw1 is 10.10.10.1 and Sw2 is 10.10.10.2 (SVI10)
I am using SVI 10 for VPC on both switches.
Here pertaining configs.
Vpc domain 10
role priority 10 (20 for switch2)
peer-keepalive dest 10.10.10.2 source 10.10.10.1
interface eth 1/45-46 (both switches)
Des vPC Peer-Link
channel-group 10 mode active
switchport 
switchport mode trunk
vpc peer-link
interface port-channel 10 (both switches)
description vPC Peer-Link
no shut
switchport 
switchport mode trunk
vpc peer-link
exit
Any help guys?
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
04-12-2018 09:55 AM
You should not use vPC peer-link and vPC keep-alive over the vPC peer link.
vPC peer-link is for all your vlans on both switches. For vPC keep-alive, you can use the mgmt0 interface on each switch. These interfaces should be connected to a 3rd switch in the same vlan.
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 12:36 PM
Oh ok.
I have the peer-keepalive configure under the vpc domain 10 config.
I am a bit confused. Should I configure that under:
int mgmt0
peer-keepalive dest 10.10.10.2 source 10.10.10.1
?
Also, if I don't have another switch to use can I configure that under let's say svi10?
I am not at the devices right now to try any commands.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 12:48 PM
Have a look at this document and look at figure-3 (vPC concept).
It shows how the peer-link and keep-alive links should be connected and configured.
You want to configure it under mgmt0 interface but that interface needs to be up and running first.
https://www.cisco.com/c/en/us/products/collateral/switches/nexus-5000-series-switches
/configuration_guide_c07-543563.html
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 01:46 PM
I noticed in step 6 it is simply:
N5k-1(config-vpc-domain)# peer-keepalive destination 172.25.182.52
but step 4 of table one has:
| Step 4 | peer-keepalive destination ipaddress [hold-timeout secs \| interval msecs {timeoutsecs} \| {precedence {prec-value \| network \| internet \| critical \| flash-override \| flash \| immediate priority \| routine}} \| tos {tos-value \| max-reliability \| max-throughput \| min-delay \| min-monetary-cost \| normal}} \|tos-byte tos-byte-value} \| source ipaddress \| vrf {management \| default}]  Example: Management interface for peer keepalive link: switch(config-vpc-domain)# peer-keepalive destination 172.28.230.85 switch(config-vpc-domain)# SVI for peer keepalive link: switch(config-vpc-domain)#peer-keepalive destination 172.28.1.100 source 172.28.1.120 vrf default | 
Thanks bro!
Couple questions (not at the nexus right now to try configs).
Do I really need to configure all those different variables when configuring keep-alive?
Also, it states I can use an SVI for peer keep alive. If I am using let's say SVI does the following look correct for lets say SVI10 (10.10.10.1 for Sw1 and 10.10.10.2 for Sw2)?:
switch1(config-vpc-domain)#peer-keepalive destination 10.10.10.2 source 10.10.10.1 vrf default
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 01:57 PM
Yes, you can use an SVI.
Here are all the commands you need for each device.
Primary switch
vpc domain xx
peer-switch
role priority 100
peer-keepalive destination 172.28.1.100 source 172.28.1.120
delay restore 150
peer-gateway
ip arp synchronize
Secondary switch:
vpc domain xx
peer-switch
peer-keepalive destination 172.28.1.120 source 172.28.1.100
delay restore 150
peer-gateway
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-16-2018 01:56 PM
Thanks!
Is there anyway to use an SVI on a 3548X Nexus for peer keep-alives instead of using the management IP and adding another switch?
I tried to create:
vrf context VPC700
int1/45
but won't take "vrf member VPC700" command.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-18-2018 07:26 AM
I still can't get everything to work for some reason.
