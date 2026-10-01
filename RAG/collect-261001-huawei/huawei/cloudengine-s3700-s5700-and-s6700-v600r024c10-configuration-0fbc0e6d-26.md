---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-26
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [3154, 3313]
sha256: 0c89ad0d2259075ca5c2637ff18f0dcd247e7b8322b1a3f4f9ff26265e449f63
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      LSRB
                        #
                        sysname LSRB
                        #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            53
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.252
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.252
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                          network 10.2.1.0 0.0.0.3
                        #
                        return

                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        mpls ldp remote-peer LSRA
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.252
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.2.1.0 0.0.0.3
                        #
                        return



3.7 Adjusting LDP Sessions



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        54
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration


3.7.1 Adjusting the LSR IDs of LDP Sessions
Context
                 By default, all LDP sessions of an LSR, including local and remote LDP sessions,
                 use the LSR ID of the LDP instance. However, if all LDP sessions use the same LSR
                 ID in L2VPN/L3VPN over LDP scenarios, VPN services fail to be isolated. To address
                 this issue, configure an exclusive LSR ID for each LDP session. Perform the
                 configuration after an LDP session is created to implement service isolation.

Procedure
                 ●      Configure an LSR ID for a local session.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the interface view of an existing LDP session.
                             interface interface-type interface-number
                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Specify the primary IP address of an interface as the LSR ID of the current
                             LDP session.
                             mpls ldp local-lsr-id interface-type interface-number

                             interface-type interface-number: configures the device to use the primary
                             IP address of the specified interface as the LSR ID.

                                   NOTE

                                  ● If multiple links directly connect an LSR pair, the LSR ID configured on the
                                    interface of each link must be the same. Otherwise, the LDP session uses the
                                    LSR ID of the link that first finds the adjacency, while other links with
                                    different LSR IDs cannot be bound to the LDP session. As a result, LDP LSPs
                                    fail to be established on these links.
                                  ● If both local and remote LDP sessions are to be established between an LSR
                                    pair, LSR IDs configured for the two sessions must be the same. Otherwise,
                                    only the LDP session that first finds the adjacency can be established.
                                  ● The specified interface must be configured with an IP address. Otherwise, LDP
                                    uses 0.0.0.0 as the LSR ID when the specified interface does not have a
                                    primary IP address.
                                  ● Running this command in the interface view resets the adjacency of the
                                    corresponding link.
                                  ● If you run this command multiple times, only the latest configuration takes
                                    effect.
                 ●      Configure an LSR ID for a remote session.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the remote MPLS-LDP peer view.
                             mpls ldp remote-peer remote-peer-name
                        c.   Specify the primary IP address of an interface as the LSR ID of the current
                             LDP session.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        55
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration

                             mpls ldp local-lsr-id interface-type interface-number

                             interface-type interface-number: configures the device to use the primary
                             IP address of the specified interface as the LSR ID.

                                   NOTE

                                  ● If both local and remote LDP sessions are to be established between an LSR
                                    pair, LSR IDs configured for the two sessions must be the same. Otherwise,
                                    only the LDP session that first finds the adjacency can be established.
                                  ● The specified interface must be configured with an IP address. Otherwise, LDP
                                    uses 0.0.0.0 as the LSR ID when the specified interface does not have a
                                    primary IP address.
                                  ● After this command is run, the current remote LDP session will be reset, and a
                                    new LSR ID will be used during the session reset.
                                  ● If you run this command multiple times, only the latest configuration takes
                                    effect.

                 ----End

