---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-27
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3064, 3209]
sha256: fd95ea872d28022834039f832181ba93be3092a506e73ad34de46a56ee92e924
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   Percentage                                 ● Traffic suppression for broadcast
                                                                packets: 10%
                                                              ● Traffic suppression for unknown
                                                                multicast or unicast packets: 100%
                                                              ● Traffic suppression associated with
                                                                MAC address flapping for unknown
                                                                unicast packets: 50%

                   Traffic suppression in the outbound        Disabled
                   direction of an interface

                   Traffic suppression in a VLAN or BD        Disabled

                   Traffic suppression for Internet Control   Enabled
                   Message Protocol (ICMP) packets

                   Traffic suppression threshold for ICMP     1500 pps
                   messages

                   Traffic suppression associated with        Enabled
                   MAC address flapping

                   Mode of traffic suppression associated     Percentage of bandwidth occupied
                   with MAC address flapping

                   Storm control                              Disabled

                   Log recording and alarm report             Disabled

                   Storm detection interval                   5s




4.4 Configuring Traffic Suppression
4.4.1 Understanding Traffic Suppression
                  Traffic suppression limits broadcast packets, unknown multicast packets, or
                  unknown unicast packets in the following modes:
                  ●      In the inbound direction of an interface, the device supports traffic
                         suppression for broadcast packets, unknown multicast packets, or unknown
                         unicast packets based on the percentage of bandwidth occupied or packet
                         rate.
                         When the rate of incoming packets reaches the threshold, the device discards
                         excess packets.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            53
Security Configuration
Security Configuration                                                            4 Storm Suppression Configuration


                  ●      In the outbound direction of an interface, the device can block broadcast
                         packets, unknown multicast packets, or unknown unicast packets.
                  ●      In the VLAN or BD view, the device supports traffic suppression for broadcast
                         packets, unknown multicast packets, or unknown unicast packets based on
                         the bit rate.
                         The device monitors the rates of various types of packets in the same VLAN or
                         BD and discards excess packets when the traffic rate configured for the VLAN
                         or BD exceeds the threshold.

                  In addition, the device supports the following traffic suppression functions:
                  ●      Traffic suppression for ICMP messages: You can set a rate limit for ICMP
                         messages to prevent a large number of ICMP messages from being sent to
                         the CPU, as this may affect other service functions.
                  ●      Traffic suppression associated with MAC address flapping: If the device
                         enabled with MAC address flapping detection detects MAC address flapping
                         on an interface, traffic suppression is triggered on the interface.

4.4.2 Configuring Traffic Suppression on an Interface in the
Inbound Direction

Context
                  To prevent broadcast storms, you can configure traffic suppression on an interface
                  in the inbound direction. The device supports traffic suppression for broadcast
                  packets, unknown multicast packets, or unknown unicast packets by percentage of
                  bandwidth occupied or packet rate. When the traffic volume of any of these
                  packet types exceeds the threshold, the system discards excess packets to reduce
                  the traffic volume to within an appropriate range.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 2.
                  portswitch

                  This step is supported only on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                  S6750E-S, S6750-S, S5755-H, S5755E-H, S5755-S and S5732-H-V2. Determine
                  whether to perform this step based on the current interface working mode.

         Step 4 Configure traffic suppression on an interface in the inbound direction.
                  storm suppression { broadcast | multicast | unknown-unicast } { percent-value | cir cir-value [ gbps |
                  kbps | mbps ] [ cbs cbs-value [ bytes | kbytes | mbytes ] ] | packets packets-per-second }

                  If traffic suppression is configured for packets of the same type on an interface in
                  the inbound direction for multiple times and parameters of percent-value and cir
                  cir-value are specified, only the latest configuration takes effect.

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                54
Security Configuration
Security Configuration                                                          4 Storm Suppression Configuration


Verifying the Configuration
                  Run the display storm suppression { broadcast | multicast | unknown-unicast }
                  [ interface interface-type interface-number ] command to check the configured
                  and actual traffic suppression thresholds on an interface in the inbound direction.

                          NOTE

                         The rate limit threshold of traffic suppression and the actual rate limit may differ, in which
                         case the actual rate limit of packets is used.


4.4.3 Configuring Traffic Suppression on an Interface in the
Outbound Direction
Context
                  If some interfaces do not need to receive any broadcast packets, unknown
                  multicast packets, or unknown unicast packets (for example, the interfaces are
                  connected to fixed user hosts and demand high security), configure traffic
                  suppression on the interfaces in the outbound direction to block those packets.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the interface view.
                  interface interface-type interface-number

         Step 3 Switch the interface working mode to Layer 2.
                  portswitch

                  This step is supported only on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                  S6750E-S, S6750-S, S5755-H, S5755E-H, S5755-S and S5732-H-V2. Determine
                  whether to perform this step based on the current interface working mode.
         Step 4 Configure traffic suppression on an interface in the outbound direction.
                  storm suppression { broadcast | multicast | unknown-unicast } block outbound

                  ----End

