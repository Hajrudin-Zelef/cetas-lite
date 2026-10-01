---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-15
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1329, 1452]
sha256: adf113fcbcb5062ccc2b7a83a835a910411e8359f3b7172823cfef2bb46475a3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

                  ●      Run the display cpu-defend dynamic-adjust history-record [ packet-type
                         { arp-reply | arp-request | arp-request-uc | dhcp-request | dhcp-reply | nd |
                         igmp | pim | pim-mc | dhcp-discovery } ] { all | slot slot-id } command to
                         check the records of adaptive CPCAR adjustments for protocol packets.
                         The display cpu-defend dynamic-adjust history-record command is
                         supported only by the following models: S6780-H, S6750-H, S6730E-H-V2,
                         S6730-H-V2, S5732-H-V2, S5755-S, S6750E-S, S6750-S, S5755-H. The dhcp-
                         discovery protocol is not supported by the S6750-H, S5755-S, and S6780-H.

                  ●      Run the display cpu-defend drop-packet record [ packet-type packet-type ]
                         [ slot slot-id ] command to check the records of packet loss caused by rate
                         limiting on protocol packets.
                         The display cpu-defend drop-packet record command is supported only by
                         the following models: S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-
                         H-V2, S6750E-S, S6750-S, S5755-H.
                  ●      Run the display cpu-defend linkup statistics [ packet-type packet-type ]
                         { all | slot slot-id } command to check statistics on application layer
                         association.
                  ●      Run the display cpu-defend linkup configuration [ packet-type packet-
                         type ] { all | slot slot-id } command to check the configuration of application
                         layer association.
                  ●      Run the display cpu-defend local-host anti-attack [ slot slot-id ] command
                         to check statistics on the packets matching hardware-based ACLs after host
                         attack defense is enabled as well as the numbers and status of the ACLs
                         bound to protocols.
                  ●      Run the display cpu-defend rate [ packet-type packet-type ] { all | slot slot-
                         id | mcu } command to check the CPCAR value configured for protocol
                         packets.
                  ●      Run the display cpu-defend statistics [ packet-type packet-type ] { all | slot
                         slot-id | mcu } command to check statistics on the packets sent to the CPU.
                         When the ICMP fast reply function is enabled, statistics on ICMPv4 and
                         ICMPv6 packets are not differentiated. The total number of ICMPv4 and
                         ICMPv6 packets is recorded in the icmp field.
                  ●      Run the display cpu-defend filter statistics [ slot slot-id ] command to
                         check statistics on the packets discarded based on filters.
                         For loopback packets, only the number of discarded packets is counted, and
                         the number of discarded bytes is not.
                  ----End

3.4.8 Example for Configuring CPU Attack Defense
Networking Requirements
                  In Figure 3-2, a large number of users access the Internet through DeviceA. The
                  administrator determines that an attacker is sending a significant number of ARP
                  Request packets to DeviceA, which impacts normal operation of the device's CPU.
                  As such, the administrator must alleviate any adverse impact of ARP packets on
                  the CPU.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            19
Security Configuration
Security Configuration                                                                  3 Local Attack Defense Configuration


                  Figure 3-2 Networking diagram of local attack defense
                          NOTE

                         In this example, interface 1, interface 2, interface 3, and interface 4 represent 10GE 1/0/1,
                         10GE 1/0/2, 10GE 1/0/3, and 10GE 1/0/4, respectively.




Procedure
         Step 1 Configure an attack defense policy.

                  # Create an attack defense policy.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] cpu-defend policy test1

                  # Set the rate limit of ARP Request packets sent to the CPU.
                  [DeviceA-cpu-defend-policy-test1] car packet-type arp-request pps 128
                  [DeviceA-cpu-defend-policy-test1] quit

         Step 2 Apply the attack defense policy globally.
                  [DeviceA] cpu-defend-policy test1

                  ----End


Verifying the Configuration
                  # Check the configured attack defense policy.
                  [DeviceA] display cpu-defend policy test1
                  ==============================================
                  Policy name: test1
                  Policy applies on slot: <1>
                  Car packet-type arp-request(pps) : 128
                  ==============================================

                  # Check the CAR setting.
                  [DeviceA] display cpu-defend configuration all
                  Car configurations on mcu :
                  ----------------------------------------------------------------------
                  PacketType            Status      Current(pps) Default(pps)           Queue
                  ----------------------------------------------------------------------
                  arp-miss             Enabled             1536        1536         13
                  arp-reply            Enabled            2048         2048         23
                  arp-request           Enabled             128         2048         23
                  arp-request-uc         Enabled             2048         2048         23
                  ...


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                            20
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration


Configuration Scripts
                  DeviceA
                  #
                  sysname DeviceA
                  #
                  cpu-defend policy test1
                   car packet-type arp-request pps 128
                  #
                  cpu-defend-policy test1
                  #
                  return



3.5 Configuring Port Attack Defense

