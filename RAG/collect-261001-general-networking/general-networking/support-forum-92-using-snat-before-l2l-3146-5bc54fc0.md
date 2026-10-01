---
id: collect-261001-general-networking/general-networking/support-forum-92-using-snat-before-l2l-3146-5bc54fc0
title: "support-forum-92-using-snat-before-l2l-3146-5bc54fc0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-using-snat-before-l2l-3146-5bc54fc0.md
source_anchor: ""
source_lines: [1, 32]
sha256: f31089043e17de3a4969c4d3aa50035d35b23311c2056db12818da1780541c1b
---

# support-forum-92-using-snat-before-l2l-3146-5bc54fc0

Using SNAT before L2L
Hello Fortinet Forum!
I've ran into an issue with NAT and usually it just works. I hope you can help me out.
I have two Fortigates (A and B) with site-to-site VPN between them.
Fortigate A needs to be NATed from 192.168.100.0/24 to 10.10.100.253. Therefore I've created a IP Pool as follows;
Type: Overload
External IP Range: 10.10.100.253 - 10.10.100.253
ARP reply: tick
Fortigate B have LAN 192.168.254.0/24
My policy is as follows;
Incoming Interface: Interface mastering 192.168.100.0/24
Outgoing Interface: site-to-site interface
Source: 192.168.100.0/24
Destination: 192.168.254.0/24
Service: ALL
Action: Accept
NAT: Enabled
IP Pool Configuration: Use Dynamic IP Pool (The overload mentioned above)
The site-to-site phase 2 have the following
Local Address: 10.10.100.0/24 (The NATed address)
Remote Address: 192.168.254.0/24
The site-to-site is state UP.
I run the following diagnose command to troubleshoot;
diagnose debug enable
diagnose debug flow filter addr 192.168.100.1
diagnose debug flow trace start 100
The outcome of those diagnose commands is;
id=20085 trace_id=246 func=print_pkt_detail line=5519 msg="vd-root:0 received a packet(proto=1, 192.168.100.1:14848->192.168.254.1:2048) from local. type=8, code=0, id=14848, seq=0." id=20085 trace_id=246 func=init_ip_session_common line=5684 msg="allocate a new session-01ce3079" id=20085 trace_id=246 func=ipsecdev_hard_start_xmit line=692 msg="enter IPsec interface-L2L" id=20085 trace_id=246 func=ipsec_common_output4 line=804 msg="No matching IPsec selector, drop"
I know the "No matching IPsec selector, drop" means that it doesn't match in my IPsec phase 2 but it would if the NAT worked as I was told, it would match.
As far as I know the network packet should match the source and destination in the policy. Then be NATed before entering the IPsec tunnel but in the diagnose output it doesn't look like it's being NATed.
received a packet(proto=1, 192.168.100.1:14848->192.168.254.1:2048
Can someone give me a hint?
