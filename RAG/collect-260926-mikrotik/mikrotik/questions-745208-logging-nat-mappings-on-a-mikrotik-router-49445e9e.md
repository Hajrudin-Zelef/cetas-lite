---
id: collect-260926-mikrotik/mikrotik/questions-745208-logging-nat-mappings-on-a-mikrotik-router-49445e9e
title: "questions-745208-logging-nat-mappings-on-a-mikrotik-router-49445e9e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-745208-logging-nat-mappings-on-a-mikrotik-router-49445e9e.md
source_anchor: ""
source_lines: [1, 12]
sha256: 770f0d2fa5a488c4001abcdd45040246c9ef261f0f909e3df87bdc5f63700735
---

# questions-745208-logging-nat-mappings-on-a-mikrotik-router-49445e9e

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am looking for suggestions on how to configure a Mikrotik router to log NAT session setup and ideally teardown. However, setup only is sufficient for my immediate needs.
I found a method to log all packets, and I might be able to reduce this to only logging the TCP SYN, FIN, and RST packets, but I have not yet found that method. This doesn't cover UDP, which I would also like to track. For UDP, it's fine if a session times out and re-appears as a new session again.
This answer has been awarded bounties worth 50 reputation by Michael Graff
Show activity on this post.
You haven't mentioned which method you have found, as standard logging function in RouterOS will not log such info.
These information are only available through Connection Tracking, if you meant that, and since Mikrotik doesn't offer any option to export it fully or to send it to remote syslog host, so I guess, there is very little chance to implement what you are trying to.
