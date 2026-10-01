---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-70
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9796, 9945]
sha256: 151cbb7494dce4a56f4e3b8be9ff71d7bcad01d0a8b92d5369b70d0fd39eba65
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Timers associated with LDP GR are as follows:
                 ●      Reconnect timer: After the GR restarter performs an active/standby
                        switchover, the GR Helper detects that the LDP session with the GR Restarter
                        fails, starts the Reconnect timer, and waits for the reestablishment of the LDP
                        session.
                        –    If the Reconnect timer expires before the LDP session between the GR
                             Helper and Restarter is established, the GR Helper immediately deletes
                             MPLS forwarding entries associated with the GR Restarter and exits from
                             the GR Helper process.
                        –    If the LDP session between the GR Helper and the GR Restarter is
                             established before the Reconnect timer times out, the GR Helper deletes
                             the timer and starts the Recovery timer.
                 ●      Recovery timer: After an LDP session is reestablished, the GR Helper starts the
                        Recovery timer and waits for the LSP to recover.
                        –    If the Recovery timer expires, the GR Helper considers that the GR
                             process on the neighbor is complete and deletes non-restored LSPs.
                        –    If all LSPs are restored before the Recovery timer expires, the GR Helper
                             considers that the GR process is complete on the neighbor after the
                             Recovery timer expires.
                 ●      Neighbor-liveness timer: indicates the LDP GR time.
                         NOTE

                        Changing the value of an LDP GR timer causes an LDP session to be reestablished.


3.19.2 Configuring the LDP GR Helper
Context
                 LDP GR Helper must be enabled on the GR Restarter and its neighboring nodes.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 165
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Enabling LDP GR
                 graceful-restart

                 By default, LDP GR is not enabled.

                         NOTE

                        ● Enabling or disabling LDP GR causes an LDP session to be reestablished. If LDP sessions
                          do not need to be reestablished when LDP GR is enabled or disabled, run the no-
                          renegotiate session-parameter-change graceful-restart command.
                        ● The undo mpls ldp and reset mpls ldp commands cannot be run during LDP GR.

         Step 4 (Optional) Configure GR Helper timers.
                         NOTE

                        Changing the value of an LDP GR timer causes an LDP session to be reestablished.
                 ●      Set a value for the Reconnect timer.
                        graceful-restart timer reconnect time

                        The Reconnect timer value that takes effect is the smaller value between the
                        Neighbor-liveness timer value configured on the GR Helper and the Reconnect
                        timer value configured on the GR Restarter.
                        By default, the value of the Reconnect timer is 300 seconds.
                 ●      Set a value for the Recovery timer.
                        graceful-restart timer recovery time

                        The Recovery timer value that takes effect is the smaller value between the
                        Recovery timer value configured on the GR Helper and the Recovery timer
                        value configured on the GR Restarter.
                        By default, the value of the Recovery timer is 300 seconds.
                 ●      Set a value for the Neighbor-liveness timer.
                        graceful-restart timer neighbor-liveness time

                        When negotiating the reconnection time of an LDP session during LDP GR,
                        the device uses the smaller value between the Neighbor-liveness timer value
                        configured on the GR helper and the Reconnect timer value configured on the
                        GR restarter.
                        By default, the value of the Neighbor-liveness timer is 300 seconds.

                 ----End


Result
                 ●      Run the display mpls ldp command to check LDP information.
                 ●      Run the display mpls ldp session [ all ] [ verbose ] command to check LDP
                        session information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    166
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


3.19.3 Example for Configuring LDP GR
Networking Requirements
                 On the network shown in Figure 3-34, LSRA, LSRB, and LSRC each contain a
                 single main control board. During an active/standby switchover or system
                 upgrade, if GR is not enabled, the neighbor deletes the LSP because the session
                 goes down. As a result, traffic and services are interrupted for a short time. In this
                 case, you can configure LDP GR to ensure that labels remain unchanged before
                 and after an unexpected active/standby switchover or protocol restart. In addition,
                 you can restore the establishment of LDP sessions and LSPs after the active/
                 standby switchover or system upgrade is complete. This ensures uninterrupted
                 MPLS forwarding and does not affect services.

                 Figure 3-34 Configuring LDP GR
                         NOTE

                        In this example, interfaces 1 and 2 represent VLANIF100 and VLANIF200, respectively.




Precautions
                 During the configuration, note the following:
                 ●      Enabling or disabling LDP GR causes an LDP session to be reestablished.
                 ●      Changing the value of an LDP GR timer causes an LDP session to be
                        reestablished.

Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface and configure OSPF to advertise the
                        route to the network segment to which each interface is connected and the
                        host route to each LSR ID.
                 2.     Enable MPLS and MPLS LDP globally on each node.
                 3.     Enable MPLS and MPLS LDP on each interface.
                 4.     Enable LDP GR.
                 5.     Set LDP GR parameters on a GR Restarter.

Procedure
         Step 1 Assign an IP address to each interface and configure OSPF to advertise the route
                to the network segment to which each interface is connected and the host route
                to each LSR ID.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     167
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


         Step 2 Enable MPLS and MPLS LDP globally on each node.
                 # Configure LSRA.
                 [LSRA] mpls lsr-id 1.1.1.9
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit

