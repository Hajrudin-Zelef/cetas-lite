---
id: collect-261001-general-networking/general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872-5
title: "t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872.md
source_anchor: ""
source_lines: [801, 1016]
sha256: dfdeb2f36438cb75376dda2162a7f0be1e1ec080851ac7e5f86c794785f575ca
---

# t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872

1002 fddi-default act/unsup
1003 token-ring-default act/unsup
1004 fddinet-default act/unsup
1005 trnet-default act/unsup
SW2
SW2#show run
Building configuration...
Current configuration : 1760 bytes
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
switchport access vlan 20
switchport mode access
no ip address
!
interface FastEthernet0/2
switchport trunk encapsulation dot1q
switchport trunk allowed vlan 1,10,20,99
switchport mode trunk
no ip address
!
-----Omitted lines-----
!
interface Vlan1
no ip address
shutdown
!
interface Vlan99
ip address 192.168.99.2 255.255.255.0
!
ip default-gateway 192.168.99.1
ip classless
ip http server
!
!
!
line con 0
logging synchronous
line vty 0 4
login
line vty 5 15
login
!
End
SW2#show vlan br
VLAN Name Status Ports
---- -------------------------------- --------- -------------------------------
1 default active Fa0/3, Fa0/4, Fa0/5, Fa0/6
Fa0/7, Fa0/8, Fa0/9, Fa0/10
Fa0/11, Fa0/12, Fa0/13, Fa0/14
Fa0/15, Fa0/16, Fa0/17, Fa0/18
Fa0/19, Fa0/20, Fa0/21, Fa0/22
Fa0/23, Fa0/24, Gi0/1, Gi0/2
20 VLAN20 active Fa0/1
99 MANAGEMENT active
1002 fddi-default active
1003 token-ring-default active
1004 fddinet-default active
1005 trnet-default active
R1
R1#show run
Building configuration...
Current configuration : 1072 bytes
!
version 12.4
service timestamps debug datetime msec
service timestamps log datetime msec
no service password-encryption
!
hostname R1
!
boot-start-marker
boot-end-marker
!
!
no aaa new-model
!
resource policy
!
ip cef
!
!
!
!
!
!
!
!
!
!
!
!
interface FastEthernet0/0
no ip address
duplex auto
speed auto
!
interface FastEthernet0/0.10
encapsulation dot1Q 10
ip address 192.168.10.1 255.255.255.0
no snmp trap link-status
!
interface FastEthernet0/0.20
encapsulation dot1Q 20
ip address 192.168.20.1 255.255.255.0
no snmp trap link-status
!
interface FastEthernet0/0.99
encapsulation dot1Q 99
ip address 192.168.99.1 255.255.255.0
no snmp trap link-status
!
interface FastEthernet0/1
no ip address
shutdown
duplex auto
speed auto
!
interface Serial0/0/0
no ip address
shutdown
no fair-queue
clock rate 2000000
!
!
!
ip http server
no ip http secure-server
!
!
!
!
!
control-plane
!
!
!
line con 0
logging synchronous
line aux 0
line vty 0 4
logging synchronous
login
line vty 5 15
logging synchronous
login
!
scheduler allocate 20000 1000
end
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 01:33 PM
From the PC in vlan 10 can you ping 192.168.20.1 ?
If you can then from the same PC can you ping both 192.168.99.2 and .3.
If you can do that try pinging either PC from the router and let us know.
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 02:35 PM
From the PC in VLAN10 i can only ping 192.168.10.1. And from the pc in VLAN20 i cant ping anywhere.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 03:00 PM
While logged into the router can you ping 192.168.99.1, 192.168.99.2 and 192.168.99.3? Also could you post the show vlan-switch command from the router? Finally could you also post the show ip route command from the router?
Regards,
Sam
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 03:09 PM
Yes, I remember I could ping to 99.2 & 99.3 from 99.1. But I cant access my lab at the moment. I'll reply with the show commands asap.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 03:17 PM
When you do get access can you do a ping from sw2 and ping both 192.168.10.1 and 192.168.20.1 which would test the routing between vlans.
Just trying to narrow down the problem.
Jon
