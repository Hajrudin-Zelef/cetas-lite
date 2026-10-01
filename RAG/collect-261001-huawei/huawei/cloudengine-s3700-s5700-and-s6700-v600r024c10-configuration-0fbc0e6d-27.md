---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-27
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3314, 3456]
sha256: 29e1461a571166ac369a32bc0411b188ece3aa3e138921a0f2167d2d76f7f0c4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.7.2 Adjusting the LDP Transport Address

Context
                 An LDP session is established over a TCP connection. To set up an LDP session, two
                 LSRs need to confirm each other's LDP transport address before they can set up a
                 TCP connection. Generally, the LSR ID (loopback interface address) is used as the
                 transport address.

                         NOTE

                        ● If multiple links exist between two LSRs and LDP sessions need to be established over
                          these links, you are advised to use the same transport address to establish LDP sessions.
                        ● Modifying the LDP transport address interrupts the LDP session. Therefore, exercise
                          caution when performing this operation.
                        ● Modifying the LDP transport address is not recommended.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the interface view of an existing LDP session.
                 interface interface-type interface-number

         Step 3 Configure LDP to use the IP address of a specified interface as the LDP transport
                address.
                 mpls ldp transport-address { interface-type interface-number | interface }

                 ●      interface-type interface-number: specifies the type and number of an
                        interface. This parameter configures LDP to use the address of the specified
                        interface as the TCP transport address.
                 ●      interface: configures LDP to use the IP address of the current interface as the
                        TCP transport address.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         56
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


                 By default, the LSR ID of a node is used as the LDP transport address.

                 ----End

3.7.3 Adjusting LDP Session Timers

Context
                 Table 3-2 lists the timers used by an LDP session. Using default values of these
                 timers is recommended.

                         NOTE

                        When a local session and a remote session coexist, the Keepalive send timer values and
                        Keepalive hold timer values of the local and remote sessions must be the same.


                 Table 3-2 Timers for LDP sessions

                  Timer                    Description                          Usage Suggestion

                  Hello send timers        Determine the interval at            On an unstable network,
                  ● Link Hello             which an LSR sends Hello             set a smaller value for the
                    send timer,            messages to notify its peer          Hello send timer to speed
                    that is, the           LSR of the local LSR's               up network fault detection.
                    Hello send             presence and establish Hello
                    timer in the           adjacencies.
                    local LDP
                    session
                  ● Target Hello
                    send timer,
                    that is, the
                    Hello send
                    timer in the
                    remote LDP
                    session

                  Hello hold timers        Determines the interval at           On a network where the
                  ● Link Hello             which LDP peers exchange             link status is unstable or a
                    hold timer,            Hello messages to maintain           large number of packets are
                    that is, the           their adjacency. If no Hello         sent, set a larger value for
                    Hello hold             message is received after the        the Hello hold timer to
                    timer in the           Hello hold timer expires, the        prevent sessions from being
                    local LDP              Hello adjacency is deleted.          frequently torn down and
                    session                                                     established.
                  ● Target Hello
                    hold timer,
                    that is, the
                    Hello hold
                    timer in the
                    remote LDP
                    session


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       57
MPLS Configuration
MPLS Configuration                                                               3 MPLS LDP Configuration


                  Timer                     Description                      Usage Suggestion

                  Keepalive send            Determines the interval at       On an unstable network,
                  timer                     which LSRs exchange              set a smaller value for the
                                            Keepalive messages to            Keepalive send timer to
                                            maintain their LDP session.      speed up network fault
                                                                             detection.

                  Keepalive hold            Determines the interval at       On a network with unstable
                  timer                     which LDP peers exchange         links, set a larger value for
                                            LDP PDUs to maintain the         the Keepalive hold timer to
                                            remote LDP session. If no        prevent session flapping.
                                            LDP PDU is received after the
                                            Keepalive hold timer expires,
                                            the TCP connection is closed
                                            and the remote LDP session
                                            is terminated.

                  Exponential               After an LSR fails to process    ● Set larger initial and
                  backoff timer             an LDP Initialization message      maximum values to
                                            or is informed that the peer       allow a longer interval
                                            LSR rejects the received LDP       between attempts to
                                            Initialization message, the        establish an LDP session
                                            LSR starts the Exponential         during device upgrade.
                                            backoff timer and periodically   ● Set smaller initial and
                                            resends an LDP Initialization      maximum values to
                                            message to initiate an LDP         allow a shorter interval
                                            session before the                 between attempts to
                                            Exponential backoff timer          establish an LDP session
                                            expires.                           if intermittent service
                                                                               interruptions occur.




Procedure
                 ●      Configure timers for a local LDP session.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the interface view of an existing LDP session.
                             interface interface-type interface-number
                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

