---
id: collect-261001-meraki/meraki/questions-39932-configure-vlan-on-cisco-sg-200-switch-connected-to-a-meraki-swit-5e8faa44
title: "questions-39932-configure-vlan-on-cisco-sg-200-switch-connected-to-a-meraki-swit-5e8faa44"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-39932-configure-vlan-on-cisco-sg-200-switch-connected-to-a-meraki-swit-5e8faa44.md
source_anchor: ""
source_lines: [1, 13]
sha256: c3585d0ebef6ef19dc4a2aef7f06eca45cf7d43ae59fa8be8ff92aa828d93685
---

# questions-39932-configure-vlan-on-cisco-sg-200-switch-connected-to-a-meraki-swit-5e8faa44

I have a Cisco SG-200-26 switch connected to a port on a Meraki switch.
The Meraki switch shows a list of all devices detected on that switch port (the Cisco switch and all of the devices connected to the Cisco switch).
I have VLANs configured on a Meraki L3 switch upstream from the Meraki switch.
The Meraki port is configured to be a Trunk port using native VLAN 10; allowed VLANs: all
The Cisco switch is configured to have a static IP of 10.128.10.12, which is in VLAN 10, and a default gateway of 10.128.10.1.
How do I accomplish all of the following things:
- Configure the port on the Meraki switch to be Trunk port using native VLAN 1, and have the Cisco switch still work with a static IP in VLAN 10.
When I set the port on the Meraki switch to be Trunk port using native VLAN 1, the Cisco switch loses connectivity because it's static IP is set to be 10.128.10.12.
The only way I am able to get in to the Cisco web interface is to set the Meraki switch port back to Trunk port using native VLAN 10.
I would like to have the Meraki switch port to be Trunk port using native VLAN 1, because all of the other Meraki switches I have are configured that way.
- Configure a port on the Cisco switch to be an access port on VLAN 20.
Basically, what I need to be able to do is connect a device to a port on the Cisco switch, and have it be put into VLAN 20, so that in the Meraki interface, it shows both the switch (in VLAN 10), and the device (in VLAN 20).
I followed the instructions on this page to configure the port to be Tagged for VLAN 20, but it doesn't work - the device does not have network connectivity, and does not show up in the Meraki switch port client list.
