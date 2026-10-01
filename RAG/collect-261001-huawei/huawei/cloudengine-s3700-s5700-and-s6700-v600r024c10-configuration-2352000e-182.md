---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-182
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [26414, 26577]
sha256: 23a6bd30f6e125d4050d8a1eb97c14f53706b1350cbc2efe320f3131531e789c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        c.   (Optional) Configure a description for the remote CCC connection.
                             ccc cccName description text

                             By default, no description is configured for a CCC connection.
                    ●   Configure the P.

                        Perform the following operations on the P along the VC.

                        a.   Enter the system view.
                             system-view

                        b.   Configure the P as the transit LSR for a static CR-LSP.
                             static-cr-lsp transit lsp-name [ incoming-interface { incoming-interfacename | { incoming-
                             interfacetype incoming-interfacenum } } ] in-label in-label { outgoing-interface { outgoing-
                             interfacename | { outgoing-interfacetype outgoing-interfacenum } } | nexthop next-hop-
                             address } * out-label out-label [ ingress-lsrid ingress-lsrid egress-lsrid egress-lsrid tunnel-id
                             tunnel-id ] [ bandwidth { [ ct0 ] bandwidth } ] *

                             Bidirectional static CR-LSPs must be configured on all Ps between the PEs
                             for remote CCC connections to transmit CCC data only.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                  419
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration


                              If the outbound interface specified on a P is a non-P2P interface (such as
                              an Ethernet interface), you must configure nexthop in this command to
                              specify a next-hop IP address.

                    ----End

5.5.4 Verifying the Configuration

Procedure
                    ●   Run the display vll ccc [ ccc-name | type { local | remote } ] command to
                        check CCC connection information.
                    ●   Run the display l2vpn ccc-interface vc-type ccc [ down | up ] command to
                        check information about interfaces used by the CCC connection.
                    ●   Run the display mpls l2vpn vpws [ interface { interface-name | interface-
                        type interface-number } [ verbose ] ] command to check VPWS service
                        information.

                    ----End

5.5.5 Example for Configuring a Local CCC Connection

Networking Requirements
                    In Figure 5-5, CEs (CE1 and CE2) are connected to the PE through VLANIF
                    interfaces.

                    A local CCC connection needs to be established between CE1 and CE2.

                    Figure 5-5 Network diagram of configuring a local CCC connection
                         NOTE

                        In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     420
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


                    1.   Configure basic MPLS functions and enable MPLS L2VPN on the PE.
                    2.   Create a local CCC connection between CE1 and CE2 on the PE. Because a
                         local CCC connection is bidirectional, only one connection is required.
                                NOTE

                         By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device. If a
                         VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts with LNP.
                         In this case, run the lnp disable command in the system view to disable LNP.


Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 10
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 10.1.1.1 24
                    [CE1-Vlanif10] quit

                    # Configure CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 20
                    [CE2] interface 10ge 1/0/2
                    [CE2-10GE1/0/2] port link-type trunk
                    [CE2-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE2-10GE1/0/2] quit
                    [CE2] interface vlanif 20
                    [CE2-Vlanif20] ip address 10.1.1.2 24
                    [CE2-Vlanif20] quit

         Step 2 Configure an LSR ID and enable MPLS and MPLS L2VPN on the PE.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE
                    [PE] vlan batch 10 20
                    [PE] interface loopback 1
                    [PE-LoopBack1] ip address 1.1.1.9 32
                    [PE-LoopBack1] quit
                    [PE] mpls lsr-id 1.1.1.9
                    [PE] mpls
                    [PE-mpls] quit
                    [PE] mpls l2vpn
                    [PE-l2vpn] quit
                    [PE] interface 10ge 1/0/1
                    [PE-10GE1/0/1] port link-type trunk
                    [PE-10GE1/0/1] port trunk allow-pass vlan 10
                    [PE-10GE1/0/1] quit
                    [PE] interface 10ge 1/0/2
                    [PE-10GE1/0/2] port link-type trunk
                    [PE-10GE1/0/2] port trunk allow-pass vlan 20
                    [PE-10GE1/0/2] quit

         Step 3 Establish a local CCC connection between CE1 and CE2.
                    [PE] ccc ce1-ce2 interface Vlanif 10 out-interface Vlanif 20

                    ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     421
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


Verifying the Configuration
                    # Check CCC connection information on the PE. The command output shows that
                    a local CCC connection has been established and is in the up state.
                    <PE> display vll ccc
                    total ccc vc : 1
                    local ccc vc : 1, 1 up
                    remote ccc vc : 0, 0 up

                    name: ce1-ce2, type: local, state: up,
                    intf1: Vlanif10 (up),access-port: false

                    intf2: Vlanif20 (up),access-port: false
                    VC last up time : 2024/03/08 05:13:10
                    VC total up time: 0 days, 0 hours, 0 minutes, 14 seconds

                    # Check information about the interfaces used for the CCC connections on PEs.
                    The command output shows that the VC type is CCC and the VC status is up.
                    <PE> display l2vpn ccc-interface vc-type all
                    Total ccc-interface of CCC: 2
                    up (2), down (0)
                    Interface               Encap Type         State   VC Type
                    Vlanif10                ethernet          up     ccc
                    Vlanif20                ethernet          up     ccc

