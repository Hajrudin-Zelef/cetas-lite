---
id: collect-261001-fortinet/fortinet/document-fortigate-7-2-0-sd-wan-sd-branch-architecture-for-mssps-151899-basic-sd-76069c5b
title: "document-fortigate-7-2-0-sd-wan-sd-branch-architecture-for-mssps-151899-basic-sd-76069c5b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-2-0-sd-wan-sd-branch-architecture-for-mssps-151899-basic-sd-76069c5b.md
source_anchor: ""
source_lines: [1, 7]
sha256: 57c5231c7138c5110d7e804e09b264e675367f4c874d56aa61e883e1e51d54bb
---

# document-fortigate-7-2-0-sd-wan-sd-branch-architecture-for-mssps-151899-basic-sd-76069c5b

Basic SD-WAN/ADVPN design
Basic SD-WAN/ADVPN design
We have already noted that the fundamental building block of our SD-WAN/ADVPN solution is the Hub-and-Spoke overlay topology that securely interconnects the SD-WAN sites:
We call each Hub-and-Spoke block a region. Every region is typically served by either one or two Hubs. Dual-Hub regions are the most recommended, for redundancy reasons:
It is common for the entire SD-WAN network to consist out of a single region, but large-scale deployments (with thousands of sites) will be multi-regional:
In a multi-regional deployment, the Hubs will typically build a Full Mesh between them. ADVPN can be enabled within each region, but it can also stretch across the regions, allowing to build inter-regional shortcuts, both Spoke-to-Spoke (between the Spokes in different regions) and Spoke-to-Hub (towards the Hubs serving other regions), as demonstrated on the above diagram.
Let's see how the SD-WAN nodes are configured:
