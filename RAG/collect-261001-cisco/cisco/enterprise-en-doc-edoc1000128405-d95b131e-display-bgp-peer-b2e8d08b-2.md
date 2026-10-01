---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1000128405-d95b131e-display-bgp-peer-b2e8d08b-2
title: "Display peer information."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1000128405-d95b131e-display-bgp-peer-b2e8d08b.md
source_anchor: ""
source_lines: [125, 149]
sha256: dad69303b27e46993a5fb69e6888459c332666acfd58aed264d469a9fa591a2f
---

# Display peer information.

| Minimum route advertisement interval is 15 seconds | Indicates the minimum interval between route advertisements.  | 
| Optional capabilities | (Optional) Indicates the peer-supported capabilities. | 
| Route refresh capability has been enabled | Indicates that route refreshing has been enabled. | 
| 4-byte-as capability has been enabled | 4-byte-As is enabled | 
| Listen-only has been configured | Indicates that only connection requests are snooped and no connections will be initiated proactively. | 
| Peer Preferred Value | Indicates the preferred value of the peer. | 
| Routing policy configured | Indicates the configured routing policy. | 
| Peer's BFD has been enabled | Indicates that BFP has been enabled on the peer. | 
<HUAWEI> display bgp peer 10.1.1.2 log-info
Peer : 10.1.1.2 
 Date/Time     : 2011/13/06 11:53:21
 State         : Up
 Date/Time     : 2011/13/06 11:53:09
 State         : Down
 Error Code    : 6(CEASE)
 Error Subcode : 4(Administrative Reset)
 Notification  : Receive Notification
 Date/Time     : 2011/13/06 10:34:05
 State         : Up
| Item | Description | 
|---|---|
| Error Code | Error code | 
| Error Subcode | Error subcode | 
| Notification | Notification packet sent or received by a peer | 
Select the content with the mouse pointer to quickly report the problem.
