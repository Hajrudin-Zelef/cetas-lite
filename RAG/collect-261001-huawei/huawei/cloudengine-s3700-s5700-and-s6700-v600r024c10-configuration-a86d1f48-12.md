---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-12
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [921, 1045]
sha256: d1557dd38373dfcaeba4610179ace4ee093e38bb15039d6021bbf4e51eeef580
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  Table 3-4 Default settings for attack source tracing
                   Parameter                                 Default Setting

                   Applicable packet types                   For details about the supported packet
                                                             types, run the auto-defend protocol
                                                             command.

                   Status                                    Enabled

                   Rate threshold                            The default rate threshold for
                                                             reporting attack source tracing events
                                                             is 128 pps on the S6780-H, S6750-H,
                                                             S6730E-H-V2, S6730-H-V2, S5732-H-
                                                             V2, S6750E-S, S6750-S, S5755-S,
                                                             S5755E-H, and S5755-H.
                                                             The default rate threshold for
                                                             reporting attack source tracing events
                                                             is 60 pps on the S5735I-L-V2, S5735I-
                                                             S-V2, S5735I-H-V2, S5735R-L-V2,
                                                             S5735E-L-V2, S5735-L-V2, S5735R-S-
                                                             V2, S5735E-S-V2, and S5735-S-V2.

                   Packet sampling ratio                     8 (One out of eight packets is
                                                             sampled.)

                   Source tracing mode                       Based on source IP addresses and
                                                             source MAC addresses

                   Punishment action                         Disabled




                  Table 3-5 Default settings for defense against malformed packet attacks,
                  fragmentation attacks, TCP SYN flood attacks, UDP flood attacks, and ICMP flood
                  attacks
                   Parameter                      Default Setting

                   Defense against malformed      Enabled
                   packet attacks

                   Defense against                Enabled
                   fragmentation attacks

                   Rate limit of fragments        155000000 bit/s

                   Defense against TCP SYN        Enabled
                   flood attacks

                   Rate limit of TCP SYN flood    155000000 bit/s
                   packets

                   Defense against UDP flood      Enabled
                   attacks




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              12
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration


                   Parameter                       Default Setting

                   Defense against ICMP flood      Enabled
                   attacks

                   Rate limit of ICMP flood        155000000 bit/s
                   packets




3.4 Configuring CPU Attack Defense

3.4.1 Understanding CPU Attack Defense
                  CPCAR is the core of CPU attack defense. You can run commands to change the
                  CPCAR values for each type of protocol packets, all packets sent to the CPU, and
                  application layer association. In addition, when the default CPCAR cannot meet
                  service requirements, adaptive adjustment of the default CPCAR value for protocol
                  packets is supported. In addition, CPU attack defense provides the filter function to
                  process the packets that meet characteristics based on the defined ACL.
                  Application layer association
                  Application layer association protects data transmitted over sessions of a specific
                  protocol. After such a session is established, the protocol-based default CPCAR
                  value no longer takes effect. Instead, the device uses the CPCAR value configured
                  for the protocol in application layer association to rate-limit the transmitted
                  packets. Typically, the CPCAR value configured for each protocol in application
                  layer association is much greater than the protocol-based default CPCAR value,
                  ensuring the reliability and stability of service operations.
                  For example, when FTP is enabled but no file has yet been transferred, the device
                  uses the default CPCAR value configured for FTP to rate-limit FTP packets. When
                  the device is transferring files, it begins using the CPCAR value configured for FTP
                  in application layer association once a successful FTP session establishment is
                  detected. File transfer may fail if the default CPCAR value is still used in this
                  scenario, as burst traffic is most likely to exceed the default value, which is usually
                  much smaller than the CPCAR value configured for application layer association.
                  Adaptive adjustment of the default CPCAR value for protocol packets
                  Adaptive adjustment of the default CPCAR value for protocol packets applies to
                  protocol packets related to user access, covering scenarios where a fixed default
                  CPCAR value cannot meet rate requirements. After this function is enabled, the
                  device dynamically adjusts the protocol-based default CPCAR value based on the
                  packet loss rate and CPU usage of protocol packets.
                  For example, when a large number of users initiate authentication by sending ARP
                  request packets at a rate exceeding the default CPCAR value, the device discards
                  excess ARP request packets and then adjusts the CPCAR value for ARP packets
                  based on the packet loss rate and CPU usage.
                  Filter function

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                               13
Security Configuration
Security Configuration                                                             3 Local Attack Defense Configuration


                  The filter function allows you to configure an ACL in an attack defense policy,
                  enabling the device to filter and process packets matching rules in the specified
                  ACL. If the action in an ACL rule is deny, the device directly discards the packets
                  matching this ACL rule. If the action is permit, the device increases the priority of
                  the matching packets before forwarding them.

3.4.2 Configuring the CPCAR Value

