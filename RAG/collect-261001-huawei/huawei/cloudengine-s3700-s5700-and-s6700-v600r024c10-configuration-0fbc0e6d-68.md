---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-68
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9521, 9660]
sha256: 0810949d9cf5ce1a74e8c6b9b7856ec2946e5ceae821ea1c7242df73588f366f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Select parameters based on networking requirements:
                 ●      If an IGP carries only LDP services, specify the infinite parameter to ensure
                        that the behavior for IGP routes is always consistent with that for an LDP LSP.
                 ●      If an IGP carries multiple types of services including LDP services, set the
                        value parameter to ensure that an LDP session or adjacency teardown does
                        not affect IGP route selection or other services.

Procedure
                 ●      Set a value for the Hold-max-cost timer in the interface view.

                        If OSPF is used, perform the following configuration on the local and remote
                        interfaces of the link between the node where the primary and backup links
                        diverge from each other and the LDP peer on the primary link.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the OSPF interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Set a value for the Hold-max-cost timer.

                             ▪    OSPF single-area
                                  ospf timer ldp-sync hold-max-cost { value | infinite }

                             ▪    OSPF multi-area
                                  ospf timer ldp-sync hold-max-cost { value | infinite } multi-area { area-id | area-id-
                                  ipv4 }


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                 160
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


                             The Hold-max-cost timer value determines the interval at which the node
                             advertises the maximum route cost in local LSAs. By default, OSPF always
                             advertises the maximum cost in LSAs on the local device.
                        If IS-IS is used, perform the following configuration on the local and remote
                        interfaces of the link between the node where the primary and backup links
                        diverge from each other and the LDP peer on the primary link.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the interface view.
                             interface interface-type interface-number
                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Set a value for the Hold-max-cost timer.
                             isis timer ldp-sync hold-max-cost { value | infinite }

                             The Hold-max-cost timer value determines the interval at which the node
                             advertises the maximum route cost in local LSAs. By default, IS-IS
                             permanently advertises the maximum route cost in LSAs sent by the local
                             device.
                 ●      Set a value for the Hold-max-cost timer in the IGP process.
                        If OSPF is used, perform the following configuration on the node where the
                        primary and backup links diverge from each other.
                        a.   Enter the system view.
                             system-view
                        b.   Start a specified OSPF process and enter the OSPF view.
                             ospf [ process-id ]

                             process-id specifies an OSPF process. If process-id is not specified, the
                             default process ID 1 is used. To associate an OSPF process with a VPN
                             instance, run the ospf [ process-id | vpn-instance vpn-instance-name ]*
                             command. If a VPN instance is specified, the OSPF process belongs to the
                             specified VPN instance. If no VPN instance is specified, the OSPF process
                             belongs to the global instance.
                        c.   Enter the OSPF area view.
                             area area-id

                             area-id specifies an OSPF area ID. If area-id is not specified, the default
                             area ID 0 is used. Determine whether to perform this step based on the
                             current interface configuration.
                        d.   Set an interval at which all interfaces in the OSPF area advertise the
                             maximum route cost in local LSAs.
                             timer ldp-sync hold-max-cost { infinite | interval }

                        If IS-IS is used, perform the following configuration on the node where the
                        primary and backup links diverge from each other.
                        a.   Enter the system view.
                             system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                161
MPLS Configuration
MPLS Configuration                                                                  3 MPLS LDP Configuration


                        b.   Start a specified IS-IS process and enter the IS-IS view.
                             isis [ process-id ]

                        c.   Set an interval at which all interfaces in the IS-IS process advertise the
                             maximum route cost in local LSPs.
                             timer ldp-sync hold-max-cost { infinite | interval }

                 ----End

3.18.6 (Optional) Setting a Value for the Delay Timer

Context
                 When a faulty link recovers and an LDP session is reestablished on the link, LDP
                 starts the Delay timer to wait for the establishment of an LSP. After the Delay
                 timer expires, LDP notifies the IGP that the synchronization process is complete.


Procedure
                 ●      Perform the following configuration in the MPLS-LDP view.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Set a value for the Delay timer to determine the period during which the
                             device waits for LSP establishment after LDP session is established.
                             igp-sync-delay timer value

                             By default, the device waits for 10s to establish an LSP after an LDP
                             session is established.
                 ●      Perform the following configuration in the interface view.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

