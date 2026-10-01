---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-28
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3210, 3382]
sha256: c7247a4771aee630593e9d057671663f3bbe14f5a9e7182501d996c7e7044871
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

4.4.4 Configuring Traffic Suppression in a VLAN
Context
                  To rate-limit incoming broadcast packets, unknown multicast packets, or unknown
                  unicast packets in a VLAN so as to prevent broadcast storms, you can configure
                  traffic suppression for the corresponding type of packets in the VLAN. Once the
                  rate reaches the configured threshold, the device will then discard excess packets.

Procedure
         Step 1 Enter the system view.
                  system-view


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          55
Security Configuration
Security Configuration                                                             4 Storm Suppression Configuration


         Step 2 Enter the VLAN view.
                  vlan vlan-id

         Step 3 Configure traffic suppression in a VLAN.
                  For the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5755-S, S6750E-S, S6750-
                  S, S5755-H, and S5732-H-V2:
                  storm suppression { broadcast | multicast | unknown-unicast } cir cir-value [ gbps | kbps | mbps ] [ cbs
                  cbs-value [ bytes | kbytes | mbytes ] ]

                  For the S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735R-L-V2, S5735E-L-V2, S5735-
                  L-V2, S5735I-L-V2, S5735I-H-V2, and S5735I-S-V2:
                  storm suppression broadcast cir cir-value [ gbps | kbps | mbps ] [ cbs cbs-value [ bytes | kbytes |
                  mbytes ] ]

                  ----End

4.4.5 Configuring Traffic Suppression in a BD
Context
                  To rate-limit incoming broadcast packets, unknown multicast packets, or unknown
                  unicast packets in a BD so as to prevent broadcast storms, you can configure
                  traffic suppression for the corresponding type of packets in the BD. Once the rate
                  reaches the configured threshold, the device will then discard excess packets.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enter the BD view.
                  bridge-domain bd-id

         Step 3 Configure traffic suppression in a BD.
                  storm suppression { broadcast | multicast | unknown-unicast } cir cir-value [ gbps | kbps | mbps ] [ cbs
                  cbs-value [ bytes | kbytes | mbytes ] ]

                          NOTE

                         This command is supported only on the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2,
                         S6750E-S, S6750-S, S5755-S, S5755E-H, S5755-H, and S5732-H-V2.

                  ----End

4.4.6 Configuring Traffic Suppression in a VSI

Context
                  You can rate-limit broadcast, unknown multicast, or unknown unicast packets in a
                  VSI to prevent broadcast storms. Perform this configuration in the VSI to
                  implement traffic suppression for packets of a specified type. After traffic
                  suppression is configured in a VSI, the device rate-limits the broadcast, unknown
                  multicast, or unknown unicast packets in the VSI and discards excess packets. Only
                  the S6780-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755E-H, S5755-H,
                  S5732-H-V2, and S6750-H support this configuration.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              56
Security Configuration
Security Configuration                                                            4 Storm Suppression Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure MPLS.
                  mpls
                  quit

         Step 3 Configure MPLS L2VPN.
                  mpls l2vpn
                  quit

         Step 4 Enter the VSI view.
                  vsi vsi-id

         Step 5 Configure traffic suppression in the VSI.
                  storm suppression { broadcast | multicast | unknown-unicast } cir cir-value [ gbps | kbps | mbps ] [ cbs
                  cbs-value [ bytes | kbytes | mbytes ] ]

                  ----End

4.4.7 Configuring Traffic Suppression Associated with MAC
Address Flapping
Context
                  If the device enabled with MAC address flapping detection detects MAC address
                  flapping on an interface, traffic suppression is triggered on the interface. When
                  traffic suppression associated with MAC address flapping is configured, the device
                  can use the CIR or the percentage of bandwidth occupied as the traffic
                  suppression threshold and forcibly forward packets based on the threshold.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Configure the threshold for traffic suppression associated with MAC address
                flapping.
                  Configure the threshold for traffic suppression associated with MAC address
                  flapping on interfaces.
                  storm suppression mac-address flapping { percent-value | cir cir-value [ kbps | mbps | gbps ] } [ force ]

                  By default, the threshold for traffic suppression associated with MAC address
                  flapping is the percentage of bandwidth occupied, and its value is 50%.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               57
Security Configuration
Security Configuration                                                          4 Storm Suppression Configuration


                             NOTE

                         When MAC address flapping occurs on an interface configured with traffic suppression:
                         ●     If this command is configured and force is specified, traffic suppression associated
                               with MAC address flapping takes effect.
                         ●     If this command is not configured or force is not specified, traffic suppression takes
                               effect on the interface.
                         Traffic suppression associated with MAC address flapping does not take effect in the
                         following scenarios:
                         ●     If storm control is configured on an interface, traffic suppression associated with MAC
                               address flapping does not take effect on the interface.

                  ----End

4.4.8 Example for Configuring Traffic Suppression on an
Interface in the Inbound Direction
Networking Requirements
                  On the network shown in Figure 4-1, DeviceA connects a Layer 2 network to a
                  Layer 3 device. Traffic suppression needs to be configured on one of DeviceA's
                  interfaces in the inbound direction to rate-limit broadcast packets, unknown
                  multicast packets, and unknown unicast packets forwarded at Layer 2, ultimately
                  preventing broadcast storms.

                  Figure 4-1 Networking diagram of traffic suppression on an interface in the
                  inbound direction
                             NOTE

                         In this example, interface 1 represents 10GE1/0/1.




Procedure
         Step 1 Enter the interface view.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10GE1/0/1] portswitch

         Step 2 Set the CIR of broadcast packets to 100 kbit/s.
                  [DeviceA-10GE1/0/1] storm suppression broadcast cir 100

         Step 3 Set the percentage of bandwidth occupied by unknown multicast packets to 80%.
                  [DeviceA-10GE1/0/1] storm suppression multicast 80

