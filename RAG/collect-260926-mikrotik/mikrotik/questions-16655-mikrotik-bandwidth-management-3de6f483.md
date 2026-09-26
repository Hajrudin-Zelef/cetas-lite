---
id: collect-260926-mikrotik/mikrotik/questions-16655-mikrotik-bandwidth-management-3de6f483
title: "questions-16655-mikrotik-bandwidth-management-3de6f483"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-16655-mikrotik-bandwidth-management-3de6f483.md
source_anchor: ""
source_lines: [1, 16]
sha256: 47aa01470ff8e080eaa4527676a546c8acef4f509b275f57a135708a502f46cf
---

# questions-16655-mikrotik-bandwidth-management-3de6f483

The only real control over traffic that you have is traffic on your network, or the traffic leaving your network. You don't have any real control over how fast an outside server sends traffic to your network. That's why DoS attacks work. By the time the traffic reaches your network the WAN bandwidth is already used.
There are ways to try to mitigate this:
- As JoeriBe pointed out, MikroTik has a was to prevent the incoming
traffic from being delivered to a user, but that doesn't directly
prevent the WAN bandwidth from be monopolized by the incoming
traffic, it just prevents the user from receiving all of it. This may
or may not have any effect. If the incoming traffic is TCP, it can
have the effect of changing the TCP window size, and that can slow
down the server to a degree. If the traffic is UDP, it may not
have any effect at all, except to create problems for the user.
- You can also come to an agreement with your ISP to limit incoming
traffic from a single source to a certain percentage of your WAN
bandwidth. This way, the traffic from a single source cannot
overwhelm your incoming bandwidth. This is the only sure way, but it
may cost some money, and you have to decide if it is worth it to your
business.
