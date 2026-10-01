---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-197
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [28776, 28928]
sha256: edb8cf12b19192f76e62779e110f9721e9aa9872760f2e661d8939d9227d4db9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         ● Flow label-based load balancing can be configured successfully only when at least one
                           of the following conditions is true:
                             –    The receive parameter is configured on the local end, and the send parameter is
                                  configured on the remote end.
                             –    The send parameter is configured on the local end, and the receive parameter is
                                  configured on the remote end.
                             –    Both the send and receive parameters are configured on the local and remote
                                  ends.
                         ● If secondary is configured, the flow label-based load balancing function takes effect
                           only for the secondary PW on the interface. If secondary is not configured, the flow
                           label-based load balancing function takes effect only for the primary PW on the
                           interface.
                         ● If static is configured, the flow label-based load balancing function is configured
                           statically. The two PEs on both ends deliver the flow label-based load balancing
                           function, regardless of whether the other end has this function enabled. For dynamic
                           PWs, if static is not configured, the flow label-based load balancing function is
                           negotiated between the local end and remote end. For static PWs, the flow label-based
                           load balancing function is statically configured, regardless of whether static is
                           configured.
                             If the static flow label-based load balancing configuration does not match on both ends,
                             the device discards packets carrying a flow label, causing packet loss.

                    ----End

5.7.5 Verifying the Configuration
Procedure
                    ●    Run the display mpls l2vc [ vc-id | interface interface-type interface-
                         number ] command on a PE to check local VPWS connection information.
                    ●    Run the display mpls l2vc remote-info [ vc-id | unmatch | verbose ]
                         command on a PE to check remote VPWS connection information.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      461
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


                    ●    Run the display pw-template [ pw-template-name ] command to check
                         information about a specified PW template.
                    ●    Run the display mpls label-stack vll interface interface-type interface-
                         number command to check label stack information.

                    ----End

5.7.6 Example for Configuring LDP VPWS (Using an LSP)

Networking Requirements
                    In Figure 5-24, CE1 and CE2 are connected to PE1 and PE2, respectively. PE1 and
                    PE2 are connected over an MPLS backbone network.

                    An LDP VPWS connection needs to be established between PE1 and PE2 over an
                    LSP.


                    Figure 5-24 Network diagram of configuring LDP VPWS (using an LSP)
                          NOTE

                         In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure an Interior Gateway Protocol (IGP) on the backbone network to
                         enable communication between devices on the backbone network.
                    2.   Configure basic MPLS functions on the backbone network and establish LSPs.
                         Establish a remote MPLS LDP peer relationship between PEs at both ends of a
                         PW.
                    3.   Create a VPWS connection between PEs.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     462
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


                                NOTE

                              By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device.
                              If a VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts
                              with LNP. In this case, run the lnp disable command in the system view to disable
                              LNP.


Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 20
                    [CE1] interface 10ge 1/0/2
                    [CE1-10GE1/0/2] port link-type trunk
                    [CE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE1-10GE1/0/2] quit
                    [CE1] interface vlanif 20
                    [CE1-Vlanif20] ip address 10.10.1.1 24
                    [CE1-Vlanif20] quit

                    # Configure CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 10
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] port link-type trunk
                    [CE2-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE2-10GE1/0/1] quit
                    [CE2] interface vlanif 10
                    [CE2-Vlanif10] ip address 10.10.1.2 24
                    [CE2-Vlanif10] quit

                    The configurations of PE1, P, and PE2 are similar to those of CEs. For detailed
                    configurations, see Configuration Scripts.
         Step 2 Configure an IGP on the MPLS backbone network.
                    # Configure PE1.
                    [PE1] interface loopback 0
                    [PE1-LoopBack0] ip address 192.168.2.2 32
                    [PE1-LoopBack0] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 192.168.2.2 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 10.1.1.1 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
         Step 3 Enable MPLS on the MPLS backbone network and establish an LSP and an LDP
                session between PEs.
                    # Configure PE1.
                    [PE1] mpls lsr-id 192.168.2.2
                    [PE1] mpls
                    [PE1-mpls] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     463
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration

                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] mpls
                    [PE1-Vlanif10] mpls ldp
                    [PE1-Vlanif10] quit
                    [PE1] mpls ldp remote-peer 192.168.3.3
                    [PE1-mpls-ldp-remote-192.168.3.3] remote-ip 192.168.3.3
                    [PE1-mpls-ldp-remote-192.168.3.3] quit

