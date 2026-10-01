---
id: collect-261001-meraki/meraki/questions-735664-meraki-switch-uplink-1bf56fd8
title: "questions-735664-meraki-switch-uplink-1bf56fd8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-735664-meraki-switch-uplink-1bf56fd8.md
source_anchor: ""
source_lines: [1, 7]
sha256: 823cec92b65dcbbda727b0e1c403086fd03de135da081e69da8f3824340b6e27
---

# questions-735664-meraki-switch-uplink-1bf56fd8

I have a new switch from Cisco Meraki (MS220-8P) and i am trying to configure a Dual Up-link to a 48 port Catalyst 3750 (which is provided by our phone & IS provider). The Catalyst 3750 has been configured to receive the dual up-link but I cannot find where to configure a dual up-link in the Meraki Control Panel.
2 Answers 2
Link Aggregation is configurable on the Switch Ports (Switch -> Monitor -> Switch Ports) page. From my experience the ports have to be unused when configuring them.
Select two (or up to 8) and the "aggregation" button appears. Click to configure It uses LACP to perform aggregation.
To make this the up-link all you do is connect the up link to these port and it will detect and assign the priority automatically.
For Meraki, select the ports to be aggregated by checking their respective boxes (under Switching > Monitor > Switch Ports page) and then select the Aggregate option at the top. this will create an LACP port group running mode: active.
Port-Channel,LACP,Etherchannel,bondingorlink aggregation. But you can connect two cables in an active/standby failover design withSpanning-Tree, or two cables, each running different VLANs, which would needVLANconfiguration on the individual ports.
