---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-31
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [3648, 3799]
sha256: 8512e7f80238ef5cbc395011ccd38c9bdeae39bfb9a1b368064255b5ac30e463
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Configuration Scripts
                  DeviceA
                  #
                  sysname DeviceA
                  #
                  interface 10GE1/0/1
                   storm control broadcast min-rate 1000 max-rate 2000
                   storm control multicast min-rate 1000 max-rate 2000
                   storm control unknown-unicast min-rate 1000 max-rate 2000
                   storm control interval 90
                   storm control action block
                   storm control enable log
                  #
                  return



4.6 Troubleshooting Storm Suppression

4.6.1 Traffic Suppression Has No Effect in the Inbound
Direction of an Interface
Fault Symptom
                  After traffic suppression for broadcast packets, unknown multicast packets, or
                  unknown unicast packets is configured on an interface, broadcast storms caused
                  by packets of the corresponding type still occur, interrupting the flow of traffic.

Possible Causes
                  ●      Traffic suppression is not configured for the corresponding type of packets on
                         the interface or the configured traffic suppression threshold is high.
                  ●      Packets of the corresponding type are not discarded in the inbound direction
                         of the interface.

Procedure
         Step 1 Check whether traffic suppression is correctly configured in the inbound direction
                of the interface.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            63
Security Configuration
Security Configuration                                                 4 Storm Suppression Configuration


                  ●      Run the display storm suppression { broadcast | multicast | unknown-
                         unicast } [ interface interface-type interface-number ] command in any view
                         to check traffic suppression information, or run the display this command in
                         the interface view to check the traffic suppression configuration of the
                         interface.
                         a.   Check whether traffic suppression is configured for the corresponding
                              type of packets.
                         b.   Check whether the traffic suppression threshold is high.

                              ▪   If the traffic suppression threshold is high, run the storm
                                  suppression { broadcast | multicast | unknown-unicast } { percent-
                                  value | cir cir-value [ gbps | kbps | mbps ] [ cbs cbs-value [ bytes |
                                  kbytes | mbytes ] ] | packets packets-per-second } command in the
                                  interface view to modify traffic suppression parameters.

                              ▪   If the traffic suppression threshold is within a reasonable range, go to
                                  the next step.
         Step 2 Check whether packets are discarded in the inbound direction of the interface
                using either of the following methods:
                  ●      Run the display interface interface-type interface-number command in the
                         user view to check whether the value of output bandwidth utilization changes
                         significantly after traffic suppression is configured. Normally, after traffic
                         suppression is configured, the bandwidth utilization on an interface decreases
                         if the interface discards excess packets once the threshold is reached. If the
                         bandwidth utilization does not change or changes only slightly, go to the next
                         step.
                  ●      Add another interface (interface B) to the same VLAN as the interface that is
                         configured with traffic suppression (interface A). Then check whether the
                         volume of the outgoing traffic on interface B is the same as the volume of the
                         traffic on interface A. If they differ, packets are not discarded in the inbound
                         direction of interface A. In this case, go to the next step.
         Step 3 Collect the following information and contact technical support personnel:
                  ●      Results of the preceding troubleshooting procedure
                  ●      Configuration, log, and trap information

                  ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            64
Security Configuration
Security Configuration                                                            5 IPSG Configuration




                                                  5        IPSG Configuration


                  5.1 Overview of IPSG
                  5.2 Understanding IPSG
                  5.3 Configuration Precautions for IPSG
                  5.4 Default Settings for IPSG
                  5.5 Configuring IPSG Based on a Static Binding Table
                  5.6 Configuring IPSG Based on a Dynamic Binding Table
                  5.7 (Optional) Configuring the IP Packet Check Alarm Function
                  5.8 (Optional) Configuring the Function of Discarding IP Packets with Identical
                  Source and Destination IP Addresses
                  5.9 Maintaining IPSG
                  5.10 Troubleshooting IPSG


5.1 Overview of IPSG
Definition
                  IP Source Guard (IPSG) implements source IP address filtering based on Layer 2
                  interfaces to prevent unauthorized hosts from using IP addresses of authorized
                  hosts or specified IP addresses to access or attack the network.

Purpose
                  As the network scale continues to grow, many attackers are forging source IP
                  addresses to initiate IP address spoofing attacks. In order to obtain network access
                  rights and access networks, attackers forge the IP addresses of authorized users.
                  As a result, authorized users cannot access networks and sensitive information
                  may be leaked. IPSG provides a mechanism to effectively defend against IP
                  address spoofing attacks.
                  Figure 5-1 illustrates how IPSG defends against attacks from an unauthorized
                  host, which obtains network access rights by forging an authorized host's IP

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             65
Security Configuration
Security Configuration                                                           5 IPSG Configuration


                  address. IPSG is configured on the Device's user-side interface or VLAN. With this
                  configuration, the Device checks the IP packets received by the interface and
                  discards the packets from unauthorized hosts, preventing IP address spoofing
                  attacks.

                  Figure 5-1 Typical implementation of IPSG




Benefits
                  ●      IPSG reduces costs for ensuring normal network operations and information
                         security.
                  ●      IPSG provides secure network environments and more stable network services.


5.2 Understanding IPSG
IPSG Fundamentals
                  IPSG checks IP packets on Layer 2 interfaces against a binding table that contains
                  the bindings of source IP addresses, source MAC addresses, VLANs, masks, and
                  inbound interfaces. Only packets matching binding entries are forwarded, and
                  other packets are discarded.
                  Table 5-1 describes two types of binding tables: static and dynamic binding tables.

