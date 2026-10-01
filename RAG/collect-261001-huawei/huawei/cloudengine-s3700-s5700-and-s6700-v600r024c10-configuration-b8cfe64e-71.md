---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-71
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [10046, 10221]
sha256: d0d030ccf7f95ed07c7ae0757d80828a20ebfa169f554d2fe611741cde5a3e7e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     Tu    Whether the            Whether the      QoS                Wheth     Applica
                     nn    Ingress Node           Egress Node      Scheduling         er the    tion
                     el    Trusts the Original    Trusts the       Behavior on        Origin    Scenari
                     Mo    CoS Value Carried      EXP Value or     the Egress         al CoS    o
                     de    in an IP Packet        the Original     Node               Value
                                                  CoS Value                           Carried
                                                  Carried in an                       in an
                                                  IP Packet                           IP
                                                                                      Packet
                                                                                      Chang
                                                                                      es

                     Sh    Untrusted. A new       The original     Scheduling         No
                     ort   value can be           CoS value        based on
                     pip   assigned to the EXP    carried in an    mapping of
                     e     field in the outer     IP packet is     the original
                           MPLS label.            trusted and      CoS value
                                                  retained.        carried in a
                                                                   packet




12.3 Configuration Precautions for MPLS QoS

12.4 Default Settings for MPLS QoS
                    Table 12-2 describes the default mappings of EXP values to PHBs and colors when
                    packets enter the device.

                    Table 12-2 Default mappings of EXP values to PHBs and colors in the inbound
                    direction of the device
                     EXP Value                     PHB                        Color

                     0                             BE                         Green

                     1                             AF1                        Green

                     2                             AF2                        Green

                     3                             AF3                        Green

                     4                             AF4                        Green

                     5                             EF                         Green

                     6                             CS6                        Green

                     7                             CS7                        Green




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           184
QoS Configuration
QoS Configuration                                                          12 MPLS QoS Configuration


                    Table 12-3 describes the default mappings of PHBs and colors to EXP values when
                    packets leave the device.

                    Table 12-3 Default mappings of PHBs and colors to EXP values in the outbound
                    direction of the device
                     PHB                         Color                     EXP Value

                     BE                          Green                     0

                     BE                          Yellow                    0

                     BE                          Red                       0

                     AF1                         Green                     1

                     AF1                         Yellow                    1

                     AF1                         Red                       1

                     AF2                         Green                     2

                     AF2                         Yellow                    2

                     AF2                         Red                       2

                     AF3                         Green                     3

                     AF3                         Yellow                    3

                     AF3                         Red                       3

                     AF4                         Green                     4

                     AF4                         Yellow                    4

                     AF4                         Red                       4

                     EF                          Green                     5

                     EF                          Yellow                    5

                     EF                          Red                       5

                     CS6                         Green                     6

                     CS6                         Yellow                    6

                     CS6                         Red                       6

                     CS7                         Green                     7

                     CS7                         Yellow                    7

                     CS7                         Red                       7




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                        185
QoS Configuration
QoS Configuration                                                                      12 MPLS QoS Configuration




12.5 Configuring a DiffServ Mode
Prerequisites
                    Before configuring a DiffServ mode, you have completed the following tasks:

                    ●   Complete basic configurations of the MPLS network.
                    ●   Configure the VPN service.


Context
                    MPLS DiffServ has three tunnel modes: uniform, pipe, and short pipe.

                    ●   In uniform mode, the original priority carried in a packet is trusted when the
                        packet enters an MPLS network. This mode applies to a scenario where
                        priorities of different services in a VPN are differentiated.
                    ●   In pipe or short pipe mode, you can specify the EXP value of the private
                        network label when packets enter the MPLS network. These modes apply to a
                        scenario where priorities of different services in different VPNs are
                        differentiated.

                    If you do not want to change the priority carried in original packets, you are
                    advised to use pipe or short pipe mode. This is because, in uniform mode, the
                    priority carried in original packets may be changed. In addition, the egress node
                    selects PHBs based on EXP values of packets in uniform and pipe modes, and
                    selects PHBs based on DSCP values of packets in short pipe mode.


Procedure
                    ●   Configuring a DiffServ mode for L3VPN
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VPN instance view.
                             ip vpn-instance vpn-instance-name

                        c.   Select a type and enter the corresponding address family view.

                             ▪    Enter the IPv4 address family view.
                                  ipv4-family

                             ▪    Enter the IPv6 address family view.
                                  ipv6-family

                        d.   Configure a DiffServ mode for the VPN instance.
                             diffserv-mode { pipe { mpls-exp exp-value | domain ds-name } | short-pipe [ mpls-exp exp-
                             value ] domain ds-name | uniform [ domain ds-name ] }

                             By default, the DiffServ mode of a VPN instance is uniform and the
                             default domain is used.




