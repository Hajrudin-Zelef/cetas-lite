---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-98
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agent", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [12221, 12374]
sha256: 2cbd630686cbbfadfd59c8a8423f72ac2e18d10388cefb8043f54905d5766e09
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  If the device is configured to process PPPoE reply packets from a PPPoE server, the concurrent
                  PPPoE user access rate is reduced.
                  The policy for processing server-side PPPoE packets takes effect only after PPPoE+ is enabled
                  globally. To modify the policy, disable PPPoE+ first.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure whether the device transparently transmits PPPoE reply packets from a
                PPPoE server.
                  pppoe intermediate-agent information ignore-reply disable

                  By default, the device transparently transmits PPPoE reply packets from a PPPoE
                  server without processing them.

                  ----End

11.5.5 Verifying the Configuration

Procedure
         Step 1 Check the globally configured formats of information fields circuit-id and remote-
                id.
                  display pppoe intermediate-agent information format

         Step 2 Check the information fields and vendor ID added to PPPoE packets.
                  display pppoe intermediate-agent information encapsulation

         Step 3 Check the globally configured policies for processing original information fields in
                user-side and server-side PPPoE packets.
                  display pppoe intermediate-agent information policy

         Step 4 Check PPPoE+ configurations.
                  display pppoe intermediate-agent configuration

                  ----End

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        224
Security Configuration
Security Configuration                                                         11 PPPoE+ Configuration


11.5.6 Example for Configuring PPPoE+

Networking Requirements
                  In Figure 11-3, the Device is connected to an upstream BRAS, which functions as a
                  PPPoE server, and to downstream user hosts. Attackers may intercept authorized
                  users' PPPoE packets and compromise their accounts. The administrator wants to
                  prevent these problems and ensure user account security.

                  Figure 11-3 Network diagram of configuring PPPoE+




Configuration Roadmap
                  The configuration roadmap is as follows:

                  1.     Enable PPPoE+ globally to authenticate user accounts together with the
                         access interfaces, preventing user accounts from being compromised.
                  2.     Configure the interface connecting the Device to the PPPoE server as a trusted
                         interface, preventing PPPoE packets from being forwarded to a non-PPPoE
                         service interface.
                  3.     Configure the policy for processing original information fields in user-side
                         PPPoE packets on the Device, enabling the Device to communicate correctly
                         with the PPPoE server.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           225
Security Configuration
Security Configuration                                                                 11 PPPoE+ Configuration


Procedure
         Step 1 Enable PPPoE+.
                  <HUAWEI> system-view
                  [HUAWEI] pppoe intermediate-agent information enable

                          NOTE

                         When enabled globally, PPPoE+ is enabled on all interfaces.

         Step 2 Configure the interface as a trusted interface.
                  <HUAWEI> interface 10GE1/0/1
                  [HUAWEI-10GE1/0/1] pppoe uplink-port trusted
                  [HUAWEI-10GE1/0/1] quit

         Step 3 Configure the policy for processing original information fields in user-side PPPoE
                packets to replace on all interfaces. This ensures that the Device replaces original
                information fields in user-side PPPoE packets with fields circuit-id and remote-id
                of the Device.
                  [HUAWEI] pppoe intermediate-agent information policy replace

         Step 4 Set the format of circuit-id to extend.
                  [HUAWEI] pppoe intermediate-agent information format circuit-id extend

         Step 5 Verify the configuration.
                  # Run the display pppoe intermediate-agent information policy command to
                  verify the policy for processing original information fields in user-side PPPoE
                  packets.
                  [HUAWEI] display pppoe intermediate-agent information policy
                   The current information Policy :REPLACE
                   The current ignore-reply Policy:ENABLE

                  # Run the display pppoe intermediate-agent information format command to
                  verify the format of circuit-id.
                  [HUAWEI] display pppoe intermediate-agent information format
                   The current information format :
                    Circuit ID : EXTEND
                    Remote ID : COMMON
                   For example:
                    interface 10GE1/0/1 SVLAN:200 CVLAN:100
                    The PPPOE Intermediate Agent information follow:
                    Circuit ID:00 04 00 c8 00 00
                    Remote ID:0025-9efb-494a

                  ----End

Configuration Scripts
                  Device
                  #
                  pppoe intermediate-agent information enable
                  pppoe intermediate-agent information format circuit-id extend
                  #
                  interface 10GE1/0/1
                   pppoe uplink-port trusted
                  #
                  return


11.5.7 Troubleshooting PPPoE+

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               226
Security Configuration
Security Configuration                                                             11 PPPoE+ Configuration


11.5.7.1 PPPoE Users Cannot Go Online After PPPoE+ Is Configured

Fault Symptom
                  PPPoE users cannot go online after PPPoE+ is configured.

Possible Causes
                  Possible causes are as follows:
                  ●      The network-side interface is not configured as a trusted interface.
                  ●      The policy for processing original information fields in user-side PPPoE
                         packets does not meet service requirements.
                  ●      The formats of the information fields added to PPPoE packets are different
                         from those required by the PPPoE server.

