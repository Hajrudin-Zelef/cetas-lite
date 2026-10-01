---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-55
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [7572, 7716]
sha256: c6faff98d177fe5fe66eafec24064c7930173ec20ea7f866af16925c66a9d420
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 If an LSP is used as a reverse path to notify the ingress of a fault, the process-pst
                 command can be run to allow the reverse path to switch traffic if the BFD session
                 goes down. If an IP link is used as a reverse path, this command can be executed
                 only when the IP link has only one hop. This command does not apply to a multi-
                 hop IP link used as a reverse path.

                 ----End

3.15.5 Verifying the Configuration
                 After configuring static BFD for LDP LSP, you can check the BFD session
                 information, such as the BFD session type and status.

Procedure
                 ●      Run the display bfd session { all | static | discriminator discr-value }
                        [ verbose ] command to check BFD session information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         127
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


                 ●      Run the display bfd statistics session { all | static | discriminator discr-value
                        | peer-ip peer-ip } command to check statistics about BFD sessions.
                 ----End

3.15.6 Example for Configuring Static BFD for LDP LSP
                 This section provides an example for configuring static BFD for LDP LSP. The
                 configuration involves enabling MPLS and MPLS LDP globally and for specific
                 interfaces and enabling BFD on two ends of a link to be monitored.

Networking Requirements
                 On the network shown in Figure 3-25, establish an LDP LSP along the path PE1 ->
                 P1 -> PE2, and an IP link along the path PE2 -> P2 -> PE1. Configure static BFD to
                 monitor the LDP LSP.

                 Figure 3-25 Networking diagram of BFD for LDP LSP
                         NOTE

                        In this example, interfaces 1 and 2 represent VLANIF100 and VLANIF200, respectively.




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure OSPF in the MPLS domain to implement network layer connectivity.
                 2.     Establish an LDP LSP along the path PE1 -> P1 -> PE2.
                 3.     On PE1, configure a BFD session and bind it to the LDP LSP.
                 4.     On PE2, configure a BFD session and bind it to the IP link, enabling PE2 to
                        notify PE1 of LDP LSP faults.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     128
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


Procedure
         Step 1 Assign an IP address to each interface and configure OSPF.

                 Assign an IP address and a mask to each interface (including loopback interfaces)
                 according to Figure 3-25.

                 Configure OSPF on all nodes to advertise the host route of each loopback
                 interface. For configuration details, see Configuration Scripts.

                 After completing the configuration, ping the LSR ID of each peer to check that the
                 LSRs interwork successfully. Run the display ip routing-table command on each
                 LSR to view the routes to the other LSRs.
                 <PE1> display ip routing-table
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 16        Routes : 16

                 Destination/Mask     Proto Pre Cost       Flags NextHop         Interface

                       1.1.1.1/32 Direct 0 0          D 127.0.0.1     LoopBack1
                       2.2.2.2/32 OSPF 10 2            D 10.1.1.2      Vlanif100
                       3.3.3.3/32 OSPF 10 2            D 10.1.2.2      Vlanif200
                       4.4.4.4/32 OSPF 10 3            D 10.1.1.2      Vlanif100
                                OSPF 10 3            D 10.1.2.2    Vlanif200
                      10.1.1.0/24 Direct 0 0          D 10.1.1.1      Vlanif100
                      10.1.1.1/32 Direct 0 0          D 127.0.0.1     Vlanif100
                    10.1.1.255/32 Direct 0 0           D 127.0.0.1      Vlanif100
                      10.1.2.0/24 Direct 0 0          D 10.1.2.1      Vlanif200
                      10.1.2.1/32 Direct 0 0          D 127.0.0.1     Vlanif200
                    10.1.2.255/32 Direct 0 0           D 127.0.0.1      Vlanif200
                      10.1.4.0/24 OSPF 10 3            D 10.1.2.2      Vlanif200
                      10.1.5.0/24 OSPF 10 3            D 10.1.1.2      Vlanif100
                     127.0.0.0/8 Direct 0 0           D 127.0.0.1     InLoopBack0
                     127.0.0.1/32 Direct 0 0           D 127.0.0.1     InLoopBack0
                 127.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                 255.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0

         Step 2 Establish an LDP LSP along the path PE1 -> P1 -> PE2.

                 # Configure PE1.
                 <PE1> system-view
                 [PE1] mpls lsr-id 1.1.1.1
                 [PE1] mpls
                 [PE1-mpls] quit
                 [PE1] mpls ldp
                 [PE1-mpls] quit
                 [PE1]interface vlanif 100
                 [PE1-Vlanif100] mpls
                 [PE1-Vlanif100] mpls ldp
                 [PE1-Vlanif100] quit

                 # Configure P1.
                 <P1> system-view
                 [P1] mpls lsr-id 2.2.2.2
                 [P1] mpls
                 [P1-mpls] quit
                 [P1] mpls ldp
                 [P1-mpls] quit
                 [P1]interface vlanif 100
                 [P1-Vlanif100] mpls
                 [P1-Vlanif100] mpls ldp
                 [P1-Vlanif100] quit


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         129
MPLS Configuration
MPLS Configuration                                                                                    3 MPLS LDP Configuration

                 [P1]interface vlanif 200
                 [P1-Vlanif200] mpls
                 [P1-Vlanif200] mpls ldp
                 [P1-Vlanif200] quit

                 # Configure PE2.
                 <PE2> system-view
                 [PE2] mpls lsr-id 4.4.4.4
                 [PE2] mpls
                 [PE2-mpls] quit
                 [PE2] mpls ldp
                 [PE2-mpls] quit
                 [PE2]interface vlanif 100
                 [PE2-Vlanif100] mpls
                 [PE2-Vlanif100] mpls ldp
                 [PE2-Vlanif100] quit

