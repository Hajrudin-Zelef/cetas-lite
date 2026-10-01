---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4.md
source_anchor: ""
source_lines: [713, 783]
sha256: ecd60a56a9d8184fb06e80c117c7e396fca798ddfbe6b4ab5aa2485279146b77
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-5cfddbe4

summary [entry-number] |
neighbors}
Displays MPLS label switched
path (LSP) Health Monitor operations.
show ip sla
reaction-configuration [entry-number]
Displays the configured
proactive threshold monitoring settings for all IP SLA operations or a specific
operation.
show ip sla
reaction-trigger [entry-number]
Displays the reaction trigger
information for all IP SLA operations or a specific operation.
show ip sla
responder
Displays information about
the IP SLA responder.
show ip sla statistics
[entry-number |
aggregated |
details]
Displays current or
aggregated operational status and statistics.
Monitoring IP SLA Operation Examples
The following example shows all IP SLAs by application:
Device# show ip sla application
IP Service Level Agreements
Version: Round Trip Time MIB 2.2.0, Infrastructure Engine-III
Supported Operation Types:
icmpEcho, path-echo, path-jitter, udpEcho, tcpConnect, http
dns, udpJitter, dhcp, ftp, udpApp, wspApp
Supported Features:
IPSLAs Event Publisher
IP SLAs low memory water mark: 33299323
Estimated system max number of entries: 24389
Estimated number of configurable operations: 24389
Number of Entries configured : 0
Number of active Entries : 0
Number of pending Entries : 0
Number of inactive Entries : 0
Time of last change in whole IP SLAs: *13:04:37.668 UTC Wed Dec 19 2012
The following example shows all IP SLA distribution statistics:
Device# show ip sla enhanced-history distribution-statistics
Point by point Enhanced History
Entry = Entry Number
Int = Aggregation Interval
BucI = Bucket Index
StartT = Aggregation Start Time
Pth = Path index
Hop = Hop in path index
Comps = Operations completed
OvrTh = Operations completed over thresholds
SumCmp = Sum of RTT (milliseconds)
SumCmp2L = Sum of RTT squared low 32 bits (milliseconds)
SumCmp2H = Sum of RTT squared high 32 bits (milliseconds)
TMax = RTT maximum (milliseconds)
TMin = RTT minimum (milliseconds)
Entry Int BucI StartT Pth Hop Comps OvrTh SumCmp SumCmp2L SumCmp2H T
Max TMin
The Cisco
Support website provides extensive online resources, including documentation
and tools for troubleshooting and resolving technical issues with Cisco
products and technologies.
To receive
security and technical information about your products, you can subscribe to
various services, such as the Product Alert Tool (accessed from Field Notices),
the Cisco Technical Services Newsletter, and Really Simple Syndication (RSS)
Feeds.
Access to
most tools on the Cisco Support website requires a Cisco.com user ID and
password.
