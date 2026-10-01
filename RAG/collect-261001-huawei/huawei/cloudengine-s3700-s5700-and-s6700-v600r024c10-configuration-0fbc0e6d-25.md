---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-25
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [2979, 3153]
sha256: 6902d5af626ff654d0f9d7fd3b131fb1e2f4467fc714f6f4564fe8c49462b6bd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Networking Requirements
                 As shown in Figure 3-8, LSRA and LSRC function as edge devices on the backbone
                 network. Configure a remote LDP session between LSRA and LSRC to establish an
                 LSP for the VPN service.

                 Figure 3-8 Configuring a remote LDP session
                         NOTE

                        In this example, interface1 and interface2 represent VLANIF100 and VLANIF200,
                        respectively.




Precautions
                 During the configuration, note the following:

                 ●      LSR IDs must be set before other MPLS commands are run.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      Using the IP address of a reachable loopback interface on an LSR as the LSR
                        ID is recommended.
                 ●      The IP address of a remote peer must be its LSR ID. When an LDP LSR ID is
                        different from an MPLS LSR ID, the LDP LSR ID must be used.

Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Assign an IP address to each interface and configure OSPF to advertise the
                        route to the network segment to which each interface is connected and the
                        host route to each LSR ID.
                 2.     Enable MPLS and MPLS LDP globally on the LSRs of a remote LDP session.
                 3.     Specify the name and remote peer IP address on the LSRs of a remote LDP
                        session.

Data Preparation
                 To complete the configuration, prepare the following data:

                 ●      IP address of each interface on each node (as shown in Figure 3-8), OSPF
                        process ID, and area ID
                 ●      LSR ID of each node

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               51
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


                 ●      Name and IP address of each remote peer of a remote LDP session

Procedure
         Step 1 Assign an IP address to each interface.

                 Assign an IP address to each interface (as shown in Figure 3-8), including the
                 loopback interfaces. Configure OSPF to advertise the route to the network
                 segment to which each interface is connected and the host route to each LSR ID.

         Step 2 Enable MPLS and MPLS LDP globally on each LSR.

                 # Configure LSRA.
                 <LSRA> system-view
                 [LSRA] mpls lsr-id 1.1.1.9
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit

                 # Configure LSRC.
                 <LSRC> system-view
                 [LSRC] mpls lsr-id 3.3.3.9
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit

         Step 3 Specify the name and remote peer IP address on the LSRs of a remote LDP
                session.

                 # Configure LSRA.
                 [LSRA] mpls ldp remote-peer LSRC
                 [LSRA-mpls-ldp-remote-LSRC] remote-ip 3.3.3.9
                 [LSRA-mpls-ldp-remote-LSRC] quit

                 # Configure LSRC.
                 [LSRC] mpls ldp remote-peer LSRA
                 [LSRC-mpls-ldp-remote-LSRA] remote-ip 1.1.1.9
                 [LSRC-mpls-ldp-remote-LSRA] quit

         Step 4 Verify the configuration.

                 # After completing the configuration, check whether the status of the remote LDP
                 session between LSRA and LSRC is Operational.

                 The following example uses the command output on LSRA.
                 <LSRA> display mpls ldp session
                  LDP Session(s) in Public Network
                   Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                   An asterisk (*) before a session means the session is being deleted.
                  --------------------------------------------------------------------------
                   PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                  --------------------------------------------------------------------------
                  3.3.3.9:0        Operational DU Passive 0000:00:01 6/6
                  --------------------------------------------------------------------------
                  TOTAL: 1 Session(s) Found.

                 # Run the display mpls ldp remote-peer command on the LSRs of the remote
                 LDP session to check information about the remote peer.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          52
MPLS Configuration
MPLS Configuration                                                                                 3 MPLS LDP Configuration


                 The following example uses the command output on LSRA.
                 <LSRA> display mpls ldp remote-peer
                                    LDP Remote Entity Information
                  ------------------------------------------------------------------------------
                  Remote Peer Name : LSRC
                  Description        : ----
                  Remote Peer IP        : 3.3.3.9           LDP ID        : 1.1.1.9:0
                  Transport Address : 1.1.1.9               Entity Status : Active

                  Configured Keepalive Hold Timer : 45 Sec
                  Configured Keepalive Send Timer : ----
                  Configured Hello Hold Timer            : 45 Sec
                  Negotiated Hello Hold Timer            : 45 Sec
                  Configured Hello Send Timer            : ----
                  Configured Delay Timer              : 10 Sec
                  Hello Packet sent/received          : 6347/6307
                  Label Advertisement Mode               : Downstream Unsolicited
                  Auto-config                   : ----
                  Manual-config                   : effective
                  Session-Protect effect           : NO
                  Session-Protect Duration           : ----
                  Session-Protect Remain             : ----
                  ------------------------------------------------------------------------------
                  TOTAL: 1 Remote-Peer(s) Found.

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        mpls ldp remote-peer LSRC
                         remote-ip 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.252
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                        #
                        return

