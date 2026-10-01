---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-72
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10087, 10298]
sha256: a2b5b494cc68d2770b370e888ebedf4096c414803cb8c8e49519a38ad9ae5ecc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls ldp
                         graceful-restart
                         #
                          ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.3
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
                         graceful-restart
                         #
                          ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            170
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.252
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
                         graceful-restart
                         #
                          ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls ldp
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



3.20 Configuring the Uniform/Pipe Mode for the MPLS
Penultimate Hop
Prerequisites
                 Before configuring the uniform/pipe mode for the MPLS penultimate hop, you
                 have completed the following task:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       171
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                 ●      Configure the physical parameters and link attributes of interfaces to ensure
                        that they work properly.

Context
                 You can configure the uniform/pipe mode on the penultimate hop of an MPLS LSP
                 to determine whether to copy the EXP value of an outer label to the EXP value of
                 an inner label.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure the MPLS DiffServ mode of the device.
                 mpls lsp exp-mode { pipe | uniform }

                        NOTE

                 ● The mode configured using this command takes effect only on new LSPs. The reset mpls ldp
                   command can be used to reestablish the original LSPs. Then, the configured mode can take
                   effect on the re-established LSPs.
                 ● The command is run only on the penultimate hop to determine whether to copy the EXP
                   value of an outer label to the EXP value of an inner label.

                 ----End


3.21 Disabling LDP Session Flapping Suppression
Context
                 If an LDP session goes down due to a protocol or interface fault, LDP immediately
                 attempts to reestablish the LDP session to ensure the fastest LDP hard
                 convergence. For an LDP session that alternates between up and down multiple
                 times within a period of time, the involved upstream and downstream LSPs are
                 frequently created and deleted, wasting resources. To prevent this problem, LDP
                 session flapping suppression is enabled by default. For a stable LDP network, you
                 can disable LDP session flapping suppression.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enter the MPLS-LDP view.
                 mpls ldp

         Step 4 Disable LDP session flapping suppression.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                172
MPLS Configuration
MPLS Configuration                                                          3 MPLS LDP Configuration

                 session suppress disable

                 ----End


3.22 Disabling LDP Interface Flapping Suppression
Context
                 If LDP interface flapping suppression is disabled and an interface frequently flaps,
                 LDP frequently sends Address and Address Withdraw messages to all LDP sessions.
                 If there are a large number of sessions, the CPU usage of the device increases,
                 causing protocol flapping. To prevent this problem, LDP interface flapping
                 suppression is enabled by default. For a stable LDP network, you can disable LDP
                 interface flapping suppression.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enter the MPLS-LDP view.
                 mpls ldp

