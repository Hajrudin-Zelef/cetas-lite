---
id: collect-261001-general-networking/general-networking/questions-1632340-firewall-logs-show-query-replied-back-to-mdns-11385ba6
title: "Firewall logs show query replied back to mdns"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["decode", "ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-1632340-firewall-logs-show-query-replied-back-to-mdns-11385ba6.md
source_anchor: ""
source_lines: [1, 79]
sha256: b5b37dd21db13a8d49fc37d7bce5fbe0177aec54f856c2dadec5467cf23c45df
---

# Firewall logs show query replied back to mdns

*Score : 0 | Source : https://superuser.com/questions/1632340/firewall-logs-show-query-replied-back-to-mdns*

This setup is from the network (192.168.29.0/27) in VMware Workstation as in the Network Diagram
I have setup a Windows Server VM (192.168.28.47/27), and a firewall VM NIC (192.168.1.21/24) which is a Bridged interface on vmnet0 of VMware Workstation.
I have an all-traffic-allowed rule on both LAN and WAN interfaces of the firewall. The rule does not seem to be working for some reason. Server VM receives ping replies from 8.8.8.8 but cannot access internet. The NAT interface is 192.168.47.140, and NAT gateway is 192.168.47.2. I disabled the firewall and tested, the server accesses internet fine.
tcpdump on WAN without any host
root@firewallwm:~ # tcpdump -i em1 -n
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em1, link-type EN10MB (Ethernet), capture size 262144 bytes
17:53:05.670752 IP 192.168.47.140.27089 > 188.165.3.28.123: NTPv4, Client, length 48
17:53:07.652383 IP 192.168.47.140.10811 > 188.125.64.7.123: NTPv4, Client, length 48
17:53:07.673695 IP 188.125.64.7.123 > 192.168.47.140.10811: NTPv4, Server, length 48
17:53:08.624653 IP 192.168.47.140.28314 > 91.209.0.17.123: NTPv4, Client, length 48
17:53:08.672014 IP 91.209.0.17.123 > 192.168.47.140.28314: NTPv4, Server, length 48
17:53:09.680357 IP 192.168.47.140.44852 > 80.237.128.148.123: NTPv4, Client, length 48
17:53:09.680434 IP 192.168.47.140.58287 > 162.159.200.123.123: NTPv4, Client, length 48
17:53:09.698470 IP 162.159.200.123.123 > 192.168.47.140.58287: NTPv4, Server, length 48
17:53:09.720681 IP 80.237.128.148.123 > 192.168.47.140.44852: NTPv4, Server, length 48
17:53:12.633164 IP 192.168.47.140.59261 > 94.199.173.123.123: NTPv4, Client, length 48
17:53:12.671315 IP 94.199.173.123.123 > 192.168.47.140.59261: NTPv4, Server, length 48
17:53:54.738729 IP6 fe80::20c:29ff:fee8:8d15.546 > ff02::1:2.547: dhcp6 solicit
17:58:48.516746 IP 192.168.47.1.64045 > 239.255.255.250.1900: UDP, length 173
17:58:49.517887 IP 192.168.47.1.64045 > 239.255.255.250.1900: UDP, length 173
17:58:50.520981 IP 192.168.47.1.64045 > 239.255.255.250.1900: UDP, length 173
17:58:51.522420 IP 192.168.47.1.64045 > 239.255.255.250.1900: UDP, length 173
tcpdump on LAN for Server VM
root@firewallwm:~ # tcpdump -i em3 host 192.168.28.47
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on em3, link-type EN10MB (Ethernet), capture size 262144 bytes
17:37:05.247138 IP 192.168.28.47 > dns.google: ICMP echo request, id 1, seq 54, length 40
17:37:05.260352 IP dns.google > 192.168.28.47: ICMP echo reply, id 1, seq 54, length 40
17:37:06.269572 IP 192.168.28.47 > dns.google: ICMP echo request, id 1, seq 55, length 40
17:37:06.283080 IP dns.google > 192.168.28.47: ICMP echo reply, id 1, seq 55, length 40
17:37:07.289858 IP 192.168.28.47 > dns.google: ICMP echo request, id 1, seq 56, length 40
17:37:07.298693 IP dns.google > 192.168.28.47: ICMP echo reply, id 1, seq 56, length 40
17:37:08.302397 IP 192.168.28.47 > dns.google: ICMP echo request, id 1, seq 57, length 40
17:37:08.311143 IP dns.google > 192.168.28.47: ICMP echo reply, id 1, seq 57, length 40
17:37:09.866366 ARP, Request who-has firewallwm.vlab.lab (00:0c:29:e8:8d:29 (oui Unknown)) tell 192.168.28.47, length 46
17:37:09.866397 ARP, Reply firewallwm.vlab.lab is-at 00:0c:29:e8:8d:29 (oui Unknown), length 28
17:37:36.493211 IP 192.168.28.47.60468 > 239.255.255.250.1900: UDP, length 173
17:37:37.509653 IP 192.168.28.47.60468 > 239.255.255.250.1900: UDP, length 173
17:37:38.523971 IP 192.168.28.47.60468 > 239.255.255.250.1900: UDP, length 173
17:37:39.540750 IP 192.168.28.47.60468 > 239.255.255.250.1900: UDP, length 173
17:39:36.498026 IP 192.168.28.47.60469 > 239.255.255.250.1900: UDP, length 173
17:39:37.510526 IP 192.168.28.47.60469 > 239.255.255.250.1900: UDP, length 173
17:39:38.512683 IP 192.168.28.47.60469 > 239.255.255.250.1900: UDP, length 173
17:39:39.512937 IP 192.168.28.47.60469 > 239.255.255.250.1900: UDP, length 173
17:39:53.018203 IP 192.168.28.47.netbios-ns > 192.168.28.63.netbios-ns: NBT UDP PACKET(137): QUERY; REQUEST; BROADCAST
17:39:53.027039 IP 192.168.28.47.mdns > 224.0.0.251.mdns: 0 A (QM)? google.local. (30)
17:39:53.032230 IP 192.168.28.47.mdns > 224.0.0.251.mdns: 0 AAAA (QM)? google.local. (30)
17:39:53.033009 IP 192.168.28.47.56963 > 224.0.0.252.5355: UDP, length 24
17:39:53.040405 IP 192.168.28.47.62251 > 224.0.0.252.5355: UDP, length 24
17:39:53.445140 IP 192.168.28.47.56963 > 224.0.0.252.5355: UDP, length 24
17:39:53.460646 IP 192.168.28.47.62251 > 224.0.0.252.5355: UDP, length 24
17:39:53.771373 IP 192.168.28.47.netbios-ns > 192.168.28.63.netbios-ns: NBT UDP PACKET(137): QUERY; REQUEST; BROADCAST
17:39:54.038374 IP 192.168.28.47.mdns > 224.0.0.251.mdns: 0 A (QM)? google.local. (30)
17:39:54.039114 IP 192.168.28.47.mdns > 224.0.0.251.mdns: 0 AAAA (QM)? google.local. (30)
17:39:54.536041 IP 192.168.28.47.netbios-ns > 192.168.28.63.netbios-ns: NBT UDP PACKET(137): QUERY; REQUEST; BROADCAST
For some reason (or setting) the replies from the WAN 192.168.47.1 are going to 224.0.0.251 which I read is Multicast DNS, and 239.255.255.250.
Thus 2 questions, 1) How are ping replies coming from Google's DNS ?!, and 2) why aren't the packets reaching the WAN interface.
New to firewalls, so I'm definitely missing something.

---

### Reponse (acceptee) — score 1

I think you are misinterpreting your packet traces and it's making you ask the wrong questions. Let me try to clarify some things:
From the first packet trace:
IP 192.168.47.1.64045 > 239.255.255.250.1900: UDP
This is your router advertising UPnP services. This packet was created by your router itself; this packet was not forwarded from somewhere else.
From the second packet trace:
IP 192.168.28.47.60468 > 239.255.255.250.1900: UDP
This is your Windows Server VM trying to do UPnP. This packet was created by your Windows Server VM. This packet was not forwarded from somewhere else.
IP 192.168.28.47.mdns > 224.0.0.251.mdns: 0 A (QM)? google.local. (30)
IP 192.168.28.47.mdns > 224.0.0.251.mdns: 0 AAAA (QM)? google.local. (30)
This looks like something on your Windows Server VM doing a lookup for hostname "google". Note: it's not looking for "google.com.", just "google" with no dots or top-level domains. So Windows is searching for hostname "google" in all search domains it's configured for, including ".local" (multicast DNS). So it's asking, via multicasts on the local LAN, "is there any device here named 'google'"? (It does not appear to be getting an answer.)
This is not a DNS answer (reply) from a unicast DNS server or a local mDNS host. It's an mDNS query (request) that was created by your Windows Server VM.
