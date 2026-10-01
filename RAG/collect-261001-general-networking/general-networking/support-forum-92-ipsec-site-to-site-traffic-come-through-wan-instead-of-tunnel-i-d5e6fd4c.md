---
id: collect-261001-general-networking/general-networking/support-forum-92-ipsec-site-to-site-traffic-come-through-wan-instead-of-tunnel-i-d5e6fd4c
title: "support-forum-92-ipsec-site-to-site-traffic-come-through-wan-instead-of-tunnel-i-d5e6fd4c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-ipsec-site-to-site-traffic-come-through-wan-instead-of-tunnel-i-d5e6fd4c.md
source_anchor: ""
source_lines: [1, 30]
sha256: 26b362e70a45cb6921772313ab9abece79553b89638b3504a326afddb07451c6
---

# support-forum-92-ipsec-site-to-site-traffic-come-through-wan-instead-of-tunnel-i-d5e6fd4c

IPsec site to site, traffic come through wan instead of tunnel interface
Hi, I have setup IPsec s2s vpn between two site, A and B
A is behind a NAT router,
topology:
192.168.63.0/24 <-> A (VPN router) <-> NAT router <-> internet <-> B (fortigate) <-> 192.168.2.0/24
I've successfully established phase 2 IPsec tunnels between A and B, but I can't ping hosts from A to B, or B to A.
After some troubleshooting, I found somethings really weird:
If I ping from B to A, traffic is pass through the tunnel interface (tun02)
forti01 # diagnose sniffer packet tun02
interfaces=[tun02]
filters=[none]
4.162893 192.168.2.1 -> 192.168.63.110: icmp: echo request
5.178048 192.168.2.1 -> 192.168.63.110: icmp: echo request
6.201902 192.168.2.1 -> 192.168.63.110: icmp: echo request
7.225765 192.168.2.1 -> 192.168.63.110: icmp: echo request
but if I ping from A to B, traffic is coming from wan interface (wan2) instead of tunnel interface, no traffic is coming through tunnel interface
forti01 # diagnose sniffer packet tun02
interfaces=[tun02]
filters=[none]
nothing shows at all
forti01 # diagnose sniffer packet wan2 'host 192.168.2.1 and icmp'
interfaces=[wan2]
filters=[host 192.168.2.1 and icmp]
0.370676 192.168.63.110 -> 192.168.2.1: icmp: echo request
1.371506 192.168.63.110 -> 192.168.2.1: icmp: echo request
2.374358 192.168.63.110 -> 192.168.2.1: icmp: echo request
3.376188 192.168.63.110 -> 192.168.2.1: icmp: echo request
I have other production ipsec s2s tunnels connected to other sites which work properly, both way traffic is pass through tunnel interface, not wan.
Also, static route from B to A looks wired as well, the static route gateway IP is 10.0.0.6 (created with firmware 7.0.6), instead of A's public IP address. My other working IPsec tunnels, static route to target sites are all their public IP addresses (created with firmware 7.0.4).
Any hints for further troubleshooting?
