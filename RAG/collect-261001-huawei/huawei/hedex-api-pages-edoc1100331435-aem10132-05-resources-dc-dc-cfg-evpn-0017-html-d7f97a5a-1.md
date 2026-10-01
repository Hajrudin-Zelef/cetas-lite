---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-evpn-0017-html-d7f97a5a-1
title: "Configure Router1. The configurations of Router2 and Router3 are similar to that of Router1, and are not mentioned here. When OSPF is used, the 32-bit loopback address of each router must be advertised."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-evpn-0017-html-d7f97a5a.md
source_anchor: ""
source_lines: [1, 33]
sha256: 6cd955db647e20dd82296356cd84c074d0515c060d9f21cf68dd86a3b1568c87
---

# Configure Router1. The configurations of Router2 and Router3 are similar to that of Router1, and are not mentioned here. When OSPF is used, the 32-bit loopback address of each router must be advertised.

In Figure 1, Router1 and Router2 are the branch and headquarters gateways of an enterprise. As users in the headquarters and branch have different service requirements, they are planned in different network segments. PC_1 in the branch and PC_2 in the headquarters belong to VLAN 10 and VLAN 20, respectively. The enterprise requires that users in the headquarters and branch can communicate over a VXLAN tunnel dynamically established using BGP EVPN.
On AR routers, the ingress replication list cannot be dynamically discovered using BGP EVPN. To implement VXLAN using BGP EVPN on AR routers, you can only run the vni vni-id head-end peer-list ip-address &<1-10> command in the NVE interface view to configure a VXLAN network identifier (VNI) ingress replication list.
The configuration roadmap is as follows:
# Configure Router1. The configurations of Router2 and Router3 are similar to that of Router1, and are not mentioned here. When OSPF is used, the 32-bit loopback address of each router must be advertised.
<Huawei> system-view
[Huawei] sysname Router1
[Router1] interface loopback 1
[Router1-LoopBack1] ip address 10.1.1.2 32
[Router1-LoopBack1] quit
[Router1] interface ethernet 2/0/0
[Router1-Ethernet2/0/0] ip address 192.168.2.1 24
[Router1-Ethernet2/0/0] quit
[Router1] ospf 1 router-id 10.1.1.2
[Router1-ospf-1] area 0
[Router1-ospf-1-area-0.0.0.0] network 10.1.1.2 0.0.0.0
[Router1-ospf-1-area-0.0.0.0] network 192.168.2.0 0.0.0.255
[Router1-ospf-1-area-0.0.0.0] quit
[Router1-ospf-1] quit
# After OSPF is configured, the routers can learn the loopback interface address of each other and successfully ping each other. The following shows the ping result from Router1 to Router2.
[Router1] ping 10.2.2.2
  PING 10.2.2.2: 56  data bytes, press CTRL_C to break                     
    Reply from 10.2.2.2: bytes=56 Sequence=1 ttl=255 time=1 ms             
    Reply from 10.2.2.2: bytes=56 Sequence=2 ttl=255 time=5 ms             
    Reply from 10.2.2.2: bytes=56 Sequence=3 ttl=255 time=5 ms             
    Reply from 10.2.2.2: bytes=56 Sequence=4 ttl=255 time=2 ms             
    Reply from 10.2.2.2: bytes=56 Sequence=5 ttl=255 time=2 ms             
                                                                                
  --- 10.2.2.2 ping statistics ---                                         
    5 packet(s) transmitted                                                     
    5 packet(s) received                                                        
    0.00% packet loss                                                           
    round-trip min/avg/max = 1/3/5 ms                                           
                                                                                
