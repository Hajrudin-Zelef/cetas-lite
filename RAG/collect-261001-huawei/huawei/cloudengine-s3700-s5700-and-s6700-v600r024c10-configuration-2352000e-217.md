---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-217
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32019, 32159]
sha256: 15c70e75ed2d7dcae88b44ef8a68a01cb5e977ade828bdca449cf98fecfa2847
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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

                    7.   Configure the primary BGP VPWS PW.
                         connection [ ce-offset ce-offset-id ] interface interface-type interface-number [ tunnel-policy
                         tunnel-policy-name ] [ raw | tagged ]
                    8.   Configure the secondary BGP VPWS PW.
                         connection [ ce-offset ce-offset-id ] interface interface-type interface-number [ tunnel-policy
                         tunnel-policy-name ] [ raw | tagged ] secondary
                    9.   Return to the system view.
                         quit

         Step 6 (Optional) Configure physical-layer fault notification.

                    After physical-layer fault notification is enabled, the AC interface can quickly
                    detect physical layer faults and trigger fast switchover of upper-layer applications.

                    1.   Enter the AC interface view.
                         interface interface-type interface-number

                    2.   Switch the interface working mode to Layer 3.
                         undo portswitch


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               512
VPN Configuration
VPN Configuration                                                                               5 VPWS Configuration


                         Determine whether to perform this step based on the current interface
                         working mode.
                    3.   Enable physical-layer fault notification.
                         mpls l2vpn trigger if-down

                         By default, physical-layer fault notification is disabled.
                    4.   Return to the system view.
                         quit

         Step 7 (Optional) Configure a switchback policy.
                    1.   Enter the AC interface view.
                         interface interface-type interface-number

                    2.   Switch the interface working mode to Layer 3.
                         undo portswitch

                         Determine whether to perform this step based on the current interface
                         working mode.
                    3.   Configure a switchback policy.
                         mpls l2vpn reroute { { delay delayTime | immediately } [ resume resumeTime ] | never }

                         By default, delayed switchback is used in FRR or master/slave mode.
                         Switchback policies for PEs are as follows:
                         –      Immediate switchback: Traffic is immediately switched back to the
                                primary PW, and the local PE notifies the peer PE on the secondary PW of
                                fault recovery after the time specified by resumeTime.
                         –      Delayed switchback: The local PE switches traffic back to the primary PW
                                after the time specified by delayTime.
                         –      No switchback: The local PE does not switch traffic back to the primary
                                PW until the secondary PW is faulty.
                         In the asymmetric networking with ACs of the Ethernet type:
                         –      If the remote shutdown function is configured on a PE interface
                                connected to a CE, you are advised not to use the policy of immediate
                                switchback, which may lead to network flapping and traffic loss. Instead,
                                you are advised to use the policy of delayed switchback and set
                                delayTime to greater than or equal to 30 seconds.
                         –      If the Ethernet OAM function is configured on a PE interface connected
                                to a CE and a switchback policy is configured, set resumeTime to greater
                                than or equal to 1 second.
                    4.   Return to the system view.
                         quit

                    ----End

5.10.3 Configuring LDP VPWS FRR
Prerequisites
                    Before configuring LDP VPWS FRR, you have completed the following tasks:
                    ●    Configure static routes or an IGP on PEs and Ps of the MPLS backbone
                         network to ensure IP connectivity.
                    ●    Configure basic MPLS functions on the PEs and Ps of the MPLS backbone
                         network.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      513
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration


                    ●      Establish LDP/BGP sessions between PEs. If the PEs are indirectly connected,
                           establish remote LDP sessions between them.
                    ●      Establish tunnels between PEs based on tunnel policies.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Configure MPLS L2VPN and enter the MPLS L2VPN view.
                    mpls l2vpn

         Step 3 Return to the system view.
                    quit

         Step 4 Enter the AC interface view.
                    interface interface-type interface-number [ .subinterface-number ]

         Step 5 Switch the interface working mode to Layer 3.
                    undo portswitch

                    Determine whether to perform this step based on the current interface working
                    mode.

         Step 6 Create a VPWS connection.
                    mpls l2vc { ip-address | pw-template pw-template-name } * vc-id [ [ control-word | no-control-word ] |
                    [ raw | tagged ] | tunnel-policy policy-name | ignore-standby-state ] *

                            NOTE

