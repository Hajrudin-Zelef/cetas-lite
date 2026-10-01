---
id: collect-261001-meraki/meraki/questions-63786-meraki-bridge-mode-dhcp-is-untagged-8888cfc7
title: "questions-63786-meraki-bridge-mode-dhcp-is-untagged-8888cfc7"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-63786-meraki-bridge-mode-dhcp-is-untagged-8888cfc7.md
source_anchor: ""
source_lines: [1, 40]
sha256: 8496b045b6cca7e8e602c3e49624d0f3e7923a8381ca27e77acdc2d04b7c7355
---

# questions-63786-meraki-bridge-mode-dhcp-is-untagged-8888cfc7

We have the following set-up:
+------+     +--------+     +--------+
|Router|-----| switch |-----| Meraki |
+------+     +--------+     +--------+
                |
             +---------+
             |wireshark|
             +---------+
The Meraki serves a number of SSIDs, A, B and C. The Meraki is in bridge mode. For all the SSIDs, there is a separate VLAN, 511 for A, 512 for B and 513 for C. On the router (a cisco device), on all the VLANs, there is a DHCP server. In the Meraki console, the option "use VLAN tags" os on for all SSIDsThe switch ports are in trunk mode, with 510 as native VLAN.
The main purpose of the switch was to allow Wireshark on a span port.
C is an open accesspoint. When someone connects to C, he will do a DHCP. It was my understanding that, being in bridge-mode, the DHCP request should be on VLAN 513, and therefore tagged with VLAN 513. However, the router with the DHCP server sees DHCP requests on VLAN 510 (the native VLAN). With Wireshark, we see that the DHCP request is untagged. It is therefore logical that the router sees it on VLAN 510.
For Meraki: On the page "Addressing and traffic",
- we checked "Bridge mode: Make clients part of the LAN"
- after "VLAN tagging", we selected "Use VLAN tagging"
- In the table VLAN ID, we set A to VLAN 511, B to VLAN 512 and C to VLAN 513
- in that table, the column "VLAN ID" in the row "All other APs: is empty (no value entered here)
On the switch, the relevant part is:
interface GigabitEthernet2
description LAN Meraki AP
switchport trunk native vlan 510
switchport trunk allowed vlan 1,510-514,1002-1005
switchport mode trunk
no ip address
On the Router, where the DHCP server lives, we have:
ip dhcp pool meraki-base-510
network 10.10.1.0 255.255.255.0
default-router 10.10.1.1
dns-server 8.8.8.8 8.8.4.4
ip dhcp pool meraki-open-513
network 10.10.3.0 255.255.255.0
default-router 10.10.3.1
dns-server 8.8.8.8 8.8.4.4
If we ask the switch where the MAC-id of the phone lives, it is on VLAN 513. So, it will never get an IP address.
When we give the phone a fixed IP address on VLAN 513, it connects as it should.
The router DHCP server produces the following info:
DHCPD: No option 25
DHCPD: DHCPDISCOVER received from client 1111.2222.3333.44 on interface Vlan 510
DHCPD: Allocate an address without class information (10.10.3.0)
(sorry, no traces from wireshark were saved; DHCP Discover was without VLAN tag)
Why does the Meraki put the DHCP request on the native VLAN? And what should I do to get a DHCP request tagged with VLAN 513?
