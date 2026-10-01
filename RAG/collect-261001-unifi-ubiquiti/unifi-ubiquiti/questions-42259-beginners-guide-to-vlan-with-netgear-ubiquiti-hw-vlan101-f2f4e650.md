---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-42259-beginners-guide-to-vlan-with-netgear-ubiquiti-hw-vlan101-f2f4e650
title: "questions-42259-beginners-guide-to-vlan-with-netgear-ubiquiti-hw-vlan101-f2f4e650"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-42259-beginners-guide-to-vlan-with-netgear-ubiquiti-hw-vlan101-f2f4e650.md
source_anchor: ""
source_lines: [1, 15]
sha256: 0cf061ee0a6cba19e527a936534cedd04f634ebd65366c6f8838ad808088ba6f
---

# questions-42259-beginners-guide-to-vlan-with-netgear-ubiquiti-hw-vlan101-f2f4e650

I've got a few questions related to networking, more specifically VLAN'ing.
I've got a Gateway/Router - Ubiquiti USG
I've got a 48Port Netgear Switch - Netgear GS748Tv5 L2+/L3 Lite
I've got 5 Ubiquiti UniFi Wireless Access Points connected to the Netgear switch
Can someone give me some basic documentation on how I should go about setting up multiple VLAN's on both devices.
I have created the 4 Networks that I wish to operate. All have DHCP on for their respective IP ranges. Network A - Private | 192.168.15.x | VLAN ID 1
Network B - Guest | 192.168..x | VLAN ID 50
Network C - AV | 192.168..x | VLAN ID 20
Network D - POS | 192.168..x | VLAN ID 10
I need Network A to be able to access the WAN & Network D I need Network B to be able to access the WAN Only I need network C to be able to access the WAN and potentially a device on Network A I need network D to be able to access Network A and one device to access the WAN.
Each Network is going to be accessible via Wifi, and the Wifi Points are Ubiquiti devices, so they seem to integrate quite easily into the Ubiquiti controller and have had the VLAN ID's associated with the respective Wifi networks.
For wired devices I know I can set certain ports on the switch to automatically assign those devices to a certain VLAN so that will solve the issue of handling wired devices in a VLAN scenario.
Currently I'm not getting any DHCP out of the router when I connect to a Wifi network other than the default one.
I know I must be missing a step somewhere but I'm not confident on what I've missed. I have created the VLAN's in the switch but I'm not sure if the switch is letting the DHCP traffic go from the USG port to the Wifi devices?
Any help with this is greatly appreciated.
