---
id: collect-261001-huawei/huawei/walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255-2
title: "walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255.md
source_anchor: ""
source_lines: [289, 446]
sha256: a67faceffb82b7c5c021c91ccb7eb73730037a0c5a98a61e7feb09b2c9e52cab
---

# walson171-huawei-ensp-attack-defend-blob-head-mpls-bgp-mpls-bgp-md-1653b255

    Reply from 201.201.201.2: bytes=56 Sequence=4 ttl=255 time=30 ms
    Reply from 201.201.201.2: bytes=56 Sequence=5 ttl=255 time=20 ms
  --- 201.201.201.2 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 20/30/60 ms[PE1]ip vpn-instance vpn_company_B
[PE1-vpn-instance-vpn_company_B]ipv4-family
[PE1-vpn-instance-vpn_company_B-af-ipv4]route-distinguisher 300:1
[PE1-vpn-instance-vpn_company_B-af-ipv4]vpn-target 20:2 both
 IVT Assignment result: 
Info: VPN-Target assignment is successful.
 EVT Assignment result: 
Info: VPN-Target assignment is successful.
[PE1-vpn-instance-vpn_company_B-af-ipv4]q
[PE1-vpn-instance-vpn_company_B]q
[PE1]int g2/0/0
[PE1-GigabitEthernet2/0/0]ip binding vpn-instance vpn_company_B
Info: All IPv4 related configurations on this interface are removed!
Info: All IPv6 related configurations on this interface are removed!
[PE1-GigabitEthernet2/0/0]ip add 203.203.203.1 24 
[PE1-GigabitEthernet2/0/0]q
[PE1][PE2]ip vpn-instance vpn_company_A
[PE2-vpn-instance-vpn_company_A]ipv4-family	
[PE2-vpn-instance-vpn_company_A-af-ipv4]route-distinguisher 200:1
[PE2-vpn-instance-vpn_company_A-af-ipv4]vpn-target 20:1 both
 IVT Assignment result: 
Info: VPN-Target assignment is successful.
 EVT Assignment result: 
Info: VPN-Target assignment is successful.
[PE2-vpn-instance-vpn_company_A-af-ipv4]q
[PE2-vpn-instance-vpn_company_A]q
[PE2]int g0/0/1
[PE2-GigabitEthernet0/0/1]ip binding vpn-instance vpn_company_A
Info: All IPv4 related configurations on this interface are removed!
Info: All IPv6 related configurations on this interface are removed!
[PE2-GigabitEthernet0/0/1]ip add 202.202.202.1 24
[PE2-GigabitEthernet0/0/1]q
[PE2][PE2]ip vpn-instance vpn_company_B
[PE2-vpn-instance-vpn_company_B]ipv4-family 
[PE2-vpn-instance-vpn_company_B-af-ipv4]route-distinguisher 400:1
[PE2-vpn-instance-vpn_company_B-af-ipv4]vpn-target 20:2 both
 IVT Assignment result: 
Info: VPN-Target assignment is successful.
 EVT Assignment result: 
Info: VPN-Target assignment is successful.
[PE2-vpn-instance-vpn_company_B-af-ipv4]q
[PE2-vpn-instance-vpn_company_B]q
[PE2]int g2/0/0	
[PE2-GigabitEthernet2/0/0]ip binding vpn-instance vpn_company_B
Info: All IPv4 related configurations on this interface are removed!
Info: All IPv6 related configurations on this interface are removed!
[PE2-GigabitEthernet2/0/0]ip add 204.204.204.1 24
[PE2-GigabitEthernet2/0/0]q
[PE2]
CE1与PE1建立EBGP邻居关系
[CE1]bgp 100
[CE1-bgp]peer 201.201.201.1 as-number 500
[CE1-bgp]network 192.168.1.0
[PE1]bgp 500
[PE1-bgp]ipv4-family vpn-ins	
[PE1-bgp]ipv4-family vpn-instance vpn_company_A
[PE1-bgp-vpn_company_A]peer 201.201.201.2 as-number 100
[PE1-bgp-vpn_company_A][CE1-bgp]disp bgp peer
 BGP local router ID : 192.168.1.1
 Local AS number : 100
 Total number of peers : 1		  Peers in established state : 1
Peer            V          AS  MsgRcvd  MsgSent  OutQ  Up/Down       State          Pre fRcv
201.201.201.1   4          500       2        5     0  00:00:07      Established    0
[CE1-bgp]
[PE1]disp bgp peer
 BGP local router ID : 201.201.201.1
 Local AS number : 500
 Total number of peers : 1		  Peers in established state : 1
Peer            V          AS  MsgRcvd  MsgSent  OutQ  Up/Down       State          Pre fRcv
10.0.4.4        4          500      47       49     0   00:43:15     Established    0
[PE1]
[PE1]disp bgp vpnv4 vpn-instance vpn_company_A peer
 BGP local router ID : 201.201.201.1
 Local AS number : 500
 VPN-Instance vpn_company_A, Router ID 201.201.201.1:
 Total number of peers : 1		  Peers in established state : 1
Peer            V          AS  MsgRcvd  MsgSent  OutQ  Up/Down       State         Pre fRcv
201.201.201.2   4         100        4        3     0  00:01:56      Established          1
[PE1]
CE2与PE2建立EBGP邻居关系
[CE2]bgp 200
[CE2-bgp]peer 202.202.202.1 as-number 500
[CE2-bgp]network 192.168.2.0
[CE2-bgp][PE2]bgp 500
[PE2-bgp]ipv4-family vpn-instance vpn_company_A
[PE2-bgp-vpn_company_A]peer 202.202.202.2 as-number 200
[PE2-bgp-vpn_company_A][PE1]disp bgp vpnv4 vpn-instance vpn_company_A routing-table 
 BGP Local router ID is 201.201.201.1 
 Status codes: * - valid, > - best, d - damped,
               h - history,  i - internal, s - suppressed, S - Stale
               Origin : i - IGP, e - EGP, ? - incomplete
 VPN-Instance vpn_company_A, Router ID 201.201.201.1:
 Total Number of Routes: 2
      Network            NextHop        MED        LocPrf    PrefVal Path/Ogn
 *>   192.168.1.0        201.201.201.2   0                     0      100i
 *>i  192.168.2.0        10.0.4.4        0          100        0      200i
[PE1]
[PE2]disp bgp vpnv4 vpn-instance vpn_company_A routing-table 
 BGP Local router ID is 202.202.202.1 
 Status codes: * - valid, > - best, d - damped,
               h - history,  i - internal, s - suppressed, S - Stale
               Origin : i - IGP, e - EGP, ? - incomplete
 VPN-Instance vpn_company_A, Router ID 202.202.202.1:
 Total Number of Routes: 2
      Network            NextHop        MED        LocPrf    PrefVal Path/Ogn
 *>i  192.168.1.0        10.0.1.1        0          100        0      100i
 *>   192.168.2.0        202.202.202.2   0                     0      200i
[PE2]
在主机1命令行输入ping 192.168.2.10 可以连通服务器1
CE3与PE1建立EBGP
[CE3]bgp 300
[CE3-bgp]peer 203.203.203.1 as-number 500
[CE3-bgp]network 192.168.1.0 [PE1]bgp 500
[PE1-bgp]ipv4-family vpn-instance vpn_company_B
[PE1-bgp-vpn_company_B]peer 203.203.203.2 as-number 300
[PE1-bgp-vpn_company_B]
CE4与PE2建立EBGP邻居关系
[CE4]bgp 400
[CE4-bgp]peer 204.204.204.1 as-number 500
[CE4-bgp]network 192.168.2.0
[CE4-bgp][PE2]bgp 500
[PE2-bgp]ipv4-family vpn-instance vpn_company_B
[PE2-bgp-vpn_company_B]peer 204.204.204.2 as-number 400
[PE2-bgp-vpn_company_B]
在PE1和PE2查看vpn_company_B实例的VPNv4路由表
[PE1]disp bgp vpnv4 vpn-instance vpn_company_B routing-table
 BGP Local router ID is 116.64.64.1 
 Status codes: * - valid, > - best, d - damped,
               h - history,  i - internal, s - suppressed, S - Stale
               Origin : i - IGP, e - EGP, ? - incomplete
 VPN-Instance vpn_company_B, Router ID 116.64.64.1:
 Total Number of Routes: 2
      Network            NextHop        MED        LocPrf    PrefVal Path/Ogn
 *>   192.168.1.0        203.203.203.2   0                     0      300i
 *>i  192.168.2.0        10.0.4.4        0          100        0      400i
[PE1][PE2]disp bgp vpnv4 vpn-instance vpn_company_B routing-table
 BGP Local router ID is 118.16.16.2 
 Status codes: * - valid, > - best, d - damped,
               h - history,  i - internal, s - suppressed, S - Stale
               Origin : i - IGP, e - EGP, ? - incomplete
 VPN-Instance vpn_company_B, Router ID 118.16.16.2:
 Total Number of Routes: 2
      Network            NextHop        MED        LocPrf    PrefVal Path/Ogn
 *>i  192.168.1.0        10.0.1.1        0          100        0      300i
 *>   192.168.2.0        204.204.204.2   0                     0      400i
[PE2]
在主机2命令行输入命令ping 192.168.2.20,可以连通服务器2
在主机2命令行输入命令ping 192.168.2.10，不可以连通服务器1
1.可在服务器上部署IIS(Internet Information Services,互联网信息服务)时选择https协议，通过SSL证书对Web站点数据加密
2.MPLS属于2.5层VPN，不能使用IPSec(第三层协议)保护MPLS报文安全性,即不可能有MPLS Over IPSec技术。
3.路由器只能保证数据包的完整性和来源的可靠性,不对其内部具体数据负责,也不负责弥补非路由器转发导致的安全漏洞
1.对于所有企业而言，一般认为数据包在自身内网中传输是安全的，而在公网中传输是不安全的。MPLS优点是转发效率高，低延迟，但缺少安全协议，难以保证MPLS报文在公网中传输的安全性
