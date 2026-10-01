---
id: collect-261001-general-networking/general-networking/questions-5644-pppoe-with-nat-configuration-broken-746b27a9
title: "questions-5644-pppoe-with-nat-configuration-broken-746b27a9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "rotary", "voice"]
source: docs/RAG/collect-261001-general-networking/questions-5644-pppoe-with-nat-configuration-broken-746b27a9.md
source_anchor: ""
source_lines: [1, 131]
sha256: 7d76a4a437c9fcc11a08fb3f74c4e254601ad2b1356c6366165bf463b7b1f7d3
---

# questions-5644-pppoe-with-nat-configuration-broken-746b27a9

We are trying to replace our existing router with a newly purchased Cisco router:
GigabitEthernet0/0: PPPoE GigabitEthernet0/1: 192.168.2.1
The PPPoE is getting established perfectly, and I can ping the outside from inside CLI however, the clients cannot ping the outside world. As mentioned, this is replace an existing router with the same IP (ie, 192.168.2.1) therefore, the clients are all setup and working fine when not using the Cisco router.
My configuration is as follows:
    Current configuration : 2698 bytes
!
version 15.1
service timestamps debug datetime msec
service timestamps log datetime msec
no service password-encryption
!
hostname r1
!
boot-start-marker
boot-end-marker
!
!
enable secret 4 <secret>
!
aaa new-model
!
aaa session-id common
!
dot11 syslog
ip source-route
!
ip cef
!
ip domain name test.com
no ipv6 cef
!
multilink bundle-name authenticated
!
voice-card 0
!
crypto pki token default removal timeout 0
!
license udi pid CISCO3825 sn FTX1039A0AA
username admin privilege 15 secret 4 <secret>
!
redundancy
!
interface GigabitEthernet0/0
 no ip address
 ip nat outside
 ip virtual-reassembly in
 duplex auto
 speed auto
 media-type rj45
 pppoe enable group global
 pppoe-client dial-pool-number 1
 no cdp enable
!
interface GigabitEthernet0/1
 ip address 192.168.2.1 255.255.255.0
 ip nat inside
 ip virtual-reassembly in
 duplex auto
 speed auto
 media-type rj45
!
interface Dialer0
 ip address negotiated
 no ip redirects
 no ip proxy-arp
 ip mtu 1492
 ip flow ingress
 encapsulation ppp
 ip route-cache policy
 dialer pool 1
 dialer-group 1
 ppp authentication pap callin
 ppp pap sent-username <username> password 0 <password>
 ppp ipcp dns request accept
 ppp ipcp route default
 ppp ipcp address accept
 no cdp enable
!
ip forward-protocol nd
no ip http server
no ip http secure-server
!
!
ip nat pool voip-rtp 192.168.2.5 192.168.2.5 netmask 255.255.255.0 type rotary
ip nat inside source list 1 interface GigabitEthernet0/1 overload
ip nat inside source static tcp 192.168.2.10 25 77.71.31.187 25 extendable
ip nat inside source static tcp 192.168.2.5 53 77.71.31.187 53 extendable
ip nat inside source static tcp 192.168.2.15 80 77.71.31.187 80 extendable
ip nat inside source static tcp 192.168.2.15 443 77.71.31.187 443 extendable
ip nat inside source static tcp 192.168.2.10 465 77.71.31.187 465 extendable
ip nat inside source static tcp 192.168.2.10 993 77.71.31.187 993 extendable
ip nat inside source static tcp 192.168.2.10 995 77.71.31.187 995 extendable
ip nat inside source static udp 192.168.2.5 5060 77.71.31.187 5060 extendable
ip nat inside source static udp 192.168.2.20 5080 77.71.31.187 5080 extendable
ip nat inside destination list 114 pool voip-rtp
ip route 0.0.0.0 0.0.0.0 GigabitEthernet0/0
ip route 64.230.11.5 255.255.255.255 GigabitEthernet0/0
!
access-list 1 permit 192.168.2.0 0.0.0.255
access-list 114 permit udp any any range 8000 65535
!
control-plane
!
mgcp profile default
!
line con 0
line aux 0
line vty 0 4
 transport input ssh
!
scheduler allocate 20000 1000
end 
Things I tried:
Statically setting the IP route: ip route 0.0.0.0 0.0.0.0 GigabitEthernet0/0 ip route 64.230.11.5 255.255.255.255 GigabitEthernet0/0 Although this is probably not needed for auto established PPPoE connection?
Update
IP NAT Trans Output
Pro Inside global         Inside local          Outside local         Outside global
tcp 77.71.31.187:53       192.168.2.5:53        ---                   ---
udp 77.71.31.187:5060     192.168.2.5:5060      <voip supplier's ip>:5060   <voip supplier's ip>:5060
udp 77.71.31.187:5060     192.168.2.5:5060      <voip supplier's ip>:5060   <voip supplier's ip>:5060
udp 77.71.31.187:5060     192.168.2.5:5060      ---                   ---
tcp 77.71.31.187:25       192.168.2.10:25       217.160.226.122:25    217.160.226.122:25
tcp 77.71.31.187:25       192.168.2.10:25       ---                   ---
tcp 77.71.31.187:465      192.168.2.10:465      ---                   ---
tcp 77.71.31.187:993      192.168.2.10:993      ---                   ---
tcp 77.71.31.187:995      192.168.2.10:995      ---                   ---
tcp 77.71.31.187:80       192.168.2.15:80       66.249.74.2:33795     66.249.74.2:33795
tcp 77.71.31.187:80       192.168.2.15:80       66.249.74.2:50496     66.249.74.2:50496
tcp 77.71.31.187:80       192.168.2.15:80       ---                   ---
tcp 77.71.31.187:443      192.168.2.15:443      ---                   ---
udp 77.71.31.187:5080     192.168.2.20:5080     ---                   ---
