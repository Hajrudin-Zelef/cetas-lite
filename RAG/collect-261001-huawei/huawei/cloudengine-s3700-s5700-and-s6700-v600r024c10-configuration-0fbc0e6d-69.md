---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-69
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9661, 9795]
sha256: 1ab895778e184089e3380d37df7bf41e6efb7c73c19c1e5c8743e2b625ff9815
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Set a value for the Delay timer to determine the period during which an
                             interface waits for LSP establishment after LDP session is established.
                             mpls ldp timer igp-sync-delay value

                             By default, an interface waits for 10s to establish an LSP after an LDP
                             session is established.

                 ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               162
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration


3.18.7 (Optional) Configuring Graceful Deletion for LDP
Sessions
Context
                 LDP graceful deletion can be configured in LDP-IGP synchronization scenarios to
                 accelerate traffic switching. It helps implement uninterrupted traffic transmission
                 during traffic switching, which improves the reliability of the entire network.
                 If the physical and protocol status of the primary link is normal but the LDP
                 session on the primary link is down, LDP-IGP synchronization enables LDP to
                 inform the IGP of the primary link fault, and the IGP advertises the maximum cost
                 of the primary link. After that, LDP immediately instructs the upstream device to
                 withdraw labels and assigns labels to the upstream device because a new LSP is
                 established on the backup link, which prolongs LSP convergence. As a result,
                 packet loss occurs.
                 After the LDP session on the faulty link goes down, LDP does not immediately
                 instruct the upstream device to withdraw labels; instead, it keeps the labels and
                 LSP and allows traffic to be transmitted on the primary link until LSP convergence
                 is complete on the backup link. This ensures uninterrupted traffic transmission and
                 speeds up LDP-IGP synchronization.
                 Perform the following configuration on the LSR configured with LDP-IGP
                 synchronization.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Enable LDP graceful deletion.
                 graceful-delete

         Step 4 Set a value for the graceful delete timer.
                 graceful-delete timer timer

                 After the LDP session goes down, LDP does not instruct the upstream device to
                 withdraw labels until the graceful delete timer expires.

                        NOTE

                 If the value of the graceful delete timer is too large, the invalid LSP will be kept for a long time,
                 consuming system resources.

                 ----End

3.18.8 Verifying the Configuration
Procedure
                 ●      Run the display mpls ldp command to check global LDP configurations.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      163
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


                 ●      Run the display isis ldp-sync interface command to check information about
                        the synchronization status of LDP and IS-IS on interfaces.
                 ●      Run the display ospf ldp-sync interface all command to check information
                        about the synchronization status of LDP and OSPF on interfaces.

                 ----End


3.19 Configuring the LDP GR Helper

3.19.1 Understanding the LDP GR Helper

LDP GR Helper
                 LDP graceful restart (GR) ensures uninterrupted traffic forwarding on a device
                 (Restarter) that performs an active/standby switchover or protocol restart with the
                 help of its neighbor (Helper).

                 Figure 3-33 LDP GR implementation




                 During an active/standby switchover, if GR is not enabled, the neighbor deletes the
                 LSP because the session is down. As a result, traffic and services are interrupted
                 for a short time. In the preceding situation, if LDP GR is configured, forwarding
                 entries remain unchanged before and after an unexpected active/standby
                 switchover or protocol restart, implementing uninterrupted MPLS forwarding.
                 Figure 3-33 illustrates the process of LDP GR.
                 1.     Before an active/standby switchover, LDP neighbors negotiate the GR
                        capability during the establishment of an LDP session.
                 2.     After detecting that the Restarter performs an active/standby switchover or
                        LDP restarts, the Helper starts the GR Reconnect timer and retains forwarding
                        entries related to the Restarter before the timer expires. The prerequisite for
                        non-stop forwarding is that the Restarter retains MPLS forwarding entries.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           164
MPLS Configuration
MPLS Configuration                                                                  3 MPLS LDP Configuration


                 3.     If the LDP session between the Restarter and Helper is reestablished before
                        the GR Reconnect timer of the Helper expires, the Helper deletes the GR
                        Reconnect timer and starts the GR Recovery timer.
                 4.     The Helper and the Restarter help each other restore the forwarding entries
                        before the Recovery timer expires. After the timer expires, the Helper deletes
                        all Restarter-related forwarding entries that are not restored.
                 5.     After the Restarter performs an active/standby switchover or protocol restart,
                        the Restarter starts the Forwarding State Holding timer. The Restarter
                        preserves the forwarding entries before a restart and restores the forwarding
                        entries before the timer expires with the assistance of the Helper. After the
                        Forwarding State Holding timer expires, the Restarter deletes all forwarding
                        entries that are not restored.

                 The device can function as a Helper to assist the Restarter implement
                 uninterrupted forwarding during an active/standby switchover or protocol restart.

GR Helper Timers
                 Configuring GR Helper timers includes configuring the LDP session Reconnect
                 timer and LSP Recovery timer.

