---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-32
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "ethernet", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [4326, 4468]
sha256: ab95b808fdc581a511d1d7aaa5436a936cc3a8ff7d66077c8237c735ee2b03e0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

8.2 Understanding Priority Mapping
External Priority Fields
                    Certain fields in the packet header or frame header record QoS information so
                    that network devices can provide differentiated services. These fields include:
                    ●   Precedence field
                        As defined in RFC 791, the 8-bit Type of Service (ToS) field in an IP packet
                        header contains a 3-bit IP precedence field. Figure 8-1 shows the format of
                        the precedence field in an IP packet.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               76
QoS Configuration
QoS Configuration                                                         8 Priority Mapping Configuration


                        Figure 8-1 IP Precedence/DSCP field




                        Bits 0 to 2 constitute the precedence field, representing values 7, 6, 5, 4, 3, 2,
                        1 and 0, listed in descending order of priority. The highest-priority values (7
                        and 6) are reserved for routing and network control communication updates.
                        User-level applications can use only priority values 0 to 5.
                        In addition to the precedence field, a ToS field contains the following sub-
                        fields:
                        –   Bit D indicates the delay. The value 0 represents a normal delay, and
                            value 1 represents a short delay.
                        –   Bit T indicates the throughput. The value 0 represents a normal
                            throughput, and value 1 represents a high throughput.
                        –   Bit R indicates the reliability. The value 0 represents normal reliability,
                            and value 1 represents high reliability.
                    ●   DSCP field
                        The ToS field in IP packets was defined in RFC 1349, with bit C (monetary
                        cost) added. Following that, the IETF DiffServ Working Group redefined bits 0
                        to 5 of a ToS field as the DSCP field in RFC 2474. In RFC 2474, the field name
                        is changed from ToS to differentiated service (DS). Figure 8-1 shows the
                        DSCP field in packets.
                        In the DS field, the first six bits (bits 0 to 5) are the DS CodePoint (DSCP) and
                        the last two bits (bits 6 and 7) are reserved. Additionally, the first three bits
                        (bits 0 to 2) are the Class Selector CodePoint (CSCP), which represents the
                        DSCP type. A DS node selects a PHB based on the DSCP value.
                    ●   802.1p value in an Ethernet frame header
                        Layer 2 devices exchange VLAN frames. As defined in IEEE 802.1Q, the PRI
                        field (802.1p value, also known as CoS) in the Ethernet frame header
                        identifies the QoS requirement. Figure 8-2 shows the format of the 802.1p
                        value in an Ethernet frame header.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 77
QoS Configuration
QoS Configuration                                                        8 Priority Mapping Configuration


                        Figure 8-2 802.1p value in an Ethernet frame header




                        The 802.1Q header contains a 3-bit PRI field and a 1-bit CFI field. The PRI
                        field defines eight service priority values 7, 6, 5, 4, 3, 2, 1, and 0, listed in
                        descending order of priority. The CFI field defines the drop priority of packets.

Implementation
                    In Figure 8-3, voice, video, and data traffic from Host1, Host2, and Host3 traverses
                    DeviceA and DeviceB to reach the network. As voice, video, and data services have
                    priorities in descending order, the device needs to provide differentiated QoS
                    services based on their priorities.
                    Packets carry different precedence fields depending on the network type. For
                    example, packets carry the 802.1p value on a Layer 2 network and the DSCP value
                    on a Layer 3 network. When packets enter a device, the device maps their external
                    priorities to internal priorities and drop priorities, and provides differentiated QoS
                    services for the packets according to their internal priorities and drop priorities.
                    When packets leave the device, the device maps the internal priorities and drop
                    priorities to external priorities so that the network can provide corresponding QoS
                    services based on the external priorities of the packets.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              78
QoS Configuration
QoS Configuration                                                       8 Priority Mapping Configuration


                    Figure 8-3 Implementation of priority mapping




                    The implementation of priority mapping is described as follows:
                    1.   On DeviceA, configure traffic policies in the inbound direction to re-mark
                         priorities of voice, video, and data packets into different 802.1p values.
                    2.   On DeviceB, map 802.1p values to internal and drop priorities in the inbound
                         direction to provide differentiated QoS services for packets based on the
                         internal and drop priorities.
                    3.   On DeviceB, map the internal and drop priorities to the corresponding DSCP
                         values in the outbound direction so that the Layer 3 network can provide
                         differentiated QoS services for the three types of services based on the DSCP
                         values.


8.3 Configuration Precautions for Priority Mapping

8.4 Default Settings for Priority Mapping
                    Table 8-1 describes the default settings for priority mapping.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            79
QoS Configuration
QoS Configuration                                                               8 Priority Mapping Configuration


                    Table 8-1 Default settings for priority mapping

                     Parameter                                       Default Setting

                     Priority trusted on an interface                After packets arrive at an interface:
                                                                     ● Layer 2 forwarding mode
                                                                         If packets carry VLAN tags, the
                                                                         interface trusts 802.1p values;
                                                                         however, if they do not, the
                                                                         interface forwards them based on
                                                                         the default interface priority.
                                                                     ● Layer 3 forwarding mode
                                                                         By default, DSCP values are trusted.

                     Mapping PHBs to DSCP values of                  Disabled
                     outgoing packets on an interface

                     Mapping PHBs to 802.1p values of                Enabled
                     outgoing packets on an interface

