---
id: collect-260926-mikrotik/mikrotik/questions-731068-how-to-limit-bandwidth-after-a-threshold-is-reached-mikrotik-b9f72aa2
title: "questions-731068-how-to-limit-bandwidth-after-a-threshold-is-reached-mikrotik-b9f72aa2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/questions-731068-how-to-limit-bandwidth-after-a-threshold-is-reached-mikrotik-b9f72aa2.md
source_anchor: ""
source_lines: [1, 5]
sha256: 194b239dfd61d9c2bca2f3dbb6ffb350fb99d96de21f3a5c8502f214fd5e7c4f
---

# questions-731068-how-to-limit-bandwidth-after-a-threshold-is-reached-mikrotik-b9f72aa2

My router is Mikrotik RB951G-2HnD and I would like it to limit the connection speed to a particular website once some amount of traffic between the user and the website has passed. The website ip range is known, the user is set up by dhcp.
- There is a connection-bytes switch for/ip firewall mangle rules. I could start marking packets or send users to anaddress-list based on that switch. However, I can't single out the user in this case. Every user connected to the website will contribute to reaching the threshold and then everyone will get a speed limit.
- It looks like the User Manager package might be what I am looking for but it's not preinstalled on the router and I would like to avoid installing extra stuff if at all possible.
- And there is RADIUS which, judging by wikipedia, is pretty hard to understand let alone install. Looks like an overkill.
Is there a simple way to limit the speed of a particular user-to-website connection after a certain threshold has been reached? Am I missing something builtin?
