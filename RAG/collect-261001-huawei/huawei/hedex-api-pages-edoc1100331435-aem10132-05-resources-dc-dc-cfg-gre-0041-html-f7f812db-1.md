---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0041-html-f7f812db-1
title: "Configure PE1. The configurations of PE2 and P are similar to the configuration of PE1, and are not mentioned here."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0041-html-f7f812db.md
source_anchor: ""
source_lines: [1, 83]
sha256: 0d4141eb7d97468be3f81d8c3373a14350347227467b23cda6d1292d7a2a8c12
---

# Configure PE1. The configurations of PE2 and P are similar to the configuration of PE1, and are not mentioned here.

AR611-S, AR611W-S, AR611E-S, AR611-LTE4EA, AR611, AR611W, AR611W-LTE4CN, AR611W-LTE6EA, AR617VW, AR617VW-LTE4, AR617VW-LTE4EA, and AR651C do not support this function.
The RU-5G-101 does not support this function.
An ISP network provides the L2VPN service for users. Many users connect to the MPLS network through PE1 and PE2, and users on the PEs change frequently. A proper VPN solution is required to provide secure VPN services for users and to simplify configuration when new users connect to the network.
A Martini VLL connection can be set up between CE1 and CE2 to meet these requirements. By default, the system uses Label Switched Paths (LSPs) for Martini VLL, and does not perform load balancing. When the P does not provide MPLS functions, VLL cannot be implemented.
To solve the problem, apply a tunnel policy to Martini VLL to specify that VLL services are transmitted over a GRE tunnel.
The configuration roadmap is as follows:
Configure a routing protocol on the PE and P devices on the backbone network to ensure reachability between them.
Enable MPLS and MPLS LDP on PEs. Set up a remote LDP session between the PEs to exchange VC labels between the PEs.
Enable MPLS L2VPN on PEs. Enabling MPLS L2VPN is the prerequisite for VLL configuration.
Create GRE tunnel interfaces on PEs and establish a GRE tunnel between PEs.
Create VC connections on PEs. Because the P does not support MPLS functions, configure a tunnel policy and apply it when you create VC connections so that VLL services can be transmitted over a GRE tunnel.
# Configure PE1. The configurations of PE2 and P are similar to the configuration of PE1, and are not mentioned here.
<Huawei> system-view
[Huawei] sysname PE1
[PE1] interface gigabitethernet 2/0/0
[PE1-GigabitEthernet2/0/0] ip address 172.1.1.1 255.255.255.0
[PE1-GigabitEthernet2/0/0] quit
[PE1] interface loopback 1
[PE1-LoopBack1] ip address 10.10.1.1 255.255.255.255
[PE1-LoopBack1] quit
[PE1] ospf 1
[PE1-ospf-1] area 0
[PE1-ospf-1-area-0.0.0.0] network 172.1.1.0 0.0.0.255
[PE1-ospf-1-area-0.0.0.0] network 10.10.1.1 0.0.0.0
[PE1-ospf-1-area-0.0.0.0] quit
[PE1-ospf-1] quit
After the configurations are complete, OSPF neighbor relationships can be set up between PE1, P, and PE2. Run the display ospf peer command. You can see that the neighbor status is Full. Run the display ip routing-table command. You can see that PEs have learnt the routes to Loopback1 of each other.
# Configure PE1.
[PE1] mpls lsr-id 10.10.1.1
[PE1] mpls
[PE1-mpls] quit
[PE1] mpls ldp
[PE1-mpls-ldp] quit
[PE1] mpls ldp remote-peer 10.10.2.1
[PE1-mpls-ldp-remote-10.10.2.1] remote-ip 10.10.2.1
[PE1-mpls-ldp-remote-10.10.2.1] quit
# Configure PE2.
[PE2] mpls lsr-id 10.10.2.1
[PE2] mpls
[PE2-mpls] quit
[PE2] mpls ldp
[PE2-mpls-ldp] quit
[PE2] mpls ldp remote-peer 10.10.1.1
[PE2-mpls-ldp-remote-10.10.1.1] remote-ip 10.10.1.1
[PE2-mpls-ldp-remote-10.10.1.1] quit
After the configurations are complete, run the display mpls ldp session command on PE1 to view the LDP session status. You can see that an LDP session is set up between PE1 and PE2.
The display on PE1 is used as an example.
[PE1] display mpls ldp session
                                                                                
 LDP Session(s) in Public Network                                               
 Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)                  
 A '*' before a session means the session is being deleted.                     
 ------------------------------------------------------------------------------ 
 PeerID             Status      LAM  SsnRole  SsnAge      KASent/Rcv            
 ------------------------------------------------------------------------------ 
 10.10.2.1:0          Operational DU   Passive  0000:00:01  1/1                   
 ------------------------------------------------------------------------------ 
 TOTAL: 1 session(s) Found.                                                     
[PE1] mpls l2vpn
[PE1-l2vpn] quit
[PE2] mpls l2vpn
[PE2-l2vpn] quit
[PE1] interface tunnel 0/0/1
[PE1-Tunnel0/0/1] ip address 10.2.1.1 255.255.255.0 
[PE1-Tunnel0/0/1] tunnel-protocol gre
[PE1-Tunnel0/0/1] source 10.10.1.1
[PE1-Tunnel0/0/1] destination 10.10.2.1
[PE1-Tunnel0/0/1] quit
[PE2] interface tunnel 0/0/1
[PE2-Tunnel0/0/1] ip address 10.2.1.2 255.255.255.0 
[PE2-Tunnel0/0/1] tunnel-protocol gre
[PE2-Tunnel0/0/1] source 10.10.2.1
[PE2-Tunnel0/0/1] destination 10.10.1.1
[PE2-Tunnel0/0/1] quit
After the configurations are complete, the tunnel interfaces go Up and can ping each other.
[PE1] ping -a 10.2.1.1 10.2.1.2
  PING 10.2.1.2: 56  data bytes, press CTRL_C to break        
    Reply from 10.2.1.2: bytes=56 Sequence=1 ttl=255 time=1 ms
    Reply from 10.2.1.2: bytes=56 Sequence=2 ttl=255 time=1 ms
    Reply from 10.2.1.2: bytes=56 Sequence=3 ttl=255 time=1 ms
    Reply from 10.2.1.2: bytes=56 Sequence=4 ttl=255 time=1 ms
    Reply from 10.2.1.2: bytes=56 Sequence=5 ttl=255 time=1 ms
                                                              
