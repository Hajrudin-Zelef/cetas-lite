---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-operate-and-maintain-how-tos-allowing-apple-icloud-services-o-0429ab4a
title: "sase-and-sd-wan-mx-operate-and-maintain-how-tos-allowing-apple-icloud-services-o-0429ab4a"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-operate-and-maintain-how-tos-allowing-apple-icloud-services-o-0429ab4a.md
source_anchor: ""
source_lines: [1, 25]
sha256: 44aad9d0e384752f42d73d6b2551bc7f34cbdddf624a7ec93b167ad4e5621ad8
---

# sase-and-sd-wan-mx-operate-and-maintain-how-tos-allowing-apple-icloud-services-o-0429ab4a

How to Allow Apple iCloud on a Restrictive Firewall
Introduction
When using restrictive Layer 3 firewall rules for outbound traffic on the Cisco Meraki MX WAN appliance, services such as Apple iCloud can be inadvertently blocked. This article discusses the ports required for Apple iCloud to perform backups and access data stored in iCloud.
Step-by-step instructions
- Navigate to Security & SD-WAN > Configure > Firewall and review your current firewall configuration, in the Meraki dashboard.
- Evaluate the Layer 3 > Outbound rules. A restrictive rule set, as pictured below, can block iCloud traffic and many features will not function. For example: 
 a. Firewall configuration in Figure 1 blocks all outbound traffic except TPC 80 and TCP 443 (http and https).
b. Users attempting to use other protocols (such as UDP) or other ports (such as 25) will be blocked by the firewall.
Figure 1. Explicit deny rule blocking iCloud traffic.
- 
    To allow iCloud to function, Apple has a list of ports which need to be allowed for iCloud to function on your client devices. The ports used by iCloud are:
- 
    TCP 25
- 
    TCP 80
- 
    TCP 443
- 
    TCP 587
- 
    TCP 993
- 
    TCP 5223
Configure your firewall to allow connections on these specific ports.
Figure 2. iCloud firewall config.
