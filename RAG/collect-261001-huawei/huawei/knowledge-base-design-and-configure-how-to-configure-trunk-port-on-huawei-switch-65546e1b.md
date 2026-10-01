---
id: collect-261001-huawei/huawei/knowledge-base-design-and-configure-how-to-configure-trunk-port-on-huawei-switch-65546e1b
title: "knowledge-base-design-and-configure-how-to-configure-trunk-port-on-huawei-switch-65546e1b"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "United States"]
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-huawei/knowledge-base-design-and-configure-how-to-configure-trunk-port-on-huawei-switch-65546e1b.md
source_anchor: ""
source_lines: [1, 34]
sha256: e42f31c78c4246b2160ac3f6301817991f18f72ff3e29db213e3151fbaa84cfb
---

# knowledge-base-design-and-configure-how-to-configure-trunk-port-on-huawei-switch-65546e1b

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
Technology: Switching
Area: Neighbor Discovery
Vendor: Huawei
Software: eNSP
Platform: Quidway switches
A trunk interface often connects to a switch, router, AP, or voice terminal that can receive and send tagged and untagged frames simultaneously. It allows tagged frames from multiple VLANs and untagged frames from only one VLAN called native vlan.
To configure ethernet port as a trunk port and allow to pass all vlans:
<HUAWEI> system-view
[HUAWEI] interface GigabitEthernet 0/0/13
[HUAWEI-GigabitEthernet 0/0/13]port link-type trunk
[HUAWEI-GigabitEthernet 0/0/13]port trunk allow-pass vlan all
[HUAWEI-GigabitEthernet 0/0/13]quit
Check how to configure access port on Huawei switch?
