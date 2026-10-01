---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-31
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [4185, 4325]
sha256: d3cfcee927087528396e3a756d9fb9a403148e1b7fa9802e24e122ee51fd6493
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    <DeviceA> display traffic-policy applied-record
                    Total records : 1
                    --------------------------------------------------------------------------------
                    Policy Type/Name                     Apply Parameter              Slot State
                    --------------------------------------------------------------------------------
                    p1                             Global(IN)               1     success
                    --------------------------------------------------------------------------------


Configuration Scripts
                    DeviceA
                    #
                    sysname DeviceA
                    #
                    vlan batch 10 20 30
                    #
                    traffic-policy p1 global inbound
                    #
                    acl number 3001
                     rule 5 permit ip destination 10.1.4.0 0.0.0.255
                    #
                    traffic classifier c1 type or
                     if-match acl 3001
                    #
                    traffic behavior b1
                     redirect nexthop 10.1.2.2 low-precedence
                    #
                    traffic policy p1
                     classifier c1 behavior b1 precedence 5
                    #
                    interface Vlanif10
                     ip address 10.1.1.1 255.255.255.0
                    #
                    interface Vlanif20
                     ip address 10.1.2.1 255.255.255.0
                    #
                    interface Vlanif30
                     ip address 10.1.3.1 255.255.255.0
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 10
                    #
                    interface 10GE1/0/2
                     port link-type trunk
                     port trunk allow-pass vlan 20
                    #
                    interface 10GE1/0/3
                     port link-type trunk
                     port trunk allow-pass vlan 30
                    #
                    ip route-static 10.1.4.0 255.255.255.0 10.1.1.2
                    ip route-static 0.0.0.0 0 10.1.3.2
                    #
                    return




Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                                74
QoS Configuration
QoS Configuration                                                         8 Priority Mapping Configuration




                        8        Priority Mapping Configuration


                    8.1 Overview of Priority Mapping
                    8.2 Understanding Priority Mapping
                    8.3 Configuration Precautions for Priority Mapping
                    8.4 Default Settings for Priority Mapping
                    8.5 Configuring DiffServ Domain-based Priority Mapping
                    8.6 Configuring an Interface Priority
                    8.7 Configuring the Internal Priority of Protocol Packets Sent by the Local Device
                    8.8 Configuring the Mapping Between Internal Priorities and Queues
                    8.9 Configuring a Packet Priority
                    8.10 Troubleshooting Priority Mapping


8.1 Overview of Priority Mapping
Definition
                    Priority mapping maps external priorities carried in packets into internal priorities
                    on a device, and manages and records the mapping between external and internal
                    priorities using DiffServ domains. In this way, the device provides differentiated
                    services for packets based on internal priorities.

                    Common concepts of priorities are described as follows:

                    ●   External priority
                        The external priority is also known as the packet priority or QoS priority. QoS
                        information is recorded by using certain fields in packets, such as the 802.1p
                        value of VLAN packets and the Differentiated Services Code Point (DSCP)
                        value of IP packets. A device can process received packets only based on
                        internal priorities to provide differentiated QoS levels for different services. As
                        such, external priorities are mapped to internal priorities after packets enter
                        the device.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                75
QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                    ●   Internal priority
                        The internal priority is also known as the class of service (CoS), per-hop
                        behavior (PHB), or local priority. Internal priority values are CS7, CS6, EF, AF4,
                        AF3, AF2, AF1, and BE (in descending order of priority), and correspond to
                        queues 7, 6, 5, 4, 3, 2, 1, and 0, respectively. The internal priority determines
                        the queue into which packets are placed. When QoS services are configured
                        for a queue, the same QoS level is configured for all the packets forwarded
                        through that queue.
                    ●   Drop priority
                        The drop priority is also known as a color. It determines the sequence in
                        which packets are dropped when congestion occurs in a queue, without
                        affecting the mapping between internal priorities and queues. The drop
                        priorities defined by Institute of Electrical and Electronics Engineers (IEEE) are
                        green, yellow, and red in ascending order. By default, the device first discards
                        packets with a higher drop priority when congestion occurs in a queue.
                        Whether packets are discarded first is determined by QoS parameter settings.
                        For example, in a weighted random early detection (WRED) drop profile, if
                        green packets can use a maximum of 50% of the buffer and red packets can
                        use the entire buffer, the device first discards green packets when congestion
                        occurs in a queue.

Purpose
                    Packets transmitted over different networks carry different external priority fields.
                    For example, the 802.1p value is used on a VLAN, and the DSCP value is used on
                    an IP network. The device processes packets transmitted on the network as
                    follows:
                    ●   For all incoming packets, the device maps external priorities (including 802.1p
                        and DSCP values) to internal priorities, places the packets into specific queues
                        based on the mapping, and applies traffic shaping, congestion avoidance, or
                        congestion management to the queues.
                    ●   For outgoing packets, the device maps the internal priority to an external
                        priority, and other devices can provide QoS services based on the external
                        priority.


