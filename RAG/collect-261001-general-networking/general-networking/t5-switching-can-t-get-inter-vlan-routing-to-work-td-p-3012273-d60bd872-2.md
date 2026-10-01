---
id: collect-261001-general-networking/general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872-2
title: "t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["pruning", "voice"]
source: docs/RAG/collect-261001-general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872.md
source_anchor: ""
source_lines: [15, 284]
sha256: d69984ba384de0c97f16b83926f8650aa15f77d8ff5d0e3926ade84910add01a
---

# t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 09:47 AM - edited 03-08-2019 09:36 AM
Hello,
I have a home lab that consists of 2 3550's and 1 1841 router. I'm learning about VLANs and trunking, and how to configure them. I've included a topology of the exact network that I've build to give you a better understanding of what i'm trying to do. The topology has the same interface numbers, ip address ranges and device names, that i used in the physical lab. But I cannot even get the hosts to ping their vlan interfaces, default gateways, or the other host. All the pings come back as "destination host unreachable" on the hosts, what am i missing? I've included all of the devices full configuration logs, and some show commands that i think are useful. If anyone wants a specific show command, or anything, leave a comment and I will update the post as soon as possible. Thanks!
SW1
SW1#show run
Building configuration...
Current configuration : 2364 bytes
!
version 12.1
no service pad
service timestamps debug uptime
service timestamps log uptime
no service password-encryption
!
hostname SW1
!
!
ip subnet-zero
!
!
spanning-tree mode pvst
spanning-tree extend system-id
!
!
!
!
interface FastEthernet0/1
 switchport access vlan 10
 switchport mode access
!
interface FastEthernet0/2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 1,10,20
 switchport mode trunk
!
interface FastEthernet0/3
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 1,10,20
 switchport mode trunk
!
interface FastEthernet0/4
 switchport mode dynamic desirable
!
interface FastEthernet0/5
 switchport mode dynamic desirable
!
interface FastEthernet0/6
 switchport mode dynamic desirable
!
interface FastEthernet0/7
 switchport mode dynamic desirable
!
interface FastEthernet0/8
 switchport mode dynamic desirable
!
interface FastEthernet0/9
 switchport mode dynamic desirable
!
interface FastEthernet0/10
 switchport mode dynamic desirable
!
interface FastEthernet0/11
 switchport mode dynamic desirable
!
interface FastEthernet0/12
 switchport mode dynamic desirable
!
interface FastEthernet0/13
 switchport mode dynamic desirable
!
interface FastEthernet0/14
 switchport mode dynamic desirable
!
interface FastEthernet0/15
 switchport mode dynamic desirable
!
interface FastEthernet0/16
 switchport mode dynamic desirable
!
interface FastEthernet0/17
 switchport mode dynamic desirable
!
interface FastEthernet0/18
 switchport mode dynamic desirable
!
interface FastEthernet0/19
 switchport mode dynamic desirable
!
interface FastEthernet0/20
 switchport mode dynamic desirable
!
interface FastEthernet0/21
 switchport mode dynamic desirable
!
interface FastEthernet0/22
 switchport mode dynamic desirable
!
interface FastEthernet0/23
 switchport mode dynamic desirable
!
interface FastEthernet0/24
 switchport mode dynamic desirable
!
interface GigabitEthernet0/1
 switchport mode dynamic desirable
!
interface GigabitEthernet0/2
 switchport mode dynamic desirable
!
interface Vlan1
 no ip address
 shutdown
!
interface Vlan10
 description VLAN10
 ip address 192.168.10.2 255.255.255.0
!
interface Vlan20
 ip address 192.168.20.2 255.255.255.0
!
ip classless
ip http server
!
!
line con 0
 logging synchronous
line vty 0 4
 logging synchronous
 login
line vty 5 15
 logging synchronous
 login
!
!
end
SW1#show vlan br
VLAN Name                          Status   Ports
---- -------------------------------- --------- -------------------------------
1      default                           active  Fa0/4, Fa0/5, Fa0/6, Fa0/7
                                                         Fa0/8, Fa0/9, Fa0/10, Fa0/11
                                                         Fa0/12, Fa0/13, Fa0/14, Fa0/15
                                                         Fa0/16, Fa0/17, Fa0/18, Fa0/19
                                                         Fa0/20, Fa0/21, Fa0/22, Fa0/23
                                                         Fa0/24, Gi0/1, Gi0/2
10  VLAN10                          active  Fa0/1
20  VLAN20                          active
1002 fddi-default                  act/unsup
1003 token-ring-default        act/unsup
1004 fddinet-default             act/unsup
1005 trnet-default               act/unsup
SW1#show int tru
Port    Mode               Encapsulation      Status     Native vlan
Fa0/2   on                      802.1q             trunking          1
Fa0/3   on                      802.1q             trunking          1
Port Vlans allowed on trunk
Fa0/2 1,10,20
Fa0/3 1,10,20
Port Vlans allowed and active in management domain
Fa0/2 1,10,20
Fa0/3 1,10,20
Port Vlans in spanning tree forwarding state and not pruned
Fa0/2 1,10,20
Fa0/3 1,10,20
SW1#show int f0/1 switch
Name: Fa0/1
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: negotiate
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
Access Mode VLAN: 10 (VLAN10)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: ALL
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none
SW1#show int f0/2 switch
Name: Fa0/2
Switchport: Enabled
Administrative Mode: trunk
Operational Mode: trunk
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: dot1q
Negotiation of Trunking: On
Access Mode VLAN: 1 (default)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: 1,10,20
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none
SW1#show int f0/3 switch
Name: Fa0/3
Switchport: Enabled
Administrative Mode: trunk
Operational Mode: trunk
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: dot1q
Negotiation of Trunking: On
Access Mode VLAN: 1 (default)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Administrative private-vlan trunk native VLAN: none
Administrative private-vlan trunk encapsulation: dot1q
Administrative private-vlan trunk normal VLANs: none
Administrative private-vlan trunk private VLANs: none
Operational private-vlan: none
Trunking VLANs Enabled: 1,10,20
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Appliance trust: none
SW2
SW2#show run
Building configuration...
Current configuration : 1783 bytes
!
version 12.1
no service pad
service timestamps debug uptime
service timestamps log uptime
no service password-encryption
!
hostname SW2
!
!
ip subnet-zero
!
!
spanning-tree mode pvst
spanning-tree extend system-id
!
!
interface FastEthernet0/1
