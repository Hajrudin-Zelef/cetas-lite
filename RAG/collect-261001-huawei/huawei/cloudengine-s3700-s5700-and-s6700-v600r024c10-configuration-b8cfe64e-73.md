---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-73
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [10370, 10522]
sha256: 022bbf49a121c77ad90be6fb6726cc306e006ea789633e004e52ff2d890aa2f9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Procedure
                    ●   Configure the ingress device to perform priority mapping for packets
                        transmitted over a public network tunnel based on the EXP value.
                        a.   Enter the system view.
                             system-view

                        b.   Configure the ingress device to perform priority mapping for packets
                             transmitted over a public network tunnel based on the EXP value.
                             mpls-qos ingress { use vpn-label-exp | trust upstream { ds-name | default | none } }

                             By default, priority mapping is performed for packets transmitted over a
                             public network tunnel based on the EXP value according to the settings in
                             the default domain.
                    ●   Configure the transit device to perform priority mapping for packets
                        transmitted over a public network tunnel based on the EXP value.
                        a.   Enter the system view.
                             system-view

                        b.   Configure the transit device to perform priority mapping for packets
                             transmitted over a public network tunnel based on the EXP value.
                             mpls-qos transit trust upstream { none | default | ds-name }

                             By default, priority mapping is performed for packets transmitted over a
                             public network tunnel based on the EXP value according to the settings in
                             the default domain.
                    ●   Configure the egress device to perform priority mapping for packets
                        transmitted over a public network tunnel based on the EXP value.
                        a.   Enter the system view.
                             system-view


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          189
QoS Configuration
QoS Configuration                                                                       12 MPLS QoS Configuration


                        b.    Configure the egress device to perform priority mapping for packets
                              transmitted over a public network tunnel based on the EXP value.
                              mpls-qos egress trust upstream { none | default | ds-name }

                              By default, priority mapping is performed for packets transmitted over a
                              public network tunnel based on the EXP value according to the settings in
                              the default domain.
                    ----End


12.7 Checking the MPLS QoS Configuration
Procedure
                    ●   Run the display diffserv domain [ brief | name ds-domain-name ] command
                        to check the DiffServ domain configuration.
                    ●   Run the display qos configuration interface [ interface-type interface-
                        number ] command to check all QoS configurations on an interface.
                    ----End


12.8 Example for Configuring MPLS QoS
Networking Requirements
                    Enterprises A and B use BGP/MPLS IP VPN to connect their headquarters and
                    branches. In Figure 12-8, CE1 and CE3 connect to the headquarters and branch of
                    enterprise A, and CE2 and CE4 connect to the headquarters and branch of
                    enterprise B. Enterprise A uses VPN vpna, and enterprise B uses VPN vpnb.
                    Enterprise A has a high service level and requires better QoS guarantee.

                    Figure 12-8 Network diagram of MPLS QoS
                         NOTE

                        In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                        10GE1/0/3, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    190
QoS Configuration
QoS Configuration                                                              12 MPLS QoS Configuration




Configuration Roadmap
                    Configure MPLS QoS on PE1 and PE2, enable the pipe mode for VPNs vpna and
                    vpnb, and set the MPLS EXP values of VPNs vpna and vpnb to 4 and 3,
                    respectively, to provide better QoS guarantee for services of enterprise A.


Procedure
         Step 1 Configure OSPF on the MPLS backbone network so that the PEs and P on the
                backbone network can communicate with each other.

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 32
                    [PE1-LoopBack1] quit
                    [PE1] interface 10ge 1/0/3
                    [PE1-10GE1/0/3] undo portswitch
                    [PE1-10GE1/0/3] ip address 172.16.1.1 24
                    [PE1-10GE1/0/3] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 172.16.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure P.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                        191
QoS Configuration
QoS Configuration                                                              12 MPLS QoS Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname P
                    [P] interface loopback 1
                    [P-LoopBack1] ip address 2.2.2.9 32
                    [P-LoopBack1] quit
                    [P] interface 10ge 1/0/1
                    [P-10GE1/0/1] undo portswitch
                    [P-10GE1/0/1] ip address 172.16.1.2 24
                    [P-10GE1/0/1] quit
                    [P] interface 10ge 1/0/2
                    [P-10GE1/0/2] undo portswitch
                    [P-10GE1/0/2] ip address 172.17.1.1 24
                    [P-10GE1/0/2] quit
                    [P] ospf
                    [P-ospf-1] area 0
                    [P-ospf-1-area-0.0.0.0] network 172.16.1.0 0.0.0.255
                    [P-ospf-1-area-0.0.0.0] network 172.17.1.0 0.0.0.255
                    [P-ospf-1-area-0.0.0.0] network 2.2.2.9 0.0.0.0
                    [P-ospf-1-area-0.0.0.0] quit
                    [P-ospf-1] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] interface loopback 1
                    [PE2-LoopBack1] ip address 3.3.3.9 32
                    [PE2-LoopBack1] quit
                    [PE2] interface 10ge 1/0/3
                    [PE2-10GE1/0/3] undo portswitch
                    [PE2-10GE1/0/3] ip address 172.17.1.2 24
                    [PE2-10GE1/0/3] quit
                    [PE2] ospf
                    [PE2-ospf-1] area 0
                    [PE2-ospf-1-area-0.0.0.0] network 172.17.1.0 0.0.0.255
                    [PE2-ospf-1-area-0.0.0.0] network 3.3.3.9 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] quit
                    [PE2-ospf-1] quit

