---
id: collect-261001-huawei/huawei/walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255-1
title: "walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255.md
source_anchor: ""
source_lines: [1, 288]
sha256: 9d66765837b75c51c0519bb0ea09fcd8ced4ac5f2f76c48539c03ccefd432581
---

# walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255

CE1
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys CE1
[CE1]undo info en 
Info: Information center is disabled.
[CE1]user-int con 0 
[CE1-ui-console0]idle 0 0 
[CE1-ui-console0]q
[CE1]int g0/0/1
[CE1-GigabitEthernet0/0/1]ip add 192.168.1.1 24
[CE1-GigabitEthernet0/0/1]q
[CE1]int g0/0/0
[CE1-GigabitEthernet0/0/0]ip add 201.201.201.2 24
[CE1-GigabitEthernet0/0/0]q
[CE1]
CE2
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys CE2
[CE2]undo info en 
Info: Information center is disabled.
[CE2]user-int con 0 
[CE2-ui-console0]idle 0 0 
[CE2-ui-console0]q
[CE2]int g0/0/1
[CE2-GigabitEthernet0/0/1]ip add 192.168.2.1 24 
[CE2-GigabitEthernet0/0/1]q
[CE2]int g0/0/0 
[CE2-GigabitEthernet0/0/0]ip add 202.202.202.2 24 
[CE2-GigabitEthernet0/0/0]
CE3
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys CE3
[CE3]undo info en 
Info: Information center is disabled.
[CE3]user-int con 0
[CE3-ui-console0]idle 0 0 
[CE3-ui-console0]int g0/0/1
[CE3-GigabitEthernet0/0/1]ip add 192.168.1.1 24 
[CE3-GigabitEthernet0/0/1]q
[CE3]int g0/0/0
[CE3-GigabitEthernet0/0/0]ip add 203.203.203.2 24
[CE3-GigabitEthernet0/0/0]q
[CE3]
CE4
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys CE4
[CE4]undo info en 
Info: Information center is disabled.
[CE4]user-int con 0 
[CE4-ui-console0]idle 0 0 
[CE4-ui-console0]q
[CE4]int g0/0/0
[CE4-GigabitEthernet0/0/0]ip add 204.204.204.2 24
[CE4-GigabitEthernet0/0/0]q
[CE4]int g0/0/1 
[CE4-GigabitEthernet0/0/1]ip add 192.168.2.1 24 
[CE4-GigabitEthernet0/0/1]q
[CE4]
PE1
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys PE1
[PE1]undo info en 
Info: Information center is disabled.
[PE1]user-int con 0 
[PE1-ui-console0]idle 0 0 
[PE1-ui-console0]q
[PE1]int g0/0/1
[PE1-GigabitEthernet0/0/1]ip add 201.201.201.1 24 
[PE1-GigabitEthernet0/0/1]q
[PE1]int g0/0/0
[PE1-GigabitEthernet0/0/0]ip add 116.64.64.1 24
[PE1-GigabitEthernet0/0/0]q
[PE1]int g2/0/0
[PE1-GigabitEthernet2/0/0]ip add 203.203.203.1 24
[PE1-GigabitEthernet2/0/0]q
[PE1]int loopback 0
[PE1-LoopBack0]ip add 10.0.1.1 24 
[PE1-LoopBack0]
[PE1]
PE2
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys PE2
[PE2]undo info en 
Info: Information center is disabled.
[PE2]user-int con 0
[PE2-ui-console0]idle 0 0
[PE2-ui-console0]q
[PE2]int g0/0/1
[PE2-GigabitEthernet0/0/1]ip add 202.202.202.1 24 
[PE2-GigabitEthernet0/0/1]q
[PE2]int g0/0/0
[PE2-GigabitEthernet0/0/0]ip add 118.16.16.2 24 
[PE2-GigabitEthernet0/0/0]q
[PE2]int g2/0/0
[PE2-GigabitEthernet2/0/0]ip add 204.204.204.1 24 
[PE2-GigabitEthernet2/0/0]q
[PE2]int loopback 0 
[PE2-LoopBack0]ip add 10.0.4.4 24 
[PE2-LoopBack0]
P1
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys p1
[p1]undo info en 
Info: Information center is disabled.
[p1]user-int con 0 
[p1-ui-console0]idle 0 0 
[p1-ui-console0]q
[p1]int g0/0/0 
[p1-GigabitEthernet0/0/0]ip add 116.64.64.2 24 
[p1-GigabitEthernet0/0/0]q
[p1]int g0/0/1
[p1-GigabitEthernet0/0/1]ip add 117.32.32.1 24 
[p1-GigabitEthernet0/0/1]q
[p1]int loopback 0
[p1-LoopBack0]ip add 10.0.2.2 24 
[p1-LoopBack0]q
[p1]
P2
<Huawei>sys
Enter system view, return user view with Ctrl+Z.
[Huawei]sys P2
[P2]undo info en 
Info: Information center is disabled.
[P2]user-int con 0 
[P2-ui-console0]idle 0 0 
[P2-ui-console0]q
[P2]int g0/0/1
[P2-GigabitEthernet0/0/1]ip add 117.32.32.2 24 
[P2-GigabitEthernet0/0/1]q
[P2]int g0/0/0
[P2-GigabitEthernet0/0/0]ip add 118.16.16.1 24 
[P2-GigabitEthernet0/0/0]q
[P2]int loopback 0
[P2-LoopBack0]ip add 10.0.3.3 24 
[P2-LoopBack0]q
[P2]
PE1
[PE1]ospf router-id 10.0.1.1 
[PE1-ospf-1]area 0
[PE1-ospf-1-area-0.0.0.0]network 116.64.64.0 0.0.0.255
[PE1-ospf-1-area-0.0.0.0]network 10.0.1.1 0.0.0.0
[PE1-ospf-1-area-0.0.0.0]q
[PE1-ospf-1]q
[PE1]
PE2
[PE2]ospf router-id 10.0.4.4
[PE2-ospf-1]area 0
[PE2-ospf-1-area-0.0.0.0]network 118.16.16.0 0.0.0.255
[PE2-ospf-1-area-0.0.0.0]network 10.0.4.4 0.0.0.0
[PE2-ospf-1-area-0.0.0.0]q
[PE2-ospf-1]q
[PE2]
P1
[p1]ospf router-id 10.0.2.2 
[p1-ospf-1]area 0
[p1-ospf-1-area-0.0.0.0]network 116.64.64.0 0.0.0.255
[p1-ospf-1-area-0.0.0.0]network 117.32.32.0 0.0.0.255
[p1-ospf-1-area-0.0.0.0]network 10.0.2.2 0.0.0.0
[p1-ospf-1-area-0.0.0.0]q
[p1-ospf-1]q
[p1]
P2
[P2]ospf router-id 10.0.3.3
[P2-ospf-1]area 0
[P2-ospf-1-area-0.0.0.0]network 117.32.32.0 0.0.0.255
[P2-ospf-1-area-0.0.0.0]network 118.16.16.0 0.0.0.255
[P2-ospf-1-area-0.0.0.0]network 10.0.3.3 0.0.0.0 
[P2-ospf-1-area-0.0.0.0]q
[P2-ospf-1]q
[P2]
P1
P2
测试PE1的连通性
PE1与PE2通过Loopback0虚拟接口建立IBGP邻居关系
[PE1]bgp 500
[PE1-bgp]peer 10.0.4.4 as-number 500
[PE1-bgp]peer 10.0.4.4 connect-interface loopback 0
[PE1-bgp][PE2]bgp 500
[PE2-bgp]peer 10.0.1.1 as-number 500
[PE2-bgp]peer 10.0.1.1 conn	
[PE2-bgp]peer 10.0.1.1 connect-interface loopback 0
[PE2-bgp]
在AR1上查看BGP邻居关系
[PE1-bgp]disp bgp peer
 BGP local router ID : 201.201.201.1
 Local AS number : 500
 Total number of peers : 1		  Peers in established state : 1
Peer            V          AS  MsgRcvd  MsgSent  OutQ  Up/Down       State         Pre fRcv
10.0.4.4        4         500        3        5     0  00:01:14      Established          0
在PE1与PE2启用IPv4-Family子族VPNv4地址族，允许PE1与PE2之间交换VPNv4路由信息
[PE1-bgp]ipv4	
[PE1-bgp]ipv4-family vpnv4
[PE1-bgp-af-vpnv4]peer 10.0.4.4 en
[PE1-bgp-af-vpnv4]peer 10.0.4.4 advertise-community 
[PE1-bgp-af-vpnv4]q
[PE1-bgp]q
[PE1][PE2-bgp]ipv4-family vpnv4
[PE2-bgp-af-vpnv4]peer 10.0.1.1 en
[PE2-bgp-af-vpnv4]peer 10.0.1.1 advertise-community 
[PE2-bgp-af-vpnv4]q
[PE2-bgp]q
[PE2]
PE1
[PE1]mpls lsr-id 10.0.1.1 
[PE1]mpls
Info: Mpls starting, please wait... OK!
[PE1-mpls]mpls ldp
[PE1-mpls-ldp]q
[PE1]int g0/0/0
[PE1-GigabitEthernet0/0/0]mpls
[PE1-GigabitEthernet0/0/0]mpls ldp
[PE1-GigabitEthernet0/0/0]q
[PE1]
PE2
[PE2]mpls lsr-id 10.0.4.4 
[PE2]mpls
Info: Mpls starting, please wait... OK!
[PE2-mpls]mpls ldp
[PE2-mpls-ldp]q
[PE2]int g0/0/0
[PE2-GigabitEthernet0/0/0]mpls
[PE2-GigabitEthernet0/0/0]mpls ldp
[PE2-GigabitEthernet0/0/0]q
[PE2]
P1
[p1]mpls lsr-id 10.0.2.2
[p1]mpls
Info: Mpls starting, please wait... OK!
[p1-mpls]mpls ldp
[p1-mpls-ldp]q
[p1]int g0/0/0
[p1-GigabitEthernet0/0/0]mpls
[p1-GigabitEthernet0/0/0]mpls ldp
[p1-GigabitEthernet0/0/0]q
[p1]int g0/0/1
[p1-GigabitEthernet0/0/1]mpls
[p1-GigabitEthernet0/0/1]mpls ldp
[p1-GigabitEthernet0/0/1]q
[p1]
P2
[P2]mpls lsr-id 10.0.3.3
[P2]mpls 
Info: Mpls starting, please wait... OK!
[P2-mpls]mpls ldp 
[P2-mpls-ldp]q
[P2]int g0/0/0
[P2-GigabitEthernet0/0/0]mpls
[P2-GigabitEthernet0/0/0]mpls ldp
[P2-GigabitEthernet0/0/0]q
[P2]int g0/0/1
[P2-GigabitEthernet0/0/1]mpls
[P2-GigabitEthernet0/0/1]mpls ldp 
[P2-GigabitEthernet0/0/1]q
[P2]
P1
P2
在PE1查看MPLS标签转发表
在PE2查看MPLS标签转发表
在P1查看完整标签交换路径LSP表
[PE1]ip vpn-instance vpn_company_A
[PE1-vpn-instance-vpn_company_A]ipv4-family
[PE1-vpn-instance-vpn_company_A-af-ipv4]route-distinguisher 100:1
[PE1-vpn-instance-vpn_company_A-af-ipv4]vpn-target 20:1 export-extcommunity 
 EVT Assignment result: 
Info: VPN-Target assignment is successful.	
[PE1-vpn-instance-vpn_company_A-af-ipv4]vpn-target 20:1 import-extcommunity 
 IVT Assignment result: 
Info: VPN-Target assignment is successful.
[PE1-vpn-instance-vpn_company_A-af-ipv4]q
[PE1-vpn-instance-vpn_company_A]q
[PE1]int g0/0/1
[PE1-GigabitEthernet0/0/1]ip binding vpn-instance vpn_company_A
Info: All IPv4 related configurations on this interface are removed!
Info: All IPv6 related configurations on this interface are removed!
[PE1-GigabitEthernet0/0/1]ip add 201.201.201.1 24
[PE1-GigabitEthernet0/0/1]q
[PE1]ping -vpn-instance vpn_company_A 201.201.201.2
  PING 201.201.201.2: 56  data bytes, press CTRL_C to break
    Reply from 201.201.201.2: bytes=56 Sequence=1 ttl=255 time=60 ms
    Reply from 201.201.201.2: bytes=56 Sequence=2 ttl=255 time=20 ms
    Reply from 201.201.201.2: bytes=56 Sequence=3 ttl=255 time=20 ms
