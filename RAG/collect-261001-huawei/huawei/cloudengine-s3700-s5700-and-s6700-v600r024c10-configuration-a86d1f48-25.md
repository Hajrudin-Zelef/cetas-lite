---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-25
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [2780, 2923]
sha256: 295ee477dcf304328a90dc50c16f0ef005ffa8836975b1b08bc81765dbaf3517
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

         Step 3 Configure defense against flood attacks.
                  # Configure defense against TCP SYN flood attacks and set the rate limit of TCP
                  SYN packets to 15000 bit/s.
                  [DeviceA] anti-attack tcp-syn enable
                  [DeviceA] anti-attack tcp-syn car cir 15000

                  # Configure defense against UDP flood attacks to enable the device to discard
                  UDP packets sent from specified ports.
                  [DeviceA] anti-attack udp-flood enable

                  # Configure defense against ICMP flood attacks and set the rate limit of ICMP
                  flood packets to 15000 bit/s.
                  [DeviceA] anti-attack icmp-flood enable
                  [DeviceA] anti-attack icmp-flood car cir 15000

                  ----End

Verifying the Configuration
                  # After the configuration is complete, run the display anti-attack statistics
                  command to view attack defense statistics.
                  <DeviceA> display anti-attack statistics
                  Packets Statistic Information:
                  -------------------------------------------------------------------------------
                  AntiAtkType TotalPacketNum               DropPacketNum             PassPacketNum
                            (H)        (L)      (H)        (L)       (H)       (L)
                  -------------------------------------------------------------------------------
                  Abnormal        0         0        0         0        0        0
                  Fragment        0        0         0        0        0         0
                  Icmp-flood 0             0         0        0        0         0
                  Tcp-syn       0         58        0        0        0         58
                  Udp-flood       0        0         0        0        0         0
                  -------------------------------------------------------------------------------


Configuration Scripts
                  DeviceA
                  #
                  sysname DeviceA
                  #
                  anti-attack abnormal enable
                  anti-attack fragment enable
                  anti-attack fragment car cir 15000
                  anti-attack tcp-syn enable
                  anti-attack tcp-syn car cir 15000
                  anti-attack udp-flood enable


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          47
Security Configuration
Security Configuration                                                3 Local Attack Defense Configuration

                  anti-attack icmp-flood enable
                  anti-attack icmp-flood car cir 15000
                  #
                  return



3.15 Troubleshooting Local Attack Defense

3.15.1 Attack Source Tracing Does Not Take Effect

Fault Symptom
                  The configured attack source tracing function does not take effect.

Possible Causes
                  Possible causes are as follows:
                  ●      The attack defense policy is not applied.
                  ●      A large rate threshold is set, resulting in failure of the attack source tracing
                         function to identify attack packets.

Procedure
                  1.     Run the display current-configuration command to check whether the
                         attack defense policy is applied.
                         –   If the command output contains cpu-defend-policy, the attack defense
                             policy has been applied. In this case, go to 2.
                         –   If the command output does not contain cpu-defend-policy, the attack
                             defense policy is not applied. In this case, you need to run the cpu-
                             defend-policy command in the system view to apply the attack defense
                             policy.
                  2.     Check whether a high rate threshold is set for attack source tracing.
                         Run the display auto-defend configuration command to check the value of
                         the auto-defend threshold field. If the value is large, run the auto-defend
                         threshold command in the attack defense policy view to reduce the value.

3.15.2 Protocol Packets Are Not Sent to the CPU

Fault Symptom
                  Protocol packets are not sent to the CPU after CPU attack defense is configured.

Possible Causes
                  Possible causes are as follows:
                  ●      A deny rule is configured to discard specified protocol packets to be sent to
                         the CPU.
                  ●      The CPU is attacked by illegitimate packets.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                 48
Security Configuration
Security Configuration                                                    3 Local Attack Defense Configuration


Procedure
                  1.     Check whether a deny rule has been configured to discard specified protocol
                         packets to be sent to the CPU.
                         a.   Run the display current-configuration command in the system view to
                              check the configured attack defense policy.
                         b.   Run the display cpu-defend policy [ policy-name ] command to check
                              whether the deny rule is configured in the attack defense policy for the
                              protocol packets to be sent to the CPU.
                              If so, run the car command in the attack defense policy view to set the
                              rate limit for the protocol packets sent to the CPU.
                              If not, go to the next step.
                  2.     Check statistics on the packets sent to the CPU.
                         Run the display cpu-defend statistics command to check statistics on the
                         packets sent to the CPU. If a large number of protocol packets are discarded,
                         check whether these are attack packets (identified using the attack source
                         tracing function, for example). If so, use a filter or traffic policy to prevent
                         such packets from being sent to the CPU.

3.15.3 How Can the CPU Be Protected from DHCPv6
Messages?
                  Run the display cpu-defend statistics command to check statistics on packets
                  controlled by CPCAR. If a large number of DHCPv6 messages are discarded, check
                  whether IPv6 is required. If IPv6 is not required, run the cpu-defend policy
                  command to configure an attack defense policy to directly discard DHCPv6
                  messages.

3.15.4 How Can I Handle Excessive ARP Reply Packets Sent to
the CPU?
                  Excessive ARP Reply packets overload the CPU. To check whether excess ARP Reply
                  packets are sent to the CPU, run the display cpu-defend configuration packet-
                  type arp-reply all or display cpu-defend statistics packet-type arp-reply all
                  command.
                  In the display cpu-defend statistics packet-type arp-reply all command output,
                  if the value of the Drop (Bytes) field is large, excess ARP Reply packets are sent to
                  the CPU.
                  In this case, adjust the CPCAR value for ARP Reply packets. If the CPU is attacked,
                  obtain the packet header or enable debugging to check the attack source and
                  configure a filter to block the attack source.

                          NOTE

