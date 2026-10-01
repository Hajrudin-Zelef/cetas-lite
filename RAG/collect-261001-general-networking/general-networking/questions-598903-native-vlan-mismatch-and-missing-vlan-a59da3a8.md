---
id: collect-261001-general-networking/general-networking/questions-598903-native-vlan-mismatch-and-missing-vlan-a59da3a8
title: "questions-598903-native-vlan-mismatch-and-missing-vlan-a59da3a8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-598903-native-vlan-mismatch-and-missing-vlan-a59da3a8.md
source_anchor: ""
source_lines: [1, 97]
sha256: 5da0ef0aaf22f5df567b29521d7cf3d7a46a8cc62dd18dee1b84d5b7dfc04559
---

# questions-598903-native-vlan-mismatch-and-missing-vlan-a59da3a8

I'm trying to wrap my head around what exactly is going on here with a new site's configuration of their networking stack. This particular piece I am working is pretty simple but I having a hard time figuring out what the original intention was. There is a Cisco Catalyst 3750x with three Port Channels (each with four interfaces a piece) going to three ESXi hosts. The Catalyst is connected to the rest of the network via a Meraki MS42 via a single interface (no Port Channel). VLAN 100 carries the networking traffic, the other VLANs are dedicated to things like vMotion or isolated networks. I think a large part of my difficulty here is I don't speak Cisco-ese.
The Setup
Port-Channel 1
interface Port-channel1
 switchport access vlan 100
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
interface GigabitEthernet1/0/1
 description ESX1
 switchport access vlan 100
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 1 mode on
!
interface GigabitEthernet1/0/2
 description ESX1
 switchport access vlan 100
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 1 mode on
!
interface GigabitEthernet1/0/3
 description ESX1
 switchport access vlan 100
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 1 mode on
Port-Channel 2 (I'm leaving out Port-Channel 3 since it is identical in configuration to Port-Channel 2)
interface Port-channel2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
!
interface GigabitEthernet1/0/5
 description ESX2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 2 mode on
!
interface GigabitEthernet1/0/6
 description ESX2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 2 mode on
!
interface GigabitEthernet1/0/7
 description ESX2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 2 mode on
!
interface GigabitEthernet1/0/8
 description ESX2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 100,101,172,192
 switchport mode trunk
 switchport nonegotiate
 spanning-tree portfast trunk
 channel-group 2 mode on
Uplink Ports
On the Catalyst:
interface GigabitEthernet1/0/24
 description Uplink
 switchport access vlan 100
 switchport trunk native vlan 2
!
On the Meraki:
Trunk port using native VLAN 1; allowed VLANs: all
The Question/s
- The combination of switchport access andswitch port trunk allowed makes theswitchport access configuration a no-op, right? You cannot have a port in access mode and trunk mode unless I am mistaken. Can someone confirm this for me?
- It is my understanding that once you add a port to Port Channel all of the VLAN an STP configuration is done per Port Channel and not per port. If I create a Port Channel out of Fa 1/10 and Fa 1/11, I configure them as trunks using their assigned Port Channel and not their individual ports (at least this is what I do with ProCurves). Is this correct?
- If the last item is correct that means all of the per-port configuration of Port Channel members is either a no-op or was done prior to that port being made a Port Channel member. Is this a reasonable assumption?
- How the heck does the traffic from VLAN 100 get across the uplink (I can reach the VMs hosted on the ESXi hosts)? VLAN 100 disappears once it hits the Meraki and the native VLAN tags are different. Things are working but I can't help but feel something is weird with this setup and it would be preferable to push VLAN 100 all the way through to the rest of stack. To make things even stranger VLAN 2 terminates at Port 41 on the Meraki as well, everything else is set to Native VLAN 1.
Moving forward I am inclined to abandon VLAN 100 or reconfigure the rest of our stack so that the subnet that rides on VLAN 100 doesn't use multiple VLANs (100 and 1) and resolve the Native VLAN tag mismatch on the uplink (Port 41 -- Gi 1/0/24). Thoughts on this plan?
