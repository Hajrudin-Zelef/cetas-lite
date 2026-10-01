---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-283
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41587, 41722]
sha256: 5bd3595b0785a3b9150bfcf2c9560e7435c8d1393488dd7558e717a32091739b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    PW between a UPE and an SPE is called U-PW, and the PW between SPEs is called
                    S-PW.
                    The following describes the packet forwarding process when CE1 sends a packet to
                    CE2:
                    1.   CE1 sends a packet with the destination MAC address of CE2 to UPE1.
                    2.   When receiving the packet from CE1, UPE1 encapsulates the packet with two
                         MPLS labels and forwards the resulting packet to SPE1. The outer label
                         identifies the LSP between UPE1 and SPE1, and the inner label identifies the
                         VC between UPE1 and SPE1.
                    3.   The LSR between UPE1 and SPE1 transmits the packet and swaps the labels
                         of the packet. The outer label is popped at the penultimate hop.
                    4.   When receiving the packet, SPE1 determines the VSI to which the packet
                         belongs according to the inner MPLS label, which is VSI1 in this example.
                    5.   SPE1 removes the inner MPLS label that is added to the packet by UPE1.
                    6.   SPE1 searches for a matching VSI entry according to the destination MAC
                         address of the packet, and finds that this packet needs to be sent to SPE2.
                         Then SPE1 adds two MPLS labels to the packet. The outer label identifies the
                         LSP between SPE1 and SPE2, and the inner label identifies the VC between
                         SPE1 and SPE2.
                    7.   The LSR between SPE1 and SPE2 transmits the packet and swaps the labels of
                         the packet. The outer label is popped at the penultimate hop.
                    8.   When receiving the packet from the S-PW side, SPE2 determines the VSI to
                         which the packet belongs according to the inner MPLS label, which is VSI1 in
                         this example. SPE2 then removes the inner MPLS label that is added to the
                         packet by SPE1.
                    9.   SPE2 adds two MPLS labels to the packet and forwards the packet. The outer
                         label identifies the LSP between SPE2 and UPE2, and the inner label identifies
                         the VC between UPE2 and SPE2.
                    10. The LSR between SPE2 and UPE2 transmits the packet and swaps the labels
                        of the packet. The outer label is popped at the penultimate hop.
                    11. When receiving the packet, UPE2 removes the inner MPLS label that is added
                        to the packet by itself, searches for a matching VSI entry according to the
                        destination MAC address of the packet, and finds that the packet needs to be
                        sent to CE2. UPE2 then forwards the packet to CE2.
                    In Figure 6-21, CE1 and CE4 exchange data locally. UPE1 directly forwards the
                    packet between CE1 and CE4 without sending the packet to SPE1, because UPE1
                    has the bridging function. However, for the first packet or a broadcast packet with
                    an unknown destination MAC address sent from CE1, UPE1 broadcasts the packet
                    to CE4 and forwards the packet to SPE1 through the U-PW. SPE1 then copies the
                    packet and forwards it to each peer CE.

6.9.2 Configuring an SPE
Context
                    On an HVPLS network, SPEs are fully meshed. You need to configure VSI peer
                    relationships between SPEs, and between SPEs and UPEs.
                    Perform the following steps on SPEs.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            668
VPN Configuration
VPN Configuration                                                                                     6 VPLS Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VSI.
                    vsi vsi-name [ auto | static ]

         Step 3 Configure LDP as the signaling protocol of the VSI and enter the VSI-LDP view.
                    pwsignal ldp

         Step 4 Configure a VSI ID.
                    vsi-id vsi-id

         Step 5 Configure VSI peer relationships between SPEs.
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ]

                    Or
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] static-npe trans transmit-label
                    recv receive-label

         Step 6 Configure VSI peer relationships between SPEs and UPEs.
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] upe

                    Or
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] static-upe trans transmit-label
                    recv receive-label

         Step 7 (Optional) Enable an SPE to forward the LDP MAC Withdraw messages received
                from NPEs to UPEs.
                    npe-upe mac-withdraw enable

                    By default, an SPE does not forward the LDP MAC Withdraw messages received
                    from NPEs to UPEs.
         Step 8 (Optional) Enable an NPE to forward the LDP MAC Withdraw messages received
                from a UPE to other UPEs.
                    upe-upe mac-withdraw enable

                    By default, an NPE does not forward the LDP MAC Withdraw messages received
                    from a UPE to other UPEs.
         Step 9 (Optional) Enable an NPE to forward the LDP MAC Withdraw messages received
                from UPEs to other NPEs.
                    upe-npe mac-withdraw enable

                    By default, an NPE does not forward the LDP MAC Withdraw messages received
                    from UPEs to other NPEs.

                    ----End

6.9.3 Configuring a UPE
Context
                    The configurations of UPEs are similar to the configurations of PEs on a full-mesh
                    VPLS network. The difference is that UPEs are connected only to SPEs and no

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                  669
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration


                    special configuration is required for HVPLS. For configuration details, see 6.6
                    Configuring LDP VPLS.

6.9.4 Verifying the Configuration
Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                        information.
                    ●   Run the display vsi [ name vsi-name ] mac-withdraw loop-detect command
                        to check information about MAC Withdraw loop detection.
                    ----End

6.9.5 Example for Configuring LDP HVPLS
Networking Requirements
                    Figure 6-22 shows a backbone network built by an enterprise. Site1 and Site2
                    connect to the backbone network by connecting CE1 and CE2, respectively, to the
                    UPE, and Site3 connects to the backbone network by connecting CE3 to PE1. Users
                    at Site1, Site2, and Site3 need to communicate at Layer 2, and user information
                    needs to be retained in Layer 2 packets when the packets are transmitted over the
                    backbone network. Additionally, the UPE and SPE need to be deployed at different
                    layers of the backbone network.

                    Figure 6-22 Network diagram of configuring LDP HVPLS
                         NOTE

