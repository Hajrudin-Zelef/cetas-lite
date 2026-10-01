---
id: collect-261001-general-networking/general-networking/support-forum-92-ip-pool-nat-one-to-one-not-working-186993-e1bbb010
title: "support-forum-92-ip-pool-nat-one-to-one-not-working-186993-e1bbb010"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-ip-pool-nat-one-to-one-not-working-186993-e1bbb010.md
source_anchor: ""
source_lines: [1, 54]
sha256: 7ce3db6aa04d3a157dfc31a6da60dce4351681d61ddedc0602517cffa56d234c
---

# support-forum-92-ip-pool-nat-one-to-one-not-working-186993-e1bbb010

IP-Pool-Nat one to one not working
Hello everyone,
I’m currently experiencing some issues with our Site-to-Site VPN (fortiOS 7.0.12) that was previously functioning without any problems. It seems that the NAT IP pool is not properly translating the source address, which is causing issues during the Phase 2 negotiation.
As a result, the remote site is unable to establish a proper connection to exit the tunnel. I suspect that this misconfiguration might be affecting the traffic routing and connectivity.
If anyone has encountered a similar issue or has suggestions on how to troubleshoot this, I would greatly appreciate your input!
_______CONFIG SNIPPET_________
edit "H_IPSEC_192.168.110.11"
set uuid xxxxxxxxxx
set subnet 192.168.110.12 255.255.255.255
next
edit "IPSEC-192.168.110.12"
set phase1name "VPN-IPSEC"
set proposal aes256-md5
set dhgrp 5
set keylifeseconds 3600
set src-subnet 10.0.11.6 255.255.255.255
set dst-subnet 192.168.110.12 255.255.255.255
next
config firewall ippool
edit "IP-POOL-NAT"
set startip 10.0.11.0
set endip 10.0.11.254
next
end
edit 17
set name "To VPN-IPSEC"
set uuid xxxxxxx
set srcintf "port2"
set dstintf "VPN-IPSEC"
set action accept
set srcaddr "H_10.0.1.6"
set dstaddr  "H_192.168.110.12"
set schedule "always"
set service "ALL"
set utm-status enable
set nat enable
set ippool enable
set
___________________________
FGTAZ-VM01 # diagnose debug reset
FGTAZ-VM01 # diagnose debug flow filter clear
FGTAZ-VM01 # diagnose debug flow filter addr 192.168.110.11
FGTAZ-VM01 # diagnose debug flow show function-name enable
show function name
FGTAZ-VM01 # diagnose debug flow trace start 100
FGTAZ-VM01 # diagnose debug enable
FGTAZ-VM01 # id=20085 trace_id=2 func=print_pkt_detail line=5844 msg="vd-root:0 received a packet(proto=1, 10.0.1.6:7390->192.168.110.11:2048) tun_id=0.0.0.0 from port2. type=8, code=0, id=7390, seq=814."
id=20085 trace_id=2 func=resolve_ip_tuple_fast line=5930 msg="Find an existing session, id-000000ed, original direction"
id=20085 trace_id=2 func=ipv4_fast_cb line=53 msg="enter fast path"
id=20085 trace_id=2 func=ip_session_run_all_tuple line=7156 msg="SNAT 10.0.1.6->10.0.11.17:7390"
id=20085 trace_id=2 func=ipsecdev_hard_start_xmit line=669 msg="enter IPSec interface VPN-IPSEC, tun_id=0.0.0.0"
id=20085 trace_id=2 func=_do_ipsecdev_hard_start_xmit line=229 msg="output to IPSec tunnel VPN-IPSEC"
id=20085 trace_id=2 func=ipsec_common_output4 line=778 msg="No matching IPsec selector, drop"
Thank you in advance for your help!
