---
id: collect-261001-meraki/meraki/questions-22414-sonicwall-tz-190-enhance-w-meraki-mr12-configure-opt-port-for-gu-9014747f
title: "questions-22414-sonicwall-tz-190-enhance-w-meraki-mr12-configure-opt-port-for-gu-9014747f"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-22414-sonicwall-tz-190-enhance-w-meraki-mr12-configure-opt-port-for-gu-9014747f.md
source_anchor: ""
source_lines: [1, 7]
sha256: 07af2f6c19d36023b34a94d3e30fc9dabd42d579939044772d19c1563db50cd4
---

# questions-22414-sonicwall-tz-190-enhance-w-meraki-mr12-configure-opt-port-for-gu-9014747f

You likely don't have a DHCP scope configured for the OPT interface (in the DMZ Zone). I don't have a TZ 190 or a device with that older SonicOS 4.x firmware but you can do one of two things:
- Configure a DHCP scope so the Meraki AP can obtain an IP and access the internet: go to Network, DHCP Server and see if you can add a Dynamic DHCP scope for the OPT interface.
- Configure a Static IP on the Meraki AP so it can access the internet. You may need to connect it to your LAN port first, allow it to obtain (via DHCP) an IP address, then configure the Static IP (on the OPT subnet) in the Meraki cloud config (or, I think depending on the firmware, log into that AP directly using the IP it has on the LAN and assign it an IP in the OPT subnet)
Meraki APs need IPs with internet access as they connect to the cloud-based controller network. Without a Static IP or DHCP on the OPT subnet it has no IP.
Whether you configure the NAT on the Meraki (for the Guest Network) won't change, you can either leave the Meraki clients on the OPT network and manage them with the SonicWALL, OR configure the Meraki Guest network (which enabled a mini-router on the Meraki AP which does NAT and creates a 10.x.x.x Guest network) and manage the guests on the Meraki (this would be Double NAT). 
If you manage the Guests on the Meraki you might as well leave the Meraki on the LAN subnet since it will keep them isolated from the local management network. But if you wanted to have, say, your internal Wireless users on the OPT/DMZ subnet BUT ALSO have Guests on a separate subnet then putting the Meraki on OPT and configure Meraki controlled Guests is the right idea.
Or you can configure VLAN subinterfaces if you have a managed switch and run several subnets that way, all managed from your SonicWALL.
