---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-11
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [789, 920]
sha256: 230332e80dc7399ffdb9fd972375e6bd73e742f5f11eb1c20d6cb9df30f5f086
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  Level 1: Discarding malicious packets sent to the CPU through a filter or the
                  punishment function of attack source tracing.
                  Level 2: Control Plane Committed Access Rate (CPCAR) based on protocol packets,
                  whereby the device rate-limits the packets sent to the CPU based on the protocol
                  type, preventing excess packets of a protocol from being sent to the CPU.
                  CPCAR control is the core of CPU attack defense, rate-limiting the protocol packets
                  on a per device basis. In contrast, user-level rate limiting rate-limits protocol
                  packets based on the MAC address of the user who initiates an attack. Both
                  approaches will be depicted in the subsequent sections.
                  Level 3: Queue-based scheduling and rate limiting. After the rate of protocol
                  packets is limited by CPCAR, the device can allocate a queue to each type of
                  protocol packets, and schedules these queues based on weights or priorities. When
                  a conflict occurs, the device preferentially processes high-priority queues. Rate
                  limiting can also be implemented to limit the maximum rate of packets in each
                  queue sent to the CPU, with the device discarding the protocol packets that exceed
                  the rate limit in a queue.
                  In port attack defense, rate limiting is implemented by moving protocol packets to
                  low-priority queues for processing.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              9
Security Configuration
Security Configuration                                             3 Local Attack Defense Configuration


                  Level 4: Rate limiting on all packets. At this level, the total number of packets
                  processed by the CPU is limited to ensure that the CPU can process as many
                  packets as possible within its processing capability.
                  Before rate-limiting all packets, the device analyzes the contents and behaviors of
                  the packets sent to the CPU to determine whether these packets are attack
                  packets. If so, the device takes defense approaches such as discarding or rate-
                  limiting the packets. Packet attack defense includes defense against malformed
                  packet attacks, fragmentation attacks, TCP SYN flood attacks, UDP flood attacks,
                  and ICMP flood attacks.


3.2 Configuration Precautions for Local Attack Defense

3.3 Default Settings for Local Attack Defense
                  The following tables describe the default settings for local attack defense.

                  Table 3-1 Default settings for CPU attack defense
                   Parameter                                  Default Setting

                   Attack defense policy                      Named default, and has been applied
                                                              by default

                   CPCAR value for each type of protocol      Value in the default policy, and can be
                   packets                                    obtained using the display cpu-
                                                              defend configuration command

                   CPCAR value for all packets sent to the    You can run the display cpu-defend
                   CPU                                        configuration command to view the
                                                              default CAR of packets sent to the
                                                              CPU.

                   Application layer association              For details about the supported packet
                                                              types, run the application-apperceive
                                                              enable command.

                   Adaptive adjustment of the default         For details about the supported packet
                   CPCAR value for protocol packets           types, run the cpu-defend dynamic-
                                                              adjust enable command.

                   Filter                                     Not configured

                   Host attack defense                        Disabled




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              10
Security Configuration
Security Configuration                                               3 Local Attack Defense Configuration


                  Table 3-2 Default settings for port attack defense
                   Parameter                                  Default Setting

                   Applicable packet types                    For details about the supported packet
                                                              types, run the auto-port-defend
                                                              protocol disable command.

                   Status                                     Enabled

                   Rate threshold                             Rate threshold for port attack defense.
                                                              By default, the value for XLDP is 10%
                                                              of the default CAR, and that for other
                                                              protocols is 80% of the default CAR.
                                                              By default, the rate threshold on
                                                              common models for port attack
                                                              defense against IP fragmentation is
                                                              1228 pps.
                                                              By default, the rate threshold on the
                                                              S5735I-L-V2, S5735I-S-V2, S5735I-H-
                                                              V2, S5735R-L-V2, S5735E-L-V2, S5735-
                                                              L-V2, S5735R-S-V2, S5735E-S-V2,
                                                              S5735-S-V2, and S3710-H for port
                                                              attack defense against IP
                                                              fragmentation is 320 pps.
                                                              By default, the rate threshold on the
                                                              S5755-S for port attack defense
                                                              against IP fragmentation is 614 pps.

                   Packet sampling ratio                      8 (One out of eight packets is
                                                              sampled.)

                   Aging time                                 300s




                  Table 3-3 Default settings for user-level rate limiting
                   Parameter                         Default Setting

                   Applicable packet types           For details about the supported packet types,
                                                     run the cpu-defend host-car command.

                   Status                            Enabled globally and on interfaces

                   Threshold                         20 pps




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             11
Security Configuration
Security Configuration                                            3 Local Attack Defense Configuration


