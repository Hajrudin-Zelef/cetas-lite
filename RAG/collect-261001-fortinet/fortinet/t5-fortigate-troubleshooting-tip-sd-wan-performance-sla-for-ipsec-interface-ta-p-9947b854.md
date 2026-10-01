---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-sd-wan-performance-sla-for-ipsec-interface-ta-p-9947b854
title: "t5-fortigate-troubleshooting-tip-sd-wan-performance-sla-for-ipsec-interface-ta-p-9947b854"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-sd-wan-performance-sla-for-ipsec-interface-ta-p-9947b854.md
source_anchor: ""
source_lines: [1, 214]
sha256: a5fd877b393edd8481f9cdd7d650e805cc6fdf7be56bfe2d40e189f440a8c06b
---

# t5-fortigate-troubleshooting-tip-sd-wan-performance-sla-for-ipsec-interface-ta-p-9947b854

**Description**


This article describes how to fix an issue where the SD-WAN performance SLA shows as 'down' by adding the VPN interface to the SD-WAN and how to configure the SD-WAN performance SLA for the VPN interface.


**Scope**


Instructions differ between FortiOS v6.4 and above and versions before FortiOS v6.4.**Solution**


**After FortiOS v6.4.**


In FortiOS v6.4 and above, there is a wizard to facilitate easy VPN creation under the SD-WAN zone. See the following article:

Technical Tip: Configure IPsec VPN with SD-WAN


**Before FortiOS v6.4.**


Adding the VPN interfaces to the SD-WAN is necessary to load balance the remote network Traffic.

After adding the VPN interface to the SD-WAN, when the performance SLA is created for the VPN interface, the performance SLA status shows as 'down'. As a result, if 'set update-static-route' is enabled, the associated route becomes inactive.


However, the VPN tunnel status shows as 'up' in the IPsec monitor:

**diagnose vpn tunnel list**

name=HO_port1_2 ver=1 serial=1 10.40.16.57:0->10.40.16.20:0

bound_if=3 lgwy=static/1 tun=intf/0 mode=auto/1 encap=none/0

proxyid_num=1 child_num=0 refcnt=20 ilast=0 olast=0 ad=/0

stat: rxp=352 txp=453 rxb=42240 txb=27180

dpd: mode=on-demand on=1 idle=20000ms retry=3 count=0 seqno=5

natt: mode=none draft=0 interval=0 remote_port=0

proxyid=HO_port1_2 proto=0 sa=1 ref=2 serial=2

**src: 0:172.31.192.0/255.255.240.0:0**

dst: 0:172.31.128.0/255.255.240.0:0

SA: ref=3 options=10226 type=00 soft=0 mtu=1438 expire=42225/0B replaywin=2048

seqno=1c6 esn=0 replaywin_lastseq=00000161 itn=0

life: type=01 bytes=0/0 timeout=42901/43200

dec: spi=e04e9478 esp=aes key=16 73cb7c9658e2b3bf3f9eddbf3e966514

ah=sha1 key=20 1c990b12c0ef9df9721f550924cc69e2ceadcbcc

enc: spi=276713b0 esp=aes key=16 83199cc34061a388045e59a157827281

ah=sha1 key=20 6ebe8070d5f7105ec4a70da6eb4e1e467a701463

dec:pkts/bytes=352/21120, enc:pkts/bytes=453/54360

------------------------------------------------------

name=HO_port2_3 ver=1 serial=2 10.40.48.57:0->10.40.48.20:0

bound_if=4 lgwy=static/1 tun=intf/0 mode=auto/1 encap=none/0

proxyid_num=1 child_num=0 refcnt=17 ilast=0 olast=0 ad=/0

stat: rxp=343 txp=443 rxb=41160 txb=26580

dpd: mode=on-demand on=1 idle=20000ms retry=3 count=0 seqno=4

natt: mode=none draft=0 interval=0 remote_port=0

proxyid=HO_port2_3 proto=0 sa=1 ref=2 serial=2

src: 0:172.31.192.0/255.255.240.0:0

dst: 0:172.31.128.0/255.255.240.0:0

SA: ref=3 options=10226 type=00 soft=0 mtu=1438 expire=42224/0B replaywin=2048

seqno=1bc esn=0 replaywin_lastseq=00000158 itn=0

life: type=01 bytes=0/0 timeout=42897/43200

dec: spi=e04e9479 esp=aes key=16 02b54e277607488456b277146f4d1429

ah=sha1 key=20 eaf940619730b89e85224178d7f2de33c3550e9e

enc: spi=276713b1 esp=aes key=16 9007a2942009fae3673649b9d87efe60

ah=sha1 key=20 36d1fa1945112b4baa003b6610d562608e093f07

dec:pkts/bytes=343/20580, enc:pkts/bytes=443/53160

**diagnose sniffer packet any "host 172.31.128.20 and icmp" 4**

interfaces=[any]

filters=[host 172.31.128.20 and icmp]

0.630359 HO_port1_2 out 10.40.16.57 -> 172.31.128.20: icmp: echo request

0.630388 HO_port2_3 out 10.40.16.57 -> 172.31.128.20: icmp: echo request

1.640399 HO_port1_2 out 10.40.16.57 -> 172.31.128.20: icmp: echo request

1.640444 HO_port2_3 out 10.40.16.57 -> 172.31.128.20: icmp: echo request

id=20085 trace_id=3 func=print_pkt_detail line=5517 msg="vd-root:0 received a packet(proto=1, 10.40.16.57:257->172.31.128.20:2048) from local. type=8, code=0, id=257, seq=50."

id=20085 trace_id=3 func=resolve_ip_tuple_fast line=5597 msg="Find an existing session, id-00001308, original direction"

id=20085 trace_id=3 func=ipsecdev_hard_start_xmit line=692 msg="enter IPsec interface-HO_port1_2"

id=20085 trace_id=3 func=ipsec_common_output4 line=804 msg=

**"No matching IPsec selector, drop"**<-----

id=20085 trace_id=4 func=print_pkt_detail line=5517 msg="vd-root:0 received a packet(proto=1, 10.40.16.57:258->172.31.128.20:2048) from local. type=8, code=0, id=258, seq=33."

id=20085 trace_id=4 func=resolve_ip_tuple_fast line=5597 msg="Find an existing session, id-0000130a, original direction"

id=20085 trace_id=4 func=ipsecdev_hard_start_xmit line=692 msg="enter IPsec interface-HO_port2_3"

**id=20085 trace_id=4 func=ipsec_common_output4 line=804 msg="No matching IPsec selector, drop" <-----**

**Solution.**

**After FortiOS v6.4.**

A simple ping test with a source defined as an IPsec VPN tunnel IP should make it clear whether the destination server is reachable through that interface.


**execute ping-options source 10.10.10.2**

**execute ping 192.168.1.20**

PING 192.168.1.20 (192.168.1.20): 56 data bytes

64 bytes from 192.168.1.20: icmp_seq=0 ttl=255 time=2.2 ms

64 bytes from 192.168.1.20: icmp_seq=1 ttl=255 time=0.9 ms

64 bytes from 192.168.1.20: icmp_seq=2 ttl=255 time=0.3 ms

64 bytes from 192.168.1.20: icmp_seq=3 ttl=255 time=0.5 ms

64 bytes from 192.168.1.20: icmp_seq=4 ttl=255 time=0.7 ms

--- 192.168.1.20 ping statistics ---

5 packets transmitted, 5 packets received, 0% packet loss

round-trip min/avg/max = 0.3/0.9/2.2 ms


**Note**: If there is a Performance SLA enabled where the Detect Server is a public IP, ensure that only the WAN interfaces are specified in the Participants and not select 'All SDWAN members'. This is because the public detects that server-bound probes will fail via VPN tunnels since there is no route to the Internet through them. Since the Performance SLA will fail, the VPN tunnels will go down.

 **Before FortiOS v6.4:**

To solve this issue, configure a source IP for the VPN interface in SD-WAN settings.

Add the FortiGate local interface IP as a source IP for the VPN in SD-WAN, and make sure that it is part of the phase2 selectors.

**config system virtual-wan-link
config members**

config members

edit <id>

set source x.x.x.x <----- LAN interface IP.

next

end

**config system virtual-wan-link**

**config members**

1 port1

2 HO_port1_2

3 HO_port2_3

4 port2(members)

**config system sdwan**

**config health-check**

**diagnose sniffer packet any "host 172.31.128.20 and icmp" 4**

interfaces=[any]

filters=[host 172.31.128.20 and icmp]

1.679274 HO_port1_2 out 172.31.192.57 -> 172.31.128.20: icmp: echo request

1.679735 HO_port1_2 in 172.31.128.20 -> 172.31.192.57: icmp: echo reply

1.799248 HO_port2_3 out 172.31.192.57 -> 172.31.128.20: icmp: echo request

1.799707 HO_port2_3 in 172.31.128.20 -> 172.31.192.57: icmp: echo reply
