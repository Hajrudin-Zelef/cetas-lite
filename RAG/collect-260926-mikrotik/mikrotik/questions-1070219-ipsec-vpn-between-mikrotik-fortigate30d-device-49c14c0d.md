---
id: collect-260926-mikrotik/mikrotik/questions-1070219-ipsec-vpn-between-mikrotik-fortigate30d-device-49c14c0d
title: "questions-1070219-ipsec-vpn-between-mikrotik-fortigate30d-device-49c14c0d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-1070219-ipsec-vpn-between-mikrotik-fortigate30d-device-49c14c0d.md
source_anchor: ""
source_lines: [1, 5]
sha256: 09d1fb888034b304116c5cf7a0c31f785bdfc2e496410513aa6f5aacabe4cfe6
---

# questions-1070219-ipsec-vpn-between-mikrotik-fortigate30d-device-49c14c0d

I want to setup an IPSEC VPN tunnel between a mikrotik 912UAG device and a fortigate30D device so that anything connected on my mikrotik LAN is able to communicate to anything connected on my fortigate LAN. My mikrotik device has an ip address of 172.16.99.1(LAN) and my fortigate has an ip address of 10.0.0.1(LAN)
A dialup connection is being used between the two devices where the mikrotik is utilising a 3g connection and the fortigate is connected to a local ADSL line through a modem.
The modem(192.168.1.1) is connected to the WAN interface of the fortigate30D device and assigning it an ip address of 192.168.1.10.
I have setup the mikrotik device to use the external IP address of my modem and have also setup port forwarding on my modem to forward UDP 500/UDP 4500 to my fortigate30D device. However, it appears that none of the VPN packets are getting to my fortigate30D device.
Can anyone provide any insight/help please?
