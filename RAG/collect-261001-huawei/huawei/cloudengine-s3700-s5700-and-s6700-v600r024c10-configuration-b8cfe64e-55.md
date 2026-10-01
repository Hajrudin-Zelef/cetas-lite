---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-55
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [7634, 7774]
sha256: 66b345efbad70f270c2080b96055ed709090d42d2b16d0852843974f1eeea2fc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Configuration Scripts
                    ●    DeviceB
                         #
                         sysname DeviceB
                         #
                         vlan batch 10
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10
                          trust 8021p outer
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 10
                          qos lr cir 10000 outbound
                          qos queue 2 shaping cir 2000 kbps pir 3000 kbps
                          qos queue 5 shaping cir 5000 kbps pir 8000 kbps
                          qos queue 6 shaping cir 3000 kbps pir 5000 kbps
                         #
                         return



9.7 Configuring Interface-based Rate Limiting
Context
                    Interface-based rate limiting controls the total rate of all packets passing through
                    an interface to ensure that the bandwidth utilization of the interface remains
                    within the allowed range. You can configure interface-based rate limiting in either
                    the inbound or outbound direction, or both.

9.7.1 Configuring Traffic Policing to Implement Interface-
based Rate Limiting
Context
                    If the rate of traffic sent from users is not limited, burst data continuously sent by
                    users may cause the network to become congested. Therefore, to limit the rate of
                    traffic entering an interface, traffic policing can be configured in the inbound
                    direction of the interface on a network device. When the rate of received packets
                    exceeds the specified traffic policing rate, the device discards excess packets.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 (Optional) Disable the device from counting the inter-frame gaps and preambles
                when the device calculates the traffic policing rate.
                    qos car ifg disable


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                             136
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration


         Step 3 Create and configure a CAR profile.
                    qos car car-name { percent percent-value | cir cir-value [ kbps | mbps | gbps ] [ cbs cbs-value [ bytes |
                    kbytes | mbytes ] [ pbs pbs-value [ bytes | kbytes | mbytes ] ] | pir pir-value [ kbps | mbps | gbps ] [ cbs
                    cbs-value [ bytes | kbytes | mbytes ] pbs pbs-value [ bytes | kbytes | mbytes ] ] ] }

                    The CAR profile defines the rate limit of traffic policing.

         Step 4 Enter the interface view.
                    interface { interface-type interface-number | interface-name }

         Step 5 Apply the CAR profile to the inbound direction of the interface.
                    qos car inbound car-name

                    After the CAR profile is applied to the inbound direction of the interface, the
                    device rate-limits all incoming service traffic on the interface.

                    ----End


Verifying the Configuration
                    ●    Run the display qos car [ name car-name ] command to check the CAR
                         profile configuration.
                    ●    Run the display qos car statistics interface { interface-type interface-number
                         | interface-name } inbound command to check statistics about packets
                         forwarded and discarded on a specified interface on which traffic policing is
                         configured.

9.7.2 Configuring Traffic Shaping to Implement Interface-
based Rate Limiting

Context
                    When a large amount of traffic is sent from the downstream device to its
                    upstream device, to prevent congestion or packet loss, configure traffic shaping on
                    the outbound interface of the device. Traffic shaping adjusts the rate of outgoing
                    traffic on the interface to reduce traffic bursts so that outgoing packets can be
                    transmitted at a stable rate. With traffic shaping, packets exceeding the rate limit
                    enter the buffer queue, and are only sent out at an even rate when there are
                    sufficient tokens in the token bucket. In addition, if the buffer queue is full, the
                    system discards new packets.

                    If packets are encapsulated through VPN tunnels, the calculated packet length
                    during rate limiting or traffic statistics collection on a WAN-side interface includes
                    the lengths of link-layer and physical-layer packet headers and the lengths of
                    tunnel protocol encapsulation headers. As a result, the actual user bandwidth
                    calculated using this method is smaller than the expected value. To minimize the
                    difference between the actual bandwidth and the expected bandwidth, you can
                    disable the device from counting the inter-frame gaps and preambles when it
                    calculates the packet length, and configure the function of excluding the tunnel
                    encapsulation length of packets on the tunnel interface. In this way, the calculated
                    packet length excludes the lengths of the physical-layer packet header and tunnel
                    protocol encapsulation header.


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                137
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration


                          NOTE

                         Currently, traffic shaping can only be used to rate-limit outgoing traffic on an interface.
                         The system will generate an alarm when the rate of outgoing traffic on an interface
                         exceeds the alarm threshold for rate-limiting outgoing traffic.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the interface view.
                    interface { interface-type interface-number | interface-name }

         Step 3 Configure the rate limit for outgoing traffic on the interface.
                    qos lr cir cir-value [ kbps | mbps | gbps ] [ cbs cbs-value [ bytes | kbytes | mbytes ] ] [ outbound ]

                    By default, the rate limit on an interface is the maximum bandwidth of the
                    interface.

                          NOTE

                         Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                         S5755E-H, S5755-S and S5755-H series support the mbytes parameter.

                    ----End

