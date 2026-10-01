---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-298
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [44063, 44222]
sha256: 7f5b0c9b5ac84879cd0cb98efffc14af0ac97e5ad93aa6c10f2db12837b4061e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 2 Establish BGP AD PWs. The configuration is similar to that in "Configuring BGP AD
                VPLS." For details, see 6.8.3 Creating a VSI and Configuring BGP AD Signaling.
                         NOTE

                        ● LDP PWs and BGP AD PWs must be configured in the same VSI.
                        ● If you run the vsi vsi-name [ static ] command to create a VSI, LDP PWs must be
                          established prior to BGP AD PWs. If you run this command without specifying the static
                          keyword, there is no requirement on the PW configuration sequence.

                    ----End

6.12.3 Configuring Interworking Between LDP VPLS and BGP
VPLS

Context
                    On the network shown in Figure 6-30, an LDP VPLS network is deployed between
                    PE1 and the SPE, and a BGP VPLS network is deployed among the SPE, PE2, and
                    PE3. Interworking between LDP VPLS and BGP VPLS needs to be configured on the
                    SPE for CE1, CE2, and CE3 to communicate.

                    Figure 6-30 Interworking between LDP VPLS and BGP VPLS




Procedure
         Step 1 Configure LDP VPLS from PE1 to the SPE. For configuration details, see 6.6
                Configuring LDP VPLS.

         Step 2 Perform the following configurations on the SPE:
                    ●   Configure LDP VPLS from the SPE to PE1. For configuration details, see 6.6
                        Configuring LDP VPLS.
                    ●   Configure BGP VPLS from the SPE to PE2 and PE3. For configuration details,
                        see 6.7 Configuring BGP VPLS.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  708
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                         NOTE

                        On the SPE, LDP VPLS and BGP VPLS must be configured in the same VSI.

         Step 3 Configure BGP VPLS from PE2 to the SPE and PE3. For configuration details, see
                6.7 Configuring BGP VPLS.

         Step 4 Configure BGP VPLS from PE3 to the SPE and PE2. For configuration details, see
                6.7 Configuring BGP VPLS.

                    ----End

6.12.4 Verifying the Configuration

Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                        information.
                    ●   Run the display vsi name vsi-name statistics command to check statistics
                        about a specified VSI, such as the PW status and AC status.
                    ●   Run the display vpls connection [ vsi vsi-name ] [ down | up ] [ verbose ]
                        command to check VPLS connections.

                    ----End

6.12.5 Example for Configuring Interworking Between LDP
VPLS and BGP AD VPLS in HVPLS Mode

Networking Requirements
                    On the network shown in Figure 6-31, PE3, PE4, and PE5 need to be fully meshed
                    using BGP AD VPLS connections. LDP VPLS needs to be deployed between PE1 and
                    PE2. PE1 and PE2 support only LDP VPLS, and PE3 supports both LDP VPLS and
                    BGP AD VPLS. CE1 and CE2 must be able to communicate with each other, and
                    PE1 and PE2 must function as UPEs to access PE3 in HVPLS mode.

                    The specific PW deployment requirements are as follows:
                    ●   Establish an LDP PW from PE1 to PE2 and from PE1 to PE3.
                    ●   Establish an LDP PW from PE2 to PE1 and from PE2 to PE3.
                    ●   Establish an LDP PW from PE3 to PE1 and from PE3 to PE2, and establish a
                        BGP AD PW from PE3 to PE4 and from PE3 to PE5.
                    ●   Establish a BGP AD PW from PE4 to PE3 and from PE4 to PE5.
                    ●   Establish a BGP AD PW from PE5 to PE3 and from PE5 to PE4.

                         NOTE

                        To avoid loops in this scenario, ensure that all connected interfaces have STP disabled and
                        are removed from VLAN 1. If STP is enabled and VLANIF interfaces of switches are used to
                        construct a Layer 3 ring network, an interface on the network will be blocked. As a result,
                        Layer 3 services on the network cannot run properly.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      709
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration


                    Figure 6-31 Interworking between LDP VPLS and BGP AD VPLS in HVPLS mode
                          NOTE

                         In this example, interface1, interface2, interface3 and interface4 represent 10GE1/0/1,
                         10GE1/0/2, 10GE1/0/3, and 10GE1/0/4, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure an IP address and a routing protocol for each involved interface so
                         that PEs can communicate at the network layer. This example uses OSPF as
                         the routing protocol.
                    2.   Configure MPLS and public network tunnels to carry PWs. In this example,
                         LDP LSPs are used between PEs.
                    3.   Configure PE1, PE2, and PE3 to form an LDP VPLS network.
                               NOTE

                         When configuring an LDP PW on PE3, you need to specify the peer end as the UPE.
                    4.   Configure BGP AD VPLS on PE3, PE4, and PE5.

Procedure
         Step 1 Configure VLANs to which interfaces belong and assign IP addresses to the
                corresponding VLANIF interfaces.

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan 10
                    [CE1-vlan10] quit
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 192.168.10.1 255.255.255.0
                    [CE1-Vlanif10] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       710
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    The configurations of CE2, PE1, PE2, PE3, PE4, and PE5 are similar to the
                    configuration of CE1. For detailed configurations, see Configuration Scripts.
         Step 2 Configure a routing protocol for communication between devices.
                    In this example, OSPF is configured.
                    # Configure PE1.
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 192.168.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 192.168.2.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

