---
id: collect-261001-huawei/huawei/questions-46265-1c4b56a7
title: "Traffic Statistics per VLAN"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-46265-1c4b56a7.md
source_anchor: ""
source_lines: [1, 19]
sha256: f18ce48211fbd13c1f6fbe68b4096140b3cd5c9ff9d515622f15da9498ef4f3a
---

# Traffic Statistics per VLAN

Question asked by Shinomoto Asakura | Score: 1 | Tags: vlan, huawei, traffic, statistics

## Question

I'm using Huawei s6720 and I would like to monitor traffic per vlan, example: 

```
`# 
Interface GigabitEthernet0/0/6 
port link-type trunk 
port trunk allow-pass vlan 200 300 400 500 
# 
`

I can verify the traffic on interface, however I don't know the traffic on VLANs.

Thanks in advance
