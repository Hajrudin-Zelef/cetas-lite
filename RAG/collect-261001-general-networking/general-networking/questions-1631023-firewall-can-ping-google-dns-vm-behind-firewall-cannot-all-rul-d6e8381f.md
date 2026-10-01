---
id: collect-261001-general-networking/general-networking/questions-1631023-firewall-can-ping-google-dns-vm-behind-firewall-cannot-all-rul-d6e8381f
title: "Firewall Can Ping Google DNS, VM behind Firewall cannot, all rules allowed"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["decode", "ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-1631023-firewall-can-ping-google-dns-vm-behind-firewall-cannot-all-rul-d6e8381f.md
source_anchor: ""
source_lines: [1, 66]
sha256: daa38afca67dfee0fe3317cdecc44e8e45ad6e5ac25c2f1eb446f7623e692bbf
---

# Firewall Can Ping Google DNS, VM behind Firewall cannot, all rules allowed

*Score : 1 | Source : https://superuser.com/questions/1631023/firewall-can-ping-google-dns-vm-behind-firewall-cannot-all-rules-allowed*

I have the following network setup https://i.ibb.co/wwPLH2H/Network.png
All traffic from 10.0.64.0 / 27 behind FirewallB (firewallsm) reaches 192.168.28.0 / 27 network via the LAN interface of FirewallA (firewallwm), and the same traffic also reaches internet in the same way, as follows:
10.0.64.42 (VM) > FirewallB (LAN) > FirewallA (LAN) > FirewallA (WAN) > Laptop's Wireless NIC > Wifi Router
Strangely FirewallB (firewallsm) can ping Google DNS but the VM 10.0.64.42 for some reason cannot ping Google DNS. I have set all protocols, ports as allowed on FirewallB (firewallsm) to reach FirewallA (firewallwm).
FirewallA (firewallwm)
Gateway - https://i.ibb.co/bRC8P8G/Firewall-A-GW.png
LAN Interface Rule - https://i.ibb.co/XWSnLRd/Firewall-A-1-dell-Rule.png
WAN Interface Rule - https://i.ibb.co/zZwcnjJ/Firewall-A-2-WAN-Rule.png
FirewallA (firewallsm) logs show 10.0.64.42 traffic is allowed through its WAN
Log - https://i.ibb.co/9cqPSW7/Firewall-A-Packet-log.png
Log - https://i.ibb.co/kVTX31B/Firewall-A-Packet-log-2.png
FirewallB (firewallsm)
Gateway - https://i.ibb.co/pPQC1p8/Firewall-B-GW.png
LAN Rule - https://i.ibb.co/VY9vFVL/Firewall-B-1-dell-Rule.png
tcpdump for 10.0.64.42 VM on FirewallB LAN (em0)
root@firewallsm:~ # tcpdump -i vmx0 host 10.0.64.42 and host 8.8.8.8 and icmp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on vmx0, link-type EN10MB (Ethernet), capture size 262144 bytes
06:17:09.324409 IP 10.0.64.42 > dns.google: ICMP echo request, id 1, seq 1, length 40
06:17:13.853917 IP 10.0.64.42 > dns.google: ICMP echo request, id 1, seq 2, length 40
06:17:18.858484 IP 10.0.64.42 > dns.google: ICMP echo request, id 1, seq 3, length 40
tcpdump for 10.0.64.42 VM on FirewallA LAN (em0)
root@firewallwm:~ # tcpdump -i em0 host 8.8.8.8 and icmp -n
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em0, link-type EN10MB (Ethernet), capture size 262144 bytes
06:17:09.335331 IP 10.0.64.42 > 8.8.8.8: ICMP echo request, id 1, seq 1, length 40
06:17:13.865408 IP 10.0.64.42 > 8.8.8.8: ICMP echo request, id 1, seq 2, length 40
06:17:23.870819 IP 10.0.64.42 > 8.8.8.8: ICMP echo request, id 1, seq 4, length 40
tcpdump for 10.0.64.42 VM on FirewallA WAN (em1)
root@firewallwm:~ # tcpdump -i em1 host 8.8.8.8 and icmp -n
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em1, link-type EN10MB (Ethernet), capture size 262144 bytes
Ping from FirewallB successful in pinging Google DNS
root@firewallsm:~ # ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: icmp_seq=0 ttl=127 time=13.196 ms
64 bytes from 8.8.8.8: icmp_seq=1 ttl=127 time=12.625 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=127 time=12.609 ms
^C
--- 8.8.8.8 ping statistics ---
16 packets transmitted, 16 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 11.705/13.026/15.910/1.111 ms
tcpdump of FirewallA LAN (em0)
root@firewallwm:~ # tcpdump -i em0 host 8.8.8.8 and icmp -n
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em0, link-type EN10MB (Ethernet), capture size 262144 bytes
06:22:21.816908 IP 192.168.10.7 > 8.8.8.8: ICMP echo request, id 23594, seq 0, length 64
06:22:21.827095 IP 8.8.8.8 > 192.168.10.7: ICMP echo reply, id 23594, seq 0, length 64
06:22:22.876598 IP 192.168.10.7 > 8.8.8.8: ICMP echo request, id 23594, seq 1, length 64
06:22:22.886317 IP 8.8.8.8 > 192.168.10.7: ICMP echo reply, id 23594, seq 1, length 64
06:22:23.948947 IP 192.168.10.7 > 8.8.8.8: ICMP echo request, id 23594, seq 2, length 64
06:22:23.957978 IP 8.8.8.8 > 192.168.10.7: ICMP echo reply, id 23594, seq 2, length 64
tcpdump of FirewallA WAN (em1)
root@firewallwm:~ # tcpdump -i em1 host 8.8.8.8 and icmp -n
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em1, link-type EN10MB (Ethernet), capture size 262144 bytes
06:22:21.817029 IP 192.168.47.132 > 8.8.8.8: ICMP echo request, id 24689, seq 0, length 64
06:22:21.826993 IP 8.8.8.8 > 192.168.47.132: ICMP echo reply, id 24689, seq 0, length 64
06:22:22.876700 IP 192.168.47.132 > 8.8.8.8: ICMP echo request, id 24689, seq 1, length 64
06:22:22.886219 IP 8.8.8.8 > 192.168.47.132: ICMP echo reply, id 24689, seq 1, length 64
06:22:23.949057 IP 192.168.47.132 > 8.8.8.8: ICMP echo request, id 24689, seq 2, length 64
06:22:23.957845 IP 8.8.8.8 > 192.168.47.132: ICMP echo reply, id 24689, seq 2, length 64
