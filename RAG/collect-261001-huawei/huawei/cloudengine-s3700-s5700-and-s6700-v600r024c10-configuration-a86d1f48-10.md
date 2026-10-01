---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-10
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [654, 788]
sha256: 4edcb7e0ea1902283584a981595e11192f4888e4239cb9d8d590fb79be125013
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Management Plane
                  The key to security of the management plane is to ensure that devices can be
                  managed only by authorized users, including who can log in to a device and what
                  operations they can perform. As shown in Figure 2-1, the security of the
                  management plane is mainly achieved by allowing only the administrator to log in
                  to the device. This is implemented by configuring user names, passwords, and
                  ACLs to control login of users, enabling the administrator to log in through
                  STelnet, and configuring the user level to control the operations users can
                  perform. For details, see CLI-based Device Login Configuration.


Control Plane
                  The control plane controls data forwarding based on the CPU. The CPU is the
                  brain of a device, controlling operations of device components. Therefore, security
                  of the CPU must be guaranteed so that devices and protocols function normally. If
                  too many protocol packets are sent to the CPU, the CPU will become overloaded,
                  leading to poor device performance and service interruptions. As the core
                  component of a device, the CPU is one of the major targets of attacks from
                  unauthorized users. Devices support security protection services on the control
                  plane, including local attack defense.

                  As shown in Figure 2-2, the default Central Processing Unit-committed Access
                  Rate (CPCAR) value limits the rate of CPU-bound protocol packets to ensure that
                  the CPU runs normally. If the CPU usage is still high after CPCAR limiting is
                  performed, perform the following steps:

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             6
Security Configuration
Security Configuration                                                            2 Overview of Security


                  ●      Adjust the CPCAR value: Decrease the CPCAR value to reduce the number of
                         CPU-bound protocol packets.
                  ●      Trace attack sources: Analyze the CPU-bound protocol packets, and set the
                         threshold for the number of packets sent to the CPU. Take measures to
                         handle packets exceeding the threshold, for example, dropping the packets,
                         shutting down interfaces, and configuring a blacklist.

                  Figure 2-2 CPU defense




Forwarding Plane
                  The forwarding plane searches for forwarding entries to perform data forwarding.
                  There are two types of attacks targeting the forwarding plane:
                  ●      Exhausting forwarding entries: prevents forwarding entries of authorized users
                         from being learned and their traffic from being forwarded.
                  ●      Tampering with forwarding entries: causes traffic of authorized users to be
                         forwarded to incorrect destinations.
                  A device can defend against attacks at Layer 2 and Layer 3.
                  ●      Layer 2 network: Layer 2 data is forwarded based on the MAC address table.
                         That is, the device needs to search for MAC address entries to forward data
                         traffic. Therefore, the MAC address table is prone to attacks. An attack can be
                         in the form of unauthorized users sending a large number of packets, which
                         are then recorded as MAC address entries. Also, when no MAC address entry
                         is found for a packet, the packet is broadcast, consuming bandwidth resources
                         and potentially causing broadcast storms. The device provides a wide variety
                         of security mechanisms to protect MAC address tables, including DHCP
                         snooping and storm suppression.
                  ●      Layer 3 network: Layer 3 data is forwarded based on the ARP table and
                         routing table. The entries in the routing table are generated through routing
                         protocol negotiation, so they are difficult to attack. The ARP entries are
                         generated by protocol packet exchange. Attackers may send a large number
                         of protocol packets or forged protocol packets to attack the ARP table.
                         Therefore, the ARP table must be properly protected in Layer 3 forwarding.
                         The device prevents the preceding attacks through ARP security, DAI, and
                         IPSG and URPF.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                7
Security Configuration
Security Configuration                                              3 Local Attack Defense Configuration




            3            Local Attack Defense Configuration


                  3.1 Overview of Local Attack Defense
                  3.2 Configuration Precautions for Local Attack Defense
                  3.3 Default Settings for Local Attack Defense
                  3.4 Configuring CPU Attack Defense
                  3.5 Configuring Port Attack Defense
                  3.6 Configuring User-Level Rate Limiting
                  3.7 Configuring Attack Source Tracing
                  3.8 Configuring Defense Against Malformed Packet Attacks
                  3.9 Configuring Defense Against Fragmentation Attacks
                  3.10 Configuring Defense Against TCP SYN Flood Attacks
                  3.11 Configuring Defense Against UDP Flood Attacks
                  3.12 Configuring Defense Against ICMP Flood Attacks
                  3.13 Maintaining Local Attack Defense
                  3.14 Configuration Examples for Local Attack Defense
                  3.15 Troubleshooting Local Attack Defense


3.1 Overview of Local Attack Defense
Definition
                  Local attack defense is a central processing unit (CPU) protection mechanism
                  designed to ensure that the CPU can properly process normal services. The CPU of
                  a device may receive both normal service packets and malicious packets targeting
                  the CPU.
                  ●      If a large number of normal service packets are sent to the CPU, its usage
                         surges. This severely impacts device performance, ultimately disrupting
                         services.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             8
Security Configuration
Security Configuration                                             3 Local Attack Defense Configuration


                  ●      If the CPU is busy processing attack packets for an extended period, normal
                         services will be interrupted, and in some cases even the entire system will
                         crash.
                  To solve these issues, the local attack defense function is introduced. With this
                  function enabled, the CPU can run properly when receiving a large number of
                  normal service packets or attack packets, ensuring normal service running.

Basic Functions
                  Basic functions of local attack defense include CPU attack defense, port attack
                  defense, user-level rate limiting, attack defense, and attack source tracing. As
                  shown in Figure 3-1, local attack defense uses a multi-level security mechanism to
                  provide hierarchical protection for the device.

                  Figure 3-1 Hierarchical protection of local attack defense




