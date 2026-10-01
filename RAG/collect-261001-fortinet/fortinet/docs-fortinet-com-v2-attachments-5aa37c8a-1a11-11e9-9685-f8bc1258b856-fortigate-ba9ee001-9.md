---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-9
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [1373, 1588]
sha256: b0c6745e6eb10f42a5d790eca658e35d2fd06c23f52c297f7935fd72fdf23926
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Example 2: Remote Sites on the Same Subnet IPsec VPN in Transparent Mode
set dstintf "port2"
set srcaddr "10.1.1.0/24"
set dstaddr "Server2"
set action ipsec
set schedule "always"
set service "ANY"
set inbound enable
set outbound enable
set vpntunnel "to_FGT2"
next
edit 3
set srcintf "port1"
set dstintf "port2"
set srcaddr "10.1.1.0/24"
set dstaddr "gateway"
set action ipsec
set schedule "always"
set service "ANY"
set inbound enable
set outbound enable
set vpntunnel "to_FGT2"
next
end
Firewall Policy 3 is not mandatory and is only used to allow PC1 to test a ping
reachability to its default gateway 10.1.1.254.
Configuration of FortiGate 2 (FGT2):
Only relevant parts of configuration are provided.
config system settings
set opmode transparent
set manageip 10.1.1.200/255.255.255.0
end
config router static
edit 1
set gateway 10.1.1.252
next
edit 2
set dst 10.3.3.0 255.255.255.0
set gateway 10.1.1.254
next
end
config system mac-address-table
edit 00:50:56:00:76:03
set interface wan1
next
end
config firewall address
40 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode Example 2: Remote Sites on the Same Subnet
edit "all"
next
edit "PC1"
set subnet 10.1.1.10 255.255.255.255
next
edit "10.1.1.0/24"
set subnet 10.1.1.0 255.255.255.0
next
edit "10.3.3.0/24"
set subnet 10.3.3.0 255.255.255.0
next
end
config vpn ipsec phase1
edit "to_FGT1"
set proposal 3des-sha1 aes128-sha1 des-md5
set remote-gw 10.1.1.100
set psksecret fortinet
next
end
config vpn ipsec phase2
edit "to_FGT1"
set keepalive enable
set phase1name "to_FGT1"
set proposal 3des-sha1 aes128-sha1
set dst-subnet 10.1.1.0 255.255.255.0
next
end
config firewall policy
edit 1
set srcintf "internal"
set dstintf "wan1"
set srcaddr "10.1.1.0/24" set dstaddr "PC1"
set action ipsec
set schedule "always"
set service "ANY" set inbound enable
set outbound enable
set vpntunnel "to_FGT1"
next
edit 2
set srcintf "internal"
set dstintf "wan1"
set srcaddr "10.3.3.0/24" set dstaddr "PC1"
set action ipsec
set schedule "always"
set service "ANY" set inbound enable
set outbound enable
set vpntunnel "to_FGT1"
next
end
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
41

Example 2: Remote Sites on the Same Subnet IPsec VPN in Transparent Mode
Troubleshooting procedure
Check the ARP entries of PC1
C:\ arp -a
Interface: 10.1.1.10 --- 0x20003
Internet Address Physical Address Type
10.1.1.20 00-50-56-00-76-04 dynamic
10.1.1.254 00-09-0f-85-3f-c8 dynamic
MAC address 00-09-0f-85-3f-c8is the FGT3 interface in subnet 10.1.1.0/24.
FDB entries of FGT1
FGT1 (global) # diagnose netlink brctl name host Vdom1.b
show bridge control interface Vdom1.b host. fdb:
size=256, used=6, num=6, depth=1
Bridge Vdom1.b host table
port no device devname mac addr ttl attributes
1 10 port1 00:50:56:00:76:03 0
2 9 port2 00:50:56:00:76:04 44 static
2 9 port2 00:09:0f:85:3f:c8 13
1 10 port1 00:09:0f:88:2f:69  0 Local Static
2 9 port2 00:09:0f:88:2f:68 0 Local Static
2 9 port2 00:09:0f:23:01:d6 0
MAC address 00:09:0f:23:01:d6is “internal” port MAC address of FGT2
00:09:0F:23:01:D6. This is the MAC address used for management in the
Transparent mode VDOM of FGT2, chosen between the lowest MAC address
between wan1 (00:09:0F:78:00:74) and internal (00:09:0F:23:01:D6).
ARP entries of FGT2
FGT2 (TP) # get system arp
Address Age(min) Hardware Addr Interface
10.1.1.20 82 00:50:56:00:76:04 TP.b
10.1.1.100 13 00:09:0f:88:2f:68 TP.b
10.1.1.254 76 00:09:0f:85:3f:c8    TP.b
it is important to have the entry for 10.1.1.254 which is the route to 10.3.3.0/24 .
42 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode Example 2: Remote Sites on the Same Subnet
IPsec Tunnel verification on FGT1
FGT1 (Vdom1) # diagnose vpn tunnel list
list all ipsec tunnel in vd 3
------------------------------------------------------
name=to_FGT2 10.1.1.100:0->10.1.1.200:0 lgwy=dyn tun=tunnel mode=auto bound_if=0
proxyid_num=1 child_num=0 refcnt=10 ilast=0 olast=0
stat: rxp=2754 txp=2945 rxb=308448 txb=176700
dpd: mode=active on=1 idle=5000ms retry=3 count=0 seqno=166 natt:
mode=none draft=0 interval=0 remote_port=0
proxyid=to_FGT2 proto=0 sa=1 ref=2 auto_negotiate=0 serial=1
src: 10.1.1.0/255.255.255.0:0
dst: 0.0.0.0/0.0.0.0:0
SA: ref=3 options=00000009 type=00 soft=0 mtu=1436 expire=1271 replaywin=0 seqno=1e1
life:type=01 bytes=0/0 timeout=1750/1800
dec: spi=3f148cb7 esp=3des key=24 834832201a0dbbf60b0098106f08380538dbd94cacd1ad31
ah=sha1 key=20 b0257a135cba745b956bef3d4b8a6e65934c074b
enc: spi=1895305e esp=3des key=24 4d3092f0b3f84184d4779f85a9953230bf9bc28bd93c0afa
ah=sha1 key=20 0c70acf6ad2193ec5934e2a4332fd09f32016e60
npu_flag=00 npu_rgwy=10.1.1.200 npu_lgwy=10.1.1.100 npu_selid=0
Sniffer trace on FGT1 when PC1 pings all 3 remote destinations
FGT1 (Vdom1) # diagnose sniffer packet any "icmp" 4
interfaces=[any]
filters=[icmp]
0.342268 port1 in 10.1.1.10 -> 10.3.3.30: icmp: echo request
0.342844 port2 in 10.3.3.30 -> 10.1.1.10: icmp: echo reply
0.342884 port1 out 10.3.3.30 -> 10.1.1.10: icmp: echo reply
0.771700 port1 in 10.1.1.10 -> 10.1.1.20: icmp: echo request
0.772504 port2 in 10.1.1.20 -> 10.1.1.10: icmp: echo reply
0.772539 port1 out 10.1.1.20 -> 10.1.1.10: icmp: echo reply
0.907377 port1 in 10.1.1.10 -> 10.1.1.254: icmp: echo request
0.907850 port2 in 10.1.1.254 -> 10.1.1.10: icmp: echo reply
0.907883 port1 out 10.1.1.254 -> 10.1.1.10: icmp: echo reply
Sniffer trace on FGT1 filtered on IPsec protocol
FGT1 (Vdom1) # diagnose sniffer packet port2 "proto 50" 6
interfaces=[port2]
filters=[proto 50]
pcap_lookupnet: port2: no IPv4 address assigned
1.249003 port2 -- 10.1.1.100 -> 10.1.1.200: ip-proto-50 92
0x0000 0009 0f23 01d6 0009 0f88 2f68 0800 4500 ...# ..... /h..E.
0x0010 0070 c9e6 0000 3f32 9a48 0a01 0164 0a01 .p ... ?2.H...d..
0x0020 01c8 1895 305f 0000 01e2 02b6 37b6 8b2c ....0_ ..... 7..,
1.249478 port2 -- 10.1.1.200 -> 10.1.1.100: ip-proto-50 92
0x0000 0009 0f88 2f68 0009 0f23 01d6 0800 4500 ..../h...# ... E.
0x0010 0070 2e31 0000 3f32 35fe 0a01 01c8 0a01 .p.1..?25 ......
0x0020 0164 3f14 8cb8 0000 01e2 324d 66e2 9236 .d? ...... 2Mf..6
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
43

Example 2: Remote Sites on the Same Subnet IPsec VPN in Transparent Mode
From the above trace, the MAC address 0009 0f88 2f68is the MAC address of FGT1
port2 . This is the MAC address used for management in the Transparent mode
VDOM of FGT1, chosen between the lowest MAC address between port1
(00:09:0F:88:2F:69) and port2 ( (00:09:0F:88:2F:68).
Debug flow on FGT1 filtered on Server3
FGT1 (Vdom1) # diagnose debug flow filter addr 10.3.3.30
FGT1 (Vdom1) # diagnose debug flow show console enable
FGT1 (Vdom1) # diagnose debug enable
FGT1 (Vdom1) # diagnose debug flow trace start 10
id=20085 trace_id=11 msg="vd-Vdom1 received a packet(proto=1, 10.1.1.10:512->10.3.3.30:8)
from port1."
id=20085 trace_id=11 msg="Find an existing session, id-00004e85, original direction"
id=20085 trace_id=11 msg="enter IPsec tunnel-to_FGT2 "
id=20085 trace_id=11 msg="encrypted, and send to 10.1.1.200 with source 10.1.1.100"
id=20085 trace_id=11 msg="send out via dev-port2, dst-mac-00:09:0f:23:01:d6 "
id=20085 trace_id=12 msg="vd-Vdom1 received a packet(proto=1, 10.3.3.30:512->10.1.1.10:0)
from port2."
id=20085 trace_id=12 msg="Find an existing session, id-00004e85, reply direction"
id=20085 trace_id=12 msg="send out via dev-port1, dst-mac-00:50:56:00:76:03"
From the trace above, dst-mac-00:09:0f:23:01:d6is “internal” port MAC address of
FGT2 00:09:0F:23:01:D6. This is the MAC address used for management in the
Transparent mode VDOM of FGT2, chosen between the lowest MAC address
between wan1 (00:09:0F:78:00:74) and internal (00:09:0F:23:01:D6).
44 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

