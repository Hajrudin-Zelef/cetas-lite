---
id: collect-261001-huawei/huawei/questions-46959-2ab0e90c
title: "Does Cisco-IPSLA work with Huawei-NQA?"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-46959-2ab0e90c.md
source_anchor: ""
source_lines: [1, 41]
sha256: 874f6a0ff5b20b0169d938d6ac092d2e45367e05a867ba284ac6c1d7c7f3a737
---

# Does Cisco-IPSLA work with Huawei-NQA?

Question asked by Kiatbordin Jumratanet | Score: 1 | Tags: cisco, udp, monitoring, huawei

## Question

I want to monitor udp-jitter using IP SLA (Cisco Router). The Destination Router was Huawei Router. I'm not sure whether Huawei Router doesn't support IP SLA responder. Is there anyway to use NQA with IPSLA ? 

This is my configuration.

Cisco Router

```
`ip sla 100
udp-jitter 10.10.10.2 65001 
frequency 30
request-datasize 1500

ip sla schedule 100 start now life forever
`

Huawei Router

```
`system-view
nqa-server udpecho 10.10.10.1 65001
`

Kindly suggest,

Thank you.

## Answer 1

By Jota | Score: 1

To enable the router to respond to UDP-Echo or UDP-Jitter packets sent by a third-party device:

ip nqa-compatible responder enable

ip nqa-compatible auto
