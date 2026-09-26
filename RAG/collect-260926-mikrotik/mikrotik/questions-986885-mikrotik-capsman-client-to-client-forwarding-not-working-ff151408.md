---
id: collect-260926-mikrotik/mikrotik/questions-986885-mikrotik-capsman-client-to-client-forwarding-not-working-ff151408
title: "questions-986885-mikrotik-capsman-client-to-client-forwarding-not-working-ff151408"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/questions-986885-mikrotik-capsman-client-to-client-forwarding-not-working-ff151408.md
source_anchor: ""
source_lines: [1, 15]
sha256: 9dca5648f66cbff70aa016de5d7a53ff1a78f9f3a7e1422310331dc761a23d49
---

# questions-986885-mikrotik-capsman-client-to-client-forwarding-not-working-ff151408

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to set Mikrotik CAPsMAN to allow wifi users communicate between each other
On CAPs server I have allowed "Local forwarding" and "Client to Client forwarding", pushed configuration but users still can not communicate.
What else is needed to allow communicate users together over wifi?
Connection to DMZ is fine, connection to Internet is fine.
Issue is only while clients try to reach each other by WiFi
CAP is reporting that local forwarding is working. However users in the same CAP and same subnet cannot reach. Only ping works...
On the wireless configuration page, set "Multicast Helper" to "full".
Also make sure the bridge specified in your settings exists on the AP and the traffic is routed appropriately. "Local Forwarding" actually means to use the bridge selected on the AP.
If you're using "Access Lists" make sure individual client settings are not overriding the global CAPsMAN settings.
