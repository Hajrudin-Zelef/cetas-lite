---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-28
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3457, 3573]
sha256: 1dad60b5a2e7703efdaf83645bcbfdc7f314403f544ce6ed1bc4a095a128fb29
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Configure the link Hello send timer.
                             mpls ldp timer hello-send interval

                             By default, the value of the link Hello send timer is one third the value of
                             the link Hello hold timer.
                             Effective link Hello send timer value = Min {Configured link Hello send
                             timer value, 1/3 of the link Hello hold timer value}

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 58
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                        e.   Configure the link Hello hold timer:
                             mpls ldp timer hello-hold interval

                             The default value of the link Hello hold timer is 15, in seconds.

                             The value of the Hello hold timer configured on the local LSR may not be
                             the actual effective value. The actual effective value is the smaller of the
                             two values configured on the LSRs of a local LDP session.

                             The configured timer value must be greater than or equal to the time
                             required for an active/standby switchover. Otherwise, protocol flapping
                             may occur during an active/standby switchover. The default value is
                             recommended. If the Hello hold timer value negotiated by the two ends
                             of an LDP session is less than 9 seconds, the effective timer value is 9
                             seconds.
                        f.   Configure the Keepalive send timer.
                             mpls ldp timer keepalive-send interval

                             By default, the value of the Keepalive send timer of a local LDP session is
                             one third the value of the Keepalive hold timer.

                             Effective Keepalive send timer value = Min {Configured Keepalive send
                             timer value, 1/3 of the Keepalive hold timer value}
                        g.   Configure the Keepalive hold timer for the local LDP session.
                             mpls ldp timer keepalive-hold interval

                             The default value of the Keepalive hold timer is 45, in seconds.

                             The value of the Keepalive hold timer configured on the local LSR may
                             not be the actual effective value. The actual effective value is the smaller
                             of the two values configured on the LSRs of a local LDP session.

                             For a local multi-link LDP session, the Keepalive hold timer values of all
                             links must be the same. Otherwise, the LDP session or LSP cannot be
                             established.

                             For a coexistent local and remote LDP session, the Keepalive hold timer
                             values of all links and remote peers must be the same. Otherwise, the
                             LDP session or LSP cannot be established.

                             If the Keepalive hold timer values configured on multiple links and
                             remote peers are different, LDP establishes a session based only on the
                             configuration of the adjacency discovered for the first time. Other
                             adjacencies with different configurations cannot be bound to the session.
                             As a result, LSPs cannot be established on these adjacencies.

                             If the Keepalive hold timer negotiated by the two ends of an LDP session
                             is less than 9 seconds, the effective timer value is 9 seconds.


                                 NOTICE

                             Changing the value of the Keepalive hold timer may cause session
                             reestablishment for the related LDP instance, interrupting MPLS services.

                        h.   Configure the Exponential backoff timer.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             59
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration

                             quit
                             mpls ldp
                             backoff timer init max

                             By default, the initial value is 15 and the maximum value is 120, in
                             seconds.
                 ●      Configure timers for a remote LDP session.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the remote MPLS-LDP peer view.
                             mpls ldp remote-peer remote-peer-name
                        c.   Configure the target Hello send timer.
                             mpls ldp timer hello-send interval

                             By default, the value of the target Hello send timer is one third the value
                             of the target Hello hold timer.
                             Effective target Hello send timer value = Min {Configured target Hello
                             send timer value, 1/3 of the target Hello hold timer value}
                        d.   Configure the target Hello hold timer.
                             mpls ldp timer hello-hold interval

                             The default value of the target Hello hold timer is 45, in seconds.
                             The value of the Hello hold timer configured on the local LSR may not be
                             the actual effective value. The actual effective value is the smaller of the
                             two values configured on the two ends of a remote LDP session.
                             The configured timer value must be greater than or equal to the time
                             required for an active/standby switchover. Otherwise, protocol flapping
                             may occur during an active/standby switchover. The default value is
                             recommended.
                        e.   Configure the Keepalive send timer.
                             mpls ldp timer keepalive-send interval

                             By default, the value of the Keepalive send timer of a remote LDP session
                             is one third the value of the Keepalive hold timer.
                             Effective Keepalive send timer value = Min {Configured Keepalive send
                             timer value, 1/3 of the Keepalive hold timer value}
                        f.   Configure the Keepalive hold timer for the remote LDP session.
                             mpls ldp timer keepalive-hold interval

