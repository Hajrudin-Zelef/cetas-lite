---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-245
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [35974, 36106]
sha256: 970bd70b66338da7c5998d47a46012c128232f5992ade59c6a9926b5b289dc36
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Primary/Secondary and Active/Inactive
                    PW redundancy involves the following terms:

                    ●   Primary/Secondary: indicates the PW forwarding priority and can be
                        configured.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              569
VPN Configuration
VPN Configuration                                                                        5 VPWS Configuration


                        The primary PW is preferentially used to forward traffic, and the secondary
                        PW is used to protect the primary PW. The primary PW forwards traffic when
                        the primary and secondary PWs work in the same forwarding state. Currently,
                        only one secondary PW can be configured for a primary PW. A bypass PW can
                        also be considered as a secondary PW.
                    ●   Active/Inactive: indicates the PW forwarding status (namely, the PW operating
                        status) and cannot be configured.
                        Traffic can only be forwarded along PWs in the active state. The local and
                        remote signaling status and configured forwarding priority (primary/
                        secondary) determine the PW forwarding status (active/inactive). Only the
                        PW with the optimal signaling status and the highest priority is in the active
                        state and forwards traffic, while other PWs are in the inactive state. PWs in
                        the inactive state cannot forward traffic but can be configured to receive
                        traffic.

PW Redundancy Operating Modes
                    The PW redundancy operating mode can be specified on a PE, on which the
                    primary and secondary PWs are configured. If the PW redundancy operating mode
                    is not specified, PWE3 FRR is used.

                         NOTE

                        In PWE3 FRR, a PE locally determines the PW primary/secondary state and does not notify
                        a remote PE of the primary/secondary state. Therefore, the remote PE is unaware of the PW
                        primary/secondary state. This solution is a private implementation and not recommended.

                    Master/Slave mode

                    In this mode, a PE locally determines the PW primary/secondary state and
                    advertises the state to a remote PE through signaling so that the remote PE can
                    detect the PW primary/secondary state. The primary/secondary state on the PW
                    and AC sides does not affect each other, and therefore PW and AC failures are
                    isolated.

                    Figure 5-39 PW redundancy networking




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   570
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


                    In the PW redundancy scenario shown in Figure 5-39, if the master/slave mode is
                    used, the local end determines the PW primary/secondary state. If the master/
                    slave mode is deployed on PE1, PE1 selects a PW to forward traffic based on the
                    configured priority and PW up/down status. If the primary PW is up, PE1 always
                    selects the configured primary PW to forward traffic. In master/slave PW
                    redundancy mode, PE1 notifies the remote PEs (PE2 and PE3) of the PW primary/
                    secondary state. PE2 and PE3 then determine whether to forward traffic over this
                    primary PW based on the PW primary/secondary state notified by PE1.
                    Independent mode
                    In this mode, the primary/secondary state of the local PWs is determined after
                    negotiation with the remote AC side; the remote end notifies the local end of the
                    primary/secondary state. If a fault occurs on the AC side, it triggers protection
                    switching on both the AC and PW sides, failing to isolate faults.
                    In the PW redundancy scenario shown in Figure 5-39, if the independent mode is
                    used, the remote end determines the PW primary/secondary state. If the
                    independent mode is deployed on PE1, PE1 determines which PW is used to
                    forward traffic based on the PW primary/secondary state notified by the remote
                    PEs (PE2 and PE3), instead of based on the locally configured PW primary/
                    secondary priority. If PE2 notifies its active PW state to PE1, PE1 selects PW1
                    connected to PE2 to forward traffic. If PE3 notifies its active PW state to PE1, PE1
                    selects PW2 connected to PE3 to forward traffic. If both remote PEs notify PE1 of
                    their active PW states, PE2 selects the PW whose status is notified later.

                         NOTE

                        The independent PW redundancy mode is recommended in PWE3 networking to ensure
                        high protection switching performance.

5.12.2 Configuring PW Redundancy in Mater/Slave Mode
Prerequisites
                    Before configuring PW redundancy in master/slave mode, you have completed the
                    following tasks:
                    ●   Configure IP addresses and an IGP on PEs and superstratum provider edges
                        (SPEs).
                    ●   Configure basic MPLS functions on PEs.
                    ●   Establish a public network tunnel between PE1 and an SPE and between PE2
                        and PE3. The public network tunnel can be an LDP or a TE tunnel.
                         NOTE

                        ● Only PWE3 VPWS supports PW redundancy. After the mpls l2vpn default martini
                          command is run, VPWS does not support PW redundancy.
                        ● When configuring PW redundancy, ensure that the parameter settings for the primary,
                          secondary, and bypass PWs are consistent. Parameter setting inconsistencies may result
                          in the secondary or bypass PW failing to take over traffic if the primary PW fails, leading
                          to service interruption.


Context
                    When creating PWs on the local end, configure the PWs that are connected to the
                    two remote PEs. The primary/secondary PW status is manually specified on the

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       571
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration


                    local and remote ends and notified to the remote end. The combined use of PW
                    redundancy in master/slave mode and the bypass PW can prevent network-side
                    faults from affecting the AC side and AC-side faults from affecting the network
                    side.

Procedure
         Step 1 Configure the primary and secondary VPWS connections. For details, see 5.10.3
                Configuring LDP VPWS FRR.

         Step 2 Configure PW redundancy in master/slave mode.
                    1.   Enter the system view.
                         system-view

                    2.   Enter the AC interface view.
                         interface interface-type interface-number

                    3.   Switch the interface working mode to Layer 3.
                         undo portswitch

