---
id: collect-261001-huawei/huawei/knowledge-base-design-and-configure-layer-3-subinterface-configuration-huawei-2a5d5d91
title: "knowledge-base-design-and-configure-layer-3-subinterface-configuration-huawei-2a5d5d91"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "United States"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/knowledge-base-design-and-configure-layer-3-subinterface-configuration-huawei-2a5d5d91.md
source_anchor: ""
source_lines: [1, 35]
sha256: f54c74d819026dd73d430d40f21f4a36dc646c3a9d450b9a7310f3ec92453f14
---

# knowledge-base-design-and-configure-layer-3-subinterface-configuration-huawei-2a5d5d91

Poland
GRANDMETRIC Sp. z o.o.
ul. Metalowa 5, 60-118 Poznań, Poland
NIP 7792433527
+48 61 271 04 43
info@grandmetric.com
UK
Grandmetric LTD
Office 584b
182-184 High Street North
London
E6 2JA
+44 20 3321 5276
info@grandmetric.com
US Region
Grandmetric LLC
Lewes DE 19958
16192 Coastal Hwy USA
EIN: 98-1615498
+1 302 691 94 10 
info@grandmetric.com
Technology: Network
Area: Configuration General
Vendor: Huawei
Software: eNSP
Platform: AR120, S9300&S9300E&S9300X V200R010C00
Sub-interfaces are multiple logical interfaces configured on a main (physical) interface to allow to communicate within subnets on a trunk link. Sub-interfaces can share physical layer parameters of their main interface or be configured with their respective link layer parameters and network layer parameters. Disabling or activating sub-interfaces does not affect the main interface status, but the main interface status change affects the status of sub-interfaces. Sub-interfaces work properly only when their main interface is in Up-state. Associating a sub-interface with a VLAN implements inter-VLAN communication and applies to Dot1q termination and QinQ termination scenarios.
To enter the System configuration mode run:
system-view
After System-View enter the subinterface mode:
interface interface-type interface-number.subinterface-number
To configure the IP address:
ip address ip-address { mask | mask-length } [ subnet ]
To configure the sub-interface for dot1q VLAN Run:
dot1q termination vid low-pe-vid [ to high-pe-vid ]
