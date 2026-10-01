---
id: collect-261001-fortinet/fortinet/questions-452521-how-to-delete-ipsec-vpn-tunnel-from-fortigate-60-8b276ad2
title: "questions-452521-how-to-delete-ipsec-vpn-tunnel-from-fortigate-60-8b276ad2"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-452521-how-to-delete-ipsec-vpn-tunnel-from-fortigate-60-8b276ad2.md
source_anchor: ""
source_lines: [1, 13]
sha256: e482a70932e8c6d466863125f14781e0c8bc94fe2aca1eb0f9ddca1c89b2755c
---

# questions-452521-how-to-delete-ipsec-vpn-tunnel-from-fortigate-60-8b276ad2

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have had a IPSEC connection setup between two firewalls. Now I want to remove the tunnel in my firewall, a "Fortigate 60".
There are two phases, "Phase 1" and "Phase 2" for each IPSEC connection. I can delete the "Phase 2" entry by clicking the trashcan icon (in the web interface), but there is not such icon for "Phase 1". Is it possible to delete that?
When I look at the log it alerts about this tunnel not working (after deleting "Phase 2") and it would be nice not ta have loads of such events in the log.
Whenever you can't delete something in the FortiGate, there usually is a reference to that object somewhere. Normally the references are easy to track as they appear on the UI adjacent to the object.
Two notable exceptions:
* A bug did exist in the past (an old version) where the delete button on the phase object wouldn't work
* Chrome browser doesn't always refresh that condition (not sure if a Chrome or Fortinet issue). Refreshing the sessions; close/reopen; empty cache or whatever works for you.
