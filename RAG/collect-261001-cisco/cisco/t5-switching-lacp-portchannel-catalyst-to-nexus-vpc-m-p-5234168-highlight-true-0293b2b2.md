---
id: collect-261001-cisco/cisco/t5-switching-lacp-portchannel-catalyst-to-nexus-vpc-m-p-5234168-highlight-true-0293b2b2
title: "t5-switching-lacp-portchannel-catalyst-to-nexus-vpc-m-p-5234168-highlight-true-0293b2b2"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/t5-switching-lacp-portchannel-catalyst-to-nexus-vpc-m-p-5234168-highlight-true-0293b2b2.md
source_anchor: ""
source_lines: [1, 100]
sha256: d50d6de6e8fe3bc2354432c1bf67a8d0c837c2640c044240ce5eafe7bce229f5
---

# t5-switching-lacp-portchannel-catalyst-to-nexus-vpc-m-p-5234168-highlight-true-0293b2b2

LACP Portchannel Catalyst to Nexus vPC
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-08-2024 09:43 AM
Dear community,
I have a LACP Portchannel configured on a Catalyst connected to a Nexus vPC.
When I shutdown one Port of the vPC the connected port on the catalyst goes into lacp suspended state.
That´s expected, as the port doesnt receiev any lacp pdus anymore.
But then a ping from a host connected to the catalyst fails and the traffic isn´t going over the other link in the channel.
Only when i shutdown the port(which was suspended) then the ping works again.
Do I miss any command to influence this behaviour?
I want a proper failover when one port is suspended.
Thanks in advance
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
12-08-2024 10:44 AM
Depending on the configuration you applied between Catalyst and nexus
On the nexus side, you must have a vPC configuration under the port channel to work as expected. If the catalyst switch is connected to a dual home with a Nexus dual-configured vPC setup.
In most cases, it should not suspend the port. (what reason is it suspending ?)
https://community.cisco.com/t5/networking-knowledge-base/nexus-vpc-recommendations/ta-p/3130797
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-08-2024 12:12 PM
Hello @SlaSla
In a properly functioning LACP PortChannel, traffic should automatically fail over to the active link when one link is suspended. The problem suggests there may be a configuration mismatch or a default setting that needs to be adjusted...
Firstly, ensure that the Po is configured with the correct minimum links setting. LACP allows you to specify the minimum number of operational links required for the PortChannel to remain active. If this value is higher than 1, the Po might become inactive when one link is lost. Verify and configure the setting using the command:
interface port-channel X
 port-channel min-links 1
This ensures that the PortChannel remains operational as long as at least one link is active.
Secondly, check the load balancing configuration on both the Catalyst and Nexus switches. Inconsistent or improper 'hash' algorithms can lead to uneven traffic distribution or incorrect forwarding. On the Catalyst switch, configure load balancing with an appropriate algorithm, such as src-dst-ip, to ensure traffic is properly distributed across the active links:
port-channel load-balance src-dst-ip
Lastly, verify that LACP mode is consistently configured on both ends of the Po. For example, ensure that both the Catalyst and Nexus are using active mode for LACP. A mismatch in LACP settings can result in one side treating a link as active while the other treats it as suspended.
--Use the command show etherchannel summary on the Catalyst and show port-channel summary on your Nexus to confirm operational consistency.
You could use debugging commands such as debug lacp on the Catalyst. This can help identify if there are missed PDUs or other protocol-level issues contributing to the behavior.
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-08-2024 01:39 PM
Hi, thanks for your helpful information.
I can reproduce the issue in a much simpler topology ( see screenshot)
The portchannel is established properly.
I am pinging 10.1.1.9 > 10.1.1.10
When I shutdown Gi0/0 on SW7 the ping fails.
The reason is that even Gi0/0 is suspended the echo replys are sent via this port.
Perhaps the virtual switches don´t behave correct regarding lacp !?!?
Switch config regarding lacp is identical on both switches:
Thanks
Stephan
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-09-2024 12:01 AM
in SW try add below
no lacp suspend-individual
MHM
