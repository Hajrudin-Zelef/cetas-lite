---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-45
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5778, 5955]
sha256: 4ca2f01e02867bf9104f5edff312fedec9890f81debff70a439d8a6936b6735e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  ●      Strict mode
                         In strict mode, the device allows a packet to pass through only when there is
                         a route to the source IP address of the packet in its FIB table and the inbound
                         interface of the packet is the same as the outbound interface of the route.
                         You are advised to use the strict mode when you are sure that the routing
                         paths recorded on the local and remote devices are the same. For example, if
                         there is only one path between two network edge devices, strict mode can be
                         used to ensure network security.
                  ●      Loose mode
                         In loose mode, the device allows a packet to pass through as long as there is
                         a route to the source IP address of the packet in its FIB table. In contrast to
                         strict mode, the inbound interface of the packet does not need to be the
                         same as the outbound interface of the route.
                         You are advised to use the loose mode when the routing paths recorded on
                         the local and remote devices may be different. For example, if there are

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            103
Security Configuration
Security Configuration                                                          7 URPF Configuration


                         multiple paths between two network edge devices, the loose mode can be
                         used to effectively protect the device against network attacks while
                         preventing valid packets from being discarded.


Implementation
                  Figure 7-2 shows how URPF is implemented.


                  Figure 7-2 URPF implementation




7.4 Default Settings for URPF
                  Table 7-1 describes the default settings for URPF.


                  Table 7-1 Default settings for URPF

                   Parameter                                  Default Setting

                   URPF check                                 Disabled

                   URPF check mode                            Loose mode




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         104
Security Configuration
Security Configuration                                                            7 URPF Configuration




7.5 Configuring URPF

7.5.1 Enabling URPF on an Interface

Prerequisites
                  Before configuring URPF, you have completed the following tasks:
                  Configure link layer protocol parameters for interfaces to ensure that the link layer
                  protocol status of the interfaces is up.

Context
                  When configuring URPF, you need to enable the URPF function on an interface.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 3.
                  undo portswitch

                  This step is supported only on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                  S6750E-S, S6750-S, S5755-H, S5755E-H, S5755-S and S5732-H-V2. Determine
                  whether to perform this step based on the current interface working mode.
         Step 4 Enable URPF on the interface.
                  ip urpf enable

                  By default, URPF is disabled on an interface.

                  ----End

7.5.2 Configuring the URPF Check Mode

Context
                  URPF check can be performed in strict or loose mode. You can configure the allow
                  default-route parameter to allow packets to match the default route.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the URPF check mode.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       105
Security Configuration
Security Configuration                                                                     7 URPF Configuration


                  1.     Enter the interface view.
                         interface interface-type interface-number

                  2.     Switch the interface working mode to Layer 3.
                         undo portswitch

                         This step is supported only on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-
                         V2, S6750E-S, S6750-S, S5755-H, S5755E-H, S5755-S and S5732-H-V2.
                         Determine whether to perform this step based on the current interface
                         working mode.
                  3.     Configure the URPF check mode on the interfaces.
                         ip urpf { loose | strict }

                         By default, URPF is disabled on an interface. If URPF is enabled on an
                         interface, the default URPF check mode is loose.
                  4.     Configure the default route to participate in URPF check.
                         ip urpf allow default-route

                         By default, no default route is configured to participate in URPF check.

                  ----End

7.5.3 (Optional) Disabling URPF for a Specified Flow

Context
                  After URPF is configured on an interface, the device performs URPF check on all
                  incoming packets on the interface. To prevent certain packets from being
                  discarded (for example, enable the device to trust all packets from a server and
                  not perform URPF check for such packets), you can disable URPF for the specified
                  flow. The configuration procedure is as follows:

                  1.     Configure a traffic classifier. The traffic classifier defines a group of matching
                         rules to classify packets that do not require URPF check. For details, see
                         "Configuring a Traffic Classifier" under "MQC Configuration" in Configuration
                         Guide > QoS Configuration.
                  2.     Configure a traffic behavior and disable URPF in the traffic behavior. For
                         details, see "Procedure" in this section.
                  3.     Configure a traffic policy, bind the traffic classifier to the traffic behavior, and
                         disable URPF for the classified packets. For details, see "Configuring a Traffic
                         Policy" under "MQC Configuration" in Configuration Guide > QoS
                         Configuration.
                  4.     Apply the traffic policy in the corresponding view as required. For details, see
                         "Applying a Traffic Policy" under "MQC Configuration" in Configuration Guide
                         > QoS Configuration.

                          NOTE

                         Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                         S5755E-H, S5755-S and S5755-H series support the function of disabling URPF for specified
                         flows.


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                 106
Security Configuration
Security Configuration                                                          7 URPF Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create a traffic behavior and enter the traffic behavior view, or enter the view of
                an existing traffic behavior.
                  traffic behavior behavior-name

         Step 3 Disable URPF check for a specified flow.
                  ip urpf disable

