---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-11240-why-cant-i-nat-to-a-local-network-address-on-my-ubnt-edgerouter-5927676f
title: "questions-11240-why-cant-i-nat-to-a-local-network-address-on-my-ubnt-edgerouter-5927676f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-11240-why-cant-i-nat-to-a-local-network-address-on-my-ubnt-edgerouter-5927676f.md
source_anchor: ""
source_lines: [1, 8]
sha256: 3e747ff9b1d1c8da0d9af9a3d0cb5557abf86001c2c150009b5d4ce45508594b
---

# questions-11240-why-cant-i-nat-to-a-local-network-address-on-my-ubnt-edgerouter-5927676f

I am trying to set up a very simple network.
I have an Ubiquiti Edgerouter on 192.168.1.1 with eth0 as WAN port. I also have an Ubiquiti Rocket m5 on 192.168.1.2, both the router and the Rocket have a web interface. I have set the https-port from the router on port 8443. The Rocket has a web interface on port 80 & 443.
What i did:
I have made a working connection with my router, both ssh and web interface on 8443.
Add a firewall rule for both port 80 and 443.
This seems to be working. I also made a NAT rule to forward to the Rocket on 192.168.1.2 1 for port 80 and 1 for port 443. Like so:
As you can see, the "count" is still 0. So i'm assuming this is the problem.
Could anyone help me here? I'm looking for a very simple way to forward to 192.168.1.2. I am quite sure the web interface is enabled, because, if i "join" the network by plugging a cable directly into the router and into my PC, i can access the web interface.
