---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-43
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5630, 5835]
sha256: 13391ef350bbdca9aca50d617d1e1f697ebe4c2fe5dea2e2a9f1edae4991b534
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack0
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                        #
                        ip ip-prefix prefix1 index 10 permit 1.1.1.9 32
                        #
                        return
                 ●      LSRB
                        #
                        sysname LSRB
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack0
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                        #
                        return
                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        95
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack0
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return

                 ●      DSLAM
                        #
                        sysname DSLAM
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.3.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack0
                         ip address 4.4.4.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                        #
                        return



3.11 Configuring IGP-based LDP Automatic
Deployment
                 IGP-based automatic LDP deployment reduces the configuration workload and
                 ensures configuration correctness.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        96
MPLS Configuration
MPLS Configuration                                                              3 MPLS LDP Configuration


Prerequisites
                 Before configuring IGP-based automatic LDP deployment, complete the following
                 tasks:
                 ●      Configure basic IGP functions.
                 ●      Enable MPLS and MPLS LDP globally.

Context
                 To configure IGP-based MPLS LDP, you need to enable MPLS LDP globally and
                 then enable MPLS LDP on all interfaces that require the function. If a large
                 number of interfaces require the function, this configuration method is time-
                 consuming and prone to configuration errors.
                 To address this issue, configure IGP-based automatic LDP deployment, allowing
                 MPLS LDP to be enabled automatically on IGP-capable interfaces after MPLS LDP
                 is enabled globally.

Procedure
                 ●      Configure IS-IS-based automatic LDP deployment.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the IS-IS view.
                             isis [ process-id ]

                        c.   Configure IS-IS-based automatic LDP deployment.
                             mpls ldp auto-config

                             After this command is run, MPLS LDP is automatically enabled on all
                             interfaces that can establish IS-IS neighbor relationships in the IS-IS
                             process. To disable MPLS LDP on an interface, run the isis mpls ldp auto-
                             config disable command in the interface view.
                 ●      Configure OSPF-based automatic LDP deployment.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the OSPF view.
                             ospf

                        c.   Configure OSPF-based automatic LDP deployment.
                             mpls ldp auto-config

                             After this command is run, MPLS LDP is automatically enabled on all
                             interfaces that can establish OSPF neighbor relationships in the OSPF
                             process. To disable MPLS LDP on an interface, run the ospf mpls ldp
                             auto-config disable command in the interface view.
                 ----End




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                        97
MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration




