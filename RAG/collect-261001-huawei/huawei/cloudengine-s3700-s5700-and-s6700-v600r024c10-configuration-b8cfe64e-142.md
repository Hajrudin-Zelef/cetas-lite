---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-142
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-12-26", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15823, 15939]
sha256: d2904d95eb88232df43fc5a9f54310a64ab3f7efb3b71addc343e4ae207e8d58
---

                    # Configure sampling for both incoming and outgoing traffic on the interface
                    configured with sa enable. The recommended sampling rate is 256.
                    [DeviceB] interface 10ge 1/0/2
                    [DeviceB-10GE1/0/2] netstream inbound ip
                    [DeviceB-10GE1/0/2] netstream outbound ip
                    [DeviceB-10GE1/0/2] netstream record sac_v4 ip inbound
                    [DeviceB-10GE1/0/2] netstream record sac_v4 ip outbound
                    [DeviceB-10GE1/0/2] netstream record sac_vxlan vxlan inner-ip inbound
                    [DeviceB-10GE1/0/2] netstream record sac_vxlan vxlan inner-ip outbound
                    [DeviceB-10GE1/0/2] netstream sampler ip random-packets 256 inbound
                    [DeviceB-10GE1/0/2] netstream sampler ip random-packets 256 outbound
                    [DeviceB-10GE1/0/2] quit
                    [DeviceB] interface 10ge 1/0/3
                    [DeviceB-10GE1/0/3] netstream inbound ip
                    [DeviceB-10GE1/0/3] netstream outbound ip
                    [DeviceB-10GE1/0/3] netstream record sac_v4 ip inbound
                    [DeviceB-10GE1/0/3] netstream record sac_v4 ip outbound
                    [DeviceB-10GE1/0/3] netstream record sac_vxlan vxlan inner-ip inbound
                    [DeviceB-10GE1/0/3] netstream record sac_vxlan vxlan inner-ip outbound
                    [DeviceB-10GE1/0/3] netstream sampler ip random-packets 256 inbound
                    [DeviceB-10GE1/0/3] netstream sampler ip random-packets 256 outbound
                    [DeviceB-10GE1/0/3] quit

                    ----End

Verifying the Configuration
                    After the preceding configurations are complete, check the experience assurance
                    configuration on DeviceB.
                    # Check the interfaces on which application identification is enabled.
                    <DeviceB> display interface sa-configuration
                    ------------------
                    InterfaceName
                    ------------------
                    10GE 1/0/2
                    10GE 1/0/3
                    ------------------


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 290
QoS Configuration
QoS Configuration                                                                  13 Experience Assurance Configuration


                    # Check the global configuration of application identification.
                    <DeviceB> display sa global configuration
                    Application statistic                   : Enable
                    SA flow-table aging-time (sec)               : 300
                    SA flow-table statistic report-cycle (sec)     : 120
                    SA flow-table usage-alarm lower percentage          : 50
                    SA flow-table usage-alarm upper percentage           : 80

                    # Check information about the application identification flow table.
                    <DeviceB> display sa flow-table all slot 1
                    This operation may take a long time. Please wait....done.
                    ------------------------------------------------------------
                     Source IPv4 address                 : 192.168.1.2
                     Source port                   : 2354
                     Destination IPv4 address              : 192.168.1.1
                     Destination port                 : 53
                     Protocol                    : TCP
                     VPN instance                    : _public_
                     Vrf ID                     :0
                     Create time                   : 2024-12-26 16:31:15
                     Latest update time                 : 2024-12-26 16:49:15
                     Application name                    : DNS
                     Application ID                 : 425
                     Total ingress packets              : 246558976
                     Total ingress bytes              : 2337919488
                     Total egress packets               :0
                     Total egress bytes               :0
                    ------------------------------------------------------------


Configuration Script
                    DeviceB
                    #
                    sysname DeviceB
                    #
                    vlan batch 10 20 30
                    #
                    interface 10GE1/0/1
                     port link-type trunk
                     port trunk allow-pass vlan 30
                    #
                    interface 10GE1/0/2
                     port link-type access
                     port default vlan 10
                     sa enable
                     traffic-policy p1 inbound
                     netstream inbound ip
                     netstream outbound ip
                     netstream record sac_v4 ip inbound
                     netstream record sac_v4 ip outbound
                     netstream record sac_vxlan vxlan inner-ip inbound
                     netstream record sac_vxlan vxlan inner-ip outbound
                     netstream sampler ip random-packets 256 inbound
                     netstream sampler ip random-packets 256 outbound
                    #
                    interface 10GE1/0/3
                     port link-type access
                     port default vlan 20
                     sa enable
                     traffic-policy p1 inbound
                     netstream inbound ip
                     netstream outbound ip
                     netstream record sac_v4 ip inbound
                     netstream record sac_v4 ip outbound
                     netstream record sac_vxlan vxlan inner-ip inbound
                     netstream record sac_vxlan vxlan inner-ip outbound


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                      291
QoS Configuration
QoS Configuration                                                       13 Experience Assurance Configuration

