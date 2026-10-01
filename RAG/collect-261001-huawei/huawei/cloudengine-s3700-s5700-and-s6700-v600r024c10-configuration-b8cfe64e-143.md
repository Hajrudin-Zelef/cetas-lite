---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-143
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [15940, 15985]
sha256: bc8d695bf5e38aaf09ad1ad1680fa2d3bd968f76407e2e2c4dc6346473e41e1d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                     netstream sampler ip random-packets 256 inbound
                     netstream sampler ip random-packets 256 outbound
                    #
                    interface vlanif 10
                     ip address 192.168.10.1 255.255.255.0
                    #
                    interface vlanif 20
                     ip address 192.168.20.1 255.255.255.0
                    #
                    interface vlanif 30
                     ip address 192.168.100.1 255.255.255.0
                    #
                    sa application-statistic enable
                     assign forward enp netstream enable
                     netstream timeout ip active 300
                     netstream timeout vxlan inner-ip active 300
                     netstream timeout ip inactive 180
                     netstream timeout vxlan inner-ip inactive 180
                     netstream timeout ip tcp-session
                     netstream timeout vxlan inner-ip tcp-session
                     netstream record sac_v4 ip
                     collect counter bytes
                     collect counter packets
                     collect interface sampler-info
                     match ip destination-address
                     match ip destination-port
                     match ip protocol
                     match ip source-address
                     match ip source-port
                     netstream record sac_vxlan vxlan inner-ip
                     collect counter bytes
                     collect counter packets
                     collect interface sampler-info
                     match inner-ip destination-address
                     match inner-ip destination-port
                     match inner-ip protocol
                     match inner-ip source-address
                     match inner-ip source-port
                    #
                    return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                              292

