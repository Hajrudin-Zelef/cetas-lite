---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-23
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [2442, 2624]
sha256: 4649421b3a0395583beb09f246a79bb9bcdc0f5ac690dcf0b3fc036f32b5657b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

3.10.1 Understanding Defense Against TCP SYN Flood Attacks
                  A TCP SYN flood attack exploits the vulnerability of the TCP three-way handshake.
                  As shown in Figure 3-9, during the TCP three-way handshake, when the receiver
                  (target device) receives the initial SYN packet from the sender (attacker), it
                  returns a SYN+ACK packet to the sender. The connection remains in a half-open
                  state while the receiver waits for the final ACK packet from the sender. If the
                  receiver does not receive the ACK packet, it retransmits a SYN+ACK packet to the
                  sender. Finally, after several retransmission attempts, the receiver shuts down the
                  session and then updates the session in memory. The period from the first SYN
                  +ACK packet being sent to session teardown is approximately 30 seconds.
                  During this period, an attacker may send thousands of SYN packets to all open
                  interfaces, and never responds to SYN+ACK packets from the receiver. This causes
                  memory overloading on the receiver and prevents the receiver from accepting new
                  connection requests. The receiver then disconnects all existing connections.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 39
Security Configuration
Security Configuration                                           3 Local Attack Defense Configuration


                  Figure 3-9 TCP SYN flood attack




                  Figure 3-10 shows how the device defends against TCP SYN attacks. With TCP
                  SYN flood attack defense enabled, the device rate-limits TCP SYN packets to
                  ensure that system resources are not exhausted if an attack occurs.


                  Figure 3-10 Defense against TCP SYN flood attacks




3.10.2 Configuring Defense Against TCP SYN Flood Attacks

Context
                  With defense against TCP SYN flood attacks enabled, the device rate-limits the
                  received TCP SYN packets, discarding any that exceeds the limit.

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                            40
Security Configuration
Security Configuration                                                3 Local Attack Defense Configuration


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable defense against TCP SYN flood attacks.
                  anti-attack tcp-syn enable

                         NOTE

                  You can also run the anti-attack enable command in the system view to enable attack defense
                  against all attack packets, including TCP SYN flood attack packets.

         Step 3 Configure the limit rate for TCP SYN packets.
                  anti-attack tcp-syn car cir cir-num

                  ----End

Verifying the Configuration
                  Run the display anti-attack statistics tcp-syn command to check statistics
                  relating to defense against TCP SYN flood attacks.


3.11 Configuring Defense Against UDP Flood Attacks

3.11.1 Understanding Defense Against UDP Flood Attacks
                  A UDP flood attack is an attack in which a large number of UDP packets are sent
                  to a target device within a short time, causing the target device to be busy with
                  these UDP packets and fail to process normal services. UDP flood attacks are
                  classified into two types:
                  ●      Fraggle attack
                         Figure 3-11 shows how a Fraggle attack works. An attacker sends UDP
                         packets where the source address is the target device's address, the
                         destination address is the broadcast address of the target network, and the
                         destination port is port 7. If multiple hosts use UDP echo services on the
                         broadcast network, the target device receives excessive response packets from
                         these hosts and becomes occupied while processing these packets.
                         The device with flood attack defense enabled considers packets from UDP
                         port 7 as attack packets and discards them.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 41
Security Configuration
Security Configuration                                                3 Local Attack Defense Configuration


                         Figure 3-11 Fraggle attack




                  ●      UDP diagnosis port attack
                         An attacker sends a large number of UDP request packets to the target
                         device's UDP diagnosis ports (such as 7-echo, 13-daytime, and 19-Chargen),
                         causing flooding and consuming network bandwidth resources. In addition, as
                         the target device consumes CPU resources when responding to these UDP
                         request packets, the system becomes overloaded and cannot process normal
                         services.
                         The device with flood attack defense enabled considers packets from UDP
                         ports 7, 13, and 19 as attack packets and discards them.

3.11.2 Configuring Defense Against UDP Flood Attacks

Context
                  With defense against UDP flood attacks enabled, the device discards the packets
                  received from ports 7, 13, and 19.


Procedure
         Step 1 Enter the system view.
                  system-view

         Step 2 Enable defense against UDP flood attacks.
                  anti-attack udp-flood enable

                         NOTE

                  You can also run the anti-attack enable command in the system view to enable attack defense
                  against all attack packets, including UDP flood attack packets.

                  ----End



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                  42
Security Configuration
Security Configuration                                             3 Local Attack Defense Configuration


Verifying the Configuration
                  Run the display anti-attack statistics udp-flood command to check statistics
                  relating to defense against UDP flood attacks.


3.12 Configuring Defense Against ICMP Flood Attacks

3.12.1 Understanding Defense Against ICMP Flood Attacks
                  A network administrator typically monitors a network and rectifies faults using the
                  ping tool as follows:
                  1.     The source host sends an ICMP Echo message to a target device.
                  2.     Upon receiving the ICMP Echo message, the target device sends an ICMP Echo
                         Reply message to the source host.
                  As shown in Figure 3-12, if the target device receives many ICMP Echo messages
                  from an attacker, it becomes occupied and unable to process other data packets.
                  As a result, normal services are affected.

                  Figure 3-12 ICMP flood attack




                  Figure 3-13 shows how the device resists ICMP flood attacks. With defense
                  against ICMP flood attacks enabled, the device rate-limits ICMP messages to
                  ensure that system resources are not exhausted when the device is attacked.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           43
Security Configuration
Security Configuration                                                3 Local Attack Defense Configuration


                  Figure 3-13 Defense against ICMP flood attacks




