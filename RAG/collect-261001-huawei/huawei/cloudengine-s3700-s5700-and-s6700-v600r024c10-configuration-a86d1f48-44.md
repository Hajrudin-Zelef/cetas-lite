---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-a86d1f48-44
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48.md
source_anchor: ""
source_lines: [5605, 5777]
sha256: 1559d16cfa4b7c50102ad5fed3ec1df4dc0e9e9234bcfd7b34f0135e149dcb8b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--a86d1f48

6.5.6 Example for Configuring Port Security
Networking Requirements
                  As shown in Figure 6-1, PC1, PC2, and PC3 can communicate with each other in
                  VLAN 10, and connect to the company network through DeviceA. For security

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         99
Security Configuration
Security Configuration                                                              6 Port Security Configuration


                  purposes, only PC1, PC2, and PC3 can access the company network, and external
                  users cannot access the company network.

                  Figure 6-1 Network diagram of port security
                          NOTE

                         In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                         10GE1/0/3, respectively.




Configuration Roadmap
                  1.     Configure a VLAN to enable employee PCs to communicate with each other.
                  2.     Enable port security and limit the number of MAC addresses learned on an
                         interface, so that external users cannot access the company network.

Procedure
         Step 1 Configure a VLAN.
                  <HUAWEI> system-view
                  [HUAWEI] sysname DeviceA
                  [DeviceA] vlan batch 10
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10GE1/0/1] portswitch
                  [DeviceA-10GE1/0/1] port link-type access
                  [DeviceA-10GE1/0/1] port default vlan 10
                  [DeviceA-10GE1/0/1] quit
                  [DeviceA] interface 10ge 1/0/2
                  [DeviceA-10GE1/0/2] portswitch
                  [DeviceA-10GE1/0/2] port link-type access
                  [DeviceA-10GE1/0/2] port default vlan 10
                  [DeviceA-10GE1/0/2] quit
                  [DeviceA] interface 10ge 1/0/3
                  [DeviceA-10GE1/0/3] portswitch
                  [DeviceA-10GE1/0/3] port link-type access
                  [DeviceA-10GE1/0/3] port default vlan 10
                  [DeviceA-10GE1/0/3] quit

         Step 2 Configure port security.
                  [DeviceA] interface 10ge 1/0/1
                  [DeviceA-10GE1/0/1] port-security enable maximum 1
                  [DeviceA-10GE1/0/1] port-security mac-address sticky


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     100
Security Configuration
Security Configuration                                                        6 Port Security Configuration

                  [DeviceA-10GE1/0/1] quit
                  [DeviceA] interface 10ge 1/0/2
                  [DeviceA-10GE1/0/2] port-security enable maximum 1
                  [DeviceA-10GE1/0/2] port-security mac-address sticky
                  [DeviceA-10GE1/0/2] quit
                  [DeviceA] interface 10ge 1/0/3
                  [DeviceA-10GE1/0/3] port-security enable maximum 1
                  [DeviceA-10GE1/0/3] port-security mac-address sticky
                  [DeviceA-10GE1/0/3] quit

                  ----End

Verifying the Configuration
                  Only PC1, PC2, and PC3 connected to the three interfaces on DeviceA can access
                  the company network.

Configuration Scripts
                  #
                   sysname DeviceA
                  #
                  vlan batch 10
                  #
                  interface 10GE1/0/1
                   port link-type access
                   port default vlan 10
                   port-security enable maximum 1
                   port-security mac-address sticky
                  #
                  interface 10GE1/0/2
                   port link-type access
                   port default vlan 10
                   port-security enable maximum 1
                   port-security mac-address sticky
                  #
                  interface 10GE1/0/3
                   port link-type access
                   port default vlan 10
                   port-security enable maximum 1
                   port-security mac-address sticky
                  #
                  return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            101
Security Configuration
Security Configuration                                                          7 URPF Configuration




                                               7           URPF Configuration


                  7.1 Overview of URPF
                  7.2 Configuration Precautions for URPF
                  7.3 Understanding URPF
                  7.4 Default Settings for URPF
                  7.5 Configuring URPF


7.1 Overview of URPF
Definition
                  Unicast Reverse Path Forwarding (URPF) defends against network attacks
                  launched through source IP address spoofing.
                  URPF allows the device to search its Forwarding Information Base (FIB) table for a
                  route to the source IP address of a packet and checks whether the inbound
                  interface of the packet is the same as the outbound interface of the route. If no
                  route to the source IP address exists in the FIB table or the inbound interface of
                  the packet is different from the outbound interface of the matching route, the
                  packet is discarded. This effectively protects the device against malicious attacks
                  that change source IP addresses of packets.

Purpose
                  A Denial of Service (DoS) attack disables users from connecting to a server. Such
                  an attack aims to occupy excessive resources by sending a large number of valid
                  or forged connection requests, preventing authorized users from receiving
                  responses from the server. URPF effectively prevents DoS attacks that use spoofed
                  source IP addresses.
                  In Figure 7-1, PC_A sends request packets with the spoofed source IP address
                  10.2.2.2 to the server. If URPF is disabled on DeviceA, the server sends response
                  packets to PC_B (with the IP address 10.2.2.2) after receiving the request packets.
                  In this way, PC_A attacks both the server and PC_B by sending the request packets.
                  If URPF is enabled on DeviceA, DeviceA checks the inbound interface of the

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                          102
Security Configuration
Security Configuration                                                              7 URPF Configuration


                  packets received from PC_A and finds that packets with the source IP address
                  10.2.2.2 must reach DeviceA through interface 2. Therefore, DeviceA considers that
                  the source IP address of the packets is a spoofed IP address and discards the
                  packets. Packets from PC_B to the server, in contrast, are allowed to pass through
                  after passing the URPF check.

                  Figure 7-1 Preventing source address spoofing attacks through URPF




7.2 Configuration Precautions for URPF

7.3 Understanding URPF
Working Mode
                  On a complex network, the routing paths recorded on a remote device may be
                  different from those recorded on the local device. A URPF-enabled device on such
                  a network may discard packets received from valid paths. To solve this problem,
                  the device provides two URPF modes:

