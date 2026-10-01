---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-13
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1046, 1174]
sha256: f0ffcd76527050b24e678b8e3773139ddb62999f0c5e8b48fde2b984c68533b8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Context
                  To reduce the total number of packets reaching the CPU, and to mitigate the
                  impact that different types of packets have on one another in order to protect the
                  CPU, the device can rate-limit packets to be sent to the CPU at different levels.
                  This includes CPCAR based on protocol packets, rate limiting on all packets sent to
                  the CPU, and rate limiting based on protocol association.
                  ●      Rate limiting based on protocol association has the highest priority and
                         therefore takes precedence over CPCAR based on protocol packets as well as
                         rate limiting on all protocol packets sent to the CPU.
                  ●      If only CPCAR based on protocol packets and rate limiting on all protocol
                         packets sent to the CPU are configured, the device rate-limits packets based
                         on the smaller value of the two.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create an attack defense policy and enter the attack defense policy view.
                  cpu-defend policy policy-name

         Step 3 Configure the method used to rate-limit packets sent to the CPU.
                  ●      Set a CPCAR value for protocol packets of a specified type.
                         car packet-type packet-type pps pps-value

                  ●      Set a CPCAR value for all packets sent to the CPU.
                         car all-packets pps pps-value

                  ●      Configure rate limiting based on protocol association.
                         a.   Enable protocol association.
                              application-apperceive { bgp | bgp4plus | isis | ftp | ssh | telnet | tftp | ospf | ospfv3 | http |
                              https | https-client | m-lag | m-lag-sync | mpls-ldp } enable

                              The S5735-L-V2, S5735E-L-V2, S5735R-L-V2, and S5735I-L-V2 do not
                              support the following parameters: bgp, bgp4plus, isis, m-lag, m-lag-sync,
                              mpls-ldp.
                              The S5735-S-V2, S5735R-S-V2, S5735E-S-V2, and S5735I-S-V2 do not
                              support the following parameters: m-lag, m-lag-sync, mpls-ldp.
                              The S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2,
                              S5732-H-V2, S5755E-H, S5755-H, and S5755-S do not support https-
                              client.
                              The S5735I-H-V2 does not support mpls-ldp.
                              The S3710-H does not support the following parameters: bgp, bgp4plus,
                              isis, m-lag, m-lag-sync, ospf, ospfv3, mpls-ldp.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                    14
Security Configuration
Security Configuration                                                            3 Local Attack Defense Configuration


                              The S5755-S does not support mpls-ldp.
                         b.   Set a CPCAR value for packets of a specific protocol upon the
                              establishment of the protocol connection.
                              linkup-car packet-type { bgp | bgp4plus | isis | ftp | ssh | telnet | tftp | m-lag | m-lag-sync |
                              ospf | ospfv3 | http | https | https-client | mpls-ldp } pps pps-value

                              The S5735-L-V2, S5735E-L-V2, S5735R-L-V2, and S5735I-L-V2 do not
                              support the following parameters: bgp, bgp4plus, isis, m-lag, m-lag-sync,
                              mpls-ldp.
                              The S5735-S-V2, S5735R-S-V2, S5735E-S-V2, and S5735I-S-V2 do not
                              support the following parameters: m-lag, m-lag-sync, mpls-ldp.
                              The S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2,
                              S5732-H-V2, S5755E-H, S5755-H, and S5755-S do not support https-
                              client.
                              The S5735I-H-V2 does not support mpls-ldp.
                              The S3710-H does not support the following parameters: bgp, bgp4plus,
                              isis, m-lag, m-lag-sync, ospf, ospfv3, mpls-ldp.
                              The S5755-S does not support mpls-ldp.
                         c.   Configure the penalty ratio threshold for protocol association.
                              linkup session anti-attack ratio-threshold rate-value-percent

                              By default, the threshold is 50%.
         Step 4 (Optional) Configure the device to discard packets to be sent to the CPU.
                  deny packet-type packet-type

                  By default, the device does not discard the packets to be sent to the CPU.
         Step 5 (Optional) Configure the description of the attack defense policy.
                  description description

                  By default, no description is configured for an attack defense policy.
         Step 6 Return to the system view.
                  quit

         Step 7 Apply the attack defense policy.
                  ●      Configure attack defense policies in batches.
                         cpu-defend-policy policy-name batch slot { start-slot [ to end-slot ] } &<1-12>

                  ●      Configure an attack defense policy separately.
                         cpu-defend-policy policy-name [ slot slot-id | mcu ]

                  After an attack defense policy is created, you must apply the policy in the system
                  view. Otherwise, the policy does not take effect.

                  ----End

3.4.3 (Optional) Configuring Adaptive Adjustment of the
Default CPCAR Value for Protocol Packets

Context
                  If a fixed default CPCAR value for protocol packets does not meet rate
                  requirements, configure adaptive adjustment of the default CPCAR value. You can

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                      15
Security Configuration
Security Configuration                                                       3 Local Attack Defense Configuration


                  run the display cpu-defend dynamic-adjust history-record command to view the
                  records of adaptive CPCAR adjustments.

                  The following table lists the types of protocol packets that support adaptive
                  CPCAR adjustment and the maximum CPCAR value after adjustment.

                   Protocol Packet Type               Protocol Packet                   Maximum CPCAR Value
                                                      Description                       After Adjustment

                   arp-reply                          ARP reply                         Twice the default value

                   arp-request                        ARP request                       Twice the default value

                   arp-request-uc                     Unicast ARP request               Twice the default value

                   dhcp-reply                         DHCP reply                        1.5 times the default
                                                                                        value

                   dhcp-request                       DHCP request                      1.5 times the default
                                                                                        value

