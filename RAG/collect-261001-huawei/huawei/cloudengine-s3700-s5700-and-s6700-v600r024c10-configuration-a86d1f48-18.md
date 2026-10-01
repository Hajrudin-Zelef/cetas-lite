---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-18
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1716, 1866]
sha256: ffe51754d5738057b4c118f6fda4057206a67d00747f36ca9837585ff9218e85
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  By default, user-level rate limiting is enabled on interfaces.
                  When user-level rate limiting is enabled globally, user-level rate limiting is also
                  enabled on interfaces. In this case, you can run the host-car disable command in
                  the interface view to disable user-level rate limiting on interfaces as needed.

                  ----End

Example
                  Enable user-level rate limiting, set the user-level rate limit to 15 pps, and limit the
                  rate of only ARP packets.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         25
Security Configuration
Security Configuration                                                   3 Local Attack Defense Configuration

                  <HUAWEI> system-view
                  [HUAWEI] cpu-defend host-car enable
                  [HUAWEI] cpu-defend host-car pps 15
                  [HUAWEI] cpu-defend host-car arp


3.6.3 (Optional) Configuring Packet Loss Monitoring for User-
Level Rate Limiting
Context

                  After user-level rate limiting is enabled, the switch discards the excess packets if
                  the rate of packets from the same source MAC address exceeds the rate limit
                  within a specified period of time. You can configure packet loss monitoring for
                  user-level rate limiting to check which packets are discarded.

                          NOTE

                         This configuration is supported only by the following models: S6750-H, S6730E-H-V2,
                         S6730-H-V2, S6780-H, S6750E-S, S6750-S, S5732-H-V2, S5755-H. Packet loss monitoring for
                         protocol packet rate limiting is supported.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable monitoring for packets discarded due to user-level rate limiting.
                  undo cpu-defend host-car drop-packet monitor disable

                  By default, packet loss monitoring for user-level rate limiting is enabled.
         Step 3 The rate limit for packets discarded due to user-level rate limiting is set.
                  cpu-defend host-car drop-packet pps pps-value

                  By default, the packet loss monitoring threshold for user-level rate limiting is 64
                  pps.

                  ----End

3.6.4 Verifying the Configuration
Procedure
                  ●      Run the display cpu-defend host-car [ mac-address mac-address ] statistics
                         [ slot slot-id ] command to check the number of packets discarded due to
                         user-level rate limiting.
                  ●      Run the display cpu-defend host-car drop-packet record [ car-id car-id ]
                         [ slot slot-id ] command to check records of packet loss caused by user-level
                         rate limiting.
                  ----End


3.7 Configuring Attack Source Tracing

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  26
Security Configuration
Security Configuration                                             3 Local Attack Defense Configuration



3.7.1 Understanding Attack Source Tracing

                  Attack source tracing can defend against DoS attacks. As shown in Figure 3-3,
                  attack source tracing involves four steps:
                  1.     The device parses packets to be sent to the CPU based on IP addresses, MAC
                         addresses, and ports. A port is identified by physical port plus VLAN.
                  2.     The device counts the number of received protocol packets based on IP
                         addresses, MAC addresses, or ports.
                  3.     When the number of packets sent to the CPU in a unit time exceeds the
                         threshold, the device considers that an attack has occurred.
                  4.     When detecting an attack, the device sends a log or alarm to notify the
                         administrator or directly implements punishment (for example, discarding
                         attack packets).


                  Figure 3-3 Attack source tracing process




                  Attack source tracing also provides the whitelist function. After legitimate users
                  are added to a whitelist, the device does not perform source tracing or
                  punishment actions on packets from these whitelisted users, so that such packets
                  can be sent to the CPU for processing. A whitelist can be flexibly configured based
                  on ACLs or ports.

3.7.2 Configuring Attack Source Tracing

Context
                  After attack source tracing is configured, the device analyzes packets sent to the
                  CPU to determine if the CPU is under attack. If so, the device traces the attack
                  source and notifies the administrator through logs or alarms, who can then take
                  appropriate measures to defend against the attack source.



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              27
Security Configuration
Security Configuration                                                          3 Local Attack Defense Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create an attack defense policy and enter the attack defense policy view.
                  cpu-defend policy policy-name

         Step 3 Enable attack source tracing.
                  auto-defend enable

         Step 4 Set the rate threshold of attack source tracing.
                  auto-defend threshold threshold-value

         Step 5 Set the packet sampling ratio for attack source tracing.
                  auto-defend attack-packet sample sample-value

         Step 6 Specify the type of packets to which attack source tracing is applied.
                  auto-defend protocol { { arp | icmp | dhcp | ttl-expired | tcp | udp | udpv6 | 8021x | telnet | dhcpv6 | dns
                  | nd | icmpv6 | tcpv6 | igmp | mld } * | all }

         Step 7 Set the attack source tracing mode.
                  auto-defend trace-type { source-mac | source-ip | source-portvlan } *

                  The source tracing modes are listed in descending order of priority as follows:
                  MAC address-based > IP address-based > interface- and VLAN-based. If multiple
                  source tracing modes are configured, the configured modes take effect according
                  to this order of priorities.
         Step 8 (Optional) Configure a whitelist for attack source tracing.
                  auto-defend whitelist whitelist-id { acl acl-number | acl ipv6 ipv6-acl-number | interface interface-type
                  interface-number }

                  By default, no whitelist is configured for attack source tracing.

                          NOTE

