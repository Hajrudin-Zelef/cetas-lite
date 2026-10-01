---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-29
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3574, 3714]
sha256: b6522b46b5f197e8e5795964a7b908fd23a43abf25c054a6070f5860fa5389f3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             The default value of the Keepalive hold timer is 45, in seconds.
                             The value of the Keepalive hold timer configured on the local LSR may
                             not be the actual effective value. The actual effective value is the smaller
                             of the two values configured on the two ends of a remote LDP session.
                             For a local multi-link LDP session, you are advised to set the same
                             Keepalive hold time for all links.
                             For a coexistent local and remote LDP session, you are advised to set the
                             same Keepalive hold time for all links and remote peers.
                             If multiple links and remote peers are configured with different Keepalive
                             hold time, they can be bound to the same session. In this case, LDP

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             60
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


                             establishes a session based only on the configuration of the adjacency
                             discovered for the first time. In the following situations:

                             ▪    If a session has only one adjacency, changing the Keepalive hold time
                                  causes the session to be reestablished.

                             ▪    If a session has more than one adjacency, changing the Keepalive
                                  hold time of an adjacency does not change the Keepalive hold time
                                  of the session nor cause session reestablishment.

                             ▪    If a session has only one adjacency (local or remote) left and the
                                  Keepalive hold time of the adjacency is different from that of the
                                  session, the session is reestablished based on the configuration of the
                                  current adjacency.

                             If the Keepalive hold timer negotiated by the two ends of an LDP session
                             is less than 9 seconds, the effective timer value is 9 seconds.
                        g.   Configure the global Keepalive hold timer for the remote LDP session.
                             quit
                             mpls ldp
                             timer auto-remote keepalive-hold interval

                             By default, the value of the global Keepalive hold timer is 45s. The
                             default value is recommended. On a network with unstable links, increase
                             the value of a Keepalive hold timer to prevent LDP session flapping.

                             The value of the Keepalive hold timer configured on the local LSR may
                             not be the actual effective value. The actual effective value is the smaller
                             of the two values configured on the two ends of a remote LDP session.
                             The Keepalive hold timer configured in the remote MPLS-LDP peer view
                             takes precedence over the global Keepalive hold timer. If the Keepalive
                             hold timer is configured both globally and in the remote MPLS-LDP peer
                             view, the Keepalive hold timer configured in the remote MPLS-LDP peer
                             view takes effect.

                             If there is more than one LDP link between two LSRs, the values of the
                             Keepalive hold timers set for the links must be the same. Otherwise, the
                             LDP session may be unstable.

                             If multiple links exist between two devices to maintain a session or a
                             local and remote coexistent session exists, the values of the Keepalive
                             timers for all the links or for the local and remote devices must be the
                             same.

                                   NOTE

                                 Changing the value of the Keepalive hold timer may cause session
                                 reestablishment for the related LDP instance, interrupting MPLS services.
                        h.   Configure the Exponential backoff timer.
                             backoff timer init max

                             By default, the initial value is 15 and the maximum value is 120, in
                             seconds.

                 ----End

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   61
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration




3.8 Configuring the Dynamic LDP Advertisement
Capability
                 On devices enabled with global LDP, the dynamic LDP advertisement capability
                 allows extended LDP functions to be dynamically enabled or disabled when the
                 LDP session is working properly, ensuring stable LSP operation.

Prerequisites
                 Before configuring the dynamic LDP advertisement capability, complete the
                 following task:
                 ●      3.6.2 Configuring MPLS LDP Globally

Context
                 On a device disabled from dynamic LDP advertisement, if an extended LDP
                 function is enabled after an LDP session is created, the LDP session will be
                 interrupted and the extended LDP function will be negotiated, affecting LSP
                 stability. After the dynamic LDP advertisement capability is enabled, the LDP
                 features that support the dynamic LDP advertisement capability can be
                 dynamically enabled or disabled without interrupting sessions, improving the
                 stability of LSPs.

                         NOTE

                        The dynamic LDP advertisement capability does not affect existing LDP functions. You are
                        advised to enable this function immediately after LDP is enabled globally, as it facilitates
                        dynamic advertisement of new extended functions.
                        Before enabling the dynamic LDP advertisement capability, enable MPLS and MPLS LDP
                        globally.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Enable the dynamic LDP advertisement capability.
                 capability-announcement

                         NOTE

                        Enabling dynamic LDP advertisement after an LDP session is established will result in
                        reestablishment of the LDP session.

                 ----End

Result
                 After configuring a local LDP session, run the display mpls ldp command to check
                 whether the dynamic LDP advertisement capability has been enabled. The

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        62
MPLS Configuration
MPLS Configuration                                                              3 MPLS LDP Configuration


                 capability has been enabled if the Capability-Announcement field in the
                 command output is On.


