---
id: collect-261001-general-networking/general-networking/manual-etpro-telemetry-html-5f853fc3-2
title: "manual-etpro-telemetry-html-5f853fc3"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-etpro-telemetry-html-5f853fc3.md
source_anchor: ""
source_lines: [100, 106]
sha256: 80b4c07d6f9560c9254059795857ad4757a7359616d91cb2250ee97bcb294713
---

# manual-etpro-telemetry-html-5f853fc3

| Suricata status | Reports if the sensor is active, when not active, no detection/telemetry can be provided. | 
| System Time | If the system time is not correct, it will impact the timestamps of messages, so knowing what time the system thinks it has will help reconcile the actual time. | 
| Active Ruleset Version | The active ruleset version should match what is published. If sensors do not have the active version then they either haven’t configured scheduled updates or there is another issue. This will help Proofpoint to identify if there are widespread issues with updates. | 
| Number of rules enabled | Helps to gain a better understanding about the number of rules people use on top of the ones provided by Proofpoint. | 
| Number of ETPro Telemetry Rules enabled | Because users can control what rules they enable, they may not want to enable all ETPro Telemetry rules, if this is the case it would help Proofpoint understand how the rules are being leveraged so they can better write / tune rules | 
| Mode (IDS or IPS) | This is helpful to understand how the system is deployed and is useful to development purposes to determine what rules we should be focusing on based on how our customers are using them. | 
| Suricata Log Stats | For QA purposes, some fields with general stats are collected from /var/log/suricata/stats.log (capture.kernel_packets, decoder.pkts, decoder.bytes, decoder.ipv4, decoder.ipv6, flow.tcp, flow.udp, detect.alert) |
