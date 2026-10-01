---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-14
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1175, 1328]
sha256: 56b176f4908b0383e42e7aac5d9cf2ba3dc0b52a1d78b21a7c79070d60597422
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                   dhcp-discovery                     DHCP discovery                    1.5 times the default
                   NOTE                                                                 value
                    This parameter is not
                    supported by the S6750-H,
                    S5755-S, and S6780-H.

                   nd                                 IPv6 ND                           Twice the default value

                   pim                                PIM unicast                       Twice the default value

                   pim-mc                             PIM multicast                     Twice the default value

                   igmp                               IGMP                              Twice the default value




                          NOTE

                         Adaptive adjustment of the default CPCAR value for protocol packets is supported only by
                         the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5755-S, S6750E-S, S6750-S, S5755-H,
                         and S5732-H-V2 series.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable adaptive adjustment of the default CPCAR value for protocol packets.
                  cpu-defend dynamic-adjust [ packet-type { arp-reply | arp-request | arp-request-uc | dhcp-reply | dhcp-
                  request | nd | igmp | dhcp-discovery | pim | pim-mc } ] enable

                  If you do not specify the packet-type parameter, adaptive adjustment of the
                  default CPCAR value is enabled for all types of protocol packets that support this
                  function.

                  ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                           16
Security Configuration
Security Configuration                                                   3 Local Attack Defense Configuration


3.4.4 (Optional) Configuring Packet Loss Monitoring for
Protocol Packet Rate Limiting
Context
                  When the rate of protocol packets exceeds the CPCAR, the device discards excess
                  packets. To check the information about discarded packets, you can enable packet
                  loss monitoring for protocol packet rate limiting.

                          NOTE

                         This configuration is supported only by the following models: S6750-H, S6730E-H-V2,
                         S6730-H-V2, S6780-H, S6750E-S, S6750-S, S5732-H-V2, S5755-H. Packet loss monitoring for
                         protocol packet rate limiting is supported.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable packet loss monitoring for protocol packet rate limiting.
                  undo cpu-defend drop-packet monitor disable

                  By default, packet loss monitoring for protocol packet rate limiting is enabled.
         Step 3 Configure a rate limit for the packets that are to be sent to the CPU after they are
                discarded due to threshold crossing of protocol packet rate limiting.
                  cpu-defend drop-packet pps pps-value

                  By default, the rate limit for the packets that are to be sent to the CPU after they
                  are discarded due to threshold crossing of protocol packet rate limiting is 64 pps.

                  ----End

3.4.5 (Optional) Configuring a Filter

Context
                  Filters can be flexibly defined using ACLs. Note the following during the filter
                  configuration:
                  When a filter references an ACL, the device takes the relevant actions configured
                  in the ACL rules on packets that match those rules.
                  The ACL referenced by a filter does not support the following parameters. If these
                  parameters are configured, the filter is invalid.
                  ●      Basic ACL: vpn-instance
                  ●      Advanced ACL: vpn-instance, icmp-type, igmp-type, source-pool, source-
                         port-pool, destination-pool, and destination-port-pool
                  ●      Layer 2 ACL: 802.3
                  ●      Basic ACL6: vpn-instance
                  ●      Advanced ACL6: destination, vpn-instance, and icmpv6-type

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  17
Security Configuration
Security Configuration                                                           3 Local Attack Defense Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create an attack defense policy and enter the attack defense policy view.
                  cpu-defend policy policy-name

         Step 3 Configure a filter.
                  filter filter-id acl { acl-number | ipv6 ipv6-acl-number } [ interface { interface-type interface-number1 [ to
                  interface-type interface-number2 ] } &<1-8> ] [ vlan { vlan-id1 [ to vlan-id2 ] } &<1-8> ]

         Step 4 Return to the system view.
                  quit

         Step 5 Apply the attack defense policy.
                  ●      Configure attack defense policies in batches.
                         cpu-defend-policy policy-name batch slot { start-slot [ to end-slot ] } &<1-12>
                  ●      Configure an attack defense policy separately.
                         cpu-defend-policy policy-name [ slot slot-id | mcu ]

                  After an attack defense policy is created, you must apply the policy in the system
                  view. Otherwise, the policy does not take effect.

                  ----End

3.4.6 (Optional) Configuring Host Attack Defense

Context
                  After the ssh server acl and telnet server acl commands are configured, SSH and
                  Telnet packets are sent to the CPU. If host attack defense is configured, these
                  packets will match hardware-based ACLs. If packets match the deny rule in a
                  hardware-based ACL, they are directly discarded before they are sent to the CPU,
                  ensuring that normal packets can be successfully sent.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable host attack defense.
                  cpu-defend local-host anti-attack enable

                  ----End

3.4.7 Verifying the Configuration
Procedure
                  ●      Run the display cpu-defend policy [ policy-name ] command to check the
                         attack defense policy configuration.
                  ●      Run the display cpu-defend configuration [ packet-type packet-type ] { all |
                         slot slot-id | mcu } command to check the rate configuration for protocol
                         packets sent to the CPU.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                  18
Security Configuration
Security Configuration                                               3 Local Attack Defense Configuration


