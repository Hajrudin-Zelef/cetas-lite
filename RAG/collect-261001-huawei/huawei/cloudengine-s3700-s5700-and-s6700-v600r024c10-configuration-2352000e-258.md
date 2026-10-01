---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-258
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [37865, 38005]
sha256: 8057f3be36d0d33a54df9a29082fb5401491e369441728e79a13624681bd35b7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    1.   After PE1 is associated with a VSI, has PE2 configured as its peer, and
                         establishes an LDP session with PE2, PE1 sends a Label Mapping message to
                         PE2 in DU mode. The Label Mapping message carries information required to
                         establish a PW, such as the PW ID, VC label, and interface parameters.
                    2.   After receiving the message, PE2 checks whether itself has been associated
                         with the VSI. If PE2 has been associated with the VSI and parameters such as
                         the encapsulation type on PE1 and PE2 are the same, PE1 and PE2 belong to
                         the same VPN. PE2 then establishes a unidirectional VC named VC1 with PE1.
                         In addition, PE2 sends a Label Mapping message to PE1. After receiving the
                         message, PE1 checks the message and takes the same actions as PE2, and
                         then establishes a unidirectional VC named VC2 with PE2.

                    Figure 6-7 shows the process for tearing down a PW using LDP signaling.

                    Figure 6-7 Process for tearing down a PW using LDP signaling




                    1.   After the peer configuration of PE2 is deleted from PE1, PE1 sends a Label
                         Withdrawal message to PE2. After receiving the message, PE2 withdraws its
                         local VC label, tears down VC1, and sends a Label Release message to PE1.
                    2.   After receiving the Label Release message, PE1 withdraws its local VC label
                         and tears down VC2.

Application Scenario
                    LDP VPLS applies to networks that have a small number of sites or do not need to
                    span ASs, especially when PEs do not run BGP.
                    LDP VPLS applies to scenarios where PEs can use LDP as the VPLS signaling
                    protocol. In Figure 6-8, PEs establish full-mesh PWs and forward packets based on
                    split horizon to prevent packet loops. Specifically, a PE forwards packets, including
                    unicast packets, unknown unicast packets, multicast packets, and broadcast
                    packets, received through a PW to only its connected CE, not to other PEs.
                    All the PEs need to establish LDP sessions with each other for PW signaling
                    negotiation. Packets from CEs are sent to the VPLS network through AC interfaces
                    on PEs.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            605
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    Figure 6-8 VPLS networking




Benefits
                    LDP VPLS offers the following benefits:
                    ●     Easy configuration
                    ●     Label resource saving

6.6.2 Creating a VSI and Configuring LDP Signaling
Context
                    When configuring LDP VPLS, configure a VSI ID and a VSI peer. VSI IDs are used to
                    identify VSIs in PW signaling negotiation.
                    Perform the following steps on the PEs at both ends of a PW.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VSI and enter the VSI view.
                    vsi vsi-name [ static | auto ]

                    The names of different VSIs on a device cannot be the same.
         Step 3 (Optional) Configure the encapsulation type of the VSI.
                    encapsulation { ethernet | vlan }

                    By default, the encapsulation type of a VSI is VLAN.
                    The encapsulation types of the VSIs bound to the two ends of a PW must be the
                    same.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                      606
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration


         Step 4 Configure LDP as the signaling protocol of the VSI and enter the VSI-LDP view.
                    pwsignal ldp

                    By default, no signaling is configured for a VSI.
         Step 5 Configure a VSI ID.
                    vsi-id vsi-id
                    ●     The VSI IDs of the two ends of a PW must be the same; otherwise, the VSI
                          cannot be created. If the VSI IDs of the two ends are different, specify the
                          negotiation-vc-id vc-id parameter in the peer command to set the VSI ID
                          used for PW negotiation.
                    ●     VSIs exist only on PEs. A PE can be configured with multiple VSIs, but the ID
                          of each VSI must be unique on a PE.
                    By default, no ID is set for a VSI.
         Step 6 Configure a VSI peer.
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] [ ignore-standby-state ]

                    A VPLS network is a P2MP L2VPN and therefore multiple peers can be specified
                    for one VSI.
         Step 7 (Optional) Delete the VCCV byte following the interface parameter in a Label
                Mapping message.
                    undo interface-parameter-type vccv

                    Run this command if LDP VPLS is configured for a device running VRP V800R006
                    or later and this device communicates with another device running VRP V300R001
                    or any branch version of VRP V300R001. By default, a Label Mapping message
                    carries the VCCV byte following the interface parameter.
         Step 8 (Optional) Configure the current VSI as the mVSI.
                    admin-vsi

                    By default, no mVSI is configured.

                    ----End

Follow-up Procedure
                    If the specified VSI is an mVSI, you must run the track admin-vsi vsi-name
                    command in the service VSI view to bind the service VSI to the mVSI.

6.6.3 Binding an AC Interface to a VSI
Context
                    You can bind an AC interface to a VSI in different views depending on the type of
                    the link between a PE and a CE:
                    ●     Binding an Ethernet interface to a VSI: applies to scenarios where a PE uses
                          a GE interface to connect to a CE.
                    ●     Binding an Ethernet sub-interface to a VSI: applies to scenarios where a PE
                          uses a GE sub-interface to connect to a CE.
                    ●     Binding a VLANIF interface to a VSI: applies to scenarios where a PE uses a
                          VLANIF interface to connect to a CE

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                           607
VPN Configuration
VPN Configuration                                                                          6 VPLS Configuration


