---
id: collect-261001-general-networking/general-networking/questions-1144523-inter-vlan-connection-issues-when-devices-use-wi-fi-and-opnsen-f388a250
title: "Inter-VLAN connection issues when devices use Wi-Fi and OPNsense router"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1144523-inter-vlan-connection-issues-when-devices-use-wi-fi-and-opnsen-f388a250.md
source_anchor: ""
source_lines: [1, 31]
sha256: e08fc1738203c9ed60d482a400b9ad465725a381f25d7413cb33c36425a7604e
---

# Inter-VLAN connection issues when devices use Wi-Fi and OPNsense router

*Score : 1 | Source : https://serverfault.com/questions/1144523/inter-vlan-connection-issues-when-devices-use-wi-fi-and-opnsense-router*

I am trying to segregate devices in my home network with 2 different VLANs: HOME and IOT. I have the following network devices:
This is how the wired connections are laid out:
Here's a diagram I created to help visualize the network layout
In OPNsense, I have created the following VLANs:
DHCP is enabled in the LAN and both VLANS, as follows:
LAN:
HOME:
IOT:
In OPNsense, I have created firewall rules to allow:
In the switch, I have configured VLANs and ports as follows using Advanced 802.1Q VLAN:
In the PVID table of the switch, all ports have ID 1, except for port 5, which has ID 3.
In the Access Point, I have created 3 SSIDs:
SSID1:
SSID2:
SSID3:
To the SSID2, I have a Windows laptop connected. To the SSID3, I have a Wi-Fi IP camera connected.
All of this seem to work fine for the most part. All devices get assigned to their respective VLAN, with correct DHCP assignments and they all can access the internet. Also, any inter VLAN communication between wired and wi-fi devices (where firewall rules allow) work correctly. For example, I am able to connect to the IOT Wired IP camera from the HOME Wi-fi laptop and likewise I am able to connect to the IOT Wi-Fi camera from the LAN Wired PC desktop. Inter VLAN wired to wired communication also works fine (again, where firewall rules allow).
The issue only arises when I attempt a connection between Wi-Fi devices in different VLANS. If I try to access the Wi-fi IOT camera from the Wi-Fi HOME laptop, the connection cannot be established.
In an attempt to troubleshoot the issue, I connected a linux laptop running an Nginx webserver to the Wi-Fi SSID3 (IOT VLAN). In this laptop, I run tcpstat to show me incoming and outgoing connections. When I try to access the home page hosted in the linux laptop from the SSID2 (HOME VLAN) Windows laptop, the page never loads. tcpstat shows the incoming connection from the Windows laptop, however it stays stuck in SYN_RECV and never reaches ESTABLISHED. Accessing the page from the Wired Windows PC works just fine.
At this point I am at a loss why there is this seemingly inter-VLAN routing issue when both devices are Wi-Fi, but they work correctly when at least one device is wired. Any tips or insights here are highly appreciated.
PS: both AP and the switch are running the latest firmware versions from Netgear.

---

### Reponse (acceptee) — score 1

So turns out the issue is with the WAX630E AP. I contacted Netgear support and they acknowledge this AP doesn't support inter-VLAN communications when using a single AP. They don't plan to release a fix either, which kind of baffles me, but what can I do. So this one is being returned.
