---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-78
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10674, 10833]
sha256: 5b743243e23c472b8c962053dbbea11138d4a346f3d545265bedfbed50657c12
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

3.13.8 Configuring Route Exchange Between the CE and ASBR

Context
                    The configuration of route exchange between a CE and an ASBR is similar to that
                    between a CE and a PE on a basic IPv4 L3VPN.

Procedure
         Step 1 Determine whether to use BGP, IGP, or static routes between a PE and a CE as
                required. For details, see 3.6 Configuring Basic IPv4 L3VPN over MPLS.

                    ----End

3.13.9 Configuring Route Exchange Between the CE and PE

Context
                    BGP, IGP, or static routes (including the default routes) can be used between a CE
                    and a PE. Determine which one to use as required.

Procedure
         Step 1 Determine whether to use BGP, IGP, or static routes between a PE and a CE as
                required. For details, see 3.6 Configuring Basic IPv4 L3VPN over MPLS.

                    ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                            169
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


3.13.10 Verifying the Configuration

Prerequisites
                    Inter-AS VPN Option B (ASBRs also functioning as PEs) has been configured.

Procedure
                    ●    Run the display bgp vpnv4 all peer command on PEs or ASBRs to check the
                         establishment of all BGP peer relationships.
                    ●    Run the display bgp vpnv4 all routing-table command on PEs or ASBRs to
                         check VPNv4 routes.
                    ●    Run the display ip routing-table vpn-instance vpn-instance-name command
                         on PEs or ASBRs to check VPN routing table information.
                    ----End

3.13.11 Example for Configuring IPv4 L3VPN over MPLS Inter-
AS Option B (ASBRs Also Functioning as PEs)

Networking Requirements
                    In a scenario where the backbone network spans two ASs, ASBRs need to
                    advertise labeled VPN-IPv4 routes through MP-EBGP and ASBRs also function as
                    PEs.
                    In inter-AS VPN Option B, ASBRs also function as PEs to manage VPN routes in
                    addition to advertising VPNv4 routes between ASs. Inter-AS VPN Option B (with
                    ASBRs also functioning as PEs) reduces the number of required PEs, but poses
                    higher performance requirements for ASBRs.
                    On the network shown in Figure 3-31, it is required that inter-AS VPN Option B
                    be configured and ASBRs also function as PEs to allow CEs to communicate.

                    Figure 3-31 Inter-AS VPN Option B (with ASBRs also functioning as PEs)
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     170
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration




Precautions
                    Note the following during the configuration:

                    ●    Create VPN instances on ASBRs and configure routing information exchange
                         between ASBRs and CEs.
                    ●    ASBRs do not filter received VPNv4 routes based on VPN targets.

Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure IGP on the backbone network for IP connectivity between the ASBR
                         and PE in the same AS, and establish an MPLS LDP LSP between the ASBR
                         and PE in the same AS.
                    2.   Establish an MP-IBGP peer relationship between the ASBR and PE in the same
                         AS.
                    3.   Configure VPN instances on PEs and ASBRs, and establish EBGP peer
                         relationships between PEs, ASBRs, and CEs.
                    4.   Enable MPLS on the interfaces that connect ASBRs to each other, and
                         establish an MP-EBGP peer relationship between the ASBRs.

Procedure
         Step 1 Configure IGP on the MPLS backbone networks in AS 100 and AS 200 for
                communication between PEs on each backbone network. OSPF is used as IGP in
                this example. For detailed configurations, see Configuration Scripts.

                    After completing the configuration, run the display ospf peer command on an
                    ASBR or PE. The command output shows that the OSPF neighbor relationship is in

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         171
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration


                    the Full state, indicating that the OSPF neighbor relationship has been established
                    between the ASBR and PE in the same AS.
                    The ASBR and PE in the same AS can learn and ping the address of each other's
                    loopback interface.
         Step 2 Configure basic MPLS functions and MPLS LDP on the MPLS backbone networks in
                AS 100 and AS 200 respectively to establish MPLS LDP LSPs.
                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] mpls
                    [PE1-Vlanif100] mpls ldp
                    [PE1-Vlanif100] quit

                    The configuration of PE2 is similar to the configuration of PE1. For detailed
                    configurations, see Configuration Scripts.
                    # Configure ASBR1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR1
                    [ASBR1] mpls lsr-id 2.2.2.9
                    [ASBR1] mpls
                    [ASBR1-mpls] quit
                    [ASBR1] mpls ldp
                    [ASBR1-mpls-ldp] quit
                    [ASBR1] interface Vlanif 100
                    [ASBR1-Vlanif100] mpls
                    [ASBR1-Vlanif100] mpls ldp
                    [ASBR1-Vlanif100] quit

                    The configuration of ASBR2 is similar to the configuration of ASBR1. For detailed
                    configurations, see Configuration Scripts.
                    After the configuration is complete, an LDP session can be established between
                    the PE and ASBR. Run the display mpls ldp session command on each device. The
                    command output shows that the session status is Operational. The following
                    example uses the command output on PE1.
                    <PE1> display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     -------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    --------------------------------------------------------------------------
                     2.2.2.9:0        Operational DU Passive 0000:00:01 5/5
                    --------------------------------------------------------------------------
                    TOTAL: 1 Session(s) Found.

