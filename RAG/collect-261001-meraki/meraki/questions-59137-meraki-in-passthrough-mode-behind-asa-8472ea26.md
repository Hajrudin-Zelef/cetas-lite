---
id: collect-261001-meraki/meraki/questions-59137-meraki-in-passthrough-mode-behind-asa-8472ea26
title: "questions-59137-meraki-in-passthrough-mode-behind-asa-8472ea26"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/questions-59137-meraki-in-passthrough-mode-behind-asa-8472ea26.md
source_anchor: ""
source_lines: [1, 221]
sha256: 8e209710f5a22785aaf5635d140516007f4ccfc94c8c2b6c3757d1ab54382dea
---

# questions-59137-meraki-in-passthrough-mode-behind-asa-8472ea26

I am having trouble getting bidirectional traffic through a Meraki box behind a ASA 5508-X. I can see requests from the client VPN going from a client device to a server on our LAN, the server responds, but the Meraki never sends it to the client. (The Meraki is in passthrough mode.) Not sure if it's necessary to define the remote subnets more clearly, I previously had the Client subnet defined locally, but that was no better. For reference, 192.168.3.0/24 is Clients, and 192.168.129.0/24 is another site.
I have the proper ports forwarded, I believe. I've added default routes for both Site-To-Site and Client. I think it's probably a NAT problem. Here's the ASA Config.
ASA Version 9.8(2) 
!
hostname ciscoasa
enable password $xxx
names
!
interface GigabitEthernet1/1
 nameif outside
 security-level 0
 ip address xxx.xxx.xxx.xxx 255.255.255.252 
!
interface GigabitEthernet1/2
 nameif inside
 security-level 100
 ip address 192.168.0.254 255.255.254.0 
!
interface GigabitEthernet1/2.2
 description Isolated Guest traffic
 vlan 2
 nameif Guest
 security-level 60
 ip address 192.168.2.1 255.255.255.0 
!
interface GigabitEthernet1/3
 nameif VOIP
 security-level 100
 ip address 172.16.0.99 255.255.255.0 
!
interface GigabitEthernet1/4
 shutdown
 no nameif
 no security-level
 no ip address
!
interface GigabitEthernet1/5
 shutdown
 no nameif
 no security-level
 no ip address
!
interface GigabitEthernet1/6
 shutdown
 no nameif
 no security-level
 no ip address
!
interface GigabitEthernet1/7
 shutdown
 no nameif
 no security-level
 no ip address
!
interface GigabitEthernet1/8
 shutdown
 no nameif
 no security-level
 no ip address
!
interface Management1/1
 management-only
 no nameif
 no security-level
 no ip address
!
ftp mode passive
clock timezone CST -6
clock summer-time CDT recurring
same-security-traffic permit inter-interface
same-security-traffic permit intra-interface
object network inside
 subnet 192.168.0.0 255.255.254.0
object network phones
 subnet 172.16.0.0 255.255.255.0
object network Guests
 subnet 192.168.2.0 255.255.255.0
 description Isolated guest traffic
object network Meraki
 host 192.168.0.1
object network Meraki_natt
 host 192.168.0.1
object network bossVPN
 host 192.168.3.0
 description 255.255.255.0
access-list global_access extended permit udp any object Meraki eq isakmp 
access-list global_access extended permit udp any object Meraki eq 4500 
access-list internal extended permit ip any any 
access-list internal extended permit icmp any any 
access-list internal extended permit udp 192.168.0.0 255.255.254.0 host 8.8.8.8 eq domain 
access-list internal extended permit udp 192.168.0.0 255.255.254.0 host 8.8.4.4 eq domain 
access-list internal extended permit udp 192.168.0.0 255.255.254.0 host 9.9.9.9 eq domain 
access-list internal extended permit udp 192.168.0.0 255.255.254.0 host 75.75.76.76 eq domain 
access-list internal extended permit udp object Meraki any eq isakmp 
access-list internal extended permit udp object Meraki any eq 4500 
access-list external extended permit ip any any 
access-list external extended permit icmp any any 
access-list Guest_access_in extended deny ip any 192.168.0.0 255.255.254.0 
access-list Guest_access_in extended permit ip any any 
access-list Guest_access_in extended permit icmp any any 
pager lines 24
logging enable
logging asdm notifications
mtu outside 1500
mtu inside 1500
mtu Guest 1500
mtu VOIP 1500
no failover
no monitor-interface service-module 
icmp unreachable rate-limit 1 burst-size 1
no asdm history enable
arp timeout 14400
no arp permit-nonconnected
arp rate-limit 16384
!
object network inside
 nat (inside,outside) dynamic interface
object network phones
 nat (VOIP,outside) dynamic interface
object network Guests
 nat (Guest,outside) dynamic interface
object network Meraki
 nat (any,outside) static interface service udp isakmp isakmp 
object network Meraki_natt
 nat (any,outside) static interface service udp 4500 4500 
access-group external in interface outside
access-group internal in interface inside
access-group Guest_access_in in interface Guest
access-group internal in interface VOIP
route outside 0.0.0.0 0.0.0.0 xxx.xxx.xxx.xxx 1
route inside 192.168.3.0 255.255.255.0 192.168.0.1 1
route inside 192.168.129.0 255.255.255.0 192.168.0.1 1
timeout xlate 3:00:00
timeout pat-xlate 0:00:30
timeout conn 1:00:00 half-closed 0:10:00 udp 0:02:00 sctp 0:02:00 icmp 0:00:02
timeout sunrpc 0:10:00 h323 0:05:00 h225 1:00:00 mgcp 0:05:00 mgcp-pat 0:05:00
timeout sip 0:30:00 sip_media 0:02:00 sip-invite 0:03:00 sip-disconnect 0:02:00
timeout sip-provisional-media 0:02:00 uauth 0:05:00 absolute
timeout tcp-proxy-reassembly 0:01:00
timeout floating-conn 0:00:00
timeout conn-holddown 0:00:15
timeout igp stale-route 0:01:10
user-identity default-domain LOCAL
aaa authentication login-history
http server enable
http 192.168.0.0 255.255.254.0 inside
no snmp-server location
no snmp-server contact
service sw-reset-button
crypto ipsec security-association pmtu-aging infinite
crypto ca trustpool policy
telnet timeout 5
ssh stricthostkeycheck
ssh timeout 5
ssh key-exchange group dh-group1-sha1
console timeout 0
dhcpd auto_config outside
!
dhcpd address 192.168.1.2-192.168.1.249 inside
dhcpd dns 8.8.8.8 interface inside
dhcpd lease 86400 interface inside
dhcpd domain BII interface inside
dhcpd enable inside
!
dhcpd address 192.168.2.100-192.168.2.149 Guest
dhcpd dns 9.9.9.9 8.8.8.8 interface Guest
dhcpd lease 86400 interface Guest
dhcpd domain Guest interface Guest
dhcpd enable Guest
!
dhcpd address 172.16.0.120-172.16.0.219 VOIP
dhcpd dns 8.8.8.8 interface VOIP
dhcpd lease 86400 interface VOIP
dhcpd domain voip interface VOIP
dhcpd enable VOIP
!
threat-detection basic-threat
threat-detection statistics
threat-detection statistics tcp-intercept rate-interval 30 burst-rate 400 average-rate 200
dynamic-access-policy-record DfltAccessPolicy
!
class-map inspection_default
 match default-inspection-traffic
!
!
policy-map type inspect dns preset_dns_map
 parameters
  message-length maximum client auto
  message-length maximum 512
  no tcp-inspection
policy-map global_policy
 class inspection_default
  inspect dns preset_dns_map 
  inspect ftp 
  inspect h323 h225 
  inspect h323 ras 
  inspect rsh 
  inspect rtsp 
  inspect esmtp 
  inspect sqlnet 
  inspect skinny  
  inspect sunrpc 
  inspect xdmcp 
  inspect netbios 
  inspect tftp 
  inspect ip-options 
  inspect icmp 
 class class-default
  user-statistics accounting
!
service-policy global_policy global
prompt hostname context 
no call-home reporting anonymous
EDIT Adding a static route on a server on our internal network to the client vpn subnet allowed that server to work with the VPN. I'd rather the static route on the ASA mattered.
Curiously, if I load the PBX (172.16.0.35), I see two way traffic on the 'Client VPN pcap', but none of the 192.168.0.0/23 traffic will work.
Do routes from the ASA have to propagate to work? Reboots of connected devices or something?
EDIT 2 Some context. The Meraki is on GigabitEthernet1/2 between the ASA and our 52port L2 switch.
EDIT 3 I'm now leaning towards NAT being the problem again.
Meraki WAN traffic (limited by client VPN IP address)
Meraki LAN traffic (limited by client VPN IP address)
Can anyone help me with the NAT statement(s) to send all 192.168.3.0/24 traffic through 192.168.0.1?
