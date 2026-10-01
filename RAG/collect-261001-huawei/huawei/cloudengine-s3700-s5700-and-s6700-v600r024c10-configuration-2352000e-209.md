---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-209
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [30687, 30845]
sha256: 559d06cd972a674dcb4bdb310b82a46682d8c9fc44996c973547629641287aa7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    On the network shown in Figure 5-29, ASBR1 in AS100 views ASBR2 in AS200 as
                    a CE. Similarly, ASBR2 also views ASBR1 as a CE.
                    The characteristics of Option A are as follows:
                    Easy implementation MPLS forwarding is not required between the PEs
                    functioning as ASBRs; instead, common IP forwarding is used. No special inter-AS
                    configuration is required.
                    Poor scalability
                    ●   The PEs functioning as ASBRs need to manage information about all L2VPNs.
                    ●   An AC interface must be reserved for each PW, because the PE functions as an
                        ASBR in the local AS.
                    ●   If users need to communicate across multiple ASs, transit ASs must support
                        L2VPN. This leads to a heavy configuration workload and significantly affects
                        the transit ASs.
                    Inter-AS Option A applies to scenarios where an L2VPN spans only a few ASs.

5.9.2 Configuring Inter-AS BGP VPWS Option A
Prerequisites
                    Before configuring inter-AS BGP VPWS Option A, you have completed the
                    following tasks:
                    ●   Configure an IGP for each AS on the MPLS backbone network to ensure IP
                        connectivity of the backbone network within each AS.
                    ●   Configure basic MPLS functions on the MPLS backbone network of each AS.
                    ●   Configure MPLS LDP and establish LDP LSPs for the MPLS backbone network
                        of each AS.

Procedure
                    ●   Configure a remote BGP VPWS connection on the PE and ASBR in each AS.
                        For details, see Configuring a Remote BGP VPWS Connection.
                        An ASBR regards its peer ASBR as a local CE and the interface that connects
                        the ASBR to its peer ASBR as an AC interface.
                    ----End

5.9.3 Configuring Inter-AS LDP VPWS Option A
Prerequisites
                    Before configuring inter-AS LDP VPWS Option A, you have completed the
                    following tasks:
                    ●   Configure an IGP for each AS on the MPLS backbone network to ensure IP
                        connectivity of the backbone network within each AS.
                    ●   Configure basic MPLS functions on the MPLS backbone network of each AS.
                    ●   Configure MPLS LDP and establish LDP LSPs for the MPLS backbone network
                        of each AS.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        491
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration


Procedure
                    ●   Configure LDP VPWS on the PE and ASBR in each AS. For details, see
                        Configuring LDP VPWS.
                        Each ASBR regards its peer ASBR a local CE.
                    ----End

5.9.4 Verifying the Configuration
Procedure
                    ●   Run the display mpls l2vc [ vc-id | brief | interface interface-type interface-
                        number ] command to check PW information on the local end.
                    ●   Run the display mpls l2vc remote-info [ vc-id | unmatch | verbose ]
                        command to check PW information on the remote end.
                    ●   Run the display mpls l2vpn connection l2vpn-name [ remote-ce remote-ce-
                        id | down | up | verbose ] command to check VPWS connection information.
                    ----End

5.9.5 Example for Configuring Inter-AS BGP VPWS Option A
Networking Requirements
                    In Figure 5-30, CE1 and CE2 access the backbone network through PE1 in AS100
                    and PE2 in AS200, respectively.
                    Inter-AS BGP VPWS Option A needs to be configured for CE1 and CE2 to
                    communicate because the number of VPWS PWs is small and interfaces between
                    ASBRs need to be used as AC interfaces.

                    Figure 5-30 Network diagram of configuring inter-AS BGP VPWS Option A
                         NOTE

                        In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     492
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure an IGP on the backbone network to ensure IP connectivity within
                         each AS.
                    2.   Configure basic MPLS functions on the backbone network and establish a
                         dynamic LSP between the PE and ASBR in the same AS. If the PE and ASBR
                         are not directly connected, you also need to establish a remote LDP session
                         between them.
                    3.   Configure BGP peers to exchange VPWS information.
                    4.   Establish a BGP VPWS connection between the PE and ASBR in the same AS.

Procedure
         Step 1 Configure IP addresses for interfaces on each device.

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] undo portswitch
                    [CE1-10GE1/0/1] ip address 10.1.1.1 24
                    [CE1-10GE1/0/1] quit

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] interface loopback1
                    [PE1-Loopback1] ip address 1.1.1.9 32
                    [PE1-Loopback1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] undo portswitch
                    [PE1-10GE1/0/2] ip address 10.10.1.1 24
                    [PE1-10GE1/0/2] quit

                    # Configure ASBR1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR1
                    [ASBR1] interface loopback1
                    [ASBR1-Loopback1] ip address 2.2.2.9 32
                    [ASBR1-Loopback1] quit
                    [ASBR1] interface 10ge 1/0/1
                    [ASBR1-10GE1/0/1] undo portswitch
                    [ASBR1-10GE1/0/1] ip address 10.10.1.2 24
                    [ASBR1-10GE1/0/1] quit

                    # Configure ASBR2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR2
                    [ASBR2] interface loopback1
                    [ASBR2-Loopback1] ip address 3.3.3.9 32
                    [ASBR2-Loopback1] quit
                    [ASBR2] interface 10ge 1/0/2
                    [ASBR2-10GE1/0/2] undo portswitch
                    [ASBR2-10GE1/0/2] ip address 10.20.1.1 24
                    [ASBR2-10GE1/0/2] quit

                    # Configure PE2.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       493
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] interface loopback1
                    [PE2-Loopback1] ip address 4.4.4.9 32
                    [PE2-Loopback1] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] undo portswitch
                    [PE2-10GE1/0/1] ip address 10.20.1.2 24
                    [PE2-10GE1/0/1] quit

