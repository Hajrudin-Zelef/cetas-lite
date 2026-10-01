---
id: collect-261001-meraki/meraki/questions-54862-meraki-mx84-vpn-question-2331e9ec
title: "questions-54862-meraki-mx84-vpn-question-2331e9ec"
domain: meraki
role: reference
task: reference
actors: []
dates: ["2018-12-25"]
keywords: []
source: docs/RAG/collect-261001-meraki/questions-54862-meraki-mx84-vpn-question-2331e9ec.md
source_anchor: ""
source_lines: [1, 9]
sha256: 4a9b2ea814785d142f67dfe493c945f1ea6e2532a36daab5aeb52e954e62de5f
---

# questions-54862-meraki-mx84-vpn-question-2331e9ec

I'm looking to replace our Cisco 2900 router, and one of the choices is the MX84. My question is that can I use the AnyConnect VPN client for the MX84 or do I need a Cisco ASA?
- 
        2I removed the question for opinion based answers to make the question on topic here.Teun Vink– Teun Vink2018-11-20 14:06:36 +00:00Commented Nov 20, 2018 at 14:06
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can provide and accept your own answer.Ron Maupin– Ron Maupin ♦2018-12-25 10:04:20 +00:00Commented Dec 25, 2018 at 10:04
2 Answers 2
Meraki MX routers do not support Cisco Anyconnect. If Anyconnect is a requirement I'm afraid you'll have to stick with Cisco branded routers.
Not currently supported.
This is likely due to the fact that AnyConnect establishes an IPSec Tunnel based on IKE v2 whereas the MX appliance (afaik) still uses IKE v1 for negotation.
