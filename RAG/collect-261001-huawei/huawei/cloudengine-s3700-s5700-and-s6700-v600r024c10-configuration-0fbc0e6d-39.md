---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-39
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5004, 5161]
sha256: 3f7115db552a39272572f89aea5285de614b2bd577741c0a9a4a296f51560b0a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             ▪    ip-prefix: allows only the FECs on routes in a specified IP prefix list. If
                                  this parameter is specified, the device sends Label Mapping messages
                                  only for IGP routes in the specified IP prefix list to specified peers.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                84
MPLS Configuration
MPLS Configuration                                                                     3 MPLS LDP Configuration


                             To apply a policy associated with the same FEC range to an LDP peer
                             group or all LDP peers sending Label Mapping messages, specify either
                             peer-group peer-group-name or all in the command.

                                   NOTE

                                 If multiple outbound policies are configured for a specified LDP peer, the earliest
                                 configuration takes effect. For example, the following configurations are
                                 performed in sequence:
                                 outbound peer 2.2.2.2 fec host
                                 outbound peer peer-group group1 fec none
                                 As group1 also contains an LDP peer with peer-id of 2.2.2.2, the following
                                 outbound policy takes effect for the peer:
                                 outbound peer 2.2.2.2 fec host
                                 If two outbound policies are configured in sequence and the peer parameters in
                                 the two commands are the same, the latter configuration overwrites the former.
                                 For example, the following configurations are performed, in this sequence:
                                 outbound peer 2.2.2.2 fec host
                                 outbound peer 2.2.2.2 fec none
                                 The second configuration overwrites the first one. This means that the following
                                 outbound policy takes effect for the LDP peer with peer-id of 2.2.2.2:
                                 outbound peer 2.2.2.2 fec none
                                 MPLS and MPLS LDP must be enabled globally before an outbound policy is
                                 configured.
                                 To delete all outbound policies simultaneously, run the undo outbound peer all
                                 command.

                 ----End

Result
                 Run the display mpls ldp lsp inbound-policy command to check the liberal LSPs
                 for which an inbound policy has taken effect.

3.10.6 Configuring a Load Balancing Mode for LDP Packets
Context
                 In real-world applications, you need to configure a proper load balancing mode on
                 transit nodes based on MPLS service traffic characteristics. You can configure an
                 MPLS load balancing mode in the ECMP view. If a service traffic parameter
                 changes frequently, it is easier to load balance traffic if you use the load balancing
                 mode based on this frequently changing parameter.

                         NOTE

                        If you change the load balancing mode of LDP packets forwarded through LSPs, the change
                        takes effect only on newly established LSPs. If you want the change to take effect on
                        previously created LSPs, run the reset mpls ldp command to restart the LSPs.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable MPLS globally and enter the MPLS view.
                 mpls


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        85
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


         Step 3 Return to the system view.
                 quit

         Step 4 Enter the ECMP view.
                 load-balance ecmp

         Step 5 Configure ECMP.
                 mpls load-balance { top-label | 2nd-label | 3rd-label | src-ip | dst-ip | l4-src-port | l4-dst-port }

                 By default, load balancing is performed based on labels in packets. A maximum of
                 three layers of labels are supported.

                 ----End

3.10.7 Example for Configuring an Inbound LDP Policy
Networking Requirements
                 MPLS LDP services are deployed on the network shown in Figure 3-15. LSRD is a
                 low-performance DSLAM for user access. By default, LSRD receives Label Mapping
                 messages from all peers and uses the routing information in these messages to
                 establish a large number of LSPs. As a result, memory on LSRD is overused and
                 LSRD is overburdened. Configure an inbound LDP policy to allow LSRD to receive
                 only Label Mapping messages destined for LSRC. This ensures that LSRD
                 establishes LSPs only to LSRC, reducing resource consumption.

                 Figure 3-15 Configuring an inbound LDP policy
                         NOTE

                        Interfaces 1 through 3 in this example represent VLANIF100, VLANIF200, and VLANIF300,
                        respectively.




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface, including the loopback interface on
                        each node.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                               86
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


                 2.     Configure OSPF to advertise the route to the network segment of each
                        interface and to advertise the host route to each LSR ID.
                 3.     Enable MPLS and MPLS LDP globally and on the interfaces of each node.
                 4.     Configure an inbound LDP policy.

Procedure
         Step 1 Assign an IP address to each interface and configure an IGP.
                 Assign an IP address and mask to each interface (as shown in Figure 3-15),
                 including the loopback interfaces. Configure OSPF to advertise the route to the
                 network segment to which each interface is connected and the host route to each
                 LSR ID.
         Step 2 Enable MPLS and MPLS LDP globally and on the interfaces of each node.
                 # Configure LSRA.
                 [LSRA] mpls lsr-id 1.1.1.1
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit

                 # Configure LSRB.
                 [LSRB] mpls lsr-id 2.2.2.2
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] quit
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] mpls ldp
                 [LSRB-Vlanif200] quit
                 [LSRB] interface vlanif 300
                 [LSRB-Vlanif300] mpls
                 [LSRB-Vlanif300] mpls ldp
                 [LSRB-Vlanif300] quit

