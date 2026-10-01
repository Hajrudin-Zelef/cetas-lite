---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-79
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11115, 11266]
sha256: 6cadae7eb028b22d1bff43ab0731b1872455188d7d41339b4fb1d045010173ba
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Cleared LDP statistics cannot be restored. Exercise caution when performing this
                 operation.


Procedure
                 ●      To clear statistics about LDP error protocol packets, run the reset mpls ldp
                        error packet { tcp | udp | all } command in the user view.
                 ●      To clear statistics about LDP adjacencies that are down, run the reset mpls
                        ldp event adjacency-down command in the user view.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         187
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


                 ●      To clear statistics about LDP sessions that are down, run the reset mpls ldp
                        event session-down command in the user view.
                 ----End


3.26 Troubleshooting MPLS LDP
3.26.1 LDP Session Flapping
Fault Symptom
                 An LDP session frequently alternates between up and down on an MPLS LDP
                 network.

Possible Causes
                 ●      The LDP GR timer value is added, changed, or deleted.
                 ●      The LDP Keepalive timer value is added, changed, or deleted.
                 ●      A transport address is added, changed, or deleted.
                 ●      The interface alternates between up and down.
                 ●      Route flapping occurs.

Procedure
         Step 1 Check whether LDP GR, the Keepalive timer, or a transport address is configured.
                 1.     Run the display this command in the LDP view to check whether LDP GR is
                        configured.
                        Check whether the command output contains the following information:
                        mpls ldp
                        graceful-restart
                        If so, LDP GR is configured.
                 2.     Run the display this command in the interface view to check whether the
                        LDP Keepalive timer or an LDP transport address is configured.
                        –    Check whether the command output contains information similar to the
                             following:
                             mpls ldp
                             mpls ldp timer keepalive-hold 30
                             If so, the LDP Keepalive timer is configured.
                        –    Check whether the command output contains information similar to the
                             following:
                             mpls ldp
                             mpls ldp transport-address interface
                             If so, an LDP transport address is configured.
                 ●      If any of the preceding items has been configured, wait 10 seconds and then
                        check whether the LDP session alternates between up and down.
                 ●      If none of the preceding items has been configured, go to Step 2.
         Step 2 Check whether the interface alternates between down and up.
                 Run the display ip interface brief command and check the Physical and Protocol
                 fields. If both the Physical and Protocol fields display Up, the interface is up;

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        188
MPLS Configuration
MPLS Configuration                                                             3 MPLS LDP Configuration


                 otherwise, the interface is down. If the field value keeps switching between Up
                 and Down, the interface alternates between up and down.

                 ●      If the interface alternates between up and down, troubleshoot the interface
                        flapping fault.
                 ●      If the interface does not alternate between up and down, go to Step 3.

         Step 3 Check whether route flapping occurs.

                 Run the display ip routing-table command to check route information. You are
                 advised to check the information several times over a short period. If a route is
                 available to the session, the route information is displayed. If no route is available
                 to the session, no related route information is displayed. If the route information is
                 displayed on some occasions but not on others, route flapping occurs.

                 ●      If the route flaps or does not exist, troubleshoot the IGP route fault.
                 ●      If route flapping does not occur, go to Step 4.

         Step 4 Collect the following information and contact technical support personnel:
                 ●      Results of the preceding steps
                 ●      Configuration file, logs, and alarms of the device

                 ----End

3.26.2 LDP Session Down

Fault Symptom
                 An LDP session goes down.

Possible Causes
                 ●      The interface on which the LDP session is established is shut down.
                 ●      The undo mpls, undo mpls ldp, or undo mpls ldp remote peer command is
                        run.
                 ●      A route does not exist.
                 ●      The LDP Keepalive-hold timer expires.
                 ●      The LDP Hello-hold timer expires.

Procedure
         Step 1 Check whether the interface on which the LDP session is established is shut down.

                 Run the display this command in the interface view and check whether the
                 following information is displayed:
                 shutdown

                 The command output shows that the interface is shut down.

                 ●      If the interface is shut down, run the undo shutdown command in the
                        interface view to enable the interface.
                 ●      If the interface is not shut down, go to Step 2.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          189
MPLS Configuration
MPLS Configuration                                                               3 MPLS LDP Configuration


         Step 2 Check whether commands for canceling MPLS configurations are executed.

                 Run the display current-configuration command to check whether commands
                 for canceling MPLS configurations are executed.

                 ●      If the command output does not contain the following information:
                        mpls
                        The MPLS configurations are canceled.
                 ●      If the command output does not contain the following information:
                        mpls ldp
                        The MPLS LDP configurations are canceled.
                 ●      If the command output does not contain the following information:
                        mpls ldp remote-peer
                        Remote LDP session configurations are deleted.
                 ●      If commands for canceling MPLS configurations are executed, run the
                        corresponding commands to restore the configurations.
                 ●      If no commands for canceling MPLS configurations are executed, go to Step 3.

         Step 3 Check whether routes are available.

                 Run the display ip routing-table command to check whether a route to the peer
                 exists based on the Destination/Mask field value. If the route does not exist, a
                 TCP connection cannot be established.

                 ●      If the route does not exist, troubleshoot the IGP route fault.
                 ●      If the route exists, go to Step 4.

         Step 4 Check whether the LDP Hello-hold timer expires.

