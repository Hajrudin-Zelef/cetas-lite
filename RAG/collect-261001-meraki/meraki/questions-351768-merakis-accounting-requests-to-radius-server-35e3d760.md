---
id: collect-261001-meraki/meraki/questions-351768-merakis-accounting-requests-to-radius-server-35e3d760
title: "questions-351768-merakis-accounting-requests-to-radius-server-35e3d760"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "research"]
source: docs/RAG/collect-261001-meraki/questions-351768-merakis-accounting-requests-to-radius-server-35e3d760.md
source_anchor: ""
source_lines: [1, 10]
sha256: cb02da7204e035c1387b18e85f61a27568d4e0a16ce85d6a6f753b9ee7fdbb85
---

# questions-351768-merakis-accounting-requests-to-radius-server-35e3d760

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm running a RADIUS server with some Meraki APs, the process of Authentications is fine... But it seems that the Meraki Cloud Controller is just sending the authentication packets and not the accounting requests. I've tested the RADIUS sending accounting requests with the radclient tool (locally) and it worked.
I think that maybe my RADIUS server is ignoring the accounting requests from the MCC because there are some Vendor Specific Attributes that my RADIUS doesn't know. should I add a Meraki's dictionary to my RADIUS configurations?
run the radsniff utility that is included in the standard set of utilities, to verify that the server is actually receiving the accounting packets (or start the daemon debug mode -X). FreeRADIUS will not ignore well formed VSAs even if it does not have the dictionaries to decode them.
The problem was that the configuration for Accounting in the MCC is hidden by default (there are two separate configuration: Auth and Accounting) . You have to talk to Meraki's support and ask them to enable it.
