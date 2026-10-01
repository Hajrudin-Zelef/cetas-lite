---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-342836-sd-wan-rules-lowest-cost-sl-a2b4492a-2
title: "document-fortigate-7-4-3-administration-guide-342836-sd-wan-rules-lowest-cost-sl-a2b4492a"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-342836-sd-wan-rules-lowest-cost-sl-a2b4492a.md
source_anchor: ""
source_lines: [142, 155]
sha256: 746aa40c8a147cbeb1dfcf91ad536e4574c30d2a4fe932845296871fde64e6eb
---

# document-fortigate-7-4-3-administration-guide-342836-sd-wan-rules-lowest-cost-sl-a2b4492a

FGT # diagnose sys sdwan health-check status
Health Check(google):
Seq(1): state(alive), packet-loss(0.000%) latency(14.563), jitter(4.334) sla_map=0x0
Seq(2): state(alive), packet-loss(0.000%) latency(12.633), jitter(6.265) sla_map=0x0
FGT # diagnose sys sdwan service 1
Service(1): Address Mode(IPV4) flags=0x0
    TOS(0x0/0x0), Protocol(0: 1->65535), Mode(load-balance)
    Members:<<BR>>
        1: Seq_num(1), alive, sla(0x1), num of pass(1), selected
        2: Seq_num(2), alive, sla(0x1), num of pass(1), selected
    Internet Service: Google.Gmail(65646)
                                            When both wan1 and wan2 meet the SLA requirements, Gmail traffic will use both wan1 and wan2. If only one of the interfaces meets the SLA requirements, Gmail traffic will only use that interface.
If neither interface meets the requirements but the health-check is still alive, then wan1 and wan2 tie. The traffic will try to balance between wan1 and wan2, using both interfaces to forward traffic.
|  | The maximize bandwidth (load-balance ) strategy used  prior to FortiOS 7.4.1 is now known as the load balancing strategy. This strategy can be configured under the manual mode and the lowest cost (SLA) strategies.  |
