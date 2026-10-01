---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-670836-disable-edgerouter-lite-gui-on-dmz-interface-5345acd4
title: "questions-670836-disable-edgerouter-lite-gui-on-dmz-interface-5345acd4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-670836-disable-edgerouter-lite-gui-on-dmz-interface-5345acd4.md
source_anchor: ""
source_lines: [1, 10]
sha256: 257f99c8047a0d1c289e6484734c6634b598584c20cfb5456f809c838294890c
---

# questions-670836-disable-edgerouter-lite-gui-on-dmz-interface-5345acd4

I'm using the Ubiquiti EdgeRouter in a SOHO network. I've configured it using the "WAN+2LAN2" wizard, without bridging the LAN ports. This sets up the three port router with eth0 assigned to the WAN, and eth1 and eth2 serving two separate subnets, 10.0.0.1/24 and 10.0.1.1/24 respectively. I plan to add a switch to eth1 to serve my LAN. I plan to connect a web-server to eth2, and port-forward port 80 to creating a DMZ (perimeter network). I would like to disable access to the router's web GUI on eth2. This will ensure that, if my DMZ is compromised, the router cannot be bridged to allow access to the remainder of my LAN. Unfortunately, there is no obvious way to set this up in the web GUI. How can I disable GUI access for one of the LAN interfaces?
1 Answer 1
I've been able to do this by changing the ip address where the gui listens. I'm sure it might be good to additionally set up some firewall rules for this. Here is what worked for me on the CLI.
configure
set service gui listen-address 10.0.0.1
set service ssh listen-address 10.0.0.1
commit
save
exit
By default, the IP address for eth1 is 192.168.1.1, so you may have to change the above to point to the ip address of your LAN's gateway (interface's IP address).
