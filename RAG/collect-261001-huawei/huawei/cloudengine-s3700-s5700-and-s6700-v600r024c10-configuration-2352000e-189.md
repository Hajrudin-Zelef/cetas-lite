---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-189
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27485, 27638]
sha256: 7f3b7b694fc5d211e2decb5e1a4a253a0b30c049d73f04d13f0a9c176aebbbcf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 7 (Optional) Disable MTU match check for the MPLS L2VPN instance.
                    ignore-mtu-match

                    By default, PEs perform MTU match check for an MPLS L2VPN instance.

                    If the MTUs of the same MPLS L2VPN instance on PEs at both ends of a VC are
                    different, the VC cannot go up. If a non-Huawei device does not support MTU
                    match check for MPLS L2VPN instances, perform this step to ignore MTU match
                    check.

         Step 8 Create a CE in the MPLS L2VPN instance.
                    ce ce-name [ id ce-id [ range ce-range ] [ default-offset ce-offset ] ]

         Step 9 Configure a BGP VPWS connection.
                    connection [ ce-offset ce-offset-id ] interface interface-type interface-number [ tunnel-policy tunnel-
                    policy-name ] [ raw | tagged ] [ secondary ]

                    To configure a tunnel policy for a VPWS connection, configure this tunnel policy
                    and then set the tunnel-policy policy-name parameter to reference the tunnel
                    policy. For details about how to configure a tunnel policy, see Configuring a
                    Tunnel Policy.

                    ----End

5.6.3 Configuring a Remote BGP VPWS Connection

Prerequisites
                    Before configuring a remote BGP VPWS connection, you have completed the
                    following tasks:

                    ●    Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                         network to ensure IP connectivity.
                    ●    Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                         network.
                    ●    Establish tunnels between PEs based on tunnel policies.

Context
                    In Figure 5-16, CE1 and CE2 are connected to different PEs. A remote BGP VPWS
                    connection needs to be established between PEs for the two CEs to communicate.

                    Figure 5-16 Network diagram of a remote BGP VPWS connection




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                 439
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure MPLS L2VPN.
                    mpls l2vpn

         Step 3 Return to the system view.
                    quit

         Step 4 Configure BGP peers to exchange VPWS information.
                    1.     Enter the BGP view.
                           bgp as-number

                    2.     Configure a BGP peer and specify its AS number.
                           peer ipv4-address as-number peer-as

                    3.     Configure the source interface for sending BGP packets.
                           peer ipv4-address connect-interface interface-type interface-number

                    4.     Enter the L2VPN-AD address family view.
                           l2vpn-ad-family

                    5.     Enable the route exchange capability between peers.
                           peer ipv4-address enable

                    6.     Enable BGP VPWS.
                           –      Set the signaling mode of all peers to VPWS.
                                  signaling vpws

                           –      Set the signaling mode of a specified peer to VPWS.
                                  peer ipv4-address signaling vpws

                    7.     Return to the BGP view.
                           quit

                    8.     Return to the system view.
                           quit

         Step 5 Configure a remote VPWS connection.
                    1.     Create a BGP VPWS instance and enter the MPLS L2VPN instance view.
                           mpls l2vpn l2vpn-name [ encapsulation { ethernet | vlan } [ control-word | no-control-word ] ]

                    2.     Configure an RD for the MPLS L2VPN instance.
                           route-distinguisher route-distinguisher

                    3.     (Optional) Configure an MTU for the MPLS L2VPN instance.
                           mtu mtu-value

                           The MTU determines the maximum packet size allowed by a VPWS network.
                           If the MTU exceeds the maximum packet size allowed by a VPWS network or
                           an intermediate node (P), there will be packet fragmentation or even dropped
                           packets, which will increase the network transmission load. The MTU is one of
                           VPWS negotiation parameters. If the MTUs of the same VPN instance on the
                           PEs at both ends are different, the two PEs cannot exchange reachability
                           information or establish a PW. An appropriate MTU must be configured for an
                           MPLS L2VPN instance based on the MTU of the interface bound to the L2VPN
                           instance. Specifically, the MTU of an MPLS L2VPN instance cannot exceed the
                           MTU of the interface bound to the L2VPN instance. By default, the MTU of an
                           MPLS L2VPN instance is 1500 bytes.

Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                         440
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration


                    4.   (Optional) Disable MTU match check for the MPLS L2VPN instance.
                         ignore-mtu-match

                         By default, PEs perform MTU match check for an MPLS L2VPN instance.
                         If the MTUs of the same MPLS L2VPN instance on PEs at both ends of a VC
                         are different, the VC cannot go up. If a non-Huawei device does not support
                         MTU match check for MPLS L2VPN instances, perform this step to ignore
                         MTU match check.
                    5.   Configure VPN targets.
                         vpn-target { vpn-target } & <1-16> [ both | export-extcommunity | import-extcommunity ]
                    6.   Create a CE in the MPLS L2VPN instance.
                         ce ce-name [ id ce-id [ range ce-range ] [ default-offset ce-offset ] ]
                    7.   Configure a BGP VPWS connection.
                         connection [ ce-offset ce-offset-id ] interface interface-type interface-number [ tunnel-policy
                         tunnel-policy-name ] [ raw | tagged ] [ secondary ]

                         To configure a tunnel policy for a VPWS connection, configure this tunnel
                         policy and then set the tunnel-policy policy-name parameter to reference the
                         tunnel policy. For details about how to configure a tunnel policy, see
                         Configuring a Tunnel Policy.
                    8.   Return to the system view.
                         quit

                    ----End

5.6.4 Verifying the Configuration
Procedure
                    ●    Run the display mpls l2vpn [ l2vpn-name [ local-ce | remote-ce ] ]
                         command to check BGP VPWS information.
                    ●    Run the display mpls l2vpn connection l2vpn-name [ remote-ce remote-ce-
                         id | down | up | verbose ] command to check BGP VPWS connection
                         information.
                    ●    Run the display mpls l2vpn { export-route-target-list | import-route-
                         target-list } command to check the VPN target list of BGP VPWS.
                    ----End

