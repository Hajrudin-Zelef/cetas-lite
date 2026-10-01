---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-66
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9229, 9373]
sha256: 25cdac3feed276d5f0420d70d53d0f7d4a2cb2985ae006038cb195e6444edf61
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.18.2 Enabling LDP-IGP Synchronization
Context
                 LDP-IGP synchronization can be configured in either of the following modes:
                 ●      Enable LDP-IGP synchronization in the interface view.
                        Enable LDP-IGP synchronization on specified interfaces if a few interfaces
                        need to support this function.
                 ●      Enable LDP-IGP synchronization in an IGP process.
                        After LDP-IGP synchronization is enabled in an IGP process, it is automatically
                        enabled on all interfaces in the process. If LDP is enabled on all IGP-enabled
                        interfaces of a node, this configuration mode is recommended.

Procedure
                 ●      Enable LDP-IGP synchronization in the interface view.
                        If OSPF is used, perform the following configuration on the local and remote
                        interfaces of the link between the node where the primary and backup links
                        diverge from each other and the LDP peer on the primary link.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            155
MPLS Configuration
MPLS Configuration                                                                       3 MPLS LDP Configuration


                        a.   Enter the system view.
                             system-view

                        b.   Enter the interface view.
                             interface interface-type interface-number

                        c.   Switch the interface working mode to Layer 3.
                             undo portswitch

                             Determine whether to perform this step based on the current interface
                             working mode.
                        d.   Enable LDP-OSPF synchronization on the protected interface.

                             ▪     OSPF single-area
                                   ospf ldp-sync

                             ▪     OSPF multi-area
                                   ospf ldp-sync multi-area { area-id | area-id-ipv4 }

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
                        d.   Enable IS-IS.
                             isis enable process-id

                        e.   Enable LDP-IS-IS synchronization on the protected interface.
                             isis ldp-sync

                             NOTE

                        If both LDP-IGP synchronization and LDP GTSM are configured on an interface, the number
                        of GTSM hops must be set based on the actual hop count because LDP needs to establish
                        an LDP session over an indirect link. The number of GTSM hops cannot be set to 1. If the
                        value is set to 1, LDP sessions cannot be established. As a result, routes cannot be switched
                        back or LDP-IGP synchronization fails.
                 ●      Enable LDP-IGP synchronization in an IGP process.
                        If OSPF is used, perform the following configuration on the node where the
                        primary and backup links diverge from each other and on its LDP peer on the
                        primary link.
                        a.   Enter the system view.
                             system-view

                        b.   Start a specified OSPF process and enter the OSPF view.
                             ospf [ process-id ]

                             process-id specifies an OSPF process. If process-id is not specified, the
                             default process ID 1 is used. To associate an OSPF process with a VPN

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      156
MPLS Configuration
MPLS Configuration                                                               3 MPLS LDP Configuration


                             instance, run the ospf [ process-id | vpn-instance vpn-instance-name ]*
                             command. If a VPN instance is specified, the OSPF process belongs to the
                             specified VPN instance. If no VPN instance is specified, the OSPF process
                             belongs to the global instance.
                        c.   Enter the OSPF area view.
                             area area-id

                             area-id specifies an OSPF area ID. If area-id is not specified, the default
                             area ID 0 is used. Determine whether to perform this step based on the
                             current interface configuration.
                        d.   Enable LDP-OSPF synchronization.
                             ldp-sync enable

                        If IS-IS is used, perform the following configuration on the node where the
                        primary and backup links diverge from each other and on its LDP peer on the
                        primary link.
                        a.   Enter the system view.
                             system-view

                        b.   Start a specified IS-IS process and enter the IS-IS view.
                             isis [ process-id ]

                             process-id specifies an IS-IS process ID. If process-id is not specified, the
                             default process ID 1 is used. To associate an IS-IS process with a VPN
                             instance, run the isis [ process-id ] [ vpn-instance vpn-instance-name ]
                             command.
                        c.   Enable LDP-IS-IS synchronization.
                             ldp-sync enable

                 ----End

3.18.3 (Optional) Blocking LDP-IGP Synchronization on an
Interface
Context
                 After LDP-IGP synchronization is enabled in an IGP process, it is enabled on all
                 interfaces whose neighbor status is up on a P2P network, enabled on all interfaces
                 whose neighbor status is up between a DR and a non-DR/BDR on an OSPF-
                 enabled broadcast network, and enabled on all interfaces whose neighbor status
                 is up between a DIS and a non-DIS on an IS-IS-enabled broadcast network.
                 If the interfaces on a device carry key services, ensure that the backup path does
                 not pass through this device. If you do not want to run LDP-IGP synchronization
                 on certain interfaces, you can block the function on these interfaces.

Procedure
                 ●      If OSPF is used, perform the following configuration on the local and remote
                        interfaces of the link between the node where the primary and backup links
                        diverge from each other and the LDP peer on the primary link.
                        a.   Enter the system view.
                             system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             157
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


