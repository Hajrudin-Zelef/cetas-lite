---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-275
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40408, 40561]
sha256: a8866de91cfd3ab71a01fca93f17c384747415347a2d9c404b2b254ab315d2a0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 3 Configure a BGP peer and specify its AS number.
                    peer ipv4-address as-number as-number


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                      648
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration


                    By default, no BGP peer is configured, and no AS number is specified for a peer.

         Step 4 Enter the L2VPN AD address family view.
                    l2vpn-ad-family

         Step 5 Enable the route exchange capability between peers.
                    peer ipv4-address enable

                    By default, peers do not exchange route information.

         Step 6 Enable BGP VPLS.

                    By default, after the peer enable command is run in the L2VPN AD address family
                    view, the BGP AD signaling mode is enabled.

                    ----End

6.8.3 Creating a VSI and Configuring BGP AD Signaling

Context
                    When configuring BGP AD VPLS, you need to create VSIs on PEs, set automatic
                    VPLS member discovery and PW deployment for the VSIs, configure BGP AD
                    signaling on the PEs, and set VPLS IDs and VPN targets for the VSIs in the VSI-BGP
                    AD view.

                    Perform the following steps on the PEs at both ends of a PW.

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

         Step 4 Create or enter the BGP AD view.
                    bgp-ad

         Step 5 Specify the ID of a VPLS domain to which the VSI belongs.
                    vpls-id vplsIdValue

                    By default, no VPLS ID is configured for a VSI. You must configure a VPLS ID when
                    creating a BGP AD VSI.

         Step 6 Associate the VSI with one or more VPN targets.
                    vpn-target vpn-target &<1-16> [ both | export-extcommunity | import-extcommunity ]

                    By default, a VSI is not associated with any VPN target.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                 649
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


                    In a BGP AD VPLS domain, a PW can be established between two PEs only when
                    the import VPN target of the local PE is the same as the export VPN target of the
                    remote PE.
         Step 7 (Optional) Configure all PWs of a BGP AD VSI as spoke PWs.
                    pw spoke-mode

                    By default, the PWs of a BGP AD VSI are hub PWs.
                    When BGP AD VPLS is applied to a star or tree network (only one PE functions as
                    the hub device and the others are spoke devices), the hub device can be a server
                    or an authentication device. In this case, you need to run this command to
                    configure all PWs of a VSI as spoke PWs on the hub device, so that split horizon is
                    disabled for PWs.

                    ----End

6.8.4 Binding an AC Interface to a VSI
Context
                    You can bind an AC interface to a VSI in different views depending on the type of
                    the link between a PE and a CE:
                    ●   Binding an Ethernet interface to a VSI: applies to scenarios where a PE uses
                        a GE interface to connect to a CE.
                    ●   Binding an Ethernet sub-interface to a VSI: applies to scenarios where a PE
                        uses a GE sub-interface to connect to a CE.
                    ●   Binding a VLANIF interface to a VSI: applies to scenarios where a PE uses a
                        VLANIF interface to connect to a CE
                    ●   Binding an Eth-Trunk interface to a VSI: applies to scenarios where a PE
                        uses an Eth-Trunk interface to connect to a CE.
                    ●   Binding an Eth-Trunk sub-interface to a VSI: applies to scenarios where a PE
                        uses an Eth-Trunk sub-interface to connect to a CE.
                    ●   Binding a dot1q VLAN tag termination sub-interface to a VSI: applies to
                        scenarios where a PE uses a dot1q VLAN tag termination sub-interface to
                        connect to a CE.
                    ●   Binding a QinQ VLAN tag termination sub-interface to a VSI: applies to
                        scenarios where a PE uses a QinQ VLAN tag termination sub-interface to
                        connect to a CE.
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

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     650
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                        Perform the following steps on the PEs at both ends of a PW.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   (Optional) Run the following commands as required:

                             ▪    Configure the default VLAN of the main interface.
                                  mpls l2vpn default vlan vlanid

                                  By default, no default VLAN is configured for a main interface.

                             ▪    Add a VLAN tag to the packets passing through the main interface.
                                  mpls l2vpn vlan-stacking stack-vlan vlanid

                                  By default, the system does not add a VLAN tag to a packet passing
                                  through a main interface.
                                   NOTE

