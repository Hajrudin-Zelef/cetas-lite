---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-24
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [2625, 2779]
sha256: a9e80d07dd89184ab9bd1be636ab3957ac04a36e306991a8a5c2075ff4b54c22
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

3.12.2 Configuring Defense Against ICMP Flood Attacks

Context
                  With defense against ICMP flood attacks enabled, the device rate-limits the
                  received ICMP messages, discarding any that exceeds the limit.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable defense against ICMP flood attacks.
                  anti-attack icmp-flood enable

                         NOTE

                  You can also run the anti-attack enable command in the system view to enable attack defense
                  against all attack packets, including ICMP flood attack packets.

         Step 3 Configure the rate limit for ICMP flood attack packets.
                  anti-attack icmp-flood car cir cir-num

                  ----End


Verifying the Configuration
                  Run the display anti-attack statistics icmp-flood command to check statistics
                  relating to defense against ICMP flood attacks.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 44
Security Configuration
Security Configuration                                               3 Local Attack Defense Configuration




3.13 Maintaining Local Attack Defense
                  During routine maintenance, you can also clear statistics on local attack defense.
                  Run the following commands in the user view as required.


                         NOTICE

                  Cleared statistics cannot be restored. Exercise caution when running reset
                  commands.


                  Table 3-7 Clearing local attack defense statistics
                   Operation                                  Command

                   Clear statistics about the packets         reset cpu-defend statistics [ packet-
                   sent to the CPU                            type packet-type ] { all | slot slot-id |
                                                              mcu }

                   Clear statistics about application         reset cpu-defend linkup statistics
                   layer association                          [ packet-type packet-type ] { slot slot-
                                                              id | all }
                   Clear the records of packet loss           reset cpu-defend drop-packet record
                   caused by rate limiting on protocol        [ packet-type packet-type ] [ slot slot-
                   packets                                    id ]
                                                              This command is supported only by the
                                                              following models: S6750-H, S6730E-H-
                                                              V2, S6730-H-V2, S6780-H, S6750E-S,
                                                              S6750-S, S5732-H-V2, S5755-H.

                   Clear statistics about the packets         reset cpu-defend filter statistics [ slot
                   sent to the CPU based on the filter        slot-id ]
                   Clear packet statistics about port         reset cpu-defend auto-port-defend
                   attack defense                             statistics [ slot slot-id ]

                   Clear source tracing information           reset cpu-defend auto-port-defend
                   about port attack defense                  attack-source [ slot slot-id ]

                   Clear packet statistics about user-        reset cpu-defend host-car [ mac-
                   level rate limiting                        address mac-address ] statistics [ slot
                                                              slot-id ]
                                                              This command is supported only by the
                                                              following models: S6750-H, S6730E-H-
                                                              V2, S6730-H-V2, S6780-H, S6750E-S,
                                                              S6750-S, S5732-H-V2, S5755-H.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               45
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration


                   Operation                                  Command

                   Clear records of packet loss caused        reset cpu-defend host-car drop-packet
                   by user-level rate limiting                record [ car-id car-id ] [ slot slot-id ]
                                                              This command is supported only by the
                                                              following models: S6750-H, S6730E-H-
                                                              V2, S6730-H-V2, S6780-H, S6750E-S,
                                                              S6750-S, S5732-H-V2, S5755-H.

                   Clear statistics about attack source       reset auto-defend attack-source [ slot
                   tracing                                    slot-id ]
                   Clear historical statistics about attack   reset auto-defend attack-source
                   source tracing                             history [ slot slot-id ]

                   Clear attack source tracing statistics     reset auto-defend attack-source trace-
                   based on the source tracing mode           type { source-mac [ mac-address ] |
                                                              source-ip [ ip-address | ipv6-address ] |
                                                              source-portvlan [ interface interface-
                                                              type interface-number vlan vlan-id
                                                              [ inner-vlan inner-vlan-id ] ] } [ slot
                                                              slot-id ]
                   Clear packet statistics about attack       reset anti-attack statistics [ abnormal
                   defense                                    | fragment | tcp-syn | udp-flood | icmp-
                                                              flood ]




3.14 Configuration Examples for Local Attack Defense

3.14.1 Example for Configuring Attack Defense

Networking Requirements
                  As shown in Figure 1, if an attacker on the Internet launches a malformed packet
                  attack, a fragmentation attack, or a flood attack on DeviceA, DeviceA may break
                  down. To prevent such issues, the network administrator needs to deploy attack
                  defense measures on DeviceA, so as to secure the network environment and
                  ensure services run as normal.


                  Figure 3-14 Network diagram of defense against malformed packet attacks,
                  fragmentation attacks, and flood attacks




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              46
Security Configuration
Security Configuration                                                               3 Local Attack Defense Configuration


Procedure
         Step 1 Configure defense against malformed packet attacks.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] anti-attack abnormal enable

         Step 2 Configure defense against fragmentation attacks and set the rate limit of packet
                fragments to 15000 bit/s.
                  [DeviceA] anti-attack fragment enable
                  [DeviceA] anti-attack fragment car cir 15000

