---
id: collect-260926-mikrotik/mikrotik/mikrotik-failover-routing-description-living-document-3
title: "Static routes"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-failover-routing-description-living-document.md
source_anchor: ""
source_lines: [283, 286]
sha256: 6c0918683847ec3a3ddf5fc19b8e44293bb16fb2b1ab2ab9757f0c287e0a454c
---

# Static routes

- Notification/reset scripts in case of link failover.
- Routing for certain connections that will always go out via a given link, regardless of the failover state.

Great explanation, thank you. This works perfectly for me. My Primary is PPPOE to a WISP and backup via wlan station to my phone hotspot (DHCP). I implemented it at ROS 6.49 and was hoping it would all be translated when I upgraded to ROS 7.12. It wasn't. Is there an updated config that will work for Ros 7?
