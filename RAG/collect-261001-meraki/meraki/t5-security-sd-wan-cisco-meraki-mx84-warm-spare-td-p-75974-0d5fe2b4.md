---
id: collect-261001-meraki/meraki/t5-security-sd-wan-cisco-meraki-mx84-warm-spare-td-p-75974-0d5fe2b4
title: "t5-security-sd-wan-cisco-meraki-mx84-warm-spare-td-p-75974-0d5fe2b4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-security-sd-wan-cisco-meraki-mx84-warm-spare-td-p-75974-0d5fe2b4.md
source_anchor: ""
source_lines: [1, 116]
sha256: e4def643756165b4890639ece3ca3502f383725c19fe6d6c8dd2595582c245e9
---

# t5-security-sd-wan-cisco-meraki-mx84-warm-spare-td-p-75974-0d5fe2b4

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2020 01:49 AM
Any body done Cisco Meraki MX warm spare.?
i got these two documents but not get cleared much?
https://www.willette.works/mx-warm-spare/
My doubt are
1.Do in need to register Meraki warmspare Box to dashboard before starting these procedures.
2.where can i add ip address of Meraki warmspare box.
Anybody have a configuration sample to share or much detailed configuration example guide??
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2020 02:04 AM
It is very simple, we have 9x HA pairs of MXs
Note: You need 2 or 3 IP addresses for each WAN/internet connection from your provider.
- Warm spare MX needs to be in inventory, but not in a network
- Go to Security & SD-WAN tab for the existing appliance and in the Warm spare box click on Configure warm spare
- There is a box with Disabled selected, click on Enabled
- Select the MX serial number that you want to become the warm-spare, it must be the same model as the existing.
- Depending on how many IP addresses you get from your WAN/internet providers choose Use MX uplink IPs from the Uplink IPs dropdown if you only have two, or choose Use virtual uplink IPs if you have three and enter them below.
- If you don't have DHCP available to plug the new warm-spare MX into, go to the local status page and enter the WAN IPs, subnet masks and gateways for the WAN link(s).
- If you do have DHCP that gives the MX internet access then you can configure the WAN IPs from the dashboard by clicking on the page below and then going to the Uplink tab
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2020 01:59 AM
I can recall having overthought this one too.
Yes, you‘ll have to add the Spare first of all. Then, simply go to the MX configuration page and create a Warm Spare pair. You‘ll be presented with the newly added spare MX and be led though the process. The IP will also be configured within this one.
Sorry that this is not a concise Howto document, but it‘s really simple.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-12-2020 02:04 AM
It is very simple, we have 9x HA pairs of MXs
Note: You need 2 or 3 IP addresses for each WAN/internet connection from your provider.
- Warm spare MX needs to be in inventory, but not in a network
- Go to Security & SD-WAN tab for the existing appliance and in the Warm spare box click on Configure warm spare
- There is a box with Disabled selected, click on Enabled
- Select the MX serial number that you want to become the warm-spare, it must be the same model as the existing.
- Depending on how many IP addresses you get from your WAN/internet providers choose Use MX uplink IPs from the Uplink IPs dropdown if you only have two, or choose Use virtual uplink IPs if you have three and enter them below.
- If you don't have DHCP available to plug the new warm-spare MX into, go to the local status page and enter the WAN IPs, subnet masks and gateways for the WAN link(s).
- If you do have DHCP that gives the MX internet access then you can configure the WAN IPs from the dashboard by clicking on the page below and then going to the Uplink tab
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2020 06:37 AM
Hello cmr,
i have some doubts in step 5.
For Use MX Uplinks IPs
What i understand is i just need to add the secondary WAN in Configure warm Spare>>>WAN 1 shared IP.
And for Use Virtual IPs i can add two more WAN IPs.
So actually what is difference??
Also how can i add heatbeat to check Primary MX is down??
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2020 07:32 AM
MX Uplink IPs
Existing MX WAN ports have an IP address from the provider that they are connected to.
New MX WAN ports need an additional IP address from each provider
Virtual uplink IPs
Existing MX WAN ports have an IP address from the provider that they are connected to.
New MX WAN ports need an additional IP address from each provider
A floating IP address that is held by the active MX is needed from each provider
The main benefits are:
Outbound traffic always comes from this address, not one MX and then sometimes the other.
Inbound traffic to internal services can go to the virtual IP, otherwise you need to know which MX is active by using dynamic DNS or similar with the inherent delays and complexity, it can work though.
Regarding the heartbeat, the dashboard will send emails about MX failover and WAN link failover
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2020 08:10 PM
I mean configuring part of heartbeat...
i can add a new vlan for warm spare heartbeat.And network 1.1.1.0/24 .Mx Ip 1.1.1.1 .
assign this vlan to port 4
connect port 4 between MX appliances.
thats all correct??
do I need to add any IP address 1.1.1.2 in mx warm spare box ?? Or does it automatically syn??
