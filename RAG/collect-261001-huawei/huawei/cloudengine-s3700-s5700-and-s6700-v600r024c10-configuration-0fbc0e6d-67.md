---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-67
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9374, 9520]
sha256: 07cde78f9e3858b217078ae46c2238ca7aea340535092080ba1db2711b372e4f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        b.   Enter the OSPF interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Block LDP-OSPF synchronization on the interface.

                             ▪     OSPF single-area
                                   ospf ldp-sync block

                             ▪     OSPF multi-area
                                   ospf ldp-sync block multi-area { area-id | area-id-ipv4 }

                 ●      If IS-IS is used, perform the following configuration on the local and remote
                        interfaces of the link between the node where the primary and backup links
                        diverge from each other and the LDP peer on the primary link.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the IS-IS interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Block LDP-IS-IS synchronization on the interface.
                             isis ldp-sync block

                 ----End

3.18.4 (Optional) Setting a Value for the Hold-down Timer
Context
                 On a device that has LDP-IGP synchronization enabled, if the primary physical link
                 recovers, an IGP enters the Hold-down state, and a Hold-down timer starts. Before
                 the Hold-down timer expires, IGP does not establish a neighbor relationship until
                 an LDP session and an LDP adjacency are established. In this manner, LDP and IGP
                 traffic can be switched back to the primary link at the same time.

                         NOTE

                        If IS-IS is used, you can set a value for the Hold-down timer on a specified interface or on
                        all IS-IS interfaces in the IS-IS view.
                        The setting on an interface takes precedence over that in an IS-IS process. If the two
                        settings are different, the setting on the interface takes effect.


Procedure
                 ●      Set a value for the Hold-down timer on a specified OSPF interface.
                        a.   Enter the system view.
                             system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        158
MPLS Configuration
MPLS Configuration                                                                         3 MPLS LDP Configuration


                        b.   Enter the OSPF interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Set a value for the Hold-down timer to determine the interval at which
                             the interface waits for the establishment of an LDP session and an LDP
                             adjacency without establishing an OSPF neighbor relationship.

                             ▪     OSPF single-area
                                   ospf timer ldp-sync hold-down value

                             ▪     OSPF multi-area
                                   ospf timer ldp-sync hold-down value multi-area { area-id | area-id-ipv4 }

                             By default, the value of the Hold-down timer is 10 seconds.
                 ●      Set a value for the Hold-down timer on all IS-IS interfaces.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the IS-IS view.
                             isis [ process-id ]

                        c.   Set a value for the Hold-down timer to determine the interval at which
                             all IS-IS interfaces wait for the establishment of an LDP session and an
                             LDP adjacency without establishing IS-IS neighbor relationships.
                             timer ldp-sync hold-down value

                             By default, the interval for waiting for the establishment of an LDP
                             session and an LDP adjacency is 10 seconds.
                 ●      Set a value for the Hold-down timer on a specified IS-IS interface.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Enable IS-IS on the interface and specify the ID of the IS-IS process to be
                             associated with the interface.
                             isis enable process-id

                        e.   Set a value for the Hold-down timer to determine the interval at which
                             the interface waits for the establishment of an LDP session and an LDP
                             adjacency without establishing an IS-IS neighbor relationship.
                             isis timer ldp-sync hold-down value

                             By default, the interval for waiting for the establishment of an LDP
                             session and an LDP adjacency is 10 seconds.

                 ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    159
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration


3.18.5 (Optional) Setting a Value for the Hold-max-cost Timer

Context
                 After LDP-IGP synchronization is enabled, if the LDP session and LDP adjacency on
                 the primary link fail but the IGP protocol is normal, the IGP on the local node
                 advertises the maximum route cost over Link State Announcements (LSAs) to
                 ensure that the IGP and LDP simultaneously switch traffic to the backup link. You
                 can set a value for the Hold-max-cost timer to adjust the interval at which an IGP
                 advertises the maximum route cost.

                 You can set a value for the Hold-max-cost timer in either of the following ways:
                 ●      Set a value for the Hold-max-cost timer in the interface view.
                        This method is recommended if a Hold-max-cost timer value needs to be set
                        only on a few interfaces.
                 ●      Set a value for the Hold-max-cost timer in the IGP process.
                        If a Hold-max-cost timer value is set in an IGP process, the value takes effect
                        for all interfaces in the process. This method is recommended if a Hold-max-
                        cost timer value needs to be set on many interfaces on the same node.

