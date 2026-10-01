---
id: collect-261001-huawei/huawei/questions-82585-bbcc4221
title: "What's the differences between sticky-config MAC & sticky MAC on Huawei SW？"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-82585-bbcc4221.md
source_anchor: ""
source_lines: [1, 30]
sha256: 6261b29ec3a89f12dced0a6e770041c75c147c8cf2c9a895146cc52f83271a40
---

# What's the differences between sticky-config MAC & sticky MAC on Huawei SW？

Question asked by Gilbert Wong | Score: 4 | Tags: mac-address, port-security, huawei

## Question

In Port-security configuration，

```
`port-security mac-address sticky-config xxxx-xxxx-xxx vlan x
port-security mac-address sticky xxxx-xxxx-xxx vlan x
`

1.All of them do not disappear after SW restarting. What's differences between the two commands？

2.Are sticky-config MAC & sticky MAC called security static MAC？

## Answer 1 (ACCEPTED)

By Zac67 | Score: 1

From the support articles I found

- port-security mac-address sticky

- port-security mac-address sticky-config

my understanding is that `sticky` is a basically dynamic mode where you can add MAC addresses manually as well. `sticky-config` is a static mode where there's no learning and all addresses need to be configured.

Both commands need to survive a reboot to make sense.
