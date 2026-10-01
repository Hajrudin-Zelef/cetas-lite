---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-246
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [36107, 36251]
sha256: ce25e555a7e7af2ff2a9652ecb11fa46b71c175450b6279d21dfc4cd66648061
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         Determine whether to perform this step based on the current interface
                         working mode.
                    4.   Set the PW redundancy negotiation mode to master/slave.
                         mpls l2vpn redundancy master

                    5.   Return to the system view.
                         quit

         Step 3 Configure a bypass PW.
                    1.   Enter the AC interface view.
                         interface interface-type interface-number

                    2.   Switch the interface working mode to Layer 3.
                         undo portswitch

                         Determine whether to perform this step based on the current interface
                         working mode.
                    3.   Configure a bypass PW.
                         mpls l2vc { ip-address | pw-template template-name } * vc-id [ [ control-word | no-control-word ] |
                         [ raw | tagged ] | tunnel-policy tnl-policy-name | ignore-standby-state ] * bypass

         Step 4 (Optional) Configure both the primary and secondary PWs on the CSG to receive
                packets, preventing packet loss during a PW switchback.
                    mpls l2vpn stream-dual-receiving

                    By default, the secondary PW cannot receive packets.

         Step 5 (Optional) Manually switch service traffic from the primary PW to the secondary
                PW.
                    mpls l2vpn switchover

                    If PW redundancy in master/slave mode is configured, in normal cases, traffic is
                    transmitted through the primary PW, and the secondary PW functions properly. To
                    transmit traffic over the secondary PW (for example, because of a device upgrade
                    or service re-deployment), you can run this command to switch traffic from the
                    primary PW to the secondary PW. After the device upgrade is complete or service
                    traffic needs to be switched back to the primary PW, you can undo this command

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            572
VPN Configuration
VPN Configuration                                                                            5 VPWS Configuration


                    to switch traffic from the secondary PW back to the primary PW. By default, traffic
                    switchover is not configured.

                    ----End

5.12.3 Configuring PW Redundancy in Independent Mode
Prerequisites
                    Before configuring PW redundancy in independent mode, you have completed the
                    following tasks:
                    ●    Configure IP addresses and an IGP on PEs and SPEs.
                    ●    Configure basic MPLS functions on PEs.
                    ●    Establish a public network tunnel between PE1 and an SPE and between PE2
                         and PE3. The public network tunnel can be an LDP or a TE tunnel.
                          NOTE

                         ● Only PWE3 VPWS supports PW redundancy. After the mpls l2vpn default martini
                           command is run, VPWS does not support PW redundancy.
                         ● When configuring PW redundancy, ensure that the parameter settings for the primary,
                           secondary, and bypass PWs are consistent. Parameter setting inconsistencies may result
                           in the secondary or bypass PW failing to take over traffic if the primary PW fails, leading
                           to service interruption.


Context
                    In independent PW redundancy mode, the primary/secondary status of PWs is
                    determined through signaling-based negotiation. When creating PWs on the local
                    end, configure two PWs that connect to two remote PEs (one PW per PE). After
                    the PWs are configured, they are both in the primary state on the local end. The
                    remote PEs determine their master/backup roles through the management Virtual
                    Router Redundancy Protocol (mVRRP). In this case, one PW is in the active state
                    and the other in the standby state. Finally, the PW endpoints determine the
                    primary and secondary PWs through negotiation.

Procedure
         Step 1 Configure the primary and secondary VPWS connections. For details, see 5.10.3
                Configuring LDP VPWS FRR.
         Step 2 Configure the independent PW redundancy mode.
                    1.   Enter the system view.
                         system-view

                    2.   Enter the AC interface view.
                         interface interface-type interface-number

                    3.   Switch the interface working mode to Layer 3.
                         undo portswitch

                         Determine whether to perform this step based on the current interface
                         working mode.
                    4.   Set the PW redundancy mode to independent.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      573
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration

                         mpls l2vpn redundancy independent

                    5.   Return to the system view.
                         quit

         Step 3 Configure a bypass PW.
                    1.   Enter the AC interface view.
                         interface interface-type interface-number

                    2.   Switch the interface working mode to Layer 3.
                         undo portswitch

                         Determine whether to perform this step based on the current interface
                         working mode.
                    3.   Configure a bypass PW.
                         mpls l2vc { ip-address | pw-template template-name } * vc-id [ [ control-word | no-control-word ] |
                         [ raw | tagged ] | tunnel-policy tnl-policy-name | ignore-standby-state ] * bypass

         Step 4 (Optional) Configure both the primary and secondary PWs on the CSG to receive
                packets, preventing packet loss during a PW switchback.
                    mpls l2vpn stream-dual-receiving

                    By default, the secondary PW cannot receive packets.

         Step 5 Bind a service PW to the mVRRP group on the remote device.
                    mpls l2vc track admin-vrrp interface interface-type interface-number vrid virtual-router-id pw-
                    redundancy

                    The pw-redundancy parameter must be configured. If this parameter is not
                    configured, after the PW is bound to the mVRRP group, packet forwarding on the
                    local PW will be interrupted if the mVRRP state is backup. If the parameter is
                    configured, whether packets can be forwarded depends on the primary/secondary
                    status negotiated between the local and remote ends.

                    ----End

5.12.4 (Optional) Configuring BFD to Detect Link Failures

Context
                    If a failure occurs on a VPWS network, devices can detect link failures through
                    route convergence. However, the detection is slow and cannot meet the
                    requirements of delay-sensitive services, such as VoIP. To speed up failure
                    detection on the VPWS network, deploy BFD on PEs to rapidly detect PW failures
                    and trigger a rapid switchover of upper-layer applications. BFD provides low-
                    overhead failure detection within milliseconds.

