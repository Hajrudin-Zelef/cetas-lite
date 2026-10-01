---
id: collect-261001-general-networking/general-networking/t5-support-forum-advpn-sd-wan-shortcut-tunnel-underlay-routing-td-p-255968-b51a9f25
title: "t5-support-forum-advpn-sd-wan-shortcut-tunnel-underlay-routing-td-p-255968-b51a9f25"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-support-forum-advpn-sd-wan-shortcut-tunnel-underlay-routing-td-p-255968-b51a9f25.md
source_anchor: ""
source_lines: [1, 4]
sha256: fe8cb4ff75cf2a1b49909da46e3372c5de1043452e3fa06f592ae14e216555e6
---

# t5-support-forum-advpn-sd-wan-shortcut-tunnel-underlay-routing-td-p-255968-b51a9f25

ADVPN SD-WAN Shortcut Tunnel (underlay routing)
When an ADVPN shortcut is created, obviosuly the connection will go from whatever ISP(s) you're using to the other Spokes Public IP(s). This creates the transit connection for the tunnel to be created in the first place.
So, do you have to configure a default route for each ISP/Connection that you use, and then how does that work if you receive a default over a tunnel interface with ADVPN as a default route?
I can always create a static route for a Public IP of the other end of the shortcut tunnel, but that seems a bit tedious and not the best way to do things imo, considering if IPs and NAT'd or change, you can't really keep up with that.
