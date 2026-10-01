---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-266
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [39035, 39203]
sha256: 036c99994b2151c18bd3729ff76d8cc5f79b12fc2543ee773960ca5355f32e12
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         of PE1 is the same as the export VPN target of PE2 and PE3, and that the
                         export VPN target of PE1 is the same as the import VPN target of PE2 and
                         PE3. The VPN targets of PE2 and PE3 do not need to match.

                    Figure 6-13 BGP VPLS networking




Benefits
                    BGP VPLS offers the following benefits:
                    ●    A dynamic discovery mechanism is used to discover VPN members, simplifying
                         user operations.
                    ●    RRs are used to reduce the number of BGP connections, increasing network
                         scalability.

6.7.2 Enabling MPLS L2VPN
Context
                    Before configuring VPLS, you need to enable MPLS L2VPN and configure label
                    resources.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enable MPLS L2VPN and enter the MPLS L2VPN view.
                    mpls l2vpn

                    ----End

6.7.3 Enabling BGP Peers to Exchange VPLS Information
Context
                    BGP VPLS shares a TCP connection with BGP. Most BGP VPLS configurations are
                    the same as BGP configurations. A major difference between BGP and BGP VPLS is

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         625
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    that the latter requires PEs to function as BGP peers to exchange VPLS label block
                    information in the L2VPN AD address family view.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Configure a BGP peer and specify its AS number.
                    peer ipv4-address as-number as-number

                    By default, no BGP peer is configured, and no AS number is specified for a peer.
         Step 4 Enter the L2VPN AD address family view.
                    l2vpn-ad-family

         Step 5 Enable the route exchange capability between peers.
                    peer ipv4-address enable

                    By default, peers do not exchange route information.
         Step 6 Enable BGP VPLS.
                    ●     Set the signaling mode of all peers or peer groups to VPLS.
                          signaling vpls

                    ●     Set the signaling mode of a specified peer to VPLS.
                          peer ip-address signaling vpls

                    By default, after the peer enable command is run in the L2VPN AD address family
                    view, the BGP AD signaling mode is enabled.

                    ----End

6.7.4 Configuring a VSI and BGP Signaling
Context
                    When configuring BGP VPLS, you need to configure a VSI, which involves
                    configuring BGP signaling, RD, VPN targets, and a site ID.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VSI.
                    vsi vsi-name [ auto | static ]

         Step 3 (Optional) Configure the encapsulation type of the VSI.
                    encapsulation { ethernet | vlan }

                    By default, the encapsulation type of a VSI is VLAN.
                    The encapsulation types of the VSIs bound to the two ends of a PW must be the
                    same.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                       626
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration


         Step 4 Configure BGP signaling for the VSI.
                    pwsignal bgp

                    By default, no signaling is configured for a VSI.
         Step 5 Configure an RD for the VSI.
                    route-distinguisher route-distinguisher

                    By default, no RD is configured for a VSI.
                    On a BGP VPLS network, you must configure an RD for a VSI after creating the
                    VSI, so as to identify the VSI on a PE.
         Step 6 Associate the VSI with one or more VPN targets.
                    vpn-target vpn-target &<1-16> [ both | export-extcommunity | import-extcommunity ]

                    By default, a VSI is not associated with any VPN target.
                    In a BGP VPLS domain, a PW can be established between two PEs only when the
                    import VPN target of the local PE is the same as the export VPN target of the
                    remote PE.
         Step 7 Configure a site ID for the VSI.
                    site site-id [ range site-range ] [ default-offset { 0 | 1 } ]

                    By default, no site ID is configured for a VSI.
                    The site IDs on PEs in the same VSI must be different. The local site ID must be
                    smaller than the remote site-range plus remote default-offset, and greater than
                    or equal to the remote default-offset.

                    ----End

6.7.5 Binding an AC Interface to a VSI
Context
                    You can bind an AC interface to a VSI in different views depending on the type of
                    the link between a PE and a CE:
                    ●     Binding an Ethernet interface to a VSI: applies to scenarios where a PE uses
                          a GE interface to connect to a CE.
                    ●     Binding an Ethernet sub-interface to a VSI: applies to scenarios where a PE
                          uses a GE sub-interface to connect to a CE.
                    ●     Binding a VLANIF interface to a VSI: applies to scenarios where a PE uses a
                          VLANIF interface to connect to a CE
                    ●     Binding an Eth-Trunk interface to a VSI: applies to scenarios where a PE
                          uses an Eth-Trunk interface to connect to a CE.
                    ●     Binding an Eth-Trunk sub-interface to a VSI: applies to scenarios where a PE
                          uses an Eth-Trunk sub-interface to connect to a CE.
                    ●     Binding a dot1q VLAN tag termination sub-interface to a VSI: applies to
                          scenarios where a PE uses a dot1q VLAN tag termination sub-interface to
                          connect to a CE.
                    ●     Binding a QinQ VLAN tag termination sub-interface to a VSI: applies to
                          scenarios where a PE uses a QinQ VLAN tag termination sub-interface to
                          connect to a CE.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                627
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


                    ●   Binding a QinQ stacking sub-interface to a VSI: applies to scenarios where
                        a PE uses a QinQ stacking sub-interface to connect to a CE.
                         NOTE

                        In VPLS applications, different CEs are transparently connected to each other on the same
                        network segment of a LAN through VSIs, and therefore the IP addresses of the CEs must be
                        different. The IP addresses of PE interfaces connected to CEs and those of the CEs must
                        belong to different network segments. Otherwise, local CEs may learn incorrect ARP entries,
                        causing traffic loss between CEs in the same VSI.


Procedure
                    ●   Bind an Ethernet interface to a VSI.

                        Perform the following steps on the PEs at both ends of a PW.

