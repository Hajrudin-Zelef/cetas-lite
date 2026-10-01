---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-21
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [2148, 2282]
sha256: b0b8094a77aeb4fb938902bf28a7106d74108442364278da16d9e5e99f67b051
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

Flood Attack from IP Null Payload Packets
                  An IP null payload packet has only a 20-byte IP header, and no data section.
                  Attackers often construct a large number of packets with only IP headers, known
                  as IP null payload packets, which constitute a flood attack. Processing such
                  packets causes devices to encounter an error or even crash, impacting services.

                  With defense against malformed packet attacks enabled, the device directly
                  discards the received IP null payload packets, thereby preventing service
                  interruption.

Attack from IGMP Null Payload Packets
                  A normal IGMP packet consists of a 20-byte IP header and an 8-byte data part,
                  equating to a total length of 28 bytes. An IGMP null payload packet, on the other
                  hand, consists of less than 28 bytes. Processing IGMP null payload packets causes
                  devices to encounter an error or even crash, impacting services.

                  With defense against malformed packet attacks enabled, the device directly
                  discards the received IGMP null payload packets, thereby preventing service
                  interruption.

LAND Attack
                  A Local Area Network Denial (LAND) attacker targets the defects in the three-way
                  handshake mechanism of TCP, sending a SYN packet in which both the source and
                  destination addresses are the target host's address and the source and destination
                  ports are the same. After receiving the SYN packet, the target host creates a null
                  TCP connection by using its own address as both the source and destination
                  addresses, and retains this connection until it times out. The target host will create
                  many null TCP connections after receiving a large number of such SYN packets,
                  wasting network resources or even crashing the system.

                  With defense against malformed packet attacks enabled, the device checks source
                  and destination addresses or ports in TCP SYN packets, and considers TCP SYN

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                              33
Security Configuration
Security Configuration                                                 3 Local Attack Defense Configuration


                  packets with the same source and destination addresses or ports as malformed
                  packets and discards them.


Smurf Attack
                  An attacker sends an ICMP Request packet with a source address as the target
                  host's address and a destination address as the broadcast address of the target
                  network. As such, all hosts on the target network receive the ICMP Request
                  packet, after which they send ICMP Reply packets to the target host. This
                  inevitably leads to the target host receiving an excessive number of packets,
                  consuming excessive resources crashing the system or network.

                  With defense against malformed packet attacks enabled, the device checks
                  whether the destination addresses in ICMP Request packets are the broadcast or
                  subnet broadcast addresses, discarding them if so.


Attack from Packets with Invalid TCP Flag Bits
                  A TCP packet contains six flag bits: URG, ACK, PSH, RST, SYN, and FIN. Different
                  systems respond differently to the combination of these flag bits.

                  ●      If the six flag bits are all 1s, the attack is a Christmas tree attack. A device
                         subject to a Christmas tree attack may crash.
                  ●      An attacker sends a TCP packet in which SYN and FIN are 1 to a target host. If
                         the receiving interface is disabled, the receiver replies with an RST | ACK
                         message. If the receiving interface is enabled, the receiver replies with an SYN
                         | ACK message. Such an attack is used to detect whether a host is online or
                         offline and whether an interface is enabled or disabled.
                  ●      An attacker sends a TCP packet in which the six flag bits are all 0s to a target
                         host. If the receiving interface is disabled, the receiver replies with an RST |
                         ACK message, which can be used by the attacker to detect whether the host is
                         online or offline. If the receiving interface is enabled, the target host does not
                         respond if running Linux or UNIX but replies with an RST | ACK message if
                         running Windows. This attack is used to detect the type of operating system
                         on the target host.

                  With defense against malformed packet attacks enabled, the device checks each
                  flag bit in TCP packets to prevent attacks from packets with invalid TCP flag bits. If
                  any of the following conditions is met, the device discards the TCP packets:

                  ●      The six flag bits are all 1s.
                  ●      Both SYN and FIN are 1.
                  ●      The six flag bits are all 0s.

3.8.2 Configuring Defense Against Malformed Packet Attacks

Context
                  With defense against malformed packet attacks enabled, the device analyzes the
                  received packets sent to the CPU and determines whether the packets are one of
                  the several types of malformed packets, discarding them if so.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  34
Security Configuration
Security Configuration                                                3 Local Attack Defense Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable defense against malformed packet attacks.
                  anti-attack abnormal enable

                         NOTE

                  You can also run the anti-attack enable command in the system view to enable attack defense
                  against all attack packets, including malformed packets.

                  ----End

Verifying the Configuration
                  Run the display anti-attack statistics abnormal command to check statistics
                  relating to defense against malformed packet attacks on the device.


3.9 Configuring Defense Against Fragmentation
Attacks

3.9.1 Understanding Defense Against Fragmentation Attacks
                  A fragmentation attack is an attack in which error fragments are sent to a target
                  device, which then crashes, restarts, or consumes a large amount of CPU resources
                  when processing such fragments, ultimately impacting services. Fragmentation
                  attack defense enables a device to detect packet fragments in real time and
                  discard or rate-limit them to protect the device.
                  Fragmentation attacks are classified into the following types.

Excess-Fragment Attack
                  The offset of IP packets is measured in units of 8-byte blocks. Normally, an IP
                  header has 20 bytes and the maximum payload of an IP packet is 65515 bytes.
                  Therefore, an IP packet can be fragmented into up to 8189 (65515/8) fragments,
                  above which the device would consume a large amount of CPU resources when
                  reassembling packets.
                  The device considers a packet with over 8189 fragments malicious and discards all
                  of its fragments.

