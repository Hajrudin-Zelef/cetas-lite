---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-29
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3383, 3510]
sha256: 64ad70a02a7c29adef1422b5cdc30971c0ee4df06d091b56af21f5c02f612028
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 4 Set the CIR of unknown unicast packets to 100 kbit/s.
                  [DeviceA-10GE1/0/1] storm suppression unknown-unicast cir 100
                  [DeviceA-10GE1/0/1] quit

                  ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                           58
Security Configuration
Security Configuration                                                                     4 Storm Suppression Configuration


Verifying the Configuration
                  # Check the traffic suppression configuration in the inbound direction of the
                  interface.
                  [DeviceA] display storm suppression broadcast interface 10ge 1/0/1
                  ------------------------------------------------------------------------------------------------
                                        Configured                        Current
                  interface       percent(%) cir(kbps) cbs(bytes)          pps percent(%) cir(kbps) cbs(bytes)       pps
                  ------------------------------------------------------------------------------------------------
                  10GE1/0/1                --      100      18800      --        --     100       18800      --
                  ------------------------------------------------------------------------------------------------
                  [DeviceA] display storm suppression multicast interface 10ge 1/0/1
                  ------------------------------------------------------------------------------------------------
                                        Configured                        Current
                  interface       percent(%) cir(kbps) cbs(bytes)          pps percent(%) cir(kbps) cbs(bytes)       pps
                  ------------------------------------------------------------------------------------------------
                  10GE1/0/1                80       --       --     --       80       --        --    --
                  ------------------------------------------------------------------------------------------------
                  [DeviceA] display storm suppression unknown-unicast interface 10ge 1/0/1
                  ------------------------------------------------------------------------------------------------
                                        Configured                        Current
                  interface       percent(%) cir(kbps) cbs(bytes)          pps percent(%) cir(kbps) cbs(bytes)       pps
                  ------------------------------------------------------------------------------------------------
                  10GE1/0/1                --      100      18800      --        --     100       18800      --
                  ------------------------------------------------------------------------------------------------

                  The Configured field displays the configured traffic suppression percentage, CIR,
                  and committed burst size (CBS). The Current field displays the effective traffic
                  suppression percentage, CIR, and CBS. The preceding command output shows that
                  the maximum rate of broadcast packets is 100 kbit/s, the percentage of interface
                  bandwidth occupied by unknown multicast packets is 80%, and the maximum rate
                  of unknown unicast packets is 100 kbit/s on 10GE1/0/1 of DeviceA in the inbound
                  direction.

Configuration Scripts
                  DeviceA
                  #
                  sysname DeviceA
                  #
                  interface 10GE1/0/1
                   storm suppression broadcast cir 100 kbps
                   storm suppression multicast 80
                   storm suppression unknown-unicast cir 100 kbps
                  #
                  return



4.5 Configuring Storm Control

4.5.1 Understanding Storm Control
                  Storm control prevents broadcast storms caused by broadcast packets, unknown
                  multicast packets, and unknown unicast packets.
                  In a detection interval, the device monitors the average rate of these types of
                  packets received on an interface and compares it with the upper and lower
                  thresholds. If the average rate is greater than the upper threshold, the device
                  performs storm control on the interface and takes the configured storm control

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              59
Security Configuration
Security Configuration                                                          4 Storm Suppression Configuration


                  action. Then, if the average rate falls below the lower threshold, the interface
                  starts to forward packets again.

                  Storm control actions include shutting down an interface, blocking packets, and
                  suppressing packets.
                  ●      If the action is to shut down an interface, you need to manually unblock the
                         interface or enable the interface to automatically return to the up state.
                  ●      If the action is to block packets, when the average rates of incoming packets
                         on a blocked interface fall below the lower thresholds, the interface starts
                         forwarding packets again.
                  ●      If the action is to suppress packets, when the average rate of packets received
                         on the interface exceeds the configured upper threshold, the system discards
                         the excess traffic until the average rate of the packets falls below the
                         threshold.

4.5.2 Configuring Storm Control

Context
                  To rate-limit broadcast packets, unknown multicast packets, or unknown unicast
                  packets on an interface so as to prevent broadcast storms, configure storm control
                  for these types of packets on the interface.

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

         Step 4 Configure the upper and lower thresholds for broadcast packets, unknown
                multicast packets, or unknown unicast packets on the interface. The device
                supports the following configuration modes based on the measurement units of
                the upper and lower thresholds:
                  ●      Specify the lower threshold min-rate-value and upper threshold max-rate-
                         value, which is expressed in pps.
                         storm control { broadcast | multicast | unknown-unicast } min-rate min-rate-value max-rate max-
                         rate-value
                  ●      Specify the lower threshold min-rate-value-kbps and upper threshold max-
                         rate-value-kbps, which is expressed in kbit/s.
                         storm control { broadcast | multicast | unknown-unicast } min-rate kbps min-rate-value-kbps max-
                         rate kbps max-rate-value-kbps

                  ●      Specify the lower threshold min-rate-value-percent and upper threshold max-
                         rate-value-percent. The threshold is expressed in percentage, that is, the
                         percentage of the interface bandwidth occupied by packets.

