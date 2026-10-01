---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-5570-trying-to-configure-hp-procurve-vlans-to-segment-unifi-guest-traf-35365449
title: "questions-5570-trying-to-configure-hp-procurve-vlans-to-segment-unifi-guest-traf-35365449"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-5570-trying-to-configure-hp-procurve-vlans-to-segment-unifi-guest-traf-35365449.md
source_anchor: ""
source_lines: [1, 6]
sha256: 217fa54f8d0bb1dabc7709c5a357477eab396c9dd513ff7be677d51503750e45
---

# questions-5570-trying-to-configure-hp-procurve-vlans-to-segment-unifi-guest-traf-35365449

My goal is to put all guest network traffic from the Ubiquiti Unifi access points onto their own subnet using VLAN 20 on a series of HP Procurve 2610-48-PWR switches.
The Unifi Access controller allows me to set the guest network to work with a given VLAN ID, which I obviously set to 20.
I want all guest traffic to be filtered to our secondary ISP on a separate firewall that will also issue DHCP addressing. The ISP runs to a firewall that runs to a web filter, then to the central Procurve. The Procurve then splits to the other two Procurves in the other wings of the campus.
On the Procurves, I configured each to use VLAN 20 over the designated subnet and tagged every port that has an access point plugged into it. All three switches have IP addresses that are contained within the desired subnet.
I was under the impression that configuring a VLAN with the same ID on each connected switch, as well as supplying an IP address within the correct addressing scheme would allow them to communicate. The switches can communicate with each other over the separate subnet, but they can't reach the filter or the firewall. They also can't reach a client plugged into the switch with a statically assigned IP address in the correct subnet.
I have been at this for a few hours and am running out of ideas. Any advice would be much appreciated.
