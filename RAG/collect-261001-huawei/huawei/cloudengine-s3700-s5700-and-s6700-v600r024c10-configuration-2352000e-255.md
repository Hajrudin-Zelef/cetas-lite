---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-255
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [37395, 37551]
sha256: 01f9840b551ea4719a91f29a81221d6f42a76adac5001dcef539078f66d50f85
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     BGP AD      BGP AD VPLS      BGP AD VPLS supports automatic    BGP AD VPLS
                     VPLS        uses extended    VPLS member discovery and VPLS    integrates the
                                 BGP Update       PW establishment:                 advantages of
                                 messages to      ● Compared with LDP VPLS, BGP     BGP VPLS and
                                 automatically      AD VPLS requires less           LDP VPLS.
                                 discover           configuration to add new
                                 members in a       nodes.
                                 VPLS domain
                                 and uses LDP     ● Compared with BGP VPLS, BGP
                                 forwarding         AD VPLS saves local label
                                 equivalence        resources and is compatible
                                 class (FEC)        with pseudowire emulation
                                 129 for local      edge-to-edge (PWE3).
                                 and remote
                                 VSIs to
                                 automatically
                                 negotiate and
                                 establish VPLS
                                 PWs.




6.3 Configuration Precautions for VPLS

6.4 Default Settings for VPLS
                    Table 6-5 Default settings for VPLS
                     Parameter                                 Default Setting

                     Flow label function                       Disabled

                     MAC address learning                      Enabled

                     VPLS service isolation                    Disabled

                     Sending LDP MAC Withdraw messages         Disabled
                     by a VSI

                     Collecting statistics about the public    Disabled
                     network traffic on a specified VPLS PW




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        597
VPN Configuration
VPN Configuration                                                                                      6 VPLS Configuration




6.5 Configuring Static VPLS
Prerequisites
                    Before configuring static VPLS, you have completed the following tasks:

                    ●     Configure IP addresses and an IGP on PEs and providers (Ps).
                    ●     Configure label switching router (LSR) IDs and enable MPLS and MPLS LDP
                          on PEs and Ps.
                    ●     Enable MPLS L2VPN on PEs.
                    ●     Establish tunnels between PEs to carry L2VPN services.

6.5.1 Creating a VSI and Configuring a Static PW

Context
                    If devices are incapable of establishing a large number of LDP sessions or you
                    want to manually manage and allocate VC labels, configure static VPLS.
                    Configuring static VPLS enables you to manually configure VC labels to establish
                    VPLS PWs, so that there is no need to transmit Layer 2 VC and link information
                    using LDP.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VSI and specify the static member discovery mode for the VSI.
                    vsi vsi-name static

                    The names of different VSIs on a device cannot be the same.

         Step 3 Configure LDP as the signaling protocol of the VSI and enter the VSI-LDP view.
                    pwsignal ldp

         Step 4 Configure a VSI ID.
                    vsi-id vsi-id

                    ●     The VSI IDs of the two ends of a PW must be the same; otherwise, the VSI
                          cannot be created. If the VSI IDs of the two ends are different, specify the
                          negotiation-vc-id vc-id parameter in the peer command to set the VSI ID
                          used for PW negotiation.
                    ●     VSIs exist only on PEs. A PE can be configured with multiple VSIs, but the ID
                          of each VSI must be unique on a PE.

         Step 5 Configure a static VPLS PW.
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] { static-npe | static-upe } trans
                    transmit-label recv receive-label

                    ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                598
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


6.5.2 Binding an AC Interface to a VSI
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


