---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-80
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11267, 11414]
sha256: 81db5e8b365683be50e812cec1803e6f2ba05b83e149b6e4167c95ca94b00472
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Run the display mpls ldp interface command to check whether both ends of the
                 LDP session can normally send Hello messages. You are advised to run the display
                 mpls ldp interface command once every three seconds. If the number of sent or
                 received Hello messages remain unchanged after the command is run for several
                 times, the Hello-hold timer expires.

                 ●      If the Hello-hold timer expires, troubleshoot the high CPU usage fault.
                 ●      If the Hello-hold timer does not expire, go to Step 5.

         Step 5 Check whether the LDP Keepalive-hold timer expires.

                 Run the display mpls ldp session command to check whether the two ends of the
                 session can normally send Keepalive messages. You are advised to run the display
                 mpls ldp session command once every five seconds. If the number of sent or
                 received Keepalive messages remain unchanged after the command is run several
                 times, the Keepalive-hold timer expires.

                 ●      If the Keepalive-hold timer expires, troubleshoot the message forwarding
                        fault.
                 ●      If the Keepalive-hold timer does not expire, go to Step 6.

         Step 6 Collect the following information and contact technical support personnel:
                 ●      Results of the preceding steps
                 ●      Configuration file, logs, and alarms of the device

                 ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         190
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


3.26.3 LDP LSP Flapping

Fault Symptom
                 An LDP LSP frequently alternates between up and down.

Possible Causes
                 ●      Route flapping occurs.
                 ●      The LDP session flaps.

Procedure
         Step 1 Check whether route flapping occurs.

                 Run the display ip routing-table command to check for a route to the destination
                 address of the LSP. Run this command every second for 5 to 10 times. If such a
                 route exists, the route information is displayed in the command output. If such a
                 route does not exist, the route information is not displayed in the command
                 output. If the route information is displayed on some occasions but not on others,
                 route flapping occurs.

                 ●      If the route flaps or does not exist, troubleshoot the IGP route fault.
                 ●      If route flapping does not occur, go to Step 2.

         Step 2 Check whether the LDP session flaps.

                 Run the display mpls ldp session command and check the value of the Status
                 field in the command output. Run this command every second for 5 to 10 times. If
                 the value of this field alternates between Operational and other values, the LDP
                 session flaps.

                 ●      If the LDP session flaps, see 3.26.1 LDP Session Flapping to troubleshoot the
                        fault.
                 ●      If the LDP session does not flap, go to Step 3.

         Step 3 Collect the following information and contact technical support personnel:
                 ●      Results of the preceding steps
                 ●      Configuration file, logs, and alarms of the device

                 ----End

3.26.4 LDP LSP Down

Fault Symptom
                 An LDP LSP goes down.

Possible Causes
                 ●      A route does not exist.
                 ●      The LDP session is down.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          191
MPLS Configuration
MPLS Configuration                                                                              3 MPLS LDP Configuration


                 ●      Resources are insufficient. For example, the number of labels reaches the
                        upper limit or the memory is insufficient.
                 ●      The policy for establishing an LSP is configured.

Procedure
         Step 1 Check whether a route exists.

                 Run the display ip routing-table ip-address mask-length verbose command to
                 check for a route to the destination address of the LSP.

                 ip-address mask-length indicates the destination address of the LSP.
                 If routing information is displayed and State is Active Adv, the route for the LSP
                 exists and is active. If the route is a public network BGP route, check whether the
                 route carries a label based on the Label field in the command output. If Label is
                 not NULL, the route carries a label.

                 ●      If the route does not exist, or is inactive, or does not carry a label if it is a BGP
                        route, troubleshoot the IGP route fault.
                 ●      If the route exists, is active, and carries a label if it is a BGP route, go to Step
                        2.

         Step 2 Check whether the LDP session is successfully established.

                 Run the display mpls ldp session command and check the Status field in the
                 command output. If the field value is Operational, the LDP session is established
                 and is up. If the field value is not Operational, the LDP session is not established.

                 ●      If the LDP session fails to be established, see 3.26.2 LDP Session Down to
                        troubleshoot the fault.
                 ●      If the LDP session is established successfully, go to Step 3.

         Step 3 Check whether resources are insufficient, for example, the memory is insufficient
                or the number of LSPs reaches the upper limit.

                 Perform the following checks:

                 1.     Check whether the system memory is insufficient.

                        Run the display health command to check whether the system memory is
                        insufficient. If the system memory is insufficient, delete unnecessary LSPs.
                 2.     Check whether the number of LSPs exceeds the upper limit.

                        Run the display mpls ldp lsp statistics command to check whether the
                        number of LSPs in the system exceeds the upper limit. Unnecessary LSPs can
                        be deleted.

                 If resources are abundant, go to Step 4.

         Step 4 Check whether a policy for establishing LSPs is configured.
                 ●      Run the display this command in the MPLS view and check whether
                        information similar to the following is displayed:
                        lsp-trigger ip-prefix abc (The value varies according to the actual situation.)
                        If so, check whether some LSPs are not included in the IP-prefix-based policy
                        abc.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         192
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration


