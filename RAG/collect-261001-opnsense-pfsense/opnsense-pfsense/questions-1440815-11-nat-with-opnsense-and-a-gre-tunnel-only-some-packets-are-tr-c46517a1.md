---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1440815-11-nat-with-opnsense-and-a-gre-tunnel-only-some-packets-are-tr-c46517a1
title: "1:1 NAT with OPNsense and a GRE tunnel, only some packets are translated"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1440815-11-nat-with-opnsense-and-a-gre-tunnel-only-some-packets-are-tr-c46517a1.md
source_anchor: ""
source_lines: [1, 35]
sha256: 2a8f4b2aa8a2dfc5a97607d3edad57089e7e328ef4ee57f23ef6828148c2d714
---

# 1:1 NAT with OPNsense and a GRE tunnel, only some packets are translated

*Score : 0 | Source : https://superuser.com/questions/1440815/11-nat-with-opnsense-and-a-gre-tunnel-only-some-packets-are-translated*

My home router is a virtual machine running OPNsense. I have a dedicated server with an IPv4 subnet routed to it, and I have set up a GRE tunnel from my dedicated server to my home router so that I can use the IP addresses at home.
On the dedicated server, I have set up the subnet to be routed through the GRE tunnel.
In OPNsense, I have set up the GRE tunnel, and added the Virtual IPs.
The issue is, when I set up 1:1 NAT, only some packets have their source address translated. On the client which as 1:1 NAT set up, the internal IP is 192.168.254.106 and the external IP is (for example), 176.9.x.x  
Whenever I run ping 8.8.8.8 on the client, only the first ping returns.
ubuntu@ubuntu:~$ ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=55 time=32.8 ms
^C
--- 8.8.8.8 ping statistics ---
6 packets transmitted, 1 received, 83% packet loss, time 5081ms
rtt min/avg/max/mdev = 32.875/32.875/32.875/0.000 ms
This is what I see when I run tcpdump -i gre0 icmp on the OPNsense router. (gre0 being the GRE tunnel)
root@core1:~ # tcpdump -i gre0 icmp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on gre0, link-type NULL (BSD loopback), capture size 262144 bytes
12:26:35.361531 IP 176.9.x.x > google-public-dns-a.google.com: ICMP echo request, id 1784, seq 1, length 64
12:26:35.393693 IP google-public-dns-a.google.com > 176.9.x.x: ICMP echo reply, id 1784, seq 1, length 64
12:26:36.362530 IP 192.168.254.106 > google-public-dns-a.google.com: ICMP echo request, id 1784, seq 2, length 64
12:26:37.371299 IP 192.168.254.106 > google-public-dns-a.google.com: ICMP echo request, id 1784, seq 3, length 64
12:26:38.395302 IP 192.168.254.106 > google-public-dns-a.google.com: ICMP echo request, id 1784, seq 4, length 64
12:26:39.419105 IP 192.168.254.106 > google-public-dns-a.google.com: ICMP echo request, id 1784, seq 5, length 64
12:26:40.443423 IP 192.168.254.106 > google-public-dns-a.google.com: ICMP echo request, id 1784, seq 6, length 64
Is there a bug, or did I do something wrong? It is supposed to translate the source IP of all packets, but it's only doing it for some. (if you let the ping running for a while, one packet will be translated again and then it'll stop translating the packets after that.)

---

### Reponse (acceptee) — score 1

Unfortunately there is only a workaround that I was able to successfully use.
At the time the question was asked -and still today-, there is a bug in the modems from my ISP that cause them to break when dealing with GRE or IPIP. The solution is Foo over UDP with GRE in it, which works well. However, OPNsense cannot work with that as far as I know, so I switched to a dedicated Ubuntu VM for it. Now everything works fine.
