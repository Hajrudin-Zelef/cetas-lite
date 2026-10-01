---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-16
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [1453, 1572]
sha256: 8b447861e9977a6c6ad39093792be2c4869f6d4e90344854b830d129ca807a66
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

3.5.1 Understanding Port Attack Defense
                  Port attack defense is an anti-Denial of Service (DoS) method. It defends against
                  attacks based on ports, preventing service interruptions caused by a failure to send
                  protocol packets from normal ports to the CPU, as protocol packets from attacked
                  ports may exhaust the bandwidth.
                  The process of port attack defense is as follows:
                  1.     The device analyzes packets by port, and counts the number of protocol
                         packets received on the port where port attack defense is applied.
                  2.     The device considers a port to be under attack if the number of packets sent
                         to the CPU in a unit time exceeds the rate threshold.
                  3.     When detecting an attack, the device sends a log and moves the packets
                         within the protocol rate limit (similar to the CPCAR value of protocol packets
                         in an attack defense policy) to a low-priority queue, before sending them to
                         the CPU for processing. The device discards the excess packets if any.
                         The rate limiting actions of port attack defense have less of an impact on
                         services than the punishment actions in attack source tracing.
                  Port attack defense also provides whitelist, aging detection, and port attack
                  defense event reporting functions.
                  Whitelist: After legitimate users are added to a whitelist, their packets are not
                  processed by the device based on port attack defense and are instead sent to the
                  CPU for processing. A whitelist can be flexibly configured based on ACLs or ports.
                  Aging detection: Upon detection of a port under attack, the device continuously
                  moves attack packets on the port to low-priority queues for processing within the
                  aging time. When the aging time expires, the device calculates the protocol packet
                  rate on the port again. If the rate is still above the rate threshold (meaning the
                  attack persists), the device continues moving packets to low-priority queues.
                  Otherwise, the device sends them to the CPU.
                  A proper aging time for port attack defense should be set based on the CPU usage
                  and service running status. If the aging time is too short, the device frequently
                  detects protocol packet rates on ports, consuming CPU resources unnecessarily. If
                  the aging time is too long, protocol packets cannot be promptly processed by the
                  CPU, impacting services.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            21
Security Configuration
Security Configuration                                                         3 Local Attack Defense Configuration


                  Port attack defense event reporting: When a port is attacked, the device reports
                  an event to the administrators so they can take appropriate measures to protect
                  the device.

3.5.2 Configuring Port Attack Defense

Context
                  The port attack defense function effectively limits the number of packets sent
                  from a port to the CPU.

Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Create an attack defense policy and enter the attack defense policy view.
                  cpu-defend policy policy-name

         Step 3 Enable port attack defense.
                  auto-port-defend enable

         Step 4 (Optional) Disable port attack defense for a specified protocol.
                  auto-port-defend protocol { 8021x | 8021x-ident | 8021x-start | arp-request | arp-request-uc | arp-reply
                  | dhcp-discovery | dhcp-request | dhcp-reply | dhcpv6-discovery | dhcpv6-reply | dhcpv6-request | icmp |
                  igmp | ip-fragment | isis | isis-overlay | lacp | nac-arp-reply | nac-arp-request | nac-dhcp | nac-dhcp-
                  discovery |nac-dhcpv6 | nac-dhcpv6-discovery | nac-nd | nd | ospf | ospf-hello | ospf-overlay | ospf-hello-
                  overlay | ospfv3 | ospfv3-overlay | pim | pim-mc | xldp| vrrp | vrrp6 } disable

                  The S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-H-V2, and S5735I-S-V2 do not
                  support dhcp-discovery, dhcpv6-discovery, isis-overlay, nac-dhcp-discovery, nac-
                  dhcpv6-discovery, ospf-overlay, ospf-hello-overlay, or ospfv3-overlay.

                  The S5735-L-V2, S5735E-L-V2, S5735R-L-V2, and S5735I-L-V2 do not support dhcp-
                  discovery, dhcpv6-discovery, isis, isis-overlay, nac-dhcp-discovery, nac-dhcpv6-
                  discovery, ospf-overlay, ospf-hello-overlay, or ospfv3-overlay.

                  On the S6780-H and S6750-H, dhcp-discovery and nac-dhcp-discovery are not
                  supported.

                  The S3710-H does not support dhcp-discovery, dhcpv6-discovery, isis, isis-overlay,
                  nac-dhcp-discovery, nac-dhcpv6-discovery, ospf, ospf-hello, ospfv3, ospf-overlay,
                  ospf-hello-overlay, or ospfv3-overlay.

                  The S5755-S does not support dhcp-discovery, nac-dhcp-discovery, or nac-dhcpv6-
                  discovery.

                  If only a small proportion of all packets exceeding the rate threshold are attack
                  packets, you can disable port attack defense for specific protocols, preventing
                  normal services from being impacted due to excessive rate limiting.

         Step 5 (Optional) Set the rate threshold for port attack defense.
                  auto-port-defend protocol { 8021x | 8021x-ident | 8021x-start | arp-request | arp-request-uc | arp-reply
                  | dhcp-discovery | dhcp-request | dhcp-reply | dhcpv6-discovery | dhcpv6-reply | dhcpv6-request | icmp |
                  igmp | ip-fragment | isis | isis-overlay | lacp | nac-arp-reply | nac-arp-request | nac-dhcp | nac-dhcp-
                  discovery |nac-dhcpv6 | nac-dhcpv6-discovery | nac-nd | nd | ospf | ospf-hello | ospf-overlay | ospf-hello-
                  overlay | ospfv3 | ospfv3-overlay | pim | pim-mc | xldp| vrrp | vrrp6 } threshold threshold-value


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               22
Security Configuration
Security Configuration                                                          3 Local Attack Defense Configuration


                  The S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-H-V2, and S5735I-S-V2 do not
                  support dhcp-discovery, dhcpv6-discovery, isis-overlay, nac-dhcp-discovery, nac-
                  dhcpv6-discovery, ospf-overlay, ospf-hello-overlay, or ospfv3-overlay.

                  The S5735-L-V2, S5735E-L-V2, S5735R-L-V2, and S5735I-L-V2 do not support dhcp-
                  discovery, dhcpv6-discovery, isis, isis-overlay, nac-dhcp-discovery, nac-dhcpv6-
                  discovery, ospf-overlay, ospf-hello-overlay, or ospfv3-overlay.

                  On the S6780-H and S6750-H, dhcp-discovery and nac-dhcp-discovery are not
                  supported.

                  The S3710-H does not support dhcp-discovery, dhcpv6-discovery, isis, isis-overlay,
                  nac-dhcp-discovery, nac-dhcpv6-discovery, ospf, ospf-hello, ospfv3, ospf-overlay,
                  ospf-hello-overlay, or ospfv3-overlay.

                  The S5755-S does not support dhcp-discovery, nac-dhcp-discovery, or nac-dhcpv6-
                  discovery.

