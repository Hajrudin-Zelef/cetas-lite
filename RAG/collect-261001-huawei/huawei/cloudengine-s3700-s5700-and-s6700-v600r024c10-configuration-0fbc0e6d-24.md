---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-24
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [2773, 2978]
sha256: 6527b45c2c7a4193c3d1abc65282d811c93af07553639ea0ed82614d8d7d89fc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
         Step 1 Assign an IP address to each interface and configure OSPF to advertise the route
                to the network segment to which each interface is connected and the host route
                to each LSR ID.
                 Assign an IP address to each interface (as shown in Figure 3-7), including the
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

                 # Configure LSRB.
                 <LSRB> system-view
                 [LSRB] mpls lsr-id 2.2.2.9
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] quit

                 # Configure LSRC.
                 <LSRC> system-view
                 [LSRC] mpls lsr-id 3.3.3.9
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit

         Step 3 Enable MPLS and MPLS LDP on the interfaces of each LSR.
                 # Configure LSRA.
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit

                 # Configure LSRB.
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] mpls ldp
                 [LSRB-Vlanif200] quit

                 # Configure LSRC.
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] mpls ldp
                 [LSRC-Vlanif100] quit

         Step 4 Verify the configuration.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        48
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration


                 # After completing the configuration, check whether the status of the local LDP
                 sessions between LSRA and LSRB and between LSRB and LSRC is Operational.
                 The following example uses the command output on LSRA.
                 <LSRA> display mpls ldp session
                  LDP Session(s) in Public Network
                  Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                  An asterisk (*) before a session means the session is being deleted.
                 --------------------------------------------------------------------------
                  PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                 --------------------------------------------------------------------------
                 2.2.2.9:0         Operational DU Passive 0000:00:22 91/91
                 --------------------------------------------------------------------------
                 TOTAL: 1 Session(s) Found.

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
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.252
                         mpls
                         mpls ldp
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
                 ●      LSRB
                        #
                        sysname LSRB
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        49
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                          network 10.2.1.0 0.0.0.3
                        #
                        return

                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.2.1.0 0.0.0.3
                        #
                        return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        50
MPLS Configuration
MPLS Configuration                                                                  3 MPLS LDP Configuration


3.6.5 Example for Configuring a Remote LDP Session

