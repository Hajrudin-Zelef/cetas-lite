---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-19
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1867, 1983]
sha256: d1de36aef4895a3059edb6835677e08c13cd4a2c257c904d1b1ca33085a6b00a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                         Note the following when configuring an ACL-based whitelist for attack source tracing:
                         ● Before referencing an ACL in a whitelist, create the ACL and configure rules. If an ACL
                           has no rule, the whitelist that references the ACL does not take effect.
                         ● The ACL referenced can be a basic ACL, advanced ACL, Layer 2 ACL, basic ACL6, or
                           advanced ACL6.
                         ● All packets matching an ACL referenced by a whitelist are considered legitimate, and
                           will not be source-traced or punished regardless of whether the ACL rule is permit or
                           deny.
                         ● If an ACL rule is defined by a protocol, ensure that the attack source tracing function
                           supports this protocol.
                         ● Insufficient ACL resources may lead to an ineffective whitelist.

         Step 9 (Optional) Configure the event reporting function for attack source tracing.
                  1.     Enable the event reporting function for attack source tracing.
                         auto-defend alarm enable

                         By default, the event reporting function for attack source tracing is disabled.
                  2.     Configure the rate threshold for triggering event reporting in attack source
                         tracing.
                         auto-defend alarm threshold alarm-threshold


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                   28
Security Configuration
Security Configuration                                                          3 Local Attack Defense Configuration


                         On the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S6750E-S, S6750-S,
                         S5755-S, S5755E-H, S5755-H, and S5732-H-V2, the default rate threshold for
                         reporting attack source tracing events is 128 pps.

                         On the S5735I-L-V2, S5735I-S-V2, S5735I-H-V2, S5735R-L-V2, S5735E-L-V2,
                         S5735-L-V2, S5735R-S-V2, S5735E-S-V2, and S5735-S-V2, the default rate
                         threshold for reporting attack source tracing events is 60 pps.

        Step 10 Configure the punishment action for attack source tracing.
                  auto-defend action { deny [ timeout timeout-num ] | error-down }

                          NOTE

                         ● A device can set the status of an interface to error-down when it detects a fault on the
                           interface. An interface in error-down state cannot receive or send packets and the
                           interface indicator is off.
                            If the punishment action results in the interface that receives attack packets being set to
                            the error-down state, services of authorized users on this interface will be interrupted.
                            Exercise caution when setting this punishment action.
                         ● The device does not take punishment actions on whitelisted users.

        Step 11 Return to the system view.
                  quit

        Step 12 Apply the attack defense policy.
                  ●      Configure attack defense policies in batches.
                         cpu-defend-policy policy-name batch slot { start-slot [ to end-slot ] } &<1-12>

                  ●      Configure an attack defense policy separately.
                         cpu-defend-policy policy-name [ slot slot-id | mcu ]

                  After an attack defense policy is created, you must apply the policy in the system
                  view. Otherwise, the policy does not take effect.

                  ----End


Example
                  In the attack defense policy named test, enable attack source tracing; and set the
                  rate threshold to 200 pps, sampling ratio to 7, attack source tracing mode to
                  source IP address-based, and punishment action to discarding attack packets.
                  <HUAWEI> system-view
                  [HUAWEI] cpu-defend policy test
                  [HUAWEI-cpu-defend-policy-test] auto-defend enable
                  [HUAWEI-cpu-defend-policy-test] auto-defend threshold 200
                  [HUAWEI-cpu-defend-policy-test] auto-defend attack-packet sample 7
                  [HUAWEI-cpu-defend-policy-test] auto-defend trace-type source-ip
                  [HUAWEI-cpu-defend-policy-test] auto-defend action deny
                  [HUAWEI-cpu-defend-policy-test] quit
                  [HUAWEI] cpu-defend-policy test


Follow-up Procedure
                  If the punishment action for attack source tracing is set to error-down, the device
                  sets the status of the interface that receives attack packets to down after
                  identifying the attack source. After the interface is set to down, you are advised to
                  eliminate attacks before recovering the interface to the up state.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        29
Security Configuration
Security Configuration                                               3 Local Attack Defense Configuration


                  Table 3-6 Methods of recovering an interface to the up state
                   Method       Application Scenario               Procedure

                   Manual       ● A small number of                Run the shutdown and undo
                   recovery       interfaces are expected to       shutdown commands, or run the
                                  go down.                         restart command in the interface
                                ● The interface has been set       view to restart the interface.
                                  to down.

                   Automa       ● A large number of                Run the error-down auto-recovery
                   tic            interfaces are expected to       cause auto-defend interval
                   recovery       go down. Manually                command in the system view to
                                  restoring the interface          enable automatic interface recovery
                                  status one by one is time-       after a specified delay. You can run
                                  consuming and may result         the display error-down recovery
                                  in omissions.                    command to view information
                                ● The interface has not been       about automatic interface recovery.
                                  set to down. This mode does
                                  not take effect on interfaces
                                  that are already in down
                                  state.




