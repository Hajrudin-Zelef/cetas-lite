---
id: collect-261001-general-networking/general-networking/questions-1040169-how-to-configure-the-att-arris-bgw-210-router-for-ip-passthrou-f7189db4
title: "questions-1040169-how-to-configure-the-att-arris-bgw-210-router-for-ip-passthrou-f7189db4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1040169-how-to-configure-the-att-arris-bgw-210-router-for-ip-passthrou-f7189db4.md
source_anchor: ""
source_lines: [1, 20]
sha256: 40a9a94e26c5da5818a038d5bed3a7b3ea1e7fba4000dd8c18ac184826b61a72
---

# questions-1040169-how-to-configure-the-att-arris-bgw-210-router-for-ip-passthrou-f7189db4

We are setting up AT&T fiber internet with 5 usable static IPs and the Ubiquity UniFi Dream Machine Pro (UDM-Pro). I would like to configure the BGW-210 to act as a bridge to the UDM-Pro.
I found this article on how to configure the BGW-210 in IP Passthrough mode (similar to bridge), but some of the details are a bit unclear and I need to adjust this setup process to use one or more of my static IP addresses on the UDM-Pro.
In one paragraph, the article said DHCP is not needed for Passthrough mode:
The DHCP Server option can be turned off if you're doing IP Passthrough, but you must leave it on if you are doing Default Server...
But later on it said that you are still using DHCP:
It is worth mentioning that this is still a DHCP address that your internal device is getting...
Which leaves some confusion on whether or not DHCP server should be configured or disabled.
Here are the things I'm fairly certain of:
- Set the "Public LAN Subnet" different than the UDM-Pro LAN subnet.
- Setup the IP addresses provided by AT&T under the "Public Subnet" section. I did this and we can connect to the Internet.
- I need to enable "Allocation Mode" to Passthrough .
- I need to set the "Passthrough Mode" to DHCPS-fixed .
- I need to enter the MAC address of the UDM-Pro in "Passthrough Fixed Mac Address".
- I need to setup the UDM-Pro to get its WAN address from a DHCP server.
What I'm unclear about is:
- Under "Public Subnet" section, do I leave "Public Subnet Mode" On and "Allow Inbound Traffic"Off ?
- Do I leave "DHCP Server Enable" On and what IP address ranges should be there? The author of the post seems to mix the Default Server instructions with the Passthrough instructions.
- After putting the BGW-210 in Passthrough mode, do I still need to turn off packet filtering and firewall features or does Passthrough mode bypass these automatically?
Again, the goal is to "bridge" the AT&T router and have the UDM-Pro manage all routing and security.
Thank you.
