---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-47
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [6484, 6646]
sha256: 29857c7f10e0a01d1c4acd59b76f0897e3cc7263947e01db3437cbd7cf10dcd6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                        Classifier: c3
                         Type: OR
                        Behavior: b3
                         Committed Access Rate:
                           CIR 8000 (Kbps), PIR 8000 (Kbps), CBS 64000 (Bytes), PBS 64000 (Bytes)
                           Color Mode: color blind
                           Conform Action: pass
                           Yellow Action: pass
                           Exceed Action: discard
                         Statistics: enable

                    # Check statistics about the traffic policy applied to 10GE 1/0/1.
                    [DeviceB] display traffic-policy statistics interface 10ge 1/0/1 inbound
                    Traffic policy: p1, inbound
                    --------------------------------------------------------------------------------
                     Slot: 1
                     Item                Packets              Bytes          pps          bps
                     -------------------------------------------------------------------------------
                     Matched                363949175            46585494400          8460795 8663854896
                      Passed              363949175            46585494400          8460795 8663854896
                      Dropped                   0                0           0          0
                       Filter               0                 0          0           0
                       CAR                    0                0           0          0
                     -------------------------------------------------------------------------------

                    The preceding command output shows that the traffic policy p1 is applied to 10GE
                    1/0/1.

Configuration Scripts
                    ●     DeviceB
                          #
                          sysname DeviceB
                          #
                          vlan batch 10 20 30


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         117
QoS Configuration                                                  9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                  based Rate Limiting Configuration

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
                         car cir 2000 kbps
                        #
                        traffic behavior b2
                         statistics enable
                         car cir 4000 kbps
                        #
                        traffic behavior b3
                         statistics enable
                         car cir 8000 kbps
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
                        return

                    ●   DeviceA
                        #
                        sysname DeviceA
                        #
                        vlan batch 10 20 30
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 20
                        #
                        interface 10GE1/0/3
                         port link-type access
                         port default vlan 30
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 10 20 30
                        #
                        return


9.5.5 Example for Configuring Traffic Policing (Level-2 CAR)

Networking Requirements
                         NOTE

                        Only the S6780-H, S6750-S, S6730-H-V2, S6750-H, S5732-H-V2, S6750E-S, S6730E-H-V2,
                        S5755E-H, S5755-S and S5755-H series support this example.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       118
QoS Configuration                                             9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                             based Rate Limiting Configuration


                    In Figure 9-8, DeviceB is connected to DeviceC through interface 2, and the
                    enterprise can access the network through DeviceB and DeviceC.

                    Figure 9-8 Network diagram for configuring traffic policing (level-2 CAR)
                         NOTE

                        In this example, interface 1 and interface 2 represent 10GE 1/0/1 and 10GE 1/0/2,
                        respectively.




                    On this network, the network-side bandwidth is lower than the enterprise's LAN
                    bandwidth. As a result, network congestion may occur on network-side interfaces,
                    causing data loss. Therefore, the total egress bandwidth needs to be limited to
                    12000 kbit/s. In addition, traffic policing needs to be performed on voice, video,
                    and data services to limit the traffic rates within proper ranges.

                    Voice, video, and data services are transmitted in VLAN 10, VLAN 30, and VLAN 20
                    respectively, and have different QoS requirements in descending order of priority.
                    Therefore, DeviceB needs to re-mark DSCP priorities of different service packets so
                    that DeviceC can process the packets based on their priorities, ensuring QoS.

                    Table 9-8 describes the configuration requirements.

                    Table 9-8 QoS guarantee provided by DeviceB for uplink traffic

                     Traffic        CIR (kbit/s)              PIR (kbit/s)                DSCP Priority
                     Type

                     Voice          2000                      10000                       46

                     Data           4000                      10000                       14

                     Video          4000                      10000                       30




Procedure
         Step 1 Create VLANs and configure interfaces.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   119
QoS Configuration                                                    9 Traffic Policing, Traffic Shaping, and Interface-
QoS Configuration                                                                    based Rate Limiting Configuration


