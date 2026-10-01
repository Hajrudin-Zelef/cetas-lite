---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-22
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [2470, 2611]
sha256: a14d163ba2717632a8abc423c093f36a696b870989d357e6b83e441ef7646547
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.6.1 Understanding LDP Sessions
LDP Discovery Mechanism
                 An LDP discovery mechanism is used by LSRs to discover potential LDP peers.
                 There are two types of LDP discovery mechanisms
                 ●      Basic discovery mechanism: used to discover directly connected LDP peers on
                        a link.
                        An LSR periodically sends LDP Hello messages to discover LDP peers and
                        establish local LDP sessions with the peers.
                        LDP Hello messages are encapsulated in UDP packets with a specific multicast
                        destination address and are sent over the LDP port 646. An LDP Hello
                        message carries an LDP identifier and other information, such as the Hello
                        hold time and transport address. If an LSR receives an LDP Hello message on
                        an interface, it can be determined that an LDP peer exists.
                 ●      Extended discovery mechanism: used to discover indirectly connected LDP
                        peers on a link.
                        An LSR periodically sends Targeted Hello messages to a destination address to
                        implement the extended discovery mechanism and to establish a remote LDP
                        session.
                        Targeted Hello messages are encapsulated in UDP packets, carry unicast
                        destination addresses, and are sent over the LDP port 646. A Targeted Hello
                        message carries an LDP identifier and other information, such as the Hello
                        hold time and transport address. If an LSR receives a Targeted Hello message,
                        it can be determined that the LSR has a potential LDP peer.

Process of Establishing an LDP Session
                 Two LSRs exchange Hello messages to establish an LDP session.
                 Figure 3-6 demonstrates the process of establishing an LDP session.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     42
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


                 Figure 3-6 Process of establishing an LDP session




                 1.     Two LSRs send Hello messages to each other. A Hello message contains a
                        transport address used to establish an LDP session. The LSR with the larger
                        transport address serves as the active peer and initiates a TCP connection
                        request. In Figure 3-6, LSR-A functioning as the active peer initiates a TCP
                        connection request to LSR-B functioning as the passive peer.
                 2.     After the TCP connection is successfully established, LSR-A sends an
                        Initialization message to negotiate parameters used to establish an LDP
                        session with LSR-B. The main parameters include the LDP version, label
                        distribution mode, Keepalive hold timer value, maximum PDU length, and
                        label space.
                 3.     Upon receipt of the Initialization message, LSR-B performs operations as
                        follows: If LSR-B rejects some parameters, it sends a Notification message to
                        terminate LDP session establishment. If LSR-B accepts all parameters, it sends
                        an Initialization message and a Keepalive message to LSR-A.
                 4.     Upon receipt of the Initialization message, LSR-A performs operations as
                        follows: If LSR-A rejects some parameters, it sends a Notification message to
                        terminate LDP session establishment. If LSR-A accepts all parameters, it sends
                        a Keepalive message to LSR-B.
                 After each LSR receives the Keepalive messages from the other, the LDP session is
                 successfully established.

Dynamic LDP Capability Advertisement
                 Without dynamic LDP capability advertisement, after an LDP session is configured
                 and an LDP LSP is established, services will be interrupted when features that
                 require session teardown are configured. As such, you are advised to configure
                 dynamic LDP capability advertisement. After this function is deployed, the device
                 dynamically enables or disables the LDP features that support dynamic LDP
                 capability advertisement, thereby ensuring LSP stability. Currently, the following
                 features support dynamic LDP capability advertisement:

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              43
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


                 ●      Global mLDP
                 ●      mLDP MBB

Automatic Remote LDP Session Establishment
                 A common remote LDP session is manually configured on a device at each end of
                 the session. In some scenarios, a local device needs to automatically establish
                 remote LDP sessions with its peers.

                 After a remote LFA-enabled LSR receives a Targeted Hello message with the R bit
                 of 1, the LSR automatically establishes a remote LDP peer relationship with its
                 peer and replies with a Targeted Hello message with the R bit of 0, which triggers
                 the establishment of a remote LDP session. The R bit in the Targeted Hello
                 message indicates whether the receive end needs to periodically reply with a
                 Targeted Hello message: a value of 1 indicates it does; a value of 0 indicates it
                 does not. If the LSR does not receive a Targeted Hello message with the R bit of 1,
                 the LSR deletes the established remote LDP session.

3.6.2 Configuring MPLS LDP Globally
Context
                 MPLS LDP must be configured globally before any MPLS LDP-related features can
                 be configured. Perform the following configuration on each node in an MPLS
                 domain.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Configure an LSR ID for the local node.
                 mpls lsr-id lsr-id

                 Note the following during LSR ID configuration:
                 ●      LSR IDs must be set before other MPLS commands are run.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      Using the address of a loopback interface as the LSR ID is recommended.

         Step 3 Enable MPLS globally and enter the MPLS view.
                 mpls




                        NOTICE

                 Running the undo mpls command deletes all MPLS configurations, including
                 established LDP sessions and LSPs.

         Step 4 Enable LDP globally and enter the MPLS-LDP view.
                 mpls ldp

         Step 5 (Optional) Configure an LSR ID for the LDP instance.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        44
MPLS Configuration
MPLS Configuration                                                                        3 MPLS LDP Configuration

                 lsr-id lsr-id

                 ----End

Result
                 Run the display mpls ldp interface [ interface-type interface-name | verbose |
                 all ] command to check information about MPLS LDP-enabled interfaces.

