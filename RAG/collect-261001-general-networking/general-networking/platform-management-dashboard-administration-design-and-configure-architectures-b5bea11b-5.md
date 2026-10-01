---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b5bea11b-5
title: "platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b.md
source_anchor: ""
source_lines: [192, 198]
sha256: 77b192bc9c1cef448724c5905ad211fd39bd5dfb23297986f56c9b03ef2e456b
---

# platform-management-dashboard-administration-design-and-configure-architectures--b5bea11b

Use traffic shaping to offer application traffic the necessary bandwidth. It is important to ensure that the application has enough bandwidth as estimated in the capacity planning section. Traffic shaping rules can be implemented to allow real-time voice and video traffic to use additional bandwidth, and the rules can be used to block or throttle applications such as P2P, social networks.
- Go to Wireless > Configure > Firewall & traffic shaping and choose the SSID from the SSID drop-down menu at the top of the screen.
- Click the drop down menu next to Shape traffic and choose Shape traffic on this SSID, then click Create a new rule.
- Click Add + and select 'All voice & video conferencing'
- Set Per-client bandwidth limit to 'Ignore SSID per-client limit (unlimited)' and click Save changes.
Convert Multicast to Unicast
Cisco Meraki APs automatically perform a multicast-to-unicast packet conversion using the IGMP protocol. The unicast frames are then sent at the client negotiated data rates rather than the minimum mandatory data rates, ensuring high-quality video transmission to large numbers of clients. This can be especially valuables in instances such as classrooms, where multiple students may be watching a high-definition video as part a classroom learning experience.
