---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-216
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [31866, 32018]
sha256: 0504af86641a8c5fc668cd27d5f5a32c7274ea1ebbe8de895c6a6b485eed87f5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Fault detection
                        If a fault occurs on a VPWS network, devices can detect link faults through
                        route convergence. To speed up fault detection on the VPWS network, BFD is
                        deployed on PEs to rapidly detect PW faults. For details about BFD for VPWS,
                        see 5.11 Configuring BFD for VPWS.
                    ●   Primary and secondary PWs
                        A secondary PW can be established on a VPWS network. The secondary PW
                        remains in the up state to ensure that traffic can be quickly switched from the
                        primary PW, preventing traffic loss. When the primary PW is working properly,
                        the secondary PW does not transmit data. When the primary PW is faulty,
                        traffic is quickly switched to the secondary PW.
                        In the networking where CEs are asymmetrically connected to PEs, as shown
                        in Figure 5-33, the link PE1-P1-PE2 is the primary PW, and the link PE1-P2-
                        PE3 is the secondary PW. When the system detects that the primary PW or
                        PE2 to which CE2 is dual-homed is faulty, PE1 switches traffic to the
                        secondary PW, ensuring traffic transmission from CE1 to CE2.
                        In this scenario, PE1 and CE2 terminate fault notification. When PE1 detects a
                        fault, it triggers traffic switching and does not notify the fault to CE1. When
                        CE2 receives the fault notification from PE2, CE2 switches traffic to the
                        secondary PW.

                        Figure 5-33 Switching between primary and secondary PWs in the
                        networking where CEs are asymmetrically connected to PEs




                        In the networking where CEs are symmetrically connected to PEs, as shown in
                        Figure 5-34, CEs terminate fault notification. Fault information must be
                        transmitted to CEs for traffic switching, no matter whether the fault occurs on
                        an AC interface or a PW.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           509
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


                        Figure 5-34 Switching between primary and secondary PWs in the
                        networking where CEs are symmetrically connected to PEs




                    ●   Fast fault notification
                        The following example uses the networking where CEs are asymmetrically
                        connected to PEs to describe the implementation of fast fault notification.
                        In Figure 5-33, switching between primary and secondary PWs ensures the
                        transmission of traffic from CE1 to CE2, and fast fault notification ensures the
                        transmission of traffic from CE2 to CE1.
                        Physical-layer fault notification implements the transfer of fault information
                        between a PW and an AC. PE2 on the primary PW shuts down the
                        corresponding AC interface when it detects a PW failure. As such, CE2 detects
                        the link failure and switches traffic, as shown in Figure 5-35.


                        Figure 5-35 Fast fault notification




                    ●   PW switchback policy
                        On the network where CEs are asymmetrically connected to PEs, when PE1 is
                        notified that the primary PW recovers, PE1 takes action based on the
                        configured PW switchback policy.
                        PW switchback policies are as follows:
                        –   No switchback: Traffic is not switched back to the primary PW.
                        –   Immediate switchback: Traffic is immediately switched back to the
                            primary PW.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             510
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration


                           –      Delayed switchback: Traffic is switched back to the primary PW after a
                                  delay period.
                           After the switchback, the PE immediately notifies the remote PE on the
                           secondary PW of the fault. In addition, either after a delay period or
                           immediately, the PE notifies the remote PE on the secondary PW of fault
                           recovery, preventing packet loss due to transmission delay between PEs.

5.10.2 Configuring BGP VPWS FRR
Prerequisites
                    Before configuring BGP VPWS FRR, you have completed the following tasks:
                    ●      Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                           network to ensure IP connectivity.
                    ●      Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                           network.
                    ●      Establish LDP/BGP sessions between PEs. If the PEs are indirectly connected,
                           establish remote LDP sessions between them.
                    ●      Establish tunnels between PEs based on tunnel policies.

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


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                  511
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration


                    8.   Return to the system view.
                         quit

         Step 5 Configure a remote VPWS connection.
                    1.   Create a BGP VPWS instance and enter the MPLS L2VPN instance view.
                         mpls l2vpn l2vpn-name [ encapsulation { ethernet | vlan } [ control-word | no-control-word ] ]

                    2.   Configure an RD for the MPLS L2VPN instance.
                         route-distinguisher route-distinguisher

                    3.   (Optional) Configure an MTU for the MPLS L2VPN instance.
                         mtu mtu-value

