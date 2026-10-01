---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1803186-how-can-i-block-a-printer-from-guest-users-on-a-ubiquiti-netwo-1196a9ec
title: "How can I block a printer from guest users on a Ubiquiti network?"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1803186-how-can-i-block-a-printer-from-guest-users-on-a-ubiquiti-netwo-1196a9ec.md
source_anchor: ""
source_lines: [1, 13]
sha256: 7fd4e863bf8fa2b676c5a0541c829b4b01d026c0cca189c9c6fffcfe1d84e326
---

# How can I block a printer from guest users on a Ubiquiti network?

*Score : 1 | Source : https://superuser.com/questions/1803186/how-can-i-block-a-printer-from-guest-users-on-a-ubiquiti-network*

I'm trying to block a printer on a Ubiquiti network with 3 LANs: one is Public or Guest(open) and two are p/w protected.
Will a network rule work or should I use vlans? Having trouble with network rules.

---

### Reponse — score 1

I think separate vlans will work.
I am also using ubiquiti and I use the simpler approach of using a different subnet for guest connections. That way its robust and I dont have to keep protecting additional devices.
