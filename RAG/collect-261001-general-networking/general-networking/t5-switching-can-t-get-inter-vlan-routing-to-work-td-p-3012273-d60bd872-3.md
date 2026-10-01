---
id: collect-261001-general-networking/general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872-3
title: "t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["full-duplex", "pruning", "voice"]
source: docs/RAG/collect-261001-general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872.md
source_anchor: ""
source_lines: [285, 592]
sha256: 521659c9bfc1096e1826e3055fe5db392394aeed51ea444d32fa7f5efd3990c0
---

# t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872

 switchport access vlan 20
 switchport mode access
 no ip address
!
interface FastEthernet0/2
 switchport trunk encapsulation dot1q
 switchport trunk allowed vlan 1,10,20
 switchport mode trunk
 no ip address
!
interface FastEthernet0/3
 no ip address
!
interface FastEthernet0/4
 no ip address
!
interface FastEthernet0/5
 no ip address
!
interface FastEthernet0/6
 no ip address
!
interface FastEthernet0/7
 no ip address
!
interface FastEthernet0/8
 no ip address
!
interface FastEthernet0/9
 no ip address
!
interface FastEthernet0/10
 no ip address
!
interface FastEthernet0/11
 no ip address
!
interface FastEthernet0/12
 no ip address
!
interface FastEthernet0/13
 no ip address
!
interface FastEthernet0/14
 no ip address
!
interface FastEthernet0/15
 no ip address
!
interface FastEthernet0/16
 no ip address
!
interface FastEthernet0/17
 no ip address
!
interface FastEthernet0/18
 no ip address
!
interface FastEthernet0/19
 no ip address
!
interface FastEthernet0/20
 no ip address
!
interface FastEthernet0/21
 no ip address
!
interface FastEthernet0/22
 no ip address
!
interface FastEthernet0/23
 no ip address
!
interface FastEthernet0/24
 no ip address
!
interface GigabitEthernet0/1
 no ip address
!
interface GigabitEthernet0/2
 no ip address
!
interface Vlan1
 no ip address
 shutdown
!
interface Vlan10
 ip address 192.168.10.3 255.255.255.0
!
interface Vlan20
 ip address 192.168.20.3 255.255.255.0
!
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
end
SW2#show vlan br
VLAN Name                          Status  Ports
---- -------------------------------- --------- -------------------------------
1    defaultn                           active  Fa0/3, Fa0/4, Fa0/5, Fa0/6
                                                         Fa0/7, Fa0/8, Fa0/9, Fa0/10
                                                         Fa0/11, Fa0/12, Fa0/13, Fa0/14
                                                         Fa0/15, Fa0/16, Fa0/17, Fa0/18
                                                         Fa0/19, Fa0/20, Fa0/21, Fa0/22
                                                         Fa0/23, Fa0/24, Gi0/1, Gi0/2
10 VLAN10                            active
20 VLAN20                            active  Fa0/1
1002 fddi-default                    active
1003 token-ring-default         active
1004 fddinet-default               active
1005 trnet-default                  active
SW2#show int tru
Port    Mode      Encapsulation   Status     Native vlan
Fa0/2   on             802.1q           trunking          1
Port Vlans allowed on trunk
Fa0/2 1,10,20
Port Vlans allowed and active in management domain
Fa0/2 1,10,20
Port Vlans in spanning tree forwarding state and not pruned
Fa0/2 1,10,20
SW2#show int f0/1 switch
Name: Fa0/1
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: negotiate
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
Access Mode VLAN: 20 (VLAN20)
Trunking Native Mode VLAN: 1 (default)
Voice VLAN: none
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none
Operational private-vlan: none
Trunking VLANs Enabled: ALL
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Voice VLAN: none (Inactive)
Appliance trust: none
SW2#show int f0/2 switch
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
Operational private-vlan: none
Trunking VLANs Enabled: 1,10,20
Pruning VLANs Enabled: 2-1001
Capture Mode Disabled
Capture VLANs Allowed: ALL
Protected: false
Unknown unicast blocked: disabled
Unknown multicast blocked: disabled
Voice VLAN: none (Inactive)
Appliance trust: none
R1
R1#show run
Building configuration...
Current configuration : 952 bytes
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
R1#show int f0/0
FastEthernet0/0 is up, line protocol is up
 Hardware is Gt96k FE, address is 001f.ca8c.37ec (bia 001f.ca8c.37ec)
 MTU 1500 bytes, BW 100000 Kbit, DLY 100 usec,
 reliability 255/255, txload 1/255, rxload 1/255
 Encapsulation 802.1Q Virtual LAN, Vlan ID 1., loopback not set
 Keepalive set (10 sec)
 Full-duplex, 100Mb/s, 100BaseTX/FX
 ARP type: ARPA, ARP Timeout 04:00:00
 Last input 00:00:02, output 00:00:01, output hang never
 Last clearing of "show interface" counters never
 Input queue: 0/75/0/0 (size/max/drops/flushes); Total output drops: 0
 Queueing strategy: fifo
 Output queue: 0/40 (size/max)
 5 minute input rate 0 bits/sec, 0 packets/sec
 5 minute output rate 0 bits/sec, 0 packets/sec
 1063 packets input, 188027 bytes
 Received 826 broadcasts, 0 runts, 0 giants, 0 throttles
 0 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
 0 watchdog
 0 input packets with dribble condition detected
 991 packets output, 111728 bytes, 0 underruns
 0 output errors, 0 collisions, 7 interface resets
 0 babbles, 0 late collision, 0 deferred
 0 lost carrier, 0 no carrier
R1#show int f0/0.10
FastEthernet0/0.10 is up, line protocol is up
 Hardware is Gt96k FE, address is 001f.ca8c.37ec (bia 001f.ca8c.37ec)
 Internet address is 192.168.10.1/24
 MTU 1500 bytes, BW 100000 Kbit, DLY 100 usec,
 reliability 255/255, txload 1/255, rxload 1/255
 Encapsulation 802.1Q Virtual LAN, Vlan ID 10.
 ARP type: ARPA, ARP Timeout 04:00:00
 Last clearing of "show interface" counters never
R1#show int f0/0.20
FastEthernet0/0.20 is up, line protocol is up
 Hardware is Gt96k FE, address is 001f.ca8c.37ec (bia 001f.ca8c.37ec)
 Internet address is 192.168.20.1/24
 MTU 1500 bytes, BW 100000 Kbit, DLY 100 usec,
 reliability 255/255, txload 1/255, rxload 1/255
 Encapsulation 802.1Q Virtual LAN, Vlan ID 20.
 ARP type: ARPA, ARP Timeout 04:00:00
 Last clearing of "show interface" counters never
- Labels:
- 
						
							
		
