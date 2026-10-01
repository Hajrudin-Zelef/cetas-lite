---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1787519-problems-after-swapping-ubiquiti-edgeroutuer-for-usg-897b7c3e
title: "questions-1787519-problems-after-swapping-ubiquiti-edgeroutuer-for-usg-897b7c3e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1787519-problems-after-swapping-ubiquiti-edgeroutuer-for-usg-897b7c3e.md
source_anchor: ""
source_lines: [1, 14]
sha256: 32f4e4ba3ad2ea412dd6894396c2535485295c4d93118b19098a2fd6969e1ffe
---

# questions-1787519-problems-after-swapping-ubiquiti-edgeroutuer-for-usg-897b7c3e

I know this is less that concise, but I am struggling a bit with where to even start.
My EdgeRouter X died, and I got a UniFi USG (3p) to replace. Most things work fine, but I have thus far noticed a few weird things (that were not a problem with my old EdgeRouter):
- My TrueNAS (10.0.30.27 ) can't ping google from shell (is accessible at FQDN via NGINX).
- My Nextcloud (10.0.30.67 ) running in a jail on the TrueNAS can't access the Nextcloud servers to check for updates (also accessible at FQDN via NGINX).
- An IoT device (10.0.20.36 ) (a hardwired hub for integrating some pet accessories) can't connect to the manufacturer's servers. If I connect a computer to the same port (on the same VLAN), it has internet access. One possible hint: The manufacturer for the device said it needs access to outbound ports443 and8883 .
I haven't changed any settings in my PiHole/DNSMasq, and I have created the same VLAN topology and assigned the same DHCP reservations for devices, so I'm assuming that it's something in the UniFi controller related to the USG (the only new device on network). But other than that, I have no clue where to even start with this (hence the vague title).
I am not finished setting up, so far I have only done the following:
- Set up my VLANs (Default [only for UniFi devices]: 10.0.0.0/24 , VLAN10:10.0.10.0/23 , VLAN20:10.0.20.0/23 , VLAN30:10.0.30.0/24 , and VLAN40:10.0.40.0/24 , VPN:10.10.10.0/24 )
- Set DHCP reservations for devices
- Added mDNS for Default, VLAN10, 20, 30, and VPN
- Set DNS to 10.0.30.43 (my PiHole/DNSMasq, added Cloudflare as backup)
- Added port forwards for 80 ,443 (for NGINX),94 (for VPN), and one more for remote ssh access to a server
- Checked Network Discovery andInform Host (setting the IP to that of UniFi Controller;10.0.30.61 )
I will off course add a bunch of firewall rules like I had on the EdgeRouter, but I wanted to get the basics to work first. And it seems something is acting up. Any ideas what would be appreciated.
