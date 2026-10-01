---
id: collect-261001-cisco/cisco/questions-69604-best-practice-advice-for-cisco-3560-behind-virtualised-opnsense-6d43c3f2
title: "Best practice / advice for cisco 3560 behind Virtualised OPNsense"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-69604-best-practice-advice-for-cisco-3560-behind-virtualised-opnsense-6d43c3f2.md
source_anchor: ""
source_lines: [1, 207]
sha256: 06a161b432cf5433c80282c0b31d09a6f55680b6b68b1c7a694614387d660b71
---

# Best practice / advice for cisco 3560 behind Virtualised OPNsense

*Score : 1 | Source : https://networkengineering.stackexchange.com/questions/69604/best-practice-advice-for-cisco-3560-behind-virtualised-opnsense*

Currently I run the below scheme. The problem is that when ESXi is down it takes down all the network with it, so I have to boot up another router to log in to it.
 WAN > OPNsense (in ESXi)> switch (C3560G)
OPNsense is DHCP server
Thing is that i would like the switch to deal with the DHCP, to avoid being naked when ESXi needs maintenance.
Also need to setup a VPN on OPNsense but will that be possible if 3560 is acting as the DHCP server?
Since the 3560 is L3 i could have it act as DHCP server and handle the VLANs, there's some uncertainty in an OpenVPN VLAN creation by the OPNsense.
edit1:
This is a mockup of current/future network. Some stuff is left out but are of no importance.
Current ISPmodem cannot be set to bridge, waiting for one with different FW, that will allow bridge mode.
There is a lot of space for improvement and while the office is not active for the next 2-3 weeks there is no actual rush. All comments are welcome.
About the OPNsense situation, I'm sorry if I overstepped here, I am aware of the general idea for rules here but as at my last place (~150 on site, ~20 remote users) we had a pfSense running, it actually didn't even cross my mind that this could potentially be off topic. Next edit will be with the mockup of the switch configuration with which I will require quite a lot of assistance.
So here's the very basic switch config
    Current configuration : 2717 bytes
!
version 12.2(37)SE1
no service timestamps log datetime msec
no service timestamps debug datetime msec
no service password-encryption
!
hostname c3560g
!
!
!
ip dhcp excluded-address 10.0.0.0 10.0.0.10
ip dhcp excluded-address 10.0.10.0 10.0.10.10
ip dhcp excluded-address 10.0.20.0 10.0.20.10
ip dhcp excluded-address 10.0.30.0 10.0.30.10
!
ip dhcp pool mgmt
 network 10.0.0.0 255.255.255.0
 default-router 10.0.0.1
 dns-server 10.0.0.2
 domain-name LAN
ip dhcp pool trusted
 network 10.0.10.0 255.255.255.0
 default-router 10.0.10.1
 dns-server 10.0.0.2
ip dhcp pool untrusted
 network 10.0.20.0 255.255.255.0
 default-router 10.0.20.1
 dns-server 10.0.0.2
ip dhcp pool guest
 network 10.0.30.0 255.255.255.0
 default-router 10.0.30.1
 dns-server 10.0.0.2
!
!
ip routing
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
!
!
!
spanning-tree mode pvst
!
!
!
!
!
!
interface FastEthernet0/1
 switchport mode access
 switchport nonegotiate
 spanning-tree portfast
!
interface FastEthernet0/2
 switchport mode access
 switchport nonegotiate
 spanning-tree portfast
!
interface FastEthernet0/3
 switchport access vlan 10
 switchport mode access
 switchport nonegotiate
 spanning-tree portfast
!
interface FastEthernet0/4
 switchport access vlan 10
 switchport mode access
 switchport nonegotiate
 spanning-tree portfast
!
interface FastEthernet0/5
 switchport access vlan 20
 switchport mode access
 switchport nonegotiate
 spanning-tree portfast
!
interface FastEthernet0/6
 switchport access vlan 30
 switchport mode access
 switchport nonegotiate
 spanning-tree portfast
!
interface FastEthernet0/7
!
interface FastEthernet0/8
!
interface FastEthernet0/9
!
interface FastEthernet0/10
!
interface FastEthernet0/11
!
interface FastEthernet0/12
!
interface FastEthernet0/13
!
interface FastEthernet0/14
!
interface FastEthernet0/15
!
interface FastEthernet0/16
!
interface FastEthernet0/17
!
interface FastEthernet0/18
!
interface FastEthernet0/19
!
interface FastEthernet0/20
!
interface FastEthernet0/21
!
interface FastEthernet0/22
!
interface FastEthernet0/23
 switchport trunk encapsulation dot1q
 switchport mode trunk
!
interface FastEthernet0/24
 switchport trunk encapsulation dot1q
 switchport mode trunk
!
interface GigabitEthernet0/1
!
interface GigabitEthernet0/2
!
interface Vlan1
 ip address 10.0.0.1 255.255.255.0
!
interface Vlan10
 mac-address 0090.0c64.7301
 ip address 10.0.10.1 255.255.255.0
!
interface Vlan20
 mac-address 0090.0c64.7302
 ip address 10.0.20.1 255.255.255.0
!
interface Vlan30
 mac-address 0090.0c64.7303
 ip address 10.0.30.1 255.255.255.0
!
ip default-gateway 10.0.0.2
ip classless
!
ip flow-export version 9
!
!
!
!
!
!
!
!
line con 0
 logging synchronous
!
line aux 0
!
line vty 0 4
 login
!
!
!
!
end
Now I will need to add some access list, hopefully avoiding restricting each vlan one by one, but will come to it if nothing pops up.

---

### Reponse — score 0

Yes you can setup a VPN server and not acting as your internal network's DHCP. Your VPN system will be a DHCP and offer IPs to externally connected clients. Your L3 switch will be offering IPs to your internal network.
For this scenario to work, I have assigned another IP subnet for VPN use only, no VLAN. Of course this subnet has to be routed. I used static routes for communication between the VPN IP subnet and the rest of the IP subnets/VLANs. OpenVPN is installed on a Debian Linux VM, I can't tell if you have the flexibility to configure everything needed on OPNsense OpenVPN server, but technically it's possible to achieve what you're asking.
I don't have any proposal if you need your VPN clients to be offered IP on the same subnets and VLANs as your internal network clients.

---

### Reponse — score 0

My recommendation would be to improve the reliability of the virtual machine hosting platform, not to move the DHCP service to the Cisco switch. I find the DHCP service capabilities of the Catalyst switch platform limited compared to a fully featured DHCP server (hopefully Opnsense has those features). If you need to make the router and DHCP functions more reliable, simply make the router more reliable by using a proper design for that system. If using a virtual machine hosting solution, have dual hosts (or more) to fit your needs and use a failover design for the Opnsense router/firewall setup or move to a firewall solution that provides that functionality. I believe Opnsense can provide that function and you can have each instance on a separate host so a host can be offline for maintenance and you will not lose your internet or DHCP service availability.
Let the switch act as a switch (Layer 3 if it helps your design) but keep IP address provisioning and DNS and associated services on a server that is designed for that functionality.
