---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-19
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "ethernet", "throughput"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [2282, 2417]
sha256: b960554f4f12b226bfc4a876995482b84556e8fcd6c14462cb9d49b5ecc1f0aa
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                    6.1 Overview of Re-marking
                    6.2 Re-marking
                    6.3 Configuration Precautions for Re-marking
                    6.4 Configuring MQC-based Priority Re-marking
                    6.5 Example for Configuring Re-marking to Distinguish Users
                    6.6 Example for Configuring Re-marking to Distinguish Services


6.1 Overview of Re-marking
                    Re-marking technology allows the device to re-mark a field (such as the priority)
                    of packets so that the packets can be scheduled or forwarded based on the re-
                    marked information. For example, the device re-marks the priority of packets
                    matching traffic classification rules. Packets of services that are sensitive to delay
                    and service quality can be re-marked with a high priority so that they can be
                    preferentially scheduled or forwarded. Similarly, priorities of services with no
                    special requirements on delay or service quality can be reduced to save network
                    resources for high-priority packets. For example, re-marking the priority of VLAN
                    packets is to re-mark the 802.1p value of the packets. The packets then can be
                    scheduled and forwarded based on the new 802.1p value.


6.2 Re-marking
Precedence Fields
                    In order for devices to provide differentiated services, certain fields in the packet
                    header or frame header are used to record QoS information. These fields include:
                    ●   802.1p value in an Ethernet frame header
                        Layer 2 devices communicate by exchanging Ethernet frames. In the Ethernet
                        frame header, the PRI field (802.1p value, also called CoS) is used to identify
                        the QoS requirements, as defined in IEEE 802.1Q. Figure 6-1 shows the
                        format of the 802.1p value in an Ethernet frame header.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               40
QoS Configuration
QoS Configuration                                                                6 Re-marking Configuration


                        Figure 6-1 802.1p value in an Ethernet frame header




                        As shown in the figure, the 802.1Q header contains a 3-bit PRI field and a 1-
                        bit CFI field. The PRI field defines eight service priority values 7, 6, 5, 4, 3, 2, 1,
                        and 0, listed in descending order of priority. The CFI field defines the drop
                        priority of packets.
                    ●   Precedence field in an IP packet
                        As defined in RFC 791, the 8-bit Type of Service (ToS) field in an IP packet
                        header contains a 3-bit IP precedence field. Figure 6-2 shows the format of
                        the precedence field in an IP packet.

                        Figure 6-2 IP precedence in the DSCP field




                        Bits 0 to 2 represent the precedence field, representing values 7, 6, 5, 4, 3, 2, 1
                        and 0, listed in descending order of priority. The highest-priority values (7 and
                        6) are reserved for routing and network control communication updates.
                        User-level applications can use only priority values 0 to 5.
                        In addition to the precedence field, a ToS field contains the following sub-
                        fields:
                        –   Bit D indicates the delay. The value 0 represents a normal delay and the
                            value 1 represents a short delay.
                        –   Bit T indicates the throughput. The value 0 represents a normal
                            throughput and the value 1 represents a high throughput.
                        –   Bit R indicates the reliability. The value 0 represents normal reliability and
                            the value 1 represents high reliability.
                    ●   DSCP field in an IP packet
                        The ToS field in IP packets was defined in RFC 1349, with bit C (monetary
                        cost) added. Following that, the IETF DiffServ Working Group redefined bits 0

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                   41
QoS Configuration
QoS Configuration                                                                   6 Re-marking Configuration


                         to 5 of a ToS field as the DS CodePoint (DSCP) field in RFC 2474. In RFC 2474,
                         the field name is changed from ToS to differentiated service (DS). Figure 6-2
                         shows the DSCP field in an IP packet.
                         In the DS field, the first six bits (bits 0 to 5) are the DSCP and the last two
                         bits (bits 6 and 7) are reserved. The first three bits (0 to 2) are the Class
                         Selector CodePoint (CSCP), which represents the DSCP type. A DS node selects
                         a Per-Hop Behavior (PHB) based on the DSCP value.

Packet Re-marking
                    Packet re-marking involves setting the preceding QoS-related packet fields and
                    redefining the scheduling and transmission modes of packets. Currently, the device
                    supports the following packet re-marking modes:
                    ●    Re-marking 802.1p values of VLAN packets
                    ●    Re-marking DSCP values of IP packets
                    ●    Re-marking internal priorities of packets
                    ●    Re-marking local IDs


6.3 Configuration Precautions for Re-marking

6.4 Configuring MQC-based Priority Re-marking
Context
                    After re-marking is configured, the device re-marks the packets matching traffic
                    classification rules so that the packets can be scheduled or forwarded based on re-
                    marked priorities.

                          NOTE

                         ● Both remark 8021p and remark dscp can be bound to the same traffic policy and take
                           effect simultaneously.
                         ● For details about MQC-related configuration precautions, see "Configuration Precautions
                           for MQC" in MQC Configuration.


Procedure
         Step 1 Configure a traffic classifier.

                    For details about how to configure a traffic classifier, see 3.4 Configuring a Traffic
                    Classifier in "MQC Configuration".

         Step 2 Configure a traffic behavior.
                    1.   Create a traffic behavior and enter the traffic behavior view, or enter the view
                         of an existing traffic behavior.
                         traffic behavior behavior-name

                    2.   Configure a traffic behavior as needed.
                         –    Re-marking 802.1p values of VLAN packets

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    42
QoS Configuration
QoS Configuration                                                                             6 Re-marking Configuration


                                For the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S,
                                S6730E-H-V2, S5755E-H, S5755-S and S5755-H series:
                                remark 8021p { 8021p-value | inner-8021p }

