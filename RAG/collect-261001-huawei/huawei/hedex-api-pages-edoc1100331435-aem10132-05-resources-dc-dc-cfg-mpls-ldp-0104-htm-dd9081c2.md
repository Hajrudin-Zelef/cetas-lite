---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-mpls-ldp-0104-htm-dd9081c2
title: "Configure LSRA. The configurations of LSRB and LSRC are similar to the configuration of LSRA, and are not mentioned here."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-mpls-ldp-0104-htm-dd9081c2.md
source_anchor: ""
source_lines: [1, 135]
sha256: 82beb9015b1bd978194db763a7035fed0281d1777fcc234c47fd45f9f05815cb
---

# Configure LSRA. The configurations of LSRB and LSRC are similar to the configuration of LSRA, and are not mentioned here.

As shown in Figure 1, LSRA and LSRC are PEs of the IP/MPLS backbone network. MPLS L2VPN services need to be deployed on LSRA and LSRC to connect VPN sites at Layer 2, so remote LDP sessions need to be deployed between LSRA and LSRC to implement VC label exchange.
If LSRA is directly connected to LSRC, local LDP sessions established on LSRs can be used to set up LDP LSPs to transmit services and exchange VC labels. In this example, LSRA is indirectly connected to LSRC, so remote LDP sessions must be configured. The configuration roadmap is as follows:
# Configure LSRA. The configurations of LSRB and LSRC are similar to the configuration of LSRA, and are not mentioned here.
<Huawei> system-view
[Huawei] sysname LSRA
[LSRA] interface loopback 0
[LSRA-LoopBack0] ip address 10.10.1.1 32
[LSRA-LoopBack0] quit
[LSRA] interface gigabitethernet 1/0/0
[LSRA-GigabitEthernet1/0/0] ip address 10.1.1.1 24
[LSRA-GigabitEthernet1/0/0] quit
[LSRA] ospf 1
[LSRA-ospf-1] area 0
[LSRA-ospf-1-area-0.0.0.0] network 10.10.1.1 0.0.0.0
[LSRA-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
[LSRA-ospf-1-area-0.0.0.0] quit
[LSRA-ospf-1] quit
# Configure LSRA.
[LSRA] mpls lsr-id 10.10.1.1 
[LSRA] mpls
[LSRA-mpls] quit 
[LSRA] mpls ldp
[LSRA-mpls-ldp] quit
# Configure LSRC.
[LSRC] mpls lsr-id 10.10.1.3 
[LSRC] mpls 
[LSRC-mpls] quit 
[LSRC] mpls ldp
[LSRC-mpls-ldp] quit
[LSRA] mpls ldp remote-peer lsrc
[LSRA-mpls-ldp-remote-lsrc] remote-ip 10.10.1.3 
[LSRA-mpls-ldp-remote-lsrc] quit
[LSRC] mpls ldp remote-peer lsra 
[LSRC-mpls-ldp-remote-lsra] remote-ip 10.10.1.1
[LSRC-mpls-ldp-remote-lsra] quit
# After the configuration is complete, run the display mpls ldp session command on the node. The command output shows that the status of the remote LDP session between LSRA and LSRC is Operational.
LSRA is used as an example.
[LSRA] display mpls ldp session
 LDP Session(s) in Public Network                                           
 Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)                  
 A '*' before a session means the session is being deleted. 
 ------------------------------------------------------------------------------
 PeerID             Status      LAM  SsnRole  SsnAge      KASent/Rcv
 ------------------------------------------------------------------------------
 10.10.1.3:0        Operational DU   Passive  0000:00:01   6/6
 ------------------------------------------------------------------------------
 TOTAL: 1 session(s) Found.
# Run the display mpls ldp remote-peer command on the two LSRs of the remote LDP session to view information about the remote peer.
[LSRA] display mpls ldp remote-peer
                         
                       LDP Remote Entity Information
 ------------------------------------------------------------------------------
 Remote Peer Name: lsrc
 Remote Peer IP    : 10.10.1.3              LDP ID        : 10.10.1.1:0
 Transport Address : 10.10.1.1              Entity Status : Active
 Configured Keepalive Hold Timer : 45 Sec
 Configured Keepalive Send Timer : ----
 Configured Hello Hold Timer     : 45 Sec
 Negotiated Hello Hold Timer     : ----
 Configured Hello Send Timer     : 15 Sec
 Configured Delay Timer          : 10 Sec
 Hello Packet sent/received      : 6347/6307
 Label Advertisement Mode        : Downstream Unsolicited                       
 Remote Peer Deletion Status     : No                                           
 Auto-config                     : ---  
 ------------------------------------------------------------------------------
 TOTAL: 1 Peer(s) Found.
LSRA configuration file
#
 sysname LSRA
#
mpls lsr-id 10.10.1.1
mpls
#
mpls ldp
#
mpls ldp remote-peer lsrc
 remote-ip 10.10.1.3
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.1 255.255.255.0
#
interface LoopBack0
 ip address 10.10.1.1 255.255.255.255
#
ospf 1
 area 0.0.0.0
  network 10.10.1.1 0.0.0.0
  network 10.1.1.0 0.0.0.255
#
return
LSRB configuration file
#
 sysname LSRB
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.2 255.255.255.0
#
interface GigabitEthernet2/0/0
 ip address 10.2.1.1 255.255.255.0
#
interface LoopBack0
 ip address 10.10.1.2 255.255.255.255
#
ospf 1
 area 0.0.0.0
  network 10.10.1.2 0.0.0.0
  network 10.1.1.0 0.0.0.255
  network 10.2.1.0 0.0.0.255
#
return
LSRC configuration file
#
 sysname LSRC
#
mpls lsr-id 10.10.1.3
mpls
#
mpls ldp
#
mpls ldp remote-peer lsra
 remote-ip 10.10.1.1
#
interface GigabitEthernet1/0/0
 ip address 10.2.1.2 255.255.255.0
#
interface LoopBack0
 ip address 10.10.1.3 255.255.255.255
#
ospf 1
 area 0.0.0.0
  network 10.10.1.3 0.0.0.0
  network 10.2.1.0 0.0.0.255
#
return
