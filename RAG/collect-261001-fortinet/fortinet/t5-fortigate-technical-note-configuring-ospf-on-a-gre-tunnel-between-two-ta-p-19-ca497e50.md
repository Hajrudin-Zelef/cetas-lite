---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-note-configuring-ospf-on-a-gre-tunnel-between-two-ta-p-19-ca497e50
title: "t5-fortigate-technical-note-configuring-ospf-on-a-gre-tunnel-between-two-ta-p-19-ca497e50"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-note-configuring-ospf-on-a-gre-tunnel-between-two-ta-p-19-ca497e50.md
source_anchor: ""
source_lines: [1, 83]
sha256: 24988761154f20648c3cfe5adb5657dd999176a9006a739c2a2f42fd9e0c630f
---

# t5-fortigate-technical-note-configuring-ospf-on-a-gre-tunnel-between-two-ta-p-19-ca497e50

**Description**


This article describes how to configure OSPF on a GRE tunnel between two FortiGates.

Additional information about OSPF or GRE is available in the related articles at the end of this document or in the FortiGate CLI Reference or Administration guide at http://docs.forticare.com/**Scope**


FortiGate or VDOM in NAT mode.

**Diagram**


(Internet)

**I**

**GRE tunnel**
I

**Expectations, Requirements**

- Establish a GRE tunnel between both FortiGates to be able to reach each remote LAN 10.x.x.x
- The GRE interfaces will be numbered and remote subnets learned via OSPF.
- OSPF will be enabled on all 10.x.x.x/8 interfaces.

**Configuration:**


| CLI configuration of FortiGate 1 config system gre-tunneledit "toFG2" set interface "port1" set local-gw 198.51.100.1 set remote-gw 203.0.113.2 next end config system interface    edit "port1"         set ip 198.51.100.1 255.255.255.0         set alias Internet     next    edit "port2"         set ip 10.1.1.254 255.255.255.0         set alias LAN     next set ip 10.0.0.1 255.255.255.255 set type tunnel set remote-ip 10.0.0.2 set interface "port1"     next end config system gre-tunneledit "toFG2" set interface "port1" set local-gw 198.51.100.1 set remote-gw 203.0.113.2 next end config firewall policyedit 0 set srcintf "port2" set dstintf "toFG2" set srcaddr "all" set dstaddr "all" set action accept set schedule "always" set service "ALL" next      edit 0set srcintf "toFG2" set dstintf "port2" set srcaddr "all" set dstaddr "all" set action accept set schedule "always" set service "ALL" next end  config router ospfconfig area edit 0.0.0.0 next end config network edit 1 set prefix 10.0.0.0 255.0.0.0 next end set router-id 10.1.1.254 end | 


| CLI configuration of FortiGate 2 config system gre-tunnel     edit "toFG1" set interface "port1" set local-gw 203.0.113.2 set remote-gw 198.51.100.1 next end config system interface      edit "port1"         set ip 203.0.113.2 255.255.255.0         set alias Internet     next      edit "port2"         set ip 10.2.2.254 255.255.255.0         set alias LAN     next     edit "toFG1"set ip 10.0.0.2 255.255.255.255 set type tunnel set remote-ip 10.0.0.1 set interface "port1" next end config firewall policy     edit 0set srcintf "port2" set dstintf "toFG1" set srcaddr "all" set dstaddr "all" set action accept set schedule "always" set service "ALL" next      edit 0 set srcintf "toFG1" set dstintf "port2" set srcaddr "all" set dstaddr "all" set action accept set schedule "always" set service "ALL" next end config router ospfconfig area edit 0.0.0.0 next end config network edit 1 set prefix 10.0.0.0 255.0.0.0 next end set router-id 10.2.2.254 end | 

**Verification**

1.  Verify the GRE tunnels:

| vd=0 devname=toFG1 devindex=3 ifindex=22 saddr=203.0.113.2 daddr=198.51.100.1 ref=0 key=0/0 flags=0/0 total tunnel = 1 | 

| if=toFG1 family=00 type=778 index=22 mtu=1476 link=0 master=0 ref=12 state=off start fw_flags=0 flags=**up**  p2p**run** noarp multicast | 

| [...] == [ toFG1 ] name: toFG1 ip: 0.0.0.0 0.0.0.0  **status: up** netbios-forward: disable    type: tunnel   netflow-sampler: disable    sflow-sampler: disable    scan-botnet-connections: disable    explicit-web-proxy: disable    explicit-ftp-proxy: disable    wccp: disable | 


2. Verify if OSPF has established full adjacency between both FortiGates:

FG1 # get router info  ospf  neighbor

| OSPF process 0: Neighbor ID Pri State Dead Time Address Interface 10.2.2.254 1 Full/ - 00:00:36 10.0.0.2 toFG2 | 


FG2 # get router info ospf  neighbor

| OSPF process 0: Neighbor ID Pri State Dead Time Address Interface 10.1.1.254 1 Full/ - 00:00:37 10.0.0.1 toFG1 | 


3. Verify the routing tables to see if the remote subnets are advertised in OSPF:

FG1 # get router info routing-table all

| Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP O - OSPF, IA - OSPF inter area N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2 E1 - OSPF external type 1, E2 - OSPF external type 2 i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area * - candidate default S* 0.0.0.0/0 [10/0] via 198.51.100.254, port1 C 10.0.0.1/32 is directly connected, toFG2 C 10.0.0.2/32 is directly connected, toFG2 C 10.1.1.0/24 is directly connected, port2 O **10.2.2.0/24** [110/200] via 10.0.0.2, toFG2, 00:04:10C 198.51.100.0/24 is directly connected, port1 | 


FG2 # get router info routing-table all

| Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP O - OSPF, IA - OSPF inter area N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2 E1 - OSPF external type 1, E2 - OSPF external type 2 i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area * - candidate default S* 0.0.0.0/0 [10/0] via 203.0.113.254, port1 C 10.0.0.1/32 is directly connected, toFG1 C 10.0.0.2/32 is directly connected, toFG1 C 10.2.2.0/24 is directly connected, port2 **O       10.1.1.0/24 [110/110] via 10.0.0.1, toFG1, 00:04:58**C 203.0.113.0/24 is directly connected, port1 | 

**Troubleshooting**

While starting a ping from PC1 to PC2, take a sniffer trace on either FortiGate to see if the traffic reaches and is forwarded on all interfaces (see also related article about sniffer on GRE interface) :

FG1 # diagnose sniffer packet  any "icmp" 4

| interfaces=[any] filters=[icmp] 2.901412 port2 in 10.1.1.1 -> 10.2.2.2: icmp: echo request 2.901429 toFG2 out 10.1.1.1 -> 10.2.2.2: icmp: echo request 2.901954 toFG2 in 10.2.2.2 -> 10.1.1.1: icmp: echo reply 2.901979 port2 out 10.2.2.2 -> 10.1.1.1: icmp: echo reply | 


FG2 # diagnose sniffer packet  any "icmp" 4

| interfaces=[any] filters=[icmp] 7.241465 toFG1 in 10.1.1.1 -> 10.2.2.2: icmp: echo request 7.241529 internal out 10.1.1.1 -> 10.2.2.2: icmp: echo request 7.241815 internal in 10.2.2.2 -> 10.1.1.1: icmp: echo reply 7.241836 toFG1 out 10.2.2.2 -> 10.1.1.1: icmp: echo reply | 


**Related Articles**
