---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-219
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32252, 32382]
sha256: 4382a61881f6d6f5ca149706ad209667a7c0a23d1b2bbd804b247828a0c125c0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         Switchback policies for PEs are as follows:
                         –      Immediate switchback: Traffic is immediately switched back to the
                                primary PW, and the local PE notifies the peer PE on the secondary PW of
                                fault recovery after the time specified by resumeTime.
                         –      Delayed switchback: The local PE switches traffic back to the primary PW
                                after the time specified by delayTime.
                         –      No switchback: The local PE does not switch traffic back to the primary
                                PW until the secondary PW is faulty.
                         In the asymmetric networking with ACs of the Ethernet type:
                         –      If the remote shutdown function is configured on a PE interface
                                connected to a CE, you are advised not to use the policy of immediate
                                switchback, which may lead to network flapping and traffic loss. Instead,
                                you are advised to use the policy of delayed switchback and set
                                delayTime to greater than or equal to 30 seconds.
                         –      If the Ethernet OAM function is configured on a PE interface connected
                                to a CE and a switchback policy is configured, set resumeTime to greater
                                than or equal to 1 second.
                    4.   Return to the system view.
                         quit

                    ----End

5.10.4 Verifying the Configuration

Procedure
                    ●    Run the display mpls l2vpn [ l2vpn-name [ local-ce | remote-ce ] ]
                         command to check BGP VPWS information.
                    ●    Run the display mpls l2vpn connection l2vpn-name [ remote-ce remote-ce-
                         id | down | up | verbose ] command to check BGP VPWS connection
                         information.
                    ●    Run the display mpls l2vpn { export-route-target-list | import-route-
                         target-list } command to check the VPN target list of BGP VPWS.
                    ●    Run the display mpls l2vc [ vc-id | interface interface-type interface-
                         number ] command on a PE to check local VPWS connection information.
                    ●    Run the display mpls l2vc remote-info [ vc-id | unmatch | verbose ]
                         command on a PE to check remote VPWS connection information.

                    ----End

5.10.5 Example for Configuring LDP VPWS FRR - Asymmetric
Access of CEs to PEs

Networking Requirements
                    In Figure 5-36, the MPLS network of an ISP provides the L2VPN service for users.
                    Many users are connected to the MPLS network through PE1, PE2, and PE3, and
                    new sites will be added in the future. A proper VPN solution is required to provide
                    secure VPN services for users and to simplify configuration when new users are

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          516
VPN Configuration
VPN Configuration                                                                            5 VPWS Configuration


                    connected to the network. In addition, this solution must ensure highly stable
                    communication between CE1 and CE2.

                          NOTE

                         In this scenario, to avoid loops, ensure that connected interfaces have STP disabled and are
                         removed from VLAN 1. On an STP-enabled ring network, if VLANIF interfaces are used to
                         construct a Layer 3 network, an interface on the network will be blocked. As a result, Layer
                         3 services on the network cannot run normally.


                    Figure 5-36 Network diagram of configuring LDP VPWS FRR - asymmetric access
                    of CEs to PEs
                          NOTE

                         In this example, interface1, interface2, interface3 and interface4 represent VLANIF 10,
                         VLANIF 20, VLANIF 30, and 10GE1/0/1, respectively.




Configuration Roadmap
                    VPWS FRR can be configured to ensure highly stable communication between CE1
                    and CE2. If only a small number of new sites will be added in the future, LDP
                    VPWS FRR can be configured.

                    The configuration roadmap is as follows:

                    1.   Configure OSPF on the backbone network.
                    2.   Establish an MPLS TE tunnel between PE1 and PE3, and an LSP between PE1
                         and PE2. The PW between PE1 and PE3 is the primary PW and uses the MPLS
                         TE tunnel.
                    3.   Establish an MPLS LDP session between PE1 and PE2, and establish a remote
                         MPLS LDP session between PE1 and PE3. The PW between PE1 and PE2 is the
                         secondary PW and uses an MPLS LSP.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        517
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


                    4.   Configure PWs on PEs using PW templates. The primary PW uses an MPLS TE
                         tunnel, so a tunnel policy needs to be used during the configuration of the
                         primary PW.
                    5.   Establish BFD for PW sessions between PE1 and PE2, and between PE1 and
                         PE3 to detect PW failures.
                    6.   Enable physical-layer fault notification on PE2 and PE3. When BFD detects a
                         fault of the primary PW, the AC on the dual-homed CE goes down. L2VPN
                         traffic is then quickly switched to the secondary PW. When BFD detects failure
                         recovery of the primary PW, L2VPN traffic can be switched back to the
                         primary PW.

Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 20 30
                    [CE1] interface vlanif 30
                    [CE1-Vlanif30] ip address 10.1.1.1 255.255.255.252
                    [CE1-Vlanif30] ip address 10.1.2.1 255.255.255.252 sub
                    [CE1-Vlanif30] quit
                    [CE1] interface vlanif 20
                    [CE1-Vlanif20] ip address 10.1.3.1 255.255.255.0
                    [CE1-Vlanif20] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk pvid vlan 30
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 30
                    [CE1-10GE1/0/1] quit
                    [CE1] interface 10ge 1/0/2
                    [CE1-10GE1/0/2] port link-type trunk
                    [CE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE1-10GE1/0/2] quit

