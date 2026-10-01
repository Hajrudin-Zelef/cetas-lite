---
id: collect-261001-general-networking/general-networking/questions-57335-cannot-ping-firewall-after-changing-sfp-48acdf7f
title: "questions-57335-cannot-ping-firewall-after-changing-sfp-48acdf7f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-57335-cannot-ping-firewall-after-changing-sfp-48acdf7f.md
source_anchor: ""
source_lines: [1, 248]
sha256: 555839f38d30b106af36eabb949d0f608063ebe27dbad9c58adf6e7bd5502fb7
---

# questions-57335-cannot-ping-firewall-after-changing-sfp-48acdf7f

I am currently working on an issue that has me stumped... I have a network 172.16.144.0/20 which is connected to our fortigate 300D firewall via an Cisco 3850 external switch. Due to hardware limitations the port being used has a 100mbps SFP the rest of the ports are using 1gbps.
The issue is that I am able to ping to all my devices from the switch but yet I am unable to ping the firewall nor am I able to ping the switch from the firewall. Using the same IP and firewall port as the switch, I used a laptop and was able to ping the firewall.
I'm assuming the SFP is the culprit, not sure how though...
Current configuration : 17202 bytes
!
! Last configuration change at 20:16:42 UTC Fri Mar 1 2019
!
version 15.2
no service pad
service tcp-keepalives-in
service tcp-keepalives-out
service timestamps debug datetime msec localtime show-timezone
service timestamps log datetime msec localtime show-timezone
service password-encryption
service compress-config
no service dhcp
service unsupported-transceiver
!
hostname SW
!
boot-start-marker
boot-end-marker
!
!
vrf definition Mgmt-vrf
 --More-- 
 address-family ipv4
 exit-address-family
 !
 address-family ipv6
 exit-address-family
!
logging console critical
logging monitor critical
!
!
aaa session-id common
switch 1 provision ws-c3850-12s
!
!
no ip source-route
no ip gratuitous-arps
ip icmp rate-limit unreachable 1000
!
ip domain-name
ip name-server 172.16.201.101
!
!
qos queue-softmax-multiplier 100
vtp domain
vtp mode transparent
udld aggressive
!
!
errdisable recovery cause udld
errdisable recovery cause bpduguard
errdisable recovery cause security-violation
errdisable recovery cause channel-misconfig
errdisable recovery cause pagp-flap
errdisable recovery cause dtp-flap
errdisable recovery cause link-flap
errdisable recovery cause sfp-config-mismatch
errdisable recovery cause gbic-invalid
errdisable recovery cause psecure-violation
errdisable recovery cause port-mode-failure
errdisable recovery cause dhcp-rate-limit
errdisable recovery cause mac-limit
errdisable recovery cause vmps
errdisable recovery cause storm-control
errdisable recovery cause inline-power
errdisable recovery cause loopback
diagnostic bootup level minimal
!
spanning-tree mode rapid-pvst
spanning-tree loopguard default
spanning-tree portfast default
spanning-tree portfast bpduguard default
spanning-tree extend system-id
spanning-tree vlan 32,101,172,201 priority 4096
hw-switch switch 1 logging onboard message level 3
!
redundancy
 mode sso
!
vlan 3
!
vlan 5
 !
vlan 6
 !
vlan 2
 !
Vlan 8
 !
vlan 11
 !
vlan 12
 !
vlan 5
!
vlan 21
 name UNUSED
no cdp run
!
ip tcp synwait-time 10
ip ssh time-out 30
ip ssh version 2
!
!
! 
!
interface Null0
 no ip unreachables
!
 interface GigabitEthernet0/0
 vrf forwarding Mgmt-vrf
 no ip address
 shutdown
 negotiation auto
!
interface GigabitEthernet1/0/1
 description spare
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/2
  switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
interface GigabitEthernet1/0/3
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/4
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/5
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/6
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/7
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
 interface GigabitEthernet1/0/8
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 speed 100
 duplex full
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/9
switchport access vlan 8
 switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/10
 switchport access vlan 8
switchport mode access
 no logging event link-status
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/11
 switchport access vlan 8
 switchport mode access
 no logging event link-status
 speed 100
 duplex full
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
 spanning-tree portfast
!
interface GigabitEthernet1/0/12
 switchport trunk native vlan 55
 switchport trunk allowed vlan 8
 switchport mode trunk
 switchport nonegotiate
 no logging event link-status
 duplex full
 storm-control broadcast level 50.00 20.00
 storm-control multicast level 5.00 2.00
!
!
interface Vlan1
 no ip address
 no ip route-cache
 shutdown
!
interface Vlan8
 ip address 172.16.150.200 255.255.240.0
!
interface Vlan3
 no ip address
 no ip redirects
 no ip unreachables
 no ip proxy-arp
 no ip route-cache
!
ip default-gateway 172.16.201.27
ip forward-protocol nd
no ip http server
no ip http secure-server
!
ip access-list extended ALL_IP_TRAFFIC
 permit ip any any
!
!
speed 100, duplex fulllines in that config. Remember that if you do that, you absolutely will have to make sure that the device at the other end of the given cable has the same fixed setting, lest you risk duplex mismatches. To restrict a Gigabit (switch)port to 100Mbps,speed auto 100(orspeed auto 100 10to also support 10Mbps) is the less cumbersome choice.
