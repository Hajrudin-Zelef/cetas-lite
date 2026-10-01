---
id: collect-261001-cisco/cisco/questions-1912-how-to-host-server-in-dmz-zone-of-an-asa-82d203ba
title: "questions-1912-how-to-host-server-in-dmz-zone-of-an-asa-82d203ba"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-1912-how-to-host-server-in-dmz-zone-of-an-asa-82d203ba.md
source_anchor: ""
source_lines: [1, 38]
sha256: d8e7e6fcaf1ddfd31f593d56cc4613cefcf6114b808a55d374ef69c14d5feb14
---

# questions-1912-how-to-host-server-in-dmz-zone-of-an-asa-82d203ba

I have an ASA5525-X with 9.1.2. On it are several interfaces, but primarily I'm looking at:
(fake subnets)
- inside 10.0.0.0/24, security level 100
- outside 10.0.200.0/24, security level 0
- DMZ 10.0.100.0/24, security level 50
I have a DNS server in DMZ, 10.0.100.1 which I can access from inside without issue. However, I want it to show up as 10.0.200.95 (not a real IP for this example) to people on the internet. I have what I thought was needed for this to work, but when I test it out, the packets are being dropped by the default acl.
Pertinent config pieces:
interface GigabitEthernet0/0
 nameif outside
 security-level 0
 ip address 10.0.200.194 255.255.255.192 
interface GigabitEthernet0/6
 nameif DMZ
 security-level 50
 ip address 10.0.100.254 255.255.255.0 
interface GigabitEthernet0/7
 nameif inside
 security-level 100
 ip address 10.0.0.254 255.255.255.0 
object network DMZ-DNS-Server-1
 host 10.0.100.1
nat (inside,outside-backup) source static Company-MPLS-Network-Group Company-MPLS-Network-Group destination static VPN VPN no-proxy-arp route-lookup
nat (inside,outside) source static Company-MPLS-Network-Group Company-MPLS-Network-Group destination static VPN VPN no-proxy-arp route-lookup
nat (inside,outside) source static Company-MPLS-Network-Group Company-MPLS-Network-Group destination static Site-Site-VPN Site-Site-VPN no-proxy-arp route-lookup
nat (inside,outside-backup) source static Company-MPLS-Network-Group Company-MPLS-Network-Group destination static Site-Site-VPN Site-Site-VPN no-proxy-arp route-lookup
nat (inside,DMZ) source static Company-MPLS-Network-Group Company-MPLS-Network-Group destination static DMZ DMZ no-proxy-arp route-lookup
nat (DMZ,outside) source static DMZ DMZ destination static Site-Site-VPN Site-Site-VPN no-proxy-arp route-lookup
nat (DMZ,outside-backup) source static DMZ DMZ destination static Site-Site-VPN Site-Site-VPN no-proxy-arp route-lookup
object network DMZ-DNS-Server-1
 nat (DMZ,outside) static 10.0.200.195 net-to-net
nat (inside,outside-backup) after-auto source dynamic any interface
nat (inside,outside) after-auto source dynamic any interface
nat (DMZ,outside) after-auto source dynamic any interface
nat (DMZ,outside-backup) after-auto source dynamic any interface
access-list traffic-in-outside extended permit tcp any host 10.0.200.195 eq domain 
access-list traffic-in-outside extended permit udp any host 10.0.200.195 eq domain 
access-group traffic-in-outside in interface outside
Any ideas?
