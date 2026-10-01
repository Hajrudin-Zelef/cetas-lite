---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-49
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6792, 6925]
sha256: 39c22efbcd833619b6e2f324bb8c4c431c7d4c75bb8f944a2b68d2badb9029a9
---

                    # Check statistics about the traffic policy applied to 10GE 1/0/1.
                    [DeviceB] display traffic-policy statistics interface 10ge 1/0/1 inbound
                    Traffic policy: p1, inbound
                    --------------------------------------------------------------------------------
                     Slot: 1
                     Item               Packets               Bytes          pps          bps
                    -------------------------------------------------------------------------------
                     Matched                   0                 0          0           0
                      Passed                 0                 0          0           0
                      Dropped                  0                0           0          0
                       Filter              0                 0          0           0
                       CAR                   0                0           0          0
                    -------------------------------------------------------------------------------

                    # Take voice packets as an example. Send voice packets with VLAN ID 10 to 10GE
                    1/0/1 at the rates of 1500 kbit/s and 12500 kbit/s. Then, run the display traffic-
                    policy statistics interface 10ge 1/0/1 inbound command to view traffic statistics
                    on the interface. If the configuration is successful, all packets with VLAN ID 10 and
                    sent at the rate of 1500 kbit/s can pass through 10GE1/0/1, and no packet is
                    discarded. Packets with VLAN ID 10 and sent at the rate of 12500 kbit/s pass
                    through 10GE 1/0/1 at the rate of 10000 kbit/s, and no packet is discarded. If
                    packets with VLAN ID 10, packets with VLAN ID 30, and packets with VLAN ID 20
                    are sent to 10GE 1/0/1 at the rates of 3000 kbit/s, 5000 kbit/s, and 5000 kbit/s
                    respectively, these packets pass through the interface at the rate of 12000 kbit/s,
                    and excess packets are discarded.

Configuration Scripts
                    ●     DeviceB
                          #
                          sysname DeviceB
                          #
                          vlan batch 10 20 30
                          #
                          qos car car1 cir 12000 kbps
                          #
                          traffic classifier c1 type or
                           if-match vlan 10
                          #
                          traffic classifier c2 type or
                           if-match vlan 20
                          #
                          traffic classifier c3 type or
                           if-match vlan 30
                          #
                          traffic behavior b1
                           statistics enable
                           car cir 2000 kbps pir 10000 kbps green pass
                           car car1 share
                           remark dscp ef
                          #
                          traffic behavior b2
                           statistics enable
                           car cir 4000 kbps pir 10000 kbps green pass
                           remark dscp af13
                          #
                          traffic behavior b3
                           statistics enable
                           car cir 4000 kbps pir 10000 kbps green pass
                           remark dscp af33


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                          122
QoS Configuration                                                  9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                  based Rate Limiting Configuration

                        #
                        traffic policy p1
                         classifier c1 behavior b1 precedence 5
                         classifier c2 behavior b2 precedence 10
                         classifier c3 behavior b3 precedence 15
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10 20 30
                         traffic-policy p1 inbound
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 10 20 30
                        #
                        return


9.5.6 Example for Configuring Traffic Policing to Limit the
Rate of Each IP Address on a Network Segment
Networking Requirements
                    In Figure 9-9, users on an enterprise network send packets through DeviceA and
                    DeviceB, and access the external network through DeviceC. Users reside on two
                    different network segments. It is required that the rate of traffic from each IP
                    address on network segment 192.168.1.0/24 be limited to 64 kbit/s and the rate of
                    traffic from each IP address on network segment 192.168.2.0/24 be limited to 128
                    kbit/s.

                    Figure 9-9 Network diagram for configuring traffic policing to limit the rate of
                    each IP address on a network segment
                         NOTE

                        In this example, interface 1 and interface 2 represent 10GE 1/0/1 and 10GE 1/0/2,
                        respectively.




Procedure
         Step 1 Create VLANs and configure interfaces so that enterprise users can access the
                network through DeviceB.
                    # Create VLANs 10 and 20 on DeviceB and add 10GE 1/0/1 to the VLANs.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       123
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration

                    [DeviceB] vlan batch 10 20
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] portswitch
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 10 20
                    [DeviceB-10GE1/0/1] quit

                    # Configure VLANIF interfaces on DeviceB and configure IP addresses for them.
                    [DeviceB] interface vlanif 10
                    [DeviceB-Vlanif10] ip address 192.168.1.1 24
                    [DeviceB-Vlanif10] quit
                    [DeviceB] interface vlanif 20
                    [DeviceB-Vlanif20] ip address 192.168.2.1 24
                    [DeviceB-Vlanif20] quit

         Step 2 Configure traffic classifiers.

